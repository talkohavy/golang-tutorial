package rectangle

import (
	"testing"

	"go_tutorial/internal/testkit"
)

func TestRectangleType(t *testing.T) {
	r := Rectangle{Width: 2.5, Height: 4}
	var w, h float64 = r.Width, r.Height
	if w != 2.5 || h != 4 {
		t.Errorf("got %+v, want {Width:2.5 Height:4}", r)
	}
}

func TestArea(t *testing.T) {
	tests := []struct {
		r    Rectangle
		want float64
	}{
		{Rectangle{Width: 3, Height: 4}, 12},
		{Rectangle{Width: 2.5, Height: 2}, 5},
		{Rectangle{Width: 0, Height: 10}, 0},
	}
	for _, tt := range tests {
		area := testkit.Method[func() float64](t, &tt.r, "Area")
		if got := area(); got != tt.want {
			t.Errorf("%+v.Area() = %v, want %v", tt.r, got, tt.want)
		}
	}
}

func TestPerimeter(t *testing.T) {
	tests := []struct {
		r    Rectangle
		want float64
	}{
		{Rectangle{Width: 3, Height: 4}, 14},
		{Rectangle{Width: 2.5, Height: 2}, 9},
		{Rectangle{Width: 0, Height: 0}, 0},
	}
	for _, tt := range tests {
		perimeter := testkit.Method[func() float64](t, &tt.r, "Perimeter")
		if got := perimeter(); got != tt.want {
			t.Errorf("%+v.Perimeter() = %v, want %v", tt.r, got, tt.want)
		}
	}
}

func TestScale(t *testing.T) {
	r := Rectangle{Width: 2, Height: 5}
	scale := testkit.Method[func(float64)](t, &r, "Scale")
	scale(3)
	switch {
	case r == Rectangle{Width: 2, Height: 5}:
		t.Errorf("after Scale(3) the rectangle is unchanged: did Scale modify a copy?")
	case r != Rectangle{Width: 6, Height: 15}:
		t.Errorf("after Scale(3) the rectangle is %+v, want {Width:6 Height:15}", r)
	}
}
