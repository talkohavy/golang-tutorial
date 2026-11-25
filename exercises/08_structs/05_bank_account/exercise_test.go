package bankaccount

import (
	"testing"

	"go_tutorial/internal/testkit"
)

func TestBankAccountType(t *testing.T) {
	testkit.AllFieldsUnexported[BankAccount](t)
}

func TestNewBankAccount(t *testing.T) {
	if got := testkit.Value[BankAccount](t, NewBankAccount("Alice", 100)); got == (BankAccount{}) {
		t.Error("NewBankAccount(\"Alice\", 100) returned an empty account: did you set its fields?")
	}
}

func TestOwner(t *testing.T) {
	acc := testkit.Ptr(t, NewBankAccount("Alice", 100))
	owner := testkit.Method[func() string](t, acc, "Owner")
	if got := owner(); got != "Alice" {
		t.Errorf("Owner() = %q, want %q", got, "Alice")
	}
}

func TestBalance(t *testing.T) {
	acc := testkit.Ptr(t, NewBankAccount("Alice", 100))
	balance := testkit.Method[func() float64](t, acc, "Balance")
	if got := balance(); got != 100 {
		t.Errorf("Balance() of a new account opened with 100 = %v, want 100", got)
	}
}

func TestDeposit(t *testing.T) {
	acc := testkit.Ptr(t, NewBankAccount("Alice", 100))
	deposit := testkit.Method[func(float64) bool](t, acc, "Deposit")
	balance := testkit.Method[func() float64](t, acc, "Balance")

	if !deposit(50) {
		t.Error("Deposit(50) returned false, want true")
	}
	switch got := balance(); got {
	case 150:
	case 100:
		t.Fatal("after Deposit(50) the balance is still 100: did Deposit modify a copy?")
	default:
		t.Fatalf("after Deposit(50) Balance() = %v, want 150", got)
	}

	for _, amount := range []float64{0, -10} {
		if deposit(amount) {
			t.Errorf("Deposit(%v) returned true, want false", amount)
		}
		if got := balance(); got != 150 {
			t.Errorf("Deposit(%v) changed the balance to %v, want it to stay 150", amount, got)
		}
	}
}

func TestWithdraw(t *testing.T) {
	acc := testkit.Ptr(t, NewBankAccount("Alice", 100))
	withdraw := testkit.Method[func(float64) bool](t, acc, "Withdraw")
	balance := testkit.Method[func() float64](t, acc, "Balance")

	if !withdraw(30) {
		t.Error("Withdraw(30) returned false, want true")
	}
	switch got := balance(); got {
	case 70:
	case 100:
		t.Fatal("after Withdraw(30) the balance is still 100: did Withdraw modify a copy?")
	default:
		t.Fatalf("after Withdraw(30) Balance() = %v, want 70", got)
	}

	for _, amount := range []float64{0, -10, 70.01, 1000} {
		if withdraw(amount) {
			t.Errorf("Withdraw(%v) with a balance of 70 returned true, want false", amount)
		}
		if got := balance(); got != 70 {
			t.Errorf("Withdraw(%v) changed the balance to %v, want it to stay 70", amount, got)
		}
	}

	if !withdraw(70) {
		t.Error("Withdraw(70) with a balance of 70 returned false, want true")
	}
	if got := balance(); got != 0 {
		t.Errorf("after withdrawing everything Balance() = %v, want 0", got)
	}
}
