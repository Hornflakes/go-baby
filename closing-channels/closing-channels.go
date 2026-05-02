package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("off i go, then", j)
			} else {
				fmt.Println("job's done")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("yes, milord?", j)
	}
	close(jobs)
	fmt.Println("sent all commands")

	<-done

	_, ok := <-jobs
	fmt.Println("more commands:", ok)
}
