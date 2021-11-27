package binaryserrch

// 二分查找
func BinarySearch(target int, arr []int) int {
	// arr 应为升序
	left := 0
	right := len(arr) - 1
	for left <= right {
		// 数轴上线段中点坐标是多少？？
		mid := (right + left) / 2
		if target < arr[mid] {
			right = mid - 1
		} else if target > arr[mid] {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
