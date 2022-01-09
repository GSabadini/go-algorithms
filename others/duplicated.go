package others

import "sort"

func LinearSearchDuplicatedInt(array []int, target int) bool {
	for _, value := range array {
		if value == target {
			return true
		}
	}

	return false
}

func BinarySearchDuplicatedInt(array []int, target int) bool {
	sort.Ints(array)

	var low = 0
	var high = len(array) - 1
	var mid int

	for low <= high {
		mid = low + (high-low)/2

		if array[mid] > target {
			high = mid - 1
			continue
		}

		if array[mid] < target {
			low = mid + 1
			continue
		}

		return true
	}

	return false
}
