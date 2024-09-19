package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}
	digitCount := [10]int{}
	if n == 0 {
		digitCount[0] = 1
	} else {
		temp := n
		for temp > 0 {
			digit := temp % 10
			digitCount[digit]++
			temp /= 10
		}
	}
	for digit := 0; digit < 10; digit++ {
		for count := 0; count < digitCount[digit]; count++ {
			z01.PrintRune(rune('0' + digit))
		}
	}
}
