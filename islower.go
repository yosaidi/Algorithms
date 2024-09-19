package piscine

func IsLower(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, char := range s {
		if char < 'a' || char > 'z' {
			return false
		}
	}
	return true
}
