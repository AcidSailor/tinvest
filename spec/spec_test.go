package spec_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/acidsailor/tinvest/spec"
)

func TestSpecDerefJSON_NoResidualRef(t *testing.T) {
	require.NotEmpty(t, spec.SpecDerefJSON)
	// Full dereferencing is the contract: a residual $ref means a downstream
	// resolver would dangle.
	assert.False(
		t,
		bytes.Contains(spec.SpecDerefJSON, []byte(`"$ref"`)),
		`SpecDerefJSON contains "$ref"; expected a fully dereferenced document`,
	)
}
