package piscine

func Split(s, sep string) []string {
	if len(s) == 0 {
		return nil
	}
	start := 0
	lenSep := len(sep)
	result := make([]string, 0)
	for i := 0; i <= len(s)-lenSep; i++ {
		if s[i:i+lenSep] == sep {
			result = append(result, s[start:i])
			start = i + lenSep
		}
	}
	result = append(result, s[start:])
	return result
}
