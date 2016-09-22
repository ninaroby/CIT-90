package main

import (
	"bufio"
	"fmt"
	"io"
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
		s := bufio.NewScanner(conn)
		for s.Scan() {
			line := s.Text()
			fmt.Fprintln(conn, line)
			if line == "" {
				break
			}
		}

		// write to connection
		io.WriteString(conn, "I made a connection with you! Ha Ha!")

		conn.Close()
	}
}
