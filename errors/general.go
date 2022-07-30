package errors

import "fmt"

func NoConfigFile() error {
	return fmt.Errorf("no config file found")
}
