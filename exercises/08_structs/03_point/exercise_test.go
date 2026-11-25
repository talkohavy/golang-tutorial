package point

import (
	"math"
	"testing"

	"go_tutorial/internal/testkit"
)

func TestPointType(t *testing.T) {
	p := Point{X: 1.5, Y: -2}
	var x, y float64 = p.X, p.Y
	if x != 1.5 || y != -2 {
		t.Errorf("got %+v, want {X:1.5 Y:-2}", p)
	}
}

func TestDistanceTo(t *testing.T) {
	tests := []struct {
		p, q Point
		want float64
	}{
		{Point{X: 0, Y: 0}, Point{X: 3, Y: 4}, 5},
		{Point{X: 1, Y: 1}, Point{X: -2, Y: -3}, 5},
		{Point{X: -2, Y: -3}, Point{X: 1, Y: 1}, 5},
		{Point{X: 7, Y: 7}, Point{X: 7, Y: 7}, 0},
	}
	for _, tt := range tests {
		distanceTo := testkit.Method[func(Point) float64](t, &tt.p, "DistanceTo")
		if got := distanceTo(tt.q); math.Abs(got-tt.want) > 1e-9 {
			t.Errorf("%+v.DistanceTo(%+v) = %v, want %v", tt.p, tt.q, got, tt.want)
		}
	}
}

func TestAdd(t *testing.T) {
	p, q := Point{X: 1, Y: 2}, Point{X: 3, Y: 4}
	add := testkit.Method[func(Point) Point](t, &p, "Add")
	if got, want := add(q), (Point{X: 4, Y: 6}); got != want {
		t.Errorf("{X:1 Y:2}.Add({X:3 Y:4}) = %+v, want %+v", got, want)
	}
	if p != (Point{X: 1, Y: 2}) {
		t.Errorf("Add modified the point it was called on: it's now %+v, want {X:1 Y:2}", p)
	}
}
