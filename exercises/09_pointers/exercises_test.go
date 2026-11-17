package pointers

import (
	"slices"
	"testing"
)

// buildList and listValues are written independently of the exercises so
// each linked-list test only depends on the function it tests.
func buildList(nums ...int) *Node {
	var head *Node
	for i := len(nums) - 1; i >= 0; i-- {
		head = &Node{Value: nums[i], Next: head}
	}
	return head
}

func listValues(t *testing.T, head *Node) []int {
	t.Helper()
	var out []int
	for n := head; n != nil; n = n.Next {
		out = append(out, n.Value)
		if len(out) > 1000 {
			t.Fatal("list has more than 1000 nodes: it probably contains a cycle")
		}
	}
	return out
}

func TestIncrement(t *testing.T) {
	x := 5
	Increment(&x)
	Increment(&x)
	if x != 7 {
		t.Errorf("after two Increment(&x) calls x = %d, want 7", x)
	}
}

func TestSwap(t *testing.T) {
	x, y := 1, 2
	Swap(&x, &y)
	if x != 2 || y != 1 {
		t.Errorf("after Swap(&x, &y): x = %d, y = %d, want 2, 1", x, y)
	}
}

func TestNewInt(t *testing.T) {
	a := NewInt(5)
	b := NewInt(42)
	if a == nil || b == nil {
		t.Fatal("NewInt returned nil")
	}
	if *a != 5 || *b != 42 {
		t.Errorf("*NewInt(5) = %d, *NewInt(42) = %d", *a, *b)
	}
	if a == b {
		t.Error("two calls to NewInt returned the same pointer")
	}
}

func TestValueOrDefault(t *testing.T) {
	x := 7
	if got := ValueOrDefault(&x, 99); got != 7 {
		t.Errorf("ValueOrDefault(&7, 99) = %d, want 7", got)
	}
	zero := 0
	if got := ValueOrDefault(&zero, 99); got != 0 {
		t.Errorf("ValueOrDefault(&0, 99) = %d, want 0", got)
	}
	if got := ValueOrDefault(nil, 99); got != 99 {
		t.Errorf("ValueOrDefault(nil, 99) = %d, want 99", got)
	}
}

func TestBirthday(t *testing.T) {
	u := User{Name: "Ann", Age: 30}
	Birthday(&u)
	if u.Age != 31 {
		t.Errorf("after Birthday(&u) u.Age = %d, want 31", u.Age)
	}
}

func TestBirthdayAll(t *testing.T) {
	users := []User{{"Ann", 30}, {"Bob", 40}}
	BirthdayAll(users)
	if users[0].Age != 31 || users[1].Age != 41 {
		t.Errorf("after BirthdayAll users = %+v, want ages 31 and 41", users)
	}
}

func TestFindUser(t *testing.T) {
	users := []User{{"Ann", 30}, {"Bob", 40}}
	p := FindUser(users, "Bob")
	if p == nil {
		t.Fatal("FindUser(users, \"Bob\") returned nil")
	}
	p.Age = 99
	if users[1].Age != 99 {
		t.Errorf("modifying the returned pointer did not change the slice: users = %+v (did you return a pointer to a copy?)", users)
	}
	if p := FindUser(users, "Zed"); p != nil {
		t.Errorf("FindUser(users, \"Zed\") = %+v, want nil", p)
	}
}

func TestAppendTo(t *testing.T) {
	nums := []int{1}
	AppendTo(&nums, 2)
	AppendTo(&nums, 3)
	if !slices.Equal(nums, []int{1, 2, 3}) {
		t.Errorf("after AppendTo(&nums, 2) and AppendTo(&nums, 3) nums = %v, want [1 2 3]", nums)
	}
}

func TestListFromSlice(t *testing.T) {
	if got := listValues(t, ListFromSlice([]int{1, 2, 3})); !slices.Equal(got, []int{1, 2, 3}) {
		t.Errorf("ListFromSlice([1 2 3]) produced a list with values %v", got)
	}
	if got := ListFromSlice(nil); got != nil {
		t.Errorf("ListFromSlice(nil) = %+v, want nil", got)
	}
}

func TestListToSlice(t *testing.T) {
	if got := ListToSlice(buildList(4, 5, 6)); !slices.Equal(got, []int{4, 5, 6}) {
		t.Errorf("ListToSlice(4 -> 5 -> 6) = %v, want [4 5 6]", got)
	}
	if got := ListToSlice(nil); len(got) != 0 {
		t.Errorf("ListToSlice(nil) = %v, want []", got)
	}
}

func TestListLength(t *testing.T) {
	if got := ListLength(buildList(1, 2, 3, 4)); got != 4 {
		t.Errorf("ListLength(1 -> 2 -> 3 -> 4) = %d, want 4", got)
	}
	if got := ListLength(nil); got != 0 {
		t.Errorf("ListLength(nil) = %d, want 0", got)
	}
}

func TestReverseList(t *testing.T) {
	if got := listValues(t, ReverseList(buildList(1, 2, 3))); !slices.Equal(got, []int{3, 2, 1}) {
		t.Errorf("ReverseList(1 -> 2 -> 3) has values %v, want [3 2 1]", got)
	}
	if got := listValues(t, ReverseList(buildList(7))); !slices.Equal(got, []int{7}) {
		t.Errorf("ReverseList(7) has values %v, want [7]", got)
	}
	if got := ReverseList(nil); got != nil {
		t.Errorf("ReverseList(nil) = %+v, want nil", got)
	}
}
