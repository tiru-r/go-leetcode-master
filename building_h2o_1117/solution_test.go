package building_h2o_1117

import (
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// ------------------------------------------------------------
// Helper: collect output in the order it is actually executed
// ------------------------------------------------------------
type collector struct {
	mu   sync.Mutex
	buf  strings.Builder
	atom int32 // running index for ordering
}

func (c *collector) write(ch byte) {
	atomic.AddInt32(&c.atom, 1)
	c.mu.Lock()
	defer c.mu.Unlock()
	c.buf.WriteByte(ch)
}

func (c *collector) String() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.buf.String()
}

// ------------------------------------------------------------
// Functional correctness: small deterministic inputs
// ------------------------------------------------------------
func TestH2O_small(t *testing.T) {
	cases := []struct {
		input string // H=H, O=O
	}{
		{"HOH"},
		{"OOHHHH"},
		{"HHHOOOHHHHHH"},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			h2o := New()
			var wg sync.WaitGroup
			out := &collector{}

			for _, ch := range tc.input {
				wg.Add(1)
				switch ch {
				case 'H':
					go func() {
						defer wg.Done()
						h2o.Hydrogen(func() { out.write('H') })
					}()
				case 'O':
					go func() {
						defer wg.Done()
						h2o.Oxygen(func() { out.write('O') })
					}()
				}
			}
			wg.Wait()

			got := out.String()
			if len(got) != len(tc.input) {
				t.Fatalf("wrong length: got %d, want %d", len(got), len(tc.input))
			}

			// validate 2H + 1O per molecule
			for i := 0; i < len(got); i += 3 {
				if i+3 > len(got) {
					t.Fatalf("molecule %d incomplete", i/3)
				}
				mol := got[i : i+3]
				h, o := 0, 0
				for _, c := range mol {
					switch c {
					case 'H':
						h++
					case 'O':
						o++
					}
				}
				if h != 2 || o != 1 {
					t.Fatalf("wrong molecule %q at %d", mol, i)
				}
			}
		})
	}
}

// ------------------------------------------------------------
// Stress & deadlock test: 100 k molecules (300 k goroutines)
// ------------------------------------------------------------
func TestH2O_stress(t *testing.T) {
	const molecules = 100_000
	h2o := New()
	var wg sync.WaitGroup
	out := &collector{}

	for i := 0; i < molecules; i++ {
		wg.Add(3)
		go func() {
			defer wg.Done()
			h2o.Hydrogen(func() { out.write('H') })
		}()
		go func() {
			defer wg.Done()
			h2o.Hydrogen(func() { out.write('H') })
		}()
		go func() {
			defer wg.Done()
			h2o.Oxygen(func() { out.write('O') })
		}()
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		// ok
	case <-time.After(10 * time.Second):
		t.Fatal("deadlock or too slow")
	}

	if len(out.String()) != molecules*3 {
		t.Fatalf("wrong total length after stress")
	}
}

// ------------------------------------------------------------
// Goroutine leak check (optional but cheap)
// ------------------------------------------------------------
func TestH2O_goroutines(t *testing.T) {
	before := runtime.NumGoroutine()
	TestH2O_small(t) // reuse correctness test
	after := runtime.NumGoroutine()
	if after-before > 5 { // allow small scheduling jitter
		t.Errorf("possible goroutine leak: %d -> %d", before, after)
	}
}
