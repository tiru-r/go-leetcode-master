package decrypt_string_from_alphabet_to_integer_mapping_1309

func freqAlphabets(s string) string {
	buf := make([]byte, 0, len(s)>>1+1)
	
	for i := len(s) - 1; i >= 0; {
		if s[i] == '#' {
			num := (s[i-2]-'0')*10 + (s[i-1]-'0')
			buf = append(buf, 'a'+byte(num-1))
			i -= 3
		} else {
			buf = append(buf, 'a'+byte(s[i]-'1'))
			i--
		}
	}
	
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	
	return string(buf)
}
