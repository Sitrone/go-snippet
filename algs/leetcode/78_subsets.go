package leetcode

// 思路，可以看成从集合从取出i个不同的元素的取法其中 i>=0 && i<= len(nums)
func subsets(nums []int) [][]int {
	ans := make([][]int, 0)

	var backtrace func(pos, num int, cur []int)
	backtrace = func(pos, num int, cur []int) {
		if len(cur) == num {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			ans = append(ans, tmp)
			return
		}

		for i := pos; i < len(nums); i++ {
			cur = append(cur, nums[i])
			backtrace(i+1, num, cur)
			cur = cur[:len(cur)-1]
		}
	}

	for i := 0; i <= len(nums); i++ {
		cur := make([]int, 0, i)
		backtrace(0, i, cur)
	}

	return ans
}

// 动态规划:
// 初始状态：dp[0]=[[]];
// dp[i+1]={dp[i],[dp[i],nums[i-1]]}
// FIXME 参考解法
func subsets1(nums []int) [][]int {
	result := make([][]int, 0)
	result = append(result, []int{})

	//遍历nums中的数，每遍历到一个就将该值依次添加到前面的所有子集中，形成新的子集
	for i := 0; i < len(nums); i++ {
		l := len(result)
		for j := 0; j < l; j++ {
			tmp := make([]int, 0)
			tmp = append(tmp, result[j]...)
			tmp = append(tmp, nums[i])
			result = append(result, tmp)
		}
	}
	return result
}
