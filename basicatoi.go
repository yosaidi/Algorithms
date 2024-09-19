package piscine

func BasicAtoi(s string) int {
	result := 0
	for _, char := range s {
		digit := int(char - '0')   // Convert the character to its corresponding integer value
		result = result*10 + digit // Accumulate the result
	}
	return result
}
