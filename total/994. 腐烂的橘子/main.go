package main

import "fmt"

// BFS 广度 借助队列
func orangesRotting(grid [][]int) int {
	M, N := len(grid), len(grid[0])
	var queue [][]int

	count := 0 // count 表示新鲜橘子的数量
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == 1 {
				count++
			} else if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			}
		}
	}

	//fmt.Println(queue)

	round := 0
	for count > 0 && len(queue) > 0 {
		round++
		n := len(queue)
		for i := 0; i < n; i++ {
			orange := queue[0]
			queue = queue[1:]

			r, c := orange[0], orange[1]

			if r-1 >= 0 && grid[r-1][c] == 1 {
				grid[r-1][c] = 2
				count--
				queue = append(queue, []int{r - 1, c})
			}

			if r+1 < M && grid[r+1][c] == 1 {
				grid[r+1][c] = 2
				count--
				queue = append(queue, []int{r + 1, c})
			}

			if c-1 >= 0 && grid[r][c-1] == 1 {
				grid[r][c-1] = 2
				count--
				queue = append(queue, []int{r, c - 1})
			}

			if c+1 < N && grid[r][c+1] == 1 {
				grid[r][c+1] = 2
				count--
				queue = append(queue, []int{r, c + 1})
			}

		}
	}

	if count > 0 {
		return -1
	}

	return round
}

func main() {
	g := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	res := orangesRotting(g)
	fmt.Println(res)
}
