package main

import (
	"fmt"

	catalyst "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"
)

// client is Catalyst Center API client
var client *catalyst.Client

func main() {
	var err error
	fmt.Println("Authenticating...")
	client, err = catalyst.NewClientWithOptions("https://192.168.196.2/",
		"altus", "Altus123",
		"true", "false", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Printing name spaces ===> ")
	namespaces, _, err := client.File.GetListOfAvailableNamespaces()
	if err != nil {
		fmt.Println(err)
		return
	}

	var fileId string
	if len(namespaces.Response) <= 0 {
		fmt.Println("There are not avaible namespaces for the test")
		return
	}
	canDownload := false
	fmt.Println("Looking for a valid fileID ===> ")
	if namespaces.Response != nil {
		for _, name := range namespaces.Response {
			list, _, err := client.File.GetListOfFiles(name)
			if err != nil {
				fmt.Println(err)
			}

			if list.Response != nil && len(*list.Response) > 0 {
				fmt.Println("Finded!")
				fileId = (*list.Response)[0].ID
				canDownload = true
				break
			}
		}
	}

	if !canDownload {
		fmt.Println("There are no files for testing")
		return
	}

	fmt.Println("Testing Download ===>")
	fileResponse, _, err := client.File.DownloadAFileByFileID(fileId)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("File ===>")
	fmt.Println(fileResponse)

}
