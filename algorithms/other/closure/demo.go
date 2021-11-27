package closure

import "log"

// 考察闭包

func func1() func() int {
	sum := 0
	return func() int {
		// 闭包使用sum
		// sum 再此函数内部为 指针
		sum += 1
		return sum
	}
}

func GetSums() (int, int) {
	getSum := func1()
	sum1 := getSum()
	sum2 := getSum()
	log.Println("test")
	return sum1, sum2
}
