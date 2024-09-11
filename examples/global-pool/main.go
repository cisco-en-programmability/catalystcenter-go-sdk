package main

import (
	"fmt"

	catalyst "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"
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

	queryParams := &catalyst.GetGlobalPoolQueryParams{
		Offset: 1,
		Limit:  5,
	}

	nResponse, _, err := client.NetworkSettings.GetGlobalPool(queryParams)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(nResponse.Response)
}
