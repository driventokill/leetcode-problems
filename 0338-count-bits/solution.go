package count_bits

func countBits1(n int) []int {
	if n == 0 {
		return []int{0}
	}

	ans := make([]int, n+1)

	var bits []bool
	var o_count int

	for i := 0; i <= n; i++ {
		bits,o_count = bin_inc(bits, o_count)
		ans[i] = o_count
	}

	return ans
}

func count_bin(bits []bool) int{
	var ret int
	for _, v := range bits {
		if v {
			ret++
		}
	}

	return ret
}

func bin_inc(origin []bool, o_count int) ([]bool, int) {
	if len(origin) == 0 {
		return []bool{false}, 0
	}

	var enlarge = false

	for i, v := range origin {
		if !v {
			origin[i] = true
			o_count++
			break
		}
		origin[i] = false
		o_count--

		if i == len(origin) - 1{
			enlarge = true
		}
	}

	if enlarge {
		return append(origin, true), o_count+1
	}

	return origin, o_count
}

func countBits(n int) []int {
	if n == 0 {
		return []int{0}
	}

	ans := make([]int, n+1)

	for i := 0; i <= n; i++ {
		ans[i] = count(toBits(i))
	}

	return ans
}

func toBits(n int) []int {
	var bits []int
	remain := n
	for {
		old := remain
		remain = old / 2
		bits = append(bits, old - remain * 2)
		if remain < 2 {
			bits = append(bits, remain)
			break
		}
	}

	return bits
}

func count(ints []int) int {
	var ret int
	for _,v := range ints {
		if v > 0 {
			ret++
		}
	}

	return ret
}
