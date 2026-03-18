package heicupload

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

// returns the new filepath
func ConvertToPNG(imagePath string) (string, error) {
	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	// read image
	err := mw.ReadImage(imagePath)
	if err != nil {
		return "", err
	}

	mw.SetFormat("PNG")
	mw.SetImageCompressionQuality(TARGET_COMPRESSION_QUALITY)

	// resize such that width is 500 but aspect is maintained
	width := mw.GetImageWidth()
	height := mw.GetImageHeight()

	convFactor := uint(TARGET_WIDTH / width)

	mw.ResizeImage(uint(width*convFactor), uint(height*convFactor), imagick.FILTER_POINT)

	// write new png image
	// name will be same but extended as png, e.g. "image39.heic" --> "image39.png"
	newFilename := fmt.Sprintf("%s.png",
		strings.TrimSuffix(filepath.Base(imagePath), filepath.Ext(imagePath)))
	newFilepath := filepath.Join(filepath.Dir(imagePath), newFilename)

	err = mw.WriteImage(newFilepath)
	if err != nil {
		return "", err
	}

	// get rid of original heic image
	os.Remove(imagePath)

	fmt.Printf("Successfully converted %s to %s\n", imagePath, newFilepath)

	return newFilepath, nil
}

func ResizePNG(imagePath string) (string, error) {
	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	mw.SetFormat("PNG")

	// read image
	err := mw.ReadImage(imagePath)
	if err != nil {
		return "", err
	}

	// only allow png
	format := mw.GetImageFormat()
	if format != "PNG" {
		return "", fmt.Errorf("Only PNG images are allowed (received %s)\n", format)
	}

	// resize
	width := mw.GetImageWidth()
	height := mw.GetImageHeight()

	var convFactor float32 = TARGET_WIDTH / float32(width)

	mw.ResizeImage(
		uint(float32(width)*convFactor),
		uint(float32(height)*convFactor),
		imagick.FILTER_POINT)

	// write
	err = mw.WriteImage(imagePath)
	if err != nil {
		return "", err
	}

	fmt.Printf("Resized %s\n", imagePath)

	return imagePath, nil
}
