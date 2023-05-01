package _005_longest_palindrome

func LongestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	for i := len(s); i > 0; i-- {
		for j := 0; j <= len(s)-i; j++ {
			if isPalindromic(s[j : j+i]) {
				return s[j : j+i]
			}
		}
	}

	return ""
}

func isPalindromic(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}
