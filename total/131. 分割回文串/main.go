package main

import "fmt"

func partition(s string) [][]string {
	var res [][]string

	var dfs func(s string, path []string)
	dfs = func(s string, path []string) {
		if len(s) == 0 {
			res = append(res, path)
			return
		}

		for i := 1; i < len(s); i++ {
			pre := s[0:i]
			if isPalindrome(pre) {
				path = append(path, pre)
				dfs(s[i:], path)
				path = path[:len(path)-1]
			}
		}
	}

	dfs(s, []string{})
	return res
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	start, end := 0, len(s)-1
	for start <= end {
		if s[start] != s[end] {
			return false
		}

		start++
		end--
	}

	return true
}

func main() {
	s := "aab"
	res := partition(s)
	fmt.Println(res)
}
