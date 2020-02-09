package leetcode

import "sort"

func subsetsWithDup(nums []int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(nums)

	var backtrace func(pos, num int, cur []int)
	backtrace = func(pos, num int, cur []int) {
		if len(cur) == num {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			ans = append(ans, tmp)
			return
		}

		for i := pos; i < len(nums); i++ {
			// 前面那个重复元素已经使用，则不再使用，在78题的基础上面剪枝
			if i > pos && nums[i] == nums[i-1] {
				continue
			}

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
