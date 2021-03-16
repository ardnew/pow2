package pow2

import (
	"math/rand"
	"testing"
)

var testCase = map[Int]bool{}

func init() {
	for i, p := Int(0), Int(1); i < Bits; i, p = i+1, p<<1 {
		// add entries: (2^i)-1, (2^i)+1
		//  but only if they don't exist already
		pm, pp := p-1, p+1
		if _, ok := testCase[pm]; !ok {
			testCase[pm] = false
		}
		if _, ok := testCase[pp]; !ok {
			testCase[pp] = false
		}
		// add our current power: 2^i
		testCase[p] = true
	}
	// initialize the RNG with a crypto source
	rand.Seed(seed())
}

// -----------------------------------------------------------------------------
//  UNIT TESTS
// -----------------------------------------------------------------------------

func TestAndMinus1(t *testing.T) {
	for k, v := range testCase {
		if e := AndMinus1(k); e != v {
			t.Errorf("AndMinus1(%d) = %v, expected = %v", k, e, v)
		}
	}
}

func TestOnesCount(t *testing.T) {
	for k, v := range testCase {
		if e := OnesCount(k); e != v {
			t.Errorf("OnesCount(%d) = %v, expected = %v", k, e, v)
		}
	}
}

func TestForCountAll(t *testing.T) {
	for k, v := range testCase {
		if e := ForCountAll(k); e != v {
			t.Errorf("ForCountAll(%d) = %v, expected = %v", k, e, v)
		}
	}
}

func TestForCountOnce(t *testing.T) {
	for k, v := range testCase {
		if e := ForCountOnce(k); e != v {
			t.Errorf("ForCountOnce(%d) = %v, expected = %v", k, e, v)
		}
	}
}

func TestInSlice(t *testing.T) {
	for k, v := range testCase {
		if e := InSlice(k); e != v {
			t.Errorf("InSlice(%d) = %v, expected = %v", k, e, v)
		}
	}
}

func TestInMap(t *testing.T) {
	for k, v := range testCase {
		if e := InMap(k); e != v {
			t.Errorf("InMap(%d) = %v, expected = %v", k, e, v)
		}
	}
}

// -----------------------------------------------------------------------------
//  BENCHMARKS
// -----------------------------------------------------------------------------

// TestsPerIteration is the number of test cases to run per benchmark iteration.
var TestsPerIteration = 1000

// benchmarkResult is an artificial store to prevent unwanted optimizations in
// the benchmarks.
var benchmarkResult int

func BenchmarkAndMinus1(b *testing.B) {
	// always run benchmark across a different range to minimize cache hits
	val := Int(rand.Int63())
	res := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for n := 0; n < TestsPerIteration; n++ {
			if val++; AndMinus1(val) {
				res++
			}
		}
	}
	benchmarkResult = res
}

func BenchmarkOnesCount(b *testing.B) {
	// always run benchmark across a different range to minimize cache hits
	val := Int(rand.Int63())
	res := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for n := 0; n < TestsPerIteration; n++ {
			if val++; OnesCount(val) {
				res++
			}
		}
	}
	benchmarkResult = res
}

func BenchmarkForCountAll(b *testing.B) {
	// always run benchmark across a different range to minimize cache hits
	val := Int(rand.Int63())
	res := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for n := 0; n < TestsPerIteration; n++ {
			if val++; ForCountAll(val) {
				res++
			}
		}
	}
	benchmarkResult = res
}

func BenchmarkForCountOnce(b *testing.B) {
	// always run benchmark across a different range to minimize cache hits
	val := Int(rand.Int63())
	res := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for n := 0; n < TestsPerIteration; n++ {
			if val++; ForCountOnce(val) {
				res++
			}
		}
	}
	benchmarkResult = res
}

func BenchmarkInSlice(b *testing.B) {
	// always run benchmark across a different range to minimize cache hits
	val := Int(rand.Int63())
	res := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for n := 0; n < TestsPerIteration; n++ {
			if val++; InSlice(val) {
				res++
			}
		}
	}
	benchmarkResult = res
}

func BenchmarkInMap(b *testing.B) {
	// always run benchmark across a different range to minimize cache hits
	val := Int(rand.Int63())
	res := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for n := 0; n < TestsPerIteration; n++ {
			if val++; InMap(val) {
				res++
			}
		}
	}
	benchmarkResult = res
}
