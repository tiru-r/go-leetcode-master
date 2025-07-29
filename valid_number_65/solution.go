package valid_number_65

// isNumber validates if a string represents a valid number using single-pass parsing.
// Valid formats: integer, decimal, scientific notation (e.g., "123", "3.14", "1e10", "-2.5E-3")
func isNumber(input string) bool {
	if len(input) == 0 {
		return false
	}

	// Manual trimming for better performance than strings.TrimSpace
	start, end := 0, len(input)-1
	for start <= end && isWhitespace(input[start]) {
		start++
	}
	for end >= start && isWhitespace(input[end]) {
		end--
	}

	if start > end {
		return false // Only whitespace
	}

	position := start
	length := end + 1

	// Parse optional leading sign
	if position < length && isSignCharacter(input[position]) {
		position++
	}

	// State tracking for number components
	hasMainDigits := false
	hasDecimalPoint := false

	// Parse main number part (digits with optional decimal point)
	for position < length {
		currentChar := input[position]
		switch {
		case isDigit(currentChar):
			hasMainDigits = true
			position++
		case currentChar == '.':
			if hasDecimalPoint {
				return false // Multiple decimal points
			}
			hasDecimalPoint = true
			position++
		default:
			break // End of main number part
		}
	}

	// Must have at least one digit in main part
	if !hasMainDigits {
		return false
	}

	// Parse optional exponent part
	if position < length && isExponentCharacter(input[position]) {
		position++

		// Optional exponent sign
		if position < length && isSignCharacter(input[position]) {
			position++
		}

		// Exponent must have digits
		hasExponentDigits := false
		for position < length && isDigit(input[position]) {
			hasExponentDigits = true
			position++
		}

		if !hasExponentDigits {
			return false // Exponent without digits
		}
	}

	// All characters must be consumed
	return position == length
}

// Helper functions for character classification

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func isSignCharacter(char byte) bool {
	return char == '+' || char == '-'
}

func isExponentCharacter(char byte) bool {
	return char == 'e' || char == 'E'
}

func isWhitespace(char byte) bool {
	return char == ' ' || char == '\t' || char == '\n' || char == '\r'
}
