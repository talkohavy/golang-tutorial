//go:build solutions

// SPOILER ALERT: reference solution. Try the exercise first!

package counter

// The zero value of an int is 0, so a zero Counter is already a valid counter
// at 0: no constructor needed.
type Counter struct {
	value int
}

// Value only reads, but it uses a pointer receiver anyway: Increment and Reset
// need one, and the rule of thumb is to keep all methods of a type consistent.
func (c *Counter) Value() int {
	return c.value
}

func (c *Counter) Increment() {
	c.value++
}

func (c *Counter) Reset() {
	c.value = 0
}
