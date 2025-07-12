package letter_combinations_of_a_phone_number_17

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	// constant lookup table indexed by digit-'0'
	keys := [...]string{
		2: "abc",
		3: "def",
		4: "ghi",
		5: "jkl",
		6: "mno",
		7: "pqrs",
		8: "tuv",
		9: "wxyz",
	}

	// validation + capacity
	cap := 1
	for i := 0; i < len(digits); i++ {
		d := digits[i] - '0'
		if d < 2 || d > 9 {
			return nil
		}
		cap *= len(keys[d])
	}

	// pre-allocate result and buffer
	res := make([]string, 0, cap)
	buf := make([]byte, len(digits))

	var dfs func(int)
	dfs = func(pos int) {
		if pos == len(digits) {
			res = append(res, string(buf))
			return
		}
		letters := keys[digits[pos]-'0']
		for i := 0; i < len(letters); i++ {
			buf[pos] = letters[i]
			dfs(pos + 1)
		}
	}

	dfs(0)
	return res
}
