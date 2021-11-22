package array

/*给定两个数组，编写一个函数来计算它们的交集。
输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2,2]
输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[4,9]*/

// Intersect 对于未排序的两个数组
func Intersect(nums1 []int, nums2 []int) []int {
	var smallArr, bigArr []int
	if len(nums1) < len(nums2) {
		smallArr = nums1
		bigArr = nums2
	} else {
		smallArr = nums2
		bigArr = nums1
	}
	var arr []int
	const null = -999999999
	for _, num1 := range smallArr {
		for i := 0; i < len(bigArr); i++ {
			if num1 == bigArr[i] {
				arr = append(arr, num1)
				bigArr[i] = null
				break
			}
		}
	}
	return arr
}

// InterSectSorted 如果nums1 和 nums2 已经排好序
func InterSectSorted(nums1, nums2 []int) []int {
	nums1Len, nums2Len := len(nums1), len(nums2)
	var arr []int
	i, j := 0, 0
	for i < nums1Len && j < nums2Len {
		switch {
		case nums1[i] == nums2[j]:
			arr = append(arr, nums1[i])
			i++
			j++
		case nums1[i] < nums2[j]:
			i++
		case nums1[i] > nums2[j]:
			j++
		}
	}
	return arr
}
