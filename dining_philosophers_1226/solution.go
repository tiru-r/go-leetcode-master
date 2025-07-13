package dining_philosophers_1226

import "sync"

type DiningPhilosophers struct {
	forks [5]sync.Mutex
	sem   chan struct{}
}

func NewDiningPhilosophers() *DiningPhilosophers {
	return &DiningPhilosophers{
		sem: make(chan struct{}, 4),
	}
}

func (dp *DiningPhilosophers) WantsToEat(
	philosopher int,
	pickLeft, pickRight, eat, putLeft, putRight func(),
) {
	dp.sem <- struct{}{}
	
	left, right := philosopher, (philosopher+1)%5
	if left > right {
		left, right = right, left
	}
	
	dp.forks[left].Lock()
	dp.forks[right].Lock()
	
	pickLeft()
	pickRight()
	eat()
	putLeft()
	putRight()
	
	dp.forks[right].Unlock()
	dp.forks[left].Unlock()
	<-dp.sem
}
