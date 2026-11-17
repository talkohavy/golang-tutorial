package challenges

import (
	"slices"
	"strings"
	"testing"
)

func equal2D(a, b [][]int) bool {
	return slices.EqualFunc(a, b, slices.Equal[[]int])
}

func TestFizzBuzz(t *testing.T) {
	want := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	if got := FizzBuzz(15); !slices.Equal(got, want) {
		t.Errorf("FizzBuzz(15) = %q, want %q", got, want)
	}
	if got := FizzBuzz(0); len(got) != 0 {
		t.Errorf("FizzBuzz(0) = %q, want []", got)
	}
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		i, j   int
	}{
		{[]int{2, 7, 11, 15}, 9, 0, 1},
		{[]int{3, 2, 4}, 6, 1, 2},
		{[]int{3, 3}, 6, 0, 1},
		{[]int{1, 2, 3}, 7, -1, -1},
	}
	for _, tt := range tests {
		i, j := TwoSum(tt.nums, tt.target)
		if i != tt.i || j != tt.j {
			t.Errorf("TwoSum(%v, %d) = %d, %d, want %d, %d", tt.nums, tt.target, i, j, tt.i, tt.j)
		}
	}
}

var romanTests = []struct {
	roman string
	n     int
}{
	{"III", 3}, {"IV", 4}, {"IX", 9}, {"LVIII", 58}, {"XL", 40}, {"MCMXCIV", 1994}, {"MMMCMXCIX", 3999},
}

func TestRomanToInt(t *testing.T) {
	for _, tt := range romanTests {
		if got := RomanToInt(tt.roman); got != tt.n {
			t.Errorf("RomanToInt(%q) = %d, want %d", tt.roman, got, tt.n)
		}
	}
}

func TestIntToRoman(t *testing.T) {
	for _, tt := range romanTests {
		if got := IntToRoman(tt.n); got != tt.roman {
			t.Errorf("IntToRoman(%d) = %q, want %q", tt.n, got, tt.roman)
		}
	}
}

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"{[()]}()", true}, {"", true}, {"f(x[0]) { return }", true},
		{"([)]", false}, {"((", false}, {"())", false}, {"}", false},
	}
	for _, tt := range tests {
		if got := IsBalanced(tt.s); got != tt.want {
			t.Errorf("IsBalanced(%q) = %t, want %t", tt.s, got, tt.want)
		}
	}
}

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5}, {[]int{7, 6, 4, 3, 1}, 0}, {[]int{2, 4, 1, 7}, 6}, {[]int{5}, 0}, {[]int{}, 0},
	}
	for _, tt := range tests {
		if got := MaxProfit(tt.prices); got != tt.want {
			t.Errorf("MaxProfit(%v) = %d, want %d", tt.prices, got, tt.want)
		}
	}
}

func TestMissingNumber(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 0, 1}, 2}, {[]int{0, 1}, 2}, {[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8}, {[]int{1}, 0},
	}
	for _, tt := range tests {
		if got := MissingNumber(tt.nums); got != tt.want {
			t.Errorf("MissingNumber(%v) = %d, want %d", tt.nums, got, tt.want)
		}
	}
}

func TestTopWords(t *testing.T) {
	tests := []struct {
		text string
		n    int
		want []string
	}{
		{"b a c b a b", 2, []string{"b", "a"}},
		{"Go go GO rust Rust zig", 3, []string{"go", "rust", "zig"}},
		{"x y z", 2, []string{"x", "y"}},
		{"solo", 5, []string{"solo"}},
	}
	for _, tt := range tests {
		if got := TopWords(tt.text, tt.n); !slices.Equal(got, tt.want) {
			t.Errorf("TopWords(%q, %d) = %q, want %q", tt.text, tt.n, got, tt.want)
		}
	}
}

func TestBinarySearch(t *testing.T) {
	nums := []int{-5, 1, 3, 8, 13, 21, 34}
	for i, v := range nums {
		if got := BinarySearch(nums, v); got != i {
			t.Errorf("BinarySearch(%v, %d) = %d, want %d", nums, v, got, i)
		}
	}
	for _, missing := range []int{-10, 2, 100} {
		if got := BinarySearch(nums, missing); got != -1 {
			t.Errorf("BinarySearch(%v, %d) = %d, want -1", nums, missing, got)
		}
	}
	if got := BinarySearch(nil, 1); got != -1 {
		t.Errorf("BinarySearch(nil, 1) = %d, want -1", got)
	}
}

