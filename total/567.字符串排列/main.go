package main

import "fmt"

func checkInclusion(t string, s string) bool {
	need := make(map[byte]int)
	window := make(map[byte]int)

	for _, c := range []byte(t) {
		need[c]++
	}

	//fmt.Println(need)

	left, right, valid := 0, 0, 0

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
		//fmt.Println(window)
		//fmt.Printf("window:[%d, %d] %d\n", left, right, valid)
		/********************/

		// 判断左侧窗口是否要收缩
		for right-left >= len(t) {
			if valid == len(need) {
				return true
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

	return false
}

func main() {
	res := checkInclusion("ab", "eidbeaooo")
	fmt.Println(res)
}
