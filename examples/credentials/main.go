package main

import (
	"fmt"

	catalyst "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"
)

// client is Catalyst Center API client
var client *catalyst.Client

func main() {
	fmt.Println("Authenticating...")
	var err error
	client, err = catalyst.NewClientWithOptions("https://192.168.196.2/",
		"altus", "Altus123",
		"false", "false", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Creating new SNMPv3 credentials...")
	snmpv3Credentials := &catalyst.RequestDiscoveryCreateSNMPv3Credentials{
		catalyst.RequestItemDiscoveryCreateSNMPv3Credentials{
			AuthType:        "SHA",
			AuthPassword:    "CATALYST-2020",
			SNMPMode:        "AUTHPRIV",
			Username:        "catalyst-guide",
			PrivacyType:     "AES128",
			PrivacyPassword: "CATALYST-PRIV-2020",
		},
	}

	snmpv3CredentialsResponse, _, err := client.Discovery.CreateSNMPv3Credentials(snmpv3Credentials)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(snmpv3CredentialsResponse)

	fmt.Println("Printing SNMPv3 credentials...")
	getGlobalCredentialsQueryParams := &catalyst.GetGlobalCredentialsQueryParams{
		CredentialSubType: "SNMPV3",
	}
	credentialsListResponse, _, err := client.Discovery.GetGlobalCredentials(getGlobalCredentialsQueryParams)
	if err != nil {
		fmt.Println(err)
	}
	if credentialsListResponse.Response != nil {
		for id, credential := range *credentialsListResponse.Response {
			fmt.Println("GET: ", id, credential.ID, credential.Description, credential.CredentialType)
			_, _, err := client.Discovery.DeleteGlobalCredentialsByID(credential.ID)
			if err != nil {
				continue
			}
		}
	}

	fmt.Println("Creating new HTTP Write credentials...")
	port := 443
	secure := true
	httpWriteCredentials := &catalyst.RequestDiscoveryCreateHTTPWriteCredentials{
		catalyst.RequestItemDiscoveryCreateHTTPWriteCredentials{
			Comments:    "Catalyst Center HTTP Credentials",
			Description: "HTTP Creds",
			Password:    "HTTP-cr3d$",
			Port:        &port,
			Username:    "catalyst-http-user",
			Secure:      &secure,
		},
	}

	httpWriteCredentialsResponse, _, err := client.Discovery.CreateHTTPWriteCredentials(httpWriteCredentials)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(httpWriteCredentialsResponse)

	fmt.Println("Printing HTTP Write credentials...")
	getGlobalCredentialsQueryParams = &catalyst.GetGlobalCredentialsQueryParams{
		CredentialSubType: "HTTP_WRITE",
	}
	credentialsListResponse, _, err = client.Discovery.GetGlobalCredentials(getGlobalCredentialsQueryParams)
	if err != nil {
		fmt.Println(err)
	}
	if credentialsListResponse.Response != nil {
		for id, credential := range *credentialsListResponse.Response {
			fmt.Println("GET: ", id, credential.ID, credential.Description, credential.CredentialType)
			_, _, err := client.Discovery.DeleteGlobalCredentialsByID(credential.ID)
			if err != nil {
				continue
			}
		}
	}

}
