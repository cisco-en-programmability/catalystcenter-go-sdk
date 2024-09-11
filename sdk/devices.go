package catalyst

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type DevicesService service

type QueryAssuranceEventsQueryParams struct {
	DeviceFamily      string  `url:"deviceFamily,omitempty"`      //Device family. Please note that multiple families across network device type and client type is not allowed. For example, choosing 'Routers' along with 'Wireless Client' or 'Unified AP' is not supported. Examples: 'deviceFamily=Switches and Hubs' (single deviceFamily requested) 'deviceFamily=Switches and Hubs&deviceFamily=Routers' (multiple deviceFamily requested)
	StartTime         float64 `url:"startTime,omitempty"`         //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If 'startTime' is not provided, API will default to current time minus 24 hours.
	EndTime           float64 `url:"endTime,omitempty"`           //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If 'endTime' is not provided, API will default to current time.
	MessageType       string  `url:"messageType,omitempty"`       //Message type for the event. Examples: 'messageType=Syslog' (single messageType requested) 'messageType=Trap&messageType=Syslog' (multiple messageType requested)
	Severity          float64 `url:"severity,omitempty"`          //Severity of the event between 0 and 6. This is applicable only for events related to network devices (other than AP) and 'Wired Client' events. | Value | Severity    | | ----- | ----------- | | 0     | Emergency   | | 1     | Alert       | | 2     | Critical    | | 3     | Error       | | 4     | Warning     | | 5     | Notice      | | 6     | Info        | Examples: 'severity=0' (single severity requested) 'severity=0&severity=1' (multiple severity requested)
	SiteID            string  `url:"siteId,omitempty"`            //The UUID of the site. (Ex. 'flooruuid') Examples: '?siteId=id1' (single siteId requested) '?siteId=id1&siteId=id2&siteId=id3' (multiple siteId requested)
	SiteHierarchyID   string  `url:"siteHierarchyId,omitempty"`   //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. 'globalUuid/areaUuid/buildingUuid/floorUuid') This field supports wildcard asterisk ('*') character search support. E.g. '*uuid*, *uuid, uuid*' Examples: '?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid '(single siteHierarchyId requested) '?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2' (multiple siteHierarchyId requested)
	NetworkDeviceName string  `url:"networkDeviceName,omitempty"` //Network device name. This parameter is applicable for network device related families. This field supports wildcard ('*') character-based search. Ex: '*Branch*' or 'Branch*' or '*Branch' Examples: 'networkDeviceName=Branch-3-Gateway' (single networkDeviceName requested) 'networkDeviceName=Branch-3-Gateway&networkDeviceName=Branch-3-Switch' (multiple networkDeviceName requested)
	NetworkDeviceID   string  `url:"networkDeviceId,omitempty"`   //The list of Network Device Uuids. (Ex. '6bef213c-19ca-4170-8375-b694e251101c') Examples: 'networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c' (single networkDeviceId requested) 'networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c&networkDeviceId=32219612-819e-4b5e-a96b-cf22aca13dd9&networkDeviceId=2541e9a7-b80d-4955-8aa2-79b233318ba0' (multiple networkDeviceId with & separator)
	ApMac             string  `url:"apMac,omitempty"`             //MAC address of the access point. This parameter is applicable for 'Unified AP' and 'Wireless Client' events. This field supports wildcard ('*') character-based search. Ex: '*50:0F*' or '50:0F*' or '*50:0F' Examples: 'apMac=50:0F:80:0F:F7:E0' (single apMac requested) 'apMac=50:0F:80:0F:F7:E0&apMac=18:80:90:AB:7E:A0' (multiple apMac requested)
	ClientMac         string  `url:"clientMac,omitempty"`         //MAC address of the client. This parameter is applicable for 'Wired Client' and 'Wireless Client' events. This field supports wildcard ('*') character-based search. Ex: '*66:2B*' or '66:2B*' or '*66:2B' Examples: 'clientMac=66:2B:B8:D2:01:56' (single clientMac requested) 'clientMac=66:2B:B8:D2:01:56&clientMac=DC:A6:32:F5:5A:89' (multiple clientMac requested)
	Attribute         string  `url:"attribute,omitempty"`         //The list of attributes that needs to be included in the response. If this parameter is not provided, then basic attributes ('id', 'name', 'timestamp', 'details', 'messageType', 'siteHierarchyId', 'siteHierarchy', 'deviceFamily', 'networkDeviceId', 'networkDeviceName', 'managementIpAddress') would be part of the response.  Examples:  'attribute=name' (single attribute requested) 'attribute=name&attribute=networkDeviceName' (multiple attribute requested)
	View              string  `url:"view,omitempty"`              //The list of events views. Please refer to 'EventViews' for the supported list  Examples:  'view=network' (single view requested) 'view=network&view=ap' (multiple view requested)
	Offset            float64 `url:"offset,omitempty"`            //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	Limit             float64 `url:"limit,omitempty"`             //Maximum number of records to return
	SortBy            string  `url:"sortBy,omitempty"`            //A field within the response to sort by.
	Order             string  `url:"order,omitempty"`             //The sort order of the field ascending or descending.
}
type QueryAssuranceEventsHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type CountTheNumberOfEventsQueryParams struct {
	DeviceFamily      string `url:"deviceFamily,omitempty"`      //Device family. Please note that multiple families across network device type and client type is not allowed. For example, choosing 'Routers' along with 'Wireless Client' or 'Unified AP' is not supported. Examples: 'deviceFamily=Switches and Hubs' (single deviceFamily requested) 'deviceFamily=Switches and Hubs&deviceFamily=Routers' (multiple deviceFamily requested)
	StartTime         string `url:"startTime,omitempty"`         //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If 'startTime' is not provided, API will default to current time minus 24 hours.
	EndTime           string `url:"endTime,omitempty"`           //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If 'endTime' is not provided, API will default to current time.
	MessageType       string `url:"messageType,omitempty"`       //Message type for the event. Examples: 'messageType=Syslog' (single messageType requested) 'messageType=Trap&messageType=Syslog' (multiple messageType requested)
	Severity          string `url:"severity,omitempty"`          //Severity of the event between 0 and 6. This is applicable only for events related to network devices (other than AP) and 'Wired Client' events. | Value | Severity    | | ----- | ----------- | | 0     | Emergency   | | 1     | Alert       | | 2     | Critical    | | 3     | Error       | | 4     | Warning     | | 5     | Notice      | | 6     | Info        | Examples: 'severity=0' (single severity requested) 'severity=0&severity=1' (multiple severity requested)
	SiteID            string `url:"siteId,omitempty"`            //The UUID of the site. (Ex. 'flooruuid') Examples: '?siteId=id1' (single siteId requested) '?siteId=id1&siteId=id2&siteId=id3' (multiple siteId requested)
	SiteHierarchyID   string `url:"siteHierarchyId,omitempty"`   //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. 'globalUuid/areaUuid/buildingUuid/floorUuid') This field supports wildcard asterisk ('*') character search support. E.g. '*uuid*, *uuid, uuid*' Examples: '?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid '(single siteHierarchyId requested) '?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2' (multiple siteHierarchyId requested)
	NetworkDeviceName string `url:"networkDeviceName,omitempty"` //Network device name. This parameter is applicable for network device related families. This field supports wildcard ('*') character-based search. Ex: '*Branch*' or 'Branch*' or '*Branch' Examples: 'networkDeviceName=Branch-3-Gateway' (single networkDeviceName requested) 'networkDeviceName=Branch-3-Gateway&networkDeviceName=Branch-3-Switch' (multiple networkDeviceName requested)
	NetworkDeviceID   string `url:"networkDeviceId,omitempty"`   //The list of Network Device Uuids. (Ex. '6bef213c-19ca-4170-8375-b694e251101c') Examples: 'networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c' (single networkDeviceId requested) 'networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c&networkDeviceId=32219612-819e-4b5e-a96b-cf22aca13dd9&networkDeviceId=2541e9a7-b80d-4955-8aa2-79b233318ba0' (multiple networkDeviceId requested)
	ApMac             string `url:"apMac,omitempty"`             //MAC address of the access point. This parameter is applicable for 'Unified AP' and 'Wireless Client' events. This field supports wildcard ('*') character-based search. Ex: '*50:0F*' or '50:0F*' or '*50:0F' Examples: 'apMac=50:0F:80:0F:F7:E0' (single apMac requested) 'apMac=50:0F:80:0F:F7:E0&apMac=18:80:90:AB:7E:A0' (multiple apMac requested)
	ClientMac         string `url:"clientMac,omitempty"`         //MAC address of the client. This parameter is applicable for 'Wired Client' and 'Wireless Client' events. This field supports wildcard ('*') character-based search. Ex: '*66:2B*' or '66:2B*' or '*66:2B' Examples: 'clientMac=66:2B:B8:D2:01:56' (single clientMac requested) 'clientMac=66:2B:B8:D2:01:56&clientMac=DC:A6:32:F5:5A:89' (multiple clientMac requested)
}
type CountTheNumberOfEventsHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type QueryAssuranceEventsWithFiltersHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type CountTheNumberOfEventsWithFiltersHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetDetailsOfASingleAssuranceEventQueryParams struct {
	Attribute string `url:"attribute,omitempty"` //The list of attributes that needs to be included in the response. If this parameter is not provided, then basic attributes ('id', 'name', 'timestamp', 'details', 'messageType', 'siteHierarchyId', 'siteHierarchy', 'deviceFamily', 'networkDeviceId', 'networkDeviceName', 'managementIpAddress') would be part of the response.  Examples:  'attribute=name' (single attribute requested) 'attribute=name&attribute=networkDeviceName' (multiple attribute requested)
	View      string `url:"view,omitempty"`      //The list of events views. Please refer to 'EventViews' for the supported list  Examples:  'view=network' (single view requested) 'view=network&view=ap' (multiple view requested)
}
type GetDetailsOfASingleAssuranceEventHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetListOfChildEventsForTheGivenWirelessClientEventHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetsInterfacesAlongWithStatisticsDataFromAllNetworkDevicesQueryParams struct {
	StartTime               float64 `url:"startTime,omitempty"`               //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If 'startTime' is not provided, API will default to current time.
	EndTime                 float64 `url:"endTime,omitempty"`                 //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	Limit                   float64 `url:"limit,omitempty"`                   //Maximum number of records to return
	Offset                  float64 `url:"offset,omitempty"`                  //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	SortBy                  string  `url:"sortBy,omitempty"`                  //A field within the response to sort by.
	Order                   string  `url:"order,omitempty"`                   //The sort order of the field ascending or descending.
	SiteHierarchy           string  `url:"siteHierarchy,omitempty"`           //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. 'Global/AreaName/BuildingName/FloorName') This field supports wildcard asterisk ('*') character search support. E.g. '*/San*, */San, /San*' Examples: '?siteHierarchy=Global/AreaName/BuildingName/FloorName' (single siteHierarchy requested) '?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2' (multiple siteHierarchies requested)
	SiteHierarchyID         string  `url:"siteHierarchyId,omitempty"`         //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. 'globalUuid/areaUuid/buildingUuid/floorUuid') This field supports wildcard asterisk ('*') character search support. E.g. '*uuid*, *uuid, uuid*' Examples: '?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid '(single siteHierarchyId requested) '?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2' (multiple siteHierarchyIds requested)
	SiteID                  string  `url:"siteId,omitempty"`                  //The UUID of the site. (Ex. 'flooruuid') Examples: '?siteId=id1' (single id requested) '?siteId=id1&siteId=id2&siteId=id3' (multiple ids requested)
	View                    string  `url:"view,omitempty"`                    //The specific summary view being requested. This is an optional parameter which can be passed to get one or more of the specific view associated fields. The default view is ''configuration''. ### Response data proviced by each view:   1. **configuration** [id,adminStatus,description,duplexConfig,duplexOper,interfaceIfIndex,interfaceType,ipv4Address,ipv6AddressList,isL3Interface,isWan,macAddress,mediaType,name,operStatus, portChannelId,portMode, portType,speed,timestamp,vlanId,networkDeviceId,networkDeviceIpAddress,networkDeviceMacAddress,siteName,siteHierarchy,siteHierarchyId]   2. **statistics** [id,name,rxDiscards,rxError,rxRate,rxUtilization,txDiscards,txError,txRate,txUtilization,networkDeviceId,networkDeviceIpAddress,networkDeviceMacAddress,siteName,siteHierarchy,siteHierarchyId]   3. **stackPort** [id,name,peerStackMember,peerStackPort,stackPortType,networkDeviceId,networkDeviceIpAddress,networkDeviceMacAddress,siteName,siteHierarchy,siteHierarchyId]   The default view is configuration, If need to access an additional view, simply include the view name in the query parameter. Examples: view=configuration (single view requested) view=configuration&view=statistic&stackPort (multiple views requested)
	Attribute               string  `url:"attribute,omitempty"`               //The following list of attributes can be provided in the attribute field [id,adminStatus, description,duplexConfig,duplexOper,interfaceIfIndex,interfaceType,ipv4Address,ipv6AddressList,isL3Interface,isWan,macAddress,mediaType,name,operStatus,peerStackMember,peerStackPort, portChannelId,portMode, portType,rxDiscards,rxError,rxRate,rxUtilization,speed,stackPortType,timestamp,txDiscards,txError,txRate,txUtilization,vlanId,networkDeviceId,networkDeviceIpAddress,networkDeviceMacAddress,siteName,siteHierarchy,siteHierarchyId] If length of attribute list is too long, please use 'views' param instead. Examples: attributes=name (single attribute requested) attributes=name,description,duplexOper (multiple attributes with comma separator)
	NetworkDeviceID         string  `url:"networkDeviceId,omitempty"`         //The list of Network Device Uuids. (Ex. '6bef213c-19ca-4170-8375-b694e251101c') Examples: 'networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c' (single networkDeviceId requested) 'networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c&networkDeviceId=32219612-819e-4b5e-a96b-cf22aca13dd9&networkDeviceId=2541e9a7-b80d-4955-8aa2-79b233318ba0' (multiple networkDeviceIds with & separator)
	NetworkDeviceIPAddress  string  `url:"networkDeviceIpAddress,omitempty"`  //The list of Network Device management IP Address. (Ex. '121.1.1.10') This field supports wildcard ('*') character-based search.  Ex: '*1.1*' or '1.1*' or '*1.1' Examples: 'networkDeviceIpAddress=121.1.1.10' 'networkDeviceIpAddress=121.1.1.10&networkDeviceIpAddress=172.20.1.10&networkDeviceIpAddress=10.10.20.10' (multiple networkDevice IP Address with & separator)
	NetworkDeviceMacAddress string  `url:"networkDeviceMacAddress,omitempty"` //The list of Network Device MAC Address. (Ex. '64:f6:9d:07:9a:00') This field supports wildcard ('*') character-based search.  Ex: '*AB:AB:AB*' or 'AB:AB:AB*' or '*AB:AB:AB' Examples: 'networkDeviceMacAddress=64:f6:9d:07:9a:00' 'networkDeviceMacAddress=64:f6:9d:07:9a:00&networkDeviceMacAddress=70:56:9d:07:ac:77' (multiple networkDevice MAC addresses with & separator)
	InterfaceID             string  `url:"interfaceId,omitempty"`             //The list of Interface Uuids. (Ex. '6bef213c-19ca-4170-8375-b694e251101c') Examples: 'interfaceId=6bef213c-19ca-4170-8375-b694e251101c' (single interface uuid ) 'interfaceId=6bef213c-19ca-4170-8375-b694e251101c&32219612-819e-4b5e-a96b-cf22aca13dd9&2541e9a7-b80d-4955-8aa2-79b233318ba0' (multiple Interface uuid with & separator)
	InterfaceName           string  `url:"interfaceName,omitempty"`           //The list of Interface name (Ex. 'GigabitEthernet1/0/1') This field supports wildcard ('*') character-based search.  Ex: '*1/0/1*' or '1/0/1*' or '*1/0/1' Examples: 'interfaceNames=GigabitEthernet1/0/1' (single interface name) 'interfaceNames=GigabitEthernet1/0/1&GigabitEthernet2/0/1&GigabitEthernet3/0/1' (multiple interface names with & separator)
}
type GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountQueryParams struct {
	StartTime               float64 `url:"startTime,omitempty"`               //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If 'startTime' is not provided, API will default to current time.
	EndTime                 float64 `url:"endTime,omitempty"`                 //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	SiteHierarchy           string  `url:"siteHierarchy,omitempty"`           //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. 'Global/AreaName/BuildingName/FloorName') This field supports wildcard asterisk ('*') character search support. E.g. '*/San*, */San, /San*' Examples: '?siteHierarchy=Global/AreaName/BuildingName/FloorName' (single siteHierarchy requested) '?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2' (multiple siteHierarchies requested)
	SiteHierarchyID         string  `url:"siteHierarchyId,omitempty"`         //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. 'globalUuid/areaUuid/buildingUuid/floorUuid') This field supports wildcard asterisk ('*') character search support. E.g. '*uuid*, *uuid, uuid*' Examples: '?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid '(single siteHierarchyId requested) '?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2' (multiple siteHierarchyIds requested)
	SiteID                  string  `url:"siteId,omitempty"`                  //The UUID of the site. (Ex. 'flooruuid') Examples: '?siteId=id1' (single id requested) '?siteId=id1&siteId=id2&siteId=id3' (multiple ids requested)
	NetworkDeviceID         string  `url:"networkDeviceId,omitempty"`         //The list of Network Device Uuids. (Ex. '6bef213c-19ca-4170-8375-b694e251101c') Examples: 'networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c' (single networkDeviceId requested) 'networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c&networkDeviceId=32219612-819e-4b5e-a96b-cf22aca13dd9&networkDeviceId=2541e9a7-b80d-4955-8aa2-79b233318ba0' (multiple networkDeviceIds with & separator)
	NetworkDeviceIPAddress  string  `url:"networkDeviceIpAddress,omitempty"`  //The list of Network Device management IP Address. (Ex. '121.1.1.10') This field supports wildcard ('*') character-based search.  Ex: '*1.1*' or '1.1*' or '*1.1' Examples: 'networkDeviceIpAddress=121.1.1.10' 'networkDeviceIpAddress=121.1.1.10&networkDeviceIpAddress=172.20.1.10&networkDeviceIpAddress=10.10.20.10' (multiple networkDevice IP Address with & separator)
	NetworkDeviceMacAddress string  `url:"networkDeviceMacAddress,omitempty"` //The list of Network Device MAC Address. (Ex. '64:f6:9d:07:9a:00') This field supports wildcard ('*') character-based search.  Ex: '*AB:AB:AB*' or 'AB:AB:AB*' or '*AB:AB:AB' Examples: 'networkDeviceMacAddress=64:f6:9d:07:9a:00' 'networkDeviceMacAddress=64:f6:9d:07:9a:00&networkDeviceMacAddress=70:56:9d:07:ac:77' (multiple networkDevice MAC addresses with & separator)
	InterfaceID             string  `url:"interfaceId,omitempty"`             //The list of Interface Uuids. (Ex. '6bef213c-19ca-4170-8375-b694e251101c') Examples: 'interfaceId=6bef213c-19ca-4170-8375-b694e251101c' (single interface uuid ) 'interfaceId=6bef213c-19ca-4170-8375-b694e251101c&32219612-819e-4b5e-a96b-cf22aca13dd9&2541e9a7-b80d-4955-8aa2-79b233318ba0' (multiple Interface uuid with & separator)
	InterfaceName           string  `url:"interfaceName,omitempty"`           //The list of Interface name (Ex. 'GigabitEthernet1/0/1') This field supports wildcard ('*') character-based search.  Ex: '*1/0/1*' or '1/0/1*' or '*1/0/1' Examples: 'interfaceNames=GigabitEthernet1/0/1' (single interface name) 'interfaceNames=GigabitEthernet1/0/1&GigabitEthernet2/0/1&GigabitEthernet3/0/1' (multiple interface names with & separator)
}
type GetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsDataQueryParams struct {
	StartTime float64 `url:"startTime,omitempty"` //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If 'startTime' is not provided, API will default to current time.
	EndTime   float64 `url:"endTime,omitempty"`   //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	View      string  `url:"view,omitempty"`      //Interface data model views
	Attribute string  `url:"attribute,omitempty"` //The following list of attributes can be provided in the attribute field [id,adminStatus, description,duplexConfig,duplexOper,interfaceIfIndex,interfaceType,ipv4Address,ipv6AddressList,isL3Interface,isWan,macAddress,mediaType,name,operStatus,peerStackMember,peerStackPort, portChannelId,portMode, portType,rxDiscards,rxError,rxRate,rxUtilization,speed,stackPortType,timestamp,txDiscards,txError,txRate,txUtilization,vlanId,networkDeviceId,networkDeviceIpAddress,networkDeviceMacAddress,siteName,siteHierarchy,siteHierarchyId] If length of attribute list is too long, please use 'views' param instead. Examples: attributes=name (single attribute requested) attributes=name,description,duplexOper (multiple attributes with comma separator)
}
type GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersQueryParams struct {
	StartTime           float64 `url:"startTime,omitempty"`           //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If 'startTime' is not provided, API will default to current time.
	EndTime             float64 `url:"endTime,omitempty"`             //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	Limit               float64 `url:"limit,omitempty"`               //Maximum number of records to return
	Offset              float64 `url:"offset,omitempty"`              //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	SortBy              string  `url:"sortBy,omitempty"`              //A field within the response to sort by.
	Order               string  `url:"order,omitempty"`               //The sort order of the field ascending or descending.
	SiteHierarchy       string  `url:"siteHierarchy,omitempty"`       //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. 'Global/AreaName/BuildingName/FloorName') This field supports wildcard asterisk (*) character search support. E.g. */San*, */San, /San* Examples: '?siteHierarchy=Global/AreaName/BuildingName/FloorName' (single siteHierarchy requested) '?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2' (multiple siteHierarchies requested)
	SiteHierarchyID     string  `url:"siteHierarchyId,omitempty"`     //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. 'globalUuid/areaUuid/buildingUuid/floorUuid') This field supports wildcard asterisk (*) character search support. E.g. '*uuid*, *uuid, uuid* Examples: '?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid '(single siteHierarchyId requested) '?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2' (multiple siteHierarchyIds requested)
	SiteID              string  `url:"siteId,omitempty"`              //The UUID of the site. (Ex. 'flooruuid') This field supports wildcard asterisk (*) character search support. E.g.*flooruuid*, *flooruuid, flooruuid* Examples: '?siteId=id1' (single id requested) '?siteId=id1&siteId=id2&siteId=id3' (multiple ids requested)
	ID                  string  `url:"id,omitempty"`                  //The list of entity Uuids. (Ex."6bef213c-19ca-4170-8375-b694e251101c") Examples: id=6bef213c-19ca-4170-8375-b694e251101c (single entity uuid requested) id=6bef213c-19ca-4170-8375-b694e251101c&id=32219612-819e-4b5e-a96b-cf22aca13dd9&id=2541e9a7-b80d-4955-8aa2-79b233318ba0 (multiple entity uuid with '&' separator)
	ManagementIPAddress string  `url:"managementIpAddress,omitempty"` //The list of entity management IP Address. It can be either Ipv4 or Ipv6 address or combination of both(Ex. "121.1.1.10") This field supports wildcard ('*') character-based search.  Ex: '*1.1*' or '1.1*' or '*1.1' Examples: managementIpAddresses=121.1.1.10 managementIpAddresses=121.1.1.10&managementIpAddresses=172.20.1.10&managementIpAddresses=200:10&=managementIpAddresses172.20.3.4 (multiple entity IP Address with & separator)
	MacAddress          string  `url:"macAddress,omitempty"`          //The macAddress of the network device or client This field supports wildcard ('*') character-based search.  Ex: '*AB:AB:AB*' or 'AB:AB:AB*' or '*AB:AB:AB' Examples: 'macAddress=AB:AB:AB:CD:CD:CD' (single macAddress requested) 'macAddress=AB:AB:AB:CD:CD:DC&macAddress=AB:AB:AB:CD:CD:FE' (multiple macAddress requested)
	Family              string  `url:"family,omitempty"`              //The list of network device family names Examples:family=Switches and Hubs (single network device family name )family=Switches and Hubs&family=Router&family=Wireless Controller (multiple Network device family names with & separator). This field is not case sensitive.
	Type                string  `url:"type,omitempty"`                //The list of network device type This field supports wildcard ('*') character-based search. Ex: '*9407R*' or '*9407R' or '9407R*' Examples: type=SwitchesCisco Catalyst 9407R Switch (single network device types ) type=Cisco Catalyst 38xx stack-able ethernet switch&type=Cisco 3945 Integrated Services Router G2 (multiple Network device types with & separator)
	Role                string  `url:"role,omitempty"`                //The list of network device role. Examples:role=CORE, role=CORE&role=ACCESS&role=ROUTER (multiple Network device roles with & separator). This field is not case sensitive.
	SerialNumber        string  `url:"serialNumber,omitempty"`        //The list of network device serial numbers. This field supports wildcard ('*') character-based search.  Ex: '*MS1SV*' or 'MS1SV*' or '*MS1SV' Examples: serialNumber=9FUFMS1SVAX serialNumber=9FUFMS1SVAX&FCW2333Q0BY&FJC240617JX(multiple Network device serial number with & separator)
	MaintenanceMode     bool    `url:"maintenanceMode,omitempty"`     //The device maintenanceMode status true or false
	SoftwareVersion     string  `url:"softwareVersion,omitempty"`     //The list of network device software version This field supports wildcard ('*') character-based search. Ex: '*17.8*' or '*17.8' or '17.8*' Examples: softwareVersion=2.3.4.0 (single network device software version ) softwareVersion=17.9.3.23&softwareVersion=17.7.1.2&softwareVersion=*.17.7 (multiple Network device software versions with & separator)
	HealthScore         string  `url:"healthScore,omitempty"`         //The list of entity health score categories Examples: healthScore=good, healthScore=good&healthScore=fair (multiple entity healthscore values with & separator). This field is not case sensitive.
	View                string  `url:"view,omitempty"`                //The List of Network Device model views. Please refer to '''NetworkDeviceView''' for the supported list
	Attribute           string  `url:"attribute,omitempty"`           //The List of Network Device model attributes. This is helps to specify the interested fields in the request.
}
type GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersQueryParams struct {
	StartTime           float64 `url:"startTime,omitempty"`           //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If 'startTime' is not provided, API will default to current time.
	EndTime             float64 `url:"endTime,omitempty"`             //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	ID                  string  `url:"id,omitempty"`                  //The list of entity Uuids. (Ex."6bef213c-19ca-4170-8375-b694e251101c") Examples: id=6bef213c-19ca-4170-8375-b694e251101c (single entity uuid requested) id=6bef213c-19ca-4170-8375-b694e251101c&id=32219612-819e-4b5e-a96b-cf22aca13dd9&id=2541e9a7-b80d-4955-8aa2-79b233318ba0 (multiple entity uuid with '&' separator)
	SiteHierarchy       string  `url:"siteHierarchy,omitempty"`       //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. 'Global/AreaName/BuildingName/FloorName') This field supports wildcard asterisk (*) character search support. E.g. */San*, */San, /San* Examples: '?siteHierarchy=Global/AreaName/BuildingName/FloorName' (single siteHierarchy requested) '?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2' (multiple siteHierarchies requested)
	SiteHierarchyID     string  `url:"siteHierarchyId,omitempty"`     //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. 'globalUuid/areaUuid/buildingUuid/floorUuid') This field supports wildcard asterisk (*) character search support. E.g. '*uuid*, *uuid, uuid* Examples: '?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid '(single siteHierarchyId requested) '?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2' (multiple siteHierarchyIds requested)
	SiteID              string  `url:"siteId,omitempty"`              //The UUID of the site. (Ex. 'flooruuid') This field supports wildcard asterisk (*) character search support. E.g.*flooruuid*, *flooruuid, flooruuid* Examples: '?siteId=id1' (single id requested) '?siteId=id1&siteId=id2&siteId=id3' (multiple ids requested)
	ManagementIPAddress string  `url:"managementIpAddress,omitempty"` //The list of entity management IP Address. It can be either Ipv4 or Ipv6 address or combination of both(Ex. "121.1.1.10") This field supports wildcard ('*') character-based search.  Ex: '*1.1*' or '1.1*' or '*1.1' Examples: managementIpAddresses=121.1.1.10 managementIpAddresses=121.1.1.10&managementIpAddresses=172.20.1.10&managementIpAddresses=200:10&=managementIpAddresses172.20.3.4 (multiple entity IP Address with & separator)
	MacAddress          string  `url:"macAddress,omitempty"`          //The macAddress of the network device or client This field supports wildcard ('*') character-based search.  Ex: '*AB:AB:AB*' or 'AB:AB:AB*' or '*AB:AB:AB' Examples: 'macAddress=AB:AB:AB:CD:CD:CD' (single macAddress requested) 'macAddress=AB:AB:AB:CD:CD:DC&macAddress=AB:AB:AB:CD:CD:FE' (multiple macAddress requested)
	Family              string  `url:"family,omitempty"`              //The list of network device family names Examples:family=Switches and Hubs (single network device family name )family=Switches and Hubs&family=Router&family=Wireless Controller (multiple Network device family names with & separator). This field is not case sensitive.
	Type                string  `url:"type,omitempty"`                //The list of network device type This field supports wildcard ('*') character-based search. Ex: '*9407R*' or '*9407R' or '9407R*'Examples:type=SwitchesCisco Catalyst 9407R Switch (single network device types )type=Cisco Catalyst 38xx stack-able ethernet switch&type=Cisco 3945 Integrated Services Router G2 (multiple Network device types with & separator)
	Role                string  `url:"role,omitempty"`                //The list of network device role. Examples:role=CORE, role=CORE&role=ACCESS&role=ROUTER (multiple Network device roles with & separator). This field is not case sensitive.
	SerialNumber        string  `url:"serialNumber,omitempty"`        //The list of network device serial numbers. This field supports wildcard ('*') character-based search.  Ex: '*MS1SV*' or 'MS1SV*' or '*MS1SV' Examples: serialNumber=9FUFMS1SVAX serialNumber=9FUFMS1SVAX&FCW2333Q0BY&FJC240617JX(multiple Network device serial number with & separator)
	MaintenanceMode     bool    `url:"maintenanceMode,omitempty"`     //The device maintenanceMode status true or false
	SoftwareVersion     string  `url:"softwareVersion,omitempty"`     //The list of network device software version This field supports wildcard ('*') character-based search. Ex: '*17.8*' or '*17.8' or '17.8*' Examples: softwareVersion=2.3.4.0 (single network device software version ) softwareVersion=17.9.3.23&softwareVersion=17.7.1.2&softwareVersion=*.17.7 (multiple Network device software versions with & separator)
	HealthScore         string  `url:"healthScore,omitempty"`         //The list of entity health score categories Examples:healthScore=good,healthScore=good&healthScore=fair (multiple entity healthscore values with & separator). This field is not case sensitive.
	View                string  `url:"view,omitempty"`                //The List of Network Device model views. Please refer to '''NetworkDeviceView''' for the supported list
	Attribute           string  `url:"attribute,omitempty"`           //The List of Network Device model attributes. This is helps to specify the interested fields in the request.
}
type GetTheDeviceDataForTheGivenDeviceIDUUIDQueryParams struct {
	StartTime float64 `url:"startTime,omitempty"` //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If 'startTime' is not provided, API will default to current time.
	EndTime   float64 `url:"endTime,omitempty"`   //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	View      string  `url:"view,omitempty"`      //The List of Network Device model views. Please refer to '''NetworkDeviceView''' for the supported list
	Attribute string  `url:"attribute,omitempty"` //The List of Network Device model attributes. This is helps to specify the interested fields in the request.
}
type GetPlannedAccessPointsForBuildingQueryParams struct {
	Limit  float64 `url:"limit,omitempty"`  //The page size limit for the response, e.g. limit=100 will return a maximum of 100 records
	Offset float64 `url:"offset,omitempty"` //The page offset for the response. E.g. if limit=100, offset=0 will return first 100 records, offset=1 will return next 100 records, etc.
	Radios bool    `url:"radios,omitempty"` //Whether to include the planned radio details of the planned access points
}
type GetDeviceDetailQueryParams struct {
	Timestamp  float64 `url:"timestamp,omitempty"`  //UTC timestamp of device data in milliseconds
	IDentifier string  `url:"identifier,omitempty"` //One of "macAddress", "nwDeviceName", "uuid" (case insensitive)
	SearchBy   string  `url:"searchBy,omitempty"`   //MAC Address, device name, or UUID of the network device
}
type GetDeviceEnrichmentDetailsHeaderParams struct {
	EntityType        string `url:"entity_type,omitempty"`         //Expects type string. Device enrichment details can be fetched based on either Device ID or Device MAC address or Device IP Address. This parameter value must either be device_id/mac_address/ip_address
	EntityValue       string `url:"entity_value,omitempty"`        //Expects type string. Contains the actual value for the entity type that has been defined
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type bool.
}
type DevicesQueryParams struct {
	DeviceRole string  `url:"deviceRole,omitempty"` //CORE, ACCESS, DISTRIBUTION, ROUTER, WLC, or AP (case insensitive)
	SiteID     string  `url:"siteId,omitempty"`     //CATALYST site UUID
	Health     string  `url:"health,omitempty"`     //CATALYST health catagory: POOR, FAIR, or GOOD (case insensitive)
	StartTime  float64 `url:"startTime,omitempty"`  //UTC epoch time in milliseconds
	EndTime    float64 `url:"endTime,omitempty"`    //UTC epoch time in milliseconds
	Limit      float64 `url:"limit,omitempty"`      //Max number of device entries in the response (default to 50. Max at 500)
	Offset     float64 `url:"offset,omitempty"`     //The offset of the first device in the returned data (Mutiple of 'limit' + 1)
}
type GetPlannedAccessPointsForFloorQueryParams struct {
	Limit  float64 `url:"limit,omitempty"`  //The page size limit for the response, e.g. limit=100 will return a maximum of 100 records
	Offset float64 `url:"offset,omitempty"` //The page offset for the response. E.g. if limit=100, offset=0 will return first 100 records, offset=1 will return next 100 records, etc.
	Radios bool    `url:"radios,omitempty"` //Whether to include the planned radio details of the planned access points
}
type GetAllHealthScoreDefinitionsForGivenFiltersQueryParams struct {
	DeviceType              string  `url:"deviceType,omitempty"`              //These are the device families supported for health score definitions. If no input is made on device family, all device families are considered.
	ID                      string  `url:"id,omitempty"`                      //The definition identifier. Examples: id=015d9cba-4f53-4087-8317-7e49e5ffef46 (single entity id request) id=015d9cba-4f53-4087-8317-7e49e5ffef46&id=015d9cba-4f53-4087-8317-7e49e5ffef47 (multiple ids in the query param)
	IncludeForOverallHealth bool    `url:"includeForOverallHealth,omitempty"` //The inclusion status of the issue definition, either true or false. true indicates that particular health metric is included in overall health computation, otherwise false. By default it's set to true.
	Attribute               string  `url:"attribute,omitempty"`               //These are the attributes supported in health score definitions response. By default, all properties are sent in response.
	Offset                  float64 `url:"offset,omitempty"`                  //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	Limit                   float64 `url:"limit,omitempty"`                   //Maximum number of records to return
}
type GetAllHealthScoreDefinitionsForGivenFiltersHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type UpdateHealthScoreDefinitionsHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetHealthScoreDefinitionForTheGivenIDHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetAllInterfacesQueryParams struct {
	Offset         int    `url:"offset,omitempty"`         //Offset
	Limit          int    `url:"limit,omitempty"`          //Limit
	LastInputTime  string `url:"lastInputTime,omitempty"`  //Last Input Time
	LastOutputTime string `url:"lastOutputTime,omitempty"` //Last Output Time
}
type GetInterfaceDetailsByDeviceIDAndInterfaceNameQueryParams struct {
	Name string `url:"name,omitempty"` //Interface name
}
type UpdateInterfaceDetailsQueryParams struct {
	DeploymentMode string `url:"deploymentMode,omitempty"` //Preview/Deploy ['Preview' means the configuration is not pushed to the device. 'Deploy' makes the configuration pushed to the device]
}
type ClearMacAddressTableQueryParams struct {
	DeploymentMode string `url:"deploymentMode,omitempty"` //Preview/Deploy ['Preview' means the configuration is not pushed to the device. 'Deploy' makes the configuration pushed to the device]
}
type GetDeviceListQueryParams struct {
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
	Limit                     int      `url:"limit,omitempty"`                      //1 <= limit <= 500 [max. no. of devices to be returned in the result]
}
type GetDeviceValuesThatMatchFullyOrPartiallyAnAttributeQueryParams struct {
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
	Limit                     int    `url:"limit,omitempty"`                     //limit
}
type GetDeviceCount2QueryParams struct {
	Hostname            []string `url:"hostname,omitempty"`            //hostname
	ManagementIPAddress []string `url:"managementIpAddress,omitempty"` //managementIpAddress
	MacAddress          []string `url:"macAddress,omitempty"`          //macAddress
	LocationName        []string `url:"locationName,omitempty"`        //locationName
}
type GetFunctionalCapabilityForDevicesQueryParams struct {
	DeviceID     string   `url:"deviceId,omitempty"`     //Accepts comma separated deviceid's and return list of functional-capabilities for the given id's. If invalid or not-found id's are provided, null entry will be returned in the list.
	FunctionName []string `url:"functionName,omitempty"` //functionName
}
type InventoryInsightDeviceLinkMismatchAPIQueryParams struct {
	Offset   int    `url:"offset,omitempty"`   //Row Number.  Default value is 1
	Limit    int    `url:"limit,omitempty"`    //Default value is 500
	Category string `url:"category,omitempty"` //Links mismatch category.  Value can be speed-duplex or vlan.
	SortBy   string `url:"sortBy,omitempty"`   //Sort By
	Order    string `url:"order,omitempty"`    //Order.  Value can be asc or desc.  Default value is asc
}
type GetModulesQueryParams struct {
	DeviceID                 string   `url:"deviceId,omitempty"`                 //deviceId
	Limit                    int      `url:"limit,omitempty"`                    //limit
	Offset                   int      `url:"offset,omitempty"`                   //offset
	NameList                 []string `url:"nameList,omitempty"`                 //nameList
	VendorEquipmentTypeList  []string `url:"vendorEquipmentTypeList,omitempty"`  //vendorEquipmentTypeList
	PartNumberList           []string `url:"partNumberList,omitempty"`           //partNumberList
	OperationalStateCodeList []string `url:"operationalStateCodeList,omitempty"` //operationalStateCodeList
}
type GetModuleCountQueryParams struct {
	DeviceID                 string   `url:"deviceId,omitempty"`                 //deviceId
	NameList                 []string `url:"nameList,omitempty"`                 //nameList
	VendorEquipmentTypeList  []string `url:"vendorEquipmentTypeList,omitempty"`  //vendorEquipmentTypeList
	PartNumberList           []string `url:"partNumberList,omitempty"`           //partNumberList
	OperationalStateCodeList []string `url:"operationalStateCodeList,omitempty"` //operationalStateCodeList
}
type SyncDevicesQueryParams struct {
	ForceSync bool `url:"forceSync,omitempty"` //forceSync
}
type GetDevicesRegisteredForWsaNotificationQueryParams struct {
	SerialNumber string `url:"serialNumber,omitempty"` //Serial number of the device
	Macaddress   string `url:"macaddress,omitempty"`   //Mac addres of the device
}
type GetAllUserDefinedFieldsQueryParams struct {
	ID   string `url:"id,omitempty"`   //Comma-seperated id(s) used for search/filtering
	Name string `url:"name,omitempty"` //Comma-seperated name(s) used for search/filtering
}
type RemoveUserDefinedFieldFromDeviceQueryParams struct {
	Name string `url:"name,omitempty"` //Name of UDF to be removed
}
type GetTheDetailsOfPhysicalComponentsOfTheGivenDeviceQueryParams struct {
	Type string `url:"type,omitempty"` //Type value can be PowerSupply, Fan, Chassis, Backplane, Module, PROCESSOR, Other, SFP. If no type is mentioned, All equipments are fetched for the device.
}
type ReturnsPoeInterfaceDetailsForTheDeviceQueryParams struct {
	InterfaceNameList string `url:"interfaceNameList,omitempty"` //comma seperated interface names
}
type DeleteDeviceByIDQueryParams struct {
	CleanConfig bool `url:"cleanConfig,omitempty"` //Selecting the clean up configuration option will attempt to remove device settings that were configured during the addition of the device to the inventory and site assignment. Please note that this operation is different from deprovisioning. It does not remove configurations that were pushed during device provisioning.
}
type GetDeviceInterfaceVLANsQueryParams struct {
	InterfaceType string `url:"interfaceType,omitempty"` //Vlan associated with sub-interface. If no interfaceType mentioned it will return all types of Vlan interfaces. If interfaceType is selected but not specified then it will take default value.
}
type GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersQueryParams struct {
	DeviceType              string `url:"deviceType,omitempty"`              //These are the device families supported for health score definitions. If no input is made on device family, all device families are considered.
	ID                      string `url:"id,omitempty"`                      //The definition identifier. Examples: id=015d9cba-4f53-4087-8317-7e49e5ffef46 (single entity id request) id=015d9cba-4f53-4087-8317-7e49e5ffef46&id=015d9cba-4f53-4087-8317-7e49e5ffef47 (multiple ids in the query param)
	IncludeForOverallHealth bool   `url:"includeForOverallHealth,omitempty"` //The inclusion status of the issue definition, either true or false. true indicates that particular health metric is included in overall health computation, otherwise false. By default it's set to true.
}
type GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}

type ResponseDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions struct {
	Response *ResponseDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponse `json:"response,omitempty"` //
	Version  string                                                                                                          `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseDevicesQueryAssuranceEvents struct {
	Response *[]ResponseDevicesQueryAssuranceEventsResponse `json:"response,omitempty"` //
	Version  string                                         `json:"version,omitempty"`  // Version
	Page     *ResponseDevicesQueryAssuranceEventsPage       `json:"page,omitempty"`     //
}
type ResponseDevicesQueryAssuranceEventsResponse struct {
	OldRadioChannelWidth         string                                                           `json:"oldRadioChannelWidth,omitempty"`         // Old Radio Channel Width
	ClientMac                    string                                                           `json:"clientMac,omitempty"`                    // Client Mac
	SwitchNumber                 string                                                           `json:"switchNumber,omitempty"`                 // Switch Number
	AssocRssi                    *int                                                             `json:"assocRssi,omitempty"`                    // Assoc Rssi
	AffectedClients              []string                                                         `json:"affectedClients,omitempty"`              // Affected Clients
	IsPrivateMac                 *bool                                                            `json:"isPrivateMac,omitempty"`                 // Is Private Mac
	Frequency                    string                                                           `json:"frequency,omitempty"`                    // Frequency
	ApRole                       string                                                           `json:"apRole,omitempty"`                       // Ap Role
	ReplacingDeviceSerialNumber  string                                                           `json:"replacingDeviceSerialNumber,omitempty"`  // Replacing Device Serial Number
	MessageType                  string                                                           `json:"messageType,omitempty"`                  // Message Type
	FailureCategory              string                                                           `json:"failureCategory,omitempty"`              // Failure Category
	ApSwitchName                 string                                                           `json:"apSwitchName,omitempty"`                 // Ap Switch Name
	ApSwitchID                   string                                                           `json:"apSwitchId,omitempty"`                   // Ap Switch Id
	RadioChannelUtilization      string                                                           `json:"radioChannelUtilization,omitempty"`      // Radio Channel Utilization
	Mnemonic                     string                                                           `json:"mnemonic,omitempty"`                     // Mnemonic
	RadioChannelSlot             *int                                                             `json:"radioChannelSlot,omitempty"`             // Radio Channel Slot
	Details                      string                                                           `json:"details,omitempty"`                      // Details
	ID                           string                                                           `json:"id,omitempty"`                           // Id
	LastApDisconnectReason       string                                                           `json:"lastApDisconnectReason,omitempty"`       // Last Ap Disconnect Reason
	NetworkDeviceName            string                                                           `json:"networkDeviceName,omitempty"`            // Network Device Name
	IDentifier                   string                                                           `json:"identifier,omitempty"`                   // Identifier
	ReasonDescription            string                                                           `json:"reasonDescription,omitempty"`            // Reason Description
	VLANID                       string                                                           `json:"vlanId,omitempty"`                       // Vlan Id
	UdnID                        string                                                           `json:"udnId,omitempty"`                        // Udn Id
	AuditSessionID               string                                                           `json:"auditSessionId,omitempty"`               // Audit Session Id
	ApMac                        string                                                           `json:"apMac,omitempty"`                        // Ap Mac
	DeviceFamily                 string                                                           `json:"deviceFamily,omitempty"`                 // Device Family
	RadioNoise                   string                                                           `json:"radioNoise,omitempty"`                   // Radio Noise
	WlcName                      string                                                           `json:"wlcName,omitempty"`                      // Wlc Name
	ApRadioOperationState        string                                                           `json:"apRadioOperationState,omitempty"`        // Ap Radio Operation State
	Name                         string                                                           `json:"name,omitempty"`                         // Name
	FailureIPAddress             string                                                           `json:"failureIpAddress,omitempty"`             // Failure Ip Address
	NewRadioChannelList          string                                                           `json:"newRadioChannelList,omitempty"`          // New Radio Channel List
	Duid                         string                                                           `json:"duid,omitempty"`                         // Duid
	RoamType                     string                                                           `json:"roamType,omitempty"`                     // Roam Type
	CandidateAPs                 *[]ResponseDevicesQueryAssuranceEventsResponseCandidateAPs       `json:"candidateAPs,omitempty"`                 //
	ReplacedDeviceSerialNumber   string                                                           `json:"replacedDeviceSerialNumber,omitempty"`   // Replaced Device Serial Number
	OldRadioChannelList          string                                                           `json:"oldRadioChannelList,omitempty"`          // Old Radio Channel List
	SSID                         string                                                           `json:"ssid,omitempty"`                         // Ssid
	SubReasonDescription         string                                                           `json:"subReasonDescription,omitempty"`         // Sub Reason Description
	WirelessClientEventEndTime   *int                                                             `json:"wirelessClientEventEndTime,omitempty"`   // Wireless Client Event End Time
	IPv4                         string                                                           `json:"ipv4,omitempty"`                         // Ipv4
	WlcID                        string                                                           `json:"wlcId,omitempty"`                        // Wlc Id
	IPv6                         string                                                           `json:"ipv6,omitempty"`                         // Ipv6
	MissingResponseAPs           *[]ResponseDevicesQueryAssuranceEventsResponseMissingResponseAPs `json:"missingResponseAPs,omitempty"`           //
	Timestamp                    *int                                                             `json:"timestamp,omitempty"`                    // Timestamp
	Severity                     *int                                                             `json:"severity,omitempty"`                     // Severity
	CurrentRadioPowerLevel       *int                                                             `json:"currentRadioPowerLevel,omitempty"`       // Current Radio Power Level
	NewRadioChannelWidth         string                                                           `json:"newRadioChannelWidth,omitempty"`         // New Radio Channel Width
	AssocSnr                     *int                                                             `json:"assocSnr,omitempty"`                     // Assoc Snr
	AuthServerIP                 string                                                           `json:"authServerIp,omitempty"`                 // Auth Server Ip
	ChildEvents                  *[]ResponseDevicesQueryAssuranceEventsResponseChildEvents        `json:"childEvents,omitempty"`                  //
	ConnectedInterfaceName       string                                                           `json:"connectedInterfaceName,omitempty"`       // Connected Interface Name
	DhcpServerIP                 string                                                           `json:"dhcpServerIp,omitempty"`                 // Dhcp Server Ip
	ManagementIPAddress          string                                                           `json:"managementIpAddress,omitempty"`          // Management Ip Address
	PreviousRadioPowerLevel      *int                                                             `json:"previousRadioPowerLevel,omitempty"`      // Previous Radio Power Level
	ResultStatus                 string                                                           `json:"resultStatus,omitempty"`                 // Result Status
	RadioInterference            string                                                           `json:"radioInterference,omitempty"`            // Radio Interference
	NetworkDeviceID              string                                                           `json:"networkDeviceId,omitempty"`              // Network Device Id
	SiteHierarchy                string                                                           `json:"siteHierarchy,omitempty"`                // Site Hierarchy
	EventStatus                  string                                                           `json:"eventStatus,omitempty"`                  // Event Status
	WirelessClientEventStartTime *int                                                             `json:"wirelessClientEventStartTime,omitempty"` // Wireless Client Event Start Time
	SiteHierarchyID              string                                                           `json:"siteHierarchyId,omitempty"`              // Site Hierarchy Id
	UdnName                      string                                                           `json:"udnName,omitempty"`                      // Udn Name
	Facility                     string                                                           `json:"facility,omitempty"`                     // Facility
	LastApResetType              string                                                           `json:"lastApResetType,omitempty"`              // Last Ap Reset Type
	InvalidIeAPs                 *[]ResponseDevicesQueryAssuranceEventsResponseInvalidIeAPs       `json:"invalidIeAPs,omitempty"`                 //
	Username                     string                                                           `json:"username,omitempty"`                     // Username
}
type ResponseDevicesQueryAssuranceEventsResponseCandidateAPs struct {
	APID   string `json:"apId,omitempty"`   // Ap Id
	ApName string `json:"apName,omitempty"` // Ap Name
	ApMac  string `json:"apMac,omitempty"`  // Ap Mac
	Bssid  string `json:"bssid,omitempty"`  // Bssid
	Rssi   *int   `json:"rssi,omitempty"`   // Rssi
}
type ResponseDevicesQueryAssuranceEventsResponseMissingResponseAPs struct {
	APID      string `json:"apId,omitempty"`      // Ap Id
	ApName    string `json:"apName,omitempty"`    // Ap Name
	ApMac     string `json:"apMac,omitempty"`     // Ap Mac
	Bssid     string `json:"bssid,omitempty"`     // Bssid
	Type      string `json:"type,omitempty"`      // Type
	FrameType string `json:"frameType,omitempty"` // Frame Type
}
type ResponseDevicesQueryAssuranceEventsResponseChildEvents struct {
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
type ResponseDevicesQueryAssuranceEventsResponseInvalidIeAPs struct {
	APID      string `json:"apId,omitempty"`      // Ap Id
	ApName    string `json:"apName,omitempty"`    // Ap Name
	ApMac     string `json:"apMac,omitempty"`     // Ap Mac
	Bssid     string `json:"bssid,omitempty"`     // Bssid
	Type      string `json:"type,omitempty"`      // Type
	FrameType string `json:"frameType,omitempty"` // Frame Type
	Ies       string `json:"ies,omitempty"`       // Ies
}
type ResponseDevicesQueryAssuranceEventsPage struct {
	Limit  *int                                             `json:"limit,omitempty"`  // Limit
	Offset *int                                             `json:"offset,omitempty"` // Offset
	Count  *int                                             `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesQueryAssuranceEventsPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesQueryAssuranceEventsPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseDevicesCountTheNumberOfEvents struct {
	Response *ResponseDevicesCountTheNumberOfEventsResponse `json:"response,omitempty"` //
	Version  string                                         `json:"version,omitempty"`  // Version
}
type ResponseDevicesCountTheNumberOfEventsResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseDevicesQueryAssuranceEventsWithFilters struct {
	Response *[]ResponseDevicesQueryAssuranceEventsWithFiltersResponse `json:"response,omitempty"` //
	Version  string                                                    `json:"version,omitempty"`  // Version
	Page     *ResponseDevicesQueryAssuranceEventsWithFiltersPage       `json:"page,omitempty"`     //
}
type ResponseDevicesQueryAssuranceEventsWithFiltersResponse struct {
	OldRadioChannelWidth         string                                                                      `json:"oldRadioChannelWidth,omitempty"`         // Old Radio Channel Width
	ClientMac                    string                                                                      `json:"clientMac,omitempty"`                    // Client Mac
	SwitchNumber                 string                                                                      `json:"switchNumber,omitempty"`                 // Switch Number
	AssocRssi                    *int                                                                        `json:"assocRssi,omitempty"`                    // Assoc Rssi
	AffectedClients              []string                                                                    `json:"affectedClients,omitempty"`              // Affected Clients
	IsPrivateMac                 *bool                                                                       `json:"isPrivateMac,omitempty"`                 // Is Private Mac
	Frequency                    string                                                                      `json:"frequency,omitempty"`                    // Frequency
	ApRole                       string                                                                      `json:"apRole,omitempty"`                       // Ap Role
	ReplacingDeviceSerialNumber  string                                                                      `json:"replacingDeviceSerialNumber,omitempty"`  // Replacing Device Serial Number
	MessageType                  string                                                                      `json:"messageType,omitempty"`                  // Message Type
	FailureCategory              string                                                                      `json:"failureCategory,omitempty"`              // Failure Category
	ApSwitchName                 string                                                                      `json:"apSwitchName,omitempty"`                 // Ap Switch Name
	ApSwitchID                   string                                                                      `json:"apSwitchId,omitempty"`                   // Ap Switch Id
	RadioChannelUtilization      string                                                                      `json:"radioChannelUtilization,omitempty"`      // Radio Channel Utilization
	Mnemonic                     string                                                                      `json:"mnemonic,omitempty"`                     // Mnemonic
	RadioChannelSlot             *int                                                                        `json:"radioChannelSlot,omitempty"`             // Radio Channel Slot
	Details                      string                                                                      `json:"details,omitempty"`                      // Details
	ID                           string                                                                      `json:"id,omitempty"`                           // Id
	LastApDisconnectReason       string                                                                      `json:"lastApDisconnectReason,omitempty"`       // Last Ap Disconnect Reason
	NetworkDeviceName            string                                                                      `json:"networkDeviceName,omitempty"`            // Network Device Name
	IDentifier                   string                                                                      `json:"identifier,omitempty"`                   // Identifier
	ReasonDescription            string                                                                      `json:"reasonDescription,omitempty"`            // Reason Description
	VLANID                       string                                                                      `json:"vlanId,omitempty"`                       // Vlan Id
	UdnID                        string                                                                      `json:"udnId,omitempty"`                        // Udn Id
	AuditSessionID               string                                                                      `json:"auditSessionId,omitempty"`               // Audit Session Id
	ApMac                        string                                                                      `json:"apMac,omitempty"`                        // Ap Mac
	DeviceFamily                 string                                                                      `json:"deviceFamily,omitempty"`                 // Device Family
	RadioNoise                   string                                                                      `json:"radioNoise,omitempty"`                   // Radio Noise
	WlcName                      string                                                                      `json:"wlcName,omitempty"`                      // Wlc Name
	ApRadioOperationState        string                                                                      `json:"apRadioOperationState,omitempty"`        // Ap Radio Operation State
	Name                         string                                                                      `json:"name,omitempty"`                         // Name
	FailureIPAddress             string                                                                      `json:"failureIpAddress,omitempty"`             // Failure Ip Address
	NewRadioChannelList          string                                                                      `json:"newRadioChannelList,omitempty"`          // New Radio Channel List
	Duid                         string                                                                      `json:"duid,omitempty"`                         // Duid
	RoamType                     string                                                                      `json:"roamType,omitempty"`                     // Roam Type
	CandidateAPs                 *[]ResponseDevicesQueryAssuranceEventsWithFiltersResponseCandidateAPs       `json:"candidateAPs,omitempty"`                 //
	ReplacedDeviceSerialNumber   string                                                                      `json:"replacedDeviceSerialNumber,omitempty"`   // Replaced Device Serial Number
	OldRadioChannelList          string                                                                      `json:"oldRadioChannelList,omitempty"`          // Old Radio Channel List
	SSID                         string                                                                      `json:"ssid,omitempty"`                         // Ssid
	SubReasonDescription         string                                                                      `json:"subReasonDescription,omitempty"`         // Sub Reason Description
	WirelessClientEventEndTime   *int                                                                        `json:"wirelessClientEventEndTime,omitempty"`   // Wireless Client Event End Time
	IPv4                         string                                                                      `json:"ipv4,omitempty"`                         // Ipv4
	WlcID                        string                                                                      `json:"wlcId,omitempty"`                        // Wlc Id
	IPv6                         string                                                                      `json:"ipv6,omitempty"`                         // Ipv6
	MissingResponseAPs           *[]ResponseDevicesQueryAssuranceEventsWithFiltersResponseMissingResponseAPs `json:"missingResponseAPs,omitempty"`           //
	Timestamp                    *int                                                                        `json:"timestamp,omitempty"`                    // Timestamp
	Severity                     *int                                                                        `json:"severity,omitempty"`                     // Severity
	CurrentRadioPowerLevel       *int                                                                        `json:"currentRadioPowerLevel,omitempty"`       // Current Radio Power Level
	NewRadioChannelWidth         string                                                                      `json:"newRadioChannelWidth,omitempty"`         // New Radio Channel Width
	AssocSnr                     *int                                                                        `json:"assocSnr,omitempty"`                     // Assoc Snr
	AuthServerIP                 string                                                                      `json:"authServerIp,omitempty"`                 // Auth Server Ip
	ChildEvents                  *[]ResponseDevicesQueryAssuranceEventsWithFiltersResponseChildEvents        `json:"childEvents,omitempty"`                  //
	ConnectedInterfaceName       string                                                                      `json:"connectedInterfaceName,omitempty"`       // Connected Interface Name
	DhcpServerIP                 string                                                                      `json:"dhcpServerIp,omitempty"`                 // Dhcp Server Ip
	ManagementIPAddress          string                                                                      `json:"managementIpAddress,omitempty"`          // Management Ip Address
	PreviousRadioPowerLevel      *int                                                                        `json:"previousRadioPowerLevel,omitempty"`      // Previous Radio Power Level
	ResultStatus                 string                                                                      `json:"resultStatus,omitempty"`                 // Result Status
	RadioInterference            string                                                                      `json:"radioInterference,omitempty"`            // Radio Interference
	NetworkDeviceID              string                                                                      `json:"networkDeviceId,omitempty"`              // Network Device Id
	SiteHierarchy                string                                                                      `json:"siteHierarchy,omitempty"`                // Site Hierarchy
	EventStatus                  string                                                                      `json:"eventStatus,omitempty"`                  // Event Status
	WirelessClientEventStartTime *int                                                                        `json:"wirelessClientEventStartTime,omitempty"` // Wireless Client Event Start Time
	SiteHierarchyID              string                                                                      `json:"siteHierarchyId,omitempty"`              // Site Hierarchy Id
	UdnName                      string                                                                      `json:"udnName,omitempty"`                      // Udn Name
	Facility                     string                                                                      `json:"facility,omitempty"`                     // Facility
	LastApResetType              string                                                                      `json:"lastApResetType,omitempty"`              // Last Ap Reset Type
	InvalidIeAPs                 *[]ResponseDevicesQueryAssuranceEventsWithFiltersResponseInvalidIeAPs       `json:"invalidIeAPs,omitempty"`                 //
	Username                     string                                                                      `json:"username,omitempty"`                     // Username
}
type ResponseDevicesQueryAssuranceEventsWithFiltersResponseCandidateAPs struct {
	APID   string `json:"apId,omitempty"`   // Ap Id
	ApName string `json:"apName,omitempty"` // Ap Name
	ApMac  string `json:"apMac,omitempty"`  // Ap Mac
	Bssid  string `json:"bssid,omitempty"`  // Bssid
	Rssi   *int   `json:"rssi,omitempty"`   // Rssi
}
type ResponseDevicesQueryAssuranceEventsWithFiltersResponseMissingResponseAPs struct {
	APID      string `json:"apId,omitempty"`      // Ap Id
	ApName    string `json:"apName,omitempty"`    // Ap Name
	ApMac     string `json:"apMac,omitempty"`     // Ap Mac
	Bssid     string `json:"bssid,omitempty"`     // Bssid
	Type      string `json:"type,omitempty"`      // Type
	FrameType string `json:"frameType,omitempty"` // Frame Type
}
type ResponseDevicesQueryAssuranceEventsWithFiltersResponseChildEvents struct {
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
type ResponseDevicesQueryAssuranceEventsWithFiltersResponseInvalidIeAPs struct {
	APID      string `json:"apId,omitempty"`      // Ap Id
	ApName    string `json:"apName,omitempty"`    // Ap Name
	ApMac     string `json:"apMac,omitempty"`     // Ap Mac
	Bssid     string `json:"bssid,omitempty"`     // Bssid
	Type      string `json:"type,omitempty"`      // Type
	FrameType string `json:"frameType,omitempty"` // Frame Type
	Ies       string `json:"ies,omitempty"`       // Ies
}
type ResponseDevicesQueryAssuranceEventsWithFiltersPage struct {
	Limit  *int                                                        `json:"limit,omitempty"`  // Limit
	Offset *int                                                        `json:"offset,omitempty"` // Offset
	Count  *int                                                        `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesQueryAssuranceEventsWithFiltersPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesQueryAssuranceEventsWithFiltersPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseDevicesCountTheNumberOfEventsWithFilters struct {
	Response *ResponseDevicesCountTheNumberOfEventsWithFiltersResponse `json:"response,omitempty"` //
	Version  string                                                    `json:"version,omitempty"`  // Version
}
type ResponseDevicesCountTheNumberOfEventsWithFiltersResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseDevicesGetDetailsOfASingleAssuranceEvent struct {
	Response *ResponseDevicesGetDetailsOfASingleAssuranceEventResponse `json:"response,omitempty"` //
	Version  string                                                    `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetDetailsOfASingleAssuranceEventResponse struct {
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
	CandidateAPs                 *[]ResponseDevicesGetDetailsOfASingleAssuranceEventResponseCandidateAPs       `json:"candidateAPs,omitempty"`                 //
	ReplacedDeviceSerialNumber   string                                                                        `json:"replacedDeviceSerialNumber,omitempty"`   // Replaced Device Serial Number
	OldRadioChannelList          string                                                                        `json:"oldRadioChannelList,omitempty"`          // Old Radio Channel List
	SSID                         string                                                                        `json:"ssid,omitempty"`                         // Ssid
	SubReasonDescription         string                                                                        `json:"subReasonDescription,omitempty"`         // Sub Reason Description
	WirelessClientEventEndTime   *int                                                                          `json:"wirelessClientEventEndTime,omitempty"`   // Wireless Client Event End Time
	IPv4                         string                                                                        `json:"ipv4,omitempty"`                         // Ipv4
	WlcID                        string                                                                        `json:"wlcId,omitempty"`                        // Wlc Id
	IPv6                         string                                                                        `json:"ipv6,omitempty"`                         // Ipv6
	MissingResponseAPs           *[]ResponseDevicesGetDetailsOfASingleAssuranceEventResponseMissingResponseAPs `json:"missingResponseAPs,omitempty"`           //
	Timestamp                    *int                                                                          `json:"timestamp,omitempty"`                    // Timestamp
	Severity                     *int                                                                          `json:"severity,omitempty"`                     // Severity
	CurrentRadioPowerLevel       *int                                                                          `json:"currentRadioPowerLevel,omitempty"`       // Current Radio Power Level
	NewRadioChannelWidth         string                                                                        `json:"newRadioChannelWidth,omitempty"`         // New Radio Channel Width
	AssocSnr                     *int                                                                          `json:"assocSnr,omitempty"`                     // Assoc Snr
	AuthServerIP                 string                                                                        `json:"authServerIp,omitempty"`                 // Auth Server Ip
	ChildEvents                  *[]ResponseDevicesGetDetailsOfASingleAssuranceEventResponseChildEvents        `json:"childEvents,omitempty"`                  //
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
	InvalidIeAPs                 *[]ResponseDevicesGetDetailsOfASingleAssuranceEventResponseInvalidIeAPs       `json:"invalidIeAPs,omitempty"`                 //
	Username                     string                                                                        `json:"username,omitempty"`                     // Username
}
type ResponseDevicesGetDetailsOfASingleAssuranceEventResponseCandidateAPs struct {
	APID   string `json:"apId,omitempty"`   // Ap Id
	ApName string `json:"apName,omitempty"` // Ap Name
	ApMac  string `json:"apMac,omitempty"`  // Ap Mac
	Bssid  string `json:"bssid,omitempty"`  // Bssid
	Rssi   *int   `json:"rssi,omitempty"`   // Rssi
}
type ResponseDevicesGetDetailsOfASingleAssuranceEventResponseMissingResponseAPs struct {
	APID      string `json:"apId,omitempty"`      // Ap Id
	ApName    string `json:"apName,omitempty"`    // Ap Name
	ApMac     string `json:"apMac,omitempty"`     // Ap Mac
	Bssid     string `json:"bssid,omitempty"`     // Bssid
	Type      string `json:"type,omitempty"`      // Type
	FrameType string `json:"frameType,omitempty"` // Frame Type
}
type ResponseDevicesGetDetailsOfASingleAssuranceEventResponseChildEvents struct {
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
type ResponseDevicesGetDetailsOfASingleAssuranceEventResponseInvalidIeAPs struct {
	APID      string `json:"apId,omitempty"`      // Ap Id
	ApName    string `json:"apName,omitempty"`    // Ap Name
	ApMac     string `json:"apMac,omitempty"`     // Ap Mac
	Bssid     string `json:"bssid,omitempty"`     // Bssid
	Type      string `json:"type,omitempty"`      // Type
	FrameType string `json:"frameType,omitempty"` // Frame Type
	Ies       string `json:"ies,omitempty"`       // Ies
}
type ResponseDevicesGetListOfChildEventsForTheGivenWirelessClientEvent struct {
	Response *[]ResponseDevicesGetListOfChildEventsForTheGivenWirelessClientEventResponse `json:"response,omitempty"` //
	Version  string                                                                       `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetListOfChildEventsForTheGivenWirelessClientEventResponse struct {
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
type ResponseDevicesGetsInterfacesAlongWithStatisticsDataFromAllNetworkDevices struct {
	Response *[]ResponseDevicesGetsInterfacesAlongWithStatisticsDataFromAllNetworkDevicesResponse `json:"response,omitempty"` //
	Page     *ResponseDevicesGetsInterfacesAlongWithStatisticsDataFromAllNetworkDevicesPage       `json:"page,omitempty"`     //
	Version  string                                                                               `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetsInterfacesAlongWithStatisticsDataFromAllNetworkDevicesResponse struct {
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
	SiteName                string   `json:"siteName,omitempty"`                // Site Name
	SiteHierarchy           string   `json:"siteHierarchy,omitempty"`           // Site Hierarchy
	SiteHierarchyID         string   `json:"siteHierarchyId,omitempty"`         // Site Hierarchy Id
}
type ResponseDevicesGetsInterfacesAlongWithStatisticsDataFromAllNetworkDevicesPage struct {
	Limit  *int                                                                                   `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                   `json:"offset,omitempty"` // Offset
	Count  *int                                                                                   `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesGetsInterfacesAlongWithStatisticsDataFromAllNetworkDevicesPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesGetsInterfacesAlongWithStatisticsDataFromAllNetworkDevicesPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseDevicesGetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCount struct {
	Response *ResponseDevicesGetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountResponse `json:"response,omitempty"` //
	Version  string                                                                                                                                                           `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions struct {
	Response *[]ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponse `json:"response,omitempty"` //
	Page     *ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPage       `json:"page,omitempty"`     //
	Version  string                                                                                                                          `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponse struct {
	ID                      string                                                                                                                                             `json:"id,omitempty"`                      // Id
	AdminStatus             string                                                                                                                                             `json:"adminStatus,omitempty"`             // Admin Status
	Description             string                                                                                                                                             `json:"description,omitempty"`             // Description
	DuplexConfig            string                                                                                                                                             `json:"duplexConfig,omitempty"`            // Duplex Config
	DuplexOper              string                                                                                                                                             `json:"duplexOper,omitempty"`              // Duplex Oper
	InterfaceIfIndex        *int                                                                                                                                               `json:"interfaceIfIndex,omitempty"`        // Interface If Index
	InterfaceType           string                                                                                                                                             `json:"interfaceType,omitempty"`           // Interface Type
	IPv4Address             string                                                                                                                                             `json:"ipv4Address,omitempty"`             // Ipv4 Address
	IPv6AddressList         []string                                                                                                                                           `json:"ipv6AddressList,omitempty"`         // Ipv6 Address List
	IsL3Interface           *bool                                                                                                                                              `json:"isL3Interface,omitempty"`           // Is L3 Interface
	IsWan                   *bool                                                                                                                                              `json:"isWan,omitempty"`                   // Is Wan
	MacAddr                 string                                                                                                                                             `json:"macAddr,omitempty"`                 // Mac Addr
	MediaType               string                                                                                                                                             `json:"mediaType,omitempty"`               // Media Type
	Name                    string                                                                                                                                             `json:"name,omitempty"`                    // Name
	OperStatus              string                                                                                                                                             `json:"operStatus,omitempty"`              // Oper Status
	PeerStackMember         *int                                                                                                                                               `json:"peerStackMember,omitempty"`         // Peer Stack Member
	PeerStackPort           string                                                                                                                                             `json:"peerStackPort,omitempty"`           // Peer Stack Port
	PortChannelID           string                                                                                                                                             `json:"portChannelId,omitempty"`           // Port Channel Id
	PortMode                string                                                                                                                                             `json:"portMode,omitempty"`                // Port Mode
	PortType                string                                                                                                                                             `json:"portType,omitempty"`                // Port Type
	RxDiscards              *float64                                                                                                                                           `json:"rxDiscards,omitempty"`              // Rx Discards
	RxError                 *int                                                                                                                                               `json:"rxError,omitempty"`                 // Rx Error
	RxRate                  *float64                                                                                                                                           `json:"rxRate,omitempty"`                  // Rx Rate
	RxUtilization           *float64                                                                                                                                           `json:"rxUtilization,omitempty"`           // Rx Utilization
	Speed                   string                                                                                                                                             `json:"speed,omitempty"`                   // Speed
	StackPortType           string                                                                                                                                             `json:"stackPortType,omitempty"`           // Stack Port Type
	Timestamp               *int                                                                                                                                               `json:"timestamp,omitempty"`               // Timestamp
	TxDiscards              *float64                                                                                                                                           `json:"txDiscards,omitempty"`              // Tx Discards
	TxError                 *int                                                                                                                                               `json:"txError,omitempty"`                 // Tx Error
	TxRate                  *float64                                                                                                                                           `json:"txRate,omitempty"`                  // Tx Rate
	TxUtilization           *float64                                                                                                                                           `json:"txUtilization,omitempty"`           // Tx Utilization
	VLANID                  string                                                                                                                                             `json:"vlanId,omitempty"`                  // Vlan Id
	NetworkDeviceID         string                                                                                                                                             `json:"networkDeviceId,omitempty"`         // Network Device Id
	NetworkDeviceIPAddress  string                                                                                                                                             `json:"networkDeviceIpAddress,omitempty"`  // Network Device Ip Address
	NetworkDeviceMacAddress string                                                                                                                                             `json:"networkDeviceMacAddress,omitempty"` // Network Device Mac Address
	SiteName                string                                                                                                                                             `json:"siteName,omitempty"`                // Site Name
	SiteHierarchy           string                                                                                                                                             `json:"siteHierarchy,omitempty"`           // Site Hierarchy
	SiteHierarchyID         string                                                                                                                                             `json:"siteHierarchyId,omitempty"`         // Site Hierarchy Id
	AggregateAttributes     *[]ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseAggregateAttributes `json:"aggregateAttributes,omitempty"`     //
}
type ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseAggregateAttributes struct {
	Name   string                                                                                                                                                   `json:"name,omitempty"`   // Name
	Values *[]ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseAggregateAttributesValues `json:"values,omitempty"` //
}
type ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseAggregateAttributesValues struct {
	Key   string   `json:"key,omitempty"`   // Key
	Value *float64 `json:"value,omitempty"` // Value
}
type ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPage struct {
	Limit  *int                                                                                                                              `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                                                              `json:"offset,omitempty"` // Offset
	Count  *int                                                                                                                              `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseDevicesTheTotalInterfacesCountAcrossTheNetworkDevices struct {
	Response *ResponseDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesResponse `json:"response,omitempty"` //
	Version  string                                                                 `json:"version,omitempty"`  // Version
}
type ResponseDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseDevicesGetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsData struct {
	Response *ResponseDevicesGetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsDataResponse `json:"response,omitempty"` //
	Version  string                                                                                                  `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsDataResponse struct {
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
	SiteName                string   `json:"siteName,omitempty"`                // Site Name
	SiteHierarchy           string   `json:"siteHierarchy,omitempty"`           // Site Hierarchy
	SiteHierarchyID         string   `json:"siteHierarchyId,omitempty"`         // Site Hierarchy Id
}
type ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParameters struct {
	Response *[]ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersResponse `json:"response,omitempty"` //
	Page     *ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersPage       `json:"page,omitempty"`     //
	Version  string                                                                                 `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersResponse struct {
	ID                         string                                                                                                    `json:"id,omitempty"`                         // Id
	Name                       string                                                                                                    `json:"name,omitempty"`                       // Name
	ManagementIPAddress        string                                                                                                    `json:"managementIpAddress,omitempty"`        // Management Ip Address
	PlatformID                 string                                                                                                    `json:"platformId,omitempty"`                 // Platform Id
	DeviceFamily               string                                                                                                    `json:"deviceFamily,omitempty"`               // Device Family
	SerialNumber               string                                                                                                    `json:"serialNumber,omitempty"`               // Serial Number
	MacAddress                 string                                                                                                    `json:"macAddress,omitempty"`                 // Mac Address
	DeviceSeries               string                                                                                                    `json:"deviceSeries,omitempty"`               // Device Series
	SoftwareVersion            string                                                                                                    `json:"softwareVersion,omitempty"`            // Software Version
	ProductVendor              string                                                                                                    `json:"productVendor,omitempty"`              // Product Vendor
	DeviceRole                 string                                                                                                    `json:"deviceRole,omitempty"`                 // Device Role
	DeviceType                 string                                                                                                    `json:"deviceType,omitempty"`                 // Device Type
	CommunicationState         string                                                                                                    `json:"communicationState,omitempty"`         // Communication State
	CollectionStatus           string                                                                                                    `json:"collectionStatus,omitempty"`           // Collection Status
	HaStatus                   string                                                                                                    `json:"haStatus,omitempty"`                   // Ha Status
	LastBootTime               *int                                                                                                      `json:"lastBootTime,omitempty"`               // Last Boot Time
	SiteHierarchyID            string                                                                                                    `json:"siteHierarchyId,omitempty"`            // Site Hierarchy Id
	SiteHierarchy              string                                                                                                    `json:"siteHierarchy,omitempty"`              // Site Hierarchy
	SiteID                     string                                                                                                    `json:"siteId,omitempty"`                     // Site Id
	DeviceGroupHierarchyID     string                                                                                                    `json:"deviceGroupHierarchyId,omitempty"`     // Device Group Hierarchy Id
	TagNames                   []string                                                                                                  `json:"tagNames,omitempty"`                   // Tag Names
	StackType                  string                                                                                                    `json:"stackType,omitempty"`                  // Stack Type
	OsType                     string                                                                                                    `json:"osType,omitempty"`                     // Os Type
	RingStatus                 *bool                                                                                                     `json:"ringStatus,omitempty"`                 // Ring Status
	MaintenanceModeEnabled     *bool                                                                                                     `json:"maintenanceModeEnabled,omitempty"`     // Maintenance Mode Enabled
	UpTime                     *int                                                                                                      `json:"upTime,omitempty"`                     // Up Time
	IPv4Address                string                                                                                                    `json:"ipv4Address,omitempty"`                // Ipv4 Address
	IPv6Address                string                                                                                                    `json:"ipv6Address,omitempty"`                // Ipv6 Address
	RedundancyMode             string                                                                                                    `json:"redundancyMode,omitempty"`             // Redundancy Mode
	FeatureFlagList            []string                                                                                                  `json:"featureFlagList,omitempty"`            // Feature Flag List
	HaLastResetReason          string                                                                                                    `json:"haLastResetReason,omitempty"`          // Ha Last Reset Reason
	RedundancyPeerStateDerived string                                                                                                    `json:"redundancyPeerStateDerived,omitempty"` // Redundancy Peer State Derived
	RedundancyPeerState        string                                                                                                    `json:"redundancyPeerState,omitempty"`        // Redundancy Peer State
	RedundancyStateDerived     string                                                                                                    `json:"redundancyStateDerived,omitempty"`     // Redundancy State Derived
	RedundancyState            string                                                                                                    `json:"redundancyState,omitempty"`            // Redundancy State
	WiredClientCount           *int                                                                                                      `json:"wiredClientCount,omitempty"`           // Wired Client Count
	WirelessClientCount        *int                                                                                                      `json:"wirelessClientCount,omitempty"`        // Wireless Client Count
	PortCount                  *int                                                                                                      `json:"portCount,omitempty"`                  // Port Count
	ClientCount                *int                                                                                                      `json:"clientCount,omitempty"`                // Client Count
	ApDetails                  *ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersResponseApDetails             `json:"apDetails,omitempty"`                  //
	MetricsDetails             *ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersResponseMetricsDetails        `json:"metricsDetails,omitempty"`             //
	FabricDetails              *ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersResponseFabricDetails         `json:"fabricDetails,omitempty"`              //
	AggregateAttributes        *[]ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersResponseAggregateAttributes `json:"aggregateAttributes,omitempty"`        //
}
type ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersResponseApDetails struct {
	ConnectedWlcName     string                                                                                                `json:"connectedWlcName,omitempty"`     // Connected Wlc Name
	PolicyTagName        string                                                                                                `json:"policyTagName,omitempty"`        // Policy Tag Name
	ApOperationalState   string                                                                                                `json:"apOperationalState,omitempty"`   // Ap Operational State
	PowerSaveMode        string                                                                                                `json:"powerSaveMode,omitempty"`        // Power Save Mode
	OperationalMode      string                                                                                                `json:"operationalMode,omitempty"`      // Operational Mode
	ResetReason          string                                                                                                `json:"resetReason,omitempty"`          // Reset Reason
	Protocol             string                                                                                                `json:"protocol,omitempty"`             // Protocol
	PowerMode            string                                                                                                `json:"powerMode,omitempty"`            // Power Mode
	ConnectedTime        *int                                                                                                  `json:"connectedTime,omitempty"`        // Connected Time
	LedFlashEnabled      *bool                                                                                                 `json:"ledFlashEnabled,omitempty"`      // Led Flash Enabled
	LedFlashSeconds      *int                                                                                                  `json:"ledFlashSeconds,omitempty"`      // Led Flash Seconds
	SubMode              string                                                                                                `json:"subMode,omitempty"`              // Sub Mode
	HomeApEnabled        *bool                                                                                                 `json:"homeApEnabled,omitempty"`        // Home Ap Enabled
	PowerType            string                                                                                                `json:"powerType,omitempty"`            // Power Type
	ApType               string                                                                                                `json:"apType,omitempty"`               // Ap Type
	AdminState           string                                                                                                `json:"adminState,omitempty"`           // Admin State
	IcapCapability       string                                                                                                `json:"icapCapability,omitempty"`       // Icap Capability
	RegulatoryDomain     string                                                                                                `json:"regulatoryDomain,omitempty"`     // Regulatory Domain
	EthernetMac          string                                                                                                `json:"ethernetMac,omitempty"`          // Ethernet Mac
	RfTagName            string                                                                                                `json:"rfTagName,omitempty"`            // Rf Tag Name
	SiteTagName          string                                                                                                `json:"siteTagName,omitempty"`          // Site Tag Name
	PowerSaveModeCapable string                                                                                                `json:"powerSaveModeCapable,omitempty"` // Power Save Mode Capable
	PowerProfile         string                                                                                                `json:"powerProfile,omitempty"`         // Power Profile
	FlexGroup            string                                                                                                `json:"flexGroup,omitempty"`            // Flex Group
	PowerCalendarProfile string                                                                                                `json:"powerCalendarProfile,omitempty"` // Power Calendar Profile
	ApGroup              string                                                                                                `json:"apGroup,omitempty"`              // Ap Group
	Radios               *[]ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersResponseApDetailsRadios `json:"radios,omitempty"`               //
}
type ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersResponseApDetailsRadios struct {
	ID           string   `json:"id,omitempty"`           // Id
	Band         string   `json:"band,omitempty"`         // Band
	Noise        *int     `json:"noise,omitempty"`        // Noise
	AirQuality   *float64 `json:"airQuality,omitempty"`   // Air Quality
	Interference *float64 `json:"interference,omitempty"` // Interference
	TrafficUtil  *int     `json:"trafficUtil,omitempty"`  // Traffic Util
	Utilization  *float64 `json:"utilization,omitempty"`  // Utilization
	ClientCount  *int     `json:"clientCount,omitempty"`  // Client Count
}
type ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersResponseMetricsDetails struct {
	OverallHealthScore                 *int     `json:"overallHealthScore,omitempty"`                 // Overall Health Score
	OverallFabricScore                 *int     `json:"overallFabricScore,omitempty"`                 // Overall Fabric Score
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
type ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersResponseFabricDetails struct {
	FabricRole     []string `json:"fabricRole,omitempty"`     // Fabric Role
	FabricSiteName string   `json:"fabricSiteName,omitempty"` // Fabric Site Name
	TransitFabrics []string `json:"transitFabrics,omitempty"` // Transit Fabrics
}
type ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersResponseAggregateAttributes struct {
	Name     string   `json:"name,omitempty"`     // Name
	Function string   `json:"function,omitempty"` // Function
	Value    *float64 `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersPage struct {
	Limit  *int   `json:"limit,omitempty"`  // Limit
	Offset *int   `json:"offset,omitempty"` // Offset
	Count  *int   `json:"count,omitempty"`  // Count
	SortBy string `json:"sortBy,omitempty"` // Sort By
	Order  string `json:"order,omitempty"`  // Order
}
type ResponseDevicesGetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParameters struct {
	Response *ResponseDevicesGetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersResponse `json:"response,omitempty"` //
	Version  string                                                                                   `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions struct {
	Response *[]ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponse `json:"response,omitempty"` //
	Page     *ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPage       `json:"page,omitempty"`     //
	Version  string                                                                                                       `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponse struct {
	ID                         string                                                                                                                          `json:"id,omitempty"`                         // Id
	Name                       string                                                                                                                          `json:"name,omitempty"`                       // Name
	ManagementIPAddress        string                                                                                                                          `json:"managementIpAddress,omitempty"`        // Management Ip Address
	PlatformID                 string                                                                                                                          `json:"platformId,omitempty"`                 // Platform Id
	DeviceFamily               string                                                                                                                          `json:"deviceFamily,omitempty"`               // Device Family
	SerialNumber               string                                                                                                                          `json:"serialNumber,omitempty"`               // Serial Number
	MacAddress                 string                                                                                                                          `json:"macAddress,omitempty"`                 // Mac Address
	DeviceSeries               string                                                                                                                          `json:"deviceSeries,omitempty"`               // Device Series
	SoftwareVersion            string                                                                                                                          `json:"softwareVersion,omitempty"`            // Software Version
	ProductVendor              string                                                                                                                          `json:"productVendor,omitempty"`              // Product Vendor
	DeviceRole                 string                                                                                                                          `json:"deviceRole,omitempty"`                 // Device Role
	DeviceType                 string                                                                                                                          `json:"deviceType,omitempty"`                 // Device Type
	CommunicationState         string                                                                                                                          `json:"communicationState,omitempty"`         // Communication State
	CollectionStatus           string                                                                                                                          `json:"collectionStatus,omitempty"`           // Collection Status
	HaStatus                   string                                                                                                                          `json:"haStatus,omitempty"`                   // Ha Status
	LastBootTime               *int                                                                                                                            `json:"lastBootTime,omitempty"`               // Last Boot Time
	SiteHierarchyID            string                                                                                                                          `json:"siteHierarchyId,omitempty"`            // Site Hierarchy Id
	SiteHierarchy              string                                                                                                                          `json:"siteHierarchy,omitempty"`              // Site Hierarchy
	SiteID                     string                                                                                                                          `json:"siteId,omitempty"`                     // Site Id
	DeviceGroupHierarchyID     string                                                                                                                          `json:"deviceGroupHierarchyId,omitempty"`     // Device Group Hierarchy Id
	TagNames                   []string                                                                                                                        `json:"tagNames,omitempty"`                   // Tag Names
	StackType                  string                                                                                                                          `json:"stackType,omitempty"`                  // Stack Type
	OsType                     string                                                                                                                          `json:"osType,omitempty"`                     // Os Type
	RingStatus                 *bool                                                                                                                           `json:"ringStatus,omitempty"`                 // Ring Status
	MaintenanceModeEnabled     *bool                                                                                                                           `json:"maintenanceModeEnabled,omitempty"`     // Maintenance Mode Enabled
	UpTime                     *int                                                                                                                            `json:"upTime,omitempty"`                     // Up Time
	IPv4Address                string                                                                                                                          `json:"ipv4Address,omitempty"`                // Ipv4 Address
	IPv6Address                string                                                                                                                          `json:"ipv6Address,omitempty"`                // Ipv6 Address
	RedundancyMode             string                                                                                                                          `json:"redundancyMode,omitempty"`             // Redundancy Mode
	FeatureFlagList            []string                                                                                                                        `json:"featureFlagList,omitempty"`            // Feature Flag List
	HaLastResetReason          string                                                                                                                          `json:"haLastResetReason,omitempty"`          // Ha Last Reset Reason
	RedundancyPeerStateDerived string                                                                                                                          `json:"redundancyPeerStateDerived,omitempty"` // Redundancy Peer State Derived
	RedundancyPeerState        string                                                                                                                          `json:"redundancyPeerState,omitempty"`        // Redundancy Peer State
	RedundancyStateDerived     string                                                                                                                          `json:"redundancyStateDerived,omitempty"`     // Redundancy State Derived
	RedundancyState            string                                                                                                                          `json:"redundancyState,omitempty"`            // Redundancy State
	WiredClientCount           *int                                                                                                                            `json:"wiredClientCount,omitempty"`           // Wired Client Count
	WirelessClientCount        *int                                                                                                                            `json:"wirelessClientCount,omitempty"`        // Wireless Client Count
	PortCount                  *int                                                                                                                            `json:"portCount,omitempty"`                  // Port Count
	ClientCount                *int                                                                                                                            `json:"clientCount,omitempty"`                // Client Count
	ApDetails                  *ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseApDetails             `json:"apDetails,omitempty"`                  //
	MetricsDetails             *ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseMetricsDetails        `json:"metricsDetails,omitempty"`             //
	FabricDetails              *ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseFabricDetails         `json:"fabricDetails,omitempty"`              //
	AggregateAttributes        *[]ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseAggregateAttributes `json:"aggregateAttributes,omitempty"`        //
}
type ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseApDetails struct {
	ConnectedWlcName     string                                                                                                                      `json:"connectedWlcName,omitempty"`     // Connected Wlc Name
	PolicyTagName        string                                                                                                                      `json:"policyTagName,omitempty"`        // Policy Tag Name
	ApOperationalState   string                                                                                                                      `json:"apOperationalState,omitempty"`   // Ap Operational State
	PowerSaveMode        string                                                                                                                      `json:"powerSaveMode,omitempty"`        // Power Save Mode
	OperationalMode      string                                                                                                                      `json:"operationalMode,omitempty"`      // Operational Mode
	ResetReason          string                                                                                                                      `json:"resetReason,omitempty"`          // Reset Reason
	Protocol             string                                                                                                                      `json:"protocol,omitempty"`             // Protocol
	PowerMode            string                                                                                                                      `json:"powerMode,omitempty"`            // Power Mode
	ConnectedTime        *int                                                                                                                        `json:"connectedTime,omitempty"`        // Connected Time
	LedFlashEnabled      *bool                                                                                                                       `json:"ledFlashEnabled,omitempty"`      // Led Flash Enabled
	LedFlashSeconds      *int                                                                                                                        `json:"ledFlashSeconds,omitempty"`      // Led Flash Seconds
	SubMode              string                                                                                                                      `json:"subMode,omitempty"`              // Sub Mode
	HomeApEnabled        *bool                                                                                                                       `json:"homeApEnabled,omitempty"`        // Home Ap Enabled
	PowerType            string                                                                                                                      `json:"powerType,omitempty"`            // Power Type
	ApType               string                                                                                                                      `json:"apType,omitempty"`               // Ap Type
	AdminState           string                                                                                                                      `json:"adminState,omitempty"`           // Admin State
	IcapCapability       string                                                                                                                      `json:"icapCapability,omitempty"`       // Icap Capability
	RegulatoryDomain     string                                                                                                                      `json:"regulatoryDomain,omitempty"`     // Regulatory Domain
	EthernetMac          string                                                                                                                      `json:"ethernetMac,omitempty"`          // Ethernet Mac
	RfTagName            string                                                                                                                      `json:"rfTagName,omitempty"`            // Rf Tag Name
	SiteTagName          string                                                                                                                      `json:"siteTagName,omitempty"`          // Site Tag Name
	PowerSaveModeCapable string                                                                                                                      `json:"powerSaveModeCapable,omitempty"` // Power Save Mode Capable
	PowerProfile         string                                                                                                                      `json:"powerProfile,omitempty"`         // Power Profile
	FlexGroup            string                                                                                                                      `json:"flexGroup,omitempty"`            // Flex Group
	PowerCalendarProfile string                                                                                                                      `json:"powerCalendarProfile,omitempty"` // Power Calendar Profile
	ApGroup              string                                                                                                                      `json:"apGroup,omitempty"`              // Ap Group
	Radios               *[]ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseApDetailsRadios `json:"radios,omitempty"`               //
}
type ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseApDetailsRadios struct {
	ID           string   `json:"id,omitempty"`           // Id
	Band         string   `json:"band,omitempty"`         // Band
	Noise        *int     `json:"noise,omitempty"`        // Noise
	AirQuality   *float64 `json:"airQuality,omitempty"`   // Air Quality
	Interference *float64 `json:"interference,omitempty"` // Interference
	TrafficUtil  *int     `json:"trafficUtil,omitempty"`  // Traffic Util
	Utilization  *float64 `json:"utilization,omitempty"`  // Utilization
	ClientCount  *int     `json:"clientCount,omitempty"`  // Client Count
}
type ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseMetricsDetails struct {
	OverallHealthScore                 *int     `json:"overallHealthScore,omitempty"`                 // Overall Health Score
	OverallFabricScore                 *int     `json:"overallFabricScore,omitempty"`                 // Overall Fabric Score
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
type ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseFabricDetails struct {
	FabricRole     []string `json:"fabricRole,omitempty"`     // Fabric Role
	FabricSiteName string   `json:"fabricSiteName,omitempty"` // Fabric Site Name
	TransitFabrics []string `json:"transitFabrics,omitempty"` // Transit Fabrics
}
type ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseAggregateAttributes struct {
	Name     string   `json:"name,omitempty"`     // Name
	Function string   `json:"function,omitempty"` // Function
	Value    *float64 `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPage struct {
	Limit  *int   `json:"limit,omitempty"`  // Limit
	Offset *int   `json:"offset,omitempty"` // Offset
	Count  *int   `json:"count,omitempty"`  // Count
	SortBy string `json:"sortBy,omitempty"` // Sort By
	Order  string `json:"order,omitempty"`  // Order
}
type ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevices struct {
	Response *ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesResponse `json:"response,omitempty"` //
	Page     *[]ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesPage   `json:"page,omitempty"`     //
}
type ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesResponse struct {
	Attributes          *[]ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesResponseAttributes          `json:"attributes,omitempty"`          // Attributes
	AggregateAttributes *[]ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` // Aggregate Attributes
	Groups              *[]ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesResponseGroups              `json:"groups,omitempty"`              //
}
type ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesResponseAttributes interface{}
type ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesResponseAggregateAttributes interface{}
type ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesResponseGroups struct {
	ID                  string                                                                                                `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesResponseGroupsAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesResponseGroupsAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesResponseGroupsAggregateAttributes struct {
	Name     string   `json:"name,omitempty"`     // Name
	Function string   `json:"function,omitempty"` // Function
	Value    *float64 `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesPage struct {
	Limit  *int                                                                           `json:"limit,omitempty"`  // Limit
	Offset *int                                                                           `json:"offset,omitempty"` // Offset
	Count  *int                                                                           `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseDevicesGetsTheTrendAnalyticsData struct {
	Response *[]ResponseDevicesGetsTheTrendAnalyticsDataResponse `json:"response,omitempty"` //
	Page     *ResponseDevicesGetsTheTrendAnalyticsDataPage       `json:"page,omitempty"`     //
	Version  string                                              `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetsTheTrendAnalyticsDataResponse struct {
	Timestamp           *float64                                                               `json:"timestamp,omitempty"`           // Timestamp
	Attributes          *[]ResponseDevicesGetsTheTrendAnalyticsDataResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetsTheTrendAnalyticsDataResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
	Groups              *[]ResponseDevicesGetsTheTrendAnalyticsDataResponseGroups              `json:"groups,omitempty"`              //
}
type ResponseDevicesGetsTheTrendAnalyticsDataResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetsTheTrendAnalyticsDataResponseAggregateAttributes struct {
	Name     string   `json:"name,omitempty"`     // Name
	Function string   `json:"function,omitempty"` // Function
	Value    *float64 `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetsTheTrendAnalyticsDataResponseGroups struct {
	ID                  string                                                                       `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseDevicesGetsTheTrendAnalyticsDataResponseGroupsAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetsTheTrendAnalyticsDataResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetsTheTrendAnalyticsDataResponseGroupsAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetsTheTrendAnalyticsDataResponseGroupsAggregateAttributes struct {
	Name     string   `json:"name,omitempty"`     // Name
	Function string   `json:"function,omitempty"` // Function
	Value    *float64 `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetsTheTrendAnalyticsDataPage struct {
	Limit          *int   `json:"limit,omitempty"`          // Limit
	Offset         *int   `json:"offset,omitempty"`         // Offset
	Count          *int   `json:"count,omitempty"`          // Count
	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUID struct {
	Response *ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDResponse `json:"response,omitempty"` //
	Version  string                                                          `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDResponse struct {
	ID                         string                                                                               `json:"id,omitempty"`                         // Id
	Name                       string                                                                               `json:"name,omitempty"`                       // Name
	ManagementIPAddress        string                                                                               `json:"managementIpAddress,omitempty"`        // Management Ip Address
	PlatformID                 string                                                                               `json:"platformId,omitempty"`                 // Platform Id
	DeviceFamily               string                                                                               `json:"deviceFamily,omitempty"`               // Device Family
	SerialNumber               string                                                                               `json:"serialNumber,omitempty"`               // Serial Number
	MacAddress                 string                                                                               `json:"macAddress,omitempty"`                 // Mac Address
	DeviceSeries               string                                                                               `json:"deviceSeries,omitempty"`               // Device Series
	SoftwareVersion            string                                                                               `json:"softwareVersion,omitempty"`            // Software Version
	ProductVendor              string                                                                               `json:"productVendor,omitempty"`              // Product Vendor
	DeviceRole                 string                                                                               `json:"deviceRole,omitempty"`                 // Device Role
	DeviceType                 string                                                                               `json:"deviceType,omitempty"`                 // Device Type
	CommunicationState         string                                                                               `json:"communicationState,omitempty"`         // Communication State
	CollectionStatus           string                                                                               `json:"collectionStatus,omitempty"`           // Collection Status
	HaStatus                   string                                                                               `json:"haStatus,omitempty"`                   // Ha Status
	LastBootTime               *int                                                                                 `json:"lastBootTime,omitempty"`               // Last Boot Time
	SiteHierarchyID            string                                                                               `json:"siteHierarchyId,omitempty"`            // Site Hierarchy Id
	SiteHierarchy              string                                                                               `json:"siteHierarchy,omitempty"`              // Site Hierarchy
	SiteID                     string                                                                               `json:"siteId,omitempty"`                     // Site Id
	DeviceGroupHierarchyID     string                                                                               `json:"deviceGroupHierarchyId,omitempty"`     // Device Group Hierarchy Id
	TagNames                   []string                                                                             `json:"tagNames,omitempty"`                   // Tag Names
	StackType                  string                                                                               `json:"stackType,omitempty"`                  // Stack Type
	OsType                     string                                                                               `json:"osType,omitempty"`                     // Os Type
	RingStatus                 *bool                                                                                `json:"ringStatus,omitempty"`                 // Ring Status
	MaintenanceModeEnabled     *bool                                                                                `json:"maintenanceModeEnabled,omitempty"`     // Maintenance Mode Enabled
	UpTime                     *int                                                                                 `json:"upTime,omitempty"`                     // Up Time
	IPv4Address                string                                                                               `json:"ipv4Address,omitempty"`                // Ipv4 Address
	IPv6Address                string                                                                               `json:"ipv6Address,omitempty"`                // Ipv6 Address
	RedundancyMode             string                                                                               `json:"redundancyMode,omitempty"`             // Redundancy Mode
	FeatureFlagList            []string                                                                             `json:"featureFlagList,omitempty"`            // Feature Flag List
	HaLastResetReason          string                                                                               `json:"haLastResetReason,omitempty"`          // Ha Last Reset Reason
	RedundancyPeerStateDerived string                                                                               `json:"redundancyPeerStateDerived,omitempty"` // Redundancy Peer State Derived
	RedundancyPeerState        string                                                                               `json:"redundancyPeerState,omitempty"`        // Redundancy Peer State
	RedundancyStateDerived     string                                                                               `json:"redundancyStateDerived,omitempty"`     // Redundancy State Derived
	RedundancyState            string                                                                               `json:"redundancyState,omitempty"`            // Redundancy State
	WiredClientCount           *int                                                                                 `json:"wiredClientCount,omitempty"`           // Wired Client Count
	WirelessClientCount        *int                                                                                 `json:"wirelessClientCount,omitempty"`        // Wireless Client Count
	PortCount                  *int                                                                                 `json:"portCount,omitempty"`                  // Port Count
	ClientCount                *int                                                                                 `json:"clientCount,omitempty"`                // Client Count
	ApDetails                  *ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDResponseApDetails             `json:"apDetails,omitempty"`                  //
	MetricsDetails             *ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDResponseMetricsDetails        `json:"metricsDetails,omitempty"`             //
	FabricDetails              *ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDResponseFabricDetails         `json:"fabricDetails,omitempty"`              //
	AggregateAttributes        *[]ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDResponseAggregateAttributes `json:"aggregateAttributes,omitempty"`        //
}
type ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDResponseApDetails struct {
	ConnectedWlcName     string                                                                           `json:"connectedWlcName,omitempty"`     // Connected Wlc Name
	PolicyTagName        string                                                                           `json:"policyTagName,omitempty"`        // Policy Tag Name
	ApOperationalState   string                                                                           `json:"apOperationalState,omitempty"`   // Ap Operational State
	PowerSaveMode        string                                                                           `json:"powerSaveMode,omitempty"`        // Power Save Mode
	OperationalMode      string                                                                           `json:"operationalMode,omitempty"`      // Operational Mode
	ResetReason          string                                                                           `json:"resetReason,omitempty"`          // Reset Reason
	Protocol             string                                                                           `json:"protocol,omitempty"`             // Protocol
	PowerMode            string                                                                           `json:"powerMode,omitempty"`            // Power Mode
	ConnectedTime        *int                                                                             `json:"connectedTime,omitempty"`        // Connected Time
	LedFlashEnabled      *bool                                                                            `json:"ledFlashEnabled,omitempty"`      // Led Flash Enabled
	LedFlashSeconds      *int                                                                             `json:"ledFlashSeconds,omitempty"`      // Led Flash Seconds
	SubMode              string                                                                           `json:"subMode,omitempty"`              // Sub Mode
	HomeApEnabled        *bool                                                                            `json:"homeApEnabled,omitempty"`        // Home Ap Enabled
	PowerType            string                                                                           `json:"powerType,omitempty"`            // Power Type
	ApType               string                                                                           `json:"apType,omitempty"`               // Ap Type
	AdminState           string                                                                           `json:"adminState,omitempty"`           // Admin State
	IcapCapability       string                                                                           `json:"icapCapability,omitempty"`       // Icap Capability
	RegulatoryDomain     string                                                                           `json:"regulatoryDomain,omitempty"`     // Regulatory Domain
	EthernetMac          string                                                                           `json:"ethernetMac,omitempty"`          // Ethernet Mac
	RfTagName            string                                                                           `json:"rfTagName,omitempty"`            // Rf Tag Name
	SiteTagName          string                                                                           `json:"siteTagName,omitempty"`          // Site Tag Name
	PowerSaveModeCapable string                                                                           `json:"powerSaveModeCapable,omitempty"` // Power Save Mode Capable
	PowerProfile         string                                                                           `json:"powerProfile,omitempty"`         // Power Profile
	FlexGroup            string                                                                           `json:"flexGroup,omitempty"`            // Flex Group
	PowerCalendarProfile string                                                                           `json:"powerCalendarProfile,omitempty"` // Power Calendar Profile
	ApGroup              string                                                                           `json:"apGroup,omitempty"`              // Ap Group
	Radios               *[]ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDResponseApDetailsRadios `json:"radios,omitempty"`               //
}
type ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDResponseApDetailsRadios struct {
	ID           string   `json:"id,omitempty"`           // Id
	Band         string   `json:"band,omitempty"`         // Band
	Noise        *int     `json:"noise,omitempty"`        // Noise
	AirQuality   *float64 `json:"airQuality,omitempty"`   // Air Quality
	Interference *float64 `json:"interference,omitempty"` // Interference
	TrafficUtil  *int     `json:"trafficUtil,omitempty"`  // Traffic Util
	Utilization  *float64 `json:"utilization,omitempty"`  // Utilization
	ClientCount  *int     `json:"clientCount,omitempty"`  // Client Count
}
type ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDResponseMetricsDetails struct {
	OverallHealthScore                 *int     `json:"overallHealthScore,omitempty"`                 // Overall Health Score
	OverallFabricScore                 *int     `json:"overallFabricScore,omitempty"`                 // Overall Fabric Score
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
type ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDResponseFabricDetails struct {
	FabricRole     []string `json:"fabricRole,omitempty"`     // Fabric Role
	FabricSiteName string   `json:"fabricSiteName,omitempty"` // Fabric Site Name
	TransitFabrics []string `json:"transitFabrics,omitempty"` // Transit Fabrics
}
type ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDResponseAggregateAttributes struct {
	Name     string   `json:"name,omitempty"`     // Name
	Function string   `json:"function,omitempty"` // Function
	Value    *float64 `json:"value,omitempty"`    // Value
}
type ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRange struct {
	Response *[]ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeResponse `json:"response,omitempty"` //
	Page     *ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangePage       `json:"page,omitempty"`     //
	Version  string                                                                                    `json:"version,omitempty"`  // Version
}
type ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeResponse struct {
	Timestamp           *float64                                                                                                     `json:"timestamp,omitempty"`           // Timestamp
	Attributes          *[]ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
	Groups              *[]ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeResponseGroups              `json:"groups,omitempty"`              //
}
type ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeResponseAggregateAttributes struct {
	Name     string   `json:"name,omitempty"`     // Name
	Function string   `json:"function,omitempty"` // Function
	Value    *float64 `json:"value,omitempty"`    // Value
}
type ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeResponseGroups struct {
	ID                  string                                                                                                             `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeResponseGroupsAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeResponseGroupsAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeResponseGroupsAggregateAttributes struct {
	Name     string   `json:"name,omitempty"`     // Name
	Function string   `json:"function,omitempty"` // Function
	Value    *float64 `json:"value,omitempty"`    // Value
}
type ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangePage struct {
	Limit          *int   `json:"limit,omitempty"`          // Limit
	Offset         *int   `json:"offset,omitempty"`         // Offset
	Count          *int   `json:"count,omitempty"`          // Count
	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type ResponseDevicesGetPlannedAccessPointsForBuilding struct {
	Response *[]ResponseDevicesGetPlannedAccessPointsForBuildingResponse `json:"response,omitempty"` //
	Version  *int                                                        `json:"version,omitempty"`  // Version of the api response model
	Total    *int                                                        `json:"total,omitempty"`    // Total number of the planned access points
}
type ResponseDevicesGetPlannedAccessPointsForBuildingResponse struct {
	Attributes *ResponseDevicesGetPlannedAccessPointsForBuildingResponseAttributes `json:"attributes,omitempty"` //
	Location   *ResponseDevicesGetPlannedAccessPointsForBuildingResponseLocation   `json:"location,omitempty"`   //
	Position   *ResponseDevicesGetPlannedAccessPointsForBuildingResponsePosition   `json:"position,omitempty"`   //
	RadioCount *int                                                                `json:"radioCount,omitempty"` // Number of radios of the planned access point
	Radios     *[]ResponseDevicesGetPlannedAccessPointsForBuildingResponseRadios   `json:"radios,omitempty"`     //
	IsSensor   *bool                                                               `json:"isSensor,omitempty"`   // Determines if the planned access point is sensor or not
}
type ResponseDevicesGetPlannedAccessPointsForBuildingResponseAttributes struct {
	ID            *float64 `json:"id,omitempty"`            // Unique id of the planned access point
	InstanceUUID  string   `json:"instanceUuid,omitempty"`  // Instance uuid of the planned access point
	Name          string   `json:"name,omitempty"`          // Display name of the planned access point
	TypeString    string   `json:"typeString,omitempty"`    // Type string representation of the planned access point
	Domain        string   `json:"domain,omitempty"`        // Service domain to which the planned access point belongs
	HeirarchyName string   `json:"heirarchyName,omitempty"` // Hierarchy name of the planned access point
	Source        string   `json:"source,omitempty"`        // Source of the data used to create the planned access point
	CreateDate    *float64 `json:"createDate,omitempty"`    // Created date of the planned access point
	MacAddress    string   `json:"macAddress,omitempty"`    // MAC address of the planned access point
}
type ResponseDevicesGetPlannedAccessPointsForBuildingResponseLocation struct {
	Altitude   *float64 `json:"altitude,omitempty"`   // Altitude of the planned access point's location
	Lattitude  *float64 `json:"lattitude,omitempty"`  // Latitude of the planned access point's location
	Longtitude *float64 `json:"longtitude,omitempty"` // Longitude of the planned access point's location
}
type ResponseDevicesGetPlannedAccessPointsForBuildingResponsePosition struct {
	X *float64 `json:"x,omitempty"` // x-coordinate of the planned access point on the map, 0,0 point being the top-left corner
	Y *float64 `json:"y,omitempty"` // y-coordinate of the planned access point on the map, 0,0 point being the top-left corner
	Z *float64 `json:"z,omitempty"` // z-coordinate, or height, of the planned access point on the map
}
type ResponseDevicesGetPlannedAccessPointsForBuildingResponseRadios struct {
	Attributes *ResponseDevicesGetPlannedAccessPointsForBuildingResponseRadiosAttributes `json:"attributes,omitempty"` //
	Antenna    *ResponseDevicesGetPlannedAccessPointsForBuildingResponseRadiosAntenna    `json:"antenna,omitempty"`    //
	IsSensor   *bool                                                                     `json:"isSensor,omitempty"`   // Determines if it is sensor or not
}
type ResponseDevicesGetPlannedAccessPointsForBuildingResponseRadiosAttributes struct {
	ID            *int     `json:"id,omitempty"`            // Id of the radio
	InstanceUUID  string   `json:"instanceUuid,omitempty"`  // Instance Uuid of the radio
	SlotID        *float64 `json:"slotId,omitempty"`        // Slot number in which the radio resides in the parent access point
	IfTypeString  string   `json:"ifTypeString,omitempty"`  // String representation of native band
	IfTypeSubband string   `json:"ifTypeSubband,omitempty"` // Sub band type of the radio
	Channel       *float64 `json:"channel,omitempty"`       // Channel in which the radio operates
	ChannelString string   `json:"channelString,omitempty"` // Channel string representation
	IfMode        string   `json:"ifMode,omitempty"`        // IF mode of the radio
	TxPowerLevel  *float64 `json:"txPowerLevel,omitempty"`  // Tx Power at which this radio operates (in dBm)
}
type ResponseDevicesGetPlannedAccessPointsForBuildingResponseRadiosAntenna struct {
	Name           string   `json:"name,omitempty"`           // Name of the antenna
	Type           string   `json:"type,omitempty"`           // Type of the antenna associated with this radio
	Mode           string   `json:"mode,omitempty"`           // Mode of the antenna associated with this radio
	AzimuthAngle   *float64 `json:"azimuthAngle,omitempty"`   // Azimuth angle of the antenna
	ElevationAngle *float64 `json:"elevationAngle,omitempty"` // Elevation angle of the antenna
	Gain           *float64 `json:"gain,omitempty"`           // Gain of the antenna
}
type ResponseDevicesGetDeviceDetail struct {
	Response *ResponseDevicesGetDeviceDetailResponse `json:"response,omitempty"` //
}
type ResponseDevicesGetDeviceDetailResponse struct {
	NoiseScore                 *int                                               `json:"noiseScore,omitempty"`                 // Device (AP) WIFI signal noise health score
	PolicyTagName              string                                             `json:"policyTagName,omitempty"`              // Device (AP) policy tag
	InterferenceScore          *int                                               `json:"interferenceScore,omitempty"`          // Device (AP) WIFI signal interference health score
	OpState                    string                                             `json:"opState,omitempty"`                    // Operation state of device (AP)
	PowerSaveMode              string                                             `json:"powerSaveMode,omitempty"`              // Device power save mode
	Mode                       string                                             `json:"mode,omitempty"`                       // Device mode (AP)
	ResetReason                string                                             `json:"resetReason,omitempty"`                // Device reset reason
	NwDeviceRole               string                                             `json:"nwDeviceRole,omitempty"`               // Device role
	Protocol                   string                                             `json:"protocol,omitempty"`                   // Protocol code
	PowerMode                  string                                             `json:"powerMode,omitempty"`                  // Device's power mode
	ConnectedTime              string                                             `json:"connectedTime,omitempty"`              // UTC timestamp
	RingStatus                 *bool                                              `json:"ringStatus,omitempty"`                 // Device's ring status
	LedFlashSeconds            string                                             `json:"ledFlashSeconds,omitempty"`            // LED flash seconds
	IPAddrManagementIPAddr     string                                             `json:"ip_addr_managementIpAddr,omitempty"`   // Device's management IP address
	StackType                  string                                             `json:"stackType,omitempty"`                  // Device stack type (applicable for stackable devices)
	SubMode                    string                                             `json:"subMode,omitempty"`                    // Device submode
	SerialNumber               string                                             `json:"serialNumber,omitempty"`               // Device serial number
	NwDeviceName               string                                             `json:"nwDeviceName,omitempty"`               // Device name
	DeviceGroupHierarchyID     string                                             `json:"deviceGroupHierarchyId,omitempty"`     // Device group site hierarchy UUID
	CPU                        *float64                                           `json:"cpu,omitempty"`                        // Device CPU utilization
	Utilization                string                                             `json:"utilization,omitempty"`                // Device utilization
	NwDeviceID                 string                                             `json:"nwDeviceId,omitempty"`                 // Device's UUID
	SiteHierarchyGraphID       string                                             `json:"siteHierarchyGraphId,omitempty"`       // Site hierarchy UUID in which device is assigned to
	NwDeviceFamily             string                                             `json:"nwDeviceFamily,omitempty"`             // Device faimly string
	MacAddress                 string                                             `json:"macAddress,omitempty"`                 // Device MAC address
	HomeApEnabled              string                                             `json:"homeApEnabled,omitempty"`              // Home Ap Enabled
	DeviceSeries               string                                             `json:"deviceSeries,omitempty"`               // Device series string
	CollectionStatus           string                                             `json:"collectionStatus,omitempty"`           // Device's telemetry data collection status for CATALYST
	UtilizationScore           *int                                               `json:"utilizationScore,omitempty"`           // Device utilization health score
	MaintenanceMode            *bool                                              `json:"maintenanceMode,omitempty"`            // Whether device is in maintenance mode
	Interference               string                                             `json:"interference,omitempty"`               // Device (AP) WIFI signal interference
	SoftwareVersion            string                                             `json:"softwareVersion,omitempty"`            // Device's software version string
	TagIDList                  *[]ResponseDevicesGetDeviceDetailResponseTagIDList `json:"tagIdList,omitempty"`                  // Tag ID List
	PowerType                  string                                             `json:"powerType,omitempty"`                  // Device (AP) power type
	OverallHealth              *int                                               `json:"overallHealth,omitempty"`              // Device's overall health score
	ManagementIPAddr           string                                             `json:"managementIpAddr,omitempty"`           // Management IP address of the device
	Memory                     string                                             `json:"memory,omitempty"`                     // Device memory utilization
	CommunicationState         string                                             `json:"communicationState,omitempty"`         // Device communication state
	ApType                     string                                             `json:"apType,omitempty"`                     // Ap Type
	AdminState                 string                                             `json:"adminState,omitempty"`                 // Device (AP) admin state
	Noise                      string                                             `json:"noise,omitempty"`                      // Device (AP) WIFI signal noise
	IcapCapability             string                                             `json:"icapCapability,omitempty"`             // Device (AP) ICAP capability bit values
	RegulatoryDomain           string                                             `json:"regulatoryDomain,omitempty"`           // Device (AP) WIFI domain
	EthernetMac                string                                             `json:"ethernetMac,omitempty"`                // Device (AP) ethernet MAC address
	NwDeviceType               string                                             `json:"nwDeviceType,omitempty"`               // Device type
	AirQuality                 string                                             `json:"airQuality,omitempty"`                 // Device (AP) WIFI air quality
	RfTagName                  string                                             `json:"rfTagName,omitempty"`                  // Device (AP) RF tag name
	SiteTagName                string                                             `json:"siteTagName,omitempty"`                // Device (AP) site tag name
	PlatformID                 string                                             `json:"platformId,omitempty"`                 // Device's platform ID
	UpTime                     string                                             `json:"upTime,omitempty"`                     // Device up time
	MemoryScore                *int                                               `json:"memoryScore,omitempty"`                // Device's memory usage score
	PowerSaveModeCapable       string                                             `json:"powerSaveModeCapable,omitempty"`       // Device (AP) power save mode capability
	PowerProfile               string                                             `json:"powerProfile,omitempty"`               // Device (AP) power profile name
	AirQualityScore            *int                                               `json:"airQualityScore,omitempty"`            // Device (AP) air quality health score
	Location                   string                                             `json:"location,omitempty"`                   // Device's site hierarchy UUID
	FlexGroup                  string                                             `json:"flexGroup,omitempty"`                  // Deivce (A) flexconnect group
	LastBootTime               *float64                                           `json:"lastBootTime,omitempty"`               // Device's last boot UTC timestamp
	PowerCalendarProfile       string                                             `json:"powerCalendarProfile,omitempty"`       // Device (AP) power calendar profile name
	ConnectivityStatus         *int                                               `json:"connectivityStatus,omitempty"`         // Device connectivity status
	LedFlashEnabled            string                                             `json:"ledFlashEnabled,omitempty"`            // Device (AP) LED flash
	CPUScore                   *int                                               `json:"cpuScore,omitempty"`                   // Device's CPU usage score
	AvgTemperature             *float64                                           `json:"avgTemperature,omitempty"`             // Device's average temperature
	MaxTemperature             *float64                                           `json:"maxTemperature,omitempty"`             // Device's max temperature
	HaStatus                   string                                             `json:"haStatus,omitempty"`                   // Device's HA status
	OsType                     string                                             `json:"osType,omitempty"`                     // Device's OS type
	Timestamp                  *int                                               `json:"timestamp,omitempty"`                  // UTC timestamp of the device health data
	ApGroup                    string                                             `json:"apGroup,omitempty"`                    // Device (AP) AP group
	RedundancyMode             string                                             `json:"redundancyMode,omitempty"`             // Device redundancy mode
	FeatureFlagList            []string                                           `json:"featureFlagList,omitempty"`            // List of device feature capabilities
	FreeMbufScore              *int                                               `json:"freeMbufScore,omitempty"`              // Free memory buffer health score
	HALastResetReason          string                                             `json:"HALastResetReason,omitempty"`          // Last HA reset reason
	WqeScore                   *int                                               `json:"wqeScore,omitempty"`                   // WQE health score
	RedundancyPeerStateDerived string                                             `json:"redundancyPeerStateDerived,omitempty"` // Redundancy Peer State Derived
	FreeTimerScore             *int                                               `json:"freeTimerScore,omitempty"`             // Free Timer Score
	RedundancyPeerState        string                                             `json:"redundancyPeerState,omitempty"`        // Redundancy Peer State
	RedundancyStateDerived     string                                             `json:"redundancyStateDerived,omitempty"`     // Derived redundancy state
	RedundancyState            string                                             `json:"redundancyState,omitempty"`            // Redundancy state
	PacketPoolScore            *int                                               `json:"packetPoolScore,omitempty"`            // Device packet pool health score
	FreeTimer                  *float64                                           `json:"freeTimer,omitempty"`                  // Free timer of the device
	PacketPool                 *float64                                           `json:"packetPool,omitempty"`                 // Packet pool of the device
	Wqe                        *float64                                           `json:"wqe,omitempty"`                        // WQE of the device
	FreeMbuf                   *float64                                           `json:"freeMbuf,omitempty"`                   // Free memory buffer of the device
}
type ResponseDevicesGetDeviceDetailResponseTagIDList interface{}
type ResponseDevicesGetDeviceEnrichmentDetails []ResponseItemDevicesGetDeviceEnrichmentDetails // Array of ResponseDevicesGetDeviceEnrichmentDetails
type ResponseItemDevicesGetDeviceEnrichmentDetails struct {
	DeviceDetails *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetails `json:"deviceDetails,omitempty"` //
}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetails struct {
	Family                    string                                                                        `json:"family,omitempty"`                    // Device Family
	Type                      string                                                                        `json:"type,omitempty"`                      // Device Type
	Location                  *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsLocation           `json:"location,omitempty"`                  // Device location - Site hierarchy
	ErrorCode                 string                                                                        `json:"errorCode,omitempty"`                 // Inventory status error code
	MacAddress                string                                                                        `json:"macAddress,omitempty"`                // Device MAC address
	Role                      string                                                                        `json:"role,omitempty"`                      // Device role
	ApManagerInterfaceIP      string                                                                        `json:"apManagerInterfaceIp,omitempty"`      // IP address of WLC on AP manager interface
	AssociatedWlcIP           string                                                                        `json:"associatedWlcIp,omitempty"`           // Associated WLC IP address of the AP device
	BootDateTime              string                                                                        `json:"bootDateTime,omitempty"`              // Device's last boot UTC timestamp
	CollectionStatus          string                                                                        `json:"collectionStatus,omitempty"`          // Device's telemetry data collection status for CATALYST
	InterfaceCount            string                                                                        `json:"interfaceCount,omitempty"`            // Number of interfaces on the device
	LineCardCount             string                                                                        `json:"lineCardCount,omitempty"`             // Number of linecards on the device
	LineCardID                string                                                                        `json:"lineCardId,omitempty"`                // IDs of linecards of the device
	ManagementIPAddress       string                                                                        `json:"managementIpAddress,omitempty"`       // Device Management Ip Address
	MemorySize                string                                                                        `json:"memorySize,omitempty"`                // Processor memory size
	PlatformID                string                                                                        `json:"platformId,omitempty"`                // Device's platform ID
	ReachabilityFailureReason string                                                                        `json:"reachabilityFailureReason,omitempty"` // Failure reason for unreachable devices
	ReachabilityStatus        string                                                                        `json:"reachabilityStatus,omitempty"`        // Reachability Status of the Device(Reachable/Unreachable)
	SNMPContact               string                                                                        `json:"snmpContact,omitempty"`               // SNMP contact on device
	SNMPLocation              string                                                                        `json:"snmpLocation,omitempty"`              // SNMP location on device
	TunnelUDPPort             *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsTunnelUDPPort      `json:"tunnelUdpPort,omitempty"`             // Mobility protocol port is stored as tunneludpport for WLC
	WaasDeviceMode            *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsWaasDeviceMode     `json:"waasDeviceMode,omitempty"`            // WAAS device mode
	Series                    string                                                                        `json:"series,omitempty"`                    // Device Series
	InventoryStatusDetail     string                                                                        `json:"inventoryStatusDetail,omitempty"`     // Status detail of inventory sync
	CollectionInterval        string                                                                        `json:"collectionInterval,omitempty"`        // Re sync Interval of the device
	SerialNumber              string                                                                        `json:"serialNumber,omitempty"`              // Device Serial Number
	SoftwareVersion           string                                                                        `json:"softwareVersion,omitempty"`           // Device Software Version
	RoleSource                string                                                                        `json:"roleSource,omitempty"`                // Role source as manual / auto
	Hostname                  string                                                                        `json:"hostname,omitempty"`                  // Device Hostname
	UpTime                    string                                                                        `json:"upTime,omitempty"`                    // Device's uptime
	LastUpdateTime            *int                                                                          `json:"lastUpdateTime,omitempty"`            // Time in epoch when the network device info last got updated
	ErrorDescription          string                                                                        `json:"errorDescription,omitempty"`          // Inventory status description
	LocationName              *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsLocationName       `json:"locationName,omitempty"`              // [Deprecated] Name of the associated location
	TagCount                  string                                                                        `json:"tagCount,omitempty"`                  // Number of tags associated with the device
	LastUpdated               string                                                                        `json:"lastUpdated,omitempty"`               // Time when the network device info last got updated
	InstanceUUID              string                                                                        `json:"instanceUuid,omitempty"`              // Instance Uuid of the device
	ID                        string                                                                        `json:"id,omitempty"`                        // Device's UUID
	NeighborTopology          *[]ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopology `json:"neighborTopology,omitempty"`          //
}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsLocation interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsTunnelUDPPort interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsWaasDeviceMode interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsLocationName interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopology struct {
	Nodes *[]ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodes `json:"nodes,omitempty"` //
	Links *[]ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyLinks `json:"links,omitempty"` //
}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodes struct {
	Role            string                                                                                          `json:"role,omitempty"`            // Role of the Node
	Name            string                                                                                          `json:"name,omitempty"`            // Hostname of the Node
	ID              string                                                                                          `json:"id,omitempty"`              // Id of the Node
	Description     string                                                                                          `json:"description,omitempty"`     // Description of the Node
	DeviceType      string                                                                                          `json:"deviceType,omitempty"`      // Device type of the node, like switch, AP, WCL,GateWay
	PlatformID      string                                                                                          `json:"platformId,omitempty"`      // Type of platform
	Family          string                                                                                          `json:"family,omitempty"`          // Device Family of the Node
	IP              string                                                                                          `json:"ip,omitempty"`              // IP Address of the Node
	SoftwareVersion string                                                                                          `json:"softwareVersion,omitempty"` // Software Version of the Node
	UserID          *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesUserID          `json:"userId,omitempty"`          // User Id of the Node
	NodeType        string                                                                                          `json:"nodeType,omitempty"`        // Type of the Node
	RadioFrequency  *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesRadioFrequency  `json:"radioFrequency,omitempty"`  // Frequency of wireless radio channel
	Clients         *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesClients         `json:"clients,omitempty"`         // Number of clients
	Count           *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesCount           `json:"count,omitempty"`           // The number of group nodes (for ap sepecifically)
	HealthScore     *int                                                                                            `json:"healthScore,omitempty"`     // The total health score of the node
	Level           *float64                                                                                        `json:"level,omitempty"`           // The level index to be used by UI widget (starts from 0)
	FabricGroup     *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesFabricGroup     `json:"fabricGroup,omitempty"`     // Fabric device group name
	ConnectedDevice *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesConnectedDevice `json:"connectedDevice,omitempty"` // The connected device to show the connected switch to wlc
}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesUserID interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesRadioFrequency interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesClients interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesCount interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesFabricGroup interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesConnectedDevice interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyLinks struct {
	Source          string                                                                                          `json:"source,omitempty"`          // Edge line starting node
	LinkStatus      string                                                                                          `json:"linkStatus,omitempty"`      // The status of the link (up/down)
	Label           *[]ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyLinksLabel         `json:"label,omitempty"`           // The details of the edge
	Target          string                                                                                          `json:"target,omitempty"`          // End node of the edge line
	ID              *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyLinksID              `json:"id,omitempty"`              // Id of the node
	PortUtilization *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyLinksPortUtilization `json:"portUtilization,omitempty"` // Number of clients
}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyLinksLabel interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyLinksID interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyLinksPortUtilization interface{}
type ResponseDevicesDevices struct {
	Version    string                            `json:"version,omitempty"`    // Response data's version string
	TotalCount *int                              `json:"totalCount,omitempty"` // Total number of devices
	Response   *[]ResponseDevicesDevicesResponse `json:"response,omitempty"`   //
}
type ResponseDevicesDevicesResponse struct {
	DeviceType                 string                                            `json:"deviceType,omitempty"`                 // Device type
	CPUUtilization             *float64                                          `json:"cpuUtilization,omitempty"`             // Device's CPU utilization
	OverallHealth              *int                                              `json:"overallHealth,omitempty"`              // Overall health score
	UtilizationHealth          *ResponseDevicesDevicesResponseUtilizationHealth  `json:"utilizationHealth,omitempty"`          //
	AirQualityHealth           *ResponseDevicesDevicesResponseAirQualityHealth   `json:"airQualityHealth,omitempty"`           //
	IPAddress                  string                                            `json:"ipAddress,omitempty"`                  // Management IP address of the device
	CPUHealth                  *int                                              `json:"cpuHealth,omitempty"`                  // Device CPU health score
	DeviceFamily               string                                            `json:"deviceFamily,omitempty"`               // Device family
	IssueCount                 *int                                              `json:"issueCount,omitempty"`                 // Number of issues
	MacAddress                 string                                            `json:"macAddress,omitempty"`                 // MAC address of the device
	NoiseHealth                *ResponseDevicesDevicesResponseNoiseHealth        `json:"noiseHealth,omitempty"`                //
	OsVersion                  string                                            `json:"osVersion,omitempty"`                  // Device OS version string
	Name                       string                                            `json:"name,omitempty"`                       // Device name
	InterfaceLinkErrHealth     *int                                              `json:"interfaceLinkErrHealth,omitempty"`     // Device (AP) error health score
	MemoryUtilization          *float64                                          `json:"memoryUtilization,omitempty"`          // Device memory utilization
	InterDeviceLinkAvailHealth *int                                              `json:"interDeviceLinkAvailHealth,omitempty"` // Device connectivity status
	InterferenceHealth         *ResponseDevicesDevicesResponseInterferenceHealth `json:"interferenceHealth,omitempty"`         //
	Model                      string                                            `json:"model,omitempty"`                      // Device model string
	Location                   string                                            `json:"location,omitempty"`                   // Site location in which this device is assigned to
	ReachabilityHealth         string                                            `json:"reachabilityHealth,omitempty"`         // Device reachability in the network
	Band                       *ResponseDevicesDevicesResponseBand               `json:"band,omitempty"`                       //
	MemoryUtilizationHealth    *int                                              `json:"memoryUtilizationHealth,omitempty"`    // Device memory utilization health score
	ClientCount                *ResponseDevicesDevicesResponseClientCount        `json:"clientCount,omitempty"`                //
	AvgTemperature             *float64                                          `json:"avgTemperature,omitempty"`             // Average device (switch) temperature
	MaxTemperature             *float64                                          `json:"maxTemperature,omitempty"`             // Max device (switch) temperature
	InterDeviceLinkAvailFabric *int                                              `json:"interDeviceLinkAvailFabric,omitempty"` // Device uplink health
	ApCount                    *int                                              `json:"apCount,omitempty"`                    // Number of AP count
	FreeTimerScore             *int                                              `json:"freeTimerScore,omitempty"`             // Device free timer health score
	FreeTimer                  *float64                                          `json:"freeTimer,omitempty"`                  // Device free timer
	PacketPoolHealth           *int                                              `json:"packetPoolHealth,omitempty"`           // Device packet pool
	PacketPool                 *int                                              `json:"packetPool,omitempty"`                 // Device packet pool
	FreeMemoryBufferHealth     *int                                              `json:"freeMemoryBufferHealth,omitempty"`     // Device free memory buffer health
	FreeMemoryBuffer           *float64                                          `json:"freeMemoryBuffer,omitempty"`           // Device free memory
	WqePoolsHealth             *int                                              `json:"wqePoolsHealth,omitempty"`             // Device WQE pool health
	WqePools                   *float64                                          `json:"wqePools,omitempty"`                   // Device WQE pool
	WanLinkUtilization         *float64                                          `json:"wanLinkUtilization,omitempty"`         // WLAN link utilization
	CPUUlitilization           *float64                                          `json:"cpuUlitilization,omitempty"`           // Device's CPU utilization
	UUID                       string                                            `json:"uuid,omitempty"`                       // Device UUID
}
type ResponseDevicesDevicesResponseUtilizationHealth struct {
	Radio0 *int `json:"radio0,omitempty"` // Radio0
	Radio1 *int `json:"radio1,omitempty"` // Radio1
	Radio2 *int `json:"radio2,omitempty"` // Radio2
	Radio3 *int `json:"radio3,omitempty"` // Radio3
	Ghz24  *int `json:"Ghz24,omitempty"`  // Ghz24
	Ghz50  *int `json:"Ghz50,omitempty"`  // Ghz50
}
type ResponseDevicesDevicesResponseAirQualityHealth struct {
	Radio0 *int `json:"radio0,omitempty"` // Radio0
	Radio1 *int `json:"radio1,omitempty"` // Radio1
	Radio2 *int `json:"radio2,omitempty"` // Radio2
	Radio3 *int `json:"radio3,omitempty"` // Radio3
	Ghz24  *int `json:"Ghz24,omitempty"`  // Ghz24
	Ghz50  *int `json:"Ghz50,omitempty"`  // Ghz50
}
type ResponseDevicesDevicesResponseNoiseHealth struct {
	Radio0 *int `json:"radio0,omitempty"` // Radio0
	Radio1 *int `json:"radio1,omitempty"` // Radio1
	Radio2 *int `json:"radio2,omitempty"` // Radio2
	Radio3 *int `json:"radio3,omitempty"` // Radio3
	Ghz24  *int `json:"Ghz24,omitempty"`  // Ghz24
	Ghz50  *int `json:"Ghz50,omitempty"`  // Ghz50
}
type ResponseDevicesDevicesResponseInterferenceHealth struct {
	Radio0 *int `json:"radio0,omitempty"` // Radio0
	Radio1 *int `json:"radio1,omitempty"` // Radio1
	Radio2 *int `json:"radio2,omitempty"` // Radio2
	Radio3 *int `json:"radio3,omitempty"` // Radio3
	Ghz24  *int `json:"Ghz24,omitempty"`  // Ghz24
	Ghz50  *int `json:"Ghz50,omitempty"`  // Ghz50
}
type ResponseDevicesDevicesResponseBand struct {
	Radio0 string `json:"radio0,omitempty"` // Radio0
	Radio1 string `json:"radio1,omitempty"` // Radio1
	Radio2 string `json:"radio2,omitempty"` // Radio2
	Radio3 *int   `json:"radio3,omitempty"` // Radio3
}
type ResponseDevicesDevicesResponseClientCount struct {
	Radio0 *int `json:"radio0,omitempty"` // Radio0
	Radio1 *int `json:"radio1,omitempty"` // Radio1
	Radio2 *int `json:"radio2,omitempty"` // Radio2
	Radio3 *int `json:"radio3,omitempty"` // Radio3
	Ghz24  *int `json:"Ghz24,omitempty"`  // Ghz24
	Ghz50  *int `json:"Ghz50,omitempty"`  // Ghz50
}
type ResponseDevicesUpdatePlannedAccessPointForFloor struct {
	Version  string                                                   `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseDevicesUpdatePlannedAccessPointForFloorResponse `json:"response,omitempty"` //
}
type ResponseDevicesUpdatePlannedAccessPointForFloorResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseDevicesCreatePlannedAccessPointForFloor struct {
	Version  string                                                   `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseDevicesCreatePlannedAccessPointForFloorResponse `json:"response,omitempty"` //
}
type ResponseDevicesCreatePlannedAccessPointForFloorResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseDevicesGetPlannedAccessPointsForFloor struct {
	Response *[]ResponseDevicesGetPlannedAccessPointsForFloorResponse `json:"response,omitempty"` //
	Version  *int                                                     `json:"version,omitempty"`  // Version of the api response model
	Total    *int                                                     `json:"total,omitempty"`    // Total number of the planned access points
}
type ResponseDevicesGetPlannedAccessPointsForFloorResponse struct {
	Attributes *ResponseDevicesGetPlannedAccessPointsForFloorResponseAttributes `json:"attributes,omitempty"` //
	Location   *ResponseDevicesGetPlannedAccessPointsForFloorResponseLocation   `json:"location,omitempty"`   //
	Position   *ResponseDevicesGetPlannedAccessPointsForFloorResponsePosition   `json:"position,omitempty"`   //
	RadioCount *int                                                             `json:"radioCount,omitempty"` // Number of radios of the planned access point
	Radios     *[]ResponseDevicesGetPlannedAccessPointsForFloorResponseRadios   `json:"radios,omitempty"`     //
	IsSensor   *bool                                                            `json:"isSensor,omitempty"`   // Determines if the planned access point is sensor or not
}
type ResponseDevicesGetPlannedAccessPointsForFloorResponseAttributes struct {
	ID            *float64 `json:"id,omitempty"`            // Unique id of the planned access point
	InstanceUUID  string   `json:"instanceUuid,omitempty"`  // Instance uuid of the planned access point
	Name          string   `json:"name,omitempty"`          // Display name of the planned access point
	TypeString    string   `json:"typeString,omitempty"`    // Type string representation of the planned access point
	Domain        string   `json:"domain,omitempty"`        // Service domain to which the planned access point belongs
	HeirarchyName string   `json:"heirarchyName,omitempty"` // Hierarchy name of the planned access point
	Source        string   `json:"source,omitempty"`        // Source of the data used to create the planned access point
	CreateDate    *float64 `json:"createDate,omitempty"`    // Created date of the planned access point
	MacAddress    string   `json:"macAddress,omitempty"`    // MAC address of the planned access point
}
type ResponseDevicesGetPlannedAccessPointsForFloorResponseLocation struct {
	Altitude   *float64 `json:"altitude,omitempty"`   // Altitude of the planned access point's location
	Lattitude  *float64 `json:"lattitude,omitempty"`  // Latitude of the planned access point's location
	Longtitude *float64 `json:"longtitude,omitempty"` // Longitude of the planned access point's location
}
type ResponseDevicesGetPlannedAccessPointsForFloorResponsePosition struct {
	X *float64 `json:"x,omitempty"` // x-coordinate of the planned access point on the map, 0,0 point being the top-left corner
	Y *float64 `json:"y,omitempty"` // y-coordinate of the planned access point on the map, 0,0 point being the top-left corner
	Z *float64 `json:"z,omitempty"` // z-coordinate, or height, of the planned access point on the map
}
type ResponseDevicesGetPlannedAccessPointsForFloorResponseRadios struct {
	Attributes *ResponseDevicesGetPlannedAccessPointsForFloorResponseRadiosAttributes `json:"attributes,omitempty"` //
	Antenna    *ResponseDevicesGetPlannedAccessPointsForFloorResponseRadiosAntenna    `json:"antenna,omitempty"`    //
	IsSensor   *bool                                                                  `json:"isSensor,omitempty"`   // Determines if it is sensor or not
}
type ResponseDevicesGetPlannedAccessPointsForFloorResponseRadiosAttributes struct {
	ID            *int     `json:"id,omitempty"`            // Id of the radio
	InstanceUUID  string   `json:"instanceUuid,omitempty"`  // Instance Uuid of the radio
	SlotID        *float64 `json:"slotId,omitempty"`        // Slot number in which the radio resides in the parent access point
	IfTypeString  string   `json:"ifTypeString,omitempty"`  // String representation of native band
	IfTypeSubband string   `json:"ifTypeSubband,omitempty"` // Sub band type of the radio
	Channel       *float64 `json:"channel,omitempty"`       // Channel in which the radio operates
	ChannelString string   `json:"channelString,omitempty"` // Channel string representation
	IfMode        string   `json:"ifMode,omitempty"`        // IF mode of the radio
	TxPowerLevel  *float64 `json:"txPowerLevel,omitempty"`  // Tx Power at which this radio operates (in dBm)
}
type ResponseDevicesGetPlannedAccessPointsForFloorResponseRadiosAntenna struct {
	Name           string   `json:"name,omitempty"`           // Name of the antenna
	Type           string   `json:"type,omitempty"`           // Type of the antenna associated with this radio
	Mode           string   `json:"mode,omitempty"`           // Mode of the antenna associated with this radio
	AzimuthAngle   *float64 `json:"azimuthAngle,omitempty"`   // Azimuth angle of the antenna
	ElevationAngle *float64 `json:"elevationAngle,omitempty"` // Elevation angle of the antenna
	Gain           *float64 `json:"gain,omitempty"`           // Gain of the antenna
}
type ResponseDevicesDeletePlannedAccessPointForFloor struct {
	Version  string                                                   `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseDevicesDeletePlannedAccessPointForFloorResponse `json:"response,omitempty"` //
}
type ResponseDevicesDeletePlannedAccessPointForFloorResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseDevicesGetAllHealthScoreDefinitionsForGivenFilters struct {
	Response *[]ResponseDevicesGetAllHealthScoreDefinitionsForGivenFiltersResponse `json:"response,omitempty"` //
}
type ResponseDevicesGetAllHealthScoreDefinitionsForGivenFiltersResponse struct {
	ID                          string   `json:"id,omitempty"`                          // Id
	Name                        string   `json:"name,omitempty"`                        // Name
	DisplayName                 string   `json:"displayName,omitempty"`                 // Display Name
	DeviceFamily                string   `json:"deviceFamily,omitempty"`                // Device Family
	Description                 string   `json:"description,omitempty"`                 // Description
	IncludeForOverallHealth     *bool    `json:"includeForOverallHealth,omitempty"`     // Include For Overall Health
	DefinitionStatus            string   `json:"definitionStatus,omitempty"`            // Definition Status
	ThresholdValue              *float64 `json:"thresholdValue,omitempty"`              // Threshold Value
	SynchronizeToIssueThreshold *bool    `json:"synchronizeToIssueThreshold,omitempty"` // Synchronize To Issue Threshold
	LastModified                string   `json:"lastModified,omitempty"`                // Last Modified
}
type ResponseDevicesUpdateHealthScoreDefinitions struct {
	Response *[]ResponseDevicesUpdateHealthScoreDefinitionsResponse `json:"response,omitempty"` //
	Version  string                                                 `json:"version,omitempty"`  // Version
}
type ResponseDevicesUpdateHealthScoreDefinitionsResponse struct {
	ID                          string   `json:"id,omitempty"`                          // Id
	Name                        string   `json:"name,omitempty"`                        // Name
	DisplayName                 string   `json:"displayName,omitempty"`                 // Display Name
	DeviceFamily                string   `json:"deviceFamily,omitempty"`                // Device Family
	Description                 string   `json:"description,omitempty"`                 // Description
	IncludeForOverallHealth     *bool    `json:"includeForOverallHealth,omitempty"`     // Include For Overall Health
	DefinitionStatus            string   `json:"definitionStatus,omitempty"`            // Definition Status
	ThresholdValue              *float64 `json:"thresholdValue,omitempty"`              // Threshold Value
	SynchronizeToIssueThreshold *bool    `json:"synchronizeToIssueThreshold,omitempty"` // Synchronize To Issue Threshold
	LastModified                string   `json:"lastModified,omitempty"`                // Last Modified
}
type ResponseDevicesGetHealthScoreDefinitionForTheGivenID struct {
	Response *[]ResponseDevicesGetHealthScoreDefinitionForTheGivenIDResponse `json:"response,omitempty"` //
}
type ResponseDevicesGetHealthScoreDefinitionForTheGivenIDResponse struct {
	ID                          string   `json:"id,omitempty"`                          // Id
	Name                        string   `json:"name,omitempty"`                        // Name
	DisplayName                 string   `json:"displayName,omitempty"`                 // Display Name
	DeviceFamily                string   `json:"deviceFamily,omitempty"`                // Device Family
	Description                 string   `json:"description,omitempty"`                 // Description
	IncludeForOverallHealth     *bool    `json:"includeForOverallHealth,omitempty"`     // Include For Overall Health
	DefinitionStatus            string   `json:"definitionStatus,omitempty"`            // Definition Status
	ThresholdValue              *float64 `json:"thresholdValue,omitempty"`              // Threshold Value
	SynchronizeToIssueThreshold *bool    `json:"synchronizeToIssueThreshold,omitempty"` // Synchronize To Issue Threshold
	LastModified                string   `json:"lastModified,omitempty"`                // Last Modified
}
type ResponseDevicesUpdateHealthScoreDefinitionForTheGivenID struct {
	Response *ResponseDevicesUpdateHealthScoreDefinitionForTheGivenIDResponse `json:"response,omitempty"` //
	Version  string                                                           `json:"version,omitempty"`  // Version
}
type ResponseDevicesUpdateHealthScoreDefinitionForTheGivenIDResponse struct {
	ID                          string   `json:"id,omitempty"`                          // Id
	Name                        string   `json:"name,omitempty"`                        // Name
	DisplayName                 string   `json:"displayName,omitempty"`                 // Display Name
	DeviceFamily                string   `json:"deviceFamily,omitempty"`                // Device Family
	Description                 string   `json:"description,omitempty"`                 // Description
	IncludeForOverallHealth     *bool    `json:"includeForOverallHealth,omitempty"`     // Include For Overall Health
	DefinitionStatus            string   `json:"definitionStatus,omitempty"`            // Definition Status
	ThresholdValue              *float64 `json:"thresholdValue,omitempty"`              // Threshold Value
	SynchronizeToIssueThreshold *bool    `json:"synchronizeToIssueThreshold,omitempty"` // Synchronize To Issue Threshold
	LastModified                string   `json:"lastModified,omitempty"`                // Last Modified
}
type ResponseDevicesGetAllInterfaces struct {
	Response *[]ResponseDevicesGetAllInterfacesResponse `json:"response,omitempty"` //
	Version  string                                     `json:"version,omitempty"`  //
}
type ResponseDevicesGetAllInterfacesResponse struct {
	Addresses                   *[]ResponseDevicesGetAllInterfacesResponseAddresses `json:"addresses,omitempty"`                   //
	AdminStatus                 string                                              `json:"adminStatus,omitempty"`                 // Admin status as ('UP'/'DOWN')
	ClassName                   string                                              `json:"className,omitempty"`                   // Classifies the port as switch port ,loopback interface etc.
	Description                 string                                              `json:"description,omitempty"`                 // Description for the Interface
	Name                        string                                              `json:"name,omitempty"`                        // Name for the interface
	DeviceID                    string                                              `json:"deviceId,omitempty"`                    // Device Id of the device
	Duplex                      string                                              `json:"duplex,omitempty"`                      // Interface duplex as AutoNegotiate or FullDuplex
	ID                          string                                              `json:"id,omitempty"`                          // ID of the Interface
	IfIndex                     string                                              `json:"ifIndex,omitempty"`                     // Interface index
	InstanceTenantID            string                                              `json:"instanceTenantId,omitempty"`            // Instance Tenant Id of the Interface
	InstanceUUID                string                                              `json:"instanceUuid,omitempty"`                // Instance Uuid of the Interface
	InterfaceType               string                                              `json:"interfaceType,omitempty"`               // Interface type as Physical or Virtual
	IPv4Address                 string                                              `json:"ipv4Address,omitempty"`                 // IPV4 Address of the device
	IPv4Mask                    string                                              `json:"ipv4Mask,omitempty"`                    // IPV4 Mask of the device
	IsisSupport                 string                                              `json:"isisSupport,omitempty"`                 // Flag for ISIS enabled / disabled
	LastOutgoingPacketTime      *float64                                            `json:"lastOutgoingPacketTime,omitempty"`      // Time, in milliseconds since UNIX epoch, when the last packet was sent from this interface
	LastIncomingPacketTime      *float64                                            `json:"lastIncomingPacketTime,omitempty"`      // Time, in milliseconds since UNIX epoch, when the last packet was received on this interface
	LastUpdated                 string                                              `json:"lastUpdated,omitempty"`                 // Time when the device interface info last got updated
	MacAddress                  string                                              `json:"macAddress,omitempty"`                  // MAC address of interface
	MappedPhysicalInterfaceID   string                                              `json:"mappedPhysicalInterfaceId,omitempty"`   // ID of physical interface mapped with the virtual interface of WLC
	MappedPhysicalInterfaceName string                                              `json:"mappedPhysicalInterfaceName,omitempty"` // Physical interface name mapped with the virtual interface of WLC
	MediaType                   string                                              `json:"mediaType,omitempty"`                   // Media Type of the interface
	Mtu                         string                                              `json:"mtu,omitempty"`                         // MTU Information of Interface
	NativeVLANID                string                                              `json:"nativeVlanId,omitempty"`                // Vlan to receive untagged frames on trunk port
	OspfSupport                 string                                              `json:"ospfSupport,omitempty"`                 // Flag for OSPF enabled / disabled
	Pid                         string                                              `json:"pid,omitempty"`                         // Platform ID of the device
	PortMode                    string                                              `json:"portMode,omitempty"`                    // Port mode as access, trunk, routed
	PortName                    string                                              `json:"portName,omitempty"`                    // Interface name
	PortType                    string                                              `json:"portType,omitempty"`                    // Port type as Ethernet Port / Ethernet SVI / Ethernet Sub Interface
	SerialNo                    string                                              `json:"serialNo,omitempty"`                    // Serial number of the device
	Series                      string                                              `json:"series,omitempty"`                      // Series of the device
	Speed                       string                                              `json:"speed,omitempty"`                       // Speed of the interface
	Status                      string                                              `json:"status,omitempty"`                      // Interface status as Down / Up
	VLANID                      string                                              `json:"vlanId,omitempty"`                      // Vlan ID of interface
	VoiceVLAN                   string                                              `json:"voiceVlan,omitempty"`                   // Vlan information of the interface
}
type ResponseDevicesGetAllInterfacesResponseAddresses struct {
	Address *ResponseDevicesGetAllInterfacesResponseAddressesAddress `json:"address,omitempty"` //
	Type    string                                                   `json:"type,omitempty"`    // Type of the interface. For e.g. IPv4, IPv6 (with unicast, multicast, anycast, etc.)
}
type ResponseDevicesGetAllInterfacesResponseAddressesAddress struct {
	IPAddress     *ResponseDevicesGetAllInterfacesResponseAddressesAddressIPAddress `json:"ipAddress,omitempty"`     //
	IPMask        *ResponseDevicesGetAllInterfacesResponseAddressesAddressIPMask    `json:"ipMask,omitempty"`        //
	IsInverseMask *bool                                                             `json:"isInverseMask,omitempty"` // Inverse Mask of the IP address is enabled or not
}
type ResponseDevicesGetAllInterfacesResponseAddressesAddressIPAddress struct {
	Address string `json:"address,omitempty"` // IP address of the interface
}
type ResponseDevicesGetAllInterfacesResponseAddressesAddressIPMask struct {
	Address string `json:"address,omitempty"` // IP Mask of the interface
}
type ResponseDevicesGetDeviceInterfaceCountForMultipleDevices struct {
	Response *int   `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}
type ResponseDevicesGetInterfaceByIP struct {
	Response *[]ResponseDevicesGetInterfaceByIPResponse `json:"response,omitempty"` //
	Version  string                                     `json:"version,omitempty"`  //
}
type ResponseDevicesGetInterfaceByIPResponse struct {
	Addresses                   *[]ResponseDevicesGetInterfaceByIPResponseAddresses `json:"addresses,omitempty"`                   //
	AdminStatus                 string                                              `json:"adminStatus,omitempty"`                 // Admin status as ('UP'/'DOWN')
	ClassName                   string                                              `json:"className,omitempty"`                   // Classifies the port as switch port ,loopback interface etc.
	Description                 string                                              `json:"description,omitempty"`                 // Description for the Interface
	Name                        string                                              `json:"name,omitempty"`                        // Name for the interface
	DeviceID                    string                                              `json:"deviceId,omitempty"`                    // Device Id of the device
	Duplex                      string                                              `json:"duplex,omitempty"`                      // Interface duplex as AutoNegotiate or FullDuplex
	ID                          string                                              `json:"id,omitempty"`                          // ID of the Interface
	IfIndex                     string                                              `json:"ifIndex,omitempty"`                     // Interface index
	InstanceTenantID            string                                              `json:"instanceTenantId,omitempty"`            // Instance Tenant Id of the Interface
	InstanceUUID                string                                              `json:"instanceUuid,omitempty"`                // Instance Uuid of the Interface
	InterfaceType               string                                              `json:"interfaceType,omitempty"`               // Interface type as Physical or Virtual
	IPv4Address                 string                                              `json:"ipv4Address,omitempty"`                 // IPV4 Address of the device
	IPv4Mask                    string                                              `json:"ipv4Mask,omitempty"`                    // IPV4 Mask of the device
	IsisSupport                 string                                              `json:"isisSupport,omitempty"`                 // Flag for ISIS enabled / disabled
	LastOutgoingPacketTime      *float64                                            `json:"lastOutgoingPacketTime,omitempty"`      // Time, in milliseconds since UNIX epoch, when the last packet was sent from this interface
	LastIncomingPacketTime      *float64                                            `json:"lastIncomingPacketTime,omitempty"`      // Time, in milliseconds since UNIX epoch, when the last packet was received on this interface
	LastUpdated                 string                                              `json:"lastUpdated,omitempty"`                 // Time when the device interface info last got updated
	MacAddress                  string                                              `json:"macAddress,omitempty"`                  // MAC address of interface
	MappedPhysicalInterfaceID   string                                              `json:"mappedPhysicalInterfaceId,omitempty"`   // ID of physical interface mapped with the virtual interface of WLC
	MappedPhysicalInterfaceName string                                              `json:"mappedPhysicalInterfaceName,omitempty"` // Physical interface name mapped with the virtual interface of WLC
	MediaType                   string                                              `json:"mediaType,omitempty"`                   // Media Type of the interface
	Mtu                         string                                              `json:"mtu,omitempty"`                         // MTU Information of Interface
	NativeVLANID                string                                              `json:"nativeVlanId,omitempty"`                // Vlan to receive untagged frames on trunk port
	OspfSupport                 string                                              `json:"ospfSupport,omitempty"`                 // Flag for OSPF enabled / disabled
	Pid                         string                                              `json:"pid,omitempty"`                         // Platform ID of the device
	PortMode                    string                                              `json:"portMode,omitempty"`                    // Port mode as access, trunk, routed
	PortName                    string                                              `json:"portName,omitempty"`                    // Interface name
	PortType                    string                                              `json:"portType,omitempty"`                    // Port type as Ethernet Port / Ethernet SVI / Ethernet Sub Interface
	SerialNo                    string                                              `json:"serialNo,omitempty"`                    // Serial number of the device
	Series                      string                                              `json:"series,omitempty"`                      // Series of the device
	Speed                       string                                              `json:"speed,omitempty"`                       // Speed of the interface
	Status                      string                                              `json:"status,omitempty"`                      // Interface status as Down / Up
	VLANID                      string                                              `json:"vlanId,omitempty"`                      // Vlan ID of interface
	VoiceVLAN                   string                                              `json:"voiceVlan,omitempty"`                   // Vlan information of the interface
}
type ResponseDevicesGetInterfaceByIPResponseAddresses struct {
	Address *ResponseDevicesGetInterfaceByIPResponseAddressesAddress `json:"address,omitempty"` //
	Type    string                                                   `json:"type,omitempty"`    // Type of the interface. For e.g. IPv4, IPv6 (with unicast, multicast, anycast, etc.)
}
type ResponseDevicesGetInterfaceByIPResponseAddressesAddress struct {
	IPAddress     *ResponseDevicesGetInterfaceByIPResponseAddressesAddressIPAddress `json:"ipAddress,omitempty"`     //
	IPMask        *ResponseDevicesGetInterfaceByIPResponseAddressesAddressIPMask    `json:"ipMask,omitempty"`        //
	IsInverseMask *bool                                                             `json:"isInverseMask,omitempty"` // Inverse Mask of the IP address is enabled or not
}
type ResponseDevicesGetInterfaceByIPResponseAddressesAddressIPAddress struct {
	Address string `json:"address,omitempty"` // IP address of the interface
}
type ResponseDevicesGetInterfaceByIPResponseAddressesAddressIPMask struct {
	Address string `json:"address,omitempty"` // IP Mask of the interface
}
type ResponseDevicesGetIsisInterfaces struct {
	Response *[]ResponseDevicesGetIsisInterfacesResponse `json:"response,omitempty"` //
	Version  string                                      `json:"version,omitempty"`  //
}
type ResponseDevicesGetIsisInterfacesResponse struct {
	Addresses                   *[]ResponseDevicesGetIsisInterfacesResponseAddresses `json:"addresses,omitempty"`                   //
	AdminStatus                 string                                               `json:"adminStatus,omitempty"`                 // Admin status as ('UP'/'DOWN')
	ClassName                   string                                               `json:"className,omitempty"`                   // Classifies the port as switch port ,loopback interface etc.
	Description                 string                                               `json:"description,omitempty"`                 // Description for the Interface
	Name                        string                                               `json:"name,omitempty"`                        // Name for the interface
	DeviceID                    string                                               `json:"deviceId,omitempty"`                    // Device Id of the device
	Duplex                      string                                               `json:"duplex,omitempty"`                      // Interface duplex as AutoNegotiate or FullDuplex
	ID                          string                                               `json:"id,omitempty"`                          // ID of the Interface
	IfIndex                     string                                               `json:"ifIndex,omitempty"`                     // Interface index
	InstanceTenantID            string                                               `json:"instanceTenantId,omitempty"`            // Instance Tenant Id of the Interface
	InstanceUUID                string                                               `json:"instanceUuid,omitempty"`                // Instance Uuid of the Interface
	InterfaceType               string                                               `json:"interfaceType,omitempty"`               // Interface type as Physical or Virtual
	IPv4Address                 string                                               `json:"ipv4Address,omitempty"`                 // IPV4 Address of the device
	IPv4Mask                    string                                               `json:"ipv4Mask,omitempty"`                    // IPV4 Mask of the device
	IsisSupport                 string                                               `json:"isisSupport,omitempty"`                 // Flag for ISIS enabled / disabled
	LastOutgoingPacketTime      *float64                                             `json:"lastOutgoingPacketTime,omitempty"`      // Time, in milliseconds since UNIX epoch, when the last packet was sent from this interface
	LastIncomingPacketTime      *float64                                             `json:"lastIncomingPacketTime,omitempty"`      // Time, in milliseconds since UNIX epoch, when the last packet was received on this interface
	LastUpdated                 string                                               `json:"lastUpdated,omitempty"`                 // Time when the device interface info last got updated
	MacAddress                  string                                               `json:"macAddress,omitempty"`                  // MAC address of interface
	MappedPhysicalInterfaceID   string                                               `json:"mappedPhysicalInterfaceId,omitempty"`   // ID of physical interface mapped with the virtual interface of WLC
	MappedPhysicalInterfaceName string                                               `json:"mappedPhysicalInterfaceName,omitempty"` // Physical interface name mapped with the virtual interface of WLC
	MediaType                   string                                               `json:"mediaType,omitempty"`                   // Media Type of the interface
	Mtu                         string                                               `json:"mtu,omitempty"`                         // MTU Information of Interface
	NativeVLANID                string                                               `json:"nativeVlanId,omitempty"`                // Vlan to receive untagged frames on trunk port
	OspfSupport                 string                                               `json:"ospfSupport,omitempty"`                 // Flag for OSPF enabled / disabled
	Pid                         string                                               `json:"pid,omitempty"`                         // Platform ID of the device
	PortMode                    string                                               `json:"portMode,omitempty"`                    // Port mode as access, trunk, routed
	PortName                    string                                               `json:"portName,omitempty"`                    // Interface name
	PortType                    string                                               `json:"portType,omitempty"`                    // Port type as Ethernet Port / Ethernet SVI / Ethernet Sub Interface
	SerialNo                    string                                               `json:"serialNo,omitempty"`                    // Serial number of the device
	Series                      string                                               `json:"series,omitempty"`                      // Series of the device
	Speed                       string                                               `json:"speed,omitempty"`                       // Speed of the interface
	Status                      string                                               `json:"status,omitempty"`                      // Interface status as Down / Up
	VLANID                      string                                               `json:"vlanId,omitempty"`                      // Vlan ID of interface
	VoiceVLAN                   string                                               `json:"voiceVlan,omitempty"`                   // Vlan information of the interface
}
type ResponseDevicesGetIsisInterfacesResponseAddresses struct {
	Address *ResponseDevicesGetIsisInterfacesResponseAddressesAddress `json:"address,omitempty"` //
	Type    string                                                    `json:"type,omitempty"`    // Type of the interface. For e.g. IPv4, IPv6 (with unicast, multicast, anycast, etc.)
}
type ResponseDevicesGetIsisInterfacesResponseAddressesAddress struct {
	IPAddress     *ResponseDevicesGetIsisInterfacesResponseAddressesAddressIPAddress `json:"ipAddress,omitempty"`     //
	IPMask        *ResponseDevicesGetIsisInterfacesResponseAddressesAddressIPMask    `json:"ipMask,omitempty"`        //
	IsInverseMask *bool                                                              `json:"isInverseMask,omitempty"` // Inverse Mask of the IP address is enabled or not
}
type ResponseDevicesGetIsisInterfacesResponseAddressesAddressIPAddress struct {
	Address string `json:"address,omitempty"` // IP address of the interface
}
type ResponseDevicesGetIsisInterfacesResponseAddressesAddressIPMask struct {
	Address string `json:"address,omitempty"` // IP Mask of the interface
}
type ResponseDevicesGetInterfaceInfoByID struct {
	Response *[]ResponseDevicesGetInterfaceInfoByIDResponse `json:"response,omitempty"` //
	Version  string                                         `json:"version,omitempty"`  //
}
type ResponseDevicesGetInterfaceInfoByIDResponse struct {
	Addresses                   *[]ResponseDevicesGetInterfaceInfoByIDResponseAddresses `json:"addresses,omitempty"`                   //
	AdminStatus                 string                                                  `json:"adminStatus,omitempty"`                 // Admin status as ('UP'/'DOWN')
	ClassName                   string                                                  `json:"className,omitempty"`                   // Classifies the port as switch port ,loopback interface etc.
	Description                 string                                                  `json:"description,omitempty"`                 // Description for the Interface
	Name                        string                                                  `json:"name,omitempty"`                        // Name for the interface
	DeviceID                    string                                                  `json:"deviceId,omitempty"`                    // Device Id of the device
	Duplex                      string                                                  `json:"duplex,omitempty"`                      // Interface duplex as AutoNegotiate or FullDuplex
	ID                          string                                                  `json:"id,omitempty"`                          // ID of the Interface
	IfIndex                     string                                                  `json:"ifIndex,omitempty"`                     // Interface index
	InstanceTenantID            string                                                  `json:"instanceTenantId,omitempty"`            // Instance Tenant Id of the Interface
	InstanceUUID                string                                                  `json:"instanceUuid,omitempty"`                // Instance Uuid of the Interface
	InterfaceType               string                                                  `json:"interfaceType,omitempty"`               // Interface type as Physical or Virtual
	IPv4Address                 string                                                  `json:"ipv4Address,omitempty"`                 // IPV4 Address of the device
	IPv4Mask                    string                                                  `json:"ipv4Mask,omitempty"`                    // IPV4 Mask of the device
	IsisSupport                 string                                                  `json:"isisSupport,omitempty"`                 // Flag for ISIS enabled / disabled
	LastOutgoingPacketTime      *float64                                                `json:"lastOutgoingPacketTime,omitempty"`      // Time, in milliseconds since UNIX epoch, when the last packet was sent from this interface
	LastIncomingPacketTime      *float64                                                `json:"lastIncomingPacketTime,omitempty"`      // Time, in milliseconds since UNIX epoch, when the last packet was received on this interface
	LastUpdated                 string                                                  `json:"lastUpdated,omitempty"`                 // Time when the device interface info last got updated
	MacAddress                  string                                                  `json:"macAddress,omitempty"`                  // MAC address of interface
	MappedPhysicalInterfaceID   string                                                  `json:"mappedPhysicalInterfaceId,omitempty"`   // ID of physical interface mapped with the virtual interface of WLC
	MappedPhysicalInterfaceName string                                                  `json:"mappedPhysicalInterfaceName,omitempty"` // Physical interface name mapped with the virtual interface of WLC
	MediaType                   string                                                  `json:"mediaType,omitempty"`                   // Media Type of the interface
	Mtu                         string                                                  `json:"mtu,omitempty"`                         // MTU Information of Interface
	NativeVLANID                string                                                  `json:"nativeVlanId,omitempty"`                // Vlan to receive untagged frames on trunk port
	OspfSupport                 string                                                  `json:"ospfSupport,omitempty"`                 // Flag for OSPF enabled / disabled
	Pid                         string                                                  `json:"pid,omitempty"`                         // Platform ID of the device
	PortMode                    string                                                  `json:"portMode,omitempty"`                    // Port mode as access, trunk, routed
	PortName                    string                                                  `json:"portName,omitempty"`                    // Interface name
	PortType                    string                                                  `json:"portType,omitempty"`                    // Port type as Ethernet Port / Ethernet SVI / Ethernet Sub Interface
	SerialNo                    string                                                  `json:"serialNo,omitempty"`                    // Serial number of the device
	Series                      string                                                  `json:"series,omitempty"`                      // Series of the device
	Speed                       string                                                  `json:"speed,omitempty"`                       // Speed of the interface
	Status                      string                                                  `json:"status,omitempty"`                      // Interface status as Down / Up
	VLANID                      string                                                  `json:"vlanId,omitempty"`                      // Vlan ID of interface
	VoiceVLAN                   string                                                  `json:"voiceVlan,omitempty"`                   // Vlan information of the interface
}
type ResponseDevicesGetInterfaceInfoByIDResponseAddresses struct {
	Address *ResponseDevicesGetInterfaceInfoByIDResponseAddressesAddress `json:"address,omitempty"` //
	Type    string                                                       `json:"type,omitempty"`    // Type of the interface. For e.g. IPv4, IPv6 (with unicast, multicast, anycast, etc.)
}
type ResponseDevicesGetInterfaceInfoByIDResponseAddressesAddress struct {
	IPAddress     *ResponseDevicesGetInterfaceInfoByIDResponseAddressesAddressIPAddress `json:"ipAddress,omitempty"`     //
	IPMask        *ResponseDevicesGetInterfaceInfoByIDResponseAddressesAddressIPMask    `json:"ipMask,omitempty"`        //
	IsInverseMask *bool                                                                 `json:"isInverseMask,omitempty"` // Inverse Mask of the IP address is enabled or not
}
type ResponseDevicesGetInterfaceInfoByIDResponseAddressesAddressIPAddress struct {
	Address string `json:"address,omitempty"` // IP address of the interface
}
type ResponseDevicesGetInterfaceInfoByIDResponseAddressesAddressIPMask struct {
	Address string `json:"address,omitempty"` // IP Mask of the interface
}
type ResponseDevicesGetDeviceInterfaceCount struct {
	Response *int   `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}
type ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceName struct {
	Response *ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameResponse `json:"response,omitempty"` //
	Version  string                                                                `json:"version,omitempty"`  //
}
type ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameResponse struct {
	Addresses                   *[]ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameResponseAddresses `json:"addresses,omitempty"`                   //
	AdminStatus                 string                                                                           `json:"adminStatus,omitempty"`                 // Admin status as ('UP'/'DOWN')
	ClassName                   string                                                                           `json:"className,omitempty"`                   // Classifies the port as switch port ,loopback interface etc.
	Description                 string                                                                           `json:"description,omitempty"`                 // Description for the Interface
	Name                        string                                                                           `json:"name,omitempty"`                        // Name for the interface
	DeviceID                    string                                                                           `json:"deviceId,omitempty"`                    // Device Id of the device
	Duplex                      string                                                                           `json:"duplex,omitempty"`                      // Interface duplex as AutoNegotiate or FullDuplex
	ID                          string                                                                           `json:"id,omitempty"`                          // ID of the Interface
	IfIndex                     string                                                                           `json:"ifIndex,omitempty"`                     // Interface index
	InstanceTenantID            string                                                                           `json:"instanceTenantId,omitempty"`            // Instance Tenant Id of the Interface
	InstanceUUID                string                                                                           `json:"instanceUuid,omitempty"`                // Instance Uuid of the Interface
	InterfaceType               string                                                                           `json:"interfaceType,omitempty"`               // Interface type as Physical or Virtual
	IPv4Address                 string                                                                           `json:"ipv4Address,omitempty"`                 // IPV4 Address of the device
	IPv4Mask                    string                                                                           `json:"ipv4Mask,omitempty"`                    // IPV4 Mask of the device
	IsisSupport                 string                                                                           `json:"isisSupport,omitempty"`                 // Flag for ISIS enabled / disabled
	LastOutgoingPacketTime      *float64                                                                         `json:"lastOutgoingPacketTime,omitempty"`      // Time, in milliseconds since UNIX epoch, when the last packet was sent from this interface
	LastIncomingPacketTime      *float64                                                                         `json:"lastIncomingPacketTime,omitempty"`      // Time, in milliseconds since UNIX epoch, when the last packet was received on this interface
	LastUpdated                 string                                                                           `json:"lastUpdated,omitempty"`                 // Time when the device interface info last got updated
	MacAddress                  string                                                                           `json:"macAddress,omitempty"`                  // MAC address of interface
	MappedPhysicalInterfaceID   string                                                                           `json:"mappedPhysicalInterfaceId,omitempty"`   // ID of physical interface mapped with the virtual interface of WLC
	MappedPhysicalInterfaceName string                                                                           `json:"mappedPhysicalInterfaceName,omitempty"` // Physical interface name mapped with the virtual interface of WLC
	MediaType                   string                                                                           `json:"mediaType,omitempty"`                   // Media Type of the interface
	Mtu                         string                                                                           `json:"mtu,omitempty"`                         // MTU Information of Interface
	NativeVLANID                string                                                                           `json:"nativeVlanId,omitempty"`                // Vlan to receive untagged frames on trunk port
	OspfSupport                 string                                                                           `json:"ospfSupport,omitempty"`                 // Flag for OSPF enabled / disabled
	Pid                         string                                                                           `json:"pid,omitempty"`                         // Platform ID of the device
	PortMode                    string                                                                           `json:"portMode,omitempty"`                    // Port mode as access, trunk, routed
	PortName                    string                                                                           `json:"portName,omitempty"`                    // Interface name
	PortType                    string                                                                           `json:"portType,omitempty"`                    // Port type as Ethernet Port / Ethernet SVI / Ethernet Sub Interface
	SerialNo                    string                                                                           `json:"serialNo,omitempty"`                    // Serial number of the device
	Series                      string                                                                           `json:"series,omitempty"`                      // Series of the device
	Speed                       string                                                                           `json:"speed,omitempty"`                       // Speed of the interface
	Status                      string                                                                           `json:"status,omitempty"`                      // Interface status as Down / Up
	VLANID                      string                                                                           `json:"vlanId,omitempty"`                      // Vlan ID of interface
	VoiceVLAN                   string                                                                           `json:"voiceVlan,omitempty"`                   // Vlan information of the interface
	Poweroverethernet           string                                                                           `json:"poweroverethernet,omitempty"`           // This is internal attribute.  Not to be used.  Deprecated
	NetworkdeviceID             string                                                                           `json:"networkdevice_id,omitempty"`            // This is internal attribute.  Not to be used.  Deprecated
	ManagedComputeElement       string                                                                           `json:"managedComputeElement,omitempty"`       // This is internal attribute.  Not to be used.  Deprecated
	ManagedNetworkElement       string                                                                           `json:"managedNetworkElement,omitempty"`       // This is internal attribute.  Not to be used.  Deprecated
	ManagedNetworkElementURL    string                                                                           `json:"managedNetworkElementUrl,omitempty"`    // This is internal attribute.  Not to be used.  Deprecated
	ManagedComputeElementURL    string                                                                           `json:"managedComputeElementUrl,omitempty"`    // This is internal attribute.  Not to be used.  Deprecated
}
type ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameResponseAddresses struct {
	Address *ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameResponseAddressesAddress `json:"address,omitempty"` //
	Type    string                                                                                `json:"type,omitempty"`    // Type of the interface. For e.g. IPv4, IPv6 (with unicast, multicast, anycast, etc.)
}
type ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameResponseAddressesAddress struct {
	IPAddress     *ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameResponseAddressesAddressIPAddress `json:"ipAddress,omitempty"`     //
	IPMask        *ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameResponseAddressesAddressIPMask    `json:"ipMask,omitempty"`        //
	IsInverseMask *bool                                                                                          `json:"isInverseMask,omitempty"` // Inverse Mask of the IP address is enabled or not
}
type ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameResponseAddressesAddressIPAddress struct {
	Address string `json:"address,omitempty"` // IP address of the interface
}
type ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameResponseAddressesAddressIPMask struct {
	Address string `json:"address,omitempty"` // IP Mask of the interface
}
type ResponseDevicesGetDeviceInterfacesBySpecifiedRange struct {
	Response *[]ResponseDevicesGetDeviceInterfacesBySpecifiedRangeResponse `json:"response,omitempty"` //
	Version  string                                                        `json:"version,omitempty"`  //
}
type ResponseDevicesGetDeviceInterfacesBySpecifiedRangeResponse struct {
	Addresses                   *[]ResponseDevicesGetDeviceInterfacesBySpecifiedRangeResponseAddresses `json:"addresses,omitempty"`                   //
	AdminStatus                 string                                                                 `json:"adminStatus,omitempty"`                 // Admin status as ('UP'/'DOWN')
	ClassName                   string                                                                 `json:"className,omitempty"`                   // Classifies the port as switch port ,loopback interface etc.
	Description                 string                                                                 `json:"description,omitempty"`                 // Description for the Interface
	Name                        string                                                                 `json:"name,omitempty"`                        // Name for the interface
	DeviceID                    string                                                                 `json:"deviceId,omitempty"`                    // Device Id of the device
	Duplex                      string                                                                 `json:"duplex,omitempty"`                      // Interface duplex as AutoNegotiate or FullDuplex
	ID                          string                                                                 `json:"id,omitempty"`                          // ID of the Interface
	IfIndex                     string                                                                 `json:"ifIndex,omitempty"`                     // Interface index
	InstanceTenantID            string                                                                 `json:"instanceTenantId,omitempty"`            // Instance Tenant Id of the Interface
	InstanceUUID                string                                                                 `json:"instanceUuid,omitempty"`                // Instance Uuid of the Interface
	InterfaceType               string                                                                 `json:"interfaceType,omitempty"`               // Interface type as Physical or Virtual
	IPv4Address                 string                                                                 `json:"ipv4Address,omitempty"`                 // IPV4 Address of the device
	IPv4Mask                    string                                                                 `json:"ipv4Mask,omitempty"`                    // IPV4 Mask of the device
	IsisSupport                 string                                                                 `json:"isisSupport,omitempty"`                 // Flag for ISIS enabled / disabled
	LastOutgoingPacketTime      *float64                                                               `json:"lastOutgoingPacketTime,omitempty"`      // Time, in milliseconds since UNIX epoch, when the last packet was sent from this interface
	LastIncomingPacketTime      *float64                                                               `json:"lastIncomingPacketTime,omitempty"`      // Time, in milliseconds since UNIX epoch, when the last packet was received on this interface
	LastUpdated                 string                                                                 `json:"lastUpdated,omitempty"`                 // Time when the device interface info last got updated
	MacAddress                  string                                                                 `json:"macAddress,omitempty"`                  // MAC address of interface
	MappedPhysicalInterfaceID   string                                                                 `json:"mappedPhysicalInterfaceId,omitempty"`   // ID of physical interface mapped with the virtual interface of WLC
	MappedPhysicalInterfaceName string                                                                 `json:"mappedPhysicalInterfaceName,omitempty"` // Physical interface name mapped with the virtual interface of WLC
	MediaType                   string                                                                 `json:"mediaType,omitempty"`                   // Media Type of the interface
	Mtu                         string                                                                 `json:"mtu,omitempty"`                         // MTU Information of Interface
	NativeVLANID                string                                                                 `json:"nativeVlanId,omitempty"`                // Vlan to receive untagged frames on trunk port
	OspfSupport                 string                                                                 `json:"ospfSupport,omitempty"`                 // Flag for OSPF enabled / disabled
	Pid                         string                                                                 `json:"pid,omitempty"`                         // Platform ID of the device
	PortMode                    string                                                                 `json:"portMode,omitempty"`                    // Port mode as access, trunk, routed
	PortName                    string                                                                 `json:"portName,omitempty"`                    // Interface name
	PortType                    string                                                                 `json:"portType,omitempty"`                    // Port type as Ethernet Port / Ethernet SVI / Ethernet Sub Interface
	SerialNo                    string                                                                 `json:"serialNo,omitempty"`                    // Serial number of the device
	Series                      string                                                                 `json:"series,omitempty"`                      // Series of the device
	Speed                       string                                                                 `json:"speed,omitempty"`                       // Speed of the interface
	Status                      string                                                                 `json:"status,omitempty"`                      // Interface status as Down / Up
	VLANID                      string                                                                 `json:"vlanId,omitempty"`                      // Vlan ID of interface
	VoiceVLAN                   string                                                                 `json:"voiceVlan,omitempty"`                   // Vlan information of the interface
}
type ResponseDevicesGetDeviceInterfacesBySpecifiedRangeResponseAddresses struct {
	Address *ResponseDevicesGetDeviceInterfacesBySpecifiedRangeResponseAddressesAddress `json:"address,omitempty"` //
	Type    string                                                                      `json:"type,omitempty"`    // Type of the interface. For e.g. IPv4, IPv6 (with unicast, multicast, anycast, etc.)
}
type ResponseDevicesGetDeviceInterfacesBySpecifiedRangeResponseAddressesAddress struct {
	IPAddress     *ResponseDevicesGetDeviceInterfacesBySpecifiedRangeResponseAddressesAddressIPAddress `json:"ipAddress,omitempty"`     //
	IPMask        *ResponseDevicesGetDeviceInterfacesBySpecifiedRangeResponseAddressesAddressIPMask    `json:"ipMask,omitempty"`        //
	IsInverseMask *bool                                                                                `json:"isInverseMask,omitempty"` // Inverse Mask of the IP address is enabled or not
}
type ResponseDevicesGetDeviceInterfacesBySpecifiedRangeResponseAddressesAddressIPAddress struct {
	Address string `json:"address,omitempty"` // IP address of the interface
}
type ResponseDevicesGetDeviceInterfacesBySpecifiedRangeResponseAddressesAddressIPMask struct {
	Address string `json:"address,omitempty"` // IP Mask of the interface
}
type ResponseDevicesGetOspfInterfaces struct {
	Response *[]ResponseDevicesGetOspfInterfacesResponse `json:"response,omitempty"` //
	Version  string                                      `json:"version,omitempty"`  //
}
type ResponseDevicesGetOspfInterfacesResponse struct {
	Addresses                   *[]ResponseDevicesGetOspfInterfacesResponseAddresses `json:"addresses,omitempty"`                   //
	AdminStatus                 string                                               `json:"adminStatus,omitempty"`                 // Admin status as ('UP'/'DOWN')
	ClassName                   string                                               `json:"className,omitempty"`                   // Classifies the port as switch port ,loopback interface etc.
	Description                 string                                               `json:"description,omitempty"`                 // Description for the Interface
	Name                        string                                               `json:"name,omitempty"`                        // Name for the interface
	DeviceID                    string                                               `json:"deviceId,omitempty"`                    // Device Id of the device
	Duplex                      string                                               `json:"duplex,omitempty"`                      // Interface duplex as AutoNegotiate or FullDuplex
	ID                          string                                               `json:"id,omitempty"`                          // ID of the Interface
	IfIndex                     string                                               `json:"ifIndex,omitempty"`                     // Interface index
	InstanceTenantID            string                                               `json:"instanceTenantId,omitempty"`            // Instance Tenant Id of the Interface
	InstanceUUID                string                                               `json:"instanceUuid,omitempty"`                // Instance Uuid of the Interface
	InterfaceType               string                                               `json:"interfaceType,omitempty"`               // Interface type as Physical or Virtual
	IPv4Address                 string                                               `json:"ipv4Address,omitempty"`                 // IPV4 Address of the device
	IPv4Mask                    string                                               `json:"ipv4Mask,omitempty"`                    // IPV4 Mask of the device
	IsisSupport                 string                                               `json:"isisSupport,omitempty"`                 // Flag for ISIS enabled / disabled
	LastOutgoingPacketTime      *float64                                             `json:"lastOutgoingPacketTime,omitempty"`      // Time, in milliseconds since UNIX epoch, when the last packet was sent from this interface
	LastIncomingPacketTime      *float64                                             `json:"lastIncomingPacketTime,omitempty"`      // Time, in milliseconds since UNIX epoch, when the last packet was received on this interface
	LastUpdated                 string                                               `json:"lastUpdated,omitempty"`                 // Time when the device interface info last got updated
	MacAddress                  string                                               `json:"macAddress,omitempty"`                  // MAC address of interface
	MappedPhysicalInterfaceID   string                                               `json:"mappedPhysicalInterfaceId,omitempty"`   // ID of physical interface mapped with the virtual interface of WLC
	MappedPhysicalInterfaceName string                                               `json:"mappedPhysicalInterfaceName,omitempty"` // Physical interface name mapped with the virtual interface of WLC
	MediaType                   string                                               `json:"mediaType,omitempty"`                   // Media Type of the interface
	Mtu                         string                                               `json:"mtu,omitempty"`                         // MTU Information of Interface
	NativeVLANID                string                                               `json:"nativeVlanId,omitempty"`                // Vlan to receive untagged frames on trunk port
	OspfSupport                 string                                               `json:"ospfSupport,omitempty"`                 // Flag for OSPF enabled / disabled
	Pid                         string                                               `json:"pid,omitempty"`                         // Platform ID of the device
	PortMode                    string                                               `json:"portMode,omitempty"`                    // Port mode as access, trunk, routed
	PortName                    string                                               `json:"portName,omitempty"`                    // Interface name
	PortType                    string                                               `json:"portType,omitempty"`                    // Port type as Ethernet Port / Ethernet SVI / Ethernet Sub Interface
	SerialNo                    string                                               `json:"serialNo,omitempty"`                    // Serial number of the device
	Series                      string                                               `json:"series,omitempty"`                      // Series of the device
	Speed                       string                                               `json:"speed,omitempty"`                       // Speed of the interface
	Status                      string                                               `json:"status,omitempty"`                      // Interface status as Down / Up
	VLANID                      string                                               `json:"vlanId,omitempty"`                      // Vlan ID of interface
	VoiceVLAN                   string                                               `json:"voiceVlan,omitempty"`                   // Vlan information of the interface
}
type ResponseDevicesGetOspfInterfacesResponseAddresses struct {
	Address *ResponseDevicesGetOspfInterfacesResponseAddressesAddress `json:"address,omitempty"` //
	Type    string                                                    `json:"type,omitempty"`    // Type of the interface. For e.g. IPv4, IPv6 (with unicast, multicast, anycast, etc.)
}
type ResponseDevicesGetOspfInterfacesResponseAddressesAddress struct {
	IPAddress     *ResponseDevicesGetOspfInterfacesResponseAddressesAddressIPAddress `json:"ipAddress,omitempty"`     //
	IPMask        *ResponseDevicesGetOspfInterfacesResponseAddressesAddressIPMask    `json:"ipMask,omitempty"`        //
	IsInverseMask *bool                                                              `json:"isInverseMask,omitempty"` // Inverse Mask of the IP address is enabled or not
}
type ResponseDevicesGetOspfInterfacesResponseAddressesAddressIPAddress struct {
	Address string `json:"address,omitempty"` // IP address of the interface
}
type ResponseDevicesGetOspfInterfacesResponseAddressesAddressIPMask struct {
	Address string `json:"address,omitempty"` // IP Mask of the interface
}
type ResponseDevicesGetInterfaceByID struct {
	Response *ResponseDevicesGetInterfaceByIDResponse `json:"response,omitempty"` //
	Version  string                                   `json:"version,omitempty"`  //
}
type ResponseDevicesGetInterfaceByIDResponse struct {
	Addresses                   *[]ResponseDevicesGetInterfaceByIDResponseAddresses `json:"addresses,omitempty"`                   //
	AdminStatus                 string                                              `json:"adminStatus,omitempty"`                 // Admin status as ('UP'/'DOWN')
	ClassName                   string                                              `json:"className,omitempty"`                   // Classifies the port as switch port ,loopback interface etc.
	Description                 string                                              `json:"description,omitempty"`                 // Description for the Interface
	Name                        string                                              `json:"name,omitempty"`                        // Name for the interface
	DeviceID                    string                                              `json:"deviceId,omitempty"`                    // Device Id of the device
	Duplex                      string                                              `json:"duplex,omitempty"`                      // Interface duplex as AutoNegotiate or FullDuplex
	ID                          string                                              `json:"id,omitempty"`                          // ID of the Interface
	IfIndex                     string                                              `json:"ifIndex,omitempty"`                     // Interface index
	InstanceTenantID            string                                              `json:"instanceTenantId,omitempty"`            // Instance Tenant Id of the Interface
	InstanceUUID                string                                              `json:"instanceUuid,omitempty"`                // Instance Uuid of the Interface
	InterfaceType               string                                              `json:"interfaceType,omitempty"`               // Interface type as Physical or Virtual
	IPv4Address                 string                                              `json:"ipv4Address,omitempty"`                 // IPV4 Address of the device
	IPv4Mask                    string                                              `json:"ipv4Mask,omitempty"`                    // IPV4 Mask of the device
	IsisSupport                 string                                              `json:"isisSupport,omitempty"`                 // Flag for ISIS enabled / disabled
	LastOutgoingPacketTime      *float64                                            `json:"lastOutgoingPacketTime,omitempty"`      // Time, in milliseconds since UNIX epoch, when the last packet was sent from this interface
	LastIncomingPacketTime      *float64                                            `json:"lastIncomingPacketTime,omitempty"`      // Time, in milliseconds since UNIX epoch, when the last packet was received on this interface
	LastUpdated                 string                                              `json:"lastUpdated,omitempty"`                 // Time when the device interface info last got updated
	MacAddress                  string                                              `json:"macAddress,omitempty"`                  // MAC address of interface
	MappedPhysicalInterfaceID   string                                              `json:"mappedPhysicalInterfaceId,omitempty"`   // ID of physical interface mapped with the virtual interface of WLC
	MappedPhysicalInterfaceName string                                              `json:"mappedPhysicalInterfaceName,omitempty"` // Physical interface name mapped with the virtual interface of WLC
	MediaType                   string                                              `json:"mediaType,omitempty"`                   // Media Type of the interface
	Mtu                         string                                              `json:"mtu,omitempty"`                         // MTU Information of Interface
	NativeVLANID                string                                              `json:"nativeVlanId,omitempty"`                // Vlan to receive untagged frames on trunk port
	OspfSupport                 string                                              `json:"ospfSupport,omitempty"`                 // Flag for OSPF enabled / disabled
	Pid                         string                                              `json:"pid,omitempty"`                         // Platform ID of the device
	PortMode                    string                                              `json:"portMode,omitempty"`                    // Port mode as access, trunk, routed
	PortName                    string                                              `json:"portName,omitempty"`                    // Interface name
	PortType                    string                                              `json:"portType,omitempty"`                    // Port type as Ethernet Port / Ethernet SVI / Ethernet Sub Interface
	SerialNo                    string                                              `json:"serialNo,omitempty"`                    // Serial number of the device
	Series                      string                                              `json:"series,omitempty"`                      // Series of the device
	Speed                       string                                              `json:"speed,omitempty"`                       // Speed of the interface
	Status                      string                                              `json:"status,omitempty"`                      // Interface status as Down / Up
	VLANID                      string                                              `json:"vlanId,omitempty"`                      // Vlan ID of interface
	VoiceVLAN                   string                                              `json:"voiceVlan,omitempty"`                   // Vlan information of the interface
}
type ResponseDevicesGetInterfaceByIDResponseAddresses struct {
	Address *ResponseDevicesGetInterfaceByIDResponseAddressesAddress `json:"address,omitempty"` //
	Type    string                                                   `json:"type,omitempty"`    // Type of the interface. For e.g. IPv4, IPv6 (with unicast, multicast, anycast, etc.)
}
type ResponseDevicesGetInterfaceByIDResponseAddressesAddress struct {
	IPAddress     *ResponseDevicesGetInterfaceByIDResponseAddressesAddressIPAddress `json:"ipAddress,omitempty"`     //
	IPMask        *ResponseDevicesGetInterfaceByIDResponseAddressesAddressIPMask    `json:"ipMask,omitempty"`        //
	IsInverseMask *bool                                                             `json:"isInverseMask,omitempty"` // Inverse Mask of the IP address is enabled or not
}
type ResponseDevicesGetInterfaceByIDResponseAddressesAddressIPAddress struct {
	Address string `json:"address,omitempty"` // IP address of the interface
}
type ResponseDevicesGetInterfaceByIDResponseAddressesAddressIPMask struct {
	Address string `json:"address,omitempty"` // IP Mask of the interface
}
type ResponseDevicesUpdateInterfaceDetails struct {
	Response *ResponseDevicesUpdateInterfaceDetailsResponse `json:"response,omitempty"` //
	Version  *ResponseDevicesUpdateInterfaceDetailsVersion  `json:"version,omitempty"`  //
}
type ResponseDevicesUpdateInterfaceDetailsResponse struct {
	Type       string                                                   `json:"type,omitempty"`       // Type
	Properties *ResponseDevicesUpdateInterfaceDetailsResponseProperties `json:"properties,omitempty"` //
	Required   []string                                                 `json:"required,omitempty"`   // Required
}
type ResponseDevicesUpdateInterfaceDetailsResponseProperties struct {
	TaskID *ResponseDevicesUpdateInterfaceDetailsResponsePropertiesTaskID `json:"taskId,omitempty"` //
	URL    *ResponseDevicesUpdateInterfaceDetailsResponsePropertiesURL    `json:"url,omitempty"`    //
}
type ResponseDevicesUpdateInterfaceDetailsResponsePropertiesTaskID struct {
	Type string `json:"type,omitempty"` // Type
}
type ResponseDevicesUpdateInterfaceDetailsResponsePropertiesURL struct {
	Type string `json:"type,omitempty"` // Type
}
type ResponseDevicesUpdateInterfaceDetailsVersion struct {
	Type string `json:"type,omitempty"` // Type
}
type ResponseDevicesLegitOperationsForInterface struct {
	Response *ResponseDevicesLegitOperationsForInterfaceResponse `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  // Version
}
type ResponseDevicesLegitOperationsForInterfaceResponse struct {
	InterfaceUUID string                                                          `json:"interfaceUuid,omitempty"` // Id of the Interface
	Properties    *[]ResponseDevicesLegitOperationsForInterfaceResponseProperties `json:"properties,omitempty"`    //
	Operations    *[]ResponseDevicesLegitOperationsForInterfaceResponseOperations `json:"operations,omitempty"`    //
}
type ResponseDevicesLegitOperationsForInterfaceResponseProperties struct {
	Name          string `json:"name,omitempty"`          // Name of the Property
	Applicable    string `json:"applicable,omitempty"`    // Checks if property is applicable to interface
	FailureReason string `json:"failureReason,omitempty"` // Failure reason of the Property
}
type ResponseDevicesLegitOperationsForInterfaceResponseOperations struct {
	Name          string `json:"name,omitempty"`          // Name of the Operation
	Applicable    string `json:"applicable,omitempty"`    // Checks if operation is applicable to interface
	FailureReason string `json:"failureReason,omitempty"` // Failure reason of the Operation
}
type ResponseDevicesClearMacAddressTable struct {
	Response *ResponseDevicesClearMacAddressTableResponse `json:"response,omitempty"` //
	Version  string                                       `json:"version,omitempty"`  // Version
}
type ResponseDevicesClearMacAddressTableResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseDevicesGetDeviceList struct {
	Response *[]ResponseDevicesGetDeviceListResponse `json:"response,omitempty"` //
	Version  string                                  `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetDeviceListResponse struct {
	ReachabilityFailureReason     string   `json:"reachabilityFailureReason,omitempty"`     // Failure reason for unreachable devices
	ReachabilityStatus            string   `json:"reachabilityStatus,omitempty"`            // Device reachability status as Reachable / Unreachable
	Series                        string   `json:"series,omitempty"`                        // Device series
	SNMPContact                   string   `json:"snmpContact,omitempty"`                   // SNMP contact on device
	SNMPLocation                  string   `json:"snmpLocation,omitempty"`                  // SNMP location on device
	TagCount                      string   `json:"tagCount,omitempty"`                      // Number of tags associated with the device
	TunnelUDPPort                 string   `json:"tunnelUdpPort,omitempty"`                 // Mobility protocol port is stored as tunneludpport for WLC
	UptimeSeconds                 *float64 `json:"uptimeSeconds,omitempty"`                 // Uptime in Seconds
	WaasDeviceMode                string   `json:"waasDeviceMode,omitempty"`                // WAAS device mode
	SerialNumber                  string   `json:"serialNumber,omitempty"`                  // Serial number of device
	LastUpdateTime                *float64 `json:"lastUpdateTime,omitempty"`                // Time in epoch when the network device info last got updated
	MacAddress                    string   `json:"macAddress,omitempty"`                    // MAC address of device
	UpTime                        string   `json:"upTime,omitempty"`                        // Time that shows for how long the device has been up
	DeviceSupportLevel            string   `json:"deviceSupportLevel,omitempty"`            // Support level of the device
	Hostname                      string   `json:"hostname,omitempty"`                      // Device name
	Type                          string   `json:"type,omitempty"`                          // Type of device as switch, router, wireless lan controller, accesspoints
	MemorySize                    string   `json:"memorySize,omitempty"`                    // Processor memory size
	Family                        string   `json:"family,omitempty"`                        // Family of device as switch, router, wireless lan controller, accesspoints
	ErrorCode                     string   `json:"errorCode,omitempty"`                     // Inventory status error code
	SoftwareType                  string   `json:"softwareType,omitempty"`                  // Software type on the device
	SoftwareVersion               string   `json:"softwareVersion,omitempty"`               // Software version on the device
	Description                   string   `json:"description,omitempty"`                   // System description
	RoleSource                    string   `json:"roleSource,omitempty"`                    // Role source as manual / auto
	Location                      string   `json:"location,omitempty"`                      // [Deprecated] Location ID that is associated with the device
	Role                          string   `json:"role,omitempty"`                          // Role of device as access, distribution, border router, core
	CollectionInterval            string   `json:"collectionInterval,omitempty"`            // Re sync Interval of the device
	InventoryStatusDetail         string   `json:"inventoryStatusDetail,omitempty"`         // Status detail of inventory sync
	ApEthernetMacAddress          string   `json:"apEthernetMacAddress,omitempty"`          // AccessPoint Ethernet MacAddress of AP device
	ApManagerInterfaceIP          string   `json:"apManagerInterfaceIp,omitempty"`          // IP address of WLC on AP manager interface
	AssociatedWlcIP               string   `json:"associatedWlcIp,omitempty"`               // Associated Wlc Ip address of the AP device
	BootDateTime                  string   `json:"bootDateTime,omitempty"`                  // Device boot time
	CollectionStatus              string   `json:"collectionStatus,omitempty"`              // Collection status as Synchronizing, Could not synchronize, Not manageable, Managed, Partial Collection Failure, Incomplete, Unreachable, Wrong credential, Reachable, In Progress
	ErrorDescription              string   `json:"errorDescription,omitempty"`              // Inventory status description
	InterfaceCount                string   `json:"interfaceCount,omitempty"`                // Number of interfaces on the device
	LastUpdated                   string   `json:"lastUpdated,omitempty"`                   // Time when the network device info last got updated
	LineCardCount                 string   `json:"lineCardCount,omitempty"`                 // Number of linecards on the device
	LineCardID                    string   `json:"lineCardId,omitempty"`                    // IDs of linecards of the device
	LocationName                  string   `json:"locationName,omitempty"`                  // [Deprecated] Name of the associated location
	ManagedAtleastOnce            *bool    `json:"managedAtleastOnce,omitempty"`            // Indicates if device went into Managed state atleast once
	ManagementIPAddress           string   `json:"managementIpAddress,omitempty"`           // IP address of the device
	PlatformID                    string   `json:"platformId,omitempty"`                    // Platform ID of device
	ManagementState               string   `json:"managementState,omitempty"`               // Represents the current management state of the network element: managed, unmanaged, under maintenance, and so on.
	PendingSyncRequestsCount      string   `json:"pendingSyncRequestsCount,omitempty"`      // Count of pending sync requests , if any
	ReasonsForDeviceResync        string   `json:"reasonsForDeviceResync,omitempty"`        // Reason for last/ongoing sync
	ReasonsForPendingSyncRequests string   `json:"reasonsForPendingSyncRequests,omitempty"` // Reasons for pending sync requests , if any
	SyncRequestedByApp            string   `json:"syncRequestedByApp,omitempty"`            // Applications which requested for the resync of network device
	LastManagedResyncReasons      string   `json:"lastManagedResyncReasons,omitempty"`      // Reasons for last successful sync
	DNSResolvedManagementAddress  string   `json:"dnsResolvedManagementAddress,omitempty"`  // Specifies the resolved ip address of dns name
	LastDeviceResyncStartTime     string   `json:"lastDeviceResyncStartTime,omitempty"`     // Start time for last/ongoing sync
	InstanceTenantID              string   `json:"instanceTenantId,omitempty"`              // Instance Tenant Id of the device
	InstanceUUID                  string   `json:"instanceUuid,omitempty"`                  // Instance Uuid of the device
	ID                            string   `json:"id,omitempty"`                            // Instance Uuid of the device
}
type ResponseDevicesAddDevice2 struct {
	Response *ResponseDevicesAddDevice2Response `json:"response,omitempty"` //
	Version  string                             `json:"version,omitempty"`  //
}
type ResponseDevicesAddDevice2Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDevicesUpdateDeviceDetails struct {
	Response *ResponseDevicesUpdateDeviceDetailsResponse `json:"response,omitempty"` //
	Version  string                                      `json:"version,omitempty"`  //
}
type ResponseDevicesUpdateDeviceDetailsResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDevicesGetDeviceValuesThatMatchFullyOrPartiallyAnAttribute interface{}
type ResponseDevicesUpdateDeviceRole struct {
	Response *ResponseDevicesUpdateDeviceRoleResponse `json:"response,omitempty"` //
	Version  string                                   `json:"version,omitempty"`  //
}
type ResponseDevicesUpdateDeviceRoleResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDevicesGetPollingIntervalForAllDevices struct {
	Response *int   `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}
type ResponseDevicesGetDeviceConfigForAllDevices struct {
	Response *[]ResponseDevicesGetDeviceConfigForAllDevicesResponse `json:"response,omitempty"` //
	Version  string                                                 `json:"version,omitempty"`  //
}
type ResponseDevicesGetDeviceConfigForAllDevicesResponse struct {
	AttributeInfo   *ResponseDevicesGetDeviceConfigForAllDevicesResponseAttributeInfo `json:"attributeInfo,omitempty"`   //
	CdpNeighbors    string                                                            `json:"cdpNeighbors,omitempty"`    //
	HealthMonitor   string                                                            `json:"healthMonitor,omitempty"`   //
	ID              string                                                            `json:"id,omitempty"`              //
	IntfDescription string                                                            `json:"intfDescription,omitempty"` //
	Inventory       string                                                            `json:"inventory,omitempty"`       //
	IPIntfBrief     string                                                            `json:"ipIntfBrief,omitempty"`     //
	MacAddressTable string                                                            `json:"macAddressTable,omitempty"` //
	RunningConfig   string                                                            `json:"runningConfig,omitempty"`   //
	SNMP            string                                                            `json:"snmp,omitempty"`            //
	Version         string                                                            `json:"version,omitempty"`         //
}
type ResponseDevicesGetDeviceConfigForAllDevicesResponseAttributeInfo interface{}
type ResponseDevicesGetDeviceConfigCount struct {
	Response *int   `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}
type ResponseDevicesGetDeviceCount2 struct {
	Response *int   `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}
type ResponseDevicesExportDeviceList struct {
	Response *ResponseDevicesExportDeviceListResponse `json:"response,omitempty"` //
	Version  string                                   `json:"version,omitempty"`  //
}
type ResponseDevicesExportDeviceListResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDevicesGetFunctionalCapabilityForDevices struct {
	Response *[]ResponseDevicesGetFunctionalCapabilityForDevicesResponse `json:"response,omitempty"` //
	Version  string                                                      `json:"version,omitempty"`  //
}
type ResponseDevicesGetFunctionalCapabilityForDevicesResponse struct {
	AttributeInfo        *ResponseDevicesGetFunctionalCapabilityForDevicesResponseAttributeInfo          `json:"attributeInfo,omitempty"`        // Deprecated
	DeviceID             string                                                                          `json:"deviceId,omitempty"`             // Device Id of the device
	FunctionalCapability *[]ResponseDevicesGetFunctionalCapabilityForDevicesResponseFunctionalCapability `json:"functionalCapability,omitempty"` //
	ID                   string                                                                          `json:"id,omitempty"`                   // Deprecated
}
type ResponseDevicesGetFunctionalCapabilityForDevicesResponseAttributeInfo interface{}
type ResponseDevicesGetFunctionalCapabilityForDevicesResponseFunctionalCapability struct {
	AttributeInfo   *ResponseDevicesGetFunctionalCapabilityForDevicesResponseFunctionalCapabilityAttributeInfo     `json:"attributeInfo,omitempty"`   // Deprecated
	FunctionDetails *[]ResponseDevicesGetFunctionalCapabilityForDevicesResponseFunctionalCapabilityFunctionDetails `json:"functionDetails,omitempty"` //
	FunctionName    string                                                                                         `json:"functionName,omitempty"`    // Name of the function
	FunctionOpState string                                                                                         `json:"functionOpState,omitempty"` // Operational state of the function
	ID              string                                                                                         `json:"id,omitempty"`              // Id of the function
}
type ResponseDevicesGetFunctionalCapabilityForDevicesResponseFunctionalCapabilityAttributeInfo interface{}
type ResponseDevicesGetFunctionalCapabilityForDevicesResponseFunctionalCapabilityFunctionDetails struct {
	AttributeInfo *ResponseDevicesGetFunctionalCapabilityForDevicesResponseFunctionalCapabilityFunctionDetailsAttributeInfo `json:"attributeInfo,omitempty"` // Deprecated
	ID            string                                                                                                    `json:"id,omitempty"`            // Deprecated
	PropertyName  string                                                                                                    `json:"propertyName,omitempty"`  // Property Name of the function
	StringValue   string                                                                                                    `json:"stringValue,omitempty"`   // Value for the property
}
type ResponseDevicesGetFunctionalCapabilityForDevicesResponseFunctionalCapabilityFunctionDetailsAttributeInfo interface{}
type ResponseDevicesGetFunctionalCapabilityByID struct {
	Response *ResponseDevicesGetFunctionalCapabilityByIDResponse `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  //
}
type ResponseDevicesGetFunctionalCapabilityByIDResponse struct {
	AttributeInfo   *ResponseDevicesGetFunctionalCapabilityByIDResponseAttributeInfo     `json:"attributeInfo,omitempty"`   // Deprecated
	FunctionDetails *[]ResponseDevicesGetFunctionalCapabilityByIDResponseFunctionDetails `json:"functionDetails,omitempty"` //
	FunctionName    string                                                               `json:"functionName,omitempty"`    // Name of the function
	FunctionOpState string                                                               `json:"functionOpState,omitempty"` // Operational state of the function
	ID              string                                                               `json:"id,omitempty"`              // Id of the function
}
type ResponseDevicesGetFunctionalCapabilityByIDResponseAttributeInfo interface{}
type ResponseDevicesGetFunctionalCapabilityByIDResponseFunctionDetails struct {
	AttributeInfo *ResponseDevicesGetFunctionalCapabilityByIDResponseFunctionDetailsAttributeInfo `json:"attributeInfo,omitempty"` // Deprecated
	ID            string                                                                          `json:"id,omitempty"`            // Deprecated
	PropertyName  string                                                                          `json:"propertyName,omitempty"`  // Property Name of the function
	StringValue   string                                                                          `json:"stringValue,omitempty"`   // Value for the property
}
type ResponseDevicesGetFunctionalCapabilityByIDResponseFunctionDetailsAttributeInfo interface{}
type ResponseDevicesInventoryInsightDeviceLinkMismatchAPI struct {
	Response *[]ResponseDevicesInventoryInsightDeviceLinkMismatchApIResponse `json:"response,omitempty"` //
	Version  string                                                          `json:"version,omitempty"`  // Api version
}
type ResponseDevicesInventoryInsightDeviceLinkMismatchApIResponse struct {
	EndPortAllowedVLANIDs   string   `json:"endPortAllowedVlanIds,omitempty"`   // End port allowed vlan ids
	EndPortNativeVLANID     string   `json:"endPortNativeVlanId,omitempty"`     // End port native vlan id
	StartPortAllowedVLANIDs string   `json:"startPortAllowedVlanIds,omitempty"` // Start port allowed vlan ids
	StartPortNativeVLANID   string   `json:"startPortNativeVlanId,omitempty"`   // Start port native vlan id
	LinkStatus              string   `json:"linkStatus,omitempty"`              // Link status
	EndDeviceHostName       string   `json:"endDeviceHostName,omitempty"`       // End device hostname
	EndDeviceID             string   `json:"endDeviceId,omitempty"`             // End device id
	EndDeviceIPAddress      string   `json:"endDeviceIpAddress,omitempty"`      // End device ip address
	EndPortAddress          string   `json:"endPortAddress,omitempty"`          // End port address
	EndPortDuplex           string   `json:"endPortDuplex,omitempty"`           // End port duplex
	EndPortID               string   `json:"endPortId,omitempty"`               // End port id
	EndPortMask             string   `json:"endPortMask,omitempty"`             // End port mask
	EndPortName             string   `json:"endPortName,omitempty"`             // End port name
	EndPortPepID            string   `json:"endPortPepId,omitempty"`            // End port pep id
	EndPortSpeed            string   `json:"endPortSpeed,omitempty"`            // End port speed
	StartDeviceHostName     string   `json:"startDeviceHostName,omitempty"`     // Start device hostname
	StartDeviceID           string   `json:"startDeviceId,omitempty"`           // Start device id
	StartDeviceIPAddress    string   `json:"startDeviceIpAddress,omitempty"`    // Start device ip address
	StartPortAddress        string   `json:"startPortAddress,omitempty"`        // Start port address
	StartPortDuplex         string   `json:"startPortDuplex,omitempty"`         // Start port duplex
	StartPortID             string   `json:"startPortId,omitempty"`             // Start port id
	StartPortMask           string   `json:"startPortMask,omitempty"`           // Start port mask
	StartPortName           string   `json:"startPortName,omitempty"`           // Start port name
	StartPortPepID          string   `json:"startPortPepId,omitempty"`          // Start port pep id
	StartPortSpeed          string   `json:"startPortSpeed,omitempty"`          // Start port speed
	LastUpdated             string   `json:"lastUpdated,omitempty"`             // Last updated
	NumUpdates              *float64 `json:"numUpdates,omitempty"`              // Number updates
	AvgUpdateFrequency      *float64 `json:"avgUpdateFrequency,omitempty"`      // Average update frequency
	Type                    string   `json:"type,omitempty"`                    // Type
	InstanceUUID            string   `json:"instanceUuid,omitempty"`            // Unique instance id
	InstanceTenantID        string   `json:"instanceTenantId,omitempty"`        // Instance tenant id
}
type ResponseDevicesGetNetworkDeviceByIP struct {
	Response *ResponseDevicesGetNetworkDeviceByIPResponse `json:"response,omitempty"` //
	Version  string                                       `json:"version,omitempty"`  //
}
type ResponseDevicesGetNetworkDeviceByIPResponse struct {
	ApManagerInterfaceIP          string   `json:"apManagerInterfaceIp,omitempty"`          // IP address of WLC on AP manager interface
	AssociatedWlcIP               string   `json:"associatedWlcIp,omitempty"`               // Associated Wlc Ip address of the AP device
	BootDateTime                  string   `json:"bootDateTime,omitempty"`                  // Device boot time
	CollectionInterval            string   `json:"collectionInterval,omitempty"`            // Re sync Interval of the device
	CollectionStatus              string   `json:"collectionStatus,omitempty"`              // Collection status as Synchronizing, Could not synchronize, Not manageable, Managed, Partial Collection Failure, Incomplete, Unreachable, Wrong credential, Reachable, In Progress
	ErrorCode                     string   `json:"errorCode,omitempty"`                     // Inventory status error code
	ErrorDescription              string   `json:"errorDescription,omitempty"`              // Inventory status description
	Family                        string   `json:"family,omitempty"`                        // Family of device as switch, router, wireless lan controller, accesspoints
	Hostname                      string   `json:"hostname,omitempty"`                      // Device name
	ID                            string   `json:"id,omitempty"`                            // Instance Uuid of the device
	InstanceTenantID              string   `json:"instanceTenantId,omitempty"`              // Instance Tenant Id of the device
	InstanceUUID                  string   `json:"instanceUuid,omitempty"`                  // Instance Uuid of the device
	InterfaceCount                string   `json:"interfaceCount,omitempty"`                // Number of interfaces on the device
	InventoryStatusDetail         string   `json:"inventoryStatusDetail,omitempty"`         // Status detail of inventory sync
	LastUpdateTime                *float64 `json:"lastUpdateTime,omitempty"`                // Time in epoch when the network device info last got updated
	LastUpdated                   string   `json:"lastUpdated,omitempty"`                   // Time when the network device info last got updated
	LineCardCount                 string   `json:"lineCardCount,omitempty"`                 // Number of linecards on the device
	LineCardID                    string   `json:"lineCardId,omitempty"`                    // IDs of linecards of the device
	Location                      string   `json:"location,omitempty"`                      // [Deprecated] Location ID that is associated with the device
	LocationName                  string   `json:"locationName,omitempty"`                  // [Deprecated] Name of the associated location
	MacAddress                    string   `json:"macAddress,omitempty"`                    // MAC address of device
	ManagementIPAddress           string   `json:"managementIpAddress,omitempty"`           // IP address of the device
	MemorySize                    string   `json:"memorySize,omitempty"`                    // Processor memory size
	PlatformID                    string   `json:"platformId,omitempty"`                    // Platform ID of device
	ReachabilityFailureReason     string   `json:"reachabilityFailureReason,omitempty"`     // Failure reason for unreachable devices
	ReachabilityStatus            string   `json:"reachabilityStatus,omitempty"`            // Device reachability status as Reachable / Unreachable
	Role                          string   `json:"role,omitempty"`                          // Role of device as access, distribution, border router, core
	RoleSource                    string   `json:"roleSource,omitempty"`                    // Role source as manual / auto
	SerialNumber                  string   `json:"serialNumber,omitempty"`                  // Serial number of device
	Series                        string   `json:"series,omitempty"`                        // Device series
	SNMPContact                   string   `json:"snmpContact,omitempty"`                   // SNMP contact on device
	SNMPLocation                  string   `json:"snmpLocation,omitempty"`                  // SNMP location on device
	SoftwareType                  string   `json:"softwareType,omitempty"`                  // Software type on the device
	SoftwareVersion               string   `json:"softwareVersion,omitempty"`               // Software version on the device
	TagCount                      string   `json:"tagCount,omitempty"`                      // Number of tags associated with the device
	TunnelUDPPort                 string   `json:"tunnelUdpPort,omitempty"`                 // Mobility protocol port is stored as tunneludpport for WLC
	Type                          string   `json:"type,omitempty"`                          // Type of device as switch, router, wireless lan controller, accesspoints
	UpTime                        string   `json:"upTime,omitempty"`                        // Time that shows for how long the device has been up
	WaasDeviceMode                string   `json:"waasDeviceMode,omitempty"`                // WAAS device mode
	DNSResolvedManagementAddress  string   `json:"dnsResolvedManagementAddress,omitempty"`  // Specifies the resolved ip address of dns name
	ApEthernetMacAddress          string   `json:"apEthernetMacAddress,omitempty"`          // AccessPoint Ethernet MacAddress of AP device
	Vendor                        string   `json:"vendor,omitempty"`                        // Vendor details
	ReasonsForPendingSyncRequests string   `json:"reasonsForPendingSyncRequests,omitempty"` // Reasons for pending sync requests , if any
	PendingSyncRequestsCount      string   `json:"pendingSyncRequestsCount,omitempty"`      // Count of pending sync requests , if any
	ReasonsForDeviceResync        string   `json:"reasonsForDeviceResync,omitempty"`        // Reason for last/ongoing sync
	LastDeviceResyncStartTime     string   `json:"lastDeviceResyncStartTime,omitempty"`     // Start time for last/ongoing sync
	UptimeSeconds                 *float64 `json:"uptimeSeconds,omitempty"`                 // Uptime in Seconds
	ManagedAtleastOnce            *bool    `json:"managedAtleastOnce,omitempty"`            // Indicates if device went into Managed state atleast once
	DeviceSupportLevel            string   `json:"deviceSupportLevel,omitempty"`            // Support level of the device
	ManagementState               string   `json:"managementState,omitempty"`               // Represents the current management state of the network element: managed, unmanaged, under maintenance, and so on.
	Description                   string   `json:"description,omitempty"`                   // System description
}
type ResponseDevicesGetModules struct {
	Response *[]ResponseDevicesGetModulesResponse `json:"response,omitempty"` //
	Version  string                               `json:"version,omitempty"`  //
}
type ResponseDevicesGetModulesResponse struct {
	AssemblyNumber           string                                          `json:"assemblyNumber,omitempty"`           // Assembly Number of the module
	AssemblyRevision         string                                          `json:"assemblyRevision,omitempty"`         // Assembly Revision of the module
	AttributeInfo            *ResponseDevicesGetModulesResponseAttributeInfo `json:"attributeInfo,omitempty"`            // Deprecated
	ContainmentEntity        string                                          `json:"containmentEntity,omitempty"`        // Containment Entity of the module
	Description              string                                          `json:"description,omitempty"`              // Description of the module
	EntityPhysicalIndex      string                                          `json:"entityPhysicalIndex,omitempty"`      // Entity Physical Index of the module
	ID                       string                                          `json:"id,omitempty"`                       // ID of the module
	IsFieldReplaceable       string                                          `json:"isFieldReplaceable,omitempty"`       // To mention if field is replaceable
	IsReportingAlarmsAllowed string                                          `json:"isReportingAlarmsAllowed,omitempty"` // To mention if reporting alarms are allowed
	Manufacturer             string                                          `json:"manufacturer,omitempty"`             // Manufacturer of the module
	ModuleIndex              *int                                            `json:"moduleIndex,omitempty"`              // Index of the module
	Name                     string                                          `json:"name,omitempty"`                     // Name of the module
	OperationalStateCode     string                                          `json:"operationalStateCode,omitempty"`     // Operational state of the module
	PartNumber               string                                          `json:"partNumber,omitempty"`               // Part number of the module
	SerialNumber             string                                          `json:"serialNumber,omitempty"`             // Serial number of the module
	VendorEquipmentType      string                                          `json:"vendorEquipmentType,omitempty"`      // Vendor euipment type of the module
}
type ResponseDevicesGetModulesResponseAttributeInfo interface{}
type ResponseDevicesGetModuleCount struct {
	Response *int   `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}
type ResponseDevicesGetModuleInfoByID struct {
	Response *ResponseDevicesGetModuleInfoByIDResponse `json:"response,omitempty"` //
	Version  string                                    `json:"version,omitempty"`  //
}
type ResponseDevicesGetModuleInfoByIDResponse struct {
	AssemblyNumber           string                                                 `json:"assemblyNumber,omitempty"`           // Assembly number of the module
	AssemblyRevision         string                                                 `json:"assemblyRevision,omitempty"`         // Assembly revision of the module
	AttributeInfo            *ResponseDevicesGetModuleInfoByIDResponseAttributeInfo `json:"attributeInfo,omitempty"`            // Deprecated
	ContainmentEntity        string                                                 `json:"containmentEntity,omitempty"`        // Containment entity of the module
	Description              string                                                 `json:"description,omitempty"`              // Description of the module
	EntityPhysicalIndex      string                                                 `json:"entityPhysicalIndex,omitempty"`      // Entity physical index of the module
	ID                       string                                                 `json:"id,omitempty"`                       // Id of the module
	IsFieldReplaceable       string                                                 `json:"isFieldReplaceable,omitempty"`       // To mention if field is replaceable
	IsReportingAlarmsAllowed string                                                 `json:"isReportingAlarmsAllowed,omitempty"` // To mention if reporting alarms are allowed
	Manufacturer             string                                                 `json:"manufacturer,omitempty"`             // Manufacturer of the module
	ModuleIndex              *int                                                   `json:"moduleIndex,omitempty"`              // Index of the module
	Name                     string                                                 `json:"name,omitempty"`                     // Name of the module
	OperationalStateCode     string                                                 `json:"operationalStateCode,omitempty"`     // Operational state of the module
	PartNumber               string                                                 `json:"partNumber,omitempty"`               // Part number of the module
	SerialNumber             string                                                 `json:"serialNumber,omitempty"`             // Serial number of the modules
	VendorEquipmentType      string                                                 `json:"vendorEquipmentType,omitempty"`      // Vendor equipment type of the module
}
type ResponseDevicesGetModuleInfoByIDResponseAttributeInfo interface{}
type ResponseDevicesGetDeviceBySerialNumber struct {
	Response *ResponseDevicesGetDeviceBySerialNumberResponse `json:"response,omitempty"` //
	Version  string                                          `json:"version,omitempty"`  //
}
type ResponseDevicesGetDeviceBySerialNumberResponse struct {
	ApManagerInterfaceIP          string   `json:"apManagerInterfaceIp,omitempty"`          // IP address of WLC on AP manager interface
	AssociatedWlcIP               string   `json:"associatedWlcIp,omitempty"`               // Associated Wlc Ip address of the AP device
	BootDateTime                  string   `json:"bootDateTime,omitempty"`                  // Device boot time
	CollectionInterval            string   `json:"collectionInterval,omitempty"`            // Re sync Interval of the device
	CollectionStatus              string   `json:"collectionStatus,omitempty"`              // Collection status as Synchronizing, Could not synchronize, Not manageable, Managed, Partial Collection Failure, Incomplete, Unreachable, Wrong credential, Reachable, In Progress
	ErrorCode                     string   `json:"errorCode,omitempty"`                     // Inventory status error code
	ErrorDescription              string   `json:"errorDescription,omitempty"`              // Inventory status description
	Family                        string   `json:"family,omitempty"`                        // Family of device as switch, router, wireless lan controller, accesspoints
	Hostname                      string   `json:"hostname,omitempty"`                      // Device name
	ID                            string   `json:"id,omitempty"`                            // Instance Uuid of the device
	InstanceTenantID              string   `json:"instanceTenantId,omitempty"`              // Instance Tenant Id of the device
	InstanceUUID                  string   `json:"instanceUuid,omitempty"`                  // Instance Uuid of the device
	InterfaceCount                string   `json:"interfaceCount,omitempty"`                // Number of interfaces on the device
	InventoryStatusDetail         string   `json:"inventoryStatusDetail,omitempty"`         // Status detail of inventory sync
	LastUpdateTime                *float64 `json:"lastUpdateTime,omitempty"`                // Time in epoch when the network device info last got updated
	LastUpdated                   string   `json:"lastUpdated,omitempty"`                   // Time when the network device info last got updated
	LineCardCount                 string   `json:"lineCardCount,omitempty"`                 // Number of linecards on the device
	LineCardID                    string   `json:"lineCardId,omitempty"`                    // IDs of linecards of the device
	Location                      string   `json:"location,omitempty"`                      // [Deprecated] Location ID that is associated with the device
	LocationName                  string   `json:"locationName,omitempty"`                  // [Deprecated] Name of the associated location
	MacAddress                    string   `json:"macAddress,omitempty"`                    // MAC address of device
	ManagementIPAddress           string   `json:"managementIpAddress,omitempty"`           // IP address of the device
	MemorySize                    string   `json:"memorySize,omitempty"`                    // Processor memory size
	PlatformID                    string   `json:"platformId,omitempty"`                    // Platform ID of device
	ReachabilityFailureReason     string   `json:"reachabilityFailureReason,omitempty"`     // Failure reason for unreachable devices
	ReachabilityStatus            string   `json:"reachabilityStatus,omitempty"`            // Device reachability status as Reachable / Unreachable
	Role                          string   `json:"role,omitempty"`                          // Role of device as access, distribution, border router, core
	RoleSource                    string   `json:"roleSource,omitempty"`                    // Role source as manual / auto
	SerialNumber                  string   `json:"serialNumber,omitempty"`                  // Serial number of device
	Series                        string   `json:"series,omitempty"`                        // Device series
	SNMPContact                   string   `json:"snmpContact,omitempty"`                   // SNMP contact on device
	SNMPLocation                  string   `json:"snmpLocation,omitempty"`                  // SNMP location on device
	SoftwareType                  string   `json:"softwareType,omitempty"`                  // Software type on the device
	SoftwareVersion               string   `json:"softwareVersion,omitempty"`               // Software version on the device
	TagCount                      string   `json:"tagCount,omitempty"`                      // Number of tags associated with the device
	TunnelUDPPort                 string   `json:"tunnelUdpPort,omitempty"`                 // Mobility protocol port is stored as tunneludpport for WLC
	Type                          string   `json:"type,omitempty"`                          // Type of device as switch, router, wireless lan controller, accesspoints
	UpTime                        string   `json:"upTime,omitempty"`                        // Time that shows for how long the device has been up
	WaasDeviceMode                string   `json:"waasDeviceMode,omitempty"`                // WAAS device mode
	DNSResolvedManagementAddress  string   `json:"dnsResolvedManagementAddress,omitempty"`  // Specifies the resolved ip address of dns name
	ApEthernetMacAddress          string   `json:"apEthernetMacAddress,omitempty"`          // AccessPoint Ethernet MacAddress of AP device
	Vendor                        string   `json:"vendor,omitempty"`                        // Vendor details
	ReasonsForPendingSyncRequests string   `json:"reasonsForPendingSyncRequests,omitempty"` // Reasons for pending sync requests , if any
	PendingSyncRequestsCount      string   `json:"pendingSyncRequestsCount,omitempty"`      // Count of pending sync requests , if any
	ReasonsForDeviceResync        string   `json:"reasonsForDeviceResync,omitempty"`        // Reason for last/ongoing sync
	LastDeviceResyncStartTime     string   `json:"lastDeviceResyncStartTime,omitempty"`     // Start time for last/ongoing sync
	UptimeSeconds                 *float64 `json:"uptimeSeconds,omitempty"`                 // Uptime in Seconds
	ManagedAtleastOnce            *bool    `json:"managedAtleastOnce,omitempty"`            // Indicates if device went into Managed state atleast once
	DeviceSupportLevel            string   `json:"deviceSupportLevel,omitempty"`            // Support level of the device
	ManagementState               string   `json:"managementState,omitempty"`               // Represents the current management state of the network element: managed, unmanaged, under maintenance, and so on.
	Description                   string   `json:"description,omitempty"`                   // System description
}
type ResponseDevicesSyncDevices struct {
	Response *ResponseDevicesSyncDevicesResponse `json:"response,omitempty"` //
	Version  string                              `json:"version,omitempty"`  //
}
type ResponseDevicesSyncDevicesResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDevicesGetDevicesRegisteredForWsaNotification struct {
	Response *ResponseDevicesGetDevicesRegisteredForWsaNotificationResponse `json:"response,omitempty"` //
	Version  string                                                         `json:"version,omitempty"`  //
}
type ResponseDevicesGetDevicesRegisteredForWsaNotificationResponse struct {
	MacAddress   string `json:"macAddress,omitempty"`   // MAC address of device
	ModelNumber  string `json:"modelNumber,omitempty"`  // Model number of the device
	Name         string `json:"name,omitempty"`         // Name of the device
	SerialNumber string `json:"serialNumber,omitempty"` // Serial Number of the device
	TenantID     string `json:"tenantId,omitempty"`     // Tenant Id of the device
}
type ResponseDevicesGetAllUserDefinedFields struct {
	Response *[]ResponseDevicesGetAllUserDefinedFieldsResponse `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetAllUserDefinedFieldsResponse struct {
	ID          string `json:"id,omitempty"`          // DeviceId of the Device
	Name        string `json:"name,omitempty"`        // UDF name
	Description string `json:"description,omitempty"` // Description for UDF
}
type ResponseDevicesCreateUserDefinedField struct {
	Response *ResponseDevicesCreateUserDefinedFieldResponse `json:"response,omitempty"` //
	Version  string                                         `json:"version,omitempty"`  // Version
}
type ResponseDevicesCreateUserDefinedFieldResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseDevicesUpdateUserDefinedField struct {
	Response *ResponseDevicesUpdateUserDefinedFieldResponse `json:"response,omitempty"` //
	Version  string                                         `json:"version,omitempty"`  // Version
}
type ResponseDevicesUpdateUserDefinedFieldResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseDevicesDeleteUserDefinedField struct {
	Response *ResponseDevicesDeleteUserDefinedFieldResponse `json:"response,omitempty"` //
	Version  string                                         `json:"version,omitempty"`  // Version
}
type ResponseDevicesDeleteUserDefinedFieldResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseDevicesGetChassisDetailsForDevice struct {
	Response *[]ResponseDevicesGetChassisDetailsForDeviceResponse `json:"response,omitempty"` //
	Version  string                                               `json:"version,omitempty"`  //
}
type ResponseDevicesGetChassisDetailsForDeviceResponse struct {
	AssemblyNumber           string `json:"assemblyNumber,omitempty"`           // Assembly Number of the chassis
	AssemblyRevision         string `json:"assemblyRevision,omitempty"`         // Assembly Revision of the chassis
	ContainmentEntity        string `json:"containmentEntity,omitempty"`        // Containment Entity of the chassis
	Description              string `json:"description,omitempty"`              // Description of the chassis
	EntityPhysicalIndex      string `json:"entityPhysicalIndex,omitempty"`      // Entity Physical Index of the chassis
	HardwareVersion          string `json:"hardwareVersion,omitempty"`          // Hardware Version of the chassis
	InstanceUUID             string `json:"instanceUuid,omitempty"`             // ID of the chassis
	IsFieldReplaceable       string `json:"isFieldReplaceable,omitempty"`       // To mention if field is replaceable
	IsReportingAlarmsAllowed string `json:"isReportingAlarmsAllowed,omitempty"` // To mention if reporting alarms are allowed
	Manufacturer             string `json:"manufacturer,omitempty"`             // Manufacturer of the chassis
	Name                     string `json:"name,omitempty"`                     // Name of the chassis
	PartNumber               string `json:"partNumber,omitempty"`               // Part Number of the chassis
	SerialNumber             string `json:"serialNumber,omitempty"`             // Serial Number of the chassis
	VendorEquipmentType      string `json:"vendorEquipmentType,omitempty"`      // Vendor Equipment Type of the chassis
}
type ResponseDevicesGetStackDetailsForDevice struct {
	Response *ResponseDevicesGetStackDetailsForDeviceResponse `json:"response,omitempty"` //
	Version  string                                           `json:"version,omitempty"`  //
}
type ResponseDevicesGetStackDetailsForDeviceResponse struct {
	DeviceID        string                                                            `json:"deviceId,omitempty"`        // Device ID
	StackPortInfo   *[]ResponseDevicesGetStackDetailsForDeviceResponseStackPortInfo   `json:"stackPortInfo,omitempty"`   //
	StackSwitchInfo *[]ResponseDevicesGetStackDetailsForDeviceResponseStackSwitchInfo `json:"stackSwitchInfo,omitempty"` //
	SvlSwitchInfo   *[]ResponseDevicesGetStackDetailsForDeviceResponseSvlSwitchInfo   `json:"svlSwitchInfo,omitempty"`   //
}
type ResponseDevicesGetStackDetailsForDeviceResponseStackPortInfo struct {
	IsSynchOk               string `json:"isSynchOk,omitempty"`               // If link partner sends valid protocol message
	LinkActive              *bool  `json:"linkActive,omitempty"`              // If stack port is in same state as link partner
	LinkOk                  *bool  `json:"linkOk,omitempty"`                  // If link is stable
	Name                    string `json:"name,omitempty"`                    // Name of the stack port
	NeighborPort            string `json:"neighborPort,omitempty"`            // Neighbor's member number and stack port number
	NrLinkOkChanges         *int   `json:"nrLinkOkChanges,omitempty"`         // Relative stability of the link
	StackCableLengthInfo    string `json:"stackCableLengthInfo,omitempty"`    // Cable length
	StackPortOperStatusInfo string `json:"stackPortOperStatusInfo,omitempty"` // Port opearation status
	SwitchPort              string `json:"switchPort,omitempty"`              // Member number and stack port number
}
type ResponseDevicesGetStackDetailsForDeviceResponseStackSwitchInfo struct {
	EntPhysicalIndex  string `json:"entPhysicalIndex,omitempty"`  //
	HwPriority        *int   `json:"hwPriority,omitempty"`        // Hardware priority of the switch
	MacAddress        string `json:"macAddress,omitempty"`        // Mac address of the switch
	NumNextReload     *int   `json:"numNextReload,omitempty"`     // Stack member number to be used in next reload
	PlatformID        string `json:"platformId,omitempty"`        // Platform Id
	Role              string `json:"role,omitempty"`              // Function of the switch
	SerialNumber      string `json:"serialNumber,omitempty"`      // Serial number
	SoftwareImage     string `json:"softwareImage,omitempty"`     // Software image type running on the switch
	StackMemberNumber *int   `json:"stackMemberNumber,omitempty"` // Switch member number
	State             string `json:"state,omitempty"`             // Current state of the switch
	SwitchPriority    *int   `json:"switchPriority,omitempty"`    // Priority of the switch
}
type ResponseDevicesGetStackDetailsForDeviceResponseSvlSwitchInfo struct {
	DadProtocol              string                                                                       `json:"dadProtocol,omitempty"`              // Stackwise virtual dual active detection config
	DadRecoveryReloadEnabled *bool                                                                        `json:"dadRecoveryReloadEnabled,omitempty"` // If dad recovery reload enabled
	DomainNumber             *int                                                                         `json:"domainNumber,omitempty"`             // Stackwise virtual switch domain number
	InDadRecoveryMode        *bool                                                                        `json:"inDadRecoveryMode,omitempty"`        // If in dad recovery mode
	SwVirtualStatus          string                                                                       `json:"swVirtualStatus,omitempty"`          // Stackwise virtual status
	SwitchMembers            *[]ResponseDevicesGetStackDetailsForDeviceResponseSvlSwitchInfoSwitchMembers `json:"switchMembers,omitempty"`            //
}
type ResponseDevicesGetStackDetailsForDeviceResponseSvlSwitchInfoSwitchMembers struct {
	Bandwidth            string                                                                                           `json:"bandwidth,omitempty"`            // Bandwidth
	SvlMemberEndPoints   *[]ResponseDevicesGetStackDetailsForDeviceResponseSvlSwitchInfoSwitchMembersSvlMemberEndPoints   `json:"svlMemberEndPoints,omitempty"`   //
	SvlMemberNumber      *int                                                                                             `json:"svlMemberNumber,omitempty"`      // Switch member number
	SvlMemberPepSettings *[]ResponseDevicesGetStackDetailsForDeviceResponseSvlSwitchInfoSwitchMembersSvlMemberPepSettings `json:"svlMemberPepSettings,omitempty"` //
}
type ResponseDevicesGetStackDetailsForDeviceResponseSvlSwitchInfoSwitchMembersSvlMemberEndPoints struct {
	SvlMemberEndPointPorts *[]ResponseDevicesGetStackDetailsForDeviceResponseSvlSwitchInfoSwitchMembersSvlMemberEndPointsSvlMemberEndPointPorts `json:"svlMemberEndPointPorts,omitempty"` //
	SvlNumber              *int                                                                                                                 `json:"svlNumber,omitempty"`              // Stackwise virtual link number
	SvlStatus              string                                                                                                               `json:"svlStatus,omitempty"`              // Stackwise virtual status
}
type ResponseDevicesGetStackDetailsForDeviceResponseSvlSwitchInfoSwitchMembersSvlMemberEndPointsSvlMemberEndPointPorts struct {
	SvlProtocolStatus string `json:"svlProtocolStatus,omitempty"` // Stackwise virtual protocol status
	SwLocalInterface  string `json:"swLocalInterface,omitempty"`  // Stackwise virtual local interface
	SwRemoteInterface string `json:"swRemoteInterface,omitempty"` // Stackwise virtual remote interface
}
type ResponseDevicesGetStackDetailsForDeviceResponseSvlSwitchInfoSwitchMembersSvlMemberPepSettings struct {
	DadEnabled       *bool  `json:"dadEnabled,omitempty"`       // If dadInterface is configured for dual active detection
	DadInterfaceName string `json:"dadInterfaceName,omitempty"` // Interface for dual active detection
}
type ResponseDevicesRemoveUserDefinedFieldFromDevice struct {
	Response *ResponseDevicesRemoveUserDefinedFieldFromDeviceResponse `json:"response,omitempty"` //
	Version  string                                                   `json:"version,omitempty"`  // Version
}
type ResponseDevicesRemoveUserDefinedFieldFromDeviceResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseDevicesAddUserDefinedFieldToDevice struct {
	Response *ResponseDevicesAddUserDefinedFieldToDeviceResponse `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  // Version
}
type ResponseDevicesAddUserDefinedFieldToDeviceResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseDevicesGetTheDetailsOfPhysicalComponentsOfTheGivenDevice struct {
	Response *[]ResponseDevicesGetTheDetailsOfPhysicalComponentsOfTheGivenDeviceResponse `json:"response,omitempty"` //
	Version  string                                                                      `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetTheDetailsOfPhysicalComponentsOfTheGivenDeviceResponse struct {
	OperationalStateCode string `json:"operationalStateCode,omitempty"` // Operational State Code
	ProductID            string `json:"productId,omitempty"`            // Product Id
	SerialNumber         string `json:"serialNumber,omitempty"`         // Serial Number
	VendorEquipmentType  string `json:"vendorEquipmentType,omitempty"`  // Vendor Equipment Type
	Description          string `json:"description,omitempty"`          // Description
	InstanceUUID         string `json:"instanceUuid,omitempty"`         // Instance Uuid
	Name                 string `json:"name,omitempty"`                 // Name
	Manufacturer         string `json:"manufacturer,omitempty"`         // Manufacturer
}
type ResponseDevicesReturnsPoeInterfaceDetailsForTheDevice struct {
	Version  string                                                           `json:"version,omitempty"`  // Version
	Response *[]ResponseDevicesReturnsPoeInterfaceDetailsForTheDeviceResponse `json:"response,omitempty"` //
}
type ResponseDevicesReturnsPoeInterfaceDetailsForTheDeviceResponse struct {
	AdminStatus    string `json:"adminStatus,omitempty"`    // Administration Status. Possible values: AUTO, STATIC, NEVER
	OperStatus     string `json:"operStatus,omitempty"`     // Operational Status. Possible values: ON, OFF, FAULTY, POWER_DENY
	InterfaceName  string `json:"interfaceName,omitempty"`  // Name of the interface
	MaxPortPower   string `json:"maxPortPower,omitempty"`   // Maximum power (in Watts) that port can hold
	AllocatedPower string `json:"allocatedPower,omitempty"` // Power (in Watts) allocated for a given interface
	PortPowerDrawn string `json:"portPowerDrawn,omitempty"` // Power (in Watts) that the port has drawn so far
}
type ResponseDevicesGetConnectedDeviceDetail struct {
	Response *ResponseDevicesGetConnectedDeviceDetailResponse `json:"response,omitempty"` //
	Version  string                                           `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetConnectedDeviceDetailResponse struct {
	NeighborDevice string   `json:"neighborDevice,omitempty"` // Info about the devices connected to the interface
	NeighborPort   string   `json:"neighborPort,omitempty"`   // Info about the connected interface
	Capabilities   []string `json:"capabilities,omitempty"`   // Info about capabilities of the connected device
}
type ResponseDevicesGetLinecardDetails struct {
	Response *[]ResponseDevicesGetLinecardDetailsResponse `json:"response,omitempty"` //
	Version  string                                       `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetLinecardDetailsResponse struct {
	Serialno string `json:"serialno,omitempty"` // Serial number of the line card
	Partno   string `json:"partno,omitempty"`   // Part number of the line card
	Switchno string `json:"switchno,omitempty"` // Switch number of the line card
	Slotno   string `json:"slotno,omitempty"`   // Slot number of line card
}
type ResponseDevicesPoeDetails struct {
	Response *ResponseDevicesPoeDetailsResponse `json:"response,omitempty"` //
	Version  string                             `json:"version,omitempty"`  // Version
}
type ResponseDevicesPoeDetailsResponse struct {
	PowerAllocated string `json:"powerAllocated,omitempty"` // Total power available on the switch on all interfaces combined in Watts
	PowerConsumed  string `json:"powerConsumed,omitempty"`  // Total power being currently drawn by all interfaces combined in Watts
	PowerRemaining string `json:"powerRemaining,omitempty"` // Total power remaining in Watts (powerConsumed - powerAllocated)
}
type ResponseDevicesGetSupervisorCardDetail struct {
	Response *[]ResponseDevicesGetSupervisorCardDetailResponse `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetSupervisorCardDetailResponse struct {
	Serialno string `json:"serialno,omitempty"` // Serial number of the supervisor card
	Partno   string `json:"partno,omitempty"`   // Part number of the supervisor card
	Switchno string `json:"switchno,omitempty"` // Switch number of the supervisor card
	Slotno   string `json:"slotno,omitempty"`   // Slot number of supervisor card
}
type ResponseDevicesUpdateDeviceManagementAddress struct {
	Response *ResponseDevicesUpdateDeviceManagementAddressResponse `json:"response,omitempty"` //
	Version  string                                                `json:"version,omitempty"`  // Version
}
type ResponseDevicesUpdateDeviceManagementAddressResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseDevicesGetDeviceByID struct {
	Response *ResponseDevicesGetDeviceByIDResponse `json:"response,omitempty"` //
	Version  string                                `json:"version,omitempty"`  //
}
type ResponseDevicesGetDeviceByIDResponse struct {
	ApManagerInterfaceIP          string   `json:"apManagerInterfaceIp,omitempty"`          // IP address of WLC on AP manager interface
	AssociatedWlcIP               string   `json:"associatedWlcIp,omitempty"`               // Associated Wlc Ip address of the AP device
	BootDateTime                  string   `json:"bootDateTime,omitempty"`                  // Device boot time
	CollectionInterval            string   `json:"collectionInterval,omitempty"`            // Re sync Interval of the device
	CollectionStatus              string   `json:"collectionStatus,omitempty"`              // Collection status as Synchronizing, Could not synchronize, Not manageable, Managed, Partial Collection Failure, Incomplete, Unreachable, Wrong credential, Reachable, In Progress
	ErrorCode                     string   `json:"errorCode,omitempty"`                     // Inventory status error code
	ErrorDescription              string   `json:"errorDescription,omitempty"`              // Inventory status description
	Family                        string   `json:"family,omitempty"`                        // Family of device as switch, router, wireless lan controller, accesspoints
	Hostname                      string   `json:"hostname,omitempty"`                      // Device name
	ID                            string   `json:"id,omitempty"`                            // Instance Uuid of the device
	InstanceTenantID              string   `json:"instanceTenantId,omitempty"`              // Instance Tenant Id of the device
	InstanceUUID                  string   `json:"instanceUuid,omitempty"`                  // Instance Uuid of the device
	InterfaceCount                string   `json:"interfaceCount,omitempty"`                // Number of interfaces on the device
	InventoryStatusDetail         string   `json:"inventoryStatusDetail,omitempty"`         // Status detail of inventory sync
	LastUpdateTime                *float64 `json:"lastUpdateTime,omitempty"`                // Time in epoch when the network device info last got updated
	LastUpdated                   string   `json:"lastUpdated,omitempty"`                   // Time when the network device info last got updated
	LineCardCount                 string   `json:"lineCardCount,omitempty"`                 // Number of linecards on the device
	LineCardID                    string   `json:"lineCardId,omitempty"`                    // IDs of linecards of the device
	Location                      string   `json:"location,omitempty"`                      // [Deprecated] Location ID that is associated with the device
	LocationName                  string   `json:"locationName,omitempty"`                  // [Deprecated] Name of the associated location
	MacAddress                    string   `json:"macAddress,omitempty"`                    // MAC address of device
	ManagementIPAddress           string   `json:"managementIpAddress,omitempty"`           // IP address of the device
	MemorySize                    string   `json:"memorySize,omitempty"`                    // Processor memory size
	PlatformID                    string   `json:"platformId,omitempty"`                    // Platform ID of device
	ReachabilityFailureReason     string   `json:"reachabilityFailureReason,omitempty"`     // Failure reason for unreachable devices
	ReachabilityStatus            string   `json:"reachabilityStatus,omitempty"`            // Device reachability status as Reachable / Unreachable
	Role                          string   `json:"role,omitempty"`                          // Role of device as access, distribution, border router, core
	RoleSource                    string   `json:"roleSource,omitempty"`                    // Role source as manual / auto
	SerialNumber                  string   `json:"serialNumber,omitempty"`                  // Serial number of device
	Series                        string   `json:"series,omitempty"`                        // Device series
	SNMPContact                   string   `json:"snmpContact,omitempty"`                   // SNMP contact on device
	SNMPLocation                  string   `json:"snmpLocation,omitempty"`                  // SNMP location on device
	SoftwareType                  string   `json:"softwareType,omitempty"`                  // Software type on the device
	SoftwareVersion               string   `json:"softwareVersion,omitempty"`               // Software version on the device
	TagCount                      string   `json:"tagCount,omitempty"`                      // Number of tags associated with the device
	TunnelUDPPort                 string   `json:"tunnelUdpPort,omitempty"`                 // Mobility protocol port is stored as tunneludpport for WLC
	Type                          string   `json:"type,omitempty"`                          // Type of device as switch, router, wireless lan controller, accesspoints
	UpTime                        string   `json:"upTime,omitempty"`                        // Time that shows for how long the device has been up
	WaasDeviceMode                string   `json:"waasDeviceMode,omitempty"`                // WAAS device mode
	DNSResolvedManagementAddress  string   `json:"dnsResolvedManagementAddress,omitempty"`  // Specifies the resolved ip address of dns name
	ApEthernetMacAddress          string   `json:"apEthernetMacAddress,omitempty"`          // AccessPoint Ethernet MacAddress of AP device
	Vendor                        string   `json:"vendor,omitempty"`                        // Vendor details
	ReasonsForPendingSyncRequests string   `json:"reasonsForPendingSyncRequests,omitempty"` // Reasons for pending sync requests , if any
	PendingSyncRequestsCount      string   `json:"pendingSyncRequestsCount,omitempty"`      // Count of pending sync requests , if any
	ReasonsForDeviceResync        string   `json:"reasonsForDeviceResync,omitempty"`        // Reason for last/ongoing sync
	LastDeviceResyncStartTime     string   `json:"lastDeviceResyncStartTime,omitempty"`     // Start time for last/ongoing sync
	UptimeSeconds                 *float64 `json:"uptimeSeconds,omitempty"`                 // Uptime in Seconds
	ManagedAtleastOnce            *bool    `json:"managedAtleastOnce,omitempty"`            // Indicates if device went into Managed state atleast once
	DeviceSupportLevel            string   `json:"deviceSupportLevel,omitempty"`            // Support level of the device
	ManagementState               string   `json:"managementState,omitempty"`               // Represents the current management state of the network element: managed, unmanaged, under maintenance, and so on.
	Description                   string   `json:"description,omitempty"`                   // System description
}
type ResponseDevicesDeleteDeviceByID struct {
	Response *ResponseDevicesDeleteDeviceByIDResponse `json:"response,omitempty"` //
	Version  string                                   `json:"version,omitempty"`  //
}
type ResponseDevicesDeleteDeviceByIDResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDevicesGetDeviceSummary struct {
	Response *ResponseDevicesGetDeviceSummaryResponse `json:"response,omitempty"` //
	Version  string                                   `json:"version,omitempty"`  //
}
type ResponseDevicesGetDeviceSummaryResponse struct {
	ID         string `json:"id,omitempty"`         // Unique identifier of the network device
	Role       string `json:"role,omitempty"`       // Role of device as access, distribution, border router, core
	RoleSource string `json:"roleSource,omitempty"` // Role source as manual / auto
}
type ResponseDevicesGetPollingIntervalByID struct {
	Response *int   `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}
type ResponseDevicesGetOrganizationListForMeraki struct {
	Response []string `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}
type ResponseDevicesGetDeviceInterfaceVLANs struct {
	Response *[]ResponseDevicesGetDeviceInterfaceVLANsResponse `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  //
}
type ResponseDevicesGetDeviceInterfaceVLANsResponse struct {
	InterfaceName  string `json:"interfaceName,omitempty"`  // Interface name
	IPAddress      string `json:"ipAddress,omitempty"`      // IP address
	Mask           *int   `json:"mask,omitempty"`           // Mask IP
	NetworkAddress string `json:"networkAddress,omitempty"` // Network addresses
	NumberOfIPs    *int   `json:"numberOfIPs,omitempty"`    // Number of Ip addresses
	Prefix         string `json:"prefix,omitempty"`         // Prefix associated with the IP address
	VLANNumber     *int   `json:"vlanNumber,omitempty"`     // Vlan Number
	VLANType       string `json:"vlanType,omitempty"`       // [Deprecated] Description of the interface VLAN
}
type ResponseDevicesGetWirelessLanControllerDetailsByID struct {
	AdminEnabledPorts        *[]int `json:"adminEnabledPorts,omitempty"`        // Admin Enabled Ports of the Device
	ApGroupName              string `json:"apGroupName,omitempty"`              // Name of the AP Group that Access point assigned
	DeviceID                 string `json:"deviceId,omitempty"`                 // Device Id
	EthMacAddress            string `json:"ethMacAddress,omitempty"`            // Ethernet MacAddress of the Device
	FlexGroupName            string `json:"flexGroupName,omitempty"`            // Name of the Flex Group that Access point assigned
	ID                       string `json:"id,omitempty"`                       // Id of the Device
	InstanceTenantID         string `json:"instanceTenantId,omitempty"`         // TenantId of the Device
	InstanceUUID             string `json:"instanceUuid,omitempty"`             // Instance UUID of the Device
	LagModeEnabled           *bool  `json:"lagModeEnabled,omitempty"`           // LagMode status of the Device
	NetconfEnabled           *bool  `json:"netconfEnabled,omitempty"`           // Netconf Status of the Device
	WirelessLicenseInfo      string `json:"wirelessLicenseInfo,omitempty"`      // License type of Wireless Device
	WirelessPackageInstalled *bool  `json:"wirelessPackageInstalled,omitempty"` // Status of the Wireless Package on the Device
}
type ResponseDevicesGetDeviceConfigByID struct {
	Response string `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}
type ResponseDevicesGetNetworkDeviceByPaginationRange struct {
	Response *[]ResponseDevicesGetNetworkDeviceByPaginationRangeResponse `json:"response,omitempty"` //
	Version  string                                                      `json:"version,omitempty"`  //
}
type ResponseDevicesGetNetworkDeviceByPaginationRangeResponse struct {
	ApManagerInterfaceIP          string   `json:"apManagerInterfaceIp,omitempty"`          // IP address of WLC on AP manager interface
	AssociatedWlcIP               string   `json:"associatedWlcIp,omitempty"`               // Associated Wlc Ip address of the AP device
	BootDateTime                  string   `json:"bootDateTime,omitempty"`                  // Device boot time
	CollectionInterval            string   `json:"collectionInterval,omitempty"`            // Re sync Interval of the device
	CollectionStatus              string   `json:"collectionStatus,omitempty"`              // Collection status as Synchronizing, Could not synchronize, Not manageable, Managed, Partial Collection Failure, Incomplete, Unreachable, Wrong credential, Reachable, In Progress
	ErrorCode                     string   `json:"errorCode,omitempty"`                     // Inventory status error code
	ErrorDescription              string   `json:"errorDescription,omitempty"`              // Inventory status description
	Family                        string   `json:"family,omitempty"`                        // Family of device as switch, router, wireless lan controller, accesspoints
	Hostname                      string   `json:"hostname,omitempty"`                      // Device name
	ID                            string   `json:"id,omitempty"`                            // Instance Uuid of the device
	InstanceTenantID              string   `json:"instanceTenantId,omitempty"`              // Instance Tenant Id of the device
	InstanceUUID                  string   `json:"instanceUuid,omitempty"`                  // Instance Uuid of the device
	InterfaceCount                string   `json:"interfaceCount,omitempty"`                // Number of interfaces on the device
	InventoryStatusDetail         string   `json:"inventoryStatusDetail,omitempty"`         // Status detail of inventory sync
	LastUpdateTime                *float64 `json:"lastUpdateTime,omitempty"`                // Time in epoch when the network device info last got updated
	LastUpdated                   string   `json:"lastUpdated,omitempty"`                   // Time when the network device info last got updated
	LineCardCount                 string   `json:"lineCardCount,omitempty"`                 // Number of linecards on the device
	LineCardID                    string   `json:"lineCardId,omitempty"`                    // IDs of linecards of the device
	Location                      string   `json:"location,omitempty"`                      // [Deprecated] Location ID that is associated with the device
	LocationName                  string   `json:"locationName,omitempty"`                  // [Deprecated] Name of the associated location
	MacAddress                    string   `json:"macAddress,omitempty"`                    // MAC address of device
	ManagementIPAddress           string   `json:"managementIpAddress,omitempty"`           // IP address of the device
	MemorySize                    string   `json:"memorySize,omitempty"`                    // Processor memory size
	PlatformID                    string   `json:"platformId,omitempty"`                    // Platform ID of device
	ReachabilityFailureReason     string   `json:"reachabilityFailureReason,omitempty"`     // Failure reason for unreachable devices
	ReachabilityStatus            string   `json:"reachabilityStatus,omitempty"`            // Device reachability status as Reachable / Unreachable
	Role                          string   `json:"role,omitempty"`                          // Role of device as access, distribution, border router, core
	RoleSource                    string   `json:"roleSource,omitempty"`                    // Role source as manual / auto
	SerialNumber                  string   `json:"serialNumber,omitempty"`                  // Serial number of device
	Series                        string   `json:"series,omitempty"`                        // Device series
	SNMPContact                   string   `json:"snmpContact,omitempty"`                   // SNMP contact on device
	SNMPLocation                  string   `json:"snmpLocation,omitempty"`                  // SNMP location on device
	SoftwareType                  string   `json:"softwareType,omitempty"`                  // Software type on the device
	SoftwareVersion               string   `json:"softwareVersion,omitempty"`               // Software version on the device
	TagCount                      string   `json:"tagCount,omitempty"`                      // Number of tags associated with the device
	TunnelUDPPort                 string   `json:"tunnelUdpPort,omitempty"`                 // Mobility protocol port is stored as tunneludpport for WLC
	Type                          string   `json:"type,omitempty"`                          // Type of device as switch, router, wireless lan controller, accesspoints
	UpTime                        string   `json:"upTime,omitempty"`                        // Time that shows for how long the device has been up
	WaasDeviceMode                string   `json:"waasDeviceMode,omitempty"`                // WAAS device mode
	DNSResolvedManagementAddress  string   `json:"dnsResolvedManagementAddress,omitempty"`  // Specifies the resolved ip address of dns name
	ApEthernetMacAddress          string   `json:"apEthernetMacAddress,omitempty"`          // AccessPoint Ethernet MacAddress of AP device
	Vendor                        string   `json:"vendor,omitempty"`                        // Vendor details
	ReasonsForPendingSyncRequests string   `json:"reasonsForPendingSyncRequests,omitempty"` // Reasons for pending sync requests , if any
	PendingSyncRequestsCount      string   `json:"pendingSyncRequestsCount,omitempty"`      // Count of pending sync requests , if any
	ReasonsForDeviceResync        string   `json:"reasonsForDeviceResync,omitempty"`        // Reason for last/ongoing sync
	LastDeviceResyncStartTime     string   `json:"lastDeviceResyncStartTime,omitempty"`     // Start time for last/ongoing sync
	UptimeSeconds                 *float64 `json:"uptimeSeconds,omitempty"`                 // Uptime in Seconds
	ManagedAtleastOnce            *bool    `json:"managedAtleastOnce,omitempty"`            // Indicates if device went into Managed state atleast once
	DeviceSupportLevel            string   `json:"deviceSupportLevel,omitempty"`            // Support level of the device
	ManagementState               string   `json:"managementState,omitempty"`               // Represents the current management state of the network element: managed, unmanaged, under maintenance, and so on.
	Description                   string   `json:"description,omitempty"`                   // System description
}
type ResponseDevicesUpdateGlobalResyncInterval struct {
	Response *ResponseDevicesUpdateGlobalResyncIntervalResponse `json:"response,omitempty"` //
	Version  string                                             `json:"version,omitempty"`  // Version of the response
}
type ResponseDevicesUpdateGlobalResyncIntervalResponse struct {
	TaskID string `json:"taskId,omitempty"` // Unique identifier for the task
	URL    string `json:"url,omitempty"`    // URL for the task
}
type ResponseDevicesOverrideResyncInterval struct {
	Response *ResponseDevicesOverrideResyncIntervalResponse `json:"response,omitempty"` //
	Version  string                                         `json:"version,omitempty"`  // Version of the response
}
type ResponseDevicesOverrideResyncIntervalResponse struct {
	TaskID string `json:"taskId,omitempty"` // Unique identifier for the task
	URL    string `json:"url,omitempty"`    // URL for the task
}
type ResponseDevicesUpdateResyncIntervalForTheNetworkDevice struct {
	Response *ResponseDevicesUpdateResyncIntervalForTheNetworkDeviceResponse `json:"response,omitempty"` //
	Version  string                                                          `json:"version,omitempty"`  // Version of the response
}
type ResponseDevicesUpdateResyncIntervalForTheNetworkDeviceResponse struct {
	TaskID string `json:"taskId,omitempty"` // Unique identifier for the task
	URL    string `json:"url,omitempty"`    // URL for the task
}
type ResponseDevicesGetResyncIntervalForTheNetworkDevice struct {
	Response *ResponseDevicesGetResyncIntervalForTheNetworkDeviceResponse `json:"response,omitempty"` //
	Version  string                                                       `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetResyncIntervalForTheNetworkDeviceResponse struct {
	Interval *int `json:"interval,omitempty"` // Resync interval of the device
}
type ResponseDevicesGetDeviceInterfaceStatsInfo struct {
	Version    string                                                `json:"version,omitempty"`    // Version
	TotalCount *float64                                              `json:"totalCount,omitempty"` // The total count
	Response   *[]ResponseDevicesGetDeviceInterfaceStatsInfoResponse `json:"response,omitempty"`   //
	Page       *ResponseDevicesGetDeviceInterfaceStatsInfoPage       `json:"page,omitempty"`       //
}
type ResponseDevicesGetDeviceInterfaceStatsInfoResponse struct {
	ID     string                                                    `json:"id,omitempty"`     // Interface Instance Id
	Values *ResponseDevicesGetDeviceInterfaceStatsInfoResponseValues `json:"values,omitempty"` //
}
type ResponseDevicesGetDeviceInterfaceStatsInfoResponseValues struct {
	AdminStatus     string   `json:"adminStatus,omitempty"`     // The desired state of the interface
	DeviceID        string   `json:"deviceId,omitempty"`        // Device Id
	DuplexConfig    string   `json:"duplexConfig,omitempty"`    // Interface duplex config status
	DuplexOper      string   `json:"duplexOper,omitempty"`      // Interface duplex operational status
	InterfaceID     string   `json:"interfaceId,omitempty"`     // Interface ifIndex
	InterfaceType   string   `json:"interfaceType,omitempty"`   // Physical or Virtual type
	InstanceID      string   `json:"instanceId,omitempty"`      // Interface InstanceId
	IPv4Address     string   `json:"ipv4Address,omitempty"`     // Interface IPV4 Address
	IPv6AddressList []string `json:"ipv6AddressList,omitempty"` // List of interface IPV6 Address
	IsL3Interface   string   `json:"isL3Interface,omitempty"`   // Interface is L3 or not
	IsWan           string   `json:"isWan,omitempty"`           // nterface is WAN link or not
	MacAddr         string   `json:"macAddr,omitempty"`         // Interface MAC Address
	MediaType       string   `json:"mediaType,omitempty"`       // Interface media type
	Name            string   `json:"name,omitempty"`            // Name of the interface
	OperStatus      string   `json:"operStatus,omitempty"`      // Interface operational status
	PeerStackMember string   `json:"peerStackMember,omitempty"` // Interface peer stack member Id
	PeerStackPort   string   `json:"peerStackPort,omitempty"`   // Interface peer stack member port
	PortChannelID   string   `json:"portChannelId,omitempty"`   // Interface Port-Channel Id
	PortMode        string   `json:"portMode,omitempty"`        // Interface Port Mode
	PortType        string   `json:"portType,omitempty"`        // Interface ifType
	Description     string   `json:"description,omitempty"`     // Interface description
	RxDiscards      string   `json:"rxDiscards,omitempty"`      // Rx Discards in %
	RxError         string   `json:"rxError,omitempty"`         // Rx Errors in %
	RxRate          string   `json:"rxRate,omitempty"`          // Rx rate in bps
	RxUtilization   string   `json:"rxUtilization,omitempty"`   // Rx Utilization in %
	Speed           string   `json:"speed,omitempty"`           // Speed of the Interface in kbps
	StackPortType   string   `json:"stackPortType,omitempty"`   // Interface stack port type. SVL or DAD
	Timestamp       string   `json:"timestamp,omitempty"`       // Interface stats collected timestamp
	TxDiscards      string   `json:"txDiscards,omitempty"`      // Tx Discards in %
	TxError         string   `json:"txError,omitempty"`         // Tx Errors in %
	TxRate          string   `json:"txRate,omitempty"`          // Tx Rate in bps
	TxUtilization   string   `json:"txUtilization,omitempty"`   // Tx  Utilization in %
	VLANID          string   `json:"vlanId,omitempty"`          // Interface VLAN Id
}
type ResponseDevicesGetDeviceInterfaceStatsInfoPage struct {
	Limit  *int     `json:"limit,omitempty"`  // Limit
	Offset *float64 `json:"offset,omitempty"` // Offset
	Count  *int     `json:"count,omitempty"`  // Count
}
type ResponseDevicesGetTheCountOfHealthScoreDefinitionsBasedOnProvidedFilters struct {
	Response *ResponseDevicesGetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersResponse `json:"response,omitempty"` //
	Version  string                                                                            `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions struct {
	StartTime           *int                                                                                                                        `json:"startTime,omitempty"`           // Start Time
	EndTime             *int                                                                                                                        `json:"endTime,omitempty"`             // End Time
	Views               []string                                                                                                                    `json:"views,omitempty"`               // Views
	Attributes          []string                                                                                                                    `json:"attributes,omitempty"`          // Attributes
	Filters             *[]RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFilters             `json:"filters,omitempty"`             //
	AggregateAttributes *[]RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
	Page                *RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPage                  `json:"page,omitempty"`                //
}
type RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFilters struct {
	Key      string `json:"key,omitempty"`      // Key
	Operator string `json:"operator,omitempty"` // Operator
	Value    string `json:"value,omitempty"`    // Value
}
type RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPage struct {
	Limit  *int                                                                                                               `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                                               `json:"offset,omitempty"` // Offset
	SortBy *[]RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPageSortBy `json:"sortBy,omitempty"` //
}
type RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type RequestDevicesQueryAssuranceEventsWithFilters struct {
	DeviceFamily []string                                                `json:"deviceFamily,omitempty"` // Device Family
	StartTime    *int                                                    `json:"startTime,omitempty"`    // Start Time
	EndTime      *int                                                    `json:"endTime,omitempty"`      // End Time
	Attributes   []string                                                `json:"attributes,omitempty"`   // Attributes
	Views        []string                                                `json:"views,omitempty"`        // Views
	Filters      *[]RequestDevicesQueryAssuranceEventsWithFiltersFilters `json:"filters,omitempty"`      //
	Page         *RequestDevicesQueryAssuranceEventsWithFiltersPage      `json:"page,omitempty"`         //
}
type RequestDevicesQueryAssuranceEventsWithFiltersFilters struct {
	Key      string `json:"key,omitempty"`      // Key
	Operator string `json:"operator,omitempty"` // Operator
	Value    string `json:"value,omitempty"`    // Value
}
type RequestDevicesQueryAssuranceEventsWithFiltersPage struct {
	Offset *int                                                       `json:"offset,omitempty"` // Offset
	Limit  *int                                                       `json:"limit,omitempty"`  // Limit
	SortBy *[]RequestDevicesQueryAssuranceEventsWithFiltersPageSortBy `json:"sortBy,omitempty"` //
}
type RequestDevicesQueryAssuranceEventsWithFiltersPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type RequestDevicesCountTheNumberOfEventsWithFilters struct {
	DeviceFamily []string                                                  `json:"deviceFamily,omitempty"` // Device Family
	StartTime    *int                                                      `json:"startTime,omitempty"`    // Start Time
	EndTime      *int                                                      `json:"endTime,omitempty"`      // End Time
	Filters      *[]RequestDevicesCountTheNumberOfEventsWithFiltersFilters `json:"filters,omitempty"`      //
}
type RequestDevicesCountTheNumberOfEventsWithFiltersFilters struct {
	Key      string `json:"key,omitempty"`      // Key
	Operator string `json:"operator,omitempty"` // Operator
	Value    string `json:"value,omitempty"`    // Value
}
type RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions struct {
	StartTime           *int                                                                                                                                      `json:"startTime,omitempty"`           // Start Time
	EndTime             *int                                                                                                                                      `json:"endTime,omitempty"`             // End Time
	Views               []string                                                                                                                                  `json:"views,omitempty"`               // Views
	Attributes          []string                                                                                                                                  `json:"attributes,omitempty"`          // Attributes
	Filters             *[]RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFilters             `json:"filters,omitempty"`             //
	AggregateAttributes *[]RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
	Page                *RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPage                  `json:"page,omitempty"`                //
}
type RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFilters struct {
	Key             string                                                                                                                               `json:"key,omitempty"`             // Key
	Operator        string                                                                                                                               `json:"operator,omitempty"`        // Operator
	LogicalOperator string                                                                                                                               `json:"logicalOperator,omitempty"` // Logical Operator
	Value           *RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFiltersValue     `json:"value,omitempty"`           // Value
	Filters         *[]RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFiltersFilters `json:"filters,omitempty"`         //
}
type RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFiltersValue interface{}
type RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFiltersFilters struct {
	Key             string                                                                                                                                  `json:"key,omitempty"`             // Key
	Operator        string                                                                                                                                  `json:"operator,omitempty"`        // Operator
	LogicalOperator string                                                                                                                                  `json:"logicalOperator,omitempty"` // Logical Operator
	Value           *RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFiltersFiltersValue `json:"value,omitempty"`           // Value
	Filters         []string                                                                                                                                `json:"filters,omitempty"`         // Filters
}
type RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFiltersFiltersValue interface{}
type RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPage struct {
	Limit  *int                                                                                                                             `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                                                             `json:"offset,omitempty"` // Offset
	SortBy *[]RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPageSortBy `json:"sortBy,omitempty"` //
}
type RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevices struct {
	StartTime           *int                                                                               `json:"startTime,omitempty"`           // Start Time
	EndTime             *int                                                                               `json:"endTime,omitempty"`             // End Time
	Views               []string                                                                           `json:"views,omitempty"`               // Views
	Attributes          []string                                                                           `json:"attributes,omitempty"`          // Attributes
	Filters             *[]RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesFilters             `json:"filters,omitempty"`             //
	AggregateAttributes *[]RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesAggregateAttributes `json:"aggregateAttributes,omitempty"` //
	Page                *RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesPage                  `json:"page,omitempty"`                //
}
type RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesFilters struct {
	Key             string                                                                    `json:"key,omitempty"`             // Key
	Operator        string                                                                    `json:"operator,omitempty"`        // Operator
	LogicalOperator string                                                                    `json:"logicalOperator,omitempty"` // Logical Operator
	Value           *RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesFiltersValue `json:"value,omitempty"`           // Value
	Filters         []string                                                                  `json:"filters,omitempty"`         // Filters
}
type RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesFiltersValue interface{}
type RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesPage struct {
	Limit  *int                                                                      `json:"limit,omitempty"`  // Limit
	Offset *int                                                                      `json:"offset,omitempty"` // Offset
	SortBy *[]RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesPageSortBy `json:"sortBy,omitempty"` //
}
type RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions struct {
	StartTime           *int                                                                                                                   `json:"startTime,omitempty"`           // Start Time
	EndTime             *int                                                                                                                   `json:"endTime,omitempty"`             // End Time
	Views               []string                                                                                                               `json:"views,omitempty"`               // Views
	Attributes          []string                                                                                                               `json:"attributes,omitempty"`          // Attributes
	Filters             *[]RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFilters             `json:"filters,omitempty"`             //
	AggregateAttributes *[]RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
	Page                *RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPage                  `json:"page,omitempty"`                //
}
type RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFilters struct {
	Key      string `json:"key,omitempty"`      // Key
	Operator string `json:"operator,omitempty"` // Operator
	Value    string `json:"value,omitempty"`    // Value
}
type RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPage struct {
	Limit  *int   `json:"limit,omitempty"`  // Limit
	Offset *int   `json:"offset,omitempty"` // Offset
	Count  *int   `json:"count,omitempty"`  // Count
	SortBy string `json:"sortBy,omitempty"` // Sort By
}
type RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevices struct {
	StartTime           *int                                                                                   `json:"startTime,omitempty"`           // Start Time
	EndTime             *int                                                                                   `json:"endTime,omitempty"`             // End Time
	GroupBy             []string                                                                               `json:"groupBy,omitempty"`             // Group By
	Attributes          []string                                                                               `json:"attributes,omitempty"`          // Attributes
	Filters             *[]RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesFilters             `json:"filters,omitempty"`             //
	AggregateAttributes *[]RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesAggregateAttributes `json:"aggregateAttributes,omitempty"` //
	Page                *RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesPage                  `json:"page,omitempty"`                //
}
type RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesFilters struct {
	Key      string `json:"key,omitempty"`      // Key
	Operator string `json:"operator,omitempty"` // Operator
	Value    string `json:"value,omitempty"`    // Value
}
type RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesPage struct {
	Limit  *int                                                                          `json:"limit,omitempty"`  // Limit
	Offset *int                                                                          `json:"offset,omitempty"` // Offset
	SortBy *[]RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesPageSortBy `json:"sortBy,omitempty"` //
}
type RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type RequestDevicesGetsTheTrendAnalyticsData struct {
	StartTime           *int                                                          `json:"startTime,omitempty"`           // Start Time
	EndTime             *int                                                          `json:"endTime,omitempty"`             // End Time
	TrendInterval       string                                                        `json:"trendInterval,omitempty"`       // Trend Interval
	GroupBy             *[]RequestDevicesGetsTheTrendAnalyticsDataGroupBy             `json:"groupBy,omitempty"`             // Group By
	Attributes          []string                                                      `json:"attributes,omitempty"`          // Attributes
	Filters             *[]RequestDevicesGetsTheTrendAnalyticsDataFilters             `json:"filters,omitempty"`             //
	AggregateAttributes *[]RequestDevicesGetsTheTrendAnalyticsDataAggregateAttributes `json:"aggregateAttributes,omitempty"` // Aggregate Attributes
	Page                *RequestDevicesGetsTheTrendAnalyticsDataPage                  `json:"page,omitempty"`                //
}
type RequestDevicesGetsTheTrendAnalyticsDataGroupBy interface{}
type RequestDevicesGetsTheTrendAnalyticsDataFilters struct {
	Key      string `json:"key,omitempty"`      // Key
	Operator string `json:"operator,omitempty"` // Operator
	Value    string `json:"value,omitempty"`    // Value
}
type RequestDevicesGetsTheTrendAnalyticsDataAggregateAttributes interface{}
type RequestDevicesGetsTheTrendAnalyticsDataPage struct {
	Limit          *int   `json:"limit,omitempty"`          // Limit
	Offset         *int   `json:"offset,omitempty"`         // Offset
	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRange struct {
	StartTime              *int                                                                                                `json:"startTime,omitempty"`              // Start Time
	EndTime                *int                                                                                                `json:"endTime,omitempty"`                // End Time
	TrendIntervalInMinutes *int                                                                                                `json:"trendIntervalInMinutes,omitempty"` // Trend Interval In Minutes
	GroupBy                []string                                                                                            `json:"groupBy,omitempty"`                // Group By
	Filters                *[]RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeFilters             `json:"filters,omitempty"`                //
	Attributes             []string                                                                                            `json:"attributes,omitempty"`             // Attributes
	AggregateAttributes    *[]RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeAggregateAttributes `json:"aggregateAttributes,omitempty"`    //
	Page                   *RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangePage                  `json:"page,omitempty"`                   //
}
type RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeFilters struct {
	Key             string                                                                                     `json:"key,omitempty"`             // Key
	Operator        string                                                                                     `json:"operator,omitempty"`        // Operator
	LogicalOperator string                                                                                     `json:"logicalOperator,omitempty"` // Logical Operator
	Value           *RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeFiltersValue `json:"value,omitempty"`           // Value
	Filters         []string                                                                                   `json:"filters,omitempty"`         // Filters
}
type RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeFiltersValue interface{}
type RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangePage struct {
	Limit          *int   `json:"limit,omitempty"`          // Limit
	Offset         *int   `json:"offset,omitempty"`         // Offset
	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type RequestDevicesUpdatePlannedAccessPointForFloor struct {
	Attributes *RequestDevicesUpdatePlannedAccessPointForFloorAttributes `json:"attributes,omitempty"` //
	IsSensor   *bool                                                     `json:"isSensor,omitempty"`   // Indicates that PAP is a sensor
	Location   *RequestDevicesUpdatePlannedAccessPointForFloorLocation   `json:"location,omitempty"`   //
	Position   *RequestDevicesUpdatePlannedAccessPointForFloorPosition   `json:"position,omitempty"`   //
	RadioCount *int                                                      `json:"radioCount,omitempty"` // Number of radios of the planned access point
	Radios     *[]RequestDevicesUpdatePlannedAccessPointForFloorRadios   `json:"radios,omitempty"`     //
}
type RequestDevicesUpdatePlannedAccessPointForFloorAttributes struct {
	CreateDate    *float64 `json:"createDate,omitempty"`    // Created date of the planned access point
	Domain        string   `json:"domain,omitempty"`        // Service domain to which the planned access point belongs
	HeirarchyName string   `json:"heirarchyName,omitempty"` // Hierarchy name of the planned access point
	ID            *float64 `json:"id,omitempty"`            // Unique id of the planned access point
	InstanceUUID  string   `json:"instanceUuid,omitempty"`  // Instance uuid of the planned access point
	MacAddress    string   `json:"macAddress,omitempty"`    // MAC address of the planned access point
	Name          string   `json:"name,omitempty"`          // Display name of the planned access point
	Source        string   `json:"source,omitempty"`        // Source of the data used to create the planned access point
	TypeString    string   `json:"typeString,omitempty"`    // Type string representation of the planned access point
}
type RequestDevicesUpdatePlannedAccessPointForFloorLocation struct {
	Altitude   *float64 `json:"altitude,omitempty"`   // Altitude of the planned access point's location
	Lattitude  *float64 `json:"lattitude,omitempty"`  // Latitude of the planned access point's location
	Longtitude *float64 `json:"longtitude,omitempty"` // Longitude of the planned access point's location
}
type RequestDevicesUpdatePlannedAccessPointForFloorPosition struct {
	X *float64 `json:"x,omitempty"` // x-coordinate of the planned access point on the map, 0,0 point being the top-left corner
	Y *float64 `json:"y,omitempty"` // y-coordinate of the planned access point on the map, 0,0 point being the top-left corner
	Z *float64 `json:"z,omitempty"` // z-coordinate, or height, of the planned access point on the map
}
type RequestDevicesUpdatePlannedAccessPointForFloorRadios struct {
	Antenna    *RequestDevicesUpdatePlannedAccessPointForFloorRadiosAntenna    `json:"antenna,omitempty"`    //
	Attributes *RequestDevicesUpdatePlannedAccessPointForFloorRadiosAttributes `json:"attributes,omitempty"` //
	IsSensor   *bool                                                           `json:"isSensor,omitempty"`   // Determines if it is sensor or not
}
type RequestDevicesUpdatePlannedAccessPointForFloorRadiosAntenna struct {
	AzimuthAngle   *float64 `json:"azimuthAngle,omitempty"`   // Azimuth angle of the antenna
	ElevationAngle *float64 `json:"elevationAngle,omitempty"` // Elevation angle of the antenna
	Gain           *float64 `json:"gain,omitempty"`           // Gain of the antenna
	Mode           string   `json:"mode,omitempty"`           // Mode of the antenna associated with this radio
	Name           string   `json:"name,omitempty"`           // Name of the antenna
	Type           string   `json:"type,omitempty"`           // Type of the antenna associated with this radio
}
type RequestDevicesUpdatePlannedAccessPointForFloorRadiosAttributes struct {
	Channel       *float64 `json:"channel,omitempty"`       // Channel in which this radio operates
	ChannelString string   `json:"channelString,omitempty"` // Channel string representation
	ID            *int     `json:"id,omitempty"`            // Id of the radio
	IfMode        string   `json:"ifMode,omitempty"`        // IF mode of the radio
	IfTypeString  string   `json:"ifTypeString,omitempty"`  // String representation of native band
	IfTypeSubband string   `json:"ifTypeSubband,omitempty"` // Sub band of the radio
	InstanceUUID  string   `json:"instanceUuid,omitempty"`  // Instance Uuid of the radio
	SlotID        *float64 `json:"slotId,omitempty"`        // Slot number in which the radio resides in the parent access point
	TxPowerLevel  *float64 `json:"txPowerLevel,omitempty"`  // Tx Power at which this radio operates (in dBm)
}
type RequestDevicesCreatePlannedAccessPointForFloor struct {
	Attributes *RequestDevicesCreatePlannedAccessPointForFloorAttributes `json:"attributes,omitempty"` //
	IsSensor   *bool                                                     `json:"isSensor,omitempty"`   // Indicates that PAP is a sensor
	Location   *RequestDevicesCreatePlannedAccessPointForFloorLocation   `json:"location,omitempty"`   //
	Position   *RequestDevicesCreatePlannedAccessPointForFloorPosition   `json:"position,omitempty"`   //
	RadioCount *int                                                      `json:"radioCount,omitempty"` // Number of radios of the planned access point
	Radios     *[]RequestDevicesCreatePlannedAccessPointForFloorRadios   `json:"radios,omitempty"`     //
}
type RequestDevicesCreatePlannedAccessPointForFloorAttributes struct {
	CreateDate    *float64 `json:"createDate,omitempty"`    // Created date of the planned access point
	Domain        string   `json:"domain,omitempty"`        // Service domain to which the planned access point belongs
	HeirarchyName string   `json:"heirarchyName,omitempty"` // Hierarchy name of the planned access point
	ID            *float64 `json:"id,omitempty"`            // Unique id of the planned access point
	InstanceUUID  string   `json:"instanceUuid,omitempty"`  // Instance uuid of the planned access point
	MacAddress    string   `json:"macAddress,omitempty"`    // MAC address of the planned access point
	Name          string   `json:"name,omitempty"`          // Display name of the planned access point
	Source        string   `json:"source,omitempty"`        // Source of the data used to create the planned access point
	TypeString    string   `json:"typeString,omitempty"`    // Type string representation of the planned access point
}
type RequestDevicesCreatePlannedAccessPointForFloorLocation struct {
	Altitude   *float64 `json:"altitude,omitempty"`   // Altitude of the planned access point's location
	Lattitude  *float64 `json:"lattitude,omitempty"`  // Latitude of the planned access point's location
	Longtitude *float64 `json:"longtitude,omitempty"` // Longitude of the planned access point's location
}
type RequestDevicesCreatePlannedAccessPointForFloorPosition struct {
	X *float64 `json:"x,omitempty"` // x-coordinate of the planned access point on the map, 0,0 point being the top-left corner
	Y *float64 `json:"y,omitempty"` // y-coordinate of the planned access point on the map, 0,0 point being the top-left corner
	Z *float64 `json:"z,omitempty"` // z-coordinate, or height, of the planned access point on the map
}
type RequestDevicesCreatePlannedAccessPointForFloorRadios struct {
	Antenna    *RequestDevicesCreatePlannedAccessPointForFloorRadiosAntenna    `json:"antenna,omitempty"`    //
	Attributes *RequestDevicesCreatePlannedAccessPointForFloorRadiosAttributes `json:"attributes,omitempty"` //
	IsSensor   *bool                                                           `json:"isSensor,omitempty"`   // Determines if it is sensor or not
}
type RequestDevicesCreatePlannedAccessPointForFloorRadiosAntenna struct {
	AzimuthAngle   *float64 `json:"azimuthAngle,omitempty"`   // Azimuth angle of the antenna
	ElevationAngle *float64 `json:"elevationAngle,omitempty"` // Elevation angle of the antenna
	Gain           *float64 `json:"gain,omitempty"`           // Gain of the antenna
	Mode           string   `json:"mode,omitempty"`           // Mode of the antenna associated with this radio
	Name           string   `json:"name,omitempty"`           // Name of the antenna
	Type           string   `json:"type,omitempty"`           // Type of the antenna associated with this radio
}
type RequestDevicesCreatePlannedAccessPointForFloorRadiosAttributes struct {
	Channel       *float64 `json:"channel,omitempty"`       // Channel in which this radio operates
	ChannelString string   `json:"channelString,omitempty"` // Channel string representation
	ID            *int     `json:"id,omitempty"`            // Id of the radio
	IfMode        string   `json:"ifMode,omitempty"`        // IF mode of the radio
	IfTypeString  string   `json:"ifTypeString,omitempty"`  // String representation of native band
	IfTypeSubband string   `json:"ifTypeSubband,omitempty"` // Sub band of the radio
	InstanceUUID  string   `json:"instanceUuid,omitempty"`  // Instance Uuid of the radio
	SlotID        *float64 `json:"slotId,omitempty"`        // Slot number in which the radio resides in the parent access point
	TxPowerLevel  *float64 `json:"txPowerLevel,omitempty"`  // Tx Power at which this radio operates (in dBm)
}
type RequestDevicesUpdateHealthScoreDefinitions []RequestItemDevicesUpdateHealthScoreDefinitions // Array of RequestDevicesUpdateHealthScoreDefinitions
type RequestItemDevicesUpdateHealthScoreDefinitions struct {
	ID                          string   `json:"id,omitempty"`                          // Id
	IncludeForOverallHealth     *bool    `json:"includeForOverallHealth,omitempty"`     // Include For Overall Health
	ThresholdValue              *float64 `json:"thresholdValue,omitempty"`              // Threshold Value
	SynchronizeToIssueThreshold *bool    `json:"synchronizeToIssueThreshold,omitempty"` // Synchronize To Issue Threshold
}
type RequestDevicesUpdateHealthScoreDefinitionForTheGivenID struct {
	IncludeForOverallHealth     *bool    `json:"includeForOverallHealth,omitempty"`     // Include For Overall Health
	ThresholdValue              *float64 `json:"thresholdValue,omitempty"`              // Thresehold Value
	SynchronizeToIssueThreshold *bool    `json:"synchronizeToIssueThreshold,omitempty"` // Synchronize To Issue Threshold
}
type RequestDevicesUpdateInterfaceDetails struct {
	Description string `json:"description,omitempty"` // Description for the Interface
	AdminStatus string `json:"adminStatus,omitempty"` // Admin status as ('UP'/'DOWN')
	VLANID      *int   `json:"vlanId,omitempty"`      // VLAN Id to be Updated
	VoiceVLANID *int   `json:"voiceVlanId,omitempty"` // Voice Vlan Id to be Updated
}
type RequestDevicesClearMacAddressTable struct {
	Operation string                                     `json:"operation,omitempty"` // Operation needs to be specified as 'ClearMacAddress'.
	Payload   *RequestDevicesClearMacAddressTablePayload `json:"payload,omitempty"`   // Payload is not applicable
}
type RequestDevicesClearMacAddressTablePayload interface{}
type RequestDevicesAddDevice2 struct {
	CliTransport            string                                             `json:"cliTransport,omitempty"`            // CLI transport. Supported values: telnet, ssh. Required if type is NETWORK_DEVICE.
	ComputeDevice           *bool                                              `json:"computeDevice,omitempty"`           // Compute Device or not. Options are true / false.
	EnablePassword          string                                             `json:"enablePassword,omitempty"`          // CLI enable password of the device. Required if device is configured to use enable password.
	ExtendedDiscoveryInfo   string                                             `json:"extendedDiscoveryInfo,omitempty"`   // This field holds that info as whether to add device with canned data or not. Supported values: DISCOVER_WITH_CANNED_DATA.
	HTTPPassword            string                                             `json:"httpPassword,omitempty"`            // HTTP password of the device / API key for Meraki Dashboard. Required if type is MERAKI_DASHBOARD or COMPUTE_DEVICE.
	HTTPPort                string                                             `json:"httpPort,omitempty"`                // HTTP port of the device. Required if type is COMPUTE_DEVICE.
	HTTPSecure              *bool                                              `json:"httpSecure,omitempty"`              // Flag to select HTTP / HTTPS protocol. Options are true / false. true for HTTPS and false for HTTP. Default is true.
	HTTPUserName            string                                             `json:"httpUserName,omitempty"`            // HTTP Username of the device. Required if type is COMPUTE_DEVICE.
	IPAddress               []string                                           `json:"ipAddress,omitempty"`               // IP Address of the device. Required if type is NETWORK_DEVICE, COMPUTE_DEVICE or THIRD_PARTY_DEVICE.
	MerakiOrgID             []string                                           `json:"merakiOrgId,omitempty"`             // Selected Meraki organization for which the devices needs to be imported. Required if type is MERAKI_DASHBOARD.
	NetconfPort             string                                             `json:"netconfPort,omitempty"`             // Netconf Port of the device. cliTransport must be 'ssh' if netconf is provided.
	Password                string                                             `json:"password,omitempty"`                // CLI Password of the device. Required if type is NETWORK_DEVICE.
	SerialNumber            string                                             `json:"serialNumber,omitempty"`            // Serial Number of the Device. Required if extendedDiscoveryInfo is 'DISCOVER_WITH_CANNED_DATA'.
	SNMPAuthPassphrase      string                                             `json:"snmpAuthPassphrase,omitempty"`      // SNMPv3 auth passphrase of the device. Required if snmpMode is authNoPriv or authPriv.
	SNMPAuthProtocol        string                                             `json:"snmpAuthProtocol,omitempty"`        // SNMPv3 auth protocol. Supported values: sha, md5. Required if snmpMode is authNoPriv or authPriv.
	SNMPMode                string                                             `json:"snmpMode,omitempty"`                // SNMPv3 mode. Supported values: noAuthnoPriv, authNoPriv, authPriv. Required if snmpVersion is v3.
	SNMPPrivPassphrase      string                                             `json:"snmpPrivPassphrase,omitempty"`      // SNMPv3 priv passphrase. Required if snmpMode is authPriv.
	SNMPPrivProtocol        string                                             `json:"snmpPrivProtocol,omitempty"`        // SNMPv3 priv protocol. Supported values: AES128. Required if snmpMode is authPriv.
	SNMPROCommunity         string                                             `json:"snmpROCommunity,omitempty"`         // SNMP Read Community of the device. If snmpVersion is v2, at least one of snmpROCommunity and snmpRWCommunity is required.
	SNMPRWCommunity         string                                             `json:"snmpRWCommunity,omitempty"`         // SNMP Write Community of the device. If snmpVersion is v2, at least one of snmpROCommunity and snmpRWCommunity is required.
	SNMPRetry               *int                                               `json:"snmpRetry,omitempty"`               // SNMP retry count. Max value supported is 3. Default is Global SNMP retry (if exists) or 3.
	SNMPTimeout             *int                                               `json:"snmpTimeout,omitempty"`             // SNMP timeout in seconds. Max value supported is 300. Default is Global SNMP timeout (if exists) or 5.
	SNMPUserName            string                                             `json:"snmpUserName,omitempty"`            // SNMPV3 user name of the device. Required if snmpVersion is v3.
	SNMPVersion             string                                             `json:"snmpVersion,omitempty"`             // SNMP version. Values supported: v2, v3. Required if type is NETWORK_DEVICE, COMPUTE_DEVICE or THIRD_PARTY_DEVICE.
	Type                    string                                             `json:"type,omitempty"`                    // Type of device being added. Default is NETWORK_DEVICE.
	UpdateMgmtIPaddressList *[]RequestDevicesAddDevice2UpdateMgmtIPaddressList `json:"updateMgmtIPaddressList,omitempty"` //
	UserName                string                                             `json:"userName,omitempty"`                // CLI user name of the device. Required if type is NETWORK_DEVICE.
}
type RequestDevicesAddDevice2UpdateMgmtIPaddressList struct {
	ExistMgmtIPAddress string `json:"existMgmtIpAddress,omitempty"` //
	NewMgmtIPAddress   string `json:"newMgmtIpAddress,omitempty"`   //
}
type RequestDevicesUpdateDeviceDetails struct {
	CliTransport            string                                                      `json:"cliTransport,omitempty"`            // CLI transport. Supported values: telnet, ssh. Use NO!$DATA!$ if no change is required. Required if type is NETWORK_DEVICE.
	ComputeDevice           *bool                                                       `json:"computeDevice,omitempty"`           // Compute Device or not. Options are true / false.
	EnablePassword          string                                                      `json:"enablePassword,omitempty"`          // CLI enable password of the device. Required if device is configured to use enable password. Use NO!$DATA!$ if no change is required.
	ExtendedDiscoveryInfo   string                                                      `json:"extendedDiscoveryInfo,omitempty"`   // This field holds that info as whether to add device with canned data or not. Supported values: DISCOVER_WITH_CANNED_DATA.
	HTTPPassword            string                                                      `json:"httpPassword,omitempty"`            // HTTP password of the device / API key for Meraki Dashboard. Required if type is MERAKI_DASHBOARD or COMPUTE_DEVICE. Use NO!$DATA!$ if no change is required.
	HTTPPort                string                                                      `json:"httpPort,omitempty"`                // HTTP port of the device. Required if type is COMPUTE_DEVICE.
	HTTPSecure              *bool                                                       `json:"httpSecure,omitempty"`              // Flag to select HTTP / HTTPS protocol. Options are true / false. true for HTTPS and false for HTTP.
	HTTPUserName            string                                                      `json:"httpUserName,omitempty"`            // HTTP Username of the device. Required if type is COMPUTE_DEVICE. Use NO!$DATA!$ if no change is required.
	IPAddress               []string                                                    `json:"ipAddress,omitempty"`               // IP Address of the device. Required. Use 'api.meraki.com' for Meraki Dashboard.
	MerakiOrgID             []string                                                    `json:"merakiOrgId,omitempty"`             // Selected Meraki organization for which the devices needs to be imported. Required if type is MERAKI_DASHBOARD.
	NetconfPort             string                                                      `json:"netconfPort,omitempty"`             // Netconf Port of the device. cliTransport must be 'ssh' if netconf is provided.
	Password                string                                                      `json:"password,omitempty"`                // CLI Password of the device. Required if type is NETWORK_DEVICE. Use NO!$DATA!$ if no change is required.
	SerialNumber            string                                                      `json:"serialNumber,omitempty"`            // Serial Number of the Device. Required if extendedDiscoveryInfo is 'DISCOVER_WITH_CANNED_DATA'.
	SNMPAuthPassphrase      string                                                      `json:"snmpAuthPassphrase,omitempty"`      // SNMPv3 auth passphrase of the device. Required if snmpMode is authNoPriv or authPriv. Use NO!$DATA!$ if no change is required.
	SNMPAuthProtocol        string                                                      `json:"snmpAuthProtocol,omitempty"`        // SNMPv3 auth protocol. Supported values: sha, md5.  Required if snmpMode is authNoPriv or authPriv. Use NODATACHANGE if no change is required.
	SNMPMode                string                                                      `json:"snmpMode,omitempty"`                // SNMPv3 mode. Supported values: noAuthnoPriv, authNoPriv, authPriv. Required if snmpVersion is v3. Use NODATACHANGE if no change is required.
	SNMPPrivPassphrase      string                                                      `json:"snmpPrivPassphrase,omitempty"`      // SNMPv3 priv passphrase. Required if snmpMode is authPriv. Use NO!$DATA!$ if no change is required.
	SNMPPrivProtocol        string                                                      `json:"snmpPrivProtocol,omitempty"`        // SNMPv3 priv protocol. Supported values: AES128. Required if snmpMode is authPriv. Use NODATACHANGE if no change is required.
	SNMPROCommunity         string                                                      `json:"snmpROCommunity,omitempty"`         // SNMP Read Community of the device. If snmpVersion is v2, at least one of snmpROCommunity and snmpRWCommunity is required. Use NO!$DATA!$ if no change is required.
	SNMPRWCommunity         string                                                      `json:"snmpRWCommunity,omitempty"`         // SNMP Write Community of the device. If snmpVersion is v2, at least one of snmpROCommunity and snmpRWCommunity is required. Use NO!$DATA!$ if no change is required.
	SNMPRetry               *int                                                        `json:"snmpRetry,omitempty"`               // SNMP retry count. Max value supported is 3. Default is Global SNMP retry (if exists) or 3.
	SNMPTimeout             *int                                                        `json:"snmpTimeout,omitempty"`             // SNMP timeout in seconds. Max value supported is 300. Default is Global SNMP timeout (if exists) or 5.
	SNMPUserName            string                                                      `json:"snmpUserName,omitempty"`            // SNMPV3 user name of the device. Required if snmpVersion is v3. Use NO!$DATA!$ if no change is required.
	SNMPVersion             string                                                      `json:"snmpVersion,omitempty"`             // SNMP version. Values supported: v2, v3. Required if type is NETWORK_DEVICE, COMPUTE_DEVICE or THIRD_PARTY_DEVICE. Use NODATACHANGE if no change is required.
	Type                    string                                                      `json:"type,omitempty"`                    // Type of device being edited. Default is NETWORK_DEVICE.
	UpdateMgmtIPaddressList *[]RequestDevicesUpdateDeviceDetailsUpdateMgmtIPaddressList `json:"updateMgmtIPaddressList,omitempty"` //
	UserName                string                                                      `json:"userName,omitempty"`                // CLI user name of the device. Required if type is NETWORK_DEVICE. Use NO!$DATA!$ if no change is required.
}
type RequestDevicesUpdateDeviceDetailsUpdateMgmtIPaddressList struct {
	ExistMgmtIPAddress string `json:"existMgmtIpAddress,omitempty"` // existMgmtIpAddress IP Address of the device.
	NewMgmtIPAddress   string `json:"newMgmtIpAddress,omitempty"`   // New IP Address to be Updated.
}
type RequestDevicesUpdateDeviceRole struct {
	ID         string `json:"id,omitempty"`         // DeviceId of the Device
	Role       string `json:"role,omitempty"`       // Role of device as ACCESS, CORE, DISTRIBUTION, BORDER ROUTER
	RoleSource string `json:"roleSource,omitempty"` // Role source as MANUAL / AUTO
}
type RequestDevicesExportDeviceList struct {
	DeviceUUIDs   []string `json:"deviceUuids,omitempty"`   // List of device uuids
	OperationEnum string   `json:"operationEnum,omitempty"` // 0 to export Device Credential Details Or 1 to export Device Details
	Parameters    []string `json:"parameters,omitempty"`    // List of device parameters that needs to be exported to file
	Password      string   `json:"password,omitempty"`      // Password is required when the operationEnum value is 0
}
type RequestDevicesSyncDevices []string // Array of RequestDevicesSyncDevices
type RequestDevicesCreateUserDefinedField struct {
	Name        string `json:"name,omitempty"`        // Name of UDF
	Description string `json:"description,omitempty"` // Description of UDF
}
type RequestDevicesUpdateUserDefinedField struct {
	Name        string `json:"name,omitempty"`        // Name of UDF
	Description string `json:"description,omitempty"` // Description of UDF
}
type RequestDevicesAddUserDefinedFieldToDevice []RequestItemDevicesAddUserDefinedFieldToDevice // Array of RequestDevicesAddUserDefinedFieldToDevice
type RequestItemDevicesAddUserDefinedFieldToDevice struct {
	Name  string `json:"name,omitempty"`  // Name of the User Defined Field
	Value string `json:"value,omitempty"` // Value of the User Defined Field that will be assigned to the device
}
type RequestDevicesUpdateDeviceManagementAddress struct {
	NewIP string `json:"newIP,omitempty"` // New IP Address of the device to be Updated
}
type RequestDevicesUpdateGlobalResyncInterval struct {
	Interval *int `json:"interval,omitempty"` // Resync Interval should be between 25 to 1440 minutes
}
type RequestDevicesUpdateResyncIntervalForTheNetworkDevice struct {
	Interval *int `json:"interval,omitempty"` // Resync interval in minutes. To disable periodic resync, set interval as '0'. To use global settings, set interval as 'null'.
}
type RequestDevicesGetDeviceInterfaceStatsInfo struct {
	StartTime *int                                            `json:"startTime,omitempty"` // UTC epoch timestamp in milliseconds
	EndTime   *int                                            `json:"endTime,omitempty"`   // UTC epoch timestamp in milliseconds
	Query     *RequestDevicesGetDeviceInterfaceStatsInfoQuery `json:"query,omitempty"`     //
}
type RequestDevicesGetDeviceInterfaceStatsInfoQuery struct {
	Fields  *[]RequestDevicesGetDeviceInterfaceStatsInfoQueryFields  `json:"fields,omitempty"`  // Required field names, default ALL
	Filters *[]RequestDevicesGetDeviceInterfaceStatsInfoQueryFilters `json:"filters,omitempty"` //
	Page    *RequestDevicesGetDeviceInterfaceStatsInfoQueryPage      `json:"page,omitempty"`    //
}
type RequestDevicesGetDeviceInterfaceStatsInfoQueryFields interface{}
type RequestDevicesGetDeviceInterfaceStatsInfoQueryFilters struct {
	Key      string `json:"key,omitempty"`      // Name of the field that the filter should be applied to
	Operator string `json:"operator,omitempty"` // Supported operators are eq,in,like
	Value    string `json:"value,omitempty"`    // Value of the field
}
type RequestDevicesGetDeviceInterfaceStatsInfoQueryPage struct {
	Limit   *int                                                         `json:"limit,omitempty"`   // Number of records, Max is 1000
	Offset  *float64                                                     `json:"offset,omitempty"`  // Record offset value, default 0
	OrderBy *[]RequestDevicesGetDeviceInterfaceStatsInfoQueryPageOrderBy `json:"orderBy,omitempty"` //
}
type RequestDevicesGetDeviceInterfaceStatsInfoQueryPageOrderBy struct {
	Name  string `json:"name,omitempty"`  // Name of the field used to sort
	Order string `json:"order,omitempty"` // Possible values asc, des
}

//QueryAssuranceEvents Query assurance events - 15a5-3b2c-4908-8ba3
/* Returns the list of events discovered by Catalyst Center, determined by the complex filters. Please refer to the 'API Support Documentation' section to understand which fields are supported. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceEvents-1.0.0-resolved.yaml


@param QueryAssuranceEventsHeaderParams Custom header parameters
@param QueryAssuranceEventsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!query-assurance-events
*/
func (s *DevicesService) QueryAssuranceEvents(QueryAssuranceEventsHeaderParams *QueryAssuranceEventsHeaderParams, QueryAssuranceEventsQueryParams *QueryAssuranceEventsQueryParams) (*ResponseDevicesQueryAssuranceEvents, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceEvents"

	queryString, _ := query.Values(QueryAssuranceEventsQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if QueryAssuranceEventsHeaderParams != nil {

		if QueryAssuranceEventsHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", QueryAssuranceEventsHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesQueryAssuranceEvents{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.QueryAssuranceEvents(QueryAssuranceEventsHeaderParams, QueryAssuranceEventsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation QueryAssuranceEvents")
	}

	result := response.Result().(*ResponseDevicesQueryAssuranceEvents)
	return result, response, err

}

//CountTheNumberOfEvents Count the number of events - 349f-a9d8-4a6a-b951
/* API to fetch the count of assurance events that match the filter criteria. Please refer to the 'API Support Documentation' section to understand which fields are supported. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceEvents-1.0.0-resolved.yaml


@param CountTheNumberOfEventsHeaderParams Custom header parameters
@param CountTheNumberOfEventsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!count-the-number-of-events
*/
func (s *DevicesService) CountTheNumberOfEvents(CountTheNumberOfEventsHeaderParams *CountTheNumberOfEventsHeaderParams, CountTheNumberOfEventsQueryParams *CountTheNumberOfEventsQueryParams) (*ResponseDevicesCountTheNumberOfEvents, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceEvents/count"

	queryString, _ := query.Values(CountTheNumberOfEventsQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if CountTheNumberOfEventsHeaderParams != nil {

		if CountTheNumberOfEventsHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", CountTheNumberOfEventsHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesCountTheNumberOfEvents{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.CountTheNumberOfEvents(CountTheNumberOfEventsHeaderParams, CountTheNumberOfEventsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation CountTheNumberOfEvents")
	}

	result := response.Result().(*ResponseDevicesCountTheNumberOfEvents)
	return result, response, err

}

//GetDetailsOfASingleAssuranceEvent Get details of a single assurance event - 039e-2909-449a-8f51
/* API to fetch the details of an assurance event using event 'id'. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceEvents-1.0.0-resolved.yaml


@param id id path parameter. Unique identifier for the event

@param GetDetailsOfASingleAssuranceEventHeaderParams Custom header parameters
@param GetDetailsOfASingleAssuranceEventQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-details-of-a-single-assurance-event
*/
func (s *DevicesService) GetDetailsOfASingleAssuranceEvent(id string, GetDetailsOfASingleAssuranceEventHeaderParams *GetDetailsOfASingleAssuranceEventHeaderParams, GetDetailsOfASingleAssuranceEventQueryParams *GetDetailsOfASingleAssuranceEventQueryParams) (*ResponseDevicesGetDetailsOfASingleAssuranceEvent, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceEvents/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetDetailsOfASingleAssuranceEventQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetDetailsOfASingleAssuranceEventHeaderParams != nil {

		if GetDetailsOfASingleAssuranceEventHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetDetailsOfASingleAssuranceEventHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetDetailsOfASingleAssuranceEvent{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDetailsOfASingleAssuranceEvent(id, GetDetailsOfASingleAssuranceEventHeaderParams, GetDetailsOfASingleAssuranceEventQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDetailsOfASingleAssuranceEvent")
	}

	result := response.Result().(*ResponseDevicesGetDetailsOfASingleAssuranceEvent)
	return result, response, err

}

//GetListOfChildEventsForTheGivenWirelessClientEvent Get list of child events for the given wireless client event - d78f-7acc-4a88-b616
/* Wireless client event could have child events and this API can be used to fetch the same using parent event 'id' as the input. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceEvents-1.0.0-resolved.yaml


@param id id path parameter. Unique identifier for the event

@param GetListOfChildEventsForTheGivenWirelessClientEventHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-list-of-child-events-for-the-given-wireless-client-event
*/
func (s *DevicesService) GetListOfChildEventsForTheGivenWirelessClientEvent(id string, GetListOfChildEventsForTheGivenWirelessClientEventHeaderParams *GetListOfChildEventsForTheGivenWirelessClientEventHeaderParams) (*ResponseDevicesGetListOfChildEventsForTheGivenWirelessClientEvent, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceEvents/{id}/childEvents"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetListOfChildEventsForTheGivenWirelessClientEventHeaderParams != nil {

		if GetListOfChildEventsForTheGivenWirelessClientEventHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetListOfChildEventsForTheGivenWirelessClientEventHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetResult(&ResponseDevicesGetListOfChildEventsForTheGivenWirelessClientEvent{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetListOfChildEventsForTheGivenWirelessClientEvent(id, GetListOfChildEventsForTheGivenWirelessClientEventHeaderParams)
		}
		return nil, response, fmt.Errorf("error with operation GetListOfChildEventsForTheGivenWirelessClientEvent")
	}

	result := response.Result().(*ResponseDevicesGetListOfChildEventsForTheGivenWirelessClientEvent)
	return result, response, err

}

//GetsInterfacesAlongWithStatisticsDataFromAllNetworkDevices Gets interfaces along with statistics data from all network devices. - 9898-9b5a-445b-884f
/* Retrieves the list of the interfaces from all network devices based on the provided query parameters. The latest interfaces data in the specified start and end time range will be returned. When there is no start and end time specified returns the latest available data.
The elements are grouped and sorted by deviceUuid first, and are then sorted by the given sort field, or by the default value: name.

 The supported sorting options are: name, adminStatus, description, duplexConfig, duplexOper,interfaceIfIndex,interfaceType, macAddress,mediaType, operStatus,portChannelId, portMode, portType,speed, vlanId. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-interfaces-1.0.2-resolved.yaml


@param GetsInterfacesAlongWithStatisticsDataFromAllNetworkDevicesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-interfaces-along-with-statistics-data-from-all-network-devices
*/
func (s *DevicesService) GetsInterfacesAlongWithStatisticsDataFromAllNetworkDevices(GetsInterfacesAlongWithStatisticsDataFromAllNetworkDevicesQueryParams *GetsInterfacesAlongWithStatisticsDataFromAllNetworkDevicesQueryParams) (*ResponseDevicesGetsInterfacesAlongWithStatisticsDataFromAllNetworkDevices, *resty.Response, error) {
	path := "/dna/data/api/v1/interfaces"

	queryString, _ := query.Values(GetsInterfacesAlongWithStatisticsDataFromAllNetworkDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetsInterfacesAlongWithStatisticsDataFromAllNetworkDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsInterfacesAlongWithStatisticsDataFromAllNetworkDevices(GetsInterfacesAlongWithStatisticsDataFromAllNetworkDevicesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetsInterfacesAlongWithStatisticsDataFromAllNetworkDevices")
	}

	result := response.Result().(*ResponseDevicesGetsInterfacesAlongWithStatisticsDataFromAllNetworkDevices)
	return result, response, err

}

//GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCount Gets the total Network device interface counts in the specified time range. When there is no start and end time specified returns the latest interfaces total count. - 40ab-799f-465a-82f4
/* Gets the total Network device interface counts. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-interfaces-1.0.2-resolved.yaml


@param GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-the-total-network-device-interface-counts-in-the-specified-time-range-when-there-is-no-start-and-end-time-specified-returns-the-latest-interfaces-total-count
*/
func (s *DevicesService) GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCount(GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountQueryParams *GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountQueryParams) (*ResponseDevicesGetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCount, *resty.Response, error) {
	path := "/dna/data/api/v1/interfaces/count"

	queryString, _ := query.Values(GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCount(GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCount")
	}

	result := response.Result().(*ResponseDevicesGetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCount)
	return result, response, err

}

//GetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsData Get the interface data for the given interface id (instance Uuid) along with the statistics data - c08d-d95c-4c7b-8283
/* Returns the interface data for the given interface instance Uuid along with the statistics data. The latest interface data in the specified start and end time range will be returned. When there is no start and end time specified returns the latest available data for the given interface Id. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-interfaces-1.0.2-resolved.yaml


@param id id path parameter. The interface Uuid

@param GetTheInterfaceDataForTheGivenInterfaceIdinstanceUuidAlongWithTheStatisticsDataQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-the-interface-data-for-the-given-interface-idinstance-uuid-along-with-the-statistics-data
*/
func (s *DevicesService) GetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsData(id string, GetTheInterfaceDataForTheGivenInterfaceIdinstanceUuidAlongWithTheStatisticsDataQueryParams *GetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsDataQueryParams) (*ResponseDevicesGetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsData, *resty.Response, error) {
	path := "/dna/data/api/v1/interfaces/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetTheInterfaceDataForTheGivenInterfaceIdinstanceUuidAlongWithTheStatisticsDataQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsData{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsData(id, GetTheInterfaceDataForTheGivenInterfaceIdinstanceUuidAlongWithTheStatisticsDataQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTheInterfaceDataForTheGivenInterfaceIdinstanceUuidAlongWithTheStatisticsData")
	}

	result := response.Result().(*ResponseDevicesGetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsData)
	return result, response, err

}

//GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParameters Gets the Network Device details based on the provided query parameters. - c8b4-f894-4c3a-932f
/* Gets the Network Device details based on the provided query parameters.  When there is no start and end time specified returns the latest device details. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceNetworkDevices-1.0.2-resolved.yaml


@param GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-the-network-device-details-based-on-the-provided-query-parameters
*/
func (s *DevicesService) GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParameters(GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersQueryParams *GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersQueryParams) (*ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParameters, *resty.Response, error) {
	path := "/dna/data/api/v1/networkDevices"

	queryString, _ := query.Values(GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParameters{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParameters(GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParameters")
	}

	result := response.Result().(*ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParameters)
	return result, response, err

}

//GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParameters Gets the total Network device counts based on the provided query parameters. - f0a6-e96b-44fb-a549
/* Gets the total Network device counts. When there is no start and end time specified returns the latest interfaces total count. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceNetworkDevices-1.0.2-resolved.yaml


@param GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-the-total-network-device-counts-based-on-the-provided-query-parameters
*/
func (s *DevicesService) GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParameters(GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersQueryParams *GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersQueryParams) (*ResponseDevicesGetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParameters, *resty.Response, error) {
	path := "/dna/data/api/v1/networkDevices/count"

	queryString, _ := query.Values(GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParameters{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParameters(GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParameters")
	}

	result := response.Result().(*ResponseDevicesGetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParameters)
	return result, response, err

}

//GetTheDeviceDataForTheGivenDeviceIDUUID Get the device data for the given device id (Uuid) - 5a93-1957-475b-95b3
/* Returns the device data for the given device Uuid in the specified start and end time range. When there is no start and end time specified returns the latest available data for the given Id. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceNetworkDevices-1.0.2-resolved.yaml


@param id id path parameter. The device Uuid

@param GetTheDeviceDataForTheGivenDeviceIdUuidQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-the-device-data-for-the-given-device-id-uuid
*/
func (s *DevicesService) GetTheDeviceDataForTheGivenDeviceIDUUID(id string, GetTheDeviceDataForTheGivenDeviceIdUuidQueryParams *GetTheDeviceDataForTheGivenDeviceIDUUIDQueryParams) (*ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUID, *resty.Response, error) {
	path := "/dna/data/api/v1/networkDevices/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetTheDeviceDataForTheGivenDeviceIdUuidQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTheDeviceDataForTheGivenDeviceIDUUID(id, GetTheDeviceDataForTheGivenDeviceIdUuidQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTheDeviceDataForTheGivenDeviceIdUuid")
	}

	result := response.Result().(*ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUID)
	return result, response, err

}

//GetPlannedAccessPointsForBuilding Get Planned Access Points for Building - b699-9b85-4e3b-acdd
/* Provides a list of Planned Access Points for the Building it is requested for


@param buildingID buildingId path parameter. The instance UUID of the building hierarchy element

@param GetPlannedAccessPointsForBuildingQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-planned-access-points-for-building
*/
func (s *DevicesService) GetPlannedAccessPointsForBuilding(buildingID string, GetPlannedAccessPointsForBuildingQueryParams *GetPlannedAccessPointsForBuildingQueryParams) (*ResponseDevicesGetPlannedAccessPointsForBuilding, *resty.Response, error) {
	path := "/dna/intent/api/v1/buildings/{buildingId}/planned-access-points"
	path = strings.Replace(path, "{buildingId}", fmt.Sprintf("%v", buildingID), -1)

	queryString, _ := query.Values(GetPlannedAccessPointsForBuildingQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetPlannedAccessPointsForBuilding{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPlannedAccessPointsForBuilding(buildingID, GetPlannedAccessPointsForBuildingQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetPlannedAccessPointsForBuilding")
	}

	result := response.Result().(*ResponseDevicesGetPlannedAccessPointsForBuilding)
	return result, response, err

}

//GetDeviceDetail Get Device Detail - ca98-fac4-4b08-895c
/* Returns detailed Network Device information retrieved by Mac Address, Device Name or UUID for any given point of time.


@param GetDeviceDetailQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-detail
*/
func (s *DevicesService) GetDeviceDetail(GetDeviceDetailQueryParams *GetDeviceDetailQueryParams) (*ResponseDevicesGetDeviceDetail, *resty.Response, error) {
	path := "/dna/intent/api/v1/device-detail"

	queryString, _ := query.Values(GetDeviceDetailQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetDeviceDetail{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceDetail(GetDeviceDetailQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceDetail")
	}

	result := response.Result().(*ResponseDevicesGetDeviceDetail)
	return result, response, err

}

//GetDeviceEnrichmentDetails Get Device Enrichment Details - e0b5-599b-4f29-97b7
/* Enriches a given network device context (device id or device Mac Address or device management IP address) with details about the device and neighbor topology


@param GetDeviceEnrichmentDetailsHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-enrichment-details
*/
func (s *DevicesService) GetDeviceEnrichmentDetails(GetDeviceEnrichmentDetailsHeaderParams *GetDeviceEnrichmentDetailsHeaderParams) (*ResponseDevicesGetDeviceEnrichmentDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/device-enrichment-details"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetDeviceEnrichmentDetailsHeaderParams != nil {

		if GetDeviceEnrichmentDetailsHeaderParams.EntityType != "" {
			clientRequest = clientRequest.SetHeader("entity_type", GetDeviceEnrichmentDetailsHeaderParams.EntityType)
		}

		if GetDeviceEnrichmentDetailsHeaderParams.EntityValue != "" {
			clientRequest = clientRequest.SetHeader("entity_value", GetDeviceEnrichmentDetailsHeaderParams.EntityValue)
		}

		if GetDeviceEnrichmentDetailsHeaderParams.Persistbapioutput != "" {
			clientRequest = clientRequest.SetHeader("__persistbapioutput", GetDeviceEnrichmentDetailsHeaderParams.Persistbapioutput)
		}

	}

	response, err = clientRequest.
		SetResult(&ResponseDevicesGetDeviceEnrichmentDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceEnrichmentDetails(GetDeviceEnrichmentDetailsHeaderParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceEnrichmentDetails")
	}

	result := response.Result().(*ResponseDevicesGetDeviceEnrichmentDetails)
	return result, response, err

}

//Devices Devices - 3ab2-bb64-4cca-81ee
/* Intent API for accessing Catalyst Center Assurance Device object for generating reports, creating dashboards or creating additional value added services.


@param DevicesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!devices
*/
func (s *DevicesService) Devices(DevicesQueryParams *DevicesQueryParams) (*ResponseDevicesDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/device-health"

	queryString, _ := query.Values(DevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.Devices(DevicesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation Devices")
	}

	result := response.Result().(*ResponseDevicesDevices)
	return result, response, err

}

//GetPlannedAccessPointsForFloor Get Planned Access Points for Floor - 6780-6977-4589-9a54
/* Provides a list of Planned Access Points for the Floor it is requested for


@param floorID floorId path parameter. The instance UUID of the floor hierarchy element

@param GetPlannedAccessPointsForFloorQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-planned-access-points-for-floor
*/
func (s *DevicesService) GetPlannedAccessPointsForFloor(floorID string, GetPlannedAccessPointsForFloorQueryParams *GetPlannedAccessPointsForFloorQueryParams) (*ResponseDevicesGetPlannedAccessPointsForFloor, *resty.Response, error) {
	path := "/dna/intent/api/v1/floors/{floorId}/planned-access-points"
	path = strings.Replace(path, "{floorId}", fmt.Sprintf("%v", floorID), -1)

	queryString, _ := query.Values(GetPlannedAccessPointsForFloorQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetPlannedAccessPointsForFloor{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPlannedAccessPointsForFloor(floorID, GetPlannedAccessPointsForFloorQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetPlannedAccessPointsForFloor")
	}

	result := response.Result().(*ResponseDevicesGetPlannedAccessPointsForFloor)
	return result, response, err

}

//GetAllHealthScoreDefinitionsForGivenFilters Get all health score definitions for given filters. - 9bb6-ea87-4ffb-b492
/* Get all health score defintions.
Supported filters are id, name and overall health include status. A health score definition can be different across device type. So, deviceType in the query param is important and default is all device types.
By default all supported attributes are listed in response. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-issueAndHealthDefinitions-1.0.0-resolved.yaml


@param GetAllHealthScoreDefinitionsForGivenFiltersHeaderParams Custom header parameters
@param GetAllHealthScoreDefinitionsForGivenFiltersQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-all-health-score-definitions-for-given-filters
*/
func (s *DevicesService) GetAllHealthScoreDefinitionsForGivenFilters(GetAllHealthScoreDefinitionsForGivenFiltersHeaderParams *GetAllHealthScoreDefinitionsForGivenFiltersHeaderParams, GetAllHealthScoreDefinitionsForGivenFiltersQueryParams *GetAllHealthScoreDefinitionsForGivenFiltersQueryParams) (*ResponseDevicesGetAllHealthScoreDefinitionsForGivenFilters, *resty.Response, error) {
	path := "/dna/intent/api/v1/healthScoreDefinitions"

	queryString, _ := query.Values(GetAllHealthScoreDefinitionsForGivenFiltersQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetAllHealthScoreDefinitionsForGivenFiltersHeaderParams != nil {

		if GetAllHealthScoreDefinitionsForGivenFiltersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetAllHealthScoreDefinitionsForGivenFiltersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetAllHealthScoreDefinitionsForGivenFilters{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAllHealthScoreDefinitionsForGivenFilters(GetAllHealthScoreDefinitionsForGivenFiltersHeaderParams, GetAllHealthScoreDefinitionsForGivenFiltersQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAllHealthScoreDefinitionsForGivenFilters")
	}

	result := response.Result().(*ResponseDevicesGetAllHealthScoreDefinitionsForGivenFilters)
	return result, response, err

}

//GetHealthScoreDefinitionForTheGivenID Get health score definition for the given id. - 99b5-d81a-4408-94c3
/* Get health score defintion for the given id. Definition includes all properties from HealthScoreDefinition schema by default. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-issueAndHealthDefinitions-1.0.0-resolved.yaml


@param id id path parameter. Health score definition id.

@param GetHealthScoreDefinitionForTheGivenIdHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-health-score-definition-for-the-given-id
*/
func (s *DevicesService) GetHealthScoreDefinitionForTheGivenID(id string, GetHealthScoreDefinitionForTheGivenIdHeaderParams *GetHealthScoreDefinitionForTheGivenIDHeaderParams) (*ResponseDevicesGetHealthScoreDefinitionForTheGivenID, *resty.Response, error) {
	path := "/dna/intent/api/v1/healthScoreDefinitions/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetHealthScoreDefinitionForTheGivenIdHeaderParams != nil {

		if GetHealthScoreDefinitionForTheGivenIdHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetHealthScoreDefinitionForTheGivenIdHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetResult(&ResponseDevicesGetHealthScoreDefinitionForTheGivenID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetHealthScoreDefinitionForTheGivenID(id, GetHealthScoreDefinitionForTheGivenIdHeaderParams)
		}
		return nil, response, fmt.Errorf("error with operation GetHealthScoreDefinitionForTheGivenId")
	}

	result := response.Result().(*ResponseDevicesGetHealthScoreDefinitionForTheGivenID)
	return result, response, err

}

//GetAllInterfaces Get all interfaces - f594-7a4c-439a-8bf0
/* Returns all available interfaces. This endpoint can return a maximum of 500 interfaces


@param GetAllInterfacesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-all-interfaces
*/
func (s *DevicesService) GetAllInterfaces(GetAllInterfacesQueryParams *GetAllInterfacesQueryParams) (*ResponseDevicesGetAllInterfaces, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface"

	queryString, _ := query.Values(GetAllInterfacesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetAllInterfaces{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAllInterfaces(GetAllInterfacesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAllInterfaces")
	}

	result := response.Result().(*ResponseDevicesGetAllInterfaces)
	return result, response, err

}

//GetDeviceInterfaceCountForMultipleDevices Get Device Interface Count for Multiple Devices - 3d92-3b18-4dc9-a4ca
/* Returns the count of interfaces for all devices



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-interface-count-for-multiple-devices
*/
func (s *DevicesService) GetDeviceInterfaceCountForMultipleDevices() (*ResponseDevicesGetDeviceInterfaceCountForMultipleDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceInterfaceCountForMultipleDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceInterfaceCountForMultipleDevices()
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceInterfaceCountForMultipleDevices")
	}

	result := response.Result().(*ResponseDevicesGetDeviceInterfaceCountForMultipleDevices)
	return result, response, err

}

//GetInterfaceByIP Get Interface by IP - cd84-69e6-47ca-ab0e
/* Returns list of interfaces for specified device management IP address


@param ipAddress ipAddress path parameter. IP address of the interface


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-interface-by-ip
*/
func (s *DevicesService) GetInterfaceByIP(ipAddress string) (*ResponseDevicesGetInterfaceByIP, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/ip-address/{ipAddress}"
	path = strings.Replace(path, "{ipAddress}", fmt.Sprintf("%v", ipAddress), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetInterfaceByIP{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetInterfaceByIP(ipAddress)
		}
		return nil, response, fmt.Errorf("error with operation GetInterfaceByIp")
	}

	result := response.Result().(*ResponseDevicesGetInterfaceByIP)
	return result, response, err

}

//GetIsisInterfaces Get ISIS interfaces - 84ad-8b0e-42ca-b48a
/* Returns the interfaces that has ISIS enabled



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-isis-interfaces
*/
func (s *DevicesService) GetIsisInterfaces() (*ResponseDevicesGetIsisInterfaces, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/isis"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetIsisInterfaces{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetIsisInterfaces()
		}
		return nil, response, fmt.Errorf("error with operation GetIsisInterfaces")
	}

	result := response.Result().(*ResponseDevicesGetIsisInterfaces)
	return result, response, err

}

//GetInterfaceInfoByID Get Interface info by Id - ba9d-c85b-4b8a-9a17
/* Returns list of interfaces by specified device


@param deviceID deviceId path parameter. Device ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-interface-info-by-id
*/
func (s *DevicesService) GetInterfaceInfoByID(deviceID string) (*ResponseDevicesGetInterfaceInfoByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/network-device/{deviceId}"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetInterfaceInfoByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetInterfaceInfoByID(deviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetInterfaceInfoById")
	}

	result := response.Result().(*ResponseDevicesGetInterfaceInfoByID)
	return result, response, err

}

//GetDeviceInterfaceCount Get Device Interface count - 5b86-3922-4cd8-8ea7
/* Returns the interface count for the given device


@param deviceID deviceId path parameter. Device ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-interface-count
*/
func (s *DevicesService) GetDeviceInterfaceCount(deviceID string) (*ResponseDevicesGetDeviceInterfaceCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/network-device/{deviceId}/count"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceInterfaceCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceInterfaceCount(deviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceInterfaceCount")
	}

	result := response.Result().(*ResponseDevicesGetDeviceInterfaceCount)
	return result, response, err

}

//GetInterfaceDetailsByDeviceIDAndInterfaceName Get Interface details by device Id and interface name - 4eb5-6a61-4cc9-a2d2
/* Returns interface by specified device Id and interface name


@param deviceID deviceId path parameter. Device ID

@param GetInterfaceDetailsByDeviceIdAndInterfaceNameQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-interface-details-by-device-id-and-interface-name
*/
func (s *DevicesService) GetInterfaceDetailsByDeviceIDAndInterfaceName(deviceID string, GetInterfaceDetailsByDeviceIdAndInterfaceNameQueryParams *GetInterfaceDetailsByDeviceIDAndInterfaceNameQueryParams) (*ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceName, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/network-device/{deviceId}/interface-name"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	queryString, _ := query.Values(GetInterfaceDetailsByDeviceIdAndInterfaceNameQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetInterfaceDetailsByDeviceIDAndInterfaceName(deviceID, GetInterfaceDetailsByDeviceIdAndInterfaceNameQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetInterfaceDetailsByDeviceIdAndInterfaceName")
	}

	result := response.Result().(*ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceName)
	return result, response, err

}

//GetDeviceInterfacesBySpecifiedRange Get Device Interfaces by specified range - 349c-8884-43b8-9a58
/* Returns the list of interfaces for the device for the specified range


@param deviceID deviceId path parameter. Device ID

@param startIndex startIndex path parameter. Start index

@param recordsToReturn recordsToReturn path parameter. Number of records to return


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-interfaces-by-specified-range
*/
func (s *DevicesService) GetDeviceInterfacesBySpecifiedRange(deviceID string, startIndex int, recordsToReturn int) (*ResponseDevicesGetDeviceInterfacesBySpecifiedRange, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/network-device/{deviceId}/{startIndex}/{recordsToReturn}"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)
	path = strings.Replace(path, "{startIndex}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{recordsToReturn}", fmt.Sprintf("%v", recordsToReturn), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceInterfacesBySpecifiedRange{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceInterfacesBySpecifiedRange(deviceID, startIndex, recordsToReturn)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceInterfacesBySpecifiedRange")
	}

	result := response.Result().(*ResponseDevicesGetDeviceInterfacesBySpecifiedRange)
	return result, response, err

}

//GetOspfInterfaces Get OSPF interfaces - 70ad-3976-49e9-b4d3
/* Returns the interfaces that has OSPF enabled



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-ospf-interfaces
*/
func (s *DevicesService) GetOspfInterfaces() (*ResponseDevicesGetOspfInterfaces, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/ospf"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetOspfInterfaces{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetOspfInterfaces()
		}
		return nil, response, fmt.Errorf("error with operation GetOspfInterfaces")
	}

	result := response.Result().(*ResponseDevicesGetOspfInterfaces)
	return result, response, err

}

//GetInterfaceByID Get Interface by Id - b888-792d-43ba-ba46
/* Returns the interface for the given interface ID


@param id id path parameter. Interface ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-interface-by-id
*/
func (s *DevicesService) GetInterfaceByID(id string) (*ResponseDevicesGetInterfaceByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetInterfaceByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetInterfaceByID(id)
		}
		return nil, response, fmt.Errorf("error with operation GetInterfaceById")
	}

	result := response.Result().(*ResponseDevicesGetInterfaceByID)
	return result, response, err

}

//LegitOperationsForInterface Legit operations for interface - 87a3-3a52-46ea-a40e
/* Get list of all properties & operations valid for an interface.


@param interfaceUUID interfaceUuid path parameter. Interface ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!legit-operations-for-interface
*/
func (s *DevicesService) LegitOperationsForInterface(interfaceUUID string) (*ResponseDevicesLegitOperationsForInterface, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/{interfaceUuid}/legit-operation"
	path = strings.Replace(path, "{interfaceUuid}", fmt.Sprintf("%v", interfaceUUID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesLegitOperationsForInterface{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.LegitOperationsForInterface(interfaceUUID)
		}
		return nil, response, fmt.Errorf("error with operation LegitOperationsForInterface")
	}

	result := response.Result().(*ResponseDevicesLegitOperationsForInterface)
	return result, response, err

}

//GetDeviceList Get Device list - 20b1-9b52-464b-8972
/* Returns list of network devices based on filter criteria such as management IP address, mac address, hostname, etc. You can use the .* in any value to conduct a wildcard search. For example, to find all hostnames beginning with myhost in the IP address range 192.25.18.n, issue the following request: GET /dna/intent/api/v1/network-device?hostname=myhost.*&managementIpAddress=192.25.18..*
If id parameter is provided with comma separated ids, it will return the list of network-devices for the given ids and ignores the other request parameters. You can also specify offset & limit to get the required list.


@param GetDeviceListQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-list
*/
func (s *DevicesService) GetDeviceList(GetDeviceListQueryParams *GetDeviceListQueryParams) (*ResponseDevicesGetDeviceList, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device"

	queryString, _ := query.Values(GetDeviceListQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetDeviceList{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceList(GetDeviceListQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceList")
	}

	result := response.Result().(*ResponseDevicesGetDeviceList)
	return result, response, err

}

//GetDeviceValuesThatMatchFullyOrPartiallyAnAttribute Get Device Values that match fully or partially an Attribute - ffa7-48cc-44e9-a437
/* Returns the list of values of the first given required parameter. You can use the .* in any value to conduct a wildcard search. For example, to get all the devices with the management IP address starting with 10.10. , issue the following request: GET /dna/inten/api/v1/network-device/autocomplete?managementIpAddress=10.10..* It will return the device management IP addresses that match fully or partially the provided attribute. {[10.10.1.1, 10.10.20.2, …]}.


@param GetDeviceValuesThatMatchFullyOrPartiallyAnAttributeQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-values-that-match-fully-or-partially-an-attribute
*/
func (s *DevicesService) GetDeviceValuesThatMatchFullyOrPartiallyAnAttribute(GetDeviceValuesThatMatchFullyOrPartiallyAnAttributeQueryParams *GetDeviceValuesThatMatchFullyOrPartiallyAnAttributeQueryParams) (*resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/autocomplete"

	queryString, _ := query.Values(GetDeviceValuesThatMatchFullyOrPartiallyAnAttributeQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceValuesThatMatchFullyOrPartiallyAnAttribute(GetDeviceValuesThatMatchFullyOrPartiallyAnAttributeQueryParams)
		}
		return response, fmt.Errorf("error with operation GetDeviceValuesThatMatchFullyOrPartiallyAnAttribute")
	}

	return response, err

}

//GetPollingIntervalForAllDevices Get Polling Interval for all devices - 38bd-0b88-4b89-a785
/* Returns polling interval of all devices



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-polling-interval-for-all-devices
*/
func (s *DevicesService) GetPollingIntervalForAllDevices() (*ResponseDevicesGetPollingIntervalForAllDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/collection-schedule/global"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetPollingIntervalForAllDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPollingIntervalForAllDevices()
		}
		return nil, response, fmt.Errorf("error with operation GetPollingIntervalForAllDevices")
	}

	result := response.Result().(*ResponseDevicesGetPollingIntervalForAllDevices)
	return result, response, err

}

//GetDeviceConfigForAllDevices Get Device Config for all devices - b7bc-aa08-4e2b-90d0
/* Returns the config for all devices. This API has been deprecated and will not be available in a Cisco Catalyst Center release after Nov 1st 2024 23:59:59 GMT.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-config-for-all-devices
*/
func (s *DevicesService) GetDeviceConfigForAllDevices() (*ResponseDevicesGetDeviceConfigForAllDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/config"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceConfigForAllDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceConfigForAllDevices()
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceConfigForAllDevices")
	}

	result := response.Result().(*ResponseDevicesGetDeviceConfigForAllDevices)
	return result, response, err

}

//GetDeviceConfigCount Get Device Config Count - 888f-585c-49b8-8441
/* Returns the count of device configs



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-config-count
*/
func (s *DevicesService) GetDeviceConfigCount() (*ResponseDevicesGetDeviceConfigCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/config/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceConfigCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceConfigCount()
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceConfigCount")
	}

	result := response.Result().(*ResponseDevicesGetDeviceConfigCount)
	return result, response, err

}

//GetDeviceCount2 Get Device Count - 5db2-1b8e-43fa-b7d8
/* Returns the count of network devices based on the filter criteria by management IP address, mac address, hostname and location name


@param GetDeviceCount2QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-count2
*/
func (s *DevicesService) GetDeviceCount2(GetDeviceCount2QueryParams *GetDeviceCount2QueryParams) (*ResponseDevicesGetDeviceCount2, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/count"

	queryString, _ := query.Values(GetDeviceCount2QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetDeviceCount2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceCount2(GetDeviceCount2QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceCount2")
	}

	result := response.Result().(*ResponseDevicesGetDeviceCount2)
	return result, response, err

}

//GetFunctionalCapabilityForDevices Get Functional Capability for devices - c3b3-c9ef-4e6b-8a09
/* Returns the functional-capability for given devices


@param GetFunctionalCapabilityForDevicesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-functional-capability-for-devices
*/
func (s *DevicesService) GetFunctionalCapabilityForDevices(GetFunctionalCapabilityForDevicesQueryParams *GetFunctionalCapabilityForDevicesQueryParams) (*ResponseDevicesGetFunctionalCapabilityForDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/functional-capability"

	queryString, _ := query.Values(GetFunctionalCapabilityForDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetFunctionalCapabilityForDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFunctionalCapabilityForDevices(GetFunctionalCapabilityForDevicesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetFunctionalCapabilityForDevices")
	}

	result := response.Result().(*ResponseDevicesGetFunctionalCapabilityForDevices)
	return result, response, err

}

//GetFunctionalCapabilityByID Get Functional Capability by Id - 81bb-4804-405a-8d2f
/* Returns functional capability with given Id


@param id id path parameter. Functional Capability UUID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-functional-capability-by-id
*/
func (s *DevicesService) GetFunctionalCapabilityByID(id string) (*ResponseDevicesGetFunctionalCapabilityByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/functional-capability/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetFunctionalCapabilityByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFunctionalCapabilityByID(id)
		}
		return nil, response, fmt.Errorf("error with operation GetFunctionalCapabilityById")
	}

	result := response.Result().(*ResponseDevicesGetFunctionalCapabilityByID)
	return result, response, err

}

//InventoryInsightDeviceLinkMismatchApI Inventory Insight Device Link Mismatch API - 5792-59d8-4208-8190
/* Find all devices with link mismatch (speed /  vlan)


@param siteID siteId path parameter.
@param InventoryInsightDeviceLinkMismatchApIQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!inventory-insight-device-link-mismatch-api
*/
func (s *DevicesService) InventoryInsightDeviceLinkMismatchAPI(siteID string, InventoryInsightDeviceLinkMismatchAPIQueryParams *InventoryInsightDeviceLinkMismatchAPIQueryParams) (*ResponseDevicesInventoryInsightDeviceLinkMismatchAPI, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/insight/{siteId}/device-link"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	queryString, _ := query.Values(InventoryInsightDeviceLinkMismatchAPIQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesInventoryInsightDeviceLinkMismatchAPI{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.InventoryInsightDeviceLinkMismatchAPI(siteID, InventoryInsightDeviceLinkMismatchAPIQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation InventoryInsightDeviceLinkMismatchApi")
	}

	result := response.Result().(*ResponseDevicesInventoryInsightDeviceLinkMismatchAPI)
	return result, response, err

}

//GetNetworkDeviceByIP Get Network Device by IP - d0a4-b881-45aa-bb51
/* Returns the network device by specified IP address


@param ipAddress ipAddress path parameter. Device IP address


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-device-by-ip
*/
func (s *DevicesService) GetNetworkDeviceByIP(ipAddress string) (*ResponseDevicesGetNetworkDeviceByIP, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/ip-address/{ipAddress}"
	path = strings.Replace(path, "{ipAddress}", fmt.Sprintf("%v", ipAddress), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetNetworkDeviceByIP{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetNetworkDeviceByIP(ipAddress)
		}
		return nil, response, fmt.Errorf("error with operation GetNetworkDeviceByIp")
	}

	result := response.Result().(*ResponseDevicesGetNetworkDeviceByIP)
	return result, response, err

}

//GetModules Get Modules - eb82-49e3-4f69-b0f1
/* Returns modules by specified device id


@param GetModulesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-modules
*/
func (s *DevicesService) GetModules(GetModulesQueryParams *GetModulesQueryParams) (*ResponseDevicesGetModules, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/module"

	queryString, _ := query.Values(GetModulesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetModules{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetModules(GetModulesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetModules")
	}

	result := response.Result().(*ResponseDevicesGetModules)
	return result, response, err

}

//GetModuleCount Get Module count - 8db9-3974-4649-a782
/* Returns Module Count


@param GetModuleCountQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-module-count
*/
func (s *DevicesService) GetModuleCount(GetModuleCountQueryParams *GetModuleCountQueryParams) (*ResponseDevicesGetModuleCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/module/count"

	queryString, _ := query.Values(GetModuleCountQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetModuleCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetModuleCount(GetModuleCountQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetModuleCount")
	}

	result := response.Result().(*ResponseDevicesGetModuleCount)
	return result, response, err

}

//GetModuleInfoByID Get Module Info by Id - 0db7-da74-4c0b-83d8
/* Returns Module info by 'module id'


@param id id path parameter. Module id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-module-info-by-id
*/
func (s *DevicesService) GetModuleInfoByID(id string) (*ResponseDevicesGetModuleInfoByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/module/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetModuleInfoByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetModuleInfoByID(id)
		}
		return nil, response, fmt.Errorf("error with operation GetModuleInfoById")
	}

	result := response.Result().(*ResponseDevicesGetModuleInfoByID)
	return result, response, err

}

//GetDeviceBySerialNumber Get Device by Serial number - d888-ab6d-4d59-a8c1
/* Returns the network device with given serial number


@param serialNumber serialNumber path parameter. Device serial number


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-by-serial-number
*/
func (s *DevicesService) GetDeviceBySerialNumber(serialNumber string) (*ResponseDevicesGetDeviceBySerialNumber, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/serial-number/{serialNumber}"
	path = strings.Replace(path, "{serialNumber}", fmt.Sprintf("%v", serialNumber), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceBySerialNumber{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceBySerialNumber(serialNumber)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceBySerialNumber")
	}

	result := response.Result().(*ResponseDevicesGetDeviceBySerialNumber)
	return result, response, err

}

//GetDevicesRegisteredForWsaNotification Get Devices registered for WSA Notification - c980-9b67-44f8-a502
/* It fetches devices which are registered to receive WSA notifications. The device serial number and/or MAC address are required to be provided as query parameters.


@param GetDevicesRegisteredForWSANotificationQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-devices-registered-for-wsa-notification
*/
func (s *DevicesService) GetDevicesRegisteredForWsaNotification(GetDevicesRegisteredForWSANotificationQueryParams *GetDevicesRegisteredForWsaNotificationQueryParams) (*ResponseDevicesGetDevicesRegisteredForWsaNotification, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/tenantinfo/macaddress"

	queryString, _ := query.Values(GetDevicesRegisteredForWSANotificationQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetDevicesRegisteredForWsaNotification{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDevicesRegisteredForWsaNotification(GetDevicesRegisteredForWSANotificationQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDevicesRegisteredForWsaNotification")
	}

	result := response.Result().(*ResponseDevicesGetDevicesRegisteredForWsaNotification)
	return result, response, err

}

//GetAllUserDefinedFields Get All User-Defined-Fields - 058d-2a92-4899-b7bb
/* Gets existing global User Defined Fields. If no input is given, it fetches ALL the Global UDFs. Filter/search is supported by UDF Id(s) or UDF name(s) or both.


@param GetAllUserDefinedFieldsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-all-user-defined-fields
*/
func (s *DevicesService) GetAllUserDefinedFields(GetAllUserDefinedFieldsQueryParams *GetAllUserDefinedFieldsQueryParams) (*ResponseDevicesGetAllUserDefinedFields, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/user-defined-field"

	queryString, _ := query.Values(GetAllUserDefinedFieldsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetAllUserDefinedFields{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAllUserDefinedFields(GetAllUserDefinedFieldsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAllUserDefinedFields")
	}

	result := response.Result().(*ResponseDevicesGetAllUserDefinedFields)
	return result, response, err

}

//GetChassisDetailsForDevice Get Chassis Details for Device - 0486-9b26-49ab-b579
/* Returns chassis details for given device ID


@param deviceID deviceId path parameter. Device ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-chassis-details-for-device
*/
func (s *DevicesService) GetChassisDetailsForDevice(deviceID string) (*ResponseDevicesGetChassisDetailsForDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceId}/chassis"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetChassisDetailsForDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetChassisDetailsForDevice(deviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetChassisDetailsForDevice")
	}

	result := response.Result().(*ResponseDevicesGetChassisDetailsForDevice)
	return result, response, err

}

//GetStackDetailsForDevice Get Stack Details for Device - 78a7-7ab0-4d5a-8a10
/* Retrieves complete stack details for given device ID


@param deviceID deviceId path parameter. Device ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-stack-details-for-device
*/
func (s *DevicesService) GetStackDetailsForDevice(deviceID string) (*ResponseDevicesGetStackDetailsForDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceId}/stack"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetStackDetailsForDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetStackDetailsForDevice(deviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetStackDetailsForDevice")
	}

	result := response.Result().(*ResponseDevicesGetStackDetailsForDevice)
	return result, response, err

}

//GetTheDetailsOfPhysicalComponentsOfTheGivenDevice Get the Details of Physical Components of the Given Device. - 20b1-9b52-464b-897a
/* Return all types of equipment details like PowerSupply, Fan, Chassis, Backplane, Module, PROCESSOR, Other and SFP for the Given device.


@param deviceUUID deviceUuid path parameter.
@param GetTheDetailsOfPhysicalComponentsOfTheGivenDeviceQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-the-details-of-physical-components-of-the-given-device
*/
func (s *DevicesService) GetTheDetailsOfPhysicalComponentsOfTheGivenDevice(deviceUUID string, GetTheDetailsOfPhysicalComponentsOfTheGivenDeviceQueryParams *GetTheDetailsOfPhysicalComponentsOfTheGivenDeviceQueryParams) (*ResponseDevicesGetTheDetailsOfPhysicalComponentsOfTheGivenDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceUuid}/equipment"
	path = strings.Replace(path, "{deviceUuid}", fmt.Sprintf("%v", deviceUUID), -1)

	queryString, _ := query.Values(GetTheDetailsOfPhysicalComponentsOfTheGivenDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetTheDetailsOfPhysicalComponentsOfTheGivenDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTheDetailsOfPhysicalComponentsOfTheGivenDevice(deviceUUID, GetTheDetailsOfPhysicalComponentsOfTheGivenDeviceQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTheDetailsOfPhysicalComponentsOfTheGivenDevice")
	}

	result := response.Result().(*ResponseDevicesGetTheDetailsOfPhysicalComponentsOfTheGivenDevice)
	return result, response, err

}

//ReturnsPoeInterfaceDetailsForTheDevice Returns POE interface details for the device. - 20b5-48af-42da-a337
/* Returns POE interface details for the device, where deviceuuid is mandatory & accepts comma seperated interface names which is optional and returns information for that particular interfaces where(operStatus = operationalStatus)


@param deviceUUID deviceUuid path parameter. uuid of the device

@param ReturnsPOEInterfaceDetailsForTheDeviceQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!returns-poe-interface-details-for-the-device
*/
func (s *DevicesService) ReturnsPoeInterfaceDetailsForTheDevice(deviceUUID string, ReturnsPOEInterfaceDetailsForTheDeviceQueryParams *ReturnsPoeInterfaceDetailsForTheDeviceQueryParams) (*ResponseDevicesReturnsPoeInterfaceDetailsForTheDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceUuid}/interface/poe-detail"
	path = strings.Replace(path, "{deviceUuid}", fmt.Sprintf("%v", deviceUUID), -1)

	queryString, _ := query.Values(ReturnsPOEInterfaceDetailsForTheDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesReturnsPoeInterfaceDetailsForTheDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReturnsPoeInterfaceDetailsForTheDevice(deviceUUID, ReturnsPOEInterfaceDetailsForTheDeviceQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReturnsPoeInterfaceDetailsForTheDevice")
	}

	result := response.Result().(*ResponseDevicesReturnsPoeInterfaceDetailsForTheDevice)
	return result, response, err

}

//GetConnectedDeviceDetail Get connected device detail - a8aa-ca21-4c09-8388
/* Get connected device detail for given deviceUuid and interfaceUuid


@param deviceUUID deviceUuid path parameter. instanceuuid of Device

@param interfaceUUID interfaceUuid path parameter. instanceuuid of interface


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-connected-device-detail
*/
func (s *DevicesService) GetConnectedDeviceDetail(deviceUUID string, interfaceUUID string) (*ResponseDevicesGetConnectedDeviceDetail, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceUuid}/interface/{interfaceUuid}/neighbor"
	path = strings.Replace(path, "{deviceUuid}", fmt.Sprintf("%v", deviceUUID), -1)
	path = strings.Replace(path, "{interfaceUuid}", fmt.Sprintf("%v", interfaceUUID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetConnectedDeviceDetail{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetConnectedDeviceDetail(deviceUUID, interfaceUUID)
		}
		return nil, response, fmt.Errorf("error with operation GetConnectedDeviceDetail")
	}

	result := response.Result().(*ResponseDevicesGetConnectedDeviceDetail)
	return result, response, err

}

//GetLinecardDetails Get Linecard details - 46a1-4b02-48fb-8fbf
/* Get line card detail for a given deviceuuid.  Response will contain serial no, part no, switch no and slot no.


@param deviceUUID deviceUuid path parameter. instanceuuid of device


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-linecard-details
*/
func (s *DevicesService) GetLinecardDetails(deviceUUID string) (*ResponseDevicesGetLinecardDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceUuid}/line-card"
	path = strings.Replace(path, "{deviceUuid}", fmt.Sprintf("%v", deviceUUID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetLinecardDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetLinecardDetails(deviceUUID)
		}
		return nil, response, fmt.Errorf("error with operation GetLinecardDetails")
	}

	result := response.Result().(*ResponseDevicesGetLinecardDetails)
	return result, response, err

}

//PoeDetails POE details  - 8ba6-7932-4ed9-abae
/* Returns POE details for device.


@param deviceUUID deviceUuid path parameter. UUID of the device


Documentation Link: https://developer.cisco.com/docs/dna-center/#!poe-details
*/
func (s *DevicesService) PoeDetails(deviceUUID string) (*ResponseDevicesPoeDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceUuid}/poe"
	path = strings.Replace(path, "{deviceUuid}", fmt.Sprintf("%v", deviceUUID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesPoeDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.PoeDetails(deviceUUID)
		}
		return nil, response, fmt.Errorf("error with operation PoeDetails")
	}

	result := response.Result().(*ResponseDevicesPoeDetails)
	return result, response, err

}

//GetSupervisorCardDetail Get Supervisor card detail - 88aa-1b52-4a38-bf97
/* Get supervisor card detail for a given deviceuuid. Response will contain serial no, part no, switch no and slot no.


@param deviceUUID deviceUuid path parameter. instanceuuid of device


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-supervisor-card-detail
*/
func (s *DevicesService) GetSupervisorCardDetail(deviceUUID string) (*ResponseDevicesGetSupervisorCardDetail, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceUuid}/supervisor-card"
	path = strings.Replace(path, "{deviceUuid}", fmt.Sprintf("%v", deviceUUID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetSupervisorCardDetail{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSupervisorCardDetail(deviceUUID)
		}
		return nil, response, fmt.Errorf("error with operation GetSupervisorCardDetail")
	}

	result := response.Result().(*ResponseDevicesGetSupervisorCardDetail)
	return result, response, err

}

//GetDeviceByID Get Device by ID - 8fa8-eb40-4a4a-8d96
/* Returns the network device details for the given device ID


@param id id path parameter. Device ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-by-id
*/
func (s *DevicesService) GetDeviceByID(id string) (*ResponseDevicesGetDeviceByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceByID(id)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceById")
	}

	result := response.Result().(*ResponseDevicesGetDeviceByID)
	return result, response, err

}

//GetDeviceSummary Get Device Summary - 819f-9aa5-4fea-b7bf
/* Returns brief summary of device info such as hostname, management IP address for the given device Id


@param id id path parameter. Device ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-summary
*/
func (s *DevicesService) GetDeviceSummary(id string) (*ResponseDevicesGetDeviceSummary, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{id}/brief"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceSummary{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceSummary(id)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceSummary")
	}

	result := response.Result().(*ResponseDevicesGetDeviceSummary)
	return result, response, err

}

//GetPollingIntervalByID Get Polling Interval by Id - 8291-8a1b-4d28-9c5c
/* Returns polling interval by device id


@param id id path parameter. Device ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-polling-interval-by-id
*/
func (s *DevicesService) GetPollingIntervalByID(id string) (*ResponseDevicesGetPollingIntervalByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{id}/collection-schedule"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetPollingIntervalByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPollingIntervalByID(id)
		}
		return nil, response, fmt.Errorf("error with operation GetPollingIntervalById")
	}

	result := response.Result().(*ResponseDevicesGetPollingIntervalByID)
	return result, response, err

}

//GetOrganizationListForMeraki Get Organization list for Meraki - 84b3-7ae5-4c59-ab28
/* Returns list of organizations for meraki dashboard


@param id id path parameter. Device Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-list-for-meraki
*/
func (s *DevicesService) GetOrganizationListForMeraki(id string) (*ResponseDevicesGetOrganizationListForMeraki, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{id}/meraki-organization"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetOrganizationListForMeraki{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetOrganizationListForMeraki(id)
		}
		return nil, response, fmt.Errorf("error with operation GetOrganizationListForMeraki")
	}

	result := response.Result().(*ResponseDevicesGetOrganizationListForMeraki)
	return result, response, err

}

//GetDeviceInterfaceVLANs Get Device Interface VLANs - 288d-f949-4f2a-9746
/* Returns Device Interface VLANs. If parameter value is null or empty, it won't return any value in response.


@param id id path parameter.
@param GetDeviceInterfaceVLANsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-interface-vlans
*/
func (s *DevicesService) GetDeviceInterfaceVLANs(id string, GetDeviceInterfaceVLANsQueryParams *GetDeviceInterfaceVLANsQueryParams) (*ResponseDevicesGetDeviceInterfaceVLANs, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{id}/vlan"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetDeviceInterfaceVLANsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetDeviceInterfaceVLANs{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceInterfaceVLANs(id, GetDeviceInterfaceVLANsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceInterfaceVlans")
	}

	result := response.Result().(*ResponseDevicesGetDeviceInterfaceVLANs)
	return result, response, err

}

//GetWirelessLanControllerDetailsByID Get wireless lan controller details by Id - f682-6a8e-41bb-a242
/* Returns the wireless lan controller info with given device ID


@param id id path parameter. Device ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-wireless-lan-controller-details-by-id
*/
func (s *DevicesService) GetWirelessLanControllerDetailsByID(id string) (*ResponseDevicesGetWirelessLanControllerDetailsByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{id}/wireless-info"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetWirelessLanControllerDetailsByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetWirelessLanControllerDetailsByID(id)
		}
		return nil, response, fmt.Errorf("error with operation GetWirelessLanControllerDetailsById")
	}

	result := response.Result().(*ResponseDevicesGetWirelessLanControllerDetailsByID)
	return result, response, err

}

//GetDeviceConfigByID Get Device Config by Id - 84b3-3a9e-480a-bcaf
/* Returns the device config by specified device ID


@param networkDeviceID networkDeviceId path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-config-by-id
*/
func (s *DevicesService) GetDeviceConfigByID(networkDeviceID string) (*ResponseDevicesGetDeviceConfigByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{networkDeviceId}/config"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceConfigByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceConfigByID(networkDeviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceConfigById")
	}

	result := response.Result().(*ResponseDevicesGetDeviceConfigByID)
	return result, response, err

}

//GetNetworkDeviceByPaginationRange Get Network Device by pagination range - f495-48c5-4be8-a3e2
/* Returns the list of network devices for the given pagination range. The maximum number of records that can be retrieved is 500


@param startIndex startIndex path parameter. Start index [>=1]

@param recordsToReturn recordsToReturn path parameter. Number of records to return [1<= recordsToReturn <= 500]


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-device-by-pagination-range
*/
func (s *DevicesService) GetNetworkDeviceByPaginationRange(startIndex int, recordsToReturn int) (*ResponseDevicesGetNetworkDeviceByPaginationRange, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{startIndex}/{recordsToReturn}"
	path = strings.Replace(path, "{startIndex}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{recordsToReturn}", fmt.Sprintf("%v", recordsToReturn), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetNetworkDeviceByPaginationRange{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetNetworkDeviceByPaginationRange(startIndex, recordsToReturn)
		}
		return nil, response, fmt.Errorf("error with operation GetNetworkDeviceByPaginationRange")
	}

	result := response.Result().(*ResponseDevicesGetNetworkDeviceByPaginationRange)
	return result, response, err

}

//GetResyncIntervalForTheNetworkDevice Get resync interval for the network device - 4783-7a87-4aea-91e6
/* Fetch the reysnc interval for the given network device id.


@param id id path parameter. The id of the network device.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-resync-interval-for-the-network-device
*/
func (s *DevicesService) GetResyncIntervalForTheNetworkDevice(id string) (*ResponseDevicesGetResyncIntervalForTheNetworkDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/{id}/resyncIntervalSettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetResyncIntervalForTheNetworkDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetResyncIntervalForTheNetworkDevice(id)
		}
		return nil, response, fmt.Errorf("error with operation GetResyncIntervalForTheNetworkDevice")
	}

	result := response.Result().(*ResponseDevicesGetResyncIntervalForTheNetworkDevice)
	return result, response, err

}

//GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFilters Get the count of health score definitions based on provided filters. - 49aa-bb2c-46ca-b58a
/* Get the count of health score definitions based on provided filters. Supported filters are id, name and overall health include status. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-issueAndHealthDefinitions-1.0.0-resolved.yaml


@param GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersHeaderParams Custom header parameters
@param GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-the-count-of-health-score-definitions-based-on-provided-filters
*/
func (s *DevicesService) GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFilters(GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersHeaderParams *GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersHeaderParams, GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersQueryParams *GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersQueryParams) (*ResponseDevicesGetTheCountOfHealthScoreDefinitionsBasedOnProvidedFilters, *resty.Response, error) {
	path := "/intent/api/v1/healthScoreDefinitions/count"

	queryString, _ := query.Values(GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersHeaderParams != nil {

		if GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetTheCountOfHealthScoreDefinitionsBasedOnProvidedFilters{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFilters(GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersHeaderParams, GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFilters")
	}

	result := response.Result().(*ResponseDevicesGetTheCountOfHealthScoreDefinitionsBasedOnProvidedFilters)
	return result, response, err

}

//GetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions Gets the total number Network Devices based on the provided complex filters and aggregation functions. - 278f-1a5c-40ab-b65a
/* Gets the total number Network Devices based on the provided complex filters and aggregation functions. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceNetworkDevices-1.0.2-resolved.yaml



Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-the-total-number-network-devices-based-on-the-provided-complex-filters-and-aggregation-functions
*/
func (s *DevicesService) GetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions(requestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions *RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions) (*ResponseDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions, *resty.Response, error) {
	path := "/data/api/v1/networkDevices/query/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions).
		SetResult(&ResponseDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions(requestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions)
		}

		return nil, response, fmt.Errorf("error with operation GetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions")
	}

	result := response.Result().(*ResponseDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions)
	return result, response, err

}

//QueryAssuranceEventsWithFilters Query assurance events with filters - c5b7-0a69-4409-9b5b
/* Returns the list of events discovered by Catalyst Center, determined by the complex filters. Please refer to the 'API Support Documentation' section to understand which fields are supported. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceEvents-1.0.0-resolved.yaml


@param QueryAssuranceEventsWithFiltersHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!query-assurance-events-with-filters
*/
func (s *DevicesService) QueryAssuranceEventsWithFilters(requestDevicesQueryAssuranceEventsWithFilters *RequestDevicesQueryAssuranceEventsWithFilters, QueryAssuranceEventsWithFiltersHeaderParams *QueryAssuranceEventsWithFiltersHeaderParams) (*ResponseDevicesQueryAssuranceEventsWithFilters, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceEvents/query"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if QueryAssuranceEventsWithFiltersHeaderParams != nil {

		if QueryAssuranceEventsWithFiltersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", QueryAssuranceEventsWithFiltersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesQueryAssuranceEventsWithFilters).
		SetResult(&ResponseDevicesQueryAssuranceEventsWithFilters{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.QueryAssuranceEventsWithFilters(requestDevicesQueryAssuranceEventsWithFilters, QueryAssuranceEventsWithFiltersHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation QueryAssuranceEventsWithFilters")
	}

	result := response.Result().(*ResponseDevicesQueryAssuranceEventsWithFilters)
	return result, response, err

}

//CountTheNumberOfEventsWithFilters Count the number of events with filters - d685-3aeb-4878-a8fd
/* API to fetch the count of assurance events for the given complex query. Please refer to the 'API Support Documentation' section to understand which fields are supported. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceEvents-1.0.0-resolved.yaml


@param CountTheNumberOfEventsWithFiltersHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!count-the-number-of-events-with-filters
*/
func (s *DevicesService) CountTheNumberOfEventsWithFilters(requestDevicesCountTheNumberOfEventsWithFilters *RequestDevicesCountTheNumberOfEventsWithFilters, CountTheNumberOfEventsWithFiltersHeaderParams *CountTheNumberOfEventsWithFiltersHeaderParams) (*ResponseDevicesCountTheNumberOfEventsWithFilters, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceEvents/query/count"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if CountTheNumberOfEventsWithFiltersHeaderParams != nil {

		if CountTheNumberOfEventsWithFiltersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", CountTheNumberOfEventsWithFiltersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesCountTheNumberOfEventsWithFilters).
		SetResult(&ResponseDevicesCountTheNumberOfEventsWithFilters{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CountTheNumberOfEventsWithFilters(requestDevicesCountTheNumberOfEventsWithFilters, CountTheNumberOfEventsWithFiltersHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation CountTheNumberOfEventsWithFilters")
	}

	result := response.Result().(*ResponseDevicesCountTheNumberOfEventsWithFilters)
	return result, response, err

}

//GetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions Gets the list of interfaces across the Network Devices based on the provided complex filters and aggregation functions - 45b8-ba96-4daa-843c
/* Gets the list of interfaces across the Network Devices based on the provided complex filters and aggregation functions
The elements are grouped and sorted by deviceUuid first, and are then sorted by the given sort field, or by the default value: name.
The supported sorting options are: name, adminStatus, description, duplexConfig, duplexOper, interfaceIfIndex,interfaceType, macAddress,mediaType, operStatus, portChannelId, portMode, portType,speed, vlanId. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-interfaces-1.0.2-resolved.yaml



Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-the-list-of-interfaces-across-the-network-devices-based-on-the-provided-complex-filters-and-aggregation-functions
*/
func (s *DevicesService) GetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions(requestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions *RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions) (*ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions, *resty.Response, error) {
	path := "/dna/data/api/v1/interfaces/query"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions).
		SetResult(&ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions(requestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions)
		}

		return nil, response, fmt.Errorf("error with operation GetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions")
	}

	result := response.Result().(*ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions)
	return result, response, err

}

//TheTotalInterfacesCountAcrossTheNetworkDevices The Total interfaces count across the Network devices. - a0bb-1bed-4529-98b1
/* Gets the total number of interfaces across the Network devices based on the provided complex filters and aggregation functions. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-interfaces-1.0.2-resolved.yaml



Documentation Link: https://developer.cisco.com/docs/dna-center/#!the-total-interfaces-count-across-the-network-devices
*/
func (s *DevicesService) TheTotalInterfacesCountAcrossTheNetworkDevices(requestDevicesTheTotalInterfacesCountAcrossTheNetworkDevices *RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevices) (*ResponseDevicesTheTotalInterfacesCountAcrossTheNetworkDevices, *resty.Response, error) {
	path := "/dna/data/api/v1/interfaces/query/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesTheTotalInterfacesCountAcrossTheNetworkDevices).
		SetResult(&ResponseDevicesTheTotalInterfacesCountAcrossTheNetworkDevices{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.TheTotalInterfacesCountAcrossTheNetworkDevices(requestDevicesTheTotalInterfacesCountAcrossTheNetworkDevices)
		}

		return nil, response, fmt.Errorf("error with operation TheTotalInterfacesCountAcrossTheNetworkDevices")
	}

	result := response.Result().(*ResponseDevicesTheTotalInterfacesCountAcrossTheNetworkDevices)
	return result, response, err

}

//GetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions Gets the list of Network Devices based on the provided complex filters and aggregation functions. - e794-1a90-428b-b583
/* Gets the list of Network Devices based on the provided complex filters and aggregation functions. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceNetworkDevices-1.0.2-resolved.yaml



Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-the-list-of-network-devices-based-on-the-provided-complex-filters-and-aggregation-functions
*/
func (s *DevicesService) GetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions(requestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions *RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions) (*ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions, *resty.Response, error) {
	path := "/dna/data/api/v1/networkDevices/query"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions).
		SetResult(&ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions(requestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions)
		}

		return nil, response, fmt.Errorf("error with operation GetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions")
	}

	result := response.Result().(*ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions)
	return result, response, err

}

//GetsTheSummaryAnalyticsDataRelatedToNetworkDevices Gets the summary analytics data related to network devices. - 15be-c9ed-4cba-8f91
/* Gets the summary analytics data related to network devices based on the provided input data. This endpoint helps to obtain the consolidated insights into the performance and status of the monitored network devices. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceNetworkDevices-1.0.2-resolved.yaml



Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-the-summary-analytics-data-related-to-network-devices
*/
func (s *DevicesService) GetsTheSummaryAnalyticsDataRelatedToNetworkDevices(requestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevices *RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevices) (*ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevices, *resty.Response, error) {
	path := "/dna/data/api/v1/networkDevices/summaryAnalytics"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevices).
		SetResult(&ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevices{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsTheSummaryAnalyticsDataRelatedToNetworkDevices(requestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevices)
		}

		return nil, response, fmt.Errorf("error with operation GetsTheSummaryAnalyticsDataRelatedToNetworkDevices")
	}

	result := response.Result().(*ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevices)
	return result, response, err

}

//GetsTheTrendAnalyticsData Gets the Trend analytics data. - 0c93-595e-451b-910e
/* Gets the Trend analytics Network device data for the given time range. The data will be grouped based on the given trend time Interval. The required property for this API is 'trendInterval'. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceNetworkDevices-1.0.2-resolved.yaml



Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-the-trend-analytics-data
*/
func (s *DevicesService) GetsTheTrendAnalyticsData(requestDevicesGetsTheTrendAnalyticsData *RequestDevicesGetsTheTrendAnalyticsData) (*ResponseDevicesGetsTheTrendAnalyticsData, *resty.Response, error) {
	path := "/dna/data/api/v1/networkDevices/trendAnalytics"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesGetsTheTrendAnalyticsData).
		SetResult(&ResponseDevicesGetsTheTrendAnalyticsData{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsTheTrendAnalyticsData(requestDevicesGetsTheTrendAnalyticsData)
		}

		return nil, response, fmt.Errorf("error with operation GetsTheTrendAnalyticsData")
	}

	result := response.Result().(*ResponseDevicesGetsTheTrendAnalyticsData)
	return result, response, err

}

//TheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRange The Trend analytics data for the network Device in the specified time range - 00ba-7a81-431b-93e7
/* The Trend analytics data for the network Device in the specified time range. The data is grouped based on the trend time Interval, other input parameters like attribute and aggregate attributes. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceNetworkDevices-1.0.2-resolved.yaml


@param id id path parameter. The device Uuid


Documentation Link: https://developer.cisco.com/docs/dna-center/#!the-trend-analytics-data-for-the-network-device-in-the-specified-time-range
*/
func (s *DevicesService) TheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRange(id string, requestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRange *RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRange) (*ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRange, *resty.Response, error) {
	path := "/dna/data/api/v1/networkDevices/{id}/trendAnalytics"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRange).
		SetResult(&ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRange{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.TheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRange(id, requestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRange)
		}

		return nil, response, fmt.Errorf("error with operation TheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRange")
	}

	result := response.Result().(*ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRange)
	return result, response, err

}

//CreatePlannedAccessPointForFloor Create Planned Access Point for Floor - 7eaa-8b15-454a-8c1d
/* Allows creation of a new planned access point on an existing floor map including its planned radio and antenna details.  Use the Get variant of this API to fetch any existing planned access points for the floor.  The payload to create a planned access point is in the same format, albeit a single object instead of a list, of that API.


@param floorID floorId path parameter. The instance UUID of the floor hierarchy element


Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-planned-access-point-for-floor
*/
func (s *DevicesService) CreatePlannedAccessPointForFloor(floorID string, requestDevicesCreatePlannedAccessPointForFloor *RequestDevicesCreatePlannedAccessPointForFloor) (*ResponseDevicesCreatePlannedAccessPointForFloor, *resty.Response, error) {
	path := "/dna/intent/api/v1/floors/{floorId}/planned-access-points"
	path = strings.Replace(path, "{floorId}", fmt.Sprintf("%v", floorID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesCreatePlannedAccessPointForFloor).
		SetResult(&ResponseDevicesCreatePlannedAccessPointForFloor{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreatePlannedAccessPointForFloor(floorID, requestDevicesCreatePlannedAccessPointForFloor)
		}

		return nil, response, fmt.Errorf("error with operation CreatePlannedAccessPointForFloor")
	}

	result := response.Result().(*ResponseDevicesCreatePlannedAccessPointForFloor)
	return result, response, err

}

//UpdateHealthScoreDefinitions Update health score definitions. - 1aab-193d-40bb-9d2f
/* Update health thresholds, include status of overall health status for each metric.
And also to synchronize with global profile issue thresholds of the definition for given metric. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-issueAndHealthDefinitions-1.0.0-resolved.yaml


@param UpdateHealthScoreDefinitionsHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!update-health-score-definitions
*/
func (s *DevicesService) UpdateHealthScoreDefinitions(requestDevicesUpdateHealthScoreDefinitions *RequestDevicesUpdateHealthScoreDefinitions, UpdateHealthScoreDefinitionsHeaderParams *UpdateHealthScoreDefinitionsHeaderParams) (*ResponseDevicesUpdateHealthScoreDefinitions, *resty.Response, error) {
	path := "/dna/intent/api/v1/healthScoreDefinitions/bulkUpdate"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if UpdateHealthScoreDefinitionsHeaderParams != nil {

		if UpdateHealthScoreDefinitionsHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", UpdateHealthScoreDefinitionsHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesUpdateHealthScoreDefinitions).
		SetResult(&ResponseDevicesUpdateHealthScoreDefinitions{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateHealthScoreDefinitions(requestDevicesUpdateHealthScoreDefinitions, UpdateHealthScoreDefinitionsHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation UpdateHealthScoreDefinitions")
	}

	result := response.Result().(*ResponseDevicesUpdateHealthScoreDefinitions)
	return result, response, err

}

//ClearMacAddressTable Clear Mac-Address table - 24be-a97f-43f9-bc65
/* Clear mac-address on an individual port. In request body, operation needs to be specified as 'ClearMacAddress'. In the future more possible operations will be added to this API


@param interfaceUUID interfaceUuid path parameter. Interface Id

@param ClearMacAddressTableQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!clear-mac-address-table
*/
func (s *DevicesService) ClearMacAddressTable(interfaceUUID string, requestDevicesClearMacAddressTable *RequestDevicesClearMacAddressTable, ClearMacAddressTableQueryParams *ClearMacAddressTableQueryParams) (*ResponseDevicesClearMacAddressTable, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/{interfaceUuid}/operation"
	path = strings.Replace(path, "{interfaceUuid}", fmt.Sprintf("%v", interfaceUUID), -1)

	queryString, _ := query.Values(ClearMacAddressTableQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetBody(requestDevicesClearMacAddressTable).
		SetResult(&ResponseDevicesClearMacAddressTable{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ClearMacAddressTable(interfaceUUID, requestDevicesClearMacAddressTable, ClearMacAddressTableQueryParams)
		}

		return nil, response, fmt.Errorf("error with operation ClearMacAddressTable")
	}

	result := response.Result().(*ResponseDevicesClearMacAddressTable)
	return result, response, err

}

//AddDevice2 Add Device - 4bb2-2af0-46fa-8f08
/* Adds the device with given credential



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-device2
*/
func (s *DevicesService) AddDevice2(requestDevicesAddDevice2 *RequestDevicesAddDevice2) (*ResponseDevicesAddDevice2, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesAddDevice2).
		SetResult(&ResponseDevicesAddDevice2{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddDevice2(requestDevicesAddDevice2)
		}

		return nil, response, fmt.Errorf("error with operation AddDevice2")
	}

	result := response.Result().(*ResponseDevicesAddDevice2)
	return result, response, err

}

//ExportDeviceList Export Device list - cd98-780f-4888-a66d
/* Exports the selected network device to a file



Documentation Link: https://developer.cisco.com/docs/dna-center/#!export-device-list
*/
func (s *DevicesService) ExportDeviceList(requestDevicesExportDeviceList *RequestDevicesExportDeviceList) (*ResponseDevicesExportDeviceList, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/file"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesExportDeviceList).
		SetResult(&ResponseDevicesExportDeviceList{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ExportDeviceList(requestDevicesExportDeviceList)
		}

		return nil, response, fmt.Errorf("error with operation ExportDeviceList")
	}

	result := response.Result().(*ResponseDevicesExportDeviceList)
	return result, response, err

}

//CreateUserDefinedField Create User-Defined-Field - 0a9c-18e7-4caa-8b07
/* Creates a new global User Defined Field, which can be assigned to devices



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-user-defined-field
*/
func (s *DevicesService) CreateUserDefinedField(requestDevicesCreateUserDefinedField *RequestDevicesCreateUserDefinedField) (*ResponseDevicesCreateUserDefinedField, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/user-defined-field"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesCreateUserDefinedField).
		SetResult(&ResponseDevicesCreateUserDefinedField{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateUserDefinedField(requestDevicesCreateUserDefinedField)
		}

		return nil, response, fmt.Errorf("error with operation CreateUserDefinedField")
	}

	result := response.Result().(*ResponseDevicesCreateUserDefinedField)
	return result, response, err

}

//OverrideResyncInterval Override resync interval - 42ac-59bd-41db-a4fe
/* Overrides the global resync interval on all network devices. This essentially removes device specific intervals if set.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!override-resync-interval
*/
func (s *DevicesService) OverrideResyncInterval() (*ResponseDevicesOverrideResyncInterval, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/resyncIntervalSettings/override"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesOverrideResyncInterval{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.OverrideResyncInterval()
		}

		return nil, response, fmt.Errorf("error with operation OverrideResyncInterval")
	}

	result := response.Result().(*ResponseDevicesOverrideResyncInterval)
	return result, response, err

}

//GetDeviceInterfaceStatsInfo Get Device Interface Stats Info - 76bb-5957-49ab-8a3b
/* This API returns the Interface Stats for the given Device Id. Please refer to the Feature tab for the Request Body usage and the API filtering support.


@param deviceID deviceId path parameter. Network Device Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-interface-stats-info
*/
func (s *DevicesService) GetDeviceInterfaceStatsInfo(deviceID string, requestDevicesGetDeviceInterfaceStatsInfo *RequestDevicesGetDeviceInterfaceStatsInfo) (*ResponseDevicesGetDeviceInterfaceStatsInfo, *resty.Response, error) {
	path := "/dna/intent/api/v2/networkDevices/{deviceId}/interfaces/query"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesGetDeviceInterfaceStatsInfo).
		SetResult(&ResponseDevicesGetDeviceInterfaceStatsInfo{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceInterfaceStatsInfo(deviceID, requestDevicesGetDeviceInterfaceStatsInfo)
		}

		return nil, response, fmt.Errorf("error with operation GetDeviceInterfaceStatsInfo")
	}

	result := response.Result().(*ResponseDevicesGetDeviceInterfaceStatsInfo)
	return result, response, err

}

//UpdatePlannedAccessPointForFloor Update Planned Access Point for Floor - 399f-596d-4f69-a080
/* Allows updating a planned access point on an existing floor map including its planned radio and antenna details.  Use the Get variant of this API to fetch the existing planned access points for the floor.  The payload to update a planned access point is in the same format, albeit a single object instead of a list, of that API.


@param floorID floorId path parameter. The instance UUID of the floor hierarchy element

*/
func (s *DevicesService) UpdatePlannedAccessPointForFloor(floorID string, requestDevicesUpdatePlannedAccessPointForFloor *RequestDevicesUpdatePlannedAccessPointForFloor) (*ResponseDevicesUpdatePlannedAccessPointForFloor, *resty.Response, error) {
	path := "/dna/intent/api/v1/floors/{floorId}/planned-access-points"
	path = strings.Replace(path, "{floorId}", fmt.Sprintf("%v", floorID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdatePlannedAccessPointForFloor).
		SetResult(&ResponseDevicesUpdatePlannedAccessPointForFloor{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdatePlannedAccessPointForFloor(floorID, requestDevicesUpdatePlannedAccessPointForFloor)
		}
		return nil, response, fmt.Errorf("error with operation UpdatePlannedAccessPointForFloor")
	}

	result := response.Result().(*ResponseDevicesUpdatePlannedAccessPointForFloor)
	return result, response, err

}

//UpdateHealthScoreDefinitionForTheGivenID Update health score definition for the given id. - f295-190f-4f08-bbe0
/* Update health threshold, include status of overall health status.
And also to synchronize with global profile issue thresholds of the definition for given id. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-issueAndHealthDefinitions-1.0.0-resolved.yaml


@param id id path parameter. Health score definition id.

*/
func (s *DevicesService) UpdateHealthScoreDefinitionForTheGivenID(id string, requestDevicesUpdateHealthScoreDefinitionForTheGivenId *RequestDevicesUpdateHealthScoreDefinitionForTheGivenID) (*ResponseDevicesUpdateHealthScoreDefinitionForTheGivenID, *resty.Response, error) {
	path := "/dna/intent/api/v1/healthScoreDefinitions/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdateHealthScoreDefinitionForTheGivenId).
		SetResult(&ResponseDevicesUpdateHealthScoreDefinitionForTheGivenID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateHealthScoreDefinitionForTheGivenID(id, requestDevicesUpdateHealthScoreDefinitionForTheGivenId)
		}
		return nil, response, fmt.Errorf("error with operation UpdateHealthScoreDefinitionForTheGivenId")
	}

	result := response.Result().(*ResponseDevicesUpdateHealthScoreDefinitionForTheGivenID)
	return result, response, err

}

//UpdateInterfaceDetails Update Interface details - 868b-5a60-4be8-a2d7
/* Add/Update Interface description, VLAN membership, Voice VLAN and change Interface admin status ('UP'/'DOWN') from Request body.


@param interfaceUUID interfaceUuid path parameter. Interface ID

@param UpdateInterfaceDetailsQueryParams Filtering parameter
*/
func (s *DevicesService) UpdateInterfaceDetails(interfaceUUID string, requestDevicesUpdateInterfaceDetails *RequestDevicesUpdateInterfaceDetails, UpdateInterfaceDetailsQueryParams *UpdateInterfaceDetailsQueryParams) (*ResponseDevicesUpdateInterfaceDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/{interfaceUuid}"
	path = strings.Replace(path, "{interfaceUuid}", fmt.Sprintf("%v", interfaceUUID), -1)

	queryString, _ := query.Values(UpdateInterfaceDetailsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetBody(requestDevicesUpdateInterfaceDetails).
		SetResult(&ResponseDevicesUpdateInterfaceDetails{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateInterfaceDetails(interfaceUUID, requestDevicesUpdateInterfaceDetails, UpdateInterfaceDetailsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation UpdateInterfaceDetails")
	}

	result := response.Result().(*ResponseDevicesUpdateInterfaceDetails)
	return result, response, err

}

//UpdateDeviceDetails Update Device Details - aeb9-eb67-460b-92df
/* Update the credentials, management IP address of a given device (or a set of devices) in Catalyst Center and trigger an inventory sync.


 */
func (s *DevicesService) UpdateDeviceDetails(requestDevicesUpdateDeviceDetails *RequestDevicesUpdateDeviceDetails) (*ResponseDevicesUpdateDeviceDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdateDeviceDetails).
		SetResult(&ResponseDevicesUpdateDeviceDetails{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateDeviceDetails(requestDevicesUpdateDeviceDetails)
		}
		return nil, response, fmt.Errorf("error with operation UpdateDeviceDetails")
	}

	result := response.Result().(*ResponseDevicesUpdateDeviceDetails)
	return result, response, err

}

//UpdateDeviceRole Update Device role - b985-5ad5-4ae9-8156
/* Updates the role of the device as access, core, distribution, border router


 */
func (s *DevicesService) UpdateDeviceRole(requestDevicesUpdateDeviceRole *RequestDevicesUpdateDeviceRole) (*ResponseDevicesUpdateDeviceRole, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/brief"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdateDeviceRole).
		SetResult(&ResponseDevicesUpdateDeviceRole{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateDeviceRole(requestDevicesUpdateDeviceRole)
		}
		return nil, response, fmt.Errorf("error with operation UpdateDeviceRole")
	}

	result := response.Result().(*ResponseDevicesUpdateDeviceRole)
	return result, response, err

}

//SyncDevices Sync Devices - 3b9e-f967-4429-be4c
/* Synchronizes the devices. If forceSync param is false (default) then the sync would run in normal priority thread. If forceSync param is true then the sync would run in high priority thread if available, else the sync will fail. Result can be seen in the child task of each device


@param SyncDevicesQueryParams Filtering parameter
*/
func (s *DevicesService) SyncDevices(requestDevicesSyncDevices *RequestDevicesSyncDevices, SyncDevicesQueryParams *SyncDevicesQueryParams) (*ResponseDevicesSyncDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/sync"

	queryString, _ := query.Values(SyncDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetBody(requestDevicesSyncDevices).
		SetResult(&ResponseDevicesSyncDevices{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SyncDevices(requestDevicesSyncDevices, SyncDevicesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation SyncDevices")
	}

	result := response.Result().(*ResponseDevicesSyncDevices)
	return result, response, err

}

//UpdateUserDefinedField Update User-Defined-Field - aa8c-ea8f-41aa-a346
/* Updates an existing global User Defined Field, using it's id.


@param id id path parameter. UDF id

*/
func (s *DevicesService) UpdateUserDefinedField(id string, requestDevicesUpdateUserDefinedField *RequestDevicesUpdateUserDefinedField) (*ResponseDevicesUpdateUserDefinedField, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/user-defined-field/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdateUserDefinedField).
		SetResult(&ResponseDevicesUpdateUserDefinedField{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateUserDefinedField(id, requestDevicesUpdateUserDefinedField)
		}
		return nil, response, fmt.Errorf("error with operation UpdateUserDefinedField")
	}

	result := response.Result().(*ResponseDevicesUpdateUserDefinedField)
	return result, response, err

}

//AddUserDefinedFieldToDevice Add User-Defined-Field to device - d3af-395c-4669-adaf
/* Assigns an existing Global User-Defined-Field to a device. If the UDF is already assigned to the specific device, then it updates the device UDF value accordingly. Please note that the assigning UDF 'name' must be an existing global UDF. Otherwise error shall be shown.


@param deviceID deviceId path parameter. UUID of device to which UDF has to be added

*/
func (s *DevicesService) AddUserDefinedFieldToDevice(deviceID string, requestDevicesAddUserDefinedFieldToDevice *RequestDevicesAddUserDefinedFieldToDevice) (*ResponseDevicesAddUserDefinedFieldToDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceId}/user-defined-field"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesAddUserDefinedFieldToDevice).
		SetResult(&ResponseDevicesAddUserDefinedFieldToDevice{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddUserDefinedFieldToDevice(deviceID, requestDevicesAddUserDefinedFieldToDevice)
		}
		return nil, response, fmt.Errorf("error with operation AddUserDefinedFieldToDevice")
	}

	result := response.Result().(*ResponseDevicesAddUserDefinedFieldToDevice)
	return result, response, err

}

//UpdateDeviceManagementAddress Update Device Management Address - af93-b807-4feb-a985
/* This is a simple PUT API to edit the management IP Address of the device.


@param deviceid deviceid path parameter. The UUID of the device whose management IP address is to be updated.

*/
func (s *DevicesService) UpdateDeviceManagementAddress(deviceid string, requestDevicesUpdateDeviceManagementAddress *RequestDevicesUpdateDeviceManagementAddress) (*ResponseDevicesUpdateDeviceManagementAddress, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceid}/management-address"
	path = strings.Replace(path, "{deviceid}", fmt.Sprintf("%v", deviceid), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdateDeviceManagementAddress).
		SetResult(&ResponseDevicesUpdateDeviceManagementAddress{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateDeviceManagementAddress(deviceid, requestDevicesUpdateDeviceManagementAddress)
		}
		return nil, response, fmt.Errorf("error with operation UpdateDeviceManagementAddress")
	}

	result := response.Result().(*ResponseDevicesUpdateDeviceManagementAddress)
	return result, response, err

}

//UpdateGlobalResyncInterval Update global resync interval - 25b5-39b4-4609-9e2a
/* Updates the resync interval (in minutes) globally for devices which do not have custom resync interval. To override this setting for all network devices refer to [/networkDevices/resyncIntervalSettings/override]


 */
func (s *DevicesService) UpdateGlobalResyncInterval(requestDevicesUpdateGlobalResyncInterval *RequestDevicesUpdateGlobalResyncInterval) (*ResponseDevicesUpdateGlobalResyncInterval, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/resyncIntervalSettings"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdateGlobalResyncInterval).
		SetResult(&ResponseDevicesUpdateGlobalResyncInterval{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateGlobalResyncInterval(requestDevicesUpdateGlobalResyncInterval)
		}
		return nil, response, fmt.Errorf("error with operation UpdateGlobalResyncInterval")
	}

	result := response.Result().(*ResponseDevicesUpdateGlobalResyncInterval)
	return result, response, err

}

//UpdateResyncIntervalForTheNetworkDevice Update resync interval for the network device - 92a0-db6c-428a-92d9
/* Update the resync interval (in minutes) for the given network device id.
To disable periodic resync, set interval as '0'.
To use global settings, set interval as 'null'.


@param id id path parameter. The id of the network device.

*/
func (s *DevicesService) UpdateResyncIntervalForTheNetworkDevice(id string, requestDevicesUpdateResyncIntervalForTheNetworkDevice *RequestDevicesUpdateResyncIntervalForTheNetworkDevice) (*ResponseDevicesUpdateResyncIntervalForTheNetworkDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/{id}/resyncIntervalSettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdateResyncIntervalForTheNetworkDevice).
		SetResult(&ResponseDevicesUpdateResyncIntervalForTheNetworkDevice{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateResyncIntervalForTheNetworkDevice(id, requestDevicesUpdateResyncIntervalForTheNetworkDevice)
		}
		return nil, response, fmt.Errorf("error with operation UpdateResyncIntervalForTheNetworkDevice")
	}

	result := response.Result().(*ResponseDevicesUpdateResyncIntervalForTheNetworkDevice)
	return result, response, err

}

//DeletePlannedAccessPointForFloor Delete Planned Access Point for Floor - 6dad-1aac-4b3a-9e67
/* Allow to delete a planned access point from an existing floor map including its planned radio and antenna details.  Use the Get variant of this API to fetch the existing planned access points for the floor.  The instanceUUID listed in each of the planned access point attributes acts as the path param input to this API to delete that specific instance.


@param floorID floorId path parameter. The instance UUID of the floor hierarchy element

@param plannedAccessPointUUID plannedAccessPointUuid path parameter. The instance UUID of the planned access point to delete


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-planned-access-point-for-floor
*/
func (s *DevicesService) DeletePlannedAccessPointForFloor(floorID string, plannedAccessPointUUID string) (*ResponseDevicesDeletePlannedAccessPointForFloor, *resty.Response, error) {
	//floorID string,plannedAccessPointUUID string
	path := "/dna/intent/api/v1/floors/{floorId}/planned-access-points/{plannedAccessPointUuid}"
	path = strings.Replace(path, "{floorId}", fmt.Sprintf("%v", floorID), -1)
	path = strings.Replace(path, "{plannedAccessPointUuid}", fmt.Sprintf("%v", plannedAccessPointUUID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesDeletePlannedAccessPointForFloor{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletePlannedAccessPointForFloor(floorID, plannedAccessPointUUID)
		}
		return nil, response, fmt.Errorf("error with operation DeletePlannedAccessPointForFloor")
	}

	result := response.Result().(*ResponseDevicesDeletePlannedAccessPointForFloor)
	return result, response, err

}

//DeleteUserDefinedField Delete User-Defined-Field - 78a3-c8b1-4799-892e
/* Deletes an existing Global User-Defined-Field using it's id.


@param id id path parameter. UDF id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-user-defined-field
*/
func (s *DevicesService) DeleteUserDefinedField(id string) (*ResponseDevicesDeleteUserDefinedField, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/network-device/user-defined-field/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesDeleteUserDefinedField{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteUserDefinedField(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteUserDefinedField")
	}

	result := response.Result().(*ResponseDevicesDeleteUserDefinedField)
	return result, response, err

}

//RemoveUserDefinedFieldFromDevice Remove User-Defined-Field from device - 8c9f-d9e8-4cab-bf96
/* Remove a User-Defined-Field from device. Name of UDF has to be passed as the query parameter. Please note that Global UDF will not be deleted by this operation.


@param deviceID deviceId path parameter. UUID of device from which UDF has to be removed

@param RemoveUserDefinedFieldFromDeviceQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!remove-user-defined-field-from-device
*/
func (s *DevicesService) RemoveUserDefinedFieldFromDevice(deviceID string, RemoveUserDefinedFieldFromDeviceQueryParams *RemoveUserDefinedFieldFromDeviceQueryParams) (*ResponseDevicesRemoveUserDefinedFieldFromDevice, *resty.Response, error) {
	//deviceID string,RemoveUserDefinedFieldFromDeviceQueryParams *RemoveUserDefinedFieldFromDeviceQueryParams
	path := "/dna/intent/api/v1/network-device/{deviceId}/user-defined-field"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	queryString, _ := query.Values(RemoveUserDefinedFieldFromDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesRemoveUserDefinedFieldFromDevice{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RemoveUserDefinedFieldFromDevice(deviceID, RemoveUserDefinedFieldFromDeviceQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RemoveUserDefinedFieldFromDevice")
	}

	result := response.Result().(*ResponseDevicesRemoveUserDefinedFieldFromDevice)
	return result, response, err

}

//DeleteDeviceByID Delete Device by Id - 1c89-4b58-48ea-b214
/* This API allows any network device that is not currently provisioned to be removed from the inventory. Important: Devices currently provisioned cannot be deleted. To delete a provisioned device, the device must be first deprovisioned.


@param id id path parameter. Device ID

@param DeleteDeviceByIdQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-device-by-id
*/
func (s *DevicesService) DeleteDeviceByID(id string, DeleteDeviceByIdQueryParams *DeleteDeviceByIDQueryParams) (*ResponseDevicesDeleteDeviceByID, *resty.Response, error) {
	//id string,DeleteDeviceByIdQueryParams *DeleteDeviceByIDQueryParams
	path := "/dna/intent/api/v1/network-device/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(DeleteDeviceByIdQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesDeleteDeviceByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteDeviceByID(id, DeleteDeviceByIdQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteDeviceById")
	}

	result := response.Result().(*ResponseDevicesDeleteDeviceByID)
	return result, response, err

}