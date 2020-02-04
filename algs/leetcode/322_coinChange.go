package leetcode

// 递归
func coinChange(coins []int, amount int) int {
	if len(coins) == 0 || amount == 0 {
		return 0
	}
	if amount < 0 {
		return 0
	}

	min := func(x, y int) int {
		if x > y {
			return y
		}

		return x
	}

	max := amount + 1
	var dp = make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = max
	}

	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	} else {
		return dp[amount]
	}

	//// 贪心解法，有时候会求不出最优解
	//sort.Ints(coins)
	//
	//var total int
	//for amount > 0 {
	//	var i = len(coins) - 1
	//	for ; i >= 0; i-- {
	//		if amount >= coins[i] {
	//			break
	//		}
	//	}
	//
	//	if i >= 0 {
	//		total++
	//		amount -= coins[i]
	//	} else {
	//		return -1
	//	}
	//}
	//
	//return total
}
