package embedding

import (
	"testing"

	"go_tutorial/internal/testkit"
)

func TestAnimalType(t *testing.T) {
	a := Animal{Name: "Tom", Sound: "Meow"}
	var name, sound string = a.Name, a.Sound
	if name != "Tom" || sound != "Meow" {
		t.Errorf("got %+v, want {Name:Tom Sound:Meow}", a)
	}
}

func TestSpeak(t *testing.T) {
	a := Animal{Name: "Tom", Sound: "Meow"}
	speak := testkit.Method[func() string](t, &a, "Speak")
	if got := speak(); got != "Tom says Meow" {
		t.Errorf("Speak() = %q, want %q", got, "Tom says Meow")
	}
}

func TestDogType(t *testing.T) {
	d := Dog{Animal: Animal{Name: "Rex", Sound: "Woof"}, Breed: "Beagle"}
	var breed string = d.Breed
	if d.Name != "Rex" || d.Sound != "Woof" || breed != "Beagle" {
		t.Errorf("got %+v, want {Animal:{Name:Rex Sound:Woof} Breed:Beagle}", d)
	}
}

func TestNewDog(t *testing.T) {
	got := testkit.Value[Dog](t, NewDog("Rex", "Beagle"))
	want := Dog{Animal: Animal{Name: "Rex", Sound: "Woof"}, Breed: "Beagle"}
	if got != want {
		t.Errorf("NewDog(\"Rex\", \"Beagle\") = %+v, want %+v", got, want)
	}
}

func TestDogSpeaks(t *testing.T) {
	d := testkit.Ptr(t, NewDog("Rex", "Beagle"))
	speak := testkit.Method[func() string](t, d, "Speak")
	if got := speak(); got != "Rex says Woof" {
		t.Errorf("Speak() = %q, want %q", got, "Rex says Woof")
	}
}
