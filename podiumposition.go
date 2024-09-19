package piscine

func PodiumPosition(podium [][]string) [][]string {
	for i := 0; i+1 < len(podium); i++ {
		if podium[i][0] > podium[i+1][0] {
			podium[i][0], podium[i+1][0] = podium[i+1][0], podium[i][0]
			i = -1
		}
	}
	return podium
}
