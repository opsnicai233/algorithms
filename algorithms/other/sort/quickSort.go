package sort

func QuickSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	var lArr, rArr, midArr []int
	for i := 0; i < n; i++ {
		if arr[i] < arr[0] {
			lArr = append(lArr, arr[i])
		}
		if arr[i] > arr[0] {
			rArr = append(rArr, arr[i])
		}
		if arr[i] == arr[0] {
			midArr = append(midArr, arr[i])
		}
	}
	var result []int
	result = append(result, QuickSort(lArr)...)
	result = append(result, midArr...)
	result = append(result, QuickSort(rArr)...)
	return result
}
