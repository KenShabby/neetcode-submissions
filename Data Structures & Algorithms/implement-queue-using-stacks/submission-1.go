type MyQueue struct {
	Queue []int
}

func Constructor() MyQueue {
	return MyQueue{Queue: nil}
}

func (this *MyQueue) Push(x int) {
	this.Queue = append(this.Queue, x)
}

func (this *MyQueue) Pop() int {
	temp := this.Queue[0]
	this.Queue = this.Queue[1:]
	return temp
}

func (this *MyQueue) Peek() int {
	return this.Queue[0]
}

func (this *MyQueue) Empty() bool {
	if len(this.Queue) == 0 {
		return true
	} else {
		return false
	}
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Peek();
 * param4 := obj.Empty();
 */
