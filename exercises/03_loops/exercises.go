//go:build !solutions

package loops

// Exercise 1: SumTo
//
// Return 1 + 2 + ... + n. Return 0 if n < 1.
//
//	SumTo(5) → 15
func SumTo(n int) int {
	// TODO: implement
	return 0
}

// Exercise 2: Factorial
//
// Return n! (1 × 2 × ... × n). By definition 0! = 1.
//
//	Factorial(5) → 120
func Factorial(n int) int {
	// TODO: implement
	return 0
}

// Exercise 3: CountDown
//
// Return a slice counting down from n to 1. Return an empty slice if n < 1.
//
//	CountDown(3) → [3 2 1]
//
// Concepts: append.
func CountDown(n int) []int {
	// TODO: implement
	return nil
}

// Exercise 4: SumSlice
//
// Return the sum of all the numbers in the slice.
//
//	SumSlice([]int{1, 2, 3}) → 6
//
// Concepts: for ... range over a slice.
func SumSlice(nums []int) int {
	// TODO: implement
	return 0
}

// Exercise 5: CountVowels
//
// Count the vowels (a, e, i, o, u, case-insensitive) in s.
//
//	CountVowels("Hello World") → 3
//
// Concepts: ranging over a string yields runes.
func CountVowels(s string) int {
	// TODO: implement
	return 0
}

// Exercise 6: IsPrime
//
// Report whether n is a prime number. Numbers below 2 are not prime.
//
//	IsPrime(7) → true
//	IsPrime(9) → false
//
// Tip: you only need to check divisors up to √n (i*i <= n).
func IsPrime(n int) bool {
	// TODO: implement
	return false
}

// Exercise 7: Fibonacci
//
// Return the first n Fibonacci numbers, starting from 0.
//
//	Fibonacci(7) → [0 1 1 2 3 5 8]
//	Fibonacci(0) → []
func Fibonacci(n int) []int {
	// TODO: implement
	return nil
}

// Exercise 8: MultiplicationTable
//
// Return an n×n table where table[i][j] = (i+1) × (j+1).
//
//	MultiplicationTable(3) → [[1 2 3] [2 4 6] [3 6 9]]
//
// Concepts: nested loops, slices of slices, make([]int, n).
func MultiplicationTable(n int) [][]int {
	// TODO: implement
	return nil
}

// Exercise 9: IndexOfFirstNegative
//
// Return the index of the first negative number, or -1 if there is none.
//
//	IndexOfFirstNegative([]int{3, 0, -2, -5}) → 2
//
// Concepts: returning (or breaking) early from a loop.
func IndexOfFirstNegative(nums []int) int {
	// TODO: implement
	return 0
}

// Exercise 10: SumOfEvens
//
// Return the sum of the even numbers in nums. Practice using `continue` to
// skip odd numbers.
//
//	SumOfEvens([]int{1, 2, 3, 4}) → 6
func SumOfEvens(nums []int) int {
	// TODO: implement
	return 0
}

// Exercise 11: CollatzSteps
//
// Starting from n (n >= 1), repeatedly apply: if n is even, n = n / 2,
// otherwise n = 3n + 1. Return how many steps it takes to reach 1.
//
//	CollatzSteps(1) → 0
//	CollatzSteps(6) → 8   (6 3 10 5 16 8 4 2 1)
//
// Concepts: "while" style loop: for n != 1 { ... }
func CollatzSteps(n int) int {
	// TODO: implement
	return 0
}

// Exercise 12: Staircase
//
// Return n lines, where line i (starting at 1) contains i asterisks.
// Build each line with an inner loop.
//
//	Staircase(3) → ["*", "**", "***"]
func Staircase(n int) []string {
	// TODO: implement
	return nil
}

// Exercise 13: FindPair
//
// Return the indexes i < j of the first pair whose values add up to target
// (scan i from left to right, then j from i+1). Return -1, -1 if no pair
// exists.
//
//	FindPair([]int{2, 7, 11, 15}, 9) → 0, 1
//
// Concepts: nested loops; try a labeled break instead of returning directly.
func FindPair(nums []int, target int) (int, int) {
	// TODO: implement
	return 0, 0
}

// Exercise 14: DigitSum
//
// Return the sum of the decimal digits of n. Negative numbers are treated as
// positive.
//
//	DigitSum(1234) → 10
//	DigitSum(-56)  → 11
//
// Concepts: n % 10 gives the last digit, n / 10 removes it.
func DigitSum(n int) int {
	// TODO: implement
	return 0
}

// Exercise 15: Squares
//
// Return the squares of 0..n-1.
//
//	Squares(4) → [0 1 4 9]
//
// Concepts: `for i := range n` (range over an integer, Go 1.22+).
func Squares(n int) []int {
	// TODO: implement
	return nil
}
