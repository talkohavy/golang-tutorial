//go:build solutions

// SPOILER ALERT: reference solution. Try the exercise first!

package embedding

type Animal struct {
	Name  string
	Sound string
}

func (a Animal) Speak() string {
	return a.Name + " says " + a.Sound
}

// Dog has no Speak method of its own: d.Speak() is promoted from the embedded
// Animal, and is shorthand for d.Animal.Speak(). That's why step 5 needs no
// code. It's composition, not inheritance: Speak's receiver is the Animal
// inside the Dog, so it can't see Breed.
type Dog struct {
	Animal
	Breed string
}

func NewDog(name, breed string) Dog {
	return Dog{Animal: Animal{Name: name, Sound: "Woof"}, Breed: breed}
}
