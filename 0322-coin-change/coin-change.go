// Package _322_coin_change https://leetcode-cn.com/problems/coin-change/
package _322_coin_change

import "math"

func coinChange(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}

	if amount == 0 {
		return 0
	}

	result := math.MaxInt32

	for _, c := range coins {
		sub := coinChange(coins, amount-c)

		// 无解
		if sub == -1 {
			continue
		}

		if result > sub+1 {
			result = sub + 1
		}
	}

	if result != math.MaxInt32 {
		return result
	} else {
		// 无解
		return -1
	}
}

func coinChange1(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		dp[i] = amount + 1 // 初始化一个不可能的数值，如果无法凑满就是这个数值
	}

	dp[0] = 0 //凑0只需要0个硬币

	for i := 0; i < len(dp); i++ {
		for _, c := range coins {
			if i < c { // 无解
				continue
			}

			if dp[i] > 1+dp[i-c] { // 找到1个更优的组合
				dp[i] = 1 + dp[i-c]
			}
		}
	}

	if dp[amount] == amount+1 { // 无解
		return -1
	} else {
		return dp[amount]
	}
}


// coins : [2,5]
// amount : 10
// dp: [0, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11]
// dp: [0, 11,  1, 11,  2, 11,  3,  2,  4, 11, 5]


