package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func intstr(s int) {
	nb := '0'
	for i := 0; i < s/10; i++ {
		nb++
	}
	z01.PrintRune(nb)
	nb = '0'
	for i := 0; i < s%10; i++ {
		nb++
	}
	z01.PrintRune(nb)
}

func main() {
	points := &point{}

	setPoint(points)

	result := "x = a, y = b"
	for _, i := range result {
		if i == 'a' {
			intstr(points.x)
		} else if i == 'b' {
			intstr(points.y)
		} else {
			z01.PrintRune(i)
		}
	}
	z01.PrintRune('\n')
}
