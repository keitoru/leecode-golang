package sliderWindow

import "fmt"

// 现在开始套模板，只需要思考以下四个问题：
//
//1、当移动 right 扩大窗口，即加入字符时，应该更新哪些数据？
//
//2、什么条件下，窗口应该暂停扩大，开始移动 left 缩小窗口？
//
//3、当移动 left 缩小窗口，即移出字符时，应该更新哪些数据？
//
//4、我们要的结果应该在扩大窗口时还是缩小窗口时进行更新？

// 滑动窗口框架
func slidingWindow(s string, t string) {
	need := make(map[byte]int)
	window := make(map[byte]int)

	for _, c := range []byte(t) {
		need[c]++
	}

	left, right, valid := 0, 0, 0

	for right < len(s) {
		// c 是将移入窗口的字符
		c := s[right]
		// 右移窗口
		right++
		// 进行窗口内数据的一系列更新
		//...

		/*** debug 输出的位置 ***/
		fmt.Printf("window:[%d, %d]\n", left, right)
		/********************/

		// 判断左侧窗口是否要收缩
		for {
			d := s[left]
			// 左移窗口
			left++
			// 进行窗口内数据的一系列更新
			// ...
		}
	}
}
