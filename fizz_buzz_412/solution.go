package fizz_buzz_412

import (
	"strconv"
	"strings"
)

func fizzBuzz(n int) []string {
	a := make([]string, n)
	for i := 1; i <= n; i++ {
		var builder strings.Builder
		
		if i%3 == 0 {
			builder.WriteString("Fizz")
		}

		if i%5 == 0 {
			builder.WriteString("Buzz")
		}

		if builder.Len() == 0 {
			a[i-1] = strconv.Itoa(i)
		} else {
			a[i-1] = builder.String()
		}
	}

	return a
}
