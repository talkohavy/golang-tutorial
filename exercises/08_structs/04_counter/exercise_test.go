package counter

import (
	"testing"

	"go_tutorial/internal/testkit"
)

func TestCounterType(t *testing.T) {
	testkit.AllFieldsUnexported[Counter](t)
}

func TestValue(t *testing.T) {
	var c Counter
	value := testkit.Method[func() int](t, &c, "Value")
	if got := value(); got != 0 {
		t.Errorf("Value() of a new Counter = %d, want 0", got)
	}
}

func TestIncrement(t *testing.T) {
	var c Counter
	increment := testkit.Method[func()](t, &c, "Increment")
	value := testkit.Method[func() int](t, &c, "Value")
	increment()
	increment()
	increment()
	switch got := value(); got {
	case 3:
	case 0:
		t.Errorf("after 3 increments Value() = 0: did Increment modify a copy?")
	default:
		t.Errorf("after 3 increments Value() = %d, want 3", got)
	}
}

func TestReset(t *testing.T) {
	var c Counter
	increment := testkit.Method[func()](t, &c, "Increment")
	reset := testkit.Method[func()](t, &c, "Reset")
	value := testkit.Method[func() int](t, &c, "Value")
	increment()
	increment()
	reset()
	if got := value(); got != 0 {
		t.Errorf("after Reset() Value() = %d, want 0", got)
	}
	increment()
	if got := value(); got != 1 {
		t.Errorf("after Reset() and one Increment() Value() = %d, want 1", got)
	}
}
