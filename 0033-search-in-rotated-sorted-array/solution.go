package searchinrotatedsortedarray

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid:=left+(right-left)/2
		if nums[mid] == target {
			return mid
		}

		if nums[left] > nums[mid] {
			// the right side of mid is sorted
			if nums[mid] < target && target <= nums[right] {
				// search in the right side
				left = mid + 1
			} else {
				// search in the left side which is rotated sorted
				right = mid - 1
			}
		} else {
			// the left side of mid is sorted
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}

		}
	}

	return -1
}
