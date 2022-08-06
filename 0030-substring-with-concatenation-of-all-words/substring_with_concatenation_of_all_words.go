package substring_with_concatenation_of_all_words

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}

	clen := len(words[0])
	if len(s) < clen*len(words) {
		return []int{}
	}

	candidates := make(map[string]int)

	// initial candidates
	for _, w := range words {
		v, exists := candidates[w]
		if exists {
			candidates[w] = v + 1
		} else {
			candidates[w] = 1
		}
	}

	ret := make([]int, 0)

	for i := 0; i <= len(s)-clen*len(words); i++ {
		sub := s[i : i+clen*len(words)]
		if isMatch(sub, candidates, clen, len(words)) {
			ret = append(ret, i)
		}
	}

	return ret
}

func isMatch(s string, candidates map[string]int, clen int, splits int) bool {
	subs := make(map[string]int)
	for i := 0; i < splits; i++ {
		begin := i * clen
		end := begin + clen

		sub := s[begin:end]

		candidateCount, exists := candidates[sub]

		if !exists {
			return false
		}

		v1, exists := subs[sub]

		if exists {
			if v1 >= candidateCount {
				return false
			}
			subs[sub] = v1 + 1
		} else {
			subs[sub] = 1
		}
	}

	return true
}
