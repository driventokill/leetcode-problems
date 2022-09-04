// https://leetcode.com/problems/partition-to-k-equal-sum-subsets/
package partitiontokequalsumsubsets

func canPartitionKSubsets(nums []int, k int) bool {
	if k > len(nums) {
		return false
	}

	var sum int
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}

	target := sum / k

	return byBucket(nums, k, target)
}

// ----------- traverse by bucket -----------------------------------

func byBucket(nums []int, k, target int) bool {
	used := make([]bool, len(nums))
	memo := make(map[uint16]bool)
	return backtrackByBucket(k, 0, nums, 0, used, target, memo)
}

// k: current bucket to be filled
func backtrackByBucket(k, bucket int, nums []int, start int, used []bool, target int, memo map[uint16]bool) bool {
	if k == 0 {
		// every bucket is fullfilled
		return true
	}

	if bucket == target {
		// current bucket is full
		// select nums for the next bucket
		res := backtrackByBucket(k-1, 0, nums, 0, used, target, memo)
		// remember current state
		memo[tobits(used)] = res
		return res
	}

	// use former result if exists
	if res, exists := memo[tobits(used)]; exists {
		return res
	}

	// find usable num from start
	for i := start; i < len(nums); i++ {
		if used[i] {
			continue
		}

		if nums[i]+bucket > target {
			continue
		}

		// make a decision
		used[i] = true
		bucket += nums[i]
		// decide if the next number should be filled
		if backtrackByBucket(k, bucket, nums, i+1, used, target, memo) {
			return true
		}

		// revert
		used[i] = false
		bucket -= nums[i]
	}

	return false
}

func tobits(used []bool) uint16 {
	var a uint16
	for i, v := range used {
		if v {
			a += 1 << i
		}
	}
	return a
}

// ----------- traverse by num --------------------------------------

func byNum(nums []int, k, target int) bool {
	buckets := make([]int, k)
	return backtrackByNum(buckets, nums, 0, target)
}

func backtrackByNum(sums []int, nums []int, index, target int) bool {
	if index == len(nums) {
		for i := range sums {
			if sums[i] != target {
				return false
			}
		}
		return true
	}

	for i := range sums {
		if sums[i]+nums[index] > target {
			continue
		}
		sums[i] += nums[index]
		if backtrackByNum(sums, nums, index+1, target) {
			return true
		}
		sums[i] -= nums[index]
	}

	return false
}
