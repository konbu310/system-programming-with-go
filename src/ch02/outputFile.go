package main

import (
	"os"
)

func main() {
	file, crtErr := os.Create("text.txt")

	if crtErr != nil {
		panic(crtErr)
	}

	file.Write([]byte("My name is konbu310"))

	file.Close()
}
