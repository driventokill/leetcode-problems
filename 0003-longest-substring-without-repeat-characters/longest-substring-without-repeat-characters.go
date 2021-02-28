package longestsubstringwithoutrepeatcharacters

// 解法一：位图
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	var bitSet [256]bool

	// result 记录最大的字符串长度
	// left、right 分别表示 substring 的左右指针长度
	result, left, right := 0, 0, 0

	for left < len(s) {
		// 右侧字符对应的 bitSet 为 true，说明此字符在之前已经被 right 指针扫过，发生了重复，需要左侧指针路过一次以后才会将该位置标记为 false 消除重复
		if bitSet[s[right]] {
			bitSet[s[left]] = false
			left++
		} else {
			bitSet[s[right]] = true
			right++
		}

		if result < right-left {
			result = right - left
		}

		if left+result >= len(s) || right >= len(s) {
			break
		}
	}

	return result
}

// 解法二：滑动窗口
func lengthOfLongestSubstring1(s string) int {
	if len(s) == 0 {
		return 0
	}

	var freq [256]int
	result, left, right := 0, 0, -1

	for left < len(s) {
		// 右侧指针的字符出现次数为0，说明当前子串没有重复字符
		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
			freq[s[right+1]-'a']++
			right++
		} else {
			freq[s[left]-'a']--
			left++
		}
		result = max(result, right-left+1)
	}

	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}
