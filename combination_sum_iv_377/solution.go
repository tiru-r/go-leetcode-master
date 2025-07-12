package combination_sum_iv_377

import (
	"math"
	"sort"
	"sync"
)

// dpArrayPool recycles DP slices to reduce GC pressure.
var dpArrayPool = sync.Pool{
	New: func() any { return make([]int, 1_001) }, // common case
}

// CombinationSumInPlace computes the number of combinations summing to target.
// buf must have length ≥ target+1; it is overwritten and left intact for reuse.
func CombinationSumInPlace(nums []int, target int, buf []int) int {
	if target == 0 {
		return 1
	}
	if target < 0 {
		return 0
	}

	// 1. Fast zero-out for small slices
	for i := 0; i <= target; i++ {
		buf[i] = 0
	}
	buf[0] = 1

	// 2. In-place filter & sort
	n := 0
	for _, v := range nums {
		if 0 < v && v <= target {
			nums[n] = v
			n++
		}
	}
	if n == 0 {
		return 0
	}
	coins := nums[:n]
	sort.Ints(coins)

	// 3. Bottom-up DP with early-break
	for i := 1; i <= target; i++ {
		for _, c := range coins {
			if c > i {
				break
			}
			// 4. Safe 64-bit arithmetic, clamp to math.MaxInt
			if buf[i-c] > math.MaxInt-buf[i] {
				buf[i] = math.MaxInt
			} else {
				buf[i] += buf[i-c]
			}
		}
	}
	return buf[target]
}

// CombinationSum is the public helper using the pool.
func CombinationSum(nums []int, target int) int {
	if target < 0 {
		return 0
	}
	for _, v := range nums {
		if v < 0 {
			return 0 // invalid input
		}
	}

	buf := dpArrayPool.Get().([]int)
	if cap(buf) < target+1 {
		buf = make([]int, target+1)
	} else {
		buf = buf[:target+1]
	}
	res := CombinationSumInPlace(nums, target, buf)
	dpArrayPool.Put(buf) // avoid defer in hot path
	return res
}
