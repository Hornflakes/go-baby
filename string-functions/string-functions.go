package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p("Contains:  ", s.Contains("stranger things", "r t"))
	p("Count:     ", s.Count("stranger things", "t"))
	p("HasPrefix: ", s.HasPrefix("stranger things", "st"))
	p("HasSuffix: ", s.HasSuffix("stranger things", "gs"))
	p("Index:     ", s.Index("stranger things", "t"))
	p("Join:      ", s.Join([]string{"demo", "gorgon"}, ""))
	p("Repeat:    ", s.Repeat("e", 7))
	p("Replace:   ", s.ReplaceAll("foo fighters", "o", "0"))
	p("Replace:   ", s.Replace("foo fighters", "o", "0", 1))
	p("Split:     ", s.Split("a-e-o-u", "-"))
	p("ToLower:   ", s.ToLower("LOW"))
	p("ToUpper:   ", s.ToUpper("up"))
}
