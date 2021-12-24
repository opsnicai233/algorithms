package array

// Rotate 给你一个数组，将数组中的元素向右轮转 k个位置，其中k是非负数。
//
//输入: nums = [1,2,3,4,5,6,7], k = 3
//输出: [5,6,7,1,2,3,4]
func Rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	start := n - k
	// o(N)
	//newnums := append(nums[start:],nums[:start]...)

	newnums := make([]int, n)
	for i := 0; i < n; i++ {
		if i < start {
			newnums[i+k] = nums[i]
		} else {
			newnums[i-start] = nums[i]
		}
	}
	copy(nums, newnums)
}

func RotateProMax(nums []int, k int) {
	n := len(nums)
	k = k % n
	newnums := make([]int, n)
	for i := 0; i < n; i++ {
		newnums[(i+k)%n] = nums[i]
	}
	copy(nums, newnums)
}
