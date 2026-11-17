package slicesx

import (
	"slices"
	"testing"
)

func equal2D(a, b [][]int) bool {
	return slices.EqualFunc(a, b, slices.Equal[[]int])
}

func TestMakeRange(t *testing.T) {
	tests := []struct {
		start, end int
		want       []int
	}{
		{2, 6, []int{2, 3, 4, 5}}, {0, 1, []int{0}}, {-2, 1, []int{-2, -1, 0}}, {5, 5, []int{}}, {5, 2, []int{}},
	}
	for _, tt := range tests {
		if got := MakeRange(tt.start, tt.end); !slices.Equal(got, tt.want) {
			t.Errorf("MakeRange(%d, %d) = %v, want %v", tt.start, tt.end, got, tt.want)
		}
	}
}

func TestLast(t *testing.T) {
	if v, ok := Last([]int{1, 2, 3}); v != 3 || !ok {
		t.Errorf("Last([1 2 3]) = %d, %t, want 3, true", v, ok)
	}
	if v, ok := Last([]int{0}); v != 0 || !ok {
		t.Errorf("Last([0]) = %d, %t, want 0, true", v, ok)
	}
	if v, ok := Last(nil); v != 0 || ok {
		t.Errorf("Last([]) = %d, %t, want 0, false", v, ok)
	}
}

func TestRemoveAt(t *testing.T) {
	input := []int{10, 20, 30, 40}
	if got, want := RemoveAt(input, 1), []int{10, 30, 40}; !slices.Equal(got, want) {
		t.Errorf("RemoveAt([10 20 30 40], 1) = %v, want %v", got, want)
	}
	if !slices.Equal(input, []int{10, 20, 30, 40}) {
		t.Errorf("RemoveAt modified its input: it is now %v", input)
	}
	if got, want := RemoveAt([]int{1, 2, 3}, 0), []int{2, 3}; !slices.Equal(got, want) {
		t.Errorf("RemoveAt([1 2 3], 0) = %v, want %v", got, want)
	}
	if got, want := RemoveAt([]int{1, 2, 3}, 2), []int{1, 2}; !slices.Equal(got, want) {
		t.Errorf("RemoveAt([1 2 3], 2) = %v, want %v", got, want)
	}
}

func TestInsertAt(t *testing.T) {
	input := make([]int, 3, 10)
	copy(input, []int{1, 2, 4})
	if got, want := InsertAt(input, 2, 3), []int{1, 2, 3, 4}; !slices.Equal(got, want) {
		t.Errorf("InsertAt([1 2 4], 2, 3) = %v, want %v", got, want)
	}
	if !slices.Equal(input, []int{1, 2, 4}) || !slices.Equal(input[:4], []int{1, 2, 4, 0}) {
		t.Errorf("InsertAt modified its input (or its spare capacity): now %v", input[:4])
	}
	if got, want := InsertAt([]int{2}, 0, 1), []int{1, 2}; !slices.Equal(got, want) {
		t.Errorf("InsertAt([2], 0, 1) = %v, want %v", got, want)
	}
	if got, want := InsertAt([]int{1}, 1, 2), []int{1, 2}; !slices.Equal(got, want) {
		t.Errorf("InsertAt([1], 1, 2) = %v, want %v", got, want)
	}
}

