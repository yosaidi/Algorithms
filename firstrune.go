package piscine

func FirstRune(s string) rune {
	if len(s) == 0 {
		return 0
	}
	MySlice := []rune(s)
	return MySlice[0]
}
