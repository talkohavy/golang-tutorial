//go:build !solutions

package slicesx

// Exercise 1: MakeRange
//
// Return the integers from start (inclusive) to end (exclusive).
// Return an empty slice when start >= end.
//
//	MakeRange(2, 6) → [2 3 4 5]
func MakeRange(start, end int) []int {
	// TODO: implement
	return nil
}

// Exercise 2: Last
//
// Return the last element and true, or 0 and false for an empty slice.
//
//	Last([]int{1, 2, 3}) → 3, true
func Last(nums []int) (int, bool) {
	// TODO: implement
	return 0, false
}

// Exercise 3: RemoveAt
//
// Return a NEW slice without the element at index i. The input slice must
// not be modified.
//
//	RemoveAt([]int{10, 20, 30, 40}, 1) → [10 30 40]
//
// Gotcha: append(nums[:i], nums[i+1:]...) overwrites the input!
func RemoveAt(nums []int, i int) []int {
	// TODO: implement
	return nil
}

// Exercise 4: InsertAt
//
// Return a NEW slice with v inserted at index i (0 <= i <= len(nums)).
// The input slice must not be modified.
//
//	InsertAt([]int{1, 2, 4}, 2, 3) → [1 2 3 4]
func InsertAt(nums []int, i, v int) []int {
	// TODO: implement
	return nil
}

// Exercise 5: Evens
//
// Return only the even numbers, in their original order.
//
//	Evens([]int{1, 2, 3, 4, 6}) → [2 4 6]
func Evens(nums []int) []int {
	// TODO: implement
	return nil
}

// Exercise 6: ReverseInPlace
//
// Reverse the slice in place (modify it, return nothing). Don't use
// slices.Reverse; swap elements yourself.
//
//	s := []int{1, 2, 3}; ReverseInPlace(s) → s is now [3 2 1]
//
// Concepts: a slice passed to a function still points to the same array.
func ReverseInPlace(nums []int) {
	// TODO: implement
}

// Exercise 7: Chunk
//
// Split nums into chunks of the given size (size >= 1). The last chunk may be
// shorter.
//
//	Chunk([]int{1, 2, 3, 4, 5}, 2) → [[1 2] [3 4] [5]]
func Chunk(nums []int, size int) [][]int {
	// TODO: implement
	return nil
}

// Exercise 8: Unique
//
// Return the numbers without duplicates, keeping the first occurrence order.
//
//	Unique([]int{3, 1, 3, 2, 1}) → [3 1 2]
func Unique(nums []int) []int {
	// TODO: implement
	return nil
}

// Exercise 9: Flatten
//
// Concatenate all rows into a single slice.
//
//	Flatten([][]int{{1, 2}, {3}, {}, {4, 5}}) → [1 2 3 4 5]
func Flatten(matrix [][]int) []int {
	// TODO: implement
	return nil
}

// Exercise 10: Transpose
//
// Transpose a rectangular matrix: rows become columns.
//
//	Transpose([][]int{{1, 2, 3}, {4, 5, 6}}) → [[1 4] [2 5] [3 6]]
func Transpose(matrix [][]int) [][]int {
	// TODO: implement
	return nil
}

// Exercise 11: RotateLeft
//
// Return a NEW slice rotated left by k positions. k may be larger than the
// length.
//
//	RotateLeft([]int{1, 2, 3, 4, 5}, 2) → [3 4 5 1 2]
func RotateLeft(nums []int, k int) []int {
	// TODO: implement
	return nil
}

// Exercise 12: DoubledCopy
//
// Return a copy of nums with every element doubled. The input must be left
// untouched.
//
//	DoubledCopy([]int{1, 2, 3}) → [2 4 6]
//
// Concepts: make + copy, or slices.Clone.
func DoubledCopy(nums []int) []int {
	// TODO: implement
	return nil
}

// Exercise 13: SortedDesc
//
// Return a NEW slice with the numbers sorted from largest to smallest. The
// input must be left untouched.
//
//	SortedDesc([]int{3, 1, 2}) → [3 2 1]
//
// Concepts: slices.Clone, slices.Sort, slices.Reverse (or sort.Slice).
func SortedDesc(nums []int) []int {
	// TODO: implement
	return nil
}

// Exercise 14: ContainsAll
//
// Report whether every element of needles appears in haystack.
// An empty needles slice is always contained.
//
//	ContainsAll([]int{1, 2, 3}, []int{3, 1}) → true
//
// Concepts: slices.Contains.
func ContainsAll(haystack, needles []int) bool {
	// TODO: implement
	return false
}

// Exercise 15: DoubleArray
//
// Return a new array with every element doubled.
//
//	DoubleArray([3]int{1, 2, 3}) → [2 4 6]
//
// Concepts: arrays are VALUES: the function receives a copy, so modifying
// arr inside the function never affects the caller's array.
func DoubleArray(arr [3]int) [3]int {
	// TODO: implement
	return [3]int{}
}

// Exercise 16: WithCapacity
//
// Return an empty slice (length 0) whose capacity is n, then append the
// numbers 1..n to it and return it. Appending within capacity never
// reallocates.
//
//	s := WithCapacity(3) → s == [1 2 3], cap(s) == 3
//
// Concepts: make([]int, 0, n).
func WithCapacity(n int) []int {
	// TODO: implement
	return nil
}
