package longest_repeating_character_replacement_424

func characterReplacement(s string, k int) int {
	// Use byte keys instead of string keys for better performance
	counts := make(map[byte]int)
	var start, maxLen, maxCnt int
	for end := 0; end < len(s); end++ {
		// Direct byte operations avoid string allocation
		counts[s[end]]++

		// max is the largest count of a single, unique character in the window
		maxCnt = max(maxCnt, counts[s[end]])

		// if the amount left to modify is greater than the k
		// modifications we can use, then shrink the window.
		windowSize := end - start + 1
		leftToModify := windowSize - maxCnt
		if leftToModify > k {
			counts[s[start]]--
			start++
		}

		maxLen = max(maxLen, end-start+1)
	}

	return maxLen
}

// Really close, but too slow.
func characterReplacement1(s string, k int) int {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}

	longest := 0
	pick := 0
	left := pick - 1
	right := pick + 1
	for pick < len(s) {
		// go left and use up all k
		ll, kl := expandLeft(s, k, left, pick)
		longest = max(longest, ll)

		// use extra kl to expand right
		if kl > 0 {
			lr, _ := expandRight(s, kl, right, pick)
			longest = max(longest, ll+lr-1)
		}

		// go right from pick using k
		lr, kr := expandRight(s, k, right, pick)
		longest = max(longest, lr)

		// use extra kr to expand left
		if kr > 0 {
			ll, _ := expandLeft(s, kr, left, pick)
			longest = max(longest, lr+ll-1)
		}

		pick++
		left = pick - 1
		right = pick + 1
	}

	return longest
}

func expandLeft(s string, kl int, left int, pick int) (int, int) {
	longest := 0
	for left > -1 && (kl > 0 || s[left] == s[pick]) {
		longest = max(longest, pick-left+1)

		if s[left] != s[pick] {
			kl--
		}
		left--
	}
	return longest, kl
}

func expandRight(s string, kr int, right int, pick int) (int, int) {
	longest := 0
	for right < len(s) && (kr > 0 || s[right] == s[pick]) {
		longest = max(longest, right-pick+1)

		if s[right] != s[pick] {
			kr--
		}
		right++
	}
	return longest, kr
}

// Note: Close but incorrect
func characterReplacement0(s string, k int) int {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}

	longest := 0
	start := 0
	end := start + 1
	for start < len(s) && end < len(s) {
		// repeating?
		if s[start] == s[end] {
			longest = max(longest, end-start+1)
			end++
		} else {
			// have an operation left to use?
			if k > 0 {
				longest = max(longest, end-start+1)
				k--
				end++
			} else {
				start++
				end = start + 1
			}
		}
	}

	return longest
}
