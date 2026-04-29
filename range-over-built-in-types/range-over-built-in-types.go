package main

import "fmt"

func main() {
	nums := []int{3, 7, 17, 23}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 7 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"b": "brie", "c": "camembert"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go baby" {
		fmt.Println(i, c)
	}
}
