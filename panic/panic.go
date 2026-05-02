package main

import (
	"os"
	"path/filepath"
)

func main() {
	// panic("problem of war")

	path := filepath.Join(os.TempDir(), "file")
	_, err := os.Create(path)
	if err != nil {
		panic(err)
	}
}
