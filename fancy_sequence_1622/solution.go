package fancy_sequence_1622

const mod = 1_000_000_007

type Fancy struct {
	values     []int
	addOffset  int
	multFactor int
}

func NewFancy() *Fancy {
	return &Fancy{
		multFactor: 1,
	}
}

func (f *Fancy) Append(val int) {
	// Store the original value adjusted for current operations
	invMult := modInverse(f.multFactor, mod)
	adjustedVal := ((val-f.addOffset)%mod + mod) % mod
	adjustedVal = (adjustedVal * invMult) % mod
	f.values = append(f.values, adjustedVal)
}

func (f *Fancy) AddAll(inc int) {
	f.addOffset = ((f.addOffset + inc) % mod + mod) % mod
}

func (f *Fancy) MultAll(m int) {
	f.multFactor = (f.multFactor * m) % mod
	f.addOffset = ((f.addOffset * m) % mod + mod) % mod
}

func (f *Fancy) GetIndex(idx int) int {
	if idx < 0 || idx >= len(f.values) {
		return -1
	}
	
	val := f.values[idx]
	val = (val * f.multFactor) % mod
	val = ((val + f.addOffset) % mod + mod) % mod
	return val
}

func modInverse(a, mod int) int {
	return modPow(a, mod-2, mod)
}

func modPow(base, exp, mod int) int {
	result := 1
	base %= mod
	for exp > 0 {
		if exp&1 == 1 {
			result = (result * base) % mod
		}
		base = (base * base) % mod
		exp >>= 1
	}
	return result
}
