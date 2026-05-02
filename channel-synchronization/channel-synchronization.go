package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("bz bz...")
	time.Sleep(time.Second)
	fmt.Println("honey")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	<-done
}
