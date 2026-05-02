package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("g([a-z]+)s", "gambas")
	fmt.Println(match)

	r, _ := regexp.Compile("g([a-z]+)s")

	fmt.Println(r.MatchString("gambas"))

	fmt.Println(r.FindString("gambas al ajillo"))

	fmt.Println("idx:", r.FindStringIndex("gambas al ajillo"))

	fmt.Println(r.FindStringSubmatch("gambas al ajillo"))

	fmt.Println(r.FindStringSubmatchIndex("gambas al ajillo"))

	fmt.Println(r.FindAllString("gambas gas grass", -1))

	fmt.Println(r.FindAllStringSubmatchIndex("gambas gas grass", -1))

	fmt.Println(r.FindAllString("gambas gas grass", 2))

	fmt.Println(r.Match([]byte("gambas")))

	r = regexp.MustCompile("g([a-z]+)s")
	fmt.Println("regexp:", r)

	fmt.Println(r.ReplaceAllString("las gambas", "<seafood>"))

	in := []byte("las gambas")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
