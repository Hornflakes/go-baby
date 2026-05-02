package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "puffin 1"
	messages <- "puffin 2"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
