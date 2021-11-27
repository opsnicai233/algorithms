package fib

// 斐波那切 数列
func Fib(N int, a, b int) int {
	if N == 1 {
		return a
	}
	return Fib(N-1, b, a+b)
}
