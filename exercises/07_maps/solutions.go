//go:build solutions

// SPOILER ALERT: reference solutions. Try the exercises first!

package mapsx

import (
	"maps"
	"slices"
	"strings"
)

func WordCount(s string) map[string]int {
	counts := map[string]int{}
	for _, w := range strings.Fields(s) {
		counts[strings.ToLower(w)]++
	}
	return counts
}

func GetOrDefault(m map[string]int, key string, def int) int {
	if v, ok := m[key]; ok {
		return v
	}
	return def
}

func Invert(m map[string]int) map[int]string {
	result := make(map[int]string, len(m))
	for k, v := range m {
		result[v] = k
	}
	return result
}

func SortedKeys(m map[string]int) []string {
	return slices.Sorted(maps.Keys(m))
}

func MergeSum(a, b map[string]int) map[string]int {
	result := maps.Clone(a)
	if result == nil {
		result = map[string]int{}
	}
	for k, v := range b {
		result[k] += v
	}
	return result
}

func GroupByLength(words []string) map[int][]string {
	groups := map[int][]string{}
	for _, w := range words {
		groups[len(w)] = append(groups[len(w)], w)
	}
	return groups
}

func Intersection(a, b []int) []int {
	inA := map[int]struct{}{}
	for _, n := range a {
		inA[n] = struct{}{}
	}
	both := map[int]struct{}{}
	for _, n := range b {
		if _, ok := inA[n]; ok {
			both[n] = struct{}{}
		}
	}
	return slices.Sorted(maps.Keys(both))
}

func MostFrequent(nums []int) int {
	counts := map[int]int{}
	for _, n := range nums {
		counts[n]++
	}
	best, bestCount := 0, 0
	for n, c := range counts {
		if c > bestCount || (c == bestCount && n < best) {
			best, bestCount = n, c
		}
	}
	return best
}

func IsAnagram(a, b string) bool {
	counts := map[rune]int{}
	for _, r := range strings.ToLower(a) {
		if r != ' ' {
			counts[r]++
		}
	}
	for _, r := range strings.ToLower(b) {
		if r != ' ' {
			counts[r]--
		}
	}
	for _, c := range counts {
		if c != 0 {
			return false
		}
	}
	return true
}

func RemoveKeys(m map[string]int, keys ...string) {
	for _, k := range keys {
		delete(m, k)
	}
}

func CharCount(s string) map[rune]int {
	counts := map[rune]int{}
	for _, r := range s {
		counts[r]++
	}
	return counts
}

func FirstNonRepeating(s string) (rune, bool) {
	counts := CharCount(s)
	for _, r := range s {
		if counts[r] == 1 {
			return r, true
		}
	}
	return 0, false
}

func Averages(grades map[string][]int) map[string]float64 {
	result := make(map[string]float64, len(grades))
	for name, gs := range grades {
		if len(gs) == 0 {
			result[name] = 0
			continue
		}
		sum := 0
		for _, g := range gs {
			sum += g
		}
		result[name] = float64(sum) / float64(len(gs))
	}
	return result
}

func AddStock(inventory map[string]map[string]int, warehouse, item string, qty int) {
	if inventory[warehouse] == nil {
		inventory[warehouse] = map[string]int{}
	}
	inventory[warehouse][item] += qty
}

func HasDuplicates(items []string) bool {
	seen := map[string]struct{}{}
	for _, it := range items {
		if _, ok := seen[it]; ok {
			return true
		}
		seen[it] = struct{}{}
	}
	return false
}
