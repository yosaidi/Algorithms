package piscine

func ToLower(s string) string {
	lower := []rune(s)
	difference := 'a' - 'A'
	for i, char := range s {
		if char >= 'A' && char <= 'Z' {
			lower[i] = char + difference
		}
	}
	return string(lower)
}
