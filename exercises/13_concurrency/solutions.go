//go:build solutions

// SPOILER ALERT: reference solutions. Try the exercises first!

package concurrency

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

func ParallelSum(nums []int, parts int) int {
	chunk := (len(nums) + parts - 1) / parts
	results := make(chan int, parts)
	var wg sync.WaitGroup
	for start := 0; start < len(nums); start += chunk {
		end := min(start+chunk, len(nums))
		wg.Go(func() {
			sum := 0
			for _, n := range nums[start:end] {
				sum += n
			}
			results <- sum
		})
	}
	wg.Wait()
	close(results)
	total := 0
	for s := range results {
		total += s
	}
	return total
}

func Generate(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range n {
			out <- i
		}
	}()
	return out
}

func SquareAll(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			out <- v * v
		}
	}()
	return out
}

func Merge(inputs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	for _, in := range inputs {
		wg.Go(func() {
			for v := range in {
				out <- v
			}
		})
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func ProcessAll(jobs []int, workers int, f func(int) int) []int {
	results := make([]int, len(jobs))
	indexes := make(chan int)
	var wg sync.WaitGroup
	for range workers {
		wg.Go(func() {
			for i := range indexes {
				results[i] = f(jobs[i])
			}
		})
	}
	for i := range jobs {
		indexes <- i
	}
	close(indexes)
	wg.Wait()
	return results
}

type SafeCounter struct {
	mu     sync.Mutex
	counts map[string]int
}

func NewSafeCounter() *SafeCounter {
	return &SafeCounter{counts: map[string]int{}}
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counts[key]++
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counts[key]
}

func WithTimeout(f func() int, d time.Duration) (int, bool) {
	result := make(chan int, 1)
	go func() { result <- f() }()
	select {
	case v := <-result:
		return v, true
	case <-time.After(d):
		return 0, false
	}
}

func FetchAll(urls []string, fetch func(string) string) map[string]string {
	results := make(map[string]string, len(urls))
	var mu sync.Mutex
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Go(func() {
			body := fetch(url)
			mu.Lock()
			results[url] = body
			mu.Unlock()
		})
	}
	wg.Wait()
	return results
}

func SumUntilDone(ctx context.Context, in <-chan int) int {
	sum := 0
	for {
		select {
		case <-ctx.Done():
			return sum
		case v, ok := <-in:
			if !ok {
				return sum
			}
			sum += v
		}
	}
}

func FilledBuffer(n int) <-chan int {
	ch := make(chan int, n)
	for i := 1; i <= n; i++ {
		ch <- i
	}
	close(ch)
	return ch
}

func ParallelCount(n int) int {
	var count atomic.Int64
	var wg sync.WaitGroup
	for range n {
		wg.Go(func() { count.Add(1) })
	}
	wg.Wait()
	return int(count.Load())
}
