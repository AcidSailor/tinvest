package rest_test

import (
	"encoding/json"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/acidsailor/tinvest/spec"
)

// parsePackageFiles parses every non-test .go file in the package into ASTs.
func parsePackageFiles(t *testing.T) []*ast.File {
	t.Helper()

	matches, err := filepath.Glob("*.go")
	require.NoError(t, err)

	fset := token.NewFileSet()
	var files []*ast.File
	for _, name := range matches {
		if strings.HasSuffix(name, "_test.go") {
			continue
		}
		file, err := parser.ParseFile(fset, name, nil, 0)
		require.NoError(t, err)
		files = append(files, file)
	}
	return files
}

// specUnaryPaths returns every non-streaming gateway path the upstream contract
// exposes. Streaming services are out of scope for the unary REST client, which
// models only request/response calls.
func specUnaryPaths(t *testing.T) map[string]bool {
	t.Helper()

	var doc struct {
		Paths map[string]map[string]json.RawMessage `json:"paths"`
	}
	require.NoError(t, json.Unmarshal(spec.SpecDerefJSON, &doc))

	paths := map[string]bool{}
	for path := range doc.Paths {
		if strings.Contains(path, "StreamService/") {
			continue
		}
		paths[path] = true
	}
	return paths
}

// invokedRESTPaths parses the package and returns every gateway path actually
// dispatched through a do[...] call, resolved from the path constant each call
// references. A path constant with no method binding never reaches a do[...]
// call, so it is not reported as implemented — this is what makes the test
// stronger than a plain source substring match.
func invokedRESTPaths(t *testing.T) map[string]bool {
	t.Helper()

	files := parsePackageFiles(t)

	constPath := map[string]string{} // const name -> gateway path literal
	for _, file := range files {
		for _, decl := range file.Decls {
			gd, ok := decl.(*ast.GenDecl)
			if !ok || gd.Tok != token.CONST {
				continue
			}
			for _, s := range gd.Specs {
				vs := s.(*ast.ValueSpec)
				for i, name := range vs.Names {
					if i >= len(vs.Values) {
						continue
					}
					if lit, ok := vs.Values[i].(*ast.BasicLit); ok &&
						lit.Kind == token.STRING {
						constPath[name.Name] = strings.Trim(lit.Value, `"`)
					}
				}
			}
		}
	}

	invoked := map[string]bool{}
	for _, file := range files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok || !isDoCall(call.Fun) {
				return true
			}
			for _, arg := range call.Args {
				id, ok := arg.(*ast.Ident)
				if !ok {
					continue
				}
				if p, ok := constPath[id.Name]; ok &&
					strings.HasPrefix(p, "/tinkoff") {
					invoked[p] = true
				}
			}
			return true
		})
	}
	return invoked
}

// isDoCall reports whether fun is a call to the generic do[...] dispatch helper.
func isDoCall(fun ast.Expr) bool {
	switch f := fun.(type) {
	case *ast.IndexExpr: // do[T](...)
		id, ok := f.X.(*ast.Ident)
		return ok && id.Name == "do"
	case *ast.IndexListExpr: // do[T, U](...) — defensive
		id, ok := f.X.(*ast.Ident)
		return ok && id.Name == "do"
	default:
		return false
	}
}

// TestUnaryEndpointsMatchSpec asserts the hand-written REST surface is in exact
// parity with the upstream contract: every non-streaming spec operation is
// reachable through a real method (a do[...] dispatch), and every dispatched
// path exists in the spec (catching typos and removed-upstream endpoints).
func TestUnaryEndpointsMatchSpec(t *testing.T) {
	specPaths := specUnaryPaths(t)
	invoked := invokedRESTPaths(t)

	var missing []string
	for p := range specPaths {
		if !invoked[p] {
			missing = append(missing, p)
		}
	}
	slices.Sort(missing)
	require.Emptyf(
		t,
		missing,
		"%d unary spec endpoints have no method binding:\n%s",
		len(missing),
		strings.Join(missing, "\n"),
	)

	var stale []string
	for p := range invoked {
		if !specPaths[p] {
			stale = append(stale, p)
		}
	}
	slices.Sort(stale)
	require.Emptyf(
		t,
		stale,
		"%d dispatched endpoints are absent from the spec:\n%s",
		len(stale),
		strings.Join(stale, "\n"),
	)
}
