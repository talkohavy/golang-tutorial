package generics

import (
	"maps"
	"slices"
	"strconv"
	"strings"
	"testing"
)

func TestMap(t *testing.T) {
	got := Map([]int{1, 2, 3}, func(n int) string { return strconv.Itoa(n * 2) })
	if want := []string{"2", "4", "6"}; !slices.Equal(got, want) {
		t.Errorf("Map(ints → doubled strings) = %q, want %q", got, want)
	}
	lengths := Map([]string{"go", "rust"}, func(s string) int { return len(s) })
	if want := []int{2, 4}; !slices.Equal(lengths, want) {
		t.Errorf("Map(strings → lengths) = %v, want %v", lengths, want)
	}
}

func TestFilter(t *testing.T) {
	got := Filter([]string{"go", "rust", "c"}, func(s string) bool { return len(s) > 1 })
	if want := []string{"go", "rust"}; !slices.Equal(got, want) {
		t.Errorf("Filter(len > 1) = %q, want %q", got, want)
	}
	evens := Filter([]int{1, 2, 3, 4}, func(n int) bool { return n%2 == 0 })
	if want := []int{2, 4}; !slices.Equal(evens, want) {
		t.Errorf("Filter(even) = %v, want %v", evens, want)
	}
}

func TestReduce(t *testing.T) {
	sum := Reduce([]int{1, 2, 3}, 0, func(acc, n int) int { return acc + n })
	if sum != 6 {
		t.Errorf("Reduce(sum) = %d, want 6", sum)
	}
	joined := Reduce([]string{"a", "b", "c"}, ">", func(acc, s string) string { return acc + s })
	if joined != ">abc" {
		t.Errorf("Reduce(concat) = %q, want %q", joined, ">abc")
	}
	totalLen := Reduce([]string{"go", "gopher"}, 0, func(acc int, s string) int { return acc + len(s) })
	if totalLen != 8 {
		t.Errorf("Reduce(total length) = %d, want 8", totalLen)
	}
}

func TestContains(t *testing.T) {
	if !Contains([]int{1, 2, 3}, 2) || Contains([]int{1, 2, 3}, 5) {
		t.Error("Contains is wrong for ints")
	}
	if !Contains([]string{"a", "b"}, "b") || Contains([]string{}, "a") {
		t.Error("Contains is wrong for strings")
	}
}

type Celsius float64

func TestSum(t *testing.T) {
	if got := Sum([]int{1, 2, 3}); got != 6 {
		t.Errorf("Sum([]int{1, 2, 3}) = %v, want 6", got)
	}
	if got := Sum([]float64{1.5, 2.5}); got != 4 {
		t.Errorf("Sum([]float64{1.5, 2.5}) = %v, want 4", got)
	}
	if got := Sum([]Celsius{20, 1.5}); got != 21.5 {
		t.Errorf("Sum([]Celsius{20, 1.5}) = %v, want 21.5", got)
	}
}

func TestMax(t *testing.T) {
	if got, ok := Max([]int{3, 9, -1}); got != 9 || !ok {
		t.Errorf("Max([3 9 -1]) = %v, %t, want 9, true", got, ok)
	}
	if got, ok := Max([]string{"pear", "zebra", "apple"}); got != "zebra" || !ok {
		t.Errorf("Max(strings) = %q, %t, want \"zebra\", true", got, ok)
	}
	if got, ok := Max([]float64{-2.5, -1.5}); got != -1.5 || !ok {
		t.Errorf("Max([-2.5 -1.5]) = %v, %t, want -1.5, true", got, ok)
	}
	if got, ok := Max([]int{}); got != 0 || ok {
		t.Errorf("Max([]) = %v, %t, want 0, false", got, ok)
	}
}

func TestKeys(t *testing.T) {
	got := Keys(map[string]int{"b": 2, "a": 1, "c": 3})
	slices.Sort(got)
	if want := []string{"a", "b", "c"}; !slices.Equal(got, want) {
		t.Errorf("Keys(...) (sorted) = %q, want %q", got, want)
	}
	ints := Keys(map[int]bool{5: true, 1: false})
	slices.Sort(ints)
	if want := []int{1, 5}; !slices.Equal(ints, want) {
		t.Errorf("Keys(...) (sorted) = %v, want %v", ints, want)
	}
}

func TestGroupBy(t *testing.T) {
	got := GroupBy([]string{"apple", "avocado", "banana", "cherry", "blueberry"}, func(s string) byte { return s[0] })
	want := map[byte][]string{'a': {"apple", "avocado"}, 'b': {"banana", "blueberry"}, 'c': {"cherry"}}
	if !maps.EqualFunc(got, want, slices.Equal[[]string]) {
		t.Errorf("GroupBy(first letter) = %q, want %q", got, want)
	}
	parity := GroupBy([]int{1, 2, 3, 4, 5}, func(n int) bool { return n%2 == 0 })
	if !slices.Equal(parity[true], []int{2, 4}) || !slices.Equal(parity[false], []int{1, 3, 5}) {
		t.Errorf("GroupBy(is even) = %v", parity)
	}
}

