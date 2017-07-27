package settings

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/wincam/go-filter/filter"
	"gopkg.in/yaml.v2"
)

//ConfigList contains all settings contained in the config yml
type ConfigList struct {
	FilterConfigs []filter.Config `yaml:"config,flow"`
}

//Config is the configuration yaml instance for go-filter
var Config ConfigList

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
