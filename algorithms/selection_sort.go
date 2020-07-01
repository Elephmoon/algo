package algorithms

func SelectionSort(input []int) []int {
	var n = len(input)
	for i := range input {
		var minIdx = i
		for j := i; j < n; j++ {
			if input[j] < input[minIdx] {
				minIdx = j
			}
		}
		input[i], input[minIdx] = input[minIdx], input[i]
	}
	return input
}
