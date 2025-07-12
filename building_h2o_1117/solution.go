package building_h2o_1117

type H2O struct {
	h chan struct{}
	o chan struct{}
}

func New() *H2O {
	return &H2O{
		h: make(chan struct{}, 2),
		o: make(chan struct{}, 1),
	}
}

func (h *H2O) Hydrogen(releaseHydrogen func()) {
	h.h <- struct{}{} // enqueue H
	select {          // wait for a full molecule
	case <-h.o: // got O, need one more H
		<-h.h
		releaseHydrogen()
		releaseHydrogen()
	case <-h.h: // got second H, need O
		<-h.o
		releaseHydrogen()
		releaseHydrogen()
	}
}

func (h *H2O) Oxygen(releaseOxygen func()) {
	h.o <- struct{}{} // enqueue O
	<-h.h             // wait for two H
	<-h.h
	releaseOxygen()
}
