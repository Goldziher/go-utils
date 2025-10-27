// This package includes utility functions for error handling and exceptions.
// It provides helpers for common error handling patterns including panic-on-error,
// error recovery, and safe error checking.

package exc

import (
	"errors"
	"fmt"
)

// MustResult panics if the error is not nil, otherwise returns the result.
// This is useful for cases where an error should never occur and you want to fail fast.
// The optional messages are formatted and included in the panic message.
func MustResult[T any](value T, err error, messages ...string) T {
	if err != nil {
		msg := formatMessages(messages)
		if msg != "" {
			panic(fmt.Errorf("%s: %w", msg, err))
		}
		panic(err)
	}
	return value
}

// Must panics if the error is not nil.
// This is useful for initialization or setup code where errors should never occur.
// The optional messages are formatted and included in the panic message.
func Must(err error, messages ...string) {
	if err != nil {
		msg := formatMessages(messages)
		if msg != "" {
			panic(fmt.Errorf("%s: %w", msg, err))
		}
		panic(err)
	}
}

// Try executes a function and returns its error.
// This is useful for defer statements or cleanup operations where you want to capture errors.
func Try(fn func() error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = e
			} else {
				err = fmt.Errorf("panic: %v", r)
			}
		}
	}()
	return fn()
}

// Catch executes a function and recovers from any panics, converting them to errors.
// Returns the function's error if any, or an error created from a recovered panic.
func Catch(fn func() error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = e
			} else {
				err = fmt.Errorf("panic: %v", r)
			}
		}
	}()
	return fn()
}

// IgnoreErr executes a function that returns an error and ignores the error.
// This is useful for cleanup operations where errors can be safely ignored.
func IgnoreErr(fn func() error) {
	_ = fn() //nolint:errcheck // explicitly ignoring error by design
}

// FirstErr returns the first non-nil error from a list of errors.
// Returns nil if all errors are nil.
func FirstErr(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

// ReturnAnyErr returns the first error found in the list of values, if any.
// This is useful when dealing with variadic functions that may return errors.
func ReturnAnyErr(values ...any) error {
	for _, value := range values {
		if err, ok := value.(error); ok && err != nil {
			return err
		}
	}
	return nil
}

// ReturnNotNil panics if the value is nil, otherwise returns the value.
// This is useful for asserting that a pointer must not be nil.
// The optional messages are formatted and included in the panic message.
func ReturnNotNil[T any](value *T, messages ...string) *T {
	if value == nil {
		msg := formatMessages(messages)
		if msg == "" {
			msg = "unexpected nil value"
		}
		panic(errors.New(msg))
	}
	return value
}

// MustNotNil panics with a formatted error message if the value is nil.
// Returns the dereferenced value if not nil.
func MustNotNil[T any](value *T, messages ...string) T {
	if value == nil {
		msg := formatMessages(messages)
		if msg == "" {
			msg = "unexpected nil value"
		}
		panic(errors.New(msg))
	}
	return *value
}

// AllErr returns an error containing all non-nil errors from the list.
// Returns nil if all errors are nil.
// Multiple errors are joined using errors.Join (Go 1.20+).
func AllErr(errs ...error) error {
	var nonNilErrs []error
	for _, err := range errs {
		if err != nil {
			nonNilErrs = append(nonNilErrs, err)
		}
	}
	if len(nonNilErrs) == 0 {
		return nil
	}
	return errors.Join(nonNilErrs...)
}

// formatMessages joins multiple messages into a single string.
func formatMessages(messages []string) string {
	if len(messages) == 0 {
		return ""
	}
	result := messages[0]
	for i := 1; i < len(messages); i++ {
		if messages[i] != "" {
			result += " " + messages[i]
		}
	}
	return result
}
