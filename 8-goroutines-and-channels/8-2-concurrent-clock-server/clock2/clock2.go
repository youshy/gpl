// Clock1 is a TCP server that periodically writes the time
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var port = flag.Int("port", 8080, "listen port")

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
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
