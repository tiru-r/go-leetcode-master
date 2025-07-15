package reorder_data_in_log_files_937

import (
	"cmp"
	"slices"
	"strings"
)

func reorderLogFiles(logs []string) []string {
	slices.SortFunc(logs, compare)
	return logs
}

func compare(a, b string) int {
	aId, aContent, _ := strings.Cut(a, " ")
	bId, bContent, _ := strings.Cut(b, " ")

	aIsDigit := aContent[0] >= '0' && aContent[0] <= '9'
	bIsDigit := bContent[0] >= '0' && bContent[0] <= '9'

	switch {
	case aIsDigit && !bIsDigit:
		return 1
	case !aIsDigit && bIsDigit:
		return -1
	case aIsDigit && bIsDigit:
		return 0
	default:
		return cmp.Or(
			strings.Compare(aContent, bContent),
			strings.Compare(aId, bId),
		)
	}
}
