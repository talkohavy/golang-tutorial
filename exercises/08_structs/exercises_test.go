package structs

import (
	"math"
	"slices"
	"testing"
)

func TestNewPerson(t *testing.T) {
	got := NewPerson("Ada", "Lovelace", 36)
	want := Person{FirstName: "Ada", LastName: "Lovelace", Age: 36}
	if got != want {
		t.Errorf("NewPerson(\"Ada\", \"Lovelace\", 36) = %+v, want %+v", got, want)
	}
}

func TestFullName(t *testing.T) {
	p := Person{FirstName: "Ada", LastName: "Lovelace"}
	if got := p.FullName(); got != "Ada Lovelace" {
		t.Errorf("FullName() = %q, want %q", got, "Ada Lovelace")
	}
}

func TestIsAdult(t *testing.T) {
	tests := []struct {
		age  int
		want bool
	}{{17, false}, {18, true}, {40, true}, {0, false}}
	for _, tt := range tests {
		if got := (Person{Age: tt.age}).IsAdult(); got != tt.want {
			t.Errorf("Person{Age: %d}.IsAdult() = %t, want %t", tt.age, got, tt.want)
		}
	}
}

func TestAreaAndPerimeter(t *testing.T) {
	r := Rectangle{Width: 3, Height: 4}
	if got := r.Area(); got != 12 {
		t.Errorf("%+v.Area() = %v, want 12", r, got)
	}
	if got := r.Perimeter(); got != 14 {
		t.Errorf("%+v.Perimeter() = %v, want 14", r, got)
	}
}

func TestScale(t *testing.T) {
	r := Rectangle{Width: 2, Height: 5}
	r.Scale(3)
	if r.Width != 6 || r.Height != 15 {
		t.Errorf("after Scale(3) the rectangle is %+v, want {Width:6 Height:15}", r)
	}
}

func TestDistanceToAndAdd(t *testing.T) {
	if got := (Point{0, 0}).DistanceTo(Point{3, 4}); math.Abs(got-5) > 1e-9 {
		t.Errorf("Point{0, 0}.DistanceTo(Point{3, 4}) = %v, want 5", got)
	}
	if got := (Point{1, 1}).DistanceTo(Point{-2, -3}); math.Abs(got-5) > 1e-9 {
		t.Errorf("Point{1, 1}.DistanceTo(Point{-2, -3}) = %v, want 5", got)
	}
	if got := (Point{1, 2}).Add(Point{3, 4}); got != (Point{4, 6}) {
		t.Errorf("Point{1, 2}.Add(Point{3, 4}) = %+v, want {X:4 Y:6}", got)
	}
}

func TestCounter(t *testing.T) {
	var c Counter
	c.Increment()
	c.Increment()
	c.Increment()
	if got := c.Value(); got != 3 {
		t.Errorf("after 3 increments Value() = %d, want 3", got)
	}
	c.Reset()
	if got := c.Value(); got != 0 {
		t.Errorf("after Reset() Value() = %d, want 0", got)
	}
	c.Increment()
	if got := c.Value(); got != 1 {
		t.Errorf("after Reset() and 1 increment Value() = %d, want 1", got)
	}
}

func TestBankAccount(t *testing.T) {
	a := NewBankAccount("Alice", 100)
	if a == nil {
		t.Fatal("NewBankAccount returned nil")
	}
	if a.Owner() != "Alice" || a.Balance() != 100 {
		t.Fatalf("new account: Owner() = %q, Balance() = %v, want \"Alice\", 100", a.Owner(), a.Balance())
	}
	if !a.Deposit(50) || a.Balance() != 150 {
		t.Errorf("Deposit(50): balance = %v, want 150 (and Deposit should return true)", a.Balance())
	}
	if a.Deposit(-10) || a.Balance() != 150 {
		t.Errorf("Deposit(-10) must return false and not change the balance (balance = %v)", a.Balance())
	}
	if !a.Withdraw(30) || a.Balance() != 120 {
		t.Errorf("Withdraw(30): balance = %v, want 120 (and Withdraw should return true)", a.Balance())
	}
	if a.Withdraw(1000) || a.Balance() != 120 {
		t.Errorf("Withdraw(1000) must return false and not change the balance (balance = %v)", a.Balance())
	}
	if a.Withdraw(0) {
		t.Errorf("Withdraw(0) must return false")
	}
}

