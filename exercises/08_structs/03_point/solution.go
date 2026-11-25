//go:build solutions

// SPOILER ALERT: reference solution. Try the exercise first!

package point

import "math"

type Point struct {
	X, Y float64
}

func (p Point) DistanceTo(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Add uses a value receiver, so it can't modify p even by accident: it works
// on a copy and returns a brand new Point.
func (p Point) Add(q Point) Point {
	return Point{X: p.X + q.X, Y: p.Y + q.Y}
}
