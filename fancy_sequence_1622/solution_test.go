package fancy_sequence_1622

import "testing"

func TestFancy(t *testing.T) {
	fancy := Constructor()
	
	fancy.Append(2)
	fancy.AddAll(3)
	fancy.Append(7)
	fancy.MultAll(2)
	
	if got := fancy.GetIndex(0); got != 10 {
		t.Errorf("GetIndex(0) = %d; want 10", got)
	}
	
	fancy.AddAll(3)
	
	if got := fancy.GetIndex(0); got != 13 {
		t.Errorf("GetIndex(0) = %d; want 13", got)
	}
	
	if got := fancy.GetIndex(1); got != 17 {
		t.Errorf("GetIndex(1) = %d; want 17", got)
	}
	
	if got := fancy.GetIndex(2); got != -1 {
		t.Errorf("GetIndex(2) = %d; want -1", got)
	}
}

func TestFancyLargeNumbers(t *testing.T) {
	fancy := Constructor()
	
	fancy.Append(1000000000)
	fancy.MultAll(1000000000)
	fancy.AddAll(1000000000)
	
	result := fancy.GetIndex(0)
	if result < 0 || result >= MOD {
		t.Errorf("Result should be in range [0, %d), got %d", MOD, result)
	}
}

func TestFancyEmpty(t *testing.T) {
	fancy := Constructor()
	
	if got := fancy.GetIndex(0); got != -1 {
		t.Errorf("GetIndex(0) on empty = %d; want -1", got)
	}
}

func BenchmarkFancy(b *testing.B) {
	fancy := Constructor()
	
	b.ResetTimer()
	for range b.N {
		fancy.Append(1)
		fancy.AddAll(1)
		fancy.MultAll(2)
		fancy.GetIndex(0)
	}
}
