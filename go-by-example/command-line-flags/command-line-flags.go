package main

import (
	"flag"
	"fmt"
)

func main() {
	// Run with e.g. :  ./command-line-flags -word=opt -numb=7 -fork -svar=flag
	wordPtr := flag.String("word", "foo", "a string")

	numPtr := flag.Int("numb", 42, "an integer")

	forkPtr := flag.Bool("fork", false, " a boolean")

	var svar string

	flag.StringVar(&svar, "svar", "bar", "a string variable")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())

}
