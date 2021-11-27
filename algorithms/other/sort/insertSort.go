package sort

// InsertSort 插入排序
func InsertSort(arr []int) []int {
	// 默认 arr 第0个元素已经排好序
	for i := 1; i < len(arr); i++ {
		// 后面的元素作为待检查的元素判断是否要插入到前面元素中
		checkItem := arr[i] // 后面元素
		j := i - 1          // 前面元素的索引
		if checkItem < arr[j] {
			// 前面元素向后移腾出位置
			for j >= 0 && checkItem < arr[j] {
				arr[j+1] = arr[j]
				j = j - 1
			}
			arr[j+1] = checkItem
		}
	}
	return arr
}
