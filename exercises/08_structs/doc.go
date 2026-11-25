// Package structs is the cheat sheet for topic 08. The exercises live in the
// numbered folders next to this file, one small package each:
//
//	01_person        struct types, struct literals, constructors, methods
//	02_rectangle     value receivers vs pointer receivers
//	03_point         methods that take and return their own type
//	04_counter       encapsulation, a useful zero value
//	05_bank_account  constructors, methods that enforce rules
//	06_int_stack     a struct wrapping a slice
//	07_embedding     composition instead of inheritance
//	08_inventory     a struct wrapping a map (and why it needs a constructor)
//	09_people        methods on a named slice type
//	10_json_tags     struct tags and encoding/json
//
// Each folder has:
//
//	exercise.go       your blank canvas, with a description of each step
//	exercise_test.go  sanity checks, one test per step
//	solution.go       a reference solution (spoilers!)
//
// You write every type, constructor and method yourself. Do the steps in
// order and run each step's test as you go. A compile error such as
// "undefined: Person" means the tests are waiting for something you haven't
// written yet; "Person has no Area method yet" means the same for a method.
//
// Go has no classes and no inheritance. Instead:
//   - a struct holds the data
//   - methods are functions with a "receiver", declared outside the struct
//   - a "constructor" is just a regular function, by convention NewXxx
//   - encapsulation is per package: lowercase names are private, Uppercase are exported
//   - code reuse comes from embedding (composition), not inheritance
//
// Syntax reminder:
//
//	type Circle struct { Radius float64 }
//	c := Circle{Radius: 2}                                 // struct literal
//	func (c Circle) Diameter() float64 { return 2 * c.Radius } // value receiver: works on a copy
//	func (c *Circle) Grow(by float64) { c.Radius += by }       // pointer receiver: can modify c
//	func NewCircle(r float64) *Circle { return &Circle{Radius: r} }
//
//	type Engine struct { HP int }
//	type Car struct { Engine; Model string }               // embedding: Car gets Engine's fields & methods
//	car := Car{Engine: Engine{HP: 300}, Model: "GT"}       // literal: set the embedded struct by its type name
//	car.HP                                                 // promoted field: same as car.Engine.HP
//
// Rule of thumb: if any method needs a pointer receiver, use pointer receivers
// for all methods of that type.
package structs
