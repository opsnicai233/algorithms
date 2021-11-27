package string

//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//注意：若s 和 t中每个字符出现的次数都相同，则称s 和 t互为字母异位词。
//示例1:
//输入: s = "anagram", t = "nagaram"
//输出: true
//示例 2:
//输入: s = "rat", t = "car"
//输出: false
//提示:
//1 <= s.length, t.length <= 5 * 104
//s 和 t仅包含小写字母

func isAnagram(s string, t string) bool {
	n := len(s)
	if n != len(t) {
		return false
	}
	// 当s t 为unicode字符时，个人觉得可以使用map统计每个字符出现的次数
	sArr, tArr := [26]int{}, [26]int{}
	for i := 0; i < n; i++ {
		sArr[s[i]-'a']++
		tArr[t[i]-'a']++
	}
	return sArr == tArr
}
