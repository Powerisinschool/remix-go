package main

import (
	"image-converter/converter"
	"image-converter/server"
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

		_, err := converter.Convert(args[1], secondArg)

		if err != nil {
			help(err)
		}

		// Handle
		if len(args) > 3 {
			if args[3] == "--open" {
				server.ServeImages(secondArg)
			}
		}
	}

	if len(args) == 2 {
		panic("Coming soon...\nInput at least 1 argument")
	}
}

func help(err error) {
	panic(err.Error())
}
