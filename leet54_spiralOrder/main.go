package main

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	upSide := 0
	leftSide := 0
	rightSide := n - 1
	downSide := m - 1

	res := []int{}

	for len(res) < m*n {
		if upSide <= downSide {
			for j := leftSide; j <= rightSide; j++ {
				res = append(res, matrix[upSide][j])
			}
			upSide++
		}
		if leftSide <= rightSide {
			for j := upSide; j <= downSide; j++ {
				res = append(res, matrix[j][rightSide])
			}
			rightSide--
		}
		if upSide <= downSide {
			for j := rightSide; j >= leftSide; j-- {
				res = append(res, matrix[downSide][j])
			}
			downSide--
		}
		if leftSide <= rightSide {
			for j := downSide; j >= upSide; j-- {
				res = append(res, matrix[j][leftSide])
			}
			leftSide++
		}
	}
	return res
}
