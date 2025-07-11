package building_h2o_1117

import (
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Helper function to calculate the minimum of two integers.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Custom assertion to validate H2O output (allows flexible order within molecules).
func isValidH2OOutput(t *testing.T, actual, expected string) {
	numH := strings.Count(actual, "H")
	numO := strings.Count(actual, "O")
	expectedH := strings.Count(expected, "H")
	expectedO := strings.Count(expected, "O")
	assert.Equal(t, expectedH, numH, "Number of H atoms should match")
	assert.Equal(t, expectedO, numO, "Number of O atoms should match")
	assert.Equal(t, len(actual)%3, 0, "Output length should be multiple of 3 (HHO)")
}

func Test_H2O_Building(t *testing.T) {
	tests := []struct {
		name           string
		inputAtoms     string
		expectedOutput string
		expectTimeout  bool
	}{
		{
			name:           "Empty input",
			inputAtoms:     "",
			expectedOutput: "",
		},
		{
			name:           "One molecule (HOH)",
			inputAtoms:     "HOH",
			expectedOutput: "HHO",
		},
		{
			name:           "Two molecules (HHOHHO)",
			inputAtoms:     "HHOHHO",
			expectedOutput: "HHOHHO",
		},
		{
			name:           "Three molecules (HOHOHOH)",
			inputAtoms:     "HOHOHOH",
			expectedOutput: "HHOHHO",
		},
		{
			name:           "Mixed order (OHH)",
			inputAtoms:     "OHH",
			expectedOutput: "HHO",
		},
		{
			name:           "Many molecules, H first (HHHHHOO)",
			inputAtoms:     "HHHHHOO",
			expectedOutput: "HHOHHO",
		},
		{
			name:           "Many molecules, O first (OOHHHH)",
			inputAtoms:     "OOHHHH",
			expectedOutput: "HHOHHO",
		},
		{
			name:           "Large number of molecules (10 H2O)",
			inputAtoms:     strings.Repeat("H", 20) + strings.Repeat("O", 10),
			expectedOutput: strings.Repeat("HHO", 10),
		},
		{
			name:           "Only H atoms (should not form molecule, blocks)",
			inputAtoms:     "HHHH",
			expectedOutput: "",
			expectTimeout:  true,
		},
		{
			name:           "Only O atoms (should not form molecule, blocks)",
			inputAtoms:     "OO",
			expectedOutput: "",
			expectTimeout:  true,
		},
		{
			name:           "More H than needed (e.g., 3H, 1O, 1H blocks)",
			inputAtoms:     "HHHOH",
			expectedOutput: "HHO",
			expectTimeout:  true,
		},
		{
			name:           "More O than needed (e.g., 2H, 2O, 1O blocks)",
			inputAtoms:     "HHOO",
			expectedOutput: "HHO",
			expectTimeout:  true,
		},
		{
			name:           "Simultaneous arrivals of all atoms for one molecule (HHO)",
			inputAtoms:     "HHO",
			expectedOutput: "HHO",
		},
		{
			name:           "Alternating atoms (HOHOHOHO - 4H, 4O -> 2 H2O)",
			inputAtoms:     "HOHOHOHO",
			expectedOutput: "HHOHHO",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			h2o := NewH2O()

			var outputBuilder strings.Builder
			var outputMutex sync.Mutex
			var wg sync.WaitGroup

			releaseHydrogen := func() {
				outputMutex.Lock()
				outputBuilder.WriteString("H")
				outputMutex.Unlock()
			}

			releaseOxygen := func() {
				outputMutex.Lock()
				outputBuilder.WriteString("O")
				outputMutex.Unlock()
			}

			for _, char := range tt.inputAtoms {
				wg.Add(1)
				if char == 'H' {
					go func() {
						defer wg.Done()
						h2o.Hydrogen(releaseHydrogen)
					}()
				} else if char == 'O' {
					go func() {
						defer wg.Done()
						h2o.Oxygen(releaseOxygen)
					}()
				}
			}

			done := make(chan struct{})
			go func() {
				wg.Wait()
				close(done)
			}()

			select {
			case <-done:
				if tt.expectTimeout {
					t.Errorf("Test '%s' was expected to timeout but completed", tt.name)
				}
			case <-time.After(500 * time.Millisecond):
				if !tt.expectTimeout {
					t.Errorf("Test '%s' timed out (500ms). Partial output: %q", tt.name, outputBuilder.String())
				} else {
					t.Logf("Test '%s' correctly timed out as expected", tt.name)
				}
			}

			numHInput := strings.Count(tt.inputAtoms, "H")
			numOInput := strings.Count(tt.inputAtoms, "O")
			maxMolecules := min(numHInput/2, numOInput/1)
			expectedFinalOutput := strings.Repeat("HHO", maxMolecules)
			finalOutput := outputBuilder.String()
			isValidH2OOutput(t, finalOutput, expectedFinalOutput)
		})
	}
}
