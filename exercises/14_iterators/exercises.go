//go:build !solutions

// Package iterators covers range-over-func iterators (Go 1.23+).
//
// An iterator is just a function that calls yield for every value. If yield
// returns false (because the caller used `break`), the iterator must stop.
//
//	type Seq[V any]     func(yield func(V) bool)       // from package iter
//	type Seq2[K, V any] func(yield func(K, V) bool)
//
//	func CountTo(n int) iter.Seq[int] {
//		return func(yield func(int) bool) {
//			for i := 1; i <= n; i++ {
//				if !yield(i) {
//					return              // the consumer said stop
//				}
//			}
//		}
//	}
//
//	for v := range CountTo(3) { }       // consuming an iterator
//	slices.Collect(seq)                 // iterator → slice
//	maps.Keys(m), slices.Values(s)      // standard library iterators
package iterators

import (
	"cmp"
	"iter"
)

// Exercise 1: CountTo
//
// Yield 1, 2, ..., n.
func CountTo(n int) iter.Seq[int] {
	// TODO: implement
	return func(yield func(int) bool) {}
}

// Exercise 2: Evens
//
// Yield the even numbers from 0 up to and including limit.
//
//	Evens(7) → 0, 2, 4, 6
func Evens(limit int) iter.Seq[int] {
	// TODO: implement
	return func(yield func(int) bool) {}
}

// Exercise 3: Enumerate
//
// Yield (index, value) pairs for the slice.
//
//	for i, v := range Enumerate([]string{"a", "b"}) { }   // (0, "a"), (1, "b")
//
// Concepts: iter.Seq2 for two values per iteration.
func Enumerate[T any](s []T) iter.Seq2[int, T] {
	// TODO: implement
	return func(yield func(int, T) bool) {}
}

// Exercise 4: Filter
//
// Yield only the values of seq for which keep returns true.
//
// Concepts: iterators can wrap other iterators (`for v := range seq`
// works inside an iterator too).
func Filter[T any](seq iter.Seq[T], keep func(T) bool) iter.Seq[T] {
	// TODO: implement
	return func(yield func(T) bool) {}
}

// Exercise 5: Map
//
// Yield f(v) for every value of seq.
func Map[T, U any](seq iter.Seq[T], f func(T) U) iter.Seq[U] {
	// TODO: implement
	return func(yield func(U) bool) {}
}

// Exercise 6: Take
//
// Yield at most the first n values of seq.
func Take[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	// TODO: implement
	return func(yield func(T) bool) {}
}

// Exercise 7: Fibonacci
//
// Yield the Fibonacci numbers 0, 1, 1, 2, 3, 5, ... FOREVER. The consumer
// decides when to stop (with break, or by using Take), so you must respect
// yield's return value.
func Fibonacci() iter.Seq[int] {
	// TODO: implement
	return func(yield func(int) bool) {}
}

// Exercise 8: SortedByKey
//
// Yield the map's (key, value) pairs ordered by key.
//
//	for k, v := range SortedByKey(map[string]int{"b": 2, "a": 1}) { }  // ("a", 1), ("b", 2)
func SortedByKey[K cmp.Ordered, V any](m map[K]V) iter.Seq2[K, V] {
	// TODO: implement
	return func(yield func(K, V) bool) {}
}
