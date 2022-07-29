package backtrack

func Permute(nums []int) [][]int {

	res := [][]int{}

	permute0(nums, &res)

	return res
}

func permute0(nums []int, res *[][]int) {
	track := []int{}
	used := make([]bool, len(nums))
	backtrack(nums, track, used, res)
}

// backtrack
// 路径：记录在track中
// 选择列表： nums中不存在于 track 的那些元素
// 结束条件：nums 中的元素全部在track中出现
func backtrack(nums []int, track []int, used []bool, res *[][]int) {
	if len(track) == len(nums) {
		path := make([]int, len(track))
		copy(path, track)
		*res = append(*res, path)
		return
	}

	for i, v := range nums {
		if used[i] {
			continue
		}

		oTrack := track
		track = append(track, v)
		used[i] = true
		// 进入下一层决策树
		backtrack(nums, track, used, res)
		// 取消选择
		track = oTrack
		used[i] = false
	}
}
