package main

type NumMatrix struct {
	preSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix)
	n := len(matrix[0])
	preSum := make([][]int, m+1)
	for p, _ := range preSum {
		preSum[p] = make([]int, n+1)
	}
	for j := 1; j < m+1; j++ {
		for k := 1; k < n+1; k++ {
			preSum[j][k] = preSum[j-1][k] + preSum[j][k-1] + matrix[j-1][j-1] - preSum[j-1][j-1]
		}
	}
	return NumMatrix{preSum: preSum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.preSum[row2+1][col2+1] - this.preSum[row1][col2+1] - this.preSum[row2+1][col1] + this.preSum[row1][col1]
}
