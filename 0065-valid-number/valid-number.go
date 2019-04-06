// https://leetcode.com/problems/valid-number/

package validnumber

type ValidNum struct {
	symSet   bool
	valueSet bool
	float    bool
	floatSet bool
}

type ValidExp struct {
	is  bool
	exp ValidNum
}

func (n *ValidNum) valid() bool {
	if n.float {
		return n.floatSet || n.valueSet
	} else {
		return n.valueSet
	}
}

func (n *ValidNum) initial() bool {
	return !(n.valueSet || n.floatSet || n.symSet || n.float)
}

func strip(s string) string {
	prefixIdx := 0
	subIdx := len(s) - 1
	for i, ch := range s {
		if ch == ' ' {
			continue
		} else {
			prefixIdx = i
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			continue
		} else {
			subIdx = i
			break
		}
	}
	return s[prefixIdx : subIdx+1]
}

func isNumber(s string) bool {
	validNum := ValidNum{false, false, false, false}
	validExp := ValidExp{false, ValidNum{false, false, false, false}}

	for _, ch := range strip(s) {
		v := ch - '0'

		actualN := &validNum
		if validExp.is {
			actualN = &validExp.exp
		}

		if ch == ' ' {
			return false
		} else if ch == '-' && actualN.initial() {
			actualN.symSet = true
			continue
		} else if ch == '+' && actualN.initial() {
			actualN.symSet = true
			continue
		} else if ch == 'e' && !validExp.is {
			validExp.is = true
			continue
		} else if ch == '.' && !validExp.is && !actualN.float {
			actualN.float = true
			continue
		} else if v >= 0 && v <= 9 {
			if actualN.float {
				actualN.floatSet = true
			} else {
				actualN.valueSet = true
			}
			continue
		} else {
			return false
		}

	}
	if validNum.valid() {
		if validExp.is {
			return validExp.exp.valueSet
		} else {
			return true
		}
	} else {
		return false
	}

}
