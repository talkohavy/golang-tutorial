//go:build solutions

// SPOILER ALERT: reference solutions. Try the challenges first!

package challenges

import (
	"cmp"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func FizzBuzz(n int) []string {
	result := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			result = append(result, "FizzBuzz")
		case i%3 == 0:
			result = append(result, "Fizz")
		case i%5 == 0:
			result = append(result, "Buzz")
		default:
			result = append(result, strconv.Itoa(i))
		}
	}
	return result
}

func TwoSum(nums []int, target int) (int, int) {
	seen := map[int]int{}
	for j, n := range nums {
		if i, ok := seen[target-n]; ok {
			return i, j
		}
		if _, ok := seen[n]; !ok {
			seen[n] = j
		}
	}
	return -1, -1
}

var romanValues = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func RomanToInt(s string) int {
	total := 0
	for i := 0; i < len(s); i++ {
		v := romanValues[s[i]]
		if i+1 < len(s) && v < romanValues[s[i+1]] {
			total -= v
		} else {
			total += v
		}
	}
	return total
}

func IntToRoman(n int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var b strings.Builder
	for i, v := range values {
		for n >= v {
			b.WriteString(symbols[i])
			n -= v
		}
	}
	return b.String()
}

func IsBalanced(s string) bool {
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}
	var stack []rune
	for _, r := range s {
		switch r {
		case '(', '[', '{':
			stack = append(stack, r)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[r] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	best, lowest := 0, prices[0]
	for _, p := range prices[1:] {
		best = max(best, p-lowest)
		lowest = min(lowest, p)
	}
	return best
}

func MissingNumber(nums []int) int {
	n := len(nums)
	expected := n * (n + 1) / 2
	for _, v := range nums {
		expected -= v
	}
	return expected
}

func TopWords(text string, n int) []string {
	counts := map[string]int{}
	for _, w := range strings.Fields(strings.ToLower(text)) {
		counts[w]++
	}
	words := make([]string, 0, len(counts))
	for w := range counts {
		words = append(words, w)
	}
	slices.SortFunc(words, func(a, b string) int {
		return cmp.Or(cmp.Compare(counts[b], counts[a]), cmp.Compare(a, b))
	})
	return words[:min(n, len(words))]
}

func BinarySearch(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] < target:
			lo = mid + 1
		default:
			hi = mid - 1
		}
	}
	return -1
}

func MergeSorted(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	result = append(result, a[i:]...)
	return append(result, b[j:]...)
}

func SecondLargest(nums []int) (int, bool) {
	var first, second int
	hasFirst, hasSecond := false, false
	for _, n := range nums {
		switch {
		case !hasFirst || n > first:
			if hasFirst {
				second, hasSecond = first, true
			}
			first, hasFirst = n, true
		case n < first && (!hasSecond || n > second):
			second, hasSecond = n, true
		}
	}
	return second, hasSecond
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b int) int {
	return a / GCD(a, b) * b
}

func PascalTriangle(n int) [][]int {
	rows := make([][]int, 0, n)
	for i := range n {
		row := make([]int, i+1)
		row[0], row[i] = 1, 1
		for j := 1; j < i; j++ {
			row[j] = rows[i-1][j-1] + rows[i-1][j]
		}
		rows = append(rows, row)
	}
	return rows
}

func SpiralOrder(matrix [][]int) []int {
	result := []int{}
	if len(matrix) == 0 {
		return result
	}
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	for top <= bottom && left <= right {
		for c := left; c <= right; c++ {
			result = append(result, matrix[top][c])
		}
		top++
		for r := top; r <= bottom; r++ {
			result = append(result, matrix[r][right])
		}
		right--
		if top <= bottom {
			for c := right; c >= left; c-- {
				result = append(result, matrix[bottom][c])
			}
			bottom--
		}
		if left <= right {
			for r := bottom; r >= top; r-- {
				result = append(result, matrix[r][left])
			}
			left++
		}
	}
	return result
}

func RunLengthDecode(s string) string {
	var b strings.Builder
	runes := []rune(s)
	for i := 0; i < len(runes); {
		ch := runes[i]
		i++
		start := i
		for i < len(runes) && unicode.IsDigit(runes[i]) {
			i++
		}
		count, _ := strconv.Atoi(string(runes[start:i]))
		b.WriteString(strings.Repeat(string(ch), count))
	}
	return b.String()
}

func LongestCommonPrefix(words []string) string {
	if len(words) == 0 {
		return ""
	}
	prefix := words[0]
	for _, w := range words[1:] {
		for !strings.HasPrefix(w, prefix) {
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}

func RotateMatrix(m [][]int) [][]int {
	n := len(m)
	result := make([][]int, n)
	for r := range n {
		result[r] = make([]int, n)
		for c := range n {
			result[r][c] = m[n-1-c][r]
		}
	}
	return result
}

func WordWrap(text string, width int) []string {
	lines := []string{}
	current := ""
	for _, w := range strings.Fields(text) {
		switch {
		case current == "":
			current = w
		case len(current)+1+len(w) <= width:
			current += " " + w
		default:
			lines = append(lines, current)
			current = w
		}
	}
	if current != "" {
		lines = append(lines, current)
	}
	return lines
}
