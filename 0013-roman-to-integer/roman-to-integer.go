// https://leetcode.com/problems/roman-to-integer/

package romantointeger

func romanToInt(s string) int {
	sym := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	if len(s) == 0 {
		return 0
	}

	lastValue := sym[s[0]]
	result := lastValue
	for i := 1; i < len(s); i++ {
		value := sym[s[i]]
		if value > lastValue {
			result += value - lastValue*2
		} else {
			result += value
		}
		lastValue = value
	}
	return result
}
