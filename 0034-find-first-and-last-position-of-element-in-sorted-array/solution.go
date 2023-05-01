package findfirstandlastpositionofelementinsortedarray

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	left, right := 0, len(nums)-1

	// 1. Find the right border
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] >= target && (mid == len(nums)-1 || nums[mid+1] > target) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	// 2. Can't find right border
	if nums[left] != target {
		return []int{-1, -1}
	}

	rightBorderIdx := left

	left, right = 0, rightBorderIdx

	// 3. Find the left border
	for left < right {
		mid := left + (right-left)/2

		if nums[mid] >= target && (mid == 0 || nums[mid-1] <= target) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return []int{left, rightBorderIdx}
}
