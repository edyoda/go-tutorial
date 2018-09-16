package main

import (
	"io"
	"log"
	"net"
	"time"
)

func handleConnection(c net.Conn) {
	defer c.Close()

	for {
		_, err := io.WriteString(c, "Hello\n")
		if err != nil {
			log.Fatal("err")
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	listen, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal("err")
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal("err")
		}
		go handleConnection(conn)
	}
}
