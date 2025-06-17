package main

import (
	"fmt"

	catalyst "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"
)

// Client is Catalyst Center API client
var client *catalyst.Client

func main() {
	var err error
	fmt.Println("Authenticating")
	client, err = catalyst.NewClientWithOptions("https://192.168.196.2/",
		"altus", "Altus123",
		"true", "false", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	getClientEnrichmentDetailsHeaderParams := &catalyst.GetClientEnrichmentDetailsHeaderParams{
		EntityType:    "network_user_id",
		EntityValue:   "test",
		IssueCategory: "test",
	}
	nResponse, _, err := client.Clients.GetClientEnrichmentDetails(getClientEnrichmentDetailsHeaderParams)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(nResponse)

}
