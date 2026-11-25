//go:build solutions

// SPOILER ALERT: reference solution. Try the exercise first!

package intstack

// The zero value of a slice is nil, and len and append both work on a nil
// slice, so a zero IntStack is already an empty, usable stack.
type IntStack struct {
	items []int
}

func (s *IntStack) Len() int {
	return len(s.items)
}

func (s *IntStack) Push(v int) {
	s.items = append(s.items, v)
}

func (s *IntStack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *IntStack) Peek() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	return s.items[len(s.items)-1], true
}

func (s *IntStack) Pop() (int, bool) {
	top, ok := s.Peek()
	if ok {
		s.items = s.items[:len(s.items)-1]
	}
	return top, ok
}
