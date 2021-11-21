package array

// 一个数组中是否包含重复元素
func containsDuplicate(nums []int) bool {
	numMap := map[int]struct{}{}
	for _, i := range nums {
		if _, ok := numMap[i]; ok {
			return true
		}
		numMap[i] = struct{}{}
	}
	return false
}
