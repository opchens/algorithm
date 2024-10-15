package main

func transpose(matrix [][]int) [][]int {
	res := make([][]int, len(matrix[0]))

	for i := 0; i < len(matrix[0]); i++ {
		res[i] = make([]int, len(matrix))
	}

	for j, _ := range matrix {
		for k, _ := range matrix[j] {
			res[k][j] = matrix[j][k]
		}
	}
	return res
}
