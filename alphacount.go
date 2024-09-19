package piscine

func AlphaCount(s string) int {
	if len(s) == 0 {
		return 0
	}
	count := 0
	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			count++
		}
	}
	return count
}
