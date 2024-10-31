package catalyst

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type ClientsService service

type RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1QueryParams struct {
	StartTime                  float64 `url:"startTime,omitempty"`                  //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to current time.
	EndTime                    float64 `url:"endTime,omitempty"`                    //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	Limit                      float64 `url:"limit,omitempty"`                      //Maximum number of records to return
	Offset                     float64 `url:"offset,omitempty"`                     //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	SortBy                     string  `url:"sortBy,omitempty"`                     //A field within the response to sort by.
	Order                      string  `url:"order,omitempty"`                      //The sort order of the field ascending or descending.
	Type                       string  `url:"type,omitempty"`                       //The client device type whether client is connected to network through Wired or Wireless medium.
	OsType                     string  `url:"osType,omitempty"`                     //Client device operating system type. This field supports wildcard (`*`) character-based search. If the value contains the (`*`) character, please use the /query API for regex search.  Ex: `*iOS*` or `iOS*` or `*iOS` Examples: `osType=iOS` (single osType requested) `osType=iOS&osType=Android` (multiple osType requested)
	OsVersion                  string  `url:"osVersion,omitempty"`                  //Client device operating system version This field supports wildcard (`*`) character-based search. If the value contains the (`*`) character, please use the /query API for regex search.  Ex: `*14.3*` or `14.3*` or `*14.3` Examples: `osVersion=14.3` (single osVersion requested) `osVersion=14.3&osVersion=10.1` (multiple osVersion requested)
	SiteHierarchy              string  `url:"siteHierarchy,omitempty"`              //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. "Global/AreaName/BuildingName/FloorName") This field supports wildcard (`*`) character-based search. If the value contains the (`*`) character, please use the /query API for regex search.  Ex: `*BuildingName*` or `BuildingName*` or `*BuildingName` Examples: `siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `siteHierarchy=Global/AreaName/BuildingName1/FloorName1&siteHierarchy=Global/AreaName/BuildingName1/FloorName2` (multiple siteHierarchy requested)
	SiteHierarchyID            string  `url:"siteHierarchyId,omitempty"`            //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. "globalUuid/areaUuid/buildingUuid/floorUuid") This field supports wildcard (`*`) character-based search.  Ex: `*buildingUuid*` or `buildingUuid*` or `*buildingUuid` Examples: `siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid` (single siteHierarchyId requested) `siteHierarchyId=globalUuid/areaUuid/buildingUuid1/floorUuid1&siteHierarchyId=globalUuid/areaUuid/buildingUuid1/floorUuid2` (multiple siteHierarchyId requested)
	SiteID                     string  `url:"siteId,omitempty"`                     //The site UUID without the top level hierarchy. (Ex."floorUuid") Examples: `siteId=floorUuid` (single siteId requested) `siteId=floorUuid1&siteId=floorUuid2` (multiple siteId requested)
	IPv4Address                string  `url:"ipv4Address,omitempty"`                //IPv4 Address of the network entity either network device or client This field supports wildcard (`*`) character-based search.  Ex: `*1.1*` or `1.1*` or `*1.1` Examples: `ipv4Address=1.1.1.1` (single ipv4Address requested) `ipv4Address=1.1.1.1&ipv4Address=2.2.2.2` (multiple ipv4Address requested)
	IPv6Address                string  `url:"ipv6Address,omitempty"`                //IPv6 Address of the network entity either network device or client This field supports wildcard (`*`) character-based search. Ex: `*2001:db8*` or `2001:db8*` or `*2001:db8` Examples: `ipv6Address=2001:db8:0:0:0:0:2:1` (single ipv6Address requested) `ipv6Address=2001:db8:0:0:0:0:2:1&ipv6Address=2001:db8:85a3:8d3:1319:8a2e:370:7348` (multiple ipv6Address requested)
	MacAddress                 string  `url:"macAddress,omitempty"`                 //The macAddress of the network device or client This field supports wildcard (`*`) character-based search.  Ex: `*AB:AB:AB*` or `AB:AB:AB*` or `*AB:AB:AB` Examples: `macAddress=AB:AB:AB:CD:CD:CD` (single macAddress requested) `macAddress=AB:AB:AB:CD:CD:DC&macAddress=AB:AB:AB:CD:CD:FE` (multiple macAddress requested)
	WlcName                    string  `url:"wlcName,omitempty"`                    //Wireless Controller name that reports the wireless client. This field supports wildcard (`*`) character-based search. If the value contains the (`*`) character, please use the /query API for regex search. Ex: `*wlc-25*` or `wlc-25*` or `*wlc-25` Examples: `wlcName=wlc-25` (single wlcName requested) `wlcName=wlc-25&wlc-34` (multiple wlcName requested)
	ConnectedNetworkDeviceName string  `url:"connectedNetworkDeviceName,omitempty"` //Name of the neighbor network device that client is connected to. This field supports wildcard (`*`) character-based search. If the value contains the (`*`) character, please use the /query API for regex search. Ex: `*ap-25*` or `ap-25*` or `*ap-25` Examples: `connectedNetworkDeviceName=ap-25` (single connectedNetworkDeviceName requested) `connectedNetworkDeviceName=ap-25&ap-34` (multiple connectedNetworkDeviceName requested)
	SSID                       string  `url:"ssid,omitempty"`                       //SSID is the name of wireless network to which client connects to. It is also referred to as WLAN ID - Wireless Local Area Network Identifier. This field supports wildcard (`*`) character-based search. If the value contains the (`*`) character, please use the /query API for regex search.  Ex: `*Alpha*` or `Alpha*` or `*Alpha` Examples: `ssid=Alpha` (single ssid requested) `ssid=Alpha&ssid=Guest` (multiple ssid requested)
	Band                       string  `url:"band,omitempty"`                       //WiFi frequency band that client or Access Point operates. Band value is represented in Giga Hertz - GHz Examples: `band=5GHZ` (single band requested) `band=2.4GHZ&band=6GHZ` (multiple band requested)
	View                       string  `url:"view,omitempty"`                       //Client related Views Refer to ClientView schema for list of views supported Examples: `view=Wireless` (single view requested) `view=WirelessHealth&view=WirelessTraffic` (multiple view requested)
	Attribute                  string  `url:"attribute,omitempty"`                  //List of attributes related to resource that can be requested to only be part of the response along with the required attributes. Refer to ClientAttribute schema for list of attributes supported Examples: `attribute=band` (single attribute requested) `attribute=band&attribute=ssid&attribute=overallScore` (multiple attribute requested)
}
type RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheTotalCountOfClientsByApplyingBasicFilteringV1QueryParams struct {
	StartTime                  float64 `url:"startTime,omitempty"`                  //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to current time.
	EndTime                    float64 `url:"endTime,omitempty"`                    //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	Type                       string  `url:"type,omitempty"`                       //The client device type whether client is connected to network through Wired or Wireless medium.
	OsType                     string  `url:"osType,omitempty"`                     //Client device operating system type. This field supports wildcard (`*`) character-based search. If the value contains the (`*`) character, please use the /query API for regex search.  Ex: `*iOS*` or `iOS*` or `*iOS` Examples: `osType=iOS` (single osType requested) `osType=iOS&osType=Android` (multiple osType requested)
	OsVersion                  string  `url:"osVersion,omitempty"`                  //Client device operating system version This field supports wildcard (`*`) character-based search. If the value contains the (`*`) character, please use the /query API for regex search.  Ex: `*14.3*` or `14.3*` or `*14.3` Examples: `osVersion=14.3` (single osVersion requested) `osVersion=14.3&osVersion=10.1` (multiple osVersion requested)
	SiteHierarchy              string  `url:"siteHierarchy,omitempty"`              //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. "Global/AreaName/BuildingName/FloorName") This field supports wildcard (`*`) character-based search. If the value contains the (`*`) character, please use the /query API for regex search.  Ex: `*BuildingName*` or `BuildingName*` or `*BuildingName` Examples: `siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `siteHierarchy=Global/AreaName/BuildingName1/FloorName1&siteHierarchy=Global/AreaName/BuildingName1/FloorName2` (multiple siteHierarchy requested)
	SiteHierarchyID            string  `url:"siteHierarchyId,omitempty"`            //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. "globalUuid/areaUuid/buildingUuid/floorUuid") This field supports wildcard (`*`) character-based search.  Ex: `*buildingUuid*` or `buildingUuid*` or `*buildingUuid` Examples: `siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid` (single siteHierarchyId requested) `siteHierarchyId=globalUuid/areaUuid/buildingUuid1/floorUuid1&siteHierarchyId=globalUuid/areaUuid/buildingUuid1/floorUuid2` (multiple siteHierarchyId requested)
	SiteID                     string  `url:"siteId,omitempty"`                     //The site UUID without the top level hierarchy. (Ex."floorUuid") Examples: `siteId=floorUuid` (single siteId requested) `siteId=floorUuid1&siteId=floorUuid2` (multiple siteId requested)
	IPv4Address                string  `url:"ipv4Address,omitempty"`                //IPv4 Address of the network entity either network device or client This field supports wildcard (`*`) character-based search.  Ex: `*1.1*` or `1.1*` or `*1.1` Examples: `ipv4Address=1.1.1.1` (single ipv4Address requested) `ipv4Address=1.1.1.1&ipv4Address=2.2.2.2` (multiple ipv4Address requested)
	IPv6Address                string  `url:"ipv6Address,omitempty"`                //IPv6 Address of the network entity either network device or client This field supports wildcard (`*`) character-based search. Ex: `*2001:db8*` or `2001:db8*` or `*2001:db8` Examples: `ipv6Address=2001:db8:0:0:0:0:2:1` (single ipv6Address requested) `ipv6Address=2001:db8:0:0:0:0:2:1&ipv6Address=2001:db8:85a3:8d3:1319:8a2e:370:7348` (multiple ipv6Address requested)
	MacAddress                 string  `url:"macAddress,omitempty"`                 //The macAddress of the network device or client This field supports wildcard (`*`) character-based search.  Ex: `*AB:AB:AB*` or `AB:AB:AB*` or `*AB:AB:AB` Examples: `macAddress=AB:AB:AB:CD:CD:CD` (single macAddress requested) `macAddress=AB:AB:AB:CD:CD:DC&macAddress=AB:AB:AB:CD:CD:FE` (multiple macAddress requested)
	WlcName                    string  `url:"wlcName,omitempty"`                    //Wireless Controller name that reports the wireless client. This field supports wildcard (`*`) character-based search. If the value contains the (`*`) character, please use the /query API for regex search. Ex: `*wlc-25*` or `wlc-25*` or `*wlc-25` Examples: `wlcName=wlc-25` (single wlcName requested) `wlcName=wlc-25&wlc-34` (multiple wlcName requested)
	ConnectedNetworkDeviceName string  `url:"connectedNetworkDeviceName,omitempty"` //Name of the neighbor network device that client is connected to. This field supports wildcard (`*`) character-based search. If the value contains the (`*`) character, please use the /query API for regex search. Ex: `*ap-25*` or `ap-25*` or `*ap-25` Examples: `connectedNetworkDeviceName=ap-25` (single connectedNetworkDeviceName requested) `connectedNetworkDeviceName=ap-25&ap-34` (multiple connectedNetworkDeviceName requested)
	SSID                       string  `url:"ssid,omitempty"`                       //SSID is the name of wireless network to which client connects to. It is also referred to as WLAN ID - Wireless Local Area Network Identifier. This field supports wildcard (`*`) character-based search. If the value contains the (`*`) character, please use the /query API for regex search.  Ex: `*Alpha*` or `Alpha*` or `*Alpha` Examples: `ssid=Alpha` (single ssid requested) `ssid=Alpha&ssid=Guest` (multiple ssid requested)
	Band                       string  `url:"band,omitempty"`                       //WiFi frequency band that client or Access Point operates. Band value is represented in Giga Hertz - GHz Examples: `band=5GHZ` (single band requested) `band=2.4GHZ&band=6GHZ` (multiple band requested)
}
type RetrievesTheTotalCountOfClientsByApplyingBasicFilteringV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1HeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheNumberOfClientsByApplyingComplexFiltersV1HeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesSummaryAnalyticsDataRelatedToClientsV1HeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheTopNAnalyticsDataRelatedToClientsV1HeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheTrendAnalyticsDataRelatedToClientsV1HeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesSpecificClientInformationMatchingTheMacaddressV1QueryParams struct {
	StartTime float64 `url:"startTime,omitempty"` //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to current time.
	EndTime   float64 `url:"endTime,omitempty"`   //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	View      string  `url:"view,omitempty"`      //Client related Views Refer to ClientView schema for list of views supported Examples: `view=Wireless` (single view requested) `view=WirelessHealth&view=WirelessTraffic` (multiple view requested)
	Attribute string  `url:"attribute,omitempty"` //List of attributes related to resource that can be requested to only be part of the response along with the required attributes. Refer to ClientAttribute schema for list of attributes supported Examples: `attribute=band` (single attribute requested) `attribute=band&attribute=ssid&attribute=overallScore` (multiple attribute requested)
}
type RetrievesSpecificClientInformationMatchingTheMacaddressV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1HeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetClientDetailV1QueryParams struct {
	MacAddress string  `url:"macAddress,omitempty"` //MAC Address of the client
	Timestamp  float64 `url:"timestamp,omitempty"`  //Epoch time(in milliseconds) when the Client health data is required
}
type GetClientEnrichmentDetailsV1HeaderParams struct {
	EntityType        string `url:"entity_type,omitempty"`         //Expects type string. Client enrichment details can be fetched based on either User ID or Client MAC address. This parameter value must either be network_user_id/mac_address
	EntityValue       string `url:"entity_value,omitempty"`        //Expects type string. Contains the actual value for the entity type that has been defined
	IssueCategory     string `url:"issueCategory,omitempty"`       //Expects type string. The category of the DNA event based on which the underlying issues need to be fetched
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type bool.
}
type GetOverallClientHealthV1QueryParams struct {
	Timestamp float64 `url:"timestamp,omitempty"` //Epoch time(in milliseconds) when the Client health data is required
}
type ClientProximityV1QueryParams struct {
	Username       string  `url:"username,omitempty"`        //Wireless client username for which proximity information is required
	NumberDays     float64 `url:"number_days,omitempty"`     //Number of days to track proximity until current date. Defaults and maximum up to 14 days.
	TimeResolution float64 `url:"time_resolution,omitempty"` //Time interval (in minutes) to measure proximity. Defaults to 15 minutes with a minimum 5 minutes.
}

type ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1 struct {
	Response *[]ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1Response `json:"response,omitempty"` //
	Page     *ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1Page       `json:"page,omitempty"`     //
	Version  string                                                                                                     `json:"version,omitempty"`  // Version
}
type ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1Response struct {
	ID                     string                                                                                                                         `json:"id,omitempty"`                     // Id
	MacAddress             string                                                                                                                         `json:"macAddress,omitempty"`             // Mac Address
	Type                   string                                                                                                                         `json:"type,omitempty"`                   // Type
	Name                   string                                                                                                                         `json:"name,omitempty"`                   // Name
	UserID                 string                                                                                                                         `json:"userId,omitempty"`                 // User Id
	Username               string                                                                                                                         `json:"username,omitempty"`               // Username
	IPv4Address            string                                                                                                                         `json:"ipv4Address,omitempty"`            // Ipv4 Address
	IPv6Addresses          []string                                                                                                                       `json:"ipv6Addresses,omitempty"`          // Ipv6 Addresses
	Vendor                 string                                                                                                                         `json:"vendor,omitempty"`                 // Vendor
	OsType                 string                                                                                                                         `json:"osType,omitempty"`                 // Os Type
	OsVersion              string                                                                                                                         `json:"osVersion,omitempty"`              // Os Version
	FormFactor             string                                                                                                                         `json:"formFactor,omitempty"`             // Form Factor
	SiteHierarchy          string                                                                                                                         `json:"siteHierarchy,omitempty"`          // Site Hierarchy
	SiteHierarchyID        string                                                                                                                         `json:"siteHierarchyId,omitempty"`        // Site Hierarchy Id
	SiteID                 string                                                                                                                         `json:"siteId,omitempty"`                 // Site Id
	LastUpdatedTime        *int                                                                                                                           `json:"lastUpdatedTime,omitempty"`        // Last Updated Time
	ConnectionStatus       string                                                                                                                         `json:"connectionStatus,omitempty"`       // Connection Status
	Tracked                string                                                                                                                         `json:"tracked,omitempty"`                // Tracked
	IsPrivateMacAddress    *bool                                                                                                                          `json:"isPrivateMacAddress,omitempty"`    // Is Private Mac Address
	Health                 *ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1ResponseHealth                 `json:"health,omitempty"`                 //
	Traffic                *ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1ResponseTraffic                `json:"traffic,omitempty"`                //
	ConnectedNetworkDevice *ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1ResponseConnectedNetworkDevice `json:"connectedNetworkDevice,omitempty"` //
	Connection             *ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1ResponseConnection             `json:"connection,omitempty"`             //
	Onboarding             *ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1ResponseOnboarding             `json:"onboarding,omitempty"`             //
	Latency                *ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1ResponseLatency                `json:"latency,omitempty"`                //
}
type ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1ResponseHealth struct {
	OverallScore                 *int  `json:"overallScore,omitempty"`                 // Overall Score
	OnboardingScore              *int  `json:"onboardingScore,omitempty"`              // Onboarding Score
	ConnectedScore               *int  `json:"connectedScore,omitempty"`               // Connected Score
	LinkErrorPercentageThreshold *int  `json:"linkErrorPercentageThreshold,omitempty"` // Link Error Percentage Threshold
	IsLinkErrorIncluded          *bool `json:"isLinkErrorIncluded,omitempty"`          // Is Link Error Included
	RssiThreshold                *int  `json:"rssiThreshold,omitempty"`                // Rssi Threshold
	SnrThreshold                 *int  `json:"snrThreshold,omitempty"`                 // Snr Threshold
	IsRssiIncluded               *bool `json:"isRssiIncluded,omitempty"`               // Is Rssi Included
	IsSnrIncluded                *bool `json:"isSnrIncluded,omitempty"`                // Is Snr Included
}
type ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1ResponseTraffic struct {
	TxBytes               *int     `json:"txBytes,omitempty"`               // Tx Bytes
	RxBytes               *int     `json:"rxBytes,omitempty"`               // Rx Bytes
	Usage                 *int     `json:"usage,omitempty"`                 // Usage
	RxPackets             *int     `json:"rxPackets,omitempty"`             // Rx Packets
	TxPackets             *int     `json:"txPackets,omitempty"`             // Tx Packets
	RxRate                *float64 `json:"rxRate,omitempty"`                // Rx Rate
	TxRate                *int     `json:"txRate,omitempty"`                // Tx Rate
	RxLinkErrorPercentage *float64 `json:"rxLinkErrorPercentage,omitempty"` // Rx Link Error Percentage
	TxLinkErrorPercentage *float64 `json:"txLinkErrorPercentage,omitempty"` // Tx Link Error Percentage
	RxRetries             *int     `json:"rxRetries,omitempty"`             // Rx Retries
	RxRetryPercentage     *float64 `json:"rxRetryPercentage,omitempty"`     // Rx Retry Percentage
	TxDrops               *int     `json:"txDrops,omitempty"`               // Tx Drops
	TxDropPercentage      *int     `json:"txDropPercentage,omitempty"`      // Tx Drop Percentage
	DNSRequestCount       *int     `json:"dnsRequestCount,omitempty"`       // Dns Request Count
	DNSResponseCount      *int     `json:"dnsResponseCount,omitempty"`      // Dns Response Count
}
type ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1ResponseConnectedNetworkDevice struct {
	ConnectedNetworkDeviceID           string `json:"connectedNetworkDeviceId,omitempty"`           // Connected Network Device Id
	ConnectedNetworkDeviceName         string `json:"connectedNetworkDeviceName,omitempty"`         // Connected Network Device Name
	ConnectedNetworkDeviceManagementIP string `json:"connectedNetworkDeviceManagementIp,omitempty"` // Connected Network Device Management Ip
	ConnectedNetworkDeviceMac          string `json:"connectedNetworkDeviceMac,omitempty"`          // Connected Network Device Mac
	ConnectedNetworkDeviceType         string `json:"connectedNetworkDeviceType,omitempty"`         // Connected Network Device Type
	InterfaceName                      string `json:"interfaceName,omitempty"`                      // Interface Name
	InterfaceSpeed                     *int   `json:"interfaceSpeed,omitempty"`                     // Interface Speed
	DuplexMode                         string `json:"duplexMode,omitempty"`                         // Duplex Mode
}
type ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1ResponseConnection struct {
	VLANID                string `json:"vlanId,omitempty"`                // Vlan Id
	SessionDuration       *int   `json:"sessionDuration,omitempty"`       // Session Duration
	VnID                  string `json:"vnId,omitempty"`                  // Vn Id
	L2Vn                  string `json:"l2Vn,omitempty"`                  // L2 Vn
	L3Vn                  string `json:"l3Vn,omitempty"`                  // L3 Vn
	SecurityGroupTag      string `json:"securityGroupTag,omitempty"`      // Security Group Tag
	LinkSpeed             *int   `json:"linkSpeed,omitempty"`             // Link Speed
	BridgeVMMode          string `json:"bridgeVMMode,omitempty"`          // Bridge V M Mode
	Band                  string `json:"band,omitempty"`                  // Band
	SSID                  string `json:"ssid,omitempty"`                  // Ssid
	AuthType              string `json:"authType,omitempty"`              // Auth Type
	WlcName               string `json:"wlcName,omitempty"`               // Wlc Name
	WlcID                 string `json:"wlcId,omitempty"`                 // Wlc Id
	ApMac                 string `json:"apMac,omitempty"`                 // Ap Mac
	ApEthernetMac         string `json:"apEthernetMac,omitempty"`         // Ap Ethernet Mac
	ApMode                string `json:"apMode,omitempty"`                // Ap Mode
	RadioID               *int   `json:"radioId,omitempty"`               // Radio Id
	Channel               string `json:"channel,omitempty"`               // Channel
	ChannelWidth          string `json:"channelWidth,omitempty"`          // Channel Width
	Protocol              string `json:"protocol,omitempty"`              // Protocol
	ProtocolCapability    string `json:"protocolCapability,omitempty"`    // Protocol Capability
	UpnID                 string `json:"upnId,omitempty"`                 // Upn Id
	UpnName               string `json:"upnName,omitempty"`               // Upn Name
	UpnOwner              string `json:"upnOwner,omitempty"`              // Upn Owner
	UpnDuid               string `json:"upnDuid,omitempty"`               // Upn Duid
	Rssi                  *int   `json:"rssi,omitempty"`                  // Rssi
	Snr                   *int   `json:"snr,omitempty"`                   // Snr
	DataRate              *int   `json:"dataRate,omitempty"`              // Data Rate
	IsIosAnalyticsCapable *bool  `json:"isIosAnalyticsCapable,omitempty"` // Is Ios Analytics Capable
}
type ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1ResponseOnboarding struct {
	AvgRunDuration         *int   `json:"avgRunDuration,omitempty"`         // Avg Run Duration
	MaxRunDuration         *int   `json:"maxRunDuration,omitempty"`         // Max Run Duration
	AvgAssocDuration       *int   `json:"avgAssocDuration,omitempty"`       // Avg Assoc Duration
	MaxAssocDuration       *int   `json:"maxAssocDuration,omitempty"`       // Max Assoc Duration
	AvgAuthDuration        *int   `json:"avgAuthDuration,omitempty"`        // Avg Auth Duration
	MaxAuthDuration        *int   `json:"maxAuthDuration,omitempty"`        // Max Auth Duration
	AvgDhcpDuration        *int   `json:"avgDhcpDuration,omitempty"`        // Avg Dhcp Duration
	MaxDhcpDuration        *int   `json:"maxDhcpDuration,omitempty"`        // Max Dhcp Duration
	MaxRoamingDuration     *int   `json:"maxRoamingDuration,omitempty"`     // Max Roaming Duration
	AAAServerIP            string `json:"aaaServerIp,omitempty"`            // Aaa Server Ip
	DhcpServerIP           string `json:"dhcpServerIp,omitempty"`           // Dhcp Server Ip
	OnboardingTime         *int   `json:"onboardingTime,omitempty"`         // Onboarding Time
	AuthDoneTime           *int   `json:"authDoneTime,omitempty"`           // Auth Done Time
	AssocDoneTime          *int   `json:"assocDoneTime,omitempty"`          // Assoc Done Time
	DhcpDoneTime           *int   `json:"dhcpDoneTime,omitempty"`           // Dhcp Done Time
	RoamingTime            *int   `json:"roamingTime,omitempty"`            // Roaming Time
	FailedRoamingCount     *int   `json:"failedRoamingCount,omitempty"`     // Failed Roaming Count
	SuccessfulRoamingCount *int   `json:"successfulRoamingCount,omitempty"` // Successful Roaming Count
	TotalRoamingAttempts   *int   `json:"totalRoamingAttempts,omitempty"`   // Total Roaming Attempts
	AssocFailureReason     string `json:"assocFailureReason,omitempty"`     // Assoc Failure Reason
	AAAFailureReason       string `json:"aaaFailureReason,omitempty"`       // Aaa Failure Reason
	DhcpFailureReason      string `json:"dhcpFailureReason,omitempty"`      // Dhcp Failure Reason
	OtherFailureReason     string `json:"otherFailureReason,omitempty"`     // Other Failure Reason
	LatestFailureReason    string `json:"latestFailureReason,omitempty"`    // Latest Failure Reason
}
type ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1ResponseLatency struct {
	Video      *int `json:"video,omitempty"`      // Video
	Voice      *int `json:"voice,omitempty"`      // Voice
	BestEffort *int `json:"bestEffort,omitempty"` // Best Effort
	Background *int `json:"background,omitempty"` // Background
}
type ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1Page struct {
	Limit  *int                                                                                                         `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                                         `json:"offset,omitempty"` // Offset
	Count  *int                                                                                                         `json:"count,omitempty"`  // Count
	SortBy *[]ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1PageSortBy `json:"sortBy,omitempty"` //
}
type ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1PageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseClientsRetrievesTheTotalCountOfClientsByApplyingBasicFilteringV1 struct {
	Response *ResponseClientsRetrievesTheTotalCountOfClientsByApplyingBasicFilteringV1Response `json:"response,omitempty"` //
	Version  string                                                                            `json:"version,omitempty"`  // Version
}
type ResponseClientsRetrievesTheTotalCountOfClientsByApplyingBasicFilteringV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1 struct {
	Response *[]ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1Response `json:"response,omitempty"` //
	Page     *ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1Page       `json:"page,omitempty"`     //
	Version  string                                                                                                              `json:"version,omitempty"`  // Version
}
type ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1Response struct {
	ID                     string                                                                                                                                  `json:"id,omitempty"`                     // Id
	MacAddress             string                                                                                                                                  `json:"macAddress,omitempty"`             // Mac Address
	Type                   string                                                                                                                                  `json:"type,omitempty"`                   // Type
	Name                   string                                                                                                                                  `json:"name,omitempty"`                   // Name
	UserID                 string                                                                                                                                  `json:"userId,omitempty"`                 // User Id
	Username               string                                                                                                                                  `json:"username,omitempty"`               // Username
	IPv4Address            string                                                                                                                                  `json:"ipv4Address,omitempty"`            // Ipv4 Address
	IPv6Addresses          []string                                                                                                                                `json:"ipv6Addresses,omitempty"`          // Ipv6 Addresses
	Vendor                 string                                                                                                                                  `json:"vendor,omitempty"`                 // Vendor
	OsType                 string                                                                                                                                  `json:"osType,omitempty"`                 // Os Type
	OsVersion              string                                                                                                                                  `json:"osVersion,omitempty"`              // Os Version
	FormFactor             string                                                                                                                                  `json:"formFactor,omitempty"`             // Form Factor
	SiteHierarchy          string                                                                                                                                  `json:"siteHierarchy,omitempty"`          // Site Hierarchy
	SiteHierarchyID        string                                                                                                                                  `json:"siteHierarchyId,omitempty"`        // Site Hierarchy Id
	SiteID                 string                                                                                                                                  `json:"siteId,omitempty"`                 // Site Id
	LastUpdatedTime        *int                                                                                                                                    `json:"lastUpdatedTime,omitempty"`        // Last Updated Time
	ConnectionStatus       string                                                                                                                                  `json:"connectionStatus,omitempty"`       // Connection Status
	Tracked                string                                                                                                                                  `json:"tracked,omitempty"`                // Tracked
	IsPrivateMacAddress    *bool                                                                                                                                   `json:"isPrivateMacAddress,omitempty"`    // Is Private Mac Address
	Health                 *ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1ResponseHealth                 `json:"health,omitempty"`                 //
	Traffic                *ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1ResponseTraffic                `json:"traffic,omitempty"`                //
	ConnectedNetworkDevice *ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1ResponseConnectedNetworkDevice `json:"connectedNetworkDevice,omitempty"` //
	Connection             *ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1ResponseConnection             `json:"connection,omitempty"`             //
	Onboarding             *ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1ResponseOnboarding             `json:"onboarding,omitempty"`             //
	Latency                *ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1ResponseLatency                `json:"latency,omitempty"`                //
	AggregateAttributes    *[]ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1ResponseAggregateAttributes  `json:"aggregateAttributes,omitempty"`    //
}
type ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1ResponseHealth struct {
	OverallScore                 *int  `json:"overallScore,omitempty"`                 // Overall Score
	OnboardingScore              *int  `json:"onboardingScore,omitempty"`              // Onboarding Score
	ConnectedScore               *int  `json:"connectedScore,omitempty"`               // Connected Score
	LinkErrorPercentageThreshold *int  `json:"linkErrorPercentageThreshold,omitempty"` // Link Error Percentage Threshold
	IsLinkErrorIncluded          *bool `json:"isLinkErrorIncluded,omitempty"`          // Is Link Error Included
	RssiThreshold                *int  `json:"rssiThreshold,omitempty"`                // Rssi Threshold
	SnrThreshold                 *int  `json:"snrThreshold,omitempty"`                 // Snr Threshold
	IsRssiIncluded               *bool `json:"isRssiIncluded,omitempty"`               // Is Rssi Included
	IsSnrIncluded                *bool `json:"isSnrIncluded,omitempty"`                // Is Snr Included
}
type ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1ResponseTraffic struct {
	TxBytes               *int     `json:"txBytes,omitempty"`               // Tx Bytes
	RxBytes               *int     `json:"rxBytes,omitempty"`               // Rx Bytes
	Usage                 *int     `json:"usage,omitempty"`                 // Usage
	RxPackets             *int     `json:"rxPackets,omitempty"`             // Rx Packets
	TxPackets             *int     `json:"txPackets,omitempty"`             // Tx Packets
	RxRate                *float64 `json:"rxRate,omitempty"`                // Rx Rate
	TxRate                *int     `json:"txRate,omitempty"`                // Tx Rate
	RxLinkErrorPercentage *float64 `json:"rxLinkErrorPercentage,omitempty"` // Rx Link Error Percentage
	TxLinkErrorPercentage *float64 `json:"txLinkErrorPercentage,omitempty"` // Tx Link Error Percentage
	RxRetries             *int     `json:"rxRetries,omitempty"`             // Rx Retries
	RxRetryPercentage     *float64 `json:"rxRetryPercentage,omitempty"`     // Rx Retry Percentage
	TxDrops               *int     `json:"txDrops,omitempty"`               // Tx Drops
	TxDropPercentage      *int     `json:"txDropPercentage,omitempty"`      // Tx Drop Percentage
	DNSRequestCount       *int     `json:"dnsRequestCount,omitempty"`       // Dns Request Count
	DNSResponseCount      *int     `json:"dnsResponseCount,omitempty"`      // Dns Response Count
}
type ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1ResponseConnectedNetworkDevice struct {
	ConnectedNetworkDeviceID           string `json:"connectedNetworkDeviceId,omitempty"`           // Connected Network Device Id
	ConnectedNetworkDeviceName         string `json:"connectedNetworkDeviceName,omitempty"`         // Connected Network Device Name
	ConnectedNetworkDeviceManagementIP string `json:"connectedNetworkDeviceManagementIp,omitempty"` // Connected Network Device Management Ip
	ConnectedNetworkDeviceMac          string `json:"connectedNetworkDeviceMac,omitempty"`          // Connected Network Device Mac
	ConnectedNetworkDeviceType         string `json:"connectedNetworkDeviceType,omitempty"`         // Connected Network Device Type
	InterfaceName                      string `json:"interfaceName,omitempty"`                      // Interface Name
	InterfaceSpeed                     *int   `json:"interfaceSpeed,omitempty"`                     // Interface Speed
	DuplexMode                         string `json:"duplexMode,omitempty"`                         // Duplex Mode
}
type ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1ResponseConnection struct {
	VLANID                string `json:"vlanId,omitempty"`                // Vlan Id
	SessionDuration       *int   `json:"sessionDuration,omitempty"`       // Session Duration
	VnID                  string `json:"vnId,omitempty"`                  // Vn Id
	L2Vn                  string `json:"l2Vn,omitempty"`                  // L2 Vn
	L3Vn                  string `json:"l3Vn,omitempty"`                  // L3 Vn
	SecurityGroupTag      string `json:"securityGroupTag,omitempty"`      // Security Group Tag
	LinkSpeed             *int   `json:"linkSpeed,omitempty"`             // Link Speed
	BridgeVMMode          string `json:"bridgeVMMode,omitempty"`          // Bridge V M Mode
	Band                  string `json:"band,omitempty"`                  // Band
	SSID                  string `json:"ssid,omitempty"`                  // Ssid
	AuthType              string `json:"authType,omitempty"`              // Auth Type
	WlcName               string `json:"wlcName,omitempty"`               // Wlc Name
	WlcID                 string `json:"wlcId,omitempty"`                 // Wlc Id
	ApMac                 string `json:"apMac,omitempty"`                 // Ap Mac
	ApEthernetMac         string `json:"apEthernetMac,omitempty"`         // Ap Ethernet Mac
	ApMode                string `json:"apMode,omitempty"`                // Ap Mode
	RadioID               *int   `json:"radioId,omitempty"`               // Radio Id
	Channel               string `json:"channel,omitempty"`               // Channel
	ChannelWidth          string `json:"channelWidth,omitempty"`          // Channel Width
	Protocol              string `json:"protocol,omitempty"`              // Protocol
	ProtocolCapability    string `json:"protocolCapability,omitempty"`    // Protocol Capability
	UpnID                 string `json:"upnId,omitempty"`                 // Upn Id
	UpnName               string `json:"upnName,omitempty"`               // Upn Name
	UpnOwner              string `json:"upnOwner,omitempty"`              // Upn Owner
	UpnDuid               string `json:"upnDuid,omitempty"`               // Upn Duid
	Rssi                  *int   `json:"rssi,omitempty"`                  // Rssi
	Snr                   *int   `json:"snr,omitempty"`                   // Snr
	DataRate              *int   `json:"dataRate,omitempty"`              // Data Rate
	IsIosAnalyticsCapable *bool  `json:"isIosAnalyticsCapable,omitempty"` // Is Ios Analytics Capable
}
type ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1ResponseOnboarding struct {
	AvgRunDuration         *int   `json:"avgRunDuration,omitempty"`         // Avg Run Duration
	MaxRunDuration         *int   `json:"maxRunDuration,omitempty"`         // Max Run Duration
	AvgAssocDuration       *int   `json:"avgAssocDuration,omitempty"`       // Avg Assoc Duration
	MaxAssocDuration       *int   `json:"maxAssocDuration,omitempty"`       // Max Assoc Duration
	AvgAuthDuration        *int   `json:"avgAuthDuration,omitempty"`        // Avg Auth Duration
	MaxAuthDuration        *int   `json:"maxAuthDuration,omitempty"`        // Max Auth Duration
	AvgDhcpDuration        *int   `json:"avgDhcpDuration,omitempty"`        // Avg Dhcp Duration
	MaxDhcpDuration        *int   `json:"maxDhcpDuration,omitempty"`        // Max Dhcp Duration
	MaxRoamingDuration     *int   `json:"maxRoamingDuration,omitempty"`     // Max Roaming Duration
	AAAServerIP            string `json:"aaaServerIp,omitempty"`            // Aaa Server Ip
	DhcpServerIP           string `json:"dhcpServerIp,omitempty"`           // Dhcp Server Ip
	OnboardingTime         *int   `json:"onboardingTime,omitempty"`         // Onboarding Time
	AuthDoneTime           *int   `json:"authDoneTime,omitempty"`           // Auth Done Time
	AssocDoneTime          *int   `json:"assocDoneTime,omitempty"`          // Assoc Done Time
	DhcpDoneTime           *int   `json:"dhcpDoneTime,omitempty"`           // Dhcp Done Time
	RoamingTime            *int   `json:"roamingTime,omitempty"`            // Roaming Time
	FailedRoamingCount     *int   `json:"failedRoamingCount,omitempty"`     // Failed Roaming Count
	SuccessfulRoamingCount *int   `json:"successfulRoamingCount,omitempty"` // Successful Roaming Count
	TotalRoamingAttempts   *int   `json:"totalRoamingAttempts,omitempty"`   // Total Roaming Attempts
	AssocFailureReason     string `json:"assocFailureReason,omitempty"`     // Assoc Failure Reason
	AAAFailureReason       string `json:"aaaFailureReason,omitempty"`       // Aaa Failure Reason
	DhcpFailureReason      string `json:"dhcpFailureReason,omitempty"`      // Dhcp Failure Reason
	OtherFailureReason     string `json:"otherFailureReason,omitempty"`     // Other Failure Reason
	LatestFailureReason    string `json:"latestFailureReason,omitempty"`    // Latest Failure Reason
}
type ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1ResponseLatency struct {
	Video      *int `json:"video,omitempty"`      // Video
	Voice      *int `json:"voice,omitempty"`      // Voice
	BestEffort *int `json:"bestEffort,omitempty"` // Best Effort
	Background *int `json:"background,omitempty"` // Background
}
type ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1ResponseAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1Page struct {
	Limit  *int                                                                                                                  `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                                                  `json:"offset,omitempty"` // Offset
	Count  *int                                                                                                                  `json:"count,omitempty"`  // Count
	SortBy *[]ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1PageSortBy `json:"sortBy,omitempty"` //
}
type ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1PageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseClientsRetrievesTheNumberOfClientsByApplyingComplexFiltersV1 struct {
	Response *ResponseClientsRetrievesTheNumberOfClientsByApplyingComplexFiltersV1Response `json:"response,omitempty"` //
	Version  string                                                                        `json:"version,omitempty"`  // Version
}
type ResponseClientsRetrievesTheNumberOfClientsByApplyingComplexFiltersV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseClientsRetrievesSummaryAnalyticsDataRelatedToClientsV1 struct {
	Response *ResponseClientsRetrievesSummaryAnalyticsDataRelatedToClientsV1Response `json:"response,omitempty"` //
	Page     *ResponseClientsRetrievesSummaryAnalyticsDataRelatedToClientsV1Page     `json:"page,omitempty"`     //
	Version  string                                                                  `json:"version,omitempty"`  // Version
}
type ResponseClientsRetrievesSummaryAnalyticsDataRelatedToClientsV1Response struct {
	Groups *[]ResponseClientsRetrievesSummaryAnalyticsDataRelatedToClientsV1ResponseGroups `json:"groups,omitempty"` //
}
type ResponseClientsRetrievesSummaryAnalyticsDataRelatedToClientsV1ResponseGroups struct {
	ID                  string                                                                                             `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseClientsRetrievesSummaryAnalyticsDataRelatedToClientsV1ResponseGroupsAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseClientsRetrievesSummaryAnalyticsDataRelatedToClientsV1ResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseClientsRetrievesSummaryAnalyticsDataRelatedToClientsV1ResponseGroupsAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value *int   `json:"value,omitempty"` // Value
}
type ResponseClientsRetrievesSummaryAnalyticsDataRelatedToClientsV1ResponseGroupsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseClientsRetrievesSummaryAnalyticsDataRelatedToClientsV1Page struct {
	Limit  *int                                                                        `json:"limit,omitempty"`  // Limit
	Cursor string                                                                      `json:"cursor,omitempty"` // Cursor
	Count  *int                                                                        `json:"count,omitempty"`  // Count
	SortBy *[]ResponseClientsRetrievesSummaryAnalyticsDataRelatedToClientsV1PageSortBy `json:"sortBy,omitempty"` //
}
type ResponseClientsRetrievesSummaryAnalyticsDataRelatedToClientsV1PageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseClientsRetrievesTheTopNAnalyticsDataRelatedToClientsV1 struct {
	Response *[]ResponseClientsRetrievesTheTopNAnalyticsDataRelatedToClientsV1Response `json:"response,omitempty"` //
	Page     *ResponseClientsRetrievesTheTopNAnalyticsDataRelatedToClientsV1Page       `json:"page,omitempty"`     //
	Version  string                                                                    `json:"version,omitempty"`  // Version
}
type ResponseClientsRetrievesTheTopNAnalyticsDataRelatedToClientsV1Response struct {
	ID                  string                                                                                       `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseClientsRetrievesTheTopNAnalyticsDataRelatedToClientsV1ResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseClientsRetrievesTheTopNAnalyticsDataRelatedToClientsV1ResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseClientsRetrievesTheTopNAnalyticsDataRelatedToClientsV1ResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value *int   `json:"value,omitempty"` // Value
}
type ResponseClientsRetrievesTheTopNAnalyticsDataRelatedToClientsV1ResponseAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseClientsRetrievesTheTopNAnalyticsDataRelatedToClientsV1Page struct {
	Limit  *int                                                                        `json:"limit,omitempty"`  // Limit
	Cursor string                                                                      `json:"cursor,omitempty"` // Cursor
	Count  *int                                                                        `json:"count,omitempty"`  // Count
	SortBy *[]ResponseClientsRetrievesTheTopNAnalyticsDataRelatedToClientsV1PageSortBy `json:"sortBy,omitempty"` //
}
type ResponseClientsRetrievesTheTopNAnalyticsDataRelatedToClientsV1PageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseClientsRetrievesTheTrendAnalyticsDataRelatedToClientsV1 struct {
	Response *[]ResponseClientsRetrievesTheTrendAnalyticsDataRelatedToClientsV1Response `json:"response,omitempty"` //
	Page     *ResponseClientsRetrievesTheTrendAnalyticsDataRelatedToClientsV1Page       `json:"page,omitempty"`     //
	Version  string                                                                     `json:"version,omitempty"`  // Version
}
type ResponseClientsRetrievesTheTrendAnalyticsDataRelatedToClientsV1Response struct {
	Timestamp *int                                                                             `json:"timestamp,omitempty"` // Timestamp
	Groups    *[]ResponseClientsRetrievesTheTrendAnalyticsDataRelatedToClientsV1ResponseGroups `json:"groups,omitempty"`    //
}
type ResponseClientsRetrievesTheTrendAnalyticsDataRelatedToClientsV1ResponseGroups struct {
	ID                  string                                                                                              `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseClientsRetrievesTheTrendAnalyticsDataRelatedToClientsV1ResponseGroupsAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseClientsRetrievesTheTrendAnalyticsDataRelatedToClientsV1ResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseClientsRetrievesTheTrendAnalyticsDataRelatedToClientsV1ResponseGroupsAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value *int   `json:"value,omitempty"` // Value
}
type ResponseClientsRetrievesTheTrendAnalyticsDataRelatedToClientsV1ResponseGroupsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseClientsRetrievesTheTrendAnalyticsDataRelatedToClientsV1Page struct {
	Limit         *int   `json:"limit,omitempty"`         // Limit
	Cursor        string `json:"cursor,omitempty"`        // Cursor
	Count         *int   `json:"count,omitempty"`         // Count
	TimeSortOrder string `json:"timeSortOrder,omitempty"` // Time Sort Order
}
type ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressV1 struct {
	Response *ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressV1Response `json:"response,omitempty"` //
	Version  string                                                                            `json:"version,omitempty"`  // Version
}
type ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressV1Response struct {
	ID                     string                                                                                                  `json:"id,omitempty"`                     // Id
	MacAddress             string                                                                                                  `json:"macAddress,omitempty"`             // Mac Address
	Type                   string                                                                                                  `json:"type,omitempty"`                   // Type
	Name                   string                                                                                                  `json:"name,omitempty"`                   // Name
	UserID                 string                                                                                                  `json:"userId,omitempty"`                 // User Id
	Username               string                                                                                                  `json:"username,omitempty"`               // Username
	IPv4Address            string                                                                                                  `json:"ipv4Address,omitempty"`            // Ipv4 Address
	IPv6Addresses          []string                                                                                                `json:"ipv6Addresses,omitempty"`          // Ipv6 Addresses
	Vendor                 string                                                                                                  `json:"vendor,omitempty"`                 // Vendor
	OsType                 string                                                                                                  `json:"osType,omitempty"`                 // Os Type
	OsVersion              string                                                                                                  `json:"osVersion,omitempty"`              // Os Version
	FormFactor             string                                                                                                  `json:"formFactor,omitempty"`             // Form Factor
	SiteHierarchy          string                                                                                                  `json:"siteHierarchy,omitempty"`          // Site Hierarchy
	SiteHierarchyID        string                                                                                                  `json:"siteHierarchyId,omitempty"`        // Site Hierarchy Id
	SiteID                 string                                                                                                  `json:"siteId,omitempty"`                 // Site Id
	LastUpdatedTime        *int                                                                                                    `json:"lastUpdatedTime,omitempty"`        // Last Updated Time
	ConnectionStatus       string                                                                                                  `json:"connectionStatus,omitempty"`       // Connection Status
	Tracked                string                                                                                                  `json:"tracked,omitempty"`                // Tracked
	IsPrivateMacAddress    *bool                                                                                                   `json:"isPrivateMacAddress,omitempty"`    // Is Private Mac Address
	Health                 *ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressV1ResponseHealth                 `json:"health,omitempty"`                 //
	Traffic                *ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressV1ResponseTraffic                `json:"traffic,omitempty"`                //
	ConnectedNetworkDevice *ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressV1ResponseConnectedNetworkDevice `json:"connectedNetworkDevice,omitempty"` //
	Connection             *ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressV1ResponseConnection             `json:"connection,omitempty"`             //
	Onboarding             *ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressV1ResponseOnboarding             `json:"onboarding,omitempty"`             //
	Latency                *ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressV1ResponseLatency                `json:"latency,omitempty"`                //
}
type ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressV1ResponseHealth struct {
	OverallScore                 *int  `json:"overallScore,omitempty"`                 // Overall Score
	OnboardingScore              *int  `json:"onboardingScore,omitempty"`              // Onboarding Score
	ConnectedScore               *int  `json:"connectedScore,omitempty"`               // Connected Score
	LinkErrorPercentageThreshold *int  `json:"linkErrorPercentageThreshold,omitempty"` // Link Error Percentage Threshold
	IsLinkErrorIncluded          *bool `json:"isLinkErrorIncluded,omitempty"`          // Is Link Error Included
	RssiThreshold                *int  `json:"rssiThreshold,omitempty"`                // Rssi Threshold
	SnrThreshold                 *int  `json:"snrThreshold,omitempty"`                 // Snr Threshold
	IsRssiIncluded               *bool `json:"isRssiIncluded,omitempty"`               // Is Rssi Included
	IsSnrIncluded                *bool `json:"isSnrIncluded,omitempty"`                // Is Snr Included
}
type ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressV1ResponseTraffic struct {
	TxBytes               *int     `json:"txBytes,omitempty"`               // Tx Bytes
	RxBytes               *int     `json:"rxBytes,omitempty"`               // Rx Bytes
	Usage                 *int     `json:"usage,omitempty"`                 // Usage
	RxPackets             *int     `json:"rxPackets,omitempty"`             // Rx Packets
	TxPackets             *int     `json:"txPackets,omitempty"`             // Tx Packets
	RxRate                *float64 `json:"rxRate,omitempty"`                // Rx Rate
	TxRate                *int     `json:"txRate,omitempty"`                // Tx Rate
	RxLinkErrorPercentage *float64 `json:"rxLinkErrorPercentage,omitempty"` // Rx Link Error Percentage
	TxLinkErrorPercentage *float64 `json:"txLinkErrorPercentage,omitempty"` // Tx Link Error Percentage
	RxRetries             *int     `json:"rxRetries,omitempty"`             // Rx Retries
	RxRetryPercentage     *float64 `json:"rxRetryPercentage,omitempty"`     // Rx Retry Percentage
	TxDrops               *int     `json:"txDrops,omitempty"`               // Tx Drops
	TxDropPercentage      *int     `json:"txDropPercentage,omitempty"`      // Tx Drop Percentage
	DNSRequestCount       *int     `json:"dnsRequestCount,omitempty"`       // Dns Request Count
	DNSResponseCount      *int     `json:"dnsResponseCount,omitempty"`      // Dns Response Count
}
type ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressV1ResponseConnectedNetworkDevice struct {
	ConnectedNetworkDeviceID           string `json:"connectedNetworkDeviceId,omitempty"`           // Connected Network Device Id
	ConnectedNetworkDeviceName         string `json:"connectedNetworkDeviceName,omitempty"`         // Connected Network Device Name
	ConnectedNetworkDeviceManagementIP string `json:"connectedNetworkDeviceManagementIp,omitempty"` // Connected Network Device Management Ip
	ConnectedNetworkDeviceMac          string `json:"connectedNetworkDeviceMac,omitempty"`          // Connected Network Device Mac
	ConnectedNetworkDeviceType         string `json:"connectedNetworkDeviceType,omitempty"`         // Connected Network Device Type
	InterfaceName                      string `json:"interfaceName,omitempty"`                      // Interface Name
	InterfaceSpeed                     *int   `json:"interfaceSpeed,omitempty"`                     // Interface Speed
	DuplexMode                         string `json:"duplexMode,omitempty"`                         // Duplex Mode
}
type ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressV1ResponseConnection struct {
	VLANID                string `json:"vlanId,omitempty"`                // Vlan Id
	SessionDuration       *int   `json:"sessionDuration,omitempty"`       // Session Duration
	VnID                  string `json:"vnId,omitempty"`                  // Vn Id
	L2Vn                  string `json:"l2Vn,omitempty"`                  // L2 Vn
	L3Vn                  string `json:"l3Vn,omitempty"`                  // L3 Vn
	SecurityGroupTag      string `json:"securityGroupTag,omitempty"`      // Security Group Tag
	LinkSpeed             *int   `json:"linkSpeed,omitempty"`             // Link Speed
	BridgeVMMode          string `json:"bridgeVMMode,omitempty"`          // Bridge V M Mode
	Band                  string `json:"band,omitempty"`                  // Band
	SSID                  string `json:"ssid,omitempty"`                  // Ssid
	AuthType              string `json:"authType,omitempty"`              // Auth Type
	WlcName               string `json:"wlcName,omitempty"`               // Wlc Name
	WlcID                 string `json:"wlcId,omitempty"`                 // Wlc Id
	ApMac                 string `json:"apMac,omitempty"`                 // Ap Mac
	ApEthernetMac         string `json:"apEthernetMac,omitempty"`         // Ap Ethernet Mac
	ApMode                string `json:"apMode,omitempty"`                // Ap Mode
	RadioID               *int   `json:"radioId,omitempty"`               // Radio Id
	Channel               string `json:"channel,omitempty"`               // Channel
	ChannelWidth          string `json:"channelWidth,omitempty"`          // Channel Width
	Protocol              string `json:"protocol,omitempty"`              // Protocol
	ProtocolCapability    string `json:"protocolCapability,omitempty"`    // Protocol Capability
	UpnID                 string `json:"upnId,omitempty"`                 // Upn Id
	UpnName               string `json:"upnName,omitempty"`               // Upn Name
	UpnOwner              string `json:"upnOwner,omitempty"`              // Upn Owner
	UpnDuid               string `json:"upnDuid,omitempty"`               // Upn Duid
	Rssi                  *int   `json:"rssi,omitempty"`                  // Rssi
	Snr                   *int   `json:"snr,omitempty"`                   // Snr
	DataRate              *int   `json:"dataRate,omitempty"`              // Data Rate
	IsIosAnalyticsCapable *bool  `json:"isIosAnalyticsCapable,omitempty"` // Is Ios Analytics Capable
}
type ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressV1ResponseOnboarding struct {
	AvgRunDuration         *int   `json:"avgRunDuration,omitempty"`         // Avg Run Duration
	MaxRunDuration         *int   `json:"maxRunDuration,omitempty"`         // Max Run Duration
	AvgAssocDuration       *int   `json:"avgAssocDuration,omitempty"`       // Avg Assoc Duration
	MaxAssocDuration       *int   `json:"maxAssocDuration,omitempty"`       // Max Assoc Duration
	AvgAuthDuration        *int   `json:"avgAuthDuration,omitempty"`        // Avg Auth Duration
	MaxAuthDuration        *int   `json:"maxAuthDuration,omitempty"`        // Max Auth Duration
	AvgDhcpDuration        *int   `json:"avgDhcpDuration,omitempty"`        // Avg Dhcp Duration
	MaxDhcpDuration        *int   `json:"maxDhcpDuration,omitempty"`        // Max Dhcp Duration
	MaxRoamingDuration     *int   `json:"maxRoamingDuration,omitempty"`     // Max Roaming Duration
	AAAServerIP            string `json:"aaaServerIp,omitempty"`            // Aaa Server Ip
	DhcpServerIP           string `json:"dhcpServerIp,omitempty"`           // Dhcp Server Ip
	OnboardingTime         *int   `json:"onboardingTime,omitempty"`         // Onboarding Time
	AuthDoneTime           *int   `json:"authDoneTime,omitempty"`           // Auth Done Time
	AssocDoneTime          *int   `json:"assocDoneTime,omitempty"`          // Assoc Done Time
	DhcpDoneTime           *int   `json:"dhcpDoneTime,omitempty"`           // Dhcp Done Time
	RoamingTime            *int   `json:"roamingTime,omitempty"`            // Roaming Time
	FailedRoamingCount     *int   `json:"failedRoamingCount,omitempty"`     // Failed Roaming Count
	SuccessfulRoamingCount *int   `json:"successfulRoamingCount,omitempty"` // Successful Roaming Count
	TotalRoamingAttempts   *int   `json:"totalRoamingAttempts,omitempty"`   // Total Roaming Attempts
	AssocFailureReason     string `json:"assocFailureReason,omitempty"`     // Assoc Failure Reason
	AAAFailureReason       string `json:"aaaFailureReason,omitempty"`       // Aaa Failure Reason
	DhcpFailureReason      string `json:"dhcpFailureReason,omitempty"`      // Dhcp Failure Reason
	OtherFailureReason     string `json:"otherFailureReason,omitempty"`     // Other Failure Reason
	LatestFailureReason    string `json:"latestFailureReason,omitempty"`    // Latest Failure Reason
}
type ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressV1ResponseLatency struct {
	Video      *int `json:"video,omitempty"`      // Video
	Voice      *int `json:"voice,omitempty"`      // Voice
	BestEffort *int `json:"bestEffort,omitempty"` // Best Effort
	Background *int `json:"background,omitempty"` // Background
}
type ResponseClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1 struct {
	Response *[]ResponseClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1Response `json:"response,omitempty"` //
	Page     *ResponseClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1Page       `json:"page,omitempty"`     //
	Version  string                                                                                   `json:"version,omitempty"`  // Version
}
type ResponseClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1Response struct {
	Timestamp *int                                                                                           `json:"timestamp,omitempty"` // Timestamp
	Groups    *[]ResponseClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1ResponseGroups `json:"groups,omitempty"`    //
}
type ResponseClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1ResponseGroups struct {
	ID                  string                                                                                                            `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1ResponseGroupsAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1ResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1ResponseGroupsAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value *int   `json:"value,omitempty"` // Value
}
type ResponseClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1ResponseGroupsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1Page struct {
	Limit         *int   `json:"limit,omitempty"`         // Limit
	Cursor        string `json:"cursor,omitempty"`        // Cursor
	Count         *int   `json:"count,omitempty"`         // Count
	TimeSortOrder string `json:"timeSortOrder,omitempty"` // Time Sort Order
}
type ResponseClientsGetClientDetailV1 struct {
	Detail         *ResponseClientsGetClientDetailV1Detail         `json:"detail,omitempty"`         //
	ConnectionInfo *ResponseClientsGetClientDetailV1ConnectionInfo `json:"connectionInfo,omitempty"` //
	Topology       *ResponseClientsGetClientDetailV1Topology       `json:"topology,omitempty"`       //
}
type ResponseClientsGetClientDetailV1Detail struct {
	ID                           string                                                   `json:"id,omitempty"`                           // Unique identifier representing a specific host
	ConnectionStatus             string                                                   `json:"connectionStatus,omitempty"`             // The client is connected, connecting, or disconnected
	Tracked                      string                                                   `json:"tracked,omitempty"`                      // Tracking status of this host
	HostType                     string                                                   `json:"hostType,omitempty"`                     // WIRED or WIRELESS
	UserID                       string                                                   `json:"userId,omitempty"`                       // The user ID of this host
	Duid                         string                                                   `json:"duid,omitempty"`                         // Device UID for MAC
	IDentifier                   string                                                   `json:"identifier,omitempty"`                   // The host's unique identifier, which is populated by and in order of userId, hostName, hostIpV4, hostIpV6, or hostMac
	HostName                     string                                                   `json:"hostName,omitempty"`                     // The hostname of the host
	HostOs                       string                                                   `json:"hostOs,omitempty"`                       // The OS of host
	HostVersion                  string                                                   `json:"hostVersion,omitempty"`                  // The version of OS of host
	SubType                      string                                                   `json:"subType,omitempty"`                      // The device type of host
	FirmwareVersion              string                                                   `json:"firmwareVersion,omitempty"`              // The firmware version of the host
	DeviceVendor                 string                                                   `json:"deviceVendor,omitempty"`                 // The device vendor string
	DeviceForm                   string                                                   `json:"deviceForm,omitempty"`                   // The device form of the host (e.g. Phone/Tablet)
	SalesCode                    string                                                   `json:"salesCode,omitempty"`                    // The Sales Code of the host
	CountryCode                  string                                                   `json:"countryCode,omitempty"`                  // The country code of the host
	LastUpdated                  *int                                                     `json:"lastUpdated,omitempty"`                  // Epoch/Unix time in milliseconds
	HealthScore                  *[]ResponseClientsGetClientDetailV1DetailHealthScore     `json:"healthScore,omitempty"`                  //
	HostMac                      string                                                   `json:"hostMac,omitempty"`                      // MAC address of the interface
	HostIPV4                     string                                                   `json:"hostIpV4,omitempty"`                     // IPv4 address of the interface
	HostIPV6                     []string                                                 `json:"hostIpV6,omitempty"`                     // List of IPv6 addresses
	AuthType                     string                                                   `json:"authType,omitempty"`                     // Authentication mechanism of the client
	VLANID                       *int                                                     `json:"vlanId,omitempty"`                       // VLAN ID for the host
	L3VirtualNetwork             string                                                   `json:"l3VirtualNetwork,omitempty"`             // Comma separated Level 3 virtual network names
	L2VirtualNetwork             string                                                   `json:"l2VirtualNetwork,omitempty"`             // Comma separated Level 2 virtual network names
	Vnid                         *int                                                     `json:"vnid,omitempty"`                         // VNID of the host
	UpnID                        string                                                   `json:"upnId,omitempty"`                        // Registered UPN ID of the host
	UpnName                      string                                                   `json:"upnName,omitempty"`                      // Registered UPN name of the host
	SSID                         string                                                   `json:"ssid,omitempty"`                         // WLAN SSID to which the client is connected
	Frequency                    string                                                   `json:"frequency,omitempty"`                    // Frequency band to which the client is connected
	Channel                      string                                                   `json:"channel,omitempty"`                      // Channel to which the client is connected
	ApGroup                      string                                                   `json:"apGroup,omitempty"`                      // AP group to which the client belongs
	Sgt                          string                                                   `json:"sgt,omitempty"`                          // Security group tag
	Location                     string                                                   `json:"location,omitempty"`                     // Site location of client
	ClientConnection             string                                                   `json:"clientConnection,omitempty"`             // AP/Switch to which the client is connected
	ConnectedDevice              *[]ResponseClientsGetClientDetailV1DetailConnectedDevice `json:"connectedDevice,omitempty"`              //
	IssueCount                   *int                                                     `json:"issueCount,omitempty"`                   // Issue count for a device
	Rssi                         string                                                   `json:"rssi,omitempty"`                         // Min RSSI value for the client
	RssiThreshold                string                                                   `json:"rssiThreshold,omitempty"`                // RSSI threshold
	RssiIsInclude                string                                                   `json:"rssiIsInclude,omitempty"`                // RSSI include/exclude flag
	AvgRssi                      string                                                   `json:"avgRssi,omitempty"`                      // Average RSSI value for the client
	Snr                          string                                                   `json:"snr,omitempty"`                          // Min signal to noise ratio for the client
	SnrThreshold                 string                                                   `json:"snrThreshold,omitempty"`                 // SNR threshold
	SnrIsInclude                 string                                                   `json:"snrIsInclude,omitempty"`                 // SNR include/exclude flag
	AvgSnr                       string                                                   `json:"avgSnr,omitempty"`                       // Average signal to noise ratio for a client
	DataRate                     string                                                   `json:"dataRate,omitempty"`                     // MCS data rates for a client
	TxBytes                      string                                                   `json:"txBytes,omitempty"`                      // total transmitted bytes for a client
	RxBytes                      string                                                   `json:"rxBytes,omitempty"`                      // Total received bytes for a client
	DNSResponse                  string                                                   `json:"dnsResponse,omitempty"`                  // DNS response attempts for a client
	DNSRequest                   string                                                   `json:"dnsRequest,omitempty"`                   // DNS request attempts for a client
	Onboarding                   *ResponseClientsGetClientDetailV1DetailOnboarding        `json:"onboarding,omitempty"`                   //
	ClientType                   string                                                   `json:"clientType,omitempty"`                   // OLD or NEW
	OnboardingTime               *int                                                     `json:"onboardingTime,omitempty"`               // Epoch/Unix time in milliseconds
	Port                         string                                                   `json:"port,omitempty"`                         // switch port for client
	IosCapable                   *bool                                                    `json:"iosCapable,omitempty"`                   // IOS Capable
	Usage                        *float64                                                 `json:"usage,omitempty"`                        // Usage of txBytes and rxBytes of client
	LinkSpeed                    *float64                                                 `json:"linkSpeed,omitempty"`                    // The speed of wired client
	LinkThreshold                string                                                   `json:"linkThreshold,omitempty"`                // Link error threshold of wired client
	RemoteEndDuplexMode          string                                                   `json:"remoteEndDuplexMode,omitempty"`          // The remote end duplex mode of wired client
	TxLinkError                  *float64                                                 `json:"txLinkError,omitempty"`                  // The error of tx link
	RxLinkError                  *float64                                                 `json:"rxLinkError,omitempty"`                  // The error of rx link
	TxRate                       *float64                                                 `json:"txRate,omitempty"`                       // The rate of tx
	RxRate                       *float64                                                 `json:"rxRate,omitempty"`                       // The rate of rx
	RxRetryPct                   string                                                   `json:"rxRetryPct,omitempty"`                   // The retry count as percentage wrt to total rx packets
	VersionTime                  *int                                                     `json:"versionTime,omitempty"`                  // The metric modification time of the new version
	Dot11Protocol                string                                                   `json:"dot11Protocol,omitempty"`                // Description of dot11 protocol
	SlotID                       *int                                                     `json:"slotId,omitempty"`                       // Slot ID of AP which client is connected
	Dot11ProtocolCapability      string                                                   `json:"dot11ProtocolCapability,omitempty"`      // description of dot11 protocol capability
	PrivateMac                   *bool                                                    `json:"privateMac,omitempty"`                   // Private Mac
	DhcpServerIP                 string                                                   `json:"dhcpServerIp,omitempty"`                 // The DHCP server IP
	AAAServerIP                  string                                                   `json:"aaaServerIp,omitempty"`                  // The AAA server IP
	AAAServerTransaction         *int                                                     `json:"aaaServerTransaction,omitempty"`         // The number of AAA server transactions
	AAAServerFailedTransaction   *int                                                     `json:"aaaServerFailedTransaction,omitempty"`   // The number of failed AAA server transactions
	AAAServerSuccessTransaction  *int                                                     `json:"aaaServerSuccessTransaction,omitempty"`  // The number of successful AAA server transactions
	AAAServerLatency             *float64                                                 `json:"aaaServerLatency,omitempty"`             // The AAA server latency
	AAAServerMABLatency          *float64                                                 `json:"aaaServerMABLatency,omitempty"`          // The AAA server MAB latency
	AAAServerEApLatency          *float64                                                 `json:"aaaServerEAPLatency,omitempty"`          // The AAA server EAP latency
	DhcpServerTransaction        *int                                                     `json:"dhcpServerTransaction,omitempty"`        // The number of DHCP server transactions
	DhcpServerFailedTransaction  *int                                                     `json:"dhcpServerFailedTransaction,omitempty"`  // The number of failed DHCP server transactions
	DhcpServerSuccessTransaction *int                                                     `json:"dhcpServerSuccessTransaction,omitempty"` // The number of successful DHCP server transactions
	DhcpServerLatency            *float64                                                 `json:"dhcpServerLatency,omitempty"`            // The DHCP server latency
	DhcpServerDOLatency          *float64                                                 `json:"dhcpServerDOLatency,omitempty"`          // The DHCP server DO latency
	DhcpServerRALatency          *float64                                                 `json:"dhcpServerRALatency,omitempty"`          // The DHCP RA latency
	MaxRoamingDuration           string                                                   `json:"maxRoamingDuration,omitempty"`           // Max roaming duration for a client
	UpnOwner                     string                                                   `json:"upnOwner,omitempty"`                     // Owner of registered UPN name
	ConnectedUpn                 string                                                   `json:"connectedUpn,omitempty"`                 // connected UPN ID
	ConnectedUpnOwner            string                                                   `json:"connectedUpnOwner,omitempty"`            // Connected UPN owner
	ConnectedUpnID               string                                                   `json:"connectedUpnId,omitempty"`               // Connected UPN ID
	IsGuestUPNEndpoint           *bool                                                    `json:"isGuestUPNEndpoint,omitempty"`           // Whether it is guest UPN endpoint
	WlcName                      string                                                   `json:"wlcName,omitempty"`                      // The name of the connected wireless controller
	WlcUUID                      string                                                   `json:"wlcUuid,omitempty"`                      // UUID of the WLC the client connected to
	SessionDuration              string                                                   `json:"sessionDuration,omitempty"`              // Time duration the session took from run time to delete time
	IntelCapable                 *bool                                                    `json:"intelCapable,omitempty"`                 // Whether support Intel device analytics
	HwModel                      string                                                   `json:"hwModel,omitempty"`                      // Hardware model
	PowerType                    string                                                   `json:"powerType,omitempty"`                    // AC/DC voltage
	ModelName                    string                                                   `json:"modelName,omitempty"`                    // System model
	BridgeVMMode                 string                                                   `json:"bridgeVMMode,omitempty"`                 // Bridge VM mode
	DhcpNakIP                    string                                                   `json:"dhcpNakIp,omitempty"`                    // DHCP NAK IP
	DhcpDeclineIP                string                                                   `json:"dhcpDeclineIp,omitempty"`                // DHCP decline IP
	PortDescription              string                                                   `json:"portDescription,omitempty"`              // Port desctiption of wired client
	LatencyVoice                 *float64                                                 `json:"latencyVoice,omitempty"`                 // Voice latency
	LatencyVideo                 *float64                                                 `json:"latencyVideo,omitempty"`                 // Video latency
	LatencyBg                    *float64                                                 `json:"latencyBg,omitempty"`                    // Background latency
	LatencyBe                    *float64                                                 `json:"latencyBe,omitempty"`                    // Best-effort latency
	TrustScore                   string                                                   `json:"trustScore,omitempty"`                   // Trust score of Client received from EndPoint Analytics
	TrustDetails                 string                                                   `json:"trustDetails,omitempty"`                 // Trust details explaining reason for corresponding Trust score
}
type ResponseClientsGetClientDetailV1DetailHealthScore struct {
	HealthType string `json:"healthType,omitempty"` // Type of device health
	Reason     string `json:"reason,omitempty"`     // Reason for the health score value
	Score      *int   `json:"score,omitempty"`      // health score of client device in the range of 1 to 10. Value 0 for a client represents an IDLE client
}
type ResponseClientsGetClientDetailV1DetailConnectedDevice struct {
	Type      string `json:"type,omitempty"`       // Type of device (AP or SWITCH)
	Name      string `json:"name,omitempty"`       // Name of the device
	Mac       string `json:"mac,omitempty"`        // MAC address of the access point
	ID        string `json:"id,omitempty"`         // Unique identifier of the device
	IPaddress string `json:"ip address,omitempty"` // Management IP address of the connected device.  (deprecated soon in favor of 'mgmtIp')
	MgmtIP    string `json:"mgmtIp,omitempty"`     // The IP address of the connected device
	Band      string `json:"band,omitempty"`       // Band of the AP
	Mode      string `json:"mode,omitempty"`       // The mode of the access point
}
type ResponseClientsGetClientDetailV1DetailOnboarding struct {
	AverageRunDuration   string   `json:"averageRunDuration,omitempty"`   // Average run Duration for a client
	MaxRunDuration       string   `json:"maxRunDuration,omitempty"`       // Max run duration for a client
	AverageAssocDuration string   `json:"averageAssocDuration,omitempty"` // Average association duration for a client
	MaxAssocDuration     string   `json:"maxAssocDuration,omitempty"`     // Max association duration for a client
	AverageAuthDuration  string   `json:"averageAuthDuration,omitempty"`  // Average auth duration for a client
	MaxAuthDuration      string   `json:"maxAuthDuration,omitempty"`      // Max auth duration for a client
	AverageDhcpDuration  string   `json:"averageDhcpDuration,omitempty"`  // Average DHCP duration for a client
	MaxDhcpDuration      string   `json:"maxDhcpDuration,omitempty"`      // Max DHCP duration for a client
	AAAServerIP          string   `json:"aaaServerIp,omitempty"`          // AAA server IP for a client
	DhcpServerIP         string   `json:"dhcpServerIp,omitempty"`         // DHCP server IP for a client
	AuthDoneTime         *int     `json:"authDoneTime,omitempty"`         // Epoch/Unix time in milliseconds
	AssocDoneTime        *int     `json:"assocDoneTime,omitempty"`        // Epoch/Unix time in milliseconds
	DhcpDoneTime         *int     `json:"dhcpDoneTime,omitempty"`         // Epoch/Unix time in milliseconds
	AssocRootcauseList   []string `json:"assocRootcauseList,omitempty"`   // Root cause list of ASSOC failure category
	AAARootcauseList     []string `json:"aaaRootcauseList,omitempty"`     // Root cause list of AAA failure category
	DhcpRootcauseList    []string `json:"dhcpRootcauseList,omitempty"`    // Root cause list of DHCP failure category
	OtherRootcauseList   []string `json:"otherRootcauseList,omitempty"`   // Root cause list of other failure category
	LatestRootCauseList  []string `json:"latestRootCauseList,omitempty"`  // Root cause list of latest root cause category
}
type ResponseClientsGetClientDetailV1ConnectionInfo struct {
	HostType      string `json:"hostType,omitempty"`      // Host Type - WIRELESS or WIRED
	NwDeviceName  string `json:"nwDeviceName,omitempty"`  // Name of the network device it is connected to. In case of wireless, it would be an AccessPoint
	NwDeviceMac   string `json:"nwDeviceMac,omitempty"`   // Device MAC address
	Protocol      string `json:"protocol,omitempty"`      // Connection Protocol used. This information is present for wireless hosts only
	Band          string `json:"band,omitempty"`          // The band at which the host is connected. This information is present for wireless hosts only
	SpatialStream string `json:"spatialStream,omitempty"` // The spatial stream of host. This information is present for wireless hosts only
	Channel       string `json:"channel,omitempty"`       // The channel used by the host. This information is present for wireless hosts only
	ChannelWidth  string `json:"channelWidth,omitempty"`  // The channel width used by the host. This information is present for wireless hosts only
	Wmm           string `json:"wmm,omitempty"`           // The wmm of the host. This information is present for wireless hosts only
	Uapsd         string `json:"uapsd,omitempty"`         // The UAPSD of the host. This information is present for wireless hosts only
	Timestamp     *int   `json:"timestamp,omitempty"`     // Epoch/Unix time in milliseconds
}
type ResponseClientsGetClientDetailV1Topology struct {
	Nodes *[]ResponseClientsGetClientDetailV1TopologyNodes `json:"nodes,omitempty"` //
	Links *[]ResponseClientsGetClientDetailV1TopologyLinks `json:"links,omitempty"` //
}
type ResponseClientsGetClientDetailV1TopologyNodes struct {
	Role            string   `json:"role,omitempty"`            // Device role
	Name            string   `json:"name,omitempty"`            // Device name
	ID              string   `json:"id,omitempty"`              // User ID, hostname, IP address, or MAC address
	Description     string   `json:"description,omitempty"`     // Description of the topology node
	DeviceType      string   `json:"deviceType,omitempty"`      // Device type
	PlatformID      string   `json:"platformId,omitempty"`      // Device platform ID
	Family          string   `json:"family,omitempty"`          // Device family
	IP              string   `json:"ip,omitempty"`              // Device IP
	IPv6            []string `json:"ipv6,omitempty"`            // Device IPv6
	SoftwareVersion string   `json:"softwareVersion,omitempty"` // Device software version
	UserID          string   `json:"userId,omitempty"`          // User ID
	NodeType        string   `json:"nodeType,omitempty"`        // Node type
	RadioFrequency  string   `json:"radioFrequency,omitempty"`  // Radio frequency
	Clients         *float64 `json:"clients,omitempty"`         // Number of clients
	Count           *int     `json:"count,omitempty"`           // Count
	HealthScore     *float64 `json:"healthScore,omitempty"`     // Device health score
	Level           *float64 `json:"level,omitempty"`           // Level in the topology
	FabricGroup     string   `json:"fabricGroup,omitempty"`     // Fabric Group
	FabricRole      []string `json:"fabricRole,omitempty"`      // Fabric Role
	ConnectedDevice string   `json:"connectedDevice,omitempty"` // Connected Device
	StackType       string   `json:"stackType,omitempty"`       // Stack Type
}
type ResponseClientsGetClientDetailV1TopologyLinks struct {
	Source              string                                                           `json:"source,omitempty"`              // Edge line starting node
	LinkStatus          string                                                           `json:"linkStatus,omitempty"`          // Link status of the link
	SourceLinkStatus    string                                                           `json:"sourceLinkStatus,omitempty"`    // The status of the link
	TargetLinkStatus    string                                                           `json:"targetLinkStatus,omitempty"`    // The status of the link
	Label               []string                                                         `json:"label,omitempty"`               // The details of the edge
	Target              string                                                           `json:"target,omitempty"`              // End node of the edge line
	ID                  string                                                           `json:"id,omitempty"`                  // Identifier of the node
	PortUtilization     *float64                                                         `json:"portUtilization,omitempty"`     // Port utilization
	SourceInterfaceName string                                                           `json:"sourceInterfaceName,omitempty"` // The interface name of the source
	TargetInterfaceName string                                                           `json:"targetInterfaceName,omitempty"` // The interface name of the target
	SourceDuplexInfo    string                                                           `json:"sourceDuplexInfo,omitempty"`    // The duplex info of the source
	TargetDuplexInfo    string                                                           `json:"targetDuplexInfo,omitempty"`    // The duplex info of the target
	SourcePortMode      string                                                           `json:"sourcePortMode,omitempty"`      // The port mode of the source
	TargetPortMode      string                                                           `json:"targetPortMode,omitempty"`      // The port mode of the target
	SourceAdminStatus   string                                                           `json:"sourceAdminStatus,omitempty"`   // The admin status of the source
	TargetAdminStatus   string                                                           `json:"targetAdminStatus,omitempty"`   // The admin status of the target
	ApRadioAdminStatus  string                                                           `json:"apRadioAdminStatus,omitempty"`  // The admin status of the radio
	ApRadioOperStatus   string                                                           `json:"apRadioOperStatus,omitempty"`   // The oper status of the radio
	SourcePortVLANInfo  string                                                           `json:"sourcePortVLANInfo,omitempty"`  // List of VLANs configured on the source port
	TargetPortVLANInfo  string                                                           `json:"targetPortVLANInfo,omitempty"`  // List of VLANs configured on the target port
	InterfaceDetails    *[]ResponseClientsGetClientDetailV1TopologyLinksInterfaceDetails `json:"interfaceDetails,omitempty"`    //
}
type ResponseClientsGetClientDetailV1TopologyLinksInterfaceDetails struct {
	ClientMacAddress       string `json:"clientMacAddress,omitempty"`       // The MAC address of the client device
	ConnectedDeviceIntName string `json:"connectedDeviceIntName,omitempty"` // The interface name of the network device
	Duplex                 string `json:"duplex,omitempty"`                 // The duplex info of the network device interface
	PortMode               string `json:"portMode,omitempty"`               // The port mode info of network device interface
	AdminStatus            string `json:"adminStatus,omitempty"`            // The admin status of network device interface
}
type ResponseClientsGetClientEnrichmentDetailsV1 []ResponseItemClientsGetClientEnrichmentDetailsV1 // Array of ResponseClientsGetClientEnrichmentDetailsV1
type ResponseItemClientsGetClientEnrichmentDetailsV1 struct {
	UserDetails     *ResponseItemClientsGetClientEnrichmentDetailsV1UserDetails       `json:"userDetails,omitempty"`     //
	ConnectedDevice *[]ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDevice `json:"connectedDevice,omitempty"` //
	IssueDetails    *ResponseItemClientsGetClientEnrichmentDetailsV1IssueDetails      `json:"issueDetails,omitempty"`    //
}
type ResponseItemClientsGetClientEnrichmentDetailsV1UserDetails struct {
	ID               string                                                                       `json:"id,omitempty"`               // Id
	ConnectionStatus string                                                                       `json:"connectionStatus,omitempty"` // Connection Status
	HostType         string                                                                       `json:"hostType,omitempty"`         // Host Type
	UserID           string                                                                       `json:"userId,omitempty"`           // User Id
	HostName         *ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsHostName          `json:"hostName,omitempty"`         // Host Name
	HostOs           *ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsHostOs            `json:"hostOs,omitempty"`           // Host Os
	HostVersion      *ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsHostVersion       `json:"hostVersion,omitempty"`      // Host Version
	SubType          *ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsSubType           `json:"subType,omitempty"`          // Sub Type
	LastUpdated      *int                                                                         `json:"lastUpdated,omitempty"`      // Last Updated
	HealthScore      *[]ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsHealthScore     `json:"healthScore,omitempty"`      //
	HostMac          string                                                                       `json:"hostMac,omitempty"`          // Host Mac
	HostIPV4         string                                                                       `json:"hostIpV4,omitempty"`         // Host Ip V4
	HostIPV6         *[]ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsHostIPV6        `json:"hostIpV6,omitempty"`         // Host Ip V6
	AuthType         *ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsAuthType          `json:"authType,omitempty"`         // Auth Type
	VLANID           string                                                                       `json:"vlanId,omitempty"`           // Vlan Id
	SSID             *ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsSSID              `json:"ssid,omitempty"`             // Ssid
	Location         *ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsLocation          `json:"location,omitempty"`         // Location
	ClientConnection string                                                                       `json:"clientConnection,omitempty"` // Client Connection
	ConnectedDevice  *[]ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsConnectedDevice `json:"connectedDevice,omitempty"`  // Connected Device
	IssueCount       *float64                                                                     `json:"issueCount,omitempty"`       // Issue Count
	Rssi             *ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsRssi              `json:"rssi,omitempty"`             // Rssi
	Snr              *ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsSnr               `json:"snr,omitempty"`              // Snr
	DataRate         *ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsDataRate          `json:"dataRate,omitempty"`         // Data Rate
	Port             *ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsPort              `json:"port,omitempty"`             // Port
}
type ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsHostName interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsHostOs interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsHostVersion interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsSubType interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsHealthScore struct {
	HealthType string `json:"healthType,omitempty"` // Health Type
	Reason     string `json:"reason,omitempty"`     // Reason
	Score      *int   `json:"score,omitempty"`      // Score
}
type ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsHostIPV6 interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsAuthType interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsSSID interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsLocation interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsConnectedDevice interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsRssi interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsSnr interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsDataRate interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsPort interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDevice struct {
	DeviceDetails *ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetails `json:"deviceDetails,omitempty"` //
}
type ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetails struct {
	Family                    string                                                                                         `json:"family,omitempty"`                    // Family
	Type                      string                                                                                         `json:"type,omitempty"`                      // Type
	Location                  *ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsLocation           `json:"location,omitempty"`                  // Location
	ErrorCode                 string                                                                                         `json:"errorCode,omitempty"`                 // Error Code
	MacAddress                string                                                                                         `json:"macAddress,omitempty"`                // Mac Address
	Role                      string                                                                                         `json:"role,omitempty"`                      // Role
	ApManagerInterfaceIP      string                                                                                         `json:"apManagerInterfaceIp,omitempty"`      // Ap Manager Interface Ip
	AssociatedWlcIP           string                                                                                         `json:"associatedWlcIp,omitempty"`           // Associated Wlc Ip
	BootDateTime              *ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsBootDateTime       `json:"bootDateTime,omitempty"`              // Boot Date Time
	CollectionStatus          string                                                                                         `json:"collectionStatus,omitempty"`          // Collection Status
	InterfaceCount            *ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsInterfaceCount     `json:"interfaceCount,omitempty"`            // Interface Count
	LineCardCount             *ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsLineCardCount      `json:"lineCardCount,omitempty"`             // Line Card Count
	LineCardID                *ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsLineCardID         `json:"lineCardId,omitempty"`                // Line Card Id
	ManagementIPAddress       string                                                                                         `json:"managementIpAddress,omitempty"`       // Management Ip Address
	MemorySize                string                                                                                         `json:"memorySize,omitempty"`                // Memory Size
	PlatformID                string                                                                                         `json:"platformId,omitempty"`                // Platform Id
	ReachabilityFailureReason string                                                                                         `json:"reachabilityFailureReason,omitempty"` // Reachability Failure Reason
	ReachabilityStatus        string                                                                                         `json:"reachabilityStatus,omitempty"`        // Reachability Status
	SNMPContact               string                                                                                         `json:"snmpContact,omitempty"`               // Snmp Contact
	SNMPLocation              string                                                                                         `json:"snmpLocation,omitempty"`              // Snmp Location
	TunnelUDPPort             string                                                                                         `json:"tunnelUdpPort,omitempty"`             // Tunnel Udp Port
	WaasDeviceMode            *ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsWaasDeviceMode     `json:"waasDeviceMode,omitempty"`            // Waas Device Mode
	Series                    string                                                                                         `json:"series,omitempty"`                    // Series
	InventoryStatusDetail     string                                                                                         `json:"inventoryStatusDetail,omitempty"`     // Inventory Status Detail
	CollectionInterval        string                                                                                         `json:"collectionInterval,omitempty"`        // Collection Interval
	SerialNumber              string                                                                                         `json:"serialNumber,omitempty"`              // Serial Number
	SoftwareVersion           string                                                                                         `json:"softwareVersion,omitempty"`           // Software Version
	RoleSource                string                                                                                         `json:"roleSource,omitempty"`                // Role Source
	Hostname                  string                                                                                         `json:"hostname,omitempty"`                  // Hostname
	UpTime                    string                                                                                         `json:"upTime,omitempty"`                    // Up Time
	LastUpdateTime            *int                                                                                           `json:"lastUpdateTime,omitempty"`            // Last Update Time
	ErrorDescription          *ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsErrorDescription   `json:"errorDescription,omitempty"`          // Error Description
	LocationName              *ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsLocationName       `json:"locationName,omitempty"`              // Location Name
	TagCount                  string                                                                                         `json:"tagCount,omitempty"`                  // Tag Count
	LastUpdated               string                                                                                         `json:"lastUpdated,omitempty"`               // Last Updated
	InstanceUUID              string                                                                                         `json:"instanceUuid,omitempty"`              // Instance Uuid
	ID                        string                                                                                         `json:"id,omitempty"`                        // Id
	NeighborTopology          *[]ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopology `json:"neighborTopology,omitempty"`          //
	Cisco360View              string                                                                                         `json:"cisco360view,omitempty"`              // Cisco360view
}
type ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsLocation interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsBootDateTime interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsInterfaceCount interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsLineCardCount interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsLineCardID interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsWaasDeviceMode interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsErrorDescription interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsLocationName interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopology struct {
	Nodes *[]ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyNodes `json:"nodes,omitempty"` //
	Links *[]ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyLinks `json:"links,omitempty"` //
}
type ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyNodes struct {
	Role            string                                                                                                           `json:"role,omitempty"`            // Role
	Name            string                                                                                                           `json:"name,omitempty"`            // Name
	ID              string                                                                                                           `json:"id,omitempty"`              // Id
	Description     string                                                                                                           `json:"description,omitempty"`     // Description
	DeviceType      *ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyNodesDeviceType      `json:"deviceType,omitempty"`      // Device Type
	PlatformID      *ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyNodesPlatformID      `json:"platformId,omitempty"`      // Platform Id
	Family          *ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyNodesFamily          `json:"family,omitempty"`          // Family
	IP              *ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyNodesIP              `json:"ip,omitempty"`              // Ip
	SoftwareVersion *ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyNodesSoftwareVersion `json:"softwareVersion,omitempty"` // Software Version
	UserID          *ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyNodesUserID          `json:"userId,omitempty"`          // User Id
	NodeType        *ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyNodesNodeType        `json:"nodeType,omitempty"`        // Node Type
	RadioFrequency  *ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyNodesRadioFrequency  `json:"radioFrequency,omitempty"`  // Radio Frequency
	Clients         *float64                                                                                                         `json:"clients,omitempty"`         // Clients
	Count           *ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyNodesCount           `json:"count,omitempty"`           // Count
	HealthScore     *ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyNodesHealthScore     `json:"healthScore,omitempty"`     // Health Score
	Level           *float64                                                                                                         `json:"level,omitempty"`           // Level
	FabricGroup     *ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyNodesFabricGroup     `json:"fabricGroup,omitempty"`     // Fabric Group
}
type ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyNodesDeviceType interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyNodesPlatformID interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyNodesFamily interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyNodesIP interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyNodesSoftwareVersion interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyNodesUserID interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyNodesNodeType interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyNodesRadioFrequency interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyNodesCount interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyNodesHealthScore interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyNodesFabricGroup interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyLinks struct {
	Source          string                                                                                                           `json:"source,omitempty"`          // Source
	LinkStatus      string                                                                                                           `json:"linkStatus,omitempty"`      // Link Status
	Label           *[]ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyLinksLabel         `json:"label,omitempty"`           // Label
	Target          string                                                                                                           `json:"target,omitempty"`          // Target
	ID              *ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyLinksID              `json:"id,omitempty"`              // Id
	PortUtilization *ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyLinksPortUtilization `json:"portUtilization,omitempty"` // Port Utilization
}
type ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyLinksLabel interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyLinksID interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyLinksPortUtilization interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1IssueDetails struct {
	Issue *[]ResponseItemClientsGetClientEnrichmentDetailsV1IssueDetailsIssue `json:"issue,omitempty"` //
}
type ResponseItemClientsGetClientEnrichmentDetailsV1IssueDetailsIssue struct {
	IssueID          string                                                                              `json:"issueId,omitempty"`          // Issue Id
	IssueSource      string                                                                              `json:"issueSource,omitempty"`      // Issue Source
	IssueCategory    string                                                                              `json:"issueCategory,omitempty"`    // Issue Category
	IssueName        string                                                                              `json:"issueName,omitempty"`        // Issue Name
	IssueDescription string                                                                              `json:"issueDescription,omitempty"` // Issue Description
	IssueEntity      string                                                                              `json:"issueEntity,omitempty"`      // Issue Entity
	IssueEntityValue string                                                                              `json:"issueEntityValue,omitempty"` // Issue Entity Value
	IssueSeverity    string                                                                              `json:"issueSeverity,omitempty"`    // Issue Severity
	IssuePriority    string                                                                              `json:"issuePriority,omitempty"`    // Issue Priority
	IssueSummary     string                                                                              `json:"issueSummary,omitempty"`     // Issue Summary
	IssueTimestamp   *int                                                                                `json:"issueTimestamp,omitempty"`   // Issue Timestamp
	SuggestedActions *[]ResponseItemClientsGetClientEnrichmentDetailsV1IssueDetailsIssueSuggestedActions `json:"suggestedActions,omitempty"` //
	ImpactedHosts    *[]ResponseItemClientsGetClientEnrichmentDetailsV1IssueDetailsIssueImpactedHosts    `json:"impactedHosts,omitempty"`    //
}
type ResponseItemClientsGetClientEnrichmentDetailsV1IssueDetailsIssueSuggestedActions struct {
	Message string                                                                                   `json:"message,omitempty"` // Message
	Steps   *[]ResponseItemClientsGetClientEnrichmentDetailsV1IssueDetailsIssueSuggestedActionsSteps `json:"steps,omitempty"`   // Steps
}
type ResponseItemClientsGetClientEnrichmentDetailsV1IssueDetailsIssueSuggestedActionsSteps interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1IssueDetailsIssueImpactedHosts struct {
	HostType           string                                                                                 `json:"hostType,omitempty"`           // Host Type
	HostName           string                                                                                 `json:"hostName,omitempty"`           // Host Name
	HostOs             string                                                                                 `json:"hostOs,omitempty"`             // Host Os
	SSID               string                                                                                 `json:"ssid,omitempty"`               // Ssid
	ConnectedInterface string                                                                                 `json:"connectedInterface,omitempty"` // Connected Interface
	MacAddress         string                                                                                 `json:"macAddress,omitempty"`         // Mac Address
	FailedAttempts     *int                                                                                   `json:"failedAttempts,omitempty"`     // Failed Attempts
	Location           *ResponseItemClientsGetClientEnrichmentDetailsV1IssueDetailsIssueImpactedHostsLocation `json:"location,omitempty"`           //
	Timestamp          *int                                                                                   `json:"timestamp,omitempty"`          // Timestamp
}
type ResponseItemClientsGetClientEnrichmentDetailsV1IssueDetailsIssueImpactedHostsLocation struct {
	SiteID      string                                                                                              `json:"siteId,omitempty"`      // Site Id
	SiteType    string                                                                                              `json:"siteType,omitempty"`    // Site Type
	Area        string                                                                                              `json:"area,omitempty"`        // Area
	Building    string                                                                                              `json:"building,omitempty"`    // Building
	Floor       *ResponseItemClientsGetClientEnrichmentDetailsV1IssueDetailsIssueImpactedHostsLocationFloor         `json:"floor,omitempty"`       // Floor
	ApsImpacted *[]ResponseItemClientsGetClientEnrichmentDetailsV1IssueDetailsIssueImpactedHostsLocationApsImpacted `json:"apsImpacted,omitempty"` // Aps Impacted
}
type ResponseItemClientsGetClientEnrichmentDetailsV1IssueDetailsIssueImpactedHostsLocationFloor interface{}
type ResponseItemClientsGetClientEnrichmentDetailsV1IssueDetailsIssueImpactedHostsLocationApsImpacted interface{}
type ResponseClientsGetOverallClientHealthV1 struct {
	Version  string                                             `json:"version,omitempty"`  // Response output version
	Response *[]ResponseClientsGetOverallClientHealthV1Response `json:"response,omitempty"` //
}
type ResponseClientsGetOverallClientHealthV1Response struct {
	SiteID      string                                                        `json:"siteId,omitempty"`      // Site UUID or 'global'
	ScoreDetail *[]ResponseClientsGetOverallClientHealthV1ResponseScoreDetail `json:"scoreDetail,omitempty"` //
}
type ResponseClientsGetOverallClientHealthV1ResponseScoreDetail struct {
	ScoreCategory                  *ResponseClientsGetOverallClientHealthV1ResponseScoreDetailScoreCategory `json:"scoreCategory,omitempty"`                  //
	ScoreValue                     *float64                                                                 `json:"scoreValue,omitempty"`                     // Percentage of GOOD health score in the category.  (-1 means not applicable for the category)
	ClientCount                    *int                                                                     `json:"clientCount,omitempty"`                    // Total client count
	ClientUniqueCount              *int                                                                     `json:"clientUniqueCount,omitempty"`              // Total unique client count
	MaintenanceAffectedClientCount *int                                                                     `json:"maintenanceAffectedClientCount,omitempty"` // Total client count affected by maintenance
	RandomMacCount                 *int                                                                     `json:"randomMacCount,omitempty"`                 // Total client count with random MAC count
	DuidCount                      *int                                                                     `json:"duidCount,omitempty"`                      // Device UUID count
	Starttime                      *int                                                                     `json:"starttime,omitempty"`                      // UTC timestamp of data start time
	Endtime                        *int                                                                     `json:"endtime,omitempty"`                        // UTC timestamp of data end time
	ConnectedToUdnCount            *int                                                                     `json:"connectedToUdnCount,omitempty"`            // Total connected to UDN count
	UnconnectedToUdnCount          *int                                                                     `json:"unconnectedToUdnCount,omitempty"`          // Total unconnected to UDN count
	ScoreList                      *[]ResponseClientsGetOverallClientHealthV1ResponseScoreDetailScoreList   `json:"scoreList,omitempty"`                      //
}
type ResponseClientsGetOverallClientHealthV1ResponseScoreDetailScoreCategory struct {
	ScoreCategory string `json:"scoreCategory,omitempty"` // Health score category
	Value         string `json:"value,omitempty"`         // Health score category value
}
type ResponseClientsGetOverallClientHealthV1ResponseScoreDetailScoreList struct {
	ScoreCategory                  *ResponseClientsGetOverallClientHealthV1ResponseScoreDetailScoreListScoreCategory `json:"scoreCategory,omitempty"`                  //
	ScoreValue                     *float64                                                                          `json:"scoreValue,omitempty"`                     // Percentage of GOOD health score in the category.  (-1 means not applicable for the category)
	ClientCount                    *int                                                                              `json:"clientCount,omitempty"`                    // Total client count
	ClientUniqueCount              *int                                                                              `json:"clientUniqueCount,omitempty"`              // Total unique client count
	MaintenanceAffectedClientCount *int                                                                              `json:"maintenanceAffectedClientCount,omitempty"` // Total client count affected by maintenance
	RandomMacCount                 *int                                                                              `json:"randomMacCount,omitempty"`                 // Total client count with random MAC count
	DuidCount                      *int                                                                              `json:"duidCount,omitempty"`                      // Device UUID count
	Starttime                      *int                                                                              `json:"starttime,omitempty"`                      // UTC timestamp of data start time
	Endtime                        *int                                                                              `json:"endtime,omitempty"`                        // UTC timestamp of data end time
	ConnectedToUdnCount            *int                                                                              `json:"connectedToUdnCount,omitempty"`            // Total connected to UDN count
	UnconnectedToUdnCount          *int                                                                              `json:"unconnectedToUdnCount,omitempty"`          // Total unconnected to UDN count
}
type ResponseClientsGetOverallClientHealthV1ResponseScoreDetailScoreListScoreCategory struct {
	ScoreCategory string `json:"scoreCategory,omitempty"` // Category of the overall health score
	Value         string `json:"value,omitempty"`         // Health score category value
}
type ResponseClientsClientProximityV1 struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1 struct {
	StartTime           *int                                                                                                                          `json:"startTime,omitempty"`           // Start Time
	EndTime             *int                                                                                                                          `json:"endTime,omitempty"`             // End Time
	Views               []string                                                                                                                      `json:"views,omitempty"`               // Views
	Attributes          []string                                                                                                                      `json:"attributes,omitempty"`          // Attributes
	Filters             *[]RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1Filters             `json:"filters,omitempty"`             //
	AggregateAttributes *[]RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1AggregateAttributes `json:"aggregateAttributes,omitempty"` //
	Page                *RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1Page                  `json:"page,omitempty"`                //
}
type RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1Filters struct {
	Key      string `json:"key,omitempty"`      // Key
	Operator string `json:"operator,omitempty"` // Operator
	Value    *int   `json:"value,omitempty"`    // Value
}
type RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1AggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
}
type RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1Page struct {
	Limit  *int                                                                                                                 `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                                                 `json:"offset,omitempty"` // Offset
	SortBy *[]RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1PageSortBy `json:"sortBy,omitempty"` //
}
type RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1PageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type RequestClientsRetrievesTheNumberOfClientsByApplyingComplexFiltersV1 struct {
	StartTime *int                                                                          `json:"startTime,omitempty"` // Start Time
	EndTime   *int                                                                          `json:"endTime,omitempty"`   // End Time
	Filters   *[]RequestClientsRetrievesTheNumberOfClientsByApplyingComplexFiltersV1Filters `json:"filters,omitempty"`   //
}
type RequestClientsRetrievesTheNumberOfClientsByApplyingComplexFiltersV1Filters struct {
	Key      string `json:"key,omitempty"`      // Key
	Operator string `json:"operator,omitempty"` // Operator
	Value    *int   `json:"value,omitempty"`    // Value
}
type RequestClientsRetrievesSummaryAnalyticsDataRelatedToClientsV1 struct {
	StartTime           *int                                                                                `json:"startTime,omitempty"`           // Start Time
	EndTime             *int                                                                                `json:"endTime,omitempty"`             // End Time
	GroupBy             []string                                                                            `json:"groupBy,omitempty"`             // Group By
	Attributes          []string                                                                            `json:"attributes,omitempty"`          // Attributes
	Filters             *[]RequestClientsRetrievesSummaryAnalyticsDataRelatedToClientsV1Filters             `json:"filters,omitempty"`             //
	AggregateAttributes *[]RequestClientsRetrievesSummaryAnalyticsDataRelatedToClientsV1AggregateAttributes `json:"aggregateAttributes,omitempty"` //
	Page                *RequestClientsRetrievesSummaryAnalyticsDataRelatedToClientsV1Page                  `json:"page,omitempty"`                //
}
type RequestClientsRetrievesSummaryAnalyticsDataRelatedToClientsV1Filters struct {
	Key      string `json:"key,omitempty"`      // Key
	Operator string `json:"operator,omitempty"` // Operator
	Value    *int   `json:"value,omitempty"`    // Value
}
type RequestClientsRetrievesSummaryAnalyticsDataRelatedToClientsV1AggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
}
type RequestClientsRetrievesSummaryAnalyticsDataRelatedToClientsV1Page struct {
	Limit  *int                                                                       `json:"limit,omitempty"`  // Limit
	Cursor string                                                                     `json:"cursor,omitempty"` // Cursor
	SortBy *[]RequestClientsRetrievesSummaryAnalyticsDataRelatedToClientsV1PageSortBy `json:"sortBy,omitempty"` //
}
type RequestClientsRetrievesSummaryAnalyticsDataRelatedToClientsV1PageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsV1 struct {
	StartTime           *int                                                                                `json:"startTime,omitempty"`           // Start Time
	EndTime             *int                                                                                `json:"endTime,omitempty"`             // End Time
	TopN                *int                                                                                `json:"topN,omitempty"`                // Top N
	GroupBy             []string                                                                            `json:"groupBy,omitempty"`             // Group By
	Attributes          []string                                                                            `json:"attributes,omitempty"`          // Attributes
	Filters             *[]RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsV1Filters             `json:"filters,omitempty"`             //
	AggregateAttributes *[]RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsV1AggregateAttributes `json:"aggregateAttributes,omitempty"` //
	Page                *RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsV1Page                  `json:"page,omitempty"`                //
}
type RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsV1Filters struct {
	Key      string `json:"key,omitempty"`      // Key
	Operator string `json:"operator,omitempty"` // Operator
	Value    *int   `json:"value,omitempty"`    // Value
}
type RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsV1AggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
}
type RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsV1Page struct {
	Limit  *int                                                                       `json:"limit,omitempty"`  // Limit
	Cursor string                                                                     `json:"cursor,omitempty"` // Cursor
	SortBy *[]RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsV1PageSortBy `json:"sortBy,omitempty"` //
}
type RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsV1PageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type RequestClientsRetrievesTheTrendAnalyticsDataRelatedToClientsV1 struct {
	StartTime           *int                                                                                 `json:"startTime,omitempty"`           // Start Time
	EndTime             *int                                                                                 `json:"endTime,omitempty"`             // End Time
	TrendInterval       string                                                                               `json:"trendInterval,omitempty"`       // Trend Interval
	GroupBy             []string                                                                             `json:"groupBy,omitempty"`             // Group By
	Attributes          []string                                                                             `json:"attributes,omitempty"`          // Attributes
	Filters             *[]RequestClientsRetrievesTheTrendAnalyticsDataRelatedToClientsV1Filters             `json:"filters,omitempty"`             //
	AggregateAttributes *[]RequestClientsRetrievesTheTrendAnalyticsDataRelatedToClientsV1AggregateAttributes `json:"aggregateAttributes,omitempty"` //
	Page                *RequestClientsRetrievesTheTrendAnalyticsDataRelatedToClientsV1Page                  `json:"page,omitempty"`                //
}
type RequestClientsRetrievesTheTrendAnalyticsDataRelatedToClientsV1Filters struct {
	Key      string `json:"key,omitempty"`      // Key
	Operator string `json:"operator,omitempty"` // Operator
	Value    *int   `json:"value,omitempty"`    // Value
}
type RequestClientsRetrievesTheTrendAnalyticsDataRelatedToClientsV1AggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
}
type RequestClientsRetrievesTheTrendAnalyticsDataRelatedToClientsV1Page struct {
	Limit         *int   `json:"limit,omitempty"`         // Limit
	Cursor        string `json:"cursor,omitempty"`        // Cursor
	TimeSortOrder string `json:"timeSortOrder,omitempty"` // Time Sort Order
}
type RequestClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1 struct {
	StartTime           *int                                                                                               `json:"startTime,omitempty"`           // Start Time
	EndTime             *int                                                                                               `json:"endTime,omitempty"`             // End Time
	TrendInterval       string                                                                                             `json:"trendInterval,omitempty"`       // Trend Interval
	GroupBy             []string                                                                                           `json:"groupBy,omitempty"`             // Group By
	Attributes          []string                                                                                           `json:"attributes,omitempty"`          // Attributes
	Filters             *[]RequestClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1Filters             `json:"filters,omitempty"`             //
	AggregateAttributes *[]RequestClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1AggregateAttributes `json:"aggregateAttributes,omitempty"` //
	Page                *RequestClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1Page                  `json:"page,omitempty"`                //
}
type RequestClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1Filters struct {
	Key      string `json:"key,omitempty"`      // Key
	Operator string `json:"operator,omitempty"` // Operator
	Value    *int   `json:"value,omitempty"`    // Value
}
type RequestClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1AggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
}
type RequestClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1Page struct {
	Limit         *int   `json:"limit,omitempty"`         // Limit
	Cursor        string `json:"cursor,omitempty"`        // Cursor
	TimeSortOrder string `json:"timeSortOrder,omitempty"` // Time Sort Order
}

//RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1 Retrieves the list of clients, while also offering basic filtering and sorting capabilities. - ecb7-ab7e-47eb-8793
/* Retrieves the list of clients, while also offering basic filtering and sorting capabilities. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-clients1-1.0.0-resolved.yaml


@param RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1HeaderParams Custom header parameters
@param RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-list-of-clients-while-also-offering-basic-filtering-and-sorting-capabilities-v1
*/
func (s *ClientsService) RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1(RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1HeaderParams *RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1HeaderParams, RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1QueryParams *RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1QueryParams) (*ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1, *resty.Response, error) {
	path := "/dna/data/api/v1/clients"

	queryString, _ := query.Values(RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1HeaderParams != nil {

		if RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1(RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1HeaderParams, RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1")
	}

	result := response.Result().(*ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1)
	return result, response, err

}

//RetrievesTheTotalCountOfClientsByApplyingBasicFilteringV1 Retrieves the total count of clients by applying basic filtering - a486-4bef-4cab-9548
/* Retrieves the number of clients by applying basic filtering. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-clients1-1.0.0-resolved.yaml


@param RetrievesTheTotalCountOfClientsByApplyingBasicFilteringV1HeaderParams Custom header parameters
@param RetrievesTheTotalCountOfClientsByApplyingBasicFilteringV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-total-count-of-clients-by-applying-basic-filtering-v1
*/
func (s *ClientsService) RetrievesTheTotalCountOfClientsByApplyingBasicFilteringV1(RetrievesTheTotalCountOfClientsByApplyingBasicFilteringV1HeaderParams *RetrievesTheTotalCountOfClientsByApplyingBasicFilteringV1HeaderParams, RetrievesTheTotalCountOfClientsByApplyingBasicFilteringV1QueryParams *RetrievesTheTotalCountOfClientsByApplyingBasicFilteringV1QueryParams) (*ResponseClientsRetrievesTheTotalCountOfClientsByApplyingBasicFilteringV1, *resty.Response, error) {
	path := "/dna/data/api/v1/clients/count"

	queryString, _ := query.Values(RetrievesTheTotalCountOfClientsByApplyingBasicFilteringV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheTotalCountOfClientsByApplyingBasicFilteringV1HeaderParams != nil {

		if RetrievesTheTotalCountOfClientsByApplyingBasicFilteringV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheTotalCountOfClientsByApplyingBasicFilteringV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseClientsRetrievesTheTotalCountOfClientsByApplyingBasicFilteringV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheTotalCountOfClientsByApplyingBasicFilteringV1(RetrievesTheTotalCountOfClientsByApplyingBasicFilteringV1HeaderParams, RetrievesTheTotalCountOfClientsByApplyingBasicFilteringV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheTotalCountOfClientsByApplyingBasicFilteringV1")
	}

	result := response.Result().(*ResponseClientsRetrievesTheTotalCountOfClientsByApplyingBasicFilteringV1)
	return result, response, err

}

//RetrievesSpecificClientInformationMatchingTheMacaddressV1 Retrieves specific client information matching the MAC address. - 829f-8b65-4779-9069
/* Retrieves specific client information matching the MAC address. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-clients1-1.0.0-resolved.yaml


@param id id path parameter. id is the client mac address. It can be specified is any notational conventions  01:23:45:67:89:AB or 01-23-45-67-89-AB or 0123.4567.89AB and is case insensitive

@param RetrievesSpecificClientInformationMatchingTheMACAddressV1HeaderParams Custom header parameters
@param RetrievesSpecificClientInformationMatchingTheMACAddressV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-specific-client-information-matching-the-macaddress-v1
*/
func (s *ClientsService) RetrievesSpecificClientInformationMatchingTheMacaddressV1(id string, RetrievesSpecificClientInformationMatchingTheMACAddressV1HeaderParams *RetrievesSpecificClientInformationMatchingTheMacaddressV1HeaderParams, RetrievesSpecificClientInformationMatchingTheMACAddressV1QueryParams *RetrievesSpecificClientInformationMatchingTheMacaddressV1QueryParams) (*ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressV1, *resty.Response, error) {
	path := "/dna/data/api/v1/clients/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(RetrievesSpecificClientInformationMatchingTheMACAddressV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesSpecificClientInformationMatchingTheMACAddressV1HeaderParams != nil {

		if RetrievesSpecificClientInformationMatchingTheMACAddressV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesSpecificClientInformationMatchingTheMACAddressV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesSpecificClientInformationMatchingTheMacaddressV1(id, RetrievesSpecificClientInformationMatchingTheMACAddressV1HeaderParams, RetrievesSpecificClientInformationMatchingTheMACAddressV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesSpecificClientInformationMatchingTheMacaddressV1")
	}

	result := response.Result().(*ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressV1)
	return result, response, err

}

//GetClientDetailV1 Get Client Detail - 1980-1996-4389-9d65
/* Returns detailed Client information retrieved by Mac Address for any given point of time.


@param GetClientDetailV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-client-detail-v1
*/
func (s *ClientsService) GetClientDetailV1(GetClientDetailV1QueryParams *GetClientDetailV1QueryParams) (*ResponseClientsGetClientDetailV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/client-detail"

	queryString, _ := query.Values(GetClientDetailV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseClientsGetClientDetailV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetClientDetailV1(GetClientDetailV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetClientDetailV1")
	}

	result := response.Result().(*ResponseClientsGetClientDetailV1)
	return result, response, err

}

//GetClientEnrichmentDetailsV1 Get Client Enrichment Details - b199-685d-4d08-9a67
/* Enriches a given network End User context (a network user-id or end user’s device Mac Address) with details about the user, the devices that the user is connected to and the assurance issues that the user is impacted by


@param GetClientEnrichmentDetailsV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-client-enrichment-details-v1
*/
func (s *ClientsService) GetClientEnrichmentDetailsV1(GetClientEnrichmentDetailsV1HeaderParams *GetClientEnrichmentDetailsV1HeaderParams) (*ResponseClientsGetClientEnrichmentDetailsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/client-enrichment-details"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetClientEnrichmentDetailsV1HeaderParams != nil {

		if GetClientEnrichmentDetailsV1HeaderParams.EntityType != "" {
			clientRequest = clientRequest.SetHeader("entity_type", GetClientEnrichmentDetailsV1HeaderParams.EntityType)
		}

		if GetClientEnrichmentDetailsV1HeaderParams.EntityValue != "" {
			clientRequest = clientRequest.SetHeader("entity_value", GetClientEnrichmentDetailsV1HeaderParams.EntityValue)
		}

		if GetClientEnrichmentDetailsV1HeaderParams.IssueCategory != "" {
			clientRequest = clientRequest.SetHeader("issueCategory", GetClientEnrichmentDetailsV1HeaderParams.IssueCategory)
		}

		if GetClientEnrichmentDetailsV1HeaderParams.Persistbapioutput != "" {
			clientRequest = clientRequest.SetHeader("__persistbapioutput", GetClientEnrichmentDetailsV1HeaderParams.Persistbapioutput)
		}

	}

	response, err = clientRequest.
		SetResult(&ResponseClientsGetClientEnrichmentDetailsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetClientEnrichmentDetailsV1(GetClientEnrichmentDetailsV1HeaderParams)
		}
		return nil, response, fmt.Errorf("error with operation GetClientEnrichmentDetailsV1")
	}

	result := response.Result().(*ResponseClientsGetClientEnrichmentDetailsV1)
	return result, response, err

}

//GetOverallClientHealthV1 Get Overall Client Health - 3f9f-d80e-4df9-863c
/* Returns Overall Client Health information by Client type (Wired and Wireless) for any given point of time


@param GetOverallClientHealthV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-overall-client-health-v1
*/
func (s *ClientsService) GetOverallClientHealthV1(GetOverallClientHealthV1QueryParams *GetOverallClientHealthV1QueryParams) (*ResponseClientsGetOverallClientHealthV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/client-health"

	queryString, _ := query.Values(GetOverallClientHealthV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseClientsGetOverallClientHealthV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetOverallClientHealthV1(GetOverallClientHealthV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetOverallClientHealthV1")
	}

	result := response.Result().(*ResponseClientsGetOverallClientHealthV1)
	return result, response, err

}

//ClientProximityV1 Client Proximity - 4497-ebe2-4c88-84a1
/* This intent API will provide client proximity information for a specific wireless user. Proximity is defined as presence on the same floor at the same time as the specified wireless user. The Proximity workflow requires the subscription to the following event (via the Event Notification workflow) prior to making this API call: NETWORK-CLIENTS-3-506 Client Proximity Report.


@param ClientProximityV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!client-proximity-v1
*/
func (s *ClientsService) ClientProximityV1(ClientProximityV1QueryParams *ClientProximityV1QueryParams) (*ResponseClientsClientProximityV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/client-proximity"

	queryString, _ := query.Values(ClientProximityV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseClientsClientProximityV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ClientProximityV1(ClientProximityV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ClientProximityV1")
	}

	result := response.Result().(*ResponseClientsClientProximityV1)
	return result, response, err

}

//RetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1 Retrieves the list of clients by applying complex filters while also supporting aggregate attributes. - e982-db36-48c9-a651
/* Retrieves the list of clients by applying complex filters while also supporting aggregate attributes. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-clients1-1.0.0-resolved.yaml


@param RetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-list-of-clients-by-applying-complex-filters-while-also-supporting-aggregate-attributes-v1
*/
func (s *ClientsService) RetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1(requestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1 *RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1, RetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1HeaderParams *RetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1HeaderParams) (*ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1, *resty.Response, error) {
	path := "/dna/data/api/v1/clients/query"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1HeaderParams != nil {

		if RetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1).
		SetResult(&ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1(requestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1, RetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation RetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1")
	}

	result := response.Result().(*ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1)
	return result, response, err

}

//RetrievesTheNumberOfClientsByApplyingComplexFiltersV1 Retrieves the number of clients by applying complex filters. - b596-a8d5-40ba-9c5c
/* Retrieves the number of clients by applying complex filters. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-clients1-1.0.0-resolved.yaml


@param RetrievesTheNumberOfClientsByApplyingComplexFiltersV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-number-of-clients-by-applying-complex-filters-v1
*/
func (s *ClientsService) RetrievesTheNumberOfClientsByApplyingComplexFiltersV1(requestClientsRetrievesTheNumberOfClientsByApplyingComplexFiltersV1 *RequestClientsRetrievesTheNumberOfClientsByApplyingComplexFiltersV1, RetrievesTheNumberOfClientsByApplyingComplexFiltersV1HeaderParams *RetrievesTheNumberOfClientsByApplyingComplexFiltersV1HeaderParams) (*ResponseClientsRetrievesTheNumberOfClientsByApplyingComplexFiltersV1, *resty.Response, error) {
	path := "/dna/data/api/v1/clients/query/count"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheNumberOfClientsByApplyingComplexFiltersV1HeaderParams != nil {

		if RetrievesTheNumberOfClientsByApplyingComplexFiltersV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheNumberOfClientsByApplyingComplexFiltersV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestClientsRetrievesTheNumberOfClientsByApplyingComplexFiltersV1).
		SetResult(&ResponseClientsRetrievesTheNumberOfClientsByApplyingComplexFiltersV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheNumberOfClientsByApplyingComplexFiltersV1(requestClientsRetrievesTheNumberOfClientsByApplyingComplexFiltersV1, RetrievesTheNumberOfClientsByApplyingComplexFiltersV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation RetrievesTheNumberOfClientsByApplyingComplexFiltersV1")
	}

	result := response.Result().(*ResponseClientsRetrievesTheNumberOfClientsByApplyingComplexFiltersV1)
	return result, response, err

}

//RetrievesSummaryAnalyticsDataRelatedToClientsV1 Retrieves summary analytics data related to clients. - d7a5-8ae6-4ecb-aab9
/* Retrieves summary analytics data related to clients while applying complex filtering, aggregate functions, and grouping. This API facilitates obtaining consolidated insights into the performance and status of the clients. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-clients1-1.0.0-resolved.yaml


@param RetrievesSummaryAnalyticsDataRelatedToClientsV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-summary-analytics-data-related-to-clients-v1
*/
func (s *ClientsService) RetrievesSummaryAnalyticsDataRelatedToClientsV1(requestClientsRetrievesSummaryAnalyticsDataRelatedToClientsV1 *RequestClientsRetrievesSummaryAnalyticsDataRelatedToClientsV1, RetrievesSummaryAnalyticsDataRelatedToClientsV1HeaderParams *RetrievesSummaryAnalyticsDataRelatedToClientsV1HeaderParams) (*ResponseClientsRetrievesSummaryAnalyticsDataRelatedToClientsV1, *resty.Response, error) {
	path := "/dna/data/api/v1/clients/summaryAnalytics"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesSummaryAnalyticsDataRelatedToClientsV1HeaderParams != nil {

		if RetrievesSummaryAnalyticsDataRelatedToClientsV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesSummaryAnalyticsDataRelatedToClientsV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestClientsRetrievesSummaryAnalyticsDataRelatedToClientsV1).
		SetResult(&ResponseClientsRetrievesSummaryAnalyticsDataRelatedToClientsV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesSummaryAnalyticsDataRelatedToClientsV1(requestClientsRetrievesSummaryAnalyticsDataRelatedToClientsV1, RetrievesSummaryAnalyticsDataRelatedToClientsV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation RetrievesSummaryAnalyticsDataRelatedToClientsV1")
	}

	result := response.Result().(*ResponseClientsRetrievesSummaryAnalyticsDataRelatedToClientsV1)
	return result, response, err

}

//RetrievesTheTopNAnalyticsDataRelatedToClientsV1 Retrieves the Top-N analytics data related to clients. - dd98-6b95-401a-90e7
/* Retrieves the top N analytics data related to clients based on the provided input data. This API facilitates obtaining insights into the top-performing or most impacted clients. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-clients1-1.0.0-resolved.yaml


@param RetrievesTheTopNAnalyticsDataRelatedToClientsV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-top-n-analytics-data-related-to-clients-v1
*/
func (s *ClientsService) RetrievesTheTopNAnalyticsDataRelatedToClientsV1(requestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsV1 *RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsV1, RetrievesTheTopNAnalyticsDataRelatedToClientsV1HeaderParams *RetrievesTheTopNAnalyticsDataRelatedToClientsV1HeaderParams) (*ResponseClientsRetrievesTheTopNAnalyticsDataRelatedToClientsV1, *resty.Response, error) {
	path := "/dna/data/api/v1/clients/topNAnalytics"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheTopNAnalyticsDataRelatedToClientsV1HeaderParams != nil {

		if RetrievesTheTopNAnalyticsDataRelatedToClientsV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheTopNAnalyticsDataRelatedToClientsV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsV1).
		SetResult(&ResponseClientsRetrievesTheTopNAnalyticsDataRelatedToClientsV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheTopNAnalyticsDataRelatedToClientsV1(requestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsV1, RetrievesTheTopNAnalyticsDataRelatedToClientsV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation RetrievesTheTopNAnalyticsDataRelatedToClientsV1")
	}

	result := response.Result().(*ResponseClientsRetrievesTheTopNAnalyticsDataRelatedToClientsV1)
	return result, response, err

}

//RetrievesTheTrendAnalyticsDataRelatedToClientsV1 Retrieves the Trend analytics data related to clients. - 5d8e-eb0c-4988-af1f
/* Retrieves the trend analytics of client data for the specified time range. The data will be grouped based on the given trend time interval. This API facilitates obtaining consolidated insights into the performance and status of the clients over the specified start and end time. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-clients1-1.0.0-resolved.yaml


@param RetrievesTheTrendAnalyticsDataRelatedToClientsV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-trend-analytics-data-related-to-clients-v1
*/
func (s *ClientsService) RetrievesTheTrendAnalyticsDataRelatedToClientsV1(requestClientsRetrievesTheTrendAnalyticsDataRelatedToClientsV1 *RequestClientsRetrievesTheTrendAnalyticsDataRelatedToClientsV1, RetrievesTheTrendAnalyticsDataRelatedToClientsV1HeaderParams *RetrievesTheTrendAnalyticsDataRelatedToClientsV1HeaderParams) (*ResponseClientsRetrievesTheTrendAnalyticsDataRelatedToClientsV1, *resty.Response, error) {
	path := "/dna/data/api/v1/clients/trendAnalytics"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheTrendAnalyticsDataRelatedToClientsV1HeaderParams != nil {

		if RetrievesTheTrendAnalyticsDataRelatedToClientsV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheTrendAnalyticsDataRelatedToClientsV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestClientsRetrievesTheTrendAnalyticsDataRelatedToClientsV1).
		SetResult(&ResponseClientsRetrievesTheTrendAnalyticsDataRelatedToClientsV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheTrendAnalyticsDataRelatedToClientsV1(requestClientsRetrievesTheTrendAnalyticsDataRelatedToClientsV1, RetrievesTheTrendAnalyticsDataRelatedToClientsV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation RetrievesTheTrendAnalyticsDataRelatedToClientsV1")
	}

	result := response.Result().(*ResponseClientsRetrievesTheTrendAnalyticsDataRelatedToClientsV1)
	return result, response, err

}

//RetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1 Retrieves specific client information over a specified period of time. - d086-9874-4ffa-a996
/* Retrieves the time series information of a specific client by applying complex filters, aggregate functions, and grouping. The data will be grouped based on the specified trend time interval. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-clients1-1.0.0-resolved.yaml


@param id id path parameter. id is the client mac address. It can be specified in one of the notational conventions  01:23:45:67:89:AB or 01-23-45-67-89-AB or 0123.4567.89AB and is case insensitive

@param RetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-specific-client-information-over-a-specified-period-of-time-v1
*/
func (s *ClientsService) RetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1(id string, requestClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1 *RequestClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1, RetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1HeaderParams *RetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1HeaderParams) (*ResponseClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1, *resty.Response, error) {
	path := "/dna/data/api/v1/clients/{id}/trendAnalytics"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1HeaderParams != nil {

		if RetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1).
		SetResult(&ResponseClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1(id, requestClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1, RetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation RetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1")
	}

	result := response.Result().(*ResponseClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1)
	return result, response, err

}

// Alias Function
func (s *ClientsService) GetClientDetail(GetClientDetailV1QueryParams *GetClientDetailV1QueryParams) (*ResponseClientsGetClientDetailV1, *resty.Response, error) {
	return s.GetClientDetailV1(GetClientDetailV1QueryParams)
}

// Alias Function
func (s *ClientsService) GetClientEnrichmentDetails(GetClientEnrichmentDetailsV1HeaderParams *GetClientEnrichmentDetailsV1HeaderParams) (*ResponseClientsGetClientEnrichmentDetailsV1, *resty.Response, error) {
	return s.GetClientEnrichmentDetailsV1(GetClientEnrichmentDetailsV1HeaderParams)
}

// Alias Function
func (s *ClientsService) RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilities(RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1HeaderParams *RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1HeaderParams, RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1QueryParams *RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1QueryParams) (*ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1, *resty.Response, error) {
	return s.RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1(RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1HeaderParams, RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesV1QueryParams)
}

// Alias Function
func (s *ClientsService) ClientProximity(ClientProximityV1QueryParams *ClientProximityV1QueryParams) (*ResponseClientsClientProximityV1, *resty.Response, error) {
	return s.ClientProximityV1(ClientProximityV1QueryParams)
}

// Alias Function
func (s *ClientsService) RetrievesTheNumberOfClientsByApplyingComplexFilters(requestClientsRetrievesTheNumberOfClientsByApplyingComplexFiltersV1 *RequestClientsRetrievesTheNumberOfClientsByApplyingComplexFiltersV1, RetrievesTheNumberOfClientsByApplyingComplexFiltersV1HeaderParams *RetrievesTheNumberOfClientsByApplyingComplexFiltersV1HeaderParams) (*ResponseClientsRetrievesTheNumberOfClientsByApplyingComplexFiltersV1, *resty.Response, error) {
	return s.RetrievesTheNumberOfClientsByApplyingComplexFiltersV1(requestClientsRetrievesTheNumberOfClientsByApplyingComplexFiltersV1, RetrievesTheNumberOfClientsByApplyingComplexFiltersV1HeaderParams)
}

// Alias Function
func (s *ClientsService) RetrievesSummaryAnalyticsDataRelatedToClients(requestClientsRetrievesSummaryAnalyticsDataRelatedToClientsV1 *RequestClientsRetrievesSummaryAnalyticsDataRelatedToClientsV1, RetrievesSummaryAnalyticsDataRelatedToClientsV1HeaderParams *RetrievesSummaryAnalyticsDataRelatedToClientsV1HeaderParams) (*ResponseClientsRetrievesSummaryAnalyticsDataRelatedToClientsV1, *resty.Response, error) {
	return s.RetrievesSummaryAnalyticsDataRelatedToClientsV1(requestClientsRetrievesSummaryAnalyticsDataRelatedToClientsV1, RetrievesSummaryAnalyticsDataRelatedToClientsV1HeaderParams)
}

// Alias Function
func (s *ClientsService) RetrievesTheTopNAnalyticsDataRelatedToClients(requestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsV1 *RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsV1, RetrievesTheTopNAnalyticsDataRelatedToClientsV1HeaderParams *RetrievesTheTopNAnalyticsDataRelatedToClientsV1HeaderParams) (*ResponseClientsRetrievesTheTopNAnalyticsDataRelatedToClientsV1, *resty.Response, error) {
	return s.RetrievesTheTopNAnalyticsDataRelatedToClientsV1(requestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsV1, RetrievesTheTopNAnalyticsDataRelatedToClientsV1HeaderParams)
}

// Alias Function
func (s *ClientsService) RetrievesSpecificClientInformationMatchingTheMacaddress(id string, RetrievesSpecificClientInformationMatchingTheMACAddressV1HeaderParams *RetrievesSpecificClientInformationMatchingTheMacaddressV1HeaderParams, RetrievesSpecificClientInformationMatchingTheMACAddressV1QueryParams *RetrievesSpecificClientInformationMatchingTheMacaddressV1QueryParams) (*ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressV1, *resty.Response, error) {
	return s.RetrievesSpecificClientInformationMatchingTheMacaddressV1(id, RetrievesSpecificClientInformationMatchingTheMACAddressV1HeaderParams, RetrievesSpecificClientInformationMatchingTheMACAddressV1QueryParams)
}

// Alias Function
func (s *ClientsService) RetrievesSpecificClientInformationOverASpecifiedPeriodOfTime(id string, requestClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1 *RequestClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1, RetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1HeaderParams *RetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1HeaderParams) (*ResponseClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1, *resty.Response, error) {
	return s.RetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1(id, requestClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1, RetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeV1HeaderParams)
}

// Alias Function
func (s *ClientsService) RetrievesTheTrendAnalyticsDataRelatedToClients(requestClientsRetrievesTheTrendAnalyticsDataRelatedToClientsV1 *RequestClientsRetrievesTheTrendAnalyticsDataRelatedToClientsV1, RetrievesTheTrendAnalyticsDataRelatedToClientsV1HeaderParams *RetrievesTheTrendAnalyticsDataRelatedToClientsV1HeaderParams) (*ResponseClientsRetrievesTheTrendAnalyticsDataRelatedToClientsV1, *resty.Response, error) {
	return s.RetrievesTheTrendAnalyticsDataRelatedToClientsV1(requestClientsRetrievesTheTrendAnalyticsDataRelatedToClientsV1, RetrievesTheTrendAnalyticsDataRelatedToClientsV1HeaderParams)
}

// Alias Function
func (s *ClientsService) RetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributes(requestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1 *RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1, RetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1HeaderParams *RetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1HeaderParams) (*ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1, *resty.Response, error) {
	return s.RetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1(requestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1, RetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesV1HeaderParams)
}

// Alias Function
func (s *ClientsService) GetOverallClientHealth(GetOverallClientHealthV1QueryParams *GetOverallClientHealthV1QueryParams) (*ResponseClientsGetOverallClientHealthV1, *resty.Response, error) {
	return s.GetOverallClientHealthV1(GetOverallClientHealthV1QueryParams)
}

// Alias Function
func (s *ClientsService) RetrievesTheTotalCountOfClientsByApplyingBasicFiltering(RetrievesTheTotalCountOfClientsByApplyingBasicFilteringV1HeaderParams *RetrievesTheTotalCountOfClientsByApplyingBasicFilteringV1HeaderParams, RetrievesTheTotalCountOfClientsByApplyingBasicFilteringV1QueryParams *RetrievesTheTotalCountOfClientsByApplyingBasicFilteringV1QueryParams) (*ResponseClientsRetrievesTheTotalCountOfClientsByApplyingBasicFilteringV1, *resty.Response, error) {
	return s.RetrievesTheTotalCountOfClientsByApplyingBasicFilteringV1(RetrievesTheTotalCountOfClientsByApplyingBasicFilteringV1HeaderParams, RetrievesTheTotalCountOfClientsByApplyingBasicFilteringV1QueryParams)
}
