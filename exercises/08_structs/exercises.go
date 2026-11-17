//go:build !solutions

// Package structs covers structs, methods, constructors, embedding and struct
// tags: everything Go offers instead of classes.
//
// Go has no classes and no inheritance. Instead:
//   - a struct holds the data
//   - methods are functions with a "receiver", declared outside the struct
//   - a "constructor" is just a regular function, by convention NewXxx
//   - encapsulation is per package: lowercase names are private, Uppercase are exported
//   - code reuse comes from embedding (composition), not inheritance
//
// Quick syntax reminder:
//
//	type Point struct { X, Y int }
//	p := Point{X: 1, Y: 2}                       // struct literal
//	func (p Point) Sum() int { return p.X + p.Y } // value receiver: works on a copy
//	func (p *Point) Move(dx int) { p.X += dx }    // pointer receiver: can modify p
//	func NewPoint(x, y int) *Point { return &Point{X: x, Y: y} }
//
//	type Animal struct { Name string }
//	type Dog struct { Animal; Breed string }      // embedding: Dog gets Animal's fields & methods
//
// Rule of thumb: if any method needs a pointer receiver, use pointer receivers
// for all methods of that type.
package structs

// Person is a simple data structure.
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// Exercise 1: NewPerson
//
// Build a Person from its fields.
//
// Concepts: struct literals with field names.
func NewPerson(first, last string, age int) Person {
	// TODO: implement
	return Person{}
}

// Exercise 2: FullName
//
// Return "<FirstName> <LastName>".
//
//	Person{FirstName: "Ada", LastName: "Lovelace"}.FullName() → "Ada Lovelace"
//
// Concepts: methods with a value receiver.
func (p Person) FullName() string {
	// TODO: implement
	return ""
}

// Exercise 3: IsAdult
//
// Report whether the person is 18 or older.
func (p Person) IsAdult() bool {
	// TODO: implement
	return false
}

// Rectangle has a width and a height.
type Rectangle struct {
	Width, Height float64
}

// Exercise 4: Area and Perimeter
//
// Implement both methods.
//
//	Rectangle{Width: 3, Height: 4}.Area()      → 12
//	Rectangle{Width: 3, Height: 4}.Perimeter() → 14
func (r Rectangle) Area() float64 {
	// TODO: implement
	return 0
}

func (r Rectangle) Perimeter() float64 {
	// TODO: implement
	return 0
}

// Exercise 5: Scale
//
// Multiply both dimensions by factor, modifying the rectangle itself.
//
// Concepts: pointer receivers. Try changing the receiver to (r Rectangle)
// and watch the test fail: you'd be scaling a copy.
func (r *Rectangle) Scale(factor float64) {
	// TODO: implement
}

// Point is a position on a 2D plane.
type Point struct {
	X, Y float64
}

// Exercise 6: DistanceTo and Add
//
// DistanceTo returns the euclidean distance between p and q.
// Add returns a new Point that is the sum of p and q.
//
//	Point{0, 0}.DistanceTo(Point{3, 4}) → 5
//	Point{1, 2}.Add(Point{3, 4})        → Point{4, 6}
func (p Point) DistanceTo(q Point) float64 {
	// TODO: implement
	return 0
}

func (p Point) Add(q Point) Point {
	// TODO: implement
	return Point{}
}

// Counter keeps its count in an unexported (private) field, so code outside
// this package can only change it through its methods.
type Counter struct {
	value int
}

// Exercise 7: Counter methods
//
// Increment adds one, Reset sets the count back to zero, and Value returns
// the current count.
//
// Concepts: encapsulation with unexported fields.
func (c *Counter) Increment() {
	// TODO: implement
}

func (c *Counter) Reset() {
	// TODO: implement
}

func (c *Counter) Value() int {
	// TODO: implement
	return 0
}

// BankAccount is a classic "class" example.
type BankAccount struct {
	owner   string
	balance float64
}

// Exercise 8: BankAccount
//
//   - NewBankAccount creates an account with an initial balance.
//   - Deposit adds a positive amount; it returns false (and does nothing) if
//     amount <= 0.
//   - Withdraw removes a positive amount; it returns false (and does nothing)
//     if amount <= 0 or if the balance is insufficient.
//   - Balance and Owner are "getters".
//
// Concepts: constructors returning a pointer, methods that enforce rules.
func NewBankAccount(owner string, initial float64) *BankAccount {
	// TODO: implement
	return &BankAccount{}
}

