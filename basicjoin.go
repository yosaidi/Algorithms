package piscine

func BasicJoin(elems []string) string {
	if len(elems) == 0 {
		return ""
	}
	result := elems[0]
	for i := 1; i < len(elems); i++ {
		result += elems[i]
	}
	return result
}
