package scramblestring

func CharCount(s string) map[rune]int {
	m := make(map[rune]int)
	for _, ch := range s {
		if cnt, exist := m[ch]; exist {
			m[ch] = cnt + 1
		} else {
			m[ch] = 1
		}
	}
	return m
}

func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	} else if len(s1) == 0 {
		return true
	} else if len(s1) == 1 {
		return s1 == s2
	} else if s1 == s2 {
		return true
	}

	m1 := CharCount(s1)
	m2 := CharCount(s2)

	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	l := len(s1)
	for i := 1; i < l; i++ {
		j := l - i
		isScramble := (isScramble(s1[0:i], s2[0:i]) &&
			isScramble(s1[i:l], s2[i:l])) ||
			(isScramble(s1[0:i], s2[j:l]) &&
				isScramble(s1[i:l], s2[0:j]))
		if isScramble {
			return true
		} else {
			continue
		}
	}
	return false
}
