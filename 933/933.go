package main

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{}

}

func (this *RecentCounter) Ping(t int) int {
	this.queue = append(this.queue, t)
	for len(this.queue) > 0 && this.queue[0] < t-3000 {
		this.queue = this.queue[1:]
	}
	return len(this.queue)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
