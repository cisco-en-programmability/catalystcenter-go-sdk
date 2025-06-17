package main

import (
	"fmt"

	catalyst "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"
)

// client is Catalyst Center API client
var client *catalyst.Client

func main() {
	var err error
	fmt.Println("Authenticating...")
	deviceUUID := ""
	client, err = catalyst.NewClientWithOptions("https://192.168.196.2/",
		"altus", "Altus123",
		"true", "false", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Printing ComplianceDetails")
	getComplianceDetailQueryParams := &catalyst.GetComplianceDetailQueryParams{}
	respComplianceDetail, _, err := client.Compliance.GetComplianceDetail(getComplianceDetailQueryParams)
	if err != nil {
		fmt.Println(err)
		return
	}

	if respComplianceDetail != nil {
		fmt.Println(respComplianceDetail)
		deviceUUID = (*respComplianceDetail.Response)[0].DeviceUUID
	} else {
		fmt.Println("There is no data on response")
		return
	}

	fmt.Println("Post ConfArchive")
	reqBody := &catalyst.RequestConfigurationArchiveExportDeviceConfigurations{
		DeviceID: []string{deviceUUID},
		Password: "C1sco123!",
	}

	resp, _, err := client.ConfigurationArchive.ExportDeviceConfigurations(reqBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	if resp != nil {
		fmt.Println(resp)
	} else {
		fmt.Println("There is no data on response")
		return
	}

}
