//go:build !solutions

// Exercise 08.02: Rectangle
//
// Everything in this package is yours to write. Do the steps in order and run
// each step's test as you go:
//
//	go test ./exercises/08_structs/02_rectangle -run '^TestRectangleType$'
package rectangle

// Step 1: the Rectangle type
//
// Declare a struct type Rectangle with two float64 fields: Width and Height.
//
// Test: TestRectangleType

// Step 2: Area
//
// Give Rectangle a method Area() float64. A 3 by 4 rectangle has an area of 12.
//
// Test: TestArea

// Step 3: Perimeter
//
// Give Rectangle a method Perimeter() float64. A 3 by 4 rectangle has a
// perimeter of 14.
//
// Test: TestPerimeter

// Step 4: Scale
//
// Give Rectangle a method Scale(factor float64) that multiplies both
// dimensions by factor. It modifies the rectangle it's called on and returns
// nothing: after r.Scale(2), r itself is twice as wide and twice as tall.
//
// Think about which kind of receiver makes that possible. Once the test
// passes, switch to the other kind and watch it fail.
//
// Test: TestScale
