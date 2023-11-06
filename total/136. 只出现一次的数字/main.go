package main

// 异或，只能用来判断一个出现一次的数字
func singleNumber(nums []int) int {
	x := 0
	for _, num := range nums {
		x ^= num
	}

	return x
}

func main() {

}
