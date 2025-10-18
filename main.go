package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"pgoipc/client"
	"pgoipc/server"
	"strings"
)

func main_handleConn(conn net.Conn) {
	reader := bufio.NewReader(conn)

	data, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	data = strings.TrimSuffix(data, "\n")

	fmt.Println("Recv:", data)

	if _, err := fmt.Fprintln(conn, "Repeat:", data); err != nil {
		panic(err)
	}
}

func main_client(conn net.Conn) {
	if _, err := fmt.Fprintln(conn, "Hello World!"); err != nil {
		panic(err)
	}

	reader := bufio.NewReader(conn)

	data, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	data = strings.TrimSuffix(data, "\n")

	fmt.Println("Recv:", data)
}

func main() {
	ipcName := "pgoipc_debug"

	switch os.Args[1] {

	case "s":
		server.Serv(ipcName, main_handleConn)

	case "c":
		main_client(client.Connect(ipcName))

	}
}
