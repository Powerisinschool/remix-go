package main

import (
	// "fmt"
	"image-converter/converter"
	// "image-converter/match"
	_ "net/http"
	"os"
)

// © 2022 Tolulope Olagunju

func main() {
	args := os.Args

	if len(args) < 2 {
		panic("Not yet supported\nInput at least 1 argument")
	}

	if len(args) > 2 {
		secondArg := args[2]
		if secondArg == "x" {
			secondArg = ""
		}
		// if match.IsMatch(args[1], "*.png") {
		// 	Convert()
		// } else {
		// 	help()
		// }

		err := converter.Convert(args[1], secondArg)

		if err != nil {
			help(err)
		}
	}

	if len(args) == 2 {
		panic("Coming soon...\nInput at least 1 argument")
	}

	// You can register another format here
	// image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	// image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)

	// file, err := os.Open("./image.png")

	// if err != nil {
	// 	fmt.Println("Error: File could not be opened")
	// 	os.Exit(1)
	// }

	// defer file.Close()

	// fmt.Println(converter.PNGToJPEG("./image.png", "image.jpeg"))
	// fmt.Println(converter.PNGToJPEG("./image.png", "image.gif"))
	// fmt.Println(converter.PNGToWEBP("./image.png", "image.webp"))

	// if err != nil {
	// 	fmt.Println("Error: Image could not be decoded")
	// 	os.Exit(1)
	// }
}

func help(err error) {
	panic(err.Error())
}
