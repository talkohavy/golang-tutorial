package mapsx

import (
	"maps"
	"slices"
	"testing"
)

func TestWordCount(t *testing.T) {
	tests := []struct {
		in   string
		want map[string]int
	}{
		{"Go is fun go", map[string]int{"go": 2, "is": 1, "fun": 1}},
		{"a a a", map[string]int{"a": 3}},
		{"", map[string]int{}},
	}
	for _, tt := range tests {
		got := WordCount(tt.in)
		if !maps.Equal(got, tt.want) {
			t.Errorf("WordCount(%q) = %v, want %v", tt.in, got, tt.want)
		}
		if got == nil {
			t.Errorf("WordCount(%q) returned a nil map, want an empty (non-nil) map", tt.in)
		}
	}
}

func TestGetOrDefault(t *testing.T) {
	m := map[string]int{"a": 1, "zero": 0}
	tests := []struct {
		key       string
		def, want int
	}{
		{"a", 99, 1}, {"zero", 99, 0}, {"missing", 99, 99}, {"missing", -1, -1},
	}
	for _, tt := range tests {
		if got := GetOrDefault(m, tt.key, tt.def); got != tt.want {
			t.Errorf("GetOrDefault(%v, %q, %d) = %d, want %d", m, tt.key, tt.def, got, tt.want)
		}
	}
}

func TestInvert(t *testing.T) {
	in := map[string]int{"a": 1, "b": 2, "c": 3}
	want := map[int]string{1: "a", 2: "b", 3: "c"}
	if got := Invert(in); !maps.Equal(got, want) {
		t.Errorf("Invert(%v) = %v, want %v", in, got, want)
	}
}

func TestSortedKeys(t *testing.T) {
	in := map[string]int{"pear": 1, "apple": 2, "fig": 3, "banana": 4}
	want := []string{"apple", "banana", "fig", "pear"}
	if got := SortedKeys(in); !slices.Equal(got, want) {
		t.Errorf("SortedKeys(%v) = %q, want %q", in, got, want)
	}
	if got := SortedKeys(map[string]int{}); len(got) != 0 {
		t.Errorf("SortedKeys(empty) = %q, want []", got)
	}
}

func TestMergeSum(t *testing.T) {
	a := map[string]int{"x": 1, "y": 2}
	b := map[string]int{"y": 10, "z": 5}
	want := map[string]int{"x": 1, "y": 12, "z": 5}
	if got := MergeSum(a, b); !maps.Equal(got, want) {
		t.Errorf("MergeSum(%v, %v) = %v, want %v", a, b, got, want)
	}
	if !maps.Equal(a, map[string]int{"x": 1, "y": 2}) || !maps.Equal(b, map[string]int{"y": 10, "z": 5}) {
		t.Errorf("MergeSum modified its inputs: a=%v b=%v", a, b)
	}
	if got := MergeSum(nil, map[string]int{"k": 1}); !maps.Equal(got, map[string]int{"k": 1}) {
		t.Errorf("MergeSum(nil, {k:1}) = %v, want map[k:1]", got)
	}
}

func TestGroupByLength(t *testing.T) {
	in := []string{"go", "is", "fun", "a", "map", "ok"}
	want := map[int][]string{1: {"a"}, 2: {"go", "is", "ok"}, 3: {"fun", "map"}}
	got := GroupByLength(in)
	if !maps.EqualFunc(got, want, slices.Equal[[]string]) {
		t.Errorf("GroupByLength(%q) = %v, want %v", in, got, want)
	}
}

func TestIntersection(t *testing.T) {
	tests := []struct{ a, b, want []int }{
		{[]int{1, 2, 2, 3}, []int{2, 3, 4}, []int{2, 3}},
		{[]int{5, 1, 9}, []int{9, 5, 5}, []int{5, 9}},
		{[]int{1, 2}, []int{3, 4}, []int{}},
	}
	for _, tt := range tests {
		if got := Intersection(tt.a, tt.b); !slices.Equal(got, tt.want) {
			t.Errorf("Intersection(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestMostFrequent(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{[]int{1, 3, 3, 2, 1, 3}, 3}, {[]int{7}, 7}, {[]int{4, 2, 4, 2}, 2}, {[]int{-1, -1, 5}, -1},
	}
	for _, tt := range tests {
		if got := MostFrequent(tt.in); got != tt.want {
			t.Errorf("MostFrequent(%v) = %d, want %d", tt.in, got, tt.want)
		}
	}
}

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		a, b string
		want bool
	}{
		{"Listen", "Silent", true}, {"Dormitory", "dirty room", true},
		{"hello", "world", false}, {"aab", "abb", false}, {"abc", "abcd", false},
	}
	for _, tt := range tests {
		if got := IsAnagram(tt.a, tt.b); got != tt.want {
			t.Errorf("IsAnagram(%q, %q) = %t, want %t", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestRemoveKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	RemoveKeys(m, "a", "c", "nope")
	if want := map[string]int{"b": 2}; !maps.Equal(m, want) {
		t.Errorf("after RemoveKeys(m, a, c, nope) m = %v, want %v", m, want)
	}
}

func TestCharCount(t *testing.T) {
	want := map[rune]int{'h': 1, 'e': 1, 'l': 2, 'o': 1}
	if got := CharCount("hello"); !maps.Equal(got, want) {
		t.Errorf("CharCount(%q) = %v, want %v", "hello", got, want)
	}
	want = map[rune]int{'世': 2, '!': 1}
	if got := CharCount("世世!"); !maps.Equal(got, want) {
		t.Errorf("CharCount(%q) = %v, want %v", "世世!", got, want)
	}
}

func TestFirstNonRepeating(t *testing.T) {
	tests := []struct {
		in   string
		want rune
		ok   bool
	}{
		{"swiss", 'w', true}, {"aabbc", 'c', true}, {"x", 'x', true}, {"aabb", 0, false}, {"", 0, false},
	}
	for _, tt := range tests {
		got, ok := FirstNonRepeating(tt.in)
		if got != tt.want || ok != tt.ok {
			t.Errorf("FirstNonRepeating(%q) = %q, %t, want %q, %t", tt.in, got, ok, tt.want, tt.ok)
		}
	}
}

func TestAverages(t *testing.T) {
	in := map[string][]int{"ann": {90, 80}, "bob": {}, "cat": {70, 75, 77}}
	want := map[string]float64{"ann": 85, "bob": 0, "cat": 74}
	if got := Averages(in); !maps.Equal(got, want) {
		t.Errorf("Averages(%v) = %v, want %v", in, got, want)
	}
}

func TestAddStock(t *testing.T) {
	inv := map[string]map[string]int{}
	AddStock(inv, "TLV", "apple", 3)
	AddStock(inv, "TLV", "apple", 2)
	AddStock(inv, "TLV", "pear", 1)
	AddStock(inv, "NYC", "apple", 7)
	want := map[string]map[string]int{
		"TLV": {"apple": 5, "pear": 1},
		"NYC": {"apple": 7},
	}
	if !maps.EqualFunc(inv, want, maps.Equal[map[string]int]) {
		t.Errorf("inventory = %v, want %v", inv, want)
	}
}

func TestHasDuplicates(t *testing.T) {
	tests := []struct {
		in   []string
		want bool
	}{
		{[]string{"a", "b", "a"}, true}, {[]string{"a", "b", "c"}, false}, {[]string{}, false}, {[]string{"x", "x"}, true},
	}
	for _, tt := range tests {
		if got := HasDuplicates(tt.in); got != tt.want {
			t.Errorf("HasDuplicates(%q) = %t, want %t", tt.in, got, tt.want)
		}
	}
}
