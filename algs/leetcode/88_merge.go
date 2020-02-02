package leetcode

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	var (
		p1 = m - 1
		p2 = n - 1
		p  = m + n - 1
	)

	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] >= nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}

	for p1 >= 0 {
		nums1[p] = nums1[p1]
		p--
		p1--
	}
	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p--
		p2--
	}
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	for i, j := m, 0; i < m+n && j < n; i++ {
		nums1[i] = nums2[j]
		j++
	}
	sort.Ints(nums1)
}
