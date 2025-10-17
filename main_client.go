package main

import (
	"fmt"
	"io"
	"net"
)

func main_client() {
	conn, err := net.Dial("unix", ipcFile)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	if _, err := fmt.Fprint(conn, "Hello World!"); err != nil {
		panic(err)
	}
	conn.(*net.UnixConn).CloseWrite()

	data, err := io.ReadAll(conn)
	if err != nil {
		panic(err)
	}

	fmt.Println("Recv:", string(data))
}
