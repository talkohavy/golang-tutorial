package loops

import (
	"slices"
	"testing"
)

func TestSumTo(t *testing.T) {
	tests := []struct{ n, want int }{{5, 15}, {1, 1}, {100, 5050}, {0, 0}, {-3, 0}}
	for _, tt := range tests {
		if got := SumTo(tt.n); got != tt.want {
			t.Errorf("SumTo(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestFactorial(t *testing.T) {
	tests := []struct{ n, want int }{{0, 1}, {1, 1}, {5, 120}, {10, 3628800}}
	for _, tt := range tests {
		if got := Factorial(tt.n); got != tt.want {
			t.Errorf("Factorial(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestCountDown(t *testing.T) {
	tests := []struct {
		n    int
		want []int
	}{
		{3, []int{3, 2, 1}}, {1, []int{1}}, {5, []int{5, 4, 3, 2, 1}}, {0, []int{}},
	}
	for _, tt := range tests {
		if got := CountDown(tt.n); !slices.Equal(got, tt.want) {
			t.Errorf("CountDown(%d) = %v, want %v", tt.n, got, tt.want)
		}
	}
}

func TestSumSlice(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3}, 6}, {[]int{}, 0}, {[]int{-5, 5, 10}, 10}, {[]int{42}, 42},
	}
	for _, tt := range tests {
		if got := SumSlice(tt.nums); got != tt.want {
			t.Errorf("SumSlice(%v) = %d, want %d", tt.nums, got, tt.want)
		}
	}
}

func TestCountVowels(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"Hello World", 3}, {"AEIOU", 5}, {"rhythm", 0}, {"", 0}, {"Gopher Go", 3},
	}
	for _, tt := range tests {
		if got := CountVowels(tt.s); got != tt.want {
			t.Errorf("CountVowels(%q) = %d, want %d", tt.s, got, tt.want)
		}
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{-7, false}, {0, false}, {1, false}, {2, true}, {3, true}, {4, false},
		{9, false}, {17, true}, {25, false}, {97, true},
	}
	for _, tt := range tests {
		if got := IsPrime(tt.n); got != tt.want {
			t.Errorf("IsPrime(%d) = %t, want %t", tt.n, got, tt.want)
		}
	}
}

func TestFibonacci(t *testing.T) {
	tests := []struct {
		n    int
		want []int
	}{
		{0, []int{}}, {1, []int{0}}, {2, []int{0, 1}}, {7, []int{0, 1, 1, 2, 3, 5, 8}},
	}
	for _, tt := range tests {
		if got := Fibonacci(tt.n); !slices.Equal(got, tt.want) {
			t.Errorf("Fibonacci(%d) = %v, want %v", tt.n, got, tt.want)
		}
	}
}

func TestMultiplicationTable(t *testing.T) {
	want := [][]int{{1, 2, 3}, {2, 4, 6}, {3, 6, 9}}
	got := MultiplicationTable(3)
	if !slices.EqualFunc(got, want, slices.Equal[[]int]) {
		t.Errorf("MultiplicationTable(3) = %v, want %v", got, want)
	}
	if got := MultiplicationTable(1); len(got) != 1 || len(got[0]) != 1 || got[0][0] != 1 {
		t.Errorf("MultiplicationTable(1) = %v, want [[1]]", got)
	}
}

func TestIndexOfFirstNegative(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 0, -2, -5}, 2}, {[]int{-1}, 0}, {[]int{1, 2, 3}, -1}, {[]int{}, -1},
	}
	for _, tt := range tests {
		if got := IndexOfFirstNegative(tt.nums); got != tt.want {
			t.Errorf("IndexOfFirstNegative(%v) = %d, want %d", tt.nums, got, tt.want)
		}
	}
}

func TestSumOfEvens(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 4}, 6}, {[]int{1, 3, 5}, 0}, {[]int{-2, -4, 1}, -6}, {[]int{10}, 10},
	}
	for _, tt := range tests {
		if got := SumOfEvens(tt.nums); got != tt.want {
			t.Errorf("SumOfEvens(%v) = %d, want %d", tt.nums, got, tt.want)
		}
	}
}

func TestCollatzSteps(t *testing.T) {
	tests := []struct{ n, want int }{{1, 0}, {2, 1}, {6, 8}, {27, 111}}
	for _, tt := range tests {
		if got := CollatzSteps(tt.n); got != tt.want {
			t.Errorf("CollatzSteps(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestStaircase(t *testing.T) {
	tests := []struct {
		n    int
		want []string
	}{
		{3, []string{"*", "**", "***"}}, {1, []string{"*"}}, {0, []string{}},
	}
	for _, tt := range tests {
		if got := Staircase(tt.n); !slices.Equal(got, tt.want) {
			t.Errorf("Staircase(%d) = %q, want %q", tt.n, got, tt.want)
		}
	}
}

func TestFindPair(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		i, j   int
	}{
		{[]int{2, 7, 11, 15}, 9, 0, 1},
		{[]int{3, 2, 4}, 6, 1, 2},
		{[]int{1, 5, 1, 5}, 6, 0, 1},
		{[]int{1, 2, 3}, 100, -1, -1},
		{[]int{}, 1, -1, -1},
	}
	for _, tt := range tests {
		i, j := FindPair(tt.nums, tt.target)
		if i != tt.i || j != tt.j {
			t.Errorf("FindPair(%v, %d) = %d, %d, want %d, %d", tt.nums, tt.target, i, j, tt.i, tt.j)
		}
	}
}

func TestDigitSum(t *testing.T) {
	tests := []struct{ n, want int }{{1234, 10}, {-56, 11}, {0, 0}, {9, 9}, {1000, 1}}
	for _, tt := range tests {
		if got := DigitSum(tt.n); got != tt.want {
			t.Errorf("DigitSum(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestSquares(t *testing.T) {
	tests := []struct {
		n    int
		want []int
	}{
		{4, []int{0, 1, 4, 9}}, {1, []int{0}}, {0, []int{}},
	}
	for _, tt := range tests {
		if got := Squares(tt.n); !slices.Equal(got, tt.want) {
			t.Errorf("Squares(%d) = %v, want %v", tt.n, got, tt.want)
		}
	}
}
