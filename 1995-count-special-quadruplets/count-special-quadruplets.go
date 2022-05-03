package countspecialquadruplets

func countQuadruplets(nums []int) int {
	var cnt int

	for i := len(nums) - 1; i >= 3; i-- {
		cnt = cnt + len(kIndexSum(nums, 0, i, nums[i], 3))
	}

	return cnt
}

// kIndexSum given a slice range [start, end) of ints, try to pick `k` indices which satisfied:
// nums[k[0]] + nums[k[1]] +...+ nums[k[k-1]] = target
func kIndexSum(nums []int, start int, end int, target int, k int) [][]int {
	if k == 1 {
		return eqs(nums, start, end, target)
	}

	if end-start < k {
		return [][]int{}
	}

	ret := make([][]int, 0)

	for i := end - 1; i >= start; i-- {
		if nums[i] > target {
			continue
		}

		nextRet := kIndexSum(nums, start, i, target-nums[i], k-1)

		for _, r := range nextRet {
			ret = append(ret, append(r, i))
		}
	}

	return ret
}

func eqs(nums []int, start int, end int, target int) [][]int {
	ret := make([][]int, 0)

	if start > end {
		return [][]int{}
	}

	for i := start; i < end; i++ {
		if nums[i] == target {
			ret = append(ret, []int{i})
		}
	}

	return ret
}

// func twoIndexSum(nums []int, start int, end int, target int) [][]int {
// 	if end-start < 2 {
// 		return [][]int{}
// 	}

// 	i := start
// 	j := end - 1

// 	ret := make([][]int, 0)

// 	moveI := false

// 	for i < j {
// 		if nums[i]+nums[j] == target {
// 			ret = append(ret, []int{i, j})
// 		}

// 		if moveI {
// 			resetI := i
// 			for nums[i+1] == nums[i] && i+1 < j {
// 				ret = append(ret, []int{i + 1, j})
// 				i++
// 			}
// 			i = resetI
// 			j--
// 		} else {
// 			resetJ := j
// 			for nums[j-1] == nums[j] && i < j-1 {
// 				ret = append(ret, []int{i, j - 1})
// 				j--
// 			}
// 			j = resetJ
// 			i++
// 		}
// 		moveI = !moveI
// 	}

// 	return ret
// }
