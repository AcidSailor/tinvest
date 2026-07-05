package tinvest

import (
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// rootCASHA256 pins the embedded root to the copy published by Gosuslugi at
// https://gu-st.ru/content/Other/doc/russian_trusted_root_ca.cer — any change to
// certs/russian_trusted_root_ca.pem must be a verified re-fetch, not a swap.
const rootCASHA256 = "d26d2d0231b7c39f92cc738512" +
	"ba54103519e4405d68b5bd703e9788ca8ecf31"

func TestEmbeddedRootCA(t *testing.T) {
	block, _ := pem.Decode(russianTrustedRootCA)
	require.NotNil(t, block, "embedded PEM did not decode")

	cert, err := x509.ParseCertificate(block.Bytes)
	require.NoError(t, err)

	assert.Equal(t, "Russian Trusted Root CA", cert.Subject.CommonName)
	assert.True(t, cert.IsCA, "embedded cert is not a CA")

	sum := sha256.Sum256(cert.Raw)
	assert.Equal(t, rootCASHA256, hex.EncodeToString(sum[:]))
}

func TestRootCAs(t *testing.T) {
	pool, err := RootCAs()
	require.NoError(t, err)
	require.NotNil(t, pool)
}
