package catalyst

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type UsersService service

type GetUserEnrichmentDetailsV1HeaderParams struct {
	EntityType        string `url:"entity_type,omitempty"`         //Expects type string. User enrichment details can be fetched based on either User ID or Client MAC address. This parameter value must either be network_user_id/mac_address
	EntityValue       string `url:"entity_value,omitempty"`        //Expects type string. Contains the actual value for the entity type that has been defined
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type bool.
}

type ResponseUsersGetUserEnrichmentDetailsV1 []ResponseItemUsersGetUserEnrichmentDetailsV1 // Array of ResponseUsersGetUserEnrichmentDetailsV1
type ResponseItemUsersGetUserEnrichmentDetailsV1 struct {
	UserDetails     *ResponseItemUsersGetUserEnrichmentDetailsV1UserDetails       `json:"userDetails,omitempty"`     //
	ConnectedDevice *[]ResponseItemUsersGetUserEnrichmentDetailsV1ConnectedDevice `json:"connectedDevice,omitempty"` //
}
type ResponseItemUsersGetUserEnrichmentDetailsV1UserDetails struct {
	ID               string                                                                   `json:"id,omitempty"`               // Id
	ConnectionStatus string                                                                   `json:"connectionStatus,omitempty"` // Connection Status
	HostType         string                                                                   `json:"hostType,omitempty"`         // Host Type
	UserID           *ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsUserID            `json:"userId,omitempty"`           // User Id
	HostName         *ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsHostName          `json:"hostName,omitempty"`         // Host Name
	HostOs           *ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsHostOs            `json:"hostOs,omitempty"`           // Host Os
	HostVersion      *ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsHostVersion       `json:"hostVersion,omitempty"`      // Host Version
	SubType          string                                                                   `json:"subType,omitempty"`          // Sub Type
	LastUpdated      *int                                                                     `json:"lastUpdated,omitempty"`      // Last Updated
	HealthScore      *[]ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsHealthScore     `json:"healthScore,omitempty"`      //
	HostMac          string                                                                   `json:"hostMac,omitempty"`          // Host Mac
	HostIPV4         string                                                                   `json:"hostIpV4,omitempty"`         // Host Ip V4
	HostIPV6         *[]ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsHostIPV6        `json:"hostIpV6,omitempty"`         // Host Ip V6
	AuthType         *ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsAuthType          `json:"authType,omitempty"`         // Auth Type
	VLANID           string                                                                   `json:"vlanId,omitempty"`           // Vlan Id
	SSID             *ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsSSID              `json:"ssid,omitempty"`             // Ssid
	Frequency        *ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsFrequency         `json:"frequency,omitempty"`        // Frequency
	Channel          *ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsChannel           `json:"channel,omitempty"`          // Channel
	ApGroup          *ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsApGroup           `json:"apGroup,omitempty"`          // Ap Group
	Location         *ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsLocation          `json:"location,omitempty"`         // Location
	ClientConnection string                                                                   `json:"clientConnection,omitempty"` // Client Connection
	ConnectedDevice  *[]ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsConnectedDevice `json:"connectedDevice,omitempty"`  // Connected Device
	IssueCount       *float64                                                                 `json:"issueCount,omitempty"`       // Issue Count
	Rssi             *ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsRssi              `json:"rssi,omitempty"`             // Rssi
	AvgRssi          *ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsAvgRssi           `json:"avgRssi,omitempty"`          // Avg Rssi
	Snr              *ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsSnr               `json:"snr,omitempty"`              // Snr
	AvgSnr           *ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsAvgSnr            `json:"avgSnr,omitempty"`           // Avg Snr
	DataRate         *ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsDataRate          `json:"dataRate,omitempty"`         // Data Rate
	TxBytes          *ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsTxBytes           `json:"txBytes,omitempty"`          // Tx Bytes
	RxBytes          *ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsRxBytes           `json:"rxBytes,omitempty"`          // Rx Bytes
	DNSSuccess       *ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsDNSSuccess        `json:"dnsSuccess,omitempty"`       // Dns Success
	DNSFailure       *ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsDNSFailure        `json:"dnsFailure,omitempty"`       // Dns Failure
	Onboarding       *ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsOnboarding        `json:"onboarding,omitempty"`       //
	OnboardingTime   *ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsOnboardingTime    `json:"onboardingTime,omitempty"`   // Onboarding Time
	Port             *ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsPort              `json:"port,omitempty"`             // Port
}
type ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsUserID interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsHostName interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsHostOs interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsHostVersion interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsHealthScore struct {
	HealthType string `json:"healthType,omitempty"` // Health Type
	Reason     string `json:"reason,omitempty"`     // Reason
	Score      *int   `json:"score,omitempty"`      // Score
}
type ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsHostIPV6 interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsAuthType interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsSSID interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsFrequency interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsChannel interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsApGroup interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsLocation interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsConnectedDevice interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsRssi interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsAvgRssi interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsSnr interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsAvgSnr interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsDataRate interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsTxBytes interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsRxBytes interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsDNSSuccess interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsDNSFailure interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsOnboarding struct {
	AverageRunDuration   *ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsOnboardingAverageRunDuration   `json:"averageRunDuration,omitempty"`   // Average Run Duration
	MaxRunDuration       *ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsOnboardingMaxRunDuration       `json:"maxRunDuration,omitempty"`       // Max Run Duration
	AverageAssocDuration *ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsOnboardingAverageAssocDuration `json:"averageAssocDuration,omitempty"` // Average Assoc Duration
	MaxAssocDuration     *ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsOnboardingMaxAssocDuration     `json:"maxAssocDuration,omitempty"`     // Max Assoc Duration
	AverageAuthDuration  *ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsOnboardingAverageAuthDuration  `json:"averageAuthDuration,omitempty"`  // Average Auth Duration
	MaxAuthDuration      *ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsOnboardingMaxAuthDuration      `json:"maxAuthDuration,omitempty"`      // Max Auth Duration
	AverageDhcpDuration  *ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsOnboardingAverageDhcpDuration  `json:"averageDhcpDuration,omitempty"`  // Average Dhcp Duration
	MaxDhcpDuration      *ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsOnboardingMaxDhcpDuration      `json:"maxDhcpDuration,omitempty"`      // Max Dhcp Duration
	AAAServerIP          *ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsOnboardingAAAServerIP          `json:"aaaServerIp,omitempty"`          // Aaa Server Ip
	DhcpServerIP         *ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsOnboardingDhcpServerIP         `json:"dhcpServerIp,omitempty"`         // Dhcp Server Ip
}
type ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsOnboardingAverageRunDuration interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsOnboardingMaxRunDuration interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsOnboardingAverageAssocDuration interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsOnboardingMaxAssocDuration interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsOnboardingAverageAuthDuration interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsOnboardingMaxAuthDuration interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsOnboardingAverageDhcpDuration interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsOnboardingMaxDhcpDuration interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsOnboardingAAAServerIP interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsOnboardingDhcpServerIP interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsOnboardingTime interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsPort interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1ConnectedDevice struct {
	DeviceDetails *ResponseItemUsersGetUserEnrichmentDetailsV1ConnectedDeviceDeviceDetails `json:"deviceDetails,omitempty"` //
}
type ResponseItemUsersGetUserEnrichmentDetailsV1ConnectedDeviceDeviceDetails struct {
	Family                    string                                                                                     `json:"family,omitempty"`                    // Family
	Type                      string                                                                                     `json:"type,omitempty"`                      // Type
	Location                  *ResponseItemUsersGetUserEnrichmentDetailsV1ConnectedDeviceDeviceDetailsLocation           `json:"location,omitempty"`                  // Location
	ErrorCode                 *ResponseItemUsersGetUserEnrichmentDetailsV1ConnectedDeviceDeviceDetailsErrorCode          `json:"errorCode,omitempty"`                 // Error Code
	MacAddress                string                                                                                     `json:"macAddress,omitempty"`                // Mac Address
	Role                      string                                                                                     `json:"role,omitempty"`                      // Role
	ApManagerInterfaceIP      string                                                                                     `json:"apManagerInterfaceIp,omitempty"`      // Ap Manager Interface Ip
	AssociatedWlcIP           string                                                                                     `json:"associatedWlcIp,omitempty"`           // Associated Wlc Ip
	BootDateTime              string                                                                                     `json:"bootDateTime,omitempty"`              // Boot Date Time
	CollectionStatus          string                                                                                     `json:"collectionStatus,omitempty"`          // Collection Status
	InterfaceCount            string                                                                                     `json:"interfaceCount,omitempty"`            // Interface Count
	LineCardCount             string                                                                                     `json:"lineCardCount,omitempty"`             // Line Card Count
	LineCardID                string                                                                                     `json:"lineCardId,omitempty"`                // Line Card Id
	ManagementIPAddress       string                                                                                     `json:"managementIpAddress,omitempty"`       // Management Ip Address
	MemorySize                string                                                                                     `json:"memorySize,omitempty"`                // Memory Size
	PlatformID                string                                                                                     `json:"platformId,omitempty"`                // Platform Id
	ReachabilityFailureReason string                                                                                     `json:"reachabilityFailureReason,omitempty"` // Reachability Failure Reason
	ReachabilityStatus        string                                                                                     `json:"reachabilityStatus,omitempty"`        // Reachability Status
	SNMPContact               string                                                                                     `json:"snmpContact,omitempty"`               // Snmp Contact
	SNMPLocation              string                                                                                     `json:"snmpLocation,omitempty"`              // Snmp Location
	TunnelUDPPort             *ResponseItemUsersGetUserEnrichmentDetailsV1ConnectedDeviceDeviceDetailsTunnelUDPPort      `json:"tunnelUdpPort,omitempty"`             // Tunnel Udp Port
	WaasDeviceMode            *ResponseItemUsersGetUserEnrichmentDetailsV1ConnectedDeviceDeviceDetailsWaasDeviceMode     `json:"waasDeviceMode,omitempty"`            // Waas Device Mode
	Series                    string                                                                                     `json:"series,omitempty"`                    // Series
	InventoryStatusDetail     string                                                                                     `json:"inventoryStatusDetail,omitempty"`     // Inventory Status Detail
	CollectionInterval        string                                                                                     `json:"collectionInterval,omitempty"`        // Collection Interval
	SerialNumber              string                                                                                     `json:"serialNumber,omitempty"`              // Serial Number
	SoftwareVersion           string                                                                                     `json:"softwareVersion,omitempty"`           // Software Version
	RoleSource                string                                                                                     `json:"roleSource,omitempty"`                // Role Source
	Hostname                  string                                                                                     `json:"hostname,omitempty"`                  // Hostname
	UpTime                    string                                                                                     `json:"upTime,omitempty"`                    // Up Time
	LastUpdateTime            *int                                                                                       `json:"lastUpdateTime,omitempty"`            // Last Update Time
	ErrorDescription          *ResponseItemUsersGetUserEnrichmentDetailsV1ConnectedDeviceDeviceDetailsErrorDescription   `json:"errorDescription,omitempty"`          // Error Description
	LocationName              *ResponseItemUsersGetUserEnrichmentDetailsV1ConnectedDeviceDeviceDetailsLocationName       `json:"locationName,omitempty"`              // Location Name
	TagCount                  string                                                                                     `json:"tagCount,omitempty"`                  // Tag Count
	LastUpdated               string                                                                                     `json:"lastUpdated,omitempty"`               // Last Updated
	InstanceUUID              string                                                                                     `json:"instanceUuid,omitempty"`              // Instance Uuid
	ID                        string                                                                                     `json:"id,omitempty"`                        // Id
	NeighborTopology          *[]ResponseItemUsersGetUserEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopology `json:"neighborTopology,omitempty"`          //
}
type ResponseItemUsersGetUserEnrichmentDetailsV1ConnectedDeviceDeviceDetailsLocation interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1ConnectedDeviceDeviceDetailsErrorCode interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1ConnectedDeviceDeviceDetailsTunnelUDPPort interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1ConnectedDeviceDeviceDetailsWaasDeviceMode interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1ConnectedDeviceDeviceDetailsErrorDescription interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1ConnectedDeviceDeviceDetailsLocationName interface{}
type ResponseItemUsersGetUserEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopology struct {
	ErrorCode *int   `json:"errorCode,omitempty"` // Error Code
	Message   string `json:"message,omitempty"`   // Message
	Detail    string `json:"detail,omitempty"`    // Detail
}

