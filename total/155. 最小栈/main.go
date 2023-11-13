package main

import "math"

// 辅助栈
type MinStack struct {
	xStack   []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		minStack: []int{math.MaxInt},
	}
}

func (this *MinStack) Push(val int) {
	this.xStack = append(this.xStack, val)

	minNum := this.minStack[len(this.minStack)-1]
	this.minStack = append(this.minStack, min(minNum, val))
}

func (this *MinStack) Pop() {
	this.xStack = this.xStack[:len(this.xStack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	if len(this.xStack) == 0 {
		return 0
	}

	return this.xStack[len(this.xStack)-1]

}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
