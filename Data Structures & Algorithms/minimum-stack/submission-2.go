type MinStack struct {
	val []int
	min []int 
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.val = append(this.val, val)

	if len(this.min) == 0 || val < this.min[len(this.min)-1] {
		this.min = append(this.min, val)
	} else {
		this.min = append(this.min, this.min[len(this.min)-1])
	}
}

func (this *MinStack) Pop() {
	top := len(this.val) - 1
	this.val = this.val[:top]
	this.min = this.min[:top]
}

func (this *MinStack) Top() int {
	return this.val[len(this.val)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}