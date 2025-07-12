package dining_philosophers_1226

import "sync"

type DiningPhilosophers struct {
	forks [5]sync.Mutex
	slot  chan struct{}
}

func Constructor() DiningPhilosophers {
	return DiningPhilosophers{
		slot: make(chan struct{}, 4), // at most 4 eaters
	}
}

// WantsToEat handles the whole life-cycle.
func (d *DiningPhilosophers) WantsToEat(
	phil int,
	pickLeft, pickRight, eat, putLeft, putRight func(),
) {
	d.slot <- struct{}{}        // acquire global seat
	defer func() { <-d.slot }() // release seat

	left, right := phil, (phil+1)%5
	// always lock lower-number fork first to avoid deadlock
	if left > right {
		left, right = right, left
	}

	d.forks[left].Lock()
	d.forks[right].Lock()

	pickLeft()
	pickRight()
	eat()
	putRight()
	putLeft()

	d.forks[right].Unlock()
	d.forks[left].Unlock()
}
