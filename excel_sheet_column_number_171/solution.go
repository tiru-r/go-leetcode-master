package excel_sheet_column_number_171

// titleToNumber converts an Excel column name to its 1-based index.
func titleToNumber(s string) int {
	res := 0
	for _, ch := range s {
		res = res*26 + int(ch-'A') + 1
	}
	return res
}
