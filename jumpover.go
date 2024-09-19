package piscine

func JumpOver(str string) string {
	if len(str) == 0 || len(str) < 3 {
		return "\n"
	}
	s := ""
	for i := 2; i < len(str); i += 3 {
		s += string(str[i])
	}
	return string(s) + "\n"
}
