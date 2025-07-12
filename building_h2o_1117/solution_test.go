package building_h2o_1117

import (
	"strings"
	"sync"
	"testing"
	"time"
)

// TestH2O tests the H2O molecule formation under various scenarios.
func TestH2O(t *testing.T) {
	// Helper function to collect output and validate it
	collectOutput := func(hCount, oCount int, t *testing.T, expectTimeout bool) string {
		var mu sync.Mutex
		var output strings.Builder
		h2o := New()

		// Define release functions that append to output
		releaseHydrogen := func() {
			mu.Lock()
			output.WriteString("H")
			mu.Unlock()
		}
		releaseOxygen := func() {
			mu.Lock()
			output.WriteString("O")
			mu.Unlock()
		}

		var wg sync.WaitGroup
		// Start goroutines in a mixed order to simulate varied arrivals
		starts := make([]string, hCount+oCount)
		for i := 0; i < hCount; i++ {
			starts[i] = "H"
		}
		for i := hCount; i < hCount+oCount; i++ {
			starts[i] = "O"
		}
		// Shuffle starts (simple Fisher-Yates shuffle)
		for i := len(starts) - 1; i > 0; i-- {
			j := int(time.Now().UnixNano()) % (i + 1)
			starts[i], starts[j] = starts[j], starts[i]
		}

		for _, atom := range starts {
			wg.Add(1)
			if atom == "H" {
				go func() {
					defer wg.Done()
					time.Sleep(time.Duration(randInt(0, 50)) * time.Millisecond)
					h2o.Hydrogen(releaseHydrogen)
				}()
			} else {
				go func() {
					defer wg.Done()
					time.Sleep(time.Duration(randInt(0, 50)) * time.Millisecond)
					h2o.Oxygen(releaseOxygen)
				}()
			}
		}

		// Wait for all goroutines to complete or timeout
		done := make(chan struct{})
		go func() {
			wg.Wait()
			close(done)
		}()

		select {
		case <-done:
			if expectTimeout {
				t.Errorf("Expected timeout, but test completed")
			}
			mu.Lock()
			result := output.String()
			mu.Unlock()
			return result
		case <-time.After(1 * time.Second):
			if !expectTimeout {
				t.Fatal("Test timed out, expected completion")
			}
			mu.Lock()
			result := output.String()
			mu.Unlock()
			return result
		}
	}

	// Helper function to validate output
	validateOutput := func(t *testing.T, output string, expectedMolecules int) {
		t.Helper()
		// Expected length: each molecule is "HHO" (3 characters)
		if len(output) != expectedMolecules*3 {
			t.Errorf("Expected output length %d, got %d", expectedMolecules*3, len(output))
		}

		// Count H and O
		hCount := strings.Count(output, "H")
		oCount := strings.Count(output, "O")
		if hCount != expectedMolecules*2 || oCount != expectedMolecules {
			t.Errorf("Expected %d H and %d O, got %d H and %d O", expectedMolecules*2, expectedMolecules, hCount, oCount)
		}

		// Validate strict HHO order for each molecule
		for i := 0; i < len(output)-2; i += 3 {
			if i+3 <= len(output) {
				chunk := output[i : i+3]
				if chunk != "HHO" {
					t.Errorf("Invalid molecule order at position %d: %s, expected HHO", i, chunk)
				}
			}
		}
	}

	// Test cases
	tests := []struct {
		name              string
		hCount, oCount    int
		expectedMolecules int
		expectTimeout     bool
	}{
		{
			name:              "SingleMolecule",
			hCount:            2,
			oCount:            1,
			expectedMolecules: 1,
			expectTimeout:     false,
		},
		{
			name:              "MultipleMolecules",
			hCount:            6,
			oCount:            3,
			expectedMolecules: 3,
			expectTimeout:     false,
		},
		{
			name:              "ExcessHydrogen",
			hCount:            4,
			oCount:            1,
			expectedMolecules: 1,
			expectTimeout:     true, // Excess H atoms remain blocked
		},
		{
			name:              "ExcessOxygen",
			hCount:            2,
			oCount:            2,
			expectedMolecules: 1,
			expectTimeout:     true, // Excess O atom remains blocked
		},
		{
			name:              "NoMolecules",
			hCount:            1,
			oCount:            1,
			expectedMolecules: 0,
			expectTimeout:     true, // Insufficient atoms, all remain blocked
		},
		{
			name:              "StressTest",
			hCount:            20,
			oCount:            10,
			expectedMolecules: 10,
			expectTimeout:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := collectOutput(tt.hCount, tt.oCount, t, tt.expectTimeout)
			if !tt.expectTimeout {
				validateOutput(t, output, tt.expectedMolecules)
			} else if len(output) != tt.expectedMolecules*3 {
				t.Errorf("Expected output length %d for timeout case, got %d", tt.expectedMolecules*3, len(output))
			}
		})
	}
}

// randInt generates a random integer between min and max (inclusive).
func randInt(min, max int) int {
	return min + int(time.Now().UnixNano())%(max-min+1)
}
