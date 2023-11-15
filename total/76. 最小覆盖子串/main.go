package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	need := map[byte]int{}
	// 每个字符所需的个数
	for _, c := range []byte(t) {
		if _, ok := need[c]; !ok {
			need[c] = 0
		}
		need[c]++
	}
	// 所需字符的个数
	needCnt := len(t)
	minBegin, minLen := 0, 0

	left := 0
	for right, c := range []byte(s) {
		if need[c] > 0 {
			needCnt--
		}
		need[c]--

		// 窗口里包含t的所有元素
		if needCnt == 0 {
			// 左指针移动，缩小滑动窗口，直到碰到一个必须包含的元素
			for need[s[left]] < 0 {
				need[s[left]]++
				left++
			}

			// 当前窗口大小
			length := right - left + 1
			if minLen == 0 || length < minLen {
				minBegin = left
				minLen = length
			}

			// 让左指针在增加一个位置，开始寻找下一个满足条件的滑动窗口
			need[s[left]]++
			needCnt++
			left++
		}
	}

	return s[minBegin : minBegin+minLen]

}

func minWindow1(s string, t string) string {
	need := make(map[byte]int)
	window := make(map[byte]int)

	for _, c := range []byte(t) {
		need[c]++
	}

	left, right, valid := 0, 0, 0
	start, length := 0, math.MaxInt

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
		fmt.Printf("window:[%d, %d]\n", left, right)
		/********************/

		// 判断左侧窗口是否要收缩
		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
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

	if length == math.MaxInt {
		return ""
	}
	return s[start : start+length]
}

func main() {
	res := minWindow("EBBANCF", "ABC")
	fmt.Println(res)
}
