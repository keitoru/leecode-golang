package main

func maxSlidingWindow(nums []int, k int) []int {

	n := len(nums)
	if n < 2 {
		return nums
	}

	queue := []int{}

	res := make([]int, n-k+1)

	for i := 0; i < n; i++ {
		// 当前遍历的数比队尾的值大，则需要弹出队尾值，直到队列重新满足从大到小的要求
		for len(queue) != 0 && nums[queue[len(queue)-1]] <= nums[i] {
			queue = queue[:len(queue)-1]
		}

		// 数组下标入队
		queue = append(queue, i)
		// 如果队首的数组下标不在[L,R]中，则需要弹出队首的值
		if queue[0] <= i-k {
			queue = queue[1:]
		}

		// i-k+1=0时窗口形成
		if i+1 >= k {
			res[i+1-k] = nums[queue[0]]
		}

	}

	return res
}
