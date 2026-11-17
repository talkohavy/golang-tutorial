package basics

import (
	"math"
	"testing"
)

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) < 1e-9
}

func TestGreeting(t *testing.T) {
	tests := []struct{ name, want string }{
		{"Gopher", "Hello, Gopher!"},
		{"Alice", "Hello, Alice!"},
		{"", "Hello, !"},
	}
	for _, tt := range tests {
		if got := Greeting(tt.name); got != tt.want {
			t.Errorf("Greeting(%q) = %q, want %q", tt.name, got, tt.want)
		}
	}
}

func TestSwap(t *testing.T) {
	tests := []struct{ a, b int }{{1, 2}, {-5, 10}, {7, 7}}
	for _, tt := range tests {
		x, y := Swap(tt.a, tt.b)
		if x != tt.b || y != tt.a {
			t.Errorf("Swap(%d, %d) = %d, %d, want %d, %d", tt.a, tt.b, x, y, tt.b, tt.a)
		}
	}
}

func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct{ c, want float64 }{
		{0, 32}, {100, 212}, {-40, -40}, {37, 98.6},
	}
	for _, tt := range tests {
		if got := CelsiusToFahrenheit(tt.c); !almostEqual(got, tt.want) {
			t.Errorf("CelsiusToFahrenheit(%v) = %v, want %v", tt.c, got, tt.want)
		}
	}
}

func TestIntDivision(t *testing.T) {
	tests := []struct{ a, b, q, r int }{
		{17, 5, 3, 2}, {10, 2, 5, 0}, {-7, 2, -3, -1}, {3, 7, 0, 3},
	}
	for _, tt := range tests {
		q, r := IntDivision(tt.a, tt.b)
		if q != tt.q || r != tt.r {
			t.Errorf("IntDivision(%d, %d) = %d, %d, want %d, %d", tt.a, tt.b, q, r, tt.q, tt.r)
		}
	}
}

func TestAverage(t *testing.T) {
	tests := []struct {
		a, b, c int
		want    float64
	}{
		{1, 2, 3, 2}, {1, 2, 2, 5.0 / 3}, {0, 0, 1, 1.0 / 3}, {-3, 0, 3, 0},
	}
	for _, tt := range tests {
		if got := Average(tt.a, tt.b, tt.c); !almostEqual(got, tt.want) {
			t.Errorf("Average(%d, %d, %d) = %v, want %v", tt.a, tt.b, tt.c, got, tt.want)
		}
	}
}

func TestIsEven(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{0, true}, {4, true}, {7, false}, {-2, true}, {-3, false},
	}
	for _, tt := range tests {
		if got := IsEven(tt.n); got != tt.want {
			t.Errorf("IsEven(%d) = %t, want %t", tt.n, got, tt.want)
		}
	}
}

func TestSecondsToHMS(t *testing.T) {
	tests := []struct{ total, h, m, s int }{
		{3725, 1, 2, 5}, {59, 0, 0, 59}, {3600, 1, 0, 0}, {86399, 23, 59, 59},
	}
	for _, tt := range tests {
		h, m, s := SecondsToHMS(tt.total)
		if h != tt.h || m != tt.m || s != tt.s {
			t.Errorf("SecondsToHMS(%d) = %d, %d, %d, want %d, %d, %d", tt.total, h, m, s, tt.h, tt.m, tt.s)
		}
	}
}

func TestCircleArea(t *testing.T) {
	tests := []struct{ r, want float64 }{
		{1, math.Pi}, {2, 4 * math.Pi}, {0, 0},
	}
	for _, tt := range tests {
		if got := CircleArea(tt.r); !almostEqual(got, tt.want) {
			t.Errorf("CircleArea(%v) = %v, want %v", tt.r, got, tt.want)
		}
	}
}

func TestHypotenuse(t *testing.T) {
	tests := []struct{ a, b, want float64 }{
		{3, 4, 5}, {5, 12, 13}, {1, 1, math.Sqrt2},
	}
	for _, tt := range tests {
		if got := Hypotenuse(tt.a, tt.b); !almostEqual(got, tt.want) {
			t.Errorf("Hypotenuse(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestFormatPrice(t *testing.T) {
	tests := []struct {
		cents int
		want  string
	}{
		{1234, "$12.34"}, {5, "$0.05"}, {10000, "$100.00"}, {99, "$0.99"},
	}
	for _, tt := range tests {
		if got := FormatPrice(tt.cents); got != tt.want {
			t.Errorf("FormatPrice(%d) = %q, want %q", tt.cents, got, tt.want)
		}
	}
}

func TestDescribe(t *testing.T) {
	tests := []struct {
		name string
		age  int
		want string
	}{
		{"Alice", 30, "Alice is 30 years old"},
		{"Bob", 7, "Bob is 7 years old"},
	}
	for _, tt := range tests {
		if got := Describe(tt.name, tt.age); got != tt.want {
			t.Errorf("Describe(%q, %d) = %q, want %q", tt.name, tt.age, got, tt.want)
		}
	}
}

func TestIsWeekend(t *testing.T) {
	want := map[Weekday]bool{
		Sunday: true, Monday: false, Tuesday: false, Wednesday: false,
		Thursday: false, Friday: false, Saturday: true,
	}
	for day, w := range want {
		if got := IsWeekend(day); got != w {
			t.Errorf("IsWeekend(%d) = %t, want %t", day, got, w)
		}
	}
}

func TestNextUint8(t *testing.T) {
	tests := []struct{ n, want uint8 }{
		{41, 42}, {0, 1}, {254, 255}, {255, 0},
	}
	for _, tt := range tests {
		if got := NextUint8(tt.n); got != tt.want {
			t.Errorf("NextUint8(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}
