package algorithms

import "math"

func FastPower(a, n uint32) uint32 {
	if n == 0 {
		return 1
	}
	x := math.Pow(float64(a), float64(n/2))
	if (n % 2) == 0 {
		return uint32(x * x)
	}
	return a * uint32(x*x)
}
