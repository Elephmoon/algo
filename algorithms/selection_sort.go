package algorithms

func SelectionSort(input []int) []int {
	for i := range input {
		smallest := i
		for i2, j := range input {
			if j < input[smallest] {
				smallest = i2
			}
			input[i2], input[smallest] = input[smallest], input[i2]
		}
	}
	return input
}
