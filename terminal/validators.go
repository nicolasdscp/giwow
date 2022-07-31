package terminal

import "github.com/nicolasdscp/giwow/errors"

func NotEmpty(inputName string) func(string) error {
	return func(input string) error {
		if input == "" {
			return errors.ErrStringEmpty(inputName)
		}
		return nil
	}
}
