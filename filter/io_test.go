package filter

import (
	"testing"
)

//TestExtToFormat tests the extToFormat function
func TestExtToFormat(t *testing.T) {
	extToFormatTestFactory := func(ext string, expectedFormat ImageFormat, errorExpected bool) func(t *testing.T) {
		return func(t *testing.T) {
			format, err := extToFormat(ext)
			if err != nil && !errorExpected {
				t.Error("error thrown when not expected")
			} else if err == nil && errorExpected {
				t.Error("error not thrown when expected")
			} else if expectedFormat != format {
				t.Error(expectedFormat + " does not match " + format)
			}
		}
	}

	// run tests
	t.Run("fake extension", extToFormatTestFactory(".fake", "", true))
	t.Run(".tiff", extToFormatTestFactory("tiff", TIFF, false))
	t.Run(".gif", extToFormatTestFactory("gif", GIF, false))
	t.Run(".jpg", extToFormatTestFactory("jpg", JPEG, false))
}

//TestFormatToExt tests the formatToExt function
func TestFormatToExt(t *testing.T) {
	formatToExtTestFactory := func(format ImageFormat, expectedExt string, errorExpected bool) func(t *testing.T) {
		return func(t *testing.T) {
			ext, err := formatToExt(format)
			if err != nil && !errorExpected {
				t.Error("error thrown when not expected")
			} else if err == nil && errorExpected {
				t.Error("error not thrown when expected")
			} else if expectedExt != ext {
				t.Error(expectedExt + " does not match " + ext)
			}
		}
	}

	t.Run("fake format", formatToExtTestFactory("FAKE", "", true))
	t.Run(".tiff", formatToExtTestFactory(TIFF, "tiff", false))
	t.Run(".gif", formatToExtTestFactory(GIF, "gif", false))
	t.Run(".jpg", formatToExtTestFactory(JPEG, "jpg", false))
}
