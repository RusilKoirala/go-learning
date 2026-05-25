package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	for {
		fmt.Printf("Select \n[1] Read file\n[2] Write file\n[3] Exit\n")
		reader := bufio.NewReader(os.Stdin)
		choice, _ := reader.ReadString('\n')
		switch choice {
		case "1\n":
			fmt.Print("Enter file name : ")
			fileName, _ := reader.ReadString('\n')
			ReadFile(fileName + ".txt")
		case "2\n":
			fmt.Print("Enter file name : ")
			fileName, _ := reader.ReadString('\n')
			file := MakeFile(fileName)
			fmt.Print("Enter content : ")
			content, _ := reader.ReadString('\n')
			WriteFile(file, content)
		case "3\n":
			break
		}
	}
}

func ReadFile(file string) {
	content, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	println(content)
}

func MakeFile(FileName string) *os.File {
	file, err := os.Create(FileName + ".txt")
	if err != nil {
		panic(err)
	}
	return file
}

func WriteFile(file *os.File, content string) {
	len, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println(len)
	defer file.Close()
}
