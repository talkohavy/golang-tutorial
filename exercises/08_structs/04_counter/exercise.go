//go:build !solutions

// Exercise 08.04: Counter
//
// Everything in this package is yours to write. Do the steps in order and run
// each step's test as you go:
//
//	go test ./exercises/08_structs/04_counter -run '^TestCounterType$'
package counter

// Step 1: the Counter type
//
// Declare a struct type Counter that holds a count (an int). Code in other
// packages must not be able to change the count directly, only through the
// methods below, so the field must be unexported.
//
// There's no constructor: the zero value (var c Counter) is a counter at 0,
// ready to use.
//
// Test: TestCounterType

// Step 2: Value
//
// Give Counter a method Value() int that returns the current count.
//
// Test: TestValue

// Step 3: Increment
//
// Give Counter a method Increment() that adds one to the count.
//
// Test: TestIncrement

// Step 4: Reset
//
// Give Counter a method Reset() that sets the count back to zero.
//
// Test: TestReset
