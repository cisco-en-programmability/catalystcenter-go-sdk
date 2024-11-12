package catalyst

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type DeviceOnboardingPnpService service

type GetDeviceListSiteManagementV1QueryParams struct {
	Limit            int      `url:"limit,omitempty"`            //Limits number of results
	Offset           int      `url:"offset,omitempty"`           //Index of first result
	Sort             []string `url:"sort,omitempty"`             //Comma seperated list of fields to sort on
	SortOrder        string   `url:"sortOrder,omitempty"`        //Sort Order Ascending (asc) or Descending (des)
	SerialNumber     []string `url:"serialNumber,omitempty"`     //Device Serial Number
	State            []string `url:"state,omitempty"`            //Device State
	OnbState         []string `url:"onbState,omitempty"`         //Device Onboarding State
	Name             []string `url:"name,omitempty"`             //Device Name
	Pid              []string `url:"pid,omitempty"`              //Device ProductId
	Source           []string `url:"source,omitempty"`           //Device Source
	WorkflowID       []string `url:"workflowId,omitempty"`       //Device Workflow Id
	WorkflowName     []string `url:"workflowName,omitempty"`     //Device Workflow Name
	SmartAccountID   []string `url:"smartAccountId,omitempty"`   //Device Smart Account
	VirtualAccountID []string `url:"virtualAccountId,omitempty"` //Device Virtual Account
	LastContact      bool     `url:"lastContact,omitempty"`      //Device Has Contacted lastContact > 0
	MacAddress       string   `url:"macAddress,omitempty"`       //Device Mac Address
	Hostname         string   `url:"hostname,omitempty"`         //Device Hostname
	SiteName         string   `url:"siteName,omitempty"`         //Device Site Name
}
type GetDeviceCountV1QueryParams struct {
	SerialNumber     []string `url:"serialNumber,omitempty"`     //Device Serial Number
	State            []string `url:"state,omitempty"`            //Device State
	OnbState         []string `url:"onbState,omitempty"`         //Device Onboarding State
	Name             []string `url:"name,omitempty"`             //Device Name
	Pid              []string `url:"pid,omitempty"`              //Device ProductId
	Source           []string `url:"source,omitempty"`           //Device Source
	WorkflowID       []string `url:"workflowId,omitempty"`       //Device Workflow Id
	WorkflowName     []string `url:"workflowName,omitempty"`     //Device Workflow Name
	SmartAccountID   []string `url:"smartAccountId,omitempty"`   //Device Smart Account
	VirtualAccountID []string `url:"virtualAccountId,omitempty"` //Device Virtual Account
	LastContact      bool     `url:"lastContact,omitempty"`      //Device Has Contacted lastContact > 0
}
type GetDeviceHistoryV1QueryParams struct {
	SerialNumber string   `url:"serialNumber,omitempty"` //Device Serial Number
	Sort         []string `url:"sort,omitempty"`         //Comma seperated list of fields to sort on
	SortOrder    string   `url:"sortOrder,omitempty"`    //Sort Order Ascending (asc) or Descending (des)
}
type DeregisterVirtualAccountV1QueryParams struct {
	Domain string `url:"domain,omitempty"` //Smart Account Domain
	Name   string `url:"name,omitempty"`   //Virtual Account Name
}
type GetWorkflowsV1QueryParams struct {
	Limit     int      `url:"limit,omitempty"`     //Limits number of results
	Offset    int      `url:"offset,omitempty"`    //Index of first result
	Sort      []string `url:"sort,omitempty"`      //Comma seperated lost of fields to sort on
	SortOrder string   `url:"sortOrder,omitempty"` //Sort Order Ascending (asc) or Descending (des)
	Type      []string `url:"type,omitempty"`      //Workflow Type
	Name      []string `url:"name,omitempty"`      //Workflow Name
}
type GetWorkflowCountV1QueryParams struct {
	Name []string `url:"name,omitempty"` //Workflow Name
}

