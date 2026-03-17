package heicupload

import (
	"fmt"
	"gopkg.in/gographics/imagick.v3/imagick"
	"os"
	"path/filepath"
	"strings"
)

func HEICToPNG(imagePath string) {
	mw := imagick.NewMagickWand()

	// read image
	err := mw.ReadImage(imagePath)
	if err != nil {
		fmt.Printf("HEICToJPG: %s", err.Error())
		return
	}

	mw.SetFormat("PNG")
	mw.SetImageCompressionQuality(50)

	/*
	   -resize 500x500 \
	   -extent 1:1 \
	   -gravity Center \
	   -quality 50 \
	*/

	mw.ResizeImage(500, 500, imagick.FILTER_LANCZOS)

	// write new png image
	// name will be same but extended as png, e.g. "image39.heic" --> "image39.png"
	newFilename := fmt.Sprintf("%s.png", strings.TrimSuffix(filepath.Base(imagePath), filepath.Ext(imagePath)))
	newPath := filepath.Join(filepath.Dir(imagePath), newFilename)

	err = mw.WriteImage(newPath)
	if err != nil {
		fmt.Printf("HEICToJPG: %s", err.Error())
		return
	}

	// get rid of original heic image
	os.Remove(imagePath)

	fmt.Printf("Successfully converted %s to %s", imagePath, newPath)
}
