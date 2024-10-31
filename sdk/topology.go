package catalyst

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type TopologyService service

type GetOverallNetworkHealthV1QueryParams struct {
	Timestamp float64 `url:"timestamp,omitempty"` //UTC timestamp of network health data in milliseconds
}
type GetPhysicalTopologyV1QueryParams struct {
	NodeType string `url:"nodeType,omitempty"` //nodeType
}

type ResponseTopologyGetOverallNetworkHealthV1 struct {
	Version                    string                                                          `json:"version,omitempty"`                    // This output's version string
	Response                   *[]ResponseTopologyGetOverallNetworkHealthV1Response            `json:"response,omitempty"`                   //
	MeasuredBy                 string                                                          `json:"measuredBy,omitempty"`                 // Overall network health measure by 'global'
	LatestMeasuredByEntity     string                                                          `json:"latestMeasuredByEntity,omitempty"`     // Latest measured by entity
	LatestHealthScore          *int                                                            `json:"latestHealthScore,omitempty"`          // Latest health score value
	MonitoredDevices           *int                                                            `json:"monitoredDevices,omitempty"`           // Number of monitored devices
	MonitoredHealthyDevices    *int                                                            `json:"monitoredHealthyDevices,omitempty"`    // Number of healthy devices
	MonitoredUnHealthyDevices  *int                                                            `json:"monitoredUnHealthyDevices,omitempty"`  // Number of unhealthy devices
	UnMonitoredDevices         *int                                                            `json:"unMonitoredDevices,omitempty"`         // Number of un-monitored devices
	NoHealthDevices            *int                                                            `json:"noHealthDevices,omitempty"`            // Number of un-monitored devices
	TotalDevices               *int                                                            `json:"totalDevices,omitempty"`               // Total number of devices
	MonitoredPoorHealthDevices *int                                                            `json:"monitoredPoorHealthDevices,omitempty"` // Number of poor health devices
	MonitoredFairHealthDevices *int                                                            `json:"monitoredFairHealthDevices,omitempty"` // Number of fair health devices
	HealthContributingDevices  *int                                                            `json:"healthContributingDevices,omitempty"`  // Number of health contributing devices
	HealthDistirubution        *[]ResponseTopologyGetOverallNetworkHealthV1HealthDistirubution `json:"healthDistirubution,omitempty"`        //
}
type ResponseTopologyGetOverallNetworkHealthV1Response struct {
	Time                 string `json:"time,omitempty"`                 // Date-time string
	HealthScore          *int   `json:"healthScore,omitempty"`          // Health score
	TotalCount           *int   `json:"totalCount,omitempty"`           // Total health count
	GoodCount            *int   `json:"goodCount,omitempty"`            // Total good health count
	NoHealthCount        *int   `json:"noHealthCount,omitempty"`        // Total no health count
	UnmonCount           *int   `json:"unmonCount,omitempty"`           // Total no health count
	FairCount            *int   `json:"fairCount,omitempty"`            // Total fair health count
	BadCount             *int   `json:"badCount,omitempty"`             // Total bad health count
	MaintenanceModeCount *int   `json:"maintenanceModeCount,omitempty"` // Total maintenance mode count
	Entity               string `json:"entity,omitempty"`               // Entity of the health data
	TimeinMillis         *int   `json:"timeinMillis,omitempty"`         // UTC time value of property 'time' in milliseconds
}
type ResponseTopologyGetOverallNetworkHealthV1HealthDistirubution struct {
	Category              string                                                                    `json:"category,omitempty"`              // Device category in this health data
	TotalCount            *int                                                                      `json:"totalCount,omitempty"`            // Total device count
	HealthScore           *int                                                                      `json:"healthScore,omitempty"`           // Health score
	GoodPercentage        *float64                                                                  `json:"goodPercentage,omitempty"`        // Good health percent
	BadPercentage         *float64                                                                  `json:"badPercentage,omitempty"`         // Poor health percent
	FairPercentage        *float64                                                                  `json:"fairPercentage,omitempty"`        // Fair health percent
	NoHealthPercentage    *float64                                                                  `json:"noHealthPercentage,omitempty"`    // No health percent
	UnmonPercentage       *float64                                                                  `json:"unmonPercentage,omitempty"`       // No health percent
	GoodCount             *float64                                                                  `json:"goodCount,omitempty"`             // Good health count
	BadCount              *float64                                                                  `json:"badCount,omitempty"`              // Poor health count
	FairCount             *float64                                                                  `json:"fairCount,omitempty"`             // Fair health count
	NoHealthCount         *float64                                                                  `json:"noHealthCount,omitempty"`         // No health count
	UnmonCount            *float64                                                                  `json:"unmonCount,omitempty"`            // No health count
	ThirdPartyDeviceCount *float64                                                                  `json:"thirdPartyDeviceCount,omitempty"` // Third party device count
	KpiMetrics            *[]ResponseTopologyGetOverallNetworkHealthV1HealthDistirubutionKpiMetrics `json:"kpiMetrics,omitempty"`            //
}
type ResponseTopologyGetOverallNetworkHealthV1HealthDistirubutionKpiMetrics struct {
	Key   string `json:"key,omitempty"`   // Health key
	Value string `json:"value,omitempty"` // Health value
}
type ResponseTopologyGetTopologyDetailsV1 struct {
	Response *ResponseTopologyGetTopologyDetailsV1Response `json:"response,omitempty"` //
	Version  string                                        `json:"version,omitempty"`  //
}
type ResponseTopologyGetTopologyDetailsV1Response struct {
	ID    string                                               `json:"id,omitempty"`    // [Deprecated]
	Links *[]ResponseTopologyGetTopologyDetailsV1ResponseLinks `json:"links,omitempty"` //
	Nodes *[]ResponseTopologyGetTopologyDetailsV1ResponseNodes `json:"nodes,omitempty"` //
}
type ResponseTopologyGetTopologyDetailsV1ResponseLinks struct {
	AdditionalInfo       *ResponseTopologyGetTopologyDetailsV1ResponseLinksAdditionalInfo `json:"additionalInfo,omitempty"`       // Additional information about the link
	EndPortID            string                                                           `json:"endPortID,omitempty"`            // Device port ID corresponding to the end device
	EndPortIPv4Address   string                                                           `json:"endPortIpv4Address,omitempty"`   // Interface port IPv4 address corresponding to the end device
	EndPortIPv4Mask      string                                                           `json:"endPortIpv4Mask,omitempty"`      // Interface port IPv4 mask corresponding to the end device
	EndPortName          string                                                           `json:"endPortName,omitempty"`          // Interface port name corresponding to the end device
	EndPortSpeed         string                                                           `json:"endPortSpeed,omitempty"`         // Interface port speed corresponding to end device
	GreyOut              *bool                                                            `json:"greyOut,omitempty"`              // Indicates if the link is greyed out
	ID                   string                                                           `json:"id,omitempty"`                   // Id of the link
	LinkStatus           string                                                           `json:"linkStatus,omitempty"`           // Indicates whether link is up or down
	Source               string                                                           `json:"source,omitempty"`               // Device ID corresponding to the source device
	StartPortID          string                                                           `json:"startPortID,omitempty"`          // Device port ID corresponding to start device
	StartPortIPv4Address string                                                           `json:"startPortIpv4Address,omitempty"` // Interface port IPv4 address corresponding to start device
	StartPortIPv4Mask    string                                                           `json:"startPortIpv4Mask,omitempty"`    // Interface port IPv4 mask corresponding to start device
	StartPortName        string                                                           `json:"startPortName,omitempty"`        // Interface port name corresponding to start device
	StartPortSpeed       string                                                           `json:"startPortSpeed,omitempty"`       // Interface port speed corresponding to start device
	Tag                  string                                                           `json:"tag,omitempty"`                  // [Deprecated]
	Target               string                                                           `json:"target,omitempty"`               // Device ID corresponding to the target device
}
type ResponseTopologyGetTopologyDetailsV1ResponseLinksAdditionalInfo interface{}
type ResponseTopologyGetTopologyDetailsV1ResponseNodes struct {
	ACLApplied        *bool                                                            `json:"aclApplied,omitempty"`        // Indicates if the Access Control List (ACL) is applied on the device
	AdditionalInfo    *ResponseTopologyGetTopologyDetailsV1ResponseNodesAdditionalInfo `json:"additionalInfo,omitempty"`    // Additional information about the node
	CustomParam       *ResponseTopologyGetTopologyDetailsV1ResponseNodesCustomParam    `json:"customParam,omitempty"`       //
	ConnectedDeviceID string                                                           `json:"connectedDeviceId,omitempty"` // ID of the connected device when the nodeType is HOST
	DataPathID        string                                                           `json:"dataPathId,omitempty"`        // ID of the path between devices
	DeviceType        string                                                           `json:"deviceType,omitempty"`        // Type of the device.
	DeviceSeries      string                                                           `json:"deviceSeries,omitempty"`      // The series of the device
	Family            string                                                           `json:"family,omitempty"`            // The product family of the device
	Fixed             *bool                                                            `json:"fixed,omitempty"`             // Boolean value indicating whether the position is fixed or will use auto layout
	GreyOut           *bool                                                            `json:"greyOut,omitempty"`           // Boolean value indicating whether the node is active for the topology view.
	ID                string                                                           `json:"id,omitempty"`                // Unique identifier for the device
	IP                string                                                           `json:"ip,omitempty"`                // IP address of the device
	Label             string                                                           `json:"label,omitempty"`             // Label of the node, typically the hostname of the device
	NetworkType       string                                                           `json:"networkType,omitempty"`       // Type of the network
	NodeType          string                                                           `json:"nodeType,omitempty"`          // Type of the node can be 'device' or 'HOST'
	Order             *int                                                             `json:"order,omitempty"`             // Device order by link number
	OsType            string                                                           `json:"osType,omitempty"`            // OS type of the device
	PlatformID        string                                                           `json:"platformId,omitempty"`        // Platform description of the device
	Role              string                                                           `json:"role,omitempty"`              // Role of the device
	RoleSource        string                                                           `json:"roleSource,omitempty"`        // Indicates whether the role is assigned manually or automatically
	SoftwareVersion   string                                                           `json:"softwareVersion,omitempty"`   // Device OS version
	Tags              []string                                                         `json:"tags,omitempty"`              // [Deprecated]
	UpperNode         string                                                           `json:"upperNode,omitempty"`         // ID of the start node
	UserID            string                                                           `json:"userId,omitempty"`            // ID of the host
	VLANID            string                                                           `json:"vlanId,omitempty"`            // VLAN ID
	X                 *int                                                             `json:"x,omitempty"`                 // [Deprecated] Please refer to customParam.x
	Y                 *int                                                             `json:"y,omitempty"`                 // [Deprecated] Please refer to customerParam.y
}
type ResponseTopologyGetTopologyDetailsV1ResponseNodesAdditionalInfo interface{}
type ResponseTopologyGetTopologyDetailsV1ResponseNodesCustomParam struct {
	ID           string `json:"id,omitempty"`           // [Deprecated] Please refer to nodes.id
	Label        string `json:"label,omitempty"`        // Label of the node
	ParentNodeID string `json:"parentNodeId,omitempty"` // Id of the parent node
	X            *int   `json:"x,omitempty"`            // X coordinate for this node in the topology view
	Y            *int   `json:"y,omitempty"`            // Y coordinate for this node in the topology view
}
type ResponseTopologyGetL3TopologyDetailsV1 struct {
	Response *ResponseTopologyGetL3TopologyDetailsV1Response `json:"response,omitempty"` //
	Version  string                                          `json:"version,omitempty"`  //
}
type ResponseTopologyGetL3TopologyDetailsV1Response struct {
	ID    string                                                 `json:"id,omitempty"`    // [Deprecated]
	Links *[]ResponseTopologyGetL3TopologyDetailsV1ResponseLinks `json:"links,omitempty"` //
	Nodes *[]ResponseTopologyGetL3TopologyDetailsV1ResponseNodes `json:"nodes,omitempty"` //
}
type ResponseTopologyGetL3TopologyDetailsV1ResponseLinks struct {
	AdditionalInfo       *ResponseTopologyGetL3TopologyDetailsV1ResponseLinksAdditionalInfo `json:"additionalInfo,omitempty"`       // Additional information about the link
	EndPortID            string                                                             `json:"endPortID,omitempty"`            // Device port ID corresponding to the end device
	EndPortIPv4Address   string                                                             `json:"endPortIpv4Address,omitempty"`   // Interface port IPv4 address corresponding to the end device
	EndPortIPv4Mask      string                                                             `json:"endPortIpv4Mask,omitempty"`      // Interface port IPv4 mask corresponding to the end device
	EndPortName          string                                                             `json:"endPortName,omitempty"`          // Interface port name corresponding to the end device
	EndPortSpeed         string                                                             `json:"endPortSpeed,omitempty"`         // Interface port speed corresponding to end device
	GreyOut              *bool                                                              `json:"greyOut,omitempty"`              // Indicates if the link is greyed out
	ID                   string                                                             `json:"id,omitempty"`                   // Id of the link
	LinkStatus           string                                                             `json:"linkStatus,omitempty"`           // Indicates whether link is up or down
	Source               string                                                             `json:"source,omitempty"`               // Device ID corresponding to the source device
	StartPortID          string                                                             `json:"startPortID,omitempty"`          // Device port ID corresponding to start device
	StartPortIPv4Address string                                                             `json:"startPortIpv4Address,omitempty"` // Interface port IPv4 address corresponding to start device
	StartPortIPv4Mask    string                                                             `json:"startPortIpv4Mask,omitempty"`    // Interface port IPv4 mask corresponding to start device
	StartPortName        string                                                             `json:"startPortName,omitempty"`        // Interface port name corresponding to start device
	StartPortSpeed       string                                                             `json:"startPortSpeed,omitempty"`       // Interface port speed corresponding to start device
	Tag                  string                                                             `json:"tag,omitempty"`                  // [Deprecated]
	Target               string                                                             `json:"target,omitempty"`               // Device ID corresponding to the target device
}
type ResponseTopologyGetL3TopologyDetailsV1ResponseLinksAdditionalInfo interface{}
type ResponseTopologyGetL3TopologyDetailsV1ResponseNodes struct {
	ACLApplied        *bool                                                              `json:"aclApplied,omitempty"`        // Indicates if the Access Control List (ACL) is applied on the device
	AdditionalInfo    *ResponseTopologyGetL3TopologyDetailsV1ResponseNodesAdditionalInfo `json:"additionalInfo,omitempty"`    // Additional information about the node
	CustomParam       *ResponseTopologyGetL3TopologyDetailsV1ResponseNodesCustomParam    `json:"customParam,omitempty"`       //
	ConnectedDeviceID string                                                             `json:"connectedDeviceId,omitempty"` // ID of the connected device when the nodeType is HOST
	DataPathID        string                                                             `json:"dataPathId,omitempty"`        // ID of the path between devices
	DeviceType        string                                                             `json:"deviceType,omitempty"`        // Type of the device.
	DeviceSeries      string                                                             `json:"deviceSeries,omitempty"`      // The series of the device
	Family            string                                                             `json:"family,omitempty"`            // The product family of the device
	Fixed             *bool                                                              `json:"fixed,omitempty"`             // Boolean value indicating whether the position is fixed or will use auto layout
	GreyOut           *bool                                                              `json:"greyOut,omitempty"`           // Boolean value indicating whether the node is active for the topology view.
	ID                string                                                             `json:"id,omitempty"`                // Unique identifier for the device
	IP                string                                                             `json:"ip,omitempty"`                // IP address of the device
	Label             string                                                             `json:"label,omitempty"`             // Label of the node, typically the hostname of the device
	NetworkType       string                                                             `json:"networkType,omitempty"`       // Type of the network
	NodeType          string                                                             `json:"nodeType,omitempty"`          // Type of the node can be 'device' or 'HOST'
	Order             *int                                                               `json:"order,omitempty"`             // Device order by link number
	OsType            string                                                             `json:"osType,omitempty"`            // OS type of the device
	PlatformID        string                                                             `json:"platformId,omitempty"`        // Platform description of the device
	Role              string                                                             `json:"role,omitempty"`              // Role of the device
	RoleSource        string                                                             `json:"roleSource,omitempty"`        // Indicates whether the role is assigned manually or automatically
	SoftwareVersion   string                                                             `json:"softwareVersion,omitempty"`   // Device OS version
	Tags              []string                                                           `json:"tags,omitempty"`              // [Deprecated]
	UpperNode         string                                                             `json:"upperNode,omitempty"`         // ID of the start node
	UserID            string                                                             `json:"userId,omitempty"`            // ID of the host
	VLANID            string                                                             `json:"vlanId,omitempty"`            // VLAN ID
	X                 *int                                                               `json:"x,omitempty"`                 // [Deprecated] Please refer to customParam.x
	Y                 *int                                                               `json:"y,omitempty"`                 // [Deprecated] Please refer to customerParam.y
}
type ResponseTopologyGetL3TopologyDetailsV1ResponseNodesAdditionalInfo interface{}
type ResponseTopologyGetL3TopologyDetailsV1ResponseNodesCustomParam struct {
	ID           string `json:"id,omitempty"`           // [Deprecated] Please refer to nodes.id
	Label        string `json:"label,omitempty"`        // Label of the node
	ParentNodeID string `json:"parentNodeId,omitempty"` // Id of the parent node
	X            *int   `json:"x,omitempty"`            // X coordinate for this node in the topology view
	Y            *int   `json:"y,omitempty"`            // Y coordinate for this node in the topology view
}
type ResponseTopologyGetPhysicalTopologyV1 struct {
	Response *ResponseTopologyGetPhysicalTopologyV1Response `json:"response,omitempty"` //
	Version  string                                         `json:"version,omitempty"`  //
}
type ResponseTopologyGetPhysicalTopologyV1Response struct {
	ID    string                                                `json:"id,omitempty"`    // [Deprecated]
	Links *[]ResponseTopologyGetPhysicalTopologyV1ResponseLinks `json:"links,omitempty"` //
	Nodes *[]ResponseTopologyGetPhysicalTopologyV1ResponseNodes `json:"nodes,omitempty"` //
}
type ResponseTopologyGetPhysicalTopologyV1ResponseLinks struct {
	AdditionalInfo       *ResponseTopologyGetPhysicalTopologyV1ResponseLinksAdditionalInfo `json:"additionalInfo,omitempty"`       // Additional information about the link
	EndPortID            string                                                            `json:"endPortID,omitempty"`            // Device port ID corresponding to the end device
	EndPortIPv4Address   string                                                            `json:"endPortIpv4Address,omitempty"`   // Interface port IPv4 address corresponding to the end device
	EndPortIPv4Mask      string                                                            `json:"endPortIpv4Mask,omitempty"`      // Interface port IPv4 mask corresponding to the end device
	EndPortName          string                                                            `json:"endPortName,omitempty"`          // Interface port name corresponding to the end device
	EndPortSpeed         string                                                            `json:"endPortSpeed,omitempty"`         // Interface port speed corresponding to end device
	GreyOut              *bool                                                             `json:"greyOut,omitempty"`              // Indicates if the link is greyed out
	ID                   string                                                            `json:"id,omitempty"`                   // Id of the link
	LinkStatus           string                                                            `json:"linkStatus,omitempty"`           // Indicates whether link is up or down
	Source               string                                                            `json:"source,omitempty"`               // Device ID corresponding to the source device
	StartPortID          string                                                            `json:"startPortID,omitempty"`          // Device port ID corresponding to start device
	StartPortIPv4Address string                                                            `json:"startPortIpv4Address,omitempty"` // Interface port IPv4 address corresponding to start device
	StartPortIPv4Mask    string                                                            `json:"startPortIpv4Mask,omitempty"`    // Interface port IPv4 mask corresponding to start device
	StartPortName        string                                                            `json:"startPortName,omitempty"`        // Interface port name corresponding to start device
	StartPortSpeed       string                                                            `json:"startPortSpeed,omitempty"`       // Interface port speed corresponding to start device
	Tag                  string                                                            `json:"tag,omitempty"`                  // [Deprecated]
	Target               string                                                            `json:"target,omitempty"`               // Device ID corresponding to the target device
}
type ResponseTopologyGetPhysicalTopologyV1ResponseLinksAdditionalInfo interface{}
type ResponseTopologyGetPhysicalTopologyV1ResponseNodes struct {
	ACLApplied        *bool                                                             `json:"aclApplied,omitempty"`        // Indicates if the Access Control List (ACL) is applied on the device
	AdditionalInfo    *ResponseTopologyGetPhysicalTopologyV1ResponseNodesAdditionalInfo `json:"additionalInfo,omitempty"`    // Additional information about the node
	CustomParam       *ResponseTopologyGetPhysicalTopologyV1ResponseNodesCustomParam    `json:"customParam,omitempty"`       //
	ConnectedDeviceID string                                                            `json:"connectedDeviceId,omitempty"` // ID of the connected device when the nodeType is HOST
	DataPathID        string                                                            `json:"dataPathId,omitempty"`        // ID of the path between devices
	DeviceType        string                                                            `json:"deviceType,omitempty"`        // Type of the device.
	DeviceSeries      string                                                            `json:"deviceSeries,omitempty"`      // The series of the device
	Family            string                                                            `json:"family,omitempty"`            // The product family of the device
	Fixed             *bool                                                             `json:"fixed,omitempty"`             // Boolean value indicating whether the position is fixed or will use auto layout
	GreyOut           *bool                                                             `json:"greyOut,omitempty"`           // Boolean value indicating whether the node is active for the topology view.
	ID                string                                                            `json:"id,omitempty"`                // Unique identifier for the device
	IP                string                                                            `json:"ip,omitempty"`                // IP address of the device
	Label             string                                                            `json:"label,omitempty"`             // Label of the node, typically the hostname of the device
	NetworkType       string                                                            `json:"networkType,omitempty"`       // Type of the network
	NodeType          string                                                            `json:"nodeType,omitempty"`          // Type of the node can be 'device' or 'HOST'
	Order             *int                                                              `json:"order,omitempty"`             // Device order by link number
	OsType            string                                                            `json:"osType,omitempty"`            // OS type of the device
	PlatformID        string                                                            `json:"platformId,omitempty"`        // Platform description of the device
	Role              string                                                            `json:"role,omitempty"`              // Role of the device
	RoleSource        string                                                            `json:"roleSource,omitempty"`        // Indicates whether the role is assigned manually or automatically
	SoftwareVersion   string                                                            `json:"softwareVersion,omitempty"`   // Device OS version
	Tags              []string                                                          `json:"tags,omitempty"`              // [Deprecated]
	UpperNode         string                                                            `json:"upperNode,omitempty"`         // ID of the start node
	UserID            string                                                            `json:"userId,omitempty"`            // ID of the host
	VLANID            string                                                            `json:"vlanId,omitempty"`            // VLAN ID
	X                 *int                                                              `json:"x,omitempty"`                 // [Deprecated] Please refer to customParam.x
	Y                 *int                                                              `json:"y,omitempty"`                 // [Deprecated] Please refer to customerParam.y
}
type ResponseTopologyGetPhysicalTopologyV1ResponseNodesAdditionalInfo interface{}
type ResponseTopologyGetPhysicalTopologyV1ResponseNodesCustomParam struct {
	ID           string `json:"id,omitempty"`           // [Deprecated] Please refer to nodes.id
	Label        string `json:"label,omitempty"`        // Label of the node
	ParentNodeID string `json:"parentNodeId,omitempty"` // Id of the parent node
	X            *int   `json:"x,omitempty"`            // X coordinate for this node in the topology view
	Y            *int   `json:"y,omitempty"`            // Y coordinate for this node in the topology view
}
type ResponseTopologyGetSiteTopologyV1 struct {
	Response *ResponseTopologyGetSiteTopologyV1Response `json:"response,omitempty"` //
	Version  string                                     `json:"version,omitempty"`  //
}
type ResponseTopologyGetSiteTopologyV1Response struct {
	Sites *[]ResponseTopologyGetSiteTopologyV1ResponseSites `json:"sites,omitempty"` //
}
type ResponseTopologyGetSiteTopologyV1ResponseSites struct {
	DisplayName        string `json:"displayName,omitempty"`        // Group id of the site
	GroupNameHierarchy string `json:"groupNameHierarchy,omitempty"` // Hierarchy of the site names from the root site to the current site. Each site name is separated by a '/'. Eg. 'Global/Site1/Building1/Floor1'
	ID                 string `json:"id,omitempty"`                 // Unique identifier of the site
	Latitude           string `json:"latitude,omitempty"`           // Latitude of the site
	LocationAddress    string `json:"locationAddress,omitempty"`    // Address of the site
	LocationCountry    string `json:"locationCountry,omitempty"`    // Country corresponding to the address of the site
	LocationType       string `json:"locationType,omitempty"`       // Type of site, eg. 'building', 'area' or 'floor'
	Longitude          string `json:"longitude,omitempty"`          // Longitude of the site
	Name               string `json:"name,omitempty"`               // Name of the site
	ParentID           string `json:"parentId,omitempty"`           // Unique identifier of the parent site
}
type ResponseTopologyGetVLANDetailsV1 struct {
	Response []string `json:"response,omitempty"` // Lists of all available VLAN names
	Version  string   `json:"version,omitempty"`  //
}

