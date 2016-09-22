package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		io.WriteString(c, fmt.Sprint("Hi from the Server: Hi", time.Now(), "\n"))

		c.Close()
	}
}
