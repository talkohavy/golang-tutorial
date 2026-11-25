package inventory

import (
	"testing"

	"go_tutorial/internal/testkit"
)

func TestInventoryType(t *testing.T) {
	testkit.AllFieldsUnexported[Inventory](t)
}

func TestNewInventory(t *testing.T) {
	testkit.MapsInitialized(t, NewInventory())
}

func TestCount(t *testing.T) {
	inv := testkit.Ptr(t, NewInventory())
	count := testkit.Method[func(string) int](t, inv, "Count")
	if got := count("apple"); got != 0 {
		t.Errorf("Count(\"apple\") on a new inventory = %d, want 0", got)
	}
}

func TestAdd(t *testing.T) {
	inv := testkit.Ptr(t, NewInventory())
	add := testkit.Method[func(string, int)](t, inv, "Add")
	count := testkit.Method[func(string) int](t, inv, "Count")
	add("apple", 5)
	add("apple", 3)
	add("pear", 2)
	if got := count("apple"); got != 8 {
		t.Errorf("after adding 5 then 3 apples Count(\"apple\") = %d, want 8", got)
	}
	if got := count("pear"); got != 2 {
		t.Errorf("after adding 2 pears Count(\"pear\") = %d, want 2", got)
	}
}

func TestItems(t *testing.T) {
	inv := testkit.Ptr(t, NewInventory())
	items := testkit.Method[func() int](t, inv, "Items")
	add := testkit.Method[func(string, int)](t, inv, "Add")
	if got := items(); got != 0 {
		t.Errorf("Items() of a new inventory = %d, want 0", got)
	}
	add("apple", 5)
	add("pear", 2)
	add("apple", 1)
	if got := items(); got != 2 {
		t.Errorf("after adding apples and pears Items() = %d, want 2", got)
	}
}

func TestRemove(t *testing.T) {
	inv := testkit.Ptr(t, NewInventory())
	remove := testkit.Method[func(string, int) bool](t, inv, "Remove")
	add := testkit.Method[func(string, int)](t, inv, "Add")
	count := testkit.Method[func(string) int](t, inv, "Count")
	items := testkit.Method[func() int](t, inv, "Items")
	add("apple", 8)
	add("pear", 2)

	if remove("apple", 100) {
		t.Error("Remove(\"apple\", 100) with 8 in stock returned true, want false")
	}
	if got := count("apple"); got != 8 {
		t.Errorf("a failed Remove changed Count(\"apple\") to %d, want it to stay 8", got)
	}
	if remove("banana", 1) {
		t.Error("Remove(\"banana\", 1) for an unknown item returned true, want false")
	}

	if !remove("apple", 3) {
		t.Error("Remove(\"apple\", 3) with 8 in stock returned false, want true")
	}
	if got := count("apple"); got != 5 {
		t.Errorf("after removing 3 of 8 apples Count(\"apple\") = %d, want 5", got)
	}

	if !remove("apple", 5) {
		t.Error("Remove(\"apple\", 5) with 5 in stock returned false, want true")
	}
	if got := count("apple"); got != 0 {
		t.Errorf("after removing every apple Count(\"apple\") = %d, want 0", got)
	}
	if got := items(); got != 1 {
		t.Errorf("after removing every apple Items() = %d, want 1 (only pears are left)", got)
	}
}
