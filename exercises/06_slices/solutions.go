//go:build solutions

// SPOILER ALERT: reference solutions. Try the exercises first!

package slicesx

import "slices"

func MakeRange(start, end int) []int {
	result := []int{}
	for i := start; i < end; i++ {
		result = append(result, i)
	}
	return result
}

func Last(nums []int) (int, bool) {
	if len(nums) == 0 {
		return 0, false
	}
	return nums[len(nums)-1], true
}

func RemoveAt(nums []int, i int) []int {
	result := make([]int, 0, len(nums)-1)
	result = append(result, nums[:i]...)
	return append(result, nums[i+1:]...)
}

func InsertAt(nums []int, i, v int) []int {
	result := make([]int, 0, len(nums)+1)
	result = append(result, nums[:i]...)
	result = append(result, v)
	return append(result, nums[i:]...)
}

func Evens(nums []int) []int {
	result := []int{}
	for _, n := range nums {
		if n%2 == 0 {
			result = append(result, n)
		}
	}
	return result
}

func ReverseInPlace(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func Chunk(nums []int, size int) [][]int {
	chunks := [][]int{}
	for start := 0; start < len(nums); start += size {
		end := min(start+size, len(nums))
		chunks = append(chunks, nums[start:end])
	}
	return chunks
}

func Unique(nums []int) []int {
	seen := map[int]bool{}
	result := []int{}
	for _, n := range nums {
		if !seen[n] {
			seen[n] = true
			result = append(result, n)
		}
	}
	return result
}

func Flatten(matrix [][]int) []int {
	result := []int{}
	for _, row := range matrix {
		result = append(result, row...)
	}
	return result
}

func Transpose(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return [][]int{}
	}
	rows, cols := len(matrix), len(matrix[0])
	result := make([][]int, cols)
	for c := range cols {
		result[c] = make([]int, rows)
		for r := range rows {
			result[c][r] = matrix[r][c]
		}
	}
	return result
}

func RotateLeft(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	// Solution 1:
	k %= len(nums)
	result := make([]int, 0, len(nums))
	result = append(result, nums[k:]...)
	return append(result, nums[:k]...)
	// Solution 2:
	// sliceResult := []int{}
	// for i := range nums {
	// 	sliceResult = append(sliceResult, nums[(i+k)%len(nums)])
	// }
	// return sliceResult
}

func DoubledCopy(nums []int) []int {
	result := slices.Clone(nums)
	for i := range result {
		result[i] *= 2
	}
	return result
}

func SortedDesc(nums []int) []int {
	result := slices.Clone(nums)
	// Solution 2:
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})
	return result
}

func ContainsAll(haystack, needles []int) bool {
	// Solution 1: self implementation
	for _, needle := range needles {
		isFound := false
		for _, hayNeedle := range haystack {
			if needle == hayNeedle {
				isFound = true
				break
			}
		}
		if !isFound {
			return false
		}
	}
	return true
	// Solution 2: using slices.Contains
	// for _, n := range needles {
	// 	if !slices.Contains(haystack, n) {
	// 		return false
	// 	}
	// }
	// return true
}

func DoubleArray(arr [3]int) [3]int {
	for i := range arr {
		arr[i] *= 2
	}
	return arr
}

func WithCapacity(n int) []int {
	s := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		s = append(s, i)
	}
	return s
}
