package piscine

func TrimAtoi(s string) int {
	var sign int = 1
	var result int = 0
	var foundNumber bool = false

	for i := 0; i < len(s); i++ {
		char := s[i]

		if char == '-' {
			if foundNumber {
				continue
			}
			sign = -1
		} else if char == '+' {
			if foundNumber {
				continue
			}
			sign = 1
		} else if char >= '0' && char <= '9' {
			foundNumber = true
			digit := int(char - '0')
			result = result*10 + digit
		}
	}
	if !foundNumber {
		return 0
	}
	return sign * result
}
