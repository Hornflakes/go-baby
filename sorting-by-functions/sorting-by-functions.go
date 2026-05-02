package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	jams := []string{"peach", "cherry", "orange", "blueberry", "strawberry"}

	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(jams, lenCmp)
	fmt.Println(jams)

	type Band struct {
		name string
		rank int
	}

	bands := []Band{
		{name: "queen", rank: 1},
		{name: "evanescence", rank: 0},
		{name: "pink floyd", rank: 3},
		{name: "linkin park", rank: 2},
	}

	slices.SortFunc(bands,
		func(a, b Band) int {
			return cmp.Compare(a.rank, b.rank)
		})
	fmt.Println(bands)
}
