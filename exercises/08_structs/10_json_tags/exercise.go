//go:build !solutions

// Exercise 08.10: JSON tags
//
// Everything in this package is yours to write. Do the steps in order and run
// each step's test as you go:
//
//	go test ./exercises/08_structs/10_json_tags -run '^TestBookType$'
package jsontags

// Step 1: the Book type
//
// Declare a struct type Book with three fields: Title and Author (strings) and
// Pages (an int). Use struct tags so that, when encoding/json encodes a Book,
// the JSON keys are "title", "author" and "pages", and "pages" is left out
// entirely when it's 0:
//
//	Book{Title: "Go", Author: "Rob"}             → {"title":"Go","author":"Rob"}
//	Book{Title: "Go", Author: "Rob", Pages: 300} → {"title":"Go","author":"Rob","pages":300}
//
// Test: TestBookType

// Step 2: ToJSON
//
// Give Book a method ToJSON() (string, error) that returns the book encoded as
// JSON.
//
// Test: TestToJSON

// Step 3: FromJSON
//
// Give Book a method FromJSON(s string) error that decodes the JSON in s into
// the book it's called on:
//
//	var b Book
//	err := b.FromJSON(`{"title":"Dune","author":"Frank","pages":412}`)
//	// b is now {Title:Dune Author:Frank Pages:412}
//
// It returns an error if s isn't valid JSON.
//
// Test: TestFromJSON
