package main

import (
	"fmt"
	"unicode/utf8"
)

func examineRune(r rune) {
	switch r {
	case 't':
		fmt.Println("found tee")
	case 'ส':
		fmt.Println("found so sua")
	}
}

func main() {
	const s = "สวัสดี"

	fmt.Println("len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("rune count:", utf8.RuneCountInString(s))

	for idx, runeVal := range s {
		fmt.Printf("%#U starts at %d\n", runeVal, idx)
	}

	fmt.Println("\nusing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeVal, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeVal, i)
		w = width

		examineRune(runeVal)
	}
}
