//go:build unix

package server

import (
	"fmt"
	"net"
	"pgoipc/ipcstr"
)

func Serv(ipcName string, onConnect func(conn net.Conn)) {
	ipcString := ipcstr.GetIPCString(ipcName)

	listener, err := net.Listen("unix", ipcString)
	if err != nil {
		panic(fmt.Errorf("不能监听 unix 域: %v", err))
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
