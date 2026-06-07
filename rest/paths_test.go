package rest_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/acidsailor/tinvest/spec"
)

func TestEveryUnaryPathHasAMethod(t *testing.T) {
	var doc struct {
		Paths map[string]map[string]json.RawMessage `json:"paths"`
	}
	require.NoError(t, json.Unmarshal(spec.SpecDerefJSON, &doc))

	files, err := filepath.Glob("*.go")
	require.NoError(t, err)
	var src strings.Builder
	for _, f := range files {
		b, err := os.ReadFile(f)
		require.NoError(t, err)
		src.Write(b)
	}
	source := src.String()

	var missing []string
	for path := range doc.Paths {
		if strings.Contains(path, "StreamService/") {
			continue // streaming is out of scope
		}
		if !strings.Contains(source, `"`+path+`"`) {
			missing = append(missing, path)
		}
	}
	require.Emptyf(t, missing,
		"%d gateway paths have no method/path const:\n%s",
		len(missing), strings.Join(missing, "\n"))
}
