package server

import (
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

func TestServerPact_Verification(t *testing.T) {
	// initialize PACT DSL
	pact := dsl.Pact{
		Provider: "example-server",
	}

	// verify Contract on server side
	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL: "http://localhost:8082",
		PactURLs:        []string{"../client/pacts/example-client1-example-server1.json"},
	})

	if err != nil {
		t.Log(err)
	}
}
