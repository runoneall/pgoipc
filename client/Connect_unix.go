//go:build unix

package client

import (
	"fmt"
	"net"

	"github.com/runoneall/pgoipc/ipcstr"
)

func Connect(ipcName string) net.Conn {
	ipcString := ipcstr.GetIPCString(ipcName, false)

	conn, err := net.Dial("unix", ipcString)
	if err != nil {
		panic(fmt.Errorf("不能连接到 Unix 域: %v", err))
	}

	return conn
}
