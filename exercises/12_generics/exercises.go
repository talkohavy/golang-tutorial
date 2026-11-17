//go:build !solutions

// Package generics covers type parameters (Go 1.18+).
//
// Quick syntax reminder:
//
//	func Map[T, U any](s []T, f func(T) U) []U { ... }   // generic function
//	Map([]int{1, 2}, strconv.Itoa)                       // type arguments are inferred
//
//	type Stack[T any] struct { items []T }               // generic type
//	func (s *Stack[T]) Push(v T) { ... }                 // its methods
//	s := &Stack[int]{}
//
// Constraints say what a type parameter must support:
//
//	any                      // anything
//	comparable               // supports == and != (needed for map keys)
//	cmp.Ordered              // supports < <= > >= (numbers and strings)
//	interface{ ~int | ~float64 }  // a union; ~int also allows `type MyInt int`
//
//	var zero T               // the zero value of any type parameter
package generics

import "cmp"

// Number is a constraint that allows any integer or float type, including
// custom types built on them (that's what the ~ means).
type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

// Pair holds two values of possibly different types.
type Pair[A, B any] struct {
	First  A
	Second B
}

// Exercise 1: Map
//
// Return a new slice with f applied to every element.
//
//	Map([]int{1, 2, 3}, func(n int) string { return strconv.Itoa(n * 2) }) → ["2" "4" "6"]
func Map[T, U any](s []T, f func(T) U) []U {
	// TODO: implement
	return nil
}

// Exercise 2: Filter
//
// Return the elements for which keep returns true.
//
//	Filter([]string{"go", "rust", "c"}, func(s string) bool { return len(s) > 1 }) → ["go" "rust"]
func Filter[T any](s []T, keep func(T) bool) []T {
	// TODO: implement
	return nil
}

// Exercise 3: Reduce
//
// Fold the slice into a single value, starting from initial.
//
//	Reduce([]int{1, 2, 3}, 0, func(acc, n int) int { return acc + n })  → 6
//	Reduce([]string{"a", "b"}, "", func(acc, s string) string { return acc + s }) → "ab"
func Reduce[T, A any](s []T, initial A, f func(A, T) A) A {
	// TODO: implement
	return initial
}

// Exercise 4: Contains
//
// Report whether v is in s.
//
// Concepts: the comparable constraint allows ==.
func Contains[T comparable](s []T, v T) bool {
	// TODO: implement
	return false
}

// Exercise 5: Sum
//
// Return the sum of the numbers.
//
//	Sum([]float64{1.5, 2.5}) → 4.0
//
// Concepts: custom constraints (see Number above).
func Sum[T Number](nums []T) T {
	// TODO: implement
	var zero T
	return zero
}

// Exercise 6: Max
//
// Return the largest element and true, or the zero value and false if s is
// empty. Don't use slices.Max.
//
// Concepts: cmp.Ordered allows the < and > operators.
func Max[T cmp.Ordered](s []T) (T, bool) {
	// TODO: implement
	var zero T
	return zero, false
}

// Exercise 7: Keys
//
// Return the keys of the map, in any order.
func Keys[K comparable, V any](m map[K]V) []K {
	// TODO: implement
	return nil
}

// Exercise 8: GroupBy
//
// Group elements by the key returned by keyFn, preserving order in groups.
//
//	GroupBy([]string{"apple", "avocado", "banana"}, func(s string) byte { return s[0] })
//	→ map['a':["apple" "avocado"] 'b':["banana"]]
func GroupBy[T any, K comparable](s []T, keyFn func(T) K) map[K][]T {
	// TODO: implement
	return nil
}

// Exercise 9: Zip
//
// Pair up elements at the same index. The result is as long as the shorter
// slice.
//
//	Zip([]string{"a", "b", "c"}, []int{1, 2}) → [{a 1} {b 2}]
func Zip[A, B any](a []A, b []B) []Pair[A, B] {
	// TODO: implement
	return nil
}

// Stack is a generic last-in-first-out stack. Its zero value is ready to use.
type Stack[T any] struct {
	items []T
}

// Exercise 10: Stack[T]
//
// Push adds to the top. Pop removes and returns the top (zero value and
// false if empty). Len returns the number of items.
func (s *Stack[T]) Push(v T) {
	// TODO: implement
}

func (s *Stack[T]) Pop() (T, bool) {
	// TODO: implement
	var zero T
	return zero, false
}

func (s *Stack[T]) Len() int {
	// TODO: implement
	return 0
}

// Set is a generic set.
type Set[T comparable] struct {
	items map[T]struct{}
}

// Exercise 11: Set[T]
//
//   - NewSet creates a set containing the given items.
//   - Add inserts an item. Has reports membership. Len returns the size.
//   - Union returns a NEW set with the items of both sets.
func NewSet[T comparable](items ...T) *Set[T] {
	// TODO: implement
	return &Set[T]{}
}

func (s *Set[T]) Add(v T) {
	// TODO: implement
}

func (s *Set[T]) Has(v T) bool {
	// TODO: implement
	return false
}

func (s *Set[T]) Len() int {
	// TODO: implement
	return 0
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	// TODO: implement
	return &Set[T]{}
}

// Exercise 12: Reverse
//
// Return a NEW slice with the elements in reverse order.
func Reverse[T any](s []T) []T {
	// TODO: implement
	return nil
}

// Exercise 13: Unique
//
// Return the elements without duplicates, keeping first-occurrence order.
func Unique[T comparable](s []T) []T {
	// TODO: implement
	return nil
}

// Exercise 14: OrDefault
//
// Return v, unless v is the zero value of its type, in which case return def.
//
//	OrDefault("", "anonymous") → "anonymous"
//	OrDefault(0, 8080)         → 8080
//	OrDefault(3000, 8080)      → 3000
//
// Concepts: `var zero T` gives you the zero value to compare against.
func OrDefault[T comparable](v, def T) T {
	// TODO: implement
	var zero T
	return zero
}
