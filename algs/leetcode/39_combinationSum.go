package leetcode

func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	cur := make([]int, 0, 2)

	largerThanAny := func(s int) bool {
		for _, candidate := range candidates {
			if s >= candidate {
				return true
			}
		}
		return false
	}

	var dfs = func(reminder, start int) {}
	dfs = func(reminder, start int) {
		if reminder == 0 {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			ans = append(ans, tmp)
			return
		}

		if reminder > 0 && largerThanAny(reminder) {
			for i := start; i < len(candidates); i++ {
				cur = append(cur, candidates[i])
				dfs(reminder-candidates[i], i) // 可以重复则从当前i开始，不能重复则从i+1开始
				cur = cur[:len(cur)-1]
			}
		}
	}

	dfs(target, 0)
	return ans
}
