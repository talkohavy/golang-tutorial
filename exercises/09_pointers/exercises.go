//go:build !solutions

// Package pointers covers pointers and the difference between values and
// references.
//
// Quick syntax reminder:
//
//	x := 5
//	p := &x        // p is a *int that points to x
//	*p = 10        // dereference: x is now 10
//	var q *int     // nil pointer: dereferencing it panics
//	n := new(int)  // allocates a zero int and returns its address
//	u := &User{}   // pointer to a struct literal
//	u.Name = "Al"  // Go auto-dereferences struct pointers (no (*u).Name needed)
//
// Everything in Go is passed BY VALUE (copied). To let a function modify the
// caller's variable, pass a pointer to it. Slices and maps feel like
// references because they contain a pointer internally, but the slice header
// itself (pointer, length, capacity) is still copied.
package pointers

// User is used by several exercises below.
type User struct {
	Name string
	Age  int
}

// Node is one element of a singly linked list.
type Node struct {
	Value int
	Next  *Node
}

// Exercise 1: Increment
//
// Add one to the int that p points to.
//
//	x := 5; Increment(&x) → x == 6
func Increment(p *int) {
	// TODO: implement
}

// Exercise 2: Swap
//
// Swap the values that a and b point to.
//
//	x, y := 1, 2; Swap(&x, &y) → x == 2, y == 1
func Swap(a, b *int) {
	// TODO: implement
}

// Exercise 3: NewInt
//
// Return a pointer to a new int holding v. Each call must return a different
// pointer.
//
// Concepts: returning the address of a local variable is safe in Go (the
// compiler moves it to the heap).
func NewInt(v int) *int {
	// TODO: implement
	return new(int)
}

// Exercise 4: ValueOrDefault
//
// Return the value p points to, or def if p is nil.
func ValueOrDefault(p *int, def int) int {
	// TODO: implement
	return 0
}

// Exercise 5: Birthday
//
// Increase the user's age by one.
func Birthday(u *User) {
	// TODO: implement
}

// Exercise 6: BirthdayAll
//
// Increase every user's age by one, modifying the slice's elements.
//
// Gotcha: in `for _, u := range users`, u is a COPY of the element.
// Modifying u changes nothing. Use the index instead.
func BirthdayAll(users []User) {
	// TODO: implement
}

// Exercise 7: FindUser
//
// Return a pointer to the user with the given name INSIDE the slice (so the
// caller can modify it through the pointer), or nil if not found.
//
//	p := FindUser(users, "Bob"); p.Age = 99 → users[i].Age == 99
func FindUser(users []User, name string) *User {
	// TODO: implement
	return nil
}

// Exercise 8: AppendTo
//
// Append v to the slice that s points to, so the caller sees the new element.
//
//	nums := []int{1}; AppendTo(&nums, 2) → nums == [1 2]
//
// Concepts: append may return a new slice header, so to change the caller's
// slice you need a pointer to it.
func AppendTo(s *[]int, v int) {
	// TODO: implement
}

// Exercise 9: ListFromSlice
//
// Build a linked list with the same values, in order, and return its head
// (nil for an empty slice).
//
//	ListFromSlice([]int{1, 2, 3}) → 1 -> 2 -> 3 -> nil
func ListFromSlice(nums []int) *Node {
	// TODO: implement
	return nil
}

// Exercise 10: ListToSlice
//
// Walk the list from head and collect its values.
//
//	for n := head; n != nil; n = n.Next { ... }
func ListToSlice(head *Node) []int {
	// TODO: implement
	return nil
}

// Exercise 11: ListLength
//
// Return the number of nodes in the list.
func ListLength(head *Node) int {
	// TODO: implement
	return 0
}

// Exercise 12: ReverseList
//
// Reverse the list in place (by re-pointing Next fields) and return the new
// head.
//
//	1 -> 2 -> 3  becomes  3 -> 2 -> 1
func ReverseList(head *Node) *Node {
	// TODO: implement
	return nil
}
