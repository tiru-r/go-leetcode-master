package building_h2o_1117

type H2O struct {
	hSem chan struct{}
	oSem chan struct{}
}

func NewH2O() *H2O {
	return &H2O{
		hSem: make(chan struct{}, 2),
		oSem: make(chan struct{}, 1),
	}
}

func (h *H2O) Hydrogen(releaseHydrogen func()) {
	h.hSem <- struct{}{}
	releaseHydrogen()

	// wait until all three atoms are ready
	if len(h.hSem) == 2 && len(h.oSem) == 1 {
		// all slots filled → molecule complete
		<-h.hSem // consume two H permits
		<-h.hSem
		<-h.oSem // consume one O permit
	}
}

func (h *H2O) Oxygen(releaseOxygen func()) {
	h.oSem <- struct{}{}
	releaseOxygen()

	// same wait logic as above
	if len(h.hSem) == 2 && len(h.oSem) == 1 {
		<-h.hSem
		<-h.hSem
		<-h.oSem
	}
}
