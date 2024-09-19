package piscine

func Capitalize(s string) string {
	var result []rune
	inWord := false

	for _, char := range s {
		if IsAlphanumeric(char) {
			if !inWord {
				result = append(result, Upper(char))
				inWord = true
			} else {
				result = append(result, Lower(char))
			}
		} else {
			result = append(result, char)
			inWord = false
		}
	}
	return string(result)
}

func Lower(char rune) rune {
	if char >= 'A' && char <= 'Z' {
		return char + ('a' - 'A')
	}
	return char
}

func Upper(char rune) rune {
	if char >= 'a' && char <= 'z' {
		return char - ('a' - 'A')
	}
	return char
}
