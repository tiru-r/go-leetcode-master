package excel_sheet_column_number_171

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTitleToNumber(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"single A", "A", 1},
		{"two-letter AB", "AB", 28},
		{"two-letter ZY", "ZY", 701},
		{"two-letter YB", "YB", 652},
		{"three-letter AAA", "AAA", 703},
		{"edge Z", "Z", 26},
		{"edge AZ", "AZ", 52},
		{"edge BZ", "BZ", 78},
		{"long AZY", "AZY", 1378},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, titleToNumber(tt.s))
		})
	}
}
