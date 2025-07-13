package count_and_say_38

func countAndSay(n int) string {
	if n < 1 || n > 30 {
		return ""
	}
	
	result := "1"
	
	for range n - 1 {
		result = say(result)
	}
	
	return result
}

func say(s string) string {
	if len(s) == 0 {
		return ""
	}
	
	var result []byte
	
	for i := 0; i < len(s); {
		char := s[i]
		count := 1
		
		for i+count < len(s) && s[i+count] == char {
			count++
		}
		
		result = append(result, byte(count+'0'), char)
		i += count
	}
	
	return string(result)
}
