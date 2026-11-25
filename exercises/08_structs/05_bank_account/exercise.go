//go:build !solutions

// Exercise 08.05: BankAccount
//
// Everything in this package is yours to write. Do the steps in order and run
// each step's test as you go:
//
//	go test ./exercises/08_structs/05_bank_account -run '^TestBankAccountType$'
package bankaccount

// Step 1: the BankAccount type
//
// Declare a struct type BankAccount that holds the owner's name (a string) and
// the balance (a float64). Both fields must be unexported: code in other
// packages can only read or change them through the methods below.
//
// Test: TestBankAccountType

// Step 2: NewBankAccount
//
// Write a constructor NewBankAccount(owner string, initial float64) that
// creates an account for owner, with initial as its balance.
//
// Should it return a BankAccount or a *BankAccount? The test accepts both.
// Unlike a Person, an account gets modified by its methods, and everyone
// holding "Alice's account" should see the same balance. What does that
// suggest?
//
// Test: TestNewBankAccount

// Step 3: Owner
//
// Give BankAccount a method Owner() string that returns the owner's name.
//
// Test: TestOwner

// Step 4: Balance
//
// Give BankAccount a method Balance() float64 that returns the current
// balance.
//
// Test: TestBalance

// Step 5: Deposit
//
// Give BankAccount a method Deposit(amount float64) bool that adds amount to
// the balance and returns true. If amount isn't positive, it changes nothing
// and returns false.
//
// Test: TestDeposit

// Step 6: Withdraw
//
// Give BankAccount a method Withdraw(amount float64) bool that removes amount
// from the balance and returns true. If amount isn't positive, or the balance
// is too low, it changes nothing and returns false.
//
// Test: TestWithdraw
