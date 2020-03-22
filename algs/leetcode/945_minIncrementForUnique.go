package leetcode

import "sort"

func minIncrementForUnique(A []int) int {
	if len(A) < 2 {
		return 0
	}

	sort.Ints(A)

	var move int
	for i := 1; i < len(A); i++ {
		// 遍历数组，若当前元素小于等于它的前一个元素，则将其变为前一个数+1
		if A[i] <= A[i-1] {
			pre := A[i]
			A[i] = A[i-1] + 1
			move += A[i] - pre
		}
	}

	return move
}

func minIncrementForUnique1(A []int) int {
	if len(A) < 2 {
		return 0
	}

	arr := make([]int, 80001)
	for i := 0; i < len(A); i++ {
		arr[A[i]]++
	}

	var ans, taken int
	for i := 0; i < len(arr); i++ {
		if arr[i] >= 2 {
			taken += arr[i] - 1
			ans -= i * (arr[i] - 1)
		} else if taken > 0 && arr[i] == 0 {
			taken--
			ans += i
		}

	}

	return ans
}
