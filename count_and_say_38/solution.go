package count_and_say_38

import "strings"

func countAndSay(n int) string {
	result := "1"
	var builder strings.Builder
	
	for i := 1; i < n; i++ {
		builder.Reset()
		builder.Grow(len(result) * 2)
		
		ch := result[0]
		freq := 1
		
		for j := 1; j < len(result); j++ {
			if result[j] == ch {
				freq++
			} else {
				builder.WriteByte(byte(freq) + '0')
				builder.WriteByte(ch)
				freq = 1
				ch = result[j]
			}
		}
		
		builder.WriteByte(byte(freq) + '0')
		builder.WriteByte(ch)
		result = builder.String()
	}
	
	return result
}
