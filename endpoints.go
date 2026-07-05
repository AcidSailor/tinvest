package tinvest

const (
	// AppName is the default x-app-name header value identifying this client library.
	AppName = "github.com/acidsailor/tinvest"
	// EndpointProduction is the T-Invest live trading API endpoint.
	EndpointProduction = "invest-public-api.tbank.ru:443"
	// EndpointSandbox is the T-Invest sandbox API endpoint for testing without real money.
	EndpointSandbox = "sandbox-invest-public-api.tbank.ru:443"
	// EndpointProductionREST is the T-Invest REST gateway (live trading).
	EndpointProductionREST = "https://" + EndpointProduction + "/rest"
	// EndpointSandboxREST is the T-Invest REST gateway sandbox.
	EndpointSandboxREST = "https://" + EndpointSandbox + "/rest"
)
