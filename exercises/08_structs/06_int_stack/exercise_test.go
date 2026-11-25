package intstack

import (
	"testing"

	"go_tutorial/internal/testkit"
)

func TestIntStackType(t *testing.T) {
	testkit.AllFieldsUnexported[IntStack](t)
}

func TestLen(t *testing.T) {
	var s IntStack
	length := testkit.Method[func() int](t, &s, "Len")
	if got := length(); got != 0 {
		t.Errorf("Len() of a new stack = %d, want 0", got)
	}
}

func TestPush(t *testing.T) {
	var s IntStack
	push := testkit.Method[func(int)](t, &s, "Push")
	length := testkit.Method[func() int](t, &s, "Len")
	push(10)
	push(20)
	push(30)
	switch got := length(); got {
	case 3:
	case 0:
		t.Errorf("after 3 pushes Len() = 0: did Push modify a copy?")
	default:
		t.Errorf("after 3 pushes Len() = %d, want 3", got)
	}
}

func TestIsEmpty(t *testing.T) {
	var s IntStack
	isEmpty := testkit.Method[func() bool](t, &s, "IsEmpty")
	push := testkit.Method[func(int)](t, &s, "Push")
	if !isEmpty() {
		t.Error("IsEmpty() of a new stack = false, want true")
	}
	push(1)
	if isEmpty() {
		t.Error("IsEmpty() after a push = true, want false")
	}
}

func TestPeek(t *testing.T) {
	var s IntStack
	peek := testkit.Method[func() (int, bool)](t, &s, "Peek")
	push := testkit.Method[func(int)](t, &s, "Push")
	length := testkit.Method[func() int](t, &s, "Len")

	if v, ok := peek(); v != 0 || ok {
		t.Errorf("Peek() on an empty stack = %d, %t, want 0, false", v, ok)
	}
	push(1)
	push(2)
	if v, ok := peek(); v != 2 || !ok {
		t.Errorf("Peek() after pushing 1, 2 = %d, %t, want 2, true", v, ok)
	}
	if got := length(); got != 2 {
		t.Errorf("Peek() changed Len() to %d, want it to stay 2", got)
	}
}

func TestPop(t *testing.T) {
	var s IntStack
	pop := testkit.Method[func() (int, bool)](t, &s, "Pop")
	push := testkit.Method[func(int)](t, &s, "Push")
	length := testkit.Method[func() int](t, &s, "Len")

	if v, ok := pop(); v != 0 || ok {
		t.Errorf("Pop() on an empty stack = %d, %t, want 0, false", v, ok)
	}
	push(1)
	push(2)
	push(3)
	for _, want := range []int{3, 2, 1} {
		if v, ok := pop(); v != want || !ok {
			t.Errorf("Pop() = %d, %t, want %d, true", v, ok, want)
		}
	}
	if got := length(); got != 0 {
		t.Errorf("after popping everything Len() = %d, want 0", got)
	}
	if v, ok := pop(); v != 0 || ok {
		t.Errorf("Pop() on an emptied stack = %d, %t, want 0, false", v, ok)
	}
}
