package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	vals := make([]int, 0, 100)

	curNode := head
	for curNode != nil {
		vals = append(vals, curNode.Val)
		curNode = curNode.Next
	}

	front, back := 0, len(vals)-1
	for front < back {
		if vals[front] != vals[back] {
			return false
		}

		front++
		back--
	}

	return true
}

func main() {

}
