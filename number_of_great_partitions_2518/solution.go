package number_of_great_partitions_2518

func WaysToPartition(nums []int, k int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	total := sum(nums)
	result := countBasePartitions(nums, total)

	left, right := make(map[int]int), make(map[int]int)
	initRightDiffs(nums, total, right)

	prefixSum := 0
	for i := 0; i < n; i++ {
		delta := k - nums[i]
		result += left[-2*delta] + right[2*delta]

		if i < n-1 {
			prefixSum += nums[i]
			diff := prefixSum*2 - total
			moveDiff(left, right, diff)
		}
	}

	return result
}

func sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func countBasePartitions(nums []int, total int) int {
	count, prefixSum := 0, 0
	for i := 0; i < len(nums)-1; i++ {
		prefixSum += nums[i]
		if prefixSum*2 == total {
			count++
		}
	}
	return count
}

func initRightDiffs(nums []int, total int, right map[int]int) {
	prefixSum := 0
	for i := 0; i < len(nums)-1; i++ {
		prefixSum += nums[i]
		right[prefixSum*2-total]++
	}
}

func moveDiff(left, right map[int]int, diff int) {
	right[diff]--
	if right[diff] == 0 {
		delete(right, diff)
	}
	left[diff]++
}