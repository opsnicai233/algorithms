package GreatestCommonDivisor

// 欧几里得算法求最大公约数

// 自然语言描述
// 计算两个非负整数a、b的最大公约数
// 若b为0，则最大公约数为 a
// a和b的余数为r，则a和b的最大公约数即为b和r的最大公约数
func Gcd(a, b uint32) uint32 {
	if b == 0 {
		return a
	}
	r := a % b
	return Gcd(b, r)
}
