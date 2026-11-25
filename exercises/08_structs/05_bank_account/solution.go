//go:build solutions

// SPOILER ALERT: reference solution. Try the exercise first!

package bankaccount

type BankAccount struct {
	owner   string
	balance float64
}

// NewBankAccount returns a pointer: Deposit and Withdraw modify the account,
// and every holder of the pointer shares that one account. Returning a value
// would let callers make independent copies of "the same" account, each with
// its own balance.
func NewBankAccount(owner string, initial float64) *BankAccount {
	return &BankAccount{owner: owner, balance: initial}
}

func (a *BankAccount) Owner() string {
	return a.owner
}

func (a *BankAccount) Balance() float64 {
	return a.balance
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
