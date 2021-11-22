package array

/*给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
你可以假设除了整数 0 之外，这个整数不会以零开头。

输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。*/

func PlusOne(digits []int) []int {
	var result []int
	isALLNine := true
	for _, digit := range digits {
		if digit != 9 {
			isALLNine = false
			break
		}
	}
	n := len(digits)
	if isALLNine {
		// 如果digits全是9
		result = append(result, 1)
		for i := 0; i < n; i++ {
			result = append(result, 0)
		}
	} else {
		for i := n - 1; i >= 0; i-- {
			// 若第i位不为9，把后面几位依次加一，第i位加一
			if digits[i] != 9 {
				digits[i] += 1
				return digits
			} else {
				digits[i] = 0
			}
		}
	}
	return result
}
