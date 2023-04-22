package _167

func TwoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	left, right := 0, len(nums)-1

	for left < right {
		sum := nums[left] + nums[right]
		if sum > target {
			right = right - 1
		} else if sum < target {
			left = left + 1
		} else {
			return []int{nums[left], nums[right]}
		}
	}

	return []int{}
}
