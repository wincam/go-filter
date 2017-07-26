package settings

import (
	"path/filepath"
	"testing"
)

func TestReadFlags(t *testing.T) {
	configAddress, err := readFlags()
	if err != nil {
		t.Error("Readflag produced error when not expected")
	}
	if filepath.Base(configAddress) != "example.config.yml" {
		t.Error("Readflag produced " + configAddress + " when example.config.yml is expected")
	}
}

func TestReadConfig(t *testing.T) {
	readConfigTestFactory := func(configAddress string, expectedError bool) func(t *testing.T) {
		return func(t *testing.T) {
			err := readConfig(configAddress)

			// expected error check
			if (err != nil) != expectedError {
				if expectedError {
					t.Error("Error not detected when expected")
				} else {
					t.Error("Error detected when not expected")
				}
			}

		}
	}

	// run subtests
	configAddress, _ := filepath.Abs("../example.config.yml")
	t.Run("example config", readConfigTestFactory(configAddress, false))
	configAddress, _ = filepath.Abs("nonexistant.config.yml")
	t.Run("fake config", readConfigTestFactory(configAddress, true))

}
