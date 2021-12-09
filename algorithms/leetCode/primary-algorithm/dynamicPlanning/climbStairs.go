package dynamicPlanning

//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//注意：给定 n 是一个正整数。

//// 执行超时
//func ClimbStairs(n int) int {
//	return goDeep(0, 0, n)
//}
//
//func goDeep(direction, sum, target int) int {
//	sum += direction
//	if sum > target {
//		return 0
//	}
//	if sum == target {
//		return 1
//	}
//	//sum += direction
//	return goDeep(1, sum, target) + goDeep(2, sum, target)
//}

// 如何爬上第n阶台阶？从第n-1阶爬一阶或从第n-2阶爬两阶
// 所以爬n阶所需的方法数 = 爬n-1阶所需的方法数 + 爬n-2阶所需的方法数
// f(n) = f(n-1) + f(n-2)  --> 斐波那契数列
// f(0)=0； f(1)=1

func ClimbStairs(n int) int {
	a, b, i := 0, 1, 0
	for i < n {
		a, b = b, a+b
		i++
	}
	return b
}
