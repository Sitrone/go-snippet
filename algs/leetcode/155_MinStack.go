package leetcode

type MinStack struct {
	eles    []int
	minNums []int
}

// 思路：双栈，一个存储正常的值，一个存储最小值（有数据冗余）
/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		eles:    make([]int, 0),
		minNums: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.eles = append(this.eles, x)

	if len(this.minNums) == 0 {
		this.minNums = append(this.minNums, x)
	} else {
		if x > this.GetMin() {
			this.minNums = append(this.minNums, this.GetMin())
		} else {
			this.minNums = append(this.minNums, x)
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.eles) == 0 {
		return
	}

	this.eles = this.eles[:len(this.eles)-1]
	this.minNums = this.minNums[:len(this.minNums)-1]
}

func (this *MinStack) Top() int {
	if len(this.eles) == 0 {
		return -1
	}

	return this.eles[len(this.eles)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.minNums) == 0 {
		return -1
	}

	return this.minNums[len(this.minNums)-1]
}
