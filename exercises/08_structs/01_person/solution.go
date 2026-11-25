//go:build solutions

// SPOILER ALERT: reference solution. Try the exercise first!

package person

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// NewPerson returns a Person value rather than a pointer: a Person is small,
// none of its methods modify it, and copying it is cheap and safe. The
// standard library does the same for small data (time.Date returns a
// time.Time). Return a pointer when the methods modify the value, when it's
// large, or when everyone must share one instance (see 05_bank_account).
func NewPerson(first, last string, age int) Person {
	return Person{FirstName: first, LastName: last, Age: age}
}

func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

func (p Person) IsAdult() bool {
	return p.Age >= 18
}
