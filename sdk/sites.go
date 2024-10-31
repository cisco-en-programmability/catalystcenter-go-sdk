package catalyst

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SitesService service

type ReadListOfSiteHealthSummariesV1QueryParams struct {
	StartTime       float64 `url:"startTime,omitempty"`       //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to current time.
	EndTime         float64 `url:"endTime,omitempty"`         //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	Limit           float64 `url:"limit,omitempty"`           //Maximum number of records to return
	Offset          float64 `url:"offset,omitempty"`          //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	SortBy          string  `url:"sortBy,omitempty"`          //A field within the response to sort by.
	Order           string  `url:"order,omitempty"`           //The sort order of the field ascending or descending.
	SiteHierarchy   string  `url:"siteHierarchy,omitempty"`   //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. `Global/AreaName/BuildingName/FloorName`) This field supports wildcard asterisk (`*`) character search support. E.g. `*/San*, */San, /San*` Examples: `?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)
	SiteHierarchyID string  `url:"siteHierarchyId,omitempty"` //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*` Examples: `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	SiteType        string  `url:"siteType,omitempty"`        //The type of the site. A site can be an area, building, or floor. Default when not provided will be `[floor,building,area]` Examples: `?siteType=area` (single siteType requested) `?siteType=area&siteType=building&siteType=floor` (multiple siteTypes requested)
	ID              string  `url:"id,omitempty"`              //The list of entity Uuids. (Ex."6bef213c-19ca-4170-8375-b694e251101c") Examples: id=6bef213c-19ca-4170-8375-b694e251101c (single entity uuid requested) id=6bef213c-19ca-4170-8375-b694e251101c&id=32219612-819e-4b5e-a96b-cf22aca13dd9&id=2541e9a7-b80d-4955-8aa2-79b233318ba0 (multiple entity uuid with '&' separator)
	View            string  `url:"view,omitempty"`            //The specific summary view being requested. This is an optional parameter which can be passed to get one or more of the specific health data summaries associated with sites. ### Response data proviced by each view:   1. **site** [id, siteHierarchy, siteHierarchyId, siteType, latitude, longitude]   2. **network** [id, networkDeviceCount, networkDeviceGoodHealthCount,wirelessDeviceCount, wirelessDeviceGoodHealthCount, accessDeviceCount, accessDeviceGoodHealthCount, coreDeviceCount, coreDeviceGoodHealthCount, distributionDeviceCount, distributionDeviceGoodHealthCount, routerDeviceCount, routerDeviceGoodHealthCount, apDeviceCount, apDeviceGoodHealthCount, wlcDeviceCount, wlcDeviceGoodHealthCount, switchDeviceCount, switchDeviceGoodHealthCount, networkDeviceGoodHealthPercentage, accessDeviceGoodHealthPercentage, coreDeviceGoodHealthPercentage, distributionDeviceGoodHealthPercentage, routerDeviceGoodHealthPercentage, apDeviceGoodHealthPercentage, wlcDeviceGoodHealthPercentage, switchDeviceGoodHealthPercentage, wirelessDeviceGoodHealthPercentage]   3. **client** [id, clientCount, clientGoodHealthCount, wiredClientCount, wirelessClientCount, wiredClientGoodHealthCount, wirelessClientGoodHealthCount, clientGoodHealthPercentage, wiredClientGoodHealthPercentage, wirelessClientGoodHealthPercentage, clientDataUsage]   4. **issue** [id, p1IssueCount, p2IssueCount, p3IssueCount, p4IssueCount, issueCount]   When this query parameter is not added the default summaries are:   **[site,client,network,issue]** Examples: view=client (single view requested) view=client&view=network&view=issue (multiple views requested)
	Attribute       string  `url:"attribute,omitempty"`       //Supported Attributes: [id, siteHierarchy, siteHierarchyId, siteType, latitude, longitude, networkDeviceCount, networkDeviceGoodHealthCount,wirelessDeviceCount, wirelessDeviceGoodHealthCount, accessDeviceCount, accessDeviceGoodHealthCount, coreDeviceCount, coreDeviceGoodHealthCount, distributionDeviceCount, distributionDeviceGoodHealthCount, routerDeviceCount, routerDeviceGoodHealthCount, apDeviceCount, apDeviceGoodHealthCount, wlcDeviceCount, wlcDeviceGoodHealthCount, switchDeviceCount, switchDeviceGoodHealthCount, networkDeviceGoodHealthPercentage, accessDeviceGoodHealthPercentage, coreDeviceGoodHealthPercentage, distributionDeviceGoodHealthPercentage, routerDeviceGoodHealthPercentage, apDeviceGoodHealthPercentage, wlcDeviceGoodHealthPercentage, switchDeviceGoodHealthPercentage, wirelessDeviceGoodHealthPercentage, clientCount, clientGoodHealthCount, wiredClientCount, wirelessClientCount, wiredClientGoodHealthCount, wirelessClientGoodHealthCount, clientGoodHealthPercentage, wiredClientGoodHealthPercentage, wirelessClientGoodHealthPercentage, clientDataUsage, p1IssueCount, p2IssueCount, p3IssueCount, p4IssueCount, issueCount] If length of attribute list is too long, please use 'view' param instead. Examples: attribute=siteHierarchy (single attribute requested) attribute=siteHierarchy&attribute=clientCount (multiple attributes requested)
}
type ReadListOfSiteHealthSummariesV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type ReadSiteCountV1QueryParams struct {
	EndTime         float64 `url:"endTime,omitempty"`         //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	SiteHierarchy   string  `url:"siteHierarchy,omitempty"`   //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. `Global/AreaName/BuildingName/FloorName`) This field supports wildcard asterisk (`*`) character search support. E.g. `*/San*, */San, /San*` Examples: `?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)
	SiteHierarchyID string  `url:"siteHierarchyId,omitempty"` //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*` Examples: `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	SiteType        string  `url:"siteType,omitempty"`        //The type of the site. A site can be an area, building, or floor. Default when not provided will be `[floor,building,area]` Examples: `?siteType=area` (single siteType requested) `?siteType=area&siteType=building&siteType=floor` (multiple siteTypes requested)
	ID              string  `url:"id,omitempty"`              //The list of entity Uuids. (Ex."6bef213c-19ca-4170-8375-b694e251101c") Examples: id=6bef213c-19ca-4170-8375-b694e251101c (single entity uuid requested) id=6bef213c-19ca-4170-8375-b694e251101c&id=32219612-819e-4b5e-a96b-cf22aca13dd9&id=2541e9a7-b80d-4955-8aa2-79b233318ba0 (multiple entity uuid with '&' separator)
}
type ReadSiteCountV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type ReadAnAggregatedSummaryOfSiteHealthDataV1QueryParams struct {
	StartTime       float64 `url:"startTime,omitempty"`       //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to current time.
	EndTime         float64 `url:"endTime,omitempty"`         //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	SiteHierarchy   string  `url:"siteHierarchy,omitempty"`   //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. `Global/AreaName/BuildingName/FloorName`) This field supports wildcard asterisk (`*`) character search support. E.g. `*/San*, */San, /San*` Examples: `?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)
	SiteHierarchyID string  `url:"siteHierarchyId,omitempty"` //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*` Examples: `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	SiteType        string  `url:"siteType,omitempty"`        //The type of the site. A site can be an area, building, or floor. Default when not provided will be `[floor,building,area]` Examples: `?siteType=area` (single siteType requested) `?siteType=area&siteType=building&siteType=floor` (multiple siteTypes requested)
	ID              string  `url:"id,omitempty"`              //The list of entity Uuids. (Ex."6bef213c-19ca-4170-8375-b694e251101c") Examples: id=6bef213c-19ca-4170-8375-b694e251101c (single entity uuid requested) id=6bef213c-19ca-4170-8375-b694e251101c&id=32219612-819e-4b5e-a96b-cf22aca13dd9&id=2541e9a7-b80d-4955-8aa2-79b233318ba0 (multiple entity uuid with '&' separator)
	View            string  `url:"view,omitempty"`            //The specific summary view being requested. This is an optional parameter which can be passed to get one or more of the specific health data summaries associated with sites. ### Response data proviced by each view:   1. **site** [id, siteHierarchy, siteHierarchyId, siteType, latitude, longitude]   2. **network** [id, networkDeviceCount, networkDeviceGoodHealthCount,wirelessDeviceCount, wirelessDeviceGoodHealthCount, accessDeviceCount, accessDeviceGoodHealthCount, coreDeviceCount, coreDeviceGoodHealthCount, distributionDeviceCount, distributionDeviceGoodHealthCount, routerDeviceCount, routerDeviceGoodHealthCount, apDeviceCount, apDeviceGoodHealthCount, wlcDeviceCount, wlcDeviceGoodHealthCount, switchDeviceCount, switchDeviceGoodHealthCount, networkDeviceGoodHealthPercentage, accessDeviceGoodHealthPercentage, coreDeviceGoodHealthPercentage, distributionDeviceGoodHealthPercentage, routerDeviceGoodHealthPercentage, apDeviceGoodHealthPercentage, wlcDeviceGoodHealthPercentage, switchDeviceGoodHealthPercentage, wirelessDeviceGoodHealthPercentage]   3. **client** [id, clientCount, clientGoodHealthCount, wiredClientCount, wirelessClientCount, wiredClientGoodHealthCount, wirelessClientGoodHealthCount, clientGoodHealthPercentage, wiredClientGoodHealthPercentage, wirelessClientGoodHealthPercentage, clientDataUsage]   4. **issue** [id, p1IssueCount, p2IssueCount, p3IssueCount, p4IssueCount, issueCount]   When this query parameter is not added the default summaries are:   **[site,client,network,issue]** Examples: view=client (single view requested) view=client&view=network&view=issue (multiple views requested)
	Attribute       string  `url:"attribute,omitempty"`       //Supported Attributes: [id, siteHierarchy, siteHierarchyId, siteType, latitude, longitude, networkDeviceCount, networkDeviceGoodHealthCount,wirelessDeviceCount, wirelessDeviceGoodHealthCount, accessDeviceCount, accessDeviceGoodHealthCount, coreDeviceCount, coreDeviceGoodHealthCount, distributionDeviceCount, distributionDeviceGoodHealthCount, routerDeviceCount, routerDeviceGoodHealthCount, apDeviceCount, apDeviceGoodHealthCount, wlcDeviceCount, wlcDeviceGoodHealthCount, switchDeviceCount, switchDeviceGoodHealthCount, networkDeviceGoodHealthPercentage, accessDeviceGoodHealthPercentage, coreDeviceGoodHealthPercentage, distributionDeviceGoodHealthPercentage, routerDeviceGoodHealthPercentage, apDeviceGoodHealthPercentage, wlcDeviceGoodHealthPercentage, switchDeviceGoodHealthPercentage, wirelessDeviceGoodHealthPercentage, clientCount, clientGoodHealthCount, wiredClientCount, wirelessClientCount, wiredClientGoodHealthCount, wirelessClientGoodHealthCount, clientGoodHealthPercentage, wiredClientGoodHealthPercentage, wirelessClientGoodHealthPercentage, clientDataUsage, p1IssueCount, p2IssueCount, p3IssueCount, p4IssueCount, issueCount] If length of attribute list is too long, please use 'view' param instead. Examples: attribute=siteHierarchy (single attribute requested) attribute=siteHierarchy&attribute=clientCount (multiple attributes requested)
}
type ReadAnAggregatedSummaryOfSiteHealthDataV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type QueryAnAggregatedSummaryOfSiteHealthDataV1QueryParams struct {
	SiteHierarchy   string `url:"siteHierarchy,omitempty"`   //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. `Global/AreaName/BuildingName/FloorName`) This field supports wildcard asterisk (`*`) character search support. E.g. `*/San*, */San, /San*` Examples: `?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)
	SiteHierarchyID string `url:"siteHierarchyId,omitempty"` //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*` Examples: `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	SiteType        string `url:"siteType,omitempty"`        //The type of the site. A site can be an area, building, or floor. Default when not provided will be `[floor,building,area]` Examples: `?siteType=area` (single siteType requested) `?siteType=area&siteType=building&siteType=floor` (multiple siteTypes requested)
	ID              string `url:"id,omitempty"`              //The list of entity Uuids. (Ex."6bef213c-19ca-4170-8375-b694e251101c") Examples: id=6bef213c-19ca-4170-8375-b694e251101c (single entity uuid requested) id=6bef213c-19ca-4170-8375-b694e251101c&id=32219612-819e-4b5e-a96b-cf22aca13dd9&id=2541e9a7-b80d-4955-8aa2-79b233318ba0 (multiple entity uuid with '&' separator)
}
type ReadSiteHealthSummaryDataBySiteIDV1QueryParams struct {
	StartTime float64 `url:"startTime,omitempty"` //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to current time.
	EndTime   float64 `url:"endTime,omitempty"`   //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	View      string  `url:"view,omitempty"`      //The specific summary view being requested. This is an optional parameter which can be passed to get one or more of the specific health data summaries associated with sites. ### Response data proviced by each view:   1. **site** [id, siteHierarchy, siteHierarchyId, siteType, latitude, longitude]   2. **network** [id, networkDeviceCount, networkDeviceGoodHealthCount,wirelessDeviceCount, wirelessDeviceGoodHealthCount, accessDeviceCount, accessDeviceGoodHealthCount, coreDeviceCount, coreDeviceGoodHealthCount, distributionDeviceCount, distributionDeviceGoodHealthCount, routerDeviceCount, routerDeviceGoodHealthCount, apDeviceCount, apDeviceGoodHealthCount, wlcDeviceCount, wlcDeviceGoodHealthCount, switchDeviceCount, switchDeviceGoodHealthCount, networkDeviceGoodHealthPercentage, accessDeviceGoodHealthPercentage, coreDeviceGoodHealthPercentage, distributionDeviceGoodHealthPercentage, routerDeviceGoodHealthPercentage, apDeviceGoodHealthPercentage, wlcDeviceGoodHealthPercentage, switchDeviceGoodHealthPercentage, wirelessDeviceGoodHealthPercentage]   3. **client** [id, clientCount, clientGoodHealthCount, wiredClientCount, wirelessClientCount, wiredClientGoodHealthCount, wirelessClientGoodHealthCount, clientGoodHealthPercentage, wiredClientGoodHealthPercentage, wirelessClientGoodHealthPercentage, clientDataUsage]   4. **issue** [id, p1IssueCount, p2IssueCount, p3IssueCount, p4IssueCount, issueCount]   When this query parameter is not added the default summaries are:   **[site,client,network,issue]** Examples: view=client (single view requested) view=client&view=network&view=issue (multiple views requested)
	Attribute string  `url:"attribute,omitempty"` //Supported Attributes: [id, siteHierarchy, siteHierarchyId, siteType, latitude, longitude, networkDeviceCount, networkDeviceGoodHealthCount,wirelessDeviceCount, wirelessDeviceGoodHealthCount, accessDeviceCount, accessDeviceGoodHealthCount, coreDeviceCount, coreDeviceGoodHealthCount, distributionDeviceCount, distributionDeviceGoodHealthCount, routerDeviceCount, routerDeviceGoodHealthCount, apDeviceCount, apDeviceGoodHealthCount, wlcDeviceCount, wlcDeviceGoodHealthCount, switchDeviceCount, switchDeviceGoodHealthCount, networkDeviceGoodHealthPercentage, accessDeviceGoodHealthPercentage, coreDeviceGoodHealthPercentage, distributionDeviceGoodHealthPercentage, routerDeviceGoodHealthPercentage, apDeviceGoodHealthPercentage, wlcDeviceGoodHealthPercentage, switchDeviceGoodHealthPercentage, wirelessDeviceGoodHealthPercentage, clientCount, clientGoodHealthCount, wiredClientCount, wirelessClientCount, wiredClientGoodHealthCount, wirelessClientGoodHealthCount, clientGoodHealthPercentage, wiredClientGoodHealthPercentage, wirelessClientGoodHealthPercentage, clientDataUsage, p1IssueCount, p2IssueCount, p3IssueCount, p4IssueCount, issueCount] If length of attribute list is too long, please use 'view' param instead. Examples: attribute=siteHierarchy (single attribute requested) attribute=siteHierarchy&attribute=clientCount (multiple attributes requested)
}
type ReadSiteHealthSummaryDataBySiteIDV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type AssignDevicesToSiteV1HeaderParams struct {
	Runsync           string `url:"__runsync,omitempty"`           //Expects type bool. Enable this parameter to execute the API and return a response synchronously
	Timeout           string `url:"__timeout,omitempty"`           //Expects type float64. During synchronous execution, this defines the maximum time to wait for a response, before the API execution is terminated
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type bool. Persist bapi sync response
}
type GetMembershipV1QueryParams struct {
	Offset       float64 `url:"offset,omitempty"`       //offset/starting row
	Limit        float64 `url:"limit,omitempty"`        //Number of sites to be retrieved
	DeviceFamily string  `url:"deviceFamily,omitempty"` //Device family name
	SerialNumber string  `url:"serialNumber,omitempty"` //Device serial number
}
type CreateSiteV1HeaderParams struct {
	Runsync           string `url:"__runsync,omitempty"`           //Expects type bool. Enable this parameter to execute the API and return a response synchronously
	Timeout           string `url:"__timeout,omitempty"`           //Expects type float64. During synchronous execution, this defines the maximum time to wait for a response, before the API execution is terminated
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type bool. Persist bapi sync response
}
type GetSiteV1QueryParams struct {
	Name   string `url:"name,omitempty"`   //Site name hierarchy (E.g Global/USA/CA)
	SiteID string `url:"siteId,omitempty"` //Site Id
	Type   string `url:"type,omitempty"`   //Site type (Ex: area, building, floor)
	Offset int    `url:"offset,omitempty"` //Offset/starting index for pagination. Indexed from 1.
	Limit  int    `url:"limit,omitempty"`  //Number of sites to be listed
}
type GetSiteHealthV1QueryParams struct {
	SiteType  string  `url:"siteType,omitempty"`  //site type: AREA or BUILDING (case insensitive)
	Offset    float64 `url:"offset,omitempty"`    //Offset of the first returned data set entry (Multiple of 'limit' + 1)
	Limit     float64 `url:"limit,omitempty"`     //Max number of data entries in the returned data set [1,50].  Default is 25
	Timestamp float64 `url:"timestamp,omitempty"` //Epoch time(in milliseconds) when the Site Hierarchy data is required
}
type GetDevicesThatAreAssignedToASiteV1QueryParams struct {
	Offset     string `url:"offset,omitempty"`     //Offset/starting index for pagination
	Limit      string `url:"limit,omitempty"`      //Number of devices to be listed. Default and max supported value is 500
	MemberType string `url:"memberType,omitempty"` //Member type (This API only supports the 'networkdevice' type)
	Level      string `url:"level,omitempty"`      //Depth of site hierarchy to be considered to list the devices. If the provided value is -1, devices for all child sites will be listed.
}
type GetSiteCountV1QueryParams struct {
	SiteID string `url:"siteId,omitempty"` //Site instance UUID
}
type UpdateSiteV1HeaderParams struct {
	Runsync           string `url:"__runsync,omitempty"`           //Expects type bool. Enable this parameter to execute the API and return a response synchronously
	Timeout           string `url:"__timeout,omitempty"`           //Expects type float64. During synchronous execution, this defines the maximum time to wait for a response, before the API execution is terminated
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type bool. Persist bapi sync response
}
type GetSiteV2QueryParams struct {
	GroupNameHierarchy string `url:"groupNameHierarchy,omitempty"` //Site name hierarchy (E.g. Global/USA/CA)
	ID                 string `url:"id,omitempty"`                 //Site Id
	Type               string `url:"type,omitempty"`               //Site type (Acceptable values: area, building, floor)
	Offset             string `url:"offset,omitempty"`             //Offset/starting index for pagination
	Limit              string `url:"limit,omitempty"`              //Number of sites to be listed. Default and max supported value is 500
}
type GetSiteCountV2QueryParams struct {
	ID string `url:"id,omitempty"` //Site instance UUID
}

