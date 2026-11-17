package functions

import (
	"slices"
	"testing"
)

func TestMinMax(t *testing.T) {
	tests := []struct {
		nums   []int
		lo, hi int
	}{
		{[]int{3, 1, 4, 1, 5}, 1, 5}, {[]int{7}, 7, 7}, {[]int{-2, -9, -4}, -9, -2},
	}
	for _, tt := range tests {
		lo, hi := MinMax(tt.nums)
		if lo != tt.lo || hi != tt.hi {
			t.Errorf("MinMax(%v) = %d, %d, want %d, %d", tt.nums, lo, hi, tt.lo, tt.hi)
		}
	}
}

func TestSum(t *testing.T) {
	if got := Sum(); got != 0 {
		t.Errorf("Sum() = %d, want 0", got)
	}
	if got := Sum(1, 2, 3); got != 6 {
		t.Errorf("Sum(1, 2, 3) = %d, want 6", got)
	}
	nums := []int{10, 20, 30, 40}
	if got := Sum(nums...); got != 100 {
		t.Errorf("Sum(%v...) = %d, want 100", nums, got)
	}
}

func TestApply(t *testing.T) {
	times10 := func(n int) int { return n * 10 }
	if got, want := Apply([]int{1, 2, 3}, times10), []int{10, 20, 30}; !slices.Equal(got, want) {
		t.Errorf("Apply([1 2 3], times10) = %v, want %v", got, want)
	}
	negate := func(n int) int { return -n }
	if got, want := Apply([]int{5, -1}, negate), []int{-5, 1}; !slices.Equal(got, want) {
		t.Errorf("Apply([5 -1], negate) = %v, want %v", got, want)
	}
	if got := Apply([]int{}, negate); len(got) != 0 {
		t.Errorf("Apply([], negate) = %v, want []", got)
	}
}

func TestMakeMultiplier(t *testing.T) {
	triple := MakeMultiplier(3)
	half := MakeMultiplier(-2)
	if got := triple(5); got != 15 {
		t.Errorf("MakeMultiplier(3)(5) = %d, want 15", got)
	}
	if got := half(4); got != -8 {
		t.Errorf("MakeMultiplier(-2)(4) = %d, want -8", got)
	}
}

func TestNewCounter(t *testing.T) {
	a := NewCounter()
	b := NewCounter()
	for want := 1; want <= 3; want++ {
		if got := a(); got != want {
			t.Fatalf("counter a: call #%d returned %d, want %d", want, got, want)
		}
	}
	if got := b(); got != 1 {
		t.Errorf("counter b: first call returned %d, want 1 (counters must be independent)", got)
	}
}

func TestCompose(t *testing.T) {
	addOne := func(x int) int { return x + 1 }
	double := func(x int) int { return x * 2 }
	if got := Compose(addOne, double)(5); got != 11 {
		t.Errorf("Compose(addOne, double)(5) = %d, want 11", got)
	}
	if got := Compose(double, addOne)(5); got != 12 {
		t.Errorf("Compose(double, addOne)(5) = %d, want 12", got)
	}
}

func TestPower(t *testing.T) {
	tests := []struct{ base, exp, want int }{
		{2, 10, 1024}, {3, 0, 1}, {5, 1, 5}, {-2, 3, -8},
	}
	for _, tt := range tests {
		if got := Power(tt.base, tt.exp); got != tt.want {
			t.Errorf("Power(%d, %d) = %d, want %d", tt.base, tt.exp, got, tt.want)
		}
	}
}

func TestDivide(t *testing.T) {
	if got, ok := Divide(10, 4); got != 2.5 || !ok {
		t.Errorf("Divide(10, 4) = %v, %t, want 2.5, true", got, ok)
	}
	if got, ok := Divide(-9, 3); got != -3 || !ok {
		t.Errorf("Divide(-9, 3) = %v, %t, want -3, true", got, ok)
	}
	if got, ok := Divide(1, 0); got != 0 || ok {
		t.Errorf("Divide(1, 0) = %v, %t, want 0, false", got, ok)
	}
}

func TestJoinWith(t *testing.T) {
	tests := []struct {
		sep   string
		parts []string
		want  string
	}{
		{"-", []string{"a", "b", "c"}, "a-b-c"},
		{", ", []string{"Go"}, "Go"},
		{", ", nil, ""},
		{"", []string{"x", "y"}, "xy"},
	}
	for _, tt := range tests {
		if got := JoinWith(tt.sep, tt.parts...); got != tt.want {
			t.Errorf("JoinWith(%q, %q...) = %q, want %q", tt.sep, tt.parts, got, tt.want)
		}
	}
}

func TestMemoize(t *testing.T) {
	calls := 0
	square := func(n int) int {
		calls++
		return n * n
	}
	m := Memoize(square)
	for _, n := range []int{4, 4, 5, 4, 5} {
		if got := m(n); got != n*n {
			t.Fatalf("memoized square(%d) = %d, want %d", n, got, n*n)
		}
	}
	if calls != 2 {
		t.Errorf("the original function was called %d times, want 2 (once per distinct argument)", calls)
	}
}

func TestDeferredSteps(t *testing.T) {
	want := []string{"start", "end", "deferred 2", "deferred 1"}
	if got := DeferredSteps(); !slices.Equal(got, want) {
		t.Errorf("DeferredSteps() = %q, want %q", got, want)
	}
}

func TestApplyN(t *testing.T) {
	double := func(x int) int { return x * 2 }
	inc := func(x int) int { return x + 1 }
	if got := ApplyN(double, 3, 1); got != 8 {
		t.Errorf("ApplyN(double, 3, 1) = %d, want 8", got)
	}
	if got := ApplyN(inc, 0, 7); got != 7 {
		t.Errorf("ApplyN(inc, 0, 7) = %d, want 7", got)
	}
	if got := ApplyN(inc, 5, 10); got != 15 {
		t.Errorf("ApplyN(inc, 5, 10) = %d, want 15", got)
	}
}

func TestCountIf(t *testing.T) {
	greaterThan2 := func(n int) bool { return n > 2 }
	if got := CountIf([]int{1, 2, 3, 4}, greaterThan2); got != 2 {
		t.Errorf("CountIf([1 2 3 4], >2) = %d, want 2", got)
	}
	var isEven Predicate = func(n int) bool { return n%2 == 0 }
	if got := CountIf([]int{2, 4, 6, 7}, isEven); got != 3 {
		t.Errorf("CountIf([2 4 6 7], isEven) = %d, want 3", got)
	}
}
