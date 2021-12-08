package sort_search

//给你两个按 非递减顺序 排列的整数数组nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
//
//请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
//
//注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，
//其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。

// Merge 递归写法
func Merge(nums1 []int, m int, nums2 []int, n int) {
	putVal(nums1, nums2, m-1, n-1)
}

func putVal(nums1, nums2 []int, i, j int) {
	if i < 0 && j < 0 {
		return
	}
	if j < 0 {
		return
	}
	if i < 0 {
		for k := 0; k <= j; k++ {
			nums1[k] = nums2[k]
		}
		return
	}
	if nums1[i] <= nums2[j] {
		nums1[i+j+1] = nums2[j]
		j--
	} else {
		nums1[i+j+1], nums1[i] = nums1[i], nums1[i+j+1]
		i--
	}
	putVal(nums1, nums2, i, j)
}

func Merge2(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	for i >= 0 || j >= 0 {
		if i < 0 {
			for k := 0; k <= j; k++ {
				nums1[k] = nums2[k]
			}
			return
		}
		if nums1[i] <= nums2[j] {
			nums1[i+j+1] = nums2[j]
			j--
		} else {
			nums1[i+j+1], nums1[i] = nums1[i], nums1[i+j+1]
			i--
		}
	}
}