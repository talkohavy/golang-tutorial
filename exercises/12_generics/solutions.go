//go:build solutions

// SPOILER ALERT: reference solutions. Try the exercises first!

package generics

import "cmp"

type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

type Pair[A, B any] struct {
	First  A
	Second B
}

func Map[T, U any](s []T, f func(T) U) []U {
	result := make([]U, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

func Filter[T any](s []T, keep func(T) bool) []T {
	var result []T
	for _, v := range s {
		if keep(v) {
			result = append(result, v)
		}
	}
	return result
}

func Reduce[T, A any](s []T, initial A, f func(A, T) A) A {
	acc := initial
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

func Contains[T comparable](s []T, v T) bool {
	for _, x := range s {
		if x == v {
			return true
		}
	}
	return false
}

func Sum[T Number](nums []T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}

func Max[T cmp.Ordered](s []T) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	best := s[0]
	for _, v := range s[1:] {
		if v > best {
			best = v
		}
	}
	return best, true
}

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func GroupBy[T any, K comparable](s []T, keyFn func(T) K) map[K][]T {
	groups := map[K][]T{}
	for _, v := range s {
		k := keyFn(v)
		groups[k] = append(groups[k], v)
	}
	return groups
}

func Zip[A, B any](a []A, b []B) []Pair[A, B] {
	n := min(len(a), len(b))
	result := make([]Pair[A, B], n)
	for i := range n {
		result[i] = Pair[A, B]{First: a[i], Second: b[i]}
	}
	return result
}

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}

type Set[T comparable] struct {
	items map[T]struct{}
}

func NewSet[T comparable](items ...T) *Set[T] {
	s := &Set[T]{items: map[T]struct{}{}}
	for _, it := range items {
		s.Add(it)
	}
	return s
}

func (s *Set[T]) Add(v T) {
	s.items[v] = struct{}{}
}

func (s *Set[T]) Has(v T) bool {
	_, ok := s.items[v]
	return ok
}

func (s *Set[T]) Len() int {
	return len(s.items)
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for v := range s.items {
		result.Add(v)
	}
	for v := range other.items {
		result.Add(v)
	}
	return result
}

func Reverse[T any](s []T) []T {
	result := make([]T, len(s))
	for i, v := range s {
		result[len(s)-1-i] = v
	}
	return result
}

func Unique[T comparable](s []T) []T {
	seen := map[T]bool{}
	var result []T
	for _, v := range s {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

func OrDefault[T comparable](v, def T) T {
	var zero T
	if v == zero {
		return def
	}
	return v
}
