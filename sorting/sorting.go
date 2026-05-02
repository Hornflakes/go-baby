package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	doggos := []string{"dalmatian", "australian shepherd", "border collie", "chihuahua"}
	slices.Sort(doggos)
	fmt.Println("doggos:", strings.Join(doggos, ", "))

	nums := []int{36, 34, 164}
	slices.Sort(nums)
	fmt.Println("nums:  ")

	s := slices.IsSorted(nums)
	fmt.Println("sorted:", s)
}
