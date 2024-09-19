package piscine

func IsAlpha(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, char := range s {
		if !IsAlphanumeric(char) {
			return false
		}
	}
	return true
}

func IsAlphanumeric(a rune) bool {
	return (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z') || (a >= '0' && a <= '9')
}
