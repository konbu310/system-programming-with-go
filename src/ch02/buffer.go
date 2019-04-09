package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer Example\n"))
	buffer.Write([]byte("This is Second"))
	fmt.Println(buffer.String())
}
