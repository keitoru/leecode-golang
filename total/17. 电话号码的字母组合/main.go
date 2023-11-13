package main

func letterCombinations(digits string) []string {

	letterMap := []string{" ", "*", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	var res []string

	if len(digits) == 0 {
		return res
	}

	var iterStr func(str string, letter []byte, index int)
	iterStr = func(str string, letter []byte, index int) {
		if index == len(str) {
			res = append(res, string(letter))
			return
		}

		c := []byte(str)[index]
		pos := c - '0'

		mapString := []byte(letterMap[pos])
		for i := 0; i < len(mapString); i++ {
			letter = append(letter, mapString[i])

			iterStr(str, letter, index+1)

			letter = letter[:len(letter)-1]
		}
	}

	iterStr(digits, []byte{}, 0)

	return res
}
