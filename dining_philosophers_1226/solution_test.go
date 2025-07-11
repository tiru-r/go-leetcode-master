package dining_philosophers_1226

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Global test state
var (
	eventLog       []string     // Logs actions for debugging
	eventLogMutex  sync.Mutex   // Protects eventLog
	eatCounts      [5]int       // Tracks eating counts per philosopher
	eatCountsMutex sync.Mutex   // Protects eatCounts
	eating         [5]bool      // Tracks which philosophers are eating
	eatingMutex    sync.Mutex   // Protects eating array
	concurrentEats atomic.Int32 // Tracks total concurrent eating events
	rng            = rand.New(rand.NewSource(time.Now().UnixNano()))
	rngMutex       sync.Mutex   // Protects rng access
)

// resetTestState clears global state before each test
func resetTestState() {
	eventLogMutex.Lock()
	eventLog = []string{}
	eventLogMutex.Unlock()

	eatCountsMutex.Lock()
	for i := range eatCounts {
		eatCounts[i] = 0
	}
	eatCountsMutex.Unlock()

	eatingMutex.Lock()
	for i := range eating {
		eating[i] = false
	}
	eatingMutex.Unlock()

	concurrentEats.Store(0)
}

// Mock functions for philosopher actions
func mockPickLeft(philosopherID int) func() {
	return func() {
		eventLogMutex.Lock()
		eventLog = append(eventLog, fmt.Sprintf("P%d picks left", philosopherID))
		eventLogMutex.Unlock()
	}
}

func mockPickRight(philosopherID int) func() {
	return func() {
		eventLogMutex.Lock()
		eventLog = append(eventLog, fmt.Sprintf("P%d picks right", philosopherID))
		eventLogMutex.Unlock()
	}
}

func mockEat(philosopherID int) func() {
	return func() {
		eatingMutex.Lock()
		eating[philosopherID] = true
		leftNeighbor := (philosopherID + 4) % 5
		rightNeighbor := (philosopherID + 1) % 5
		if eating[leftNeighbor] || eating[rightNeighbor] {
			// Log error instead of using undefined t
			eventLogMutex.Lock()
			eventLog = append(eventLog, fmt.Sprintf("ERROR: P%d eating concurrently with neighbor %d or %d", philosopherID, leftNeighbor, rightNeighbor))
			eventLogMutex.Unlock()
		}
		concurrentEats.Add(1)
		eatCountsMutex.Lock()
		eatCounts[philosopherID]++
		eventLogMutex.Lock()
		eventLog = append(eventLog, fmt.Sprintf("P%d EATS (count=%d, concurrent=%d)", philosopherID, eatCounts[philosopherID], concurrentEats.Load()))
		eatCountsMutex.Unlock()
		eatingMutex.Unlock()

		// Randomized delay to simulate eating and expose concurrency issues
		rngMutex.Lock()
		delay := time.Duration(5+rng.Intn(10)) * time.Millisecond
		rngMutex.Unlock()
		time.Sleep(delay)

		eatingMutex.Lock()
		eating[philosopherID] = false
		eatingMutex.Unlock()
	}
}

func mockPutLeft(philosopherID int) func() {
	return func() {
		eventLogMutex.Lock()
		eventLog = append(eventLog, fmt.Sprintf("P%d puts left", philosopherID))
		eventLogMutex.Unlock()
	}
}

func mockPutRight(philosopherID int) func() {
	return func() {
		eventLogMutex.Lock()
		eventLog = append(eventLog, fmt.Sprintf("P%d puts right", philosopherID))
		eventLogMutex.Unlock()
	}
}

