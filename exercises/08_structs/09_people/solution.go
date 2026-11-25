//go:build solutions

// SPOILER ALERT: reference solution. Try the exercise first!

package people

import (
	"cmp"
	"slices"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type People []Person

func (ps People) Oldest() (Person, bool) {
	if len(ps) == 0 {
		return Person{}, false
	}
	oldest := ps[0]
	for _, p := range ps[1:] {
		if p.Age > oldest.Age {
			oldest = p
		}
	}
	return oldest, true
}

// SortByAge works with a value receiver: the copy it receives is only a slice
// header pointing at the same backing array, so sorting the copy reorders the
// caller's elements. A pointer receiver would only be needed to change the
// slice itself (its length or backing array), as append does.
func (ps People) SortByAge() {
	slices.SortFunc(ps, func(a, b Person) int {
		return cmp.Or(cmp.Compare(a.Age, b.Age), cmp.Compare(a.FirstName, b.FirstName))
	})
}
