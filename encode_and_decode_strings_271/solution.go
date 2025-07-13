package encode_and_decode_strings_271

import (
	"strconv"
	"strings"
)

type Codec struct{}

func (c *Codec) Encode(strs []string) string {
	var buf strings.Builder
	for _, s := range strs {
		buf.WriteString(strconv.Itoa(len(s)))
		buf.WriteByte('#')
		buf.WriteString(s)
	}
	return buf.String()
}

func (c *Codec) Decode(s string) []string {
	if s == "" {
		return nil
	}
	
	var res []string
	i := 0
	
	for i < len(s) {
		j := i
		for s[j] != '#' {
			j++
		}
		length, _ := strconv.Atoi(s[i:j])
		i = j + 1
		res = append(res, s[i:i+length])
		i += length
	}
	
	return res
}
