type MinStack struct {
	val []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.val = append(this.val, val)
}

func (this *MinStack) Pop() {
	top := len(this.val) - 1
	this.val = this.val[:top]
}

func (this *MinStack) Top() int {
	topVal := this.val[len(this.val)-1]
	return topVal
}

func (this *MinStack) GetMin() int {
	currMin := 1 << 32
	for _, num := range this.val {
		if num < currMin {
			currMin = num
		}
	}
	return currMin
}
