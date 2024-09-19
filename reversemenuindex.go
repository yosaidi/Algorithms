package piscine

func ReverseMenuIndex(menu []string) []string {
	Slice := []string{}
	for i := len(menu) - 1; i >= 0; i-- {
		Slice = append(Slice, menu[i])
	}
	return Slice
}
