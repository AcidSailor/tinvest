// Package spec embeds the dereferenced T-Invest OpenAPI document used by the
// MCP server to assemble per-tool JSON schemas.
package spec

import _ "embed"

// SpecDerefJSON is the fully dereferenced T-Invest OpenAPI document. See the
// `deref` task. Every $ref is inlined while components.schemas is retained.
//
//go:embed spec-deref.json
var SpecDerefJSON []byte
