package main

import (
	"fmt"

	catalyst "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"
)

// Client is Catalyst Center API client
var Client *catalyst.Client

func main() {
	fmt.Println("Authenticating")
	var err error
	Client, err = catalyst.NewClientWithOptions("https://100.119.103.190",
		"cloverhound_user", "LABchsys!23$",
		"true", "false", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	resourcePath := fmt.Sprintf("%s/api/v1/siteprofile", Client.RestyClient().BaseURL)
	a, err := Client.CustomCall.GetCustomCall(resourcePath, nil)

	print(a.String())

}
