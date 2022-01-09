package others

func MaxInt(array []int) int {
	var max int

	for _, value := range array {
		if value > max {
			max = value
		}
	}
	return max
}
