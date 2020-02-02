package leetcode

// 从右往左遍历，标记max
func replaceElements(arr []int) []int {
	var (
		ans = make([]int, len(arr))
		max = -1
	)
	ans[len(arr)-1] = -1
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i+1] > max {
			max = arr[i+1]
		}
		ans[i] = max
	}

	return ans
}

func replaceElements1(arr []int) []int {
	var ans = make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		var max = -1
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > max {
				max = arr[j]
			}
		}

		ans[i] = max
	}

	return ans
}
