// Command progress runs the exercise tests and shows how many exercises you
// have solved in each topic.
//
//	go run ./cmd/progress              # summary of every topic
//	go run ./cmd/progress 03_loops     # exercise-by-exercise detail for one topic
//	go run ./cmd/progress -solutions   # run against the reference solutions
//
// A topic is either a single package (exercises/03_loops) or a folder of
// exercise packages (exercises/08_structs/01_person, ...).
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"slices"
	"strings"
)

const exercisesDir = "exercises"

// testRef identifies a top-level test. pkg is the exercise folder inside the
// topic, or "" when the topic is a single package.
type testRef struct {
	pkg  string
	name string
}

type pkgResult struct {
	buildFailed bool
	panicked    bool
	panicFrame  bool
	output      []string // compiler errors or panic details
}

type topic struct {
	name   string
	nested bool               // exercises live in sub-packages
	tests  []testRef          // in source order
	status map[testRef]string // "pass" / "fail" / "skip"
	pkgs   map[string]*pkgResult
}

func (t *topic) pkg(name string) *pkgResult {
	r := t.pkgs[name]
	if r == nil {
		r = &pkgResult{}
		t.pkgs[name] = r
	}
	return r
}

func (t *topic) passed() int {
	n := 0
	for _, ref := range t.tests {
		if t.status[ref] == "pass" {
			n++
		}
	}
	return n
}

// pkgNames lists the packages holding tests, in source order.
func (t *topic) pkgNames() []string {
	var names []string
	for _, ref := range t.tests {
		if !slices.Contains(names, ref.pkg) {
			names = append(names, ref.pkg)
		}
	}
	return names
}

func (t *topic) anyPanicked() bool {
	for _, r := range t.pkgs {
		if r.panicked {
			return true
		}
	}
	return false
}

func (t *topic) dir(pkg string) string {
	return path.Join(exercisesDir, t.name, pkg)
}

func (t *topic) file(pkg string) string {
	if pkg == "" {
		return path.Join(t.dir(pkg), "exercises.go")
	}
	return path.Join(t.dir(pkg), "exercise.go")
}

// testEvent is one line of `go test -json` output.
type testEvent struct {
	Action      string
	Package     string
	ImportPath  string
	Test        string
	Output      string
	FailedBuild string
}

func main() {
	solutions := flag.Bool("solutions", false, "run the tests against the reference solutions")
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "usage: go run ./cmd/progress [-solutions] [topic]")
		flag.PrintDefaults()
	}
	flag.Parse()

	only := strings.TrimSuffix(strings.TrimPrefix(flag.Arg(0), exercisesDir+"/"), "/")

	topics, err := discoverTopics(only)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	if len(topics) == 0 {
		fmt.Fprintf(os.Stderr, "no topic named %q in %s/\n", only, exercisesDir)
		os.Exit(1)
	}

	if err := runTests(topics, *solutions); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	if only != "" {
		printDetail(topics[0])
	} else {
		printSummary(topics)
	}
}

// discoverTopics lists the exercise packages and statically extracts their
// test names, so exercises that never ran (because of a compile error or a
// panic) are still counted.
func discoverTopics(only string) ([]*topic, error) {
	entries, err := os.ReadDir(exercisesDir)
	if err != nil {
		return nil, fmt.Errorf("run this command from the repository root: %w", err)
	}
	var topics []*topic
	for _, e := range entries {
		if !e.IsDir() || (only != "" && e.Name() != only) {
			continue
		}
		dir := filepath.Join(exercisesDir, e.Name())
		flat, _ := filepath.Glob(filepath.Join(dir, "*_test.go"))
		nested, _ := filepath.Glob(filepath.Join(dir, "*", "*_test.go"))
		files := slices.Concat(flat, nested)
		slices.Sort(files)
		t := &topic{
			name:   e.Name(),
			nested: len(nested) > 0,
			status: map[testRef]string{},
			pkgs:   map[string]*pkgResult{},
		}
		fset := token.NewFileSet()
		for _, f := range files {
			pkg, err := filepath.Rel(dir, filepath.Dir(f))
			if err != nil {
				return nil, err
			}
			if pkg == "." {
				pkg = ""
			}
			file, err := parser.ParseFile(fset, f, nil, parser.SkipObjectResolution)
			if err != nil {
				return nil, err
			}
			for _, decl := range file.Decls {
				fn, ok := decl.(*ast.FuncDecl)
				if ok && fn.Recv == nil && strings.HasPrefix(fn.Name.Name, "Test") && fn.Name.Name != "TestMain" {
					t.tests = append(t.tests, testRef{pkg: filepath.ToSlash(pkg), name: fn.Name.Name})
				}
			}
		}
		topics = append(topics, t)
	}
	return topics, nil
}

