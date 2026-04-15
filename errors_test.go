package tinvest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrClient(t *testing.T) {
	assert.NotNil(t, ErrClient)
	assert.Equal(t, "tinvest client", ErrClient.Error())
}

func TestSubSentinelErrors(t *testing.T) {
	assert.Equal(t, "nil", ErrNil.Error())
	assert.Equal(t, "invalid config", ErrInvalidConfig.Error())
	assert.Equal(t, "overflow", ErrOverflow.Error())
	assert.Equal(t, "conversion", ErrConversion.Error())
}
