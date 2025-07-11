package rotate_string_796

import "strings"

func rotateString(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}

	if A == B {
		return true
	}

	// Optimized: use strings.Builder to avoid O(n²) string concatenation
	var builder strings.Builder
	builder.Grow(len(A))

	a := A
	for i := 0; i < len(A); i++ {
		if a == B {
			return true
		}

		// rotate a by prepending end to front using efficient builder
		builder.Reset()
		builder.WriteString(a[len(A)-1:])
		builder.WriteString(a[0 : len(A)-1])
		a = builder.String()
	}

	return false
}
