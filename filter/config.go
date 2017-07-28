package filter

import (
	"io/ioutil"
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
	processDirectory() error
	cd(dirName string)
	isDir() bool
}

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
				newFilter.cd(file.Name())
				newFilter.processDirectory()
			}
			fmt.Println(file.Name())
		}*/

	return nil
}

//cd changes the filter directory to dirName
func (filter Config) cd(dirName string) {
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
