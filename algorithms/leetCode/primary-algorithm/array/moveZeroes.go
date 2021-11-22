package array

/*给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
输入: [0,1,0,3,12]
输出: [1,3,12,0,0]*/

// MoveZeroes
//思路一：
// 遍历数组，把 0 直接删除，并记下删了几个0，
// 把 0 添加到数组末尾
// 效率很低
func MoveZeroes(nums []int) {
	n, i := 0, 0
	for i < len(nums) {
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			n++
		} else {
			i++
		}

	}
	for j := 0; j < n; j++ {
		nums = append(nums, 0)
	}
}

func MoveZeroesPro(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	slow, fast := 0, 0
	for fast < n {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
}

func MoveZeroesProMax(nums []int) {
	n := len(nums)
	for i, j := 0, 0; i < n; i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}
