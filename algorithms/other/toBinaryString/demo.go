package toBinaryString

import "fmt"

func ToBinaryString(N int) string {
	s := ""
	for N != 0 {
		s = fmt.Sprintf("%d", N%2) + s
		N /= 2
	}
	return s
}
