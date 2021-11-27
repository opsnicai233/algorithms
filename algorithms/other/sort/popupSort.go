package sort

// PopupSortFunc 冒泡排序：
// 每次将一个数和它相邻的数比较，若该数较大，则交换位置;
// 每轮都把把该轮最大元素放在末尾;
// 不断缩小每轮冒泡元素个数;
func PopupSortFunc(arr []int) []int {
	for i := len(arr) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		//fmt.Printf("%d,", arr[i])
	}
	return arr
}

func swap(a, b *int) {
	*a, *b = *b, *a
}
