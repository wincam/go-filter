package filter

import (
	"os"
	"testing"
)

func TestIsDirectory(t *testing.T) {
	isDirectoryTestFactory := func(dirAddress string, dirExpected bool) func(t *testing.T) {
		return func(t *testing.T) {
			isDir := isDirectory(dirAddress)

			// expected error check
			if isDir != dirExpected {
				if dirExpected {
					t.Error("Unexpected dir at " + dirAddress)
				} else {
					t.Error("No dir at " + dirAddress)
				}
			}

		}
	}

	dir, err := os.Getwd()
	if err != nil {
		t.Error("Cannot get PWD")
	}
	// run tests
	t.Run("PWD test", isDirectoryTestFactory(dir, true))
	t.Run("Fake dir test", isDirectoryTestFactory("fake dir", false))
}
