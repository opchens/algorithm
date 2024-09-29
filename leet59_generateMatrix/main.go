package main

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	upside := 0
	leftside := 0
	downside := n - 1
	rightside := n - 1

	num := 1

	for num <= n*n {
		if upside <= downside {
			for j := leftside; j <= rightside; j++ {
				matrix[upside][j] = num
				num++
			}
			upside++
		}
		if leftside <= rightside {
			for j := upside; j <= downside; j++ {
				matrix[j][rightside] = num
				num++
			}
			rightside--
		}
		if upside <= downside {
			for j := rightside; j >= leftside; j-- {
				matrix[downside][j] = num
				num++
			}
			downside--
		}
		if leftside <= rightside {
			for j := downside; j >= upside; j-- {
				matrix[j][leftside] = num
				num++
			}
			leftside++
		}
	}
	return matrix
}
