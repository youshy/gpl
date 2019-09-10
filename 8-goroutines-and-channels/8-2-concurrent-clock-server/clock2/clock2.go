// Clock1 is a TCP server that periodically writes the time
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // f.e. connection aborted
			continue       // even if error occurs, continue
		}
		// this will introduce concurrency
		// in this case, it will be able to handle
		// multiple connections
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // f.e. client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
