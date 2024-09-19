package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	myslice := []rune(base)
	for _, char := range base {
		if char == '-' || char == '+' || !UniqueChars(myslice) || len(myslice) < 2 {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
	}
	if nbr == -9223372036854775808 { // Edge case for minimum int value
		z01.PrintRune('-')
		Remainder(922337203685477580, base)
		z01.PrintRune('8') // Manually handle the last digit
		return
	}
	if nbr < 0 {
		z01.PrintRune('-')
		nbr = -nbr
	}
	Remainder(nbr, base)
}

func Remainder(nbr int, base string) {
	mySlice := []rune(base)
	lenbase := len(base)
	var result []int
	for nbr > 0 {
		rem := nbr % lenbase
		result = append(result, rem)
		nbr /= lenbase
	}
	if len(result) == 0 {
		z01.PrintRune(mySlice[0])
	}
	for i := len(result) - 1; i >= 0; i-- {
		z01.PrintRune(mySlice[result[i]])
	}
}

func UniqueChars(slice []rune) bool {
	charmap := make(map[rune]bool)
	for _, char := range slice {
		if _, exist := charmap[char]; exist {
			return false
		}
		charmap[char] = true
	}
	return true
}
