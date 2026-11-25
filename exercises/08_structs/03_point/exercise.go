//go:build !solutions

// Exercise 08.03: Point
//
// Everything in this package is yours to write. Do the steps in order and run
// each step's test as you go:
//
//	go test ./exercises/08_structs/03_point -run '^TestPointType$'
package point

// Step 1: the Point type
//
// Declare a struct type Point, a position on a 2D plane, with two float64
// fields: X and Y.
//
// Test: TestPointType

// Step 2: DistanceTo
//
// Give Point a method DistanceTo(q Point) float64 that returns the straight-
// line (euclidean) distance between the point and q. From (0, 0) to (3, 4),
// the distance is 5.
//
// Hint: the math package can do the square root, or even the whole thing.
//
// Test: TestDistanceTo

// Step 3: Add
//
// Give Point a method Add(q Point) Point that returns a new Point: the sum of
// the point and q. (1, 2) + (3, 4) = (4, 6). Neither the point nor q changes.
//
// Test: TestAdd