func TestZip(t *testing.T) {
	got := Zip([]string{"a", "b", "c"}, []int{1, 2})
	want := []Pair[string, int]{{"a", 1}, {"b", 2}}
	if !slices.Equal(got, want) {
		t.Errorf("Zip([a b c], [1 2]) = %v, want %v", got, want)
	}
	if got := Zip([]int{}, []bool{true}); len(got) != 0 {
		t.Errorf("Zip([], [true]) = %v, want []", got)
	}
}

func TestStack(t *testing.T) {
	var s Stack[string]
	if _, ok := s.Pop(); ok {
		t.Fatal("Pop() on an empty stack returned ok = true")
	}
	s.Push("a")
	s.Push("b")
	if s.Len() != 2 {
		t.Fatalf("Len() = %d after 2 pushes, want 2", s.Len())
	}
	if v, ok := s.Pop(); v != "b" || !ok {
		t.Errorf("Pop() = %q, %t, want \"b\", true", v, ok)
	}
	if v, ok := s.Pop(); v != "a" || !ok {
		t.Errorf("Pop() = %q, %t, want \"a\", true", v, ok)
	}
	if s.Len() != 0 {
		t.Errorf("Len() = %d after popping everything, want 0", s.Len())
	}

	nums := &Stack[int]{}
	nums.Push(42)
	if v, _ := nums.Pop(); v != 42 {
		t.Errorf("Stack[int]: Pop() = %d, want 42", v)
	}
}

func TestSet(t *testing.T) {
	s := NewSet("go", "rust", "go")
	if s.Len() != 2 || !s.Has("go") || !s.Has("rust") || s.Has("java") {
		t.Fatalf("NewSet(go, rust, go): Len() = %d, Has(go) = %t, Has(rust) = %t, Has(java) = %t",
			s.Len(), s.Has("go"), s.Has("rust"), s.Has("java"))
	}
	s.Add("zig")
	s.Add("zig")
	if s.Len() != 3 || !s.Has("zig") {
		t.Errorf("after Add(zig) twice: Len() = %d, Has(zig) = %t, want 3, true", s.Len(), s.Has("zig"))
	}
	u := NewSet(1, 2).Union(NewSet(2, 3))
	if u.Len() != 3 || !u.Has(1) || !u.Has(2) || !u.Has(3) {
		t.Errorf("{1 2} ∪ {2 3} has Len() = %d, want 3 containing 1, 2, 3", u.Len())
	}
	empty := NewSet[int]()
	empty.Add(7)
	if !empty.Has(7) {
		t.Error("NewSet[int]() then Add(7): Has(7) = false")
	}
}

func TestReverse(t *testing.T) {
	in := []string{"a", "b", "c"}
	if got := Reverse(in); !slices.Equal(got, []string{"c", "b", "a"}) {
		t.Errorf("Reverse([a b c]) = %q, want [c b a]", got)
	}
	if !slices.Equal(in, []string{"a", "b", "c"}) {
		t.Errorf("Reverse modified its input: %q", in)
	}
	if got := Reverse([]int{}); len(got) != 0 {
		t.Errorf("Reverse([]) = %v, want []", got)
	}
}

func TestUnique(t *testing.T) {
	if got := Unique([]int{3, 1, 3, 2, 1}); !slices.Equal(got, []int{3, 1, 2}) {
		t.Errorf("Unique([3 1 3 2 1]) = %v, want [3 1 2]", got)
	}
	words := strings.Fields("the cat and the hat and the bat")
	if got, want := Unique(words), []string{"the", "cat", "and", "hat", "bat"}; !slices.Equal(got, want) {
		t.Errorf("Unique(words) = %q, want %q", got, want)
	}
}

func TestOrDefault(t *testing.T) {
	if got := OrDefault("", "anonymous"); got != "anonymous" {
		t.Errorf("OrDefault(\"\", \"anonymous\") = %q", got)
	}
	if got := OrDefault("bob", "anonymous"); got != "bob" {
		t.Errorf("OrDefault(\"bob\", \"anonymous\") = %q", got)
	}
	if got := OrDefault(0, 8080); got != 8080 {
		t.Errorf("OrDefault(0, 8080) = %d", got)
	}
	if got := OrDefault(3000, 8080); got != 3000 {
		t.Errorf("OrDefault(3000, 8080) = %d", got)
	}
}
