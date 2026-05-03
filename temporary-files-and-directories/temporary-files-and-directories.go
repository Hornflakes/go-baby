package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.CreateTemp("", "porsche")
	check(err)

	fmt.Println("temp file name:", f.Name())

	defer os.Remove(f.Name())

	_, err = f.Write([]byte{9, 1, 1})
	check(err)

	dname, err := os.MkdirTemp("", "cars")
	check(err)
	fmt.Println("temp dir name:", dname)

	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "land rover defender")
	err = os.WriteFile(fname, []byte{9, 0}, 0666)
	check(err)
}
