package main

func searchMatrix(matrix [][]int, target int) bool {
	rows, colums := len(matrix)-1, 0
	for rows >= 0 && colums < len(matrix[0]) {
		num := matrix[rows][colums]

		if num == target {
			return true
		} else if num > target {
			rows--
		} else {
			colums++
		}
	}

	return false
}
