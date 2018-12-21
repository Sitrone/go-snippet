package algs

import "math/rand"

type Shuffle struct {
	ori []int
}

func Constructor(nums []int) Shuffle {
	s := Shuffle{ori: nums}
	copy(s.ori, nums)
	return s
}

/** Resets the array to its original configuration and return it. */
func (this *Shuffle) Reset() []int {
	return this.ori
}

/** Returns a random shuffling of the array. */
func (this *Shuffle) Shuffle() []int {
	ori := make([]int, len(this.ori))
	copy(ori, this.ori)
	rd := make([]int, len(ori))
	for i := len(rd) - 1; i >= 0; i-- {
		pos := rand.Intn(i + 1)
		rd[i] = ori[pos]
		ori[pos], ori[i] = ori[i], ori[pos]
	}
	return rd
}
