package ers

import (
	"fmt"
	"strings"
)

// StackLine represents a single stack trace entry containing the file path
// and line number where an error occurred or was wrapped.
type StackLine struct {
	File string
	Line int
}

// String formats the StackLine into a readable string showing file and line number.
func (s StackLine) String() string {
	return fmt.Sprintf("(%s:%d)", s.File, s.Line)
}

// Error implements the error interface and provides enhanced error handling
// capabilities. It maintains the original error, a stack trace of StackLine
// entries, and additional context information provided during error wrapping.
type Error struct {
	error
	stackTrace     []StackLine
	additionalInfo []string
}

// Error returns a formatted string containing the original error message,
// additional context information if any exists, and the complete stack trace.
func (e *Error) Error() string {
	stackTrace := fmt.Sprintf("• stack trace (most recent call last) •\n%s", e.StackTrace())
	if len(e.additionalInfo) > 0 {
		additionalInfo := fmt.Sprintf("~ additional context (most recent last) ~\n%s", e.AdditionalInfo())
		return fmt.Sprintf("%s\n\n%s\n%s", e.error.Error(), additionalInfo, stackTrace)
	}
	out := fmt.Sprintf("%s\n\n%s", e.error.Error(), stackTrace)
	return out
}

// Unwrap returns the underlying error, enabling compatibility with Go's
// error unwrapping conventions.
func (e *Error) Unwrap() error {
	return e.error
}

// StackTrace formats the complete stack trace into a readable string,
// with each line prefixed by an arrow.
func (e *Error) StackTrace() string {
	trace := ""
	for _, line := range e.stackTrace {
		trace += fmt.Sprintf("-> %s\n", line.String())
	}
	return strings.TrimSuffix(trace, "\n")
}

// AddInfo appends a new context message to the error's additional information.
func (e *Error) AddInfo(info ...string) {
	e.additionalInfo = append(e.additionalInfo, info...)
}

// AdditionalInfo formats all context messages into a readable string,
// with each message prefixed by a plus sign and enclosed in brackets.
func (e *Error) AdditionalInfo() string {
	info := ""
	for _, line := range e.additionalInfo {
		info += fmt.Sprintf("+[%s]\n", line)
	}
	return strings.TrimSuffix(info, "\n")
}
