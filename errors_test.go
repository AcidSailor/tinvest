package tinvest

import (
	"testing"

	"github.com/acidsailor/tinvest/money"
	"github.com/stretchr/testify/assert"
)

func TestSentinelErrors(t *testing.T) {
	assert.Equal(t, "overflow", money.ErrOverflow.Error())
	assert.Equal(t, "conversion", money.ErrConversion.Error())
}
