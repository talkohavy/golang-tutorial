//go:build solutions

// SPOILER ALERT: reference solutions. Try the exercises first!

package basics

import (
	"fmt"
	"math"
)

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

func Greeting(name string) string {
	return "Hello, " + name + "!"
}

func Swap(a, b int) (int, int) {
	return b, a
}

func CelsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

func IntDivision(a, b int) (quotient, remainder int) {
	return a / b, a % b
}

func Average(a, b, c int) float64 {
	return float64(a+b+c) / 3
}

func IsEven(n int) bool {
	return n%2 == 0
}

func SecondsToHMS(total int) (h, m, s int) {
	h = total / 3600
	m = total % 3600 / 60
	s = total % 60
	return h, m, s
}

func CircleArea(r float64) float64 {
	return math.Pi * r * r
}

func Hypotenuse(a, b float64) float64 {
	return math.Sqrt(a*a + b*b)
}

func FormatPrice(cents int) string {
	return fmt.Sprintf("$%d.%02d", cents/100, cents%100)
}

func Describe(name string, age int) string {
	return fmt.Sprintf("%s is %d years old", name, age)
}

func IsWeekend(d Weekday) bool {
	return d == Saturday || d == Sunday
}

func NextUint8(n uint8) uint8 {
	return n + 1
}
