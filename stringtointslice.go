package piscine

func StringToIntSlice(str string) []int {
	s := []rune(str)
	result := []int{}
	if len(str) == 0 {
		return nil
	}
	for _, char := range s {
		result = append(result, int(char))
	}
	return result
}
