package string

//实现strStr()函数。
//
//给你两个字符串haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。
//如果不存在，则返回 -1 。

func StrStr(haystack string, needle string) int {
	nh, nn := len(haystack), len(needle)
	if nn == 0 {
		return 0
	}
	if nn > nh {
		return -1
	}
	for i := 0; i < nh; i++ {
		if i+nn <= nh && haystack[i:i+nn] == needle {
			return i
		}
	}
	return -1
}
