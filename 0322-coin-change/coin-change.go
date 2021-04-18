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
		dp[i] = amount + 1
	}

	dp[0] = 0

	for i := 0; i < len(dp); i++ {
		for _, c := range coins {
			if i < c { // 无解
				continue
			}

			if dp[i] > 1 + dp[i - c] {
				dp[i] = 1 + dp[i-c]
			}
		}
	}

	if dp[amount] == amount + 1 { // 无解
		return -1
	} else {
		return dp[amount]
	}
}
