package heicupload

import (
	"fmt"
	"gopkg.in/gographics/imagick.v3/imagick"
	"os"
)

func HEICToJPG(filepath string) {
	mw := imagick.MagickWand{}

	// check that image file exists
	_, err := os.Stat(filepath)
	if err != nil {
		fmt.Printf("HEICToJPG: %s does not exist\n", filepath)
		return
	}

	mw.SetFormat("PNG")
	mw.SetImageCompressionQuality(50)
	mw.WriteImage("output.png")
	mw.ReadImage(filepath)
}
