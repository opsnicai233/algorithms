package string

// ReverseString 反转字符串
func ReverseString(s []byte) {
	last := len(s) - 1
	for start := 0; start <= last; start++ {
		if start >= last {
			break
		}
		s[start], s[last] = s[last], s[start]
		last--
	}
}
