package filter

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/wincam/go-filter/settings"
)

//RunFilter takes a filter config and runs that filter
func RunFilter(filter settings.FilterConfig) error {
	if !isDirectory(filter.Input.Directory) {
		return errors.New(filter.Input.Directory + " is not a directory.")
	}

	err := processDirectory(filter)
	if err != nil {
		return err
	}

	return nil
}

//isDirectory determines if
func isDirectory(dir string) bool {
	// check if directory exists and is directory
	dirInfo, err := os.Stat(dir)
	if err != nil {
		return false
	}
	return dirInfo.IsDir()
}

//processDirectory runs a filter on a directory and all it's children
func processDirectory(filter settings.FilterConfig) error {
	//get all files in directory
	files, err := ioutil.ReadDir(filter.Input.Directory)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Println(file.Name() + "/")
			newFilter := filter
			newFilter.Output.Directory = path.Join(newFilter.Output.Directory, file.Name())
			newFilter.Input.Directory = path.Join(newFilter.Input.Directory, file.Name())
			processDirectory(newFilter)
		}
		fmt.Println(file.Name())
	}

	return nil
}
