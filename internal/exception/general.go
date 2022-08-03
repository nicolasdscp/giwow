package exception

import (
	"github.com/pkg/errors"
)

func NoConfigFile() error {
	return errors.New("no config file found")
}

func StringEmpty(strName string) error {
	return errors.Errorf("%s cannot be empty", strName)
}
