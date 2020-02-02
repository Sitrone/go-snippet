package heap

import "errors"

var (
	EmptyErr = errors.New("heap is empty")
)

// 大顶堆，数组实现，可以自动扩容缩容，eles[0]不存储元素
type MaxHeap struct {
	n    int
	eles []int
}

type IHeap interface {
	Insert(key int)
	Size() int
	Top() (int, error)
	DelTop() (int, error)
}

func NewMaxHeap(cap int) *MaxHeap {
	return &MaxHeap{
		n:    0,
		eles: make([]int, cap+1),
	}
}

func (h *MaxHeap) IsEmpty() bool {
	return h.Size() == 0
}

func (h *MaxHeap) isFull() bool {
	return h.Size() == cap(h.eles)-1
}

func (h *MaxHeap) Insert(key int) {
	if h.isFull() {
		h.resize(h.n + h.n/2)
	}
	h.n++
	h.eles[h.n] = key

	h.swin(h.n)
}

func (h *MaxHeap) Size() int {
	return h.n
}

func (h *MaxHeap) Top() (int, error) {
	if h.IsEmpty() {
		return 0, EmptyErr
	}

	return h.eles[1], nil
}

// 交换top和最后一个元素，top元素下沉
func (h *MaxHeap) DelTop() (int, error) {
	if h.IsEmpty() {
		return 0, EmptyErr
	}

	top := h.eles[1]
	h.eles[1], h.eles[h.n] = h.eles[h.n], h.eles[1]
	h.n--
	if h.n > 0 && h.n < h.n/4 {
		h.resize(h.n / 2)
	}

	h.sink(1)

	return top, nil
}

// 堆下沉
func (h *MaxHeap) sink(k int) {
	for k*2 < h.n {
		j := k * 2
		// 找到两个子节点的最大值
		if j < h.n && h.eles[j] < h.eles[j+1] {
			j++
		}
		// 如果大于两个子节点的最大值，则交换
		if !(h.eles[k] < h.eles[j]) {
			break
		}

		h.eles[k], h.eles[j] = h.eles[j], h.eles[k]
		k = j
	}
}

// 堆上浮
func (h *MaxHeap) swin(k int) {
	for k > 1 && h.eles[k/2] < h.eles[k] {
		h.eles[k/2], h.eles[k] = h.eles[k], h.eles[k/2]
		k /= 2
	}
}

func (h *MaxHeap) resize(len int) {
	tmp := make([]int, len+1)
	for i := 1; i < h.n+1; i++ {
		tmp[i] = h.eles[i]
	}

	h.eles = tmp
}