func TestMergeSorted(t *testing.T) {
	tests := []struct{ a, b, want []int }{
		{[]int{1, 4, 9}, []int{2, 3, 10}, []int{1, 2, 3, 4, 9, 10}},
		{[]int{}, []int{1, 2}, []int{1, 2}},
		{[]int{1, 1}, []int{1}, []int{1, 1, 1}},
		{[]int{5, 6}, []int{}, []int{5, 6}},
	}
	for _, tt := range tests {
		if got := MergeSorted(tt.a, tt.b); !slices.Equal(got, tt.want) {
			t.Errorf("MergeSorted(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestSecondLargest(t *testing.T) {
	tests := []struct {
		nums []int
		want int
		ok   bool
	}{
		{[]int{5, 1, 5, 3}, 3, true}, {[]int{1, 2}, 1, true}, {[]int{-1, -3, -2}, -2, true},
		{[]int{7, 7}, 0, false}, {[]int{4}, 0, false}, {[]int{}, 0, false},
	}
	for _, tt := range tests {
		got, ok := SecondLargest(tt.nums)
		if got != tt.want || ok != tt.ok {
			t.Errorf("SecondLargest(%v) = %d, %t, want %d, %t", tt.nums, got, ok, tt.want, tt.ok)
		}
	}
}

func TestGCDAndLCM(t *testing.T) {
	tests := []struct{ a, b, gcd, lcm int }{
		{12, 18, 6, 36}, {4, 6, 2, 12}, {7, 13, 1, 91}, {10, 10, 10, 10}, {1, 9, 1, 9},
	}
	for _, tt := range tests {
		if got := GCD(tt.a, tt.b); got != tt.gcd {
			t.Errorf("GCD(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.gcd)
		}
		if got := LCM(tt.a, tt.b); got != tt.lcm {
			t.Errorf("LCM(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.lcm)
		}
	}
}

func TestPascalTriangle(t *testing.T) {
	want := [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}
	if got := PascalTriangle(5); !equal2D(got, want) {
		t.Errorf("PascalTriangle(5) = %v, want %v", got, want)
	}
	if got := PascalTriangle(0); len(got) != 0 {
		t.Errorf("PascalTriangle(0) = %v, want []", got)
	}
}

func TestSpiralOrder(t *testing.T) {
	tests := []struct {
		m    [][]int
		want []int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
		{[][]int{{1}, {2}, {3}}, []int{1, 2, 3}},
		{[][]int{{1, 2, 3}}, []int{1, 2, 3}},
		{[][]int{}, []int{}},
	}
	for _, tt := range tests {
		if got := SpiralOrder(tt.m); !slices.Equal(got, tt.want) {
			t.Errorf("SpiralOrder(%v) = %v, want %v", tt.m, got, tt.want)
		}
	}
}

func TestRunLengthDecode(t *testing.T) {
	tests := []struct{ in, want string }{
		{"a3b1c2", "aaabcc"}, {"x1", "x"}, {"", ""}, {"z12", strings.Repeat("z", 12)}, {"a2b10a1", "aabbbbbbbbbba"},
	}
	for _, tt := range tests {
		if got := RunLengthDecode(tt.in); got != tt.want {
			t.Errorf("RunLengthDecode(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		words []string
		want  string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"}, {[]string{"dog", "racecar", "car"}, ""},
		{[]string{"gopher"}, "gopher"}, {[]string{"go", "gopher", "golang"}, "go"}, {nil, ""},
	}
	for _, tt := range tests {
		if got := LongestCommonPrefix(tt.words); got != tt.want {
			t.Errorf("LongestCommonPrefix(%q) = %q, want %q", tt.words, got, tt.want)
		}
	}
}

func TestRotateMatrix(t *testing.T) {
	tests := []struct{ in, want [][]int }{
		{[][]int{{1, 2}, {3, 4}}, [][]int{{3, 1}, {4, 2}}},
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}},
		{[][]int{{5}}, [][]int{{5}}},
	}
	for _, tt := range tests {
		original := [][]int{}
		for _, row := range tt.in {
			original = append(original, slices.Clone(row))
		}
		if got := RotateMatrix(tt.in); !equal2D(got, tt.want) {
			t.Errorf("RotateMatrix(%v) = %v, want %v", original, got, tt.want)
		}
		if !equal2D(tt.in, original) {
			t.Errorf("RotateMatrix modified its input: %v became %v", original, tt.in)
		}
	}
}

func TestWordWrap(t *testing.T) {
	tests := []struct {
		text  string
		width int
		want  []string
	}{
		{"the quick brown fox jumps", 10, []string{"the quick", "brown fox", "jumps"}},
		{"a b c d", 3, []string{"a b", "c d"}},
		{"supercalifragilistic is long", 8, []string{"supercalifragilistic", "is long"}},
		{"", 5, []string{}},
	}
	for _, tt := range tests {
		if got := WordWrap(tt.text, tt.width); !slices.Equal(got, tt.want) {
			t.Errorf("WordWrap(%q, %d) = %q, want %q", tt.text, tt.width, got, tt.want)
		}
	}
}