type ResponseDeviceOnboardingPnpAuthorizeDeviceV1 struct {
	JSONResponse      *ResponseDeviceOnboardingPnpAuthorizeDeviceV1JSONResponse `json:"jsonResponse,omitempty"`      //
	Message           string                                                    `json:"message,omitempty"`           // Message
	StatusCode        *float64                                                  `json:"statusCode,omitempty"`        // Status Code
	JSONArrayResponse []string                                                  `json:"jsonArrayResponse,omitempty"` // Json Array Response
}
type ResponseDeviceOnboardingPnpAuthorizeDeviceV1JSONResponse struct {
	Empty *bool `json:"empty,omitempty"` // Empty
}
type ResponseDeviceOnboardingPnpAddDeviceV1 struct {
	TypeID               string                                                      `json:"_id,omitempty"`                  // Id
	DeviceInfo           *ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfo           `json:"deviceInfo,omitempty"`           //
	SystemResetWorkflow  *ResponseDeviceOnboardingPnpAddDeviceV1SystemResetWorkflow  `json:"systemResetWorkflow,omitempty"`  //
	SystemWorkflow       *ResponseDeviceOnboardingPnpAddDeviceV1SystemWorkflow       `json:"systemWorkflow,omitempty"`       //
	Workflow             *ResponseDeviceOnboardingPnpAddDeviceV1Workflow             `json:"workflow,omitempty"`             //
	RunSummaryList       *[]ResponseDeviceOnboardingPnpAddDeviceV1RunSummaryList     `json:"runSummaryList,omitempty"`       //
	WorkflowParameters   *ResponseDeviceOnboardingPnpAddDeviceV1WorkflowParameters   `json:"workflowParameters,omitempty"`   //
	DayZeroConfig        *ResponseDeviceOnboardingPnpAddDeviceV1DayZeroConfig        `json:"dayZeroConfig,omitempty"`        //
	DayZeroConfigPreview *ResponseDeviceOnboardingPnpAddDeviceV1DayZeroConfigPreview `json:"dayZeroConfigPreview,omitempty"` // Day Zero Config Preview
	Version              *float64                                                    `json:"version,omitempty"`              // Version
	TenantID             string                                                      `json:"tenantId,omitempty"`             // Tenant Id
}
type ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfo struct {
	Source                    string                                                                  `json:"source,omitempty"`                    // Source
	SerialNumber              string                                                                  `json:"serialNumber,omitempty"`              // Serial Number
	Stack                     *bool                                                                   `json:"stack,omitempty"`                     // Stack
	Mode                      string                                                                  `json:"mode,omitempty"`                      // Mode
	State                     string                                                                  `json:"state,omitempty"`                     // State
	Location                  *ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoLocation               `json:"location,omitempty"`                  //
	Description               string                                                                  `json:"description,omitempty"`               // Description
	OnbState                  string                                                                  `json:"onbState,omitempty"`                  // Onb State
	AuthenticatedMicNumber    string                                                                  `json:"authenticatedMicNumber,omitempty"`    // Authenticated Mic Number
	AuthenticatedSudiSerialNo string                                                                  `json:"authenticatedSudiSerialNo,omitempty"` // Authenticated Sudi Serial No
	CapabilitiesSupported     []string                                                                `json:"capabilitiesSupported,omitempty"`     // Capabilities Supported
	FeaturesSupported         []string                                                                `json:"featuresSupported,omitempty"`         // Features Supported
	CmState                   string                                                                  `json:"cmState,omitempty"`                   // Cm State
	FirstContact              *float64                                                                `json:"firstContact,omitempty"`              // First Contact
	LastContact               *float64                                                                `json:"lastContact,omitempty"`               // Last Contact
	MacAddress                string                                                                  `json:"macAddress,omitempty"`                // Mac Address
	Pid                       string                                                                  `json:"pid,omitempty"`                       // Pid
	DeviceSudiSerialNos       []string                                                                `json:"deviceSudiSerialNos,omitempty"`       // Device Sudi Serial Nos
	LastUpdateOn              *float64                                                                `json:"lastUpdateOn,omitempty"`              // Last Update On
	WorkflowID                string                                                                  `json:"workflowId,omitempty"`                // Workflow Id
	WorkflowName              string                                                                  `json:"workflowName,omitempty"`              // Workflow Name
	ProjectID                 string                                                                  `json:"projectId,omitempty"`                 // Project Id
	ProjectName               string                                                                  `json:"projectName,omitempty"`               // Project Name
	DeviceType                string                                                                  `json:"deviceType,omitempty"`                // Device Type
	AgentType                 string                                                                  `json:"agentType,omitempty"`                 // Agent Type
	ImageVersion              string                                                                  `json:"imageVersion,omitempty"`              // Image Version
	FileSystemList            *[]ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoFileSystemList       `json:"fileSystemList,omitempty"`            //
	PnpProfileList            *[]ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoPnpProfileList       `json:"pnpProfileList,omitempty"`            //
	ImageFile                 string                                                                  `json:"imageFile,omitempty"`                 // Image File
	HTTPHeaders               *[]ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoHTTPHeaders          `json:"httpHeaders,omitempty"`               //
	NeighborLinks             *[]ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoNeighborLinks        `json:"neighborLinks,omitempty"`             //
	LastSyncTime              *float64                                                                `json:"lastSyncTime,omitempty"`              // Last Sync Time
	IPInterfaces              *[]ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoIPInterfaces         `json:"ipInterfaces,omitempty"`              //
	Hostname                  string                                                                  `json:"hostname,omitempty"`                  // Hostname
	AuthStatus                string                                                                  `json:"authStatus,omitempty"`                // Auth Status
	StackInfo                 *ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoStackInfo              `json:"stackInfo,omitempty"`                 //
	ReloadRequested           *bool                                                                   `json:"reloadRequested,omitempty"`           // Reload Requested
	AddedOn                   *float64                                                                `json:"addedOn,omitempty"`                   // Added On
	SiteID                    string                                                                  `json:"siteId,omitempty"`                    // Site Id
	AAACredentials            *ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoAAACredentials         `json:"aaaCredentials,omitempty"`            //
	UserMicNumbers            []string                                                                `json:"userMicNumbers,omitempty"`            // User Mic Numbers
	UserSudiSerialNos         []string                                                                `json:"userSudiSerialNos,omitempty"`         // User Sudi Serial Nos
	AddnMacAddrs              []string                                                                `json:"addnMacAddrs,omitempty"`              // Addn Mac Addrs
	PreWorkflowCliOuputs      *[]ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoPreWorkflowCliOuputs `json:"preWorkflowCliOuputs,omitempty"`      //
	Tags                      *ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoTags                   `json:"tags,omitempty"`                      // Tags
	SudiRequired              *bool                                                                   `json:"sudiRequired,omitempty"`              // Sudi Required
	SmartAccountID            string                                                                  `json:"smartAccountId,omitempty"`            // Smart Account Id
	VirtualAccountID          string                                                                  `json:"virtualAccountId,omitempty"`          // Virtual Account Id
	PopulateInventory         *bool                                                                   `json:"populateInventory,omitempty"`         // Populate Inventory
	SiteName                  string                                                                  `json:"siteName,omitempty"`                  // Site Name
	Name                      string                                                                  `json:"name,omitempty"`                      // Name
}
type ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoLocation struct {
	SiteID    string `json:"siteId,omitempty"`    // Site Id
	Address   string `json:"address,omitempty"`   // Address
	Latitude  string `json:"latitude,omitempty"`  // Latitude
	Longitude string `json:"longitude,omitempty"` // Longitude
	Altitude  string `json:"altitude,omitempty"`  // Altitude
}
type ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoFileSystemList struct {
	Type      string   `json:"type,omitempty"`      // Type
	Writeable *bool    `json:"writeable,omitempty"` // Writeable
	Freespace *float64 `json:"freespace,omitempty"` // Freespace
	Name      string   `json:"name,omitempty"`      // Name
	Readable  *bool    `json:"readable,omitempty"`  // Readable
	Size      *float64 `json:"size,omitempty"`      // Size
}
type ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoPnpProfileList struct {
	ProfileName       string                                                                           `json:"profileName,omitempty"`       // Profile Name
	DiscoveryCreated  *bool                                                                            `json:"discoveryCreated,omitempty"`  // Discovery Created
	CreatedBy         string                                                                           `json:"createdBy,omitempty"`         // Created By
	PrimaryEndpoint   *ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoPnpProfileListPrimaryEndpoint   `json:"primaryEndpoint,omitempty"`   //
	SecondaryEndpoint *ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoPnpProfileListSecondaryEndpoint `json:"secondaryEndpoint,omitempty"` //
}
type ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoPnpProfileListPrimaryEndpoint struct {
	Port        *float64                                                                                  `json:"port,omitempty"`        // Port
	Protocol    string                                                                                    `json:"protocol,omitempty"`    // Protocol
	IPv4Address *ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoPnpProfileListPrimaryEndpointIPv4Address `json:"ipv4Address,omitempty"` // Ipv4 Address
	IPv6Address *ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoPnpProfileListPrimaryEndpointIPv6Address `json:"ipv6Address,omitempty"` // Ipv6 Address
	Fqdn        string                                                                                    `json:"fqdn,omitempty"`        // Fqdn
	Certificate string                                                                                    `json:"certificate,omitempty"` // Certificate
}
type ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoPnpProfileListPrimaryEndpointIPv4Address interface{}
type ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoPnpProfileListPrimaryEndpointIPv6Address interface{}
type ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoPnpProfileListSecondaryEndpoint struct {
	Port        *float64                                                                                    `json:"port,omitempty"`        // Port
	Protocol    string                                                                                      `json:"protocol,omitempty"`    // Protocol
	IPv4Address *ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoPnpProfileListSecondaryEndpointIPv4Address `json:"ipv4Address,omitempty"` // Ipv4 Address
	IPv6Address *ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoPnpProfileListSecondaryEndpointIPv6Address `json:"ipv6Address,omitempty"` // Ipv6 Address
	Fqdn        string                                                                                      `json:"fqdn,omitempty"`        // Fqdn
	Certificate string                                                                                      `json:"certificate,omitempty"` // Certificate
}
type ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoPnpProfileListSecondaryEndpointIPv4Address interface{}
type ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoPnpProfileListSecondaryEndpointIPv6Address interface{}
type ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoHTTPHeaders struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoNeighborLinks struct {
	LocalInterfaceName       string `json:"localInterfaceName,omitempty"`       // Local Interface Name
	LocalShortInterfaceName  string `json:"localShortInterfaceName,omitempty"`  // Local Short Interface Name
	LocalMacAddress          string `json:"localMacAddress,omitempty"`          // Local Mac Address
	RemoteInterfaceName      string `json:"remoteInterfaceName,omitempty"`      // Remote Interface Name
	RemoteShortInterfaceName string `json:"remoteShortInterfaceName,omitempty"` // Remote Short Interface Name
	RemoteMacAddress         string `json:"remoteMacAddress,omitempty"`         // Remote Mac Address
	RemoteDeviceName         string `json:"remoteDeviceName,omitempty"`         // Remote Device Name
	RemotePlatform           string `json:"remotePlatform,omitempty"`           // Remote Platform
	RemoteVersion            string `json:"remoteVersion,omitempty"`            // Remote Version
}
type ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoIPInterfaces struct {
	Status          string                                                                         `json:"status,omitempty"`          // Status
	MacAddress      string                                                                         `json:"macAddress,omitempty"`      // Mac Address
	IPv4Address     *ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoIPInterfacesIPv4Address       `json:"ipv4Address,omitempty"`     // Ipv4 Address
	IPv6AddressList *[]ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoIPInterfacesIPv6AddressList `json:"ipv6AddressList,omitempty"` // Ipv6 Address List
	Name            string                                                                         `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoIPInterfacesIPv4Address interface{}
type ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoIPInterfacesIPv6AddressList interface{}
type ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoStackInfo struct {
	SupportsStackWorkflows *bool                                                                       `json:"supportsStackWorkflows,omitempty"` // Supports Stack Workflows
	IsFullRing             *bool                                                                       `json:"isFullRing,omitempty"`             // Is Full Ring
	StackMemberList        *[]ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoStackInfoStackMemberList `json:"stackMemberList,omitempty"`        //
	StackRingProtocol      string                                                                      `json:"stackRingProtocol,omitempty"`      // Stack Ring Protocol
	ValidLicenseLevels     []string                                                                    `json:"validLicenseLevels,omitempty"`     // Valid License Levels
	TotalMemberCount       *float64                                                                    `json:"totalMemberCount,omitempty"`       // Total Member Count
}
type ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoStackInfoStackMemberList struct {
	SerialNumber     string   `json:"serialNumber,omitempty"`     // Serial Number
	State            string   `json:"state,omitempty"`            // State
	Role             string   `json:"role,omitempty"`             // Role
	MacAddress       string   `json:"macAddress,omitempty"`       // Mac Address
	Pid              string   `json:"pid,omitempty"`              // Pid
	LicenseLevel     string   `json:"licenseLevel,omitempty"`     // License Level
	LicenseType      string   `json:"licenseType,omitempty"`      // License Type
	SudiSerialNumber string   `json:"sudiSerialNumber,omitempty"` // Sudi Serial Number
	HardwareVersion  string   `json:"hardwareVersion,omitempty"`  // Hardware Version
	StackNumber      *float64 `json:"stackNumber,omitempty"`      // Stack Number
	SoftwareVersion  string   `json:"softwareVersion,omitempty"`  // Software Version
	Priority         *float64 `json:"priority,omitempty"`         // Priority
}
type ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoAAACredentials struct {
	Password string `json:"password,omitempty"` // Password
	Username string `json:"username,omitempty"` // Username
}
type ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoPreWorkflowCliOuputs struct {
	Cli       string `json:"cli,omitempty"`       // Cli
	CliOutput string `json:"cliOutput,omitempty"` // Cli Output
}
type ResponseDeviceOnboardingPnpAddDeviceV1DeviceInfoTags interface{}
type ResponseDeviceOnboardingPnpAddDeviceV1SystemResetWorkflow struct {
	TypeID         string                                                            `json:"_id,omitempty"`            // Id
	State          string                                                            `json:"state,omitempty"`          // State
	Type           string                                                            `json:"type,omitempty"`           // Type
	Description    string                                                            `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                                          `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                            `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                                          `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                                          `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpAddDeviceV1SystemResetWorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                             `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                            `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                                          `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                                          `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                                          `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                            `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                            `json:"configId,omitempty"`       // Config Id
	Name           string                                                            `json:"name,omitempty"`           // Name
	Version        *float64                                                          `json:"version,omitempty"`        // Version
	TenantID       string                                                            `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpAddDeviceV1SystemResetWorkflowTasks struct {
	State           string                                                                        `json:"state,omitempty"`           // State
	Type            string                                                                        `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                                      `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                                      `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                                      `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                                      `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpAddDeviceV1SystemResetWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                                      `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                                        `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpAddDeviceV1SystemResetWorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpAddDeviceV1SystemWorkflow struct {
	TypeID         string                                                       `json:"_id,omitempty"`            // Id
	State          string                                                       `json:"state,omitempty"`          // State
	Type           string                                                       `json:"type,omitempty"`           // Type
	Description    string                                                       `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                                     `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                       `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                                     `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                                     `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpAddDeviceV1SystemWorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                        `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                       `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                                     `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                                     `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                                     `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                       `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                       `json:"configId,omitempty"`       // Config Id
	Name           string                                                       `json:"name,omitempty"`           // Name
	Version        *float64                                                     `json:"version,omitempty"`        // Version
	TenantID       string                                                       `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpAddDeviceV1SystemWorkflowTasks struct {
	State           string                                                                   `json:"state,omitempty"`           // State
	Type            string                                                                   `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                                 `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                                 `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                                 `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                                 `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpAddDeviceV1SystemWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                                 `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                                   `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpAddDeviceV1SystemWorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpAddDeviceV1Workflow struct {
	TypeID         string                                                 `json:"_id,omitempty"`            // Id
	State          string                                                 `json:"state,omitempty"`          // State
	Type           string                                                 `json:"type,omitempty"`           // Type
	Description    string                                                 `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                               `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                 `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                               `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                               `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpAddDeviceV1WorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                  `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                 `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                               `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                               `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                               `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                 `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                 `json:"configId,omitempty"`       // Config Id
	Name           string                                                 `json:"name,omitempty"`           // Name
	Version        *float64                                               `json:"version,omitempty"`        // Version
	TenantID       string                                                 `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpAddDeviceV1WorkflowTasks struct {
	State           string                                                             `json:"state,omitempty"`           // State
	Type            string                                                             `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                           `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                           `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                           `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                           `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpAddDeviceV1WorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                           `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                             `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpAddDeviceV1WorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpAddDeviceV1RunSummaryList struct {
	Details         string                                                               `json:"details,omitempty"`         // Details
	HistoryTaskInfo *ResponseDeviceOnboardingPnpAddDeviceV1RunSummaryListHistoryTaskInfo `json:"historyTaskInfo,omitempty"` //
	ErrorFlag       *bool                                                                `json:"errorFlag,omitempty"`       // Error Flag
	Timestamp       *float64                                                             `json:"timestamp,omitempty"`       // Timestamp
}
type ResponseDeviceOnboardingPnpAddDeviceV1RunSummaryListHistoryTaskInfo struct {
	Type         string                                                                             `json:"type,omitempty"`         // Type
	WorkItemList *[]ResponseDeviceOnboardingPnpAddDeviceV1RunSummaryListHistoryTaskInfoWorkItemList `json:"workItemList,omitempty"` //
	TimeTaken    *float64                                                                           `json:"timeTaken,omitempty"`    // Time Taken
	AddnDetails  *[]ResponseDeviceOnboardingPnpAddDeviceV1RunSummaryListHistoryTaskInfoAddnDetails  `json:"addnDetails,omitempty"`  //
	Name         string                                                                             `json:"name,omitempty"`         // Name
}
type ResponseDeviceOnboardingPnpAddDeviceV1RunSummaryListHistoryTaskInfoWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpAddDeviceV1RunSummaryListHistoryTaskInfoAddnDetails struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseDeviceOnboardingPnpAddDeviceV1WorkflowParameters struct {
	TopOfStackSerialNumber string                                                                `json:"topOfStackSerialNumber,omitempty"` // Top Of Stack Serial Number
	LicenseLevel           string                                                                `json:"licenseLevel,omitempty"`           // License Level
	LicenseType            string                                                                `json:"licenseType,omitempty"`            // License Type
	ConfigList             *[]ResponseDeviceOnboardingPnpAddDeviceV1WorkflowParametersConfigList `json:"configList,omitempty"`             //
}
type ResponseDeviceOnboardingPnpAddDeviceV1WorkflowParametersConfigList struct {
	ConfigParameters *[]ResponseDeviceOnboardingPnpAddDeviceV1WorkflowParametersConfigListConfigParameters `json:"configParameters,omitempty"` //
	ConfigID         string                                                                                `json:"configId,omitempty"`         // Config Id
}
type ResponseDeviceOnboardingPnpAddDeviceV1WorkflowParametersConfigListConfigParameters struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseDeviceOnboardingPnpAddDeviceV1DayZeroConfig struct {
	Config string `json:"config,omitempty"` // Config
}
type ResponseDeviceOnboardingPnpAddDeviceV1DayZeroConfigPreview interface{}
type ResponseDeviceOnboardingPnpGetDeviceListSiteManagementV1 []ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1 // Array of ResponseDeviceOnboardingPnpGetDeviceListSiteManagementV1
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1 struct {
	DeviceInfo           *ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfo           `json:"deviceInfo,omitempty"`           //
	SystemResetWorkflow  *ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1SystemResetWorkflow  `json:"systemResetWorkflow,omitempty"`  //
	SystemWorkflow       *ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1SystemWorkflow       `json:"systemWorkflow,omitempty"`       //
	Workflow             *ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1Workflow             `json:"workflow,omitempty"`             //
	RunSummaryList       *[]ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1RunSummaryList     `json:"runSummaryList,omitempty"`       //
	WorkflowParameters   *ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1WorkflowParameters   `json:"workflowParameters,omitempty"`   //
	DayZeroConfig        *ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DayZeroConfig        `json:"dayZeroConfig,omitempty"`        //
	DayZeroConfigPreview *ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DayZeroConfigPreview `json:"dayZeroConfigPreview,omitempty"` // Day Zero Config Preview
	Version              *float64                                                                          `json:"version,omitempty"`              // Version
	TenantID             string                                                                            `json:"tenantId,omitempty"`             // Tenant Id
	Progress             *ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1Progress             `json:"progress,omitempty"`             //
	ID                   string                                                                            `json:"id,omitempty"`                   // Id
}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfo struct {
	Source                    string                                                                                        `json:"source,omitempty"`                    // Source
	SerialNumber              string                                                                                        `json:"serialNumber,omitempty"`              // Serial Number
	Stack                     *bool                                                                                         `json:"stack,omitempty"`                     // Stack
	Mode                      string                                                                                        `json:"mode,omitempty"`                      // Mode
	State                     string                                                                                        `json:"state,omitempty"`                     // State
	Location                  *ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoLocation               `json:"location,omitempty"`                  //
	Description               string                                                                                        `json:"description,omitempty"`               // Description
	OnbState                  string                                                                                        `json:"onbState,omitempty"`                  // Onb State
	AuthenticatedMicNumber    string                                                                                        `json:"authenticatedMicNumber,omitempty"`    // Authenticated Mic Number
	AuthenticatedSudiSerialNo string                                                                                        `json:"authenticatedSudiSerialNo,omitempty"` // Authenticated Sudi Serial No
	CapabilitiesSupported     []string                                                                                      `json:"capabilitiesSupported,omitempty"`     // Capabilities Supported
	FeaturesSupported         []string                                                                                      `json:"featuresSupported,omitempty"`         // Features Supported
	CmState                   string                                                                                        `json:"cmState,omitempty"`                   // Cm State
	FirstContact              *float64                                                                                      `json:"firstContact,omitempty"`              // First Contact
	LastContact               *float64                                                                                      `json:"lastContact,omitempty"`               // Last Contact
	MacAddress                string                                                                                        `json:"macAddress,omitempty"`                // Mac Address
	Pid                       string                                                                                        `json:"pid,omitempty"`                       // Pid
	DeviceSudiSerialNos       []string                                                                                      `json:"deviceSudiSerialNos,omitempty"`       // Device Sudi Serial Nos
	LastUpdateOn              *float64                                                                                      `json:"lastUpdateOn,omitempty"`              // Last Update On
	WorkflowID                string                                                                                        `json:"workflowId,omitempty"`                // Workflow Id
	WorkflowName              string                                                                                        `json:"workflowName,omitempty"`              // Workflow Name
	ProjectID                 string                                                                                        `json:"projectId,omitempty"`                 // Project Id
	ProjectName               string                                                                                        `json:"projectName,omitempty"`               // Project Name
	DeviceType                string                                                                                        `json:"deviceType,omitempty"`                // Device Type
	AgentType                 string                                                                                        `json:"agentType,omitempty"`                 // Agent Type
	ImageVersion              string                                                                                        `json:"imageVersion,omitempty"`              // Image Version
	FileSystemList            *[]ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoFileSystemList       `json:"fileSystemList,omitempty"`            //
	PnpProfileList            *[]ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoPnpProfileList       `json:"pnpProfileList,omitempty"`            //
	ImageFile                 string                                                                                        `json:"imageFile,omitempty"`                 // Image File
	HTTPHeaders               *[]ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoHTTPHeaders          `json:"httpHeaders,omitempty"`               //
	NeighborLinks             *[]ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoNeighborLinks        `json:"neighborLinks,omitempty"`             //
	LastSyncTime              *float64                                                                                      `json:"lastSyncTime,omitempty"`              // Last Sync Time
	IPInterfaces              *[]ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoIPInterfaces         `json:"ipInterfaces,omitempty"`              //
	Hostname                  string                                                                                        `json:"hostname,omitempty"`                  // Hostname
	AuthStatus                string                                                                                        `json:"authStatus,omitempty"`                // Auth Status
	StackInfo                 *ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoStackInfo              `json:"stackInfo,omitempty"`                 //
	ReloadRequested           *bool                                                                                         `json:"reloadRequested,omitempty"`           // Reload Requested
	AddedOn                   *float64                                                                                      `json:"addedOn,omitempty"`                   // Added On
	SiteID                    string                                                                                        `json:"siteId,omitempty"`                    // Site Id
	AAACredentials            *ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoAAACredentials         `json:"aaaCredentials,omitempty"`            //
	UserMicNumbers            []string                                                                                      `json:"userMicNumbers,omitempty"`            // User Mic Numbers
	UserSudiSerialNos         []string                                                                                      `json:"userSudiSerialNos,omitempty"`         // User Sudi Serial Nos
	AddnMacAddrs              []string                                                                                      `json:"addnMacAddrs,omitempty"`              // Addn Mac Addrs
	PreWorkflowCliOuputs      *[]ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoPreWorkflowCliOuputs `json:"preWorkflowCliOuputs,omitempty"`      //
	Tags                      *ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoTags                   `json:"tags,omitempty"`                      // Tags
	SudiRequired              *bool                                                                                         `json:"sudiRequired,omitempty"`              // Sudi Required
	SmartAccountID            string                                                                                        `json:"smartAccountId,omitempty"`            // Smart Account Id
	VirtualAccountID          string                                                                                        `json:"virtualAccountId,omitempty"`          // Virtual Account Id
	PopulateInventory         *bool                                                                                         `json:"populateInventory,omitempty"`         // Populate Inventory
	SiteName                  string                                                                                        `json:"siteName,omitempty"`                  // Site Name
	Name                      string                                                                                        `json:"name,omitempty"`                      // Name
}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoLocation struct {
	SiteID    string `json:"siteId,omitempty"`    // Site Id
	Address   string `json:"address,omitempty"`   // Address
	Latitude  string `json:"latitude,omitempty"`  // Latitude
	Longitude string `json:"longitude,omitempty"` // Longitude
	Altitude  string `json:"altitude,omitempty"`  // Altitude
}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoFileSystemList struct {
	Type      string   `json:"type,omitempty"`      // Type
	Writeable *bool    `json:"writeable,omitempty"` // Writeable
	Freespace *float64 `json:"freespace,omitempty"` // Freespace
	Name      string   `json:"name,omitempty"`      // Name
	Readable  *bool    `json:"readable,omitempty"`  // Readable
	Size      *float64 `json:"size,omitempty"`      // Size
}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoPnpProfileList struct {
	ProfileName       string                                                                                                 `json:"profileName,omitempty"`       // Profile Name
	DiscoveryCreated  *bool                                                                                                  `json:"discoveryCreated,omitempty"`  // Discovery Created
	CreatedBy         string                                                                                                 `json:"createdBy,omitempty"`         // Created By
	PrimaryEndpoint   *ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoPnpProfileListPrimaryEndpoint   `json:"primaryEndpoint,omitempty"`   //
	SecondaryEndpoint *ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoPnpProfileListSecondaryEndpoint `json:"secondaryEndpoint,omitempty"` //
}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoPnpProfileListPrimaryEndpoint struct {
	Port        *float64                                                                                                        `json:"port,omitempty"`        // Port
	Protocol    string                                                                                                          `json:"protocol,omitempty"`    // Protocol
	IPv4Address *ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoPnpProfileListPrimaryEndpointIPv4Address `json:"ipv4Address,omitempty"` // Ipv4 Address
	IPv6Address *ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoPnpProfileListPrimaryEndpointIPv6Address `json:"ipv6Address,omitempty"` // Ipv6 Address
	Fqdn        string                                                                                                          `json:"fqdn,omitempty"`        // Fqdn
	Certificate string                                                                                                          `json:"certificate,omitempty"` // Certificate
}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoPnpProfileListPrimaryEndpointIPv4Address interface{}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoPnpProfileListPrimaryEndpointIPv6Address interface{}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoPnpProfileListSecondaryEndpoint struct {
	Port        *float64                                                                                                          `json:"port,omitempty"`        // Port
	Protocol    string                                                                                                            `json:"protocol,omitempty"`    // Protocol
	IPv4Address *ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoPnpProfileListSecondaryEndpointIPv4Address `json:"ipv4Address,omitempty"` // Ipv4 Address
	IPv6Address *ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoPnpProfileListSecondaryEndpointIPv6Address `json:"ipv6Address,omitempty"` // Ipv6 Address
	Fqdn        string                                                                                                            `json:"fqdn,omitempty"`        // Fqdn
	Certificate string                                                                                                            `json:"certificate,omitempty"` // Certificate
}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoPnpProfileListSecondaryEndpointIPv4Address interface{}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoPnpProfileListSecondaryEndpointIPv6Address interface{}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoHTTPHeaders struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoNeighborLinks struct {
	LocalInterfaceName       string `json:"localInterfaceName,omitempty"`       // Local Interface Name
	LocalShortInterfaceName  string `json:"localShortInterfaceName,omitempty"`  // Local Short Interface Name
	LocalMacAddress          string `json:"localMacAddress,omitempty"`          // Local Mac Address
	RemoteInterfaceName      string `json:"remoteInterfaceName,omitempty"`      // Remote Interface Name
	RemoteShortInterfaceName string `json:"remoteShortInterfaceName,omitempty"` // Remote Short Interface Name
	RemoteMacAddress         string `json:"remoteMacAddress,omitempty"`         // Remote Mac Address
	RemoteDeviceName         string `json:"remoteDeviceName,omitempty"`         // Remote Device Name
	RemotePlatform           string `json:"remotePlatform,omitempty"`           // Remote Platform
	RemoteVersion            string `json:"remoteVersion,omitempty"`            // Remote Version
}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoIPInterfaces struct {
	Status          string                                                                                               `json:"status,omitempty"`          // Status
	MacAddress      string                                                                                               `json:"macAddress,omitempty"`      // Mac Address
	IPv4Address     *ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoIPInterfacesIPv4Address       `json:"ipv4Address,omitempty"`     // Ipv4 Address
	IPv6AddressList *[]ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoIPInterfacesIPv6AddressList `json:"ipv6AddressList,omitempty"` // Ipv6 Address List
	Name            string                                                                                               `json:"name,omitempty"`            // Name
}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoIPInterfacesIPv4Address interface{}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoIPInterfacesIPv6AddressList interface{}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoStackInfo struct {
	SupportsStackWorkflows *bool                                                                                             `json:"supportsStackWorkflows,omitempty"` // Supports Stack Workflows
	IsFullRing             *bool                                                                                             `json:"isFullRing,omitempty"`             // Is Full Ring
	StackMemberList        *[]ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoStackInfoStackMemberList `json:"stackMemberList,omitempty"`        //
	StackRingProtocol      string                                                                                            `json:"stackRingProtocol,omitempty"`      // Stack Ring Protocol
	ValidLicenseLevels     []string                                                                                          `json:"validLicenseLevels,omitempty"`     // Valid License Levels
	TotalMemberCount       *float64                                                                                          `json:"totalMemberCount,omitempty"`       // Total Member Count
}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoStackInfoStackMemberList struct {
	SerialNumber     string   `json:"serialNumber,omitempty"`     // Serial Number
	State            string   `json:"state,omitempty"`            // State
	Role             string   `json:"role,omitempty"`             // Role
	MacAddress       string   `json:"macAddress,omitempty"`       // Mac Address
	Pid              string   `json:"pid,omitempty"`              // Pid
	LicenseLevel     string   `json:"licenseLevel,omitempty"`     // License Level
	LicenseType      string   `json:"licenseType,omitempty"`      // License Type
	SudiSerialNumber string   `json:"sudiSerialNumber,omitempty"` // Sudi Serial Number
	HardwareVersion  string   `json:"hardwareVersion,omitempty"`  // Hardware Version
	StackNumber      *float64 `json:"stackNumber,omitempty"`      // Stack Number
	SoftwareVersion  string   `json:"softwareVersion,omitempty"`  // Software Version
	Priority         *float64 `json:"priority,omitempty"`         // Priority
}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoAAACredentials struct {
	Password string `json:"password,omitempty"` // Password
	Username string `json:"username,omitempty"` // Username
}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoPreWorkflowCliOuputs struct {
	Cli       string `json:"cli,omitempty"`       // Cli
	CliOutput string `json:"cliOutput,omitempty"` // Cli Output
}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DeviceInfoTags interface{}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1SystemResetWorkflow struct {
	TypeID         string                                                                                  `json:"_id,omitempty"`            // Id
	State          string                                                                                  `json:"state,omitempty"`          // State
	Type           string                                                                                  `json:"type,omitempty"`           // Type
	Description    string                                                                                  `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                                                                `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                                                  `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                                                                `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                                                                `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1SystemResetWorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                                                   `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                                                  `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                                                                `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                                                                `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                                                                `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                                                  `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                                                  `json:"configId,omitempty"`       // Config Id
	Name           string                                                                                  `json:"name,omitempty"`           // Name
	Version        *float64                                                                                `json:"version,omitempty"`        // Version
	TenantID       string                                                                                  `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1SystemResetWorkflowTasks struct {
	State           string                                                                                              `json:"state,omitempty"`           // State
	Type            string                                                                                              `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                                                            `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                                                            `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                                                            `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                                                            `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1SystemResetWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                                                            `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                                                              `json:"name,omitempty"`            // Name
}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1SystemResetWorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1SystemWorkflow struct {
	TypeID         string                                                                             `json:"_id,omitempty"`            // Id
	State          string                                                                             `json:"state,omitempty"`          // State
	Type           string                                                                             `json:"type,omitempty"`           // Type
	Description    string                                                                             `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                                                           `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                                             `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                                                           `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                                                           `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1SystemWorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                                              `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                                             `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                                                           `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                                                           `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                                                           `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                                             `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                                             `json:"configId,omitempty"`       // Config Id
	Name           string                                                                             `json:"name,omitempty"`           // Name
	Version        *float64                                                                           `json:"version,omitempty"`        // Version
	TenantID       string                                                                             `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1SystemWorkflowTasks struct {
	State           string                                                                                         `json:"state,omitempty"`           // State
	Type            string                                                                                         `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                                                       `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                                                       `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                                                       `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                                                       `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1SystemWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                                                       `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                                                         `json:"name,omitempty"`            // Name
}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1SystemWorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1Workflow struct {
	TypeID         string                                                                       `json:"_id,omitempty"`            // Id
	State          string                                                                       `json:"state,omitempty"`          // State
	Type           string                                                                       `json:"type,omitempty"`           // Type
	Description    string                                                                       `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                                                     `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                                       `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                                                     `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                                                     `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1WorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                                        `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                                       `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                                                     `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                                                     `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                                                     `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                                       `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                                       `json:"configId,omitempty"`       // Config Id
	Name           string                                                                       `json:"name,omitempty"`           // Name
	Version        *float64                                                                     `json:"version,omitempty"`        // Version
	TenantID       string                                                                       `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1WorkflowTasks struct {
	State           string                                                                                   `json:"state,omitempty"`           // State
	Type            string                                                                                   `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                                                 `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                                                 `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                                                 `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                                                 `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1WorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                                                 `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                                                   `json:"name,omitempty"`            // Name
}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1WorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1RunSummaryList struct {
	Details         string                                                                                     `json:"details,omitempty"`         // Details
	HistoryTaskInfo *ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1RunSummaryListHistoryTaskInfo `json:"historyTaskInfo,omitempty"` //
	ErrorFlag       *bool                                                                                      `json:"errorFlag,omitempty"`       // Error Flag
	Timestamp       *float64                                                                                   `json:"timestamp,omitempty"`       // Timestamp
}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1RunSummaryListHistoryTaskInfo struct {
	Type         string                                                                                                   `json:"type,omitempty"`         // Type
	WorkItemList *[]ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1RunSummaryListHistoryTaskInfoWorkItemList `json:"workItemList,omitempty"` //
	TimeTaken    *float64                                                                                                 `json:"timeTaken,omitempty"`    // Time Taken
	AddnDetails  *[]ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1RunSummaryListHistoryTaskInfoAddnDetails  `json:"addnDetails,omitempty"`  //
	Name         string                                                                                                   `json:"name,omitempty"`         // Name
}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1RunSummaryListHistoryTaskInfoWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1RunSummaryListHistoryTaskInfoAddnDetails struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1WorkflowParameters struct {
	TopOfStackSerialNumber string                                                                                      `json:"topOfStackSerialNumber,omitempty"` // Top Of Stack Serial Number
	LicenseLevel           string                                                                                      `json:"licenseLevel,omitempty"`           // License Level
	LicenseType            string                                                                                      `json:"licenseType,omitempty"`            // License Type
	ConfigList             *[]ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1WorkflowParametersConfigList `json:"configList,omitempty"`             //
}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1WorkflowParametersConfigList struct {
	ConfigParameters *[]ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1WorkflowParametersConfigListConfigParameters `json:"configParameters,omitempty"` //
	ConfigID         string                                                                                                      `json:"configId,omitempty"`         // Config Id
}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1WorkflowParametersConfigListConfigParameters struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DayZeroConfig struct {
	Config string `json:"config,omitempty"` // Config
}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1DayZeroConfigPreview interface{}
type ResponseDeviceOnboardingPnpClaimDeviceV1 struct {
	JSONArrayResponse *[]ResponseDeviceOnboardingPnpClaimDeviceV1JSONArrayResponse `json:"jsonArrayResponse,omitempty"` // Json Array Response
	JSONResponse      *ResponseDeviceOnboardingPnpClaimDeviceV1JSONResponse        `json:"jsonResponse,omitempty"`      // Json Response
	Message           string                                                       `json:"message,omitempty"`           // Message
	StatusCode        *float64                                                     `json:"statusCode,omitempty"`        // Status Code
}
type ResponseItemDeviceOnboardingPnpGetDeviceListSiteManagementV1Progress struct {
	Message         string `json:"message,omitempty"`
	InProgress      bool   `json:"inProgress,omitempty"`
	ProgressPercent int    `json:"progressPercent,omitempty"`
}
type ResponseDeviceOnboardingPnpClaimDeviceV1JSONArrayResponse interface{}
type ResponseDeviceOnboardingPnpClaimDeviceV1JSONResponse interface{}
type ResponseDeviceOnboardingPnpGetDeviceCountV1 struct {
	Response *float64 `json:"response,omitempty"` // Response
}
type ResponseDeviceOnboardingPnpGetDeviceHistoryV1 struct {
	Response   *[]ResponseDeviceOnboardingPnpGetDeviceHistoryV1Response `json:"response,omitempty"`   //
	StatusCode *float64                                                 `json:"statusCode,omitempty"` // Status Code
}
type ResponseDeviceOnboardingPnpGetDeviceHistoryV1Response struct {
	Timestamp       *float64                                                              `json:"timestamp,omitempty"`       // Timestamp
	Details         string                                                                `json:"details,omitempty"`         // Details
	HistoryTaskInfo *ResponseDeviceOnboardingPnpGetDeviceHistoryV1ResponseHistoryTaskInfo `json:"historyTaskInfo,omitempty"` //
	ErrorFlag       *bool                                                                 `json:"errorFlag,omitempty"`       // Error Flag
}
type ResponseDeviceOnboardingPnpGetDeviceHistoryV1ResponseHistoryTaskInfo struct {
	Name         string                                                                              `json:"name,omitempty"`         // Name
	Type         string                                                                              `json:"type,omitempty"`         // Type
	TimeTaken    *float64                                                                            `json:"timeTaken,omitempty"`    // Time Taken
	WorkItemList *[]ResponseDeviceOnboardingPnpGetDeviceHistoryV1ResponseHistoryTaskInfoWorkItemList `json:"workItemList,omitempty"` //
	AddnDetails  *[]ResponseDeviceOnboardingPnpGetDeviceHistoryV1ResponseHistoryTaskInfoAddnDetails  `json:"addnDetails,omitempty"`  //
}
type ResponseDeviceOnboardingPnpGetDeviceHistoryV1ResponseHistoryTaskInfoWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
}
type ResponseDeviceOnboardingPnpGetDeviceHistoryV1ResponseHistoryTaskInfoAddnDetails struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1 struct {
	SuccessList *[]ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessList `json:"successList,omitempty"` //
	FailureList *[]ResponseDeviceOnboardingPnpImportDevicesInBulkV1FailureList `json:"failureList,omitempty"` //
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessList struct {
	ID                   string                                                                           `json:"id,omitempty"`                   // Id
	DeviceInfo           *ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfo           `json:"deviceInfo,omitempty"`           //
	SystemResetWorkflow  *ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListSystemResetWorkflow  `json:"systemResetWorkflow,omitempty"`  //
	SystemWorkflow       *ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListSystemWorkflow       `json:"systemWorkflow,omitempty"`       //
	Workflow             *ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListWorkflow             `json:"workflow,omitempty"`             //
	RunSummaryList       *[]ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListRunSummaryList     `json:"runSummaryList,omitempty"`       //
	WorkflowParameters   *ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListWorkflowParameters   `json:"workflowParameters,omitempty"`   //
	DayZeroConfig        *ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDayZeroConfig        `json:"dayZeroConfig,omitempty"`        //
	DayZeroConfigPreview *ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDayZeroConfigPreview `json:"dayZeroConfigPreview,omitempty"` // Day Zero Config Preview
	Version              *float64                                                                         `json:"version,omitempty"`              // Version
	TenantID             string                                                                           `json:"tenantId,omitempty"`             // Tenant Id
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfo struct {
	Source                    string                                                                                       `json:"source,omitempty"`                    // Source
	SerialNumber              string                                                                                       `json:"serialNumber,omitempty"`              // Serial Number
	Stack                     *bool                                                                                        `json:"stack,omitempty"`                     // Stack
	Mode                      string                                                                                       `json:"mode,omitempty"`                      // Mode
	State                     string                                                                                       `json:"state,omitempty"`                     // State
	Location                  *ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoLocation               `json:"location,omitempty"`                  //
	Description               string                                                                                       `json:"description,omitempty"`               // Description
	OnbState                  string                                                                                       `json:"onbState,omitempty"`                  // Onb State
	AuthenticatedMicNumber    string                                                                                       `json:"authenticatedMicNumber,omitempty"`    // Authenticated Mic Number
	AuthenticatedSudiSerialNo string                                                                                       `json:"authenticatedSudiSerialNo,omitempty"` // Authenticated Sudi Serial No
	CapabilitiesSupported     []string                                                                                     `json:"capabilitiesSupported,omitempty"`     // Capabilities Supported
	FeaturesSupported         []string                                                                                     `json:"featuresSupported,omitempty"`         // Features Supported
	CmState                   string                                                                                       `json:"cmState,omitempty"`                   // Cm State
	FirstContact              *float64                                                                                     `json:"firstContact,omitempty"`              // First Contact
	LastContact               *float64                                                                                     `json:"lastContact,omitempty"`               // Last Contact
	MacAddress                string                                                                                       `json:"macAddress,omitempty"`                // Mac Address
	Pid                       string                                                                                       `json:"pid,omitempty"`                       // Pid
	DeviceSudiSerialNos       []string                                                                                     `json:"deviceSudiSerialNos,omitempty"`       // Device Sudi Serial Nos
	LastUpdateOn              *float64                                                                                     `json:"lastUpdateOn,omitempty"`              // Last Update On
	WorkflowID                string                                                                                       `json:"workflowId,omitempty"`                // Workflow Id
	WorkflowName              string                                                                                       `json:"workflowName,omitempty"`              // Workflow Name
	ProjectID                 string                                                                                       `json:"projectId,omitempty"`                 // Project Id
	ProjectName               string                                                                                       `json:"projectName,omitempty"`               // Project Name
	DeviceType                string                                                                                       `json:"deviceType,omitempty"`                // Device Type
	AgentType                 string                                                                                       `json:"agentType,omitempty"`                 // Agent Type
	ImageVersion              string                                                                                       `json:"imageVersion,omitempty"`              // Image Version
	FileSystemList            *[]ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoFileSystemList       `json:"fileSystemList,omitempty"`            //
	PnpProfileList            *[]ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoPnpProfileList       `json:"pnpProfileList,omitempty"`            //
	ImageFile                 string                                                                                       `json:"imageFile,omitempty"`                 // Image File
	HTTPHeaders               *[]ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoHTTPHeaders          `json:"httpHeaders,omitempty"`               //
	NeighborLinks             *[]ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoNeighborLinks        `json:"neighborLinks,omitempty"`             //
	LastSyncTime              *float64                                                                                     `json:"lastSyncTime,omitempty"`              // Last Sync Time
	IPInterfaces              *[]ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoIPInterfaces         `json:"ipInterfaces,omitempty"`              //
	Hostname                  string                                                                                       `json:"hostname,omitempty"`                  // Hostname
	AuthStatus                string                                                                                       `json:"authStatus,omitempty"`                // Auth Status
	StackInfo                 *ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoStackInfo              `json:"stackInfo,omitempty"`                 //
	ReloadRequested           *bool                                                                                        `json:"reloadRequested,omitempty"`           // Reload Requested
	AddedOn                   *float64                                                                                     `json:"addedOn,omitempty"`                   // Added On
	SiteID                    string                                                                                       `json:"siteId,omitempty"`                    // Site Id
	AAACredentials            *ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoAAACredentials         `json:"aaaCredentials,omitempty"`            //
	UserMicNumbers            []string                                                                                     `json:"userMicNumbers,omitempty"`            // User Mic Numbers
	UserSudiSerialNos         []string                                                                                     `json:"userSudiSerialNos,omitempty"`         // User Sudi Serial Nos
	AddnMacAddrs              []string                                                                                     `json:"addnMacAddrs,omitempty"`              // Addn Mac Addrs
	PreWorkflowCliOuputs      *[]ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoPreWorkflowCliOuputs `json:"preWorkflowCliOuputs,omitempty"`      //
	Tags                      *ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoTags                   `json:"tags,omitempty"`                      // Tags
	SudiRequired              *bool                                                                                        `json:"sudiRequired,omitempty"`              // Sudi Required
	SmartAccountID            string                                                                                       `json:"smartAccountId,omitempty"`            // Smart Account Id
	VirtualAccountID          string                                                                                       `json:"virtualAccountId,omitempty"`          // Virtual Account Id
	PopulateInventory         *bool                                                                                        `json:"populateInventory,omitempty"`         // Populate Inventory
	SiteName                  string                                                                                       `json:"siteName,omitempty"`                  // Site Name
	Name                      string                                                                                       `json:"name,omitempty"`                      // Name
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoLocation struct {
	SiteID    string `json:"siteId,omitempty"`    // Site Id
	Address   string `json:"address,omitempty"`   // Address
	Latitude  string `json:"latitude,omitempty"`  // Latitude
	Longitude string `json:"longitude,omitempty"` // Longitude
	Altitude  string `json:"altitude,omitempty"`  // Altitude
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoFileSystemList struct {
	Type      string   `json:"type,omitempty"`      // Type
	Writeable *bool    `json:"writeable,omitempty"` // Writeable
	Freespace *float64 `json:"freespace,omitempty"` // Freespace
	Name      string   `json:"name,omitempty"`      // Name
	Readable  *bool    `json:"readable,omitempty"`  // Readable
	Size      *float64 `json:"size,omitempty"`      // Size
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoPnpProfileList struct {
	ProfileName       string                                                                                                `json:"profileName,omitempty"`       // Profile Name
	DiscoveryCreated  *bool                                                                                                 `json:"discoveryCreated,omitempty"`  // Discovery Created
	CreatedBy         string                                                                                                `json:"createdBy,omitempty"`         // Created By
	PrimaryEndpoint   *ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoPnpProfileListPrimaryEndpoint   `json:"primaryEndpoint,omitempty"`   //
	SecondaryEndpoint *ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoPnpProfileListSecondaryEndpoint `json:"secondaryEndpoint,omitempty"` //
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoPnpProfileListPrimaryEndpoint struct {
	Port        *float64                                                                                                       `json:"port,omitempty"`        // Port
	Protocol    string                                                                                                         `json:"protocol,omitempty"`    // Protocol
	IPv4Address *ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoPnpProfileListPrimaryEndpointIPv4Address `json:"ipv4Address,omitempty"` // Ipv4 Address
	IPv6Address *ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoPnpProfileListPrimaryEndpointIPv6Address `json:"ipv6Address,omitempty"` // Ipv6 Address
	Fqdn        string                                                                                                         `json:"fqdn,omitempty"`        // Fqdn
	Certificate string                                                                                                         `json:"certificate,omitempty"` // Certificate
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoPnpProfileListPrimaryEndpointIPv4Address interface{}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoPnpProfileListPrimaryEndpointIPv6Address interface{}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoPnpProfileListSecondaryEndpoint struct {
	Port        *float64                                                                                                         `json:"port,omitempty"`        // Port
	Protocol    string                                                                                                           `json:"protocol,omitempty"`    // Protocol
	IPv4Address *ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoPnpProfileListSecondaryEndpointIPv4Address `json:"ipv4Address,omitempty"` // Ipv4 Address
	IPv6Address *ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoPnpProfileListSecondaryEndpointIPv6Address `json:"ipv6Address,omitempty"` // Ipv6 Address
	Fqdn        string                                                                                                           `json:"fqdn,omitempty"`        // Fqdn
	Certificate string                                                                                                           `json:"certificate,omitempty"` // Certificate
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoPnpProfileListSecondaryEndpointIPv4Address interface{}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoPnpProfileListSecondaryEndpointIPv6Address interface{}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoHTTPHeaders struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoNeighborLinks struct {
	LocalInterfaceName       string `json:"localInterfaceName,omitempty"`       // Local Interface Name
	LocalShortInterfaceName  string `json:"localShortInterfaceName,omitempty"`  // Local Short Interface Name
	LocalMacAddress          string `json:"localMacAddress,omitempty"`          // Local Mac Address
	RemoteInterfaceName      string `json:"remoteInterfaceName,omitempty"`      // Remote Interface Name
	RemoteShortInterfaceName string `json:"remoteShortInterfaceName,omitempty"` // Remote Short Interface Name
	RemoteMacAddress         string `json:"remoteMacAddress,omitempty"`         // Remote Mac Address
	RemoteDeviceName         string `json:"remoteDeviceName,omitempty"`         // Remote Device Name
	RemotePlatform           string `json:"remotePlatform,omitempty"`           // Remote Platform
	RemoteVersion            string `json:"remoteVersion,omitempty"`            // Remote Version
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoIPInterfaces struct {
	Status          string                                                                                              `json:"status,omitempty"`          // Status
	MacAddress      string                                                                                              `json:"macAddress,omitempty"`      // Mac Address
	IPv4Address     *ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoIPInterfacesIPv4Address       `json:"ipv4Address,omitempty"`     // Ipv4 Address
	IPv6AddressList *[]ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoIPInterfacesIPv6AddressList `json:"ipv6AddressList,omitempty"` // Ipv6 Address List
	Name            string                                                                                              `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoIPInterfacesIPv4Address interface{}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoIPInterfacesIPv6AddressList interface{}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoStackInfo struct {
	SupportsStackWorkflows *bool                                                                                            `json:"supportsStackWorkflows,omitempty"` // Supports Stack Workflows
	IsFullRing             *bool                                                                                            `json:"isFullRing,omitempty"`             // Is Full Ring
	StackMemberList        *[]ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoStackInfoStackMemberList `json:"stackMemberList,omitempty"`        //
	StackRingProtocol      string                                                                                           `json:"stackRingProtocol,omitempty"`      // Stack Ring Protocol
	ValidLicenseLevels     []string                                                                                         `json:"validLicenseLevels,omitempty"`     // Valid License Levels
	TotalMemberCount       *float64                                                                                         `json:"totalMemberCount,omitempty"`       // Total Member Count
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoStackInfoStackMemberList struct {
	SerialNumber     string   `json:"serialNumber,omitempty"`     // Serial Number
	State            string   `json:"state,omitempty"`            // State
	Role             string   `json:"role,omitempty"`             // Role
	MacAddress       string   `json:"macAddress,omitempty"`       // Mac Address
	Pid              string   `json:"pid,omitempty"`              // Pid
	LicenseLevel     string   `json:"licenseLevel,omitempty"`     // License Level
	LicenseType      string   `json:"licenseType,omitempty"`      // License Type
	SudiSerialNumber string   `json:"sudiSerialNumber,omitempty"` // Sudi Serial Number
	HardwareVersion  string   `json:"hardwareVersion,omitempty"`  // Hardware Version
	StackNumber      *float64 `json:"stackNumber,omitempty"`      // Stack Number
	SoftwareVersion  string   `json:"softwareVersion,omitempty"`  // Software Version
	Priority         *float64 `json:"priority,omitempty"`         // Priority
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoAAACredentials struct {
	Password string `json:"password,omitempty"` // Password
	Username string `json:"username,omitempty"` // Username
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoPreWorkflowCliOuputs struct {
	Cli       string `json:"cli,omitempty"`       // Cli
	CliOutput string `json:"cliOutput,omitempty"` // Cli Output
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDeviceInfoTags interface{}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListSystemResetWorkflow struct {
	TypeID         string                                                                                 `json:"_id,omitempty"`            // Id
	State          string                                                                                 `json:"state,omitempty"`          // State
	Type           string                                                                                 `json:"type,omitempty"`           // Type
	Description    string                                                                                 `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                                                               `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                                                 `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                                                               `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                                                               `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListSystemResetWorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                                                  `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                                                 `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                                                               `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                                                               `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                                                               `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                                                 `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                                                 `json:"configId,omitempty"`       // Config Id
	Name           string                                                                                 `json:"name,omitempty"`           // Name
	Version        *float64                                                                               `json:"version,omitempty"`        // Version
	TenantID       string                                                                                 `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListSystemResetWorkflowTasks struct {
	State           string                                                                                             `json:"state,omitempty"`           // State
	Type            string                                                                                             `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                                                           `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                                                           `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                                                           `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                                                           `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListSystemResetWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                                                           `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                                                             `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListSystemResetWorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListSystemWorkflow struct {
	TypeID         string                                                                            `json:"_id,omitempty"`            // Id
	State          string                                                                            `json:"state,omitempty"`          // State
	Type           string                                                                            `json:"type,omitempty"`           // Type
	Description    string                                                                            `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                                                          `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                                            `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                                                          `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                                                          `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListSystemWorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                                             `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                                            `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                                                          `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                                                          `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                                                          `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                                            `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                                            `json:"configId,omitempty"`       // Config Id
	Name           string                                                                            `json:"name,omitempty"`           // Name
	Version        *float64                                                                          `json:"version,omitempty"`        // Version
	TenantID       string                                                                            `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListSystemWorkflowTasks struct {
	State           string                                                                                        `json:"state,omitempty"`           // State
	Type            string                                                                                        `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                                                      `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                                                      `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                                                      `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                                                      `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListSystemWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                                                      `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                                                        `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListSystemWorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListWorkflow struct {
	TypeID         string                                                                      `json:"_id,omitempty"`            // Id
	State          string                                                                      `json:"state,omitempty"`          // State
	Type           string                                                                      `json:"type,omitempty"`           // Type
	Description    string                                                                      `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                                                    `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                                      `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                                                    `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                                                    `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListWorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                                       `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                                      `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                                                    `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                                                    `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                                                    `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                                      `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                                      `json:"configId,omitempty"`       // Config Id
	Name           string                                                                      `json:"name,omitempty"`           // Name
	Version        *float64                                                                    `json:"version,omitempty"`        // Version
	TenantID       string                                                                      `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListWorkflowTasks struct {
	State           string                                                                                  `json:"state,omitempty"`           // State
	Type            string                                                                                  `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                                                `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                                                `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                                                `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                                                `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                                                `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                                                  `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListWorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListRunSummaryList struct {
	Details         string                                                                                    `json:"details,omitempty"`         // Details
	HistoryTaskInfo *ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListRunSummaryListHistoryTaskInfo `json:"historyTaskInfo,omitempty"` //
	ErrorFlag       *bool                                                                                     `json:"errorFlag,omitempty"`       // Error Flag
	Timestamp       *float64                                                                                  `json:"timestamp,omitempty"`       // Timestamp
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListRunSummaryListHistoryTaskInfo struct {
	Type         string                                                                                                  `json:"type,omitempty"`         // Type
	WorkItemList *[]ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListRunSummaryListHistoryTaskInfoWorkItemList `json:"workItemList,omitempty"` //
	TimeTaken    *float64                                                                                                `json:"timeTaken,omitempty"`    // Time Taken
	AddnDetails  *[]ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListRunSummaryListHistoryTaskInfoAddnDetails  `json:"addnDetails,omitempty"`  //
	Name         string                                                                                                  `json:"name,omitempty"`         // Name
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListRunSummaryListHistoryTaskInfoWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListRunSummaryListHistoryTaskInfoAddnDetails struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListWorkflowParameters struct {
	TopOfStackSerialNumber string                                                                                     `json:"topOfStackSerialNumber,omitempty"` // Top Of Stack Serial Number
	LicenseLevel           string                                                                                     `json:"licenseLevel,omitempty"`           // License Level
	LicenseType            string                                                                                     `json:"licenseType,omitempty"`            // License Type
	ConfigList             *[]ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListWorkflowParametersConfigList `json:"configList,omitempty"`             //
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListWorkflowParametersConfigList struct {
	ConfigParameters *[]ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListWorkflowParametersConfigListConfigParameters `json:"configParameters,omitempty"` //
	ConfigID         string                                                                                                     `json:"configId,omitempty"`         // Config Id
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListWorkflowParametersConfigListConfigParameters struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDayZeroConfig struct {
	Config string `json:"config,omitempty"` // Config
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1SuccessListDayZeroConfigPreview interface{}
type ResponseDeviceOnboardingPnpImportDevicesInBulkV1FailureList struct {
	Index     *float64 `json:"index,omitempty"`     // Index
	SerialNum string   `json:"serialNum,omitempty"` // Serial Num
	ID        string   `json:"id,omitempty"`        // Id
	Msg       string   `json:"msg,omitempty"`       // Msg
}
type ResponseDeviceOnboardingPnpResetDeviceV1 struct {
	JSONArrayResponse *[]ResponseDeviceOnboardingPnpResetDeviceV1JSONArrayResponse `json:"jsonArrayResponse,omitempty"` // Json Array Response
	JSONResponse      *ResponseDeviceOnboardingPnpResetDeviceV1JSONResponse        `json:"jsonResponse,omitempty"`      // Json Response
	Message           string                                                       `json:"message,omitempty"`           // Message
	StatusCode        *float64                                                     `json:"statusCode,omitempty"`        // Status Code
}
type ResponseDeviceOnboardingPnpResetDeviceV1JSONArrayResponse interface{}
type ResponseDeviceOnboardingPnpResetDeviceV1JSONResponse interface{}
type ResponseDeviceOnboardingPnpGetSyncResultForVirtualAccountV1 struct {
	VirtualAccountID string                                                                 `json:"virtualAccountId,omitempty"` // Virtual Account Id
	AutoSyncPeriod   *float64                                                               `json:"autoSyncPeriod,omitempty"`   // Auto Sync Period
	SyncResultStr    string                                                                 `json:"syncResultStr,omitempty"`    // Sync Result Str
	Profile          *ResponseDeviceOnboardingPnpGetSyncResultForVirtualAccountV1Profile    `json:"profile,omitempty"`          //
	CcoUser          string                                                                 `json:"ccoUser,omitempty"`          // Cco User
	SyncResult       *ResponseDeviceOnboardingPnpGetSyncResultForVirtualAccountV1SyncResult `json:"syncResult,omitempty"`       //
	Token            string                                                                 `json:"token,omitempty"`            // Token
	SyncStartTime    *float64                                                               `json:"syncStartTime,omitempty"`    // Sync Start Time
	LastSync         *float64                                                               `json:"lastSync,omitempty"`         // Last Sync
	TenantID         string                                                                 `json:"tenantId,omitempty"`         // Tenant Id
	SmartAccountID   string                                                                 `json:"smartAccountId,omitempty"`   // Smart Account Id
	Expiry           *float64                                                               `json:"expiry,omitempty"`           // Expiry
	SyncStatus       string                                                                 `json:"syncStatus,omitempty"`       // Sync Status
}
type ResponseDeviceOnboardingPnpGetSyncResultForVirtualAccountV1Profile struct {
	Proxy       *bool    `json:"proxy,omitempty"`       // Proxy
	MakeDefault *bool    `json:"makeDefault,omitempty"` // Make Default
	Port        *float64 `json:"port,omitempty"`        // Port
	ProfileID   string   `json:"profileId,omitempty"`   // Profile Id
	Name        string   `json:"name,omitempty"`        // Name
	AddressIPV4 string   `json:"addressIpV4,omitempty"` // Address Ip V4
	Cert        string   `json:"cert,omitempty"`        // Cert
	AddressFqdn string   `json:"addressFqdn,omitempty"` // Address Fqdn
}
type ResponseDeviceOnboardingPnpGetSyncResultForVirtualAccountV1SyncResult struct {
	SyncList *[]ResponseDeviceOnboardingPnpGetSyncResultForVirtualAccountV1SyncResultSyncList `json:"syncList,omitempty"` //
	SyncMsg  string                                                                           `json:"syncMsg,omitempty"`  // Sync Msg
}
type ResponseDeviceOnboardingPnpGetSyncResultForVirtualAccountV1SyncResultSyncList struct {
	SyncType     string   `json:"syncType,omitempty"`     // Sync Type
	DeviceSnList []string `json:"deviceSnList,omitempty"` // Device Sn List
}
type ResponseDeviceOnboardingPnpClaimADeviceToASiteV1 struct {
	Response string `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}
type ResponseDeviceOnboardingPnpPreviewConfigV1 struct {
	Response *ResponseDeviceOnboardingPnpPreviewConfigV1Response `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  //
}
type ResponseDeviceOnboardingPnpPreviewConfigV1Response struct {
	Complete      *bool  `json:"complete,omitempty"`      //
	Config        string `json:"config,omitempty"`        //
	Error         *bool  `json:"error,omitempty"`         //
	ErrorMessage  string `json:"errorMessage,omitempty"`  //
	ExpiredTime   *int   `json:"expiredTime,omitempty"`   //
	RfProfile     string `json:"rfProfile,omitempty"`     //
	SensorProfile string `json:"sensorProfile,omitempty"` //
	SiteID        string `json:"siteId,omitempty"`        //
	StartTime     *int   `json:"startTime,omitempty"`     //
	TaskID        string `json:"taskId,omitempty"`        //
}
type ResponseDeviceOnboardingPnpUnClaimDeviceV1 struct {
	JSONArrayResponse *[]ResponseDeviceOnboardingPnpUnClaimDeviceV1JSONArrayResponse `json:"jsonArrayResponse,omitempty"` // Json Array Response
	JSONResponse      *ResponseDeviceOnboardingPnpUnClaimDeviceV1JSONResponse        `json:"jsonResponse,omitempty"`      // Json Response
	Message           string                                                         `json:"message,omitempty"`           // Message
	StatusCode        *float64                                                       `json:"statusCode,omitempty"`        // Status Code
}
type ResponseDeviceOnboardingPnpUnClaimDeviceV1JSONArrayResponse interface{}
type ResponseDeviceOnboardingPnpUnClaimDeviceV1JSONResponse interface{}
type ResponseDeviceOnboardingPnpSyncVirtualAccountDevicesV1 struct {
	VirtualAccountID string                                                            `json:"virtualAccountId,omitempty"` // Virtual Account Id
	AutoSyncPeriod   *float64                                                          `json:"autoSyncPeriod,omitempty"`   // Auto Sync Period
	SyncResultStr    string                                                            `json:"syncResultStr,omitempty"`    // Sync Result Str
	Profile          *ResponseDeviceOnboardingPnpSyncVirtualAccountDevicesV1Profile    `json:"profile,omitempty"`          //
	CcoUser          string                                                            `json:"ccoUser,omitempty"`          // Cco User
	SyncResult       *ResponseDeviceOnboardingPnpSyncVirtualAccountDevicesV1SyncResult `json:"syncResult,omitempty"`       //
	Token            string                                                            `json:"token,omitempty"`            // Token
	SyncStartTime    *float64                                                          `json:"syncStartTime,omitempty"`    // Sync Start Time
	LastSync         *float64                                                          `json:"lastSync,omitempty"`         // Last Sync
	TenantID         string                                                            `json:"tenantId,omitempty"`         // Tenant Id
	SmartAccountID   string                                                            `json:"smartAccountId,omitempty"`   // Smart Account Id
	Expiry           *float64                                                          `json:"expiry,omitempty"`           // Expiry
	SyncStatus       string                                                            `json:"syncStatus,omitempty"`       // Sync Status
}
type ResponseDeviceOnboardingPnpSyncVirtualAccountDevicesV1Profile struct {
	Proxy       *bool    `json:"proxy,omitempty"`       // Proxy
	MakeDefault *bool    `json:"makeDefault,omitempty"` // Make Default
	Port        *float64 `json:"port,omitempty"`        // Port
	ProfileID   string   `json:"profileId,omitempty"`   // Profile Id
	Name        string   `json:"name,omitempty"`        // Name
	AddressIPV4 string   `json:"addressIpV4,omitempty"` // Address Ip V4
	Cert        string   `json:"cert,omitempty"`        // Cert
	AddressFqdn string   `json:"addressFqdn,omitempty"` // Address Fqdn
}
type ResponseDeviceOnboardingPnpSyncVirtualAccountDevicesV1SyncResult struct {
	SyncList *[]ResponseDeviceOnboardingPnpSyncVirtualAccountDevicesV1SyncResultSyncList `json:"syncList,omitempty"` //
	SyncMsg  string                                                                      `json:"syncMsg,omitempty"`  // Sync Msg
}
type ResponseDeviceOnboardingPnpSyncVirtualAccountDevicesV1SyncResultSyncList struct {
	SyncType     string   `json:"syncType,omitempty"`     // Sync Type
	DeviceSnList []string `json:"deviceSnList,omitempty"` // Device Sn List
}
type ResponseDeviceOnboardingPnpUpdateDeviceV1 struct {
	ID                   string                                                         `json:"id,omitempty"`                   // Id
	DeviceInfo           *ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfo           `json:"deviceInfo,omitempty"`           //
	SystemResetWorkflow  *ResponseDeviceOnboardingPnpUpdateDeviceV1SystemResetWorkflow  `json:"systemResetWorkflow,omitempty"`  //
	SystemWorkflow       *ResponseDeviceOnboardingPnpUpdateDeviceV1SystemWorkflow       `json:"systemWorkflow,omitempty"`       //
	Workflow             *ResponseDeviceOnboardingPnpUpdateDeviceV1Workflow             `json:"workflow,omitempty"`             //
	RunSummaryList       *[]ResponseDeviceOnboardingPnpUpdateDeviceV1RunSummaryList     `json:"runSummaryList,omitempty"`       //
	WorkflowParameters   *ResponseDeviceOnboardingPnpUpdateDeviceV1WorkflowParameters   `json:"workflowParameters,omitempty"`   //
	DayZeroConfig        *ResponseDeviceOnboardingPnpUpdateDeviceV1DayZeroConfig        `json:"dayZeroConfig,omitempty"`        //
	DayZeroConfigPreview *ResponseDeviceOnboardingPnpUpdateDeviceV1DayZeroConfigPreview `json:"dayZeroConfigPreview,omitempty"` // Day Zero Config Preview
	Version              *float64                                                       `json:"version,omitempty"`              // Version
	TenantID             string                                                         `json:"tenantId,omitempty"`             // Tenant Id
}
type ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfo struct {
	Source                    string                                                                     `json:"source,omitempty"`                    // Source
	SerialNumber              string                                                                     `json:"serialNumber,omitempty"`              // Serial Number
	Stack                     *bool                                                                      `json:"stack,omitempty"`                     // Stack
	Mode                      string                                                                     `json:"mode,omitempty"`                      // Mode
	State                     string                                                                     `json:"state,omitempty"`                     // State
	Location                  *ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoLocation               `json:"location,omitempty"`                  //
	Description               string                                                                     `json:"description,omitempty"`               // Description
	OnbState                  string                                                                     `json:"onbState,omitempty"`                  // Onb State
	AuthenticatedMicNumber    string                                                                     `json:"authenticatedMicNumber,omitempty"`    // Authenticated Mic Number
	AuthenticatedSudiSerialNo string                                                                     `json:"authenticatedSudiSerialNo,omitempty"` // Authenticated Sudi Serial No
	CapabilitiesSupported     []string                                                                   `json:"capabilitiesSupported,omitempty"`     // Capabilities Supported
	FeaturesSupported         []string                                                                   `json:"featuresSupported,omitempty"`         // Features Supported
	CmState                   string                                                                     `json:"cmState,omitempty"`                   // Cm State
	FirstContact              *float64                                                                   `json:"firstContact,omitempty"`              // First Contact
	LastContact               *float64                                                                   `json:"lastContact,omitempty"`               // Last Contact
	MacAddress                string                                                                     `json:"macAddress,omitempty"`                // Mac Address
	Pid                       string                                                                     `json:"pid,omitempty"`                       // Pid
	DeviceSudiSerialNos       []string                                                                   `json:"deviceSudiSerialNos,omitempty"`       // Device Sudi Serial Nos
	LastUpdateOn              *float64                                                                   `json:"lastUpdateOn,omitempty"`              // Last Update On
	WorkflowID                string                                                                     `json:"workflowId,omitempty"`                // Workflow Id
	WorkflowName              string                                                                     `json:"workflowName,omitempty"`              // Workflow Name
	ProjectID                 string                                                                     `json:"projectId,omitempty"`                 // Project Id
	ProjectName               string                                                                     `json:"projectName,omitempty"`               // Project Name
	DeviceType                string                                                                     `json:"deviceType,omitempty"`                // Device Type
	AgentType                 string                                                                     `json:"agentType,omitempty"`                 // Agent Type
	ImageVersion              string                                                                     `json:"imageVersion,omitempty"`              // Image Version
	FileSystemList            *[]ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoFileSystemList       `json:"fileSystemList,omitempty"`            //
	PnpProfileList            *[]ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoPnpProfileList       `json:"pnpProfileList,omitempty"`            //
	ImageFile                 string                                                                     `json:"imageFile,omitempty"`                 // Image File
	HTTPHeaders               *[]ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoHTTPHeaders          `json:"httpHeaders,omitempty"`               //
	NeighborLinks             *[]ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoNeighborLinks        `json:"neighborLinks,omitempty"`             //
	LastSyncTime              *float64                                                                   `json:"lastSyncTime,omitempty"`              // Last Sync Time
	IPInterfaces              *[]ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoIPInterfaces         `json:"ipInterfaces,omitempty"`              //
	Hostname                  string                                                                     `json:"hostname,omitempty"`                  // Hostname
	AuthStatus                string                                                                     `json:"authStatus,omitempty"`                // Auth Status
	StackInfo                 *ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoStackInfo              `json:"stackInfo,omitempty"`                 //
	ReloadRequested           *bool                                                                      `json:"reloadRequested,omitempty"`           // Reload Requested
	AddedOn                   *float64                                                                   `json:"addedOn,omitempty"`                   // Added On
	SiteID                    string                                                                     `json:"siteId,omitempty"`                    // Site Id
	AAACredentials            *ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoAAACredentials         `json:"aaaCredentials,omitempty"`            //
	UserMicNumbers            []string                                                                   `json:"userMicNumbers,omitempty"`            // User Mic Numbers
	UserSudiSerialNos         []string                                                                   `json:"userSudiSerialNos,omitempty"`         // User Sudi Serial Nos
	AddnMacAddrs              []string                                                                   `json:"addnMacAddrs,omitempty"`              // Addn Mac Addrs
	PreWorkflowCliOuputs      *[]ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoPreWorkflowCliOuputs `json:"preWorkflowCliOuputs,omitempty"`      //
	Tags                      *ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoTags                   `json:"tags,omitempty"`                      // Tags
	SudiRequired              *bool                                                                      `json:"sudiRequired,omitempty"`              // Sudi Required
	SmartAccountID            string                                                                     `json:"smartAccountId,omitempty"`            // Smart Account Id
	VirtualAccountID          string                                                                     `json:"virtualAccountId,omitempty"`          // Virtual Account Id
	PopulateInventory         *bool                                                                      `json:"populateInventory,omitempty"`         // Populate Inventory
	SiteName                  string                                                                     `json:"siteName,omitempty"`                  // Site Name
	Name                      string                                                                     `json:"name,omitempty"`                      // Name
}
type ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoLocation struct {
	SiteID    string `json:"siteId,omitempty"`    // Site Id
	Address   string `json:"address,omitempty"`   // Address
	Latitude  string `json:"latitude,omitempty"`  // Latitude
	Longitude string `json:"longitude,omitempty"` // Longitude
	Altitude  string `json:"altitude,omitempty"`  // Altitude
}
type ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoFileSystemList struct {
	Type      string   `json:"type,omitempty"`      // Type
	Writeable *bool    `json:"writeable,omitempty"` // Writeable
	Freespace *float64 `json:"freespace,omitempty"` // Freespace
	Name      string   `json:"name,omitempty"`      // Name
	Readable  *bool    `json:"readable,omitempty"`  // Readable
	Size      *float64 `json:"size,omitempty"`      // Size
}
type ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoPnpProfileList struct {
	ProfileName       string                                                                              `json:"profileName,omitempty"`       // Profile Name
	DiscoveryCreated  *bool                                                                               `json:"discoveryCreated,omitempty"`  // Discovery Created
	CreatedBy         string                                                                              `json:"createdBy,omitempty"`         // Created By
	PrimaryEndpoint   *ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoPnpProfileListPrimaryEndpoint   `json:"primaryEndpoint,omitempty"`   //
	SecondaryEndpoint *ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoPnpProfileListSecondaryEndpoint `json:"secondaryEndpoint,omitempty"` //
}
type ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoPnpProfileListPrimaryEndpoint struct {
	Port        *float64                                                                                     `json:"port,omitempty"`        // Port
	Protocol    string                                                                                       `json:"protocol,omitempty"`    // Protocol
	IPv4Address *ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoPnpProfileListPrimaryEndpointIPv4Address `json:"ipv4Address,omitempty"` // Ipv4 Address
	IPv6Address *ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoPnpProfileListPrimaryEndpointIPv6Address `json:"ipv6Address,omitempty"` // Ipv6 Address
	Fqdn        string                                                                                       `json:"fqdn,omitempty"`        // Fqdn
	Certificate string                                                                                       `json:"certificate,omitempty"` // Certificate
}
type ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoPnpProfileListPrimaryEndpointIPv4Address interface{}
type ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoPnpProfileListPrimaryEndpointIPv6Address interface{}
type ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoPnpProfileListSecondaryEndpoint struct {
	Port        *float64                                                                                       `json:"port,omitempty"`        // Port
	Protocol    string                                                                                         `json:"protocol,omitempty"`    // Protocol
	IPv4Address *ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoPnpProfileListSecondaryEndpointIPv4Address `json:"ipv4Address,omitempty"` // Ipv4 Address
	IPv6Address *ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoPnpProfileListSecondaryEndpointIPv6Address `json:"ipv6Address,omitempty"` // Ipv6 Address
	Fqdn        string                                                                                         `json:"fqdn,omitempty"`        // Fqdn
	Certificate string                                                                                         `json:"certificate,omitempty"` // Certificate
}
type ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoPnpProfileListSecondaryEndpointIPv4Address interface{}
type ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoPnpProfileListSecondaryEndpointIPv6Address interface{}
type ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoHTTPHeaders struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoNeighborLinks struct {
	LocalInterfaceName       string `json:"localInterfaceName,omitempty"`       // Local Interface Name
	LocalShortInterfaceName  string `json:"localShortInterfaceName,omitempty"`  // Local Short Interface Name
	LocalMacAddress          string `json:"localMacAddress,omitempty"`          // Local Mac Address
	RemoteInterfaceName      string `json:"remoteInterfaceName,omitempty"`      // Remote Interface Name
	RemoteShortInterfaceName string `json:"remoteShortInterfaceName,omitempty"` // Remote Short Interface Name
	RemoteMacAddress         string `json:"remoteMacAddress,omitempty"`         // Remote Mac Address
	RemoteDeviceName         string `json:"remoteDeviceName,omitempty"`         // Remote Device Name
	RemotePlatform           string `json:"remotePlatform,omitempty"`           // Remote Platform
	RemoteVersion            string `json:"remoteVersion,omitempty"`            // Remote Version
}
type ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoIPInterfaces struct {
	Status          string                                                                            `json:"status,omitempty"`          // Status
	MacAddress      string                                                                            `json:"macAddress,omitempty"`      // Mac Address
	IPv4Address     *ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoIPInterfacesIPv4Address       `json:"ipv4Address,omitempty"`     // Ipv4 Address
	IPv6AddressList *[]ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoIPInterfacesIPv6AddressList `json:"ipv6AddressList,omitempty"` // Ipv6 Address List
	Name            string                                                                            `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoIPInterfacesIPv4Address interface{}
type ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoIPInterfacesIPv6AddressList interface{}
type ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoStackInfo struct {
	SupportsStackWorkflows *bool                                                                          `json:"supportsStackWorkflows,omitempty"` // Supports Stack Workflows
	IsFullRing             *bool                                                                          `json:"isFullRing,omitempty"`             // Is Full Ring
	StackMemberList        *[]ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoStackInfoStackMemberList `json:"stackMemberList,omitempty"`        //
	StackRingProtocol      string                                                                         `json:"stackRingProtocol,omitempty"`      // Stack Ring Protocol
	ValidLicenseLevels     []string                                                                       `json:"validLicenseLevels,omitempty"`     // Valid License Levels
	TotalMemberCount       *float64                                                                       `json:"totalMemberCount,omitempty"`       // Total Member Count
}
type ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoStackInfoStackMemberList struct {
	SerialNumber     string   `json:"serialNumber,omitempty"`     // Serial Number
	State            string   `json:"state,omitempty"`            // State
	Role             string   `json:"role,omitempty"`             // Role
	MacAddress       string   `json:"macAddress,omitempty"`       // Mac Address
	Pid              string   `json:"pid,omitempty"`              // Pid
	LicenseLevel     string   `json:"licenseLevel,omitempty"`     // License Level
	LicenseType      string   `json:"licenseType,omitempty"`      // License Type
	SudiSerialNumber string   `json:"sudiSerialNumber,omitempty"` // Sudi Serial Number
	HardwareVersion  string   `json:"hardwareVersion,omitempty"`  // Hardware Version
	StackNumber      *float64 `json:"stackNumber,omitempty"`      // Stack Number
	SoftwareVersion  string   `json:"softwareVersion,omitempty"`  // Software Version
	Priority         *float64 `json:"priority,omitempty"`         // Priority
}
type ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoAAACredentials struct {
	Password string `json:"password,omitempty"` // Password
	Username string `json:"username,omitempty"` // Username
}
type ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoPreWorkflowCliOuputs struct {
	Cli       string `json:"cli,omitempty"`       // Cli
	CliOutput string `json:"cliOutput,omitempty"` // Cli Output
}
type ResponseDeviceOnboardingPnpUpdateDeviceV1DeviceInfoTags interface{}
type ResponseDeviceOnboardingPnpUpdateDeviceV1SystemResetWorkflow struct {
	TypeID         string                                                               `json:"_id,omitempty"`            // Id
	State          string                                                               `json:"state,omitempty"`          // State
	Type           string                                                               `json:"type,omitempty"`           // Type
	Description    string                                                               `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                                             `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                               `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                                             `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                                             `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpUpdateDeviceV1SystemResetWorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                                `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                               `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                                             `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                                             `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                                             `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                               `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                               `json:"configId,omitempty"`       // Config Id
	Name           string                                                               `json:"name,omitempty"`           // Name
	Version        *float64                                                             `json:"version,omitempty"`        // Version
	TenantID       string                                                               `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpUpdateDeviceV1SystemResetWorkflowTasks struct {
	State           string                                                                           `json:"state,omitempty"`           // State
	Type            string                                                                           `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                                         `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                                         `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                                         `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                                         `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpUpdateDeviceV1SystemResetWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                                         `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                                           `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpUpdateDeviceV1SystemResetWorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpUpdateDeviceV1SystemWorkflow struct {
	TypeID         string                                                          `json:"_id,omitempty"`            // Id
	State          string                                                          `json:"state,omitempty"`          // State
	Type           string                                                          `json:"type,omitempty"`           // Type
	Description    string                                                          `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                                        `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                          `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                                        `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                                        `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpUpdateDeviceV1SystemWorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                           `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                          `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                                        `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                                        `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                                        `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                          `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                          `json:"configId,omitempty"`       // Config Id
	Name           string                                                          `json:"name,omitempty"`           // Name
	Version        *float64                                                        `json:"version,omitempty"`        // Version
	TenantID       string                                                          `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpUpdateDeviceV1SystemWorkflowTasks struct {
	State           string                                                                      `json:"state,omitempty"`           // State
	Type            string                                                                      `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                                    `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                                    `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                                    `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                                    `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpUpdateDeviceV1SystemWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                                    `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                                      `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpUpdateDeviceV1SystemWorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpUpdateDeviceV1Workflow struct {
	TypeID         string                                                    `json:"_id,omitempty"`            // Id
	State          string                                                    `json:"state,omitempty"`          // State
	Type           string                                                    `json:"type,omitempty"`           // Type
	Description    string                                                    `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                                  `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                    `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                                  `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                                  `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpUpdateDeviceV1WorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                     `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                    `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                                  `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                                  `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                                  `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                    `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                    `json:"configId,omitempty"`       // Config Id
	Name           string                                                    `json:"name,omitempty"`           // Name
	Version        *float64                                                  `json:"version,omitempty"`        // Version
	TenantID       string                                                    `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpUpdateDeviceV1WorkflowTasks struct {
	State           string                                                                `json:"state,omitempty"`           // State
	Type            string                                                                `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                              `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                              `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                              `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                              `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpUpdateDeviceV1WorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                              `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                                `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpUpdateDeviceV1WorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpUpdateDeviceV1RunSummaryList struct {
	Details         string                                                                  `json:"details,omitempty"`         // Details
	HistoryTaskInfo *ResponseDeviceOnboardingPnpUpdateDeviceV1RunSummaryListHistoryTaskInfo `json:"historyTaskInfo,omitempty"` //
	ErrorFlag       *bool                                                                   `json:"errorFlag,omitempty"`       // Error Flag
	Timestamp       *float64                                                                `json:"timestamp,omitempty"`       // Timestamp
}
type ResponseDeviceOnboardingPnpUpdateDeviceV1RunSummaryListHistoryTaskInfo struct {
	Type         string                                                                                `json:"type,omitempty"`         // Type
	WorkItemList *[]ResponseDeviceOnboardingPnpUpdateDeviceV1RunSummaryListHistoryTaskInfoWorkItemList `json:"workItemList,omitempty"` //
	TimeTaken    *float64                                                                              `json:"timeTaken,omitempty"`    // Time Taken
	AddnDetails  *[]ResponseDeviceOnboardingPnpUpdateDeviceV1RunSummaryListHistoryTaskInfoAddnDetails  `json:"addnDetails,omitempty"`  //
	Name         string                                                                                `json:"name,omitempty"`         // Name
}
type ResponseDeviceOnboardingPnpUpdateDeviceV1RunSummaryListHistoryTaskInfoWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpUpdateDeviceV1RunSummaryListHistoryTaskInfoAddnDetails struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseDeviceOnboardingPnpUpdateDeviceV1WorkflowParameters struct {
	TopOfStackSerialNumber string                                                                   `json:"topOfStackSerialNumber,omitempty"` // Top Of Stack Serial Number
	LicenseLevel           string                                                                   `json:"licenseLevel,omitempty"`           // License Level
	LicenseType            string                                                                   `json:"licenseType,omitempty"`            // License Type
	ConfigList             *[]ResponseDeviceOnboardingPnpUpdateDeviceV1WorkflowParametersConfigList `json:"configList,omitempty"`             //
}
type ResponseDeviceOnboardingPnpUpdateDeviceV1WorkflowParametersConfigList struct {
	ConfigParameters *[]ResponseDeviceOnboardingPnpUpdateDeviceV1WorkflowParametersConfigListConfigParameters `json:"configParameters,omitempty"` //
	ConfigID         string                                                                                   `json:"configId,omitempty"`         // Config Id
}
type ResponseDeviceOnboardingPnpUpdateDeviceV1WorkflowParametersConfigListConfigParameters struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseDeviceOnboardingPnpUpdateDeviceV1DayZeroConfig struct {
	Config string `json:"config,omitempty"` // Config
}
type ResponseDeviceOnboardingPnpUpdateDeviceV1DayZeroConfigPreview interface{}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1 struct {
	ID                   string                                                                    `json:"id,omitempty"`                   // Id
	DeviceInfo           *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfo           `json:"deviceInfo,omitempty"`           //
	SystemResetWorkflow  *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1SystemResetWorkflow  `json:"systemResetWorkflow,omitempty"`  //
	SystemWorkflow       *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1SystemWorkflow       `json:"systemWorkflow,omitempty"`       //
	Workflow             *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1Workflow             `json:"workflow,omitempty"`             //
	RunSummaryList       *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1RunSummaryList     `json:"runSummaryList,omitempty"`       //
	WorkflowParameters   *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1WorkflowParameters   `json:"workflowParameters,omitempty"`   //
	DayZeroConfig        *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DayZeroConfig        `json:"dayZeroConfig,omitempty"`        //
	DayZeroConfigPreview *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DayZeroConfigPreview `json:"dayZeroConfigPreview,omitempty"` // Day Zero Config Preview
	Version              *float64                                                                  `json:"version,omitempty"`              // Version
	TenantID             string                                                                    `json:"tenantId,omitempty"`             // Tenant Id
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfo struct {
	Source                    string                                                                                `json:"source,omitempty"`                    // Source
	SerialNumber              string                                                                                `json:"serialNumber,omitempty"`              // Serial Number
	Stack                     *bool                                                                                 `json:"stack,omitempty"`                     // Stack
	Mode                      string                                                                                `json:"mode,omitempty"`                      // Mode
	State                     string                                                                                `json:"state,omitempty"`                     // State
	Location                  *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoLocation               `json:"location,omitempty"`                  //
	Description               string                                                                                `json:"description,omitempty"`               // Description
	OnbState                  string                                                                                `json:"onbState,omitempty"`                  // Onb State
	AuthenticatedMicNumber    string                                                                                `json:"authenticatedMicNumber,omitempty"`    // Authenticated Mic Number
	AuthenticatedSudiSerialNo string                                                                                `json:"authenticatedSudiSerialNo,omitempty"` // Authenticated Sudi Serial No
	CapabilitiesSupported     []string                                                                              `json:"capabilitiesSupported,omitempty"`     // Capabilities Supported
	FeaturesSupported         []string                                                                              `json:"featuresSupported,omitempty"`         // Features Supported
	CmState                   string                                                                                `json:"cmState,omitempty"`                   // Cm State
	FirstContact              *float64                                                                              `json:"firstContact,omitempty"`              // First Contact
	LastContact               *float64                                                                              `json:"lastContact,omitempty"`               // Last Contact
	MacAddress                string                                                                                `json:"macAddress,omitempty"`                // Mac Address
	Pid                       string                                                                                `json:"pid,omitempty"`                       // Pid
	DeviceSudiSerialNos       []string                                                                              `json:"deviceSudiSerialNos,omitempty"`       // Device Sudi Serial Nos
	LastUpdateOn              *float64                                                                              `json:"lastUpdateOn,omitempty"`              // Last Update On
	WorkflowID                string                                                                                `json:"workflowId,omitempty"`                // Workflow Id
	WorkflowName              string                                                                                `json:"workflowName,omitempty"`              // Workflow Name
	ProjectID                 string                                                                                `json:"projectId,omitempty"`                 // Project Id
	ProjectName               string                                                                                `json:"projectName,omitempty"`               // Project Name
	DeviceType                string                                                                                `json:"deviceType,omitempty"`                // Device Type
	AgentType                 string                                                                                `json:"agentType,omitempty"`                 // Agent Type
	ImageVersion              string                                                                                `json:"imageVersion,omitempty"`              // Image Version
	FileSystemList            *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoFileSystemList       `json:"fileSystemList,omitempty"`            //
	PnpProfileList            *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoPnpProfileList       `json:"pnpProfileList,omitempty"`            //
	ImageFile                 string                                                                                `json:"imageFile,omitempty"`                 // Image File
	HTTPHeaders               *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoHTTPHeaders          `json:"httpHeaders,omitempty"`               //
	NeighborLinks             *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoNeighborLinks        `json:"neighborLinks,omitempty"`             //
	LastSyncTime              *float64                                                                              `json:"lastSyncTime,omitempty"`              // Last Sync Time
	IPInterfaces              *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoIPInterfaces         `json:"ipInterfaces,omitempty"`              //
	Hostname                  string                                                                                `json:"hostname,omitempty"`                  // Hostname
	AuthStatus                string                                                                                `json:"authStatus,omitempty"`                // Auth Status
	StackInfo                 *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoStackInfo              `json:"stackInfo,omitempty"`                 //
	ReloadRequested           *bool                                                                                 `json:"reloadRequested,omitempty"`           // Reload Requested
	AddedOn                   *float64                                                                              `json:"addedOn,omitempty"`                   // Added On
	SiteID                    string                                                                                `json:"siteId,omitempty"`                    // Site Id
	AAACredentials            *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoAAACredentials         `json:"aaaCredentials,omitempty"`            //
	UserMicNumbers            []string                                                                              `json:"userMicNumbers,omitempty"`            // User Mic Numbers
	UserSudiSerialNos         []string                                                                              `json:"userSudiSerialNos,omitempty"`         // User Sudi Serial Nos
	AddnMacAddrs              []string                                                                              `json:"addnMacAddrs,omitempty"`              // Addn Mac Addrs
	PreWorkflowCliOuputs      *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoPreWorkflowCliOuputs `json:"preWorkflowCliOuputs,omitempty"`      //
	Tags                      *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoTags                   `json:"tags,omitempty"`                      // Tags
	SudiRequired              *bool                                                                                 `json:"sudiRequired,omitempty"`              // Sudi Required
	SmartAccountID            string                                                                                `json:"smartAccountId,omitempty"`            // Smart Account Id
	VirtualAccountID          string                                                                                `json:"virtualAccountId,omitempty"`          // Virtual Account Id
	PopulateInventory         *bool                                                                                 `json:"populateInventory,omitempty"`         // Populate Inventory
	SiteName                  string                                                                                `json:"siteName,omitempty"`                  // Site Name
	Name                      string                                                                                `json:"name,omitempty"`                      // Name
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoLocation struct {
	SiteID    string `json:"siteId,omitempty"`    // Site Id
	Address   string `json:"address,omitempty"`   // Address
	Latitude  string `json:"latitude,omitempty"`  // Latitude
	Longitude string `json:"longitude,omitempty"` // Longitude
	Altitude  string `json:"altitude,omitempty"`  // Altitude
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoFileSystemList struct {
	Type      string   `json:"type,omitempty"`      // Type
	Writeable *bool    `json:"writeable,omitempty"` // Writeable
	Freespace *float64 `json:"freespace,omitempty"` // Freespace
	Name      string   `json:"name,omitempty"`      // Name
	Readable  *bool    `json:"readable,omitempty"`  // Readable
	Size      *float64 `json:"size,omitempty"`      // Size
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoPnpProfileList struct {
	ProfileName       string                                                                                         `json:"profileName,omitempty"`       // Profile Name
	DiscoveryCreated  *bool                                                                                          `json:"discoveryCreated,omitempty"`  // Discovery Created
	CreatedBy         string                                                                                         `json:"createdBy,omitempty"`         // Created By
	PrimaryEndpoint   *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoPnpProfileListPrimaryEndpoint   `json:"primaryEndpoint,omitempty"`   //
	SecondaryEndpoint *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoPnpProfileListSecondaryEndpoint `json:"secondaryEndpoint,omitempty"` //
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoPnpProfileListPrimaryEndpoint struct {
	Port        *float64                                                                                                `json:"port,omitempty"`        // Port
	Protocol    string                                                                                                  `json:"protocol,omitempty"`    // Protocol
	IPv4Address *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoPnpProfileListPrimaryEndpointIPv4Address `json:"ipv4Address,omitempty"` // Ipv4 Address
	IPv6Address *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoPnpProfileListPrimaryEndpointIPv6Address `json:"ipv6Address,omitempty"` // Ipv6 Address
	Fqdn        string                                                                                                  `json:"fqdn,omitempty"`        // Fqdn
	Certificate string                                                                                                  `json:"certificate,omitempty"` // Certificate
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoPnpProfileListPrimaryEndpointIPv4Address interface{}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoPnpProfileListPrimaryEndpointIPv6Address interface{}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoPnpProfileListSecondaryEndpoint struct {
	Port        *float64                                                                                                  `json:"port,omitempty"`        // Port
	Protocol    string                                                                                                    `json:"protocol,omitempty"`    // Protocol
	IPv4Address *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoPnpProfileListSecondaryEndpointIPv4Address `json:"ipv4Address,omitempty"` // Ipv4 Address
	IPv6Address *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoPnpProfileListSecondaryEndpointIPv6Address `json:"ipv6Address,omitempty"` // Ipv6 Address
	Fqdn        string                                                                                                    `json:"fqdn,omitempty"`        // Fqdn
	Certificate string                                                                                                    `json:"certificate,omitempty"` // Certificate
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoPnpProfileListSecondaryEndpointIPv4Address interface{}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoPnpProfileListSecondaryEndpointIPv6Address interface{}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoHTTPHeaders struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoNeighborLinks struct {
	LocalInterfaceName       string `json:"localInterfaceName,omitempty"`       // Local Interface Name
	LocalShortInterfaceName  string `json:"localShortInterfaceName,omitempty"`  // Local Short Interface Name
	LocalMacAddress          string `json:"localMacAddress,omitempty"`          // Local Mac Address
	RemoteInterfaceName      string `json:"remoteInterfaceName,omitempty"`      // Remote Interface Name
	RemoteShortInterfaceName string `json:"remoteShortInterfaceName,omitempty"` // Remote Short Interface Name
	RemoteMacAddress         string `json:"remoteMacAddress,omitempty"`         // Remote Mac Address
	RemoteDeviceName         string `json:"remoteDeviceName,omitempty"`         // Remote Device Name
	RemotePlatform           string `json:"remotePlatform,omitempty"`           // Remote Platform
	RemoteVersion            string `json:"remoteVersion,omitempty"`            // Remote Version
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoIPInterfaces struct {
	Status          string                                                                                       `json:"status,omitempty"`          // Status
	MacAddress      string                                                                                       `json:"macAddress,omitempty"`      // Mac Address
	IPv4Address     *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoIPInterfacesIPv4Address       `json:"ipv4Address,omitempty"`     // Ipv4 Address
	IPv6AddressList *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoIPInterfacesIPv6AddressList `json:"ipv6AddressList,omitempty"` // Ipv6 Address List
	Name            string                                                                                       `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoIPInterfacesIPv4Address interface{}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoIPInterfacesIPv6AddressList interface{}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoStackInfo struct {
	SupportsStackWorkflows *bool                                                                                     `json:"supportsStackWorkflows,omitempty"` // Supports Stack Workflows
	IsFullRing             *bool                                                                                     `json:"isFullRing,omitempty"`             // Is Full Ring
	StackMemberList        *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoStackInfoStackMemberList `json:"stackMemberList,omitempty"`        //
	StackRingProtocol      string                                                                                    `json:"stackRingProtocol,omitempty"`      // Stack Ring Protocol
	ValidLicenseLevels     []string                                                                                  `json:"validLicenseLevels,omitempty"`     // Valid License Levels
	TotalMemberCount       *float64                                                                                  `json:"totalMemberCount,omitempty"`       // Total Member Count
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoStackInfoStackMemberList struct {
	SerialNumber     string   `json:"serialNumber,omitempty"`     // Serial Number
	State            string   `json:"state,omitempty"`            // State
	Role             string   `json:"role,omitempty"`             // Role
	MacAddress       string   `json:"macAddress,omitempty"`       // Mac Address
	Pid              string   `json:"pid,omitempty"`              // Pid
	LicenseLevel     string   `json:"licenseLevel,omitempty"`     // License Level
	LicenseType      string   `json:"licenseType,omitempty"`      // License Type
	SudiSerialNumber string   `json:"sudiSerialNumber,omitempty"` // Sudi Serial Number
	HardwareVersion  string   `json:"hardwareVersion,omitempty"`  // Hardware Version
	StackNumber      *float64 `json:"stackNumber,omitempty"`      // Stack Number
	SoftwareVersion  string   `json:"softwareVersion,omitempty"`  // Software Version
	Priority         *float64 `json:"priority,omitempty"`         // Priority
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoAAACredentials struct {
	Password string `json:"password,omitempty"` // Password
	Username string `json:"username,omitempty"` // Username
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoPreWorkflowCliOuputs struct {
	Cli       string `json:"cli,omitempty"`       // Cli
	CliOutput string `json:"cliOutput,omitempty"` // Cli Output
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DeviceInfoTags interface{}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1SystemResetWorkflow struct {
	TypeID         string                                                                          `json:"_id,omitempty"`            // Id
	State          string                                                                          `json:"state,omitempty"`          // State
	Type           string                                                                          `json:"type,omitempty"`           // Type
	Description    string                                                                          `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                                                        `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                                          `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                                                        `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                                                        `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1SystemResetWorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                                           `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                                          `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                                                        `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                                                        `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                                                        `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                                          `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                                          `json:"configId,omitempty"`       // Config Id
	Name           string                                                                          `json:"name,omitempty"`           // Name
	Version        *float64                                                                        `json:"version,omitempty"`        // Version
	TenantID       string                                                                          `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1SystemResetWorkflowTasks struct {
	State           string                                                                                      `json:"state,omitempty"`           // State
	Type            string                                                                                      `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                                                    `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                                                    `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                                                    `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                                                    `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1SystemResetWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                                                    `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                                                      `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1SystemResetWorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1SystemWorkflow struct {
	TypeID         string                                                                     `json:"_id,omitempty"`            // Id
	State          string                                                                     `json:"state,omitempty"`          // State
	Type           string                                                                     `json:"type,omitempty"`           // Type
	Description    string                                                                     `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                                                   `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                                     `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                                                   `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                                                   `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1SystemWorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                                      `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                                     `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                                                   `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                                                   `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                                                   `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                                     `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                                     `json:"configId,omitempty"`       // Config Id
	Name           string                                                                     `json:"name,omitempty"`           // Name
	Version        *float64                                                                   `json:"version,omitempty"`        // Version
	TenantID       string                                                                     `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1SystemWorkflowTasks struct {
	State           string                                                                                 `json:"state,omitempty"`           // State
	Type            string                                                                                 `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                                               `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                                               `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                                               `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                                               `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1SystemWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                                               `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                                                 `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1SystemWorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1Workflow struct {
	TypeID         string                                                               `json:"_id,omitempty"`            // Id
	State          string                                                               `json:"state,omitempty"`          // State
	Type           string                                                               `json:"type,omitempty"`           // Type
	Description    string                                                               `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                                             `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                               `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                                             `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                                             `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1WorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                                `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                               `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                                             `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                                             `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                                             `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                               `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                               `json:"configId,omitempty"`       // Config Id
	Name           string                                                               `json:"name,omitempty"`           // Name
	Version        *float64                                                             `json:"version,omitempty"`        // Version
	TenantID       string                                                               `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1WorkflowTasks struct {
	State           string                                                                           `json:"state,omitempty"`           // State
	Type            string                                                                           `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                                         `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                                         `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                                         `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                                         `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1WorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                                         `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                                           `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1WorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1RunSummaryList struct {
	Details         string                                                                             `json:"details,omitempty"`         // Details
	HistoryTaskInfo *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1RunSummaryListHistoryTaskInfo `json:"historyTaskInfo,omitempty"` //
	ErrorFlag       *bool                                                                              `json:"errorFlag,omitempty"`       // Error Flag
	Timestamp       *float64                                                                           `json:"timestamp,omitempty"`       // Timestamp
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1RunSummaryListHistoryTaskInfo struct {
	Type         string                                                                                           `json:"type,omitempty"`         // Type
	WorkItemList *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1RunSummaryListHistoryTaskInfoWorkItemList `json:"workItemList,omitempty"` //
	TimeTaken    *float64                                                                                         `json:"timeTaken,omitempty"`    // Time Taken
	AddnDetails  *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1RunSummaryListHistoryTaskInfoAddnDetails  `json:"addnDetails,omitempty"`  //
	Name         string                                                                                           `json:"name,omitempty"`         // Name
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1RunSummaryListHistoryTaskInfoWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1RunSummaryListHistoryTaskInfoAddnDetails struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1WorkflowParameters struct {
	TopOfStackSerialNumber string                                                                              `json:"topOfStackSerialNumber,omitempty"` // Top Of Stack Serial Number
	LicenseLevel           string                                                                              `json:"licenseLevel,omitempty"`           // License Level
	LicenseType            string                                                                              `json:"licenseType,omitempty"`            // License Type
	ConfigList             *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1WorkflowParametersConfigList `json:"configList,omitempty"`             //
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1WorkflowParametersConfigList struct {
	ConfigParameters *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1WorkflowParametersConfigListConfigParameters `json:"configParameters,omitempty"` //
	ConfigID         string                                                                                              `json:"configId,omitempty"`         // Config Id
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1WorkflowParametersConfigListConfigParameters struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DayZeroConfig struct {
	Config string `json:"config,omitempty"` // Config
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1DayZeroConfigPreview interface{}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1 struct {
	TypeID               string                                                          `json:"_id,omitempty"`                  // Id
	DeviceInfo           *ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfo           `json:"deviceInfo,omitempty"`           //
	SystemResetWorkflow  *ResponseDeviceOnboardingPnpGetDeviceByIDV1SystemResetWorkflow  `json:"systemResetWorkflow,omitempty"`  //
	SystemWorkflow       *ResponseDeviceOnboardingPnpGetDeviceByIDV1SystemWorkflow       `json:"systemWorkflow,omitempty"`       //
	Workflow             *ResponseDeviceOnboardingPnpGetDeviceByIDV1Workflow             `json:"workflow,omitempty"`             //
	RunSummaryList       *[]ResponseDeviceOnboardingPnpGetDeviceByIDV1RunSummaryList     `json:"runSummaryList,omitempty"`       //
	WorkflowParameters   *ResponseDeviceOnboardingPnpGetDeviceByIDV1WorkflowParameters   `json:"workflowParameters,omitempty"`   //
	DayZeroConfig        *ResponseDeviceOnboardingPnpGetDeviceByIDV1DayZeroConfig        `json:"dayZeroConfig,omitempty"`        //
	DayZeroConfigPreview *ResponseDeviceOnboardingPnpGetDeviceByIDV1DayZeroConfigPreview `json:"dayZeroConfigPreview,omitempty"` // Day Zero Config Preview
	Version              *float64                                                        `json:"version,omitempty"`              // Version
	TenantID             string                                                          `json:"tenantId,omitempty"`             // Tenant Id
}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfo struct {
	Source                    string                                                                      `json:"source,omitempty"`                    // Source
	SerialNumber              string                                                                      `json:"serialNumber,omitempty"`              // Serial Number
	Stack                     *bool                                                                       `json:"stack,omitempty"`                     // Stack
	Mode                      string                                                                      `json:"mode,omitempty"`                      // Mode
	State                     string                                                                      `json:"state,omitempty"`                     // State
	Location                  *ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoLocation               `json:"location,omitempty"`                  //
	Description               string                                                                      `json:"description,omitempty"`               // Description
	OnbState                  string                                                                      `json:"onbState,omitempty"`                  // Onb State
	AuthenticatedMicNumber    string                                                                      `json:"authenticatedMicNumber,omitempty"`    // Authenticated Mic Number
	AuthenticatedSudiSerialNo string                                                                      `json:"authenticatedSudiSerialNo,omitempty"` // Authenticated Sudi Serial No
	CapabilitiesSupported     []string                                                                    `json:"capabilitiesSupported,omitempty"`     // Capabilities Supported
	FeaturesSupported         []string                                                                    `json:"featuresSupported,omitempty"`         // Features Supported
	CmState                   string                                                                      `json:"cmState,omitempty"`                   // Cm State
	FirstContact              *float64                                                                    `json:"firstContact,omitempty"`              // First Contact
	LastContact               *float64                                                                    `json:"lastContact,omitempty"`               // Last Contact
	MacAddress                string                                                                      `json:"macAddress,omitempty"`                // Mac Address
	Pid                       string                                                                      `json:"pid,omitempty"`                       // Pid
	DeviceSudiSerialNos       []string                                                                    `json:"deviceSudiSerialNos,omitempty"`       // Device Sudi Serial Nos
	LastUpdateOn              *float64                                                                    `json:"lastUpdateOn,omitempty"`              // Last Update On
	WorkflowID                string                                                                      `json:"workflowId,omitempty"`                // Workflow Id
	WorkflowName              string                                                                      `json:"workflowName,omitempty"`              // Workflow Name
	ProjectID                 string                                                                      `json:"projectId,omitempty"`                 // Project Id
	ProjectName               string                                                                      `json:"projectName,omitempty"`               // Project Name
	DeviceType                string                                                                      `json:"deviceType,omitempty"`                // Device Type
	AgentType                 string                                                                      `json:"agentType,omitempty"`                 // Agent Type
	ImageVersion              string                                                                      `json:"imageVersion,omitempty"`              // Image Version
	FileSystemList            *[]ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoFileSystemList       `json:"fileSystemList,omitempty"`            //
	PnpProfileList            *[]ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoPnpProfileList       `json:"pnpProfileList,omitempty"`            //
	ImageFile                 string                                                                      `json:"imageFile,omitempty"`                 // Image File
	HTTPHeaders               *[]ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoHTTPHeaders          `json:"httpHeaders,omitempty"`               //
	NeighborLinks             *[]ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoNeighborLinks        `json:"neighborLinks,omitempty"`             //
	LastSyncTime              *float64                                                                    `json:"lastSyncTime,omitempty"`              // Last Sync Time
	IPInterfaces              *[]ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoIPInterfaces         `json:"ipInterfaces,omitempty"`              //
	Hostname                  string                                                                      `json:"hostname,omitempty"`                  // Hostname
	AuthStatus                string                                                                      `json:"authStatus,omitempty"`                // Auth Status
	StackInfo                 *ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoStackInfo              `json:"stackInfo,omitempty"`                 //
	ReloadRequested           *bool                                                                       `json:"reloadRequested,omitempty"`           // Reload Requested
	AddedOn                   *float64                                                                    `json:"addedOn,omitempty"`                   // Added On
	SiteID                    string                                                                      `json:"siteId,omitempty"`                    // Site Id
	AAACredentials            *ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoAAACredentials         `json:"aaaCredentials,omitempty"`            //
	UserMicNumbers            []string                                                                    `json:"userMicNumbers,omitempty"`            // User Mic Numbers
	UserSudiSerialNos         []string                                                                    `json:"userSudiSerialNos,omitempty"`         // User Sudi Serial Nos
	AddnMacAddrs              []string                                                                    `json:"addnMacAddrs,omitempty"`              // Addn Mac Addrs
	PreWorkflowCliOuputs      *[]ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoPreWorkflowCliOuputs `json:"preWorkflowCliOuputs,omitempty"`      //
	Tags                      *ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoTags                   `json:"tags,omitempty"`                      // Tags
	SudiRequired              *bool                                                                       `json:"sudiRequired,omitempty"`              // Sudi Required
	SmartAccountID            string                                                                      `json:"smartAccountId,omitempty"`            // Smart Account Id
	VirtualAccountID          string                                                                      `json:"virtualAccountId,omitempty"`          // Virtual Account Id
	PopulateInventory         *bool                                                                       `json:"populateInventory,omitempty"`         // Populate Inventory
	SiteName                  string                                                                      `json:"siteName,omitempty"`                  // Site Name
	Name                      string                                                                      `json:"name,omitempty"`                      // Name
}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoLocation struct {
	SiteID    string `json:"siteId,omitempty"`    // Site Id
	Address   string `json:"address,omitempty"`   // Address
	Latitude  string `json:"latitude,omitempty"`  // Latitude
	Longitude string `json:"longitude,omitempty"` // Longitude
	Altitude  string `json:"altitude,omitempty"`  // Altitude
}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoFileSystemList struct {
	Type      string   `json:"type,omitempty"`      // Type
	Writeable *bool    `json:"writeable,omitempty"` // Writeable
	Freespace *float64 `json:"freespace,omitempty"` // Freespace
	Name      string   `json:"name,omitempty"`      // Name
	Readable  *bool    `json:"readable,omitempty"`  // Readable
	Size      *float64 `json:"size,omitempty"`      // Size
}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoPnpProfileList struct {
	ProfileName       string                                                                               `json:"profileName,omitempty"`       // Profile Name
	DiscoveryCreated  *bool                                                                                `json:"discoveryCreated,omitempty"`  // Discovery Created
	CreatedBy         string                                                                               `json:"createdBy,omitempty"`         // Created By
	PrimaryEndpoint   *ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoPnpProfileListPrimaryEndpoint   `json:"primaryEndpoint,omitempty"`   //
	SecondaryEndpoint *ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoPnpProfileListSecondaryEndpoint `json:"secondaryEndpoint,omitempty"` //
}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoPnpProfileListPrimaryEndpoint struct {
	Port        *float64                                                                                      `json:"port,omitempty"`        // Port
	Protocol    string                                                                                        `json:"protocol,omitempty"`    // Protocol
	IPv4Address *ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoPnpProfileListPrimaryEndpointIPv4Address `json:"ipv4Address,omitempty"` // Ipv4 Address
	IPv6Address *ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoPnpProfileListPrimaryEndpointIPv6Address `json:"ipv6Address,omitempty"` // Ipv6 Address
	Fqdn        string                                                                                        `json:"fqdn,omitempty"`        // Fqdn
	Certificate string                                                                                        `json:"certificate,omitempty"` // Certificate
}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoPnpProfileListPrimaryEndpointIPv4Address interface{}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoPnpProfileListPrimaryEndpointIPv6Address interface{}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoPnpProfileListSecondaryEndpoint struct {
	Port        *float64                                                                                        `json:"port,omitempty"`        // Port
	Protocol    string                                                                                          `json:"protocol,omitempty"`    // Protocol
	IPv4Address *ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoPnpProfileListSecondaryEndpointIPv4Address `json:"ipv4Address,omitempty"` // Ipv4 Address
	IPv6Address *ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoPnpProfileListSecondaryEndpointIPv6Address `json:"ipv6Address,omitempty"` // Ipv6 Address
	Fqdn        string                                                                                          `json:"fqdn,omitempty"`        // Fqdn
	Certificate string                                                                                          `json:"certificate,omitempty"` // Certificate
}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoPnpProfileListSecondaryEndpointIPv4Address interface{}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoPnpProfileListSecondaryEndpointIPv6Address interface{}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoHTTPHeaders struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoNeighborLinks struct {
	LocalInterfaceName       string `json:"localInterfaceName,omitempty"`       // Local Interface Name
	LocalShortInterfaceName  string `json:"localShortInterfaceName,omitempty"`  // Local Short Interface Name
	LocalMacAddress          string `json:"localMacAddress,omitempty"`          // Local Mac Address
	RemoteInterfaceName      string `json:"remoteInterfaceName,omitempty"`      // Remote Interface Name
	RemoteShortInterfaceName string `json:"remoteShortInterfaceName,omitempty"` // Remote Short Interface Name
	RemoteMacAddress         string `json:"remoteMacAddress,omitempty"`         // Remote Mac Address
	RemoteDeviceName         string `json:"remoteDeviceName,omitempty"`         // Remote Device Name
	RemotePlatform           string `json:"remotePlatform,omitempty"`           // Remote Platform
	RemoteVersion            string `json:"remoteVersion,omitempty"`            // Remote Version
}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoIPInterfaces struct {
	Status          string                                                                             `json:"status,omitempty"`          // Status
	MacAddress      string                                                                             `json:"macAddress,omitempty"`      // Mac Address
	IPv4Address     *ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoIPInterfacesIPv4Address       `json:"ipv4Address,omitempty"`     // Ipv4 Address
	IPv6AddressList *[]ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoIPInterfacesIPv6AddressList `json:"ipv6AddressList,omitempty"` // Ipv6 Address List
	Name            string                                                                             `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoIPInterfacesIPv4Address interface{}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoIPInterfacesIPv6AddressList interface{}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoStackInfo struct {
	SupportsStackWorkflows *bool                                                                           `json:"supportsStackWorkflows,omitempty"` // Supports Stack Workflows
	IsFullRing             *bool                                                                           `json:"isFullRing,omitempty"`             // Is Full Ring
	StackMemberList        *[]ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoStackInfoStackMemberList `json:"stackMemberList,omitempty"`        //
	StackRingProtocol      string                                                                          `json:"stackRingProtocol,omitempty"`      // Stack Ring Protocol
	ValidLicenseLevels     []string                                                                        `json:"validLicenseLevels,omitempty"`     // Valid License Levels
	TotalMemberCount       *float64                                                                        `json:"totalMemberCount,omitempty"`       // Total Member Count
}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoStackInfoStackMemberList struct {
	SerialNumber     string   `json:"serialNumber,omitempty"`     // Serial Number
	State            string   `json:"state,omitempty"`            // State
	Role             string   `json:"role,omitempty"`             // Role
	MacAddress       string   `json:"macAddress,omitempty"`       // Mac Address
	Pid              string   `json:"pid,omitempty"`              // Pid
	LicenseLevel     string   `json:"licenseLevel,omitempty"`     // License Level
	LicenseType      string   `json:"licenseType,omitempty"`      // License Type
	SudiSerialNumber string   `json:"sudiSerialNumber,omitempty"` // Sudi Serial Number
	HardwareVersion  string   `json:"hardwareVersion,omitempty"`  // Hardware Version
	StackNumber      *float64 `json:"stackNumber,omitempty"`      // Stack Number
	SoftwareVersion  string   `json:"softwareVersion,omitempty"`  // Software Version
	Priority         *float64 `json:"priority,omitempty"`         // Priority
}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoAAACredentials struct {
	Password string `json:"password,omitempty"` // Password
	Username string `json:"username,omitempty"` // Username
}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoPreWorkflowCliOuputs struct {
	Cli       string `json:"cli,omitempty"`       // Cli
	CliOutput string `json:"cliOutput,omitempty"` // Cli Output
}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1DeviceInfoTags interface{}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1SystemResetWorkflow struct {
	TypeID         string                                                                `json:"_id,omitempty"`            // Id
	State          string                                                                `json:"state,omitempty"`          // State
	Type           string                                                                `json:"type,omitempty"`           // Type
	Description    string                                                                `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                                              `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                                `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                                              `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                                              `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpGetDeviceByIDV1SystemResetWorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                                 `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                                `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                                              `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                                              `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                                              `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                                `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                                `json:"configId,omitempty"`       // Config Id
	Name           string                                                                `json:"name,omitempty"`           // Name
	Version        *float64                                                              `json:"version,omitempty"`        // Version
	TenantID       string                                                                `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1SystemResetWorkflowTasks struct {
	State           string                                                                            `json:"state,omitempty"`           // State
	Type            string                                                                            `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                                          `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                                          `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                                          `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                                          `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpGetDeviceByIDV1SystemResetWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                                          `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                                            `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1SystemResetWorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1SystemWorkflow struct {
	TypeID         string                                                           `json:"_id,omitempty"`            // Id
	State          string                                                           `json:"state,omitempty"`          // State
	Type           string                                                           `json:"type,omitempty"`           // Type
	Description    string                                                           `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                                         `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                           `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                                         `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                                         `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpGetDeviceByIDV1SystemWorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                            `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                           `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                                         `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                                         `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                                         `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                           `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                           `json:"configId,omitempty"`       // Config Id
	Name           string                                                           `json:"name,omitempty"`           // Name
	Version        *float64                                                         `json:"version,omitempty"`        // Version
	TenantID       string                                                           `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1SystemWorkflowTasks struct {
	State           string                                                                       `json:"state,omitempty"`           // State
	Type            string                                                                       `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                                     `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                                     `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                                     `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                                     `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpGetDeviceByIDV1SystemWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                                     `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                                       `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1SystemWorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1Workflow struct {
	TypeID         string                                                     `json:"_id,omitempty"`            // Id
	State          string                                                     `json:"state,omitempty"`          // State
	Type           string                                                     `json:"type,omitempty"`           // Type
	Description    string                                                     `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                                   `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                     `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                                   `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                                   `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpGetDeviceByIDV1WorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                      `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                     `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                                   `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                                   `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                                   `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                     `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                     `json:"configId,omitempty"`       // Config Id
	Name           string                                                     `json:"name,omitempty"`           // Name
	Version        *float64                                                   `json:"version,omitempty"`        // Version
	TenantID       string                                                     `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1WorkflowTasks struct {
	State           string                                                                 `json:"state,omitempty"`           // State
	Type            string                                                                 `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                               `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                               `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                               `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                               `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpGetDeviceByIDV1WorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                               `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                                 `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1WorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1RunSummaryList struct {
	Details         string                                                                   `json:"details,omitempty"`         // Details
	HistoryTaskInfo *ResponseDeviceOnboardingPnpGetDeviceByIDV1RunSummaryListHistoryTaskInfo `json:"historyTaskInfo,omitempty"` //
	ErrorFlag       *bool                                                                    `json:"errorFlag,omitempty"`       // Error Flag
	Timestamp       *float64                                                                 `json:"timestamp,omitempty"`       // Timestamp
}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1RunSummaryListHistoryTaskInfo struct {
	Type         string                                                                                 `json:"type,omitempty"`         // Type
	WorkItemList *[]ResponseDeviceOnboardingPnpGetDeviceByIDV1RunSummaryListHistoryTaskInfoWorkItemList `json:"workItemList,omitempty"` //
	TimeTaken    *float64                                                                               `json:"timeTaken,omitempty"`    // Time Taken
	AddnDetails  *[]ResponseDeviceOnboardingPnpGetDeviceByIDV1RunSummaryListHistoryTaskInfoAddnDetails  `json:"addnDetails,omitempty"`  //
	Name         string                                                                                 `json:"name,omitempty"`         // Name
}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1RunSummaryListHistoryTaskInfoWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1RunSummaryListHistoryTaskInfoAddnDetails struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1WorkflowParameters struct {
	TopOfStackSerialNumber string                                                                    `json:"topOfStackSerialNumber,omitempty"` // Top Of Stack Serial Number
	LicenseLevel           string                                                                    `json:"licenseLevel,omitempty"`           // License Level
	LicenseType            string                                                                    `json:"licenseType,omitempty"`            // License Type
	ConfigList             *[]ResponseDeviceOnboardingPnpGetDeviceByIDV1WorkflowParametersConfigList `json:"configList,omitempty"`             //
}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1WorkflowParametersConfigList struct {
	ConfigParameters *[]ResponseDeviceOnboardingPnpGetDeviceByIDV1WorkflowParametersConfigListConfigParameters `json:"configParameters,omitempty"` //
	ConfigID         string                                                                                    `json:"configId,omitempty"`         // Config Id
}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1WorkflowParametersConfigListConfigParameters struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1DayZeroConfig struct {
	Config string `json:"config,omitempty"` // Config
}
type ResponseDeviceOnboardingPnpGetDeviceByIDV1DayZeroConfigPreview interface{}
type ResponseDeviceOnboardingPnpUpdatePnpGlobalSettingsV1 struct {
	SavaMappingList *[]ResponseDeviceOnboardingPnpUpdatePnpGlobalSettingsV1SavaMappingList `json:"savaMappingList,omitempty"` //
	TaskTimeOuts    *ResponseDeviceOnboardingPnpUpdatePnpGlobalSettingsV1TaskTimeOuts      `json:"taskTimeOuts,omitempty"`    //
	TenantID        string                                                                 `json:"tenantId,omitempty"`        // Tenant Id
	AAACredentials  *ResponseDeviceOnboardingPnpUpdatePnpGlobalSettingsV1AAACredentials    `json:"aaaCredentials,omitempty"`  //
	DefaultProfile  *ResponseDeviceOnboardingPnpUpdatePnpGlobalSettingsV1DefaultProfile    `json:"defaultProfile,omitempty"`  //
	AcceptEula      *bool                                                                  `json:"acceptEula,omitempty"`      // Accept Eula
	ID              string                                                                 `json:"id,omitempty"`              // Id
	Version         *float64                                                               `json:"version,omitempty"`         // Version
}
type ResponseDeviceOnboardingPnpUpdatePnpGlobalSettingsV1SavaMappingList struct {
	SyncStatus       string                                                                         `json:"syncStatus,omitempty"`       // Sync Status
	SyncStartTime    *float64                                                                       `json:"syncStartTime,omitempty"`    // Sync Start Time
	SyncResult       *ResponseDeviceOnboardingPnpUpdatePnpGlobalSettingsV1SavaMappingListSyncResult `json:"syncResult,omitempty"`       //
	LastSync         *float64                                                                       `json:"lastSync,omitempty"`         // Last Sync
	TenantID         string                                                                         `json:"tenantId,omitempty"`         // Tenant Id
	Profile          *ResponseDeviceOnboardingPnpUpdatePnpGlobalSettingsV1SavaMappingListProfile    `json:"profile,omitempty"`          //
	Token            string                                                                         `json:"token,omitempty"`            // Token
	Expiry           *float64                                                                       `json:"expiry,omitempty"`           // Expiry
	CcoUser          string                                                                         `json:"ccoUser,omitempty"`          // Cco User
	SmartAccountID   string                                                                         `json:"smartAccountId,omitempty"`   // Smart Account Id
	VirtualAccountID string                                                                         `json:"virtualAccountId,omitempty"` // Virtual Account Id
	AutoSyncPeriod   *float64                                                                       `json:"autoSyncPeriod,omitempty"`   // Auto Sync Period
	SyncResultStr    string                                                                         `json:"syncResultStr,omitempty"`    // Sync Result Str
}
type ResponseDeviceOnboardingPnpUpdatePnpGlobalSettingsV1SavaMappingListSyncResult struct {
	SyncList *[]ResponseDeviceOnboardingPnpUpdatePnpGlobalSettingsV1SavaMappingListSyncResultSyncList `json:"syncList,omitempty"` //
	SyncMsg  string                                                                                   `json:"syncMsg,omitempty"`  // Sync Msg
}
type ResponseDeviceOnboardingPnpUpdatePnpGlobalSettingsV1SavaMappingListSyncResultSyncList struct {
	SyncType     string   `json:"syncType,omitempty"`     // Sync Type
	DeviceSnList []string `json:"deviceSnList,omitempty"` // Device Sn List
}
type ResponseDeviceOnboardingPnpUpdatePnpGlobalSettingsV1SavaMappingListProfile struct {
	Port        *float64 `json:"port,omitempty"`        // Port
	AddressIPV4 string   `json:"addressIpV4,omitempty"` // Address Ip V4
	AddressFqdn string   `json:"addressFqdn,omitempty"` // Address Fqdn
	ProfileID   string   `json:"profileId,omitempty"`   // Profile Id
	Proxy       *bool    `json:"proxy,omitempty"`       // Proxy
	MakeDefault *bool    `json:"makeDefault,omitempty"` // Make Default
	Cert        string   `json:"cert,omitempty"`        // Cert
	Name        string   `json:"name,omitempty"`        // Name
}
type ResponseDeviceOnboardingPnpUpdatePnpGlobalSettingsV1TaskTimeOuts struct {
	ImageDownloadTimeOut *float64 `json:"imageDownloadTimeOut,omitempty"` // Image Download Time Out
	ConfigTimeOut        *float64 `json:"configTimeOut,omitempty"`        // Config Time Out
	GeneralTimeOut       *float64 `json:"generalTimeOut,omitempty"`       // General Time Out
}
type ResponseDeviceOnboardingPnpUpdatePnpGlobalSettingsV1AAACredentials struct {
	Password string `json:"password,omitempty"` // Password
	Username string `json:"username,omitempty"` // Username
}
type ResponseDeviceOnboardingPnpUpdatePnpGlobalSettingsV1DefaultProfile struct {
	FqdnAddresses []string `json:"fqdnAddresses,omitempty"` // Fqdn Addresses
	Proxy         *bool    `json:"proxy,omitempty"`         // Proxy
	Cert          string   `json:"cert,omitempty"`          // Cert
	IPAddresses   []string `json:"ipAddresses,omitempty"`   // Ip Addresses
	Port          *float64 `json:"port,omitempty"`          // Port
}
type ResponseDeviceOnboardingPnpGetPnpGlobalSettingsV1 struct {
	SavaMappingList *[]ResponseDeviceOnboardingPnpGetPnpGlobalSettingsV1SavaMappingList `json:"savaMappingList,omitempty"` //
	TaskTimeOuts    *ResponseDeviceOnboardingPnpGetPnpGlobalSettingsV1TaskTimeOuts      `json:"taskTimeOuts,omitempty"`    //
	TenantID        string                                                              `json:"tenantId,omitempty"`        // Tenant Id
	AAACredentials  *ResponseDeviceOnboardingPnpGetPnpGlobalSettingsV1AAACredentials    `json:"aaaCredentials,omitempty"`  //
	DefaultProfile  *ResponseDeviceOnboardingPnpGetPnpGlobalSettingsV1DefaultProfile    `json:"defaultProfile,omitempty"`  //
	AcceptEula      *bool                                                               `json:"acceptEula,omitempty"`      // Accept Eula
	ID              string                                                              `json:"id,omitempty"`              // Id
	Version         *float64                                                            `json:"version,omitempty"`         // Version
}
type ResponseDeviceOnboardingPnpGetPnpGlobalSettingsV1SavaMappingList struct {
	SyncStatus       string                                                                      `json:"syncStatus,omitempty"`       // Sync Status
	SyncStartTime    *float64                                                                    `json:"syncStartTime,omitempty"`    // Sync Start Time
	SyncResult       *ResponseDeviceOnboardingPnpGetPnpGlobalSettingsV1SavaMappingListSyncResult `json:"syncResult,omitempty"`       //
	LastSync         *float64                                                                    `json:"lastSync,omitempty"`         // Last Sync
	TenantID         string                                                                      `json:"tenantId,omitempty"`         // Tenant Id
	Profile          *ResponseDeviceOnboardingPnpGetPnpGlobalSettingsV1SavaMappingListProfile    `json:"profile,omitempty"`          //
	Token            string                                                                      `json:"token,omitempty"`            // Token
	Expiry           *float64                                                                    `json:"expiry,omitempty"`           // Expiry
	CcoUser          string                                                                      `json:"ccoUser,omitempty"`          // Cco User
	SmartAccountID   string                                                                      `json:"smartAccountId,omitempty"`   // Smart Account Id
	VirtualAccountID string                                                                      `json:"virtualAccountId,omitempty"` // Virtual Account Id
	AutoSyncPeriod   *float64                                                                    `json:"autoSyncPeriod,omitempty"`   // Auto Sync Period
	SyncResultStr    string                                                                      `json:"syncResultStr,omitempty"`    // Sync Result Str
}
type ResponseDeviceOnboardingPnpGetPnpGlobalSettingsV1SavaMappingListSyncResult struct {
	SyncList *[]ResponseDeviceOnboardingPnpGetPnpGlobalSettingsV1SavaMappingListSyncResultSyncList `json:"syncList,omitempty"` //
	SyncMsg  string                                                                                `json:"syncMsg,omitempty"`  // Sync Msg
}
type ResponseDeviceOnboardingPnpGetPnpGlobalSettingsV1SavaMappingListSyncResultSyncList struct {
	SyncType     string   `json:"syncType,omitempty"`     // Sync Type
	DeviceSnList []string `json:"deviceSnList,omitempty"` // Device Sn List
}
type ResponseDeviceOnboardingPnpGetPnpGlobalSettingsV1SavaMappingListProfile struct {
	Port        *float64 `json:"port,omitempty"`        // Port
	AddressIPV4 string   `json:"addressIpV4,omitempty"` // Address Ip V4
	AddressFqdn string   `json:"addressFqdn,omitempty"` // Address Fqdn
	ProfileID   string   `json:"profileId,omitempty"`   // Profile Id
	Proxy       *bool    `json:"proxy,omitempty"`       // Proxy
	MakeDefault *bool    `json:"makeDefault,omitempty"` // Make Default
	Cert        string   `json:"cert,omitempty"`        // Cert
	Name        string   `json:"name,omitempty"`        // Name
}
type ResponseDeviceOnboardingPnpGetPnpGlobalSettingsV1TaskTimeOuts struct {
	ImageDownloadTimeOut *float64 `json:"imageDownloadTimeOut,omitempty"` // Image Download Time Out
	ConfigTimeOut        *float64 `json:"configTimeOut,omitempty"`        // Config Time Out
	GeneralTimeOut       *float64 `json:"generalTimeOut,omitempty"`       // General Time Out
}
type ResponseDeviceOnboardingPnpGetPnpGlobalSettingsV1AAACredentials struct {
	Password string `json:"password,omitempty"` // Password
	Username string `json:"username,omitempty"` // Username
}
type ResponseDeviceOnboardingPnpGetPnpGlobalSettingsV1DefaultProfile struct {
	FqdnAddresses []string `json:"fqdnAddresses,omitempty"` // Fqdn Addresses
	Proxy         *bool    `json:"proxy,omitempty"`         // Proxy
	Cert          string   `json:"cert,omitempty"`          // Cert
	IPAddresses   []string `json:"ipAddresses,omitempty"`   // Ip Addresses
	Port          *float64 `json:"port,omitempty"`          // Port
}
type ResponseDeviceOnboardingPnpGetSmartAccountListV1 []string   // Array of ResponseDeviceOnboardingPnpGetSmartAccountListV1
type ResponseDeviceOnboardingPnpGetVirtualAccountListV1 []string // Array of ResponseDeviceOnboardingPnpGetVirtualAccountListV1
type ResponseDeviceOnboardingPnpAddVirtualAccountV1 struct {
	VirtualAccountID string                                                 `json:"virtualAccountId,omitempty"` // Virtual Account Id
	AutoSyncPeriod   *float64                                               `json:"autoSyncPeriod,omitempty"`   // Auto Sync Period
	Profile          *ResponseDeviceOnboardingPnpAddVirtualAccountV1Profile `json:"profile,omitempty"`          //
	CcoUser          string                                                 `json:"ccoUser,omitempty"`          // Cco User
	SyncStartTime    *float64                                               `json:"syncStartTime,omitempty"`    // Sync Start Time
	LastSync         *float64                                               `json:"lastSync,omitempty"`         // Last Sync
	TenantID         string                                                 `json:"tenantId,omitempty"`         // Tenant Id
	SmartAccountID   string                                                 `json:"smartAccountId,omitempty"`   // Smart Account Id
	Expiry           *float64                                               `json:"expiry,omitempty"`           // Expiry
	SyncStatus       string                                                 `json:"syncStatus,omitempty"`       // Sync Status
}
type ResponseDeviceOnboardingPnpAddVirtualAccountV1Profile struct {
	Proxy       *bool    `json:"proxy,omitempty"`       // Proxy
	MakeDefault *bool    `json:"makeDefault,omitempty"` // Make Default
	Port        *float64 `json:"port,omitempty"`        // Port
	ProfileID   string   `json:"profileId,omitempty"`   // Profile Id
	Name        string   `json:"name,omitempty"`        // Name
	AddressIPV4 string   `json:"addressIpV4,omitempty"` // Address Ip V4
	Cert        string   `json:"cert,omitempty"`        // Cert
	AddressFqdn string   `json:"addressFqdn,omitempty"` // Address Fqdn
}
type ResponseDeviceOnboardingPnpUpdatePnpServerProfileV1 struct {
	VirtualAccountID string                                                         `json:"virtualAccountId,omitempty"` // Virtual Account Id
	AutoSyncPeriod   *float64                                                       `json:"autoSyncPeriod,omitempty"`   // Auto Sync Period
	SyncResultStr    string                                                         `json:"syncResultStr,omitempty"`    // Sync Result Str
	Profile          *ResponseDeviceOnboardingPnpUpdatePnpServerProfileV1Profile    `json:"profile,omitempty"`          //
	CcoUser          string                                                         `json:"ccoUser,omitempty"`          // Cco User
	SyncResult       *ResponseDeviceOnboardingPnpUpdatePnpServerProfileV1SyncResult `json:"syncResult,omitempty"`       //
	Token            string                                                         `json:"token,omitempty"`            // Token
	SyncStartTime    *float64                                                       `json:"syncStartTime,omitempty"`    // Sync Start Time
	LastSync         *float64                                                       `json:"lastSync,omitempty"`         // Last Sync
	TenantID         string                                                         `json:"tenantId,omitempty"`         // Tenant Id
	SmartAccountID   string                                                         `json:"smartAccountId,omitempty"`   // Smart Account Id
	Expiry           *float64                                                       `json:"expiry,omitempty"`           // Expiry
	SyncStatus       string                                                         `json:"syncStatus,omitempty"`       // Sync Status
}
type ResponseDeviceOnboardingPnpUpdatePnpServerProfileV1Profile struct {
	Proxy       *bool    `json:"proxy,omitempty"`       // Proxy
	MakeDefault *bool    `json:"makeDefault,omitempty"` // Make Default
	Port        *float64 `json:"port,omitempty"`        // Port
	ProfileID   string   `json:"profileId,omitempty"`   // Profile Id
	Name        string   `json:"name,omitempty"`        // Name
	AddressIPV4 string   `json:"addressIpV4,omitempty"` // Address Ip V4
	Cert        string   `json:"cert,omitempty"`        // Cert
	AddressFqdn string   `json:"addressFqdn,omitempty"` // Address Fqdn
}
type ResponseDeviceOnboardingPnpUpdatePnpServerProfileV1SyncResult struct {
	SyncList *[]ResponseDeviceOnboardingPnpUpdatePnpServerProfileV1SyncResultSyncList `json:"syncList,omitempty"` //
	SyncMsg  string                                                                   `json:"syncMsg,omitempty"`  // Sync Msg
}
type ResponseDeviceOnboardingPnpUpdatePnpServerProfileV1SyncResultSyncList struct {
	SyncType     string   `json:"syncType,omitempty"`     // Sync Type
	DeviceSnList []string `json:"deviceSnList,omitempty"` // Device Sn List
}
type ResponseDeviceOnboardingPnpDeregisterVirtualAccountV1 struct {
	SmartAccountID   string                                                        `json:"smartAccountId,omitempty"`   // Smart Account Id
	VirtualAccountID string                                                        `json:"virtualAccountId,omitempty"` // Virtual Account Id
	LastSync         *float64                                                      `json:"lastSync,omitempty"`         // Last Sync
	CcoUser          string                                                        `json:"ccoUser,omitempty"`          // Cco User
	Expiry           *float64                                                      `json:"expiry,omitempty"`           // Expiry
	AutoSyncPeriod   *int                                                          `json:"autoSyncPeriod,omitempty"`   // Auto Sync Period
	Profile          *ResponseDeviceOnboardingPnpDeregisterVirtualAccountV1Profile `json:"profile,omitempty"`          //
	SyncStatus       string                                                        `json:"syncStatus,omitempty"`       // Sync Status
	SyncStartTime    *float64                                                      `json:"syncStartTime,omitempty"`    // Sync Start Time
	TenantID         string                                                        `json:"tenantId,omitempty"`         // Tenant Id
}
type ResponseDeviceOnboardingPnpDeregisterVirtualAccountV1Profile struct {
	Name        string `json:"name,omitempty"`        // Name
	ProfileID   string `json:"profileId,omitempty"`   // Profile Id
	MakeDefault *bool  `json:"makeDefault,omitempty"` // Make Default
	AddressIPV4 string `json:"addressIpV4,omitempty"` // Address Ip V4
	AddressIPV6 string `json:"addressIpV6,omitempty"` // Address Ip V6
	AddressFqdn string `json:"addressFqdn,omitempty"` // Address Fqdn
	Port        *int   `json:"port,omitempty"`        // Port
	Cert        string `json:"cert,omitempty"`        // Cert
	Proxy       *bool  `json:"proxy,omitempty"`       // Proxy
}
type ResponseDeviceOnboardingPnpGetWorkflowsV1 []ResponseItemDeviceOnboardingPnpGetWorkflowsV1 // Array of ResponseDeviceOnboardingPnpGetWorkflows
type ResponseItemDeviceOnboardingPnpGetWorkflowsV1 struct {
	TypeID         string                                                `json:"_id,omitempty"`            // Id
	State          string                                                `json:"state,omitempty"`          // State
	Type           string                                                `json:"type,omitempty"`           // Type
	Description    string                                                `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                              `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                              `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                              `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseItemDeviceOnboardingPnpGetWorkflowsV1Tasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                 `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                              `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                              `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                              `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                `json:"configId,omitempty"`       // Config Id
	Name           string                                                `json:"name,omitempty"`           // Name
	Version        *float64                                              `json:"version,omitempty"`        // Version
	TenantID       string                                                `json:"tenantId,omitempty"`       // Tenant Id
}

type ResponseDeviceOnboardingPnpGetWorkflowCount struct {
	Response *float64 `json:"response,omitempty"` // Response
}
type ResponseDeviceOnboardingPnpGetWorkflowByID struct {
	TypeID         string                                             `json:"_id,omitempty"`            // Id
	State          string                                             `json:"state,omitempty"`          // State
	Type           string                                             `json:"type,omitempty"`           // Type
	Description    string                                             `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                           `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                             `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                           `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                           `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpGetWorkflowByIDTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                              `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                             `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                           `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                           `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                           `json:"startTime,omitempty"`      // Start Time
	UseState       string                                             `json:"useState,omitempty"`       // Use State
	ConfigID       string                                             `json:"configId,omitempty"`       // Config Id
	Name           string                                             `json:"name,omitempty"`           // Name
	Version        *float64                                           `json:"version,omitempty"`        // Version
	TenantID       string                                             `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpGetWorkflowByIDTasks struct {
	State           string                                                         `json:"state,omitempty"`           // State
	Type            string                                                         `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                       `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                       `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                       `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                       `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpGetWorkflowByIDTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                       `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                         `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpGetWorkflowByIDTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpDeleteWorkflowV1ByID struct {
	TypeID         string                                                `json:"_id,omitempty"`            // Id
	State          string                                                `json:"state,omitempty"`          // State
	Type           string                                                `json:"type,omitempty"`           // Type
	Description    string                                                `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                              `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                              `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                              `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseItemDeviceOnboardingPnpGetWorkflowsV1Tasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                 `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                              `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                              `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                              `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                `json:"configId,omitempty"`       // Config Id
	Name           string                                                `json:"name,omitempty"`           // Name
	Version        *float64                                              `json:"version,omitempty"`        // Version
	TenantID       string                                                `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseItemDeviceOnboardingPnpGetWorkflowsV1Tasks struct {
	State           string                                                            `json:"state,omitempty"`           // State
	Type            string                                                            `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                          `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                          `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                          `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                          `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseItemDeviceOnboardingPnpGetWorkflowsV1TasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                          `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                            `json:"name,omitempty"`            // Name
}
type ResponseItemDeviceOnboardingPnpGetWorkflowsV1TasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpAddAWorkflowV1 struct {
	TypeID         string                                            `json:"_id,omitempty"`            // Id
	State          string                                            `json:"state,omitempty"`          // State
	Type           string                                            `json:"type,omitempty"`           // Type
	Description    string                                            `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                          `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                            `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                          `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                          `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpAddAWorkflowV1Tasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                             `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                            `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                          `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                          `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                          `json:"startTime,omitempty"`      // Start Time
	UseState       string                                            `json:"useState,omitempty"`       // Use State
	ConfigID       string                                            `json:"configId,omitempty"`       // Config Id
	Name           string                                            `json:"name,omitempty"`           // Name
	Version        *float64                                          `json:"version,omitempty"`        // Version
	TenantID       string                                            `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpAddAWorkflowV1Tasks struct {
	State           string                                                        `json:"state,omitempty"`           // State
	Type            string                                                        `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                      `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                      `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                      `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                      `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpAddAWorkflowV1TasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                      `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                        `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpAddAWorkflowV1TasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpGetWorkflowCountV1 struct {
	Response *float64 `json:"response,omitempty"` // Response
}
type ResponseDeviceOnboardingPnpGetWorkflowByIDV1 struct {
	TypeID         string                                               `json:"_id,omitempty"`            // Id
	State          string                                               `json:"state,omitempty"`          // State
	Type           string                                               `json:"type,omitempty"`           // Type
	Description    string                                               `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                             `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                               `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                             `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                             `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpGetWorkflowByIDV1Tasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                               `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                             `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                             `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                             `json:"startTime,omitempty"`      // Start Time
	UseState       string                                               `json:"useState,omitempty"`       // Use State
	ConfigID       string                                               `json:"configId,omitempty"`       // Config Id
	Name           string                                               `json:"name,omitempty"`           // Name
	Version        *float64                                             `json:"version,omitempty"`        // Version
	TenantID       string                                               `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpGetWorkflowByIDV1Tasks struct {
	State           string                                                           `json:"state,omitempty"`           // State
	Type            string                                                           `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                         `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                         `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                         `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                         `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpGetWorkflowByIDV1TasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                         `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                           `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpGetWorkflowByIDV1TasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpDeleteWorkflowByIDV1 struct {
	TypeID         string                                                  `json:"_id,omitempty"`            // Id
	State          string                                                  `json:"state,omitempty"`          // State
	Type           string                                                  `json:"type,omitempty"`           // Type
	Description    string                                                  `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                                `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                  `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                                `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                                `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpDeleteWorkflowByIDV1Tasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                   `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                  `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                                `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                                `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                                `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                  `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                  `json:"configId,omitempty"`       // Config Id
	Name           string                                                  `json:"name,omitempty"`           // Name
	Version        *float64                                                `json:"version,omitempty"`        // Version
	TenantID       string                                                  `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpDeleteWorkflowByIDV1Tasks struct {
	State           string                                                              `json:"state,omitempty"`           // State
	Type            string                                                              `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                            `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                            `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                            `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                            `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpDeleteWorkflowByIDV1TasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                            `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                              `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpDeleteWorkflowByIDV1TasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpUpdateWorkflowV1 struct {
	TypeID         string                                              `json:"_id,omitempty"`            // Id
	State          string                                              `json:"state,omitempty"`          // State
	Type           string                                              `json:"type,omitempty"`           // Type
	Description    string                                              `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                            `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                              `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                            `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                            `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpUpdateWorkflowV1Tasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                               `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                              `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                            `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                            `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                            `json:"startTime,omitempty"`      // Start Time
	UseState       string                                              `json:"useState,omitempty"`       // Use State
	ConfigID       string                                              `json:"configId,omitempty"`       // Config Id
	Name           string                                              `json:"name,omitempty"`           // Name
	Version        *float64                                            `json:"version,omitempty"`        // Version
	TenantID       string                                              `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpUpdateWorkflowV1Tasks struct {
	State           string                                                          `json:"state,omitempty"`           // State
	Type            string                                                          `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                        `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                        `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                        `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                        `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpUpdateWorkflowV1TasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                        `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                          `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpUpdateWorkflowV1TasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type RequestDeviceOnboardingPnpAuthorizeDeviceV1 struct {
	DeviceIDList []string `json:"deviceIdList,omitempty"` // Device Id List
}
type RequestDeviceOnboardingPnpAddDeviceV1 struct {
	DeviceInfo *RequestDeviceOnboardingPnpAddDeviceV1DeviceInfo `json:"deviceInfo,omitempty"` //
}
type RequestDeviceOnboardingPnpAddDeviceV1DeviceInfo struct {
	SerialNumber        string                                                    `json:"serialNumber,omitempty"`        // Serial Number
	Stack               *bool                                                     `json:"stack,omitempty"`               // Stack
	Description         string                                                    `json:"description,omitempty"`         // Description
	MacAddress          string                                                    `json:"macAddress,omitempty"`          // Mac Address
	Pid                 string                                                    `json:"pid,omitempty"`                 // Pid
	SiteID              string                                                    `json:"siteId,omitempty"`              // Site Id
	SudiRequired        *bool                                                     `json:"sudiRequired,omitempty"`        // Is Sudi Required
	DeviceSudiSerialNos []string                                                  `json:"deviceSudiSerialNos,omitempty"` // Device Sudi Serial Nos
	UserMicNumbers      []string                                                  `json:"userMicNumbers,omitempty"`      // User Mic Numbers
	UserSudiSerialNos   []string                                                  `json:"userSudiSerialNos,omitempty"`   // List of Secure Unique Device Identifier (SUDI) serial numbers to perform SUDI authorization, Required if sudiRequired is true.
	WorkflowID          string                                                    `json:"workflowId,omitempty"`          // Workflow Id
	WorkflowName        string                                                    `json:"workflowName,omitempty"`        // Workflow Name
	Hostname            string                                                    `json:"hostname,omitempty"`            // Hostname
	StackInfo           *RequestDeviceOnboardingPnpAddDeviceV1DeviceInfoStackInfo `json:"stackInfo,omitempty"`           //
}
type RequestDeviceOnboardingPnpAddDeviceV1DeviceInfoStackInfo struct {
	SupportsStackWorkflows *bool                                                                      `json:"supportsStackWorkflows,omitempty"` // Supports Stack Workflows
	IsFullRing             *bool                                                                      `json:"isFullRing,omitempty"`             // Is Full Ring
	StackMemberList        *[]RequestDeviceOnboardingPnpAddDeviceV1DeviceInfoStackInfoStackMemberList `json:"stackMemberList,omitempty"`        //
	StackRingProtocol      string                                                                     `json:"stackRingProtocol,omitempty"`      // Stack Ring Protocol
	ValidLicenseLevels     []string                                                                   `json:"validLicenseLevels,omitempty"`     // Valid License Levels
	TotalMemberCount       *float64                                                                   `json:"totalMemberCount,omitempty"`       // Total Member Count
}
type RequestDeviceOnboardingPnpAddDeviceV1DeviceInfoStackInfoStackMemberList struct {
	SerialNumber     string   `json:"serialNumber,omitempty"`     // Serial Number
	State            string   `json:"state,omitempty"`            // State
	Role             string   `json:"role,omitempty"`             // Role
	MacAddress       string   `json:"macAddress,omitempty"`       // Mac Address
	Pid              string   `json:"pid,omitempty"`              // Pid
	LicenseLevel     string   `json:"licenseLevel,omitempty"`     // License Level
	LicenseType      string   `json:"licenseType,omitempty"`      // License Type
	SudiSerialNumber string   `json:"sudiSerialNumber,omitempty"` // Sudi Serial Number
	HardwareVersion  string   `json:"hardwareVersion,omitempty"`  // Hardware Version
	StackNumber      *float64 `json:"stackNumber,omitempty"`      // Stack Number
	SoftwareVersion  string   `json:"softwareVersion,omitempty"`  // Software Version
	Priority         *float64 `json:"priority,omitempty"`         // Priority
}
type RequestDeviceOnboardingPnpClaimDeviceV1 struct {
	ConfigFileURL       string                                                    `json:"configFileUrl,omitempty"`       //
	ConfigID            string                                                    `json:"configId,omitempty"`            //
	DeviceClaimList     *[]RequestDeviceOnboardingPnpClaimDeviceV1DeviceClaimList `json:"deviceClaimList,omitempty"`     //
	FileServiceID       string                                                    `json:"fileServiceId,omitempty"`       //
	ImageID             string                                                    `json:"imageId,omitempty"`             //
	ImageURL            string                                                    `json:"imageUrl,omitempty"`            //
	PopulateInventory   *bool                                                     `json:"populateInventory,omitempty"`   //
	ProjectID           string                                                    `json:"projectId,omitempty"`           //
	WorkflowID          string                                                    `json:"workflowId,omitempty"`          //
	AuthorizationNeeded *bool                                                     `json:"authorizationNeeded,omitempty"` // Flag to enable/disable PnP device authorization. (true means enable)
}
type RequestDeviceOnboardingPnpClaimDeviceV1DeviceClaimList struct {
	ConfigList             *[]RequestDeviceOnboardingPnpClaimDeviceV1DeviceClaimListConfigList `json:"configList,omitempty"`             //
	DeviceID               string                                                              `json:"deviceId,omitempty"`               //
	LicenseLevel           string                                                              `json:"licenseLevel,omitempty"`           //
	LicenseType            string                                                              `json:"licenseType,omitempty"`            //
	TopOfStackSerialNumber string                                                              `json:"topOfStackSerialNumber,omitempty"` //
}
type RequestDeviceOnboardingPnpClaimDeviceV1DeviceClaimListConfigList struct {
	ConfigID         string                                                                              `json:"configId,omitempty"`         //
	ConfigParameters *[]RequestDeviceOnboardingPnpClaimDeviceV1DeviceClaimListConfigListConfigParameters `json:"configParameters,omitempty"` //
}
type RequestDeviceOnboardingPnpClaimDeviceV1DeviceClaimListConfigListConfigParameters struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}
type RequestDeviceOnboardingPnpImportDevicesInBulkV1 []RequestItemDeviceOnboardingPnpImportDevicesInBulkV1 // Array of RequestDeviceOnboardingPnpImportDevicesInBulkV1
type RequestItemDeviceOnboardingPnpImportDevicesInBulkV1 struct {
	TypeID     string                                                         `json:"_id,omitempty"`        // Id
	DeviceInfo *RequestItemDeviceOnboardingPnpImportDevicesInBulkV1DeviceInfo `json:"deviceInfo,omitempty"` //
}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkV1DeviceInfo struct {
	SerialNumber        string                                                                  `json:"serialNumber,omitempty"`        // Serial Number
	Stack               *bool                                                                   `json:"stack,omitempty"`               // Stack
	Description         string                                                                  `json:"description,omitempty"`         // Description
	MacAddress          string                                                                  `json:"macAddress,omitempty"`          // Mac Address
	Pid                 string                                                                  `json:"pid,omitempty"`                 // Pid
	SiteID              string                                                                  `json:"siteId,omitempty"`              // Site Id
	SudiRequired        *bool                                                                   `json:"sudiRequired,omitempty"`        // Is Sudi Required
	DeviceSudiSerialNos []string                                                                `json:"deviceSudiSerialNos,omitempty"` // Device Sudi Serial Nos
	UserMicNumbers      []string                                                                `json:"userMicNumbers,omitempty"`      // User Mic Numbers
	UserSudiSerialNos   []string                                                                `json:"userSudiSerialNos,omitempty"`   // User Sudi Serial Nos
	WorkflowID          string                                                                  `json:"workflowId,omitempty"`          // Workflow Id
	WorkflowName        string                                                                  `json:"workflowName,omitempty"`        // Workflow Name
	Hostname            string                                                                  `json:"hostname,omitempty"`            // Hostname
	StackInfo           *RequestItemDeviceOnboardingPnpImportDevicesInBulkV1DeviceInfoStackInfo `json:"stackInfo,omitempty"`           //
}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkV1DeviceInfoStackInfo struct {
	SupportsStackWorkflows *bool                                                                                    `json:"supportsStackWorkflows,omitempty"` // Supports Stack Workflows
	IsFullRing             *bool                                                                                    `json:"isFullRing,omitempty"`             // Is Full Ring
	StackMemberList        *[]RequestItemDeviceOnboardingPnpImportDevicesInBulkV1DeviceInfoStackInfoStackMemberList `json:"stackMemberList,omitempty"`        //
	StackRingProtocol      string                                                                                   `json:"stackRingProtocol,omitempty"`      // Stack Ring Protocol
	ValidLicenseLevels     []string                                                                                 `json:"validLicenseLevels,omitempty"`     // Valid License Levels
	TotalMemberCount       *float64                                                                                 `json:"totalMemberCount,omitempty"`       // Total Member Count
}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkV1DeviceInfoStackInfoStackMemberList struct {
	SerialNumber     string   `json:"serialNumber,omitempty"`     // Serial Number
	State            string   `json:"state,omitempty"`            // State
	Role             string   `json:"role,omitempty"`             // Role
	MacAddress       string   `json:"macAddress,omitempty"`       // Mac Address
	Pid              string   `json:"pid,omitempty"`              // Pid
	LicenseLevel     string   `json:"licenseLevel,omitempty"`     // License Level
	LicenseType      string   `json:"licenseType,omitempty"`      // License Type
	SudiSerialNumber string   `json:"sudiSerialNumber,omitempty"` // Sudi Serial Number
	HardwareVersion  string   `json:"hardwareVersion,omitempty"`  // Hardware Version
	StackNumber      *float64 `json:"stackNumber,omitempty"`      // Stack Number
	SoftwareVersion  string   `json:"softwareVersion,omitempty"`  // Software Version
	Priority         *float64 `json:"priority,omitempty"`         // Priority
}
type RequestDeviceOnboardingPnpResetDeviceV1 struct {
	DeviceResetList *[]RequestDeviceOnboardingPnpResetDeviceV1DeviceResetList `json:"deviceResetList,omitempty"` //
	ProjectID       string                                                    `json:"projectId,omitempty"`       //
	WorkflowID      string                                                    `json:"workflowId,omitempty"`      //
}
type RequestDeviceOnboardingPnpResetDeviceV1DeviceResetList struct {
	ConfigList             *[]RequestDeviceOnboardingPnpResetDeviceV1DeviceResetListConfigList `json:"configList,omitempty"`             //
	DeviceID               string                                                              `json:"deviceId,omitempty"`               //
	LicenseLevel           string                                                              `json:"licenseLevel,omitempty"`           //
	LicenseType            string                                                              `json:"licenseType,omitempty"`            //
	TopOfStackSerialNumber string                                                              `json:"topOfStackSerialNumber,omitempty"` //
}
type RequestDeviceOnboardingPnpResetDeviceV1DeviceResetListConfigList struct {
	ConfigID         string                                                                              `json:"configId,omitempty"`         //
	ConfigParameters *[]RequestDeviceOnboardingPnpResetDeviceV1DeviceResetListConfigListConfigParameters `json:"configParameters,omitempty"` //
}
type RequestDeviceOnboardingPnpResetDeviceV1DeviceResetListConfigListConfigParameters struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}
type RequestDeviceOnboardingPnpClaimADeviceToASiteV1 struct {
	DeviceID        string                                                     `json:"deviceId,omitempty"`        // Device Id
	SiteID          string                                                     `json:"siteId,omitempty"`          // Site Id
	Type            string                                                     `json:"type,omitempty"`            // Type
	ImageInfo       *RequestDeviceOnboardingPnpClaimADeviceToASiteV1ImageInfo  `json:"imageInfo,omitempty"`       //
	ConfigInfo      *RequestDeviceOnboardingPnpClaimADeviceToASiteV1ConfigInfo `json:"configInfo,omitempty"`      //
	RfProfile       string                                                     `json:"rfProfile,omitempty"`       // for Access Points
	StaticIP        string                                                     `json:"staticIP,omitempty"`        // for CatalystWLC/MobilityExpress
	SubnetMask      string                                                     `json:"subnetMask,omitempty"`      // for CatalystWLC/MobilityExpress
	Gateway         string                                                     `json:"gateway,omitempty"`         // for CatalystWLC/MobilityExpress
	VLANID          string                                                     `json:"vlanId,omitempty"`          // for Catalyst 9800 WLC
	IPInterfaceName string                                                     `json:"ipInterfaceName,omitempty"` // for Catalyst 9800 WLC
	SensorProfile   string                                                     `json:"sensorProfile,omitempty"`   // for Sensors
	Hostname        string                                                     `json:"hostname,omitempty"`        // hostname to configure on Device.
}
type RequestDeviceOnboardingPnpClaimADeviceToASiteV1ImageInfo struct {
	ImageID string `json:"imageId,omitempty"` // Image Id
	Skip    *bool  `json:"skip,omitempty"`    // Skip
}
type RequestDeviceOnboardingPnpClaimADeviceToASiteV1ConfigInfo struct {
	ConfigID         string                                                                       `json:"configId,omitempty"`         // Config Id
	ConfigParameters *[]RequestDeviceOnboardingPnpClaimADeviceToASiteV1ConfigInfoConfigParameters `json:"configParameters,omitempty"` //
}
type RequestDeviceOnboardingPnpClaimADeviceToASiteV1ConfigInfoConfigParameters struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type RequestDeviceOnboardingPnpPreviewConfigV1 struct {
	DeviceID string `json:"deviceId,omitempty"` //
	SiteID   string `json:"siteId,omitempty"`   //
	Type     string `json:"type,omitempty"`     //
}
type RequestDeviceOnboardingPnpUnClaimDeviceV1 struct {
	DeviceIDList []string `json:"deviceIdList,omitempty"` //
}
type RequestDeviceOnboardingPnpSyncVirtualAccountDevicesV1 struct {
	AutoSyncPeriod   *int                                                             `json:"autoSyncPeriod,omitempty"`   //
	CcoUser          string                                                           `json:"ccoUser,omitempty"`          //
	Expiry           *int                                                             `json:"expiry,omitempty"`           //
	LastSync         *int                                                             `json:"lastSync,omitempty"`         //
	Profile          *RequestDeviceOnboardingPnpSyncVirtualAccountDevicesV1Profile    `json:"profile,omitempty"`          //
	SmartAccountID   string                                                           `json:"smartAccountId,omitempty"`   //
	SyncResult       *RequestDeviceOnboardingPnpSyncVirtualAccountDevicesV1SyncResult `json:"syncResult,omitempty"`       //
	SyncResultStr    string                                                           `json:"syncResultStr,omitempty"`    //
	SyncStartTime    *int                                                             `json:"syncStartTime,omitempty"`    //
	SyncStatus       string                                                           `json:"syncStatus,omitempty"`       //
	TenantID         string                                                           `json:"tenantId,omitempty"`         //
	Token            string                                                           `json:"token,omitempty"`            //
	VirtualAccountID string                                                           `json:"virtualAccountId,omitempty"` //
}
type RequestDeviceOnboardingPnpSyncVirtualAccountDevicesV1Profile struct {
	AddressFqdn string `json:"addressFqdn,omitempty"` //
	AddressIPV4 string `json:"addressIpV4,omitempty"` //
	Cert        string `json:"cert,omitempty"`        //
	MakeDefault *bool  `json:"makeDefault,omitempty"` //
	Name        string `json:"name,omitempty"`        //
	Port        *int   `json:"port,omitempty"`        //
	ProfileID   string `json:"profileId,omitempty"`   //
	Proxy       *bool  `json:"proxy,omitempty"`       //
}
type RequestDeviceOnboardingPnpSyncVirtualAccountDevicesV1SyncResult struct {
	SyncList *[]RequestDeviceOnboardingPnpSyncVirtualAccountDevicesV1SyncResultSyncList `json:"syncList,omitempty"` //
	SyncMsg  string                                                                     `json:"syncMsg,omitempty"`  //
}
type RequestDeviceOnboardingPnpSyncVirtualAccountDevicesV1SyncResultSyncList struct {
	DeviceSnList []string `json:"deviceSnList,omitempty"` //
	SyncType     string   `json:"syncType,omitempty"`     //
}
type RequestDeviceOnboardingPnpUpdateDeviceV1 struct {
	ID         string                                              `json:"id,omitempty"`         // Id
	DeviceInfo *RequestDeviceOnboardingPnpUpdateDeviceV1DeviceInfo `json:"deviceInfo,omitempty"` //
}
type RequestDeviceOnboardingPnpUpdateDeviceV1DeviceInfo struct {
	Hostname          string   `json:"hostname,omitempty"`          // Hostname
	SerialNumber      string   `json:"serialNumber,omitempty"`      // Serial Number
	Pid               string   `json:"pid,omitempty"`               // Pid
	SudiRequired      *bool    `json:"sudiRequired,omitempty"`      // Sudi Required
	UserSudiSerialNos []string `json:"userSudiSerialNos,omitempty"` // List of Secure Unique Device Identifier (SUDI) serial numbers to perform SUDI authorization, Required if sudiRequired is true.
	Stack             *bool    `json:"stack,omitempty"`             // Stack
}
type RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsV1 struct {
	ID              string                                                                `json:"id,omitempty"`              // Id
	AcceptEula      string                                                                `json:"acceptEula,omitempty"`      // Accept Eula
	DefaultProfile  *RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsV1DefaultProfile    `json:"defaultProfile,omitempty"`  //
	SavaMappingList *[]RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsV1SavaMappingList `json:"savaMappingList,omitempty"` //
}
type RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsV1DefaultProfile struct {
	Cert          string   `json:"cert,omitempty"`          // Cert
	FqdnAddresses []string `json:"fqdnAddresses,omitempty"` // Fqdn Addresses
	IPAddresses   []string `json:"ipAddresses,omitempty"`   // Ip Addresses
	Port          string   `json:"port,omitempty"`          // Port
	Proxy         string   `json:"proxy,omitempty"`         // Proxy
}
type RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsV1SavaMappingList struct {
	CcoUser          string                                                                     `json:"ccoUser,omitempty"`          // Cco User
	Expiry           string                                                                     `json:"expiry,omitempty"`           // Expiry
	Profile          *RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsV1SavaMappingListProfile `json:"profile,omitempty"`          //
	SmartAccountID   string                                                                     `json:"smartAccountId,omitempty"`   // Smart Account Id
	VirtualAccountID string                                                                     `json:"virtualAccountId,omitempty"` // Virtual Account Id
}
type RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsV1SavaMappingListProfile struct {
	AddressFqdn string `json:"addressFqdn,omitempty"` // Address Fqdn
	AddressIPV4 string `json:"addressIpV4,omitempty"` // Address Ip V4
	Cert        string `json:"cert,omitempty"`        // Cert
	MakeDefault string `json:"makeDefault,omitempty"` // Make Default
	Name        string `json:"name,omitempty"`        // Name
	Port        string `json:"port,omitempty"`        // Port
	ProfileID   string `json:"profileId,omitempty"`   // Profile Id
	Proxy       string `json:"proxy,omitempty"`       // Proxy
}
type RequestDeviceOnboardingPnpAddVirtualAccountV1 struct {
	AutoSyncPeriod   *int                                                     `json:"autoSyncPeriod,omitempty"`   //
	CcoUser          string                                                   `json:"ccoUser,omitempty"`          //
	Expiry           *int                                                     `json:"expiry,omitempty"`           //
	LastSync         *int                                                     `json:"lastSync,omitempty"`         //
	Profile          *RequestDeviceOnboardingPnpAddVirtualAccountV1Profile    `json:"profile,omitempty"`          //
	SmartAccountID   string                                                   `json:"smartAccountId,omitempty"`   //
	SyncResult       *RequestDeviceOnboardingPnpAddVirtualAccountV1SyncResult `json:"syncResult,omitempty"`       //
	SyncResultStr    string                                                   `json:"syncResultStr,omitempty"`    // Represent internal state and SHOULD not be used or relied upon. (Deprecated)
	SyncStartTime    *int                                                     `json:"syncStartTime,omitempty"`    //
	SyncStatus       string                                                   `json:"syncStatus,omitempty"`       // Represent internal state and SHOULD not be used or relied upon. (Deprecated)
	TenantID         string                                                   `json:"tenantId,omitempty"`         // Represent internal state and SHOULD not be used or relied upon. (Deprecated)
	Token            string                                                   `json:"token,omitempty"`            // Represent internal state and SHOULD not be used or relied upon. (Deprecated)
	VirtualAccountID string                                                   `json:"virtualAccountId,omitempty"` //
}
type RequestDeviceOnboardingPnpAddVirtualAccountV1Profile struct {
	AddressFqdn string `json:"addressFqdn,omitempty"` // Required when cluster is configured with fully qualified domain name (FQDN)
	AddressIPV4 string `json:"addressIpV4,omitempty"` // Required when cluster is configured with IPv4
	AddressIPV6 string `json:"addressIpV6,omitempty"` // Required when cluster is configured with IPv6
	Cert        string `json:"cert,omitempty"`        //
	MakeDefault *bool  `json:"makeDefault,omitempty"` //
	Name        string `json:"name,omitempty"`        //
	Port        *int   `json:"port,omitempty"`        //
	ProfileID   string `json:"profileId,omitempty"`   //
	Proxy       *bool  `json:"proxy,omitempty"`       //
}
type RequestDeviceOnboardingPnpAddVirtualAccountV1SyncResult struct {
	SyncList *[]RequestDeviceOnboardingPnpAddVirtualAccountV1SyncResultSyncList `json:"syncList,omitempty"` //
	SyncMsg  string                                                             `json:"syncMsg,omitempty"`  //
}
type RequestDeviceOnboardingPnpAddVirtualAccountV1SyncResultSyncList struct {
	DeviceSnList []string `json:"deviceSnList,omitempty"` //
	SyncType     string   `json:"syncType,omitempty"`     //
}
type RequestDeviceOnboardingPnpUpdatePnpServerProfileV1 struct {
	SmartAccountID   string                                                     `json:"smartAccountId,omitempty"`   // Smart Account Id
	VirtualAccountID string                                                     `json:"virtualAccountId,omitempty"` // Virtual Account Id
	Profile          *RequestDeviceOnboardingPnpUpdatePnpServerProfileV1Profile `json:"profile,omitempty"`          //
	CcoUser          string                                                     `json:"ccoUser,omitempty"`          // Cco User
}
type RequestDeviceOnboardingPnpUpdatePnpServerProfileV1Profile struct {
	Proxy       *bool    `json:"proxy,omitempty"`       // Proxy
	MakeDefault *bool    `json:"makeDefault,omitempty"` // Make Default
	Port        *float64 `json:"port,omitempty"`        // Port
	ProfileID   string   `json:"profileId,omitempty"`   // Profile Id
	Name        string   `json:"name,omitempty"`        // Name
	AddressFqdn string   `json:"addressFqdn,omitempty"` // Required when cluster is configured with fully qualified domain name (FQDN)
	AddressIPV4 string   `json:"addressIpV4,omitempty"` // Required when cluster is configured with IPv4
	AddressIPV6 string   `json:"addressIpV6,omitempty"` // Required when cluster is configured with IPv6
	Cert        string   `json:"cert,omitempty"`        // Cert
}
type RequestDeviceOnboardingPnpAddAWorkflowV1 struct {
	TypeID         string                                           `json:"_id,omitempty"`            //
	AddToInventory *bool                                            `json:"addToInventory,omitempty"` //
	AddedOn        *int                                             `json:"addedOn,omitempty"`        //
	ConfigID       string                                           `json:"configId,omitempty"`       //
	CurrTaskIDx    *int                                             `json:"currTaskIdx,omitempty"`    //
	Description    string                                           `json:"description,omitempty"`    //
	EndTime        *int                                             `json:"endTime,omitempty"`        //
	ExecTime       *int                                             `json:"execTime,omitempty"`       //
	ImageID        string                                           `json:"imageId,omitempty"`        //
	InstanceType   string                                           `json:"instanceType,omitempty"`   //
	LastupdateOn   *int                                             `json:"lastupdateOn,omitempty"`   //
	Name           string                                           `json:"name,omitempty"`           //
	StartTime      *int                                             `json:"startTime,omitempty"`      //
	State          string                                           `json:"state,omitempty"`          //
	Tasks          *[]RequestDeviceOnboardingPnpAddAWorkflowV1Tasks `json:"tasks,omitempty"`          //
	TenantID       string                                           `json:"tenantId,omitempty"`       //
	Type           string                                           `json:"type,omitempty"`           //
	UseState       string                                           `json:"useState,omitempty"`       //
	Version        *int                                             `json:"version,omitempty"`        //
}
type RequestDeviceOnboardingPnpAddAWorkflowV1Tasks struct {
	CurrWorkItemIDx *int                                                         `json:"currWorkItemIdx,omitempty"` //
	EndTime         *int                                                         `json:"endTime,omitempty"`         //
	Name            string                                                       `json:"name,omitempty"`            //
	StartTime       *int                                                         `json:"startTime,omitempty"`       //
	State           string                                                       `json:"state,omitempty"`           //
	TaskSeqNo       *int                                                         `json:"taskSeqNo,omitempty"`       //
	TimeTaken       *int                                                         `json:"timeTaken,omitempty"`       //
	Type            string                                                       `json:"type,omitempty"`            //
	WorkItemList    *[]RequestDeviceOnboardingPnpAddAWorkflowV1TasksWorkItemList `json:"workItemList,omitempty"`    //
}
type RequestDeviceOnboardingPnpAddAWorkflowV1TasksWorkItemList struct {
	Command   string `json:"command,omitempty"`   //
	EndTime   *int   `json:"endTime,omitempty"`   //
	OutputStr string `json:"outputStr,omitempty"` //
	StartTime *int   `json:"startTime,omitempty"` //
	State     string `json:"state,omitempty"`     //
	TimeTaken *int   `json:"timeTaken,omitempty"` //
}
type RequestDeviceOnboardingPnpUpdateWorkflowV1 struct {
	TypeID         string                                             `json:"_id,omitempty"`            //
	AddToInventory *bool                                              `json:"addToInventory,omitempty"` //
	AddedOn        *int                                               `json:"addedOn,omitempty"`        //
	ConfigID       string                                             `json:"configId,omitempty"`       //
	CurrTaskIDx    *int                                               `json:"currTaskIdx,omitempty"`    //
	Description    string                                             `json:"description,omitempty"`    //
	EndTime        *int                                               `json:"endTime,omitempty"`        //
	ExecTime       *int                                               `json:"execTime,omitempty"`       //
	ImageID        string                                             `json:"imageId,omitempty"`        //
	InstanceType   string                                             `json:"instanceType,omitempty"`   //
	LastupdateOn   *int                                               `json:"lastupdateOn,omitempty"`   //
	Name           string                                             `json:"name,omitempty"`           //
	StartTime      *int                                               `json:"startTime,omitempty"`      //
	State          string                                             `json:"state,omitempty"`          //
	Tasks          *[]RequestDeviceOnboardingPnpUpdateWorkflowV1Tasks `json:"tasks,omitempty"`          //
	TenantID       string                                             `json:"tenantId,omitempty"`       //
	Type           string                                             `json:"type,omitempty"`           //
	UseState       string                                             `json:"useState,omitempty"`       //
	Version        *int                                               `json:"version,omitempty"`        //
}
type RequestDeviceOnboardingPnpUpdateWorkflowV1Tasks struct {
	CurrWorkItemIDx *int                                                           `json:"currWorkItemIdx,omitempty"` //
	EndTime         *int                                                           `json:"endTime,omitempty"`         //
	Name            string                                                         `json:"name,omitempty"`            //
	StartTime       *int                                                           `json:"startTime,omitempty"`       //
	State           string                                                         `json:"state,omitempty"`           //
	TaskSeqNo       *int                                                           `json:"taskSeqNo,omitempty"`       //
	TimeTaken       *int                                                           `json:"timeTaken,omitempty"`       //
	Type            string                                                         `json:"type,omitempty"`            //
	WorkItemList    *[]RequestDeviceOnboardingPnpUpdateWorkflowV1TasksWorkItemList `json:"workItemList,omitempty"`    //
}
type RequestDeviceOnboardingPnpUpdateWorkflowV1TasksWorkItemList struct {
	Command   string `json:"command,omitempty"`   //
	EndTime   *int   `json:"endTime,omitempty"`   //
	OutputStr string `json:"outputStr,omitempty"` //
	StartTime *int   `json:"startTime,omitempty"` //
	State     string `json:"state,omitempty"`     //
	TimeTaken *int   `json:"timeTaken,omitempty"` //
}

//GetDeviceListSiteManagementV1 Get Device list - e6b3-db80-46c9-9654
/* Returns list of devices from Plug & Play based on filter criteria. Returns 50 devices by default. This endpoint supports Pagination and Sorting.


@param GetDeviceListSiteManagementV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-list-site-management-v1
*/
func (s *DeviceOnboardingPnpService) GetDeviceListSiteManagementV1(GetDeviceListSiteManagementV1QueryParams *GetDeviceListSiteManagementV1QueryParams) (*ResponseDeviceOnboardingPnpGetDeviceListSiteManagementV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-device"

	queryString, _ := query.Values(GetDeviceListSiteManagementV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDeviceOnboardingPnpGetDeviceListSiteManagementV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceListSiteManagementV1(GetDeviceListSiteManagementV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceListSiteManagementV1")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpGetDeviceListSiteManagementV1)
	return result, response, err

}

//GetDeviceCountV1 Get Device Count - d9a1-fa9c-4068-b23c
/* Returns the device count based on filter criteria. This is useful for pagination


@param GetDeviceCountV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-count-v1
*/
func (s *DeviceOnboardingPnpService) GetDeviceCountV1(GetDeviceCountV1QueryParams *GetDeviceCountV1QueryParams) (*ResponseDeviceOnboardingPnpGetDeviceCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-device/count"

	queryString, _ := query.Values(GetDeviceCountV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDeviceOnboardingPnpGetDeviceCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceCountV1(GetDeviceCountV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceCountV1")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpGetDeviceCountV1)
	return result, response, err

}

//GetDeviceHistoryV1 Get Device History - f093-1967-4049-a7d4
/* Returns history for a specific device. Serial number is a required parameter


@param GetDeviceHistoryV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-history-v1
*/
func (s *DeviceOnboardingPnpService) GetDeviceHistoryV1(GetDeviceHistoryV1QueryParams *GetDeviceHistoryV1QueryParams) (*ResponseDeviceOnboardingPnpGetDeviceHistoryV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-device/history"

	queryString, _ := query.Values(GetDeviceHistoryV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDeviceOnboardingPnpGetDeviceHistoryV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceHistoryV1(GetDeviceHistoryV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceHistoryV1")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpGetDeviceHistoryV1)
	return result, response, err

}

//GetSyncResultForVirtualAccountV1 Get Sync Result for Virtual Account - 0a9c-9884-45cb-91c8
/* Returns the summary of devices synced from the given smart account & virtual account with PnP (Deprecated)


@param domain domain path parameter. Smart Account Domain

@param name name path parameter. Virtual Account Name


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-sync-result-for-virtual-account-v1
*/
func (s *DeviceOnboardingPnpService) GetSyncResultForVirtualAccountV1(domain string, name string) (*ResponseDeviceOnboardingPnpGetSyncResultForVirtualAccountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-device/sacct/{domain}/vacct/{name}/sync-result"
	path = strings.Replace(path, "{domain}", fmt.Sprintf("%v", domain), -1)
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceOnboardingPnpGetSyncResultForVirtualAccountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSyncResultForVirtualAccountV1(domain, name)
		}
		return nil, response, fmt.Errorf("error with operation GetSyncResultForVirtualAccountV1")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpGetSyncResultForVirtualAccountV1)
	return result, response, err

}

//GetDeviceByIDV1 Get Device by Id - bab6-c9e5-4408-85cc
/* Returns device details specified by device id


@param id id path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-by-id-v1
*/
func (s *DeviceOnboardingPnpService) GetDeviceByIDV1(id string) (*ResponseDeviceOnboardingPnpGetDeviceByIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-device/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceOnboardingPnpGetDeviceByIDV1{}).
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

	result := response.Result().(*ResponseDeviceOnboardingPnpGetDeviceByIDV1)
	return result, response, err

}

//GetPnpGlobalSettingsV1 Get PnP global settings - 7e92-f9eb-46db-8320
/* Returns global PnP settings of the user



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-pnp-global-settings-v1
*/
func (s *DeviceOnboardingPnpService) GetPnpGlobalSettingsV1() (*ResponseDeviceOnboardingPnpGetPnpGlobalSettingsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-settings"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceOnboardingPnpGetPnpGlobalSettingsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPnpGlobalSettingsV1()
		}
		return nil, response, fmt.Errorf("error with operation GetPnpGlobalSettingsV1")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpGetPnpGlobalSettingsV1)
	return result, response, err

}

//GetSmartAccountListV1 Get Smart Account List - 3cb2-4acb-486b-89d2
/* Returns the list of Smart Account domains



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-smart-account-list-v1
*/
func (s *DeviceOnboardingPnpService) GetSmartAccountListV1() (*ResponseDeviceOnboardingPnpGetSmartAccountListV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-settings/sacct"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceOnboardingPnpGetSmartAccountListV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSmartAccountListV1()
		}
		return nil, response, fmt.Errorf("error with operation GetSmartAccountListV1")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpGetSmartAccountListV1)
	return result, response, err

}

//GetVirtualAccountListV1 Get Virtual Account List - 70a4-79a6-462a-9496
/* Returns list of virtual accounts associated with the specified smart account


@param domain domain path parameter. Smart Account Domain


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-virtual-account-list-v1
*/
func (s *DeviceOnboardingPnpService) GetVirtualAccountListV1(domain string) (*ResponseDeviceOnboardingPnpGetVirtualAccountListV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-settings/sacct/{domain}/vacct"
	path = strings.Replace(path, "{domain}", fmt.Sprintf("%v", domain), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceOnboardingPnpGetVirtualAccountListV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetVirtualAccountListV1(domain)
		}
		return nil, response, fmt.Errorf("error with operation GetVirtualAccountListV1")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpGetVirtualAccountListV1)
	return result, response, err

}

//GetWorkflowsV1 Get Workflows - aeb4-dad0-4a99-bbe3
/* Returns the list of workflows based on filter criteria. If a limit is not specified, it will default to return 50 workflows. Pagination and sorting are also supported by this endpoint


@param GetWorkflowsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-workflows-v1
*/
func (s *DeviceOnboardingPnpService) GetWorkflowsV1(GetWorkflowsV1QueryParams *GetWorkflowsV1QueryParams) (*ResponseDeviceOnboardingPnpGetWorkflowsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-workflow"

	queryString, _ := query.Values(GetWorkflowsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDeviceOnboardingPnpGetWorkflowsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetWorkflowsV1(GetWorkflowsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetWorkflowsV1")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpGetWorkflowsV1)
	return result, response, err

}

//GetWorkflowCountV1 Get Workflow Count - 7989-f868-46fa-af99
/* Returns the workflow count


@param GetWorkflowCountV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-workflow-count-v1
*/
func (s *DeviceOnboardingPnpService) GetWorkflowCountV1(GetWorkflowCountV1QueryParams *GetWorkflowCountV1QueryParams) (*ResponseDeviceOnboardingPnpGetWorkflowCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-workflow/count"

	queryString, _ := query.Values(GetWorkflowCountV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDeviceOnboardingPnpGetWorkflowCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetWorkflowCountV1(GetWorkflowCountV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetWorkflowCountV1")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpGetWorkflowCountV1)
	return result, response, err

}

//GetWorkflowByIDV1 Get Workflow by Id - 80ac-b88e-4ac9-ac6d
/* Returns a workflow specified by id


@param id id path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-workflow-by-id-v1
*/
func (s *DeviceOnboardingPnpService) GetWorkflowByIDV1(id string) (*ResponseDeviceOnboardingPnpGetWorkflowByIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-workflow/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceOnboardingPnpGetWorkflowByIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetWorkflowByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation GetWorkflowByIdV1")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpGetWorkflowByIDV1)
	return result, response, err

}

//AuthorizeDeviceV1 Authorize Device - 2897-4ae4-4ae9-a1dc
/* Authorizes one of more devices. A device can only be authorized if Authorization is set in Device Settings.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!authorize-device-v1
*/
func (s *DeviceOnboardingPnpService) AuthorizeDeviceV1(requestDeviceOnboardingPnpAuthorizeDeviceV1 *RequestDeviceOnboardingPnpAuthorizeDeviceV1) (*ResponseDeviceOnboardingPnpAuthorizeDeviceV1, *resty.Response, error) {
	path := "/api/v1/onboarding/pnp-device/authorize"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceOnboardingPnpAuthorizeDeviceV1).
		SetResult(&ResponseDeviceOnboardingPnpAuthorizeDeviceV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AuthorizeDeviceV1(requestDeviceOnboardingPnpAuthorizeDeviceV1)
		}

		return nil, response, fmt.Errorf("error with operation AuthorizeDeviceV1")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpAuthorizeDeviceV1)
	return result, response, err

}

//AddDeviceV1 Add Device - f3b2-6b55-44ca-bab9
/* Adds a device to the PnP database.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-device-v1
*/
func (s *DeviceOnboardingPnpService) AddDeviceV1(requestDeviceOnboardingPnpAddDeviceV1 *RequestDeviceOnboardingPnpAddDeviceV1) (*ResponseDeviceOnboardingPnpAddDeviceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-device"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceOnboardingPnpAddDeviceV1).
		SetResult(&ResponseDeviceOnboardingPnpAddDeviceV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddDeviceV1(requestDeviceOnboardingPnpAddDeviceV1)
		}

		return nil, response, fmt.Errorf("error with operation AddDeviceV1")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpAddDeviceV1)
	return result, response, err

}

//ClaimDeviceV1 Claim Device - d8a6-1997-4a8a-8c48
/* Claims one of more devices with specified workflow



Documentation Link: https://developer.cisco.com/docs/dna-center/#!claim-device-v1
*/
func (s *DeviceOnboardingPnpService) ClaimDeviceV1(requestDeviceOnboardingPnpClaimDeviceV1 *RequestDeviceOnboardingPnpClaimDeviceV1) (*ResponseDeviceOnboardingPnpClaimDeviceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-device/claim"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceOnboardingPnpClaimDeviceV1).
		SetResult(&ResponseDeviceOnboardingPnpClaimDeviceV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ClaimDeviceV1(requestDeviceOnboardingPnpClaimDeviceV1)
		}

		return nil, response, fmt.Errorf("error with operation ClaimDeviceV1")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpClaimDeviceV1)
	return result, response, err

}

//ImportDevicesInBulkV1 Import Devices in bulk - 21a6-db25-4029-8f55
/* Add devices to PnP in bulk



Documentation Link: https://developer.cisco.com/docs/dna-center/#!import-devices-in-bulk-v1
*/
func (s *DeviceOnboardingPnpService) ImportDevicesInBulkV1(requestDeviceOnboardingPnpImportDevicesInBulkV1 *RequestDeviceOnboardingPnpImportDevicesInBulkV1) (*ResponseDeviceOnboardingPnpImportDevicesInBulkV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-device/import"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceOnboardingPnpImportDevicesInBulkV1).
		SetResult(&ResponseDeviceOnboardingPnpImportDevicesInBulkV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ImportDevicesInBulkV1(requestDeviceOnboardingPnpImportDevicesInBulkV1)
		}

		return nil, response, fmt.Errorf("error with operation ImportDevicesInBulkV1")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpImportDevicesInBulkV1)
	return result, response, err

}

//ResetDeviceV1 Reset Device - 9e85-7b5a-4a0b-bcdb
/* Recovers a device from a Workflow Execution Error state



Documentation Link: https://developer.cisco.com/docs/dna-center/#!reset-device-v1
*/
func (s *DeviceOnboardingPnpService) ResetDeviceV1(requestDeviceOnboardingPnpResetDeviceV1 *RequestDeviceOnboardingPnpResetDeviceV1) (*ResponseDeviceOnboardingPnpResetDeviceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-device/reset"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceOnboardingPnpResetDeviceV1).
		SetResult(&ResponseDeviceOnboardingPnpResetDeviceV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ResetDeviceV1(requestDeviceOnboardingPnpResetDeviceV1)
		}

		return nil, response, fmt.Errorf("error with operation ResetDeviceV1")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpResetDeviceV1)
	return result, response, err

}

//ClaimADeviceToASiteV1 Claim a Device to a Site - 5889-fb84-4939-a13b
/* Claim a device based on Catalyst Center Site-based design process. Some required parameters differ based on device platform:
Default/StackSwitch: imageInfo, configInfo.
AccessPoints: rfProfile.
Sensors: sensorProfile.
CatalystWLC/MobilityExpress/EWC: staticIP, subnetMask, gateway. vlanId and ipInterfaceName are also allowed for Catalyst 9800 WLCs.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!claim-a-device-to-a-site-v1
*/
func (s *DeviceOnboardingPnpService) ClaimADeviceToASiteV1(requestDeviceOnboardingPnpClaimADeviceToASiteV1 *RequestDeviceOnboardingPnpClaimADeviceToASiteV1) (*ResponseDeviceOnboardingPnpClaimADeviceToASiteV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-device/site-claim"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceOnboardingPnpClaimADeviceToASiteV1).
		SetResult(&ResponseDeviceOnboardingPnpClaimADeviceToASiteV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ClaimADeviceToASiteV1(requestDeviceOnboardingPnpClaimADeviceToASiteV1)
		}

		return nil, response, fmt.Errorf("error with operation ClaimADeviceToASiteV1")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpClaimADeviceToASiteV1)
	return result, response, err

}

//PreviewConfigV1 Preview Config - cf94-1823-4d9a-b37e
/* Triggers a preview for site-based Day 0 Configuration



Documentation Link: https://developer.cisco.com/docs/dna-center/#!preview-config-v1
*/
func (s *DeviceOnboardingPnpService) PreviewConfigV1(requestDeviceOnboardingPnpPreviewConfigV1 *RequestDeviceOnboardingPnpPreviewConfigV1) (*ResponseDeviceOnboardingPnpPreviewConfigV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-device/site-config-preview"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceOnboardingPnpPreviewConfigV1).
		SetResult(&ResponseDeviceOnboardingPnpPreviewConfigV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.PreviewConfigV1(requestDeviceOnboardingPnpPreviewConfigV1)
		}

		return nil, response, fmt.Errorf("error with operation PreviewConfigV1")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpPreviewConfigV1)
	return result, response, err

}

//UnClaimDeviceV1 Un-Claim Device - 0b83-6b7b-4b6a-9fd5
/* Un-Claims one of more devices with specified workflow (Deprecated).



Documentation Link: https://developer.cisco.com/docs/dna-center/#!un-claim-device-v1
*/
func (s *DeviceOnboardingPnpService) UnClaimDeviceV1(requestDeviceOnboardingPnpUnClaimDeviceV1 *RequestDeviceOnboardingPnpUnClaimDeviceV1) (*ResponseDeviceOnboardingPnpUnClaimDeviceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-device/unclaim"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceOnboardingPnpUnClaimDeviceV1).
		SetResult(&ResponseDeviceOnboardingPnpUnClaimDeviceV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.UnClaimDeviceV1(requestDeviceOnboardingPnpUnClaimDeviceV1)
		}

		return nil, response, fmt.Errorf("error with operation UnClaimDeviceV1")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpUnClaimDeviceV1)
	return result, response, err

}

//SyncVirtualAccountDevicesV1 Sync Virtual Account Devices - a4b6-c87a-4ffb-9efa
/* Synchronizes the device info from the given smart account & virtual account with the PnP database. The response payload returns a list of synced devices (Deprecated).



Documentation Link: https://developer.cisco.com/docs/dna-center/#!sync-virtual-account-devices-v1
*/
func (s *DeviceOnboardingPnpService) SyncVirtualAccountDevicesV1(requestDeviceOnboardingPnpSyncVirtualAccountDevicesV1 *RequestDeviceOnboardingPnpSyncVirtualAccountDevicesV1) (*ResponseDeviceOnboardingPnpSyncVirtualAccountDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-device/vacct-sync"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceOnboardingPnpSyncVirtualAccountDevicesV1).
		SetResult(&ResponseDeviceOnboardingPnpSyncVirtualAccountDevicesV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.SyncVirtualAccountDevicesV1(requestDeviceOnboardingPnpSyncVirtualAccountDevicesV1)
		}

		return nil, response, fmt.Errorf("error with operation SyncVirtualAccountDevicesV1")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpSyncVirtualAccountDevicesV1)
	return result, response, err

}

//AddVirtualAccountV1 Add Virtual Account - 1e96-2af3-45b8-b59f
/* Registers a Smart Account, Virtual Account and the relevant server profile info with the PnP System & database. The devices present in the registered virtual account are synced with the PnP database as well. The response payload returns the new profile



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-virtual-account-v1
*/
func (s *DeviceOnboardingPnpService) AddVirtualAccountV1(requestDeviceOnboardingPnpAddVirtualAccountV1 *RequestDeviceOnboardingPnpAddVirtualAccountV1) (*ResponseDeviceOnboardingPnpAddVirtualAccountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-settings/savacct"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceOnboardingPnpAddVirtualAccountV1).
		SetResult(&ResponseDeviceOnboardingPnpAddVirtualAccountV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddVirtualAccountV1(requestDeviceOnboardingPnpAddVirtualAccountV1)
		}

		return nil, response, fmt.Errorf("error with operation AddVirtualAccountV1")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpAddVirtualAccountV1)
	return result, response, err

}

//AddAWorkflowV1 Add a Workflow - 848b-5a7b-4f9b-8c12
/* Adds a PnP Workflow along with the relevant tasks in the workflow into the PnP database



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-a-workflow-v1
*/
func (s *DeviceOnboardingPnpService) AddAWorkflowV1(requestDeviceOnboardingPnpAddAWorkflowV1 *RequestDeviceOnboardingPnpAddAWorkflowV1) (*ResponseDeviceOnboardingPnpAddAWorkflowV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-workflow"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceOnboardingPnpAddAWorkflowV1).
		SetResult(&ResponseDeviceOnboardingPnpAddAWorkflowV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddAWorkflowV1(requestDeviceOnboardingPnpAddAWorkflowV1)
		}

		return nil, response, fmt.Errorf("error with operation AddAWorkflowV1")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpAddAWorkflowV1)
	return result, response, err

}

//UpdateDeviceV1 Update Device - 09b0-f9ce-4239-ae10
/* Updates device details specified by device id in PnP database


@param id id path parameter.
*/
func (s *DeviceOnboardingPnpService) UpdateDeviceV1(id string, requestDeviceOnboardingPnpUpdateDeviceV1 *RequestDeviceOnboardingPnpUpdateDeviceV1) (*ResponseDeviceOnboardingPnpUpdateDeviceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-device/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceOnboardingPnpUpdateDeviceV1).
		SetResult(&ResponseDeviceOnboardingPnpUpdateDeviceV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateDeviceV1(id, requestDeviceOnboardingPnpUpdateDeviceV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateDeviceV1")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpUpdateDeviceV1)
	return result, response, err

}

//UpdatePnpGlobalSettingsV1 Update PnP global settings - 8da0-3919-4708-8a5a
/* Updates the user's list of global PnP settings


 */
func (s *DeviceOnboardingPnpService) UpdatePnpGlobalSettingsV1(requestDeviceOnboardingPnpUpdatePnPGlobalSettingsV1 *RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsV1) (*ResponseDeviceOnboardingPnpUpdatePnpGlobalSettingsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-settings"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceOnboardingPnpUpdatePnPGlobalSettingsV1).
		SetResult(&ResponseDeviceOnboardingPnpUpdatePnpGlobalSettingsV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdatePnpGlobalSettingsV1(requestDeviceOnboardingPnpUpdatePnPGlobalSettingsV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdatePnpGlobalSettingsV1")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpUpdatePnpGlobalSettingsV1)
	return result, response, err

}

//UpdatePnpServerProfileV1 Update PnP Server Profile - 6f98-19e8-4178-870c
/* Updates the PnP Server profile in a registered Virtual Account in the PnP database. The response payload returns the updated smart & virtual account info


 */
func (s *DeviceOnboardingPnpService) UpdatePnpServerProfileV1(requestDeviceOnboardingPnpUpdatePnPServerProfileV1 *RequestDeviceOnboardingPnpUpdatePnpServerProfileV1) (*ResponseDeviceOnboardingPnpUpdatePnpServerProfileV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-settings/savacct"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceOnboardingPnpUpdatePnPServerProfileV1).
		SetResult(&ResponseDeviceOnboardingPnpUpdatePnpServerProfileV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdatePnpServerProfileV1(requestDeviceOnboardingPnpUpdatePnPServerProfileV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdatePnpServerProfileV1")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpUpdatePnpServerProfileV1)
	return result, response, err

}

//UpdateWorkflowV1 Update Workflow - 3086-c962-4f49-8b85
/* Updates an existing workflow


@param id id path parameter.
*/
func (s *DeviceOnboardingPnpService) UpdateWorkflowV1(id string, requestDeviceOnboardingPnpUpdateWorkflowV1 *RequestDeviceOnboardingPnpUpdateWorkflowV1) (*ResponseDeviceOnboardingPnpUpdateWorkflowV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-workflow/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceOnboardingPnpUpdateWorkflowV1).
		SetResult(&ResponseDeviceOnboardingPnpUpdateWorkflowV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateWorkflowV1(id, requestDeviceOnboardingPnpUpdateWorkflowV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateWorkflowV1")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpUpdateWorkflowV1)
	return result, response, err

}

//DeleteDeviceByIDFromPnpV1 Delete Device by Id from PnP - cdab-9b47-4899-ae06
/* Deletes specified device from PnP database


@param id id path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-device-by-id-from-pnp-v1
*/
func (s *DeviceOnboardingPnpService) DeleteDeviceByIDFromPnpV1(id string) (*ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/onboarding/pnp-device/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteDeviceByIDFromPnpV1(
				id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteDeviceByIdFromPnpV1")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1)
	return result, response, err

}

//DeregisterVirtualAccountV1 Deregister Virtual Account - 2499-e9ad-42e8-ae5b
/* Deregisters the specified smart account & virtual account info and the associated device information from the PnP System & database. The devices associated with the deregistered virtual account are removed from the PnP database as well. The response payload contains the deregistered smart & virtual account information


@param DeregisterVirtualAccountV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!deregister-virtual-account-v1
*/
func (s *DeviceOnboardingPnpService) DeregisterVirtualAccountV1(DeregisterVirtualAccountV1QueryParams *DeregisterVirtualAccountV1QueryParams) (*ResponseDeviceOnboardingPnpDeregisterVirtualAccountV1, *resty.Response, error) {
	//DeregisterVirtualAccountV1QueryParams *DeregisterVirtualAccountV1QueryParams
	path := "/dna/intent/api/v1/onboarding/pnp-settings/vacct"

	queryString, _ := query.Values(DeregisterVirtualAccountV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDeviceOnboardingPnpDeregisterVirtualAccountV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeregisterVirtualAccountV1(
				DeregisterVirtualAccountV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeregisterVirtualAccountV1")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpDeregisterVirtualAccountV1)
	return result, response, err

}

//DeleteWorkflowByIDV1 Delete Workflow By Id - af8d-7b0e-470b-8ae2
/* Deletes a workflow specified by id


@param id id path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-workflow-by-id-v1
*/
func (s *DeviceOnboardingPnpService) DeleteWorkflowByIDV1(id string) (*ResponseDeviceOnboardingPnpDeleteWorkflowByIDV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/onboarding/pnp-workflow/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceOnboardingPnpDeleteWorkflowByIDV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteWorkflowByIDV1(
				id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteWorkflowByIdV1")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpDeleteWorkflowByIDV1)
	return result, response, err

}

// Alias Function
/*
This method acts as an alias for the method `AddDeviceV1`
*/
func (s *DeviceOnboardingPnpService) AddDevice(requestDeviceOnboardingPnpAddDeviceV1 *RequestDeviceOnboardingPnpAddDeviceV1) (*ResponseDeviceOnboardingPnpAddDeviceV1, *resty.Response, error) {
	return s.AddDeviceV1(requestDeviceOnboardingPnpAddDeviceV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetWorkflowsV1`
*/
func (s *DeviceOnboardingPnpService) GetWorkflows(GetWorkflowsV1QueryParams *GetWorkflowsV1QueryParams) (*ResponseDeviceOnboardingPnpGetWorkflowsV1, *resty.Response, error) {
	return s.GetWorkflowsV1(GetWorkflowsV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetWorkflowByIDV1`
*/
func (s *DeviceOnboardingPnpService) GetWorkflowByID(id string) (*ResponseDeviceOnboardingPnpGetWorkflowByIDV1, *resty.Response, error) {
	return s.GetWorkflowByIDV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `ImportDevicesInBulkV1`
*/
func (s *DeviceOnboardingPnpService) ImportDevicesInBulk(requestDeviceOnboardingPnpImportDevicesInBulkV1 *RequestDeviceOnboardingPnpImportDevicesInBulkV1) (*ResponseDeviceOnboardingPnpImportDevicesInBulkV1, *resty.Response, error) {
	return s.ImportDevicesInBulkV1(requestDeviceOnboardingPnpImportDevicesInBulkV1)
}

// Alias Function
/*
This method acts as an alias for the method `AddVirtualAccountV1`
*/
func (s *DeviceOnboardingPnpService) AddVirtualAccount(requestDeviceOnboardingPnpAddVirtualAccountV1 *RequestDeviceOnboardingPnpAddVirtualAccountV1) (*ResponseDeviceOnboardingPnpAddVirtualAccountV1, *resty.Response, error) {
	return s.AddVirtualAccountV1(requestDeviceOnboardingPnpAddVirtualAccountV1)
}

// Alias Function
/*
This method acts as an alias for the method `UpdatePnpGlobalSettingsV1`
*/
func (s *DeviceOnboardingPnpService) UpdatePnpGlobalSettings(requestDeviceOnboardingPnpUpdatePnPGlobalSettingsV1 *RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsV1) (*ResponseDeviceOnboardingPnpUpdatePnpGlobalSettingsV1, *resty.Response, error) {
	return s.UpdatePnpGlobalSettingsV1(requestDeviceOnboardingPnpUpdatePnPGlobalSettingsV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetVirtualAccountListV1`
*/
func (s *DeviceOnboardingPnpService) GetVirtualAccountList(domain string) (*ResponseDeviceOnboardingPnpGetVirtualAccountListV1, *resty.Response, error) {
	return s.GetVirtualAccountListV1(domain)
}

// Alias Function
/*
This method acts as an alias for the method `PreviewConfigV1`
*/
func (s *DeviceOnboardingPnpService) PreviewConfig(requestDeviceOnboardingPnpPreviewConfigV1 *RequestDeviceOnboardingPnpPreviewConfigV1) (*ResponseDeviceOnboardingPnpPreviewConfigV1, *resty.Response, error) {
	return s.PreviewConfigV1(requestDeviceOnboardingPnpPreviewConfigV1)
}

// Alias Function
/*
This method acts as an alias for the method `ResetDeviceV1`
*/
func (s *DeviceOnboardingPnpService) ResetDevice(requestDeviceOnboardingPnpResetDeviceV1 *RequestDeviceOnboardingPnpResetDeviceV1) (*ResponseDeviceOnboardingPnpResetDeviceV1, *resty.Response, error) {
	return s.ResetDeviceV1(requestDeviceOnboardingPnpResetDeviceV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetSmartAccountListV1`
*/
func (s *DeviceOnboardingPnpService) GetSmartAccountList() (*ResponseDeviceOnboardingPnpGetSmartAccountListV1, *resty.Response, error) {
	return s.GetSmartAccountListV1()
}

// Alias Function
/*
This method acts as an alias for the method `UpdatePnpServerProfileV1`
*/
func (s *DeviceOnboardingPnpService) UpdatePnpServerProfile(requestDeviceOnboardingPnpUpdatePnPServerProfileV1 *RequestDeviceOnboardingPnpUpdatePnpServerProfileV1) (*ResponseDeviceOnboardingPnpUpdatePnpServerProfileV1, *resty.Response, error) {
	return s.UpdatePnpServerProfileV1(requestDeviceOnboardingPnpUpdatePnPServerProfileV1)
}

// Alias Function
/*
This method acts as an alias for the method `UpdateDeviceV1`
*/
func (s *DeviceOnboardingPnpService) UpdateDevice(id string, requestDeviceOnboardingPnpUpdateDeviceV1 *RequestDeviceOnboardingPnpUpdateDeviceV1) (*ResponseDeviceOnboardingPnpUpdateDeviceV1, *resty.Response, error) {
	return s.UpdateDeviceV1(id, requestDeviceOnboardingPnpUpdateDeviceV1)
}

// Alias Function
/*
This method acts as an alias for the method `AuthorizeDeviceV1`
*/
func (s *DeviceOnboardingPnpService) AuthorizeDevice(requestDeviceOnboardingPnpAuthorizeDeviceV1 *RequestDeviceOnboardingPnpAuthorizeDeviceV1) (*ResponseDeviceOnboardingPnpAuthorizeDeviceV1, *resty.Response, error) {
	return s.AuthorizeDeviceV1(requestDeviceOnboardingPnpAuthorizeDeviceV1)
}

// Alias Function
/*
This method acts as an alias for the method `DeleteWorkflowByIDV1`
*/
func (s *DeviceOnboardingPnpService) DeleteWorkflowByID(id string) (*ResponseDeviceOnboardingPnpDeleteWorkflowByIDV1, *resty.Response, error) {
	return s.DeleteWorkflowByIDV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `GetDeviceHistoryV1`
*/
func (s *DeviceOnboardingPnpService) GetDeviceHistory(GetDeviceHistoryV1QueryParams *GetDeviceHistoryV1QueryParams) (*ResponseDeviceOnboardingPnpGetDeviceHistoryV1, *resty.Response, error) {
	return s.GetDeviceHistoryV1(GetDeviceHistoryV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `AddAWorkflowV1`
*/
func (s *DeviceOnboardingPnpService) AddAWorkflow(requestDeviceOnboardingPnpAddAWorkflowV1 *RequestDeviceOnboardingPnpAddAWorkflowV1) (*ResponseDeviceOnboardingPnpAddAWorkflowV1, *resty.Response, error) {
	return s.AddAWorkflowV1(requestDeviceOnboardingPnpAddAWorkflowV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetDeviceCountV1`
*/
func (s *DeviceOnboardingPnpService) GetDeviceCount(GetDeviceCountV1QueryParams *GetDeviceCountV1QueryParams) (*ResponseDeviceOnboardingPnpGetDeviceCountV1, *resty.Response, error) {
	return s.GetDeviceCountV1(GetDeviceCountV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `ClaimDeviceV1`
*/
func (s *DeviceOnboardingPnpService) ClaimDevice(requestDeviceOnboardingPnpClaimDeviceV1 *RequestDeviceOnboardingPnpClaimDeviceV1) (*ResponseDeviceOnboardingPnpClaimDeviceV1, *resty.Response, error) {
	return s.ClaimDeviceV1(requestDeviceOnboardingPnpClaimDeviceV1)
}

// Alias Function
/*
This method acts as an alias for the method `UpdateWorkflowV1`
*/
func (s *DeviceOnboardingPnpService) UpdateWorkflow(id string, requestDeviceOnboardingPnpUpdateWorkflowV1 *RequestDeviceOnboardingPnpUpdateWorkflowV1) (*ResponseDeviceOnboardingPnpUpdateWorkflowV1, *resty.Response, error) {
	return s.UpdateWorkflowV1(id, requestDeviceOnboardingPnpUpdateWorkflowV1)
}

// Alias Function
/*
This method acts as an alias for the method `SyncVirtualAccountDevicesV1`
*/
func (s *DeviceOnboardingPnpService) SyncVirtualAccountDevices(requestDeviceOnboardingPnpSyncVirtualAccountDevicesV1 *RequestDeviceOnboardingPnpSyncVirtualAccountDevicesV1) (*ResponseDeviceOnboardingPnpSyncVirtualAccountDevicesV1, *resty.Response, error) {
	return s.SyncVirtualAccountDevicesV1(requestDeviceOnboardingPnpSyncVirtualAccountDevicesV1)
}

// Alias Function
/*
This method acts as an alias for the method `UnClaimDeviceV1`
*/
func (s *DeviceOnboardingPnpService) UnClaimDevice(requestDeviceOnboardingPnpUnClaimDeviceV1 *RequestDeviceOnboardingPnpUnClaimDeviceV1) (*ResponseDeviceOnboardingPnpUnClaimDeviceV1, *resty.Response, error) {
	return s.UnClaimDeviceV1(requestDeviceOnboardingPnpUnClaimDeviceV1)
}

// Alias Function
/*
This method acts as an alias for the method `DeleteDeviceByIDFromPnpV1`
*/
func (s *DeviceOnboardingPnpService) DeleteDeviceByIDFromPnp(id string) (*ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpV1, *resty.Response, error) {
	return s.DeleteDeviceByIDFromPnpV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `GetSyncResultForVirtualAccountV1`
*/
func (s *DeviceOnboardingPnpService) GetSyncResultForVirtualAccount(domain string, name string) (*ResponseDeviceOnboardingPnpGetSyncResultForVirtualAccountV1, *resty.Response, error) {
	return s.GetSyncResultForVirtualAccountV1(domain, name)
}

// Alias Function
/*
This method acts as an alias for the method `DeregisterVirtualAccountV1`
*/
func (s *DeviceOnboardingPnpService) DeregisterVirtualAccount(DeregisterVirtualAccountV1QueryParams *DeregisterVirtualAccountV1QueryParams) (*ResponseDeviceOnboardingPnpDeregisterVirtualAccountV1, *resty.Response, error) {
	return s.DeregisterVirtualAccountV1(DeregisterVirtualAccountV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetDeviceListSiteManagementV1`
*/
func (s *DeviceOnboardingPnpService) GetDeviceListSiteManagement(GetDeviceListSiteManagementV1QueryParams *GetDeviceListSiteManagementV1QueryParams) (*ResponseDeviceOnboardingPnpGetDeviceListSiteManagementV1, *resty.Response, error) {
	return s.GetDeviceListSiteManagementV1(GetDeviceListSiteManagementV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetPnpGlobalSettingsV1`
*/
func (s *DeviceOnboardingPnpService) GetPnpGlobalSettings() (*ResponseDeviceOnboardingPnpGetPnpGlobalSettingsV1, *resty.Response, error) {
	return s.GetPnpGlobalSettingsV1()
}

// Alias Function
/*
This method acts as an alias for the method `GetDeviceByIDV1`
*/
func (s *DeviceOnboardingPnpService) GetDeviceByID(id string) (*ResponseDeviceOnboardingPnpGetDeviceByIDV1, *resty.Response, error) {
	return s.GetDeviceByIDV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `ClaimADeviceToASiteV1`
*/
func (s *DeviceOnboardingPnpService) ClaimADeviceToASite(requestDeviceOnboardingPnpClaimADeviceToASiteV1 *RequestDeviceOnboardingPnpClaimADeviceToASiteV1) (*ResponseDeviceOnboardingPnpClaimADeviceToASiteV1, *resty.Response, error) {
	return s.ClaimADeviceToASiteV1(requestDeviceOnboardingPnpClaimADeviceToASiteV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetWorkflowCountV1`
*/
func (s *DeviceOnboardingPnpService) GetWorkflowCount(GetWorkflowCountV1QueryParams *GetWorkflowCountV1QueryParams) (*ResponseDeviceOnboardingPnpGetWorkflowCountV1, *resty.Response, error) {
	return s.GetWorkflowCountV1(GetWorkflowCountV1QueryParams)
}
