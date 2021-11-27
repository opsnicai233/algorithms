package rescuvie

func Rescuvie(N int) int {
	if N == 1 {
		return 1
	}
	// 每次运算都会加长运算链条，
	// 不仅会拖慢速度，进行了 2N 次运算
	// 还会造成栈溢出
	return N * Rescuvie(N-1)
}
