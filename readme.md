# ers - Enhanced Error Handling for Go

`ers` is a custom error-handling package for Go that makes it easier to trace the origin of errors. It allows you to create errors, trace them through the call stack, and wrap existing errors with additional context, while automatically appending the source file and line number to the error messages.

## Features

- **Enhanced error creation**: Automatically includes the file name and line number of where the error was generated.
- **Error tracing**: Easily trace errors through the call stack by adding location information at each point where the error is propagated.
- **Error wrapping**: Wrap errors with additional context without losing the original error details.

## Installation

To install the package, use the following command:

```bash
go get github.com/ovila98/ers
```

Then, import it in your project:

```go
import "github.com/ovila98/ers"
```

## Usage

Here’s how to use `ers` to improve error handling in your Go projects.

### 1. Creating New Errors

You can create a new error using the `New` function, which automatically includes the file and line number where the error was generated.

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
-> (/path/to/file.go:10)
```

### 2. Tracing Errors

When an error is returned from a function and passed up the call stack, you can add trace information using the `Trace` function.

```go
func innerFunc() error {
    return ers.New("Initial error")
}

func outerFunc() error {
    err := innerFunc()
    if err != nil {
        return ers.Trace(err)
    }
    return nil
}

func main() {
    err := outerFunc()
    fmt.Println(err)
}
```

Output:

```
Initial error
-> (/path/to/file.go:5)
-> (/path/to/file.go:10)
```

### 3. Wrapping Errors with Context

You can add more context to an error using the `Wrap` function. It keeps the original error message and stack trace, while adding a new message.

```go
func fetchData() error {
    return ers.New("Failed to fetch data")
}

func processData() error {
    err := fetchData()
    if err != nil {
        return ers.Wrap(err, "While processing data")
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
+[While processing data]
Failed to fetch data
-> (/path/to/file.go:5)
-> (/path/to/file.go:10)
```

## Method Chaining Example

You can chain multiple calls to `Trace` and `Wrap` to create a comprehensive error chain:

```go
func funcA() error {
    return ers.New("Error in funcA")
}

func funcB() error {
    err := funcA()
    if err != nil {
        return ers.Trace(err)
    }
    return nil
}

func funcC() error {
    err := funcB()
    if err != nil {
        return ers.Wrap(err, "Additional context in funcC")
    }
    return nil
}

func main() {
    err := funcC()
    fmt.Println(err)
}
```

Output:

```
+[Additional context in funcC]
Error in funcA
-> (/path/to/file.go:5)
-> (/path/to/file.go:10)
-> (/path/to/file.go:15)
```

## API Reference

### `New(message string) error`

- **Description**: Creates a new error with the given message and appends the file and line number of the caller.
- **Parameters**: `message` - The error message to include.
- **Returns**: An `error` that includes the message and the caller’s file and line number.

### `Trace(err error) error`

- **Description**: Appends the caller’s file and line number to an existing error.
- **Parameters**: `err` - The existing error to trace.
- **Returns**: An `error` that wraps the original error and adds location context.

### `Wrap(err error, message string) error`

- **Description**: Wraps an existing error with a new message and appends the caller’s file and line number.
- **Parameters**:
  - `err` - The existing error to wrap.
  - `message` - Additional context to include.
- **Returns**: An `error` that wraps the original error with additional context and location information.

## Why Use `ers`?

1. **Easier Debugging**: Trace errors back to their source with minimal effort. The package automatically includes file and line numbers, so you don't have to manually add that information.
2. **Cleaner Error Handling**: Instead of cluttering your code with manual stack tracing, `ers` does it for you.
3. **Clear Contextual Errors**: Wrap errors with additional messages that provide more insight into what went wrong, while still preserving the original error.

## Contributions

Contributions are welcome! If you want to improve this package, please check out the [CONTRIBUTING.md](CONTRIBUTING.md) file for more information.

## License

This project is licensed under the [Apache 2.0 License](LICENSE).
