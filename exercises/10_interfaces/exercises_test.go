package interfaces

import (
	"errors"
	"fmt"
	"io"
	"math"
	"slices"
	"sort"
	"strings"
	"testing"
)

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) < 1e-9
}

func TestCircle(t *testing.T) {
	var s Shape = Circle{Radius: 2}
	if got := s.Area(); !almostEqual(got, 4*math.Pi) {
		t.Errorf("Circle{2}.Area() = %v, want %v", got, 4*math.Pi)
	}
	if got := s.Perimeter(); !almostEqual(got, 4*math.Pi) {
		t.Errorf("Circle{2}.Perimeter() = %v, want %v", got, 4*math.Pi)
	}
	if got := (Circle{Radius: 1}).Perimeter(); !almostEqual(got, 2*math.Pi) {
		t.Errorf("Circle{1}.Perimeter() = %v, want %v", got, 2*math.Pi)
	}
}

func TestRect(t *testing.T) {
	var s Shape = Rect{Width: 3, Height: 4}
	if got := s.Area(); got != 12 {
		t.Errorf("Rect{3, 4}.Area() = %v, want 12", got)
	}
	if got := s.Perimeter(); got != 14 {
		t.Errorf("Rect{3, 4}.Perimeter() = %v, want 14", got)
	}
}

// fixedShape lets the TotalArea and Largest tests work even before Circle
// and Rect are implemented.
type fixedShape struct{ area float64 }

func (f fixedShape) Area() float64      { return f.area }
func (f fixedShape) Perimeter() float64 { return 0 }

func TestTotalArea(t *testing.T) {
	shapes := []Shape{fixedShape{1.5}, fixedShape{2}, fixedShape{10}}
	if got := TotalArea(shapes); !almostEqual(got, 13.5) {
		t.Errorf("TotalArea(areas 1.5, 2, 10) = %v, want 13.5", got)
	}
	if got := TotalArea(nil); got != 0 {
		t.Errorf("TotalArea(nil) = %v, want 0", got)
	}
}

func TestLargest(t *testing.T) {
	shapes := []Shape{fixedShape{1}, fixedShape{7}, fixedShape{3}}
	if got := Largest(shapes); got != (fixedShape{7}) {
		t.Errorf("Largest(areas 1, 7, 3) = %v, want the shape with area 7", got)
	}
	if got := Largest(nil); got != nil {
		t.Errorf("Largest(nil) = %v, want nil", got)
	}
}

func TestString(t *testing.T) {
	p := Point{1, -2}
	if got := p.String(); got != "(1, -2)" {
		t.Errorf("Point{1, -2}.String() = %q, want %q", got, "(1, -2)")
	}
	if got := fmt.Sprintf("point: %v", Point{3, 4}); got != "point: (3, 4)" {
		t.Errorf("fmt.Sprintf(\"point: %%v\", Point{3, 4}) = %q, want %q", got, "point: (3, 4)")
	}
	var _ fmt.Stringer = p
}

func TestSpeak(t *testing.T) {
	tests := []struct {
		s    Speaker
		want string
	}{
		{Dog{"Rex"}, "Woof! I'm Rex"}, {Cat{"Tom"}, "Meow! I'm Tom"}, {Robot{"T-800"}, "Beep boop. Model T-800"},
	}
	for _, tt := range tests {
		if got := tt.s.Speak(); got != tt.want {
			t.Errorf("%#v.Speak() = %q, want %q", tt.s, got, tt.want)
		}
	}
}

type parrot struct{ phrase string }

func (p parrot) Speak() string { return p.phrase }

func TestChorus(t *testing.T) {
	speakers := []Speaker{parrot{"hello"}, parrot{"pretty bird"}, parrot{"cracker"}}
	if got, want := Chorus(speakers), "hello | pretty bird | cracker"; got != want {
		t.Errorf("Chorus(...) = %q, want %q", got, want)
	}
	if got := Chorus([]Speaker{parrot{"solo"}}); got != "solo" {
		t.Errorf("Chorus(one speaker) = %q, want %q", got, "solo")
	}
}

func TestDescribe(t *testing.T) {
	tests := []struct {
		in   any
		want string
	}{
		{nil, "nil"},
		{42, "int: 42"},
		{"hi", `string: "hi"`},
		{true, "bool: true"},
		{[]int{1, 2, 3}, "[]int of length 3"},
		{errors.New("boom"), "error: boom"},
		{3.14, "unknown"},
		{[]string{"x"}, "unknown"},
	}
	for _, tt := range tests {
		if got := Describe(tt.in); got != tt.want {
			t.Errorf("Describe(%#v) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestSumNumeric(t *testing.T) {
	values := []any{1, 2.5, "3", int64(4), nil, true, []int{5}}
	if got := SumNumeric(values); !almostEqual(got, 7.5) {
		t.Errorf("SumNumeric(%v) = %v, want 7.5", values, got)
	}
}

func TestByLength(t *testing.T) {
	words := []string{"banana", "kiwi", "fig", "cherry", "apple"}
	var _ sort.Interface = ByLength(words)
	sort.Stable(ByLength(words))
	if want := []string{"fig", "kiwi", "apple", "banana", "cherry"}; !slices.Equal(words, want) {
		t.Errorf("after sorting by length: %q, want %q", words, want)
	}
}

func TestQueue(t *testing.T) {
	var rw ReadWriter = &Queue{}
	if _, ok := rw.Read(); ok {
		t.Fatal("Read() on an empty queue returned ok = true")
	}
	rw.Write("a")
	rw.Write("b")
	rw.Write("c")
	for _, want := range []string{"a", "b", "c"} {
		if got, ok := rw.Read(); got != want || !ok {
			t.Errorf("Read() = %q, %t, want %q, true", got, ok, want)
		}
	}
	if _, ok := rw.Read(); ok {
		t.Error("Read() after reading everything returned ok = true")
	}
}

func TestCountingWriter(t *testing.T) {
	cw := &CountingWriter{}
	var w io.Writer = cw
	fmt.Fprintf(w, "hello")
	fmt.Fprintf(w, ", %s!", "gopher")
	if cw.Count != 14 {
		t.Errorf("Count = %d after writing \"hello, gopher!\", want 14", cw.Count)
	}
	n, err := io.Copy(cw, strings.NewReader("12345"))
	if n != 5 || err != nil || cw.Count != 19 {
		t.Errorf("io.Copy of 5 bytes: n = %d, err = %v, Count = %d, want 5, nil, 19", n, err, cw.Count)
	}
}

func TestScale(t *testing.T) {
	c := &Circle{Radius: 2}
	r := &Rect{Width: 1, Height: 3}
	var s1, s2 Scaler = c, r
	s1.Scale(1.5)
	s2.Scale(2)
	if c.Radius != 3 {
		t.Errorf("after Scale(1.5) Radius = %v, want 3", c.Radius)
	}
	if r.Width != 2 || r.Height != 6 {
		t.Errorf("after Scale(2) Rect = %+v, want {Width:2 Height:6}", *r)
	}
}

func TestScaleAll(t *testing.T) {
	c := &Circle{Radius: 1}
	r := &Rect{Width: 1, Height: 1}
	shapes := []Shape{c, fixedShape{5}, r}
	if got := ScaleAll(shapes, 10); got != 2 {
		t.Errorf("ScaleAll returned %d, want 2 (fixedShape is not a Scaler)", got)
	}
	if c.Radius != 10 || r.Width != 10 || r.Height != 10 {
		t.Errorf("after ScaleAll(10): circle = %+v, rect = %+v", *c, *r)
	}
}
