package iterators

import (
	"iter"
	"slices"
	"testing"
)

// first calls seq directly (instead of with range) and collects at most max
// values. It fails the test if the iterator keeps calling yield after being
// told to stop, which would otherwise hang or crash the whole test run.
func first[T any](t *testing.T, seq iter.Seq[T], max int) []T {
	t.Helper()
	var out []T
	stopped := false
	seq(func(v T) bool {
		if stopped {
			t.Fatalf("yield was called again after it returned false (collected so far: %v)", out)
		}
		if len(out) == max {
			stopped = true
			return false
		}
		out = append(out, v)
		return true
	})
	return out
}

func TestCountTo(t *testing.T) {
	if got := first(t, CountTo(5), 100); !slices.Equal(got, []int{1, 2, 3, 4, 5}) {
		t.Errorf("CountTo(5) yielded %v, want [1 2 3 4 5]", got)
	}
	if got := first(t, CountTo(5), 2); !slices.Equal(got, []int{1, 2}) {
		t.Errorf("CountTo(5) stopped after 2 yielded %v, want [1 2]", got)
	}
	if got := first(t, CountTo(0), 100); len(got) != 0 {
		t.Errorf("CountTo(0) yielded %v, want nothing", got)
	}
}

func TestEvens(t *testing.T) {
	if got := first(t, Evens(7), 100); !slices.Equal(got, []int{0, 2, 4, 6}) {
		t.Errorf("Evens(7) yielded %v, want [0 2 4 6]", got)
	}
	if got := first(t, Evens(8), 100); !slices.Equal(got, []int{0, 2, 4, 6, 8}) {
		t.Errorf("Evens(8) yielded %v, want [0 2 4 6 8]", got)
	}
}

func TestEnumerate(t *testing.T) {
	var idx []int
	var vals []string
	for i, v := range Enumerate([]string{"a", "b", "c"}) {
		idx = append(idx, i)
		vals = append(vals, v)
	}
	if !slices.Equal(idx, []int{0, 1, 2}) || !slices.Equal(vals, []string{"a", "b", "c"}) {
		t.Errorf("Enumerate([a b c]) yielded indexes %v and values %q", idx, vals)
	}
	calls := 0
	Enumerate([]int{1, 2, 3})(func(int, int) bool {
		calls++
		if calls > 1 {
			t.Fatal("yield was called again after it returned false")
		}
		return false
	})
}

func TestFilter(t *testing.T) {
	isOdd := func(n int) bool { return n%2 == 1 }
	src := slices.Values([]int{1, 2, 3, 4, 5, 6, 7})
	if got := first(t, Filter(src, isOdd), 100); !slices.Equal(got, []int{1, 3, 5, 7}) {
		t.Errorf("Filter(1..7, isOdd) yielded %v, want [1 3 5 7]", got)
	}
	if got := first(t, Filter(src, isOdd), 2); !slices.Equal(got, []int{1, 3}) {
		t.Errorf("Filter(1..7, isOdd) stopped after 2 yielded %v, want [1 3]", got)
	}
}

func TestMap(t *testing.T) {
	src := slices.Values([]string{"go", "is", "fun"})
	if got := first(t, Map(src, func(s string) int { return len(s) }), 100); !slices.Equal(got, []int{2, 2, 3}) {
		t.Errorf("Map(len) yielded %v, want [2 2 3]", got)
	}
	if got := first(t, Map(src, func(s string) string { return s + "!" }), 1); !slices.Equal(got, []string{"go!"}) {
		t.Errorf("Map(add !) stopped after 1 yielded %q, want [go!]", got)
	}
}

func TestTake(t *testing.T) {
	src := slices.Values([]int{10, 20, 30, 40})
	if got := first(t, Take(src, 2), 100); !slices.Equal(got, []int{10, 20}) {
		t.Errorf("Take(src, 2) yielded %v, want [10 20]", got)
	}
	if got := first(t, Take(src, 10), 100); !slices.Equal(got, []int{10, 20, 30, 40}) {
		t.Errorf("Take(src, 10) yielded %v, want [10 20 30 40]", got)
	}
	if got := first(t, Take(src, 0), 100); len(got) != 0 {
		t.Errorf("Take(src, 0) yielded %v, want nothing", got)
	}
}

func TestFibonacci(t *testing.T) {
	want := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	if got := first(t, Fibonacci(), 10); !slices.Equal(got, want) {
		t.Errorf("first 10 of Fibonacci() = %v, want %v", got, want)
	}
	var withBreak []int
	for v := range Fibonacci() {
		if v > 50 {
			break
		}
		withBreak = append(withBreak, v)
	}
	if !slices.Equal(withBreak, want) {
		t.Errorf("Fibonacci() values up to 50 = %v, want %v", withBreak, want)
	}
}

func TestSortedByKey(t *testing.T) {
	m := map[string]int{"pear": 3, "apple": 1, "fig": 2}
	var keys []string
	var vals []int
	for k, v := range SortedByKey(m) {
		keys = append(keys, k)
		vals = append(vals, v)
	}
	if !slices.Equal(keys, []string{"apple", "fig", "pear"}) || !slices.Equal(vals, []int{1, 2, 3}) {
		t.Errorf("SortedByKey yielded keys %q and values %v, want [apple fig pear] and [1 2 3]", keys, vals)
	}
}
