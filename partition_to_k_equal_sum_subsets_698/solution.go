package partition_to_k_equal_sum_subsets_698

import "slices"

// Note: study again. Good problem!
func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}

	// if we can't evenly distribute the sum of each num over k partitions
	if sum%k != 0 {
		return false
	}

	// the target sum for each partition
	target := sum / k

	// if the largest value in the array is greater than the target sum, return false
	// ex: [5,2,1] k=2 target=4 (we can't use 5 for any partition)
	slices.Sort(nums)
	end := len(nums) - 1
	if nums[end] > target {
		return false
	}

	// clip off values that would fit directly into a partition and reduce k
	for end >= 0 && nums[end] == target {
		end--
		k--
	}

	return fill(make([]int, k), end, nums, target)
}

func fill(subsets []int, idx int, nums []int, target int) bool {
	if idx < 0 {
		return true
	}

	// choose a value from nums to try to slot into each subset
	pick := nums[idx]
	idx--

	// try to fill each subset
	for i := 0; i < len(subsets); i++ {
		if subsets[i]+pick <= target {
			subsets[i] += pick

			// explore if the subsets current value + chosen value is <= target
			if fill(subsets, idx, nums, target) {
				// if call to fill is true, we were able to partition successfully
				// if not, continue on the search
				return true
			}

			// un-choose the value for the current subset
			subsets[i] -= pick
		}

		// if current subset became 0 from the un-choose, then break
		// this keeps all unfilled subsets on the end and reduces work
		if subsets[i] == 0 {
			break
		}
	}

	return false
}
