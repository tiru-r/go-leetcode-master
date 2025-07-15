package reverse_bits_190

import "math/bits"

func reverseBits(num uint32) uint32 {
	return bits.Reverse32(num)
}

func reverseBitsManual(num uint32) uint32 {
	result := uint32(0)
	for range 32 {
		result = (result << 1) | (num & 1)
		num >>= 1
	}
	return result
}