type ResponseSitesReadListOfSiteHealthSummariesV1 struct {
	Response *[]ResponseSitesReadListOfSiteHealthSummariesV1Response `json:"response,omitempty"` //
	Page     *ResponseSitesReadListOfSiteHealthSummariesV1Page       `json:"page,omitempty"`     //
	Version  string                                                  `json:"version,omitempty"`  // Version
}
type ResponseSitesReadListOfSiteHealthSummariesV1Response struct {
	ID                                     string   `json:"id,omitempty"`                                     // Id
	SiteHierarchy                          string   `json:"siteHierarchy,omitempty"`                          // Site Hierarchy
	SiteHierarchyID                        string   `json:"siteHierarchyId,omitempty"`                        // Site Hierarchy Id
	SiteType                               string   `json:"siteType,omitempty"`                               // Site Type
	Latitude                               *float64 `json:"latitude,omitempty"`                               // Latitude
	Longitude                              *float64 `json:"longitude,omitempty"`                              // Longitude
	NetworkDeviceGoodHealthPercentage      *int     `json:"networkDeviceGoodHealthPercentage,omitempty"`      // Network Device Good Health Percentage
	NetworkDeviceGoodHealthCount           *int     `json:"networkDeviceGoodHealthCount,omitempty"`           // Network Device Good Health Count
	ClientGoodHealthCount                  *int     `json:"clientGoodHealthCount,omitempty"`                  // Client Good Health Count
	ClientGoodHealthPercentage             *int     `json:"clientGoodHealthPercentage,omitempty"`             // Client Good Health Percentage
	WiredClientGoodHealthPercentage        *int     `json:"wiredClientGoodHealthPercentage,omitempty"`        // Wired Client Good Health Percentage
	WirelessClientGoodHealthPercentage     *int     `json:"wirelessClientGoodHealthPercentage,omitempty"`     // Wireless Client Good Health Percentage
	ClientCount                            *int     `json:"clientCount,omitempty"`                            // Client Count
	WiredClientCount                       *int     `json:"wiredClientCount,omitempty"`                       // Wired Client Count
	WirelessClientCount                    *int     `json:"wirelessClientCount,omitempty"`                    // Wireless Client Count
	WiredClientGoodHealthCount             *int     `json:"wiredClientGoodHealthCount,omitempty"`             // Wired Client Good Health Count
	WirelessClientGoodHealthCount          *int     `json:"wirelessClientGoodHealthCount,omitempty"`          // Wireless Client Good Health Count
	NetworkDeviceCount                     *int     `json:"networkDeviceCount,omitempty"`                     // Network Device Count
	AccessDeviceCount                      *int     `json:"accessDeviceCount,omitempty"`                      // Access Device Count
	AccessDeviceGoodHealthCount            *int     `json:"accessDeviceGoodHealthCount,omitempty"`            // Access Device Good Health Count
	CoreDeviceCount                        *int     `json:"coreDeviceCount,omitempty"`                        // Core Device Count
	CoreDeviceGoodHealthCount              *int     `json:"coreDeviceGoodHealthCount,omitempty"`              // Core Device Good Health Count
	DistributionDeviceCount                *int     `json:"distributionDeviceCount,omitempty"`                // Distribution Device Count
	DistributionDeviceGoodHealthCount      *int     `json:"distributionDeviceGoodHealthCount,omitempty"`      // Distribution Device Good Health Count
	RouterDeviceCount                      *int     `json:"routerDeviceCount,omitempty"`                      // Router Device Count
	RouterDeviceGoodHealthCount            *int     `json:"routerDeviceGoodHealthCount,omitempty"`            // Router Device Good Health Count
	WirelessDeviceCount                    *int     `json:"wirelessDeviceCount,omitempty"`                    // Wireless Device Count
	WirelessDeviceGoodHealthCount          *int     `json:"wirelessDeviceGoodHealthCount,omitempty"`          // Wireless Device Good Health Count
	ApDeviceCount                          *int     `json:"apDeviceCount,omitempty"`                          // Ap Device Count
	ApDeviceGoodHealthCount                *int     `json:"apDeviceGoodHealthCount,omitempty"`                // Ap Device Good Health Count
	WlcDeviceCount                         *int     `json:"wlcDeviceCount,omitempty"`                         // Wlc Device Count
	WlcDeviceGoodHealthCount               *int     `json:"wlcDeviceGoodHealthCount,omitempty"`               // Wlc Device Good Health Count
	SwitchDeviceCount                      *int     `json:"switchDeviceCount,omitempty"`                      // Switch Device Count
	SwitchDeviceGoodHealthCount            *int     `json:"switchDeviceGoodHealthCount,omitempty"`            // Switch Device Good Health Count
	AccessDeviceGoodHealthPercentage       *int     `json:"accessDeviceGoodHealthPercentage,omitempty"`       // Access Device Good Health Percentage
	CoreDeviceGoodHealthPercentage         *int     `json:"coreDeviceGoodHealthPercentage,omitempty"`         // Core Device Good Health Percentage
	DistributionDeviceGoodHealthPercentage *int     `json:"distributionDeviceGoodHealthPercentage,omitempty"` // Distribution Device Good Health Percentage
	RouterDeviceGoodHealthPercentage       *int     `json:"routerDeviceGoodHealthPercentage,omitempty"`       // Router Device Good Health Percentage
	ApDeviceGoodHealthPercentage           *int     `json:"apDeviceGoodHealthPercentage,omitempty"`           // Ap Device Good Health Percentage
	WlcDeviceGoodHealthPercentage          *int     `json:"wlcDeviceGoodHealthPercentage,omitempty"`          // Wlc Device Good Health Percentage
	SwitchDeviceGoodHealthPercentage       *int     `json:"switchDeviceGoodHealthPercentage,omitempty"`       // Switch Device Good Health Percentage
	WirelessDeviceGoodHealthPercentage     *int     `json:"wirelessDeviceGoodHealthPercentage,omitempty"`     // Wireless Device Good Health Percentage
	ClientDataUsage                        *float64 `json:"clientDataUsage,omitempty"`                        // Client Data Usage
	P1IssueCount                           *int     `json:"p1IssueCount,omitempty"`                           // P1 Issue Count
	P2IssueCount                           *int     `json:"p2IssueCount,omitempty"`                           // P2 Issue Count
	P3IssueCount                           *int     `json:"p3IssueCount,omitempty"`                           // P3 Issue Count
	P4IssueCount                           *int     `json:"p4IssueCount,omitempty"`                           // P4 Issue Count
	IssueCount                             *int     `json:"issueCount,omitempty"`                             // Issue Count
}
type ResponseSitesReadListOfSiteHealthSummariesV1Page struct {
	Limit  *int                                                      `json:"limit,omitempty"`  // Limit
	Offset *int                                                      `json:"offset,omitempty"` // Offset
	Count  *int                                                      `json:"count,omitempty"`  // Count
	SortBy *[]ResponseSitesReadListOfSiteHealthSummariesV1PageSortBy `json:"sortBy,omitempty"` //
}
type ResponseSitesReadListOfSiteHealthSummariesV1PageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseSitesReadSiteCountV1 struct {
	Response *ResponseSitesReadSiteCountV1Response `json:"response,omitempty"` //
	Version  string                                `json:"version,omitempty"`  // Version
}
type ResponseSitesReadSiteCountV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseSitesReadAnAggregatedSummaryOfSiteHealthDataV1 struct {
	Response *ResponseSitesReadAnAggregatedSummaryOfSiteHealthDataV1Response `json:"response,omitempty"` //
	Version  string                                                          `json:"version,omitempty"`  // Version
}
type ResponseSitesReadAnAggregatedSummaryOfSiteHealthDataV1Response struct {
	ID                                     string   `json:"id,omitempty"`                                     // Id
	SiteHierarchy                          string   `json:"siteHierarchy,omitempty"`                          // Site Hierarchy
	SiteHierarchyID                        string   `json:"siteHierarchyId,omitempty"`                        // Site Hierarchy Id
	SiteType                               string   `json:"siteType,omitempty"`                               // Site Type
	Latitude                               *float64 `json:"latitude,omitempty"`                               // Latitude
	Longitude                              *float64 `json:"longitude,omitempty"`                              // Longitude
	NetworkDeviceGoodHealthPercentage      *int     `json:"networkDeviceGoodHealthPercentage,omitempty"`      // Network Device Good Health Percentage
	NetworkDeviceGoodHealthCount           *int     `json:"networkDeviceGoodHealthCount,omitempty"`           // Network Device Good Health Count
	ClientGoodHealthCount                  *int     `json:"clientGoodHealthCount,omitempty"`                  // Client Good Health Count
	ClientGoodHealthPercentage             *int     `json:"clientGoodHealthPercentage,omitempty"`             // Client Good Health Percentage
	WiredClientGoodHealthPercentage        *int     `json:"wiredClientGoodHealthPercentage,omitempty"`        // Wired Client Good Health Percentage
	WirelessClientGoodHealthPercentage     *int     `json:"wirelessClientGoodHealthPercentage,omitempty"`     // Wireless Client Good Health Percentage
	ClientCount                            *int     `json:"clientCount,omitempty"`                            // Client Count
	WiredClientCount                       *int     `json:"wiredClientCount,omitempty"`                       // Wired Client Count
	WirelessClientCount                    *int     `json:"wirelessClientCount,omitempty"`                    // Wireless Client Count
	WiredClientGoodHealthCount             *int     `json:"wiredClientGoodHealthCount,omitempty"`             // Wired Client Good Health Count
	WirelessClientGoodHealthCount          *int     `json:"wirelessClientGoodHealthCount,omitempty"`          // Wireless Client Good Health Count
	NetworkDeviceCount                     *int     `json:"networkDeviceCount,omitempty"`                     // Network Device Count
	AccessDeviceCount                      *int     `json:"accessDeviceCount,omitempty"`                      // Access Device Count
	AccessDeviceGoodHealthCount            *int     `json:"accessDeviceGoodHealthCount,omitempty"`            // Access Device Good Health Count
	CoreDeviceCount                        *int     `json:"coreDeviceCount,omitempty"`                        // Core Device Count
	CoreDeviceGoodHealthCount              *int     `json:"coreDeviceGoodHealthCount,omitempty"`              // Core Device Good Health Count
	DistributionDeviceCount                *int     `json:"distributionDeviceCount,omitempty"`                // Distribution Device Count
	DistributionDeviceGoodHealthCount      *int     `json:"distributionDeviceGoodHealthCount,omitempty"`      // Distribution Device Good Health Count
	RouterDeviceCount                      *int     `json:"routerDeviceCount,omitempty"`                      // Router Device Count
	RouterDeviceGoodHealthCount            *int     `json:"routerDeviceGoodHealthCount,omitempty"`            // Router Device Good Health Count
	WirelessDeviceCount                    *int     `json:"wirelessDeviceCount,omitempty"`                    // Wireless Device Count
	WirelessDeviceGoodHealthCount          *int     `json:"wirelessDeviceGoodHealthCount,omitempty"`          // Wireless Device Good Health Count
	ApDeviceCount                          *int     `json:"apDeviceCount,omitempty"`                          // Ap Device Count
	ApDeviceGoodHealthCount                *int     `json:"apDeviceGoodHealthCount,omitempty"`                // Ap Device Good Health Count
	WlcDeviceCount                         *int     `json:"wlcDeviceCount,omitempty"`                         // Wlc Device Count
	WlcDeviceGoodHealthCount               *int     `json:"wlcDeviceGoodHealthCount,omitempty"`               // Wlc Device Good Health Count
	SwitchDeviceCount                      *int     `json:"switchDeviceCount,omitempty"`                      // Switch Device Count
	SwitchDeviceGoodHealthCount            *int     `json:"switchDeviceGoodHealthCount,omitempty"`            // Switch Device Good Health Count
	AccessDeviceGoodHealthPercentage       *int     `json:"accessDeviceGoodHealthPercentage,omitempty"`       // Access Device Good Health Percentage
	CoreDeviceGoodHealthPercentage         *int     `json:"coreDeviceGoodHealthPercentage,omitempty"`         // Core Device Good Health Percentage
	DistributionDeviceGoodHealthPercentage *int     `json:"distributionDeviceGoodHealthPercentage,omitempty"` // Distribution Device Good Health Percentage
	RouterDeviceGoodHealthPercentage       *int     `json:"routerDeviceGoodHealthPercentage,omitempty"`       // Router Device Good Health Percentage
	ApDeviceGoodHealthPercentage           *int     `json:"apDeviceGoodHealthPercentage,omitempty"`           // Ap Device Good Health Percentage
	WlcDeviceGoodHealthPercentage          *int     `json:"wlcDeviceGoodHealthPercentage,omitempty"`          // Wlc Device Good Health Percentage
	SwitchDeviceGoodHealthPercentage       *int     `json:"switchDeviceGoodHealthPercentage,omitempty"`       // Switch Device Good Health Percentage
	WirelessDeviceGoodHealthPercentage     *int     `json:"wirelessDeviceGoodHealthPercentage,omitempty"`     // Wireless Device Good Health Percentage
	ClientDataUsage                        *float64 `json:"clientDataUsage,omitempty"`                        // Client Data Usage
	P1IssueCount                           *int     `json:"p1IssueCount,omitempty"`                           // P1 Issue Count
	P2IssueCount                           *int     `json:"p2IssueCount,omitempty"`                           // P2 Issue Count
	P3IssueCount                           *int     `json:"p3IssueCount,omitempty"`                           // P3 Issue Count
	P4IssueCount                           *int     `json:"p4IssueCount,omitempty"`                           // P4 Issue Count
	IssueCount                             *int     `json:"issueCount,omitempty"`                             // Issue Count
}
type ResponseSitesQueryAnAggregatedSummaryOfSiteHealthDataV1 struct {
	Response *ResponseSitesQueryAnAggregatedSummaryOfSiteHealthDataV1Response `json:"response,omitempty"` //
	Version  string                                                           `json:"version,omitempty"`  // Version
}
type ResponseSitesQueryAnAggregatedSummaryOfSiteHealthDataV1Response struct {
	ID                                     string   `json:"id,omitempty"`                                     // Id
	SiteHierarchy                          string   `json:"siteHierarchy,omitempty"`                          // Site Hierarchy
	SiteHierarchyID                        string   `json:"siteHierarchyId,omitempty"`                        // Site Hierarchy Id
	SiteType                               string   `json:"siteType,omitempty"`                               // Site Type
	Latitude                               *float64 `json:"latitude,omitempty"`                               // Latitude
	Longitude                              *float64 `json:"longitude,omitempty"`                              // Longitude
	NetworkDeviceGoodHealthPercentage      *int     `json:"networkDeviceGoodHealthPercentage,omitempty"`      // Network Device Good Health Percentage
	NetworkDeviceGoodHealthCount           *int     `json:"networkDeviceGoodHealthCount,omitempty"`           // Network Device Good Health Count
	ClientGoodHealthCount                  *int     `json:"clientGoodHealthCount,omitempty"`                  // Client Good Health Count
	ClientGoodHealthPercentage             *int     `json:"clientGoodHealthPercentage,omitempty"`             // Client Good Health Percentage
	WiredClientGoodHealthPercentage        *int     `json:"wiredClientGoodHealthPercentage,omitempty"`        // Wired Client Good Health Percentage
	WirelessClientGoodHealthPercentage     *int     `json:"wirelessClientGoodHealthPercentage,omitempty"`     // Wireless Client Good Health Percentage
	ClientCount                            *int     `json:"clientCount,omitempty"`                            // Client Count
	WiredClientCount                       *int     `json:"wiredClientCount,omitempty"`                       // Wired Client Count
	WirelessClientCount                    *int     `json:"wirelessClientCount,omitempty"`                    // Wireless Client Count
	WiredClientGoodHealthCount             *int     `json:"wiredClientGoodHealthCount,omitempty"`             // Wired Client Good Health Count
	WirelessClientGoodHealthCount          *int     `json:"wirelessClientGoodHealthCount,omitempty"`          // Wireless Client Good Health Count
	NetworkDeviceCount                     *int     `json:"networkDeviceCount,omitempty"`                     // Network Device Count
	AccessDeviceCount                      *int     `json:"accessDeviceCount,omitempty"`                      // Access Device Count
	AccessDeviceGoodHealthCount            *int     `json:"accessDeviceGoodHealthCount,omitempty"`            // Access Device Good Health Count
	CoreDeviceCount                        *int     `json:"coreDeviceCount,omitempty"`                        // Core Device Count
	CoreDeviceGoodHealthCount              *int     `json:"coreDeviceGoodHealthCount,omitempty"`              // Core Device Good Health Count
	DistributionDeviceCount                *int     `json:"distributionDeviceCount,omitempty"`                // Distribution Device Count
	DistributionDeviceGoodHealthCount      *int     `json:"distributionDeviceGoodHealthCount,omitempty"`      // Distribution Device Good Health Count
	RouterDeviceCount                      *int     `json:"routerDeviceCount,omitempty"`                      // Router Device Count
	RouterDeviceGoodHealthCount            *int     `json:"routerDeviceGoodHealthCount,omitempty"`            // Router Device Good Health Count
	WirelessDeviceCount                    *int     `json:"wirelessDeviceCount,omitempty"`                    // Wireless Device Count
	WirelessDeviceGoodHealthCount          *int     `json:"wirelessDeviceGoodHealthCount,omitempty"`          // Wireless Device Good Health Count
	ApDeviceCount                          *int     `json:"apDeviceCount,omitempty"`                          // Ap Device Count
	ApDeviceGoodHealthCount                *int     `json:"apDeviceGoodHealthCount,omitempty"`                // Ap Device Good Health Count
	WlcDeviceCount                         *int     `json:"wlcDeviceCount,omitempty"`                         // Wlc Device Count
	WlcDeviceGoodHealthCount               *int     `json:"wlcDeviceGoodHealthCount,omitempty"`               // Wlc Device Good Health Count
	SwitchDeviceCount                      *int     `json:"switchDeviceCount,omitempty"`                      // Switch Device Count
	SwitchDeviceGoodHealthCount            *int     `json:"switchDeviceGoodHealthCount,omitempty"`            // Switch Device Good Health Count
	AccessDeviceGoodHealthPercentage       *int     `json:"accessDeviceGoodHealthPercentage,omitempty"`       // Access Device Good Health Percentage
	CoreDeviceGoodHealthPercentage         *int     `json:"coreDeviceGoodHealthPercentage,omitempty"`         // Core Device Good Health Percentage
	DistributionDeviceGoodHealthPercentage *int     `json:"distributionDeviceGoodHealthPercentage,omitempty"` // Distribution Device Good Health Percentage
	RouterDeviceGoodHealthPercentage       *int     `json:"routerDeviceGoodHealthPercentage,omitempty"`       // Router Device Good Health Percentage
	ApDeviceGoodHealthPercentage           *int     `json:"apDeviceGoodHealthPercentage,omitempty"`           // Ap Device Good Health Percentage
	WlcDeviceGoodHealthPercentage          *int     `json:"wlcDeviceGoodHealthPercentage,omitempty"`          // Wlc Device Good Health Percentage
	SwitchDeviceGoodHealthPercentage       *int     `json:"switchDeviceGoodHealthPercentage,omitempty"`       // Switch Device Good Health Percentage
	WirelessDeviceGoodHealthPercentage     *int     `json:"wirelessDeviceGoodHealthPercentage,omitempty"`     // Wireless Device Good Health Percentage
	ClientDataUsage                        *float64 `json:"clientDataUsage,omitempty"`                        // Client Data Usage
	P1IssueCount                           *int     `json:"p1IssueCount,omitempty"`                           // P1 Issue Count
	P2IssueCount                           *int     `json:"p2IssueCount,omitempty"`                           // P2 Issue Count
	P3IssueCount                           *int     `json:"p3IssueCount,omitempty"`                           // P3 Issue Count
	P4IssueCount                           *int     `json:"p4IssueCount,omitempty"`                           // P4 Issue Count
	IssueCount                             *int     `json:"issueCount,omitempty"`                             // Issue Count
}
type ResponseSitesReadSiteHealthSummaryDataBySiteIDV1 struct {
	Response *ResponseSitesReadSiteHealthSummaryDataBySiteIDV1Response `json:"response,omitempty"` //
	Version  string                                                    `json:"version,omitempty"`  // Version
}
type ResponseSitesReadSiteHealthSummaryDataBySiteIDV1Response struct {
	ID                                     string   `json:"id,omitempty"`                                     // Id
	SiteHierarchy                          string   `json:"siteHierarchy,omitempty"`                          // Site Hierarchy
	SiteHierarchyID                        string   `json:"siteHierarchyId,omitempty"`                        // Site Hierarchy Id
	SiteType                               string   `json:"siteType,omitempty"`                               // Site Type
	Latitude                               *float64 `json:"latitude,omitempty"`                               // Latitude
	Longitude                              *float64 `json:"longitude,omitempty"`                              // Longitude
	NetworkDeviceGoodHealthPercentage      *int     `json:"networkDeviceGoodHealthPercentage,omitempty"`      // Network Device Good Health Percentage
	NetworkDeviceGoodHealthCount           *int     `json:"networkDeviceGoodHealthCount,omitempty"`           // Network Device Good Health Count
	ClientGoodHealthCount                  *int     `json:"clientGoodHealthCount,omitempty"`                  // Client Good Health Count
	ClientGoodHealthPercentage             *int     `json:"clientGoodHealthPercentage,omitempty"`             // Client Good Health Percentage
	WiredClientGoodHealthPercentage        *int     `json:"wiredClientGoodHealthPercentage,omitempty"`        // Wired Client Good Health Percentage
	WirelessClientGoodHealthPercentage     *int     `json:"wirelessClientGoodHealthPercentage,omitempty"`     // Wireless Client Good Health Percentage
	ClientCount                            *int     `json:"clientCount,omitempty"`                            // Client Count
	WiredClientCount                       *int     `json:"wiredClientCount,omitempty"`                       // Wired Client Count
	WirelessClientCount                    *int     `json:"wirelessClientCount,omitempty"`                    // Wireless Client Count
	WiredClientGoodHealthCount             *int     `json:"wiredClientGoodHealthCount,omitempty"`             // Wired Client Good Health Count
	WirelessClientGoodHealthCount          *int     `json:"wirelessClientGoodHealthCount,omitempty"`          // Wireless Client Good Health Count
	NetworkDeviceCount                     *int     `json:"networkDeviceCount,omitempty"`                     // Network Device Count
	AccessDeviceCount                      *int     `json:"accessDeviceCount,omitempty"`                      // Access Device Count
	AccessDeviceGoodHealthCount            *int     `json:"accessDeviceGoodHealthCount,omitempty"`            // Access Device Good Health Count
	CoreDeviceCount                        *int     `json:"coreDeviceCount,omitempty"`                        // Core Device Count
	CoreDeviceGoodHealthCount              *int     `json:"coreDeviceGoodHealthCount,omitempty"`              // Core Device Good Health Count
	DistributionDeviceCount                *int     `json:"distributionDeviceCount,omitempty"`                // Distribution Device Count
	DistributionDeviceGoodHealthCount      *int     `json:"distributionDeviceGoodHealthCount,omitempty"`      // Distribution Device Good Health Count
	RouterDeviceCount                      *int     `json:"routerDeviceCount,omitempty"`                      // Router Device Count
	RouterDeviceGoodHealthCount            *int     `json:"routerDeviceGoodHealthCount,omitempty"`            // Router Device Good Health Count
	WirelessDeviceCount                    *int     `json:"wirelessDeviceCount,omitempty"`                    // Wireless Device Count
	WirelessDeviceGoodHealthCount          *int     `json:"wirelessDeviceGoodHealthCount,omitempty"`          // Wireless Device Good Health Count
	ApDeviceCount                          *int     `json:"apDeviceCount,omitempty"`                          // Ap Device Count
	ApDeviceGoodHealthCount                *int     `json:"apDeviceGoodHealthCount,omitempty"`                // Ap Device Good Health Count
	WlcDeviceCount                         *int     `json:"wlcDeviceCount,omitempty"`                         // Wlc Device Count
	WlcDeviceGoodHealthCount               *int     `json:"wlcDeviceGoodHealthCount,omitempty"`               // Wlc Device Good Health Count
	SwitchDeviceCount                      *int     `json:"switchDeviceCount,omitempty"`                      // Switch Device Count
	SwitchDeviceGoodHealthCount            *int     `json:"switchDeviceGoodHealthCount,omitempty"`            // Switch Device Good Health Count
	AccessDeviceGoodHealthPercentage       *int     `json:"accessDeviceGoodHealthPercentage,omitempty"`       // Access Device Good Health Percentage
	CoreDeviceGoodHealthPercentage         *int     `json:"coreDeviceGoodHealthPercentage,omitempty"`         // Core Device Good Health Percentage
	DistributionDeviceGoodHealthPercentage *int     `json:"distributionDeviceGoodHealthPercentage,omitempty"` // Distribution Device Good Health Percentage
	RouterDeviceGoodHealthPercentage       *int     `json:"routerDeviceGoodHealthPercentage,omitempty"`       // Router Device Good Health Percentage
	ApDeviceGoodHealthPercentage           *int     `json:"apDeviceGoodHealthPercentage,omitempty"`           // Ap Device Good Health Percentage
	WlcDeviceGoodHealthPercentage          *int     `json:"wlcDeviceGoodHealthPercentage,omitempty"`          // Wlc Device Good Health Percentage
	SwitchDeviceGoodHealthPercentage       *int     `json:"switchDeviceGoodHealthPercentage,omitempty"`       // Switch Device Good Health Percentage
	WirelessDeviceGoodHealthPercentage     *int     `json:"wirelessDeviceGoodHealthPercentage,omitempty"`     // Wireless Device Good Health Percentage
	ClientDataUsage                        *float64 `json:"clientDataUsage,omitempty"`                        // Client Data Usage
	P1IssueCount                           *int     `json:"p1IssueCount,omitempty"`                           // P1 Issue Count
	P2IssueCount                           *int     `json:"p2IssueCount,omitempty"`                           // P2 Issue Count
	P3IssueCount                           *int     `json:"p3IssueCount,omitempty"`                           // P3 Issue Count
	P4IssueCount                           *int     `json:"p4IssueCount,omitempty"`                           // P4 Issue Count
	IssueCount                             *int     `json:"issueCount,omitempty"`                             // Issue Count
}
type ResponseSitesAssignDevicesToSiteV1 struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseSitesExportMapArchiveV1 struct {
	Response *ResponseSitesExportMapArchiveV1Response `json:"response,omitempty"` //
	Version  string                                   `json:"version,omitempty"`  //
}
type ResponseSitesExportMapArchiveV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
} // # Review unknown case
type ResponseSitesImportMapArchiveCancelAnImportV1 interface{}
type ResponseSitesImportMapArchivePerformImportV1 interface{}
type ResponseSitesImportMapArchiveImportStatusV1 struct {
	AuditLog *ResponseSitesImportMapArchiveImportStatusV1AuditLog `json:"auditLog,omitempty"` //
	Status   string                                               `json:"status,omitempty"`   //
	UUID     *ResponseSitesImportMapArchiveImportStatusV1UUID     `json:"uuid,omitempty"`     //
}
type ResponseSitesImportMapArchiveImportStatusV1AuditLog struct {
	Children                   *[]ResponseSitesImportMapArchiveImportStatusV1AuditLogChildren              `json:"children,omitempty"`                   //
	EntitiesCount              *[]ResponseSitesImportMapArchiveImportStatusV1AuditLogEntitiesCount         `json:"entitiesCount,omitempty"`              //
	EntityName                 string                                                                      `json:"entityName,omitempty"`                 //
	EntityType                 string                                                                      `json:"entityType,omitempty"`                 //
	ErrorEntitiesCount         *[]ResponseSitesImportMapArchiveImportStatusV1AuditLogErrorEntitiesCount    `json:"errorEntitiesCount,omitempty"`         //
	Errors                     *[]ResponseSitesImportMapArchiveImportStatusV1AuditLogErrors                `json:"errors,omitempty"`                     //
	Infos                      *[]ResponseSitesImportMapArchiveImportStatusV1AuditLogInfos                 `json:"infos,omitempty"`                      //
	MatchingEntitiesCount      *[]ResponseSitesImportMapArchiveImportStatusV1AuditLogMatchingEntitiesCount `json:"matchingEntitiesCount,omitempty"`      //
	SubTasksRootTaskID         string                                                                      `json:"subTasksRootTaskId,omitempty"`         //
	SuccessfullyImportedFloors []string                                                                    `json:"successfullyImportedFloors,omitempty"` //
	Warnings                   *[]ResponseSitesImportMapArchiveImportStatusV1AuditLogWarnings              `json:"warnings,omitempty"`                   //
}
type ResponseSitesImportMapArchiveImportStatusV1AuditLogChildren interface{}
type ResponseSitesImportMapArchiveImportStatusV1AuditLogEntitiesCount struct {
	Key *int `json:"key,omitempty"` //
}
type ResponseSitesImportMapArchiveImportStatusV1AuditLogErrorEntitiesCount struct {
	Key *int `json:"key,omitempty"` //
}
type ResponseSitesImportMapArchiveImportStatusV1AuditLogErrors struct {
	Message string `json:"message,omitempty"` //
}
type ResponseSitesImportMapArchiveImportStatusV1AuditLogInfos struct {
	Message string `json:"message,omitempty"` //
}
type ResponseSitesImportMapArchiveImportStatusV1AuditLogMatchingEntitiesCount struct {
	Key *int `json:"key,omitempty"` //
}
type ResponseSitesImportMapArchiveImportStatusV1AuditLogWarnings struct {
	Message string `json:"message,omitempty"` //
}
type ResponseSitesImportMapArchiveImportStatusV1UUID struct {
	LeastSignificantBits *int `json:"leastSignificantBits,omitempty"` //
	MostSignificantBits  *int `json:"mostSignificantBits,omitempty"`  //
}
type ResponseSitesMapsSupportedAccessPointsV1 []ResponseItemSitesMapsSupportedAccessPointsV1 // Array of ResponseSitesMapsSupportedAccessPointsV1
type ResponseItemSitesMapsSupportedAccessPointsV1 struct {
	AntennaPatterns *[]ResponseItemSitesMapsSupportedAccessPointsV1AntennaPatterns `json:"antennaPatterns,omitempty"` //
	ApType          string                                                         `json:"apType,omitempty"`          //
}
type ResponseItemSitesMapsSupportedAccessPointsV1AntennaPatterns struct {
	Band  string   `json:"band,omitempty"`  //
	Names []string `json:"names,omitempty"` //
}
type ResponseSitesGetMembershipV1 struct {
	Site   *ResponseSitesGetMembershipV1Site     `json:"site,omitempty"`   //
	Device *[]ResponseSitesGetMembershipV1Device `json:"device,omitempty"` //
}
type ResponseSitesGetMembershipV1Site struct {
	Response *[]ResponseSitesGetMembershipV1SiteResponse `json:"response,omitempty"` // Response
	Version  string                                      `json:"version,omitempty"`  // Version
}
type ResponseSitesGetMembershipV1SiteResponse interface{}
type ResponseSitesGetMembershipV1Device struct {
	Response *[]ResponseSitesGetMembershipV1DeviceResponse `json:"response,omitempty"` // Response
	Version  string                                        `json:"version,omitempty"`  // Version
	SiteID   string                                        `json:"siteId,omitempty"`   // Site Id
}
type ResponseSitesGetMembershipV1DeviceResponse interface{}
type ResponseSitesCreateSiteV1 struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseSitesGetSiteV1 struct {
	Response *[]ResponseSitesGetSiteV1Response `json:"response,omitempty"` //
}
type ResponseSitesGetSiteV1Response struct {
	ParentID          string                                         `json:"parentId,omitempty"`          // Parent Id
	Name              string                                         `json:"name,omitempty"`              // Name
	AdditionalInfo    []ResponseSitesGetSiteV1ResponseAdditionalInfo `json:"additionalInfo,omitempty"`    //
	SiteHierarchy     string                                         `json:"siteHierarchy,omitempty"`     // Site Hierarchy
	SiteNameHierarchy string                                         `json:"siteNameHierarchy,omitempty"` // Site Name Hierarchy
	InstanceTenantID  string                                         `json:"instanceTenantId,omitempty"`  // Instance Tenant Id
	ID                string                                         `json:"id,omitempty"`                // Id
}
type ResponseSitesGetSiteV1ResponseAdditionalInfo struct {
	Namespace  string                                                 `json:"nameSpace,omitempty"`  //
	Attributes ResponseSitesGetSiteV1ResponseAdditionalInfoAttributes `json:"attributes,omitempty"` //
}

