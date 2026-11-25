package person

import (
	"testing"

	"go_tutorial/internal/testkit"
)

func TestPersonType(t *testing.T) {
	p := Person{FirstName: "Ada", LastName: "Lovelace", Age: 36}
	var first, last string = p.FirstName, p.LastName
	var age int = p.Age
	if first != "Ada" || last != "Lovelace" || age != 36 {
		t.Errorf("got %+v, want {FirstName:Ada LastName:Lovelace Age:36}", p)
	}
}

func TestNewPerson(t *testing.T) {
	got := testkit.Value[Person](t, NewPerson("Ada", "Lovelace", 36))
	want := Person{FirstName: "Ada", LastName: "Lovelace", Age: 36}
	if got != want {
		t.Errorf("NewPerson(\"Ada\", \"Lovelace\", 36) = %+v, want %+v", got, want)
	}
}

func TestFullName(t *testing.T) {
	p := Person{FirstName: "Ada", LastName: "Lovelace", Age: 36}
	fullName := testkit.Method[func() string](t, &p, "FullName")
	if got := fullName(); got != "Ada Lovelace" {
		t.Errorf("FullName() = %q, want %q", got, "Ada Lovelace")
	}
}

func TestIsAdult(t *testing.T) {
	tests := []struct {
		age  int
		want bool
	}{{0, false}, {17, false}, {18, true}, {40, true}}
	for _, tt := range tests {
		p := Person{FirstName: "Sam", Age: tt.age}
		isAdult := testkit.Method[func() bool](t, &p, "IsAdult")
		if got := isAdult(); got != tt.want {
			t.Errorf("IsAdult() with Age %d = %t, want %t", tt.age, got, tt.want)
		}
	}
}
