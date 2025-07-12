package building_h2o_1117

import "sync"

type H2O struct {
	hReady   chan struct{} // Buffer for 2 H atoms
	oReady   chan struct{} // Buffer for 1 O atom
	hRelease chan struct{} // Unbuffered for H release
	oRelease chan struct{} // Unbuffered for O release
	hDone    chan struct{} // Track H completion for ordering
	hCount   int           // Count of available H atoms
	oCount   int           // Count of available O atoms
	mu       sync.Mutex
}

// New creates and initializes a new H2O instance.
func New() *H2O {
	return &H2O{
		hReady:   make(chan struct{}, 2),
		oReady:   make(chan struct{}, 1),
		hRelease: make(chan struct{}),
		oRelease: make(chan struct{}),
		hDone:    make(chan struct{}, 2),
	}
}

// Hydrogen signals readiness and waits for release.
func (h *H2O) Hydrogen(releaseHydrogen func()) {
	h.hReady <- struct{}{} // Signal H atom arrival

	h.mu.Lock()
	h.hCount++
	// Check if we can form a molecule
	if h.hCount >= 2 && h.oCount >= 1 {
		// Consume ready signals
		<-h.hReady
		<-h.hReady
		<-h.oReady
		h.hCount -= 2
		h.oCount -= 1
		h.mu.Unlock()

		// Signal release in strict H, H, O order (use goroutine to avoid blocking)
		go func() {
			h.hRelease <- struct{}{} // First H
			h.hRelease <- struct{}{} // Second H
			<-h.hDone                // Wait for first H to complete
			<-h.hDone                // Wait for second H to complete
			h.oRelease <- struct{}{} // Then release O
		}()
	} else {
		h.mu.Unlock()
	}

	<-h.hRelease // Wait for release signal
	releaseHydrogen()
	h.hDone <- struct{}{} // Signal completion
}

// Oxygen coordinates molecule formation and release.
func (h *H2O) Oxygen(releaseOxygen func()) {
	h.oReady <- struct{}{} // Signal O atom arrival

	h.mu.Lock()
	h.oCount++
	// Check if we can form a molecule
	if h.hCount >= 2 && h.oCount >= 1 {
		// Consume ready signals
		<-h.hReady
		<-h.hReady
		<-h.oReady
		h.hCount -= 2
		h.oCount -= 1
		h.mu.Unlock()

		// Signal release in strict H, H, O order (use goroutine to avoid blocking)
		go func() {
			h.hRelease <- struct{}{} // First H
			h.hRelease <- struct{}{} // Second H
			<-h.hDone                // Wait for first H to complete
			<-h.hDone                // Wait for second H to complete
			h.oRelease <- struct{}{} // Then release O
		}()
	} else {
		h.mu.Unlock()
	}

	<-h.oRelease // Wait for release signal
	releaseOxygen()
}
