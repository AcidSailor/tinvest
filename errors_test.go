package tinvest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrTInvestClient(t *testing.T) {
	assert.NotNil(t, ErrTInvestClient)
	assert.Equal(t, "tinvest client", ErrTInvestClient.Error())
}

func TestSubSentinelErrors(t *testing.T) {
	assert.Equal(t, "nil", ErrNil.Error())
	assert.Equal(t, "invalid config", ErrInvalidConfig.Error())
	assert.Equal(t, "overflow", ErrOverflow.Error())
	assert.Equal(t, "conversion", ErrConversion.Error())
}