func (a *BankAccount) Deposit(amount float64) bool {
	// TODO: implement
	return false
}

func (a *BankAccount) Withdraw(amount float64) bool {
	// TODO: implement
	return false
}

func (a *BankAccount) Balance() float64 {
	// TODO: implement
	return 0
}

func (a *BankAccount) Owner() string {
	// TODO: implement
	return ""
}

// IntStack is a last-in-first-out stack of ints. Its zero value is an empty
// stack that is ready to use.
type IntStack struct {
	items []int
}

// Exercise 9: IntStack
//
// Push adds to the top. Pop removes and returns the top (false if empty).
// Peek returns the top without removing it (false if empty). Len returns the
// number of items and IsEmpty reports whether there are none.
func (s *IntStack) Push(v int) {
	// TODO: implement
}

func (s *IntStack) Pop() (int, bool) {
	// TODO: implement
	return 0, false
}

func (s *IntStack) Peek() (int, bool) {
	// TODO: implement
	return 0, false
}

func (s *IntStack) Len() int {
	// TODO: implement
	return 0
}

func (s *IntStack) IsEmpty() bool {
	// TODO: implement
	return false
}

// Animal is embedded into Dog below.
type Animal struct {
	Name  string
	Sound string
}

// Dog embeds Animal: a Dog has Name, Sound and Speak() "for free".
type Dog struct {
	Animal
	Breed string
}

// Exercise 10: Embedding
//
// Speak returns "<Name> says <Sound>".
// NewDog builds a Dog whose Sound is always "Woof".
//
//	d := NewDog("Rex", "Beagle")
//	d.Name    → "Rex"          (promoted field from Animal)
//	d.Speak() → "Rex says Woof" (promoted method from Animal)
func (a Animal) Speak() string {
	// TODO: implement
	return ""
}

func NewDog(name, breed string) Dog {
	// TODO: implement
	return Dog{}
}

// Inventory tracks item quantities.
type Inventory struct {
	items map[string]int
}

// Exercise 11: Inventory
//
//   - NewInventory returns a ready-to-use inventory (initialize the map!).
//   - Add increases the quantity of an item.
//   - Remove decreases it and returns true, or returns false (changing
//     nothing) if there isn't enough stock. Items that reach 0 are deleted.
//   - Count returns the quantity of an item (0 if unknown).
//   - Items returns the number of distinct items in stock.
func NewInventory() *Inventory {
	// TODO: implement
	return &Inventory{}
}

func (inv *Inventory) Add(item string, qty int) {
	// TODO: implement
}

func (inv *Inventory) Remove(item string, qty int) bool {
	// TODO: implement
	return false
}

func (inv *Inventory) Count(item string) int {
	// TODO: implement
	return 0
}

func (inv *Inventory) Items() int {
	// TODO: implement
	return 0
}

// Exercise 12: Oldest
//
// Return the oldest person and true, or a zero Person and false if the slice
// is empty. On a tie, return the first one.
func Oldest(people []Person) (Person, bool) {
	// TODO: implement
	return Person{}, false
}

// Exercise 13: SortByAge
//
// Sort people in place by age, youngest first. People with the same age are
// ordered by FirstName.
//
// Concepts: slices.SortFunc (or sort.Slice) with a comparison function.
func SortByAge(people []Person) {
	// TODO: implement
}

// Book uses struct tags to control its JSON field names.
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages,omitempty"`
}

// Exercise 14: JSON
//
// BookToJSON encodes a Book as JSON. BookFromJSON decodes one.
//
//	BookToJSON(Book{Title: "Go", Author: "Rob"}) → `{"title":"Go","author":"Rob"}`
//	(pages is omitted because it's zero and the tag says omitempty)
//
// Concepts: encoding/json: json.Marshal and json.Unmarshal; struct tags.
func BookToJSON(b Book) (string, error) {
	// TODO: implement
	return "", nil
}

func BookFromJSON(s string) (Book, error) {
	// TODO: implement
	return Book{}, nil
}
