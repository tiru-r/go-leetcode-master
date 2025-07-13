package excel_sheet_column_number_171

func titleToNumber(s string) int {
	result := 0
	for i := range len(s) {
		result = result*26 + int(s[i]-'A') + 1
	}
	return result
}
