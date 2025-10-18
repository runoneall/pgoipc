//go:build windows

package client

import (
	"fmt"
	"net"

	"github.com/runoneall/pgoipc/ipcstr"

	"github.com/Microsoft/go-winio"
)

func Connect(ipcName string) net.Conn {
	ipcString := ipcstr.GetIPCString(ipcName)

	conn, err := winio.DialPipe(ipcString, nil)
	if err != nil {
		panic(fmt.Errorf("不能连接到 Named Pipe: %v", err))
	}

	return conn
}
