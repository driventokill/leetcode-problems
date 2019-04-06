// https://leetcode.com/problems/string-to-integer-atoi

package atoi

func myAtoi(str string) int {
	const WS byte = ' '
	var leadingSpace = true
	var n = 0
	var neg = false
	const IntMax = 1<<31 - 1
	const IntMin = -1 << 31
	for i := 0; i < len(str); i++ {
		if leadingSpace {
			if WS == str[i] {
				continue
			} else if '+' == str[i] {
				leadingSpace = false
				continue
			} else if '-' == str[i] {
				neg = true
				leadingSpace = false
				continue
			}
		}
		leadingSpace = false

		ch := str[i] - '0'
		if ch > 9 || ch < 0 {
			return n
		} else if neg {
			if IntMin+int(ch) < n*10 {
				n = n*10 - int(ch)
			} else {
				return IntMin
			}
		} else {
			if IntMax-(n*10)-int(ch) > 0 {
				n = n*10 + int(ch)
			} else {
				return IntMax
			}
		}
	}
	return n
}
