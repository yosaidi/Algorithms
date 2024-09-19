package piscine

func AtoiBase(s string, base string) int {
	basechars := make(map[rune]bool)
	for _, char := range base {
		if char == '-' || char == '+' || len(base) < 2 || basechars[char] {
			return 0
		}
		basechars[char] = true
	}
	for _, char := range base {
		if !basechars[char] {
			return 0
		}
	}
	baseValue := make(map[rune]int)
	for i, char := range base {
		baseValue[char] = i
	}
	baselen := len(base)
	result := 0
	for _, char := range s {
		result = result*baselen + baseValue[char]
	}
	return result
}
