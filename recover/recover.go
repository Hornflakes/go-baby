package main

import "fmt"

func panicAtTheDisco() {
	panic("panic at the disco")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered. error:", r)
		}
	}()

	panicAtTheDisco()

	fmt.Println("after panicAtTheDisco()")
}
