package bonzoline

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestLocalHostname(t *testing.T) {
	var s string
	var err error
	s, err = LocalHostname()
	assert.NoError(t, err)
	assert.NotEqual(t, 0, len(s))
	assert.Equal(t, true, strings.Contains(s, "."))
}
