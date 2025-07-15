package partition_to_k_equal_sum_subsets_698

import "slices"

func canPartitionKSubsets(nums []int, k int) bool {
	total := sum(nums)
	if total%k != 0 {
		return false
	}

	target := total / k
	slices.Sort(nums)
	
	if nums[len(nums)-1] > target {
		return false
	}

	end := trimExact(nums, target, &k)
	return backtrack(make([]int, k), end, nums, target)
}

func sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func trimExact(nums []int, target int, k *int) int {
	end := len(nums) - 1
	for end >= 0 && nums[end] == target {
		end--
		*k--
	}
	return end
}

func backtrack(subsets []int, idx int, nums []int, target int) bool {
	if idx < 0 {
		return true
	}

	num := nums[idx]
	for i := range subsets {
		if subsets[i]+num <= target {
			subsets[i] += num
			if backtrack(subsets, idx-1, nums, target) {
				return true
			}
			subsets[i] -= num
		}
		
		if subsets[i] == 0 {
			break
		}
	}

	return false
}