// Test_DiningPhilosophers tests the DiningPhilosophers implementation
func Test_DiningPhilosophers(t *testing.T) {
	// Initialize test with modern Go features

	tests := []struct {
		name             string
		numPhilosophers  int
		eatAttempts      int
		expectedDeadlock bool
		minEatsPerPhilo  int // Minimum eats to consider test successful (for stress tests)
		timeout          time.Duration
	}{
		{
			name:             "Standard case: 5 philosophers, 10 eats each",
			numPhilosophers:  5,
			eatAttempts:      10,
			expectedDeadlock: false,
			minEatsPerPhilo:  10,
			timeout:          30 * time.Second,
		},
		{
			name:             "Stress test: 5 philosophers, 50 eats each",
			numPhilosophers:  5,
			eatAttempts:      50,
			expectedDeadlock: false,
			minEatsPerPhilo:  50,
			timeout:          60 * time.Second,
		},
		{
			name:             "Minimal case: 5 philosophers, 1 eat each",
			numPhilosophers:  5,
			eatAttempts:      1,
			expectedDeadlock: false,
			minEatsPerPhilo:  1,
			timeout:          10 * time.Second,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetTestState()
			dp := Constructor()
			var wg sync.WaitGroup

			// Launch goroutines for each philosopher
			for i := 0; i < tt.numPhilosophers; i++ {
				philosopherID := i
				wg.Add(1)
				go func() {
					defer wg.Done()
					for j := 0; j < tt.eatAttempts; j++ {
						dp.WantsToEat(
							philosopherID,
							mockPickLeft(philosopherID),
							mockPickRight(philosopherID),
							mockEat(philosopherID),
							mockPutLeft(philosopherID),
							mockPutRight(philosopherID),
						)
						// Random delay between attempts to increase contention
						rngMutex.Lock()
						delay := time.Duration(rng.Intn(5)) * time.Millisecond
						rngMutex.Unlock()
						time.Sleep(delay)
					}
				}()
			}

			// Deadlock detection
			done := make(chan struct{})
			go func() {
				wg.Wait()
				close(done)
			}()

			select {
			case <-done:
				if tt.expectedDeadlock {
					t.Errorf("Test '%s' expected to deadlock but completed", tt.name)
				}
				t.Logf("Test '%s' completed successfully (no deadlock)", tt.name)
			case <-time.After(tt.timeout):
				if !tt.expectedDeadlock {
					var underfed []string
					eatCountsMutex.Lock()
					for i, count := range eatCounts {
						if count < tt.minEatsPerPhilo {
							underfed = append(underfed, fmt.Sprintf("P%d: %d/%d", i, count, tt.minEatsPerPhilo))
						}
					}
					eatCountsMutex.Unlock()
					t.Errorf("Test '%s' timed out after %v! Likely deadlock or starvation. Underfed philosophers: %s, Eat counts: %v",
						tt.name, tt.timeout, strings.Join(underfed, ", "), eatCounts)
				} else {
					t.Logf("Test '%s' correctly timed out as expected", tt.name)
				}
				t.FailNow()
			}

			// Verify eating progress
			if !tt.expectedDeadlock {
				expectedCounts := [5]int{}
				for i := 0; i < tt.numPhilosophers; i++ {
					expectedCounts[i] = tt.minEatsPerPhilo
				}
				eatCountsMutex.Lock()
				assert.Equal(t, expectedCounts, eatCounts, "Each philosopher should have eaten the expected number of times")
				eatCountsMutex.Unlock()

				// Verify fairness (no philosopher starved)
				for i, count := range eatCounts {
					if count < tt.minEatsPerPhilo {
						t.Errorf("Philosopher %d only ate %d times, expected at least %d", i, count, tt.minEatsPerPhilo)
					}
				}
			}

			// Log event summary for debugging
			eventLogMutex.Lock()
			if len(eventLog) > 0 {
				t.Logf("\n--- Event Log Summary for %s (showing first/last 10 entries) ---", tt.name)
				for i, event := range eventLog {
					if i < 10 || i >= len(eventLog)-10 {
						t.Logf("  %s", event)
					} else if i == 10 {
						t.Logf("  ... (%d events skipped) ...", len(eventLog)-20)
					}
				}
			}
			eventLogMutex.Unlock()
		})
	}
}
