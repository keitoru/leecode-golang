package main

func longestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	left, right, length := 0, 0, 1
	maxStart, maxLen := 0, 0

	for i := 0; i < n; i++ {
		left = i - 1
		right = i + 1

		for left >= 0 && s[left] == s[i] {
			length++
			left--
		}

		for right < n && s[right] == s[i] {
			length++
			right++
		}

		for left >= 0 && right < n && s[right] == s[left] {
			length += 2
			left--
			right++
		}

		if length > maxLen {
			maxLen = length
			maxStart = left
		}

		length = 1
	}

	return s[maxStart+1 : maxStart+maxLen+1]

}
