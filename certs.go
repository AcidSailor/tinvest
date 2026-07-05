package tinvest

import (
	"crypto/x509"
	_ "embed"
	"fmt"
)

// russianTrustedRootCA is the PEM-encoded "Russian Trusted Root CA" operated by
// the Russian Ministry of Digital Development (НУЦ Минцифры РФ). The T-Invest
// API endpoints serve TLS certificates that chain to it, and it is absent from
// standard OS and browser trust stores — so TLS verification against the API
// fails ("certificate verify failed") without it.
//
// Verified byte-identical to the copy published by Gosuslugi at
// https://gu-st.ru/content/Other/doc/russian_trusted_root_ca.cer
// SHA-256 d26d2d0231b7c39f92cc738512ba54103519e4405d68b5bd703e9788ca8ecf31
//
//go:embed certs/russian_trusted_root_ca.pem
var russianTrustedRootCA []byte

// RootCAs returns a certificate pool containing the host's system roots plus the
// Russian Trusted Root CA that the T-Invest API endpoints chain to (see the
// note on russianTrustedRootCA). Both transports use it to configure TLS.
//
// Appending is additive: it grants trust for this one verified government root
// only, on connections made through this pool, and never mutates the
// process-wide system trust store.
func RootCAs() (*x509.CertPool, error) {
	pool, err := x509.SystemCertPool()
	if err != nil {
		return nil, fmt.Errorf("tinvest: load system cert pool: %w", err)
	}
	if pool == nil {
		pool = x509.NewCertPool()
	}
	if !pool.AppendCertsFromPEM(russianTrustedRootCA) {
		return nil, fmt.Errorf("tinvest: append Russian Trusted Root CA")
	}
	return pool, nil
}
