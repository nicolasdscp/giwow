package errors

import "fmt"

func NoConfigFile() error {
	return fmt.Errorf("no config file found")
}

func ErrStringEmpty(str string) error {
	return fmt.Errorf("%s cannot be empty", str)
}
