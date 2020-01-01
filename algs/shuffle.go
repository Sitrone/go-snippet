package algs

import "math/rand"

// ref rand.Perm(int)
type Shuffle struct {
	ori []int
}

func Constructor(nums []int) Shuffle {
	s := Shuffle{ori: nums}
	copy(s.ori, nums)
	return s
}

/** Resets the array to its original configuration and return it. */
func (s *Shuffle) Reset() []int {
	return s.ori
}

/** Returns a random shuffling of the array. */
func (s *Shuffle) Shuffle() []int {
	ori := make([]int, len(s.ori))
	copy(ori, s.ori)
	rd := make([]int, len(ori))
	for i := len(rd) - 1; i >= 0; i-- {
		pos := rand.Intn(i + 1)
		rd[i] = ori[pos]
		ori[pos], ori[i] = ori[i], ori[pos]
	}
	return rd
}
