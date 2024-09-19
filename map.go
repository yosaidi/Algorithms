package piscine

func Map(f func(int) bool, a []int) []bool {
	result := make([]bool, len(a))
	for i, x := range a {
		result[i] = f(x)
	}
	return result
}
