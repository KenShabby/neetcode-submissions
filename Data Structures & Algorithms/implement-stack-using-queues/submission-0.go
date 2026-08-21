type MyStack struct {
	q1 []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.q1 = append(this.q1, x)
}

func (this *MyStack) Pop() int {
	tail := len(this.q1) - 1
	temp := this.q1[tail] 
	this.q1 = this.q1[:tail]
	return temp
}

func (this *MyStack) Top() int {
	return this.q1[len(this.q1) - 1]
}

func (this *MyStack) Empty() bool {
	if len(this.q1) == 0 {
		return true
	} else {
		return false
	}
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Top();
 * param4 := obj.Empty();
 */
