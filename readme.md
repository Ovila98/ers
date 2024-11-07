# ers - Simple Error Handling for Go

`ers` is a custom error-handling package for Go that makes it easier to trace the origin of errors. It implements the error interface with custom Error and StackLine structs, providing automatic stack trace collection and error wrapping capabilities.

## Features

- **Enhanced error creation**: Automatically includes the file name and line number of where the error was generated
- **Error wrapping**: Wrap errors with additional context while maintaining the complete stack trace
- **Unwrap support**: Full support for Go's error unwrapping conventions
- **Stack trace collection**: Automatic collection of file and line information at each wrapping point

## Installation v1.0.0

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

Create a new error using the `New` function, which automatically includes the file and line number where the error was generated.

```go
package main

import (
    "fmt"
    "github.com/ovila98/ers"
)

func main() {
    err := ers.New("Something went wrong")
    fmt.Println(err)
}
```

Output:

```
Something went wrong

• stack trace (most recent call last) •
-> (/path/to/file.go:10)
```

### 2. Wrapping Errors with Context

Add context to an error using the `Wrap` function. It preserves the original error and stack trace while adding new context and location information.

```go
func fetchData() error {
    return ers.New("Failed to fetch data")
}

func processData() error {
    err := fetchData()
    if err != nil {
        return ers.Wrap(err, "During data processing", "Additional detail")
    }
    return nil
}

func main() {
    err := processData()
    fmt.Println(err)
}
```

Output:

```
Failed to fetch data

~ additional context (most recent last) ~
+[During data processing]
+[Additional detail]
• stack trace (most recent call last) •
-> (/path/to/file.go:5)
-> (/path/to/file.go:10)
```

## API Reference

### `New(message string) error`

- Creates a new error with the given message and current stack location
- Returns an `Error` struct implementing the error interface

### `Wrap(err error, details ...string) error`

- Wraps an existing error with additional context details
- Automatically adds the current stack location
- Supports multiple detail strings
- Returns an `Error` struct that can be unwrapped

### Error Types

#### `Error` struct

- Implements the error interface
- Contains the original error
- Maintains stack trace information
- Stores additional context details
- Supports unwrapping via `Unwrap()`

#### `StackLine` struct

- Stores file and line information for stack traces
- Used internally to track error locations

## Why Use `ers`?

1. **Modern Error Handling**: Full support for Go's error wrapping conventions
2. **Automatic Stack Traces**: Built-in collection of stack information
3. **Multiple Context Details**: Add multiple pieces of context when wrapping errors
4. **Clean Implementation**: Leverages Go's error interface effectively

## Contributions

Contributions are welcome! Please check out the [CONTRIBUTING.md](CONTRIBUTING.md) file for more information.

## License

This project is licensed under the [Apache 2.0 License](LICENSE).
