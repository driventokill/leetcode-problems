package numberofwaystobuypencilsandpens

func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	var solutions int64
	for pens := 0; pens*cost1 <= total; pens++ {
		remain := total - pens*cost1
		solutions += int64(remain/cost2) + 1
	}

	return solutions
}
