//go:build !solutions

// Package functions covers multiple/named return values, variadic functions,
// functions as values, closures, recursion and defer.
//
// Quick syntax reminder:
//
//	func add(a, b int) int { return a + b }
//	func divmod(a, b int) (q, r int) { q = a / b; r = a % b; return }  // named returns
//	func sum(nums ...int) int { }          // variadic: nums is a []int
//	sum(1, 2, 3); sum(slice...)            // calling a variadic function
//	f := func(x int) int { return x * 2 }  // function literal (anonymous function)
//	type Op func(int, int) int             // a function type
//	defer fmt.Println("runs when the surrounding function returns")
package functions

// Exercise 1: MinMax
//
// Return the smallest and largest numbers in a non-empty slice.
//
//	MinMax([]int{3, 1, 4, 1, 5}) → 1, 5
//
// Concepts: named return values (lo, hi).
func MinMax(nums []int) (lo, hi int) {
	// TODO: implement
	return 0, 0
}

// Exercise 2: Sum
//
// Return the sum of any number of ints.
//
//	Sum()        → 0
//	Sum(1, 2, 3) → 6
//
// Concepts: variadic parameters.
func Sum(nums ...int) int {
	// TODO: implement
	return 0
}

// Exercise 3: Apply
//
// Return a new slice with f applied to every element.
//
//	Apply([]int{1, 2, 3}, func(n int) int { return n * 10 }) → [10 20 30]
//
// Concepts: functions as parameters.
func Apply(nums []int, f func(int) int) []int {
	// TODO: implement
	return nil
}

// Exercise 4: MakeMultiplier
//
// Return a function that multiplies its argument by factor.
//
//	triple := MakeMultiplier(3)
//	triple(5) → 15
//
// Concepts: returning functions, closures capturing variables.
func MakeMultiplier(factor int) func(int) int {
	// TODO: implement
	return func(n int) int { return 0 }
}

// Exercise 5: NewCounter
//
// Return a function that returns 1 the first time it's called, then 2, 3, ...
// Each counter returned by NewCounter must be independent.
//
//	next := NewCounter()
//	next() → 1
//	next() → 2
//
// Concepts: closures keep their captured variables alive between calls.
func NewCounter() func() int {
	// TODO: implement
	return func() int { return 0 }
}

// Exercise 6: Compose
//
// Return a function that computes f(g(x)).
//
//	addOne := func(x int) int { return x + 1 }
//	double := func(x int) int { return x * 2 }
//	Compose(addOne, double)(5) → 11   (double first, then addOne)
func Compose(f, g func(int) int) func(int) int {
	// TODO: implement
	return func(x int) int { return 0 }
}

// Exercise 7: Power
//
// Return base raised to exp (exp >= 0). Solve it recursively:
// base^0 = 1, base^exp = base × base^(exp-1).
//
//	Power(2, 10) → 1024
//
// Concepts: recursion.
func Power(base, exp int) int {
	// TODO: implement
	return 0
}

// Exercise 8: Divide
//
// Return a / b and true, or 0 and false when b is 0.
//
//	Divide(10, 4) → 2.5, true
//	Divide(1, 0)  → 0, false
//
// Concepts: the "value, ok" idiom.
func Divide(a, b float64) (float64, bool) {
	// TODO: implement
	return 0, false
}

// Exercise 9: JoinWith
//
// Join the parts using sep, without using strings.Join.
//
//	JoinWith("-", "a", "b", "c") → "a-b-c"
//	JoinWith(", ")               → ""
//
// Concepts: a regular parameter followed by a variadic one.
func JoinWith(sep string, parts ...string) string {
	// TODO: implement
	return ""
}

// Exercise 10: Memoize
//
// Return a function that behaves like f, but remembers results: calling it
// again with an argument it has already seen must NOT call f again.
//
// Concepts: closures capturing a map.
func Memoize(f func(int) int) func(int) int {
	// TODO: implement
	return func(n int) int { return 0 }
}

// Exercise 11: DeferredSteps
//
// Using defer, make this function return exactly:
//
//	["start", "end", "deferred 2", "deferred 1"]
//
// Steps: append "start"; defer a func that appends "deferred 1"; defer a func
// that appends "deferred 2"; append "end"; return.
//
// Concepts: deferred calls run in LIFO order *after* the return statement,
// and they can modify named return values.
func DeferredSteps() (steps []string) {
	// TODO: implement
	return nil
}

// Exercise 12: ApplyN
//
// Apply f to x, n times.
//
//	ApplyN(func(x int) int { return x * 2 }, 3, 1) → 8
//	ApplyN(f, 0, 7) → 7
func ApplyN(f func(int) int, n int, x int) int {
	// TODO: implement
	return 0
}

// Predicate is a named function type.
type Predicate func(int) bool

// Exercise 13: CountIf
//
// Return how many numbers satisfy p.
//
//	CountIf([]int{1, 2, 3, 4}, func(n int) bool { return n > 2 }) → 2
//
// Concepts: named function types.
func CountIf(nums []int, p Predicate) int {
	// TODO: implement
	return 0
}
