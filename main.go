package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	httptransport "github.com/go-openapi/runtime/client"
	apiclient "github.com/uempfel/tkgi-client-go/client"
	"github.com/uempfel/tkgi-client-go/client/uaa"
)

func main() {

	host := os.Getenv("TKGI_HOSTNAME")
	if host == "" {
		log.Fatal("environment variable TKGI_HOSTNAME not set")
	}

	// export TKGI_REFRESH_TOKEN=$(yq -r .refresh_token < ~/.pks/creds.yml)
	refreshToken := os.Getenv("TKGI_REFRESH_TOKEN")
	if refreshToken == "" {
		log.Fatal("environment variable TKGI_REFRESH_TOKEN not set")
	}
	
	//fmt.Println(uaa.GetTokenRemainingValidity(refreshToken))	
	accesToken := uaa.RenewAccessToken(host, refreshToken)
	client := apiclient.NewHTTPClientWithConfig(nil, apiclient.DefaultTransportConfig().WithHost(fmt.Sprintf("%s:9021", host)))

	bearerTokenAuth := httptransport.BearerToken(accesToken)
	re, err := client.Cluster.ListClusters(nil, bearerTokenAuth)
	if err != nil {
		log.Fatal(err)
	}
	c, err := json.Marshal(re.Payload)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(c))

}
