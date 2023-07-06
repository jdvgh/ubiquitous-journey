package main

import (
	"fmt"
	"os"
)

func main() {

	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
	readFile(f)

}

func createFile(p string) *os.File {
	fmt.Println("creating file")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing file")
	fmt.Fprintln(f, "data")
}

func readFile(f *os.File) {
	dat, err := os.ReadFile(f.Name())
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
	fmt.Printf("Read from file : %s\n", string(dat[:]))
}

func closeFile(f *os.File) {
	fmt.Println("closing file")
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
