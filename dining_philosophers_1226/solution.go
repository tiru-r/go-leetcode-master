package dining_philosophers_1226

import "sync"

type DiningPhilosophers struct {
	forks [5]sync.Mutex
	seats chan struct{}
}

func NewDiningPhilosophers() *DiningPhilosophers {
	return &DiningPhilosophers{
		seats: make(chan struct{}, 4), // Allow at most 4 philosophers at table
	}
}

func (dp *DiningPhilosophers) WantsToEat(
	philosopher int,
	pickLeft, pickRight, eat, putLeft, putRight func(),
) {
	// Acquire seat at table (deadlock prevention: limit concurrent access)
	dp.seats <- struct{}{}
	defer func() { <-dp.seats }()

	leftFork := philosopher
	rightFork := (philosopher + 1) % 5

	// Deadlock prevention: always acquire lower-numbered fork first
	if leftFork > rightFork {
		leftFork, rightFork = rightFork, leftFork
	}

	// Acquire both forks atomically
	dp.forks[leftFork].Lock()
	dp.forks[rightFork].Lock()

	// Perform eating sequence
	pickLeft()
	pickRight()
	eat()
	putLeft()
	putRight()

	// Release forks in reverse order
	dp.forks[rightFork].Unlock()
	dp.forks[leftFork].Unlock()
}
