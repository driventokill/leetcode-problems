package permutations

func permute(nums []int) [][]int {
	track := make([]int, 0)
	used := make([]bool, len(nums))

	res := make([][]int, 0)
	backtrack(&res, nums, track, used)

	return res
}

func backtrack(res *[][]int, nums []int, track []int, used []bool) {
	if len(track) == len(nums) {
		d := make([]int, len(track))
		copy(d, track)
		*res = append(*res, d)
		return
	}

	for i := range nums {
		if used[i] {
			continue
		}
		track = append(track, nums[i])
		used[i] = true
		backtrack(res, nums, track, used)
		track = track[0 : len(track)-1]
		used[i] = false
	}
}
