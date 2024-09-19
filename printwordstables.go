package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	if len(a) == 0 {
		return
	}
	for _, arg := range a {
		for _, char := range arg {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
