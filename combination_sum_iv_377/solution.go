package combination_sum_iv_377

import (
	"math"
	"sync"
)

// dpArrayPool recycles DP slices to cut GC pressure.
var dpArrayPool = sync.Pool{
	New: func() any { return make([]int, 1000) },
}

// CombinationSum4InPlace computes the answer in the provided buffer.
// The caller must guarantee len(buf) ≥ target+1.
// On return, buf is left intact for reuse.
func CombinationSumInPlace(nums []int, target int, buf []int) int {
	if target == 0 {
		return 1
	}
	if len(nums) == 0 || target < 0 {
		return 0
	}

	// zero the slice once (faster than clear on Go ≤1.20)
	for i := range buf[:target+1] {
		buf[i] = 0
	}
	buf[0] = 1

	// in-place filter (zero-copy)
	coins := nums
	nCoins := 0
	for _, v := range coins {
		if v > 0 && v <= target {
			coins[nCoins] = v
			nCoins++
		}
	}
	if nCoins == 0 {
		return 0
	}
	coins = coins[:nCoins]

	for i := 1; i <= target; i++ {
		for _, c := range coins {
			if c <= i {
				buf[i] += buf[i-c]
			}
		}
		// optional overflow clamp
		if buf[i] < 0 {
			buf[i] = math.MaxInt32
		}
	}
	return buf[target]
}

// Wrapper for the old signature; allocates only once per goroutine via sync.Pool.
func CombinationSum(nums []int, target int) int {
	if target < 0 {
		return 0
	}
	buf := dpArrayPool.Get().([]int)
	defer dpArrayPool.Put(buf)
	if target+1 > len(buf) {
		buf = make([]int, target+1)
	} else {
		buf = buf[:target+1]
	}
	return CombinationSumInPlace(nums, target, buf)
}
