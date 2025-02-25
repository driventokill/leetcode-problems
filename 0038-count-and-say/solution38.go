package countandsay

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	var lastChar = rune('0')
	var cnt = 0
	var ret []rune
	for _, c := range countAndSay(n - 1) {
		if lastChar == rune('0') {
			lastChar = c
			cnt++
			continue
		}

		if c == lastChar {
			cnt++
			continue
		}

		ret = append(ret, rune('0'+cnt), lastChar)

		cnt = 1
		lastChar = c
	}

	ret = append(ret, rune('0'+cnt), lastChar)
	return string(ret)
}
