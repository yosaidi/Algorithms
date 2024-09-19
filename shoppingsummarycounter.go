package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	m := make(map[string]int)

	myslice := Slice(str)

	for i := 0; i < len(myslice); i++ {
		m[myslice[i]]++
	}
	return m
}

func Slice(str string) []string {
	myslice := []string{}
	s := ""
	for _, char := range str {
		if char == ' ' {
			myslice = append(myslice, s)
			s = ""
		} else {
			s += string(char)
		}
	}
	myslice = append(myslice, s)
	return myslice
}
