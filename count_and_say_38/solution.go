package count_and_say_38

func countAndSay(n int) string {
	if n < 1 || n > 30 {
		return ""
	}

	// Two alternating buffers (max length < 1e5 for n ≤ 30).
	var cur, next []byte
	cur = []byte("1")

	for k := 2; k <= n; k++ {
		// Estimate max possible length: len(cur)*2 is always enough.
		next = make([]byte, 0, len(cur)*2)

		i := 0
		for i < len(cur) {
			c := cur[i]
			cnt := 1
			for i+1 < len(cur) && cur[i+1] == c {
				i++
				cnt++
			}
			next = append(next, byte(cnt+'0'), c)
			i++
		}
		cur = next
	}
	return string(cur)
}
