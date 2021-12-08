package dynamicPlanning

// 执行超时
func ClimbStairs(n int) int {
	return goDeep(0, 0, n)
}

func goDeep(direction, sum, target int) int {
	sum += direction
	if sum > target {
		return 0
	}
	if sum == target {
		return 1
	}
	//sum += direction
	return goDeep(1, sum, target) + goDeep(2, sum, target)
}
