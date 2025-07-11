package single_number_136

// NOTE: bitwise XOR arithmetic operator used to find single digit
// n XOR 0 = n
// n XOR n = 0
//
// So, for [2,2,1]:
// a = 0 ^ 2 = 2
// a = 2 ^ 2 = 0
// a = 0 ^ 1 = 1 = answer

// [2,1,2]
// a = 0 ^ 2 = 2
// a = 2 ^ 1 = 3 (because 00000010 XOR 000000001 -> 00000011 -> 3)
// a = 3 ^ 2 = 1 (because 00000011 XOR 00000010 -> 00000001 -> 1)
func singleNumber(nums []int) int {
	a := 0
	for _, num := range nums {
		a ^= num
	}
	return a
}
