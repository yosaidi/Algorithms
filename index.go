package piscine

func Index(s string, toFind string) int {
	if len(s) == 0 {
		return 0
	}
	if len(toFind) == 0 {
		return 0
	}
	slen := len(s)
	tofindlen := len(toFind)

	for i := 0; i <= slen-tofindlen; i++ {
		if s[i:i+tofindlen] == toFind {
			return i
		}
	}
	return -1
}
