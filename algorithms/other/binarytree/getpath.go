package binarytree

const null = -99999999

// bug: 收集节点错误
func GoDeep(i, sum, target int, arr []int) (result []int) {
	n := len(arr)
	if i < n && arr[i] != null {
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
