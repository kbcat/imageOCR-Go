package main

import (
	"fmt"
	"log"

	"github.com/disintegration/imaging"
	"github.com/otiai10/gosseract/v2"
)

const (
	inputImg     = "level-final.png"
	processedImg = "processed.png"
)

func main() {
	err := processImage()
	if err != nil {
		log.Fatal(err)
	}

	err = OCRImage()
	if err != nil {
		log.Fatal(err)
	}
	return
}

func processImage() error {
	file, err := imaging.Open(inputImg, imaging.AutoOrientation(true))
	if err != nil {
		return err
	}
	greyImg := imaging.Grayscale(file)
	resamplerImg := imaging.Resize(greyImg, 0, 70, imaging.Gaussian)
	sharpenImg := imaging.Sharpen(resamplerImg, 0.5)
	constrastImg := imaging.AdjustContrast(sharpenImg, 100)
	// saturationImg := imaging.AdjustSaturation(constrastImg, 100)
	// brightnessImg := imaging.AdjustBrightness(constrastImg, -10)
	err = imaging.Save(constrastImg, processedImg, imaging.JPEGQuality(100))
	if err != nil {
		return err
	}

	return nil
}

func OCRImage() error {
	client := gosseract.NewClient()
	defer client.Close()
	err := client.SetImage(processedImg)
	if err != nil {
		return err
	}
	text, err := client.Text()
	if err != nil {
		return err
	}
	fmt.Println(text)
	return nil
}
