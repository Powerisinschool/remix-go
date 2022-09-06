package converter

import (
	"errors"
	"fmt"
	"image"
	"image-converter/match"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"strings"

	ewebp "github.com/chai2010/webp"
	"golang.org/x/image/webp"
)

var matchers = []string{"*.png", "*.jpeg", "*.webp", "*.gif"}

// Get the bi-dimensional pixel array
func getPixels(file io.Reader) ([][]Pixel, error) {
	img, _, err := image.Decode(file)

	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][]Pixel
	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			row = append(row, _rgbaToPixel(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}

	return pixels, nil
}

// img.At(x, y).RGBA() returns four uint32 values; we want a Pixel
func _rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{uint8(r / 257), uint8(g / 257), uint8(b / 257), uint8(a / 257)}
}

// Pixel struct example
type Pixel struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func Convert(path string, newPath string) (image.Image, error) {
	// Handle Errors
	openers := []string{"--open", "--open-server"}
	if match.IsMatching(newPath, openers) {
		newPath = ""
	}

	matchers := []string{"*.png", "*.jpeg", "*.webp", "*.gif"}
	if (!match.IsMatching(path, matchers) || !match.IsMatching(newPath, matchers)) && newPath != "" {
		return nil, errors.New("unsupported extension used")
	}

	// End Error Handlers

	if newPath == "" {
		newPath = "out.webp"
		fmt.Printf("\nNo output file specified\nDefaulted to " + newPath + "\n\n")
	}

	mainFormatStr := strings.Split(path, ".")[len(strings.Split(path, "."))-1]
	outFormatStr := strings.Split(newPath, ".")[len(strings.Split(newPath, "."))-1]

	if mainFormatStr == outFormatStr {
		return nil, errors.New("nothing to convert (same format)")
	}

	if newPath == "" {
		newPath = "dist/" + newPath
	}

	switch mainFormatStr {
	case "png":
		image.RegisterFormat(mainFormatStr, mainFormatStr, png.Decode, png.DecodeConfig)
	case "jpeg":
		image.RegisterFormat(mainFormatStr, mainFormatStr, jpeg.Decode, jpeg.DecodeConfig)
	case "webp":
		image.RegisterFormat(mainFormatStr, mainFormatStr, webp.Decode, webp.DecodeConfig)
	case "gif":
		image.RegisterFormat(mainFormatStr, mainFormatStr, gif.Decode, gif.DecodeConfig)
	}

	switch outFormatStr {
	case "png":
		image.RegisterFormat(mainFormatStr, mainFormatStr, png.Decode, png.DecodeConfig)
	case "jpeg":
		image.RegisterFormat(mainFormatStr, mainFormatStr, jpeg.Decode, jpeg.DecodeConfig)
	case "webp":
		image.RegisterFormat(mainFormatStr, mainFormatStr, ewebp.Decode, ewebp.DecodeConfig)
	case "gif":
		image.RegisterFormat(mainFormatStr, mainFormatStr, gif.Decode, gif.DecodeConfig)
	}

	file, err := os.Open(path)
	if err != nil {
		log.Panic(err.Error())
	}

	defer file.Close()

	img, _, err := image.Decode(file)

	defer fmt.Println("Converted " + path + " to " + newPath)

	if err != nil {
		log.Panic(err.Error())
	}

	bounds := img.Bounds()

	filer, err := os.Open(path)

	if err != nil {
		log.Panic(err)
	}

	defer filer.Close()

	imgPixels, err := getPixels(filer)
	if err != nil {
		log.Panic(err)
	}

	imagee := image.NewRGBA(image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y))
	if err != nil {
		log.Panic(err.Error())
	}
	x := 0
	fmt.Printf("Converting ...\r")
	for i := 0; i < int(len(imgPixels)*len(imgPixels[0])); i++ {
		indexY := int(i % len(imgPixels))
		if i%len(imgPixels) == 0 {
			x++
			continue
		}
		pixel := imgPixels[indexY][x-1]
		imagee.Set(x-1, int(i%len(imgPixels)), color.RGBA{pixel.R, pixel.G, pixel.B, pixel.A})
	}
	processedFile, _ := os.Create(newPath)
	defer processedFile.Close()
	switch outFormatStr {
	case "png":
		png.Encode(processedFile, imagee)
	case "jpeg":
		jpeg.Encode(processedFile, imagee, &jpeg.Options{Quality: 100})
	case "webp":
		ewebp.Encode(processedFile, imagee, &ewebp.Options{Lossless: true})
	case "gif":
		gif.Encode(processedFile, imagee, nil)
	}

	return imagee, nil
}

func ConvertDir(sDir, xDir, outputFormat string) error {
	files, err := os.ReadDir(sDir)
	if err != nil {
		return err
	}

	for _, f := range files {
		file := f.Name()
		fmt.Println("Converting", file)
		filer := strings.Join(strings.Split(file, ".")[0:len(strings.Split(file, "."))-1], ".")
		if match.IsMatching(file, matchers) {
			_, err = Convert(sDir+file, strings.Join([]string{xDir + filer, outputFormat}, "."))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
