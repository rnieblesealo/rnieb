package heicupload

import (
	"fmt"
	"gopkg.in/gographics/imagick.v3/imagick"
	"os"
	"path/filepath"
)

func HEICToPNG(imagePath string) {
	mw := imagick.MagickWand{}

	// check that image file exists
	_, err := os.Stat(imagePath)
	if err != nil {
		fmt.Printf("HEICToJPG: %s does not exist\n", imagePath)
		return
	}

	mw.ReadImage(imagePath)

	mw.SetFormat("PNG")
	mw.SetImageCompressionQuality(50)

	// write new png image
	// name will be same but extended as png, e.g. "image39.heic" --> "image39.png"
	newFilename := fmt.Sprintf("%s.png", filepath.Base(imagePath))
	mw.WriteImage(newFilename)

	// get rid of original heic image
	os.Remove(imagePath)

	fmt.Printf("Successfully converted %s to %s", imagePath, newFilename)
}
