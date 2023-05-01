package searchinsertposition

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > target {
			right = mid
		} else if nums[mid] == target {
			return mid
		} else {
			left = mid + 1
		}
	}

	if right == len(nums) - 1 && nums[right] < target {
		return right + 1
	}

	return right
}
