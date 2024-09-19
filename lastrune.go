package piscine

func LastRune(s string) rune {
	if len(s) == 0 {
		return 0
	}
	MySlice := []rune(s)
	a := len(s)
	return MySlice[a-1]
}
