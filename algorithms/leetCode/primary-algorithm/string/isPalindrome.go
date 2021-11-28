package string

import (
	"strings"
)

/*给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:
输入: "A man, a plan, a canal: Panama"
输出: true
解释："amanaplanacanalpanama" 是回文串*/

func IsPalindrome(s string) bool {
	n := len(s)
	if n == 0 {
		return true
	}
	s = strings.ToLower(s)
	i, j := 0, n-1
	for i < j {
		if !isValidLetter(s[i]) {
			i++
			continue
		}
		if !isValidLetter(s[j]) {
			j--
			continue
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// isPalindrome
func isValidLetter(letter uint8) bool {
	return (letter >= 'a' && letter <= 'z') || (letter >= '0' && letter <= '9')
}

// 递归写法，占用内存较高
func IsPalindrome2(s string) bool {
	s = strings.ToLower(s)
	return checkSubstring(s, 0, len(s)-1)
}

func checkSubstring(s string, first, last int) bool {
	for first < last && !isValidLetter(s[first]) {
		first++
	}
	for last > 0 && !isValidLetter(s[last]) {
		last--
	}
	if last <= first {
		return true
	}
	if s[first] != s[last] {
		return false
	}
	first++
	last--
	return checkSubstring(s, first, last)
}
