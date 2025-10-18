//go:build unix

package ipcstr

import "os"

func remove(path string) {
	if _, err := os.Stat(path); err == nil {
		if err := os.RemoveAll(path); err != nil {
			panic(err)
		}
	}
}

func isexist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
