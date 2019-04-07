// https://leetcode.com/problems/palindrome-number/

package palindromenumber

import "strconv"

func isPalindrome(x int) bool {
	origin := strconv.Itoa(x)
	runes := []rune(origin)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes) == origin
}