func TestEvens(t *testing.T) {
	tests := []struct{ in, want []int }{
		{[]int{1, 2, 3, 4, 6}, []int{2, 4, 6}}, {[]int{1, 3}, []int{}}, {[]int{-2, 0, 7}, []int{-2, 0}},
	}
	for _, tt := range tests {
		if got := Evens(tt.in); !slices.Equal(got, tt.want) {
			t.Errorf("Evens(%v) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestReverseInPlace(t *testing.T) {
	tests := []struct{ in, want []int }{
		{[]int{1, 2, 3}, []int{3, 2, 1}}, {[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}}, {[]int{9}, []int{9}}, {[]int{}, []int{}},
	}
	for _, tt := range tests {
		s := slices.Clone(tt.in)
		ReverseInPlace(s)
		if !slices.Equal(s, tt.want) {
			t.Errorf("after ReverseInPlace(%v) the slice is %v, want %v", tt.in, s, tt.want)
		}
	}
}

func TestChunk(t *testing.T) {
	tests := []struct {
		in   []int
		size int
		want [][]int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, [][]int{{1, 2}, {3, 4}, {5}}},
		{[]int{1, 2, 3}, 3, [][]int{{1, 2, 3}}},
		{[]int{1, 2}, 5, [][]int{{1, 2}}},
		{[]int{}, 2, [][]int{}},
	}
	for _, tt := range tests {
		if got := Chunk(tt.in, tt.size); !equal2D(got, tt.want) {
			t.Errorf("Chunk(%v, %d) = %v, want %v", tt.in, tt.size, got, tt.want)
		}
	}
}

func TestUnique(t *testing.T) {
	tests := []struct{ in, want []int }{
		{[]int{3, 1, 3, 2, 1}, []int{3, 1, 2}}, {[]int{1, 1, 1}, []int{1}}, {[]int{1, 2}, []int{1, 2}}, {[]int{}, []int{}},
	}
	for _, tt := range tests {
		if got := Unique(tt.in); !slices.Equal(got, tt.want) {
			t.Errorf("Unique(%v) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestFlatten(t *testing.T) {
	in := [][]int{{1, 2}, {3}, {}, {4, 5}}
	if got, want := Flatten(in), []int{1, 2, 3, 4, 5}; !slices.Equal(got, want) {
		t.Errorf("Flatten(%v) = %v, want %v", in, got, want)
	}
	if got := Flatten(nil); len(got) != 0 {
		t.Errorf("Flatten(nil) = %v, want []", got)
	}
}

func TestTranspose(t *testing.T) {
	tests := []struct{ in, want [][]int }{
		{[][]int{{1, 2, 3}, {4, 5, 6}}, [][]int{{1, 4}, {2, 5}, {3, 6}}},
		{[][]int{{1, 2}, {3, 4}}, [][]int{{1, 3}, {2, 4}}},
		{[][]int{{7}}, [][]int{{7}}},
		{[][]int{}, [][]int{}},
	}
	for _, tt := range tests {
		if got := Transpose(tt.in); !equal2D(got, tt.want) {
			t.Errorf("Transpose(%v) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestRotateLeft(t *testing.T) {
	tests := []struct {
		in   []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{3, 4, 5, 1, 2}},
		{[]int{1, 2, 3}, 0, []int{1, 2, 3}},
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}},
		{[]int{1, 2, 3}, 7, []int{2, 3, 1}},
		{[]int{}, 4, []int{}},
	}
	for _, tt := range tests {
		input := slices.Clone(tt.in)
		if got := RotateLeft(input, tt.k); !slices.Equal(got, tt.want) {
			t.Errorf("RotateLeft(%v, %d) = %v, want %v", tt.in, tt.k, got, tt.want)
		}
		if !slices.Equal(input, tt.in) {
			t.Errorf("RotateLeft modified its input: %v became %v", tt.in, input)
		}
	}
}

func TestDoubledCopy(t *testing.T) {
	input := []int{1, 2, 3}
	if got, want := DoubledCopy(input), []int{2, 4, 6}; !slices.Equal(got, want) {
		t.Errorf("DoubledCopy([1 2 3]) = %v, want %v", got, want)
	}
	if !slices.Equal(input, []int{1, 2, 3}) {
		t.Errorf("DoubledCopy modified its input: it is now %v", input)
	}
}

func TestSortedDesc(t *testing.T) {
	input := []int{3, 1, 2, 5, 4}
	if got, want := SortedDesc(input), []int{5, 4, 3, 2, 1}; !slices.Equal(got, want) {
		t.Errorf("SortedDesc([3 1 2 5 4]) = %v, want %v", got, want)
	}
	if !slices.Equal(input, []int{3, 1, 2, 5, 4}) {
		t.Errorf("SortedDesc modified its input: it is now %v", input)
	}
}

func TestContainsAll(t *testing.T) {
	tests := []struct {
		haystack, needles []int
		want              bool
	}{
		{[]int{1, 2, 3}, []int{3, 1}, true}, {[]int{1, 2, 3}, []int{4}, false},
		{[]int{1, 2, 3}, []int{}, true}, {[]int{}, []int{1}, false}, {[]int{5, 5}, []int{5, 5, 5}, true},
	}
	for _, tt := range tests {
		if got := ContainsAll(tt.haystack, tt.needles); got != tt.want {
			t.Errorf("ContainsAll(%v, %v) = %t, want %t", tt.haystack, tt.needles, got, tt.want)
		}
	}
}

func TestDoubleArray(t *testing.T) {
	arr := [3]int{1, 2, 3}
	if got, want := DoubleArray(arr), [3]int{2, 4, 6}; got != want {
		t.Errorf("DoubleArray(%v) = %v, want %v", arr, got, want)
	}
	if arr != [3]int{1, 2, 3} {
		t.Errorf("the caller's array changed to %v (arrays are copied, so this is impossible!)", arr)
	}
}

func TestWithCapacity(t *testing.T) {
	for _, n := range []int{0, 1, 5} {
		got := WithCapacity(n)
		want := oneTo(n)
		if !slices.Equal(got, want) || cap(got) != n {
			t.Errorf("WithCapacity(%d) = %v with cap %d, want %v with cap %d", n, got, cap(got), want, n)
		}
	}
}

func oneTo(n int) []int {
	var s []int
	for i := 1; i <= n; i++ {
		s = append(s, i)
	}
	return s
}
