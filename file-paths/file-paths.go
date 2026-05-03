package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	p := filepath.Join("misc", "games", "the last of us")
	fmt.Println("p:", p)

	fmt.Println(filepath.Join("games//", "sims"))
	fmt.Println(filepath.Join("games/../games", "sims"))

	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	fmt.Println(filepath.IsAbs("games/assassins creed"))
	fmt.Println(filepath.IsAbs("/games/assassins creed"))

	cwd, _ := os.Getwd()
	absPath := filepath.Join(cwd, "games", "assassins creed")
	fmt.Println(filepath.IsAbs(absPath))

	filename := "games.json"

	ext := filepath.Ext(filename)
	fmt.Println(ext)

	fmt.Println(strings.TrimSuffix(filename, ext))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
