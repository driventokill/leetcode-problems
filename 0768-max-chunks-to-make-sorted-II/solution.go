package maxchunkstomakesortesii

import "fmt"

func maxChunksToSorted(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	var cnt int = 0
	maxOfLeft := make([]int, len(arr))
	minOfRight := make([]int, len(arr))

	for i := range minOfRight {
		minOfRight[i] = 100000000 + 1
	}

	for i := range arr {
		if i > 0 {
			maxOfLeft[i] = maxOfLeft[i-1]
		}
		if maxOfLeft[i] <= arr[i] {
			maxOfLeft[i] = arr[i]
		}
	}

	for i := len(arr) - 1; i > 0; i-- {
		minOfRight[i-1] = minOfRight[i]
		if arr[i] < minOfRight[i-1] {
			minOfRight[i-1] = arr[i]
		}
	}

	fmt.Printf("maxOfLeft: %v, minOfRight: %v\n", maxOfLeft, minOfRight)

	for i := range arr {
		if maxOfLeft[i] <= minOfRight[i] {
			cnt++
		}
	}

	return cnt
}
