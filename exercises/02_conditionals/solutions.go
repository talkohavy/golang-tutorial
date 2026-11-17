//go:build solutions

// SPOILER ALERT: reference solutions. Try the exercises first!

package conditionals

import (
	"strconv"
	"strings"
)

func Sign(n int) string {
	if n > 0 {
		return "positive"
	} else if n < 0 {
		return "negative"
	}
	return "zero"
}

func Max3(a, b, c int) int {
	largest := a
	if b > largest {
		largest = b
	}
	if c > largest {
		largest = c
	}
	return largest
}

func Grade(score int) string {
	switch {
	case score < 0 || score > 100:
		return "invalid"
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}

func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func DayType(day string) string {
	switch day {
	case "Saturday", "Sunday":
		return "weekend"
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		return "weekday"
	default:
		return "unknown"
	}
}

func FizzBuzz(n int) string {
	switch {
	case n%15 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(n)
	}
}

func Clamp(n, lo, hi int) int {
	if n < lo {
		return lo
	}
	if n > hi {
		return hi
	}
	return n
}

func TriangleType(a, b, c int) string {
	if a <= 0 || b <= 0 || c <= 0 || a+b <= c || a+c <= b || b+c <= a {
		return "invalid"
	}
	if a == b && b == c {
		return "equilateral"
	}
	if a == b || b == c || a == c {
		return "isosceles"
	}
	return "scalene"
}

func TicketPrice(age int, isStudent bool) int {
	switch {
	case age < 3:
		return 0
	case age <= 12:
		return 5
	case age >= 65:
		return 7
	case isStudent:
		return 8
	default:
		return 12
	}
}

func Season(month int) string {
	switch month {
	case 12, 1, 2:
		return "winter"
	case 3, 4, 5:
		return "spring"
	case 6, 7, 8:
		return "summer"
	case 9, 10, 11:
		return "autumn"
	default:
		return "invalid"
	}
}

func ParseYesNo(s string) (value bool, ok bool) {
	switch strings.ToLower(s) {
	case "y", "yes", "true", "1":
		return true, true
	case "n", "no", "false", "0":
		return false, true
	default:
		return false, false
	}
}

func BMICategory(weightKg, heightM float64) string {
	if bmi := weightKg / (heightM * heightM); bmi < 18.5 {
		return "underweight"
	} else if bmi < 25 {
		return "normal"
	} else if bmi < 30 {
		return "overweight"
	}
	return "obese"
}
