package bonzoline

import (
	"crypto/rand"
	"io"
)

// The alphabet to use
var alphabet = []byte("0123456789abcdefghijklmnopqrstuvwxyz")

// 36^6 < 2^32 < 36^7 so for each 32-bit word from the RNG, we can
// extract up to 6 symbols without bias.
const symbolsPerWord = 6

// Populate ident with a random lower-case alphanumeric ident
// string.
func identFromReader(rng io.Reader, n int) (ident string, err error) {
	// Get enough raw bytes
	rawBytes := make([]byte, (n+symbolsPerWord-1)/symbolsPerWord*4)
	var bytesRead int
	if bytesRead, err = rng.Read(rawBytes); err != nil {
		return
	}
	if bytesRead != len(rawBytes) {
		panic("RNG did not return enough bytes")
	}
	// Synthesize the ident
	left := 0
	var word uint32
	nb := 0
	identBytes := make([]byte, n)
	for i := 0; i < n; i++ {
		// Every 6 output bytes, we pick another word out of
		// the raw random data.
		// (We use little-endian so that the compiler can
		// optimize it to a single load. Empirically you need
		// | rather than + for this to work.)
		if left == 0 {
			word = uint32(rawBytes[nb+0])
			word |= uint32(rawBytes[nb+1]) << 8
			word |= uint32(rawBytes[nb+2]) << 16
			word |= uint32(rawBytes[nb+3]) << 24
			nb += 4
			left = symbolsPerWord
		}
		identBytes[i] = alphabet[word%36]
		word /= 36
		left -= 1
	}
	ident = string(identBytes)
	return
}

// Return a random lower-case ident string of length n.
func RandomIdent(n int) (s string, err error) {
	if s, err = identFromReader(rand.Reader, n); err != nil {
		return
	}
	return
}
