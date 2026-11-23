//go:build !solutions

package basics

// Exercise 1: Greeting
//
// Return a greeting in the form "Hello, <name>!".
//
//	Greeting("Gopher") → "Hello, Gopher!"
//
// Concepts: string concatenation with +.
func Greeting(name string) string {
	// TODO: implement
	return ""
}

// Exercise 2: Swap
//
// Return the two values in reverse order.
//
//	Swap(1, 2) → 2, 1
//
// Concepts: multiple return values.
func Swap(a, b int) (int, int) {
	// TODO: implement
	return 0, 0
}

// Exercise 3: CelsiusToFahrenheit
//
// Convert a temperature from Celsius to Fahrenheit: F = C × 9/5 + 32.
//
//	CelsiusToFahrenheit(100) → 212
//
// Concepts: float64 arithmetic.
func CelsiusToFahrenheit(c float64) float64 {
	// TODO: implement
	return 0
}

// Exercise 4: IntDivision
//
// Return the integer quotient and the remainder of a / b.
//
//	IntDivision(17, 5) → 3, 2
//	IntDivision(-7, 2) → -3, -1   (Go's integer division truncates toward zero)
//
// Concepts: the / and % operators on ints, named return values.
func IntDivision(a, b int) (quotient, remainder int) {
	// TODO: implement
	return 0, 0
}

// Exercise 5: Average
//
// Return the average of three ints as a float64.
//
//	Average(1, 2, 2) → 1.6666666666666667
//
// Watch out: (a+b+c)/3 is *integer* division. Convert before dividing.
//
// Concepts: type conversion with float64(...).
func Average(a, b, c int) float64 {
	// TODO: implement
	return 0
}

// Exercise 6: IsEven
//
// Report whether n is even.
//
//	IsEven(4) → true
//	IsEven(7) → false
//
// Concepts: the % operator and boolean expressions.
func IsEven(n int) bool {
	// TODO: implement
	return false
}

// Exercise 7: SecondsToHMS
//
// Split a number of seconds into hours, minutes and seconds.
//
//	SecondsToHMS(3725) → 1, 2, 5
//
// Concepts: integer division and remainder.
func SecondsToHMS(total int) (h, m, s int) {
	// TODO: implement
	return 0, 0, 0
}

// Exercise 8: CircleArea
//
// Return the area of a circle with radius r (π × r²).
//
//	CircleArea(2) → 12.566370614359172
//
// Concepts: the math package (math.Pi). You'll need to add an import.
func CircleArea(r float64) float64 {
	// TODO: implement
	return 0
}

// Exercise 9: Hypotenuse
//
// Return the length of the hypotenuse of a right triangle with sides a and b.
//
//	Hypotenuse(3, 4) → 5
//
// Concepts: math.Sqrt.
func Hypotenuse(a, b float64) float64 {
	// TODO: implement
	return 0
}

// Exercise 10: FormatPrice
//
// Format an amount of cents as a dollar price with exactly two decimals.
//
//	FormatPrice(1234) → "$12.34"
//	FormatPrice(5)    → "$0.05"
//
// Concepts: fmt.Sprintf with %d and zero padding (%02d).
func FormatPrice(cents int) string {
	// TODO: implement
	return ""
}

// Exercise 11: Describe
//
// Build a sentence describing a person.
//
//	Describe("Alice", 30) → "Alice is 30 years old"
//
// Concepts: fmt.Sprintf with %s and %d.
func Describe(name string, age int) string {
	// TODO: implement
	return ""
}

// Weekday is a custom type whose underlying type is int. The constants below
// use iota, which starts at 0 and increments by one on every line of the
// const block. This is Go's idiomatic way to declare enums.
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// Exercise 12: IsWeekend
//
// Report whether the day is Saturday or Sunday.
//
//	IsWeekend(Saturday) → true
//	IsWeekend(Monday)   → false
//
// Concepts: custom types and constants declared with iota (see Weekday above).
func IsWeekend(d Weekday) bool {
	// TODO: implement
	return false
}

// Exercise 13: NextUint8
//
// Return n + 1 as a uint8.
//
//	NextUint8(41)  → 42
//	NextUint8(255) → 0   (uint8 holds 0..255, so it wraps around)
//
// Concepts: sized integer types and overflow.
func NextUint8(n uint8) uint8 {
	// TODO: implement
	return 0
}
