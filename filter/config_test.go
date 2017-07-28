package filter

import (
	"os"
	"testing"
)

//TestIsDir tests the isDir fuction
func TestIsDir(t *testing.T) {
	isDirectoryTestFactory := func(dirAddress string, dirExpected bool) func(t *testing.T) {
		return func(t *testing.T) {
			testConfig := Config{Input: DirectoryConfig{Directory: dirAddress}, Output: DirectoryConfig{Directory: dirAddress}}
			isDir := testConfig.isDir()

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

//Test for the Cd function
func TestCd(t *testing.T) {
	testConfig := Config{Input: DirectoryConfig{Directory: "/testInputDir"}, Output: DirectoryConfig{Directory: "/testOutputDir"}}

	testConfig.cd("testSubDir")
	if testConfig.Input.Directory != "/testInputDir/testSubDir" {
		t.Error("Input dir not correct: " + testConfig.Input.Directory)
	}

	if testConfig.Output.Directory != "/testOutputDir/testSubDir" {
		t.Error("Output dir not correct: " + testConfig.Output.Directory)
	}
}
