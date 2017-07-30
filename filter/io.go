package filter

import (
	"errors"
	"image"
)

//IOConfig contains all data about how go-filter will interact with a directory
type IOConfig struct {
	Directory string        `yaml:"directory"`
	Formats   []ImageFormat `yaml:"formats,flow"`
}

//extToFormat translates extensions to a image format
func extToFormat(ext string) (form ImageFormat, err error) {
	formats := map[string]ImageFormat{
		"png":  PNG,
		"gif":  GIF,
		"jpg":  JPEG,
		"jpeg": JPEG,
		"bmp":  BMP,
		"tiff": TIFF,
		"tif":  TIFF}

	// get formats
	format, formatExists := formats[ext]
	if !formatExists {
		return "", errors.New("format doesn't exist")
	}
	return format, nil
}

//formatToExt translates images formats to a extension
func formatToExt(form ImageFormat) (ext string, err error) {
	extensions := map[ImageFormat]string{
		PNG:  "png",
		GIF:  "gif",
		JPEG: "jpg",
		BMP:  "bmp",
		TIFF: "tiff"}

	// get extensions
	ext, extExists := extensions[form]
	if !extExists {
		return "", errors.New("extension doesn't exist")
	}
	return ext, nil
}

//TODO
func (io IOConfig) GetInput() image.Image {
	return nil
}
