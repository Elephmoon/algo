package main

import (
	"algo/algorithms"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arrSize := 100000000
	list := make([]int, arrSize)
	elem := rand.Intn(arrSize)

	for i := 0; i < arrSize; i++ {
		list[i] = i
	}

	startBinary := time.Now()
	sortedPosition, found := algorithms.BinarySearchInt(list, elem)
	fmt.Println("Binary search time = ", time.Since(startBinary))
	if found {
		fmt.Println(sortedPosition)
	} else {
		fmt.Println("Not found ", elem)
	}

	startStupid := time.Now()
	sortedPosition, found = algorithms.StupidSearchInt(list, elem)
	fmt.Println("Stupid search time = ", time.Since(startStupid))
	if found {
		fmt.Println(sortedPosition)
	} else {
		fmt.Println("Not found", elem)
	}
}
