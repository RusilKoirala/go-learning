package main

import (
	"fmt"
	"os"
)

func main() {

	// File open
	file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	// Files Stat
	fileinfo, err := file.Stat()

	if err != nil {
		panic(err)
	}
	fmt.Println(fileinfo.Name())
	fmt.Println(fileinfo.IsDir())
	fmt.Println(fileinfo.Size())

	// Buffer
	buf := make([]byte, fileinfo.Size())
	file.Read(buf)
	fmt.Println(string(buf))

	// Create new file
	func() {
		newFile, err := os.Create("home.txt")
		if err != nil {
			panic(err)
		}
		os.WriteFile("home.txt", buf, os.ModeAppend)
		fmt.Println(newFile)
		fmt.Println("The file is succesfully created")
	}()

	defer file.Close()
}
