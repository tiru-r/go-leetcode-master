package sum_of_two_integers_371

func getSum(a, b int) int {
	for b != 0 {
		a, b = a^b, (a&b)<<1
	}
	return a
}
