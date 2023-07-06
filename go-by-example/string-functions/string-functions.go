package main

import (
	"fmt"
	"strings"
)

var p = fmt.Println

func main() {

	p("Contains: ", strings.Contains("test", "es"))
	p("Count: ", strings.Count("test", "t"))
	p("HasPrefix: ", strings.HasPrefix("test", "te"))
	p("HasSuffix: ", strings.HasSuffix("test", "st"))
	p("Index: ", strings.Index("test", "s"))
	p("Join: ", strings.Join([]string{"test", "es"}, "-"))
	p("Repeat: ", strings.Repeat("test", 5))
	p("Replace: ", strings.Replace("test", "es", "35", -1))
	p("Replace: ", strings.Replace("testtest", "s", "5", 1))
	p("Split: ", strings.Split("test-abcd-efgh", "-"))
	p("ToLower: ", strings.ToLower("test"))
	p("ToUpper: ", strings.ToUpper("test"))
}
