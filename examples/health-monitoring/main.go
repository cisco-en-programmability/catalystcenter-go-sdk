package main

import (
	"fmt"

	catalyst "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"
)

// client is Catalyst Center API client
var client *catalyst.Client

func main() {
	fmt.Println("Authenticating...")
	var err error
	client, err = catalyst.NewClientWithOptions("https://sandboxdnac.cisco.com",
		"devnetuser", "Cisco123!",
		"true", "false", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < 10; i++ {
		fmt.Println("Printing site health...")
		// now := time.Now() // current local time
		// sec := now.UnixNano()
		getSiteHealthQueryParams := &catalyst.GetSiteHealthV1QueryParams{
			// Timestamp: strconv.Itoa(int(sec) / 1000000),
		}
		siteHealth, _, err := client.Sites.GetSiteHealth(getSiteHealthQueryParams)
		if err != nil {
			fmt.Println(err)
		}
		print(siteHealth)
		// if siteHealth.Response != nil {
		// 	for id, site := range *siteHealth.Response {
		// 		fmt.Println(fmt.Sprintf("Site --> ID: %d, Name: %s, Health: %d", id, site.SiteName, site.NetworkHealthAverage))
		// 	}
		// }
	}

	// getOverallNetworkHealthQueryParams := &catalyst.GetOverallNetworkHealthQueryParams{
	// 	Timestamp: strconv.Itoa(int(sec) / 1000000),
	// }

	// networkHealth, _, err := client.Topology.GetOverallNetworkHealth(getOverallNetworkHealthQueryParams)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// if networkHealth.Response != nil {
	// 	for _, network := range *networkHealth.Response {
	// 		fmt.Println(fmt.Sprintf("Network Health --> Good Count: %d, Bad Count: %d, Health Score: %d", network.GoodCount, network.BadCount, network.HealthScore))
	// 	}
	// }

	// getOverallClientHealthQueryParams := &catalyst.GetOverallClientHealthQueryParams{
	// 	Timestamp: strconv.Itoa(int(sec) / 1000000),
	// }

	// clientHealth, _, err := client.Clients.GetOverallClientHealth(getOverallClientHealthQueryParams)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// if clientHealth.Response != nil {
	// 	for id := range *clientHealth.Response {
	// 		fmt.Println(id, (*clientHealth.Response)[id])
	// 	}
	// 	fmt.Println((*clientHealth.Response)[0])
	// }

}
