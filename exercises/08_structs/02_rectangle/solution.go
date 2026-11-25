//go:build solutions

// SPOILER ALERT: reference solution. Try the exercise first!

package rectangle

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Scale needs a pointer receiver: with (r Rectangle) it would receive a copy,
// scale the copy, and leave the caller's rectangle untouched.
//
// Mixing receiver kinds goes against the rule of thumb in ../doc.go. It's kept
// here so you can compare both kinds side by side; in real code you'd likely
// make Area and Perimeter pointer receivers too.
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}