//type ResponseSitesGetSiteResponseAdditionalInfoAttributes map[string]string

type ResponseSitesGetSiteV1ResponseAdditionalInfoAttributes struct {
	Country              string `json:"country,omitempty"`              //
	Address              string `json:"address,omitempty"`              //
	Latitude             string `json:"latitude,omitempty"`             //
	AddressInheritedFrom string `json:"addressInheritedFrom,omitempty"` //
	Type                 string `json:"type,omitempty"`                 //
	Longitude            string `json:"longitude,omitempty"`            //
	OffsetX              string `json:"offsetX,omitempty"`              //
	OffsetY              string `json:"offsetY,omitempty"`              //
	Length               string `json:"length,omitempty"`               //
	Width                string `json:"width,omitempty"`                //
	Height               string `json:"height,omitempty"`               //
	RfModel              string `json:"rfModel,omitempty"`              //
	FloorIndex           string `json:"floorIndex,omitempty"`           //
}

// Area
type ResponseSitesGetAreaV1 struct {
	Response *[]ResponseSitesGetAreaV1Response `json:"response,omitempty"` //
}
type ResponseSitesGetAreaV1Response struct {
	ParentID          string                                         `json:"parentId,omitempty"`          // Parent Id
	Name              string                                         `json:"name,omitempty"`              // Name
	AdditionalInfo    []ResponseSitesGetAreaV1ResponseAdditionalInfo `json:"additionalInfo,omitempty"`    //
	SiteHierarchy     string                                         `json:"siteHierarchy,omitempty"`     // Site Hierarchy
	SiteNameHierarchy string                                         `json:"siteNameHierarchy,omitempty"` // Site Name Hierarchy
	InstanceTenantID  string                                         `json:"instanceTenantId,omitempty"`  // Instance Tenant Id
	ID                string                                         `json:"id,omitempty"`
	ParentName        string                                         `json:"parent_name,omitempty"` // Id
}
type ResponseSitesGetAreaV1ResponseAdditionalInfo struct {
	Namespace  string                                                 `json:"nameSpace,omitempty"`  //
	Attributes ResponseSitesGetAreaV1ResponseAdditionalInfoAttributes `json:"attributes,omitempty"` //
}

