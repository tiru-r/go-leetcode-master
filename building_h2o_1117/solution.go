package building_h2o_1117

type H2O struct {
	h chan func()
	o chan func()
}

func New() *H2O {
	h2o := &H2O{
		h: make(chan func()),
		o: make(chan func()),
	}
	
	go func() {
		for {
			h1 := <-h2o.h
			h2 := <-h2o.h
			o := <-h2o.o
			
			h1()
			h2()
			o()
		}
	}()
	
	return h2o
}

func (h *H2O) Hydrogen(releaseHydrogen func()) {
	done := make(chan struct{})
	h.h <- func() {
		releaseHydrogen()
		done <- struct{}{}
	}
	<-done
}

func (h *H2O) Oxygen(releaseOxygen func()) {
	done := make(chan struct{})
	h.o <- func() {
		releaseOxygen()
		done <- struct{}{}
	}
	<-done
}
