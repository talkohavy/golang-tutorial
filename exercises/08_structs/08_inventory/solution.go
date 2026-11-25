//go:build solutions

// SPOILER ALERT: reference solution. Try the exercise first!

package inventory

type Inventory struct {
	items map[string]int
}

func NewInventory() *Inventory {
	return &Inventory{items: map[string]int{}}
}

// Count relies on a missing map key reading as the zero value, 0.
func (inv *Inventory) Count(item string) int {
	return inv.items[item]
}

func (inv *Inventory) Add(item string, qty int) {
	inv.items[item] += qty
}

func (inv *Inventory) Items() int {
	return len(inv.items)
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
