package array

import "log"

/*
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
输入：nums = [0,0,1,1,1,2,2,3,3,4]
输出：5, nums = [0,1,2,3,4]
解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。
*/
func RemoveDuplicates(arr []int) int {
	if len(arr) <= 1{
		return len(arr)
	}
	var numMap = map[int]int{}
	numMap[arr[0]] = 1
	j := 0
	for i:=0;i<len(arr);i++{
		if _, ok:= numMap[arr[i]]; !ok {
			numMap[arr[i]] = 1
			j += 1
			arr[j] = arr[i]
		}
	}
	log.Println(arr[:j+1])
	return len(arr[:j+1])
}
