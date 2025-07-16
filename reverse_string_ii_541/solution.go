package reverse_string_ii_541

import "slices"

func reverseStr(s string, k int) string {
	r := []rune(s)
	for i := 0; i < len(r); i += 2 * k {
		end := min(i+k, len(r))
		slices.Reverse(r[i:end])
	}
	return string(r)
}
