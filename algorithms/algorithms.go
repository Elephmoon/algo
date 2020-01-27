package algorithms

func BinarySearchInt(input []int, item int) (position int, found bool) {
	low := 0
	high := len(input) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := input[mid]
		if guess == item {
			return mid, true
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return 0, false
}

func StupidSearchInt(input []int, item int) (position int, found bool) {
	for i, j := range input {
		if j == item {
			return i, true
		}
	}
	return 0, false
}
