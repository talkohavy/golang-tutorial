//go:build solutions

// SPOILER ALERT: reference solutions. Try the exercises first!

package structs

import (
	"cmp"
	"encoding/json"
	"math"
	"slices"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func NewPerson(first, last string, age int) Person {
	return Person{FirstName: first, LastName: last, Age: age}
}

func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

func (p Person) IsAdult() bool {
	return p.Age >= 18
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

type Point struct {
	X, Y float64
}

func (p Point) DistanceTo(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Add(q Point) Point {
	return Point{X: p.X + q.X, Y: p.Y + q.Y}
}

type Counter struct {
	value int
}

func (c *Counter) Increment() {
	c.value++
}

func (c *Counter) Reset() {
	c.value = 0
}

func (c *Counter) Value() int {
	return c.value
}

type BankAccount struct {
	owner   string
	balance float64
}

func NewBankAccount(owner string, initial float64) *BankAccount {
	return &BankAccount{owner: owner, balance: initial}
}

func (a *BankAccount) Deposit(amount float64) bool {
	if amount <= 0 {
		return false
	}
	a.balance += amount
	return true
}

func (a *BankAccount) Withdraw(amount float64) bool {
	if amount <= 0 || amount > a.balance {
		return false
	}
	a.balance -= amount
	return true
}

func (a *BankAccount) Balance() float64 {
	return a.balance
}

func (a *BankAccount) Owner() string {
	return a.owner
}

type IntStack struct {
	items []int
}

func (s *IntStack) Push(v int) {
	s.items = append(s.items, v)
}

func (s *IntStack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}

func (s *IntStack) Peek() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	return s.items[len(s.items)-1], true
}

func (s *IntStack) Len() int {
	return len(s.items)
}

func (s *IntStack) IsEmpty() bool {
	return len(s.items) == 0
}

type Animal struct {
	Name  string
	Sound string
}

type Dog struct {
	Animal
	Breed string
}

func (a Animal) Speak() string {
	return a.Name + " says " + a.Sound
}

func NewDog(name, breed string) Dog {
	return Dog{Animal: Animal{Name: name, Sound: "Woof"}, Breed: breed}
}

type Inventory struct {
	items map[string]int
}

func NewInventory() *Inventory {
	return &Inventory{items: map[string]int{}}
}

func (inv *Inventory) Add(item string, qty int) {
	inv.items[item] += qty
}

func (inv *Inventory) Remove(item string, qty int) bool {
	if inv.items[item] < qty {
		return false
	}
	inv.items[item] -= qty
	if inv.items[item] == 0 {
		delete(inv.items, item)
	}
	return true
}

func (inv *Inventory) Count(item string) int {
	return inv.items[item]
}

func (inv *Inventory) Items() int {
	return len(inv.items)
}

func Oldest(people []Person) (Person, bool) {
	if len(people) == 0 {
		return Person{}, false
	}
	oldest := people[0]
	for _, p := range people[1:] {
		if p.Age > oldest.Age {
			oldest = p
		}
	}
	return oldest, true
}

func SortByAge(people []Person) {
	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Or(cmp.Compare(a.Age, b.Age), cmp.Compare(a.FirstName, b.FirstName))
	})
}

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages,omitempty"`
}

func BookToJSON(b Book) (string, error) {
	data, err := json.Marshal(b)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func BookFromJSON(s string) (Book, error) {
	var b Book
	err := json.Unmarshal([]byte(s), &b)
	return b, err
}
