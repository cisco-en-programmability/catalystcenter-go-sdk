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
	nResponse, _, err := client.Task.GetBusinessAPIExecutionDetails("a919fe4c-70c2-4023-a063-404e2705c277")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(nResponse)

}
