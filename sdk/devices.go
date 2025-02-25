package catalyst

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type DevicesService service

type RetrievesTheListOfAAAServicesForGivenParametersV1QueryParams struct {
	StartTime             float64 `url:"startTime,omitempty"`             //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to current time.
	EndTime               float64 `url:"endTime,omitempty"`               //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	Limit                 float64 `url:"limit,omitempty"`                 //Maximum number of records to return
	Offset                float64 `url:"offset,omitempty"`                //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	SortBy                string  `url:"sortBy,omitempty"`                //Field name on which sorting needs to be done
	Order                 string  `url:"order,omitempty"`                 //The sort order of the field ascending or descending.
	ServerIP              string  `url:"serverIp,omitempty"`              //IP Address of the AAA Server. This parameter supports wildcard (`*`) character -based search. Example: `10.76.81.*` or `*56.78*` or `*50.28` Examples: serverIp=10.42.3.31 (single IP Address is requested) serverIp=10.42.3.31&serverIp=name2&fabricVnName=name3 (multiple IP Addresses are requested)
	DeviceID              string  `url:"deviceId,omitempty"`              //The device UUID.   Examples:  `deviceId=6bef213c-19ca-4170-8375-b694e251101c` (single deviceId is requested)  `deviceId=6bef213c-19ca-4170-8375-b694e251101c&deviceId=32219612-819e-4b5e-a96b-cf22aca13dd9 (multiple networkDeviceIds with & separator)
	DeviceName            string  `url:"deviceName,omitempty"`            //Name of the device. This parameter supports wildcard (`*`) character -based search. Example: `wnbu-sjc*` or `*wnbu-sjc*` or `*wnbu-sjc` Examples: deviceName=wnbu-sjc24.cisco.com (single device name is requested) deviceName=wnbu-sjc24.cisco.com&deviceName=wnbu-sjc22.cisco.com (multiple device names are requested)
	SiteHierarchy         string  `url:"siteHierarchy,omitempty"`         //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. `Global/AreaName/BuildingName/FloorName`) This field supports wildcard asterisk (`*`) character search support. E.g. `*/San*, */San, /San*` Examples: `?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)
	DeviceSiteHierarchyID string  `url:"deviceSiteHierarchyId,omitempty"` //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*` Examples: `?deviceSiteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?deviceSiteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&deviceSiteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	SiteID                string  `url:"siteId,omitempty"`                //The UUID of the site. (Ex. `flooruuid`) Examples: `?siteId=id1` (single id requested) `?siteId=id1&siteId=id2&siteId=id3` (multiple ids requested)
}
type RetrievesTheListOfAAAServicesForGivenParametersV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheTotalNumberOfAAAServicesForGivenParametersV1QueryParams struct {
	StartTime             float64 `url:"startTime,omitempty"`             //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime               float64 `url:"endTime,omitempty"`               //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	ServerIP              string  `url:"serverIp,omitempty"`              //IP Address of the AAA Server. This parameter supports wildcard (`*`) character -based search. Example: `10.76.81.*` or `*56.78*` or `*50.28` Examples: serverIp=10.42.3.31 (single IP Address is requested) serverIp=10.42.3.31&serverIp=name2&fabricVnName=name3 (multiple IP Addresses are requested)
	DeviceID              string  `url:"deviceId,omitempty"`              //The device UUID.   Examples:  `deviceId=6bef213c-19ca-4170-8375-b694e251101c` (single deviceId is requested)  `deviceId=6bef213c-19ca-4170-8375-b694e251101c&deviceId=32219612-819e-4b5e-a96b-cf22aca13dd9 (multiple networkDeviceIds with & separator)
	DeviceName            string  `url:"deviceName,omitempty"`            //Name of the device. This parameter supports wildcard (`*`) character -based search. Example: `wnbu-sjc*` or `*wnbu-sjc*` or `*wnbu-sjc` Examples: deviceName=wnbu-sjc24.cisco.com (single device name is requested) deviceName=wnbu-sjc24.cisco.com&deviceName=wnbu-sjc22.cisco.com (multiple device names are requested)
	DeviceSiteHierarchy   string  `url:"deviceSiteHierarchy,omitempty"`   //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. `Global/AreaName/BuildingName/FloorName`) This field supports wildcard asterisk (`*`) character search support. E.g. `*/San*, */San, /San*` Examples: `?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `?deviceSiteHierarchy=Global/AreaName/BuildingName/FloorName&deviceSiteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)
	DeviceSiteHierarchyID string  `url:"deviceSiteHierarchyId,omitempty"` //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*` Examples: `?deviceSiteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?deviceSiteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&deviceSiteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	DeviceSiteID          string  `url:"deviceSiteId,omitempty"`          //The UUID of the site. (Ex. `flooruuid`) Examples: `?deviceSiteIds=id1` (single id requested) `?deviceSiteIds=id1&deviceSiteIds=id2&siteId=id3` (multiple ids requested)
}
type RetrievesTheTotalNumberOfAAAServicesForGivenParametersV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIDOfTheServiceV1QueryParams struct {
	StartTime float64 `url:"startTime,omitempty"` //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime   float64 `url:"endTime,omitempty"`   //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
}
type RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIDOfTheServiceV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1HeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type QueryAssuranceEventsV1QueryParams struct {
	DeviceFamily      string  `url:"deviceFamily,omitempty"`      //Device family. Please note that multiple families across network device type and client type is not allowed. For example, choosing `Routers` along with `Wireless Client` or `Unified AP` is not supported. Examples: `deviceFamily=Switches and Hubs` (single deviceFamily requested) `deviceFamily=Switches and Hubs&deviceFamily=Routers` (multiple deviceFamily requested)
	StartTime         float64 `url:"startTime,omitempty"`         //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to current time minus 24 hours.
	EndTime           float64 `url:"endTime,omitempty"`           //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `endTime` is not provided, API will default to current time.
	MessageType       string  `url:"messageType,omitempty"`       //Message type for the event. Examples: `messageType=Syslog` (single messageType requested) `messageType=Trap&messageType=Syslog` (multiple messageType requested)
	Severity          float64 `url:"severity,omitempty"`          //Severity of the event between 0 and 6. This is applicable only for events related to network devices (other than AP) and `Wired Client` events. | Value | Severity    | | ----- | ----------- | | 0     | Emergency   | | 1     | Alert       | | 2     | Critical    | | 3     | Error       | | 4     | Warning     | | 5     | Notice      | | 6     | Info        | Examples: `severity=0` (single severity requested) `severity=0&severity=1` (multiple severity requested)
	SiteID            string  `url:"siteId,omitempty"`            //The UUID of the site. (Ex. `flooruuid`) Examples: `?siteId=id1` (single siteId requested) `?siteId=id1&siteId=id2&siteId=id3` (multiple siteId requested)
	SiteHierarchyID   string  `url:"siteHierarchyId,omitempty"`   //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*` Examples: `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyId requested)
	NetworkDeviceName string  `url:"networkDeviceName,omitempty"` //Network device name. This parameter is applicable for network device related families. This field supports wildcard (`*`) character-based search. Ex: `*Branch*` or `Branch*` or `*Branch` Examples: `networkDeviceName=Branch-3-Gateway` (single networkDeviceName requested) `networkDeviceName=Branch-3-Gateway&networkDeviceName=Branch-3-Switch` (multiple networkDeviceName requested)
	NetworkDeviceID   string  `url:"networkDeviceId,omitempty"`   //The list of Network Device Uuids. (Ex. `6bef213c-19ca-4170-8375-b694e251101c`) Examples: `networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c` (single networkDeviceId requested) `networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c&networkDeviceId=32219612-819e-4b5e-a96b-cf22aca13dd9&networkDeviceId=2541e9a7-b80d-4955-8aa2-79b233318ba0` (multiple networkDeviceId with & separator)
	ApMac             string  `url:"apMac,omitempty"`             //MAC address of the access point. This parameter is applicable for `Unified AP` and `Wireless Client` events. This field supports wildcard (`*`) character-based search. Ex: `*50:0F*` or `50:0F*` or `*50:0F` Examples: `apMac=50:0F:80:0F:F7:E0` (single apMac requested) `apMac=50:0F:80:0F:F7:E0&apMac=18:80:90:AB:7E:A0` (multiple apMac requested)
	ClientMac         string  `url:"clientMac,omitempty"`         //MAC address of the client. This parameter is applicable for `Wired Client` and `Wireless Client` events. This field supports wildcard (`*`) character-based search. Ex: `*66:2B*` or `66:2B*` or `*66:2B` Examples: `clientMac=66:2B:B8:D2:01:56` (single clientMac requested) `clientMac=66:2B:B8:D2:01:56&clientMac=DC:A6:32:F5:5A:89` (multiple clientMac requested)
	Attribute         string  `url:"attribute,omitempty"`         //The list of attributes that needs to be included in the response. If this parameter is not provided, then basic attributes (`id`, `name`, `timestamp`, `details`, `messageType`, `siteHierarchyId`, `siteHierarchy`, `deviceFamily`, `networkDeviceId`, `networkDeviceName`, `managementIpAddress`) would be part of the response.  Examples:  `attribute=name` (single attribute requested) `attribute=name&attribute=networkDeviceName` (multiple attribute requested)
	View              string  `url:"view,omitempty"`              //The list of events views. Please refer to `EventViews` for the supported list  Examples:  `view=network` (single view requested) `view=network&view=ap` (multiple view requested)
	Offset            float64 `url:"offset,omitempty"`            //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	Limit             float64 `url:"limit,omitempty"`             //Maximum number of records to return
	SortBy            string  `url:"sortBy,omitempty"`            //A field within the response to sort by.
	Order             string  `url:"order,omitempty"`             //The sort order of the field ascending or descending.
}
type QueryAssuranceEventsV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type CountTheNumberOfEventsV1QueryParams struct {
	DeviceFamily      string `url:"deviceFamily,omitempty"`      //Device family. Please note that multiple families across network device type and client type is not allowed. For example, choosing `Routers` along with `Wireless Client` or `Unified AP` is not supported. Examples: `deviceFamily=Switches and Hubs` (single deviceFamily requested) `deviceFamily=Switches and Hubs&deviceFamily=Routers` (multiple deviceFamily requested)
	StartTime         string `url:"startTime,omitempty"`         //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to current time minus 24 hours.
	EndTime           string `url:"endTime,omitempty"`           //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `endTime` is not provided, API will default to current time.
	MessageType       string `url:"messageType,omitempty"`       //Message type for the event. Examples: `messageType=Syslog` (single messageType requested) `messageType=Trap&messageType=Syslog` (multiple messageType requested)
	Severity          string `url:"severity,omitempty"`          //Severity of the event between 0 and 6. This is applicable only for events related to network devices (other than AP) and `Wired Client` events. | Value | Severity    | | ----- | ----------- | | 0     | Emergency   | | 1     | Alert       | | 2     | Critical    | | 3     | Error       | | 4     | Warning     | | 5     | Notice      | | 6     | Info        | Examples: `severity=0` (single severity requested) `severity=0&severity=1` (multiple severity requested)
	SiteID            string `url:"siteId,omitempty"`            //The UUID of the site. (Ex. `flooruuid`) Examples: `?siteId=id1` (single siteId requested) `?siteId=id1&siteId=id2&siteId=id3` (multiple siteId requested)
	SiteHierarchyID   string `url:"siteHierarchyId,omitempty"`   //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*` Examples: `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyId requested)
	NetworkDeviceName string `url:"networkDeviceName,omitempty"` //Network device name. This parameter is applicable for network device related families. This field supports wildcard (`*`) character-based search. Ex: `*Branch*` or `Branch*` or `*Branch` Examples: `networkDeviceName=Branch-3-Gateway` (single networkDeviceName requested) `networkDeviceName=Branch-3-Gateway&networkDeviceName=Branch-3-Switch` (multiple networkDeviceName requested)
	NetworkDeviceID   string `url:"networkDeviceId,omitempty"`   //The list of Network Device Uuids. (Ex. `6bef213c-19ca-4170-8375-b694e251101c`) Examples: `networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c` (single networkDeviceId requested) `networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c&networkDeviceId=32219612-819e-4b5e-a96b-cf22aca13dd9&networkDeviceId=2541e9a7-b80d-4955-8aa2-79b233318ba0` (multiple networkDeviceId requested)
	ApMac             string `url:"apMac,omitempty"`             //MAC address of the access point. This parameter is applicable for `Unified AP` and `Wireless Client` events. This field supports wildcard (`*`) character-based search. Ex: `*50:0F*` or `50:0F*` or `*50:0F` Examples: `apMac=50:0F:80:0F:F7:E0` (single apMac requested) `apMac=50:0F:80:0F:F7:E0&apMac=18:80:90:AB:7E:A0` (multiple apMac requested)
	ClientMac         string `url:"clientMac,omitempty"`         //MAC address of the client. This parameter is applicable for `Wired Client` and `Wireless Client` events. This field supports wildcard (`*`) character-based search. Ex: `*66:2B*` or `66:2B*` or `*66:2B` Examples: `clientMac=66:2B:B8:D2:01:56` (single clientMac requested) `clientMac=66:2B:B8:D2:01:56&clientMac=DC:A6:32:F5:5A:89` (multiple clientMac requested)
}
type CountTheNumberOfEventsV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type QueryAssuranceEventsWithFiltersV1HeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type CountTheNumberOfEventsWithFiltersV1HeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetDetailsOfASingleAssuranceEventV1QueryParams struct {
	Attribute string `url:"attribute,omitempty"` //The list of attributes that needs to be included in the response. If this parameter is not provided, then basic attributes (`id`, `name`, `timestamp`, `details`, `messageType`, `siteHierarchyId`, `siteHierarchy`, `deviceFamily`, `networkDeviceId`, `networkDeviceName`, `managementIpAddress`) would be part of the response.  Examples:  `attribute=name` (single attribute requested) `attribute=name&attribute=networkDeviceName` (multiple attribute requested)
	View      string `url:"view,omitempty"`      //The list of events views. Please refer to `EventViews` for the supported list  Examples:  `view=network` (single view requested) `view=network&view=ap` (multiple view requested)
}
type GetDetailsOfASingleAssuranceEventV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetListOfChildEventsForTheGivenWirelessClientEventV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheListOfDHCPServicesForGivenParametersV1QueryParams struct {
	StartTime             float64 `url:"startTime,omitempty"`             //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime               float64 `url:"endTime,omitempty"`               //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	Limit                 float64 `url:"limit,omitempty"`                 //Maximum number of records to return
	Offset                float64 `url:"offset,omitempty"`                //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	SortBy                string  `url:"sortBy,omitempty"`                //Field name on which sorting needs to be done
	Order                 string  `url:"order,omitempty"`                 //The sort order of the field ascending or descending.
	ServerIP              string  `url:"serverIp,omitempty"`              //IP Address of the DHCP Server. This parameter supports wildcard (`*`) character -based search. Example: `10.76.81.*` or `*56.78*` or `*50.28` Examples: serverIp=10.42.3.31 (single IP Address is requested) serverIp=10.42.3.31&serverIp=name2&fabricVnName=name3 (multiple IP Addresses are requested)
	DeviceID              string  `url:"deviceId,omitempty"`              //The device UUID.   Examples:  `deviceId=6bef213c-19ca-4170-8375-b694e251101c` (single deviceId is requested)  `deviceId=6bef213c-19ca-4170-8375-b694e251101c&deviceId=32219612-819e-4b5e-a96b-cf22aca13dd9 (multiple networkDeviceIds with & separator)
	DeviceName            string  `url:"deviceName,omitempty"`            //Name of the device. This parameter supports wildcard (`*`) character -based search. Example: `wnbu-sjc*` or `*wnbu-sjc*` or `*wnbu-sjc` Examples: deviceName=wnbu-sjc24.cisco.com (single device name is requested) deviceName=wnbu-sjc24.cisco.com&deviceName=wnbu-sjc22.cisco.com (multiple device names are requested)
	DeviceSiteHierarchy   string  `url:"deviceSiteHierarchy,omitempty"`   //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. `Global/AreaName/BuildingName/FloorName`) This field supports wildcard asterisk (`*`) character search support. E.g. `*/San*, */San, /San*` Examples: `?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `?deviceSiteHierarchy=Global/AreaName/BuildingName/FloorName&deviceSiteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)
	DeviceSiteHierarchyID string  `url:"deviceSiteHierarchyId,omitempty"` //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*` Examples: `?deviceSiteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?deviceSiteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&deviceSiteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	DeviceSiteID          string  `url:"deviceSiteId,omitempty"`          //The UUID of the site. (Ex. `flooruuid`) Examples: `?deviceSiteIds=id1` (single id requested) `?deviceSiteIds=id1&deviceSiteIds=id2&siteId=id3` (multiple ids requested)
}
type RetrievesTheListOfDHCPServicesForGivenParametersV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheTotalNumberOfDHCPServicesForGivenParametersV1QueryParams struct {
	StartTime             float64 `url:"startTime,omitempty"`             //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime               float64 `url:"endTime,omitempty"`               //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	ServerIP              string  `url:"serverIp,omitempty"`              //IP Address of the DHCP Server. This parameter supports wildcard (`*`) character -based search. Example: `10.76.81.*` or `*56.78*` or `*50.28` Examples: serverIp=10.42.3.31 (single IP Address is requested) serverIp=10.42.3.31&serverIp=name2&fabricVnName=name3 (multiple IP Addresses are requested)
	DeviceID              string  `url:"deviceId,omitempty"`              //The device UUID.   Examples:  `deviceId=6bef213c-19ca-4170-8375-b694e251101c` (single deviceId is requested)  `deviceId=6bef213c-19ca-4170-8375-b694e251101c&deviceId=32219612-819e-4b5e-a96b-cf22aca13dd9 (multiple networkDeviceIds with & separator)
	DeviceName            string  `url:"deviceName,omitempty"`            //Name of the device. This parameter supports wildcard (`*`) character -based search. Example: `wnbu-sjc*` or `*wnbu-sjc*` or `*wnbu-sjc` Examples: deviceName=wnbu-sjc24.cisco.com (single device name is requested) deviceName=wnbu-sjc24.cisco.com&deviceName=wnbu-sjc22.cisco.com (multiple device names are requested)
	DeviceSiteHierarchy   string  `url:"deviceSiteHierarchy,omitempty"`   //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. `Global/AreaName/BuildingName/FloorName`) This field supports wildcard asterisk (`*`) character search support. E.g. `*/San*, */San, /San*` Examples: `?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `?deviceSiteHierarchy=Global/AreaName/BuildingName/FloorName&deviceSiteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)
	DeviceSiteHierarchyID string  `url:"deviceSiteHierarchyId,omitempty"` //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*` Examples: `?deviceSiteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?deviceSiteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&deviceSiteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	DeviceSiteID          string  `url:"deviceSiteId,omitempty"`          //The UUID of the site. (Ex. `flooruuid`) Examples: `?deviceSiteIds=id1` (single id requested) `?deviceSiteIds=id1&deviceSiteIds=id2&siteId=id3` (multiple ids requested)
}
type RetrievesTheTotalNumberOfDHCPServicesForGivenParametersV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheServiceV1QueryParams struct {
	StartTime float64 `url:"startTime,omitempty"` //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime   float64 `url:"endTime,omitempty"`   //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
}
type RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheServiceV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1HeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheListOfDNSServicesForGivenParametersV1QueryParams struct {
	StartTime             float64 `url:"startTime,omitempty"`             //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime               float64 `url:"endTime,omitempty"`               //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	Limit                 float64 `url:"limit,omitempty"`                 //Maximum number of records to return
	Offset                float64 `url:"offset,omitempty"`                //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	SortBy                string  `url:"sortBy,omitempty"`                //Field name on which sorting needs to be done
	Order                 string  `url:"order,omitempty"`                 //The sort order of the field ascending or descending.
	ServerIP              string  `url:"serverIp,omitempty"`              //IP Address of the DNS Server. This parameter supports wildcard (`*`) character -based search. Example: `10.76.81.*` or `*56.78*` or `*50.28` Examples: serverIp=10.42.3.31 (single IP Address is requested) serverIp=10.42.3.31&serverIp=name2&fabricVnName=name3 (multiple IP Addresses are requested)
	DeviceID              string  `url:"deviceId,omitempty"`              //The device UUID.   Examples:  `deviceId=6bef213c-19ca-4170-8375-b694e251101c` (single deviceId is requested)  `deviceId=6bef213c-19ca-4170-8375-b694e251101c&deviceId=32219612-819e-4b5e-a96b-cf22aca13dd9 (multiple networkDeviceIds with & separator)
	DeviceSiteHierarchyID string  `url:"deviceSiteHierarchyId,omitempty"` //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*` Examples: `?deviceSiteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?deviceSiteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&deviceSiteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	DeviceSiteID          string  `url:"deviceSiteId,omitempty"`          //The UUID of the site. (Ex. `flooruuid`) Examples: `?deviceSiteIds=id1` (single id requested) `?deviceSiteIds=id1&deviceSiteIds=id2&siteId=id3` (multiple ids requested)
	SSID                  string  `url:"ssid,omitempty"`                  //SSID is the name of wireless network to which client connects to. It is also referred to as WLAN ID - Wireless Local Area Network Identifier. This field supports wildcard (`*`) character-based search. If the field contains the (`*`) character, please use the /query API for search. Ex: `*Alpha*` or `Alpha*` or `*Alpha` Examples: `ssid=Alpha` (single ssid requested) `ssid=Alpha&ssid=Guest` (multiple ssid requested)
}
type RetrievesTheListOfDNSServicesForGivenParametersV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheTotalNumberOfDNSServicesForGivenParametersV1QueryParams struct {
	StartTime             float64 `url:"startTime,omitempty"`             //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime               float64 `url:"endTime,omitempty"`               //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	ServerIP              string  `url:"serverIp,omitempty"`              //IP Address of the DNS Server. This parameter supports wildcard (`*`) character -based search. Example: `10.76.81.*` or `*56.78*` or `*50.28` Examples: serverIp=10.42.3.31 (single IP Address is requested) serverIp=10.42.3.31&serverIp=name2&fabricVnName=name3 (multiple IP Addresses are requested)
	DeviceID              string  `url:"deviceId,omitempty"`              //The device UUID.   Examples:  `deviceId=6bef213c-19ca-4170-8375-b694e251101c` (single deviceId is requested)  `deviceId=6bef213c-19ca-4170-8375-b694e251101c&deviceId=32219612-819e-4b5e-a96b-cf22aca13dd9 (multiple networkDeviceIds with & separator)
	DeviceSiteHierarchyID string  `url:"deviceSiteHierarchyId,omitempty"` //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*` Examples: `?deviceSiteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?deviceSiteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&deviceSiteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	DeviceSiteID          string  `url:"deviceSiteId,omitempty"`          //The UUID of the site. (Ex. `flooruuid`) Examples: `?deviceSiteIds=id1` (single id requested) `?deviceSiteIds=id1&deviceSiteIds=id2&siteId=id3` (multiple ids requested)
	SSID                  string  `url:"ssid,omitempty"`                  //SSID is the name of wireless network to which client connects to. It is also referred to as WLAN ID - Wireless Local Area Network Identifier. This field supports wildcard (`*`) character-based search. If the field contains the (`*`) character, please use the /query API for search. Ex: `*Alpha*` or `Alpha*` or `*Alpha` Examples: `ssid=Alpha` (single ssid requested) `ssid=Alpha&ssid=Guest` (multiple ssid requested)
}
type RetrievesTheTotalNumberOfDNSServicesForGivenParametersV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIDOfTheServiceV1QueryParams struct {
	StartTime float64 `url:"startTime,omitempty"` //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime   float64 `url:"endTime,omitempty"`   //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
}
type RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIDOfTheServiceV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1HeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesV1QueryParams struct {
	StartTime               float64 `url:"startTime,omitempty"`               //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to current time.
	EndTime                 float64 `url:"endTime,omitempty"`                 //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	Limit                   float64 `url:"limit,omitempty"`                   //Maximum number of records to return
	Offset                  float64 `url:"offset,omitempty"`                  //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	SortBy                  string  `url:"sortBy,omitempty"`                  //A field within the response to sort by.
	Order                   string  `url:"order,omitempty"`                   //The sort order of the field ascending or descending.
	SiteHierarchy           string  `url:"siteHierarchy,omitempty"`           //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. `Global/AreaName/BuildingName/FloorName`) This field supports wildcard asterisk (`*`) character search support. E.g. `*/San*, */San, /San*` Examples: `?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)
	SiteHierarchyID         string  `url:"siteHierarchyId,omitempty"`         //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*` Examples: `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	SiteID                  string  `url:"siteId,omitempty"`                  //The UUID of the site. (Ex. `flooruuid`) Examples: `?siteId=id1` (single id requested) `?siteId=id1&siteId=id2&siteId=id3` (multiple ids requested)
	View                    string  `url:"view,omitempty"`                    //Views which are supported by this API. Each view represents a specific         data set.           ### Response data provided by each view:             1. **configuration**          [id,name,adminStatus,description,duplexConfig,duplexOper,interfaceIfIndex,interfaceType,ipv4Address,ipv6AddressList,isL3Interface,isWan,macAddress,mediaType,name,operStatus,         portChannelId,portMode,         portType,speed,timestamp,vlanId,networkDeviceId,networkDeviceIpAddress,networkDeviceMacAddress,siteName,siteHierarchy,siteHierarchyId]             2. **statistics**          [id,name,rxDiscards,rxError,rxRate,rxUtilization,txDiscards,txError,txRate,txUtilization,networkDeviceId,networkDeviceIpAddress,networkDeviceMacAddress,siteName,siteHierarchy,siteHierarchyId]            3. **stackPort**          [id,name,peerStackMember,peerStackPort,stackPortType,networkDeviceId,networkDeviceIpAddress,networkDeviceMacAddress,siteName,siteHierarchy,siteHierarchyId]                   4. **poE**              [id, name,rxDiscards,rxError,rxRate,rxUtilization,txDiscards,txError,txRate,txUtilization,networkDeviceId,networkDeviceIpAddress,networkDeviceMacAddress,siteName,siteHierarchy,siteHierarchyId]              When this query parameter is not added by default all configuration attributes will be         available in the response.     **[configuration,statistics,stackPort]**
	Attribute               string  `url:"attribute,omitempty"`               //The following list of attributes can be provided in the attribute field          [id,adminStatus, description,duplexConfig,duplexOper,interfaceIfIndex,interfaceType,ipv4Address,ipv6AddressList,isL3Interface,isWan,macAddress,mediaType,name,operStatus,peerStackMember,peerStackPort, portChannelId,portMode, portType,rxDiscards,rxError,rxRate,rxUtilization,speed,stackPortType,timestamp,txDiscards,txError,txRate,txUtilization,vlanId,networkDeviceId,networkDeviceIpAddress,networkDeviceMacAddress,siteName,siteHierarchy,siteHierarchyId,poeAdminStatus,poeOperStatus,chassisId,moduleId,pdClassSignal,pdClassSpare,pdDeviceType,pdDeviceModel,pdPowerAdminMaxInWatt,pdPowerBudgetInWatt,pdPowerConsumedInWatt,pdPowerRemainingInWatt,pdMaxPowerDrawn,pdConnectedDeviceList,poeOperPriority,fastPoEEnabled,perpetualPoEEnabled,policingPoEEnabled,upoePlusEnabled,fourPairEnabled,poeDataTimestamp,pdLocation,pdDeviceName,pdConnectedSwitch,connectedSwitchUuid,ieeeCompliant,connectedSwitchType]          If length of attribute list is too long, please use 'views' param instead.          Examples:          attributes=name (single attribute requested)          attributes=name&description&duplexOper (multiple attributes with comma separator)
	NetworkDeviceID         string  `url:"networkDeviceId,omitempty"`         //The list of Network Device Uuids. (Ex. `6bef213c-19ca-4170-8375-b694e251101c`) Examples: `networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c` (single networkDeviceId requested) `networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c&networkDeviceId=32219612-819e-4b5e-a96b-cf22aca13dd9&networkDeviceId=2541e9a7-b80d-4955-8aa2-79b233318ba0` (multiple networkDeviceIds with & separator)
	NetworkDeviceIPAddress  string  `url:"networkDeviceIpAddress,omitempty"`  //The list of Network Device management IP Address. (Ex. `121.1.1.10`) This field supports wildcard (`*`) character-based search. Ex: `*1.1*` or `1.1*` or `*1.1` Examples: `networkDeviceIpAddress=121.1.1.10` `networkDeviceIpAddress=121.1.1.10&networkDeviceIpAddress=172.20.1.10&networkDeviceIpAddress=10.10.20.10` (multiple networkDevice IP Address with & separator)
	NetworkDeviceMacAddress string  `url:"networkDeviceMacAddress,omitempty"` //The list of Network Device MAC Address. (Ex. `64:f6:9d:07:9a:00`) This field supports wildcard (`*`) character-based search. Ex: `*AB:AB:AB*` or `AB:AB:AB*` or `*AB:AB:AB` Examples: `networkDeviceMacAddress=64:f6:9d:07:9a:00` `networkDeviceMacAddress=64:f6:9d:07:9a:00&networkDeviceMacAddress=70:56:9d:07:ac:77` (multiple networkDevice MAC addresses with & separator)
	InterfaceID             string  `url:"interfaceId,omitempty"`             //The list of Interface Uuids. (Ex. `6bef213c-19ca-4170-8375-b694e251101c`) Examples: `interfaceId=6bef213c-19ca-4170-8375-b694e251101c` (single interface uuid ) `interfaceId=6bef213c-19ca-4170-8375-b694e251101c&32219612-819e-4b5e-a96b-cf22aca13dd9&2541e9a7-b80d-4955-8aa2-79b233318ba0` (multiple Interface uuid with & separator)
	InterfaceName           string  `url:"interfaceName,omitempty"`           //The list of Interface name (Ex. `GigabitEthernet1/0/1`) This field supports wildcard (`*`) character-based search. Ex: `*1/0/1*` or `1/0/1*` or `*1/0/1` Examples: `interfaceNames=GigabitEthernet1/0/1` (single interface name) `interfaceNames=GigabitEthernet1/0/1&GigabitEthernet2/0/1&GigabitEthernet3/0/1` (multiple interface names with & separator)
}
type GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountV1QueryParams struct {
	StartTime               float64 `url:"startTime,omitempty"`               //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to current time.
	EndTime                 float64 `url:"endTime,omitempty"`                 //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	SiteHierarchy           string  `url:"siteHierarchy,omitempty"`           //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. `Global/AreaName/BuildingName/FloorName`) This field supports wildcard asterisk (`*`) character search support. E.g. `*/San*, */San, /San*` Examples: `?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)
	SiteHierarchyID         string  `url:"siteHierarchyId,omitempty"`         //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*` Examples: `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	SiteID                  string  `url:"siteId,omitempty"`                  //The UUID of the site. (Ex. `flooruuid`) Examples: `?siteId=id1` (single id requested) `?siteId=id1&siteId=id2&siteId=id3` (multiple ids requested)
	NetworkDeviceID         string  `url:"networkDeviceId,omitempty"`         //The list of Network Device Uuids. (Ex. `6bef213c-19ca-4170-8375-b694e251101c`) Examples: `networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c` (single networkDeviceId requested) `networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c&networkDeviceId=32219612-819e-4b5e-a96b-cf22aca13dd9&networkDeviceId=2541e9a7-b80d-4955-8aa2-79b233318ba0` (multiple networkDeviceIds with & separator)
	NetworkDeviceIPAddress  string  `url:"networkDeviceIpAddress,omitempty"`  //The list of Network Device management IP Address. (Ex. `121.1.1.10`) This field supports wildcard (`*`) character-based search. Ex: `*1.1*` or `1.1*` or `*1.1` Examples: `networkDeviceIpAddress=121.1.1.10` `networkDeviceIpAddress=121.1.1.10&networkDeviceIpAddress=172.20.1.10&networkDeviceIpAddress=10.10.20.10` (multiple networkDevice IP Address with & separator)
	NetworkDeviceMacAddress string  `url:"networkDeviceMacAddress,omitempty"` //The list of Network Device MAC Address. (Ex. `64:f6:9d:07:9a:00`) This field supports wildcard (`*`) character-based search. Ex: `*AB:AB:AB*` or `AB:AB:AB*` or `*AB:AB:AB` Examples: `networkDeviceMacAddress=64:f6:9d:07:9a:00` `networkDeviceMacAddress=64:f6:9d:07:9a:00&networkDeviceMacAddress=70:56:9d:07:ac:77` (multiple networkDevice MAC addresses with & separator)
	InterfaceID             string  `url:"interfaceId,omitempty"`             //The list of Interface Uuids. (Ex. `6bef213c-19ca-4170-8375-b694e251101c`) Examples: `interfaceId=6bef213c-19ca-4170-8375-b694e251101c` (single interface uuid ) `interfaceId=6bef213c-19ca-4170-8375-b694e251101c&32219612-819e-4b5e-a96b-cf22aca13dd9&2541e9a7-b80d-4955-8aa2-79b233318ba0` (multiple Interface uuid with & separator)
	InterfaceName           string  `url:"interfaceName,omitempty"`           //The list of Interface name (Ex. `GigabitEthernet1/0/1`) This field supports wildcard (`*`) character-based search. Ex: `*1/0/1*` or `1/0/1*` or `*1/0/1` Examples: `interfaceNames=GigabitEthernet1/0/1` (single interface name) `interfaceNames=GigabitEthernet1/0/1&GigabitEthernet2/0/1&GigabitEthernet3/0/1` (multiple interface names with & separator)
}
type GetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsAndPoeDataV1QueryParams struct {
	StartTime float64 `url:"startTime,omitempty"` //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to current time.
	EndTime   float64 `url:"endTime,omitempty"`   //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	View      string  `url:"view,omitempty"`      //Interface data model views
	Attribute string  `url:"attribute,omitempty"` //The following list of attributes can be provided in the attribute field          [id,adminStatus, description,duplexConfig,duplexOper,interfaceIfIndex,interfaceType,ipv4Address,ipv6AddressList,isL3Interface,isWan,macAddress,mediaType,name,operStatus,peerStackMember,peerStackPort, portChannelId,portMode, portType,rxDiscards,rxError,rxRate,rxUtilization,speed,stackPortType,timestamp,txDiscards,txError,txRate,txUtilization,vlanId,networkDeviceId,networkDeviceIpAddress,networkDeviceMacAddress,siteName,siteHierarchy,siteHierarchyId,poeAdminStatus,poeOperStatus,chassisId,moduleId,pdClassSignal,pdClassSpare,pdDeviceType,pdDeviceModel,pdPowerAdminMaxInWatt,pdPowerBudgetInWatt,pdPowerConsumedInWatt,pdPowerRemainingInWatt,pdMaxPowerDrawn,pdConnectedDeviceList,poeOperPriority,fastPoEEnabled,perpetualPoEEnabled,policingPoEEnabled,upoePlusEnabled,fourPairEnabled,poeDataTimestamp,pdLocation,pdDeviceName,pdConnectedSwitch,connectedSwitchUuid,ieeeCompliant,connectedSwitchType]          If length of attribute list is too long, please use 'views' param instead.          Examples:          attributes=name (single attribute requested)          attributes=name&description&duplexOper (multiple attributes with comma separator)
}
type GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1QueryParams struct {
	StartTime           float64 `url:"startTime,omitempty"`           //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to current time.
	EndTime             float64 `url:"endTime,omitempty"`             //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	Limit               float64 `url:"limit,omitempty"`               //Maximum number of records to return
	Offset              float64 `url:"offset,omitempty"`              //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	SortBy              string  `url:"sortBy,omitempty"`              //A field within the response to sort by.
	Order               string  `url:"order,omitempty"`               //The sort order of the field ascending or descending.
	SiteHierarchy       string  `url:"siteHierarchy,omitempty"`       //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. `Global/AreaName/BuildingName/FloorName`) This field supports wildcard asterisk (*) character search support. E.g. */San*, */San, /San* Examples: `?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)
	SiteHierarchyID     string  `url:"siteHierarchyId,omitempty"`     //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (*) character search support. E.g. `*uuid*, *uuid, uuid* Examples: `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	SiteID              string  `url:"siteId,omitempty"`              //The UUID of the site. (Ex. `flooruuid`) This field supports wildcard asterisk (*) character search support. E.g.*flooruuid*, *flooruuid, flooruuid* Examples: `?siteId=id1` (single id requested) `?siteId=id1&siteId=id2&siteId=id3` (multiple ids requested)
	ID                  string  `url:"id,omitempty"`                  //The list of entity Uuids. (Ex."6bef213c-19ca-4170-8375-b694e251101c") Examples: id=6bef213c-19ca-4170-8375-b694e251101c (single entity uuid requested) id=6bef213c-19ca-4170-8375-b694e251101c&id=32219612-819e-4b5e-a96b-cf22aca13dd9&id=2541e9a7-b80d-4955-8aa2-79b233318ba0 (multiple entity uuid with '&' separator)
	ManagementIPAddress string  `url:"managementIpAddress,omitempty"` //The list of entity management IP Address. It can be either Ipv4 or Ipv6 address or combination of both(Ex. "121.1.1.10") This field supports wildcard (`*`) character-based search.  Ex: `*1.1*` or `1.1*` or `*1.1` Examples: managementIpAddresses=121.1.1.10 managementIpAddresses=121.1.1.10&managementIpAddresses=172.20.1.10&managementIpAddresses=200:10&=managementIpAddresses172.20.3.4 (multiple entity IP Address with & separator)
	MacAddress          string  `url:"macAddress,omitempty"`          //The macAddress of the network device or client This field supports wildcard (`*`) character-based search. Ex: `*AB:AB:AB*` or `AB:AB:AB*` or `*AB:AB:AB` Examples: `macAddress=AB:AB:AB:CD:CD:CD` (single macAddress requested) `macAddress=AB:AB:AB:CD:CD:DC&macAddress=AB:AB:AB:CD:CD:FE` (multiple macAddress requested)
	Family              string  `url:"family,omitempty"`              //The list of network device family names Examples:family=Switches and Hubs (single network device family name )family=Switches and Hubs&family=Router&family=Wireless Controller (multiple Network device family names with & separator). This field is not case sensitive.
	Type                string  `url:"type,omitempty"`                //The list of network device type This field supports wildcard (`*`) character-based search. Ex: `*9407R*` or `*9407R` or `9407R*` Examples: type=SwitchesCisco Catalyst 9407R Switch (single network device types ) type=Cisco Catalyst 38xx stack-able ethernet switch&type=Cisco 3945 Integrated Services Router G2 (multiple Network device types with & separator)
	Role                string  `url:"role,omitempty"`                //The list of network device role. Examples:role=CORE, role=CORE&role=ACCESS&role=ROUTER (multiple Network device roles with & separator). This field is not case sensitive.
	SerialNumber        string  `url:"serialNumber,omitempty"`        //The list of network device serial numbers. This field supports wildcard (`*`) character-based search.  Ex: `*MS1SV*` or `MS1SV*` or `*MS1SV` Examples: serialNumber=9FUFMS1SVAX serialNumber=9FUFMS1SVAX&FCW2333Q0BY&FJC240617JX(multiple Network device serial number with & separator)
	MaintenanceMode     bool    `url:"maintenanceMode,omitempty"`     //The device maintenanceMode status true or false
	SoftwareVersion     string  `url:"softwareVersion,omitempty"`     //The list of network device software version This field supports wildcard (`*`) character-based search. Ex: `*17.8*` or `*17.8` or `17.8*` Examples: softwareVersion=2.3.4.0 (single network device software version ) softwareVersion=17.9.3.23&softwareVersion=17.7.1.2&softwareVersion=*.17.7 (multiple Network device software versions with & separator)
	HealthScore         string  `url:"healthScore,omitempty"`         //The list of entity health score categories Examples: healthScore=good, healthScore=good&healthScore=fair (multiple entity healthscore values with & separator). This field is not case sensitive.
	View                string  `url:"view,omitempty"`                //The List of Network Device model views. Please refer to ```NetworkDeviceView``` section in the Open API specification document mentioned in the description.
	Attribute           string  `url:"attribute,omitempty"`           //The List of Network Device model attributes. Please refer to ```NetworkDeviceAttribute``` section in the Open API specification document mentioned in the description.
	FabricSiteID        string  `url:"fabricSiteId,omitempty"`        //The fabric site Id or list to fabric site Ids to filter the data  This field supports wildcard asterisk (*) character search support. E.g. *uuid*, *uuid, uuid*  Examples:  `?fabricSiteId=fabricSiteUuid)  ?fabricSiteId=fabricSiteUuid1&fabricSiteId=fabricSiteUuid2 (multiple fabricSiteIds requested)
	L2Vn                string  `url:"l2Vn,omitempty"`                //The L2 Virtual Network Id or list to Virtual Network Ids to filter the data  This field supports wildcard asterisk (*) character search support. E.g. *uuid*, *uuid, uuid*  Examples:  `?l2Vn=virtualNetworkId  ?l2Vn=virtualNetworkId1&l2Vn=virtualNetworkId2 (multiple virtualNetworkId's requested)
	L3Vn                string  `url:"l3Vn,omitempty"`                //The L3 Virtual Network Id or list to Virtual Network Ids to filter the data  This field supports wildcard asterisk (*) character search support. E.g. *uuid*, *uuid, uuid*  Examples:  `?l3Vn=virtualNetworkId  ?l3Vn=virtualNetworkId1&l3Vn=virtualNetworkId2 (multiple virtualNetworkId's requested)
	TransitNetworkID    string  `url:"transitNetworkId,omitempty"`    //The Transit Network Id or list to Transit Network Ids to filter the data  This field supports wildcard asterisk (*) character search support. E.g. *uuid*, *uuid, uuid*  Examples:  `?transitNetworkId=transitNetworkId  ?transitNetworkId=transitNetworkuuid1&transitNetworkId=transitNetworkuuid1 (multiple transitNetworkIds requested
	FabricRole          string  `url:"fabricRole,omitempty"`          //The list of fabric device role. Examples: fabricRole=BORDER, fabricRole=BORDER&fabricRole=EDGE (multiple fabric device roles with & separator)  Available values : BORDER, EDGE, MAP-SERVER, LEAF, SPINE, TRANSIT-CP, EXTENDED-NODE, WLC, UNIFIED-AP
}
type GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersV1QueryParams struct {
	StartTime           float64 `url:"startTime,omitempty"`           //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to current time.
	EndTime             float64 `url:"endTime,omitempty"`             //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	ID                  string  `url:"id,omitempty"`                  //The list of entity Uuids. (Ex."6bef213c-19ca-4170-8375-b694e251101c") Examples: id=6bef213c-19ca-4170-8375-b694e251101c (single entity uuid requested) id=6bef213c-19ca-4170-8375-b694e251101c&id=32219612-819e-4b5e-a96b-cf22aca13dd9&id=2541e9a7-b80d-4955-8aa2-79b233318ba0 (multiple entity uuid with '&' separator)
	SiteHierarchy       string  `url:"siteHierarchy,omitempty"`       //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. `Global/AreaName/BuildingName/FloorName`) This field supports wildcard asterisk (*) character search support. E.g. */San*, */San, /San* Examples: `?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)
	SiteHierarchyID     string  `url:"siteHierarchyId,omitempty"`     //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (*) character search support. E.g. `*uuid*, *uuid, uuid* Examples: `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	SiteID              string  `url:"siteId,omitempty"`              //The UUID of the site. (Ex. `flooruuid`) This field supports wildcard asterisk (*) character search support. E.g.*flooruuid*, *flooruuid, flooruuid* Examples: `?siteId=id1` (single id requested) `?siteId=id1&siteId=id2&siteId=id3` (multiple ids requested)
	ManagementIPAddress string  `url:"managementIpAddress,omitempty"` //The list of entity management IP Address. It can be either Ipv4 or Ipv6 address or combination of both(Ex. "121.1.1.10") This field supports wildcard (`*`) character-based search.  Ex: `*1.1*` or `1.1*` or `*1.1` Examples: managementIpAddresses=121.1.1.10 managementIpAddresses=121.1.1.10&managementIpAddresses=172.20.1.10&managementIpAddresses=200:10&=managementIpAddresses172.20.3.4 (multiple entity IP Address with & separator)
	MacAddress          string  `url:"macAddress,omitempty"`          //The macAddress of the network device or client This field supports wildcard (`*`) character-based search. Ex: `*AB:AB:AB*` or `AB:AB:AB*` or `*AB:AB:AB` Examples: `macAddress=AB:AB:AB:CD:CD:CD` (single macAddress requested) `macAddress=AB:AB:AB:CD:CD:DC&macAddress=AB:AB:AB:CD:CD:FE` (multiple macAddress requested)
	Family              string  `url:"family,omitempty"`              //The list of network device family names Examples:family=Switches and Hubs (single network device family name )family=Switches and Hubs&family=Router&family=Wireless Controller (multiple Network device family names with & separator). This field is not case sensitive.
	Type                string  `url:"type,omitempty"`                //The list of network device type This field supports wildcard (`*`) character-based search. Ex: `*9407R*` or `*9407R` or `9407R*`Examples:type=SwitchesCisco Catalyst 9407R Switch (single network device types )type=Cisco Catalyst 38xx stack-able ethernet switch&type=Cisco 3945 Integrated Services Router G2 (multiple Network device types with & separator)
	Role                string  `url:"role,omitempty"`                //The list of network device role. Examples:role=CORE, role=CORE&role=ACCESS&role=ROUTER (multiple Network device roles with & separator). This field is not case sensitive.
	SerialNumber        string  `url:"serialNumber,omitempty"`        //The list of network device serial numbers. This field supports wildcard (`*`) character-based search.  Ex: `*MS1SV*` or `MS1SV*` or `*MS1SV` Examples: serialNumber=9FUFMS1SVAX serialNumber=9FUFMS1SVAX&FCW2333Q0BY&FJC240617JX(multiple Network device serial number with & separator)
	MaintenanceMode     bool    `url:"maintenanceMode,omitempty"`     //The device maintenanceMode status true or false
	SoftwareVersion     string  `url:"softwareVersion,omitempty"`     //The list of network device software version This field supports wildcard (`*`) character-based search. Ex: `*17.8*` or `*17.8` or `17.8*` Examples: softwareVersion=2.3.4.0 (single network device software version ) softwareVersion=17.9.3.23&softwareVersion=17.7.1.2&softwareVersion=*.17.7 (multiple Network device software versions with & separator)
	HealthScore         string  `url:"healthScore,omitempty"`         //The list of entity health score categories Examples:healthScore=good,healthScore=good&healthScore=fair (multiple entity healthscore values with & separator). This field is not case sensitive.
	FabricSiteID        string  `url:"fabricSiteId,omitempty"`        //The fabric site Id or list to fabric site Ids to filter the data  This field supports wildcard asterisk (*) character search support. E.g. *uuid*, *uuid, uuid*  Examples:  `?fabricSiteId=fabricSiteUuid)  ?fabricSiteId=fabricSiteUuid1&fabricSiteId=fabricSiteUuid2 (multiple fabricSiteIds requested)
	L2Vn                string  `url:"l2Vn,omitempty"`                //The L2 Virtual Network Id or list to Virtual Network Ids to filter the data  This field supports wildcard asterisk (*) character search support. E.g. *uuid*, *uuid, uuid*  Examples:  `?l2Vn=virtualNetworkId  ?l2Vn=virtualNetworkId1&l2Vn=virtualNetworkId2 (multiple virtualNetworkId's requested)
	L3Vn                string  `url:"l3Vn,omitempty"`                //The L3 Virtual Network Id or list to Virtual Network Ids to filter the data  This field supports wildcard asterisk (*) character search support. E.g. *uuid*, *uuid, uuid*  Examples:  `?l3Vn=virtualNetworkId  ?l3Vn=virtualNetworkId1&l3Vn=virtualNetworkId2 (multiple virtualNetworkId's requested)
	TransitNetworkID    string  `url:"transitNetworkId,omitempty"`    //The Transit Network Id or list to Transit Network Ids to filter the data  This field supports wildcard asterisk (*) character search support. E.g. *uuid*, *uuid, uuid*  Examples:  `?transitNetworkId=transitNetworkId  ?transitNetworkId=transitNetworkuuid1&transitNetworkId=transitNetworkuuid1 (multiple transitNetworkIds requested)
	FabricRole          string  `url:"fabricRole,omitempty"`          //The list of fabric device role. Examples: fabricRole=BORDER, fabricRole=BORDER&fabricRole=EDGE (multiple fabric device roles with & separator)  Available values : BORDER, EDGE, MAP-SERVER, LEAF, SPINE, TRANSIT-CP, EXTENDED-NODE, WLC, UNIFIED-AP
}
type GetTheDeviceDataForTheGivenDeviceIDUUIDV1QueryParams struct {
	StartTime float64 `url:"startTime,omitempty"` //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to current time.
	EndTime   float64 `url:"endTime,omitempty"`   //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	View      string  `url:"view,omitempty"`      //The List of Network Device model views. Please refer to ```NetworkDeviceView``` section in the Open API specification document mentioned in the description.
	Attribute string  `url:"attribute,omitempty"` //The List of Network Device model attributes. Please refer to ```NetworkDeviceAttribute``` section in the Open API specification document mentioned in the description.
}
type GetPlannedAccessPointsForBuildingV1QueryParams struct {
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page;The minimum is 1, and the maximum is 500.
	Offset float64 `url:"offset,omitempty"` //The page offset for the response. E.g. if limit=100, offset=0 will return first 100 records, offset=1 will return next 100 records, etc.
	Radios bool    `url:"radios,omitempty"` //Whether to include the planned radio details of the planned access points
}
type GetDeviceDetailV1QueryParams struct {
	Timestamp  float64 `url:"timestamp,omitempty"`  //UTC timestamp of device data in milliseconds
	IDentifier string  `url:"identifier,omitempty"` //One of "macAddress", "nwDeviceName", "uuid" (case insensitive)
	SearchBy   string  `url:"searchBy,omitempty"`   //MAC Address, device name, or UUID of the network device
}
type GetDeviceEnrichmentDetailsV1HeaderParams struct {
	EntityType        string `url:"entity_type,omitempty"`         //Expects type string. Device enrichment details can be fetched based on either Device ID or Device MAC address or Device IP Address. This parameter value must either be device_id/mac_address/ip_address
	EntityValue       string `url:"entity_value,omitempty"`        //Expects type string. Contains the actual value for the entity type that has been defined
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type bool. For the enrichment details to be made available as part of the API response, this header must be set to true. This header must be explicitly passed when called from client applications outside Catalyst Center
}
type DevicesV1QueryParams struct {
	DeviceRole string  `url:"deviceRole,omitempty"` //CORE, ACCESS, DISTRIBUTION, ROUTER, WLC, or AP (case insensitive)
	SiteID     string  `url:"siteId,omitempty"`     //DNAC site UUID
	Health     string  `url:"health,omitempty"`     //DNAC health catagory: POOR, FAIR, or GOOD (case insensitive)
	StartTime  float64 `url:"startTime,omitempty"`  //UTC epoch time in milliseconds
	EndTime    float64 `url:"endTime,omitempty"`    //UTC epoch time in milliseconds
	Limit      float64 `url:"limit,omitempty"`      //Max number of device entries in the response (default to 50. Max at 500)
	Offset     float64 `url:"offset,omitempty"`     //The offset of the first device in the returned data (Mutiple of 'limit' + 1)
}
type GetPlannedAccessPointsForFloorV1QueryParams struct {
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page;The minimum is 1, and the maximum is 500.
	Offset float64 `url:"offset,omitempty"` //The page offset for the response. E.g. if limit=100, offset=0 will return first 100 records, offset=1 will return next 100 records, etc.
	Radios bool    `url:"radios,omitempty"` //Whether to include the planned radio details of the planned access points
}
type GetAllHealthScoreDefinitionsForGivenFiltersV1QueryParams struct {
	DeviceType              string  `url:"deviceType,omitempty"`              //These are the device families supported for health score definitions. If no input is made on device family, all device families are considered.
	ID                      string  `url:"id,omitempty"`                      //The definition identifier. Examples: id=015d9cba-4f53-4087-8317-7e49e5ffef46 (single entity id request) id=015d9cba-4f53-4087-8317-7e49e5ffef46&id=015d9cba-4f53-4087-8317-7e49e5ffef47 (multiple ids in the query param)
	IncludeForOverallHealth bool    `url:"includeForOverallHealth,omitempty"` //The inclusion status of the issue definition, either true or false. true indicates that particular health metric is included in overall health computation, otherwise false. By default it's set to true.
	Attribute               string  `url:"attribute,omitempty"`               //These are the attributes supported in health score definitions response. By default, all properties are sent in response.
	Offset                  float64 `url:"offset,omitempty"`                  //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	Limit                   float64 `url:"limit,omitempty"`                   //Maximum number of records to return
}
type GetAllHealthScoreDefinitionsForGivenFiltersV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type UpdateHealthScoreDefinitionsV1HeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersV1QueryParams struct {
	DeviceType              string `url:"deviceType,omitempty"`              //These are the device families supported for health score definitions. If no input is made on device family, all device families are considered.
	ID                      string `url:"id,omitempty"`                      //The definition identifier. Examples: id=015d9cba-4f53-4087-8317-7e49e5ffef46 (single entity id request) id=015d9cba-4f53-4087-8317-7e49e5ffef46&id=015d9cba-4f53-4087-8317-7e49e5ffef47 (multiple ids in the query param)
	IncludeForOverallHealth bool   `url:"includeForOverallHealth,omitempty"` //The inclusion status of the issue definition, either true or false. true indicates that particular health metric is included in overall health computation, otherwise false. By default it's set to true.
}
type GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetHealthScoreDefinitionForTheGivenIDV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetAllInterfacesV1QueryParams struct {
	Offset         int    `url:"offset,omitempty"`         //Offset
	Limit          int    `url:"limit,omitempty"`          //The number of records to show for this page. Min: 1, Max: 500
	LastInputTime  string `url:"lastInputTime,omitempty"`  //Last Input Time
	LastOutputTime string `url:"lastOutputTime,omitempty"` //Last Output Time
}
type GetInterfaceDetailsByDeviceIDAndInterfaceNameV1QueryParams struct {
	Name string `url:"name,omitempty"` //Interface name
}
type UpdateInterfaceDetailsV1QueryParams struct {
	DeploymentMode string `url:"deploymentMode,omitempty"` //Preview/Deploy ['Preview' means the configuration is not pushed to the device. 'Deploy' makes the configuration pushed to the device]
}
type ClearMacAddressTableV1QueryParams struct {
	DeploymentMode string `url:"deploymentMode,omitempty"` //Preview/Deploy ['Preview' means the configuration is not pushed to the device. 'Deploy' makes the configuration pushed to the device]
}
type GetDeviceListV1QueryParams struct {
	Hostname                  []string `url:"hostname,omitempty"`                   //hostname
	ManagementIPAddress       []string `url:"managementIpAddress,omitempty"`        //managementIpAddress
	MacAddress                []string `url:"macAddress,omitempty"`                 //macAddress
	LocationName              []string `url:"locationName,omitempty"`               //locationName
	SerialNumber              []string `url:"serialNumber,omitempty"`               //serialNumber
	Location                  []string `url:"location,omitempty"`                   //location
	Family                    []string `url:"family,omitempty"`                     //family
	Type                      []string `url:"type,omitempty"`                       //type
	Series                    []string `url:"series,omitempty"`                     //series
	CollectionStatus          []string `url:"collectionStatus,omitempty"`           //collectionStatus
	CollectionInterval        []string `url:"collectionInterval,omitempty"`         //collectionInterval
	NotSyncedForMinutes       []string `url:"notSyncedForMinutes,omitempty"`        //notSyncedForMinutes
	ErrorCode                 []string `url:"errorCode,omitempty"`                  //errorCode
	ErrorDescription          []string `url:"errorDescription,omitempty"`           //errorDescription
	SoftwareVersion           []string `url:"softwareVersion,omitempty"`            //softwareVersion
	SoftwareType              []string `url:"softwareType,omitempty"`               //softwareType
	PlatformID                []string `url:"platformId,omitempty"`                 //platformId
	Role                      []string `url:"role,omitempty"`                       //role
	ReachabilityStatus        []string `url:"reachabilityStatus,omitempty"`         //reachabilityStatus
	UpTime                    []string `url:"upTime,omitempty"`                     //upTime
	AssociatedWlcIP           []string `url:"associatedWlcIp,omitempty"`            //associatedWlcIp
	Licensename               []string `url:"license.name,omitempty"`               //licenseName
	Licensetype               []string `url:"license.type,omitempty"`               //licenseType
	Licensestatus             []string `url:"license.status,omitempty"`             //licenseStatus
	Modulename                []string `url:"module+name,omitempty"`                //moduleName
	Moduleequpimenttype       []string `url:"module+equpimenttype,omitempty"`       //moduleEqupimentType
	Moduleservicestate        []string `url:"module+servicestate,omitempty"`        //moduleServiceState
	Modulevendorequipmenttype []string `url:"module+vendorequipmenttype,omitempty"` //moduleVendorEquipmentType
	Modulepartnumber          []string `url:"module+partnumber,omitempty"`          //modulePartNumber
	Moduleoperationstatecode  []string `url:"module+operationstatecode,omitempty"`  //moduleOperationStateCode
	ID                        string   `url:"id,omitempty"`                         //Accepts comma separated ids and return list of network-devices for the given ids. If invalid or not-found ids are provided, null entry will be returned in the list.
	DeviceSupportLevel        string   `url:"deviceSupportLevel,omitempty"`         //deviceSupportLevel
	Offset                    int      `url:"offset,omitempty"`                     //offset >= 1 [X gives results from Xth device onwards]
	Limit                     int      `url:"limit,omitempty"`                      //The number of records to show for this page. Min: 1, Max: 500
}
type GetDeviceValuesThatMatchFullyOrPartiallyAnAttributeV1QueryParams struct {
	VrfName                   string `url:"vrfName,omitempty"`                   //vrfName
	ManagementIPAddress       string `url:"managementIpAddress,omitempty"`       //managementIpAddress
	Hostname                  string `url:"hostname,omitempty"`                  //hostname
	MacAddress                string `url:"macAddress,omitempty"`                //macAddress
	Family                    string `url:"family,omitempty"`                    //family
	CollectionStatus          string `url:"collectionStatus,omitempty"`          //collectionStatus
	CollectionInterval        string `url:"collectionInterval,omitempty"`        //collectionInterval
	SoftwareVersion           string `url:"softwareVersion,omitempty"`           //softwareVersion
	SoftwareType              string `url:"softwareType,omitempty"`              //softwareType
	ReachabilityStatus        string `url:"reachabilityStatus,omitempty"`        //reachabilityStatus
	ReachabilityFailureReason string `url:"reachabilityFailureReason,omitempty"` //reachabilityFailureReason
	ErrorCode                 string `url:"errorCode,omitempty"`                 //errorCode
	PlatformID                string `url:"platformId,omitempty"`                //platformId
	Series                    string `url:"series,omitempty"`                    //series
	Type                      string `url:"type,omitempty"`                      //type
	SerialNumber              string `url:"serialNumber,omitempty"`              //serialNumber
	UpTime                    string `url:"upTime,omitempty"`                    //upTime
	Role                      string `url:"role,omitempty"`                      //role
	RoleSource                string `url:"roleSource,omitempty"`                //roleSource
	AssociatedWlcIP           string `url:"associatedWlcIp,omitempty"`           //associatedWlcIp
	Offset                    int    `url:"offset,omitempty"`                    //offset
	Limit                     int    `url:"limit,omitempty"`                     //The number of records to show for this page. Min: 1, Max: 500
}
type GetDeviceCountKnowYourNetworkV1QueryParams struct {
	Hostname            []string `url:"hostname,omitempty"`            //hostname
	ManagementIPAddress []string `url:"managementIpAddress,omitempty"` //managementIpAddress
	MacAddress          []string `url:"macAddress,omitempty"`          //macAddress
	LocationName        []string `url:"locationName,omitempty"`        //locationName
}
type GetFunctionalCapabilityForDevicesV1QueryParams struct {
	DeviceID     string   `url:"deviceId,omitempty"`     //Accepts comma separated deviceid's and return list of functional-capabilities for the given id's. If invalid or not-found id's are provided, null entry will be returned in the list.
	FunctionName []string `url:"functionName,omitempty"` //functionName
}
type InventoryInsightDeviceLinkMismatchAPIV1QueryParams struct {
	Offset   int    `url:"offset,omitempty"`   //Row Number.  Default value is 1
	Limit    int    `url:"limit,omitempty"`    //The number of records to show for this page. Min: 1, Max: 500
	Category string `url:"category,omitempty"` //Links mismatch category.  Value can be speed-duplex or vlan.
	SortBy   string `url:"sortBy,omitempty"`   //Sort By
	Order    string `url:"order,omitempty"`    //Order.  Value can be asc or desc.  Default value is asc
}
type GetModulesV1QueryParams struct {
	DeviceID                 string   `url:"deviceId,omitempty"`                 //deviceId
	Limit                    int      `url:"limit,omitempty"`                    //The number of records to show for this page. Min: 1, Max: 500
	Offset                   int      `url:"offset,omitempty"`                   //offset
	NameList                 []string `url:"nameList,omitempty"`                 //nameList
	VendorEquipmentTypeList  []string `url:"vendorEquipmentTypeList,omitempty"`  //vendorEquipmentTypeList
	PartNumberList           []string `url:"partNumberList,omitempty"`           //partNumberList
	OperationalStateCodeList []string `url:"operationalStateCodeList,omitempty"` //operationalStateCodeList
}
type GetModuleCountV1QueryParams struct {
	DeviceID                 string   `url:"deviceId,omitempty"`                 //deviceId
	NameList                 []string `url:"nameList,omitempty"`                 //nameList
	VendorEquipmentTypeList  []string `url:"vendorEquipmentTypeList,omitempty"`  //vendorEquipmentTypeList
	PartNumberList           []string `url:"partNumberList,omitempty"`           //partNumberList
	OperationalStateCodeList []string `url:"operationalStateCodeList,omitempty"` //operationalStateCodeList
}
type SyncDevicesV1QueryParams struct {
	ForceSync bool `url:"forceSync,omitempty"` //forceSync
}
type GetDevicesRegisteredForWsaNotificationV1QueryParams struct {
	SerialNumber string `url:"serialNumber,omitempty"` //Serial number of the device
	Macaddress   string `url:"macaddress,omitempty"`   //Mac addres of the device
}
type GetAllUserDefinedFieldsV1QueryParams struct {
	ID   string `url:"id,omitempty"`   //Comma-seperated id(s) used for search/filtering
	Name string `url:"name,omitempty"` //Comma-seperated name(s) used for search/filtering
}
type RemoveUserDefinedFieldFromDeviceV1QueryParams struct {
	Name string `url:"name,omitempty"` //Name of UDF to be removed
}
type GetTheDetailsOfPhysicalComponentsOfTheGivenDeviceV1QueryParams struct {
	Type string `url:"type,omitempty"` //Type value can be PowerSupply, Fan, Chassis, Backplane, Module, PROCESSOR, Other, SFP. If no type is mentioned, All equipments are fetched for the device.
}
type ReturnsPoeInterfaceDetailsForTheDeviceV1QueryParams struct {
	InterfaceNameList string `url:"interfaceNameList,omitempty"` //comma seperated interface names
}
type DeleteDeviceByIDV1QueryParams struct {
	CleanConfig bool `url:"cleanConfig,omitempty"` //Selecting the clean up configuration option will attempt to remove device settings that were configured during the addition of the device to the inventory and site assignment. Please note that this operation is different from deprovisioning. It does not remove configurations that were pushed during device provisioning.
}
type GetDeviceInterfaceVLANsV1QueryParams struct {
	InterfaceType string `url:"interfaceType,omitempty"` //Vlan associated with sub-interface. If no interfaceType mentioned it will return all types of Vlan interfaces. If interfaceType is selected but not specified then it will take default value.
}
type RetrieveScheduledMaintenanceWindowsForNetworkDevicesV1QueryParams struct {
	NetworkDeviceIDs string `url:"networkDeviceIds,omitempty"` //List of network device ids.
	Status           string `url:"status,omitempty"`           //The status of the maintenance schedule. Possible values are: UPCOMING, IN_PROGRESS, COMPLETED, FAILED. Refer features for more details.
	Limit            string `url:"limit,omitempty"`            //The number of records to show for this page. Min: 1, Max: 500
	Offset           string `url:"offset,omitempty"`           //The first record to show for this page; the first record is numbered 1.
	SortBy           string `url:"sortBy,omitempty"`           //A property within the response to sort by.
	Order            string `url:"order,omitempty"`            //Whether ascending or descending order should be used to sort the response.
}
type RetrieveTheTotalNumberOfScheduledMaintenanceWindowsV1QueryParams struct {
	NetworkDeviceIDs string `url:"networkDeviceIds,omitempty"` //List of network device ids.
	Status           string `url:"status,omitempty"`           //The status of the maintenance schedule. Possible values are: UPCOMING, IN_PROGRESS, COMPLETED, FAILED. Refer features for more details.
}
type RetrieveNetworkDevicesV1QueryParams struct {
	ID                 string `url:"id,omitempty"`                 //Network device Id
	ManagementAddress  string `url:"managementAddress,omitempty"`  //Management address of the network device
	SerialNumber       string `url:"serialNumber,omitempty"`       //Serial number of the network device
	Family             string `url:"family,omitempty"`             //Product family of the network device. For example, Switches, Routers, etc.
	StackDevice        string `url:"stackDevice,omitempty"`        //Flag indicating if the device is a stack device
	Role               string `url:"role,omitempty"`               //Role assigned to the network device. Available values : BORDER_ROUTER, CORE, DISTRIBUTION, ACCESS, UNKNOWN
	Status             string `url:"status,omitempty"`             //Inventory related status of the network device. Available values : MANAGED, SYNC_NOT_STARTED, SYNC_INIT_FAILED, SYNC_PRECHECK_FAILED, SYNC_IN_PROGRESS, SYNC_INTERNAL_ERROR, SYNC_DISABLED, DELETING_DEVICE, UNDER_MAINTENANCE, QUARANTINED, UNASSOCIATED, UNREACHABLE, UNKNOWN. Refer features for more details.
	ReachabilityStatus string `url:"reachabilityStatus,omitempty"` //Reachability status of the network device. Available values : REACHABLE, ONLY_PING_REACHABLE, UNREACHABLE, UNKNOWN. Refer features for more details.
	ManagementState    string `url:"managementState,omitempty"`    //The status of the network device's manageability. Available statuses are MANAGED, UNDER_MAINTENANCE, NEVER_MANAGED. Refer features for more details.
	Views              string `url:"views,omitempty"`              //The specific views being requested. This is an optional parameter which can be passed to get one or more of the network device data. If this is not provided, then it will default to BASIC views. If multiple views are provided, the response will contain the union of the views. Refer features for more details. Available values : BASIC, RESYNC, USER_DEFINED_FIELDS.
	Limit              string `url:"limit,omitempty"`              //The number of records to show for this page. Min: 1, Max: 500
	Offset             string `url:"offset,omitempty"`             //The first record to show for this page; the first record is numbered 1.
	SortBy             string `url:"sortBy,omitempty"`             //A property within the response to sort by. Available values : id, managementAddress, dnsResolvedManagementIpAddress, hostname, macAddress, type, family, series, platformids, softwareType, softwareVersion, vendor, bootTime, role, roleSource, apEthernetMacAddress, apManagerInterfaceIpAddress, apWlcIpAddress, deviceSupportLevel, reachabilityFailureReason, resyncStartTime, resyncEndTime, resyncReasons, pendingResyncRequestCount, pendingResyncRequestReasons, resyncIntervalSource, resyncIntervalMinutes
	Order              string `url:"order,omitempty"`              //Whether ascending or descending order should be used to sort the response.
}
type CountTheNumberOfNetworkDevicesV1QueryParams struct {
	ID                 string `url:"id,omitempty"`                 //Network device Id
	ManagementAddress  string `url:"managementAddress,omitempty"`  //Management address of the network device
	SerialNumber       string `url:"serialNumber,omitempty"`       //Serial number of the network device
	Family             string `url:"family,omitempty"`             //Product family of the network device. For example, Switches, Routers, etc.
	StackDevice        string `url:"stackDevice,omitempty"`        //Flag indicating if the device is a stack device
	Role               string `url:"role,omitempty"`               //Role assigned to the network device. Available values : BORDER_ROUTER, CORE, DISTRIBUTION, ACCESS, UNKNOWN
	Status             string `url:"status,omitempty"`             //Inventory related status of the network device. Available values : MANAGED, SYNC_NOT_STARTED, SYNC_INIT_FAILED, SYNC_PRECHECK_FAILED, SYNC_IN_PROGRESS, SYNC_INTERNAL_ERROR, SYNC_DISABLED, DELETING_DEVICE, UNDER_MAINTENANCE, QUARANTINED, UNASSOCIATED, UNREACHABLE, UNKNOWN. Refer features for more details.
	ReachabilityStatus string `url:"reachabilityStatus,omitempty"` //Reachability status of the network device. Available values : REACHABLE, ONLY_PING_REACHABLE, UNREACHABLE, UNKNOWN. Refer features for more details.
	ManagementState    string `url:"managementState,omitempty"`    //The status of the network device's manageability. Available values : MANAGED, UNDER_MAINTENANCE, NEVER_MANAGED. Refer features for more details.
}
type GetDetailsOfASingleNetworkDeviceV1QueryParams struct {
	Views string `url:"views,omitempty"` //The specific views being requested. This is an optional parameter which can be passed to get one or more of the network device data. If this is not provided, then it will default to BASIC views. If multiple views are provided, the response will contain the union of the views. Available values : BASIC, RESYNC, USER_DEFINED_FIELDS
}
type GetAllowedMacAddressV1QueryParams struct {
	Offset float64 `url:"offset,omitempty"` //The offset of the first item in the collection to return.
	Limit  float64 `url:"limit,omitempty"`  //The maximum number of entries to return. If the value exceeds the total count, then the maximum entries will be returned.
}

type ResponseDevicesRetrievesTheListOfAAAServicesForGivenParametersV1 struct {
	Response *[]ResponseDevicesRetrievesTheListOfAAAServicesForGivenParametersV1Response `json:"response,omitempty"` //
	Page     *ResponseDevicesRetrievesTheListOfAAAServicesForGivenParametersV1Page       `json:"page,omitempty"`     //
	Version  string                                                                      `json:"version,omitempty"`  // Version
}
type ResponseDevicesRetrievesTheListOfAAAServicesForGivenParametersV1Response struct {
	ID                        string `json:"id,omitempty"`                        // Id
	ServerIP                  string `json:"serverIp,omitempty"`                  // Server Ip
	DeviceID                  string `json:"deviceId,omitempty"`                  // Device Id
	DeviceName                string `json:"deviceName,omitempty"`                // Device Name
	DeviceFamily              string `json:"deviceFamily,omitempty"`              // Device Family
	DeviceSiteHierarchy       string `json:"deviceSiteHierarchy,omitempty"`       // Device Site Hierarchy
	DeviceSiteID              string `json:"deviceSiteId,omitempty"`              // Device Site Id
	DeviceSiteHierarchyID     string `json:"deviceSiteHierarchyId,omitempty"`     // Device Site Hierarchy Id
	Transactions              *int   `json:"transactions,omitempty"`              // Transactions
	FailedTransactions        *int   `json:"failedTransactions,omitempty"`        // Failed Transactions
	SuccessfulTransactions    *int   `json:"successfulTransactions,omitempty"`    // Successful Transactions
	EapTransactions           *int   `json:"eapTransactions,omitempty"`           // Eap Transactions
	EapFailedTransactions     *int   `json:"eapFailedTransactions,omitempty"`     // Eap Failed Transactions
	EapSuccessfulTransactions *int   `json:"eapSuccessfulTransactions,omitempty"` // Eap Successful Transactions
	MabTransactions           *int   `json:"mabTransactions,omitempty"`           // Mab Transactions
	MabFailedTransactions     *int   `json:"mabFailedTransactions,omitempty"`     // Mab Failed Transactions
	MabSuccessfulTransactions *int   `json:"mabSuccessfulTransactions,omitempty"` // Mab Successful Transactions
	Latency                   *int   `json:"latency,omitempty"`                   // Latency
	EapLatency                *int   `json:"eapLatency,omitempty"`                // Eap Latency
	MabLatency                *int   `json:"mabLatency,omitempty"`                // Mab Latency
}
type ResponseDevicesRetrievesTheListOfAAAServicesForGivenParametersV1Page struct {
	Limit  *int                                                                          `json:"limit,omitempty"`  // Limit
	Offset *int                                                                          `json:"offset,omitempty"` // Offset
	Count  *int                                                                          `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesRetrievesTheListOfAAAServicesForGivenParametersV1PageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesRetrievesTheListOfAAAServicesForGivenParametersV1PageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseDevicesRetrievesTheTotalNumberOfAAAServicesForGivenParametersV1 struct {
	Response *ResponseDevicesRetrievesTheTotalNumberOfAAAServicesForGivenParametersV1Response `json:"response,omitempty"` //
	Version  string                                                                           `json:"version,omitempty"`  // Version
}
type ResponseDevicesRetrievesTheTotalNumberOfAAAServicesForGivenParametersV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1 struct {
	Response *[]ResponseDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1Response `json:"response,omitempty"` //
	Page     *ResponseDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1Page       `json:"page,omitempty"`     //
	Version  string                                                                               `json:"version,omitempty"`  // Version
}
type ResponseDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1Response struct {
	ID                        string `json:"id,omitempty"`                        // Id
	ServerIP                  string `json:"serverIp,omitempty"`                  // Server Ip
	DeviceID                  string `json:"deviceId,omitempty"`                  // Device Id
	DeviceName                string `json:"deviceName,omitempty"`                // Device Name
	DeviceFamily              string `json:"deviceFamily,omitempty"`              // Device Family
	DeviceSiteHierarchy       string `json:"deviceSiteHierarchy,omitempty"`       // Device Site Hierarchy
	DeviceSiteID              string `json:"deviceSiteId,omitempty"`              // Device Site Id
	DeviceSiteHierarchyID     string `json:"deviceSiteHierarchyId,omitempty"`     // Device Site Hierarchy Id
	Transactions              *int   `json:"transactions,omitempty"`              // Transactions
	FailedTransactions        *int   `json:"failedTransactions,omitempty"`        // Failed Transactions
	SuccessfulTransactions    *int   `json:"successfulTransactions,omitempty"`    // Successful Transactions
	EapTransactions           *int   `json:"eapTransactions,omitempty"`           // Eap Transactions
	EapFailedTransactions     *int   `json:"eapFailedTransactions,omitempty"`     // Eap Failed Transactions
	EapSuccessfulTransactions *int   `json:"eapSuccessfulTransactions,omitempty"` // Eap Successful Transactions
	MabTransactions           *int   `json:"mabTransactions,omitempty"`           // Mab Transactions
	MabFailedTransactions     *int   `json:"mabFailedTransactions,omitempty"`     // Mab Failed Transactions
	MabSuccessfulTransactions *int   `json:"mabSuccessfulTransactions,omitempty"` // Mab Successful Transactions
	Latency                   *int   `json:"latency,omitempty"`                   // Latency
	EapLatency                *int   `json:"eapLatency,omitempty"`                // Eap Latency
	MabLatency                *int   `json:"mabLatency,omitempty"`                // Mab Latency
}
type ResponseDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1Page struct {
	Limit  *int                                                                                   `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                   `json:"offset,omitempty"` // Offset
	Count  *int                                                                                   `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1PageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1PageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseDevicesRetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersV1 struct {
	Response *ResponseDevicesRetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersV1Response `json:"response,omitempty"` //
	Version  string                                                                                    `json:"version,omitempty"`  // Version
}
type ResponseDevicesRetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1 struct {
	Version  string                                                                                    `json:"version,omitempty"`  // Version
	Response *ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Response `json:"response,omitempty"` //
	Page     *ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Page     `json:"page,omitempty"`     //
}
type ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Response struct {
	Groups              *[]ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ResponseGroups              `json:"groups,omitempty"`              //
	Attributes          *[]ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ResponseGroups struct {
	ID                  string                                                                                                               `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ResponseGroupsAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ResponseGroupsAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ResponseGroupsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ResponseAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Page struct {
	Limit  *int                                                                                          `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                          `json:"offset,omitempty"` // Offset
	Count  string                                                                                        `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1PageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1PageSortBy struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Order    string `json:"order,omitempty"`    // Order
}
type ResponseDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1 struct {
	Version  string                                                                                   `json:"version,omitempty"`  // Version
	Response *[]ResponseDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Response `json:"response,omitempty"` //
	Page     *ResponseDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Page       `json:"page,omitempty"`     //
}
type ResponseDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Response struct {
	ID                  string                                                                                                      `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ResponseAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Page struct {
	Limit  *int                                                                                       `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                       `json:"offset,omitempty"` // Offset
	Count  string                                                                                     `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1PageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1PageSortBy struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Order    string `json:"order,omitempty"`    // Order
}
type ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1 struct {
	Version  string                                                                                    `json:"version,omitempty"`  // Version
	Response *[]ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Response `json:"response,omitempty"` //
	Page     *ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Page       `json:"page,omitempty"`     //
}
type ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Response struct {
	Timestamp           *int                                                                                                         `json:"timestamp,omitempty"`           // Timestamp
	Groups              *[]ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ResponseGroups              `json:"groups,omitempty"`              //
	Attributes          *[]ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ResponseGroups struct {
	ID                  string                                                                                                             `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ResponseGroupsAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ResponseGroupsAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ResponseGroupsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ResponseAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Page struct {
	Limit          *int   `json:"limit,omitempty"`          // Limit
	Offset         *int   `json:"offset,omitempty"`         // Offset
	Count          *int   `json:"count,omitempty"`          // Count
	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type ResponseDevicesRetrievesTheDetailsOfASpecificAAAServiceMatchingTheIDOfTheServiceV1 struct {
	Response *ResponseDevicesRetrievesTheDetailsOfASpecificAAAServiceMatchingTheIDOfTheServiceV1Response `json:"response,omitempty"` //
	Version  string                                                                                      `json:"version,omitempty"`  // Version
}
type ResponseDevicesRetrievesTheDetailsOfASpecificAAAServiceMatchingTheIDOfTheServiceV1Response struct {
	ID                        string `json:"id,omitempty"`                        // Id
	ServerIP                  string `json:"serverIp,omitempty"`                  // Server Ip
	DeviceID                  string `json:"deviceId,omitempty"`                  // Device Id
	DeviceName                string `json:"deviceName,omitempty"`                // Device Name
	DeviceFamily              string `json:"deviceFamily,omitempty"`              // Device Family
	DeviceSiteHierarchy       string `json:"deviceSiteHierarchy,omitempty"`       // Device Site Hierarchy
	DeviceSiteID              string `json:"deviceSiteId,omitempty"`              // Device Site Id
	DeviceSiteHierarchyID     string `json:"deviceSiteHierarchyId,omitempty"`     // Device Site Hierarchy Id
	Transactions              *int   `json:"transactions,omitempty"`              // Transactions
	FailedTransactions        *int   `json:"failedTransactions,omitempty"`        // Failed Transactions
	SuccessfulTransactions    *int   `json:"successfulTransactions,omitempty"`    // Successful Transactions
	EapTransactions           *int   `json:"eapTransactions,omitempty"`           // Eap Transactions
	EapFailedTransactions     *int   `json:"eapFailedTransactions,omitempty"`     // Eap Failed Transactions
	EapSuccessfulTransactions *int   `json:"eapSuccessfulTransactions,omitempty"` // Eap Successful Transactions
	MabTransactions           *int   `json:"mabTransactions,omitempty"`           // Mab Transactions
	MabFailedTransactions     *int   `json:"mabFailedTransactions,omitempty"`     // Mab Failed Transactions
	MabSuccessfulTransactions *int   `json:"mabSuccessfulTransactions,omitempty"` // Mab Successful Transactions
	Latency                   *int   `json:"latency,omitempty"`                   // Latency
	EapLatency                *int   `json:"eapLatency,omitempty"`                // Eap Latency
	MabLatency                *int   `json:"mabLatency,omitempty"`                // Mab Latency
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1 struct {
	Version  string                                                                                        `json:"version,omitempty"`  // Version
	Response *[]ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1Response `json:"response,omitempty"` //
	Page     *ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1Page       `json:"page,omitempty"`     //
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1Response struct {
	Timestamp *int `json:"timestamp,omitempty"` // Timestamp

	Groups *[]ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1ResponseGroups `json:"groups,omitempty"` //

	Attributes *[]ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1ResponseAttributes `json:"attributes,omitempty"` //

	AggregateAttributes *[]ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1ResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1ResponseGroups struct {
	ID string `json:"id,omitempty"` // Id

	Attributes *[]ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1ResponseGroupsAttributes `json:"attributes,omitempty"` //

	AggregateAttributes *[]ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1ResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1ResponseGroupsAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1ResponseGroupsAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function

	Value *int `json:"value,omitempty"` // Value
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1ResponseAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1ResponseAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function

	Value *int `json:"value,omitempty"` // Value
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	Count *int `json:"count,omitempty"` // Count

	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type ResponseDevicesQueryAssuranceEventsV1 struct {
	Response *[]ResponseDevicesQueryAssuranceEventsV1Response `json:"response,omitempty"` //
	Version  string                                           `json:"version,omitempty"`  // Version
	Page     *ResponseDevicesQueryAssuranceEventsV1Page       `json:"page,omitempty"`     //
}
type ResponseDevicesQueryAssuranceEventsV1Response struct {
	OldRadioChannelWidth         string                                                             `json:"oldRadioChannelWidth,omitempty"`         // Old Radio Channel Width
	ClientMac                    string                                                             `json:"clientMac,omitempty"`                    // Client Mac
	SwitchNumber                 string                                                             `json:"switchNumber,omitempty"`                 // Switch Number
	AssocRssi                    *int                                                               `json:"assocRssi,omitempty"`                    // Assoc Rssi
	AffectedClients              []string                                                           `json:"affectedClients,omitempty"`              // Affected Clients
	IsPrivateMac                 *bool                                                              `json:"isPrivateMac,omitempty"`                 // Is Private Mac
	Frequency                    string                                                             `json:"frequency,omitempty"`                    // Frequency
	ApRole                       string                                                             `json:"apRole,omitempty"`                       // Ap Role
	ReplacingDeviceSerialNumber  string                                                             `json:"replacingDeviceSerialNumber,omitempty"`  // Replacing Device Serial Number
	MessageType                  string                                                             `json:"messageType,omitempty"`                  // Message Type
	FailureCategory              string                                                             `json:"failureCategory,omitempty"`              // Failure Category
	ApSwitchName                 string                                                             `json:"apSwitchName,omitempty"`                 // Ap Switch Name
	ApSwitchID                   string                                                             `json:"apSwitchId,omitempty"`                   // Ap Switch Id
	RadioChannelUtilization      string                                                             `json:"radioChannelUtilization,omitempty"`      // Radio Channel Utilization
	Mnemonic                     string                                                             `json:"mnemonic,omitempty"`                     // Mnemonic
	RadioChannelSlot             *int                                                               `json:"radioChannelSlot,omitempty"`             // Radio Channel Slot
	Details                      string                                                             `json:"details,omitempty"`                      // Details
	ID                           string                                                             `json:"id,omitempty"`                           // Id
	LastApDisconnectReason       string                                                             `json:"lastApDisconnectReason,omitempty"`       // Last Ap Disconnect Reason
	NetworkDeviceName            string                                                             `json:"networkDeviceName,omitempty"`            // Network Device Name
	IDentifier                   string                                                             `json:"identifier,omitempty"`                   // Identifier
	ReasonDescription            string                                                             `json:"reasonDescription,omitempty"`            // Reason Description
	VLANID                       string                                                             `json:"vlanId,omitempty"`                       // Vlan Id
	UdnID                        string                                                             `json:"udnId,omitempty"`                        // Udn Id
	AuditSessionID               string                                                             `json:"auditSessionId,omitempty"`               // Audit Session Id
	ApMac                        string                                                             `json:"apMac,omitempty"`                        // Ap Mac
	DeviceFamily                 string                                                             `json:"deviceFamily,omitempty"`                 // Device Family
	RadioNoise                   string                                                             `json:"radioNoise,omitempty"`                   // Radio Noise
	WlcName                      string                                                             `json:"wlcName,omitempty"`                      // Wlc Name
	ApRadioOperationState        string                                                             `json:"apRadioOperationState,omitempty"`        // Ap Radio Operation State
	Name                         string                                                             `json:"name,omitempty"`                         // Name
	FailureIPAddress             string                                                             `json:"failureIpAddress,omitempty"`             // Failure Ip Address
	NewRadioChannelList          string                                                             `json:"newRadioChannelList,omitempty"`          // New Radio Channel List
	Duid                         string                                                             `json:"duid,omitempty"`                         // Duid
	RoamType                     string                                                             `json:"roamType,omitempty"`                     // Roam Type
	CandidateAPs                 *[]ResponseDevicesQueryAssuranceEventsV1ResponseCandidateAPs       `json:"candidateAPs,omitempty"`                 //
	ReplacedDeviceSerialNumber   string                                                             `json:"replacedDeviceSerialNumber,omitempty"`   // Replaced Device Serial Number
	OldRadioChannelList          string                                                             `json:"oldRadioChannelList,omitempty"`          // Old Radio Channel List
	SSID                         string                                                             `json:"ssid,omitempty"`                         // Ssid
	SubReasonDescription         string                                                             `json:"subReasonDescription,omitempty"`         // Sub Reason Description
	WirelessClientEventEndTime   *int                                                               `json:"wirelessClientEventEndTime,omitempty"`   // Wireless Client Event End Time
	IPv4                         string                                                             `json:"ipv4,omitempty"`                         // Ipv4
	WlcID                        string                                                             `json:"wlcId,omitempty"`                        // Wlc Id
	IPv6                         string                                                             `json:"ipv6,omitempty"`                         // Ipv6
	MissingResponseAPs           *[]ResponseDevicesQueryAssuranceEventsV1ResponseMissingResponseAPs `json:"missingResponseAPs,omitempty"`           //
	Timestamp                    *int                                                               `json:"timestamp,omitempty"`                    // Timestamp
	Severity                     *int                                                               `json:"severity,omitempty"`                     // Severity
	CurrentRadioPowerLevel       *int                                                               `json:"currentRadioPowerLevel,omitempty"`       // Current Radio Power Level
	NewRadioChannelWidth         string                                                             `json:"newRadioChannelWidth,omitempty"`         // New Radio Channel Width
	AssocSnr                     *int                                                               `json:"assocSnr,omitempty"`                     // Assoc Snr
	AuthServerIP                 string                                                             `json:"authServerIp,omitempty"`                 // Auth Server Ip
	ChildEvents                  *[]ResponseDevicesQueryAssuranceEventsV1ResponseChildEvents        `json:"childEvents,omitempty"`                  //
	ConnectedInterfaceName       string                                                             `json:"connectedInterfaceName,omitempty"`       // Connected Interface Name
	DhcpServerIP                 string                                                             `json:"dhcpServerIp,omitempty"`                 // Dhcp Server Ip
	ManagementIPAddress          string                                                             `json:"managementIpAddress,omitempty"`          // Management Ip Address
	PreviousRadioPowerLevel      *int                                                               `json:"previousRadioPowerLevel,omitempty"`      // Previous Radio Power Level
	ResultStatus                 string                                                             `json:"resultStatus,omitempty"`                 // Result Status
	RadioInterference            string                                                             `json:"radioInterference,omitempty"`            // Radio Interference
	NetworkDeviceID              string                                                             `json:"networkDeviceId,omitempty"`              // Network Device Id
	SiteHierarchy                string                                                             `json:"siteHierarchy,omitempty"`                // Site Hierarchy
	EventStatus                  string                                                             `json:"eventStatus,omitempty"`                  // Event Status
	WirelessClientEventStartTime *int                                                               `json:"wirelessClientEventStartTime,omitempty"` // Wireless Client Event Start Time
	SiteHierarchyID              string                                                             `json:"siteHierarchyId,omitempty"`              // Site Hierarchy Id
	UdnName                      string                                                             `json:"udnName,omitempty"`                      // Udn Name
	Facility                     string                                                             `json:"facility,omitempty"`                     // Facility
	LastApResetType              string                                                             `json:"lastApResetType,omitempty"`              // Last Ap Reset Type
	InvalidIeAPs                 *[]ResponseDevicesQueryAssuranceEventsV1ResponseInvalidIeAPs       `json:"invalidIeAPs,omitempty"`                 //
	Username                     string                                                             `json:"username,omitempty"`                     // Username
}
type ResponseDevicesQueryAssuranceEventsV1ResponseCandidateAPs struct {
	APID   string `json:"apId,omitempty"`   // Ap Id
	ApName string `json:"apName,omitempty"` // Ap Name
	ApMac  string `json:"apMac,omitempty"`  // Ap Mac
	Bssid  string `json:"bssid,omitempty"`  // Bssid
	Rssi   *int   `json:"rssi,omitempty"`   // Rssi
}
type ResponseDevicesQueryAssuranceEventsV1ResponseMissingResponseAPs struct {
	APID      string `json:"apId,omitempty"`      // Ap Id
	ApName    string `json:"apName,omitempty"`    // Ap Name
	ApMac     string `json:"apMac,omitempty"`     // Ap Mac
	Bssid     string `json:"bssid,omitempty"`     // Bssid
	Type      string `json:"type,omitempty"`      // Type
	FrameType string `json:"frameType,omitempty"` // Frame Type
}
type ResponseDevicesQueryAssuranceEventsV1ResponseChildEvents struct {
	ID                   string `json:"id,omitempty"`                   // Id
	Name                 string `json:"name,omitempty"`                 // Name
	Timestamp            *int   `json:"timestamp,omitempty"`            // Timestamp
	WirelessEventType    *int   `json:"wirelessEventType,omitempty"`    // Wireless Event Type
	Details              string `json:"details,omitempty"`              // Details
	ReasonCode           string `json:"reasonCode,omitempty"`           // Reason Code
	ReasonDescription    string `json:"reasonDescription,omitempty"`    // Reason Description
	SubReasonCode        string `json:"subReasonCode,omitempty"`        // Sub Reason Code
	SubReasonDescription string `json:"subReasonDescription,omitempty"` // Sub Reason Description
	ResultStatus         string `json:"resultStatus,omitempty"`         // Result Status
	FailureCategory      string `json:"failureCategory,omitempty"`      // Failure Category
}
type ResponseDevicesQueryAssuranceEventsV1ResponseInvalidIeAPs struct {
	APID      string `json:"apId,omitempty"`      // Ap Id
	ApName    string `json:"apName,omitempty"`    // Ap Name
	ApMac     string `json:"apMac,omitempty"`     // Ap Mac
	Bssid     string `json:"bssid,omitempty"`     // Bssid
	Type      string `json:"type,omitempty"`      // Type
	FrameType string `json:"frameType,omitempty"` // Frame Type
	Ies       string `json:"ies,omitempty"`       // Ies
}
type ResponseDevicesQueryAssuranceEventsV1Page struct {
	Limit  *int                                               `json:"limit,omitempty"`  // Limit
	Offset *int                                               `json:"offset,omitempty"` // Offset
	Count  *int                                               `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesQueryAssuranceEventsV1PageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesQueryAssuranceEventsV1PageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseDevicesCountTheNumberOfEventsV1 struct {
	Response *ResponseDevicesCountTheNumberOfEventsV1Response `json:"response,omitempty"` //
	Version  string                                           `json:"version,omitempty"`  // Version
}
type ResponseDevicesCountTheNumberOfEventsV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseDevicesQueryAssuranceEventsWithFiltersV1 struct {
	Response *[]ResponseDevicesQueryAssuranceEventsWithFiltersV1Response `json:"response,omitempty"` //
	Version  string                                                      `json:"version,omitempty"`  // Version
	Page     *ResponseDevicesQueryAssuranceEventsWithFiltersV1Page       `json:"page,omitempty"`     //
}
type ResponseDevicesQueryAssuranceEventsWithFiltersV1Response struct {
	OldRadioChannelWidth         string                                                                        `json:"oldRadioChannelWidth,omitempty"`         // Old Radio Channel Width
	ClientMac                    string                                                                        `json:"clientMac,omitempty"`                    // Client Mac
	SwitchNumber                 string                                                                        `json:"switchNumber,omitempty"`                 // Switch Number
	AssocRssi                    *int                                                                          `json:"assocRssi,omitempty"`                    // Assoc Rssi
	AffectedClients              []string                                                                      `json:"affectedClients,omitempty"`              // Affected Clients
	IsPrivateMac                 *bool                                                                         `json:"isPrivateMac,omitempty"`                 // Is Private Mac
	Frequency                    string                                                                        `json:"frequency,omitempty"`                    // Frequency
	ApRole                       string                                                                        `json:"apRole,omitempty"`                       // Ap Role
	ReplacingDeviceSerialNumber  string                                                                        `json:"replacingDeviceSerialNumber,omitempty"`  // Replacing Device Serial Number
	MessageType                  string                                                                        `json:"messageType,omitempty"`                  // Message Type
	FailureCategory              string                                                                        `json:"failureCategory,omitempty"`              // Failure Category
	ApSwitchName                 string                                                                        `json:"apSwitchName,omitempty"`                 // Ap Switch Name
	ApSwitchID                   string                                                                        `json:"apSwitchId,omitempty"`                   // Ap Switch Id
	RadioChannelUtilization      string                                                                        `json:"radioChannelUtilization,omitempty"`      // Radio Channel Utilization
	Mnemonic                     string                                                                        `json:"mnemonic,omitempty"`                     // Mnemonic
	RadioChannelSlot             *int                                                                          `json:"radioChannelSlot,omitempty"`             // Radio Channel Slot
	Details                      string                                                                        `json:"details,omitempty"`                      // Details
	ID                           string                                                                        `json:"id,omitempty"`                           // Id
	LastApDisconnectReason       string                                                                        `json:"lastApDisconnectReason,omitempty"`       // Last Ap Disconnect Reason
	NetworkDeviceName            string                                                                        `json:"networkDeviceName,omitempty"`            // Network Device Name
	IDentifier                   string                                                                        `json:"identifier,omitempty"`                   // Identifier
	ReasonDescription            string                                                                        `json:"reasonDescription,omitempty"`            // Reason Description
	VLANID                       string                                                                        `json:"vlanId,omitempty"`                       // Vlan Id
	UdnID                        string                                                                        `json:"udnId,omitempty"`                        // Udn Id
	AuditSessionID               string                                                                        `json:"auditSessionId,omitempty"`               // Audit Session Id
	ApMac                        string                                                                        `json:"apMac,omitempty"`                        // Ap Mac
	DeviceFamily                 string                                                                        `json:"deviceFamily,omitempty"`                 // Device Family
	RadioNoise                   string                                                                        `json:"radioNoise,omitempty"`                   // Radio Noise
	WlcName                      string                                                                        `json:"wlcName,omitempty"`                      // Wlc Name
	ApRadioOperationState        string                                                                        `json:"apRadioOperationState,omitempty"`        // Ap Radio Operation State
	Name                         string                                                                        `json:"name,omitempty"`                         // Name
	FailureIPAddress             string                                                                        `json:"failureIpAddress,omitempty"`             // Failure Ip Address
	NewRadioChannelList          string                                                                        `json:"newRadioChannelList,omitempty"`          // New Radio Channel List
	Duid                         string                                                                        `json:"duid,omitempty"`                         // Duid
	RoamType                     string                                                                        `json:"roamType,omitempty"`                     // Roam Type
	CandidateAPs                 *[]ResponseDevicesQueryAssuranceEventsWithFiltersV1ResponseCandidateAPs       `json:"candidateAPs,omitempty"`                 //
	ReplacedDeviceSerialNumber   string                                                                        `json:"replacedDeviceSerialNumber,omitempty"`   // Replaced Device Serial Number
	OldRadioChannelList          string                                                                        `json:"oldRadioChannelList,omitempty"`          // Old Radio Channel List
	SSID                         string                                                                        `json:"ssid,omitempty"`                         // Ssid
	SubReasonDescription         string                                                                        `json:"subReasonDescription,omitempty"`         // Sub Reason Description
	WirelessClientEventEndTime   *int                                                                          `json:"wirelessClientEventEndTime,omitempty"`   // Wireless Client Event End Time
	IPv4                         string                                                                        `json:"ipv4,omitempty"`                         // Ipv4
	WlcID                        string                                                                        `json:"wlcId,omitempty"`                        // Wlc Id
	IPv6                         string                                                                        `json:"ipv6,omitempty"`                         // Ipv6
	MissingResponseAPs           *[]ResponseDevicesQueryAssuranceEventsWithFiltersV1ResponseMissingResponseAPs `json:"missingResponseAPs,omitempty"`           //
	Timestamp                    *int                                                                          `json:"timestamp,omitempty"`                    // Timestamp
	Severity                     *int                                                                          `json:"severity,omitempty"`                     // Severity
	CurrentRadioPowerLevel       *int                                                                          `json:"currentRadioPowerLevel,omitempty"`       // Current Radio Power Level
	NewRadioChannelWidth         string                                                                        `json:"newRadioChannelWidth,omitempty"`         // New Radio Channel Width
	AssocSnr                     *int                                                                          `json:"assocSnr,omitempty"`                     // Assoc Snr
	AuthServerIP                 string                                                                        `json:"authServerIp,omitempty"`                 // Auth Server Ip
	ChildEvents                  *[]ResponseDevicesQueryAssuranceEventsWithFiltersV1ResponseChildEvents        `json:"childEvents,omitempty"`                  //
	ConnectedInterfaceName       string                                                                        `json:"connectedInterfaceName,omitempty"`       // Connected Interface Name
	DhcpServerIP                 string                                                                        `json:"dhcpServerIp,omitempty"`                 // Dhcp Server Ip
	ManagementIPAddress          string                                                                        `json:"managementIpAddress,omitempty"`          // Management Ip Address
	PreviousRadioPowerLevel      *int                                                                          `json:"previousRadioPowerLevel,omitempty"`      // Previous Radio Power Level
	ResultStatus                 string                                                                        `json:"resultStatus,omitempty"`                 // Result Status
	RadioInterference            string                                                                        `json:"radioInterference,omitempty"`            // Radio Interference
	NetworkDeviceID              string                                                                        `json:"networkDeviceId,omitempty"`              // Network Device Id
	SiteHierarchy                string                                                                        `json:"siteHierarchy,omitempty"`                // Site Hierarchy
	EventStatus                  string                                                                        `json:"eventStatus,omitempty"`                  // Event Status
	WirelessClientEventStartTime *int                                                                          `json:"wirelessClientEventStartTime,omitempty"` // Wireless Client Event Start Time
	SiteHierarchyID              string                                                                        `json:"siteHierarchyId,omitempty"`              // Site Hierarchy Id
	UdnName                      string                                                                        `json:"udnName,omitempty"`                      // Udn Name
	Facility                     string                                                                        `json:"facility,omitempty"`                     // Facility
	LastApResetType              string                                                                        `json:"lastApResetType,omitempty"`              // Last Ap Reset Type
	InvalidIeAPs                 *[]ResponseDevicesQueryAssuranceEventsWithFiltersV1ResponseInvalidIeAPs       `json:"invalidIeAPs,omitempty"`                 //
	Username                     string                                                                        `json:"username,omitempty"`                     // Username
}
type ResponseDevicesQueryAssuranceEventsWithFiltersV1ResponseCandidateAPs struct {
	APID   string `json:"apId,omitempty"`   // Ap Id
	ApName string `json:"apName,omitempty"` // Ap Name
	ApMac  string `json:"apMac,omitempty"`  // Ap Mac
	Bssid  string `json:"bssid,omitempty"`  // Bssid
	Rssi   *int   `json:"rssi,omitempty"`   // Rssi
}
type ResponseDevicesQueryAssuranceEventsWithFiltersV1ResponseMissingResponseAPs struct {
	APID      string `json:"apId,omitempty"`      // Ap Id
	ApName    string `json:"apName,omitempty"`    // Ap Name
	ApMac     string `json:"apMac,omitempty"`     // Ap Mac
	Bssid     string `json:"bssid,omitempty"`     // Bssid
	Type      string `json:"type,omitempty"`      // Type
	FrameType string `json:"frameType,omitempty"` // Frame Type
}
type ResponseDevicesQueryAssuranceEventsWithFiltersV1ResponseChildEvents struct {
	ID                   string `json:"id,omitempty"`                   // Id
	Name                 string `json:"name,omitempty"`                 // Name
	Timestamp            *int   `json:"timestamp,omitempty"`            // Timestamp
	WirelessEventType    *int   `json:"wirelessEventType,omitempty"`    // Wireless Event Type
	Details              string `json:"details,omitempty"`              // Details
	ReasonCode           string `json:"reasonCode,omitempty"`           // Reason Code
	ReasonDescription    string `json:"reasonDescription,omitempty"`    // Reason Description
	SubReasonCode        string `json:"subReasonCode,omitempty"`        // Sub Reason Code
	SubReasonDescription string `json:"subReasonDescription,omitempty"` // Sub Reason Description
	ResultStatus         string `json:"resultStatus,omitempty"`         // Result Status
	FailureCategory      string `json:"failureCategory,omitempty"`      // Failure Category
}
type ResponseDevicesQueryAssuranceEventsWithFiltersV1ResponseInvalidIeAPs struct {
	APID      string `json:"apId,omitempty"`      // Ap Id
	ApName    string `json:"apName,omitempty"`    // Ap Name
	ApMac     string `json:"apMac,omitempty"`     // Ap Mac
	Bssid     string `json:"bssid,omitempty"`     // Bssid
	Type      string `json:"type,omitempty"`      // Type
	FrameType string `json:"frameType,omitempty"` // Frame Type
	Ies       string `json:"ies,omitempty"`       // Ies
}
type ResponseDevicesQueryAssuranceEventsWithFiltersV1Page struct {
	Limit  *int                                                          `json:"limit,omitempty"`  // Limit
	Offset *int                                                          `json:"offset,omitempty"` // Offset
	Count  *int                                                          `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesQueryAssuranceEventsWithFiltersV1PageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesQueryAssuranceEventsWithFiltersV1PageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseDevicesCountTheNumberOfEventsWithFiltersV1 struct {
	Response *ResponseDevicesCountTheNumberOfEventsWithFiltersV1Response `json:"response,omitempty"` //
	Version  string                                                      `json:"version,omitempty"`  // Version
}
type ResponseDevicesCountTheNumberOfEventsWithFiltersV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseDevicesGetDetailsOfASingleAssuranceEventV1 struct {
	Response *ResponseDevicesGetDetailsOfASingleAssuranceEventV1Response `json:"response,omitempty"` //
	Version  string                                                      `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetDetailsOfASingleAssuranceEventV1Response struct {
	OldRadioChannelWidth         string                                                                          `json:"oldRadioChannelWidth,omitempty"`         // Old Radio Channel Width
	ClientMac                    string                                                                          `json:"clientMac,omitempty"`                    // Client Mac
	SwitchNumber                 string                                                                          `json:"switchNumber,omitempty"`                 // Switch Number
	AssocRssi                    *int                                                                            `json:"assocRssi,omitempty"`                    // Assoc Rssi
	AffectedClients              []string                                                                        `json:"affectedClients,omitempty"`              // Affected Clients
	IsPrivateMac                 *bool                                                                           `json:"isPrivateMac,omitempty"`                 // Is Private Mac
	Frequency                    string                                                                          `json:"frequency,omitempty"`                    // Frequency
	ApRole                       string                                                                          `json:"apRole,omitempty"`                       // Ap Role
	ReplacingDeviceSerialNumber  string                                                                          `json:"replacingDeviceSerialNumber,omitempty"`  // Replacing Device Serial Number
	MessageType                  string                                                                          `json:"messageType,omitempty"`                  // Message Type
	FailureCategory              string                                                                          `json:"failureCategory,omitempty"`              // Failure Category
	ApSwitchName                 string                                                                          `json:"apSwitchName,omitempty"`                 // Ap Switch Name
	ApSwitchID                   string                                                                          `json:"apSwitchId,omitempty"`                   // Ap Switch Id
	RadioChannelUtilization      string                                                                          `json:"radioChannelUtilization,omitempty"`      // Radio Channel Utilization
	Mnemonic                     string                                                                          `json:"mnemonic,omitempty"`                     // Mnemonic
	RadioChannelSlot             *int                                                                            `json:"radioChannelSlot,omitempty"`             // Radio Channel Slot
	Details                      string                                                                          `json:"details,omitempty"`                      // Details
	ID                           string                                                                          `json:"id,omitempty"`                           // Id
	LastApDisconnectReason       string                                                                          `json:"lastApDisconnectReason,omitempty"`       // Last Ap Disconnect Reason
	NetworkDeviceName            string                                                                          `json:"networkDeviceName,omitempty"`            // Network Device Name
	IDentifier                   string                                                                          `json:"identifier,omitempty"`                   // Identifier
	ReasonDescription            string                                                                          `json:"reasonDescription,omitempty"`            // Reason Description
	VLANID                       string                                                                          `json:"vlanId,omitempty"`                       // Vlan Id
	UdnID                        string                                                                          `json:"udnId,omitempty"`                        // Udn Id
	AuditSessionID               string                                                                          `json:"auditSessionId,omitempty"`               // Audit Session Id
	ApMac                        string                                                                          `json:"apMac,omitempty"`                        // Ap Mac
	DeviceFamily                 string                                                                          `json:"deviceFamily,omitempty"`                 // Device Family
	RadioNoise                   string                                                                          `json:"radioNoise,omitempty"`                   // Radio Noise
	WlcName                      string                                                                          `json:"wlcName,omitempty"`                      // Wlc Name
	ApRadioOperationState        string                                                                          `json:"apRadioOperationState,omitempty"`        // Ap Radio Operation State
	Name                         string                                                                          `json:"name,omitempty"`                         // Name
	FailureIPAddress             string                                                                          `json:"failureIpAddress,omitempty"`             // Failure Ip Address
	NewRadioChannelList          string                                                                          `json:"newRadioChannelList,omitempty"`          // New Radio Channel List
	Duid                         string                                                                          `json:"duid,omitempty"`                         // Duid
	RoamType                     string                                                                          `json:"roamType,omitempty"`                     // Roam Type
	CandidateAPs                 *[]ResponseDevicesGetDetailsOfASingleAssuranceEventV1ResponseCandidateAPs       `json:"candidateAPs,omitempty"`                 //
	ReplacedDeviceSerialNumber   string                                                                          `json:"replacedDeviceSerialNumber,omitempty"`   // Replaced Device Serial Number
	OldRadioChannelList          string                                                                          `json:"oldRadioChannelList,omitempty"`          // Old Radio Channel List
	SSID                         string                                                                          `json:"ssid,omitempty"`                         // Ssid
	SubReasonDescription         string                                                                          `json:"subReasonDescription,omitempty"`         // Sub Reason Description
	WirelessClientEventEndTime   *int                                                                            `json:"wirelessClientEventEndTime,omitempty"`   // Wireless Client Event End Time
	IPv4                         string                                                                          `json:"ipv4,omitempty"`                         // Ipv4
	WlcID                        string                                                                          `json:"wlcId,omitempty"`                        // Wlc Id
	IPv6                         string                                                                          `json:"ipv6,omitempty"`                         // Ipv6
	MissingResponseAPs           *[]ResponseDevicesGetDetailsOfASingleAssuranceEventV1ResponseMissingResponseAPs `json:"missingResponseAPs,omitempty"`           //
	Timestamp                    *int                                                                            `json:"timestamp,omitempty"`                    // Timestamp
	Severity                     *int                                                                            `json:"severity,omitempty"`                     // Severity
	CurrentRadioPowerLevel       *int                                                                            `json:"currentRadioPowerLevel,omitempty"`       // Current Radio Power Level
	NewRadioChannelWidth         string                                                                          `json:"newRadioChannelWidth,omitempty"`         // New Radio Channel Width
	AssocSnr                     *int                                                                            `json:"assocSnr,omitempty"`                     // Assoc Snr
	AuthServerIP                 string                                                                          `json:"authServerIp,omitempty"`                 // Auth Server Ip
	ChildEvents                  *[]ResponseDevicesGetDetailsOfASingleAssuranceEventV1ResponseChildEvents        `json:"childEvents,omitempty"`                  //
	ConnectedInterfaceName       string                                                                          `json:"connectedInterfaceName,omitempty"`       // Connected Interface Name
	DhcpServerIP                 string                                                                          `json:"dhcpServerIp,omitempty"`                 // Dhcp Server Ip
	ManagementIPAddress          string                                                                          `json:"managementIpAddress,omitempty"`          // Management Ip Address
	PreviousRadioPowerLevel      *int                                                                            `json:"previousRadioPowerLevel,omitempty"`      // Previous Radio Power Level
	ResultStatus                 string                                                                          `json:"resultStatus,omitempty"`                 // Result Status
	RadioInterference            string                                                                          `json:"radioInterference,omitempty"`            // Radio Interference
	NetworkDeviceID              string                                                                          `json:"networkDeviceId,omitempty"`              // Network Device Id
	SiteHierarchy                string                                                                          `json:"siteHierarchy,omitempty"`                // Site Hierarchy
	EventStatus                  string                                                                          `json:"eventStatus,omitempty"`                  // Event Status
	WirelessClientEventStartTime *int                                                                            `json:"wirelessClientEventStartTime,omitempty"` // Wireless Client Event Start Time
	SiteHierarchyID              string                                                                          `json:"siteHierarchyId,omitempty"`              // Site Hierarchy Id
	UdnName                      string                                                                          `json:"udnName,omitempty"`                      // Udn Name
	Facility                     string                                                                          `json:"facility,omitempty"`                     // Facility
	LastApResetType              string                                                                          `json:"lastApResetType,omitempty"`              // Last Ap Reset Type
	InvalidIeAPs                 *[]ResponseDevicesGetDetailsOfASingleAssuranceEventV1ResponseInvalidIeAPs       `json:"invalidIeAPs,omitempty"`                 //
	Username                     string                                                                          `json:"username,omitempty"`                     // Username
}
type ResponseDevicesGetDetailsOfASingleAssuranceEventV1ResponseCandidateAPs struct {
	APID   string `json:"apId,omitempty"`   // Ap Id
	ApName string `json:"apName,omitempty"` // Ap Name
	ApMac  string `json:"apMac,omitempty"`  // Ap Mac
	Bssid  string `json:"bssid,omitempty"`  // Bssid
	Rssi   *int   `json:"rssi,omitempty"`   // Rssi
}
type ResponseDevicesGetDetailsOfASingleAssuranceEventV1ResponseMissingResponseAPs struct {
	APID      string `json:"apId,omitempty"`      // Ap Id
	ApName    string `json:"apName,omitempty"`    // Ap Name
	ApMac     string `json:"apMac,omitempty"`     // Ap Mac
	Bssid     string `json:"bssid,omitempty"`     // Bssid
	Type      string `json:"type,omitempty"`      // Type
	FrameType string `json:"frameType,omitempty"` // Frame Type
}
type ResponseDevicesGetDetailsOfASingleAssuranceEventV1ResponseChildEvents struct {
	ID                   string `json:"id,omitempty"`                   // Id
	Name                 string `json:"name,omitempty"`                 // Name
	Timestamp            *int   `json:"timestamp,omitempty"`            // Timestamp
	WirelessEventType    *int   `json:"wirelessEventType,omitempty"`    // Wireless Event Type
	Details              string `json:"details,omitempty"`              // Details
	ReasonCode           string `json:"reasonCode,omitempty"`           // Reason Code
	ReasonDescription    string `json:"reasonDescription,omitempty"`    // Reason Description
	SubReasonCode        string `json:"subReasonCode,omitempty"`        // Sub Reason Code
	SubReasonDescription string `json:"subReasonDescription,omitempty"` // Sub Reason Description
	ResultStatus         string `json:"resultStatus,omitempty"`         // Result Status
	FailureCategory      string `json:"failureCategory,omitempty"`      // Failure Category
}
type ResponseDevicesGetDetailsOfASingleAssuranceEventV1ResponseInvalidIeAPs struct {
	APID      string `json:"apId,omitempty"`      // Ap Id
	ApName    string `json:"apName,omitempty"`    // Ap Name
	ApMac     string `json:"apMac,omitempty"`     // Ap Mac
	Bssid     string `json:"bssid,omitempty"`     // Bssid
	Type      string `json:"type,omitempty"`      // Type
	FrameType string `json:"frameType,omitempty"` // Frame Type
	Ies       string `json:"ies,omitempty"`       // Ies
}
type ResponseDevicesGetListOfChildEventsForTheGivenWirelessClientEventV1 struct {
	Response *[]ResponseDevicesGetListOfChildEventsForTheGivenWirelessClientEventV1Response `json:"response,omitempty"` //
	Version  string                                                                         `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetListOfChildEventsForTheGivenWirelessClientEventV1Response struct {
	ID                   string `json:"id,omitempty"`                   // Id
	Name                 string `json:"name,omitempty"`                 // Name
	Timestamp            *int   `json:"timestamp,omitempty"`            // Timestamp
	WirelessEventType    *int   `json:"wirelessEventType,omitempty"`    // Wireless Event Type
	Details              string `json:"details,omitempty"`              // Details
	ReasonCode           string `json:"reasonCode,omitempty"`           // Reason Code
	SubreasonCode        string `json:"subreasonCode,omitempty"`        // Subreason Code
	ResultStatus         string `json:"resultStatus,omitempty"`         // Result Status
	ReasonDescription    string `json:"reasonDescription,omitempty"`    // Reason Description
	SubReasonDescription string `json:"subReasonDescription,omitempty"` // Sub Reason Description
	FailureCategory      string `json:"failureCategory,omitempty"`      // Failure Category
}
type ResponseDevicesRetrievesTheListOfDHCPServicesForGivenParametersV1 struct {
	Response *[]ResponseDevicesRetrievesTheListOfDHCPServicesForGivenParametersV1Response `json:"response,omitempty"` //
	Page     *ResponseDevicesRetrievesTheListOfDHCPServicesForGivenParametersV1Page       `json:"page,omitempty"`     //
	Version  string                                                                       `json:"version,omitempty"`  // Version
}
type ResponseDevicesRetrievesTheListOfDHCPServicesForGivenParametersV1Response struct {
	ID                        string `json:"id,omitempty"`                        // Id
	ServerIP                  string `json:"serverIp,omitempty"`                  // Server Ip
	DeviceID                  string `json:"deviceId,omitempty"`                  // Device Id
	DeviceName                string `json:"deviceName,omitempty"`                // Device Name
	DeviceFamily              string `json:"deviceFamily,omitempty"`              // Device Family
	DeviceSiteHierarchy       string `json:"deviceSiteHierarchy,omitempty"`       // Device Site Hierarchy
	DeviceSiteID              string `json:"deviceSiteId,omitempty"`              // Device Site Id
	DeviceSiteHierarchyID     string `json:"deviceSiteHierarchyId,omitempty"`     // Device Site Hierarchy Id
	Transactions              *int   `json:"transactions,omitempty"`              // Transactions
	FailedTransactions        *int   `json:"failedTransactions,omitempty"`        // Failed Transactions
	SuccessfulTransactions    *int   `json:"successfulTransactions,omitempty"`    // Successful Transactions
	Latency                   *int   `json:"latency,omitempty"`                   // Latency
	DiscoverOfferLatency      *int   `json:"discoverOfferLatency,omitempty"`      // Discover Offer Latency
	RequestAcknowledgeLatency *int   `json:"requestAcknowledgeLatency,omitempty"` // Request Acknowledge Latency
}
type ResponseDevicesRetrievesTheListOfDHCPServicesForGivenParametersV1Page struct {
	Limit  *int                                                                           `json:"limit,omitempty"`  // Limit
	Offset *int                                                                           `json:"offset,omitempty"` // Offset
	Count  *int                                                                           `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesRetrievesTheListOfDHCPServicesForGivenParametersV1PageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesRetrievesTheListOfDHCPServicesForGivenParametersV1PageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenParametersV1 struct {
	Response *ResponseDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenParametersV1Response `json:"response,omitempty"` //
	Version  string                                                                            `json:"version,omitempty"`  // Version
}
type ResponseDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenParametersV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1 struct {
	Response *[]ResponseDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1Response `json:"response,omitempty"` //
	Page     *ResponseDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1Page       `json:"page,omitempty"`     //
	Version  string                                                                                `json:"version,omitempty"`  // Version
}
type ResponseDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1Response struct {
	ID                        string `json:"id,omitempty"`                        // Id
	ServerIP                  string `json:"serverIp,omitempty"`                  // Server Ip
	DeviceID                  string `json:"deviceId,omitempty"`                  // Device Id
	DeviceName                string `json:"deviceName,omitempty"`                // Device Name
	DeviceFamily              string `json:"deviceFamily,omitempty"`              // Device Family
	DeviceSiteHierarchy       string `json:"deviceSiteHierarchy,omitempty"`       // Device Site Hierarchy
	DeviceSiteID              string `json:"deviceSiteId,omitempty"`              // Device Site Id
	DeviceSiteHierarchyID     string `json:"deviceSiteHierarchyId,omitempty"`     // Device Site Hierarchy Id
	Transactions              *int   `json:"transactions,omitempty"`              // Transactions
	FailedTransactions        *int   `json:"failedTransactions,omitempty"`        // Failed Transactions
	SuccessfulTransactions    *int   `json:"successfulTransactions,omitempty"`    // Successful Transactions
	Latency                   *int   `json:"latency,omitempty"`                   // Latency
	DiscoverOfferLatency      *int   `json:"discoverOfferLatency,omitempty"`      // Discover Offer Latency
	RequestAcknowledgeLatency *int   `json:"requestAcknowledgeLatency,omitempty"` // Request Acknowledge Latency
}
type ResponseDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1Page struct {
	Limit  *int                                                                                    `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                    `json:"offset,omitempty"` // Offset
	Count  *int                                                                                    `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1PageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1PageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersV1 struct {
	Response *ResponseDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersV1Response `json:"response,omitempty"` //
	Version  string                                                                                     `json:"version,omitempty"`  // Version
}
type ResponseDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1 struct {
	Version  string                                                                                     `json:"version,omitempty"`  // Version
	Response *ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1Response `json:"response,omitempty"` //
	Page     *ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1Page     `json:"page,omitempty"`     //
}
type ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1Response struct {
	Groups              *[]ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1ResponseGroups              `json:"groups,omitempty"`              //
	Attributes          *[]ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1ResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1ResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1ResponseGroups struct {
	ID                  string                                                                                                                `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1ResponseGroupsAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1ResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1ResponseGroupsAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1ResponseGroupsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1ResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1ResponseAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1Page struct {
	Limit  *int                                                                                           `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                           `json:"offset,omitempty"` // Offset
	Count  string                                                                                         `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1PageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1PageSortBy struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Order    string `json:"order,omitempty"`    // Order
}
type ResponseDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1 struct {
	Version  string                                                                                    `json:"version,omitempty"`  // Version
	Response *[]ResponseDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1Response `json:"response,omitempty"` //
	Page     *ResponseDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1Page       `json:"page,omitempty"`     //
}
type ResponseDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1Response struct {
	ID                  string                                                                                                       `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1ResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1ResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1ResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1ResponseAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1Page struct {
	Limit  *int                                                                                        `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                        `json:"offset,omitempty"` // Offset
	Count  string                                                                                      `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1PageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1PageSortBy struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Order    string `json:"order,omitempty"`    // Order
}
type ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1 struct {
	Version  string                                                                                     `json:"version,omitempty"`  // Version
	Response *[]ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1Response `json:"response,omitempty"` //
	Page     *ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1Page       `json:"page,omitempty"`     //
}
type ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1Response struct {
	Timestamp           *int                                                                                                          `json:"timestamp,omitempty"`           // Timestamp
	Groups              *[]ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1ResponseGroups              `json:"groups,omitempty"`              //
	Attributes          *[]ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1ResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1ResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1ResponseGroups struct {
	ID                  string                                                                                                              `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1ResponseGroupsAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1ResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1ResponseGroupsAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1ResponseGroupsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1ResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1ResponseAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1Page struct {
	Limit          *int   `json:"limit,omitempty"`          // Limit
	Offset         *int   `json:"offset,omitempty"`         // Offset
	Count          *int   `json:"count,omitempty"`          // Count
	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type ResponseDevicesRetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheServiceV1 struct {
	Response *ResponseDevicesRetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheServiceV1Response `json:"response,omitempty"` //
	Version  string                                                                                       `json:"version,omitempty"`  // Version
}
type ResponseDevicesRetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheServiceV1Response struct {
	ID                        string `json:"id,omitempty"`                        // Id
	ServerIP                  string `json:"serverIp,omitempty"`                  // Server Ip
	DeviceID                  string `json:"deviceId,omitempty"`                  // Device Id
	DeviceName                string `json:"deviceName,omitempty"`                // Device Name
	DeviceFamily              string `json:"deviceFamily,omitempty"`              // Device Family
	DeviceSiteHierarchy       string `json:"deviceSiteHierarchy,omitempty"`       // Device Site Hierarchy
	DeviceSiteID              string `json:"deviceSiteId,omitempty"`              // Device Site Id
	DeviceSiteHierarchyID     string `json:"deviceSiteHierarchyId,omitempty"`     // Device Site Hierarchy Id
	Transactions              *int   `json:"transactions,omitempty"`              // Transactions
	FailedTransactions        *int   `json:"failedTransactions,omitempty"`        // Failed Transactions
	SuccessfulTransactions    *int   `json:"successfulTransactions,omitempty"`    // Successful Transactions
	Latency                   *int   `json:"latency,omitempty"`                   // Latency
	DiscoverOfferLatency      *int   `json:"discoverOfferLatency,omitempty"`      // Discover Offer Latency
	RequestAcknowledgeLatency *int   `json:"requestAcknowledgeLatency,omitempty"` // Request Acknowledge Latency
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1 struct {
	Version  string                                                                                         `json:"version,omitempty"`  // Version
	Response *[]ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1Response `json:"response,omitempty"` //
	Page     *ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1Page       `json:"page,omitempty"`     //
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1Response struct {
	Timestamp           *int                                                                                                              `json:"timestamp,omitempty"`           // Timestamp
	Groups              *[]ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1ResponseGroups              `json:"groups,omitempty"`              //
	Attributes          *[]ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1ResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1ResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1ResponseGroups struct {
	ID                  string                                                                                                                  `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1ResponseGroupsAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1ResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1ResponseGroupsAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1ResponseGroupsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1ResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1ResponseAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1Page struct {
	Limit          *int   `json:"limit,omitempty"`          // Limit
	Offset         *int   `json:"offset,omitempty"`         // Offset
	Count          *int   `json:"count,omitempty"`          // Count
	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type ResponseDevicesRetrievesTheListOfDNSServicesForGivenParametersV1 struct {
	Response *[]ResponseDevicesRetrievesTheListOfDNSServicesForGivenParametersV1Response `json:"response,omitempty"` //
	Page     *ResponseDevicesRetrievesTheListOfDNSServicesForGivenParametersV1Page       `json:"page,omitempty"`     //
	Version  string                                                                      `json:"version,omitempty"`  // Version
}
type ResponseDevicesRetrievesTheListOfDNSServicesForGivenParametersV1Response struct {
	ID                     string                                                                              `json:"id,omitempty"`                     // Id
	ServerIP               string                                                                              `json:"serverIp,omitempty"`               // Server Ip
	DeviceID               string                                                                              `json:"deviceId,omitempty"`               // Device Id
	DeviceName             string                                                                              `json:"deviceName,omitempty"`             // Device Name
	DeviceFamily           string                                                                              `json:"deviceFamily,omitempty"`           // Device Family
	DeviceSiteHierarchy    string                                                                              `json:"deviceSiteHierarchy,omitempty"`    // Device Site Hierarchy
	DeviceSiteID           string                                                                              `json:"deviceSiteId,omitempty"`           // Device Site Id
	DeviceSiteHierarchyID  string                                                                              `json:"deviceSiteHierarchyId,omitempty"`  // Device Site Hierarchy Id
	Transactions           *int                                                                                `json:"transactions,omitempty"`           // Transactions
	FailedTransactions     *int                                                                                `json:"failedTransactions,omitempty"`     // Failed Transactions
	Failures               *[]ResponseDevicesRetrievesTheListOfDNSServicesForGivenParametersV1ResponseFailures `json:"failures,omitempty"`               //
	SuccessfulTransactions *int                                                                                `json:"successfulTransactions,omitempty"` // Successful Transactions
	Latency                *int                                                                                `json:"latency,omitempty"`                // Latency
	SSID                   string                                                                              `json:"ssid,omitempty"`                   // Ssid
}
type ResponseDevicesRetrievesTheListOfDNSServicesForGivenParametersV1ResponseFailures struct {
	FailureResponseCode *int   `json:"failureResponseCode,omitempty"` // Failure Response Code
	FailureDescription  string `json:"failureDescription,omitempty"`  // Failure Description
	FailedTransactions  *int   `json:"failedTransactions,omitempty"`  // Failed Transactions
}
type ResponseDevicesRetrievesTheListOfDNSServicesForGivenParametersV1Page struct {
	Limit  *int                                                                          `json:"limit,omitempty"`  // Limit
	Offset *int                                                                          `json:"offset,omitempty"` // Offset
	Count  *int                                                                          `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesRetrievesTheListOfDNSServicesForGivenParametersV1PageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesRetrievesTheListOfDNSServicesForGivenParametersV1PageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseDevicesRetrievesTheTotalNumberOfDNSServicesForGivenParametersV1 struct {
	Response *ResponseDevicesRetrievesTheTotalNumberOfDNSServicesForGivenParametersV1Response `json:"response,omitempty"` //
	Version  string                                                                           `json:"version,omitempty"`  // Version
}
type ResponseDevicesRetrievesTheTotalNumberOfDNSServicesForGivenParametersV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1 struct {
	Response *[]ResponseDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1Response `json:"response,omitempty"` //
	Page     *ResponseDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1Page       `json:"page,omitempty"`     //
	Version  string                                                                               `json:"version,omitempty"`  // Version
}
type ResponseDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1Response struct {
	ID                     string                                                                                       `json:"id,omitempty"`                     // Id
	ServerIP               string                                                                                       `json:"serverIp,omitempty"`               // Server Ip
	DeviceID               string                                                                                       `json:"deviceId,omitempty"`               // Device Id
	DeviceName             string                                                                                       `json:"deviceName,omitempty"`             // Device Name
	DeviceFamily           string                                                                                       `json:"deviceFamily,omitempty"`           // Device Family
	DeviceSiteHierarchy    string                                                                                       `json:"deviceSiteHierarchy,omitempty"`    // Device Site Hierarchy
	DeviceSiteID           string                                                                                       `json:"deviceSiteId,omitempty"`           // Device Site Id
	DeviceSiteHierarchyID  string                                                                                       `json:"deviceSiteHierarchyId,omitempty"`  // Device Site Hierarchy Id
	Transactions           *int                                                                                         `json:"transactions,omitempty"`           // Transactions
	FailedTransactions     *int                                                                                         `json:"failedTransactions,omitempty"`     // Failed Transactions
	Failures               *[]ResponseDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1ResponseFailures `json:"failures,omitempty"`               //
	SuccessfulTransactions *int                                                                                         `json:"successfulTransactions,omitempty"` // Successful Transactions
	Latency                *int                                                                                         `json:"latency,omitempty"`                // Latency
	SSID                   string                                                                                       `json:"ssid,omitempty"`                   // Ssid
}
type ResponseDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1ResponseFailures struct {
	FailureResponseCode *int   `json:"failureResponseCode,omitempty"` // Failure Response Code
	FailureDescription  string `json:"failureDescription,omitempty"`  // Failure Description
	FailedTransactions  *int   `json:"failedTransactions,omitempty"`  // Failed Transactions
}
type ResponseDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1Page struct {
	Limit  *int                                                                                   `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                   `json:"offset,omitempty"` // Offset
	Count  *int                                                                                   `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1PageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1PageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1 struct {
	Response *ResponseDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1Response `json:"response,omitempty"` //
	Version  string                                                                                    `json:"version,omitempty"`  // Version
}
type ResponseDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1 struct {
	Version  string                                                                                    `json:"version,omitempty"`  // Version
	Response *ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1Response `json:"response,omitempty"` //
	Page     *ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1Page     `json:"page,omitempty"`     //
}
type ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1Response struct {
	Groups              *[]ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1ResponseGroups              `json:"groups,omitempty"`              //
	Attributes          *[]ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1ResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1ResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1ResponseGroups struct {
	ID                  string                                                                                                               `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1ResponseGroupsAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1ResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1ResponseGroupsAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1ResponseGroupsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1ResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1ResponseAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1Page struct {
	Limit  *int                                                                                          `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                          `json:"offset,omitempty"` // Offset
	Count  string                                                                                        `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1PageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1PageSortBy struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Order    string `json:"order,omitempty"`    // Order
}
type ResponseDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1 struct {
	Version  string                                                                                   `json:"version,omitempty"`  // Version
	Response *[]ResponseDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1Response `json:"response,omitempty"` //
	Page     *ResponseDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1Page       `json:"page,omitempty"`     //
}
type ResponseDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1Response struct {
	ID                  string                                                                                                      `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1ResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1ResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1ResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1ResponseAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1Page struct {
	Limit  *int                                                                                       `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                       `json:"offset,omitempty"` // Offset
	Count  string                                                                                     `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1PageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1PageSortBy struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Order    string `json:"order,omitempty"`    // Order
}
type ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1 struct {
	Version  string                                                                                    `json:"version,omitempty"`  // Version
	Response *[]ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1Response `json:"response,omitempty"` //
	Page     *ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1Page       `json:"page,omitempty"`     //
}
type ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1Response struct {
	Timestamp           *int                                                                                                         `json:"timestamp,omitempty"`           // Timestamp
	Groups              *[]ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1ResponseGroups              `json:"groups,omitempty"`              //
	Attributes          *[]ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1ResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1ResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1ResponseGroups struct {
	ID                  string                                                                                                             `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1ResponseGroupsAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1ResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1ResponseGroupsAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1ResponseGroupsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1ResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1ResponseAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1Page struct {
	Limit          *int   `json:"limit,omitempty"`          // Limit
	Offset         *int   `json:"offset,omitempty"`         // Offset
	Count          *int   `json:"count,omitempty"`          // Count
	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type ResponseDevicesRetrievesTheDetailsOfASpecificDNSServiceMatchingTheIDOfTheServiceV1 struct {
	Response *ResponseDevicesRetrievesTheDetailsOfASpecificDNSServiceMatchingTheIDOfTheServiceV1Response `json:"response,omitempty"` //
	Version  string                                                                                      `json:"version,omitempty"`  // Version
}
type ResponseDevicesRetrievesTheDetailsOfASpecificDNSServiceMatchingTheIDOfTheServiceV1Response struct {
	ID                     string                                                                                                `json:"id,omitempty"`                     // Id
	ServerIP               string                                                                                                `json:"serverIp,omitempty"`               // Server Ip
	DeviceID               string                                                                                                `json:"deviceId,omitempty"`               // Device Id
	DeviceName             string                                                                                                `json:"deviceName,omitempty"`             // Device Name
	DeviceFamily           string                                                                                                `json:"deviceFamily,omitempty"`           // Device Family
	DeviceSiteHierarchy    string                                                                                                `json:"deviceSiteHierarchy,omitempty"`    // Device Site Hierarchy
	DeviceSiteID           string                                                                                                `json:"deviceSiteId,omitempty"`           // Device Site Id
	DeviceSiteHierarchyID  string                                                                                                `json:"deviceSiteHierarchyId,omitempty"`  // Device Site Hierarchy Id
	Transactions           *int                                                                                                  `json:"transactions,omitempty"`           // Transactions
	FailedTransactions     *int                                                                                                  `json:"failedTransactions,omitempty"`     // Failed Transactions
	Failures               *[]ResponseDevicesRetrievesTheDetailsOfASpecificDNSServiceMatchingTheIDOfTheServiceV1ResponseFailures `json:"failures,omitempty"`               //
	SuccessfulTransactions *int                                                                                                  `json:"successfulTransactions,omitempty"` // Successful Transactions
	Latency                *int                                                                                                  `json:"latency,omitempty"`                // Latency
	SSID                   string                                                                                                `json:"ssid,omitempty"`                   // Ssid
}
type ResponseDevicesRetrievesTheDetailsOfASpecificDNSServiceMatchingTheIDOfTheServiceV1ResponseFailures struct {
	FailureResponseCode *int   `json:"failureResponseCode,omitempty"` // Failure Response Code
	FailureDescription  string `json:"failureDescription,omitempty"`  // Failure Description
	FailedTransactions  *int   `json:"failedTransactions,omitempty"`  // Failed Transactions
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1 struct {
	Version  string                                                                                        `json:"version,omitempty"`  // Version
	Response *[]ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1Response `json:"response,omitempty"` //
	Page     *ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1Page       `json:"page,omitempty"`     //
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1Response struct {
	Timestamp           *int                                                                                                             `json:"timestamp,omitempty"`           // Timestamp
	Groups              *[]ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1ResponseGroups              `json:"groups,omitempty"`              //
	Attributes          *[]ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1ResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1ResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1ResponseGroups struct {
	ID                  string                                                                                                                 `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1ResponseGroupsAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1ResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1ResponseGroupsAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1ResponseGroupsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1ResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1ResponseAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1Page struct {
	Limit          *int   `json:"limit,omitempty"`          // Limit
	Offset         *int   `json:"offset,omitempty"`         // Offset
	Count          *int   `json:"count,omitempty"`          // Count
	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type ResponseDevicesGetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesV1 struct {
	Response *[]ResponseDevicesGetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesV1Response `json:"response,omitempty"` //
	Page     *ResponseDevicesGetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesV1Page       `json:"page,omitempty"`     //
	Version  string                                                                                       `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesV1Response struct {
	ID                      string   `json:"id,omitempty"`                      // Id
	AdminStatus             string   `json:"adminStatus,omitempty"`             // Admin Status
	Description             string   `json:"description,omitempty"`             // Description
	DuplexConfig            string   `json:"duplexConfig,omitempty"`            // Duplex Config
	DuplexOper              string   `json:"duplexOper,omitempty"`              // Duplex Oper
	InterfaceIfIndex        *int     `json:"interfaceIfIndex,omitempty"`        // Interface If Index
	InterfaceType           string   `json:"interfaceType,omitempty"`           // Interface Type
	IPv4Address             string   `json:"ipv4Address,omitempty"`             // Ipv4 Address
	IPv6AddressList         []string `json:"ipv6AddressList,omitempty"`         // Ipv6 Address List
	IsL3Interface           *bool    `json:"isL3Interface,omitempty"`           // Is L3 Interface
	IsWan                   *bool    `json:"isWan,omitempty"`                   // Is Wan
	MacAddr                 string   `json:"macAddr,omitempty"`                 // Mac Addr
	MediaType               string   `json:"mediaType,omitempty"`               // Media Type
	Name                    string   `json:"name,omitempty"`                    // Name
	OperStatus              string   `json:"operStatus,omitempty"`              // Oper Status
	PeerStackMember         *int     `json:"peerStackMember,omitempty"`         // Peer Stack Member
	PeerStackPort           string   `json:"peerStackPort,omitempty"`           // Peer Stack Port
	PortChannelID           string   `json:"portChannelId,omitempty"`           // Port Channel Id
	PortMode                string   `json:"portMode,omitempty"`                // Port Mode
	PortType                string   `json:"portType,omitempty"`                // Port Type
	RxDiscards              *float64 `json:"rxDiscards,omitempty"`              // Rx Discards
	RxError                 *int     `json:"rxError,omitempty"`                 // Rx Error
	RxRate                  *float64 `json:"rxRate,omitempty"`                  // Rx Rate
	RxUtilization           *float64 `json:"rxUtilization,omitempty"`           // Rx Utilization
	Speed                   string   `json:"speed,omitempty"`                   // Speed
	StackPortType           string   `json:"stackPortType,omitempty"`           // Stack Port Type
	Timestamp               *int     `json:"timestamp,omitempty"`               // Timestamp
	TxDiscards              *float64 `json:"txDiscards,omitempty"`              // Tx Discards
	TxError                 *int     `json:"txError,omitempty"`                 // Tx Error
	TxRate                  *float64 `json:"txRate,omitempty"`                  // Tx Rate
	TxUtilization           *float64 `json:"txUtilization,omitempty"`           // Tx Utilization
	VLANID                  string   `json:"vlanId,omitempty"`                  // Vlan Id
	NetworkDeviceID         string   `json:"networkDeviceId,omitempty"`         // Network Device Id
	NetworkDeviceIPAddress  string   `json:"networkDeviceIpAddress,omitempty"`  // Network Device Ip Address
	NetworkDeviceMacAddress string   `json:"networkDeviceMacAddress,omitempty"` // Network Device Mac Address
	SiteHierarchy           string   `json:"siteHierarchy,omitempty"`           // Site Hierarchy
	SiteHierarchyID         string   `json:"siteHierarchyId,omitempty"`         // Site Hierarchy Id
	PoeAdminStatus          string   `json:"poeAdminStatus,omitempty"`          // Poe Admin Status
	PoeOperStatus           string   `json:"poeOperStatus,omitempty"`           // Poe Oper Status
	ChassisID               *int     `json:"chassisId,omitempty"`               // Chassis Id
	ModuleID                *int     `json:"moduleId,omitempty"`                // Module Id
	PdClassSignal           string   `json:"pdClassSignal,omitempty"`           // Pd Class Signal
	PdClassSpare            string   `json:"pdClassSpare,omitempty"`            // Pd Class Spare
	PdDeviceType            string   `json:"pdDeviceType,omitempty"`            // Pd Device Type
	PdDeviceModel           string   `json:"pdDeviceModel,omitempty"`           // Pd Device Model
	PdPowerAdminMaxInWatt   string   `json:"pdPowerAdminMaxInWatt,omitempty"`   // Pd Power Admin Max In Watt
	PdPowerBudgetInWatt     string   `json:"pdPowerBudgetInWatt,omitempty"`     // Pd Power Budget In Watt
	PdPowerConsumedInWatt   string   `json:"pdPowerConsumedInWatt,omitempty"`   // Pd Power Consumed In Watt
	PdPowerRemainingInWatt  string   `json:"pdPowerRemainingInWatt,omitempty"`  // Pd Power Remaining In Watt
	PdMaxPowerDrawn         string   `json:"pdMaxPowerDrawn,omitempty"`         // Pd Max Power Drawn
	PdConnectedDeviceList   []string `json:"pdConnectedDeviceList,omitempty"`   // Pd Connected Device List
	PoeOperPriority         string   `json:"poeOperPriority,omitempty"`         // Poe Oper Priority
	FastPoEEnabled          *bool    `json:"fastPoEEnabled,omitempty"`          // Fast Po E Enabled
	PerpetualPoEEnabled     *bool    `json:"perpetualPoEEnabled,omitempty"`     // Perpetual Po E Enabled
	PolicingPoEEnabled      *bool    `json:"policingPoEEnabled,omitempty"`      // Policing Po E Enabled
	UpoePlusEnabled         *bool    `json:"upoePlusEnabled,omitempty"`         // Upoe Plus Enabled
	FourPairEnabled         *bool    `json:"fourPairEnabled,omitempty"`         // Four Pair Enabled
	PoeDataTimestamp        *int     `json:"poeDataTimestamp,omitempty"`        // Poe Data Timestamp
	PdLocation              string   `json:"pdLocation,omitempty"`              // Pd Location
	PdDeviceName            string   `json:"pdDeviceName,omitempty"`            // Pd Device Name
	PdConnectedSwitch       string   `json:"pdConnectedSwitch,omitempty"`       // Pd Connected Switch
	ConnectedSwitchUUID     string   `json:"connectedSwitchUuid,omitempty"`     // Connected Switch Uuid
	IeeeCompliant           *bool    `json:"ieeeCompliant,omitempty"`           // Ieee Compliant
	ConnectedSwitchType     string   `json:"connectedSwitchType,omitempty"`     // Connected Switch Type
	SiteName                string   `json:"siteName,omitempty"`
}
type ResponseDevicesGetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesV1Page struct {
	Limit  *int                                                                                           `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                           `json:"offset,omitempty"` // Offset
	Count  *int                                                                                           `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesGetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesV1PageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesGetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesV1PageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseDevicesGetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountV1 struct {
	Response *ResponseDevicesGetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountV1Response `json:"response,omitempty"` //
	Version  string                                                                                                                                                             `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1 struct {
	Response *[]ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1Response `json:"response,omitempty"` //
	Page     *ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1Page       `json:"page,omitempty"`     //
	Version  string                                                                                                                            `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1Response struct {
	ID                      string                                                                                                                                               `json:"id,omitempty"`                      // Id
	AdminStatus             string                                                                                                                                               `json:"adminStatus,omitempty"`             // Admin Status
	Description             string                                                                                                                                               `json:"description,omitempty"`             // Description
	DuplexConfig            string                                                                                                                                               `json:"duplexConfig,omitempty"`            // Duplex Config
	DuplexOper              string                                                                                                                                               `json:"duplexOper,omitempty"`              // Duplex Oper
	InterfaceIfIndex        *int                                                                                                                                                 `json:"interfaceIfIndex,omitempty"`        // Interface If Index
	InterfaceType           string                                                                                                                                               `json:"interfaceType,omitempty"`           // Interface Type
	IPv4Address             string                                                                                                                                               `json:"ipv4Address,omitempty"`             // Ipv4 Address
	IPv6AddressList         []string                                                                                                                                             `json:"ipv6AddressList,omitempty"`         // Ipv6 Address List
	IsL3Interface           *bool                                                                                                                                                `json:"isL3Interface,omitempty"`           // Is L3 Interface
	IsWan                   *bool                                                                                                                                                `json:"isWan,omitempty"`                   // Is Wan
	MacAddr                 string                                                                                                                                               `json:"macAddr,omitempty"`                 // Mac Addr
	MediaType               string                                                                                                                                               `json:"mediaType,omitempty"`               // Media Type
	Name                    string                                                                                                                                               `json:"name,omitempty"`                    // Name
	OperStatus              string                                                                                                                                               `json:"operStatus,omitempty"`              // Oper Status
	PeerStackMember         *int                                                                                                                                                 `json:"peerStackMember,omitempty"`         // Peer Stack Member
	PeerStackPort           string                                                                                                                                               `json:"peerStackPort,omitempty"`           // Peer Stack Port
	PortChannelID           string                                                                                                                                               `json:"portChannelId,omitempty"`           // Port Channel Id
	PortMode                string                                                                                                                                               `json:"portMode,omitempty"`                // Port Mode
	PortType                string                                                                                                                                               `json:"portType,omitempty"`                // Port Type
	RxDiscards              *float64                                                                                                                                             `json:"rxDiscards,omitempty"`              // Rx Discards
	RxError                 *int                                                                                                                                                 `json:"rxError,omitempty"`                 // Rx Error
	RxRate                  *float64                                                                                                                                             `json:"rxRate,omitempty"`                  // Rx Rate
	RxUtilization           *float64                                                                                                                                             `json:"rxUtilization,omitempty"`           // Rx Utilization
	Speed                   string                                                                                                                                               `json:"speed,omitempty"`                   // Speed
	StackPortType           string                                                                                                                                               `json:"stackPortType,omitempty"`           // Stack Port Type
	Timestamp               *int                                                                                                                                                 `json:"timestamp,omitempty"`               // Timestamp
	TxDiscards              *float64                                                                                                                                             `json:"txDiscards,omitempty"`              // Tx Discards
	TxError                 *int                                                                                                                                                 `json:"txError,omitempty"`                 // Tx Error
	TxRate                  *float64                                                                                                                                             `json:"txRate,omitempty"`                  // Tx Rate
	TxUtilization           *float64                                                                                                                                             `json:"txUtilization,omitempty"`           // Tx Utilization
	VLANID                  string                                                                                                                                               `json:"vlanId,omitempty"`                  // Vlan Id
	NetworkDeviceID         string                                                                                                                                               `json:"networkDeviceId,omitempty"`         // Network Device Id
	NetworkDeviceIPAddress  string                                                                                                                                               `json:"networkDeviceIpAddress,omitempty"`  // Network Device Ip Address
	NetworkDeviceMacAddress string                                                                                                                                               `json:"networkDeviceMacAddress,omitempty"` // Network Device Mac Address
	SiteName                string                                                                                                                                               `json:"siteName,omitempty"`                // Site Name
	SiteHierarchy           string                                                                                                                                               `json:"siteHierarchy,omitempty"`           // Site Hierarchy
	SiteHierarchyID         string                                                                                                                                               `json:"siteHierarchyId,omitempty"`         // Site Hierarchy Id
	AggregateAttributes     *[]ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1ResponseAggregateAttributes `json:"aggregateAttributes,omitempty"`     //
}
type ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1ResponseAggregateAttributes struct {
	Name   string                                                                                                                                                     `json:"name,omitempty"`   // Name
	Values *[]ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1ResponseAggregateAttributesValues `json:"values,omitempty"` //
}
type ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1ResponseAggregateAttributesValues struct {
	Key   string   `json:"key,omitempty"`   // Key
	Value *float64 `json:"value,omitempty"` // Value
}
type ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1Page struct {
	Limit  *int                                                                                                                                `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                                                                `json:"offset,omitempty"` // Offset
	Count  *int                                                                                                                                `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1PageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1PageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesV1 struct {
	Response *ResponseDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesV1Response `json:"response,omitempty"` //
	Version  string                                                                   `json:"version,omitempty"`  // Version
}
type ResponseDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseDevicesGetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsAndPoeDataV1 struct {
	Response *ResponseDevicesGetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsAndPoeDataV1Response `json:"response,omitempty"` //
	Version  string                                                                                                          `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsAndPoeDataV1Response struct {
	ID                      string   `json:"id,omitempty"`                      // Id
	AdminStatus             string   `json:"adminStatus,omitempty"`             // Admin Status
	Description             string   `json:"description,omitempty"`             // Description
	DuplexConfig            string   `json:"duplexConfig,omitempty"`            // Duplex Config
	DuplexOper              string   `json:"duplexOper,omitempty"`              // Duplex Oper
	InterfaceIfIndex        *int     `json:"interfaceIfIndex,omitempty"`        // Interface If Index
	InterfaceType           string   `json:"interfaceType,omitempty"`           // Interface Type
	IPv4Address             string   `json:"ipv4Address,omitempty"`             // Ipv4 Address
	IPv6AddressList         []string `json:"ipv6AddressList,omitempty"`         // Ipv6 Address List
	IsL3Interface           *bool    `json:"isL3Interface,omitempty"`           // Is L3 Interface
	IsWan                   *bool    `json:"isWan,omitempty"`                   // Is Wan
	MacAddr                 string   `json:"macAddr,omitempty"`                 // Mac Addr
	MediaType               string   `json:"mediaType,omitempty"`               // Media Type
	Name                    string   `json:"name,omitempty"`                    // Name
	OperStatus              string   `json:"operStatus,omitempty"`              // Oper Status
	PeerStackMember         *int     `json:"peerStackMember,omitempty"`         // Peer Stack Member
	PeerStackPort           string   `json:"peerStackPort,omitempty"`           // Peer Stack Port
	PortChannelID           string   `json:"portChannelId,omitempty"`           // Port Channel Id
	PortMode                string   `json:"portMode,omitempty"`                // Port Mode
	PortType                string   `json:"portType,omitempty"`                // Port Type
	RxDiscards              *float64 `json:"rxDiscards,omitempty"`              // Rx Discards
	RxError                 *int     `json:"rxError,omitempty"`                 // Rx Error
	RxRate                  *float64 `json:"rxRate,omitempty"`                  // Rx Rate
	RxUtilization           *float64 `json:"rxUtilization,omitempty"`           // Rx Utilization
	Speed                   string   `json:"speed,omitempty"`                   // Speed
	StackPortType           string   `json:"stackPortType,omitempty"`           // Stack Port Type
	Timestamp               *int     `json:"timestamp,omitempty"`               // Timestamp
	TxDiscards              *float64 `json:"txDiscards,omitempty"`              // Tx Discards
	TxError                 *int     `json:"txError,omitempty"`                 // Tx Error
	TxRate                  *float64 `json:"txRate,omitempty"`                  // Tx Rate
	TxUtilization           *float64 `json:"txUtilization,omitempty"`           // Tx Utilization
	VLANID                  string   `json:"vlanId,omitempty"`                  // Vlan Id
	NetworkDeviceID         string   `json:"networkDeviceId,omitempty"`         // Network Device Id
	NetworkDeviceIPAddress  string   `json:"networkDeviceIpAddress,omitempty"`  // Network Device Ip Address
	NetworkDeviceMacAddress string   `json:"networkDeviceMacAddress,omitempty"` // Network Device Mac Address
	SiteHierarchy           string   `json:"siteHierarchy,omitempty"`           // Site Hierarchy
	SiteHierarchyID         string   `json:"siteHierarchyId,omitempty"`         // Site Hierarchy Id
	PoeAdminStatus          string   `json:"poeAdminStatus,omitempty"`          // Poe Admin Status
	PoeOperStatus           string   `json:"poeOperStatus,omitempty"`           // Poe Oper Status
	ChassisID               *int     `json:"chassisId,omitempty"`               // Chassis Id
	ModuleID                *int     `json:"moduleId,omitempty"`                // Module Id
	PdClassSignal           string   `json:"pdClassSignal,omitempty"`           // Pd Class Signal
	PdClassSpare            string   `json:"pdClassSpare,omitempty"`            // Pd Class Spare
	PdDeviceType            string   `json:"pdDeviceType,omitempty"`            // Pd Device Type
	PdDeviceModel           string   `json:"pdDeviceModel,omitempty"`           // Pd Device Model
	PdPowerAdminMaxInWatt   string   `json:"pdPowerAdminMaxInWatt,omitempty"`   // Pd Power Admin Max In Watt
	PdPowerBudgetInWatt     string   `json:"pdPowerBudgetInWatt,omitempty"`     // Pd Power Budget In Watt
	PdPowerConsumedInWatt   string   `json:"pdPowerConsumedInWatt,omitempty"`   // Pd Power Consumed In Watt
	PdPowerRemainingInWatt  string   `json:"pdPowerRemainingInWatt,omitempty"`  // Pd Power Remaining In Watt
	PdMaxPowerDrawn         string   `json:"pdMaxPowerDrawn,omitempty"`         // Pd Max Power Drawn
	PdConnectedDeviceList   []string `json:"pdConnectedDeviceList,omitempty"`   // Pd Connected Device List
	PoeOperPriority         string   `json:"poeOperPriority,omitempty"`         // Poe Oper Priority
	FastPoEEnabled          *bool    `json:"fastPoEEnabled,omitempty"`          // Fast Po E Enabled
	PerpetualPoEEnabled     *bool    `json:"perpetualPoEEnabled,omitempty"`     // Perpetual Po E Enabled
	PolicingPoEEnabled      *bool    `json:"policingPoEEnabled,omitempty"`      // Policing Po E Enabled
	UpoePlusEnabled         *bool    `json:"upoePlusEnabled,omitempty"`         // Upoe Plus Enabled
	FourPairEnabled         *bool    `json:"fourPairEnabled,omitempty"`         // Four Pair Enabled
	PoeDataTimestamp        *int     `json:"poeDataTimestamp,omitempty"`        // Poe Data Timestamp
	PdLocation              string   `json:"pdLocation,omitempty"`              // Pd Location
	PdDeviceName            string   `json:"pdDeviceName,omitempty"`            // Pd Device Name
	PdConnectedSwitch       string   `json:"pdConnectedSwitch,omitempty"`       // Pd Connected Switch
	ConnectedSwitchUUID     string   `json:"connectedSwitchUuid,omitempty"`     // Connected Switch Uuid
	IeeeCompliant           *bool    `json:"ieeeCompliant,omitempty"`           // Ieee Compliant
	ConnectedSwitchType     string   `json:"connectedSwitchType,omitempty"`     // Connected Switch Type
	SiteName                string   `json:"siteName,omitempty"`
}
type ResponseDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeV1 struct {
	Response       *[]ResponseDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeV1Response `json:"response,omitempty"`       //
	TimestampOrder string                                                                                   `json:"timestampOrder,omitempty"` // Timestamp Order
}
type ResponseDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeV1Response struct {
	Timestamp           *int                                                                                                        `json:"timestamp,omitempty"`           // Timestamp
	Attributes          *[]ResponseDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeV1ResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeV1ResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeV1ResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeV1ResponseAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name
}
type ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1 struct {
	Response *[]ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1Response `json:"response,omitempty"` //
	Page     *ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1Page       `json:"page,omitempty"`     //
	Version  string                                                                                   `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1Response struct {
	ID                         string                                                                                                      `json:"id,omitempty"`                         // Id
	Name                       string                                                                                                      `json:"name,omitempty"`                       // Name
	ManagementIPAddress        string                                                                                                      `json:"managementIpAddress,omitempty"`        // Management Ip Address
	PlatformID                 string                                                                                                      `json:"platformId,omitempty"`                 // Platform Id
	DeviceFamily               string                                                                                                      `json:"deviceFamily,omitempty"`               // Device Family
	SerialNumber               string                                                                                                      `json:"serialNumber,omitempty"`               // Serial Number
	MacAddress                 string                                                                                                      `json:"macAddress,omitempty"`                 // Mac Address
	DeviceSeries               string                                                                                                      `json:"deviceSeries,omitempty"`               // Device Series
	SoftwareVersion            string                                                                                                      `json:"softwareVersion,omitempty"`            // Software Version
	ProductVendor              string                                                                                                      `json:"productVendor,omitempty"`              // Product Vendor
	DeviceRole                 string                                                                                                      `json:"deviceRole,omitempty"`                 // Device Role
	DeviceType                 string                                                                                                      `json:"deviceType,omitempty"`                 // Device Type
	CommunicationState         string                                                                                                      `json:"communicationState,omitempty"`         // Communication State
	CollectionStatus           string                                                                                                      `json:"collectionStatus,omitempty"`           // Collection Status
	HaStatus                   string                                                                                                      `json:"haStatus,omitempty"`                   // Ha Status
	LastBootTime               *int                                                                                                        `json:"lastBootTime,omitempty"`               // Last Boot Time
	SiteHierarchyID            string                                                                                                      `json:"siteHierarchyId,omitempty"`            // Site Hierarchy Id
	SiteHierarchy              string                                                                                                      `json:"siteHierarchy,omitempty"`              // Site Hierarchy
	SiteID                     string                                                                                                      `json:"siteId,omitempty"`                     // Site Id
	DeviceGroupHierarchyID     string                                                                                                      `json:"deviceGroupHierarchyId,omitempty"`     // Device Group Hierarchy Id
	TagNames                   []string                                                                                                    `json:"tagNames,omitempty"`                   // Tag Names
	StackType                  string                                                                                                      `json:"stackType,omitempty"`                  // Stack Type
	OsType                     string                                                                                                      `json:"osType,omitempty"`                     // Os Type
	RingStatus                 *bool                                                                                                       `json:"ringStatus,omitempty"`                 // Ring Status
	MaintenanceModeEnabled     *bool                                                                                                       `json:"maintenanceModeEnabled,omitempty"`     // Maintenance Mode Enabled
	UpTime                     *int                                                                                                        `json:"upTime,omitempty"`                     // Up Time
	IPv4Address                string                                                                                                      `json:"ipv4Address,omitempty"`                // Ipv4 Address
	IPv6Address                string                                                                                                      `json:"ipv6Address,omitempty"`                // Ipv6 Address
	RedundancyMode             string                                                                                                      `json:"redundancyMode,omitempty"`             // Redundancy Mode
	FeatureFlagList            []string                                                                                                    `json:"featureFlagList,omitempty"`            // Feature Flag List
	HaLastResetReason          string                                                                                                      `json:"haLastResetReason,omitempty"`          // Ha Last Reset Reason
	RedundancyPeerStateDerived string                                                                                                      `json:"redundancyPeerStateDerived,omitempty"` // Redundancy Peer State Derived
	RedundancyPeerState        string                                                                                                      `json:"redundancyPeerState,omitempty"`        // Redundancy Peer State
	RedundancyStateDerived     string                                                                                                      `json:"redundancyStateDerived,omitempty"`     // Redundancy State Derived
	RedundancyState            string                                                                                                      `json:"redundancyState,omitempty"`            // Redundancy State
	WiredClientCount           *int                                                                                                        `json:"wiredClientCount,omitempty"`           // Wired Client Count
	WirelessClientCount        *int                                                                                                        `json:"wirelessClientCount,omitempty"`        // Wireless Client Count
	PortCount                  *int                                                                                                        `json:"portCount,omitempty"`                  // Port Count
	PhysicalPortCount          *int                                                                                                        `json:"physicalPortCount,omitempty"`          // Physical Port Count
	VirtualPortCount           *int                                                                                                        `json:"virtualPortCount,omitempty"`           // Virtual Port Count
	ClientCount                *int                                                                                                        `json:"clientCount,omitempty"`                // Client Count
	ApDetails                  *ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1ResponseApDetails             `json:"apDetails,omitempty"`                  //
	MetricsDetails             *ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1ResponseMetricsDetails        `json:"metricsDetails,omitempty"`             //
	FabricDetails              *ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1ResponseFabricDetails         `json:"fabricDetails,omitempty"`              //
	SwitchPoeDetails           *ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1ResponseSwitchPoeDetails      `json:"switchPoeDetails,omitempty"`           //
	FabricMetricsDetails       *ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1ResponseFabricMetricsDetails  `json:"fabricMetricsDetails,omitempty"`       //
	AggregateAttributes        *[]ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1ResponseAggregateAttributes `json:"aggregateAttributes,omitempty"`        //
}
type ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1ResponseApDetails struct {
	ConnectedWlcName     string                                                                                                  `json:"connectedWlcName,omitempty"`     // Connected Wlc Name
	PolicyTagName        string                                                                                                  `json:"policyTagName,omitempty"`        // Policy Tag Name
	ApOperationalState   string                                                                                                  `json:"apOperationalState,omitempty"`   // Ap Operational State
	PowerSaveMode        string                                                                                                  `json:"powerSaveMode,omitempty"`        // Power Save Mode
	OperationalMode      string                                                                                                  `json:"operationalMode,omitempty"`      // Operational Mode
	ResetReason          string                                                                                                  `json:"resetReason,omitempty"`          // Reset Reason
	Protocol             string                                                                                                  `json:"protocol,omitempty"`             // Protocol
	PowerMode            string                                                                                                  `json:"powerMode,omitempty"`            // Power Mode
	ConnectedTime        *int                                                                                                    `json:"connectedTime,omitempty"`        // Connected Time
	LedFlashEnabled      *bool                                                                                                   `json:"ledFlashEnabled,omitempty"`      // Led Flash Enabled
	LedFlashSeconds      *int                                                                                                    `json:"ledFlashSeconds,omitempty"`      // Led Flash Seconds
	SubMode              string                                                                                                  `json:"subMode,omitempty"`              // Sub Mode
	HomeApEnabled        *bool                                                                                                   `json:"homeApEnabled,omitempty"`        // Home Ap Enabled
	PowerType            string                                                                                                  `json:"powerType,omitempty"`            // Power Type
	ApType               string                                                                                                  `json:"apType,omitempty"`               // Ap Type
	AdminState           string                                                                                                  `json:"adminState,omitempty"`           // Admin State
	IcapCapability       string                                                                                                  `json:"icapCapability,omitempty"`       // Icap Capability
	RegulatoryDomain     string                                                                                                  `json:"regulatoryDomain,omitempty"`     // Regulatory Domain
	EthernetMac          string                                                                                                  `json:"ethernetMac,omitempty"`          // Ethernet Mac
	RfTagName            string                                                                                                  `json:"rfTagName,omitempty"`            // Rf Tag Name
	SiteTagName          string                                                                                                  `json:"siteTagName,omitempty"`          // Site Tag Name
	PowerSaveModeCapable string                                                                                                  `json:"powerSaveModeCapable,omitempty"` // Power Save Mode Capable
	PowerProfile         string                                                                                                  `json:"powerProfile,omitempty"`         // Power Profile
	FlexGroup            string                                                                                                  `json:"flexGroup,omitempty"`            // Flex Group
	PowerCalendarProfile string                                                                                                  `json:"powerCalendarProfile,omitempty"` // Power Calendar Profile
	ApGroup              string                                                                                                  `json:"apGroup,omitempty"`              // Ap Group
	Radios               *[]ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1ResponseApDetailsRadios `json:"radios,omitempty"`               //
}
type ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1ResponseApDetailsRadios struct {
	ID           string   `json:"id,omitempty"`           // Id
	Band         string   `json:"band,omitempty"`         // Band
	Noise        *int     `json:"noise,omitempty"`        // Noise
	AirQuality   *float64 `json:"airQuality,omitempty"`   // Air Quality
	Interference *float64 `json:"interference,omitempty"` // Interference
	TrafficUtil  *int     `json:"trafficUtil,omitempty"`  // Traffic Util
	Utilization  *float64 `json:"utilization,omitempty"`  // Utilization
	ClientCount  *int     `json:"clientCount,omitempty"`  // Client Count
}
type ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1ResponseMetricsDetails struct {
	OverallHealthScore                 *int     `json:"overallHealthScore,omitempty"`                 // Overall Health Score
	CPUUtilization                     *float64 `json:"cpuUtilization,omitempty"`                     // Cpu Utilization
	CPUScore                           *int     `json:"cpuScore,omitempty"`                           // Cpu Score
	MemoryUtilization                  *float64 `json:"memoryUtilization,omitempty"`                  // Memory Utilization
	MemoryScore                        *int     `json:"memoryScore,omitempty"`                        // Memory Score
	AvgTemperature                     *float64 `json:"avgTemperature,omitempty"`                     // Avg Temperature
	MaxTemperature                     *float64 `json:"maxTemperature,omitempty"`                     // Max Temperature
	DiscardScore                       *int     `json:"discardScore,omitempty"`                       // Discard Score
	DiscardInterfaces                  []string `json:"discardInterfaces,omitempty"`                  // Discard Interfaces
	ErrorScore                         *int     `json:"errorScore,omitempty"`                         // Error Score
	ErrorInterfaces                    []string `json:"errorInterfaces,omitempty"`                    // Error Interfaces
	InterDeviceLinkScore               *int     `json:"interDeviceLinkScore,omitempty"`               // Inter Device Link Score
	InterDeviceConnectedDownInterfaces []string `json:"interDeviceConnectedDownInterfaces,omitempty"` // Inter Device Connected Down Interfaces
	LinkUtilizationScore               *int     `json:"linkUtilizationScore,omitempty"`               // Link Utilization Score
	HighLinkUtilizationInterfaces      []string `json:"highLinkUtilizationInterfaces,omitempty"`      // High Link Utilization Interfaces
	FreeTimerScore                     *int     `json:"freeTimerScore,omitempty"`                     // Free Timer Score
	FreeTimer                          *float64 `json:"freeTimer,omitempty"`                          // Free Timer
	PacketPoolScore                    *int     `json:"packetPoolScore,omitempty"`                    // Packet Pool Score
	PacketPool                         *int     `json:"packetPool,omitempty"`                         // Packet Pool
	FreeMemoryBufferScore              *int     `json:"freeMemoryBufferScore,omitempty"`              // Free Memory Buffer Score
	FreeMemoryBuffer                   *float64 `json:"freeMemoryBuffer,omitempty"`                   // Free Memory Buffer
	WqePoolScore                       *int     `json:"wqePoolScore,omitempty"`                       // Wqe Pool Score
	WqePool                            *int     `json:"wqePool,omitempty"`                            // Wqe Pool
	ApCount                            *int     `json:"apCount,omitempty"`                            // Ap Count
	NoiseScore                         *int     `json:"noiseScore,omitempty"`                         // Noise Score
	UtilizationScore                   *int     `json:"utilizationScore,omitempty"`                   // Utilization Score
	InterferenceScore                  *int     `json:"interferenceScore,omitempty"`                  // Interference Score
	AirQualityScore                    *int     `json:"airQualityScore,omitempty"`                    // Air Quality Score
}
type ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1ResponseFabricDetails struct {
	FabricRole      []string `json:"fabricRole,omitempty"`      // Fabric Role
	FabricSiteName  string   `json:"fabricSiteName,omitempty"`  // Fabric Site Name
	TransitFabrics  []string `json:"transitFabrics,omitempty"`  // Transit Fabrics
	L2Vns           []string `json:"l2Vns,omitempty"`           // L2 Vns
	L3Vns           []string `json:"l3Vns,omitempty"`           // L3 Vns
	FabricSiteID    string   `json:"fabricSiteId,omitempty"`    // Fabric Site Id
	NetworkProtocol string   `json:"networkProtocol,omitempty"` // Network Protocol
}
type ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1ResponseSwitchPoeDetails struct {
	PortCount            *int                                                                                                                  `json:"portCount,omitempty"`            // Port Count
	UsedPortCount        *int                                                                                                                  `json:"usedPortCount,omitempty"`        // Used Port Count
	FreePortCount        *int                                                                                                                  `json:"freePortCount,omitempty"`        // Free Port Count
	PowerConsumed        *float64                                                                                                              `json:"powerConsumed,omitempty"`        // Power Consumed
	PoePowerConsumed     *int                                                                                                                  `json:"poePowerConsumed,omitempty"`     // Poe Power Consumed
	SystemPowerConsumed  *float64                                                                                                              `json:"systemPowerConsumed,omitempty"`  // System Power Consumed
	PowerBudget          *int                                                                                                                  `json:"powerBudget,omitempty"`          // Power Budget
	PoePowerAllocated    *float64                                                                                                              `json:"poePowerAllocated,omitempty"`    // Poe Power Allocated
	SystemPowerAllocated *int                                                                                                                  `json:"systemPowerAllocated,omitempty"` // System Power Allocated
	PowerRemaining       *float64                                                                                                              `json:"powerRemaining,omitempty"`       // Power Remaining
	PoeVersion           string                                                                                                                `json:"poeVersion,omitempty"`           // Poe Version
	ChassisCount         *int                                                                                                                  `json:"chassisCount,omitempty"`         // Chassis Count
	ModuleCount          *int                                                                                                                  `json:"moduleCount,omitempty"`          // Module Count
	ModuleDetails        *[]ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1ResponseSwitchPoeDetailsModuleDetails `json:"moduleDetails,omitempty"`        //
}
type ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1ResponseSwitchPoeDetailsModuleDetails struct {
	ModuleID                   string   `json:"moduleId,omitempty"`                   // Module Id
	ChassisID                  string   `json:"chassisId,omitempty"`                  // Chassis Id
	ModulePortCount            *int     `json:"modulePortCount,omitempty"`            // Module Port Count
	ModuleUsedPortCount        *int     `json:"moduleUsedPortCount,omitempty"`        // Module Used Port Count
	ModuleFreePortCount        *int     `json:"moduleFreePortCount,omitempty"`        // Module Free Port Count
	ModulePowerConsumed        *float64 `json:"modulePowerConsumed,omitempty"`        // Module Power Consumed
	ModulePoePowerConsumed     *int     `json:"modulePoePowerConsumed,omitempty"`     // Module Poe Power Consumed
	ModuleSystemPowerConsumed  *float64 `json:"moduleSystemPowerConsumed,omitempty"`  // Module System Power Consumed
	ModulePowerBudget          *int     `json:"modulePowerBudget,omitempty"`          // Module Power Budget
	ModulePoePowerAllocated    *float64 `json:"modulePoePowerAllocated,omitempty"`    // Module Poe Power Allocated
	ModuleSystemPowerAllocated *int     `json:"moduleSystemPowerAllocated,omitempty"` // Module System Power Allocated
	ModulePowerRemaining       *float64 `json:"modulePowerRemaining,omitempty"`       // Module Power Remaining
	InterfacePowerMax          *int     `json:"interfacePowerMax,omitempty"`          // Interface Power Max
}
type ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1ResponseFabricMetricsDetails struct {
	OverallFabricScore       *int `json:"overallFabricScore,omitempty"`       // Overall Fabric Score
	FabricTransitScore       *int `json:"fabricTransitScore,omitempty"`       // Fabric Transit Score
	FabricSiteScore          *int `json:"fabricSiteScore,omitempty"`          // Fabric Site Score
	FabricVnScore            *int `json:"fabricVnScore,omitempty"`            // Fabric Vn Score
	FabsiteFcpScore          *int `json:"fabsiteFcpScore,omitempty"`          // Fabsite Fcp Score
	FabsiteInfraScore        *int `json:"fabsiteInfraScore,omitempty"`        // Fabsite Infra Score
	FabsiteFsconnScore       *int `json:"fabsiteFsconnScore,omitempty"`       // Fabsite Fsconn Score
	VnExitScore              *int `json:"vnExitScore,omitempty"`              // Vn Exit Score
	VnFcpScore               *int `json:"vnFcpScore,omitempty"`               // Vn Fcp Score
	VnStatusScore            *int `json:"vnStatusScore,omitempty"`            // Vn Status Score
	VnServiceScore           *int `json:"vnServiceScore,omitempty"`           // Vn Service Score
	TransitControlPlaneScore *int `json:"transitControlPlaneScore,omitempty"` // Transit Control Plane Score
	TransitServicesScore     *int `json:"transitServicesScore,omitempty"`     // Transit Services Score
	TCPConnScore             *int `json:"tcpConnScore,omitempty"`             // Tcp Conn Score
	BgpBgpSiteScore          *int `json:"bgpBgpSiteScore,omitempty"`          // Bgp Bgp Site Score
	VniStatusScore           *int `json:"vniStatusScore,omitempty"`           // Vni Status Score
	PubsubTransitConnScore   *int `json:"pubsubTransitConnScore,omitempty"`   // Pubsub Transit Conn Score
	BgpPeerInfraVnScore      *int `json:"bgpPeerInfraVnScore,omitempty"`      // Bgp Peer Infra Vn Score
	InternetAvailScore       *int `json:"internetAvailScore,omitempty"`       // Internet Avail Score
	BgpEvpnScore             *int `json:"bgpEvpnScore,omitempty"`             // Bgp Evpn Score
	LispTransitConnScore     *int `json:"lispTransitConnScore,omitempty"`     // Lisp Transit Conn Score
	CtsEnvDataDownloadScore  *int `json:"ctsEnvDataDownloadScore,omitempty"`  // Cts Env Data Download Score
	PubsubInfraVnScore       *int `json:"pubsubInfraVnScore,omitempty"`       // Pubsub Infra Vn Score
	PeerScore                *int `json:"peerScore,omitempty"`                // Peer Score
	BgpPeerScore             *int `json:"bgpPeerScore,omitempty"`             // Bgp Peer Score
	RemoteInternetAvailScore *int `json:"remoteInternetAvailScore,omitempty"` // Remote Internet Avail Score
	BgpTCPScore              *int `json:"bgpTcpScore,omitempty"`              // Bgp Tcp Score
	PubsubSessionScore       *int `json:"pubsubSessionScore,omitempty"`       // Pubsub Session Score
	AAAStatusScore           *int `json:"aaaStatusScore,omitempty"`           // Aaa Status Score
	LispCpConnScore          *int `json:"lispCpConnScore,omitempty"`          // Lisp Cp Conn Score
	BgpPubsubSiteScore       *int `json:"bgpPubsubSiteScore,omitempty"`       // Bgp Pubsub Site Score
	McastScore               *int `json:"mcastScore,omitempty"`               // Mcast Score
	PortChannelScore         *int `json:"portChannelScore,omitempty"`         // Port Channel Score
}
type ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1ResponseAggregateAttributes struct {
	Name     string   `json:"name,omitempty"`     // Name
	Function string   `json:"function,omitempty"` // Function
	Value    *float64 `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1Page struct {
	Limit  *int   `json:"limit,omitempty"`  // Limit
	Offset *int   `json:"offset,omitempty"` // Offset
	Count  *int   `json:"count,omitempty"`  // Count
	SortBy string `json:"sortBy,omitempty"` // Sort By
	Order  string `json:"order,omitempty"`  // Order
}
type ResponseDevicesGetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersV1 struct {
	Response *ResponseDevicesGetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersV1Response `json:"response,omitempty"` //
	Version  string                                                                                     `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1 struct {
	Response *[]ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1Response `json:"response,omitempty"` //
	Page     *ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1Page       `json:"page,omitempty"`     //
	Version  string                                                                                                         `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1Response struct {
	ID string `json:"id,omitempty"` // Id

	Name string `json:"name,omitempty"` // Name

	ManagementIPAddress string `json:"managementIpAddress,omitempty"` // Management Ip Address

	PlatformID string `json:"platformId,omitempty"` // Platform Id

	DeviceFamily string `json:"deviceFamily,omitempty"` // Device Family

	SerialNumber string `json:"serialNumber,omitempty"` // Serial Number

	MacAddress string `json:"macAddress,omitempty"` // Mac Address

	DeviceSeries string `json:"deviceSeries,omitempty"` // Device Series

	SoftwareVersion string `json:"softwareVersion,omitempty"` // Software Version

	ProductVendor string `json:"productVendor,omitempty"` // Product Vendor

	DeviceRole string `json:"deviceRole,omitempty"` // Device Role

	DeviceType string `json:"deviceType,omitempty"` // Device Type

	CommunicationState string `json:"communicationState,omitempty"` // Communication State

	CollectionStatus string `json:"collectionStatus,omitempty"` // Collection Status

	HaStatus string `json:"haStatus,omitempty"` // Ha Status

	LastBootTime *int `json:"lastBootTime,omitempty"` // Last Boot Time

	SiteHierarchyID string `json:"siteHierarchyId,omitempty"` // Site Hierarchy Id

	SiteHierarchy string `json:"siteHierarchy,omitempty"` // Site Hierarchy

	SiteID string `json:"siteId,omitempty"` // Site Id

	DeviceGroupHierarchyID string `json:"deviceGroupHierarchyId,omitempty"` // Device Group Hierarchy Id

	TagNames []string `json:"tagNames,omitempty"` // Tag Names

	StackType string `json:"stackType,omitempty"` // Stack Type

	OsType string `json:"osType,omitempty"` // Os Type

	RingStatus *bool `json:"ringStatus,omitempty"` // Ring Status

	MaintenanceModeEnabled *bool `json:"maintenanceModeEnabled,omitempty"` // Maintenance Mode Enabled

	UpTime *int `json:"upTime,omitempty"` // Up Time

	IPv4Address string `json:"ipv4Address,omitempty"` // Ipv4 Address

	IPv6Address string `json:"ipv6Address,omitempty"` // Ipv6 Address

	RedundancyMode string `json:"redundancyMode,omitempty"` // Redundancy Mode

	FeatureFlagList []string `json:"featureFlagList,omitempty"` // Feature Flag List

	HaLastResetReason string `json:"haLastResetReason,omitempty"` // Ha Last Reset Reason

	RedundancyPeerStateDerived string `json:"redundancyPeerStateDerived,omitempty"` // Redundancy Peer State Derived

	RedundancyPeerState string `json:"redundancyPeerState,omitempty"` // Redundancy Peer State

	RedundancyStateDerived string `json:"redundancyStateDerived,omitempty"` // Redundancy State Derived

	RedundancyState string `json:"redundancyState,omitempty"` // Redundancy State

	WiredClientCount *int `json:"wiredClientCount,omitempty"` // Wired Client Count

	WirelessClientCount *int `json:"wirelessClientCount,omitempty"` // Wireless Client Count

	PortCount *int `json:"portCount,omitempty"` // Port Count

	PhysicalPortCount *int `json:"physicalPortCount,omitempty"` // Physical Port Count

	VirtualPortCount *int `json:"virtualPortCount,omitempty"` // Virtual Port Count

	ClientCount *int `json:"clientCount,omitempty"` // Client Count

	ApDetails *ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1ResponseApDetails `json:"apDetails,omitempty"` //

	MetricsDetails *ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1ResponseMetricsDetails `json:"metricsDetails,omitempty"` //

	FabricDetails *ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1ResponseFabricDetails `json:"fabricDetails,omitempty"` //

	SwitchPoeDetails *ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1ResponseSwitchPoeDetails `json:"switchPoeDetails,omitempty"` //

	FabricMetricsDetails *ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1ResponseFabricMetricsDetails `json:"fabricMetricsDetails,omitempty"` //

	AggregateAttributes *[]ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1ResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1ResponseApDetails struct {
	ConnectedWlcName     string                                                                                                                        `json:"connectedWlcName,omitempty"`     // Connected Wlc Name
	PolicyTagName        string                                                                                                                        `json:"policyTagName,omitempty"`        // Policy Tag Name
	ApOperationalState   string                                                                                                                        `json:"apOperationalState,omitempty"`   // Ap Operational State
	PowerSaveMode        string                                                                                                                        `json:"powerSaveMode,omitempty"`        // Power Save Mode
	OperationalMode      string                                                                                                                        `json:"operationalMode,omitempty"`      // Operational Mode
	ResetReason          string                                                                                                                        `json:"resetReason,omitempty"`          // Reset Reason
	Protocol             string                                                                                                                        `json:"protocol,omitempty"`             // Protocol
	PowerMode            string                                                                                                                        `json:"powerMode,omitempty"`            // Power Mode
	ConnectedTime        *int                                                                                                                          `json:"connectedTime,omitempty"`        // Connected Time
	LedFlashEnabled      *bool                                                                                                                         `json:"ledFlashEnabled,omitempty"`      // Led Flash Enabled
	LedFlashSeconds      *int                                                                                                                          `json:"ledFlashSeconds,omitempty"`      // Led Flash Seconds
	SubMode              string                                                                                                                        `json:"subMode,omitempty"`              // Sub Mode
	HomeApEnabled        *bool                                                                                                                         `json:"homeApEnabled,omitempty"`        // Home Ap Enabled
	PowerType            string                                                                                                                        `json:"powerType,omitempty"`            // Power Type
	ApType               string                                                                                                                        `json:"apType,omitempty"`               // Ap Type
	AdminState           string                                                                                                                        `json:"adminState,omitempty"`           // Admin State
	IcapCapability       string                                                                                                                        `json:"icapCapability,omitempty"`       // Icap Capability
	RegulatoryDomain     string                                                                                                                        `json:"regulatoryDomain,omitempty"`     // Regulatory Domain
	EthernetMac          string                                                                                                                        `json:"ethernetMac,omitempty"`          // Ethernet Mac
	RfTagName            string                                                                                                                        `json:"rfTagName,omitempty"`            // Rf Tag Name
	SiteTagName          string                                                                                                                        `json:"siteTagName,omitempty"`          // Site Tag Name
	PowerSaveModeCapable string                                                                                                                        `json:"powerSaveModeCapable,omitempty"` // Power Save Mode Capable
	PowerProfile         string                                                                                                                        `json:"powerProfile,omitempty"`         // Power Profile
	FlexGroup            string                                                                                                                        `json:"flexGroup,omitempty"`            // Flex Group
	PowerCalendarProfile string                                                                                                                        `json:"powerCalendarProfile,omitempty"` // Power Calendar Profile
	ApGroup              string                                                                                                                        `json:"apGroup,omitempty"`              // Ap Group
	Radios               *[]ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1ResponseApDetailsRadios `json:"radios,omitempty"`               //
}
type ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1ResponseApDetailsRadios struct {
	ID           string   `json:"id,omitempty"`           // Id
	Band         string   `json:"band,omitempty"`         // Band
	Noise        *int     `json:"noise,omitempty"`        // Noise
	AirQuality   *float64 `json:"airQuality,omitempty"`   // Air Quality
	Interference *float64 `json:"interference,omitempty"` // Interference
	TrafficUtil  *int     `json:"trafficUtil,omitempty"`  // Traffic Util
	Utilization  *float64 `json:"utilization,omitempty"`  // Utilization
	ClientCount  *int     `json:"clientCount,omitempty"`  // Client Count
}
type ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1ResponseMetricsDetails struct {
	OverallHealthScore                 *int     `json:"overallHealthScore,omitempty"` // Overall Health Score
	OverallFabricScore                 *int     `json:"overallFabricScore,omitempty"`
	CPUUtilization                     *float64 `json:"cpuUtilization,omitempty"`                     // Cpu Utilization
	CPUScore                           *int     `json:"cpuScore,omitempty"`                           // Cpu Score
	MemoryUtilization                  *float64 `json:"memoryUtilization,omitempty"`                  // Memory Utilization
	MemoryScore                        *int     `json:"memoryScore,omitempty"`                        // Memory Score
	AvgTemperature                     *float64 `json:"avgTemperature,omitempty"`                     // Avg Temperature
	MaxTemperature                     *float64 `json:"maxTemperature,omitempty"`                     // Max Temperature
	DiscardScore                       *int     `json:"discardScore,omitempty"`                       // Discard Score
	DiscardInterfaces                  []string `json:"discardInterfaces,omitempty"`                  // Discard Interfaces
	ErrorScore                         *int     `json:"errorScore,omitempty"`                         // Error Score
	ErrorInterfaces                    []string `json:"errorInterfaces,omitempty"`                    // Error Interfaces
	InterDeviceLinkScore               *int     `json:"interDeviceLinkScore,omitempty"`               // Inter Device Link Score
	InterDeviceConnectedDownInterfaces []string `json:"interDeviceConnectedDownInterfaces,omitempty"` // Inter Device Connected Down Interfaces
	LinkUtilizationScore               *int     `json:"linkUtilizationScore,omitempty"`               // Link Utilization Score
	HighLinkUtilizationInterfaces      []string `json:"highLinkUtilizationInterfaces,omitempty"`      // High Link Utilization Interfaces
	FreeTimerScore                     *int     `json:"freeTimerScore,omitempty"`                     // Free Timer Score
	FreeTimer                          *float64 `json:"freeTimer,omitempty"`                          // Free Timer
	PacketPoolScore                    *int     `json:"packetPoolScore,omitempty"`                    // Packet Pool Score
	PacketPool                         *int     `json:"packetPool,omitempty"`                         // Packet Pool
	FreeMemoryBufferScore              *int     `json:"freeMemoryBufferScore,omitempty"`              // Free Memory Buffer Score
	FreeMemoryBuffer                   *float64 `json:"freeMemoryBuffer,omitempty"`                   // Free Memory Buffer
	WqePoolScore                       *int     `json:"wqePoolScore,omitempty"`                       // Wqe Pool Score
	WqePool                            *int     `json:"wqePool,omitempty"`                            // Wqe Pool
	ApCount                            *int     `json:"apCount,omitempty"`                            // Ap Count
	NoiseScore                         *int     `json:"noiseScore,omitempty"`                         // Noise Score
	UtilizationScore                   *int     `json:"utilizationScore,omitempty"`                   // Utilization Score
	InterferenceScore                  *int     `json:"interferenceScore,omitempty"`                  // Interference Score
	AirQualityScore                    *int     `json:"airQualityScore,omitempty"`                    // Air Quality Score
}
type ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1ResponseFabricDetails struct {
	FabricRole []string `json:"fabricRole,omitempty"` // Fabric Role

	FabricSiteName string `json:"fabricSiteName,omitempty"` // Fabric Site Name

	TransitFabrics []string `json:"transitFabrics,omitempty"` // Transit Fabrics

	L2Vns []string `json:"l2Vns,omitempty"` // L2 Vns

	L3Vns []string `json:"l3Vns,omitempty"` // L3 Vns

	FabricSiteID string `json:"fabricSiteId,omitempty"` // Fabric Site Id

	NetworkProtocol string `json:"networkProtocol,omitempty"` // Network Protocol
}
type ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1ResponseSwitchPoeDetails struct {
	PortCount *int `json:"portCount,omitempty"` // Port Count

	UsedPortCount *int `json:"usedPortCount,omitempty"` // Used Port Count

	FreePortCount *int `json:"freePortCount,omitempty"` // Free Port Count

	PowerConsumed *float64 `json:"powerConsumed,omitempty"` // Power Consumed

	PoePowerConsumed *int `json:"poePowerConsumed,omitempty"` // Poe Power Consumed

	SystemPowerConsumed *float64 `json:"systemPowerConsumed,omitempty"` // System Power Consumed

	PowerBudget *int `json:"powerBudget,omitempty"` // Power Budget

	PoePowerAllocated *float64 `json:"poePowerAllocated,omitempty"` // Poe Power Allocated

	SystemPowerAllocated *int `json:"systemPowerAllocated,omitempty"` // System Power Allocated

	PowerRemaining *float64 `json:"powerRemaining,omitempty"` // Power Remaining

	PoeVersion string `json:"poeVersion,omitempty"` // Poe Version

	ChassisCount *int `json:"chassisCount,omitempty"` // Chassis Count

	ModuleCount *int `json:"moduleCount,omitempty"` // Module Count

	ModuleDetails *[]ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1ResponseSwitchPoeDetailsModuleDetails `json:"moduleDetails,omitempty"` //
}
type ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1ResponseSwitchPoeDetailsModuleDetails struct {
	ModuleID string `json:"moduleId,omitempty"` // Module Id

	ChassisID string `json:"chassisId,omitempty"` // Chassis Id

	ModulePortCount *int `json:"modulePortCount,omitempty"` // Module Port Count

	ModuleUsedPortCount *int `json:"moduleUsedPortCount,omitempty"` // Module Used Port Count

	ModuleFreePortCount *int `json:"moduleFreePortCount,omitempty"` // Module Free Port Count

	ModulePowerConsumed *float64 `json:"modulePowerConsumed,omitempty"` // Module Power Consumed

	ModulePoePowerConsumed *int `json:"modulePoePowerConsumed,omitempty"` // Module Poe Power Consumed

	ModuleSystemPowerConsumed *float64 `json:"moduleSystemPowerConsumed,omitempty"` // Module System Power Consumed

	ModulePowerBudget *int `json:"modulePowerBudget,omitempty"` // Module Power Budget

	ModulePoePowerAllocated *float64 `json:"modulePoePowerAllocated,omitempty"` // Module Poe Power Allocated

	ModuleSystemPowerAllocated *int `json:"moduleSystemPowerAllocated,omitempty"` // Module System Power Allocated

	ModulePowerRemaining *float64 `json:"modulePowerRemaining,omitempty"` // Module Power Remaining

	InterfacePowerMax *int `json:"interfacePowerMax,omitempty"` // Interface Power Max
}
type ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1ResponseFabricMetricsDetails struct {
	OverallFabricScore *int `json:"overallFabricScore,omitempty"` // Overall Fabric Score

	FabricTransitScore *int `json:"fabricTransitScore,omitempty"` // Fabric Transit Score

	FabricSiteScore *int `json:"fabricSiteScore,omitempty"` // Fabric Site Score

	FabricVnScore *int `json:"fabricVnScore,omitempty"` // Fabric Vn Score

	FabsiteFcpScore *int `json:"fabsiteFcpScore,omitempty"` // Fabsite Fcp Score

	FabsiteInfraScore *int `json:"fabsiteInfraScore,omitempty"` // Fabsite Infra Score

	FabsiteFsconnScore *int `json:"fabsiteFsconnScore,omitempty"` // Fabsite Fsconn Score

	VnExitScore *int `json:"vnExitScore,omitempty"` // Vn Exit Score

	VnFcpScore *int `json:"vnFcpScore,omitempty"` // Vn Fcp Score

	VnStatusScore *int `json:"vnStatusScore,omitempty"` // Vn Status Score

	VnServiceScore *int `json:"vnServiceScore,omitempty"` // Vn Service Score

	TransitControlPlaneScore *int `json:"transitControlPlaneScore,omitempty"` // Transit Control Plane Score

	TransitServicesScore *int `json:"transitServicesScore,omitempty"` // Transit Services Score

	TCPConnScore *int `json:"tcpConnScore,omitempty"` // Tcp Conn Score

	BgpBgpSiteScore *int `json:"bgpBgpSiteScore,omitempty"` // Bgp Bgp Site Score

	VniStatusScore *int `json:"vniStatusScore,omitempty"` // Vni Status Score

	PubsubTransitConnScore *int `json:"pubsubTransitConnScore,omitempty"` // Pubsub Transit Conn Score

	BgpPeerInfraVnScore *int `json:"bgpPeerInfraVnScore,omitempty"` // Bgp Peer Infra Vn Score

	InternetAvailScore *int `json:"internetAvailScore,omitempty"` // Internet Avail Score

	BgpEvpnScore *int `json:"bgpEvpnScore,omitempty"` // Bgp Evpn Score

	LispTransitConnScore *int `json:"lispTransitConnScore,omitempty"` // Lisp Transit Conn Score

	CtsEnvDataDownloadScore *int `json:"ctsEnvDataDownloadScore,omitempty"` // Cts Env Data Download Score

	PubsubInfraVnScore *int `json:"pubsubInfraVnScore,omitempty"` // Pubsub Infra Vn Score

	PeerScore *int `json:"peerScore,omitempty"` // Peer Score

	BgpPeerScore *int `json:"bgpPeerScore,omitempty"` // Bgp Peer Score

	RemoteInternetAvailScore *int `json:"remoteInternetAvailScore,omitempty"` // Remote Internet Avail Score

	BgpTCPScore *int `json:"bgpTcpScore,omitempty"` // Bgp Tcp Score

	PubsubSessionScore *int `json:"pubsubSessionScore,omitempty"` // Pubsub Session Score

	AAAStatusScore *int `json:"aaaStatusScore,omitempty"` // Aaa Status Score

	LispCpConnScore *int `json:"lispCpConnScore,omitempty"` // Lisp Cp Conn Score

	BgpPubsubSiteScore *int `json:"bgpPubsubSiteScore,omitempty"` // Bgp Pubsub Site Score

	McastScore *int `json:"mcastScore,omitempty"` // Mcast Score

	PortChannelScore *int `json:"portChannelScore,omitempty"` // Port Channel Score
}
type ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1ResponseAggregateAttributes struct {
	Name     string   `json:"name,omitempty"`     // Name
	Function string   `json:"function,omitempty"` // Function
	Value    *float64 `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	Count *int `json:"count,omitempty"` // Count

	SortBy string `json:"sortBy,omitempty"` // Sort By

	Order string `json:"order,omitempty"` // Order
}
type ResponseDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1 struct {
	Response *ResponseDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1 struct {
	Response *ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1Response `json:"response,omitempty"` //

	Page *[]ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1Page `json:"page,omitempty"` //
}
type ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1Response struct {
	Attributes *[]ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1ResponseAttributes `json:"attributes,omitempty"` // Attributes

	AggregateAttributes *[]ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1ResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` // Aggregate Attributes

	Groups *[]ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1ResponseGroups `json:"groups,omitempty"` //
}
type ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1ResponseAttributes interface{}
type ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1ResponseAggregateAttributes interface{}
type ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1ResponseGroups struct {
	ID string `json:"id,omitempty"` // Id

	Attributes *[]ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1ResponseGroupsAttributes `json:"attributes,omitempty"` //

	AggregateAttributes *[]ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1ResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1ResponseGroupsAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1ResponseGroupsAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function

	Value *float64 `json:"value,omitempty"` // Value
}
type ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	Count *int `json:"count,omitempty"` // Count

	SortBy *[]ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1PageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1PageSortBy struct {
	Name string `json:"name,omitempty"` // Name

	Order string `json:"order,omitempty"` // Order
}
type ResponseDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1 struct {
	Response *[]ResponseDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1Response `json:"response,omitempty"` //

	Page *[]ResponseDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1Page `json:"page,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1Response struct {
	ID string `json:"id,omitempty"` // Id

	Attributes *[]ResponseDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1ResponseAttributes `json:"attributes,omitempty"` //

	AggregateAttributes *[]ResponseDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1ResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1ResponseAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1ResponseAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function

	Value *float64 `json:"value,omitempty"` // Value
}
type ResponseDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	Count *int `json:"count,omitempty"` // Count

	SortBy *[]ResponseDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1PageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1PageSortBy struct {
	Name string `json:"name,omitempty"` // Name

	Order string `json:"order,omitempty"` // Order

	Function string `json:"function,omitempty"` // Function
}
type ResponseDevicesGetsTheTrendAnalyticsDataV1 struct {
	Response *[]ResponseDevicesGetsTheTrendAnalyticsDataV1Response `json:"response,omitempty"` //

	Page *ResponseDevicesGetsTheTrendAnalyticsDataV1Page `json:"page,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesGetsTheTrendAnalyticsDataV1Response struct {
	Timestamp *float64 `json:"timestamp,omitempty"` // Timestamp

	Attributes *[]ResponseDevicesGetsTheTrendAnalyticsDataV1ResponseAttributes `json:"attributes,omitempty"` //

	AggregateAttributes *[]ResponseDevicesGetsTheTrendAnalyticsDataV1ResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Groups *[]ResponseDevicesGetsTheTrendAnalyticsDataV1ResponseGroups `json:"groups,omitempty"` //
}
type ResponseDevicesGetsTheTrendAnalyticsDataV1ResponseAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetsTheTrendAnalyticsDataV1ResponseAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function

	Value *float64 `json:"value,omitempty"` // Value
}
type ResponseDevicesGetsTheTrendAnalyticsDataV1ResponseGroups struct {
	ID string `json:"id,omitempty"` // Id

	Attributes *[]ResponseDevicesGetsTheTrendAnalyticsDataV1ResponseGroupsAttributes `json:"attributes,omitempty"` //

	AggregateAttributes *[]ResponseDevicesGetsTheTrendAnalyticsDataV1ResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetsTheTrendAnalyticsDataV1ResponseGroupsAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetsTheTrendAnalyticsDataV1ResponseGroupsAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function

	Value *float64 `json:"value,omitempty"` // Value
}
type ResponseDevicesGetsTheTrendAnalyticsDataV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	Count *int `json:"count,omitempty"` // Count

	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDV1 struct {
	Response *ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDV1Response struct {
	ID string `json:"id,omitempty"` // Id

	Name string `json:"name,omitempty"` // Name

	ManagementIPAddress string `json:"managementIpAddress,omitempty"` // Management Ip Address

	PlatformID string `json:"platformId,omitempty"` // Platform Id

	DeviceFamily string `json:"deviceFamily,omitempty"` // Device Family

	SerialNumber string `json:"serialNumber,omitempty"` // Serial Number

	MacAddress string `json:"macAddress,omitempty"` // Mac Address

	DeviceSeries string `json:"deviceSeries,omitempty"` // Device Series

	SoftwareVersion string `json:"softwareVersion,omitempty"` // Software Version

	ProductVendor string `json:"productVendor,omitempty"` // Product Vendor

	DeviceRole string `json:"deviceRole,omitempty"` // Device Role

	DeviceType string `json:"deviceType,omitempty"` // Device Type

	CommunicationState string `json:"communicationState,omitempty"` // Communication State

	CollectionStatus string `json:"collectionStatus,omitempty"` // Collection Status

	HaStatus string `json:"haStatus,omitempty"` // Ha Status

	LastBootTime *int `json:"lastBootTime,omitempty"` // Last Boot Time

	SiteHierarchyID string `json:"siteHierarchyId,omitempty"` // Site Hierarchy Id

	SiteHierarchy string `json:"siteHierarchy,omitempty"` // Site Hierarchy

	SiteID string `json:"siteId,omitempty"` // Site Id

	DeviceGroupHierarchyID string `json:"deviceGroupHierarchyId,omitempty"` // Device Group Hierarchy Id

	TagNames []string `json:"tagNames,omitempty"` // Tag Names

	StackType string `json:"stackType,omitempty"` // Stack Type

	OsType string `json:"osType,omitempty"` // Os Type

	RingStatus *bool `json:"ringStatus,omitempty"` // Ring Status

	MaintenanceModeEnabled *bool `json:"maintenanceModeEnabled,omitempty"` // Maintenance Mode Enabled

	UpTime *int `json:"upTime,omitempty"` // Up Time

	IPv4Address string `json:"ipv4Address,omitempty"` // Ipv4 Address

	IPv6Address string `json:"ipv6Address,omitempty"` // Ipv6 Address

	RedundancyMode string `json:"redundancyMode,omitempty"` // Redundancy Mode

	FeatureFlagList []string `json:"featureFlagList,omitempty"` // Feature Flag List

	HaLastResetReason string `json:"haLastResetReason,omitempty"` // Ha Last Reset Reason

	RedundancyPeerStateDerived string `json:"redundancyPeerStateDerived,omitempty"` // Redundancy Peer State Derived

	RedundancyPeerState string `json:"redundancyPeerState,omitempty"` // Redundancy Peer State

	RedundancyStateDerived string `json:"redundancyStateDerived,omitempty"` // Redundancy State Derived

	RedundancyState string `json:"redundancyState,omitempty"` // Redundancy State

	WiredClientCount *int `json:"wiredClientCount,omitempty"` // Wired Client Count

	WirelessClientCount *int `json:"wirelessClientCount,omitempty"` // Wireless Client Count

	PortCount *int `json:"portCount,omitempty"` // Port Count

	PhysicalPortCount *int `json:"physicalPortCount,omitempty"` // Physical Port Count

	VirtualPortCount *int `json:"virtualPortCount,omitempty"` // Virtual Port Count

	ClientCount *int `json:"clientCount,omitempty"` // Client Count

	ApDetails *ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDV1ResponseApDetails `json:"apDetails,omitempty"` //

	MetricsDetails *ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDV1ResponseMetricsDetails `json:"metricsDetails,omitempty"` //

	FabricDetails *ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDV1ResponseFabricDetails `json:"fabricDetails,omitempty"` //

	SwitchPoeDetails *ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDV1ResponseSwitchPoeDetails `json:"switchPoeDetails,omitempty"` //

	FabricMetricsDetails *ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDV1ResponseFabricMetricsDetails `json:"fabricMetricsDetails,omitempty"` //

	AggregateAttributes *[]ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDV1ResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDV1ResponseApDetails struct {
	ConnectedWlcName string `json:"connectedWlcName,omitempty"` // Connected Wlc Name

	PolicyTagName string `json:"policyTagName,omitempty"` // Policy Tag Name

	ApOperationalState string `json:"apOperationalState,omitempty"` // Ap Operational State

	PowerSaveMode string `json:"powerSaveMode,omitempty"` // Power Save Mode

	OperationalMode string `json:"operationalMode,omitempty"` // Operational Mode

	ResetReason string `json:"resetReason,omitempty"` // Reset Reason

	Protocol string `json:"protocol,omitempty"` // Protocol

	PowerMode string `json:"powerMode,omitempty"` // Power Mode

	ConnectedTime *int `json:"connectedTime,omitempty"` // Connected Time

	LedFlashEnabled *bool `json:"ledFlashEnabled,omitempty"` // Led Flash Enabled

	LedFlashSeconds *int `json:"ledFlashSeconds,omitempty"` // Led Flash Seconds

	SubMode string `json:"subMode,omitempty"` // Sub Mode

	HomeApEnabled *bool `json:"homeApEnabled,omitempty"` // Home Ap Enabled

	PowerType string `json:"powerType,omitempty"` // Power Type

	ApType string `json:"apType,omitempty"` // Ap Type

	AdminState string `json:"adminState,omitempty"` // Admin State

	IcapCapability string `json:"icapCapability,omitempty"` // Icap Capability

	RegulatoryDomain string `json:"regulatoryDomain,omitempty"` // Regulatory Domain

	EthernetMac string `json:"ethernetMac,omitempty"` // Ethernet Mac

	RfTagName string `json:"rfTagName,omitempty"` // Rf Tag Name

	SiteTagName string `json:"siteTagName,omitempty"` // Site Tag Name

	PowerSaveModeCapable string `json:"powerSaveModeCapable,omitempty"` // Power Save Mode Capable

	PowerProfile string `json:"powerProfile,omitempty"` // Power Profile

	FlexGroup string `json:"flexGroup,omitempty"` // Flex Group

	PowerCalendarProfile string `json:"powerCalendarProfile,omitempty"` // Power Calendar Profile

	ApGroup string `json:"apGroup,omitempty"` // Ap Group

	Radios *[]ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDV1ResponseApDetailsRadios `json:"radios,omitempty"` //
}
type ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDV1ResponseApDetailsRadios struct {
	ID string `json:"id,omitempty"` // Id

	Band string `json:"band,omitempty"` // Band

	Noise *int `json:"noise,omitempty"` // Noise

	AirQuality *float64 `json:"airQuality,omitempty"` // Air Quality

	Interference *float64 `json:"interference,omitempty"` // Interference

	TrafficUtil *int `json:"trafficUtil,omitempty"` // Traffic Util

	Utilization *float64 `json:"utilization,omitempty"` // Utilization

	ClientCount *int `json:"clientCount,omitempty"` // Client Count
}
type ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDV1ResponseMetricsDetails struct {
	OverallHealthScore *int `json:"overallHealthScore,omitempty"` // Overall Health Score

	CPUUtilization *float64 `json:"cpuUtilization,omitempty"` // Cpu Utilization

	CPUScore *int `json:"cpuScore,omitempty"` // Cpu Score

	MemoryUtilization *float64 `json:"memoryUtilization,omitempty"` // Memory Utilization

	MemoryScore *int `json:"memoryScore,omitempty"` // Memory Score

	AvgTemperature *float64 `json:"avgTemperature,omitempty"` // Avg Temperature

	MaxTemperature *float64 `json:"maxTemperature,omitempty"` // Max Temperature

	DiscardScore *int `json:"discardScore,omitempty"` // Discard Score

	DiscardInterfaces []string `json:"discardInterfaces,omitempty"` // Discard Interfaces

	ErrorScore *int `json:"errorScore,omitempty"` // Error Score

	ErrorInterfaces []string `json:"errorInterfaces,omitempty"` // Error Interfaces

	InterDeviceLinkScore *int `json:"interDeviceLinkScore,omitempty"` // Inter Device Link Score

	InterDeviceConnectedDownInterfaces []string `json:"interDeviceConnectedDownInterfaces,omitempty"` // Inter Device Connected Down Interfaces

	LinkUtilizationScore *int `json:"linkUtilizationScore,omitempty"` // Link Utilization Score

	HighLinkUtilizationInterfaces []string `json:"highLinkUtilizationInterfaces,omitempty"` // High Link Utilization Interfaces

	FreeTimerScore *int `json:"freeTimerScore,omitempty"` // Free Timer Score

	FreeTimer *float64 `json:"freeTimer,omitempty"` // Free Timer

	PacketPoolScore *int `json:"packetPoolScore,omitempty"` // Packet Pool Score

	PacketPool *int `json:"packetPool,omitempty"` // Packet Pool

	FreeMemoryBufferScore *int `json:"freeMemoryBufferScore,omitempty"` // Free Memory Buffer Score

	FreeMemoryBuffer *float64 `json:"freeMemoryBuffer,omitempty"` // Free Memory Buffer

	WqePoolScore *int `json:"wqePoolScore,omitempty"` // Wqe Pool Score

	WqePool *int `json:"wqePool,omitempty"` // Wqe Pool

	ApCount *int `json:"apCount,omitempty"` // Ap Count

	NoiseScore *int `json:"noiseScore,omitempty"` // Noise Score

	UtilizationScore *int `json:"utilizationScore,omitempty"` // Utilization Score

	InterferenceScore *int `json:"interferenceScore,omitempty"` // Interference Score

	AirQualityScore *int `json:"airQualityScore,omitempty"` // Air Quality Score
}
type ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDV1ResponseFabricDetails struct {
	FabricRole []string `json:"fabricRole,omitempty"` // Fabric Role

	FabricSiteName string `json:"fabricSiteName,omitempty"` // Fabric Site Name

	TransitFabrics []string `json:"transitFabrics,omitempty"` // Transit Fabrics

	L2Vns []string `json:"l2Vns,omitempty"` // L2 Vns

	L3Vns []string `json:"l3Vns,omitempty"` // L3 Vns

	FabricSiteID string `json:"fabricSiteId,omitempty"` // Fabric Site Id

	NetworkProtocol string `json:"networkProtocol,omitempty"` // Network Protocol
}
type ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDV1ResponseSwitchPoeDetails struct {
	PortCount *int `json:"portCount,omitempty"` // Port Count

	UsedPortCount *int `json:"usedPortCount,omitempty"` // Used Port Count

	FreePortCount *int `json:"freePortCount,omitempty"` // Free Port Count

	PowerConsumed *float64 `json:"powerConsumed,omitempty"` // Power Consumed

	PoePowerConsumed *int `json:"poePowerConsumed,omitempty"` // Poe Power Consumed

	SystemPowerConsumed *float64 `json:"systemPowerConsumed,omitempty"` // System Power Consumed

	PowerBudget *int `json:"powerBudget,omitempty"` // Power Budget

	PoePowerAllocated *float64 `json:"poePowerAllocated,omitempty"` // Poe Power Allocated

	SystemPowerAllocated *int `json:"systemPowerAllocated,omitempty"` // System Power Allocated

	PowerRemaining *float64 `json:"powerRemaining,omitempty"` // Power Remaining

	PoeVersion string `json:"poeVersion,omitempty"` // Poe Version

	ChassisCount *int `json:"chassisCount,omitempty"` // Chassis Count

	ModuleCount *int `json:"moduleCount,omitempty"` // Module Count

	ModuleDetails *[]ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDV1ResponseSwitchPoeDetailsModuleDetails `json:"moduleDetails,omitempty"` //
}
type ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDV1ResponseSwitchPoeDetailsModuleDetails struct {
	ModuleID string `json:"moduleId,omitempty"` // Module Id

	ChassisID string `json:"chassisId,omitempty"` // Chassis Id

	ModulePortCount *int `json:"modulePortCount,omitempty"` // Module Port Count

	ModuleUsedPortCount *int `json:"moduleUsedPortCount,omitempty"` // Module Used Port Count

	ModuleFreePortCount *int `json:"moduleFreePortCount,omitempty"` // Module Free Port Count

	ModulePowerConsumed *float64 `json:"modulePowerConsumed,omitempty"` // Module Power Consumed

	ModulePoePowerConsumed *int `json:"modulePoePowerConsumed,omitempty"` // Module Poe Power Consumed

	ModuleSystemPowerConsumed *float64 `json:"moduleSystemPowerConsumed,omitempty"` // Module System Power Consumed

	ModulePowerBudget *int `json:"modulePowerBudget,omitempty"` // Module Power Budget

	ModulePoePowerAllocated *float64 `json:"modulePoePowerAllocated,omitempty"` // Module Poe Power Allocated

	ModuleSystemPowerAllocated *int `json:"moduleSystemPowerAllocated,omitempty"` // Module System Power Allocated

	ModulePowerRemaining *float64 `json:"modulePowerRemaining,omitempty"` // Module Power Remaining

	InterfacePowerMax *int `json:"interfacePowerMax,omitempty"` // Interface Power Max
}
type ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDV1ResponseFabricMetricsDetails struct {
	OverallFabricScore *int `json:"overallFabricScore,omitempty"` // Overall Fabric Score

	FabricTransitScore *int `json:"fabricTransitScore,omitempty"` // Fabric Transit Score

	FabricSiteScore *int `json:"fabricSiteScore,omitempty"` // Fabric Site Score

	FabricVnScore *int `json:"fabricVnScore,omitempty"` // Fabric Vn Score

	FabsiteFcpScore *int `json:"fabsiteFcpScore,omitempty"` // Fabsite Fcp Score

	FabsiteInfraScore *int `json:"fabsiteInfraScore,omitempty"` // Fabsite Infra Score

	FabsiteFsconnScore *int `json:"fabsiteFsconnScore,omitempty"` // Fabsite Fsconn Score

	VnExitScore *int `json:"vnExitScore,omitempty"` // Vn Exit Score

	VnFcpScore *int `json:"vnFcpScore,omitempty"` // Vn Fcp Score

	VnStatusScore *int `json:"vnStatusScore,omitempty"` // Vn Status Score

	VnServiceScore *int `json:"vnServiceScore,omitempty"` // Vn Service Score

	TransitControlPlaneScore *int `json:"transitControlPlaneScore,omitempty"` // Transit Control Plane Score

	TransitServicesScore *int `json:"transitServicesScore,omitempty"` // Transit Services Score

	TCPConnScore *int `json:"tcpConnScore,omitempty"` // Tcp Conn Score

	BgpBgpSiteScore *int `json:"bgpBgpSiteScore,omitempty"` // Bgp Bgp Site Score

	VniStatusScore *int `json:"vniStatusScore,omitempty"` // Vni Status Score

	PubsubTransitConnScore *int `json:"pubsubTransitConnScore,omitempty"` // Pubsub Transit Conn Score

	BgpPeerInfraVnScore *int `json:"bgpPeerInfraVnScore,omitempty"` // Bgp Peer Infra Vn Score

	InternetAvailScore *int `json:"internetAvailScore,omitempty"` // Internet Avail Score

	BgpEvpnScore *int `json:"bgpEvpnScore,omitempty"` // Bgp Evpn Score

	LispTransitConnScore *int `json:"lispTransitConnScore,omitempty"` // Lisp Transit Conn Score

	CtsEnvDataDownloadScore *int `json:"ctsEnvDataDownloadScore,omitempty"` // Cts Env Data Download Score

	PubsubInfraVnScore *int `json:"pubsubInfraVnScore,omitempty"` // Pubsub Infra Vn Score

	PeerScore *int `json:"peerScore,omitempty"` // Peer Score

	BgpPeerScore *int `json:"bgpPeerScore,omitempty"` // Bgp Peer Score

	RemoteInternetAvailScore *int `json:"remoteInternetAvailScore,omitempty"` // Remote Internet Avail Score

	BgpTCPScore *int `json:"bgpTcpScore,omitempty"` // Bgp Tcp Score

	PubsubSessionScore *int `json:"pubsubSessionScore,omitempty"` // Pubsub Session Score

	AAAStatusScore *int `json:"aaaStatusScore,omitempty"` // Aaa Status Score

	LispCpConnScore *int `json:"lispCpConnScore,omitempty"` // Lisp Cp Conn Score

	BgpPubsubSiteScore *int `json:"bgpPubsubSiteScore,omitempty"` // Bgp Pubsub Site Score

	McastScore *int `json:"mcastScore,omitempty"` // Mcast Score

	PortChannelScore *int `json:"portChannelScore,omitempty"` // Port Channel Score
}
type ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDV1ResponseAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function

	Value *float64 `json:"value,omitempty"` // Value
}
type ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1 struct {
	Response *[]ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1Response `json:"response,omitempty"` //

	Page *ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1Page `json:"page,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1Response struct {
	Timestamp *float64 `json:"timestamp,omitempty"` // Timestamp

	Attributes *[]ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1ResponseAttributes `json:"attributes,omitempty"` //

	AggregateAttributes *[]ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1ResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Groups *[]ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1ResponseGroups `json:"groups,omitempty"` //
}
type ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1ResponseAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1ResponseAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function

	Value *float64 `json:"value,omitempty"` // Value
}
type ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1ResponseGroups struct {
	ID string `json:"id,omitempty"` // Id

	Attributes *[]ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1ResponseGroupsAttributes `json:"attributes,omitempty"` //

	AggregateAttributes *[]ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1ResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1ResponseGroupsAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1ResponseGroupsAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function

	Value *float64 `json:"value,omitempty"` // Value
}
type ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	Count *int `json:"count,omitempty"` // Count

	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type ResponseDevicesGetPlannedAccessPointsForBuildingV1 struct {
	Response *[]ResponseDevicesGetPlannedAccessPointsForBuildingV1Response `json:"response,omitempty"` //

	Version *int `json:"version,omitempty"` // Version of the api response model

	Total *int `json:"total,omitempty"` // Total number of the planned access points
}
type ResponseDevicesGetPlannedAccessPointsForBuildingV1Response struct {
	Attributes *ResponseDevicesGetPlannedAccessPointsForBuildingV1ResponseAttributes `json:"attributes,omitempty"` //

	Location *ResponseDevicesGetPlannedAccessPointsForBuildingV1ResponseLocation `json:"location,omitempty"` //

	Position *ResponseDevicesGetPlannedAccessPointsForBuildingV1ResponsePosition `json:"position,omitempty"` //

	RadioCount *int `json:"radioCount,omitempty"` // Number of radios of the planned access point

	Radios *[]ResponseDevicesGetPlannedAccessPointsForBuildingV1ResponseRadios `json:"radios,omitempty"` //

	IsSensor *bool `json:"isSensor,omitempty"` // Determines if the planned access point is sensor or not
}
type ResponseDevicesGetPlannedAccessPointsForBuildingV1ResponseAttributes struct {
	ID *float64 `json:"id,omitempty"` // Unique id of the planned access point

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance uuid of the planned access point

	Name string `json:"name,omitempty"` // Display name of the planned access point

	TypeString string `json:"typeString,omitempty"` // Type string representation of the planned access point

	Domain string `json:"domain,omitempty"` // Service domain to which the planned access point belongs

	HeirarchyName string `json:"heirarchyName,omitempty"` // Hierarchy name of the planned access point

	Source string `json:"source,omitempty"` // Source of the data used to create the planned access point

	CreateDate *float64 `json:"createDate,omitempty"` // Created date of the planned access point

	MacAddress string `json:"macAddress,omitempty"` // MAC address of the planned access point
}
type ResponseDevicesGetPlannedAccessPointsForBuildingV1ResponseLocation struct {
	Altitude *float64 `json:"altitude,omitempty"` // Altitude of the planned access point's location

	Lattitude *float64 `json:"lattitude,omitempty"` // Latitude of the planned access point's location

	Longtitude *float64 `json:"longtitude,omitempty"` // Longitude of the planned access point's location
}
type ResponseDevicesGetPlannedAccessPointsForBuildingV1ResponsePosition struct {
	X *float64 `json:"x,omitempty"` // x-coordinate of the planned access point on the map, 0,0 point being the top-left corner

	Y *float64 `json:"y,omitempty"` // y-coordinate of the planned access point on the map, 0,0 point being the top-left corner

	Z *float64 `json:"z,omitempty"` // z-coordinate, or height, of the planned access point on the map
}
type ResponseDevicesGetPlannedAccessPointsForBuildingV1ResponseRadios struct {
	Attributes *ResponseDevicesGetPlannedAccessPointsForBuildingV1ResponseRadiosAttributes `json:"attributes,omitempty"` //

	Antenna *ResponseDevicesGetPlannedAccessPointsForBuildingV1ResponseRadiosAntenna `json:"antenna,omitempty"` //

	IsSensor *bool `json:"isSensor,omitempty"` // Determines if it is sensor or not
}
type ResponseDevicesGetPlannedAccessPointsForBuildingV1ResponseRadiosAttributes struct {
	ID *int `json:"id,omitempty"` // Id of the radio

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid of the radio

	SlotID *float64 `json:"slotId,omitempty"` // Slot number in which the radio resides in the parent access point

	IfTypeString string `json:"ifTypeString,omitempty"` // String representation of native band

	IfTypeSubband string `json:"ifTypeSubband,omitempty"` // Sub band type of the radio

	Channel *float64 `json:"channel,omitempty"` // Channel in which the radio operates

	ChannelString string `json:"channelString,omitempty"` // Channel string representation

	IfMode string `json:"ifMode,omitempty"` // IF mode of the radio

	TxPowerLevel *float64 `json:"txPowerLevel,omitempty"` // Tx Power at which this radio operates (in dBm)
}
type ResponseDevicesGetPlannedAccessPointsForBuildingV1ResponseRadiosAntenna struct {
	Name string `json:"name,omitempty"` // Name of the antenna

	Type string `json:"type,omitempty"` // Type of the antenna associated with this radio

	Mode string `json:"mode,omitempty"` // Mode of the antenna associated with this radio

	AzimuthAngle *float64 `json:"azimuthAngle,omitempty"` // Azimuth angle of the antenna

	ElevationAngle *float64 `json:"elevationAngle,omitempty"` // Elevation angle of the antenna

	Gain *float64 `json:"gain,omitempty"` // Gain of the antenna
}
type ResponseDevicesGetDeviceDetailV1 struct {
	Response *ResponseDevicesGetDeviceDetailV1Response `json:"response,omitempty"` //
}
type ResponseDevicesGetDeviceDetailV1Response struct {
	NoiseScore *int `json:"noiseScore,omitempty"` // Device (AP) WIFI signal noise health score

	PolicyTagName string `json:"policyTagName,omitempty"` // Device (AP) policy tag

	InterferenceScore *int `json:"interferenceScore,omitempty"` // Device (AP) WIFI signal interference health score

	OpState string `json:"opState,omitempty"` // Operation state of device (AP)

	PowerSaveMode string `json:"powerSaveMode,omitempty"` // Device power save mode

	Mode string `json:"mode,omitempty"` // Device mode (AP)

	ResetReason string `json:"resetReason,omitempty"` // Device reset reason

	NwDeviceRole string `json:"nwDeviceRole,omitempty"` // Device role

	Protocol string `json:"protocol,omitempty"` // Protocol code

	PowerMode string `json:"powerMode,omitempty"` // Device's power mode

	ConnectedTime string `json:"connectedTime,omitempty"` // UTC timestamp

	RingStatus *bool `json:"ringStatus,omitempty"` // Device's ring status

	LedFlashSeconds string `json:"ledFlashSeconds,omitempty"` // LED flash seconds

	IPAddrManagementIPAddr string `json:"ip_addr_managementIpAddr,omitempty"` // Device's management IP address

	StackType string `json:"stackType,omitempty"` // Device stack type (applicable for stackable devices)

	SubMode string `json:"subMode,omitempty"` // Device submode

	SerialNumber string `json:"serialNumber,omitempty"` // Device serial number

	NwDeviceName string `json:"nwDeviceName,omitempty"` // Device name

	DeviceGroupHierarchyID string `json:"deviceGroupHierarchyId,omitempty"` // Device group site hierarchy UUID

	CPU *float64 `json:"cpu,omitempty"` // Device CPU utilization

	Utilization string `json:"utilization,omitempty"` // Device utilization

	NwDeviceID string `json:"nwDeviceId,omitempty"` // Device's UUID

	SiteHierarchyGraphID string `json:"siteHierarchyGraphId,omitempty"` // Site hierarchy UUID in which device is assigned to

	NwDeviceFamily string `json:"nwDeviceFamily,omitempty"` // Device faimly string

	MacAddress string `json:"macAddress,omitempty"` // Device MAC address

	HomeApEnabled string `json:"homeApEnabled,omitempty"` // Home Ap Enabled

	DeviceSeries string `json:"deviceSeries,omitempty"` // Device series string

	CollectionStatus string `json:"collectionStatus,omitempty"` // Device's telemetry data collection status for DNAC

	UtilizationScore *int `json:"utilizationScore,omitempty"` // Device utilization health score

	MaintenanceMode *bool `json:"maintenanceMode,omitempty"` // Whether device is in maintenance mode

	Interference string `json:"interference,omitempty"` // Device (AP) WIFI signal interference

	SoftwareVersion string `json:"softwareVersion,omitempty"` // Device's software version string

	TagIDList *[]ResponseDevicesGetDeviceDetailV1ResponseTagIDList `json:"tagIdList,omitempty"` // Tag ID List

	PowerType string `json:"powerType,omitempty"` // Device (AP) power type

	OverallHealth *int `json:"overallHealth,omitempty"` // Device's overall health score

	ManagementIPAddr string `json:"managementIpAddr,omitempty"` // Management IP address of the device

	Memory string `json:"memory,omitempty"` // Device memory utilization

	CommunicationState string `json:"communicationState,omitempty"` // Device communication state

	ApType string `json:"apType,omitempty"` // Ap Type

	AdminState string `json:"adminState,omitempty"` // Device (AP) admin state

	Noise string `json:"noise,omitempty"` // Device (AP) WIFI signal noise

	IcapCapability string `json:"icapCapability,omitempty"` // Device (AP) ICAP capability bit values

	RegulatoryDomain string `json:"regulatoryDomain,omitempty"` // Device (AP) WIFI domain

	EthernetMac string `json:"ethernetMac,omitempty"` // Device (AP) ethernet MAC address

	NwDeviceType string `json:"nwDeviceType,omitempty"` // Device type

	AirQuality string `json:"airQuality,omitempty"` // Device (AP) WIFI air quality

	RfTagName string `json:"rfTagName,omitempty"` // Device (AP) RF tag name

	SiteTagName string `json:"siteTagName,omitempty"` // Device (AP) site tag name

	PlatformID string `json:"platformId,omitempty"` // Device's platform ID

	UpTime string `json:"upTime,omitempty"` // Device up time

	MemoryScore *int `json:"memoryScore,omitempty"` // Device's memory usage score

	PowerSaveModeCapable string `json:"powerSaveModeCapable,omitempty"` // Device (AP) power save mode capability

	PowerProfile string `json:"powerProfile,omitempty"` // Device (AP) power profile name

	AirQualityScore *int `json:"airQualityScore,omitempty"` // Device (AP) air quality health score

	Location string `json:"location,omitempty"` // Device's site hierarchy UUID

	FlexGroup string `json:"flexGroup,omitempty"` // Deivce (A) flexconnect group

	LastBootTime *float64 `json:"lastBootTime,omitempty"` // Device's last boot UTC timestamp

	PowerCalendarProfile string `json:"powerCalendarProfile,omitempty"` // Device (AP) power calendar profile name

	ConnectivityStatus *int `json:"connectivityStatus,omitempty"` // Device connectivity status

	LedFlashEnabled string `json:"ledFlashEnabled,omitempty"` // Device (AP) LED flash

	CPUScore *int `json:"cpuScore,omitempty"` // Device's CPU usage score

	AvgTemperature *float64 `json:"avgTemperature,omitempty"` // Device's average temperature

	MaxTemperature *float64 `json:"maxTemperature,omitempty"` // Device's max temperature

	HaStatus string `json:"haStatus,omitempty"` // Device's HA status

	OsType string `json:"osType,omitempty"` // Device's OS type

	Timestamp *int `json:"timestamp,omitempty"` // UTC timestamp of the device health data

	ApGroup string `json:"apGroup,omitempty"` // Device (AP) AP group

	RedundancyMode string `json:"redundancyMode,omitempty"` // Device redundancy mode

	FeatureFlagList []string `json:"featureFlagList,omitempty"` // List of device feature capabilities

	FreeMbufScore *int `json:"freeMbufScore,omitempty"` // Free memory buffer health score

	HALastResetReason string `json:"HALastResetReason,omitempty"` // Last HA reset reason

	WqeScore *int `json:"wqeScore,omitempty"` // WQE health score

	RedundancyPeerStateDerived string `json:"redundancyPeerStateDerived,omitempty"` // Redundancy Peer State Derived

	FreeTimerScore *int `json:"freeTimerScore,omitempty"` // Free Timer Score

	RedundancyPeerState string `json:"redundancyPeerState,omitempty"` // Redundancy Peer State

	RedundancyStateDerived string `json:"redundancyStateDerived,omitempty"` // Derived redundancy state

	RedundancyState string `json:"redundancyState,omitempty"` // Redundancy state

	PacketPoolScore *int `json:"packetPoolScore,omitempty"` // Device packet pool health score

	FreeTimer *float64 `json:"freeTimer,omitempty"` // Free timer of the device

	PacketPool *float64 `json:"packetPool,omitempty"` // Packet pool of the device

	Wqe *float64 `json:"wqe,omitempty"` // WQE of the device

	FreeMbuf *float64 `json:"freeMbuf,omitempty"` // Free memory buffer of the device
}
type ResponseDevicesGetDeviceDetailV1ResponseTagIDList interface{}
type ResponseDevicesGetDeviceEnrichmentDetailsV1 []ResponseItemDevicesGetDeviceEnrichmentDetailsV1 // Array of ResponseDevicesGetDeviceEnrichmentDetailsV1
type ResponseItemDevicesGetDeviceEnrichmentDetailsV1 struct {
	DeviceDetails *ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetails `json:"deviceDetails,omitempty"` //
}
type ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetails struct {
	Family string `json:"family,omitempty"` // Device Family

	Type string `json:"type,omitempty"` // Device Type

	Location *ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsLocation `json:"location,omitempty"` // Device location - Site hierarchy

	ErrorCode string `json:"errorCode,omitempty"` // Inventory status error code

	MacAddress string `json:"macAddress,omitempty"` // Device MAC address

	Role string `json:"role,omitempty"` // Device role

	ApManagerInterfaceIP string `json:"apManagerInterfaceIp,omitempty"` // IP address of WLC on AP manager interface

	AssociatedWlcIP string `json:"associatedWlcIp,omitempty"` // Associated WLC IP address of the AP device

	BootDateTime string `json:"bootDateTime,omitempty"` // Device's last boot UTC timestamp

	CollectionStatus string `json:"collectionStatus,omitempty"` // Device's telemetry data collection status for DNAC

	InterfaceCount string `json:"interfaceCount,omitempty"` // Number of interfaces on the device

	LineCardCount string `json:"lineCardCount,omitempty"` // Number of linecards on the device

	LineCardID string `json:"lineCardId,omitempty"` // IDs of linecards of the device

	ManagementIPAddress string `json:"managementIpAddress,omitempty"` // Device Management Ip Address

	MemorySize string `json:"memorySize,omitempty"` // Processor memory size

	PlatformID string `json:"platformId,omitempty"` // Device's platform ID

	ReachabilityFailureReason string `json:"reachabilityFailureReason,omitempty"` // Failure reason for unreachable devices

	ReachabilityStatus string `json:"reachabilityStatus,omitempty"` // Reachability Status of the Device(Reachable/Unreachable)

	SNMPContact string `json:"snmpContact,omitempty"` // SNMP contact on device

	SNMPLocation string `json:"snmpLocation,omitempty"` // SNMP location on device

	TunnelUDPPort *ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsTunnelUDPPort `json:"tunnelUdpPort,omitempty"` // Mobility protocol port is stored as tunneludpport for WLC

	WaasDeviceMode *ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsWaasDeviceMode `json:"waasDeviceMode,omitempty"` // WAAS device mode

	Series string `json:"series,omitempty"` // Device Series

	InventoryStatusDetail string `json:"inventoryStatusDetail,omitempty"` // Status detail of inventory sync

	CollectionInterval string `json:"collectionInterval,omitempty"` // Re sync Interval of the device

	SerialNumber string `json:"serialNumber,omitempty"` // Device Serial Number

	SoftwareVersion string `json:"softwareVersion,omitempty"` // Device Software Version

	RoleSource string `json:"roleSource,omitempty"` // Role source as manual / auto

	Hostname string `json:"hostname,omitempty"` // Device Hostname

	UpTime string `json:"upTime,omitempty"` // Device's uptime

	LastUpdateTime *int `json:"lastUpdateTime,omitempty"` // Time in epoch when the network device info last got updated

	ErrorDescription string `json:"errorDescription,omitempty"` // Inventory status description

	LocationName *ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsLocationName `json:"locationName,omitempty"` // [Deprecated] Name of the associated location

	TagCount string `json:"tagCount,omitempty"` // Number of tags associated with the device

	LastUpdated string `json:"lastUpdated,omitempty"` // Time when the network device info last got updated

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid of the device

	ID string `json:"id,omitempty"` // Device's UUID

	NeighborTopology *[]ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsNeighborTopology `json:"neighborTopology,omitempty"` //
}
type ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsLocation interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsTunnelUDPPort interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsWaasDeviceMode interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsLocationName interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsNeighborTopology struct {
	Nodes *[]ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsNeighborTopologyNodes `json:"nodes,omitempty"` //

	Links *[]ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsNeighborTopologyLinks `json:"links,omitempty"` //
}
type ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsNeighborTopologyNodes struct {
	Role string `json:"role,omitempty"` // Role of the Node

	Name string `json:"name,omitempty"` // Hostname of the Node

	ID string `json:"id,omitempty"` // Id of the Node

	Description string `json:"description,omitempty"` // Description of the Node

	DeviceType string `json:"deviceType,omitempty"` // Device type of the node, like switch, AP, WCL,GateWay

	PlatformID string `json:"platformId,omitempty"` // Type of platform

	Family string `json:"family,omitempty"` // Device Family of the Node

	IP string `json:"ip,omitempty"` // IP Address of the Node

	SoftwareVersion string `json:"softwareVersion,omitempty"` // Software Version of the Node

	UserID *ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsNeighborTopologyNodesUserID `json:"userId,omitempty"` // User Id of the Node

	NodeType string `json:"nodeType,omitempty"` // Type of the Node

	RadioFrequency *ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsNeighborTopologyNodesRadioFrequency `json:"radioFrequency,omitempty"` // Frequency of wireless radio channel

	Clients *ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsNeighborTopologyNodesClients `json:"clients,omitempty"` // Number of clients

	Count *ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsNeighborTopologyNodesCount `json:"count,omitempty"` // The number of group nodes (for ap sepecifically)

	HealthScore *int `json:"healthScore,omitempty"` // The total health score of the node

	Level *float64 `json:"level,omitempty"` // The level index to be used by UI widget (starts from 0)

	FabricGroup *ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsNeighborTopologyNodesFabricGroup `json:"fabricGroup,omitempty"` // Fabric device group name

	ConnectedDevice *ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsNeighborTopologyNodesConnectedDevice `json:"connectedDevice,omitempty"` // The connected device to show the connected switch to wlc
}
type ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsNeighborTopologyNodesUserID interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsNeighborTopologyNodesRadioFrequency interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsNeighborTopologyNodesClients interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsNeighborTopologyNodesCount interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsNeighborTopologyNodesFabricGroup interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsNeighborTopologyNodesConnectedDevice interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsNeighborTopologyLinks struct {
	Source string `json:"source,omitempty"` // Edge line starting node

	LinkStatus string `json:"linkStatus,omitempty"` // The status of the link (up/down)

	Label *[]ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsNeighborTopologyLinksLabel `json:"label,omitempty"` // The details of the edge

	Target string `json:"target,omitempty"` // End node of the edge line

	ID *ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsNeighborTopologyLinksID `json:"id,omitempty"` // Id of the node

	PortUtilization *ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsNeighborTopologyLinksPortUtilization `json:"portUtilization,omitempty"` // Number of clients
}
type ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsNeighborTopologyLinksLabel interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsNeighborTopologyLinksID interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsNeighborTopologyLinksPortUtilization interface{}
type ResponseDevicesDevicesV1 struct {
	Version string `json:"version,omitempty"` // Response data's version string

	TotalCount *int `json:"totalCount,omitempty"` // Total number of devices

	Response *[]ResponseDevicesDevicesV1Response `json:"response,omitempty"` //
}
type ResponseDevicesDevicesV1Response struct {
	DeviceType string `json:"deviceType,omitempty"` // Device type

	CPUUtilization *float64 `json:"cpuUtilization,omitempty"` // Device's CPU utilization

	OverallHealth *int `json:"overallHealth,omitempty"` // Overall health score

	UtilizationHealth *ResponseDevicesDevicesV1ResponseUtilizationHealth `json:"utilizationHealth,omitempty"` //

	AirQualityHealth *ResponseDevicesDevicesV1ResponseAirQualityHealth `json:"airQualityHealth,omitempty"` //

	IPAddress string `json:"ipAddress,omitempty"` // Management IP address of the device

	CPUHealth *int `json:"cpuHealth,omitempty"` // Device CPU health score

	DeviceFamily string `json:"deviceFamily,omitempty"` // Device family

	IssueCount *int `json:"issueCount,omitempty"` // Number of issues

	MacAddress string `json:"macAddress,omitempty"` // MAC address of the device

	NoiseHealth *ResponseDevicesDevicesV1ResponseNoiseHealth `json:"noiseHealth,omitempty"` //

	OsVersion string `json:"osVersion,omitempty"` // Device OS version string

	Name string `json:"name,omitempty"` // Device name

	InterfaceLinkErrHealth *int `json:"interfaceLinkErrHealth,omitempty"` // Device (AP) error health score

	MemoryUtilization *float64 `json:"memoryUtilization,omitempty"` // Device memory utilization

	InterDeviceLinkAvailHealth *int `json:"interDeviceLinkAvailHealth,omitempty"` // Device connectivity status

	InterferenceHealth *ResponseDevicesDevicesV1ResponseInterferenceHealth `json:"interferenceHealth,omitempty"` //

	Model string `json:"model,omitempty"` // Device model string

	Location string `json:"location,omitempty"` // Site location in which this device is assigned to

	ReachabilityHealth string `json:"reachabilityHealth,omitempty"` // Device reachability in the network

	Band *ResponseDevicesDevicesV1ResponseBand `json:"band,omitempty"` //

	MemoryUtilizationHealth *int `json:"memoryUtilizationHealth,omitempty"` // Device memory utilization health score

	ClientCount *ResponseDevicesDevicesV1ResponseClientCount `json:"clientCount,omitempty"` //

	AvgTemperature *float64 `json:"avgTemperature,omitempty"` // Average device (switch) temperature

	MaxTemperature *float64 `json:"maxTemperature,omitempty"` // Max device (switch) temperature

	InterDeviceLinkAvailFabric *int `json:"interDeviceLinkAvailFabric,omitempty"` // Device uplink health

	ApCount *int `json:"apCount,omitempty"` // Number of AP count

	FreeTimerScore *int `json:"freeTimerScore,omitempty"` // Device free timer health score

	FreeTimer *float64 `json:"freeTimer,omitempty"` // Device free timer

	PacketPoolHealth *int `json:"packetPoolHealth,omitempty"` // Device packet pool

	PacketPool *int `json:"packetPool,omitempty"` // Device packet pool

	FreeMemoryBufferHealth *int `json:"freeMemoryBufferHealth,omitempty"` // Device free memory buffer health

	FreeMemoryBuffer *float64 `json:"freeMemoryBuffer,omitempty"` // Device free memory

	WqePoolsHealth *int `json:"wqePoolsHealth,omitempty"` // Device WQE pool health

	WqePools *float64 `json:"wqePools,omitempty"` // Device WQE pool

	WanLinkUtilization *float64 `json:"wanLinkUtilization,omitempty"` // WLAN link utilization

	CPUUlitilization *float64 `json:"cpuUlitilization,omitempty"` // Device's CPU utilization

	UUID string `json:"uuid,omitempty"` // Device UUID
}
type ResponseDevicesDevicesV1ResponseUtilizationHealth struct {
	Radio0 *int `json:"radio0,omitempty"` // Radio0

	Radio1 *int `json:"radio1,omitempty"` // Radio1

	Radio2 *int `json:"radio2,omitempty"` // Radio2

	Radio3 *int `json:"radio3,omitempty"` // Radio3

	Ghz24 *int `json:"Ghz24,omitempty"` // Ghz24

	Ghz50 *int `json:"Ghz50,omitempty"` // Ghz50
}
type ResponseDevicesDevicesV1ResponseAirQualityHealth struct {
	Radio0 *int `json:"radio0,omitempty"` // Radio0

	Radio1 *int `json:"radio1,omitempty"` // Radio1

	Radio2 *int `json:"radio2,omitempty"` // Radio2

	Radio3 *int `json:"radio3,omitempty"` // Radio3

	Ghz24 *int `json:"Ghz24,omitempty"` // Ghz24

	Ghz50 *int `json:"Ghz50,omitempty"` // Ghz50
}
type ResponseDevicesDevicesV1ResponseNoiseHealth struct {
	Radio0 *int `json:"radio0,omitempty"` // Radio0

	Radio1 *int `json:"radio1,omitempty"` // Radio1

	Radio2 *int `json:"radio2,omitempty"` // Radio2

	Radio3 *int `json:"radio3,omitempty"` // Radio3

	Ghz24 *int `json:"Ghz24,omitempty"` // Ghz24

	Ghz50 *int `json:"Ghz50,omitempty"` // Ghz50
}
type ResponseDevicesDevicesV1ResponseInterferenceHealth struct {
	Radio0 *int `json:"radio0,omitempty"` // Radio0

	Radio1 *int `json:"radio1,omitempty"` // Radio1

	Radio2 *int `json:"radio2,omitempty"` // Radio2

	Radio3 *int `json:"radio3,omitempty"` // Radio3

	Ghz24 *int `json:"Ghz24,omitempty"` // Ghz24

	Ghz50 *int `json:"Ghz50,omitempty"` // Ghz50
}
type ResponseDevicesDevicesV1ResponseBand struct {
	Radio0 string `json:"radio0,omitempty"` // Radio0

	Radio1 string `json:"radio1,omitempty"` // Radio1

	Radio2 string `json:"radio2,omitempty"` // Radio2

	Radio3 *int `json:"radio3,omitempty"` // Radio3
}
type ResponseDevicesDevicesV1ResponseClientCount struct {
	Radio0 *int `json:"radio0,omitempty"` // Radio0

	Radio1 *int `json:"radio1,omitempty"` // Radio1

	Radio2 *int `json:"radio2,omitempty"` // Radio2

	Radio3 *int `json:"radio3,omitempty"` // Radio3

	Ghz24 *int `json:"Ghz24,omitempty"` // Ghz24

	Ghz50 *int `json:"Ghz50,omitempty"` // Ghz50
}
type ResponseDevicesUpdatePlannedAccessPointForFloorV1 struct {
	Version string `json:"version,omitempty"` // Response Version e.g. : 1.0

	Response *ResponseDevicesUpdatePlannedAccessPointForFloorV1Response `json:"response,omitempty"` //
}
type ResponseDevicesUpdatePlannedAccessPointForFloorV1Response struct {
	URL string `json:"url,omitempty"` // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5

	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseDevicesCreatePlannedAccessPointForFloorV1 struct {
	Version string `json:"version,omitempty"` // Response Version e.g. : 1.0

	Response *ResponseDevicesCreatePlannedAccessPointForFloorV1Response `json:"response,omitempty"` //
}
type ResponseDevicesCreatePlannedAccessPointForFloorV1Response struct {
	URL string `json:"url,omitempty"` // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5

	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseDevicesGetPlannedAccessPointsForFloorV1 struct {
	Response *[]ResponseDevicesGetPlannedAccessPointsForFloorV1Response `json:"response,omitempty"` //

	Version *int `json:"version,omitempty"` // Version of the api response model

	Total *int `json:"total,omitempty"` // Total number of the planned access points
}
type ResponseDevicesGetPlannedAccessPointsForFloorV1Response struct {
	Attributes *ResponseDevicesGetPlannedAccessPointsForFloorV1ResponseAttributes `json:"attributes,omitempty"` //

	Location *ResponseDevicesGetPlannedAccessPointsForFloorV1ResponseLocation `json:"location,omitempty"` //

	Position *ResponseDevicesGetPlannedAccessPointsForFloorV1ResponsePosition `json:"position,omitempty"` //

	RadioCount *int `json:"radioCount,omitempty"` // Number of radios of the planned access point

	Radios *[]ResponseDevicesGetPlannedAccessPointsForFloorV1ResponseRadios `json:"radios,omitempty"` //

	IsSensor *bool `json:"isSensor,omitempty"` // Determines if the planned access point is sensor or not
}
type ResponseDevicesGetPlannedAccessPointsForFloorV1ResponseAttributes struct {
	ID *float64 `json:"id,omitempty"` // Unique id of the planned access point

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance uuid of the planned access point

	Name string `json:"name,omitempty"` // Display name of the planned access point

	TypeString string `json:"typeString,omitempty"` // Type string representation of the planned access point

	Domain string `json:"domain,omitempty"` // Service domain to which the planned access point belongs

	HeirarchyName string `json:"heirarchyName,omitempty"` // Hierarchy name of the planned access point

	Source string `json:"source,omitempty"` // Source of the data used to create the planned access point

	CreateDate *float64 `json:"createDate,omitempty"` // Created date of the planned access point

	MacAddress string `json:"macAddress,omitempty"` // MAC address of the planned access point
}
type ResponseDevicesGetPlannedAccessPointsForFloorV1ResponseLocation struct {
	Altitude *float64 `json:"altitude,omitempty"` // Altitude of the planned access point's location

	Lattitude *float64 `json:"lattitude,omitempty"` // Latitude of the planned access point's location

	Longtitude *float64 `json:"longtitude,omitempty"` // Longitude of the planned access point's location
}
type ResponseDevicesGetPlannedAccessPointsForFloorV1ResponsePosition struct {
	X *float64 `json:"x,omitempty"` // x-coordinate of the planned access point on the map, 0,0 point being the top-left corner

	Y *float64 `json:"y,omitempty"` // y-coordinate of the planned access point on the map, 0,0 point being the top-left corner

	Z *float64 `json:"z,omitempty"` // z-coordinate, or height, of the planned access point on the map
}
type ResponseDevicesGetPlannedAccessPointsForFloorV1ResponseRadios struct {
	Attributes *ResponseDevicesGetPlannedAccessPointsForFloorV1ResponseRadiosAttributes `json:"attributes,omitempty"` //

	Antenna *ResponseDevicesGetPlannedAccessPointsForFloorV1ResponseRadiosAntenna `json:"antenna,omitempty"` //

	IsSensor *bool `json:"isSensor,omitempty"` // Determines if it is sensor or not
}
type ResponseDevicesGetPlannedAccessPointsForFloorV1ResponseRadiosAttributes struct {
	ID *int `json:"id,omitempty"` // Id of the radio

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid of the radio

	SlotID *float64 `json:"slotId,omitempty"` // Slot number in which the radio resides in the parent access point

	IfTypeString string `json:"ifTypeString,omitempty"` // String representation of native band

	IfTypeSubband string `json:"ifTypeSubband,omitempty"` // Sub band type of the radio

	Channel *float64 `json:"channel,omitempty"` // Channel in which the radio operates

	ChannelString string `json:"channelString,omitempty"` // Channel string representation

	IfMode string `json:"ifMode,omitempty"` // IF mode of the radio

	TxPowerLevel *float64 `json:"txPowerLevel,omitempty"` // Tx Power at which this radio operates (in dBm)
}
type ResponseDevicesGetPlannedAccessPointsForFloorV1ResponseRadiosAntenna struct {
	Name string `json:"name,omitempty"` // Name of the antenna

	Type string `json:"type,omitempty"` // Type of the antenna associated with this radio

	Mode string `json:"mode,omitempty"` // Mode of the antenna associated with this radio

	AzimuthAngle *float64 `json:"azimuthAngle,omitempty"` // Azimuth angle of the antenna

	ElevationAngle *float64 `json:"elevationAngle,omitempty"` // Elevation angle of the antenna

	Gain *float64 `json:"gain,omitempty"` // Gain of the antenna
}
type ResponseDevicesDeletePlannedAccessPointForFloorV1 struct {
	Version string `json:"version,omitempty"` // Response Version e.g. : 1.0

	Response *ResponseDevicesDeletePlannedAccessPointForFloorV1Response `json:"response,omitempty"` //
}
type ResponseDevicesDeletePlannedAccessPointForFloorV1Response struct {
	URL string `json:"url,omitempty"` // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5

	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseDevicesGetAllHealthScoreDefinitionsForGivenFiltersV1 struct {
	Response *[]ResponseDevicesGetAllHealthScoreDefinitionsForGivenFiltersV1Response `json:"response,omitempty"` //
}
type ResponseDevicesGetAllHealthScoreDefinitionsForGivenFiltersV1Response struct {
	ID string `json:"id,omitempty"` // Id

	Name string `json:"name,omitempty"` // Name

	DisplayName string `json:"displayName,omitempty"` // Display Name

	DeviceFamily string `json:"deviceFamily,omitempty"` // Device Family

	Description string `json:"description,omitempty"` // Description

	IncludeForOverallHealth *bool `json:"includeForOverallHealth,omitempty"` // Include For Overall Health

	DefinitionStatus string `json:"definitionStatus,omitempty"` // Definition Status

	ThresholdValue *float64 `json:"thresholdValue,omitempty"` // Threshold Value

	SynchronizeToIssueThreshold *bool `json:"synchronizeToIssueThreshold,omitempty"` // Synchronize To Issue Threshold

	LastModified string `json:"lastModified,omitempty"` // Last Modified
}
type ResponseDevicesUpdateHealthScoreDefinitionsV1 struct {
	Response *[]ResponseDevicesUpdateHealthScoreDefinitionsV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesUpdateHealthScoreDefinitionsV1Response struct {
	ID string `json:"id,omitempty"` // Id

	Name string `json:"name,omitempty"` // Name

	DisplayName string `json:"displayName,omitempty"` // Display Name

	DeviceFamily string `json:"deviceFamily,omitempty"` // Device Family

	Description string `json:"description,omitempty"` // Description

	IncludeForOverallHealth *bool `json:"includeForOverallHealth,omitempty"` // Include For Overall Health

	DefinitionStatus string `json:"definitionStatus,omitempty"` // Definition Status

	ThresholdValue *float64 `json:"thresholdValue,omitempty"` // Threshold Value

	SynchronizeToIssueThreshold *bool `json:"synchronizeToIssueThreshold,omitempty"` // Synchronize To Issue Threshold

	LastModified string `json:"lastModified,omitempty"` // Last Modified
}
type ResponseDevicesGetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersV1 struct {
	Response *ResponseDevicesGetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesGetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseDevicesGetHealthScoreDefinitionForTheGivenIDV1 struct {
	Response *[]ResponseDevicesGetHealthScoreDefinitionForTheGivenIDV1Response `json:"response,omitempty"` //
}
type ResponseDevicesGetHealthScoreDefinitionForTheGivenIDV1Response struct {
	ID string `json:"id,omitempty"` // Id

	Name string `json:"name,omitempty"` // Name

	DisplayName string `json:"displayName,omitempty"` // Display Name

	DeviceFamily string `json:"deviceFamily,omitempty"` // Device Family

	Description string `json:"description,omitempty"` // Description

	IncludeForOverallHealth *bool `json:"includeForOverallHealth,omitempty"` // Include For Overall Health

	DefinitionStatus string `json:"definitionStatus,omitempty"` // Definition Status

	ThresholdValue *float64 `json:"thresholdValue,omitempty"` // Threshold Value

	SynchronizeToIssueThreshold *bool `json:"synchronizeToIssueThreshold,omitempty"` // Synchronize To Issue Threshold

	LastModified string `json:"lastModified,omitempty"` // Last Modified
}
type ResponseDevicesUpdateHealthScoreDefinitionForTheGivenIDV1 struct {
	Response *ResponseDevicesUpdateHealthScoreDefinitionForTheGivenIDV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesUpdateHealthScoreDefinitionForTheGivenIDV1Response struct {
	ID string `json:"id,omitempty"` // Id

	Name string `json:"name,omitempty"` // Name

	DisplayName string `json:"displayName,omitempty"` // Display Name

	DeviceFamily string `json:"deviceFamily,omitempty"` // Device Family

	Description string `json:"description,omitempty"` // Description

	IncludeForOverallHealth *bool `json:"includeForOverallHealth,omitempty"` // Include For Overall Health

	DefinitionStatus string `json:"definitionStatus,omitempty"` // Definition Status

	ThresholdValue *float64 `json:"thresholdValue,omitempty"` // Threshold Value

	SynchronizeToIssueThreshold *bool `json:"synchronizeToIssueThreshold,omitempty"` // Synchronize To Issue Threshold

	LastModified string `json:"lastModified,omitempty"` // Last Modified
}
type ResponseDevicesGetAllInterfacesV1 struct {
	Response *[]ResponseDevicesGetAllInterfacesV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetAllInterfacesV1Response struct {
	Addresses *[]ResponseDevicesGetAllInterfacesV1ResponseAddresses `json:"addresses,omitempty"` //

	AdminStatus string `json:"adminStatus,omitempty"` // Admin status as ('UP'/'DOWN')

	ClassName string `json:"className,omitempty"` // Classifies the port as switch port ,loopback interface etc.

	Description string `json:"description,omitempty"` // Description for the Interface

	Name string `json:"name,omitempty"` // Name for the interface

	DeviceID string `json:"deviceId,omitempty"` // Device Id of the device

	Duplex string `json:"duplex,omitempty"` // Interface duplex as AutoNegotiate or FullDuplex

	ID string `json:"id,omitempty"` // ID of the Interface

	IfIndex string `json:"ifIndex,omitempty"` // Interface index

	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id of the Interface

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid of the Interface

	InterfaceType string `json:"interfaceType,omitempty"` // Interface type as Physical or Virtual

	IPv4Address string `json:"ipv4Address,omitempty"` // IPV4 Address of the device

	IPv4Mask string `json:"ipv4Mask,omitempty"` // IPV4 Mask of the device

	IsisSupport string `json:"isisSupport,omitempty"` // Flag for ISIS enabled / disabled

	LastOutgoingPacketTime *float64 `json:"lastOutgoingPacketTime,omitempty"` // Time, in milliseconds since UNIX epoch, when the last packet was sent from this interface

	LastIncomingPacketTime *float64 `json:"lastIncomingPacketTime,omitempty"` // Time, in milliseconds since UNIX epoch, when the last packet was received on this interface

	LastUpdated string `json:"lastUpdated,omitempty"` // Time when the device interface info last got updated

	MacAddress string `json:"macAddress,omitempty"` // MAC address of interface

	MappedPhysicalInterfaceID string `json:"mappedPhysicalInterfaceId,omitempty"` // ID of physical interface mapped with the virtual interface of WLC

	MappedPhysicalInterfaceName string `json:"mappedPhysicalInterfaceName,omitempty"` // Physical interface name mapped with the virtual interface of WLC

	MediaType string `json:"mediaType,omitempty"` // Media Type of the interface

	Mtu string `json:"mtu,omitempty"` // MTU Information of Interface

	NativeVLANID string `json:"nativeVlanId,omitempty"` // Vlan to receive untagged frames on trunk port

	OspfSupport string `json:"ospfSupport,omitempty"` // Flag for OSPF enabled / disabled

	Pid string `json:"pid,omitempty"` // Platform ID of the device

	PortMode string `json:"portMode,omitempty"` // Port mode as access, trunk, routed

	PortName string `json:"portName,omitempty"` // Interface name

	PortType string `json:"portType,omitempty"` // Port type as Ethernet Port / Ethernet SVI / Ethernet Sub Interface

	SerialNo string `json:"serialNo,omitempty"` // Serial number of the device

	Series string `json:"series,omitempty"` // Series of the device

	Speed string `json:"speed,omitempty"` // Speed of the interface

	Status string `json:"status,omitempty"` // Interface status as Down / Up

	VLANID string `json:"vlanId,omitempty"` // Vlan ID of interface

	VoiceVLAN string `json:"voiceVlan,omitempty"` // Vlan information of the interface
}
type ResponseDevicesGetAllInterfacesV1ResponseAddresses struct {
	Address *ResponseDevicesGetAllInterfacesV1ResponseAddressesAddress `json:"address,omitempty"` //

	Type string `json:"type,omitempty"` // Type of the interface. For e.g. IPv4, IPv6 (with unicast, multicast, anycast, etc.)
}
type ResponseDevicesGetAllInterfacesV1ResponseAddressesAddress struct {
	IPAddress *ResponseDevicesGetAllInterfacesV1ResponseAddressesAddressIPAddress `json:"ipAddress,omitempty"` //

	IPMask *ResponseDevicesGetAllInterfacesV1ResponseAddressesAddressIPMask `json:"ipMask,omitempty"` //

	IsInverseMask *bool `json:"isInverseMask,omitempty"` // Inverse Mask of the IP address is enabled or not
}
type ResponseDevicesGetAllInterfacesV1ResponseAddressesAddressIPAddress struct {
	Address string `json:"address,omitempty"` // IP address of the interface
}
type ResponseDevicesGetAllInterfacesV1ResponseAddressesAddressIPMask struct {
	Address string `json:"address,omitempty"` // IP Mask of the interface
}
type ResponseDevicesGetDeviceInterfaceCountForMultipleDevicesV1 struct {
	Response *int `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetInterfaceByIPV1 struct {
	Response *[]ResponseDevicesGetInterfaceByIPV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetInterfaceByIPV1Response struct {
	Addresses *[]ResponseDevicesGetInterfaceByIPV1ResponseAddresses `json:"addresses,omitempty"` //

	AdminStatus string `json:"adminStatus,omitempty"` // Admin status as ('UP'/'DOWN')

	ClassName string `json:"className,omitempty"` // Classifies the port as switch port ,loopback interface etc.

	Description string `json:"description,omitempty"` // Description for the Interface

	Name string `json:"name,omitempty"` // Name for the interface

	DeviceID string `json:"deviceId,omitempty"` // Device Id of the device

	Duplex string `json:"duplex,omitempty"` // Interface duplex as AutoNegotiate or FullDuplex

	ID string `json:"id,omitempty"` // ID of the Interface

	IfIndex string `json:"ifIndex,omitempty"` // Interface index

	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id of the Interface

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid of the Interface

	InterfaceType string `json:"interfaceType,omitempty"` // Interface type as Physical or Virtual

	IPv4Address string `json:"ipv4Address,omitempty"` // IPV4 Address of the device

	IPv4Mask string `json:"ipv4Mask,omitempty"` // IPV4 Mask of the device

	IsisSupport string `json:"isisSupport,omitempty"` // Flag for ISIS enabled / disabled

	LastOutgoingPacketTime *float64 `json:"lastOutgoingPacketTime,omitempty"` // Time, in milliseconds since UNIX epoch, when the last packet was sent from this interface

	LastIncomingPacketTime *float64 `json:"lastIncomingPacketTime,omitempty"` // Time, in milliseconds since UNIX epoch, when the last packet was received on this interface

	LastUpdated string `json:"lastUpdated,omitempty"` // Time when the device interface info last got updated

	MacAddress string `json:"macAddress,omitempty"` // MAC address of interface

	MappedPhysicalInterfaceID string `json:"mappedPhysicalInterfaceId,omitempty"` // ID of physical interface mapped with the virtual interface of WLC

	MappedPhysicalInterfaceName string `json:"mappedPhysicalInterfaceName,omitempty"` // Physical interface name mapped with the virtual interface of WLC

	MediaType string `json:"mediaType,omitempty"` // Media Type of the interface

	Mtu string `json:"mtu,omitempty"` // MTU Information of Interface

	NativeVLANID string `json:"nativeVlanId,omitempty"` // Vlan to receive untagged frames on trunk port

	OspfSupport string `json:"ospfSupport,omitempty"` // Flag for OSPF enabled / disabled

	Pid string `json:"pid,omitempty"` // Platform ID of the device

	PortMode string `json:"portMode,omitempty"` // Port mode as access, trunk, routed

	PortName string `json:"portName,omitempty"` // Interface name

	PortType string `json:"portType,omitempty"` // Port type as Ethernet Port / Ethernet SVI / Ethernet Sub Interface

	SerialNo string `json:"serialNo,omitempty"` // Serial number of the device

	Series string `json:"series,omitempty"` // Series of the device

	Speed string `json:"speed,omitempty"` // Speed of the interface

	Status string `json:"status,omitempty"` // Interface status as Down / Up

	VLANID string `json:"vlanId,omitempty"` // Vlan ID of interface

	VoiceVLAN string `json:"voiceVlan,omitempty"` // Vlan information of the interface
}
type ResponseDevicesGetInterfaceByIPV1ResponseAddresses struct {
	Address *ResponseDevicesGetInterfaceByIPV1ResponseAddressesAddress `json:"address,omitempty"` //

	Type string `json:"type,omitempty"` // Type of the interface. For e.g. IPv4, IPv6 (with unicast, multicast, anycast, etc.)
}
type ResponseDevicesGetInterfaceByIPV1ResponseAddressesAddress struct {
	IPAddress *ResponseDevicesGetInterfaceByIPV1ResponseAddressesAddressIPAddress `json:"ipAddress,omitempty"` //

	IPMask *ResponseDevicesGetInterfaceByIPV1ResponseAddressesAddressIPMask `json:"ipMask,omitempty"` //

	IsInverseMask *bool `json:"isInverseMask,omitempty"` // Inverse Mask of the IP address is enabled or not
}
type ResponseDevicesGetInterfaceByIPV1ResponseAddressesAddressIPAddress struct {
	Address string `json:"address,omitempty"` // IP address of the interface
}
type ResponseDevicesGetInterfaceByIPV1ResponseAddressesAddressIPMask struct {
	Address string `json:"address,omitempty"` // IP Mask of the interface
}
type ResponseDevicesGetIsisInterfacesV1 struct {
	Response *[]ResponseDevicesGetIsisInterfacesV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetIsisInterfacesV1Response struct {
	Addresses *[]ResponseDevicesGetIsisInterfacesV1ResponseAddresses `json:"addresses,omitempty"` //

	AdminStatus string `json:"adminStatus,omitempty"` // Admin status as ('UP'/'DOWN')

	ClassName string `json:"className,omitempty"` // Classifies the port as switch port ,loopback interface etc.

	Description string `json:"description,omitempty"` // Description for the Interface

	Name string `json:"name,omitempty"` // Name for the interface

	DeviceID string `json:"deviceId,omitempty"` // Device Id of the device

	Duplex string `json:"duplex,omitempty"` // Interface duplex as AutoNegotiate or FullDuplex

	ID string `json:"id,omitempty"` // ID of the Interface

	IfIndex string `json:"ifIndex,omitempty"` // Interface index

	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id of the Interface

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid of the Interface

	InterfaceType string `json:"interfaceType,omitempty"` // Interface type as Physical or Virtual

	IPv4Address string `json:"ipv4Address,omitempty"` // IPV4 Address of the device

	IPv4Mask string `json:"ipv4Mask,omitempty"` // IPV4 Mask of the device

	IsisSupport string `json:"isisSupport,omitempty"` // Flag for ISIS enabled / disabled

	LastOutgoingPacketTime *float64 `json:"lastOutgoingPacketTime,omitempty"` // Time, in milliseconds since UNIX epoch, when the last packet was sent from this interface

	LastIncomingPacketTime *float64 `json:"lastIncomingPacketTime,omitempty"` // Time, in milliseconds since UNIX epoch, when the last packet was received on this interface

	LastUpdated string `json:"lastUpdated,omitempty"` // Time when the device interface info last got updated

	MacAddress string `json:"macAddress,omitempty"` // MAC address of interface

	MappedPhysicalInterfaceID string `json:"mappedPhysicalInterfaceId,omitempty"` // ID of physical interface mapped with the virtual interface of WLC

	MappedPhysicalInterfaceName string `json:"mappedPhysicalInterfaceName,omitempty"` // Physical interface name mapped with the virtual interface of WLC

	MediaType string `json:"mediaType,omitempty"` // Media Type of the interface

	Mtu string `json:"mtu,omitempty"` // MTU Information of Interface

	NativeVLANID string `json:"nativeVlanId,omitempty"` // Vlan to receive untagged frames on trunk port

	OspfSupport string `json:"ospfSupport,omitempty"` // Flag for OSPF enabled / disabled

	Pid string `json:"pid,omitempty"` // Platform ID of the device

	PortMode string `json:"portMode,omitempty"` // Port mode as access, trunk, routed

	PortName string `json:"portName,omitempty"` // Interface name

	PortType string `json:"portType,omitempty"` // Port type as Ethernet Port / Ethernet SVI / Ethernet Sub Interface

	SerialNo string `json:"serialNo,omitempty"` // Serial number of the device

	Series string `json:"series,omitempty"` // Series of the device

	Speed string `json:"speed,omitempty"` // Speed of the interface

	Status string `json:"status,omitempty"` // Interface status as Down / Up

	VLANID string `json:"vlanId,omitempty"` // Vlan ID of interface

	VoiceVLAN string `json:"voiceVlan,omitempty"` // Vlan information of the interface
}
type ResponseDevicesGetIsisInterfacesV1ResponseAddresses struct {
	Address *ResponseDevicesGetIsisInterfacesV1ResponseAddressesAddress `json:"address,omitempty"` //

	Type string `json:"type,omitempty"` // Type of the interface. For e.g. IPv4, IPv6 (with unicast, multicast, anycast, etc.)
}
type ResponseDevicesGetIsisInterfacesV1ResponseAddressesAddress struct {
	IPAddress *ResponseDevicesGetIsisInterfacesV1ResponseAddressesAddressIPAddress `json:"ipAddress,omitempty"` //

	IPMask *ResponseDevicesGetIsisInterfacesV1ResponseAddressesAddressIPMask `json:"ipMask,omitempty"` //

	IsInverseMask *bool `json:"isInverseMask,omitempty"` // Inverse Mask of the IP address is enabled or not
}
type ResponseDevicesGetIsisInterfacesV1ResponseAddressesAddressIPAddress struct {
	Address string `json:"address,omitempty"` // IP address of the interface
}
type ResponseDevicesGetIsisInterfacesV1ResponseAddressesAddressIPMask struct {
	Address string `json:"address,omitempty"` // IP Mask of the interface
}
type ResponseDevicesGetInterfaceInfoByIDV1 struct {
	Response *[]ResponseDevicesGetInterfaceInfoByIDV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetInterfaceInfoByIDV1Response struct {
	Addresses *[]ResponseDevicesGetInterfaceInfoByIDV1ResponseAddresses `json:"addresses,omitempty"` //

	AdminStatus string `json:"adminStatus,omitempty"` // Admin status as ('UP'/'DOWN')

	ClassName string `json:"className,omitempty"` // Classifies the port as switch port ,loopback interface etc.

	Description string `json:"description,omitempty"` // Description for the Interface

	Name string `json:"name,omitempty"` // Name for the interface

	DeviceID string `json:"deviceId,omitempty"` // Device Id of the device

	Duplex string `json:"duplex,omitempty"` // Interface duplex as AutoNegotiate or FullDuplex

	ID string `json:"id,omitempty"` // ID of the Interface

	IfIndex string `json:"ifIndex,omitempty"` // Interface index

	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id of the Interface

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid of the Interface

	InterfaceType string `json:"interfaceType,omitempty"` // Interface type as Physical or Virtual

	IPv4Address string `json:"ipv4Address,omitempty"` // IPV4 Address of the device

	IPv4Mask string `json:"ipv4Mask,omitempty"` // IPV4 Mask of the device

	IsisSupport string `json:"isisSupport,omitempty"` // Flag for ISIS enabled / disabled

	LastOutgoingPacketTime *float64 `json:"lastOutgoingPacketTime,omitempty"` // Time, in milliseconds since UNIX epoch, when the last packet was sent from this interface

	LastIncomingPacketTime *float64 `json:"lastIncomingPacketTime,omitempty"` // Time, in milliseconds since UNIX epoch, when the last packet was received on this interface

	LastUpdated string `json:"lastUpdated,omitempty"` // Time when the device interface info last got updated

	MacAddress string `json:"macAddress,omitempty"` // MAC address of interface

	MappedPhysicalInterfaceID string `json:"mappedPhysicalInterfaceId,omitempty"` // ID of physical interface mapped with the virtual interface of WLC

	MappedPhysicalInterfaceName string `json:"mappedPhysicalInterfaceName,omitempty"` // Physical interface name mapped with the virtual interface of WLC

	MediaType string `json:"mediaType,omitempty"` // Media Type of the interface

	Mtu string `json:"mtu,omitempty"` // MTU Information of Interface

	NativeVLANID string `json:"nativeVlanId,omitempty"` // Vlan to receive untagged frames on trunk port

	OspfSupport string `json:"ospfSupport,omitempty"` // Flag for OSPF enabled / disabled

	Pid string `json:"pid,omitempty"` // Platform ID of the device

	PortMode string `json:"portMode,omitempty"` // Port mode as access, trunk, routed

	PortName string `json:"portName,omitempty"` // Interface name

	PortType string `json:"portType,omitempty"` // Port type as Ethernet Port / Ethernet SVI / Ethernet Sub Interface

	SerialNo string `json:"serialNo,omitempty"` // Serial number of the device

	Series string `json:"series,omitempty"` // Series of the device

	Speed string `json:"speed,omitempty"` // Speed of the interface

	Status string `json:"status,omitempty"` // Interface status as Down / Up

	VLANID string `json:"vlanId,omitempty"` // Vlan ID of interface

	VoiceVLAN string `json:"voiceVlan,omitempty"` // Vlan information of the interface
}
type ResponseDevicesGetInterfaceInfoByIDV1ResponseAddresses struct {
	Address *ResponseDevicesGetInterfaceInfoByIDV1ResponseAddressesAddress `json:"address,omitempty"` //

	Type string `json:"type,omitempty"` // Type of the interface. For e.g. IPv4, IPv6 (with unicast, multicast, anycast, etc.)
}
type ResponseDevicesGetInterfaceInfoByIDV1ResponseAddressesAddress struct {
	IPAddress *ResponseDevicesGetInterfaceInfoByIDV1ResponseAddressesAddressIPAddress `json:"ipAddress,omitempty"` //

	IPMask *ResponseDevicesGetInterfaceInfoByIDV1ResponseAddressesAddressIPMask `json:"ipMask,omitempty"` //

	IsInverseMask *bool `json:"isInverseMask,omitempty"` // Inverse Mask of the IP address is enabled or not
}
type ResponseDevicesGetInterfaceInfoByIDV1ResponseAddressesAddressIPAddress struct {
	Address string `json:"address,omitempty"` // IP address of the interface
}
type ResponseDevicesGetInterfaceInfoByIDV1ResponseAddressesAddressIPMask struct {
	Address string `json:"address,omitempty"` // IP Mask of the interface
}
type ResponseDevicesGetDeviceInterfaceCountV1 struct {
	Response *int `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameV1 struct {
	Response *ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameV1Response struct {
	Addresses *[]ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameV1ResponseAddresses `json:"addresses,omitempty"` //

	AdminStatus string `json:"adminStatus,omitempty"` // Admin status as ('UP'/'DOWN')

	ClassName string `json:"className,omitempty"` // Classifies the port as switch port ,loopback interface etc.

	Description string `json:"description,omitempty"` // Description for the Interface

	Name string `json:"name,omitempty"` // Name for the interface

	DeviceID string `json:"deviceId,omitempty"` // Device Id of the device

	Duplex string `json:"duplex,omitempty"` // Interface duplex as AutoNegotiate or FullDuplex

	ID string `json:"id,omitempty"` // ID of the Interface

	IfIndex string `json:"ifIndex,omitempty"` // Interface index

	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id of the Interface

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid of the Interface

	InterfaceType string `json:"interfaceType,omitempty"` // Interface type as Physical or Virtual

	IPv4Address string `json:"ipv4Address,omitempty"` // IPV4 Address of the device

	IPv4Mask string `json:"ipv4Mask,omitempty"` // IPV4 Mask of the device

	IsisSupport string `json:"isisSupport,omitempty"` // Flag for ISIS enabled / disabled

	LastOutgoingPacketTime *float64 `json:"lastOutgoingPacketTime,omitempty"` // Time, in milliseconds since UNIX epoch, when the last packet was sent from this interface

	LastIncomingPacketTime *float64 `json:"lastIncomingPacketTime,omitempty"` // Time, in milliseconds since UNIX epoch, when the last packet was received on this interface

	LastUpdated string `json:"lastUpdated,omitempty"` // Time when the device interface info last got updated

	MacAddress string `json:"macAddress,omitempty"` // MAC address of interface

	MappedPhysicalInterfaceID string `json:"mappedPhysicalInterfaceId,omitempty"` // ID of physical interface mapped with the virtual interface of WLC

	MappedPhysicalInterfaceName string `json:"mappedPhysicalInterfaceName,omitempty"` // Physical interface name mapped with the virtual interface of WLC

	MediaType string `json:"mediaType,omitempty"` // Media Type of the interface

	Mtu string `json:"mtu,omitempty"` // MTU Information of Interface

	NativeVLANID string `json:"nativeVlanId,omitempty"` // Vlan to receive untagged frames on trunk port

	OspfSupport string `json:"ospfSupport,omitempty"` // Flag for OSPF enabled / disabled

	Pid string `json:"pid,omitempty"` // Platform ID of the device

	PortMode string `json:"portMode,omitempty"` // Port mode as access, trunk, routed

	PortName string `json:"portName,omitempty"` // Interface name

	PortType string `json:"portType,omitempty"` // Port type as Ethernet Port / Ethernet SVI / Ethernet Sub Interface

	SerialNo string `json:"serialNo,omitempty"` // Serial number of the device

	Series string `json:"series,omitempty"` // Series of the device

	Speed string `json:"speed,omitempty"` // Speed of the interface

	Status string `json:"status,omitempty"` // Interface status as Down / Up

	VLANID string `json:"vlanId,omitempty"` // Vlan ID of interface

	VoiceVLAN string `json:"voiceVlan,omitempty"` // Vlan information of the interface

	Poweroverethernet string `json:"poweroverethernet,omitempty"` // This is internal attribute.  Not to be used.  Deprecated

	NetworkdeviceID string `json:"networkdevice_id,omitempty"` // This is internal attribute.  Not to be used.  Deprecated

	ManagedComputeElement string `json:"managedComputeElement,omitempty"` // This is internal attribute.  Not to be used.  Deprecated

	ManagedNetworkElement string `json:"managedNetworkElement,omitempty"` // This is internal attribute.  Not to be used.  Deprecated

	ManagedNetworkElementURL string `json:"managedNetworkElementUrl,omitempty"` // This is internal attribute.  Not to be used.  Deprecated

	ManagedComputeElementURL string `json:"managedComputeElementUrl,omitempty"` // This is internal attribute.  Not to be used.  Deprecated
}
type ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameV1ResponseAddresses struct {
	Address *ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameV1ResponseAddressesAddress `json:"address,omitempty"` //

	Type string `json:"type,omitempty"` // Type of the interface. For e.g. IPv4, IPv6 (with unicast, multicast, anycast, etc.)
}
type ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameV1ResponseAddressesAddress struct {
	IPAddress *ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameV1ResponseAddressesAddressIPAddress `json:"ipAddress,omitempty"` //

	IPMask *ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameV1ResponseAddressesAddressIPMask `json:"ipMask,omitempty"` //

	IsInverseMask *bool `json:"isInverseMask,omitempty"` // Inverse Mask of the IP address is enabled or not
}
type ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameV1ResponseAddressesAddressIPAddress struct {
	Address string `json:"address,omitempty"` // IP address of the interface
}
type ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameV1ResponseAddressesAddressIPMask struct {
	Address string `json:"address,omitempty"` // IP Mask of the interface
}
type ResponseDevicesGetDeviceInterfacesBySpecifiedRangeV1 struct {
	Response *[]ResponseDevicesGetDeviceInterfacesBySpecifiedRangeV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetDeviceInterfacesBySpecifiedRangeV1Response struct {
	Addresses *[]ResponseDevicesGetDeviceInterfacesBySpecifiedRangeV1ResponseAddresses `json:"addresses,omitempty"` //

	AdminStatus string `json:"adminStatus,omitempty"` // Admin status as ('UP'/'DOWN')

	ClassName string `json:"className,omitempty"` // Classifies the port as switch port ,loopback interface etc.

	Description string `json:"description,omitempty"` // Description for the Interface

	Name string `json:"name,omitempty"` // Name for the interface

	DeviceID string `json:"deviceId,omitempty"` // Device Id of the device

	Duplex string `json:"duplex,omitempty"` // Interface duplex as AutoNegotiate or FullDuplex

	ID string `json:"id,omitempty"` // ID of the Interface

	IfIndex string `json:"ifIndex,omitempty"` // Interface index

	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id of the Interface

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid of the Interface

	InterfaceType string `json:"interfaceType,omitempty"` // Interface type as Physical or Virtual

	IPv4Address string `json:"ipv4Address,omitempty"` // IPV4 Address of the device

	IPv4Mask string `json:"ipv4Mask,omitempty"` // IPV4 Mask of the device

	IsisSupport string `json:"isisSupport,omitempty"` // Flag for ISIS enabled / disabled

	LastOutgoingPacketTime *float64 `json:"lastOutgoingPacketTime,omitempty"` // Time, in milliseconds since UNIX epoch, when the last packet was sent from this interface

	LastIncomingPacketTime *float64 `json:"lastIncomingPacketTime,omitempty"` // Time, in milliseconds since UNIX epoch, when the last packet was received on this interface

	LastUpdated string `json:"lastUpdated,omitempty"` // Time when the device interface info last got updated

	MacAddress string `json:"macAddress,omitempty"` // MAC address of interface

	MappedPhysicalInterfaceID string `json:"mappedPhysicalInterfaceId,omitempty"` // ID of physical interface mapped with the virtual interface of WLC

	MappedPhysicalInterfaceName string `json:"mappedPhysicalInterfaceName,omitempty"` // Physical interface name mapped with the virtual interface of WLC

	MediaType string `json:"mediaType,omitempty"` // Media Type of the interface

	Mtu string `json:"mtu,omitempty"` // MTU Information of Interface

	NativeVLANID string `json:"nativeVlanId,omitempty"` // Vlan to receive untagged frames on trunk port

	OspfSupport string `json:"ospfSupport,omitempty"` // Flag for OSPF enabled / disabled

	Pid string `json:"pid,omitempty"` // Platform ID of the device

	PortMode string `json:"portMode,omitempty"` // Port mode as access, trunk, routed

	PortName string `json:"portName,omitempty"` // Interface name

	PortType string `json:"portType,omitempty"` // Port type as Ethernet Port / Ethernet SVI / Ethernet Sub Interface

	SerialNo string `json:"serialNo,omitempty"` // Serial number of the device

	Series string `json:"series,omitempty"` // Series of the device

	Speed string `json:"speed,omitempty"` // Speed of the interface

	Status string `json:"status,omitempty"` // Interface status as Down / Up

	VLANID string `json:"vlanId,omitempty"` // Vlan ID of interface

	VoiceVLAN string `json:"voiceVlan,omitempty"` // Vlan information of the interface
}
type ResponseDevicesGetDeviceInterfacesBySpecifiedRangeV1ResponseAddresses struct {
	Address *ResponseDevicesGetDeviceInterfacesBySpecifiedRangeV1ResponseAddressesAddress `json:"address,omitempty"` //

	Type string `json:"type,omitempty"` // Type of the interface. For e.g. IPv4, IPv6 (with unicast, multicast, anycast, etc.)
}
type ResponseDevicesGetDeviceInterfacesBySpecifiedRangeV1ResponseAddressesAddress struct {
	IPAddress *ResponseDevicesGetDeviceInterfacesBySpecifiedRangeV1ResponseAddressesAddressIPAddress `json:"ipAddress,omitempty"` //

	IPMask *ResponseDevicesGetDeviceInterfacesBySpecifiedRangeV1ResponseAddressesAddressIPMask `json:"ipMask,omitempty"` //

	IsInverseMask *bool `json:"isInverseMask,omitempty"` // Inverse Mask of the IP address is enabled or not
}
type ResponseDevicesGetDeviceInterfacesBySpecifiedRangeV1ResponseAddressesAddressIPAddress struct {
	Address string `json:"address,omitempty"` // IP address of the interface
}
type ResponseDevicesGetDeviceInterfacesBySpecifiedRangeV1ResponseAddressesAddressIPMask struct {
	Address string `json:"address,omitempty"` // IP Mask of the interface
}
type ResponseDevicesGetOspfInterfacesV1 struct {
	Response *[]ResponseDevicesGetOspfInterfacesV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetOspfInterfacesV1Response struct {
	Addresses *[]ResponseDevicesGetOspfInterfacesV1ResponseAddresses `json:"addresses,omitempty"` //

	AdminStatus string `json:"adminStatus,omitempty"` // Admin status as ('UP'/'DOWN')

	ClassName string `json:"className,omitempty"` // Classifies the port as switch port ,loopback interface etc.

	Description string `json:"description,omitempty"` // Description for the Interface

	Name string `json:"name,omitempty"` // Name for the interface

	DeviceID string `json:"deviceId,omitempty"` // Device Id of the device

	Duplex string `json:"duplex,omitempty"` // Interface duplex as AutoNegotiate or FullDuplex

	ID string `json:"id,omitempty"` // ID of the Interface

	IfIndex string `json:"ifIndex,omitempty"` // Interface index

	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id of the Interface

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid of the Interface

	InterfaceType string `json:"interfaceType,omitempty"` // Interface type as Physical or Virtual

	IPv4Address string `json:"ipv4Address,omitempty"` // IPV4 Address of the device

	IPv4Mask string `json:"ipv4Mask,omitempty"` // IPV4 Mask of the device

	IsisSupport string `json:"isisSupport,omitempty"` // Flag for ISIS enabled / disabled

	LastOutgoingPacketTime *float64 `json:"lastOutgoingPacketTime,omitempty"` // Time, in milliseconds since UNIX epoch, when the last packet was sent from this interface

	LastIncomingPacketTime *float64 `json:"lastIncomingPacketTime,omitempty"` // Time, in milliseconds since UNIX epoch, when the last packet was received on this interface

	LastUpdated string `json:"lastUpdated,omitempty"` // Time when the device interface info last got updated

	MacAddress string `json:"macAddress,omitempty"` // MAC address of interface

	MappedPhysicalInterfaceID string `json:"mappedPhysicalInterfaceId,omitempty"` // ID of physical interface mapped with the virtual interface of WLC

	MappedPhysicalInterfaceName string `json:"mappedPhysicalInterfaceName,omitempty"` // Physical interface name mapped with the virtual interface of WLC

	MediaType string `json:"mediaType,omitempty"` // Media Type of the interface

	Mtu string `json:"mtu,omitempty"` // MTU Information of Interface

	NativeVLANID string `json:"nativeVlanId,omitempty"` // Vlan to receive untagged frames on trunk port

	OspfSupport string `json:"ospfSupport,omitempty"` // Flag for OSPF enabled / disabled

	Pid string `json:"pid,omitempty"` // Platform ID of the device

	PortMode string `json:"portMode,omitempty"` // Port mode as access, trunk, routed

	PortName string `json:"portName,omitempty"` // Interface name

	PortType string `json:"portType,omitempty"` // Port type as Ethernet Port / Ethernet SVI / Ethernet Sub Interface

	SerialNo string `json:"serialNo,omitempty"` // Serial number of the device

	Series string `json:"series,omitempty"` // Series of the device

	Speed string `json:"speed,omitempty"` // Speed of the interface

	Status string `json:"status,omitempty"` // Interface status as Down / Up

	VLANID string `json:"vlanId,omitempty"` // Vlan ID of interface

	VoiceVLAN string `json:"voiceVlan,omitempty"` // Vlan information of the interface
}
type ResponseDevicesGetOspfInterfacesV1ResponseAddresses struct {
	Address *ResponseDevicesGetOspfInterfacesV1ResponseAddressesAddress `json:"address,omitempty"` //

	Type string `json:"type,omitempty"` // Type of the interface. For e.g. IPv4, IPv6 (with unicast, multicast, anycast, etc.)
}
type ResponseDevicesGetOspfInterfacesV1ResponseAddressesAddress struct {
	IPAddress *ResponseDevicesGetOspfInterfacesV1ResponseAddressesAddressIPAddress `json:"ipAddress,omitempty"` //

	IPMask *ResponseDevicesGetOspfInterfacesV1ResponseAddressesAddressIPMask `json:"ipMask,omitempty"` //

	IsInverseMask *bool `json:"isInverseMask,omitempty"` // Inverse Mask of the IP address is enabled or not
}
type ResponseDevicesGetOspfInterfacesV1ResponseAddressesAddressIPAddress struct {
	Address string `json:"address,omitempty"` // IP address of the interface
}
type ResponseDevicesGetOspfInterfacesV1ResponseAddressesAddressIPMask struct {
	Address string `json:"address,omitempty"` // IP Mask of the interface
}
type ResponseDevicesGetInterfaceByIDV1 struct {
	Response *ResponseDevicesGetInterfaceByIDV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetInterfaceByIDV1Response struct {
	Addresses *[]ResponseDevicesGetInterfaceByIDV1ResponseAddresses `json:"addresses,omitempty"` //

	AdminStatus string `json:"adminStatus,omitempty"` // Admin status as ('UP'/'DOWN')

	ClassName string `json:"className,omitempty"` // Classifies the port as switch port ,loopback interface etc.

	Description string `json:"description,omitempty"` // Description for the Interface

	Name string `json:"name,omitempty"` // Name for the interface

	DeviceID string `json:"deviceId,omitempty"` // Device Id of the device

	Duplex string `json:"duplex,omitempty"` // Interface duplex as AutoNegotiate or FullDuplex

	ID string `json:"id,omitempty"` // ID of the Interface

	IfIndex string `json:"ifIndex,omitempty"` // Interface index

	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id of the Interface

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid of the Interface

	InterfaceType string `json:"interfaceType,omitempty"` // Interface type as Physical or Virtual

	IPv4Address string `json:"ipv4Address,omitempty"` // IPV4 Address of the device

	IPv4Mask string `json:"ipv4Mask,omitempty"` // IPV4 Mask of the device

	IsisSupport string `json:"isisSupport,omitempty"` // Flag for ISIS enabled / disabled

	LastOutgoingPacketTime *float64 `json:"lastOutgoingPacketTime,omitempty"` // Time, in milliseconds since UNIX epoch, when the last packet was sent from this interface

	LastIncomingPacketTime *float64 `json:"lastIncomingPacketTime,omitempty"` // Time, in milliseconds since UNIX epoch, when the last packet was received on this interface

	LastUpdated string `json:"lastUpdated,omitempty"` // Time when the device interface info last got updated

	MacAddress string `json:"macAddress,omitempty"` // MAC address of interface

	MappedPhysicalInterfaceID string `json:"mappedPhysicalInterfaceId,omitempty"` // ID of physical interface mapped with the virtual interface of WLC

	MappedPhysicalInterfaceName string `json:"mappedPhysicalInterfaceName,omitempty"` // Physical interface name mapped with the virtual interface of WLC

	MediaType string `json:"mediaType,omitempty"` // Media Type of the interface

	Mtu string `json:"mtu,omitempty"` // MTU Information of Interface

	NativeVLANID string `json:"nativeVlanId,omitempty"` // Vlan to receive untagged frames on trunk port

	OspfSupport string `json:"ospfSupport,omitempty"` // Flag for OSPF enabled / disabled

	Pid string `json:"pid,omitempty"` // Platform ID of the device

	PortMode string `json:"portMode,omitempty"` // Port mode as access, trunk, routed

	PortName string `json:"portName,omitempty"` // Interface name

	PortType string `json:"portType,omitempty"` // Port type as Ethernet Port / Ethernet SVI / Ethernet Sub Interface

	SerialNo string `json:"serialNo,omitempty"` // Serial number of the device

	Series string `json:"series,omitempty"` // Series of the device

	Speed string `json:"speed,omitempty"` // Speed of the interface

	Status string `json:"status,omitempty"` // Interface status as Down / Up

	VLANID string `json:"vlanId,omitempty"` // Vlan ID of interface

	VoiceVLAN string `json:"voiceVlan,omitempty"` // Vlan information of the interface
}
type ResponseDevicesGetInterfaceByIDV1ResponseAddresses struct {
	Address *ResponseDevicesGetInterfaceByIDV1ResponseAddressesAddress `json:"address,omitempty"` //

	Type string `json:"type,omitempty"` // Type of the interface. For e.g. IPv4, IPv6 (with unicast, multicast, anycast, etc.)
}
type ResponseDevicesGetInterfaceByIDV1ResponseAddressesAddress struct {
	IPAddress *ResponseDevicesGetInterfaceByIDV1ResponseAddressesAddressIPAddress `json:"ipAddress,omitempty"` //

	IPMask *ResponseDevicesGetInterfaceByIDV1ResponseAddressesAddressIPMask `json:"ipMask,omitempty"` //

	IsInverseMask *bool `json:"isInverseMask,omitempty"` // Inverse Mask of the IP address is enabled or not
}
type ResponseDevicesGetInterfaceByIDV1ResponseAddressesAddressIPAddress struct {
	Address string `json:"address,omitempty"` // IP address of the interface
}
type ResponseDevicesGetInterfaceByIDV1ResponseAddressesAddressIPMask struct {
	Address string `json:"address,omitempty"` // IP Mask of the interface
}
type ResponseDevicesUpdateInterfaceDetailsV1 struct {
	Response *ResponseDevicesUpdateInterfaceDetailsV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesUpdateInterfaceDetailsV1Response struct {
	TaskID string `json:"taskId,omitempty"` //

	URL string `json:"url,omitempty"` //
}
type ResponseDevicesLegitOperationsForInterfaceV1 struct {
	Response *ResponseDevicesLegitOperationsForInterfaceV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesLegitOperationsForInterfaceV1Response struct {
	InterfaceUUID string `json:"interfaceUuid,omitempty"` // Id of the Interface

	Properties *[]ResponseDevicesLegitOperationsForInterfaceV1ResponseProperties `json:"properties,omitempty"` //

	Operations *[]ResponseDevicesLegitOperationsForInterfaceV1ResponseOperations `json:"operations,omitempty"` //
}
type ResponseDevicesLegitOperationsForInterfaceV1ResponseProperties struct {
	Name string `json:"name,omitempty"` // Name of the Property

	Applicable string `json:"applicable,omitempty"` // Checks if property is applicable to interface

	FailureReason string `json:"failureReason,omitempty"` // Failure reason of the Property
}
type ResponseDevicesLegitOperationsForInterfaceV1ResponseOperations struct {
	Name string `json:"name,omitempty"` // Name of the Operation

	Applicable string `json:"applicable,omitempty"` // Checks if operation is applicable to interface

	FailureReason string `json:"failureReason,omitempty"` // Failure reason of the Operation
}
type ResponseDevicesClearMacAddressTableV1 struct {
	Response *ResponseDevicesClearMacAddressTableV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesClearMacAddressTableV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	URL string `json:"url,omitempty"` // Url
}
type ResponseDevicesGetDeviceListV1 struct {
	Response *[]ResponseDevicesGetDeviceListV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesGetDeviceListV1Response struct {
	ReachabilityFailureReason string `json:"reachabilityFailureReason,omitempty"` // Failure reason for unreachable devices

	ReachabilityStatus string `json:"reachabilityStatus,omitempty"` // Device reachability status as Reachable / Unreachable

	Series string `json:"series,omitempty"` // Device series

	SNMPContact string `json:"snmpContact,omitempty"` // SNMP contact on device

	SNMPLocation string `json:"snmpLocation,omitempty"` // SNMP location on device

	TagCount string `json:"tagCount,omitempty"` // Number of tags associated with the device

	TunnelUDPPort string `json:"tunnelUdpPort,omitempty"` // Mobility protocol port is stored as tunneludpport for WLC

	UptimeSeconds *float64 `json:"uptimeSeconds,omitempty"` // Uptime in Seconds

	WaasDeviceMode string `json:"waasDeviceMode,omitempty"` // WAAS device mode

	SerialNumber string `json:"serialNumber,omitempty"` // Serial number of device

	LastUpdateTime *float64 `json:"lastUpdateTime,omitempty"` // Time in epoch when the network device info last got updated

	MacAddress string `json:"macAddress,omitempty"` // MAC address of device

	UpTime string `json:"upTime,omitempty"` // Time that shows for how long the device has been up

	DeviceSupportLevel string `json:"deviceSupportLevel,omitempty"` // Support level of the device

	Hostname string `json:"hostname,omitempty"` // Device name

	Type string `json:"type,omitempty"` // Type of device as switch, router, wireless lan controller, accesspoints

	MemorySize string `json:"memorySize,omitempty"` // Processor memory size

	Family string `json:"family,omitempty"` // Family of device as switch, router, wireless lan controller, accesspoints

	ErrorCode string `json:"errorCode,omitempty"` // Inventory status error code

	SoftwareType string `json:"softwareType,omitempty"` // Software type on the device

	SoftwareVersion string `json:"softwareVersion,omitempty"` // Software version on the device

	Description string `json:"description,omitempty"` // System description

	RoleSource string `json:"roleSource,omitempty"` // Role source as manual / auto

	Location string `json:"location,omitempty"` // [Deprecated] Location ID that is associated with the device

	Role string `json:"role,omitempty"` // Role of device as access, distribution, border router, core

	CollectionInterval string `json:"collectionInterval,omitempty"` // Re sync Interval of the device

	InventoryStatusDetail string `json:"inventoryStatusDetail,omitempty"` // Status detail of inventory sync

	ApEthernetMacAddress string `json:"apEthernetMacAddress,omitempty"` // AccessPoint Ethernet MacAddress of AP device

	ApManagerInterfaceIP string `json:"apManagerInterfaceIp,omitempty"` // IP address of WLC on AP manager interface

	AssociatedWlcIP string `json:"associatedWlcIp,omitempty"` // Associated Wlc Ip address of the AP device

	BootDateTime string `json:"bootDateTime,omitempty"` // Device boot time

	CollectionStatus string `json:"collectionStatus,omitempty"` // Collection status as Synchronizing, Could not synchronize, Not manageable, Managed, Partial Collection Failure, Incomplete, Unreachable, Wrong credential, Reachable, In Progress

	ErrorDescription string `json:"errorDescription,omitempty"` // Inventory status description

	InterfaceCount string `json:"interfaceCount,omitempty"` // Number of interfaces on the device

	LastUpdated string `json:"lastUpdated,omitempty"` // Time when the network device info last got updated

	LineCardCount string `json:"lineCardCount,omitempty"` // Number of linecards on the device

	LineCardID string `json:"lineCardId,omitempty"` // IDs of linecards of the device

	LocationName string `json:"locationName,omitempty"` // [Deprecated] Name of the associated location

	ManagedAtleastOnce *bool `json:"managedAtleastOnce,omitempty"` // Indicates if device went into Managed state atleast once

	ManagementIPAddress string `json:"managementIpAddress,omitempty"` // IP address of the device

	PlatformID string `json:"platformId,omitempty"` // Platform ID of device

	ManagementState string `json:"managementState,omitempty"` // Represents the current management state of the network element: managed, unmanaged, under maintenance, and so on.

	PendingSyncRequestsCount string `json:"pendingSyncRequestsCount,omitempty"` // Count of pending sync requests , if any

	ReasonsForDeviceResync string `json:"reasonsForDeviceResync,omitempty"` // Reason for last/ongoing sync

	ReasonsForPendingSyncRequests string `json:"reasonsForPendingSyncRequests,omitempty"` // Reasons for pending sync requests , if any

	SyncRequestedByApp string `json:"syncRequestedByApp,omitempty"` // Applications which requested for the resync of network device

	LastManagedResyncReasons string `json:"lastManagedResyncReasons,omitempty"` // Reasons for last successful sync

	DNSResolvedManagementAddress string `json:"dnsResolvedManagementAddress,omitempty"` // Specifies the resolved ip address of dns name

	LastDeviceResyncStartTime string `json:"lastDeviceResyncStartTime,omitempty"` // Start time for last/ongoing sync

	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id of the device

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid of the device

	ID string `json:"id,omitempty"` // Instance Uuid of the device
}
type ResponseDevicesAddDeviceKnowYourNetworkV1 struct {
	Response *ResponseDevicesAddDeviceKnowYourNetworkV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesAddDeviceKnowYourNetworkV1Response struct {
	TaskID string `json:"taskId,omitempty"` //

	URL string `json:"url,omitempty"` //
}
type ResponseDevicesUpdateDeviceDetailsV1 struct {
	Response *ResponseDevicesUpdateDeviceDetailsV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesUpdateDeviceDetailsV1Response struct {
	TaskID string `json:"taskId,omitempty"` //

	URL string `json:"url,omitempty"` //
}
type ResponseDevicesGetDeviceValuesThatMatchFullyOrPartiallyAnAttributeV1 struct {
	Response []string `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesUpdateDeviceRoleV1 struct {
	Response *ResponseDevicesUpdateDeviceRoleV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesUpdateDeviceRoleV1Response struct {
	TaskID string `json:"taskId,omitempty"` //

	URL string `json:"url,omitempty"` //
}
type ResponseDevicesGetPollingIntervalForAllDevicesV1 struct {
	Response *int `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetDeviceConfigForAllDevicesV1 struct {
	Response *[]ResponseDevicesGetDeviceConfigForAllDevicesV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetDeviceConfigForAllDevicesV1Response struct {
	AttributeInfo *ResponseDevicesGetDeviceConfigForAllDevicesV1ResponseAttributeInfo `json:"attributeInfo,omitempty"` //

	CdpNeighbors string `json:"cdpNeighbors,omitempty"` //

	HealthMonitor string `json:"healthMonitor,omitempty"` //

	ID string `json:"id,omitempty"` //

	IntfDescription string `json:"intfDescription,omitempty"` //

	Inventory string `json:"inventory,omitempty"` //

	IPIntfBrief string `json:"ipIntfBrief,omitempty"` //

	MacAddressTable string `json:"macAddressTable,omitempty"` //

	RunningConfig string `json:"runningConfig,omitempty"` //

	SNMP string `json:"snmp,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetDeviceConfigForAllDevicesV1ResponseAttributeInfo interface{}
type ResponseDevicesGetDeviceConfigCountV1 struct {
	Response *int `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetDeviceCountKnowYourNetworkV1 struct {
	Response *int `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesExportDeviceListV1 struct {
	Response *ResponseDevicesExportDeviceListV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesExportDeviceListV1Response struct {
	TaskID string `json:"taskId,omitempty"` //

	URL string `json:"url,omitempty"` //
}
type ResponseDevicesGetFunctionalCapabilityForDevicesV1 struct {
	Response *[]ResponseDevicesGetFunctionalCapabilityForDevicesV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetFunctionalCapabilityForDevicesV1Response struct {
	AttributeInfo *ResponseDevicesGetFunctionalCapabilityForDevicesV1ResponseAttributeInfo `json:"attributeInfo,omitempty"` // Deprecated

	DeviceID string `json:"deviceId,omitempty"` // Device Id of the device

	FunctionalCapability *[]ResponseDevicesGetFunctionalCapabilityForDevicesV1ResponseFunctionalCapability `json:"functionalCapability,omitempty"` //

	ID string `json:"id,omitempty"` // Deprecated
}
type ResponseDevicesGetFunctionalCapabilityForDevicesV1ResponseAttributeInfo interface{}
type ResponseDevicesGetFunctionalCapabilityForDevicesV1ResponseFunctionalCapability struct {
	AttributeInfo *ResponseDevicesGetFunctionalCapabilityForDevicesV1ResponseFunctionalCapabilityAttributeInfo `json:"attributeInfo,omitempty"` // Deprecated

	FunctionDetails *[]ResponseDevicesGetFunctionalCapabilityForDevicesV1ResponseFunctionalCapabilityFunctionDetails `json:"functionDetails,omitempty"` //

	FunctionName string `json:"functionName,omitempty"` // Name of the function

	FunctionOpState string `json:"functionOpState,omitempty"` // Operational state of the function

	ID string `json:"id,omitempty"` // Id of the function
}
type ResponseDevicesGetFunctionalCapabilityForDevicesV1ResponseFunctionalCapabilityAttributeInfo interface{}
type ResponseDevicesGetFunctionalCapabilityForDevicesV1ResponseFunctionalCapabilityFunctionDetails struct {
	AttributeInfo *ResponseDevicesGetFunctionalCapabilityForDevicesV1ResponseFunctionalCapabilityFunctionDetailsAttributeInfo `json:"attributeInfo,omitempty"` // Deprecated

	ID string `json:"id,omitempty"` // Deprecated

	PropertyName string `json:"propertyName,omitempty"` // Property Name of the function

	StringValue string `json:"stringValue,omitempty"` // Value for the property
}
type ResponseDevicesGetFunctionalCapabilityForDevicesV1ResponseFunctionalCapabilityFunctionDetailsAttributeInfo interface{}
type ResponseDevicesGetFunctionalCapabilityByIDV1 struct {
	Response *ResponseDevicesGetFunctionalCapabilityByIDV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetFunctionalCapabilityByIDV1Response struct {
	AttributeInfo *ResponseDevicesGetFunctionalCapabilityByIDV1ResponseAttributeInfo `json:"attributeInfo,omitempty"` // Deprecated

	FunctionDetails *[]ResponseDevicesGetFunctionalCapabilityByIDV1ResponseFunctionDetails `json:"functionDetails,omitempty"` //

	FunctionName string `json:"functionName,omitempty"` // Name of the function

	FunctionOpState string `json:"functionOpState,omitempty"` // Operational state of the function

	ID string `json:"id,omitempty"` // Id of the function
}
type ResponseDevicesGetFunctionalCapabilityByIDV1ResponseAttributeInfo interface{}
type ResponseDevicesGetFunctionalCapabilityByIDV1ResponseFunctionDetails struct {
	AttributeInfo *ResponseDevicesGetFunctionalCapabilityByIDV1ResponseFunctionDetailsAttributeInfo `json:"attributeInfo,omitempty"` // Deprecated

	ID string `json:"id,omitempty"` // Deprecated

	PropertyName string `json:"propertyName,omitempty"` // Property Name of the function

	StringValue string `json:"stringValue,omitempty"` // Value for the property
}
type ResponseDevicesGetFunctionalCapabilityByIDV1ResponseFunctionDetailsAttributeInfo interface{}
type ResponseDevicesInventoryInsightDeviceLinkMismatchAPIV1 struct {
	Response *[]ResponseDevicesInventoryInsightDeviceLinkMismatchAPIV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Api version
}
type ResponseDevicesInventoryInsightDeviceLinkMismatchAPIV1Response struct {
	EndPortAllowedVLANIDs string `json:"endPortAllowedVlanIds,omitempty"` // End port allowed vlan ids

	EndPortNativeVLANID string `json:"endPortNativeVlanId,omitempty"` // End port native vlan id

	StartPortAllowedVLANIDs string `json:"startPortAllowedVlanIds,omitempty"` // Start port allowed vlan ids

	StartPortNativeVLANID string `json:"startPortNativeVlanId,omitempty"` // Start port native vlan id

	LinkStatus string `json:"linkStatus,omitempty"` // Link status

	EndDeviceHostName string `json:"endDeviceHostName,omitempty"` // End device hostname

	EndDeviceID string `json:"endDeviceId,omitempty"` // End device id

	EndDeviceIPAddress string `json:"endDeviceIpAddress,omitempty"` // End device ip address

	EndPortAddress string `json:"endPortAddress,omitempty"` // End port address

	EndPortDuplex string `json:"endPortDuplex,omitempty"` // End port duplex

	EndPortID string `json:"endPortId,omitempty"` // End port id

	EndPortMask string `json:"endPortMask,omitempty"` // End port mask

	EndPortName string `json:"endPortName,omitempty"` // End port name

	EndPortPepID string `json:"endPortPepId,omitempty"` // End port pep id

	EndPortSpeed string `json:"endPortSpeed,omitempty"` // End port speed

	StartDeviceHostName string `json:"startDeviceHostName,omitempty"` // Start device hostname

	StartDeviceID string `json:"startDeviceId,omitempty"` // Start device id

	StartDeviceIPAddress string `json:"startDeviceIpAddress,omitempty"` // Start device ip address

	StartPortAddress string `json:"startPortAddress,omitempty"` // Start port address

	StartPortDuplex string `json:"startPortDuplex,omitempty"` // Start port duplex

	StartPortID string `json:"startPortId,omitempty"` // Start port id

	StartPortMask string `json:"startPortMask,omitempty"` // Start port mask

	StartPortName string `json:"startPortName,omitempty"` // Start port name

	StartPortPepID string `json:"startPortPepId,omitempty"` // Start port pep id

	StartPortSpeed string `json:"startPortSpeed,omitempty"` // Start port speed

	LastUpdated string `json:"lastUpdated,omitempty"` // Last updated

	NumUpdates *float64 `json:"numUpdates,omitempty"` // Number updates

	AvgUpdateFrequency *float64 `json:"avgUpdateFrequency,omitempty"` // Average update frequency

	Type string `json:"type,omitempty"` // Type

	InstanceUUID string `json:"instanceUuid,omitempty"` // Unique instance id

	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance tenant id
}
type ResponseDevicesGetNetworkDeviceByIPV1 struct {
	Response *ResponseDevicesGetNetworkDeviceByIPV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetNetworkDeviceByIPV1Response struct {
	ApManagerInterfaceIP string `json:"apManagerInterfaceIp,omitempty"` // IP address of WLC on AP manager interface

	AssociatedWlcIP string `json:"associatedWlcIp,omitempty"` // Associated Wlc Ip address of the AP device

	BootDateTime string `json:"bootDateTime,omitempty"` // Device boot time

	CollectionInterval string `json:"collectionInterval,omitempty"` // Re sync Interval of the device

	CollectionStatus string `json:"collectionStatus,omitempty"` // Collection status as Synchronizing, Could not synchronize, Not manageable, Managed, Partial Collection Failure, Incomplete, Unreachable, Wrong credential, Reachable, In Progress

	ErrorCode string `json:"errorCode,omitempty"` // Inventory status error code

	ErrorDescription string `json:"errorDescription,omitempty"` // Inventory status description

	Family string `json:"family,omitempty"` // Family of device as switch, router, wireless lan controller, accesspoints

	Hostname string `json:"hostname,omitempty"` // Device name

	ID string `json:"id,omitempty"` // Instance Uuid of the device

	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id of the device

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid of the device

	InterfaceCount string `json:"interfaceCount,omitempty"` // Number of interfaces on the device

	InventoryStatusDetail string `json:"inventoryStatusDetail,omitempty"` // Status detail of inventory sync

	LastUpdateTime *float64 `json:"lastUpdateTime,omitempty"` // Time in epoch when the network device info last got updated

	LastUpdated string `json:"lastUpdated,omitempty"` // Time when the network device info last got updated

	LineCardCount string `json:"lineCardCount,omitempty"` // Number of linecards on the device

	LineCardID string `json:"lineCardId,omitempty"` // IDs of linecards of the device

	Location string `json:"location,omitempty"` // [Deprecated] Location ID that is associated with the device

	LocationName string `json:"locationName,omitempty"` // [Deprecated] Name of the associated location

	MacAddress string `json:"macAddress,omitempty"` // MAC address of device

	ManagementIPAddress string `json:"managementIpAddress,omitempty"` // IP address of the device

	MemorySize string `json:"memorySize,omitempty"` // Processor memory size

	PlatformID string `json:"platformId,omitempty"` // Platform ID of device

	ReachabilityFailureReason string `json:"reachabilityFailureReason,omitempty"` // Failure reason for unreachable devices

	ReachabilityStatus string `json:"reachabilityStatus,omitempty"` // Device reachability status as Reachable / Unreachable

	Role string `json:"role,omitempty"` // Role of device as access, distribution, border router, core

	RoleSource string `json:"roleSource,omitempty"` // Role source as manual / auto

	SerialNumber string `json:"serialNumber,omitempty"` // Serial number of device

	Series string `json:"series,omitempty"` // Device series

	SNMPContact string `json:"snmpContact,omitempty"` // SNMP contact on device

	SNMPLocation string `json:"snmpLocation,omitempty"` // SNMP location on device

	SoftwareType string `json:"softwareType,omitempty"` // Software type on the device

	SoftwareVersion string `json:"softwareVersion,omitempty"` // Software version on the device

	TagCount string `json:"tagCount,omitempty"` // Number of tags associated with the device

	TunnelUDPPort string `json:"tunnelUdpPort,omitempty"` // Mobility protocol port is stored as tunneludpport for WLC

	Type string `json:"type,omitempty"` // Type of device as switch, router, wireless lan controller, accesspoints

	UpTime string `json:"upTime,omitempty"` // Time that shows for how long the device has been up

	WaasDeviceMode string `json:"waasDeviceMode,omitempty"` // WAAS device mode

	DNSResolvedManagementAddress string `json:"dnsResolvedManagementAddress,omitempty"` // Specifies the resolved ip address of dns name

	ApEthernetMacAddress string `json:"apEthernetMacAddress,omitempty"` // AccessPoint Ethernet MacAddress of AP device

	Vendor string `json:"vendor,omitempty"` // Vendor details

	ReasonsForPendingSyncRequests string `json:"reasonsForPendingSyncRequests,omitempty"` // Reasons for pending sync requests , if any

	PendingSyncRequestsCount string `json:"pendingSyncRequestsCount,omitempty"` // Count of pending sync requests , if any

	ReasonsForDeviceResync string `json:"reasonsForDeviceResync,omitempty"` // Reason for last/ongoing sync

	LastDeviceResyncStartTime string `json:"lastDeviceResyncStartTime,omitempty"` // Start time for last/ongoing sync

	UptimeSeconds *float64 `json:"uptimeSeconds,omitempty"` // Uptime in Seconds

	ManagedAtleastOnce *bool `json:"managedAtleastOnce,omitempty"` // Indicates if device went into Managed state atleast once

	DeviceSupportLevel string `json:"deviceSupportLevel,omitempty"` // Support level of the device

	ManagementState string `json:"managementState,omitempty"` // Represents the current management state of the network element: managed, unmanaged, under maintenance, and so on.

	Description string `json:"description,omitempty"` // System description
}
type ResponseDevicesGetModulesV1 struct {
	Response *[]ResponseDevicesGetModulesV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetModulesV1Response struct {
	AssemblyNumber string `json:"assemblyNumber,omitempty"` // Assembly Number of the module

	AssemblyRevision string `json:"assemblyRevision,omitempty"` // Assembly Revision of the module

	AttributeInfo *ResponseDevicesGetModulesV1ResponseAttributeInfo `json:"attributeInfo,omitempty"` // Deprecated

	ContainmentEntity string `json:"containmentEntity,omitempty"` // Containment Entity of the module

	Description string `json:"description,omitempty"` // Description of the module

	EntityPhysicalIndex string `json:"entityPhysicalIndex,omitempty"` // Entity Physical Index of the module

	ID string `json:"id,omitempty"` // ID of the module

	IsFieldReplaceable string `json:"isFieldReplaceable,omitempty"` // To mention if field is replaceable

	IsReportingAlarmsAllowed string `json:"isReportingAlarmsAllowed,omitempty"` // To mention if reporting alarms are allowed

	Manufacturer string `json:"manufacturer,omitempty"` // Manufacturer of the module

	ModuleIndex *int `json:"moduleIndex,omitempty"` // Index of the module

	Name string `json:"name,omitempty"` // Name of the module

	OperationalStateCode string `json:"operationalStateCode,omitempty"` // Operational state of the module

	PartNumber string `json:"partNumber,omitempty"` // Part number of the module

	SerialNumber string `json:"serialNumber,omitempty"` // Serial number of the module

	VendorEquipmentType string `json:"vendorEquipmentType,omitempty"` // Vendor euipment type of the module
}
type ResponseDevicesGetModulesV1ResponseAttributeInfo interface{}
type ResponseDevicesGetModuleCountV1 struct {
	Response *int `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetModuleInfoByIDV1 struct {
	Response *ResponseDevicesGetModuleInfoByIDV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetModuleInfoByIDV1Response struct {
	AssemblyNumber string `json:"assemblyNumber,omitempty"` // Assembly number of the module

	AssemblyRevision string `json:"assemblyRevision,omitempty"` // Assembly revision of the module

	AttributeInfo *ResponseDevicesGetModuleInfoByIDV1ResponseAttributeInfo `json:"attributeInfo,omitempty"` // Deprecated

	ContainmentEntity string `json:"containmentEntity,omitempty"` // Containment entity of the module

	Description string `json:"description,omitempty"` // Description of the module

	EntityPhysicalIndex string `json:"entityPhysicalIndex,omitempty"` // Entity physical index of the module

	ID string `json:"id,omitempty"` // Id of the module

	IsFieldReplaceable string `json:"isFieldReplaceable,omitempty"` // To mention if field is replaceable

	IsReportingAlarmsAllowed string `json:"isReportingAlarmsAllowed,omitempty"` // To mention if reporting alarms are allowed

	Manufacturer string `json:"manufacturer,omitempty"` // Manufacturer of the module

	ModuleIndex *int `json:"moduleIndex,omitempty"` // Index of the module

	Name string `json:"name,omitempty"` // Name of the module

	OperationalStateCode string `json:"operationalStateCode,omitempty"` // Operational state of the module

	PartNumber string `json:"partNumber,omitempty"` // Part number of the module

	SerialNumber string `json:"serialNumber,omitempty"` // Serial number of the modules

	VendorEquipmentType string `json:"vendorEquipmentType,omitempty"` // Vendor equipment type of the module
}
type ResponseDevicesGetModuleInfoByIDV1ResponseAttributeInfo interface{}
type ResponseDevicesGetDeviceBySerialNumberV1 struct {
	Response *ResponseDevicesGetDeviceBySerialNumberV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetDeviceBySerialNumberV1Response struct {
	ApManagerInterfaceIP string `json:"apManagerInterfaceIp,omitempty"` // IP address of WLC on AP manager interface

	AssociatedWlcIP string `json:"associatedWlcIp,omitempty"` // Associated Wlc Ip address of the AP device

	BootDateTime string `json:"bootDateTime,omitempty"` // Device boot time

	CollectionInterval string `json:"collectionInterval,omitempty"` // Re sync Interval of the device

	CollectionStatus string `json:"collectionStatus,omitempty"` // Collection status as Synchronizing, Could not synchronize, Not manageable, Managed, Partial Collection Failure, Incomplete, Unreachable, Wrong credential, Reachable, In Progress

	ErrorCode string `json:"errorCode,omitempty"` // Inventory status error code

	ErrorDescription string `json:"errorDescription,omitempty"` // Inventory status description

	Family string `json:"family,omitempty"` // Family of device as switch, router, wireless lan controller, accesspoints

	Hostname string `json:"hostname,omitempty"` // Device name

	ID string `json:"id,omitempty"` // Instance Uuid of the device

	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id of the device

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid of the device

	InterfaceCount string `json:"interfaceCount,omitempty"` // Number of interfaces on the device

	InventoryStatusDetail string `json:"inventoryStatusDetail,omitempty"` // Status detail of inventory sync

	LastUpdateTime *float64 `json:"lastUpdateTime,omitempty"` // Time in epoch when the network device info last got updated

	LastUpdated string `json:"lastUpdated,omitempty"` // Time when the network device info last got updated

	LineCardCount string `json:"lineCardCount,omitempty"` // Number of linecards on the device

	LineCardID string `json:"lineCardId,omitempty"` // IDs of linecards of the device

	Location string `json:"location,omitempty"` // [Deprecated] Location ID that is associated with the device

	LocationName string `json:"locationName,omitempty"` // [Deprecated] Name of the associated location

	MacAddress string `json:"macAddress,omitempty"` // MAC address of device

	ManagementIPAddress string `json:"managementIpAddress,omitempty"` // IP address of the device

	MemorySize string `json:"memorySize,omitempty"` // Processor memory size

	PlatformID string `json:"platformId,omitempty"` // Platform ID of device

	ReachabilityFailureReason string `json:"reachabilityFailureReason,omitempty"` // Failure reason for unreachable devices

	ReachabilityStatus string `json:"reachabilityStatus,omitempty"` // Device reachability status as Reachable / Unreachable

	Role string `json:"role,omitempty"` // Role of device as access, distribution, border router, core

	RoleSource string `json:"roleSource,omitempty"` // Role source as manual / auto

	SerialNumber string `json:"serialNumber,omitempty"` // Serial number of device

	Series string `json:"series,omitempty"` // Device series

	SNMPContact string `json:"snmpContact,omitempty"` // SNMP contact on device

	SNMPLocation string `json:"snmpLocation,omitempty"` // SNMP location on device

	SoftwareType string `json:"softwareType,omitempty"` // Software type on the device

	SoftwareVersion string `json:"softwareVersion,omitempty"` // Software version on the device

	TagCount string `json:"tagCount,omitempty"` // Number of tags associated with the device

	TunnelUDPPort string `json:"tunnelUdpPort,omitempty"` // Mobility protocol port is stored as tunneludpport for WLC

	Type string `json:"type,omitempty"` // Type of device as switch, router, wireless lan controller, accesspoints

	UpTime string `json:"upTime,omitempty"` // Time that shows for how long the device has been up

	WaasDeviceMode string `json:"waasDeviceMode,omitempty"` // WAAS device mode

	DNSResolvedManagementAddress string `json:"dnsResolvedManagementAddress,omitempty"` // Specifies the resolved ip address of dns name

	ApEthernetMacAddress string `json:"apEthernetMacAddress,omitempty"` // AccessPoint Ethernet MacAddress of AP device

	Vendor string `json:"vendor,omitempty"` // Vendor details

	ReasonsForPendingSyncRequests string `json:"reasonsForPendingSyncRequests,omitempty"` // Reasons for pending sync requests , if any

	PendingSyncRequestsCount string `json:"pendingSyncRequestsCount,omitempty"` // Count of pending sync requests , if any

	ReasonsForDeviceResync string `json:"reasonsForDeviceResync,omitempty"` // Reason for last/ongoing sync

	LastDeviceResyncStartTime string `json:"lastDeviceResyncStartTime,omitempty"` // Start time for last/ongoing sync

	UptimeSeconds *float64 `json:"uptimeSeconds,omitempty"` // Uptime in Seconds

	ManagedAtleastOnce *bool `json:"managedAtleastOnce,omitempty"` // Indicates if device went into Managed state atleast once

	DeviceSupportLevel string `json:"deviceSupportLevel,omitempty"` // Support level of the device

	ManagementState string `json:"managementState,omitempty"` // Represents the current management state of the network element: managed, unmanaged, under maintenance, and so on.

	Description string `json:"description,omitempty"` // System description
}
type ResponseDevicesSyncDevicesV1 struct {
	Response *ResponseDevicesSyncDevicesV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesSyncDevicesV1Response struct {
	TaskID string `json:"taskId,omitempty"` //

	URL string `json:"url,omitempty"` //
}
type ResponseDevicesGetDevicesRegisteredForWsaNotificationV1 struct {
	Response *ResponseDevicesGetDevicesRegisteredForWsaNotificationV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetDevicesRegisteredForWsaNotificationV1Response struct {
	MacAddress string `json:"macAddress,omitempty"` // MAC address of device

	ModelNumber string `json:"modelNumber,omitempty"` // Model number of the device

	Name string `json:"name,omitempty"` // Name of the device

	SerialNumber string `json:"serialNumber,omitempty"` // Serial Number of the device

	TenantID string `json:"tenantId,omitempty"` // Tenant Id of the device
}
type ResponseDevicesGetAllUserDefinedFieldsV1 struct {
	Response *[]ResponseDevicesGetAllUserDefinedFieldsV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesGetAllUserDefinedFieldsV1Response struct {
	ID string `json:"id,omitempty"` // DeviceId of the Device

	Name string `json:"name,omitempty"` // UDF name

	Description string `json:"description,omitempty"` // Description for UDF
}
type ResponseDevicesCreateUserDefinedFieldV1 struct {
	Response *ResponseDevicesCreateUserDefinedFieldV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesCreateUserDefinedFieldV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	URL string `json:"url,omitempty"` // Url
}
type ResponseDevicesUpdateUserDefinedFieldV1 struct {
	Response *ResponseDevicesUpdateUserDefinedFieldV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesUpdateUserDefinedFieldV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	URL string `json:"url,omitempty"` // Url
}
type ResponseDevicesDeleteUserDefinedFieldV1 struct {
	Response *ResponseDevicesDeleteUserDefinedFieldV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesDeleteUserDefinedFieldV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	URL string `json:"url,omitempty"` // Url
}
type ResponseDevicesGetChassisDetailsForDeviceV1 struct {
	Response *[]ResponseDevicesGetChassisDetailsForDeviceV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetChassisDetailsForDeviceV1Response struct {
	AssemblyNumber string `json:"assemblyNumber,omitempty"` // Assembly Number of the chassis

	AssemblyRevision string `json:"assemblyRevision,omitempty"` // Assembly Revision of the chassis

	ContainmentEntity string `json:"containmentEntity,omitempty"` // Containment Entity of the chassis

	Description string `json:"description,omitempty"` // Description of the chassis

	EntityPhysicalIndex string `json:"entityPhysicalIndex,omitempty"` // Entity Physical Index of the chassis

	HardwareVersion string `json:"hardwareVersion,omitempty"` // Hardware Version of the chassis

	InstanceUUID string `json:"instanceUuid,omitempty"` // ID of the chassis

	IsFieldReplaceable string `json:"isFieldReplaceable,omitempty"` // To mention if field is replaceable

	IsReportingAlarmsAllowed string `json:"isReportingAlarmsAllowed,omitempty"` // To mention if reporting alarms are allowed

	Manufacturer string `json:"manufacturer,omitempty"` // Manufacturer of the chassis

	Name string `json:"name,omitempty"` // Name of the chassis

	PartNumber string `json:"partNumber,omitempty"` // Part Number of the chassis

	SerialNumber string `json:"serialNumber,omitempty"` // Serial Number of the chassis

	VendorEquipmentType string `json:"vendorEquipmentType,omitempty"` // Vendor Equipment Type of the chassis
}
type ResponseDevicesGetStackDetailsForDeviceV1 struct {
	Response *ResponseDevicesGetStackDetailsForDeviceV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetStackDetailsForDeviceV1Response struct {
	DeviceID string `json:"deviceId,omitempty"` // Device ID

	StackPortInfo *[]ResponseDevicesGetStackDetailsForDeviceV1ResponseStackPortInfo `json:"stackPortInfo,omitempty"` //

	StackSwitchInfo *[]ResponseDevicesGetStackDetailsForDeviceV1ResponseStackSwitchInfo `json:"stackSwitchInfo,omitempty"` //

	SvlSwitchInfo *[]ResponseDevicesGetStackDetailsForDeviceV1ResponseSvlSwitchInfo `json:"svlSwitchInfo,omitempty"` //
}
type ResponseDevicesGetStackDetailsForDeviceV1ResponseStackPortInfo struct {
	IsSynchOk string `json:"isSynchOk,omitempty"` // If link partner sends valid protocol message

	LinkActive *bool `json:"linkActive,omitempty"` // If stack port is in same state as link partner

	LinkOk *bool `json:"linkOk,omitempty"` // If link is stable

	Name string `json:"name,omitempty"` // Name of the stack port

	NeighborPort string `json:"neighborPort,omitempty"` // Neighbor's member number and stack port number

	NrLinkOkChanges *int `json:"nrLinkOkChanges,omitempty"` // Relative stability of the link

	StackCableLengthInfo string `json:"stackCableLengthInfo,omitempty"` // Cable length

	StackPortOperStatusInfo string `json:"stackPortOperStatusInfo,omitempty"` // Port opearation status

	SwitchPort string `json:"switchPort,omitempty"` // Member number and stack port number
}
type ResponseDevicesGetStackDetailsForDeviceV1ResponseStackSwitchInfo struct {
	EntPhysicalIndex string `json:"entPhysicalIndex,omitempty"` //

	HwPriority *int `json:"hwPriority,omitempty"` // Hardware priority of the switch

	MacAddress string `json:"macAddress,omitempty"` // Mac address of the switch

	NumNextReload *int `json:"numNextReload,omitempty"` // Stack member number to be used in next reload

	PlatformID string `json:"platformId,omitempty"` // Platform Id

	Role string `json:"role,omitempty"` // Function of the switch

	SerialNumber string `json:"serialNumber,omitempty"` // Serial number

	SoftwareImage string `json:"softwareImage,omitempty"` // Software image type running on the switch

	StackMemberNumber *int `json:"stackMemberNumber,omitempty"` // Switch member number

	State string `json:"state,omitempty"` // Current state of the switch

	SwitchPriority *int `json:"switchPriority,omitempty"` // Priority of the switch
}
type ResponseDevicesGetStackDetailsForDeviceV1ResponseSvlSwitchInfo struct {
	DadProtocol string `json:"dadProtocol,omitempty"` // Stackwise virtual dual active detection config

	DadRecoveryReloadEnabled *bool `json:"dadRecoveryReloadEnabled,omitempty"` // If dad recovery reload enabled

	DomainNumber *int `json:"domainNumber,omitempty"` // Stackwise virtual switch domain number

	InDadRecoveryMode *bool `json:"inDadRecoveryMode,omitempty"` // If in dad recovery mode

	SwVirtualStatus string `json:"swVirtualStatus,omitempty"` // Stackwise virtual status

	SwitchMembers *[]ResponseDevicesGetStackDetailsForDeviceV1ResponseSvlSwitchInfoSwitchMembers `json:"switchMembers,omitempty"` //
}
type ResponseDevicesGetStackDetailsForDeviceV1ResponseSvlSwitchInfoSwitchMembers struct {
	Bandwidth string `json:"bandwidth,omitempty"` // Bandwidth

	SvlMemberEndPoints *[]ResponseDevicesGetStackDetailsForDeviceV1ResponseSvlSwitchInfoSwitchMembersSvlMemberEndPoints `json:"svlMemberEndPoints,omitempty"` //

	SvlMemberNumber *int `json:"svlMemberNumber,omitempty"` // Switch member number

	SvlMemberPepSettings *[]ResponseDevicesGetStackDetailsForDeviceV1ResponseSvlSwitchInfoSwitchMembersSvlMemberPepSettings `json:"svlMemberPepSettings,omitempty"` //
}
type ResponseDevicesGetStackDetailsForDeviceV1ResponseSvlSwitchInfoSwitchMembersSvlMemberEndPoints struct {
	SvlMemberEndPointPorts *[]ResponseDevicesGetStackDetailsForDeviceV1ResponseSvlSwitchInfoSwitchMembersSvlMemberEndPointsSvlMemberEndPointPorts `json:"svlMemberEndPointPorts,omitempty"` //

	SvlNumber *int `json:"svlNumber,omitempty"` // Stackwise virtual link number

	SvlStatus string `json:"svlStatus,omitempty"` // Stackwise virtual status
}
type ResponseDevicesGetStackDetailsForDeviceV1ResponseSvlSwitchInfoSwitchMembersSvlMemberEndPointsSvlMemberEndPointPorts struct {
	SvlProtocolStatus string `json:"svlProtocolStatus,omitempty"` // Stackwise virtual protocol status

	SwLocalInterface string `json:"swLocalInterface,omitempty"` // Stackwise virtual local interface

	SwRemoteInterface string `json:"swRemoteInterface,omitempty"` // Stackwise virtual remote interface
}
type ResponseDevicesGetStackDetailsForDeviceV1ResponseSvlSwitchInfoSwitchMembersSvlMemberPepSettings struct {
	DadEnabled *bool `json:"dadEnabled,omitempty"` // If dadInterface is configured for dual active detection

	DadInterfaceName string `json:"dadInterfaceName,omitempty"` // Interface for dual active detection
}
type ResponseDevicesRemoveUserDefinedFieldFromDeviceV1 struct {
	Response *ResponseDevicesRemoveUserDefinedFieldFromDeviceV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesRemoveUserDefinedFieldFromDeviceV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	URL string `json:"url,omitempty"` // Url
}
type ResponseDevicesAddUserDefinedFieldToDeviceV1 struct {
	Response *ResponseDevicesAddUserDefinedFieldToDeviceV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesAddUserDefinedFieldToDeviceV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	URL string `json:"url,omitempty"` // Url
}
type ResponseDevicesGetTheDetailsOfPhysicalComponentsOfTheGivenDeviceV1 struct {
	Response *[]ResponseDevicesGetTheDetailsOfPhysicalComponentsOfTheGivenDeviceV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesGetTheDetailsOfPhysicalComponentsOfTheGivenDeviceV1Response struct {
	OperationalStateCode string `json:"operationalStateCode,omitempty"` // Operational State Code

	ProductID string `json:"productId,omitempty"` // Product Id

	SerialNumber string `json:"serialNumber,omitempty"` // Serial Number

	VendorEquipmentType string `json:"vendorEquipmentType,omitempty"` // Vendor Equipment Type

	Description string `json:"description,omitempty"` // Description

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid

	Name string `json:"name,omitempty"` // Name

	Manufacturer string `json:"manufacturer,omitempty"` // Manufacturer
}
type ResponseDevicesReturnsPoeInterfaceDetailsForTheDeviceV1 struct {
	Version string `json:"version,omitempty"` // Version

	Response *[]ResponseDevicesReturnsPoeInterfaceDetailsForTheDeviceV1Response `json:"response,omitempty"` //
}
type ResponseDevicesReturnsPoeInterfaceDetailsForTheDeviceV1Response struct {
	AdminStatus string `json:"adminStatus,omitempty"` // Administration Status. Possible values: AUTO, STATIC, NEVER

	OperStatus string `json:"operStatus,omitempty"` // Operational Status. Possible values: ON, OFF, FAULTY, POWER_DENY

	InterfaceName string `json:"interfaceName,omitempty"` // Name of the interface

	MaxPortPower string `json:"maxPortPower,omitempty"` // Maximum power (in Watts) that port can hold

	AllocatedPower string `json:"allocatedPower,omitempty"` // Power (in Watts) allocated for a given interface

	PortPowerDrawn string `json:"portPowerDrawn,omitempty"` // Power (in Watts) that the port has drawn so far
}
type ResponseDevicesGetConnectedDeviceDetailV1 struct {
	Response *ResponseDevicesGetConnectedDeviceDetailV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesGetConnectedDeviceDetailV1Response struct {
	NeighborDevice string `json:"neighborDevice,omitempty"` // Info about the devices connected to the interface

	NeighborPort string `json:"neighborPort,omitempty"` // Info about the connected interface

	Capabilities []string `json:"capabilities,omitempty"` // Info about capabilities of the connected device
}
type ResponseDevicesGetLinecardDetailsV1 struct {
	Response *[]ResponseDevicesGetLinecardDetailsV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesGetLinecardDetailsV1Response struct {
	Serialno string `json:"serialno,omitempty"` // Serial number of the line card

	Partno string `json:"partno,omitempty"` // Part number of the line card

	Switchno string `json:"switchno,omitempty"` // Switch number of the line card

	Slotno string `json:"slotno,omitempty"` // Slot number of line card
}
type ResponseDevicesPoeDetailsV1 struct {
	Response *ResponseDevicesPoeDetailsV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesPoeDetailsV1Response struct {
	PowerAllocated string `json:"powerAllocated,omitempty"` // Total power available on the switch on all interfaces combined in Watts

	PowerConsumed string `json:"powerConsumed,omitempty"` // Total power being currently drawn by all interfaces combined in Watts

	PowerRemaining string `json:"powerRemaining,omitempty"` // Total power remaining in Watts (powerConsumed - powerAllocated)
}
type ResponseDevicesGetSupervisorCardDetailV1 struct {
	Response *[]ResponseDevicesGetSupervisorCardDetailV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesGetSupervisorCardDetailV1Response struct {
	Serialno string `json:"serialno,omitempty"` // Serial number of the supervisor card

	Partno string `json:"partno,omitempty"` // Part number of the supervisor card

	Switchno string `json:"switchno,omitempty"` // Switch number of the supervisor card

	Slotno string `json:"slotno,omitempty"` // Slot number of supervisor card
}
type ResponseDevicesUpdateDeviceManagementAddressV1 struct {
	Response *ResponseDevicesUpdateDeviceManagementAddressV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesUpdateDeviceManagementAddressV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	URL string `json:"url,omitempty"` // Url
}
type ResponseDevicesGetDeviceByIDV1 struct {
	Response *ResponseDevicesGetDeviceByIDV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetDeviceByIDV1Response struct {
	ApManagerInterfaceIP string `json:"apManagerInterfaceIp,omitempty"` // IP address of WLC on AP manager interface

	AssociatedWlcIP string `json:"associatedWlcIp,omitempty"` // Associated Wlc Ip address of the AP device

	BootDateTime string `json:"bootDateTime,omitempty"` // Device boot time

	CollectionInterval string `json:"collectionInterval,omitempty"` // Re sync Interval of the device

	CollectionStatus string `json:"collectionStatus,omitempty"` // Collection status as Synchronizing, Could not synchronize, Not manageable, Managed, Partial Collection Failure, Incomplete, Unreachable, Wrong credential, Reachable, In Progress

	ErrorCode string `json:"errorCode,omitempty"` // Inventory status error code

	ErrorDescription string `json:"errorDescription,omitempty"` // Inventory status description

	Family string `json:"family,omitempty"` // Family of device as switch, router, wireless lan controller, accesspoints

	Hostname string `json:"hostname,omitempty"` // Device name

	ID string `json:"id,omitempty"` // Instance Uuid of the device

	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id of the device

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid of the device

	InterfaceCount string `json:"interfaceCount,omitempty"` // Number of interfaces on the device

	InventoryStatusDetail string `json:"inventoryStatusDetail,omitempty"` // Status detail of inventory sync

	LastUpdateTime *float64 `json:"lastUpdateTime,omitempty"` // Time in epoch when the network device info last got updated

	LastUpdated string `json:"lastUpdated,omitempty"` // Time when the network device info last got updated

	LineCardCount string `json:"lineCardCount,omitempty"` // Number of linecards on the device

	LineCardID string `json:"lineCardId,omitempty"` // IDs of linecards of the device

	Location string `json:"location,omitempty"` // [Deprecated] Location ID that is associated with the device

	LocationName string `json:"locationName,omitempty"` // [Deprecated] Name of the associated location

	MacAddress string `json:"macAddress,omitempty"` // MAC address of device

	ManagementIPAddress string `json:"managementIpAddress,omitempty"` // IP address of the device

	MemorySize string `json:"memorySize,omitempty"` // Processor memory size

	PlatformID string `json:"platformId,omitempty"` // Platform ID of device

	ReachabilityFailureReason string `json:"reachabilityFailureReason,omitempty"` // Failure reason for unreachable devices

	ReachabilityStatus string `json:"reachabilityStatus,omitempty"` // Device reachability status as Reachable / Unreachable

	Role string `json:"role,omitempty"` // Role of device as access, distribution, border router, core

	RoleSource string `json:"roleSource,omitempty"` // Role source as manual / auto

	SerialNumber string `json:"serialNumber,omitempty"` // Serial number of device

	Series string `json:"series,omitempty"` // Device series

	SNMPContact string `json:"snmpContact,omitempty"` // SNMP contact on device

	SNMPLocation string `json:"snmpLocation,omitempty"` // SNMP location on device

	SoftwareType string `json:"softwareType,omitempty"` // Software type on the device

	SoftwareVersion string `json:"softwareVersion,omitempty"` // Software version on the device

	TagCount string `json:"tagCount,omitempty"` // Number of tags associated with the device

	TunnelUDPPort string `json:"tunnelUdpPort,omitempty"` // Mobility protocol port is stored as tunneludpport for WLC

	Type string `json:"type,omitempty"` // Type of device as switch, router, wireless lan controller, accesspoints

	UpTime string `json:"upTime,omitempty"` // Time that shows for how long the device has been up

	WaasDeviceMode string `json:"waasDeviceMode,omitempty"` // WAAS device mode

	DNSResolvedManagementAddress string `json:"dnsResolvedManagementAddress,omitempty"` // Specifies the resolved ip address of dns name

	ApEthernetMacAddress string `json:"apEthernetMacAddress,omitempty"` // AccessPoint Ethernet MacAddress of AP device

	Vendor string `json:"vendor,omitempty"` // Vendor details

	ReasonsForPendingSyncRequests string `json:"reasonsForPendingSyncRequests,omitempty"` // Reasons for pending sync requests , if any

	PendingSyncRequestsCount string `json:"pendingSyncRequestsCount,omitempty"` // Count of pending sync requests , if any

	ReasonsForDeviceResync string `json:"reasonsForDeviceResync,omitempty"` // Reason for last/ongoing sync

	LastDeviceResyncStartTime string `json:"lastDeviceResyncStartTime,omitempty"` // Start time for last/ongoing sync

	UptimeSeconds *float64 `json:"uptimeSeconds,omitempty"` // Uptime in Seconds

	ManagedAtleastOnce *bool `json:"managedAtleastOnce,omitempty"` // Indicates if device went into Managed state atleast once

	DeviceSupportLevel string `json:"deviceSupportLevel,omitempty"` // Support level of the device

	ManagementState string `json:"managementState,omitempty"` // Represents the current management state of the network element: managed, unmanaged, under maintenance, and so on.

	Description string `json:"description,omitempty"` // System description
}
type ResponseDevicesDeleteDeviceByIDV1 struct {
	Response *ResponseDevicesDeleteDeviceByIDV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesDeleteDeviceByIDV1Response struct {
	TaskID string `json:"taskId,omitempty"` //

	URL string `json:"url,omitempty"` //
}
type ResponseDevicesGetDeviceSummaryV1 struct {
	Response *ResponseDevicesGetDeviceSummaryV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetDeviceSummaryV1Response struct {
	ID string `json:"id,omitempty"` // Unique identifier of the network device

	Role string `json:"role,omitempty"` // Role of device as access, distribution, border router, core

	RoleSource string `json:"roleSource,omitempty"` // Role source as manual / auto
}
type ResponseDevicesGetPollingIntervalByIDV1 struct {
	Response *int `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetOrganizationListForMerakiV1 struct {
	Response []string `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetDeviceInterfaceVLANsV1 struct {
	Response *[]ResponseDevicesGetDeviceInterfaceVLANsV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetDeviceInterfaceVLANsV1Response struct {
	InterfaceName string `json:"interfaceName,omitempty"` // Interface name

	IPAddress string `json:"ipAddress,omitempty"` // IP address

	Mask *int `json:"mask,omitempty"` // Mask IP

	NetworkAddress string `json:"networkAddress,omitempty"` // Network addresses

	NumberOfIPs *int `json:"numberOfIPs,omitempty"` // Number of Ip addresses

	Prefix string `json:"prefix,omitempty"` // Prefix associated with the IP address

	VLANNumber *int `json:"vlanNumber,omitempty"` // Vlan Number

	VLANType string `json:"vlanType,omitempty"` // [Deprecated] Description of the interface VLAN
}
type ResponseDevicesGetWirelessLanControllerDetailsByIDV1 struct {
	AdminEnabledPorts *[]int `json:"adminEnabledPorts,omitempty"` // Admin Enabled Ports of the Device

	ApGroupName string `json:"apGroupName,omitempty"` // Name of the AP Group that Access point assigned

	DeviceID string `json:"deviceId,omitempty"` // Device Id

	EthMacAddress string `json:"ethMacAddress,omitempty"` // Ethernet MacAddress of the Device

	FlexGroupName string `json:"flexGroupName,omitempty"` // Name of the Flex Group that Access point assigned

	ID string `json:"id,omitempty"` // Id of the Device

	InstanceTenantID string `json:"instanceTenantId,omitempty"` // TenantId of the Device

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance UUID of the Device

	LagModeEnabled *bool `json:"lagModeEnabled,omitempty"` // LagMode status of the Device

	NetconfEnabled *bool `json:"netconfEnabled,omitempty"` // Netconf Status of the Device

	WirelessLicenseInfo string `json:"wirelessLicenseInfo,omitempty"` // License type of Wireless Device

	WirelessPackageInstalled *bool `json:"wirelessPackageInstalled,omitempty"` // Status of the Wireless Package on the Device
}
type ResponseDevicesGetDeviceConfigByIDV1 struct {
	Response string `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetNetworkDeviceByPaginationRangeV1 struct {
	Response *[]ResponseDevicesGetNetworkDeviceByPaginationRangeV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetNetworkDeviceByPaginationRangeV1Response struct {
	ApManagerInterfaceIP string `json:"apManagerInterfaceIp,omitempty"` // IP address of WLC on AP manager interface

	AssociatedWlcIP string `json:"associatedWlcIp,omitempty"` // Associated Wlc Ip address of the AP device

	BootDateTime string `json:"bootDateTime,omitempty"` // Device boot time

	CollectionInterval string `json:"collectionInterval,omitempty"` // Re sync Interval of the device

	CollectionStatus string `json:"collectionStatus,omitempty"` // Collection status as Synchronizing, Could not synchronize, Not manageable, Managed, Partial Collection Failure, Incomplete, Unreachable, Wrong credential, Reachable, In Progress

	ErrorCode string `json:"errorCode,omitempty"` // Inventory status error code

	ErrorDescription string `json:"errorDescription,omitempty"` // Inventory status description

	Family string `json:"family,omitempty"` // Family of device as switch, router, wireless lan controller, accesspoints

	Hostname string `json:"hostname,omitempty"` // Device name

	ID string `json:"id,omitempty"` // Instance Uuid of the device

	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id of the device

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid of the device

	InterfaceCount string `json:"interfaceCount,omitempty"` // Number of interfaces on the device

	InventoryStatusDetail string `json:"inventoryStatusDetail,omitempty"` // Status detail of inventory sync

	LastUpdateTime *float64 `json:"lastUpdateTime,omitempty"` // Time in epoch when the network device info last got updated

	LastUpdated string `json:"lastUpdated,omitempty"` // Time when the network device info last got updated

	LineCardCount string `json:"lineCardCount,omitempty"` // Number of linecards on the device

	LineCardID string `json:"lineCardId,omitempty"` // IDs of linecards of the device

	Location string `json:"location,omitempty"` // [Deprecated] Location ID that is associated with the device

	LocationName string `json:"locationName,omitempty"` // [Deprecated] Name of the associated location

	MacAddress string `json:"macAddress,omitempty"` // MAC address of device

	ManagementIPAddress string `json:"managementIpAddress,omitempty"` // IP address of the device

	MemorySize string `json:"memorySize,omitempty"` // Processor memory size

	PlatformID string `json:"platformId,omitempty"` // Platform ID of device

	ReachabilityFailureReason string `json:"reachabilityFailureReason,omitempty"` // Failure reason for unreachable devices

	ReachabilityStatus string `json:"reachabilityStatus,omitempty"` // Device reachability status as Reachable / Unreachable

	Role string `json:"role,omitempty"` // Role of device as access, distribution, border router, core

	RoleSource string `json:"roleSource,omitempty"` // Role source as manual / auto

	SerialNumber string `json:"serialNumber,omitempty"` // Serial number of device

	Series string `json:"series,omitempty"` // Device series

	SNMPContact string `json:"snmpContact,omitempty"` // SNMP contact on device

	SNMPLocation string `json:"snmpLocation,omitempty"` // SNMP location on device

	SoftwareType string `json:"softwareType,omitempty"` // Software type on the device

	SoftwareVersion string `json:"softwareVersion,omitempty"` // Software version on the device

	TagCount string `json:"tagCount,omitempty"` // Number of tags associated with the device

	TunnelUDPPort string `json:"tunnelUdpPort,omitempty"` // Mobility protocol port is stored as tunneludpport for WLC

	Type string `json:"type,omitempty"` // Type of device as switch, router, wireless lan controller, accesspoints

	UpTime string `json:"upTime,omitempty"` // Time that shows for how long the device has been up

	WaasDeviceMode string `json:"waasDeviceMode,omitempty"` // WAAS device mode

	DNSResolvedManagementAddress string `json:"dnsResolvedManagementAddress,omitempty"` // Specifies the resolved ip address of dns name

	ApEthernetMacAddress string `json:"apEthernetMacAddress,omitempty"` // AccessPoint Ethernet MacAddress of AP device

	Vendor string `json:"vendor,omitempty"` // Vendor details

	ReasonsForPendingSyncRequests string `json:"reasonsForPendingSyncRequests,omitempty"` // Reasons for pending sync requests , if any

	PendingSyncRequestsCount string `json:"pendingSyncRequestsCount,omitempty"` // Count of pending sync requests , if any

	ReasonsForDeviceResync string `json:"reasonsForDeviceResync,omitempty"` // Reason for last/ongoing sync

	LastDeviceResyncStartTime string `json:"lastDeviceResyncStartTime,omitempty"` // Start time for last/ongoing sync

	UptimeSeconds *float64 `json:"uptimeSeconds,omitempty"` // Uptime in Seconds

	ManagedAtleastOnce *bool `json:"managedAtleastOnce,omitempty"` // Indicates if device went into Managed state atleast once

	DeviceSupportLevel string `json:"deviceSupportLevel,omitempty"` // Support level of the device

	ManagementState string `json:"managementState,omitempty"` // Represents the current management state of the network element: managed, unmanaged, under maintenance, and so on.

	Description string `json:"description,omitempty"` // System description
}
type ResponseDevicesCreateMaintenanceScheduleForNetworkDevicesV1 struct {
	Response *ResponseDevicesCreateMaintenanceScheduleForNetworkDevicesV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version of the response
}
type ResponseDevicesCreateMaintenanceScheduleForNetworkDevicesV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Unique identifier for the task

	URL string `json:"url,omitempty"` // URL for the task
}
type ResponseDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesV1 struct {
	Response *[]ResponseDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number of the response
}
type ResponseDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesV1Response struct {
	ID string `json:"id,omitempty"` // Id of the schedule maintenance window

	Description string `json:"description,omitempty"` // A brief narrative describing the maintenance schedule.

	MaintenanceSchedule *ResponseDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesV1ResponseMaintenanceSchedule `json:"maintenanceSchedule,omitempty"` //

	NetworkDeviceIDs []string `json:"networkDeviceIds,omitempty"` // List of network device ids. This field is applicable only during creation of schedules; for updates, it is read-only.
}
type ResponseDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesV1ResponseMaintenanceSchedule struct {
	StartID string `json:"startId,omitempty"` // Activity id of start schedule of the maintenance window. To check the status of the start schedule, use GET /dna/intent/api/v1/activities/{id}. startId remains same for every occurrence of recurrence instance.

	EndID string `json:"endId,omitempty"` // Activity id of end schedule of the maintenance window. To check the status of the end schedule, use GET /dna/intent/api/v1/activities/{id}. endId remains same for every occurrence of recurrence instance.

	StartTime *float64 `json:"startTime,omitempty"` // Start time indicates the beginning of the maintenance window in Unix epoch time in milliseconds.

	EndTime *float64 `json:"endTime,omitempty"` // End time indicates the ending of the maintenance window in Unix epoch time in milliseconds.

	Recurrence *ResponseDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesV1ResponseMaintenanceScheduleRecurrence `json:"recurrence,omitempty"` //

	Status string `json:"status,omitempty"` // The status of the maintenance schedule.
}
type ResponseDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesV1ResponseMaintenanceScheduleRecurrence struct {
	Interval *int `json:"interval,omitempty"` // Interval for recurrence in days. The interval must be longer than the duration of the schedules. The maximum allowed interval is 365 days.

	RecurrenceEndTime *float64 `json:"recurrenceEndTime,omitempty"` // The end date for the recurrence in Unix epoch time in milliseconds. Recurrence end time should be greater than maintenance end date/time.
}
type ResponseDevicesRetrieveTheTotalNumberOfScheduledMaintenanceWindowsV1 struct {
	Response *ResponseDevicesRetrieveTheTotalNumberOfScheduledMaintenanceWindowsV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // The version of the response
}
type ResponseDevicesRetrieveTheTotalNumberOfScheduledMaintenanceWindowsV1Response struct {
	Count *int `json:"count,omitempty"` // Count of scheduled maintenance windows
}
type ResponseDevicesUpdatesTheMaintenanceScheduleInformationV1 struct {
	Response *ResponseDevicesUpdatesTheMaintenanceScheduleInformationV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version of the response
}
type ResponseDevicesUpdatesTheMaintenanceScheduleInformationV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Unique identifier for the task

	URL string `json:"url,omitempty"` // URL for the task
}
type ResponseDevicesRetrievesTheMaintenanceScheduleInformationV1 struct {
	Response *ResponseDevicesRetrievesTheMaintenanceScheduleInformationV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number of the response
}
type ResponseDevicesRetrievesTheMaintenanceScheduleInformationV1Response struct {
	ID string `json:"id,omitempty"` // Id of the schedule maintenance window

	Description string `json:"description,omitempty"` // A brief narrative describing the maintenance schedule.

	MaintenanceSchedule *ResponseDevicesRetrievesTheMaintenanceScheduleInformationV1ResponseMaintenanceSchedule `json:"maintenanceSchedule,omitempty"` //

	NetworkDeviceIDs []string `json:"networkDeviceIds,omitempty"` // List of network device ids. This field is applicable only during creation of schedules; for updates, it is read-only.
}
type ResponseDevicesRetrievesTheMaintenanceScheduleInformationV1ResponseMaintenanceSchedule struct {
	StartID string `json:"startId,omitempty"` // Activity id of start schedule of the maintenance window. To check the status of the start schedule, use GET /intent/api/v1/activities/{id}. startId remains same for every occurrence of recurrence instance.

	EndID string `json:"endId,omitempty"` // Activity id of end schedule of the maintenance window. To check the status of the end schedule, use GET /intent/api/v1/activities/{id}. endId remains same for every occurrence of recurrence instance.

	StartTime *float64 `json:"startTime,omitempty"` // Start time indicates the beginning of the maintenance window in Unix epoch time in milliseconds.

	EndTime *float64 `json:"endTime,omitempty"` // End time indicates the ending of the maintenance window in Unix epoch time in milliseconds.

	Recurrence *ResponseDevicesRetrievesTheMaintenanceScheduleInformationV1ResponseMaintenanceScheduleRecurrence `json:"recurrence,omitempty"` //

	Status string `json:"status,omitempty"` // The status of the maintenance schedule.
}
type ResponseDevicesRetrievesTheMaintenanceScheduleInformationV1ResponseMaintenanceScheduleRecurrence struct {
	Interval *int `json:"interval,omitempty"` // Interval for recurrence in days. The interval must be longer than the duration of the schedules. The maximum allowed interval is 365 days.

	RecurrenceEndTime *float64 `json:"recurrenceEndTime,omitempty"` // The end date for the recurrence in Unix epoch time in milliseconds. Recurrence end time should be greater than maintenance end date/time.
}
type ResponseDevicesDeleteMaintenanceScheduleV1 struct {
	Response *ResponseDevicesDeleteMaintenanceScheduleV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version of the response
}
type ResponseDevicesDeleteMaintenanceScheduleV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Unique identifier for the task

	URL string `json:"url,omitempty"` // URL for the task
}
type ResponseDevicesRetrieveNetworkDevicesV1 struct {
	Response *[]ResponseDevicesRetrieveNetworkDevicesV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number of the response
}
type ResponseDevicesRetrieveNetworkDevicesV1Response struct {
	ID string `json:"id,omitempty"` // Unique identifier of the network device

	ManagementAddress string `json:"managementAddress,omitempty"` // Management address of the network device

	DNSResolvedManagementIPAddress string `json:"dnsResolvedManagementIpAddress,omitempty"` // DNS-resolved management IP address of the network device

	Hostname string `json:"hostname,omitempty"` // Hostname of the network device

	MacAddress string `json:"macAddress,omitempty"` // MAC address of the network device

	SerialNumbers []string `json:"serialNumbers,omitempty"` // Serial number of the network device. In case of stack device, there will be multiple serial numbers

	Type string `json:"type,omitempty"` // Type of the network device. This list of types can be obtained from the API intent/networkDeviceProductNames productName field.

	Family string `json:"family,omitempty"` // Product family of the network device. For example, Switches, Routers, etc

	Series string `json:"series,omitempty"` // The model range or series of the network device

	Status string `json:"status,omitempty"` // Inventory related status of the network device. Refer features for more details

	PlatformIDs []string `json:"platformIds,omitempty"` // Platform identifier of the network device

	SoftwareType string `json:"softwareType,omitempty"` // Type of software running on the network device. For example, IOS-XE, etc.

	SoftwareVersion string `json:"softwareVersion,omitempty"` // Version of the software running on the network device

	Vendor string `json:"vendor,omitempty"` // Vendor of the network device

	StackDevice *bool `json:"stackDevice,omitempty"` // Flag indicating if the network device is a stack device

	BootTime *float64 `json:"bootTime,omitempty"` // The time at which the network device was last rebooted or powered on represented as epoch in milliseconds

	Role string `json:"role,omitempty"` // Role assigned to the network device

	RoleSource string `json:"roleSource,omitempty"` // Indicates whether the network device's role was assigned automatically by the software or manually by an administrator.

	ApEthernetMacAddress string `json:"apEthernetMacAddress,omitempty"` // Ethernet MAC address of the AP network device

	ApManagerInterfaceIPAddress string `json:"apManagerInterfaceIpAddress,omitempty"` // Management IP address of the AP network device

	ApWlcIPAddress string `json:"apWlcIpAddress,omitempty"` // Management IP address of the WLC on which AP is associated to

	DeviceSupportLevel string `json:"deviceSupportLevel,omitempty"` // The level of support Catalyst Center provides for the network device.

	SNMPLocation string `json:"snmpLocation,omitempty"` // SNMP location of the network device

	SNMPContact string `json:"snmpContact,omitempty"` // SNMP contact of the network device

	ReachabilityStatus string `json:"reachabilityStatus,omitempty"` // Reachability status of the network device. Refer features for more details

	ReachabilityFailureReason string `json:"reachabilityFailureReason,omitempty"` // Reason for reachability failure. This message that provides more information about the reachability failure.

	ManagementState string `json:"managementState,omitempty"` // The status of the network device's manageability. Refer features for more details.

	LastSuccessfulResyncReasons []string `json:"lastSuccessfulResyncReasons,omitempty"` // List of reasons for the last successful resync of the device. If multiple resync requests are made before the device can start the resync, all the reasons will be captured. Possible values: ADD_DEVICE_SYNC, LINK_UP_DOWN, CONFIG_CHANGE, DEVICE_UPDATED_SYNC, AP_EVENT_BASED_SYNC, APP_REQUESTED_SYNC, PERIODIC_SYNC, UI_SYNC, CUSTOM, UNKNOWN, REFRESH_OBJECTS_FEATURE_BASED_SYNC

	ResyncStartTime *float64 `json:"resyncStartTime,omitempty"` // Start time for the last/ongoing resync represented as epoch in milliseconds

	ResyncEndTime *float64 `json:"resyncEndTime,omitempty"` // End time for the last resync represented as epoch in milliseconds

	ResyncReasons []string `json:"resyncReasons,omitempty"` // List of reasons for the ongoing/last resync on the device. If multiple resync requests were made before the resync could start, all the reasons will be captured as an array. Possible values: ADD_DEVICE_SYNC, LINK_UP_DOWN, CONFIG_CHANGE, DEVICE_UPDATED_SYNC, AP_EVENT_BASED_SYNC, APP_REQUESTED_SYNC, PERIODIC_SYNC, UI_SYNC, CUSTOM, UNKNOWN, REFRESH_OBJECTS_FEATURE_BASED_SYNC

	ResyncRequestedByApps []string `json:"resyncRequestedByApps,omitempty"` // List of applications that requested the last/ongoing resync on the device

	PendingResyncRequestCount *int `json:"pendingResyncRequestCount,omitempty"` // Number of pending resync requests for the device

	PendingResyncRequestReasons []string `json:"pendingResyncRequestReasons,omitempty"` // List of reasons for the pending resync requests. Possible values: ADD_DEVICE_SYNC, LINK_UP_DOWN, CONFIG_CHANGE, DEVICE_UPDATED_SYNC, AP_EVENT_BASED_SYNC, APP_REQUESTED_SYNC, PERIODIC_SYNC, UI_SYNC, CUSTOM, UNKNOWN, REFRESH_OBJECTS_FEATURE_BASED_SYNC

	ResyncIntervalSource string `json:"resyncIntervalSource,omitempty"` // Source of the resync interval. Note: Please refer to PUT /dna/intent/api/v1/networkDevices/resyncIntervalSettings API to update the global resync interval.

	ResyncIntervalMinutes *int `json:"resyncIntervalMinutes,omitempty"` // The duration in minutes between the periodic resync attempts for the device

	ErrorCode string `json:"errorCode,omitempty"` // Error code indicating the reason for the last resync failure

	ErrorDescription string `json:"errorDescription,omitempty"` // Additional information regarding the reason for resync failure. This is a human-readable error message and should not be expected programmatically.

	UserDefinedFields *ResponseDevicesRetrieveNetworkDevicesV1ResponseUserDefinedFields `json:"userDefinedFields,omitempty"` // Map of all user defined fields and their values associated with the device. Refer to /dna/intent/api/v1/network-device/user-defined-field API to fetch all the user defined fields.
}
type ResponseDevicesRetrieveNetworkDevicesV1ResponseUserDefinedFields interface{}
type ResponseDevicesCountTheNumberOfNetworkDevicesV1 struct {
	Response *ResponseDevicesCountTheNumberOfNetworkDevicesV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // The version of the response
}
type ResponseDevicesCountTheNumberOfNetworkDevicesV1Response struct {
	Count *int `json:"count,omitempty"` // The reported count.
}
type ResponseDevicesDeleteNetworkDeviceWithConfigurationCleanupV1 struct {
	Response *ResponseDevicesDeleteNetworkDeviceWithConfigurationCleanupV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version of the response
}
type ResponseDevicesDeleteNetworkDeviceWithConfigurationCleanupV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Unique identifier for the task

	URL string `json:"url,omitempty"` // URL for the task
}
type ResponseDevicesDeleteANetworkDeviceWithoutConfigurationCleanupV1 struct {
	Response *ResponseDevicesDeleteANetworkDeviceWithoutConfigurationCleanupV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version of the response
}
type ResponseDevicesDeleteANetworkDeviceWithoutConfigurationCleanupV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Unique identifier for the task

	URL string `json:"url,omitempty"` // URL for the task
}
type ResponseDevicesQueryNetworkDevicesWithFiltersV1 struct {
	Response *[]ResponseDevicesQueryNetworkDevicesWithFiltersV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number of the response
}
type ResponseDevicesQueryNetworkDevicesWithFiltersV1Response struct {
	ID string `json:"id,omitempty"` // Unique identifier of the network device

	ManagementAddress string `json:"managementAddress,omitempty"` // Management address of the network device

	DNSResolvedManagementIPAddress string `json:"dnsResolvedManagementIpAddress,omitempty"` // DNS-resolved management IP address of the network device

	Hostname string `json:"hostname,omitempty"` // Hostname of the network device

	MacAddress string `json:"macAddress,omitempty"` // MAC address of the network device

	SerialNumbers []string `json:"serialNumbers,omitempty"` // Serial number of the network device. In case of stack device, there will be multiple serial numbers

	Type string `json:"type,omitempty"` // Type of the network device. This list of types can be obtained from the API intent/networkDeviceProductNames productName field.

	Family string `json:"family,omitempty"` // Product family of the network device. For example, Switches, Routers, etc

	Series string `json:"series,omitempty"` // The model range or series of the network device

	Status string `json:"status,omitempty"` // Inventory related status of the network device. Refer features for more details

	PlatformIDs []string `json:"platformIds,omitempty"` // Platform identifier of the network device

	SoftwareType string `json:"softwareType,omitempty"` // Type of software running on the network device. For example, IOS-XE, etc.

	SoftwareVersion string `json:"softwareVersion,omitempty"` // Version of the software running on the network device

	Vendor string `json:"vendor,omitempty"` // Vendor of the network device

	StackDevice *bool `json:"stackDevice,omitempty"` // Flag indicating if the network device is a stack device

	BootTime *float64 `json:"bootTime,omitempty"` // The time at which the network device was last rebooted or powered on represented as epoch in milliseconds

	Role string `json:"role,omitempty"` // Role assigned to the network device

	RoleSource string `json:"roleSource,omitempty"` // Indicates whether the network device's role was assigned automatically by the software or manually by an administrator.

	ApEthernetMacAddress string `json:"apEthernetMacAddress,omitempty"` // Ethernet MAC address of the AP network device

	ApManagerInterfaceIPAddress string `json:"apManagerInterfaceIpAddress,omitempty"` // Management IP address of the AP network device

	ApWlcIPAddress string `json:"apWlcIpAddress,omitempty"` // Management IP address of the WLC on which AP is associated to

	DeviceSupportLevel string `json:"deviceSupportLevel,omitempty"` // The level of support Catalyst Center provides for the network device.

	SNMPLocation string `json:"snmpLocation,omitempty"` // SNMP location of the network device

	SNMPContact string `json:"snmpContact,omitempty"` // SNMP contact of the network device

	ReachabilityStatus string `json:"reachabilityStatus,omitempty"` // Reachability status of the network device. Refer features for more details

	ReachabilityFailureReason string `json:"reachabilityFailureReason,omitempty"` // Reason for reachability failure. This message that provides more information about the reachability failure.

	ManagementState string `json:"managementState,omitempty"` // The status of the network device's manageability. Refer features for more details.

	LastSuccessfulResyncReasons []string `json:"lastSuccessfulResyncReasons,omitempty"` // List of reasons for the last successful resync of the device. If multiple resync requests are made before the device can start the resync, all the reasons will be captured. Possible values: ADD_DEVICE_SYNC, LINK_UP_DOWN, CONFIG_CHANGE, DEVICE_UPDATED_SYNC, AP_EVENT_BASED_SYNC, APP_REQUESTED_SYNC, PERIODIC_SYNC, UI_SYNC, CUSTOM, UNKNOWN, REFRESH_OBJECTS_FEATURE_BASED_SYNC

	ResyncStartTime *float64 `json:"resyncStartTime,omitempty"` // Start time for the last/ongoing resync represented as epoch in milliseconds

	ResyncEndTime *float64 `json:"resyncEndTime,omitempty"` // End time for the last resync represented as epoch in milliseconds

	ResyncReasons []string `json:"resyncReasons,omitempty"` // List of reasons for the ongoing/last resync on the device. If multiple resync requests were made before the resync could start, all the reasons will be captured as an array. Possible values: ADD_DEVICE_SYNC, LINK_UP_DOWN, CONFIG_CHANGE, DEVICE_UPDATED_SYNC, AP_EVENT_BASED_SYNC, APP_REQUESTED_SYNC, PERIODIC_SYNC, UI_SYNC, CUSTOM, UNKNOWN, REFRESH_OBJECTS_FEATURE_BASED_SYNC

	ResyncRequestedByApps []string `json:"resyncRequestedByApps,omitempty"` // List of applications that requested the last/ongoing resync on the device

	PendingResyncRequestCount *int `json:"pendingResyncRequestCount,omitempty"` // Number of pending resync requests for the device

	PendingResyncRequestReasons []string `json:"pendingResyncRequestReasons,omitempty"` // List of reasons for the pending resync requests. Possible values: ADD_DEVICE_SYNC, LINK_UP_DOWN, CONFIG_CHANGE, DEVICE_UPDATED_SYNC, AP_EVENT_BASED_SYNC, APP_REQUESTED_SYNC, PERIODIC_SYNC, UI_SYNC, CUSTOM, UNKNOWN, REFRESH_OBJECTS_FEATURE_BASED_SYNC

	ResyncIntervalSource string `json:"resyncIntervalSource,omitempty"` // Source of the resync interval. Note: Please refer to PUT /dna/intent/api/v1/networkDevices/resyncIntervalSettings API to update the global resync interval.

	ResyncIntervalMinutes *int `json:"resyncIntervalMinutes,omitempty"` // The duration in minutes between the periodic resync attempts for the device

	ErrorCode string `json:"errorCode,omitempty"` // Error code indicating the reason for the last resync failure

	ErrorDescription string `json:"errorDescription,omitempty"` // Additional information regarding the reason for resync failure. This is a human-readable error message and should not be expected programmatically.

	UserDefinedFields *ResponseDevicesQueryNetworkDevicesWithFiltersV1ResponseUserDefinedFields `json:"userDefinedFields,omitempty"` // Map of all user defined fields and their values associated with the device. Refer to /dna/intent/api/v1/network-device/user-defined-field API to fetch all the user defined fields.
}
type ResponseDevicesQueryNetworkDevicesWithFiltersV1ResponseUserDefinedFields interface{}
type ResponseDevicesCountTheNumberOfNetworkDevicesWithFiltersV1 struct {
	Response *ResponseDevicesCountTheNumberOfNetworkDevicesWithFiltersV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // The version of the response
}
type ResponseDevicesCountTheNumberOfNetworkDevicesWithFiltersV1Response struct {
	Count *int `json:"count,omitempty"` // The reported count.
}
type ResponseDevicesUpdateGlobalResyncIntervalV1 struct {
	Response *ResponseDevicesUpdateGlobalResyncIntervalV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version of the response
}
type ResponseDevicesUpdateGlobalResyncIntervalV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Unique identifier for the task

	URL string `json:"url,omitempty"` // URL for the task
}
type ResponseDevicesOverrideResyncIntervalV1 struct {
	Response *ResponseDevicesOverrideResyncIntervalV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version of the response
}
type ResponseDevicesOverrideResyncIntervalV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Unique identifier for the task

	URL string `json:"url,omitempty"` // URL for the task
}
type ResponseDevicesGetDetailsOfASingleNetworkDeviceV1 struct {
	Response *ResponseDevicesGetDetailsOfASingleNetworkDeviceV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number of the response
}
type ResponseDevicesGetDetailsOfASingleNetworkDeviceV1Response struct {
	ID string `json:"id,omitempty"` // Unique identifier of the network device

	ManagementAddress string `json:"managementAddress,omitempty"` // Management address of the network device

	DNSResolvedManagementIPAddress string `json:"dnsResolvedManagementIpAddress,omitempty"` // DNS-resolved management IP address of the network device

	Hostname string `json:"hostname,omitempty"` // Hostname of the network device

	MacAddress string `json:"macAddress,omitempty"` // MAC address of the network device

	SerialNumbers []string `json:"serialNumbers,omitempty"` // Serial number of the network device. In case of stack device, there will be multiple serial numbers

	Type string `json:"type,omitempty"` // Type of the network device. This list of types can be obtained from the API intent/networkDeviceProductNames productName field.

	Family string `json:"family,omitempty"` // Product family of the network device. For example, Switches, Routers, etc

	Series string `json:"series,omitempty"` // The model range or series of the network device

	Status string `json:"status,omitempty"` // Inventory related status of the network device. Refer features for more details

	PlatformIDs []string `json:"platformIds,omitempty"` // Platform identifier of the network device

	SoftwareType string `json:"softwareType,omitempty"` // Type of software running on the network device. For example, IOS-XE, etc.

	SoftwareVersion string `json:"softwareVersion,omitempty"` // Version of the software running on the network device

	Vendor string `json:"vendor,omitempty"` // Vendor of the network device

	StackDevice *bool `json:"stackDevice,omitempty"` // Flag indicating if the network device is a stack device

	BootTime *float64 `json:"bootTime,omitempty"` // The time at which the network device was last rebooted or powered on represented as epoch in milliseconds

	Role string `json:"role,omitempty"` // Role assigned to the network device

	RoleSource string `json:"roleSource,omitempty"` // Indicates whether the network device's role was assigned automatically by the software or manually by an administrator.

	ApEthernetMacAddress string `json:"apEthernetMacAddress,omitempty"` // Ethernet MAC address of the AP network device

	ApManagerInterfaceIPAddress string `json:"apManagerInterfaceIpAddress,omitempty"` // Management IP address of the AP network device

	ApWlcIPAddress string `json:"apWlcIpAddress,omitempty"` // Management IP address of the WLC on which AP is associated to

	DeviceSupportLevel string `json:"deviceSupportLevel,omitempty"` // The level of support Catalyst Center provides for the network device.

	SNMPLocation string `json:"snmpLocation,omitempty"` // SNMP location of the network device

	SNMPContact string `json:"snmpContact,omitempty"` // SNMP contact of the network device

	ReachabilityStatus string `json:"reachabilityStatus,omitempty"` // Reachability status of the network device. Refer features for more details

	ReachabilityFailureReason string `json:"reachabilityFailureReason,omitempty"` // Reason for reachability failure. This message that provides more information about the reachability failure.

	ManagementState string `json:"managementState,omitempty"` // The status of the network device's manageability. Refer features for more details.

	LastSuccessfulResyncReasons []string `json:"lastSuccessfulResyncReasons,omitempty"` // List of reasons for the last successful resync of the device. If multiple resync requests are made before the device can start the resync, all the reasons will be captured. Possible values: ADD_DEVICE_SYNC, LINK_UP_DOWN, CONFIG_CHANGE, DEVICE_UPDATED_SYNC, AP_EVENT_BASED_SYNC, APP_REQUESTED_SYNC, PERIODIC_SYNC, UI_SYNC, CUSTOM, UNKNOWN, REFRESH_OBJECTS_FEATURE_BASED_SYNC

	ResyncStartTime *float64 `json:"resyncStartTime,omitempty"` // Start time for the last/ongoing resync represented as epoch in milliseconds

	ResyncEndTime *float64 `json:"resyncEndTime,omitempty"` // End time for the last resync represented as epoch in milliseconds

	ResyncReasons []string `json:"resyncReasons,omitempty"` // List of reasons for the ongoing/last resync on the device. If multiple resync requests were made before the resync could start, all the reasons will be captured as an array. Possible values: ADD_DEVICE_SYNC, LINK_UP_DOWN, CONFIG_CHANGE, DEVICE_UPDATED_SYNC, AP_EVENT_BASED_SYNC, APP_REQUESTED_SYNC, PERIODIC_SYNC, UI_SYNC, CUSTOM, UNKNOWN, REFRESH_OBJECTS_FEATURE_BASED_SYNC

	ResyncRequestedByApps []string `json:"resyncRequestedByApps,omitempty"` // List of applications that requested the last/ongoing resync on the device

	PendingResyncRequestCount *int `json:"pendingResyncRequestCount,omitempty"` // Number of pending resync requests for the device

	PendingResyncRequestReasons []string `json:"pendingResyncRequestReasons,omitempty"` // List of reasons for the pending resync requests. Possible values: ADD_DEVICE_SYNC, LINK_UP_DOWN, CONFIG_CHANGE, DEVICE_UPDATED_SYNC, AP_EVENT_BASED_SYNC, APP_REQUESTED_SYNC, PERIODIC_SYNC, UI_SYNC, CUSTOM, UNKNOWN, REFRESH_OBJECTS_FEATURE_BASED_SYNC

	ResyncIntervalSource string `json:"resyncIntervalSource,omitempty"` // Source of the resync interval. Note: Please refer to PUT /dna/intent/api/v1/networkDevices/resyncIntervalSettings API to update the global resync interval.

	ResyncIntervalMinutes *int `json:"resyncIntervalMinutes,omitempty"` // The duration in minutes between the periodic resync attempts for the device

	ErrorCode string `json:"errorCode,omitempty"` // Error code indicating the reason for the last resync failure

	ErrorDescription string `json:"errorDescription,omitempty"` // Additional information regarding the reason for resync failure. This is a human-readable error message and should not be expected programmatically.

	UserDefinedFields *ResponseDevicesGetDetailsOfASingleNetworkDeviceV1ResponseUserDefinedFields `json:"userDefinedFields,omitempty"` // Map of all user defined fields and their values associated with the device. Refer to /dna/intent/api/v1/network-device/user-defined-field API to fetch all the user defined fields.
}
type ResponseDevicesGetDetailsOfASingleNetworkDeviceV1ResponseUserDefinedFields interface{}
type ResponseDevicesUpdateResyncIntervalForTheNetworkDeviceV1 struct {
	Response *ResponseDevicesUpdateResyncIntervalForTheNetworkDeviceV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version of the response
}
type ResponseDevicesUpdateResyncIntervalForTheNetworkDeviceV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Unique identifier for the task

	URL string `json:"url,omitempty"` // URL for the task
}
type ResponseDevicesGetResyncIntervalForTheNetworkDeviceV1 struct {
	Response *ResponseDevicesGetResyncIntervalForTheNetworkDeviceV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesGetResyncIntervalForTheNetworkDeviceV1Response struct {
	Interval *int `json:"interval,omitempty"` // Resync interval of the device
}
type ResponseDevicesRogueAdditionalDetailsV1 struct {
	Response *[]ResponseDevicesRogueAdditionalDetailsV1Response `json:"response,omitempty"` //

	TotalCount *int `json:"totalCount,omitempty"` // Total Count

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesRogueAdditionalDetailsV1Response struct {
	MacAddress string `json:"macAddress,omitempty"` // MAC Address of the Rogue BSSID

	MldMacAddress string `json:"mldMacAddress,omitempty"` // MLD MAC Address of the Rogue BSSID, this is applicable only for Wi-Fi 7 Rogues

	UpdatedTime *int `json:"updatedTime,omitempty"` // Last time when the Rogue is seen in the network

	CreatedTime *int `json:"createdTime,omitempty"` // First time when the Rogue is seen in the network

	ThreatType string `json:"threatType,omitempty"` // Type of the Rogue Threat

	ThreatLevel string `json:"threatLevel,omitempty"` // Level of the Rogue Threat

	ApName string `json:"apName,omitempty"` // Detecting AP Name

	DetectingApMac string `json:"detectingAPMac,omitempty"` // MAC Address of the Detecting AP

	SSID string `json:"ssid,omitempty"` // Rogue SSID

	Containment string `json:"containment,omitempty"` // Containment Status of the Rogue

	RadioType string `json:"radioType,omitempty"` // Radio Type on which Rogue is detected

	ControllerIP string `json:"controllerIp,omitempty"` // IP Address of the Controller detecting this Rogue

	ControllerName string `json:"controllerName,omitempty"` // Name of the Controller detecting this Rogue

	ChannelNumber string `json:"channelNumber,omitempty"` // Channel Number on which the Rogue is detected

	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Site Hierarchy of the Rogue

	Encryption string `json:"encryption,omitempty"` // Security status of the Rogue SSID

	SwitchIP string `json:"switchIp,omitempty"` // IP Address of the Switch on which the Rogue is connected. This will be filled only in case of Rogue on Wire Threat Type

	SwitchName string `json:"switchName,omitempty"` // Name of the Switch on which the Rogue is connected. This will be filled only in case of Rogue on Wire Threat Type

	PortDescription string `json:"portDescription,omitempty"` // Port information of the Switch on which the Rogue is connected. This will be filled only in case of Rogue on Wire Threat Type
}
type ResponseDevicesRogueAdditionalDetailCountV1 struct {
	Response *int `json:"response,omitempty"` // Response

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesStartWirelessRogueApContainmentV1 struct {
	Response *ResponseDevicesStartWirelessRogueApContainmentV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesStartWirelessRogueApContainmentV1Response struct {
	MacAddress string `json:"macAddress,omitempty"` // Mac Address

	Type *int `json:"type,omitempty"` // Type

	InitiatedOnWlcIP string `json:"initiatedOnWlcIp,omitempty"` // Initiated On Wlc Ip

	TaskID string `json:"taskId,omitempty"` // Task Id

	TaskType string `json:"taskType,omitempty"` // Task Type

	InitiatedOnBssid []string `json:"initiatedOnBssid,omitempty"` // Initiated On Bssid
}
type ResponseDevicesWirelessRogueApContainmentStatusV1 struct {
	Response *[]ResponseDevicesWirelessRogueApContainmentStatusV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesWirelessRogueApContainmentStatusV1Response struct {
	MacAddress string `json:"macAddress,omitempty"` // Mac Address

	Type *int `json:"type,omitempty"` // Type

	Classification string `json:"classification,omitempty"` // Classification

	ContainmentStatus string `json:"containmentStatus,omitempty"` // Containment Status

	ContainedByWlcIP []string `json:"containedByWlcIp,omitempty"` // Contained By Wlc Ip

	LastSeen *int `json:"lastSeen,omitempty"` // Last Seen

	StrongestDetectingWlcIP string `json:"strongestDetectingWlcIp,omitempty"` // Strongest Detecting Wlc Ip

	LastTaskDetail *ResponseDevicesWirelessRogueApContainmentStatusV1ResponseLastTaskDetail `json:"lastTaskDetail,omitempty"` //

	BssidContainmentStatus *[]ResponseDevicesWirelessRogueApContainmentStatusV1ResponseBssidContainmentStatus `json:"bssidContainmentStatus,omitempty"` //
}
type ResponseDevicesWirelessRogueApContainmentStatusV1ResponseLastTaskDetail struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	TaskType string `json:"taskType,omitempty"` // Task Type

	TaskState string `json:"taskState,omitempty"` // Task State

	TaskStartTime *int `json:"taskStartTime,omitempty"` // Task Start Time

	InitiatedOnWlcIP string `json:"initiatedOnWlcIp,omitempty"` // Initiated On Wlc Ip

	InitiatedOnBssid []string `json:"initiatedOnBssid,omitempty"` // Initiated On Bssid
}
type ResponseDevicesWirelessRogueApContainmentStatusV1ResponseBssidContainmentStatus struct {
	Bssid string `json:"bssid,omitempty"` // Bssid

	SSID string `json:"ssid,omitempty"` // Ssid

	RadioType string `json:"radioType,omitempty"` // Radio Type

	ContainmentStatus string `json:"containmentStatus,omitempty"` // Containment Status

	ContainedByWlcIP string `json:"containedByWlcIp,omitempty"` // Contained By Wlc Ip

	IsAdhoc *bool `json:"isAdhoc,omitempty"` // Is Adhoc
}
type ResponseDevicesStopWirelessRogueApContainmentV1 struct {
	Response *ResponseDevicesStopWirelessRogueApContainmentV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesStopWirelessRogueApContainmentV1Response struct {
	MacAddress string `json:"macAddress,omitempty"` // Mac Address

	Type *int `json:"type,omitempty"` // Type

	InitiatedOnWlcIP string `json:"initiatedOnWlcIp,omitempty"` // Initiated On Wlc Ip

	TaskID string `json:"taskId,omitempty"` // Task Id

	TaskType string `json:"taskType,omitempty"` // Task Type

	InitiatedOnBssid []string `json:"initiatedOnBssid,omitempty"` // Initiated On Bssid
}
type ResponseDevicesThreatDetailsV1 struct {
	Response *[]ResponseDevicesThreatDetailsV1Response `json:"response,omitempty"` //

	TotalCount *int `json:"totalCount,omitempty"` // Total Count

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesThreatDetailsV1Response struct {
	MacAddress string `json:"macAddress,omitempty"` // Mac Address

	UpdatedTime *int `json:"updatedTime,omitempty"` // Updated Time

	Vendor string `json:"vendor,omitempty"` // Vendor

	ThreatType string `json:"threatType,omitempty"` // Threat Type

	ThreatLevel string `json:"threatLevel,omitempty"` // Threat Level

	ApName string `json:"apName,omitempty"` // Ap Name

	DetectingApMac string `json:"detectingAPMac,omitempty"` // Detecting A P Mac

	SiteID string `json:"siteId,omitempty"` // Site Id

	Rssi string `json:"rssi,omitempty"` // Rssi

	SSID string `json:"ssid,omitempty"` // Ssid

	Containment string `json:"containment,omitempty"` // Containment

	State string `json:"state,omitempty"` // State

	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Site Name Hierarchy
}
type ResponseDevicesThreatDetailCountV1 struct {
	Response *int `json:"response,omitempty"` // Response

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesGetThreatLevelsV1 struct {
	Response *[]ResponseDevicesGetThreatLevelsV1Response `json:"response,omitempty"` //
}
type ResponseDevicesGetThreatLevelsV1Response struct {
	Name string `json:"name,omitempty"` // Name

	Value *int `json:"value,omitempty"` // Value
}
type ResponseDevicesAddAllowedMacAddressV1 struct {
	Response string `json:"response,omitempty"` // Response

	Error *ResponseDevicesAddAllowedMacAddressV1Error `json:"error,omitempty"` // Error
}
type ResponseDevicesAddAllowedMacAddressV1Error interface{}
type ResponseDevicesGetAllowedMacAddressV1 []ResponseItemDevicesGetAllowedMacAddressV1 // Array of ResponseDevicesGetAllowedMacAddressV1
type ResponseItemDevicesGetAllowedMacAddressV1 struct {
	MacAddress string `json:"macAddress,omitempty"` // Mac Address

	Category *int `json:"category,omitempty"` // Category

	LastModified *int `json:"lastModified,omitempty"` // Last Modified
}
type ResponseDevicesGetAllowedMacAddressCountV1 struct {
	Response *int `json:"response,omitempty"` // Response

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesRemoveAllowedMacAddressV1 struct {
	Response string `json:"response,omitempty"` // Response

	Error *ResponseDevicesRemoveAllowedMacAddressV1Error `json:"error,omitempty"` // Error
}
type ResponseDevicesRemoveAllowedMacAddressV1Error interface{}
type ResponseDevicesThreatSummaryV1 struct {
	Response *[]ResponseDevicesThreatSummaryV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesThreatSummaryV1Response struct {
	Timestamp *int `json:"timestamp,omitempty"` // Timestamp

	ThreatData *[]ResponseDevicesThreatSummaryV1ResponseThreatData `json:"threatData,omitempty"` //
}
type ResponseDevicesThreatSummaryV1ResponseThreatData struct {
	ThreatType string `json:"threatType,omitempty"` // Threat Type

	ThreatLevel string `json:"threatLevel,omitempty"` // Threat Level

	ThreatCount *int `json:"threatCount,omitempty"` // Threat Count
}
type ResponseDevicesGetThreatTypesV1 struct {
	Response *[]ResponseDevicesGetThreatTypesV1Response `json:"response,omitempty"` //
}
type ResponseDevicesGetThreatTypesV1Response struct {
	Value *int `json:"value,omitempty"` // Value

	Name string `json:"name,omitempty"` // Name

	Label string `json:"label,omitempty"` // Label

	IsCustom *bool `json:"isCustom,omitempty"` // Is Custom

	IsDeleted *bool `json:"isDeleted,omitempty"` // Is Deleted
}
type ResponseDevicesGetDeviceInterfaceStatsInfoV2 struct {
	Version string `json:"version,omitempty"` // Version

	TotalCount *float64 `json:"totalCount,omitempty"` // The total count

	Response *[]ResponseDevicesGetDeviceInterfaceStatsInfoV2Response `json:"response,omitempty"` //

	Page *ResponseDevicesGetDeviceInterfaceStatsInfoV2Page `json:"page,omitempty"` //
}
type ResponseDevicesGetDeviceInterfaceStatsInfoV2Response struct {
	ID string `json:"id,omitempty"` // Interface Instance Id

	Values *ResponseDevicesGetDeviceInterfaceStatsInfoV2ResponseValues `json:"values,omitempty"` //
}
type ResponseDevicesGetDeviceInterfaceStatsInfoV2ResponseValues struct {
	AdminStatus string `json:"adminStatus,omitempty"` // The desired state of the interface

	DeviceID string `json:"deviceId,omitempty"` // Device Id

	DuplexConfig string `json:"duplexConfig,omitempty"` // Interface duplex config status

	DuplexOper string `json:"duplexOper,omitempty"` // Interface duplex operational status

	InterfaceID string `json:"interfaceId,omitempty"` // Interface ifIndex

	InterfaceType string `json:"interfaceType,omitempty"` // Physical or Virtual type

	InstanceID string `json:"instanceId,omitempty"` // Interface InstanceId

	IPv4Address string `json:"ipv4Address,omitempty"` // Interface IPV4 Address

	IPv6AddressList []string `json:"ipv6AddressList,omitempty"` // List of interface IPV6 Address

	IsL3Interface string `json:"isL3Interface,omitempty"` // Interface is L3 or not

	IsWan string `json:"isWan,omitempty"` // nterface is WAN link or not

	MacAddr string `json:"macAddr,omitempty"` // Interface MAC Address

	MediaType string `json:"mediaType,omitempty"` // Interface media type

	Name string `json:"name,omitempty"` // Name of the interface

	OperStatus string `json:"operStatus,omitempty"` // Interface operational status

	PeerStackMember string `json:"peerStackMember,omitempty"` // Interface peer stack member Id

	PeerStackPort string `json:"peerStackPort,omitempty"` // Interface peer stack member port

	PortChannelID string `json:"portChannelId,omitempty"` // Interface Port-Channel Id

	PortMode string `json:"portMode,omitempty"` // Interface Port Mode

	PortType string `json:"portType,omitempty"` // Interface ifType

	Description string `json:"description,omitempty"` // Interface description

	RxDiscards string `json:"rxDiscards,omitempty"` // Rx Discards in %

	RxError string `json:"rxError,omitempty"` // Rx Errors in %

	RxRate string `json:"rxRate,omitempty"` // Rx rate in bps

	RxUtilization string `json:"rxUtilization,omitempty"` // Rx Utilization in %

	Speed string `json:"speed,omitempty"` // Speed of the Interface in kbps

	StackPortType string `json:"stackPortType,omitempty"` // Interface stack port type. SVL or DAD

	Timestamp string `json:"timestamp,omitempty"` // Interface stats collected timestamp

	TxDiscards string `json:"txDiscards,omitempty"` // Tx Discards in %

	TxError string `json:"txError,omitempty"` // Tx Errors in %

	TxRate string `json:"txRate,omitempty"` // Tx Rate in bps

	TxUtilization string `json:"txUtilization,omitempty"` // Tx  Utilization in %

	VLANID string `json:"vlanId,omitempty"` // Interface VLAN Id
}
type ResponseDevicesGetDeviceInterfaceStatsInfoV2Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *float64 `json:"offset,omitempty"` // Offset

	Count *int `json:"count,omitempty"` // Count
}
type RequestDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1 struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	Filters *[]RequestDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1Filters `json:"filters,omitempty"` //

	Page *RequestDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1Page `json:"page,omitempty"` //
}
type RequestDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1Filters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	Value []string `json:"value,omitempty"` // Value
}
type RequestDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	SortBy *[]RequestDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1PageSortBy `json:"sortBy,omitempty"` //
}
type RequestDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1PageSortBy struct {
	Name string `json:"name,omitempty"` // Name

	Order string `json:"order,omitempty"` // Order
}
type RequestDevicesRetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersV1 struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	Filters *[]RequestDevicesRetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersV1Filters `json:"filters,omitempty"` //
}
type RequestDevicesRetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersV1Filters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	Value []string `json:"value,omitempty"` // Value
}
type RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1 struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	GroupBy []string `json:"groupBy,omitempty"` // Group By

	Filters *[]RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Filters `json:"filters,omitempty"` //

	Attributes []string `json:"attributes,omitempty"` // Attributes

	AggregateAttributes *[]RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1AggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Page `json:"page,omitempty"` //
}
type RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Filters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	LogicalOperator string `json:"logicalOperator,omitempty"` // Logical Operator

	Value *RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1FiltersValue `json:"value,omitempty"` // Value

	Filters []string `json:"filters,omitempty"` // Filters
}
type RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1FiltersValue interface{}
type RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1AggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	SortBy *[]RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1PageSortBy `json:"sortBy,omitempty"` //
}
type RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1PageSortBy struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function

	Order string `json:"order,omitempty"` // Order
}
type RequestDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1 struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	TopN *int `json:"topN,omitempty"` // Top N

	GroupBy []string `json:"groupBy,omitempty"` // Group By

	Filters *[]RequestDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Filters `json:"filters,omitempty"` //

	Attributes []string `json:"attributes,omitempty"` // Attributes

	AggregateAttributes *[]RequestDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1AggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Page `json:"page,omitempty"` //
}
type RequestDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Filters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	LogicalOperator string `json:"logicalOperator,omitempty"` // Logical Operator

	Value *RequestDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1FiltersValue `json:"value,omitempty"` // Value

	Filters []string `json:"filters,omitempty"` // Filters
}
type RequestDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1FiltersValue interface{}
type RequestDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1AggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	SortBy *[]RequestDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1PageSortBy `json:"sortBy,omitempty"` //
}
type RequestDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1PageSortBy struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function

	Order string `json:"order,omitempty"` // Order
}
type RequestDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1 struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	TrendInterval string `json:"trendInterval,omitempty"` // Trend Interval

	GroupBy []string `json:"groupBy,omitempty"` // Group By

	Filters *[]RequestDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Filters `json:"filters,omitempty"` //

	Attributes []string `json:"attributes,omitempty"` // Attributes

	AggregateAttributes *[]RequestDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1AggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Page `json:"page,omitempty"` //
}
type RequestDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Filters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	LogicalOperator string `json:"logicalOperator,omitempty"` // Logical Operator

	Value *RequestDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1FiltersValue `json:"value,omitempty"` // Value

	Filters []string `json:"filters,omitempty"` // Filters
}
type RequestDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1FiltersValue interface{}
type RequestDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1AggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type RequestDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1 struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	TrendInterval string `json:"trendInterval,omitempty"` // Trend Interval

	GroupBy []string `json:"groupBy,omitempty"` // Group By

	Filters *[]RequestDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1Filters `json:"filters,omitempty"` //

	Attributes []string `json:"attributes,omitempty"` // Attributes

	AggregateAttributes *[]RequestDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1AggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1Page `json:"page,omitempty"` //
}
type RequestDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1Filters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	LogicalOperator string `json:"logicalOperator,omitempty"` // Logical Operator

	Value *RequestDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1FiltersValue `json:"value,omitempty"` // Value

	Filters []string `json:"filters,omitempty"` // Filters
}
type RequestDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1FiltersValue interface{}
type RequestDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1AggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type RequestDevicesQueryAssuranceEventsWithFiltersV1 struct {
	DeviceFamily []string `json:"deviceFamily,omitempty"` // Device Family

	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	Attributes []string `json:"attributes,omitempty"` // Attributes

	Views []string `json:"views,omitempty"` // Views

	Filters *[]RequestDevicesQueryAssuranceEventsWithFiltersV1Filters `json:"filters,omitempty"` //

	Page *RequestDevicesQueryAssuranceEventsWithFiltersV1Page `json:"page,omitempty"` //
}
type RequestDevicesQueryAssuranceEventsWithFiltersV1Filters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	Value string `json:"value,omitempty"` // Value
}
type RequestDevicesQueryAssuranceEventsWithFiltersV1Page struct {
	Offset *int `json:"offset,omitempty"` // Offset

	Limit *int `json:"limit,omitempty"` // Limit

	SortBy *[]RequestDevicesQueryAssuranceEventsWithFiltersV1PageSortBy `json:"sortBy,omitempty"` //
}
type RequestDevicesQueryAssuranceEventsWithFiltersV1PageSortBy struct {
	Name string `json:"name,omitempty"` // Name

	Order string `json:"order,omitempty"` // Order
}
type RequestDevicesCountTheNumberOfEventsWithFiltersV1 struct {
	DeviceFamily []string `json:"deviceFamily,omitempty"` // Device Family

	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	Filters *[]RequestDevicesCountTheNumberOfEventsWithFiltersV1Filters `json:"filters,omitempty"` //
}
type RequestDevicesCountTheNumberOfEventsWithFiltersV1Filters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	Value string `json:"value,omitempty"` // Value
}
type RequestDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1 struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	Filters *[]RequestDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1Filters `json:"filters,omitempty"` //

	Page *RequestDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1Page `json:"page,omitempty"` //
}
type RequestDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1Filters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	Value []string `json:"value,omitempty"` // Value
}
type RequestDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	SortBy *[]RequestDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1PageSortBy `json:"sortBy,omitempty"` //
}
type RequestDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1PageSortBy struct {
	Name string `json:"name,omitempty"` // Name

	Order string `json:"order,omitempty"` // Order
}
type RequestDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersV1 struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	Filters *[]RequestDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersV1Filters `json:"filters,omitempty"` //
}
type RequestDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersV1Filters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	Value []string `json:"value,omitempty"` // Value
}
type RequestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1 struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	GroupBy []string `json:"groupBy,omitempty"` // Group By

	Filters *[]RequestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1Filters `json:"filters,omitempty"` //

	Attributes []string `json:"attributes,omitempty"` // Attributes

	AggregateAttributes *[]RequestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1AggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1Page `json:"page,omitempty"` //
}
type RequestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1Filters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	LogicalOperator string `json:"logicalOperator,omitempty"` // Logical Operator

	Value *RequestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1FiltersValue `json:"value,omitempty"` // Value

	Filters *[]RequestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1FiltersFilters `json:"filters,omitempty"` //
}
type RequestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1FiltersValue interface{}
type RequestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1FiltersFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	LogicalOperator string `json:"logicalOperator,omitempty"` // Logical Operator

	Value *RequestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1FiltersFiltersValue `json:"value,omitempty"` // Value

	Filters []string `json:"filters,omitempty"` // Filters
}
type RequestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1FiltersFiltersValue interface{}
type RequestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1AggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	SortBy *[]RequestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1PageSortBy `json:"sortBy,omitempty"` //
}
type RequestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1PageSortBy struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function

	Order string `json:"order,omitempty"` // Order
}
type RequestDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1 struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	TopN *int `json:"topN,omitempty"` // Top N

	GroupBy []string `json:"groupBy,omitempty"` // Group By

	Filters *[]RequestDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1Filters `json:"filters,omitempty"` //

	Attributes []string `json:"attributes,omitempty"` // Attributes

	AggregateAttributes *[]RequestDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1AggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1Page `json:"page,omitempty"` //
}
type RequestDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1Filters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	LogicalOperator string `json:"logicalOperator,omitempty"` // Logical Operator

	Value *RequestDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1FiltersValue `json:"value,omitempty"` // Value

	Filters []string `json:"filters,omitempty"` // Filters
}
type RequestDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1FiltersValue interface{}
type RequestDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1AggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	SortBy *[]RequestDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1PageSortBy `json:"sortBy,omitempty"` //
}
type RequestDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1PageSortBy struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function

	Order string `json:"order,omitempty"` // Order
}
type RequestDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1 struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	TrendInterval string `json:"trendInterval,omitempty"` // Trend Interval

	GroupBy []string `json:"groupBy,omitempty"` // Group By

	Filters *[]RequestDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1Filters `json:"filters,omitempty"` //

	Attributes []string `json:"attributes,omitempty"` // Attributes

	AggregateAttributes *[]RequestDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1AggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1Page `json:"page,omitempty"` //
}
type RequestDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1Filters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	LogicalOperator string `json:"logicalOperator,omitempty"` // Logical Operator

	Value *RequestDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1FiltersValue `json:"value,omitempty"` // Value

	Filters []string `json:"filters,omitempty"` // Filters
}
type RequestDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1FiltersValue interface{}
type RequestDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1AggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type RequestDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1 struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	TrendInterval string `json:"trendInterval,omitempty"` // Trend Interval

	GroupBy []string `json:"groupBy,omitempty"` // Group By

	Filters *[]RequestDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1Filters `json:"filters,omitempty"` //

	Attributes []string `json:"attributes,omitempty"` // Attributes

	AggregateAttributes *[]RequestDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1AggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1Page `json:"page,omitempty"` //
}
type RequestDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1Filters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	LogicalOperator string `json:"logicalOperator,omitempty"` // Logical Operator

	Value *RequestDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1FiltersValue `json:"value,omitempty"` // Value

	Filters []string `json:"filters,omitempty"` // Filters
}
type RequestDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1FiltersValue interface{}
type RequestDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1AggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type RequestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1 struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	Filters *[]RequestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1Filters `json:"filters,omitempty"` //

	Page *RequestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1Page `json:"page,omitempty"` //
}
type RequestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1Filters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	Value []string `json:"value,omitempty"` // Value
}
type RequestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	SortBy *[]RequestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1PageSortBy `json:"sortBy,omitempty"` //
}
type RequestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1PageSortBy struct {
	Name string `json:"name,omitempty"` // Name

	Order string `json:"order,omitempty"` // Order
}
type RequestDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1 struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	Filters *[]RequestDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1Filters `json:"filters,omitempty"` //
}
type RequestDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1Filters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	Value []string `json:"value,omitempty"` // Value
}
type RequestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1 struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	GroupBy []string `json:"groupBy,omitempty"` // Group By

	Filters *[]RequestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1Filters `json:"filters,omitempty"` //

	Attributes []string `json:"attributes,omitempty"` // Attributes

	AggregateAttributes *[]RequestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1AggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1Page `json:"page,omitempty"` //
}
type RequestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1Filters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	LogicalOperator string `json:"logicalOperator,omitempty"` // Logical Operator

	Value *RequestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1FiltersValue `json:"value,omitempty"` // Value

	Filters *[]RequestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1FiltersFilters `json:"filters,omitempty"` //
}
type RequestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1FiltersValue interface{}
type RequestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1FiltersFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	LogicalOperator string `json:"logicalOperator,omitempty"` // Logical Operator

	Value *RequestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1FiltersFiltersValue `json:"value,omitempty"` // Value

	Filters []string `json:"filters,omitempty"` // Filters
}
type RequestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1FiltersFiltersValue interface{}
type RequestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1AggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	SortBy *[]RequestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1PageSortBy `json:"sortBy,omitempty"` //
}
type RequestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1PageSortBy struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function

	Order string `json:"order,omitempty"` // Order
}
type RequestDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1 struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	TopN *int `json:"topN,omitempty"` // Top N

	GroupBy []string `json:"groupBy,omitempty"` // Group By

	Filters *[]RequestDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1Filters `json:"filters,omitempty"` //

	Attributes []string `json:"attributes,omitempty"` // Attributes

	AggregateAttributes *[]RequestDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1AggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1Page `json:"page,omitempty"` //
}
type RequestDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1Filters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	LogicalOperator string `json:"logicalOperator,omitempty"` // Logical Operator

	Value *RequestDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1FiltersValue `json:"value,omitempty"` // Value

	Filters []string `json:"filters,omitempty"` // Filters
}
type RequestDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1FiltersValue interface{}
type RequestDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1AggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	SortBy *[]RequestDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1PageSortBy `json:"sortBy,omitempty"` //
}
type RequestDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1PageSortBy struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function

	Order string `json:"order,omitempty"` // Order
}
type RequestDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1 struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	TrendInterval string `json:"trendInterval,omitempty"` // Trend Interval

	GroupBy []string `json:"groupBy,omitempty"` // Group By

	Filters *[]RequestDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1Filters `json:"filters,omitempty"` //

	Attributes []string `json:"attributes,omitempty"` // Attributes

	AggregateAttributes *[]RequestDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1AggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1Page `json:"page,omitempty"` //
}
type RequestDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1Filters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	LogicalOperator string `json:"logicalOperator,omitempty"` // Logical Operator

	Value *RequestDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1FiltersValue `json:"value,omitempty"` // Value

	Filters []string `json:"filters,omitempty"` // Filters
}
type RequestDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1FiltersValue interface{}
type RequestDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1AggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type RequestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1 struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	TrendInterval string `json:"trendInterval,omitempty"` // Trend Interval

	GroupBy []string `json:"groupBy,omitempty"` // Group By

	Filters *[]RequestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1Filters `json:"filters,omitempty"` //

	Attributes []string `json:"attributes,omitempty"` // Attributes

	AggregateAttributes *[]RequestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1AggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1Page `json:"page,omitempty"` //
}
type RequestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1Filters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	LogicalOperator string `json:"logicalOperator,omitempty"` // Logical Operator

	Value *RequestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1FiltersValue `json:"value,omitempty"` // Value

	Filters []string `json:"filters,omitempty"` // Filters
}
type RequestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1FiltersValue interface{}
type RequestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1AggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1 struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	Views []string `json:"views,omitempty"` // Views

	Attributes []string `json:"attributes,omitempty"` // Attributes

	Filters *[]RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1Filters `json:"filters,omitempty"` //

	AggregateAttributes *[]RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1AggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1Page `json:"page,omitempty"` //
}
type RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1Filters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	LogicalOperator string `json:"logicalOperator,omitempty"` // Logical Operator

	Value *RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1FiltersValue `json:"value,omitempty"` // Value

	Filters *[]RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1FiltersFilters `json:"filters,omitempty"` //
}
type RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1FiltersValue interface{}
type RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1FiltersFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	LogicalOperator string `json:"logicalOperator,omitempty"` // Logical Operator

	Value *RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1FiltersFiltersValue `json:"value,omitempty"` // Value

	Filters []string `json:"filters,omitempty"` // Filters
}
type RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1FiltersFiltersValue interface{}
type RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1AggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	SortBy *[]RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1PageSortBy `json:"sortBy,omitempty"` //
}
type RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1PageSortBy struct {
	Name string `json:"name,omitempty"` // Name

	Order string `json:"order,omitempty"` // Order
}
type RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesV1 struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	Views []string `json:"views,omitempty"` // Views

	Attributes []string `json:"attributes,omitempty"` // Attributes

	Filters *[]RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesV1Filters `json:"filters,omitempty"` //

	AggregateAttributes *[]RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesV1AggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesV1Page `json:"page,omitempty"` //
}
type RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesV1Filters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	LogicalOperator string `json:"logicalOperator,omitempty"` // Logical Operator

	Value *RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesV1FiltersValue `json:"value,omitempty"` // Value

	Filters []string `json:"filters,omitempty"` // Filters
}
type RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesV1FiltersValue interface{}
type RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesV1AggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	SortBy *[]RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesV1PageSortBy `json:"sortBy,omitempty"` //
}
type RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesV1PageSortBy struct {
	Name string `json:"name,omitempty"` // Name

	Order string `json:"order,omitempty"` // Order
}
type RequestDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeV1 struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	TrendIntervalInMinutes *int `json:"trendIntervalInMinutes,omitempty"` // Trend Interval In Minutes

	Attributes []string `json:"attributes,omitempty"` // Attributes

	Filters *[]RequestDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeV1Filters `json:"filters,omitempty"` //

	AggregateAttributes *[]RequestDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeV1AggregateAttributes `json:"aggregateAttributes,omitempty"` //

	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type RequestDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeV1Filters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	Value string `json:"value,omitempty"` // Value
}
type RequestDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeV1AggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1 struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	Views []string `json:"views,omitempty"` // Views

	Attributes []string `json:"attributes,omitempty"` // Attributes

	Filters *[]RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1Filters `json:"filters,omitempty"` //

	AggregateAttributes *[]RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1AggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1Page `json:"page,omitempty"` //
}
type RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1Filters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	Value string `json:"value,omitempty"` // Value
}
type RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1AggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	Count *int `json:"count,omitempty"` // Count

	SortBy string `json:"sortBy,omitempty"` // Sort By
}
type RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1 struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	Filters *[]RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1Filters `json:"filters,omitempty"` //

	Views []string `json:"views,omitempty"`

	Attributes []string `json:"attributes,omitempty"`

	AggregateAttributes []RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1AggregateAttributes `json:"aggregateAttributes,omitempty"`

	Page *RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1Page `json:"page,omitempty"`
}
type RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1Page struct {
	Limit  *int `json:"limit,omitempty"`
	Offset *int `json:"offset,omitempty"`
	SortBy *[]RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1PageSortBy
}
type RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1PageSortBy struct {
	Name  string `json:"name,omitempty"`
	Order string `json:"order,omitempty"`
}

type RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1AggregateAttributes struct {
	Name     string `json:"name,omitempty"`
	Function string `json:"function,omitempty"`
}
type RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1Filters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	Value string `json:"value,omitempty"` // Value
}
type RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1 struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	GroupBy []string `json:"groupBy,omitempty"` // Group By

	Attributes []string `json:"attributes,omitempty"` // Attributes

	Filters *[]RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1Filters `json:"filters,omitempty"` //

	AggregateAttributes *[]RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1AggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1Page `json:"page,omitempty"` //
}
type RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1Filters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	Value string `json:"value,omitempty"` // Value
}
type RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1AggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	SortBy *[]RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1PageSortBy `json:"sortBy,omitempty"` //
}
type RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1PageSortBy struct {
	Name string `json:"name,omitempty"` // Name

	Order string `json:"order,omitempty"` // Order
}
type RequestDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1 struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	TopN *int `json:"topN,omitempty"` // Top N

	GroupBy []string `json:"groupBy,omitempty"` // Group By

	Attributes *[]RequestDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1Attributes `json:"attributes,omitempty"` // Attributes

	Filters *[]RequestDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1Filters `json:"filters,omitempty"` //

	AggregateAttributes *[]RequestDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1AggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1Page `json:"page,omitempty"` //
}
type RequestDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1Attributes interface{}
type RequestDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1Filters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	Value string `json:"value,omitempty"` // Value
}
type RequestDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1AggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	SortBy *[]RequestDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1PageSortBy `json:"sortBy,omitempty"` //
}
type RequestDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1PageSortBy struct {
	Name string `json:"name,omitempty"` // Name

	Order string `json:"order,omitempty"` // Order

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetsTheTrendAnalyticsDataV1 struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	TrendInterval string `json:"trendInterval,omitempty"` // Trend Interval

	GroupBy *[]RequestDevicesGetsTheTrendAnalyticsDataV1GroupBy `json:"groupBy,omitempty"` // Group By

	Attributes []string `json:"attributes,omitempty"` // Attributes

	Filters *[]RequestDevicesGetsTheTrendAnalyticsDataV1Filters `json:"filters,omitempty"` //

	AggregateAttributes *[]RequestDevicesGetsTheTrendAnalyticsDataV1AggregateAttributes `json:"aggregateAttributes,omitempty"` // Aggregate Attributes

	Page *RequestDevicesGetsTheTrendAnalyticsDataV1Page `json:"page,omitempty"` //
}
type RequestDevicesGetsTheTrendAnalyticsDataV1GroupBy interface{}
type RequestDevicesGetsTheTrendAnalyticsDataV1Filters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	Value string `json:"value,omitempty"` // Value
}
type RequestDevicesGetsTheTrendAnalyticsDataV1AggregateAttributes interface{}
type RequestDevicesGetsTheTrendAnalyticsDataV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1 struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	TrendIntervalInMinutes *int `json:"trendIntervalInMinutes,omitempty"` // Trend Interval In Minutes

	GroupBy []string `json:"groupBy,omitempty"` // Group By

	Filters *[]RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1Filters `json:"filters,omitempty"` //

	Attributes []string `json:"attributes,omitempty"` // Attributes

	AggregateAttributes *[]RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1AggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1Page `json:"page,omitempty"` //
}
type RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1Filters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	LogicalOperator string `json:"logicalOperator,omitempty"` // Logical Operator

	Value *RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1FiltersValue `json:"value,omitempty"` // Value

	Filters []string `json:"filters,omitempty"` // Filters
}
type RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1FiltersValue interface{}
type RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1AggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type RequestDevicesUpdatePlannedAccessPointForFloorV1 struct {
	Attributes *RequestDevicesUpdatePlannedAccessPointForFloorV1Attributes `json:"attributes,omitempty"` //

	IsSensor *bool `json:"isSensor,omitempty"` // Indicates that PAP is a sensor

	Location *RequestDevicesUpdatePlannedAccessPointForFloorV1Location `json:"location,omitempty"` //

	Position *RequestDevicesUpdatePlannedAccessPointForFloorV1Position `json:"position,omitempty"` //

	RadioCount *int `json:"radioCount,omitempty"` // Number of radios of the planned access point

	Radios *[]RequestDevicesUpdatePlannedAccessPointForFloorV1Radios `json:"radios,omitempty"` //
}
type RequestDevicesUpdatePlannedAccessPointForFloorV1Attributes struct {
	CreateDate *float64 `json:"createDate,omitempty"` // Created date of the planned access point

	Domain string `json:"domain,omitempty"` // Service domain to which the planned access point belongs

	HeirarchyName string `json:"heirarchyName,omitempty"` // Hierarchy name of the planned access point

	ID *float64 `json:"id,omitempty"` // Unique id of the planned access point

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance uuid of the planned access point

	MacAddress string `json:"macAddress,omitempty"` // MAC address of the planned access point

	Name string `json:"name,omitempty"` // Display name of the planned access point

	Source string `json:"source,omitempty"` // Source of the data used to create the planned access point

	TypeString string `json:"typeString,omitempty"` // Type string representation of the planned access point
}
type RequestDevicesUpdatePlannedAccessPointForFloorV1Location struct {
	Altitude *float64 `json:"altitude,omitempty"` // Altitude of the planned access point's location

	Lattitude *float64 `json:"lattitude,omitempty"` // Latitude of the planned access point's location

	Longtitude *float64 `json:"longtitude,omitempty"` // Longitude of the planned access point's location
}
type RequestDevicesUpdatePlannedAccessPointForFloorV1Position struct {
	X *float64 `json:"x,omitempty"` // x-coordinate of the planned access point on the map, 0,0 point being the top-left corner

	Y *float64 `json:"y,omitempty"` // y-coordinate of the planned access point on the map, 0,0 point being the top-left corner

	Z *float64 `json:"z,omitempty"` // z-coordinate, or height, of the planned access point on the map
}
type RequestDevicesUpdatePlannedAccessPointForFloorV1Radios struct {
	Antenna *RequestDevicesUpdatePlannedAccessPointForFloorV1RadiosAntenna `json:"antenna,omitempty"` //

	Attributes *RequestDevicesUpdatePlannedAccessPointForFloorV1RadiosAttributes `json:"attributes,omitempty"` //

	IsSensor *bool `json:"isSensor,omitempty"` // Determines if it is sensor or not
}
type RequestDevicesUpdatePlannedAccessPointForFloorV1RadiosAntenna struct {
	AzimuthAngle *float64 `json:"azimuthAngle,omitempty"` // Azimuth angle of the antenna

	ElevationAngle *float64 `json:"elevationAngle,omitempty"` // Elevation angle of the antenna

	Gain *float64 `json:"gain,omitempty"` // Gain of the antenna

	Mode string `json:"mode,omitempty"` // Mode of the antenna associated with this radio

	Name string `json:"name,omitempty"` // Name of the antenna

	Type string `json:"type,omitempty"` // Type of the antenna associated with this radio
}
type RequestDevicesUpdatePlannedAccessPointForFloorV1RadiosAttributes struct {
	Channel *float64 `json:"channel,omitempty"` // Channel in which this radio operates

	ChannelString string `json:"channelString,omitempty"` // Channel string representation

	ID *int `json:"id,omitempty"` // Id of the radio

	IfMode string `json:"ifMode,omitempty"` // IF mode of the radio

	IfTypeString string `json:"ifTypeString,omitempty"` // String representation of native band

	IfTypeSubband string `json:"ifTypeSubband,omitempty"` // Sub band of the radio

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid of the radio

	SlotID *float64 `json:"slotId,omitempty"` // Slot number in which the radio resides in the parent access point

	TxPowerLevel *float64 `json:"txPowerLevel,omitempty"` // Tx Power at which this radio operates (in dBm)
}
type RequestDevicesCreatePlannedAccessPointForFloorV1 struct {
	Attributes *RequestDevicesCreatePlannedAccessPointForFloorV1Attributes `json:"attributes,omitempty"` //

	IsSensor *bool `json:"isSensor,omitempty"` // Indicates that PAP is a sensor

	Location *RequestDevicesCreatePlannedAccessPointForFloorV1Location `json:"location,omitempty"` //

	Position *RequestDevicesCreatePlannedAccessPointForFloorV1Position `json:"position,omitempty"` //

	RadioCount *int `json:"radioCount,omitempty"` // Number of radios of the planned access point

	Radios *[]RequestDevicesCreatePlannedAccessPointForFloorV1Radios `json:"radios,omitempty"` //
}
type RequestDevicesCreatePlannedAccessPointForFloorV1Attributes struct {
	CreateDate *float64 `json:"createDate,omitempty"` // Created date of the planned access point

	Domain string `json:"domain,omitempty"` // Service domain to which the planned access point belongs

	HeirarchyName string `json:"heirarchyName,omitempty"` // Hierarchy name of the planned access point

	ID *float64 `json:"id,omitempty"` // Unique id of the planned access point

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance uuid of the planned access point

	MacAddress string `json:"macAddress,omitempty"` // MAC address of the planned access point

	Name string `json:"name,omitempty"` // Display name of the planned access point

	Source string `json:"source,omitempty"` // Source of the data used to create the planned access point

	TypeString string `json:"typeString,omitempty"` // Type string representation of the planned access point
}
type RequestDevicesCreatePlannedAccessPointForFloorV1Location struct {
	Altitude *float64 `json:"altitude,omitempty"` // Altitude of the planned access point's location

	Lattitude *float64 `json:"lattitude,omitempty"` // Latitude of the planned access point's location

	Longtitude *float64 `json:"longtitude,omitempty"` // Longitude of the planned access point's location
}
type RequestDevicesCreatePlannedAccessPointForFloorV1Position struct {
	X *float64 `json:"x,omitempty"` // x-coordinate of the planned access point on the map, 0,0 point being the top-left corner

	Y *float64 `json:"y,omitempty"` // y-coordinate of the planned access point on the map, 0,0 point being the top-left corner

	Z *float64 `json:"z,omitempty"` // z-coordinate, or height, of the planned access point on the map
}
type RequestDevicesCreatePlannedAccessPointForFloorV1Radios struct {
	Antenna *RequestDevicesCreatePlannedAccessPointForFloorV1RadiosAntenna `json:"antenna,omitempty"` //

	Attributes *RequestDevicesCreatePlannedAccessPointForFloorV1RadiosAttributes `json:"attributes,omitempty"` //

	IsSensor *bool `json:"isSensor,omitempty"` // Determines if it is sensor or not
}
type RequestDevicesCreatePlannedAccessPointForFloorV1RadiosAntenna struct {
	AzimuthAngle *float64 `json:"azimuthAngle,omitempty"` // Azimuth angle of the antenna

	ElevationAngle *float64 `json:"elevationAngle,omitempty"` // Elevation angle of the antenna

	Gain *float64 `json:"gain,omitempty"` // Gain of the antenna

	Mode string `json:"mode,omitempty"` // Mode of the antenna associated with this radio

	Name string `json:"name,omitempty"` // Name of the antenna

	Type string `json:"type,omitempty"` // Type of the antenna associated with this radio
}
type RequestDevicesCreatePlannedAccessPointForFloorV1RadiosAttributes struct {
	Channel *float64 `json:"channel,omitempty"` // Channel in which this radio operates

	ChannelString string `json:"channelString,omitempty"` // Channel string representation

	ID *int `json:"id,omitempty"` // Id of the radio

	IfMode string `json:"ifMode,omitempty"` // IF mode of the radio

	IfTypeString string `json:"ifTypeString,omitempty"` // String representation of native band

	IfTypeSubband string `json:"ifTypeSubband,omitempty"` // Sub band of the radio

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid of the radio

	SlotID *float64 `json:"slotId,omitempty"` // Slot number in which the radio resides in the parent access point

	TxPowerLevel *float64 `json:"txPowerLevel,omitempty"` // Tx Power at which this radio operates (in dBm)
}
type RequestDevicesUpdateHealthScoreDefinitionsV1 []RequestItemDevicesUpdateHealthScoreDefinitionsV1 // Array of RequestDevicesUpdateHealthScoreDefinitionsV1
type RequestItemDevicesUpdateHealthScoreDefinitionsV1 struct {
	ID string `json:"id,omitempty"` // Id

	IncludeForOverallHealth *bool `json:"includeForOverallHealth,omitempty"` // Include For Overall Health

	ThresholdValue *float64 `json:"thresholdValue,omitempty"` // Threshold Value

	SynchronizeToIssueThreshold *bool `json:"synchronizeToIssueThreshold,omitempty"` // Synchronize To Issue Threshold
}
type RequestDevicesUpdateHealthScoreDefinitionForTheGivenIDV1 struct {
	IncludeForOverallHealth *bool `json:"includeForOverallHealth,omitempty"` // Include For Overall Health

	ThresholdValue *float64 `json:"thresholdValue,omitempty"` // Thresehold Value

	SynchronizeToIssueThreshold *bool `json:"synchronizeToIssueThreshold,omitempty"` // Synchronize To Issue Threshold
}
type RequestDevicesUpdateInterfaceDetailsV1 struct {
	Description string `json:"description,omitempty"` // Description for the Interface

	AdminStatus string `json:"adminStatus,omitempty"` // Admin status as ('UP'/'DOWN')

	VLANID *int `json:"vlanId,omitempty"` // VLAN Id to be Updated

	VoiceVLANID *int `json:"voiceVlanId,omitempty"` // Voice Vlan Id to be Updated
}
type RequestDevicesClearMacAddressTableV1 struct {
	Operation string `json:"operation,omitempty"` // Operation needs to be specified as 'ClearMacAddress'.

	Payload *RequestDevicesClearMacAddressTableV1Payload `json:"payload,omitempty"` // Payload is not applicable
}
type RequestDevicesClearMacAddressTableV1Payload interface{}
type RequestDevicesAddDeviceKnowYourNetworkV1 struct {
	CliTransport string `json:"cliTransport,omitempty"` // CLI transport. Supported values: telnet, ssh. Required if type is NETWORK_DEVICE.

	ComputeDevice *bool `json:"computeDevice,omitempty"` // Compute Device or not. Options are true / false.

	EnablePassword string `json:"enablePassword,omitempty"` // CLI enable password of the device. Required if device is configured to use enable password.

	ExtendedDiscoveryInfo string `json:"extendedDiscoveryInfo,omitempty"` // This field holds that info as whether to add device with canned data or not. Supported values: DISCOVER_WITH_CANNED_DATA.

	HTTPPassword string `json:"httpPassword,omitempty"` // HTTP password of the device / API key for Meraki Dashboard. Required if type is MERAKI_DASHBOARD or COMPUTE_DEVICE.

	HTTPPort string `json:"httpPort,omitempty"` // HTTP port of the device. Required if type is COMPUTE_DEVICE.

	HTTPSecure *bool `json:"httpSecure,omitempty"` // Flag to select HTTP / HTTPS protocol. Options are true / false. true for HTTPS and false for HTTP. Default is true.

	HTTPUserName string `json:"httpUserName,omitempty"` // HTTP Username of the device. Required if type is COMPUTE_DEVICE.

	IPAddress []string `json:"ipAddress,omitempty"` // IP Address of the device. Required if type is NETWORK_DEVICE, COMPUTE_DEVICE or THIRD_PARTY_DEVICE.

	MerakiOrgID []string `json:"merakiOrgId,omitempty"` // Selected Meraki organization for which the devices needs to be imported. Required if type is MERAKI_DASHBOARD.

	NetconfPort string `json:"netconfPort,omitempty"` // Netconf Port of the device. cliTransport must be 'ssh' if netconf is provided. Netconf port is required for eWLC.

	Password string `json:"password,omitempty"` // CLI Password of the device. Required if type is NETWORK_DEVICE.

	SerialNumber string `json:"serialNumber,omitempty"` // Serial Number of the Device. Required if extendedDiscoveryInfo is 'DISCOVER_WITH_CANNED_DATA'.

	SNMPAuthPassphrase string `json:"snmpAuthPassphrase,omitempty"` // SNMPv3 auth passphrase of the device. Required if snmpMode is authNoPriv or authPriv.

	SNMPAuthProtocol string `json:"snmpAuthProtocol,omitempty"` // SNMPv3 auth protocol. Supported values: sha, md5. Required if snmpMode is authNoPriv or authPriv.

	SNMPMode string `json:"snmpMode,omitempty"` // SNMPv3 mode. Supported values: noAuthnoPriv, authNoPriv, authPriv. Required if snmpVersion is v3.

	SNMPPrivPassphrase string `json:"snmpPrivPassphrase,omitempty"` // SNMPv3 priv passphrase. Required if snmpMode is authPriv.

	SNMPPrivProtocol string `json:"snmpPrivProtocol,omitempty"` // SNMPv3 priv protocol. Supported values: AES128. Required if snmpMode is authPriv.

	SNMPROCommunity string `json:"snmpROCommunity,omitempty"` // SNMP Read Community of the device. If snmpVersion is v2, at least one of snmpROCommunity and snmpRWCommunity is required.

	SNMPRWCommunity string `json:"snmpRWCommunity,omitempty"` // SNMP Write Community of the device. If snmpVersion is v2, at least one of snmpROCommunity and snmpRWCommunity is required.

	SNMPRetry *int `json:"snmpRetry,omitempty"` // SNMP retry count. Max value supported is 3. Default is Global SNMP retry (if exists) or 3.

	SNMPTimeout *int `json:"snmpTimeout,omitempty"` // SNMP timeout in seconds. Max value supported is 300. Default is Global SNMP timeout (if exists) or 5.

	SNMPUserName string `json:"snmpUserName,omitempty"` // SNMPV3 user name of the device. Required if snmpVersion is v3.

	SNMPVersion string `json:"snmpVersion,omitempty"` // SNMP version. Values supported: v2, v3. Required if type is NETWORK_DEVICE, COMPUTE_DEVICE or THIRD_PARTY_DEVICE.

	Type string `json:"type,omitempty"` // Type of device being added. Default is NETWORK_DEVICE.

	UserName string `json:"userName,omitempty"` // CLI user name of the device. Required if type is NETWORK_DEVICE.
}
type RequestDevicesUpdateDeviceDetailsV1 struct {
	CliTransport string `json:"cliTransport,omitempty"` // CLI transport. Supported values: telnet, ssh. Use NO!$DATA!$ if no change is required. Required if type is NETWORK_DEVICE.

	ComputeDevice *bool `json:"computeDevice,omitempty"` // Compute Device or not. Options are true / false.

	EnablePassword string `json:"enablePassword,omitempty"` // CLI enable password of the device. Required if device is configured to use enable password. Use NO!$DATA!$ if no change is required.

	ExtendedDiscoveryInfo string `json:"extendedDiscoveryInfo,omitempty"` // This field holds that info as whether to add device with canned data or not. Supported values: DISCOVER_WITH_CANNED_DATA.

	HTTPPassword string `json:"httpPassword,omitempty"` // HTTP password of the device / API key for Meraki Dashboard. Required if type is MERAKI_DASHBOARD or COMPUTE_DEVICE. Use NO!$DATA!$ if no change is required.

	HTTPPort string `json:"httpPort,omitempty"` // HTTP port of the device. Required if type is COMPUTE_DEVICE.

	HTTPSecure *bool `json:"httpSecure,omitempty"` // Flag to select HTTP / HTTPS protocol. Options are true / false. true for HTTPS and false for HTTP.

	HTTPUserName string `json:"httpUserName,omitempty"` // HTTP Username of the device. Required if type is COMPUTE_DEVICE. Use NO!$DATA!$ if no change is required.

	IPAddress []string `json:"ipAddress,omitempty"` // IP Address of the device. Required. Use 'api.meraki.com' for Meraki Dashboard.

	MerakiOrgID []string `json:"merakiOrgId,omitempty"` // Selected Meraki organization for which the devices needs to be imported. Required if type is MERAKI_DASHBOARD.

	NetconfPort string `json:"netconfPort,omitempty"` // Netconf Port of the device. cliTransport must be 'ssh' if netconf is provided. Netconf port is required for eWLC.

	Password string `json:"password,omitempty"` // CLI Password of the device. Required if type is NETWORK_DEVICE. Use NO!$DATA!$ if no change is required.

	SerialNumber string `json:"serialNumber,omitempty"` // Serial Number of the Device. Required if extendedDiscoveryInfo is 'DISCOVER_WITH_CANNED_DATA'.

	SNMPAuthPassphrase string `json:"snmpAuthPassphrase,omitempty"` // SNMPv3 auth passphrase of the device. Required if snmpMode is authNoPriv or authPriv. Use NO!$DATA!$ if no change is required.

	SNMPAuthProtocol string `json:"snmpAuthProtocol,omitempty"` // SNMPv3 auth protocol. Supported values: sha, md5.  Required if snmpMode is authNoPriv or authPriv. Use NODATACHANGE if no change is required.

	SNMPMode string `json:"snmpMode,omitempty"` // SNMPv3 mode. Supported values: noAuthnoPriv, authNoPriv, authPriv. Required if snmpVersion is v3. Use NODATACHANGE if no change is required.

	SNMPPrivPassphrase string `json:"snmpPrivPassphrase,omitempty"` // SNMPv3 priv passphrase. Required if snmpMode is authPriv. Use NO!$DATA!$ if no change is required.

	SNMPPrivProtocol string `json:"snmpPrivProtocol,omitempty"` // SNMPv3 priv protocol. Supported values: AES128. Required if snmpMode is authPriv. Use NODATACHANGE if no change is required.

	SNMPROCommunity string `json:"snmpROCommunity,omitempty"` // SNMP Read Community of the device. If snmpVersion is v2, at least one of snmpROCommunity and snmpRWCommunity is required. Use NO!$DATA!$ if no change is required.

	SNMPRWCommunity string `json:"snmpRWCommunity,omitempty"` // SNMP Write Community of the device. If snmpVersion is v2, at least one of snmpROCommunity and snmpRWCommunity is required. Use NO!$DATA!$ if no change is required.

	SNMPRetry *int `json:"snmpRetry,omitempty"` // SNMP retry count. Max value supported is 3. Default is Global SNMP retry (if exists) or 3.

	SNMPTimeout *int `json:"snmpTimeout,omitempty"` // SNMP timeout in seconds. Max value supported is 300. Default is Global SNMP timeout (if exists) or 5.

	SNMPUserName string `json:"snmpUserName,omitempty"` // SNMPV3 user name of the device. Required if snmpVersion is v3. Use NO!$DATA!$ if no change is required.

	SNMPVersion string `json:"snmpVersion,omitempty"` // SNMP version. Values supported: v2, v3. Required if type is NETWORK_DEVICE, COMPUTE_DEVICE or THIRD_PARTY_DEVICE. Use NODATACHANGE if no change is required.

	Type string `json:"type,omitempty"` // Type of device being edited. Default is NETWORK_DEVICE.

	UpdateMgmtIPaddressList *[]RequestDevicesUpdateDeviceDetailsV1UpdateMgmtIPaddressList `json:"updateMgmtIPaddressList,omitempty"` //

	UserName string `json:"userName,omitempty"` // CLI user name of the device. Required if type is NETWORK_DEVICE. Use NO!$DATA!$ if no change is required.
}
type RequestDevicesUpdateDeviceDetailsV1UpdateMgmtIPaddressList struct {
	ExistMgmtIPAddress string `json:"existMgmtIpAddress,omitempty"` // existMgmtIpAddress IP Address of the device.

	NewMgmtIPAddress string `json:"newMgmtIpAddress,omitempty"` // New IP Address to be Updated.
}
type RequestDevicesUpdateDeviceRoleV1 struct {
	ID string `json:"id,omitempty"` // DeviceId of the Device

	Role string `json:"role,omitempty"` // Role of device as ACCESS, CORE, DISTRIBUTION, BORDER ROUTER

	RoleSource string `json:"roleSource,omitempty"` // Role source as MANUAL / AUTO
}
type RequestDevicesExportDeviceListV1 struct {
	DeviceUUIDs []string `json:"deviceUuids,omitempty"` // List of device uuids

	OperationEnum string `json:"operationEnum,omitempty"` // 0 to export Device Credential Details Or 1 to export Device Details

	Parameters []string `json:"parameters,omitempty"` // List of device parameters that needs to be exported to file

	Password string `json:"password,omitempty"` // Password is required when the operationEnum value is 0
}
type RequestDevicesSyncDevicesV1 []string // Array of RequestDevicesSyncDevicesV1
type RequestDevicesCreateUserDefinedFieldV1 struct {
	Name string `json:"name,omitempty"` // Name of UDF

	Description string `json:"description,omitempty"` // Description of UDF
}
type RequestDevicesUpdateUserDefinedFieldV1 struct {
	Name string `json:"name,omitempty"` // Name of UDF

	Description string `json:"description,omitempty"` // Description of UDF
}
type RequestDevicesAddUserDefinedFieldToDeviceV1 []RequestItemDevicesAddUserDefinedFieldToDeviceV1 // Array of RequestDevicesAddUserDefinedFieldToDeviceV1
type RequestItemDevicesAddUserDefinedFieldToDeviceV1 struct {
	Name string `json:"name,omitempty"` // Name of the User Defined Field

	Value string `json:"value,omitempty"` // Value of the User Defined Field that will be assigned to the device
}
type RequestDevicesUpdateDeviceManagementAddressV1 struct {
	NewIP string `json:"newIP,omitempty"` // New IP Address of the device to be Updated
}
type RequestDevicesCreateMaintenanceScheduleForNetworkDevicesV1 struct {
	Description string `json:"description,omitempty"` // A brief narrative describing the maintenance schedule.

	MaintenanceSchedule *RequestDevicesCreateMaintenanceScheduleForNetworkDevicesV1MaintenanceSchedule `json:"maintenanceSchedule,omitempty"` //

	NetworkDeviceIDs []string `json:"networkDeviceIds,omitempty"` // List of network device ids. This field is applicable only during creation of schedules; for updates, it is read-only.
}
type RequestDevicesCreateMaintenanceScheduleForNetworkDevicesV1MaintenanceSchedule struct {
	StartTime *float64 `json:"startTime,omitempty"` // Start time indicates the beginning of the maintenance window in Unix epoch time in milliseconds.

	EndTime *float64 `json:"endTime,omitempty"` // End time indicates the ending of the maintenance window in Unix epoch time in milliseconds.

	Recurrence *RequestDevicesCreateMaintenanceScheduleForNetworkDevicesV1MaintenanceScheduleRecurrence `json:"recurrence,omitempty"` //
}
type RequestDevicesCreateMaintenanceScheduleForNetworkDevicesV1MaintenanceScheduleRecurrence struct {
	Interval *int `json:"interval,omitempty"` // Interval for recurrence in days. The interval must be longer than the duration of the schedules. The maximum allowed interval is 365 days.

	RecurrenceEndTime *float64 `json:"recurrenceEndTime,omitempty"` // The end date for the recurrence in Unix epoch time in milliseconds. Recurrence end time should be greater than maintenance end date/time.
}
type RequestDevicesUpdatesTheMaintenanceScheduleInformationV1 struct {
	Description string `json:"description,omitempty"` // A brief narrative describing the maintenance schedule.

	MaintenanceSchedule *RequestDevicesUpdatesTheMaintenanceScheduleInformationV1MaintenanceSchedule `json:"maintenanceSchedule,omitempty"` //

	NetworkDeviceIDs []string `json:"networkDeviceIds,omitempty"` // List of network device ids. This field is applicable only during creation of schedules; for updates, it is read-only.
}
type RequestDevicesUpdatesTheMaintenanceScheduleInformationV1MaintenanceSchedule struct {
	StartTime *float64 `json:"startTime,omitempty"` // Start time indicates the beginning of the maintenance window in Unix epoch time in milliseconds.

	EndTime *float64 `json:"endTime,omitempty"` // End time indicates the ending of the maintenance window in Unix epoch time in milliseconds.

	Recurrence *RequestDevicesUpdatesTheMaintenanceScheduleInformationV1MaintenanceScheduleRecurrence `json:"recurrence,omitempty"` //
}
type RequestDevicesUpdatesTheMaintenanceScheduleInformationV1MaintenanceScheduleRecurrence struct {
	Interval *int `json:"interval,omitempty"` // Interval for recurrence in days. The interval must be longer than the duration of the schedules. The maximum allowed interval is 365 days.

	RecurrenceEndTime *float64 `json:"recurrenceEndTime,omitempty"` // The end date for the recurrence in Unix epoch time in milliseconds. Recurrence end time should be greater than maintenance end date/time.
}
type RequestDevicesDeleteNetworkDeviceWithConfigurationCleanupV1 struct {
	ID string `json:"id,omitempty"` // The unique identifier of the network device to be deleted
}
type RequestDevicesDeleteANetworkDeviceWithoutConfigurationCleanupV1 struct {
	ID string `json:"id,omitempty"` // The unique identifier of the network device to be deleted
}
type RequestDevicesQueryNetworkDevicesWithFiltersV1 struct {
	Filter *RequestDevicesQueryNetworkDevicesWithFiltersV1Filter `json:"filter,omitempty"` //

	Views []string `json:"views,omitempty"` // The specific views being requested. This is an optional parameter which can be passed to get one or more of the network device data. If this is not provided, then it will default to BASIC views. If multiple views are provided, the response will contain the union of the views.

	Page *RequestDevicesQueryNetworkDevicesWithFiltersV1Page `json:"page,omitempty"` //
}
type RequestDevicesQueryNetworkDevicesWithFiltersV1Filter struct {
	LogicalOperator string `json:"logicalOperator,omitempty"` // The logical operator to use for combining the filter criteria. If not provided, the default value is AND.

	Filters *[]RequestDevicesQueryNetworkDevicesWithFiltersV1FilterFilters `json:"filters,omitempty"` //
}
type RequestDevicesQueryNetworkDevicesWithFiltersV1FilterFilters struct {
	Key string `json:"key,omitempty"` // The key to filter by

	Operator string `json:"operator,omitempty"` // The operator to use for filtering the values

	Value *RequestDevicesQueryNetworkDevicesWithFiltersV1FilterFiltersValue `json:"value,omitempty"` // Value to filter by. For `in` operator, the value should be a list of values.
}
type RequestDevicesQueryNetworkDevicesWithFiltersV1FilterFiltersValue interface{}
type RequestDevicesQueryNetworkDevicesWithFiltersV1Page struct {
	SortBy *RequestDevicesQueryNetworkDevicesWithFiltersV1PageSortBy `json:"sortBy,omitempty"` //

	Limit *int `json:"limit,omitempty"` // The number of records to show for this page. Min: 1, Max: 500

	Offset *int `json:"offset,omitempty"` // The first record to show for this page; the first record is numbered 1.
}
type RequestDevicesQueryNetworkDevicesWithFiltersV1PageSortBy struct {
	Name string `json:"name,omitempty"` // The field to sort by. Default is hostname.

	Order string `json:"order,omitempty"` // The order to sort by.
}
type RequestDevicesCountTheNumberOfNetworkDevicesWithFiltersV1 struct {
	Filter *RequestDevicesCountTheNumberOfNetworkDevicesWithFiltersV1Filter `json:"filter,omitempty"` //
}
type RequestDevicesCountTheNumberOfNetworkDevicesWithFiltersV1Filter struct {
	LogicalOperator string `json:"logicalOperator,omitempty"` // The logical operator to use for combining the filter criteria. If not provided, the default value is AND.

	Filters *[]RequestDevicesCountTheNumberOfNetworkDevicesWithFiltersV1FilterFilters `json:"filters,omitempty"` //
}
type RequestDevicesCountTheNumberOfNetworkDevicesWithFiltersV1FilterFilters struct {
	Key string `json:"key,omitempty"` // The key to filter by

	Operator string `json:"operator,omitempty"` // The operator to use for filtering the values

	Value *RequestDevicesCountTheNumberOfNetworkDevicesWithFiltersV1FilterFiltersValue `json:"value,omitempty"` // Value to filter by. For `in` operator, the value should be a list of values.
}
type RequestDevicesCountTheNumberOfNetworkDevicesWithFiltersV1FilterFiltersValue interface{}
type RequestDevicesUpdateGlobalResyncIntervalV1 struct {
	Interval *int `json:"interval,omitempty"` // Resync Interval should be between 25 to 1440 minutes
}
type RequestDevicesUpdateResyncIntervalForTheNetworkDeviceV1 struct {
	Interval *int `json:"interval,omitempty"` // Resync interval in minutes. To disable periodic resync, set interval as `0`. To use global settings, set interval as `null`.
}
type RequestDevicesRogueAdditionalDetailsV1 struct {
	Offset *float64 `json:"offset,omitempty"` // The offset of the first item in the collection to return. Default value is 1

	Limit *float64 `json:"limit,omitempty"` // The maximum number of entries to return. Default value is 1000

	StartTime *float64 `json:"startTime,omitempty"` // This is the epoch start time in milliseconds from which data need to be fetched. Default value is 24 hours earlier to endTime

	EndTime *float64 `json:"endTime,omitempty"` // This is the epoch end time in milliseconds upto which data need to be fetched. Default value is current time

	SiteID []string `json:"siteId,omitempty"` // Filter Rogues by location. Site IDs information can be fetched from "Get Site" API

	ThreatLevel []string `json:"threatLevel,omitempty"` // Filter Rogues by Threat Level. Threat Level information can be fetched from "Get Threat Levels" API

	ThreatType []string `json:"threatType,omitempty"` // Filter Rogues by Threat Type. Threat Type information can be fetched from "Get Threat Types" API
}
type RequestDevicesRogueAdditionalDetailCountV1 struct {
	StartTime *float64 `json:"startTime,omitempty"` // This is the epoch start time in milliseconds from which data need to be fetched. Default value is 24 hours earlier to endTime

	EndTime *float64 `json:"endTime,omitempty"` // This is the epoch end time in milliseconds upto which data need to be fetched. Default value is current time

	SiteID []string `json:"siteId,omitempty"` // Filter Rogues by location. Site IDs information can be fetched from "Get Site" API

	ThreatLevel []string `json:"threatLevel,omitempty"` // This information can be fetched from "Get Threat Levels" API

	ThreatType []string `json:"threatType,omitempty"` // This information can be fetched from "Get Threat Types" API
}
type RequestDevicesStartWirelessRogueApContainmentV1 struct {
	MacAddress string `json:"macAddress,omitempty"` // Mac Address

	Type *int `json:"type,omitempty"` // Type
}
type RequestDevicesStopWirelessRogueApContainmentV1 struct {
	MacAddress string `json:"macAddress,omitempty"` // Mac Address

	Type *int `json:"type,omitempty"` // Type

	WlcIP string `json:"wlcIp,omitempty"` // Wlc Ip
}
type RequestDevicesThreatDetailsV1 struct {
	Offset *int `json:"offset,omitempty"` // Offset

	Limit *int `json:"limit,omitempty"` // Limit

	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	SiteID []string `json:"siteId,omitempty"` // Site Id

	ThreatLevel []string `json:"threatLevel,omitempty"` // Threat Level

	ThreatType []string `json:"threatType,omitempty"` // Threat Type

	IsNewThreat *bool `json:"isNewThreat,omitempty"` // Is New Threat
}
type RequestDevicesThreatDetailCountV1 struct {
	Offset *int `json:"offset,omitempty"` // Offset

	Limit *int `json:"limit,omitempty"` // Limit

	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	SiteID []string `json:"siteId,omitempty"` // Site Id

	ThreatLevel []string `json:"threatLevel,omitempty"` // Threat Level

	ThreatType []string `json:"threatType,omitempty"` // Threat Type

	IsNewThreat *bool `json:"isNewThreat,omitempty"` // Is New Threat
}
type RequestDevicesAddAllowedMacAddressV1 []RequestItemDevicesAddAllowedMacAddressV1 // Array of RequestDevicesAddAllowedMacAddressV1
type RequestItemDevicesAddAllowedMacAddressV1 struct {
	MacAddress string `json:"macAddress,omitempty"` // Mac Address

	Category *int `json:"category,omitempty"` // Category
}
type RequestDevicesThreatSummaryV1 struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	SiteID []string `json:"siteId,omitempty"` // Site Id

	ThreatLevel []string `json:"threatLevel,omitempty"` // Threat Level

	ThreatType []string `json:"threatType,omitempty"` // Threat Type
}
type RequestDevicesGetDeviceInterfaceStatsInfoV2 struct {
	StartTime *int `json:"startTime,omitempty"` // UTC epoch timestamp in milliseconds

	EndTime *int `json:"endTime,omitempty"` // UTC epoch timestamp in milliseconds

	Query *RequestDevicesGetDeviceInterfaceStatsInfoV2Query `json:"query,omitempty"` //
}
type RequestDevicesGetDeviceInterfaceStatsInfoV2Query struct {
	Fields *[]RequestDevicesGetDeviceInterfaceStatsInfoV2QueryFields `json:"fields,omitempty"` // Required field names, default ALL

	Filters *[]RequestDevicesGetDeviceInterfaceStatsInfoV2QueryFilters `json:"filters,omitempty"` //

	Page *RequestDevicesGetDeviceInterfaceStatsInfoV2QueryPage `json:"page,omitempty"` //
}
type RequestDevicesGetDeviceInterfaceStatsInfoV2QueryFields interface{}
type RequestDevicesGetDeviceInterfaceStatsInfoV2QueryFilters struct {
	Key string `json:"key,omitempty"` // Name of the field that the filter should be applied to

	Operator string `json:"operator,omitempty"` // Supported operators are eq,in,like

	Value string `json:"value,omitempty"` // Value of the field
}
type RequestDevicesGetDeviceInterfaceStatsInfoV2QueryPage struct {
	Limit *int `json:"limit,omitempty"` // Number of records, Max is 1000

	Offset *float64 `json:"offset,omitempty"` // Record offset value, default 0

	OrderBy *[]RequestDevicesGetDeviceInterfaceStatsInfoV2QueryPageOrderBy `json:"orderBy,omitempty"` //
}
type RequestDevicesGetDeviceInterfaceStatsInfoV2QueryPageOrderBy struct {
	Name string `json:"name,omitempty"` // Name of the field used to sort

	Order string `json:"order,omitempty"` // Possible values asc, des
}

//RetrievesTheListOfAAAServicesForGivenParametersV1 Retrieves the list of AAA Services for given parameters. - b990-0822-4c08-a780
/* Retrieves the list of AAA Services and offers basic filtering and sorting capabilities. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AAAServices-1.0.0-resolved.yaml


@param RetrievesTheListOfAAAServicesForGivenParametersV1HeaderParams Custom header parameters
@param RetrievesTheListOfAAAServicesForGivenParametersV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-list-of-a-a-a-services-for-given-parameters
*/
func (s *DevicesService) RetrievesTheListOfAAAServicesForGivenParametersV1(RetrievesTheListOfAAAServicesForGivenParametersV1HeaderParams *RetrievesTheListOfAAAServicesForGivenParametersV1HeaderParams, RetrievesTheListOfAAAServicesForGivenParametersV1QueryParams *RetrievesTheListOfAAAServicesForGivenParametersV1QueryParams) (*ResponseDevicesRetrievesTheListOfAAAServicesForGivenParametersV1, *resty.Response, error) {
	path := "/dna/data/api/v1/aaaServices"

	queryString, _ := query.Values(RetrievesTheListOfAAAServicesForGivenParametersV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheListOfAAAServicesForGivenParametersV1HeaderParams != nil {

		if RetrievesTheListOfAAAServicesForGivenParametersV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheListOfAAAServicesForGivenParametersV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesRetrievesTheListOfAAAServicesForGivenParametersV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheListOfAAAServicesForGivenParametersV1(RetrievesTheListOfAAAServicesForGivenParametersV1HeaderParams, RetrievesTheListOfAAAServicesForGivenParametersV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheListOfAAAServicesForGivenParametersV1")
	}

	result := response.Result().(*ResponseDevicesRetrievesTheListOfAAAServicesForGivenParametersV1)
	return result, response, err

}

//RetrievesTheTotalNumberOfAAAServicesForGivenParametersV1 Retrieves the total number of AAA Services for given parameters. - c393-f961-4939-b53c
/* Retrieves the total number of AAA Services for given parameters. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AAAServices-1.0.0-resolved.yaml


@param RetrievesTheTotalNumberOfAAAServicesForGivenParametersV1HeaderParams Custom header parameters
@param RetrievesTheTotalNumberOfAAAServicesForGivenParametersV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-total-number-of-a-a-a-services-for-given-parameters
*/
func (s *DevicesService) RetrievesTheTotalNumberOfAAAServicesForGivenParametersV1(RetrievesTheTotalNumberOfAAAServicesForGivenParametersV1HeaderParams *RetrievesTheTotalNumberOfAAAServicesForGivenParametersV1HeaderParams, RetrievesTheTotalNumberOfAAAServicesForGivenParametersV1QueryParams *RetrievesTheTotalNumberOfAAAServicesForGivenParametersV1QueryParams) (*ResponseDevicesRetrievesTheTotalNumberOfAAAServicesForGivenParametersV1, *resty.Response, error) {
	path := "/dna/data/api/v1/aaaServices/count"

	queryString, _ := query.Values(RetrievesTheTotalNumberOfAAAServicesForGivenParametersV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheTotalNumberOfAAAServicesForGivenParametersV1HeaderParams != nil {

		if RetrievesTheTotalNumberOfAAAServicesForGivenParametersV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheTotalNumberOfAAAServicesForGivenParametersV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesRetrievesTheTotalNumberOfAAAServicesForGivenParametersV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheTotalNumberOfAAAServicesForGivenParametersV1(RetrievesTheTotalNumberOfAAAServicesForGivenParametersV1HeaderParams, RetrievesTheTotalNumberOfAAAServicesForGivenParametersV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheTotalNumberOfAAAServicesForGivenParametersV1")
	}

	result := response.Result().(*ResponseDevicesRetrievesTheTotalNumberOfAAAServicesForGivenParametersV1)
	return result, response, err

}

//RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIDOfTheServiceV1 Retrieves the details of a specific AAA Service matching the id of the Service. - 35a1-3ae4-4cbb-ae6f
/* Retrieves the details of the AAA Service matching the given id. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AAAServices-1.0.0-resolved.yaml


@param id id path parameter. Unique id of the AAA Service. It is the combination of AAA Server IP (`serverIp`) and Device UUID (`deviceId`) separated by underscore (`_`). Example: If `serverIp` is `10.76.81.33` and `deviceId` is `6bef213c-19ca-4170-8375-b694e251101c`, then the `id` would be `10.76.81.33_6bef213c-19ca-4170-8375-b694e251101c`

@param RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIdOfTheServiceV1HeaderParams Custom header parameters
@param RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIdOfTheServiceV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-details-of-a-specific-a-a-a-service-matching-the-id-of-the-service
*/
func (s *DevicesService) RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIDOfTheServiceV1(id string, RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIdOfTheServiceV1HeaderParams *RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIDOfTheServiceV1HeaderParams, RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIdOfTheServiceV1QueryParams *RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIDOfTheServiceV1QueryParams) (*ResponseDevicesRetrievesTheDetailsOfASpecificAAAServiceMatchingTheIDOfTheServiceV1, *resty.Response, error) {
	path := "/dna/data/api/v1/aaaServices/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIdOfTheServiceV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIdOfTheServiceV1HeaderParams != nil {

		if RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIdOfTheServiceV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIdOfTheServiceV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesRetrievesTheDetailsOfASpecificAAAServiceMatchingTheIDOfTheServiceV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIDOfTheServiceV1(id, RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIdOfTheServiceV1HeaderParams, RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIdOfTheServiceV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIdOfTheServiceV1")
	}

	result := response.Result().(*ResponseDevicesRetrievesTheDetailsOfASpecificAAAServiceMatchingTheIDOfTheServiceV1)
	return result, response, err

}

//QueryAssuranceEventsV1 Query assurance events - 15a5-3b2c-4908-8ba3
/* Returns the list of events discovered by Catalyst Center, determined by the complex filters. Please refer to the 'API Support Documentation' section to understand which fields are supported. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceEvents-1.0.0-resolved.yaml


@param QueryAssuranceEventsV1HeaderParams Custom header parameters
@param QueryAssuranceEventsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!query-assurance-events
*/
func (s *DevicesService) QueryAssuranceEventsV1(QueryAssuranceEventsV1HeaderParams *QueryAssuranceEventsV1HeaderParams, QueryAssuranceEventsV1QueryParams *QueryAssuranceEventsV1QueryParams) (*ResponseDevicesQueryAssuranceEventsV1, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceEvents"

	queryString, _ := query.Values(QueryAssuranceEventsV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if QueryAssuranceEventsV1HeaderParams != nil {

		if QueryAssuranceEventsV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", QueryAssuranceEventsV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesQueryAssuranceEventsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.QueryAssuranceEventsV1(QueryAssuranceEventsV1HeaderParams, QueryAssuranceEventsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation QueryAssuranceEventsV1")
	}

	result := response.Result().(*ResponseDevicesQueryAssuranceEventsV1)
	return result, response, err

}

//CountTheNumberOfEventsV1 Count the number of events - 349f-a9d8-4a6a-b951
/* API to fetch the count of assurance events that match the filter criteria. Please refer to the 'API Support Documentation' section to understand which fields are supported. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceEvents-1.0.0-resolved.yaml


@param CountTheNumberOfEventsV1HeaderParams Custom header parameters
@param CountTheNumberOfEventsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!count-the-number-of-events
*/
func (s *DevicesService) CountTheNumberOfEventsV1(CountTheNumberOfEventsV1HeaderParams *CountTheNumberOfEventsV1HeaderParams, CountTheNumberOfEventsV1QueryParams *CountTheNumberOfEventsV1QueryParams) (*ResponseDevicesCountTheNumberOfEventsV1, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceEvents/count"

	queryString, _ := query.Values(CountTheNumberOfEventsV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if CountTheNumberOfEventsV1HeaderParams != nil {

		if CountTheNumberOfEventsV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", CountTheNumberOfEventsV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesCountTheNumberOfEventsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.CountTheNumberOfEventsV1(CountTheNumberOfEventsV1HeaderParams, CountTheNumberOfEventsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation CountTheNumberOfEventsV1")
	}

	result := response.Result().(*ResponseDevicesCountTheNumberOfEventsV1)
	return result, response, err

}

//GetDetailsOfASingleAssuranceEventV1 Get details of a single assurance event - 039e-2909-449a-8f51
/* API to fetch the details of an assurance event using event `id`. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceEvents-1.0.0-resolved.yaml


@param id id path parameter. Unique identifier for the event

@param GetDetailsOfASingleAssuranceEventV1HeaderParams Custom header parameters
@param GetDetailsOfASingleAssuranceEventV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-details-of-a-single-assurance-event
*/
func (s *DevicesService) GetDetailsOfASingleAssuranceEventV1(id string, GetDetailsOfASingleAssuranceEventV1HeaderParams *GetDetailsOfASingleAssuranceEventV1HeaderParams, GetDetailsOfASingleAssuranceEventV1QueryParams *GetDetailsOfASingleAssuranceEventV1QueryParams) (*ResponseDevicesGetDetailsOfASingleAssuranceEventV1, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceEvents/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetDetailsOfASingleAssuranceEventV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetDetailsOfASingleAssuranceEventV1HeaderParams != nil {

		if GetDetailsOfASingleAssuranceEventV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetDetailsOfASingleAssuranceEventV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetDetailsOfASingleAssuranceEventV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDetailsOfASingleAssuranceEventV1(id, GetDetailsOfASingleAssuranceEventV1HeaderParams, GetDetailsOfASingleAssuranceEventV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDetailsOfASingleAssuranceEventV1")
	}

	result := response.Result().(*ResponseDevicesGetDetailsOfASingleAssuranceEventV1)
	return result, response, err

}

//GetListOfChildEventsForTheGivenWirelessClientEventV1 Get list of child events for the given wireless client event - d78f-7acc-4a88-b616
/* Wireless client event could have child events and this API can be used to fetch the same using parent event `id` as the input. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceEvents-1.0.0-resolved.yaml


@param id id path parameter. Unique identifier for the event

@param GetListOfChildEventsForTheGivenWirelessClientEventV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-list-of-child-events-for-the-given-wireless-client-event
*/
func (s *DevicesService) GetListOfChildEventsForTheGivenWirelessClientEventV1(id string, GetListOfChildEventsForTheGivenWirelessClientEventV1HeaderParams *GetListOfChildEventsForTheGivenWirelessClientEventV1HeaderParams) (*ResponseDevicesGetListOfChildEventsForTheGivenWirelessClientEventV1, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceEvents/{id}/childEvents"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetListOfChildEventsForTheGivenWirelessClientEventV1HeaderParams != nil {

		if GetListOfChildEventsForTheGivenWirelessClientEventV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetListOfChildEventsForTheGivenWirelessClientEventV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetResult(&ResponseDevicesGetListOfChildEventsForTheGivenWirelessClientEventV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetListOfChildEventsForTheGivenWirelessClientEventV1(id, GetListOfChildEventsForTheGivenWirelessClientEventV1HeaderParams)
		}
		return nil, response, fmt.Errorf("error with operation GetListOfChildEventsForTheGivenWirelessClientEventV1")
	}

	result := response.Result().(*ResponseDevicesGetListOfChildEventsForTheGivenWirelessClientEventV1)
	return result, response, err

}

//RetrievesTheListOfDHCPServicesForGivenParametersV1 Retrieves the list of DHCP Services for given parameters. - bfa0-ebff-418a-b093
/* Retrieves the list of DHCP Services and offers basic filtering and sorting capabilities. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DHCPServices-1.0.0-resolved.yaml


@param RetrievesTheListOfDHCPServicesForGivenParametersV1HeaderParams Custom header parameters
@param RetrievesTheListOfDHCPServicesForGivenParametersV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-list-of-d-h-c-p-services-for-given-parameters
*/
func (s *DevicesService) RetrievesTheListOfDHCPServicesForGivenParametersV1(RetrievesTheListOfDHCPServicesForGivenParametersV1HeaderParams *RetrievesTheListOfDHCPServicesForGivenParametersV1HeaderParams, RetrievesTheListOfDHCPServicesForGivenParametersV1QueryParams *RetrievesTheListOfDHCPServicesForGivenParametersV1QueryParams) (*ResponseDevicesRetrievesTheListOfDHCPServicesForGivenParametersV1, *resty.Response, error) {
	path := "/dna/data/api/v1/dhcpServices"

	queryString, _ := query.Values(RetrievesTheListOfDHCPServicesForGivenParametersV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheListOfDHCPServicesForGivenParametersV1HeaderParams != nil {

		if RetrievesTheListOfDHCPServicesForGivenParametersV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheListOfDHCPServicesForGivenParametersV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesRetrievesTheListOfDHCPServicesForGivenParametersV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheListOfDHCPServicesForGivenParametersV1(RetrievesTheListOfDHCPServicesForGivenParametersV1HeaderParams, RetrievesTheListOfDHCPServicesForGivenParametersV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheListOfDHCPServicesForGivenParametersV1")
	}

	result := response.Result().(*ResponseDevicesRetrievesTheListOfDHCPServicesForGivenParametersV1)
	return result, response, err

}

//RetrievesTheTotalNumberOfDHCPServicesForGivenParametersV1 Retrieves the total number of DHCP Services for given parameters. - 8eaf-6891-4319-9f95
/* Retrieves the total number of DHCP Services for given parameters. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DHCPServices-1.0.0-resolved.yaml


@param RetrievesTheTotalNumberOfDHCPServicesForGivenParametersV1HeaderParams Custom header parameters
@param RetrievesTheTotalNumberOfDHCPServicesForGivenParametersV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-total-number-of-d-h-c-p-services-for-given-parameters
*/
func (s *DevicesService) RetrievesTheTotalNumberOfDHCPServicesForGivenParametersV1(RetrievesTheTotalNumberOfDHCPServicesForGivenParametersV1HeaderParams *RetrievesTheTotalNumberOfDHCPServicesForGivenParametersV1HeaderParams, RetrievesTheTotalNumberOfDHCPServicesForGivenParametersV1QueryParams *RetrievesTheTotalNumberOfDHCPServicesForGivenParametersV1QueryParams) (*ResponseDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenParametersV1, *resty.Response, error) {
	path := "/dna/data/api/v1/dhcpServices/count"

	queryString, _ := query.Values(RetrievesTheTotalNumberOfDHCPServicesForGivenParametersV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheTotalNumberOfDHCPServicesForGivenParametersV1HeaderParams != nil {

		if RetrievesTheTotalNumberOfDHCPServicesForGivenParametersV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheTotalNumberOfDHCPServicesForGivenParametersV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenParametersV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheTotalNumberOfDHCPServicesForGivenParametersV1(RetrievesTheTotalNumberOfDHCPServicesForGivenParametersV1HeaderParams, RetrievesTheTotalNumberOfDHCPServicesForGivenParametersV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheTotalNumberOfDHCPServicesForGivenParametersV1")
	}

	result := response.Result().(*ResponseDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenParametersV1)
	return result, response, err

}

//RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheServiceV1 Retrieves the details of a specific DHCP Service matching the id of the Service. - 3287-8874-4319-8db1
/* Retrieves the details of the DHCP Service matching the given id. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DHCPServices-1.0.0-resolved.yaml


@param id id path parameter. Unique id of the DHCP Service. It is the combination of DHCP Server IP (`serverIp`) and Device UUID (`deviceId`) separated by underscore (`_`). Example: If `serverIp` is `10.76.81.33` and `deviceId` is `6bef213c-19ca-4170-8375-b694e251101c`, then the `id` would be `10.76.81.33_6bef213c-19ca-4170-8375-b694e251101c`

@param RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIdOfTheServiceV1HeaderParams Custom header parameters
@param RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIdOfTheServiceV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-details-of-a-specific-d-h-c-p-service-matching-the-id-of-the-service
*/
func (s *DevicesService) RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheServiceV1(id string, RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIdOfTheServiceV1HeaderParams *RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheServiceV1HeaderParams, RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIdOfTheServiceV1QueryParams *RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheServiceV1QueryParams) (*ResponseDevicesRetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheServiceV1, *resty.Response, error) {
	path := "/dna/data/api/v1/dhcpServices/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIdOfTheServiceV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIdOfTheServiceV1HeaderParams != nil {

		if RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIdOfTheServiceV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIdOfTheServiceV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesRetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheServiceV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheServiceV1(id, RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIdOfTheServiceV1HeaderParams, RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIdOfTheServiceV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIdOfTheServiceV1")
	}

	result := response.Result().(*ResponseDevicesRetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheServiceV1)
	return result, response, err

}

//RetrievesTheListOfDNSServicesForGivenParametersV1 Retrieves the list of DNS Services for given parameters. - 0bbd-2bd5-4f9b-9a57
/* Retrieves the list of DNS Services and offers basic filtering and sorting capabilities. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DNSServices-1.0.0-resolved.yaml


@param RetrievesTheListOfDNSServicesForGivenParametersV1HeaderParams Custom header parameters
@param RetrievesTheListOfDNSServicesForGivenParametersV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-list-of-d-n-s-services-for-given-parameters
*/
func (s *DevicesService) RetrievesTheListOfDNSServicesForGivenParametersV1(RetrievesTheListOfDNSServicesForGivenParametersV1HeaderParams *RetrievesTheListOfDNSServicesForGivenParametersV1HeaderParams, RetrievesTheListOfDNSServicesForGivenParametersV1QueryParams *RetrievesTheListOfDNSServicesForGivenParametersV1QueryParams) (*ResponseDevicesRetrievesTheListOfDNSServicesForGivenParametersV1, *resty.Response, error) {
	path := "/dna/data/api/v1/dnsServices"

	queryString, _ := query.Values(RetrievesTheListOfDNSServicesForGivenParametersV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheListOfDNSServicesForGivenParametersV1HeaderParams != nil {

		if RetrievesTheListOfDNSServicesForGivenParametersV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheListOfDNSServicesForGivenParametersV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesRetrievesTheListOfDNSServicesForGivenParametersV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheListOfDNSServicesForGivenParametersV1(RetrievesTheListOfDNSServicesForGivenParametersV1HeaderParams, RetrievesTheListOfDNSServicesForGivenParametersV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheListOfDNSServicesForGivenParametersV1")
	}

	result := response.Result().(*ResponseDevicesRetrievesTheListOfDNSServicesForGivenParametersV1)
	return result, response, err

}

//RetrievesTheTotalNumberOfDNSServicesForGivenParametersV1 Retrieves the total number of DNS Services for given parameters. - 4385-991e-43a9-9561
/* Retrieves the total number of DNS Services for given parameters. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DNSServices-1.0.0-resolved.yaml


@param RetrievesTheTotalNumberOfDNSServicesForGivenParametersV1HeaderParams Custom header parameters
@param RetrievesTheTotalNumberOfDNSServicesForGivenParametersV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-total-number-of-d-n-s-services-for-given-parameters
*/
func (s *DevicesService) RetrievesTheTotalNumberOfDNSServicesForGivenParametersV1(RetrievesTheTotalNumberOfDNSServicesForGivenParametersV1HeaderParams *RetrievesTheTotalNumberOfDNSServicesForGivenParametersV1HeaderParams, RetrievesTheTotalNumberOfDNSServicesForGivenParametersV1QueryParams *RetrievesTheTotalNumberOfDNSServicesForGivenParametersV1QueryParams) (*ResponseDevicesRetrievesTheTotalNumberOfDNSServicesForGivenParametersV1, *resty.Response, error) {
	path := "/dna/data/api/v1/dnsServices/count"

	queryString, _ := query.Values(RetrievesTheTotalNumberOfDNSServicesForGivenParametersV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheTotalNumberOfDNSServicesForGivenParametersV1HeaderParams != nil {

		if RetrievesTheTotalNumberOfDNSServicesForGivenParametersV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheTotalNumberOfDNSServicesForGivenParametersV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesRetrievesTheTotalNumberOfDNSServicesForGivenParametersV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheTotalNumberOfDNSServicesForGivenParametersV1(RetrievesTheTotalNumberOfDNSServicesForGivenParametersV1HeaderParams, RetrievesTheTotalNumberOfDNSServicesForGivenParametersV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheTotalNumberOfDNSServicesForGivenParametersV1")
	}

	result := response.Result().(*ResponseDevicesRetrievesTheTotalNumberOfDNSServicesForGivenParametersV1)
	return result, response, err

}

//RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIDOfTheServiceV1 Retrieves the details of a specific DNS Service matching the id of the Service. - 84ab-b9c3-498a-b6a7
/* Retrieves the details of the DNS Service matching the given id. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DNSServices-1.0.0-resolved.yaml


@param id id path parameter. Unique id of the DNS Service. It is the combination of DNS Server IP (`serverIp`) and Device UUID (`deviceId`) separated by underscore (`_`). Example: If `serverIp` is `10.76.81.33` and `deviceId` is `6bef213c-19ca-4170-8375-b694e251101c`, then the `id` would be `10.76.81.33_6bef213c-19ca-4170-8375-b694e251101c`

@param RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIdOfTheServiceV1HeaderParams Custom header parameters
@param RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIdOfTheServiceV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-details-of-a-specific-d-n-s-service-matching-the-id-of-the-service
*/
func (s *DevicesService) RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIDOfTheServiceV1(id string, RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIdOfTheServiceV1HeaderParams *RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIDOfTheServiceV1HeaderParams, RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIdOfTheServiceV1QueryParams *RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIDOfTheServiceV1QueryParams) (*ResponseDevicesRetrievesTheDetailsOfASpecificDNSServiceMatchingTheIDOfTheServiceV1, *resty.Response, error) {
	path := "/dna/data/api/v1/dnsServices/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIdOfTheServiceV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIdOfTheServiceV1HeaderParams != nil {

		if RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIdOfTheServiceV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIdOfTheServiceV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesRetrievesTheDetailsOfASpecificDNSServiceMatchingTheIDOfTheServiceV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIDOfTheServiceV1(id, RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIdOfTheServiceV1HeaderParams, RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIdOfTheServiceV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIdOfTheServiceV1")
	}

	result := response.Result().(*ResponseDevicesRetrievesTheDetailsOfASpecificDNSServiceMatchingTheIDOfTheServiceV1)
	return result, response, err

}

//GetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesV1 Gets interfaces along with statistics and poe data from all network devices. - 9898-9b5a-445b-884f
/* Retrieves the list of the interfaces from all network devices based on the provided query parameters. The latest interfaces data in the specified start and end time range will be returned. When there is no start and end time specified returns the latest available data.

The elements are grouped and sorted by deviceUuid first, and are then sorted by the given sort field, or by the default value: name.


The supported sorting options are:
name, adminStatus, description, duplexConfig, duplexOper, interfaceIfIndex,interfaceType, macAddress,mediaType, operStatus, portChannelId, portMode, portType,speed, vlanId



This API can paginate up to 500,000 records, please narrow matching results with additional filters beyond that value. The elements are grouped and sorted by deviceUuid first, and are then sorted by the given sort field, or by the default value: name.

 The supported sorting options are: name, adminStatus, description, duplexConfig, duplexOper,interfaceIfIndex,interfaceType, macAddress,mediaType, operStatus,portChannelId, portMode, portType,speed, vlanId,pdPowerAdminMaxInWatt,pdPowerBudgetInWatt,pdPowerConsumedInWatt,pdPowerRemainingInWatt,pdMaxPowerDrawn. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-interfaces-2.0.0-resolved.yaml


@param GetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-interfaces-along-with-statistics-and-poe-data-from-all-network-devices
*/
func (s *DevicesService) GetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesV1(GetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesV1QueryParams *GetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesV1QueryParams) (*ResponseDevicesGetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesV1, *resty.Response, error) {
	path := "/dna/data/api/v1/interfaces"

	queryString, _ := query.Values(GetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesV1(GetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesV1")
	}

	result := response.Result().(*ResponseDevicesGetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesV1)
	return result, response, err

}

//GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountV1 Gets the total Network device interface counts in the specified time range. When there is no start and end time specified returns the latest interfaces total count. - 40ab-799f-465a-82f4
/* Gets the total Network device interface counts. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-interfaces-2.0.0-resolved.yaml


@param GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-the-total-network-device-interface-counts-in-the-specified-time-range-when-there-is-no-start-and-end-time-specified-returns-the-latest-interfaces-total-count
*/
func (s *DevicesService) GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountV1(GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountV1QueryParams *GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountV1QueryParams) (*ResponseDevicesGetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountV1, *resty.Response, error) {
	path := "/dna/data/api/v1/interfaces/count"

	queryString, _ := query.Values(GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountV1(GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountV1")
	}

	result := response.Result().(*ResponseDevicesGetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountV1)
	return result, response, err

}

//GetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsAndPoeDataV1 Get the interface data for the given interface id (instance Uuid) along with the statistics and poe data - c08d-d95c-4c7b-8283
/* Returns the interface data for the given interface instance Uuid along with the statistics data. The latest interface data in the specified start and end time range will be returned. When there is no start and end time specified returns the latest available data for the given interface Id. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-interfaces-2.0.0-resolved.yaml


@param id id path parameter. The interface Uuid

@param GetTheInterfaceDataForTheGivenInterfaceIdinstanceUuidAlongWithTheStatisticsAndPoeDataV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-the-interface-data-for-the-given-interface-idinstance-uuid-along-with-the-statistics-and-poe-data
*/
func (s *DevicesService) GetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsAndPoeDataV1(id string, GetTheInterfaceDataForTheGivenInterfaceIdinstanceUuidAlongWithTheStatisticsAndPoeDataV1QueryParams *GetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsAndPoeDataV1QueryParams) (*ResponseDevicesGetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsAndPoeDataV1, *resty.Response, error) {
	path := "/dna/data/api/v1/interfaces/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetTheInterfaceDataForTheGivenInterfaceIdinstanceUuidAlongWithTheStatisticsAndPoeDataV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsAndPoeDataV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsAndPoeDataV1(id, GetTheInterfaceDataForTheGivenInterfaceIdinstanceUuidAlongWithTheStatisticsAndPoeDataV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTheInterfaceDataForTheGivenInterfaceIdinstanceUuidAlongWithTheStatisticsAndPoeDataV1")
	}

	result := response.Result().(*ResponseDevicesGetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsAndPoeDataV1)
	return result, response, err

}

//GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1 Gets the Network Device details based on the provided query parameters. - c8b4-f894-4c3a-932f
/* Gets the Network Device details based on the provided query parameters.  When there is no start and end time specified returns the latest device details. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceNetworkDevices-2.0.1-resolved.yaml


@param GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-the-network-device-details-based-on-the-provided-query-parameters
*/
func (s *DevicesService) GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1(GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1QueryParams *GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1QueryParams) (*ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1, *resty.Response, error) {
	path := "/dna/data/api/v1/networkDevices"

	queryString, _ := query.Values(GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1(GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1")
	}

	result := response.Result().(*ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1)
	return result, response, err

}

//GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersV1 Gets the total Network device counts based on the provided query parameters. - f0a6-e96b-44fb-a549
/* Gets the total Network device counts. When there is no start and end time specified returns the latest interfaces total count. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceNetworkDevices-2.0.1-resolved.yaml


@param GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-the-total-network-device-counts-based-on-the-provided-query-parameters
*/
func (s *DevicesService) GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersV1(GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersV1QueryParams *GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersV1QueryParams) (*ResponseDevicesGetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersV1, *resty.Response, error) {
	path := "/dna/data/api/v1/networkDevices/count"

	queryString, _ := query.Values(GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersV1(GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersV1")
	}

	result := response.Result().(*ResponseDevicesGetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersV1)
	return result, response, err

}

//GetTheDeviceDataForTheGivenDeviceIDUUIDV1 Get the device data for the given device id (Uuid) - 5a93-1957-475b-95b3
/* Returns the device data for the given device Uuid in the specified start and end time range. When there is no start and end time specified returns the latest available data for the given Id. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceNetworkDevices-2.0.1-resolved.yaml


@param id id path parameter. The device Uuid

@param GetTheDeviceDataForTheGivenDeviceIdUuidV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-the-device-data-for-the-given-device-id-uuid
*/
func (s *DevicesService) GetTheDeviceDataForTheGivenDeviceIDUUIDV1(id string, GetTheDeviceDataForTheGivenDeviceIdUuidV1QueryParams *GetTheDeviceDataForTheGivenDeviceIDUUIDV1QueryParams) (*ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDV1, *resty.Response, error) {
	path := "/dna/data/api/v1/networkDevices/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetTheDeviceDataForTheGivenDeviceIdUuidV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTheDeviceDataForTheGivenDeviceIDUUIDV1(id, GetTheDeviceDataForTheGivenDeviceIdUuidV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTheDeviceDataForTheGivenDeviceIdUuidV1")
	}

	result := response.Result().(*ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDV1)
	return result, response, err

}

//GetPlannedAccessPointsForBuildingV1 Get Planned Access Points for Building - b699-9b85-4e3b-acdd
/* Provides a list of Planned Access Points for the Building it is requested for


@param buildingID buildingId path parameter. The instance UUID of the building hierarchy element

@param GetPlannedAccessPointsForBuildingV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-planned-access-points-for-building
*/
func (s *DevicesService) GetPlannedAccessPointsForBuildingV1(buildingID string, GetPlannedAccessPointsForBuildingV1QueryParams *GetPlannedAccessPointsForBuildingV1QueryParams) (*ResponseDevicesGetPlannedAccessPointsForBuildingV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/buildings/{buildingId}/planned-access-points"
	path = strings.Replace(path, "{buildingId}", fmt.Sprintf("%v", buildingID), -1)

	queryString, _ := query.Values(GetPlannedAccessPointsForBuildingV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetPlannedAccessPointsForBuildingV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPlannedAccessPointsForBuildingV1(buildingID, GetPlannedAccessPointsForBuildingV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetPlannedAccessPointsForBuildingV1")
	}

	result := response.Result().(*ResponseDevicesGetPlannedAccessPointsForBuildingV1)
	return result, response, err

}

//GetDeviceDetailV1 Get Device Detail - ca98-fac4-4b08-895c
/* Returns detailed Network Device information retrieved by Mac Address, Device Name or UUID for any given point of time.


@param GetDeviceDetailV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-detail
*/
func (s *DevicesService) GetDeviceDetailV1(GetDeviceDetailV1QueryParams *GetDeviceDetailV1QueryParams) (*ResponseDevicesGetDeviceDetailV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/device-detail"

	queryString, _ := query.Values(GetDeviceDetailV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetDeviceDetailV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceDetailV1(GetDeviceDetailV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceDetailV1")
	}

	result := response.Result().(*ResponseDevicesGetDeviceDetailV1)
	return result, response, err

}

//GetDeviceEnrichmentDetailsV1 Get Device Enrichment Details - e0b5-599b-4f29-97b7
/* Enriches a given network device context (device id or device Mac Address or device management IP address) with details about the device and neighbor topology


@param GetDeviceEnrichmentDetailsV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-enrichment-details
*/
func (s *DevicesService) GetDeviceEnrichmentDetailsV1(GetDeviceEnrichmentDetailsV1HeaderParams *GetDeviceEnrichmentDetailsV1HeaderParams) (*ResponseDevicesGetDeviceEnrichmentDetailsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/device-enrichment-details"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetDeviceEnrichmentDetailsV1HeaderParams != nil {

		if GetDeviceEnrichmentDetailsV1HeaderParams.EntityType != "" {
			clientRequest = clientRequest.SetHeader("entity_type", GetDeviceEnrichmentDetailsV1HeaderParams.EntityType)
		}

		if GetDeviceEnrichmentDetailsV1HeaderParams.EntityValue != "" {
			clientRequest = clientRequest.SetHeader("entity_value", GetDeviceEnrichmentDetailsV1HeaderParams.EntityValue)
		}

		if GetDeviceEnrichmentDetailsV1HeaderParams.Persistbapioutput != "" {
			clientRequest = clientRequest.SetHeader("__persistbapioutput", GetDeviceEnrichmentDetailsV1HeaderParams.Persistbapioutput)
		}

	}

	response, err = clientRequest.
		SetResult(&ResponseDevicesGetDeviceEnrichmentDetailsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceEnrichmentDetailsV1(GetDeviceEnrichmentDetailsV1HeaderParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceEnrichmentDetailsV1")
	}

	result := response.Result().(*ResponseDevicesGetDeviceEnrichmentDetailsV1)
	return result, response, err

}

//DevicesV1 Devices - 3ab2-bb64-4cca-81ee
/* Intent API for accessing DNA Assurance Device object for generating reports, creating dashboards or creating additional value added services.


@param DevicesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!devices
*/
func (s *DevicesService) DevicesV1(DevicesV1QueryParams *DevicesV1QueryParams) (*ResponseDevicesDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/device-health"

	queryString, _ := query.Values(DevicesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesDevicesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DevicesV1(DevicesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DevicesV1")
	}

	result := response.Result().(*ResponseDevicesDevicesV1)
	return result, response, err

}

//GetPlannedAccessPointsForFloorV1 Get Planned Access Points for Floor - 6780-6977-4589-9a54
/* Provides a list of Planned Access Points for the Floor it is requested for


@param floorID floorId path parameter. The instance UUID of the floor hierarchy element

@param GetPlannedAccessPointsForFloorV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-planned-access-points-for-floor
*/
func (s *DevicesService) GetPlannedAccessPointsForFloorV1(floorID string, GetPlannedAccessPointsForFloorV1QueryParams *GetPlannedAccessPointsForFloorV1QueryParams) (*ResponseDevicesGetPlannedAccessPointsForFloorV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/floors/{floorId}/planned-access-points"
	path = strings.Replace(path, "{floorId}", fmt.Sprintf("%v", floorID), -1)

	queryString, _ := query.Values(GetPlannedAccessPointsForFloorV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetPlannedAccessPointsForFloorV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPlannedAccessPointsForFloorV1(floorID, GetPlannedAccessPointsForFloorV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetPlannedAccessPointsForFloorV1")
	}

	result := response.Result().(*ResponseDevicesGetPlannedAccessPointsForFloorV1)
	return result, response, err

}

//GetAllHealthScoreDefinitionsForGivenFiltersV1 Get all health score definitions for given filters. - 9bb6-ea87-4ffb-b492
/* Get all health score defintions.
Supported filters are id, name and overall health include status. A health score definition can be different across device type. So, deviceType in the query param is important and default is all device types.
By default all supported attributes are listed in response. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-issueAndHealthDefinitions-1.0.0-resolved.yaml


@param GetAllHealthScoreDefinitionsForGivenFiltersV1HeaderParams Custom header parameters
@param GetAllHealthScoreDefinitionsForGivenFiltersV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-all-health-score-definitions-for-given-filters
*/
func (s *DevicesService) GetAllHealthScoreDefinitionsForGivenFiltersV1(GetAllHealthScoreDefinitionsForGivenFiltersV1HeaderParams *GetAllHealthScoreDefinitionsForGivenFiltersV1HeaderParams, GetAllHealthScoreDefinitionsForGivenFiltersV1QueryParams *GetAllHealthScoreDefinitionsForGivenFiltersV1QueryParams) (*ResponseDevicesGetAllHealthScoreDefinitionsForGivenFiltersV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/healthScoreDefinitions"

	queryString, _ := query.Values(GetAllHealthScoreDefinitionsForGivenFiltersV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetAllHealthScoreDefinitionsForGivenFiltersV1HeaderParams != nil {

		if GetAllHealthScoreDefinitionsForGivenFiltersV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetAllHealthScoreDefinitionsForGivenFiltersV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetAllHealthScoreDefinitionsForGivenFiltersV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAllHealthScoreDefinitionsForGivenFiltersV1(GetAllHealthScoreDefinitionsForGivenFiltersV1HeaderParams, GetAllHealthScoreDefinitionsForGivenFiltersV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAllHealthScoreDefinitionsForGivenFiltersV1")
	}

	result := response.Result().(*ResponseDevicesGetAllHealthScoreDefinitionsForGivenFiltersV1)
	return result, response, err

}

//GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersV1 Get the count of health score definitions based on provided filters. - 49aa-bb2c-46ca-b58a
/* Get the count of health score definitions based on provided filters. Supported filters are id, name and overall health include status. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-issueAndHealthDefinitions-1.0.0-resolved.yaml


@param GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersV1HeaderParams Custom header parameters
@param GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-the-count-of-health-score-definitions-based-on-provided-filters
*/
func (s *DevicesService) GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersV1(GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersV1HeaderParams *GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersV1HeaderParams, GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersV1QueryParams *GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersV1QueryParams) (*ResponseDevicesGetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/healthScoreDefinitions/count"

	queryString, _ := query.Values(GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersV1HeaderParams != nil {

		if GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersV1(GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersV1HeaderParams, GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersV1")
	}

	result := response.Result().(*ResponseDevicesGetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersV1)
	return result, response, err

}

//GetHealthScoreDefinitionForTheGivenIDV1 Get health score definition for the given id. - 99b5-d81a-4408-94c3
/* Get health score defintion for the given id. Definition includes all properties from HealthScoreDefinition schema by default. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-issueAndHealthDefinitions-1.0.0-resolved.yaml


@param id id path parameter. Health score definition id.

@param GetHealthScoreDefinitionForTheGivenIdV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-health-score-definition-for-the-given-id
*/
func (s *DevicesService) GetHealthScoreDefinitionForTheGivenIDV1(id string, GetHealthScoreDefinitionForTheGivenIdV1HeaderParams *GetHealthScoreDefinitionForTheGivenIDV1HeaderParams) (*ResponseDevicesGetHealthScoreDefinitionForTheGivenIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/healthScoreDefinitions/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetHealthScoreDefinitionForTheGivenIdV1HeaderParams != nil {

		if GetHealthScoreDefinitionForTheGivenIdV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetHealthScoreDefinitionForTheGivenIdV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetResult(&ResponseDevicesGetHealthScoreDefinitionForTheGivenIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetHealthScoreDefinitionForTheGivenIDV1(id, GetHealthScoreDefinitionForTheGivenIdV1HeaderParams)
		}
		return nil, response, fmt.Errorf("error with operation GetHealthScoreDefinitionForTheGivenIdV1")
	}

	result := response.Result().(*ResponseDevicesGetHealthScoreDefinitionForTheGivenIDV1)
	return result, response, err

}

//GetAllInterfacesV1 Get all interfaces - f594-7a4c-439a-8bf0
/* Returns all available interfaces. This endpoint can return a maximum of 500 interfaces


@param GetAllInterfacesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-all-interfaces
*/
func (s *DevicesService) GetAllInterfacesV1(GetAllInterfacesV1QueryParams *GetAllInterfacesV1QueryParams) (*ResponseDevicesGetAllInterfacesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface"

	queryString, _ := query.Values(GetAllInterfacesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetAllInterfacesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAllInterfacesV1(GetAllInterfacesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAllInterfacesV1")
	}

	result := response.Result().(*ResponseDevicesGetAllInterfacesV1)
	return result, response, err

}

//GetDeviceInterfaceCountForMultipleDevicesV1 Get Device Interface Count for Multiple Devices - 3d92-3b18-4dc9-a4ca
/* Returns the count of interfaces for all devices



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-interface-count-for-multiple-devices
*/
func (s *DevicesService) GetDeviceInterfaceCountForMultipleDevicesV1() (*ResponseDevicesGetDeviceInterfaceCountForMultipleDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceInterfaceCountForMultipleDevicesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceInterfaceCountForMultipleDevicesV1()
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceInterfaceCountForMultipleDevicesV1")
	}

	result := response.Result().(*ResponseDevicesGetDeviceInterfaceCountForMultipleDevicesV1)
	return result, response, err

}

//GetInterfaceByIPV1 Get Interface by IP - cd84-69e6-47ca-ab0e
/* Returns list of interfaces for specified device management IP address


@param ipAddress ipAddress path parameter. IP address of the interface


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-interface-by-ip
*/
func (s *DevicesService) GetInterfaceByIPV1(ipAddress string) (*ResponseDevicesGetInterfaceByIPV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/ip-address/{ipAddress}"
	path = strings.Replace(path, "{ipAddress}", fmt.Sprintf("%v", ipAddress), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetInterfaceByIPV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetInterfaceByIPV1(ipAddress)
		}
		return nil, response, fmt.Errorf("error with operation GetInterfaceByIpV1")
	}

	result := response.Result().(*ResponseDevicesGetInterfaceByIPV1)
	return result, response, err

}

//GetIsisInterfacesV1 Get ISIS interfaces - 84ad-8b0e-42ca-b48a
/* Returns the interfaces that has ISIS enabled



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-isis-interfaces
*/
func (s *DevicesService) GetIsisInterfacesV1() (*ResponseDevicesGetIsisInterfacesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/isis"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetIsisInterfacesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetIsisInterfacesV1()
		}
		return nil, response, fmt.Errorf("error with operation GetIsisInterfacesV1")
	}

	result := response.Result().(*ResponseDevicesGetIsisInterfacesV1)
	return result, response, err

}

//GetInterfaceInfoByIDV1 Get Interface info by Id - ba9d-c85b-4b8a-9a17
/* Returns list of interfaces by specified device


@param deviceID deviceId path parameter. Device ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-interface-info-by-id
*/
func (s *DevicesService) GetInterfaceInfoByIDV1(deviceID string) (*ResponseDevicesGetInterfaceInfoByIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/network-device/{deviceId}"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetInterfaceInfoByIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetInterfaceInfoByIDV1(deviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetInterfaceInfoByIdV1")
	}

	result := response.Result().(*ResponseDevicesGetInterfaceInfoByIDV1)
	return result, response, err

}

//GetDeviceInterfaceCountV1 Get Device Interface count - 5b86-3922-4cd8-8ea7
/* Returns the interface count for the given device


@param deviceID deviceId path parameter. Device ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-interface-count
*/
func (s *DevicesService) GetDeviceInterfaceCountV1(deviceID string) (*ResponseDevicesGetDeviceInterfaceCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/network-device/{deviceId}/count"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceInterfaceCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceInterfaceCountV1(deviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceInterfaceCountV1")
	}

	result := response.Result().(*ResponseDevicesGetDeviceInterfaceCountV1)
	return result, response, err

}

//GetInterfaceDetailsByDeviceIDAndInterfaceNameV1 Get Interface details by device Id and interface name - 4eb5-6a61-4cc9-a2d2
/* Returns interface by specified device Id and interface name


@param deviceID deviceId path parameter. Device ID

@param GetInterfaceDetailsByDeviceIdAndInterfaceNameV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-interface-details-by-device-id-and-interface-name
*/
func (s *DevicesService) GetInterfaceDetailsByDeviceIDAndInterfaceNameV1(deviceID string, GetInterfaceDetailsByDeviceIdAndInterfaceNameV1QueryParams *GetInterfaceDetailsByDeviceIDAndInterfaceNameV1QueryParams) (*ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/network-device/{deviceId}/interface-name"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	queryString, _ := query.Values(GetInterfaceDetailsByDeviceIdAndInterfaceNameV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetInterfaceDetailsByDeviceIDAndInterfaceNameV1(deviceID, GetInterfaceDetailsByDeviceIdAndInterfaceNameV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetInterfaceDetailsByDeviceIdAndInterfaceNameV1")
	}

	result := response.Result().(*ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameV1)
	return result, response, err

}

//GetDeviceInterfacesBySpecifiedRangeV1 Get Device Interfaces by specified range - 349c-8884-43b8-9a58
/* Returns the list of interfaces for the device for the specified range


@param deviceID deviceId path parameter. Device ID

@param startIndex startIndex path parameter. Start index

@param recordsToReturn recordsToReturn path parameter. Number of records to return


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-interfaces-by-specified-range
*/
func (s *DevicesService) GetDeviceInterfacesBySpecifiedRangeV1(deviceID string, startIndex int, recordsToReturn int) (*ResponseDevicesGetDeviceInterfacesBySpecifiedRangeV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/network-device/{deviceId}/{startIndex}/{recordsToReturn}"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)
	path = strings.Replace(path, "{startIndex}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{recordsToReturn}", fmt.Sprintf("%v", recordsToReturn), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceInterfacesBySpecifiedRangeV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceInterfacesBySpecifiedRangeV1(deviceID, startIndex, recordsToReturn)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceInterfacesBySpecifiedRangeV1")
	}

	result := response.Result().(*ResponseDevicesGetDeviceInterfacesBySpecifiedRangeV1)
	return result, response, err

}

//GetOspfInterfacesV1 Get OSPF interfaces - 70ad-3976-49e9-b4d3
/* Returns the interfaces that has OSPF enabled



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-ospf-interfaces
*/
func (s *DevicesService) GetOspfInterfacesV1() (*ResponseDevicesGetOspfInterfacesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/ospf"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetOspfInterfacesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetOspfInterfacesV1()
		}
		return nil, response, fmt.Errorf("error with operation GetOspfInterfacesV1")
	}

	result := response.Result().(*ResponseDevicesGetOspfInterfacesV1)
	return result, response, err

}

//GetInterfaceByIDV1 Get Interface by Id - b888-792d-43ba-ba46
/* Returns the interface for the given interface ID


@param id id path parameter. Interface ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-interface-by-id
*/
func (s *DevicesService) GetInterfaceByIDV1(id string) (*ResponseDevicesGetInterfaceByIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetInterfaceByIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetInterfaceByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation GetInterfaceByIdV1")
	}

	result := response.Result().(*ResponseDevicesGetInterfaceByIDV1)
	return result, response, err

}

//LegitOperationsForInterfaceV1 Legit operations for interface - 87a3-3a52-46ea-a40e
/* Get list of all properties & operations valid for an interface.


@param interfaceUUID interfaceUuid path parameter. Interface ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!legit-operations-for-interface
*/
func (s *DevicesService) LegitOperationsForInterfaceV1(interfaceUUID string) (*ResponseDevicesLegitOperationsForInterfaceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/{interfaceUuid}/legit-operation"
	path = strings.Replace(path, "{interfaceUuid}", fmt.Sprintf("%v", interfaceUUID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesLegitOperationsForInterfaceV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.LegitOperationsForInterfaceV1(interfaceUUID)
		}
		return nil, response, fmt.Errorf("error with operation LegitOperationsForInterfaceV1")
	}

	result := response.Result().(*ResponseDevicesLegitOperationsForInterfaceV1)
	return result, response, err

}

//GetDeviceListV1 Get Device list - 20b1-9b52-464b-8972
/* Returns list of network devices based on filter criteria such as management IP address, mac address, hostname, etc. You can use the .* in any value to conduct a wildcard search. For example, to find all hostnames beginning with myhost in the IP address range 192.25.18.n, issue the following request: GET /dna/intent/api/v1/network-device?hostname=myhost.*&managementIpAddress=192.25.18..*
If id parameter is provided with comma separated ids, it will return the list of network-devices for the given ids and ignores the other request parameters. You can also specify offset & limit to get the required list.


@param GetDeviceListV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-list
*/
func (s *DevicesService) GetDeviceListV1(GetDeviceListV1QueryParams *GetDeviceListV1QueryParams) (*ResponseDevicesGetDeviceListV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device"

	queryString, _ := query.Values(GetDeviceListV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetDeviceListV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceListV1(GetDeviceListV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceListV1")
	}

	result := response.Result().(*ResponseDevicesGetDeviceListV1)
	return result, response, err

}

//GetDeviceValuesThatMatchFullyOrPartiallyAnAttributeV1 Get Device Values that match fully or partially an Attribute - ffa7-48cc-44e9-a437
/* Returns the list of values of the first given required parameter. You can use the .* in any value to conduct a wildcard search. For example, to get all the devices with the management IP address starting with 10.10. , issue the following request: GET /dna/inten/api/v1/network-device/autocomplete?managementIpAddress=10.10..* It will return the device management IP addresses that match fully or partially the provided attribute. {[10.10.1.1, 10.10.20.2, …]}.


@param GetDeviceValuesThatMatchFullyOrPartiallyAnAttributeV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-values-that-match-fully-or-partially-an-attribute
*/
func (s *DevicesService) GetDeviceValuesThatMatchFullyOrPartiallyAnAttributeV1(GetDeviceValuesThatMatchFullyOrPartiallyAnAttributeV1QueryParams *GetDeviceValuesThatMatchFullyOrPartiallyAnAttributeV1QueryParams) (*ResponseDevicesGetDeviceValuesThatMatchFullyOrPartiallyAnAttributeV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/autocomplete"

	queryString, _ := query.Values(GetDeviceValuesThatMatchFullyOrPartiallyAnAttributeV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetDeviceValuesThatMatchFullyOrPartiallyAnAttributeV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceValuesThatMatchFullyOrPartiallyAnAttributeV1(GetDeviceValuesThatMatchFullyOrPartiallyAnAttributeV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceValuesThatMatchFullyOrPartiallyAnAttributeV1")
	}

	result := response.Result().(*ResponseDevicesGetDeviceValuesThatMatchFullyOrPartiallyAnAttributeV1)
	return result, response, err

}

//GetPollingIntervalForAllDevicesV1 Get Polling Interval for all devices - 38bd-0b88-4b89-a785
/* Returns polling interval of all devices



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-polling-interval-for-all-devices
*/
func (s *DevicesService) GetPollingIntervalForAllDevicesV1() (*ResponseDevicesGetPollingIntervalForAllDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/collection-schedule/global"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetPollingIntervalForAllDevicesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPollingIntervalForAllDevicesV1()
		}
		return nil, response, fmt.Errorf("error with operation GetPollingIntervalForAllDevicesV1")
	}

	result := response.Result().(*ResponseDevicesGetPollingIntervalForAllDevicesV1)
	return result, response, err

}

//GetDeviceConfigForAllDevicesV1 Get Device Config for all devices - b7bc-aa08-4e2b-90d0
/* Returns the config for all devices. This API has been deprecated and will not be available in a Cisco Catalyst Center release after Nov 1st 2024 23:59:59 GMT.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-config-for-all-devices
*/
func (s *DevicesService) GetDeviceConfigForAllDevicesV1() (*ResponseDevicesGetDeviceConfigForAllDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/config"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceConfigForAllDevicesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceConfigForAllDevicesV1()
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceConfigForAllDevicesV1")
	}

	result := response.Result().(*ResponseDevicesGetDeviceConfigForAllDevicesV1)
	return result, response, err

}

//GetDeviceConfigCountV1 Get Device Config Count - 888f-585c-49b8-8441
/* Returns the count of device configs



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-config-count
*/
func (s *DevicesService) GetDeviceConfigCountV1() (*ResponseDevicesGetDeviceConfigCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/config/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceConfigCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceConfigCountV1()
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceConfigCountV1")
	}

	result := response.Result().(*ResponseDevicesGetDeviceConfigCountV1)
	return result, response, err

}

//GetDeviceCountKnowYourNetworkV1 Get Device Count - 5db2-1b8e-43fa-b7d8
/* Returns the count of network devices based on the filter criteria by management IP address, mac address, hostname and location name


@param GetDeviceCountKnowYourNetworkV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-count-know-your-network
*/
func (s *DevicesService) GetDeviceCountKnowYourNetworkV1(GetDeviceCountKnowYourNetworkV1QueryParams *GetDeviceCountKnowYourNetworkV1QueryParams) (*ResponseDevicesGetDeviceCountKnowYourNetworkV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/count"

	queryString, _ := query.Values(GetDeviceCountKnowYourNetworkV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetDeviceCountKnowYourNetworkV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceCountKnowYourNetworkV1(GetDeviceCountKnowYourNetworkV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceCountKnowYourNetworkV1")
	}

	result := response.Result().(*ResponseDevicesGetDeviceCountKnowYourNetworkV1)
	return result, response, err

}

//GetFunctionalCapabilityForDevicesV1 Get Functional Capability for devices - c3b3-c9ef-4e6b-8a09
/* Returns the functional-capability for given devices


@param GetFunctionalCapabilityForDevicesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-functional-capability-for-devices
*/
func (s *DevicesService) GetFunctionalCapabilityForDevicesV1(GetFunctionalCapabilityForDevicesV1QueryParams *GetFunctionalCapabilityForDevicesV1QueryParams) (*ResponseDevicesGetFunctionalCapabilityForDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/functional-capability"

	queryString, _ := query.Values(GetFunctionalCapabilityForDevicesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetFunctionalCapabilityForDevicesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFunctionalCapabilityForDevicesV1(GetFunctionalCapabilityForDevicesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetFunctionalCapabilityForDevicesV1")
	}

	result := response.Result().(*ResponseDevicesGetFunctionalCapabilityForDevicesV1)
	return result, response, err

}

//GetFunctionalCapabilityByIDV1 Get Functional Capability by Id - 81bb-4804-405a-8d2f
/* Returns functional capability with given Id


@param id id path parameter. Functional Capability UUID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-functional-capability-by-id
*/
func (s *DevicesService) GetFunctionalCapabilityByIDV1(id string) (*ResponseDevicesGetFunctionalCapabilityByIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/functional-capability/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetFunctionalCapabilityByIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFunctionalCapabilityByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation GetFunctionalCapabilityByIdV1")
	}

	result := response.Result().(*ResponseDevicesGetFunctionalCapabilityByIDV1)
	return result, response, err

}

//InventoryInsightDeviceLinkMismatchAPIV1 Inventory Insight Device Link Mismatch API - 5792-59d8-4208-8190
/* Find all devices with link mismatch (speed /  vlan)


@param siteID siteId path parameter.
@param InventoryInsightDeviceLinkMismatchAPIV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!inventory-insight-device-link-mismatch-api
*/
func (s *DevicesService) InventoryInsightDeviceLinkMismatchAPIV1(siteID string, InventoryInsightDeviceLinkMismatchAPIV1QueryParams *InventoryInsightDeviceLinkMismatchAPIV1QueryParams) (*ResponseDevicesInventoryInsightDeviceLinkMismatchAPIV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/insight/{siteId}/device-link"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	queryString, _ := query.Values(InventoryInsightDeviceLinkMismatchAPIV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesInventoryInsightDeviceLinkMismatchAPIV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.InventoryInsightDeviceLinkMismatchAPIV1(siteID, InventoryInsightDeviceLinkMismatchAPIV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation InventoryInsightDeviceLinkMismatchApiV1")
	}

	result := response.Result().(*ResponseDevicesInventoryInsightDeviceLinkMismatchAPIV1)
	return result, response, err

}

//GetNetworkDeviceByIPV1 Get Network Device by IP - d0a4-b881-45aa-bb51
/* Returns the network device by specified IP address


@param ipAddress ipAddress path parameter. Device IP address


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-device-by-ip
*/
func (s *DevicesService) GetNetworkDeviceByIPV1(ipAddress string) (*ResponseDevicesGetNetworkDeviceByIPV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/ip-address/{ipAddress}"
	path = strings.Replace(path, "{ipAddress}", fmt.Sprintf("%v", ipAddress), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetNetworkDeviceByIPV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetNetworkDeviceByIPV1(ipAddress)
		}
		return nil, response, fmt.Errorf("error with operation GetNetworkDeviceByIpV1")
	}

	result := response.Result().(*ResponseDevicesGetNetworkDeviceByIPV1)
	return result, response, err

}

//GetModulesV1 Get Modules - eb82-49e3-4f69-b0f1
/* Returns modules by specified device id


@param GetModulesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-modules
*/
func (s *DevicesService) GetModulesV1(GetModulesV1QueryParams *GetModulesV1QueryParams) (*ResponseDevicesGetModulesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/module"

	queryString, _ := query.Values(GetModulesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetModulesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetModulesV1(GetModulesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetModulesV1")
	}

	result := response.Result().(*ResponseDevicesGetModulesV1)
	return result, response, err

}

//GetModuleCountV1 Get Module count - 8db9-3974-4649-a782
/* Returns Module Count


@param GetModuleCountV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-module-count
*/
func (s *DevicesService) GetModuleCountV1(GetModuleCountV1QueryParams *GetModuleCountV1QueryParams) (*ResponseDevicesGetModuleCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/module/count"

	queryString, _ := query.Values(GetModuleCountV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetModuleCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetModuleCountV1(GetModuleCountV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetModuleCountV1")
	}

	result := response.Result().(*ResponseDevicesGetModuleCountV1)
	return result, response, err

}

//GetModuleInfoByIDV1 Get Module Info by Id - 0db7-da74-4c0b-83d8
/* Returns Module info by 'module id'


@param id id path parameter. Module id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-module-info-by-id
*/
func (s *DevicesService) GetModuleInfoByIDV1(id string) (*ResponseDevicesGetModuleInfoByIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/module/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetModuleInfoByIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetModuleInfoByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation GetModuleInfoByIdV1")
	}

	result := response.Result().(*ResponseDevicesGetModuleInfoByIDV1)
	return result, response, err

}

//GetDeviceBySerialNumberV1 Get Device by Serial number - d888-ab6d-4d59-a8c1
/* Returns the network device with given serial number


@param serialNumber serialNumber path parameter. Device serial number


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-by-serial-number
*/
func (s *DevicesService) GetDeviceBySerialNumberV1(serialNumber string) (*ResponseDevicesGetDeviceBySerialNumberV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/serial-number/{serialNumber}"
	path = strings.Replace(path, "{serialNumber}", fmt.Sprintf("%v", serialNumber), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceBySerialNumberV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceBySerialNumberV1(serialNumber)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceBySerialNumberV1")
	}

	result := response.Result().(*ResponseDevicesGetDeviceBySerialNumberV1)
	return result, response, err

}

//GetDevicesRegisteredForWsaNotificationV1 Get Devices registered for WSA Notification - c980-9b67-44f8-a502
/* It fetches devices which are registered to receive WSA notifications. The device serial number and/or MAC address are required to be provided as query parameters.


@param GetDevicesRegisteredForWSANotificationV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-devices-registered-for-wsa-notification
*/
func (s *DevicesService) GetDevicesRegisteredForWsaNotificationV1(GetDevicesRegisteredForWSANotificationV1QueryParams *GetDevicesRegisteredForWsaNotificationV1QueryParams) (*ResponseDevicesGetDevicesRegisteredForWsaNotificationV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/tenantinfo/macaddress"

	queryString, _ := query.Values(GetDevicesRegisteredForWSANotificationV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetDevicesRegisteredForWsaNotificationV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDevicesRegisteredForWsaNotificationV1(GetDevicesRegisteredForWSANotificationV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDevicesRegisteredForWsaNotificationV1")
	}

	result := response.Result().(*ResponseDevicesGetDevicesRegisteredForWsaNotificationV1)
	return result, response, err

}

//GetAllUserDefinedFieldsV1 Get All User-Defined-Fields - 058d-2a92-4899-b7bb
/* Gets existing global User Defined Fields. If no input is given, it fetches ALL the Global UDFs. Filter/search is supported by UDF Id(s) or UDF name(s) or both.


@param GetAllUserDefinedFieldsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-all-user-defined-fields
*/
func (s *DevicesService) GetAllUserDefinedFieldsV1(GetAllUserDefinedFieldsV1QueryParams *GetAllUserDefinedFieldsV1QueryParams) (*ResponseDevicesGetAllUserDefinedFieldsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/user-defined-field"

	queryString, _ := query.Values(GetAllUserDefinedFieldsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetAllUserDefinedFieldsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAllUserDefinedFieldsV1(GetAllUserDefinedFieldsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAllUserDefinedFieldsV1")
	}

	result := response.Result().(*ResponseDevicesGetAllUserDefinedFieldsV1)
	return result, response, err

}

//GetChassisDetailsForDeviceV1 Get Chassis Details for Device - 0486-9b26-49ab-b579
/* Returns chassis details for given device ID


@param deviceID deviceId path parameter. Device ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-chassis-details-for-device
*/
func (s *DevicesService) GetChassisDetailsForDeviceV1(deviceID string) (*ResponseDevicesGetChassisDetailsForDeviceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceId}/chassis"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetChassisDetailsForDeviceV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetChassisDetailsForDeviceV1(deviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetChassisDetailsForDeviceV1")
	}

	result := response.Result().(*ResponseDevicesGetChassisDetailsForDeviceV1)
	return result, response, err

}

//GetStackDetailsForDeviceV1 Get Stack Details for Device - 78a7-7ab0-4d5a-8a10
/* Retrieves complete stack details for given device ID


@param deviceID deviceId path parameter. Device ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-stack-details-for-device
*/
func (s *DevicesService) GetStackDetailsForDeviceV1(deviceID string) (*ResponseDevicesGetStackDetailsForDeviceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceId}/stack"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetStackDetailsForDeviceV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetStackDetailsForDeviceV1(deviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetStackDetailsForDeviceV1")
	}

	result := response.Result().(*ResponseDevicesGetStackDetailsForDeviceV1)
	return result, response, err

}

//GetTheDetailsOfPhysicalComponentsOfTheGivenDeviceV1 Get the Details of Physical Components of the Given Device. - 20b1-9b52-464b-897a
/* Return all types of equipment details like PowerSupply, Fan, Chassis, Backplane, Module, PROCESSOR, Other and SFP for the Given device.


@param deviceUUID deviceUuid path parameter.
@param GetTheDetailsOfPhysicalComponentsOfTheGivenDeviceV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-the-details-of-physical-components-of-the-given-device
*/
func (s *DevicesService) GetTheDetailsOfPhysicalComponentsOfTheGivenDeviceV1(deviceUUID string, GetTheDetailsOfPhysicalComponentsOfTheGivenDeviceV1QueryParams *GetTheDetailsOfPhysicalComponentsOfTheGivenDeviceV1QueryParams) (*ResponseDevicesGetTheDetailsOfPhysicalComponentsOfTheGivenDeviceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceUuid}/equipment"
	path = strings.Replace(path, "{deviceUuid}", fmt.Sprintf("%v", deviceUUID), -1)

	queryString, _ := query.Values(GetTheDetailsOfPhysicalComponentsOfTheGivenDeviceV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetTheDetailsOfPhysicalComponentsOfTheGivenDeviceV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTheDetailsOfPhysicalComponentsOfTheGivenDeviceV1(deviceUUID, GetTheDetailsOfPhysicalComponentsOfTheGivenDeviceV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTheDetailsOfPhysicalComponentsOfTheGivenDeviceV1")
	}

	result := response.Result().(*ResponseDevicesGetTheDetailsOfPhysicalComponentsOfTheGivenDeviceV1)
	return result, response, err

}

//ReturnsPoeInterfaceDetailsForTheDeviceV1 Returns POE interface details for the device. - 20b5-48af-42da-a337
/* Returns POE interface details for the device, where deviceuuid is mandatory & accepts comma seperated interface names which is optional and returns information for that particular interfaces where(operStatus = operationalStatus)


@param deviceUUID deviceUuid path parameter. uuid of the device

@param ReturnsPOEInterfaceDetailsForTheDeviceV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!returns-poe-interface-details-for-the-device
*/
func (s *DevicesService) ReturnsPoeInterfaceDetailsForTheDeviceV1(deviceUUID string, ReturnsPOEInterfaceDetailsForTheDeviceV1QueryParams *ReturnsPoeInterfaceDetailsForTheDeviceV1QueryParams) (*ResponseDevicesReturnsPoeInterfaceDetailsForTheDeviceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceUuid}/interface/poe-detail"
	path = strings.Replace(path, "{deviceUuid}", fmt.Sprintf("%v", deviceUUID), -1)

	queryString, _ := query.Values(ReturnsPOEInterfaceDetailsForTheDeviceV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesReturnsPoeInterfaceDetailsForTheDeviceV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReturnsPoeInterfaceDetailsForTheDeviceV1(deviceUUID, ReturnsPOEInterfaceDetailsForTheDeviceV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReturnsPoeInterfaceDetailsForTheDeviceV1")
	}

	result := response.Result().(*ResponseDevicesReturnsPoeInterfaceDetailsForTheDeviceV1)
	return result, response, err

}

//GetConnectedDeviceDetailV1 Get connected device detail - a8aa-ca21-4c09-8388
/* Get connected device detail for given deviceUuid and interfaceUuid


@param deviceUUID deviceUuid path parameter. instanceuuid of Device

@param interfaceUUID interfaceUuid path parameter. instanceuuid of interface


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-connected-device-detail
*/
func (s *DevicesService) GetConnectedDeviceDetailV1(deviceUUID string, interfaceUUID string) (*ResponseDevicesGetConnectedDeviceDetailV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceUuid}/interface/{interfaceUuid}/neighbor"
	path = strings.Replace(path, "{deviceUuid}", fmt.Sprintf("%v", deviceUUID), -1)
	path = strings.Replace(path, "{interfaceUuid}", fmt.Sprintf("%v", interfaceUUID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetConnectedDeviceDetailV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetConnectedDeviceDetailV1(deviceUUID, interfaceUUID)
		}
		return nil, response, fmt.Errorf("error with operation GetConnectedDeviceDetailV1")
	}

	result := response.Result().(*ResponseDevicesGetConnectedDeviceDetailV1)
	return result, response, err

}

//GetLinecardDetailsV1 Get Linecard details - 46a1-4b02-48fb-8fbf
/* Get line card detail for a given deviceuuid.  Response will contain serial no, part no, switch no and slot no.


@param deviceUUID deviceUuid path parameter. instanceuuid of device


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-linecard-details
*/
func (s *DevicesService) GetLinecardDetailsV1(deviceUUID string) (*ResponseDevicesGetLinecardDetailsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceUuid}/line-card"
	path = strings.Replace(path, "{deviceUuid}", fmt.Sprintf("%v", deviceUUID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetLinecardDetailsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetLinecardDetailsV1(deviceUUID)
		}
		return nil, response, fmt.Errorf("error with operation GetLinecardDetailsV1")
	}

	result := response.Result().(*ResponseDevicesGetLinecardDetailsV1)
	return result, response, err

}

//PoeDetailsV1 POE details  - 8ba6-7932-4ed9-abae
/* Returns POE details for device.


@param deviceUUID deviceUuid path parameter. UUID of the device


Documentation Link: https://developer.cisco.com/docs/dna-center/#!poe-details
*/
func (s *DevicesService) PoeDetailsV1(deviceUUID string) (*ResponseDevicesPoeDetailsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceUuid}/poe"
	path = strings.Replace(path, "{deviceUuid}", fmt.Sprintf("%v", deviceUUID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesPoeDetailsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.PoeDetailsV1(deviceUUID)
		}
		return nil, response, fmt.Errorf("error with operation PoeDetailsV1")
	}

	result := response.Result().(*ResponseDevicesPoeDetailsV1)
	return result, response, err

}

//GetSupervisorCardDetailV1 Get Supervisor card detail - 88aa-1b52-4a38-bf97
/* Get supervisor card detail for a given deviceuuid. Response will contain serial no, part no, switch no and slot no.


@param deviceUUID deviceUuid path parameter. instanceuuid of device


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-supervisor-card-detail
*/
func (s *DevicesService) GetSupervisorCardDetailV1(deviceUUID string) (*ResponseDevicesGetSupervisorCardDetailV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceUuid}/supervisor-card"
	path = strings.Replace(path, "{deviceUuid}", fmt.Sprintf("%v", deviceUUID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetSupervisorCardDetailV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSupervisorCardDetailV1(deviceUUID)
		}
		return nil, response, fmt.Errorf("error with operation GetSupervisorCardDetailV1")
	}

	result := response.Result().(*ResponseDevicesGetSupervisorCardDetailV1)
	return result, response, err

}

//GetDeviceByIDV1 Get Device by ID - 8fa8-eb40-4a4a-8d96
/* Returns the network device details for the given device ID


@param id id path parameter. Device ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-by-id
*/
func (s *DevicesService) GetDeviceByIDV1(id string) (*ResponseDevicesGetDeviceByIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceByIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceByIdV1")
	}

	result := response.Result().(*ResponseDevicesGetDeviceByIDV1)
	return result, response, err

}

//GetDeviceSummaryV1 Get Device Summary - 819f-9aa5-4fea-b7bf
/* Returns brief summary of device info for the given device Id


@param id id path parameter. Device ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-summary
*/
func (s *DevicesService) GetDeviceSummaryV1(id string) (*ResponseDevicesGetDeviceSummaryV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{id}/brief"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceSummaryV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceSummaryV1(id)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceSummaryV1")
	}

	result := response.Result().(*ResponseDevicesGetDeviceSummaryV1)
	return result, response, err

}

//GetPollingIntervalByIDV1 Get Polling Interval by Id - 8291-8a1b-4d28-9c5c
/* Returns polling interval by device id


@param id id path parameter. Device ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-polling-interval-by-id
*/
func (s *DevicesService) GetPollingIntervalByIDV1(id string) (*ResponseDevicesGetPollingIntervalByIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{id}/collection-schedule"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetPollingIntervalByIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPollingIntervalByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation GetPollingIntervalByIdV1")
	}

	result := response.Result().(*ResponseDevicesGetPollingIntervalByIDV1)
	return result, response, err

}

//GetOrganizationListForMerakiV1 Get Organization list for Meraki - 84b3-7ae5-4c59-ab28
/* Returns list of organizations for meraki dashboard


@param id id path parameter. Device Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-list-for-meraki
*/
func (s *DevicesService) GetOrganizationListForMerakiV1(id string) (*ResponseDevicesGetOrganizationListForMerakiV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{id}/meraki-organization"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetOrganizationListForMerakiV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetOrganizationListForMerakiV1(id)
		}
		return nil, response, fmt.Errorf("error with operation GetOrganizationListForMerakiV1")
	}

	result := response.Result().(*ResponseDevicesGetOrganizationListForMerakiV1)
	return result, response, err

}

//GetDeviceInterfaceVLANsV1 Get Device Interface VLANs - 288d-f949-4f2a-9746
/* Returns Device Interface VLANs. If parameter value is null or empty, it won't return any value in response.


@param id id path parameter.
@param GetDeviceInterfaceVLANsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-interface-vlans
*/
func (s *DevicesService) GetDeviceInterfaceVLANsV1(id string, GetDeviceInterfaceVLANsV1QueryParams *GetDeviceInterfaceVLANsV1QueryParams) (*ResponseDevicesGetDeviceInterfaceVLANsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{id}/vlan"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetDeviceInterfaceVLANsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetDeviceInterfaceVLANsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceInterfaceVLANsV1(id, GetDeviceInterfaceVLANsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceInterfaceVlansV1")
	}

	result := response.Result().(*ResponseDevicesGetDeviceInterfaceVLANsV1)
	return result, response, err

}

//GetWirelessLanControllerDetailsByIDV1 Get wireless lan controller details by Id - f682-6a8e-41bb-a242
/* Returns the wireless lan controller info with given device ID


@param id id path parameter. Device ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-wireless-lan-controller-details-by-id
*/
func (s *DevicesService) GetWirelessLanControllerDetailsByIDV1(id string) (*ResponseDevicesGetWirelessLanControllerDetailsByIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{id}/wireless-info"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetWirelessLanControllerDetailsByIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetWirelessLanControllerDetailsByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation GetWirelessLanControllerDetailsByIdV1")
	}

	result := response.Result().(*ResponseDevicesGetWirelessLanControllerDetailsByIDV1)
	return result, response, err

}

//GetDeviceConfigByIDV1 Get Device Config by Id - 84b3-3a9e-480a-bcaf
/* Returns the device config by specified device ID


@param networkDeviceID networkDeviceId path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-config-by-id
*/
func (s *DevicesService) GetDeviceConfigByIDV1(networkDeviceID string) (*ResponseDevicesGetDeviceConfigByIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{networkDeviceId}/config"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceConfigByIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceConfigByIDV1(networkDeviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceConfigByIdV1")
	}

	result := response.Result().(*ResponseDevicesGetDeviceConfigByIDV1)
	return result, response, err

}

//GetNetworkDeviceByPaginationRangeV1 Get Network Device by pagination range - f495-48c5-4be8-a3e2
/* Returns the list of network devices for the given pagination range. The maximum number of records that can be retrieved is 500


@param startIndex startIndex path parameter. Start index [>=1]

@param recordsToReturn recordsToReturn path parameter. Number of records to return [1<= recordsToReturn <= 500]


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-device-by-pagination-range
*/
func (s *DevicesService) GetNetworkDeviceByPaginationRangeV1(startIndex int, recordsToReturn int) (*ResponseDevicesGetNetworkDeviceByPaginationRangeV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{startIndex}/{recordsToReturn}"
	path = strings.Replace(path, "{startIndex}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{recordsToReturn}", fmt.Sprintf("%v", recordsToReturn), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetNetworkDeviceByPaginationRangeV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetNetworkDeviceByPaginationRangeV1(startIndex, recordsToReturn)
		}
		return nil, response, fmt.Errorf("error with operation GetNetworkDeviceByPaginationRangeV1")
	}

	result := response.Result().(*ResponseDevicesGetNetworkDeviceByPaginationRangeV1)
	return result, response, err

}

//RetrieveScheduledMaintenanceWindowsForNetworkDevicesV1 Retrieve scheduled maintenance windows for network devices - 7d9e-198c-4a9b-8971
/* This API retrieves a list of scheduled maintenance windows for network devices based on filter parameters. Each maintenance window is composed of a start schedule and end schedule, both of which have unique identifiers(`startId` and `endId`). These identifiers can be used to fetch the status of the start schedule and end schedule using the `GET /dna/intent/api/v1/activities/{id}` API. Completed maintenance schedules are automatically removed from the system after two weeks.


@param RetrieveScheduledMaintenanceWindowsForNetworkDevicesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-scheduled-maintenance-windows-for-network-devices
*/
func (s *DevicesService) RetrieveScheduledMaintenanceWindowsForNetworkDevicesV1(RetrieveScheduledMaintenanceWindowsForNetworkDevicesV1QueryParams *RetrieveScheduledMaintenanceWindowsForNetworkDevicesV1QueryParams) (*ResponseDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDeviceMaintenanceSchedules"

	queryString, _ := query.Values(RetrieveScheduledMaintenanceWindowsForNetworkDevicesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveScheduledMaintenanceWindowsForNetworkDevicesV1(RetrieveScheduledMaintenanceWindowsForNetworkDevicesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveScheduledMaintenanceWindowsForNetworkDevicesV1")
	}

	result := response.Result().(*ResponseDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesV1)
	return result, response, err

}

//RetrieveTheTotalNumberOfScheduledMaintenanceWindowsV1 Retrieve the total number of scheduled maintenance windows - 6981-69d1-44e8-9284
/* Retrieve the total count of all scheduled maintenance windows for network devices.


@param RetrieveTheTotalNumberOfScheduledMaintenanceWindowsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-the-total-number-of-scheduled-maintenance-windows
*/
func (s *DevicesService) RetrieveTheTotalNumberOfScheduledMaintenanceWindowsV1(RetrieveTheTotalNumberOfScheduledMaintenanceWindowsV1QueryParams *RetrieveTheTotalNumberOfScheduledMaintenanceWindowsV1QueryParams) (*ResponseDevicesRetrieveTheTotalNumberOfScheduledMaintenanceWindowsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDeviceMaintenanceSchedules/count"

	queryString, _ := query.Values(RetrieveTheTotalNumberOfScheduledMaintenanceWindowsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesRetrieveTheTotalNumberOfScheduledMaintenanceWindowsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveTheTotalNumberOfScheduledMaintenanceWindowsV1(RetrieveTheTotalNumberOfScheduledMaintenanceWindowsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveTheTotalNumberOfScheduledMaintenanceWindowsV1")
	}

	result := response.Result().(*ResponseDevicesRetrieveTheTotalNumberOfScheduledMaintenanceWindowsV1)
	return result, response, err

}

//RetrievesTheMaintenanceScheduleInformationV1 Retrieves the maintenance schedule information. - 5fb8-487f-4248-9f71
/* API to retrieve the maintenance schedule information for the given id.


@param id id path parameter. Unique identifier for the maintenance schedule


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-maintenance-schedule-information
*/
func (s *DevicesService) RetrievesTheMaintenanceScheduleInformationV1(id string) (*ResponseDevicesRetrievesTheMaintenanceScheduleInformationV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDeviceMaintenanceSchedules/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesRetrievesTheMaintenanceScheduleInformationV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheMaintenanceScheduleInformationV1(id)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheMaintenanceScheduleInformationV1")
	}

	result := response.Result().(*ResponseDevicesRetrievesTheMaintenanceScheduleInformationV1)
	return result, response, err

}

//RetrieveNetworkDevicesV1 Retrieve network devices - 9e97-b8b2-4539-86a3
/* API to fetch the list of network devices using basic filters. Use the `/dna/intent/api/v1/networkDevices/query` API for advanced filtering. Refer features for more details.


@param RetrieveNetworkDevicesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-network-devices
*/
func (s *DevicesService) RetrieveNetworkDevicesV1(RetrieveNetworkDevicesV1QueryParams *RetrieveNetworkDevicesV1QueryParams) (*ResponseDevicesRetrieveNetworkDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices"

	queryString, _ := query.Values(RetrieveNetworkDevicesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesRetrieveNetworkDevicesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveNetworkDevicesV1(RetrieveNetworkDevicesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveNetworkDevicesV1")
	}

	result := response.Result().(*ResponseDevicesRetrieveNetworkDevicesV1)
	return result, response, err

}

//CountTheNumberOfNetworkDevicesV1 Count the number of network devices - b988-e961-493b-8e72
/* API to fetch the count of network devices using basic filters. Use the `/dna/intent/api/v1/networkDevices/query/count` API if you need advanced filtering.


@param CountTheNumberOfNetworkDevicesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!count-the-number-of-network-devices
*/
func (s *DevicesService) CountTheNumberOfNetworkDevicesV1(CountTheNumberOfNetworkDevicesV1QueryParams *CountTheNumberOfNetworkDevicesV1QueryParams) (*ResponseDevicesCountTheNumberOfNetworkDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/count"

	queryString, _ := query.Values(CountTheNumberOfNetworkDevicesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesCountTheNumberOfNetworkDevicesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.CountTheNumberOfNetworkDevicesV1(CountTheNumberOfNetworkDevicesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation CountTheNumberOfNetworkDevicesV1")
	}

	result := response.Result().(*ResponseDevicesCountTheNumberOfNetworkDevicesV1)
	return result, response, err

}

//GetDetailsOfASingleNetworkDeviceV1 Get details of a single network device - 289a-88ec-4e49-a477
/* API to fetch the details of network device using the `id`. Use the `/dna/intent/api/v1/networkDevices/query` API for advanced filtering. The API supports views to fetch only the required fields. Refer features for more details.


@param id id path parameter. Unique identifier for the network device

@param GetDetailsOfASingleNetworkDeviceV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-details-of-a-single-network-device
*/
func (s *DevicesService) GetDetailsOfASingleNetworkDeviceV1(id string, GetDetailsOfASingleNetworkDeviceV1QueryParams *GetDetailsOfASingleNetworkDeviceV1QueryParams) (*ResponseDevicesGetDetailsOfASingleNetworkDeviceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetDetailsOfASingleNetworkDeviceV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetDetailsOfASingleNetworkDeviceV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDetailsOfASingleNetworkDeviceV1(id, GetDetailsOfASingleNetworkDeviceV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDetailsOfASingleNetworkDeviceV1")
	}

	result := response.Result().(*ResponseDevicesGetDetailsOfASingleNetworkDeviceV1)
	return result, response, err

}

//GetResyncIntervalForTheNetworkDeviceV1 Get resync interval for the network device - 4783-7a87-4aea-91e6
/* Fetch the reysnc interval for the given network device id.


@param id id path parameter. The id of the network device.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-resync-interval-for-the-network-device
*/
func (s *DevicesService) GetResyncIntervalForTheNetworkDeviceV1(id string) (*ResponseDevicesGetResyncIntervalForTheNetworkDeviceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/{id}/resyncIntervalSettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetResyncIntervalForTheNetworkDeviceV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetResyncIntervalForTheNetworkDeviceV1(id)
		}
		return nil, response, fmt.Errorf("error with operation GetResyncIntervalForTheNetworkDeviceV1")
	}

	result := response.Result().(*ResponseDevicesGetResyncIntervalForTheNetworkDeviceV1)
	return result, response, err

}

//WirelessRogueApContainmentStatusV1 Wireless Rogue AP Containment Status - a1ab-f9ae-4c38-9286
/* Intent API to check the wireless rogue access point containment status. The response includes all the details like containment status, contained by WLC, containment status of each BSSID etc. This API also includes the information of strongest detecting WLC for this rogue access point.


@param macAddress macAddress path parameter. MAC Address of the Wireless Rogue AP


Documentation Link: https://developer.cisco.com/docs/dna-center/#!wireless-rogue-ap-containment-status
*/
func (s *DevicesService) WirelessRogueApContainmentStatusV1(macAddress string) (*ResponseDevicesWirelessRogueApContainmentStatusV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/security/rogue/wireless-containment/status/{macAddress}"
	path = strings.Replace(path, "{macAddress}", fmt.Sprintf("%v", macAddress), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesWirelessRogueApContainmentStatusV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.WirelessRogueApContainmentStatusV1(macAddress)
		}
		return nil, response, fmt.Errorf("error with operation WirelessRogueApContainmentStatusV1")
	}

	result := response.Result().(*ResponseDevicesWirelessRogueApContainmentStatusV1)
	return result, response, err

}

//GetThreatLevelsV1 Get Threat Levels - 64ba-bad4-4aa9-b493
/* Intent API to fetch all threat levels defined.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-threat-levels
*/
func (s *DevicesService) GetThreatLevelsV1() (*ResponseDevicesGetThreatLevelsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/security/threats/level"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetThreatLevelsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetThreatLevelsV1()
		}
		return nil, response, fmt.Errorf("error with operation GetThreatLevelsV1")
	}

	result := response.Result().(*ResponseDevicesGetThreatLevelsV1)
	return result, response, err

}

//GetAllowedMacAddressV1 Get Allowed Mac Address - 18ae-3ab0-447a-872f
/* Intent API to fetch all the allowed mac addresses in the system.


@param GetAllowedMacAddressV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-allowed-mac-address
*/
func (s *DevicesService) GetAllowedMacAddressV1(GetAllowedMacAddressV1QueryParams *GetAllowedMacAddressV1QueryParams) (*ResponseDevicesGetAllowedMacAddressV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/security/threats/rogue/allowed-list"

	queryString, _ := query.Values(GetAllowedMacAddressV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetAllowedMacAddressV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAllowedMacAddressV1(GetAllowedMacAddressV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAllowedMacAddressV1")
	}

	result := response.Result().(*ResponseDevicesGetAllowedMacAddressV1)
	return result, response, err

}

//GetAllowedMacAddressCountV1 Get Allowed Mac Address Count - d4a1-e8c8-410a-b009
/* Intent API to fetch the count of allowed mac addresses in the system.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-allowed-mac-address-count
*/
func (s *DevicesService) GetAllowedMacAddressCountV1() (*ResponseDevicesGetAllowedMacAddressCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/security/threats/rogue/allowed-list/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetAllowedMacAddressCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAllowedMacAddressCountV1()
		}
		return nil, response, fmt.Errorf("error with operation GetAllowedMacAddressCountV1")
	}

	result := response.Result().(*ResponseDevicesGetAllowedMacAddressCountV1)
	return result, response, err

}

//GetThreatTypesV1 Get Threat Types - 519a-9b70-45c8-8b82
/* Intent API to fetch all threat types defined.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-threat-types
*/
func (s *DevicesService) GetThreatTypesV1() (*ResponseDevicesGetThreatTypesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/security/threats/type"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetThreatTypesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetThreatTypesV1()
		}
		return nil, response, fmt.Errorf("error with operation GetThreatTypesV1")
	}

	result := response.Result().(*ResponseDevicesGetThreatTypesV1)
	return result, response, err

}

//RetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1 Retrieves the list of AAA Services for given set of complex filters. - 55b1-ab25-40f8-ac6b
/* Retrieves the list of AAA Services and offers complex filtering and sorting capabilities. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AAAServices-1.0.0-resolved.yaml


@param RetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-list-of-a-a-a-services-for-given-set-of-complex-filters
*/
func (s *DevicesService) RetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1(requestDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1 *RequestDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1, RetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams *RetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams) (*ResponseDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1, *resty.Response, error) {
	path := "/dna/data/api/v1/aaaServices/query"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams != nil {

		if RetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1).
		SetResult(&ResponseDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1(requestDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1, RetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation RetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1")
	}

	result := response.Result().(*ResponseDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1)
	return result, response, err

}

//RetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersV1 Retrieves the total number of AAA Services for given set of complex filters. - f894-482b-4d28-852f
/* Retrieves the total number of AAA Services and offers complex filtering and sorting capabilities. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AAAServices-1.0.0-resolved.yaml


@param RetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-total-number-of-a-a-a-services-for-given-set-of-complex-filters
*/
func (s *DevicesService) RetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersV1(requestDevicesRetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersV1 *RequestDevicesRetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersV1, RetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams *RetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams) (*ResponseDevicesRetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersV1, *resty.Response, error) {
	path := "/dna/data/api/v1/aaaServices/query/count"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams != nil {

		if RetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesRetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersV1).
		SetResult(&ResponseDevicesRetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersV1(requestDevicesRetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersV1, RetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation RetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersV1")
	}

	result := response.Result().(*ResponseDevicesRetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersV1)
	return result, response, err

}

//GetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1 Get summary analytics data of AAA Services for given set of complex filters. - 0aaa-aa9b-4d39-90ec
/* Gets the summary analytics data related to AAA Services based on given filters and group by field. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AAAServices-1.0.0-resolved.yaml


@param GetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-summary-analytics-data-of-a-a-a-services-for-given-set-of-complex-filters
*/
func (s *DevicesService) GetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1(requestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1 *RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1, GetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams *GetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams) (*ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1, *resty.Response, error) {
	path := "/dna/data/api/v1/aaaServices/summaryAnalytics"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams != nil {

		if GetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1).
		SetResult(&ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1(requestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1, GetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1")
	}

	result := response.Result().(*ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1)
	return result, response, err

}

//GetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1 Get Top N analytics data of AAA Services for given set of complex filters. - 69ad-ebf6-4ffa-89d4
/* Gets the Top N analytics data related to AAA Services based on given filters and group by field. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AAAServices-1.0.0-resolved.yaml


@param GetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-top-n-analytics-data-of-a-a-a-services-for-given-set-of-complex-filters
*/
func (s *DevicesService) GetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1(requestDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1 *RequestDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1, GetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams *GetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams) (*ResponseDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1, *resty.Response, error) {
	path := "/dna/data/api/v1/aaaServices/topNAnalytics"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams != nil {

		if GetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1).
		SetResult(&ResponseDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1(requestDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1, GetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1")
	}

	result := response.Result().(*ResponseDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1)
	return result, response, err

}

//GetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1 Get trend analytics data of AAA Services for given set of complex filters. - 00a7-4ab0-4b3b-8e31
/* Gets the trend analytics data related to AAA Services based on given filters and group by field. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AAAServices-1.0.0-resolved.yaml


@param GetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-trend-analytics-data-of-a-a-a-services-for-given-set-of-complex-filters
*/
func (s *DevicesService) GetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1(requestDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1 *RequestDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1, GetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams *GetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams) (*ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1, *resty.Response, error) {
	path := "/dna/data/api/v1/aaaServices/trendAnalytics"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams != nil {

		if GetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1).
		SetResult(&ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1(requestDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1, GetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1")
	}

	result := response.Result().(*ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1)
	return result, response, err

}

//GetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1 Get trend analytics data for a given AAA Service matching the id of the Service. - f595-b9d5-413b-95fa
/* Gets the trend analytics data related to a particular AAA Service matching the id. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AAAServices-1.0.0-resolved.yaml


@param id id path parameter. Unique id of the AAA Service. It is the combination of AAA Server IP (`serverIp`) and Device UUID (`deviceId`) separated by underscore (`_`). Example: If `serverIp` is `10.76.81.33` and `deviceId` is `6bef213c-19ca-4170-8375-b694e251101c`, then the `id` would be `10.76.81.33_6bef213c-19ca-4170-8375-b694e251101c`

@param GetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIdOfTheServiceV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-trend-analytics-data-for-a-given-a-a-a-service-matching-the-id-of-the-service
*/
func (s *DevicesService) GetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1(id string, requestDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIdOfTheServiceV1 *RequestDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1, GetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIdOfTheServiceV1HeaderParams *GetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1HeaderParams) (*ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1, *resty.Response, error) {
	path := "/dna/data/api/v1/aaaServices/{id}/trendAnalytics"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIdOfTheServiceV1HeaderParams != nil {

		if GetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIdOfTheServiceV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIdOfTheServiceV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIdOfTheServiceV1).
		SetResult(&ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1(id, requestDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIdOfTheServiceV1, GetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIdOfTheServiceV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIdOfTheServiceV1")
	}

	result := response.Result().(*ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1)
	return result, response, err

}

//QueryAssuranceEventsWithFiltersV1 Query assurance events with filters - c5b7-0a69-4409-9b5b
/* Returns the list of events discovered by Catalyst Center, determined by the complex filters. Please refer to the 'API Support Documentation' section to understand which fields are supported. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceEvents-1.0.0-resolved.yaml


@param QueryAssuranceEventsWithFiltersV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!query-assurance-events-with-filters
*/
func (s *DevicesService) QueryAssuranceEventsWithFiltersV1(requestDevicesQueryAssuranceEventsWithFiltersV1 *RequestDevicesQueryAssuranceEventsWithFiltersV1, QueryAssuranceEventsWithFiltersV1HeaderParams *QueryAssuranceEventsWithFiltersV1HeaderParams) (*ResponseDevicesQueryAssuranceEventsWithFiltersV1, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceEvents/query"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if QueryAssuranceEventsWithFiltersV1HeaderParams != nil {

		if QueryAssuranceEventsWithFiltersV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", QueryAssuranceEventsWithFiltersV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesQueryAssuranceEventsWithFiltersV1).
		SetResult(&ResponseDevicesQueryAssuranceEventsWithFiltersV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.QueryAssuranceEventsWithFiltersV1(requestDevicesQueryAssuranceEventsWithFiltersV1, QueryAssuranceEventsWithFiltersV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation QueryAssuranceEventsWithFiltersV1")
	}

	result := response.Result().(*ResponseDevicesQueryAssuranceEventsWithFiltersV1)
	return result, response, err

}

//CountTheNumberOfEventsWithFiltersV1 Count the number of events with filters - d685-3aeb-4878-a8fd
/* API to fetch the count of assurance events for the given complex query. Please refer to the 'API Support Documentation' section to understand which fields are supported. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceEvents-1.0.0-resolved.yaml


@param CountTheNumberOfEventsWithFiltersV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!count-the-number-of-events-with-filters
*/
func (s *DevicesService) CountTheNumberOfEventsWithFiltersV1(requestDevicesCountTheNumberOfEventsWithFiltersV1 *RequestDevicesCountTheNumberOfEventsWithFiltersV1, CountTheNumberOfEventsWithFiltersV1HeaderParams *CountTheNumberOfEventsWithFiltersV1HeaderParams) (*ResponseDevicesCountTheNumberOfEventsWithFiltersV1, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceEvents/query/count"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if CountTheNumberOfEventsWithFiltersV1HeaderParams != nil {

		if CountTheNumberOfEventsWithFiltersV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", CountTheNumberOfEventsWithFiltersV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesCountTheNumberOfEventsWithFiltersV1).
		SetResult(&ResponseDevicesCountTheNumberOfEventsWithFiltersV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CountTheNumberOfEventsWithFiltersV1(requestDevicesCountTheNumberOfEventsWithFiltersV1, CountTheNumberOfEventsWithFiltersV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation CountTheNumberOfEventsWithFiltersV1")
	}

	result := response.Result().(*ResponseDevicesCountTheNumberOfEventsWithFiltersV1)
	return result, response, err

}

//RetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1 Retrieves the list of DHCP Services for given set of complex filters. - 9b95-faaf-467b-8b13
/* Retrieves the list of DHCP Services and offers complex filtering and sorting capabilities. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DHCPServices-1.0.0-resolved.yaml


@param RetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-list-of-d-h-c-p-services-for-given-set-of-complex-filters
*/
func (s *DevicesService) RetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1(requestDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1 *RequestDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1, RetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams *RetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams) (*ResponseDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1, *resty.Response, error) {
	path := "/dna/data/api/v1/dhcpServices/query"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams != nil {

		if RetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1).
		SetResult(&ResponseDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1(requestDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1, RetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation RetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1")
	}

	result := response.Result().(*ResponseDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1)
	return result, response, err

}

//RetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersV1 Retrieves the total number of DHCP Services for given set of complex filters. - 4b94-993c-459b-96a2
/* Retrieves the total number of DHCP Services and offers complex filtering and sorting capabilities. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DHCPServices-1.0.0-resolved.yaml


@param RetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-total-number-of-d-h-c-p-services-for-given-set-of-complex-filters
*/
func (s *DevicesService) RetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersV1(requestDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersV1 *RequestDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersV1, RetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams *RetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams) (*ResponseDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersV1, *resty.Response, error) {
	path := "/dna/data/api/v1/dhcpServices/query/count"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams != nil {

		if RetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersV1).
		SetResult(&ResponseDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersV1(requestDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersV1, RetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation RetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersV1")
	}

	result := response.Result().(*ResponseDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersV1)
	return result, response, err

}

//GetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1 Get summary analytics data of DHCP Services for given set of complex filters. - 84b8-3b7d-405a-95e3
/* Gets the summary analytics data related to DHCP Services based on given filters and group by field. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DHCPServices-1.0.0-resolved.yaml


@param GetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-summary-analytics-data-of-d-h-c-p-services-for-given-set-of-complex-filters
*/
func (s *DevicesService) GetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1(requestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1 *RequestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1, GetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams *GetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams) (*ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1, *resty.Response, error) {
	path := "/dna/data/api/v1/dhcpServices/summaryAnalytics"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams != nil {

		if GetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1).
		SetResult(&ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1(requestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1, GetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1")
	}

	result := response.Result().(*ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1)
	return result, response, err

}

//GetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1 Get Top N analytics data of DHCP Services for given set of complex filters. - c2b2-0a95-46ca-a4f1
/* Gets the Top N analytics data related to DHCP Services based on given filters and group by field. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DHCPServices-1.0.0-resolved.yaml


@param GetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-top-n-analytics-data-of-d-h-c-p-services-for-given-set-of-complex-filters
*/
func (s *DevicesService) GetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1(requestDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1 *RequestDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1, GetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams *GetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams) (*ResponseDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1, *resty.Response, error) {
	path := "/dna/data/api/v1/dhcpServices/topNAnalytics"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams != nil {

		if GetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1).
		SetResult(&ResponseDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1(requestDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1, GetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1")
	}

	result := response.Result().(*ResponseDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1)
	return result, response, err

}

//GetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1 Get trend analytics data of DHCP Services for given set of complex filters. - 73ad-bb60-447b-8f79
/* Gets the trend analytics data related to DHCP Services based on given filters and group by field. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DHCPServices-1.0.0-resolved.yaml


@param GetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-trend-analytics-data-of-d-h-c-p-services-for-given-set-of-complex-filters
*/
func (s *DevicesService) GetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1(requestDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1 *RequestDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1, GetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams *GetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams) (*ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1, *resty.Response, error) {
	path := "/dna/data/api/v1/dhcpServices/trendAnalytics"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams != nil {

		if GetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1).
		SetResult(&ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1(requestDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1, GetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1")
	}

	result := response.Result().(*ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1)
	return result, response, err

}

//GetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1 Get trend analytics data for a given DHCP Service matching the id of the Service. - c28b-d8c9-4e59-bf4a
/* Gets the trend analytics data related to a particular DHCP Service matching the id. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DHCPServices-1.0.0-resolved.yaml


@param id id path parameter. Unique id of the DHCP Service. It is the combination of DHCP Server IP (`serverIp`) and Device UUID (`deviceId`) separated by underscore (`_`). Example: If `serverIp` is `10.76.81.33` and `deviceId` is `6bef213c-19ca-4170-8375-b694e251101c`, then the `id` would be `10.76.81.33_6bef213c-19ca-4170-8375-b694e251101c`

@param GetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIdOfTheServiceV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-trend-analytics-data-for-a-given-d-h-c-p-service-matching-the-id-of-the-service
*/
func (s *DevicesService) GetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1(id string, requestDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIdOfTheServiceV1 *RequestDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1, GetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIdOfTheServiceV1HeaderParams *GetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1HeaderParams) (*ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1, *resty.Response, error) {
	path := "/dna/data/api/v1/dhcpServices/{id}/trendAnalytics"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIdOfTheServiceV1HeaderParams != nil {

		if GetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIdOfTheServiceV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIdOfTheServiceV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIdOfTheServiceV1).
		SetResult(&ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1(id, requestDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIdOfTheServiceV1, GetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIdOfTheServiceV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIdOfTheServiceV1")
	}

	result := response.Result().(*ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1)
	return result, response, err

}

//RetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1 Retrieves the list of DNS Services for given set of complex filters. - ccbf-bb30-4b29-b2c4
/* Retrieves the list of DNS Services and offers complex filtering and sorting capabilities. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DNSServices-1.0.0-resolved.yaml


@param RetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-list-of-d-n-s-services-for-given-set-of-complex-filters
*/
func (s *DevicesService) RetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1(requestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1 *RequestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1, RetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams *RetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams) (*ResponseDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1, *resty.Response, error) {
	path := "/dna/data/api/v1/dnsServices/query"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams != nil {

		if RetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1).
		SetResult(&ResponseDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1(requestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1, RetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation RetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1")
	}

	result := response.Result().(*ResponseDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1)
	return result, response, err

}

//RetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1 Retrieves the total number of DNS Services for given set of complex filters. - 8bb8-abac-469a-b511
/* Retrieves the total number of DNS Services and offers complex filtering and sorting capabilities. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DNSServices-1.0.0-resolved.yaml


@param RetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-total-number-of-d-n-s-services-for-given-set-of-complex-filters
*/
func (s *DevicesService) RetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1(requestDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1 *RequestDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1, RetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams *RetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams) (*ResponseDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1, *resty.Response, error) {
	path := "/dna/data/api/v1/dnsServices/query/count"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams != nil {

		if RetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1).
		SetResult(&ResponseDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1(requestDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1, RetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation RetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1")
	}

	result := response.Result().(*ResponseDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1)
	return result, response, err

}

//GetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1 Get summary analytics data of DNS Services for given set of complex filters. - 6993-98d1-456a-a499
/* Gets the summary analytics data related to DNS Services based on given filters and group by field. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DNSServices-1.0.0-resolved.yaml


@param GetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-summary-analytics-data-of-d-n-s-services-for-given-set-of-complex-filters
*/
func (s *DevicesService) GetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1(requestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1 *RequestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1, GetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams *GetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams) (*ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1, *resty.Response, error) {
	path := "/dna/data/api/v1/dnsServices/summaryAnalytics"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams != nil {

		if GetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1).
		SetResult(&ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1(requestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1, GetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1")
	}

	result := response.Result().(*ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1)
	return result, response, err

}

//GetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1 Get Top N analytics data of DNS Services for given set of complex filters. - 90a0-19ad-4668-88c8
/* Gets the Top N analytics data related to DNS Services based on given filters and group by field. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DNSServices-1.0.0-resolved.yaml


@param GetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-top-n-analytics-data-of-d-n-s-services-for-given-set-of-complex-filters
*/
func (s *DevicesService) GetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1(requestDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1 *RequestDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1, GetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams *GetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams) (*ResponseDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1, *resty.Response, error) {
	path := "/dna/data/api/v1/dnsServices/topNAnalytics"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams != nil {

		if GetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1).
		SetResult(&ResponseDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1(requestDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1, GetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1")
	}

	result := response.Result().(*ResponseDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1)
	return result, response, err

}

//GetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1 Get trend analytics data of DNS Services for given set of complex filters. - feb3-7b93-4b0b-a74d
/* Gets the trend analytics data related to DNS Services based on given filters and group by field. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DNSServices-1.0.0-resolved.yaml


@param GetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-trend-analytics-data-of-d-n-s-services-for-given-set-of-complex-filters
*/
func (s *DevicesService) GetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1(requestDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1 *RequestDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1, GetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams *GetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams) (*ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1, *resty.Response, error) {
	path := "/dna/data/api/v1/dnsServices/trendAnalytics"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams != nil {

		if GetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1).
		SetResult(&ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1(requestDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1, GetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1")
	}

	result := response.Result().(*ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1)
	return result, response, err

}

//GetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1 Get trend analytics data for a given DNS Service matching the id of the Service. - 6099-29a8-404a-bf56
/* Gets the trend analytics data related to a particular DNS Service matching the id. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DNSServices-1.0.0-resolved.yaml


@param id id path parameter. Unique id of the DNS Service. It is the combination of DNS Server IP (`serverIp`) and Device UUID (`deviceId`) separated by underscore (`_`). Example: If `serverIp` is `10.76.81.33` and `deviceId` is `6bef213c-19ca-4170-8375-b694e251101c`, then the `id` would be `10.76.81.33_6bef213c-19ca-4170-8375-b694e251101c`

@param GetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIdOfTheServiceV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-trend-analytics-data-for-a-given-d-n-s-service-matching-the-id-of-the-service
*/
func (s *DevicesService) GetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1(id string, requestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIdOfTheServiceV1 *RequestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1, GetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIdOfTheServiceV1HeaderParams *GetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1HeaderParams) (*ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1, *resty.Response, error) {
	path := "/dna/data/api/v1/dnsServices/{id}/trendAnalytics"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIdOfTheServiceV1HeaderParams != nil {

		if GetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIdOfTheServiceV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIdOfTheServiceV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIdOfTheServiceV1).
		SetResult(&ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1(id, requestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIdOfTheServiceV1, GetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIdOfTheServiceV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIdOfTheServiceV1")
	}

	result := response.Result().(*ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1)
	return result, response, err

}

//GetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1 Gets the list of interfaces across the Network Devices based on the provided complex filters and aggregation functions - 45b8-ba96-4daa-843c
/* Gets the list of interfaces across the Network Devices based on the provided complex filters and aggregation functions
The elements are grouped and sorted by deviceUuid first, and are then sorted by the given sort field, or by the default value: name.
The supported sorting options are: name, adminStatus, description, duplexConfig, duplexOper, interfaceIfIndex,interfaceType, macAddress,mediaType, operStatus, portChannelId, portMode, portType,speed, vlanId,pdPowerAdminMaxInWatt,pdPowerBudgetInWatt,pdPowerConsumedInWatt,pdPowerRemainingInWatt,pdMaxPowerDrawn. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-interfaces-2.0.0-resolved.yaml



Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-the-list-of-interfaces-across-the-network-devices-based-on-the-provided-complex-filters-and-aggregation-functions
*/
func (s *DevicesService) GetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1(requestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1 *RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1) (*ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1, *resty.Response, error) {
	path := "/dna/data/api/v1/interfaces/query"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1).
		SetResult(&ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1(requestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1)
		}

		return nil, response, fmt.Errorf("error with operation GetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1")
	}

	result := response.Result().(*ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1)
	return result, response, err

}

//TheTotalInterfacesCountAcrossTheNetworkDevicesV1 The Total interfaces count across the Network devices. - a0bb-1bed-4529-98b1
/* Gets the total number of interfaces across the Network devices based on the provided complex filters and aggregation functions. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-interfaces-2.0.0-resolved.yaml



Documentation Link: https://developer.cisco.com/docs/dna-center/#!the-total-interfaces-count-across-the-network-devices
*/
func (s *DevicesService) TheTotalInterfacesCountAcrossTheNetworkDevicesV1(requestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesV1 *RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesV1) (*ResponseDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesV1, *resty.Response, error) {
	path := "/dna/data/api/v1/interfaces/query/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesV1).
		SetResult(&ResponseDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.TheTotalInterfacesCountAcrossTheNetworkDevicesV1(requestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesV1)
		}

		return nil, response, fmt.Errorf("error with operation TheTotalInterfacesCountAcrossTheNetworkDevicesV1")
	}

	result := response.Result().(*ResponseDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesV1)
	return result, response, err

}

//TheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeV1 The Trend analytcis data for the interfaces in the specified time range - ed96-48ef-4a98-a0a7
/* The Trend analytcis data for the interface, identified by its instanceUuid, in the specified time range. The data is grouped based on the trend time Interval, other input parameters like attributes and aggregate attributes. The default time interval range is 3 hours when start and endTime is not provided.
The field trendIntervalInMinutes is requiered and either the attributes or the aggregateAttributes fields is also required.
For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-interfaces-2.0.0-resolved.yaml


@param id id path parameter. The interface instance Uuid


Documentation Link: https://developer.cisco.com/docs/dna-center/#!the-trend-analytcis-data-for-the-interfaces-in-the-specified-time-range
*/
func (s *DevicesService) TheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeV1(id string, requestDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeV1 *RequestDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeV1) (*ResponseDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeV1, *resty.Response, error) {
	path := "/dna/data/api/v1/interfaces/{id}/trendAnalytics"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeV1).
		SetResult(&ResponseDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.TheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeV1(id, requestDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeV1)
		}

		return nil, response, fmt.Errorf("error with operation TheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeV1")
	}

	result := response.Result().(*ResponseDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeV1)
	return result, response, err

}

//GetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1 Gets the list of Network Devices based on the provided complex filters and aggregation functions. - e794-1a90-428b-b583
/* Gets the list of Network Devices based on the provided complex filters and aggregation functions. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceNetworkDevices-2.0.1-resolved.yaml



Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-the-list-of-network-devices-based-on-the-provided-complex-filters-and-aggregation-functions
*/
func (s *DevicesService) GetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1(requestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1 *RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1) (*ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1, *resty.Response, error) {
	path := "/dna/data/api/v1/networkDevices/query"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1).
		SetResult(&ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1(requestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1)
		}

		return nil, response, fmt.Errorf("error with operation GetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1")
	}

	result := response.Result().(*ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1)
	return result, response, err

}

//GetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1 Gets the total number Network Devices based on the provided complex filters and aggregation functions. - 278f-1a5c-40ab-b65a
/* Gets the total number Network Devices based on the provided complex filters and aggregation functions. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceNetworkDevices-2.0.1-resolved.yaml



Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-the-total-number-network-devices-based-on-the-provided-complex-filters-and-aggregation-functions
*/
func (s *DevicesService) GetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1(requestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1 *RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1) (*ResponseDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1, *resty.Response, error) {
	path := "/dna/data/api/v1/networkDevices/query/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1).
		SetResult(&ResponseDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1(requestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1)
		}

		return nil, response, fmt.Errorf("error with operation GetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1")
	}

	result := response.Result().(*ResponseDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1)
	return result, response, err

}

//GetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1 Gets the summary analytics data related to network devices. - 15be-c9ed-4cba-8f91
/* Gets the summary analytics data related to network devices based on the provided input data. This endpoint helps to obtain the consolidated insights into the performance and status of the monitored network devices. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceNetworkDevices-2.0.1-resolved.yaml



Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-the-summary-analytics-data-related-to-network-devices
*/
func (s *DevicesService) GetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1(requestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1 *RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1) (*ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1, *resty.Response, error) {
	path := "/dna/data/api/v1/networkDevices/summaryAnalytics"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1).
		SetResult(&ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1(requestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1)
		}

		return nil, response, fmt.Errorf("error with operation GetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1")
	}

	result := response.Result().(*ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1)
	return result, response, err

}

//GetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1 Gets the Top-N analytics data related to network devices. - 5b94-cae7-4acb-a8fe
/* Gets the Top N analytics data related to network devices based on the provided input data. This endpoint is valuable to obtain the top-performing or most impacted network devices. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceNetworkDevices-2.0.1-resolved.yaml
The required properties for this API are topN and groupBy



Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-the-top-n-analytics-data-related-to-network-devices
*/
func (s *DevicesService) GetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1(requestDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1 *RequestDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1) (*ResponseDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1, *resty.Response, error) {
	path := "/dna/data/api/v1/networkDevices/topNAnalytics"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1).
		SetResult(&ResponseDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1(requestDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1)
		}

		return nil, response, fmt.Errorf("error with operation GetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1")
	}

	result := response.Result().(*ResponseDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1)
	return result, response, err

}

//GetsTheTrendAnalyticsDataV1 Gets the Trend analytics data. - 0c93-595e-451b-910e
/* Gets the Trend analytics Network device data for the given time range. The data will be grouped based on the given trend time Interval. The required property for this API is `trendInterval`. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceNetworkDevices-2.0.1-resolved.yaml



Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-the-trend-analytics-data
*/
func (s *DevicesService) GetsTheTrendAnalyticsDataV1(requestDevicesGetsTheTrendAnalyticsDataV1 *RequestDevicesGetsTheTrendAnalyticsDataV1) (*ResponseDevicesGetsTheTrendAnalyticsDataV1, *resty.Response, error) {
	path := "/dna/data/api/v1/networkDevices/trendAnalytics"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesGetsTheTrendAnalyticsDataV1).
		SetResult(&ResponseDevicesGetsTheTrendAnalyticsDataV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsTheTrendAnalyticsDataV1(requestDevicesGetsTheTrendAnalyticsDataV1)
		}

		return nil, response, fmt.Errorf("error with operation GetsTheTrendAnalyticsDataV1")
	}

	result := response.Result().(*ResponseDevicesGetsTheTrendAnalyticsDataV1)
	return result, response, err

}

//TheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1 The Trend analytics data for the network Device in the specified time range - 00ba-7a81-431b-93e7
/* The Trend analytics data for the network Device in the specified time range. The data is grouped based on the trend time Interval, other input parameters like attribute and aggregate attributes. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceNetworkDevices-2.0.1-resolved.yaml


@param id id path parameter. The device Uuid


Documentation Link: https://developer.cisco.com/docs/dna-center/#!the-trend-analytics-data-for-the-network-device-in-the-specified-time-range
*/
func (s *DevicesService) TheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1(id string, requestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1 *RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1) (*ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1, *resty.Response, error) {
	path := "/dna/data/api/v1/networkDevices/{id}/trendAnalytics"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1).
		SetResult(&ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.TheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1(id, requestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1)
		}

		return nil, response, fmt.Errorf("error with operation TheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1")
	}

	result := response.Result().(*ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1)
	return result, response, err

}

//CreatePlannedAccessPointForFloorV1 Create Planned Access Point for Floor - 7eaa-8b15-454a-8c1d
/* Allows creation of a new planned access point on an existing floor map including its planned radio and antenna details.  Use the Get variant of this API to fetch any existing planned access points for the floor.  The payload to create a planned access point is in the same format, albeit a single object instead of a list, of that API.


@param floorID floorId path parameter. The instance UUID of the floor hierarchy element


Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-planned-access-point-for-floor
*/
func (s *DevicesService) CreatePlannedAccessPointForFloorV1(floorID string, requestDevicesCreatePlannedAccessPointForFloorV1 *RequestDevicesCreatePlannedAccessPointForFloorV1) (*ResponseDevicesCreatePlannedAccessPointForFloorV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/floors/{floorId}/planned-access-points"
	path = strings.Replace(path, "{floorId}", fmt.Sprintf("%v", floorID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesCreatePlannedAccessPointForFloorV1).
		SetResult(&ResponseDevicesCreatePlannedAccessPointForFloorV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreatePlannedAccessPointForFloorV1(floorID, requestDevicesCreatePlannedAccessPointForFloorV1)
		}

		return nil, response, fmt.Errorf("error with operation CreatePlannedAccessPointForFloorV1")
	}

	result := response.Result().(*ResponseDevicesCreatePlannedAccessPointForFloorV1)
	return result, response, err

}

//UpdateHealthScoreDefinitionsV1 Update health score definitions. - 1aab-193d-40bb-9d2f
/* Update health thresholds, include status of overall health status for each metric.
And also to synchronize with global profile issue thresholds of the definition for given metric. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-issueAndHealthDefinitions-1.0.0-resolved.yaml


@param UpdateHealthScoreDefinitionsV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!update-health-score-definitions
*/
func (s *DevicesService) UpdateHealthScoreDefinitionsV1(requestDevicesUpdateHealthScoreDefinitionsV1 *RequestDevicesUpdateHealthScoreDefinitionsV1, UpdateHealthScoreDefinitionsV1HeaderParams *UpdateHealthScoreDefinitionsV1HeaderParams) (*ResponseDevicesUpdateHealthScoreDefinitionsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/healthScoreDefinitions/bulkUpdate"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if UpdateHealthScoreDefinitionsV1HeaderParams != nil {

		if UpdateHealthScoreDefinitionsV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", UpdateHealthScoreDefinitionsV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesUpdateHealthScoreDefinitionsV1).
		SetResult(&ResponseDevicesUpdateHealthScoreDefinitionsV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateHealthScoreDefinitionsV1(requestDevicesUpdateHealthScoreDefinitionsV1, UpdateHealthScoreDefinitionsV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation UpdateHealthScoreDefinitionsV1")
	}

	result := response.Result().(*ResponseDevicesUpdateHealthScoreDefinitionsV1)
	return result, response, err

}

//ClearMacAddressTableV1 Clear Mac-Address table - 24be-a97f-43f9-bc65
/* Clear mac-address on an individual port. In request body, operation needs to be specified as 'ClearMacAddress'. In the future more possible operations will be added to this API


@param interfaceUUID interfaceUuid path parameter. Interface Id

@param ClearMacAddressTableV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!clear-mac-address-table
*/
func (s *DevicesService) ClearMacAddressTableV1(interfaceUUID string, requestDevicesClearMacAddressTableV1 *RequestDevicesClearMacAddressTableV1, ClearMacAddressTableV1QueryParams *ClearMacAddressTableV1QueryParams) (*ResponseDevicesClearMacAddressTableV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/{interfaceUuid}/operation"
	path = strings.Replace(path, "{interfaceUuid}", fmt.Sprintf("%v", interfaceUUID), -1)

	queryString, _ := query.Values(ClearMacAddressTableV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetBody(requestDevicesClearMacAddressTableV1).
		SetResult(&ResponseDevicesClearMacAddressTableV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ClearMacAddressTableV1(interfaceUUID, requestDevicesClearMacAddressTableV1, ClearMacAddressTableV1QueryParams)
		}

		return nil, response, fmt.Errorf("error with operation ClearMacAddressTableV1")
	}

	result := response.Result().(*ResponseDevicesClearMacAddressTableV1)
	return result, response, err

}

//AddDeviceKnowYourNetworkV1 Add Device - 4bb2-2af0-46fa-8f08
/* Adds the device with given credential



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-device-know-your-network
*/
func (s *DevicesService) AddDeviceKnowYourNetworkV1(requestDevicesAddDeviceKnowYourNetworkV1 *RequestDevicesAddDeviceKnowYourNetworkV1) (*ResponseDevicesAddDeviceKnowYourNetworkV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesAddDeviceKnowYourNetworkV1).
		SetResult(&ResponseDevicesAddDeviceKnowYourNetworkV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddDeviceKnowYourNetworkV1(requestDevicesAddDeviceKnowYourNetworkV1)
		}

		return nil, response, fmt.Errorf("error with operation AddDeviceKnowYourNetworkV1")
	}

	result := response.Result().(*ResponseDevicesAddDeviceKnowYourNetworkV1)
	return result, response, err

}

//ExportDeviceListV1 Export Device list - cd98-780f-4888-a66d
/* Exports the selected network device to a file



Documentation Link: https://developer.cisco.com/docs/dna-center/#!export-device-list
*/
func (s *DevicesService) ExportDeviceListV1(requestDevicesExportDeviceListV1 *RequestDevicesExportDeviceListV1) (*ResponseDevicesExportDeviceListV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/file"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesExportDeviceListV1).
		SetResult(&ResponseDevicesExportDeviceListV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ExportDeviceListV1(requestDevicesExportDeviceListV1)
		}

		return nil, response, fmt.Errorf("error with operation ExportDeviceListV1")
	}

	result := response.Result().(*ResponseDevicesExportDeviceListV1)
	return result, response, err

}

//CreateUserDefinedFieldV1 Create User-Defined-Field - 0a9c-18e7-4caa-8b07
/* Creates a new global User Defined Field, which can be assigned to devices



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-user-defined-field
*/
func (s *DevicesService) CreateUserDefinedFieldV1(requestDevicesCreateUserDefinedFieldV1 *RequestDevicesCreateUserDefinedFieldV1) (*ResponseDevicesCreateUserDefinedFieldV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/user-defined-field"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesCreateUserDefinedFieldV1).
		SetResult(&ResponseDevicesCreateUserDefinedFieldV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateUserDefinedFieldV1(requestDevicesCreateUserDefinedFieldV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateUserDefinedFieldV1")
	}

	result := response.Result().(*ResponseDevicesCreateUserDefinedFieldV1)
	return result, response, err

}

//CreateMaintenanceScheduleForNetworkDevicesV1 Create maintenance schedule for network devices - 0a8a-6859-4dca-b635
/* API to create maintenance schedule for network devices. The state of network device can be queried using API `GET /dna/intent/api/v1/networkDevices`. The `managementState` attribute of the network device will be updated to `UNDER_MAINTENANCE` when the maintenance window starts.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-maintenance-schedule-for-network-devices
*/
func (s *DevicesService) CreateMaintenanceScheduleForNetworkDevicesV1(requestDevicesCreateMaintenanceScheduleForNetworkDevicesV1 *RequestDevicesCreateMaintenanceScheduleForNetworkDevicesV1) (*ResponseDevicesCreateMaintenanceScheduleForNetworkDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDeviceMaintenanceSchedules"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesCreateMaintenanceScheduleForNetworkDevicesV1).
		SetResult(&ResponseDevicesCreateMaintenanceScheduleForNetworkDevicesV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateMaintenanceScheduleForNetworkDevicesV1(requestDevicesCreateMaintenanceScheduleForNetworkDevicesV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateMaintenanceScheduleForNetworkDevicesV1")
	}

	result := response.Result().(*ResponseDevicesCreateMaintenanceScheduleForNetworkDevicesV1)
	return result, response, err

}

//DeleteNetworkDeviceWithConfigurationCleanupV1 Delete network device with configuration cleanup - 5080-5a46-4bbb-8ec0
/* This API endpoint facilitates the deletion of a network device after performing configuration cleanup on the device.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-network-device-with-configuration-cleanup
*/
func (s *DevicesService) DeleteNetworkDeviceWithConfigurationCleanupV1(requestDevicesDeleteNetworkDeviceWithConfigurationCleanupV1 *RequestDevicesDeleteNetworkDeviceWithConfigurationCleanupV1) (*ResponseDevicesDeleteNetworkDeviceWithConfigurationCleanupV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/deleteWithCleanup"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesDeleteNetworkDeviceWithConfigurationCleanupV1).
		SetResult(&ResponseDevicesDeleteNetworkDeviceWithConfigurationCleanupV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteNetworkDeviceWithConfigurationCleanupV1(requestDevicesDeleteNetworkDeviceWithConfigurationCleanupV1)
		}

		return nil, response, fmt.Errorf("error with operation DeleteNetworkDeviceWithConfigurationCleanupV1")
	}

	result := response.Result().(*ResponseDevicesDeleteNetworkDeviceWithConfigurationCleanupV1)
	return result, response, err

}

//DeleteANetworkDeviceWithoutConfigurationCleanupV1 Delete a network device without configuration cleanup - 9a88-7984-4ba8-ab1c
/* This API endpoint facilitates the deletion of a network device without performing configuration cleanup on the device.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-a-network-device-without-configuration-cleanup
*/
func (s *DevicesService) DeleteANetworkDeviceWithoutConfigurationCleanupV1(requestDevicesDeleteANetworkDeviceWithoutConfigurationCleanupV1 *RequestDevicesDeleteANetworkDeviceWithoutConfigurationCleanupV1) (*ResponseDevicesDeleteANetworkDeviceWithoutConfigurationCleanupV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/deleteWithoutCleanup"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesDeleteANetworkDeviceWithoutConfigurationCleanupV1).
		SetResult(&ResponseDevicesDeleteANetworkDeviceWithoutConfigurationCleanupV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteANetworkDeviceWithoutConfigurationCleanupV1(requestDevicesDeleteANetworkDeviceWithoutConfigurationCleanupV1)
		}

		return nil, response, fmt.Errorf("error with operation DeleteANetworkDeviceWithoutConfigurationCleanupV1")
	}

	result := response.Result().(*ResponseDevicesDeleteANetworkDeviceWithoutConfigurationCleanupV1)
	return result, response, err

}

//QueryNetworkDevicesWithFiltersV1 Query network devices with filters - 0caa-2abb-490a-8e69
/* Returns the list of network devices, determined by the filters. It is possible to filter the network devices based on various parameters, such as device type, device role, software version, etc.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!query-network-devices-with-filters
*/
func (s *DevicesService) QueryNetworkDevicesWithFiltersV1(requestDevicesQueryNetworkDevicesWithFiltersV1 *RequestDevicesQueryNetworkDevicesWithFiltersV1) (*ResponseDevicesQueryNetworkDevicesWithFiltersV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/query"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesQueryNetworkDevicesWithFiltersV1).
		SetResult(&ResponseDevicesQueryNetworkDevicesWithFiltersV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.QueryNetworkDevicesWithFiltersV1(requestDevicesQueryNetworkDevicesWithFiltersV1)
		}

		return nil, response, fmt.Errorf("error with operation QueryNetworkDevicesWithFiltersV1")
	}

	result := response.Result().(*ResponseDevicesQueryNetworkDevicesWithFiltersV1)
	return result, response, err

}

//CountTheNumberOfNetworkDevicesWithFiltersV1 Count the number of network devices with filters - 83a1-08b8-471b-9820
/* API to fetch the count of network devices for the given filter query.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!count-the-number-of-network-devices-with-filters
*/
func (s *DevicesService) CountTheNumberOfNetworkDevicesWithFiltersV1(requestDevicesCountTheNumberOfNetworkDevicesWithFiltersV1 *RequestDevicesCountTheNumberOfNetworkDevicesWithFiltersV1) (*ResponseDevicesCountTheNumberOfNetworkDevicesWithFiltersV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/query/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesCountTheNumberOfNetworkDevicesWithFiltersV1).
		SetResult(&ResponseDevicesCountTheNumberOfNetworkDevicesWithFiltersV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CountTheNumberOfNetworkDevicesWithFiltersV1(requestDevicesCountTheNumberOfNetworkDevicesWithFiltersV1)
		}

		return nil, response, fmt.Errorf("error with operation CountTheNumberOfNetworkDevicesWithFiltersV1")
	}

	result := response.Result().(*ResponseDevicesCountTheNumberOfNetworkDevicesWithFiltersV1)
	return result, response, err

}

//OverrideResyncIntervalV1 Override resync interval - 42ac-59bd-41db-a4fe
/* Overrides the global resync interval on all network devices. This essentially removes device specific intervals if set.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!override-resync-interval
*/
func (s *DevicesService) OverrideResyncIntervalV1() (*ResponseDevicesOverrideResyncIntervalV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/resyncIntervalSettings/override"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesOverrideResyncIntervalV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.OverrideResyncIntervalV1()
		}

		return nil, response, fmt.Errorf("error with operation OverrideResyncIntervalV1")
	}

	result := response.Result().(*ResponseDevicesOverrideResyncIntervalV1)
	return result, response, err

}

//RogueAdditionalDetailsV1 Rogue Additional Details - 659c-e9bd-403a-8de6
/* This API provides additional information of the rogue threats with details at BSSID level. The additional information includes Switch Port details in case of Rogue on Wire, first time when the rogue is seen in the network etc.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!rogue-additional-details
*/
func (s *DevicesService) RogueAdditionalDetailsV1(requestDevicesRogueAdditionalDetailsV1 *RequestDevicesRogueAdditionalDetailsV1) (*ResponseDevicesRogueAdditionalDetailsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/security/rogue/additional/details"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesRogueAdditionalDetailsV1).
		SetResult(&ResponseDevicesRogueAdditionalDetailsV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RogueAdditionalDetailsV1(requestDevicesRogueAdditionalDetailsV1)
		}

		return nil, response, fmt.Errorf("error with operation RogueAdditionalDetailsV1")
	}

	result := response.Result().(*ResponseDevicesRogueAdditionalDetailsV1)
	return result, response, err

}

//RogueAdditionalDetailCountV1 Rogue Additional Detail Count - 4ca7-59be-4b99-9041
/* This API returns the count for the Rogue Additional Details.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!rogue-additional-detail-count
*/
func (s *DevicesService) RogueAdditionalDetailCountV1(requestDevicesRogueAdditionalDetailCountV1 *RequestDevicesRogueAdditionalDetailCountV1) (*ResponseDevicesRogueAdditionalDetailCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/security/rogue/additional/details/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesRogueAdditionalDetailCountV1).
		SetResult(&ResponseDevicesRogueAdditionalDetailCountV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RogueAdditionalDetailCountV1(requestDevicesRogueAdditionalDetailCountV1)
		}

		return nil, response, fmt.Errorf("error with operation RogueAdditionalDetailCountV1")
	}

	result := response.Result().(*ResponseDevicesRogueAdditionalDetailCountV1)
	return result, response, err

}

//StartWirelessRogueApContainmentV1 Start Wireless Rogue AP Containment - 6998-5b93-4218-aea5
/* Intent API to start the wireless rogue access point containment. This API will initiate the containment operation on the strongest detecting WLC for the given Rogue AP. This is a resource intensive operation which has legal implications since the rogue access point on whom it is triggered, might be a valid neighbor access point.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!start-wireless-rogue-ap-containment
*/
func (s *DevicesService) StartWirelessRogueApContainmentV1(requestDevicesStartWirelessRogueAPContainmentV1 *RequestDevicesStartWirelessRogueApContainmentV1) (*ResponseDevicesStartWirelessRogueApContainmentV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/security/rogue/wireless-containment/start"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesStartWirelessRogueAPContainmentV1).
		SetResult(&ResponseDevicesStartWirelessRogueApContainmentV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.StartWirelessRogueApContainmentV1(requestDevicesStartWirelessRogueAPContainmentV1)
		}

		return nil, response, fmt.Errorf("error with operation StartWirelessRogueApContainmentV1")
	}

	result := response.Result().(*ResponseDevicesStartWirelessRogueApContainmentV1)
	return result, response, err

}

//StopWirelessRogueApContainmentV1 Stop Wireless Rogue AP Containment - b692-6b1c-4d0a-b3fb
/* Intent API to stop the wireless rogue access point containment. This API will stop the containment through single WLC. The response includes the details like WLC and BSSID on which the stop containment has been initiated.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!stop-wireless-rogue-ap-containment
*/
func (s *DevicesService) StopWirelessRogueApContainmentV1(requestDevicesStopWirelessRogueAPContainmentV1 *RequestDevicesStopWirelessRogueApContainmentV1) (*ResponseDevicesStopWirelessRogueApContainmentV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/security/rogue/wireless-containment/stop"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesStopWirelessRogueAPContainmentV1).
		SetResult(&ResponseDevicesStopWirelessRogueApContainmentV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.StopWirelessRogueApContainmentV1(requestDevicesStopWirelessRogueAPContainmentV1)
		}

		return nil, response, fmt.Errorf("error with operation StopWirelessRogueApContainmentV1")
	}

	result := response.Result().(*ResponseDevicesStopWirelessRogueApContainmentV1)
	return result, response, err

}

//ThreatDetailsV1 Threat Details - f6bf-c880-435a-ae2a
/* The details for the Rogue and aWIPS threats



Documentation Link: https://developer.cisco.com/docs/dna-center/#!threat-details
*/
func (s *DevicesService) ThreatDetailsV1(requestDevicesThreatDetailsV1 *RequestDevicesThreatDetailsV1) (*ResponseDevicesThreatDetailsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/security/threats/details"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesThreatDetailsV1).
		SetResult(&ResponseDevicesThreatDetailsV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ThreatDetailsV1(requestDevicesThreatDetailsV1)
		}

		return nil, response, fmt.Errorf("error with operation ThreatDetailsV1")
	}

	result := response.Result().(*ResponseDevicesThreatDetailsV1)
	return result, response, err

}

//ThreatDetailCountV1 Threat Detail Count - eb8c-2a83-45aa-871f
/* The details count for the Rogue and aWIPS threats



Documentation Link: https://developer.cisco.com/docs/dna-center/#!threat-detail-count
*/
func (s *DevicesService) ThreatDetailCountV1(requestDevicesThreatDetailCountV1 *RequestDevicesThreatDetailCountV1) (*ResponseDevicesThreatDetailCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/security/threats/details/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesThreatDetailCountV1).
		SetResult(&ResponseDevicesThreatDetailCountV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ThreatDetailCountV1(requestDevicesThreatDetailCountV1)
		}

		return nil, response, fmt.Errorf("error with operation ThreatDetailCountV1")
	}

	result := response.Result().(*ResponseDevicesThreatDetailCountV1)
	return result, response, err

}

//AddAllowedMacAddressV1 Add Allowed Mac Address - b6a0-887d-4fe9-9d5f
/* Intent API to add the threat mac address to allowed list.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-allowed-mac-address
*/
func (s *DevicesService) AddAllowedMacAddressV1(requestDevicesAddAllowedMacAddressV1 *RequestDevicesAddAllowedMacAddressV1) (*ResponseDevicesAddAllowedMacAddressV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/security/threats/rogue/allowed-list"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesAddAllowedMacAddressV1).
		SetResult(&ResponseDevicesAddAllowedMacAddressV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddAllowedMacAddressV1(requestDevicesAddAllowedMacAddressV1)
		}

		return nil, response, fmt.Errorf("error with operation AddAllowedMacAddressV1")
	}

	result := response.Result().(*ResponseDevicesAddAllowedMacAddressV1)
	return result, response, err

}

//ThreatSummaryV1 Threat Summary - 3b98-98f0-4cfb-b74b
/* The Threat Summary for the Rogues and aWIPS



Documentation Link: https://developer.cisco.com/docs/dna-center/#!threat-summary
*/
func (s *DevicesService) ThreatSummaryV1(requestDevicesThreatSummaryV1 *RequestDevicesThreatSummaryV1) (*ResponseDevicesThreatSummaryV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/security/threats/summary"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesThreatSummaryV1).
		SetResult(&ResponseDevicesThreatSummaryV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ThreatSummaryV1(requestDevicesThreatSummaryV1)
		}

		return nil, response, fmt.Errorf("error with operation ThreatSummaryV1")
	}

	result := response.Result().(*ResponseDevicesThreatSummaryV1)
	return result, response, err

}

//GetDeviceInterfaceStatsInfoV2 Get Device Interface Stats Info - 76bb-5957-49ab-8a3b
/* This API returns the Interface Stats for the given Device Id. Please refer to the Feature tab for the Request Body usage and the API filtering support.


@param deviceID deviceId path parameter. Network Device Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-interface-stats-info-v2
*/
func (s *DevicesService) GetDeviceInterfaceStatsInfoV2(deviceID string, requestDevicesGetDeviceInterfaceStatsInfoV2 *RequestDevicesGetDeviceInterfaceStatsInfoV2) (*ResponseDevicesGetDeviceInterfaceStatsInfoV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/networkDevices/{deviceId}/interfaces/query"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesGetDeviceInterfaceStatsInfoV2).
		SetResult(&ResponseDevicesGetDeviceInterfaceStatsInfoV2{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceInterfaceStatsInfoV2(deviceID, requestDevicesGetDeviceInterfaceStatsInfoV2)
		}

		return nil, response, fmt.Errorf("error with operation GetDeviceInterfaceStatsInfoV2")
	}

	result := response.Result().(*ResponseDevicesGetDeviceInterfaceStatsInfoV2)
	return result, response, err

}

//UpdatePlannedAccessPointForFloorV1 Update Planned Access Point for Floor - 399f-596d-4f69-a080
/* Allows updating a planned access point on an existing floor map including its planned radio and antenna details.  Use the Get variant of this API to fetch the existing planned access points for the floor.  The payload to update a planned access point is in the same format, albeit a single object instead of a list, of that API.


@param floorID floorId path parameter. The instance UUID of the floor hierarchy element

*/
func (s *DevicesService) UpdatePlannedAccessPointForFloorV1(floorID string, requestDevicesUpdatePlannedAccessPointForFloorV1 *RequestDevicesUpdatePlannedAccessPointForFloorV1) (*ResponseDevicesUpdatePlannedAccessPointForFloorV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/floors/{floorId}/planned-access-points"
	path = strings.Replace(path, "{floorId}", fmt.Sprintf("%v", floorID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdatePlannedAccessPointForFloorV1).
		SetResult(&ResponseDevicesUpdatePlannedAccessPointForFloorV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdatePlannedAccessPointForFloorV1(floorID, requestDevicesUpdatePlannedAccessPointForFloorV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdatePlannedAccessPointForFloorV1")
	}

	result := response.Result().(*ResponseDevicesUpdatePlannedAccessPointForFloorV1)
	return result, response, err

}

//UpdateHealthScoreDefinitionForTheGivenIDV1 Update health score definition for the given id. - f295-190f-4f08-bbe0
/* Update health threshold, include status of overall health status.
And also to synchronize with global profile issue thresholds of the definition for given id. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-issueAndHealthDefinitions-1.0.0-resolved.yaml


@param id id path parameter. Health score definition id.

*/
func (s *DevicesService) UpdateHealthScoreDefinitionForTheGivenIDV1(id string, requestDevicesUpdateHealthScoreDefinitionForTheGivenIdV1 *RequestDevicesUpdateHealthScoreDefinitionForTheGivenIDV1) (*ResponseDevicesUpdateHealthScoreDefinitionForTheGivenIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/healthScoreDefinitions/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdateHealthScoreDefinitionForTheGivenIdV1).
		SetResult(&ResponseDevicesUpdateHealthScoreDefinitionForTheGivenIDV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateHealthScoreDefinitionForTheGivenIDV1(id, requestDevicesUpdateHealthScoreDefinitionForTheGivenIdV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateHealthScoreDefinitionForTheGivenIdV1")
	}

	result := response.Result().(*ResponseDevicesUpdateHealthScoreDefinitionForTheGivenIDV1)
	return result, response, err

}

//UpdateInterfaceDetailsV1 Update Interface details - 868b-5a60-4be8-a2d7
/* Add/Update Interface description, VLAN membership, Voice VLAN and change Interface admin status ('UP'/'DOWN') from Request body.


@param interfaceUUID interfaceUuid path parameter. Interface ID

@param UpdateInterfaceDetailsV1QueryParams Filtering parameter
*/
func (s *DevicesService) UpdateInterfaceDetailsV1(interfaceUUID string, requestDevicesUpdateInterfaceDetailsV1 *RequestDevicesUpdateInterfaceDetailsV1, UpdateInterfaceDetailsV1QueryParams *UpdateInterfaceDetailsV1QueryParams) (*ResponseDevicesUpdateInterfaceDetailsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/{interfaceUuid}"
	path = strings.Replace(path, "{interfaceUuid}", fmt.Sprintf("%v", interfaceUUID), -1)

	queryString, _ := query.Values(UpdateInterfaceDetailsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetBody(requestDevicesUpdateInterfaceDetailsV1).
		SetResult(&ResponseDevicesUpdateInterfaceDetailsV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateInterfaceDetailsV1(interfaceUUID, requestDevicesUpdateInterfaceDetailsV1, UpdateInterfaceDetailsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation UpdateInterfaceDetailsV1")
	}

	result := response.Result().(*ResponseDevicesUpdateInterfaceDetailsV1)
	return result, response, err

}

//UpdateDeviceDetailsV1 Update Device Details - aeb9-eb67-460b-92df
/* Update the credentials, management IP address of a given device (or a set of devices) in Catalyst Center and trigger an inventory sync.


 */
func (s *DevicesService) UpdateDeviceDetailsV1(requestDevicesUpdateDeviceDetailsV1 *RequestDevicesUpdateDeviceDetailsV1) (*ResponseDevicesUpdateDeviceDetailsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdateDeviceDetailsV1).
		SetResult(&ResponseDevicesUpdateDeviceDetailsV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateDeviceDetailsV1(requestDevicesUpdateDeviceDetailsV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateDeviceDetailsV1")
	}

	result := response.Result().(*ResponseDevicesUpdateDeviceDetailsV1)
	return result, response, err

}

//UpdateDeviceRoleV1 Update Device role - b985-5ad5-4ae9-8156
/* Updates the role of the device as access, core, distribution, border router


 */
func (s *DevicesService) UpdateDeviceRoleV1(requestDevicesUpdateDeviceRoleV1 *RequestDevicesUpdateDeviceRoleV1) (*ResponseDevicesUpdateDeviceRoleV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/brief"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdateDeviceRoleV1).
		SetResult(&ResponseDevicesUpdateDeviceRoleV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateDeviceRoleV1(requestDevicesUpdateDeviceRoleV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateDeviceRoleV1")
	}

	result := response.Result().(*ResponseDevicesUpdateDeviceRoleV1)
	return result, response, err

}

//SyncDevicesV1 Sync Devices - 3b9e-f967-4429-be4c
/* Synchronizes the devices. If forceSync param is false (default) then the sync would run in normal priority thread. If forceSync param is true then the sync would run in high priority thread if available, else the sync will fail. Result can be seen in the child task of each device


@param SyncDevicesV1QueryParams Filtering parameter
*/
func (s *DevicesService) SyncDevicesV1(requestDevicesSyncDevicesV1 *RequestDevicesSyncDevicesV1, SyncDevicesV1QueryParams *SyncDevicesV1QueryParams) (*ResponseDevicesSyncDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/sync"

	queryString, _ := query.Values(SyncDevicesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetBody(requestDevicesSyncDevicesV1).
		SetResult(&ResponseDevicesSyncDevicesV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SyncDevicesV1(requestDevicesSyncDevicesV1, SyncDevicesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation SyncDevicesV1")
	}

	result := response.Result().(*ResponseDevicesSyncDevicesV1)
	return result, response, err

}

//UpdateUserDefinedFieldV1 Update User-Defined-Field - aa8c-ea8f-41aa-a346
/* Updates an existing global User Defined Field, using it's id.


@param id id path parameter. UDF id

*/
func (s *DevicesService) UpdateUserDefinedFieldV1(id string, requestDevicesUpdateUserDefinedFieldV1 *RequestDevicesUpdateUserDefinedFieldV1) (*ResponseDevicesUpdateUserDefinedFieldV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/user-defined-field/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdateUserDefinedFieldV1).
		SetResult(&ResponseDevicesUpdateUserDefinedFieldV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateUserDefinedFieldV1(id, requestDevicesUpdateUserDefinedFieldV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateUserDefinedFieldV1")
	}

	result := response.Result().(*ResponseDevicesUpdateUserDefinedFieldV1)
	return result, response, err

}

//AddUserDefinedFieldToDeviceV1 Add User-Defined-Field to device - d3af-395c-4669-adaf
/* Assigns an existing Global User-Defined-Field to a device. If the UDF is already assigned to the specific device, then it updates the device UDF value accordingly. Please note that the assigning UDF 'name' must be an existing global UDF. Otherwise error shall be shown.


@param deviceID deviceId path parameter. UUID of device to which UDF has to be added

*/
func (s *DevicesService) AddUserDefinedFieldToDeviceV1(deviceID string, requestDevicesAddUserDefinedFieldToDeviceV1 *RequestDevicesAddUserDefinedFieldToDeviceV1) (*ResponseDevicesAddUserDefinedFieldToDeviceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceId}/user-defined-field"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesAddUserDefinedFieldToDeviceV1).
		SetResult(&ResponseDevicesAddUserDefinedFieldToDeviceV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddUserDefinedFieldToDeviceV1(deviceID, requestDevicesAddUserDefinedFieldToDeviceV1)
		}
		return nil, response, fmt.Errorf("error with operation AddUserDefinedFieldToDeviceV1")
	}

	result := response.Result().(*ResponseDevicesAddUserDefinedFieldToDeviceV1)
	return result, response, err

}

//UpdateDeviceManagementAddressV1 Update Device Management Address - af93-b807-4feb-a985
/* This is a simple PUT API to edit the management IP Address of the device.


@param deviceid deviceid path parameter. The UUID of the device whose management IP address is to be updated.

*/
func (s *DevicesService) UpdateDeviceManagementAddressV1(deviceid string, requestDevicesUpdateDeviceManagementAddressV1 *RequestDevicesUpdateDeviceManagementAddressV1) (*ResponseDevicesUpdateDeviceManagementAddressV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceid}/management-address"
	path = strings.Replace(path, "{deviceid}", fmt.Sprintf("%v", deviceid), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdateDeviceManagementAddressV1).
		SetResult(&ResponseDevicesUpdateDeviceManagementAddressV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateDeviceManagementAddressV1(deviceid, requestDevicesUpdateDeviceManagementAddressV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateDeviceManagementAddressV1")
	}

	result := response.Result().(*ResponseDevicesUpdateDeviceManagementAddressV1)
	return result, response, err

}

//UpdatesTheMaintenanceScheduleInformationV1 Updates the maintenance schedule information - 8cb8-5ae3-4a2b-8a08
/* API to update the maintenance schedule for the network devices. The `maintenanceSchedule` can be updated only if the `status` value is `UPCOMING` or `IN_PROGRESS`. User can exit `IN_PROGRESS` maintenance window by setting the `endTime` to -1. This will update the endTime to the current time and exit the maintenance window immediately. When exiting the maintenance window, only the endTime will be updated while other parameters remain read-only.


@param id id path parameter. Unique identifier for the maintenance schedule

*/
func (s *DevicesService) UpdatesTheMaintenanceScheduleInformationV1(id string, requestDevicesUpdatesTheMaintenanceScheduleInformationV1 *RequestDevicesUpdatesTheMaintenanceScheduleInformationV1) (*ResponseDevicesUpdatesTheMaintenanceScheduleInformationV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDeviceMaintenanceSchedules/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdatesTheMaintenanceScheduleInformationV1).
		SetResult(&ResponseDevicesUpdatesTheMaintenanceScheduleInformationV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdatesTheMaintenanceScheduleInformationV1(id, requestDevicesUpdatesTheMaintenanceScheduleInformationV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdatesTheMaintenanceScheduleInformationV1")
	}

	result := response.Result().(*ResponseDevicesUpdatesTheMaintenanceScheduleInformationV1)
	return result, response, err

}

//UpdateGlobalResyncIntervalV1 Update global resync interval - 25b5-39b4-4609-9e2a
/* Updates the resync interval (in minutes) globally for devices which do not have custom resync interval. To override this setting for all network devices refer to [/networkDevices/resyncIntervalSettings/override]


 */
func (s *DevicesService) UpdateGlobalResyncIntervalV1(requestDevicesUpdateGlobalResyncIntervalV1 *RequestDevicesUpdateGlobalResyncIntervalV1) (*ResponseDevicesUpdateGlobalResyncIntervalV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/resyncIntervalSettings"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdateGlobalResyncIntervalV1).
		SetResult(&ResponseDevicesUpdateGlobalResyncIntervalV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateGlobalResyncIntervalV1(requestDevicesUpdateGlobalResyncIntervalV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateGlobalResyncIntervalV1")
	}

	result := response.Result().(*ResponseDevicesUpdateGlobalResyncIntervalV1)
	return result, response, err

}

//UpdateResyncIntervalForTheNetworkDeviceV1 Update resync interval for the network device - 92a0-db6c-428a-92d9
/* Update the resync interval (in minutes) for the given network device id.
To disable periodic resync, set interval as `0`.
To use global settings, set interval as `null`.


@param id id path parameter. The id of the network device.

*/
func (s *DevicesService) UpdateResyncIntervalForTheNetworkDeviceV1(id string, requestDevicesUpdateResyncIntervalForTheNetworkDeviceV1 *RequestDevicesUpdateResyncIntervalForTheNetworkDeviceV1) (*ResponseDevicesUpdateResyncIntervalForTheNetworkDeviceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/{id}/resyncIntervalSettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdateResyncIntervalForTheNetworkDeviceV1).
		SetResult(&ResponseDevicesUpdateResyncIntervalForTheNetworkDeviceV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateResyncIntervalForTheNetworkDeviceV1(id, requestDevicesUpdateResyncIntervalForTheNetworkDeviceV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateResyncIntervalForTheNetworkDeviceV1")
	}

	result := response.Result().(*ResponseDevicesUpdateResyncIntervalForTheNetworkDeviceV1)
	return result, response, err

}

//DeletePlannedAccessPointForFloorV1 Delete Planned Access Point for Floor - 6dad-1aac-4b3a-9e67
/* Allow to delete a planned access point from an existing floor map including its planned radio and antenna details.  Use the Get variant of this API to fetch the existing planned access points for the floor.  The instanceUUID listed in each of the planned access point attributes acts as the path param input to this API to delete that specific instance.


@param floorID floorId path parameter. The instance UUID of the floor hierarchy element

@param plannedAccessPointUUID plannedAccessPointUuid path parameter. The instance UUID of the planned access point to delete


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-planned-access-point-for-floor
*/
func (s *DevicesService) DeletePlannedAccessPointForFloorV1(floorID string, plannedAccessPointUUID string) (*ResponseDevicesDeletePlannedAccessPointForFloorV1, *resty.Response, error) {
	//floorID string,plannedAccessPointUUID string
	path := "/dna/intent/api/v1/floors/{floorId}/planned-access-points/{plannedAccessPointUuid}"
	path = strings.Replace(path, "{floorId}", fmt.Sprintf("%v", floorID), -1)
	path = strings.Replace(path, "{plannedAccessPointUuid}", fmt.Sprintf("%v", plannedAccessPointUUID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesDeletePlannedAccessPointForFloorV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletePlannedAccessPointForFloorV1(floorID, plannedAccessPointUUID)
		}
		return nil, response, fmt.Errorf("error with operation DeletePlannedAccessPointForFloorV1")
	}

	result := response.Result().(*ResponseDevicesDeletePlannedAccessPointForFloorV1)
	return result, response, err

}

//DeleteUserDefinedFieldV1 Delete User-Defined-Field - 78a3-c8b1-4799-892e
/* Deletes an existing Global User-Defined-Field using it's id.


@param id id path parameter. UDF id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-user-defined-field
*/
func (s *DevicesService) DeleteUserDefinedFieldV1(id string) (*ResponseDevicesDeleteUserDefinedFieldV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/network-device/user-defined-field/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesDeleteUserDefinedFieldV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteUserDefinedFieldV1(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteUserDefinedFieldV1")
	}

	result := response.Result().(*ResponseDevicesDeleteUserDefinedFieldV1)
	return result, response, err

}

//RemoveUserDefinedFieldFromDeviceV1 Remove User-Defined-Field from device - 8c9f-d9e8-4cab-bf96
/* Remove a User-Defined-Field from device. Name of UDF has to be passed as the query parameter. Please note that Global UDF will not be deleted by this operation.


@param deviceID deviceId path parameter. UUID of device from which UDF has to be removed

@param RemoveUserDefinedFieldFromDeviceV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!remove-user-defined-field-from-device
*/
func (s *DevicesService) RemoveUserDefinedFieldFromDeviceV1(deviceID string, RemoveUserDefinedFieldFromDeviceV1QueryParams *RemoveUserDefinedFieldFromDeviceV1QueryParams) (*ResponseDevicesRemoveUserDefinedFieldFromDeviceV1, *resty.Response, error) {
	//deviceID string,RemoveUserDefinedFieldFromDeviceV1QueryParams *RemoveUserDefinedFieldFromDeviceV1QueryParams
	path := "/dna/intent/api/v1/network-device/{deviceId}/user-defined-field"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	queryString, _ := query.Values(RemoveUserDefinedFieldFromDeviceV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesRemoveUserDefinedFieldFromDeviceV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RemoveUserDefinedFieldFromDeviceV1(deviceID, RemoveUserDefinedFieldFromDeviceV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RemoveUserDefinedFieldFromDeviceV1")
	}

	result := response.Result().(*ResponseDevicesRemoveUserDefinedFieldFromDeviceV1)
	return result, response, err

}

//DeleteDeviceByIDV1 Delete Device by Id - 1c89-4b58-48ea-b214
/* This API allows any network device that is not currently provisioned to be removed from the inventory. Important: Devices currently provisioned cannot be deleted. To delete a provisioned device, the device must be first deprovisioned.


@param id id path parameter. Device ID

@param DeleteDeviceByIdV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-device-by-id
*/
func (s *DevicesService) DeleteDeviceByIDV1(id string, DeleteDeviceByIdV1QueryParams *DeleteDeviceByIDV1QueryParams) (*ResponseDevicesDeleteDeviceByIDV1, *resty.Response, error) {
	//id string,DeleteDeviceByIdV1QueryParams *DeleteDeviceByIDV1QueryParams
	path := "/dna/intent/api/v1/network-device/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(DeleteDeviceByIdV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesDeleteDeviceByIDV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteDeviceByIDV1(id, DeleteDeviceByIdV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteDeviceByIdV1")
	}

	result := response.Result().(*ResponseDevicesDeleteDeviceByIDV1)
	return result, response, err

}

//DeleteMaintenanceScheduleV1 Delete maintenance schedule. - f5b4-5b47-402a-a14a
/* API to delete maintenance schedule by id. Deletion is allowed if the maintenance window is in the `UPCOMING`, `COMPLETED`, or `FAILED` state. Deletion of maintenance schedule is not allowed if the maintenance window is currently `IN_PROGRESS`. To delete the maintenance schedule while it is `IN_PROGRESS`, first exit the current maintenance window using `PUT /dna/intent/api/v1/networkDeviceMaintenanceSchedules/{id}` API, and then proceed to delete the maintenance schedule.


@param id id path parameter. Unique identifier for the maintenance schedule


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-maintenance-schedule
*/
func (s *DevicesService) DeleteMaintenanceScheduleV1(id string) (*ResponseDevicesDeleteMaintenanceScheduleV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/networkDeviceMaintenanceSchedules/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesDeleteMaintenanceScheduleV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteMaintenanceScheduleV1(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteMaintenanceScheduleV1")
	}

	result := response.Result().(*ResponseDevicesDeleteMaintenanceScheduleV1)
	return result, response, err

}

//RemoveAllowedMacAddressV1 Remove Allowed Mac Address - c8ac-a91b-4c5a-9b5c
/* Intent API to remove the threat mac address from allowed list.


@param macAddress macAddress path parameter. Threat mac address which needs to be removed from the allowed list. Multiple mac addresses will be removed if provided as comma separated values (example: 00:2A:10:51:22:43,00:2A:10:51:22:44). Note: In one request, maximum 100 mac addresses can be removed.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!remove-allowed-mac-address
*/
func (s *DevicesService) RemoveAllowedMacAddressV1(macAddress string) (*ResponseDevicesRemoveAllowedMacAddressV1, *resty.Response, error) {
	//macAddress string
	path := "/dna/intent/api/v1/security/threats/rogue/allowed-list/{macAddress}"
	path = strings.Replace(path, "{macAddress}", fmt.Sprintf("%v", macAddress), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesRemoveAllowedMacAddressV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RemoveAllowedMacAddressV1(macAddress)
		}
		return nil, response, fmt.Errorf("error with operation RemoveAllowedMacAddressV1")
	}

	result := response.Result().(*ResponseDevicesRemoveAllowedMacAddressV1)
	return result, response, err

}

// Alias Function
/*
This method acts as an alias for the method `GetInterfaceByIPV1`
*/
func (s *DevicesService) GetInterfaceByIP(ipAddress string) (*ResponseDevicesGetInterfaceByIPV1, *resty.Response, error) {
	return s.GetInterfaceByIPV1(ipAddress)

}

// Alias Function
/*
This method acts as an alias for the method `GetSupervisorCardDetailV1`
*/
func (s *DevicesService) GetSupervisorCardDetail(deviceUUID string) (*ResponseDevicesGetSupervisorCardDetailV1, *resty.Response, error) {
	return s.GetSupervisorCardDetailV1(deviceUUID)
}

// Alias Function
/*
This method acts as an alias for the method `TheTotalInterfacesCountAcrossTheNetworkDevicesV1`
*/
func (s *DevicesService) TheTotalInterfacesCountAcrossTheNetworkDevices(requestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesV1 *RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesV1) (*ResponseDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesV1, *resty.Response, error) {
	return s.TheTotalInterfacesCountAcrossTheNetworkDevicesV1(requestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesV1)
}

// Alias Function
/*
This method acts as an alias for the method `ThreatDetailsV1`
*/
func (s *DevicesService) ThreatDetails(requestDevicesThreatDetailsV1 *RequestDevicesThreatDetailsV1) (*ResponseDevicesThreatDetailsV1, *resty.Response, error) {
	return s.ThreatDetailsV1(requestDevicesThreatDetailsV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersV1`
*/
func (s *DevicesService) GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParameters(GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersV1QueryParams *GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersV1QueryParams) (*ResponseDevicesGetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersV1, *resty.Response, error) {
	return s.GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersV1(GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `PoeDetailsV1`
*/
func (s *DevicesService) PoeDetails(deviceUUID string) (*ResponseDevicesPoeDetailsV1, *resty.Response, error) {
	return s.PoeDetailsV1(deviceUUID)
}

// Alias Function
/*
This method acts as an alias for the method `GetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1`
*/
func (s *DevicesService) GetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheService(id string, requestDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIdOfTheServiceV1 *RequestDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1, GetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIdOfTheServiceV1HeaderParams *GetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1HeaderParams) (*ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1, *resty.Response, error) {
	return s.GetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceV1(id, requestDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIdOfTheServiceV1, GetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIdOfTheServiceV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesTheTotalNumberOfDNSServicesForGivenParametersV1`
*/
func (s *DevicesService) RetrievesTheTotalNumberOfDNSServicesForGivenParameters(RetrievesTheTotalNumberOfDNSServicesForGivenParametersV1HeaderParams *RetrievesTheTotalNumberOfDNSServicesForGivenParametersV1HeaderParams, RetrievesTheTotalNumberOfDNSServicesForGivenParametersV1QueryParams *RetrievesTheTotalNumberOfDNSServicesForGivenParametersV1QueryParams) (*ResponseDevicesRetrievesTheTotalNumberOfDNSServicesForGivenParametersV1, *resty.Response, error) {
	return s.RetrievesTheTotalNumberOfDNSServicesForGivenParametersV1(RetrievesTheTotalNumberOfDNSServicesForGivenParametersV1HeaderParams, RetrievesTheTotalNumberOfDNSServicesForGivenParametersV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetThreatLevelsV1`
*/
func (s *DevicesService) GetThreatLevels() (*ResponseDevicesGetThreatLevelsV1, *resty.Response, error) {
	return s.GetThreatLevelsV1()
}

// Alias Function
/*
This method acts as an alias for the method `GetConnectedDeviceDetailV1`
*/
func (s *DevicesService) GetConnectedDeviceDetail(deviceUUID string, interfaceUUID string) (*ResponseDevicesGetConnectedDeviceDetailV1, *resty.Response, error) {
	return s.GetConnectedDeviceDetailV1(deviceUUID, interfaceUUID)
}

// Alias Function
/*
This method acts as an alias for the method `ExportDeviceListV1`
*/
func (s *DevicesService) ExportDeviceList(requestDevicesExportDeviceListV1 *RequestDevicesExportDeviceListV1) (*ResponseDevicesExportDeviceListV1, *resty.Response, error) {
	return s.ExportDeviceListV1(requestDevicesExportDeviceListV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1`
*/
func (s *DevicesService) GetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions(requestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1 *RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1) (*ResponseDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1, *resty.Response, error) {
	return s.GetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1(requestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1`
*/
func (s *DevicesService) RetrievesTheListOfAAAServicesForGivenSetOfComplexFilters(requestDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1 *RequestDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1, RetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams *RetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams) (*ResponseDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1, *resty.Response, error) {
	return s.RetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1(requestDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1, RetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetOrganizationListForMerakiV1`
*/
func (s *DevicesService) GetOrganizationListForMeraki(id string) (*ResponseDevicesGetOrganizationListForMerakiV1, *resty.Response, error) {
	return s.GetOrganizationListForMerakiV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `GetWirelessLanControllerDetailsByIDV1`
*/
func (s *DevicesService) GetWirelessLanControllerDetailsByID(id string) (*ResponseDevicesGetWirelessLanControllerDetailsByIDV1, *resty.Response, error) {
	return s.GetWirelessLanControllerDetailsByIDV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesTheListOfDHCPServicesForGivenParametersV1`
*/
func (s *DevicesService) RetrievesTheListOfDHCPServicesForGivenParameters(RetrievesTheListOfDHCPServicesForGivenParametersV1HeaderParams *RetrievesTheListOfDHCPServicesForGivenParametersV1HeaderParams, RetrievesTheListOfDHCPServicesForGivenParametersV1QueryParams *RetrievesTheListOfDHCPServicesForGivenParametersV1QueryParams) (*ResponseDevicesRetrievesTheListOfDHCPServicesForGivenParametersV1, *resty.Response, error) {
	return s.RetrievesTheListOfDHCPServicesForGivenParametersV1(RetrievesTheListOfDHCPServicesForGivenParametersV1HeaderParams, RetrievesTheListOfDHCPServicesForGivenParametersV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetAllInterfacesV1`
*/
func (s *DevicesService) GetAllInterfaces(GetAllInterfacesV1QueryParams *GetAllInterfacesV1QueryParams) (*ResponseDevicesGetAllInterfacesV1, *resty.Response, error) {
	return s.GetAllInterfacesV1(GetAllInterfacesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `DeleteNetworkDeviceWithConfigurationCleanupV1`
*/
func (s *DevicesService) DeleteNetworkDeviceWithConfigurationCleanup(requestDevicesDeleteNetworkDeviceWithConfigurationCleanupV1 *RequestDevicesDeleteNetworkDeviceWithConfigurationCleanupV1) (*ResponseDevicesDeleteNetworkDeviceWithConfigurationCleanupV1, *resty.Response, error) {
	return s.DeleteNetworkDeviceWithConfigurationCleanupV1(requestDevicesDeleteNetworkDeviceWithConfigurationCleanupV1)
}

// Alias Function
/*
This method acts as an alias for the method `UpdateUserDefinedFieldV1`
*/
func (s *DevicesService) UpdateUserDefinedField(id string, requestDevicesUpdateUserDefinedFieldV1 *RequestDevicesUpdateUserDefinedFieldV1) (*ResponseDevicesUpdateUserDefinedFieldV1, *resty.Response, error) {
	return s.UpdateUserDefinedFieldV1(id, requestDevicesUpdateUserDefinedFieldV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetDeviceInterfaceStatsInfoV2`
*/
func (s *DevicesService) GetDeviceInterfaceStatsInfo(deviceID string, requestDevicesGetDeviceInterfaceStatsInfoV2 *RequestDevicesGetDeviceInterfaceStatsInfoV2) (*ResponseDevicesGetDeviceInterfaceStatsInfoV2, *resty.Response, error) {
	return s.GetDeviceInterfaceStatsInfoV2(deviceID, requestDevicesGetDeviceInterfaceStatsInfoV2)
}

// Alias Function
/*
This method acts as an alias for the method `GetNetworkDeviceByPaginationRangeV1`
*/
func (s *DevicesService) GetNetworkDeviceByPaginationRange(startIndex int, recordsToReturn int) (*ResponseDevicesGetNetworkDeviceByPaginationRangeV1, *resty.Response, error) {
	return s.GetNetworkDeviceByPaginationRangeV1(startIndex, recordsToReturn)
}

// Alias Function
/*
This method acts as an alias for the method `GetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1`
*/
func (s *DevicesService) GetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters(requestDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1 *RequestDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1, GetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams *GetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams) (*ResponseDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1, *resty.Response, error) {
	return s.GetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1(requestDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1, GetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1`
*/
func (s *DevicesService) GetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters(requestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1 *RequestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1, GetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams *GetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams) (*ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1, *resty.Response, error) {
	return s.GetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1(requestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1, GetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `StopWirelessRogueApContainmentV1`
*/
func (s *DevicesService) StopWirelessRogueApContainment(requestDevicesStopWirelessRogueAPContainmentV1 *RequestDevicesStopWirelessRogueApContainmentV1) (*ResponseDevicesStopWirelessRogueApContainmentV1, *resty.Response, error) {
	return s.StopWirelessRogueApContainmentV1(requestDevicesStopWirelessRogueAPContainmentV1)
}

// Alias Function
/*
This method acts as an alias for the method `InventoryInsightDeviceLinkMismatchAPIV1`
*/
func (s *DevicesService) InventoryInsightDeviceLinkMismatchAPI(siteID string, InventoryInsightDeviceLinkMismatchAPIV1QueryParams *InventoryInsightDeviceLinkMismatchAPIV1QueryParams) (*ResponseDevicesInventoryInsightDeviceLinkMismatchAPIV1, *resty.Response, error) {
	return s.InventoryInsightDeviceLinkMismatchAPIV1(siteID, InventoryInsightDeviceLinkMismatchAPIV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetDevicesRegisteredForWsaNotificationV1`
*/
func (s *DevicesService) GetDevicesRegisteredForWsaNotification(GetDevicesRegisteredForWSANotificationV1QueryParams *GetDevicesRegisteredForWsaNotificationV1QueryParams) (*ResponseDevicesGetDevicesRegisteredForWsaNotificationV1, *resty.Response, error) {
	return s.GetDevicesRegisteredForWsaNotificationV1(GetDevicesRegisteredForWSANotificationV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `ThreatDetailCountV1`
*/
func (s *DevicesService) ThreatDetailCount(requestDevicesThreatDetailCountV1 *RequestDevicesThreatDetailCountV1) (*ResponseDevicesThreatDetailCountV1, *resty.Response, error) {
	return s.ThreatDetailCountV1(requestDevicesThreatDetailCountV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetDeviceEnrichmentDetailsV1`
*/
func (s *DevicesService) GetDeviceEnrichmentDetails(GetDeviceEnrichmentDetailsV1HeaderParams *GetDeviceEnrichmentDetailsV1HeaderParams) (*ResponseDevicesGetDeviceEnrichmentDetailsV1, *resty.Response, error) {
	return s.GetDeviceEnrichmentDetailsV1(GetDeviceEnrichmentDetailsV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetDeviceByIDV1`
*/
func (s *DevicesService) GetDeviceByID(id string) (*ResponseDevicesGetDeviceByIDV1, *resty.Response, error) {
	return s.GetDeviceByIDV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `GetDeviceInterfaceVLANsV1`
*/
func (s *DevicesService) GetDeviceInterfaceVLANs(id string, GetDeviceInterfaceVLANsV1QueryParams *GetDeviceInterfaceVLANsV1QueryParams) (*ResponseDevicesGetDeviceInterfaceVLANsV1, *resty.Response, error) {
	return s.GetDeviceInterfaceVLANsV1(id, GetDeviceInterfaceVLANsV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `CountTheNumberOfNetworkDevicesWithFiltersV1`
*/
func (s *DevicesService) CountTheNumberOfNetworkDevicesWithFilters(requestDevicesCountTheNumberOfNetworkDevicesWithFiltersV1 *RequestDevicesCountTheNumberOfNetworkDevicesWithFiltersV1) (*ResponseDevicesCountTheNumberOfNetworkDevicesWithFiltersV1, *resty.Response, error) {
	return s.CountTheNumberOfNetworkDevicesWithFiltersV1(requestDevicesCountTheNumberOfNetworkDevicesWithFiltersV1)
}

// Alias Function
/*
This method acts as an alias for the method `CreateMaintenanceScheduleForNetworkDevicesV1`
*/
func (s *DevicesService) CreateMaintenanceScheduleForNetworkDevices(requestDevicesCreateMaintenanceScheduleForNetworkDevicesV1 *RequestDevicesCreateMaintenanceScheduleForNetworkDevicesV1) (*ResponseDevicesCreateMaintenanceScheduleForNetworkDevicesV1, *resty.Response, error) {
	return s.CreateMaintenanceScheduleForNetworkDevicesV1(requestDevicesCreateMaintenanceScheduleForNetworkDevicesV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1`
*/
func (s *DevicesService) GetsTheSummaryAnalyticsDataRelatedToNetworkDevices(requestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1 *RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1) (*ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1, *resty.Response, error) {
	return s.GetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1(requestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetLinecardDetailsV1`
*/
func (s *DevicesService) GetLinecardDetails(deviceUUID string) (*ResponseDevicesGetLinecardDetailsV1, *resty.Response, error) {
	return s.GetLinecardDetailsV1(deviceUUID)
}

// Alias Function
/*
This method acts as an alias for the method `DeletePlannedAccessPointForFloorV1`
*/
func (s *DevicesService) DeletePlannedAccessPointForFloor(floorID string, plannedAccessPointUUID string) (*ResponseDevicesDeletePlannedAccessPointForFloorV1, *resty.Response, error) {
	return s.DeletePlannedAccessPointForFloorV1(floorID, plannedAccessPointUUID)
}

// Alias Function
/*
This method acts as an alias for the method `UpdatesTheMaintenanceScheduleInformationV1`
*/
func (s *DevicesService) UpdatesTheMaintenanceScheduleInformation(id string, requestDevicesUpdatesTheMaintenanceScheduleInformationV1 *RequestDevicesUpdatesTheMaintenanceScheduleInformationV1) (*ResponseDevicesUpdatesTheMaintenanceScheduleInformationV1, *resty.Response, error) {
	return s.UpdatesTheMaintenanceScheduleInformationV1(id, requestDevicesUpdatesTheMaintenanceScheduleInformationV1)
}

// Alias Function
/*
This method acts as an alias for the method `LegitOperationsForInterfaceV1`
*/
func (s *DevicesService) LegitOperationsForInterface(interfaceUUID string) (*ResponseDevicesLegitOperationsForInterfaceV1, *resty.Response, error) {
	return s.LegitOperationsForInterfaceV1(interfaceUUID)
}

// Alias Function
/*
This method acts as an alias for the method `OverrideResyncIntervalV1`
*/
func (s *DevicesService) OverrideResyncInterval() (*ResponseDevicesOverrideResyncIntervalV1, *resty.Response, error) {
	return s.OverrideResyncIntervalV1()
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesTheMaintenanceScheduleInformationV1`
*/
func (s *DevicesService) RetrievesTheMaintenanceScheduleInformation(id string) (*ResponseDevicesRetrievesTheMaintenanceScheduleInformationV1, *resty.Response, error) {
	return s.RetrievesTheMaintenanceScheduleInformationV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `GetThreatTypesV1`
*/
func (s *DevicesService) GetThreatTypes() (*ResponseDevicesGetThreatTypesV1, *resty.Response, error) {
	return s.GetThreatTypesV1()
}

// Alias Function
/*
This method acts as an alias for the method `GetModulesV1`
*/
func (s *DevicesService) GetModules(GetModulesV1QueryParams *GetModulesV1QueryParams) (*ResponseDevicesGetModulesV1, *resty.Response, error) {
	return s.GetModulesV1(GetModulesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `UpdatePlannedAccessPointForFloorV1`
*/
func (s *DevicesService) UpdatePlannedAccessPointForFloor(floorID string, requestDevicesUpdatePlannedAccessPointForFloorV1 *RequestDevicesUpdatePlannedAccessPointForFloorV1) (*ResponseDevicesUpdatePlannedAccessPointForFloorV1, *resty.Response, error) {
	return s.UpdatePlannedAccessPointForFloorV1(floorID, requestDevicesUpdatePlannedAccessPointForFloorV1)
}

// Alias Function
/*
This method acts as an alias for the method `SyncDevicesV1`
*/
func (s *DevicesService) SyncDevices(requestDevicesSyncDevicesV1 *RequestDevicesSyncDevicesV1, SyncDevicesV1QueryParams *SyncDevicesV1QueryParams) (*ResponseDevicesSyncDevicesV1, *resty.Response, error) {
	return s.SyncDevicesV1(requestDevicesSyncDevicesV1, SyncDevicesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `CreatePlannedAccessPointForFloorV1`
*/
func (s *DevicesService) CreatePlannedAccessPointForFloor(floorID string, requestDevicesCreatePlannedAccessPointForFloorV1 *RequestDevicesCreatePlannedAccessPointForFloorV1) (*ResponseDevicesCreatePlannedAccessPointForFloorV1, *resty.Response, error) {
	return s.CreatePlannedAccessPointForFloorV1(floorID, requestDevicesCreatePlannedAccessPointForFloorV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1`
*/
func (s *DevicesService) GetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions(requestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1 *RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1) (*ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1, *resty.Response, error) {
	return s.GetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1(requestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1)
}

// Alias Function
/*
This method acts as an alias for the method `RetrieveNetworkDevicesV1`
*/
func (s *DevicesService) RetrieveNetworkDevices(RetrieveNetworkDevicesV1QueryParams *RetrieveNetworkDevicesV1QueryParams) (*ResponseDevicesRetrieveNetworkDevicesV1, *resty.Response, error) {
	return s.RetrieveNetworkDevicesV1(RetrieveNetworkDevicesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetPollingIntervalByIDV1`
*/
func (s *DevicesService) GetPollingIntervalByID(id string) (*ResponseDevicesGetPollingIntervalByIDV1, *resty.Response, error) {
	return s.GetPollingIntervalByIDV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersV1`
*/
func (s *DevicesService) GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFilters(GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersV1HeaderParams *GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersV1HeaderParams, GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersV1QueryParams *GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersV1QueryParams) (*ResponseDevicesGetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersV1, *resty.Response, error) {
	return s.GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersV1(GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersV1HeaderParams, GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `ThreatSummaryV1`
*/
func (s *DevicesService) ThreatSummary(requestDevicesThreatSummaryV1 *RequestDevicesThreatSummaryV1) (*ResponseDevicesThreatSummaryV1, *resty.Response, error) {
	return s.ThreatSummaryV1(requestDevicesThreatSummaryV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetInterfaceInfoByIDV1`
*/
func (s *DevicesService) GetInterfaceInfoByID(deviceID string) (*ResponseDevicesGetInterfaceInfoByIDV1, *resty.Response, error) {
	return s.GetInterfaceInfoByIDV1(deviceID)
}

// Alias Function
/*
This method acts as an alias for the method `TheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeV1`
*/
func (s *DevicesService) TheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRange(id string, requestDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeV1 *RequestDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeV1) (*ResponseDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeV1, *resty.Response, error) {
	return s.TheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeV1(id, requestDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetHealthScoreDefinitionForTheGivenIDV1`
*/
func (s *DevicesService) GetHealthScoreDefinitionForTheGivenID(id string, GetHealthScoreDefinitionForTheGivenIdV1HeaderParams *GetHealthScoreDefinitionForTheGivenIDV1HeaderParams) (*ResponseDevicesGetHealthScoreDefinitionForTheGivenIDV1, *resty.Response, error) {
	return s.GetHealthScoreDefinitionForTheGivenIDV1(id, GetHealthScoreDefinitionForTheGivenIdV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1`
*/
func (s *DevicesService) GetsTheTopNAnalyticsDataRelatedToNetworkDevices(requestDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1 *RequestDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1) (*ResponseDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1, *resty.Response, error) {
	return s.GetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1(requestDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesV1)
}

// Alias Function
/*
This method acts as an alias for the method `CountTheNumberOfEventsWithFiltersV1`
*/
func (s *DevicesService) CountTheNumberOfEventsWithFilters(requestDevicesCountTheNumberOfEventsWithFiltersV1 *RequestDevicesCountTheNumberOfEventsWithFiltersV1, CountTheNumberOfEventsWithFiltersV1HeaderParams *CountTheNumberOfEventsWithFiltersV1HeaderParams) (*ResponseDevicesCountTheNumberOfEventsWithFiltersV1, *resty.Response, error) {
	return s.CountTheNumberOfEventsWithFiltersV1(requestDevicesCountTheNumberOfEventsWithFiltersV1, CountTheNumberOfEventsWithFiltersV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `RemoveUserDefinedFieldFromDeviceV1`
*/
func (s *DevicesService) RemoveUserDefinedFieldFromDevice(deviceID string, RemoveUserDefinedFieldFromDeviceV1QueryParams *RemoveUserDefinedFieldFromDeviceV1QueryParams) (*ResponseDevicesRemoveUserDefinedFieldFromDeviceV1, *resty.Response, error) {
	return s.RemoveUserDefinedFieldFromDeviceV1(deviceID, RemoveUserDefinedFieldFromDeviceV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `DeleteUserDefinedFieldV1`
*/
func (s *DevicesService) DeleteUserDefinedField(id string) (*ResponseDevicesDeleteUserDefinedFieldV1, *resty.Response, error) {
	return s.DeleteUserDefinedFieldV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `QueryNetworkDevicesWithFiltersV1`
*/
func (s *DevicesService) QueryNetworkDevicesWithFilters(requestDevicesQueryNetworkDevicesWithFiltersV1 *RequestDevicesQueryNetworkDevicesWithFiltersV1) (*ResponseDevicesQueryNetworkDevicesWithFiltersV1, *resty.Response, error) {
	return s.QueryNetworkDevicesWithFiltersV1(requestDevicesQueryNetworkDevicesWithFiltersV1)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1`
*/
func (s *DevicesService) RetrievesTheListOfDNSServicesForGivenSetOfComplexFilters(requestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1 *RequestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1, RetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams *RetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams) (*ResponseDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1, *resty.Response, error) {
	return s.RetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1(requestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1, RetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1`
*/
func (s *DevicesService) RetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFilters(requestDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1 *RequestDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1, RetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams *RetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams) (*ResponseDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1, *resty.Response, error) {
	return s.RetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1(requestDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1, RetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `DeleteDeviceByIDV1`
*/
func (s *DevicesService) DeleteDeviceByID(id string, DeleteDeviceByIdV1QueryParams *DeleteDeviceByIDV1QueryParams) (*ResponseDevicesDeleteDeviceByIDV1, *resty.Response, error) {
	return s.DeleteDeviceByIDV1(id, DeleteDeviceByIdV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `UpdateGlobalResyncIntervalV1`
*/
func (s *DevicesService) UpdateGlobalResyncInterval(requestDevicesUpdateGlobalResyncIntervalV1 *RequestDevicesUpdateGlobalResyncIntervalV1) (*ResponseDevicesUpdateGlobalResyncIntervalV1, *resty.Response, error) {
	return s.UpdateGlobalResyncIntervalV1(requestDevicesUpdateGlobalResyncIntervalV1)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesTheListOfDNSServicesForGivenParametersV1`
*/
func (s *DevicesService) RetrievesTheListOfDNSServicesForGivenParameters(RetrievesTheListOfDNSServicesForGivenParametersV1HeaderParams *RetrievesTheListOfDNSServicesForGivenParametersV1HeaderParams, RetrievesTheListOfDNSServicesForGivenParametersV1QueryParams *RetrievesTheListOfDNSServicesForGivenParametersV1QueryParams) (*ResponseDevicesRetrievesTheListOfDNSServicesForGivenParametersV1, *resty.Response, error) {
	return s.RetrievesTheListOfDNSServicesForGivenParametersV1(RetrievesTheListOfDNSServicesForGivenParametersV1HeaderParams, RetrievesTheListOfDNSServicesForGivenParametersV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1`
*/
func (s *DevicesService) GetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters(requestDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1 *RequestDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1, GetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams *GetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams) (*ResponseDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1, *resty.Response, error) {
	return s.GetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1(requestDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1, GetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetModuleInfoByIDV1`
*/
func (s *DevicesService) GetModuleInfoByID(id string) (*ResponseDevicesGetModuleInfoByIDV1, *resty.Response, error) {
	return s.GetModuleInfoByIDV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `GetDeviceListV1`
*/
func (s *DevicesService) GetDeviceList(GetDeviceListV1QueryParams *GetDeviceListV1QueryParams) (*ResponseDevicesGetDeviceListV1, *resty.Response, error) {
	return s.GetDeviceListV1(GetDeviceListV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountV1`
*/
func (s *DevicesService) GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCount(GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountV1QueryParams *GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountV1QueryParams) (*ResponseDevicesGetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountV1, *resty.Response, error) {
	return s.GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountV1(GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `AddDeviceKnowYourNetworkV1`
*/
func (s *DevicesService) AddDeviceKnowYourNetwork(requestDevicesAddDeviceKnowYourNetworkV1 *RequestDevicesAddDeviceKnowYourNetworkV1) (*ResponseDevicesAddDeviceKnowYourNetworkV1, *resty.Response, error) {
	return s.AddDeviceKnowYourNetworkV1(requestDevicesAddDeviceKnowYourNetworkV1)
}

// Alias Function
/*
This method acts as an alias for the method `UpdateResyncIntervalForTheNetworkDeviceV1`
*/
func (s *DevicesService) UpdateResyncIntervalForTheNetworkDevice(id string, requestDevicesUpdateResyncIntervalForTheNetworkDeviceV1 *RequestDevicesUpdateResyncIntervalForTheNetworkDeviceV1) (*ResponseDevicesUpdateResyncIntervalForTheNetworkDeviceV1, *resty.Response, error) {
	return s.UpdateResyncIntervalForTheNetworkDeviceV1(id, requestDevicesUpdateResyncIntervalForTheNetworkDeviceV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1`
*/
func (s *DevicesService) GetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters(requestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1 *RequestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1, GetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams *GetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams) (*ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1, *resty.Response, error) {
	return s.GetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1(requestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1, GetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `TheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1`
*/
func (s *DevicesService) TheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRange(id string, requestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1 *RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1) (*ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1, *resty.Response, error) {
	return s.TheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1(id, requestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetIsisInterfacesV1`
*/
func (s *DevicesService) GetIsisInterfaces() (*ResponseDevicesGetIsisInterfacesV1, *resty.Response, error) {
	return s.GetIsisInterfacesV1()
}

// Alias Function
/*
This method acts as an alias for the method `GetDeviceBySerialNumberV1`
*/
func (s *DevicesService) GetDeviceBySerialNumber(serialNumber string) (*ResponseDevicesGetDeviceBySerialNumberV1, *resty.Response, error) {
	return s.GetDeviceBySerialNumberV1(serialNumber)
}

// Alias Function
/*
This method acts as an alias for the method `GetDeviceConfigCountV1`
*/
func (s *DevicesService) GetDeviceConfigCount() (*ResponseDevicesGetDeviceConfigCountV1, *resty.Response, error) {
	return s.GetDeviceConfigCountV1()
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIDOfTheServiceV1`
*/
func (s *DevicesService) RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIDOfTheService(id string, RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIdOfTheServiceV1HeaderParams *RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIDOfTheServiceV1HeaderParams, RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIdOfTheServiceV1QueryParams *RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIDOfTheServiceV1QueryParams) (*ResponseDevicesRetrievesTheDetailsOfASpecificDNSServiceMatchingTheIDOfTheServiceV1, *resty.Response, error) {
	return s.RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIDOfTheServiceV1(id, RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIdOfTheServiceV1HeaderParams, RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIdOfTheServiceV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetChassisDetailsForDeviceV1`
*/
func (s *DevicesService) GetChassisDetailsForDevice(deviceID string) (*ResponseDevicesGetChassisDetailsForDeviceV1, *resty.Response, error) {
	return s.GetChassisDetailsForDeviceV1(deviceID)
}

// Alias Function
/*
This method acts as an alias for the method `GetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1`
*/
func (s *DevicesService) GetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheService(id string, requestDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIdOfTheServiceV1 *RequestDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1, GetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIdOfTheServiceV1HeaderParams *GetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1HeaderParams) (*ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1, *resty.Response, error) {
	return s.GetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceV1(id, requestDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIdOfTheServiceV1, GetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIdOfTheServiceV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetDeviceConfigByIDV1`
*/
func (s *DevicesService) GetDeviceConfigByID(networkDeviceID string) (*ResponseDevicesGetDeviceConfigByIDV1, *resty.Response, error) {
	return s.GetDeviceConfigByIDV1(networkDeviceID)
}

// Alias Function
/*
This method acts as an alias for the method `RogueAdditionalDetailsV1`
*/
func (s *DevicesService) RogueAdditionalDetails(requestDevicesRogueAdditionalDetailsV1 *RequestDevicesRogueAdditionalDetailsV1) (*ResponseDevicesRogueAdditionalDetailsV1, *resty.Response, error) {
	return s.RogueAdditionalDetailsV1(requestDevicesRogueAdditionalDetailsV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetAllHealthScoreDefinitionsForGivenFiltersV1`
*/
func (s *DevicesService) GetAllHealthScoreDefinitionsForGivenFilters(GetAllHealthScoreDefinitionsForGivenFiltersV1HeaderParams *GetAllHealthScoreDefinitionsForGivenFiltersV1HeaderParams, GetAllHealthScoreDefinitionsForGivenFiltersV1QueryParams *GetAllHealthScoreDefinitionsForGivenFiltersV1QueryParams) (*ResponseDevicesGetAllHealthScoreDefinitionsForGivenFiltersV1, *resty.Response, error) {
	return s.GetAllHealthScoreDefinitionsForGivenFiltersV1(GetAllHealthScoreDefinitionsForGivenFiltersV1HeaderParams, GetAllHealthScoreDefinitionsForGivenFiltersV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetPlannedAccessPointsForBuildingV1`
*/
func (s *DevicesService) GetPlannedAccessPointsForBuilding(buildingID string, GetPlannedAccessPointsForBuildingV1QueryParams *GetPlannedAccessPointsForBuildingV1QueryParams) (*ResponseDevicesGetPlannedAccessPointsForBuildingV1, *resty.Response, error) {
	return s.GetPlannedAccessPointsForBuildingV1(buildingID, GetPlannedAccessPointsForBuildingV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `RemoveAllowedMacAddressV1`
*/
func (s *DevicesService) RemoveAllowedMacAddress(macAddress string) (*ResponseDevicesRemoveAllowedMacAddressV1, *resty.Response, error) {
	return s.RemoveAllowedMacAddressV1(macAddress)
}

// Alias Function
/*
This method acts as an alias for the method `DevicesV1`
*/
func (s *DevicesService) Devices(DevicesV1QueryParams *DevicesV1QueryParams) (*ResponseDevicesDevicesV1, *resty.Response, error) {
	return s.DevicesV1(DevicesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `DeleteANetworkDeviceWithoutConfigurationCleanupV1`
*/
func (s *DevicesService) DeleteANetworkDeviceWithoutConfigurationCleanup(requestDevicesDeleteANetworkDeviceWithoutConfigurationCleanupV1 *RequestDevicesDeleteANetworkDeviceWithoutConfigurationCleanupV1) (*ResponseDevicesDeleteANetworkDeviceWithoutConfigurationCleanupV1, *resty.Response, error) {
	return s.DeleteANetworkDeviceWithoutConfigurationCleanupV1(requestDevicesDeleteANetworkDeviceWithoutConfigurationCleanupV1)
}

// Alias Function
/*
This method acts as an alias for the method `UpdateDeviceRoleV1`
*/
func (s *DevicesService) UpdateDeviceRole(requestDevicesUpdateDeviceRoleV1 *RequestDevicesUpdateDeviceRoleV1) (*ResponseDevicesUpdateDeviceRoleV1, *resty.Response, error) {
	return s.UpdateDeviceRoleV1(requestDevicesUpdateDeviceRoleV1)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesTheTotalNumberOfDHCPServicesForGivenParametersV1`
*/
func (s *DevicesService) RetrievesTheTotalNumberOfDHCPServicesForGivenParameters(RetrievesTheTotalNumberOfDHCPServicesForGivenParametersV1HeaderParams *RetrievesTheTotalNumberOfDHCPServicesForGivenParametersV1HeaderParams, RetrievesTheTotalNumberOfDHCPServicesForGivenParametersV1QueryParams *RetrievesTheTotalNumberOfDHCPServicesForGivenParametersV1QueryParams) (*ResponseDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenParametersV1, *resty.Response, error) {
	return s.RetrievesTheTotalNumberOfDHCPServicesForGivenParametersV1(RetrievesTheTotalNumberOfDHCPServicesForGivenParametersV1HeaderParams, RetrievesTheTotalNumberOfDHCPServicesForGivenParametersV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1`
*/
func (s *DevicesService) GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParameters(GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1QueryParams *GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1QueryParams) (*ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1, *resty.Response, error) {
	return s.GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1(GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetDeviceInterfaceCountForMultipleDevicesV1`
*/
func (s *DevicesService) GetDeviceInterfaceCountForMultipleDevices() (*ResponseDevicesGetDeviceInterfaceCountForMultipleDevicesV1, *resty.Response, error) {
	return s.GetDeviceInterfaceCountForMultipleDevicesV1()
}

// Alias Function
/*
This method acts as an alias for the method `CountTheNumberOfNetworkDevicesV1`
*/
func (s *DevicesService) CountTheNumberOfNetworkDevices(CountTheNumberOfNetworkDevicesV1QueryParams *CountTheNumberOfNetworkDevicesV1QueryParams) (*ResponseDevicesCountTheNumberOfNetworkDevicesV1, *resty.Response, error) {
	return s.CountTheNumberOfNetworkDevicesV1(CountTheNumberOfNetworkDevicesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetResyncIntervalForTheNetworkDeviceV1`
*/
func (s *DevicesService) GetResyncIntervalForTheNetworkDevice(id string) (*ResponseDevicesGetResyncIntervalForTheNetworkDeviceV1, *resty.Response, error) {
	return s.GetResyncIntervalForTheNetworkDeviceV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `GetAllUserDefinedFieldsV1`
*/
func (s *DevicesService) GetAllUserDefinedFields(GetAllUserDefinedFieldsV1QueryParams *GetAllUserDefinedFieldsV1QueryParams) (*ResponseDevicesGetAllUserDefinedFieldsV1, *resty.Response, error) {
	return s.GetAllUserDefinedFieldsV1(GetAllUserDefinedFieldsV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1`
*/
func (s *DevicesService) GetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions(requestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1 *RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1) (*ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1, *resty.Response, error) {
	return s.GetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1(requestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1`
*/
func (s *DevicesService) RetrievesTheListOfDHCPServicesForGivenSetOfComplexFilters(requestDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1 *RequestDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1, RetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams *RetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams) (*ResponseDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1, *resty.Response, error) {
	return s.RetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1(requestDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1, RetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetDeviceValuesThatMatchFullyOrPartiallyAnAttributeV1`
*/
func (s *DevicesService) GetDeviceValuesThatMatchFullyOrPartiallyAnAttribute(GetDeviceValuesThatMatchFullyOrPartiallyAnAttributeV1QueryParams *GetDeviceValuesThatMatchFullyOrPartiallyAnAttributeV1QueryParams) (*ResponseDevicesGetDeviceValuesThatMatchFullyOrPartiallyAnAttributeV1, *resty.Response, error) {
	return s.GetDeviceValuesThatMatchFullyOrPartiallyAnAttributeV1(GetDeviceValuesThatMatchFullyOrPartiallyAnAttributeV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `ReturnsPoeInterfaceDetailsForTheDeviceV1`
*/
func (s *DevicesService) ReturnsPoeInterfaceDetailsForTheDevice(deviceUUID string, ReturnsPOEInterfaceDetailsForTheDeviceV1QueryParams *ReturnsPoeInterfaceDetailsForTheDeviceV1QueryParams) (*ResponseDevicesReturnsPoeInterfaceDetailsForTheDeviceV1, *resty.Response, error) {
	return s.ReturnsPoeInterfaceDetailsForTheDeviceV1(deviceUUID, ReturnsPOEInterfaceDetailsForTheDeviceV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetModuleCountV1`
*/
func (s *DevicesService) GetModuleCount(GetModuleCountV1QueryParams *GetModuleCountV1QueryParams) (*ResponseDevicesGetModuleCountV1, *resty.Response, error) {
	return s.GetModuleCountV1(GetModuleCountV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetPollingIntervalForAllDevicesV1`
*/
func (s *DevicesService) GetPollingIntervalForAllDevices() (*ResponseDevicesGetPollingIntervalForAllDevicesV1, *resty.Response, error) {
	return s.GetPollingIntervalForAllDevicesV1()
}

// Alias Function
/*
This method acts as an alias for the method `GetPlannedAccessPointsForFloorV1`
*/
func (s *DevicesService) GetPlannedAccessPointsForFloor(floorID string, GetPlannedAccessPointsForFloorV1QueryParams *GetPlannedAccessPointsForFloorV1QueryParams) (*ResponseDevicesGetPlannedAccessPointsForFloorV1, *resty.Response, error) {
	return s.GetPlannedAccessPointsForFloorV1(floorID, GetPlannedAccessPointsForFloorV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `ClearMacAddressTableV1`
*/
func (s *DevicesService) ClearMacAddressTable(interfaceUUID string, requestDevicesClearMacAddressTableV1 *RequestDevicesClearMacAddressTableV1, ClearMacAddressTableV1QueryParams *ClearMacAddressTableV1QueryParams) (*ResponseDevicesClearMacAddressTableV1, *resty.Response, error) {
	return s.ClearMacAddressTableV1(interfaceUUID, requestDevicesClearMacAddressTableV1, ClearMacAddressTableV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `WirelessRogueApContainmentStatusV1`
*/
func (s *DevicesService) WirelessRogueApContainmentStatus(macAddress string) (*ResponseDevicesWirelessRogueApContainmentStatusV1, *resty.Response, error) {
	return s.WirelessRogueApContainmentStatusV1(macAddress)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIDOfTheServiceV1`
*/
func (s *DevicesService) RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIDOfTheService(id string, RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIdOfTheServiceV1HeaderParams *RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIDOfTheServiceV1HeaderParams, RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIdOfTheServiceV1QueryParams *RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIDOfTheServiceV1QueryParams) (*ResponseDevicesRetrievesTheDetailsOfASpecificAAAServiceMatchingTheIDOfTheServiceV1, *resty.Response, error) {
	return s.RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIDOfTheServiceV1(id, RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIdOfTheServiceV1HeaderParams, RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIdOfTheServiceV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetsTheTrendAnalyticsDataV1`
*/
func (s *DevicesService) GetsTheTrendAnalyticsData(requestDevicesGetsTheTrendAnalyticsDataV1 *RequestDevicesGetsTheTrendAnalyticsDataV1) (*ResponseDevicesGetsTheTrendAnalyticsDataV1, *resty.Response, error) {
	return s.GetsTheTrendAnalyticsDataV1(requestDevicesGetsTheTrendAnalyticsDataV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetDeviceConfigForAllDevicesV1`
*/
func (s *DevicesService) GetDeviceConfigForAllDevices() (*ResponseDevicesGetDeviceConfigForAllDevicesV1, *resty.Response, error) {
	return s.GetDeviceConfigForAllDevicesV1()
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersV1`
*/
func (s *DevicesService) RetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFilters(requestDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersV1 *RequestDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersV1, RetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams *RetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams) (*ResponseDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersV1, *resty.Response, error) {
	return s.RetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersV1(requestDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersV1, RetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `AddUserDefinedFieldToDeviceV1`
*/
func (s *DevicesService) AddUserDefinedFieldToDevice(deviceID string, requestDevicesAddUserDefinedFieldToDeviceV1 *RequestDevicesAddUserDefinedFieldToDeviceV1) (*ResponseDevicesAddUserDefinedFieldToDeviceV1, *resty.Response, error) {
	return s.AddUserDefinedFieldToDeviceV1(deviceID, requestDevicesAddUserDefinedFieldToDeviceV1)
}

// Alias Function
/*
This method acts as an alias for the method `RetrieveScheduledMaintenanceWindowsForNetworkDevicesV1`
*/
func (s *DevicesService) RetrieveScheduledMaintenanceWindowsForNetworkDevices(RetrieveScheduledMaintenanceWindowsForNetworkDevicesV1QueryParams *RetrieveScheduledMaintenanceWindowsForNetworkDevicesV1QueryParams) (*ResponseDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesV1, *resty.Response, error) {
	return s.RetrieveScheduledMaintenanceWindowsForNetworkDevicesV1(RetrieveScheduledMaintenanceWindowsForNetworkDevicesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1`
*/
func (s *DevicesService) GetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheService(id string, requestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIdOfTheServiceV1 *RequestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1, GetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIdOfTheServiceV1HeaderParams *GetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1HeaderParams) (*ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1, *resty.Response, error) {
	return s.GetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceV1(id, requestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIdOfTheServiceV1, GetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIdOfTheServiceV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetStackDetailsForDeviceV1`
*/
func (s *DevicesService) GetStackDetailsForDevice(deviceID string) (*ResponseDevicesGetStackDetailsForDeviceV1, *resty.Response, error) {
	return s.GetStackDetailsForDeviceV1(deviceID)
}

// Alias Function
/*
This method acts as an alias for the method `RogueAdditionalDetailCountV1`
*/
func (s *DevicesService) RogueAdditionalDetailCount(requestDevicesRogueAdditionalDetailCountV1 *RequestDevicesRogueAdditionalDetailCountV1) (*ResponseDevicesRogueAdditionalDetailCountV1, *resty.Response, error) {
	return s.RogueAdditionalDetailCountV1(requestDevicesRogueAdditionalDetailCountV1)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesTheListOfAAAServicesForGivenParametersV1`
*/
func (s *DevicesService) RetrievesTheListOfAAAServicesForGivenParameters(RetrievesTheListOfAAAServicesForGivenParametersV1HeaderParams *RetrievesTheListOfAAAServicesForGivenParametersV1HeaderParams, RetrievesTheListOfAAAServicesForGivenParametersV1QueryParams *RetrievesTheListOfAAAServicesForGivenParametersV1QueryParams) (*ResponseDevicesRetrievesTheListOfAAAServicesForGivenParametersV1, *resty.Response, error) {
	return s.RetrievesTheListOfAAAServicesForGivenParametersV1(RetrievesTheListOfAAAServicesForGivenParametersV1HeaderParams, RetrievesTheListOfAAAServicesForGivenParametersV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `RetrieveTheTotalNumberOfScheduledMaintenanceWindowsV1`
*/
func (s *DevicesService) RetrieveTheTotalNumberOfScheduledMaintenanceWindows(RetrieveTheTotalNumberOfScheduledMaintenanceWindowsV1QueryParams *RetrieveTheTotalNumberOfScheduledMaintenanceWindowsV1QueryParams) (*ResponseDevicesRetrieveTheTotalNumberOfScheduledMaintenanceWindowsV1, *resty.Response, error) {
	return s.RetrieveTheTotalNumberOfScheduledMaintenanceWindowsV1(RetrieveTheTotalNumberOfScheduledMaintenanceWindowsV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetListOfChildEventsForTheGivenWirelessClientEventV1`
*/
func (s *DevicesService) GetListOfChildEventsForTheGivenWirelessClientEvent(id string, GetListOfChildEventsForTheGivenWirelessClientEventV1HeaderParams *GetListOfChildEventsForTheGivenWirelessClientEventV1HeaderParams) (*ResponseDevicesGetListOfChildEventsForTheGivenWirelessClientEventV1, *resty.Response, error) {
	return s.GetListOfChildEventsForTheGivenWirelessClientEventV1(id, GetListOfChildEventsForTheGivenWirelessClientEventV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetNetworkDeviceByIPV1`
*/
func (s *DevicesService) GetNetworkDeviceByIP(ipAddress string) (*ResponseDevicesGetNetworkDeviceByIPV1, *resty.Response, error) {
	return s.GetNetworkDeviceByIPV1(ipAddress)
}

// Alias Function
/*
This method acts as an alias for the method `GetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1`
*/
func (s *DevicesService) GetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters(requestDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1 *RequestDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1, GetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams *GetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams) (*ResponseDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1, *resty.Response, error) {
	return s.GetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1(requestDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1, GetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetDeviceSummaryV1`
*/
func (s *DevicesService) GetDeviceSummary(id string) (*ResponseDevicesGetDeviceSummaryV1, *resty.Response, error) {
	return s.GetDeviceSummaryV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `GetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesV1`
*/
func (s *DevicesService) GetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevices(GetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesV1QueryParams *GetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesV1QueryParams) (*ResponseDevicesGetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesV1, *resty.Response, error) {
	return s.GetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesV1(GetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `AddAllowedMacAddressV1`
*/
func (s *DevicesService) AddAllowedMacAddress(requestDevicesAddAllowedMacAddressV1 *RequestDevicesAddAllowedMacAddressV1) (*ResponseDevicesAddAllowedMacAddressV1, *resty.Response, error) {
	return s.AddAllowedMacAddressV1(requestDevicesAddAllowedMacAddressV1)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheServiceV1`
*/
func (s *DevicesService) RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheService(id string, RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIdOfTheServiceV1HeaderParams *RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheServiceV1HeaderParams, RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIdOfTheServiceV1QueryParams *RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheServiceV1QueryParams) (*ResponseDevicesRetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheServiceV1, *resty.Response, error) {
	return s.RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheServiceV1(id, RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIdOfTheServiceV1HeaderParams, RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIdOfTheServiceV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetFunctionalCapabilityByIDV1`
*/
func (s *DevicesService) GetFunctionalCapabilityByID(id string) (*ResponseDevicesGetFunctionalCapabilityByIDV1, *resty.Response, error) {
	return s.GetFunctionalCapabilityByIDV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `QueryAssuranceEventsV1`
*/
func (s *DevicesService) QueryAssuranceEvents(QueryAssuranceEventsV1HeaderParams *QueryAssuranceEventsV1HeaderParams, QueryAssuranceEventsV1QueryParams *QueryAssuranceEventsV1QueryParams) (*ResponseDevicesQueryAssuranceEventsV1, *resty.Response, error) {
	return s.QueryAssuranceEventsV1(QueryAssuranceEventsV1HeaderParams, QueryAssuranceEventsV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetAllowedMacAddressCountV1`
*/
func (s *DevicesService) GetAllowedMacAddressCount() (*ResponseDevicesGetAllowedMacAddressCountV1, *resty.Response, error) {
	return s.GetAllowedMacAddressCountV1()
}

// Alias Function
/*
This method acts as an alias for the method `UpdateHealthScoreDefinitionsV1`
*/
func (s *DevicesService) UpdateHealthScoreDefinitions(requestDevicesUpdateHealthScoreDefinitionsV1 *RequestDevicesUpdateHealthScoreDefinitionsV1, UpdateHealthScoreDefinitionsV1HeaderParams *UpdateHealthScoreDefinitionsV1HeaderParams) (*ResponseDevicesUpdateHealthScoreDefinitionsV1, *resty.Response, error) {
	return s.UpdateHealthScoreDefinitionsV1(requestDevicesUpdateHealthScoreDefinitionsV1, UpdateHealthScoreDefinitionsV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetDeviceInterfaceCountV1`
*/
func (s *DevicesService) GetDeviceInterfaceCount(deviceID string) (*ResponseDevicesGetDeviceInterfaceCountV1, *resty.Response, error) {
	return s.GetDeviceInterfaceCountV1(deviceID)
}

// Alias Function
/*
This method acts as an alias for the method `GetInterfaceByIDV1`
*/
func (s *DevicesService) GetInterfaceByID(id string) (*ResponseDevicesGetInterfaceByIDV1, *resty.Response, error) {
	return s.GetInterfaceByIDV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `GetAllowedMacAddressV1`
*/
func (s *DevicesService) GetAllowedMacAddress(GetAllowedMacAddressV1QueryParams *GetAllowedMacAddressV1QueryParams) (*ResponseDevicesGetAllowedMacAddressV1, *resty.Response, error) {
	return s.GetAllowedMacAddressV1(GetAllowedMacAddressV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1`
*/
func (s *DevicesService) GetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters(requestDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1 *RequestDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1, GetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams *GetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams) (*ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1, *resty.Response, error) {
	return s.GetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1(requestDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1, GetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesTheTotalNumberOfAAAServicesForGivenParametersV1`
*/
func (s *DevicesService) RetrievesTheTotalNumberOfAAAServicesForGivenParameters(RetrievesTheTotalNumberOfAAAServicesForGivenParametersV1HeaderParams *RetrievesTheTotalNumberOfAAAServicesForGivenParametersV1HeaderParams, RetrievesTheTotalNumberOfAAAServicesForGivenParametersV1QueryParams *RetrievesTheTotalNumberOfAAAServicesForGivenParametersV1QueryParams) (*ResponseDevicesRetrievesTheTotalNumberOfAAAServicesForGivenParametersV1, *resty.Response, error) {
	return s.RetrievesTheTotalNumberOfAAAServicesForGivenParametersV1(RetrievesTheTotalNumberOfAAAServicesForGivenParametersV1HeaderParams, RetrievesTheTotalNumberOfAAAServicesForGivenParametersV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersV1`
*/
func (s *DevicesService) RetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFilters(requestDevicesRetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersV1 *RequestDevicesRetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersV1, RetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams *RetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams) (*ResponseDevicesRetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersV1, *resty.Response, error) {
	return s.RetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersV1(requestDevicesRetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersV1, RetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetDeviceDetailV1`
*/
func (s *DevicesService) GetDeviceDetail(GetDeviceDetailV1QueryParams *GetDeviceDetailV1QueryParams) (*ResponseDevicesGetDeviceDetailV1, *resty.Response, error) {
	return s.GetDeviceDetailV1(GetDeviceDetailV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `UpdateHealthScoreDefinitionForTheGivenIDV1`
*/
func (s *DevicesService) UpdateHealthScoreDefinitionForTheGivenID(id string, requestDevicesUpdateHealthScoreDefinitionForTheGivenIdV1 *RequestDevicesUpdateHealthScoreDefinitionForTheGivenIDV1) (*ResponseDevicesUpdateHealthScoreDefinitionForTheGivenIDV1, *resty.Response, error) {
	return s.UpdateHealthScoreDefinitionForTheGivenIDV1(id, requestDevicesUpdateHealthScoreDefinitionForTheGivenIdV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1`
*/
func (s *DevicesService) GetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters(requestDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1 *RequestDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1, GetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams *GetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams) (*ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1, *resty.Response, error) {
	return s.GetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1(requestDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1, GetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1`
*/
func (s *DevicesService) GetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters(requestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1 *RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1, GetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams *GetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams) (*ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1, *resty.Response, error) {
	return s.GetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1(requestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1, GetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetDetailsOfASingleAssuranceEventV1`
*/
func (s *DevicesService) GetDetailsOfASingleAssuranceEvent(id string, GetDetailsOfASingleAssuranceEventV1HeaderParams *GetDetailsOfASingleAssuranceEventV1HeaderParams, GetDetailsOfASingleAssuranceEventV1QueryParams *GetDetailsOfASingleAssuranceEventV1QueryParams) (*ResponseDevicesGetDetailsOfASingleAssuranceEventV1, *resty.Response, error) {
	return s.GetDetailsOfASingleAssuranceEventV1(id, GetDetailsOfASingleAssuranceEventV1HeaderParams, GetDetailsOfASingleAssuranceEventV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetDeviceInterfacesBySpecifiedRangeV1`
*/
func (s *DevicesService) GetDeviceInterfacesBySpecifiedRange(deviceID string, startIndex int, recordsToReturn int) (*ResponseDevicesGetDeviceInterfacesBySpecifiedRangeV1, *resty.Response, error) {
	return s.GetDeviceInterfacesBySpecifiedRangeV1(deviceID, startIndex, recordsToReturn)
}

// Alias Function
/*
This method acts as an alias for the method `QueryAssuranceEventsWithFiltersV1`
*/
func (s *DevicesService) QueryAssuranceEventsWithFilters(requestDevicesQueryAssuranceEventsWithFiltersV1 *RequestDevicesQueryAssuranceEventsWithFiltersV1, QueryAssuranceEventsWithFiltersV1HeaderParams *QueryAssuranceEventsWithFiltersV1HeaderParams) (*ResponseDevicesQueryAssuranceEventsWithFiltersV1, *resty.Response, error) {
	return s.QueryAssuranceEventsWithFiltersV1(requestDevicesQueryAssuranceEventsWithFiltersV1, QueryAssuranceEventsWithFiltersV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetDeviceCountKnowYourNetworkV1`
*/
func (s *DevicesService) GetDeviceCountKnowYourNetwork(GetDeviceCountKnowYourNetworkV1QueryParams *GetDeviceCountKnowYourNetworkV1QueryParams) (*ResponseDevicesGetDeviceCountKnowYourNetworkV1, *resty.Response, error) {
	return s.GetDeviceCountKnowYourNetworkV1(GetDeviceCountKnowYourNetworkV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `StartWirelessRogueApContainmentV1`
*/
func (s *DevicesService) StartWirelessRogueApContainment(requestDevicesStartWirelessRogueAPContainmentV1 *RequestDevicesStartWirelessRogueApContainmentV1) (*ResponseDevicesStartWirelessRogueApContainmentV1, *resty.Response, error) {
	return s.StartWirelessRogueApContainmentV1(requestDevicesStartWirelessRogueAPContainmentV1)
}

// Alias Function
/*
This method acts as an alias for the method `CreateUserDefinedFieldV1`
*/
func (s *DevicesService) CreateUserDefinedField(requestDevicesCreateUserDefinedFieldV1 *RequestDevicesCreateUserDefinedFieldV1) (*ResponseDevicesCreateUserDefinedFieldV1, *resty.Response, error) {
	return s.CreateUserDefinedFieldV1(requestDevicesCreateUserDefinedFieldV1)
}

// Alias Function
/*
This method acts as an alias for the method `UpdateDeviceManagementAddressV1`
*/
func (s *DevicesService) UpdateDeviceManagementAddress(deviceid string, requestDevicesUpdateDeviceManagementAddressV1 *RequestDevicesUpdateDeviceManagementAddressV1) (*ResponseDevicesUpdateDeviceManagementAddressV1, *resty.Response, error) {
	return s.UpdateDeviceManagementAddressV1(deviceid, requestDevicesUpdateDeviceManagementAddressV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetInterfaceDetailsByDeviceIDAndInterfaceNameV1`
*/
func (s *DevicesService) GetInterfaceDetailsByDeviceIDAndInterfaceName(deviceID string, GetInterfaceDetailsByDeviceIdAndInterfaceNameV1QueryParams *GetInterfaceDetailsByDeviceIDAndInterfaceNameV1QueryParams) (*ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameV1, *resty.Response, error) {
	return s.GetInterfaceDetailsByDeviceIDAndInterfaceNameV1(deviceID, GetInterfaceDetailsByDeviceIdAndInterfaceNameV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetDetailsOfASingleNetworkDeviceV1`
*/
func (s *DevicesService) GetDetailsOfASingleNetworkDevice(id string, GetDetailsOfASingleNetworkDeviceV1QueryParams *GetDetailsOfASingleNetworkDeviceV1QueryParams) (*ResponseDevicesGetDetailsOfASingleNetworkDeviceV1, *resty.Response, error) {
	return s.GetDetailsOfASingleNetworkDeviceV1(id, GetDetailsOfASingleNetworkDeviceV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `UpdateDeviceDetailsV1`
*/
func (s *DevicesService) UpdateDeviceDetails(requestDevicesUpdateDeviceDetailsV1 *RequestDevicesUpdateDeviceDetailsV1) (*ResponseDevicesUpdateDeviceDetailsV1, *resty.Response, error) {
	return s.UpdateDeviceDetailsV1(requestDevicesUpdateDeviceDetailsV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetTheDetailsOfPhysicalComponentsOfTheGivenDeviceV1`
*/
func (s *DevicesService) GetTheDetailsOfPhysicalComponentsOfTheGivenDevice(deviceUUID string, GetTheDetailsOfPhysicalComponentsOfTheGivenDeviceV1QueryParams *GetTheDetailsOfPhysicalComponentsOfTheGivenDeviceV1QueryParams) (*ResponseDevicesGetTheDetailsOfPhysicalComponentsOfTheGivenDeviceV1, *resty.Response, error) {
	return s.GetTheDetailsOfPhysicalComponentsOfTheGivenDeviceV1(deviceUUID, GetTheDetailsOfPhysicalComponentsOfTheGivenDeviceV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1`
*/
func (s *DevicesService) GetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters(requestDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1 *RequestDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1, GetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams *GetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams) (*ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1, *resty.Response, error) {
	return s.GetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1(requestDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1, GetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `DeleteMaintenanceScheduleV1`
*/
func (s *DevicesService) DeleteMaintenanceSchedule(id string) (*ResponseDevicesDeleteMaintenanceScheduleV1, *resty.Response, error) {
	return s.DeleteMaintenanceScheduleV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `GetFunctionalCapabilityForDevicesV1`
*/
func (s *DevicesService) GetFunctionalCapabilityForDevices(GetFunctionalCapabilityForDevicesV1QueryParams *GetFunctionalCapabilityForDevicesV1QueryParams) (*ResponseDevicesGetFunctionalCapabilityForDevicesV1, *resty.Response, error) {
	return s.GetFunctionalCapabilityForDevicesV1(GetFunctionalCapabilityForDevicesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `UpdateInterfaceDetailsV1`
*/
func (s *DevicesService) UpdateInterfaceDetails(interfaceUUID string, requestDevicesUpdateInterfaceDetailsV1 *RequestDevicesUpdateInterfaceDetailsV1, UpdateInterfaceDetailsV1QueryParams *UpdateInterfaceDetailsV1QueryParams) (*ResponseDevicesUpdateInterfaceDetailsV1, *resty.Response, error) {
	return s.UpdateInterfaceDetailsV1(interfaceUUID, requestDevicesUpdateInterfaceDetailsV1, UpdateInterfaceDetailsV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `CountTheNumberOfEventsV1`
*/
func (s *DevicesService) CountTheNumberOfEvents(CountTheNumberOfEventsV1HeaderParams *CountTheNumberOfEventsV1HeaderParams, CountTheNumberOfEventsV1QueryParams *CountTheNumberOfEventsV1QueryParams) (*ResponseDevicesCountTheNumberOfEventsV1, *resty.Response, error) {
	return s.CountTheNumberOfEventsV1(CountTheNumberOfEventsV1HeaderParams, CountTheNumberOfEventsV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetOspfInterfacesV1`
*/
func (s *DevicesService) GetOspfInterfaces() (*ResponseDevicesGetOspfInterfacesV1, *resty.Response, error) {
	return s.GetOspfInterfacesV1()
}

// Alias Function
/*
This method acts as an alias for the method `GetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsAndPoeDataV1`
*/
func (s *DevicesService) GetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsAndPoeData(id string, GetTheInterfaceDataForTheGivenInterfaceIdinstanceUuidAlongWithTheStatisticsAndPoeDataV1QueryParams *GetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsAndPoeDataV1QueryParams) (*ResponseDevicesGetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsAndPoeDataV1, *resty.Response, error) {
	return s.GetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsAndPoeDataV1(id, GetTheInterfaceDataForTheGivenInterfaceIdinstanceUuidAlongWithTheStatisticsAndPoeDataV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetTheDeviceDataForTheGivenDeviceIDUUIDV1`
*/
func (s *DevicesService) GetTheDeviceDataForTheGivenDeviceIDUUID(id string, GetTheDeviceDataForTheGivenDeviceIdUuidV1QueryParams *GetTheDeviceDataForTheGivenDeviceIDUUIDV1QueryParams) (*ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDV1, *resty.Response, error) {
	return s.GetTheDeviceDataForTheGivenDeviceIDUUIDV1(id, GetTheDeviceDataForTheGivenDeviceIdUuidV1QueryParams)
}
