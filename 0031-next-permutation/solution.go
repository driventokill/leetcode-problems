package nextpermutation

func nextPermutation(nums []int) {
	size := len(nums)
	left := size - 2
	right := size - 1

	for left >= 0 {
		if nums[left] >= nums[left+1] {
			left--
			continue
		}

		for nums[left] >= nums[right] {
			right--
		}

		nums[left], nums[right] = nums[right], nums[left]
		break
	}

	// sort.Ints(nums[(left+1):])
	left = left + 1
	right = size - 1

	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
