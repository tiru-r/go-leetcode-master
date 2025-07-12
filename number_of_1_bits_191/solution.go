package number_of_1_bits_191

import "math/bits"

func hammingWeight(num uint32) int {
	return bits.OnesCount32(num)
}
