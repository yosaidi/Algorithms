package piscine

func NRune(s string, n int) rune {
	if n <= 0 || n > len(s) {
		return 0
	}
	if len(s) == 0 {
		return 0
	}
	MySlice := []rune(s)
	return MySlice[n-1]
}
