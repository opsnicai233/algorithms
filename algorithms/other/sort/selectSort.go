package sort

// SelectSort 简单选择排序
// 每轮都用开头元素和后面元素一一比较，若开头元素小于后面元素，则交换位置
// 每轮之后开头元素都应是该轮元素中最小值
// 不断缩小每轮冒泡的元素范围
func SelectSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				swap(&arr[i], &arr[j])
			}
		}
		//fmt.Printf("%d,", arr[i])
	}
	return arr
}
