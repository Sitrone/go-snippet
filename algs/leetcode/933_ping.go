package leetcode

type RecentCounter struct {
	Ele []int
}

//func Constructor() RecentCounter {
//	return RecentCounter{Ele: make([]int, 0, 10)}
//}

func (this *RecentCounter) Ping(t int) int {
	var min = t - 3000

	this.Ele = append(this.Ele, t)
	var i = len(this.Ele) - 1
	for ; i >= 0; i-- {
		if this.Ele[i] < min {
			break
		}
	}

	this.Ele = this.Ele[i+1:]

	return len(this.Ele)
}
