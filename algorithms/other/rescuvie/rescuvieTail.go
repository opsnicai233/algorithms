package rescuvie

func RescuvieTail(N int, total int) int {
	if N == 1 {
		return total
	}
	// 将每次运算的结果作为下次运算的参数
	// 进行N次运算
	return RescuvieTail(N-1, total*N)
}
