package main

func findDuplicate(nums []int) int {

	slow := nums[0]
	fast := nums[nums[0]]

	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	// 找入口
	pre1, pre2 := 0, slow
	for pre1 != pre2 {
		pre1 = nums[pre1]
		pre2 = nums[pre2]
	}

	return pre1

}
