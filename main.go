package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/runoneall/pgoipc/client"
	"github.com/runoneall/pgoipc/server"
)

func main_handleConn(conn net.Conn) {
	reader := bufio.NewReader(conn)
	req, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	req = strings.TrimSpace(req)

	fmt.Println("Recv:", req)

	msg := fmt.Sprintf("REQ GET %s", req)

	time.Sleep(time.Second)

	fmt.Println("Send:", msg)
	fmt.Fprintln(conn, msg)
}

func main_client(ipcName string) {
	for i := range 10 {
		client.Connect(ipcName, func(conn net.Conn) {
			msg := fmt.Sprintf("MSG IDX %d", i)

			fmt.Println("Send:", msg)
			fmt.Fprintln(conn, msg)

			reader := bufio.NewReader(conn)
			rep, err := reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			rep = strings.TrimSpace(rep)

			fmt.Println("Recv:", rep)

			time.Sleep(time.Second)
		})
	}
}

func main() {
	ipcName := "pgoipc_example"

	switch os.Args[1] {

	case "s":
		server.Serv(ipcName, main_handleConn)

	case "c":
		main_client(ipcName)

	}
}
