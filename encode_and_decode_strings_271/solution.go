package encode_and_decode_strings_271

import (
	"strconv"
	"strings"
)

type Codec struct{}

func (c *Codec) Encode(strs []string) string {
	if strs == nil {
		return "nil"
	}
	if len(strs) == 0 {
		return "empty"
	}
	var buf strings.Builder
	for _, s := range strs {
		buf.WriteString(strconv.Itoa(len(s)))
		buf.WriteByte('#')
		buf.WriteString(s)
	}
	return buf.String()
}

func (c *Codec) Decode(s string) []string {
	if s == "nil" {
		return nil
	}
	if s == "empty" {
		return []string{}
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
