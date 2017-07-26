package settings

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
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

//FilterConfig contains the rules that are applied when filtering
type FilterConfig struct {
	Input  DirectoryConfig `yaml:"input"`
	Output DirectoryConfig `yaml:"output"`
}

//DirectoryConfig contains all data about how go-filter will interact with a directory
type DirectoryConfig struct {
	Directory string        `yaml:"directory"`
	Formats   []ImageFormat `yaml:"formats,flow"`
}

//YamlConfig contains all settings contained in the config yml
type YamlConfig struct {
	FilterConfigs []FilterConfig `yaml:"config,flow"`
}

//Config is the configuration yaml instance for go-filter
var Config YamlConfig

//LoadConfig loads all arguments and the config yml
func LoadConfig() error {
	configAddress, err := readFlags()
	if err != nil {
		return err
	}

	return readConfig(configAddress)
}

//readFlags reads all flags and returns their values
func readFlags() (configAddress string, err error) {
	// process args
	configAddressPtr := flag.String("config", "example.config.yml", "The location of the config yml")
	flag.Parse()

	return filepath.Abs(*configAddressPtr)
}

// readConfig reads the config file at configAddress
func readConfig(configAddress string) (err error) {
	// read the config yml
	log.Println("opening config at " + configAddress)
	file, err := os.Open(configAddress)
	if err != nil {
		return err
	}
	defer file.Close()

	log.Println("reading config")
	configFile, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	// process the config yml
	err = yaml.Unmarshal(configFile, &Config)
	return err
}
