package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Create("multiwriter.txt")

	if err != nil {
		panic(err)
	}

	writer := io.MultiWriter(file, os.Stdout)
	io.WriteString(writer, "io.MultiWriter Example\n")
	io.WriteString(writer, "My name is konbu310\n")
}
