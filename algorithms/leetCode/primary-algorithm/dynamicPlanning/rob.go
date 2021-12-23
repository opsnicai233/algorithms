package dynamicPlanning

// leetcode 打家劫舍

func Rob(arr []int) int {
	// 必须初始化map
	var cache = make(map[int]int)
	return robHelper(len(arr)-1, arr, cache)
}

func robHelper(i int, arr []int, cache map[int]int) int {
	if i < 0 {
		return 0
	}
	var lastlast, last, curr int
	if val, ok := cache[i-2]; ok {
		lastlast = val
	} else {
		lastlast = robHelper(i-2, arr, cache)
		cache[i-2] = lastlast
	}
	if val, ok := cache[i-1]; ok {
		last = val
	} else {
		last = robHelper(i-1, arr, cache)
		cache[i-1] = last
	}
	curr = lastlast + arr[i]
	return MaxAB(last, curr)
}

func RobProMax(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	prev, curr := nums[0], MaxAB(nums[0], nums[1])
	for i := 2; i < n; i++ {
		prev, curr = curr, MaxAB(curr, prev+nums[i])
	}
	return curr
}
