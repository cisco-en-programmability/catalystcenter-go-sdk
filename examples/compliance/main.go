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

	fmt.Println("Printing ComplianceDetailCount")
	getComplianceDetailCountQueryParams := &catalyst.GetComplianceDetailCountQueryParams{}

	respComplianceDetailCount, _, err := client.Compliance.GetComplianceDetailCount(getComplianceDetailCountQueryParams)
	if err != nil {
		fmt.Println(err)
		return
	}

	if respComplianceDetailCount != nil {
		fmt.Println(respComplianceDetailCount)
	} else {
		fmt.Println("There is no data on response")
		return
	}

	fmt.Println("Printing ComplianceStatus")

	getGetComplianceStatusQueryParams := &catalyst.GetComplianceStatusQueryParams{}
	respComplianceStatus, _, err := client.Compliance.GetComplianceStatus(getGetComplianceStatusQueryParams)
	if err != nil {
		fmt.Println(err)
		return
	}

	if respComplianceStatus != nil {
		fmt.Println(respComplianceStatus)
	} else {
		fmt.Println("There is no data on response")
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

	fmt.Println("Printing ComplianceDetailsCount")
	getComplianceStatusCountQueryParams := &catalyst.GetComplianceStatusCountQueryParams{}
	respComplianceStatusCount, _, err := client.Compliance.GetComplianceStatusCount(getComplianceStatusCountQueryParams)
	if err != nil {
		fmt.Println(err)
		return
	}

	if respComplianceStatusCount != nil {
		fmt.Println(respComplianceStatusCount)
	} else {
		fmt.Println("There is no data on response")
		return
	}

	fmt.Println("Printing ComplianceDetailsOfDevice")
	getComplianceDetailsOfDeviceQueryParams := &catalyst.ComplianceDetailsOfDeviceQueryParams{}
	respComplianceDetailsOfDevice, _, err := client.Compliance.ComplianceDetailsOfDevice(deviceUUID, getComplianceDetailsOfDeviceQueryParams)
	if err != nil {
		fmt.Println(err)
		return
	}

	if respComplianceDetailsOfDevice != nil {
		fmt.Println(respComplianceDetailsOfDevice)
	} else {
		fmt.Println("There is no data on response")
		return
	}

	fmt.Println("POST RunCompiliance")
	triggerFull := true
	reqBody := &catalyst.RequestComplianceRunCompliance{
		TriggerFull: &triggerFull,
		Categories:  []string{"PSIRT"},
		DeviceUUIDs: []string{deviceUUID},
	}
	resp, _, err := client.Compliance.RunCompliance(reqBody)
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
