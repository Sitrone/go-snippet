package snippet

var (
	nums = []int{1, 5, 10, 25, 50}

	ret int
)

//CounterChange
func CounterChange(total int) int {
	counterChange(0, total, total)
	return ret
}

func counterChange(i, remain, total int) {
	if i == len(nums)-1 && remain == 0 {
		ret += 1
		return
	}
	if i > len(nums)-1 || remain < 0 {
		return
	}

	counterChange(i, remain-nums[i], total) // 使用了第i种
	counterChange(i+1, remain, total)       // 没有使用第i种
}
