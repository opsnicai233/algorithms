package dynamicPlanning

import "math"

// 最大子序和
//给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//子数组 是数组中的一个连续部分。

// 官方答案
func MaxSubArray(nums []int) int {
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1]+nums[i] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > sum {
			sum = nums[i]
		}
	}
	return sum
}

func MaxSubArrayPro(nums []int) int {
	n := len(nums)
	var dp = make([]int, n)
	dp[0] = nums[0]
	sum := nums[0]
	for i := 1; i < n; i++ {
		dp[i] = MaxAB(nums[i]+dp[i-1], nums[i])
		sum = MaxAB(dp[i], sum)
	}
	return sum
}

// 此解法的重点在于若prev对 当前值无增益作用， 则直接更新 prev
func MaxSubArrayPro2(nums []int) int {
	prev, ans := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if prev > 0 {
			prev += nums[i]
		} else {
			prev = nums[i]
		}
		ans = int(math.Max(float64(ans), float64(prev)))
	}
	return ans
}

func MaxSubArrayProMax(nums []int) int {
	n := len(nums)

	sum, prev := nums[0], nums[0]
	for i := 1; i < n; i++ {
		prev = int(math.Max(float64(nums[i]+prev), float64(nums[i])))
		sum = int(math.Max(float64(prev), float64(sum)))
	}
	return sum
}
