package array

import "log"

// SingleNumber
//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
//算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
//示例 1:
//输入: [2,2,1]
//输出: 1
func SingleNumber(nums []int) int {
	numMap := map[int]int{}
	for _, num := range nums {
		if _, ok := numMap[num]; ok {
			numMap[num] = 1
		}
		numMap[num] += 1
	}
	log.Println(numMap)
	for _, num := range nums {
		if numMap[num] == 1 {
			return num
		}
	}
	return -1
}

func SingleNumberPro(nums []int) int {
	// 利用集合的元素单一性
	// 集合所有元素 *2 - nums == singleNumber

	return 0
}

// 利用异或运算
func SingleNumberProMax(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}
