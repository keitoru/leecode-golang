package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func coinChange(coins []int, amount int) int {

	max := amount + 1
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		dp[i] = max

		for _, coin := range coins {
			if coin <= i {
				// dp[coin] = 1 所以 dp[i] = dp[i-coin]+1
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]

}

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	res := coinChange(coins, amount)
	fmt.Println(res)
}
