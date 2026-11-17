//go:build !solutions

// Package interfaces covers interfaces, type assertions and type switches.
//
// An interface is a set of method signatures. A type satisfies an interface
// IMPLICITLY, just by having those methods: there's no "implements" keyword.
//
// Quick syntax reminder:
//
//	type Shape interface { Area() float64 }
//	var s Shape = Circle{Radius: 1}   // works because Circle has Area()
//	var _ Shape = Circle{}            // compile-time check that Circle is a Shape
//
//	c, ok := s.(Circle)               // type assertion (ok is false if s isn't a Circle)
//	switch v := x.(type) {            // type switch
//	case int:    // v is an int here
//	case string: // v is a string here
//	}
//
//	any                               // alias for interface{}: any value at all
//
// Important standard interfaces: fmt.Stringer (String() string),
// error (Error() string), io.Writer (Write([]byte) (int, error)),
// sort.Interface (Len, Less, Swap).
package interfaces

// Shape is implemented by Circle and Rect below.
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

type Rect struct {
	Width, Height float64
}

// Exercise 1: Circle implements Shape
//
// Area = π × r², Perimeter = 2 × π × r.
func (c Circle) Area() float64 {
	// TODO: implement
	return 0
}

func (c Circle) Perimeter() float64 {
	// TODO: implement
	return 0
}

// Exercise 2: Rect implements Shape
func (r Rect) Area() float64 {
	// TODO: implement
	return 0
}

func (r Rect) Perimeter() float64 {
	// TODO: implement
	return 0
}

// Exercise 3: TotalArea
//
// Return the sum of the areas of all shapes, whatever their concrete type.
//
// Concepts: polymorphism through interfaces.
func TotalArea(shapes []Shape) float64 {
	// TODO: implement
	return 0
}

// Exercise 4: Largest
//
// Return the shape with the biggest area, or nil if the slice is empty.
//
// Concepts: nil is the zero value of an interface.
func Largest(shapes []Shape) Shape {
	// TODO: implement
	return nil
}

// Point implements fmt.Stringer.
type Point struct {
	X, Y int
}

// Exercise 5: String
//
// Return the point formatted as "(X, Y)". Once Point has a String() method,
// fmt.Println(p) and fmt.Sprintf("%v", p) use it automatically.
//
//	Point{1, -2}.String() → "(1, -2)"
func (p Point) String() string {
	// TODO: implement
	return ""
}

type Speaker interface {
	Speak() string
}

type Dog struct{ Name string }
type Cat struct{ Name string }
type Robot struct{ Model string }

// Exercise 6: Speak
//
//	Dog{"Rex"}.Speak()      → "Woof! I'm Rex"
//	Cat{"Tom"}.Speak()      → "Meow! I'm Tom"
//	Robot{"T-800"}.Speak()  → "Beep boop. Model T-800"
func (d Dog) Speak() string {
	// TODO: implement
	return ""
}

func (c Cat) Speak() string {
	// TODO: implement
	return ""
}

func (r Robot) Speak() string {
	// TODO: implement
	return ""
}

// Exercise 7: Chorus
//
// Join what every speaker says with " | ".
//
//	Chorus([]Speaker{Dog{"Rex"}, Robot{"R2"}}) → "Woof! I'm Rex | Beep boop. Model R2"
func Chorus(speakers []Speaker) string {
	// TODO: implement
	return ""
}

// Exercise 8: Describe
//
// Describe any value using a type switch:
//
//	nil            → "nil"
//	int            → "int: 42"
//	string         → `string: "hi"`     (quoted, use %q)
//	bool           → "bool: true"
//	[]int          → "[]int of length 3"
//	error          → "error: <the error message>"
//	anything else  → "unknown"
func Describe(v any) string {
	// TODO: implement
	return ""
}

// Exercise 9: SumNumeric
//
// Sum every int, int64 and float64 in values, ignoring everything else.
//
//	SumNumeric([]any{1, 2.5, "3", int64(4), nil}) → 7.5
func SumNumeric(values []any) float64 {
	// TODO: implement
	return 0
}

// ByLength sorts strings by length (shortest first) when used with sort.Sort.
type ByLength []string

// Exercise 10: sort.Interface
//
// Implement Len, Less and Swap so that sort.Sort(ByLength(words)) sorts the
// words by length.
func (b ByLength) Len() int {
	// TODO: implement
	return 0
}

func (b ByLength) Less(i, j int) bool {
	// TODO: implement
	return false
}

func (b ByLength) Swap(i, j int) {
	// TODO: implement
}

// Interfaces can be composed from other interfaces.
type Reader interface {
	Read() (string, bool)
}

type Writer interface {
	Write(s string)
}

type ReadWriter interface {
	Reader
	Writer
}

// Queue should satisfy ReadWriter.
type Queue struct {
	items []string
}

// Exercise 11: Queue implements ReadWriter
//
// Write appends s to the back. Read removes and returns the front item, or
// "", false if the queue is empty (first in, first out).
func (q *Queue) Write(s string) {
	// TODO: implement
}

func (q *Queue) Read() (string, bool) {
	// TODO: implement
	return "", false
}

// CountingWriter should satisfy io.Writer and count the bytes written to it.
type CountingWriter struct {
	Count int
}

// Exercise 12: io.Writer
//
// Add len(p) to Count and report that all bytes were written. Because
// *CountingWriter will then be an io.Writer, you can pass it to fmt.Fprintf,
// io.Copy, and everything else in the standard library that writes.
func (w *CountingWriter) Write(p []byte) (int, error) {
	// TODO: implement
	return 0, nil
}

// Scaler is implemented by shapes that can be resized.
type Scaler interface {
	Scale(factor float64)
}

// Exercise 13: Scale (pointer receivers)
//
// Multiply the circle's radius / the rect's sides by factor.
//
// Concepts: these methods need pointer receivers, so only *Circle and *Rect
// (not Circle and Rect) satisfy Scaler.
func (c *Circle) Scale(factor float64) {
	// TODO: implement
}

func (r *Rect) Scale(factor float64) {
	// TODO: implement
}

// Exercise 14: ScaleAll
//
// Scale every shape that also implements Scaler; leave the others alone.
// Return how many shapes were scaled.
//
// Concepts: asserting from one interface to another: s.(Scaler).
func ScaleAll(shapes []Shape, factor float64) int {
	// TODO: implement
	return 0
}
