package fancy_sequence_1622

const MOD = 1_000_000_007

type Fancy struct {
	arr  []int
	add  int
	mult int
}

func Constructor() Fancy {
	return Fancy{
		mult: 1,
	}
}

func (f *Fancy) Append(val int) {
	invMult := modInverse(f.mult)
	val = (val - f.add + MOD) % MOD
	val = (val * invMult) % MOD
	f.arr = append(f.arr, val)
}

func (f *Fancy) AddAll(inc int) {
	f.add = (f.add + inc) % MOD
}

func (f *Fancy) MultAll(m int) {
	f.mult = (f.mult * m) % MOD
	f.add = (f.add * m) % MOD
}

func (f *Fancy) GetIndex(idx int) int {
	if idx >= len(f.arr) {
		return -1
	}
	val := f.arr[idx]
	val = (val * f.mult) % MOD
	val = (val + f.add) % MOD
	return val
}

func modInverse(a int) int {
	return modPow(a, MOD-2)
}

func modPow(base, exp int) int {
	result := 1
	base %= MOD
	for exp > 0 {
		if exp&1 == 1 {
			result = (result * base) % MOD
		}
		base = (base * base) % MOD
		exp >>= 1
	}
	return result
}
