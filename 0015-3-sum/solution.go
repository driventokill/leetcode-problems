package sum

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	return nSums(nums, 3, 0)
}

func nSums(nums []int, n int, target int) [][]int {
	if n == 2 {
		return twoSums(nums, target)
	}

	// n > 2
	ret := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] != nums[i-1] {
			sub := nSums(nums[i+1:], n-1, target-nums[i])
			for _, s := range sub {
				ret = append(ret, append(s, nums[i]))
			}
		}
	}
	return ret
}

func twoSums(nums []int, target int) [][]int {
	ret := make([][]int, 0)
	m := make(map[int]int)
	rem := make(map[int]bool)
	for _, n := range nums {
		if _, ok := m[n]; ok {
			if _, ok := rem[n]; !ok {
				ret = append(ret, []int{m[n], n})
				rem[n] = true
			}
		} else {
			m[target-n] = n
		}
	}
	return ret
}
