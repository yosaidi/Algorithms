package piscine

func Join(strs []string, sep string) string {
	if len(strs) == 0 && len(sep) == 0 {
		return ""
	}
	var result1 string
	for i := 0; i < len(strs); i++ {
		result1 += strs[i]
		if i != len(strs)-1 {
			result1 += sep
		}
	}
	return result1
}
