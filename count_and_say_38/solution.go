package count_and_say_38

import "strings"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	s := countAndSay(n - 1)
	var builder strings.Builder
	ch := s[0]
	freq := 0
	for _, c := range s {
		if byte(c) == ch {
			freq++
		} else {
			builder.WriteByte(byte(freq) + '0')
			builder.WriteByte(ch)
			freq = 1
			ch = byte(c)
		}
	}

	builder.WriteByte(byte(freq) + '0')
	builder.WriteByte(ch)
	return builder.String()
}
