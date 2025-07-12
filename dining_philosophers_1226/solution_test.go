package dining_philosophers_1226

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// fake actions that just record calls
type trace struct{ calls []string }

func (t *trace) pickLeft()  { t.calls = append(t.calls, "pickL") }
func (t *trace) pickRight() { t.calls = append(t.calls, "pickR") }
func (t *trace) eat()       { t.calls = append(t.calls, "eat") }
func (t *trace) putLeft()   { t.calls = append(t.calls, "putL") }
func (t *trace) putRight()  { t.calls = append(t.calls, "putR") }

func TestDiningPhilosophers(t *testing.T) {
	dp := Constructor()
	tr := &trace{}

	// single philosopher must complete the whole sequence
	dp.WantsToEat(0, tr.pickLeft, tr.pickRight, tr.eat, tr.putLeft, tr.putRight)
	expected := []string{"pickL", "pickR", "eat", "putR", "putL"}
	assert.Equal(t, expected, tr.calls)

	// concurrent test: 5 goroutines, each runs once; no deadlock within 1 s
	done := make(chan bool, 5)
	for i := 0; i < 5; i++ {
		go func(id int) {
			dp.WantsToEat(id,
				func() {}, func() {}, func() {}, func() {}, func() {},
			)
			done <- true
		}(i)
	}

	timeout := time.After(time.Second)
	for i := 0; i < 5; i++ {
		select {
		case <-done:
		case <-timeout:
			t.Fatal("deadlock or timeout")
		}
	}
}
