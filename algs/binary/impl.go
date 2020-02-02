package binary

// 注意：数据量太小不适合二分查找，数据量太大也不适合二分查找。

// 二分查找，基本场景，查找等于给定值的元素
// 变体一：查找第一个值等于给定值的元素
// 变体二：查找最后一个值等于给定值的元素
// 变体三：查找第一个大于等于给定值的元素
// 变体四：查找最后一个小于等于给定值的元素

// 二分查找，基本场景，查找等于给定值的元素, 元素有序且没有重复元素
func bsearch(arr []int, val int) int {
	low, high := 0, len(arr)-1
	var mid int
	for low <= high {
		mid = low + (high-low)/2
		if val == arr[mid] {
			return mid
		} else if val > arr[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return low
}

// 变体一：查找第一个值等于给定值的元素
func bsearch1(arr []int, val int) int {
	low, high := 0, len(arr)-1
	var mid int
	for low <= high {
		mid = low + (high-low)/2
		if arr[mid] > val {
			high = mid - 1
		} else if arr[mid] < val {
			low = mid + 1
		} else {
			// key point
			if (mid == 0) || arr[mid-1] != val {
				return mid
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}

// 变体二：查找最后一个值等于给定值的元素
func bsearch2(arr []int, val int) int {
	low, high := 0, len(arr)-1
	var mid int
	for low <= high {
		mid = low + (high-low)/2
		if arr[mid] > val {
			high = mid - 1
		} else if arr[mid] < val {
			low = mid + 1
		} else {
			// key point
			if (mid == len(arr)-1) || arr[mid+1] != val {
				return mid
			} else {
				low = mid + 1
			}
		}
	}

	return -1
}

// 变体三：查找第一个大于等于给定值的元素
func bsearch3(arr []int, val int) int {
	low, high := 0, len(arr)-1
	var mid int
	for low <= high {
		mid = low + (high-low)/2
		if arr[mid] >= val {
			if mid == 0 || arr[mid-1] < val {
				return mid
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}

	return -1
}

// 变体四：查找最后一个小于等于给定值的元素
func bsearch4(arr []int, val int) int {
	low, high := 0, len(arr)-1
	var mid int
	for low <= high {
		mid = low + (high-low)/2
		if arr[mid] > val {
			high = mid - 1
		} else {
			if mid == len(arr)-1 || arr[mid+1] > val {
				return mid
			} else {
				low = mid + 1
			}
		}
	}

	return -1
}
