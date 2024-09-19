package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("File name missing")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error opening file: %v\n ", err)
		return
	}
	defer file.Close()

	arr := make([]byte, 1024)

	n, err := file.Read(arr)
	if err != nil && err != io.EOF {
		fmt.Printf("Error reading file: %v\n ", err)
		return
	}
	fmt.Print(string(arr[:n]))
}
