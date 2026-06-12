// Package rest error types are the shared restkit types, re-exported as aliases
// so callers match them with errors.As without importing restkit:
//
//	var re *rest.ResponseError
//	if errors.As(err, &re) { _ = re.StatusCode } // e.g. 401
//
//	var re *rest.RequestError
//	if errors.As(err, &re) { _ = re.Op }          // "send", "unmarshal", ...
//
// There are no sentinel errors: the typed error IS the category.
package rest

import "github.com/acidsailor/restkit"

// ResponseError is a non-2xx T-Invest REST response (status + raw body).
type ResponseError = restkit.ResponseError

// RequestError is a per-call failure; Op names the stage and Err wraps the cause.
type RequestError = restkit.RequestError

// ConfigError is invalid NewClient construction input.
type ConfigError = restkit.ConfigError

// RequestError.Op values, re-exported so callers can match the failed stage
// without importing restkit.
const (
	OpMarshal   = restkit.OpMarshal   // encoding the request body to JSON
	OpBuild     = restkit.OpBuild     // constructing the *http.Request
	OpHook      = restkit.OpHook      // a request hook returned an error
	OpSend      = restkit.OpSend      // the HTTP round-trip
	OpRead      = restkit.OpRead      // reading the response body
	OpUnmarshal = restkit.OpUnmarshal // decoding the 2xx body
)
