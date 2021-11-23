package array

/*给定一个整数数组 nums和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那两个整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案。

示例 1：
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。*/

func TwoSum(nums []int, target int) []int {
	var result []int
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return append(result, i, j)
			}
		}
	}
	return result
}

func TwoSum2(nums []int, target int) []int {
	var result []int
	numMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		numMap[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		v, ok := numMap[target-nums[i]]
		if ok && v != i {
			return append(result, i, v)
		}
	}

	return result
}

// TwoSumProMax 思路：
//
//
//
func TwoSumProMax(nums []int, target int) []int {
	numMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if v, ok := numMap[target-num]; ok {
			return []int{v, i}
		}
		numMap[num] = i
	}
	return nil
}
