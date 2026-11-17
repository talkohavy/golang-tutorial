//go:build solutions

// SPOILER ALERT: reference solutions. Try the exercises first!

package functions

func MinMax(nums []int) (lo, hi int) {
	lo, hi = nums[0], nums[0]
	for _, n := range nums[1:] {
		if n < lo {
			lo = n
		}
		if n > hi {
			hi = n
		}
	}
	return lo, hi
}

func Sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func Apply(nums []int, f func(int) int) []int {
	result := make([]int, len(nums))
	for i, n := range nums {
		result[i] = f(n)
	}
	return result
}

func MakeMultiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

func NewCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func Compose(f, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

func Power(base, exp int) int {
	if exp == 0 {
		return 1
	}
	return base * Power(base, exp-1)
}

func Divide(a, b float64) (float64, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

func JoinWith(sep string, parts ...string) string {
	result := ""
	for i, p := range parts {
		if i > 0 {
			result += sep
		}
		result += p
	}
	return result
}

func Memoize(f func(int) int) func(int) int {
	cache := map[int]int{}
	return func(n int) int {
		if v, ok := cache[n]; ok {
			return v
		}
		v := f(n)
		cache[n] = v
		return v
	}
}

func DeferredSteps() (steps []string) {
	steps = append(steps, "start")
	defer func() { steps = append(steps, "deferred 1") }()
	defer func() { steps = append(steps, "deferred 2") }()
	steps = append(steps, "end")
	return steps
}

func ApplyN(f func(int) int, n int, x int) int {
	for range n {
		x = f(x)
	}
	return x
}

type Predicate func(int) bool

func CountIf(nums []int, p Predicate) int {
	count := 0
	for _, n := range nums {
		if p(n) {
			count++
		}
	}
	return count
}
