package utils

import (
	"fmt"
	"strings"
)

// ErrorsWithStack returns a string contains error messages in the stack with its stack trace levels for given error
func ErrorsWithStack(err error) string {
	res := fmt.Sprintf("%+v\n", err)
	return res
}

// ErrorsWithoutStack just returns error messages without its callstack
func ErrorsWithoutStack(err error, format bool) string {
	res := fmt.Sprintf("%v\n", err)

	if format {
		var errStr string
		items := strings.Split(res, ":")
		for _, item := range items {
			errStr += fmt.Sprintf("%s\n", strings.TrimSpace(item))
		}
		return errStr
	}

	return res
}
