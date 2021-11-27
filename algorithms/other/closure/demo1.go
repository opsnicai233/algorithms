package closure

import "log"

func func2() int {
	val := 10
	// defer 在函数 return 之后执行
	defer func() {
		val += 1
		log.Println("test")
	}()

	return val
}

func func2pro() (val int) {
	// func2pro 和 func2 函数的区别在于
	// func2pro 会初始化 val，并分配内存地址
	// 之后的操作，包括闭包都会涉及到对 val 内存的修改
	val = 10
	// defer 在函数 return 之后执行
	defer func() {
		val += 1
		log.Println("test")
	}()

	return val
}

func GetFunc2Val() int {
	// func2函数只有在运行结束（return、defer）之后才会返回
	return func2()
}

func GetFunc2proVal() int {
	return func2pro()
}
