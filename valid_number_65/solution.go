package valid_number_65

// Package valid provides a fast, zero-allocation validator for determining
// whether a byte slice represents a valid number as defined by LeetCode 65
// "Valid Number".
//
// A valid number consists of:
//   - An optional leading sign ('+' or '-')
//   - Either an integer part (e.g., "123") or a decimal part (e.g., "12.34", ".34", "12.")
//   - An optional exponent ('e' or 'E') followed by a signed integer (e.g., "1e10", "1E-10")
//
// Examples of valid numbers:
//   "2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "6e-1"
//
// Examples of invalid numbers:
//   "abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"
//
// The implementation uses a finite state machine (FSM) for efficient parsing
// with zero allocations, operating directly on the input byte slice.

// IsNumber reports whether s is a valid number according to LeetCode 65.
func IsNumber(s []byte) bool {
	// Handle edge cases: nil or empty input
	if len(s) == 0 {
		return false
	}

	// FSM states
	const (
		startState         = iota // Initial state
		signState                 // After leading '+' or '-'
		integerState              // Digits before decimal point
		leadingDotState           // '.' with no digits before
		fractionState             // Digits after decimal point
		trailingDotState          // Digits followed by '.' (e.g., "42.")
		exponentState             // After 'e' or 'E'
		exponentSignState         // After exponent sign
		exponentDigitState        // Digits in exponent
	)

	// Accepting states
	accept := func(state int) bool {
		return state == integerState || state == fractionState ||
			state == trailingDotState || state == exponentDigitState
	}

	// Character constants
	const (
		dot           = '.'
		exponentLower = 'e'
		exponentUpper = 'E'
		plus          = '+'
		minus         = '-'
	)

	state := startState
	for i := 0; i < len(s); i++ {
		ch := s[i]
		switch state {
		case startState:
			switch {
			case ch == plus || ch == minus:
				state = signState
			case ch == dot:
				state = leadingDotState
			case isDigit(ch):
				state = integerState
			default:
				return false
			}

		case signState:
			switch {
			case ch == dot:
				state = leadingDotState
			case isDigit(ch):
				state = integerState
			default:
				return false
			}

		case integerState:
			switch {
			case isDigit(ch):
				// Stay in integerState
			case ch == dot:
				state = trailingDotState
			case ch == exponentLower || ch == exponentUpper:
				state = exponentState
			default:
				return false
			}

		case leadingDotState:
			if isDigit(ch) {
				state = fractionState
			} else {
				return false
			}

		case fractionState:
			if ch == exponentLower || ch == exponentUpper {
				state = exponentState
			} else if !isDigit(ch) {
				return false
			}

		case trailingDotState:
			if ch == exponentLower || ch == exponentUpper {
				state = exponentState
			} else if !isDigit(ch) {
				return false
			}

		case exponentState:
			switch {
			case ch == plus || ch == minus:
				state = exponentSignState
			case isDigit(ch):
				state = exponentDigitState
			default:
				return false
			}

		case exponentSignState:
			if isDigit(ch) {
				state = exponentDigitState
			} else {
				return false
			}

		case exponentDigitState:
			if !isDigit(ch) {
				return false
			}
		}
	}

	return accept(state)
}

// isDigit checks if a byte is a digit (0-9).
func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}
