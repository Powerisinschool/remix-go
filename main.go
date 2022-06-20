package main

import (
	"image-converter/converter"
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

		err := converter.Convert(args[1], secondArg)

		if err != nil {
			help(err)
		}
	}

	if len(args) == 2 {
		panic("Coming soon...\nInput at least 1 argument")
	}
}

func help(err error) {
	panic(err.Error())
}
