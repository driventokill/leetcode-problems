// https://leetcode.com/problems/trapping-rain-water/

package trappingrainwater

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	drops := 0
	lmax := make([]int, len(height))

	lmax[0] = height[0]

	for i := 0; i < len(height)-1; i++ {
		lmax[i+1] = max(height[i], lmax[i])
	}

	rmax := 0
	for i := len(height) - 1; i > 0; i-- {
		rmax = max(rmax, height[i])
		h := min(rmax, lmax[i])
		if h > height[i] {
			drops += h - height[i]
		}
	}
	return drops
}
