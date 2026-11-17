//go:build solutions

// SPOILER ALERT: reference solutions. Try the exercises first!

package iterators

import (
	"cmp"
	"iter"
	"maps"
	"slices"
)

func CountTo(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 1; i <= n; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

func Evens(limit int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 0; i <= limit; i += 2 {
			if !yield(i) {
				return
			}
		}
	}
}

func Enumerate[T any](s []T) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, v := range s {
			if !yield(i, v) {
				return
			}
		}
	}
}

func Filter[T any](seq iter.Seq[T], keep func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range seq {
			if keep(v) && !yield(v) {
				return
			}
		}
	}
}

func Map[T, U any](seq iter.Seq[T], f func(T) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		for v := range seq {
			if !yield(f(v)) {
				return
			}
		}
	}
}

func Take[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		if n <= 0 {
			return
		}
		taken := 0
		for v := range seq {
			if !yield(v) {
				return
			}
			taken++
			if taken == n {
				return
			}
		}
	}
}

func Fibonacci() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 0, 1
		for yield(a) {
			a, b = b, a+b
		}
	}
}

func SortedByKey[K cmp.Ordered, V any](m map[K]V) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, k := range slices.Sorted(maps.Keys(m)) {
			if !yield(k, m[k]) {
				return
			}
		}
	}
}
