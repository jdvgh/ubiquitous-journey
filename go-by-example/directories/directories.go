package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	err := os.Mkdir("subdir", 0755)
	check(err)

	defer os.RemoveAll("subdir")

	createEmptyFile := func(name string) {
		d := []byte("")
		err := os.WriteFile(name, d, 0644)
		check(err)
	}

	createEmptyFile(filepath.Join("subdir", "file1"))

	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile(filepath.Join("subdir", "parent", "file2"))
	createEmptyFile(filepath.Join("subdir", "parent", "file3"))
	createEmptyFile(filepath.Join("subdir", "parent", "child", "file4"))

	c, err := os.ReadDir(filepath.Join("subdir", "parent"))
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir(filepath.Join("subdir", "parent", "child"))

	c, err = os.ReadDir(".")

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../..")
	check(err)

	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit)
	err = filepath.WalkDir("subdir", visitDir)
}

func visit(p string, info os.FileInfo, err error) error {
	if p != "" {
		fmt.Printf("%v : Name: %v, IsDir: %v\n", p, info.Name(), info.IsDir())
	}
	if err != nil {
		return err
	}
	return nil
}

func visitDir(p string, d os.DirEntry, err error) error {
	if p != "" {
		fmt.Printf("%v : Name: %v, IsDir: %v\n", p, d.Name(), d.Type().IsDir())
	}
	if err != nil {
		return err
	}
	return nil
}
