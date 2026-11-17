package conditionals

import "testing"

func TestSign(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{5, "positive"}, {-3, "negative"}, {0, "zero"},
	}
	for _, tt := range tests {
		if got := Sign(tt.n); got != tt.want {
			t.Errorf("Sign(%d) = %q, want %q", tt.n, got, tt.want)
		}
	}
}

func TestMax3(t *testing.T) {
	tests := []struct{ a, b, c, want int }{
		{1, 7, 3, 7}, {9, 2, 3, 9}, {1, 2, 8, 8}, {-1, -2, -3, -1}, {4, 4, 4, 4},
	}
	for _, tt := range tests {
		if got := Max3(tt.a, tt.b, tt.c); got != tt.want {
			t.Errorf("Max3(%d, %d, %d) = %d, want %d", tt.a, tt.b, tt.c, got, tt.want)
		}
	}
}

func TestGrade(t *testing.T) {
	tests := []struct {
		score int
		want  string
	}{
		{100, "A"}, {90, "A"}, {89, "B"}, {80, "B"}, {75, "C"}, {60, "D"},
		{59, "F"}, {0, "F"}, {-1, "invalid"}, {101, "invalid"},
	}
	for _, tt := range tests {
		if got := Grade(tt.score); got != tt.want {
			t.Errorf("Grade(%d) = %q, want %q", tt.score, got, tt.want)
		}
	}
}

func TestIsLeapYear(t *testing.T) {
	tests := []struct {
		year int
		want bool
	}{
		{2024, true}, {2023, false}, {1900, false}, {2000, true}, {2100, false},
	}
	for _, tt := range tests {
		if got := IsLeapYear(tt.year); got != tt.want {
			t.Errorf("IsLeapYear(%d) = %t, want %t", tt.year, got, tt.want)
		}
	}
}

func TestDayType(t *testing.T) {
	tests := []struct{ day, want string }{
		{"Saturday", "weekend"}, {"Sunday", "weekend"}, {"Monday", "weekday"},
		{"Wednesday", "weekday"}, {"Friday", "weekday"}, {"Funday", "unknown"},
	}
	for _, tt := range tests {
		if got := DayType(tt.day); got != tt.want {
			t.Errorf("DayType(%q) = %q, want %q", tt.day, got, tt.want)
		}
	}
}

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{1, "1"}, {3, "Fizz"}, {5, "Buzz"}, {15, "FizzBuzz"}, {9, "Fizz"},
		{10, "Buzz"}, {30, "FizzBuzz"}, {7, "7"},
	}
	for _, tt := range tests {
		if got := FizzBuzz(tt.n); got != tt.want {
			t.Errorf("FizzBuzz(%d) = %q, want %q", tt.n, got, tt.want)
		}
	}
}

func TestClamp(t *testing.T) {
	tests := []struct{ n, lo, hi, want int }{
		{15, 0, 10, 10}, {-3, 0, 10, 0}, {5, 0, 10, 5}, {10, 0, 10, 10}, {-50, -20, -10, -20},
	}
	for _, tt := range tests {
		if got := Clamp(tt.n, tt.lo, tt.hi); got != tt.want {
			t.Errorf("Clamp(%d, %d, %d) = %d, want %d", tt.n, tt.lo, tt.hi, got, tt.want)
		}
	}
}

func TestTriangleType(t *testing.T) {
	tests := []struct {
		a, b, c int
		want    string
	}{
		{3, 3, 3, "equilateral"}, {3, 3, 5, "isosceles"}, {5, 3, 5, "isosceles"},
		{3, 4, 5, "scalene"}, {1, 2, 3, "invalid"}, {0, 1, 1, "invalid"},
		{-1, 2, 2, "invalid"}, {10, 1, 1, "invalid"},
	}
	for _, tt := range tests {
		if got := TriangleType(tt.a, tt.b, tt.c); got != tt.want {
			t.Errorf("TriangleType(%d, %d, %d) = %q, want %q", tt.a, tt.b, tt.c, got, tt.want)
		}
	}
}

func TestTicketPrice(t *testing.T) {
	tests := []struct {
		age       int
		isStudent bool
		want      int
	}{
		{1, false, 0}, {3, false, 5}, {12, true, 5}, {13, false, 12}, {20, true, 8},
		{64, true, 8}, {65, false, 7}, {80, true, 7},
	}
	for _, tt := range tests {
		if got := TicketPrice(tt.age, tt.isStudent); got != tt.want {
			t.Errorf("TicketPrice(%d, %t) = %d, want %d", tt.age, tt.isStudent, got, tt.want)
		}
	}
}

func TestSeason(t *testing.T) {
	tests := []struct {
		month int
		want  string
	}{
		{1, "winter"}, {2, "winter"}, {12, "winter"}, {3, "spring"}, {5, "spring"},
		{6, "summer"}, {8, "summer"}, {9, "autumn"}, {11, "autumn"}, {0, "invalid"}, {13, "invalid"},
	}
	for _, tt := range tests {
		if got := Season(tt.month); got != tt.want {
			t.Errorf("Season(%d) = %q, want %q", tt.month, got, tt.want)
		}
	}
}

func TestParseYesNo(t *testing.T) {
	tests := []struct {
		in        string
		value, ok bool
	}{
		{"y", true, true}, {"YES", true, true}, {"True", true, true}, {"1", true, true},
		{"n", false, true}, {"No", false, true}, {"FALSE", false, true}, {"0", false, true},
		{"maybe", false, false}, {"", false, false},
	}
	for _, tt := range tests {
		value, ok := ParseYesNo(tt.in)
		if value != tt.value || ok != tt.ok {
			t.Errorf("ParseYesNo(%q) = %t, %t, want %t, %t", tt.in, value, ok, tt.value, tt.ok)
		}
	}
}

func TestBMICategory(t *testing.T) {
	tests := []struct {
		weight, height float64
		want           string
	}{
		{50, 1.80, "underweight"}, {70, 1.75, "normal"}, {85, 1.75, "overweight"}, {110, 1.75, "obese"},
	}
	for _, tt := range tests {
		if got := BMICategory(tt.weight, tt.height); got != tt.want {
			t.Errorf("BMICategory(%v, %v) = %q, want %q", tt.weight, tt.height, got, tt.want)
		}
	}
}
