package main

import "math"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxProfit(prices []int) int {
	maxNum := math.MinInt
	in := math.MaxInt

	for _, price := range prices {
		if price < in {
			in = price
		}

		maxNum = max(maxNum, price-in)
	}

	return maxNum
}
