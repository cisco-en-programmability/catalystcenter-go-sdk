package catalyst

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SiteDesignService service

type GetSiteAssignedNetworkDevicesV1QueryParams struct {
	SiteID string  `url:"siteId,omitempty"` //Site Id. It must be area Id or building Id or floor Id.
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1.
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page.
}
type GetSiteAssignedNetworkDevicesCountV1QueryParams struct {
	SiteID string `url:"siteId,omitempty"` //Site Id. It must be area Id or building Id or floor Id.
}
type GetSiteNotAssignedNetworkDevicesV1QueryParams struct {
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1.
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page.
}
type RetrievesTheListOfNetworkProfilesForSitesV1QueryParams struct {
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1.
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page.
	SortBy string  `url:"sortBy,omitempty"` //A property within the response to sort by.
	Order  string  `url:"order,omitempty"`  //Whether ascending or descending order should be used to sort the response.
	Type   string  `url:"type,omitempty"`   //Filter responses to only include profiles of a given type
}
type RetrievesTheCountOfNetworkProfilesForSitesV1QueryParams struct {
	Type string `url:"type,omitempty"` //Filter the response to only count profiles of a given type
}
type RetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToV1QueryParams struct {
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1.
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page.
}
type UnassignsANetworkProfileForSitesFromMultipleSitesV1QueryParams struct {
	SiteID string `url:"siteId,omitempty"` //The `id` of the site, retrievable from `GET /intent/api/v1/sites`
}
type GetSitesV1QueryParams struct {
	Name           string `url:"name,omitempty"`            //Site name.
	NameHierarchy  string `url:"nameHierarchy,omitempty"`   //Site name hierarchy.
	Type           string `url:"type,omitempty"`            //Site type.
	UnitsOfMeasure string `url:"_unitsOfMeasure,omitempty"` //Floor units of measure
	Offset         int    `url:"offset,omitempty"`          //The first record to show for this page; the first record is numbered 1.
	Limit          int    `url:"limit,omitempty"`           //The number of records to show for this page.
}
type GetSitesCountV1QueryParams struct {
	Name string `url:"name,omitempty"` //Site name.
}
type RetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssignedV1QueryParams struct {
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1.
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page.
}
type GetsAFloorV2QueryParams struct {
	UnitsOfMeasure string `url:"_unitsOfMeasure,omitempty"` //Floor units of measure
}