//type ResponseSitesGetSiteResponseAdditionalInfoAttributes map[string]string

type ResponseSitesGetAreaV1ResponseAdditionalInfoAttributes struct {
	Name                 string `json:"name,omitempty"` //
	ParentName           string `json:"parent_name,omitempty"`
	AddressInheritedFrom string `json:"addressinheritedfrom,omitempty"` //
	Type                 string `json:"type,omitempty"`                 //
}

// Floor
type ResponseSitesGetFloorV1 struct {
	Response *[]ResponseSitesGetFloorV1Response `json:"response,omitempty"` //
}
type ResponseSitesGetFloorV1Response struct {
	ParentID          string                                          `json:"parentId,omitempty"`          // Parent Id
	Name              string                                          `json:"name,omitempty"`              // Name
	AdditionalInfo    []ResponseSitesGetFloorV1ResponseAdditionalInfo `json:"additionalInfo,omitempty"`    //
	SiteHierarchy     string                                          `json:"siteHierarchy,omitempty"`     // Site Hierarchy
	SiteNameHierarchy string                                          `json:"siteNameHierarchy,omitempty"` // Site Name Hierarchy
	InstanceTenantID  string                                          `json:"instanceTenantId,omitempty"`  // Instance Tenant Id
	ID                string                                          `json:"id,omitempty"`
	ParentName        string                                          `json:"parent_name,omitempty"` // Id
}
type ResponseSitesGetFloorV1ResponseAdditionalInfo struct {
	Namespace  string                                                  `json:"nameSpace,omitempty"`  //
	Attributes ResponseSitesGetFloorV1ResponseAdditionalInfoAttributes `json:"attributes,omitempty"` //
}

type ResponseSitesGetFloorV1ResponseAdditionalInfoAttributes struct {
	FloorIndex string `json:"floorIndex,omitempty"` //
	Height     string `json:"height,omitempty"`
	Length     string `json:"length,omitempty"`      //
	Name       string `json:"name,omitempty"`        //
	ParentName string `json:"parent_name,omitempty"` //
	RfModel    string `json:"rfmodel,omitempty"`     //
	Width      string `json:"width,omitempty"`       //
}

