package isPrime

// 判断一个数是否为素数

func IsPrime(a uint32) bool {
	if a < 2 {
		return false
	}
	for i := uint32(2); i < a; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}
