package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() { messages <- "pinghy" }()

	msg := <-messages
	fmt.Println(msg)
}
