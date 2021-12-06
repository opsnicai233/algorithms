package sort_search

//你是产品经理，目前正在带领一个团队开发新的产品。不幸的是，你的产品的最新版本没有通过质量检测。由于每个版本都是基于之前的版本开发的，所以错误的版本之后的所有版本都是错的。
//
//假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。
//
//你可以通过调用bool isBadVersion(version)接口来判断版本号 version 是否在单元测试中出错。实现一个函数来查找第一个错误的版本。你应该尽量减少对调用 API 的次数。

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func isBadVersion(version int) bool {
	return false
}

func firstBadVersion(n int) int {
	return check(1, n)
}

func check(low, high int) int {
	if low == high {
		return low
	}
	if high-low == 1 {
		if isBadVersion(low) {
			return low
		} else {
			return high
		}
	}
	mid := (low + high) / 2
	if isBadVersion(mid) == false {
		low = mid
	} else {
		high = mid
	}
	return check(low, high)
}

// 迭代法
func firstBadVersion2(n int) int {
	low, high := 1, n
	for low < high {
		sum := low + high
		mid := sum / 2
		if isBadVersion(mid) {
			high = mid
		} else {
			low = mid
		}
		if high-low == 1 {
			if isBadVersion(low) {
				return low
			}
			return high
		}
	}
	return low
}
