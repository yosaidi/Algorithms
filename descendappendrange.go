package piscine

func DescendAppendRange(max, min int) []int {
	mySlice := []int{}
	if max <= min {
		return mySlice
	}
	for i := max; i > min; i-- {
		mySlice = append(mySlice, i)
	}
	return mySlice
}
