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
	queryParams1 := catalyst.GetSyslogEventSubscriptionsQueryParams{}
	queryParams1.Name = "Test Terraform 2"
	nResponse, _, err := client.EventManagement.GetSyslogEventSubscriptions(&queryParams1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(nResponse)

}
