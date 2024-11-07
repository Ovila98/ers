// Package ers provides a custom error handling mechanism that implements the error
// interface using Error and StackLine structs. It offers automatic stack trace
// collection and error wrapping with context details.
package ers

import (
	"fmt"
	"runtime"
)

// getCaller retrieves the file and line number of the calling function at a given
// stack depth, represented by the 'skip' parameter. The information is stored in
// a StackLine struct for consistent tracking.
//
// Parameters:
// - skip: The number of stack frames to skip when retrieving the caller.
//
// Returns a StackLine struct containing the file path and line number, or default
// values if retrieval fails.
func getCaller(skip int) StackLine {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return StackLine{
			File: "unknown",
			Line: 0,
		}
	}
	return StackLine{
		File: file,
		Line: line,
	}
}

// New creates a new Error struct with the provided message and automatically captures
// the current stack location. The returned error implements both the error interface
// and supports unwrapping.
//
// Parameters:
// - message: The error message to be included.
//
// Returns an Error struct containing the message and current stack location.
func New(message string) error {
	stack := getCaller(2)
	return &Error{
		error: fmt.Errorf("%s", message),
		stackTrace: []StackLine{
			stack,
		},
		additionalInfo: []string{},
	}
}

// Wrap enhances an existing error with additional context details and the current
// stack location. If the input error is already an Error struct, it extends its
// stack trace and adds the new details. Otherwise, it creates a new Error struct
// wrapping the original error.
//
// Parameters:
// - err: The existing error to wrap.
// - details: Variable number of strings providing additional context.
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
		err.stackTrace = append(err.stackTrace, stack)
		err.AddInfo(details...)
		return err
	default:
		return &Error{
			error: err,
			stackTrace: []StackLine{
				stack,
			},
			additionalInfo: details,
		}
	}
}