type ResponseSiteDesignCreatesAnAreaV1 struct {
	Version  string                                     `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignCreatesAnAreaV1Response `json:"response,omitempty"` //
}
type ResponseSiteDesignCreatesAnAreaV1Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignUpdatesAnAreaV1 struct {
	Version  string                                     `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignUpdatesAnAreaV1Response `json:"response,omitempty"` //
}
type ResponseSiteDesignUpdatesAnAreaV1Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignDeletesAnAreaV1 struct {
	Version  string                                     `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignDeletesAnAreaV1Response `json:"response,omitempty"` //
}
type ResponseSiteDesignDeletesAnAreaV1Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignGetsAnAreaV1 struct {
	Response *ResponseSiteDesignGetsAnAreaV1Response `json:"response,omitempty"` //
	Version  string                                  `json:"version,omitempty"`  //
}
type ResponseSiteDesignGetsAnAreaV1Response struct {
	ID            string `json:"id,omitempty"`            // Aread Id. Read only.
	Name          string `json:"name,omitempty"`          // Area name
	NameHierarchy string `json:"nameHierarchy,omitempty"` // Area hierarchical name. Read only.
	ParentID      string `json:"parentId,omitempty"`      // Parent Id
	Type          string `json:"type,omitempty"`          // Site Type.
}
type ResponseSiteDesignAssignNetworkDevicesToASiteV1 struct {
	Version  string                                                   `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignAssignNetworkDevicesToASiteV1Response `json:"response,omitempty"` //
}
type ResponseSiteDesignAssignNetworkDevicesToASiteV1Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignGetSiteAssignedNetworkDevicesV1 struct {
	Response *[]ResponseSiteDesignGetSiteAssignedNetworkDevicesV1Response `json:"response,omitempty"` //
	Version  string                                                       `json:"version,omitempty"`  //
}
type ResponseSiteDesignGetSiteAssignedNetworkDevicesV1Response struct {
	DeviceID          string `json:"deviceId,omitempty"`          // Site assigned network device Id.
	SiteID            string `json:"siteId,omitempty"`            // Site Id where device has been assigned.
	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Site name hierarchy
	SiteType          string `json:"siteType,omitempty"`          // Type of the site where device has been assigned.
}
type ResponseSiteDesignGetSiteAssignedNetworkDevicesCountV1 struct {
	Response *ResponseSiteDesignGetSiteAssignedNetworkDevicesCountV1Response `json:"response,omitempty"` //
	Version  string                                                          `json:"version,omitempty"`  // The version of the response
}
type ResponseSiteDesignGetSiteAssignedNetworkDevicesCountV1Response struct {
	Count *int `json:"count,omitempty"` // The total number of records related to the resource
}
type ResponseSiteDesignGetDeviceControllabilitySettingsV1 struct {
	Response *ResponseSiteDesignGetDeviceControllabilitySettingsV1Response `json:"response,omitempty"` //
	Version  string                                                        `json:"version,omitempty"`  //
}
type ResponseSiteDesignGetDeviceControllabilitySettingsV1Response struct {
	AutocorrectTelemetryConfig *bool `json:"autocorrectTelemetryConfig,omitempty"` // If it is true, autocorrect telemetry config is enabled. If it is false, autocorrect telemetry config is disabled. The autocorrect telemetry config feature is supported only when device controllability is enabled.
	DeviceControllability      *bool `json:"deviceControllability,omitempty"`      // If it is true, device controllability is enabled. If it is false, device controllability is disabled.
}
type ResponseSiteDesignUpdateDeviceControllabilitySettingsV1 struct {
	Version  string                                                           `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignUpdateDeviceControllabilitySettingsV1Response `json:"response,omitempty"` //
}
type ResponseSiteDesignUpdateDeviceControllabilitySettingsV1Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignGetSiteNotAssignedNetworkDevicesV1 struct {
	Response *ResponseSiteDesignGetSiteNotAssignedNetworkDevicesV1Response `json:"response,omitempty"` //
	Version  string                                                        `json:"version,omitempty"`  //
}
type ResponseSiteDesignGetSiteNotAssignedNetworkDevicesV1Response struct {
	DeviceIDs []string `json:"deviceIds,omitempty"` // Network device Ids.
}
type ResponseSiteDesignGetSiteNotAssignedNetworkDevicesCountV1 struct {
	Response *ResponseSiteDesignGetSiteNotAssignedNetworkDevicesCountV1Response `json:"response,omitempty"` //
	Version  string                                                             `json:"version,omitempty"`  // The version of the response
}
type ResponseSiteDesignGetSiteNotAssignedNetworkDevicesCountV1Response struct {
	Count *int `json:"count,omitempty"` // The total number of records related to the resource
}
type ResponseSiteDesignUnassignNetworkDevicesFromSitesV1 struct {
	Version  string                                                       `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignUnassignNetworkDevicesFromSitesV1Response `json:"response,omitempty"` //
}
type ResponseSiteDesignUnassignNetworkDevicesFromSitesV1Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignGetSiteAssignedNetworkDeviceV1 struct {
	Response *ResponseSiteDesignGetSiteAssignedNetworkDeviceV1Response `json:"response,omitempty"` //
	Version  string                                                    `json:"version,omitempty"`  //
}
type ResponseSiteDesignGetSiteAssignedNetworkDeviceV1Response struct {
	DeviceID          string `json:"deviceId,omitempty"`          // Site assigned network device Id.
	SiteID            string `json:"siteId,omitempty"`            // Site Id where device has been assigned.
	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Site name hierarchy
	SiteType          string `json:"siteType,omitempty"`          // Type of the site where device has been assigned.
}
type ResponseSiteDesignRetrievesTheListOfNetworkProfilesForSitesV1 struct {
	Response *[]ResponseSiteDesignRetrievesTheListOfNetworkProfilesForSitesV1Response `json:"response,omitempty"` //
	Version  string                                                                   `json:"version,omitempty"`  // Version
}
type ResponseSiteDesignRetrievesTheListOfNetworkProfilesForSitesV1Response struct {
	ID   string `json:"id,omitempty"`   // The ID of this network profile.
	Name string `json:"name,omitempty"` // The name of the network profile.
	Type string `json:"type,omitempty"` // Type
}
type ResponseSiteDesignRetrievesTheCountOfNetworkProfilesForSitesV1 struct {
	Response *ResponseSiteDesignRetrievesTheCountOfNetworkProfilesForSitesV1Response `json:"response,omitempty"` //
	Version  string                                                                  `json:"version,omitempty"`  // Version
}
type ResponseSiteDesignRetrievesTheCountOfNetworkProfilesForSitesV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseSiteDesignDeletesANetworkProfileForSitesV1 struct {
	Version  string                                                      `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignDeletesANetworkProfileForSitesV1Response `json:"response,omitempty"` //
}
type ResponseSiteDesignDeletesANetworkProfileForSitesV1Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignRetrieveANetworkProfileForSitesByIDV1 struct {
	Response *ResponseSiteDesignRetrieveANetworkProfileForSitesByIDV1Response `json:"response,omitempty"` //
}
type ResponseSiteDesignRetrieveANetworkProfileForSitesByIDV1Response struct {
	ID   string `json:"id,omitempty"`   // The ID of this network profile.
	Name string `json:"name,omitempty"` // The name of the network profile.
	Type string `json:"type,omitempty"` // Type
}
type ResponseSiteDesignAssignANetworkProfileForSitesToTheGivenSiteV1 struct {
	Version  string                                                                   `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignAssignANetworkProfileForSitesToTheGivenSiteV1Response `json:"response,omitempty"` //
}
type ResponseSiteDesignAssignANetworkProfileForSitesToTheGivenSiteV1Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignRetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToV1 struct {
	Response *[]ResponseSiteDesignRetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToV1Response `json:"response,omitempty"` //
	Version  string                                                                                               `json:"version,omitempty"`  // Version
}
type ResponseSiteDesignRetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToV1Response struct {
	ID string `json:"id,omitempty"` // Id
}
type ResponseSiteDesignAssignANetworkProfileForSitesToAListOfSitesV1 struct {
	Version  string                                                                   `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignAssignANetworkProfileForSitesToAListOfSitesV1Response `json:"response,omitempty"` //
}
type ResponseSiteDesignAssignANetworkProfileForSitesToAListOfSitesV1Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignUnassignsANetworkProfileForSitesFromMultipleSitesV1 struct {
	Version  string                                                                         `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignUnassignsANetworkProfileForSitesFromMultipleSitesV1Response `json:"response,omitempty"` //
}
type ResponseSiteDesignUnassignsANetworkProfileForSitesFromMultipleSitesV1Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignRetrievesTheCountOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToV1 struct {
	Response *ResponseSiteDesignRetrievesTheCountOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToV1Response `json:"response,omitempty"` //
	Version  string                                                                                              `json:"version,omitempty"`  // Version
}
type ResponseSiteDesignRetrievesTheCountOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseSiteDesignUnassignsANetworkProfileForSitesFromASiteV1 struct {
	Version  string                                                                 `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignUnassignsANetworkProfileForSitesFromASiteV1Response `json:"response,omitempty"` //
}
type ResponseSiteDesignUnassignsANetworkProfileForSitesFromASiteV1Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignAssociateV1 struct {
	Version  string                                 `json:"version,omitempty"`  // Version
	Response *ResponseSiteDesignAssociateV1Response `json:"response,omitempty"` //
}
type ResponseSiteDesignAssociateV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseSiteDesignDisassociateV1 struct {
	Version  string                                    `json:"version,omitempty"`  // Version
	Response *ResponseSiteDesignDisassociateV1Response `json:"response,omitempty"` //
}
type ResponseSiteDesignDisassociateV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseSiteDesignGetSitesV1 struct {
	Response *[]ResponseSiteDesignGetSitesV1Response `json:"response,omitempty"` //
	Version  string                                  `json:"version,omitempty"`  // Version
}
type ResponseSiteDesignGetSitesV1Response struct {
	NameHierarchy  string   `json:"nameHierarchy,omitempty"`  // Site hierarchical name. Read only. Example: Global/USA/San Jose/Building1
	Name           string   `json:"name,omitempty"`           // Site name.
	Latitude       *float64 `json:"latitude,omitempty"`       // Building Latitude. Example: 37.403712
	Longitude      *float64 `json:"longitude,omitempty"`      // Building Longitude. Example: -121.971063
	Address        string   `json:"address,omitempty"`        // Building address. Example: 4900 Marie P. Debartolo Way, Santa Clara, California 95054, United States
	Country        string   `json:"country,omitempty"`        // Country name for the building.
	FloorNumber    *int     `json:"floorNumber,omitempty"`    // Floor number
	RfModel        string   `json:"rfModel,omitempty"`        // Floor RF Model
	Width          *float64 `json:"width,omitempty"`          // Floor width. Example : 100.5
	Length         *float64 `json:"length,omitempty"`         // Floor length. Example : 110.3
	Height         *float64 `json:"height,omitempty"`         // Floor height. Example : 10.1
	UnitsOfMeasure string   `json:"unitsOfMeasure,omitempty"` // Floor unit of measure
	Type           string   `json:"type,omitempty"`           // Type
	ID             string   `json:"id,omitempty"`             // Site Id. Read only.
	ParentID       string   `json:"parentId,omitempty"`       // Parent Id. Read only
}
type ResponseSiteDesignCreateSitesV1 struct {
	Version  string                                   `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignCreateSitesV1Response `json:"response,omitempty"` //
}
type ResponseSiteDesignCreateSitesV1Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignGetSitesCountV1 []ResponseItemSiteDesignGetSitesCountV1 // Array of ResponseSiteDesignGetSitesCountV1
type ResponseItemSiteDesignGetSitesCountV1 struct {
	Response *ResponseItemSiteDesignGetSitesCountV1Response `json:"response,omitempty"` //
	Version  string                                         `json:"version,omitempty"`  // The version of the response
}
type ResponseItemSiteDesignGetSitesCountV1Response struct {
	Count *int `json:"count,omitempty"` // The reported count.
}
type ResponseSiteDesignRetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssignedV1 struct {
	Response *[]ResponseSiteDesignRetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssignedV1Response `json:"response,omitempty"` //
	Version  string                                                                                          `json:"version,omitempty"`  // Version
}
type ResponseSiteDesignRetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssignedV1Response struct {
	ID string `json:"id,omitempty"` // Id
}
type ResponseSiteDesignRetrievesTheCountOfProfilesThatTheGivenSiteHasBeenAssignedV1 struct {
	Response *ResponseSiteDesignRetrievesTheCountOfProfilesThatTheGivenSiteHasBeenAssignedV1Response `json:"response,omitempty"` //
	Version  string                                                                                  `json:"version,omitempty"`  // Version
}
type ResponseSiteDesignRetrievesTheCountOfProfilesThatTheGivenSiteHasBeenAssignedV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseSiteDesignCreatesABuildingV2 struct {
	Version  string                                        `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignCreatesABuildingV2Response `json:"response,omitempty"` //
}
type ResponseSiteDesignCreatesABuildingV2Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignUpdatesABuildingV2 struct {
	Version  string                                        `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignUpdatesABuildingV2Response `json:"response,omitempty"` //
}
type ResponseSiteDesignUpdatesABuildingV2Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignDeletesABuildingV2 struct {
	Version  string                                        `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignDeletesABuildingV2Response `json:"response,omitempty"` //
}
type ResponseSiteDesignDeletesABuildingV2Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignGetsABuildingV2 struct {
	Response *ResponseSiteDesignGetsABuildingV2Response `json:"response,omitempty"` //
}
type ResponseSiteDesignGetsABuildingV2Response struct {
	ParentID  string   `json:"parentId,omitempty"`  // Parent Id
	Name      string   `json:"name,omitempty"`      // Building name
	Latitude  *float64 `json:"latitude,omitempty"`  // Building Latitude. Example: 37.403712
	Longitude *float64 `json:"longitude,omitempty"` // Building Longitude. Example: -121.971063
	Address   string   `json:"address,omitempty"`   // Building address. Example: 4900 Marie P. Debartolo Way, Santa Clara, California 95054, United States
	Country   string   `json:"country,omitempty"`   // Country name
	Type      string   `json:"type,omitempty"`      // Example: building
}
type ResponseSiteDesignCreatesAFloorV2 struct {
	Version  string                                     `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignCreatesAFloorV2Response `json:"response,omitempty"` //
}
type ResponseSiteDesignCreatesAFloorV2Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignUpdatesFloorSettingsV2 struct {
	Version  string                                            `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignUpdatesFloorSettingsV2Response `json:"response,omitempty"` //
}
type ResponseSiteDesignUpdatesFloorSettingsV2Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignGetFloorSettingsV2 struct {
	Response *ResponseSiteDesignGetFloorSettingsV2Response `json:"response,omitempty"` //
	Version  string                                        `json:"version,omitempty"`  // Version
}
type ResponseSiteDesignGetFloorSettingsV2Response struct {
	UnitsOfMeasure string `json:"unitsOfMeasure,omitempty"` // Floor units of measure.
}
type ResponseSiteDesignUpdatesAFloorV2 struct {
	Version  string                                     `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignUpdatesAFloorV2Response `json:"response,omitempty"` //
}
type ResponseSiteDesignUpdatesAFloorV2Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignGetsAFloorV2 struct {
	Response *ResponseSiteDesignGetsAFloorV2Response `json:"response,omitempty"` //
}
type ResponseSiteDesignGetsAFloorV2Response struct {
	ParentID       string   `json:"parentId,omitempty"`       // Parent Id.
	Name           string   `json:"name,omitempty"`           // Floor name
	FloorNumber    *int     `json:"floorNumber,omitempty"`    // Floor number
	RfModel        string   `json:"rfModel,omitempty"`        // RF Model
	Width          *float64 `json:"width,omitempty"`          // Floor width. Example : 100.5
	Length         *float64 `json:"length,omitempty"`         // Floor length. Example : 110.3
	Height         *float64 `json:"height,omitempty"`         // Floor height. Example : 10.1
	UnitsOfMeasure string   `json:"unitsOfMeasure,omitempty"` // Units Of Measure
	Type           string   `json:"type,omitempty"`           // Example : floor
	ID             string   `json:"id,omitempty"`             // Floor Id. Read only.
	NameHierarchy  string   `json:"nameHierarchy,omitempty"`  // Floor hierarchical name. Read only.
}
type ResponseSiteDesignDeletesAFloorV2 struct {
	Version  string                                     `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignDeletesAFloorV2Response `json:"response,omitempty"` //
}
type ResponseSiteDesignDeletesAFloorV2Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type RequestSiteDesignCreatesAnAreaV1 struct {
	Name     string `json:"name,omitempty"`     // Area name
	ParentID string `json:"parentId,omitempty"` // Parent Id
}
type RequestSiteDesignUpdatesAnAreaV1 struct {
	Name     string `json:"name,omitempty"`     // Area name
	ParentID string `json:"parentId,omitempty"` // Parent Id
}
type RequestSiteDesignAssignNetworkDevicesToASiteV1 struct {
	DeviceIDs []string `json:"deviceIds,omitempty"` // Unassigned network devices.
	SiteID    string   `json:"siteId,omitempty"`    // This must be building Id or floor Id. Access points, Sensors are assigned to floor. Remaining network devices are assigned to building. Site Id can be retrieved using '/intent/api/v1/sites'.
}
type RequestSiteDesignUpdateDeviceControllabilitySettingsV1 struct {
	AutocorrectTelemetryConfig *bool `json:"autocorrectTelemetryConfig,omitempty"` // If it is true, autocorrect telemetry config is enabled. If it is false, autocorrect telemetry config is disabled. The autocorrect telemetry config feature is supported only when device controllability is enabled.
	DeviceControllability      *bool `json:"deviceControllability,omitempty"`      // If it is true, device controllability is enabled. If it is false, device controllability is disabled.
}
type RequestSiteDesignUnassignNetworkDevicesFromSitesV1 struct {
	DeviceIDs []string `json:"deviceIds,omitempty"` // Network device Ids.
}
type RequestSiteDesignAssignANetworkProfileForSitesToTheGivenSiteV1 struct {
	ID string `json:"id,omitempty"` // Id
}
type RequestSiteDesignAssignANetworkProfileForSitesToAListOfSitesV1 struct {
	Type *RequestSiteDesignAssignANetworkProfileForSitesToAListOfSitesV1Type `json:"type,omitempty"` //
}
type RequestSiteDesignAssignANetworkProfileForSitesToAListOfSitesV1Type interface{}
type RequestSiteDesignCreateSitesV1 []RequestItemSiteDesignCreateSitesV1 // Array of RequestSiteDesignCreateSitesV1
type RequestItemSiteDesignCreateSitesV1 struct {
	ParentNameHierarchy string   `json:"parentNameHierarchy,omitempty"` // Parent hierarchical name. Example: Global/USA/San Jose/Building1
	Name                string   `json:"name,omitempty"`                // Site name.
	Latitude            *float64 `json:"latitude,omitempty"`            // Building Latitude. Example: 37.403712
	Longitude           *float64 `json:"longitude,omitempty"`           // Building Longitude. Example: -121.971063
	Address             string   `json:"address,omitempty"`             // Building address. Example: 4900 Marie P. Debartolo Way, Santa Clara, California 95054, United States
	Country             string   `json:"country,omitempty"`             // Country name. Required for building.
	FloorNumber         *int     `json:"floorNumber,omitempty"`         // Floor number. Required for floor.
	RfModel             string   `json:"rfModel,omitempty"`             // Floor RF Model. Required for floor.
	Width               *float64 `json:"width,omitempty"`               // Floor width. Required for floor. Example : 100.5
	Length              *float64 `json:"length,omitempty"`              // Floor length. Required for floor. Example : 110.3
	Height              *float64 `json:"height,omitempty"`              // Floor height. Required for floor. Example : 10.1
	UnitsOfMeasure      string   `json:"unitsOfMeasure,omitempty"`      // Floor unit of measure. Required for floor.
	Type                string   `json:"type,omitempty"`                // Type
}
type RequestSiteDesignCreatesABuildingV2 struct {
	ParentID  string   `json:"parentId,omitempty"`  // Parent Id
	Name      string   `json:"name,omitempty"`      // Building name
	Latitude  *float64 `json:"latitude,omitempty"`  // Building Latitude. Example: 37.403712
	Longitude *float64 `json:"longitude,omitempty"` // Building Longitude. Example: -121.971063
	Address   string   `json:"address,omitempty"`   // Building address. Example: 4900 Marie P. Debartolo Way, Santa Clara, California 95054, United States
	Country   string   `json:"country,omitempty"`   // Country name
}
type RequestSiteDesignUpdatesABuildingV2 struct {
	ParentID  string   `json:"parentId,omitempty"`  // Parent Id
	Name      string   `json:"name,omitempty"`      // Building name
	Latitude  *float64 `json:"latitude,omitempty"`  // Building Latitude. Example: 37.403712
	Longitude *float64 `json:"longitude,omitempty"` // Building Longitude. Example: -121.971063
	Address   string   `json:"address,omitempty"`   // Building address. Example: 4900 Marie P. Debartolo Way, Santa Clara, California 95054, United States
	Country   string   `json:"country,omitempty"`   // Country name
}
type RequestSiteDesignCreatesAFloorV2 struct {
	ParentID       string   `json:"parentId,omitempty"`       // Parent Id
	Name           string   `json:"name,omitempty"`           // Floor name
	FloorNumber    *int     `json:"floorNumber,omitempty"`    // Floor number
	RfModel        string   `json:"rfModel,omitempty"`        // RF Model
	Width          *float64 `json:"width,omitempty"`          // Floor width. Example : 100.5
	Length         *float64 `json:"length,omitempty"`         // Floor length. Example : 110.3
	Height         *float64 `json:"height,omitempty"`         // Floor height. Example : 10.1
	UnitsOfMeasure string   `json:"unitsOfMeasure,omitempty"` // Units Of Measure
}
type RequestSiteDesignUpdatesFloorSettingsV2 struct {
	UnitsOfMeasure string `json:"unitsOfMeasure,omitempty"` // Floor units of measure
}
type RequestSiteDesignUpdatesAFloorV2 struct {
	ParentID       string   `json:"parentId,omitempty"`       // Parent Id
	Name           string   `json:"name,omitempty"`           // Floor name
	FloorNumber    *int     `json:"floorNumber,omitempty"`    // Floor number
	RfModel        string   `json:"rfModel,omitempty"`        // RF Model
	Width          *float64 `json:"width,omitempty"`          // Floor width. Example : 100.5
	Length         *float64 `json:"length,omitempty"`         // Floor length. Example : 110.3
	Height         *float64 `json:"height,omitempty"`         // Floor height. Example : 10.1
	UnitsOfMeasure string   `json:"unitsOfMeasure,omitempty"` // Units Of Measure
}

//GetsAnAreaV1 Gets an area - d6af-ab3e-43bb-a73c
/* Gets an area in the network hierarchy.


@param id id path parameter. Area Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-an-area-v1
*/
func (s *SiteDesignService) GetsAnAreaV1(id string) (*ResponseSiteDesignGetsAnAreaV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/areas/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignGetsAnAreaV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsAnAreaV1(id)
		}
		return nil, response, fmt.Errorf("error with operation GetsAnAreaV1")
	}

	result := response.Result().(*ResponseSiteDesignGetsAnAreaV1)
	return result, response, err

}

//GetSiteAssignedNetworkDevicesV1 Get site assigned network devices - 0ea1-4875-4219-995d
/* Get all site assigned network devices. The items in the list are arranged in an order that corresponds with their internal identifiers.


@param GetSiteAssignedNetworkDevicesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-site-assigned-network-devices-v1
*/
func (s *SiteDesignService) GetSiteAssignedNetworkDevicesV1(GetSiteAssignedNetworkDevicesV1QueryParams *GetSiteAssignedNetworkDevicesV1QueryParams) (*ResponseSiteDesignGetSiteAssignedNetworkDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/assignedToSite"

	queryString, _ := query.Values(GetSiteAssignedNetworkDevicesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSiteDesignGetSiteAssignedNetworkDevicesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSiteAssignedNetworkDevicesV1(GetSiteAssignedNetworkDevicesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSiteAssignedNetworkDevicesV1")
	}

	result := response.Result().(*ResponseSiteDesignGetSiteAssignedNetworkDevicesV1)
	return result, response, err

}

//GetSiteAssignedNetworkDevicesCountV1 Get site assigned network devices count - fd93-c911-48ba-9386
/* Get all network devices count under the given site in the network hierarchy.


@param GetSiteAssignedNetworkDevicesCountV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-site-assigned-network-devices-count-v1
*/
func (s *SiteDesignService) GetSiteAssignedNetworkDevicesCountV1(GetSiteAssignedNetworkDevicesCountV1QueryParams *GetSiteAssignedNetworkDevicesCountV1QueryParams) (*ResponseSiteDesignGetSiteAssignedNetworkDevicesCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/assignedToSite/count"

	queryString, _ := query.Values(GetSiteAssignedNetworkDevicesCountV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSiteDesignGetSiteAssignedNetworkDevicesCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSiteAssignedNetworkDevicesCountV1(GetSiteAssignedNetworkDevicesCountV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSiteAssignedNetworkDevicesCountV1")
	}

	result := response.Result().(*ResponseSiteDesignGetSiteAssignedNetworkDevicesCountV1)
	return result, response, err

}

//GetDeviceControllabilitySettingsV1 Get device controllability settings - 06b2-4916-486b-aeec
/* Device Controllability is a system-level process on Catalyst Center that enforces state synchronization for some device-layer features. Its purpose is to aid in the deployment of required network settings that Catalyst Center needs to manage devices. Changes are made on network devices during discovery, when adding a device to Inventory, or when assigning a device to a site. If changes are made to any settings that are under the scope of this process, these changes are applied to the network devices during the Provision and Update Telemetry Settings operations, even if Device Controllability is disabled. The following device settings will be enabled as part of Device Controllability when devices are discovered. SNMP Credentials. NETCONF Credentials. Subsequent to discovery, devices will be added to Inventory. The following device settings will be enabled when devices are added to inventory. Cisco TrustSec (CTS) Credentials. The following device settings will be enabled when devices are assigned to a site. Some of these settings can be defined at a site level under Design > Network Settings > Telemetry & Wireless. Wired Endpoint Data Collection Enablement. Controller Certificates. SNMP Trap Server Definitions. Syslog Server Definitions. Application Visibility. Application QoS Policy. Wireless Service Assurance (WSA). Wireless Telemetry. DTLS Ciphersuite. AP Impersonation. If Device Controllability is disabled, Catalyst Center does not configure any of the preceding credentials or settings on devices during discovery, at runtime, or during site assignment. However, the telemetry settings and related configuration are pushed when the device is provisioned or when the update Telemetry Settings action is performed. Catalyst Center identifies and automatically corrects the following telemetry configuration issues on the device. SWIM certificate issue. IOS WLC NA certificate issue. PKCS12 certificate issue. IOS telemetry configuration issu



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-controllability-settings-v1
*/
func (s *SiteDesignService) GetDeviceControllabilitySettingsV1() (*ResponseSiteDesignGetDeviceControllabilitySettingsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/deviceControllability/settings"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignGetDeviceControllabilitySettingsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceControllabilitySettingsV1()
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceControllabilitySettingsV1")
	}

	result := response.Result().(*ResponseSiteDesignGetDeviceControllabilitySettingsV1)
	return result, response, err

}

//GetSiteNotAssignedNetworkDevicesV1 Get site not assigned network devices - cd89-78de-4109-8f0d
/* Get network devices that are not assigned to any site.


@param GetSiteNotAssignedNetworkDevicesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-site-not-assigned-network-devices-v1
*/
func (s *SiteDesignService) GetSiteNotAssignedNetworkDevicesV1(GetSiteNotAssignedNetworkDevicesV1QueryParams *GetSiteNotAssignedNetworkDevicesV1QueryParams) (*ResponseSiteDesignGetSiteNotAssignedNetworkDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/notAssignedToSite"

	queryString, _ := query.Values(GetSiteNotAssignedNetworkDevicesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSiteDesignGetSiteNotAssignedNetworkDevicesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSiteNotAssignedNetworkDevicesV1(GetSiteNotAssignedNetworkDevicesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSiteNotAssignedNetworkDevicesV1")
	}

	result := response.Result().(*ResponseSiteDesignGetSiteNotAssignedNetworkDevicesV1)
	return result, response, err

}

//GetSiteNotAssignedNetworkDevicesCountV1 Get site not assigned network devices count - b28e-881d-4cda-b06c
/* Get network devices count that are not assigned to any site.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-site-not-assigned-network-devices-count-v1
*/
func (s *SiteDesignService) GetSiteNotAssignedNetworkDevicesCountV1() (*ResponseSiteDesignGetSiteNotAssignedNetworkDevicesCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/notAssignedToSite/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignGetSiteNotAssignedNetworkDevicesCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSiteNotAssignedNetworkDevicesCountV1()
		}
		return nil, response, fmt.Errorf("error with operation GetSiteNotAssignedNetworkDevicesCountV1")
	}

	result := response.Result().(*ResponseSiteDesignGetSiteNotAssignedNetworkDevicesCountV1)
	return result, response, err

}

//GetSiteAssignedNetworkDeviceV1 Get site assigned network device - f08f-3b31-4bda-9c96
/* Get site assigned network device. The items in the list are arranged in an order that corresponds with their internal identifiers.


@param id id path parameter. Network Device Id.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-site-assigned-network-device-v1
*/
func (s *SiteDesignService) GetSiteAssignedNetworkDeviceV1(id string) (*ResponseSiteDesignGetSiteAssignedNetworkDeviceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/{id}/assignedToSite"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignGetSiteAssignedNetworkDeviceV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSiteAssignedNetworkDeviceV1(id)
		}
		return nil, response, fmt.Errorf("error with operation GetSiteAssignedNetworkDeviceV1")
	}

	result := response.Result().(*ResponseSiteDesignGetSiteAssignedNetworkDeviceV1)
	return result, response, err

}

//RetrievesTheListOfNetworkProfilesForSitesV1 Retrieves the list of network profiles for sites - a78d-8918-4898-9cf2
/* Retrieves the list of network profiles for sites.


@param RetrievesTheListOfNetworkProfilesForSitesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-list-of-network-profiles-for-sites-v1
*/
func (s *SiteDesignService) RetrievesTheListOfNetworkProfilesForSitesV1(RetrievesTheListOfNetworkProfilesForSitesV1QueryParams *RetrievesTheListOfNetworkProfilesForSitesV1QueryParams) (*ResponseSiteDesignRetrievesTheListOfNetworkProfilesForSitesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkProfilesForSites"

	queryString, _ := query.Values(RetrievesTheListOfNetworkProfilesForSitesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSiteDesignRetrievesTheListOfNetworkProfilesForSitesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheListOfNetworkProfilesForSitesV1(RetrievesTheListOfNetworkProfilesForSitesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheListOfNetworkProfilesForSitesV1")
	}

	result := response.Result().(*ResponseSiteDesignRetrievesTheListOfNetworkProfilesForSitesV1)
	return result, response, err

}

//RetrievesTheCountOfNetworkProfilesForSitesV1 Retrieves the count of network profiles for sites - 57a7-c9d2-4f4a-b000
/* Retrieves the count of network profiles for sites


@param RetrievesTheCountOfNetworkProfilesForSitesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-count-of-network-profiles-for-sites-v1
*/
func (s *SiteDesignService) RetrievesTheCountOfNetworkProfilesForSitesV1(RetrievesTheCountOfNetworkProfilesForSitesV1QueryParams *RetrievesTheCountOfNetworkProfilesForSitesV1QueryParams) (*ResponseSiteDesignRetrievesTheCountOfNetworkProfilesForSitesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkProfilesForSites/count"

	queryString, _ := query.Values(RetrievesTheCountOfNetworkProfilesForSitesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSiteDesignRetrievesTheCountOfNetworkProfilesForSitesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheCountOfNetworkProfilesForSitesV1(RetrievesTheCountOfNetworkProfilesForSitesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheCountOfNetworkProfilesForSitesV1")
	}

	result := response.Result().(*ResponseSiteDesignRetrievesTheCountOfNetworkProfilesForSitesV1)
	return result, response, err

}

//RetrieveANetworkProfileForSitesByIDV1 Retrieve a network profile for sites by id - 87a8-eb79-4f28-acc4
/* Retrieves a network profile for sites by id.


@param id id path parameter. The `id` of the network profile, retrievable from `GET /intent/api/v1/networkProfilesForSites`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-a-network-profile-for-sites-by-id-v1
*/
func (s *SiteDesignService) RetrieveANetworkProfileForSitesByIDV1(id string) (*ResponseSiteDesignRetrieveANetworkProfileForSitesByIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkProfilesForSites/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignRetrieveANetworkProfileForSitesByIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveANetworkProfileForSitesByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveANetworkProfileForSitesByIdV1")
	}

	result := response.Result().(*ResponseSiteDesignRetrieveANetworkProfileForSitesByIDV1)
	return result, response, err

}

//RetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToV1 Retrieves the list of sites that the given network profile for sites is assigned to - 0f84-ba73-4429-accd
/* Retrieves the list of sites that the given network profile for sites is assigned to.
The list includes the sites the profile has been directly assigned to, as well as child sites that have inherited the profile.


@param profileID profileId path parameter. The `id` of the network profile, retrievable from `GET /intent/api/v1/networkProfilesForSites`

@param RetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-list-of-sites-that-the-given-network-profile-for-sites-is-assigned-to-v1
*/
func (s *SiteDesignService) RetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToV1(profileID string, RetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToV1QueryParams *RetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToV1QueryParams) (*ResponseSiteDesignRetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkProfilesForSites/{profileId}/siteAssignments"
	path = strings.Replace(path, "{profileId}", fmt.Sprintf("%v", profileID), -1)

	queryString, _ := query.Values(RetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSiteDesignRetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToV1(profileID, RetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToV1")
	}

	result := response.Result().(*ResponseSiteDesignRetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToV1)
	return result, response, err

}

//RetrievesTheCountOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToV1 Retrieves the count of sites that the given network profile for sites is assigned to - 53ba-29dd-42eb-bd19
/* Retrieves the count of sites that the given network profile for sites is assigned to.


@param profileID profileId path parameter. The `id` of the network profile, retrievable from `GET /intent/api/v1/networkProfilesForSites`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-count-of-sites-that-the-given-network-profile-for-sites-is-assigned-to-v1
*/
func (s *SiteDesignService) RetrievesTheCountOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToV1(profileID string) (*ResponseSiteDesignRetrievesTheCountOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkProfilesForSites/{profileId}/siteAssignments/count"
	path = strings.Replace(path, "{profileId}", fmt.Sprintf("%v", profileID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignRetrievesTheCountOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheCountOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToV1(profileID)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheCountOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToV1")
	}

	result := response.Result().(*ResponseSiteDesignRetrievesTheCountOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToV1)
	return result, response, err

}

//GetSitesV1 Get sites - 4e8a-49c3-4b49-b291
/* Get sites.


@param GetSitesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-sites-v1
*/
func (s *SiteDesignService) GetSitesV1(GetSitesV1QueryParams *GetSitesV1QueryParams) (*ResponseSiteDesignGetSitesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites"

	queryString, _ := query.Values(GetSitesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSiteDesignGetSitesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSitesV1(GetSitesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSitesV1")
	}

	result := response.Result().(*ResponseSiteDesignGetSitesV1)
	return result, response, err

}

//GetSitesCountV1 Get sites count - 0fbf-482e-446a-835f
/* Get sites count.


@param GetSitesCountV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-sites-count-v1
*/
func (s *SiteDesignService) GetSitesCountV1(GetSitesCountV1QueryParams *GetSitesCountV1QueryParams) (*ResponseSiteDesignGetSitesCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/count"

	queryString, _ := query.Values(GetSitesCountV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSiteDesignGetSitesCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSitesCountV1(GetSitesCountV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSitesCountV1")
	}

	result := response.Result().(*ResponseSiteDesignGetSitesCountV1)
	return result, response, err

}

//RetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssignedV1 Retrieves the list of network profiles that the given site has been assigned - b0b4-5962-49d9-8d6b
/* Retrieves the list of profiles that the given site has been assigned.  These profiles may either be directly assigned to this site, or were assigned to a parent site and have been inherited.
These assigments can be modified via the `/dna/intent/api/v1/networkProfilesForSites/{profileId}/siteAssignments` resources.


@param siteID siteId path parameter. The `id` of the site, retrievable from `/dna/intent/api/v1/sites`

@param RetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssignedV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-list-of-network-profiles-that-the-given-site-has-been-assigned-v1
*/
func (s *SiteDesignService) RetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssignedV1(siteID string, RetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssignedV1QueryParams *RetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssignedV1QueryParams) (*ResponseSiteDesignRetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssignedV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{siteId}/profileAssignments"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	queryString, _ := query.Values(RetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssignedV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSiteDesignRetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssignedV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssignedV1(siteID, RetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssignedV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssignedV1")
	}

	result := response.Result().(*ResponseSiteDesignRetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssignedV1)
	return result, response, err

}

//RetrievesTheCountOfProfilesThatTheGivenSiteHasBeenAssignedV1 Retrieves the count of profiles that the given site has been assigned - 28be-9a3f-4688-a2d4
/* Retrieves the count of profiles that the given site has been assigned.  These profiles may either be directly assigned to this site, or were assigned to a parent site and have been inherited.


@param siteID siteId path parameter. The `id` of the site, retrievable from `/dna/intent/api/v1/sites`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-count-of-profiles-that-the-given-site-has-been-assigned-v1
*/
func (s *SiteDesignService) RetrievesTheCountOfProfilesThatTheGivenSiteHasBeenAssignedV1(siteID string) (*ResponseSiteDesignRetrievesTheCountOfProfilesThatTheGivenSiteHasBeenAssignedV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{siteId}/profileAssignments/count"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignRetrievesTheCountOfProfilesThatTheGivenSiteHasBeenAssignedV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheCountOfProfilesThatTheGivenSiteHasBeenAssignedV1(siteID)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheCountOfProfilesThatTheGivenSiteHasBeenAssignedV1")
	}

	result := response.Result().(*ResponseSiteDesignRetrievesTheCountOfProfilesThatTheGivenSiteHasBeenAssignedV1)
	return result, response, err

}

//GetsABuildingV2 Gets a building - e293-295e-4a78-bf64
/* Gets a building in the network hierarchy.


@param id id path parameter. Building Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-a-building-v2
*/
func (s *SiteDesignService) GetsABuildingV2(id string) (*ResponseSiteDesignGetsABuildingV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/buildings/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignGetsABuildingV2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsABuildingV2(id)
		}
		return nil, response, fmt.Errorf("error with operation GetsABuildingV2")
	}

	result := response.Result().(*ResponseSiteDesignGetsABuildingV2)
	return result, response, err

}

//GetFloorSettingsV2 Get floor settings - f697-a95f-4469-958f
/* Gets UI user preference for floor unit system.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-floor-settings-v2
*/
func (s *SiteDesignService) GetFloorSettingsV2() (*ResponseSiteDesignGetFloorSettingsV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/floors/settings"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignGetFloorSettingsV2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFloorSettingsV2()
		}
		return nil, response, fmt.Errorf("error with operation GetFloorSettingsV2")
	}

	result := response.Result().(*ResponseSiteDesignGetFloorSettingsV2)
	return result, response, err

}

//GetsAFloorV2 Gets a floor - ff92-2958-4bba-9288
/* Gets a floor in the network hierarchy.


@param id id path parameter. Floor Id

@param GetsAFloorV2QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-a-floor-v2
*/
func (s *SiteDesignService) GetsAFloorV2(id string, GetsAFloorV2QueryParams *GetsAFloorV2QueryParams) (*ResponseSiteDesignGetsAFloorV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/floors/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetsAFloorV2QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSiteDesignGetsAFloorV2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsAFloorV2(id, GetsAFloorV2QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetsAFloorV2")
	}

	result := response.Result().(*ResponseSiteDesignGetsAFloorV2)
	return result, response, err

}

//CreatesAnAreaV1 Creates an area - a8bc-d9dc-43ea-a7e3
/* Creates an area in the network hierarchy.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!creates-an-area-v1
*/
func (s *SiteDesignService) CreatesAnAreaV1(requestSiteDesignCreatesAnAreaV1 *RequestSiteDesignCreatesAnAreaV1) (*ResponseSiteDesignCreatesAnAreaV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/areas"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignCreatesAnAreaV1).
		SetResult(&ResponseSiteDesignCreatesAnAreaV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreatesAnAreaV1(requestSiteDesignCreatesAnAreaV1)
		}

		return nil, response, fmt.Errorf("error with operation CreatesAnAreaV1")
	}

	result := response.Result().(*ResponseSiteDesignCreatesAnAreaV1)
	return result, response, err

}

//AssignNetworkDevicesToASiteV1 Assign network devices to a site - a1b4-8949-4a6a-a6ac
/* Assign unprovisioned network devices to a site. Along with that it can also be used to assign unprovisioned network devices to a different site. If device controllability is enabled, it will be triggered once device assigned to site successfully. Device Controllability can be enabled/disabled using `/dna/intent/api/v1/networkDevices/deviceControllability/settings`.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!assign-network-devices-to-a-site-v1
*/
func (s *SiteDesignService) AssignNetworkDevicesToASiteV1(requestSiteDesignAssignNetworkDevicesToASiteV1 *RequestSiteDesignAssignNetworkDevicesToASiteV1) (*ResponseSiteDesignAssignNetworkDevicesToASiteV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/assignToSite/apply"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignAssignNetworkDevicesToASiteV1).
		SetResult(&ResponseSiteDesignAssignNetworkDevicesToASiteV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AssignNetworkDevicesToASiteV1(requestSiteDesignAssignNetworkDevicesToASiteV1)
		}

		return nil, response, fmt.Errorf("error with operation AssignNetworkDevicesToASiteV1")
	}

	result := response.Result().(*ResponseSiteDesignAssignNetworkDevicesToASiteV1)
	return result, response, err

}

//UnassignNetworkDevicesFromSitesV1 Unassign network devices from sites - 08a6-1a87-44eb-8606
/* Unassign unprovisioned network devices from their site. If device controllability is enabled, it will be triggered once device unassigned from site successfully. Device Controllability can be enabled/disabled using `/dna/intent/api/v1/networkDevices/deviceControllability/settings`.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!unassign-network-devices-from-sites-v1
*/
func (s *SiteDesignService) UnassignNetworkDevicesFromSitesV1(requestSiteDesignUnassignNetworkDevicesFromSitesV1 *RequestSiteDesignUnassignNetworkDevicesFromSitesV1) (*ResponseSiteDesignUnassignNetworkDevicesFromSitesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/unassignFromSite/apply"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignUnassignNetworkDevicesFromSitesV1).
		SetResult(&ResponseSiteDesignUnassignNetworkDevicesFromSitesV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.UnassignNetworkDevicesFromSitesV1(requestSiteDesignUnassignNetworkDevicesFromSitesV1)
		}

		return nil, response, fmt.Errorf("error with operation UnassignNetworkDevicesFromSitesV1")
	}

	result := response.Result().(*ResponseSiteDesignUnassignNetworkDevicesFromSitesV1)
	return result, response, err

}

//AssignANetworkProfileForSitesToTheGivenSiteV1 Assign a network profile for sites to the given site - 40ba-89da-420a-a21b
/* Assigns a given network profile for sites to a given site. Also assigns the profile to child sites.


@param profileID profileId path parameter. The `id` of the network profile, retrievable from `GET /intent/api/v1/networkProfilesForSites`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!assign-a-network-profile-for-sites-to-the-given-site-v1
*/
func (s *SiteDesignService) AssignANetworkProfileForSitesToTheGivenSiteV1(profileID string, requestSiteDesignAssignANetworkProfileForSitesToTheGivenSiteV1 *RequestSiteDesignAssignANetworkProfileForSitesToTheGivenSiteV1) (*ResponseSiteDesignAssignANetworkProfileForSitesToTheGivenSiteV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkProfilesForSites/{profileId}/siteAssignments"
	path = strings.Replace(path, "{profileId}", fmt.Sprintf("%v", profileID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignAssignANetworkProfileForSitesToTheGivenSiteV1).
		SetResult(&ResponseSiteDesignAssignANetworkProfileForSitesToTheGivenSiteV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AssignANetworkProfileForSitesToTheGivenSiteV1(profileID, requestSiteDesignAssignANetworkProfileForSitesToTheGivenSiteV1)
		}

		return nil, response, fmt.Errorf("error with operation AssignANetworkProfileForSitesToTheGivenSiteV1")
	}

	result := response.Result().(*ResponseSiteDesignAssignANetworkProfileForSitesToTheGivenSiteV1)
	return result, response, err

}

//AssignANetworkProfileForSitesToAListOfSitesV1 Assign a network profile for sites to a list of sites - 6ab6-e992-451b-8bc9
/* Assign a network profile for sites to a list of sites. Also assigns the profile to child sites.


@param profileID profileId path parameter. The `id` of the network profile, retrievable from `GET /intent/api/v1/networkProfilesForSites`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!assign-a-network-profile-for-sites-to-a-list-of-sites-v1
*/
func (s *SiteDesignService) AssignANetworkProfileForSitesToAListOfSitesV1(profileID string, requestSiteDesignAssignANetworkProfileForSitesToAListOfSitesV1 *RequestSiteDesignAssignANetworkProfileForSitesToAListOfSitesV1) (*ResponseSiteDesignAssignANetworkProfileForSitesToAListOfSitesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkProfilesForSites/{profileId}/siteAssignments/bulk"
	path = strings.Replace(path, "{profileId}", fmt.Sprintf("%v", profileID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignAssignANetworkProfileForSitesToAListOfSitesV1).
		SetResult(&ResponseSiteDesignAssignANetworkProfileForSitesToAListOfSitesV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AssignANetworkProfileForSitesToAListOfSitesV1(profileID, requestSiteDesignAssignANetworkProfileForSitesToAListOfSitesV1)
		}

		return nil, response, fmt.Errorf("error with operation AssignANetworkProfileForSitesToAListOfSitesV1")
	}

	result := response.Result().(*ResponseSiteDesignAssignANetworkProfileForSitesToAListOfSitesV1)
	return result, response, err

}

//AssociateV1 Associate - 308e-195d-403a-bbd4
/* Associate Site to a Network Profile


@param networkProfileID networkProfileId path parameter. Network-Profile Id to be associated

@param siteID siteId path parameter. Site Id to be associated


Documentation Link: https://developer.cisco.com/docs/dna-center/#!associate-v1
*/
func (s *SiteDesignService) AssociateV1(networkProfileID string, siteID string) (*ResponseSiteDesignAssociateV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkprofile/{networkProfileId}/site/{siteId}"
	path = strings.Replace(path, "{networkProfileId}", fmt.Sprintf("%v", networkProfileID), -1)
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignAssociateV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AssociateV1(networkProfileID, siteID)
		}

		return nil, response, fmt.Errorf("error with operation AssociateV1")
	}

	result := response.Result().(*ResponseSiteDesignAssociateV1)
	return result, response, err

}

//CreateSitesV1 Create sites - efac-69a1-4c2a-9d5e
/* Create area/building/floor together in bulk. If site already exist, then that will be ignored. Sites in the request payload need not to be ordered.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-sites-v1
*/
func (s *SiteDesignService) CreateSitesV1(requestSiteDesignCreateSitesV1 *RequestSiteDesignCreateSitesV1) (*ResponseSiteDesignCreateSitesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/bulk"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignCreateSitesV1).
		SetResult(&ResponseSiteDesignCreateSitesV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateSitesV1(requestSiteDesignCreateSitesV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateSitesV1")
	}

	result := response.Result().(*ResponseSiteDesignCreateSitesV1)
	return result, response, err

}

//CreatesABuildingV2 Creates a building - 73ae-3922-466b-adb1
/* Creates a building in the network hierarchy under area.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!creates-a-building-v2
*/
func (s *SiteDesignService) CreatesABuildingV2(requestSiteDesignCreatesABuildingV2 *RequestSiteDesignCreatesABuildingV2) (*ResponseSiteDesignCreatesABuildingV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/buildings"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignCreatesABuildingV2).
		SetResult(&ResponseSiteDesignCreatesABuildingV2{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreatesABuildingV2(requestSiteDesignCreatesABuildingV2)
		}

		return nil, response, fmt.Errorf("error with operation CreatesABuildingV2")
	}

	result := response.Result().(*ResponseSiteDesignCreatesABuildingV2)
	return result, response, err

}

//CreatesAFloorV2 Creates a floor - 8882-b8fb-450a-8528
/* Create a floor in the network hierarchy under building.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!creates-a-floor-v2
*/
func (s *SiteDesignService) CreatesAFloorV2(requestSiteDesignCreatesAFloorV2 *RequestSiteDesignCreatesAFloorV2) (*ResponseSiteDesignCreatesAFloorV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/floors"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignCreatesAFloorV2).
		SetResult(&ResponseSiteDesignCreatesAFloorV2{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreatesAFloorV2(requestSiteDesignCreatesAFloorV2)
		}

		return nil, response, fmt.Errorf("error with operation CreatesAFloorV2")
	}

	result := response.Result().(*ResponseSiteDesignCreatesAFloorV2)
	return result, response, err

}

//UploadsFloorImageV2 Uploads floor image - fca4-5804-4758-98d2
/* Uploads floor image.


@param id id path parameter. Floor Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!uploads-floor-image-v2
*/
func (s *SiteDesignService) UploadsFloorImageV2(id string) (*resty.Response, error) {
	path := "/dna/intent/api/v2/floors/{id}/uploadImage"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

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
			return s.UploadsFloorImageV2(id)
		}

		return response, fmt.Errorf("error with operation UploadsFloorImageV2")
	}

	return response, err

}

//UpdatesAnAreaV1 Updates an area - fab7-a965-4599-885f
/* Updates an area in the network hierarchy.


@param id id path parameter. Area Id

*/
func (s *SiteDesignService) UpdatesAnAreaV1(id string, requestSiteDesignUpdatesAnAreaV1 *RequestSiteDesignUpdatesAnAreaV1) (*ResponseSiteDesignUpdatesAnAreaV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/areas/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignUpdatesAnAreaV1).
		SetResult(&ResponseSiteDesignUpdatesAnAreaV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdatesAnAreaV1(id, requestSiteDesignUpdatesAnAreaV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdatesAnAreaV1")
	}

	result := response.Result().(*ResponseSiteDesignUpdatesAnAreaV1)
	return result, response, err

}

//UpdateDeviceControllabilitySettingsV1 Update device controllability settings - f58f-e8d6-4b88-912c
/* Device Controllability is a system-level process on Catalyst Center that enforces state synchronization for some device-layer features. Its purpose is to aid in the deployment of required network settings that Catalyst Center needs to manage devices. Changes are made on network devices  during discovery, when adding a device to Inventory, or when assigning a device to a site. If changes  are made to any settings that are under the scope of this process, these changes are applied to the  network devices during the Provision and Update Telemetry Settings operations, even if Device  Controllability is disabled. The following device settings will be enabled as part of  Device Controllability when devices are discovered.

  SNMP Credentials.
  NETCONF Credentials.

Subsequent to discovery, devices will be added to Inventory. The following device settings will be  enabled when devices are added to inventory.

  Cisco TrustSec (CTS) Credentials.

The following device settings will be enabled when devices are assigned to a site. Some of these  settings can be defined at a site level under Design > Network Settings > Telemetry & Wireless.

  Wired Endpoint Data Collection Enablement.
  Controller Certificates.
  SNMP Trap Server Definitions.
  Syslog Server Definitions.
  Application Visibility.
  Application QoS Policy.
  Wireless Service Assurance (WSA).
  Wireless Telemetry.
  DTLS Ciphersuite.
  AP Impersonation.

If Device Controllability is disabled, Catalyst Center does not configure any of the preceding  credentials or settings on devices during discovery, at runtime, or during site assignment. However,  the telemetry settings and related configuration are pushed when the device is provisioned or when the  update Telemetry Settings action is performed.
Catalyst Center identifies and automatically corrects the following telemetry configuration issues on  the device.

  SWIM certificate issue.
  IOS WLC NA certificate issue.
  PKCS12 certificate issue.
  IOS telemetry configuration issue.

The autocorrect telemetry config feature is supported only when Device Controllability is enabled.


*/
func (s *SiteDesignService) UpdateDeviceControllabilitySettingsV1(requestSiteDesignUpdateDeviceControllabilitySettingsV1 *RequestSiteDesignUpdateDeviceControllabilitySettingsV1) (*ResponseSiteDesignUpdateDeviceControllabilitySettingsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/deviceControllability/settings"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignUpdateDeviceControllabilitySettingsV1).
		SetResult(&ResponseSiteDesignUpdateDeviceControllabilitySettingsV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateDeviceControllabilitySettingsV1(requestSiteDesignUpdateDeviceControllabilitySettingsV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateDeviceControllabilitySettingsV1")
	}

	result := response.Result().(*ResponseSiteDesignUpdateDeviceControllabilitySettingsV1)
	return result, response, err

}

//UpdatesABuildingV2 Updates a building - 0280-1b95-4d78-845d
/* Updates a building in the network hierarchy.


@param id id path parameter. Building Id

*/
func (s *SiteDesignService) UpdatesABuildingV2(id string, requestSiteDesignUpdatesABuildingV2 *RequestSiteDesignUpdatesABuildingV2) (*ResponseSiteDesignUpdatesABuildingV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/buildings/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignUpdatesABuildingV2).
		SetResult(&ResponseSiteDesignUpdatesABuildingV2{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdatesABuildingV2(id, requestSiteDesignUpdatesABuildingV2)
		}
		return nil, response, fmt.Errorf("error with operation UpdatesABuildingV2")
	}

	result := response.Result().(*ResponseSiteDesignUpdatesABuildingV2)
	return result, response, err

}

//UpdatesFloorSettingsV2 Updates floor settings - 16b8-59a5-4b38-8cb7
/* Updates UI user preference for floor unit system. Unit sytem change will effect for all floors across all sites.


 */
func (s *SiteDesignService) UpdatesFloorSettingsV2(requestSiteDesignUpdatesFloorSettingsV2 *RequestSiteDesignUpdatesFloorSettingsV2) (*ResponseSiteDesignUpdatesFloorSettingsV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/floors/settings"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignUpdatesFloorSettingsV2).
		SetResult(&ResponseSiteDesignUpdatesFloorSettingsV2{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdatesFloorSettingsV2(requestSiteDesignUpdatesFloorSettingsV2)
		}
		return nil, response, fmt.Errorf("error with operation UpdatesFloorSettingsV2")
	}

	result := response.Result().(*ResponseSiteDesignUpdatesFloorSettingsV2)
	return result, response, err

}

//UpdatesAFloorV2 Updates a floor - ee9f-c9f2-4a3a-b7a5
/* Updates a floor in the network hierarchy.


@param id id path parameter. Floor Id

*/
func (s *SiteDesignService) UpdatesAFloorV2(id string, requestSiteDesignUpdatesAFloorV2 *RequestSiteDesignUpdatesAFloorV2) (*ResponseSiteDesignUpdatesAFloorV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/floors/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignUpdatesAFloorV2).
		SetResult(&ResponseSiteDesignUpdatesAFloorV2{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdatesAFloorV2(id, requestSiteDesignUpdatesAFloorV2)
		}
		return nil, response, fmt.Errorf("error with operation UpdatesAFloorV2")
	}

	result := response.Result().(*ResponseSiteDesignUpdatesAFloorV2)
	return result, response, err

}

//DeletesAnAreaV1 Deletes an area - 2ba4-c8f4-4ec8-b80e
/* Deletes an area in the network hierarchy. This operations fails if there are any child areas or buildings for this area.


@param id id path parameter. Area ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!deletes-an-area-v1
*/
func (s *SiteDesignService) DeletesAnAreaV1(id string) (*ResponseSiteDesignDeletesAnAreaV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/areas/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignDeletesAnAreaV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletesAnAreaV1(id)
		}
		return nil, response, fmt.Errorf("error with operation DeletesAnAreaV1")
	}

	result := response.Result().(*ResponseSiteDesignDeletesAnAreaV1)
	return result, response, err

}

//DeletesANetworkProfileForSitesV1 Deletes a network profile for sites - a48b-b959-493b-93d1
/* Deletes a network profile for sites.


@param id id path parameter. The `id` of the network profile, retrievable from `GET /intent/api/v1/networkProfilesForSites`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!deletes-a-network-profile-for-sites-v1
*/
func (s *SiteDesignService) DeletesANetworkProfileForSitesV1(id string) (*ResponseSiteDesignDeletesANetworkProfileForSitesV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/networkProfilesForSites/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignDeletesANetworkProfileForSitesV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletesANetworkProfileForSitesV1(id)
		}
		return nil, response, fmt.Errorf("error with operation DeletesANetworkProfileForSitesV1")
	}

	result := response.Result().(*ResponseSiteDesignDeletesANetworkProfileForSitesV1)
	return result, response, err

}

//UnassignsANetworkProfileForSitesFromMultipleSitesV1 Unassigns a network profile for sites from multiple sites - ec90-f973-435b-a6f5
/* Unassigns a given network profile for sites from multiple sites. The profile must be removed from the containing building first if this site is a floor.


@param profileID profileId path parameter. The `id` of the network profile, retrievable from `GET /intent/api/v1/networkProfilesForSites`

@param UnassignsANetworkProfileForSitesFromMultipleSitesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!unassigns-a-network-profile-for-sites-from-multiple-sites-v1
*/
func (s *SiteDesignService) UnassignsANetworkProfileForSitesFromMultipleSitesV1(profileID string, UnassignsANetworkProfileForSitesFromMultipleSitesV1QueryParams *UnassignsANetworkProfileForSitesFromMultipleSitesV1QueryParams) (*ResponseSiteDesignUnassignsANetworkProfileForSitesFromMultipleSitesV1, *resty.Response, error) {
	//profileID string,UnassignsANetworkProfileForSitesFromMultipleSitesV1QueryParams *UnassignsANetworkProfileForSitesFromMultipleSitesV1QueryParams
	path := "/dna/intent/api/v1/networkProfilesForSites/{profileId}/siteAssignments/bulk"
	path = strings.Replace(path, "{profileId}", fmt.Sprintf("%v", profileID), -1)

	queryString, _ := query.Values(UnassignsANetworkProfileForSitesFromMultipleSitesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSiteDesignUnassignsANetworkProfileForSitesFromMultipleSitesV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UnassignsANetworkProfileForSitesFromMultipleSitesV1(profileID, UnassignsANetworkProfileForSitesFromMultipleSitesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation UnassignsANetworkProfileForSitesFromMultipleSitesV1")
	}

	result := response.Result().(*ResponseSiteDesignUnassignsANetworkProfileForSitesFromMultipleSitesV1)
	return result, response, err

}

//UnassignsANetworkProfileForSitesFromASiteV1 Unassigns a network profile for sites from a site - 8c94-0a8e-450b-98cc
/* Unassigns a given network profile for sites from a site. The profile must be removed from parent sites first, otherwise this operation will not ulimately  unassign the profile.


@param profileID profileId path parameter. The `id` of the network profile, retrievable from `GET /intent/api/v1/networkProfilesForSites`

@param id id path parameter. The `id` of the site, retrievable from `GET /intent/api/v1/networkProfilesForSites/{id}/siteAssignments`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!unassigns-a-network-profile-for-sites-from-a-site-v1
*/
func (s *SiteDesignService) UnassignsANetworkProfileForSitesFromASiteV1(profileID string, id string) (*ResponseSiteDesignUnassignsANetworkProfileForSitesFromASiteV1, *resty.Response, error) {
	//profileID string,id string
	path := "/dna/intent/api/v1/networkProfilesForSites/{profileId}/siteAssignments/{id}"
	path = strings.Replace(path, "{profileId}", fmt.Sprintf("%v", profileID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignUnassignsANetworkProfileForSitesFromASiteV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UnassignsANetworkProfileForSitesFromASiteV1(profileID, id)
		}
		return nil, response, fmt.Errorf("error with operation UnassignsANetworkProfileForSitesFromASiteV1")
	}

	result := response.Result().(*ResponseSiteDesignUnassignsANetworkProfileForSitesFromASiteV1)
	return result, response, err

}

//DisassociateV1 Disassociate - e687-58d2-4b19-b5c6
/* Disassociate a Site from a Network Profile


@param networkProfileID networkProfileId path parameter. Network-Profile Id to be associated

@param siteID siteId path parameter. Site Id to be associated


Documentation Link: https://developer.cisco.com/docs/dna-center/#!disassociate-v1
*/
func (s *SiteDesignService) DisassociateV1(networkProfileID string, siteID string) (*ResponseSiteDesignDisassociateV1, *resty.Response, error) {
	//networkProfileID string,siteID string
	path := "/dna/intent/api/v1/networkprofile/{networkProfileId}/site/{siteId}"
	path = strings.Replace(path, "{networkProfileId}", fmt.Sprintf("%v", networkProfileID), -1)
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignDisassociateV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DisassociateV1(networkProfileID, siteID)
		}
		return nil, response, fmt.Errorf("error with operation DisassociateV1")
	}

	result := response.Result().(*ResponseSiteDesignDisassociateV1)
	return result, response, err

}

//DeletesABuildingV2 Deletes a building - 45ae-a9e4-4008-b0b6
/* Deletes building in the network hierarchy. This operations fails if there are any floors for this building, or if there are any devices assigned to this building.


@param id id path parameter. Building ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!deletes-a-building-v2
*/
func (s *SiteDesignService) DeletesABuildingV2(id string) (*ResponseSiteDesignDeletesABuildingV2, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v2/buildings/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignDeletesABuildingV2{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletesABuildingV2(id)
		}
		return nil, response, fmt.Errorf("error with operation DeletesABuildingV2")
	}

	result := response.Result().(*ResponseSiteDesignDeletesABuildingV2)
	return result, response, err

}

//DeletesAFloorV2 Deletes a floor - 6cb4-884b-47db-a808
/* Deletes a floor from the network hierarchy. This operations fails if there are any devices assigned to this floor.


@param id id path parameter. Floor ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!deletes-a-floor-v2
*/
func (s *SiteDesignService) DeletesAFloorV2(id string) (*ResponseSiteDesignDeletesAFloorV2, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v2/floors/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignDeletesAFloorV2{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletesAFloorV2(id)
		}
		return nil, response, fmt.Errorf("error with operation DeletesAFloorV2")
	}

	result := response.Result().(*ResponseSiteDesignDeletesAFloorV2)
	return result, response, err

}

// Alias Function
/*
This method acts as an alias for the method `UpdatesFloorSettingsV2`
*/
func (s *SiteDesignService) UpdatesFloorSettings(requestSiteDesignUpdatesFloorSettingsV2 *RequestSiteDesignUpdatesFloorSettingsV2) (*ResponseSiteDesignUpdatesFloorSettingsV2, *resty.Response, error) {
	return s.UpdatesFloorSettingsV2(requestSiteDesignUpdatesFloorSettingsV2)
}

// Alias Function
/*
This method acts as an alias for the method `GetSitesCountV1`
*/
func (s *SiteDesignService) GetSitesCount(GetSitesCountV1QueryParams *GetSitesCountV1QueryParams) (*ResponseSiteDesignGetSitesCountV1, *resty.Response, error) {
	return s.GetSitesCountV1(GetSitesCountV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesTheCountOfProfilesThatTheGivenSiteHasBeenAssignedV1`
*/
func (s *SiteDesignService) RetrievesTheCountOfProfilesThatTheGivenSiteHasBeenAssigned(siteID string) (*ResponseSiteDesignRetrievesTheCountOfProfilesThatTheGivenSiteHasBeenAssignedV1, *resty.Response, error) {
	return s.RetrievesTheCountOfProfilesThatTheGivenSiteHasBeenAssignedV1(siteID)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesTheListOfNetworkProfilesForSitesV1`
*/
func (s *SiteDesignService) RetrievesTheListOfNetworkProfilesForSites(RetrievesTheListOfNetworkProfilesForSitesV1QueryParams *RetrievesTheListOfNetworkProfilesForSitesV1QueryParams) (*ResponseSiteDesignRetrievesTheListOfNetworkProfilesForSitesV1, *resty.Response, error) {
	return s.RetrievesTheListOfNetworkProfilesForSitesV1(RetrievesTheListOfNetworkProfilesForSitesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetSiteNotAssignedNetworkDevicesV1`
*/
func (s *SiteDesignService) GetSiteNotAssignedNetworkDevices(GetSiteNotAssignedNetworkDevicesV1QueryParams *GetSiteNotAssignedNetworkDevicesV1QueryParams) (*ResponseSiteDesignGetSiteNotAssignedNetworkDevicesV1, *resty.Response, error) {
	return s.GetSiteNotAssignedNetworkDevicesV1(GetSiteNotAssignedNetworkDevicesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetSiteNotAssignedNetworkDevicesCountV1`
*/
func (s *SiteDesignService) GetSiteNotAssignedNetworkDevicesCount() (*ResponseSiteDesignGetSiteNotAssignedNetworkDevicesCountV1, *resty.Response, error) {
	return s.GetSiteNotAssignedNetworkDevicesCountV1()
}

// Alias Function
/*
This method acts as an alias for the method `CreatesAFloorV2`
*/
func (s *SiteDesignService) CreatesAFloor(requestSiteDesignCreatesAFloorV2 *RequestSiteDesignCreatesAFloorV2) (*ResponseSiteDesignCreatesAFloorV2, *resty.Response, error) {
	return s.CreatesAFloorV2(requestSiteDesignCreatesAFloorV2)
}

// Alias Function
/*
This method acts as an alias for the method `UpdatesAnAreaV1`
*/
func (s *SiteDesignService) UpdatesAnArea(id string, requestSiteDesignUpdatesAnAreaV1 *RequestSiteDesignUpdatesAnAreaV1) (*ResponseSiteDesignUpdatesAnAreaV1, *resty.Response, error) {
	return s.UpdatesAnAreaV1(id, requestSiteDesignUpdatesAnAreaV1)
}

// Alias Function
/*
This method acts as an alias for the method `UploadsFloorImageV2`
*/
func (s *SiteDesignService) UploadsFloorImage(id string) (*resty.Response, error) {
	return s.UploadsFloorImageV2(id)
}

// Alias Function
/*
This method acts as an alias for the method `CreatesAnAreaV1`
*/
func (s *SiteDesignService) CreatesAnArea(requestSiteDesignCreatesAnAreaV1 *RequestSiteDesignCreatesAnAreaV1) (*ResponseSiteDesignCreatesAnAreaV1, *resty.Response, error) {
	return s.CreatesAnAreaV1(requestSiteDesignCreatesAnAreaV1)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesTheCountOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToV1`
*/
func (s *SiteDesignService) RetrievesTheCountOfSitesThatTheGivenNetworkProfileForSitesIsAssignedTo(profileID string) (*ResponseSiteDesignRetrievesTheCountOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToV1, *resty.Response, error) {
	return s.RetrievesTheCountOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToV1(profileID)
}

// Alias Function
/*
This method acts as an alias for the method `UpdatesAFloorV2`
*/
func (s *SiteDesignService) UpdatesAFloor(id string, requestSiteDesignUpdatesAFloorV2 *RequestSiteDesignUpdatesAFloorV2) (*ResponseSiteDesignUpdatesAFloorV2, *resty.Response, error) {
	return s.UpdatesAFloorV2(id, requestSiteDesignUpdatesAFloorV2)
}

// Alias Function
/*
This method acts as an alias for the method `DeletesANetworkProfileForSitesV1`
*/
func (s *SiteDesignService) DeletesANetworkProfileForSites(id string) (*ResponseSiteDesignDeletesANetworkProfileForSitesV1, *resty.Response, error) {
	return s.DeletesANetworkProfileForSitesV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `GetSiteAssignedNetworkDeviceV1`
*/
func (s *SiteDesignService) GetSiteAssignedNetworkDevice(id string) (*ResponseSiteDesignGetSiteAssignedNetworkDeviceV1, *resty.Response, error) {
	return s.GetSiteAssignedNetworkDeviceV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `AssignANetworkProfileForSitesToTheGivenSiteV1`
*/
func (s *SiteDesignService) AssignANetworkProfileForSitesToTheGivenSite(profileID string, requestSiteDesignAssignANetworkProfileForSitesToTheGivenSiteV1 *RequestSiteDesignAssignANetworkProfileForSitesToTheGivenSiteV1) (*ResponseSiteDesignAssignANetworkProfileForSitesToTheGivenSiteV1, *resty.Response, error) {
	return s.AssignANetworkProfileForSitesToTheGivenSiteV1(profileID, requestSiteDesignAssignANetworkProfileForSitesToTheGivenSiteV1)
}

// Alias Function
/*
This method acts as an alias for the method `CreatesABuildingV2`
*/
func (s *SiteDesignService) CreatesABuilding(requestSiteDesignCreatesABuildingV2 *RequestSiteDesignCreatesABuildingV2) (*ResponseSiteDesignCreatesABuildingV2, *resty.Response, error) {
	return s.CreatesABuildingV2(requestSiteDesignCreatesABuildingV2)
}

// Alias Function
/*
This method acts as an alias for the method `DeletesAnAreaV1`
*/
func (s *SiteDesignService) DeletesAnArea(id string) (*ResponseSiteDesignDeletesAnAreaV1, *resty.Response, error) {
	return s.DeletesAnAreaV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `RetrieveANetworkProfileForSitesByIDV1`
*/
func (s *SiteDesignService) RetrieveANetworkProfileForSitesByID(id string) (*ResponseSiteDesignRetrieveANetworkProfileForSitesByIDV1, *resty.Response, error) {
	return s.RetrieveANetworkProfileForSitesByIDV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `GetsAFloorV2`
*/
func (s *SiteDesignService) GetsAFloor(id string, GetsAFloorV2QueryParams *GetsAFloorV2QueryParams) (*ResponseSiteDesignGetsAFloorV2, *resty.Response, error) {
	return s.GetsAFloorV2(id, GetsAFloorV2QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToV1`
*/
func (s *SiteDesignService) RetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedTo(profileID string, RetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToV1QueryParams *RetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToV1QueryParams) (*ResponseSiteDesignRetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToV1, *resty.Response, error) {
	return s.RetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToV1(profileID, RetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `AssignANetworkProfileForSitesToAListOfSitesV1`
*/
func (s *SiteDesignService) AssignANetworkProfileForSitesToAListOfSites(profileID string, requestSiteDesignAssignANetworkProfileForSitesToAListOfSitesV1 *RequestSiteDesignAssignANetworkProfileForSitesToAListOfSitesV1) (*ResponseSiteDesignAssignANetworkProfileForSitesToAListOfSitesV1, *resty.Response, error) {
	return s.AssignANetworkProfileForSitesToAListOfSitesV1(profileID, requestSiteDesignAssignANetworkProfileForSitesToAListOfSitesV1)
}

// Alias Function
/*
This method acts as an alias for the method `UpdatesABuildingV2`
*/
func (s *SiteDesignService) UpdatesABuilding(id string, requestSiteDesignUpdatesABuildingV2 *RequestSiteDesignUpdatesABuildingV2) (*ResponseSiteDesignUpdatesABuildingV2, *resty.Response, error) {
	return s.UpdatesABuildingV2(id, requestSiteDesignUpdatesABuildingV2)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesTheCountOfNetworkProfilesForSitesV1`
*/
func (s *SiteDesignService) RetrievesTheCountOfNetworkProfilesForSites(RetrievesTheCountOfNetworkProfilesForSitesV1QueryParams *RetrievesTheCountOfNetworkProfilesForSitesV1QueryParams) (*ResponseSiteDesignRetrievesTheCountOfNetworkProfilesForSitesV1, *resty.Response, error) {
	return s.RetrievesTheCountOfNetworkProfilesForSitesV1(RetrievesTheCountOfNetworkProfilesForSitesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `UnassignsANetworkProfileForSitesFromMultipleSitesV1`
*/
func (s *SiteDesignService) UnassignsANetworkProfileForSitesFromMultipleSites(profileID string, UnassignsANetworkProfileForSitesFromMultipleSitesV1QueryParams *UnassignsANetworkProfileForSitesFromMultipleSitesV1QueryParams) (*ResponseSiteDesignUnassignsANetworkProfileForSitesFromMultipleSitesV1, *resty.Response, error) {
	return s.UnassignsANetworkProfileForSitesFromMultipleSitesV1(profileID, UnassignsANetworkProfileForSitesFromMultipleSitesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `UnassignNetworkDevicesFromSitesV1`
*/
func (s *SiteDesignService) UnassignNetworkDevicesFromSites(requestSiteDesignUnassignNetworkDevicesFromSitesV1 *RequestSiteDesignUnassignNetworkDevicesFromSitesV1) (*ResponseSiteDesignUnassignNetworkDevicesFromSitesV1, *resty.Response, error) {
	return s.UnassignNetworkDevicesFromSitesV1(requestSiteDesignUnassignNetworkDevicesFromSitesV1)
}

// Alias Function
/*
This method acts as an alias for the method `UnassignsANetworkProfileForSitesFromASiteV1`
*/
func (s *SiteDesignService) UnassignsANetworkProfileForSitesFromASite(profileID string, id string) (*ResponseSiteDesignUnassignsANetworkProfileForSitesFromASiteV1, *resty.Response, error) {
	return s.UnassignsANetworkProfileForSitesFromASiteV1(profileID, id)
}

// Alias Function
/*
This method acts as an alias for the method `GetsAnAreaV1`
*/
func (s *SiteDesignService) GetsAnArea(id string) (*ResponseSiteDesignGetsAnAreaV1, *resty.Response, error) {
	return s.GetsAnAreaV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `DeletesABuildingV2`
*/
func (s *SiteDesignService) DeletesABuilding(id string) (*ResponseSiteDesignDeletesABuildingV2, *resty.Response, error) {
	return s.DeletesABuildingV2(id)
}

// Alias Function
/*
This method acts as an alias for the method `GetDeviceControllabilitySettingsV1`
*/
func (s *SiteDesignService) GetDeviceControllabilitySettings() (*ResponseSiteDesignGetDeviceControllabilitySettingsV1, *resty.Response, error) {
	return s.GetDeviceControllabilitySettingsV1()
}

// Alias Function
/*
This method acts as an alias for the method `GetSiteAssignedNetworkDevicesCountV1`
*/
func (s *SiteDesignService) GetSiteAssignedNetworkDevicesCount(GetSiteAssignedNetworkDevicesCountV1QueryParams *GetSiteAssignedNetworkDevicesCountV1QueryParams) (*ResponseSiteDesignGetSiteAssignedNetworkDevicesCountV1, *resty.Response, error) {
	return s.GetSiteAssignedNetworkDevicesCountV1(GetSiteAssignedNetworkDevicesCountV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetsABuildingV2`
*/
func (s *SiteDesignService) GetsABuilding(id string) (*ResponseSiteDesignGetsABuildingV2, *resty.Response, error) {
	return s.GetsABuildingV2(id)
}

// Alias Function
/*
This method acts as an alias for the method `AssignNetworkDevicesToASiteV1`
*/
func (s *SiteDesignService) AssignNetworkDevicesToASite(requestSiteDesignAssignNetworkDevicesToASiteV1 *RequestSiteDesignAssignNetworkDevicesToASiteV1) (*ResponseSiteDesignAssignNetworkDevicesToASiteV1, *resty.Response, error) {
	return s.AssignNetworkDevicesToASiteV1(requestSiteDesignAssignNetworkDevicesToASiteV1)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssignedV1`
*/
func (s *SiteDesignService) RetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssigned(siteID string, RetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssignedV1QueryParams *RetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssignedV1QueryParams) (*ResponseSiteDesignRetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssignedV1, *resty.Response, error) {
	return s.RetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssignedV1(siteID, RetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssignedV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `CreateSitesV1`
*/
func (s *SiteDesignService) CreateSites(requestSiteDesignCreateSitesV1 *RequestSiteDesignCreateSitesV1) (*ResponseSiteDesignCreateSitesV1, *resty.Response, error) {
	return s.CreateSitesV1(requestSiteDesignCreateSitesV1)
}

// Alias Function
/*
This method acts as an alias for the method `DeletesAFloorV2`
*/
func (s *SiteDesignService) DeletesAFloor(id string) (*ResponseSiteDesignDeletesAFloorV2, *resty.Response, error) {
	return s.DeletesAFloorV2(id)
}

// Alias Function
/*
This method acts as an alias for the method `UpdateDeviceControllabilitySettingsV1`
*/
func (s *SiteDesignService) UpdateDeviceControllabilitySettings(requestSiteDesignUpdateDeviceControllabilitySettingsV1 *RequestSiteDesignUpdateDeviceControllabilitySettingsV1) (*ResponseSiteDesignUpdateDeviceControllabilitySettingsV1, *resty.Response, error) {
	return s.UpdateDeviceControllabilitySettingsV1(requestSiteDesignUpdateDeviceControllabilitySettingsV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetFloorSettingsV2`
*/
func (s *SiteDesignService) GetFloorSettings() (*ResponseSiteDesignGetFloorSettingsV2, *resty.Response, error) {
	return s.GetFloorSettingsV2()
}

// Alias Function
/*
This method acts as an alias for the method `GetSiteAssignedNetworkDevicesV1`
*/
func (s *SiteDesignService) GetSiteAssignedNetworkDevices(GetSiteAssignedNetworkDevicesV1QueryParams *GetSiteAssignedNetworkDevicesV1QueryParams) (*ResponseSiteDesignGetSiteAssignedNetworkDevicesV1, *resty.Response, error) {
	return s.GetSiteAssignedNetworkDevicesV1(GetSiteAssignedNetworkDevicesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetSitesV1`
*/
func (s *SiteDesignService) GetSites(GetSitesV1QueryParams *GetSitesV1QueryParams) (*ResponseSiteDesignGetSitesV1, *resty.Response, error) {
	return s.GetSitesV1(GetSitesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `AssociateV1`
*/
func (s *SiteDesignService) Associate(networkProfileID string, siteID string) (*ResponseSiteDesignAssociateV1, *resty.Response, error) {
	return s.AssociateV1(networkProfileID, siteID)
}

// Alias Function
/*
This method acts as an alias for the method `DisassociateV1`
*/
func (s *SiteDesignService) Disassociate(networkProfileID string, siteID string) (*ResponseSiteDesignDisassociateV1, *resty.Response, error) {
	return s.DisassociateV1(networkProfileID, siteID)
}
