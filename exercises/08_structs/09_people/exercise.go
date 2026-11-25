//go:build !solutions

// Exercise 08.09: People
//
// Everything in this package is yours to write. Do the steps in order and run
// each step's test as you go:
//
//	go test ./exercises/08_structs/09_people -run '^TestPersonType$'
package people

// Step 1: the Person type
//
// Declare Person again, as in 01_person: FirstName and LastName (strings) and
// Age (an int). Each exercise is its own package, so you can't reuse the one
// from 01_person, and writing it again from memory is good practice.
//
// Test: TestPersonType

// Step 2: the People type
//
// Methods aren't only for structs: any type you declare can have them.
// Declare a type People whose underlying type is a slice of Person.
//
// Test: TestPeopleType

// Step 3: Oldest
//
// Give People a method Oldest() (Person, bool) that returns the oldest person
// and true. If two people share the highest age, return the first of them. If
// there's nobody, return a zero Person and false.
//
// Test: TestOldest

// Step 4: SortByAge
//
// Give People a method SortByAge() that sorts the people in place, youngest
// first. People with the same age are ordered by FirstName.
//
// Hint: look at slices.SortFunc and the cmp package. Then ask yourself which
// receiver kind SortByAge needs, and why.
//
// Test: TestSortByAge