type ResponseSitesGetSiteHealthV1 struct {
	Response *[]ResponseSitesGetSiteHealthV1Response `json:"response,omitempty"` //
}
type ResponseSitesGetSiteHealthV1Response struct {
	SiteName                           string                                                       `json:"siteName,omitempty"`                           // Name of the site
	SiteID                             string                                                       `json:"siteId,omitempty"`                             // Site UUID
	ParentSiteID                       string                                                       `json:"parentSiteId,omitempty"`                       // The parent site's UUID of this site
	ParentSiteName                     string                                                       `json:"parentSiteName,omitempty"`                     // The parent site's name of this site
	SiteType                           string                                                       `json:"siteType,omitempty"`                           // Site type of this site
	Latitude                           *float64                                                     `json:"latitude,omitempty"`                           // Site (building) location's latitude
	Longitude                          *float64                                                     `json:"longitude,omitempty"`                          // Site (building) location's longitude
	HealthyNetworkDevicePercentage     *int                                                         `json:"healthyNetworkDevicePercentage,omitempty"`     // Network health of devices on the site
	HealthyClientsPercentage           *int                                                         `json:"healthyClientsPercentage,omitempty"`           // Client health of all clients in the site
	ClientHealthWired                  *int                                                         `json:"clientHealthWired,omitempty"`                  // Health of all wired clients in the site
	ClientHealthWireless               *int                                                         `json:"clientHealthWireless,omitempty"`               // Health of all wireless clients in the site
	NumberOfClients                    *int                                                         `json:"numberOfClients,omitempty"`                    // Total number of clients in the site
	NumberOfNetworkDevice              *int                                                         `json:"numberOfNetworkDevice,omitempty"`              // Total number of network devices in the site
	NetworkHealthAverage               *int                                                         `json:"networkHealthAverage,omitempty"`               // Average network health in the site
	NetworkHealthAccess                *int                                                         `json:"networkHealthAccess,omitempty"`                // Network health for access devices in the site
	NetworkHealthCore                  *int                                                         `json:"networkHealthCore,omitempty"`                  // Network health for core devices in the site
	NetworkHealthDistribution          *int                                                         `json:"networkHealthDistribution,omitempty"`          // Network health for distribution devices in the site
	NetworkHealthRouter                *int                                                         `json:"networkHealthRouter,omitempty"`                // Network health for router devices in the site
	NetworkHealthWireless              *int                                                         `json:"networkHealthWireless,omitempty"`              // Network health for wireless devices in the site
	NetworkHealthAP                    *int                                                         `json:"networkHealthAP,omitempty"`                    // Network health for AP devices in the site
	NetworkHealthWLC                   *int                                                         `json:"networkHealthWLC,omitempty"`                   // Network health for WLC devices in the site
	NetworkHealthSwitch                *int                                                         `json:"networkHealthSwitch,omitempty"`                // Network health for switch devices in the site
	NetworkHealthOthers                *int                                                         `json:"networkHealthOthers,omitempty"`                // Network health for other devices in the site
	NumberOfWiredClients               *int                                                         `json:"numberOfWiredClients,omitempty"`               // Number of wired clients in the site
	NumberOfWirelessClients            *int                                                         `json:"numberOfWirelessClients,omitempty"`            // Number of wireless clients in the site
	TotalNumberOfConnectedWiredClients *int                                                         `json:"totalNumberOfConnectedWiredClients,omitempty"` // Number of connected wired clients in the site
	TotalNumberOfActiveWirelessClients *int                                                         `json:"totalNumberOfActiveWirelessClients,omitempty"` // Number of active wireless clients in the site
	WiredGoodClients                   *int                                                         `json:"wiredGoodClients,omitempty"`                   // Number of GOOD health wired clients in the site
	WirelessGoodClients                *int                                                         `json:"wirelessGoodClients,omitempty"`                // Number of GOOD health wireless clients in the site
	OverallGoodDevices                 *int                                                         `json:"overallGoodDevices,omitempty"`                 // Number of GOOD health devices in the site
	AccessGoodCount                    *int                                                         `json:"accessGoodCount,omitempty"`                    // Number of GOOD health access devices in the site
	AccessTotalCount                   *int                                                         `json:"accessTotalCount,omitempty"`                   // Number of access devices in the site
	CoreGoodCount                      *int                                                         `json:"coreGoodCount,omitempty"`                      // Number of GOOD health core devices in the site
	CoreTotalCount                     *int                                                         `json:"coreTotalCount,omitempty"`                     // Number of core devices in the site
	DistributionGoodCount              *int                                                         `json:"distributionGoodCount,omitempty"`              // Number of GOOD health distribution devices in the site
	DistributionTotalCount             *int                                                         `json:"distributionTotalCount,omitempty"`             // Number of distribution devices in the site
	RouterGoodCount                    *int                                                         `json:"routerGoodCount,omitempty"`                    // Number of GOOD health router in the site
	RouterTotalCount                   *int                                                         `json:"routerTotalCount,omitempty"`                   // Number of router devices in the site
	WirelessDeviceGoodCount            *int                                                         `json:"wirelessDeviceGoodCount,omitempty"`            // Number of GOOD health wireless devices in the site
	WirelessDeviceTotalCount           *int                                                         `json:"wirelessDeviceTotalCount,omitempty"`           // Number of wireless devices in the site
	ApDeviceGoodCount                  *int                                                         `json:"apDeviceGoodCount,omitempty"`                  // Number of GOOD health AP devices in the site
	ApDeviceTotalCount                 *int                                                         `json:"apDeviceTotalCount,omitempty"`                 // Number of AP devices in the site
	WlcDeviceGoodCount                 *int                                                         `json:"wlcDeviceGoodCount,omitempty"`                 // Number of GOOD health wireless controller devices in the site
	WlcDeviceTotalCount                *int                                                         `json:"wlcDeviceTotalCount,omitempty"`                // Number of wireless controller devices in the site
	SwitchDeviceGoodCount              *int                                                         `json:"switchDeviceGoodCount,omitempty"`              // Number of GOOD health switch devices in the site
	SwitchDeviceTotalCount             *int                                                         `json:"switchDeviceTotalCount,omitempty"`             // Number of switch devices in the site
	ApplicationHealth                  *int                                                         `json:"applicationHealth,omitempty"`                  // Average application health in the site
	ApplicationHealthInfo              *[]ResponseSitesGetSiteHealthV1ResponseApplicationHealthInfo `json:"applicationHealthInfo,omitempty"`              //
	ApplicationGoodCount               *int                                                         `json:"applicationGoodCount,omitempty"`               // Number of GOOD health applications int the site
	ApplicationTotalCount              *int                                                         `json:"applicationTotalCount,omitempty"`              // Number of applications int the site
	ApplicationBytesTotalCount         *float64                                                     `json:"applicationBytesTotalCount,omitempty"`         // Total application bytes
	DnacInfo                           *ResponseSitesGetSiteHealthV1ResponseDnacInfo                `json:"dnacInfo,omitempty"`                           //
	Usage                              *float64                                                     `json:"usage,omitempty"`                              // Total bits used by all clients in a site
	ApplicationHealthStats             *ResponseSitesGetSiteHealthV1ResponseApplicationHealthStats  `json:"applicationHealthStats,omitempty"`             //
}
type ResponseSitesGetSiteHealthV1ResponseApplicationHealthInfo struct {
	TrafficClass string   `json:"trafficClass,omitempty"` // Traffic class of the application
	BytesCount   *float64 `json:"bytesCount,omitempty"`   // Byte count of the application
	HealthScore  *int     `json:"healthScore,omitempty"`  // Health score of the application
}
type ResponseSitesGetSiteHealthV1ResponseDnacInfo struct {
	UUID   string `json:"uuid,omitempty"`   // UUID of the DNAC
	IP     string `json:"ip,omitempty"`     // IP address of the DNAC
	Status string `json:"status,omitempty"` // Status of the DNAC
}
type ResponseSitesGetSiteHealthV1ResponseApplicationHealthStats struct {
	AppTotalCount              *float64                                                                              `json:"appTotalCount,omitempty"`              // Total application count
	BusinessRelevantAppCount   *ResponseSitesGetSiteHealthV1ResponseApplicationHealthStatsBusinessRelevantAppCount   `json:"businessRelevantAppCount,omitempty"`   //
	BusinessIrrelevantAppCount *ResponseSitesGetSiteHealthV1ResponseApplicationHealthStatsBusinessIrrelevantAppCount `json:"businessIrrelevantAppCount,omitempty"` //
	DefaultHealthAppCount      *ResponseSitesGetSiteHealthV1ResponseApplicationHealthStatsDefaultHealthAppCount      `json:"defaultHealthAppCount,omitempty"`      //
}
type ResponseSitesGetSiteHealthV1ResponseApplicationHealthStatsBusinessRelevantAppCount struct {
	Poor *float64 `json:"poor,omitempty"` // Poor business relevant application count
	Fair *float64 `json:"fair,omitempty"` // Fair business relevant application count
	Good *float64 `json:"good,omitempty"` // Good business relevant application count
}
type ResponseSitesGetSiteHealthV1ResponseApplicationHealthStatsBusinessIrrelevantAppCount struct {
	Poor *float64 `json:"poor,omitempty"` // Poor business irrelevant application count
	Fair *float64 `json:"fair,omitempty"` // Fair business irrelevant application count
	Good *float64 `json:"good,omitempty"` // Good business irrelevant application count
}
type ResponseSitesGetSiteHealthV1ResponseApplicationHealthStatsDefaultHealthAppCount struct {
	Poor *float64 `json:"poor,omitempty"` // Poor default application count
	Fair *float64 `json:"fair,omitempty"` // Fair default application count
	Good *float64 `json:"good,omitempty"` // Good default application count
}
type ResponseSitesGetDevicesThatAreAssignedToASiteV1 struct {
	Response *[]ResponseSitesGetDevicesThatAreAssignedToASiteV1Response `json:"response,omitempty"` //
}
type ResponseSitesGetDevicesThatAreAssignedToASiteV1Response struct {
	InstanceUUID                  string `json:"instanceUuid,omitempty"`                  // Device UUID (E.g. 48eebb3e-b3fc-4928-a7df-1c80e216f930)
	InstanceID                    *int   `json:"instanceId,omitempty"`                    // Device Id (E.g. 230230)
	AuthEntityID                  *int   `json:"authEntityId,omitempty"`                  // Authentication Entity Id (Internal record)
	AuthEntityClass               *int   `json:"authEntityClass,omitempty"`               // Authentication entity class (Internal record)
	InstanceTenantID              string `json:"instanceTenantId,omitempty"`              // Device tenant Id (E.g. 64472cc32d3bc1658597669c)
	DeployPending                 string `json:"deployPending,omitempty"`                 // Deploy pending (Internal record)
	InstanceVersion               *int   `json:"instanceVersion,omitempty"`               // Instance version (Internal record)
	ApManagerInterfaceIP          string `json:"apManagerInterfaceIp,omitempty"`          // Access Point manager interface IP
	AssociatedWlcIP               string `json:"associatedWlcIp,omitempty"`               // Associated Wireless Controller IP
	BootDateTime                  string `json:"bootDateTime,omitempty"`                  // Device boot date and time
	CollectionInterval            string `json:"collectionInterval,omitempty"`            // Device resync interval type (E.g. Global Default)
	CollectionIntervalValue       string `json:"collectionIntervalValue,omitempty"`       // Device resync interval value
	CollectionStatus              string `json:"collectionStatus,omitempty"`              // Device inventory collection status (E.g. Managed)
	Description                   string `json:"description,omitempty"`                   // Device description
	DeviceSupportLevel            string `json:"deviceSupportLevel,omitempty"`            // Device support level (E.g. Supported)
	DNSResolvedManagementAddress  string `json:"dnsResolvedManagementAddress,omitempty"`  // DNS resolved management address
	Family                        string `json:"family,omitempty"`                        // Device family (E.g. Routers)
	Hostname                      string `json:"hostname,omitempty"`                      // Device hostname
	InterfaceCount                string `json:"interfaceCount,omitempty"`                // Device interface count
	InventoryStatusDetail         string `json:"inventoryStatusDetail,omitempty"`         // Device inventory collection status detail
	LastUpdateTime                *int   `json:"lastUpdateTime,omitempty"`                // Last update time
	LastUpdated                   string `json:"lastUpdated,omitempty"`                   // Last updated date and time
	LineCardCount                 string `json:"lineCardCount,omitempty"`                 // Line card count
	LineCardID                    string `json:"lineCardId,omitempty"`                    // Line card Id
	LastDeviceResyncStartTime     string `json:"lastDeviceResyncStartTime,omitempty"`     // Last device inventory resync start date and time
	MacAddress                    string `json:"macAddress,omitempty"`                    // MAC address
	ManagedAtleastOnce            *bool  `json:"managedAtleastOnce,omitempty"`            // If device managed atleast once, value will be true otherwise false
	ManagementIPAddress           string `json:"managementIpAddress,omitempty"`           // Management IP address
	ManagementState               string `json:"managementState,omitempty"`               // Device management state (E.g. Managed)
	MemorySize                    string `json:"memorySize,omitempty"`                    // Memory size
	PaddedMgmtIPAddress           string `json:"paddedMgmtIpAddress,omitempty"`           // Padded management IP address. Internal record
	PendingSyncRequestsCount      string `json:"pendingSyncRequestsCount,omitempty"`      // Pending sync requests count. Internal record
	PlatformID                    string `json:"platformId,omitempty"`                    // Device platform Id (E.g. CSR1000V)
	ReachabilityFailureReason     string `json:"reachabilityFailureReason,omitempty"`     // Device reachability failure reason
	ReachabilityStatus            string `json:"reachabilityStatus,omitempty"`            // Device reachability status (E.g. Reachable)
	ReasonsForDeviceResync        string `json:"reasonsForDeviceResync,omitempty"`        // Reasons for device resync (E.g. Periodic)
	ReasonsForPendingSyncRequests string `json:"reasonsForPendingSyncRequests,omitempty"` // Reasons for pending device sync requests
	Role                          string `json:"role,omitempty"`                          // Device role (E.g. BORDER ROUTER)
	RoleSource                    string `json:"roleSource,omitempty"`                    // Device role source. Internal record
	SerialNumber                  string `json:"serialNumber,omitempty"`                  // Device serial Number
	Series                        string `json:"series,omitempty"`                        // Device series
	SNMPContact                   string `json:"snmpContact,omitempty"`                   // Device snmp contact. Internal record
	SNMPLocation                  string `json:"snmpLocation,omitempty"`                  // Device snmp location
	SoftwareType                  string `json:"softwareType,omitempty"`                  // Device software type
	SoftwareVersion               string `json:"softwareVersion,omitempty"`               // Device software version
	TagCount                      string `json:"tagCount,omitempty"`                      // Device tag Count
	Type                          string `json:"type,omitempty"`                          // Device type (E.g. Cisco Cloud Services Router 1000V)
	UpTime                        string `json:"upTime,omitempty"`                        // Device up time (E.g. 112 days, 6:09:13.86)
	UptimeSeconds                 *int   `json:"uptimeSeconds,omitempty"`                 // Device uptime in seconds
	Vendor                        string `json:"vendor,omitempty"`                        // Vendor (E.g. Cisco)
	DisplayName                   string `json:"displayName,omitempty"`                   // Device display name
}
type ResponseSitesGetSiteCountV1 struct {
	Response *int   `json:"response,omitempty"` // Response
	Version  string `json:"version,omitempty"`  // Version
}
type ResponseSitesUpdateSiteV1 struct {
	Result   string                             `json:"result,omitempty"`   // Result
	Response *ResponseSitesUpdateSiteV1Response `json:"response,omitempty"` //
	Status   string                             `json:"status,omitempty"`   // Status
}
type ResponseSitesUpdateSiteV1Response struct {
	EndTime          string   `json:"endTime,omitempty"`          // End Time
	Version          string   `json:"version,omitempty"`          // Version
	StartTime        string   `json:"startTime,omitempty"`        // Start Time
	Progress         string   `json:"progress,omitempty"`         // Progress
	Data             string   `json:"data,omitempty"`             // Data
	ServiceType      string   `json:"serviceType,omitempty"`      // Service Type
	OperationIDList  []string `json:"operationIdList,omitempty"`  // Operation Id List
	IsError          string   `json:"isError,omitempty"`          // Is Error
	RootID           string   `json:"rootId,omitempty"`           // Root Id
	InstanceTenantID string   `json:"instanceTenantId,omitempty"` // Instance Tenant Id
	ID               string   `json:"id,omitempty"`               // Id
}
type ResponseSitesDeleteSiteV1 struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseSitesGetSiteV2 struct {
	Response *[]ResponseSitesGetSiteV2Response `json:"response,omitempty"` //
}
type ResponseSitesGetSiteV2Response struct {
	ParentID           string                                          `json:"parentId,omitempty"`           // Parent site Instance UUID (e.g. b27181bb-211b-40ec-ba5d-2603867c3f2c)
	GroupTypeList      []string                                        `json:"groupTypeList,omitempty"`      // There are different group types like 'RBAC', 'POLICY', 'SITE', 'TAG', 'PORT', 'DEVICE_TYPE'. This API is for site, so list contains 'SITE' only
	GroupHierarchy     string                                          `json:"groupHierarchy,omitempty"`     // Site hierarchy by instance UUID (e.g. b27181bb-211b-40ec-ba5d-2603867c3f2c/576c7859-e485-4073-a46f-305f475de4c5)
	AdditionalInfo     *[]ResponseSitesGetSiteV2ResponseAdditionalInfo `json:"additionalInfo,omitempty"`     //
	GroupNameHierarchy string                                          `json:"groupNameHierarchy,omitempty"` // Site hierarchy by name (e.g. Global/USA/CA/San Jose/Building4)
	Name               string                                          `json:"name,omitempty"`               // Site name (e.g. Building4)
	InstanceTenantID   string                                          `json:"instanceTenantId,omitempty"`   // Tenant instance Id where site created (e.g. 63bf047b64ec9c1c45f9019c)
	ID                 string                                          `json:"id,omitempty"`                 // Site instance UUID (e.g. bb5122ce-4527-4af5-8718-44b746a3a3d8)
}
type ResponseSitesGetSiteV2ResponseAdditionalInfo struct {
	NameSpace  string                                                  `json:"nameSpace,omitempty"`  // Site name space. Default value is 'Location'
	Attributes *ResponseSitesGetSiteV2ResponseAdditionalInfoAttributes `json:"attributes,omitempty"` //
}
type ResponseSitesGetSiteV2ResponseAdditionalInfoAttributes struct {
	AddressInheritedFrom string `json:"addressInheritedFrom,omitempty"` // Site instance UUID from where address inherited (e.g. 576c7859-e485-4073-a46f-305f475de4c5)
	Type                 string `json:"type,omitempty"`                 // Site type
	Country              string `json:"country,omitempty"`              // Site Country (e.g. United States)
	Address              string `json:"address,omitempty"`              // Site address (e.g. 269 East Tasman Drive, San Jose, California 95134, United States)
	Latitude             string `json:"latitude,omitempty"`             // Site latitude (e.g. 37.413082)
	Longitude            string `json:"longitude,omitempty"`            // Site longitude (e.g. -121.933886)
}
type ResponseSitesGetSiteCountV2 struct {
	Response *int   `json:"response,omitempty"` // Response
	Version  string `json:"version,omitempty"`  // Version
}
type RequestSitesQueryAnAggregatedSummaryOfSiteHealthDataV1 struct {
	StartTime  *int     `json:"startTime,omitempty"`  // Start Time
	EndTime    *int     `json:"endTime,omitempty"`    // End Time
	Views      []string `json:"views,omitempty"`      // Views
	Attributes []string `json:"attributes,omitempty"` // Attributes
}
type RequestSitesAssignDevicesToSiteV1 struct {
	Device *[]RequestSitesAssignDevicesToSiteV1Device `json:"device,omitempty"` //
}
type RequestSitesAssignDevicesToSiteV1Device struct {
	IP string `json:"ip,omitempty"` // Device IP. It can be either IPv4 or IPv6. IPV4 e.g., 10.104.240.64. IPV6 e.g., 2001:420:284:2004:4:181:500:183
} // # Review unknown case
type RequestSitesCreateSiteV1 struct {
	Type string                        `json:"type,omitempty"` // Type of site to create (eg: area, building, floor)
	Site *RequestSitesCreateSiteV1Site `json:"site,omitempty"` //
}
type RequestSitesCreateSiteV1Site struct {
	Area     *RequestSitesCreateSiteV1SiteArea     `json:"area,omitempty"`     //
	Building *RequestSitesCreateSiteV1SiteBuilding `json:"building,omitempty"` //
	Floor    *RequestSitesCreateSiteV1SiteFloor    `json:"floor,omitempty"`    //
}
type RequestSitesCreateSiteV1SiteArea struct {
	Name       string `json:"name,omitempty"`       // Name of the area (eg: Area1)
	ParentName string `json:"parentName,omitempty"` // Parent name of the area to be created
}
type RequestSitesCreateSiteV1SiteBuilding struct {
	Name       string   `json:"name,omitempty"`       // Name of the building (eg: building1)
	Address    string   `json:"address,omitempty"`    // Address of the building to be created
	ParentName string   `json:"parentName,omitempty"` // Parent name of building to be created
	Latitude   *float64 `json:"latitude,omitempty"`   // Latitude coordinate of the building (eg:37.338)
	Longitude  *float64 `json:"longitude,omitempty"`  // Longitude coordinate of the building (eg:-121.832)
	Country    string   `json:"country,omitempty"`    // Country (eg:United States)
}
type RequestSitesCreateSiteV1SiteFloor struct {
	Name        string   `json:"name,omitempty"`        // Name of the floor (eg:floor-1)
	ParentName  string   `json:"parentName,omitempty"`  // Parent name of the floor to be created
	RfModel     string   `json:"rfModel,omitempty"`     // Type of floor (eg: Cubes And Walled Offices0
	Width       *float64 `json:"width,omitempty"`       // Width of the floor. Unit of measure is ft. (eg: 100)
	Length      *float64 `json:"length,omitempty"`      // Length of the floor. Unit of measure is ft. (eg: 100)
	Height      *float64 `json:"height,omitempty"`      // Height of the floor. Unit of measure is ft. (eg: 15)
	FloorNumber *float64 `json:"floorNumber,omitempty"` // Floor number. (eg: 5)
}
type RequestSitesUpdateSiteV1 struct {
	Type string                        `json:"type,omitempty"` // Site type
	Site *RequestSitesUpdateSiteV1Site `json:"site,omitempty"` //
}
type RequestSitesUpdateSiteV1Site struct {
	Area     *RequestSitesUpdateSiteV1SiteArea     `json:"area,omitempty"`     //
	Building *RequestSitesUpdateSiteV1SiteBuilding `json:"building,omitempty"` //
	Floor    *RequestSitesUpdateSiteV1SiteFloor    `json:"floor,omitempty"`    //
}
type RequestSitesUpdateSiteV1SiteArea struct {
	Name       string `json:"name,omitempty"`       // Area name
	ParentName string `json:"parentName,omitempty"` // Parent hierarchical name (Example: Global/USA/CA)
}
type RequestSitesUpdateSiteV1SiteBuilding struct {
	Name       string   `json:"name,omitempty"`       // Building name
	Address    string   `json:"address,omitempty"`    // Building address (Example: 4900 Marie P. Debartolo Way, Santa Clara, California 95054, United States)
	ParentName string   `json:"parentName,omitempty"` // Parent hierarchical name (Example: Global/USA/CA/SantaClara)
	Latitude   *float64 `json:"latitude,omitempty"`   // Building latitude (Example: 37.403712)
	Longitude  *float64 `json:"longitude,omitempty"`  // Building longitude (Example: -121.971063)
	Country    string   `json:"country,omitempty"`    // Country name. This field is mandatory for air-gapped networks (Example: United States)
}
type RequestSitesUpdateSiteV1SiteFloor struct {
	Name        string   `json:"name,omitempty"`        // Floor name
	RfModel     string   `json:"rfModel,omitempty"`     // RF model (Example : Cubes And Walled Offices)
	Width       *float64 `json:"width,omitempty"`       // Floor width in feet (Example: 200)
	Length      *float64 `json:"length,omitempty"`      // Floor length in feet (Example: 100)
	Height      *float64 `json:"height,omitempty"`      // Floor height in feet (Example: 10)
	FloorNumber *float64 `json:"floorNumber,omitempty"` // Floor Number (Example: 3)
}

