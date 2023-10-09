package main

import "fmt"

func setZeroes(matrix [][]int) {
	col0_flag := false
	cols, rows := len(matrix[0]), len(matrix)

	for i := 0; i < rows; i++ {
		if matrix[i][0] == 0 {
			col0_flag = true
		}
		for j := 1; j < cols; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0], matrix[0][j] = 0, 0
			}
		}
	}

	for i := rows - 1; i >= 0; i-- {
		for j := cols - 1; j > 0; j-- {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}

		if col0_flag {
			matrix[i][0] = 0
		}
	}
}

func main() {
	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}

	setZeroes(matrix)

	fmt.Println(matrix)
}