//GetOverallNetworkHealthV1 Get Overall Network Health - 7997-6a34-4409-bfbb
/* Returns Overall Network Health information by Device category (Access, Distribution, Core, Router, Wireless) for any given point of time


@param GetOverallNetworkHealthV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-overall-network-health-v1
*/
func (s *TopologyService) GetOverallNetworkHealthV1(GetOverallNetworkHealthV1QueryParams *GetOverallNetworkHealthV1QueryParams) (*ResponseTopologyGetOverallNetworkHealthV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-health"

	queryString, _ := query.Values(GetOverallNetworkHealthV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseTopologyGetOverallNetworkHealthV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetOverallNetworkHealthV1(GetOverallNetworkHealthV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetOverallNetworkHealthV1")
	}

	result := response.Result().(*ResponseTopologyGetOverallNetworkHealthV1)
	return result, response, err

}

//GetTopologyDetailsV1 Get topology details - b9b4-8ac8-463a-8aba
/* Returns Layer 2 network topology by specified VLAN ID


@param vlanID vlanID path parameter. Vlan Name for e.g Vlan1, Vlan23 etc


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-topology-details-v1
*/
func (s *TopologyService) GetTopologyDetailsV1(vlanID string) (*ResponseTopologyGetTopologyDetailsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/topology/l2/{vlanID}"
	path = strings.Replace(path, "{vlanID}", fmt.Sprintf("%v", vlanID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTopologyGetTopologyDetailsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTopologyDetailsV1(vlanID)
		}
		return nil, response, fmt.Errorf("error with operation GetTopologyDetailsV1")
	}

	result := response.Result().(*ResponseTopologyGetTopologyDetailsV1)
	return result, response, err

}

//GetL3TopologyDetailsV1 Get L3 Topology Details - c2b5-fb76-4d88-8375
/* Returns the Layer 3 network topology by routing protocol


@param topologyType topologyType path parameter. Type of topology(OSPF,ISIS,etc)


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-l3-topology-details-v1
*/
func (s *TopologyService) GetL3TopologyDetailsV1(topologyType string) (*ResponseTopologyGetL3TopologyDetailsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/topology/l3/{topologyType}"
	path = strings.Replace(path, "{topologyType}", fmt.Sprintf("%v", topologyType), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTopologyGetL3TopologyDetailsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetL3TopologyDetailsV1(topologyType)
		}
		return nil, response, fmt.Errorf("error with operation GetL3TopologyDetailsV1")
	}

	result := response.Result().(*ResponseTopologyGetL3TopologyDetailsV1)
	return result, response, err

}