//ReadListOfSiteHealthSummariesV1 Read list of site health summaries. - e4b7-1b5e-4099-b15b
/* Get a paginated list of site health summaries. Use the available query parameters to identify a subset of sites you want health summaries for. This API provides the latest health data from a given `endTime` If data is not ready for the provided endTime, the request will fail, and the error message will indicate the recommended endTime to use to retrieve a complete data set. This behavior may occur if the provided endTime=currentTime, since we are not a real time system. When `endTime` is not provided, the API returns the latest data. This API also provides issue data. The `startTime` query param can be used to specify the beginning point of time range to retrieve the active issue counts in. When this param is not provided, the default `startTime` will be 24 hours before endTime. Valid values for `sortBy` param in this API are limited to the attributes provided in the `site` view. Default sortBy is 'siteHierarchy' in order 'asc' (ascending). For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-siteHealthSummaries-1.0.3-resolved.yaml


@param ReadListOfSiteHealthSummariesV1HeaderParams Custom header parameters
@param ReadListOfSiteHealthSummariesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!read-list-of-site-health-summaries-v1
*/
func (s *SitesService) ReadListOfSiteHealthSummariesV1(ReadListOfSiteHealthSummariesV1HeaderParams *ReadListOfSiteHealthSummariesV1HeaderParams, ReadListOfSiteHealthSummariesV1QueryParams *ReadListOfSiteHealthSummariesV1QueryParams) (*ResponseSitesReadListOfSiteHealthSummariesV1, *resty.Response, error) {
	path := "/dna/data/api/v1/siteHealthSummaries"

	queryString, _ := query.Values(ReadListOfSiteHealthSummariesV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if ReadListOfSiteHealthSummariesV1HeaderParams != nil {

		if ReadListOfSiteHealthSummariesV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", ReadListOfSiteHealthSummariesV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesReadListOfSiteHealthSummariesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReadListOfSiteHealthSummariesV1(ReadListOfSiteHealthSummariesV1HeaderParams, ReadListOfSiteHealthSummariesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReadListOfSiteHealthSummariesV1")
	}

	result := response.Result().(*ResponseSitesReadListOfSiteHealthSummariesV1)
	return result, response, err

}

//ReadSiteCountV1 Read site count. - b6ac-283b-4f39-b488
/* Get a count of sites. Use the available query parameters to get the count of a subset of sites. This API provides the latest data from a given `endTime` If data is not ready for the provided endTime, the request will fail, and the error message will indicate the recommended endTime to use to retrieve a complete data set. This behavior may occur if the provided endTime=currentTime, since we are not a real time system. When `endTime` is not provided, the API returns the latest data. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-siteHealthSummaries-1.0.3-resolved.yaml


@param ReadSiteCountV1HeaderParams Custom header parameters
@param ReadSiteCountV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!read-site-count-v1
*/
func (s *SitesService) ReadSiteCountV1(ReadSiteCountV1HeaderParams *ReadSiteCountV1HeaderParams, ReadSiteCountV1QueryParams *ReadSiteCountV1QueryParams) (*ResponseSitesReadSiteCountV1, *resty.Response, error) {
	path := "/dna/data/api/v1/siteHealthSummaries/count"

	queryString, _ := query.Values(ReadSiteCountV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if ReadSiteCountV1HeaderParams != nil {

		if ReadSiteCountV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", ReadSiteCountV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesReadSiteCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReadSiteCountV1(ReadSiteCountV1HeaderParams, ReadSiteCountV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReadSiteCountV1")
	}

	result := response.Result().(*ResponseSitesReadSiteCountV1)
	return result, response, err

}

//ReadAnAggregatedSummaryOfSiteHealthDataV1 Read an aggregated summary of site health data. - e2b7-8b97-4c78-a4f7
/* Get an aggregated summary of all site health or use the query params to get an aggregated summary of health for a subset of sites. This API provides the latest health data from a given `endTime` If data is not ready for the provided endTime, the request will fail, and the error message will indicate the recommended endTime to use to retrieve a complete data set. This behavior may occur if the provided endTime=currentTime, since we are not a real time system. When `endTime` is not provided, the API returns the latest data. This API also provides issue data. The `startTime` query param can be used to specify the beginning point of time range to retrieve the active issue counts in. When this param is not provided, the default `startTime` will be 24 hours before endTime. Aggregated response data will NOT have unique identifier data populated. List of unique identifier data: [`id`, `siteHierarchy`, `siteHierarchyId`, `siteType`, `latitude`, `longitude`]. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-siteHealthSummaries-1.0.3-resolved.yaml


@param ReadAnAggregatedSummaryOfSiteHealthDataV1HeaderParams Custom header parameters
@param ReadAnAggregatedSummaryOfSiteHealthDataV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!read-an-aggregated-summary-of-site-health-data-v1
*/
func (s *SitesService) ReadAnAggregatedSummaryOfSiteHealthDataV1(ReadAnAggregatedSummaryOfSiteHealthDataV1HeaderParams *ReadAnAggregatedSummaryOfSiteHealthDataV1HeaderParams, ReadAnAggregatedSummaryOfSiteHealthDataV1QueryParams *ReadAnAggregatedSummaryOfSiteHealthDataV1QueryParams) (*ResponseSitesReadAnAggregatedSummaryOfSiteHealthDataV1, *resty.Response, error) {
	path := "/dna/data/api/v1/siteHealthSummaries/summaryAnalytics"

	queryString, _ := query.Values(ReadAnAggregatedSummaryOfSiteHealthDataV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if ReadAnAggregatedSummaryOfSiteHealthDataV1HeaderParams != nil {

		if ReadAnAggregatedSummaryOfSiteHealthDataV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", ReadAnAggregatedSummaryOfSiteHealthDataV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesReadAnAggregatedSummaryOfSiteHealthDataV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReadAnAggregatedSummaryOfSiteHealthDataV1(ReadAnAggregatedSummaryOfSiteHealthDataV1HeaderParams, ReadAnAggregatedSummaryOfSiteHealthDataV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReadAnAggregatedSummaryOfSiteHealthDataV1")
	}

	result := response.Result().(*ResponseSitesReadAnAggregatedSummaryOfSiteHealthDataV1)
	return result, response, err

}

