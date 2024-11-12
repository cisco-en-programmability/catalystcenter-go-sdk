package catalyst

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type LanAutomationService service

type LanAutomationLogV1QueryParams struct {
	Offset float64 `url:"offset,omitempty"` //Starting index of the LAN Automation session. Minimum value is 1.
	Limit  float64 `url:"limit,omitempty"`  //Number of LAN Automation sessions to be retrieved. Limit value can range between 1 to 10.
}
type LanAutomationLogsForIndividualDevicesV1QueryParams struct {
	LogLevel string `url:"logLevel,omitempty"` //Supported levels are ERROR, INFO, WARNING, TRACE, CONFIG and ALL. Specifying ALL will display device specific logs with the exception of CONFIG logs. In order to view CONFIG logs along with the remaining logs, please leave the query parameter blank.
}
type LanAutomationStatusV1QueryParams struct {
	Offset float64 `url:"offset,omitempty"` //Starting index of the LAN Automation session. Minimum value is 1.
	Limit  float64 `url:"limit,omitempty"`  //Number of LAN Automation sessions to be retrieved. Limit value can range between 1 to 10.
}
type LanAutomationDeviceUpdateV1QueryParams struct {
	Feature string `url:"feature,omitempty"` //Feature ID for the update. Supported feature IDs include: LOOPBACK0_IPADDRESS_UPDATE, HOSTNAME_UPDATE, LINK_ADD, and LINK_DELETE.
}

