package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0

	for _, x := range tab {
		if f(x) {
			count++
		}
	}

	return count
}
