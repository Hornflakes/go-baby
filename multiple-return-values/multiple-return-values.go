package main

import "fmt"

func vals() (int, int) {
	return 34, 36
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}
