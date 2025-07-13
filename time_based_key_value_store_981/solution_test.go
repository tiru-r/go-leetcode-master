package time_based_key_value_store_981

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

type operation struct {
	kind      string
	key       string
	value     string
	timestamp int
}

func TestAllInOne(t *testing.T) {
	tests := []struct {
		name     string
		ops      []operation
		expected []string
	}{
		{"empty store", []operation{
			{"get", "a", "", 1},
		}, []string{""}},

		{"single set/get", []operation{
			{"set", "a", "v1", 1},
			{"get", "a", "", 1},
			{"get", "a", "", 0},
			{"get", "a", "", 2},
		}, []string{"v1", "", "v1"}},

		{"duplicate timestamp overwrite", []operation{
			{"set", "a", "v1", 10},
			{"set", "a", "v2", 10},
			{"get", "a", "", 10},
		}, []string{"v2"}},

		{"non-monotonic timestamps", []operation{
			{"set", "a", "v1", 5},
			{"set", "a", "v2", 3},
			{"set", "a", "v3", 7},
			{"get", "a", "", 4},
			{"get", "a", "", 6},
			{"get", "a", "", 10},
		}, []string{"v2", "v1", "v3"}},

		{"eviction not required (per key list grows)", []operation{
			{"set", "k", "v1", 1},
			{"set", "k", "v2", 2},
			{"set", "k", "v3", 3},
			{"get", "k", "", 0},
			{"get", "k", "", 1},
			{"get", "k", "", 2},
			{"get", "k", "", 3},
			{"get", "k", "", 4},
		}, []string{"", "v1", "v2", "v3", "v3"}},

		{"multiple keys isolation", []operation{
			{"set", "a", "a1", 1},
			{"set", "b", "b1", 1},
			{"set", "a", "a2", 2},
			{"get", "b", "", 1},
			{"get", "b", "", 2},
			{"get", "a", "", 2},
		}, []string{"b1", "b1", "a2"}},

		{"large monotonic insert + binary search", func() []operation {
			var out []operation
			for i := range 1_000 {
				out = append(out, operation{"set", "big", "v" + strconv.Itoa(i+1), i+1})
			}
			for range 20 {
				ts := rand.Intn(2000) + 1
				out = append(out, operation{"get", "big", "", ts})
			}
			return out
		}(), nil},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tm := NewTimeMap()
			var got []string
			for _, o := range tc.ops {
				switch o.kind {
				case "set":
					tm.Set(o.key, o.value, o.timestamp)
				case "get":
					got = append(got, tm.Get(o.key, o.timestamp))
				}
			}

			if tc.expected != nil && !eq(got, tc.expected) {
				t.Fatalf("want %v, got %v", tc.expected, got)
			}

			// For the "large" case, do a spot check instead of full want slice.
			if tc.name == "large monotonic insert + binary search" {
				for ts := 1; ts <= 1_000; ts++ {
					if v := tm.Get("big", ts); v != "v"+strconv.Itoa(ts) {
						t.Errorf("big@%d: want %q, got %q", ts, "v"+strconv.Itoa(ts), v)
					}
				}
				if v := tm.Get("big", 0); v != "" {
					t.Errorf("big@0: want \"\", got %q", v)
				}
				if v := tm.Get("big", 2000); v != "v1000" {
					t.Errorf("big@2000: want \"v1000\", got %q", v)
				}
			}
		})
	}
}

func BenchmarkSet(b *testing.B) {
	tm := NewTimeMap()
	b.ResetTimer()
	for i := range b.N {
		tm.Set("bench", "v", i)
	}
}

func BenchmarkGet(b *testing.B) {
	const size = 100_000
	tm := NewTimeMap()
	for i := range size {
		tm.Set("bench", "v", i*2)
	}
	rand.Seed(time.Now().UnixNano())
	b.ResetTimer()
	for range b.N {
		ts := rand.Intn(size * 2)
		_ = tm.Get("bench", ts)
	}
}

// helper
func eq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
