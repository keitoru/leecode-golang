package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	N := len(s)

	dp := make([]bool, N+1)
	dp[0] = true

	mWordDict := make(map[string]struct{}, N)
	for _, s := range wordDict {
		mWordDict[s] = struct{}{}
	}

	for i := 1; i <= N; i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[i] {
				break
			}

			if dp[j] == false {
				continue
			}

			suffix := s[j:i]
			if _, ok := mWordDict[suffix]; ok && dp[j] {
				dp[i] = true
				break
			}
		}
	}

	//fmt.Println(dp)

	return dp[N]
}

func main() {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	res := wordBreak(s, wordDict)
	fmt.Println(res)
}
