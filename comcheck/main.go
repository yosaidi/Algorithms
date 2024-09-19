package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		return
	}
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "01" || os.Args[i] == "galaxy" || os.Args[i] == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}
