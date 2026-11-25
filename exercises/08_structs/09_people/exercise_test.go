package people

import (
	"slices"
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

func TestPeopleType(t *testing.T) {
	ps := People{{FirstName: "Ann", Age: 30}, {FirstName: "Bob", Age: 52}}
	var asSlice []Person = ps
	if len(asSlice) != 2 || asSlice[1].FirstName != "Bob" {
		t.Errorf("got %+v, want [{FirstName:Ann Age:30} {FirstName:Bob Age:52}]", ps)
	}
}

func TestOldest(t *testing.T) {
	ps := People{
		{FirstName: "Ann", Age: 30}, {FirstName: "Bob", Age: 52}, {FirstName: "Cid", Age: 52}, {FirstName: "Dee", Age: 9},
	}
	oldest := testkit.Method[func() (Person, bool)](t, &ps, "Oldest")
	if got, ok := oldest(); got.FirstName != "Bob" || !ok {
		t.Errorf("Oldest() = %+v, %t, want Bob (the first of the two 52-year-olds), true", got, ok)
	}

	var nobody People
	oldest = testkit.Method[func() (Person, bool)](t, &nobody, "Oldest")
	if got, ok := oldest(); got != (Person{}) || ok {
		t.Errorf("Oldest() of an empty People = %+v, %t, want a zero Person, false", got, ok)
	}
}

func TestSortByAge(t *testing.T) {
	ps := People{
		{FirstName: "Zoe", Age: 30}, {FirstName: "Bob", Age: 52}, {FirstName: "Amy", Age: 30}, {FirstName: "Dee", Age: 9},
	}
	sortByAge := testkit.Method[func()](t, &ps, "SortByAge")
	sortByAge()
	var got []string
	for _, p := range ps {
		got = append(got, p.FirstName)
	}
	if want := []string{"Dee", "Amy", "Zoe", "Bob"}; !slices.Equal(got, want) {
		t.Errorf("after SortByAge() the order is %q, want %q", got, want)
	}
}
