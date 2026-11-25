# Go Playground

200+ small, test-driven exercises for getting fluent in Go syntax, from `for` loops to goroutines.

Each exercise is a function with a `// TODO` body and a test that checks it. Your job: make the tests pass.

## Quick start

```bash
go run ./cmd/progress            # see where you are and what's next
```

Then open the suggested file (e.g. `exercises/01_basics/exercises.go`), read the exercise's comment, write the code, and run its test:

```bash
go test ./exercises/01_basics -run '^TestGreeting$'   # one exercise
go test ./exercises/01_basics                         # a whole topic, with failure details
go run ./cmd/progress 01_basics                       # ✅ / ❌ per exercise in a topic
```

Tip: in Cursor / VS Code (with the Go extension), click the "run test" link that appears above each `func TestXxx` in `exercises_test.go`.

## Topics

Work through them in order; each topic builds on the previous ones.

| Topic | Exercises | What you practice |
| --- | --- | --- |
| `01_basics` | 13 | variables, constants, `iota`, types, conversions, `fmt.Sprintf` |
| `02_conditionals` | 12 | `if`/`else`, if-with-init, `switch`, tagless `switch` |
| `03_loops` | 15 | `for` in all its forms, `range`, `break`/`continue`, labels |
| `04_functions` | 13 | multiple/named returns, variadics, closures, recursion, `defer` |
| `05_strings` | 14 | bytes vs runes, `strings`, `strconv`, `unicode`, `strings.Builder` |
| `06_slices` | 16 | arrays, slices, `append`, `copy`, aliasing gotchas, `slices` package |
| `07_maps` | 15 | maps, comma-ok, sets, nested maps, sorting keys |
| `08_structs` | 45 | structs, methods, constructors, embedding, encapsulation, JSON tags (blank-canvas format, see below) |
| `09_pointers` | 12 | pointers, value vs reference semantics, linked lists |
| `10_interfaces` | 14 | interfaces, `Stringer`, `io.Writer`, `sort.Interface`, type switches |
| `11_errors` | 11 | errors, sentinel/custom errors, wrapping, `errors.Is/As`, panic/recover |
| `12_generics` | 14 | type parameters, constraints, generic types (`Stack[T]`, `Set[T]`) |
| `13_concurrency` | 11 | goroutines, channels, `select`, `WaitGroup`, `Mutex`, `context`, atomics |
| `14_iterators` | 8 | range-over-func iterators, `iter.Seq`, `yield` |
| `15_challenges` | 18 | classic puzzles: FizzBuzz, Roman numerals, two-sum, spiral matrix... |

Every topic's `exercises.go` (or `doc.go` for `08_structs`) starts with a short syntax cheat sheet.

Go has no classes: `08_structs` shows the idiomatic replacement (structs + methods + constructor functions + embedding), and `10_interfaces` covers polymorphism.

## Project layout

```text
exercises/
  01_basics/
    exercises.go         ← you edit this
    exercises_test.go    ← the tests (read them when you're stuck!)
    solutions.go         ← reference solutions (spoilers!)
  ...
  08_structs/
    doc.go               ← cheat sheet
    01_person/
      exercise.go        ← a blank canvas: you write every type, function and method
      exercise_test.go   ← one sanity-check test per step
      solution.go        ← reference solution (spoilers!)
    02_rectangle/
    ...
cmd/progress/            ← the progress tracker
internal/testkit/        ← helpers that let blank-canvas tests run one step at a time
scratch/main.go          ← free-form playground: go run ./scratch
notes/main.go            ← my original lesson notes: go run ./notes
```

It's a single Go module where each topic is its own package, which is how real Go projects are organized.

### Blank-canvas topics

In `08_structs`, nothing is pre-written: each exercise folder is its own small package, and its `exercise.go` holds only a description of each step. Do the steps in order and run each step's test as you go:

```bash
go test ./exercises/08_structs/02_rectangle -run '^TestArea$'   # one step
go test ./exercises/08_structs/02_rectangle                     # one exercise
go test ./exercises/08_structs/...                              # the whole topic
```

A compile error like `undefined: Rectangle` means the tests are waiting for something you haven't written yet. Tests look up methods by name, so a missing method only fails its own step ("Rectangle has no Perimeter method yet") instead of breaking the whole exercise.

## Solutions

Each topic has a `solutions.go` (each `08_structs` exercise folder has a `solution.go`). It's excluded from normal builds by a build tag (`//go:build solutions`), so it never interferes with your work. Peek only after you've solved an exercise (or are really stuck), to compare approaches.

```bash
go run ./cmd/progress -solutions             # proves every test passes with the reference solutions
go test -tags solutions ./exercises/03_loops
```

## Useful commands

```bash
go test -race ./exercises/13_concurrency     # detect data races (great for the concurrency topic)
go test -v ./exercises/03_loops              # verbose: list every test
go vet ./exercises/...                       # catch common mistakes
gofmt -w exercises                           # format your code
go doc strings.Fields                        # read docs from the terminal
```

`make progress`, `make progress topic=03_loops`, `make test topic=03_loops`, `make solutions`, and `make scratch` are shortcuts for the above.

## Building & running programs (refresher)

```bash
go run ./scratch                      # compile & run in one step
go build -o dist/scratch ./scratch    # compile to a binary (dist/ is git-ignored)
./dist/scratch                        # run it
```
