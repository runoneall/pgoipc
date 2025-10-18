//go:build windows

package server

import (
	"fmt"
	"net"

	"github.com/runoneall/pgoipc/ipcstr"

	"github.com/Microsoft/go-winio"
)

func Serv(ipcName string, onConnect func(conn net.Conn)) {
	ipcString := ipcstr.GetIPCString(ipcName)

	listener, err := winio.ListenPipe(ipcString, nil)
	if err != nil {
		panic(fmt.Errorf("不能监听 Named Pipe: %v", err))
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("不能接受连接:", err)
		}

		go func() {
			onConnect(conn)
			conn.Close()
		}()
	}
}
