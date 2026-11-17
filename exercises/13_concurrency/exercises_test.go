package concurrency

import (
	"context"
	"fmt"
	"maps"
	"slices"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// collect reads ch until it's closed, failing the test if that takes too long
// (for example when the channel is nil or never closed).
func collect(t *testing.T, ch <-chan int) []int {
	t.Helper()
	var out []int
	timeout := time.After(time.Second)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				return out
			}
			out = append(out, v)
		case <-timeout:
			t.Fatalf("timed out waiting for the channel to be closed (received so far: %v)", out)
		}
	}
}

// within runs f in a goroutine and fails the test if it doesn't return in d.
func within(t *testing.T, d time.Duration, what string, f func()) {
	t.Helper()
	done := make(chan struct{})
	go func() {
		defer close(done)
		f()
	}()
	select {
	case <-done:
	case <-time.After(d):
		t.Fatalf("%s did not return within %v (deadlock?)", what, d)
	}
}

func source(values ...int) <-chan int {
	ch := make(chan int, len(values))
	for _, v := range values {
		ch <- v
	}
	close(ch)
	return ch
}

func TestParallelSum(t *testing.T) {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i + 1
	}
	tests := []struct {
		nums  []int
		parts int
		want  int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 3, 21}, {[]int{1, 2, 3, 4, 5}, 2, 15}, {nums, 7, 500500}, {[]int{5}, 4, 5},
	}
	for _, tt := range tests {
		var got int
		within(t, time.Second, "ParallelSum", func() { got = ParallelSum(tt.nums, tt.parts) })
		if got != tt.want {
			t.Errorf("ParallelSum(%d numbers, %d) = %d, want %d", len(tt.nums), tt.parts, got, tt.want)
		}
	}
}

func TestGenerate(t *testing.T) {
	if got := collect(t, Generate(5)); !slices.Equal(got, []int{0, 1, 2, 3, 4}) {
		t.Errorf("Generate(5) produced %v, want [0 1 2 3 4]", got)
	}
	if got := collect(t, Generate(0)); len(got) != 0 {
		t.Errorf("Generate(0) produced %v, want nothing", got)
	}
}

func TestSquareAll(t *testing.T) {
	if got := collect(t, SquareAll(source(1, 2, 3, -4))); !slices.Equal(got, []int{1, 4, 9, 16}) {
		t.Errorf("SquareAll(1, 2, 3, -4) produced %v, want [1 4 9 16]", got)
	}
}

func TestMerge(t *testing.T) {
	got := collect(t, Merge(source(1, 2), source(3), source(), source(4, 5, 6)))
	slices.Sort(got)
	if want := []int{1, 2, 3, 4, 5, 6}; !slices.Equal(got, want) {
		t.Errorf("Merge(...) produced (sorted) %v, want %v", got, want)
	}
}

func TestProcessAll(t *testing.T) {
	var active, maxActive atomic.Int32
	slowDouble := func(n int) int {
		cur := active.Add(1)
		for {
			old := maxActive.Load()
			if cur <= old || maxActive.CompareAndSwap(old, cur) {
				break
			}
		}
		time.Sleep(50 * time.Millisecond)
		active.Add(-1)
		return n * 2
	}
	jobs := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var got []int
	start := time.Now()
	within(t, 2*time.Second, "ProcessAll", func() { got = ProcessAll(jobs, 4, slowDouble) })
	elapsed := time.Since(start)

	if want := []int{2, 4, 6, 8, 10, 12, 14, 16}; !slices.Equal(got, want) {
		t.Errorf("ProcessAll(...) = %v, want %v (in the same order as jobs)", got, want)
	}
	if m := maxActive.Load(); m > 4 {
		t.Errorf("up to %d jobs ran at the same time, but only 4 workers were allowed", m)
	} else if m < 2 {
		t.Errorf("jobs never ran concurrently (max %d at a time)", m)
	}
	if elapsed > 300*time.Millisecond {
		t.Errorf("took %v; with 4 workers it should take about 100ms", elapsed)
	}
}

