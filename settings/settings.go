package settings

import (
	"flag"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type DirectoryConfig struct {
	Directory string   `yaml:"directory"`
	Formats   []string `yaml:"formats,flow"`
}

//YamlConfig contains all settings contained in the config yml
type YamlConfig struct {
	Input  DirectoryConfig `yaml:"input"`
	Output DirectoryConfig `yaml:"output"`
}

//Config is the configuration yaml instance for go-filter
var Config YamlConfig

//LoadConfig loads all arguments and the config yml
func LoadConfig() error {
	// process args
	configAddressPtr := flag.String("config", "~/.config/go-filter-config", "The location of the config yml")
	flag.Parse()

	// read the config yml
	file, err := os.Open(*configAddressPtr)
	if err != nil {
		return err
	}
	defer file.Close()

	configText, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	// process the config yml
	err = yaml.Unmarshal(configText, &Config)
	if err != nil {
		return err
	}
	return nil
}
