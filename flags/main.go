package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) == 0 || arguments[0] == "--help" || arguments[0] == "-h" {
		PrintHelp()
		return
	}
	var insertstr string
	var order bool
	var argstr string
	for _, arg := range arguments {
		if len(arg) > 9 && arg[:9] == "--insert=" {
			insertstr = arg[9:]
		} else if len(arg) > 3 && arg[:3] == "-i=" {
			insertstr = arg[3:]
		} else if arg == "--order" || arg == "-o" {
			order = true
		} else {
			argstr = arg
		}
	}
	result := argstr + insertstr
	if order {
		result = Sorting(result)
	}
	printString(result)
}

func PrintHelp() {
	fmt.Printf("--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n")
	fmt.Printf("--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.\n")
}

func Sorting(s string) string {
	MySlice := []rune(s)
	for i := 0; i < len(MySlice)-1; i++ {
		for j := 0; j < len(MySlice)-i-1; j++ {
			if MySlice[j] > MySlice[j+1] {
				MySlice[j], MySlice[j+1] = MySlice[j+1], MySlice[j]
			}
		}
	}
	return string(MySlice)
}

func printString(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
