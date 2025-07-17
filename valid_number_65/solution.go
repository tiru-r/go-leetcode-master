package valid_number_65

func isNumber(s string) bool {
	if len(s) == 0 {
		return false
	}
	
	i := 0
	
	for i < len(s) && s[i] == ' ' {
		i++
	}
	
	if i == len(s) {
		return false
	}
	
	if i < len(s) && (s[i] == '+' || s[i] == '-') {
		i++
	}
	
	hasDigits, hasDot := false, false
	
	for i < len(s) && (isDigit(s[i]) || s[i] == '.') {
		if s[i] == '.' {
			if hasDot {
				return false
			}
			hasDot = true
		} else {
			hasDigits = true
		}
		i++
	}
	
	if !hasDigits {
		return false
	}
	
	if i < len(s) && (s[i] == 'e' || s[i] == 'E') {
		i++
		
		if i < len(s) && (s[i] == '+' || s[i] == '-') {
			i++
		}
		
		hasExponentDigits := false
		for i < len(s) && isDigit(s[i]) {
			hasExponentDigits = true
			i++
		}
		
		if !hasExponentDigits {
			return false
		}
	}
	
	for i < len(s) && s[i] == ' ' {
		i++
	}
	
	return i == len(s)
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
