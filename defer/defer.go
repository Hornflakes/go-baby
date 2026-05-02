package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "pove")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		panic(err)
	}
}

func main() {
	path := filepath.Join(os.TempDir(), "o_sa_iti_pove.txt")
	f := createFile(path)
	defer closeFile(f)
	writeFile(f)
}
