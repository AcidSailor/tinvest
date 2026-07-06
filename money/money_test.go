package money_test

import (
	"math"
	"testing"

	"github.com/quagmt/udecimal"
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

func TestUnitsNanoToDecimal_NanoOutOfRange(t *testing.T) {
	// A nano at/above one whole unit is not a valid split; it must be rejected
	// as ErrOverflow rather than silently yielding units+1.5 = wrong value.
	_, err := money.UnitsNanoToDecimal(1, 1_500_000_000)
	require.Error(t, err)
	assert.ErrorIs(t, err, money.ErrOverflow)

	require.Error(t, money.ValidateSigns(0, 1_000_000_000))
	require.Error(t, money.ValidateSigns(0, -1_000_000_000))
	require.NoError(t, money.ValidateSigns(0, 999_999_999))
}

func TestDecimalToUnitsNano_RoundTrip(t *testing.T) {
	d, err := money.UnitsNanoToDecimal(-7, -500000000)
	require.NoError(t, err)
	u, n, err := money.DecimalToUnitsNano(d)
	require.NoError(t, err)
	assert.Equal(t, int64(-7), u)
	assert.Equal(t, int32(-500000000), n)
}

func TestDecimalToUnitsNano_ExcessPrecision(t *testing.T) {
	// 10 fractional digits cannot round-trip through units/nano (9 digits) and
	// must be rejected rather than silently truncated.
	d, err := udecimal.Parse("0.0000000001")
	require.NoError(t, err)
	_, _, err = money.DecimalToUnitsNano(d)
	require.Error(t, err)
	assert.ErrorIs(t, err, money.ErrOverflow)
}

func TestFormatMoney(t *testing.T) {
	s, err := money.FormatMoney(250, 500000000, "rub")
	require.NoError(t, err)
	assert.Equal(t, "250.50 RUB", s)
}

func TestFormatMoney_Negative(t *testing.T) {
	s, err := money.FormatMoney(-250, -500000000, "usd")
	require.NoError(t, err)
	assert.Equal(t, "-250.50 USD", s)
}

func TestFormatQuotation_NoNano(t *testing.T) {
	s, err := money.FormatQuotation(89, 0)
	require.NoError(t, err)
	assert.Equal(t, "89", s)
}

func TestFormatQuotation_WithNano(t *testing.T) {
	s, err := money.FormatQuotation(89, 250_000_000)
	require.NoError(t, err)
	assert.Equal(t, "89.25", s)

	s, err = money.FormatQuotation(-1, -50_000_000)
	require.NoError(t, err)
	assert.Equal(t, "-1.05", s)
}

func TestNormalizeSign_IntMinErrors(t *testing.T) {
	_, _, _, err := money.NormalizeSign(0, math.MinInt32)
	require.Error(t, err)
	assert.ErrorIs(t, err, money.ErrOverflow)

	_, _, _, err = money.NormalizeSign(math.MinInt64, 0)
	require.Error(t, err)
	assert.ErrorIs(t, err, money.ErrOverflow)
}

func TestSentinelErrors(t *testing.T) {
	assert.Equal(t, "overflow", money.ErrOverflow.Error())
	assert.Equal(t, "conversion", money.ErrConversion.Error())
}
