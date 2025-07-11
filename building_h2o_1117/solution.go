package building_h2o_1117

// H2O struct uses buffered channels as semaphores to control atom flow.
type H2O struct {
	// hSem acts as a semaphore for Hydrogen atoms.
	// Capacity 2 means 2 H atoms can acquire a "permit" to proceed.
	hSem chan struct{}

	// oSem acts as a semaphore for Oxygen atoms.
	// Capacity 1 means 1 O atom can acquire a "permit" to proceed.
	oSem chan struct{}

	// hDone signals that a Hydrogen atom has called its release function.
	// Capacity 2 because 2 H atoms will signal this.
	hDone chan struct{}

	// oDone signals that an Oxygen atom has called its release function.
	// Capacity 1 because 1 O atom will signal this.
	oDone chan struct{}
}

// NewH2O initializes the H2O struct with the necessary channels.
// The hSem and oSem channels are initially empty. Atoms will block
// until they can acquire a permit. The hDone and oDone channels are
// used for rendezvous after atoms have released.
func NewH2O() *H2O {
	return &H2O{
		hSem:  make(chan struct{}, 2), // H atoms acquire permits here
		oSem:  make(chan struct{}, 1), // O atom acquires permit here
		hDone: make(chan struct{}, 2), // H atoms signal completion here
		oDone: make(chan struct{}, 1), // O atom signals completion here
	}
}

// Hydrogen is called by a goroutine representing a Hydrogen atom.
// It ensures that releaseHydrogen is called only when a molecule can be formed.
func (h2o *H2O) Hydrogen(releaseHydrogen func()) {
	// Step 1: Acquire a permit to be a Hydrogen atom in the current molecule.
	// This will block if 2 Hydrogen atoms are already "in progress" for a molecule
	// and haven't finished releasing and reset their slots.
	h2o.hSem <- struct{}{}

	// Step 2: Call the release function. This happens after the atom has successfully
	// acquired its slot, meaning it's ready to form part of a molecule.
	releaseHydrogen()

	// Step 3: Signal that this Hydrogen atom has completed its release step.
	// This signal is crucial for the Oxygen atom to know when both H atoms are ready.
	h2o.hDone <- struct{}{}

	// Step 4: Wait for the Oxygen atom to signal that it has also released.
	// This ensures that the H atom doesn't "reset" its slot until the entire
	// molecule has been formed and released.
	<-h2o.oDone

	// Step 5: Release the permit for this Hydrogen atom.
	// This allows another Hydrogen atom to start forming the next molecule.
	// The order of releasing permits (hSem and oSem) is important to ensure
	// that a new molecule doesn't start forming until the previous one is complete.
	<-h2o.hSem // This makes the hSem available for the NEXT molecule.
}

// Oxygen is called by a goroutine representing an Oxygen atom.
// It ensures that releaseOxygen is called only when a molecule can be formed.
func (h2o *H2O) Oxygen(releaseOxygen func()) {
	// Step 1: Acquire a permit to be an Oxygen atom in the current molecule.
	// This will block if an Oxygen atom is already "in progress" for a molecule.
	h2o.oSem <- struct{}{}

	// Step 2: Wait for both Hydrogen atoms to signal that they have completed their release step.
	// This ensures that the Oxygen atom only calls releaseOxygen when both H atoms are ready.
	<-h2o.hDone // Wait for the first H to be done
	<-h2o.hDone // Wait for the second H to be done

	// Step 3: Call the release function. This happens after all 3 atoms are ready.
	releaseOxygen()

	// Step 4: Signal that this Oxygen atom has completed its release step.
	// This signal is crucial for the Hydrogen atoms to know when the molecule is fully formed.
	h2o.oDone <- struct{}{}

	// Step 5: Release the permit for this Oxygen atom.
	// This allows another Oxygen atom to start forming the next molecule.
	<-h2o.oSem // This makes the oSem available for the NEXT molecule.
}
