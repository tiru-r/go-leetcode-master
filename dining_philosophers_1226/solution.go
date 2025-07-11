package dining_philosophers_1226

import (
	"sync"
	"time"
)

// DiningPhilosophers manages the dining philosophers problem with deadlock and starvation prevention.
type DiningPhilosophers struct {
	forks       [5]sync.Mutex // Mutex for each fork
	limit       chan struct{} // Limits concurrent eaters to 4 to prevent deadlock
	eatingCount [5]int32      // Tracks eating count per philosopher
	lastEaten   [5]time.Time  // Tracks when each philosopher last ate
	mutex       sync.RWMutex  // Protects shared state
}

// Constructor initializes the DiningPhilosophers struct.
func Constructor() DiningPhilosophers {
	now := time.Now()
	return DiningPhilosophers{
		forks:     [5]sync.Mutex{},
		limit:     make(chan struct{}, 4), // Allow at most 4 philosophers to eat concurrently
		lastEaten: [5]time.Time{now, now, now, now, now},
	}
}

// WantsToEat handles a philosopher's request to eat with guaranteed deadlock prevention.
func (dp *DiningPhilosophers) WantsToEat(philosopher int,
	pickLeftFork, pickRightFork, eat, putLeftFork, putRightFork func()) {

	// Step 1: Acquire permission to eat (max 4 philosophers)
	dp.limit <- struct{}{}
	defer func() { <-dp.limit }()

	// Step 2: Get fork indices
	leftFork := philosopher
	rightFork := (philosopher + 1) % 5

	// Step 3: Always lock in ascending order to prevent circular wait
	if leftFork < rightFork {
		dp.forks[leftFork].Lock()
		defer dp.forks[leftFork].Unlock()
		dp.forks[rightFork].Lock()
		defer dp.forks[rightFork].Unlock()
	} else {
		dp.forks[rightFork].Lock()
		defer dp.forks[rightFork].Unlock()
		dp.forks[leftFork].Lock()
		defer dp.forks[leftFork].Unlock()
	}

	// Step 4: Execute eating sequence
	pickLeftFork()
	pickRightFork()
	eat()
	putRightFork()
	putLeftFork()

	// Step 5: Update eating statistics
	dp.mutex.Lock()
	dp.eatingCount[philosopher]++
	dp.lastEaten[philosopher] = time.Now()
	dp.mutex.Unlock()
}
