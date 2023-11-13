package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)

		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
				continue
			}

			// 上边界
			if i == 0 && j != 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
				continue
			}

			// 左边界
			if i != 0 && j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
				continue
			}

			dp[i][j] = min(dp[i][j-1], dp[i-1][j]) + grid[i][j]

		}
	}

	return dp[m-1][n-1]
}

func main() {
	g := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	res := minPathSum(g)
	fmt.Println(res)
}
