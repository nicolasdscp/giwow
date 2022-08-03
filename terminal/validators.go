package terminal

import "github.com/nicolasdscp/giwow/internal/exception"

// NotEmpty returns a function that can be used as a validator.
// validates if the string is not empty, otherwise return an error.
func NotEmpty(inputName string) func(string) error {
	return func(input string) error {
		if input == "" {
			return exception.StringEmpty(inputName)
		}
		return nil
	}
}
