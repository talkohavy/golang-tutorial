//go:build solutions

// SPOILER ALERT: reference solutions. Try the exercises first!

package errorsx

import (
	"errors"
	"fmt"
	"strconv"
)

var ErrNotFound = errors.New("not found")

var users = map[int]string{1: "alice", 2: "bob"}
var emails = map[string]string{"alice": "alice@example.com"}

type ValidationError struct {
	Field   string
	Message string
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func ParseAge(s string) (int, error) {
	age, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("invalid age %q: %w", s, err)
	}
	if age < 0 || age > 150 {
		return 0, fmt.Errorf("age %d out of range", age)
	}
	return age, nil
}

func FindUser(id int) (string, error) {
	name, ok := users[id]
	if !ok {
		return "", ErrNotFound
	}
	return name, nil
}

func (e *ValidationError) Error() string {
	return e.Field + ": " + e.Message
}

func ValidateUser(name string, age int) error {
	if name == "" {
		return &ValidationError{Field: "name", Message: "must not be empty"}
	}
	if age < 0 {
		return &ValidationError{Field: "age", Message: "must not be negative"}
	}
	return nil
}

func GetEmail(id int) (string, error) {
	name, err := FindUser(id)
	if err != nil {
		return "", fmt.Errorf("get email for user %d: %w", id, err)
	}
	email, ok := emails[name]
	if !ok {
		return "", fmt.Errorf("get email for user %d: no email for %s: %w", id, name, ErrNotFound)
	}
	return email, nil
}

func SafeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered: %v", r)
		}
	}()
	return a / b, nil
}

func MustPositive(n int) int {
	if n <= 0 {
		panic(fmt.Sprintf("MustPositive: %d is not positive", n))
	}
	return n
}

func FirstError(fns ...func() error) error {
	for _, f := range fns {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}

func ValidateAll(nums []int) error {
	var errs []error
	for i, n := range nums {
		if n < 0 {
			errs = append(errs, fmt.Errorf("index %d: negative value %d", i, n))
		}
	}
	return errors.Join(errs...)
}

func Retry(attempts int, f func() error) error {
	var err error
	for range attempts {
		if err = f(); err == nil {
			return nil
		}
	}
	return fmt.Errorf("failed after %d attempts: %w", attempts, err)
}

func Classify(err error) string {
	var ve *ValidationError
	switch {
	case err == nil:
		return "none"
	case errors.Is(err, ErrNotFound):
		return "not found"
	case errors.As(err, &ve):
		return "validation"
	default:
		return "unknown"
	}
}
