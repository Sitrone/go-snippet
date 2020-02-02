package leetcode

func findNumbers(nums []int) int {
	var total int
	for _, num := range nums {
		if hasEvenDigit(num) {
			total++
		}
	}
	return total
}

func hasEvenDigit(num int) bool {
	var count int
	for num != 0 {
		num = num / 10
		count++
	}
	return (count & 0x01) == 0
}
