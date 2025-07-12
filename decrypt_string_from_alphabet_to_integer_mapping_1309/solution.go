package decrypt_string_from_alphabet_to_integer_mapping_1309

// freqAlphabets decodes the 1–26 '#' string into a-z.
func freqAlphabets(s string) string {
	buf := make([]byte, 0, len(s)) // worst-case length
	for i := 0; i < len(s); {
		if i+2 < len(s) && s[i+2] == '#' { // "10#".."26#"
			buf = append(buf, 'j'+byte(s[i]-'1')*10+byte(s[i+1]-'0')-10)
			i += 3
		} else { // "1".."9"
			buf = append(buf, 'a'+byte(s[i]-'1'))
			i++
		}
	}
	return string(buf)
}
