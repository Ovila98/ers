// Package ers provides a custom error handling mechanism that includes stack tracing
// and error wrapping with context. It simplifies error tracking by appending the
// caller's file and line number to error messages.
package ers

import (
	"fmt"
	"runtime"
)

// getCaller retrieves the file and line number of the calling function at a given
// stack depth, represented by the 'skip' parameter. This is used to append the
// source location to the error messages.
//
// Parameters:
// - skip: The number of stack frames to skip when retrieving the caller.
//
// Returns a string containing the file and line number, or a default message if
// unable to retrieve them.
func getCaller(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "(unable to get line number)"
	}
	return fmt.Sprintf("(%s:%d)", file, line)
}

// New creates a new error with a message and appends the source file and line number
// of the code that called this function. It's useful for error creation with automatic
// tracking of where the error was generated.
//
// Parameters:
// - message: The error message to be included.
//
// Returns a new error with the provided message and the caller's file/line information.
func New(message string) error {
	return fmt.Errorf("%s\n-> %s", message, getCaller(2))
}

// Trace wraps an existing error by appending the file and line number of the
// location where Trace was called. This is useful for adding context to an existing
// error without altering its message.
//
// Parameters:
// - err: The error to be traced.
//
// Returns a new error that wraps the provided error with additional context.
// If 'err' is nil, Trace returns nil.
func Trace(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%w\n-> %s", err, getCaller(2))
}

// Wrap wraps an existing error with a new message and appends the file and line
// number of the location where Wrap was called. This adds both new context
// (via the message) and stack trace information.
//
// Parameters:
// - err: The existing error to wrap.
// - message: Additional context to describe the error.
//
// Returns a new error that includes both the wrapped error and the new message.
// If 'err' is nil, Wrap returns nil.
func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("+[%s]\n%w\n-> %s", message, err, getCaller(2))
}