func runTests(topics []*topic, solutions bool) error {
	byName := map[string]*topic{}
	args := []string{"test", "-json", "-count=1"}
	if solutions {
		args = append(args, "-tags", "solutions")
	}
	for _, t := range topics {
		byName[t.name] = t
		args = append(args, "./"+exercisesDir+"/"+t.name+"/...")
	}

	cmd := exec.Command("go", args...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}

	scanner := bufio.NewScanner(stdout)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
	for scanner.Scan() {
		var ev testEvent
		if json.Unmarshal(scanner.Bytes(), &ev) != nil {
			continue
		}
		importPath := ev.Package
		if importPath == "" {
			importPath = ev.ImportPath
		}
		topicName, pkg := locate(importPath)
		t := byName[topicName]
		if t == nil {
			continue
		}
		r := t.pkg(pkg)
		switch {
		case ev.Action == "build-output":
			r.output = append(r.output, strings.TrimRight(ev.Output, "\n"))
		case ev.Action == "build-fail" || ev.FailedBuild != "":
			r.buildFailed = true
		case ev.Test != "" && !strings.Contains(ev.Test, "/"):
			switch ev.Action {
			case "pass", "fail", "skip":
				t.status[testRef{pkg: pkg, name: ev.Test}] = ev.Action
			}
		}
		if ev.Action == "output" {
			collectPanic(r, ev)
		}
	}
	// A non-zero exit status just means some tests failed, which is expected.
	_ = cmd.Wait()

	if stderr.Len() > 0 && !anyResults(topics) {
		return fmt.Errorf("go test failed:\n%s", stderr.String())
	}
	return nil
}

// locate splits an import path such as "mod/exercises/08_structs/01_person"
// into its topic ("08_structs") and exercise package ("01_person", or "" for a
// single-package topic). Build output import paths can carry a suffix, like
// "mod/exercises/01_basics [mod/exercises/01_basics.test]".
func locate(importPath string) (topic, pkg string) {
	importPath = strings.Fields(importPath + " ")[0]
	_, rest, ok := strings.Cut(importPath, "/"+exercisesDir+"/")
	if !ok {
		return "", ""
	}
	topic, pkg, _ = strings.Cut(rest, "/")
	return topic, pkg
}

// collectPanic keeps the panic message and the first stack frame located in
// the learner's exercise file, so the summary can point at the faulty line.
func collectPanic(r *pkgResult, ev testEvent) {
	line := strings.TrimSpace(ev.Output)
	switch {
	case strings.HasPrefix(line, "panic: "):
		r.panicked = true
		r.output = append(r.output, line)
	case r.panicked && !r.panicFrame && strings.Contains(line, "/"+exercisesDir+"/") &&
		strings.Contains(line, ".go:") && !strings.Contains(line, "_test.go"):
		r.panicFrame = true
		if i := strings.Index(line, exercisesDir+"/"); i >= 0 {
			line = line[i:]
		}
		r.output = append(r.output, "at "+strings.Fields(line)[0])
	}
}

func anyResults(topics []*topic) bool {
	for _, t := range topics {
		if len(t.status) > 0 {
			return true
		}
		for _, r := range t.pkgs {
			if r.buildFailed {
				return true
			}
		}
	}
	return false
}

func exerciseName(test string) string {
	return strings.TrimPrefix(test, "Test")
}

func bar(done, total, width int) string {
	filled := 0
	if total > 0 {
		filled = done * width / total
	}
	return strings.Repeat("█", filled) + strings.Repeat("░", width-filled)
}

