package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	log.Fatal(listenTCP())
}

func listenTCP() error {
	var listenAddr string
	flag.StringVar(&listenAddr, "listen-addr", "127.0.0.1:9000", "address to listen")
	flag.Parse()

	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return fmt.Errorf("error listening on %q: %s", listenAddr, err)
	}
	defer func() {
		closeErr := l.Close()
		if closeErr != nil {
			log.Printf("can't close listen socket: %s", err)
		}
	}()

	fmt.Printf("Listening on %q\n", listenAddr)

	for {
		conn, err := l.Accept()
		defer conn.Close()

		if err != nil {
			fmt.Printf("Error accepting connection %q: %s", listenAddr, err)
			time.Sleep(100 * time.Millisecond)
			continue
		}

		if err := handleRequest(conn); err != nil {
			log.Printf("Error handling request: %s", err)
			time.Sleep(100 * time.Millisecond)
			continue
		}
	}
}

func handleRequest(conn net.Conn) error {
	buf := make([]byte, 1024)
	reqLen, err := conn.Read(buf)
	if err != nil {
		return fmt.Errorf("can't read data from connection: %s", err)
	}
	log.Printf("got request of len %d bytes: %q", reqLen, buf[:reqLen])

	if _, err := conn.Write([]byte("Message received.")); err != nil {
		return fmt.Errorf("can't write to connection: %s", err)
	}

	return nil
}
