package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	cc "golang.org/x/oauth2/clientcredentials"
)

type SpecialClient struct {
	*http.Client
}

const (
	issuer = "http://localhost:1323/"
)

var (
	provider       *oidc.Provider
	oauth2Endpoint oauth2.Endpoint
)

func main() {
	ctx := context.Background()
	var err error
	provider, err = oidc.NewProvider(ctx, issuer)
	if err != nil {
		panic(err)
	}
	oauth2Endpoint = provider.Endpoint()

	client := NewClient(
		"b2b-client",
		"secret",
	)

	// the client will update its token if it's expired
	resp, err := client.Get("http://localhost:9888/")

	if err != nil {
		panic(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	// If response code is 200 it was successful
	if resp.StatusCode == 200 {
		fmt.Println("The request was successful. Response below:")
		fmt.Println(string(body))
	} else {
		fmt.Println("Could not perform request to the endpoint. Response below:")
		fmt.Println(string(body))
	}
}

func NewClient(cid, csec string) *SpecialClient {

	// this should match whatever service has given you
	// client credential access
	config := &cc.Config{
		ClientID:     cid,
		ClientSecret: csec,
		TokenURL:     oauth2Endpoint.TokenURL,
		Scopes:       []string{"artificer-ns", "a", "b", "c", "invoices", "users.read"},
	}

	// you can modify the client (for example ignoring bad certs or otherwise)
	// by modifying the context
	ctx := context.Background()
	client := config.Client(ctx)

	return &SpecialClient{client}
}
