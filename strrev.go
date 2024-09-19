package piscine

func StrRev(s string) string {
	runeArray := []rune(s) // Convert the string to a slice of runes to handle multi-byte characters
	length := len(runeArray)
	for i := 0; i < length/2; i++ {
		// Swap the characters
		runeArray[i], runeArray[length-1-i] = runeArray[length-1-i], runeArray[i]
	}
	return string(runeArray) // Convert the slice of runes back to a string and return it
}
