//go:build !solutions

// Exercise 08.07: Embedding
//
// Everything in this package is yours to write. Do the steps in order and run
// each step's test as you go:
//
//	go test ./exercises/08_structs/07_embedding -run '^TestAnimalType$'
package embedding

// Step 1: the Animal type
//
// Declare a struct type Animal with two string fields: Name and Sound.
//
// Test: TestAnimalType

// Step 2: Speak
//
// Give Animal a method Speak() string that returns "<Name> says <Sound>".
// For an animal named Tom whose sound is Meow, it returns "Tom says Meow".
//
// Test: TestSpeak

// Step 3: the Dog type
//
// Declare a struct type Dog that embeds Animal (don't give the field a name)
// and adds a Breed string field. Through embedding, a Dog gets Animal's fields
// and methods: d.Name works just like d.Animal.Name.
//
// Test: TestDogType

// Step 4: NewDog
//
// Write a constructor NewDog(name, breed string) for a Dog whose Sound is
// always "Woof". As with NewPerson, returning a Dog or a *Dog is your call.
//
// Watch out: go.mod says go 1.26, and before Go 1.27 a struct literal can't
// name a promoted field, so Dog{Name: name} won't compile. Set the embedded
// struct through its type name instead.
//
// Test: TestNewDog

// Step 5: a Dog speaks
//
// NewDog("Rex", "Beagle").Speak() should return "Rex says Woof". Don't write
// any new code until you've run the test, and make sure you can explain why it
// already passes.
//
// Test: TestDogSpeaks
