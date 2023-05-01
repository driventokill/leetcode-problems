package sum

import "sort"

func threeSum1(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)

	var ret [][]int

	for i := 0; i < len(nums)-2; i++ {
		// Skip the same numbers
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		// Now take the two pointer approach to solve the linear array problem.
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] +nums[k]

			// If sum is greater than target means the larger
			// value(from right as nums is sorted i.e, k at right)
			// is taken and it is not able to sum up to the target
			if sum > 0 {
				k--
			} else if sum < 0 {
				j++
			} else {
				ret = append(ret, []int{nums[i], nums[j], nums[k]})
				j++
				// Skip the same numbers
				for nums[j-1] == nums[j] && j < k {
					j++
				}
			}
		}
	}

	return ret
}
