package filter

import (
	"errors"
)

//RunFilter takes a filter config and runs that filter
func RunFilter(filter config) error {
	if !filter.isDir() {
		return errors.New("config does not point a directory")
	}

	err := filter.processInput()
	if err != nil {
		return err
	}

	return nil
}
