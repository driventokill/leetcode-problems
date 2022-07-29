package combinationsum

import "sort"

// combinationSum
func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}

	sort.Ints(candidates)

	nums, res := []int{}, [][]int{}
	findCombinationSum(candidates, 0, target, nums, &res)

	return res
}

func findCombinationSum(candidates []int, index int, target int, nums []int, res *[][]int) {
	if target == 0 {
		n := make([]int, len(nums))
		copy(n, nums)
		*res = append(*res, n)
		return
	}

	c := candidates[index]

	if c > target {
		return
	}

	if target >= c {
		findCombinationSum(candidates, index, target-c, append(nums, c), res)
	}

	if index < len(candidates)-1 {
		findCombinationSum(candidates, index+1, target, nums, res)
	}
}
