package main

import (
	"flag"
	"fmt"
	"github.com/Elephmoon/algo/algorithms"
	"strconv"
)

const (
	fibAlg           = 1
	selectionSortAlg = 2
)

var (
	alg = flag.String("alg", "0", "Choose an algorithm: \n "+
		"1 - Fibonacci numbers \n "+
		"2 - SelectionSort []int \n")
)

func main() {
	flag.Parse()
	inputFlag, err := strconv.Atoi(*alg)
	if err != nil {
		panic(err)
	}
	switch inputFlag {
	case fibAlg:
		fmt.Println(fibonacci())
	case selectionSortAlg:
		fmt.Println(selectionSort())
	}
}

func selectionSort() []int {
	fmt.Println("Input []int like 1 2 3 4 -0. -0 signifies the end of the array")
	var input []int
	var n int
	for {
		if _, err := fmt.Scanf("%d", &n); err != nil {
			panic(err)
		}
		if n == -0 {
			break
		}
		input = append(input, n)
	}
	return algorithms.SelectionSort(input)
}

func fibonacci() uint32 {
	fmt.Println("Enter number")
	var n uint32
	if _, err := fmt.Scanf("%d", &n); err != nil {
		panic(err)
	}
	return algorithms.Fib(uint32(n))
}
