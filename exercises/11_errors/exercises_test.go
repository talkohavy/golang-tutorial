package errorsx

import (
	"errors"
	"fmt"
	"testing"
)

func TestDivide(t *testing.T) {
	if got, err := Divide(10, 3); got != 3 || err != nil {
		t.Errorf("Divide(10, 3) = %d, %v, want 3, nil", got, err)
	}
	got, err := Divide(1, 0)
	if err == nil {
		t.Fatal("Divide(1, 0) returned no error")
	}
	if got != 0 || err.Error() != "division by zero" {
		t.Errorf("Divide(1, 0) = %d, %q, want 0, \"division by zero\"", got, err)
	}
}

func TestParseAge(t *testing.T) {
	if got, err := ParseAge("42"); got != 42 || err != nil {
		t.Errorf("ParseAge(\"42\") = %d, %v, want 42, nil", got, err)
	}
	if got, err := ParseAge("0"); got != 0 || err != nil {
		t.Errorf("ParseAge(\"0\") = %d, %v, want 0, nil", got, err)
	}
	for _, bad := range []string{"abc", "", "12.5", "-1", "151", "200"} {
		if got, err := ParseAge(bad); err == nil || got != 0 {
			t.Errorf("ParseAge(%q) = %d, %v, want 0 and an error", bad, got, err)
		}
	}
}

func TestFindUser(t *testing.T) {
	if got, err := FindUser(2); got != "bob" || err != nil {
		t.Errorf("FindUser(2) = %q, %v, want \"bob\", nil", got, err)
	}
	got, err := FindUser(42)
	if err != ErrNotFound || got != "" {
		t.Errorf("FindUser(42) = %q, %v, want \"\", ErrNotFound", got, err)
	}
}

func TestValidateUser(t *testing.T) {
	if got := (&ValidationError{Field: "age", Message: "must not be negative"}).Error(); got != "age: must not be negative" {
		t.Errorf("ValidationError.Error() = %q, want %q", got, "age: must not be negative")
	}
	if err := ValidateUser("ann", 30); err != nil {
		t.Errorf("ValidateUser(\"ann\", 30) = %v, want nil", err)
	}
	tests := []struct {
		name  string
		age   int
		field string
	}{
		{"", 30, "name"}, {"ann", -1, "age"},
	}
	for _, tt := range tests {
		err := ValidateUser(tt.name, tt.age)
		var ve *ValidationError
		if !errors.As(err, &ve) {
			t.Errorf("ValidateUser(%q, %d) = %v, want a *ValidationError", tt.name, tt.age, err)
			continue
		}
		if ve.Field != tt.field {
			t.Errorf("ValidateUser(%q, %d) has Field %q, want %q", tt.name, tt.age, ve.Field, tt.field)
		}
	}
}

func TestGetEmail(t *testing.T) {
	if got, err := GetEmail(1); got != "alice@example.com" || err != nil {
		t.Errorf("GetEmail(1) = %q, %v, want \"alice@example.com\", nil", got, err)
	}
	tests := []struct {
		id      int
		wantMsg string
	}{
		{99, "get email for user 99: not found"},
		{2, "get email for user 2: no email for bob: not found"},
	}
	for _, tt := range tests {
		_, err := GetEmail(tt.id)
		if err == nil {
			t.Errorf("GetEmail(%d) returned no error", tt.id)
			continue
		}
		if err.Error() != tt.wantMsg {
			t.Errorf("GetEmail(%d) error = %q, want %q", tt.id, err, tt.wantMsg)
		}
		if !errors.Is(err, ErrNotFound) {
			t.Errorf("GetEmail(%d): errors.Is(err, ErrNotFound) = false; did you wrap with %%w?", tt.id)
		}
	}
}

func TestSafeDivide(t *testing.T) {
	if got, err := SafeDivide(9, 3); got != 3 || err != nil {
		t.Errorf("SafeDivide(9, 3) = %d, %v, want 3, nil", got, err)
	}
	var (
		got int
		err error
	)
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("SafeDivide(1, 0) panicked (%v): it should recover and return an error", r)
			}
		}()
		got, err = SafeDivide(1, 0)
	}()
	if err == nil || got != 0 {
		t.Errorf("SafeDivide(1, 0) = %d, %v, want 0 and an error", got, err)
	}
}

func panics(f func()) (didPanic bool) {
	defer func() {
		if recover() != nil {
			didPanic = true
		}
	}()
	f()
	return false
}

func TestMustPositive(t *testing.T) {
	if got := MustPositive(5); got != 5 {
		t.Errorf("MustPositive(5) = %d, want 5", got)
	}
	for _, n := range []int{0, -3} {
		if !panics(func() { MustPositive(n) }) {
			t.Errorf("MustPositive(%d) did not panic", n)
		}
	}
}

func TestFirstError(t *testing.T) {
	var calls []string
	step := func(name string, err error) func() error {
		return func() error {
			calls = append(calls, name)
			return err
		}
	}
	boom := errors.New("boom")
	err := FirstError(step("a", nil), step("b", boom), step("c", errors.New("other")))
	if err != boom {
		t.Errorf("FirstError returned %v, want boom", err)
	}
	if fmt.Sprint(calls) != "[a b]" {
		t.Errorf("functions called: %v, want [a b] (stop after the first error)", calls)
	}
	calls = nil
	if err := FirstError(step("x", nil), step("y", nil)); err != nil || len(calls) != 2 {
		t.Errorf("FirstError(all ok) = %v with calls %v, want nil with calls [x y]", err, calls)
	}
}

func TestValidateAll(t *testing.T) {
	if err := ValidateAll([]int{1, 2, 3}); err != nil {
		t.Errorf("ValidateAll([1 2 3]) = %v, want nil", err)
	}
	err := ValidateAll([]int{1, -2, 3, -4})
	want := "index 1: negative value -2\nindex 3: negative value -4"
	if err == nil || err.Error() != want {
		t.Errorf("ValidateAll([1 -2 3 -4]) = %q, want %q", err, want)
	}
}

func TestRetry(t *testing.T) {
	calls := 0
	err := Retry(5, func() error {
		calls++
		if calls < 3 {
			return errors.New("not yet")
		}
		return nil
	})
	if err != nil || calls != 3 {
		t.Errorf("Retry succeeding on the 3rd try: err = %v, calls = %d, want nil, 3", err, calls)
	}

	calls = 0
	last := errors.New("still broken")
	err = Retry(3, func() error {
		calls++
		if calls == 3 {
			return last
		}
		return errors.New("broken")
	})
	if calls != 3 {
		t.Errorf("Retry(3, alwaysFails) called f %d times, want 3", calls)
	}
	if err == nil || err.Error() != "failed after 3 attempts: still broken" || !errors.Is(err, last) {
		t.Errorf("Retry(3, alwaysFails) = %v, want \"failed after 3 attempts: still broken\" wrapping the last error", err)
	}
}

func TestClassify(t *testing.T) {
	tests := []struct {
		err  error
		want string
	}{
		{nil, "none"},
		{ErrNotFound, "not found"},
		{fmt.Errorf("wrapped: %w", ErrNotFound), "not found"},
		{&ValidationError{Field: "x", Message: "bad"}, "validation"},
		{fmt.Errorf("ctx: %w", &ValidationError{Field: "x", Message: "bad"}), "validation"},
		{errors.New("other"), "unknown"},
	}
	for _, tt := range tests {
		if got := Classify(tt.err); got != tt.want {
			t.Errorf("Classify(%v) = %q, want %q", tt.err, got, tt.want)
		}
	}
}