func printSummary(topics []*topic) {
	fmt.Println()
	fmt.Println("  Go Playground: your progress")
	fmt.Println()
	done, total := 0, 0
	var next *topic
	var nextTest testRef
	for _, t := range topics {
		p := t.passed()
		done += p
		total += len(t.tests)
		mark := "  "
		if p == len(t.tests) && len(t.tests) > 0 {
			mark = "✅"
		}
		note := ""
		// In a nested topic, exercises you haven't started don't compile yet,
		// which is expected, so only single-package topics get this warning.
		if r := t.pkgs[""]; !t.nested && r != nil && r.buildFailed {
			note = "  ⚠️  does not compile"
		} else if t.anyPanicked() {
			note = "  💥 a test panicked"
		}
		fmt.Printf("  %s %-18s %s  %3d/%-3d%s\n", mark, t.name, bar(p, len(t.tests), 24), p, len(t.tests), note)
		if next == nil {
			for _, ref := range t.tests {
				if t.status[ref] != "pass" {
					next, nextTest = t, ref
					break
				}
			}
		}
	}
	pct := 0
	if total > 0 {
		pct = done * 100 / total
	}
	fmt.Println()
	fmt.Printf("  Total: %d/%d exercises solved (%d%%)\n", done, total, pct)
	if next != nil {
		label := exerciseName(nextTest.name)
		if nextTest.pkg != "" {
			label = nextTest.pkg + " → " + label
		}
		fmt.Println()
		fmt.Printf("  Next up: %s → %s\n", next.name, label)
		fmt.Printf("    open    %s\n", next.file(nextTest.pkg))
		fmt.Printf("    test    go test ./%s -run '^%s$'\n", next.dir(nextTest.pkg), nextTest.name)
		fmt.Printf("    detail  go run ./cmd/progress %s\n", next.name)
		if r := next.pkgs[nextTest.pkg]; r != nil && (r.buildFailed || r.panicked) {
			printOutput(r.output)
		}
	} else if total > 0 {
		fmt.Println()
		fmt.Println("  🎉 Everything passes. You're a Gopher now!")
	}
	fmt.Println()
}

func printDetail(t *topic) {
	fmt.Println()
	fmt.Printf("  %s: %d/%d solved\n", t.name, t.passed(), len(t.tests))
	if !t.nested {
		fmt.Println()
	}
	indent := "  "
	if t.nested {
		indent = "    "
	}
	pkg, n := "", 0
	for i, ref := range t.tests {
		if t.nested && (i == 0 || ref.pkg != pkg) {
			pkg, n = ref.pkg, 0
			fmt.Printf("\n  %s\n", pkg)
		}
		n++
		icon := "❌"
		switch t.status[ref] {
		case "pass":
			icon = "✅"
		case "skip":
			icon = "⏭️ "
		case "":
			icon = "💥" // never ran: an earlier test probably panicked
			if r := t.pkgs[ref.pkg]; r != nil && r.buildFailed {
				icon = "⬜" // doesn't compile (yet)
			}
		}
		fmt.Printf("%s%s %2d. %s\n", indent, icon, n, exerciseName(ref.name))
	}

	// Only the first package that doesn't compile is reported: in a nested
	// topic, the ones after it are usually exercises you haven't started.
	for _, name := range t.pkgNames() {
		if r := t.pkgs[name]; r != nil && r.buildFailed {
			fmt.Println()
			fmt.Printf("  ⬜ %s does not compile (yet), so its tests could not run:\n", t.dir(name))
			printOutput(r.output)
			break
		}
	}
	for _, name := range t.pkgNames() {
		if r := t.pkgs[name]; r != nil && r.panicked {
			fmt.Println()
			fmt.Printf("  💥 A test panicked in %s, so the tests after it did not run:\n", t.dir(name))
			printOutput(r.output)
		}
	}

	target := t.dir("")
	if t.nested {
		target += "/..."
	}
	fmt.Println()
	fmt.Printf("  See failure details with: go test ./%s\n\n", target)
}

func printOutput(lines []string) {
	fmt.Println()
	for _, line := range lines {
		if line != "" {
			fmt.Println("    " + line)
		}
	}
}
