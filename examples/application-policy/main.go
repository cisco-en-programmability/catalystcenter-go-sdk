package main

import (
	"fmt"

	catalyst "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"
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
	// client.SetCatalystWaitTimeToManyRequest(2)
	for i := 0; i < 10; i++ {
		nResponse, _, err := client.ApplicationPolicy.GetApplicationsCount()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(nResponse.Response)
	}

}
