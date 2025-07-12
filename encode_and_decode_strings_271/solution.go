package encode_and_decode_strings_271

type Codec struct{}

// Encode converts []string → single byte string using length-prefixed format.
func (c *Codec) Encode(strs []string) string {
	// pre-compute exact byte count
	n := 0
	for _, s := range strs {
		n += len(s) + 2 // 2 = len digit(s) + '_'
		if l := len(s); l >= 10 {
			for l > 0 {
				n++
				l /= 10
			}
		} else {
			n++
		}
	}
	buf := make([]byte, 0, n)
	for _, s := range strs {
		l := len(s)
		// write length as ASCII digits
		if l == 0 {
			buf = append(buf, '0', '_')
		} else {
			start := len(buf)
			for l > 0 {
				buf = append(buf, byte(l%10)+'0')
				l /= 10
			}
			// reverse digits in place
			for l, r := start, len(buf)-1; l < r; l, r = l+1, r-1 {
				buf[l], buf[r] = buf[r], buf[l]
			}
			buf = append(buf, '_')
		}
		buf = append(buf, s...)
	}
	return string(buf)
}

// Decode converts the encoded string back to []string.
func (c *Codec) Decode(s string) []string {
	if len(s) == 0 {
		return nil
	}
	b := []byte(s)
	res := make([]string, 0, 8)
	i := 0
	for i < len(b) {
		// read length
		l := 0
		for i < len(b) && b[i] != '_' {
			l = l*10 + int(b[i]-'0')
			i++
		}
		i++ // skip '_'
		// slice out the string
		res = append(res, string(b[i:i+l]))
		i += l
	}
	return res
}
