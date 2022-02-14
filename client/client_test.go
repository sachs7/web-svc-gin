package client

import (
	"errors"
	"fmt"
	"testing"
	"web-service-gin/server"

	"github.com/pact-foundation/pact-go/dsl"
)

func TestClientPact_Local(t *testing.T) {
	// initialize PACT DSL
	pact := dsl.Pact{
		Consumer: "example-client1",
		Provider: "example-server1",
	}

	// setup a PACT Mock Server
	pact.Setup(true)

	t.Run("get user by id", func(t *testing.T) {
		id := "1"

		pact.
			AddInteraction().                                 // specify PACT interaction
			Given("Album 'Blue Train' exists").               // specify Provider state
			UponReceiving("Album 'Blue Train' is requested"). // specify test case name
			WithRequest(dsl.Request{                          // specify expected request
				Method: "GET",
				Path:   dsl.Term("/albums/1", "/albums/[0-9]+"), // specify matching for endpoint
			}).
			WillRespondWith(dsl.Response{ // specify minimal expected response
				Status: 200,
				Body: dsl.Like(server.Album{ // specify matching for response body
					ID:     id,
					Title:  "Blue Train",
					Artist: "John Coltrane",
					// Price:  56.99,
				}),
			})

		// verify interaction on client side
		err := pact.Verify(func() error {
			// specify host and post of PACT Mock Server as actual server
			host := fmt.Sprintf("%s:%d", pact.Host, pact.Server.Port)

			// execute function
			album, err := GetUserByID(host, id)
			if err != nil {
				return errors.New("error is not expected")
			}

			// check if actual user is equal to expected
			if album == nil || album.ID != id {
				return fmt.Errorf("expected album with ID %s but got %v", id, album)
			}

			return err
		})

		if err != nil {
			t.Fatal(err)
		}
	})

	// write Contract into file
	if err := pact.WritePact(); err != nil {
		t.Fatal(err)
	}

	// stop PACT mock server
	pact.Teardown()
}
