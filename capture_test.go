package bonzoline

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCapture(t *testing.T) {
	var b []byte
	var err error
	b, err = Capture("echo", "whatever")
	assert.NoError(t, err)
	assert.Equal(t, []byte("whatever\n"), b)
}

func TestCaptureError(t *testing.T) {
	var b []byte
	var err error
	b, err = Capture("false")
	assert.Error(t, err)
	assert.Equal(t, 0, len(b))
}
