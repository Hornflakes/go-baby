package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(3)
	sum(3, 7, 17)

	nums := []int{3, 7, 17, 23}
	sum(nums...)
}
