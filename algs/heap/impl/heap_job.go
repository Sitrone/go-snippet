package impl

import "time"

type Job struct {
	job       func(data interface{})
	data      interface{}
	timestamp time.Time // unix 时间戳
}

func (j *Job) RunTime() time.Time {
	return j.timestamp
}

type JobHeap []*Job

func (h JobHeap) Len() int {
	return len(h)
}

func (h JobHeap) Less(i, j int) bool {
	return h[i].timestamp.Sub(h[j].RunTime()) < 0
}

func (h JobHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *JobHeap) Push(x interface{}) {
	*h = append(*h, x.(*Job))
}

func (h *JobHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]  // 返回删除的元素
	*h = (*h)[:n-1] // [n:m]不包括下标为m的元素
	return x
}