//GetUserEnrichmentDetailsV1 Get User Enrichment Details - d7a6-3928-45e8-969d
/* Enriches a given network End User context (a network user-id or end user’s device Mac Address) with details about the user and devices that the user is connected to


@param GetUserEnrichmentDetailsV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-user-enrichment-details-v1
*/
func (s *UsersService) GetUserEnrichmentDetailsV1(GetUserEnrichmentDetailsV1HeaderParams *GetUserEnrichmentDetailsV1HeaderParams) (*ResponseUsersGetUserEnrichmentDetailsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/user-enrichment-details"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetUserEnrichmentDetailsV1HeaderParams != nil {

		if GetUserEnrichmentDetailsV1HeaderParams.EntityType != "" {
			clientRequest = clientRequest.SetHeader("entity_type", GetUserEnrichmentDetailsV1HeaderParams.EntityType)
		}

		if GetUserEnrichmentDetailsV1HeaderParams.EntityValue != "" {
			clientRequest = clientRequest.SetHeader("entity_value", GetUserEnrichmentDetailsV1HeaderParams.EntityValue)
		}

		if GetUserEnrichmentDetailsV1HeaderParams.Persistbapioutput != "" {
			clientRequest = clientRequest.SetHeader("__persistbapioutput", GetUserEnrichmentDetailsV1HeaderParams.Persistbapioutput)
		}

	}

	response, err = clientRequest.
		SetResult(&ResponseUsersGetUserEnrichmentDetailsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetUserEnrichmentDetailsV1(GetUserEnrichmentDetailsV1HeaderParams)
		}
		return nil, response, fmt.Errorf("error with operation GetUserEnrichmentDetailsV1")
	}

	result := response.Result().(*ResponseUsersGetUserEnrichmentDetailsV1)
	return result, response, err

}

// Alias Function
/*
This method acts as an alias for the method `GetUserEnrichmentDetailsV1`
*/
func (s *UsersService) GetUserEnrichmentDetails(GetUserEnrichmentDetailsV1HeaderParams *GetUserEnrichmentDetailsV1HeaderParams) (*ResponseUsersGetUserEnrichmentDetailsV1, *resty.Response, error) {
	return s.GetUserEnrichmentDetailsV1(GetUserEnrichmentDetailsV1HeaderParams)
}
