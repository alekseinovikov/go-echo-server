package main

import (
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":2042")
	if err != nil {
		log.Fatalf("Error on listening: %v", err)
	}

	log.Printf("Listening on %s", listener.Addr().String())
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error on accepting connection: %v", err)
			continue
		}

		go echo(conn)
	}
}

func echo(conn net.Conn) {
	defer conn.Close()

	for {
		i, err := io.Copy(conn, conn)
		log.Printf("Echoed %d bytes", i)
		if err == io.EOF {
			log.Printf("Connection closed")
			break
		}

		if err != nil {
			log.Printf("Error on reading/writing data: %v", err)
			break
		}
	}
}