//ReadSiteHealthSummaryDataBySiteIDV1 Read site health summary data by site id. - 48aa-094f-423b-bd33
/* Get a health summary for a specific site by providing the unique site id in the url path. This API provides the latest health data from a given `endTime` If data is not ready for the provided endTime, the request will fail, and the error message will indicate the recommended endTime to use to retrieve a complete data set. This behavior may occur if the provided endTime=currentTime, since we are not a real time system. When `endTime` is not provided, the API returns the latest data. This API also provides issue data. The `startTime` query param can be used to specify the beginning point of time range to retrieve the active issue counts in. When this param is not provided, the default `startTime` will be 24 hours before endTime. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-siteHealthSummaries-1.0.3-resolved.yaml


@param id id path parameter. unique site uuid

@param ReadSiteHealthSummaryDataBySiteIdV1HeaderParams Custom header parameters
@param ReadSiteHealthSummaryDataBySiteIdV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!read-site-health-summary-data-by-site-id-v1
*/
func (s *SitesService) ReadSiteHealthSummaryDataBySiteIDV1(id string, ReadSiteHealthSummaryDataBySiteIdV1HeaderParams *ReadSiteHealthSummaryDataBySiteIDV1HeaderParams, ReadSiteHealthSummaryDataBySiteIdV1QueryParams *ReadSiteHealthSummaryDataBySiteIDV1QueryParams) (*ResponseSitesReadSiteHealthSummaryDataBySiteIDV1, *resty.Response, error) {
	path := "/dna/data/api/v1/siteHealthSummaries/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(ReadSiteHealthSummaryDataBySiteIdV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if ReadSiteHealthSummaryDataBySiteIdV1HeaderParams != nil {

		if ReadSiteHealthSummaryDataBySiteIdV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", ReadSiteHealthSummaryDataBySiteIdV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesReadSiteHealthSummaryDataBySiteIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReadSiteHealthSummaryDataBySiteIDV1(id, ReadSiteHealthSummaryDataBySiteIdV1HeaderParams, ReadSiteHealthSummaryDataBySiteIdV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReadSiteHealthSummaryDataBySiteIdV1")
	}

	result := response.Result().(*ResponseSitesReadSiteHealthSummaryDataBySiteIDV1)
	return result, response, err

}

//ImportMapArchiveImportStatusV1 Import Map Archive - Import Status - d9b0-599f-4d88-9a9a
/* Gets the status of a map archive import operation. For a map archive import that has just been initiated, will provide the result of validation of the archive and a pre-import preview of what will be performed if the import is performed.  Once an import is requested to be performed, this API will give the status of the import and upon completion a post-import summary of what was performed by the operation.


@param importContextUUID importContextUuid path parameter. The unique import context UUID given by a previous and recent call to maps/import/start API


Documentation Link: https://developer.cisco.com/docs/dna-center/#!import-map-archive-import-status-v1
*/
func (s *SitesService) ImportMapArchiveImportStatusV1(importContextUUID string) (*ResponseSitesImportMapArchiveImportStatusV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/maps/import/{importContextUuid}/status"
	path = strings.Replace(path, "{importContextUuid}", fmt.Sprintf("%v", importContextUUID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSitesImportMapArchiveImportStatusV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ImportMapArchiveImportStatusV1(importContextUUID)
		}
		return nil, response, fmt.Errorf("error with operation ImportMapArchiveImportStatusV1")
	}

	result := response.Result().(*ResponseSitesImportMapArchiveImportStatusV1)
	return result, response, err

}

//MapsSupportedAccessPointsV1 Maps Supported Access Points - 97b4-9a04-403b-bc05
/* Gets the list of supported access point types as well as valid antenna pattern names that can be used for each.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!maps-supported-access-points-v1
*/
func (s *SitesService) MapsSupportedAccessPointsV1() (*ResponseSitesMapsSupportedAccessPointsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/maps/supported-access-points"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSitesMapsSupportedAccessPointsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.MapsSupportedAccessPointsV1()
		}
		return nil, response, fmt.Errorf("error with operation MapsSupportedAccessPointsV1")
	}

	result := response.Result().(*ResponseSitesMapsSupportedAccessPointsV1)
	return result, response, err

}

//GetMembershipV1 Get Membership - eba6-6905-4e08-a60e
/* Getting the site children details and device details.


@param siteID siteId path parameter. Site id to retrieve device associated with the site.

@param GetMembershipV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-membership-v1
*/
func (s *SitesService) GetMembershipV1(siteID string, GetMembershipV1QueryParams *GetMembershipV1QueryParams) (*ResponseSitesGetMembershipV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/membership/{siteId}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	queryString, _ := query.Values(GetMembershipV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesGetMembershipV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetMembershipV1(siteID, GetMembershipV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetMembershipV1")
	}

	result := response.Result().(*ResponseSitesGetMembershipV1)
	return result, response, err

}

//GetSiteV1 Get Site - 6fb4-ab36-43fa-a80f
/* Get site(s) by site-name-hierarchy or siteId or type. List all sites if these parameters are not given as an input.


@param GetSiteV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-site-v1
*/
func (s *SitesService) GetSiteV1(GetSiteV1QueryParams *GetSiteV1QueryParams) (*ResponseSitesGetSiteV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/site"

	queryString, _ := query.Values(GetSiteV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesGetSiteV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSiteV1(GetSiteV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSiteV1")
	}

	result := response.Result().(*ResponseSitesGetSiteV1)
	return result, response, err

}

//GetSiteHealthV1 Get Site Health - 2597-2a31-43c8-8729
/* Returns Overall Health information for all sites


@param GetSiteHealthV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-site-health-v1
*/
func (s *SitesService) GetSiteHealthV1(GetSiteHealthV1QueryParams *GetSiteHealthV1QueryParams) (*ResponseSitesGetSiteHealthV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/site-health"

	queryString, _ := query.Values(GetSiteHealthV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesGetSiteHealthV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSiteHealthV1(GetSiteHealthV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSiteHealthV1")
	}

	result := response.Result().(*ResponseSitesGetSiteHealthV1)
	return result, response, err

}

//GetDevicesThatAreAssignedToASiteV1 Get devices that are assigned to a site - ae86-1be7-4d39-b0d1
/* API to get devices that are assigned to a site.


@param id id path parameter. Site Id

@param GetDevicesThatAreAssignedToASiteV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-devices-that-are-assigned-to-a-site-v1
*/
func (s *SitesService) GetDevicesThatAreAssignedToASiteV1(id string, GetDevicesThatAreAssignedToASiteV1QueryParams *GetDevicesThatAreAssignedToASiteV1QueryParams) (*ResponseSitesGetDevicesThatAreAssignedToASiteV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/site-member/{id}/member"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetDevicesThatAreAssignedToASiteV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesGetDevicesThatAreAssignedToASiteV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDevicesThatAreAssignedToASiteV1(id, GetDevicesThatAreAssignedToASiteV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDevicesThatAreAssignedToASiteV1")
	}

	result := response.Result().(*ResponseSitesGetDevicesThatAreAssignedToASiteV1)
	return result, response, err

}

//GetSiteCountV1 Get Site Count - b0b7-eabc-4f4b-9b28
/* Get the site count of the specified site's sub-hierarchy (inclusive of the provided site)


@param GetSiteCountV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-site-count-v1
*/
func (s *SitesService) GetSiteCountV1(GetSiteCountV1QueryParams *GetSiteCountV1QueryParams) (*ResponseSitesGetSiteCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/site/count"

	queryString, _ := query.Values(GetSiteCountV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesGetSiteCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSiteCountV1(GetSiteCountV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSiteCountV1")
	}

	result := response.Result().(*ResponseSitesGetSiteCountV1)
	return result, response, err

}

//GetSiteV2 Get Site V2 - 6490-7aa4-40c9-ba3a
/* API to get site(s) by site-name-hierarchy or siteId or type. List all sites if these parameters  are not given as an input.


@param GetSiteV2QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-site-v2
*/
func (s *SitesService) GetSiteV2(GetSiteV2QueryParams *GetSiteV2QueryParams) (*ResponseSitesGetSiteV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/site"

	queryString, _ := query.Values(GetSiteV2QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesGetSiteV2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSiteV2(GetSiteV2QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSiteV2")
	}

	result := response.Result().(*ResponseSitesGetSiteV2)
	return result, response, err

}

//GetSiteCountV2 Get Site Count V2 - 24a8-fb4c-4fbb-9a47
/* Get the site count of the specified site's sub-hierarchy (inclusive of the provided site)


@param GetSiteCountV2QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-site-count-v2
*/
func (s *SitesService) GetSiteCountV2(GetSiteCountV2QueryParams *GetSiteCountV2QueryParams) (*ResponseSitesGetSiteCountV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/site/count"

	queryString, _ := query.Values(GetSiteCountV2QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesGetSiteCountV2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSiteCountV2(GetSiteCountV2QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSiteCountV2")
	}

	result := response.Result().(*ResponseSitesGetSiteCountV2)
	return result, response, err

}

//QueryAnAggregatedSummaryOfSiteHealthDataV1 Query an aggregated summary of site health data. - 2782-ca59-4cc8-ad34
/* Query an aggregated summary of all site health This API provides the latest health data from a given `endTime` If data is not ready for the provided endTime, the request will fail, and the error message will indicate the recommended endTime to use to retrieve a complete data set. This behavior may occur if the provided endTime=currentTime, since we are not a real time system. When `endTime` is not provided, the API returns the latest data. This API also provides issue data. The `startTime` query param can be used to specify the beginning point of time range to retrieve the active issue counts in. When this param is not provided, the default `startTime` will be 24 hours before endTime.

 Aggregated response data will NOT have unique identifier data populated.

 List of unique identifier data: [`id`, `siteHierarchy`,
`siteHierarchyId`, `siteType`, `latitude`, `longitude`] Please refer to the 'API Support Documentation' section to understand which fields are supported. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-siteHealthSummaries-1.0.3-resolved.yaml


@param QueryAnAggregatedSummaryOfSiteHealthDataV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!query-an-aggregated-summary-of-site-health-data-v1
*/
func (s *SitesService) QueryAnAggregatedSummaryOfSiteHealthDataV1(requestSitesQueryAnAggregatedSummaryOfSiteHealthDataV1 *RequestSitesQueryAnAggregatedSummaryOfSiteHealthDataV1, QueryAnAggregatedSummaryOfSiteHealthDataV1QueryParams *QueryAnAggregatedSummaryOfSiteHealthDataV1QueryParams) (*ResponseSitesQueryAnAggregatedSummaryOfSiteHealthDataV1, *resty.Response, error) {
	path := "/dna/data/api/v1/siteHealthSummaries/summaryAnalytics"

	queryString, _ := query.Values(QueryAnAggregatedSummaryOfSiteHealthDataV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetBody(requestSitesQueryAnAggregatedSummaryOfSiteHealthDataV1).
		SetResult(&ResponseSitesQueryAnAggregatedSummaryOfSiteHealthDataV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.QueryAnAggregatedSummaryOfSiteHealthDataV1(requestSitesQueryAnAggregatedSummaryOfSiteHealthDataV1, QueryAnAggregatedSummaryOfSiteHealthDataV1QueryParams)
		}

		return nil, response, fmt.Errorf("error with operation QueryAnAggregatedSummaryOfSiteHealthDataV1")
	}

	result := response.Result().(*ResponseSitesQueryAnAggregatedSummaryOfSiteHealthDataV1)
	return result, response, err

}

//AssignDevicesToSiteV1 Assign Devices To Site - 98a8-aa5e-40cb-b90b
/* Assigns unassigned devices to a site. This API does not move assigned devices to other sites.


@param siteID siteId path parameter. Site Id where device(s) needs to be assigned

@param AssignDevicesToSiteV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!assign-devices-to-site-v1
*/
func (s *SitesService) AssignDevicesToSiteV1(siteID string, requestSitesAssignDevicesToSiteV1 *RequestSitesAssignDevicesToSiteV1, AssignDevicesToSiteV1HeaderParams *AssignDevicesToSiteV1HeaderParams) (*ResponseSitesAssignDevicesToSiteV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/assign-device-to-site/{siteId}/device"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if AssignDevicesToSiteV1HeaderParams != nil {

		if AssignDevicesToSiteV1HeaderParams.Runsync != "" {
			clientRequest = clientRequest.SetHeader("__runsync", AssignDevicesToSiteV1HeaderParams.Runsync)
		}

		if AssignDevicesToSiteV1HeaderParams.Timeout != "" {
			clientRequest = clientRequest.SetHeader("__timeout", AssignDevicesToSiteV1HeaderParams.Timeout)
		}

		if AssignDevicesToSiteV1HeaderParams.Persistbapioutput != "" {
			clientRequest = clientRequest.SetHeader("__persistbapioutput", AssignDevicesToSiteV1HeaderParams.Persistbapioutput)
		}

	}

	response, err = clientRequest.
		SetBody(requestSitesAssignDevicesToSiteV1).
		SetResult(&ResponseSitesAssignDevicesToSiteV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AssignDevicesToSiteV1(siteID, requestSitesAssignDevicesToSiteV1, AssignDevicesToSiteV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation AssignDevicesToSiteV1")
	}

	result := response.Result().(*ResponseSitesAssignDevicesToSiteV1)
	return result, response, err

}

//ExportMapArchiveV1 Export Map Archive - fb9f-2948-493b-bc68
/* Allows exporting a Map archive in an XML interchange format along with the associated images.


@param siteHierarchyUUID siteHierarchyUuid path parameter. The site hierarchy element UUID to export, all child elements starting at this hierarchy element will be included. Limited to a hierarchy that contains 500 or fewer maps.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!export-map-archive-v1
*/
func (s *SitesService) ExportMapArchiveV1(siteHierarchyUUID string) (*ResponseSitesExportMapArchiveV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/maps/export/{siteHierarchyUuid}"
	path = strings.Replace(path, "{siteHierarchyUuid}", fmt.Sprintf("%v", siteHierarchyUUID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSitesExportMapArchiveV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ExportMapArchiveV1(siteHierarchyUUID)
		}

		return nil, response, fmt.Errorf("error with operation ExportMapArchiveV1")
	}

	result := response.Result().(*ResponseSitesExportMapArchiveV1)
	return result, response, err

}

//ImportMapArchiveStartImportV1 Import Map Archive - Start Import - b485-8b4e-4f6b-8409
/* Initiates a map archive import of a tar.gz file.  The archive must consist of one xmlDir/MapsImportExport.xml map descriptor file, and 1 or more images for the map areas nested under /images folder.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!import-map-archive-start-import-v1
*/
func (s *SitesService) ImportMapArchiveStartImportV1() (*resty.Response, error) {
	path := "/dna/intent/api/v1/maps/import/start"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ImportMapArchiveStartImportV1()
		}

		return response, fmt.Errorf("error with operation ImportMapArchiveStartImportV1")
	}

	return response, err

}

//ImportMapArchivePerformImportV1 Import Map Archive - Perform Import - 5a9d-db37-4c18-9f6f
/* For a previously initatied import, approves the import to be performed, accepting that data loss may occur.  A Map import will fully replace existing Maps data for the site(s) defined in the archive. The Map Archive Import Status API /maps/import/${contextUuid}/status should always be checked to validate the pre-import validation output prior to performing the import.


@param importContextUUID importContextUuid path parameter. The unique import context UUID given by a previous call of Start Import API


Documentation Link: https://developer.cisco.com/docs/dna-center/#!import-map-archive-perform-import-v1
*/
func (s *SitesService) ImportMapArchivePerformImportV1(importContextUUID string) (*ResponseSitesImportMapArchivePerformImportV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/maps/import/{importContextUuid}/perform"
	path = strings.Replace(path, "{importContextUuid}", fmt.Sprintf("%v", importContextUUID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").

		// SetResult(&ResponseSitesImportMapArchivePerformImportV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ImportMapArchivePerformImportV1(importContextUUID)
		}

		return nil, response, fmt.Errorf("error with operation ImportMapArchivePerformImportV1")
	}

	result := response.Result().(ResponseSitesImportMapArchivePerformImportV1)

	return &result, response, err

}

//CreateSiteV1 Create Site - 50b5-89fd-4c7a-930a
/* Creates site with area/building/floor with specified hierarchy.


@param CreateSiteV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-site-v1
*/
func (s *SitesService) CreateSiteV1(requestSitesCreateSiteV1 *RequestSitesCreateSiteV1, CreateSiteV1HeaderParams *CreateSiteV1HeaderParams) (*ResponseSitesCreateSiteV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/site"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if CreateSiteV1HeaderParams != nil {

		if CreateSiteV1HeaderParams.Runsync != "" {
			clientRequest = clientRequest.SetHeader("__runsync", CreateSiteV1HeaderParams.Runsync)
		}

		if CreateSiteV1HeaderParams.Timeout != "" {
			clientRequest = clientRequest.SetHeader("__timeout", CreateSiteV1HeaderParams.Timeout)
		}

		if CreateSiteV1HeaderParams.Persistbapioutput != "" {
			clientRequest = clientRequest.SetHeader("__persistbapioutput", CreateSiteV1HeaderParams.Persistbapioutput)
		}

	}

	response, err = clientRequest.
		SetBody(requestSitesCreateSiteV1).
		SetResult(&ResponseSitesCreateSiteV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateSiteV1(requestSitesCreateSiteV1, CreateSiteV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation CreateSiteV1")
	}

	result := response.Result().(*ResponseSitesCreateSiteV1)
	return result, response, err

}

//UpdateSiteV1 Update Site - eeb7-eb4b-4bd8-a1dd
/* Update site area/building/floor with specified hierarchy and new values


@param siteID siteId path parameter. Site id to which site details to be updated.

@param UpdateSiteV1HeaderParams Custom header parameters
*/
func (s *SitesService) UpdateSiteV1(siteID string, requestSitesUpdateSiteV1 *RequestSitesUpdateSiteV1, UpdateSiteV1HeaderParams *UpdateSiteV1HeaderParams) (*ResponseSitesUpdateSiteV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/site/{siteId}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if UpdateSiteV1HeaderParams != nil {

		if UpdateSiteV1HeaderParams.Runsync != "" {
			clientRequest = clientRequest.SetHeader("__runsync", UpdateSiteV1HeaderParams.Runsync)
		}

		if UpdateSiteV1HeaderParams.Timeout != "" {
			clientRequest = clientRequest.SetHeader("__timeout", UpdateSiteV1HeaderParams.Timeout)
		}

		if UpdateSiteV1HeaderParams.Persistbapioutput != "" {
			clientRequest = clientRequest.SetHeader("__persistbapioutput", UpdateSiteV1HeaderParams.Persistbapioutput)
		}

	}

	response, err = clientRequest.
		SetBody(requestSitesUpdateSiteV1).
		SetResult(&ResponseSitesUpdateSiteV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateSiteV1(siteID, requestSitesUpdateSiteV1, UpdateSiteV1HeaderParams)
		}
		return nil, response, fmt.Errorf("error with operation UpdateSiteV1")
	}

	result := response.Result().(*ResponseSitesUpdateSiteV1)
	return result, response, err

}

//ImportMapArchiveCancelAnImportV1 Import Map Archive - Cancel an Import - 52b0-d8a3-4159-a6b7
/* Cancels a previously initatied import, allowing the system to cleanup cached resources about that import data, and ensures the import cannot accidentally be performed / approved at a later time.


@param importContextUUID importContextUuid path parameter. The unique import context UUID given by a previous call to Start Import API


Documentation Link: https://developer.cisco.com/docs/dna-center/#!import-map-archive-cancel-an-import-v1
*/
func (s *SitesService) ImportMapArchiveCancelAnImportV1(importContextUUID string) (*ResponseSitesImportMapArchiveCancelAnImportV1, *resty.Response, error) {
	//importContextUUID string
	path := "/dna/intent/api/v1/maps/import/{importContextUuid}"
	path = strings.Replace(path, "{importContextUuid}", fmt.Sprintf("%v", importContextUUID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		// SetResult(&ResponseSitesImportMapArchiveCancelAnImportV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ImportMapArchiveCancelAnImportV1(
				importContextUUID)
		}
		return nil, response, fmt.Errorf("error with operation ImportMapArchiveCancelAnImportV1")
	}

	result := response.Result().(ResponseSitesImportMapArchiveCancelAnImportV1)

	return &result, response, err

}

