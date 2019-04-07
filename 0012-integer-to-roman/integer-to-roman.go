// https://leetcode.com/problems/integer-to-roman/
package integertoroman

import "strings"

func intToRomanSym(n int, x string, v string, i string) string {
	if n > 0 {
		var sb strings.Builder
		if n > 8 {
			sb.WriteString(i)
			sb.WriteString(x)
		} else if n > 5 {
			sb.WriteString(v)
			for idx := 0; idx < n-5; idx++ {
				sb.WriteString(i)
			}
		} else if n > 3 {
			for idx := 0; idx < 5-n; idx++ {
				sb.WriteString(i)
			}
			sb.WriteString(v)
		} else {
			for idx := 0; idx < n; idx++ {
				sb.WriteString(i)
			}
		}
		return sb.String()
	} else {
		return ""
	}
}

func intToRoman(num int) string {
	var sb strings.Builder

	n := num

	if x := n / 1000; x > 0 {
		for idx := 0; idx < x; idx++ {
			sb.WriteString("M")
		}
		n -= 1000 * x
	}

	if x := n / 100; x > 0 {
		sb.WriteString(intToRomanSym(x, "M", "D", "C"))
		n -= x * 100
	}

	if x := n / 10; x > 0 {
		sb.WriteString(intToRomanSym(x, "C", "L", "X"))
		n -= x * 10
	}

	if n > 0 {
		sb.WriteString(intToRomanSym(n, "X", "V", "I"))
	}

	return sb.String()
}
