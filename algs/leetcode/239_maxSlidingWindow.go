package leetcode

// https://leetcode-cn.com/problems/sliding-window-maximum/

// 思路1 暴力法
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k > len(nums) {
		return nil
	}

	var maxEle = func(arr []int, l, h int) int {
		var m = arr[l]
		for i := l + 1; i <= h; i++ {
			if arr[i] > m {
				m = arr[i]
			}
		}
		return m
	}

	ans := make([]int, 0, len(nums)+1-k)
	start, end := 0, k-1

	for end < len(nums) {
		ans = append(ans, maxEle(nums, start, end))
		start++
		end++
	}
	return ans
}

// 思路2 双端队列 TODO
// 思路3 堆+遍历
func maxSlidingWindow1(nums []int, k int) []int {
	if len(nums) == 0 || k > len(nums) {
		return nil
	}

	var deque = make([]int, 0, k)
	for i := 1; i < k; i++ {
		for len(deque) > 0 && nums[i] > nums[deque[len(deque)-1]] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
	}
	// fmt.Println(deque)
	var res []int
	res = append(res, nums[deque[0]])
	for i := k; i < len(nums); i++ {
		if i-deque[0] == k {
			deque = deque[1:len(deque)]
		}
		for len(deque) > 0 && nums[i] > nums[deque[len(deque)-1]] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
		// fmt.Println(deque)
		res = append(res, nums[deque[0]])
	}
	return res
}
