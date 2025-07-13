package dining_philosophers_1226

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestDiningPhilosophers(t *testing.T) {
	const (
		philosophers = 5
		eatRounds    = 100
	)

	dp := NewDiningPhilosophers()

	var (
		wg              sync.WaitGroup
		meals           [philosophers]int64
		neighbourEating [philosophers]int64
	)

	other := func(p, fork int) *int64 {
		if fork == p {
			return &neighbourEating[(p+4)%philosophers]
		}
		return &neighbourEating[(p+1)%philosophers]
	}

	for p := 0; p < philosophers; p++ {
		p := p
		wg.Add(1)
		go func() {
			defer wg.Done()
			for round := 0; round < eatRounds; round++ {
				dp.WantsToEat(p,
					func() {
						left := p
						if atomic.LoadInt64(other(p, left)) == 1 {
							t.Errorf("philosopher %d: left fork already taken", p)
						}
						atomic.StoreInt64(&neighbourEating[left], 1)
					},
					func() {
						right := (p + 1) % philosophers
						if atomic.LoadInt64(other(p, right)) == 1 {
							t.Errorf("philosopher %d: right fork already taken", p)
						}
						atomic.StoreInt64(&neighbourEating[right], 1)
					},
					func() {
						atomic.AddInt64(&meals[p], 1)
						time.Sleep(1 * time.Microsecond)
					},
					func() { atomic.StoreInt64(&neighbourEating[p], 0) },
					func() { atomic.StoreInt64(&neighbourEating[(p+1)%philosophers], 0) },
				)
			}
		}()
	}

	wg.Wait()

	for p := 0; p < philosophers; p++ {
		if got := atomic.LoadInt64(&meals[p]); got != eatRounds {
			t.Errorf("philosopher %d ate %d times, want %d", p, got, eatRounds)
		}
	}
}
