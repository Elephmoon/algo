package algorithms

func Fib(n uint32) uint32 {
	if n == 0 {
		return 0
	}
	fib := make([]uint32, n+1)
	fib[1] = 1
	var i uint32
	for i = 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib[n]
}
