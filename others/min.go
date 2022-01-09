package others

func MinInt(array []int) int {
	var min = array[0]

	for _, value := range array {
		if value < min {
			min = value
		}
	}
	return min
}
