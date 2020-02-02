package leetcode

// 思路1: 尽量分配原则，每次分配查询已经有的分配的组是否满员，如果没有则加入，否则新增group，时间复杂度O(n^2),空间复杂度O(n)
func groupThePeople(groupSizes []int) [][]int {
	var (
		ret    [][]int
		exists bool
	)
	for i, size := range groupSizes {
		exists = false
		for j, group := range ret {
			if cap(group) == size && len(group) < size {
				var tmp = group
				tmp = append(tmp, i)
				ret[j] = tmp
				exists = true
				goto OUTER
			}
		}

	OUTER:
		if !exists {
			group := make([]int, 0, size)
			group = append(group, i)
			ret = append(ret, group)
		}
	}

	return ret
}

// 思路2：参考官方答案，先粗分组，再细分组
func groupThePeople2(groupSizes []int) [][]int {
	var m = make(map[int][]int, len(groupSizes)/2)
	for i, size := range groupSizes {
		if g, ok := m[size]; ok {
			g = append(g, i)
			m[size] = g
		} else {
			group := make([]int, 0, 10)
			group = append(group, i)
			m[size] = group
		}
	}

	var ret [][]int
	for size, group := range m {
		tmp := group
		for i := 0; i < len(tmp); i += size {
			ret = append(ret, tmp[i:i+size])
		}
	}

	return ret
}
