# pow2
#### Test if a number is a power of 2 (performance benchmarks)

## Methods

There are 6 methods evaluated to determine if `N` is a power of 2:

1. `AndMinusOne`: `return N != 0 and N&(N-1) == 0`
2. `OnesCount`: `return bits.OnesCount(N) == 1` (from standard library package `math/bits`)
3. `ForCountAll`: `while N > 0 { S+=1 if N&1 != 0; N >>= 1 }; return S == 1`
4. `ForCountOnce`: `while N > 0 { return N>>1 == 0 if N&1 != 0; N >>= 1 }; return false`
5. `InSlice`: `for each P in List[ 2**K | 0 <= K <= 64 ] { return true if N == P }; return false`
6. `InMap`: `return true if and only if N in Map[ { 2**K: true } | 0 <= K <= 64 ]`

## Results

These methods resulted in the following benchmarks using **Go 1.16**:

```
$ go test -bench .
goos: linux
goarch: amd64
pkg: github.com/ardnew/pow2
cpu: AMD Ryzen Threadripper 1950X 16-Core Processor
BenchmarkAndMinus1-32            2264751               515.9 ns/op
BenchmarkOnesCount-32            1714140               694.0 ns/op
BenchmarkForCountAll-32            39033             31690 ns/op
BenchmarkForCountOnce-32          581548              2096 ns/op
BenchmarkInSlice-32                25918             45959 ns/op
BenchmarkInMap-32                 100513             11501 ns/op
PASS
ok      github.com/ardnew/pow2  10.231s
```

## Usage

To check these results yourself using a different system, simply fetch this repo using `go get` or `git`, and run the test command from the root directory:

```
$ mkdir -p ardnew/pow2
$ git clone https://github.com/ardnew/pow2.git ardnew/pow2
$ cd ardnew/pow2
$ go test -bench .
```
