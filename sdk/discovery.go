package catalyst

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type DiscoveryService service

type GetDiscoveryJobsByIPV1QueryParams struct {
	Offset    int    `url:"offset,omitempty"`    //offset
	Limit     int    `url:"limit,omitempty"`     //limit
	IPAddress string `url:"ipAddress,omitempty"` //ipAddress
	Name      string `url:"name,omitempty"`      //name
}
type GetListOfDiscoveriesByDiscoveryIDV1QueryParams struct {
	Offset    int    `url:"offset,omitempty"`    //Starting index for the records
	Limit     int    `url:"limit,omitempty"`     //Number of records to fetch from the starting index
	IPAddress string `url:"ipAddress,omitempty"` //Filter records based on IP address
}
type GetDiscoveredNetworkDevicesByDiscoveryIDV1QueryParams struct {
	TaskID string `url:"taskId,omitempty"` //taskId
}
type GetDevicesDiscoveredByIDV1QueryParams struct {
	TaskID string `url:"taskId,omitempty"` //taskId
}
type GetDiscoveredDevicesByRangeV1QueryParams struct {
	TaskID string `url:"taskId,omitempty"` //taskId
}
type GetNetworkDevicesFromDiscoveryV1QueryParams struct {
	TaskID        string   `url:"taskId,omitempty"`        //taskId
	SortBy        string   `url:"sortBy,omitempty"`        //Sort by field. Available values are pingStatus, cliStatus,snmpStatus, httpStatus and netconfStatus
	SortOrder     string   `url:"sortOrder,omitempty"`     //Order of sorting based on sortBy. Available values are 'asc' and 'des'
	IPAddress     []string `url:"ipAddress,omitempty"`     //IP Address of the device
	PingStatus    []string `url:"pingStatus,omitempty"`    //Ping status for the IP during the job run. Available values are 'SUCCESS', 'FAILURE', 'NOT-PROVIDED' and 'NOT-VALIDATED'
	SNMPStatus    []string `url:"snmpStatus,omitempty"`    //SNMP status for the IP during the job run. Available values are 'SUCCESS', 'FAILURE', 'NOT-PROVIDED' and 'NOT-VALIDATED'
	Clistatus     []string `url:"cliStatus,omitempty"`     //CLI status for the IP during the job run. Available values are 'SUCCESS', 'FAILURE', 'NOT-PROVIDED' and 'NOT-VALIDATED'
	NetconfStatus []string `url:"netconfStatus,omitempty"` //NETCONF status for the IP during the job run. Available values are 'SUCCESS', 'FAILURE', 'NOT-PROVIDED' and 'NOT-VALIDATED'
	HTTPStatus    []string `url:"httpStatus,omitempty"`    //HTTP staus for the IP during the job run. Available values are 'SUCCESS', 'FAILURE', 'NOT-PROVIDED' and 'NOT-VALIDATED'
}
type GetGlobalCredentialsV1QueryParams struct {
	CredentialSubType string `url:"credentialSubType,omitempty"` //Credential type as CLI / SNMPV2_READ_COMMUNITY / SNMPV2_WRITE_COMMUNITY / SNMPV3 / HTTP_WRITE / HTTP_READ / NETCONF
	SortBy            string `url:"sortBy,omitempty"`            //Field to sort the results by. Sorts by 'instanceId' if no value is provided
	Order             string `url:"order,omitempty"`             //Order of sorting. 'asc' or 'des'
}

