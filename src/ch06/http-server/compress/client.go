package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}

	// POST文字列を送るリクエストを作成
	request, err := http.NewRequest(
		"POST",
		"http://localhost:8888",
		strings.NewReader("Hello, World!"),
	)
	if err != nil {
		panic(err)
	}

	request.Header.Set("Accept-Encoding", "gzip")



	err = request.Write(conn)
	if err != nil {
		panic(err)
	}

	// サーバーから読み込む。タイムアウトはここでエラーになるのでリトライ
	response, err := http.ReadResponse(bufio.NewReader(conn), request)
	if err != nil {
		panic(err)
	}

	// 結果を表示
	dump, err := httputil.DumpResponse(response, false)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dump))

	defer response.Body.Close()

	if response.Header.Get("Content-Encoding") == "gzip" {
		reader, err := gzip.NewReader(response.Body)
		if err != nil {
			panic(err)
		}
		io.Copy(os.Stdout, reader)
	} else {
		io.Copy(os.Stdout, response.Body)
	}

	conn.Close()
}
