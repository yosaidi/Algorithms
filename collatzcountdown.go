package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}
	result := 0
	num := start
	for num != 1 {
		if num%2 == 0 {
			num /= 2
		} else {
			num = num*3 + 1
		}
		result++
	}
	return result
}
