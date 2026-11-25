//go:build !solutions

// Exercise 08.06: IntStack
//
// Everything in this package is yours to write. Do the steps in order and run
// each step's test as you go:
//
//	go test ./exercises/08_structs/06_int_stack -run '^TestIntStackType$'
package intstack

// Step 1: the IntStack type
//
// Declare a struct type IntStack: a last-in-first-out stack of ints. Store the
// items in an unexported field, so other packages can only use the methods.
//
// There's no constructor: the zero value (var s IntStack) is an empty stack,
// ready to use.
//
// Test: TestIntStackType

// Step 2: Len
//
// Give IntStack a method Len() int that returns the number of items.
//
// Test: TestLen

// Step 3: Push
//
// Give IntStack a method Push(v int) that puts v on top of the stack.
//
// Test: TestPush

// Step 4: IsEmpty
//
// Give IntStack a method IsEmpty() bool that reports whether the stack has no
// items.
//
// Test: TestIsEmpty

// Step 5: Peek
//
// Give IntStack a method Peek() (int, bool) that returns the top item and true,
// without removing it. On an empty stack it returns 0 and false.
//
// Test: TestPeek

// Step 6: Pop
//
// Give IntStack a method Pop() (int, bool) that removes the top item and
// returns it with true. On an empty stack it returns 0 and false.
//
// Test: TestPop
