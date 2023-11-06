package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func numSquares(n int) int {
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = i

		for j := 1; i-j*j >= 0; j++ {
			// 因为dp[j * j] = dp[0] + 1 = 1
			// 所以dp[i] = 1 + dp[i - j * j]
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}

	return dp[n]
}

func main() {
	res := numSquares(14)
	fmt.Println(res)
}