type ResponseDiscoveryDeleteAllDiscoveryV1 struct {
	Response *ResponseDiscoveryDeleteAllDiscoveryV1Response `json:"response,omitempty"` //
	Version  string                                         `json:"version,omitempty"`  //
}
type ResponseDiscoveryDeleteAllDiscoveryV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDV1 struct {
	Response *ResponseDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDV1Response `json:"response,omitempty"` //
	Version  string                                                              `json:"version,omitempty"`  //
}
type ResponseDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryStartDiscoveryV1 struct {
	Response *ResponseDiscoveryStartDiscoveryV1Response `json:"response,omitempty"` //
	Version  string                                     `json:"version,omitempty"`  //
}
type ResponseDiscoveryStartDiscoveryV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryGetCountOfAllDiscoveryJobsV1 struct {
	Response *int   `json:"response,omitempty"` // The count of all available discovery jobs
	Version  string `json:"version,omitempty"`  //
}
type ResponseDiscoveryGetDiscoveryJobsByIPV1 struct {
	Response *[]ResponseDiscoveryGetDiscoveryJobsByIPV1Response `json:"response,omitempty"` //
	Version  string                                             `json:"version,omitempty"`  //
}
type ResponseDiscoveryGetDiscoveryJobsByIPV1Response struct {
	AttributeInfo               *ResponseDiscoveryGetDiscoveryJobsByIPV1ResponseAttributeInfo `json:"attributeInfo,omitempty"`               //
	Clistatus                   string                                                        `json:"cliStatus,omitempty"`                   //
	DiscoveryStatus             string                                                        `json:"discoveryStatus,omitempty"`             //
	EndTime                     string                                                        `json:"endTime,omitempty"`                     //
	HTTPStatus                  string                                                        `json:"httpStatus,omitempty"`                  //
	ID                          string                                                        `json:"id,omitempty"`                          //
	InventoryCollectionStatus   string                                                        `json:"inventoryCollectionStatus,omitempty"`   //
	InventoryReachabilityStatus string                                                        `json:"inventoryReachabilityStatus,omitempty"` //
	IPAddress                   string                                                        `json:"ipAddress,omitempty"`                   //
	JobStatus                   string                                                        `json:"jobStatus,omitempty"`                   //
	Name                        string                                                        `json:"name,omitempty"`                        //
	NetconfStatus               string                                                        `json:"netconfStatus,omitempty"`               //
	PingStatus                  string                                                        `json:"pingStatus,omitempty"`                  //
	SNMPStatus                  string                                                        `json:"snmpStatus,omitempty"`                  //
	StartTime                   string                                                        `json:"startTime,omitempty"`                   //
	TaskID                      string                                                        `json:"taskId,omitempty"`                      //
}
type ResponseDiscoveryGetDiscoveryJobsByIPV1ResponseAttributeInfo interface{}
type ResponseDiscoveryDeleteDiscoveryByIDV1 struct {
	Response *ResponseDiscoveryDeleteDiscoveryByIDV1Response `json:"response,omitempty"` //
	Version  string                                          `json:"version,omitempty"`  //
}
type ResponseDiscoveryDeleteDiscoveryByIDV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryGetDiscoveryByIDV1 struct {
	Response *ResponseDiscoveryGetDiscoveryByIDV1Response `json:"response,omitempty"` //
	Version  string                                       `json:"version,omitempty"`  //
}
type ResponseDiscoveryGetDiscoveryByIDV1Response struct {
	AttributeInfo          *ResponseDiscoveryGetDiscoveryByIDV1ResponseAttributeInfo       `json:"attributeInfo,omitempty"`          // Deprecated
	CdpLevel               *int                                                            `json:"cdpLevel,omitempty"`               // CDP level to which neighbor devices to be discovered
	DeviceIDs              string                                                          `json:"deviceIds,omitempty"`              // Ids of the devices discovered in a discovery
	DiscoveryCondition     string                                                          `json:"discoveryCondition,omitempty"`     // To indicate the discovery status. Available options: Complete or In Progress
	DiscoveryStatus        string                                                          `json:"discoveryStatus,omitempty"`        // Status of the discovery. Available options are: Active, Inactive, Edit
	DiscoveryType          string                                                          `json:"discoveryType,omitempty"`          // Type of the discovery. Available types are: 'Single', 'Range', 'CDP', 'LLDP', 'CIDR'
	EnablePasswordList     string                                                          `json:"enablePasswordList,omitempty"`     // Enable Password of the devices to be discovered
	GlobalCredentialIDList []string                                                        `json:"globalCredentialIdList,omitempty"` // List of global credential ids to be used
	HTTPReadCredential     *ResponseDiscoveryGetDiscoveryByIDV1ResponseHTTPReadCredential  `json:"httpReadCredential,omitempty"`     //
	HTTPWriteCredential    *ResponseDiscoveryGetDiscoveryByIDV1ResponseHTTPWriteCredential `json:"httpWriteCredential,omitempty"`    //
	ID                     string                                                          `json:"id,omitempty"`                     // Unique Discovery Id
	IPAddressList          string                                                          `json:"ipAddressList,omitempty"`          // List of IP address of the devices to be discovered
	IPFilterList           string                                                          `json:"ipFilterList,omitempty"`           // IP addresses of the devices to be filtered
	IsAutoCdp              *bool                                                           `json:"isAutoCdp,omitempty"`              // Flag to mention if CDP discovery or not
	LldpLevel              *int                                                            `json:"lldpLevel,omitempty"`              // LLDP level to which neighbor devices to be discovered
	Name                   string                                                          `json:"name,omitempty"`                   // Name for the discovery
	NetconfPort            string                                                          `json:"netconfPort,omitempty"`            // Netconf port on the device. Netconf will need valid sshv2 credentials for it to work
	NumDevices             *int                                                            `json:"numDevices,omitempty"`             // Number of devices discovered in the discovery
	ParentDiscoveryID      string                                                          `json:"parentDiscoveryId,omitempty"`      // Parent Discovery Id from which the discovery was initiated
	PasswordList           string                                                          `json:"passwordList,omitempty"`           // Password of the devices to be discovered
	PreferredMgmtIPMethod  string                                                          `json:"preferredMgmtIPMethod,omitempty"`  // Preferred management IP method. Available options are '' and 'UseLoopBack'
	ProtocolOrder          string                                                          `json:"protocolOrder,omitempty"`          // Order of protocol (ssh/telnet) in which device connection will be tried. Ex: 'telnet': only telnet; 'ssh,telnet': ssh with higher order than telnet
	RetryCount             *int                                                            `json:"retryCount,omitempty"`             // Number of times to try establishing connection to device
	SNMPAuthPassphrase     string                                                          `json:"snmpAuthPassphrase,omitempty"`     // Auth passphrase for SNMP
	SNMPAuthProtocol       string                                                          `json:"snmpAuthProtocol,omitempty"`       // SNMP auth protocol. SHA' or 'MD5'
	SNMPMode               string                                                          `json:"snmpMode,omitempty"`               // Mode of SNMP. 'AUTHPRIV' or 'AUTHNOPRIV' or 'NOAUTHNOPRIV'
	SNMPPrivPassphrase     string                                                          `json:"snmpPrivPassphrase,omitempty"`     // Passphrase for SNMP privacy
	SNMPPrivProtocol       string                                                          `json:"snmpPrivProtocol,omitempty"`       // SNMP privacy protocol. 'AES128'
	SNMPRoCommunity        string                                                          `json:"snmpRoCommunity,omitempty"`        // SNMP RO community of the devices to be discovered
	SNMPRoCommunityDesc    string                                                          `json:"snmpRoCommunityDesc,omitempty"`    // Description for SNMP RO community
	SNMPRwCommunity        string                                                          `json:"snmpRwCommunity,omitempty"`        // SNMP RW community of the devices to be discovered
	SNMPRwCommunityDesc    string                                                          `json:"snmpRwCommunityDesc,omitempty"`    // Description for SNMP RW community
	SNMPUserName           string                                                          `json:"snmpUserName,omitempty"`           // SNMP username of the device
	TimeOut                *int                                                            `json:"timeOut,omitempty"`                // Time to wait for device response.
	UpdateMgmtIP           *bool                                                           `json:"updateMgmtIp,omitempty"`           // Updates Maganement IP if multiple IPs are available for a device. If set to true, when a device is rediscovered with a different IP, the management IP is updated. Default value is false
	UserNameList           string                                                          `json:"userNameList,omitempty"`           // Username of the devices to be discovered
}
type ResponseDiscoveryGetDiscoveryByIDV1ResponseAttributeInfo interface{}
type ResponseDiscoveryGetDiscoveryByIDV1ResponseHTTPReadCredential struct {
	Comments         string `json:"comments,omitempty"`         // Comments to identify the credential
	CredentialType   string `json:"credentialType,omitempty"`   // Credential type to identify the application that uses the credential
	Description      string `json:"description,omitempty"`      // Description of the credential
	ID               string `json:"id,omitempty"`               // Credential Id
	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Credential Tenant Id
	InstanceUUID     string `json:"instanceUuid,omitempty"`     // Credential Id
	Password         string `json:"password,omitempty"`         // HTTP(S) password
	Port             *int   `json:"port,omitempty"`             // HTTP(S) port
	Secure           *bool  `json:"secure,omitempty"`           // Flag for HTTPS
	Username         string `json:"username,omitempty"`         // HTTP(S) username
}
type ResponseDiscoveryGetDiscoveryByIDV1ResponseHTTPWriteCredential struct {
	Comments         string `json:"comments,omitempty"`         // Comments to identify the credential
	CredentialType   string `json:"credentialType,omitempty"`   // Credential type to identify the application that uses the credential
	Description      string `json:"description,omitempty"`      // Description of the credential
	ID               string `json:"id,omitempty"`               // Credential Id
	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Credential Tenant Id
	InstanceUUID     string `json:"instanceUuid,omitempty"`     // Credential Id
	Password         string `json:"password,omitempty"`         // HTTP(S) password
	Port             *int   `json:"port,omitempty"`             // HTTP(S) port
	Secure           *bool  `json:"secure,omitempty"`           // Flag for HTTPS
	Username         string `json:"username,omitempty"`         // HTTP(S) username
}
type ResponseDiscoveryGetListOfDiscoveriesByDiscoveryIDV1 struct {
	Response *[]ResponseDiscoveryGetListOfDiscoveriesByDiscoveryIDV1Response `json:"response,omitempty"` //
	Version  string                                                          `json:"version,omitempty"`  //
}
type ResponseDiscoveryGetListOfDiscoveriesByDiscoveryIDV1Response struct {
	AttributeInfo               *ResponseDiscoveryGetListOfDiscoveriesByDiscoveryIDV1ResponseAttributeInfo `json:"attributeInfo,omitempty"`               // Deprecated
	Clistatus                   string                                                                     `json:"cliStatus,omitempty"`                   // CLI status for the IP during the job run. Available values are 'SUCCESS', 'FAILURE', 'NOT-PROVIDED' and 'NOT-VALIDATED'
	DiscoveryStatus             string                                                                     `json:"discoveryStatus,omitempty"`             // Status of the discovery. Available options are: MANAGED_DEVICES, UNMANAGED_DEVICES, DISCARDED_DEVICES
	EndTime                     string                                                                     `json:"endTime,omitempty"`                     // End time for the discovery job
	HTTPStatus                  string                                                                     `json:"httpStatus,omitempty"`                  // HTTP status for the IP during the job run. Available values are 'SUCCESS', 'FAILURE', 'NOT-PROVIDED' and 'NOT-VALIDATED'
	ID                          string                                                                     `json:"id,omitempty"`                          // Discovery Id
	InventoryCollectionStatus   string                                                                     `json:"inventoryCollectionStatus,omitempty"`   // Last known inventory collection status of the device. Available values are 'MANAGED', 'ABORTED', 'FAILED', 'PARTIAL COLLECTION FAILURE' and 'NOT-AVAILABLE'
	InventoryReachabilityStatus string                                                                     `json:"inventoryReachabilityStatus,omitempty"` // Last known reachability status of the device. Available values are : 'Reachable', 'Unreachable', 'PingReachable' and 'NOT-AVAILABLE'
	IPAddress                   string                                                                     `json:"ipAddress,omitempty"`                   // IP Address of the device
	JobStatus                   string                                                                     `json:"jobStatus,omitempty"`                   // Status of the job
	Name                        string                                                                     `json:"name,omitempty"`                        // Discovery name
	NetconfStatus               string                                                                     `json:"netconfStatus,omitempty"`               // NETCONF status for the IP during the job run. Available values are 'SUCCESS', 'FAILURE', 'NOT-PROVIDED' and 'NOT-VALIDATED'
	PingStatus                  string                                                                     `json:"pingStatus,omitempty"`                  // Ping status for the IP during the job run. Available values are 'SUCCESS', 'FAILURE', 'NOT-PROVIDED' and 'NOT-VALIDATED
	SNMPStatus                  string                                                                     `json:"snmpStatus,omitempty"`                  // SNMP status for the IP during the job run. Available values are 'SUCCESS', 'FAILURE', 'NOT-PROVIDED' and 'NOT-VALIDATED'
	StartTime                   string                                                                     `json:"startTime,omitempty"`                   // Discovery job start time
	TaskID                      string                                                                     `json:"taskId,omitempty"`                      // Discovery job task id
}
type ResponseDiscoveryGetListOfDiscoveriesByDiscoveryIDV1ResponseAttributeInfo interface{}
type ResponseDiscoveryGetDiscoveredNetworkDevicesByDiscoveryIDV1 struct {
	Response *[]ResponseDiscoveryGetDiscoveredNetworkDevicesByDiscoveryIDV1Response `json:"response,omitempty"` //
	Version  string                                                                 `json:"version,omitempty"`  //
}
type ResponseDiscoveryGetDiscoveredNetworkDevicesByDiscoveryIDV1Response struct {
	AnchorWlcForAp              string `json:"anchorWlcForAp,omitempty"`              // Connected WLC device for AP
	AuthModelID                 string `json:"authModelId,omitempty"`                 // Authentication model Id on device
	AvgUpdateFrequency          *int   `json:"avgUpdateFrequency,omitempty"`          // Frequency in which interface info gets updated
	BootDateTime                string `json:"bootDateTime,omitempty"`                // Device boot time
	Clistatus                   string `json:"cliStatus,omitempty"`                   // CLI status at the time of discovery
	DuplicateDeviceID           string `json:"duplicateDeviceId,omitempty"`           // Identifier of the duplicate ip of the same device discovered
	ErrorCode                   string `json:"errorCode,omitempty"`                   // Error code when inventory collection fails
	ErrorDescription            string `json:"errorDescription,omitempty"`            // Error description when inventory collection fails
	Family                      string `json:"family,omitempty"`                      // Family of device as switch, router, wireless lan controller, accesspoints
	Hostname                    string `json:"hostname,omitempty"`                    // Device name
	HTTPStatus                  string `json:"httpStatus,omitempty"`                  // HTTP(S) status at the time of discovery
	ID                          string `json:"id,omitempty"`                          //
	ImageName                   string `json:"imageName,omitempty"`                   // Image details on the device
	IngressQueueConfig          string `json:"ingressQueueConfig,omitempty"`          // Ingress queue config on device
	InterfaceCount              string `json:"interfaceCount,omitempty"`              // Number of interfaces on the device
	InventoryCollectionStatus   string `json:"inventoryCollectionStatus,omitempty"`   // Last known collection status of the device. Available values are : 'Deleting Device', 'Partial Collection Failure', 'Yet to Sync', 'Could Not Synchronize', 'Not Manageable', 'Managed', 'Incomplete', 'Unreachable', 'In Progress', 'Maintenance', 'Sync Disabled', 'Quarantined', 'Unassociated', 'Unknown'
	InventoryReachabilityStatus string `json:"inventoryReachabilityStatus,omitempty"` // Last known reachability status of the device. Available values are : 'Reachable', 'Unreachable', 'PingReachable' and 'NOT-AVAILABLE’
	LastUpdated                 string `json:"lastUpdated,omitempty"`                 // Time when the network device info last got updated
	LineCardCount               string `json:"lineCardCount,omitempty"`               // Number of linecards on the device
	LineCardID                  string `json:"lineCardId,omitempty"`                  // IDs of linecards of the device
	Location                    string `json:"location,omitempty"`                    // Location ID that is associated with the device
	LocationName                string `json:"locationName,omitempty"`                // Name of the associated location
	MacAddress                  string `json:"macAddress,omitempty"`                  // MAC address of device
	ManagementIPAddress         string `json:"managementIpAddress,omitempty"`         // IP address of the device
	MemorySize                  string `json:"memorySize,omitempty"`                  // Processor memory size
	NetconfStatus               string `json:"netconfStatus,omitempty"`               // NETCONF status at the time of discovery. Available values are 'SUCCESS', 'FAILURE', 'NOT-PROVIDED' and 'NOT-VALIDATED'
	NumUpdates                  *int   `json:"numUpdates,omitempty"`                  // Number of time network-device info got updated
	PingStatus                  string `json:"pingStatus,omitempty"`                  // Ping status at the time of discovery. Available values are 'SUCCESS', 'FAILURE', 'NOT-PROVIDED' and 'NOT-VALIDATED'
	PlatformID                  string `json:"platformId,omitempty"`                  // Platform ID of device
	PortRange                   string `json:"portRange,omitempty"`                   // Range of ports on device
	QosStatus                   string `json:"qosStatus,omitempty"`                   // Qos status on device
	ReachabilityFailureReason   string `json:"reachabilityFailureReason,omitempty"`   // Failure reason for unreachable devices
	ReachabilityStatus          string `json:"reachabilityStatus,omitempty"`          // Reachability status of a device as Success/Failure/Discarded
	Role                        string `json:"role,omitempty"`                        // Role of device as access, distribution, border router, core
	RoleSource                  string `json:"roleSource,omitempty"`                  // Role source as manual / auto
	SerialNumber                string `json:"serialNumber,omitempty"`                // Serial number of device
	SNMPContact                 string `json:"snmpContact,omitempty"`                 // SNMP contact on device
	SNMPLocation                string `json:"snmpLocation,omitempty"`                // SNMP location on device
	SNMPStatus                  string `json:"snmpStatus,omitempty"`                  // SNMP status at the time of discovery
	SoftwareVersion             string `json:"softwareVersion,omitempty"`             // Software version on the device
	Tag                         string `json:"tag,omitempty"`                         // Tag ID that is associated with the device
	TagCount                    *int   `json:"tagCount,omitempty"`                    // Number of tags associated with the device
	Type                        string `json:"type,omitempty"`                        // Type of device as switch, router, wireless lan controller, accesspoints
	UpTime                      string `json:"upTime,omitempty"`                      // Time that shows for how long the device has been up
	Vendor                      string `json:"vendor,omitempty"`                      // Vendor information of the device
	WlcApDeviceStatus           string `json:"wlcApDeviceStatus,omitempty"`           // Collection status of AP devices
}
type ResponseDiscoveryGetDevicesDiscoveredByIDV1 struct {
	Response *int   `json:"response,omitempty"` // The count of network devices discovered in the given discovery
	Version  string `json:"version,omitempty"`  //
}
type ResponseDiscoveryGetDiscoveredDevicesByRangeV1 struct {
	Response *[]ResponseDiscoveryGetDiscoveredDevicesByRangeV1Response `json:"response,omitempty"` //
	Version  string                                                    `json:"version,omitempty"`  //
}
type ResponseDiscoveryGetDiscoveredDevicesByRangeV1Response struct {
	AnchorWlcForAp              string `json:"anchorWlcForAp,omitempty"`              // Connected WLC device for AP
	AuthModelID                 string `json:"authModelId,omitempty"`                 // Authentication model Id on device
	AvgUpdateFrequency          *int   `json:"avgUpdateFrequency,omitempty"`          // Frequency in which interface info gets updated
	BootDateTime                string `json:"bootDateTime,omitempty"`                // Device boot time
	Clistatus                   string `json:"cliStatus,omitempty"`                   // CLI status at the time of discovery
	DuplicateDeviceID           string `json:"duplicateDeviceId,omitempty"`           // Identifier of the duplicate ip of the same device discovered
	ErrorCode                   string `json:"errorCode,omitempty"`                   // Error code when inventory collection fails
	ErrorDescription            string `json:"errorDescription,omitempty"`            // Error description when inventory collection fails
	Family                      string `json:"family,omitempty"`                      // Family of device as switch, router, wireless lan controller, accesspoints
	Hostname                    string `json:"hostname,omitempty"`                    // Device name
	HTTPStatus                  string `json:"httpStatus,omitempty"`                  // HTTP(S) status at the time of discovery
	ID                          string `json:"id,omitempty"`                          // Unique identifier of network device
	ImageName                   string `json:"imageName,omitempty"`                   // Image details on the device
	IngressQueueConfig          string `json:"ingressQueueConfig,omitempty"`          // Ingress queue config on device
	InterfaceCount              string `json:"interfaceCount,omitempty"`              // Number of interfaces on the device
	InventoryCollectionStatus   string `json:"inventoryCollectionStatus,omitempty"`   // Last known collection status of the device. Available values are : 'Deleting Device', 'Partial Collection Failure', 'Yet to Sync', 'Could Not Synchronize', 'Not Manageable', 'Managed', 'Incomplete', 'Unreachable', 'In Progress', 'Maintenance', 'Sync Disabled', 'Quarantined', 'Unassociated', 'Unknown'
	InventoryReachabilityStatus string `json:"inventoryReachabilityStatus,omitempty"` // Last known reachability status of the device. Available values are : 'Reachable', 'Unreachable', 'PingReachable' and 'NOT-AVAILABLE’
	LastUpdated                 string `json:"lastUpdated,omitempty"`                 // Time when the network device info last got updated
	LineCardCount               string `json:"lineCardCount,omitempty"`               // Number of linecards on the device
	LineCardID                  string `json:"lineCardId,omitempty"`                  // IDs of linecards of the device
	Location                    string `json:"location,omitempty"`                    // Location ID that is associated with the device
	LocationName                string `json:"locationName,omitempty"`                // Name of the associated location
	MacAddress                  string `json:"macAddress,omitempty"`                  // MAC address of device
	ManagementIPAddress         string `json:"managementIpAddress,omitempty"`         // IP address of the device
	MemorySize                  string `json:"memorySize,omitempty"`                  // Processor memory size
	NetconfStatus               string `json:"netconfStatus,omitempty"`               // NETCONF status at the time of discovery. Available values are 'SUCCESS', 'FAILURE', 'NOT-PROVIDED' and 'NOT-VALIDATED'
	NumUpdates                  *int   `json:"numUpdates,omitempty"`                  // Number of time network-device info got updated
	PingStatus                  string `json:"pingStatus,omitempty"`                  // Ping status at the time of discovery. Available values are 'SUCCESS', 'FAILURE', 'NOT-PROVIDED' and 'NOT-VALIDATED'
	PlatformID                  string `json:"platformId,omitempty"`                  // Platform ID of device
	PortRange                   string `json:"portRange,omitempty"`                   // Range of ports on device
	QosStatus                   string `json:"qosStatus,omitempty"`                   // Qos status on device
	ReachabilityFailureReason   string `json:"reachabilityFailureReason,omitempty"`   // Failure reason for unreachable devices
	ReachabilityStatus          string `json:"reachabilityStatus,omitempty"`          // Reachability status of a device as Success/Failure/Discarded
	Role                        string `json:"role,omitempty"`                        // Role of device as access, distribution, border router, core
	RoleSource                  string `json:"roleSource,omitempty"`                  // Role source as manual / auto
	SerialNumber                string `json:"serialNumber,omitempty"`                // Serial number of device
	SNMPContact                 string `json:"snmpContact,omitempty"`                 // SNMP contact on device
	SNMPLocation                string `json:"snmpLocation,omitempty"`                // SNMP location on device
	SNMPStatus                  string `json:"snmpStatus,omitempty"`                  // SNMP status at the time of discovery
	SoftwareVersion             string `json:"softwareVersion,omitempty"`             // Software version on the device
	Tag                         string `json:"tag,omitempty"`                         // Tag ID that is associated with the device
	TagCount                    *int   `json:"tagCount,omitempty"`                    // Number of tags associated with the device
	Type                        string `json:"type,omitempty"`                        // Type of device as switch, router, wireless lan controller, accesspoints
	UpTime                      string `json:"upTime,omitempty"`                      // Time that shows for how long the device has been up
	Vendor                      string `json:"vendor,omitempty"`                      // Vendor information of the device
	WlcApDeviceStatus           string `json:"wlcApDeviceStatus,omitempty"`           // Collection status of AP devices
}
type ResponseDiscoveryGetNetworkDevicesFromDiscoveryV1 struct {
	Response *int   `json:"response,omitempty"` // The number of network devices from the discovery job based on the given filters
	Version  string `json:"version,omitempty"`  //
}
type ResponseDiscoveryDeleteDiscoveryBySpecifiedRangeV1 struct {
	Response *ResponseDiscoveryDeleteDiscoveryBySpecifiedRangeV1Response `json:"response,omitempty"` //
	Version  string                                                      `json:"version,omitempty"`  //
}
type ResponseDiscoveryDeleteDiscoveryBySpecifiedRangeV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryGetDiscoveriesByRangeV1 struct {
	Response *[]ResponseDiscoveryGetDiscoveriesByRangeV1Response `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  //
}
type ResponseDiscoveryGetDiscoveriesByRangeV1Response struct {
	AttributeInfo          *ResponseDiscoveryGetDiscoveriesByRangeV1ResponseAttributeInfo       `json:"attributeInfo,omitempty"`          // Deprecated
	CdpLevel               *int                                                                 `json:"cdpLevel,omitempty"`               // CDP level to which neighbor devices to be discovered
	DeviceIDs              string                                                               `json:"deviceIds,omitempty"`              // Ids of the devices discovered in a discovery
	DiscoveryCondition     string                                                               `json:"discoveryCondition,omitempty"`     // To indicate the discovery status. Available options: Complete or In Progress
	DiscoveryStatus        string                                                               `json:"discoveryStatus,omitempty"`        // Status of the discovery. Available options are: Active, Inactive, Edit
	DiscoveryType          string                                                               `json:"discoveryType,omitempty"`          // Type of the discovery. 'Single', 'Range', 'CDP', 'LLDP', 'CIDR'
	EnablePasswordList     string                                                               `json:"enablePasswordList,omitempty"`     // Enable Password of the devices to be discovered
	GlobalCredentialIDList []string                                                             `json:"globalCredentialIdList,omitempty"` // List of global credential ids to be used
	HTTPReadCredential     *ResponseDiscoveryGetDiscoveriesByRangeV1ResponseHTTPReadCredential  `json:"httpReadCredential,omitempty"`     //
	HTTPWriteCredential    *ResponseDiscoveryGetDiscoveriesByRangeV1ResponseHTTPWriteCredential `json:"httpWriteCredential,omitempty"`    //
	ID                     string                                                               `json:"id,omitempty"`                     // Unique Discovery Id
	IPAddressList          string                                                               `json:"ipAddressList,omitempty"`          // List of IP address of the devices to be discovered
	IPFilterList           string                                                               `json:"ipFilterList,omitempty"`           // IP addresses of the devices to be filtered
	IsAutoCdp              *bool                                                                `json:"isAutoCdp,omitempty"`              // Flag to mention if CDP discovery or not
	LldpLevel              *int                                                                 `json:"lldpLevel,omitempty"`              // LLDP level to which neighbor devices to be discovered
	Name                   string                                                               `json:"name,omitempty"`                   // Name for the discovery
	NetconfPort            string                                                               `json:"netconfPort,omitempty"`            // Netconf port on the device. Netconf will need valid sshv2 credentials for it to work
	NumDevices             *int                                                                 `json:"numDevices,omitempty"`             // Number of devices discovered in the discovery
	ParentDiscoveryID      string                                                               `json:"parentDiscoveryId,omitempty"`      // Parent Discovery Id from which the discovery was initiated
	PasswordList           string                                                               `json:"passwordList,omitempty"`           // Password of the devices to be discovered
	PreferredMgmtIPMethod  string                                                               `json:"preferredMgmtIPMethod,omitempty"`  // Preferred management IP method. Available options are '' and 'UseLoopBack'
	ProtocolOrder          string                                                               `json:"protocolOrder,omitempty"`          // Order of protocol (ssh/telnet) in which device connection will be tried. Ex: 'telnet': only telnet; 'ssh,telnet': ssh with higher order than telnet
	RetryCount             *int                                                                 `json:"retryCount,omitempty"`             // Number of times to try establishing connection to device
	SNMPAuthPassphrase     string                                                               `json:"snmpAuthPassphrase,omitempty"`     // Auth passphrase for SNMP
	SNMPAuthProtocol       string                                                               `json:"snmpAuthProtocol,omitempty"`       // SNMP auth protocol. SHA' or 'MD5'
	SNMPMode               string                                                               `json:"snmpMode,omitempty"`               // Mode of SNMP. 'AUTHPRIV' or 'AUTHNOPRIV' or 'NOAUTHNOPRIV'
	SNMPPrivPassphrase     string                                                               `json:"snmpPrivPassphrase,omitempty"`     // Passphrase for SNMP privacy
	SNMPPrivProtocol       string                                                               `json:"snmpPrivProtocol,omitempty"`       // SNMP privacy protocol. 'AES128'
	SNMPRoCommunity        string                                                               `json:"snmpRoCommunity,omitempty"`        // SNMP RO community of the devices to be discovered
	SNMPRoCommunityDesc    string                                                               `json:"snmpRoCommunityDesc,omitempty"`    // Description for SNMP RO community
	SNMPRwCommunity        string                                                               `json:"snmpRwCommunity,omitempty"`        // SNMP RW community of the devices to be discovered
	SNMPRwCommunityDesc    string                                                               `json:"snmpRwCommunityDesc,omitempty"`    // Description for SNMP RW community
	SNMPUserName           string                                                               `json:"snmpUserName,omitempty"`           // SNMP username of the device
	TimeOut                *int                                                                 `json:"timeOut,omitempty"`                // Time to wait for device response.
	UpdateMgmtIP           *bool                                                                `json:"updateMgmtIp,omitempty"`           // Updates Management IP if multiple IPs are available for a device. If set to true, when a device is rediscovered with a different IP, the management IP is updated. Default value is false
	UserNameList           string                                                               `json:"userNameList,omitempty"`           // Username of the devices to be discovered
}
type ResponseDiscoveryGetDiscoveriesByRangeV1ResponseAttributeInfo interface{}
type ResponseDiscoveryGetDiscoveriesByRangeV1ResponseHTTPReadCredential struct {
	Comments         string `json:"comments,omitempty"`         // Comments to identify the credential
	CredentialType   string `json:"credentialType,omitempty"`   // Credential type to identify the application that uses the credential
	Description      string `json:"description,omitempty"`      // Description of the credential
	ID               string `json:"id,omitempty"`               // Credential Id
	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Credential Tenant Id
	InstanceUUID     string `json:"instanceUuid,omitempty"`     // Credential Id
	Password         string `json:"password,omitempty"`         // HTTP(S) password
	Port             *int   `json:"port,omitempty"`             // HTTP(S) port
	Secure           *bool  `json:"secure,omitempty"`           // Flag for HTTPS
	Username         string `json:"username,omitempty"`         // HTTP(S) username
}
type ResponseDiscoveryGetDiscoveriesByRangeV1ResponseHTTPWriteCredential struct {
	Comments         string `json:"comments,omitempty"`         // Comments to identify the credential
	CredentialType   string `json:"credentialType,omitempty"`   // Credential type to identify the application that uses the credential
	Description      string `json:"description,omitempty"`      // Description of the credential
	ID               string `json:"id,omitempty"`               // Credential Id
	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Credential Tenant Id
	InstanceUUID     string `json:"instanceUuid,omitempty"`     // Credential Id
	Password         string `json:"password,omitempty"`         // HTTP(S) password
	Port             *int   `json:"port,omitempty"`             // HTTP(S) port
	Secure           *bool  `json:"secure,omitempty"`           // Flag for HTTPS
	Username         string `json:"username,omitempty"`         // HTTP(S) username
}
type ResponseDiscoveryGetGlobalCredentialsV1 struct {
	Response *[]ResponseDiscoveryGetGlobalCredentialsV1Response `json:"response,omitempty"` //
	Version  string                                             `json:"version,omitempty"`  //
}
type ResponseDiscoveryGetGlobalCredentialsV1Response struct {
	Username         string `json:"username,omitempty"`         // CLI Username
	EnablePassword   string `json:"enablePassword,omitempty"`   // CLI Enable Password
	Password         string `json:"password,omitempty"`         // CLI Password
	NetconfPort      string `json:"netconfPort,omitempty"`      // Netconf Port
	ReadCommunity    string `json:"readCommunity,omitempty"`    // SNMP Read Community
	WriteCommunity   string `json:"writeCommunity,omitempty"`   // SNMP Write Community
	AuthPassword     string `json:"authPassword,omitempty"`     // SNMPV3 Auth Password
	AuthType         string `json:"authType,omitempty"`         // SNMPV3 Auth Type
	PrivacyPassword  string `json:"privacyPassword,omitempty"`  // SNMPV3 Privacy Password
	PrivacyType      string `json:"privacyType,omitempty"`      // SNMPV3 Privacy Type
	SNMPMode         string `json:"snmpMode,omitempty"`         // SNMP Mode
	Secure           string `json:"secure,omitempty"`           // Flag for HTTP(S)
	Port             *int   `json:"port,omitempty"`             // HTTP(S) port
	Comments         string `json:"comments,omitempty"`         // Comments to identify the Global Credential
	CredentialType   string `json:"credentialType,omitempty"`   // Credential type to identify the application that uses the Global credential
	Description      string `json:"description,omitempty"`      // Description for Global Credential
	ID               string `json:"id,omitempty"`               // Id of the Global Credential
	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id of the Global Credential
	InstanceUUID     string `json:"instanceUuid,omitempty"`     // Instance Uuid of the Global Credential
}
type ResponseDiscoveryUpdateCliCredentialsV1 struct {
	Response *ResponseDiscoveryUpdateCliCredentialsV1Response `json:"response,omitempty"` //
	Version  string                                           `json:"version,omitempty"`  //
}
type ResponseDiscoveryUpdateCliCredentialsV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryCreateCliCredentialsV1 struct {
	Response *ResponseDiscoveryCreateCliCredentialsV1Response `json:"response,omitempty"` //
	Version  string                                           `json:"version,omitempty"`  //
}
type ResponseDiscoveryCreateCliCredentialsV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryCreateHTTPReadCredentialsV1 struct {
	Response *ResponseDiscoveryCreateHTTPReadCredentialsV1Response `json:"response,omitempty"` //
	Version  string                                                `json:"version,omitempty"`  //
}
type ResponseDiscoveryCreateHTTPReadCredentialsV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryUpdateHTTPReadCredentialV1 struct {
	Response *ResponseDiscoveryUpdateHTTPReadCredentialV1Response `json:"response,omitempty"` //
	Version  string                                               `json:"version,omitempty"`  //
}
type ResponseDiscoveryUpdateHTTPReadCredentialV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryUpdateHTTPWriteCredentialsV1 struct {
	Response *ResponseDiscoveryUpdateHTTPWriteCredentialsV1Response `json:"response,omitempty"` //
	Version  string                                                 `json:"version,omitempty"`  //
}
type ResponseDiscoveryUpdateHTTPWriteCredentialsV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryCreateHTTPWriteCredentialsV1 struct {
	Response *ResponseDiscoveryCreateHTTPWriteCredentialsV1Response `json:"response,omitempty"` //
	Version  string                                                 `json:"version,omitempty"`  //
}
type ResponseDiscoveryCreateHTTPWriteCredentialsV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryUpdateNetconfCredentialsV1 struct {
	Response *ResponseDiscoveryUpdateNetconfCredentialsV1Response `json:"response,omitempty"` //
	Version  string                                               `json:"version,omitempty"`  //
}
type ResponseDiscoveryUpdateNetconfCredentialsV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryCreateNetconfCredentialsV1 struct {
	Response *ResponseDiscoveryCreateNetconfCredentialsV1Response `json:"response,omitempty"` //
	Version  string                                               `json:"version,omitempty"`  //
}
type ResponseDiscoveryCreateNetconfCredentialsV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryUpdateSNMPReadCommunityV1 struct {
	Response *ResponseDiscoveryUpdateSNMPReadCommunityV1Response `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  //
}
type ResponseDiscoveryUpdateSNMPReadCommunityV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryCreateSNMPReadCommunityV1 struct {
	Response *ResponseDiscoveryCreateSNMPReadCommunityV1Response `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  //
}
type ResponseDiscoveryCreateSNMPReadCommunityV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryCreateSNMPWriteCommunityV1 struct {
	Response *ResponseDiscoveryCreateSNMPWriteCommunityV1Response `json:"response,omitempty"` //
	Version  string                                               `json:"version,omitempty"`  //
}
type ResponseDiscoveryCreateSNMPWriteCommunityV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryUpdateSNMPWriteCommunityV1 struct {
	Response *ResponseDiscoveryUpdateSNMPWriteCommunityV1Response `json:"response,omitempty"` //
	Version  string                                               `json:"version,omitempty"`  //
}
type ResponseDiscoveryUpdateSNMPWriteCommunityV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryUpdateSNMPv3CredentialsV1 struct {
	Response *ResponseDiscoveryUpdateSNMPv3CredentialsV1Response `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  //
}
type ResponseDiscoveryUpdateSNMPv3CredentialsV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryCreateSNMPv3CredentialsV1 struct {
	Response *ResponseDiscoveryCreateSNMPv3CredentialsV1Response `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  //
}
type ResponseDiscoveryCreateSNMPv3CredentialsV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryDeleteGlobalCredentialsByIDV1 struct {
	Response *ResponseDiscoveryDeleteGlobalCredentialsByIDV1Response `json:"response,omitempty"` //
	Version  string                                                  `json:"version,omitempty"`  //
}
type ResponseDiscoveryDeleteGlobalCredentialsByIDV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryUpdateGlobalCredentialsV1 struct {
	Response *ResponseDiscoveryUpdateGlobalCredentialsV1Response `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  //
}
type ResponseDiscoveryUpdateGlobalCredentialsV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryGetCredentialSubTypeByCredentialIDV1 struct {
	Response string `json:"response,omitempty"` // Credential type as 'CLICredential', 'HTTPReadCredential', 'HTTPWriteCredential', 'NetconfCredential', 'SNMPv2ReadCommunity', 'SNMPv2WriteCommunity', 'SNMPv3Credential'
	Version  string `json:"version,omitempty"`  //
}
type ResponseDiscoveryGetSNMPPropertiesV1 struct {
	Response *[]ResponseDiscoveryGetSNMPPropertiesV1Response `json:"response,omitempty"` //
	Version  string                                          `json:"version,omitempty"`  //
}
type ResponseDiscoveryGetSNMPPropertiesV1Response struct {
	ID                 string `json:"id,omitempty"`                 // Id of the SNMP Property
	InstanceTenantID   string `json:"instanceTenantId,omitempty"`   // [Deprecated] InstanceTenantId of the SNMP Property
	InstanceUUID       string `json:"instanceUuid,omitempty"`       // Instance Uuid of the SNMP Property. It is the same as the id. It will be deprecated in future version.
	IntValue           *int   `json:"intValue,omitempty"`           // Integer Value of the SNMP 'Retry' or 'Timeout' property
	SystemPropertyName string `json:"systemPropertyName,omitempty"` // Name of the SNMP Property as 'Retry' or 'Timeout'
}
type ResponseDiscoveryCreateUpdateSNMPPropertiesV1 struct {
	Response *ResponseDiscoveryCreateUpdateSNMPPropertiesV1Response `json:"response,omitempty"` //
	Version  string                                                 `json:"version,omitempty"`  //
}
type ResponseDiscoveryCreateUpdateSNMPPropertiesV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryUpdateGlobalCredentialsV2 struct {
	Response *ResponseDiscoveryUpdateGlobalCredentialsV2Response `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  // Version
}
type ResponseDiscoveryUpdateGlobalCredentialsV2Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseDiscoveryCreateGlobalCredentialsV2 struct {
	Response *ResponseDiscoveryCreateGlobalCredentialsV2Response `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  // Version
}
type ResponseDiscoveryCreateGlobalCredentialsV2Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseDiscoveryGetAllGlobalCredentialsV2 struct {
	Response *ResponseDiscoveryGetAllGlobalCredentialsV2Response `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  // Version
}
type ResponseDiscoveryGetAllGlobalCredentialsV2Response struct {
	CliCredential *[]ResponseDiscoveryGetAllGlobalCredentialsV2ResponseCliCredential `json:"cliCredential,omitempty"` //
	SNMPV2CRead   *[]ResponseDiscoveryGetAllGlobalCredentialsV2ResponseSNMPV2CRead   `json:"snmpV2cRead,omitempty"`   //
	SNMPV2CWrite  *[]ResponseDiscoveryGetAllGlobalCredentialsV2ResponseSNMPV2CWrite  `json:"snmpV2cWrite,omitempty"`  //
	HTTPSRead     *[]ResponseDiscoveryGetAllGlobalCredentialsV2ResponseHTTPSRead     `json:"httpsRead,omitempty"`     //
	HTTPSWrite    *[]ResponseDiscoveryGetAllGlobalCredentialsV2ResponseHTTPSWrite    `json:"httpsWrite,omitempty"`    //
	SNMPV3        *[]ResponseDiscoveryGetAllGlobalCredentialsV2ResponseSNMPV3        `json:"snmpV3,omitempty"`        //
}
type ResponseDiscoveryGetAllGlobalCredentialsV2ResponseCliCredential struct {
	Password         string `json:"password,omitempty"`         // CLI Password
	Username         string `json:"username,omitempty"`         // CLI Username
	EnablePassword   string `json:"enablePassword,omitempty"`   // CLI Enable Password
	Description      string `json:"description,omitempty"`      // Description of the CLI credential
	Comments         string `json:"comments,omitempty"`         // Comments to identify the CLI credential
	CredentialType   string `json:"credentialType,omitempty"`   // Credential type to identify the application that uses the CLI credential
	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id of CLI Credential
	InstanceUUID     string `json:"instanceUuid,omitempty"`     // Instance Uuid of CLI Credential
	ID               string `json:"id,omitempty"`               // Id of the CLI Credential
}
type ResponseDiscoveryGetAllGlobalCredentialsV2ResponseSNMPV2CRead struct {
	ReadCommunity    string `json:"readCommunity,omitempty"`    // Snmp RO community
	Description      string `json:"description,omitempty"`      // Description for Snmp RO community
	Comments         string `json:"comments,omitempty"`         // Comments to identify the SNMP Read credential
	CredentialType   string `json:"credentialType,omitempty"`   // Credential type to identify the application that uses the SNMP Read credential
	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id of SNMP Read Credential
	InstanceUUID     string `json:"instanceUuid,omitempty"`     // Instance Uuid of SNMP Read Credential
	ID               string `json:"id,omitempty"`               // Id of the SNMP Read Credential
}
type ResponseDiscoveryGetAllGlobalCredentialsV2ResponseSNMPV2CWrite struct {
	WriteCommunity   string `json:"writeCommunity,omitempty"`   // Snmp RW community
	Description      string `json:"description,omitempty"`      // Description for Snmp RW community
	Comments         string `json:"comments,omitempty"`         // Comments to identify the SNMP Write credential
	CredentialType   string `json:"credentialType,omitempty"`   // Credential type to identify the application that uses the SNMP Write credential
	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id of SNMP Write Credential
	InstanceUUID     string `json:"instanceUuid,omitempty"`     // Instance Uuid of SNMP Write Credential
	ID               string `json:"id,omitempty"`               // Id of SNMP Write Credential
}
type ResponseDiscoveryGetAllGlobalCredentialsV2ResponseHTTPSRead struct {
	Password         string `json:"password,omitempty"`         // HTTP(S) Read Password
	Port             *int   `json:"port,omitempty"`             // HTTP(S) Port
	Username         string `json:"username,omitempty"`         // HTTP(S) Read Username
	Secure           *bool  `json:"secure,omitempty"`           // Flag for HTTP(S) Read
	Description      string `json:"description,omitempty"`      // Description for HTTP(S) Read Credential
	Comments         string `json:"comments,omitempty"`         // Comments to identify the HTTP(S) Read credential
	CredentialType   string `json:"credentialType,omitempty"`   // Credential type to identify the application that uses the HTTP(S) Read credential
	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id of HTTP(S) Read Credential
	InstanceUUID     string `json:"instanceUuid,omitempty"`     // Instance Uuid of HTTP(S) Read Credential
	ID               string `json:"id,omitempty"`               // Id of the HTTP(S) Read Credential
}
type ResponseDiscoveryGetAllGlobalCredentialsV2ResponseHTTPSWrite struct {
	Password         string `json:"password,omitempty"`         // HTTP(S) Write Password
	Port             *int   `json:"port,omitempty"`             // HTTP(S) Port
	Username         string `json:"username,omitempty"`         // HTTP(S) Write Username
	Secure           *bool  `json:"secure,omitempty"`           // Flag for HTTP(S) Write
	Description      string `json:"description,omitempty"`      // Description for HTTP(S) Write Credetntials
	Comments         string `json:"comments,omitempty"`         // Comments to identify the HTTP(S) Write credential
	CredentialType   string `json:"credentialType,omitempty"`   // Credential type to identify the application that uses the HTTP(S) Write credential
	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id of HTTP(S) Write Credential
	InstanceUUID     string `json:"instanceUuid,omitempty"`     // Instance Uuid of HTTP(S) Write Credential
	ID               string `json:"id,omitempty"`               // Id of the HTTP(S) Write Credential
}
type ResponseDiscoveryGetAllGlobalCredentialsV2ResponseSNMPV3 struct {
	Username         string `json:"username,omitempty"`         // SNMP V3 Username
	AuthPassword     string `json:"authPassword,omitempty"`     // Auth Password for SNMP V3
	AuthType         string `json:"authType,omitempty"`         // SNMP auth protocol. SHA' or 'MD5'
	PrivacyPassword  string `json:"privacyPassword,omitempty"`  // Privacy Password for SNMP privacy
	PrivacyType      string `json:"privacyType,omitempty"`      // SNMP privacy protocol. 'AES128','AES192','AES256'
	SNMPMode         string `json:"snmpMode,omitempty"`         // Mode of SNMP. 'AUTHPRIV' or 'AUTHNOPRIV' or 'NOAUTHNOPRIV'
	Description      string `json:"description,omitempty"`      // Description for Snmp V3 Credential
	Comments         string `json:"comments,omitempty"`         // Comments to identify the SNMP V3 credential
	CredentialType   string `json:"credentialType,omitempty"`   // Credential type to identify the application that uses the SNMP V3 credential
	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id of SNMP V3 Credential
	InstanceUUID     string `json:"instanceUuid,omitempty"`     // Instance Tenant Id of SNMP V3 Credential
	ID               string `json:"id,omitempty"`               // Id of the SNMP V3 Credential
}
type ResponseDiscoveryDeleteGlobalCredentialV2 struct {
	Response *ResponseDiscoveryDeleteGlobalCredentialV2Response `json:"response,omitempty"` //
	Version  string                                             `json:"version,omitempty"`  // Version
}
type ResponseDiscoveryDeleteGlobalCredentialV2Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDV1 struct {
	AttributeInfo          *RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDV1AttributeInfo       `json:"attributeInfo,omitempty"`          // Deprecated
	CdpLevel               *int                                                                          `json:"cdpLevel,omitempty"`               // CDP level to which neighbor devices to be discovered
	DeviceIDs              string                                                                        `json:"deviceIds,omitempty"`              // Ids of the devices discovered in a discovery
	DiscoveryCondition     string                                                                        `json:"discoveryCondition,omitempty"`     // To indicate the discovery status. Available options: Complete or In Progress
	DiscoveryStatus        string                                                                        `json:"discoveryStatus,omitempty"`        // Status of the discovery. Available options are: Active, Inactive, Edit
	DiscoveryType          string                                                                        `json:"discoveryType,omitempty"`          // Type of the discovery. 'Single', 'Range', 'Multi Range', 'CDP', 'LLDP', 'CIDR'
	EnablePasswordList     string                                                                        `json:"enablePasswordList,omitempty"`     // Enable Password of the devices to be discovered
	GlobalCredentialIDList []string                                                                      `json:"globalCredentialIdList,omitempty"` // List of global credential ids to be used
	HTTPReadCredential     *RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDV1HTTPReadCredential  `json:"httpReadCredential,omitempty"`     //
	HTTPWriteCredential    *RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDV1HTTPWriteCredential `json:"httpWriteCredential,omitempty"`    //
	ID                     string                                                                        `json:"id,omitempty"`                     // Unique Discovery Id
	IPAddressList          string                                                                        `json:"ipAddressList,omitempty"`          // List of IP address of the devices to be discovered
	IPFilterList           string                                                                        `json:"ipFilterList,omitempty"`           // IP addresses of the devices to be filtered
	IsAutoCdp              *bool                                                                         `json:"isAutoCdp,omitempty"`              // Flag to mention if CDP discovery or not
	LldpLevel              *int                                                                          `json:"lldpLevel,omitempty"`              // LLDP level to which neighbor devices to be discovered
	Name                   string                                                                        `json:"name,omitempty"`                   // Name for the discovery
	NetconfPort            string                                                                        `json:"netconfPort,omitempty"`            // Netconf port on the device. Netconf will need valid sshv2 credentials for it to work
	NumDevices             *int                                                                          `json:"numDevices,omitempty"`             // Number of devices discovered in the discovery
	ParentDiscoveryID      string                                                                        `json:"parentDiscoveryId,omitempty"`      // Parent Discovery Id from which the discovery was initiated
	PasswordList           string                                                                        `json:"passwordList,omitempty"`           // Password of the devices to be discovered
	PreferredMgmtIPMethod  string                                                                        `json:"preferredMgmtIPMethod,omitempty"`  // Preferred management IP method. Available options are '' and 'UseLoopBack'
	ProtocolOrder          string                                                                        `json:"protocolOrder,omitempty"`          // Order of protocol (ssh/telnet) in which device connection will be tried. Ex: 'telnet': only telnet; 'ssh,telnet': ssh with higher order than telnet
	RetryCount             *int                                                                          `json:"retryCount,omitempty"`             // Number of times to try establishing connection to device
	SNMPAuthPassphrase     string                                                                        `json:"snmpAuthPassphrase,omitempty"`     // Auth passphrase for SNMP
	SNMPAuthProtocol       string                                                                        `json:"snmpAuthProtocol,omitempty"`       // SNMP auth protocol. SHA' or 'MD5'
	SNMPMode               string                                                                        `json:"snmpMode,omitempty"`               // Mode of SNMP. 'AUTHPRIV' or 'AUTHNOPRIV' or 'NOAUTHNOPRIV'
	SNMPPrivPassphrase     string                                                                        `json:"snmpPrivPassphrase,omitempty"`     // Passphrase for SNMP privacy
	SNMPPrivProtocol       string                                                                        `json:"snmpPrivProtocol,omitempty"`       // SNMP privacy protocol. 'AES128'
	SNMPRoCommunity        string                                                                        `json:"snmpRoCommunity,omitempty"`        // SNMP RO community of the devices to be discovered
	SNMPRoCommunityDesc    string                                                                        `json:"snmpRoCommunityDesc,omitempty"`    // Description for SNMP RO community
	SNMPRwCommunity        string                                                                        `json:"snmpRwCommunity,omitempty"`        // SNMP RW community of the devices to be discovered
	SNMPRwCommunityDesc    string                                                                        `json:"snmpRwCommunityDesc,omitempty"`    // Description for SNMP RW community
	SNMPUserName           string                                                                        `json:"snmpUserName,omitempty"`           // SNMP username of the device
	TimeOut                *int                                                                          `json:"timeOut,omitempty"`                // Time to wait for device response.
	UpdateMgmtIP           *bool                                                                         `json:"updateMgmtIp,omitempty"`           // Updates Management IP if multiple IPs are available for a device. If set to true, when a device is rediscovered with a different IP, the management IP is updated. Default value is false
	UserNameList           string                                                                        `json:"userNameList,omitempty"`           // Username of the devices to be discovered
}
type RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDV1AttributeInfo interface{}
type RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDV1HTTPReadCredential struct {
	Comments         string `json:"comments,omitempty"`         // Comments to identify the credential
	CredentialType   string `json:"credentialType,omitempty"`   // Credential type to identify the application that uses the credential
	Description      string `json:"description,omitempty"`      // Description of the credential
	ID               string `json:"id,omitempty"`               // Credential Id
	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Credential Tenant Id
	InstanceUUID     string `json:"instanceUuid,omitempty"`     // Credential Id
	Password         string `json:"password,omitempty"`         // HTTP(S) password
	Port             *int   `json:"port,omitempty"`             // HTTP(S) port
	Secure           *bool  `json:"secure,omitempty"`           // Flag for HTTPS
	Username         string `json:"username,omitempty"`         // HTTP(S) username
}
type RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDV1HTTPWriteCredential struct {
	Comments         string `json:"comments,omitempty"`         // Comments to identify the credential
	CredentialType   string `json:"credentialType,omitempty"`   // Credential type to identify the application that uses the credential
	Description      string `json:"description,omitempty"`      // Description of the credential
	ID               string `json:"id,omitempty"`               // Credential Id
	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Credential Tenant Id
	InstanceUUID     string `json:"instanceUuid,omitempty"`     // Credential Id
	Password         string `json:"password,omitempty"`         // HTTP(S) password
	Port             *int   `json:"port,omitempty"`             // HTTP(S) port
	Secure           *bool  `json:"secure,omitempty"`           // Flag for HTTPS
	Username         string `json:"username,omitempty"`         // HTTP(S) username
}
type RequestDiscoveryStartDiscoveryV1 struct {
	CdpLevel               *int                                                 `json:"cdpLevel,omitempty"`               // CDP level to which neighbor devices are to be discovered
	DiscoveryType          string                                               `json:"discoveryType,omitempty"`          // Type of the discovery. 'Single', 'Range', 'Multi Range', 'CDP', 'LLDP', 'CIDR'
	EnablePasswordList     []string                                             `json:"enablePasswordList,omitempty"`     // Enable Password of the devices to be discovered
	GlobalCredentialIDList []string                                             `json:"globalCredentialIdList,omitempty"` // Global Credential Ids to be used for discovery
	HTTPReadCredential     *RequestDiscoveryStartDiscoveryV1HTTPReadCredential  `json:"httpReadCredential,omitempty"`     //
	HTTPWriteCredential    *RequestDiscoveryStartDiscoveryV1HTTPWriteCredential `json:"httpWriteCredential,omitempty"`    //
	IPAddressList          string                                               `json:"ipAddressList,omitempty"`          // IP Address of devices to be discovered. Ex: '172.30.0.1' for SINGLE, CDP and LLDP; '72.30.0.1-172.30.0.4' for RANGE; '72.30.0.1-172.30.0.4,172.31.0.1-172.31.0.4' for MULTI RANGE; '172.30.0.1/20' for CIDR
	IPFilterList           []string                                             `json:"ipFilterList,omitempty"`           // IP Addresses of the devices to be filtered out during discovery
	LldpLevel              *int                                                 `json:"lldpLevel,omitempty"`              // LLDP level to which neighbor devices to be discovered
	Name                   string                                               `json:"name,omitempty"`                   // Name of the discovery
	NetconfPort            string                                               `json:"netconfPort,omitempty"`            // Netconf Port. It will need valid SSH credentials to work
	PasswordList           []string                                             `json:"passwordList,omitempty"`           // Password of the devices to be discovered
	PreferredMgmtIPMethod  string                                               `json:"preferredMgmtIPMethod,omitempty"`  // Preferred Management IP Method.'' or 'UseLoopBack'. Default is ''
	ProtocolOrder          string                                               `json:"protocolOrder,omitempty"`          // Order of protocol (ssh/telnet) in which device connection will be tried. Ex: 'telnet': only telnet; 'ssh,telnet': ssh with higher order than telnet
	Retry                  *int                                                 `json:"retry,omitempty"`                  // Number of times to try establishing connection to device
	SNMPAuthPassphrase     string                                               `json:"snmpAuthPassphrase,omitempty"`     // Auth passphrase for SNMP
	SNMPAuthProtocol       string                                               `json:"snmpAuthProtocol,omitempty"`       // SNMP auth protocol. SHA' or 'MD5'
	SNMPMode               string                                               `json:"snmpMode,omitempty"`               // Mode of SNMP. 'AUTHPRIV' or 'AUTHNOPRIV' or 'NOAUTHNOPRIV'
	SNMPPrivPassphrase     string                                               `json:"snmpPrivPassphrase,omitempty"`     // Pass phrase for SNMP privacy
	SNMPPrivProtocol       string                                               `json:"snmpPrivProtocol,omitempty"`       // SNMP privacy protocol. 'AES128'
	SNMPROCommunity        string                                               `json:"snmpROCommunity,omitempty"`        // SNMP RO community of the devices to be discovered
	SNMPROCommunityDesc    string                                               `json:"snmpROCommunityDesc,omitempty"`    // Description for SNMP RO community
	SNMPRWCommunity        string                                               `json:"snmpRWCommunity,omitempty"`        // SNMP RW community of the devices to be discovered
	SNMPRWCommunityDesc    string                                               `json:"snmpRWCommunityDesc,omitempty"`    // Description for SNMP RW community
	SNMPUserName           string                                               `json:"snmpUserName,omitempty"`           // SNMP username of the device
	SNMPVersion            string                                               `json:"snmpVersion,omitempty"`            // Version of SNMP. v2 or v3
	Timeout                *int                                                 `json:"timeout,omitempty"`                // Time to wait for device response in seconds
	UserNameList           []string                                             `json:"userNameList,omitempty"`           // Username of the devices to be discovered
}
type RequestDiscoveryStartDiscoveryV1HTTPReadCredential struct {
	Password string `json:"password,omitempty"` // HTTP(S) password
	Port     *int   `json:"port,omitempty"`     // HTTP(S) port
	Secure   *bool  `json:"secure,omitempty"`   // Flag for HTTPS
	Username string `json:"username,omitempty"` // HTTP(S) username
}
type RequestDiscoveryStartDiscoveryV1HTTPWriteCredential struct {
	Password string `json:"password,omitempty"` // HTTP(S) password
	Port     *int   `json:"port,omitempty"`     // HTTP(S) port
	Secure   *bool  `json:"secure,omitempty"`   // Flag for HTTPS
	Username string `json:"username,omitempty"` // HTTP(S) username
}
type RequestDiscoveryUpdateCliCredentialsV1 struct {
	Comments         string `json:"comments,omitempty"`         // Comments to identify the CLI credential
	CredentialType   string `json:"credentialType,omitempty"`   // Credential type to identify the application that uses the CLI credential
	Description      string `json:"description,omitempty"`      // Description for CLI Credentials
	EnablePassword   string `json:"enablePassword,omitempty"`   // CLI Enable Password
	ID               string `json:"id,omitempty"`               // Id of the CLI Credential in UUID format
	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Deprecated
	InstanceUUID     string `json:"instanceUuid,omitempty"`     // Deprecated
	Password         string `json:"password,omitempty"`         // CLI Password
	Username         string `json:"username,omitempty"`         // CLI Username
}
type RequestDiscoveryCreateCliCredentialsV1 []RequestItemDiscoveryCreateCliCredentialsV1 // Array of RequestDiscoveryCreateCLICredentialsV1
type RequestItemDiscoveryCreateCliCredentialsV1 struct {
	Comments         string `json:"comments,omitempty"`         // Comments to identify the CLI credential
	CredentialType   string `json:"credentialType,omitempty"`   // Credential type to identify the application that uses the CLI credential
	Description      string `json:"description,omitempty"`      // Description for CLI Credentials
	EnablePassword   string `json:"enablePassword,omitempty"`   // CLI Enable Password
	ID               string `json:"id,omitempty"`               // Deprecated
	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Deprecated
	InstanceUUID     string `json:"instanceUuid,omitempty"`     // Deprecated
	Password         string `json:"password,omitempty"`         // CLI Password
	Username         string `json:"username,omitempty"`         // CLI Username
}
type RequestDiscoveryCreateHTTPReadCredentialsV1 []RequestItemDiscoveryCreateHTTPReadCredentialsV1 // Array of RequestDiscoveryCreateHTTPReadCredentialsV1
type RequestItemDiscoveryCreateHTTPReadCredentialsV1 struct {
	Comments         string `json:"comments,omitempty"`         // Comments to identify the HTTP(S) Read credential
	CredentialType   string `json:"credentialType,omitempty"`   // Credential type to identify the application that uses the HTTP(S) Read credential
	Description      string `json:"description,omitempty"`      // Description for HTTP(S) Read Credential
	ID               string `json:"id,omitempty"`               // Deprecated
	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Deprecated
	InstanceUUID     string `json:"instanceUuid,omitempty"`     // Deprecated
	Password         string `json:"password,omitempty"`         // HTTP(S) Read Password
	Port             *int   `json:"port,omitempty"`             // HTTP(S) Port. Valid port should be in the range of 1 to 65535.
	Secure           *bool  `json:"secure,omitempty"`           // Flag for HTTPS Read
	Username         string `json:"username,omitempty"`         // HTTP(S) Read Username
}
type RequestDiscoveryUpdateHTTPReadCredentialV1 struct {
	Comments         string `json:"comments,omitempty"`         // Comments to identify the HTTP(S) Read credential
	CredentialType   string `json:"credentialType,omitempty"`   // Credential type to identify the application that uses the HTTP(S) Read credential
	Description      string `json:"description,omitempty"`      // Description for HTTP(S) Read Credential
	ID               string `json:"id,omitempty"`               // Id of the HTTP(S) Read Credential in UUID format
	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Deprecated
	InstanceUUID     string `json:"instanceUuid,omitempty"`     // Deprecated
	Password         string `json:"password,omitempty"`         // HTTP(S) Read Password
	Port             *int   `json:"port,omitempty"`             // HTTP(S) Port. Valid port should be in the range of 1 to 65535.
	Secure           *bool  `json:"secure,omitempty"`           // Flag for HTTPS Read
	Username         string `json:"username,omitempty"`         // HTTP(S) Read Username
}
type RequestDiscoveryUpdateHTTPWriteCredentialsV1 struct {
	Comments         string `json:"comments,omitempty"`         // Comments to identify the HTTP(S) Write credential
	CredentialType   string `json:"credentialType,omitempty"`   // Credential type to identify the application that uses the HTTP(S) Write credential
	Description      string `json:"description,omitempty"`      // Description for HTTP(S) Write Credential
	ID               string `json:"id,omitempty"`               // Id of the HTTP(S) Write Credential in UUID format
	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Deprecated
	InstanceUUID     string `json:"instanceUuid,omitempty"`     // Deprecated
	Password         string `json:"password,omitempty"`         // HTTP(S) Write Password
	Port             *int   `json:"port,omitempty"`             // HTTP(S) Port. Valid port should be in the range of 1 to 65535.
	Secure           *bool  `json:"secure,omitempty"`           // Flag for HTTPS Write
	Username         string `json:"username,omitempty"`         // HTTP(S) Write Username
}
type RequestDiscoveryCreateHTTPWriteCredentialsV1 []RequestItemDiscoveryCreateHTTPWriteCredentialsV1 // Array of RequestDiscoveryCreateHTTPWriteCredentialsV1
type RequestItemDiscoveryCreateHTTPWriteCredentialsV1 struct {
	Comments         string `json:"comments,omitempty"`         // Comments to identify the HTTP(S) Write credential
	CredentialType   string `json:"credentialType,omitempty"`   // Credential type to identify the application that uses the HTTP(S) Write credential
	Description      string `json:"description,omitempty"`      // Description for HTTP(S) Write Credential
	ID               string `json:"id,omitempty"`               // Deprecated
	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Deprecated
	InstanceUUID     string `json:"instanceUuid,omitempty"`     // Deprecated
	Password         string `json:"password,omitempty"`         // HTTP(S) Write Password
	Port             *int   `json:"port,omitempty"`             // HTTP(S) Port. Valid port should be in the range of 1 to 65535.
	Secure           *bool  `json:"secure,omitempty"`           // Flag for HTTPS Write
	Username         string `json:"username,omitempty"`         // HTTP(S) Write Username
}
type RequestDiscoveryUpdateNetconfCredentialsV1 struct {
	Comments         string `json:"comments,omitempty"`         // Comments to identify the netconf credential
	CredentialType   string `json:"credentialType,omitempty"`   // Credential type to identify the application that uses the netconf credential
	Description      string `json:"description,omitempty"`      // Description for Netconf Credentials
	ID               string `json:"id,omitempty"`               // Id of the Netconf Credential in UUID format
	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Deprecated
	InstanceUUID     string `json:"instanceUuid,omitempty"`     // Deprecated
	NetconfPort      string `json:"netconfPort,omitempty"`      // Netconf port on the device. Valid port should be in the range of 1 to 65535.
}
type RequestDiscoveryCreateNetconfCredentialsV1 []RequestItemDiscoveryCreateNetconfCredentialsV1 // Array of RequestDiscoveryCreateNetconfCredentialsV1
type RequestItemDiscoveryCreateNetconfCredentialsV1 struct {
	Comments         string `json:"comments,omitempty"`         // Comments to identify the netconf credential
	CredentialType   string `json:"credentialType,omitempty"`   // Credential type to identify the application that uses the credential
	Description      string `json:"description,omitempty"`      // Description for Netconf Credentials
	ID               string `json:"id,omitempty"`               // Deprecated
	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Deprecated
	InstanceUUID     string `json:"instanceUuid,omitempty"`     // deprecated
	NetconfPort      string `json:"netconfPort,omitempty"`      // Netconf port on the device. Valid port should be in the range of 1 to 65535.
}
type RequestDiscoveryUpdateSNMPReadCommunityV1 struct {
	Comments       string `json:"comments,omitempty"`       // Comments to identify the credential
	CredentialType string `json:"credentialType,omitempty"` // Credential type to identify the application that uses the credential
	Description    string `json:"description,omitempty"`    // Name/Description of the credential
	InstanceUUID   string `json:"instanceUuid,omitempty"`   // Credential UUID
	ReadCommunity  string `json:"readCommunity,omitempty"`  // SNMP read community. NO!$DATA!$ for no value change
}
type RequestDiscoveryCreateSNMPReadCommunityV1 []RequestItemDiscoveryCreateSNMPReadCommunityV1 // Array of RequestDiscoveryCreateSNMPReadCommunityV1
type RequestItemDiscoveryCreateSNMPReadCommunityV1 struct {
	Comments       string `json:"comments,omitempty"`       // Comments to identify the credential
	CredentialType string `json:"credentialType,omitempty"` // Credential type to identify the application that uses the credential
	Description    string `json:"description,omitempty"`    // Name/Description of the credential
	ReadCommunity  string `json:"readCommunity,omitempty"`  // SNMP read community
}
type RequestDiscoveryCreateSNMPWriteCommunityV1 []RequestItemDiscoveryCreateSNMPWriteCommunityV1 // Array of RequestDiscoveryCreateSNMPWriteCommunityV1
type RequestItemDiscoveryCreateSNMPWriteCommunityV1 struct {
	Comments       string `json:"comments,omitempty"`       // Comments to identify the credential
	CredentialType string `json:"credentialType,omitempty"` // Credential type to identify the application that uses the credential
	Description    string `json:"description,omitempty"`    // Name/Description of the credential
	WriteCommunity string `json:"writeCommunity,omitempty"` // SNMP write community
}
type RequestDiscoveryUpdateSNMPWriteCommunityV1 struct {
	Comments       string `json:"comments,omitempty"`       // Comments to identify the credential
	CredentialType string `json:"credentialType,omitempty"` // Credential type to identify the application that uses the credential
	Description    string `json:"description,omitempty"`    // Name/Description of the credential
	InstanceUUID   string `json:"instanceUuid,omitempty"`   // Credential UUID
	WriteCommunity string `json:"writeCommunity,omitempty"` // SNMP write community. NO!$DATA!$ for no value change
}
type RequestDiscoveryUpdateSNMPv3CredentialsV1 struct {
	AuthPassword     string `json:"authPassword,omitempty"`     // Auth password for SNMPv3. Required if snmpMode is 'AUTHPRIV' or 'AUTHNOPRIV'. Use 'NO!$DATA!$' if no change required.
	AuthType         string `json:"authType,omitempty"`         // SNMPv3 auth protocol. 'SHA' or 'MD5'. Required if snmpMode is 'AUTHPRIV' or 'AUTHNOPRIV'.
	Comments         string `json:"comments,omitempty"`         // Comments to identify the SNMPv3 credential
	CredentialType   string `json:"credentialType,omitempty"`   // Credential type to identify the application that uses the SNMPv3 credential
	Description      string `json:"description,omitempty"`      // Description for SNMPv3 Credential
	ID               string `json:"id,omitempty"`               // Id of the SNMPv3 Credential
	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Deprecated. This attribute will be removed in a future release, should not be used, and any value supplied will be ignored.
	InstanceUUID     string `json:"instanceUuid,omitempty"`     // Deprecated. This attribute will be removed in a future release, should not be used, and any value supplied will be ignored.
	PrivacyPassword  string `json:"privacyPassword,omitempty"`  // Privacy password for SNMPv3 privacy. Required if snmpMode is 'AUTHPRIV'. Use 'NO!$DATA!$' if no change required.
	PrivacyType      string `json:"privacyType,omitempty"`      // SNMPv3 privacy protocol. Required is snmpMode is 'AUTHPRIV'.
	SNMPMode         string `json:"snmpMode,omitempty"`         // Mode of SNMPv3. 'AUTHPRIV' or 'AUTHNOPRIV' or 'NOAUTHNOPRIV'
	Username         string `json:"username,omitempty"`         // SNMPv3 username
}
type RequestDiscoveryCreateSNMPv3CredentialsV1 []RequestItemDiscoveryCreateSNMPv3CredentialsV1 // Array of RequestDiscoveryCreateSNMPv3CredentialsV1
type RequestItemDiscoveryCreateSNMPv3CredentialsV1 struct {
	AuthPassword     string `json:"authPassword,omitempty"`     // Auth password for SNMPv3. Required if snmpMode is 'AUTHPRIV' or 'AUTHNOPRIV'.
	AuthType         string `json:"authType,omitempty"`         // SNMPv3 auth protocol. 'SHA' or 'MD5'. Required if snmpMode is 'AUTHPRIV' or 'AUTHNOPRIV'.
	Comments         string `json:"comments,omitempty"`         // Comments to identify the SNMPv3 credential
	CredentialType   string `json:"credentialType,omitempty"`   // Credential type to identify the application that uses the SNMPv3 credential
	Description      string `json:"description,omitempty"`      // Description for the SNMPv3 credential
	ID               string `json:"id,omitempty"`               // Deprecated. This attribute will be removed in a future release, should not be used, and any value supplied will be ignored.
	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Deprecated. This attribute will be removed in a future release, should not be used, and any value supplied will be ignored.
	InstanceUUID     string `json:"instanceUuid,omitempty"`     // Deprecated. This attribute will be removed in a future release, should not be used, and any value supplied will be ignored.
	PrivacyPassword  string `json:"privacyPassword,omitempty"`  // Privacy password for SNMPv3 privacy. Required if snmpMode is 'AUTHPRIV'.
	PrivacyType      string `json:"privacyType,omitempty"`      // SNMPv3 privacy protocol. Required is snmpMode is 'AUTHPRIV'.
	SNMPMode         string `json:"snmpMode,omitempty"`         // Mode of SNMPv3. 'AUTHPRIV' or 'AUTHNOPRIV' or 'NOAUTHNOPRIV'.
	Username         string `json:"username,omitempty"`         // SNMPv3 username
}
type RequestDiscoveryUpdateGlobalCredentialsV1 struct {
	SiteUUIDs []string `json:"siteUuids,omitempty"` // List of siteUuids where credential is to be updated
}
type RequestDiscoveryCreateUpdateSNMPPropertiesV1 []RequestItemDiscoveryCreateUpdateSNMPPropertiesV1 // Array of RequestDiscoveryCreateUpdateSNMPPropertiesV1
type RequestItemDiscoveryCreateUpdateSNMPPropertiesV1 struct {
	ID                 string `json:"id,omitempty"`                 //
	InstanceTenantID   string `json:"instanceTenantId,omitempty"`   //
	InstanceUUID       string `json:"instanceUuid,omitempty"`       //
	IntValue           *int   `json:"intValue,omitempty"`           //
	SystemPropertyName string `json:"systemPropertyName,omitempty"` //
}
type RequestDiscoveryUpdateGlobalCredentialsV2 struct {
	CliCredential *RequestDiscoveryUpdateGlobalCredentialsV2CliCredential `json:"cliCredential,omitempty"` //
	SNMPV2CRead   *RequestDiscoveryUpdateGlobalCredentialsV2SNMPV2CRead   `json:"snmpV2cRead,omitempty"`   //
	SNMPV2CWrite  *RequestDiscoveryUpdateGlobalCredentialsV2SNMPV2CWrite  `json:"snmpV2cWrite,omitempty"`  //
	SNMPV3        *RequestDiscoveryUpdateGlobalCredentialsV2SNMPV3        `json:"snmpV3,omitempty"`        //
	HTTPSRead     *RequestDiscoveryUpdateGlobalCredentialsV2HTTPSRead     `json:"httpsRead,omitempty"`     //
	HTTPSWrite    *RequestDiscoveryUpdateGlobalCredentialsV2HTTPSWrite    `json:"httpsWrite,omitempty"`    //
}
type RequestDiscoveryUpdateGlobalCredentialsV2CliCredential struct {
	Description    string `json:"description,omitempty"`    // Description for CLI credential
	Username       string `json:"username,omitempty"`       // CLI Username
	Password       string `json:"password,omitempty"`       // CLI Password
	EnablePassword string `json:"enablePassword,omitempty"` // CLI Enable Password
	ID             string `json:"id,omitempty"`             // Id of the CLI Credential in UUID format
}
type RequestDiscoveryUpdateGlobalCredentialsV2SNMPV2CRead struct {
	Description   string `json:"description,omitempty"`   // Description for Snmp RO community
	ReadCommunity string `json:"readCommunity,omitempty"` // Snmp RO community
	ID            string `json:"id,omitempty"`            // Id of the SNMP Read Credential in UUID format
}
type RequestDiscoveryUpdateGlobalCredentialsV2SNMPV2CWrite struct {
	Description    string `json:"description,omitempty"`    // Description for Snmp RW community
	WriteCommunity string `json:"writeCommunity,omitempty"` // Snmp RW community
	ID             string `json:"id,omitempty"`             // Id of the SNMP Write Credential in UUID format
}
type RequestDiscoveryUpdateGlobalCredentialsV2SNMPV3 struct {
	AuthPassword    string `json:"authPassword,omitempty"`    // Auth Password for SNMP V3
	AuthType        string `json:"authType,omitempty"`        // SNMP auth protocol. SHA' or 'MD5'
	SNMPMode        string `json:"snmpMode,omitempty"`        // Mode of SNMP. 'AUTHPRIV' or 'AUTHNOPRIV' or 'NOAUTHNOPRIV'
	PrivacyPassword string `json:"privacyPassword,omitempty"` // Privacy Password for SNMP privacy
	PrivacyType     string `json:"privacyType,omitempty"`     // SNMP privacy protocol. 'AES128','AES192','AES256'
	Username        string `json:"username,omitempty"`        // SNMP V3 Username
	Description     string `json:"description,omitempty"`     // Description for Snmp V3 Credential
	ID              string `json:"id,omitempty"`              // Id of the SNMP V3 Credential in UUID format
}
type RequestDiscoveryUpdateGlobalCredentialsV2HTTPSRead struct {
	Description string `json:"description,omitempty"` // Description for HTTP(S) Read Credentials
	Username    string `json:"username,omitempty"`    // HTTP(S) Read Username
	Password    string `json:"password,omitempty"`    // HTTP(S) Read Password
	Port        *int   `json:"port,omitempty"`        // HTTP(S) Port
	ID          string `json:"id,omitempty"`          // Id of the HTTP(S) Read Credential in UUID format
}
type RequestDiscoveryUpdateGlobalCredentialsV2HTTPSWrite struct {
	Description string `json:"description,omitempty"` // Description for HTTP(S) Write Credentials
	Username    string `json:"username,omitempty"`    // HTTP(S) Write Username
	Password    string `json:"password,omitempty"`    // HTTP(S) Write Password
	Port        *int   `json:"port,omitempty"`        // HTTP(S) Port
	ID          string `json:"id,omitempty"`          // Id of the HTTP(S) Read Credential in UUID format
}
type RequestDiscoveryCreateGlobalCredentialsV2 struct {
	CliCredential *[]RequestDiscoveryCreateGlobalCredentialsV2CliCredential `json:"cliCredential,omitempty"` //
	SNMPV2CRead   *[]RequestDiscoveryCreateGlobalCredentialsV2SNMPV2CRead   `json:"snmpV2cRead,omitempty"`   //
	SNMPV2CWrite  *[]RequestDiscoveryCreateGlobalCredentialsV2SNMPV2CWrite  `json:"snmpV2cWrite,omitempty"`  //
	SNMPV3        *[]RequestDiscoveryCreateGlobalCredentialsV2SNMPV3        `json:"snmpV3,omitempty"`        //
	HTTPSRead     *[]RequestDiscoveryCreateGlobalCredentialsV2HTTPSRead     `json:"httpsRead,omitempty"`     //
	HTTPSWrite    *[]RequestDiscoveryCreateGlobalCredentialsV2HTTPSWrite    `json:"httpsWrite,omitempty"`    //
}
type RequestDiscoveryCreateGlobalCredentialsV2CliCredential struct {
	Description    string `json:"description,omitempty"`    // Description for CLI credential
	Username       string `json:"username,omitempty"`       // CLI Username
	Password       string `json:"password,omitempty"`       // CLI Password
	EnablePassword string `json:"enablePassword,omitempty"` // CLI Enable Password
}
type RequestDiscoveryCreateGlobalCredentialsV2SNMPV2CRead struct {
	Description   string `json:"description,omitempty"`   // Description for Snmp RO community
	ReadCommunity string `json:"readCommunity,omitempty"` // Snmp RO community
}
type RequestDiscoveryCreateGlobalCredentialsV2SNMPV2CWrite struct {
	Description    string `json:"description,omitempty"`    // Description for Snmp RW community
	WriteCommunity string `json:"writeCommunity,omitempty"` // Snmp RW community
}
type RequestDiscoveryCreateGlobalCredentialsV2SNMPV3 struct {
	Description     string `json:"description,omitempty"`     // Description for Snmp V3 Credential
	Username        string `json:"username,omitempty"`        // SNMP V3 Username
	PrivacyType     string `json:"privacyType,omitempty"`     // SNMP privacy protocol. 'AES128','AES192','AES256'
	PrivacyPassword string `json:"privacyPassword,omitempty"` // Privacy Password for SNMP privacy
	AuthType        string `json:"authType,omitempty"`        // SNMP auth protocol. SHA' or 'MD5'
	AuthPassword    string `json:"authPassword,omitempty"`    // Auth Password for SNMP
	SNMPMode        string `json:"snmpMode,omitempty"`        // Mode of SNMP. 'AUTHPRIV' or 'AUTHNOPRIV' or 'NOAUTHNOPRIV'
}
type RequestDiscoveryCreateGlobalCredentialsV2HTTPSRead struct {
	Description string `json:"description,omitempty"` // Description for HTTP(S) Read Credentials
	Username    string `json:"username,omitempty"`    // HTTP(S) Read Username
	Password    string `json:"password,omitempty"`    // HTTP(S) Read Password
	Port        *int   `json:"port,omitempty"`        // HTTP(S) Port
}
type RequestDiscoveryCreateGlobalCredentialsV2HTTPSWrite struct {
	Description string `json:"description,omitempty"` // Description for HTTP(S) Write Credentials
	Username    string `json:"username,omitempty"`    // HTTP(S) Write Username
	Password    string `json:"password,omitempty"`    // HTTP(S) Write Password
	Port        *int   `json:"port,omitempty"`        // HTTP(S) Port
}

//GetCountOfAllDiscoveryJobsV1 Get count of all discovery jobs - 069d-9823-451b-892d
/* Returns the count of all available discovery jobs



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-count-of-all-discovery-jobs-v1
*/
func (s *DiscoveryService) GetCountOfAllDiscoveryJobsV1() (*ResponseDiscoveryGetCountOfAllDiscoveryJobsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/discovery/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDiscoveryGetCountOfAllDiscoveryJobsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCountOfAllDiscoveryJobsV1()
		}
		return nil, response, fmt.Errorf("error with operation GetCountOfAllDiscoveryJobsV1")
	}

	result := response.Result().(*ResponseDiscoveryGetCountOfAllDiscoveryJobsV1)
	return result, response, err

}

//GetDiscoveryJobsByIPV1 Get Discovery jobs by IP - a496-7be6-4dfa-aa1a
/* Returns the list of discovery jobs for the given IP


@param GetDiscoveryJobsByIPV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-discovery-jobs-by-ip-v1
*/
func (s *DiscoveryService) GetDiscoveryJobsByIPV1(GetDiscoveryJobsByIPV1QueryParams *GetDiscoveryJobsByIPV1QueryParams) (*ResponseDiscoveryGetDiscoveryJobsByIPV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/discovery/job"

	queryString, _ := query.Values(GetDiscoveryJobsByIPV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDiscoveryGetDiscoveryJobsByIPV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDiscoveryJobsByIPV1(GetDiscoveryJobsByIPV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDiscoveryJobsByIpV1")
	}

	result := response.Result().(*ResponseDiscoveryGetDiscoveryJobsByIPV1)
	return result, response, err

}

//GetDiscoveryByIDV1 Get Discovery by Id - 63bb-88b7-4f59-aa17
/* Returns discovery by Discovery ID. Discovery ID can be obtained using the "Get Discoveries by range" API.


@param id id path parameter. Discovery ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-discovery-by-id-v1
*/
func (s *DiscoveryService) GetDiscoveryByIDV1(id string) (*ResponseDiscoveryGetDiscoveryByIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/discovery/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDiscoveryGetDiscoveryByIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDiscoveryByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation GetDiscoveryByIdV1")
	}

	result := response.Result().(*ResponseDiscoveryGetDiscoveryByIDV1)
	return result, response, err

}

//GetListOfDiscoveriesByDiscoveryIDV1 Get list of discoveries by discovery Id - 9987-2a13-4d0a-9fb4
/* Returns the list of discovery jobs for the given Discovery ID. The results can be optionally filtered based on IP. Discovery ID can be obtained using the "Get Discoveries by range" API.


@param id id path parameter. Discovery ID

@param GetListOfDiscoveriesByDiscoveryIdV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-list-of-discoveries-by-discovery-id-v1
*/
func (s *DiscoveryService) GetListOfDiscoveriesByDiscoveryIDV1(id string, GetListOfDiscoveriesByDiscoveryIdV1QueryParams *GetListOfDiscoveriesByDiscoveryIDV1QueryParams) (*ResponseDiscoveryGetListOfDiscoveriesByDiscoveryIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/discovery/{id}/job"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetListOfDiscoveriesByDiscoveryIdV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDiscoveryGetListOfDiscoveriesByDiscoveryIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetListOfDiscoveriesByDiscoveryIDV1(id, GetListOfDiscoveriesByDiscoveryIdV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetListOfDiscoveriesByDiscoveryIdV1")
	}

	result := response.Result().(*ResponseDiscoveryGetListOfDiscoveriesByDiscoveryIDV1)
	return result, response, err

}

//GetDiscoveredNetworkDevicesByDiscoveryIDV1 Get Discovered network devices by discovery Id - f6ac-994f-451b-a011
/* Returns the network devices discovered for the given Discovery ID. Discovery ID can be obtained using the "Get Discoveries by range" API.


@param id id path parameter. Discovery ID

@param GetDiscoveredNetworkDevicesByDiscoveryIdV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-discovered-network-devices-by-discovery-id-v1
*/
func (s *DiscoveryService) GetDiscoveredNetworkDevicesByDiscoveryIDV1(id string, GetDiscoveredNetworkDevicesByDiscoveryIdV1QueryParams *GetDiscoveredNetworkDevicesByDiscoveryIDV1QueryParams) (*ResponseDiscoveryGetDiscoveredNetworkDevicesByDiscoveryIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/discovery/{id}/network-device"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetDiscoveredNetworkDevicesByDiscoveryIdV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDiscoveryGetDiscoveredNetworkDevicesByDiscoveryIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDiscoveredNetworkDevicesByDiscoveryIDV1(id, GetDiscoveredNetworkDevicesByDiscoveryIdV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDiscoveredNetworkDevicesByDiscoveryIdV1")
	}

	result := response.Result().(*ResponseDiscoveryGetDiscoveredNetworkDevicesByDiscoveryIDV1)
	return result, response, err

}

//GetDevicesDiscoveredByIDV1 Get Devices discovered by Id - a696-5b45-4c9a-8663
/* Returns the count of network devices discovered in the given discovery. Discovery ID can be obtained using the "Get Discoveries by range" API.


@param id id path parameter. Discovery ID

@param GetDevicesDiscoveredByIdV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-devices-discovered-by-id-v1
*/
func (s *DiscoveryService) GetDevicesDiscoveredByIDV1(id string, GetDevicesDiscoveredByIdV1QueryParams *GetDevicesDiscoveredByIDV1QueryParams) (*ResponseDiscoveryGetDevicesDiscoveredByIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/discovery/{id}/network-device/count"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetDevicesDiscoveredByIdV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDiscoveryGetDevicesDiscoveredByIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDevicesDiscoveredByIDV1(id, GetDevicesDiscoveredByIdV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDevicesDiscoveredByIdV1")
	}

	result := response.Result().(*ResponseDiscoveryGetDevicesDiscoveredByIDV1)
	return result, response, err

}

//GetDiscoveredDevicesByRangeV1 Get Discovered devices by range - a6b7-98ab-4aca-a34e
/* Returns the network devices discovered for the given discovery and for the given range. The maximum number of records that can be retrieved is 500. Discovery ID can be obtained using the "Get Discoveries by range" API.


@param id id path parameter. Discovery ID

@param startIndex startIndex path parameter. Starting index for the records

@param recordsToReturn recordsToReturn path parameter. Number of records to fetch from the start index

@param GetDiscoveredDevicesByRangeV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-discovered-devices-by-range-v1
*/
func (s *DiscoveryService) GetDiscoveredDevicesByRangeV1(id string, startIndex int, recordsToReturn int, GetDiscoveredDevicesByRangeV1QueryParams *GetDiscoveredDevicesByRangeV1QueryParams) (*ResponseDiscoveryGetDiscoveredDevicesByRangeV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/discovery/{id}/network-device/{startIndex}/{recordsToReturn}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{startIndex}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{recordsToReturn}", fmt.Sprintf("%v", recordsToReturn), -1)

	queryString, _ := query.Values(GetDiscoveredDevicesByRangeV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDiscoveryGetDiscoveredDevicesByRangeV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDiscoveredDevicesByRangeV1(id, startIndex, recordsToReturn, GetDiscoveredDevicesByRangeV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDiscoveredDevicesByRangeV1")
	}

	result := response.Result().(*ResponseDiscoveryGetDiscoveredDevicesByRangeV1)
	return result, response, err

}

//GetNetworkDevicesFromDiscoveryV1 Get network devices from Discovery - 3d9b-99c3-4339-8a27
/* Returns the devices discovered in the given discovery based on given filters. Discovery ID can be obtained using the "Get Discoveries by range" API.


@param id id path parameter. Discovery ID

@param GetNetworkDevicesFromDiscoveryV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-devices-from-discovery-v1
*/
func (s *DiscoveryService) GetNetworkDevicesFromDiscoveryV1(id string, GetNetworkDevicesFromDiscoveryV1QueryParams *GetNetworkDevicesFromDiscoveryV1QueryParams) (*ResponseDiscoveryGetNetworkDevicesFromDiscoveryV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/discovery/{id}/summary"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetNetworkDevicesFromDiscoveryV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDiscoveryGetNetworkDevicesFromDiscoveryV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetNetworkDevicesFromDiscoveryV1(id, GetNetworkDevicesFromDiscoveryV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetNetworkDevicesFromDiscoveryV1")
	}

	result := response.Result().(*ResponseDiscoveryGetNetworkDevicesFromDiscoveryV1)
	return result, response, err

}

//GetDiscoveriesByRangeV1 Get Discoveries by range - 33b7-99d0-4d0a-8907
/* Returns the discoveries by specified range


@param startIndex startIndex path parameter. Starting index for the records

@param recordsToReturn recordsToReturn path parameter. Number of records to fetch from the starting index


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-discoveries-by-range-v1
*/
func (s *DiscoveryService) GetDiscoveriesByRangeV1(startIndex int, recordsToReturn int) (*ResponseDiscoveryGetDiscoveriesByRangeV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/discovery/{startIndex}/{recordsToReturn}"
	path = strings.Replace(path, "{startIndex}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{recordsToReturn}", fmt.Sprintf("%v", recordsToReturn), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDiscoveryGetDiscoveriesByRangeV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDiscoveriesByRangeV1(startIndex, recordsToReturn)
		}
		return nil, response, fmt.Errorf("error with operation GetDiscoveriesByRangeV1")
	}

	result := response.Result().(*ResponseDiscoveryGetDiscoveriesByRangeV1)
	return result, response, err

}

//GetGlobalCredentialsV1 Get Global credentials - ff81-6b8e-4358-97eb
/* Returns global credential for the given credential sub type


@param GetGlobalCredentialsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-global-credentials-v1
*/
func (s *DiscoveryService) GetGlobalCredentialsV1(GetGlobalCredentialsV1QueryParams *GetGlobalCredentialsV1QueryParams) (*ResponseDiscoveryGetGlobalCredentialsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-credential"

	queryString, _ := query.Values(GetGlobalCredentialsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDiscoveryGetGlobalCredentialsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetGlobalCredentialsV1(GetGlobalCredentialsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetGlobalCredentialsV1")
	}

	result := response.Result().(*ResponseDiscoveryGetGlobalCredentialsV1)
	return result, response, err

}

//GetCredentialSubTypeByCredentialIDV1 Get Credential sub type by credential Id - 58a3-699e-489b-9529
/* Returns the credential sub type for the given Id


@param id id path parameter. Global Credential ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-credential-sub-type-by-credential-id-v1
*/
func (s *DiscoveryService) GetCredentialSubTypeByCredentialIDV1(id string) (*ResponseDiscoveryGetCredentialSubTypeByCredentialIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-credential/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDiscoveryGetCredentialSubTypeByCredentialIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCredentialSubTypeByCredentialIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation GetCredentialSubTypeByCredentialIdV1")
	}

	result := response.Result().(*ResponseDiscoveryGetCredentialSubTypeByCredentialIDV1)
	return result, response, err

}

//GetSNMPPropertiesV1 Get SNMP properties - 4497-4ba5-435a-801d
/* Returns SNMP properties



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-snmp-properties-v1
*/
func (s *DiscoveryService) GetSNMPPropertiesV1() (*ResponseDiscoveryGetSNMPPropertiesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/snmp-property"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDiscoveryGetSNMPPropertiesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSNMPPropertiesV1()
		}
		return nil, response, fmt.Errorf("error with operation GetSnmpPropertiesV1")
	}

	result := response.Result().(*ResponseDiscoveryGetSNMPPropertiesV1)
	return result, response, err

}

//GetAllGlobalCredentialsV2 Get All Global Credentials V2 - 6088-4a03-4b5b-8252
/* API to get device credentials' details. It fetches all global credentials of all types at once, without the need to pass any input parameters.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-all-global-credentials-v2
*/
func (s *DiscoveryService) GetAllGlobalCredentialsV2() (*ResponseDiscoveryGetAllGlobalCredentialsV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/global-credential"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDiscoveryGetAllGlobalCredentialsV2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAllGlobalCredentialsV2()
		}
		return nil, response, fmt.Errorf("error with operation GetAllGlobalCredentialsV2")
	}

	result := response.Result().(*ResponseDiscoveryGetAllGlobalCredentialsV2)
	return result, response, err

}

//StartDiscoveryV1 Start discovery - 55b4-39dc-4239-b140
/* Initiates discovery with the given parameters



Documentation Link: https://developer.cisco.com/docs/dna-center/#!start-discovery-v1
*/
func (s *DiscoveryService) StartDiscoveryV1(requestDiscoveryStartDiscoveryV1 *RequestDiscoveryStartDiscoveryV1) (*ResponseDiscoveryStartDiscoveryV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/discovery"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryStartDiscoveryV1).
		SetResult(&ResponseDiscoveryStartDiscoveryV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.StartDiscoveryV1(requestDiscoveryStartDiscoveryV1)
		}

		return nil, response, fmt.Errorf("error with operation StartDiscoveryV1")
	}

	result := response.Result().(*ResponseDiscoveryStartDiscoveryV1)
	return result, response, err

}

//CreateCliCredentialsV1 Create CLI credentials - 948e-a819-4348-bc0b
/* Adds global CLI credential



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-cli-credentials-v1
*/
func (s *DiscoveryService) CreateCliCredentialsV1(requestDiscoveryCreateCLICredentialsV1 *RequestDiscoveryCreateCliCredentialsV1) (*ResponseDiscoveryCreateCliCredentialsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-credential/cli"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryCreateCLICredentialsV1).
		SetResult(&ResponseDiscoveryCreateCliCredentialsV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateCliCredentialsV1(requestDiscoveryCreateCLICredentialsV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateCliCredentialsV1")
	}

	result := response.Result().(*ResponseDiscoveryCreateCliCredentialsV1)
	return result, response, err

}

//CreateHTTPReadCredentialsV1 Create HTTP read credentials - bf85-9ac6-4a0b-a19c
/* Adds HTTP read credentials



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-http-read-credentials-v1
*/
func (s *DiscoveryService) CreateHTTPReadCredentialsV1(requestDiscoveryCreateHTTPReadCredentialsV1 *RequestDiscoveryCreateHTTPReadCredentialsV1) (*ResponseDiscoveryCreateHTTPReadCredentialsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-credential/http-read"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryCreateHTTPReadCredentialsV1).
		SetResult(&ResponseDiscoveryCreateHTTPReadCredentialsV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateHTTPReadCredentialsV1(requestDiscoveryCreateHTTPReadCredentialsV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateHttpReadCredentialsV1")
	}

	result := response.Result().(*ResponseDiscoveryCreateHTTPReadCredentialsV1)
	return result, response, err

}

//CreateHTTPWriteCredentialsV1 Create HTTP write credentials - 4d9c-a8e2-431a-8a24
/* Adds global HTTP write credentials



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-http-write-credentials-v1
*/
func (s *DiscoveryService) CreateHTTPWriteCredentialsV1(requestDiscoveryCreateHTTPWriteCredentialsV1 *RequestDiscoveryCreateHTTPWriteCredentialsV1) (*ResponseDiscoveryCreateHTTPWriteCredentialsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-credential/http-write"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryCreateHTTPWriteCredentialsV1).
		SetResult(&ResponseDiscoveryCreateHTTPWriteCredentialsV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateHTTPWriteCredentialsV1(requestDiscoveryCreateHTTPWriteCredentialsV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateHttpWriteCredentialsV1")
	}

	result := response.Result().(*ResponseDiscoveryCreateHTTPWriteCredentialsV1)
	return result, response, err

}

//CreateNetconfCredentialsV1 Create Netconf credentials - 1792-9bc7-465b-b564
/* Adds global netconf credentials



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-netconf-credentials-v1
*/
func (s *DiscoveryService) CreateNetconfCredentialsV1(requestDiscoveryCreateNetconfCredentialsV1 *RequestDiscoveryCreateNetconfCredentialsV1) (*ResponseDiscoveryCreateNetconfCredentialsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-credential/netconf"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryCreateNetconfCredentialsV1).
		SetResult(&ResponseDiscoveryCreateNetconfCredentialsV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateNetconfCredentialsV1(requestDiscoveryCreateNetconfCredentialsV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateNetconfCredentialsV1")
	}

	result := response.Result().(*ResponseDiscoveryCreateNetconfCredentialsV1)
	return result, response, err

}

//CreateSNMPReadCommunityV1 Create SNMP read community - 7aa3-da9d-4e09-8ef2
/* Adds global SNMP read community



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-snmp-read-community-v1
*/
func (s *DiscoveryService) CreateSNMPReadCommunityV1(requestDiscoveryCreateSNMPReadCommunityV1 *RequestDiscoveryCreateSNMPReadCommunityV1) (*ResponseDiscoveryCreateSNMPReadCommunityV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-credential/snmpv2-read-community"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryCreateSNMPReadCommunityV1).
		SetResult(&ResponseDiscoveryCreateSNMPReadCommunityV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateSNMPReadCommunityV1(requestDiscoveryCreateSNMPReadCommunityV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateSnmpReadCommunityV1")
	}

	result := response.Result().(*ResponseDiscoveryCreateSNMPReadCommunityV1)
	return result, response, err

}

//CreateSNMPWriteCommunityV1 Create SNMP write community - 6bac-b8d1-4639-bdc7
/* Adds global SNMP write community



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-snmp-write-community-v1
*/
func (s *DiscoveryService) CreateSNMPWriteCommunityV1(requestDiscoveryCreateSNMPWriteCommunityV1 *RequestDiscoveryCreateSNMPWriteCommunityV1) (*ResponseDiscoveryCreateSNMPWriteCommunityV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-credential/snmpv2-write-community"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryCreateSNMPWriteCommunityV1).
		SetResult(&ResponseDiscoveryCreateSNMPWriteCommunityV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateSNMPWriteCommunityV1(requestDiscoveryCreateSNMPWriteCommunityV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateSnmpWriteCommunityV1")
	}

	result := response.Result().(*ResponseDiscoveryCreateSNMPWriteCommunityV1)
	return result, response, err

}

//CreateSNMPv3CredentialsV1 Create SNMPv3 credentials - 9796-8808-4b7b-a60d
/* Adds global SNMPv3 credentials



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-snmpv3-credentials-v1
*/
func (s *DiscoveryService) CreateSNMPv3CredentialsV1(requestDiscoveryCreateSNMPv3CredentialsV1 *RequestDiscoveryCreateSNMPv3CredentialsV1) (*ResponseDiscoveryCreateSNMPv3CredentialsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-credential/snmpv3"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryCreateSNMPv3CredentialsV1).
		SetResult(&ResponseDiscoveryCreateSNMPv3CredentialsV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateSNMPv3CredentialsV1(requestDiscoveryCreateSNMPv3CredentialsV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateSnmpv3CredentialsV1")
	}

	result := response.Result().(*ResponseDiscoveryCreateSNMPv3CredentialsV1)
	return result, response, err

}

//CreateUpdateSNMPPropertiesV1 Create/Update SNMP properties - a5ac-9977-4c6b-b541
/* Adds SNMP properties



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-update-snmp-properties-v1
*/
func (s *DiscoveryService) CreateUpdateSNMPPropertiesV1(requestDiscoveryCreateUpdateSNMPPropertiesV1 *RequestDiscoveryCreateUpdateSNMPPropertiesV1) (*ResponseDiscoveryCreateUpdateSNMPPropertiesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/snmp-property"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryCreateUpdateSNMPPropertiesV1).
		SetResult(&ResponseDiscoveryCreateUpdateSNMPPropertiesV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateUpdateSNMPPropertiesV1(requestDiscoveryCreateUpdateSNMPPropertiesV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateUpdateSnmpPropertiesV1")
	}

	result := response.Result().(*ResponseDiscoveryCreateUpdateSNMPPropertiesV1)
	return result, response, err

}

//CreateGlobalCredentialsV2 Create Global Credentials V2 - 4ca1-8b14-4059-82b0
/* API to create new global credentials. Multiple credentials of various types can be passed at once. Please refer sample Request Body for more information.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-global-credentials-v2
*/
func (s *DiscoveryService) CreateGlobalCredentialsV2(requestDiscoveryCreateGlobalCredentialsV2 *RequestDiscoveryCreateGlobalCredentialsV2) (*ResponseDiscoveryCreateGlobalCredentialsV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/global-credential"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryCreateGlobalCredentialsV2).
		SetResult(&ResponseDiscoveryCreateGlobalCredentialsV2{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateGlobalCredentialsV2(requestDiscoveryCreateGlobalCredentialsV2)
		}

		return nil, response, fmt.Errorf("error with operation CreateGlobalCredentialsV2")
	}

	result := response.Result().(*ResponseDiscoveryCreateGlobalCredentialsV2)
	return result, response, err

}

//UpdatesAnExistingDiscoveryBySpecifiedIDV1 Updates an existing discovery by specified Id - 9788-b8fc-4418-831d
/* Stops or starts an existing discovery


 */
func (s *DiscoveryService) UpdatesAnExistingDiscoveryBySpecifiedIDV1(requestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIdV1 *RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDV1) (*ResponseDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/discovery"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIdV1).
		SetResult(&ResponseDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdatesAnExistingDiscoveryBySpecifiedIDV1(requestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIdV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdatesAnExistingDiscoveryBySpecifiedIdV1")
	}

	result := response.Result().(*ResponseDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDV1)
	return result, response, err

}

//UpdateCliCredentialsV1 Update CLI credentials - fba0-d807-47eb-82e8
/* Updates global CLI credentials


 */
func (s *DiscoveryService) UpdateCliCredentialsV1(requestDiscoveryUpdateCLICredentialsV1 *RequestDiscoveryUpdateCliCredentialsV1) (*ResponseDiscoveryUpdateCliCredentialsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-credential/cli"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryUpdateCLICredentialsV1).
		SetResult(&ResponseDiscoveryUpdateCliCredentialsV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateCliCredentialsV1(requestDiscoveryUpdateCLICredentialsV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateCliCredentialsV1")
	}

	result := response.Result().(*ResponseDiscoveryUpdateCliCredentialsV1)
	return result, response, err

}

//UpdateHTTPReadCredentialV1 Update HTTP read credential - 89b3-6b46-4999-9d81
/* Updates global HTTP Read credential


 */
func (s *DiscoveryService) UpdateHTTPReadCredentialV1(requestDiscoveryUpdateHTTPReadCredentialV1 *RequestDiscoveryUpdateHTTPReadCredentialV1) (*ResponseDiscoveryUpdateHTTPReadCredentialV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-credential/http-read"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryUpdateHTTPReadCredentialV1).
		SetResult(&ResponseDiscoveryUpdateHTTPReadCredentialV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateHTTPReadCredentialV1(requestDiscoveryUpdateHTTPReadCredentialV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateHttpReadCredentialV1")
	}

	result := response.Result().(*ResponseDiscoveryUpdateHTTPReadCredentialV1)
	return result, response, err

}

//UpdateHTTPWriteCredentialsV1 Update HTTP write credentials - b68a-6bd8-473a-9a25
/* Updates global HTTP write credentials


 */
func (s *DiscoveryService) UpdateHTTPWriteCredentialsV1(requestDiscoveryUpdateHTTPWriteCredentialsV1 *RequestDiscoveryUpdateHTTPWriteCredentialsV1) (*ResponseDiscoveryUpdateHTTPWriteCredentialsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-credential/http-write"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryUpdateHTTPWriteCredentialsV1).
		SetResult(&ResponseDiscoveryUpdateHTTPWriteCredentialsV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateHTTPWriteCredentialsV1(requestDiscoveryUpdateHTTPWriteCredentialsV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateHttpWriteCredentialsV1")
	}

	result := response.Result().(*ResponseDiscoveryUpdateHTTPWriteCredentialsV1)
	return result, response, err

}

//UpdateNetconfCredentialsV1 Update Netconf credentials - c5ac-d9fa-4c1a-8abc
/* Updates global netconf credentials


 */
func (s *DiscoveryService) UpdateNetconfCredentialsV1(requestDiscoveryUpdateNetconfCredentialsV1 *RequestDiscoveryUpdateNetconfCredentialsV1) (*ResponseDiscoveryUpdateNetconfCredentialsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-credential/netconf"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryUpdateNetconfCredentialsV1).
		SetResult(&ResponseDiscoveryUpdateNetconfCredentialsV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateNetconfCredentialsV1(requestDiscoveryUpdateNetconfCredentialsV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateNetconfCredentialsV1")
	}

	result := response.Result().(*ResponseDiscoveryUpdateNetconfCredentialsV1)
	return result, response, err

}

//UpdateSNMPReadCommunityV1 Update SNMP read community - 47a1-b84b-4e1b-8044
/* Updates global SNMP read community


 */
func (s *DiscoveryService) UpdateSNMPReadCommunityV1(requestDiscoveryUpdateSNMPReadCommunityV1 *RequestDiscoveryUpdateSNMPReadCommunityV1) (*ResponseDiscoveryUpdateSNMPReadCommunityV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-credential/snmpv2-read-community"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryUpdateSNMPReadCommunityV1).
		SetResult(&ResponseDiscoveryUpdateSNMPReadCommunityV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateSNMPReadCommunityV1(requestDiscoveryUpdateSNMPReadCommunityV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateSnmpReadCommunityV1")
	}

	result := response.Result().(*ResponseDiscoveryUpdateSNMPReadCommunityV1)
	return result, response, err

}

//UpdateSNMPWriteCommunityV1 Update SNMP write community - 10b0-6a6a-4f7b-b3cb
/* Updates global SNMP write community


 */
func (s *DiscoveryService) UpdateSNMPWriteCommunityV1(requestDiscoveryUpdateSNMPWriteCommunityV1 *RequestDiscoveryUpdateSNMPWriteCommunityV1) (*ResponseDiscoveryUpdateSNMPWriteCommunityV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-credential/snmpv2-write-community"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryUpdateSNMPWriteCommunityV1).
		SetResult(&ResponseDiscoveryUpdateSNMPWriteCommunityV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateSNMPWriteCommunityV1(requestDiscoveryUpdateSNMPWriteCommunityV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateSnmpWriteCommunityV1")
	}

	result := response.Result().(*ResponseDiscoveryUpdateSNMPWriteCommunityV1)
	return result, response, err

}

//UpdateSNMPv3CredentialsV1 Update SNMPv3 credentials - 1da5-ebdd-434a-acfe
/* Updates global SNMPv3 credential


 */
func (s *DiscoveryService) UpdateSNMPv3CredentialsV1(requestDiscoveryUpdateSNMPv3CredentialsV1 *RequestDiscoveryUpdateSNMPv3CredentialsV1) (*ResponseDiscoveryUpdateSNMPv3CredentialsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-credential/snmpv3"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryUpdateSNMPv3CredentialsV1).
		SetResult(&ResponseDiscoveryUpdateSNMPv3CredentialsV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateSNMPv3CredentialsV1(requestDiscoveryUpdateSNMPv3CredentialsV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateSnmpv3CredentialsV1")
	}

	result := response.Result().(*ResponseDiscoveryUpdateSNMPv3CredentialsV1)
	return result, response, err

}

//UpdateGlobalCredentialsV1 Update global credentials - 709f-da3c-42b8-877a
/* Update global credential for network devices in site(s)


@param globalCredentialID globalCredentialId path parameter. Global credential Uuid

*/
func (s *DiscoveryService) UpdateGlobalCredentialsV1(globalCredentialID string, requestDiscoveryUpdateGlobalCredentialsV1 *RequestDiscoveryUpdateGlobalCredentialsV1) (*ResponseDiscoveryUpdateGlobalCredentialsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-credential/{globalCredentialId}"
	path = strings.Replace(path, "{globalCredentialId}", fmt.Sprintf("%v", globalCredentialID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryUpdateGlobalCredentialsV1).
		SetResult(&ResponseDiscoveryUpdateGlobalCredentialsV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateGlobalCredentialsV1(globalCredentialID, requestDiscoveryUpdateGlobalCredentialsV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateGlobalCredentialsV1")
	}

	result := response.Result().(*ResponseDiscoveryUpdateGlobalCredentialsV1)
	return result, response, err

}

//UpdateGlobalCredentialsV2 Update Global Credentials V2 - a7bb-8baa-487a-acf6
/* API to update device credentials. Multiple credentials can be passed at once, but only a single credential of a given type can be passed at once. Please refer sample Request Body for more information.


 */
func (s *DiscoveryService) UpdateGlobalCredentialsV2(requestDiscoveryUpdateGlobalCredentialsV2 *RequestDiscoveryUpdateGlobalCredentialsV2) (*ResponseDiscoveryUpdateGlobalCredentialsV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/global-credential"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryUpdateGlobalCredentialsV2).
		SetResult(&ResponseDiscoveryUpdateGlobalCredentialsV2{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateGlobalCredentialsV2(requestDiscoveryUpdateGlobalCredentialsV2)
		}
		return nil, response, fmt.Errorf("error with operation UpdateGlobalCredentialsV2")
	}

	result := response.Result().(*ResponseDiscoveryUpdateGlobalCredentialsV2)
	return result, response, err

}

//DeleteAllDiscoveryV1 Delete all discovery - db8e-0923-4a98-8bab
/* Stops all the discoveries and removes them



Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-all-discovery-v1
*/
func (s *DiscoveryService) DeleteAllDiscoveryV1() (*ResponseDiscoveryDeleteAllDiscoveryV1, *resty.Response, error) {
	//
	path := "/dna/intent/api/v1/discovery"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDiscoveryDeleteAllDiscoveryV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteAllDiscoveryV1()
		}
		return nil, response, fmt.Errorf("error with operation DeleteAllDiscoveryV1")
	}

	result := response.Result().(*ResponseDiscoveryDeleteAllDiscoveryV1)
	return result, response, err

}

//DeleteDiscoveryByIDV1 Delete discovery by Id - 4c8c-ab5f-435a-80f4
/* Stops the discovery for the given Discovery ID and removes it. Discovery ID can be obtained using the "Get Discoveries by range" API.


@param id id path parameter. Discovery ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-discovery-by-id-v1
*/
func (s *DiscoveryService) DeleteDiscoveryByIDV1(id string) (*ResponseDiscoveryDeleteDiscoveryByIDV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/discovery/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDiscoveryDeleteDiscoveryByIDV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteDiscoveryByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteDiscoveryByIdV1")
	}

	result := response.Result().(*ResponseDiscoveryDeleteDiscoveryByIDV1)
	return result, response, err

}

//DeleteDiscoveryBySpecifiedRangeV1 Delete discovery by specified range - c1ba-9a42-4c08-a01b
/* Stops discovery for the given range and removes them


@param startIndex startIndex path parameter. Starting index for the records

@param recordsToDelete recordsToDelete path parameter. Number of records to delete from the starting index


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-discovery-by-specified-range-v1
*/
func (s *DiscoveryService) DeleteDiscoveryBySpecifiedRangeV1(startIndex int, recordsToDelete int) (*ResponseDiscoveryDeleteDiscoveryBySpecifiedRangeV1, *resty.Response, error) {
	//startIndex int,recordsToDelete int
	path := "/dna/intent/api/v1/discovery/{startIndex}/{recordsToDelete}"
	path = strings.Replace(path, "{startIndex}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{recordsToDelete}", fmt.Sprintf("%v", recordsToDelete), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDiscoveryDeleteDiscoveryBySpecifiedRangeV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteDiscoveryBySpecifiedRangeV1(startIndex, recordsToDelete)
		}
		return nil, response, fmt.Errorf("error with operation DeleteDiscoveryBySpecifiedRangeV1")
	}

	result := response.Result().(*ResponseDiscoveryDeleteDiscoveryBySpecifiedRangeV1)
	return result, response, err

}

//DeleteGlobalCredentialsByIDV1 Delete global credentials by Id - f5ac-590c-4ca9-975a
/* Deletes global credential for the given ID


@param globalCredentialID globalCredentialId path parameter. ID of global-credential


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-global-credentials-by-id-v1
*/
func (s *DiscoveryService) DeleteGlobalCredentialsByIDV1(globalCredentialID string) (*ResponseDiscoveryDeleteGlobalCredentialsByIDV1, *resty.Response, error) {
	//globalCredentialID string
	path := "/dna/intent/api/v1/global-credential/{globalCredentialId}"
	path = strings.Replace(path, "{globalCredentialId}", fmt.Sprintf("%v", globalCredentialID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDiscoveryDeleteGlobalCredentialsByIDV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteGlobalCredentialsByIDV1(globalCredentialID)
		}
		return nil, response, fmt.Errorf("error with operation DeleteGlobalCredentialsByIdV1")
	}

	result := response.Result().(*ResponseDiscoveryDeleteGlobalCredentialsByIDV1)
	return result, response, err

}

//DeleteGlobalCredentialV2 Delete Global Credential V2 - 6487-3b18-4f48-be33
/* Delete a global credential. Only 'id' of the credential has to be passed.


@param id id path parameter. Global Credential id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-global-credential-v2
*/
func (s *DiscoveryService) DeleteGlobalCredentialV2(id string) (*ResponseDiscoveryDeleteGlobalCredentialV2, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v2/global-credential/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDiscoveryDeleteGlobalCredentialV2{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteGlobalCredentialV2(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteGlobalCredentialV2")
	}

	result := response.Result().(*ResponseDiscoveryDeleteGlobalCredentialV2)
	return result, response, err

}

// Alias Function
func (s *DiscoveryService) CreateSNMPWriteCommunity(requestDiscoveryCreateSNMPWriteCommunityV1 *RequestDiscoveryCreateSNMPWriteCommunityV1) (*ResponseDiscoveryCreateSNMPWriteCommunityV1, *resty.Response, error) {
	return s.CreateSNMPWriteCommunityV1(requestDiscoveryCreateSNMPWriteCommunityV1)
}

// Alias Function
func (s *DiscoveryService) GetDiscoveriesByRange(startIndex int, recordsToReturn int) (*ResponseDiscoveryGetDiscoveriesByRangeV1, *resty.Response, error) {
	return s.GetDiscoveriesByRangeV1(startIndex, recordsToReturn)
}

// Alias Function
func (s *DiscoveryService) GetCountOfAllDiscoveryJobs() (*ResponseDiscoveryGetCountOfAllDiscoveryJobsV1, *resty.Response, error) {
	return s.GetCountOfAllDiscoveryJobsV1()
}

// Alias Function
func (s *DiscoveryService) GetSNMPProperties() (*ResponseDiscoveryGetSNMPPropertiesV1, *resty.Response, error) {
	return s.GetSNMPPropertiesV1()
}

// Alias Function
func (s *DiscoveryService) GetDiscoveredNetworkDevicesByDiscoveryID(id string, GetDiscoveredNetworkDevicesByDiscoveryIdV1QueryParams *GetDiscoveredNetworkDevicesByDiscoveryIDV1QueryParams) (*ResponseDiscoveryGetDiscoveredNetworkDevicesByDiscoveryIDV1, *resty.Response, error) {
	return s.GetDiscoveredNetworkDevicesByDiscoveryIDV1(id, GetDiscoveredNetworkDevicesByDiscoveryIdV1QueryParams)
}

// Alias Function
func (s *DiscoveryService) GetGlobalCredentials(GetGlobalCredentialsV1QueryParams *GetGlobalCredentialsV1QueryParams) (*ResponseDiscoveryGetGlobalCredentialsV1, *resty.Response, error) {
	return s.GetGlobalCredentialsV1(GetGlobalCredentialsV1QueryParams)
}

// Alias Function
func (s *DiscoveryService) GetDevicesDiscoveredByID(id string, GetDevicesDiscoveredByIdV1QueryParams *GetDevicesDiscoveredByIDV1QueryParams) (*ResponseDiscoveryGetDevicesDiscoveredByIDV1, *resty.Response, error) {
	return s.GetDevicesDiscoveredByIDV1(id, GetDevicesDiscoveredByIdV1QueryParams)
}

// Alias Function
func (s *DiscoveryService) DeleteDiscoveryBySpecifiedRange(startIndex int, recordsToDelete int) (*ResponseDiscoveryDeleteDiscoveryBySpecifiedRangeV1, *resty.Response, error) {
	return s.DeleteDiscoveryBySpecifiedRangeV1(startIndex, recordsToDelete)
}

// Alias Function
func (s *DiscoveryService) UpdateNetconfCredentials(requestDiscoveryUpdateNetconfCredentialsV1 *RequestDiscoveryUpdateNetconfCredentialsV1) (*ResponseDiscoveryUpdateNetconfCredentialsV1, *resty.Response, error) {
	return s.UpdateNetconfCredentialsV1(requestDiscoveryUpdateNetconfCredentialsV1)
}

// Alias Function
func (s *DiscoveryService) GetDiscoveredDevicesByRange(id string, startIndex int, recordsToReturn int, GetDiscoveredDevicesByRangeV1QueryParams *GetDiscoveredDevicesByRangeV1QueryParams) (*ResponseDiscoveryGetDiscoveredDevicesByRangeV1, *resty.Response, error) {
	return s.GetDiscoveredDevicesByRangeV1(id, startIndex, recordsToReturn, GetDiscoveredDevicesByRangeV1QueryParams)
}

// Alias Function
func (s *DiscoveryService) DeleteAllDiscovery() (*ResponseDiscoveryDeleteAllDiscoveryV1, *resty.Response, error) {
	return s.DeleteAllDiscoveryV1()
}

// Alias Function
func (s *DiscoveryService) UpdatesAnExistingDiscoveryBySpecifiedID(requestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIdV1 *RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDV1) (*ResponseDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDV1, *resty.Response, error) {
	return s.UpdatesAnExistingDiscoveryBySpecifiedIDV1(requestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIdV1)
}

// Alias Function
func (s *DiscoveryService) UpdateSNMPReadCommunity(requestDiscoveryUpdateSNMPReadCommunityV1 *RequestDiscoveryUpdateSNMPReadCommunityV1) (*ResponseDiscoveryUpdateSNMPReadCommunityV1, *resty.Response, error) {
	return s.UpdateSNMPReadCommunityV1(requestDiscoveryUpdateSNMPReadCommunityV1)
}

// Alias Function
func (s *DiscoveryService) UpdateSNMPv3Credentials(requestDiscoveryUpdateSNMPv3CredentialsV1 *RequestDiscoveryUpdateSNMPv3CredentialsV1) (*ResponseDiscoveryUpdateSNMPv3CredentialsV1, *resty.Response, error) {
	return s.UpdateSNMPv3CredentialsV1(requestDiscoveryUpdateSNMPv3CredentialsV1)
}

// Alias Function
func (s *DiscoveryService) DeleteGlobalCredentialsByID(globalCredentialID string) (*ResponseDiscoveryDeleteGlobalCredentialsByIDV1, *resty.Response, error) {
	return s.DeleteGlobalCredentialsByIDV1(globalCredentialID)
}

// Alias Function
func (s *DiscoveryService) GetListOfDiscoveriesByDiscoveryID(id string, GetListOfDiscoveriesByDiscoveryIdV1QueryParams *GetListOfDiscoveriesByDiscoveryIDV1QueryParams) (*ResponseDiscoveryGetListOfDiscoveriesByDiscoveryIDV1, *resty.Response, error) {
	return s.GetListOfDiscoveriesByDiscoveryIDV1(id, GetListOfDiscoveriesByDiscoveryIdV1QueryParams)
}

// Alias Function
func (s *DiscoveryService) DeleteDiscoveryByID(id string) (*ResponseDiscoveryDeleteDiscoveryByIDV1, *resty.Response, error) {
	return s.DeleteDiscoveryByIDV1(id)
}

// Alias Function
func (s *DiscoveryService) UpdateGlobalCredentials(globalCredentialID string, requestDiscoveryUpdateGlobalCredentialsV1 *RequestDiscoveryUpdateGlobalCredentialsV1) (*ResponseDiscoveryUpdateGlobalCredentialsV1, *resty.Response, error) {
	return s.UpdateGlobalCredentialsV1(globalCredentialID, requestDiscoveryUpdateGlobalCredentialsV1)
}

// Alias Function
func (s *DiscoveryService) GetCredentialSubTypeByCredentialID(id string) (*ResponseDiscoveryGetCredentialSubTypeByCredentialIDV1, *resty.Response, error) {
	return s.GetCredentialSubTypeByCredentialIDV1(id)
}

// Alias Function
func (s *DiscoveryService) CreateGlobalCredentials(requestDiscoveryCreateGlobalCredentialsV2 *RequestDiscoveryCreateGlobalCredentialsV2) (*ResponseDiscoveryCreateGlobalCredentialsV2, *resty.Response, error) {
	return s.CreateGlobalCredentialsV2(requestDiscoveryCreateGlobalCredentialsV2)
}

// Alias Function
func (s *DiscoveryService) GetAllGlobalCredentials() (*ResponseDiscoveryGetAllGlobalCredentialsV2, *resty.Response, error) {
	return s.GetAllGlobalCredentialsV2()
}

// Alias Function
func (s *DiscoveryService) CreateHTTPReadCredentials(requestDiscoveryCreateHTTPReadCredentialsV1 *RequestDiscoveryCreateHTTPReadCredentialsV1) (*ResponseDiscoveryCreateHTTPReadCredentialsV1, *resty.Response, error) {
	return s.CreateHTTPReadCredentialsV1(requestDiscoveryCreateHTTPReadCredentialsV1)
}

// Alias Function
func (s *DiscoveryService) GetNetworkDevicesFromDiscovery(id string, GetNetworkDevicesFromDiscoveryV1QueryParams *GetNetworkDevicesFromDiscoveryV1QueryParams) (*ResponseDiscoveryGetNetworkDevicesFromDiscoveryV1, *resty.Response, error) {
	return s.GetNetworkDevicesFromDiscoveryV1(id, GetNetworkDevicesFromDiscoveryV1QueryParams)
}

// Alias Function
func (s *DiscoveryService) UpdateCliCredentials(requestDiscoveryUpdateCLICredentialsV1 *RequestDiscoveryUpdateCliCredentialsV1) (*ResponseDiscoveryUpdateCliCredentialsV1, *resty.Response, error) {
	return s.UpdateCliCredentialsV1(requestDiscoveryUpdateCLICredentialsV1)
}

// Alias Function
func (s *DiscoveryService) StartDiscovery(requestDiscoveryStartDiscoveryV1 *RequestDiscoveryStartDiscoveryV1) (*ResponseDiscoveryStartDiscoveryV1, *resty.Response, error) {
	return s.StartDiscoveryV1(requestDiscoveryStartDiscoveryV1)
}

// Alias Function
func (s *DiscoveryService) CreateSNMPReadCommunity(requestDiscoveryCreateSNMPReadCommunityV1 *RequestDiscoveryCreateSNMPReadCommunityV1) (*ResponseDiscoveryCreateSNMPReadCommunityV1, *resty.Response, error) {
	return s.CreateSNMPReadCommunityV1(requestDiscoveryCreateSNMPReadCommunityV1)
}

// Alias Function
func (s *DiscoveryService) CreateCliCredentials(requestDiscoveryCreateCLICredentialsV1 *RequestDiscoveryCreateCliCredentialsV1) (*ResponseDiscoveryCreateCliCredentialsV1, *resty.Response, error) {
	return s.CreateCliCredentialsV1(requestDiscoveryCreateCLICredentialsV1)
}

// Alias Function
func (s *DiscoveryService) UpdateHTTPReadCredential(requestDiscoveryUpdateHTTPReadCredentialV1 *RequestDiscoveryUpdateHTTPReadCredentialV1) (*ResponseDiscoveryUpdateHTTPReadCredentialV1, *resty.Response, error) {
	return s.UpdateHTTPReadCredentialV1(requestDiscoveryUpdateHTTPReadCredentialV1)
}

// Alias Function
func (s *DiscoveryService) UpdateHTTPWriteCredentials(requestDiscoveryUpdateHTTPWriteCredentialsV1 *RequestDiscoveryUpdateHTTPWriteCredentialsV1) (*ResponseDiscoveryUpdateHTTPWriteCredentialsV1, *resty.Response, error) {
	return s.UpdateHTTPWriteCredentialsV1(requestDiscoveryUpdateHTTPWriteCredentialsV1)
}

// Alias Function
func (s *DiscoveryService) DeleteGlobalCredential(id string) (*ResponseDiscoveryDeleteGlobalCredentialV2, *resty.Response, error) {
	return s.DeleteGlobalCredentialV2(id)
}

// Alias Function
func (s *DiscoveryService) CreateHTTPWriteCredentials(requestDiscoveryCreateHTTPWriteCredentialsV1 *RequestDiscoveryCreateHTTPWriteCredentialsV1) (*ResponseDiscoveryCreateHTTPWriteCredentialsV1, *resty.Response, error) {
	return s.CreateHTTPWriteCredentialsV1(requestDiscoveryCreateHTTPWriteCredentialsV1)
}

// Alias Function
func (s *DiscoveryService) GetDiscoveryByID(id string) (*ResponseDiscoveryGetDiscoveryByIDV1, *resty.Response, error) {
	return s.GetDiscoveryByIDV1(id)
}

// Alias Function
func (s *DiscoveryService) CreateUpdateSNMPProperties(requestDiscoveryCreateUpdateSNMPPropertiesV1 *RequestDiscoveryCreateUpdateSNMPPropertiesV1) (*ResponseDiscoveryCreateUpdateSNMPPropertiesV1, *resty.Response, error) {
	return s.CreateUpdateSNMPPropertiesV1(requestDiscoveryCreateUpdateSNMPPropertiesV1)
}

// Alias Function
func (s *DiscoveryService) CreateSNMPv3Credentials(requestDiscoveryCreateSNMPv3CredentialsV1 *RequestDiscoveryCreateSNMPv3CredentialsV1) (*ResponseDiscoveryCreateSNMPv3CredentialsV1, *resty.Response, error) {
	return s.CreateSNMPv3CredentialsV1(requestDiscoveryCreateSNMPv3CredentialsV1)
}

// Alias Function
func (s *DiscoveryService) CreateNetconfCredentials(requestDiscoveryCreateNetconfCredentialsV1 *RequestDiscoveryCreateNetconfCredentialsV1) (*ResponseDiscoveryCreateNetconfCredentialsV1, *resty.Response, error) {
	return s.CreateNetconfCredentialsV1(requestDiscoveryCreateNetconfCredentialsV1)
}

// Alias Function
func (s *DiscoveryService) UpdateSNMPWriteCommunity(requestDiscoveryUpdateSNMPWriteCommunityV1 *RequestDiscoveryUpdateSNMPWriteCommunityV1) (*ResponseDiscoveryUpdateSNMPWriteCommunityV1, *resty.Response, error) {
	return s.UpdateSNMPWriteCommunityV1(requestDiscoveryUpdateSNMPWriteCommunityV1)
}

// Alias Function
func (s *DiscoveryService) GetDiscoveryJobsByIP(GetDiscoveryJobsByIPV1QueryParams *GetDiscoveryJobsByIPV1QueryParams) (*ResponseDiscoveryGetDiscoveryJobsByIPV1, *resty.Response, error) {
	return s.GetDiscoveryJobsByIPV1(GetDiscoveryJobsByIPV1QueryParams)
}
