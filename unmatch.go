package piscine

func calculOccur(a []int, nbr int) bool {
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] == nbr {
			count++
		} else {
			continue
		}
	}
	return count%2 == 0
}

func Unmatch(a []int) int {
	for i := 0; i < len(a); i++ {
		if calculOccur(a, a[i]) {
			continue
		} else {
			return a[i]
		}
	}
	return -1
}