type ResponseLanAutomationLanAutomationStartV1 struct {
	Response *ResponseLanAutomationLanAutomationStartV1Response `json:"response,omitempty"` //
	Version  string                                             `json:"version,omitempty"`  // Version
}
type ResponseLanAutomationLanAutomationStartV1Response struct {
	Message string `json:"message,omitempty"` // Status of the LAN Automation Start request
	ID      string `json:"id,omitempty"`      // LAN Automation Session Id
}
type ResponseLanAutomationLanAutomationSessionCountV1 struct {
	Response *ResponseLanAutomationLanAutomationSessionCountV1Response `json:"response,omitempty"` //
}
type ResponseLanAutomationLanAutomationSessionCountV1Response struct {
	SessionCount string `json:"sessionCount,omitempty"` // Total number of sessions executed.
}
type ResponseLanAutomationLanAutomationLogV1 struct {
	Response *[]ResponseLanAutomationLanAutomationLogV1Response `json:"response,omitempty"` //
	Version  string                                             `json:"version,omitempty"`  // Version
}
type ResponseLanAutomationLanAutomationLogV1Response struct {
	NwOrchID string                                                  `json:"nwOrchId,omitempty"` // LAN Automation session identifier.
	Entry    *[]ResponseLanAutomationLanAutomationLogV1ResponseEntry `json:"entry,omitempty"`    //
}
type ResponseLanAutomationLanAutomationLogV1ResponseEntry struct {
	LogLevel  string `json:"logLevel,omitempty"`  // Supported levels are ERROR, INFO, WARNING, TRACE and CONFIG.
	TimeStamp string `json:"timeStamp,omitempty"` // Time at which the log message is created.
	Record    string `json:"record,omitempty"`    // Detailed log message.
	DeviceID  string `json:"deviceId,omitempty"`  // Device serial number for which the log message is associated.
}
type ResponseLanAutomationLanAutomationLogByIDV1 struct {
	Response *[]ResponseLanAutomationLanAutomationLogByIDV1Response `json:"response,omitempty"` //
	Version  string                                                 `json:"version,omitempty"`  // Version
}
type ResponseLanAutomationLanAutomationLogByIDV1Response struct {
	NwOrchID string                                                      `json:"nwOrchId,omitempty"` // LAN Automation session identifier.
	Entry    *[]ResponseLanAutomationLanAutomationLogByIDV1ResponseEntry `json:"entry,omitempty"`    //
}
type ResponseLanAutomationLanAutomationLogByIDV1ResponseEntry struct {
	LogLevel  string `json:"logLevel,omitempty"`  // Supported levels are ERROR, INFO, WARNING, TRACE and CONFIG.
	TimeStamp string `json:"timeStamp,omitempty"` // Time at which the log message is created.
	Record    string `json:"record,omitempty"`    // Detailed log message.
	DeviceID  string `json:"deviceId,omitempty"`  // Device serial number for which the log message is associated.
}
type ResponseLanAutomationLanAutomationLogsForIndividualDevicesV1 struct {
	Response *[]ResponseLanAutomationLanAutomationLogsForIndividualDevicesV1Response `json:"response,omitempty"` //
	Version  string                                                                  `json:"version,omitempty"`  // Version
}
type ResponseLanAutomationLanAutomationLogsForIndividualDevicesV1Response struct {
	NwOrchID     string                                                                      `json:"nwOrchId,omitempty"`     // LAN Automation session identifier.
	Logs         *[]ResponseLanAutomationLanAutomationLogsForIndividualDevicesV1ResponseLogs `json:"logs,omitempty"`         //
	SerialNumber string                                                                      `json:"serialNumber,omitempty"` // Device serial number for which the log messages are associated.
}
type ResponseLanAutomationLanAutomationLogsForIndividualDevicesV1ResponseLogs struct {
	LogLevel  string `json:"logLevel,omitempty"`  // Supported levels are ERROR, INFO, WARNING, TRACE, CONFIG and ALL. Specifying ALL will display device specific logs with the exception of CONFIG logs. In order to view CONFIG logs along with the remaining logs, please leave the query parameter blank.
	TimeStamp string `json:"timeStamp,omitempty"` // Time at which the log message is created.
	Record    string `json:"record,omitempty"`    // Detailed log message.
}
type ResponseLanAutomationLanAutomationActiveSessionsV1 struct {
	Response *ResponseLanAutomationLanAutomationActiveSessionsV1Response `json:"response,omitempty"` //
	Version  string                                                      `json:"version,omitempty"`  // Version
}
type ResponseLanAutomationLanAutomationActiveSessionsV1Response struct {
	MaxSupportedCount string   `json:"maxSupportedCount,omitempty"` // Maximum supported parallel sessions count
	ActiveSessions    string   `json:"activeSessions,omitempty"`    // Current active sessions count
	ActiveSessionIDs  []string `json:"activeSessionIds,omitempty"`  // List of Active LAN Automation IDs
}
type ResponseLanAutomationLanAutomationStatusV1 struct {
	Response *[]ResponseLanAutomationLanAutomationStatusV1Response `json:"response,omitempty"` //
	Version  string                                                `json:"version,omitempty"`  // Version
}
type ResponseLanAutomationLanAutomationStatusV1Response struct {
	ID                                string                                                                    `json:"id,omitempty"`                                // LAN Automation session id.
	DiscoveredDeviceSiteNameHierarchy string                                                                    `json:"discoveredDeviceSiteNameHierarchy,omitempty"` // Discovered device site name.
	PrimaryDeviceManagmentIPAddress   string                                                                    `json:"primaryDeviceManagmentIPAddress,omitempty"`   // Primary seed device management IP address.
	IPPools                           *[]ResponseLanAutomationLanAutomationStatusV1ResponseIPPools              `json:"ipPools,omitempty"`                           //
	PrimaryDeviceInterfaceNames       []string                                                                  `json:"primaryDeviceInterfaceNames,omitempty"`       // The list of interfaces on primary seed via which the discovered devices are connected.
	Status                            string                                                                    `json:"status,omitempty"`                            // Status of the LAN Automation session along with the number of discovered devices.
	Action                            string                                                                    `json:"action,omitempty"`                            // State (START/STOP) of the LAN Automation session.
	CreationTime                      string                                                                    `json:"creationTime,omitempty"`                      // LAN Automation session creation time.
	MulticastEnabled                  *bool                                                                     `json:"multicastEnabled,omitempty"`                  // Shows whether underlay multicast is enabled or not.
	PeerDeviceManagmentIPAddress      string                                                                    `json:"peerDeviceManagmentIPAddress,omitempty"`      // Peer seed device management IP address.
	DiscoveredDeviceList              *[]ResponseLanAutomationLanAutomationStatusV1ResponseDiscoveredDeviceList `json:"discoveredDeviceList,omitempty"`              //
	RedistributeIsisToBgp             *bool                                                                     `json:"redistributeIsisToBgp,omitempty"`             // Shows whether advertise LAN Automation summary route into BGP is enabled or not.
	DiscoveryLevel                    *int                                                                      `json:"discoveryLevel,omitempty"`                    // Level below primary seed device upto which the new devices will be LAN Automated by this session, level + seed = tier. Supported range for level is [1-5], default level is 2.
	DiscoveryTimeout                  *int                                                                      `json:"discoveryTimeout,omitempty"`                  // Discovery timeout in minutes. Until this time, the stop processing will not be triggered. Any device contacting after the provided discovery timeout will not be processed, and a device reset and reload will be attempted to bring it back to the PnP agent state before process completion. The supported timeout range is in minutes [20-10080].
	DiscoveryDevices                  *[]ResponseLanAutomationLanAutomationStatusV1ResponseDiscoveryDevices     `json:"discoveryDevices,omitempty"`                  // Specific devices that will be LAN Automated in this session. Any other device discovered via DHCP will be attempted for a reset and reload to bring it back to the PnP agent state at the end of the LAN Automation process before process completion. The maximum supported devices that can be provided for a session is 50.
}
type ResponseLanAutomationLanAutomationStatusV1ResponseIPPools struct {
	IPPoolName string `json:"ipPoolName,omitempty"` // Name of the IP pool.
	IPPoolRole string `json:"ipPoolRole,omitempty"` // Role of the IP pool. Supported roles are: MAIN_POOL and PHYSICAL_LINK_POOL.
}
type ResponseLanAutomationLanAutomationStatusV1ResponseDiscoveredDeviceList struct {
	Name               string   `json:"name,omitempty"`               // Name of the device.
	SerialNumber       string   `json:"serialNumber,omitempty"`       // Serial number of the device.
	State              string   `json:"state,omitempty"`              // State of the device (Added to inventory/Deleted from inventory).
	IPAddressInUseList []string `json:"ipAddressInUseList,omitempty"` // List of IP address used by the device.
}
type ResponseLanAutomationLanAutomationStatusV1ResponseDiscoveryDevices interface{}
type ResponseLanAutomationLanAutomationStatusByIDV1 struct {
	Response *[]ResponseLanAutomationLanAutomationStatusByIDV1Response `json:"response,omitempty"` //
	Version  string                                                    `json:"version,omitempty"`  // Version
}
type ResponseLanAutomationLanAutomationStatusByIDV1Response struct {
	ID                                string                                                                        `json:"id,omitempty"`                                // LAN Automation session id.
	DiscoveredDeviceSiteNameHierarchy string                                                                        `json:"discoveredDeviceSiteNameHierarchy,omitempty"` // Discovered device site name.
	PrimaryDeviceManagmentIPAddress   string                                                                        `json:"primaryDeviceManagmentIPAddress,omitempty"`   // Primary seed device management IP address.
	IPPools                           *[]ResponseLanAutomationLanAutomationStatusByIDV1ResponseIPPools              `json:"ipPools,omitempty"`                           //
	PrimaryDeviceInterfaceNames       []string                                                                      `json:"primaryDeviceInterfaceNames,omitempty"`       // The list of interfaces on primary seed via which the discovered devices are connected.
	Status                            string                                                                        `json:"status,omitempty"`                            // Status of the LAN Automation session along with the number of discovered devices.
	Action                            string                                                                        `json:"action,omitempty"`                            // State (START/STOP) of the LAN Automation session.
	CreationTime                      string                                                                        `json:"creationTime,omitempty"`                      // LAN Automation session creation time.
	MulticastEnabled                  *bool                                                                         `json:"multicastEnabled,omitempty"`                  // Shows whether underlay multicast is enabled or not.
	PeerDeviceManagmentIPAddress      string                                                                        `json:"peerDeviceManagmentIPAddress,omitempty"`      // Peer seed device management IP address.
	DiscoveredDeviceList              *[]ResponseLanAutomationLanAutomationStatusByIDV1ResponseDiscoveredDeviceList `json:"discoveredDeviceList,omitempty"`              //
	RedistributeIsisToBgp             *bool                                                                         `json:"redistributeIsisToBgp,omitempty"`             // Shows whether advertise LAN Automation summary route into BGP is enabled or not.
	DiscoveryLevel                    *int                                                                          `json:"discoveryLevel,omitempty"`                    // Level below primary seed device upto which the new devices will be LAN Automated by this session, level + seed = tier. Supported range for level is [1-5], default level is 2.
	DiscoveryTimeout                  *int                                                                          `json:"discoveryTimeout,omitempty"`                  // Discovery timeout in minutes. Until this time, the stop processing will not be triggered. Any device contacting after the provided discovery timeout will not be processed, and a device reset and reload will be attempted to bring it back to the PnP agent state before process completion. The supported timeout range is in minutes [20-10080].
	DiscoveryDevices                  *[]ResponseLanAutomationLanAutomationStatusByIDV1ResponseDiscoveryDevices     `json:"discoveryDevices,omitempty"`                  // Specific devices that will be LAN Automated in this session. Any other device discovered via DHCP will be attempted for a reset and reload to bring it back to the PnP agent state at the end of the LAN Automation process before process completion. The maximum supported devices that can be provided for a session is 50.
}
type ResponseLanAutomationLanAutomationStatusByIDV1ResponseIPPools struct {
	IPPoolName string `json:"ipPoolName,omitempty"` // Name of the IP pool.
	IPPoolRole string `json:"ipPoolRole,omitempty"` // Role of the IP pool. Supported roles are: MAIN_POOL and PHYSICAL_LINK_POOL.
}
type ResponseLanAutomationLanAutomationStatusByIDV1ResponseDiscoveredDeviceList struct {
	Name               string   `json:"name,omitempty"`               // Name of the device.
	SerialNumber       string   `json:"serialNumber,omitempty"`       // Serial number of the device.
	State              string   `json:"state,omitempty"`              // State of the device (Added to inventory/Deleted from inventory).
	IPAddressInUseList []string `json:"ipAddressInUseList,omitempty"` // List of IP address used by the device.
}
type ResponseLanAutomationLanAutomationStatusByIDV1ResponseDiscoveryDevices interface{}
type ResponseLanAutomationLanAutomationDeviceUpdateV1 struct {
	Response *ResponseLanAutomationLanAutomationDeviceUpdateV1Response `json:"response,omitempty"` //
	Version  string                                                    `json:"version,omitempty"`  // version
}
type ResponseLanAutomationLanAutomationDeviceUpdateV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID
	URL    string `json:"url,omitempty"`    // url to check the status of task
}
type ResponseLanAutomationLanAutomationStopV1 struct {
	Response *ResponseLanAutomationLanAutomationStopV1Response `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  // Version
}
type ResponseLanAutomationLanAutomationStopV1Response struct {
	ErrorCode string `json:"errorCode,omitempty"` // Error code value.
	Message   string `json:"message,omitempty"`   // Description of the error code.
	Detail    string `json:"detail,omitempty"`    // Detailed information of the error code.
}
type ResponseLanAutomationLanAutomationStopAndUpdateDevicesV1 struct {
	Response *ResponseLanAutomationLanAutomationStopAndUpdateDevicesV1Response `json:"response,omitempty"` //
	Version  string                                                            `json:"version,omitempty"`  // version
}
type ResponseLanAutomationLanAutomationStopAndUpdateDevicesV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID
	URL    string `json:"url,omitempty"`    // url to check the status of task
}
type ResponseLanAutomationLanAutomationStartV2 struct {
	Response *ResponseLanAutomationLanAutomationStartV2Response `json:"response,omitempty"` //
	Version  string                                             `json:"version,omitempty"`  // version
}
type ResponseLanAutomationLanAutomationStartV2Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID
	URL    string `json:"url,omitempty"`    // url to check the status of task
}
type ResponseLanAutomationLanAutomationStopAndUpdateDevicesV2 struct {
	Response *ResponseLanAutomationLanAutomationStopAndUpdateDevicesV2Response `json:"response,omitempty"` //
	Version  string                                                            `json:"version,omitempty"`  // version
}
type ResponseLanAutomationLanAutomationStopAndUpdateDevicesV2Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID
	URL    string `json:"url,omitempty"`    // url to check the status of task
}
type RequestLanAutomationLanAutomationStartV1 []RequestItemLanAutomationLanAutomationStartV1 // Array of RequestLanAutomationLANAutomationStartV1
type RequestItemLanAutomationLanAutomationStartV1 struct {
	DiscoveredDeviceSiteNameHierarchy string                                                 `json:"discoveredDeviceSiteNameHierarchy,omitempty"` // Discovered device site name.
	PrimaryDeviceManagmentIPAddress   string                                                 `json:"primaryDeviceManagmentIPAddress,omitempty"`   // Primary seed management IP address.
	PeerDeviceManagmentIPAddress      string                                                 `json:"peerDeviceManagmentIPAddress,omitempty"`      // Peer seed management IP address.
	PrimaryDeviceInterfaceNames       []string                                               `json:"primaryDeviceInterfaceNames,omitempty"`       // The list of interfaces on primary seed via which the discovered devices are connected.
	IPPools                           *[]RequestItemLanAutomationLanAutomationStartV1IPPools `json:"ipPools,omitempty"`                           //
	MulitcastEnabled                  *bool                                                  `json:"mulitcastEnabled,omitempty"`                  // To enable underlay native multicast.
	HostNamePrefix                    string                                                 `json:"hostNamePrefix,omitempty"`                    // Host name prefix which shall be assigned to the discovered device.
	HostNameFileID                    string                                                 `json:"hostNameFileId,omitempty"`                    // Use /dna/intent/api/v1/file/namespace/nw_orch api to get the file id for the already uploaded file in nw_orch namespace.
	IsisDomainPwd                     string                                                 `json:"isisDomainPwd,omitempty"`                     // IS-IS domain password in plain text.
	RedistributeIsisToBgp             *bool                                                  `json:"redistributeIsisToBgp,omitempty"`             // Advertise LAN Automation summary route into BGP.
}
type RequestItemLanAutomationLanAutomationStartV1IPPools struct {
	IPPoolName string `json:"ipPoolName,omitempty"` // Name of the IP pool.
	IPPoolRole string `json:"ipPoolRole,omitempty"` // Role of the IP pool. Supported roles are: MAIN_POOL and PHYSICAL_LINK_POOL.
}
type RequestLanAutomationLanAutomationDeviceUpdateV1 struct {
	LoopbackUpdateDeviceList *[]RequestLanAutomationLanAutomationDeviceUpdateV1LoopbackUpdateDeviceList `json:"loopbackUpdateDeviceList,omitempty"` //
	LinkUpdate               *RequestLanAutomationLanAutomationDeviceUpdateV1LinkUpdate                 `json:"linkUpdate,omitempty"`               //
	HostnameUpdateDevices    *[]RequestLanAutomationLanAutomationDeviceUpdateV1HostnameUpdateDevices    `json:"hostnameUpdateDevices,omitempty"`    //
}
type RequestLanAutomationLanAutomationDeviceUpdateV1LoopbackUpdateDeviceList struct {
	DeviceManagementIPAddress string `json:"deviceManagementIPAddress,omitempty"` // Device Management IP Address
	NewLoopback0IPAddress     string `json:"newLoopback0IPAddress,omitempty"`     // New Loopback0 IP Address from LAN Pool of Device Discovery Site(Shared pool should not be used).
}
type RequestLanAutomationLanAutomationDeviceUpdateV1LinkUpdate struct {
	SourceDeviceManagementIPAddress      string `json:"sourceDeviceManagementIPAddress,omitempty"`      // Source Device Management IP Address
	SourceDeviceInterfaceName            string `json:"sourceDeviceInterfaceName,omitempty"`            // Source Device Interface Name
	DestinationDeviceManagementIPAddress string `json:"destinationDeviceManagementIPAddress,omitempty"` // Destination Device Management IP Address
	DestinationDeviceInterfaceName       string `json:"destinationDeviceInterfaceName,omitempty"`       // Destination Device Interface Name
	IPPoolName                           string `json:"ipPoolName,omitempty"`                           // Name of the IP LAN Pool, required for Link Add should be from discovery site of source and destination device.
}
type RequestLanAutomationLanAutomationDeviceUpdateV1HostnameUpdateDevices struct {
	DeviceManagementIPAddress string `json:"deviceManagementIPAddress,omitempty"` // Device Management IP Address
	NewHostName               string `json:"newHostName,omitempty"`               // New hostname for the device
}
type RequestLanAutomationLanAutomationStopAndUpdateDevicesV1 []RequestItemLanAutomationLanAutomationStopAndUpdateDevicesV1 // Array of RequestLanAutomationLANAutomationStopAndUpdateDevicesV1
type RequestItemLanAutomationLanAutomationStopAndUpdateDevicesV1 struct {
	DeviceManagementIPAddress string `json:"deviceManagementIPAddress,omitempty"` // Device Management IP Address
	NewLoopback0IPAddress     string `json:"newLoopback0IPAddress,omitempty"`     // New Loopback0 IP Address from LAN pool of Device Discovery Site.
}
type RequestLanAutomationLanAutomationStartV2 []RequestItemLanAutomationLanAutomationStartV2 // Array of RequestLanAutomationLANAutomationStartV2
type RequestItemLanAutomationLanAutomationStartV2 struct {
	DiscoveredDeviceSiteNameHierarchy string                                                          `json:"discoveredDeviceSiteNameHierarchy,omitempty"` // Discovered device site name.
	PrimaryDeviceManagmentIPAddress   string                                                          `json:"primaryDeviceManagmentIPAddress,omitempty"`   // Primary seed management IP address.
	PeerDeviceManagmentIPAddress      string                                                          `json:"peerDeviceManagmentIPAddress,omitempty"`      // Peer seed management IP address.
	PrimaryDeviceInterfaceNames       []string                                                        `json:"primaryDeviceInterfaceNames,omitempty"`       // The list of interfaces on primary seed via which the discovered devices are connected.
	IPPools                           *[]RequestItemLanAutomationLanAutomationStartV2IPPools          `json:"ipPools,omitempty"`                           //
	MulticastEnabled                  *bool                                                           `json:"multicastEnabled,omitempty"`                  // Enable underlay native multicast.
	HostNamePrefix                    string                                                          `json:"hostNamePrefix,omitempty"`                    // Host name prefix assigned to the discovered device.
	HostNameFileID                    string                                                          `json:"hostNameFileId,omitempty"`                    // Use /dna/intent/api/v1/file/namespace/nw_orch API to get the file ID for the already uploaded file in the nw_orch namespace.
	RedistributeIsisToBgp             *bool                                                           `json:"redistributeIsisToBgp,omitempty"`             // Advertise LAN Automation summary route into BGP.
	IsisDomainPwd                     string                                                          `json:"isisDomainPwd,omitempty"`                     // IS-IS domain password in plain text.
	DiscoveryLevel                    *int                                                            `json:"discoveryLevel,omitempty"`                    // Level below primary seed device upto which the new devices will be LAN Automated by this session, level + seed = tier. Supported range for level is [1-5], default level is 2.
	DiscoveryTimeout                  *int                                                            `json:"discoveryTimeout,omitempty"`                  // Discovery timeout in minutes. Until this time, the stop processing will not be triggered. Any device contacting after the provided discovery timeout will not be processed, and a device reset and reload will be attempted to bring it back to the PnP agent state before process completion. The supported timeout range is in minutes [20-10080]. If both timeout and discovery devices list are provided, the stop processing will be attempted whichever happens earlier. Users can always use the LAN Automation delete API to force stop processing.
	DiscoveryDevices                  *[]RequestItemLanAutomationLanAutomationStartV2DiscoveryDevices `json:"discoveryDevices,omitempty"`                  //
}
type RequestItemLanAutomationLanAutomationStartV2IPPools struct {
	IPPoolName string `json:"ipPoolName,omitempty"` // Name of the IP pool.
	IPPoolRole string `json:"ipPoolRole,omitempty"` // Role of the IP pool. Supported roles are: MAIN_POOL and PHYSICAL_LINK_POOL.
}
type RequestItemLanAutomationLanAutomationStartV2DiscoveryDevices struct {
	DeviceSerialNumber        string `json:"deviceSerialNumber,omitempty"`        // Serial number of the device
	DeviceHostName            string `json:"deviceHostName,omitempty"`            // Hostname of the device
	DeviceSiteNameHierarchy   string `json:"deviceSiteNameHierarchy,omitempty"`   // Site name hierarchy for the device, must be a child site of the discoveredDeviceSiteNameHierarchy or same if it’s not area type.
	DeviceManagementIPAddress string `json:"deviceManagementIPAddress,omitempty"` // Management IP Address of the device
}
type RequestLanAutomationLanAutomationStopAndUpdateDevicesV2 []RequestItemLanAutomationLanAutomationStopAndUpdateDevicesV2 // Array of RequestLanAutomationLANAutomationStopAndUpdateDevicesV2
type RequestItemLanAutomationLanAutomationStopAndUpdateDevicesV2 struct {
	DeviceManagementIPAddress string `json:"deviceManagementIPAddress,omitempty"` // Device Management IP Address
	NewLoopback0IPAddress     string `json:"newLoopback0IPAddress,omitempty"`     // New Loopback0 IP Address from LAN pool of Device Discovery Site.
	NewHostName               string `json:"newHostName,omitempty"`               // New hostname to be assigned to the device
}

//LanAutomationSessionCountV1 LAN Automation Session Count - b08b-6b11-4669-a12b
/* Invoke this API to get the total count of LAN Automation sessions.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!lan-automation-session-count-v1
*/
func (s *LanAutomationService) LanAutomationSessionCountV1() (*ResponseLanAutomationLanAutomationSessionCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/lan-automation/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLanAutomationLanAutomationSessionCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.LanAutomationSessionCountV1()
		}
		return nil, response, fmt.Errorf("error with operation LanAutomationSessionCountV1")
	}

	result := response.Result().(*ResponseLanAutomationLanAutomationSessionCountV1)
	return result, response, err

}

//LanAutomationLogV1 LAN Automation Log  - 93a9-68c2-480a-85d1
/* Invoke this API to get the LAN Automation session logs.


@param LANAutomationLogV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!lan-automation-log-v1
*/
func (s *LanAutomationService) LanAutomationLogV1(LANAutomationLogV1QueryParams *LanAutomationLogV1QueryParams) (*ResponseLanAutomationLanAutomationLogV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/lan-automation/log"

	queryString, _ := query.Values(LANAutomationLogV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseLanAutomationLanAutomationLogV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.LanAutomationLogV1(LANAutomationLogV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation LanAutomationLogV1")
	}

	result := response.Result().(*ResponseLanAutomationLanAutomationLogV1)
	return result, response, err

}

//LanAutomationLogByIDV1 LAN Automation Log by Id - 55b5-eb50-440a-a123
/* Invoke this API to get the LAN Automation session logs based on the given LAN Automation session id.


@param id id path parameter. LAN Automation session identifier.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!lan-automation-log-by-id-v1
*/
func (s *LanAutomationService) LanAutomationLogByIDV1(id string) (*ResponseLanAutomationLanAutomationLogByIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/lan-automation/log/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLanAutomationLanAutomationLogByIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.LanAutomationLogByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation LanAutomationLogByIdV1")
	}

	result := response.Result().(*ResponseLanAutomationLanAutomationLogByIDV1)
	return result, response, err

}

//LanAutomationLogsForIndividualDevicesV1 LAN Automation Logs for Individual Devices - b2ac-5af7-45d8-8c4e
/* Invoke this API to get the LAN Automation session logs for individual devices based on the given LAN Automation session id and device serial number.


@param id id path parameter. LAN Automation session identifier.

@param serialNumber serialNumber path parameter. Device serial number.

@param LANAutomationLogsForIndividualDevicesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!lan-automation-logs-for-individual-devices-v1
*/
func (s *LanAutomationService) LanAutomationLogsForIndividualDevicesV1(id string, serialNumber string, LANAutomationLogsForIndividualDevicesV1QueryParams *LanAutomationLogsForIndividualDevicesV1QueryParams) (*ResponseLanAutomationLanAutomationLogsForIndividualDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/lan-automation/log/{id}/{serialNumber}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{serialNumber}", fmt.Sprintf("%v", serialNumber), -1)

	queryString, _ := query.Values(LANAutomationLogsForIndividualDevicesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseLanAutomationLanAutomationLogsForIndividualDevicesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.LanAutomationLogsForIndividualDevicesV1(id, serialNumber, LANAutomationLogsForIndividualDevicesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation LanAutomationLogsForIndividualDevicesV1")
	}

	result := response.Result().(*ResponseLanAutomationLanAutomationLogsForIndividualDevicesV1)
	return result, response, err

}

//LanAutomationActiveSessionsV1 LAN Automation Active Sessions - c1bf-69fb-4ad8-979c
/* Invoke this API to get the LAN Automation active session information



Documentation Link: https://developer.cisco.com/docs/dna-center/#!lan-automation-active-sessions-v1
*/
func (s *LanAutomationService) LanAutomationActiveSessionsV1() (*ResponseLanAutomationLanAutomationActiveSessionsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/lan-automation/sessions"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLanAutomationLanAutomationActiveSessionsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.LanAutomationActiveSessionsV1()
		}
		return nil, response, fmt.Errorf("error with operation LanAutomationActiveSessionsV1")
	}

	result := response.Result().(*ResponseLanAutomationLanAutomationActiveSessionsV1)
	return result, response, err

}

//LanAutomationStatusV1 LAN Automation Status - a4ab-087e-4ed9-a3bb
/* Invoke this API to get the LAN Automation session status.


@param LANAutomationStatusV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!lan-automation-status-v1
*/
func (s *LanAutomationService) LanAutomationStatusV1(LANAutomationStatusV1QueryParams *LanAutomationStatusV1QueryParams) (*ResponseLanAutomationLanAutomationStatusV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/lan-automation/status"

	queryString, _ := query.Values(LANAutomationStatusV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseLanAutomationLanAutomationStatusV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.LanAutomationStatusV1(LANAutomationStatusV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation LanAutomationStatusV1")
	}

	result := response.Result().(*ResponseLanAutomationLanAutomationStatusV1)
	return result, response, err

}

//LanAutomationStatusByIDV1 LAN Automation Status by Id - 5b99-8b6e-47b8-9882
/* Invoke this API to get the LAN Automation session status based on the given Lan Automation session id.


@param id id path parameter. LAN Automation session identifier.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!lan-automation-status-by-id-v1
*/
func (s *LanAutomationService) LanAutomationStatusByIDV1(id string) (*ResponseLanAutomationLanAutomationStatusByIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/lan-automation/status/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLanAutomationLanAutomationStatusByIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.LanAutomationStatusByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation LanAutomationStatusByIdV1")
	}

	result := response.Result().(*ResponseLanAutomationLanAutomationStatusByIDV1)
	return result, response, err

}

//LanAutomationStartV1 LAN Automation Start - 9795-f927-469a-a6d2
/* Invoke this API to start LAN Automation for the given site.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!lan-automation-start-v1
*/
func (s *LanAutomationService) LanAutomationStartV1(requestLanAutomationLANAutomationStartV1 *RequestLanAutomationLanAutomationStartV1) (*ResponseLanAutomationLanAutomationStartV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/lan-automation"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestLanAutomationLANAutomationStartV1).
		SetResult(&ResponseLanAutomationLanAutomationStartV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.LanAutomationStartV1(requestLanAutomationLANAutomationStartV1)
		}

		return nil, response, fmt.Errorf("error with operation LanAutomationStartV1")
	}

	result := response.Result().(*ResponseLanAutomationLanAutomationStartV1)
	return result, response, err

}

//LanAutomationStartV2 LAN Automation Start V2 - 51ba-8921-46da-9bec
/* Invoke V2 LAN Automation Start API, which supports optional auto-stop processing feature based on the provided timeout or a specific device list, or both. The stop processing will be executed automatically when either of the cases is satisfied, without specifically calling the stop API. The V2 API behaves similarly to V1 if no timeout or device list is provided, and the user needs to call the stop API for LAN Automation stop processing. With the V2 API, the user can also specify the level up to which the devices can be LAN automated.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!lan-automation-start-v2
*/
func (s *LanAutomationService) LanAutomationStartV2(requestLanAutomationLANAutomationStartV2 *RequestLanAutomationLanAutomationStartV2) (*ResponseLanAutomationLanAutomationStartV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/lan-automation"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestLanAutomationLANAutomationStartV2).
		SetResult(&ResponseLanAutomationLanAutomationStartV2{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.LanAutomationStartV2(requestLanAutomationLANAutomationStartV2)
		}

		return nil, response, fmt.Errorf("error with operation LanAutomationStartV2")
	}

	result := response.Result().(*ResponseLanAutomationLanAutomationStartV2)
	return result, response, err

}

//LanAutomationDeviceUpdateV1 LAN Automation Device Update - 1190-5ac3-4e88-bd5e
/* Invoke this API to perform a DAY-N update on LAN Automation-related devices. Supported features include Loopback0 IP update, hostname update, link addition, and link deletion.


@param LANAutomationDeviceUpdateV1QueryParams Filtering parameter
*/
func (s *LanAutomationService) LanAutomationDeviceUpdateV1(requestLanAutomationLANAutomationDeviceUpdateV1 *RequestLanAutomationLanAutomationDeviceUpdateV1, LANAutomationDeviceUpdateV1QueryParams *LanAutomationDeviceUpdateV1QueryParams) (*ResponseLanAutomationLanAutomationDeviceUpdateV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/lan-automation/updateDevice"

	queryString, _ := query.Values(LANAutomationDeviceUpdateV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetBody(requestLanAutomationLANAutomationDeviceUpdateV1).
		SetResult(&ResponseLanAutomationLanAutomationDeviceUpdateV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.LanAutomationDeviceUpdateV1(requestLanAutomationLANAutomationDeviceUpdateV1, LANAutomationDeviceUpdateV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation LanAutomationDeviceUpdateV1")
	}

	result := response.Result().(*ResponseLanAutomationLanAutomationDeviceUpdateV1)
	return result, response, err

}

//LanAutomationStopAndUpdateDevicesV1 LAN Automation Stop and Update Devices - 0780-4a4c-44cb-bae8
/* Invoke this API to stop LAN Automation and Update Loopback0 IP Address of Devices, discovered in the current session


@param id id path parameter. LAN Automation id can be obtained from /dna/intent/api/v1/lan-automation/status.

*/
func (s *LanAutomationService) LanAutomationStopAndUpdateDevicesV1(id string, requestLanAutomationLANAutomationStopAndUpdateDevicesV1 *RequestLanAutomationLanAutomationStopAndUpdateDevicesV1) (*ResponseLanAutomationLanAutomationStopAndUpdateDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/lan-automation/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestLanAutomationLANAutomationStopAndUpdateDevicesV1).
		SetResult(&ResponseLanAutomationLanAutomationStopAndUpdateDevicesV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.LanAutomationStopAndUpdateDevicesV1(id, requestLanAutomationLANAutomationStopAndUpdateDevicesV1)
		}
		return nil, response, fmt.Errorf("error with operation LanAutomationStopAndUpdateDevicesV1")
	}

	result := response.Result().(*ResponseLanAutomationLanAutomationStopAndUpdateDevicesV1)
	return result, response, err

}

//LanAutomationStopAndUpdateDevicesV2 LAN Automation Stop and Update Devices V2 - 9381-ba28-42c9-9e6a
/* Invoke this API to stop LAN Automation and update device parameters such as Loopback0 IP address and/or hostname discovered in the current session.


@param id id path parameter. LAN Automation id can be obtained from /dna/intent/api/v1/lan-automation/status.

*/
func (s *LanAutomationService) LanAutomationStopAndUpdateDevicesV2(id string, requestLanAutomationLANAutomationStopAndUpdateDevicesV2 *RequestLanAutomationLanAutomationStopAndUpdateDevicesV2) (*ResponseLanAutomationLanAutomationStopAndUpdateDevicesV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/lan-automation/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestLanAutomationLANAutomationStopAndUpdateDevicesV2).
		SetResult(&ResponseLanAutomationLanAutomationStopAndUpdateDevicesV2{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.LanAutomationStopAndUpdateDevicesV2(id, requestLanAutomationLANAutomationStopAndUpdateDevicesV2)
		}
		return nil, response, fmt.Errorf("error with operation LanAutomationStopAndUpdateDevicesV2")
	}

	result := response.Result().(*ResponseLanAutomationLanAutomationStopAndUpdateDevicesV2)
	return result, response, err

}

//LanAutomationStopV1 LAN Automation Stop - e6a0-da69-4adb-8929
/* Invoke this API to stop LAN Automation for the given site.


@param id id path parameter. LAN Automation id can be obtained from /dna/intent/api/v1/lan-automation/status.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!lan-automation-stop-v1
*/
func (s *LanAutomationService) LanAutomationStopV1(id string) (*ResponseLanAutomationLanAutomationStopV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/lan-automation/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLanAutomationLanAutomationStopV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.LanAutomationStopV1(
				id)
		}
		return nil, response, fmt.Errorf("error with operation LanAutomationStopV1")
	}

	result := response.Result().(*ResponseLanAutomationLanAutomationStopV1)
	return result, response, err

}

// Alias Function
/*
This method acts as an alias for the method `LanAutomationStatusByIDV1`
*/
func (s *LanAutomationService) LanAutomationStatusByID(id string) (*ResponseLanAutomationLanAutomationStatusByIDV1, *resty.Response, error) {
	return s.LanAutomationStatusByIDV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `LanAutomationStopV1`
*/
func (s *LanAutomationService) LanAutomationStop(id string) (*ResponseLanAutomationLanAutomationStopV1, *resty.Response, error) {
	return s.LanAutomationStopV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `LanAutomationLogV1`
*/
func (s *LanAutomationService) LanAutomationLog(LANAutomationLogV1QueryParams *LanAutomationLogV1QueryParams) (*ResponseLanAutomationLanAutomationLogV1, *resty.Response, error) {
	return s.LanAutomationLogV1(LANAutomationLogV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `LanAutomationLogByIDV1`
*/
func (s *LanAutomationService) LanAutomationLogByID(id string) (*ResponseLanAutomationLanAutomationLogByIDV1, *resty.Response, error) {
	return s.LanAutomationLogByIDV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `LanAutomationLogsForIndividualDevicesV1`
*/
func (s *LanAutomationService) LanAutomationLogsForIndividualDevices(id string, serialNumber string, LANAutomationLogsForIndividualDevicesV1QueryParams *LanAutomationLogsForIndividualDevicesV1QueryParams) (*ResponseLanAutomationLanAutomationLogsForIndividualDevicesV1, *resty.Response, error) {
	return s.LanAutomationLogsForIndividualDevicesV1(id, serialNumber, LANAutomationLogsForIndividualDevicesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `LanAutomationStopAndUpdateDevicesV1`
*/
func (s *LanAutomationService) LanAutomationStopAndUpdateDevices(id string, requestLanAutomationLANAutomationStopAndUpdateDevicesV1 *RequestLanAutomationLanAutomationStopAndUpdateDevicesV1) (*ResponseLanAutomationLanAutomationStopAndUpdateDevicesV1, *resty.Response, error) {
	return s.LanAutomationStopAndUpdateDevicesV1(id, requestLanAutomationLANAutomationStopAndUpdateDevicesV1)
}

// Alias Function
/*
This method acts as an alias for the method `LanAutomationStartV1`
*/
func (s *LanAutomationService) LanAutomationStart(requestLanAutomationLANAutomationStartV1 *RequestLanAutomationLanAutomationStartV1) (*ResponseLanAutomationLanAutomationStartV1, *resty.Response, error) {
	return s.LanAutomationStartV1(requestLanAutomationLANAutomationStartV1)
}

// Alias Function
/*
This method acts as an alias for the method `LanAutomationStatusV1`
*/
func (s *LanAutomationService) LanAutomationStatus(LANAutomationStatusV1QueryParams *LanAutomationStatusV1QueryParams) (*ResponseLanAutomationLanAutomationStatusV1, *resty.Response, error) {
	return s.LanAutomationStatusV1(LANAutomationStatusV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `LanAutomationDeviceUpdateV1`
*/
func (s *LanAutomationService) LanAutomationDeviceUpdate(requestLanAutomationLANAutomationDeviceUpdateV1 *RequestLanAutomationLanAutomationDeviceUpdateV1, LANAutomationDeviceUpdateV1QueryParams *LanAutomationDeviceUpdateV1QueryParams) (*ResponseLanAutomationLanAutomationDeviceUpdateV1, *resty.Response, error) {
	return s.LanAutomationDeviceUpdateV1(requestLanAutomationLANAutomationDeviceUpdateV1, LANAutomationDeviceUpdateV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `LanAutomationSessionCountV1`
*/
func (s *LanAutomationService) LanAutomationSessionCount() (*ResponseLanAutomationLanAutomationSessionCountV1, *resty.Response, error) {
	return s.LanAutomationSessionCountV1()
}

// Alias Function
/*
This method acts as an alias for the method `LanAutomationActiveSessionsV1`
*/
func (s *LanAutomationService) LanAutomationActiveSessions() (*ResponseLanAutomationLanAutomationActiveSessionsV1, *resty.Response, error) {
	return s.LanAutomationActiveSessionsV1()
}
