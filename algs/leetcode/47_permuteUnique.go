package leetcode

import (
	"sort"
)

// 类似LeetCode 40题
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)

	ans := make([][]int, 0)

	var dfs func(pos int, cur []int, used map[int]struct{})
	dfs = func(pos int, cur []int, used map[int]struct{}) {
		if pos == len(nums) {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			ans = append(ans, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			// 已经使用的元素不再使用，剪枝
			if _, ok := used[i]; ok {
				continue
			}

			// 当当前元素和前一个元素值相同, 并且前一个元素被使用过，剪枝
			if i > 0 && nums[i] == nums[i-1] {
				if _, ok := used[i-1]; ok {
					continue
				}
			}

			cur = append(cur, nums[i])
			used[i] = struct{}{}
			dfs(pos+1, cur, used)
			cur = cur[:len(cur)-1]
			delete(used, i)
		}
	}

	cur := make([]int, 0, len(nums))
	used := make(map[int]struct{}, len(nums))
	dfs(0, cur, used)
	return ans
}
