package main

import (
	"image-converter/converter"
	"github.com/Powerisinschool/pxl"
	"image-converter/server"
	"image-converter/match"
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
		if len(args) > 2 {
			consider := "";
			openers := []string{"--open", "--open-server"}

			if match.IsMatching(args[2], openers) {
				secondArg = "out.webp"
			}
			
			if (len(args) > 3) {
				consider = args[3]
			} else {
				consider = args[2]
			}
			if consider == "--open" {
				pxl.Render([]string{secondArg})
				panic("functionality is broken for now!")
			} else if consider == "--open-server" {
				server.ServeImages(secondArg)
			} else {
				panic("Invalid args")
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