//GetPhysicalTopologyV1 Get Physical Topology - b2b8-cb91-459a-a58f
/* Returns the raw physical topology by specified criteria of nodeType


@param GetPhysicalTopologyV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-physical-topology-v1
*/
func (s *TopologyService) GetPhysicalTopologyV1(GetPhysicalTopologyV1QueryParams *GetPhysicalTopologyV1QueryParams) (*ResponseTopologyGetPhysicalTopologyV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/topology/physical-topology"

	queryString, _ := query.Values(GetPhysicalTopologyV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseTopologyGetPhysicalTopologyV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPhysicalTopologyV1(GetPhysicalTopologyV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetPhysicalTopologyV1")
	}

	result := response.Result().(*ResponseTopologyGetPhysicalTopologyV1)
	return result, response, err

}

//GetSiteTopologyV1 Get Site Topology - 9ba1-4a9e-441b-8a60
/* Returns site topology



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-site-topology-v1
*/
func (s *TopologyService) GetSiteTopologyV1() (*ResponseTopologyGetSiteTopologyV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/topology/site-topology"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTopologyGetSiteTopologyV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSiteTopologyV1()
		}
		return nil, response, fmt.Errorf("error with operation GetSiteTopologyV1")
	}

	result := response.Result().(*ResponseTopologyGetSiteTopologyV1)
	return result, response, err

}

