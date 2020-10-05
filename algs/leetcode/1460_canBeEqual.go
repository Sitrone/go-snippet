package leetcode

func canBeEqual(target []int, arr []int) bool {
	if len(target) != len(arr) {
		return false
	}

	if len(target) == 0 {
		return true
	}

	targetM := make(map[int]int, len(target))
	for i := 0; i < len(target); i++ {
		targetM[target[i]]++
	}

	for i := 0; i < len(arr); i++ {
		if val, ok := targetM[arr[i]]; !ok || val <= 0 {
			return false
		} else {
			targetM[arr[i]] = val - 1
		}
	}

	return true
}
