// Package ers provides a custom error handling mechanism that implements the error
// interface using Error and StackLine structs. It offers automatic stack trace
// collection and error wrapping with context details.
package ers

import (
	"fmt"
)

// New creates a new Error struct with the provided message and automatically captures
// the current stack location. It supports fmt-style formatting through formatTags.
//
// Parameters:
// - fmessage: The error message format string
// - formatTags: Optional formatting arguments for the message
//
// Returns an Error struct containing the formatted message and current stack location.
func New(fmessage string, formatTags ...any) error {
	stack := getCaller(2)
	return &Error{
		error: fmt.Errorf(fmessage, formatTags...),
		stackTrace: []StackLine{
			stack,
		},
		contexts: []string{},
	}
}

// Wrap enhances an existing error with additional context details and the current
// stack location. If the input error is already an Error struct, it extends its
// stack trace and adds the new details. Otherwise, it creates a new Error struct
// wrapping the original error.
//
// Parameters:
// - err: The existing error to wrap
// - details: Variable number of strings providing additional context
//
// Returns an Error struct that can be unwrapped to access the original error.
// Returns nil if the input error is nil.
func Wrap(err error, details ...string) error {
	if err == nil {
		return nil
	}
	stack := getCaller(2)
	switch err := err.(type) {
	case *Error:
		err.addToStack(stack)
		err.AddContext(details...)
		return err
	default:
		return &Error{
			error: err,
			stackTrace: []StackLine{
				stack,
			},
			contexts: details,
		}
	}
}

// Wrapf wraps an error with a formatted message. It behaves like Wrap but accepts
// fmt-style formatting for the context message.
//
// Parameters:
// - err: The existing error to wrap
// - fmessage: Format string for the context message
// - formatTags: Optional formatting arguments
//
// Returns an Error struct with the formatted context message.
// Returns nil if the input error is nil.
func Wrapf(err error, fmessage string, formatTags ...any) error {
	if err == nil {
		return nil
	}
	stack := getCaller(2)
	switch err := err.(type) {
	case *Error:
		err.addToStack(stack)
		err.AddContext(fmt.Sprintf(fmessage, formatTags...))
		return err
	default:
		return &Error{
			error: err,
			stackTrace: []StackLine{
				stack,
			},
			contexts: []string{
				fmt.Sprintf(fmessage, formatTags...),
			},
		}
	}
}
