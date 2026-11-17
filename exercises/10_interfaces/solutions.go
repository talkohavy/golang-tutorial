//go:build solutions

// SPOILER ALERT: reference solutions. Try the exercises first!

package interfaces

import (
	"fmt"
	"math"
	"strings"
)

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

var (
	_ Shape  = Circle{}
	_ Shape  = Rect{}
	_ Scaler = (*Circle)(nil)
)

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rect) Area() float64 {
	return r.Width * r.Height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func TotalArea(shapes []Shape) float64 {
	total := 0.0
	for _, s := range shapes {
		total += s.Area()
	}
	return total
}

func Largest(shapes []Shape) Shape {
	var largest Shape
	for _, s := range shapes {
		if largest == nil || s.Area() > largest.Area() {
			largest = s
		}
	}
	return largest
}

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

type Speaker interface {
	Speak() string
}

type Dog struct{ Name string }
type Cat struct{ Name string }
type Robot struct{ Model string }

func (d Dog) Speak() string {
	return "Woof! I'm " + d.Name
}

func (c Cat) Speak() string {
	return "Meow! I'm " + c.Name
}

func (r Robot) Speak() string {
	return "Beep boop. Model " + r.Model
}

func Chorus(speakers []Speaker) string {
	lines := make([]string, len(speakers))
	for i, s := range speakers {
		lines[i] = s.Speak()
	}
	return strings.Join(lines, " | ")
}

func Describe(v any) string {
	switch x := v.(type) {
	case nil:
		return "nil"
	case int:
		return fmt.Sprintf("int: %d", x)
	case string:
		return fmt.Sprintf("string: %q", x)
	case bool:
		return fmt.Sprintf("bool: %t", x)
	case []int:
		return fmt.Sprintf("[]int of length %d", len(x))
	case error:
		return "error: " + x.Error()
	default:
		return "unknown"
	}
}

func SumNumeric(values []any) float64 {
	sum := 0.0
	for _, v := range values {
		switch n := v.(type) {
		case int:
			sum += float64(n)
		case int64:
			sum += float64(n)
		case float64:
			sum += n
		}
	}
	return sum
}

type ByLength []string

func (b ByLength) Len() int           { return len(b) }
func (b ByLength) Less(i, j int) bool { return len(b[i]) < len(b[j]) }
func (b ByLength) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

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

type Queue struct {
	items []string
}

func (q *Queue) Write(s string) {
	q.items = append(q.items, s)
}

func (q *Queue) Read() (string, bool) {
	if len(q.items) == 0 {
		return "", false
	}
	front := q.items[0]
	q.items = q.items[1:]
	return front, true
}

type CountingWriter struct {
	Count int
}

func (w *CountingWriter) Write(p []byte) (int, error) {
	w.Count += len(p)
	return len(p), nil
}

type Scaler interface {
	Scale(factor float64)
}

func (c *Circle) Scale(factor float64) {
	c.Radius *= factor
}

func (r *Rect) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

func ScaleAll(shapes []Shape, factor float64) int {
	scaled := 0
	for _, s := range shapes {
		if sc, ok := s.(Scaler); ok {
			sc.Scale(factor)
			scaled++
		}
	}
	return scaled
}
