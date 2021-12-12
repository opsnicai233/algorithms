package dynamicPlanning

// 最大子序和
//给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//子数组 是数组中的一个连续部分。

// error: 算法逻辑有问题

import "math"

func MaxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return math.MinInt64
	}
	if n == 1 {
		return nums[0]
	}
	leftIndex := 0
	for i := 0; i < n; i++ {
		if nums[i] < 0 {
			leftIndex = i
			break
		}
	}
	rightIndex := n
	for i := n - 1; i >= 0; i-- {
		if nums[i] < 0 {
			rightIndex = i
			break
		}
	}
	return fn(nums, leftIndex+1, rightIndex, n)
}

func fn(nums []int, leftIndex, rightIndex, n int) int {
	if leftIndex >= rightIndex {
		return math.MinInt64
	}
	leftSum := sumOfArray(nums[:leftIndex+1])
	rightSum := sumOfArray(nums[rightIndex:])

	li := leftIndex
	for i := leftIndex; i < n; i++ {
		if nums[i] < 0 {
			li = i
			break
		}
	}
	ri := rightIndex
	for i := rightIndex - 1; i >= 0; i-- {
		if nums[i] < 0 {
			ri = i
			break
		}
	}
	midSum := fn(nums[leftIndex+1:rightIndex], li, ri, rightIndex-leftIndex+1)
	return MaxAB(MaxAB(leftSum, rightSum), midSum)
}

func sumOfArray(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
