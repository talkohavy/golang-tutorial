//go:build !solutions

// Exercise 08.08: Inventory
//
// Everything in this package is yours to write. Do the steps in order and run
// each step's test as you go:
//
//	go test ./exercises/08_structs/08_inventory -run '^TestInventoryType$'
package inventory

// Step 1: the Inventory type
//
// Declare a struct type Inventory that tracks a quantity (an int) per item
// name (a string). Store them in an unexported map field.
//
// Test: TestInventoryType

// Step 2: NewInventory
//
// Write a constructor NewInventory() that returns an empty inventory, ready to
// use. Returning an Inventory or an *Inventory is your call.
//
// Unlike Counter and IntStack, the zero value of Inventory isn't usable: the
// zero value of a map is nil, and writing to a nil map panics. That's why this
// type needs a constructor.
//
// Test: TestNewInventory

// Step 3: Count
//
// Give Inventory a method Count(item string) int that returns the quantity of
// an item, or 0 for an item it has never seen.
//
// Test: TestCount

// Step 4: Add
//
// Give Inventory a method Add(item string, qty int) that increases the
// quantity of an item by qty.
//
// Test: TestAdd

// Step 5: Items
//
// Give Inventory a method Items() int that returns the number of distinct
// items in stock.
//
// Test: TestItems

// Step 6: Remove
//
// Give Inventory a method Remove(item string, qty int) bool that decreases the
// quantity of an item by qty and returns true. If there isn't enough stock, it
// changes nothing and returns false. An item whose quantity reaches 0 is no
// longer in stock, so Items stops counting it.
//
// Test: TestRemove
