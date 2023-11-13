package main

func exist(board [][]byte, word string) bool {
	words := []byte(word)
	N := len(board)
	M := len(board[0])

	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		if i >= N || i < 0 || j >= M || j < 0 || board[i][j] != word[k] {
			return false
		}

		if k == len(words)-1 {
			return true
		}

		board[i][j] = 0
		res := dfs(i+1, j, k+1) || dfs(i-1, j, k+1) ||
			dfs(i, j+1, k+1) || dfs(i, j-1, k+1)

		board[i][j] = word[k]
		return res
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}
