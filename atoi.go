package piscine

// Atoi converts a string of digits (with an optional leading + or -) into an integer
func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	sign := 1
	start := 0

	if s[0] == '+' {
		start = 1
	} else if s[0] == '-' {
		sign = -1
		start = 1
	}

	result := 0
	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' { // Check if the character is not a digit
			return 0
		}
		digit := int(s[i] - '0')   // Convert the character to its corresponding integer value
		result = result*10 + digit // Accumulate the result
	}

	return sign * result
}
