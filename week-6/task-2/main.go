package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
		}

		// read from connection
		bs, err := ioutil.ReadAll(conn)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(string(bs))

		// write to connection
		io.WriteString(conn, "I see you connected")

		conn.Close()
	}
}
