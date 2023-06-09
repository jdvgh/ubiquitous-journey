package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	data, err := os.ReadFile("/tmp/data")
	check(err)
	fmt.Print(string(data))

	file, err := os.Open("/tmp/data")
	check(err)

	b1 := make([]byte, 5)
	n1, err := file.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	o2, err := file.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := file.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	o3, err := file.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(file, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = file.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(file)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	file.Close()
}
