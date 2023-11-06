package main

import "fmt"

func generate(numRows int) [][]int {
	list := make([][]int, 0, numRows)

	for i := 0; i < numRows; i++ {
		row := make([]int, 0, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				row = append(row, 1)
			} else {
				row = append(row, list[i-1][j-1]+list[i-1][j])
			}
		}

		list = append(list, row)
	}

	return list
}

func main() {
	res := generate(5)
	fmt.Println(res)
}
