package sicp

func fRecursive(n int) int {
	if n < 3 {
		return n
	}
	return fRecursive(n-1) + fRecursive(n-2) + fRecursive(n-3)
}

func fWithIter(n int) int {
	return fIter(2, 1, 0, n)
}

func fIter(a, b, c, count int) int {
	if count == 0 {
		return c
	}
	return fIter(a+b+c, a, b, count-1)
}
