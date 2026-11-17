//go:build solutions

// SPOILER ALERT: reference solutions. Try the exercises first!

package loops

import "strings"

func SumTo(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func Factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func CountDown(n int) []int {
	result := []int{}
	for i := n; i >= 1; i-- {
		result = append(result, i)
	}
	return result
}

func SumSlice(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func CountVowels(s string) int {
	count := 0
	for _, r := range strings.ToLower(s) {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Fibonacci(n int) []int {
	result := []int{}
	a, b := 0, 1
	for range n {
		result = append(result, a)
		a, b = b, a+b
	}
	return result
}

func MultiplicationTable(n int) [][]int {
	table := make([][]int, n)
	for i := range n {
		table[i] = make([]int, n)
		for j := range n {
			table[i][j] = (i + 1) * (j + 1)
		}
	}
	return table
}

func IndexOfFirstNegative(nums []int) int {
	for i, n := range nums {
		if n < 0 {
			return i
		}
	}
	return -1
}

func SumOfEvens(nums []int) int {
	sum := 0
	for _, n := range nums {
		if n%2 != 0 {
			continue
		}
		sum += n
	}
	return sum
}

func CollatzSteps(n int) int {
	steps := 0
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		steps++
	}
	return steps
}

func Staircase(n int) []string {
	lines := []string{}
	for i := 1; i <= n; i++ {
		line := ""
		for range i {
			line += "*"
		}
		lines = append(lines, line)
	}
	return lines
}

func FindPair(nums []int, target int) (int, int) {
	first, second := -1, -1
outer:
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				first, second = i, j
				break outer
			}
		}
	}
	return first, second
}

func DigitSum(n int) int {
	if n < 0 {
		n = -n
	}
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func Squares(n int) []int {
	result := []int{}
	for i := range n {
		result = append(result, i*i)
	}
	return result
}
