package filter

import (
	"io/ioutil"
	"log"
	"os"
	"path"
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

type config interface {
	processInput() error
	isDir() bool
}

//processDirectory runs a filter on a directory and all it's children
func (filter Config) processInput() error {
	if filter.isDir() {
		//get all files in directory
		files, err := ioutil.ReadDir(filter.Input.Directory)
		if err != nil {
			return err
		}

		// process children of input
		for _, file := range files {
			// point to children
			newFilter := filter
			cd(file.Name(), &newFilter)
			err = newFilter.processInput()
			if err != nil {
				log.Print(err)
			}
		}
	} else {
		//TODO process
	}

	return nil
}

//cd changes the filter directory to dirName
func cd(dirName string, filter *Config) {
	filter.Output.Directory = path.Join(filter.Output.Directory, dirName)
	filter.Input.Directory = path.Join(filter.Input.Directory, dirName)
}

func (filter Config) isDir() bool {
	isDirectory := func(dir string) bool {
		// check if directory exists and is directory
		dirInfo, err := os.Stat(dir)
		if err != nil {
			return false
		}
		return dirInfo.IsDir()
	}
	return isDirectory(filter.Input.Directory) && isDirectory(filter.Output.Directory)
}

//Config contains the rules that are applied when filtering
type Config struct {
	Input  IOConfig `yaml:"input"`
	Output IOConfig `yaml:"output"`
}
