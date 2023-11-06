package main

import (
	"fmt"
	"strconv"
	"strings"
)

func decodeString(s string) string {
	stack := make([]string, 0, len(s))
	res := ""
	multi := 0

	for _, c := range []byte(s) {
		if string(c) == "[" {
			stack = append(stack, []string{strconv.Itoa(multi), res}...)
			res, multi = "", 0
		} else if string(c) == "]" {
			last_res := stack[len(stack)-1:][0]
			stack = stack[:len(stack)-1]

			cur_multi := stack[len(stack)-1:][0]
			stack = stack[:len(stack)-1]

			nums, _ := strconv.Atoi(cur_multi)
			res = last_res + strings.Repeat(res, nums)

		} else if string(c) >= "0" && string(c) <= "9" {
			num, _ := strconv.Atoi(string(c))
			multi = multi*10 + num
		} else {
			res += string(c)
		}

	}

	return res
}

func main() {
	res := decodeString("2[abc]3[cd]ef")
	fmt.Println(res)
}
