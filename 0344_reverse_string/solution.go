package _344_reverse_string

func ReverseString(s []byte) []byte {
	if len(s) == 0 {
		return s
	}

	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}

	return s
}
