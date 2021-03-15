package pow2

import "math/bits"

type (
	// Evaluate 64-bit powers of 2
	Int uint64

	IsPow2Func = func(n Int) bool
)

const Bits = Int(64)

// Various containers holding all powers of 2
var (
	powSlice = make([]Int, 0, Bits)
	powMap   = make(map[Int]bool, Bits)
)

func init() {
	// Initialize the powers of 2 containers
	for i, p := Int(0), Int(1); i < Bits; i, p = i+1, p<<1 {
		powSlice = append(powSlice, p)
		powMap[p] = true
	}
}

func AndMinus1(n Int) bool {
	return n != 0 && n&(n-1) == 0
}

func OnesCount(n Int) bool {
	return bits.OnesCount64(uint64(n)) == 1
}

func ForCountAll(n Int) bool {
	s := Int(0)
	for n > 0 {
		s += n & 1
		n >>= 1
	}
	return s == 1
}

func ForCountOnce(n Int) bool {
	for n > 0 {
		if 0 != n&1 {
			return 0 == n>>1
		}
		n >>= 1
	}
	return false
}

func InSlice(n Int) bool {
	for _, p := range powSlice {
		if n == p {
			return true
		}
	}
	return false
}

func InMap(n Int) bool {
	return powMap[n]
}