func TestIntStack(t *testing.T) {
	var s IntStack
	if !s.IsEmpty() || s.Len() != 0 {
		t.Fatalf("new stack: IsEmpty() = %t, Len() = %d, want true, 0", s.IsEmpty(), s.Len())
	}
	if _, ok := s.Pop(); ok {
		t.Fatal("Pop() on an empty stack returned ok = true")
	}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	if s.Len() != 3 || s.IsEmpty() {
		t.Fatalf("after 3 pushes: Len() = %d, IsEmpty() = %t, want 3, false", s.Len(), s.IsEmpty())
	}
	if v, ok := s.Peek(); v != 3 || !ok {
		t.Errorf("Peek() = %d, %t, want 3, true", v, ok)
	}
	for _, want := range []int{3, 2, 1} {
		if v, ok := s.Pop(); v != want || !ok {
			t.Errorf("Pop() = %d, %t, want %d, true", v, ok, want)
		}
	}
	if !s.IsEmpty() {
		t.Error("stack should be empty after popping everything")
	}
	if _, ok := s.Peek(); ok {
		t.Error("Peek() on an empty stack returned ok = true")
	}
}

func TestEmbedding(t *testing.T) {
	cat := Animal{Name: "Tom", Sound: "Meow"}
	if got := cat.Speak(); got != "Tom says Meow" {
		t.Errorf("Animal.Speak() = %q, want %q", got, "Tom says Meow")
	}
	d := NewDog("Rex", "Beagle")
	if d.Name != "Rex" || d.Breed != "Beagle" || d.Sound != "Woof" {
		t.Errorf("NewDog(\"Rex\", \"Beagle\") = %+v", d)
	}
	if got := d.Speak(); got != "Rex says Woof" {
		t.Errorf("Dog.Speak() = %q, want %q", got, "Rex says Woof")
	}
}

func TestInventory(t *testing.T) {
	inv := NewInventory()
	inv.Add("apple", 5)
	inv.Add("apple", 3)
	inv.Add("pear", 2)
	if got := inv.Count("apple"); got != 8 {
		t.Errorf("Count(apple) = %d, want 8", got)
	}
	if got := inv.Items(); got != 2 {
		t.Errorf("Items() = %d, want 2", got)
	}
	if inv.Remove("apple", 100) {
		t.Error("Remove(apple, 100) returned true with only 8 in stock")
	}
	if !inv.Remove("apple", 8) {
		t.Error("Remove(apple, 8) returned false with 8 in stock")
	}
	if got := inv.Count("apple"); got != 0 {
		t.Errorf("Count(apple) after removing all = %d, want 0", got)
	}
	if got := inv.Items(); got != 1 {
		t.Errorf("Items() after removing all apples = %d, want 1 (items at 0 must be deleted)", got)
	}
	if inv.Remove("banana", 1) {
		t.Error("Remove(banana, 1) returned true for an unknown item")
	}
}

func TestOldest(t *testing.T) {
	people := []Person{
		{FirstName: "Ann", Age: 30}, {FirstName: "Bob", Age: 52}, {FirstName: "Cid", Age: 52}, {FirstName: "Dee", Age: 9},
	}
	if got, ok := Oldest(people); got.FirstName != "Bob" || !ok {
		t.Errorf("Oldest() = %+v, %t, want Bob, true", got, ok)
	}
	if got, ok := Oldest(nil); got != (Person{}) || ok {
		t.Errorf("Oldest(nil) = %+v, %t, want zero Person, false", got, ok)
	}
}

func TestSortByAge(t *testing.T) {
	people := []Person{
		{FirstName: "Zoe", Age: 30}, {FirstName: "Bob", Age: 52}, {FirstName: "Amy", Age: 30}, {FirstName: "Dee", Age: 9},
	}
	SortByAge(people)
	var got []string
	for _, p := range people {
		got = append(got, p.FirstName)
	}
	if want := []string{"Dee", "Amy", "Zoe", "Bob"}; !slices.Equal(got, want) {
		t.Errorf("after SortByAge the order is %q, want %q", got, want)
	}
}

func TestJSON(t *testing.T) {
	got, err := BookToJSON(Book{Title: "Go", Author: "Rob"})
	if want := `{"title":"Go","author":"Rob"}`; got != want || err != nil {
		t.Errorf("BookToJSON(...) = %s, %v, want %s, nil", got, err, want)
	}
	got, err = BookToJSON(Book{Title: "Go", Author: "Rob", Pages: 300})
	if want := `{"title":"Go","author":"Rob","pages":300}`; got != want || err != nil {
		t.Errorf("BookToJSON(...) = %s, %v, want %s, nil", got, err, want)
	}

	b, err := BookFromJSON(`{"title":"Dune","author":"Frank","pages":412}`)
	if want := (Book{Title: "Dune", Author: "Frank", Pages: 412}); b != want || err != nil {
		t.Errorf("BookFromJSON(...) = %+v, %v, want %+v, nil", b, err, want)
	}
	if _, err := BookFromJSON(`{not json`); err == nil {
		t.Error("BookFromJSON(`{not json`) returned no error")
	}
}
