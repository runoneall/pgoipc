//go:build unix

package ipcstr

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetIPCString(ipcName string, autoClean bool) string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	baseDir := filepath.Join(homeDir, ".pgoipc")
	if !isexist(baseDir) {
		if err := os.MkdirAll(baseDir, 0750); err != nil {
			panic(err)
		}
	}

	ipcFile := filepath.Join(baseDir, fmt.Sprintf("%s.sock", ipcName))
	if isexist(ipcFile) && autoClean {
		remove(ipcFile)
	}

	return ipcFile
}
