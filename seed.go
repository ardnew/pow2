package pow2

import (
	"crypto/rand"
	"encoding/binary"
)

// binDecoder returns a function that will convert a given byte slice to int64.
// The length of the slice (the number of bytes), given as argument to the
// returned function, is expected to equal Bits/8.
func binDecoder() func([]byte) int64 {
	switch Bits {
	case 8:
		return func(b []byte) int64 { return int64(b[0]) }
	case 16:
		return func(b []byte) int64 { return int64(binary.LittleEndian.Uint16(b)) }
	case 32:
		return func(b []byte) int64 { return int64(binary.LittleEndian.Uint32(b)) }
	case 64:
		return func(b []byte) int64 { return int64(binary.LittleEndian.Uint64(b)) }
	}
	return nil
}

// seed returns a random, non-negative integer (width = Bits) as int64, making
// it a suitable argument for math/rand.Seed().
// Otherwise, a random integer could not be generated and -1 is returned.
func seed() int64 {
	const tryMax = 10
	b := make([]byte, Bits/8)
	for try := 0; try < tryMax; try++ {
		if _, err := rand.Read(b); nil == err {
			return binDecoder()(b)
		}
	}
	return -1
}
