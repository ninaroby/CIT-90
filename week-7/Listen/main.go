package main

import (
	"io"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer li.Close()

	conn, err := li.Accept()
	if err != nil {
		panic(err)
	}

	io.WriteString(conn, "I Accept!\n")

	conn.Close()
}
