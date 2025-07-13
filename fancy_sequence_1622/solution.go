package fancy_sequence_1622

const mod = 1_000_000_007

type Fancy struct {
	vals []int
	add  int
	mult int
}

func NewFancy() *Fancy {
	return &Fancy{mult: 1}
}

func (f *Fancy) Append(val int) {
	inv := pow(f.mult, mod-2)
	adj := ((val-f.add)%mod + mod) % mod
	f.vals = append(f.vals, adj*inv%mod)
}

func (f *Fancy) AddAll(inc int) {
	f.add = (f.add + inc%mod + mod) % mod
}

func (f *Fancy) MultAll(m int) {
	f.mult = f.mult * m % mod
	f.add = f.add * m % mod
}

func (f *Fancy) GetIndex(idx int) int {
	if idx < 0 || idx >= len(f.vals) {
		return -1
	}
	return (f.vals[idx]*f.mult + f.add) % mod
}

func pow(base, exp int) int {
	result := 1
	base %= mod
	for exp > 0 {
		if exp&1 == 1 {
			result = result * base % mod
		}
		base = base * base % mod
		exp >>= 1
	}
	return result
}
