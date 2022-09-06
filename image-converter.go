package main

import (
	"errors"
	"fmt"
	"image-converter/converter"

	// "github.com/Powerisinschool/pxl"
	"image-converter/match"
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

		var err error

		if secondArg[len(secondArg)-1] == '/' {
			if len(args) < 4 {
				help(errors.New("no output format"))
			}
			err = converter.ConvertDir(args[1], secondArg, args[3])
		} else {
			_, err = converter.Convert(args[1], secondArg)
		}

		if err != nil {
			help(err)
		}

		// Handle
		if len(args) > 2 && err != nil {
			consider := ""
			openers := []string{"--open", "--open-server"}
			if match.IsMatching(args[2], openers) {
				secondArg = "out.webp"
			}
			if len(args) > 3 {
				consider = args[3]
			} else {
				consider = args[2]
			}
			if consider == "--open" || consider == "--open-server" {
				server.ServeImages(secondArg)
			} else {
				panic("Invalid args")
			}
		}
	}

	if len(args) == 2 {
		if match.IsMatching(args[1], []string{"--version", "version"}) {
			fmt.Println("v0.1.0")
		} else {
			panic("Coming soon...\nInput at least 1 argument")
		}

	}
}

func help(err error) {
	panic(err.Error())
}