//GetVLANDetailsV1 Get VLAN details - 6284-db46-49aa-8d31
/* Returns the list of VLAN names that are involved in a loop as identified by the Spanning Tree Protocol



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-vlan-details-v1
*/
func (s *TopologyService) GetVLANDetailsV1() (*ResponseTopologyGetVLANDetailsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/topology/vlan/vlan-names"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTopologyGetVLANDetailsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetVLANDetailsV1()
		}
		return nil, response, fmt.Errorf("error with operation GetVlanDetailsV1")
	}

	result := response.Result().(*ResponseTopologyGetVLANDetailsV1)
	return result, response, err

}

// Alias Function
func (s *TopologyService) GetSiteTopology() (*ResponseTopologyGetSiteTopologyV1, *resty.Response, error) {
	return s.GetSiteTopologyV1()
}

// Alias Function
func (s *TopologyService) GetPhysicalTopology(GetPhysicalTopologyV1QueryParams *GetPhysicalTopologyV1QueryParams) (*ResponseTopologyGetPhysicalTopologyV1, *resty.Response, error) {
	return s.GetPhysicalTopologyV1(GetPhysicalTopologyV1QueryParams)
}

// Alias Function
func (s *TopologyService) GetTopologyDetails(vlanID string) (*ResponseTopologyGetTopologyDetailsV1, *resty.Response, error) {
	return s.GetTopologyDetailsV1(vlanID)
}

// Alias Function
func (s *TopologyService) GetL3TopologyDetails(topologyType string) (*ResponseTopologyGetL3TopologyDetailsV1, *resty.Response, error) {
	return s.GetL3TopologyDetailsV1(topologyType)
}

// Alias Function
func (s *TopologyService) GetVLANDetails() (*ResponseTopologyGetVLANDetailsV1, *resty.Response, error) {
	return s.GetVLANDetailsV1()
}

// Alias Function
func (s *TopologyService) GetOverallNetworkHealth(GetOverallNetworkHealthV1QueryParams *GetOverallNetworkHealthV1QueryParams) (*ResponseTopologyGetOverallNetworkHealthV1, *resty.Response, error) {
	return s.GetOverallNetworkHealthV1(GetOverallNetworkHealthV1QueryParams)
}
