package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main_server_handleConn(conn net.Conn) {
	defer conn.Close()

	data, err := io.ReadAll(conn)
	if err != nil {
		panic(err)
	}
	dataStr := string(data)

	fmt.Println("Recv:", dataStr)

	if _, err := fmt.Fprintf(conn, "Repeat: %s", dataStr); err != nil {
		panic(err)
	}
	conn.(*net.UnixConn).CloseWrite()
}

func main_server() {
	if _, err := os.Stat(ipcFile); err == nil {
		if err := os.RemoveAll(ipcFile); err != nil {
			panic(err)
		}
	}

	listener, err := net.Listen("unix", ipcFile)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go main_server_handleConn(conn)
	}
}
