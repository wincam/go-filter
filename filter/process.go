package filter

import (
	"errors"
	"io/ioutil"
	"os"
)

//ImageFormat defines types of images
type ImageFormat string

// supported image format identifiers
const (
	PNG  ImageFormat = "PNG"
	GIF  ImageFormat = "GIF"
	JPEG ImageFormat = "JPEG"
	BMP  ImageFormat = "BMP"
	TIFF ImageFormat = "TIFF"
)

//Config contains the rules that are applied when filtering
type Config struct {
	Input  DirectoryConfig `yaml:"input"`
	Output DirectoryConfig `yaml:"output"`
}

//DirectoryConfig contains all data about how go-filter will interact with a directory
type DirectoryConfig struct {
	Directory string        `yaml:"directory"`
	Formats   []ImageFormat `yaml:"formats,flow"`
}

//RunFilter takes a filter config and runs that filter
func RunFilter(filter Config) error {
	if !isDirectory(filter.Input.Directory) {
		return errors.New(filter.Input.Directory + " is not a directory.")
	}

	err := filter.processDirectory()
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
func (filter Config) processDirectory() error {
	//get all files in directory
	_, err := ioutil.ReadDir(filter.Input.Directory)
	if err != nil {
		return err
	}

	/*
		for _, file := range files {
			if file.IsDir() {
				fmt.Println(file.Name() + "/")
				newFilter := filter
				newFilter.Output.Directory = path.Join(newFilter.Output.Directory, file.Name())
				newFilter.Input.Directory = path.Join(newFilter.Input.Directory, file.Name())
				newFilter.processDirectory()
			}
			fmt.Println(file.Name())
		}*/

	return nil
}
