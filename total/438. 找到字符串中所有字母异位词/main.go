package main

import "fmt"

func findAnagrams(s string, p string) []int {
	need := make(map[byte]int)
	window := make(map[byte]int)

	for _, c := range []byte(p) {
		need[c]++
	}

	left, right, valid := 0, 0, 0

	res := make([]int, 0, len(s))

	for right < len(s) {
		// c 是将移入窗口的字符
		c := s[right]
		// 右移窗口
		right++
		// 进行窗口内数据的一系列更新
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		/*** debug 输出的位置 ***/
		//fmt.Printf("window:[%d, %d]\n", left, right)
		/********************/

		// 判断左侧窗口是否要收缩
		for right-left >= len(p) {
			if valid == len(need) {
				res = append(res, left)
			}
			d := s[left]
			// 左移窗口
			left++
			// 进行窗口内数据的一系列更新
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	return res
}

func main() {
	res := findAnagrams("cbaebabacd", "cba")
	fmt.Println(res)
}
