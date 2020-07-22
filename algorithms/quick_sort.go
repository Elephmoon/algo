package algorithms

import "math/rand"

func QuickSort(data []int32) []int32 {
	if len(data) < 2 {
		return data
	}
	left := 0
	right := len(data) - 1
	pivot := rand.Int() % len(data)
	data[pivot], data[right] = data[right], data[pivot]
	for i := range data {
		if data[i] < data[right] {
			data[left], data[i] = data[i], data[left]
			left++
		}
	}
	data[left], data[right] = data[right], data[left]
	QuickSort(data[:left])
	QuickSort(data[left+1:])
	return data
}
