package leetcode

import "sort"

// 奇数位递增，奇数位前一半大小的数字，偶数位模拟填充
func deckRevealedIncreasing(deck []int) []int {
	sort.Ints(deck)
	mid := len(deck) / 2

	var ret = make([]int, len(deck))
	for i, j := 0, 0; i < len(deck) && j <= mid; i += 2 {
		ret[i] = deck[j]
		j++
	}

	// 填充标记位
	for i := 1; i < len(deck); i += 2 {
		ret[i] = i
	}

	// 模拟填充
	tmp := make([]int, len(deck))
	copy(tmp, ret)
	for i := 0; len(tmp) != 0; i++ {
		if tmp[0] != deck[i] {
			ret[tmp[0]] = deck[i]
		}

		if len(tmp) == 1 {
			break
		}

		if len(tmp) >= 2 {
			toLast := tmp[1]
			tmp = tmp[2:]
			tmp = append(tmp, toLast)
		}
	}

	return ret
}

func deckRevealedIncreasing1(deck []int) []int {
	sort.Ints(deck)

	var indexs = make([]int, len(deck))
	for i := 0; i < len(deck); i++ {
		indexs[i] = i
	}

	var (
		ans = make([]int, len(deck))
		idx int
	)
	for _, v := range deck {
		idx, indexs = indexs[0], indexs[1:]
		ans[idx] = v
		if len(indexs) != 0 {
			first := indexs[0]
			indexs = append(indexs[1:], first)
		}
	}

	return ans
}
