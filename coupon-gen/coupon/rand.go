package coupon

import (
	crand "crypto/rand"
	"encoding/binary"
	"log"
	"math/rand"
)

var (
	srand = rand.New(&cryptoSource{})
)

type cryptoSource struct{}

func (s cryptoSource) Seed(seed int64) {}

func (s cryptoSource) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func (s cryptoSource) Uint64() (v uint64) {
	err := binary.Read(crand.Reader, binary.BigEndian, &v)
	if err != nil {
		log.Fatal(err)
	}
	return v
}

// 从给定的大数组中每次随机选择一个元素，每次选择的元素不允许重复，要求时间复杂度O(1)
// 思路：以空间换时间，用map记录已经交换的数据
// 对于从大的数据范围取随机数，可以使用，如果给定的数据源范围不大，可以直接用洗牌算法洗牌然后取数，代码更加简单易读
func randomGet(nums []int) chan int {
	var ans = make(chan int)

	go func() {
		defer close(ans)

		total := len(nums)
		m := make(map[int]int, total)

		for ; total > 0; total-- {
			n := rnd.Intn(total)
			vn, ok1 := m[n]
			tn, ok2 := m[total-1]

			switch {
			case ok1 && ok2:
				ans <- vn
				m[n], m[total-1] = m[total-1], m[n]
			case !ok1 && ok2:
				ans <- nums[n]
				m[n] = tn
			case ok1 && !ok2:
				ans <- vn
				m[n] = total - 1
			case !ok1 && !ok2:
				ans <- nums[n]
				m[n] = total - 1
			}
		}
	}()

	return ans
}
