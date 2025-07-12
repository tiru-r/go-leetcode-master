package encode_and_decode_strings_271

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodecRoundTrip(t *testing.T) {
	tests := [][]string{
		nil,
		{},
		{""},
		{"", "", ""},
		{"a"},
		{"hello", "world"},
		{"should", "encode", "and", ""},
		{"63/Rc", "h", "BmI3FS~J9#vmk", "7uBZ?7*/", "24h+X", "O "},
		{"long" + strings.Repeat("x", 1e3), "short"},
	}

	for _, tc := range tests {
		name := strings.Join(tc, "|")
		t.Run(name, func(t *testing.T) {
			var c Codec
			enc := c.Encode(tc)
			dec := c.Decode(enc)
			assert.Equal(t, tc, dec)
		})
	}
}
