//go:build windows

package ipcstr

import "fmt"

func GetIPCString(ipcName string) string {
	return fmt.Sprintf(`\\.\pipe\%s`, ipcName)
}
