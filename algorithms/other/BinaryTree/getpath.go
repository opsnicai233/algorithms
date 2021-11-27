package BinaryTree

const Null = -99999999

// 目标：找到二叉树从根节点到叶子节点的某个路径，使得该路径下所有节点和为target
// 放弃
// bug: 收集节点错误
/*
试图在不使用二叉树这个数据结构的前提下，直接通过array实现目标
很难很难！
因此数据结构和算法是分不开的！
*/
func GoDeep(i, sum, target int, arr []int) (result []int) {
	n := len(arr)
	if i < n && arr[i] != Null {
		sum += arr[i]
		result = append(result, arr[i])
		if 2*i+1 <= n {
			result = append(result, GoDeep(2*i+1, sum, target, arr)...)
		}
		if 2*i+2 <= n {
			result = append(result, GoDeep(2*i+1, sum, target, arr)...)
		}
	} else {
		if sum == target {
			return result
		}
	}
	return
}
