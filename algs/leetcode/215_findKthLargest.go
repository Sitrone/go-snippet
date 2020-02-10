package leetcode

import "container/heap"

// 思路1： 利用快排算法的思想切分 TODO 代码还有问题
// 思路2： 小顶堆
func findKthLargest(nums []int, k int) int {
	l, h := 0, len(nums)-1

	idx := len(nums) - k
	for l < h {
		p := partition(nums, l, h)
		if p == idx {
			return nums[p]
		} else if p > idx {
			h = p - 1
		} else {
			l = p + 1
		}
	}

	return -1
}

func partition(nums []int, l, h int) int {
	i, j := l, h+1
	v := nums[l]
	for {
		for i++; nums[i] < v; i++ {
			if i == h {
				break
			}
		}
		for j--; nums[j] > v; j-- {
			if j == l {
				break
			}
		}

		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			break
		}
	}

	nums[l], nums[j] = nums[j], nums[l]
	return j
}

func findKthLargest1(nums []int, k int) int {
	if k > len(nums) {
		return 0
	}

	h := &MinHeap{}
	heap.Init(h)
	for i := 0; i < len(nums); i++ {
		heap.Push(h, nums[i])
		if h.Len() > k {
			heap.Pop(h)
		}

	}
	return heap.Pop(h).(int)
}

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}
func (h MinHeap) Less(i, j int) bool {
	return h[i] <= h[j]
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := old.Len() - 1
	x := old[n]
	*h = old[0:n]
	return x
}
