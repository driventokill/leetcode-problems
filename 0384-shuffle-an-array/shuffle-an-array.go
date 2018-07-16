package shuffleAnArray

import (
	"math/rand"
)

// Store the origin and shuffed
type Solution struct {
	origin   []int
	shuffled []int
}

// Construct a solution
func Constructor(nums []int) Solution {
	origin := make([]int, len(nums))
	shuffled := make([]int, len(nums))

	copy(origin, nums)
	copy(shuffled, nums)

	return Solution{
		origin:   nums,
		shuffled: shuffled,
	}
}

// Return the original array
func (Solution *Solution) Reset() []int {
	return Solution.origin
}

// Shuffle and return the shuffed array
func (Solution *Solution) Shuffle() []int {

	size := len(Solution.shuffled)
	for i := 1; i < size+1; i++ {
		n := rand.Intn(i)
		Solution.shuffled[i-1], Solution.shuffled[n] = Solution.shuffled[n], Solution.shuffled[i-1]
	}

	return Solution.shuffled
}
