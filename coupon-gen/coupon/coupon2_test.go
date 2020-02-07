package coupon

import (
	"fmt"
	"sort"
	"testing"
)

func TestCoupon2Data_Gen(t *testing.T) {
	coupon := NewData2()
	coupon.Alphabet = []rune("023456789ABCDEFGHJKMNOPQRSTUVWXYZ")
	//coupon.Alphabet = []rune("0123456789")
	coupon.Len = 6
	ret, _ := coupon.Gen(500_000)

	for i := 0; i < 100; i++ {
		fmt.Printf("%s ", <-ret)
	}
	fmt.Println()
}

func TestCoupon2Data_Gen2(t *testing.T) {
	var (
		total int64 = 10
		m           = make(map[int64]int64, total)
		ans         = make([]int, 0, total)
	)

	for ; total > 5; total-- {
		n := rnd.Int63n(total)
		vn, ok1 := m[n]
		tn, ok2 := m[total-1]

		switch {
		case ok1 && ok2:
			ans = append(ans, int(vn))
			m[n], m[total-1] = m[total-1], m[n]
		case !ok1 && ok2:
			ans = append(ans, int(n))
			m[n] = tn
		case ok1 && !ok2:
			ans = append(ans, int(vn))
			m[n] = total - 1
		case !ok1 && !ok2:
			ans = append(ans, int(n))
			m[n] = total - 1
		}
	}

	fmt.Println(m)
	fmt.Println(ans)
	sort.Ints(ans)
	fmt.Println(ans)
}

func TestCoupon2Data_Gen3(t *testing.T) {
	var nums = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	ansChan := randomGet(nums)
	for i := 0; i < len(nums); i++ {
		fmt.Printf("%d ", <-ansChan)
	}
	fmt.Println()
}
