package string

import (
	"strconv"
)

// CountAndSay https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnpvdm/
// CountAndSay 性能有待提升
func CountAndSay(n int) string {
	i, s := 1, "1"
	for i < n {
		s = f(s)
		i++
	}
	return s
}

func f(s string) string {
	result, nums := "", 1
	for i := 0; i < len(s); i += nums {
		nums = count(i, s)
		result += strconv.Itoa(nums) + string(s[i])
	}
	return result
}

func count(i int, s string) int {
	nums := 1
	for i+1 < len(s) && s[i] == s[i+1] {
		nums += 1
		i += 1
	}
	return nums
}

func CountAndSay2(n int) string {
	return getResult(1, n, "1")
}

func getResult(i, n int, result string) string {
	if i == n {
		return result
	}
	s, nums := "", 1
	for i := 0; i < len(result); i += nums {
		nums = count(i, result)
		result += strconv.Itoa(nums) + string(s[i])
	}
	return getResult(i+1, n, s)
}
