package jianzhi_offer

// 递归
func Fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return Fib(n-1) + Fib(n-2)
}

// dp
func Fib1(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	var dp = make([]int, n+1)
	dp[0], dp[1] = 0, 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// dp 优化
func Fib2(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}

	return a
}

// 矩阵乘法

//type IntStack []int
//
//func (s *IntStack) Push(val int) {
//	*s = append(*s, val)
//}
//
//func (s *IntStack) Pop() int {
//	ans := (*s)[len(*s)-1]
//	*s = (*s)[:len(*s)-1]
//	return ans
//}
//
//// 用两个栈实现队列
//type CQueue struct {
//	st1 *IntStack
//	st2 *IntStack
//}
//
//func Constructor() CQueue {
//	return CQueue{
//		st1: &IntStack{},
//		st2: &IntStack{},
//	}
//}
//
//func (this *CQueue) AppendTail(value int) {
//	this.st1.Push(value)
//}
//
//func (this *CQueue) DeleteHead() int {
//	if len(*this.st2) == 0 {
//		if len(*this.st1) == 0 {
//			return -1
//		}
//
//		for len(*this.st1) > 0 {
//			this.st2.Push(this.st1.Pop())
//		}
//	}
//
//	return this.st2.Pop()
//}

// 双100%
type CQueue struct {
	st1 []int
	st2 []int
}

func Constructor() CQueue {
	return CQueue{
		st1: make([]int, 0, 10),
		st2: make([]int, 0, 10),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.st1 = append(this.st1, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.st2) == 0 {
		if len(this.st1) == 0 {
			return -1
		}

		for i := len(this.st1) - 1; i >= 0; i-- {
			this.st2 = append(this.st2, this.st1[i])
		}

		this.st1 = this.st1[:0] // clear stack1
	}

	ans := this.st2[len(this.st2)-1]
	this.st2 = this.st2[:len(this.st2)-1]

	return ans
}
