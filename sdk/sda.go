package catalyst

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SdaService service

type GetDefaultAuthenticationProfileFromSdaFabricV1QueryParams struct {
	SiteNameHierarchy        string `url:"siteNameHierarchy,omitempty"`        //siteNameHierarchy
	AuthenticateTemplateName string `url:"authenticateTemplateName,omitempty"` //authenticateTemplateName
}
type DeleteDefaultAuthenticationProfileFromSdaFabricV1QueryParams struct {
	SiteNameHierarchy string `url:"siteNameHierarchy,omitempty"` //siteNameHierarchy
}
type GetBorderDeviceDetailFromSdaFabricV1QueryParams struct {
	DeviceManagementIPAddress string `url:"deviceManagementIpAddress,omitempty"` //deviceManagementIpAddress
}
type DeleteBorderDeviceFromSdaFabricV1QueryParams struct {
	DeviceManagementIPAddress string `url:"deviceManagementIpAddress,omitempty"` //deviceManagementIpAddress
}
type DeleteControlPlaneDeviceInSdaFabricV1QueryParams struct {
	DeviceManagementIPAddress string `url:"deviceManagementIpAddress,omitempty"` //deviceManagementIpAddress
}
type GetControlPlaneDeviceFromSdaFabricV1QueryParams struct {
	DeviceManagementIPAddress string `url:"deviceManagementIpAddress,omitempty"` //deviceManagementIpAddress
}
type GetDeviceInfoFromSdaFabricV1QueryParams struct {
	DeviceManagementIPAddress string `url:"deviceManagementIpAddress,omitempty"` //deviceManagementIpAddress
}
type GetDeviceRoleInSdaFabricV1QueryParams struct {
	DeviceManagementIPAddress string `url:"deviceManagementIpAddress,omitempty"` //Device Management IP Address
}
type DeleteEdgeDeviceFromSdaFabricV1QueryParams struct {
	DeviceManagementIPAddress string `url:"deviceManagementIpAddress,omitempty"` //deviceManagementIpAddress
}
type GetEdgeDeviceFromSdaFabricV1QueryParams struct {
	DeviceManagementIPAddress string `url:"deviceManagementIpAddress,omitempty"` //deviceManagementIpAddress
}
type GetSiteFromSdaFabricV1QueryParams struct {
	SiteNameHierarchy string `url:"siteNameHierarchy,omitempty"` //Site Name Hierarchy
}
type DeleteSiteFromSdaFabricV1QueryParams struct {
	SiteNameHierarchy string `url:"siteNameHierarchy,omitempty"` //Site Name Hierarchy
}
type DeletePortAssignmentForAccessPointInSdaFabricV1QueryParams struct {
	DeviceManagementIPAddress string `url:"deviceManagementIpAddress,omitempty"` //deviceManagementIpAddress
	InterfaceName             string `url:"interfaceName,omitempty"`             //interfaceName
}
type GetPortAssignmentForAccessPointInSdaFabricV1QueryParams struct {
	DeviceManagementIPAddress string `url:"deviceManagementIpAddress,omitempty"` //deviceManagementIpAddress
	InterfaceName             string `url:"interfaceName,omitempty"`             //interfaceName
}
type DeletePortAssignmentForUserDeviceInSdaFabricV1QueryParams struct {
	DeviceManagementIPAddress string `url:"deviceManagementIpAddress,omitempty"` //deviceManagementIpAddress
	InterfaceName             string `url:"interfaceName,omitempty"`             //interfaceName
}
type GetPortAssignmentForUserDeviceInSdaFabricV1QueryParams struct {
	DeviceManagementIPAddress string `url:"deviceManagementIpAddress,omitempty"` //deviceManagementIpAddress
	InterfaceName             string `url:"interfaceName,omitempty"`             //interfaceName
}
type GetMulticastDetailsFromSdaFabricV1QueryParams struct {
	SiteNameHierarchy string `url:"siteNameHierarchy,omitempty"` //fabric site name hierarchy
}
type DeleteMulticastFromSdaFabricV1QueryParams struct {
	SiteNameHierarchy string `url:"siteNameHierarchy,omitempty"` //siteNameHierarchy
}
type DeleteProvisionedWiredDeviceV1QueryParams struct {
	DeviceManagementIPAddress string `url:"deviceManagementIpAddress,omitempty"` //Valid IP address of the device currently provisioned in a fabric site
}
type GetProvisionedWiredDeviceV1QueryParams struct {
	DeviceManagementIPAddress string `url:"deviceManagementIpAddress,omitempty"` //deviceManagementIpAddress
}
type DeleteTransitPeerNetworkV1QueryParams struct {
	TransitPeerNetworkName string `url:"transitPeerNetworkName,omitempty"` //Transit Peer Network Name
}
type GetTransitPeerNetworkInfoV1QueryParams struct {
	TransitPeerNetworkName string `url:"transitPeerNetworkName,omitempty"` //Transit or Peer Network Name
}
type DeleteVnFromSdaFabricV1QueryParams struct {
	VirtualNetworkName string `url:"virtualNetworkName,omitempty"` //virtualNetworkName
	SiteNameHierarchy  string `url:"siteNameHierarchy,omitempty"`  //siteNameHierarchy
}
type GetVnFromSdaFabricV1QueryParams struct {
	VirtualNetworkName string `url:"virtualNetworkName,omitempty"` //virtualNetworkName
	SiteNameHierarchy  string `url:"siteNameHierarchy,omitempty"`  //siteNameHierarchy
}
type GetVirtualNetworkSummaryV1QueryParams struct {
	SiteNameHierarchy string `url:"siteNameHierarchy,omitempty"` //Complete fabric siteNameHierarchy Path
}
type GetIPPoolFromSdaVirtualNetworkV1QueryParams struct {
	SiteNameHierarchy  string `url:"siteNameHierarchy,omitempty"`  //siteNameHierarchy
	VirtualNetworkName string `url:"virtualNetworkName,omitempty"` //virtualNetworkName
	IPPoolName         string `url:"ipPoolName,omitempty"`         //ipPoolName. Note: Use vlanName as a value for this parameter if same ip pool is assigned to multiple virtual networks (e.g.. ipPoolName=vlan1021)
}
type DeleteIPPoolFromSdaVirtualNetworkV1QueryParams struct {
	SiteNameHierarchy  string `url:"siteNameHierarchy,omitempty"`  //siteNameHierarchy
	VirtualNetworkName string `url:"virtualNetworkName,omitempty"` //virtualNetworkName
	IPPoolName         string `url:"ipPoolName,omitempty"`         //ipPoolName
}
type GetAnycastGatewaysV1QueryParams struct {
	ID                 string  `url:"id,omitempty"`                 //ID of the anycast gateway.
	FabricID           string  `url:"fabricId,omitempty"`           //ID of the fabric the anycast gateway is assigned to.
	VirtualNetworkName string  `url:"virtualNetworkName,omitempty"` //Name of the virtual network associated with the anycast gateways.
	IPPoolName         string  `url:"ipPoolName,omitempty"`         //Name of the IP pool associated with the anycast gateways.
	VLANName           string  `url:"vlanName,omitempty"`           //VLAN name of the anycast gateways.
	VLANID             float64 `url:"vlanId,omitempty"`             //VLAN ID of the anycast gateways. The allowed range for vlanId is [2-4093] except for reserved VLANs [1002-1005], 2046, and 4094.
	Offset             float64 `url:"offset,omitempty"`             //Starting record for pagination.
	Limit              float64 `url:"limit,omitempty"`              //Maximum number of records to return.
}
type GetAnycastGatewayCountV1QueryParams struct {
	FabricID           string  `url:"fabricId,omitempty"`           //ID of the fabric the anycast gateway is assigned to.
	VirtualNetworkName string  `url:"virtualNetworkName,omitempty"` //Name of the virtual network associated with the anycast gateways.
	IPPoolName         string  `url:"ipPoolName,omitempty"`         //Name of the IP pool associated with the anycast gateways.
	VLANName           string  `url:"vlanName,omitempty"`           //VLAN name of the anycast gateways.
	VLANID             float64 `url:"vlanId,omitempty"`             //VLAN ID of the anycast gateways. The allowed range for vlanId is [2-4093] except for reserved VLANs [1002-1005], 2046, and 4094.
}
type GetAuthenticationProfilesV1QueryParams struct {
	FabricID                  string  `url:"fabricId,omitempty"`                  //ID of the fabric the authentication profile is assigned to.
	AuthenticationProfileName string  `url:"authenticationProfileName,omitempty"` //Return only the authentication profiles with this specified name. Note that 'No Authentication' is not a valid option for this parameter.
	Offset                    float64 `url:"offset,omitempty"`                    //Starting record for pagination.
	Limit                     float64 `url:"limit,omitempty"`                     //Maximum number of records to return.
}
type DeleteExtranetPoliciesV1QueryParams struct {
	ExtranetPolicyName string `url:"extranetPolicyName,omitempty"` //Name of the extranet policy.
}
type GetExtranetPoliciesV1QueryParams struct {
	ExtranetPolicyName string  `url:"extranetPolicyName,omitempty"` //Name of the extranet policy.
	Offset             float64 `url:"offset,omitempty"`             //Starting record for pagination.
	Limit              float64 `url:"limit,omitempty"`              //Maximum number of records to return.
}
type GetFabricDevicesV1QueryParams struct {
	FabricID        string  `url:"fabricId,omitempty"`        //ID of the fabric this device belongs to.
	NetworkDeviceID string  `url:"networkDeviceId,omitempty"` //Network device ID of the fabric device.
	DeviceRoles     string  `url:"deviceRoles,omitempty"`     //Device roles of the fabric device. Allowed values are [CONTROL_PLANE_NODE, EDGE_NODE, BORDER_NODE, WIRELESS_CONTROLLER_NODE, EXTENDED_NODE].
	Offset          float64 `url:"offset,omitempty"`          //Starting record for pagination.
	Limit           float64 `url:"limit,omitempty"`           //Maximum number of records to return.
}
type DeleteFabricDevicesV1QueryParams struct {
	FabricID        string `url:"fabricId,omitempty"`        //ID of the fabric this device belongs to.
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` //Network device ID of the fabric device.
	DeviceRoles     string `url:"deviceRoles,omitempty"`     //Device roles of the fabric device. Allowed values are [CONTROL_PLANE_NODE, EDGE_NODE, BORDER_NODE, WIRELESS_CONTROLLER_NODE].
}
type GetFabricDevicesCountV1QueryParams struct {
	FabricID        string `url:"fabricId,omitempty"`        //ID of the fabric this device belongs to.
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` //Network device ID of the fabric device.
	DeviceRoles     string `url:"deviceRoles,omitempty"`     //Device roles of the fabric device. Allowed values are [CONTROL_PLANE_NODE, EDGE_NODE, BORDER_NODE, WIRELESS_CONTROLLER_NODE, EXTENDED_NODE].
}
type DeleteFabricDeviceLayer2HandoffsV1QueryParams struct {
	FabricID        string `url:"fabricId,omitempty"`        //ID of the fabric this device belongs to.
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` //Network device ID of the fabric device.
}
type GetFabricDevicesLayer2HandoffsV1QueryParams struct {
	FabricID        string  `url:"fabricId,omitempty"`        //ID of the fabric this device belongs to.
	NetworkDeviceID string  `url:"networkDeviceId,omitempty"` //Network device ID of the fabric device.
	Offset          float64 `url:"offset,omitempty"`          //Starting record for pagination.
	Limit           float64 `url:"limit,omitempty"`           //Maximum number of records to return.
}
type GetFabricDevicesLayer2HandoffsCountV1QueryParams struct {
	FabricID        string `url:"fabricId,omitempty"`        //ID of the fabric this device belongs to.
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` //Network device ID of the fabric device.
}
type DeleteFabricDeviceLayer3HandoffsWithIPTransitV1QueryParams struct {
	FabricID        string `url:"fabricId,omitempty"`        //ID of the fabric this device belongs to.
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` //Network device ID of the fabric device.
}
type GetFabricDevicesLayer3HandoffsWithIPTransitV1QueryParams struct {
	FabricID        string  `url:"fabricId,omitempty"`        //ID of the fabric this device belongs to.
	NetworkDeviceID string  `url:"networkDeviceId,omitempty"` //Network device ID of the fabric device.
	Offset          float64 `url:"offset,omitempty"`          //Starting record for pagination.
	Limit           float64 `url:"limit,omitempty"`           //Maximum number of records to return.
}
type GetFabricDevicesLayer3HandoffsWithIPTransitCountV1QueryParams struct {
	FabricID        string `url:"fabricId,omitempty"`        //ID of the fabric this device belongs to.
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` //Network device ID of the fabric device.
}
type GetFabricDevicesLayer3HandoffsWithSdaTransitV1QueryParams struct {
	FabricID        string  `url:"fabricId,omitempty"`        //ID of the fabric this device belongs to.
	NetworkDeviceID string  `url:"networkDeviceId,omitempty"` //Network device ID of the fabric device.
	Offset          float64 `url:"offset,omitempty"`          //Starting record for pagination.
	Limit           float64 `url:"limit,omitempty"`           //Maximum number of records to return.
}
type DeleteFabricDeviceLayer3HandoffsWithSdaTransitV1QueryParams struct {
	FabricID        string `url:"fabricId,omitempty"`        //ID of the fabric this device belongs to.
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` //Network device ID of the fabric device.
}
type GetFabricDevicesLayer3HandoffsWithSdaTransitCountV1QueryParams struct {
	FabricID        string `url:"fabricId,omitempty"`        //ID of the fabric this device belongs to.
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` //Network device ID of the fabric device.
}
type GetFabricSitesV1QueryParams struct {
	ID     string  `url:"id,omitempty"`     //ID of the fabric site.
	SiteID string  `url:"siteId,omitempty"` //ID of the network hierarchy associated with the fabric site.
	Offset float64 `url:"offset,omitempty"` //Starting record for pagination.
	Limit  float64 `url:"limit,omitempty"`  //Maximum number of records to return.
}
type GetFabricZonesV1QueryParams struct {
	ID     string  `url:"id,omitempty"`     //ID of the fabric zone.
	SiteID string  `url:"siteId,omitempty"` //ID of the network hierarchy associated with the fabric zone.
	Offset float64 `url:"offset,omitempty"` //Starting record for pagination.
	Limit  float64 `url:"limit,omitempty"`  //Maximum number of records to return.
}
type DeleteLayer2VirtualNetworksV1QueryParams struct {
	FabricID                           string  `url:"fabricId,omitempty"`                           //ID of the fabric the layer 2 virtual network is assigned to.
	VLANName                           string  `url:"vlanName,omitempty"`                           //The vlan name of the layer 2 virtual network.
	VLANID                             float64 `url:"vlanId,omitempty"`                             //The vlan ID of the layer 2 virtual network.
	TrafficType                        string  `url:"trafficType,omitempty"`                        //The traffic type of the layer 2 virtual network.
	AssociatedLayer3VirtualNetworkName string  `url:"associatedLayer3VirtualNetworkName,omitempty"` //Name of the associated layer 3 virtual network.
}
type GetLayer2VirtualNetworksV1QueryParams struct {
	ID                                 string  `url:"id,omitempty"`                                 //ID of the layer 2 virtual network.
	FabricID                           string  `url:"fabricId,omitempty"`                           //ID of the fabric the layer 2 virtual network is assigned to.
	VLANName                           string  `url:"vlanName,omitempty"`                           //The vlan name of the layer 2 virtual network.
	VLANID                             float64 `url:"vlanId,omitempty"`                             //The vlan ID of the layer 2 virtual network.
	TrafficType                        string  `url:"trafficType,omitempty"`                        //The traffic type of the layer 2 virtual network.
	AssociatedLayer3VirtualNetworkName string  `url:"associatedLayer3VirtualNetworkName,omitempty"` //Name of the associated layer 3 virtual network.
	Offset                             float64 `url:"offset,omitempty"`                             //Starting record for pagination.
	Limit                              float64 `url:"limit,omitempty"`                              //Maximum number of records to return.
}
type GetLayer2VirtualNetworkCountV1QueryParams struct {
	FabricID                           string  `url:"fabricId,omitempty"`                           //ID of the fabric the layer 2 virtual network is assigned to.
	VLANName                           string  `url:"vlanName,omitempty"`                           //The vlan name of the layer 2 virtual network.
	VLANID                             float64 `url:"vlanId,omitempty"`                             //The vlan ID of the layer 2 virtual network.
	TrafficType                        string  `url:"trafficType,omitempty"`                        //The traffic type of the layer 2 virtual network.
	AssociatedLayer3VirtualNetworkName string  `url:"associatedLayer3VirtualNetworkName,omitempty"` //Name of the associated layer 3 virtual network.
}
type GetLayer3VirtualNetworksV1QueryParams struct {
	VirtualNetworkName string  `url:"virtualNetworkName,omitempty"` //Name of the layer 3 virtual network.
	FabricID           string  `url:"fabricId,omitempty"`           //ID of the fabric the layer 3 virtual network is assigned to.
	AnchoredSiteID     string  `url:"anchoredSiteId,omitempty"`     //Fabric ID of the fabric site the layer 3 virtual network is anchored at.
	Offset             float64 `url:"offset,omitempty"`             //Starting record for pagination.
	Limit              float64 `url:"limit,omitempty"`              //Maximum number of records to return.
}
type DeleteLayer3VirtualNetworksV1QueryParams struct {
	VirtualNetworkName string `url:"virtualNetworkName,omitempty"` //Name of the layer 3 virtual network.
}
type GetLayer3VirtualNetworksCountV1QueryParams struct {
	FabricID       string `url:"fabricId,omitempty"`       //ID of the fabric the layer 3 virtual network is assigned to.
	AnchoredSiteID string `url:"anchoredSiteId,omitempty"` //Fabric ID of the fabric site the layer 3 virtual network is anchored at.
}
type GetMulticastV1QueryParams struct {
	FabricID string  `url:"fabricId,omitempty"` //ID of the fabric site where multicast is configured.
	Offset   float64 `url:"offset,omitempty"`   //Starting record for pagination.
	Limit    float64 `url:"limit,omitempty"`    //Maximum number of records to return.
}
type GetMulticastVirtualNetworksV1QueryParams struct {
	FabricID           string  `url:"fabricId,omitempty"`           //ID of the fabric site where multicast is configured.
	VirtualNetworkName string  `url:"virtualNetworkName,omitempty"` //Name of the virtual network associated to the multicast configuration.
	Offset             float64 `url:"offset,omitempty"`             //Starting record for pagination.
	Limit              float64 `url:"limit,omitempty"`              //Maximum number of records to return.
}
type GetMulticastVirtualNetworkCountV1QueryParams struct {
	FabricID string `url:"fabricId,omitempty"` //ID of the fabric site the multicast configuration is associated with.
}
type GetPortAssignmentsV1QueryParams struct {
	FabricID        string  `url:"fabricId,omitempty"`        //ID of the fabric the device is assigned to.
	NetworkDeviceID string  `url:"networkDeviceId,omitempty"` //Network device ID of the port assignment.
	InterfaceName   string  `url:"interfaceName,omitempty"`   //Interface name of the port assignment.
	DataVLANName    string  `url:"dataVlanName,omitempty"`    //Data VLAN name of the port assignment.
	VoiceVLANName   string  `url:"voiceVlanName,omitempty"`   //Voice VLAN name of the port assignment.
	Offset          float64 `url:"offset,omitempty"`          //Starting record for pagination.
	Limit           float64 `url:"limit,omitempty"`           //Maximum number of records to return.
}
type DeletePortAssignmentsV1QueryParams struct {
	FabricID        string `url:"fabricId,omitempty"`        //ID of the fabric the device is assigned to.
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` //Network device ID of the port assignment.
	InterfaceName   string `url:"interfaceName,omitempty"`   //Interface name of the port assignment.
	DataVLANName    string `url:"dataVlanName,omitempty"`    //Data VLAN name of the port assignment.
	VoiceVLANName   string `url:"voiceVlanName,omitempty"`   //Voice VLAN name of the port assignment.
}
type GetPortAssignmentCountV1QueryParams struct {
	FabricID        string `url:"fabricId,omitempty"`        //ID of the fabric the device is assigned to.
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` //Network device ID of the port assignment.
	InterfaceName   string `url:"interfaceName,omitempty"`   //Interface name of the port assignment.
	DataVLANName    string `url:"dataVlanName,omitempty"`    //Data VLAN name of the port assignment.
	VoiceVLANName   string `url:"voiceVlanName,omitempty"`   //Voice VLAN name of the port assignment.
}
type GetPortChannelsV1QueryParams struct {
	FabricID            string  `url:"fabricId,omitempty"`            //ID of the fabric the device is assigned to.
	NetworkDeviceID     string  `url:"networkDeviceId,omitempty"`     //ID of the network device.
	PortChannelName     string  `url:"portChannelName,omitempty"`     //Name of the port channel.
	ConnectedDeviceType string  `url:"connectedDeviceType,omitempty"` //Connected device type of the port channel. The allowed values are [TRUNK, EXTENDED_NODE].
	Offset              float64 `url:"offset,omitempty"`              //Starting record for pagination.
	Limit               float64 `url:"limit,omitempty"`               //Maximum number of records to return.
}
type DeletePortChannelsV1QueryParams struct {
	FabricID            string `url:"fabricId,omitempty"`            //ID of the fabric the device is assigned to.
	NetworkDeviceID     string `url:"networkDeviceId,omitempty"`     //ID of the network device.
	PortChannelName     string `url:"portChannelName,omitempty"`     //Name of the port channel.
	ConnectedDeviceType string `url:"connectedDeviceType,omitempty"` //Connected device type of the port channel. The allowed values are [TRUNK, EXTENDED_NODE].
}
type GetPortChannelCountV1QueryParams struct {
	FabricID            string `url:"fabricId,omitempty"`            //ID of the fabric the device is assigned to.
	NetworkDeviceID     string `url:"networkDeviceId,omitempty"`     //ID of the network device.
	PortChannelName     string `url:"portChannelName,omitempty"`     //Name of the port channel.
	ConnectedDeviceType string `url:"connectedDeviceType,omitempty"` //Connected device type of the port channel. The allowed values are [TRUNK, EXTENDED_NODE].
}
type DeleteProvisionedDevicesV1QueryParams struct {
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` //ID of the network device.
	SiteID          string `url:"siteId,omitempty"`          //ID of the site hierarchy.
}
type GetProvisionedDevicesV1QueryParams struct {
	ID              string  `url:"id,omitempty"`              //ID of the provisioned device.
	NetworkDeviceID string  `url:"networkDeviceId,omitempty"` //ID of the network device.
	SiteID          string  `url:"siteId,omitempty"`          //ID of the site hierarchy.
	Offset          float64 `url:"offset,omitempty"`          //Starting record for pagination.
	Limit           float64 `url:"limit,omitempty"`           //Maximum number of devices to return.
}
type GetProvisionedDevicesCountV1QueryParams struct {
	SiteID string `url:"siteId,omitempty"` //ID of the site hierarchy.
}
type GetTransitNetworksV1QueryParams struct {
	ID     string  `url:"id,omitempty"`     //ID of the transit network.
	Name   string  `url:"name,omitempty"`   //Name of the transit network.
	Type   string  `url:"type,omitempty"`   //Type of the transit network. Allowed values are [IP_BASED_TRANSIT, SDA_LISP_PUB_SUB_TRANSIT, SDA_LISP_BGP_TRANSIT].
	Offset float64 `url:"offset,omitempty"` //Starting record for pagination.
	Limit  float64 `url:"limit,omitempty"`  //Maximum number of records to return.
}
type GetTransitNetworksCountV1QueryParams struct {
	Type string `url:"type,omitempty"` //Type of the transit network. Allowed values are [IP_BASED_TRANSIT, SDA_LISP_PUB_SUB_TRANSIT, SDA_LISP_BGP_TRANSIT].
}
type DeleteVirtualNetworkWithScalableGroupsV1QueryParams struct {
	VirtualNetworkName string `url:"virtualNetworkName,omitempty"` //virtualNetworkName
}
type GetVirtualNetworkWithScalableGroupsV1QueryParams struct {
	VirtualNetworkName string `url:"virtualNetworkName,omitempty"` //virtualNetworkName
}

type ResponseSdaGetDefaultAuthenticationProfileFromSdaFabricV1 struct {
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         // Path of sda Fabric Site
	AuthenticateTemplateName  string `json:"authenticateTemplateName,omitempty"`  // Authenticate Template Name
	AuthenticationOrder       string `json:"authenticationOrder,omitempty"`       // Authentication Order
	Dot1XToMabFallbackTimeout string `json:"dot1xToMabFallbackTimeout,omitempty"` // Dot1x To Mab Fallback Timeout
	WakeOnLan                 *bool  `json:"wakeOnLan,omitempty"`                 // Wake On Lan
	NumberOfHosts             string `json:"numberOfHosts,omitempty"`             // Number Of Hosts
	Status                    string `json:"status,omitempty"`                    // Status
	Description               string `json:"description,omitempty"`               // Authenticate Template info reterieved successfully in sda fabric site
	ExecutionID               string `json:"executionId,omitempty"`               // Execution Id
}
type ResponseSdaAddDefaultAuthenticationTemplateInSdaFabricV1 struct {
	Status             string `json:"status,omitempty"`             // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.
	Description        string `json:"description,omitempty"`        // provides detailed information for API success or failure.
	TaskID             string `json:"taskId,omitempty"`             // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available
	TaskStatusURL      string `json:"taskStatusUrl,omitempty"`      // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>
	ExecutionID        string `json:"executionId,omitempty"`        // uuid for API execution status
}
type ResponseSdaUpdateDefaultAuthenticationProfileInSdaFabricV1 struct {
	Status             string `json:"status,omitempty"`             // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.
	Description        string `json:"description,omitempty"`        // provides detailed information for API success or failure.
	TaskID             string `json:"taskId,omitempty"`             // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available
	TaskStatusURL      string `json:"taskStatusUrl,omitempty"`      // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>
	ExecutionID        string `json:"executionId,omitempty"`        // uuid for API execution status
}
type ResponseSdaDeleteDefaultAuthenticationProfileFromSdaFabricV1 struct {
	Status             string `json:"status,omitempty"`             // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.
	Description        string `json:"description,omitempty"`        // provides detailed information for API success or failure.
	TaskID             string `json:"taskId,omitempty"`             // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available
	TaskStatusURL      string `json:"taskStatusUrl,omitempty"`      // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>
	ExecutionID        string `json:"executionId,omitempty"`        // uuid for API execution status
}
type ResponseSdaAddBorderDeviceInSdaFabricV1 struct {
	Status             string `json:"status,omitempty"`             // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.
	Description        string `json:"description,omitempty"`        // provides detailed information for API success or failure.
	TaskID             string `json:"taskId,omitempty"`             // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available
	TaskStatusURL      string `json:"taskStatusUrl,omitempty"`      // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>
	ExecutionID        string `json:"executionId,omitempty"`        // uuid for API execution status
}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricV1 struct {
	Status      string                                                  `json:"status,omitempty"`      // Status
	Description string                                                  `json:"description,omitempty"` // Description
	Payload     *ResponseSdaGetBorderDeviceDetailFromSdaFabricV1Payload `json:"payload,omitempty"`     //
}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricV1Payload struct {
	ID                             string                                                                       `json:"id,omitempty"`                             // Id
	InstanceID                     *int                                                                         `json:"instanceId,omitempty"`                     // Instance Id
	AuthEntityID                   *int                                                                         `json:"authEntityId,omitempty"`                   // Auth Entity Id
	DisplayName                    string                                                                       `json:"displayName,omitempty"`                    // Display Name
	AuthEntityClass                *int                                                                         `json:"authEntityClass,omitempty"`                // Auth Entity Class
	InstanceTenantID               string                                                                       `json:"instanceTenantId,omitempty"`               // Instance Tenant Id
	DeployPending                  string                                                                       `json:"deployPending,omitempty"`                  // Deploy Pending
	InstanceVersion                *int                                                                         `json:"instanceVersion,omitempty"`                // Instance Version
	CreateTime                     *int                                                                         `json:"createTime,omitempty"`                     // Create Time
	Deployed                       *bool                                                                        `json:"deployed,omitempty"`                       // Deployed
	IsSeeded                       *bool                                                                        `json:"isSeeded,omitempty"`                       // Is Seeded
	IsStale                        *bool                                                                        `json:"isStale,omitempty"`                        // Is Stale
	LastUpdateTime                 *int                                                                         `json:"lastUpdateTime,omitempty"`                 // Last Update Time
	Name                           string                                                                       `json:"name,omitempty"`                           // Name
	Namespace                      string                                                                       `json:"namespace,omitempty"`                      // Namespace
	ProvisioningState              string                                                                       `json:"provisioningState,omitempty"`              // Provisioning State
	ResourceVersion                *int                                                                         `json:"resourceVersion,omitempty"`                // Resource Version
	TargetIDList                   *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadTargetIDList        `json:"targetIdList,omitempty"`                   // Target Id List
	Type                           string                                                                       `json:"type,omitempty"`                           // Type
	CfsChangeInfo                  *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadCfsChangeInfo       `json:"cfsChangeInfo,omitempty"`                  // Cfs Change Info
	CustomProvisions               *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadCustomProvisions    `json:"customProvisions,omitempty"`               // Custom Provisions
	Configs                        *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadConfigs             `json:"configs,omitempty"`                        // Configs
	ManagedSites                   *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadManagedSites        `json:"managedSites,omitempty"`                   // Managed Sites
	NetworkDeviceID                string                                                                       `json:"networkDeviceId,omitempty"`                // Network Device Id
	Roles                          []string                                                                     `json:"roles,omitempty"`                          // Roles
	SaveWanConnectivityDetailsOnly *bool                                                                        `json:"saveWanConnectivityDetailsOnly,omitempty"` // Save Wan Connectivity Details Only
	SiteID                         string                                                                       `json:"siteId,omitempty"`                         // Site Id
	AkcSettingsCfs                 *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadAkcSettingsCfs      `json:"akcSettingsCfs,omitempty"`                 // Akc Settings Cfs
	DeviceInterfaceInfo            *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadDeviceInterfaceInfo `json:"deviceInterfaceInfo,omitempty"`            // Device Interface Info
	DeviceSettings                 *ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadDeviceSettings        `json:"deviceSettings,omitempty"`                 //
	NetworkWidesettings            *ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadNetworkWidesettings   `json:"networkWideSettings,omitempty"`            //
	OtherDevice                    *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadOtherDevice         `json:"otherDevice,omitempty"`                    // Other Device
	TransitNetworks                *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadTransitNetworks     `json:"transitNetworks,omitempty"`                //
	VirtualNetwork                 *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadVirtualNetwork      `json:"virtualNetwork,omitempty"`                 // Virtual Network
	WLAN                           *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadWLAN                `json:"wlan,omitempty"`                           // Wlan
}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadTargetIDList interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadCfsChangeInfo interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadCustomProvisions interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadConfigs interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadManagedSites interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadAkcSettingsCfs interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadDeviceInterfaceInfo interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadDeviceSettings struct {
	ID                            string                                                                                         `json:"id,omitempty"`                            // Id
	InstanceID                    *int                                                                                           `json:"instanceId,omitempty"`                    // Instance Id
	DisplayName                   string                                                                                         `json:"displayName,omitempty"`                   // Display Name
	InstanceTenantID              string                                                                                         `json:"instanceTenantId,omitempty"`              // Instance Tenant Id
	DeployPending                 string                                                                                         `json:"deployPending,omitempty"`                 // Deploy Pending
	InstanceVersion               *int                                                                                           `json:"instanceVersion,omitempty"`               // Instance Version
	ConnectedTo                   *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadDeviceSettingsConnectedTo             `json:"connectedTo,omitempty"`                   // Connected To
	CPU                           *float64                                                                                       `json:"cpu,omitempty"`                           // Cpu
	DhcpEnabled                   *bool                                                                                          `json:"dhcpEnabled,omitempty"`                   // Dhcp Enabled
	ExternalConnectivityIPPool    string                                                                                         `json:"externalConnectivityIpPool,omitempty"`    // External Connectivity Ip Pool
	ExternalDomainRoutingProtocol string                                                                                         `json:"externalDomainRoutingProtocol,omitempty"` // External Domain Routing Protocol
	InternalDomainProtocolNumber  string                                                                                         `json:"internalDomainProtocolNumber,omitempty"`  // Internal Domain Protocol Number
	Memory                        *float64                                                                                       `json:"memory,omitempty"`                        // Memory
	NodeType                      []string                                                                                       `json:"nodeType,omitempty"`                      // Node Type
	Storage                       *float64                                                                                       `json:"storage,omitempty"`                       // Storage
	ExtConnectivitySettings       *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadDeviceSettingsExtConnectivitySettings `json:"extConnectivitySettings,omitempty"`       //
}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadDeviceSettingsConnectedTo interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadDeviceSettingsExtConnectivitySettings struct {
	ID                           string                                                                                                  `json:"id,omitempty"`                           // Id
	InstanceID                   *int                                                                                                    `json:"instanceId,omitempty"`                   // Instance Id
	DisplayName                  string                                                                                                  `json:"displayName,omitempty"`                  // Display Name
	InstanceTenantID             string                                                                                                  `json:"instanceTenantId,omitempty"`             // Instance Tenant Id
	DeployPending                string                                                                                                  `json:"deployPending,omitempty"`                // Deploy Pending
	InstanceVersion              *int                                                                                                    `json:"instanceVersion,omitempty"`              // Instance Version
	ExternalDomainProtocolNumber string                                                                                                  `json:"externalDomainProtocolNumber,omitempty"` // External Domain Protocol Number
	InterfaceUUID                string                                                                                                  `json:"interfaceUuid,omitempty"`                // Interface Uuid
	PolicyPropagationEnabled     *bool                                                                                                   `json:"policyPropagationEnabled,omitempty"`     // Policy Propagation Enabled
	PolicySgtTag                 *float64                                                                                                `json:"policySgtTag,omitempty"`                 // Policy Sgt Tag
	L2Handoff                    *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadDeviceSettingsExtConnectivitySettingsL2Handoff `json:"l2Handoff,omitempty"`                    // L2 Handoff
	L3Handoff                    *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadDeviceSettingsExtConnectivitySettingsL3Handoff `json:"l3Handoff,omitempty"`                    //
}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadDeviceSettingsExtConnectivitySettingsL2Handoff interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadDeviceSettingsExtConnectivitySettingsL3Handoff struct {
	ID               string                                                                                                              `json:"id,omitempty"`               // Id
	InstanceID       *int                                                                                                                `json:"instanceId,omitempty"`       // Instance Id
	DisplayName      string                                                                                                              `json:"displayName,omitempty"`      // Display Name
	InstanceTenantID string                                                                                                              `json:"instanceTenantId,omitempty"` // Instance Tenant Id
	DeployPending    string                                                                                                              `json:"deployPending,omitempty"`    // Deploy Pending
	InstanceVersion  *float64                                                                                                            `json:"instanceVersion,omitempty"`  // Instance Version
	LocalIPAddress   string                                                                                                              `json:"localIpAddress,omitempty"`   // Local Ip Address
	RemoteIPAddress  string                                                                                                              `json:"remoteIpAddress,omitempty"`  // Remote Ip Address
	VLANID           *int                                                                                                                `json:"vlanId,omitempty"`           // Vlan Id
	VirtualNetwork   *ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadDeviceSettingsExtConnectivitySettingsL3HandoffVirtualNetwork `json:"virtualNetwork,omitempty"`   //
}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadDeviceSettingsExtConnectivitySettingsL3HandoffVirtualNetwork struct {
	IDRef string `json:"idRef,omitempty"` // Id Ref
}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadNetworkWidesettings struct {
	ID               string                                                                                 `json:"id,omitempty"`               // Id
	InstanceID       *int                                                                                   `json:"instanceId,omitempty"`       // Instance Id
	DisplayName      string                                                                                 `json:"displayName,omitempty"`      // Display Name
	InstanceTenantID string                                                                                 `json:"instanceTenantId,omitempty"` // Instance Tenant Id
	DeployPending    string                                                                                 `json:"deployPending,omitempty"`    // Deploy Pending
	InstanceVersion  *int                                                                                   `json:"instanceVersion,omitempty"`  // Instance Version
	AAA              *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadNetworkWidesettingsAAA        `json:"aaa,omitempty"`              // Aaa
	Cmx              *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadNetworkWidesettingsCmx        `json:"cmx,omitempty"`              // Cmx
	Dhcp             *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadNetworkWidesettingsDhcp       `json:"dhcp,omitempty"`             //
	DNS              *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadNetworkWidesettingsDNS        `json:"dns,omitempty"`              //
	Ldap             *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadNetworkWidesettingsLdap       `json:"ldap,omitempty"`             // Ldap
	NativeVLAN       *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadNetworkWidesettingsNativeVLAN `json:"nativeVlan,omitempty"`       // Native Vlan
	Netflow          *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadNetworkWidesettingsNetflow    `json:"netflow,omitempty"`          // Netflow
	Ntp              *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadNetworkWidesettingsNtp        `json:"ntp,omitempty"`              // Ntp
	SNMP             *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadNetworkWidesettingsSNMP       `json:"snmp,omitempty"`             // Snmp
	Syslogs          *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadNetworkWidesettingsSyslogs    `json:"syslogs,omitempty"`          // Syslogs
}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadNetworkWidesettingsAAA interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadNetworkWidesettingsCmx interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadNetworkWidesettingsDhcp struct {
	ID        string                                                                                  `json:"id,omitempty"`        // Id
	IPAddress *ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadNetworkWidesettingsDhcpIPAddress `json:"ipAddress,omitempty"` //
}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadNetworkWidesettingsDhcpIPAddress struct {
	ID            string `json:"id,omitempty"`            // Id
	PaddedAddress string `json:"paddedAddress,omitempty"` // Padded Address
	AddressType   string `json:"addressType,omitempty"`   // Address Type
	Address       string `json:"address,omitempty"`       // Address
}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadNetworkWidesettingsDNS struct {
	ID         string                                                                          `json:"id,omitempty"`         // Id
	DomainName string                                                                          `json:"domainName,omitempty"` // Domain Name
	IP         *ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadNetworkWidesettingsDNSIP `json:"ip,omitempty"`         //
}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadNetworkWidesettingsDNSIP struct {
	ID            string `json:"id,omitempty"`            // Id
	PaddedAddress string `json:"paddedAddress,omitempty"` // Padded Address
	AddressType   string `json:"addressType,omitempty"`   // Address Type
	Address       string `json:"address,omitempty"`       // Address
}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadNetworkWidesettingsLdap interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadNetworkWidesettingsNativeVLAN interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadNetworkWidesettingsNetflow interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadNetworkWidesettingsNtp interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadNetworkWidesettingsSNMP interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadNetworkWidesettingsSyslogs interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadOtherDevice interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadTransitNetworks struct {
	IDRef string `json:"idRef,omitempty"` // Id Ref
}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadVirtualNetwork interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricV1PayloadWLAN interface{}
type ResponseSdaDeleteBorderDeviceFromSdaFabricV1 struct {
	Status             string `json:"status,omitempty"`             // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.
	Description        string `json:"description,omitempty"`        // provides detailed information for API success or failure.
	TaskID             string `json:"taskId,omitempty"`             // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available
	TaskStatusURL      string `json:"taskStatusUrl,omitempty"`      // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>
	ExecutionID        string `json:"executionId,omitempty"`        // uuid for API execution status
}
type ResponseSdaDeleteControlPlaneDeviceInSdaFabricV1 struct {
	Status             string `json:"status,omitempty"`             // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.
	Description        string `json:"description,omitempty"`        // provides detailed information for API success or failure.
	TaskID             string `json:"taskId,omitempty"`             // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available
	TaskStatusURL      string `json:"taskStatusUrl,omitempty"`      // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>
	ExecutionID        string `json:"executionId,omitempty"`        // uuid for API execution status
}
type ResponseSdaGetControlPlaneDeviceFromSdaFabricV1 struct {
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` // Management Ip Address of the Device which is provisioned successfully
	DeviceName                string `json:"deviceName,omitempty"`                // Device Name
	Roles                     string `json:"roles,omitempty"`                     // Assigned roles
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         // Site Name Hierarchy
	RouteDistributionProtocol string `json:"routeDistributionProtocol,omitempty"` // routeDistributionProtocol
	Status                    string `json:"status,omitempty"`                    // Status
	Description               string `json:"description,omitempty"`               // Control plane device info retrieved successfully in sda fabric
}
type ResponseSdaAddControlPlaneDeviceInSdaFabricV1 struct {
	Status             string `json:"status,omitempty"`             // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.
	Description        string `json:"description,omitempty"`        // provides detailed information for API success or failure.
	TaskID             string `json:"taskId,omitempty"`             // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available
	TaskStatusURL      string `json:"taskStatusUrl,omitempty"`      // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>
	ExecutionID        string `json:"executionId,omitempty"`        // uuid for API execution status
}
type ResponseSdaGetDeviceInfoFromSdaFabricV1 struct {
	Status                    string   `json:"status,omitempty"`                    // Status
	Description               string   `json:"description,omitempty"`               // Description
	Name                      string   `json:"name,omitempty"`                      // Name
	Roles                     []string `json:"roles,omitempty"`                     // Roles
	DeviceManagementIPAddress string   `json:"deviceManagementIpAddress,omitempty"` // Device Management Ip Address
	SiteHierarchy             string   `json:"siteHierarchy,omitempty"`             // Site Hierarchy
}
type ResponseSdaGetDeviceRoleInSdaFabricV1 struct {
	Roles       []string `json:"roles,omitempty"`       // Assigned device roles. Possible roles are [Edge Node, Control Plane, Border Node, Extended Node, Wireless Controller, Transit Control Plane]
	Status      string   `json:"status,omitempty"`      // Status indicates if API failed or passed.
	Description string   `json:"description,omitempty"` // Device role successfully retrieved from sda fabric.
}
type ResponseSdaAddEdgeDeviceInSdaFabricV1 struct {
	Status             string `json:"status,omitempty"`             // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.
	Description        string `json:"description,omitempty"`        // provides detailed information for API success or failure.
	TaskID             string `json:"taskId,omitempty"`             // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available
	TaskStatusURL      string `json:"taskStatusUrl,omitempty"`      // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>
	ExecutionID        string `json:"executionId,omitempty"`        // uuid for API execution status
}
type ResponseSdaDeleteEdgeDeviceFromSdaFabricV1 struct {
	Status             string `json:"status,omitempty"`             // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.
	Description        string `json:"description,omitempty"`        // provides detailed information for API success or failure.
	TaskID             string `json:"taskId,omitempty"`             // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available
	TaskStatusURL      string `json:"taskStatusUrl,omitempty"`      // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>
	ExecutionID        string `json:"executionId,omitempty"`        // uuid for API execution status
}
type ResponseSdaGetEdgeDeviceFromSdaFabricV1 struct {
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` // Management Ip Address of the Device which is provisioned successfully
	DeviceName                string `json:"deviceName,omitempty"`                // Device Name
	Roles                     string `json:"roles,omitempty"`                     // Assigned roles
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         // Site Name Hierarchy
	FabricSiteNameHierarchy   string `json:"fabricSiteNameHierarchy,omitempty"`   // Fabric Site Name Hierarchy
	Status                    string `json:"status,omitempty"`                    // Status
	Description               string `json:"description,omitempty"`               // Edge device info retrieved successfully in sda fabric
}
type ResponseSdaGetSiteFromSdaFabricV1 struct {
	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Site Name Hierarchy
	FabricName        string `json:"fabricName,omitempty"`        // Fabric Name
	FabricType        string `json:"fabricType,omitempty"`        // Fabric Type
	FabricDomainType  string `json:"fabricDomainType,omitempty"`  // Fabric Domain Type
	Status            string `json:"status,omitempty"`            // Status
	Description       string `json:"description,omitempty"`       // Fabric Site info successfully retrieved from sda fabric
}
type ResponseSdaDeleteSiteFromSdaFabricV1 struct {
	Status             string `json:"status,omitempty"`             // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.
	Description        string `json:"description,omitempty"`        // provides detailed information for API success or failure.
	TaskID             string `json:"taskId,omitempty"`             // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available
	TaskStatusURL      string `json:"taskStatusUrl,omitempty"`      // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>
	ExecutionID        string `json:"executionId,omitempty"`        // uuid for API execution status
}
type ResponseSdaAddSiteInSdaFabricV1 struct {
	Status             string `json:"status,omitempty"`             // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.
	Description        string `json:"description,omitempty"`        // provides detailed information for API success or failure.
	TaskID             string `json:"taskId,omitempty"`             // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available
	TaskStatusURL      string `json:"taskStatusUrl,omitempty"`      // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>
	ExecutionID        string `json:"executionId,omitempty"`        // uuid for API execution status
}
type ResponseSdaAddPortAssignmentForAccessPointInSdaFabricV1 struct {
	Status             string `json:"status,omitempty"`             // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.
	Description        string `json:"description,omitempty"`        // provides detailed information for API success or failure.
	TaskID             string `json:"taskId,omitempty"`             // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available
	TaskStatusURL      string `json:"taskStatusUrl,omitempty"`      // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>
	ExecutionID        string `json:"executionId,omitempty"`        // uuid for API execution status
}
type ResponseSdaDeletePortAssignmentForAccessPointInSdaFabricV1 struct {
	Status             string `json:"status,omitempty"`             // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.
	Description        string `json:"description,omitempty"`        // provides detailed information for API success or failure.
	TaskID             string `json:"taskId,omitempty"`             // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available
	TaskStatusURL      string `json:"taskStatusUrl,omitempty"`      // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>
	ExecutionID        string `json:"executionId,omitempty"`        // uuid for API execution status
}
type ResponseSdaGetPortAssignmentForAccessPointInSdaFabricV1 struct {
	Status                    string `json:"status,omitempty"`                    // Status
	Description               string `json:"description,omitempty"`               // Description
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         // Site Name Hierarchy
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` // Device Management Ip Address
	InterfaceName             string `json:"interfaceName,omitempty"`             // Interface Name
	DataIPAddressPoolName     string `json:"dataIpAddressPoolName,omitempty"`     // Data Ip Address Pool Name
	VoiceIPAddressPoolName    string `json:"voiceIpAddressPoolName,omitempty"`    // Voice Ip Address Pool Name
	ScalableGroupName         string `json:"scalableGroupName,omitempty"`         // Scalable Group Name
	AuthenticateTemplateName  string `json:"authenticateTemplateName,omitempty"`  // Authenticate Template Name
}
type ResponseSdaDeletePortAssignmentForUserDeviceInSdaFabricV1 struct {
	Status             string `json:"status,omitempty"`             // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.
	Description        string `json:"description,omitempty"`        // provides detailed information for API success or failure.
	TaskID             string `json:"taskId,omitempty"`             // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available
	TaskStatusURL      string `json:"taskStatusUrl,omitempty"`      // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>
	ExecutionID        string `json:"executionId,omitempty"`        // uuid for API execution status
}
type ResponseSdaAddPortAssignmentForUserDeviceInSdaFabricV1 struct {
	Status             string `json:"status,omitempty"`             // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.
	Description        string `json:"description,omitempty"`        // provides detailed information for API success or failure.
	TaskID             string `json:"taskId,omitempty"`             // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available
	TaskStatusURL      string `json:"taskStatusUrl,omitempty"`      // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>
	ExecutionID        string `json:"executionId,omitempty"`        // uuid for API execution status
}
type ResponseSdaGetPortAssignmentForUserDeviceInSdaFabricV1 struct {
	Status                    string `json:"status,omitempty"`                    // Status
	Description               string `json:"description,omitempty"`               // Description
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         // Site Name Hierarchy
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` // Device Management Ip Address
	InterfaceName             string `json:"interfaceName,omitempty"`             // Interface Name
	DataIPAddressPoolName     string `json:"dataIpAddressPoolName,omitempty"`     // Data Ip Address Pool Name
	VoiceIPAddressPoolName    string `json:"voiceIpAddressPoolName,omitempty"`    // Voice Ip Address Pool Name
	ScalableGroupName         string `json:"scalableGroupName,omitempty"`         // Scalable Group Name
	AuthenticateTemplateName  string `json:"authenticateTemplateName,omitempty"`  // Authenticate Template Name
}
type ResponseSdaAddMulticastInSdaFabricV1 struct {
	Status             string `json:"status,omitempty"`             // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.
	Description        string `json:"description,omitempty"`        // provides detailed information for API success or failure.
	TaskID             string `json:"taskId,omitempty"`             // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available
	TaskStatusURL      string `json:"taskStatusUrl,omitempty"`      // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>
	ExecutionID        string `json:"executionId,omitempty"`        // uuid for API execution status
}
type ResponseSdaGetMulticastDetailsFromSdaFabricV1 struct {
	MulticastMethod string                                                          `json:"multicastMethod,omitempty"` // Multicast Method
	MulticastType   string                                                          `json:"multicastType,omitempty"`   // Multicast Type
	MulticastVnInfo *[]ResponseSdaGetMulticastDetailsFromSdaFabricV1MulticastVnInfo `json:"multicastVnInfo,omitempty"` //
	Status          string                                                          `json:"status,omitempty"`          // Status
	Description     string                                                          `json:"description,omitempty"`     // multicast configuration info retrieved successfully from sda fabric
}
type ResponseSdaGetMulticastDetailsFromSdaFabricV1MulticastVnInfo struct {
	VirtualNetworkName  string                                                                 `json:"virtualNetworkName,omitempty"`  // Virtual Network Name, that is associated to Fabric Site
	IPPoolName          string                                                                 `json:"ipPoolName,omitempty"`          // Ip Pool Name, that is reserved to Fabric Site
	InternalRpIPAddress []string                                                               `json:"internalRpIpAddress,omitempty"` // InternalRpIpAddress
	ExternalRpIPAddress string                                                                 `json:"externalRpIpAddress,omitempty"` // ExternalRpIpAddress
	SsmInfo             *[]ResponseSdaGetMulticastDetailsFromSdaFabricV1MulticastVnInfoSsmInfo `json:"ssmInfo,omitempty"`             //
}
type ResponseSdaGetMulticastDetailsFromSdaFabricV1MulticastVnInfoSsmInfo struct {
	SsmGroupRange   string `json:"ssmGroupRange,omitempty"`   // SSM group range
	SsmWildcardMask string `json:"ssmWildcardMask,omitempty"` // SSM Wildcard Mask
}
type ResponseSdaDeleteMulticastFromSdaFabricV1 struct {
	Status             string `json:"status,omitempty"`             // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.
	Description        string `json:"description,omitempty"`        // provides detailed information for API success or failure.
	TaskID             string `json:"taskId,omitempty"`             // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available
	TaskStatusURL      string `json:"taskStatusUrl,omitempty"`      // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>
	ExecutionID        string `json:"executionId,omitempty"`        // uuid for API execution status
}
type ResponseSdaDeleteProvisionedWiredDeviceV1 struct {
	Status             string `json:"status,omitempty"`             // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.
	Description        string `json:"description,omitempty"`        // provides detailed information for API success or failure.
	TaskID             string `json:"taskId,omitempty"`             // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available
	TaskStatusURL      string `json:"taskStatusUrl,omitempty"`      // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>
	ExecutionID        string `json:"executionId,omitempty"`        // uuid for API execution status
}
type ResponseSdaReProvisionWiredDeviceV1 struct {
	Status             string `json:"status,omitempty"`             // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.
	Description        string `json:"description,omitempty"`        // provides detailed information for API success or failure.
	TaskID             string `json:"taskId,omitempty"`             // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available
	TaskStatusURL      string `json:"taskStatusUrl,omitempty"`      // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>
	ExecutionID        string `json:"executionId,omitempty"`        // uuid for API execution status
}
type ResponseSdaProvisionWiredDeviceV1 struct {
	Status             string `json:"status,omitempty"`             // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.
	Description        string `json:"description,omitempty"`        // provides detailed information for API success or failure.
	TaskID             string `json:"taskId,omitempty"`             // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available
	TaskStatusURL      string `json:"taskStatusUrl,omitempty"`      // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>
	ExecutionID        string `json:"executionId,omitempty"`        // uuid for API execution status
}
type ResponseSdaGetProvisionedWiredDeviceV1 struct {
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` // Management Ip Address of the device to be provisioned
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         // Site Name Hierarchy for device location(only building / floor level)
	Status                    string `json:"status,omitempty"`                    // Status
	Description               string `json:"description,omitempty"`               // Wired Provisioned device detail retrieved successfully
}
type ResponseSdaDeleteTransitPeerNetworkV1 struct {
	Status             string `json:"status,omitempty"`             // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.
	Description        string `json:"description,omitempty"`        // provides detailed information for API success or failure.
	TaskID             string `json:"taskId,omitempty"`             // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available
	TaskStatusURL      string `json:"taskStatusUrl,omitempty"`      // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>
	ExecutionID        string `json:"executionId,omitempty"`        // uuid for API execution status
}
type ResponseSdaGetTransitPeerNetworkInfoV1 struct {
	TransitPeerNetworkName string                                                    `json:"transitPeerNetworkName,omitempty"` // Transit Peer Network Name
	TransitPeerNetworkType string                                                    `json:"transitPeerNetworkType,omitempty"` // Transit Peer Network Type
	IPTransitSettings      *ResponseSdaGetTransitPeerNetworkInfoV1IPTransitSettings  `json:"ipTransitSettings,omitempty"`      //
	SdaTransitSettings     *ResponseSdaGetTransitPeerNetworkInfoV1SdaTransitSettings `json:"sdaTransitSettings,omitempty"`     //
	Status                 string                                                    `json:"status,omitempty"`                 // status
	Description            string                                                    `json:"description,omitempty"`            // Transit Peer network info retrieved successfully
	TransitPeerNetworkID   string                                                    `json:"transitPeerNetworkId,omitempty"`   // Transit Peer Network Id
}
type ResponseSdaGetTransitPeerNetworkInfoV1IPTransitSettings struct {
	RoutingProtocolName    string `json:"routingProtocolName,omitempty"`    // Routing Protocol Name
	AutonomousSystemNumber string `json:"autonomousSystemNumber,omitempty"` // Autonomous System Number
}
type ResponseSdaGetTransitPeerNetworkInfoV1SdaTransitSettings struct {
	TransitControlPlaneSettings *[]ResponseSdaGetTransitPeerNetworkInfoV1SdaTransitSettingsTransitControlPlaneSettings `json:"transitControlPlaneSettings,omitempty"` //
}
type ResponseSdaGetTransitPeerNetworkInfoV1SdaTransitSettingsTransitControlPlaneSettings struct {
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         // Site Name Hierarchy
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` // Device Management Ip Address
}
type ResponseSdaAddTransitPeerNetworkV1 struct {
	Status             string `json:"status,omitempty"`             // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.
	Description        string `json:"description,omitempty"`        // provides detailed information for API success or failure.
	TaskID             string `json:"taskId,omitempty"`             // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available
	TaskStatusURL      string `json:"taskStatusUrl,omitempty"`      // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>
	ExecutionID        string `json:"executionId,omitempty"`        // uuid for API execution status
}
type ResponseSdaDeleteVnFromSdaFabricV1 struct {
	Status             string `json:"status,omitempty"`             // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.
	Description        string `json:"description,omitempty"`        // provides detailed information for API success or failure.
	TaskID             string `json:"taskId,omitempty"`             // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available
	TaskStatusURL      string `json:"taskStatusUrl,omitempty"`      // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>
	ExecutionID        string `json:"executionId,omitempty"`        // uuid for API execution status
}
type ResponseSdaGetVnFromSdaFabricV1 struct {
	SiteNameHierarchy       string `json:"siteNameHierarchy,omitempty"`       // Path of sda Fabric Site
	VirtualNetworkName      string `json:"virtualNetworkName,omitempty"`      // Virtual Network Name
	FabricName              string `json:"fabricName,omitempty"`              // Fabric Name
	IsInfraVn               *bool  `json:"isInfraVN,omitempty"`               // Infra VN
	IsDefaultVn             *bool  `json:"isDefaultVN,omitempty"`             // Default VN
	VirtualNetworkContextID string `json:"virtualNetworkContextId,omitempty"` // Virtual Network Context Id
	VirtualNetworkID        string `json:"virtualNetworkId,omitempty"`        // Virtual Network Id
	Status                  string `json:"status,omitempty"`                  // Status
	Description             string `json:"description,omitempty"`             // Virtual Network info retrieved successfully from SDA Fabric
	ExecutionID             string `json:"executionId,omitempty"`             // Execution Id
}
type ResponseSdaAddVnInFabricV1 struct {
	Status             string `json:"status,omitempty"`             // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.
	Description        string `json:"description,omitempty"`        // provides detailed information for API success or failure.
	TaskID             string `json:"taskId,omitempty"`             // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available
	TaskStatusURL      string `json:"taskStatusUrl,omitempty"`      // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>
	ExecutionID        string `json:"executionId,omitempty"`        // uuid for API execution status
}
type ResponseSdaGetVirtualNetworkSummaryV1 struct {
	VirtualNetworkCount   *int                                                          `json:"virtualNetworkCount,omitempty"`   // Virtual Networks Count
	VirtualNetworkSummary *[]ResponseSdaGetVirtualNetworkSummaryV1VirtualNetworkSummary `json:"virtualNetworkSummary,omitempty"` //
	Status                string                                                        `json:"status,omitempty"`                // Status
	Description           string                                                        `json:"description,omitempty"`           // Virtual Network summary retrieved successfully from SDA Fabric
	ExecutionID           string                                                        `json:"executionId,omitempty"`           // Execution Id
}
type ResponseSdaGetVirtualNetworkSummaryV1VirtualNetworkSummary struct {
	VirtualNetworkContextID string `json:"virtualNetworkContextId,omitempty"` // Virtual Network Context Id
	VirtualNetworkID        string `json:"virtualNetworkId,omitempty"`        // Virtual Network Id
	SiteNameHierarchy       string `json:"siteNameHierarchy,omitempty"`       // Site Name Hierarchy
	VirtualNetworkName      string `json:"virtualNetworkName,omitempty"`      // Virtual Network Name
	Layer3Instance          *int   `json:"layer3Instance,omitempty"`          // layer3 Instance
	VirtualNetworkStatus    string `json:"virtualNetworkStatus,omitempty"`    // Virtual Network Status
}
type ResponseSdaGetIPPoolFromSdaVirtualNetworkV1 struct {
	Status                   string `json:"status,omitempty"`                   // Status
	Description              string `json:"description,omitempty"`              // Description
	VirtualNetworkName       string `json:"virtualNetworkName,omitempty"`       // Virtual Network Name
	IPPoolName               string `json:"ipPoolName,omitempty"`               // Ip Pool Name
	AuthenticationPolicyName string `json:"authenticationPolicyName,omitempty"` // Authentication Policy Name
	TrafficType              string `json:"trafficType,omitempty"`              // Traffic Type
	ScalableGroupName        string `json:"scalableGroupName,omitempty"`        // Scalable Group Name
	IsL2FloodingEnabled      *bool  `json:"isL2FloodingEnabled,omitempty"`      // Is L2 Flooding Enabled
	IsThisCriticalPool       *bool  `json:"isThisCriticalPool,omitempty"`       // Is This Critical Pool
}
type ResponseSdaDeleteIPPoolFromSdaVirtualNetworkV1 struct {
	Status             string `json:"status,omitempty"`             // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.
	Description        string `json:"description,omitempty"`        // provides detailed information for API success or failure.
	TaskID             string `json:"taskId,omitempty"`             // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available
	TaskStatusURL      string `json:"taskStatusUrl,omitempty"`      // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>
	ExecutionID        string `json:"executionId,omitempty"`        // uuid for API execution status
}
type ResponseSdaAddIPPoolInSdaVirtualNetworkV1 struct {
	Status             string `json:"status,omitempty"`             // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.
	Description        string `json:"description,omitempty"`        // provides detailed information for API success or failure.
	TaskID             string `json:"taskId,omitempty"`             // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available
	TaskStatusURL      string `json:"taskStatusUrl,omitempty"`      // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>
	ExecutionID        string `json:"executionId,omitempty"`        // uuid for API execution status
}
type ResponseSdaUpdateAnycastGatewaysV1 struct {
	Response *ResponseSdaUpdateAnycastGatewaysV1Response `json:"response,omitempty"` //
	Version  string                                      `json:"version,omitempty"`  // Version number.
}
type ResponseSdaUpdateAnycastGatewaysV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaAddAnycastGatewaysV1 struct {
	Response *ResponseSdaAddAnycastGatewaysV1Response `json:"response,omitempty"` //
	Version  string                                   `json:"version,omitempty"`  // Version number.
}
type ResponseSdaAddAnycastGatewaysV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaGetAnycastGatewaysV1 struct {
	Response *[]ResponseSdaGetAnycastGatewaysV1Response `json:"response,omitempty"` //
	Version  string                                     `json:"version,omitempty"`  // Version number.
}
type ResponseSdaGetAnycastGatewaysV1Response struct {
	ID                                      string `json:"id,omitempty"`                                      // ID of the anycast gateway.
	FabricID                                string `json:"fabricId,omitempty"`                                // ID of the fabric this anycast gateway is assigned to.
	VirtualNetworkName                      string `json:"virtualNetworkName,omitempty"`                      // Name of the layer 3 virtual network associated with the anycast gateway.
	IPPoolName                              string `json:"ipPoolName,omitempty"`                              // Name of the IP pool associated with the anycast gateway.
	TCPMssAdjustment                        *int   `json:"tcpMssAdjustment,omitempty"`                        // TCP maximum segment size adjustment.
	VLANName                                string `json:"vlanName,omitempty"`                                // Name of the VLAN of the anycast gateway.
	VLANID                                  *int   `json:"vlanId,omitempty"`                                  // ID of the VLAN of the anycast gateway.
	TrafficType                             string `json:"trafficType,omitempty"`                             // The type of traffic the anycast gateway serves.
	PoolType                                string `json:"poolType,omitempty"`                                // The pool type of the anycast gateway (applicable only to INFRA_VN).
	SecurityGroupName                       string `json:"securityGroupName,omitempty"`                       // Name of the associated Security Group (not applicable to INFRA_VN).
	IsCriticalPool                          *bool  `json:"isCriticalPool,omitempty"`                          // Enable/disable critical VLAN (not applicable to INFRA_VN).
	IsLayer2FloodingEnabled                 *bool  `json:"isLayer2FloodingEnabled,omitempty"`                 // Enable/disable layer 2 flooding (not applicable to INFRA_VN).
	IsWirelessPool                          *bool  `json:"isWirelessPool,omitempty"`                          // Enable/disable fabric-enabled wireless (not applicable to INFRA_VN).
	IsIPDirectedBroadcast                   *bool  `json:"isIpDirectedBroadcast,omitempty"`                   // Enable/disable IP-directed broadcast (not applicable to INFRA_VN).
	IsIntraSubnetRoutingEnabled             *bool  `json:"isIntraSubnetRoutingEnabled,omitempty"`             // Enable/disable Intra-Subnet Routing (not applicable to INFRA_VN).
	IsMultipleIPToMacAddresses              *bool  `json:"isMultipleIpToMacAddresses,omitempty"`              // Enable/disable multiple IP-to-MAC Addresses (Wireless Bridged-Network Virtual Machine; not applicable to INFRA_VN).
	IsSupplicantBasedExtendedNodeOnboarding *bool  `json:"isSupplicantBasedExtendedNodeOnboarding,omitempty"` // Enable/disable Supplicant-Based Extended Node Onboarding (applicable only to INFRA_VN).
	IsGroupBasedPolicyEnforcementEnabled    *bool  `json:"isGroupBasedPolicyEnforcementEnabled,omitempty"`    // Enable/disable Group-Based Policy Enforcement (applicable only to INFRA_VN).
}
type ResponseSdaGetAnycastGatewayCountV1 struct {
	Response *ResponseSdaGetAnycastGatewayCountV1Response `json:"response,omitempty"` //
	Version  string                                       `json:"version,omitempty"`  // Version number.
}
type ResponseSdaGetAnycastGatewayCountV1Response struct {
	Count *int `json:"count,omitempty"` // The number of anycast gateways.
}
type ResponseSdaDeleteAnycastGatewayByIDV1 struct {
	Response *ResponseSdaDeleteAnycastGatewayByIDV1Response `json:"response,omitempty"` //
	Version  string                                         `json:"version,omitempty"`  // Version number.
}
type ResponseSdaDeleteAnycastGatewayByIDV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaGetAuthenticationProfilesV1 struct {
	Response *[]ResponseSdaGetAuthenticationProfilesV1Response `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  // Version number.
}
type ResponseSdaGetAuthenticationProfilesV1Response struct {
	ID                        string `json:"id,omitempty"`                        // ID of the authentication profile.
	FabricID                  string `json:"fabricId,omitempty"`                  // ID of the fabric this authentication profile is assigned to.
	AuthenticationProfileName string `json:"authenticationProfileName,omitempty"` // The default host authentication template.
	AuthenticationOrder       string `json:"authenticationOrder,omitempty"`       // First authentication method.
	Dot1XToMabFallbackTimeout *int   `json:"dot1xToMabFallbackTimeout,omitempty"` // 802.1x Timeout.
	WakeOnLan                 *bool  `json:"wakeOnLan,omitempty"`                 // Wake on LAN.
	NumberOfHosts             string `json:"numberOfHosts,omitempty"`             // Number of Hosts.
	IsBpduGuardEnabled        *bool  `json:"isBpduGuardEnabled,omitempty"`        // Enable/disable BPDU Guard. Only applicable when authenticationProfileName is set to "Closed Authentication".
}
type ResponseSdaUpdateAuthenticationProfileV1 struct {
	Response *ResponseSdaUpdateAuthenticationProfileV1Response `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  // Version number.
}
type ResponseSdaUpdateAuthenticationProfileV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaDeleteExtranetPoliciesV1 struct {
	Response *ResponseSdaDeleteExtranetPoliciesV1Response `json:"response,omitempty"` //
	Version  string                                       `json:"version,omitempty"`  // Version number.
}
type ResponseSdaDeleteExtranetPoliciesV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaUpdateExtranetPolicyV1 struct {
	Response *ResponseSdaUpdateExtranetPolicyV1Response `json:"response,omitempty"` //
	Version  string                                     `json:"version,omitempty"`  // Version number.
}
type ResponseSdaUpdateExtranetPolicyV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaAddExtranetPolicyV1 struct {
	Response *ResponseSdaAddExtranetPolicyV1Response `json:"response,omitempty"` //
	Version  string                                  `json:"version,omitempty"`  // Version number.
}
type ResponseSdaAddExtranetPolicyV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaGetExtranetPoliciesV1 struct {
	Response *[]ResponseSdaGetExtranetPoliciesV1Response `json:"response,omitempty"` //
	Version  string                                      `json:"version,omitempty"`  // Version number.
}
type ResponseSdaGetExtranetPoliciesV1Response struct {
	ID                            string   `json:"id,omitempty"`                            // ID of the extranet policy.
	ExtranetPolicyName            string   `json:"extranetPolicyName,omitempty"`            // Name of the extranet policy.
	FabricIDs                     []string `json:"fabricIds,omitempty"`                     // IDs of the fabric sites associated with this extranet policy.
	ProviderVirtualNetworkName    string   `json:"providerVirtualNetworkName,omitempty"`    // Name of the provider virtual network.
	SubscriberVirtualNetworkNames []string `json:"subscriberVirtualNetworkNames,omitempty"` // Name of the subscriber virtual network names.
}
type ResponseSdaGetExtranetPolicyCountV1 struct {
	Response *ResponseSdaGetExtranetPolicyCountV1Response `json:"response,omitempty"` //
	Version  string                                       `json:"version,omitempty"`  // Version number.
}
type ResponseSdaGetExtranetPolicyCountV1Response struct {
	Count *int `json:"count,omitempty"` // Number of extranet policies.
}
type ResponseSdaDeleteExtranetPolicyByIDV1 struct {
	Response *ResponseSdaDeleteExtranetPolicyByIDV1Response `json:"response,omitempty"` //
	Version  string                                         `json:"version,omitempty"`  // Version number.
}
type ResponseSdaDeleteExtranetPolicyByIDV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaGetFabricDevicesV1 struct {
	Response *[]ResponseSdaGetFabricDevicesV1Response `json:"response,omitempty"` //
	Version  string                                   `json:"version,omitempty"`  // Version number.
}
type ResponseSdaGetFabricDevicesV1Response struct {
	ID                   string                                                     `json:"id,omitempty"`                   // ID of the fabric device.
	NetworkDeviceID      string                                                     `json:"networkDeviceId,omitempty"`      // Network device ID of the fabric device.
	FabricID             string                                                     `json:"fabricId,omitempty"`             // ID of the fabric of this fabric device.
	DeviceRoles          []string                                                   `json:"deviceRoles,omitempty"`          // List of the roles of the fabric device. Allowed values are [CONTROL_PLANE_NODE, EDGE_NODE, BORDER_NODE, WIRELESS_CONTROLLER_NODE, EXTENDED_NODE].
	BorderDeviceSettings *ResponseSdaGetFabricDevicesV1ResponseBorderDeviceSettings `json:"borderDeviceSettings,omitempty"` //
}
type ResponseSdaGetFabricDevicesV1ResponseBorderDeviceSettings struct {
	BorderTypes    []string                                                                 `json:"borderTypes,omitempty"`    // List of the border types of the fabric device. Allowed values are [LAYER_2, LAYER_3].
	Layer3Settings *ResponseSdaGetFabricDevicesV1ResponseBorderDeviceSettingsLayer3Settings `json:"layer3Settings,omitempty"` //
}
type ResponseSdaGetFabricDevicesV1ResponseBorderDeviceSettingsLayer3Settings struct {
	LocalAutonomousSystemNumber  string `json:"localAutonomousSystemNumber,omitempty"`  // BGP Local autonomous system number of the fabric border device.
	IsDefaultExit                *bool  `json:"isDefaultExit,omitempty"`                // Is default exit value of the fabric border device.
	ImportExternalRoutes         *bool  `json:"importExternalRoutes,omitempty"`         // Import external routes value of the fabric border device.
	BorderPriority               *int   `json:"borderPriority,omitempty"`               // Border priority of the fabric border device.  A lower value indicates higher priority. E.g., a priority of 1 takes precedence over 5.
	PrependAutonomousSystemCount *int   `json:"prependAutonomousSystemCount,omitempty"` // Prepend autonomous system count of the fabric border device.
}
type ResponseSdaUpdateFabricDevicesV1 struct {
	Response *ResponseSdaUpdateFabricDevicesV1Response `json:"response,omitempty"` //
	Version  string                                    `json:"version,omitempty"`  // Version number.
}
type ResponseSdaUpdateFabricDevicesV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaDeleteFabricDevicesV1 struct {
	Response *ResponseSdaDeleteFabricDevicesV1Response `json:"response,omitempty"` //
	Version  string                                    `json:"version,omitempty"`  // Version number.
}
type ResponseSdaDeleteFabricDevicesV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaAddFabricDevicesV1 struct {
	Response *ResponseSdaAddFabricDevicesV1Response `json:"response,omitempty"` //
	Version  string                                 `json:"version,omitempty"`  // Version number.
}
type ResponseSdaAddFabricDevicesV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaGetFabricDevicesCountV1 struct {
	Response *ResponseSdaGetFabricDevicesCountV1Response `json:"response,omitempty"` //
	Version  string                                      `json:"version,omitempty"`  // Version number.
}
type ResponseSdaGetFabricDevicesCountV1Response struct {
	Count *int `json:"count,omitempty"` // Number of fabric devices.
}
type ResponseSdaDeleteFabricDeviceLayer2HandoffsV1 struct {
	Response *ResponseSdaDeleteFabricDeviceLayer2HandoffsV1Response `json:"response,omitempty"` //
	Version  string                                                 `json:"version,omitempty"`  // Version number.
}
type ResponseSdaDeleteFabricDeviceLayer2HandoffsV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaGetFabricDevicesLayer2HandoffsV1 struct {
	Response *[]ResponseSdaGetFabricDevicesLayer2HandoffsV1Response `json:"response,omitempty"` //
	Version  string                                                 `json:"version,omitempty"`  // Version number.
}
type ResponseSdaGetFabricDevicesLayer2HandoffsV1Response struct {
	ID              string `json:"id,omitempty"`              // ID of the layer 2 handoff of a fabric device.
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Network device ID of the fabric device.
	FabricID        string `json:"fabricId,omitempty"`        // ID of the fabric this device is assigned to.
	InterfaceName   string `json:"interfaceName,omitempty"`   // Interface name of the layer 2 handoff.
	InternalVLANID  *int   `json:"internalVlanId,omitempty"`  // VLAN number associated with this fabric.
	ExternalVLANID  *int   `json:"externalVlanId,omitempty"`  // External VLAN number into which the fabric is extended.
}
type ResponseSdaAddFabricDevicesLayer2HandoffsV1 struct {
	Response *ResponseSdaAddFabricDevicesLayer2HandoffsV1Response `json:"response,omitempty"` //
	Version  string                                               `json:"version,omitempty"`  // Version number.
}
type ResponseSdaAddFabricDevicesLayer2HandoffsV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaGetFabricDevicesLayer2HandoffsCountV1 struct {
	Response *ResponseSdaGetFabricDevicesLayer2HandoffsCountV1Response `json:"response,omitempty"` //
	Version  string                                                    `json:"version,omitempty"`  // Version number.
}
type ResponseSdaGetFabricDevicesLayer2HandoffsCountV1Response struct {
	Count *int `json:"count,omitempty"` // Number of fabric device layer 2 handoffs.
}
type ResponseSdaDeleteFabricDeviceLayer2HandoffByIDV1 struct {
	Response *ResponseSdaDeleteFabricDeviceLayer2HandoffByIDV1Response `json:"response,omitempty"` //
	Version  string                                                    `json:"version,omitempty"`  // Version number.
}
type ResponseSdaDeleteFabricDeviceLayer2HandoffByIDV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaAddFabricDevicesLayer3HandoffsWithIPTransitV1 struct {
	Response *ResponseSdaAddFabricDevicesLayer3HandoffsWithIPTransitV1Response `json:"response,omitempty"` //
	Version  string                                                            `json:"version,omitempty"`  // Version number.
}
type ResponseSdaAddFabricDevicesLayer3HandoffsWithIPTransitV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaUpdateFabricDevicesLayer3HandoffsWithIPTransitV1 struct {
	Response *ResponseSdaUpdateFabricDevicesLayer3HandoffsWithIPTransitV1Response `json:"response,omitempty"` //
	Version  string                                                               `json:"version,omitempty"`  // Version number.
}
type ResponseSdaUpdateFabricDevicesLayer3HandoffsWithIPTransitV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaDeleteFabricDeviceLayer3HandoffsWithIPTransitV1 struct {
	Response *ResponseSdaDeleteFabricDeviceLayer3HandoffsWithIPTransitV1Response `json:"response,omitempty"` //
	Version  string                                                              `json:"version,omitempty"`  // Version number.
}
type ResponseSdaDeleteFabricDeviceLayer3HandoffsWithIPTransitV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaGetFabricDevicesLayer3HandoffsWithIPTransitV1 struct {
	Response *[]ResponseSdaGetFabricDevicesLayer3HandoffsWithIPTransitV1Response `json:"response,omitempty"` //
	Version  string                                                              `json:"version,omitempty"`  // Version number.
}
type ResponseSdaGetFabricDevicesLayer3HandoffsWithIPTransitV1Response struct {
	ID                             string `json:"id,omitempty"`                             // ID of the fabric device layer 3 handoff ip transit.
	NetworkDeviceID                string `json:"networkDeviceId,omitempty"`                // Network device ID of the fabric device.
	FabricID                       string `json:"fabricId,omitempty"`                       // ID of the fabric this device is assigned to.
	TransitNetworkID               string `json:"transitNetworkId,omitempty"`               // ID of the transit network of the layer 3 handoff ip transit.
	InterfaceName                  string `json:"interfaceName,omitempty"`                  // Interface name of the layer 3 handoff ip transit.
	ExternalConnectivityIPPoolName string `json:"externalConnectivityIpPoolName,omitempty"` // External connectivity ip pool is used by Catalyst Center to allocate IP address for the connection between the border node and peer.
	VirtualNetworkName             string `json:"virtualNetworkName,omitempty"`             // Name of the virtual network associated with this fabric site.
	VLANID                         *int   `json:"vlanId,omitempty"`                         // VLAN number for the Switch Virtual Interface (SVI) used to establish BGP peering with the external domain for the virtual network. Allowed VLAN range is 2-4094 except for reserved vlans (1, 1002-1005, 2046, 4094).
	TCPMssAdjustment               *int   `json:"tcpMssAdjustment,omitempty"`               // TCP maximum segment size (mss) value for the layer 3 handoff. Allowed range is [500-1440]. TCP MSS Adjustment value is applicable for the TCP sessions over both IPv4 and IPv6.
	LocalIPAddress                 string `json:"localIpAddress,omitempty"`                 // Local ipv4 address for the selected virtual network. IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). Not applicable if an external connectivity ip pool name is present.
	RemoteIPAddress                string `json:"remoteIpAddress,omitempty"`                // Remote ipv4 address for the selected virtual network. IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). Not applicable if an external connectivity ip pool name is present.
	LocalIPv6Address               string `json:"localIpv6Address,omitempty"`               // Local ipv6 address for the selected virtual network. IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). Not applicable if an external connectivity ip pool name is present.
	RemoteIPv6Address              string `json:"remoteIpv6Address,omitempty"`              // Remote ipv6 address for the selected virtual network. IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). Not applicable if an external connectivity ip pool name is present.
}
type ResponseSdaGetFabricDevicesLayer3HandoffsWithIPTransitCountV1 struct {
	Response *ResponseSdaGetFabricDevicesLayer3HandoffsWithIPTransitCountV1Response `json:"response,omitempty"` //
	Version  string                                                                 `json:"version,omitempty"`  // Version number.
}
type ResponseSdaGetFabricDevicesLayer3HandoffsWithIPTransitCountV1Response struct {
	Count *int `json:"count,omitempty"` // Number of fabric device layer 3 handoffs with IP transit.
}
type ResponseSdaDeleteFabricDeviceLayer3HandoffWithIPTransitByIDV1 struct {
	Response *ResponseSdaDeleteFabricDeviceLayer3HandoffWithIPTransitByIDV1Response `json:"response,omitempty"` //
	Version  string                                                                 `json:"version,omitempty"`  // Version number.
}
type ResponseSdaDeleteFabricDeviceLayer3HandoffWithIPTransitByIDV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransitV1 struct {
	Response *ResponseSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransitV1Response `json:"response,omitempty"` //
	Version  string                                                                `json:"version,omitempty"`  // Version number.
}
type ResponseSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransitV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaGetFabricDevicesLayer3HandoffsWithSdaTransitV1 struct {
	Response *[]ResponseSdaGetFabricDevicesLayer3HandoffsWithSdaTransitV1Response `json:"response,omitempty"` //
	Version  string                                                               `json:"version,omitempty"`  // Version number.
}
type ResponseSdaGetFabricDevicesLayer3HandoffsWithSdaTransitV1Response struct {
	NetworkDeviceID               string `json:"networkDeviceId,omitempty"`               // Network device ID of the fabric device.
	FabricID                      string `json:"fabricId,omitempty"`                      // ID of the fabric this device is assigned to.
	TransitNetworkID              string `json:"transitNetworkId,omitempty"`              // ID of the transit network of the layer 3 handoff sda transit.
	AffinityIDPrime               *int   `json:"affinityIdPrime,omitempty"`               // Affinity id prime value of the border node. It supersedes the border priority to determine border node preference. Allowed range is [0-2147483647]. The lower the relative value of affinity id prime, the higher the preference for a destination border node.
	AffinityIDDecider             *int   `json:"affinityIdDecider,omitempty"`             // Affinity id decider value of the border node. When the affinity id prime value is the same on multiple devices, the affinity id decider value is used as a tiebreaker. Allowed range is [0-2147483647]. The lower the relative value of affinity id decider, the higher the preference for a destination border node.
	ConnectedToInternet           *bool  `json:"connectedToInternet,omitempty"`           // True value for this allows associated site to provide internet access to other sites through sd-access.
	IsMulticastOverTransitEnabled *bool  `json:"isMulticastOverTransitEnabled,omitempty"` // True value for this configures native multicast over multiple sites that are connected to an sd-access transit.
}
type ResponseSdaDeleteFabricDeviceLayer3HandoffsWithSdaTransitV1 struct {
	Response *ResponseSdaDeleteFabricDeviceLayer3HandoffsWithSdaTransitV1Response `json:"response,omitempty"` //
	Version  string                                                               `json:"version,omitempty"`  // Version number.
}
type ResponseSdaDeleteFabricDeviceLayer3HandoffsWithSdaTransitV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaAddFabricDevicesLayer3HandoffsWithSdaTransitV1 struct {
	Response *ResponseSdaAddFabricDevicesLayer3HandoffsWithSdaTransitV1Response `json:"response,omitempty"` //
	Version  string                                                             `json:"version,omitempty"`  // Version number.
}
type ResponseSdaAddFabricDevicesLayer3HandoffsWithSdaTransitV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaGetFabricDevicesLayer3HandoffsWithSdaTransitCountV1 struct {
	Response *ResponseSdaGetFabricDevicesLayer3HandoffsWithSdaTransitCountV1Response `json:"response,omitempty"` //
	Version  string                                                                  `json:"version,omitempty"`  // Version number.
}
type ResponseSdaGetFabricDevicesLayer3HandoffsWithSdaTransitCountV1Response struct {
	Count *int `json:"count,omitempty"` // Number of fabric device layer 3 handoffs with sda transit.
}
type ResponseSdaDeleteFabricDeviceByIDV1 struct {
	Response *ResponseSdaDeleteFabricDeviceByIDV1Response `json:"response,omitempty"` //
	Version  string                                       `json:"version,omitempty"`  // Version number.
}
type ResponseSdaDeleteFabricDeviceByIDV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaGetFabricSitesV1 struct {
	Response *[]ResponseSdaGetFabricSitesV1Response `json:"response,omitempty"` //
	Version  string                                 `json:"version,omitempty"`  // Version number.
}
type ResponseSdaGetFabricSitesV1Response struct {
	ID                        string `json:"id,omitempty"`                        // ID of the fabric site.
	SiteID                    string `json:"siteId,omitempty"`                    // ID of the network hierarchy.
	AuthenticationProfileName string `json:"authenticationProfileName,omitempty"` // Authentication profile used for this fabric.
	IsPubSubEnabled           *bool  `json:"isPubSubEnabled,omitempty"`           // Specifies whether this fabric site will use pub/sub for control nodes.
}
type ResponseSdaAddFabricSiteV1 struct {
	Response *ResponseSdaAddFabricSiteV1Response `json:"response,omitempty"` //
	Version  string                              `json:"version,omitempty"`  // Version number.
}
type ResponseSdaAddFabricSiteV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaUpdateFabricSiteV1 struct {
	Response *ResponseSdaUpdateFabricSiteV1Response `json:"response,omitempty"` //
	Version  string                                 `json:"version,omitempty"`  // Version number.
}
type ResponseSdaUpdateFabricSiteV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaGetFabricSiteCountV1 struct {
	Response *ResponseSdaGetFabricSiteCountV1Response `json:"response,omitempty"` //
	Version  string                                   `json:"version,omitempty"`  // Version number.
}
type ResponseSdaGetFabricSiteCountV1Response struct {
	Count *int `json:"count,omitempty"` // The number of fabric sites.
}
type ResponseSdaDeleteFabricSiteByIDV1 struct {
	Response *ResponseSdaDeleteFabricSiteByIDV1Response `json:"response,omitempty"` //
	Version  string                                     `json:"version,omitempty"`  // Version number.
}
type ResponseSdaDeleteFabricSiteByIDV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaGetFabricZonesV1 struct {
	Response *[]ResponseSdaGetFabricZonesV1Response `json:"response,omitempty"` //
	Version  string                                 `json:"version,omitempty"`  // Version number.
}
type ResponseSdaGetFabricZonesV1Response struct {
	ID                        string `json:"id,omitempty"`                        // ID of the fabric zone.
	SiteID                    string `json:"siteId,omitempty"`                    // ID of the network hierarchy.
	AuthenticationProfileName string `json:"authenticationProfileName,omitempty"` // Authentication profile used for this fabric.
}
type ResponseSdaUpdateFabricZoneV1 struct {
	Response *ResponseSdaUpdateFabricZoneV1Response `json:"response,omitempty"` //
	Version  string                                 `json:"version,omitempty"`  // Version number.
}
type ResponseSdaUpdateFabricZoneV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaAddFabricZoneV1 struct {
	Response *ResponseSdaAddFabricZoneV1Response `json:"response,omitempty"` //
	Version  string                              `json:"version,omitempty"`  // Version number.
}
type ResponseSdaAddFabricZoneV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaGetFabricZoneCountV1 struct {
	Response *ResponseSdaGetFabricZoneCountV1Response `json:"response,omitempty"` //
	Version  string                                   `json:"version,omitempty"`  // Version number.
}
type ResponseSdaGetFabricZoneCountV1Response struct {
	Count *int `json:"count,omitempty"` // The number of fabric zones.
}
type ResponseSdaDeleteFabricZoneByIDV1 struct {
	Response *ResponseSdaDeleteFabricZoneByIDV1Response `json:"response,omitempty"` //
	Version  string                                     `json:"version,omitempty"`  // Version number.
}
type ResponseSdaDeleteFabricZoneByIDV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaAddLayer2VirtualNetworksV1 struct {
	Response *ResponseSdaAddLayer2VirtualNetworksV1Response `json:"response,omitempty"` //
	Version  string                                         `json:"version,omitempty"`  // Version number.
}
type ResponseSdaAddLayer2VirtualNetworksV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaDeleteLayer2VirtualNetworksV1 struct {
	Response *ResponseSdaDeleteLayer2VirtualNetworksV1Response `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  // Version number.
}
type ResponseSdaDeleteLayer2VirtualNetworksV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaGetLayer2VirtualNetworksV1 struct {
	Response *[]ResponseSdaGetLayer2VirtualNetworksV1Response `json:"response,omitempty"` //
	Version  string                                           `json:"version,omitempty"`  // Version number.
}
type ResponseSdaGetLayer2VirtualNetworksV1Response struct {
	ID                                 string `json:"id,omitempty"`                                 // ID of the layer 2 virtual network.
	FabricID                           string `json:"fabricId,omitempty"`                           // ID of the fabric this layer 2 virtual network is assigned to.
	VLANName                           string `json:"vlanName,omitempty"`                           // Name of the VLAN of the layer 2 virtual network.
	VLANID                             *int   `json:"vlanId,omitempty"`                             // ID of the VLAN of the layer 2 virtual network.
	TrafficType                        string `json:"trafficType,omitempty"`                        // The type of traffic that is served.
	IsFabricEnabledWireless            *bool  `json:"isFabricEnabledWireless,omitempty"`            // Set to true to enable wireless.
	AssociatedLayer3VirtualNetworkName string `json:"associatedLayer3VirtualNetworkName,omitempty"` // Name of the layer 3 virtual network associated with the layer 2 virtual network. This field is provided to support requests related to virtual network anchoring.
}
type ResponseSdaUpdateLayer2VirtualNetworksV1 struct {
	Response *ResponseSdaUpdateLayer2VirtualNetworksV1Response `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  // Version number.
}
type ResponseSdaUpdateLayer2VirtualNetworksV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaGetLayer2VirtualNetworkCountV1 struct {
	Response *ResponseSdaGetLayer2VirtualNetworkCountV1Response `json:"response,omitempty"` //
	Version  string                                             `json:"version,omitempty"`  // Version number.
}
type ResponseSdaGetLayer2VirtualNetworkCountV1Response struct {
	Count *int `json:"count,omitempty"` // The number of layer 2 virtual networks
}
type ResponseSdaDeleteLayer2VirtualNetworkByIDV1 struct {
	Response *ResponseSdaDeleteLayer2VirtualNetworkByIDV1Response `json:"response,omitempty"` //
	Version  string                                               `json:"version,omitempty"`  // Version number.
}
type ResponseSdaDeleteLayer2VirtualNetworkByIDV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaAddLayer3VirtualNetworksV1 struct {
	Response *ResponseSdaAddLayer3VirtualNetworksV1Response `json:"response,omitempty"` //
	Version  string                                         `json:"version,omitempty"`  // Version number.
}
type ResponseSdaAddLayer3VirtualNetworksV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaGetLayer3VirtualNetworksV1 struct {
	Response *[]ResponseSdaGetLayer3VirtualNetworksV1Response `json:"response,omitempty"` //
	Version  string                                           `json:"version,omitempty"`  // Version number.
}
type ResponseSdaGetLayer3VirtualNetworksV1Response struct {
	ID                 string   `json:"id,omitempty"`                 // ID of the layer 3 virtual network.
	VirtualNetworkName string   `json:"virtualNetworkName,omitempty"` // Name of the layer 3 virtual network.
	FabricIDs          []string `json:"fabricIds,omitempty"`          // IDs of the fabrics this layer 3 virtual network is assigned to.
	AnchoredSiteID     string   `json:"anchoredSiteId,omitempty"`     // Fabric ID of the fabric site this layer 3 virtual network is anchored at.
}
type ResponseSdaDeleteLayer3VirtualNetworksV1 struct {
	Response *ResponseSdaDeleteLayer3VirtualNetworksV1Response `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  // Version number.
}
type ResponseSdaDeleteLayer3VirtualNetworksV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaUpdateLayer3VirtualNetworksV1 struct {
	Response *ResponseSdaUpdateLayer3VirtualNetworksV1Response `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  // Version number.
}
type ResponseSdaUpdateLayer3VirtualNetworksV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaGetLayer3VirtualNetworksCountV1 struct {
	Response *ResponseSdaGetLayer3VirtualNetworksCountV1Response `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  // Version number.
}
type ResponseSdaGetLayer3VirtualNetworksCountV1Response struct {
	Count *int `json:"count,omitempty"` // Number of layer 3 virtual networks.
}
type ResponseSdaDeleteLayer3VirtualNetworkByIDV1 struct {
	Response *ResponseSdaDeleteLayer3VirtualNetworkByIDV1Response `json:"response,omitempty"` //
	Version  string                                               `json:"version,omitempty"`  // Version number.
}
type ResponseSdaDeleteLayer3VirtualNetworkByIDV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaUpdateMulticastV1 struct {
	Response *ResponseSdaUpdateMulticastV1Response `json:"response,omitempty"` //
	Version  string                                `json:"version,omitempty"`  // Version number.
}
type ResponseSdaUpdateMulticastV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaGetMulticastV1 struct {
	Response *[]ResponseSdaGetMulticastV1Response `json:"response,omitempty"` //
	Version  string                               `json:"version,omitempty"`  // Version number.
}
type ResponseSdaGetMulticastV1Response struct {
	FabricID        string `json:"fabricId,omitempty"`        // ID of the fabric site.
	ReplicationMode string `json:"replicationMode,omitempty"` // Replication Mode deployed in the fabric site.
}
type ResponseSdaAddMulticastVirtualNetworksV1 struct {
	Response *ResponseSdaAddMulticastVirtualNetworksV1Response `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  // Version number.
}
type ResponseSdaAddMulticastVirtualNetworksV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaGetMulticastVirtualNetworksV1 struct {
	Response *[]ResponseSdaGetMulticastVirtualNetworksV1Response `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  // Version number.
}
type ResponseSdaGetMulticastVirtualNetworksV1Response struct {
	ID                 string                                                          `json:"id,omitempty"`                 // ID of the multicast configuration.
	FabricID           string                                                          `json:"fabricId,omitempty"`           // ID of the fabric site.
	VirtualNetworkName string                                                          `json:"virtualNetworkName,omitempty"` // Name of the virtual network.
	IPPoolName         string                                                          `json:"ipPoolName,omitempty"`         // Name of the IP Pool.
	IPv4SsmRanges      []string                                                        `json:"ipv4SsmRanges,omitempty"`      // IPv4 Source Specific Multicast (SSM) ranges. Allowed ranges are from 225.0.0.0/8 to 239.0.0.0/8. SSM ranges should not conflict with ranges provided for ASM multicast.
	MulticastRPs       *[]ResponseSdaGetMulticastVirtualNetworksV1ResponseMulticastRPs `json:"multicastRPs,omitempty"`       //
}
type ResponseSdaGetMulticastVirtualNetworksV1ResponseMulticastRPs struct {
	RpDeviceLocation string   `json:"rpDeviceLocation,omitempty"` // Device location of the RP.
	IPv4Address      string   `json:"ipv4Address,omitempty"`      // IPv4 address of the RP. For external RP configuration, exactly one of ipv4Address or ipv6Address must be provided. For fabric RP, this address is allocated by SDA and should not be provided during RP creation request and SDA allocated address should be retained in subsequent requests.
	IPv6Address      string   `json:"ipv6Address,omitempty"`      // IPv6 address of the RP. For external RP configuration, exactly one of ipv4Address or ipv6Address must be provided. For fabric RP, this address is allocated by SDA and should not be provided during RP creation request and SDA allocated address should be retained in subsequent requests. ipv6Address can only be provided for virtual networks with dual stack (IPv4 + IPv6) multicast pool.
	IsDefaultV4RP    *bool    `json:"isDefaultV4RP,omitempty"`    // Specifies whether it is a default IPv4 RP.
	IsDefaultV6RP    *bool    `json:"isDefaultV6RP,omitempty"`    // Specifies whether it is a default IPv6 RP.
	NetworkDeviceIDs []string `json:"networkDeviceIds,omitempty"` // IDs of the network devices. This is a required field for fabric RPs. There can be maximum of two fabric RPs for a fabric site and these are shared across all multicast virtual networks. For configuring two fabric RPs in a fabric site both devices must have border roles. Only one RP can be configured in scenarios where a fabric edge device is used as RP or a dual stack multicast pool is used.
	IPv4AsmRanges    []string `json:"ipv4AsmRanges,omitempty"`    // IPv4 Any Source Multicast ranges. Comma seperated list of IPv4 multicast group ranges that will be served by a given Multicast RP. Only IPv4 ranges can be provided. For fabric RP, both IPv4 and IPv6 ranges can be provided together. For external RP, IPv4 ranges should be provided for IPv4 external RP and IPv6 ranges should be provided for IPv6 external RP.
	IPv6AsmRanges    []string `json:"ipv6AsmRanges,omitempty"`    // IPv6 Any Source Multicast ranges. Comma seperated list of IPv6 multicast group ranges that will be served by a given Multicast RP. Only IPv6 ranges can be provided. IPv6 ranges can only be provided for dual stack multicast pool. For fabric RP, both IPv4 and IPv6 ranges can be provided together. For external RP, IPv4 ranges should be provided for IPv4 external RP and IPv6 ranges should be provided for IPv6 external RP.
}
type ResponseSdaUpdateMulticastVirtualNetworksV1 struct {
	Response *ResponseSdaUpdateMulticastVirtualNetworksV1Response `json:"response,omitempty"` //
	Version  string                                               `json:"version,omitempty"`  // Version number.
}
type ResponseSdaUpdateMulticastVirtualNetworksV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaGetMulticastVirtualNetworkCountV1 struct {
	Response *ResponseSdaGetMulticastVirtualNetworkCountV1Response `json:"response,omitempty"` //
	Version  string                                                `json:"version,omitempty"`  // Version number.
}
type ResponseSdaGetMulticastVirtualNetworkCountV1Response struct {
	Count *int `json:"count,omitempty"` // Number of multicast configurations.
}
type ResponseSdaDeleteMulticastVirtualNetworkByIDV1 struct {
	Response *ResponseSdaDeleteMulticastVirtualNetworkByIDV1Response `json:"response,omitempty"` //
	Version  string                                                  `json:"version,omitempty"`  // Version number.
}
type ResponseSdaDeleteMulticastVirtualNetworkByIDV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaAddPortAssignmentsV1 struct {
	Response *ResponseSdaAddPortAssignmentsV1Response `json:"response,omitempty"` //
	Version  string                                   `json:"version,omitempty"`  // Version number.
}
type ResponseSdaAddPortAssignmentsV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaGetPortAssignmentsV1 struct {
	Response *[]ResponseSdaGetPortAssignmentsV1Response `json:"response,omitempty"` //
	Version  string                                     `json:"version,omitempty"`  // Version number.
}
type ResponseSdaGetPortAssignmentsV1Response struct {
	ID                       string `json:"id,omitempty"`                       // ID of the port assignment.
	FabricID                 string `json:"fabricId,omitempty"`                 // ID of the fabric the device is assigned to.
	NetworkDeviceID          string `json:"networkDeviceId,omitempty"`          // Network device ID of the port assignment.
	InterfaceName            string `json:"interfaceName,omitempty"`            // Interface name of the port assignment.
	ConnectedDeviceType      string `json:"connectedDeviceType,omitempty"`      // Connected device type of the port assignment.
	DataVLANName             string `json:"dataVlanName,omitempty"`             // Data VLAN name of the port assignment.
	VoiceVLANName            string `json:"voiceVlanName,omitempty"`            // Voice VLAN name of the port assignment.
	AuthenticateTemplateName string `json:"authenticateTemplateName,omitempty"` // Authenticate template name of the port assignment.
	SecurityGroupName        string `json:"securityGroupName,omitempty"`        // Security group name of the port assignment.
	InterfaceDescription     string `json:"interfaceDescription,omitempty"`     // Interface description of the port assignment.
}
type ResponseSdaUpdatePortAssignmentsV1 struct {
	Response *ResponseSdaUpdatePortAssignmentsV1Response `json:"response,omitempty"` //
	Version  string                                      `json:"version,omitempty"`  // Version number.
}
type ResponseSdaUpdatePortAssignmentsV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaDeletePortAssignmentsV1 struct {
	Response *ResponseSdaDeletePortAssignmentsV1Response `json:"response,omitempty"` //
	Version  string                                      `json:"version,omitempty"`  // Version number.
}
type ResponseSdaDeletePortAssignmentsV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaGetPortAssignmentCountV1 struct {
	Response *ResponseSdaGetPortAssignmentCountV1Response `json:"response,omitempty"` //
	Version  string                                       `json:"version,omitempty"`  // Version number.
}
type ResponseSdaGetPortAssignmentCountV1Response struct {
	Count *int `json:"count,omitempty"` // Number of port assignments.
}
type ResponseSdaDeletePortAssignmentByIDV1 struct {
	Response *ResponseSdaDeletePortAssignmentByIDV1Response `json:"response,omitempty"` //
	Version  string                                         `json:"version,omitempty"`  // Version number.
}
type ResponseSdaDeletePortAssignmentByIDV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaGetPortChannelsV1 struct {
	Response *[]ResponseSdaGetPortChannelsV1Response `json:"response,omitempty"` //
	Version  string                                  `json:"version,omitempty"`  // Version number.
}
type ResponseSdaGetPortChannelsV1Response struct {
	ID                  string   `json:"id,omitempty"`                  // ID of the port channel.
	FabricID            string   `json:"fabricId,omitempty"`            // ID of the fabric the device is assigned to.
	NetworkDeviceID     string   `json:"networkDeviceId,omitempty"`     // ID of the network device.
	PortChannelName     string   `json:"portChannelName,omitempty"`     // Name of the port channel.
	InterfaceNames      []string `json:"interfaceNames,omitempty"`      // Interface names of this port channel.
	ConnectedDeviceType string   `json:"connectedDeviceType,omitempty"` // Connected device type of the port channel.
	Protocol            string   `json:"protocol,omitempty"`            // Protocol of the port channel.
	Description         string   `json:"description,omitempty"`         // Description of the port channel.
}
type ResponseSdaAddPortChannelsV1 struct {
	Response *ResponseSdaAddPortChannelsV1Response `json:"response,omitempty"` //
	Version  string                                `json:"version,omitempty"`  // Version number.
}
type ResponseSdaAddPortChannelsV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaUpdatePortChannelsV1 struct {
	Response *ResponseSdaUpdatePortChannelsV1Response `json:"response,omitempty"` //
	Version  string                                   `json:"version,omitempty"`  // Version number.
}
type ResponseSdaUpdatePortChannelsV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaDeletePortChannelsV1 struct {
	Response *ResponseSdaDeletePortChannelsV1Response `json:"response,omitempty"` //
	Version  string                                   `json:"version,omitempty"`  // Version number.
}
type ResponseSdaDeletePortChannelsV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaGetPortChannelCountV1 struct {
	Response *ResponseSdaGetPortChannelCountV1Response `json:"response,omitempty"` //
	Version  string                                    `json:"version,omitempty"`  // Version number.
}
type ResponseSdaGetPortChannelCountV1Response struct {
	Count *int `json:"count,omitempty"` // Number of port channels.
}
type ResponseSdaDeletePortChannelByIDV1 struct {
	Response *ResponseSdaDeletePortChannelByIDV1Response `json:"response,omitempty"` //
	Version  string                                      `json:"version,omitempty"`  // Version number.
}
type ResponseSdaDeletePortChannelByIDV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaDeleteProvisionedDevicesV1 struct {
	Response *ResponseSdaDeleteProvisionedDevicesV1Response `json:"response,omitempty"` //
	Version  string                                         `json:"version,omitempty"`  // Version number.
}
type ResponseSdaDeleteProvisionedDevicesV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaProvisionDevicesV1 struct {
	Response *ResponseSdaProvisionDevicesV1Response `json:"response,omitempty"` //
	Version  string                                 `json:"version,omitempty"`  // Version number.
}
type ResponseSdaProvisionDevicesV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaGetProvisionedDevicesV1 struct {
	Response *[]ResponseSdaGetProvisionedDevicesV1Response `json:"response,omitempty"` //
	Version  string                                        `json:"version,omitempty"`  // Version number.
}
type ResponseSdaGetProvisionedDevicesV1Response struct {
	ID              string `json:"id,omitempty"`              // ID of the provisioned device.
	SiteID          string `json:"siteId,omitempty"`          // ID of the site this device is provisioned to.
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // ID of the network device.
}
type ResponseSdaReProvisionDevicesV1 struct {
	Response *ResponseSdaReProvisionDevicesV1Response `json:"response,omitempty"` //
	Version  string                                   `json:"version,omitempty"`  // Version number.
}
type ResponseSdaReProvisionDevicesV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaGetProvisionedDevicesCountV1 struct {
	Response *ResponseSdaGetProvisionedDevicesCountV1Response `json:"response,omitempty"` //
	Version  string                                           `json:"version,omitempty"`  // Version number.
}
type ResponseSdaGetProvisionedDevicesCountV1Response struct {
	Count *int `json:"count,omitempty"` // Number of provisioned devices.
}
type ResponseSdaDeleteProvisionedDeviceByIDV1 struct {
	Response *ResponseSdaDeleteProvisionedDeviceByIDV1Response `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  // Version number.
}
type ResponseSdaDeleteProvisionedDeviceByIDV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaUpdateTransitNetworksV1 struct {
	Response *ResponseSdaUpdateTransitNetworksV1Response `json:"response,omitempty"` //
	Version  string                                      `json:"version,omitempty"`  // Version number.
}
type ResponseSdaUpdateTransitNetworksV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaGetTransitNetworksV1 struct {
	Response *[]ResponseSdaGetTransitNetworksV1Response `json:"response,omitempty"` //
	Version  string                                     `json:"version,omitempty"`  // Version number.
}
type ResponseSdaGetTransitNetworksV1Response struct {
	ID                 string                                                     `json:"id,omitempty"`                 // ID of the transit network.
	Name               string                                                     `json:"name,omitempty"`               // Name of the transit network.
	Type               string                                                     `json:"type,omitempty"`               // Type of the transit network.
	IPTransitSettings  *ResponseSdaGetTransitNetworksV1ResponseIPTransitSettings  `json:"ipTransitSettings,omitempty"`  //
	SdaTransitSettings *ResponseSdaGetTransitNetworksV1ResponseSdaTransitSettings `json:"sdaTransitSettings,omitempty"` //
}
type ResponseSdaGetTransitNetworksV1ResponseIPTransitSettings struct {
	RoutingProtocolName    string `json:"routingProtocolName,omitempty"`    // Routing Protocol Name of the IP transit network.
	AutonomousSystemNumber string `json:"autonomousSystemNumber,omitempty"` // Autonomous System Number of the IP transit network. Allowed range is [1 to 4294967295].
}
type ResponseSdaGetTransitNetworksV1ResponseSdaTransitSettings struct {
	IsMulticastOverTransitEnabled *bool    `json:"isMulticastOverTransitEnabled,omitempty"` // This indicates that multicast is enabled over SD-Access Transit. This supports Native Multicast over SD-Access Transit. This is only applicable for transit of type SDA_LISP_PUB_SUB_TRANSIT.
	ControlPlaneNetworkDeviceIDs  []string `json:"controlPlaneNetworkDeviceIds,omitempty"`  // List of network device IDs that are used as control plane nodes.
}
type ResponseSdaAddTransitNetworksV1 struct {
	Response *ResponseSdaAddTransitNetworksV1Response `json:"response,omitempty"` //
	Version  string                                   `json:"version,omitempty"`  // Version number.
}
type ResponseSdaAddTransitNetworksV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaGetTransitNetworksCountV1 struct {
	Response *ResponseSdaGetTransitNetworksCountV1Response `json:"response,omitempty"` //
	Version  string                                        `json:"version,omitempty"`  // Version number.
}
type ResponseSdaGetTransitNetworksCountV1Response struct {
	Count *int `json:"count,omitempty"` // Number of transit networks.
}
type ResponseSdaDeleteTransitNetworkByIDV1 struct {
	Response *ResponseSdaDeleteTransitNetworkByIDV1Response `json:"response,omitempty"` //
	Version  string                                         `json:"version,omitempty"`  // Version number.
}
type ResponseSdaDeleteTransitNetworkByIDV1Response struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.
	URL    string `json:"url,omitempty"`    // Task status lookup url.
}
type ResponseSdaAddVirtualNetworkWithScalableGroupsV1 struct {
	Status             string `json:"status,omitempty"`             // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.
	Description        string `json:"description,omitempty"`        // provides detailed information for API success or failure.
	TaskID             string `json:"taskId,omitempty"`             // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available
	TaskStatusURL      string `json:"taskStatusUrl,omitempty"`      // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>
	ExecutionID        string `json:"executionId,omitempty"`        // uuid for API execution status
}
type ResponseSdaDeleteVirtualNetworkWithScalableGroupsV1 struct {
	Status             string `json:"status,omitempty"`             // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.
	Description        string `json:"description,omitempty"`        // provides detailed information for API success or failure.
	TaskID             string `json:"taskId,omitempty"`             // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available
	TaskStatusURL      string `json:"taskStatusUrl,omitempty"`      // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>
	ExecutionID        string `json:"executionId,omitempty"`        // uuid for API execution status
}
type ResponseSdaGetVirtualNetworkWithScalableGroupsV1 struct {
	VirtualNetworkName      string   `json:"virtualNetworkName,omitempty"`      // Virtual Network Name to be assigned at global level
	IsGuestVirtualNetwork   *bool    `json:"isGuestVirtualNetwork,omitempty"`   // Guest Virtual Network
	ScalableGroupNames      []string `json:"scalableGroupNames,omitempty"`      // Scalable Group Names
	VManageVpnID            string   `json:"vManageVpnId,omitempty"`            // vManage vpn id for SD-WAN
	VirtualNetworkContextID string   `json:"virtualNetworkContextId,omitempty"` // Virtual Network Context Id for Global Virtual Network
	Status                  string   `json:"status,omitempty"`                  // Status
	Description             string   `json:"description,omitempty"`             // Virtual network info retrieved successfully
	ExecutionID             string   `json:"executionId,omitempty"`             // Execution Id
}
type ResponseSdaUpdateVirtualNetworkWithScalableGroupsV1 struct {
	Status             string `json:"status,omitempty"`             // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.
	Description        string `json:"description,omitempty"`        // provides detailed information for API success or failure.
	TaskID             string `json:"taskId,omitempty"`             // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available
	TaskStatusURL      string `json:"taskStatusUrl,omitempty"`      // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>
	ExecutionID        string `json:"executionId,omitempty"`        // uuid for API execution status
}
type RequestSdaAddDefaultAuthenticationTemplateInSdaFabricV1 []RequestItemSdaAddDefaultAuthenticationTemplateInSdaFabricV1 // Array of RequestSdaAddDefaultAuthenticationTemplateInSDAFabricV1
type RequestItemSdaAddDefaultAuthenticationTemplateInSdaFabricV1 struct {
	SiteNameHierarchy        string `json:"siteNameHierarchy,omitempty"`        // Path of sda Fabric Site
	AuthenticateTemplateName string `json:"authenticateTemplateName,omitempty"` // Authenticate Template Name
}
type RequestSdaUpdateDefaultAuthenticationProfileInSdaFabricV1 []RequestItemSdaUpdateDefaultAuthenticationProfileInSdaFabricV1 // Array of RequestSdaUpdateDefaultAuthenticationProfileInSDAFabricV1
type RequestItemSdaUpdateDefaultAuthenticationProfileInSdaFabricV1 struct {
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         // Path of sda Fabric Site
	AuthenticateTemplateName  string `json:"authenticateTemplateName,omitempty"`  // Authenticate Template Name
	AuthenticationOrder       string `json:"authenticationOrder,omitempty"`       // Authentication Order
	Dot1XToMabFallbackTimeout string `json:"dot1xToMabFallbackTimeout,omitempty"` // Dot1x To MabFallback Timeout( Allowed range is [3-120])
	WakeOnLan                 *bool  `json:"wakeOnLan,omitempty"`                 // Wake On Lan
	NumberOfHosts             string `json:"numberOfHosts,omitempty"`             // Number Of Hosts
}
type RequestSdaAddBorderDeviceInSdaFabricV1 []RequestItemSdaAddBorderDeviceInSdaFabricV1 // Array of RequestSdaAddBorderDeviceInSDAFabricV1
type RequestItemSdaAddBorderDeviceInSdaFabricV1 struct {
	DeviceManagementIPAddress         string                                                                    `json:"deviceManagementIpAddress,omitempty"`         // Management Ip Address of the provisioned Device
	SiteNameHierarchy                 string                                                                    `json:"siteNameHierarchy,omitempty"`                 // Site Name Hierarchy of provisioned Device(site should be part of Fabric Site)
	DeviceRole                        []string                                                                  `json:"deviceRole,omitempty"`                        // Supported Device Roles in SD-Access fabric. Allowed roles are "Border_Node","Control_Plane_Node","Edge_Node". E.g. ["Border_Node"] or ["Border_Node", "Control_Plane_Node"] or ["Border_Node", "Control_Plane_Node","Edge_Node"]
	RouteDistributionProtocol         string                                                                    `json:"routeDistributionProtocol,omitempty"`         // Route Distribution Protocol for Control Plane Device. Allowed values are "LISP_BGP" or "LISP_PUB_SUB". Default value is "LISP_BGP"
	ExternalDomainRoutingProtocolName string                                                                    `json:"externalDomainRoutingProtocolName,omitempty"` // External Domain Routing Protocol Name
	ExternalConnectivityIPPoolName    string                                                                    `json:"externalConnectivityIpPoolName,omitempty"`    // External Connectivity IpPool Name
	InternalAutonomouSystemNumber     string                                                                    `json:"internalAutonomouSystemNumber,omitempty"`     // Internal Autonomous System Number
	BorderPriority                    string                                                                    `json:"borderPriority,omitempty"`                    // Border priority associated with a given device. Allowed range for Border Priority is [1-9]. A lower value indicates higher priority. E.g., a priority of 1 takes precedence over 5. Default priority would be set to 10.
	BorderSessionType                 string                                                                    `json:"borderSessionType,omitempty"`                 // Border Session Type
	ConnectedToInternet               *bool                                                                     `json:"connectedToInternet,omitempty"`               // Connected to Internet
	SdaTransitNetworkName             string                                                                    `json:"sdaTransitNetworkName,omitempty"`             // SD-Access Transit Network Name
	BorderWithExternalConnectivity    *bool                                                                     `json:"borderWithExternalConnectivity,omitempty"`    // Border With External Connectivity (Note: True for transit and False for non-transit border)
	ExternalConnectivitySettings      *[]RequestItemSdaAddBorderDeviceInSdaFabricV1ExternalConnectivitySettings `json:"externalConnectivitySettings,omitempty"`      //
}
type RequestItemSdaAddBorderDeviceInSdaFabricV1ExternalConnectivitySettings struct {
	InterfaceName                 string                                                                             `json:"interfaceName,omitempty"`                 // Interface Name
	InterfaceDescription          string                                                                             `json:"interfaceDescription,omitempty"`          // Interface Description
	ExternalAutonomouSystemNumber string                                                                             `json:"externalAutonomouSystemNumber,omitempty"` // External Autonomous System Number peer (e.g.,1-65535)
	L3Handoff                     *[]RequestItemSdaAddBorderDeviceInSdaFabricV1ExternalConnectivitySettingsL3Handoff `json:"l3Handoff,omitempty"`                     //
	L2Handoff                     *[]RequestItemSdaAddBorderDeviceInSdaFabricV1ExternalConnectivitySettingsL2Handoff `json:"l2Handoff,omitempty"`                     //
}
type RequestItemSdaAddBorderDeviceInSdaFabricV1ExternalConnectivitySettingsL3Handoff struct {
	VirtualNetwork *RequestItemSdaAddBorderDeviceInSdaFabricV1ExternalConnectivitySettingsL3HandoffVirtualNetwork `json:"virtualNetwork,omitempty"` //
}
type RequestItemSdaAddBorderDeviceInSdaFabricV1ExternalConnectivitySettingsL3HandoffVirtualNetwork struct {
	VirtualNetworkName string `json:"virtualNetworkName,omitempty"` // Virtual Network Name, that is associated to Fabric Site
	VLANID             string `json:"vlanId,omitempty"`             // Vlan Id (e.g.,2-4096 except for reserved VLANs (1002-1005, 2046, 4095))
}
type RequestItemSdaAddBorderDeviceInSdaFabricV1ExternalConnectivitySettingsL2Handoff struct {
	VirtualNetworkName string `json:"virtualNetworkName,omitempty"` // Virtual Network Name, that is associated to Fabric Site
	VLANName           string `json:"vlanName,omitempty"`           // Vlan Name of L2 Handoff
}
type RequestSdaAddControlPlaneDeviceInSdaFabricV1 struct {
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` // Management Ip Address of the Device which is provisioned successfully
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         // siteNameHierarchy of the Provisioned Device(site should be part of Fabric Site)
	RouteDistributionProtocol string `json:"routeDistributionProtocol,omitempty"` // Route Distribution Protocol for Control Plane Device. Allowed values are "LISP_BGP" or "LISP_PUB_SUB". Default value is "LISP_BGP"
}
type RequestSdaAddEdgeDeviceInSdaFabricV1 struct {
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` // Management Ip Address of the Device which is provisioned successfully
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         // siteNameHierarchy of the Provisioned Device(site should be part of Fabric Site)
}
type RequestSdaAddSiteInSdaFabricV1 struct {
	FabricName        string `json:"fabricName,omitempty"`        // Warning - Starting DNA Center 2.2.3.5 release, this field has been deprecated. SD-Access Fabric does not need it anymore.  It will be removed in future DNA Center releases.
	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Existing site name hierarchy available at global level. For Example "Global/Chicago/Building21/Floor1"
	FabricType        string `json:"fabricType,omitempty"`        // Type of SD-Access Fabric. Allowed values are "FABRIC_SITE" or "FABRIC_ZONE".  Default value is "FABRIC_SITE".
}
type RequestSdaAddPortAssignmentForAccessPointInSdaFabricV1 struct {
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         // Path of sda Fabric Site
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` // Management Ip Address of the edge device
	InterfaceName             string `json:"interfaceName,omitempty"`             // Interface Name of the edge device
	DataIPAddressPoolName     string `json:"dataIpAddressPoolName,omitempty"`     // Ip Pool Name, that is assigned to INFRA_VN
	AuthenticateTemplateName  string `json:"authenticateTemplateName,omitempty"`  // Authenticate TemplateName associated to Fabric Site
	InterfaceDescription      string `json:"interfaceDescription,omitempty"`      // Details or note of interface port assignment
}
type RequestSdaAddPortAssignmentForUserDeviceInSdaFabricV1 struct {
	SiteNameHierarchy         string   `json:"siteNameHierarchy,omitempty"`         // Complete Path of SD-Access Fabric Site.
	DeviceManagementIPAddress string   `json:"deviceManagementIpAddress,omitempty"` // Management Ip Address of the Edge Node Device.
	InterfaceName             string   `json:"interfaceName,omitempty"`             // Interface Name on the Edge Node Device.
	InterfaceNames            []string `json:"interfaceNames,omitempty"`            // List of Interface Names on the Edge Node Device. E.g.["GigabitEthernet1/0/3","GigabitEthernet1/0/4"]
	DataIPAddressPoolName     string   `json:"dataIpAddressPoolName,omitempty"`     // Ip Pool Name, that is assigned to virtual network with traffic type as DATA(can't be empty if voiceIpAddressPoolName is empty)
	VoiceIPAddressPoolName    string   `json:"voiceIpAddressPoolName,omitempty"`    // Ip Pool Name, that is assigned to virtual network with traffic type as VOICE(can't be empty if dataIpAddressPoolName is empty)
	AuthenticateTemplateName  string   `json:"authenticateTemplateName,omitempty"`  // Authenticate TemplateName associated with siteNameHierarchy
	ScalableGroupName         string   `json:"scalableGroupName,omitempty"`         // Scalable Group name associated with VN
	InterfaceDescription      string   `json:"interfaceDescription,omitempty"`      // User defined text message for port assignment
}
type RequestSdaAddMulticastInSdaFabricV1 struct {
	SiteNameHierarchy string                                                `json:"siteNameHierarchy,omitempty"` // Full path of sda Fabric Site
	MulticastMethod   string                                                `json:"multicastMethod,omitempty"`   // Multicast Method
	MulticastType     string                                                `json:"multicastType,omitempty"`     // Multicast Type
	MulticastVnInfo   *[]RequestSdaAddMulticastInSdaFabricV1MulticastVnInfo `json:"multicastVnInfo,omitempty"`   //
}
type RequestSdaAddMulticastInSdaFabricV1MulticastVnInfo struct {
	VirtualNetworkName  string                                                       `json:"virtualNetworkName,omitempty"`  // Virtual Network Name, that is associated to Fabric Site
	IPPoolName          string                                                       `json:"ipPoolName,omitempty"`          // Ip Pool Name, that is reserved to Fabric Site
	InternalRpIPAddress []string                                                     `json:"internalRpIpAddress,omitempty"` // InternalRpIpAddress, required if multicastType is asm_with_internal_rp
	ExternalRpIPAddress string                                                       `json:"externalRpIpAddress,omitempty"` // ExternalRpIpAddress, required if multicastType is asm_with_external_rp
	SsmInfo             *[]RequestSdaAddMulticastInSdaFabricV1MulticastVnInfoSsmInfo `json:"ssmInfo,omitempty"`             //
}
type RequestSdaAddMulticastInSdaFabricV1MulticastVnInfoSsmInfo struct {
	SsmGroupRange   string `json:"ssmGroupRange,omitempty"`   // Valid SSM group range ip address(e.g., 230.0.0.0)
	SsmWildcardMask string `json:"ssmWildcardMask,omitempty"` // Valid SSM Wildcard Mask ip address(e.g.,0.255.255.255)
}
type RequestSdaReProvisionWiredDeviceV1 struct {
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` // Management Ip Address of the device to be re-provisioned
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         // siteNameHierarchy of the provisioned device
}
type RequestSdaProvisionWiredDeviceV1 struct {
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` // Management Ip Address of the device to be provisioned
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         // Site Name Hierarchy for device location(only building / floor level)
}
type RequestSdaAddTransitPeerNetworkV1 struct {
	TransitPeerNetworkName string                                               `json:"transitPeerNetworkName,omitempty"` // Transit Peer Network Name
	TransitPeerNetworkType string                                               `json:"transitPeerNetworkType,omitempty"` // Transit Peer Network Type
	IPTransitSettings      *RequestSdaAddTransitPeerNetworkV1IPTransitSettings  `json:"ipTransitSettings,omitempty"`      //
	SdaTransitSettings     *RequestSdaAddTransitPeerNetworkV1SdaTransitSettings `json:"sdaTransitSettings,omitempty"`     //
}
type RequestSdaAddTransitPeerNetworkV1IPTransitSettings struct {
	RoutingProtocolName    string `json:"routingProtocolName,omitempty"`    // Routing Protocol Name
	AutonomousSystemNumber string `json:"autonomousSystemNumber,omitempty"` // Autonomous System Number
}
type RequestSdaAddTransitPeerNetworkV1SdaTransitSettings struct {
	TransitControlPlaneSettings *[]RequestSdaAddTransitPeerNetworkV1SdaTransitSettingsTransitControlPlaneSettings `json:"transitControlPlaneSettings,omitempty"` //
}
type RequestSdaAddTransitPeerNetworkV1SdaTransitSettingsTransitControlPlaneSettings struct {
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         // Site Name Hierarchy where device is provisioned
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` // Device Management Ip Address of provisioned device
}
type RequestSdaAddVnInFabricV1 struct {
	VirtualNetworkName string `json:"virtualNetworkName,omitempty"` // Virtual Network Name, that is created at Global level
	SiteNameHierarchy  string `json:"siteNameHierarchy,omitempty"`  // Path of sda Fabric Site
}
type RequestSdaAddIPPoolInSdaVirtualNetworkV1 struct {
	SiteNameHierarchy     string `json:"siteNameHierarchy,omitempty"`     // Path of sda Fabric Site
	VirtualNetworkName    string `json:"virtualNetworkName,omitempty"`    // Virtual Network Name, that is associated to Fabric Site
	IsLayer2Only          *bool  `json:"isLayer2Only,omitempty"`          // Layer2 Only enablement flag and default value is False
	IPPoolName            string `json:"ipPoolName,omitempty"`            // Ip Pool Name, that is reserved to Fabric Site (Required for L3 and INFRA_VN)
	VLANID                string `json:"vlanId,omitempty"`                // vlan Id(applicable for L3 , L2 and  INFRA_VN)
	VLANName              string `json:"vlanName,omitempty"`              // Vlan name represent the segment name, if empty, vlanName would be auto generated by API
	AutoGenerateVLANName  *bool  `json:"autoGenerateVlanName,omitempty"`  // It will auto generate vlanName, if vlanName is empty(applicable for L3  and INFRA_VN)
	TrafficType           string `json:"trafficType,omitempty"`           // Traffic type(applicable for L3  and L2)
	ScalableGroupName     string `json:"scalableGroupName,omitempty"`     // Scalable Group Name(applicable for L3)
	IsL2FloodingEnabled   *bool  `json:"isL2FloodingEnabled,omitempty"`   // Layer2 flooding enablement flag(applicable for L3 , L2 and always true for L2 and default value is False )
	IsThisCriticalPool    *bool  `json:"isThisCriticalPool,omitempty"`    // Critical pool enablement flag(applicable for L3 and default value is False )
	IsWirelessPool        *bool  `json:"isWirelessPool,omitempty"`        // Wireless Pool enablement flag(applicable for L3  and L2 and default value is False )
	IsIPDirectedBroadcast *bool  `json:"isIpDirectedBroadcast,omitempty"` // Ip Directed Broadcast enablement flag(applicable for L3 and default value is False )
	IsCommonPool          *bool  `json:"isCommonPool,omitempty"`          // Common Pool enablement flag(applicable for L3 and L2 and default value is False )
	IsBridgeModeVm        *bool  `json:"isBridgeModeVm,omitempty"`        // Bridge Mode Vm enablement flag (applicable for L3 and L2 and default value is False )
	PoolType              string `json:"poolType,omitempty"`              // Pool Type (applicable for INFRA_VN)
}
type RequestSdaUpdateAnycastGatewaysV1 []RequestItemSdaUpdateAnycastGatewaysV1 // Array of RequestSdaUpdateAnycastGatewaysV1
type RequestItemSdaUpdateAnycastGatewaysV1 struct {
	ID                                      string `json:"id,omitempty"`                                      // ID of the anycast gateway (updating this field is not allowed).
	FabricID                                string `json:"fabricId,omitempty"`                                // ID of the fabric this anycast gateway is assigned to. Updating anycast gateways on fabric zones is not allowed--instead, update the corresponding anycast gateway on the fabric site and the updates will be applied on all applicable fabric zones (updating this field is not allowed).
	VirtualNetworkName                      string `json:"virtualNetworkName,omitempty"`                      // Name of the layer 3 virtual network associated with the anycast gateway (updating this field is not allowed).
	IPPoolName                              string `json:"ipPoolName,omitempty"`                              // Name of the IP pool associated with the anycast gateway (updating this field is not allowed).
	TCPMssAdjustment                        *int   `json:"tcpMssAdjustment,omitempty"`                        // TCP maximum segment size adjustment.
	VLANName                                string `json:"vlanName,omitempty"`                                // Name of the VLAN of the anycast gateway (updating this field is not allowed).
	VLANID                                  *int   `json:"vlanId,omitempty"`                                  // ID of the VLAN of the anycast gateway (updating this field is not allowed).
	TrafficType                             string `json:"trafficType,omitempty"`                             // The type of traffic the anycast gateway serves.
	PoolType                                string `json:"poolType,omitempty"`                                // The pool type of the anycast gateway (required for & applicable only to INFRA_VN; updating this field is not allowed).
	SecurityGroupName                       string `json:"securityGroupName,omitempty"`                       // Name of the associated Security Group (not applicable to INFRA_VN).
	IsCriticalPool                          *bool  `json:"isCriticalPool,omitempty"`                          // Enable/disable critical VLAN (not applicable to INFRA_VN; updating this field is not allowed).
	IsLayer2FloodingEnabled                 *bool  `json:"isLayer2FloodingEnabled,omitempty"`                 // Enable/disable layer 2 flooding (not applicable to INFRA_VN).
	IsWirelessPool                          *bool  `json:"isWirelessPool,omitempty"`                          // Enable/disable fabric-enabled wireless (not applicable to INFRA_VN).
	IsIPDirectedBroadcast                   *bool  `json:"isIpDirectedBroadcast,omitempty"`                   // Enable/disable IP-directed broadcast (not applicable to INFRA_VN).
	IsIntraSubnetRoutingEnabled             *bool  `json:"isIntraSubnetRoutingEnabled,omitempty"`             // Enable/disable Intra-Subnet Routing (not applicable to INFRA_VN; updating this field is not allowed).
	IsMultipleIPToMacAddresses              *bool  `json:"isMultipleIpToMacAddresses,omitempty"`              // Enable/disable multiple IP-to-MAC Addresses (Wireless Bridged-Network Virtual Machine; not applicable to INFRA_VN).
	IsSupplicantBasedExtendedNodeOnboarding *bool  `json:"isSupplicantBasedExtendedNodeOnboarding,omitempty"` // Enable/disable Supplicant-Based Extended Node Onboarding (applicable only to INFRA_VN requests; must not be null when poolType is EXTENDED_NODE).
	IsGroupBasedPolicyEnforcementEnabled    *bool  `json:"isGroupBasedPolicyEnforcementEnabled,omitempty"`    // Enable/disable Group-Based Policy Enforcement (applicable only to INFRA_VN; defaults to false).
}
type RequestSdaAddAnycastGatewaysV1 []RequestItemSdaAddAnycastGatewaysV1 // Array of RequestSdaAddAnycastGatewaysV1
type RequestItemSdaAddAnycastGatewaysV1 struct {
	FabricID                                string `json:"fabricId,omitempty"`                                // ID of the fabric this anycast gateway is to be assigned to.
	VirtualNetworkName                      string `json:"virtualNetworkName,omitempty"`                      // Name of the layer 3 virtual network associated with the anycast gateway. the virtual network must have already been added to the site before creating an anycast gateway with it.
	IPPoolName                              string `json:"ipPoolName,omitempty"`                              // Name of the IP pool associated with the anycast gateway.
	TCPMssAdjustment                        *int   `json:"tcpMssAdjustment,omitempty"`                        // TCP maximum segment size adjustment.
	VLANName                                string `json:"vlanName,omitempty"`                                // Name of the VLAN of the anycast gateway.
	VLANID                                  *int   `json:"vlanId,omitempty"`                                  // ID of the VLAN of the anycast gateway. allowed VLAN range is 2-4093 except for reserved VLANs 1002-1005, 2046, and 4094. if deploying an anycast gateway on a fabric zone, this vlanId must match the vlanId of the corresponding anycast gateway on the fabric site.
	TrafficType                             string `json:"trafficType,omitempty"`                             // The type of traffic the anycast gateway serves.
	PoolType                                string `json:"poolType,omitempty"`                                // The pool type of the anycast gateway (required for & applicable only to INFRA_VN).
	SecurityGroupName                       string `json:"securityGroupName,omitempty"`                       // Name of the associated Security Group (not applicable to INFRA_VN).
	IsCriticalPool                          *bool  `json:"isCriticalPool,omitempty"`                          // Enable/disable critical VLAN. if true, autoGenerateVlanName must also be true. (isCriticalPool is not applicable to INFRA_VN).
	IsLayer2FloodingEnabled                 *bool  `json:"isLayer2FloodingEnabled,omitempty"`                 // Enable/disable layer 2 flooding (not applicable to INFRA_VN).
	IsWirelessPool                          *bool  `json:"isWirelessPool,omitempty"`                          // Enable/disable fabric-enabled wireless (not applicable to INFRA_VN).
	IsIPDirectedBroadcast                   *bool  `json:"isIpDirectedBroadcast,omitempty"`                   // Enable/disable IP-directed broadcast (not applicable to INFRA_VN).
	IsIntraSubnetRoutingEnabled             *bool  `json:"isIntraSubnetRoutingEnabled,omitempty"`             // Enable/disable Intra-Subnet Routing (not applicable to INFRA_VN).
	IsMultipleIPToMacAddresses              *bool  `json:"isMultipleIpToMacAddresses,omitempty"`              // Enable/disable multiple IP-to-MAC Addresses (Wireless Bridged-Network Virtual Machine; not applicable to INFRA_VN).
	IsSupplicantBasedExtendedNodeOnboarding *bool  `json:"isSupplicantBasedExtendedNodeOnboarding,omitempty"` // Enable/disable Supplicant-Based Extended Node Onboarding (applicable only to INFRA_VN).
	IsGroupBasedPolicyEnforcementEnabled    *bool  `json:"isGroupBasedPolicyEnforcementEnabled,omitempty"`    // Enable/disable Group-Based Policy Enforcement (applicable only to INFRA_VN; defaults to false).
	AutoGenerateVLANName                    *bool  `json:"autoGenerateVlanName,omitempty"`                    // This field cannot be true when vlanName is provided. the vlanName will be generated as "{ipPoolGroupV4Cidr}-{virtualNetworkName}" for non-critical VLANs. for critical VLANs with DATA trafficType, vlanName will be "CRITICAL_VLAN". for critical VLANs with VOICE trafficType, vlanName will be "VOICE_VLAN".
}
type RequestSdaUpdateAuthenticationProfileV1 []RequestItemSdaUpdateAuthenticationProfileV1 // Array of RequestSdaUpdateAuthenticationProfileV1
type RequestItemSdaUpdateAuthenticationProfileV1 struct {
	ID                        string `json:"id,omitempty"`                        // ID of the authentication profile (updating this field is not allowed).
	FabricID                  string `json:"fabricId,omitempty"`                  // ID of the fabric this authentication profile is assigned to (updating this field is not allowed).
	AuthenticationProfileName string `json:"authenticationProfileName,omitempty"` // The default host authentication template (updating this field is not allowed).
	AuthenticationOrder       string `json:"authenticationOrder,omitempty"`       // First authentication method.
	Dot1XToMabFallbackTimeout *int   `json:"dot1xToMabFallbackTimeout,omitempty"` // 802.1x Timeout.
	WakeOnLan                 *bool  `json:"wakeOnLan,omitempty"`                 // Wake on LAN.
	NumberOfHosts             string `json:"numberOfHosts,omitempty"`             // Number of Hosts.
	IsBpduGuardEnabled        *bool  `json:"isBpduGuardEnabled,omitempty"`        // Enable/disable BPDU Guard. Only applicable when authenticationProfileName is set to "Closed Authentication" (defaults to true).
}
type RequestSdaUpdateExtranetPolicyV1 []RequestItemSdaUpdateExtranetPolicyV1 // Array of RequestSdaUpdateExtranetPolicyV1
type RequestItemSdaUpdateExtranetPolicyV1 struct {
	ID                            string   `json:"id,omitempty"`                            // ID of the existing extranet policy (updating this field is not allowed).
	ExtranetPolicyName            string   `json:"extranetPolicyName,omitempty"`            // Name of the existing extranet policy (updating this field is not allowed).
	FabricIDs                     []string `json:"fabricIds,omitempty"`                     // IDs of the fabric sites associated with this extranet policy.
	ProviderVirtualNetworkName    string   `json:"providerVirtualNetworkName,omitempty"`    // Name of the existing provider virtual network (updating this field is not allowed).
	SubscriberVirtualNetworkNames []string `json:"subscriberVirtualNetworkNames,omitempty"` // Name of the subscriber virtual networks.
}
type RequestSdaAddExtranetPolicyV1 []RequestItemSdaAddExtranetPolicyV1 // Array of RequestSdaAddExtranetPolicyV1
type RequestItemSdaAddExtranetPolicyV1 struct {
	ExtranetPolicyName            string   `json:"extranetPolicyName,omitempty"`            // Name of the extranet policy to be created.
	FabricIDs                     []string `json:"fabricIds,omitempty"`                     // IDs of the fabric sites to be associated with this extranet policy.
	ProviderVirtualNetworkName    string   `json:"providerVirtualNetworkName,omitempty"`    // Name of the existing provider virtual network.
	SubscriberVirtualNetworkNames []string `json:"subscriberVirtualNetworkNames,omitempty"` // Name of the subscriber virtual networks.
}
type RequestSdaUpdateFabricDevicesV1 []RequestItemSdaUpdateFabricDevicesV1 // Array of RequestSdaUpdateFabricDevicesV1
type RequestItemSdaUpdateFabricDevicesV1 struct {
	ID                   string                                                   `json:"id,omitempty"`                   // ID of the fabric device. (updating this field is not allowed).
	NetworkDeviceID      string                                                   `json:"networkDeviceId,omitempty"`      // Network device ID of the fabric device. (updating this field is not allowed).
	FabricID             string                                                   `json:"fabricId,omitempty"`             // ID of the fabric of this fabric device. (updating this field is not allowed).
	DeviceRoles          []string                                                 `json:"deviceRoles,omitempty"`          // List of the roles of the fabric device. Allowed values are [CONTROL_PLANE_NODE, EDGE_NODE, BORDER_NODE, WIRELESS_CONTROLLER_NODE]. (updating this field is not allowed).
	BorderDeviceSettings *RequestItemSdaUpdateFabricDevicesV1BorderDeviceSettings `json:"borderDeviceSettings,omitempty"` //
}
type RequestItemSdaUpdateFabricDevicesV1BorderDeviceSettings struct {
	BorderTypes    []string                                                               `json:"borderTypes,omitempty"`    // List of the border types of the fabric device. Allowed values are [LAYER_2, LAYER_3].
	Layer3Settings *RequestItemSdaUpdateFabricDevicesV1BorderDeviceSettingsLayer3Settings `json:"layer3Settings,omitempty"` //
}
type RequestItemSdaUpdateFabricDevicesV1BorderDeviceSettingsLayer3Settings struct {
	LocalAutonomousSystemNumber  string `json:"localAutonomousSystemNumber,omitempty"`  // BGP Local autonomous system number of the fabric border device. Allowed range is [1 to 4294967295]. (updating this field is not allowed).
	IsDefaultExit                *bool  `json:"isDefaultExit,omitempty"`                // Set this to make the fabric border device the gateway of last resort for this site. Any unknown traffic will be sent to this fabric border device from edge nodes. (updating this field is not allowed).
	ImportExternalRoutes         *bool  `json:"importExternalRoutes,omitempty"`         // Set this to import external routes from other routing protocols (such as BGP) to the fabric control plane. (updating this field is not allowed).
	BorderPriority               *int   `json:"borderPriority,omitempty"`               // Border priority of the fabric border device. Allowed range is [1-9]. A lower value indicates higher priority. E.g., a priority of 1 takes precedence over 5. Default priority would be set to 10.
	PrependAutonomousSystemCount *int   `json:"prependAutonomousSystemCount,omitempty"` // Prepend autonomous system count of the fabric border device. Allowed range is [1 to 10].
}
type RequestSdaAddFabricDevicesV1 []RequestItemSdaAddFabricDevicesV1 // Array of RequestSdaAddFabricDevicesV1
type RequestItemSdaAddFabricDevicesV1 struct {
	NetworkDeviceID      string                                                `json:"networkDeviceId,omitempty"`      // Network device ID of the fabric device.
	FabricID             string                                                `json:"fabricId,omitempty"`             // ID of the fabric of this fabric device.
	DeviceRoles          []string                                              `json:"deviceRoles,omitempty"`          // List of the roles of the fabric device. Allowed values are [CONTROL_PLANE_NODE, EDGE_NODE, BORDER_NODE, WIRELESS_CONTROLLER_NODE].
	BorderDeviceSettings *RequestItemSdaAddFabricDevicesV1BorderDeviceSettings `json:"borderDeviceSettings,omitempty"` //
}
type RequestItemSdaAddFabricDevicesV1BorderDeviceSettings struct {
	BorderTypes    []string                                                            `json:"borderTypes,omitempty"`    // List of the border types of the fabric device. Allowed values are [LAYER_2, LAYER_3].
	Layer3Settings *RequestItemSdaAddFabricDevicesV1BorderDeviceSettingsLayer3Settings `json:"layer3Settings,omitempty"` //
}
type RequestItemSdaAddFabricDevicesV1BorderDeviceSettingsLayer3Settings struct {
	LocalAutonomousSystemNumber  string `json:"localAutonomousSystemNumber,omitempty"`  // BGP Local autonomous system number of the fabric border device. Allowed range is [1 to 4294967295].
	IsDefaultExit                *bool  `json:"isDefaultExit,omitempty"`                // Set this to make the fabric border device the gateway of last resort for this site. Any unknown traffic will be sent to this fabric border device from edge nodes.
	ImportExternalRoutes         *bool  `json:"importExternalRoutes,omitempty"`         // Set this to import external routes from other routing protocols (such as BGP) to the fabric control plane.
	BorderPriority               *int   `json:"borderPriority,omitempty"`               // Border priority of the fabric border device. Allowed range is [1-9]. A lower value indicates higher priority. E.g., a priority of 1 takes precedence over 5. Default priority would be set to 10.
	PrependAutonomousSystemCount *int   `json:"prependAutonomousSystemCount,omitempty"` // Prepend autonomous system count of the fabric border device. Allowed range is [1 to 10].
}
type RequestSdaAddFabricDevicesLayer2HandoffsV1 []RequestItemSdaAddFabricDevicesLayer2HandoffsV1 // Array of RequestSdaAddFabricDevicesLayer2HandoffsV1
type RequestItemSdaAddFabricDevicesLayer2HandoffsV1 struct {
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Network device ID of the fabric device.
	FabricID        string `json:"fabricId,omitempty"`        // ID of the fabric this device is assigned to.
	InterfaceName   string `json:"interfaceName,omitempty"`   // Interface name of the layer 2 handoff. E.g., GigabitEthernet1/0/4
	InternalVLANID  *int   `json:"internalVlanId,omitempty"`  // VLAN number associated with this fabric. Allowed VLAN range is 2-4094 except for reserved vlans (1, 1002-1005, 2046, 4094).
	ExternalVLANID  *int   `json:"externalVlanId,omitempty"`  // External VLAN number into which the fabric must be extended. Allowed VLAN range is 2-4094 except for reserved vlans (1, 1002-1005, 2046, 4094).
}
type RequestSdaAddFabricDevicesLayer3HandoffsWithIPTransitV1 []RequestItemSdaAddFabricDevicesLayer3HandoffsWithIPTransitV1 // Array of RequestSdaAddFabricDevicesLayer3HandoffsWithIpTransitV1
type RequestItemSdaAddFabricDevicesLayer3HandoffsWithIPTransitV1 struct {
	NetworkDeviceID                string `json:"networkDeviceId,omitempty"`                // Network device ID of the fabric device.
	FabricID                       string `json:"fabricId,omitempty"`                       // ID of the fabric this device is assigned to.
	TransitNetworkID               string `json:"transitNetworkId,omitempty"`               // ID of the transit network of the layer 3 handoff ip transit.
	InterfaceName                  string `json:"interfaceName,omitempty"`                  // Interface name of the layer 3 handoff ip transit.
	ExternalConnectivityIPPoolName string `json:"externalConnectivityIpPoolName,omitempty"` // External connectivity ip pool will be used by Catalyst Center to allocate IP address for the connection between the border node and peer.
	VirtualNetworkName             string `json:"virtualNetworkName,omitempty"`             // Name of the virtual network associated with this fabric site.
	VLANID                         *int   `json:"vlanId,omitempty"`                         // VLAN number for the Switch Virtual Interface (SVI) used to establish BGP peering with the external domain for the virtual network.  Allowed VLAN range is 2-4094 except for reserved vlans (1, 1002-1005, 2046, 4094).
	TCPMssAdjustment               *int   `json:"tcpMssAdjustment,omitempty"`               // TCP maximum segment size (mss) value for the layer 3 handoff. Allowed range is [500-1440]. TCP MSS Adjustment value is applicable for the TCP sessions over both IPv4 and IPv6.
	LocalIPAddress                 string `json:"localIpAddress,omitempty"`                 // Local ipv4 address for the selected virtual network. Enter the IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). Not applicable if you have already provided an external connectivity ip pool name.
	RemoteIPAddress                string `json:"remoteIpAddress,omitempty"`                // Remote ipv4 address for the selected virtual network. Enter the IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). Not applicable if you have already provided an external connectivity ip pool name.
	LocalIPv6Address               string `json:"localIpv6Address,omitempty"`               // Local ipv6 address for the selected virtual network. Enter the IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). Not applicable if you have already provided an external connectivity ip pool name.
	RemoteIPv6Address              string `json:"remoteIpv6Address,omitempty"`              // Remote ipv6 address for the selected virtual network. Enter the IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). Not applicable if you have already provided an external connectivity ip pool name.
}
type RequestSdaUpdateFabricDevicesLayer3HandoffsWithIPTransitV1 []RequestItemSdaUpdateFabricDevicesLayer3HandoffsWithIPTransitV1 // Array of RequestSdaUpdateFabricDevicesLayer3HandoffsWithIpTransitV1
type RequestItemSdaUpdateFabricDevicesLayer3HandoffsWithIPTransitV1 struct {
	ID                             string `json:"id,omitempty"`                             // ID of the fabric device layer 3 handoff ip transit. (updating this field is not allowed).
	NetworkDeviceID                string `json:"networkDeviceId,omitempty"`                // Network device ID of the fabric device. (updating this field is not allowed).
	FabricID                       string `json:"fabricId,omitempty"`                       // ID of the fabric this device is assigned to. (updating this field is not allowed).
	TransitNetworkID               string `json:"transitNetworkId,omitempty"`               // ID of the transit network of the layer 3 handoff ip transit. (updating this field is not allowed).
	InterfaceName                  string `json:"interfaceName,omitempty"`                  // Interface name of the layer 3 handoff ip transit. (updating this field is not allowed).
	ExternalConnectivityIPPoolName string `json:"externalConnectivityIpPoolName,omitempty"` // External connectivity ip pool will be used by Catalyst Center to allocate IP address for the connection between the border node and peer. (updating this field is not allowed).
	VirtualNetworkName             string `json:"virtualNetworkName,omitempty"`             // Name of the virtual network associated with this fabric site. (updating this field is not allowed).
	VLANID                         *int   `json:"vlanId,omitempty"`                         // VLAN number for the Switch Virtual Interface (SVI) used to establish BGP peering with the external domain for the virtual network. Allowed VLAN range is 2-4094 except for reserved vlans (1, 1002-1005, 2046, 4094). (updating this field is not allowed).
	TCPMssAdjustment               *int   `json:"tcpMssAdjustment,omitempty"`               // TCP maximum segment size (mss) value for the layer 3 handoff. Allowed range is [500-1440]. TCP MSS Adjustment value is applicable for the TCP sessions over both IPv4 and IPv6.
	LocalIPAddress                 string `json:"localIpAddress,omitempty"`                 // Local ipv4 address for the selected virtual network. Enter the IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). (updating this field is not allowed).
	RemoteIPAddress                string `json:"remoteIpAddress,omitempty"`                // Remote ipv4 address for the selected virtual network. Enter the IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). (updating this field is not allowed).
	LocalIPv6Address               string `json:"localIpv6Address,omitempty"`               // Local ipv6 address for the selected virtual network. Enter the IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). If the value has already been set, it cannot be updated.
	RemoteIPv6Address              string `json:"remoteIpv6Address,omitempty"`              // Remote ipv6 address for the selected virtual network. Enter the IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). If the value has already been set, it cannot be updated.
}
type RequestSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransitV1 []RequestItemSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransitV1 // Array of RequestSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransitV1
type RequestItemSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransitV1 struct {
	NetworkDeviceID               string `json:"networkDeviceId,omitempty"`               // Network device ID of the fabric device. (updating this field is not allowed).
	FabricID                      string `json:"fabricId,omitempty"`                      // ID of the fabric this device is assigned to. (updating this field is not allowed).
	TransitNetworkID              string `json:"transitNetworkId,omitempty"`              // ID of the transit network of the layer 3 handoff sda transit. (updating this field is not allowed).
	AffinityIDPrime               *int   `json:"affinityIdPrime,omitempty"`               // Affinity id prime value of the border node. It supersedes the border priority to determine border node preference. Allowed range is [0-2147483647]. The lower the relative value of affinity id prime, the higher the preference for a destination border node.
	AffinityIDDecider             *int   `json:"affinityIdDecider,omitempty"`             // Affinity id decider value of the border node. When the affinity id prime value is the same on multiple devices, the affinity id decider value is used as a tiebreaker. Allowed range is [0-2147483647]. The lower the relative value of affinity id decider, the higher the preference for a destination border node.
	ConnectedToInternet           *bool  `json:"connectedToInternet,omitempty"`           // Set this true to allow associated site to provide internet access to other sites through sd-access.
	IsMulticastOverTransitEnabled *bool  `json:"isMulticastOverTransitEnabled,omitempty"` // Set this true to configure native multicast over multiple sites that are connected to an sd-access transit.
}
type RequestSdaAddFabricDevicesLayer3HandoffsWithSdaTransitV1 []RequestItemSdaAddFabricDevicesLayer3HandoffsWithSdaTransitV1 // Array of RequestSdaAddFabricDevicesLayer3HandoffsWithSdaTransitV1
type RequestItemSdaAddFabricDevicesLayer3HandoffsWithSdaTransitV1 struct {
	NetworkDeviceID               string `json:"networkDeviceId,omitempty"`               // Network device ID of the fabric device.
	FabricID                      string `json:"fabricId,omitempty"`                      // ID of the fabric this device is assigned to.
	TransitNetworkID              string `json:"transitNetworkId,omitempty"`              // ID of the transit network of the layer 3 handoff sda transit.
	AffinityIDPrime               *int   `json:"affinityIdPrime,omitempty"`               // Affinity id prime value of the border node. It supersedes the border priority to determine border node preference. Allowed range is [0-2147483647]. The lower the relative value of affinity id prime, the higher the preference for a destination border node.
	AffinityIDDecider             *int   `json:"affinityIdDecider,omitempty"`             // Affinity id decider value of the border node. When the affinity id prime value is the same on multiple devices, the affinity id decider value is used as a tiebreaker. Allowed range is [0-2147483647]. The lower the relative value of affinity id decider, the higher the preference for a destination border node.
	ConnectedToInternet           *bool  `json:"connectedToInternet,omitempty"`           // Set this true to allow associated site to provide internet access to other sites through sd-access.
	IsMulticastOverTransitEnabled *bool  `json:"isMulticastOverTransitEnabled,omitempty"` // Set this true to configure native multicast over multiple sites that are connected to an sd-access transit.
}
type RequestSdaAddFabricSiteV1 []RequestItemSdaAddFabricSiteV1 // Array of RequestSdaAddFabricSiteV1
type RequestItemSdaAddFabricSiteV1 struct {
	SiteID                    string `json:"siteId,omitempty"`                    // ID of the network hierarchy.
	AuthenticationProfileName string `json:"authenticationProfileName,omitempty"` // Authentication profile used for this fabric.
	IsPubSubEnabled           *bool  `json:"isPubSubEnabled,omitempty"`           // Specifies whether this fabric site will use pub/sub for control nodes.
}
type RequestSdaUpdateFabricSiteV1 []RequestItemSdaUpdateFabricSiteV1 // Array of RequestSdaUpdateFabricSiteV1
type RequestItemSdaUpdateFabricSiteV1 struct {
	ID                        string `json:"id,omitempty"`                        // ID of the fabric site (updating this field is not allowed).
	SiteID                    string `json:"siteId,omitempty"`                    // ID of the network hierarchy (updating this field is not allowed).
	AuthenticationProfileName string `json:"authenticationProfileName,omitempty"` // Authentication profile used for this fabric.
	IsPubSubEnabled           *bool  `json:"isPubSubEnabled,omitempty"`           // Specifies whether this fabric site will use pub/sub for control nodes.
}
type RequestSdaUpdateFabricZoneV1 []RequestItemSdaUpdateFabricZoneV1 // Array of RequestSdaUpdateFabricZoneV1
type RequestItemSdaUpdateFabricZoneV1 struct {
	ID                        string `json:"id,omitempty"`                        // ID of the fabric zone (updating this field is not allowed).
	SiteID                    string `json:"siteId,omitempty"`                    // ID of the network hierarchy (updating this field is not allowed).
	AuthenticationProfileName string `json:"authenticationProfileName,omitempty"` // Authentication profile used for this fabric.
}
type RequestSdaAddFabricZoneV1 []RequestItemSdaAddFabricZoneV1 // Array of RequestSdaAddFabricZoneV1
type RequestItemSdaAddFabricZoneV1 struct {
	SiteID                    string `json:"siteId,omitempty"`                    // ID of the network hierarchy.
	AuthenticationProfileName string `json:"authenticationProfileName,omitempty"` // Authentication profile used for this fabric.
}
type RequestSdaAddLayer2VirtualNetworksV1 []RequestItemSdaAddLayer2VirtualNetworksV1 // Array of RequestSdaAddLayer2VirtualNetworksV1
type RequestItemSdaAddLayer2VirtualNetworksV1 struct {
	FabricID                           string `json:"fabricId,omitempty"`                           // ID of the fabric this layer 2 virtual network is to be assigned to.
	VLANName                           string `json:"vlanName,omitempty"`                           // Name of the VLAN of the layer 2 virtual network. Must contain only alphanumeric characters, underscores, and hyphens.
	VLANID                             *int   `json:"vlanId,omitempty"`                             // ID of the VLAN of the layer 2 virtual network. Allowed VLAN range is 2-4093 except for reserved VLANs 1002-1005, and 2046. If deploying on a fabric zone, this vlanId must match the vlanId of the corresponding layer 2 virtual network on the fabric site.
	TrafficType                        string `json:"trafficType,omitempty"`                        // The type of traffic that is served.
	IsFabricEnabledWireless            *bool  `json:"isFabricEnabledWireless,omitempty"`            // Set to true to enable wireless. Default is false.
	AssociatedLayer3VirtualNetworkName string `json:"associatedLayer3VirtualNetworkName,omitempty"` // Name of the layer 3 virtual network associated with the layer 2 virtual network. This field is provided to support requests related to virtual network anchoring. The layer 3 virtual network must have already been added to the fabric before association. This field must either be present in all payload elements or none.
}
type RequestSdaUpdateLayer2VirtualNetworksV1 []RequestItemSdaUpdateLayer2VirtualNetworksV1 // Array of RequestSdaUpdateLayer2VirtualNetworksV1
type RequestItemSdaUpdateLayer2VirtualNetworksV1 struct {
	ID                                 string `json:"id,omitempty"`                                 // ID of the layer 2 virtual network (updating this field is not allowed).
	FabricID                           string `json:"fabricId,omitempty"`                           // ID of the fabric this layer 2 virtual network is assigned to (updating this field is not allowed).
	VLANName                           string `json:"vlanName,omitempty"`                           // Name of the VLAN of the layer 2 virtual network. Must contain only alphanumeric characters, underscores, and hyphens (updating this field is not allowed).
	VLANID                             *int   `json:"vlanId,omitempty"`                             // ID of the VLAN of the layer 2 virtual network (updating this field is not allowed).
	TrafficType                        string `json:"trafficType,omitempty"`                        // The type of traffic that is served.
	IsFabricEnabledWireless            *bool  `json:"isFabricEnabledWireless,omitempty"`            // Set to true to enable wireless.
	AssociatedLayer3VirtualNetworkName string `json:"associatedLayer3VirtualNetworkName,omitempty"` // Name of the layer 3 virtual network associated with the layer 2 virtual network. This field is provided to support requests related to virtual network anchoring. This field must either be present in all payload elements or none (updating this field is not allowed).
}
type RequestSdaAddLayer3VirtualNetworksV1 []RequestItemSdaAddLayer3VirtualNetworksV1 // Array of RequestSdaAddLayer3VirtualNetworksV1
type RequestItemSdaAddLayer3VirtualNetworksV1 struct {
	VirtualNetworkName string   `json:"virtualNetworkName,omitempty"` // Name of the layer 3 virtual network.
	FabricIDs          []string `json:"fabricIds,omitempty"`          // IDs of the fabrics this layer 3 virtual network is to be assigned to.
	AnchoredSiteID     string   `json:"anchoredSiteId,omitempty"`     // Fabric ID of the fabric site this layer 3 virtual network is to be anchored at.
}
type RequestSdaUpdateLayer3VirtualNetworksV1 []RequestItemSdaUpdateLayer3VirtualNetworksV1 // Array of RequestSdaUpdateLayer3VirtualNetworksV1
type RequestItemSdaUpdateLayer3VirtualNetworksV1 struct {
	ID                 string   `json:"id,omitempty"`                 // ID of the layer 3 virtual network (updating this field is not allowed).
	VirtualNetworkName string   `json:"virtualNetworkName,omitempty"` // Name of the layer 3 virtual network (updating this field is not allowed).
	FabricIDs          []string `json:"fabricIds,omitempty"`          // IDs of the fabrics this layer 3 virtual network is assigned to.
	AnchoredSiteID     string   `json:"anchoredSiteId,omitempty"`     // Fabric ID of the fabric site this layer 3 virtual network is anchored at.
}
type RequestSdaUpdateMulticastV1 []RequestItemSdaUpdateMulticastV1 // Array of RequestSdaUpdateMulticastV1
type RequestItemSdaUpdateMulticastV1 struct {
	FabricID        string `json:"fabricId,omitempty"`        // ID of the fabric site (updating this field is not allowed).
	ReplicationMode string `json:"replicationMode,omitempty"` // Replication Mode deployed in the fabric site.
}
type RequestSdaAddMulticastVirtualNetworksV1 []RequestItemSdaAddMulticastVirtualNetworksV1 // Array of RequestSdaAddMulticastVirtualNetworksV1
type RequestItemSdaAddMulticastVirtualNetworksV1 struct {
	FabricID           string                                                     `json:"fabricId,omitempty"`           // ID of the fabric site this multicast configuration is associated with.
	VirtualNetworkName string                                                     `json:"virtualNetworkName,omitempty"` // Name of the virtual network associated with the fabric site.
	IPPoolName         string                                                     `json:"ipPoolName,omitempty"`         // Name of the IP Pool associated with the fabric site.
	IPv4SsmRanges      []string                                                   `json:"ipv4SsmRanges,omitempty"`      // IPv4 Source Specific Multicast (SSM) ranges. Allowed ranges are from 225.0.0.0/8 to 239.0.0.0/8. SSM ranges should not conflict with ranges provided for ASM multicast.
	MulticastRPs       *[]RequestItemSdaAddMulticastVirtualNetworksV1MulticastRPs `json:"multicastRPs,omitempty"`       //
}
type RequestItemSdaAddMulticastVirtualNetworksV1MulticastRPs struct {
	RpDeviceLocation string   `json:"rpDeviceLocation,omitempty"` // Device location of the RP.
	IPv4Address      string   `json:"ipv4Address,omitempty"`      // IPv4 address of the RP. For external RP configuration, exactly one of ipv4Address or ipv6Address must be provided. For fabric RP, this address is allocated by SDA and should not be provided during RP creation request and SDA allocated address should be retained in subsequent requests.
	IPv6Address      string   `json:"ipv6Address,omitempty"`      // IPv6 address of the RP. For external RP configuration, exactly one of ipv4Address or ipv6Address must be provided. For fabric RP, this address is allocated by SDA and should not be provided during RP creation request and SDA allocated address should be retained in subsequent requests. ipv6Address can only be provided for virtual networks with dual stack (IPv4 + IPv6) multicast pool.
	IsDefaultV4RP    *bool    `json:"isDefaultV4RP,omitempty"`    // Specifies whether it is a default IPv4 RP.
	IsDefaultV6RP    *bool    `json:"isDefaultV6RP,omitempty"`    // Specifies whether it is a default IPv6 RP.
	NetworkDeviceIDs []string `json:"networkDeviceIds,omitempty"` // IDs of the network devices. This is a required field for fabric RPs. There can be maximum of two fabric RPs for a fabric site and these are shared across all multicast virtual networks. For configuring two fabric RPs in a fabric site both devices must have border roles. Only one RP can be configured in scenarios where a fabric edge device is used as RP or a dual stack multicast pool is used.
	IPv4AsmRanges    []string `json:"ipv4AsmRanges,omitempty"`    // IPv4 Any Source Multicast ranges. Comma seperated list of IPv4 multicast group ranges that will be served by a given Multicast RP. Only IPv4 ranges can be provided. For fabric RP, both IPv4 and IPv6 ranges can be provided together. For external RP, IPv4 ranges should be provided for IPv4 external RP and IPv6 ranges should be provided for IPv6 external RP.
	IPv6AsmRanges    []string `json:"ipv6AsmRanges,omitempty"`    // IPv6 Any Source Multicast ranges. Comma seperated list of IPv6 multicast group ranges that will be served by a given Multicast RP. Only IPv6 ranges can be provided. IPv6 ranges can only be provided for dual stack multicast pool. For fabric RP, both IPv4 and IPv6 ranges can be provided together. For external RP, IPv4 ranges should be provided for IPv4 external RP and IPv6 ranges should be provided for IPv6 external RP.
}
type RequestSdaUpdateMulticastVirtualNetworksV1 []RequestItemSdaUpdateMulticastVirtualNetworksV1 // Array of RequestSdaUpdateMulticastVirtualNetworksV1
type RequestItemSdaUpdateMulticastVirtualNetworksV1 struct {
	ID                 string                                                        `json:"id,omitempty"`                 // ID of the multicast configuration (updating this field is not allowed).
	FabricID           string                                                        `json:"fabricId,omitempty"`           // ID of the fabric site this multicast configuration is associated with (updating this field is not allowed).
	VirtualNetworkName string                                                        `json:"virtualNetworkName,omitempty"` // Name of the virtual network associated with the fabric site (updating this field is not allowed).
	IPPoolName         string                                                        `json:"ipPoolName,omitempty"`         // Name of the IP Pool associated with the fabric site (updating this field is not allowed).
	IPv4SsmRanges      []string                                                      `json:"ipv4SsmRanges,omitempty"`      // IPv4 Source Specific Multicast (SSM) ranges. Allowed ranges are from 225.0.0.0/8 to 239.0.0.0/8. SSM ranges should not conflict with ranges provided for ASM multicast.
	MulticastRPs       *[]RequestItemSdaUpdateMulticastVirtualNetworksV1MulticastRPs `json:"multicastRPs,omitempty"`       //
}
type RequestItemSdaUpdateMulticastVirtualNetworksV1MulticastRPs struct {
	RpDeviceLocation string   `json:"rpDeviceLocation,omitempty"` // Device location of the RP.
	IPv4Address      string   `json:"ipv4Address,omitempty"`      // IPv4 address of the RP. For external RP configuration, exactly one of ipv4Address or ipv6Address must be provided. For fabric RP, this address is allocated by SDA and should not be provided during RP creation request and SDA allocated address should be retained in subsequent requests.
	IPv6Address      string   `json:"ipv6Address,omitempty"`      // IPv6 address of the RP. For external RP configuration, exactly one of ipv4Address or ipv6Address must be provided. For fabric RP, this address is allocated by SDA and should not be provided during RP creation request and SDA allocated address should be retained in subsequent requests. ipv6Address can only be provided for virtual networks with dual stack (IPv4 + IPv6) multicast pool.
	IsDefaultV4RP    *bool    `json:"isDefaultV4RP,omitempty"`    // Specifies whether it is a default IPv4 RP.
	IsDefaultV6RP    *bool    `json:"isDefaultV6RP,omitempty"`    // Specifies whether it is a default IPv6 RP.
	NetworkDeviceIDs []string `json:"networkDeviceIds,omitempty"` // IDs of the network devices. This is a required field for fabric RPs. There can be maximum of two fabric RPs for a fabric site and these are shared across all multicast virtual networks. For configuring two fabric RPs in a fabric site both devices must have border roles. Only one RP can be configured in scenarios where a fabric edge device is used as RP or a dual stack multicast pool is used.
	IPv4AsmRanges    []string `json:"ipv4AsmRanges,omitempty"`    // IPv4 Any Source Multicast ranges. Comma seperated list of IPv4 multicast group ranges that will be served by a given Multicast RP. Only IPv4 ranges can be provided. For fabric RP, both IPv4 and IPv6 ranges can be provided together. For external RP, IPv4 ranges should be provided for IPv4 external RP and IPv6 ranges should be provided for IPv6 external RP.
	IPv6AsmRanges    []string `json:"ipv6AsmRanges,omitempty"`    // IPv6 Any Source Multicast ranges. Comma seperated list of IPv6 multicast group ranges that will be served by a given Multicast RP. Only IPv6 ranges can be provided. IPv6 ranges can only be provided for dual stack multicast pool. For fabric RP, both IPv4 and IPv6 ranges can be provided together. For external RP, IPv4 ranges should be provided for IPv4 external RP and IPv6 ranges should be provided for IPv6 external RP.
}
type RequestSdaAddPortAssignmentsV1 []RequestItemSdaAddPortAssignmentsV1 // Array of RequestSdaAddPortAssignmentsV1
type RequestItemSdaAddPortAssignmentsV1 struct {
	FabricID                 string `json:"fabricId,omitempty"`                 // ID of the fabric the device is assigned to.
	NetworkDeviceID          string `json:"networkDeviceId,omitempty"`          // Network device ID of the port assignment.
	InterfaceName            string `json:"interfaceName,omitempty"`            // Interface name of the port assignment.
	ConnectedDeviceType      string `json:"connectedDeviceType,omitempty"`      // Connected device type of the port assignment.
	DataVLANName             string `json:"dataVlanName,omitempty"`             // Data VLAN name of the port assignment.
	VoiceVLANName            string `json:"voiceVlanName,omitempty"`            // Voice VLAN name of the port assignment.
	AuthenticateTemplateName string `json:"authenticateTemplateName,omitempty"` // Authenticate template name of the port assignment.
	SecurityGroupName        string `json:"securityGroupName,omitempty"`        // Security group name of the port assignment.
	InterfaceDescription     string `json:"interfaceDescription,omitempty"`     // Interface description of the port assignment.
}
type RequestSdaUpdatePortAssignmentsV1 []RequestItemSdaUpdatePortAssignmentsV1 // Array of RequestSdaUpdatePortAssignmentsV1
type RequestItemSdaUpdatePortAssignmentsV1 struct {
	ID                       string `json:"id,omitempty"`                       // ID of the port assignment.
	FabricID                 string `json:"fabricId,omitempty"`                 // ID of the fabric the device is assigned to (updating this filed is not allowed).
	NetworkDeviceID          string `json:"networkDeviceId,omitempty"`          // Network device ID of the port assignment (updating this field is not allowed).
	InterfaceName            string `json:"interfaceName,omitempty"`            // Interface name of the port assignment (updating this field is not allowed).
	ConnectedDeviceType      string `json:"connectedDeviceType,omitempty"`      // Connected device type of the port assignment (updating this field is not allowed).
	DataVLANName             string `json:"dataVlanName,omitempty"`             // Data VLAN name of the port assignment.
	VoiceVLANName            string `json:"voiceVlanName,omitempty"`            // Voice VLAN name of the port assignment.
	AuthenticateTemplateName string `json:"authenticateTemplateName,omitempty"` // Authenticate template name of the port assignment.
	ScalableGroupName        string `json:"scalableGroupName,omitempty"`        // Scalable group name of the port assignment.
	InterfaceDescription     string `json:"interfaceDescription,omitempty"`     // Interface description of the port assignment.
}
type RequestSdaAddPortChannelsV1 []RequestItemSdaAddPortChannelsV1 // Array of RequestSdaAddPortChannelsV1
type RequestItemSdaAddPortChannelsV1 struct {
	FabricID            string   `json:"fabricId,omitempty"`            // ID of the fabric the device is assigned to.
	NetworkDeviceID     string   `json:"networkDeviceId,omitempty"`     // ID of the network device.
	InterfaceNames      []string `json:"interfaceNames,omitempty"`      // Interface names for this port channel (Maximum 16 ports for LACP protocol, Maximum 8 ports for PAGP and ON protocol).
	ConnectedDeviceType string   `json:"connectedDeviceType,omitempty"` // Connected device type of the port channel.
	Protocol            string   `json:"protocol,omitempty"`            // Protocol of the port channel (only PAGP is allowed if connectedDeviceType is EXTENDED_NODE).
	Description         string   `json:"description,omitempty"`         // Description of the port channel.
}
type RequestSdaUpdatePortChannelsV1 []RequestItemSdaUpdatePortChannelsV1 // Array of RequestSdaUpdatePortChannelsV1
type RequestItemSdaUpdatePortChannelsV1 struct {
	ID                  string   `json:"id,omitempty"`                  // ID of the port channel (updating this field is not allowed).
	FabricID            string   `json:"fabricId,omitempty"`            // ID of the fabric the device is assigned to (updating this field is not allowed).
	NetworkDeviceID     string   `json:"networkDeviceId,omitempty"`     // ID of the network device (updating this field is not allowed).
	PortChannelName     string   `json:"portChannelName,omitempty"`     // Name of the port channel (updating this field is not allowed).
	InterfaceNames      []string `json:"interfaceNames,omitempty"`      // Interface names for this port channel (Maximum 16 ports for LACP protocol, Maximum 8 ports for PAGP and ON protocol).
	ConnectedDeviceType string   `json:"connectedDeviceType,omitempty"` // Connected device type of the port channel.
	Protocol            string   `json:"protocol,omitempty"`            // Protocol of the port channel (updating this field is not allowed).
	Description         string   `json:"description,omitempty"`         // Description of the port channel.
}
type RequestSdaProvisionDevicesV1 []RequestItemSdaProvisionDevicesV1 // Array of RequestSdaProvisionDevicesV1
type RequestItemSdaProvisionDevicesV1 struct {
	SiteID          string `json:"siteId,omitempty"`          // ID of the site this network device needs to be provisioned.
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // ID of network device to be provisioned.
}
type RequestSdaReProvisionDevicesV1 []RequestItemSdaReProvisionDevicesV1 // Array of RequestSdaReProvisionDevicesV1
type RequestItemSdaReProvisionDevicesV1 struct {
	ID              string `json:"id,omitempty"`              // ID of the provisioned device.
	SiteID          string `json:"siteId,omitempty"`          // ID of the site this device is already provisioned to. (updating this field is not allowed).
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // ID of the network device to be re-provisioned. (updating this field is not allowed).
}
type RequestSdaUpdateTransitNetworksV1 []RequestItemSdaUpdateTransitNetworksV1 // Array of RequestSdaUpdateTransitNetworksV1
type RequestItemSdaUpdateTransitNetworksV1 struct {
	ID                 string                                                   `json:"id,omitempty"`                 // ID of the transit network (updating this field is not allowed).
	Name               string                                                   `json:"name,omitempty"`               // Name of the transit network (updating this field is not allowed).
	Type               string                                                   `json:"type,omitempty"`               // Type of the transit network (updating this field is not allowed).
	IPTransitSettings  *RequestItemSdaUpdateTransitNetworksV1IPTransitSettings  `json:"ipTransitSettings,omitempty"`  //
	SdaTransitSettings *RequestItemSdaUpdateTransitNetworksV1SdaTransitSettings `json:"sdaTransitSettings,omitempty"` //
}
type RequestItemSdaUpdateTransitNetworksV1IPTransitSettings struct {
	RoutingProtocolName    string `json:"routingProtocolName,omitempty"`    // Routing Protocol Name of the IP transit network (updating this field is not allowed).
	AutonomousSystemNumber string `json:"autonomousSystemNumber,omitempty"` // Autonomous System Number of the IP transit network. Allowed range is [1 to 4294967295] (updating this field is not allowed).
}
type RequestItemSdaUpdateTransitNetworksV1SdaTransitSettings struct {
	IsMulticastOverTransitEnabled *bool    `json:"isMulticastOverTransitEnabled,omitempty"` // Set this to true to enable multicast over SD-Access transit. This supports Native Multicast over SD-Access Transit. This is only applicable for transit of type SDA_LISP_PUB_SUB_TRANSIT.
	ControlPlaneNetworkDeviceIDs  []string `json:"controlPlaneNetworkDeviceIds,omitempty"`  // List of network device IDs that will be used as control plane nodes. Maximum 2 network device IDs can be provided for transit of type SDA_LISP_BGP_TRANSIT and maximum 4 network device IDs can be provided for transit of type SDA_LISP_PUB_SUB_TRANSIT.
}
type RequestSdaAddTransitNetworksV1 []RequestItemSdaAddTransitNetworksV1 // Array of RequestSdaAddTransitNetworksV1
type RequestItemSdaAddTransitNetworksV1 struct {
	Name               string                                                `json:"name,omitempty"`               // Name of the transit network.
	Type               string                                                `json:"type,omitempty"`               // Type of the transit network.
	IPTransitSettings  *RequestItemSdaAddTransitNetworksV1IPTransitSettings  `json:"ipTransitSettings,omitempty"`  //
	SdaTransitSettings *RequestItemSdaAddTransitNetworksV1SdaTransitSettings `json:"sdaTransitSettings,omitempty"` //
}
type RequestItemSdaAddTransitNetworksV1IPTransitSettings struct {
	RoutingProtocolName    string `json:"routingProtocolName,omitempty"`    // Routing protocol name of the IP transit network.
	AutonomousSystemNumber string `json:"autonomousSystemNumber,omitempty"` // Autonomous system number of the IP transit network. Allowed range is [1 to 4294967295].
}
type RequestItemSdaAddTransitNetworksV1SdaTransitSettings struct {
	IsMulticastOverTransitEnabled *bool    `json:"isMulticastOverTransitEnabled,omitempty"` // Set this to true to enable multicast over SD-Access transit.  This supports Native Multicast over SD-Access Transit. This is only applicable for transit of type SDA_LISP_PUB_SUB_TRANSIT.
	ControlPlaneNetworkDeviceIDs  []string `json:"controlPlaneNetworkDeviceIds,omitempty"`  // List of network device IDs that will be used as control plane nodes. Maximum 2 network device IDs can be provided for transit of type SDA_LISP_BGP_TRANSIT and maximum 4 network device IDs can be provided for transit of type SDA_LISP_PUB_SUB_TRANSIT.
}
type RequestSdaAddVirtualNetworkWithScalableGroupsV1 struct {
	VirtualNetworkName    string   `json:"virtualNetworkName,omitempty"`    // Virtual Network Name to be assigned at global level
	IsGuestVirtualNetwork *bool    `json:"isGuestVirtualNetwork,omitempty"` // Guest Virtual Network enablement flag, default value is False.
	ScalableGroupNames    []string `json:"scalableGroupNames,omitempty"`    // Scalable Group to be associated to virtual network
	VManageVpnID          string   `json:"vManageVpnId,omitempty"`          // vManage vpn id for SD-WAN
}
type RequestSdaUpdateVirtualNetworkWithScalableGroupsV1 struct {
	VirtualNetworkName    string   `json:"virtualNetworkName,omitempty"`    // Virtual Network Name to be assigned global level
	IsGuestVirtualNetwork *bool    `json:"isGuestVirtualNetwork,omitempty"` // Indicates whether to set this as guest virtual network or not, default value is False.
	ScalableGroupNames    []string `json:"scalableGroupNames,omitempty"`    // Scalable Group Name to be associated to virtual network
	VManageVpnID          string   `json:"vManageVpnId,omitempty"`          // vManage vpn id for SD-WAN
}

//GetDefaultAuthenticationProfileFromSdaFabricV1 Get default authentication profile from SDA Fabric - 8b90-8a4e-4c5a-9a23
/* Get default authentication profile from SDA Fabric


@param GetDefaultAuthenticationProfileFromSDAFabricV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-default-authentication-profile-from-sda-fabric-v1
*/
func (s *SdaService) GetDefaultAuthenticationProfileFromSdaFabricV1(GetDefaultAuthenticationProfileFromSDAFabricV1QueryParams *GetDefaultAuthenticationProfileFromSdaFabricV1QueryParams) (*ResponseSdaGetDefaultAuthenticationProfileFromSdaFabricV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/authentication-profile"

	queryString, _ := query.Values(GetDefaultAuthenticationProfileFromSDAFabricV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetDefaultAuthenticationProfileFromSdaFabricV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDefaultAuthenticationProfileFromSdaFabricV1(GetDefaultAuthenticationProfileFromSDAFabricV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDefaultAuthenticationProfileFromSdaFabricV1")
	}

	result := response.Result().(*ResponseSdaGetDefaultAuthenticationProfileFromSdaFabricV1)
	return result, response, err

}

//GetBorderDeviceDetailFromSdaFabricV1 Get border device detail from SDA Fabric - 98a3-9bf4-485a-9871
/* Get border device detail from SDA Fabric


@param GetBorderDeviceDetailFromSDAFabricV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-border-device-detail-from-sda-fabric-v1
*/
func (s *SdaService) GetBorderDeviceDetailFromSdaFabricV1(GetBorderDeviceDetailFromSDAFabricV1QueryParams *GetBorderDeviceDetailFromSdaFabricV1QueryParams) (*ResponseSdaGetBorderDeviceDetailFromSdaFabricV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/border-device"

	queryString, _ := query.Values(GetBorderDeviceDetailFromSDAFabricV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetBorderDeviceDetailFromSdaFabricV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetBorderDeviceDetailFromSdaFabricV1(GetBorderDeviceDetailFromSDAFabricV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetBorderDeviceDetailFromSdaFabricV1")
	}

	result := response.Result().(*ResponseSdaGetBorderDeviceDetailFromSdaFabricV1)
	return result, response, err

}

//GetControlPlaneDeviceFromSdaFabricV1 Get control plane device from SDA Fabric - aba4-991d-4e9b-8747
/* Get control plane device from SDA Fabric


@param GetControlPlaneDeviceFromSDAFabricV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-control-plane-device-from-sda-fabric-v1
*/
func (s *SdaService) GetControlPlaneDeviceFromSdaFabricV1(GetControlPlaneDeviceFromSDAFabricV1QueryParams *GetControlPlaneDeviceFromSdaFabricV1QueryParams) (*ResponseSdaGetControlPlaneDeviceFromSdaFabricV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/control-plane-device"

	queryString, _ := query.Values(GetControlPlaneDeviceFromSDAFabricV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetControlPlaneDeviceFromSdaFabricV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetControlPlaneDeviceFromSdaFabricV1(GetControlPlaneDeviceFromSDAFabricV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetControlPlaneDeviceFromSdaFabricV1")
	}

	result := response.Result().(*ResponseSdaGetControlPlaneDeviceFromSdaFabricV1)
	return result, response, err

}

//GetDeviceInfoFromSdaFabricV1 Get device info from SDA Fabric - 1385-18e1-4069-ab5f
/* Get device info from SDA Fabric


@param GetDeviceInfoFromSDAFabricV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-info-from-sda-fabric-v1
*/
func (s *SdaService) GetDeviceInfoFromSdaFabricV1(GetDeviceInfoFromSDAFabricV1QueryParams *GetDeviceInfoFromSdaFabricV1QueryParams) (*ResponseSdaGetDeviceInfoFromSdaFabricV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/device"

	queryString, _ := query.Values(GetDeviceInfoFromSDAFabricV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetDeviceInfoFromSdaFabricV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceInfoFromSdaFabricV1(GetDeviceInfoFromSDAFabricV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceInfoFromSdaFabricV1")
	}

	result := response.Result().(*ResponseSdaGetDeviceInfoFromSdaFabricV1)
	return result, response, err

}

//GetDeviceRoleInSdaFabricV1 Get device role in SDA Fabric - 8a92-d87c-416a-8e83
/* Get device role in SDA Fabric


@param GetDeviceRoleInSDAFabricV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-role-in-sda-fabric-v1
*/
func (s *SdaService) GetDeviceRoleInSdaFabricV1(GetDeviceRoleInSDAFabricV1QueryParams *GetDeviceRoleInSdaFabricV1QueryParams) (*ResponseSdaGetDeviceRoleInSdaFabricV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/device/role"

	queryString, _ := query.Values(GetDeviceRoleInSDAFabricV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetDeviceRoleInSdaFabricV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceRoleInSdaFabricV1(GetDeviceRoleInSDAFabricV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceRoleInSdaFabricV1")
	}

	result := response.Result().(*ResponseSdaGetDeviceRoleInSdaFabricV1)
	return result, response, err

}

//GetEdgeDeviceFromSdaFabricV1 Get edge device from SDA Fabric - 7683-f90b-4efa-b090
/* Get edge device from SDA Fabric


@param GetEdgeDeviceFromSDAFabricV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-edge-device-from-sda-fabric-v1
*/
func (s *SdaService) GetEdgeDeviceFromSdaFabricV1(GetEdgeDeviceFromSDAFabricV1QueryParams *GetEdgeDeviceFromSdaFabricV1QueryParams) (*ResponseSdaGetEdgeDeviceFromSdaFabricV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/edge-device"

	queryString, _ := query.Values(GetEdgeDeviceFromSDAFabricV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetEdgeDeviceFromSdaFabricV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetEdgeDeviceFromSdaFabricV1(GetEdgeDeviceFromSDAFabricV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetEdgeDeviceFromSdaFabricV1")
	}

	result := response.Result().(*ResponseSdaGetEdgeDeviceFromSdaFabricV1)
	return result, response, err

}

//GetSiteFromSdaFabricV1 Get Site from SDA Fabric - 80b7-f8e6-406a-8701
/* Get Site info from SDA Fabric


@param GetSiteFromSDAFabricV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-site-from-sda-fabric-v1
*/
func (s *SdaService) GetSiteFromSdaFabricV1(GetSiteFromSDAFabricV1QueryParams *GetSiteFromSdaFabricV1QueryParams) (*ResponseSdaGetSiteFromSdaFabricV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/fabric-site"

	queryString, _ := query.Values(GetSiteFromSDAFabricV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetSiteFromSdaFabricV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSiteFromSdaFabricV1(GetSiteFromSDAFabricV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSiteFromSdaFabricV1")
	}

	result := response.Result().(*ResponseSdaGetSiteFromSdaFabricV1)
	return result, response, err

}

//GetPortAssignmentForAccessPointInSdaFabricV1 Get Port assignment for access point in SDA Fabric - 5097-f8d4-45f9-8f51
/* Get Port assignment for access point in SDA Fabric


@param GetPortAssignmentForAccessPointInSDAFabricV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-port-assignment-for-access-point-in-sda-fabric-v1
*/
func (s *SdaService) GetPortAssignmentForAccessPointInSdaFabricV1(GetPortAssignmentForAccessPointInSDAFabricV1QueryParams *GetPortAssignmentForAccessPointInSdaFabricV1QueryParams) (*ResponseSdaGetPortAssignmentForAccessPointInSdaFabricV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/hostonboarding/access-point"

	queryString, _ := query.Values(GetPortAssignmentForAccessPointInSDAFabricV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetPortAssignmentForAccessPointInSdaFabricV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPortAssignmentForAccessPointInSdaFabricV1(GetPortAssignmentForAccessPointInSDAFabricV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetPortAssignmentForAccessPointInSdaFabricV1")
	}

	result := response.Result().(*ResponseSdaGetPortAssignmentForAccessPointInSdaFabricV1)
	return result, response, err

}

//GetPortAssignmentForUserDeviceInSdaFabricV1 Get Port assignment for user device in SDA Fabric - a4a1-e8ed-41cb-9653
/* Get Port assignment for user device in SDA Fabric.


@param GetPortAssignmentForUserDeviceInSDAFabricV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-port-assignment-for-user-device-in-sda-fabric-v1
*/
func (s *SdaService) GetPortAssignmentForUserDeviceInSdaFabricV1(GetPortAssignmentForUserDeviceInSDAFabricV1QueryParams *GetPortAssignmentForUserDeviceInSdaFabricV1QueryParams) (*ResponseSdaGetPortAssignmentForUserDeviceInSdaFabricV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/hostonboarding/user-device"

	queryString, _ := query.Values(GetPortAssignmentForUserDeviceInSDAFabricV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetPortAssignmentForUserDeviceInSdaFabricV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPortAssignmentForUserDeviceInSdaFabricV1(GetPortAssignmentForUserDeviceInSDAFabricV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetPortAssignmentForUserDeviceInSdaFabricV1")
	}

	result := response.Result().(*ResponseSdaGetPortAssignmentForUserDeviceInSdaFabricV1)
	return result, response, err

}

//GetMulticastDetailsFromSdaFabricV1 Get multicast details from SDA fabric - c286-f98b-47ba-9ab4
/* Get multicast details from SDA fabric


@param GetMulticastDetailsFromSDAFabricV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-multicast-details-from-sda-fabric-v1
*/
func (s *SdaService) GetMulticastDetailsFromSdaFabricV1(GetMulticastDetailsFromSDAFabricV1QueryParams *GetMulticastDetailsFromSdaFabricV1QueryParams) (*ResponseSdaGetMulticastDetailsFromSdaFabricV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/multicast"

	queryString, _ := query.Values(GetMulticastDetailsFromSDAFabricV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetMulticastDetailsFromSdaFabricV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetMulticastDetailsFromSdaFabricV1(GetMulticastDetailsFromSDAFabricV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetMulticastDetailsFromSdaFabricV1")
	}

	result := response.Result().(*ResponseSdaGetMulticastDetailsFromSdaFabricV1)
	return result, response, err

}

//GetProvisionedWiredDeviceV1 Get Provisioned Wired Device - dfbf-2ae2-42ca-a449
/* Get Provisioned Wired Device


@param GetProvisionedWiredDeviceV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-provisioned-wired-device-v1
*/
func (s *SdaService) GetProvisionedWiredDeviceV1(GetProvisionedWiredDeviceV1QueryParams *GetProvisionedWiredDeviceV1QueryParams) (*ResponseSdaGetProvisionedWiredDeviceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/provision-device"

	queryString, _ := query.Values(GetProvisionedWiredDeviceV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetProvisionedWiredDeviceV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetProvisionedWiredDeviceV1(GetProvisionedWiredDeviceV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetProvisionedWiredDeviceV1")
	}

	result := response.Result().(*ResponseSdaGetProvisionedWiredDeviceV1)
	return result, response, err

}

//GetTransitPeerNetworkInfoV1 Get Transit Peer Network Info - 16a1-bb5d-48cb-873d
/* Get Transit Peer Network Info from SD-Access


@param GetTransitPeerNetworkInfoV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-transit-peer-network-info-v1
*/
func (s *SdaService) GetTransitPeerNetworkInfoV1(GetTransitPeerNetworkInfoV1QueryParams *GetTransitPeerNetworkInfoV1QueryParams) (*ResponseSdaGetTransitPeerNetworkInfoV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/transit-peer-network"

	queryString, _ := query.Values(GetTransitPeerNetworkInfoV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetTransitPeerNetworkInfoV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTransitPeerNetworkInfoV1(GetTransitPeerNetworkInfoV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTransitPeerNetworkInfoV1")
	}

	result := response.Result().(*ResponseSdaGetTransitPeerNetworkInfoV1)
	return result, response, err

}

//GetVnFromSdaFabricV1 Get VN from SDA Fabric - 2eb1-fa1e-49ca-a2b4
/* Get virtual network (VN) from SDA Fabric


@param GetVNFromSDAFabricV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-vn-from-sda-fabric-v1
*/
func (s *SdaService) GetVnFromSdaFabricV1(GetVNFromSDAFabricV1QueryParams *GetVnFromSdaFabricV1QueryParams) (*ResponseSdaGetVnFromSdaFabricV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/virtual-network"

	queryString, _ := query.Values(GetVNFromSDAFabricV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetVnFromSdaFabricV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetVnFromSdaFabricV1(GetVNFromSDAFabricV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetVnFromSdaFabricV1")
	}

	result := response.Result().(*ResponseSdaGetVnFromSdaFabricV1)
	return result, response, err

}

//GetVirtualNetworkSummaryV1 Get Virtual Network Summary - 6fa0-f8d5-4d29-857a
/* Get Virtual Network Summary


@param GetVirtualNetworkSummaryV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-virtual-network-summary-v1
*/
func (s *SdaService) GetVirtualNetworkSummaryV1(GetVirtualNetworkSummaryV1QueryParams *GetVirtualNetworkSummaryV1QueryParams) (*ResponseSdaGetVirtualNetworkSummaryV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/virtual-network/summary"

	queryString, _ := query.Values(GetVirtualNetworkSummaryV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetVirtualNetworkSummaryV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetVirtualNetworkSummaryV1(GetVirtualNetworkSummaryV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetVirtualNetworkSummaryV1")
	}

	result := response.Result().(*ResponseSdaGetVirtualNetworkSummaryV1)
	return result, response, err

}

//GetIPPoolFromSdaVirtualNetworkV1 Get IP Pool from SDA Virtual Network - fa92-19bf-45c8-b43b
/* Get IP Pool from SDA Virtual Network


@param GetIPPoolFromSDAVirtualNetworkV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-ip-pool-from-sda-virtual-network-v1
*/
func (s *SdaService) GetIPPoolFromSdaVirtualNetworkV1(GetIPPoolFromSDAVirtualNetworkV1QueryParams *GetIPPoolFromSdaVirtualNetworkV1QueryParams) (*ResponseSdaGetIPPoolFromSdaVirtualNetworkV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/virtualnetwork/ippool"

	queryString, _ := query.Values(GetIPPoolFromSDAVirtualNetworkV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetIPPoolFromSdaVirtualNetworkV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetIPPoolFromSdaVirtualNetworkV1(GetIPPoolFromSDAVirtualNetworkV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetIpPoolFromSdaVirtualNetworkV1")
	}

	result := response.Result().(*ResponseSdaGetIPPoolFromSdaVirtualNetworkV1)
	return result, response, err

}

//GetAnycastGatewaysV1 Get anycast gateways - 5cb3-f980-670e-770a
/* Returns a list of anycast gateways that match the provided query parameters.


@param GetAnycastGatewaysV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-anycast-gateways-v1
*/
func (s *SdaService) GetAnycastGatewaysV1(GetAnycastGatewaysV1QueryParams *GetAnycastGatewaysV1QueryParams) (*ResponseSdaGetAnycastGatewaysV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/anycastGateways"

	queryString, _ := query.Values(GetAnycastGatewaysV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetAnycastGatewaysV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAnycastGatewaysV1(GetAnycastGatewaysV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAnycastGatewaysV1")
	}

	result := response.Result().(*ResponseSdaGetAnycastGatewaysV1)
	return result, response, err

}

//GetAnycastGatewayCountV1 Get anycast gateway count - e504-152d-3f53-4d07
/* Returns the count of anycast gateways that match the provided query parameters.


@param GetAnycastGatewayCountV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-anycast-gateway-count-v1
*/
func (s *SdaService) GetAnycastGatewayCountV1(GetAnycastGatewayCountV1QueryParams *GetAnycastGatewayCountV1QueryParams) (*ResponseSdaGetAnycastGatewayCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/anycastGateways/count"

	queryString, _ := query.Values(GetAnycastGatewayCountV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetAnycastGatewayCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAnycastGatewayCountV1(GetAnycastGatewayCountV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAnycastGatewayCountV1")
	}

	result := response.Result().(*ResponseSdaGetAnycastGatewayCountV1)
	return result, response, err

}

//GetAuthenticationProfilesV1 Get authentication profiles - 9eb7-1a2d-44c8-82aa
/* Returns a list of authentication profiles that match the provided query parameters.


@param GetAuthenticationProfilesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-authentication-profiles-v1
*/
func (s *SdaService) GetAuthenticationProfilesV1(GetAuthenticationProfilesV1QueryParams *GetAuthenticationProfilesV1QueryParams) (*ResponseSdaGetAuthenticationProfilesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/authenticationProfiles"

	queryString, _ := query.Values(GetAuthenticationProfilesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetAuthenticationProfilesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAuthenticationProfilesV1(GetAuthenticationProfilesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAuthenticationProfilesV1")
	}

	result := response.Result().(*ResponseSdaGetAuthenticationProfilesV1)
	return result, response, err

}

//GetExtranetPoliciesV1 Get extranet policies - 3f85-3834-4b1b-bbcb
/* Returns a list of extranet policies that match the provided query parameters.


@param GetExtranetPoliciesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-extranet-policies-v1
*/
func (s *SdaService) GetExtranetPoliciesV1(GetExtranetPoliciesV1QueryParams *GetExtranetPoliciesV1QueryParams) (*ResponseSdaGetExtranetPoliciesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/extranetPolicies"

	queryString, _ := query.Values(GetExtranetPoliciesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetExtranetPoliciesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetExtranetPoliciesV1(GetExtranetPoliciesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetExtranetPoliciesV1")
	}

	result := response.Result().(*ResponseSdaGetExtranetPoliciesV1)
	return result, response, err

}

//GetExtranetPolicyCountV1 Get extranet policy count - 35a7-4975-447a-a6b8
/* Returns the count of extranet policies that match the provided query parameters.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-extranet-policy-count-v1
*/
func (s *SdaService) GetExtranetPolicyCountV1() (*ResponseSdaGetExtranetPolicyCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/extranetPolicies/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSdaGetExtranetPolicyCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetExtranetPolicyCountV1()
		}
		return nil, response, fmt.Errorf("error with operation GetExtranetPolicyCountV1")
	}

	result := response.Result().(*ResponseSdaGetExtranetPolicyCountV1)
	return result, response, err

}

//GetFabricDevicesV1 Get fabric devices - e680-7a97-47db-99e5
/* Returns a list of fabric devices that match the provided query parameters.


@param GetFabricDevicesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-fabric-devices-v1
*/
func (s *SdaService) GetFabricDevicesV1(GetFabricDevicesV1QueryParams *GetFabricDevicesV1QueryParams) (*ResponseSdaGetFabricDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricDevices"

	queryString, _ := query.Values(GetFabricDevicesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetFabricDevicesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFabricDevicesV1(GetFabricDevicesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetFabricDevicesV1")
	}

	result := response.Result().(*ResponseSdaGetFabricDevicesV1)
	return result, response, err

}

//GetFabricDevicesCountV1 Get fabric devices count - 9ba6-7b73-44b9-bb42
/* Returns the count of fabric devices that match the provided query parameters.


@param GetFabricDevicesCountV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-fabric-devices-count-v1
*/
func (s *SdaService) GetFabricDevicesCountV1(GetFabricDevicesCountV1QueryParams *GetFabricDevicesCountV1QueryParams) (*ResponseSdaGetFabricDevicesCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricDevices/count"

	queryString, _ := query.Values(GetFabricDevicesCountV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetFabricDevicesCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFabricDevicesCountV1(GetFabricDevicesCountV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetFabricDevicesCountV1")
	}

	result := response.Result().(*ResponseSdaGetFabricDevicesCountV1)
	return result, response, err

}

//GetFabricDevicesLayer2HandoffsV1 Get fabric devices layer 2 handoffs - b7af-eb15-4409-86a4
/* Returns a list of layer 2 handoffs of fabric devices that match the provided query parameters.


@param GetFabricDevicesLayer2HandoffsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-fabric-devices-layer2-handoffs-v1
*/
func (s *SdaService) GetFabricDevicesLayer2HandoffsV1(GetFabricDevicesLayer2HandoffsV1QueryParams *GetFabricDevicesLayer2HandoffsV1QueryParams) (*ResponseSdaGetFabricDevicesLayer2HandoffsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricDevices/layer2Handoffs"

	queryString, _ := query.Values(GetFabricDevicesLayer2HandoffsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetFabricDevicesLayer2HandoffsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFabricDevicesLayer2HandoffsV1(GetFabricDevicesLayer2HandoffsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetFabricDevicesLayer2HandoffsV1")
	}

	result := response.Result().(*ResponseSdaGetFabricDevicesLayer2HandoffsV1)
	return result, response, err

}

//GetFabricDevicesLayer2HandoffsCountV1 Get fabric devices layer 2 handoffs count - 019c-791b-48f9-b1d9
/* Returns the count of layer 2 handoffs of fabric devices that match the provided query parameters.


@param GetFabricDevicesLayer2HandoffsCountV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-fabric-devices-layer2-handoffs-count-v1
*/
func (s *SdaService) GetFabricDevicesLayer2HandoffsCountV1(GetFabricDevicesLayer2HandoffsCountV1QueryParams *GetFabricDevicesLayer2HandoffsCountV1QueryParams) (*ResponseSdaGetFabricDevicesLayer2HandoffsCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricDevices/layer2Handoffs/count"

	queryString, _ := query.Values(GetFabricDevicesLayer2HandoffsCountV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetFabricDevicesLayer2HandoffsCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFabricDevicesLayer2HandoffsCountV1(GetFabricDevicesLayer2HandoffsCountV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetFabricDevicesLayer2HandoffsCountV1")
	}

	result := response.Result().(*ResponseSdaGetFabricDevicesLayer2HandoffsCountV1)
	return result, response, err

}

//GetFabricDevicesLayer3HandoffsWithIPTransitV1 Get fabric devices layer 3 handoffs with ip transit - cbb9-daa0-43a9-913b
/* Returns a list of layer 3 handoffs with ip transit of fabric devices that match the provided query parameters.


@param GetFabricDevicesLayer3HandoffsWithIpTransitV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-fabric-devices-layer3-handoffs-with-ip-transit-v1
*/
func (s *SdaService) GetFabricDevicesLayer3HandoffsWithIPTransitV1(GetFabricDevicesLayer3HandoffsWithIpTransitV1QueryParams *GetFabricDevicesLayer3HandoffsWithIPTransitV1QueryParams) (*ResponseSdaGetFabricDevicesLayer3HandoffsWithIPTransitV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricDevices/layer3Handoffs/ipTransits"

	queryString, _ := query.Values(GetFabricDevicesLayer3HandoffsWithIpTransitV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetFabricDevicesLayer3HandoffsWithIPTransitV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFabricDevicesLayer3HandoffsWithIPTransitV1(GetFabricDevicesLayer3HandoffsWithIpTransitV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetFabricDevicesLayer3HandoffsWithIpTransitV1")
	}

	result := response.Result().(*ResponseSdaGetFabricDevicesLayer3HandoffsWithIPTransitV1)
	return result, response, err

}

//GetFabricDevicesLayer3HandoffsWithIPTransitCountV1 Get fabric devices layer 3 handoffs with ip transit count - bb90-4a31-4378-9125
/* Returns the count of layer 3 handoffs with ip transit of fabric devices that match the provided query parameters.


@param GetFabricDevicesLayer3HandoffsWithIpTransitCountV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-fabric-devices-layer3-handoffs-with-ip-transit-count-v1
*/
func (s *SdaService) GetFabricDevicesLayer3HandoffsWithIPTransitCountV1(GetFabricDevicesLayer3HandoffsWithIpTransitCountV1QueryParams *GetFabricDevicesLayer3HandoffsWithIPTransitCountV1QueryParams) (*ResponseSdaGetFabricDevicesLayer3HandoffsWithIPTransitCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricDevices/layer3Handoffs/ipTransits/count"

	queryString, _ := query.Values(GetFabricDevicesLayer3HandoffsWithIpTransitCountV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetFabricDevicesLayer3HandoffsWithIPTransitCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFabricDevicesLayer3HandoffsWithIPTransitCountV1(GetFabricDevicesLayer3HandoffsWithIpTransitCountV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetFabricDevicesLayer3HandoffsWithIpTransitCountV1")
	}

	result := response.Result().(*ResponseSdaGetFabricDevicesLayer3HandoffsWithIPTransitCountV1)
	return result, response, err

}

//GetFabricDevicesLayer3HandoffsWithSdaTransitV1 Get fabric devices layer 3 handoffs with sda transit - 0d8e-d8dd-458b-9dc1
/* Returns a list of layer 3 handoffs with sda transit of fabric devices that match the provided query parameters.


@param GetFabricDevicesLayer3HandoffsWithSdaTransitV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-fabric-devices-layer3-handoffs-with-sda-transit-v1
*/
func (s *SdaService) GetFabricDevicesLayer3HandoffsWithSdaTransitV1(GetFabricDevicesLayer3HandoffsWithSdaTransitV1QueryParams *GetFabricDevicesLayer3HandoffsWithSdaTransitV1QueryParams) (*ResponseSdaGetFabricDevicesLayer3HandoffsWithSdaTransitV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricDevices/layer3Handoffs/sdaTransits"

	queryString, _ := query.Values(GetFabricDevicesLayer3HandoffsWithSdaTransitV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetFabricDevicesLayer3HandoffsWithSdaTransitV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFabricDevicesLayer3HandoffsWithSdaTransitV1(GetFabricDevicesLayer3HandoffsWithSdaTransitV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetFabricDevicesLayer3HandoffsWithSdaTransitV1")
	}

	result := response.Result().(*ResponseSdaGetFabricDevicesLayer3HandoffsWithSdaTransitV1)
	return result, response, err

}

//GetFabricDevicesLayer3HandoffsWithSdaTransitCountV1 Get fabric devices layer 3 handoffs with sda transit count - bd89-6aca-46cb-8f65
/* Returns the count of layer 3 handoffs with sda transit of fabric devices that match the provided query parameters.


@param GetFabricDevicesLayer3HandoffsWithSdaTransitCountV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-fabric-devices-layer3-handoffs-with-sda-transit-count-v1
*/
func (s *SdaService) GetFabricDevicesLayer3HandoffsWithSdaTransitCountV1(GetFabricDevicesLayer3HandoffsWithSdaTransitCountV1QueryParams *GetFabricDevicesLayer3HandoffsWithSdaTransitCountV1QueryParams) (*ResponseSdaGetFabricDevicesLayer3HandoffsWithSdaTransitCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricDevices/layer3Handoffs/sdaTransits/count"

	queryString, _ := query.Values(GetFabricDevicesLayer3HandoffsWithSdaTransitCountV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetFabricDevicesLayer3HandoffsWithSdaTransitCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFabricDevicesLayer3HandoffsWithSdaTransitCountV1(GetFabricDevicesLayer3HandoffsWithSdaTransitCountV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetFabricDevicesLayer3HandoffsWithSdaTransitCountV1")
	}

	result := response.Result().(*ResponseSdaGetFabricDevicesLayer3HandoffsWithSdaTransitCountV1)
	return result, response, err

}

//GetFabricSitesV1 Get fabric sites - b78b-fa87-49a9-b804
/* Returns a list of fabric sites that match the provided query parameters.


@param GetFabricSitesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-fabric-sites-v1
*/
func (s *SdaService) GetFabricSitesV1(GetFabricSitesV1QueryParams *GetFabricSitesV1QueryParams) (*ResponseSdaGetFabricSitesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricSites"

	queryString, _ := query.Values(GetFabricSitesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetFabricSitesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFabricSitesV1(GetFabricSitesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetFabricSitesV1")
	}

	result := response.Result().(*ResponseSdaGetFabricSitesV1)
	return result, response, err

}

//GetFabricSiteCountV1 Get fabric site count - 109a-0907-4a9b-82ef
/* Returns the count of fabric sites that match the provided query parameters.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-fabric-site-count-v1
*/
func (s *SdaService) GetFabricSiteCountV1() (*ResponseSdaGetFabricSiteCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricSites/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSdaGetFabricSiteCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFabricSiteCountV1()
		}
		return nil, response, fmt.Errorf("error with operation GetFabricSiteCountV1")
	}

	result := response.Result().(*ResponseSdaGetFabricSiteCountV1)
	return result, response, err

}

//GetFabricZonesV1 Get fabric zones - d0bc-0b5c-4fdb-839a
/* Returns a list of fabric zones that match the provided query parameters.


@param GetFabricZonesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-fabric-zones-v1
*/
func (s *SdaService) GetFabricZonesV1(GetFabricZonesV1QueryParams *GetFabricZonesV1QueryParams) (*ResponseSdaGetFabricZonesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricZones"

	queryString, _ := query.Values(GetFabricZonesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetFabricZonesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFabricZonesV1(GetFabricZonesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetFabricZonesV1")
	}

	result := response.Result().(*ResponseSdaGetFabricZonesV1)
	return result, response, err

}

//GetFabricZoneCountV1 Get fabric zone count - 15a2-da20-4758-bc78
/* Returns the count of fabric zones that match the provided query parameters.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-fabric-zone-count-v1
*/
func (s *SdaService) GetFabricZoneCountV1() (*ResponseSdaGetFabricZoneCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricZones/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSdaGetFabricZoneCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFabricZoneCountV1()
		}
		return nil, response, fmt.Errorf("error with operation GetFabricZoneCountV1")
	}

	result := response.Result().(*ResponseSdaGetFabricZoneCountV1)
	return result, response, err

}

//GetLayer2VirtualNetworksV1 Get layer 2 virtual networks - 659a-ab00-4c69-a663
/* Returns a list of layer 2 virtual networks that match the provided query parameters.


@param GetLayer2VirtualNetworksV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-layer2-virtual-networks-v1
*/
func (s *SdaService) GetLayer2VirtualNetworksV1(GetLayer2VirtualNetworksV1QueryParams *GetLayer2VirtualNetworksV1QueryParams) (*ResponseSdaGetLayer2VirtualNetworksV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/layer2VirtualNetworks"

	queryString, _ := query.Values(GetLayer2VirtualNetworksV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetLayer2VirtualNetworksV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetLayer2VirtualNetworksV1(GetLayer2VirtualNetworksV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetLayer2VirtualNetworksV1")
	}

	result := response.Result().(*ResponseSdaGetLayer2VirtualNetworksV1)
	return result, response, err

}

//GetLayer2VirtualNetworkCountV1 Get layer 2 virtual network count - 5c9f-0a6e-445b-b743
/* Returns the count of layer 2 virtual networks that match the provided query parameters.


@param GetLayer2VirtualNetworkCountV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-layer2-virtual-network-count-v1
*/
func (s *SdaService) GetLayer2VirtualNetworkCountV1(GetLayer2VirtualNetworkCountV1QueryParams *GetLayer2VirtualNetworkCountV1QueryParams) (*ResponseSdaGetLayer2VirtualNetworkCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/layer2VirtualNetworks/count"

	queryString, _ := query.Values(GetLayer2VirtualNetworkCountV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetLayer2VirtualNetworkCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetLayer2VirtualNetworkCountV1(GetLayer2VirtualNetworkCountV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetLayer2VirtualNetworkCountV1")
	}

	result := response.Result().(*ResponseSdaGetLayer2VirtualNetworkCountV1)
	return result, response, err

}

//GetLayer3VirtualNetworksV1 Get layer 3 virtual networks - 2892-e9d4-4b68-b538
/* Returns a list of layer 3 virtual networks that match the provided query parameters.


@param GetLayer3VirtualNetworksV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-layer3-virtual-networks-v1
*/
func (s *SdaService) GetLayer3VirtualNetworksV1(GetLayer3VirtualNetworksV1QueryParams *GetLayer3VirtualNetworksV1QueryParams) (*ResponseSdaGetLayer3VirtualNetworksV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/layer3VirtualNetworks"

	queryString, _ := query.Values(GetLayer3VirtualNetworksV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetLayer3VirtualNetworksV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetLayer3VirtualNetworksV1(GetLayer3VirtualNetworksV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetLayer3VirtualNetworksV1")
	}

	result := response.Result().(*ResponseSdaGetLayer3VirtualNetworksV1)
	return result, response, err

}

//GetLayer3VirtualNetworksCountV1 Get layer 3 virtual networks count - 87af-99e9-493b-9f3d
/* Returns the count of layer 3 virtual networks that match the provided query parameters.


@param GetLayer3VirtualNetworksCountV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-layer3-virtual-networks-count-v1
*/
func (s *SdaService) GetLayer3VirtualNetworksCountV1(GetLayer3VirtualNetworksCountV1QueryParams *GetLayer3VirtualNetworksCountV1QueryParams) (*ResponseSdaGetLayer3VirtualNetworksCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/layer3VirtualNetworks/count"

	queryString, _ := query.Values(GetLayer3VirtualNetworksCountV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetLayer3VirtualNetworksCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetLayer3VirtualNetworksCountV1(GetLayer3VirtualNetworksCountV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetLayer3VirtualNetworksCountV1")
	}

	result := response.Result().(*ResponseSdaGetLayer3VirtualNetworksCountV1)
	return result, response, err

}

//GetMulticastV1 Get multicast - b48d-e933-4e5b-988a
/* Returns a list of multicast configurations at a fabric site level that match the provided query parameters.


@param GetMulticastV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-multicast-v1
*/
func (s *SdaService) GetMulticastV1(GetMulticastV1QueryParams *GetMulticastV1QueryParams) (*ResponseSdaGetMulticastV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/multicast"

	queryString, _ := query.Values(GetMulticastV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetMulticastV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetMulticastV1(GetMulticastV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetMulticastV1")
	}

	result := response.Result().(*ResponseSdaGetMulticastV1)
	return result, response, err

}

//GetMulticastVirtualNetworksV1 Get multicast virtual networks - 048b-698b-4048-b3aa
/* Returns a list of multicast configurations for virtual networks that match the provided query parameters.


@param GetMulticastVirtualNetworksV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-multicast-virtual-networks-v1
*/
func (s *SdaService) GetMulticastVirtualNetworksV1(GetMulticastVirtualNetworksV1QueryParams *GetMulticastVirtualNetworksV1QueryParams) (*ResponseSdaGetMulticastVirtualNetworksV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/multicast/virtualNetworks"

	queryString, _ := query.Values(GetMulticastVirtualNetworksV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetMulticastVirtualNetworksV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetMulticastVirtualNetworksV1(GetMulticastVirtualNetworksV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetMulticastVirtualNetworksV1")
	}

	result := response.Result().(*ResponseSdaGetMulticastVirtualNetworksV1)
	return result, response, err

}

//GetMulticastVirtualNetworkCountV1 Get multicast virtual network count - 7cbb-0b86-4a39-98ab
/* Returns the count of multicast configurations associated to virtual networks that match the provided query parameters.


@param GetMulticastVirtualNetworkCountV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-multicast-virtual-network-count-v1
*/
func (s *SdaService) GetMulticastVirtualNetworkCountV1(GetMulticastVirtualNetworkCountV1QueryParams *GetMulticastVirtualNetworkCountV1QueryParams) (*ResponseSdaGetMulticastVirtualNetworkCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/multicast/virtualNetworks/count"

	queryString, _ := query.Values(GetMulticastVirtualNetworkCountV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetMulticastVirtualNetworkCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetMulticastVirtualNetworkCountV1(GetMulticastVirtualNetworkCountV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetMulticastVirtualNetworkCountV1")
	}

	result := response.Result().(*ResponseSdaGetMulticastVirtualNetworkCountV1)
	return result, response, err

}

//GetPortAssignmentsV1 Get port assignments - c199-09a2-4619-a140
/* Returns a list of port assignments that match the provided query parameters.


@param GetPortAssignmentsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-port-assignments-v1
*/
func (s *SdaService) GetPortAssignmentsV1(GetPortAssignmentsV1QueryParams *GetPortAssignmentsV1QueryParams) (*ResponseSdaGetPortAssignmentsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/portAssignments"

	queryString, _ := query.Values(GetPortAssignmentsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetPortAssignmentsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPortAssignmentsV1(GetPortAssignmentsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetPortAssignmentsV1")
	}

	result := response.Result().(*ResponseSdaGetPortAssignmentsV1)
	return result, response, err

}

//GetPortAssignmentCountV1 Get port assignment count - 4587-0827-4f1b-a2d4
/* Returns the count of port assignments that match the provided query parameters.


@param GetPortAssignmentCountV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-port-assignment-count-v1
*/
func (s *SdaService) GetPortAssignmentCountV1(GetPortAssignmentCountV1QueryParams *GetPortAssignmentCountV1QueryParams) (*ResponseSdaGetPortAssignmentCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/portAssignments/count"

	queryString, _ := query.Values(GetPortAssignmentCountV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetPortAssignmentCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPortAssignmentCountV1(GetPortAssignmentCountV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetPortAssignmentCountV1")
	}

	result := response.Result().(*ResponseSdaGetPortAssignmentCountV1)
	return result, response, err

}

//GetPortChannelsV1 Get port channels - dea6-fbe3-4469-8d79
/* Returns a list of port channels that match the provided query parameters.


@param GetPortChannelsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-port-channels-v1
*/
func (s *SdaService) GetPortChannelsV1(GetPortChannelsV1QueryParams *GetPortChannelsV1QueryParams) (*ResponseSdaGetPortChannelsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/portChannels"

	queryString, _ := query.Values(GetPortChannelsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetPortChannelsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPortChannelsV1(GetPortChannelsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetPortChannelsV1")
	}

	result := response.Result().(*ResponseSdaGetPortChannelsV1)
	return result, response, err

}

//GetPortChannelCountV1 Get port channel count - 7ebb-88ff-4c2b-989c
/* Returns the count of port channels that match the provided query parameters.


@param GetPortChannelCountV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-port-channel-count-v1
*/
func (s *SdaService) GetPortChannelCountV1(GetPortChannelCountV1QueryParams *GetPortChannelCountV1QueryParams) (*ResponseSdaGetPortChannelCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/portChannels/count"

	queryString, _ := query.Values(GetPortChannelCountV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetPortChannelCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPortChannelCountV1(GetPortChannelCountV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetPortChannelCountV1")
	}

	result := response.Result().(*ResponseSdaGetPortChannelCountV1)
	return result, response, err

}

//GetProvisionedDevicesV1 Get provisioned devices - 99b3-ba27-4fe9-9e6b
/* Returns the list of provisioned devices based on query parameters.


@param GetProvisionedDevicesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-provisioned-devices-v1
*/
func (s *SdaService) GetProvisionedDevicesV1(GetProvisionedDevicesV1QueryParams *GetProvisionedDevicesV1QueryParams) (*ResponseSdaGetProvisionedDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/provisionDevices"

	queryString, _ := query.Values(GetProvisionedDevicesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetProvisionedDevicesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetProvisionedDevicesV1(GetProvisionedDevicesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetProvisionedDevicesV1")
	}

	result := response.Result().(*ResponseSdaGetProvisionedDevicesV1)
	return result, response, err

}

//GetProvisionedDevicesCountV1 Get Provisioned Devices count - e0b3-195e-4678-aeb4
/* Returns the count of provisioned devices based on query parameters.


@param GetProvisionedDevicesCountV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-provisioned-devices-count-v1
*/
func (s *SdaService) GetProvisionedDevicesCountV1(GetProvisionedDevicesCountV1QueryParams *GetProvisionedDevicesCountV1QueryParams) (*ResponseSdaGetProvisionedDevicesCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/provisionDevices/count"

	queryString, _ := query.Values(GetProvisionedDevicesCountV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetProvisionedDevicesCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetProvisionedDevicesCountV1(GetProvisionedDevicesCountV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetProvisionedDevicesCountV1")
	}

	result := response.Result().(*ResponseSdaGetProvisionedDevicesCountV1)
	return result, response, err

}

//GetTransitNetworksV1 Get transit networks - e395-fae2-4f0a-b11e
/* Returns a list of transit networks that match the provided query parameters.


@param GetTransitNetworksV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-transit-networks-v1
*/
func (s *SdaService) GetTransitNetworksV1(GetTransitNetworksV1QueryParams *GetTransitNetworksV1QueryParams) (*ResponseSdaGetTransitNetworksV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/transitNetworks"

	queryString, _ := query.Values(GetTransitNetworksV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetTransitNetworksV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTransitNetworksV1(GetTransitNetworksV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTransitNetworksV1")
	}

	result := response.Result().(*ResponseSdaGetTransitNetworksV1)
	return result, response, err

}

//GetTransitNetworksCountV1 Get transit networks count - 9397-d838-446b-b716
/* Returns the count of transit networks that match the provided query parameters.


@param GetTransitNetworksCountV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-transit-networks-count-v1
*/
func (s *SdaService) GetTransitNetworksCountV1(GetTransitNetworksCountV1QueryParams *GetTransitNetworksCountV1QueryParams) (*ResponseSdaGetTransitNetworksCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/transitNetworks/count"

	queryString, _ := query.Values(GetTransitNetworksCountV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetTransitNetworksCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTransitNetworksCountV1(GetTransitNetworksCountV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTransitNetworksCountV1")
	}

	result := response.Result().(*ResponseSdaGetTransitNetworksCountV1)
	return result, response, err

}

//GetVirtualNetworkWithScalableGroupsV1 Get virtual network with scalable groups - ec8a-1ab5-4eba-bca7
/* Get virtual network with scalable groups


@param GetVirtualNetworkWithScalableGroupsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-virtual-network-with-scalable-groups-v1
*/
func (s *SdaService) GetVirtualNetworkWithScalableGroupsV1(GetVirtualNetworkWithScalableGroupsV1QueryParams *GetVirtualNetworkWithScalableGroupsV1QueryParams) (*ResponseSdaGetVirtualNetworkWithScalableGroupsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/virtual-network"

	queryString, _ := query.Values(GetVirtualNetworkWithScalableGroupsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetVirtualNetworkWithScalableGroupsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetVirtualNetworkWithScalableGroupsV1(GetVirtualNetworkWithScalableGroupsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetVirtualNetworkWithScalableGroupsV1")
	}

	result := response.Result().(*ResponseSdaGetVirtualNetworkWithScalableGroupsV1)
	return result, response, err

}

//AddDefaultAuthenticationTemplateInSdaFabricV1 Add default authentication template in SDA Fabric - bca3-39d8-44c8-a3c0
/* Add default authentication template in SDA Fabric



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-default-authentication-template-in-sda-fabric-v1
*/
func (s *SdaService) AddDefaultAuthenticationTemplateInSdaFabricV1(requestSdaAddDefaultAuthenticationTemplateInSDAFabricV1 *RequestSdaAddDefaultAuthenticationTemplateInSdaFabricV1) (*ResponseSdaAddDefaultAuthenticationTemplateInSdaFabricV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/authentication-profile"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddDefaultAuthenticationTemplateInSDAFabricV1).
		SetResult(&ResponseSdaAddDefaultAuthenticationTemplateInSdaFabricV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddDefaultAuthenticationTemplateInSdaFabricV1(requestSdaAddDefaultAuthenticationTemplateInSDAFabricV1)
		}

		return nil, response, fmt.Errorf("error with operation AddDefaultAuthenticationTemplateInSdaFabricV1")
	}

	result := response.Result().(*ResponseSdaAddDefaultAuthenticationTemplateInSdaFabricV1)
	return result, response, err

}

//AddBorderDeviceInSdaFabricV1 Add border device in SDA Fabric - bead-7b34-43b9-96a7
/* Add border device in SDA Fabric



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-border-device-in-sda-fabric-v1
*/
func (s *SdaService) AddBorderDeviceInSdaFabricV1(requestSdaAddBorderDeviceInSDAFabricV1 *RequestSdaAddBorderDeviceInSdaFabricV1) (*ResponseSdaAddBorderDeviceInSdaFabricV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/border-device"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddBorderDeviceInSDAFabricV1).
		SetResult(&ResponseSdaAddBorderDeviceInSdaFabricV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddBorderDeviceInSdaFabricV1(requestSdaAddBorderDeviceInSDAFabricV1)
		}

		return nil, response, fmt.Errorf("error with operation AddBorderDeviceInSdaFabricV1")
	}

	result := response.Result().(*ResponseSdaAddBorderDeviceInSdaFabricV1)
	return result, response, err

}

//AddControlPlaneDeviceInSdaFabricV1 Add control plane device in SDA Fabric - dd85-c910-4248-9a3f
/* Add control plane device in SDA Fabric



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-control-plane-device-in-sda-fabric-v1
*/
func (s *SdaService) AddControlPlaneDeviceInSdaFabricV1(requestSdaAddControlPlaneDeviceInSDAFabricV1 *RequestSdaAddControlPlaneDeviceInSdaFabricV1) (*ResponseSdaAddControlPlaneDeviceInSdaFabricV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/control-plane-device"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddControlPlaneDeviceInSDAFabricV1).
		SetResult(&ResponseSdaAddControlPlaneDeviceInSdaFabricV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddControlPlaneDeviceInSdaFabricV1(requestSdaAddControlPlaneDeviceInSDAFabricV1)
		}

		return nil, response, fmt.Errorf("error with operation AddControlPlaneDeviceInSdaFabricV1")
	}

	result := response.Result().(*ResponseSdaAddControlPlaneDeviceInSdaFabricV1)
	return result, response, err

}

//AddEdgeDeviceInSdaFabricV1 Add edge device in SDA Fabric - 87a8-ba44-4ce9-bc59
/* Add edge device in SDA Fabric



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-edge-device-in-sda-fabric-v1
*/
func (s *SdaService) AddEdgeDeviceInSdaFabricV1(requestSdaAddEdgeDeviceInSDAFabricV1 *RequestSdaAddEdgeDeviceInSdaFabricV1) (*ResponseSdaAddEdgeDeviceInSdaFabricV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/edge-device"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddEdgeDeviceInSDAFabricV1).
		SetResult(&ResponseSdaAddEdgeDeviceInSdaFabricV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddEdgeDeviceInSdaFabricV1(requestSdaAddEdgeDeviceInSDAFabricV1)
		}

		return nil, response, fmt.Errorf("error with operation AddEdgeDeviceInSdaFabricV1")
	}

	result := response.Result().(*ResponseSdaAddEdgeDeviceInSdaFabricV1)
	return result, response, err

}

//AddSiteInSdaFabricV1 Add Site in SDA Fabric - d2b4-d9d0-4a4b-884c
/* Add Site in SDA Fabric



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-site-in-sda-fabric-v1
*/
func (s *SdaService) AddSiteInSdaFabricV1(requestSdaAddSiteInSDAFabricV1 *RequestSdaAddSiteInSdaFabricV1) (*ResponseSdaAddSiteInSdaFabricV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/fabric-site"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddSiteInSDAFabricV1).
		SetResult(&ResponseSdaAddSiteInSdaFabricV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddSiteInSdaFabricV1(requestSdaAddSiteInSDAFabricV1)
		}

		return nil, response, fmt.Errorf("error with operation AddSiteInSdaFabricV1")
	}

	result := response.Result().(*ResponseSdaAddSiteInSdaFabricV1)
	return result, response, err

}

//AddPortAssignmentForAccessPointInSdaFabricV1 Add Port assignment for access point in SDA Fabric - c2a4-3ad2-4098-baa7
/* Add Port assignment for access point in SDA Fabric



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-port-assignment-for-access-point-in-sda-fabric-v1
*/
func (s *SdaService) AddPortAssignmentForAccessPointInSdaFabricV1(requestSdaAddPortAssignmentForAccessPointInSDAFabricV1 *RequestSdaAddPortAssignmentForAccessPointInSdaFabricV1) (*ResponseSdaAddPortAssignmentForAccessPointInSdaFabricV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/hostonboarding/access-point"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddPortAssignmentForAccessPointInSDAFabricV1).
		SetResult(&ResponseSdaAddPortAssignmentForAccessPointInSdaFabricV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddPortAssignmentForAccessPointInSdaFabricV1(requestSdaAddPortAssignmentForAccessPointInSDAFabricV1)
		}

		return nil, response, fmt.Errorf("error with operation AddPortAssignmentForAccessPointInSdaFabricV1")
	}

	result := response.Result().(*ResponseSdaAddPortAssignmentForAccessPointInSdaFabricV1)
	return result, response, err

}

//AddPortAssignmentForUserDeviceInSdaFabricV1 Add Port assignment for user device in SDA Fabric - 9582-ab82-4ce8-b29d
/* Add Port assignment for user device in SDA Fabric.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-port-assignment-for-user-device-in-sda-fabric-v1
*/
func (s *SdaService) AddPortAssignmentForUserDeviceInSdaFabricV1(requestSdaAddPortAssignmentForUserDeviceInSDAFabricV1 *RequestSdaAddPortAssignmentForUserDeviceInSdaFabricV1) (*ResponseSdaAddPortAssignmentForUserDeviceInSdaFabricV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/hostonboarding/user-device"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddPortAssignmentForUserDeviceInSDAFabricV1).
		SetResult(&ResponseSdaAddPortAssignmentForUserDeviceInSdaFabricV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddPortAssignmentForUserDeviceInSdaFabricV1(requestSdaAddPortAssignmentForUserDeviceInSDAFabricV1)
		}

		return nil, response, fmt.Errorf("error with operation AddPortAssignmentForUserDeviceInSdaFabricV1")
	}

	result := response.Result().(*ResponseSdaAddPortAssignmentForUserDeviceInSdaFabricV1)
	return result, response, err

}

//AddMulticastInSdaFabricV1 Add multicast in SDA fabric - ff85-3826-472a-98fb
/* Add multicast in SDA fabric



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-multicast-in-sda-fabric-v1
*/
func (s *SdaService) AddMulticastInSdaFabricV1(requestSdaAddMulticastInSDAFabricV1 *RequestSdaAddMulticastInSdaFabricV1) (*ResponseSdaAddMulticastInSdaFabricV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/multicast"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddMulticastInSDAFabricV1).
		SetResult(&ResponseSdaAddMulticastInSdaFabricV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddMulticastInSdaFabricV1(requestSdaAddMulticastInSDAFabricV1)
		}

		return nil, response, fmt.Errorf("error with operation AddMulticastInSdaFabricV1")
	}

	result := response.Result().(*ResponseSdaAddMulticastInSdaFabricV1)
	return result, response, err

}

//ProvisionWiredDeviceV1 Provision Wired Device - cf9a-5843-45fa-9399
/* Provision Wired Device



Documentation Link: https://developer.cisco.com/docs/dna-center/#!provision-wired-device-v1
*/
func (s *SdaService) ProvisionWiredDeviceV1(requestSdaProvisionWiredDeviceV1 *RequestSdaProvisionWiredDeviceV1) (*ResponseSdaProvisionWiredDeviceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/provision-device"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaProvisionWiredDeviceV1).
		SetResult(&ResponseSdaProvisionWiredDeviceV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ProvisionWiredDeviceV1(requestSdaProvisionWiredDeviceV1)
		}

		return nil, response, fmt.Errorf("error with operation ProvisionWiredDeviceV1")
	}

	result := response.Result().(*ResponseSdaProvisionWiredDeviceV1)
	return result, response, err

}

//AddTransitPeerNetworkV1 Add Transit Peer Network - 6db9-292d-4f28-a26b
/* Add Transit Peer Network in SD-Access



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-transit-peer-network-v1
*/
func (s *SdaService) AddTransitPeerNetworkV1(requestSdaAddTransitPeerNetworkV1 *RequestSdaAddTransitPeerNetworkV1) (*ResponseSdaAddTransitPeerNetworkV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/transit-peer-network"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddTransitPeerNetworkV1).
		SetResult(&ResponseSdaAddTransitPeerNetworkV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddTransitPeerNetworkV1(requestSdaAddTransitPeerNetworkV1)
		}

		return nil, response, fmt.Errorf("error with operation AddTransitPeerNetworkV1")
	}

	result := response.Result().(*ResponseSdaAddTransitPeerNetworkV1)
	return result, response, err

}

//AddVnInFabricV1 Add VN in fabric - 518c-59cd-441a-a9fc
/* Add virtual network (VN) in SDA Fabric



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-vn-in-fabric-v1
*/
func (s *SdaService) AddVnInFabricV1(requestSdaAddVNInFabricV1 *RequestSdaAddVnInFabricV1) (*ResponseSdaAddVnInFabricV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/virtual-network"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddVNInFabricV1).
		SetResult(&ResponseSdaAddVnInFabricV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddVnInFabricV1(requestSdaAddVNInFabricV1)
		}

		return nil, response, fmt.Errorf("error with operation AddVnInFabricV1")
	}

	result := response.Result().(*ResponseSdaAddVnInFabricV1)
	return result, response, err

}

//AddIPPoolInSdaVirtualNetworkV1 Add IP Pool in SDA Virtual Network - 2085-79ea-4ed9-8f4f
/* Add IP Pool in SDA Virtual Network



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-ip-pool-in-sda-virtual-network-v1
*/
func (s *SdaService) AddIPPoolInSdaVirtualNetworkV1(requestSdaAddIPPoolInSDAVirtualNetworkV1 *RequestSdaAddIPPoolInSdaVirtualNetworkV1) (*ResponseSdaAddIPPoolInSdaVirtualNetworkV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/virtualnetwork/ippool"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddIPPoolInSDAVirtualNetworkV1).
		SetResult(&ResponseSdaAddIPPoolInSdaVirtualNetworkV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddIPPoolInSdaVirtualNetworkV1(requestSdaAddIPPoolInSDAVirtualNetworkV1)
		}

		return nil, response, fmt.Errorf("error with operation AddIpPoolInSdaVirtualNetworkV1")
	}

	result := response.Result().(*ResponseSdaAddIPPoolInSdaVirtualNetworkV1)
	return result, response, err

}

//AddAnycastGatewaysV1 Add anycast gateways - a6fa-2ce6-46ac-7ac5
/* Adds anycast gateways based on user input.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-anycast-gateways-v1
*/
func (s *SdaService) AddAnycastGatewaysV1(requestSdaAddAnycastGatewaysV1 *RequestSdaAddAnycastGatewaysV1) (*ResponseSdaAddAnycastGatewaysV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/anycastGateways"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddAnycastGatewaysV1).
		SetResult(&ResponseSdaAddAnycastGatewaysV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddAnycastGatewaysV1(requestSdaAddAnycastGatewaysV1)
		}

		return nil, response, fmt.Errorf("error with operation AddAnycastGatewaysV1")
	}

	result := response.Result().(*ResponseSdaAddAnycastGatewaysV1)
	return result, response, err

}

//AddExtranetPolicyV1 Add extranet policy - 1282-78e3-45e8-aae7
/* Adds an extranet policy based on user input.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-extranet-policy-v1
*/
func (s *SdaService) AddExtranetPolicyV1(requestSdaAddExtranetPolicyV1 *RequestSdaAddExtranetPolicyV1) (*ResponseSdaAddExtranetPolicyV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/extranetPolicies"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddExtranetPolicyV1).
		SetResult(&ResponseSdaAddExtranetPolicyV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddExtranetPolicyV1(requestSdaAddExtranetPolicyV1)
		}

		return nil, response, fmt.Errorf("error with operation AddExtranetPolicyV1")
	}

	result := response.Result().(*ResponseSdaAddExtranetPolicyV1)
	return result, response, err

}

//AddFabricDevicesV1 Add fabric devices - 698a-f8d2-4fd8-b091
/* Adds fabric devices based on user input.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-fabric-devices-v1
*/
func (s *SdaService) AddFabricDevicesV1(requestSdaAddFabricDevicesV1 *RequestSdaAddFabricDevicesV1) (*ResponseSdaAddFabricDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricDevices"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddFabricDevicesV1).
		SetResult(&ResponseSdaAddFabricDevicesV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddFabricDevicesV1(requestSdaAddFabricDevicesV1)
		}

		return nil, response, fmt.Errorf("error with operation AddFabricDevicesV1")
	}

	result := response.Result().(*ResponseSdaAddFabricDevicesV1)
	return result, response, err

}

//AddFabricDevicesLayer2HandoffsV1 Add fabric devices layer 2 handoffs - 60ae-1ba0-4ebb-9ae1
/* Adds layer 2 handoffs in fabric devices based on user input.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-fabric-devices-layer2-handoffs-v1
*/
func (s *SdaService) AddFabricDevicesLayer2HandoffsV1(requestSdaAddFabricDevicesLayer2HandoffsV1 *RequestSdaAddFabricDevicesLayer2HandoffsV1) (*ResponseSdaAddFabricDevicesLayer2HandoffsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricDevices/layer2Handoffs"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddFabricDevicesLayer2HandoffsV1).
		SetResult(&ResponseSdaAddFabricDevicesLayer2HandoffsV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddFabricDevicesLayer2HandoffsV1(requestSdaAddFabricDevicesLayer2HandoffsV1)
		}

		return nil, response, fmt.Errorf("error with operation AddFabricDevicesLayer2HandoffsV1")
	}

	result := response.Result().(*ResponseSdaAddFabricDevicesLayer2HandoffsV1)
	return result, response, err

}

//AddFabricDevicesLayer3HandoffsWithIPTransitV1 Add fabric devices layer 3 handoffs with ip transit - 248b-d8b2-4c69-b397
/* Adds layer 3 handoffs with ip transit in fabric devices based on user input.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-fabric-devices-layer3-handoffs-with-ip-transit-v1
*/
func (s *SdaService) AddFabricDevicesLayer3HandoffsWithIPTransitV1(requestSdaAddFabricDevicesLayer3HandoffsWithIpTransitV1 *RequestSdaAddFabricDevicesLayer3HandoffsWithIPTransitV1) (*ResponseSdaAddFabricDevicesLayer3HandoffsWithIPTransitV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricDevices/layer3Handoffs/ipTransits"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddFabricDevicesLayer3HandoffsWithIpTransitV1).
		SetResult(&ResponseSdaAddFabricDevicesLayer3HandoffsWithIPTransitV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddFabricDevicesLayer3HandoffsWithIPTransitV1(requestSdaAddFabricDevicesLayer3HandoffsWithIpTransitV1)
		}

		return nil, response, fmt.Errorf("error with operation AddFabricDevicesLayer3HandoffsWithIpTransitV1")
	}

	result := response.Result().(*ResponseSdaAddFabricDevicesLayer3HandoffsWithIPTransitV1)
	return result, response, err

}

//AddFabricDevicesLayer3HandoffsWithSdaTransitV1 Add fabric devices layer 3 handoffs with sda transit - 61a1-aa25-40fa-8312
/* Adds layer 3 handoffs with sda transit in fabric devices based on user input.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-fabric-devices-layer3-handoffs-with-sda-transit-v1
*/
func (s *SdaService) AddFabricDevicesLayer3HandoffsWithSdaTransitV1(requestSdaAddFabricDevicesLayer3HandoffsWithSdaTransitV1 *RequestSdaAddFabricDevicesLayer3HandoffsWithSdaTransitV1) (*ResponseSdaAddFabricDevicesLayer3HandoffsWithSdaTransitV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricDevices/layer3Handoffs/sdaTransits"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddFabricDevicesLayer3HandoffsWithSdaTransitV1).
		SetResult(&ResponseSdaAddFabricDevicesLayer3HandoffsWithSdaTransitV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddFabricDevicesLayer3HandoffsWithSdaTransitV1(requestSdaAddFabricDevicesLayer3HandoffsWithSdaTransitV1)
		}

		return nil, response, fmt.Errorf("error with operation AddFabricDevicesLayer3HandoffsWithSdaTransitV1")
	}

	result := response.Result().(*ResponseSdaAddFabricDevicesLayer3HandoffsWithSdaTransitV1)
	return result, response, err

}

//AddFabricSiteV1 Add fabric site - 83af-8a6c-4b58-8f99
/* Adds a fabric site based on user input.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-fabric-site-v1
*/
func (s *SdaService) AddFabricSiteV1(requestSdaAddFabricSiteV1 *RequestSdaAddFabricSiteV1) (*ResponseSdaAddFabricSiteV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricSites"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddFabricSiteV1).
		SetResult(&ResponseSdaAddFabricSiteV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddFabricSiteV1(requestSdaAddFabricSiteV1)
		}

		return nil, response, fmt.Errorf("error with operation AddFabricSiteV1")
	}

	result := response.Result().(*ResponseSdaAddFabricSiteV1)
	return result, response, err

}

//AddFabricZoneV1 Add fabric zone - 658c-083b-4cb9-92aa
/* Adds a fabric zone based on user input.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-fabric-zone-v1
*/
func (s *SdaService) AddFabricZoneV1(requestSdaAddFabricZoneV1 *RequestSdaAddFabricZoneV1) (*ResponseSdaAddFabricZoneV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricZones"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddFabricZoneV1).
		SetResult(&ResponseSdaAddFabricZoneV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddFabricZoneV1(requestSdaAddFabricZoneV1)
		}

		return nil, response, fmt.Errorf("error with operation AddFabricZoneV1")
	}

	result := response.Result().(*ResponseSdaAddFabricZoneV1)
	return result, response, err

}

//AddLayer2VirtualNetworksV1 Add layer 2 virtual networks - f8bd-49ed-4f5a-9aec
/* Adds layer 2 virtual networks based on user input.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-layer2-virtual-networks-v1
*/
func (s *SdaService) AddLayer2VirtualNetworksV1(requestSdaAddLayer2VirtualNetworksV1 *RequestSdaAddLayer2VirtualNetworksV1) (*ResponseSdaAddLayer2VirtualNetworksV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/layer2VirtualNetworks"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddLayer2VirtualNetworksV1).
		SetResult(&ResponseSdaAddLayer2VirtualNetworksV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddLayer2VirtualNetworksV1(requestSdaAddLayer2VirtualNetworksV1)
		}

		return nil, response, fmt.Errorf("error with operation AddLayer2VirtualNetworksV1")
	}

	result := response.Result().(*ResponseSdaAddLayer2VirtualNetworksV1)
	return result, response, err

}

//AddLayer3VirtualNetworksV1 Add layer 3 virtual networks - aba8-d8b2-482a-8b53
/* Adds layer 3 virtual networks based on user input.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-layer3-virtual-networks-v1
*/
func (s *SdaService) AddLayer3VirtualNetworksV1(requestSdaAddLayer3VirtualNetworksV1 *RequestSdaAddLayer3VirtualNetworksV1) (*ResponseSdaAddLayer3VirtualNetworksV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/layer3VirtualNetworks"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddLayer3VirtualNetworksV1).
		SetResult(&ResponseSdaAddLayer3VirtualNetworksV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddLayer3VirtualNetworksV1(requestSdaAddLayer3VirtualNetworksV1)
		}

		return nil, response, fmt.Errorf("error with operation AddLayer3VirtualNetworksV1")
	}

	result := response.Result().(*ResponseSdaAddLayer3VirtualNetworksV1)
	return result, response, err

}

//AddMulticastVirtualNetworksV1 Add multicast virtual networks - 0284-eac5-4a3b-abfd
/* Adds multicast for virtual networks based on user input.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-multicast-virtual-networks-v1
*/
func (s *SdaService) AddMulticastVirtualNetworksV1(requestSdaAddMulticastVirtualNetworksV1 *RequestSdaAddMulticastVirtualNetworksV1) (*ResponseSdaAddMulticastVirtualNetworksV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/multicast/virtualNetworks"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddMulticastVirtualNetworksV1).
		SetResult(&ResponseSdaAddMulticastVirtualNetworksV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddMulticastVirtualNetworksV1(requestSdaAddMulticastVirtualNetworksV1)
		}

		return nil, response, fmt.Errorf("error with operation AddMulticastVirtualNetworksV1")
	}

	result := response.Result().(*ResponseSdaAddMulticastVirtualNetworksV1)
	return result, response, err

}

//AddPortAssignmentsV1 Add port assignments - d8bb-0923-498a-9f63
/* Adds port assignments based on user input.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-port-assignments-v1
*/
func (s *SdaService) AddPortAssignmentsV1(requestSdaAddPortAssignmentsV1 *RequestSdaAddPortAssignmentsV1) (*ResponseSdaAddPortAssignmentsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/portAssignments"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddPortAssignmentsV1).
		SetResult(&ResponseSdaAddPortAssignmentsV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddPortAssignmentsV1(requestSdaAddPortAssignmentsV1)
		}

		return nil, response, fmt.Errorf("error with operation AddPortAssignmentsV1")
	}

	result := response.Result().(*ResponseSdaAddPortAssignmentsV1)
	return result, response, err

}

//AddPortChannelsV1 Add port channels - 2ba0-7a63-43db-843c
/* Adds port channels based on user input.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-port-channels-v1
*/
func (s *SdaService) AddPortChannelsV1(requestSdaAddPortChannelsV1 *RequestSdaAddPortChannelsV1) (*ResponseSdaAddPortChannelsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/portChannels"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddPortChannelsV1).
		SetResult(&ResponseSdaAddPortChannelsV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddPortChannelsV1(requestSdaAddPortChannelsV1)
		}

		return nil, response, fmt.Errorf("error with operation AddPortChannelsV1")
	}

	result := response.Result().(*ResponseSdaAddPortChannelsV1)
	return result, response, err

}

//ProvisionDevicesV1 Provision devices - 9c83-f9d6-4bda-b655
/* Provisions network devices to respective Sites based on user input.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!provision-devices-v1
*/
func (s *SdaService) ProvisionDevicesV1(requestSdaProvisionDevicesV1 *RequestSdaProvisionDevicesV1) (*ResponseSdaProvisionDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/provisionDevices"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaProvisionDevicesV1).
		SetResult(&ResponseSdaProvisionDevicesV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ProvisionDevicesV1(requestSdaProvisionDevicesV1)
		}

		return nil, response, fmt.Errorf("error with operation ProvisionDevicesV1")
	}

	result := response.Result().(*ResponseSdaProvisionDevicesV1)
	return result, response, err

}

//AddTransitNetworksV1 Add transit networks - bd82-db95-41fb-83b9
/* Adds transit networks based on user input.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-transit-networks-v1
*/
func (s *SdaService) AddTransitNetworksV1(requestSdaAddTransitNetworksV1 *RequestSdaAddTransitNetworksV1) (*ResponseSdaAddTransitNetworksV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/transitNetworks"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddTransitNetworksV1).
		SetResult(&ResponseSdaAddTransitNetworksV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddTransitNetworksV1(requestSdaAddTransitNetworksV1)
		}

		return nil, response, fmt.Errorf("error with operation AddTransitNetworksV1")
	}

	result := response.Result().(*ResponseSdaAddTransitNetworksV1)
	return result, response, err

}

//AddVirtualNetworkWithScalableGroupsV1 Add virtual network with scalable groups - e3a8-5b19-406a-9f4e
/* Add virtual network with scalable groups at global level



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-virtual-network-with-scalable-groups-v1
*/
func (s *SdaService) AddVirtualNetworkWithScalableGroupsV1(requestSdaAddVirtualNetworkWithScalableGroupsV1 *RequestSdaAddVirtualNetworkWithScalableGroupsV1) (*ResponseSdaAddVirtualNetworkWithScalableGroupsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/virtual-network"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddVirtualNetworkWithScalableGroupsV1).
		SetResult(&ResponseSdaAddVirtualNetworkWithScalableGroupsV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddVirtualNetworkWithScalableGroupsV1(requestSdaAddVirtualNetworkWithScalableGroupsV1)
		}

		return nil, response, fmt.Errorf("error with operation AddVirtualNetworkWithScalableGroupsV1")
	}

	result := response.Result().(*ResponseSdaAddVirtualNetworkWithScalableGroupsV1)
	return result, response, err

}

//UpdateDefaultAuthenticationProfileInSdaFabricV1 Update default authentication profile in SDA Fabric - 8984-ea77-44d9-8a54
/* Update default authentication profile in SDA Fabric


 */
func (s *SdaService) UpdateDefaultAuthenticationProfileInSdaFabricV1(requestSdaUpdateDefaultAuthenticationProfileInSDAFabricV1 *RequestSdaUpdateDefaultAuthenticationProfileInSdaFabricV1) (*ResponseSdaUpdateDefaultAuthenticationProfileInSdaFabricV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/authentication-profile"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaUpdateDefaultAuthenticationProfileInSDAFabricV1).
		SetResult(&ResponseSdaUpdateDefaultAuthenticationProfileInSdaFabricV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateDefaultAuthenticationProfileInSdaFabricV1(requestSdaUpdateDefaultAuthenticationProfileInSDAFabricV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateDefaultAuthenticationProfileInSdaFabricV1")
	}

	result := response.Result().(*ResponseSdaUpdateDefaultAuthenticationProfileInSdaFabricV1)
	return result, response, err

}

//ReProvisionWiredDeviceV1 Re-Provision Wired Device - 4e95-c9a2-41ab-8889
/* Re-Provision Wired Device


 */
func (s *SdaService) ReProvisionWiredDeviceV1(requestSdaReProvisionWiredDeviceV1 *RequestSdaReProvisionWiredDeviceV1) (*ResponseSdaReProvisionWiredDeviceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/provision-device"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaReProvisionWiredDeviceV1).
		SetResult(&ResponseSdaReProvisionWiredDeviceV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReProvisionWiredDeviceV1(requestSdaReProvisionWiredDeviceV1)
		}
		return nil, response, fmt.Errorf("error with operation ReProvisionWiredDeviceV1")
	}

	result := response.Result().(*ResponseSdaReProvisionWiredDeviceV1)
	return result, response, err

}

//UpdateAnycastGatewaysV1 Update anycast gateways - 0dde-2905-a7b5-4e8d
/* Updates anycast gateways based on user input.


 */
func (s *SdaService) UpdateAnycastGatewaysV1(requestSdaUpdateAnycastGatewaysV1 *RequestSdaUpdateAnycastGatewaysV1) (*ResponseSdaUpdateAnycastGatewaysV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/anycastGateways"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaUpdateAnycastGatewaysV1).
		SetResult(&ResponseSdaUpdateAnycastGatewaysV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateAnycastGatewaysV1(requestSdaUpdateAnycastGatewaysV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateAnycastGatewaysV1")
	}

	result := response.Result().(*ResponseSdaUpdateAnycastGatewaysV1)
	return result, response, err

}

//UpdateAuthenticationProfileV1 Update authentication profile - 8296-3890-4978-bda1
/* Updates an authentication profile based on user input.


 */
func (s *SdaService) UpdateAuthenticationProfileV1(requestSdaUpdateAuthenticationProfileV1 *RequestSdaUpdateAuthenticationProfileV1) (*ResponseSdaUpdateAuthenticationProfileV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/authenticationProfiles"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaUpdateAuthenticationProfileV1).
		SetResult(&ResponseSdaUpdateAuthenticationProfileV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateAuthenticationProfileV1(requestSdaUpdateAuthenticationProfileV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateAuthenticationProfileV1")
	}

	result := response.Result().(*ResponseSdaUpdateAuthenticationProfileV1)
	return result, response, err

}

//UpdateExtranetPolicyV1 Update extranet policy - 899d-9aab-4b0a-8d5a
/* Updates an extranet policy based on user input.


 */
func (s *SdaService) UpdateExtranetPolicyV1(requestSdaUpdateExtranetPolicyV1 *RequestSdaUpdateExtranetPolicyV1) (*ResponseSdaUpdateExtranetPolicyV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/extranetPolicies"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaUpdateExtranetPolicyV1).
		SetResult(&ResponseSdaUpdateExtranetPolicyV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateExtranetPolicyV1(requestSdaUpdateExtranetPolicyV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateExtranetPolicyV1")
	}

	result := response.Result().(*ResponseSdaUpdateExtranetPolicyV1)
	return result, response, err

}

//UpdateFabricDevicesV1 Update fabric devices - ceb9-2a9a-409b-8066
/* Updates fabric devices based on user input.


 */
func (s *SdaService) UpdateFabricDevicesV1(requestSdaUpdateFabricDevicesV1 *RequestSdaUpdateFabricDevicesV1) (*ResponseSdaUpdateFabricDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricDevices"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaUpdateFabricDevicesV1).
		SetResult(&ResponseSdaUpdateFabricDevicesV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateFabricDevicesV1(requestSdaUpdateFabricDevicesV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateFabricDevicesV1")
	}

	result := response.Result().(*ResponseSdaUpdateFabricDevicesV1)
	return result, response, err

}

//UpdateFabricDevicesLayer3HandoffsWithIPTransitV1 Update fabric devices layer 3 handoffs with ip transit - ff8b-cba9-4829-8ca6
/* Updates layer 3 handoffs with ip transit of fabric devices based on user input.


 */
func (s *SdaService) UpdateFabricDevicesLayer3HandoffsWithIPTransitV1(requestSdaUpdateFabricDevicesLayer3HandoffsWithIpTransitV1 *RequestSdaUpdateFabricDevicesLayer3HandoffsWithIPTransitV1) (*ResponseSdaUpdateFabricDevicesLayer3HandoffsWithIPTransitV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricDevices/layer3Handoffs/ipTransits"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaUpdateFabricDevicesLayer3HandoffsWithIpTransitV1).
		SetResult(&ResponseSdaUpdateFabricDevicesLayer3HandoffsWithIPTransitV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateFabricDevicesLayer3HandoffsWithIPTransitV1(requestSdaUpdateFabricDevicesLayer3HandoffsWithIpTransitV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateFabricDevicesLayer3HandoffsWithIpTransitV1")
	}

	result := response.Result().(*ResponseSdaUpdateFabricDevicesLayer3HandoffsWithIPTransitV1)
	return result, response, err

}

//UpdateFabricDevicesLayer3HandoffsWithSdaTransitV1 Update fabric devices layer 3 handoffs with sda transit - b6a1-3a87-435a-a5ed
/* Updates layer 3 handoffs with sda transit of fabric devices based on user input.


 */
func (s *SdaService) UpdateFabricDevicesLayer3HandoffsWithSdaTransitV1(requestSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransitV1 *RequestSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransitV1) (*ResponseSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransitV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricDevices/layer3Handoffs/sdaTransits"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransitV1).
		SetResult(&ResponseSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransitV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateFabricDevicesLayer3HandoffsWithSdaTransitV1(requestSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransitV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateFabricDevicesLayer3HandoffsWithSdaTransitV1")
	}

	result := response.Result().(*ResponseSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransitV1)
	return result, response, err

}

//UpdateFabricSiteV1 Update fabric site - b683-b984-430b-8c74
/* Updates a fabric site based on user input.


 */
func (s *SdaService) UpdateFabricSiteV1(requestSdaUpdateFabricSiteV1 *RequestSdaUpdateFabricSiteV1) (*ResponseSdaUpdateFabricSiteV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricSites"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaUpdateFabricSiteV1).
		SetResult(&ResponseSdaUpdateFabricSiteV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateFabricSiteV1(requestSdaUpdateFabricSiteV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateFabricSiteV1")
	}

	result := response.Result().(*ResponseSdaUpdateFabricSiteV1)
	return result, response, err

}

//UpdateFabricZoneV1 Update fabric zone - f39c-1a54-4c2a-a790
/* Updates a fabric zone based on user input.


 */
func (s *SdaService) UpdateFabricZoneV1(requestSdaUpdateFabricZoneV1 *RequestSdaUpdateFabricZoneV1) (*ResponseSdaUpdateFabricZoneV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricZones"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaUpdateFabricZoneV1).
		SetResult(&ResponseSdaUpdateFabricZoneV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateFabricZoneV1(requestSdaUpdateFabricZoneV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateFabricZoneV1")
	}

	result := response.Result().(*ResponseSdaUpdateFabricZoneV1)
	return result, response, err

}

//UpdateLayer2VirtualNetworksV1 Update layer 2 virtual networks - 18ba-d85c-4dcb-bc31
/* Updates layer 2 virtual networks based on user input.


 */
func (s *SdaService) UpdateLayer2VirtualNetworksV1(requestSdaUpdateLayer2VirtualNetworksV1 *RequestSdaUpdateLayer2VirtualNetworksV1) (*ResponseSdaUpdateLayer2VirtualNetworksV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/layer2VirtualNetworks"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaUpdateLayer2VirtualNetworksV1).
		SetResult(&ResponseSdaUpdateLayer2VirtualNetworksV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateLayer2VirtualNetworksV1(requestSdaUpdateLayer2VirtualNetworksV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateLayer2VirtualNetworksV1")
	}

	result := response.Result().(*ResponseSdaUpdateLayer2VirtualNetworksV1)
	return result, response, err

}

//UpdateLayer3VirtualNetworksV1 Update layer 3 virtual networks - c995-fbe5-465b-be45
/* Updates layer 3 virtual networks based on user input.


 */
func (s *SdaService) UpdateLayer3VirtualNetworksV1(requestSdaUpdateLayer3VirtualNetworksV1 *RequestSdaUpdateLayer3VirtualNetworksV1) (*ResponseSdaUpdateLayer3VirtualNetworksV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/layer3VirtualNetworks"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaUpdateLayer3VirtualNetworksV1).
		SetResult(&ResponseSdaUpdateLayer3VirtualNetworksV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateLayer3VirtualNetworksV1(requestSdaUpdateLayer3VirtualNetworksV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateLayer3VirtualNetworksV1")
	}

	result := response.Result().(*ResponseSdaUpdateLayer3VirtualNetworksV1)
	return result, response, err

}

//UpdateMulticastV1 Update multicast - 9cb0-68b9-4e6b-8c0c
/* Updates a multicast configuration at a fabric level based on user input.


 */
func (s *SdaService) UpdateMulticastV1(requestSdaUpdateMulticastV1 *RequestSdaUpdateMulticastV1) (*ResponseSdaUpdateMulticastV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/multicast"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaUpdateMulticastV1).
		SetResult(&ResponseSdaUpdateMulticastV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateMulticastV1(requestSdaUpdateMulticastV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateMulticastV1")
	}

	result := response.Result().(*ResponseSdaUpdateMulticastV1)
	return result, response, err

}

//UpdateMulticastVirtualNetworksV1 Update multicast virtual networks - 6c8a-cbbf-4a19-b48b
/* Updates multicast configurations for virtual networks based on user input.


 */
func (s *SdaService) UpdateMulticastVirtualNetworksV1(requestSdaUpdateMulticastVirtualNetworksV1 *RequestSdaUpdateMulticastVirtualNetworksV1) (*ResponseSdaUpdateMulticastVirtualNetworksV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/multicast/virtualNetworks"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaUpdateMulticastVirtualNetworksV1).
		SetResult(&ResponseSdaUpdateMulticastVirtualNetworksV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateMulticastVirtualNetworksV1(requestSdaUpdateMulticastVirtualNetworksV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateMulticastVirtualNetworksV1")
	}

	result := response.Result().(*ResponseSdaUpdateMulticastVirtualNetworksV1)
	return result, response, err

}

//UpdatePortAssignmentsV1 Update port assignments - eab7-6a0e-469a-b21a
/* Updates port assignments based on user input.


 */
func (s *SdaService) UpdatePortAssignmentsV1(requestSdaUpdatePortAssignmentsV1 *RequestSdaUpdatePortAssignmentsV1) (*ResponseSdaUpdatePortAssignmentsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/portAssignments"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaUpdatePortAssignmentsV1).
		SetResult(&ResponseSdaUpdatePortAssignmentsV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdatePortAssignmentsV1(requestSdaUpdatePortAssignmentsV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdatePortAssignmentsV1")
	}

	result := response.Result().(*ResponseSdaUpdatePortAssignmentsV1)
	return result, response, err

}

//UpdatePortChannelsV1 Update port channels - fc9d-7a51-472b-ba5e
/* Updates port channels based on user input.


 */
func (s *SdaService) UpdatePortChannelsV1(requestSdaUpdatePortChannelsV1 *RequestSdaUpdatePortChannelsV1) (*ResponseSdaUpdatePortChannelsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/portChannels"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaUpdatePortChannelsV1).
		SetResult(&ResponseSdaUpdatePortChannelsV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdatePortChannelsV1(requestSdaUpdatePortChannelsV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdatePortChannelsV1")
	}

	result := response.Result().(*ResponseSdaUpdatePortChannelsV1)
	return result, response, err

}

//ReProvisionDevicesV1 Re-provision devices - 898a-c9a2-4ee8-8e7f
/* Re-provisions network devices to the site based on the user input.


 */
func (s *SdaService) ReProvisionDevicesV1(requestSdaReProvisionDevicesV1 *RequestSdaReProvisionDevicesV1) (*ResponseSdaReProvisionDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/provisionDevices"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaReProvisionDevicesV1).
		SetResult(&ResponseSdaReProvisionDevicesV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReProvisionDevicesV1(requestSdaReProvisionDevicesV1)
		}
		return nil, response, fmt.Errorf("error with operation ReProvisionDevicesV1")
	}

	result := response.Result().(*ResponseSdaReProvisionDevicesV1)
	return result, response, err

}

//UpdateTransitNetworksV1 Update transit networks - 6fa2-3824-49f8-a0d7
/* Updates transit networks based on user input.


 */
func (s *SdaService) UpdateTransitNetworksV1(requestSdaUpdateTransitNetworksV1 *RequestSdaUpdateTransitNetworksV1) (*ResponseSdaUpdateTransitNetworksV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/transitNetworks"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaUpdateTransitNetworksV1).
		SetResult(&ResponseSdaUpdateTransitNetworksV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateTransitNetworksV1(requestSdaUpdateTransitNetworksV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateTransitNetworksV1")
	}

	result := response.Result().(*ResponseSdaUpdateTransitNetworksV1)
	return result, response, err

}

//UpdateVirtualNetworkWithScalableGroupsV1 Update virtual network with scalable groups - c48b-2904-49bb-875f
/* Update virtual network with scalable groups


 */
func (s *SdaService) UpdateVirtualNetworkWithScalableGroupsV1(requestSdaUpdateVirtualNetworkWithScalableGroupsV1 *RequestSdaUpdateVirtualNetworkWithScalableGroupsV1) (*ResponseSdaUpdateVirtualNetworkWithScalableGroupsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/virtual-network"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaUpdateVirtualNetworkWithScalableGroupsV1).
		SetResult(&ResponseSdaUpdateVirtualNetworkWithScalableGroupsV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateVirtualNetworkWithScalableGroupsV1(requestSdaUpdateVirtualNetworkWithScalableGroupsV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateVirtualNetworkWithScalableGroupsV1")
	}

	result := response.Result().(*ResponseSdaUpdateVirtualNetworkWithScalableGroupsV1)
	return result, response, err

}

//DeleteDefaultAuthenticationProfileFromSdaFabricV1 Delete default authentication profile from SDA Fabric - 3ebc-da3e-4acb-afb7
/* Delete default authentication profile in SDA Fabric


@param DeleteDefaultAuthenticationProfileFromSDAFabricV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-default-authentication-profile-from-sda-fabric-v1
*/
func (s *SdaService) DeleteDefaultAuthenticationProfileFromSdaFabricV1(DeleteDefaultAuthenticationProfileFromSDAFabricV1QueryParams *DeleteDefaultAuthenticationProfileFromSdaFabricV1QueryParams) (*ResponseSdaDeleteDefaultAuthenticationProfileFromSdaFabricV1, *resty.Response, error) {
	//DeleteDefaultAuthenticationProfileFromSDAFabricV1QueryParams *DeleteDefaultAuthenticationProfileFromSdaFabricV1QueryParams
	path := "/dna/intent/api/v1/business/sda/authentication-profile"

	queryString, _ := query.Values(DeleteDefaultAuthenticationProfileFromSDAFabricV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteDefaultAuthenticationProfileFromSdaFabricV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteDefaultAuthenticationProfileFromSdaFabricV1(DeleteDefaultAuthenticationProfileFromSDAFabricV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteDefaultAuthenticationProfileFromSdaFabricV1")
	}

	result := response.Result().(*ResponseSdaDeleteDefaultAuthenticationProfileFromSdaFabricV1)
	return result, response, err

}

//DeleteBorderDeviceFromSdaFabricV1 Delete border device from SDA Fabric - cb81-b935-40ba-aab0
/* Delete border device from SDA Fabric


@param DeleteBorderDeviceFromSDAFabricV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-border-device-from-sda-fabric-v1
*/
func (s *SdaService) DeleteBorderDeviceFromSdaFabricV1(DeleteBorderDeviceFromSDAFabricV1QueryParams *DeleteBorderDeviceFromSdaFabricV1QueryParams) (*ResponseSdaDeleteBorderDeviceFromSdaFabricV1, *resty.Response, error) {
	//DeleteBorderDeviceFromSDAFabricV1QueryParams *DeleteBorderDeviceFromSdaFabricV1QueryParams
	path := "/dna/intent/api/v1/business/sda/border-device"

	queryString, _ := query.Values(DeleteBorderDeviceFromSDAFabricV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteBorderDeviceFromSdaFabricV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteBorderDeviceFromSdaFabricV1(DeleteBorderDeviceFromSDAFabricV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteBorderDeviceFromSdaFabricV1")
	}

	result := response.Result().(*ResponseSdaDeleteBorderDeviceFromSdaFabricV1)
	return result, response, err

}

//DeleteControlPlaneDeviceInSdaFabricV1 Delete control plane device in SDA Fabric - f6bd-6bf6-4e68-90be
/* Delete control plane device in SDA Fabric


@param DeleteControlPlaneDeviceInSDAFabricV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-control-plane-device-in-sda-fabric-v1
*/
func (s *SdaService) DeleteControlPlaneDeviceInSdaFabricV1(DeleteControlPlaneDeviceInSDAFabricV1QueryParams *DeleteControlPlaneDeviceInSdaFabricV1QueryParams) (*ResponseSdaDeleteControlPlaneDeviceInSdaFabricV1, *resty.Response, error) {
	//DeleteControlPlaneDeviceInSDAFabricV1QueryParams *DeleteControlPlaneDeviceInSdaFabricV1QueryParams
	path := "/dna/intent/api/v1/business/sda/control-plane-device"

	queryString, _ := query.Values(DeleteControlPlaneDeviceInSDAFabricV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteControlPlaneDeviceInSdaFabricV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteControlPlaneDeviceInSdaFabricV1(DeleteControlPlaneDeviceInSDAFabricV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteControlPlaneDeviceInSdaFabricV1")
	}

	result := response.Result().(*ResponseSdaDeleteControlPlaneDeviceInSdaFabricV1)
	return result, response, err

}

//DeleteEdgeDeviceFromSdaFabricV1 Delete edge device from SDA Fabric - 1fb8-f9f2-4c99-8133
/* Delete edge device from SDA Fabric.


@param DeleteEdgeDeviceFromSDAFabricV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-edge-device-from-sda-fabric-v1
*/
func (s *SdaService) DeleteEdgeDeviceFromSdaFabricV1(DeleteEdgeDeviceFromSDAFabricV1QueryParams *DeleteEdgeDeviceFromSdaFabricV1QueryParams) (*ResponseSdaDeleteEdgeDeviceFromSdaFabricV1, *resty.Response, error) {
	//DeleteEdgeDeviceFromSDAFabricV1QueryParams *DeleteEdgeDeviceFromSdaFabricV1QueryParams
	path := "/dna/intent/api/v1/business/sda/edge-device"

	queryString, _ := query.Values(DeleteEdgeDeviceFromSDAFabricV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteEdgeDeviceFromSdaFabricV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteEdgeDeviceFromSdaFabricV1(DeleteEdgeDeviceFromSDAFabricV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteEdgeDeviceFromSdaFabricV1")
	}

	result := response.Result().(*ResponseSdaDeleteEdgeDeviceFromSdaFabricV1)
	return result, response, err

}

//DeleteSiteFromSdaFabricV1 Delete Site from SDA Fabric - 5086-4acf-4ad8-b54d
/* Delete Site from SDA Fabric


@param DeleteSiteFromSDAFabricV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-site-from-sda-fabric-v1
*/
func (s *SdaService) DeleteSiteFromSdaFabricV1(DeleteSiteFromSDAFabricV1QueryParams *DeleteSiteFromSdaFabricV1QueryParams) (*ResponseSdaDeleteSiteFromSdaFabricV1, *resty.Response, error) {
	//DeleteSiteFromSDAFabricV1QueryParams *DeleteSiteFromSdaFabricV1QueryParams
	path := "/dna/intent/api/v1/business/sda/fabric-site"

	queryString, _ := query.Values(DeleteSiteFromSDAFabricV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteSiteFromSdaFabricV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteSiteFromSdaFabricV1(DeleteSiteFromSDAFabricV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteSiteFromSdaFabricV1")
	}

	result := response.Result().(*ResponseSdaDeleteSiteFromSdaFabricV1)
	return result, response, err

}

//DeletePortAssignmentForAccessPointInSdaFabricV1 Delete Port assignment for access point in SDA Fabric - 0787-4a4c-4c9a-abd9
/* Delete Port assignment for access point in SDA Fabric


@param DeletePortAssignmentForAccessPointInSDAFabricV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-port-assignment-for-access-point-in-sda-fabric-v1
*/
func (s *SdaService) DeletePortAssignmentForAccessPointInSdaFabricV1(DeletePortAssignmentForAccessPointInSDAFabricV1QueryParams *DeletePortAssignmentForAccessPointInSdaFabricV1QueryParams) (*ResponseSdaDeletePortAssignmentForAccessPointInSdaFabricV1, *resty.Response, error) {
	//DeletePortAssignmentForAccessPointInSDAFabricV1QueryParams *DeletePortAssignmentForAccessPointInSdaFabricV1QueryParams
	path := "/dna/intent/api/v1/business/sda/hostonboarding/access-point"

	queryString, _ := query.Values(DeletePortAssignmentForAccessPointInSDAFabricV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeletePortAssignmentForAccessPointInSdaFabricV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletePortAssignmentForAccessPointInSdaFabricV1(DeletePortAssignmentForAccessPointInSDAFabricV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeletePortAssignmentForAccessPointInSdaFabricV1")
	}

	result := response.Result().(*ResponseSdaDeletePortAssignmentForAccessPointInSdaFabricV1)
	return result, response, err

}

//DeletePortAssignmentForUserDeviceInSdaFabricV1 Delete Port assignment for user device in SDA Fabric - cba5-b8b1-4edb-81f4
/* Delete Port assignment for user device in SDA Fabric.


@param DeletePortAssignmentForUserDeviceInSDAFabricV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-port-assignment-for-user-device-in-sda-fabric-v1
*/
func (s *SdaService) DeletePortAssignmentForUserDeviceInSdaFabricV1(DeletePortAssignmentForUserDeviceInSDAFabricV1QueryParams *DeletePortAssignmentForUserDeviceInSdaFabricV1QueryParams) (*ResponseSdaDeletePortAssignmentForUserDeviceInSdaFabricV1, *resty.Response, error) {
	//DeletePortAssignmentForUserDeviceInSDAFabricV1QueryParams *DeletePortAssignmentForUserDeviceInSdaFabricV1QueryParams
	path := "/dna/intent/api/v1/business/sda/hostonboarding/user-device"

	queryString, _ := query.Values(DeletePortAssignmentForUserDeviceInSDAFabricV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeletePortAssignmentForUserDeviceInSdaFabricV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletePortAssignmentForUserDeviceInSdaFabricV1(DeletePortAssignmentForUserDeviceInSDAFabricV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeletePortAssignmentForUserDeviceInSdaFabricV1")
	}

	result := response.Result().(*ResponseSdaDeletePortAssignmentForUserDeviceInSdaFabricV1)
	return result, response, err

}

//DeleteMulticastFromSdaFabricV1 Delete multicast from SDA fabric - 2bb0-0be5-45cb-bc99
/* Delete multicast from SDA fabric


@param DeleteMulticastFromSDAFabricV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-multicast-from-sda-fabric-v1
*/
func (s *SdaService) DeleteMulticastFromSdaFabricV1(DeleteMulticastFromSDAFabricV1QueryParams *DeleteMulticastFromSdaFabricV1QueryParams) (*ResponseSdaDeleteMulticastFromSdaFabricV1, *resty.Response, error) {
	//DeleteMulticastFromSDAFabricV1QueryParams *DeleteMulticastFromSdaFabricV1QueryParams
	path := "/dna/intent/api/v1/business/sda/multicast"

	queryString, _ := query.Values(DeleteMulticastFromSDAFabricV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteMulticastFromSdaFabricV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteMulticastFromSdaFabricV1(DeleteMulticastFromSDAFabricV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteMulticastFromSdaFabricV1")
	}

	result := response.Result().(*ResponseSdaDeleteMulticastFromSdaFabricV1)
	return result, response, err

}

//DeleteProvisionedWiredDeviceV1 Delete provisioned Wired Device - e495-b94e-463b-ae04
/* Delete provisioned Wired Device


@param DeleteProvisionedWiredDeviceV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-provisioned-wired-device-v1
*/
func (s *SdaService) DeleteProvisionedWiredDeviceV1(DeleteProvisionedWiredDeviceV1QueryParams *DeleteProvisionedWiredDeviceV1QueryParams) (*ResponseSdaDeleteProvisionedWiredDeviceV1, *resty.Response, error) {
	//DeleteProvisionedWiredDeviceV1QueryParams *DeleteProvisionedWiredDeviceV1QueryParams
	path := "/dna/intent/api/v1/business/sda/provision-device"

	queryString, _ := query.Values(DeleteProvisionedWiredDeviceV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteProvisionedWiredDeviceV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteProvisionedWiredDeviceV1(DeleteProvisionedWiredDeviceV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteProvisionedWiredDeviceV1")
	}

	result := response.Result().(*ResponseSdaDeleteProvisionedWiredDeviceV1)
	return result, response, err

}

//DeleteTransitPeerNetworkV1 Delete Transit Peer Network - d0aa-fa69-4f4b-9d7b
/* Delete Transit Peer Network from SD-Access


@param DeleteTransitPeerNetworkV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-transit-peer-network-v1
*/
func (s *SdaService) DeleteTransitPeerNetworkV1(DeleteTransitPeerNetworkV1QueryParams *DeleteTransitPeerNetworkV1QueryParams) (*ResponseSdaDeleteTransitPeerNetworkV1, *resty.Response, error) {
	//DeleteTransitPeerNetworkV1QueryParams *DeleteTransitPeerNetworkV1QueryParams
	path := "/dna/intent/api/v1/business/sda/transit-peer-network"

	queryString, _ := query.Values(DeleteTransitPeerNetworkV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteTransitPeerNetworkV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteTransitPeerNetworkV1(DeleteTransitPeerNetworkV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteTransitPeerNetworkV1")
	}

	result := response.Result().(*ResponseSdaDeleteTransitPeerNetworkV1)
	return result, response, err

}

//DeleteVnFromSdaFabricV1 Delete VN from SDA Fabric - c78c-9ad2-45bb-9657
/* Delete virtual network (VN) from SDA Fabric


@param DeleteVNFromSDAFabricV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-vn-from-sda-fabric-v1
*/
func (s *SdaService) DeleteVnFromSdaFabricV1(DeleteVNFromSDAFabricV1QueryParams *DeleteVnFromSdaFabricV1QueryParams) (*ResponseSdaDeleteVnFromSdaFabricV1, *resty.Response, error) {
	//DeleteVNFromSDAFabricV1QueryParams *DeleteVnFromSdaFabricV1QueryParams
	path := "/dna/intent/api/v1/business/sda/virtual-network"

	queryString, _ := query.Values(DeleteVNFromSDAFabricV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteVnFromSdaFabricV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteVnFromSdaFabricV1(DeleteVNFromSDAFabricV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteVnFromSdaFabricV1")
	}

	result := response.Result().(*ResponseSdaDeleteVnFromSdaFabricV1)
	return result, response, err

}

//DeleteIPPoolFromSdaVirtualNetworkV1 Delete IP Pool from SDA Virtual Network - 549e-4aff-42bb-b52a
/* Delete IP Pool from SDA Virtual Network


@param DeleteIPPoolFromSDAVirtualNetworkV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-ip-pool-from-sda-virtual-network-v1
*/
func (s *SdaService) DeleteIPPoolFromSdaVirtualNetworkV1(DeleteIPPoolFromSDAVirtualNetworkV1QueryParams *DeleteIPPoolFromSdaVirtualNetworkV1QueryParams) (*ResponseSdaDeleteIPPoolFromSdaVirtualNetworkV1, *resty.Response, error) {
	//DeleteIPPoolFromSDAVirtualNetworkV1QueryParams *DeleteIPPoolFromSdaVirtualNetworkV1QueryParams
	path := "/dna/intent/api/v1/business/sda/virtualnetwork/ippool"

	queryString, _ := query.Values(DeleteIPPoolFromSDAVirtualNetworkV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteIPPoolFromSdaVirtualNetworkV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteIPPoolFromSdaVirtualNetworkV1(DeleteIPPoolFromSDAVirtualNetworkV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteIpPoolFromSdaVirtualNetworkV1")
	}

	result := response.Result().(*ResponseSdaDeleteIPPoolFromSdaVirtualNetworkV1)
	return result, response, err

}

//DeleteAnycastGatewayByIDV1 Delete anycast gateway by id - 4bfa-d25a-ce07-99f3
/* Deletes an anycast gateway based on id.


@param id id path parameter. ID of the anycast gateway.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-anycast-gateway-by-id-v1
*/
func (s *SdaService) DeleteAnycastGatewayByIDV1(id string) (*ResponseSdaDeleteAnycastGatewayByIDV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/sda/anycastGateways/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSdaDeleteAnycastGatewayByIDV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteAnycastGatewayByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteAnycastGatewayByIdV1")
	}

	result := response.Result().(*ResponseSdaDeleteAnycastGatewayByIDV1)
	return result, response, err

}

//DeleteExtranetPoliciesV1 Delete extranet policies - 908a-8bbf-4aeb-9382
/* Deletes extranet policies based on user input.


@param DeleteExtranetPoliciesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-extranet-policies-v1
*/
func (s *SdaService) DeleteExtranetPoliciesV1(DeleteExtranetPoliciesV1QueryParams *DeleteExtranetPoliciesV1QueryParams) (*ResponseSdaDeleteExtranetPoliciesV1, *resty.Response, error) {
	//DeleteExtranetPoliciesV1QueryParams *DeleteExtranetPoliciesV1QueryParams
	path := "/dna/intent/api/v1/sda/extranetPolicies"

	queryString, _ := query.Values(DeleteExtranetPoliciesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteExtranetPoliciesV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteExtranetPoliciesV1(DeleteExtranetPoliciesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteExtranetPoliciesV1")
	}

	result := response.Result().(*ResponseSdaDeleteExtranetPoliciesV1)
	return result, response, err

}

//DeleteExtranetPolicyByIDV1 Delete extranet policy by id - 45a7-eb82-446a-b812
/* Deletes an extranet policy based on id.


@param id id path parameter. ID of the extranet policy.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-extranet-policy-by-id-v1
*/
func (s *SdaService) DeleteExtranetPolicyByIDV1(id string) (*ResponseSdaDeleteExtranetPolicyByIDV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/sda/extranetPolicies/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSdaDeleteExtranetPolicyByIDV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteExtranetPolicyByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteExtranetPolicyByIdV1")
	}

	result := response.Result().(*ResponseSdaDeleteExtranetPolicyByIDV1)
	return result, response, err

}

//DeleteFabricDevicesV1 Delete fabric devices - 8db3-88ed-4018-810a
/* Deletes fabric devices based on user input.


@param DeleteFabricDevicesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-fabric-devices-v1
*/
func (s *SdaService) DeleteFabricDevicesV1(DeleteFabricDevicesV1QueryParams *DeleteFabricDevicesV1QueryParams) (*ResponseSdaDeleteFabricDevicesV1, *resty.Response, error) {
	//DeleteFabricDevicesV1QueryParams *DeleteFabricDevicesV1QueryParams
	path := "/dna/intent/api/v1/sda/fabricDevices"

	queryString, _ := query.Values(DeleteFabricDevicesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteFabricDevicesV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteFabricDevicesV1(DeleteFabricDevicesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteFabricDevicesV1")
	}

	result := response.Result().(*ResponseSdaDeleteFabricDevicesV1)
	return result, response, err

}

//DeleteFabricDeviceLayer2HandoffsV1 Delete fabric device layer 2 handoffs - aea7-8b07-48a9-955d
/* Deletes layer 2 handoffs of a fabric device based on user input.


@param DeleteFabricDeviceLayer2HandoffsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-fabric-device-layer2-handoffs-v1
*/
func (s *SdaService) DeleteFabricDeviceLayer2HandoffsV1(DeleteFabricDeviceLayer2HandoffsV1QueryParams *DeleteFabricDeviceLayer2HandoffsV1QueryParams) (*ResponseSdaDeleteFabricDeviceLayer2HandoffsV1, *resty.Response, error) {
	//DeleteFabricDeviceLayer2HandoffsV1QueryParams *DeleteFabricDeviceLayer2HandoffsV1QueryParams
	path := "/dna/intent/api/v1/sda/fabricDevices/layer2Handoffs"

	queryString, _ := query.Values(DeleteFabricDeviceLayer2HandoffsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteFabricDeviceLayer2HandoffsV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteFabricDeviceLayer2HandoffsV1(DeleteFabricDeviceLayer2HandoffsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteFabricDeviceLayer2HandoffsV1")
	}

	result := response.Result().(*ResponseSdaDeleteFabricDeviceLayer2HandoffsV1)
	return result, response, err

}

//DeleteFabricDeviceLayer2HandoffByIDV1 Delete fabric device layer 2 handoff by id - 61ab-38c3-4948-b928
/* Deletes a layer 2 handoff of a fabric device based on id.


@param id id path parameter. ID of the layer 2 handoff of a fabric device.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-fabric-device-layer2-handoff-by-id-v1
*/
func (s *SdaService) DeleteFabricDeviceLayer2HandoffByIDV1(id string) (*ResponseSdaDeleteFabricDeviceLayer2HandoffByIDV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/sda/fabricDevices/layer2Handoffs/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSdaDeleteFabricDeviceLayer2HandoffByIDV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteFabricDeviceLayer2HandoffByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteFabricDeviceLayer2HandoffByIdV1")
	}

	result := response.Result().(*ResponseSdaDeleteFabricDeviceLayer2HandoffByIDV1)
	return result, response, err

}

//DeleteFabricDeviceLayer3HandoffsWithIPTransitV1 Delete fabric device layer 3 handoffs with ip transit - 8a87-9bad-45db-8f8f
/* Deletes layer 3 handoffs with ip transit of a fabric device based on user input.


@param DeleteFabricDeviceLayer3HandoffsWithIpTransitV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-fabric-device-layer3-handoffs-with-ip-transit-v1
*/
func (s *SdaService) DeleteFabricDeviceLayer3HandoffsWithIPTransitV1(DeleteFabricDeviceLayer3HandoffsWithIpTransitV1QueryParams *DeleteFabricDeviceLayer3HandoffsWithIPTransitV1QueryParams) (*ResponseSdaDeleteFabricDeviceLayer3HandoffsWithIPTransitV1, *resty.Response, error) {
	//DeleteFabricDeviceLayer3HandoffsWithIpTransitV1QueryParams *DeleteFabricDeviceLayer3HandoffsWithIPTransitV1QueryParams
	path := "/dna/intent/api/v1/sda/fabricDevices/layer3Handoffs/ipTransits"

	queryString, _ := query.Values(DeleteFabricDeviceLayer3HandoffsWithIpTransitV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteFabricDeviceLayer3HandoffsWithIPTransitV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteFabricDeviceLayer3HandoffsWithIPTransitV1(DeleteFabricDeviceLayer3HandoffsWithIpTransitV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteFabricDeviceLayer3HandoffsWithIpTransitV1")
	}

	result := response.Result().(*ResponseSdaDeleteFabricDeviceLayer3HandoffsWithIPTransitV1)
	return result, response, err

}

//DeleteFabricDeviceLayer3HandoffWithIPTransitByIDV1 Delete fabric device layer 3 handoff with ip transit by id - d396-7b01-4d1b-845b
/* Deletes a layer 3 handoff with ip transit of a fabric device by id.


@param id id path parameter. ID of the layer 3 handoff with ip transit of a fabric device.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-fabric-device-layer3-handoff-with-ip-transit-by-id-v1
*/
func (s *SdaService) DeleteFabricDeviceLayer3HandoffWithIPTransitByIDV1(id string) (*ResponseSdaDeleteFabricDeviceLayer3HandoffWithIPTransitByIDV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/sda/fabricDevices/layer3Handoffs/ipTransits/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSdaDeleteFabricDeviceLayer3HandoffWithIPTransitByIDV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteFabricDeviceLayer3HandoffWithIPTransitByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteFabricDeviceLayer3HandoffWithIpTransitByIdV1")
	}

	result := response.Result().(*ResponseSdaDeleteFabricDeviceLayer3HandoffWithIPTransitByIDV1)
	return result, response, err

}

//DeleteFabricDeviceLayer3HandoffsWithSdaTransitV1 Delete fabric device layer 3 handoffs with sda transit - 8d8c-b8b6-4e9a-a432
/* Deletes layer 3 handoffs with sda transit of a fabric device based on user input.


@param DeleteFabricDeviceLayer3HandoffsWithSdaTransitV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-fabric-device-layer3-handoffs-with-sda-transit-v1
*/
func (s *SdaService) DeleteFabricDeviceLayer3HandoffsWithSdaTransitV1(DeleteFabricDeviceLayer3HandoffsWithSdaTransitV1QueryParams *DeleteFabricDeviceLayer3HandoffsWithSdaTransitV1QueryParams) (*ResponseSdaDeleteFabricDeviceLayer3HandoffsWithSdaTransitV1, *resty.Response, error) {
	//DeleteFabricDeviceLayer3HandoffsWithSdaTransitV1QueryParams *DeleteFabricDeviceLayer3HandoffsWithSdaTransitV1QueryParams
	path := "/dna/intent/api/v1/sda/fabricDevices/layer3Handoffs/sdaTransits"

	queryString, _ := query.Values(DeleteFabricDeviceLayer3HandoffsWithSdaTransitV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteFabricDeviceLayer3HandoffsWithSdaTransitV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteFabricDeviceLayer3HandoffsWithSdaTransitV1(
				DeleteFabricDeviceLayer3HandoffsWithSdaTransitV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteFabricDeviceLayer3HandoffsWithSdaTransitV1")
	}

	result := response.Result().(*ResponseSdaDeleteFabricDeviceLayer3HandoffsWithSdaTransitV1)
	return result, response, err

}

//DeleteFabricDeviceByIDV1 Delete fabric device by id - 7593-7bef-4a68-b011
/* Deletes a fabric device based on id.


@param id id path parameter. ID of the fabric device.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-fabric-device-by-id-v1
*/
func (s *SdaService) DeleteFabricDeviceByIDV1(id string) (*ResponseSdaDeleteFabricDeviceByIDV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/sda/fabricDevices/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSdaDeleteFabricDeviceByIDV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteFabricDeviceByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteFabricDeviceByIdV1")
	}

	result := response.Result().(*ResponseSdaDeleteFabricDeviceByIDV1)
	return result, response, err

}

//DeleteFabricSiteByIDV1 Delete fabric site by id - aea4-2a3c-4799-8f73
/* Deletes a fabric site based on id.


@param id id path parameter. ID of the fabric site.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-fabric-site-by-id-v1
*/
func (s *SdaService) DeleteFabricSiteByIDV1(id string) (*ResponseSdaDeleteFabricSiteByIDV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/sda/fabricSites/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSdaDeleteFabricSiteByIDV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteFabricSiteByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteFabricSiteByIdV1")
	}

	result := response.Result().(*ResponseSdaDeleteFabricSiteByIDV1)
	return result, response, err

}

//DeleteFabricZoneByIDV1 Delete fabric zone by id - 78be-d947-4eb8-8fc3
/* Deletes a fabric zone based on id.


@param id id path parameter. ID of the fabric zone.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-fabric-zone-by-id-v1
*/
func (s *SdaService) DeleteFabricZoneByIDV1(id string) (*ResponseSdaDeleteFabricZoneByIDV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/sda/fabricZones/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSdaDeleteFabricZoneByIDV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteFabricZoneByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteFabricZoneByIdV1")
	}

	result := response.Result().(*ResponseSdaDeleteFabricZoneByIDV1)
	return result, response, err

}

//DeleteLayer2VirtualNetworksV1 Delete layer 2 virtual networks - d9a4-09cb-4b08-be45
/* Deletes layer 2 virtual networks based on user input.


@param DeleteLayer2VirtualNetworksV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-layer2-virtual-networks-v1
*/
func (s *SdaService) DeleteLayer2VirtualNetworksV1(DeleteLayer2VirtualNetworksV1QueryParams *DeleteLayer2VirtualNetworksV1QueryParams) (*ResponseSdaDeleteLayer2VirtualNetworksV1, *resty.Response, error) {
	//DeleteLayer2VirtualNetworksV1QueryParams *DeleteLayer2VirtualNetworksV1QueryParams
	path := "/dna/intent/api/v1/sda/layer2VirtualNetworks"

	queryString, _ := query.Values(DeleteLayer2VirtualNetworksV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteLayer2VirtualNetworksV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteLayer2VirtualNetworksV1(DeleteLayer2VirtualNetworksV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteLayer2VirtualNetworksV1")
	}

	result := response.Result().(*ResponseSdaDeleteLayer2VirtualNetworksV1)
	return result, response, err

}

//DeleteLayer2VirtualNetworkByIDV1 Delete layer 2 virtual network by id - b081-c850-4ab9-86f6
/* Deletes a layer 2 virtual network based on id.


@param id id path parameter. ID of the layer 2 virtual network.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-layer2-virtual-network-by-id-v1
*/
func (s *SdaService) DeleteLayer2VirtualNetworkByIDV1(id string) (*ResponseSdaDeleteLayer2VirtualNetworkByIDV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/sda/layer2VirtualNetworks/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSdaDeleteLayer2VirtualNetworkByIDV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteLayer2VirtualNetworkByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteLayer2VirtualNetworkByIdV1")
	}

	result := response.Result().(*ResponseSdaDeleteLayer2VirtualNetworkByIDV1)
	return result, response, err

}

//DeleteLayer3VirtualNetworksV1 Delete layer 3 virtual networks - 49bf-69ec-4a8a-a473
/* Deletes layer 3 virtual networks based on user input.


@param DeleteLayer3VirtualNetworksV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-layer3-virtual-networks-v1
*/
func (s *SdaService) DeleteLayer3VirtualNetworksV1(DeleteLayer3VirtualNetworksV1QueryParams *DeleteLayer3VirtualNetworksV1QueryParams) (*ResponseSdaDeleteLayer3VirtualNetworksV1, *resty.Response, error) {
	//DeleteLayer3VirtualNetworksV1QueryParams *DeleteLayer3VirtualNetworksV1QueryParams
	path := "/dna/intent/api/v1/sda/layer3VirtualNetworks"

	queryString, _ := query.Values(DeleteLayer3VirtualNetworksV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteLayer3VirtualNetworksV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteLayer3VirtualNetworksV1(DeleteLayer3VirtualNetworksV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteLayer3VirtualNetworksV1")
	}

	result := response.Result().(*ResponseSdaDeleteLayer3VirtualNetworksV1)
	return result, response, err

}

//DeleteLayer3VirtualNetworkByIDV1 Delete layer 3 virtual network by id - 4cb3-2a8c-4e38-b2ea
/* Deletes a layer 3 virtual network based on id.


@param id id path parameter. ID of the layer 3 virtual network.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-layer3-virtual-network-by-id-v1
*/
func (s *SdaService) DeleteLayer3VirtualNetworkByIDV1(id string) (*ResponseSdaDeleteLayer3VirtualNetworkByIDV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/sda/layer3VirtualNetworks/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSdaDeleteLayer3VirtualNetworkByIDV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteLayer3VirtualNetworkByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteLayer3VirtualNetworkByIdV1")
	}

	result := response.Result().(*ResponseSdaDeleteLayer3VirtualNetworkByIDV1)
	return result, response, err

}

//DeleteMulticastVirtualNetworkByIDV1 Delete multicast virtual network by id - 0b91-0b73-4649-a1d2
/* Deletes a multicast configuration for a virtual network based on id.


@param id id path parameter. ID of the multicast configuration.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-multicast-virtual-network-by-id-v1
*/
func (s *SdaService) DeleteMulticastVirtualNetworkByIDV1(id string) (*ResponseSdaDeleteMulticastVirtualNetworkByIDV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/sda/multicast/virtualNetworks/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSdaDeleteMulticastVirtualNetworkByIDV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteMulticastVirtualNetworkByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteMulticastVirtualNetworkByIdV1")
	}

	result := response.Result().(*ResponseSdaDeleteMulticastVirtualNetworkByIDV1)
	return result, response, err

}

//DeletePortAssignmentsV1 Delete port assignments - bd9b-8a54-494b-a3f1
/* Deletes port assignments based on user input.


@param DeletePortAssignmentsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-port-assignments-v1
*/
func (s *SdaService) DeletePortAssignmentsV1(DeletePortAssignmentsV1QueryParams *DeletePortAssignmentsV1QueryParams) (*ResponseSdaDeletePortAssignmentsV1, *resty.Response, error) {
	//DeletePortAssignmentsV1QueryParams *DeletePortAssignmentsV1QueryParams
	path := "/dna/intent/api/v1/sda/portAssignments"

	queryString, _ := query.Values(DeletePortAssignmentsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeletePortAssignmentsV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletePortAssignmentsV1(DeletePortAssignmentsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeletePortAssignmentsV1")
	}

	result := response.Result().(*ResponseSdaDeletePortAssignmentsV1)
	return result, response, err

}

//DeletePortAssignmentByIDV1 Delete port assignment by id - fdbe-aa08-422b-9fa1
/* Deletes a port assignment based on id.


@param id id path parameter. ID of the port assignment.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-port-assignment-by-id-v1
*/
func (s *SdaService) DeletePortAssignmentByIDV1(id string) (*ResponseSdaDeletePortAssignmentByIDV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/sda/portAssignments/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSdaDeletePortAssignmentByIDV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletePortAssignmentByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation DeletePortAssignmentByIdV1")
	}

	result := response.Result().(*ResponseSdaDeletePortAssignmentByIDV1)
	return result, response, err

}

//DeletePortChannelsV1 Delete port channels - ffb2-e803-4c7b-94ad
/* Deletes port channels based on user input.


@param DeletePortChannelsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-port-channels-v1
*/
func (s *SdaService) DeletePortChannelsV1(DeletePortChannelsV1QueryParams *DeletePortChannelsV1QueryParams) (*ResponseSdaDeletePortChannelsV1, *resty.Response, error) {
	//DeletePortChannelsV1QueryParams *DeletePortChannelsV1QueryParams
	path := "/dna/intent/api/v1/sda/portChannels"

	queryString, _ := query.Values(DeletePortChannelsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeletePortChannelsV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletePortChannelsV1(DeletePortChannelsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeletePortChannelsV1")
	}

	result := response.Result().(*ResponseSdaDeletePortChannelsV1)
	return result, response, err

}

//DeletePortChannelByIDV1 Delete port channel by id - 55ab-6978-47fa-b9d8
/* Deletes a port channel based on id.


@param id id path parameter. ID of the port channel.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-port-channel-by-id-v1
*/
func (s *SdaService) DeletePortChannelByIDV1(id string) (*ResponseSdaDeletePortChannelByIDV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/sda/portChannels/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSdaDeletePortChannelByIDV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletePortChannelByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation DeletePortChannelByIdV1")
	}

	result := response.Result().(*ResponseSdaDeletePortChannelByIDV1)
	return result, response, err

}

//DeleteProvisionedDevicesV1 Delete provisioned devices - 559a-8ac2-4729-a509
/* Delete provisioned devices based on query parameters.


@param DeleteProvisionedDevicesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-provisioned-devices-v1
*/
func (s *SdaService) DeleteProvisionedDevicesV1(DeleteProvisionedDevicesV1QueryParams *DeleteProvisionedDevicesV1QueryParams) (*ResponseSdaDeleteProvisionedDevicesV1, *resty.Response, error) {
	//DeleteProvisionedDevicesV1QueryParams *DeleteProvisionedDevicesV1QueryParams
	path := "/dna/intent/api/v1/sda/provisionDevices"

	queryString, _ := query.Values(DeleteProvisionedDevicesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteProvisionedDevicesV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteProvisionedDevicesV1(DeleteProvisionedDevicesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteProvisionedDevicesV1")
	}

	result := response.Result().(*ResponseSdaDeleteProvisionedDevicesV1)
	return result, response, err

}

//DeleteProvisionedDeviceByIDV1 Delete provisioned device by Id - 8bb4-88f0-4f58-9856
/* Deletes provisioned device based on Id.


@param id id path parameter. ID of the provisioned device.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-provisioned-device-by-id-v1
*/
func (s *SdaService) DeleteProvisionedDeviceByIDV1(id string) (*ResponseSdaDeleteProvisionedDeviceByIDV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/sda/provisionDevices/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSdaDeleteProvisionedDeviceByIDV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteProvisionedDeviceByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteProvisionedDeviceByIdV1")
	}

	result := response.Result().(*ResponseSdaDeleteProvisionedDeviceByIDV1)
	return result, response, err

}

//DeleteTransitNetworkByIDV1 Delete transit network by id - 91bd-2956-4359-a935
/* Deletes a transit network based on id.


@param id id path parameter. ID of the transit network.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-transit-network-by-id-v1
*/
func (s *SdaService) DeleteTransitNetworkByIDV1(id string) (*ResponseSdaDeleteTransitNetworkByIDV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/sda/transitNetworks/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSdaDeleteTransitNetworkByIDV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteTransitNetworkByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteTransitNetworkByIdV1")
	}

	result := response.Result().(*ResponseSdaDeleteTransitNetworkByIDV1)
	return result, response, err

}

//DeleteVirtualNetworkWithScalableGroupsV1 Delete virtual network with scalable groups - c8b6-0bc3-4808-8d56
/* Delete virtual network with scalable groups


@param DeleteVirtualNetworkWithScalableGroupsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-virtual-network-with-scalable-groups-v1
*/
func (s *SdaService) DeleteVirtualNetworkWithScalableGroupsV1(DeleteVirtualNetworkWithScalableGroupsV1QueryParams *DeleteVirtualNetworkWithScalableGroupsV1QueryParams) (*ResponseSdaDeleteVirtualNetworkWithScalableGroupsV1, *resty.Response, error) {
	//DeleteVirtualNetworkWithScalableGroupsV1QueryParams *DeleteVirtualNetworkWithScalableGroupsV1QueryParams
	path := "/dna/intent/api/v1/virtual-network"

	queryString, _ := query.Values(DeleteVirtualNetworkWithScalableGroupsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteVirtualNetworkWithScalableGroupsV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteVirtualNetworkWithScalableGroupsV1(DeleteVirtualNetworkWithScalableGroupsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteVirtualNetworkWithScalableGroupsV1")
	}

	result := response.Result().(*ResponseSdaDeleteVirtualNetworkWithScalableGroupsV1)
	return result, response, err

}

// Alias Function
func (s *SdaService) GetDefaultAuthenticationProfileFromSdaFabric(GetDefaultAuthenticationProfileFromSDAFabricV1QueryParams *GetDefaultAuthenticationProfileFromSdaFabricV1QueryParams) (*ResponseSdaGetDefaultAuthenticationProfileFromSdaFabricV1, *resty.Response, error) {
	return s.GetDefaultAuthenticationProfileFromSdaFabricV1(GetDefaultAuthenticationProfileFromSDAFabricV1QueryParams)
}

// Alias Function
func (s *SdaService) GetDeviceInfoFromSdaFabric(GetDeviceInfoFromSDAFabricV1QueryParams *GetDeviceInfoFromSdaFabricV1QueryParams) (*ResponseSdaGetDeviceInfoFromSdaFabricV1, *resty.Response, error) {
	return s.GetDeviceInfoFromSdaFabricV1(GetDeviceInfoFromSDAFabricV1QueryParams)
}

// Alias Function
func (s *SdaService) AddBorderDeviceInSdaFabric(requestSdaAddBorderDeviceInSDAFabricV1 *RequestSdaAddBorderDeviceInSdaFabricV1) (*ResponseSdaAddBorderDeviceInSdaFabricV1, *resty.Response, error) {
	return s.AddBorderDeviceInSdaFabricV1(requestSdaAddBorderDeviceInSDAFabricV1)
}

// Alias Function
func (s *SdaService) GetSiteFromSdaFabric(GetSiteFromSDAFabricV1QueryParams *GetSiteFromSdaFabricV1QueryParams) (*ResponseSdaGetSiteFromSdaFabricV1, *resty.Response, error) {
	return s.GetSiteFromSdaFabricV1(GetSiteFromSDAFabricV1QueryParams)
}

// Alias Function
func (s *SdaService) AddPortAssignmentForAccessPointInSdaFabric(requestSdaAddPortAssignmentForAccessPointInSDAFabricV1 *RequestSdaAddPortAssignmentForAccessPointInSdaFabricV1) (*ResponseSdaAddPortAssignmentForAccessPointInSdaFabricV1, *resty.Response, error) {
	return s.AddPortAssignmentForAccessPointInSdaFabricV1(requestSdaAddPortAssignmentForAccessPointInSDAFabricV1)
}

// Alias Function
func (s *SdaService) UpdateAnycastGateways(requestSdaUpdateAnycastGatewaysV1 *RequestSdaUpdateAnycastGatewaysV1) (*ResponseSdaUpdateAnycastGatewaysV1, *resty.Response, error) {
	return s.UpdateAnycastGatewaysV1(requestSdaUpdateAnycastGatewaysV1)
}

// Alias Function
func (s *SdaService) GetIPPoolFromSdaVirtualNetwork(GetIPPoolFromSDAVirtualNetworkV1QueryParams *GetIPPoolFromSdaVirtualNetworkV1QueryParams) (*ResponseSdaGetIPPoolFromSdaVirtualNetworkV1, *resty.Response, error) {
	return s.GetIPPoolFromSdaVirtualNetworkV1(GetIPPoolFromSDAVirtualNetworkV1QueryParams)
}

// Alias Function
func (s *SdaService) AddAnycastGateways(requestSdaAddAnycastGatewaysV1 *RequestSdaAddAnycastGatewaysV1) (*ResponseSdaAddAnycastGatewaysV1, *resty.Response, error) {
	return s.AddAnycastGatewaysV1(requestSdaAddAnycastGatewaysV1)
}

// Alias Function
func (s *SdaService) AddDefaultAuthenticationTemplateInSdaFabric(requestSdaAddDefaultAuthenticationTemplateInSDAFabricV1 *RequestSdaAddDefaultAuthenticationTemplateInSdaFabricV1) (*ResponseSdaAddDefaultAuthenticationTemplateInSdaFabricV1, *resty.Response, error) {
	return s.AddDefaultAuthenticationTemplateInSdaFabricV1(requestSdaAddDefaultAuthenticationTemplateInSDAFabricV1)
}

// Alias Function
func (s *SdaService) DeleteIPPoolFromSdaVirtualNetwork(DeleteIPPoolFromSDAVirtualNetworkV1QueryParams *DeleteIPPoolFromSdaVirtualNetworkV1QueryParams) (*ResponseSdaDeleteIPPoolFromSdaVirtualNetworkV1, *resty.Response, error) {
	return s.DeleteIPPoolFromSdaVirtualNetworkV1(DeleteIPPoolFromSDAVirtualNetworkV1QueryParams)
}

// Alias Function
func (s *SdaService) AddIPPoolInSdaVirtualNetwork(requestSdaAddIPPoolInSDAVirtualNetworkV1 *RequestSdaAddIPPoolInSdaVirtualNetworkV1) (*ResponseSdaAddIPPoolInSdaVirtualNetworkV1, *resty.Response, error) {
	return s.AddIPPoolInSdaVirtualNetworkV1(requestSdaAddIPPoolInSDAVirtualNetworkV1)
}

// Alias Function
func (s *SdaService) DeleteFabricDeviceLayer2Handoffs(DeleteFabricDeviceLayer2HandoffsV1QueryParams *DeleteFabricDeviceLayer2HandoffsV1QueryParams) (*ResponseSdaDeleteFabricDeviceLayer2HandoffsV1, *resty.Response, error) {
	return s.DeleteFabricDeviceLayer2HandoffsV1(DeleteFabricDeviceLayer2HandoffsV1QueryParams)
}

// Alias Function
func (s *SdaService) GetAuthenticationProfiles(GetAuthenticationProfilesV1QueryParams *GetAuthenticationProfilesV1QueryParams) (*ResponseSdaGetAuthenticationProfilesV1, *resty.Response, error) {
	return s.GetAuthenticationProfilesV1(GetAuthenticationProfilesV1QueryParams)
}

// Alias Function
func (s *SdaService) GetFabricDevicesLayer3HandoffsWithIPTransitCount(GetFabricDevicesLayer3HandoffsWithIpTransitCountV1QueryParams *GetFabricDevicesLayer3HandoffsWithIPTransitCountV1QueryParams) (*ResponseSdaGetFabricDevicesLayer3HandoffsWithIPTransitCountV1, *resty.Response, error) {
	return s.GetFabricDevicesLayer3HandoffsWithIPTransitCountV1(GetFabricDevicesLayer3HandoffsWithIpTransitCountV1QueryParams)
}

// Alias Function
func (s *SdaService) AddFabricDevicesLayer3HandoffsWithIPTransit(requestSdaAddFabricDevicesLayer3HandoffsWithIpTransitV1 *RequestSdaAddFabricDevicesLayer3HandoffsWithIPTransitV1) (*ResponseSdaAddFabricDevicesLayer3HandoffsWithIPTransitV1, *resty.Response, error) {
	return s.AddFabricDevicesLayer3HandoffsWithIPTransitV1(requestSdaAddFabricDevicesLayer3HandoffsWithIpTransitV1)
}

// Alias Function
func (s *SdaService) GetLayer3VirtualNetworksCount(GetLayer3VirtualNetworksCountV1QueryParams *GetLayer3VirtualNetworksCountV1QueryParams) (*ResponseSdaGetLayer3VirtualNetworksCountV1, *resty.Response, error) {
	return s.GetLayer3VirtualNetworksCountV1(GetLayer3VirtualNetworksCountV1QueryParams)
}

// Alias Function
func (s *SdaService) AddPortAssignments(requestSdaAddPortAssignmentsV1 *RequestSdaAddPortAssignmentsV1) (*ResponseSdaAddPortAssignmentsV1, *resty.Response, error) {
	return s.AddPortAssignmentsV1(requestSdaAddPortAssignmentsV1)
}

// Alias Function
func (s *SdaService) DeleteLayer2VirtualNetworkByID(id string) (*ResponseSdaDeleteLayer2VirtualNetworkByIDV1, *resty.Response, error) {
	return s.DeleteLayer2VirtualNetworkByIDV1(id)
}

// Alias Function
func (s *SdaService) DeleteExtranetPolicies(DeleteExtranetPoliciesV1QueryParams *DeleteExtranetPoliciesV1QueryParams) (*ResponseSdaDeleteExtranetPoliciesV1, *resty.Response, error) {
	return s.DeleteExtranetPoliciesV1(DeleteExtranetPoliciesV1QueryParams)
}

// Alias Function
func (s *SdaService) GetFabricSites(GetFabricSitesV1QueryParams *GetFabricSitesV1QueryParams) (*ResponseSdaGetFabricSitesV1, *resty.Response, error) {
	return s.GetFabricSitesV1(GetFabricSitesV1QueryParams)
}

// Alias Function
func (s *SdaService) AddVirtualNetworkWithScalableGroups(requestSdaAddVirtualNetworkWithScalableGroupsV1 *RequestSdaAddVirtualNetworkWithScalableGroupsV1) (*ResponseSdaAddVirtualNetworkWithScalableGroupsV1, *resty.Response, error) {
	return s.AddVirtualNetworkWithScalableGroupsV1(requestSdaAddVirtualNetworkWithScalableGroupsV1)
}

// Alias Function
func (s *SdaService) DeletePortAssignmentForAccessPointInSdaFabric(DeletePortAssignmentForAccessPointInSDAFabricV1QueryParams *DeletePortAssignmentForAccessPointInSdaFabricV1QueryParams) (*ResponseSdaDeletePortAssignmentForAccessPointInSdaFabricV1, *resty.Response, error) {
	return s.DeletePortAssignmentForAccessPointInSdaFabricV1(DeletePortAssignmentForAccessPointInSDAFabricV1QueryParams)
}

// Alias Function
func (s *SdaService) AddMulticastInSdaFabric(requestSdaAddMulticastInSDAFabricV1 *RequestSdaAddMulticastInSdaFabricV1) (*ResponseSdaAddMulticastInSdaFabricV1, *resty.Response, error) {
	return s.AddMulticastInSdaFabricV1(requestSdaAddMulticastInSDAFabricV1)
}

// Alias Function
func (s *SdaService) UpdateTransitNetworks(requestSdaUpdateTransitNetworksV1 *RequestSdaUpdateTransitNetworksV1) (*ResponseSdaUpdateTransitNetworksV1, *resty.Response, error) {
	return s.UpdateTransitNetworksV1(requestSdaUpdateTransitNetworksV1)
}

// Alias Function
func (s *SdaService) DeleteFabricDeviceLayer3HandoffWithIPTransitByID(id string) (*ResponseSdaDeleteFabricDeviceLayer3HandoffWithIPTransitByIDV1, *resty.Response, error) {
	return s.DeleteFabricDeviceLayer3HandoffWithIPTransitByIDV1(id)
}

// Alias Function
func (s *SdaService) UpdateExtranetPolicy(requestSdaUpdateExtranetPolicyV1 *RequestSdaUpdateExtranetPolicyV1) (*ResponseSdaUpdateExtranetPolicyV1, *resty.Response, error) {
	return s.UpdateExtranetPolicyV1(requestSdaUpdateExtranetPolicyV1)
}

// Alias Function
func (s *SdaService) DeleteFabricSiteByID(id string) (*ResponseSdaDeleteFabricSiteByIDV1, *resty.Response, error) {
	return s.DeleteFabricSiteByIDV1(id)
}

// Alias Function
func (s *SdaService) GetFabricDevicesLayer2Handoffs(GetFabricDevicesLayer2HandoffsV1QueryParams *GetFabricDevicesLayer2HandoffsV1QueryParams) (*ResponseSdaGetFabricDevicesLayer2HandoffsV1, *resty.Response, error) {
	return s.GetFabricDevicesLayer2HandoffsV1(GetFabricDevicesLayer2HandoffsV1QueryParams)
}

// Alias Function
func (s *SdaService) GetFabricDevicesCount(GetFabricDevicesCountV1QueryParams *GetFabricDevicesCountV1QueryParams) (*ResponseSdaGetFabricDevicesCountV1, *resty.Response, error) {
	return s.GetFabricDevicesCountV1(GetFabricDevicesCountV1QueryParams)
}

// Alias Function
func (s *SdaService) DeleteFabricDeviceLayer2HandoffByID(id string) (*ResponseSdaDeleteFabricDeviceLayer2HandoffByIDV1, *resty.Response, error) {
	return s.DeleteFabricDeviceLayer2HandoffByIDV1(id)
}

// Alias Function
func (s *SdaService) GetPortChannels(GetPortChannelsV1QueryParams *GetPortChannelsV1QueryParams) (*ResponseSdaGetPortChannelsV1, *resty.Response, error) {
	return s.GetPortChannelsV1(GetPortChannelsV1QueryParams)
}

// Alias Function
func (s *SdaService) AddEdgeDeviceInSdaFabric(requestSdaAddEdgeDeviceInSDAFabricV1 *RequestSdaAddEdgeDeviceInSdaFabricV1) (*ResponseSdaAddEdgeDeviceInSdaFabricV1, *resty.Response, error) {
	return s.AddEdgeDeviceInSdaFabricV1(requestSdaAddEdgeDeviceInSDAFabricV1)
}

// Alias Function
func (s *SdaService) GetDeviceRoleInSdaFabric(GetDeviceRoleInSDAFabricV1QueryParams *GetDeviceRoleInSdaFabricV1QueryParams) (*ResponseSdaGetDeviceRoleInSdaFabricV1, *resty.Response, error) {
	return s.GetDeviceRoleInSdaFabricV1(GetDeviceRoleInSDAFabricV1QueryParams)
}

// Alias Function
func (s *SdaService) DeletePortChannelByID(id string) (*ResponseSdaDeletePortChannelByIDV1, *resty.Response, error) {
	return s.DeletePortChannelByIDV1(id)
}

// Alias Function
func (s *SdaService) UpdateFabricDevicesLayer3HandoffsWithIPTransit(requestSdaUpdateFabricDevicesLayer3HandoffsWithIpTransitV1 *RequestSdaUpdateFabricDevicesLayer3HandoffsWithIPTransitV1) (*ResponseSdaUpdateFabricDevicesLayer3HandoffsWithIPTransitV1, *resty.Response, error) {
	return s.UpdateFabricDevicesLayer3HandoffsWithIPTransitV1(requestSdaUpdateFabricDevicesLayer3HandoffsWithIpTransitV1)
}

// Alias Function
func (s *SdaService) UpdateMulticast(requestSdaUpdateMulticastV1 *RequestSdaUpdateMulticastV1) (*ResponseSdaUpdateMulticastV1, *resty.Response, error) {
	return s.UpdateMulticastV1(requestSdaUpdateMulticastV1)
}

// Alias Function
func (s *SdaService) GetTransitNetworksCount(GetTransitNetworksCountV1QueryParams *GetTransitNetworksCountV1QueryParams) (*ResponseSdaGetTransitNetworksCountV1, *resty.Response, error) {
	return s.GetTransitNetworksCountV1(GetTransitNetworksCountV1QueryParams)
}

// Alias Function
func (s *SdaService) AddLayer3VirtualNetworks(requestSdaAddLayer3VirtualNetworksV1 *RequestSdaAddLayer3VirtualNetworksV1) (*ResponseSdaAddLayer3VirtualNetworksV1, *resty.Response, error) {
	return s.AddLayer3VirtualNetworksV1(requestSdaAddLayer3VirtualNetworksV1)
}

// Alias Function
func (s *SdaService) GetLayer3VirtualNetworks(GetLayer3VirtualNetworksV1QueryParams *GetLayer3VirtualNetworksV1QueryParams) (*ResponseSdaGetLayer3VirtualNetworksV1, *resty.Response, error) {
	return s.GetLayer3VirtualNetworksV1(GetLayer3VirtualNetworksV1QueryParams)
}

// Alias Function
func (s *SdaService) GetPortAssignmentCount(GetPortAssignmentCountV1QueryParams *GetPortAssignmentCountV1QueryParams) (*ResponseSdaGetPortAssignmentCountV1, *resty.Response, error) {
	return s.GetPortAssignmentCountV1(GetPortAssignmentCountV1QueryParams)
}

// Alias Function
func (s *SdaService) DeleteExtranetPolicyByID(id string) (*ResponseSdaDeleteExtranetPolicyByIDV1, *resty.Response, error) {
	return s.DeleteExtranetPolicyByIDV1(id)
}

// Alias Function
func (s *SdaService) DeleteMulticastVirtualNetworkByID(id string) (*ResponseSdaDeleteMulticastVirtualNetworkByIDV1, *resty.Response, error) {
	return s.DeleteMulticastVirtualNetworkByIDV1(id)
}

// Alias Function
func (s *SdaService) DeleteVirtualNetworkWithScalableGroups(DeleteVirtualNetworkWithScalableGroupsV1QueryParams *DeleteVirtualNetworkWithScalableGroupsV1QueryParams) (*ResponseSdaDeleteVirtualNetworkWithScalableGroupsV1, *resty.Response, error) {
	return s.DeleteVirtualNetworkWithScalableGroupsV1(DeleteVirtualNetworkWithScalableGroupsV1QueryParams)
}

// Alias Function
func (s *SdaService) DeleteProvisionedDevices(DeleteProvisionedDevicesV1QueryParams *DeleteProvisionedDevicesV1QueryParams) (*ResponseSdaDeleteProvisionedDevicesV1, *resty.Response, error) {
	return s.DeleteProvisionedDevicesV1(DeleteProvisionedDevicesV1QueryParams)
}

// Alias Function
func (s *SdaService) DeleteLayer3VirtualNetworkByID(id string) (*ResponseSdaDeleteLayer3VirtualNetworkByIDV1, *resty.Response, error) {
	return s.DeleteLayer3VirtualNetworkByIDV1(id)
}

// Alias Function
func (s *SdaService) DeleteTransitPeerNetwork(DeleteTransitPeerNetworkV1QueryParams *DeleteTransitPeerNetworkV1QueryParams) (*ResponseSdaDeleteTransitPeerNetworkV1, *resty.Response, error) {
	return s.DeleteTransitPeerNetworkV1(DeleteTransitPeerNetworkV1QueryParams)
}

// Alias Function
func (s *SdaService) DeleteProvisionedDeviceByID(id string) (*ResponseSdaDeleteProvisionedDeviceByIDV1, *resty.Response, error) {
	return s.DeleteProvisionedDeviceByIDV1(id)
}

// Alias Function
func (s *SdaService) GetFabricDevices(GetFabricDevicesV1QueryParams *GetFabricDevicesV1QueryParams) (*ResponseSdaGetFabricDevicesV1, *resty.Response, error) {
	return s.GetFabricDevicesV1(GetFabricDevicesV1QueryParams)
}

// Alias Function
func (s *SdaService) AddMulticastVirtualNetworks(requestSdaAddMulticastVirtualNetworksV1 *RequestSdaAddMulticastVirtualNetworksV1) (*ResponseSdaAddMulticastVirtualNetworksV1, *resty.Response, error) {
	return s.AddMulticastVirtualNetworksV1(requestSdaAddMulticastVirtualNetworksV1)
}

// Alias Function
func (s *SdaService) GetTransitNetworks(GetTransitNetworksV1QueryParams *GetTransitNetworksV1QueryParams) (*ResponseSdaGetTransitNetworksV1, *resty.Response, error) {
	return s.GetTransitNetworksV1(GetTransitNetworksV1QueryParams)
}

// Alias Function
func (s *SdaService) DeleteVnFromSdaFabric(DeleteVNFromSDAFabricV1QueryParams *DeleteVnFromSdaFabricV1QueryParams) (*ResponseSdaDeleteVnFromSdaFabricV1, *resty.Response, error) {
	return s.DeleteVnFromSdaFabricV1(DeleteVNFromSDAFabricV1QueryParams)
}

// Alias Function
func (s *SdaService) AddFabricDevicesLayer2Handoffs(requestSdaAddFabricDevicesLayer2HandoffsV1 *RequestSdaAddFabricDevicesLayer2HandoffsV1) (*ResponseSdaAddFabricDevicesLayer2HandoffsV1, *resty.Response, error) {
	return s.AddFabricDevicesLayer2HandoffsV1(requestSdaAddFabricDevicesLayer2HandoffsV1)
}

// Alias Function
func (s *SdaService) UpdateFabricDevices(requestSdaUpdateFabricDevicesV1 *RequestSdaUpdateFabricDevicesV1) (*ResponseSdaUpdateFabricDevicesV1, *resty.Response, error) {
	return s.UpdateFabricDevicesV1(requestSdaUpdateFabricDevicesV1)
}

// Alias Function
func (s *SdaService) GetBorderDeviceDetailFromSdaFabric(GetBorderDeviceDetailFromSDAFabricV1QueryParams *GetBorderDeviceDetailFromSdaFabricV1QueryParams) (*ResponseSdaGetBorderDeviceDetailFromSdaFabricV1, *resty.Response, error) {
	return s.GetBorderDeviceDetailFromSdaFabricV1(GetBorderDeviceDetailFromSDAFabricV1QueryParams)
}

// Alias Function
func (s *SdaService) GetExtranetPolicyCount() (*ResponseSdaGetExtranetPolicyCountV1, *resty.Response, error) {
	return s.GetExtranetPolicyCountV1()
}

// Alias Function
func (s *SdaService) GetTransitPeerNetworkInfo(GetTransitPeerNetworkInfoV1QueryParams *GetTransitPeerNetworkInfoV1QueryParams) (*ResponseSdaGetTransitPeerNetworkInfoV1, *resty.Response, error) {
	return s.GetTransitPeerNetworkInfoV1(GetTransitPeerNetworkInfoV1QueryParams)
}

// Alias Function
func (s *SdaService) DeleteFabricDeviceByID(id string) (*ResponseSdaDeleteFabricDeviceByIDV1, *resty.Response, error) {
	return s.DeleteFabricDeviceByIDV1(id)
}

// Alias Function
func (s *SdaService) GetPortAssignments(GetPortAssignmentsV1QueryParams *GetPortAssignmentsV1QueryParams) (*ResponseSdaGetPortAssignmentsV1, *resty.Response, error) {
	return s.GetPortAssignmentsV1(GetPortAssignmentsV1QueryParams)
}

// Alias Function
func (s *SdaService) GetVirtualNetworkWithScalableGroups(GetVirtualNetworkWithScalableGroupsV1QueryParams *GetVirtualNetworkWithScalableGroupsV1QueryParams) (*ResponseSdaGetVirtualNetworkWithScalableGroupsV1, *resty.Response, error) {
	return s.GetVirtualNetworkWithScalableGroupsV1(GetVirtualNetworkWithScalableGroupsV1QueryParams)
}

// Alias Function
func (s *SdaService) GetMulticastDetailsFromSdaFabric(GetMulticastDetailsFromSDAFabricV1QueryParams *GetMulticastDetailsFromSdaFabricV1QueryParams) (*ResponseSdaGetMulticastDetailsFromSdaFabricV1, *resty.Response, error) {
	return s.GetMulticastDetailsFromSdaFabricV1(GetMulticastDetailsFromSDAFabricV1QueryParams)
}

// Alias Function
func (s *SdaService) DeleteProvisionedWiredDevice(DeleteProvisionedWiredDeviceV1QueryParams *DeleteProvisionedWiredDeviceV1QueryParams) (*ResponseSdaDeleteProvisionedWiredDeviceV1, *resty.Response, error) {
	return s.DeleteProvisionedWiredDeviceV1(DeleteProvisionedWiredDeviceV1QueryParams)
}

// Alias Function
func (s *SdaService) DeleteAnycastGatewayByID(id string) (*ResponseSdaDeleteAnycastGatewayByIDV1, *resty.Response, error) {
	return s.DeleteAnycastGatewayByIDV1(id)
}

// Alias Function
func (s *SdaService) GetVirtualNetworkSummary(GetVirtualNetworkSummaryV1QueryParams *GetVirtualNetworkSummaryV1QueryParams) (*ResponseSdaGetVirtualNetworkSummaryV1, *resty.Response, error) {
	return s.GetVirtualNetworkSummaryV1(GetVirtualNetworkSummaryV1QueryParams)
}

// Alias Function
func (s *SdaService) AddPortChannels(requestSdaAddPortChannelsV1 *RequestSdaAddPortChannelsV1) (*ResponseSdaAddPortChannelsV1, *resty.Response, error) {
	return s.AddPortChannelsV1(requestSdaAddPortChannelsV1)
}

// Alias Function
func (s *SdaService) DeleteEdgeDeviceFromSdaFabric(DeleteEdgeDeviceFromSDAFabricV1QueryParams *DeleteEdgeDeviceFromSdaFabricV1QueryParams) (*ResponseSdaDeleteEdgeDeviceFromSdaFabricV1, *resty.Response, error) {
	return s.DeleteEdgeDeviceFromSdaFabricV1(DeleteEdgeDeviceFromSDAFabricV1QueryParams)
}

// Alias Function
func (s *SdaService) DeleteFabricDevices(DeleteFabricDevicesV1QueryParams *DeleteFabricDevicesV1QueryParams) (*ResponseSdaDeleteFabricDevicesV1, *resty.Response, error) {
	return s.DeleteFabricDevicesV1(DeleteFabricDevicesV1QueryParams)
}

// Alias Function
func (s *SdaService) GetFabricZones(GetFabricZonesV1QueryParams *GetFabricZonesV1QueryParams) (*ResponseSdaGetFabricZonesV1, *resty.Response, error) {
	return s.GetFabricZonesV1(GetFabricZonesV1QueryParams)
}

// Alias Function
func (s *SdaService) UpdateFabricDevicesLayer3HandoffsWithSdaTransit(requestSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransitV1 *RequestSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransitV1) (*ResponseSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransitV1, *resty.Response, error) {
	return s.UpdateFabricDevicesLayer3HandoffsWithSdaTransitV1(requestSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransitV1)
}

// Alias Function
func (s *SdaService) UpdateAuthenticationProfile(requestSdaUpdateAuthenticationProfileV1 *RequestSdaUpdateAuthenticationProfileV1) (*ResponseSdaUpdateAuthenticationProfileV1, *resty.Response, error) {
	return s.UpdateAuthenticationProfileV1(requestSdaUpdateAuthenticationProfileV1)
}

// Alias Function
func (s *SdaService) DeleteControlPlaneDeviceInSdaFabric(DeleteControlPlaneDeviceInSDAFabricV1QueryParams *DeleteControlPlaneDeviceInSdaFabricV1QueryParams) (*ResponseSdaDeleteControlPlaneDeviceInSdaFabricV1, *resty.Response, error) {
	return s.DeleteControlPlaneDeviceInSdaFabricV1(DeleteControlPlaneDeviceInSDAFabricV1QueryParams)
}

// Alias Function
func (s *SdaService) GetMulticastVirtualNetworks(GetMulticastVirtualNetworksV1QueryParams *GetMulticastVirtualNetworksV1QueryParams) (*ResponseSdaGetMulticastVirtualNetworksV1, *resty.Response, error) {
	return s.GetMulticastVirtualNetworksV1(GetMulticastVirtualNetworksV1QueryParams)
}

// Alias Function
func (s *SdaService) UpdatePortAssignments(requestSdaUpdatePortAssignmentsV1 *RequestSdaUpdatePortAssignmentsV1) (*ResponseSdaUpdatePortAssignmentsV1, *resty.Response, error) {
	return s.UpdatePortAssignmentsV1(requestSdaUpdatePortAssignmentsV1)
}

// Alias Function
func (s *SdaService) GetFabricDevicesLayer3HandoffsWithSdaTransit(GetFabricDevicesLayer3HandoffsWithSdaTransitV1QueryParams *GetFabricDevicesLayer3HandoffsWithSdaTransitV1QueryParams) (*ResponseSdaGetFabricDevicesLayer3HandoffsWithSdaTransitV1, *resty.Response, error) {
	return s.GetFabricDevicesLayer3HandoffsWithSdaTransitV1(GetFabricDevicesLayer3HandoffsWithSdaTransitV1QueryParams)
}

// Alias Function
func (s *SdaService) DeleteFabricDeviceLayer3HandoffsWithSdaTransit(DeleteFabricDeviceLayer3HandoffsWithSdaTransitV1QueryParams *DeleteFabricDeviceLayer3HandoffsWithSdaTransitV1QueryParams) (*ResponseSdaDeleteFabricDeviceLayer3HandoffsWithSdaTransitV1, *resty.Response, error) {
	return s.DeleteFabricDeviceLayer3HandoffsWithSdaTransitV1(DeleteFabricDeviceLayer3HandoffsWithSdaTransitV1QueryParams)
}

// Alias Function
func (s *SdaService) AddLayer2VirtualNetworks(requestSdaAddLayer2VirtualNetworksV1 *RequestSdaAddLayer2VirtualNetworksV1) (*ResponseSdaAddLayer2VirtualNetworksV1, *resty.Response, error) {
	return s.AddLayer2VirtualNetworksV1(requestSdaAddLayer2VirtualNetworksV1)
}

// Alias Function
func (s *SdaService) AddFabricDevicesLayer3HandoffsWithSdaTransit(requestSdaAddFabricDevicesLayer3HandoffsWithSdaTransitV1 *RequestSdaAddFabricDevicesLayer3HandoffsWithSdaTransitV1) (*ResponseSdaAddFabricDevicesLayer3HandoffsWithSdaTransitV1, *resty.Response, error) {
	return s.AddFabricDevicesLayer3HandoffsWithSdaTransitV1(requestSdaAddFabricDevicesLayer3HandoffsWithSdaTransitV1)
}

// Alias Function
func (s *SdaService) DeletePortAssignmentForUserDeviceInSdaFabric(DeletePortAssignmentForUserDeviceInSDAFabricV1QueryParams *DeletePortAssignmentForUserDeviceInSdaFabricV1QueryParams) (*ResponseSdaDeletePortAssignmentForUserDeviceInSdaFabricV1, *resty.Response, error) {
	return s.DeletePortAssignmentForUserDeviceInSdaFabricV1(DeletePortAssignmentForUserDeviceInSDAFabricV1QueryParams)
}

// Alias Function
func (s *SdaService) GetFabricDevicesLayer2HandoffsCount(GetFabricDevicesLayer2HandoffsCountV1QueryParams *GetFabricDevicesLayer2HandoffsCountV1QueryParams) (*ResponseSdaGetFabricDevicesLayer2HandoffsCountV1, *resty.Response, error) {
	return s.GetFabricDevicesLayer2HandoffsCountV1(GetFabricDevicesLayer2HandoffsCountV1QueryParams)
}

// Alias Function
func (s *SdaService) AddTransitNetworks(requestSdaAddTransitNetworksV1 *RequestSdaAddTransitNetworksV1) (*ResponseSdaAddTransitNetworksV1, *resty.Response, error) {
	return s.AddTransitNetworksV1(requestSdaAddTransitNetworksV1)
}

// Alias Function
func (s *SdaService) UpdateMulticastVirtualNetworks(requestSdaUpdateMulticastVirtualNetworksV1 *RequestSdaUpdateMulticastVirtualNetworksV1) (*ResponseSdaUpdateMulticastVirtualNetworksV1, *resty.Response, error) {
	return s.UpdateMulticastVirtualNetworksV1(requestSdaUpdateMulticastVirtualNetworksV1)
}

// Alias Function
func (s *SdaService) DeleteTransitNetworkByID(id string) (*ResponseSdaDeleteTransitNetworkByIDV1, *resty.Response, error) {
	return s.DeleteTransitNetworkByIDV1(id)
}

// Alias Function
func (s *SdaService) DeleteFabricZoneByID(id string) (*ResponseSdaDeleteFabricZoneByIDV1, *resty.Response, error) {
	return s.DeleteFabricZoneByIDV1(id)
}

// Alias Function
func (s *SdaService) GetProvisionedDevicesCount(GetProvisionedDevicesCountV1QueryParams *GetProvisionedDevicesCountV1QueryParams) (*ResponseSdaGetProvisionedDevicesCountV1, *resty.Response, error) {
	return s.GetProvisionedDevicesCountV1(GetProvisionedDevicesCountV1QueryParams)
}

// Alias Function
func (s *SdaService) DeleteSiteFromSdaFabric(DeleteSiteFromSDAFabricV1QueryParams *DeleteSiteFromSdaFabricV1QueryParams) (*ResponseSdaDeleteSiteFromSdaFabricV1, *resty.Response, error) {
	return s.DeleteSiteFromSdaFabricV1(DeleteSiteFromSDAFabricV1QueryParams)
}

// Alias Function
func (s *SdaService) AddExtranetPolicy(requestSdaAddExtranetPolicyV1 *RequestSdaAddExtranetPolicyV1) (*ResponseSdaAddExtranetPolicyV1, *resty.Response, error) {
	return s.AddExtranetPolicyV1(requestSdaAddExtranetPolicyV1)
}

// Alias Function
func (s *SdaService) DeleteLayer3VirtualNetworks(DeleteLayer3VirtualNetworksV1QueryParams *DeleteLayer3VirtualNetworksV1QueryParams) (*ResponseSdaDeleteLayer3VirtualNetworksV1, *resty.Response, error) {
	return s.DeleteLayer3VirtualNetworksV1(DeleteLayer3VirtualNetworksV1QueryParams)
}

// Alias Function
func (s *SdaService) UpdateFabricZone(requestSdaUpdateFabricZoneV1 *RequestSdaUpdateFabricZoneV1) (*ResponseSdaUpdateFabricZoneV1, *resty.Response, error) {
	return s.UpdateFabricZoneV1(requestSdaUpdateFabricZoneV1)
}

// Alias Function
func (s *SdaService) GetMulticast(GetMulticastV1QueryParams *GetMulticastV1QueryParams) (*ResponseSdaGetMulticastV1, *resty.Response, error) {
	return s.GetMulticastV1(GetMulticastV1QueryParams)
}

// Alias Function
func (s *SdaService) AddFabricZone(requestSdaAddFabricZoneV1 *RequestSdaAddFabricZoneV1) (*ResponseSdaAddFabricZoneV1, *resty.Response, error) {
	return s.AddFabricZoneV1(requestSdaAddFabricZoneV1)
}

// Alias Function
func (s *SdaService) ReProvisionWiredDevice(requestSdaReProvisionWiredDeviceV1 *RequestSdaReProvisionWiredDeviceV1) (*ResponseSdaReProvisionWiredDeviceV1, *resty.Response, error) {
	return s.ReProvisionWiredDeviceV1(requestSdaReProvisionWiredDeviceV1)
}

// Alias Function
func (s *SdaService) DeletePortAssignmentByID(id string) (*ResponseSdaDeletePortAssignmentByIDV1, *resty.Response, error) {
	return s.DeletePortAssignmentByIDV1(id)
}

// Alias Function
func (s *SdaService) DeleteFabricDeviceLayer3HandoffsWithIPTransit(DeleteFabricDeviceLayer3HandoffsWithIpTransitV1QueryParams *DeleteFabricDeviceLayer3HandoffsWithIPTransitV1QueryParams) (*ResponseSdaDeleteFabricDeviceLayer3HandoffsWithIPTransitV1, *resty.Response, error) {
	return s.DeleteFabricDeviceLayer3HandoffsWithIPTransitV1(DeleteFabricDeviceLayer3HandoffsWithIpTransitV1QueryParams)
}

// Alias Function
func (s *SdaService) AddTransitPeerNetwork(requestSdaAddTransitPeerNetworkV1 *RequestSdaAddTransitPeerNetworkV1) (*ResponseSdaAddTransitPeerNetworkV1, *resty.Response, error) {
	return s.AddTransitPeerNetworkV1(requestSdaAddTransitPeerNetworkV1)
}

// Alias Function
func (s *SdaService) GetControlPlaneDeviceFromSdaFabric(GetControlPlaneDeviceFromSDAFabricV1QueryParams *GetControlPlaneDeviceFromSdaFabricV1QueryParams) (*ResponseSdaGetControlPlaneDeviceFromSdaFabricV1, *resty.Response, error) {
	return s.GetControlPlaneDeviceFromSdaFabricV1(GetControlPlaneDeviceFromSDAFabricV1QueryParams)
}

// Alias Function
func (s *SdaService) DeleteLayer2VirtualNetworks(DeleteLayer2VirtualNetworksV1QueryParams *DeleteLayer2VirtualNetworksV1QueryParams) (*ResponseSdaDeleteLayer2VirtualNetworksV1, *resty.Response, error) {
	return s.DeleteLayer2VirtualNetworksV1(DeleteLayer2VirtualNetworksV1QueryParams)
}

// Alias Function
func (s *SdaService) GetFabricZoneCount() (*ResponseSdaGetFabricZoneCountV1, *resty.Response, error) {
	return s.GetFabricZoneCountV1()
}

// Alias Function
func (s *SdaService) AddPortAssignmentForUserDeviceInSdaFabric(requestSdaAddPortAssignmentForUserDeviceInSDAFabricV1 *RequestSdaAddPortAssignmentForUserDeviceInSdaFabricV1) (*ResponseSdaAddPortAssignmentForUserDeviceInSdaFabricV1, *resty.Response, error) {
	return s.AddPortAssignmentForUserDeviceInSdaFabricV1(requestSdaAddPortAssignmentForUserDeviceInSDAFabricV1)
}

// Alias Function
func (s *SdaService) UpdatePortChannels(requestSdaUpdatePortChannelsV1 *RequestSdaUpdatePortChannelsV1) (*ResponseSdaUpdatePortChannelsV1, *resty.Response, error) {
	return s.UpdatePortChannelsV1(requestSdaUpdatePortChannelsV1)
}

// Alias Function
func (s *SdaService) GetVnFromSdaFabric(GetVNFromSDAFabricV1QueryParams *GetVnFromSdaFabricV1QueryParams) (*ResponseSdaGetVnFromSdaFabricV1, *resty.Response, error) {
	return s.GetVnFromSdaFabricV1(GetVNFromSDAFabricV1QueryParams)
}

// Alias Function
func (s *SdaService) GetFabricDevicesLayer3HandoffsWithIPTransit(GetFabricDevicesLayer3HandoffsWithIpTransitV1QueryParams *GetFabricDevicesLayer3HandoffsWithIPTransitV1QueryParams) (*ResponseSdaGetFabricDevicesLayer3HandoffsWithIPTransitV1, *resty.Response, error) {
	return s.GetFabricDevicesLayer3HandoffsWithIPTransitV1(GetFabricDevicesLayer3HandoffsWithIpTransitV1QueryParams)
}

// Alias Function
func (s *SdaService) AddFabricSite(requestSdaAddFabricSiteV1 *RequestSdaAddFabricSiteV1) (*ResponseSdaAddFabricSiteV1, *resty.Response, error) {
	return s.AddFabricSiteV1(requestSdaAddFabricSiteV1)
}

// Alias Function
func (s *SdaService) UpdateVirtualNetworkWithScalableGroups(requestSdaUpdateVirtualNetworkWithScalableGroupsV1 *RequestSdaUpdateVirtualNetworkWithScalableGroupsV1) (*ResponseSdaUpdateVirtualNetworkWithScalableGroupsV1, *resty.Response, error) {
	return s.UpdateVirtualNetworkWithScalableGroupsV1(requestSdaUpdateVirtualNetworkWithScalableGroupsV1)
}

// Alias Function
func (s *SdaService) GetAnycastGateways(GetAnycastGatewaysV1QueryParams *GetAnycastGatewaysV1QueryParams) (*ResponseSdaGetAnycastGatewaysV1, *resty.Response, error) {
	return s.GetAnycastGatewaysV1(GetAnycastGatewaysV1QueryParams)
}

// Alias Function
func (s *SdaService) ProvisionDevices(requestSdaProvisionDevicesV1 *RequestSdaProvisionDevicesV1) (*ResponseSdaProvisionDevicesV1, *resty.Response, error) {
	return s.ProvisionDevicesV1(requestSdaProvisionDevicesV1)
}

// Alias Function
func (s *SdaService) UpdateLayer3VirtualNetworks(requestSdaUpdateLayer3VirtualNetworksV1 *RequestSdaUpdateLayer3VirtualNetworksV1) (*ResponseSdaUpdateLayer3VirtualNetworksV1, *resty.Response, error) {
	return s.UpdateLayer3VirtualNetworksV1(requestSdaUpdateLayer3VirtualNetworksV1)
}

// Alias Function
func (s *SdaService) DeleteMulticastFromSdaFabric(DeleteMulticastFromSDAFabricV1QueryParams *DeleteMulticastFromSdaFabricV1QueryParams) (*ResponseSdaDeleteMulticastFromSdaFabricV1, *resty.Response, error) {
	return s.DeleteMulticastFromSdaFabricV1(DeleteMulticastFromSDAFabricV1QueryParams)
}

// Alias Function
func (s *SdaService) AddControlPlaneDeviceInSdaFabric(requestSdaAddControlPlaneDeviceInSDAFabricV1 *RequestSdaAddControlPlaneDeviceInSdaFabricV1) (*ResponseSdaAddControlPlaneDeviceInSdaFabricV1, *resty.Response, error) {
	return s.AddControlPlaneDeviceInSdaFabricV1(requestSdaAddControlPlaneDeviceInSDAFabricV1)
}

// Alias Function
func (s *SdaService) DeletePortAssignments(DeletePortAssignmentsV1QueryParams *DeletePortAssignmentsV1QueryParams) (*ResponseSdaDeletePortAssignmentsV1, *resty.Response, error) {
	return s.DeletePortAssignmentsV1(DeletePortAssignmentsV1QueryParams)
}

// Alias Function
func (s *SdaService) GetMulticastVirtualNetworkCount(GetMulticastVirtualNetworkCountV1QueryParams *GetMulticastVirtualNetworkCountV1QueryParams) (*ResponseSdaGetMulticastVirtualNetworkCountV1, *resty.Response, error) {
	return s.GetMulticastVirtualNetworkCountV1(GetMulticastVirtualNetworkCountV1QueryParams)
}

// Alias Function
func (s *SdaService) UpdateDefaultAuthenticationProfileInSdaFabric(requestSdaUpdateDefaultAuthenticationProfileInSDAFabricV1 *RequestSdaUpdateDefaultAuthenticationProfileInSdaFabricV1) (*ResponseSdaUpdateDefaultAuthenticationProfileInSdaFabricV1, *resty.Response, error) {
	return s.UpdateDefaultAuthenticationProfileInSdaFabricV1(requestSdaUpdateDefaultAuthenticationProfileInSDAFabricV1)
}

// Alias Function
func (s *SdaService) GetFabricDevicesLayer3HandoffsWithSdaTransitCount(GetFabricDevicesLayer3HandoffsWithSdaTransitCountV1QueryParams *GetFabricDevicesLayer3HandoffsWithSdaTransitCountV1QueryParams) (*ResponseSdaGetFabricDevicesLayer3HandoffsWithSdaTransitCountV1, *resty.Response, error) {
	return s.GetFabricDevicesLayer3HandoffsWithSdaTransitCountV1(GetFabricDevicesLayer3HandoffsWithSdaTransitCountV1QueryParams)
}

// Alias Function
func (s *SdaService) AddFabricDevices(requestSdaAddFabricDevicesV1 *RequestSdaAddFabricDevicesV1) (*ResponseSdaAddFabricDevicesV1, *resty.Response, error) {
	return s.AddFabricDevicesV1(requestSdaAddFabricDevicesV1)
}

// Alias Function
func (s *SdaService) GetProvisionedDevices(GetProvisionedDevicesV1QueryParams *GetProvisionedDevicesV1QueryParams) (*ResponseSdaGetProvisionedDevicesV1, *resty.Response, error) {
	return s.GetProvisionedDevicesV1(GetProvisionedDevicesV1QueryParams)
}

// Alias Function
func (s *SdaService) GetLayer2VirtualNetworkCount(GetLayer2VirtualNetworkCountV1QueryParams *GetLayer2VirtualNetworkCountV1QueryParams) (*ResponseSdaGetLayer2VirtualNetworkCountV1, *resty.Response, error) {
	return s.GetLayer2VirtualNetworkCountV1(GetLayer2VirtualNetworkCountV1QueryParams)
}

// Alias Function
func (s *SdaService) GetAnycastGatewayCount(GetAnycastGatewayCountV1QueryParams *GetAnycastGatewayCountV1QueryParams) (*ResponseSdaGetAnycastGatewayCountV1, *resty.Response, error) {
	return s.GetAnycastGatewayCountV1(GetAnycastGatewayCountV1QueryParams)
}

// Alias Function
func (s *SdaService) GetPortChannelCount(GetPortChannelCountV1QueryParams *GetPortChannelCountV1QueryParams) (*ResponseSdaGetPortChannelCountV1, *resty.Response, error) {
	return s.GetPortChannelCountV1(GetPortChannelCountV1QueryParams)
}

// Alias Function
func (s *SdaService) GetEdgeDeviceFromSdaFabric(GetEdgeDeviceFromSDAFabricV1QueryParams *GetEdgeDeviceFromSdaFabricV1QueryParams) (*ResponseSdaGetEdgeDeviceFromSdaFabricV1, *resty.Response, error) {
	return s.GetEdgeDeviceFromSdaFabricV1(GetEdgeDeviceFromSDAFabricV1QueryParams)
}

// Alias Function
func (s *SdaService) GetFabricSiteCount() (*ResponseSdaGetFabricSiteCountV1, *resty.Response, error) {
	return s.GetFabricSiteCountV1()
}

// Alias Function
func (s *SdaService) ReProvisionDevices(requestSdaReProvisionDevicesV1 *RequestSdaReProvisionDevicesV1) (*ResponseSdaReProvisionDevicesV1, *resty.Response, error) {
	return s.ReProvisionDevicesV1(requestSdaReProvisionDevicesV1)
}

// Alias Function
func (s *SdaService) ProvisionWiredDevice(requestSdaProvisionWiredDeviceV1 *RequestSdaProvisionWiredDeviceV1) (*ResponseSdaProvisionWiredDeviceV1, *resty.Response, error) {
	return s.ProvisionWiredDeviceV1(requestSdaProvisionWiredDeviceV1)
}

// Alias Function
func (s *SdaService) GetProvisionedWiredDevice(GetProvisionedWiredDeviceV1QueryParams *GetProvisionedWiredDeviceV1QueryParams) (*ResponseSdaGetProvisionedWiredDeviceV1, *resty.Response, error) {
	return s.GetProvisionedWiredDeviceV1(GetProvisionedWiredDeviceV1QueryParams)
}

// Alias Function
func (s *SdaService) GetPortAssignmentForUserDeviceInSdaFabric(GetPortAssignmentForUserDeviceInSDAFabricV1QueryParams *GetPortAssignmentForUserDeviceInSdaFabricV1QueryParams) (*ResponseSdaGetPortAssignmentForUserDeviceInSdaFabricV1, *resty.Response, error) {
	return s.GetPortAssignmentForUserDeviceInSdaFabricV1(GetPortAssignmentForUserDeviceInSDAFabricV1QueryParams)
}

// Alias Function
func (s *SdaService) DeleteDefaultAuthenticationProfileFromSdaFabric(DeleteDefaultAuthenticationProfileFromSDAFabricV1QueryParams *DeleteDefaultAuthenticationProfileFromSdaFabricV1QueryParams) (*ResponseSdaDeleteDefaultAuthenticationProfileFromSdaFabricV1, *resty.Response, error) {
	return s.DeleteDefaultAuthenticationProfileFromSdaFabricV1(DeleteDefaultAuthenticationProfileFromSDAFabricV1QueryParams)
}

// Alias Function
func (s *SdaService) UpdateFabricSite(requestSdaUpdateFabricSiteV1 *RequestSdaUpdateFabricSiteV1) (*ResponseSdaUpdateFabricSiteV1, *resty.Response, error) {
	return s.UpdateFabricSiteV1(requestSdaUpdateFabricSiteV1)
}

// Alias Function
func (s *SdaService) AddSiteInSdaFabric(requestSdaAddSiteInSDAFabricV1 *RequestSdaAddSiteInSdaFabricV1) (*ResponseSdaAddSiteInSdaFabricV1, *resty.Response, error) {
	return s.AddSiteInSdaFabricV1(requestSdaAddSiteInSDAFabricV1)
}

// Alias Function
func (s *SdaService) DeleteBorderDeviceFromSdaFabric(DeleteBorderDeviceFromSDAFabricV1QueryParams *DeleteBorderDeviceFromSdaFabricV1QueryParams) (*ResponseSdaDeleteBorderDeviceFromSdaFabricV1, *resty.Response, error) {
	return s.DeleteBorderDeviceFromSdaFabricV1(DeleteBorderDeviceFromSDAFabricV1QueryParams)
}

// Alias Function
func (s *SdaService) GetLayer2VirtualNetworks(GetLayer2VirtualNetworksV1QueryParams *GetLayer2VirtualNetworksV1QueryParams) (*ResponseSdaGetLayer2VirtualNetworksV1, *resty.Response, error) {
	return s.GetLayer2VirtualNetworksV1(GetLayer2VirtualNetworksV1QueryParams)
}

// Alias Function
func (s *SdaService) GetExtranetPolicies(GetExtranetPoliciesV1QueryParams *GetExtranetPoliciesV1QueryParams) (*ResponseSdaGetExtranetPoliciesV1, *resty.Response, error) {
	return s.GetExtranetPoliciesV1(GetExtranetPoliciesV1QueryParams)
}

// Alias Function
func (s *SdaService) GetPortAssignmentForAccessPointInSdaFabric(GetPortAssignmentForAccessPointInSDAFabricV1QueryParams *GetPortAssignmentForAccessPointInSdaFabricV1QueryParams) (*ResponseSdaGetPortAssignmentForAccessPointInSdaFabricV1, *resty.Response, error) {
	return s.GetPortAssignmentForAccessPointInSdaFabricV1(GetPortAssignmentForAccessPointInSDAFabricV1QueryParams)
}

// Alias Function
func (s *SdaService) DeletePortChannels(DeletePortChannelsV1QueryParams *DeletePortChannelsV1QueryParams) (*ResponseSdaDeletePortChannelsV1, *resty.Response, error) {
	return s.DeletePortChannelsV1(DeletePortChannelsV1QueryParams)
}

// Alias Function
func (s *SdaService) AddVnInFabric(requestSdaAddVNInFabricV1 *RequestSdaAddVnInFabricV1) (*ResponseSdaAddVnInFabricV1, *resty.Response, error) {
	return s.AddVnInFabricV1(requestSdaAddVNInFabricV1)
}

// Alias Function
func (s *SdaService) UpdateLayer2VirtualNetworks(requestSdaUpdateLayer2VirtualNetworksV1 *RequestSdaUpdateLayer2VirtualNetworksV1) (*ResponseSdaUpdateLayer2VirtualNetworksV1, *resty.Response, error) {
	return s.UpdateLayer2VirtualNetworksV1(requestSdaUpdateLayer2VirtualNetworksV1)
}