//DeleteSiteV1 Delete Site - f083-cb13-484a-8fae
/* Delete site with area/building/floor by siteId.


@param siteID siteId path parameter. Site id to which site details to be deleted.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-site-v1
*/
func (s *SitesService) DeleteSiteV1(siteID string) (*ResponseSitesDeleteSiteV1, *resty.Response, error) {
	//siteID string
	path := "/dna/intent/api/v1/site/{siteId}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSitesDeleteSiteV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteSiteV1(siteID)
		}
		return nil, response, fmt.Errorf("error with operation DeleteSiteV1")
	}

	result := response.Result().(*ResponseSitesDeleteSiteV1)
	return result, response, err

}

// Alias Function
func (s *SitesService) GetSiteHealth(GetSiteHealthV1QueryParams *GetSiteHealthV1QueryParams) (*ResponseSitesGetSiteHealthV1, *resty.Response, error) {
	return s.GetSiteHealthV1(GetSiteHealthV1QueryParams)
}

// Alias Function
func (s *SitesService) CreateSite(requestSitesCreateSiteV1 *RequestSitesCreateSiteV1, CreateSiteV1HeaderParams *CreateSiteV1HeaderParams) (*ResponseSitesCreateSiteV1, *resty.Response, error) {
	return s.CreateSiteV1(requestSitesCreateSiteV1, CreateSiteV1HeaderParams)
}

// Alias Function
func (s *SitesService) ImportMapArchivePerformImport(importContextUUID string) (*ResponseSitesImportMapArchivePerformImportV1, *resty.Response, error) {
	return s.ImportMapArchivePerformImportV1(importContextUUID)
}

// Alias Function
func (s *SitesService) ImportMapArchiveImportStatus(importContextUUID string) (*ResponseSitesImportMapArchiveImportStatusV1, *resty.Response, error) {
	return s.ImportMapArchiveImportStatusV1(importContextUUID)
}

// Alias Function
func (s *SitesService) ImportMapArchiveCancelAnImport(importContextUUID string) (*ResponseSitesImportMapArchiveCancelAnImportV1, *resty.Response, error) {
	return s.ImportMapArchiveCancelAnImportV1(importContextUUID)
}

// Alias Function
func (s *SitesService) GetSiteCount(GetSiteCountV1QueryParams *GetSiteCountV1QueryParams) (*ResponseSitesGetSiteCountV1, *resty.Response, error) {
	return s.GetSiteCountV1(GetSiteCountV1QueryParams)
}

// Alias Function
func (s *SitesService) ReadSiteCount(ReadSiteCountV1HeaderParams *ReadSiteCountV1HeaderParams, ReadSiteCountV1QueryParams *ReadSiteCountV1QueryParams) (*ResponseSitesReadSiteCountV1, *resty.Response, error) {
	return s.ReadSiteCountV1(ReadSiteCountV1HeaderParams, ReadSiteCountV1QueryParams)
}

// Alias Function
func (s *SitesService) GetMembership(siteID string, GetMembershipV1QueryParams *GetMembershipV1QueryParams) (*ResponseSitesGetMembershipV1, *resty.Response, error) {
	return s.GetMembershipV1(siteID, GetMembershipV1QueryParams)
}

// Alias Function
func (s *SitesService) ReadListOfSiteHealthSummaries(ReadListOfSiteHealthSummariesV1HeaderParams *ReadListOfSiteHealthSummariesV1HeaderParams, ReadListOfSiteHealthSummariesV1QueryParams *ReadListOfSiteHealthSummariesV1QueryParams) (*ResponseSitesReadListOfSiteHealthSummariesV1, *resty.Response, error) {
	return s.ReadListOfSiteHealthSummariesV1(ReadListOfSiteHealthSummariesV1HeaderParams, ReadListOfSiteHealthSummariesV1QueryParams)
}

// Alias Function
func (s *SitesService) GetSite(GetSiteV1QueryParams *GetSiteV1QueryParams) (*ResponseSitesGetSiteV1, *resty.Response, error) {
	return s.GetSiteV1(GetSiteV1QueryParams)
}

// Alias Function
func (s *SitesService) ReadAnAggregatedSummaryOfSiteHealthData(ReadAnAggregatedSummaryOfSiteHealthDataV1HeaderParams *ReadAnAggregatedSummaryOfSiteHealthDataV1HeaderParams, ReadAnAggregatedSummaryOfSiteHealthDataV1QueryParams *ReadAnAggregatedSummaryOfSiteHealthDataV1QueryParams) (*ResponseSitesReadAnAggregatedSummaryOfSiteHealthDataV1, *resty.Response, error) {
	return s.ReadAnAggregatedSummaryOfSiteHealthDataV1(ReadAnAggregatedSummaryOfSiteHealthDataV1HeaderParams, ReadAnAggregatedSummaryOfSiteHealthDataV1QueryParams)
}

// Alias Function
func (s *SitesService) AssignDevicesToSite(siteID string, requestSitesAssignDevicesToSiteV1 *RequestSitesAssignDevicesToSiteV1, AssignDevicesToSiteV1HeaderParams *AssignDevicesToSiteV1HeaderParams) (*ResponseSitesAssignDevicesToSiteV1, *resty.Response, error) {
	return s.AssignDevicesToSiteV1(siteID, requestSitesAssignDevicesToSiteV1, AssignDevicesToSiteV1HeaderParams)
}

// Alias Function
func (s *SitesService) ExportMapArchive(siteHierarchyUUID string) (*ResponseSitesExportMapArchiveV1, *resty.Response, error) {
	return s.ExportMapArchiveV1(siteHierarchyUUID)
}

// Alias Function
func (s *SitesService) UpdateSite(siteID string, requestSitesUpdateSiteV1 *RequestSitesUpdateSiteV1, UpdateSiteV1HeaderParams *UpdateSiteV1HeaderParams) (*ResponseSitesUpdateSiteV1, *resty.Response, error) {
	return s.UpdateSiteV1(siteID, requestSitesUpdateSiteV1, UpdateSiteV1HeaderParams)
}

// Alias Function
func (s *SitesService) GetDevicesThatAreAssignedToASite(id string, GetDevicesThatAreAssignedToASiteV1QueryParams *GetDevicesThatAreAssignedToASiteV1QueryParams) (*ResponseSitesGetDevicesThatAreAssignedToASiteV1, *resty.Response, error) {
	return s.GetDevicesThatAreAssignedToASiteV1(id, GetDevicesThatAreAssignedToASiteV1QueryParams)
}

// Alias Function
func (s *SitesService) QueryAnAggregatedSummaryOfSiteHealthData(requestSitesQueryAnAggregatedSummaryOfSiteHealthDataV1 *RequestSitesQueryAnAggregatedSummaryOfSiteHealthDataV1, QueryAnAggregatedSummaryOfSiteHealthDataV1QueryParams *QueryAnAggregatedSummaryOfSiteHealthDataV1QueryParams) (*ResponseSitesQueryAnAggregatedSummaryOfSiteHealthDataV1, *resty.Response, error) {
	return s.QueryAnAggregatedSummaryOfSiteHealthDataV1(requestSitesQueryAnAggregatedSummaryOfSiteHealthDataV1, QueryAnAggregatedSummaryOfSiteHealthDataV1QueryParams)
}

// Alias Function
func (s *SitesService) ImportMapArchiveStartImport() (*resty.Response, error) {
	return s.ImportMapArchiveStartImportV1()
}

// Alias Function
func (s *SitesService) DeleteSite(siteID string) (*ResponseSitesDeleteSiteV1, *resty.Response, error) {
	return s.DeleteSiteV1(siteID)
}

// Alias Function
func (s *SitesService) MapsSupportedAccessPoints() (*ResponseSitesMapsSupportedAccessPointsV1, *resty.Response, error) {
	return s.MapsSupportedAccessPointsV1()
}

// Alias Function
func (s *SitesService) ReadSiteHealthSummaryDataBySiteID(id string, ReadSiteHealthSummaryDataBySiteIdV1HeaderParams *ReadSiteHealthSummaryDataBySiteIDV1HeaderParams, ReadSiteHealthSummaryDataBySiteIdV1QueryParams *ReadSiteHealthSummaryDataBySiteIDV1QueryParams) (*ResponseSitesReadSiteHealthSummaryDataBySiteIDV1, *resty.Response, error) {
	return s.ReadSiteHealthSummaryDataBySiteIDV1(id, ReadSiteHealthSummaryDataBySiteIdV1HeaderParams, ReadSiteHealthSummaryDataBySiteIdV1QueryParams)
}
