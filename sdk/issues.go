package catalyst

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type IssuesService service

type GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1QueryParams struct {
	StartTime              float64 `url:"startTime,omitempty"`              //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to current time.
	EndTime                float64 `url:"endTime,omitempty"`                //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	Limit                  float64 `url:"limit,omitempty"`                  //Maximum number of issues to return
	Offset                 float64 `url:"offset,omitempty"`                 //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	SortBy                 string  `url:"sortBy,omitempty"`                 //
	Order                  string  `url:"order,omitempty"`                  //The sort order of the field ascending or descending.
	IsGlobal               bool    `url:"isGlobal,omitempty"`               //Global issues are those issues which impacts across many devices, sites. They are also displayed on Issue Dashboard in Catalyst Center UI. Non-Global issues are displayed only on Client 360 or Device 360 pages. If this flag is 'true', only global issues are returned. If it if 'false', all issues are returned.
	Priority               string  `url:"priority,omitempty"`               //Priority of the issue. Supports single priority and multiple priorities Examples: priority=P1 (single priority requested) priority=P1&priority=P2&priority=P3 (multiple priorities requested)
	Severity               string  `url:"severity,omitempty"`               //Severity of the issue. Supports single severity and multiple severities. Examples: severity=high (single severity requested) severity=high&severity=medium (multiple severities requested)
	Status                 string  `url:"status,omitempty"`                 //Status of the issue. Supports single status and multiple statuses. Examples: status=active (single status requested) status=active&status=resolved (multiple statuses requested)
	EntityType             string  `url:"entityType,omitempty"`             //Entity type of the issue. Supports single entity type and multiple entity types. Examples: entityType=networkDevice (single entity type requested) entityType=network device&entityType=client (multiple entity types requested)
	Category               string  `url:"category,omitempty"`               //Categories of the issue. Supports single category and multiple categories. Examples: category=availability (single status requested) category=availability&category=onboarding (multiple categories requested)
	DeviceType             string  `url:"deviceType,omitempty"`             //Device Type of the device to which this issue belongs to. Supports single device type and multiple device types. Examples: deviceType=wireless controller (single device type requested) deviceType=wireless controller&deviceType=core (multiple device types requested)
	Name                   string  `url:"name,omitempty"`                   //The name of the issue Examples: name=ap_down (single issue name requested) name=ap_down&name=wlc_monitor (multiple issue names requested) Issue names can be retrieved using the API - /data/api/v1/assuranceIssueConfigurations
	IssueID                string  `url:"issueId,omitempty"`                //UUID of the issue Examples: issueId=e52aecfe-b142-4287-a587-11a16ba6dd26 (single issue id requested) issueId=e52aecfe-b142-4287-a587-11a16ba6dd26&issueId=864d0421-02c0-43a6-9c52-81cad45f66d8 (multiple issue ids requested)
	EntityID               string  `url:"entityId,omitempty"`               //Id of the entity for which this issue belongs to. For example, it     could be mac address of AP or UUID of Sensor   example: 68:ca:e4:79:3f:20 4de02167-901b-43cf-8822-cffd3caa286f Examples: entityId=68:ca:e4:79:3f:20 (single entity id requested) entityId=68:ca:e4:79:3f:20&entityId=864d0421-02c0-43a6-9c52-81cad45f66d8 (multiple entity ids requested)
	UpdatedBy              string  `url:"updatedBy,omitempty"`              //The user who last updated this issue. Examples: updatedBy=admin (single updatedBy requested) updatedBy=admin&updatedBy=john (multiple updatedBy requested)
	SiteHierarchy          string  `url:"siteHierarchy,omitempty"`          //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. `Global/AreaName/BuildingName/FloorName`) This field supports wildcard asterisk (*) character search support. E.g. */San*, */San, /San* Examples: `?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)
	SiteHierarchyID        string  `url:"siteHierarchyId,omitempty"`        //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (*) character search support. E.g. `*uuid*, *uuid, uuid* Examples: `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	SiteName               string  `url:"siteName,omitempty"`               //The name of the site. (Ex. `FloorName`) This field supports wildcard asterisk (*) character search support. E.g. *San*, *San, San* Examples: `?siteName=building1` (single siteName requested) `?siteName=building1&siteName=building2&siteName=building3` (multiple siteNames requested)
	SiteID                 string  `url:"siteId,omitempty"`                 //The UUID of the site. (Ex. `flooruuid`) This field supports wildcard asterisk (*) character search support. E.g.*flooruuid*, *flooruuid, flooruuid* Examples: `?siteId=id1` (single id requested) `?siteId=id1&siteId=id2&siteId=id3` (multiple ids requested)
	FabricSiteID           string  `url:"fabricSiteId,omitempty"`           //The UUID of the fabric site. (Ex. "flooruuid") Examples: fabricSiteId=e52aecfe-b142-4287-a587-11a16ba6dd26 (single id requested) fabricSiteId=e52aecfe-b142-4287-a587-11a16ba6dd26,864d0421-02c0-43a6-9c52-81cad45f66d8 (multiple ids requested)
	FabricVnName           string  `url:"fabricVnName,omitempty"`           //The name of the fabric virtual network Examples: fabricVnName=name1 (single fabric virtual network name requested) fabricVnName=name1&fabricVnName=name2&fabricVnName=name3 (multiple fabric virtual network names requested)
	FabricTransitSiteID    string  `url:"fabricTransitSiteId,omitempty"`    //The UUID of the fabric transit site. (Ex. "flooruuid") Examples: fabricTransitSiteId=e52aecfe-b142-4287-a587-11a16ba6dd26 (single id requested) fabricTransitSiteId=e52aecfe-b142-4287-a587-11a16ba6dd26&fabricTransitSiteId=864d0421-02c0-43a6-9c52-81cad45f66d8 (multiple ids requested)
	NetworkDeviceID        string  `url:"networkDeviceId,omitempty"`        //The list of Network Device Uuids. (Ex. `6bef213c-19ca-4170-8375-b694e251101c`) Examples: `networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c` (single networkDeviceId requested) `networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c&networkDeviceId=32219612-819e-4b5e-a96b-cf22aca13dd9&networkDeviceId=2541e9a7-b80d-4955-8aa2-79b233318ba0` (multiple networkDeviceIds with & separator)
	NetworkDeviceIPAddress string  `url:"networkDeviceIpAddress,omitempty"` //The list of Network Device management IP Address. (Ex. `121.1.1.10`) This field supports wildcard (`*`) character-based search. Ex: `*1.1*` or `1.1*` or `*1.1` Examples: `networkDeviceIpAddress=121.1.1.10` `networkDeviceIpAddress=121.1.1.10&networkDeviceIpAddress=172.20.1.10&networkDeviceIpAddress=10.10.20.10` (multiple networkDevice IP Address with & separator)
	MacAddress             string  `url:"macAddress,omitempty"`             //The macAddress of the network device or client This field supports wildcard (`*`) character-based search. Ex: `*AB:AB:AB*` or `AB:AB:AB*` or `*AB:AB:AB` Examples: `macAddress=AB:AB:AB:CD:CD:CD` (single macAddress requested) `macAddress=AB:AB:AB:CD:CD:DC&macAddress=AB:AB:AB:CD:CD:FE` (multiple macAddress requested)
	View                   string  `url:"view,omitempty"`                   //The name of the View. Each view represents a specific data set. Please refer to the `IssuesView` Model for supported views. View is predefined set of attributes supported by the API. Only the attributes related to the given view will be part of the API response along with default attributes. If multiple views are provided, then response will contain attributes from all those views. If no views are specified, all attributes will be returned. | View Name | Included Attributes | | --- | --- | | `update` | updatedTime, updatedBy | | `site` | siteName, siteHierarchy, siteId, siteHierarchyId | Examples: `view=update` (single view requested) `view=update&view=site` (multiple views requested)
	Attribute              string  `url:"attribute,omitempty"`              //List of attributes related to the issue. If these are provided, then only those attributes will be part of response along with the default attributes. Please refer to the `IssuesResponseAttribute` Model for supported attributes. Examples: `attribute=deviceType` (single attribute requested) `attribute=deviceType&attribute=updatedBy` (multiple attributes requested)
	AiDriven               bool    `url:"aiDriven,omitempty"`               //Flag whether the issue is AI driven issue
	FabricDriven           bool    `url:"fabricDriven,omitempty"`           //Flag whether the issue is related to a Fabric site, a virtual network or a transit.
	FabricSiteDriven       bool    `url:"fabricSiteDriven,omitempty"`       //Flag whether the issue is Fabric site driven issue
	FabricVnDriven         bool    `url:"fabricVnDriven,omitempty"`         //Flag whether the issue is Fabric Virtual Network driven issue
	FabricTransitDriven    bool    `url:"fabricTransitDriven,omitempty"`    //Flag whether the issue is Fabric Transit driven issue
}
type GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1HeaderParams struct {
	AcceptLanguage string `url:"Accept-Language,omitempty"` //Expects type string. This header parameter can be used to specify the language in which issue description and suggested actions need to be returned. Available options are - 'en' (English), 'ja' (Japanese), 'ko' (Korean), 'zh' (Chinese). If this parameter is not present the issue details are returned in English language.
	XCaLLERID      string `url:"X-CALLER-ID,omitempty"`     //Expects type string. Caller ID can be used to trace the caller for queries executed on database. The caller id is like a optional attribute which can be added to API invocation like ui, python, postman, test-automation etc
}
type GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkV1QueryParams struct {
	StartTime              float64 `url:"startTime,omitempty"`              //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to current time.
	EndTime                float64 `url:"endTime,omitempty"`                //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	IsGlobal               bool    `url:"isGlobal,omitempty"`               //Global issues are those issues which impacts across many devices, sites. They are also displayed on Issue Dashboard in Catalyst Center UI. Non-Global issues are displayed only on Client 360 or Device 360 pages. If this flag is 'true', only global issues are returned. If it if 'false', all issues are returned.
	Priority               string  `url:"priority,omitempty"`               //Priority of the issue. Supports single priority and multiple priorities Examples: priority=P1 (single priority requested) priority=P1&priority=P2&priority=P3 (multiple priorities requested)
	Severity               string  `url:"severity,omitempty"`               //Severity of the issue. Supports single severity and multiple severities. Examples: severity=high (single severity requested) severity=high&severity=medium (multiple severities requested)
	Status                 string  `url:"status,omitempty"`                 //Status of the issue. Supports single status and multiple statuses. Examples: status=active (single status requested) status=active&status=resolved (multiple statuses requested)
	EntityType             string  `url:"entityType,omitempty"`             //Entity type of the issue. Supports single entity type and multiple entity types. Examples: entityType=networkDevice (single entity type requested) entityType=network device&entityType=client (multiple entity types requested)
	Category               string  `url:"category,omitempty"`               //Categories of the issue. Supports single category and multiple categories. Examples: category=availability (single status requested) category=availability&category=onboarding (multiple categories requested)
	DeviceType             string  `url:"deviceType,omitempty"`             //Device Type of the device to which this issue belongs to. Supports single device type and multiple device types. Examples: deviceType=wireless controller (single device type requested) deviceType=wireless controller&deviceType=core (multiple device types requested)
	Name                   string  `url:"name,omitempty"`                   //The name of the issue Examples: name=ap_down (single issue name requested) name=ap_down&name=wlc_monitor (multiple issue names requested) Issue names can be retrieved using the API - /data/api/v1/assuranceIssueConfigurations
	IssueID                string  `url:"issueId,omitempty"`                //UUID of the issue Examples: issueId=e52aecfe-b142-4287-a587-11a16ba6dd26 (single issue id requested) issueId=e52aecfe-b142-4287-a587-11a16ba6dd26&issueId=864d0421-02c0-43a6-9c52-81cad45f66d8 (multiple issue ids requested)
	EntityID               string  `url:"entityId,omitempty"`               //Id of the entity for which this issue belongs to. For example, it     could be mac address of AP or UUID of Sensor   example: 68:ca:e4:79:3f:20 4de02167-901b-43cf-8822-cffd3caa286f Examples: entityId=68:ca:e4:79:3f:20 (single entity id requested) entityId=68:ca:e4:79:3f:20&entityId=864d0421-02c0-43a6-9c52-81cad45f66d8 (multiple entity ids requested)
	UpdatedBy              string  `url:"updatedBy,omitempty"`              //The user who last updated this issue. Examples: updatedBy=admin (single updatedBy requested) updatedBy=admin&updatedBy=john (multiple updatedBy requested)
	SiteHierarchy          string  `url:"siteHierarchy,omitempty"`          //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. `Global/AreaName/BuildingName/FloorName`) This field supports wildcard asterisk (*) character search support. E.g. */San*, */San, /San* Examples: `?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)
	SiteHierarchyID        string  `url:"siteHierarchyId,omitempty"`        //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (*) character search support. E.g. `*uuid*, *uuid, uuid* Examples: `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	SiteName               string  `url:"siteName,omitempty"`               //The name of the site. (Ex. `FloorName`) This field supports wildcard asterisk (*) character search support. E.g. *San*, *San, San* Examples: `?siteName=building1` (single siteName requested) `?siteName=building1&siteName=building2&siteName=building3` (multiple siteNames requested)
	SiteID                 string  `url:"siteId,omitempty"`                 //The UUID of the site. (Ex. `flooruuid`) This field supports wildcard asterisk (*) character search support. E.g.*flooruuid*, *flooruuid, flooruuid* Examples: `?siteId=id1` (single id requested) `?siteId=id1&siteId=id2&siteId=id3` (multiple ids requested)
	FabricSiteID           string  `url:"fabricSiteId,omitempty"`           //The UUID of the fabric site. (Ex. "flooruuid") Examples: fabricSiteId=e52aecfe-b142-4287-a587-11a16ba6dd26 (single id requested) fabricSiteId=e52aecfe-b142-4287-a587-11a16ba6dd26,864d0421-02c0-43a6-9c52-81cad45f66d8 (multiple ids requested)
	FabricVnName           string  `url:"fabricVnName,omitempty"`           //The name of the fabric virtual network Examples: fabricVnName=name1 (single fabric virtual network name requested) fabricVnName=name1&fabricVnName=name2&fabricVnName=name3 (multiple fabric virtual network names requested)
	FabricTransitSiteID    string  `url:"fabricTransitSiteId,omitempty"`    //The UUID of the fabric transit site. (Ex. "flooruuid") Examples: fabricTransitSiteId=e52aecfe-b142-4287-a587-11a16ba6dd26 (single id requested) fabricTransitSiteId=e52aecfe-b142-4287-a587-11a16ba6dd26&fabricTransitSiteId=864d0421-02c0-43a6-9c52-81cad45f66d8 (multiple ids requested)
	NetworkDeviceID        string  `url:"networkDeviceId,omitempty"`        //The list of Network Device Uuids. (Ex. `6bef213c-19ca-4170-8375-b694e251101c`) Examples: `networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c` (single networkDeviceId requested) `networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c&networkDeviceId=32219612-819e-4b5e-a96b-cf22aca13dd9&networkDeviceId=2541e9a7-b80d-4955-8aa2-79b233318ba0` (multiple networkDeviceIds with & separator)
	NetworkDeviceIPAddress string  `url:"networkDeviceIpAddress,omitempty"` //The list of Network Device management IP Address. (Ex. `121.1.1.10`) This field supports wildcard (`*`) character-based search. Ex: `*1.1*` or `1.1*` or `*1.1` Examples: `networkDeviceIpAddress=121.1.1.10` `networkDeviceIpAddress=121.1.1.10&networkDeviceIpAddress=172.20.1.10&networkDeviceIpAddress=10.10.20.10` (multiple networkDevice IP Address with & separator)
	MacAddress             string  `url:"macAddress,omitempty"`             //The macAddress of the network device or client This field supports wildcard (`*`) character-based search. Ex: `*AB:AB:AB*` or `AB:AB:AB*` or `*AB:AB:AB` Examples: `macAddress=AB:AB:AB:CD:CD:CD` (single macAddress requested) `macAddress=AB:AB:AB:CD:CD:DC&macAddress=AB:AB:AB:CD:CD:FE` (multiple macAddress requested)
	AiDriven               bool    `url:"aiDriven,omitempty"`               //Flag whether the issue is AI driven issue
	FabricDriven           bool    `url:"fabricDriven,omitempty"`           //Flag whether the issue is related to a Fabric site, a virtual network or a transit.
	FabricSiteDriven       bool    `url:"fabricSiteDriven,omitempty"`       //Flag whether the issue is Fabric site driven issue
	FabricVnDriven         bool    `url:"fabricVnDriven,omitempty"`         //Flag whether the issue is Fabric Virtual Network driven issue
	FabricTransitDriven    bool    `url:"fabricTransitDriven,omitempty"`    //Flag whether the issue is Fabric Transit driven issue
}
type GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID can be used to trace the caller for queries executed on database. The caller id is like a optional attribute which can be added to API invocation like ui, python, postman, test-automation etc
}
type GetTheDetailsOfIssuesForGivenSetOfFiltersV1HeaderParams struct {
	ContentType    string `url:"Content-Type,omitempty"`    //Expects type string. Request body content type
	AcceptLanguage string `url:"Accept-Language,omitempty"` //Expects type string. This header parameter can be used to specify the language in which issue description and suggested actions need to be returned. Available options are - 'en' (English), 'ja' (Japanese), 'ko' (Korean), 'zh' (Chinese). If this parameter is not present the issue details are returned in English language.
	XCaLLERID      string `url:"X-CALLER-ID,omitempty"`     //Expects type string. Caller ID can be used to trace the caller for queries executed on database. The caller id is like a optional attribute which can be added to API invocation like ui, python, postman, test-automation etc
}
type GetTheTotalNumberOfIssuesForGivenSetOfFiltersV1HeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID can be used to trace the caller for queries executed on database. The caller id is like a optional attribute which can be added to API invocation like ui, python, postman, test-automation etc
}
type GetSummaryAnalyticsDataOfIssuesV1HeaderParams struct {
	ContentType    string `url:"Content-Type,omitempty"`    //Expects type string. Request body content type
	AcceptLanguage string `url:"Accept-Language,omitempty"` //Expects type string. This header parameter can be used to specify the language in which issue display name need to be returned. Available options are - 'en' (English), 'ja' (Japanese), 'ko' (Korean), 'zh' (Chinese). If this parameter is not present the issue display name is returned in English language.
	XCaLLERID      string `url:"X-CALLER-ID,omitempty"`     //Expects type string. Caller ID can be used to trace the caller for queries executed on database. The caller id is like a optional attribute which can be added to API invocation like ui, python, postman, test-automation etc
}
type GetTopNAnalyticsDataOfIssuesV1HeaderParams struct {
	ContentType    string `url:"Content-Type,omitempty"`    //Expects type string. Request body content type
	AcceptLanguage string `url:"Accept-Language,omitempty"` //Expects type string. This header parameter can be used to specify the language in which issue display name need to be returned. Available options are - 'en' (English), 'ja' (Japanese), 'ko' (Korean), 'zh' (Chinese). If this parameter is not present the issue display name is returned in English language.
	XCaLLERID      string `url:"X-CALLER-ID,omitempty"`     //Expects type string. Caller ID can be used to trace the caller for queries executed on database. The caller id is like a optional attribute which can be added to API invocation like ui, python, postman, test-automation etc
}
type GetTrendAnalyticsDataOfIssuesV1HeaderParams struct {
	ContentType    string `url:"Content-Type,omitempty"`    //Expects type string. Request body content type
	AcceptLanguage string `url:"Accept-Language,omitempty"` //Expects type string. This header parameter can be used to specify the language in which issue display name need to be returned. Available options are - 'en' (English), 'ja' (Japanese), 'ko' (Korean), 'zh' (Chinese). If this parameter is not present the issue display name is returned in English language.
	XCaLLERID      string `url:"X-CALLER-ID,omitempty"`     //Expects type string. Caller ID can be used to trace the caller for queries executed on database. The caller id is like a optional attribute which can be added to API invocation like ui, python, postman, test-automation etc
}
type GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDV1QueryParams struct {
	View      string `url:"view,omitempty"`      //The name of the View. Each view represents a specific data set. Please refer to the `IssuesView` Model for supported views. View is predefined set of attributes supported by the API. Only the attributes related to the given view will be part of the API response along with default attributes. If multiple views are provided, then response will contain attributes from all those views. If no views are specified, all attributes will be returned. | View Name | Included Attributes | | --- | --- | | `update` | updatedTime, updatedBy | | `site` | siteName, siteHierarchy, siteId, siteHierarchyId | Examples: `view=update` (single view requested) `view=update&view=site` (multiple views requested)
	Attribute string `url:"attribute,omitempty"` //List of attributes related to the issue. If these are provided, then only those attributes will be part of response along with the default attributes. Please refer to the `IssuesResponseAttribute` Model for supported attributes. Examples: `attribute=deviceType` (single attribute requested) `attribute=deviceType&attribute=updatedBy` (multiple attributes requested)
}
type GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDV1HeaderParams struct {
	AcceptLanguage string `url:"Accept-Language,omitempty"` //Expects type string. This header parameter can be used to specify the language in which issue description and suggested actions need to be returned. Available options are - 'en' (English), 'ja' (Japanese), 'ko' (Korean), 'zh' (Chinese). If this parameter is not present the issue details are returned in English language.
	XCaLLERID      string `url:"X-CALLER-ID,omitempty"`     //Expects type string. Caller ID can be used to trace the caller for queries executed on database. The caller id is like a optional attribute which can be added to API invocation like ui, python, postman, test-automation etc
}
type IgnoreTheGivenListOfIssuesV1HeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID can be used to trace the caller for queries executed on database. The caller id is like a optional attribute which can be added to API invocation like ui, python, postman, test-automation etc
}
type ResolveTheGivenListsOfIssuesV1HeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID can be used to trace the caller for queries executed on database. The caller id is like a optional attribute which can be added to API invocation like ui, python, postman, test-automation etc
}
type UpdateTheGivenIssueByUpdatingSelectedFieldsV1HeaderParams struct {
	ContentType    string `url:"Content-Type,omitempty"`    //Expects type string. Request body content type
	AcceptLanguage string `url:"Accept-Language,omitempty"` //Expects type string. This header parameter can be used to specify the language in which issue description and suggested actions need to be returned. Available options are - 'en' (English), 'ja' (Japanese), 'ko' (Korean), 'zh' (Chinese). If this parameter is not present the issue details are returned in English language.
	XCaLLERID      string `url:"X-CALLER-ID,omitempty"`     //Expects type string. Caller ID can be used to trace the caller for queries executed on database. The caller id is like a optional attribute which can be added to API invocation like ui, python, postman, test-automation etc
}
type CreatesANewUserDefinedIssueDefinitionsV1HeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersV1QueryParams struct {
	ID        string  `url:"id,omitempty"`        //The custom issue definition identifier and unique identifier across the profile.Examples: id=6bef213c-19ca-4170-8375-b694e251101c (single entity uuid requested) id=6bef213c-19ca-4170-8375-b694e251101c&id=19ca-4170-8375-b694e251101c-6bef213c (multiple Id request in the query param)
	ProfileID string  `url:"profileId,omitempty"` //The profile identifier to fetch the profile associated custom issue definitions. The default is global. For the custom profile, it is profile UUID. Example : 3fa85f64-5717-4562-b3fc-2c963f66afa6
	Name      string  `url:"name,omitempty"`      //The list of UDI issue names
	Priority  string  `url:"priority,omitempty"`  //The Issue priority value, possible values are P1, P2, P3, P4. P1: A critical issue that needs immediate attention and can have a wide impact on network operations. P2: A major issue that can potentially impact multiple devices or clients. P3: A minor issue that has a localized or minimal impact. P4: A warning issue that may not be an immediate problem but addressing it can optimize the network performance
	IsEnabled bool    `url:"isEnabled,omitempty"` //The enable status of the custom issue definition, either true or false.
	Severity  float64 `url:"severity,omitempty"`  //The syslog severity level. 0: Emergency 1: Alert, 2: Critical. 3: Error, 4: Warning, 5: Notice, 6: Info. Examples:severity=1&severity=2 (multi value support with & separator)
	Facility  string  `url:"facility,omitempty"`  //The syslog facility name
	Mnemonic  string  `url:"mnemonic,omitempty"`  //The syslog mnemonic name
	Limit     float64 `url:"limit,omitempty"`     //The maximum number of records to return
	Offset    float64 `url:"offset,omitempty"`    //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	SortBy    string  `url:"sortBy,omitempty"`    //A field within the response to sort by.
	Order     string  `url:"order,omitempty"`     //The sort order of the field ascending or descending.
}
type GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersV1QueryParams struct {
	ID        string  `url:"id,omitempty"`        //The custom issue definition identifier and unique identifier across the profile. Examples: id=6bef213c-19ca-4170-8375-b694e251101c (single entity uuid requested) id=6bef213c-19ca-4170-8375-b694e251101c&id=19ca-4170-8375-b694e251101c-6bef213c (multiple Id request in the query param)
	ProfileID string  `url:"profileId,omitempty"` //The profile identifier to fetch the profile associated custom issue definitions. The default is global. For the custom profile, it is profile UUID. Example : 3fa85f64-5717-4562-b3fc-2c963f66afa6
	Name      string  `url:"name,omitempty"`      //The list of UDI issue names. (Ex."TestUdiIssues")
	Priority  string  `url:"priority,omitempty"`  //The Issue priority value, possible values are P1, P2, P3, P4. P1: A critical issue that needs immediate attention and can have a wide impact on network operations. P2: A major issue that can potentially impact multiple devices or clients. P3: A minor issue that has a localized or minimal impact. P4: A warning issue that may not be an immediate problem but addressing it can optimize the network performance
	IsEnabled bool    `url:"isEnabled,omitempty"` //The enable status of the custom issue definition, either true or false.
	Severity  float64 `url:"severity,omitempty"`  //The syslog severity level. 0: Emergency 1: Alert, 2: Critical. 3: Error, 4: Warning, 5: Notice, 6: Info. Examples:severity=1&severity=2 (multi value support with & separator)
	Facility  string  `url:"facility,omitempty"`  //The syslog facility name
	Mnemonic  string  `url:"mnemonic,omitempty"`  //The syslog mnemonic name
}
type GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID can be used to trace the caller for queries executed on database. The caller id is like a optional attribute which can be added to API invocation like ui, python, postman, test-automation etc
}
type GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID can be used to trace the caller for queries executed on database. The caller id is like a optional attribute which can be added to API invocation like ui, python, postman, test-automation etc
}
type GetIssueEnrichmentDetailsV1HeaderParams struct {
	EntityType        string `url:"entity_type,omitempty"`         //Expects type string. Issue enrichment details can be fetched based on either Issue ID or Client MAC address. This parameter value must either be issue_id/mac_address
	EntityValue       string `url:"entity_value,omitempty"`        //Expects type string. Contains the actual value for the entity type that has been defined
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type bool. For the enrichment details to be made available as part of the API response, this header must be set to true. This header must be explicitly passed when called from client applications outside Catalyst Center
}
type IssuesV1QueryParams struct {
	StartTime   float64 `url:"startTime,omitempty"`   //Starting epoch time in milliseconds of query time window
	EndTime     float64 `url:"endTime,omitempty"`     //Ending epoch time in milliseconds of query time window
	SiteID      string  `url:"siteId,omitempty"`      //Assurance UUID value of the site in the issue content
	DeviceID    string  `url:"deviceId,omitempty"`    //Assurance UUID value of the device in the issue content
	MacAddress  string  `url:"macAddress,omitempty"`  //Client's device MAC address of the issue (format xx:xx:xx:xx:xx:xx)
	Priority    string  `url:"priority,omitempty"`    //The issue's priority value: P1, P2, P3, or P4 (case insensitive) (Use only when macAddress and deviceId are not provided)
	IssueStatus string  `url:"issueStatus,omitempty"` //The issue's status value: ACTIVE, IGNORED, RESOLVED (case insensitive)
	AiDriven    string  `url:"aiDriven,omitempty"`    //The issue's AI driven value: YES or NO (case insensitive) (Use only when macAddress and deviceId are not provided)
}
type ReturnsAllIssueTriggerDefinitionsForGivenFiltersV1QueryParams struct {
	DeviceType   string  `url:"deviceType,omitempty"`   //These are the device families/types supported for system issue definitions. If no input is made on device type, all device types are considered.
	ProfileID    string  `url:"profileId,omitempty"`    //The profile identier to fetch the profile associated issue defintions. The default is `global`. Please refer Network design profiles documentation for more details.
	ID           string  `url:"id,omitempty"`           //The definition identifier. Examples: id=015d9cba-4f53-4087-8317-7e49e5ffef46 (single entity id request) id=015d9cba-4f53-4087-8317-7e49e5ffef46&id=015d9cba-4f53-4087-8317-7e49e5ffef47 (multiple ids in the query param)
	Name         string  `url:"name,omitempty"`         //The list of system defined issue names. (Ex."BGP_Down") Examples: name=BGP_Down (single entity uuid requested) name=BGP_Down&name=BGP_Flap (multiple issue names separated by & operator)
	Priority     string  `url:"priority,omitempty"`     //Issue priority, possible values are P1, P2, P3, P4. `P1`: A critical issue that needs immediate attention and can have a wide impact on network operations. `P2`: A major issue that can potentially impact multiple devices or clients. `P3`: A minor issue that has a localized or minimal impact. `P4`: A warning issue that may not be an immediate problem but addressing it can optimize the network performance.
	IssueEnabled bool    `url:"issueEnabled,omitempty"` //The enablement status of the issue definition, either true or false.
	Attribute    string  `url:"attribute,omitempty"`    //These are the attributes supported in system issue definitions response. By default, all properties are sent in response.
	Offset       float64 `url:"offset,omitempty"`       //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	Limit        float64 `url:"limit,omitempty"`        //Maximum number of records to return
	SortBy       string  `url:"sortBy,omitempty"`       //A field within the response to sort by.
	Order        string  `url:"order,omitempty"`        //The sort order of the field ascending or descending.
}
type ReturnsAllIssueTriggerDefinitionsForGivenFiltersV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1QueryParams struct {
	DeviceType   string `url:"deviceType,omitempty"`   //These are the device families/types supported for system issue definitions. If no input is made on device type, all device types are considered.
	ProfileID    string `url:"profileId,omitempty"`    //The profile identier to fetch the profile associated issue defintions. The default is `global`. Please refer Network design profiles documentation for more details.
	ID           string `url:"id,omitempty"`           //The definition identifier. Examples: id=015d9cba-4f53-4087-8317-7e49e5ffef46 (single entity id request) id=015d9cba-4f53-4087-8317-7e49e5ffef46&id=015d9cba-4f53-4087-8317-7e49e5ffef47 (multiple ids in the query param)
	Name         string `url:"name,omitempty"`         //The list of system defined issue names. (Ex."BGP_Down") Examples: name=BGP_Down (single entity uuid requested) name=BGP_Down&name=BGP_Flap (multiple issue names separated by & operator)
	Priority     string `url:"priority,omitempty"`     //Issue priority, possible values are P1, P2, P3, P4. `P1`: A critical issue that needs immediate attention and can have a wide impact on network operations. `P2`: A major issue that can potentially impact multiple devices or clients. `P3`: A minor issue that has a localized or minimal impact. `P4`: A warning issue that may not be an immediate problem but addressing it can optimize the network performance.
	IssueEnabled bool   `url:"issueEnabled,omitempty"` //The enablement status of the issue definition, either true or false.
}
type GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetIssueTriggerDefinitionForGivenIDV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}

