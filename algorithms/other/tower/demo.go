package tower

import "fmt"

var Total = 0

// 汉诺塔问题
func Tower(n int, a, b, c string) {
	if n == 1 {
		fmt.Println(a, "->", c)
		Total += 1
		return
	}
	// 把A柱上的上面n-1个盘子移到B柱
	Tower(n-1, a, c, b)

	// 把A柱上的最后一个盘子移到C柱
	fmt.Println(a, "->", c)
	Total += 1
	// 把B柱上的n-1个盘子移到C柱
	Tower(n-1, b, a, c)
}
