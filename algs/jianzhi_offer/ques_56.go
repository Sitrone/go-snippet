package jianzhi_offer

// 思路1：常规思路，使用map计数
// 进阶思路，可以参考：https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/solution/san-jin-zhi-zhuang-tai-ji-by-muyids/
func singleNumber(nums []int) int {
	m := make(map[int]int, len(nums)/3+1)
	for _, num := range nums {
		m[num]++
	}

	for k, v := range m {
		if v == 1 {
			return k
		}
	}

	return -1
}
