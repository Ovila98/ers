package ers

import (
	"fmt"
	"runtime"
)

func getCaller(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "(unable to get line  number)"
	}
	return fmt.Sprintf("(%s:%d)", file, line)

}

func New(message string) error {
	return fmt.Errorf("%s\n-> %s", message, getCaller(2))
}

func Trace(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%w\n-> %s", err, getCaller(2))
}

func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("+[%s]\n%w\n-> %s", message, err, getCaller(2))
}
