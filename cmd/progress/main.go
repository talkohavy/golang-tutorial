// Command progress runs the exercise tests and shows how many exercises you
// have solved in each topic.
//
//	go run ./cmd/progress              # summary of every topic
//	go run ./cmd/progress 03_loops     # exercise-by-exercise detail for one topic
//	go run ./cmd/progress -solutions   # run against the reference solutions
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

type topic struct {
	name        string
	tests       []string          // top-level test names, in source order
	status      map[string]string // test name → "pass" / "fail" / "skip"
	buildFailed bool
	panicked    bool
	panicFrame  bool
	output      []string // compiler errors or panic details
}

func (t *topic) passed() int {
	n := 0
	for _, name := range t.tests {
		if t.status[name] == "pass" {
			n++
		}
	}
	return n
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
		files, _ := filepath.Glob(filepath.Join(exercisesDir, e.Name(), "*_test.go"))
		slices.Sort(files)
		t := &topic{name: e.Name(), status: map[string]string{}}
		fset := token.NewFileSet()
		for _, f := range files {
			file, err := parser.ParseFile(fset, f, nil, parser.SkipObjectResolution)
			if err != nil {
				return nil, err
			}
			for _, decl := range file.Decls {
				fn, ok := decl.(*ast.FuncDecl)
				if ok && fn.Recv == nil && strings.HasPrefix(fn.Name.Name, "Test") && fn.Name.Name != "TestMain" {
					t.tests = append(t.tests, fn.Name.Name)
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
		args = append(args, "./"+exercisesDir+"/"+t.name)
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
		pkg := ev.Package
		if pkg == "" {
			pkg = ev.ImportPath
		}
		// ImportPath for build output can look like "mod/exercises/01_basics [mod/exercises/01_basics.test]".
		t := byName[path.Base(strings.Fields(pkg + " ")[0])]
		if t == nil {
			continue
		}
		switch {
		case ev.Action == "build-output":
			t.output = append(t.output, strings.TrimRight(ev.Output, "\n"))
		case ev.Action == "build-fail" || ev.FailedBuild != "":
			t.buildFailed = true
		case ev.Test != "" && !strings.Contains(ev.Test, "/"):
			switch ev.Action {
			case "pass", "fail", "skip":
				t.status[ev.Test] = ev.Action
			}
		}
		if ev.Action == "output" {
			collectPanic(t, ev)
		}
	}
	// A non-zero exit status just means some tests failed, which is expected.
	_ = cmd.Wait()

	if stderr.Len() > 0 && !anyResults(topics) {
		return fmt.Errorf("go test failed:\n%s", stderr.String())
	}
	return nil
}

// collectPanic keeps the panic message and the first stack frame located in
// the learner's exercise file, so the summary can point at the faulty line.
func collectPanic(t *topic, ev testEvent) {
	line := strings.TrimSpace(ev.Output)
	switch {
	case strings.HasPrefix(line, "panic: "):
		t.panicked = true
		t.output = append(t.output, line)
	case t.panicked && !t.panicFrame && strings.Contains(line, "/"+exercisesDir+"/") &&
		strings.Contains(line, ".go:") && !strings.Contains(line, "_test.go"):
		t.panicFrame = true
		if i := strings.Index(line, exercisesDir+"/"); i >= 0 {
			line = line[i:]
		}
		t.output = append(t.output, "at "+strings.Fields(line)[0])
	}
}

func anyResults(topics []*topic) bool {
	for _, t := range topics {
		if len(t.status) > 0 || t.buildFailed {
			return true
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
	var nextTest string
	for _, t := range topics {
		p := t.passed()
		done += p
		total += len(t.tests)
		mark := "  "
		if p == len(t.tests) && len(t.tests) > 0 {
			mark = "✅"
		}
		note := ""
		if t.buildFailed {
			note = "  ⚠️  does not compile"
		} else if t.panicked {
			note = "  💥 a test panicked"
		}
		fmt.Printf("  %s %-18s %s  %3d/%-3d%s\n", mark, t.name, bar(p, len(t.tests), 24), p, len(t.tests), note)
		if next == nil {
			for _, name := range t.tests {
				if t.status[name] != "pass" {
					next, nextTest = t, name
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
		fmt.Println()
		fmt.Printf("  Next up: %s → %s\n", next.name, exerciseName(nextTest))
		fmt.Printf("    open    %s/%s/exercises.go\n", exercisesDir, next.name)
		fmt.Printf("    test    go test ./%s/%s -run '^%s$'\n", exercisesDir, next.name, nextTest)
		fmt.Printf("    detail  go run ./cmd/progress %s\n", next.name)
		if next.buildFailed || next.panicked {
			printOutput(next)
		}
	} else if total > 0 {
		fmt.Println()
		fmt.Println("  🎉 Everything passes. You're a Gopher now!")
	}
	fmt.Println()
}

func printDetail(t *topic) {
	fmt.Println()
	fmt.Printf("  %s: %d/%d solved\n\n", t.name, t.passed(), len(t.tests))
	if t.buildFailed {
		fmt.Println("  ⚠️  This package does not compile, so no test could run.")
		printOutput(t)
		fmt.Println()
		return
	}
	for i, name := range t.tests {
		icon := "❌"
		switch t.status[name] {
		case "pass":
			icon = "✅"
		case "skip":
			icon = "⏭️ "
		case "":
			icon = "💥" // never ran: an earlier test probably panicked
		}
		fmt.Printf("  %s %2d. %s\n", icon, i+1, exerciseName(name))
	}
	if len(t.output) > 0 {
		fmt.Println()
		fmt.Println("  A test panicked, so the tests after it did not run (💥):")
		printOutput(t)
	}
	fmt.Println()
	fmt.Printf("  See failure details with: go test ./%s/%s\n\n", exercisesDir, t.name)
}

func printOutput(t *topic) {
	fmt.Println()
	for _, line := range t.output {
		if line != "" {
			fmt.Println("    " + line)
		}
	}
}