func TestSafeCounter(t *testing.T) {
	c := NewSafeCounter()
	var wg sync.WaitGroup
	for i := range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc("hits")
			if i%2 == 0 {
				c.Inc("even")
			}
		}()
	}
	wg.Wait()
	if got := c.Value("hits"); got != 100 {
		t.Errorf("Value(hits) = %d, want 100", got)
	}
	if got := c.Value("even"); got != 50 {
		t.Errorf("Value(even) = %d, want 50", got)
	}
	if got := c.Value("nothing"); got != 0 {
		t.Errorf("Value(nothing) = %d, want 0", got)
	}
}

func TestWithTimeout(t *testing.T) {
	fast := func() int { return 42 }
	slow := func() int { time.Sleep(500 * time.Millisecond); return 1 }

	if v, ok := WithTimeout(fast, 200*time.Millisecond); v != 42 || !ok {
		t.Errorf("WithTimeout(fast, 200ms) = %d, %t, want 42, true", v, ok)
	}
	start := time.Now()
	v, ok := WithTimeout(slow, 50*time.Millisecond)
	if v != 0 || ok {
		t.Errorf("WithTimeout(slow, 50ms) = %d, %t, want 0, false", v, ok)
	}
	if elapsed := time.Since(start); elapsed > 300*time.Millisecond {
		t.Errorf("WithTimeout(slow, 50ms) took %v: it should give up after ~50ms", elapsed)
	}
}

func TestFetchAll(t *testing.T) {
	var urls []string
	for i := range 10 {
		urls = append(urls, fmt.Sprintf("https://example.com/%d", i))
	}
	fetch := func(url string) string {
		time.Sleep(100 * time.Millisecond)
		return "body of " + url
	}
	var got map[string]string
	start := time.Now()
	within(t, 3*time.Second, "FetchAll", func() { got = FetchAll(urls, fetch) })
	elapsed := time.Since(start)

	want := map[string]string{}
	for _, u := range urls {
		want[u] = "body of " + u
	}
	if !maps.Equal(got, want) {
		t.Errorf("FetchAll returned %d results, want %d: %v", len(got), len(want), got)
	}
	if elapsed > 500*time.Millisecond {
		t.Errorf("FetchAll took %v; fetching concurrently should take ~100ms", elapsed)
	}
}

func TestSumUntilDone(t *testing.T) {
	var got int
	within(t, time.Second, "SumUntilDone (closed channel)", func() {
		got = SumUntilDone(context.Background(), source(1, 2, 3))
	})
	if got != 6 {
		t.Errorf("SumUntilDone(closed channel with 1, 2, 3) = %d, want 6", got)
	}

	ctx, cancel := context.WithCancel(context.Background())
	in := make(chan int)
	go func() {
		in <- 10
		in <- 5
		cancel()
	}()
	within(t, time.Second, "SumUntilDone (cancelled context)", func() {
		got = SumUntilDone(ctx, in)
	})
	if got != 15 {
		t.Errorf("SumUntilDone(10, 5, then cancel) = %d, want 15", got)
	}
}

func TestFilledBuffer(t *testing.T) {
	ch := FilledBuffer(4)
	if ch == nil {
		t.Fatal("FilledBuffer(4) returned a nil channel")
	}
	if cap(ch) != 4 || len(ch) != 4 {
		t.Fatalf("FilledBuffer(4): cap = %d, len = %d, want 4, 4", cap(ch), len(ch))
	}
	if got := collect(t, ch); !slices.Equal(got, []int{1, 2, 3, 4}) {
		t.Errorf("FilledBuffer(4) contained %v, want [1 2 3 4]", got)
	}
}

func TestParallelCount(t *testing.T) {
	for _, n := range []int{1, 100, 5000} {
		var got int
		within(t, 2*time.Second, "ParallelCount", func() { got = ParallelCount(n) })
		if got != n {
			t.Errorf("ParallelCount(%d) = %d, want %d", n, got, n)
		}
	}
}
