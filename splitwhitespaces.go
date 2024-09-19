package piscine

func SplitWhiteSpaces(s string) []string {
	if len(s) == 0 {
		return nil
	}
	var result []string
	word := ""
	for _, char := range s {
		if char == ' ' || char == '\n' || char == '\t' {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		} else {
			word += string(char)
		}
	}
	if word != "" {
		result = append(result, word)
	}
	return result
}
