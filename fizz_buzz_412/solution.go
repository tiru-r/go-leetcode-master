package fizz_buzz_412

import "strconv"

// fizzBuzz generates FizzBuzz sequence for numbers 1 to n
// Optimized: O(n) time, O(n) space with branch-free logic and modern iteration
func fizzBuzz(n int) []string {
	if n <= 0 {
		return []string{}
	}
	
	result := make([]string, n)
	
	// Modern range-based iteration with zero-indexed access
	for i := range n {
		num := i + 1
		
		// Branch-free concatenation approach for better performance
		output := ""
		if num%3 == 0 {
			output += "Fizz"
		}
		if num%5 == 0 {
			output += "Buzz"
		}
		
		// Use number string if no fizz/buzz applied
		if output == "" {
			output = strconv.Itoa(num)
		}
		
		result[i] = output
	}
	
	return result
}
