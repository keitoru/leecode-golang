package main

import "fmt"

// 按照 顺时针螺旋顺序 ，返回矩阵中的所有元素
func spiralOrder(matrix [][]int) []int {
	var res []int
	rowTop, colLeft, rowDown, colRight := 0, 0, len(matrix)-1, len(matrix[0])-1

	for {
		// left -> right
		for i := colLeft; i <= colRight; i++ {
			res = append(res, matrix[rowTop][i])
		}

		rowTop++
		if rowTop > rowDown {
			break
		}

		// top -> down
		for i := rowTop; i <= rowDown; i++ {
			res = append(res, matrix[i][colRight])
		}

		colRight--
		if colRight < colLeft {
			break
		}

		// right -> left
		for i := colRight; i >= colLeft; i-- {
			res = append(res, matrix[rowDown][i])
		}

		rowDown--
		if rowDown < rowTop {
			break
		}

		// down -> top
		for i := rowDown; i >= rowTop; i-- {
			res = append(res, matrix[i][colLeft])
		}

		colLeft++
		if colLeft > colRight {
			break
		}

	}

	return res

}

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}

	res := spiralOrder(matrix)

	fmt.Println(res)
}
