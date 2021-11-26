package string

import (
	"math"
	"strconv"
)

func ReverseInteger(x int) int {
	if x == 0 {
		return 0
	}
	//for x % 10 == 0{
	//	x /=10
	//}
	start := 0
	if x < 0 {
		start = 1
	}
	numBytes := []byte(strconv.Itoa(x))[start:]
	ReverseString(numBytes)
	result, _ := strconv.Atoi(string(numBytes))
	if result >= math.MaxInt32 {
		return 0
	}
	if x < 0 {
		return result * -1
	}
	return result
}

func ReverseIntegerProMax(x int) int {
	i, num := 0, 0
	for x != 0 {
		num = x%10 + num*10
		x /= 10
		i++
	}
	if num <= math.MinInt32 || num >= math.MaxInt32 {
		return 0
	}
	return num
}
