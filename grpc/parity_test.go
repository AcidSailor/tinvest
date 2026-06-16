package grpc_test

import (
	"encoding/json"
	"reflect"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	tgrpc "github.com/acidsailor/tinvest/grpc"
	"github.com/acidsailor/tinvest/spec"
)

// protoContractPkg is the protobuf package shared by every T-Invest service.
// Gateway and gRPC method paths are "/<protoContractPkg><Service>/<Method>",
// the same namespace the REST spec paths use — which is what lets the two
// transports be checked against one source of truth.
const protoContractPkg = "tinkoff.public.invest.api.contract.v1."

// specUnaryPaths returns every non-streaming gateway path the upstream contract
// exposes. Streaming services are out of scope for both the unary REST client
// and this gRPC parity check.
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

// grpcUnaryPaths reflects over the exposed Client service fields and returns the
// gateway path of every unary RPC. A unary RPC returns a *Response pointer; a
// streaming RPC returns a stream-client interface instead, which is how the two
// are told apart without hard-coding method names.
func grpcUnaryPaths(t *testing.T) map[string]bool {
	t.Helper()

	paths := map[string]bool{}
	for field := range reflect.TypeFor[tgrpc.Client]().Fields() {
		if field.Type.Kind() != reflect.Interface {
			continue
		}
		// pb.InstrumentsServiceClient -> InstrumentsService
		service := strings.TrimSuffix(field.Type.Name(), "Client")
		for method := range field.Type.Methods() {
			sig := method.Type
			if sig.NumOut() == 0 || sig.Out(0).Kind() != reflect.Pointer {
				continue // streaming RPC: first result is a stream-client interface
			}
			paths["/"+protoContractPkg+service+"/"+method.Name] = true
		}
	}
	return paths
}

// TestGRPCClientMatchesSpec asserts the gRPC client exposes exactly the unary
// operations the upstream contract defines — the same parity the REST client is
// held to (TestUnaryEndpointsMatchSpec), so the two transports never drift apart
// in which unary endpoints they cover. Streaming RPCs are excluded by design.
func TestGRPCClientMatchesSpec(t *testing.T) {
	specPaths := specUnaryPaths(t)
	grpcPaths := grpcUnaryPaths(t)

	var missing []string
	for p := range specPaths {
		if !grpcPaths[p] {
			missing = append(missing, p)
		}
	}
	slices.Sort(missing)
	require.Emptyf(
		t,
		missing,
		"%d unary spec operations are not exposed by the gRPC client:\n%s",
		len(missing),
		strings.Join(missing, "\n"),
	)

	var extra []string
	for p := range grpcPaths {
		if !specPaths[p] {
			extra = append(extra, p)
		}
	}
	slices.Sort(extra)
	require.Emptyf(t, extra, "%d gRPC unary RPCs are absent from the spec:\n%s",
		len(extra), strings.Join(extra, "\n"))
}
