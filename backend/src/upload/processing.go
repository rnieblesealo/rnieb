package upload

import (
	"fmt"
	"gopkg.in/gographics/imagick.v3/imagick"
	"os"
	"path/filepath"
	"strings"
)

const (
	TARGET_WIDTH               = 500
	TARGET_COMPRESSION_QUALITY = 50
)

// Converts the image at the given path to PNG with TARGET_COMPRESSION_QUALITY applied
// Returns the converted PNG's filepath
// This filepath is the original filename but with a .png extension instead
func ConvertToPNG(imagePath string) (string, error) {
	// Create MagickWand
	/* One wand per edited image is expected, create + destroy MW has low overhead */

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	// Read image

	err := mw.ReadImage(imagePath)
	if err != nil {
		return "", fmt.Errorf("ConvertToPNG: Failed to read image: %s", err)
	}

	// Set conversion format

	mw.SetFormat("PNG")

	// Set compression quality

	mw.SetImageCompressionQuality(TARGET_COMPRESSION_QUALITY)

	// Write the result with the correct PNG extension

	newFilename := fmt.Sprintf("%s.png",
		strings.TrimSuffix(filepath.Base(imagePath), filepath.Ext(imagePath)))

	newFilepath := filepath.Join(filepath.Dir(imagePath), newFilename)

	err = mw.WriteImage(newFilepath)
	if err != nil {
		return "", fmt.Errorf("ConvertToPNG: Failed to write image: %s", err)
	}

	// Remove the original unconverted image

	os.Remove(imagePath)

	return newFilepath, nil
}

func ResizePNG(imagePath string) (string, error) {
	// Start wand

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	// Set format

	mw.SetFormat("PNG")

	// Read image

	err := mw.ReadImage(imagePath)
	if err != nil {
		return "", fmt.Errorf("ResizePNG: Failed to read image: %s", err)
	}

	// Ensure is PNG

	format := mw.GetImageFormat()
	if format != "PNG" {
		return "", fmt.Errorf(
			"ResizePNG: Only PNG images are allowed (received %s)\n",
			format)
	}

	// Perform resize (see README for math explanation)

	width := mw.GetImageWidth()
	height := mw.GetImageHeight()

	var convFactor float32 = TARGET_WIDTH / float32(width)

	mw.ResizeImage(
		uint(float32(width)*convFactor),
		uint(float32(height)*convFactor),
		imagick.FILTER_POINT)

	// Write result
	/* Since is same filename, no need to remove; it will be overwritten */

	err = mw.WriteImage(imagePath)
	if err != nil {
		return "", fmt.Errorf("ConvertToPNG: Failed to write image: %s", err)
	}

	return imagePath, nil
}
