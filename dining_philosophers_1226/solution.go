package dining_philosophers_1226

import "sync"

type DiningPhilosophers struct {
	forks [5]sync.Mutex
	seats chan struct{}
}

func NewDiningPhilosophers() *DiningPhilosophers {
	return &DiningPhilosophers{
		seats: make(chan struct{}, 4),
	}
}

func (dp *DiningPhilosophers) WantsToEat(
	philosopher int,
	pickLeft, pickRight, eat, putLeft, putRight func(),
) {
	dp.seats <- struct{}{}
	defer func() { <-dp.seats }()

	leftFork := philosopher
	rightFork := (philosopher + 1) % 5

	if leftFork > rightFork {
		leftFork, rightFork = rightFork, leftFork
	}

	dp.forks[leftFork].Lock()
	dp.forks[rightFork].Lock()

	pickLeft()
	pickRight()
	eat()
	putLeft()
	putRight()

	dp.forks[rightFork].Unlock()
	dp.forks[leftFork].Unlock()
}
