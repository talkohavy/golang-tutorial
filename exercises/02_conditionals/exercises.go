//go:build !solutions

package conditionals

// Exercise 1: Sign
//
// Return "positive", "negative" or "zero".
//
//	Sign(5) → "positive"
func Sign(n int) string {
	// TODO: implement
	return ""
}

// Exercise 2: Max3
//
// Return the largest of three ints. Use if statements (not the built-in max).
//
//	Max3(1, 7, 3) → 7
func Max3(a, b, c int) int {
	// TODO: implement
	return 0
}

// Exercise 3: Grade
//
// Convert a score into a letter grade:
//
//	90–100 → "A", 80–89 → "B", 70–79 → "C", 60–69 → "D", 0–59 → "F"
//	anything outside 0–100 → "invalid"
//
// Concepts: tagless switch.
func Grade(score int) string {
	// TODO: implement
	return ""
}

// Exercise 4: IsLeapYear
//
// A year is a leap year if it is divisible by 4, except years divisible by
// 100, unless they are also divisible by 400.
//
//	IsLeapYear(2024) → true
//	IsLeapYear(1900) → false
//	IsLeapYear(2000) → true
//
// Concepts: combining conditions with && and ||.
func IsLeapYear(year int) bool {
	// TODO: implement
	return false
}

// Exercise 5: DayType
//
// Return "weekend" for "Saturday"/"Sunday", "weekday" for "Monday" through
// "Friday", and "unknown" for anything else.
//
// Concepts: switch with several values per case.
func DayType(day string) string {
	// TODO: implement
	return ""
}

// Exercise 6: FizzBuzz
//
// Return "FizzBuzz" if n is divisible by 3 and 5, "Fizz" if divisible by 3,
// "Buzz" if divisible by 5, otherwise the number itself as a string.
//
//	FizzBuzz(9) → "Fizz"
//	FizzBuzz(7) → "7"
//
// Concepts: order of conditions matters; strconv.Itoa converts int → string.
func FizzBuzz(n int) string {
	// TODO: implement
	return ""
}

// Exercise 7: Clamp
//
// Restrict n to the range [lo, hi].
//
//	Clamp(15, 0, 10) → 10
//	Clamp(-3, 0, 10) → 0
//	Clamp(5, 0, 10)  → 5
func Clamp(n, lo, hi int) int {
	// TODO: implement
	return 0
}

// Exercise 8: TriangleType
//
// Classify a triangle by its side lengths:
//   - "invalid" if any side is <= 0, or the sum of any two sides is not
//     greater than the third
//   - "equilateral" if all three sides are equal
//   - "isosceles" if exactly two sides are equal
//   - "scalene" otherwise
func TriangleType(a, b, c int) string {
	// TODO: implement
	return ""
}

// Exercise 9: TicketPrice
//
// Return the ticket price for a visitor:
//   - under 3: 0
//   - 3–12: 5
//   - 65 and over: 7
//   - everyone else: 12, but students pay 8
func TicketPrice(age int, isStudent bool) int {
	// TODO: implement
	return 0
}

// Exercise 10: Season
//
// Return the (northern hemisphere) season for a month number:
// 12, 1, 2 → "winter"; 3–5 → "spring"; 6–8 → "summer"; 9–11 → "autumn";
// anything else → "invalid".
func Season(month int) string {
	// TODO: implement
	return ""
}

// Exercise 11: ParseYesNo
//
// Interpret a user's answer, case-insensitively:
//   - "y", "yes", "true", "1"  → true, true
//   - "n", "no", "false", "0"  → false, true
//   - anything else            → false, false (ok is false: not understood)
//
// Concepts: switch on strings.ToLower(s), the "value, ok" return pattern.
func ParseYesNo(s string) (value bool, ok bool) {
	// TODO: implement
	return false, false
}

// Exercise 12: BMICategory
//
// Compute BMI = weight / (height × height) (kg and meters) and return:
// below 18.5 → "underweight"; below 25 → "normal"; below 30 → "overweight";
// otherwise "obese".
//
// Concepts: if with a short statement, e.g. `if bmi := ...; bmi < 18.5 {`.
func BMICategory(weightKg, heightM float64) string {
	// TODO: implement
	return ""
}
