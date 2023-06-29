package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	match, err := regexp.MatchString("p([a-z]+)ch", "peach")

	if err != nil {
		fmt.Println("err occured: %v", err)
		return
	}
	fmt.Println(match)

	r, err := regexp.Compile("p([a-z]+)ch")

	if err != nil {
		fmt.Println("err occured: %v", err)
		return
	}

	fmt.Println(r.MatchString("peach"))

	fmt.Println(r.FindString("peach punch"))

	fmt.Println("idx: ", r.FindStringIndex("peach punch"))

	fmt.Println(r.FindStringSubmatch("peach punch"))

	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	fmt.Println(r.FindAllString("peach punch pinch", -1))

	fmt.Println(r.FindAllStringIndex("peach punch pinch", -1))

	fmt.Println(r.FindAllStringIndex("peach punch pinch", 2))

	fmt.Println(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")

	fmt.Println("mustcompile regex: ", r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))

}
