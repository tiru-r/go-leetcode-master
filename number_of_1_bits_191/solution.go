package number_of_1_bits_191

import "math/bits"

func hammingWeight(num uint32) int {
	return bits.OnesCount32(num)
}

func hammingWeightManual(num uint32) int {
	count := 0
	for num != 0 {
		num &= num - 1
		count++
	}
	return count
}
