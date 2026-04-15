package main

import (
	"io"
	"net"
	"time"
)

func main() {

	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
		//time.Sleep(10*time.Second)
	}
}

func handleConnection(c net.Conn) {

	defer c.Close()
	for {
		_, err := io.WriteString(c, "response from server")
		if err != nil {
			return
		}
		time.Sleep(10 * time.Second)
	}

}
