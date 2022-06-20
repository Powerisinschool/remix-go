package converter

import (
	"fmt"
	ewebp "github.com/chai2010/webp"
	"golang.org/x/image/webp"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"image/gif"
	"io"
	"log"
	"os"
)

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

func PNGToJPEG(path string, newFilename string) bool {
	if newFilename == "" {
		newFilename = "out.jpeg"
		fmt.Printf("No output file given\nDefaulted to " + newFilename)
	}

	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)

	file, err := os.Open(path)
	if err != nil {
		log.Panic(err.Error())
	}

	defer file.Close()

	img, _, err := image.Decode(file)

	defer fmt.Println("Converted " + path + " to " + newFilename)

	if err != nil {
		log.Panic(err.Error())
	}

	bounds := img.Bounds()
	fmt.Println(bounds)

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
	fmt.Println(len(imgPixels))
	x := 0
	fmt.Println("Compiling...")
	for i:=0; i<int(len(imgPixels)*len(imgPixels[0])); i++ {
		indexY := int(i%len(imgPixels))
		if (i%len(imgPixels) == 0) {
			x++;
			continue
		}
		pixel := imgPixels[indexY][x - 1]
		imagee.Set(x-1, int(i%len(imgPixels)), color.RGBA{pixel.R, pixel.G, pixel.B, pixel.A})
	}
	processedFile, _ := os.Create(newFilename)
	defer processedFile.Close()
	jpeg.Encode(processedFile, imagee, nil)

	return true
}

func PNGToGIF(path string, newFilename string) bool {
	if newFilename == "" {
		newFilename = "out.gif"
		fmt.Printf("No output file given\nDefaulted to " + newFilename)
	}

	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("gif", "gif", gif.Decode, gif.DecodeConfig)

	file, err := os.Open(path)
	if err != nil {
		log.Panic(err.Error())
	}

	defer file.Close()

	img, _, err := image.Decode(file)

	defer fmt.Println("Converted " + path + " to " + newFilename)

	if err != nil {
		log.Panic(err.Error())
	}

	bounds := img.Bounds()
	fmt.Println(bounds)

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
	fmt.Println(len(imgPixels))
	x := 0
	fmt.Println("Compiling...")
	for i:=0; i<int(len(imgPixels)*len(imgPixels[0])); i++ {
		indexY := int(i%len(imgPixels))
		if (i%len(imgPixels) == 0) {
			x++;
			continue
		}
		pixel := imgPixels[indexY][x - 1]
		imagee.Set(x-1, int(i%len(imgPixels)), color.RGBA{pixel.R, pixel.G, pixel.B, pixel.A})
	}
	processedFile, _ := os.Create(newFilename)
	defer processedFile.Close()
	gif.Encode(processedFile, imagee, nil)

	return true
}

func PNGToWEBP(path string, newFilename string) bool {
	if newFilename == "" {
		newFilename = "out.webp"
		fmt.Printf("No output file given\nDefaulted to " + newFilename)
	}

	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("webp", "webp", webp.Decode, webp.DecodeConfig)

	file, err := os.Open(path)
	if err != nil {
		log.Panic(err.Error())
	}

	defer file.Close()

	img, _, err := image.Decode(file)

	defer fmt.Println("Converted " + path + " to " + newFilename)

	if err != nil {
		log.Panic(err.Error())
	}

	bounds := img.Bounds()
	fmt.Println(bounds)

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
	fmt.Println(len(imgPixels))
	x := 0
	fmt.Println("Compiling...")
	for i:=0; i<int(len(imgPixels)*len(imgPixels[0])); i++ {
		indexY := int(i%len(imgPixels))
		if (i%len(imgPixels) == 0) {
			x++;
			continue
		}
		pixel := imgPixels[indexY][x - 1]
		imagee.Set(x-1, int(i%len(imgPixels)), color.RGBA{pixel.R, pixel.G, pixel.B, pixel.A})
	}
	processedFile, _ := os.Create(newFilename)
	defer processedFile.Close()
	ewebp.Encode(processedFile, imagee, &ewebp.Options{Lossless: true})

	return true
}