package array

// 旋转图像
// 给定一个 n×n 的二维矩阵matrix 表示一个图像。请你将图像顺时针旋转 90 度。
//你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。

// RotateMatrix 按“圈”旋转矩阵,矩阵最外圈为第0圈
// i控制第几圈，j 控制‘圈’ 的边长
func RotateMatrix(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-2-i; j++ {
			RotateFourItems(matrix, i, j, n, matrix[i][j], 0)
		}
	}
}

// RotateFourItems 旋转四个相关的元素
// 把当前元素matrix[i][j]放到新的位置 matrix[j][n-1-i]
func RotateFourItems(matrix [][]int, i int, j int, n, newItem int, counter int) {
	if counter >= 4 {
		return
	}
	// 元素matrix[i][j] 旋转后新的位置为 matrix[j][n-1-i]
	newitem := matrix[j][n-1-i]
	matrix[j][n-1-i] = newItem
	RotateFourItems(matrix, j, n-1-i, n, newitem, counter+1)
}
