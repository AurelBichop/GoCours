package filter

import (
	"image/jpeg"
	"os"

	"github.com/disintegration/imaging"
)

type Filter interface {
	Process(srcPath, dstPath string) error
}

type Grayscale struct{}

func (g Grayscale) Process(srcPath, dstPath string) error {

	src, err := imaging.Open(srcPath)
	if err != nil {
		return err
	}

	//Create a grayscale version of the image
	dst := imaging.Grayscale(src)

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// Save the resulting image as JPEG.
	opts := jpeg.Options{Quality: 90}
	return jpeg.Encode(dstFile, dst, &opts)
}

type Blur struct{}

func (b Blur) Process(srcPath, dstPath string) error {

	src, err := imaging.Open(srcPath)
	if err != nil {
		return err
	}

	//Create a grayscale version of the image
	dst := imaging.Blur(src, 3.5)

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// Save the resulting image as JPEG.
	opts := jpeg.Options{Quality: 90}
	return jpeg.Encode(dstFile, dst, &opts)
}
