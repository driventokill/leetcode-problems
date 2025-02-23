package maximumscorefromremovingsubstrings

// maximumGain will try to use greedy approach.
// If there is a substring only contains 'a' and 'b', it can always be removed to
// then end with a substring of empty or only 'a' or 'b'. The idea is to count the
// 'a' 'b' pairs inside the substring. The key point is to consider which kind of
// substring if greater, 'ab' or 'ba'. It will guide us to remove which substring
// first.
func maximumGain(s string, x int, y int) int {
	// a is the start of the greater substring.
	a, b := 'a', 'b'

	// 'ba' is greater than 'ab'.
	if x < y {
		x, y = y, x
		a, b = b, a
	}

	var cnt1, cnt2, ans int
	for _, c := range s {
		if c == a {
			cnt1++
		} else if c == b {
			if cnt1 > 0 {
				// Remove the greater substring and gain the point immediately
				ans += x
				cnt1--
			} else {
				cnt2++
			}
		} else {
			// Remove the remaining lesser substrings
			ans += min(cnt1, cnt2) * y
			cnt1 = 0
			cnt2 = 0
		}
	}
	ans += min(cnt1, cnt2) * y
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
