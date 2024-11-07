// Package ers provides enhanced error handling with stack traces and context information
package ers

import (
	"fmt"
	"runtime"
	"strings"
)

// StackLine represents a single stack trace entry containing the file path
// and line number where an error occurred or was wrapped.
type StackLine struct {
	// file stores the source code file path where the error occurred
	file string
	// line stores the line number in the source file
	line int
}

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
			file: "unknown",
			line: 0,
		}
	}
	return StackLine{
		file: file,
		line: line,
	}
}

// NewStackLine creates a new StackLine capturing the current file and line number.
// It skips 2 stack frames to get the actual caller's location.
func NewStackLine() StackLine {
	return getCaller(2)
}

// String formats the StackLine into a readable string showing file and line number
// in the format "(file:line)".
func (s StackLine) String() string {
	return fmt.Sprintf("(%s:%d)", s.file, s.line)
}

// File returns the file path stored in the StackLine.
func (s StackLine) File() string {
	return s.file
}

// Line returns the line number stored in the StackLine.
func (s StackLine) Line() int {
	return s.line
}

// Error implements the error interface and provides enhanced error handling
// capabilities. It maintains the original error, a stack trace of StackLine
// entries, and additional context information provided during error wrapping.
type Error struct {
	// error holds the original error being wrapped
	error
	// stackTrace maintains the stack of StackLine entries showing error progression
	stackTrace []StackLine
	// contexts stores additional context messages added during error wrapping
	contexts []string
	// Maybe add error type and error code later
}

// Error returns a formatted string containing the original error message,
// additional context information if any exists, and the complete stack trace.
//
// The format follows:
//
// original error
//
// ~ additional context (if any) ~
//
// • stack trace •
func (e *Error) Error() string {
	stackTrace := fmt.Sprintf("• stack trace (most recent call last) •\n%s", e.StackTrace())
	if len(e.contexts) > 0 {
		contexts := fmt.Sprintf("~ additional context (most recent last) ~\n%s", e.Context())
		return fmt.Sprintf("%s\n\n%s\n%s", e.error.Error(), contexts, stackTrace)
	}
	out := fmt.Sprintf("%s\n\n%s", e.error.Error(), stackTrace)
	return out
}

// Unwrap returns the underlying error, enabling compatibility with Go's
// error unwrapping conventions and errors.Is/As functionality.
func (e *Error) Unwrap() error {
	return e.error
}

// StackTrace formats the complete stack trace into a readable string,
// with each line prefixed by an arrow "->".
func (e *Error) StackTrace() string {
	trace := ""
	for _, line := range e.stackTrace {
		trace += fmt.Sprintf("-> %s\n", line.String())
	}
	return strings.TrimSuffix(trace, "\n")
}

// Stack returns the complete stack trace as a slice of StackLine entries,
// allowing programmatic access to the trace information.
func (e *Error) Stack() []StackLine {
	return e.stackTrace
}

// AddContext appends one or more context messages to the error's additional information.
// Messages are stored in order and displayed most recent last.
func (e *Error) AddContext(contexts ...string) {
	e.contexts = append(e.contexts, contexts...)
}

// Context formats all context messages into a readable string,
// with each message prefixed by "+[" and suffixed with "]".
func (e *Error) Context() string {
	info := ""
	for _, line := range e.contexts {
		info += fmt.Sprintf("+[%s]\n", line)
	}
	return strings.TrimSuffix(info, "\n")
}

// Contexts returns the slice of additional context messages associated with the error,
// allowing programmatic access to the context information.
func (e *Error) Contexts() []string {
	return e.contexts
}
