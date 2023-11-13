package main

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	var dfs func(lRemain, rRemain int, path []byte)
	dfs = func(lRemain, rRemain int, path []byte) {
		if 2*n == len(path) {
			ans = append(ans, string(path))
			return
		}
		if lRemain > 0 {
			dfs(lRemain-1, rRemain, append(path, '('))
		}
		if rRemain > lRemain {
			dfs(lRemain, rRemain-1, append(path, ')'))
		}
	}

	dfs(n, n, []byte{})
	return ans
}
