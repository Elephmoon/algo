package main

import (
	"flag"
	"fmt"
	"github.com/Elephmoon/algo/algorithms"
	"strconv"
)

const (
	fibAlg = 1
)

var (
	alg = flag.String("alg", "0", "Choose an algorithm. 1 - Fibonacci numbers")
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
	}
}

func fibonacci() uint32 {
	fmt.Println("Enter number")
	var n uint32
	if _, err := fmt.Scanf("%d", &n); err != nil {
		panic(err)
	}
	return algorithms.Fib(uint32(n))
}
