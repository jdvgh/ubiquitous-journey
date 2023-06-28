package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	data := "test-email+with_different&.symbols@testmail.com"

	enc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(enc)

	dec, _ := base64.StdEncoding.DecodeString(enc)
	fmt.Printf("%s\n\n", string(dec))

	encU := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(encU)
	decU, _ := base64.URLEncoding.DecodeString(encU)
	fmt.Println(string(decU))

}
