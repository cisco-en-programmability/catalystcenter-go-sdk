package main

import (
	"fmt"

	catalyst "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"
)

// Client is Catalyst Center API client
var Client *catalyst.Client

func main() {
	fmt.Println("Authenticating...")
	var err error
	Client, err = catalyst.NewClientWithOptions("https://sandboxdnac.cisco.com",
		"devnetuser", "Cisco123!",
		"true", "false", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Printing site topology...")
	topology, _, err := Client.Topology.GetSiteTopology()
	if err != nil {
		fmt.Println(err)
	}
	if topology.Response != nil && topology.Response.Sites != nil {
		for id, site := range *topology.Response.Sites {
			fmt.Println("GET:", id, site.ID, site.GroupNameHierarchy)
		}
	}

	fmt.Println("Printing physical topology...")

	getPhysicalTopologyQueryParams := &catalyst.GetPhysicalTopologyQueryParams{
		NodeType: "",
	}
	physicalTopology, _, err := Client.Topology.GetPhysicalTopology(getPhysicalTopologyQueryParams)
	if err != nil {
		fmt.Println(err)
	}
	if physicalTopology.Response != nil && physicalTopology.Response.Nodes != nil {
		for id, nodes := range *physicalTopology.Response.Nodes {
			fmt.Println("GET:", id, nodes.ID, nodes.IP, nodes.Label, nodes.AdditionalInfo)
		}
	}

	fmt.Println("Printing VLAN Information...")
	vlanInformation, _, err := Client.Topology.GetVLANDetails()
	if err != nil {
		fmt.Println(err)
	}
	for id, name := range vlanInformation.Response {
		fmt.Println("GET:", id, name)
	}

	fmt.Println("Printing VLAN 1 Details ...")
	vlanDetails, _, err := Client.Topology.GetTopologyDetails("1")
	if err != nil {
		fmt.Println(err)
	}
	if vlanDetails.Response != nil && vlanDetails.Response.Links != nil {
		for id, link := range *vlanDetails.Response.Links {
			fmt.Println("GET:", id, link.ID, link.Source, link.Target, link.LinkStatus)
		}
	}

	fmt.Println("Printing Network Health...")
	getOverallNetworkHealthQueryParams := &catalyst.GetOverallNetworkHealthQueryParams{
		Timestamp: 0,
	}
	networkHealth, _, err := Client.Topology.GetOverallNetworkHealth(getOverallNetworkHealthQueryParams)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(networkHealth.HealthDistirubution, networkHealth.LatestHealthScore, networkHealth.Response, networkHealth.MonitoredDevices)
}
