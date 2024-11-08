[![Go Reference](https://pkg.go.dev/badge/github.com/ovila98/ers.svg)](https://pkg.go.dev/github.com/ovila98/ers)
[![Email](https://img.shields.io/badge/contact-ovila.acolatse.dev@gmail.com-blue)](mailto:ovila.acolatse.dev@gmail.com)
[![LinkedIn](https://img.shields.io/badge/LinkedIn-Ovila_Acolatse-blue?style=flat&logo=linkedin)](https://www.linkedin.com/in/ovila-acolatse/)

# ers - Simple Error Handling for Go

`ers` is a custom error-handling package for Go that makes it easier to trace the origin of errors. It implements the error interface with custom Error and StackLine structs, providing automatic stack trace collection and error wrapping with context details.

## Features

- **Enhanced error creation**: Automatically includes the file name and line number of where the error was generated
- **Error wrapping**: Wrap errors with additional context while maintaining the complete stack trace
- **Unwrap support**: Full support for Go's error unwrapping conventions
- **Stack trace collection**: Automatic collection of file and line information at each wrapping point
- **Formatted messages**: Support for fmt-style formatting in error messages and context
- **100% Test Coverage**: Comprehensive test suite covering all functionality and edge cases

## Installation v1.2.0

To install the package, use the following command:

```bash
go get github.com/ovila98/ers
```

Then, import it in your project:

```go
import "github.com/ovila98/ers"
```

## Usage

Here's how to use `ers` to improve error handling in your Go projects.

### 1. Creating New Errors

Create new errors using the `New` function with optional formatting:

```go
package main

import (
    "fmt"
    "github.com/ovila98/ers"
)

func main() {
    // Simple error
    err1 := ers.New("Something went wrong")

    // Formatted error
    count := 5
    err2 := ers.New("Processing failed after %d attempts", count)

    fmt.Println(err1)
    fmt.Println(err2)
}
```

### 2. Wrapping Errors with Context

Add context to errors using either `Wrap` or `Wrapf`:

```go
func fetchData() error {
    return ers.New("Failed to fetch data")
}

func processData() error {
    err := fetchData()
    if err != nil {
        // Using Wrap with multiple context details
        return ers.Wrap(err, "During data processing", "Retry attempt 3")

        // Or using Wrapf for formatted context
        // return ers.Wrapf(err, "Processing failed on attempt %d", 3)
    }
    return nil
}
```

Output example:

```
Failed to fetch data

~ additional context (most recent last) ~
+[During data processing]
+[Retry attempt 3]

• stack trace (most recent call last) •
-> (/path/to/file.go:5)
-> (/path/to/file.go:10)
```

## API Reference

### `New(fmessage string, formatTags ...any) error`

- Creates a new error with formatted message and current stack location
- Supports fmt-style formatting through formatTags
- Returns an Error struct implementing the error interface

### `Wrap(err error, details ...string) error`

- Wraps an existing error with multiple context details
- Automatically adds the current stack location
- Returns an Error struct that can be unwrapped

### `Wrapf(err error, fmessage string, formatTags ...any) error`

- Wraps an error with a formatted context message
- Supports fmt-style formatting for the context
- Returns an Error struct that can be unwrapped

### Error Types

#### `Error` struct

- Implements the error interface
- Contains the original error
- Maintains stack trace information
- Stores context details
- Supports unwrapping via `Unwrap()`
- Provides access to stack and context information

#### `StackLine` struct

- Stores file and line information for stack traces
- Provides string formatting for stack entries
- Offers access to individual file and line details

## Why Use `ers`?

1. **Modern Error Handling**: Full support for Go's error wrapping conventions
2. **Automatic Stack Traces**: Built-in collection of stack information
3. **Multiple Context Details**: Add multiple pieces of context when wrapping errors
4. **Formatted Messages**: Support for fmt-style formatting in both errors and context
5. **Clean Implementation**: Leverages Go's error interface effectively
6. **100% Test Coverage**: Thoroughly tested with comprehensive test suite

## Contributions

Contributions are welcome! Please check out the [CONTRIBUTING.md](CONTRIBUTING.md) file for more information.

## License

This project is licensed under the [Apache 2.0 License](LICENSE).
