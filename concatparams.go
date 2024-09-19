package piscine

func ConcatParams(args []string) string {
	if len(args) == 0 {
		return ""
	}
	slicelen := 0
	for _, arg := range args {
		slicelen += len(arg)
	}
	result := make([]byte, 0, slicelen)
	for i, arg := range args {
		if i != 0 {
			result = append(result, '\n')
		}
		result = append(result, arg...)
	}
	return string(result)
}
