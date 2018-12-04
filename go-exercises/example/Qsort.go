package example

func Qsort(arr []int) {
	doSort(arr, 0, len(arr)-1)
}

func doSort(arr []int, left int, right int) {
	if left > right {
		return
	}
	p := partition(arr, left, right)
	doSort(arr, left, p-1)
	doSort(arr, p+1, right)
}

func partition(nums []int, left int, right int) int {
	pivot := nums[left]

	i := left
	j := right
	for i < j {
		for nums[j] >= pivot && i < j {
			j--
		}
		for nums[i] <= pivot && i < j {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[left], nums[j] = nums[j], pivot
	return j

}
