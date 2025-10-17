package main

import (
	"os"
	"path/filepath"
)

var ipcFile string

func main() {
	rootDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	ipcFile = filepath.Join(rootDir, "ipc.sock")

	switch os.Args[1] {

	case "s":
		main_server()

	case "c":
		main_client()

	}
}
