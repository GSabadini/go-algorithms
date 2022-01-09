package search

import (
	"sort"
)

func Binary(array []int, target int, low int, high int) bool {
	sort.Ints(array)

	if low > high || len(array) == 0 {
		return false
	}

	var mid = low + (high-low)/2

	if array[mid] > target {
		return Binary(array, target, low, mid-1)
	}

	if array[mid] < target {
		return Binary(array, target, mid+1, high)
	}

	if array[mid] == target {
		return true
	}

	return false
}

func BinaryNotRegression(array []int, target int) bool {
	sort.Ints(array)

	value := sort.SearchInts(array, target)

	if len(array) == value {
		return false
	}

	if target == array[value] {
		return true
	}

	return false
}

func BinaryNotRegression_2(array []int, target int) bool {
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
