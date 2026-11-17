//go:build !solutions

// Package errorsx covers error handling: returning errors, sentinel errors,
// custom error types, wrapping, errors.Is / errors.As, and panic/recover.
//
// Quick syntax reminder:
//
//	errors.New("something failed")                // a simple error
//	fmt.Errorf("parse %q: %w", s, err)            // wrap an error with context (%w)
//	if err != nil { return 0, err }               // the classic check
//	var ErrNotFound = errors.New("not found")     // sentinel error (compare with errors.Is)
//	errors.Is(err, ErrNotFound)                   // is ErrNotFound anywhere in the chain?
//	var ve *ValidationError
//	errors.As(err, &ve)                           // is there a *ValidationError in the chain?
//	errors.Join(err1, err2)                       // combine several errors
//
//	defer func() {
//		if r := recover(); r != nil { ... }    // stop a panic from crashing the program
//	}()
//
// Errors are just values that implement `interface { Error() string }`.
// By convention, error is the LAST return value.
package errorsx

import "errors"

// ErrNotFound is a sentinel error.
var ErrNotFound = errors.New("not found")

// users and emails are fake "databases" for the exercises below.
var users = map[int]string{1: "alice", 2: "bob"}
var emails = map[string]string{"alice": "alice@example.com"}

// ValidationError is a custom error type with extra information.
type ValidationError struct {
	Field   string
	Message string
}

// Exercise 1: Divide
//
// Return a / b, or an error with the message "division by zero" if b is 0.
func Divide(a, b int) (int, error) {
	// TODO: implement
	return 0, nil
}

// Exercise 2: ParseAge
//
// Convert s to an int. Return an error if s isn't a number, or if the age is
// negative or above 150.
//
//	ParseAge("42")  → 42, nil
//	ParseAge("abc") → 0, error
//	ParseAge("200") → 0, error
func ParseAge(s string) (int, error) {
	// TODO: implement
	return 0, nil
}

// Exercise 3: FindUser
//
// Look up the user's name in the users map. If the id doesn't exist, return
// the sentinel error ErrNotFound itself (so callers can use errors.Is).
func FindUser(id int) (string, error) {
	// TODO: implement
	return "", nil
}

// Exercise 4: ValidationError and ValidateUser
//
// Make *ValidationError implement the error interface; its message is
// "<Field>: <Message>", e.g. "age: must not be negative".
//
// ValidateUser returns:
//   - &ValidationError{Field: "name", Message: "must not be empty"} if name is ""
//   - &ValidationError{Field: "age", Message: "must not be negative"} if age < 0
//   - nil otherwise
func (e *ValidationError) Error() string {
	// TODO: implement
	return ""
}

func ValidateUser(name string, age int) error {
	// TODO: implement
	return nil
}

// Exercise 5: GetEmail
//
// Return the email of the user with the given id. Use FindUser, and wrap any
// error it returns with context using %w:
//
//	GetEmail(1)  → "alice@example.com", nil
//	GetEmail(99) → error "get email for user 99: not found"
//	GetEmail(2)  → error "get email for user 2: no email for bob: not found"
//
// In both error cases errors.Is(err, ErrNotFound) must be true.
func GetEmail(id int) (string, error) {
	// TODO: implement
	return "", nil
}

// Exercise 6: SafeDivide
//
// Compute a / b using the regular / operator, which PANICS when b is 0.
// Use defer + recover to turn that panic into an error (any message).
//
// Concepts: a deferred function can set named return values.
func SafeDivide(a, b int) (result int, err error) {
	// TODO: implement
	return 0, nil
}

// Exercise 7: MustPositive
//
// Return n if it's positive, otherwise panic (with any value).
//
// Concepts: the "Must" convention: functions that panic instead of returning
// an error, used when a failure is a programming bug.
func MustPositive(n int) int {
	// TODO: implement
	return 0
}

// Exercise 8: FirstError
//
// Call each function in order and return the first error. Functions after
// the failing one must NOT be called. Return nil if all succeed.
func FirstError(fns ...func() error) error {
	// TODO: implement
	return nil
}

// Exercise 9: ValidateAll
//
// Return one error per negative number, combined with errors.Join, or nil if
// there are none. Each error's message is "index <i>: negative value <n>".
//
//	ValidateAll([]int{1, -2, 3, -4}).Error() →
//	"index 1: negative value -2\nindex 3: negative value -4"
func ValidateAll(nums []int) error {
	// TODO: implement
	return nil
}

// Exercise 10: Retry
//
// Call f up to attempts times, stopping as soon as it succeeds (returns nil).
// If every attempt fails, return an error that wraps the LAST error:
//
//	"failed after 3 attempts: <last error>"
func Retry(attempts int, f func() error) error {
	// TODO: implement
	return nil
}

// Exercise 11: Classify
//
// Inspect an error and return:
//   - "none"       if err is nil
//   - "not found"  if ErrNotFound is anywhere in its chain (errors.Is)
//   - "validation" if a *ValidationError is anywhere in its chain (errors.As)
//   - "unknown"    otherwise
func Classify(err error) string {
	// TODO: implement
	return ""
}
