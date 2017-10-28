package bonzoline

import (
	"testing"
)

// A reader that just returns 0s
// In the test below we use this as an extremely bad RNG
type zeroReader struct{}

func (r zeroReader) Read(p []byte) (n int, err error) {
	for i, _ := range p {
		p[i] = 0
	}
	n = len(p)
	return
}

func TestIdentFromReader(t *testing.T) {
	if ident, err := identFromReader(zeroReader{}, 12); err != nil {
		t.Errorf("identFromReader failed: %s", err)
	} else if len(ident) != 12 {
		t.Errorf("identFromReader short: %s", ident)
	} else if ident != "000000000000" {
		t.Errorf("identFromReader wrong: '%s", ident)
	}
}

func TestRandomIdent(t *testing.T) {
	if ident, err := RandomIdent(12); err != nil {
		t.Errorf("RandomIdent failed: %s", err)
	} else if len(ident) != 12 {
		t.Errorf("RandomIdent short: %s", ident)
	} else if string(ident[:6]) == "000000" {
		t.Errorf("ident doesn't look random: %s", ident)
	} else if string(ident[6:12]) == "000000" {
		t.Errorf("ident doesn't look random: %s", ident)
	}
}
