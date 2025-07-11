package reorder_data_in_log_files_937

import (
	"cmp"
	"slices"
	"strings"
)

// Modern solution using slices.SortFunc for efficient sorting
func reorderLogFiles(logs []string) []string {
	// Use slices.SortFunc with the existing comparison function
	slices.SortFunc(logs, compareLogs)
	return logs
}

func compareLogs(a, b string) int {
	// Use strings.Cut for cleaner parsing instead of Split + TrimPrefix
	aId, aContent, _ := strings.Cut(a, " ")
	bId, bContent, _ := strings.Cut(b, " ")

	// Use direct byte comparison for better performance
	aIsDigit := aContent[0] >= '0' && aContent[0] <= '9'
	bIsDigit := bContent[0] >= '0' && bContent[0] <= '9'

	// Letter logs should come before digit logs
	if aIsDigit != bIsDigit {
		if aIsDigit {
			return 1 // a is digit, b is letter, so b comes first
		}
		return -1 // a is letter, b is digit, so a comes first
	}

	// Both are digit logs - preserve original order
	if aIsDigit && bIsDigit {
		return 0
	}

	// Both are letter logs - compare by content, then by identifier using cmp.Or
	return cmp.Or(
		strings.Compare(aContent, bContent),
		strings.Compare(aId, bId),
	)
}