type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1 struct {
	Response *[]ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1Response `json:"response,omitempty"` //
	Page     *ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1Page       `json:"page,omitempty"`     //
	Version  string                                                                              `json:"version,omitempty"`  // Version
}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1Response struct {
	IssueID                string                                                                                                  `json:"issueId,omitempty"`                // Issue Id
	Name                   string                                                                                                  `json:"name,omitempty"`                   // Name
	Description            string                                                                                                  `json:"description,omitempty"`            // Description
	Summary                string                                                                                                  `json:"summary,omitempty"`                // Summary
	Priority               string                                                                                                  `json:"priority,omitempty"`               // Priority
	Severity               string                                                                                                  `json:"severity,omitempty"`               // Severity
	DeviceType             string                                                                                                  `json:"deviceType,omitempty"`             // Device Type
	Category               string                                                                                                  `json:"category,omitempty"`               // Category
	EntityType             string                                                                                                  `json:"entityType,omitempty"`             // Entity Type
	EntityID               string                                                                                                  `json:"entityId,omitempty"`               // Entity Id
	FirstOccurredTime      *int                                                                                                    `json:"firstOccurredTime,omitempty"`      // First Occurred Time
	MostRecentOccurredTime *int                                                                                                    `json:"mostRecentOccurredTime,omitempty"` // Most Recent Occurred Time
	Status                 string                                                                                                  `json:"status,omitempty"`                 // Status
	IsGlobal               *bool                                                                                                   `json:"isGlobal,omitempty"`               // Is Global
	UpdatedBy              *ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1ResponseUpdatedBy              `json:"updatedBy,omitempty"`              // Updated By
	UpdatedTime            *ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1ResponseUpdatedTime            `json:"updatedTime,omitempty"`            // Updated Time
	Notes                  *ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1ResponseNotes                  `json:"notes,omitempty"`                  // Notes
	SiteID                 *ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1ResponseSiteID                 `json:"siteId,omitempty"`                 // Site Id
	SiteHierarchyID        *ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1ResponseSiteHierarchyID        `json:"siteHierarchyId,omitempty"`        // Site Hierarchy Id
	SiteName               *ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1ResponseSiteName               `json:"siteName,omitempty"`               // Site Name
	SiteHierarchy          *ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1ResponseSiteHierarchy          `json:"siteHierarchy,omitempty"`          // Site Hierarchy
	SuggestedActions       *[]ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1ResponseSuggestedActions     `json:"suggestedActions,omitempty"`       //
	AdditionalAttributes   *[]ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1ResponseAdditionalAttributes `json:"additionalAttributes,omitempty"`   //
}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1ResponseUpdatedBy interface{}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1ResponseUpdatedTime interface{}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1ResponseNotes interface{}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1ResponseSiteID interface{}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1ResponseSiteHierarchyID interface{}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1ResponseSiteName interface{}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1ResponseSiteHierarchy interface{}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1ResponseSuggestedActions struct {
	Message string                                                                                                   `json:"message,omitempty"` // Message
	Steps   *[]ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1ResponseSuggestedActionsSteps `json:"steps,omitempty"`   // Steps
}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1ResponseSuggestedActionsSteps interface{}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1ResponseAdditionalAttributes struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1Page struct {
	Limit  *int                                                                                  `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                  `json:"offset,omitempty"` // Offset
	Count  *int                                                                                  `json:"count,omitempty"`  // Count
	SortBy *[]ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1PageSortBy `json:"sortBy,omitempty"` //
}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1PageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkV1 struct {
	Response *ResponseIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkV1Response `json:"response,omitempty"` //
	Version  string                                                                                `json:"version,omitempty"`  // Version
}
type ResponseIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1 struct {
	Response *[]ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1Response `json:"response,omitempty"` //
	Page     *ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1Page       `json:"page,omitempty"`     //
	Version  string                                                               `json:"version,omitempty"`  // Version
}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1Response struct {
	IssueID                string                                                                                   `json:"issueId,omitempty"`                // Issue Id
	Name                   string                                                                                   `json:"name,omitempty"`                   // Name
	Description            string                                                                                   `json:"description,omitempty"`            // Description
	Summary                string                                                                                   `json:"summary,omitempty"`                // Summary
	Priority               string                                                                                   `json:"priority,omitempty"`               // Priority
	Severity               string                                                                                   `json:"severity,omitempty"`               // Severity
	DeviceType             string                                                                                   `json:"deviceType,omitempty"`             // Device Type
	Category               string                                                                                   `json:"category,omitempty"`               // Category
	EntityType             string                                                                                   `json:"entityType,omitempty"`             // Entity Type
	EntityID               string                                                                                   `json:"entityId,omitempty"`               // Entity Id
	FirstOccurredTime      *int                                                                                     `json:"firstOccurredTime,omitempty"`      // First Occurred Time
	MostRecentOccurredTime *int                                                                                     `json:"mostRecentOccurredTime,omitempty"` // Most Recent Occurred Time
	Status                 string                                                                                   `json:"status,omitempty"`                 // Status
	IsGlobal               *bool                                                                                    `json:"isGlobal,omitempty"`               // Is Global
	UpdatedBy              *ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ResponseUpdatedBy              `json:"updatedBy,omitempty"`              // Updated By
	UpdatedTime            *ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ResponseUpdatedTime            `json:"updatedTime,omitempty"`            // Updated Time
	Notes                  *ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ResponseNotes                  `json:"notes,omitempty"`                  // Notes
	SiteID                 *ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ResponseSiteID                 `json:"siteId,omitempty"`                 // Site Id
	SiteHierarchyID        *ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ResponseSiteHierarchyID        `json:"siteHierarchyId,omitempty"`        // Site Hierarchy Id
	SiteName               *ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ResponseSiteName               `json:"siteName,omitempty"`               // Site Name
	SiteHierarchy          *ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ResponseSiteHierarchy          `json:"siteHierarchy,omitempty"`          // Site Hierarchy
	SuggestedActions       *[]ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ResponseSuggestedActions     `json:"suggestedActions,omitempty"`       //
	AdditionalAttributes   *[]ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ResponseAdditionalAttributes `json:"additionalAttributes,omitempty"`   //
}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ResponseUpdatedBy interface{}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ResponseUpdatedTime interface{}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ResponseNotes interface{}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ResponseSiteID interface{}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ResponseSiteHierarchyID interface{}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ResponseSiteName interface{}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ResponseSiteHierarchy interface{}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ResponseSuggestedActions struct {
	Message string                                                                                    `json:"message,omitempty"` // Message
	Steps   *[]ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ResponseSuggestedActionsSteps `json:"steps,omitempty"`   // Steps
}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ResponseSuggestedActionsSteps interface{}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ResponseAdditionalAttributes struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1Page struct {
	Limit  *int                                                                   `json:"limit,omitempty"`  // Limit
	Offset *int                                                                   `json:"offset,omitempty"` // Offset
	Count  *int                                                                   `json:"count,omitempty"`  // Count
	SortBy *[]ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1PageSortBy `json:"sortBy,omitempty"` //
}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1PageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersV1 struct {
	Filters *[]ResponseIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersV1Filters `json:"filters,omitempty"` //
}
type ResponseIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersV1Filters struct {
	Key      string `json:"key,omitempty"`      // Key
	Value    string `json:"value,omitempty"`    // Value
	Operator string `json:"operator,omitempty"` // Operator
}
type ResponseIssuesGetSummaryAnalyticsDataOfIssuesV1 struct {
	Version  string                                                   `json:"version,omitempty"`  // Version
	Response *ResponseIssuesGetSummaryAnalyticsDataOfIssuesV1Response `json:"response,omitempty"` //
	Page     *ResponseIssuesGetSummaryAnalyticsDataOfIssuesV1Page     `json:"page,omitempty"`     //
}
type ResponseIssuesGetSummaryAnalyticsDataOfIssuesV1Response struct {
	Groups              *[]ResponseIssuesGetSummaryAnalyticsDataOfIssuesV1ResponseGroups              `json:"groups,omitempty"`              //
	Attributes          *[]ResponseIssuesGetSummaryAnalyticsDataOfIssuesV1ResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseIssuesGetSummaryAnalyticsDataOfIssuesV1ResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseIssuesGetSummaryAnalyticsDataOfIssuesV1ResponseGroups struct {
	ID                  string                                                                              `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseIssuesGetSummaryAnalyticsDataOfIssuesV1ResponseGroupsAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseIssuesGetSummaryAnalyticsDataOfIssuesV1ResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseIssuesGetSummaryAnalyticsDataOfIssuesV1ResponseGroupsAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseIssuesGetSummaryAnalyticsDataOfIssuesV1ResponseGroupsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseIssuesGetSummaryAnalyticsDataOfIssuesV1ResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseIssuesGetSummaryAnalyticsDataOfIssuesV1ResponseAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseIssuesGetSummaryAnalyticsDataOfIssuesV1Page struct {
	Limit  *int                                                         `json:"limit,omitempty"`  // Limit
	Offset *int                                                         `json:"offset,omitempty"` // Offset
	Count  *int                                                         `json:"count,omitempty"`  // Count
	SortBy *[]ResponseIssuesGetSummaryAnalyticsDataOfIssuesV1PageSortBy `json:"sortBy,omitempty"` //
}
type ResponseIssuesGetSummaryAnalyticsDataOfIssuesV1PageSortBy struct {
	Name     string                                                             `json:"name,omitempty"`     // Name
	Function *ResponseIssuesGetSummaryAnalyticsDataOfIssuesV1PageSortByFunction `json:"function,omitempty"` // Function
	Order    string                                                             `json:"order,omitempty"`    // Order
}
type ResponseIssuesGetSummaryAnalyticsDataOfIssuesV1PageSortByFunction interface{}
type ResponseIssuesGetTopNAnalyticsDataOfIssuesV1 struct {
	Version  string                                                  `json:"version,omitempty"`  // Version
	Response *[]ResponseIssuesGetTopNAnalyticsDataOfIssuesV1Response `json:"response,omitempty"` //
	Page     *ResponseIssuesGetTopNAnalyticsDataOfIssuesV1Page       `json:"page,omitempty"`     //
}
type ResponseIssuesGetTopNAnalyticsDataOfIssuesV1Response struct {
	ID                  string                                                                     `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseIssuesGetTopNAnalyticsDataOfIssuesV1ResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseIssuesGetTopNAnalyticsDataOfIssuesV1ResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseIssuesGetTopNAnalyticsDataOfIssuesV1ResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseIssuesGetTopNAnalyticsDataOfIssuesV1ResponseAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseIssuesGetTopNAnalyticsDataOfIssuesV1Page struct {
	Limit  *int                                                      `json:"limit,omitempty"`  // Limit
	Offset *int                                                      `json:"offset,omitempty"` // Offset
	Count  *int                                                      `json:"count,omitempty"`  // Count
	SortBy *[]ResponseIssuesGetTopNAnalyticsDataOfIssuesV1PageSortBy `json:"sortBy,omitempty"` //
}
type ResponseIssuesGetTopNAnalyticsDataOfIssuesV1PageSortBy struct {
	Name     string                                                          `json:"name,omitempty"`     // Name
	Function *ResponseIssuesGetTopNAnalyticsDataOfIssuesV1PageSortByFunction `json:"function,omitempty"` // Function
	Order    string                                                          `json:"order,omitempty"`    // Order
}
type ResponseIssuesGetTopNAnalyticsDataOfIssuesV1PageSortByFunction interface{}
type ResponseIssuesGetTrendAnalyticsDataOfIssuesV1 struct {
	Version  string                                                   `json:"version,omitempty"`  // Version
	Response *[]ResponseIssuesGetTrendAnalyticsDataOfIssuesV1Response `json:"response,omitempty"` //
	Page     *ResponseIssuesGetTrendAnalyticsDataOfIssuesV1Page       `json:"page,omitempty"`     //
}
type ResponseIssuesGetTrendAnalyticsDataOfIssuesV1Response struct {
	Timestamp           *int                                                                        `json:"timestamp,omitempty"`           // Timestamp
	Groups              *[]ResponseIssuesGetTrendAnalyticsDataOfIssuesV1ResponseGroups              `json:"groups,omitempty"`              //
	Attributes          *[]ResponseIssuesGetTrendAnalyticsDataOfIssuesV1ResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseIssuesGetTrendAnalyticsDataOfIssuesV1ResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseIssuesGetTrendAnalyticsDataOfIssuesV1ResponseGroups struct {
	ID                  string                                                                            `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseIssuesGetTrendAnalyticsDataOfIssuesV1ResponseGroupsAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseIssuesGetTrendAnalyticsDataOfIssuesV1ResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseIssuesGetTrendAnalyticsDataOfIssuesV1ResponseGroupsAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseIssuesGetTrendAnalyticsDataOfIssuesV1ResponseGroupsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseIssuesGetTrendAnalyticsDataOfIssuesV1ResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseIssuesGetTrendAnalyticsDataOfIssuesV1ResponseAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseIssuesGetTrendAnalyticsDataOfIssuesV1Page struct {
	Limit          *int   `json:"limit,omitempty"`          // Limit
	Offset         *int   `json:"offset,omitempty"`         // Offset
	Count          *int   `json:"count,omitempty"`          // Count
	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDV1 struct {
	Response *ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDV1Response `json:"response,omitempty"` //
	Version  string                                                                                  `json:"version,omitempty"`  // Version
}
type ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDV1Response struct {
	IssueID                string                                                                                                        `json:"issueId,omitempty"`                // Issue Id
	Name                   string                                                                                                        `json:"name,omitempty"`                   // Name
	Description            string                                                                                                        `json:"description,omitempty"`            // Description
	Summary                string                                                                                                        `json:"summary,omitempty"`                // Summary
	Priority               string                                                                                                        `json:"priority,omitempty"`               // Priority
	Severity               string                                                                                                        `json:"severity,omitempty"`               // Severity
	DeviceType             string                                                                                                        `json:"deviceType,omitempty"`             // Device Type
	Category               string                                                                                                        `json:"category,omitempty"`               // Category
	EntityType             string                                                                                                        `json:"entityType,omitempty"`             // Entity Type
	EntityID               string                                                                                                        `json:"entityId,omitempty"`               // Entity Id
	FirstOccurredTime      *int                                                                                                          `json:"firstOccurredTime,omitempty"`      // First Occurred Time
	MostRecentOccurredTime *int                                                                                                          `json:"mostRecentOccurredTime,omitempty"` // Most Recent Occurred Time
	Status                 string                                                                                                        `json:"status,omitempty"`                 // Status
	IsGlobal               *bool                                                                                                         `json:"isGlobal,omitempty"`               // Is Global
	UpdatedBy              string                                                                                                        `json:"updatedBy,omitempty"`              // Updated By
	UpdatedTime            *int                                                                                                          `json:"updatedTime,omitempty"`            // Updated Time
	Notes                  *ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDV1ResponseNotes                  `json:"notes,omitempty"`                  // Notes
	SiteID                 *ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDV1ResponseSiteID                 `json:"siteId,omitempty"`                 // Site Id
	SiteHierarchyID        *ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDV1ResponseSiteHierarchyID        `json:"siteHierarchyId,omitempty"`        // Site Hierarchy Id
	SiteName               *ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDV1ResponseSiteName               `json:"siteName,omitempty"`               // Site Name
	SiteHierarchy          *ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDV1ResponseSiteHierarchy          `json:"siteHierarchy,omitempty"`          // Site Hierarchy
	SuggestedActions       *[]ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDV1ResponseSuggestedActions     `json:"suggestedActions,omitempty"`       //
	AdditionalAttributes   *[]ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDV1ResponseAdditionalAttributes `json:"additionalAttributes,omitempty"`   //
}
type ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDV1ResponseNotes interface{}
type ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDV1ResponseSiteID interface{}
type ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDV1ResponseSiteHierarchyID interface{}
type ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDV1ResponseSiteName interface{}
type ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDV1ResponseSiteHierarchy interface{}
type ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDV1ResponseSuggestedActions struct {
	Message string                                                                                                         `json:"message,omitempty"` // Message
	Steps   *[]ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDV1ResponseSuggestedActionsSteps `json:"steps,omitempty"`   // Steps
}
type ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDV1ResponseSuggestedActionsSteps interface{}
type ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDV1ResponseAdditionalAttributes struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseIssuesIgnoreTheGivenListOfIssuesV1 struct {
	Response *ResponseIssuesIgnoreTheGivenListOfIssuesV1Response `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  // Version
}
type ResponseIssuesIgnoreTheGivenListOfIssuesV1Response struct {
	SuccessfulIssueIDs []string `json:"successfulIssueIds,omitempty"` // Successful Issue Ids
	FailureIssueIDs    []string `json:"failureIssueIds,omitempty"`    // Failure Issue Ids
}
type ResponseIssuesResolveTheGivenListsOfIssuesV1 struct {
	Response *ResponseIssuesResolveTheGivenListsOfIssuesV1Response `json:"response,omitempty"` //
	Version  string                                                `json:"version,omitempty"`  // Version
}
type ResponseIssuesResolveTheGivenListsOfIssuesV1Response struct {
	SuccessfulIssueIDs []string `json:"successfulIssueIds,omitempty"` // Successful Issue Ids
	FailureIssueIDs    []string `json:"failureIssueIds,omitempty"`    // Failure Issue Ids
}
type ResponseIssuesUpdateTheGivenIssueByUpdatingSelectedFieldsV1 struct {
	Response *ResponseIssuesUpdateTheGivenIssueByUpdatingSelectedFieldsV1Response `json:"response,omitempty"` //
	Version  string                                                               `json:"version,omitempty"`  // Version
}
type ResponseIssuesUpdateTheGivenIssueByUpdatingSelectedFieldsV1Response struct {
	IssueID                string                                                                                     `json:"issueId,omitempty"`                // Issue Id
	Name                   string                                                                                     `json:"name,omitempty"`                   // Name
	Description            string                                                                                     `json:"description,omitempty"`            // Description
	Summary                string                                                                                     `json:"summary,omitempty"`                // Summary
	Priority               string                                                                                     `json:"priority,omitempty"`               // Priority
	Severity               string                                                                                     `json:"severity,omitempty"`               // Severity
	DeviceType             string                                                                                     `json:"deviceType,omitempty"`             // Device Type
	Category               string                                                                                     `json:"category,omitempty"`               // Category
	EntityType             string                                                                                     `json:"entityType,omitempty"`             // Entity Type
	EntityID               string                                                                                     `json:"entityId,omitempty"`               // Entity Id
	FirstOccurredTime      *int                                                                                       `json:"firstOccurredTime,omitempty"`      // First Occurred Time
	MostRecentOccurredTime *int                                                                                       `json:"mostRecentOccurredTime,omitempty"` // Most Recent Occurred Time
	Status                 string                                                                                     `json:"status,omitempty"`                 // Status
	IsGlobal               *bool                                                                                      `json:"isGlobal,omitempty"`               // Is Global
	UpdatedBy              string                                                                                     `json:"updatedBy,omitempty"`              // Updated By
	UpdatedTime            *int                                                                                       `json:"updatedTime,omitempty"`            // Updated Time
	Notes                  string                                                                                     `json:"notes,omitempty"`                  // Notes
	SiteID                 string                                                                                     `json:"siteId,omitempty"`                 // Site Id
	SiteHierarchyID        string                                                                                     `json:"siteHierarchyId,omitempty"`        // Site Hierarchy Id
	SiteName               string                                                                                     `json:"siteName,omitempty"`               // Site Name
	SiteHierarchy          string                                                                                     `json:"siteHierarchy,omitempty"`          // Site Hierarchy
	SuggestedActions       *[]ResponseIssuesUpdateTheGivenIssueByUpdatingSelectedFieldsV1ResponseSuggestedActions     `json:"suggestedActions,omitempty"`       //
	AdditionalAttributes   *[]ResponseIssuesUpdateTheGivenIssueByUpdatingSelectedFieldsV1ResponseAdditionalAttributes `json:"additionalAttributes,omitempty"`   //
}
type ResponseIssuesUpdateTheGivenIssueByUpdatingSelectedFieldsV1ResponseSuggestedActions struct {
	Message string `json:"message,omitempty"` // Message
}
type ResponseIssuesUpdateTheGivenIssueByUpdatingSelectedFieldsV1ResponseAdditionalAttributes struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseIssuesCreatesANewUserDefinedIssueDefinitionsV1 struct {
	Response *ResponseIssuesCreatesANewUserDefinedIssueDefinitionsV1Response `json:"response,omitempty"` //
}
type ResponseIssuesCreatesANewUserDefinedIssueDefinitionsV1Response struct {
	ID                    string                                                                 `json:"id,omitempty"`                    // Id
	Name                  string                                                                 `json:"name,omitempty"`                  // Name
	Description           string                                                                 `json:"description,omitempty"`           // Description
	ProfileID             string                                                                 `json:"profileId,omitempty"`             // Profile Id
	TriggerID             string                                                                 `json:"triggerId,omitempty"`             // Trigger Id
	Rules                 *[]ResponseIssuesCreatesANewUserDefinedIssueDefinitionsV1ResponseRules `json:"rules,omitempty"`                 //
	IsEnabled             *bool                                                                  `json:"isEnabled,omitempty"`             // Is Enabled
	Priority              string                                                                 `json:"priority,omitempty"`              // Priority
	IsDeletable           *bool                                                                  `json:"isDeletable,omitempty"`           // Is Deletable
	IsNotificationEnabled *bool                                                                  `json:"isNotificationEnabled,omitempty"` // Is Notification Enabled
	CreatedTime           *int                                                                   `json:"createdTime,omitempty"`           // Created Time
	LastUpdatedTime       *int                                                                   `json:"lastUpdatedTime,omitempty"`       // Last Updated Time
}
type ResponseIssuesCreatesANewUserDefinedIssueDefinitionsV1ResponseRules struct {
	Type              string `json:"type,omitempty"`              // Type
	Severity          *int   `json:"severity,omitempty"`          // Severity
	Facility          string `json:"facility,omitempty"`          // Facility
	Mnemonic          string `json:"mnemonic,omitempty"`          // Mnemonic
	Pattern           string `json:"pattern,omitempty"`           // Pattern
	Occurrences       *int   `json:"occurrences,omitempty"`       // Occurrences
	DurationInMinutes *int   `json:"durationInMinutes,omitempty"` // Duration In Minutes
}
type ResponseIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersV1 struct {
	Response *[]ResponseIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersV1Response `json:"response,omitempty"` //
	Page     *ResponseIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersV1Page       `json:"page,omitempty"`     //
	Version  string                                                                           `json:"version,omitempty"`  // Version
}
type ResponseIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersV1Response struct {
	ID                    string                                                                                `json:"id,omitempty"`                    // Id
	Name                  string                                                                                `json:"name,omitempty"`                  // Name
	Description           string                                                                                `json:"description,omitempty"`           // Description
	ProfileID             string                                                                                `json:"profileId,omitempty"`             // Profile Id
	TriggerID             string                                                                                `json:"triggerId,omitempty"`             // Trigger Id
	Rules                 *[]ResponseIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersV1ResponseRules `json:"rules,omitempty"`                 //
	IsEnabled             *bool                                                                                 `json:"isEnabled,omitempty"`             // Is Enabled
	Priority              string                                                                                `json:"priority,omitempty"`              // Priority
	IsDeletable           *bool                                                                                 `json:"isDeletable,omitempty"`           // Is Deletable
	IsNotificationEnabled *bool                                                                                 `json:"isNotificationEnabled,omitempty"` // Is Notification Enabled
	CreatedTime           *int                                                                                  `json:"createdTime,omitempty"`           // Created Time
	LastUpdatedTime       *int                                                                                  `json:"lastUpdatedTime,omitempty"`       // Last Updated Time
}
type ResponseIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersV1ResponseRules struct {
	Type              string `json:"type,omitempty"`              // Type
	Severity          *int   `json:"severity,omitempty"`          // Severity
	Facility          string `json:"facility,omitempty"`          // Facility
	Mnemonic          string `json:"mnemonic,omitempty"`          // Mnemonic
	Pattern           string `json:"pattern,omitempty"`           // Pattern
	Occurrences       *int   `json:"occurrences,omitempty"`       // Occurrences
	DurationInMinutes *int   `json:"durationInMinutes,omitempty"` // Duration In Minutes
}
type ResponseIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersV1Page struct {
	Limit  *int                                                                               `json:"limit,omitempty"`  // Limit
	Offset *int                                                                               `json:"offset,omitempty"` // Offset
	Count  *int                                                                               `json:"count,omitempty"`  // Count
	SortBy *[]ResponseIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersV1PageSortBy `json:"sortBy,omitempty"` //
}
type ResponseIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersV1PageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseIssuesGetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersV1 struct {
	Response *ResponseIssuesGetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersV1Response `json:"response,omitempty"` //
	Version  string                                                                                   `json:"version,omitempty"`  // Version
}
type ResponseIssuesGetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseIssuesGetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDV1 struct {
	Response *ResponseIssuesGetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDV1Response `json:"response,omitempty"` //
}
type ResponseIssuesGetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDV1Response struct {
	ID string `json:"id,omitempty"` // Id

	Name string `json:"name,omitempty"` // Name

	Description string `json:"description,omitempty"` // Description

	ProfileID string `json:"profileId,omitempty"` // Profile Id

	TriggerID string `json:"triggerId,omitempty"` // Trigger Id

	Rules *[]ResponseIssuesGetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDV1ResponseRules `json:"rules,omitempty"` //

	IsEnabled *bool `json:"isEnabled,omitempty"` // Is Enabled

	Priority string `json:"priority,omitempty"` // Priority

	IsDeletable *bool `json:"isDeletable,omitempty"` // Is Deletable

	IsNotificationEnabled *bool `json:"isNotificationEnabled,omitempty"` // Is Notification Enabled

	CreatedTime *int `json:"createdTime,omitempty"` // Created Time

	LastUpdatedTime *int `json:"lastUpdatedTime,omitempty"` // Last Updated Time
}
type ResponseIssuesGetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDV1ResponseRules struct {
	Type string `json:"type,omitempty"` // Type

	Severity *int `json:"severity,omitempty"` // Severity

	Facility string `json:"facility,omitempty"` // Facility

	Mnemonic string `json:"mnemonic,omitempty"` // Mnemonic

	Pattern string `json:"pattern,omitempty"` // Pattern

	Occurrences *int `json:"occurrences,omitempty"` // Occurrences

	DurationInMinutes *int `json:"durationInMinutes,omitempty"` // Duration In Minutes
}
type ResponseIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIDV1 struct {
	Response *ResponseIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIDV1Response `json:"response,omitempty"` //
}
type ResponseIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIDV1Response struct {
	ID                    string                                                                                     `json:"id,omitempty"`                    // Id
	Name                  string                                                                                     `json:"name,omitempty"`                  // Name
	Description           string                                                                                     `json:"description,omitempty"`           // Description
	ProfileID             string                                                                                     `json:"profileId,omitempty"`             // Profile Id
	TriggerID             string                                                                                     `json:"triggerId,omitempty"`             // Trigger Id
	Rules                 *[]ResponseIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIDV1ResponseRules `json:"rules,omitempty"`                 //
	IsEnabled             *bool                                                                                      `json:"isEnabled,omitempty"`             // Is Enabled
	Priority              string                                                                                     `json:"priority,omitempty"`              // Priority
	IsDeletable           *bool                                                                                      `json:"isDeletable,omitempty"`           // Is Deletable
	IsNotificationEnabled *bool                                                                                      `json:"isNotificationEnabled,omitempty"` // Is Notification Enabled
	CreatedTime           *int                                                                                       `json:"createdTime,omitempty"`           // Created Time
	LastUpdatedTime       *int                                                                                       `json:"lastUpdatedTime,omitempty"`       // Last Updated Time
}
type ResponseIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIDV1ResponseRules struct {
	Type              string `json:"type,omitempty"`              // Type
	Severity          *int   `json:"severity,omitempty"`          // Severity
	Facility          string `json:"facility,omitempty"`          // Facility
	Mnemonic          string `json:"mnemonic,omitempty"`          // Mnemonic
	Pattern           string `json:"pattern,omitempty"`           // Pattern
	Occurrences       *int   `json:"occurrences,omitempty"`       // Occurrences
	DurationInMinutes *int   `json:"durationInMinutes,omitempty"` // Duration In Minutes
}
type ResponseIssuesExecuteSuggestedActionsCommandsV1 []ResponseItemIssuesExecuteSuggestedActionsCommandsV1 // Array of ResponseIssuesExecuteSuggestedActionsCommandsV1
type ResponseItemIssuesExecuteSuggestedActionsCommandsV1 struct {
	ActionInfo       string                                                            `json:"actionInfo,omitempty"`       // Actions Info
	StepsCount       *int                                                              `json:"stepsCount,omitempty"`       // Steps Count
	EntityID         string                                                            `json:"entityId,omitempty"`         // Entity Id
	Hostname         string                                                            `json:"hostname,omitempty"`         // Hostname
	StepsDescription string                                                            `json:"stepsDescription,omitempty"` // Steps Description
	Command          string                                                            `json:"command,omitempty"`          // Command
	CommandOutput    *ResponseItemIssuesExecuteSuggestedActionsCommandsV1CommandOutput `json:"commandOutput,omitempty"`    // Command Output
}
type ResponseItemIssuesExecuteSuggestedActionsCommandsV1CommandOutput interface{}
type ResponseIssuesGetIssueEnrichmentDetailsV1 struct {
	IssueDetails *ResponseIssuesGetIssueEnrichmentDetailsV1IssueDetails `json:"issueDetails,omitempty"` //
}
type ResponseIssuesGetIssueEnrichmentDetailsV1IssueDetails struct {
	Issue *[]ResponseIssuesGetIssueEnrichmentDetailsV1IssueDetailsIssue `json:"issue,omitempty"` //
}
type ResponseIssuesGetIssueEnrichmentDetailsV1IssueDetailsIssue struct {
	IssueID          string                                                                        `json:"issueId,omitempty"`          // Issue Id
	IssueSource      string                                                                        `json:"issueSource,omitempty"`      // Issue Source
	IssueCategory    string                                                                        `json:"issueCategory,omitempty"`    // Issue Category
	IssueName        string                                                                        `json:"issueName,omitempty"`        // Issue Name
	IssueDescription string                                                                        `json:"issueDescription,omitempty"` // Issue Description
	IssueEntity      string                                                                        `json:"issueEntity,omitempty"`      // Issue Entity
	IssueEntityValue string                                                                        `json:"issueEntityValue,omitempty"` // Issue Entity Value
	IssueSeverity    string                                                                        `json:"issueSeverity,omitempty"`    // Issue Severity
	IssuePriority    string                                                                        `json:"issuePriority,omitempty"`    // Issue Priority
	IssueSummary     string                                                                        `json:"issueSummary,omitempty"`     // Issue Summary
	IssueTimestamp   *int                                                                          `json:"issueTimestamp,omitempty"`   // Issue Timestamp
	SuggestedActions *[]ResponseIssuesGetIssueEnrichmentDetailsV1IssueDetailsIssueSuggestedActions `json:"suggestedActions,omitempty"` //
	ImpactedHosts    *[]ResponseIssuesGetIssueEnrichmentDetailsV1IssueDetailsIssueImpactedHosts    `json:"impactedHosts,omitempty"`    // Impacted Hosts
}
type ResponseIssuesGetIssueEnrichmentDetailsV1IssueDetailsIssueSuggestedActions struct {
	Message string                                                                             `json:"message,omitempty"` // Message
	Steps   *[]ResponseIssuesGetIssueEnrichmentDetailsV1IssueDetailsIssueSuggestedActionsSteps `json:"steps,omitempty"`   // Steps
}
type ResponseIssuesGetIssueEnrichmentDetailsV1IssueDetailsIssueSuggestedActionsSteps interface{}
type ResponseIssuesGetIssueEnrichmentDetailsV1IssueDetailsIssueImpactedHosts interface{}
type ResponseIssuesIssuesV1 struct {
	Version    string                            `json:"version,omitempty"`    // Response body's schema version string
	TotalCount string                            `json:"totalCount,omitempty"` // Total number of issues in the query time window
	Response   *[]ResponseIssuesIssuesV1Response `json:"response,omitempty"`   //
}
type ResponseIssuesIssuesV1Response struct {
	IssueID             string `json:"issueId,omitempty"`               // The issue's unique identifier
	Name                string `json:"name,omitempty"`                  // The issue's display name
	SiteID              string `json:"siteId,omitempty"`                // The site UUID where the issue occurred
	DeviceID            string `json:"deviceId,omitempty"`              // The device UUID where the issue occurred
	DeviceRole          string `json:"deviceRole,omitempty"`            // The device role
	AiDriven            string `json:"aiDriven,omitempty"`              // Whether the issue is AI driven ('Yes' or 'No')
	ClientMac           string `json:"clientMac,omitempty"`             // The client MAC address related to this issue
	IssueOccurenceCount *int   `json:"issue_occurence_count,omitempty"` // Total number of instances of this issue in the query time window
	Status              string `json:"status,omitempty"`                // The status of the issue
	Priority            string `json:"priority,omitempty"`              // Priority setting of the issue
	Category            string `json:"category,omitempty"`              // Category of the issue
	LastOccurenceTime   *int   `json:"last_occurence_time,omitempty"`   // The UTC timestamp of last occurence of this issue
}
type ResponseIssuesReturnsAllIssueTriggerDefinitionsForGivenFiltersV1 struct {
	Response *[]ResponseIssuesReturnsAllIssueTriggerDefinitionsForGivenFiltersV1Response `json:"response,omitempty"` //
}
type ResponseIssuesReturnsAllIssueTriggerDefinitionsForGivenFiltersV1Response struct {
	ID                           string   `json:"id,omitempty"`                           // Id
	Name                         string   `json:"name,omitempty"`                         // Name
	DisplayName                  string   `json:"displayName,omitempty"`                  // Display Name
	Description                  string   `json:"description,omitempty"`                  // Description
	Priority                     string   `json:"priority,omitempty"`                     // Priority
	DefaultPriority              string   `json:"defaultPriority,omitempty"`              // Default Priority
	DeviceType                   string   `json:"deviceType,omitempty"`                   // Device Type
	IssueEnabled                 *bool    `json:"issueEnabled,omitempty"`                 // Issue Enabled
	ProfileID                    string   `json:"profileId,omitempty"`                    // Profile Id
	DefinitionStatus             string   `json:"definitionStatus,omitempty"`             // Definition Status
	CategoryName                 string   `json:"categoryName,omitempty"`                 // Category Name
	SynchronizeToHealthThreshold *bool    `json:"synchronizeToHealthThreshold,omitempty"` // Synchronize To Health Threshold
	ThresholdValue               *float64 `json:"thresholdValue,omitempty"`               // Threshold Value
	LastModified                 string   `json:"lastModified,omitempty"`                 // Last Modified
}
type ResponseIssuesGetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1 struct {
	Response *ResponseIssuesGetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1Response `json:"response,omitempty"` //
	Version  string                                                                                    `json:"version,omitempty"`  // Version
}
type ResponseIssuesGetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseIssuesGetIssueTriggerDefinitionForGivenIDV1 struct {
	Response *ResponseIssuesGetIssueTriggerDefinitionForGivenIDV1Response `json:"response,omitempty"` //
	Version  string                                                       `json:"version,omitempty"`  // Version
}
type ResponseIssuesGetIssueTriggerDefinitionForGivenIDV1Response struct {
	ID                           string   `json:"id,omitempty"`                           // Id
	Name                         string   `json:"name,omitempty"`                         // Name
	DisplayName                  string   `json:"displayName,omitempty"`                  // Display Name
	Description                  string   `json:"description,omitempty"`                  // Description
	Priority                     string   `json:"priority,omitempty"`                     // Priority
	DefaultPriority              string   `json:"defaultPriority,omitempty"`              // Default Priority
	DeviceType                   string   `json:"deviceType,omitempty"`                   // Device Type
	IssueEnabled                 *bool    `json:"issueEnabled,omitempty"`                 // Issue Enabled
	ProfileID                    string   `json:"profileId,omitempty"`                    // Profile Id
	DefinitionStatus             string   `json:"definitionStatus,omitempty"`             // Definition Status
	CategoryName                 string   `json:"categoryName,omitempty"`                 // Category Name
	SynchronizeToHealthThreshold *bool    `json:"synchronizeToHealthThreshold,omitempty"` // Synchronize To Health Threshold
	ThresholdValue               *float64 `json:"thresholdValue,omitempty"`               // Threshold Value
	LastModified                 string   `json:"lastModified,omitempty"`                 // Last Modified
}
type ResponseIssuesIssueTriggerDefinitionUpdateV1 struct {
	Response *ResponseIssuesIssueTriggerDefinitionUpdateV1Response `json:"response,omitempty"` //
	Version  string                                                `json:"version,omitempty"`  // Version
}
type ResponseIssuesIssueTriggerDefinitionUpdateV1Response struct {
	ID string `json:"id,omitempty"` // Id

	Name string `json:"name,omitempty"` // Name

	DisplayName string `json:"displayName,omitempty"` // Display Name

	Description string `json:"description,omitempty"` // Description

	Priority string `json:"priority,omitempty"` // Priority

	DefaultPriority string `json:"defaultPriority,omitempty"` // Default Priority

	DeviceType string `json:"deviceType,omitempty"` // Device Type

	IssueEnabled *bool `json:"issueEnabled,omitempty"` // Issue Enabled

	ProfileID string `json:"profileId,omitempty"` // Profile Id

	DefinitionStatus string `json:"definitionStatus,omitempty"` // Definition Status

	CategoryName string `json:"categoryName,omitempty"` // Category Name

	SynchronizeToHealthThreshold *bool `json:"synchronizeToHealthThreshold,omitempty"` // Synchronize To Health Threshold

	ThresholdValue *float64 `json:"thresholdValue,omitempty"` // Threshold Value

	LastModified string `json:"lastModified,omitempty"` // Last Modified
}
type RequestIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1 struct {
	StartTime *int                                                               `json:"startTime,omitempty"` // Start Time
	EndTime   *int                                                               `json:"endTime,omitempty"`   // End Time
	Filters   *[]RequestIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1Filters `json:"filters,omitempty"`   //
}
type RequestIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1Filters struct {
	Key             string                                                                    `json:"key,omitempty"`             // Key
	Operator        string                                                                    `json:"operator,omitempty"`        // Operator
	Value           string                                                                    `json:"value,omitempty"`           // Value
	LogicalOperator string                                                                    `json:"logicalOperator,omitempty"` // Logical Operator
	Filters         *[]RequestIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1FiltersFilters `json:"filters,omitempty"`         //
}
type RequestIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1FiltersFilters struct {
	Key      string `json:"key,omitempty"`      // Key
	Operator string `json:"operator,omitempty"` // Operator
	Value    string `json:"value,omitempty"`    // Value
}
type RequestIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersV1 struct {
	StartTime *int                                                                   `json:"startTime,omitempty"` // Start Time
	EndTime   *int                                                                   `json:"endTime,omitempty"`   // End Time
	Filters   *[]RequestIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersV1Filters `json:"filters,omitempty"`   //
}
type RequestIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersV1Filters struct {
	Key             string                                                                        `json:"key,omitempty"`             // Key
	Operator        string                                                                        `json:"operator,omitempty"`        // Operator
	Value           string                                                                        `json:"value,omitempty"`           // Value
	LogicalOperator string                                                                        `json:"logicalOperator,omitempty"` // Logical Operator
	Filters         *[]RequestIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersV1FiltersFilters `json:"filters,omitempty"`         //
}
type RequestIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersV1FiltersFilters struct {
	Key      string `json:"key,omitempty"`      // Key
	Operator string `json:"operator,omitempty"` // Operator
	Value    string `json:"value,omitempty"`    // Value
}
type RequestIssuesGetSummaryAnalyticsDataOfIssuesV1 struct {
	StartTime           *int                                                                 `json:"startTime,omitempty"`           // Start Time
	EndTime             *int                                                                 `json:"endTime,omitempty"`             // End Time
	Filters             *[]RequestIssuesGetSummaryAnalyticsDataOfIssuesV1Filters             `json:"filters,omitempty"`             //
	GroupBy             []string                                                             `json:"groupBy,omitempty"`             // Group By
	Attributes          []string                                                             `json:"attributes,omitempty"`          // Attributes
	AggregateAttributes *[]RequestIssuesGetSummaryAnalyticsDataOfIssuesV1AggregateAttributes `json:"aggregateAttributes,omitempty"` //
	Page                *RequestIssuesGetSummaryAnalyticsDataOfIssuesV1Page                  `json:"page,omitempty"`                //
}
type RequestIssuesGetSummaryAnalyticsDataOfIssuesV1Filters struct {
	Key             string                                                          `json:"key,omitempty"`             // Key
	Operator        string                                                          `json:"operator,omitempty"`        // Operator
	Value           string                                                          `json:"value,omitempty"`           // Value
	LogicalOperator string                                                          `json:"logicalOperator,omitempty"` // Logical Operator
	Filters         *[]RequestIssuesGetSummaryAnalyticsDataOfIssuesV1FiltersFilters `json:"filters,omitempty"`         //
}
type RequestIssuesGetSummaryAnalyticsDataOfIssuesV1FiltersFilters struct {
	Key      string `json:"key,omitempty"`      // Key
	Operator string `json:"operator,omitempty"` // Operator
	Value    string `json:"value,omitempty"`    // Value
}
type RequestIssuesGetSummaryAnalyticsDataOfIssuesV1AggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
}
type RequestIssuesGetSummaryAnalyticsDataOfIssuesV1Page struct {
	Limit  *int                                                        `json:"limit,omitempty"`  // Limit
	Offset *int                                                        `json:"offset,omitempty"` // Offset
	SortBy *[]RequestIssuesGetSummaryAnalyticsDataOfIssuesV1PageSortBy `json:"sortBy,omitempty"` //
}
type RequestIssuesGetSummaryAnalyticsDataOfIssuesV1PageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type RequestIssuesGetTopNAnalyticsDataOfIssuesV1 struct {
	StartTime           *int                                                              `json:"startTime,omitempty"`           // Start Time
	EndTime             *int                                                              `json:"endTime,omitempty"`             // End Time
	TopN                *int                                                              `json:"topN,omitempty"`                // Top N
	Filters             *[]RequestIssuesGetTopNAnalyticsDataOfIssuesV1Filters             `json:"filters,omitempty"`             //
	GroupBy             []string                                                          `json:"groupBy,omitempty"`             // Group By
	Attributes          []string                                                          `json:"attributes,omitempty"`          // Attributes
	AggregateAttributes *[]RequestIssuesGetTopNAnalyticsDataOfIssuesV1AggregateAttributes `json:"aggregateAttributes,omitempty"` //
	Page                *RequestIssuesGetTopNAnalyticsDataOfIssuesV1Page                  `json:"page,omitempty"`                //
}
type RequestIssuesGetTopNAnalyticsDataOfIssuesV1Filters struct {
	Key             string                                                       `json:"key,omitempty"`             // Key
	Operator        string                                                       `json:"operator,omitempty"`        // Operator
	Value           string                                                       `json:"value,omitempty"`           // Value
	LogicalOperator string                                                       `json:"logicalOperator,omitempty"` // Logical Operator
	Filters         *[]RequestIssuesGetTopNAnalyticsDataOfIssuesV1FiltersFilters `json:"filters,omitempty"`         //
}
type RequestIssuesGetTopNAnalyticsDataOfIssuesV1FiltersFilters struct {
	Key      string `json:"key,omitempty"`      // Key
	Operator string `json:"operator,omitempty"` // Operator
	Value    string `json:"value,omitempty"`    // Value
}
type RequestIssuesGetTopNAnalyticsDataOfIssuesV1AggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
}
type RequestIssuesGetTopNAnalyticsDataOfIssuesV1Page struct {
	Limit  *int                                                     `json:"limit,omitempty"`  // Limit
	Offset *int                                                     `json:"offset,omitempty"` // Offset
	SortBy *[]RequestIssuesGetTopNAnalyticsDataOfIssuesV1PageSortBy `json:"sortBy,omitempty"` //
}
type RequestIssuesGetTopNAnalyticsDataOfIssuesV1PageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type RequestIssuesGetTrendAnalyticsDataOfIssuesV1 struct {
	StartTime           *int                                                               `json:"startTime,omitempty"`           // Start Time
	EndTime             *int                                                               `json:"endTime,omitempty"`             // End Time
	TrendInterval       string                                                             `json:"trendInterval,omitempty"`       // Trend Interval
	Filters             *[]RequestIssuesGetTrendAnalyticsDataOfIssuesV1Filters             `json:"filters,omitempty"`             //
	GroupBy             []string                                                           `json:"groupBy,omitempty"`             // Group By
	Attributes          []string                                                           `json:"attributes,omitempty"`          // Attributes
	AggregateAttributes *[]RequestIssuesGetTrendAnalyticsDataOfIssuesV1AggregateAttributes `json:"aggregateAttributes,omitempty"` //
	Page                *RequestIssuesGetTrendAnalyticsDataOfIssuesV1Page                  `json:"page,omitempty"`                //
}
type RequestIssuesGetTrendAnalyticsDataOfIssuesV1Filters struct {
	Key      string `json:"key,omitempty"`      // Key
	Value    string `json:"value,omitempty"`    // Value
	Operator string `json:"operator,omitempty"` // Operator
}
type RequestIssuesGetTrendAnalyticsDataOfIssuesV1AggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
}
type RequestIssuesGetTrendAnalyticsDataOfIssuesV1Page struct {
	Limit          *int   `json:"limit,omitempty"`          // Limit
	Offset         *int   `json:"offset,omitempty"`         // Offset
	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type RequestIssuesIgnoreTheGivenListOfIssuesV1 struct {
	IssueIDs []string `json:"issueIds,omitempty"` // Issue Ids
}
type RequestIssuesResolveTheGivenListsOfIssuesV1 struct {
	IssueIDs []string `json:"issueIds,omitempty"` // Issue Ids
}
type RequestIssuesUpdateTheGivenIssueByUpdatingSelectedFieldsV1 struct {
	Notes string `json:"notes,omitempty"` // Notes
}
type RequestIssuesCreatesANewUserDefinedIssueDefinitionsV1 struct {
	Name                  string                                                        `json:"name,omitempty"`                  // Name
	Description           string                                                        `json:"description,omitempty"`           // Description
	Rules                 *[]RequestIssuesCreatesANewUserDefinedIssueDefinitionsV1Rules `json:"rules,omitempty"`                 //
	IsEnabled             *bool                                                         `json:"isEnabled,omitempty"`             // Is Enabled
	Priority              string                                                        `json:"priority,omitempty"`              // Priority
	IsNotificationEnabled *bool                                                         `json:"isNotificationEnabled,omitempty"` // Is Notification Enabled
}
type RequestIssuesCreatesANewUserDefinedIssueDefinitionsV1Rules struct {
	Severity          *int   `json:"severity,omitempty"`          // Severity
	Facility          string `json:"facility,omitempty"`          // Facility
	Mnemonic          string `json:"mnemonic,omitempty"`          // Mnemonic
	Pattern           string `json:"pattern,omitempty"`           // Pattern
	Occurrences       *int   `json:"occurrences,omitempty"`       // Occurrences
	DurationInMinutes *int   `json:"durationInMinutes,omitempty"` // Duration In Minutes
}
type RequestIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIDV1 struct {
	Name                  string                                                                            `json:"name,omitempty"`                  // Name
	Description           string                                                                            `json:"description,omitempty"`           // Description
	Rules                 *[]RequestIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIDV1Rules `json:"rules,omitempty"`                 //
	IsEnabled             *bool                                                                             `json:"isEnabled,omitempty"`             // Is Enabled
	Priority              string                                                                            `json:"priority,omitempty"`              // Priority
	IsNotificationEnabled *bool                                                                             `json:"isNotificationEnabled,omitempty"` // Is Notification Enabled
}
type RequestIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIDV1Rules struct {
	Severity          *int   `json:"severity,omitempty"`          // Severity
	Facility          string `json:"facility,omitempty"`          // Facility
	Mnemonic          string `json:"mnemonic,omitempty"`          // Mnemonic
	Pattern           string `json:"pattern,omitempty"`           // Pattern
	Occurrences       *int   `json:"occurrences,omitempty"`       // Occurrences
	DurationInMinutes *int   `json:"durationInMinutes,omitempty"` // Duration In Minutes
}
type RequestIssuesExecuteSuggestedActionsCommandsV1 struct {
	EntityType  string `json:"entity_type,omitempty"`  // Commands provided as part of the suggested actions for an issue can be executed based on issue id. The value here must be issue_id
	EntityValue string `json:"entity_value,omitempty"` // Contains the actual value for the entity type that has been defined
}
type RequestIssuesIssueTriggerDefinitionUpdateV1 struct {
	SynchronizeToHealthThreshold *bool    `json:"synchronizeToHealthThreshold,omitempty"` // Synchronize To Health Threshold
	Priority                     string   `json:"priority,omitempty"`                     // Priority
	IssueEnabled                 *bool    `json:"issueEnabled,omitempty"`                 // Issue Enabled
	ThresholdValue               *float64 `json:"thresholdValue,omitempty"`               // Threshold Value
}

//GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1 Get the details of issues for given set of filters - a991-6985-476b-b271
/* Returns all details of each issue along with suggested actions for given set of filters specified in query parameters. If there is no start and/or end time, then end time will be defaulted to current time and start time will be defaulted to 24-hours ago from end time. All string type query parameters support wildcard search (using *). For example: siteHierarchy=Global/San Jose/* returns issues under all sites whole siteHierarchy starts with "Global/San Jose/". https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-IssuesList-1.0.0-resolved.yaml


@param GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1HeaderParams Custom header parameters
@param GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-the-details-of-issues-for-given-set-of-filters-know-your-network
*/
func (s *IssuesService) GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1(GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1HeaderParams *GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1HeaderParams, GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1QueryParams *GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1QueryParams) (*ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceIssues"

	queryString, _ := query.Values(GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1HeaderParams != nil {

		if GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1HeaderParams.AcceptLanguage != "" {
			clientRequest = clientRequest.SetHeader("Accept-Language", GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1HeaderParams.AcceptLanguage)
		}

		if GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1(GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1HeaderParams, GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1")
	}

	result := response.Result().(*ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1)
	return result, response, err

}

//GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkV1 Get the total number of issues for given set of filters - 049b-c87d-456a-a69b
/* Returns the total number issues for given set of filters. If there is no start and/or end time, then end time will be defaulted to current time and start time will be defaulted to 24-hours ago from end time. https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-IssuesList-1.0.0-resolved.yaml


@param GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkV1HeaderParams Custom header parameters
@param GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-the-total-number-of-issues-for-given-set-of-filters-know-your-network
*/
func (s *IssuesService) GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkV1(GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkV1HeaderParams *GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkV1HeaderParams, GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkV1QueryParams *GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkV1QueryParams) (*ResponseIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkV1, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceIssues/count"

	queryString, _ := query.Values(GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkV1HeaderParams != nil {

		if GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkV1(GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkV1HeaderParams, GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkV1")
	}

	result := response.Result().(*ResponseIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkV1)
	return result, response, err

}

//GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDV1 Get all the details and suggested actions of an issue for the given issue id - 82ae-1acd-4b6a-ab00
/* Returns all the details and suggested actions of an issue for the given issue id. https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-IssuesList-1.0.0-resolved.yaml


@param id id path parameter. The issue Uuid

@param GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIdV1HeaderParams Custom header parameters
@param GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIdV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-all-the-details-and-suggested-actions-of-an-issue-for-the-given-issue-id
*/
func (s *IssuesService) GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDV1(id string, GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIdV1HeaderParams *GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDV1HeaderParams, GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIdV1QueryParams *GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDV1QueryParams) (*ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDV1, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceIssues/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIdV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIdV1HeaderParams != nil {

		if GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIdV1HeaderParams.AcceptLanguage != "" {
			clientRequest = clientRequest.SetHeader("Accept-Language", GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIdV1HeaderParams.AcceptLanguage)
		}

		if GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIdV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIdV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDV1(id, GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIdV1HeaderParams, GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIdV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIdV1")
	}

	result := response.Result().(*ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDV1)
	return result, response, err

}

//GetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersV1 Get all the custom issue definitions based on the given filters. - 1bb9-bb87-4efa-afd2
/* Retrieve the existing syslog-based custom issue definitions. The supported filters are id, name, profileId,  definition enable status, priority, severity, facility and mnemonic. The issue definition configurations may vary across profiles, hence specifying the profile Id in the query parameter is important and the default profile is global.

  The assurance profile definitions can be obtain via the API endpoint: /api/v1/siteprofile?namespace=assurance. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceUserDefinedIssueAPIs-1.0.0-resolved.yaml


@param GetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-all-the-custom-issue-definitions-based-on-the-given-filters
*/
func (s *IssuesService) GetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersV1(GetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersV1QueryParams *GetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersV1QueryParams) (*ResponseIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/customIssueDefinitions"

	queryString, _ := query.Values(GetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersV1(GetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersV1")
	}

	result := response.Result().(*ResponseIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersV1)
	return result, response, err

}

//GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersV1 Get the total custom issue definitions count based on the provided filters. - 9b91-2a4a-4d1a-9595
/* Get the total number of Custom issue definitions count based on the provided filters. The supported filters are id, name, profileId and definition enable status, severity, facility and mnemonic. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceUserDefinedIssueAPIs-1.0.0-resolved.yaml


@param GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersV1HeaderParams Custom header parameters
@param GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-the-total-custom-issue-definitions-count-based-on-the-provided-filters
*/
func (s *IssuesService) GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersV1(GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersV1HeaderParams *GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersV1HeaderParams, GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersV1QueryParams *GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersV1QueryParams) (*ResponseIssuesGetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/customIssueDefinitions/count"

	queryString, _ := query.Values(GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersV1HeaderParams != nil {

		if GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseIssuesGetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersV1(GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersV1HeaderParams, GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersV1")
	}

	result := response.Result().(*ResponseIssuesGetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersV1)
	return result, response, err

}

//GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDV1 Get the custom issue definition for the given custom issue definition Id. - d39f-a9d8-44b8-880d
/* Get the custom issue definition for the given custom issue definition Id. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceUserDefinedIssueAPIs-1.0.0-resolved.yaml


@param id id path parameter. Get the custom issue definition for the given custom issue definition Id.

@param GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIdV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-the-custom-issue-definition-for-the-given-custom-issue-definition-id
*/
func (s *IssuesService) GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDV1(id string, GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIdV1HeaderParams *GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDV1HeaderParams) (*ResponseIssuesGetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/customIssueDefinitions/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIdV1HeaderParams != nil {

		if GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIdV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIdV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetResult(&ResponseIssuesGetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDV1(id, GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIdV1HeaderParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIdV1")
	}

	result := response.Result().(*ResponseIssuesGetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDV1)
	return result, response, err

}

//GetIssueEnrichmentDetailsV1 Get Issue Enrichment Details - 8684-39bb-4e89-a6e4
/* Enriches a given network issue context (an issue id or end user’s Mac Address) with details about the issue(s), impacted hosts and suggested actions for remediation


@param GetIssueEnrichmentDetailsV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-issue-enrichment-details
*/
func (s *IssuesService) GetIssueEnrichmentDetailsV1(GetIssueEnrichmentDetailsV1HeaderParams *GetIssueEnrichmentDetailsV1HeaderParams) (*ResponseIssuesGetIssueEnrichmentDetailsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/issue-enrichment-details"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetIssueEnrichmentDetailsV1HeaderParams != nil {

		if GetIssueEnrichmentDetailsV1HeaderParams.EntityType != "" {
			clientRequest = clientRequest.SetHeader("entity_type", GetIssueEnrichmentDetailsV1HeaderParams.EntityType)
		}

		if GetIssueEnrichmentDetailsV1HeaderParams.EntityValue != "" {
			clientRequest = clientRequest.SetHeader("entity_value", GetIssueEnrichmentDetailsV1HeaderParams.EntityValue)
		}

		if GetIssueEnrichmentDetailsV1HeaderParams.Persistbapioutput != "" {
			clientRequest = clientRequest.SetHeader("__persistbapioutput", GetIssueEnrichmentDetailsV1HeaderParams.Persistbapioutput)
		}

	}

	response, err = clientRequest.
		SetResult(&ResponseIssuesGetIssueEnrichmentDetailsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetIssueEnrichmentDetailsV1(GetIssueEnrichmentDetailsV1HeaderParams)
		}
		return nil, response, fmt.Errorf("error with operation GetIssueEnrichmentDetailsV1")
	}

	result := response.Result().(*ResponseIssuesGetIssueEnrichmentDetailsV1)
	return result, response, err

}

//IssuesV1 Issues - ecb6-7807-47c9-bc59
/* Intent API to get a list of global issues, issues for a specific device, or issue for a specific client device's MAC address.


@param IssuesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!issues
*/
func (s *IssuesService) IssuesV1(IssuesV1QueryParams *IssuesV1QueryParams) (*ResponseIssuesIssuesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/issues"

	queryString, _ := query.Values(IssuesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseIssuesIssuesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.IssuesV1(IssuesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation IssuesV1")
	}

	result := response.Result().(*ResponseIssuesIssuesV1)
	return result, response, err

}

//ReturnsAllIssueTriggerDefinitionsForGivenFiltersV1 Returns all issue trigger definitions for given filters. - 199e-880b-4dc9-95c3
/* Get all system issue defintions. The supported filters are id, name, profileId and definition enable status. An issue trigger definition can be different across the profile and device type. So, `profileId` and `deviceType` in the query param is important and default is global profile and all device type. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-issueAndHealthDefinitions-1.0.0-resolved.yaml


@param ReturnsAllIssueTriggerDefinitionsForGivenFiltersV1HeaderParams Custom header parameters
@param ReturnsAllIssueTriggerDefinitionsForGivenFiltersV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!returns-all-issue-trigger-definitions-for-given-filters
*/
func (s *IssuesService) ReturnsAllIssueTriggerDefinitionsForGivenFiltersV1(ReturnsAllIssueTriggerDefinitionsForGivenFiltersV1HeaderParams *ReturnsAllIssueTriggerDefinitionsForGivenFiltersV1HeaderParams, ReturnsAllIssueTriggerDefinitionsForGivenFiltersV1QueryParams *ReturnsAllIssueTriggerDefinitionsForGivenFiltersV1QueryParams) (*ResponseIssuesReturnsAllIssueTriggerDefinitionsForGivenFiltersV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/systemIssueDefinitions"

	queryString, _ := query.Values(ReturnsAllIssueTriggerDefinitionsForGivenFiltersV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if ReturnsAllIssueTriggerDefinitionsForGivenFiltersV1HeaderParams != nil {

		if ReturnsAllIssueTriggerDefinitionsForGivenFiltersV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", ReturnsAllIssueTriggerDefinitionsForGivenFiltersV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseIssuesReturnsAllIssueTriggerDefinitionsForGivenFiltersV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReturnsAllIssueTriggerDefinitionsForGivenFiltersV1(ReturnsAllIssueTriggerDefinitionsForGivenFiltersV1HeaderParams, ReturnsAllIssueTriggerDefinitionsForGivenFiltersV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReturnsAllIssueTriggerDefinitionsForGivenFiltersV1")
	}

	result := response.Result().(*ResponseIssuesReturnsAllIssueTriggerDefinitionsForGivenFiltersV1)
	return result, response, err

}

//GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1 Get the count of system defined issue definitions based on provided filters. - a7b5-4a48-4b5b-a680
/* Get the count of system defined issue definitions based on provided filters. Supported filters are id, name, profileId and definition enable status. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-issueAndHealthDefinitions-1.0.0-resolved.yaml


@param GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1HeaderParams Custom header parameters
@param GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-the-count-of-system-defined-issue-definitions-based-on-provided-filters
*/
func (s *IssuesService) GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1(GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1HeaderParams *GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1HeaderParams, GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1QueryParams *GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1QueryParams) (*ResponseIssuesGetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/systemIssueDefinitions/count"

	queryString, _ := query.Values(GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1HeaderParams != nil {

		if GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseIssuesGetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1(GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1HeaderParams, GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1")
	}

	result := response.Result().(*ResponseIssuesGetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1)
	return result, response, err

}

//GetIssueTriggerDefinitionForGivenIDV1 Get issue trigger definition for given id. - 71a4-aa5c-400a-a129
/* Get system issue defintion for the given id. Definition includes all properties from IssueTriggerDefinition schema by default. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-issueAndHealthDefinitions-1.0.0-resolved.yaml


@param id id path parameter. Issue trigger definition id.

@param GetIssueTriggerDefinitionForGivenIdV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-issue-trigger-definition-for-given-id
*/
func (s *IssuesService) GetIssueTriggerDefinitionForGivenIDV1(id string, GetIssueTriggerDefinitionForGivenIdV1HeaderParams *GetIssueTriggerDefinitionForGivenIDV1HeaderParams) (*ResponseIssuesGetIssueTriggerDefinitionForGivenIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/systemIssueDefinitions/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetIssueTriggerDefinitionForGivenIdV1HeaderParams != nil {

		if GetIssueTriggerDefinitionForGivenIdV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetIssueTriggerDefinitionForGivenIdV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetResult(&ResponseIssuesGetIssueTriggerDefinitionForGivenIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetIssueTriggerDefinitionForGivenIDV1(id, GetIssueTriggerDefinitionForGivenIdV1HeaderParams)
		}
		return nil, response, fmt.Errorf("error with operation GetIssueTriggerDefinitionForGivenIdV1")
	}

	result := response.Result().(*ResponseIssuesGetIssueTriggerDefinitionForGivenIDV1)
	return result, response, err

}

//GetTheDetailsOfIssuesForGivenSetOfFiltersV1 Get the details of issues for given set of filters - 82ad-186f-4848-a3dd
/* Returns all details of each issue along with suggested actions for given set of filters specified in request body. If there is no start and/or end time, then end time will be defaulted to current time and start time will be defaulted to 24-hours ago from end time. https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-IssuesList-1.0.0-resolved.yaml


@param GetTheDetailsOfIssuesForGivenSetOfFiltersV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-the-details-of-issues-for-given-set-of-filters
*/
func (s *IssuesService) GetTheDetailsOfIssuesForGivenSetOfFiltersV1(requestIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1 *RequestIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1, GetTheDetailsOfIssuesForGivenSetOfFiltersV1HeaderParams *GetTheDetailsOfIssuesForGivenSetOfFiltersV1HeaderParams) (*ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceIssues/query"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTheDetailsOfIssuesForGivenSetOfFiltersV1HeaderParams != nil {

		if GetTheDetailsOfIssuesForGivenSetOfFiltersV1HeaderParams.AcceptLanguage != "" {
			clientRequest = clientRequest.SetHeader("Accept-Language", GetTheDetailsOfIssuesForGivenSetOfFiltersV1HeaderParams.AcceptLanguage)
		}

		if GetTheDetailsOfIssuesForGivenSetOfFiltersV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTheDetailsOfIssuesForGivenSetOfFiltersV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1).
		SetResult(&ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTheDetailsOfIssuesForGivenSetOfFiltersV1(requestIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1, GetTheDetailsOfIssuesForGivenSetOfFiltersV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetTheDetailsOfIssuesForGivenSetOfFiltersV1")
	}

	result := response.Result().(*ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1)
	return result, response, err

}

//GetTheTotalNumberOfIssuesForGivenSetOfFiltersV1 Get the total number of issues for given set of filters - b3ad-493a-409b-90b4
/* Returns the total number issues for given set of filters. If there is no start and/or end time, then end time will be defaulted to current time and start time will be defaulted to 24-hours ago from end time. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-IssuesList-1.0.0-resolved.yaml


@param GetTheTotalNumberOfIssuesForGivenSetOfFiltersV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-the-total-number-of-issues-for-given-set-of-filters
*/
func (s *IssuesService) GetTheTotalNumberOfIssuesForGivenSetOfFiltersV1(requestIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersV1 *RequestIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersV1, GetTheTotalNumberOfIssuesForGivenSetOfFiltersV1HeaderParams *GetTheTotalNumberOfIssuesForGivenSetOfFiltersV1HeaderParams) (*ResponseIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersV1, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceIssues/query/count"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTheTotalNumberOfIssuesForGivenSetOfFiltersV1HeaderParams != nil {

		if GetTheTotalNumberOfIssuesForGivenSetOfFiltersV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTheTotalNumberOfIssuesForGivenSetOfFiltersV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersV1).
		SetResult(&ResponseIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTheTotalNumberOfIssuesForGivenSetOfFiltersV1(requestIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersV1, GetTheTotalNumberOfIssuesForGivenSetOfFiltersV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetTheTotalNumberOfIssuesForGivenSetOfFiltersV1")
	}

	result := response.Result().(*ResponseIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersV1)
	return result, response, err

}

//GetSummaryAnalyticsDataOfIssuesV1 Get summary analytics data of issues - afaa-2bdf-424b-9161
/* Gets the summary analytics data related to issues based on given filters and group by field. This data can be used to find issue counts grouped by different keys. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-IssuesList-1.0.1-resolved.yaml


@param GetSummaryAnalyticsDataOfIssuesV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-summary-analytics-data-of-issues
*/
func (s *IssuesService) GetSummaryAnalyticsDataOfIssuesV1(requestIssuesGetSummaryAnalyticsDataOfIssuesV1 *RequestIssuesGetSummaryAnalyticsDataOfIssuesV1, GetSummaryAnalyticsDataOfIssuesV1HeaderParams *GetSummaryAnalyticsDataOfIssuesV1HeaderParams) (*ResponseIssuesGetSummaryAnalyticsDataOfIssuesV1, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceIssues/summaryAnalytics"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetSummaryAnalyticsDataOfIssuesV1HeaderParams != nil {

		if GetSummaryAnalyticsDataOfIssuesV1HeaderParams.AcceptLanguage != "" {
			clientRequest = clientRequest.SetHeader("Accept-Language", GetSummaryAnalyticsDataOfIssuesV1HeaderParams.AcceptLanguage)
		}

		if GetSummaryAnalyticsDataOfIssuesV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetSummaryAnalyticsDataOfIssuesV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestIssuesGetSummaryAnalyticsDataOfIssuesV1).
		SetResult(&ResponseIssuesGetSummaryAnalyticsDataOfIssuesV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSummaryAnalyticsDataOfIssuesV1(requestIssuesGetSummaryAnalyticsDataOfIssuesV1, GetSummaryAnalyticsDataOfIssuesV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetSummaryAnalyticsDataOfIssuesV1")
	}

	result := response.Result().(*ResponseIssuesGetSummaryAnalyticsDataOfIssuesV1)
	return result, response, err

}

//GetTopNAnalyticsDataOfIssuesV1 Get Top N analytics data of issues - 21a7-c91a-4f5a-b54d
/* Gets the Top N analytics data related to issues based on given filters and group by field. This data can be used to find top sites which has most issues or top device types with most issue etc,. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-IssuesList-1.0.1-resolved.yaml


@param GetTopNAnalyticsDataOfIssuesV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-top-n-analytics-data-of-issues
*/
func (s *IssuesService) GetTopNAnalyticsDataOfIssuesV1(requestIssuesGetTopNAnalyticsDataOfIssuesV1 *RequestIssuesGetTopNAnalyticsDataOfIssuesV1, GetTopNAnalyticsDataOfIssuesV1HeaderParams *GetTopNAnalyticsDataOfIssuesV1HeaderParams) (*ResponseIssuesGetTopNAnalyticsDataOfIssuesV1, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceIssues/topNAnalytics"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTopNAnalyticsDataOfIssuesV1HeaderParams != nil {

		if GetTopNAnalyticsDataOfIssuesV1HeaderParams.AcceptLanguage != "" {
			clientRequest = clientRequest.SetHeader("Accept-Language", GetTopNAnalyticsDataOfIssuesV1HeaderParams.AcceptLanguage)
		}

		if GetTopNAnalyticsDataOfIssuesV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTopNAnalyticsDataOfIssuesV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestIssuesGetTopNAnalyticsDataOfIssuesV1).
		SetResult(&ResponseIssuesGetTopNAnalyticsDataOfIssuesV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTopNAnalyticsDataOfIssuesV1(requestIssuesGetTopNAnalyticsDataOfIssuesV1, GetTopNAnalyticsDataOfIssuesV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetTopNAnalyticsDataOfIssuesV1")
	}

	result := response.Result().(*ResponseIssuesGetTopNAnalyticsDataOfIssuesV1)
	return result, response, err

}

//GetTrendAnalyticsDataOfIssuesV1 Get trend analytics data of issues - f9ae-db6a-4618-b045
/* Gets the trend analytics data related to issues based on given filters and group by field. This data can be used to find issue counts in different intervals over a period of time. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-IssuesList-1.0.1-resolved.yaml


@param GetTrendAnalyticsDataOfIssuesV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-trend-analytics-data-of-issues
*/
func (s *IssuesService) GetTrendAnalyticsDataOfIssuesV1(requestIssuesGetTrendAnalyticsDataOfIssuesV1 *RequestIssuesGetTrendAnalyticsDataOfIssuesV1, GetTrendAnalyticsDataOfIssuesV1HeaderParams *GetTrendAnalyticsDataOfIssuesV1HeaderParams) (*ResponseIssuesGetTrendAnalyticsDataOfIssuesV1, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceIssues/trendAnalytics"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTrendAnalyticsDataOfIssuesV1HeaderParams != nil {

		if GetTrendAnalyticsDataOfIssuesV1HeaderParams.AcceptLanguage != "" {
			clientRequest = clientRequest.SetHeader("Accept-Language", GetTrendAnalyticsDataOfIssuesV1HeaderParams.AcceptLanguage)
		}

		if GetTrendAnalyticsDataOfIssuesV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTrendAnalyticsDataOfIssuesV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestIssuesGetTrendAnalyticsDataOfIssuesV1).
		SetResult(&ResponseIssuesGetTrendAnalyticsDataOfIssuesV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTrendAnalyticsDataOfIssuesV1(requestIssuesGetTrendAnalyticsDataOfIssuesV1, GetTrendAnalyticsDataOfIssuesV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetTrendAnalyticsDataOfIssuesV1")
	}

	result := response.Result().(*ResponseIssuesGetTrendAnalyticsDataOfIssuesV1)
	return result, response, err

}

//IgnoreTheGivenListOfIssuesV1 Ignore the given list of issues - 4b92-ca6b-4918-b9fd
/* Ignores the given list of issues. The response contains the list of issues which were successfully ignored as well as the issues which are failed to ignore. After this API returns success response, it may take few seconds for the issue status to be updated if the system is heavily loaded. Please use `GET /dna/data/api/v1/assuranceIssues/{id}` API to fetch the details of a particular issue and verify `updatedTime`. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-IssuesLifecycle-1.0.0-resolved.yaml


@param IgnoreTheGivenListOfIssuesV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!ignore-the-given-list-of-issues
*/
func (s *IssuesService) IgnoreTheGivenListOfIssuesV1(requestIssuesIgnoreTheGivenListOfIssuesV1 *RequestIssuesIgnoreTheGivenListOfIssuesV1, IgnoreTheGivenListOfIssuesV1HeaderParams *IgnoreTheGivenListOfIssuesV1HeaderParams) (*ResponseIssuesIgnoreTheGivenListOfIssuesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/assuranceIssues/ignore"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if IgnoreTheGivenListOfIssuesV1HeaderParams != nil {

		if IgnoreTheGivenListOfIssuesV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", IgnoreTheGivenListOfIssuesV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestIssuesIgnoreTheGivenListOfIssuesV1).
		SetResult(&ResponseIssuesIgnoreTheGivenListOfIssuesV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.IgnoreTheGivenListOfIssuesV1(requestIssuesIgnoreTheGivenListOfIssuesV1, IgnoreTheGivenListOfIssuesV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation IgnoreTheGivenListOfIssuesV1")
	}

	result := response.Result().(*ResponseIssuesIgnoreTheGivenListOfIssuesV1)
	return result, response, err

}

//ResolveTheGivenListsOfIssuesV1 Resolve the given lists of issues - d48f-a9ed-4929-a6dd
/* Resolves the given list of issues. The response contains the list of issues which were successfully resolved as well as the issues which are failed to resolve. After this API returns success response, it may take few seconds for the issue status to be updated if the system is heavily loaded. Please use `GET /dna/data/api/v1/assuranceIssues/{id}` API to fetch the details of a particular issue and verify `updatedTime`. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-IssuesLifecycle-1.0.0-resolved.yaml


@param ResolveTheGivenListsOfIssuesV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!resolve-the-given-lists-of-issues
*/
func (s *IssuesService) ResolveTheGivenListsOfIssuesV1(requestIssuesResolveTheGivenListsOfIssuesV1 *RequestIssuesResolveTheGivenListsOfIssuesV1, ResolveTheGivenListsOfIssuesV1HeaderParams *ResolveTheGivenListsOfIssuesV1HeaderParams) (*ResponseIssuesResolveTheGivenListsOfIssuesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/assuranceIssues/resolve"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if ResolveTheGivenListsOfIssuesV1HeaderParams != nil {

		if ResolveTheGivenListsOfIssuesV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", ResolveTheGivenListsOfIssuesV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestIssuesResolveTheGivenListsOfIssuesV1).
		SetResult(&ResponseIssuesResolveTheGivenListsOfIssuesV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ResolveTheGivenListsOfIssuesV1(requestIssuesResolveTheGivenListsOfIssuesV1, ResolveTheGivenListsOfIssuesV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation ResolveTheGivenListsOfIssuesV1")
	}

	result := response.Result().(*ResponseIssuesResolveTheGivenListsOfIssuesV1)
	return result, response, err

}

//UpdateTheGivenIssueByUpdatingSelectedFieldsV1 Update the given issue by updating selected fields - b0bc-dba1-4c19-8d7c
/* Updates selected fields in the given issue. Currently the only field that can be updated is 'notes' field. After this API returns success response, it may take few seconds for the issue details to be updated if the system is heavily loaded. Please use `GET /dna/data/api/v1/assuranceIssues/{id}` API to fetch the details of a particular issue and verify `updatedTime`. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-IssuesLifecycle-1.0.0-resolved.yaml


@param id id path parameter. The issue Uuid

@param UpdateTheGivenIssueByUpdatingSelectedFieldsV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!update-the-given-issue-by-updating-selected-fields
*/
func (s *IssuesService) UpdateTheGivenIssueByUpdatingSelectedFieldsV1(id string, requestIssuesUpdateTheGivenIssueByUpdatingSelectedFieldsV1 *RequestIssuesUpdateTheGivenIssueByUpdatingSelectedFieldsV1, UpdateTheGivenIssueByUpdatingSelectedFieldsV1HeaderParams *UpdateTheGivenIssueByUpdatingSelectedFieldsV1HeaderParams) (*ResponseIssuesUpdateTheGivenIssueByUpdatingSelectedFieldsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/assuranceIssues/{id}/update"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if UpdateTheGivenIssueByUpdatingSelectedFieldsV1HeaderParams != nil {

		if UpdateTheGivenIssueByUpdatingSelectedFieldsV1HeaderParams.AcceptLanguage != "" {
			clientRequest = clientRequest.SetHeader("Accept-Language", UpdateTheGivenIssueByUpdatingSelectedFieldsV1HeaderParams.AcceptLanguage)
		}

		if UpdateTheGivenIssueByUpdatingSelectedFieldsV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", UpdateTheGivenIssueByUpdatingSelectedFieldsV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestIssuesUpdateTheGivenIssueByUpdatingSelectedFieldsV1).
		SetResult(&ResponseIssuesUpdateTheGivenIssueByUpdatingSelectedFieldsV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateTheGivenIssueByUpdatingSelectedFieldsV1(id, requestIssuesUpdateTheGivenIssueByUpdatingSelectedFieldsV1, UpdateTheGivenIssueByUpdatingSelectedFieldsV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation UpdateTheGivenIssueByUpdatingSelectedFieldsV1")
	}

	result := response.Result().(*ResponseIssuesUpdateTheGivenIssueByUpdatingSelectedFieldsV1)
	return result, response, err

}

//CreatesANewUserDefinedIssueDefinitionsV1 Creates a new user-defined issue definitions. - 95b5-9b50-4e48-9d82
/* Create a new custom issue definition using the provided input request data. The unique identifier for this issue definition is id. Please note that the issue names cannot be duplicated. The definition is based on the syslog. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceUserDefinedIssueAPIs-1.0.0-resolved.yaml


@param CreatesANewUserDefinedIssueDefinitionsV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!creates-a-new-user-defined-issue-definitions
*/
func (s *IssuesService) CreatesANewUserDefinedIssueDefinitionsV1(requestIssuesCreatesANewUserDefinedIssueDefinitionsV1 *RequestIssuesCreatesANewUserDefinedIssueDefinitionsV1, CreatesANewUserDefinedIssueDefinitionsV1HeaderParams *CreatesANewUserDefinedIssueDefinitionsV1HeaderParams) (*ResponseIssuesCreatesANewUserDefinedIssueDefinitionsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/customIssueDefinitions"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if CreatesANewUserDefinedIssueDefinitionsV1HeaderParams != nil {

		if CreatesANewUserDefinedIssueDefinitionsV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", CreatesANewUserDefinedIssueDefinitionsV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestIssuesCreatesANewUserDefinedIssueDefinitionsV1).
		SetResult(&ResponseIssuesCreatesANewUserDefinedIssueDefinitionsV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreatesANewUserDefinedIssueDefinitionsV1(requestIssuesCreatesANewUserDefinedIssueDefinitionsV1, CreatesANewUserDefinedIssueDefinitionsV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation CreatesANewUserDefinedIssueDefinitionsV1")
	}

	result := response.Result().(*ResponseIssuesCreatesANewUserDefinedIssueDefinitionsV1)
	return result, response, err

}

//ExecuteSuggestedActionsCommandsV1 Execute Suggested Actions Commands - cfb2-ab10-4cea-bfbb
/* This API fetches the issue details and suggested actions for an issue, given the Issue Id, executes the commands associated with the suggested actions to remediate the issue



Documentation Link: https://developer.cisco.com/docs/dna-center/#!execute-suggested-actions-commands
*/
func (s *IssuesService) ExecuteSuggestedActionsCommandsV1(requestIssuesExecuteSuggestedActionsCommandsV1 *RequestIssuesExecuteSuggestedActionsCommandsV1) (*ResponseIssuesExecuteSuggestedActionsCommandsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/execute-suggested-actions-commands"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestIssuesExecuteSuggestedActionsCommandsV1).
		SetResult(&ResponseIssuesExecuteSuggestedActionsCommandsV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ExecuteSuggestedActionsCommandsV1(requestIssuesExecuteSuggestedActionsCommandsV1)
		}

		return nil, response, fmt.Errorf("error with operation ExecuteSuggestedActionsCommandsV1")
	}

	result := response.Result().(*ResponseIssuesExecuteSuggestedActionsCommandsV1)
	return result, response, err

}

//UpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIDV1 Updates an existing custom issue definition based on the provided Id. - 8b90-3b69-4c18-90ad
/* Updates an existing custom issue definition based on the provided Id. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceUserDefinedIssueAPIs-1.0.0-resolved.yaml


@param id id path parameter. The custom issue definition Identifier

@param UpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIdV1HeaderParams Custom header parameters
*/
func (s *IssuesService) UpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIDV1(id string, requestIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIdV1 *RequestIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIDV1) (*ResponseIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/customIssueDefinitions/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	response, err = clientRequest.
		SetBody(requestIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIdV1).
		SetResult(&ResponseIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIDV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIDV1(id, requestIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIdV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIdV1")
	}

	result := response.Result().(*ResponseIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIDV1)
	return result, response, err

}

//IssueTriggerDefinitionUpdateV1 Issue trigger definition update. - 099a-397b-46c8-8aa7
/* Update issue trigger threshold, priority for the given id.
Also enable or disable issue trigger for the given id. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-issueAndHealthDefinitions-1.0.0-resolved.yaml


@param id id path parameter. Issue trigger definition id.

*/
func (s *IssuesService) IssueTriggerDefinitionUpdateV1(id string, requestIssuesIssueTriggerDefinitionUpdateV1 *RequestIssuesIssueTriggerDefinitionUpdateV1) (*ResponseIssuesIssueTriggerDefinitionUpdateV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/systemIssueDefinitions/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestIssuesIssueTriggerDefinitionUpdateV1).
		SetResult(&ResponseIssuesIssueTriggerDefinitionUpdateV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.IssueTriggerDefinitionUpdateV1(id, requestIssuesIssueTriggerDefinitionUpdateV1)
		}
		return nil, response, fmt.Errorf("error with operation IssueTriggerDefinitionUpdateV1")
	}

	result := response.Result().(*ResponseIssuesIssueTriggerDefinitionUpdateV1)
	return result, response, err

}

//DeletesAnExistingCustomIssueDefinitionV1 Deletes an existing custom issue definition. - e38b-fa80-4c28-955f
/* Deletes an existing custom issue definition based on the Id. Only the Global profile issue has the access to delete the issue definition, so no profile id is required. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceUserDefinedIssueAPIs-1.0.0-resolved.yaml


@param id id path parameter. The custom issue definition unique identifier


Documentation Link: https://developer.cisco.com/docs/dna-center/#!deletes-an-existing-custom-issue-definition
*/
func (s *IssuesService) DeletesAnExistingCustomIssueDefinitionV1(id string) (*resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/customIssueDefinitions/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletesAnExistingCustomIssueDefinitionV1(id)
		}
		return response, fmt.Errorf("error with operation DeletesAnExistingCustomIssueDefinitionV1")
	}

	return response, err

}

// Alias Function
/*
This method acts as an alias for the method `GetIssueEnrichmentDetailsV1`
*/
func (s *IssuesService) GetIssueEnrichmentDetails(GetIssueEnrichmentDetailsV1HeaderParams *GetIssueEnrichmentDetailsV1HeaderParams) (*ResponseIssuesGetIssueEnrichmentDetailsV1, *resty.Response, error) {
	return s.GetIssueEnrichmentDetailsV1(GetIssueEnrichmentDetailsV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDV1`
*/
func (s *IssuesService) GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionID(id string, GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIdV1HeaderParams *GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDV1HeaderParams) (*ResponseIssuesGetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDV1, *resty.Response, error) {
	return s.GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDV1(id, GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIdV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetSummaryAnalyticsDataOfIssuesV1`
*/
func (s *IssuesService) GetSummaryAnalyticsDataOfIssues(requestIssuesGetSummaryAnalyticsDataOfIssuesV1 *RequestIssuesGetSummaryAnalyticsDataOfIssuesV1, GetSummaryAnalyticsDataOfIssuesV1HeaderParams *GetSummaryAnalyticsDataOfIssuesV1HeaderParams) (*ResponseIssuesGetSummaryAnalyticsDataOfIssuesV1, *resty.Response, error) {
	return s.GetSummaryAnalyticsDataOfIssuesV1(requestIssuesGetSummaryAnalyticsDataOfIssuesV1, GetSummaryAnalyticsDataOfIssuesV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `ExecuteSuggestedActionsCommandsV1`
*/
func (s *IssuesService) ExecuteSuggestedActionsCommands(requestIssuesExecuteSuggestedActionsCommandsV1 *RequestIssuesExecuteSuggestedActionsCommandsV1) (*ResponseIssuesExecuteSuggestedActionsCommandsV1, *resty.Response, error) {
	return s.ExecuteSuggestedActionsCommandsV1(requestIssuesExecuteSuggestedActionsCommandsV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersV1`
*/
func (s *IssuesService) GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFilters(GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersV1HeaderParams *GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersV1HeaderParams, GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersV1QueryParams *GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersV1QueryParams) (*ResponseIssuesGetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersV1, *resty.Response, error) {
	return s.GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersV1(GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersV1HeaderParams, GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetTheTotalNumberOfIssuesForGivenSetOfFiltersV1`
*/
func (s *IssuesService) GetTheTotalNumberOfIssuesForGivenSetOfFilters(requestIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersV1 *RequestIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersV1, GetTheTotalNumberOfIssuesForGivenSetOfFiltersV1HeaderParams *GetTheTotalNumberOfIssuesForGivenSetOfFiltersV1HeaderParams) (*ResponseIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersV1, *resty.Response, error) {
	return s.GetTheTotalNumberOfIssuesForGivenSetOfFiltersV1(requestIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersV1, GetTheTotalNumberOfIssuesForGivenSetOfFiltersV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetTopNAnalyticsDataOfIssuesV1`
*/
func (s *IssuesService) GetTopNAnalyticsDataOfIssues(requestIssuesGetTopNAnalyticsDataOfIssuesV1 *RequestIssuesGetTopNAnalyticsDataOfIssuesV1, GetTopNAnalyticsDataOfIssuesV1HeaderParams *GetTopNAnalyticsDataOfIssuesV1HeaderParams) (*ResponseIssuesGetTopNAnalyticsDataOfIssuesV1, *resty.Response, error) {
	return s.GetTopNAnalyticsDataOfIssuesV1(requestIssuesGetTopNAnalyticsDataOfIssuesV1, GetTopNAnalyticsDataOfIssuesV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetTrendAnalyticsDataOfIssuesV1`
*/
func (s *IssuesService) GetTrendAnalyticsDataOfIssues(requestIssuesGetTrendAnalyticsDataOfIssuesV1 *RequestIssuesGetTrendAnalyticsDataOfIssuesV1, GetTrendAnalyticsDataOfIssuesV1HeaderParams *GetTrendAnalyticsDataOfIssuesV1HeaderParams) (*ResponseIssuesGetTrendAnalyticsDataOfIssuesV1, *resty.Response, error) {
	return s.GetTrendAnalyticsDataOfIssuesV1(requestIssuesGetTrendAnalyticsDataOfIssuesV1, GetTrendAnalyticsDataOfIssuesV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetIssueTriggerDefinitionForGivenIDV1`
*/
func (s *IssuesService) GetIssueTriggerDefinitionForGivenID(id string, GetIssueTriggerDefinitionForGivenIdV1HeaderParams *GetIssueTriggerDefinitionForGivenIDV1HeaderParams) (*ResponseIssuesGetIssueTriggerDefinitionForGivenIDV1, *resty.Response, error) {
	return s.GetIssueTriggerDefinitionForGivenIDV1(id, GetIssueTriggerDefinitionForGivenIdV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `CreatesANewUserDefinedIssueDefinitionsV1`
*/
func (s *IssuesService) CreatesANewUserDefinedIssueDefinitions(requestIssuesCreatesANewUserDefinedIssueDefinitionsV1 *RequestIssuesCreatesANewUserDefinedIssueDefinitionsV1, CreatesANewUserDefinedIssueDefinitionsV1HeaderParams *CreatesANewUserDefinedIssueDefinitionsV1HeaderParams) (*ResponseIssuesCreatesANewUserDefinedIssueDefinitionsV1, *resty.Response, error) {
	return s.CreatesANewUserDefinedIssueDefinitionsV1(requestIssuesCreatesANewUserDefinedIssueDefinitionsV1, CreatesANewUserDefinedIssueDefinitionsV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetTheDetailsOfIssuesForGivenSetOfFiltersV1`
*/
func (s *IssuesService) GetTheDetailsOfIssuesForGivenSetOfFilters(requestIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1 *RequestIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1, GetTheDetailsOfIssuesForGivenSetOfFiltersV1HeaderParams *GetTheDetailsOfIssuesForGivenSetOfFiltersV1HeaderParams) (*ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1, *resty.Response, error) {
	return s.GetTheDetailsOfIssuesForGivenSetOfFiltersV1(requestIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1, GetTheDetailsOfIssuesForGivenSetOfFiltersV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `UpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIDV1`
*/
func (s *IssuesService) UpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedID(id string, requestIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIdV1 *RequestIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIDV1) (*ResponseIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIDV1, *resty.Response, error) {
	return s.UpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIDV1(id, requestIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIdV1)
}

// Alias Function
/*
This method acts as an alias for the method `IssuesV1`
*/
func (s *IssuesService) Issues(IssuesV1QueryParams *IssuesV1QueryParams) (*ResponseIssuesIssuesV1, *resty.Response, error) {
	return s.IssuesV1(IssuesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1`
*/
func (s *IssuesService) GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFilters(GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1HeaderParams *GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1HeaderParams, GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1QueryParams *GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1QueryParams) (*ResponseIssuesGetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1, *resty.Response, error) {
	return s.GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1(GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1HeaderParams, GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `ResolveTheGivenListsOfIssuesV1`
*/
func (s *IssuesService) ResolveTheGivenListsOfIssues(requestIssuesResolveTheGivenListsOfIssuesV1 *RequestIssuesResolveTheGivenListsOfIssuesV1, ResolveTheGivenListsOfIssuesV1HeaderParams *ResolveTheGivenListsOfIssuesV1HeaderParams) (*ResponseIssuesResolveTheGivenListsOfIssuesV1, *resty.Response, error) {
	return s.ResolveTheGivenListsOfIssuesV1(requestIssuesResolveTheGivenListsOfIssuesV1, ResolveTheGivenListsOfIssuesV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `IssueTriggerDefinitionUpdateV1`
*/
func (s *IssuesService) IssueTriggerDefinitionUpdate(id string, requestIssuesIssueTriggerDefinitionUpdateV1 *RequestIssuesIssueTriggerDefinitionUpdateV1) (*ResponseIssuesIssueTriggerDefinitionUpdateV1, *resty.Response, error) {
	return s.IssueTriggerDefinitionUpdateV1(id, requestIssuesIssueTriggerDefinitionUpdateV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1`
*/
func (s *IssuesService) GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetwork(GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1HeaderParams *GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1HeaderParams, GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1QueryParams *GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1QueryParams) (*ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1, *resty.Response, error) {
	return s.GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1(GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1HeaderParams, GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDV1`
*/
func (s *IssuesService) GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueID(id string, GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIdV1HeaderParams *GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDV1HeaderParams, GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIdV1QueryParams *GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDV1QueryParams) (*ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDV1, *resty.Response, error) {
	return s.GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDV1(id, GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIdV1HeaderParams, GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIdV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersV1`
*/
func (s *IssuesService) GetAllTheCustomIssueDefinitionsBasedOnTheGivenFilters(GetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersV1QueryParams *GetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersV1QueryParams) (*ResponseIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersV1, *resty.Response, error) {
	return s.GetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersV1(GetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `DeletesAnExistingCustomIssueDefinitionV1`
*/
func (s *IssuesService) DeletesAnExistingCustomIssueDefinition(id string) (*resty.Response, error) {
	return s.DeletesAnExistingCustomIssueDefinitionV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `IgnoreTheGivenListOfIssuesV1`
*/
func (s *IssuesService) IgnoreTheGivenListOfIssues(requestIssuesIgnoreTheGivenListOfIssuesV1 *RequestIssuesIgnoreTheGivenListOfIssuesV1, IgnoreTheGivenListOfIssuesV1HeaderParams *IgnoreTheGivenListOfIssuesV1HeaderParams) (*ResponseIssuesIgnoreTheGivenListOfIssuesV1, *resty.Response, error) {
	return s.IgnoreTheGivenListOfIssuesV1(requestIssuesIgnoreTheGivenListOfIssuesV1, IgnoreTheGivenListOfIssuesV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkV1`
*/
func (s *IssuesService) GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetwork(GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkV1HeaderParams *GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkV1HeaderParams, GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkV1QueryParams *GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkV1QueryParams) (*ResponseIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkV1, *resty.Response, error) {
	return s.GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkV1(GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkV1HeaderParams, GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `UpdateTheGivenIssueByUpdatingSelectedFieldsV1`
*/
func (s *IssuesService) UpdateTheGivenIssueByUpdatingSelectedFields(id string, requestIssuesUpdateTheGivenIssueByUpdatingSelectedFieldsV1 *RequestIssuesUpdateTheGivenIssueByUpdatingSelectedFieldsV1, UpdateTheGivenIssueByUpdatingSelectedFieldsV1HeaderParams *UpdateTheGivenIssueByUpdatingSelectedFieldsV1HeaderParams) (*ResponseIssuesUpdateTheGivenIssueByUpdatingSelectedFieldsV1, *resty.Response, error) {
	return s.UpdateTheGivenIssueByUpdatingSelectedFieldsV1(id, requestIssuesUpdateTheGivenIssueByUpdatingSelectedFieldsV1, UpdateTheGivenIssueByUpdatingSelectedFieldsV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `ReturnsAllIssueTriggerDefinitionsForGivenFiltersV1`
*/
func (s *IssuesService) ReturnsAllIssueTriggerDefinitionsForGivenFilters(ReturnsAllIssueTriggerDefinitionsForGivenFiltersV1HeaderParams *ReturnsAllIssueTriggerDefinitionsForGivenFiltersV1HeaderParams, ReturnsAllIssueTriggerDefinitionsForGivenFiltersV1QueryParams *ReturnsAllIssueTriggerDefinitionsForGivenFiltersV1QueryParams) (*ResponseIssuesReturnsAllIssueTriggerDefinitionsForGivenFiltersV1, *resty.Response, error) {
	return s.ReturnsAllIssueTriggerDefinitionsForGivenFiltersV1(ReturnsAllIssueTriggerDefinitionsForGivenFiltersV1HeaderParams, ReturnsAllIssueTriggerDefinitionsForGivenFiltersV1QueryParams)
}
