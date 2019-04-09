package main

import (
	"fmt"
	"strings"
)

func main() {
	var builder strings.Builder
	builder.Write([]byte("This is first line\n"))
	builder.Write([]byte("This is second line\n"))
	fmt.Println(builder.String())
}
