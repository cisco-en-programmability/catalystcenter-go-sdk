package main

import (
	"fmt"
	"io"
	"os"

	catalyst "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"
)

// Client is Catalyst Center API client
var client *catalyst.Client

func main() {
	var err error
	fmt.Println("Authenticating")
	client, err = catalyst.NewClientWithOptions("https://192.168.196.2/",
		"altus", "Altus123",
		"false", "false", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	f, err := os.Open("dnac_2-2-3-3.zip")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err = f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	var r io.Reader
	r = f

	nResponse, _, err := client.SoftwareImageManagementSwim.ImportLocalSoftwareImage(
		&catalyst.ImportLocalSoftwareImageQueryParams{},
		&catalyst.ImportLocalSoftwareImageMultipartFields{
			File:     r,
			FileName: "dnac_2-2-3-3_v2.zip",
		},
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	if nResponse.Response != nil {

		fmt.Println(nResponse.Response.TaskID)
	}

}
