package money_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/acidsailor/tinvest/money"
)

func TestUnitsNanoToDecimal(t *testing.T) {
	d, err := money.UnitsNanoToDecimal(114, 250000000)
	require.NoError(t, err)
	assert.Equal(t, "114.25", d.String())
}

func TestUnitsNanoToDecimal_MixedSign(t *testing.T) {
	_, err := money.UnitsNanoToDecimal(1, -5)
	require.Error(t, err)
	assert.ErrorIs(t, err, money.ErrConversion)
}

func TestDecimalToUnitsNano_RoundTrip(t *testing.T) {
	d, err := money.UnitsNanoToDecimal(-7, -500000000)
	require.NoError(t, err)
	u, n, err := money.DecimalToUnitsNano(d)
	require.NoError(t, err)
	assert.Equal(t, int64(-7), u)
	assert.Equal(t, int32(-500000000), n)
}

func TestFormatMoney(t *testing.T) {
	s, err := money.FormatMoney(250, 500000000, "rub")
	require.NoError(t, err)
	assert.Equal(t, "250.50 RUB", s)
}

func TestFormatQuotation_NoNano(t *testing.T) {
	s, err := money.FormatQuotation(89, 0)
	require.NoError(t, err)
	assert.Equal(t, "89", s)
}

func TestNormalizeSign_IntMinErrors(t *testing.T) {
	_, _, _, err := money.NormalizeSign(0, math.MinInt32)
	require.Error(t, err)
	assert.ErrorIs(t, err, money.ErrOverflow)

	_, _, _, err = money.NormalizeSign(math.MinInt64, 0)
	require.Error(t, err)
	assert.ErrorIs(t, err, money.ErrOverflow)
}
