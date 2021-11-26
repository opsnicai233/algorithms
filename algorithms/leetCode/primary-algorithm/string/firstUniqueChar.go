package string

func FirstUniqChar(s string) int {
	charMap := map[uint8]int{}
	for i := 0; i < len(s); i++ {
		if _, ok := charMap[s[i]]; ok {
			charMap[s[i]]++
		} else {
			charMap[s[i]] = 1
		}
	}
	for i := 0; i < len(s); i++ {
		if charMap[s[i]] == 1 {
			return i
		}
	}
	return -1
}

func FirstUniqChar2(s string) int {
	charArr := make([]int, 26)
	for i := 0; i < len(s); i++ {
		// s="a" s[0] == 97
		charArr[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if charArr[s[i]-'a'] == 1 {
			return i
		}
	}
	return 0
}
