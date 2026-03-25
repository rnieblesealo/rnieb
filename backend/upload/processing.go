package upload

import (
	"fmt"
	"github.com/u2takey/ffmpeg-go"
	"gopkg.in/gographics/imagick.v3/imagick"
	"os"
	"path/filepath"
	"strings"
)

const (
	TARGET_IMAGE_WIDTH               = 500
	TARGET_IMAGE_COMPRESSION_QUALITY = 50
	TARGET_VIDEO_WIDTH               = 400
	TARGET_VIDEO_CRF                 = 28 // This is basically compression rate
)

// All functions return the filepath of the newly processed file
// If something failed or no new file was created, the original filepath is returned

// Converts the image at the given path to PNG with TARGET_COMPRESSION_QUALITY applied
func ConvertToPNG(imagePath string) (string, error) {
	// Create MagickWand
	/* One wand per edited image is expected, create + destroy MW has low overhead */

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	// Read image

	err := mw.ReadImage(imagePath)
	if err != nil {
		return imagePath, fmt.Errorf("ConvertToPNG: Failed to read image: %s", err)
	}

	// Set conversion format

	mw.SetFormat("PNG")

	// Set compression quality

	mw.SetImageCompressionQuality(TARGET_IMAGE_COMPRESSION_QUALITY)

	// Write the result with the correct PNG extension

	newFilename := fmt.Sprintf("%s.png",
		strings.TrimSuffix(filepath.Base(imagePath), filepath.Ext(imagePath)))
	newFilepath := filepath.Join(filepath.Dir(imagePath), newFilename)

	err = mw.WriteImage(newFilepath)
	if err != nil {
		return imagePath, fmt.Errorf("ConvertToPNG: Failed to write image: %s", err)
	}

	// Remove the original unconverted image

	os.Remove(imagePath)

	return newFilepath, nil
}

// Resizes PNG to TARGET_IMAGE_WIDTH
func ResizePNG(imagePath string) (string, error) {
	// Start wand

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	// Set format

	mw.SetFormat("PNG")

	// Read image

	err := mw.ReadImage(imagePath)
	if err != nil {
		return imagePath, fmt.Errorf("ResizePNG: Failed to read image: %s", err)
	}

	// Ensure is PNG

	format := mw.GetImageFormat()
	if format != "PNG" {
		return imagePath, fmt.Errorf(
			"ResizePNG: Only PNG images are allowed (received %s)\n",
			format)
	}

	// Perform resize (see README for math explanation)

	width := mw.GetImageWidth()
	height := mw.GetImageHeight()

	var convFactor float32 = TARGET_IMAGE_WIDTH / float32(width)

	mw.ResizeImage(
		uint(float32(width)*convFactor),
		uint(float32(height)*convFactor),
		imagick.FILTER_POINT)

	// Write result
	/* Since is same filename, no need to remove; it will be overwritten */

	err = mw.WriteImage(imagePath)
	if err != nil {
		return imagePath, fmt.Errorf("ConvertToPNG: Failed to write image: %s", err)
	}

	return imagePath, nil
}

// Proportionally scale video to TARGET_VIDEO_WIDTH
// Then reencode as H264 with new CBF ( TARGET_VIDEO_CBF )
func ResizeAndCompressVideo(videoPath string) (string, error) {
	// Get new filename and path ( just changes extension )

	newFilename := fmt.Sprintf("%s.mp4", // Change ext to tell ffmpeg to reformat
		strings.TrimSuffix(filepath.Base(videoPath), filepath.Ext(videoPath)))
	newFilepath := filepath.Join(filepath.Dir(videoPath), newFilename)

	// Run ffmpeg

	err := ffmpeg_go.Input(videoPath).Filter("scale",
		ffmpeg_go.Args{
			fmt.Sprintf("w=%d:h=-2", TARGET_VIDEO_WIDTH),
			// -2 makes sure resulting height % 2 == 0 which h264 requires
		}).Output(newFilepath, ffmpeg_go.KwArgs{ // Change ext here for new format
		"vcodec": "libx264",                           // Select libx264 ( H264 codec )
		"crf":    fmt.Sprintf("%d", TARGET_VIDEO_CRF), // Set constant rate factor to 28
	}).Run()

	// NOTE: I tried H265 which is supposedly newer better but that shit dont work
	// ( QuickTime couldnt open the output )

	if err != nil {
		// ffmpeg will always create the file, if it fails it just leaves it as is
		// We must get rid of this junk file
		os.Remove(newFilepath)

		// Leave the original video since we still may wanna work with it

		return videoPath, fmt.Errorf("ResizeAndCompressVideo: %s", err)
	}

	os.Remove(videoPath) // Get rid of original video since we processed it

	return newFilepath, nil
}
