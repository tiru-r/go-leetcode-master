package encode_and_decode_strings_271

import (
	"strconv"
	"strings"
)

type Codec struct {
}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	// Calculate precise capacity for optimal performance
	totalLen := 0
	for _, s := range strs {
		totalLen += len(s) + len(strconv.Itoa(len(s))) + 1 // content + length + separator
	}
	
	var builder strings.Builder
	builder.Grow(totalLen)
	
	for _, s := range strs {
		builder.WriteString(strconv.Itoa(len(s)))
		builder.WriteByte('_')
		builder.WriteString(s)
	}
	return builder.String()
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	var result []string
	i := 0
	
	for i < len(strs) {
		// Use strings.Cut for cleaner parsing
		lengthStr, rest, found := strings.Cut(strs[i:], "_")
		if !found {
			break
		}
		
		// Parse the length
		length, err := strconv.Atoi(lengthStr)
		if err != nil || length > len(rest) {
			break
		}
		
		// Extract the string using length
		result = append(result, rest[:length])
		i += len(lengthStr) + 1 + length // length + underscore + content
	}
	
	return result
}
