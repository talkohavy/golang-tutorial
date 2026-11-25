//go:build !solutions

// Exercise 08.01: Person
//
// Everything in this package is yours to write. Do the steps in order and run
// each step's test as you go:
//
//	go test ./exercises/08_structs/01_person -run '^TestPersonType$'
package person

// Step 1: the Person type
//
// Declare a struct type Person with three fields: FirstName and LastName
// (strings) and Age (an int).
//
// Test: TestPersonType

// Step 2: NewPerson
//
// Write a constructor NewPerson(first, last string, age int) that builds a
// Person from its arguments.
//
// Should it return a Person or a *Person? Your call: the test accepts both.
// Decide which one fits a small piece of data like this before you peek at
// the solution.
//
// Test: TestNewPerson

// Step 3: FullName
//
// Give Person a method FullName() string that returns the first and last name
// separated by a space. For Ada Lovelace, it returns "Ada Lovelace".
//
// Test: TestFullName

// Step 4: IsAdult
//
// Give Person a method IsAdult() bool that reports whether the person is 18
// or older.
//
// Test: TestIsAdult
