package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func isErr(err error) {
	if err != nil {
		panic(err)
	}
}

func isGZipAcceptable(request *http.Request) bool {
	return strings.Index(
		strings.Join(request.Header["Accept-Encoding"], ","), "gzip") != -1
}

func processSession(conn net.Conn) {
	fmt.Printf("Accept %v\n", conn.RemoteAddr())

	defer conn.Close()

	request, err := http.ReadRequest(bufio.NewReader(conn))
	isErr(err)

	dump, err := httputil.DumpRequest(request, true)
	isErr(err)

	fmt.Println(string(dump))

	response := http.Response{
		StatusCode: 200,
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
	}

	if isGZipAcceptable(request) {
		content := "Hello World! (gzipped)\n"
		// コンテンツをgzip化して転送
		var buffer bytes.Buffer
		writer := gzip.NewWriter(&buffer)
		io.WriteString(writer, content)
		writer.Close()
		response.Body = ioutil.NopCloser(&buffer)
		response.Header.Set("Content-Encoding", "gzip")
	} else {
		content := "Hello World\n"
		response.Body = ioutil.NopCloser(strings.NewReader(content))
	}
	response.Write(conn)
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	isErr(err)

	fmt.Println("Server is running at http://localhost:8888")

	for {
		conn, err := listener.Accept()
		isErr(err)

		go processSession(conn)
	}
}
