package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	N := len(temperatures)

	res := make([]int, N)
	var st []int

	for i := 0; i < N; i++ {
		for len(st) > 0 && temperatures[i] > temperatures[st[len(st)-1]] {
			fmt.Println("stack:", st)
			t := st[len(st)-1]
			st = st[:len(st)-1]
			res[t] = i - t
		}

		st = append(st, i)

	}

	return res
}

func main() {
	t := []int{73, 74, 75, 71, 69, 72, 76, 73}
	res := dailyTemperatures(t)
	fmt.Println(res)
}
