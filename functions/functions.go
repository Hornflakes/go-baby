package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(6, 7)
	fmt.Println("6+7 =", res)

	res = plusPlus(3, 17, 23)
	fmt.Println("3+17+23 =", res)
}
