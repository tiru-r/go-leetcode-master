package moving_average_from_data_stream_346

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovingAverage(t *testing.T) {
	ma := Constructor(3)

	assert.Equal(t, 1.0, ma.Next(1))
	assert.Equal(t, 1.5, ma.Next(2))
	assert.Equal(t, 2.0, ma.Next(3))
	assert.Equal(t, 3.0, ma.Next(4))
	assert.Equal(t, 4.0, ma.Next(5))
}

func BenchmarkNext(b *testing.B) {
	ma := Constructor(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ma.Next(i)
	}
}
