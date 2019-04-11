package main

import (
	"fmt"
	"net"
)

func isError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Server is running at localhost:8888")

	conn, err := net.ListenPacket("udp", "localhost:8888")
	isError(err)

	defer conn.Close()

	buffer := make([]byte, 1500)

	for {
		length, remoteAddress, err := conn.ReadFrom(buffer)
		isError(err)

		fmt.Printf("Received from %v: %v\n",
			remoteAddress,
			string(buffer[:length]),
		)

		_, err = conn.WriteTo([]byte("Hello from Server"), remoteAddress)
		isError(err)
	}
}
