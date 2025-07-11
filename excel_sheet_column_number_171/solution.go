package excel_sheet_column_number_171

// Modern solution avoiding floating-point math.Pow

// Modern implementation using integer arithmetic instead of floating-point
func titleToNumber(s string) int {
	var res int
	powerOf26 := 1
	for i := len(s) - 1; i >= 0; i-- {
		charNum := int(s[i]-'A') + 1
		res += charNum * powerOf26
		powerOf26 *= 26
	}
	return res
}
