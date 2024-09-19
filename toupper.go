package piscine

func ToUpper(s string) string {
	upper := []rune(s)
	difference := 'a' - 'A'
	for i, char := range s {
		if char >= 'a' && char <= 'z' {
			upper[i] = char - difference
		}
	}
	return string(upper)
}
