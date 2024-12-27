package catalyst

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type WirelessService service

type SensorTestResultsV1QueryParams struct {
	SiteID        string  `url:"siteId,omitempty"`        //Assurance site UUID
	StartTime     float64 `url:"startTime,omitempty"`     //The epoch time in milliseconds
	EndTime       float64 `url:"endTime,omitempty"`       //The epoch time in milliseconds
	TestFailureBy string  `url:"testFailureBy,omitempty"` //Obtain failure statistics group by "area", "building", or "floor" (case insensitive)
}
type CreateAndProvisionSSIDV1HeaderParams struct {
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type string. Enable this parameter to execute the API and return a response asynchronously.
}
type DeleteSSIDAndProvisionItToDevicesV1HeaderParams struct {
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type string. Enable this parameter to execute the API and return a response asynchronously.
}
type GetAccessPointRebootTaskResultV1QueryParams struct {
	ParentTaskID string `url:"parentTaskId,omitempty"` //task id of ap reboot request
}
type GetEnterpriseSSIDV1QueryParams struct {
	SSIDName string `url:"ssidName,omitempty"` //Enter the enterprise SSID name that needs to be retrieved. If not entered, all the enterprise SSIDs will be retrieved.
}
type GetSSIDBySiteV1QueryParams struct {
	Limit      float64 `url:"limit,omitempty"`      //The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500.
	Offset     float64 `url:"offset,omitempty"`     //The first record to show for this page; the first record is numbered 1.
	SSID       string  `url:"ssid,omitempty"`       //SSID Name
	WLANType   string  `url:"wlanType,omitempty"`   //Wlan Type
	AuthType   string  `url:"authType,omitempty"`   //Auth Type
	L3AuthType string  `url:"l3authType,omitempty"` //L3 Auth Type
}
type GetSSIDCountBySiteV1QueryParams struct {
	Inherited bool `url:"_inherited,omitempty"` //This query parameter indicates whether the current SSID count at the given 'siteId' is of the SSID(s) it is inheriting or count of non-inheriting SSID(s)
}
type DeleteSSIDV1QueryParams struct {
	RemoveOverrideInHierarchy bool `url:"removeOverrideInHierarchy,omitempty"` //Remove override in hierarchy . Refer Feature tab for details
}
type GetAccessPointConfigurationCountV1QueryParams struct {
	WlcIPAddress string `url:"wlcIpAddress,omitempty"` //WLC IP Address
	ApMode       string `url:"apMode,omitempty"`       //AP Mode. Allowed values are Local, Bridge, Monitor, FlexConnect, Sniffer, Rogue Detector, SE-Connect, Flex+Bridge, Sensor.
	ApModel      string `url:"apModel,omitempty"`      //AP Model
	MeshRole     string `url:"meshRole,omitempty"`     //Mesh Role. Allowed values are RAP or MAP
	Provisioned  string `url:"provisioned,omitempty"`  //Indicate whether AP provisioned or not. Allowed values are True or False
}
type GetAccessPointConfigurationV1QueryParams struct {
	Key          string  `url:"key,omitempty"`          //The ethernet MAC address of Access point
	WlcIPAddress string  `url:"wlcIpAddress,omitempty"` //WLC IP Address
	ApMode       string  `url:"apMode,omitempty"`       //AP Mode. Allowed values are Local, Bridge, Monitor, FlexConnect, Sniffer, Rogue Detector, SE-Connect, Flex+Bridge, Sensor.
	ApModel      string  `url:"apModel,omitempty"`      //AP Model
	MeshRole     string  `url:"meshRole,omitempty"`     //Mesh Role. Allowed values are RAP or MAP
	Provisioned  string  `url:"provisioned,omitempty"`  //Indicate whether AP provisioned or not. Allowed values are True or False
	Limit        float64 `url:"limit,omitempty"`        //The number of records to show for this page. The default is 500 if not specified. The maximum allowed limit is 500.
	Offset       float64 `url:"offset,omitempty"`       //The first record to show for this page; the first record is numbered 1.
}
type ApProvisionConnectivityV1HeaderParams struct {
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type string.
}
type DeleteDynamicInterfaceV1QueryParams struct {
	InterfaceName string `url:"interfaceName,omitempty"` //valid interface-name to be deleted
}
type DeleteDynamicInterfaceV1HeaderParams struct {
	Runsync string `url:"__runsync,omitempty"` //Expects type bool. Enable this parameter to execute the API and return a response synchronously
	Timeout string `url:"__timeout,omitempty"` //Expects type float64. If __runsync is set to ‘true’, this defines the maximum time before which if the API completes its execution, then a synchronous response is returned.  If the time taken for the API to complete the execution, exceeds this time, then an asynchronous response is returned with an execution id, that can be used to get the status and response associated with the API execution
}
type GetDynamicInterfaceV1QueryParams struct {
	InterfaceName string `url:"interface-name,omitempty"` //dynamic-interface name, if not specified all the existing dynamic interfaces will be retrieved
}
type GetWirelessProfileV1QueryParams struct {
	ProfileName string `url:"profileName,omitempty"` //Wireless Network Profile Name
}
type ProvisionUpdateV1HeaderParams struct {
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type string. Enable this parameter to execute the API and return a response asynchronously.
}
type RetrieveRfProfilesV1QueryParams struct {
	RfProfileName string `url:"rf-profile-name,omitempty"` //RF Profile Name
}
type GetAccessPointsFactoryResetStatusV1QueryParams struct {
	TaskID string `url:"taskId,omitempty"` //provide the task id which is returned in the response of ap factory reset post api
}
type GetMeshApNeighboursV1QueryParams struct {
	WlcIPAddress       string `url:"wlcIpAddress,omitempty"`       //Employ this query parameter to obtain the details of the Access points corresponding to the provided WLC IP address.
	ApName             string `url:"apName,omitempty"`             //Employ this query parameter to obtain the details of the Access points corresponding to the provided ap name.
	EthernetMacAddress string `url:"ethernetMacAddress,omitempty"` //Employ this query parameter to obtain the details of the Access points corresponding to the provided EthernetMacAddress.
}
type GetMobilityGroupsV1QueryParams struct {
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` //Employ this query parameter to obtain the details of the Mobility Group corresponding to the provided networkDeviceId. Obtain the network device ID value by using the API GET call /dna/intent/api/v1/network-device/ip-address/${ipAddress}.
}
type GetAnchorManagedApLocationsForSpecificWirelessControllerV1QueryParams struct {
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500.
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1.
}
type GetPrimaryManagedApLocationsForSpecificWirelessControllerV1QueryParams struct {
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500.
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1.
}
type GetSecondaryManagedApLocationsForSpecificWirelessControllerV1QueryParams struct {
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500.
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1.
}
type GetSSIDDetailsForSpecificWirelessControllerV1QueryParams struct {
	SSIDName    string  `url:"ssidName,omitempty"`    //Employ this query parameter to obtain the details of the SSID corresponding to the provided SSID name.
	AdminStatus bool    `url:"adminStatus,omitempty"` //Utilize this query parameter to obtain the administrative status. A 'true' value signifies that the admin status of the SSID is enabled, while a 'false' value indicates that the admin status of the SSID is disabled.
	Managed     bool    `url:"managed,omitempty"`     //If value is 'true' means SSIDs are configured through design.If the value is 'false' means out of band configuration from the Wireless Controller.
	Limit       float64 `url:"limit,omitempty"`       //The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500.
	Offset      float64 `url:"offset,omitempty"`      //The first record to show for this page; the first record is numbered 1.
}
type GetSSIDCountForSpecificWirelessControllerV1QueryParams struct {
	AdminStatus bool `url:"adminStatus,omitempty"` //Utilize this query parameter to obtain the number of SSIDs according to their administrative status. A 'true' value signifies that the admin status of the SSID is enabled, while a 'false' value indicates that the admin status of the SSID is disabled.
	Managed     bool `url:"managed,omitempty"`     //If value is 'true' means SSIDs are configured through design.If the value is 'false' means out of band configuration from the Wireless Controller.
}
type GetWirelessProfilesV1QueryParams struct {
	Limit               float64 `url:"limit,omitempty"`               //The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500
	Offset              float64 `url:"offset,omitempty"`              //The first record to show for this page; the first record is numbered 1
	WirelessProfileName string  `url:"wirelessProfileName,omitempty"` //Wireless Profile Name
}
type RetrieveAllPolicyTagsForAWirelessProfileV1QueryParams struct {
	Limit         float64 `url:"limit,omitempty"`         //The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500.
	Offset        float64 `url:"offset,omitempty"`        //Offset
	PolicyTagName string  `url:"policyTagName,omitempty"` //PolicyTagName
}
type RetrieveAllSiteTagsForAWirelessProfileV1QueryParams struct {
	Limit       float64 `url:"limit,omitempty"`       //The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500.
	Offset      float64 `url:"offset,omitempty"`      //Offset
	SiteTagName string  `url:"siteTagName,omitempty"` //SiteTagName
}
type GetApAuthorizationListsV1QueryParams struct {
	ApAuthorizationListName string `url:"apAuthorizationListName,omitempty"` //Employ this query parameter to obtain the details of the AP Authorization List corresponding to the provided apAuthorizationListName.
	Offset                  string `url:"offset,omitempty"`                  //The first record to show for this page. The first record is numbered 1.
	Limit                   string `url:"limit,omitempty"`                   //The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500.
}
type GetApProfilesV1QueryParams struct {
	Limit         string `url:"limit,omitempty"`         //The number of records to show for this page. The default is 500 if not specified. The maximum allowed limit is 500.
	Offset        string `url:"offset,omitempty"`        //The first record to show for this page; the first record is numbered 1.
	ApProfileName string `url:"apProfileName,omitempty"` //Employ this query parameter to obtain the details of the apProfiles corresponding to the provided apProfileName.
}
type Get80211BeProfilesV1QueryParams struct {
	Limit            float64 `url:"limit,omitempty"`            //The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500.
	Offset           float64 `url:"offset,omitempty"`           //The first record to show for this page, the first record is numbered 1
	ProfileName      string  `url:"profileName,omitempty"`      //Profile Name
	IsOfDmaDownLink  bool    `url:"isOfDmaDownLink,omitempty"`  //OFDMA Downlink
	IsOfDmaUpLink    bool    `url:"isOfDmaUpLink,omitempty"`    //OFDMA Uplink
	IsMuMimoUpLink   bool    `url:"isMuMimoUpLink,omitempty"`   //MU-MIMO Uplink
	IsMuMimoDownLink bool    `url:"isMuMimoDownLink,omitempty"` //MU-MIMO Downlink
	IsOfDmaMultiRu   bool    `url:"isOfDmaMultiRu,omitempty"`   //OFDMA Multi-RU
}
type GetInterfacesV1QueryParams struct {
	Limit         float64 `url:"limit,omitempty"`         //The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500.
	Offset        float64 `url:"offset,omitempty"`        //The first record to show for this page. The first record is numbered 1.
	InterfaceName string  `url:"interfaceName,omitempty"` //Interface Name
	VLANID        float64 `url:"vlanId,omitempty"`        //Vlan Id
}
type GetPowerProfilesV1QueryParams struct {
	Limit       float64 `url:"limit,omitempty"`       //Limit
	Offset      float64 `url:"offset,omitempty"`      //Offset
	ProfileName string  `url:"profileName,omitempty"` //Power Profile Name
}
type GetRfProfilesV1QueryParams struct {
	Limit               float64 `url:"limit,omitempty"`               //The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500
	Offset              float64 `url:"offset,omitempty"`              //The first record to show for this page; the first record is numbered 1
	RfProfileName       string  `url:"rfProfileName,omitempty"`       //RF Profile Name
	EnableRadioTypeA    bool    `url:"enableRadioTypeA,omitempty"`    //Enable Radio TypeA
	EnableRadioTypeB    bool    `url:"enableRadioTypeB,omitempty"`    //Enable Radio TypeB
	EnableRadioType6GHz bool    `url:"enableRadioType6GHz,omitempty"` //Enable Radio Type6GHz
}
type RetrieveSitesWithOverriddenSSIDsV1QueryParams struct {
	SiteID string  `url:"siteId,omitempty"` //Site UUID
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500.
}

type ResponseWirelessSensorTestResultsV1 struct {
	Version string `json:"version,omitempty"` // Version

	Response *ResponseWirelessSensorTestResultsV1Response `json:"response,omitempty"` //
}
type ResponseWirelessSensorTestResultsV1Response struct {
	Summary *ResponseWirelessSensorTestResultsV1ResponseSummary `json:"summary,omitempty"` //

	FailureStats *[]ResponseWirelessSensorTestResultsV1ResponseFailureStats `json:"failureStats,omitempty"` //
}
type ResponseWirelessSensorTestResultsV1ResponseSummary struct {
	TotalTestCount *int `json:"totalTestCount,omitempty"` // Total test count

	OnBoarding *ResponseWirelessSensorTestResultsV1ResponseSummaryOnBoarding `json:"ONBOARDING,omitempty"` //

	PERfORMAncE *ResponseWirelessSensorTestResultsV1ResponseSummaryPERfORMAncE `json:"PERFORMANCE,omitempty"` //

	NETWORKSERVICES *ResponseWirelessSensorTestResultsV1ResponseSummaryNETWORKSERVICES `json:"NETWORK_SERVICES,omitempty"` //

	ApPCONNECTIVITY *ResponseWirelessSensorTestResultsV1ResponseSummaryApPCONNECTIVITY `json:"APP_CONNECTIVITY,omitempty"` //

	RfASSESSMENT *ResponseWirelessSensorTestResultsV1ResponseSummaryRfASSESSMENT `json:"RF_ASSESSMENT,omitempty"` //

	Email *ResponseWirelessSensorTestResultsV1ResponseSummaryEmail `json:"EMAIL,omitempty"` //
}
type ResponseWirelessSensorTestResultsV1ResponseSummaryOnBoarding struct {
	Auth *ResponseWirelessSensorTestResultsV1ResponseSummaryOnBoardingAuth `json:"AUTH,omitempty"` //

	DHCP *ResponseWirelessSensorTestResultsV1ResponseSummaryOnBoardingDHCP `json:"DHCP,omitempty"` //

	Assoc *ResponseWirelessSensorTestResultsV1ResponseSummaryOnBoardingAssoc `json:"ASSOC,omitempty"` //
}
type ResponseWirelessSensorTestResultsV1ResponseSummaryOnBoardingAuth struct {
	PassCount *int `json:"passCount,omitempty"` // Total passed test count

	FailCount *int `json:"failCount,omitempty"` // Total failed test count
}
type ResponseWirelessSensorTestResultsV1ResponseSummaryOnBoardingDHCP struct {
	PassCount *int `json:"passCount,omitempty"` // Total passed test count

	FailCount *float64 `json:"failCount,omitempty"` // Total failed test count
}
type ResponseWirelessSensorTestResultsV1ResponseSummaryOnBoardingAssoc struct {
	PassCount *int `json:"passCount,omitempty"` // Total passed test count

	FailCount *int `json:"failCount,omitempty"` // Total failed test count
}
type ResponseWirelessSensorTestResultsV1ResponseSummaryPERfORMAncE struct {
	IPSLASENDER *ResponseWirelessSensorTestResultsV1ResponseSummaryPERfORMAncEIPSLASENDER `json:"IPSLASENDER,omitempty"` //
}
type ResponseWirelessSensorTestResultsV1ResponseSummaryPERfORMAncEIPSLASENDER struct {
	PassCount *int `json:"passCount,omitempty"` // Total passed test count

	FailCount *int `json:"failCount,omitempty"` // Total failed test count
}
type ResponseWirelessSensorTestResultsV1ResponseSummaryNETWORKSERVICES struct {
	DNS *ResponseWirelessSensorTestResultsV1ResponseSummaryNETWORKSERVICESDNS `json:"DNS,omitempty"` //
}
type ResponseWirelessSensorTestResultsV1ResponseSummaryNETWORKSERVICESDNS struct {
	PassCount *int `json:"passCount,omitempty"` // Total passed test count

	FailCount *float64 `json:"failCount,omitempty"` // Total failed test count
}
type ResponseWirelessSensorTestResultsV1ResponseSummaryApPCONNECTIVITY struct {
	HOSTREACHABILITY *ResponseWirelessSensorTestResultsV1ResponseSummaryApPCONNECTIVITYHOSTREACHABILITY `json:"HOST_REACHABILITY,omitempty"` //

	WebServer *ResponseWirelessSensorTestResultsV1ResponseSummaryApPCONNECTIVITYWebServer `json:"WEBSERVER,omitempty"` //

	FileTransfer *ResponseWirelessSensorTestResultsV1ResponseSummaryApPCONNECTIVITYFileTransfer `json:"FILETRANSFER,omitempty"` //
}
type ResponseWirelessSensorTestResultsV1ResponseSummaryApPCONNECTIVITYHOSTREACHABILITY struct {
	PassCount *int `json:"passCount,omitempty"` // Total passed test count

	FailCount *float64 `json:"failCount,omitempty"` // Total failed test count
}
type ResponseWirelessSensorTestResultsV1ResponseSummaryApPCONNECTIVITYWebServer struct {
	PassCount *int `json:"passCount,omitempty"` // Total passed test count

	FailCount *int `json:"failCount,omitempty"` // Total failed test count
}
type ResponseWirelessSensorTestResultsV1ResponseSummaryApPCONNECTIVITYFileTransfer struct {
	PassCount *float64 `json:"passCount,omitempty"` // Total passed test count

	FailCount *int `json:"failCount,omitempty"` // Total failed test count
}
type ResponseWirelessSensorTestResultsV1ResponseSummaryRfASSESSMENT struct {
	DATARATE *ResponseWirelessSensorTestResultsV1ResponseSummaryRfASSESSMENTDATARATE `json:"DATA_RATE,omitempty"` //

	SNR *ResponseWirelessSensorTestResultsV1ResponseSummaryRfASSESSMENTSNR `json:"SNR,omitempty"` //
}
type ResponseWirelessSensorTestResultsV1ResponseSummaryRfASSESSMENTDATARATE struct {
	PassCount *int `json:"passCount,omitempty"` // Total passed test count

	FailCount *int `json:"failCount,omitempty"` // Total failed test count
}
type ResponseWirelessSensorTestResultsV1ResponseSummaryRfASSESSMENTSNR struct {
	PassCount *int `json:"passCount,omitempty"` // Total passed test count

	FailCount *float64 `json:"failCount,omitempty"` // Total failed test count
}
type ResponseWirelessSensorTestResultsV1ResponseSummaryEmail struct {
	MailServer *ResponseWirelessSensorTestResultsV1ResponseSummaryEmailMailServer `json:"MAILSERVER,omitempty"` //
}
type ResponseWirelessSensorTestResultsV1ResponseSummaryEmailMailServer struct {
	PassCount *float64 `json:"passCount,omitempty"` // Total passed test count

	FailCount *int `json:"failCount,omitempty"` // Total failed test count
}
type ResponseWirelessSensorTestResultsV1ResponseFailureStats struct {
	ErrorCode *int `json:"errorCode,omitempty"` // The error code

	ErrorTitle string `json:"errorTitle,omitempty"` // The error title

	TestType string `json:"testType,omitempty"` // The test type

	TestCategory string `json:"testCategory,omitempty"` // The test category
}
type ResponseWirelessCreateAndProvisionSSIDV1 struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url

	Message string `json:"message,omitempty"` // Message
}
type ResponseWirelessDeleteSSIDAndProvisionItToDevicesV1 struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url

	Message string `json:"message,omitempty"` // Message
}
type ResponseWirelessRebootAccessPointsV1 struct {
	Response *ResponseWirelessRebootAccessPointsV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseWirelessRebootAccessPointsV1Response struct {
	TaskID string `json:"taskId,omitempty"` //

	URL string `json:"url,omitempty"` //
}
type ResponseWirelessGetAccessPointRebootTaskResultV1 []ResponseItemWirelessGetAccessPointRebootTaskResultV1 // Array of ResponseWirelessGetAccessPointRebootTaskResultV1
type ResponseItemWirelessGetAccessPointRebootTaskResultV1 struct {
	WlcIP string `json:"wlcIP,omitempty"` //

	ApList *[]ResponseItemWirelessGetAccessPointRebootTaskResultV1ApList `json:"apList,omitempty"` //
}
type ResponseItemWirelessGetAccessPointRebootTaskResultV1ApList struct {
	ApName string `json:"apName,omitempty"` //

	RebootStatus string `json:"rebootStatus,omitempty"` //

	FailureReason *ResponseItemWirelessGetAccessPointRebootTaskResultV1ApListFailureReason `json:"failureReason,omitempty"` //
}
type ResponseItemWirelessGetAccessPointRebootTaskResultV1ApListFailureReason interface{}
type ResponseWirelessGetEnterpriseSSIDV1 []ResponseItemWirelessGetEnterpriseSSIDV1 // Array of ResponseWirelessGetEnterpriseSSIDV1
type ResponseItemWirelessGetEnterpriseSSIDV1 struct {
	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid

	Version *int `json:"version,omitempty"` // Version

	SSIDDetails *[]ResponseItemWirelessGetEnterpriseSSIDV1SSIDDetails `json:"ssidDetails,omitempty"` //

	GroupUUID string `json:"groupUuid,omitempty"` // Group Uuid

	InheritedGroupUUID string `json:"inheritedGroupUuid,omitempty"` // Inherited Group Uuid

	InheritedGroupName string `json:"inheritedGroupName,omitempty"` // Inherited Group Name
}
type ResponseItemWirelessGetEnterpriseSSIDV1SSIDDetails struct {
	Name string `json:"name,omitempty"` // SSID Name

	WLANType string `json:"wlanType,omitempty"` // Wlan Type

	EnableFastLane *bool `json:"enableFastLane,omitempty"` // Enable Fast Lane

	SecurityLevel string `json:"securityLevel,omitempty"` // Security Level

	AuthServer string `json:"authServer,omitempty"` // Auth Server

	Passphrase string `json:"passphrase,omitempty"` // Passphrase

	TrafficType string `json:"trafficType,omitempty"` // Traffic Type

	EnableMacFiltering *bool `json:"enableMACFiltering,omitempty"` // Enable MAC Filtering

	IsEnabled *bool `json:"isEnabled,omitempty"` // Is Enabled

	IsFabric *bool `json:"isFabric,omitempty"` // Is Fabric

	FastTransition string `json:"fastTransition,omitempty"` // Fast Transition

	RadioPolicy string `json:"radioPolicy,omitempty"` // Radio Policy

	EnableBroadcastSSID *bool `json:"enableBroadcastSSID,omitempty"` // Enable Broadcast SSID

	NasOptions []string `json:"nasOptions,omitempty"` // Nas Options

	AAAOverride *bool `json:"aaaOverride,omitempty"` // Aaa Override

	CoverageHoleDetectionEnable *bool `json:"coverageHoleDetectionEnable,omitempty"` // Coverage Hole Detection Enable

	ProtectedManagementFrame string `json:"protectedManagementFrame,omitempty"` // Protected Management Frame

	MultipSKSettings *[]ResponseItemWirelessGetEnterpriseSSIDV1SSIDDetailsMultipSKSettings `json:"multiPSKSettings,omitempty"` //

	ClientRateLimit *float64 `json:"clientRateLimit,omitempty"` // Client Rate Limit. (in bits per second)

	EnableSessionTimeOut *bool `json:"enableSessionTimeOut,omitempty"` // Enable Session Time Out

	SessionTimeOut *float64 `json:"sessionTimeOut,omitempty"` // sessionTimeOut

	EnableClientExclusion *bool `json:"enableClientExclusion,omitempty"` // Enable Client Exclusion

	ClientExclusionTimeout *float64 `json:"clientExclusionTimeout,omitempty"` // Client Exclusion Timeout

	EnableBasicServiceSetMaxIDle *bool `json:"enableBasicServiceSetMaxIdle,omitempty"` // Enable Basic Service Set Max Idle

	BasicServiceSetClientIDleTimeout *float64 `json:"basicServiceSetClientIdleTimeout,omitempty"` // Basic Service Set ClientIdle Timeout

	EnableDirectedMulticastService *bool `json:"enableDirectedMulticastService,omitempty"` // Enable Directed MulticastService

	EnableNeighborList *bool `json:"enableNeighborList,omitempty"` // Enable NeighborList

	MfpClientProtection string `json:"mfpClientProtection,omitempty"` // Mfp Client Protection
}
type ResponseItemWirelessGetEnterpriseSSIDV1SSIDDetailsMultipSKSettings struct {
	Priority *int `json:"priority,omitempty"` // Priority

	PassphraseType string `json:"passphraseType,omitempty"` // Passphrase Type

	Passphrase string `json:"passphrase,omitempty"` // Passphrase
}
type ResponseWirelessCreateEnterpriseSSIDV1 struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url

	Message string `json:"message,omitempty"` // Message
}
type ResponseWirelessUpdateEnterpriseSSIDV1 struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url

	Message string `json:"message,omitempty"` // Message
}
type ResponseWirelessDeleteEnterpriseSSIDV1 struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url

	Message string `json:"message,omitempty"` // Message
}
type ResponseWirelessCreateSSIDV1 struct {
	Response *ResponseWirelessCreateSSIDV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessCreateSSIDV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessGetSSIDBySiteV1 struct {
	Response *[]ResponseWirelessGetSSIDBySiteV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version of the response
}
type ResponseWirelessGetSSIDBySiteV1Response struct {
	SSID string `json:"ssid,omitempty"` // Name of the SSID

	AuthType string `json:"authType,omitempty"` // L2 Authentication Type (If authType is not open , then atleast one RSN Cipher Suite and corresponding valid AKM must be enabled)

	Passphrase string `json:"passphrase,omitempty"` // Passphrase (Only applicable for SSID with PERSONAL security level). Passphrase needs to be between 8 and 63 characters for ASCII type. HEX passphrase needs to be 64 characters

	IsFastLaneEnabled *bool `json:"isFastLaneEnabled,omitempty"` // When set to true, MAC Filtering will be activated, allowing control over network access based on the MAC address of the device

	IsMacFilteringEnabled *bool `json:"isMacFilteringEnabled,omitempty"` // True if MAC Filtering is enabled, else False

	SSIDRadioType string `json:"ssidRadioType,omitempty"` // Radio Policy Enum (default: Triple band operation(2.4GHz, 5GHz and 6GHz))

	IsBroadcastSSID *bool `json:"isBroadcastSSID,omitempty"` // When activated by setting it to true, the Broadcast SSID feature will make the SSID publicly visible to wireless devices searching for available networks

	FastTransition string `json:"fastTransition,omitempty"` // Fast Transition

	SessionTimeOutEnable *bool `json:"sessionTimeOutEnable,omitempty"` // Turn on the feature that imposes a time limit on user sessions

	SessionTimeOut *int `json:"sessionTimeOut,omitempty"` // This denotes the allotted time span, expressed in seconds, before a session is automatically terminated due to inactivity

	ClientExclusionEnable *bool `json:"clientExclusionEnable,omitempty"` // Activate the feature that allows for the exclusion of clients

	ClientExclusionTimeout *int `json:"clientExclusionTimeout,omitempty"` // This refers to the length of time, in seconds, a client is excluded or blocked from accessing the network after a specified number of unsuccessful attempts

	BasicServiceSetMaxIDleEnable *bool `json:"basicServiceSetMaxIdleEnable,omitempty"` // Activate the maximum idle feature for the Basic Service Set

	BasicServiceSetClientIDleTimeout *int `json:"basicServiceSetClientIdleTimeout,omitempty"` // This refers to the duration of inactivity, measured in seconds, before a client connected to the Basic Service Set is considered idle and timed out

	DirectedMulticastServiceEnable *bool `json:"directedMulticastServiceEnable,omitempty"` // The Directed Multicast Service feature becomes operational when it is set to true

	NeighborListEnable *bool `json:"neighborListEnable,omitempty"` // The Neighbor List feature is enabled when it is set to true

	ManagementFrameProtectionClientprotection string `json:"managementFrameProtectionClientprotection,omitempty"` // Management Frame Protection Client

	NasOptions []string `json:"nasOptions,omitempty"` // Pre-Defined NAS Options : AP ETH Mac Address, AP IP address, AP Location , AP MAC Address, AP Name, AP Policy Tag, AP Site Tag, SSID, System IP Address, System MAC Address, System Name.

	ProfileName string `json:"profileName,omitempty"` // WLAN Profile Name, if not passed autogenerated profile name will be assigned

	PolicyProfileName string `json:"policyProfileName,omitempty"` // Policy Profile Name. If not passed, profileName value will be used to populate this parameter

	AAAOverride *bool `json:"aaaOverride,omitempty"` // Activate the AAA Override feature when set to true

	CoverageHoleDetectionEnable *bool `json:"coverageHoleDetectionEnable,omitempty"` // Activate Coverage Hole Detection feature when set to true

	ProtectedManagementFrame string `json:"protectedManagementFrame,omitempty"` // (REQUIRED is applicable for authType WPA3_PERSONAL, WPA3_ENTERPRISE, OPEN_SECURED) and (OPTIONAL/REQUIRED is applicable for authType WPA2_WPA3_PERSONAL and WPA2_WPA3_ENTERPRISE)

	MultipSKSettings *[]ResponseWirelessGetSSIDBySiteV1ResponseMultipSKSettings `json:"multiPSKSettings,omitempty"` //

	ClientRateLimit *int `json:"clientRateLimit,omitempty"` // This pertains to the maximum data transfer rate, specified in bits per second, that a client is permitted to achieve

	RsnCipherSuiteGcmp256 *bool `json:"rsnCipherSuiteGcmp256,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite GCMP256 encryption protocol is activated

	RsnCipherSuiteCcmp256 *bool `json:"rsnCipherSuiteCcmp256,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite CCMP256 encryption protocol is activated

	RsnCipherSuiteGcmp128 *bool `json:"rsnCipherSuiteGcmp128,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite GCMP128 encryption protocol is activated

	RsnCipherSuiteCcmp128 *bool `json:"rsnCipherSuiteCcmp128,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite CCMP128 encryption protocol is activated

	Ghz6PolicyClientSteering *bool `json:"ghz6PolicyClientSteering,omitempty"` // True if 6 GHz Policy Client Steering is enabled, else False

	IsAuthKey8021X *bool `json:"isAuthKey8021x,omitempty"` // When set to true, the 802.1X authentication key is in use

	IsAuthKey8021XPlusFT *bool `json:"isAuthKey8021xPlusFT,omitempty"` // When set to true, the 802.1X-Plus-FT authentication key is in use

	IsAuthKey8021XSHA256 *bool `json:"isAuthKey8021x_SHA256,omitempty"` // When set to true, the feature that enables 802.1X authentication using the SHA256 algorithm is turned on

	IsAuthKeySae *bool `json:"isAuthKeySae,omitempty"` // When set to true, the feature enabling the Simultaneous Authentication of Equals (SAE) authentication key is activated

	IsAuthKeySaePlusFT *bool `json:"isAuthKeySaePlusFT,omitempty"` // Activating this setting by switching it to true turns on the authentication key feature that supports both Simultaneous Authentication of Equals (SAE) and Fast Transition (FT)

	IsAuthKeyPSK *bool `json:"isAuthKeyPSK,omitempty"` // When set to true, the Pre-shared Key (PSK) authentication feature is enabled

	IsAuthKeyPSKPlusFT *bool `json:"isAuthKeyPSKPlusFT,omitempty"` // When set to true, the feature that enables the combination of Pre-shared Key (PSK) and Fast Transition (FT) authentication keys is activated

	IsAuthKeyOWE *bool `json:"isAuthKeyOWE,omitempty"` // When set to true, the Opportunistic Wireless Encryption (OWE) authentication key feature is turned on

	IsAuthKeyEasyPSK *bool `json:"isAuthKeyEasyPSK,omitempty"` // When set to true, the feature that enables the use of Easy Pre-shared Key (PSK) authentication is activated

	IsAuthKeyPSKSHA256 *bool `json:"isAuthKeyPSKSHA256,omitempty"` // The feature that allows the use of Pre-shared Key (PSK) authentication with the SHA256 algorithm is enabled when it is set to true

	OpenSSID string `json:"openSsid,omitempty"` // Open SSID which is already created in the design and not associated to any other OPEN-SECURED SSID

	IsCustomNasIDOptions *bool `json:"isCustomNasIdOptions,omitempty"` // Set to true if Custom NAS ID Options provided

	WLANBandSelectEnable *bool `json:"wlanBandSelectEnable,omitempty"` // Band select is allowed only when band options selected contains at least 2.4 GHz and 5 GHz band

	IsEnabled *bool `json:"isEnabled,omitempty"` // Set SSID's admin status as 'Enabled' when set to true

	AuthServers []string `json:"authServers,omitempty"` // List of Authentication/Authorization server IpAddresses

	AcctServers []string `json:"acctServers,omitempty"` // List of Accounting server IpAddresses

	EgressQos string `json:"egressQos,omitempty"` // Egress QOS

	IngressQos string `json:"ingressQos,omitempty"` // Ingress QOS

	InheritedSiteID string `json:"inheritedSiteId,omitempty"` // Site UUID from where the SSID is inherited

	InheritedSiteName string `json:"inheritedSiteName,omitempty"` // Site Name from where the SSID is inherited

	WLANType string `json:"wlanType,omitempty"` // Wlan Type

	L3AuthType string `json:"l3AuthType,omitempty"` // L3 Authentication Type

	AuthServer string `json:"authServer,omitempty"` // Authentication Server, Mandatory for Guest SSIDs with wlanType=Guest and l3AuthType=web_auth

	ExternalAuthIPAddress string `json:"externalAuthIpAddress,omitempty"` // External WebAuth URL (Mandatory for Guest SSIDs with wlanType = Guest, l3AuthType = web_auth and authServer = auth_external)

	WebPassthrough *bool `json:"webPassthrough,omitempty"` // When set to true, the Web-Passthrough feature will be activated for the Guest SSID, allowing guests to bypass certain login requirements

	SleepingClientEnable *bool `json:"sleepingClientEnable,omitempty"` // When set to true, this will activate the timeout settings that apply to clients in sleep mode

	SleepingClientTimeout *int `json:"sleepingClientTimeout,omitempty"` // This refers to the amount of time, measured in minutes, before a sleeping (inactive) client is timed out of the network

	ACLName string `json:"aclName,omitempty"` // Pre-Auth Access Control List (ACL) Name

	IsPosturingEnabled *bool `json:"isPosturingEnabled,omitempty"` // Applicable only for Enterprise SSIDs. When set to True, Posturing will enabled. Required to be set to True if ACL needs to be mapped for Enterprise SSID.

	IsAuthKeySuiteB1X *bool `json:"isAuthKeySuiteB1x,omitempty"` // When activated by setting it to true, the SuiteB-1x authentication key feature is engaged.

	IsAuthKeySuiteB1921X *bool `json:"isAuthKeySuiteB1921x,omitempty"` // When set to true, the SuiteB192-1x authentication key feature is enabled.

	IsAuthKeySaeExt *bool `json:"isAuthKeySaeExt,omitempty"` // When set to true, the Simultaneous Authentication of Equals (SAE) Extended Authentication key feature is turned on.

	IsAuthKeySaeExtPlusFT *bool `json:"isAuthKeySaeExtPlusFT,omitempty"` // When set to true, the Simultaneous Authentication of Equals (SAE) combined with Fast Transition (FT) Authentication Key feature is enabled.

	IsApBeaconProtectionEnabled *bool `json:"isApBeaconProtectionEnabled,omitempty"` // When set to true, the Access Point (AP) Beacon Protection feature is activated, enhancing the security of the network.

	Ghz24Policy string `json:"ghz24Policy,omitempty"` // 2.4 Ghz Band Policy value. Allowed only when 2.4 Radio Band is enabled in ssidRadioType

	CckmTsfTolerance *int `json:"cckmTsfTolerance,omitempty"` // Cckm TImestamp Tolerance(in milliseconds)

	IsCckmEnabled *bool `json:"isCckmEnabled,omitempty"` // True if CCKM is enabled, else False

	IsHex *bool `json:"isHex,omitempty"` // True if passphrase is in Hex format, else False.

	IsSensorPnp *bool `json:"isSensorPnp,omitempty"` // True if SSID is a sensor SSID

	ID string `json:"id,omitempty"` // SSID ID

	IsRandomMacFilterEnabled *bool `json:"isRandomMacFilterEnabled,omitempty"` // Deny clients using randomized MAC addresses when set to true

	FastTransitionOverTheDistributedSystemEnable *bool `json:"fastTransitionOverTheDistributedSystemEnable,omitempty"` // Enable Fast Transition over the Distributed System when set to true

	InheritedSiteNameHierarchy string `json:"inheritedSiteNameHierarchy,omitempty"` // Inherited Site Name Hierarchy
}
type ResponseWirelessGetSSIDBySiteV1ResponseMultipSKSettings struct {
	Priority *int `json:"priority,omitempty"` // Priority

	PassphraseType string `json:"passphraseType,omitempty"` // Passphrase Type

	Passphrase string `json:"passphrase,omitempty"` // Passphrase needs to be between 8 and 63 characters for ASCII type. HEX passphrase needs to be 64 characters
}
type ResponseWirelessGetSSIDCountBySiteV1 struct {
	Response *ResponseWirelessGetSSIDCountBySiteV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response Version
}
type ResponseWirelessGetSSIDCountBySiteV1Response struct {
	Count *int `json:"count,omitempty"` // Count of the requested resource
}
type ResponseWirelessGetSSIDByIDV1 struct {
	Response *ResponseWirelessGetSSIDByIDV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version of the response
}
type ResponseWirelessGetSSIDByIDV1Response struct {
	SSID string `json:"ssid,omitempty"` // Name of the SSID

	AuthType string `json:"authType,omitempty"` // L2 Authentication Type (If authType is not open , then atleast one RSN Cipher Suite and corresponding valid AKM must be enabled)

	Passphrase string `json:"passphrase,omitempty"` // Passphrase (Only applicable for SSID with PERSONAL security level). Passphrase needs to be between 8 and 63 characters for ASCII type. HEX passphrase needs to be 64 characters

	IsFastLaneEnabled *bool `json:"isFastLaneEnabled,omitempty"` // True if FastLane is enabled, else False

	IsMacFilteringEnabled *bool `json:"isMacFilteringEnabled,omitempty"` // When set to true, MAC Filtering will be activated, allowing control over network access based on the MAC address of the device

	SSIDRadioType string `json:"ssidRadioType,omitempty"` // Radio Policy Enum (default: Triple band operation(2.4GHz, 5GHz and 6GHz))

	IsBroadcastSSID *bool `json:"isBroadcastSSID,omitempty"` // When activated by setting it to true, the Broadcast SSID feature will make the SSID publicly visible to wireless devices searching for available networks

	FastTransition string `json:"fastTransition,omitempty"` // Fast Transition

	SessionTimeOutEnable *bool `json:"sessionTimeOutEnable,omitempty"` // Turn on the feature that imposes a time limit on user sessions

	SessionTimeOut *int `json:"sessionTimeOut,omitempty"` // This denotes the allotted time span, expressed in seconds, before a session is automatically terminated due to inactivity

	ClientExclusionEnable *bool `json:"clientExclusionEnable,omitempty"` // Activate the feature that allows for the exclusion of clients

	ClientExclusionTimeout *int `json:"clientExclusionTimeout,omitempty"` // This refers to the length of time, in seconds, a client is excluded or blocked from accessing the network after a specified number of unsuccessful attempts

	BasicServiceSetMaxIDleEnable *bool `json:"basicServiceSetMaxIdleEnable,omitempty"` // Activate the maximum idle feature for the Basic Service Set

	BasicServiceSetClientIDleTimeout *int `json:"basicServiceSetClientIdleTimeout,omitempty"` // This refers to the duration of inactivity, measured in seconds, before a client connected to the Basic Service Set is considered idle and timed out

	DirectedMulticastServiceEnable *bool `json:"directedMulticastServiceEnable,omitempty"` // The Directed Multicast Service feature becomes operational when it is set to true

	NeighborListEnable *bool `json:"neighborListEnable,omitempty"` // The Neighbor List feature is enabled when it is set to true

	ManagementFrameProtectionClientprotection string `json:"managementFrameProtectionClientprotection,omitempty"` // Management Frame Protection Client

	NasOptions []string `json:"nasOptions,omitempty"` // Pre-Defined NAS Options : AP ETH Mac Address, AP IP address, AP Location , AP MAC Address, AP Name, AP Policy Tag, AP Site Tag, SSID, System IP Address, System MAC Address, System Name.

	ProfileName string `json:"profileName,omitempty"` // WLAN Profile Name, if not passed autogenerated profile name will be assigned

	PolicyProfileName string `json:"policyProfileName,omitempty"` // Policy Profile Name. If not passed, profileName value will be used to populate this parameter

	AAAOverride *bool `json:"aaaOverride,omitempty"` // Activate the AAA Override feature when set to true

	CoverageHoleDetectionEnable *bool `json:"coverageHoleDetectionEnable,omitempty"` // Activate Coverage Hole Detection feature when set to true

	ProtectedManagementFrame string `json:"protectedManagementFrame,omitempty"` // (REQUIRED is applicable for authType WPA3_PERSONAL, WPA3_ENTERPRISE, OPEN_SECURED) and (OPTIONAL/REQUIRED is applicable for authType WPA2_WPA3_PERSONAL and WPA2_WPA3_ENTERPRISE)

	MultipSKSettings *[]ResponseWirelessGetSSIDByIDV1ResponseMultipSKSettings `json:"multiPSKSettings,omitempty"` //

	ClientRateLimit *int `json:"clientRateLimit,omitempty"` // This pertains to the maximum data transfer rate, specified in bits per second, that a client is permitted to achieve

	RsnCipherSuiteGcmp256 *bool `json:"rsnCipherSuiteGcmp256,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite GCMP256 encryption protocol is activated

	RsnCipherSuiteCcmp256 *bool `json:"rsnCipherSuiteCcmp256,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite CCMP256 encryption protocol is activated

	RsnCipherSuiteGcmp128 *bool `json:"rsnCipherSuiteGcmp128,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite GCMP128 encryption protocol is activatedTrue if RSN Cipher Suite GCMP128 is enabled, else False

	RsnCipherSuiteCcmp128 *bool `json:"rsnCipherSuiteCcmp128,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite CCMP128 encryption protocol is activated

	Ghz6PolicyClientSteering *bool `json:"ghz6PolicyClientSteering,omitempty"` // True if 6 GHz Policy Client Steering is enabled, else False

	IsAuthKey8021X *bool `json:"isAuthKey8021x,omitempty"` // When set to true, the 802.1X authentication key is in use

	IsAuthKey8021XPlusFT *bool `json:"isAuthKey8021xPlusFT,omitempty"` // When set to true, the 802.1X-Plus-FT authentication key is in use

	IsAuthKey8021XSHA256 *bool `json:"isAuthKey8021x_SHA256,omitempty"` // When set to true, the feature that enables 802.1X authentication using the SHA256 algorithm is turned on

	IsAuthKeySae *bool `json:"isAuthKeySae,omitempty"` // When set to true, the feature enabling the Simultaneous Authentication of Equals (SAE) authentication key is activated

	IsAuthKeySaePlusFT *bool `json:"isAuthKeySaePlusFT,omitempty"` // Activating this setting by switching it to true turns on the authentication key feature that supports both Simultaneous Authentication of Equals (SAE) and Fast Transition (FT)

	IsAuthKeyPSK *bool `json:"isAuthKeyPSK,omitempty"` // When set to true, the Pre-shared Key (PSK) authentication feature is enabled

	IsAuthKeyPSKPlusFT *bool `json:"isAuthKeyPSKPlusFT,omitempty"` // When set to true, the feature that enables the combination of Pre-shared Key (PSK) and Fast Transition (FT) authentication keys is activated

	IsAuthKeyOWE *bool `json:"isAuthKeyOWE,omitempty"` // When set to true, the Opportunistic Wireless Encryption (OWE) authentication key feature is turned on

	IsAuthKeyEasyPSK *bool `json:"isAuthKeyEasyPSK,omitempty"` // When set to true, the feature that enables the use of Easy Pre-shared Key (PSK) authentication is activated

	IsAuthKeyPSKSHA256 *bool `json:"isAuthKeyPSKSHA256,omitempty"` // The feature that allows the use of Pre-shared Key (PSK) authentication with the SHA256 algorithm is enabled when it is set to true

	OpenSSID string `json:"openSsid,omitempty"` // Open SSID which is already created in the design and not associated to any other OPEN-SECURED SSID

	IsCustomNasIDOptions *bool `json:"isCustomNasIdOptions,omitempty"` // Set to true if Custom NAS ID Options provided

	WLANBandSelectEnable *bool `json:"wlanBandSelectEnable,omitempty"` // Band select is allowed only when band options selected contains at least 2.4 GHz and 5 GHz band

	IsEnabled *bool `json:"isEnabled,omitempty"` // Set SSID's admin status as 'Enabled' when set to true

	AuthServers []string `json:"authServers,omitempty"` // List of Authentication/Authorization server IpAddresses

	AcctServers []string `json:"acctServers,omitempty"` // List of Accounting server IpAddresses

	EgressQos string `json:"egressQos,omitempty"` // Egress QOS

	IngressQos string `json:"ingressQos,omitempty"` // Ingress QOS

	InheritedSiteID string `json:"inheritedSiteId,omitempty"` // Site UUID from where the SSID is inherited

	InheritedSiteName string `json:"inheritedSiteName,omitempty"` // Site Name from where the SSID is inherited

	WLANType string `json:"wlanType,omitempty"` // Wlan Type

	L3AuthType string `json:"l3AuthType,omitempty"` // L3 Authentication Type

	AuthServer string `json:"authServer,omitempty"` // Authentication Server, Mandatory for Guest SSIDs with wlanType=Guest and l3AuthType=web_auth

	ExternalAuthIPAddress string `json:"externalAuthIpAddress,omitempty"` // External WebAuth URL (Mandatory for Guest SSIDs with wlanType = Guest, l3AuthType = web_auth and authServer = auth_external)

	WebPassthrough *bool `json:"webPassthrough,omitempty"` // When set to true, the Web-Passthrough feature will be activated for the Guest SSID, allowing guests to bypass certain login requirements

	SleepingClientEnable *bool `json:"sleepingClientEnable,omitempty"` // When set to true, this will activate the timeout settings that apply to clients in sleep mode

	SleepingClientTimeout *int `json:"sleepingClientTimeout,omitempty"` // This refers to the amount of time, measured in minutes, before a sleeping (inactive) client is timed out of the network

	ACLName string `json:"aclName,omitempty"` // Pre-Auth Access Control List (ACL) Name

	IsPosturingEnabled *bool `json:"isPosturingEnabled,omitempty"` // Applicable only for Enterprise SSIDs. When set to True, Posturing will enabled. Required to be set to True if ACL needs to be mapped for Enterprise SSID.

	IsAuthKeySuiteB1X *bool `json:"isAuthKeySuiteB1x,omitempty"` // When activated by setting it to true, the SuiteB-1x authentication key feature is engaged.

	IsAuthKeySuiteB1921X *bool `json:"isAuthKeySuiteB1921x,omitempty"` // When set to true, the SuiteB192-1x authentication key feature is enabled.

	IsAuthKeySaeExt *bool `json:"isAuthKeySaeExt,omitempty"` // When set to true, the Simultaneous Authentication of Equals (SAE) Extended Authentication key feature is turned on.

	IsAuthKeySaeExtPlusFT *bool `json:"isAuthKeySaeExtPlusFT,omitempty"` // When set to true, the Simultaneous Authentication of Equals (SAE) combined with Fast Transition (FT) Authentication Key feature is enabled.

	IsApBeaconProtectionEnabled *bool `json:"isApBeaconProtectionEnabled,omitempty"` // When set to true, the Access Point (AP) Beacon Protection feature is activated, enhancing the security of the network.

	Ghz24Policy string `json:"ghz24Policy,omitempty"` // 2.4 Ghz Band Policy value. Allowed only when 2.4 Radio Band is enabled in ssidRadioType

	CckmTsfTolerance *int `json:"cckmTsfTolerance,omitempty"` // Cckm TImestamp Tolerance(in milliseconds)

	IsCckmEnabled *bool `json:"isCckmEnabled,omitempty"` // True if CCKM is enabled, else False

	IsHex *bool `json:"isHex,omitempty"` // True if passphrase is in Hex format, else False.

	IsSensorPnp *bool `json:"isSensorPnp,omitempty"` // True if SSID is a sensor SSID

	ID string `json:"id,omitempty"` // SSID ID

	IsRandomMacFilterEnabled *bool `json:"isRandomMacFilterEnabled,omitempty"` // Deny clients using randomized MAC addresses when set to true

	FastTransitionOverTheDistributedSystemEnable *bool `json:"fastTransitionOverTheDistributedSystemEnable,omitempty"` // Enable Fast Transition over the Distributed System when set to true

	InheritedSiteNameHierarchy string `json:"inheritedSiteNameHierarchy,omitempty"` // Inherited Site Name Hierarchy

	InheritedSiteUUID string `json:"inheritedSiteUUID,omitempty"` // Inherited Site UUID
}
type ResponseWirelessGetSSIDByIDV1ResponseMultipSKSettings struct {
	Priority *int `json:"priority,omitempty"` // Priority

	PassphraseType string `json:"passphraseType,omitempty"` // Passphrase Type

	Passphrase string `json:"passphrase,omitempty"` // Passphrase needs to be between 8 and 63 characters for ASCII type. HEX passphrase needs to be 64 characters
}
type ResponseWirelessUpdateSSIDV1 struct {
	Response *ResponseWirelessUpdateSSIDV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessUpdateSSIDV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessDeleteSSIDV1 struct {
	Response *ResponseWirelessDeleteSSIDV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessDeleteSSIDV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessUpdateOrOverridessidV1 struct {
	Response *ResponseWirelessUpdateOrOverridessidV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessUpdateOrOverridessidV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessDeleteWirelessProfileV1 struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url

	Message string `json:"message,omitempty"` // Message
}
type ResponseWirelessConfigureAccessPointsV1 struct {
	Response *ResponseWirelessConfigureAccessPointsV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseWirelessConfigureAccessPointsV1Response struct {
	TaskID string `json:"taskId,omitempty"` //

	URL string `json:"url,omitempty"` //
}
type ResponseWirelessGetAccessPointConfigurationCountV1 struct {
	Response *ResponseWirelessGetAccessPointConfigurationCountV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response Version
}
type ResponseWirelessGetAccessPointConfigurationCountV1Response struct {
	Count *int `json:"count,omitempty"` // Count of the requested resource
}
type ResponseWirelessGetAccessPointConfigurationTaskResultV1 []ResponseItemWirelessGetAccessPointConfigurationTaskResultV1 // Array of ResponseWirelessGetAccessPointConfigurationTaskResultV1
type ResponseItemWirelessGetAccessPointConfigurationTaskResultV1 struct {
	InstanceUUID *ResponseItemWirelessGetAccessPointConfigurationTaskResultV1InstanceUUID `json:"instanceUuid,omitempty"` //

	InstanceID *float64 `json:"instanceId,omitempty"` //

	AuthEntityID *ResponseItemWirelessGetAccessPointConfigurationTaskResultV1AuthEntityID `json:"authEntityId,omitempty"` //

	DisplayName string `json:"displayName,omitempty"` //

	AuthEntityClass *ResponseItemWirelessGetAccessPointConfigurationTaskResultV1AuthEntityClass `json:"authEntityClass,omitempty"` //

	InstanceTenantID string `json:"instanceTenantId,omitempty"` //

	OrderedListOEIndex *float64 `json:"_orderedListOEIndex,omitempty"` //

	OrderedListOEAssocName *ResponseItemWirelessGetAccessPointConfigurationTaskResultV1OrderedListOEAssocName `json:"_orderedListOEAssocName,omitempty"` //

	CreationOrderIndex *float64 `json:"_creationOrderIndex,omitempty"` //

	IsBeingChanged *bool `json:"_isBeingChanged,omitempty"` //

	DeployPending string `json:"deployPending,omitempty"` //

	InstanceCreatedOn *ResponseItemWirelessGetAccessPointConfigurationTaskResultV1InstanceCreatedOn `json:"instanceCreatedOn,omitempty"` //

	InstanceUpdatedOn *ResponseItemWirelessGetAccessPointConfigurationTaskResultV1InstanceUpdatedOn `json:"instanceUpdatedOn,omitempty"` //

	ChangeLogList *ResponseItemWirelessGetAccessPointConfigurationTaskResultV1ChangeLogList `json:"changeLogList,omitempty"` //

	InstanceOrigin *ResponseItemWirelessGetAccessPointConfigurationTaskResultV1InstanceOrigin `json:"instanceOrigin,omitempty"` //

	LazyLoadedEntities *ResponseItemWirelessGetAccessPointConfigurationTaskResultV1LazyLoadedEntities `json:"lazyLoadedEntities,omitempty"` //

	InstanceVersion *float64 `json:"instanceVersion,omitempty"` //

	ApName string `json:"apName,omitempty"` //

	ControllerName string `json:"controllerName,omitempty"` //

	LocationHeirarchy string `json:"locationHeirarchy,omitempty"` //

	MacAddress string `json:"macAddress,omitempty"` //

	Status string `json:"status,omitempty"` //

	StatusDetails string `json:"statusDetails,omitempty"` //

	InternalKey *ResponseItemWirelessGetAccessPointConfigurationTaskResultV1InternalKey `json:"internalKey,omitempty"` //
}
type ResponseItemWirelessGetAccessPointConfigurationTaskResultV1InstanceUUID interface{}
type ResponseItemWirelessGetAccessPointConfigurationTaskResultV1AuthEntityID interface{}
type ResponseItemWirelessGetAccessPointConfigurationTaskResultV1AuthEntityClass interface{}
type ResponseItemWirelessGetAccessPointConfigurationTaskResultV1OrderedListOEAssocName interface{}
type ResponseItemWirelessGetAccessPointConfigurationTaskResultV1InstanceCreatedOn interface{}
type ResponseItemWirelessGetAccessPointConfigurationTaskResultV1InstanceUpdatedOn interface{}
type ResponseItemWirelessGetAccessPointConfigurationTaskResultV1ChangeLogList interface{}
type ResponseItemWirelessGetAccessPointConfigurationTaskResultV1InstanceOrigin interface{}
type ResponseItemWirelessGetAccessPointConfigurationTaskResultV1LazyLoadedEntities interface{}
type ResponseItemWirelessGetAccessPointConfigurationTaskResultV1InternalKey struct {
	Type string `json:"type,omitempty"` //

	ID *float64 `json:"id,omitempty"` //

	LongType string `json:"longType,omitempty"` //

	URL string `json:"url,omitempty"` //
}
type ResponseWirelessGetAccessPointConfigurationV1 struct {
	InstanceUUID            *ResponseWirelessGetAccessPointConfigurationV1InstanceUUID           `json:"instanceUuid,omitempty"`            //
	InstanceID              *float64                                                             `json:"instanceId,omitempty"`              //
	AuthEntityID            *ResponseWirelessGetAccessPointConfigurationV1AuthEntityID           `json:"authEntityId,omitempty"`            //
	DisplayName             string                                                               `json:"displayName,omitempty"`             //
	AuthEntityClass         *ResponseWirelessGetAccessPointConfigurationV1AuthEntityClass        `json:"authEntityClass,omitempty"`         //
	InstanceTenantID        string                                                               `json:"instanceTenantId,omitempty"`        //
	OrderedListOEIndex      *float64                                                             `json:"_orderedListOEIndex,omitempty"`     //
	OrderedListOEAssocName  *ResponseWirelessGetAccessPointConfigurationV1OrderedListOEAssocName `json:"_orderedListOEAssocName,omitempty"` //
	CreationOrderIndex      *float64                                                             `json:"_creationOrderIndex,omitempty"`     //
	IsBeingChanged          *bool                                                                `json:"_isBeingChanged,omitempty"`         //
	DeployPending           string                                                               `json:"deployPending,omitempty"`           //
	InstanceCreatedOn       *ResponseWirelessGetAccessPointConfigurationV1InstanceCreatedOn      `json:"instanceCreatedOn,omitempty"`       //
	InstanceUpdatedOn       *ResponseWirelessGetAccessPointConfigurationV1InstanceUpdatedOn      `json:"instanceUpdatedOn,omitempty"`       //
	ChangeLogList           *ResponseWirelessGetAccessPointConfigurationV1ChangeLogList          `json:"changeLogList,omitempty"`           //
	InstanceOrigin          *ResponseWirelessGetAccessPointConfigurationV1InstanceOrigin         `json:"instanceOrigin,omitempty"`          //
	LazyLoadedEntities      *ResponseWirelessGetAccessPointConfigurationV1LazyLoadedEntities     `json:"lazyLoadedEntities,omitempty"`      //
	InstanceVersion         *float64                                                             `json:"instanceVersion,omitempty"`         //
	AdminStatus             string                                                               `json:"adminStatus,omitempty"`             //
	ApHeight                *float64                                                             `json:"apHeight,omitempty"`                //
	ApMode                  string                                                               `json:"apMode,omitempty"`                  //
	ApName                  string                                                               `json:"apName,omitempty"`                  //
	EthMac                  string                                                               `json:"ethMac,omitempty"`                  //
	FailoverPriority        string                                                               `json:"failoverPriority,omitempty"`        //
	LedBrightnessLevel      *int                                                                 `json:"ledBrightnessLevel,omitempty"`      //
	LedStatus               string                                                               `json:"ledStatus,omitempty"`               //
	Location                string                                                               `json:"location,omitempty"`                //
	MacAddress              string                                                               `json:"macAddress,omitempty"`              //
	PrimaryControllerName   string                                                               `json:"primaryControllerName,omitempty"`   //
	PrimaryIPAddress        string                                                               `json:"primaryIpAddress,omitempty"`        //
	SecondaryControllerName string                                                               `json:"secondaryControllerName,omitempty"` //
	SecondaryIPAddress      string                                                               `json:"secondaryIpAddress,omitempty"`      //
	TertiaryControllerName  string                                                               `json:"tertiaryControllerName,omitempty"`  //
	TertiaryIPAddress       string                                                               `json:"tertiaryIpAddress,omitempty"`       //
	MeshDTOs                *[]ResponseWirelessGetAccessPointConfigurationV1MeshDTOs             `json:"meshDTOs,omitempty"`                //
	RadioDTOs               *[]ResponseWirelessGetAccessPointConfigurationV1RadioDTOs            `json:"radioDTOs,omitempty"`               //
	InternalKey             *ResponseWirelessGetAccessPointConfigurationV1InternalKey            `json:"internalKey,omitempty"`             //
}
type ResponseWirelessGetAccessPointConfigurationV1InstanceUUID interface{}
type ResponseWirelessGetAccessPointConfigurationV1AuthEntityID interface{}
type ResponseWirelessGetAccessPointConfigurationV1AuthEntityClass interface{}
type ResponseWirelessGetAccessPointConfigurationV1OrderedListOEAssocName interface{}
type ResponseWirelessGetAccessPointConfigurationV1InstanceCreatedOn interface{}
type ResponseWirelessGetAccessPointConfigurationV1InstanceUpdatedOn interface{}
type ResponseWirelessGetAccessPointConfigurationV1ChangeLogList interface{}
type ResponseWirelessGetAccessPointConfigurationV1InstanceOrigin interface{}
type ResponseWirelessGetAccessPointConfigurationV1LazyLoadedEntities interface{}
type ResponseWirelessGetAccessPointConfigurationV1MeshDTOs interface{}
type ResponseWirelessGetAccessPointConfigurationV1RadioDTOs struct {
	InstanceUUID           *ResponseWirelessGetAccessPointConfigurationV1RadioDTOsInstanceUUID           `json:"instanceUuid,omitempty"`            //
	InstanceID             *float64                                                                      `json:"instanceId,omitempty"`              //
	AuthEntityID           *ResponseWirelessGetAccessPointConfigurationV1RadioDTOsAuthEntityID           `json:"authEntityId,omitempty"`            //
	DisplayName            string                                                                        `json:"displayName,omitempty"`             //
	AuthEntityClass        *ResponseWirelessGetAccessPointConfigurationV1RadioDTOsAuthEntityClass        `json:"authEntityClass,omitempty"`         //
	InstanceTenantID       string                                                                        `json:"instanceTenantId,omitempty"`        //
	OrderedListOEIndex     *float64                                                                      `json:"_orderedListOEIndex,omitempty"`     //
	OrderedListOEAssocName *ResponseWirelessGetAccessPointConfigurationV1RadioDTOsOrderedListOEAssocName `json:"_orderedListOEAssocName,omitempty"` //
	CreationOrderIndex     *float64                                                                      `json:"_creationOrderIndex,omitempty"`     //
	IsBeingChanged         *bool                                                                         `json:"_isBeingChanged,omitempty"`         //
	DeployPending          string                                                                        `json:"deployPending,omitempty"`           //
	InstanceCreatedOn      *ResponseWirelessGetAccessPointConfigurationV1RadioDTOsInstanceCreatedOn      `json:"instanceCreatedOn,omitempty"`       //
	InstanceUpdatedOn      *ResponseWirelessGetAccessPointConfigurationV1RadioDTOsInstanceUpdatedOn      `json:"instanceUpdatedOn,omitempty"`       //
	ChangeLogList          *ResponseWirelessGetAccessPointConfigurationV1RadioDTOsChangeLogList          `json:"changeLogList,omitempty"`           //
	InstanceOrigin         *ResponseWirelessGetAccessPointConfigurationV1RadioDTOsInstanceOrigin         `json:"instanceOrigin,omitempty"`          //
	LazyLoadedEntities     *ResponseWirelessGetAccessPointConfigurationV1RadioDTOsLazyLoadedEntities     `json:"lazyLoadedEntities,omitempty"`      //
	InstanceVersion        *float64                                                                      `json:"instanceVersion,omitempty"`         //
	AdminStatus            string                                                                        `json:"adminStatus,omitempty"`             //
	AntennaAngle           *float64                                                                      `json:"antennaAngle,omitempty"`            //
	AntennaElevAngle       *float64                                                                      `json:"antennaElevAngle,omitempty"`        //
	AntennaGain            *int                                                                          `json:"antennaGain,omitempty"`             //
	AntennaPatternName     string                                                                        `json:"antennaPatternName,omitempty"`      //
	ChannelAssignmentMode  string                                                                        `json:"channelAssignmentMode,omitempty"`   //
	ChannelNumber          *int                                                                          `json:"channelNumber,omitempty"`           //
	ChannelWidth           string                                                                        `json:"channelWidth,omitempty"`            //
	CleanAirSI             string                                                                        `json:"cleanAirSI,omitempty"`              //
	IfType                 *int                                                                          `json:"ifType,omitempty"`                  //
	IfTypeValue            string                                                                        `json:"ifTypeValue,omitempty"`             //
	MacAddress             string                                                                        `json:"macAddress,omitempty"`              //
	PowerAssignmentMode    string                                                                        `json:"powerAssignmentMode,omitempty"`     //
	Powerlevel             *int                                                                          `json:"powerlevel,omitempty"`              //
	RadioBand              *ResponseWirelessGetAccessPointConfigurationV1RadioDTOsRadioBand              `json:"radioBand,omitempty"`               //
	RadioRoleAssignment    *ResponseWirelessGetAccessPointConfigurationV1RadioDTOsRadioRoleAssignment    `json:"radioRoleAssignment,omitempty"`     //
	SlotID                 *int                                                                          `json:"slotId,omitempty"`                  //
	InternalKey            *ResponseWirelessGetAccessPointConfigurationV1RadioDTOsInternalKey            `json:"internalKey,omitempty"`             //
}
type ResponseWirelessGetAccessPointConfigurationV1RadioDTOsInstanceUUID interface{}
type ResponseWirelessGetAccessPointConfigurationV1RadioDTOsAuthEntityID interface{}
type ResponseWirelessGetAccessPointConfigurationV1RadioDTOsAuthEntityClass interface{}
type ResponseWirelessGetAccessPointConfigurationV1RadioDTOsOrderedListOEAssocName interface{}
type ResponseWirelessGetAccessPointConfigurationV1RadioDTOsInstanceCreatedOn interface{}
type ResponseWirelessGetAccessPointConfigurationV1RadioDTOsInstanceUpdatedOn interface{}
type ResponseWirelessGetAccessPointConfigurationV1RadioDTOsChangeLogList interface{}
type ResponseWirelessGetAccessPointConfigurationV1RadioDTOsInstanceOrigin interface{}
type ResponseWirelessGetAccessPointConfigurationV1RadioDTOsLazyLoadedEntities interface{}
type ResponseWirelessGetAccessPointConfigurationV1RadioDTOsRadioBand interface{}
type ResponseWirelessGetAccessPointConfigurationV1RadioDTOsRadioRoleAssignment interface{}
type ResponseWirelessGetAccessPointConfigurationV1RadioDTOsInternalKey struct {
	Type     string   `json:"type,omitempty"`     //
	ID       *float64 `json:"id,omitempty"`       //
	LongType string   `json:"longType,omitempty"` //
	URL      string   `json:"url,omitempty"`      //
}
type ResponseWirelessGetAccessPointConfigurationV1InternalKey struct {
	Type     string   `json:"type,omitempty"`     //
	ID       *float64 `json:"id,omitempty"`       //
	LongType string   `json:"longType,omitempty"` //
	URL      string   `json:"url,omitempty"`      //
}
type ResponseWirelessApProvisionConnectivityV1 struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status URL

	Message string `json:"message,omitempty"` // Message
}
type ResponseWirelessDeleteDynamicInterfaceV1 struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status URL

	Message string `json:"message,omitempty"` // Message
}
type ResponseWirelessCreateUpdateDynamicInterfaceV1 struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status URL

	Message string `json:"message,omitempty"` // Message
}
type ResponseWirelessGetDynamicInterfaceV1 []ResponseItemWirelessGetDynamicInterfaceV1 // Array of ResponseWirelessGetDynamicInterfaceV1
type ResponseItemWirelessGetDynamicInterfaceV1 struct {
	InterfaceName string `json:"interfaceName,omitempty"` // dynamic interface name

	VLANID *float64 `json:"vlanId,omitempty"` // Vlan id
}
type ResponseWirelessUpdateWirelessProfileV1 struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url

	Message string `json:"message,omitempty"` // Message
}
type ResponseWirelessCreateWirelessProfileV1 struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url

	Message string `json:"message,omitempty"` // Message
}
type ResponseWirelessGetWirelessProfileV1 []ResponseItemWirelessGetWirelessProfileV1 // Array of ResponseWirelessGetWirelessProfileV1
type ResponseItemWirelessGetWirelessProfileV1 struct {
	ProfileDetails *ResponseItemWirelessGetWirelessProfileV1ProfileDetails `json:"profileDetails,omitempty"` //
}
type ResponseItemWirelessGetWirelessProfileV1ProfileDetails struct {
	Name string `json:"name,omitempty"` // Profile Name

	Sites []string `json:"sites,omitempty"` // array of site name hierarchies(eg: ["Global/aaa/zzz", "Global/aaa/zzz"])

	SSIDDetails *[]ResponseItemWirelessGetWirelessProfileV1ProfileDetailsSSIDDetails `json:"ssidDetails,omitempty"` //
}
type ResponseItemWirelessGetWirelessProfileV1ProfileDetailsSSIDDetails struct {
	Name string `json:"name,omitempty"` // SSID Name

	Type string `json:"type,omitempty"` // SSID Type(enum: Enterprise/Guest)

	EnableFabric *bool `json:"enableFabric,omitempty"` // true if fabric is enabled else false

	FlexConnect *ResponseItemWirelessGetWirelessProfileV1ProfileDetailsSSIDDetailsFlexConnect `json:"flexConnect,omitempty"` //

	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name

	WLANProfileName string `json:"wlanProfileName,omitempty"` // WLAN Profile Name

	PolicyProfileName string `json:"policyProfileName,omitempty"` // Policy Profile Name
}

func (r *ResponseItemWirelessGetWirelessProfileV1ProfileDetailsSSIDDetails) UnmarshalJSON(data []byte) error {
	type Alias ResponseItemWirelessGetWirelessProfileV1ProfileDetailsSSIDDetails
	aux := &struct {
		EnableFabric interface{} `json:"enableFabric"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	switch v := aux.EnableFabric.(type) {
	case bool:
		r.EnableFabric = &v
	case string:
		if v == "true" {
			r.EnableFabric = new(bool)
			*r.EnableFabric = true
		} else if v == "false" {
			r.EnableFabric = new(bool)
			*r.EnableFabric = false
		} else {
			r.EnableFabric = nil
		}
	case nil:
		r.EnableFabric = nil
	default:
		r.EnableFabric = nil
	}

	return nil
}

type ResponseItemWirelessGetWirelessProfileV1ProfileDetailsSSIDDetailsFlexConnect struct {
	EnableFlexConnect *bool `json:"enableFlexConnect,omitempty"` // true if flex connect is enabled else false

	LocalToVLAN *int `json:"localToVlan,omitempty"` // Local To VLAN ID
}
type ResponseWirelessProvisionUpdateV1 struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url

	Message string `json:"message,omitempty"` // Message
}
type ResponseWirelessProvisionV1 struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url

	Message string `json:"message,omitempty"` // Message
}
type ResponseWirelessPSKOverrideV1 struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url

	Message string `json:"message,omitempty"` // Message
}
type ResponseWirelessRetrieveRfProfilesV1 struct {
	Name string `json:"name,omitempty"` // RF Profile Name

	DefaultRfProfile *bool `json:"defaultRfProfile,omitempty"` // is Default Rf Profile

	EnableRadioTypeA *bool `json:"enableRadioTypeA,omitempty"` // Enable Radio Type A

	EnableRadioTypeB *bool `json:"enableRadioTypeB,omitempty"` // Enable Radio Type B

	ChannelWidth string `json:"channelWidth,omitempty"` // Channel Width

	EnableCustom *bool `json:"enableCustom,omitempty"` // Enable Custom

	EnableBrownField *bool `json:"enableBrownField,omitempty"` // Enable Brown Field

	RadioTypeAProperties *ResponseWirelessRetrieveRfProfilesV1RadioTypeAProperties `json:"radioTypeAProperties,omitempty"` //

	RadioTypeBProperties *ResponseWirelessRetrieveRfProfilesV1RadioTypeBProperties `json:"radioTypeBProperties,omitempty"` //

	RadioTypeCProperties *ResponseWirelessRetrieveRfProfilesV1RadioTypeCProperties `json:"radioTypeCProperties,omitempty"` //

	EnableRadioTypeC *bool `json:"enableRadioTypeC,omitempty"` // Enable Radio Type C (6GHz)
}
type ResponseWirelessRetrieveRfProfilesV1RadioTypeAProperties struct {
	ParentProfile string `json:"parentProfile,omitempty"` // Parent Profile (Default : CUSTOM)

	RadioChannels string `json:"radioChannels,omitempty"` // Radio Channels (Default : "36,40,44,48,52,56,60,64,149,153,157,161")

	DataRates string `json:"dataRates,omitempty"` // Data Rates (Default : "6,9,12,18,24,36,48,54")

	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory Data Rates (Default: "6,12,24")

	PowerThresholdV1 *float64 `json:"powerThresholdV1,omitempty"` // Power Threshold V1 ( (Default: -70)

	RxSopThreshold string `json:"rxSopThreshold,omitempty"` // Rx Sop Threshold  (Default: "AUTO")

	MinPowerLevel *float64 `json:"minPowerLevel,omitempty"` // Rx Sop Threshold  (Default: -10)

	MaxPowerLevel *float64 `json:"maxPowerLevel,omitempty"` // Max Power Level  (Default: 30)
}
type ResponseWirelessRetrieveRfProfilesV1RadioTypeBProperties struct {
	ParentProfile string `json:"parentProfile,omitempty"` // Parent Profile (Default : CUSTOM)

	RadioChannels string `json:"radioChannels,omitempty"` // Radio Channels (Default : "9,11,12,18,24,36,48,54")

	DataRates string `json:"dataRates,omitempty"` // Data Rates  (Default: "9,11,12,18,24,36,48,54")

	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory Data Rates  (Default: "12")

	PowerThresholdV1 *float64 `json:"powerThresholdV1,omitempty"` // Power Threshold V1  (Default: -70)

	RxSopThreshold string `json:"rxSopThreshold,omitempty"` // Rx Sop Threshold (Default: "AUTO")

	MinPowerLevel *float64 `json:"minPowerLevel,omitempty"` // Min Power Level  (Default: -10)

	MaxPowerLevel *float64 `json:"maxPowerLevel,omitempty"` // Max Power Level  (Default: 30)
}
type ResponseWirelessRetrieveRfProfilesV1RadioTypeCProperties struct {
	ParentProfile string `json:"parentProfile,omitempty"` // Parent Profile (Default : CUSTOM)

	RadioChannels string `json:"radioChannels,omitempty"` // Radio Channels (Default : "5,21,37,53,69,85,101,117,133,149,165,181,197,213,229")

	DataRates string `json:"dataRates,omitempty"` // Data Rates  (Default: "6,9,12,18,24,36,48,54")

	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory Data Rates  (Default: "6,12,24")

	RxSopThreshold string `json:"rxSopThreshold,omitempty"` // Rx Sop Threshold  (Default: "AUTO")

	MinPowerLevel *float64 `json:"minPowerLevel,omitempty"` // Min Power Level  (Default: -10)

	MaxPowerLevel *float64 `json:"maxPowerLevel,omitempty"` // Max Power Level  (Default: 30)

	PowerThresholdV1 *float64 `json:"powerThresholdV1,omitempty"` // Power Threshold V1  (Default: -70)
}
type ResponseWirelessCreateOrUpdateRfProfileV1 struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url

	Message string `json:"message,omitempty"` // Message
}
type ResponseWirelessDeleteRfProfilesV1 struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url

	Message string `json:"message,omitempty"` // Message
}
type ResponseWirelessFactoryResetAccessPointsV1 struct {
	Response *ResponseWirelessFactoryResetAccessPointsV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessFactoryResetAccessPointsV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	URL string `json:"url,omitempty"` // Url
}
type ResponseWirelessGetAccessPointsFactoryResetStatusV1 struct {
	Response *[]ResponseWirelessGetAccessPointsFactoryResetStatusV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetAccessPointsFactoryResetStatusV1Response struct {
	WlcIP string `json:"wlcIP,omitempty"` // Wireless Controller IP address

	WlcName string `json:"wlcName,omitempty"` // Wireless Controller name

	ApResponseInfoList *[]ResponseWirelessGetAccessPointsFactoryResetStatusV1ResponseApResponseInfoList `json:"apResponseInfoList,omitempty"` //
}
type ResponseWirelessGetAccessPointsFactoryResetStatusV1ResponseApResponseInfoList struct {
	ApName string `json:"apName,omitempty"` // Access Point name

	ApFactoryResetStatus string `json:"apFactoryResetStatus,omitempty"` // AP factory reset status, "Success" or "Failure" or "In Progress"

	FailureReason string `json:"failureReason,omitempty"` // Reason for failure if the factory reset status is "Failure"

	RadioMacAddress string `json:"radioMacAddress,omitempty"` // AP Radio Mac Address

	EthernetMacAddress string `json:"ethernetMacAddress,omitempty"` // AP Ethernet Mac Address
}
type ResponseWirelessApProvisionV1 struct {
	Response *ResponseWirelessApProvisionV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessApProvisionV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessGetAnchorCapableDevicesV1 struct {
	DeviceIP string `json:"deviceIp,omitempty"` // Anchor Controller Ip

	DeviceName string `json:"deviceName,omitempty"` // Anchor Controller host name

	WirelessMgmtIP string `json:"wirelessMgmtIP,omitempty"` // Wireless management Ip Address
}
type ResponseWirelessGetMeshApNeighboursV1 struct {
	ID string `json:"id,omitempty"` // id

	ApName string `json:"apName,omitempty"` // Name of the Wireless Access point

	EthernetMacAddress string `json:"ethernetMacAddress,omitempty"` // AP Ethernet MacAddress mac

	NeighbourMacAddress string `json:"neighbourMacAddress,omitempty"` // AP Base Radio MacAddress mac.

	WlcIPAddress string `json:"wlcIpAddress,omitempty"` // Device wireless Management IP

	NeighbourType string `json:"neighbourType,omitempty"` // Neighbour Type

	MeshRole string `json:"meshRole,omitempty"` // Mesh Role
}
type ResponseWirelessGetMeshApNeighboursCountV1 struct {
	Response *ResponseWirelessGetMeshApNeighboursCountV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response Version
}
type ResponseWirelessGetMeshApNeighboursCountV1Response struct {
	Count *int `json:"count,omitempty"` // Count of the requested resource
}
type ResponseWirelessGetMobilityGroupsV1 struct {
	Response *[]ResponseWirelessGetMobilityGroupsV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response version.
}
type ResponseWirelessGetMobilityGroupsV1Response struct {
	MobilityGroupName string `json:"mobilityGroupName,omitempty"` // Self device Group Name. Must be alphanumeric without {!,<,space,?/'} and maximum of 31 characters.

	MacAddress string `json:"macAddress,omitempty"` // Device mobility MAC Address. Allowed formats are:0a0b.0c01.0211, 0a0b0c010211, 0a:0b:0c:01:02:11

	ManagementIP string `json:"managementIp,omitempty"` // Self device wireless Management IP.

	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Obtain the network device ID value by using the API call GET: /dna/intent/api/v1/network-device/ip-address/${ipAddress}.

	DtlsHighCipher *bool `json:"dtlsHighCipher,omitempty"` // DTLS High Cipher.

	DataLinkEncryption *bool `json:"dataLinkEncryption,omitempty"` // A secure link in which data is encrypted using CAPWAP DTLS protocol can be established between two controllers. This value will be applied to all peers during POST operation.

	MobilityPeers *[]ResponseWirelessGetMobilityGroupsV1ResponseMobilityPeers `json:"mobilityPeers,omitempty"` //
}
type ResponseWirelessGetMobilityGroupsV1ResponseMobilityPeers struct {
	MobilityGroupName string `json:"mobilityGroupName,omitempty"` // Peer device mobility group Name. Must be alphanumeric without {!,<,space,?/'} and maximum of 31 characters.

	PeerNetworkDeviceID string `json:"peerNetworkDeviceId,omitempty"` // Peer device Id. The possible values are UNKNOWN or valid UUID of Network device ID. UNKNOWN represents out of band device which is not managed internally. Valid UUID represents WLC network device ID.

	MemberMacAddress string `json:"memberMacAddress,omitempty"` // Peer device mobility MAC Address.  Allowed formats are:0a0b.0c01.0211, 0a0b0c010211, 0a:0b:0c:01:02:11

	DeviceSeries string `json:"deviceSeries,omitempty"` // Peer device mobility belongs to AireOS or IOX-XE family. 0 - indicates AireOS and 1 - indicates C9800.

	DataLinkEncryption *bool `json:"dataLinkEncryption,omitempty"` // A secure link in which data is encrypted using CAPWAP DTLS protocol can be established between two controllers.

	HashKey string `json:"hashKey,omitempty"` // SSC hash string must be 40 characters.

	Status string `json:"status,omitempty"` // Possible values are - Control and Data Path Down, Data Path Down, Control Path Down, UP.

	PeerIP string `json:"peerIp,omitempty"` // This indicates public IP address.

	PrivateIPAddress string `json:"privateIpAddress,omitempty"` // This indicates private/management IP address.
}
type ResponseWirelessGetMobilityGroupsCountV1 struct {
	Response *ResponseWirelessGetMobilityGroupsCountV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response version.
}
type ResponseWirelessGetMobilityGroupsCountV1Response struct {
	Count *int `json:"count,omitempty"` // Total number of mobility groups available.
}
type ResponseWirelessMobilityProvisionV1 struct {
	Response *ResponseWirelessMobilityProvisionV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessMobilityProvisionV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Asynchronous Task Id

	URL string `json:"url,omitempty"` // Asynchronous Task URL for further tracking
}
type ResponseWirelessMobilityResetV1 struct {
	Response *ResponseWirelessMobilityResetV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessMobilityResetV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Asynchronous Task Id

	URL string `json:"url,omitempty"` // Asynchronous Task URL for further tracking
}
type ResponseWirelessAssignManagedApLocationsForWLCV1 struct {
	Response *ResponseWirelessAssignManagedApLocationsForWLCV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessAssignManagedApLocationsForWLCV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessWirelessControllerProvisionV1 struct {
	Response *ResponseWirelessWirelessControllerProvisionV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessWirelessControllerProvisionV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessGetAnchorManagedApLocationsForSpecificWirelessControllerV1 struct {
	Response *[]ResponseWirelessGetAnchorManagedApLocationsForSpecificWirelessControllerV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetAnchorManagedApLocationsForSpecificWirelessControllerV1Response struct {
	ManagedApLocations *[]ResponseWirelessGetAnchorManagedApLocationsForSpecificWirelessControllerV1ResponseManagedApLocations `json:"managedApLocations,omitempty"` //
}
type ResponseWirelessGetAnchorManagedApLocationsForSpecificWirelessControllerV1ResponseManagedApLocations struct {
	SiteID string `json:"siteId,omitempty"` // The site id of the managed ap location.

	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // The site name hierarchy of the managed ap location.
}
type ResponseWirelessGetApAuthorizationListByNetworkDeviceIDV1 struct {
	Response *ResponseWirelessGetApAuthorizationListByNetworkDeviceIDV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetApAuthorizationListByNetworkDeviceIDV1Response struct {
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Network Device Id

	ApAuthorizationListName string `json:"apAuthorizationListName,omitempty"` // Ap Authorization List Name

	LocalAuthorization *ResponseWirelessGetApAuthorizationListByNetworkDeviceIDV1ResponseLocalAuthorization `json:"localAuthorization,omitempty"` //

	RemoteAuthorization *ResponseWirelessGetApAuthorizationListByNetworkDeviceIDV1ResponseRemoteAuthorization `json:"remoteAuthorization,omitempty"` //
}
type ResponseWirelessGetApAuthorizationListByNetworkDeviceIDV1ResponseLocalAuthorization struct {
	ApMacEntries []string `json:"apMacEntries,omitempty"` // Ap Mac Entries

	ApSerialNumberEntries []string `json:"apSerialNumberEntries,omitempty"` // Ap Serial Number Entries
}
type ResponseWirelessGetApAuthorizationListByNetworkDeviceIDV1ResponseRemoteAuthorization struct {
	AAAServers []string `json:"aaaServers,omitempty"` // Aaa Servers

	AuthorizeApWithMac *bool `json:"authorizeApWithMac,omitempty"` // Authorize Ap With Mac

	AuthorizeApWithSerialNumber *bool `json:"authorizeApWithSerialNumber,omitempty"` // Authorize Ap With Serial Number
}
type ResponseWirelessGetManagedApLocationsCountForSpecificWirelessControllerV1 struct {
	Response *ResponseWirelessGetManagedApLocationsCountForSpecificWirelessControllerV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetManagedApLocationsCountForSpecificWirelessControllerV1Response struct {
	PrimaryManagedApLocationsCount *int `json:"primaryManagedApLocationsCount,omitempty"` // The count of the Primary managed ap locations.

	SecondaryManagedApLocationsCount *int `json:"secondaryManagedApLocationsCount,omitempty"` // The count of the Secondary managed ap locations.

	AnchorManagedApLocationsCount *int `json:"anchorManagedApLocationsCount,omitempty"` // The count of the Anchor managed ap  locations.
}
type ResponseWirelessGetPrimaryManagedApLocationsForSpecificWirelessControllerV1 struct {
	Response *[]ResponseWirelessGetPrimaryManagedApLocationsForSpecificWirelessControllerV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetPrimaryManagedApLocationsForSpecificWirelessControllerV1Response struct {
	ManagedApLocations *[]ResponseWirelessGetPrimaryManagedApLocationsForSpecificWirelessControllerV1ResponseManagedApLocations `json:"managedApLocations,omitempty"` //
}
type ResponseWirelessGetPrimaryManagedApLocationsForSpecificWirelessControllerV1ResponseManagedApLocations struct {
	SiteID string `json:"siteId,omitempty"` // The site id of the managed ap location.

	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // The site name hierarchy of the managed ap location.
}
type ResponseWirelessGetSecondaryManagedApLocationsForSpecificWirelessControllerV1 struct {
	Response *[]ResponseWirelessGetSecondaryManagedApLocationsForSpecificWirelessControllerV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetSecondaryManagedApLocationsForSpecificWirelessControllerV1Response struct {
	ManagedApLocations *[]ResponseWirelessGetSecondaryManagedApLocationsForSpecificWirelessControllerV1ResponseManagedApLocations `json:"managedApLocations,omitempty"` //
}
type ResponseWirelessGetSecondaryManagedApLocationsForSpecificWirelessControllerV1ResponseManagedApLocations struct {
	SiteID string `json:"siteId,omitempty"` // The site id of the managed ap location.

	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // The site name hierarchy of the managed ap location.
}
type ResponseWirelessGetSSIDDetailsForSpecificWirelessControllerV1 struct {
	Response *[]ResponseWirelessGetSSIDDetailsForSpecificWirelessControllerV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetSSIDDetailsForSpecificWirelessControllerV1Response struct {
	SSIDName string `json:"ssidName,omitempty"` // Name of the SSID.

	WLANID *int `json:"wlanId,omitempty"` // WLAN ID.

	WLANProfileName string `json:"wlanProfileName,omitempty"` // WLAN Profile Name.

	L2Security string `json:"l2Security,omitempty"` // This represents the identifier for the Layer 2 authentication type. The authentication types supported include wpa2_enterprise, wpa2_personal, open, wpa3_enterprise, wpa3_personal, wpa2_wpa3_personal, wpa2_wpa3_enterprise, and open-secured.

	L3Security string `json:"l3Security,omitempty"` // This represents the identifier for the Layer 3 authentication type. The authentication types supported are 'open' and 'webauth'.

	RadioPolicy string `json:"radioPolicy,omitempty"` // This represents the identifier for the radio policy. The policies supported include 2.4GHz, 5GHz, and 6GHz.

	AdminStatus *bool `json:"adminStatus,omitempty"` // Utilize this query parameter to obtain the administrative status. A 'true' value signifies that the admin status of the SSID is enabled, while a 'false' value indicates that the admin status of the SSID is disabled.

	Managed *bool `json:"managed,omitempty"` // If the value is 'true,' the SSID is configured through design; if 'false,' it indicates out-of-band configuration on the Wireless LAN Controller.
}
type ResponseWirelessGetSSIDCountForSpecificWirelessControllerV1 struct {
	Response *ResponseWirelessGetSSIDCountForSpecificWirelessControllerV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetSSIDCountForSpecificWirelessControllerV1Response struct {
	Count *int `json:"count,omitempty"` // The count of the SSIDs.
}
type ResponseWirelessGetWirelessProfilesV1 struct {
	Response *[]ResponseWirelessGetWirelessProfilesV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetWirelessProfilesV1Response struct {
	WirelessProfileName string `json:"wirelessProfileName,omitempty"` // Wireless Profile Name

	SSIDDetails *[]ResponseWirelessGetWirelessProfilesV1ResponseSSIDDetails `json:"ssidDetails,omitempty"` //

	ID string `json:"id,omitempty"` // Wireless Profile Id

	AdditionalInterfaces []string `json:"additionalInterfaces,omitempty"` // Additional Interfaces

	ApZones *[]ResponseWirelessGetWirelessProfilesV1ResponseApZones `json:"apZones,omitempty"` //
}
type ResponseWirelessGetWirelessProfilesV1ResponseSSIDDetails struct {
	SSIDName string `json:"ssidName,omitempty"` // SSID Name

	FlexConnect *ResponseWirelessGetWirelessProfilesV1ResponseSSIDDetailsFlexConnect `json:"flexConnect,omitempty"` //

	EnableFabric *bool `json:"enableFabric,omitempty"` // True if fabric is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time

	WLANProfileName string `json:"wlanProfileName,omitempty"` // WLAN Profile Name

	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name

	PolicyProfileName string `json:"policyProfileName,omitempty"` // Policy Profile Name

	Dot11BeProfileID string `json:"dot11beProfileId,omitempty"` // 802.11be Profile ID

	AnchorGroupName string `json:"anchorGroupName,omitempty"` // Anchor Group Name

	VLANGroupName string `json:"vlanGroupName,omitempty"` // VLAN Group Name
}
type ResponseWirelessGetWirelessProfilesV1ResponseSSIDDetailsFlexConnect struct {
	EnableFlexConnect *bool `json:"enableFlexConnect,omitempty"` // True if flex connect is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time

	LocalToVLAN *int `json:"localToVlan,omitempty"` // Local to VLAN ID
}
type ResponseWirelessGetWirelessProfilesV1ResponseApZones struct {
	ApZoneName string `json:"apZoneName,omitempty"` // AP Zone Name

	RfProfileName string `json:"rfProfileName,omitempty"` // RF Profile Name

	SSIDs []string `json:"ssids,omitempty"` // ssids part of apZone
}
type ResponseWirelessCreateWirelessProfileConnectivityV1 struct {
	Response *ResponseWirelessCreateWirelessProfileConnectivityV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessCreateWirelessProfileConnectivityV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessGetWirelessProfilesCountV1 struct {
	Response *ResponseWirelessGetWirelessProfilesCountV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response Version
}
type ResponseWirelessGetWirelessProfilesCountV1Response struct {
	Count *int `json:"count,omitempty"` // Count of the requested resource
}
type ResponseWirelessUpdateWirelessProfileConnectivityV1 struct {
	Response *ResponseWirelessUpdateWirelessProfileConnectivityV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessUpdateWirelessProfileConnectivityV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessGetWirelessProfileByIDV1 struct {
	Response *ResponseWirelessGetWirelessProfileByIDV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetWirelessProfileByIDV1Response struct {
	WirelessProfileName string `json:"wirelessProfileName,omitempty"` // Wireless Profile Name

	SSIDDetails *[]ResponseWirelessGetWirelessProfileByIDV1ResponseSSIDDetails `json:"ssidDetails,omitempty"` //

	ID string `json:"id,omitempty"` // Wireless Profile Id

	AdditionalInterfaces []string `json:"additionalInterfaces,omitempty"` // Additional Interfaces

	ApZones *[]ResponseWirelessGetWirelessProfileByIDV1ResponseApZones `json:"apZones,omitempty"` //
}
type ResponseWirelessGetWirelessProfileByIDV1ResponseSSIDDetails struct {
	SSIDName string `json:"ssidName,omitempty"` // SSID Name

	FlexConnect *ResponseWirelessGetWirelessProfileByIDV1ResponseSSIDDetailsFlexConnect `json:"flexConnect,omitempty"` //

	EnableFabric *bool `json:"enableFabric,omitempty"` // True if fabric is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time

	WLANProfileName string `json:"wlanProfileName,omitempty"` // WLAN Profile Name

	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name

	PolicyProfileName string `json:"policyProfileName,omitempty"` // Policy Profile Name

	Dot11BeProfileID string `json:"dot11beProfileId,omitempty"` // 802.11be Profile ID

	AnchorGroupName string `json:"anchorGroupName,omitempty"` // Anchor Group Name

	VLANGroupName string `json:"vlanGroupName,omitempty"` // VLAN Group Name
}
type ResponseWirelessGetWirelessProfileByIDV1ResponseSSIDDetailsFlexConnect struct {
	EnableFlexConnect *bool `json:"enableFlexConnect,omitempty"` // True if flex connect is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time

	LocalToVLAN *int `json:"localToVlan,omitempty"` // Local to VLAN ID
}
type ResponseWirelessGetWirelessProfileByIDV1ResponseApZones struct {
	ApZoneName string `json:"apZoneName,omitempty"` // AP Zone Name

	RfProfileName string `json:"rfProfileName,omitempty"` // RF Profile Name

	SSIDs []string `json:"ssids,omitempty"` // ssids part of apZone
}
type ResponseWirelessDeleteWirelessProfileConnectivityV1 struct {
	Response *ResponseWirelessDeleteWirelessProfileConnectivityV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessDeleteWirelessProfileConnectivityV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessRetrieveAllPolicyTagsForAWirelessProfileV1 struct {
	Response *[]ResponseWirelessRetrieveAllPolicyTagsForAWirelessProfileV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessRetrieveAllPolicyTagsForAWirelessProfileV1Response struct {
	SiteIDs []string `json:"siteIds,omitempty"` // Site Ids

	PolicyTagName string `json:"policyTagName,omitempty"` // Use English letters, numbers, special characters except <, /, '.*', ? and leading/trailing space.

	ApZones []string `json:"apZones,omitempty"` // Ap Zones

	PolicyTagID string `json:"policyTagId,omitempty"` // Policy Tag Id
}
type ResponseWirelessCreateMultiplePolicyTagsForAWirelessProfileInBulkV1 struct {
	Response *ResponseWirelessCreateMultiplePolicyTagsForAWirelessProfileInBulkV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessCreateMultiplePolicyTagsForAWirelessProfileInBulkV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessRetrieveTheCountOfPolicyTagsForAWirelessProfileV1 struct {
	Response *ResponseWirelessRetrieveTheCountOfPolicyTagsForAWirelessProfileV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response Version
}
type ResponseWirelessRetrieveTheCountOfPolicyTagsForAWirelessProfileV1Response struct {
	Count *int `json:"count,omitempty"` // Count of the requested resource
}
type ResponseWirelessDeleteASpecificPolicyTagFromAWirelessProfileV1 struct {
	Response *ResponseWirelessDeleteASpecificPolicyTagFromAWirelessProfileV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessDeleteASpecificPolicyTagFromAWirelessProfileV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessUpdateASpecificPolicyTagForAWirelessProfileV1 struct {
	Response *ResponseWirelessUpdateASpecificPolicyTagForAWirelessProfileV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessUpdateASpecificPolicyTagForAWirelessProfileV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessRetrieveASpecificPolicyTagForAWirelessProfileV1 struct {
	Response *ResponseWirelessRetrieveASpecificPolicyTagForAWirelessProfileV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessRetrieveASpecificPolicyTagForAWirelessProfileV1Response struct {
	SiteIDs []string `json:"siteIds,omitempty"` // Site Ids

	PolicyTagName string `json:"policyTagName,omitempty"` // Policy Tag Name

	ApZones []string `json:"apZones,omitempty"` // Ap Zones

	PolicyTagID string `json:"policyTagId,omitempty"` // Policy Tag Id
}
type ResponseWirelessRetrieveAllSiteTagsForAWirelessProfileV1 struct {
	Response *[]ResponseWirelessRetrieveAllSiteTagsForAWirelessProfileV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessRetrieveAllSiteTagsForAWirelessProfileV1Response struct {
	SiteIDs []string `json:"siteIds,omitempty"` // Site Ids

	SiteTagName string `json:"siteTagName,omitempty"` // Site Tag Name

	FlexProfileName string `json:"flexProfileName,omitempty"` // Flex Profile Name

	ApProfileName string `json:"apProfileName,omitempty"` // Ap Profile Name

	SiteTagID string `json:"siteTagId,omitempty"` // Site Tag Id
}
type ResponseWirelessCreateMultipleSiteTagsForAWirelessProfileInBulkV1 struct {
	Response *ResponseWirelessCreateMultipleSiteTagsForAWirelessProfileInBulkV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessCreateMultipleSiteTagsForAWirelessProfileInBulkV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessRetrieveTheCountOfSiteTagsForAWirelessProfileV1 struct {
	Response *ResponseWirelessRetrieveTheCountOfSiteTagsForAWirelessProfileV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response Version
}
type ResponseWirelessRetrieveTheCountOfSiteTagsForAWirelessProfileV1Response struct {
	Count *int `json:"count,omitempty"` // Count of the requested resource
}
type ResponseWirelessUpdateASpecificSiteTagForAWirelessProfileV1 struct {
	Response *ResponseWirelessUpdateASpecificSiteTagForAWirelessProfileV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessUpdateASpecificSiteTagForAWirelessProfileV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessRetrieveASpecificSiteTagForAWirelessProfileV1 struct {
	Response *ResponseWirelessRetrieveASpecificSiteTagForAWirelessProfileV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessRetrieveASpecificSiteTagForAWirelessProfileV1Response struct {
	SiteIDs []string `json:"siteIds,omitempty"` // Site Ids

	SiteTagName string `json:"siteTagName,omitempty"` // Use English letters, numbers, special characters except <, /, '.*', ? and leading/trailing space

	FlexProfileName string `json:"flexProfileName,omitempty"` // Flex Profile Name

	ApProfileName string `json:"apProfileName,omitempty"` // Ap Profile Name

	SiteTagID string `json:"siteTagId,omitempty"` // Site Tag Id
}
type ResponseWirelessDeleteASpecificSiteTagFromAWirelessProfileV1 struct {
	Response *ResponseWirelessDeleteASpecificSiteTagFromAWirelessProfileV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessDeleteASpecificSiteTagFromAWirelessProfileV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessCreateAnchorGroupV1 struct {
	Response *ResponseWirelessCreateAnchorGroupV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessCreateAnchorGroupV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessGetAnchorGroupsV1 struct {
	ID string `json:"id,omitempty"` // Anchor Profile unique ID

	AnchorGroupName string `json:"anchorGroupName,omitempty"` // Anchor Group Name. Max length is 32 characters

	MobilityAnchors *[]ResponseWirelessGetAnchorGroupsV1MobilityAnchors `json:"mobilityAnchors,omitempty"` //
}
type ResponseWirelessGetAnchorGroupsV1MobilityAnchors struct {
	DeviceName string `json:"deviceName,omitempty"` // Peer Host Name

	IPAddress string `json:"ipAddress,omitempty"` // This indicates Mobility public IP address

	AnchorPriority string `json:"anchorPriority,omitempty"` // This indicates anchor priority.  Priority values range from 1 (high) to 3 (low). Primary, secondary or tertiary and defined priority is displayed with guest anchor. Only one priority value is allowed per anchor WLC.

	ManagedAnchorWlc *bool `json:"managedAnchorWlc,omitempty"` // This indicates whether the Wireless LAN Controller supporting Anchor is managed by the Network Controller or not. True means this is managed by Network Controller.

	PeerDeviceType string `json:"peerDeviceType,omitempty"` // Indicates peer device mobility belongs to AireOS or IOS-XE family.

	MacAddress string `json:"macAddress,omitempty"` // Peer Device mobility MAC address. Allowed formats are:0a0b.0c01.0211, 0a0b0c010211, 0a:0b:0c:01:02:11

	MobilityGroupName string `json:"mobilityGroupName,omitempty"` // Peer Device mobility group Name. Must be alphanumeric without {!,<,space,?/'} and maximum of 31 characters.

	PrivateIP string `json:"privateIp,omitempty"` // This indicates private management IP address
}
type ResponseWirelessGetCountOfAnchorGroupsV1 struct {
	Response *ResponseWirelessGetCountOfAnchorGroupsV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response Version
}
type ResponseWirelessGetCountOfAnchorGroupsV1Response struct {
	Count *int `json:"count,omitempty"` // Count of the requested resource
}
type ResponseWirelessGetAnchorGroupByIDV1 struct {
	ID string `json:"id,omitempty"` // Anchor Profile unique ID

	AnchorGroupName string `json:"anchorGroupName,omitempty"` // Anchor Group Name. Max length is 32 characters

	MobilityAnchors *[]ResponseWirelessGetAnchorGroupByIDV1MobilityAnchors `json:"mobilityAnchors,omitempty"` //
}
type ResponseWirelessGetAnchorGroupByIDV1MobilityAnchors struct {
	DeviceName string `json:"deviceName,omitempty"` // Peer Host Name

	IPAddress string `json:"ipAddress,omitempty"` // This indicates Mobility public IP address

	AnchorPriority string `json:"anchorPriority,omitempty"` // This indicates anchor priority.  Priority values range from 1 (high) to 3 (low). Primary, secondary or tertiary and defined priority is displayed with guest anchor. Only one priority value is allowed per anchor WLC.

	ManagedAnchorWlc *bool `json:"managedAnchorWlc,omitempty"` // This indicates whether the Wireless LAN Controller supporting Anchor is managed by the Network Controller or not. True means this is managed by Network Controller.

	PeerDeviceType string `json:"peerDeviceType,omitempty"` // Indicates peer device mobility belongs to AireOS or IOS-XE family.

	MacAddress string `json:"macAddress,omitempty"` // Peer Device mobility MAC address. Allowed formats are:0a0b.0c01.0211, 0a0b0c010211, 0a:0b:0c:01:02:11

	MobilityGroupName string `json:"mobilityGroupName,omitempty"` // Peer Device mobility group Name. Must be alphanumeric without {!,<,space,?/'} and maximum of 31 characters.

	PrivateIP string `json:"privateIp,omitempty"` // This indicates private management IP address
}
type ResponseWirelessDeleteAnchorGroupByIDV1 struct {
	Response *ResponseWirelessDeleteAnchorGroupByIDV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessDeleteAnchorGroupByIDV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessUpdateAnchorGroupV1 struct {
	Response *ResponseWirelessUpdateAnchorGroupV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessUpdateAnchorGroupV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessGetApAuthorizationListsV1 struct {
	Response *ResponseWirelessGetApAuthorizationListsV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetApAuthorizationListsV1Response struct {
	ID string `json:"id,omitempty"` // Id

	ApAuthorizationListName string `json:"apAuthorizationListName,omitempty"` // Ap Authorization List Name

	LocalAuthorization *ResponseWirelessGetApAuthorizationListsV1ResponseLocalAuthorization `json:"localAuthorization,omitempty"` //

	RemoteAuthorization *ResponseWirelessGetApAuthorizationListsV1ResponseRemoteAuthorization `json:"remoteAuthorization,omitempty"` //
}
type ResponseWirelessGetApAuthorizationListsV1ResponseLocalAuthorization struct {
	ApMacEntries []string `json:"apMacEntries,omitempty"` // AP Mac Addresses

	ApSerialNumberEntries []string `json:"apSerialNumberEntries,omitempty"` // AP Serial Number Entries
}
type ResponseWirelessGetApAuthorizationListsV1ResponseRemoteAuthorization struct {
	AAAServers []string `json:"aaaServers,omitempty"` // AAA Servers

	AuthorizeApWithMac *bool `json:"authorizeApWithMac,omitempty"` // Authorize AP With Mac

	AuthorizeApWithSerialNumber *bool `json:"authorizeApWithSerialNumber,omitempty"` // Authorize AP With Serial Number
}
type ResponseWirelessCreateApAuthorizationListV1 struct {
	Response *ResponseWirelessCreateApAuthorizationListV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessCreateApAuthorizationListV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessGetApAuthorizationListCountV1 struct {
	Response *ResponseWirelessGetApAuthorizationListCountV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response Version
}
type ResponseWirelessGetApAuthorizationListCountV1Response struct {
	Count *int `json:"count,omitempty"` // Count of the requested resource
}
type ResponseWirelessDeleteApAuthorizationListV1 struct {
	Response *ResponseWirelessDeleteApAuthorizationListV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessDeleteApAuthorizationListV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessUpdateApAuthorizationListV1 struct {
	Response *ResponseWirelessUpdateApAuthorizationListV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessUpdateApAuthorizationListV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessGetApAuthorizationListByIDV1 struct {
	Response *ResponseWirelessGetApAuthorizationListByIDV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetApAuthorizationListByIDV1Response struct {
	ID string `json:"id,omitempty"` // Id

	ApAuthorizationListName string `json:"apAuthorizationListName,omitempty"` // Ap Authorization List Name

	LocalAuthorization *ResponseWirelessGetApAuthorizationListByIDV1ResponseLocalAuthorization `json:"localAuthorization,omitempty"` //

	RemoteAuthorization *ResponseWirelessGetApAuthorizationListByIDV1ResponseRemoteAuthorization `json:"remoteAuthorization,omitempty"` //
}
type ResponseWirelessGetApAuthorizationListByIDV1ResponseLocalAuthorization struct {
	ApMacEntries []string `json:"apMacEntries,omitempty"` // AP Mac Addresses

	ApSerialNumberEntries []string `json:"apSerialNumberEntries,omitempty"` // AP Serial Number Entries
}
type ResponseWirelessGetApAuthorizationListByIDV1ResponseRemoteAuthorization struct {
	AAAServers []string `json:"aaaServers,omitempty"` // AAA Servers

	AuthorizeApWithMac *bool `json:"authorizeApWithMac,omitempty"` // Authorize AP With Mac

	AuthorizeApWithSerialNumber *bool `json:"authorizeApWithSerialNumber,omitempty"` // Authorize AP With Serial Number
}
type ResponseWirelessCreateApProfileV1 struct {
	Response *ResponseWirelessCreateApProfileV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessCreateApProfileV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessGetApProfilesV1 struct {
	Response *[]ResponseWirelessGetApProfilesV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetApProfilesV1Response struct {
	ID string `json:"id,omitempty"` // AP Profile unique ID

	ApProfileName string `json:"apProfileName,omitempty"` // Name of the Access Point profile. Max length is 32 characters.

	Description string `json:"description,omitempty"` // Description of the AP profile. Max length is 241 characters

	RemoteWorkerEnabled *bool `json:"remoteWorkerEnabled,omitempty"` // Indicates if remote worker mode is enabled on the AP. Remote teleworker enabled profile cannot support security features like aWIPS,Forensic Capture Enablement, Rogue Detection and Rogue Containment.

	ManagementSetting *ResponseWirelessGetApProfilesV1ResponseManagementSetting `json:"managementSetting,omitempty"` //

	AwipsEnabled *bool `json:"awipsEnabled,omitempty"` // Indicates if AWIPS is enabled on the AP.

	AwipsForensicEnabled *bool `json:"awipsForensicEnabled,omitempty"` // Indicates if AWIPS forensic is enabled on the AP. Forensic Capture is supported from IOS-XE version 17.4 and above. Forensic Capture can be activated only if aWIPS is enabled.

	RogueDetectionSetting *ResponseWirelessGetApProfilesV1ResponseRogueDetectionSetting `json:"rogueDetectionSetting,omitempty"` //

	PmfDenialEnabled *bool `json:"pmfDenialEnabled,omitempty"` // Indicates if PMF denial is active on the AP. PMF Denial is supported from IOS-XE version 17.12 and above.

	MeshEnabled *bool `json:"meshEnabled,omitempty"` // This indicates whether mesh networking is enabled on the AP. For IOS-XE devices, when mesh networking is enabled, a custom mesh profile with the configured parameters will be created and mapped to the AP join profile on the device. When mesh networking is disabled, any existing custom mesh profile will be deleted from the device, and the AP join profile will be mapped to the default mesh profile on the device.

	MeshSetting *ResponseWirelessGetApProfilesV1ResponseMeshSetting `json:"meshSetting,omitempty"` //

	ApPowerProfileName string `json:"apPowerProfileName,omitempty"` // Name of the existing AP power profile.

	CalendarPowerProfiles *ResponseWirelessGetApProfilesV1ResponseCalendarPowerProfiles `json:"calendarPowerProfiles,omitempty"` //

	CountryCode string `json:"countryCode,omitempty"` // Country Code

	TimeZone string `json:"timeZone,omitempty"` // Time zone of the AP.

	TimeZoneOffsetHour *int `json:"timeZoneOffsetHour,omitempty"` // Hour 'Delta from Controller' for the time zone. The value should be between -12 and 14.

	TimeZoneOffsetMinutes *int `json:"timeZoneOffsetMinutes,omitempty"` // Minute 'Delta from Controller' for the time zone. Value should be between 0 to 59.

	ClientLimit *int `json:"clientLimit,omitempty"` // Number of clients. Value should be between 0-1200.
}
type ResponseWirelessGetApProfilesV1ResponseManagementSetting struct {
	AuthType string `json:"authType,omitempty"` // Authentication type used in the AP profile. These setting are applicable during PnP claim and for day-N authentication of AP. Changing these settings will be service impacting for the PnP onboarded APs and will need a factory-reset for those APs.

	Dot1XUsername string `json:"dot1xUsername,omitempty"` // Username for 802.1X authentication. dot1xUsername must have a minimum of 1 character and a maximum of 32 characters.

	Dot1XPassword string `json:"dot1xPassword,omitempty"` // Password for 802.1X authentication. AP dot1x password length should not exceed 120.

	SSHEnabled *bool `json:"sshEnabled,omitempty"` // Indicates if SSH is enabled on the AP. Enable SSH add credentials for device management.

	TelnetEnabled *bool `json:"telnetEnabled,omitempty"` // Indicates if Telnet is enabled on the AP. Enable Telnet to add credentials for device management.

	ManagementUserName string `json:"managementUserName,omitempty"` // Management username must have a minimum of 1 character and a maximum of 32 characters.

	ManagementPassword string `json:"managementPassword,omitempty"` // Management password for the AP. Length must be 8-120 characters.

	ManagementEnablePassword string `json:"managementEnablePassword,omitempty"` // Enable password for managing the AP. Length must be 8-120 characters.

	CdpState *bool `json:"cdpState,omitempty"` // Indicates if CDP is enabled on the AP. Enable CDP in order to make Cisco Access Points known to its neighboring devices and vice-versa.
}
type ResponseWirelessGetApProfilesV1ResponseRogueDetectionSetting struct {
	RogueDetection *bool `json:"rogueDetection,omitempty"` // Indicates if rogue detection is enabled. Detect Access Points that have been installed on a secure network without explicit authorization from a system administrator and configure rogue general configuration parameters

	RogueDetectionMinRssi *int `json:"rogueDetectionMinRssi,omitempty"` // Minimum RSSI for rogue detection. Value should be in range -128 decibel milliwatts and -70 decibel milliwatts

	RogueDetectionTransientInterval *int `json:"rogueDetectionTransientInterval,omitempty"` // Transient interval for rogue detection. Value should be 0 or from 120 to 1800.

	RogueDetectionReportInterval *int `json:"rogueDetectionReportInterval,omitempty"` // Report interval for rogue detection. Value should be in range 10 and 300.
}
type ResponseWirelessGetApProfilesV1ResponseMeshSetting struct {
	BridgeGroupName string `json:"bridgeGroupName,omitempty"` // Name of the bridge group for mesh settings. If not configured, 'Default' Bridge group name will be used in mesh profile.

	BackhaulClientAccess *bool `json:"backhaulClientAccess,omitempty"` // Indicates if backhaul client access is enabled on the AP.

	Range *int `json:"range,omitempty"` // Range of the mesh network. Value should be between 150-132000

	Ghz5BackhaulDataRates string `json:"ghz5BackhaulDataRates,omitempty"` // 5GHz backhaul data rates.

	Ghz24BackhaulDataRates string `json:"ghz24BackhaulDataRates,omitempty"` // 2.4GHz backhaul data rates.

	RapDownlinkBackhaul string `json:"rapDownlinkBackhaul,omitempty"` // Type of downlink backhaul used.
}
type ResponseWirelessGetApProfilesV1ResponseCalendarPowerProfiles struct {
	PowerProfileName string `json:"powerProfileName,omitempty"` // Name of the existing AP power profile to be mapped to the calendar power profile. The following API is used create AP power profile. API-/intent/api/v1/wirelessSettings/powerProfiles

	SchedulerType string `json:"schedulerType,omitempty"` // Type of the scheduler.

	Duration *ResponseWirelessGetApProfilesV1ResponseCalendarPowerProfilesDuration `json:"duration,omitempty"` //
}
type ResponseWirelessGetApProfilesV1ResponseCalendarPowerProfilesDuration struct {
	SchedulerStartTime string `json:"schedulerStartTime,omitempty"` // Start time of the duration setting.

	SchedulerEndTime string `json:"schedulerEndTime,omitempty"` // End time of the duration setting.

	SchedulerDay string `json:"schedulerDay,omitempty"` // Applies every week on the selected days

	SchedulerDate string `json:"schedulerDate,omitempty"` // Start and End date of the duration setting, applicable for MONTHLY schedulers.
}
type ResponseWirelessGetApProfilesCountV1 struct {
	Response *ResponseWirelessGetApProfilesCountV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response Version
}
type ResponseWirelessGetApProfilesCountV1Response struct {
	Count *int `json:"count,omitempty"` // Count of the requested resource
}
type ResponseWirelessDeleteApProfileByIDV1 struct {
	Response *ResponseWirelessDeleteApProfileByIDV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessDeleteApProfileByIDV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessUpdateApProfileByIDV1 struct {
	Response *ResponseWirelessUpdateApProfileByIDV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessUpdateApProfileByIDV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessGetApProfileByIDV1 struct {
	Response *[]ResponseWirelessGetApProfileByIDV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetApProfileByIDV1Response struct {
	ID string `json:"id,omitempty"` // AP Profile unique ID

	ApProfileName string `json:"apProfileName,omitempty"` // Name of the Access Point profile. Max length is 32 characters.

	Description string `json:"description,omitempty"` // Description of the AP profile. Max length is 241 characters

	RemoteWorkerEnabled *bool `json:"remoteWorkerEnabled,omitempty"` // Indicates if remote worker mode is enabled on the AP. Remote teleworker enabled profile cannot support security features like aWIPS,Forensic Capture Enablement, Rogue Detection and Rogue Containment.

	ManagementSetting *ResponseWirelessGetApProfileByIDV1ResponseManagementSetting `json:"managementSetting,omitempty"` //

	AwipsEnabled *bool `json:"awipsEnabled,omitempty"` // Indicates if AWIPS is enabled on the AP.

	AwipsForensicEnabled *bool `json:"awipsForensicEnabled,omitempty"` // Indicates if AWIPS forensic is enabled on the AP. Forensic Capture is supported from IOS-XE version 17.4 and above. Forensic Capture can be activated only if aWIPS is enabled.

	RogueDetectionSetting *ResponseWirelessGetApProfileByIDV1ResponseRogueDetectionSetting `json:"rogueDetectionSetting,omitempty"` //

	PmfDenialEnabled *bool `json:"pmfDenialEnabled,omitempty"` // Indicates if PMF denial is active on the AP. PMF Denial is supported from IOS-XE version 17.12 and above.

	MeshEnabled *bool `json:"meshEnabled,omitempty"` // This indicates whether mesh networking is enabled on the AP. For IOS-XE devices, when mesh networking is enabled, a custom mesh profile with the configured parameters will be created and mapped to the AP join profile on the device. When mesh networking is disabled, any existing custom mesh profile will be deleted from the device, and the AP join profile will be mapped to the default mesh profile on the device.

	MeshSetting *ResponseWirelessGetApProfileByIDV1ResponseMeshSetting `json:"meshSetting,omitempty"` //

	ApPowerProfileName string `json:"apPowerProfileName,omitempty"` // Name of the existing AP power profile.

	CalendarPowerProfiles *ResponseWirelessGetApProfileByIDV1ResponseCalendarPowerProfiles `json:"calendarPowerProfiles,omitempty"` //

	CountryCode string `json:"countryCode,omitempty"` // Country Code

	TimeZone string `json:"timeZone,omitempty"` // In the Time Zone area, choose one of the following options.             Not Configured - APs operate in the UTC time zone.             Controller - APs operate in the Cisco Wireless Controller time zone.             Delta from Controller - APs operate in the offset time from the wireless controller time zone.

	TimeZoneOffsetHour *int `json:"timeZoneOffsetHour,omitempty"` // Enter the hour value (HH). The valid range is from -12 through 14.

	TimeZoneOffsetMinutes *int `json:"timeZoneOffsetMinutes,omitempty"` // Enter the minute value (MM). The valid range is from 0 through 59.

	ClientLimit *int `json:"clientLimit,omitempty"` // Number of clients. Value should be between 0-1200.
}
type ResponseWirelessGetApProfileByIDV1ResponseManagementSetting struct {
	AuthType string `json:"authType,omitempty"` // Authentication type used in the AP profile. These setting are applicable during PnP claim and for day-N authentication of AP. Changing these settings will be service impacting for the PnP onboarded APs and will need a factory-reset for those APs.

	Dot1XUsername string `json:"dot1xUsername,omitempty"` // Username for 802.1X authentication. dot1xUsername must have a minimum of 1 character and a maximum of 32 characters.

	Dot1XPassword string `json:"dot1xPassword,omitempty"` // Password for 802.1X authentication. AP dot1x password length should not exceed 120.

	SSHEnabled *bool `json:"sshEnabled,omitempty"` // Indicates if SSH is enabled on the AP. Enable SSH add credentials for device management.

	TelnetEnabled *bool `json:"telnetEnabled,omitempty"` // Indicates if Telnet is enabled on the AP. Enable Telnet to add credentials for device management.

	ManagementUserName string `json:"managementUserName,omitempty"` // Management username must have a minimum of 1 character and a maximum of 32 characters.

	ManagementPassword string `json:"managementPassword,omitempty"` // Management password for the AP. Length must be 8-120 characters.

	ManagementEnablePassword string `json:"managementEnablePassword,omitempty"` // Enable password for managing the AP. Length must be 8-120 characters.

	CdpState *bool `json:"cdpState,omitempty"` // Indicates if CDP is enabled on the AP. Enable CDP in order to make Cisco Access Points known to its neighboring devices and vice-versa.
}
type ResponseWirelessGetApProfileByIDV1ResponseRogueDetectionSetting struct {
	RogueDetection *bool `json:"rogueDetection,omitempty"` // Indicates if rogue detection is enabled on the AP. Detect Access Points that have been installed on a secure network without explicit authorization from a system administrator and configure rogue general configuration parameters

	RogueDetectionMinRssi *int `json:"rogueDetectionMinRssi,omitempty"` // Minimum RSSI for rogue detection. Value should be in range -128 decibel milliwatts and -70 decibel milliwatts

	RogueDetectionTransientInterval *int `json:"rogueDetectionTransientInterval,omitempty"` // Transient interval for rogue detection. Value should be 0 or from 120 to 1800.

	RogueDetectionReportInterval *int `json:"rogueDetectionReportInterval,omitempty"` // Report interval for rogue detection. Value should be in range 10 and 300.
}
type ResponseWirelessGetApProfileByIDV1ResponseMeshSetting struct {
	BridgeGroupName string `json:"bridgeGroupName,omitempty"` // Name of the bridge group for mesh settings. If not configured, 'Default' Bridge group name will be used in mesh profile.

	BackhaulClientAccess *bool `json:"backhaulClientAccess,omitempty"` // Indicates if backhaul client access is enabled on the AP.

	Range *int `json:"range,omitempty"` // Range of the mesh network. Value should be between 150-132000

	Ghz5BackhaulDataRates string `json:"ghz5BackhaulDataRates,omitempty"` // 5GHz backhaul data rates.

	Ghz24BackhaulDataRates string `json:"ghz24BackhaulDataRates,omitempty"` // 2.4GHz backhaul data rates.

	RapDownlinkBackhaul string `json:"rapDownlinkBackhaul,omitempty"` // Type of downlink backhaul used.
}
type ResponseWirelessGetApProfileByIDV1ResponseCalendarPowerProfiles struct {
	PowerProfileName string `json:"powerProfileName,omitempty"` // Name of the existing AP power profile to be mapped to the calendar power profile. API-/intent/api/v1/wirelessSettings/powerProfiles.

	SchedulerType string `json:"schedulerType,omitempty"` // Type of the scheduler.

	Duration *ResponseWirelessGetApProfileByIDV1ResponseCalendarPowerProfilesDuration `json:"duration,omitempty"` //
}
type ResponseWirelessGetApProfileByIDV1ResponseCalendarPowerProfilesDuration struct {
	SchedulerStartTime string `json:"schedulerStartTime,omitempty"` // Start time of the duration setting.

	SchedulerEndTime string `json:"schedulerEndTime,omitempty"` // End time of the duration setting.

	SchedulerDay string `json:"schedulerDay,omitempty"` // Applies every week on the selected days

	SchedulerDate string `json:"schedulerDate,omitempty"` // Start and End date of the duration setting, applicable for MONTHLY schedulers.
}
type ResponseWirelessGet80211BeProfilesV1 struct {
	Response *[]ResponseWirelessGet80211BeProfilesV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response Version
}
type ResponseWirelessGet80211BeProfilesV1Response struct {
	ID string `json:"id,omitempty"` // 802.11be Profile ID

	ProfileName string `json:"profileName,omitempty"` // 802.11be Profile Name

	OfdmaDownLink *bool `json:"ofdmaDownLink,omitempty"` // OFDMA Downlink

	OfdmaUpLink *bool `json:"ofdmaUpLink,omitempty"` // OFDMA Uplink

	MuMimoDownLink *bool `json:"muMimoDownLink,omitempty"` // MU-MIMO Downlink

	MuMimoUpLink *bool `json:"muMimoUpLink,omitempty"` // MU-MIMO Uplink

	OfdmaMultiRu *bool `json:"ofdmaMultiRu,omitempty"` // OFDMA Multi-RU

	Default *bool `json:"default,omitempty"` // 802.11be Profile is marked default or custom (Read only field)
}
type ResponseWirelessCreateA80211BeProfileV1 struct {
	Response *ResponseWirelessCreateA80211BeProfileV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessCreateA80211BeProfileV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessGet80211BeProfilesCountV1 struct {
	Response *ResponseWirelessGet80211BeProfilesCountV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response Version
}
type ResponseWirelessGet80211BeProfilesCountV1Response struct {
	Count *int `json:"count,omitempty"` // Count of the requested resource
}
type ResponseWirelessDeleteA80211BeProfileV1 struct {
	Response *ResponseWirelessDeleteA80211BeProfileV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessDeleteA80211BeProfileV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessUpdate80211BeProfileV1 struct {
	Response *ResponseWirelessUpdate80211BeProfileV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessUpdate80211BeProfileV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessGet80211BeProfileByIDV1 struct {
	Response *ResponseWirelessGet80211BeProfileByIDV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response Version
}
type ResponseWirelessGet80211BeProfileByIDV1Response struct {
	ID string `json:"id,omitempty"` // 802.11be Profile ID

	ProfileName string `json:"profileName,omitempty"` // 802.11be Profile Name

	OfdmaDownLink *bool `json:"ofdmaDownLink,omitempty"` // OFDMA Downlink

	OfdmaUpLink *bool `json:"ofdmaUpLink,omitempty"` // OFDMA Uplink

	MuMimoDownLink *bool `json:"muMimoDownLink,omitempty"` // MU-MIMO Downlink

	MuMimoUpLink *bool `json:"muMimoUpLink,omitempty"` // MU-MIMO Uplink

	OfdmaMultiRu *bool `json:"ofdmaMultiRu,omitempty"` // OFDMA Multi-RU

	Default *bool `json:"default,omitempty"` // Is 802.11be Profile marked as default in System . (Read only field)
}
type ResponseWirelessGetInterfacesV1 struct {
	Response *[]ResponseWirelessGetInterfacesV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetInterfacesV1Response struct {
	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name

	VLANID *int `json:"vlanId,omitempty"` // VLAN ID

	ID string `json:"id,omitempty"` // Interface ID
}
type ResponseWirelessCreateInterfaceV1 struct {
	Response *ResponseWirelessCreateInterfaceV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessCreateInterfaceV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessGetInterfacesCountV1 struct {
	Response *ResponseWirelessGetInterfacesCountV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response Version
}
type ResponseWirelessGetInterfacesCountV1Response struct {
	Count *int `json:"count,omitempty"` // Count of the requested resource
}
type ResponseWirelessGetInterfaceByIDV1 struct {
	Response *ResponseWirelessGetInterfaceByIDV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetInterfaceByIDV1Response struct {
	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name

	VLANID *int `json:"vlanId,omitempty"` // VLAN ID

	ID string `json:"id,omitempty"` // Interface ID
}
type ResponseWirelessDeleteInterfaceV1 struct {
	Response *ResponseWirelessDeleteInterfaceV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessDeleteInterfaceV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessUpdateInterfaceV1 struct {
	Response *ResponseWirelessUpdateInterfaceV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessUpdateInterfaceV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessCreatePowerProfileV1 struct {
	Response *ResponseWirelessCreatePowerProfileV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessCreatePowerProfileV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessGetPowerProfilesV1 struct {
	Response *[]ResponseWirelessGetPowerProfilesV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetPowerProfilesV1Response struct {
	ID string `json:"id,omitempty"` // Unique Identifier of the power profile.

	ProfileName string `json:"profileName,omitempty"` // The Name of the Power Profile.

	Description string `json:"description,omitempty"` // The description of the Power Profile.

	Rules *[]ResponseWirelessGetPowerProfilesV1ResponseRules `json:"rules,omitempty"` //
}
type ResponseWirelessGetPowerProfilesV1ResponseRules struct {
	Sequence *int `json:"sequence,omitempty"` // The sequence of the power profile rule.

	InterfaceType string `json:"interfaceType,omitempty"` // Interface Type for the rule.

	InterfaceID string `json:"interfaceId,omitempty"` // Interface Id for the rule.

	ParameterType string `json:"parameterType,omitempty"` // Parameter Type for the rule.

	ParameterValue string `json:"parameterValue,omitempty"` // Parameter Value for the rule.
}
type ResponseWirelessGetPowerProfilesCountV1 struct {
	Response *ResponseWirelessGetPowerProfilesCountV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response Version
}
type ResponseWirelessGetPowerProfilesCountV1Response struct {
	Count *int `json:"count,omitempty"` // Count of the requested resource
}
type ResponseWirelessDeletePowerProfileByIDV1 struct {
	Response *ResponseWirelessDeletePowerProfileByIDV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessDeletePowerProfileByIDV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessUpdatePowerProfileByIDV1 struct {
	Response *ResponseWirelessUpdatePowerProfileByIDV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessUpdatePowerProfileByIDV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessGetPowerProfileByIDV1 struct {
	Response *ResponseWirelessGetPowerProfileByIDV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetPowerProfileByIDV1Response struct {
	ID string `json:"id,omitempty"` // Unique Identifier of the power profile.

	ProfileName string `json:"profileName,omitempty"` // The Name of the Power Profile.

	Description string `json:"description,omitempty"` // The description of the Power Profile.

	Rules *[]ResponseWirelessGetPowerProfileByIDV1ResponseRules `json:"rules,omitempty"` //
}
type ResponseWirelessGetPowerProfileByIDV1ResponseRules struct {
	Sequence *int `json:"sequence,omitempty"` // Sequential Ordered List of rules for Power Profile.

	InterfaceType string `json:"interfaceType,omitempty"` // Interface Type for the rule.

	InterfaceID string `json:"interfaceId,omitempty"` // Interface Id for the rule.

	ParameterType string `json:"parameterType,omitempty"` // Parameter Type for the rule.

	ParameterValue string `json:"parameterValue,omitempty"` // Parameter Value for the rule.
}
type ResponseWirelessCreateRfProfileV1 struct {
	Response *ResponseWirelessCreateRfProfileV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessCreateRfProfileV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessGetRfProfilesV1 struct {
	Response *[]ResponseWirelessGetRfProfilesV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetRfProfilesV1Response struct {
	RfProfileName string `json:"rfProfileName,omitempty"` // RF Profile Name

	DefaultRfProfile *bool `json:"defaultRfProfile,omitempty"` // True if RF Profile is default, else False. Maximum of only 1 RF Profile can be marked as default at any given time

	EnableRadioTypeA *bool `json:"enableRadioTypeA,omitempty"` // True if 5 GHz radio band is enabled in the RF Profile, else False

	EnableRadioTypeB *bool `json:"enableRadioTypeB,omitempty"` // True if 2.4 GHz radio band is enabled in the RF Profile, else False

	EnableRadioType6GHz *bool `json:"enableRadioType6GHz,omitempty"` // True if 6 GHz radio band is enabled in the RF Profile, else False

	EnableCustom *bool `json:"enableCustom,omitempty"` // True if RF Profile is custom, else False for system RF profiles like Low, High and Medium (Typical)

	RadioTypeAProperties *ResponseWirelessGetRfProfilesV1ResponseRadioTypeAProperties `json:"radioTypeAProperties,omitempty"` //

	RadioTypeBProperties *ResponseWirelessGetRfProfilesV1ResponseRadioTypeBProperties `json:"radioTypeBProperties,omitempty"` //

	RadioType6GHzProperties *ResponseWirelessGetRfProfilesV1ResponseRadioType6GHzProperties `json:"radioType6GHzProperties,omitempty"` //

	ID string `json:"id,omitempty"` // RF Profile ID
}
type ResponseWirelessGetRfProfilesV1ResponseRadioTypeAProperties struct {
	ParentProfile string `json:"parentProfile,omitempty"` // Parent profile of 5 GHz radio band

	RadioChannels string `json:"radioChannels,omitempty"` // DCA channels of 5 GHz radio band passed in comma separated format without any spaces. Permissible values: 36, 40, 44, 48, 52, 56, 60, 64, 100, 104, 108, 112, 116, 120, 124, 128, 132, 136, 140, 144, 149, 153, 157, 161, 165, 169, 173

	DataRates string `json:"dataRates,omitempty"` // Data rates of 5 GHz radio band passed in comma separated format without any spaces. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54

	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory data rates of 5 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54

	PowerThresholdV1 *int `json:"powerThresholdV1,omitempty"` // Power threshold of 5 GHz radio band

	RxSopThreshold string `json:"rxSopThreshold,omitempty"` // RX-SOP threshold of 5 GHz radio band

	MinPowerLevel *int `json:"minPowerLevel,omitempty"` // Minimum power level of 5 GHz radio band

	MaxPowerLevel *int `json:"maxPowerLevel,omitempty"` // Maximum power level of 5 GHz radio band

	ChannelWidth string `json:"channelWidth,omitempty"` // Channel Width

	PreamblePuncture *bool `json:"preamblePuncture,omitempty"` // Enable or Disable Preamble Puncturing. This Wifi 7 configuration is applicable to wireless IOS devices supporting 17.15 and higher

	ZeroWaitDfsEnable *bool `json:"zeroWaitDfsEnable,omitempty"` // Zero Wait DFS is applicable only for IOS-XE based Wireless Controllers running 17.9.1 and above versions

	CustomRxSopThreshold *int `json:"customRxSopThreshold,omitempty"` // RX-SOP threshold custom configuration of 5 GHz radio band

	MaxRadioClients *int `json:"maxRadioClients,omitempty"` // Client Limit of 5 GHz radio band

	FraProperties *ResponseWirelessGetRfProfilesV1ResponseRadioTypeAPropertiesFraProperties `json:"fraProperties,omitempty"` //

	CoverageHoleDetectionProperties *ResponseWirelessGetRfProfilesV1ResponseRadioTypeAPropertiesCoverageHoleDetectionProperties `json:"coverageHoleDetectionProperties,omitempty"` //

	SpatialReuseProperties *ResponseWirelessGetRfProfilesV1ResponseRadioTypeAPropertiesSpatialReuseProperties `json:"spatialReuseProperties,omitempty"` //
}
type ResponseWirelessGetRfProfilesV1ResponseRadioTypeAPropertiesFraProperties struct {
	ClientAware *bool `json:"clientAware,omitempty"` // Client Aware of 5 GHz radio band

	ClientSelect *int `json:"clientSelect,omitempty"` // Client Select(%) of 5 GHz radio band

	ClientReset *int `json:"clientReset,omitempty"` // Client Reset(%) of 5 GHz radio band
}
type ResponseWirelessGetRfProfilesV1ResponseRadioTypeAPropertiesCoverageHoleDetectionProperties struct {
	ChdClientLevel *int `json:"chdClientLevel,omitempty"` // Coverage Hole Detection Client Level

	ChdDataRssiThreshold *int `json:"chdDataRssiThreshold,omitempty"` // Coverage Hole Detection Data Rssi Threshold

	ChdVoiceRssiThreshold *int `json:"chdVoiceRssiThreshold,omitempty"` // Coverage Hole Detection Voice Rssi Threshold

	ChdExceptionLevel *int `json:"chdExceptionLevel,omitempty"` // Coverage Hole Detection Exception Level(%)
}
type ResponseWirelessGetRfProfilesV1ResponseRadioTypeAPropertiesSpatialReuseProperties struct {
	Dot11AxNonSrgObssPacketDetect *bool `json:"dot11axNonSrgObssPacketDetect,omitempty"` // Dot11ax Non SRG OBSS PD

	Dot11AxNonSrgObssPacketDetectMaxThreshold *int `json:"dot11axNonSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax Non SRG OBSS PD Max Threshold

	Dot11AxSrgObssPacketDetect *bool `json:"dot11axSrgObssPacketDetect,omitempty"` // Dot11ax SRG OBSS PD

	Dot11AxSrgObssPacketDetectMinThreshold *int `json:"dot11axSrgObssPacketDetectMinThreshold,omitempty"` // Dot11ax SRG OBSS PD Min Threshold

	Dot11AxSrgObssPacketDetectMaxThreshold *int `json:"dot11axSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax SRG OBSS PD Max Threshold
}
type ResponseWirelessGetRfProfilesV1ResponseRadioTypeBProperties struct {
	ParentProfile string `json:"parentProfile,omitempty"` // Parent profile of 2.4 GHz radio band

	RadioChannels string `json:"radioChannels,omitempty"` // DCA channels of 2.4 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14

	DataRates string `json:"dataRates,omitempty"` // Data rates of 2.4 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 2, 5.5, 6, 9, 11, 12, 18, 24, 36, 48, 54

	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory data rates of 2.4 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 1, 2, 5.5, 6, 9, 11, 12, 18, 24, 36, 48, 54

	PowerThresholdV1 *int `json:"powerThresholdV1,omitempty"` // Power threshold of 2.4 GHz radio band

	RxSopThreshold string `json:"rxSopThreshold,omitempty"` // RX-SOP threshold of 2.4 GHz radio band

	MinPowerLevel *int `json:"minPowerLevel,omitempty"` // Minimum power level of 2.4 GHz radio band

	MaxPowerLevel *int `json:"maxPowerLevel,omitempty"` // Maximum power level of 2.4 GHz radio band

	CustomRxSopThreshold *int `json:"customRxSopThreshold,omitempty"` // RX-SOP threshold custom configuration of 2.4 GHz radio band

	MaxRadioClients *int `json:"maxRadioClients,omitempty"` // Client Limit of 2.4 GHz radio band

	CoverageHoleDetectionProperties *ResponseWirelessGetRfProfilesV1ResponseRadioTypeBPropertiesCoverageHoleDetectionProperties `json:"coverageHoleDetectionProperties,omitempty"` //

	SpatialReuseProperties *ResponseWirelessGetRfProfilesV1ResponseRadioTypeBPropertiesSpatialReuseProperties `json:"spatialReuseProperties,omitempty"` //
}
type ResponseWirelessGetRfProfilesV1ResponseRadioTypeBPropertiesCoverageHoleDetectionProperties struct {
	ChdClientLevel *int `json:"chdClientLevel,omitempty"` // Coverage Hole Detection Client Level

	ChdDataRssiThreshold *int `json:"chdDataRssiThreshold,omitempty"` // Coverage Hole Detection Data Rssi Threshold

	ChdVoiceRssiThreshold *int `json:"chdVoiceRssiThreshold,omitempty"` // Coverage Hole Detection Voice Rssi Threshold

	ChdExceptionLevel *int `json:"chdExceptionLevel,omitempty"` // Coverage Hole Detection Exception Level(%)
}
type ResponseWirelessGetRfProfilesV1ResponseRadioTypeBPropertiesSpatialReuseProperties struct {
	Dot11AxNonSrgObssPacketDetect *bool `json:"dot11axNonSrgObssPacketDetect,omitempty"` // Dot11ax Non SRG OBSS PD

	Dot11AxNonSrgObssPacketDetectMaxThreshold *int `json:"dot11axNonSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax Non SRG OBSS PD Max Threshold

	Dot11AxSrgObssPacketDetect *bool `json:"dot11axSrgObssPacketDetect,omitempty"` // Dot11ax SRG OBSS PD

	Dot11AxSrgObssPacketDetectMinThreshold *int `json:"dot11axSrgObssPacketDetectMinThreshold,omitempty"` // Dot11ax SRG OBSS PD Min Threshold

	Dot11AxSrgObssPacketDetectMaxThreshold *int `json:"dot11axSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax SRG OBSS PD Max Threshold
}
type ResponseWirelessGetRfProfilesV1ResponseRadioType6GHzProperties struct {
	ParentProfile string `json:"parentProfile,omitempty"` // Parent profile of 6 GHz radio band

	RadioChannels string `json:"radioChannels,omitempty"` // DCA channels of 6 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 5, 9, 13, 17, 21, 25, 29, 33, 37, 41, 45, 49, 53, 57, 61, 65, 69, 73, 77, 81, 85, 89, 93, 97, 101, 105, 109, 113, 117, 121, 125, 129, 133, 137, 141, 145, 149, 153, 157, 161, 165, 169, 173, 177, 181, 185, 189, 193, 197, 201, 205, 209, 213, 217, 221, 225, 229, 233

	DataRates string `json:"dataRates,omitempty"` // Data rates of 6 GHz radio band passed in comma separated format without any spaces. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54

	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory data rates of 6 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54

	PowerThresholdV1 *int `json:"powerThresholdV1,omitempty"` // Power threshold of 6 GHz radio band

	RxSopThreshold string `json:"rxSopThreshold,omitempty"` // RX-SOP threshold of 6 GHz radio band

	MinPowerLevel *int `json:"minPowerLevel,omitempty"` // Minimum power level of 6 GHz radio band

	MaxPowerLevel *int `json:"maxPowerLevel,omitempty"` // Maximum power level of 6 GHz radio band

	EnableStandardPowerService *bool `json:"enableStandardPowerService,omitempty"` // True if Standard Power Service is enabled, else False

	MultiBssidProperties *ResponseWirelessGetRfProfilesV1ResponseRadioType6GHzPropertiesMultiBssidProperties `json:"multiBssidProperties,omitempty"` //

	PreamblePuncture *bool `json:"preamblePuncture,omitempty"` // Enable or Disable Preamble Puncturing. This Wifi 7 configuration is applicable to wireless IOS devices supporting 17.15 and higher

	MinDbsWidth *int `json:"minDbsWidth,omitempty"` // Minimum DBS Width ( Permissible values : 20,40,80,160,320)

	MaxDbsWidth *int `json:"maxDbsWidth,omitempty"` // Maximum DBS Width (Permissible Values: 20,40,80,160,320)

	CustomRxSopThreshold *int `json:"customRxSopThreshold,omitempty"` // RX-SOP threshold custom configuration of 6 GHz radio band

	MaxRadioClients *int `json:"maxRadioClients,omitempty"` // Client Limit of 6 GHz radio band

	PscEnforcingEnabled *bool `json:"pscEnforcingEnabled,omitempty"` // PSC Enforcing Enable for 6 GHz radio band

	DiscoveryFrames6GHz string `json:"discoveryFrames6GHz,omitempty"` // Discovery Frames of 6 GHz radio band

	BroadcastProbeResponseInterval *int `json:"broadcastProbeResponseInterval,omitempty"` // Broadcast Probe Response Interval of 6 GHz radio band

	FraProperties *ResponseWirelessGetRfProfilesV1ResponseRadioType6GHzPropertiesFraProperties `json:"fraProperties,omitempty"` //

	CoverageHoleDetectionProperties *ResponseWirelessGetRfProfilesV1ResponseRadioType6GHzPropertiesCoverageHoleDetectionProperties `json:"coverageHoleDetectionProperties,omitempty"` //

	SpatialReuseProperties *ResponseWirelessGetRfProfilesV1ResponseRadioType6GHzPropertiesSpatialReuseProperties `json:"spatialReuseProperties,omitempty"` //
}
type ResponseWirelessGetRfProfilesV1ResponseRadioType6GHzPropertiesMultiBssidProperties struct {
	Dot11AxParameters *ResponseWirelessGetRfProfilesV1ResponseRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters `json:"dot11axParameters,omitempty"` //

	Dot11BeParameters *ResponseWirelessGetRfProfilesV1ResponseRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters `json:"dot11beParameters,omitempty"` //

	TargetWakeTime *bool `json:"targetWakeTime,omitempty"` // Target Wake Time

	TwtBroadcastSupport *bool `json:"twtBroadcastSupport,omitempty"` // TWT Broadcast Support
}
type ResponseWirelessGetRfProfilesV1ResponseRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters struct {
	OfdmaDownLink *bool `json:"ofdmaDownLink,omitempty"` // OFDMA Downlink

	OfdmaUpLink *bool `json:"ofdmaUpLink,omitempty"` // OFDMA Uplink

	MuMimoUpLink *bool `json:"muMimoUpLink,omitempty"` // MU-MIMO Uplink

	MuMimoDownLink *bool `json:"muMimoDownLink,omitempty"` // MU-MIMO Downlink
}
type ResponseWirelessGetRfProfilesV1ResponseRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters struct {
	OfdmaDownLink *bool `json:"ofdmaDownLink,omitempty"` // OFDMA Downlink

	OfdmaUpLink *bool `json:"ofdmaUpLink,omitempty"` // OFDMA Uplink

	MuMimoUpLink *bool `json:"muMimoUpLink,omitempty"` // MU-MIMO Uplink

	MuMimoDownLink *bool `json:"muMimoDownLink,omitempty"` // MU-MIMO Downlink

	OfdmaMultiRu *bool `json:"ofdmaMultiRu,omitempty"` // OFDMA Multi-RU
}
type ResponseWirelessGetRfProfilesV1ResponseRadioType6GHzPropertiesFraProperties struct {
	ClientResetCount *int `json:"clientResetCount,omitempty"` // Client Reset Count of 6 GHz radio band

	ClientUtilizationThreshold *int `json:"clientUtilizationThreshold,omitempty"` // Client Utilization Threshold of 6 GHz radio band
}
type ResponseWirelessGetRfProfilesV1ResponseRadioType6GHzPropertiesCoverageHoleDetectionProperties struct {
	ChdClientLevel *int `json:"chdClientLevel,omitempty"` // Coverage Hole Detection Client Level

	ChdDataRssiThreshold *int `json:"chdDataRssiThreshold,omitempty"` // Coverage Hole Detection Data Rssi Threshold

	ChdVoiceRssiThreshold *int `json:"chdVoiceRssiThreshold,omitempty"` // Coverage Hole Detection Voice Rssi Threshold

	ChdExceptionLevel *int `json:"chdExceptionLevel,omitempty"` // Coverage Hole Detection Exception Level(%)
}
type ResponseWirelessGetRfProfilesV1ResponseRadioType6GHzPropertiesSpatialReuseProperties struct {
	Dot11AxNonSrgObssPacketDetect *bool `json:"dot11axNonSrgObssPacketDetect,omitempty"` // Dot11ax Non SRG OBSS PD

	Dot11AxNonSrgObssPacketDetectMaxThreshold *int `json:"dot11axNonSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax Non SRG OBSS PD Max Threshold

	Dot11AxSrgObssPacketDetect *bool `json:"dot11axSrgObssPacketDetect,omitempty"` // Dot11ax SRG OBSS PD

	Dot11AxSrgObssPacketDetectMinThreshold *int `json:"dot11axSrgObssPacketDetectMinThreshold,omitempty"` // Dot11ax SRG OBSS PD Min Threshold

	Dot11AxSrgObssPacketDetectMaxThreshold *int `json:"dot11axSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax SRG OBSS PD Max Threshold
}
type ResponseWirelessGetRfProfilesCountV1 struct {
	Response *ResponseWirelessGetRfProfilesCountV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response Version
}
type ResponseWirelessGetRfProfilesCountV1Response struct {
	Count *int `json:"count,omitempty"` // Count of the requested resource
}
type ResponseWirelessDeleteRfProfileV1 struct {
	Response *ResponseWirelessDeleteRfProfileV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessDeleteRfProfileV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessGetRfProfileByIDV1 struct {
	Response *ResponseWirelessGetRfProfileByIDV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetRfProfileByIDV1Response struct {
	RfProfileName string `json:"rfProfileName,omitempty"` // RF Profile Name

	DefaultRfProfile *bool `json:"defaultRfProfile,omitempty"` // True if RF Profile is default, else False. Maximum of only 1 RF Profile can be marked as default at any given time

	EnableRadioTypeA *bool `json:"enableRadioTypeA,omitempty"` // True if 5 GHz radio band is enabled in the RF Profile, else False

	EnableRadioTypeB *bool `json:"enableRadioTypeB,omitempty"` // True if 2.4 GHz radio band is enabled in the RF Profile, else False

	EnableRadioType6GHz *bool `json:"enableRadioType6GHz,omitempty"` // True if 6 GHz radio band is enabled in the RF Profile, else False

	EnableCustom *bool `json:"enableCustom,omitempty"` // True if RF Profile is custom, else False for system RF profiles like Low, High and Medium (Typical)

	RadioTypeAProperties *ResponseWirelessGetRfProfileByIDV1ResponseRadioTypeAProperties `json:"radioTypeAProperties,omitempty"` //

	RadioTypeBProperties *ResponseWirelessGetRfProfileByIDV1ResponseRadioTypeBProperties `json:"radioTypeBProperties,omitempty"` //

	RadioType6GHzProperties *ResponseWirelessGetRfProfileByIDV1ResponseRadioType6GHzProperties `json:"radioType6GHzProperties,omitempty"` //

	ID string `json:"id,omitempty"` // RF Profile ID
}
type ResponseWirelessGetRfProfileByIDV1ResponseRadioTypeAProperties struct {
	ParentProfile string `json:"parentProfile,omitempty"` // Parent profile of 5 GHz radio band

	RadioChannels string `json:"radioChannels,omitempty"` // DCA channels of 5 GHz radio band passed in comma separated format without any spaces. Permissible values: 36, 40, 44, 48, 52, 56, 60, 64, 100, 104, 108, 112, 116, 120, 124, 128, 132, 136, 140, 144, 149, 153, 157, 161, 165, 169, 173

	DataRates string `json:"dataRates,omitempty"` // Data rates of 5 GHz radio band passed in comma separated format without any spaces. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54

	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory data rates of 5 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54

	PowerThresholdV1 *int `json:"powerThresholdV1,omitempty"` // Power threshold of 5 GHz radio band

	RxSopThreshold string `json:"rxSopThreshold,omitempty"` // RX-SOP threshold of 5 GHz radio band

	MinPowerLevel *int `json:"minPowerLevel,omitempty"` // Minimum power level of 5 GHz radio band

	MaxPowerLevel *int `json:"maxPowerLevel,omitempty"` // Maximum power level of 5 GHz radio band

	ChannelWidth string `json:"channelWidth,omitempty"` // Channel Width

	PreamblePuncture *bool `json:"preamblePuncture,omitempty"` // Enable or Disable Preamble Puncturing. This Wifi 7 configuration is applicable to wireless IOS devices supporting 17.15 and higher

	ZeroWaitDfsEnable *bool `json:"zeroWaitDfsEnable,omitempty"` // Zero Wait DFS is applicable only for IOS-XE based Wireless Controllers running 17.9.1 and above versions

	CustomRxSopThreshold *int `json:"customRxSopThreshold,omitempty"` // RX-SOP threshold custom configuration of 5 GHz radio band

	MaxRadioClients *int `json:"maxRadioClients,omitempty"` // Client Limit of 5 GHz radio band

	FraProperties *ResponseWirelessGetRfProfileByIDV1ResponseRadioTypeAPropertiesFraProperties `json:"fraProperties,omitempty"` //

	CoverageHoleDetectionProperties *ResponseWirelessGetRfProfileByIDV1ResponseRadioTypeAPropertiesCoverageHoleDetectionProperties `json:"coverageHoleDetectionProperties,omitempty"` //

	SpatialReuseProperties *ResponseWirelessGetRfProfileByIDV1ResponseRadioTypeAPropertiesSpatialReuseProperties `json:"spatialReuseProperties,omitempty"` //
}
type ResponseWirelessGetRfProfileByIDV1ResponseRadioTypeAPropertiesFraProperties struct {
	ClientAware *bool `json:"clientAware,omitempty"` // Client Aware of 5 GHz radio band

	ClientSelect *int `json:"clientSelect,omitempty"` // Client Select(%) of 5 GHz radio band

	ClientReset *int `json:"clientReset,omitempty"` // Client Reset(%) of 5 GHz radio band
}
type ResponseWirelessGetRfProfileByIDV1ResponseRadioTypeAPropertiesCoverageHoleDetectionProperties struct {
	ChdClientLevel *int `json:"chdClientLevel,omitempty"` // Coverage Hole Detection Client Level

	ChdDataRssiThreshold *int `json:"chdDataRssiThreshold,omitempty"` // Coverage Hole Detection Data Rssi Threshold

	ChdVoiceRssiThreshold *int `json:"chdVoiceRssiThreshold,omitempty"` // Coverage Hole Detection Voice Rssi Threshold

	ChdExceptionLevel *int `json:"chdExceptionLevel,omitempty"` // Coverage Hole Detection Exception Level(%)
}
type ResponseWirelessGetRfProfileByIDV1ResponseRadioTypeAPropertiesSpatialReuseProperties struct {
	Dot11AxNonSrgObssPacketDetect *bool `json:"dot11axNonSrgObssPacketDetect,omitempty"` // Dot11ax Non SRG OBSS PD

	Dot11AxNonSrgObssPacketDetectMaxThreshold *int `json:"dot11axNonSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax Non SRG OBSS PD Max Threshold

	Dot11AxSrgObssPacketDetect *bool `json:"dot11axSrgObssPacketDetect,omitempty"` // Dot11ax SRG OBSS PD

	Dot11AxSrgObssPacketDetectMinThreshold *int `json:"dot11axSrgObssPacketDetectMinThreshold,omitempty"` // Dot11ax SRG OBSS PD Min Threshold

	Dot11AxSrgObssPacketDetectMaxThreshold *int `json:"dot11axSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax SRG OBSS PD Max Threshold
}
type ResponseWirelessGetRfProfileByIDV1ResponseRadioTypeBProperties struct {
	ParentProfile string `json:"parentProfile,omitempty"` // Parent profile of 2.4 GHz radio band

	RadioChannels string `json:"radioChannels,omitempty"` // DCA channels of 2.4 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14

	DataRates string `json:"dataRates,omitempty"` // Data rates of 2.4 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 2, 5.5, 6, 9, 11, 12, 18, 24, 36, 48, 54

	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory data rates of 2.4 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 1, 2, 5.5, 6, 9, 11, 12, 18, 24, 36, 48, 54

	PowerThresholdV1 *int `json:"powerThresholdV1,omitempty"` // Power threshold of 2.4 GHz radio band

	RxSopThreshold string `json:"rxSopThreshold,omitempty"` // RX-SOP threshold of 2.4 GHz radio band

	MinPowerLevel *int `json:"minPowerLevel,omitempty"` // Minimum power level of 2.4 GHz radio band

	MaxPowerLevel *int `json:"maxPowerLevel,omitempty"` // Maximum power level of 2.4 GHz radio band

	CustomRxSopThreshold *int `json:"customRxSopThreshold,omitempty"` // RX-SOP threshold custom configuration of 2.4 GHz radio band

	MaxRadioClients *int `json:"maxRadioClients,omitempty"` // Client Limit of 2.4 GHz radio band

	CoverageHoleDetectionProperties *ResponseWirelessGetRfProfileByIDV1ResponseRadioTypeBPropertiesCoverageHoleDetectionProperties `json:"coverageHoleDetectionProperties,omitempty"` //

	SpatialReuseProperties *ResponseWirelessGetRfProfileByIDV1ResponseRadioTypeBPropertiesSpatialReuseProperties `json:"spatialReuseProperties,omitempty"` //
}
type ResponseWirelessGetRfProfileByIDV1ResponseRadioTypeBPropertiesCoverageHoleDetectionProperties struct {
	ChdClientLevel *int `json:"chdClientLevel,omitempty"` // Coverage Hole Detection Client Level

	ChdDataRssiThreshold *int `json:"chdDataRssiThreshold,omitempty"` // Coverage Hole Detection Data Rssi Threshold

	ChdVoiceRssiThreshold *int `json:"chdVoiceRssiThreshold,omitempty"` // Coverage Hole Detection Voice Rssi Threshold

	ChdExceptionLevel *int `json:"chdExceptionLevel,omitempty"` // Coverage Hole Detection Exception Level(%)
}
type ResponseWirelessGetRfProfileByIDV1ResponseRadioTypeBPropertiesSpatialReuseProperties struct {
	Dot11AxNonSrgObssPacketDetect *bool `json:"dot11axNonSrgObssPacketDetect,omitempty"` // Dot11ax Non SRG OBSS PD

	Dot11AxNonSrgObssPacketDetectMaxThreshold *int `json:"dot11axNonSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax Non SRG OBSS PD Max Threshold

	Dot11AxSrgObssPacketDetect *bool `json:"dot11axSrgObssPacketDetect,omitempty"` // Dot11ax SRG OBSS PD

	Dot11AxSrgObssPacketDetectMinThreshold *int `json:"dot11axSrgObssPacketDetectMinThreshold,omitempty"` // Dot11ax SRG OBSS PD Min Threshold

	Dot11AxSrgObssPacketDetectMaxThreshold *int `json:"dot11axSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax SRG OBSS PD Max Threshold
}
type ResponseWirelessGetRfProfileByIDV1ResponseRadioType6GHzProperties struct {
	ParentProfile string `json:"parentProfile,omitempty"` // Parent profile of 6 GHz radio band

	RadioChannels string `json:"radioChannels,omitempty"` // DCA channels of 6 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 5, 9, 13, 17, 21, 25, 29, 33, 37, 41, 45, 49, 53, 57, 61, 65, 69, 73, 77, 81, 85, 89, 93, 97, 101, 105, 109, 113, 117, 121, 125, 129, 133, 137, 141, 145, 149, 153, 157, 161, 165, 169, 173, 177, 181, 185, 189, 193, 197, 201, 205, 209, 213, 217, 221, 225, 229, 233

	DataRates string `json:"dataRates,omitempty"` // Data rates of 6 GHz radio band passed in comma separated format without any spaces. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54

	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory data rates of 6 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54

	PowerThresholdV1 *int `json:"powerThresholdV1,omitempty"` // Power threshold of 6 GHz radio band

	RxSopThreshold string `json:"rxSopThreshold,omitempty"` // RX-SOP threshold of 6 GHz radio band

	MinPowerLevel *int `json:"minPowerLevel,omitempty"` // Minimum power level of 6 GHz radio band

	MaxPowerLevel *int `json:"maxPowerLevel,omitempty"` // Maximum power level of 6 GHz radio band

	EnableStandardPowerService *bool `json:"enableStandardPowerService,omitempty"` // True if Standard Power Service is enabled, else False

	MultiBssidProperties *ResponseWirelessGetRfProfileByIDV1ResponseRadioType6GHzPropertiesMultiBssidProperties `json:"multiBssidProperties,omitempty"` //

	PreamblePuncture *bool `json:"preamblePuncture,omitempty"` // Enable or Disable Preamble Puncturing. This Wifi 7 configuration is applicable to wireless IOS devices supporting 17.15 and higher

	MinDbsWidth *int `json:"minDbsWidth,omitempty"` // Minimum DBS Width ( Permissible values : 20,40,80,160,320)

	MaxDbsWidth *int `json:"maxDbsWidth,omitempty"` // Maximum DBS Width (Permissible Values: 20,40,80,160,320)

	CustomRxSopThreshold *int `json:"customRxSopThreshold,omitempty"` // RX-SOP threshold custom configuration of 6 GHz radio band

	MaxRadioClients *int `json:"maxRadioClients,omitempty"` // Client Limit of 6 GHz radio band

	PscEnforcingEnabled *bool `json:"pscEnforcingEnabled,omitempty"` // PSC Enforcing Enable for 6 GHz radio band

	DiscoveryFrames6GHz string `json:"discoveryFrames6GHz,omitempty"` // Discovery Frames of 6 GHz radio band

	BroadcastProbeResponseInterval *int `json:"broadcastProbeResponseInterval,omitempty"` // Broadcast Probe Response Interval of 6 GHz radio band

	FraProperties *ResponseWirelessGetRfProfileByIDV1ResponseRadioType6GHzPropertiesFraProperties `json:"fraProperties,omitempty"` //

	CoverageHoleDetectionProperties *ResponseWirelessGetRfProfileByIDV1ResponseRadioType6GHzPropertiesCoverageHoleDetectionProperties `json:"coverageHoleDetectionProperties,omitempty"` //

	SpatialReuseProperties *ResponseWirelessGetRfProfileByIDV1ResponseRadioType6GHzPropertiesSpatialReuseProperties `json:"spatialReuseProperties,omitempty"` //
}
type ResponseWirelessGetRfProfileByIDV1ResponseRadioType6GHzPropertiesMultiBssidProperties struct {
	Dot11AxParameters *ResponseWirelessGetRfProfileByIDV1ResponseRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters `json:"dot11axParameters,omitempty"` //

	Dot11BeParameters *ResponseWirelessGetRfProfileByIDV1ResponseRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters `json:"dot11beParameters,omitempty"` //

	TargetWakeTime *bool `json:"targetWakeTime,omitempty"` // Target Wake Time

	TwtBroadcastSupport *bool `json:"twtBroadcastSupport,omitempty"` // TWT Broadcast Support
}
type ResponseWirelessGetRfProfileByIDV1ResponseRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters struct {
	OfdmaDownLink *bool `json:"ofdmaDownLink,omitempty"` // OFDMA Downlink

	OfdmaUpLink *bool `json:"ofdmaUpLink,omitempty"` // OFDMA Uplink

	MuMimoUpLink *bool `json:"muMimoUpLink,omitempty"` // MU-MIMO Uplink

	MuMimoDownLink *bool `json:"muMimoDownLink,omitempty"` // MU-MIMO Downlink
}
type ResponseWirelessGetRfProfileByIDV1ResponseRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters struct {
	OfdmaDownLink *bool `json:"ofdmaDownLink,omitempty"` // OFDMA Downlink

	OfdmaUpLink *bool `json:"ofdmaUpLink,omitempty"` // OFDMA Uplink

	MuMimoUpLink *bool `json:"muMimoUpLink,omitempty"` // MU-MIMO Uplink

	MuMimoDownLink *bool `json:"muMimoDownLink,omitempty"` // MU-MIMO Downlink

	OfdmaMultiRu *bool `json:"ofdmaMultiRu,omitempty"` // OFDMA Multi-RU
}
type ResponseWirelessGetRfProfileByIDV1ResponseRadioType6GHzPropertiesFraProperties struct {
	ClientResetCount *int `json:"clientResetCount,omitempty"` // Client Reset Count of 6 GHz radio band

	ClientUtilizationThreshold *int `json:"clientUtilizationThreshold,omitempty"` // Client Utilization Threshold of 6 GHz radio band
}
type ResponseWirelessGetRfProfileByIDV1ResponseRadioType6GHzPropertiesCoverageHoleDetectionProperties struct {
	ChdClientLevel *int `json:"chdClientLevel,omitempty"` // Coverage Hole Detection Client Level

	ChdDataRssiThreshold *int `json:"chdDataRssiThreshold,omitempty"` // Coverage Hole Detection Data Rssi Threshold

	ChdVoiceRssiThreshold *int `json:"chdVoiceRssiThreshold,omitempty"` // Coverage Hole Detection Voice Rssi Threshold

	ChdExceptionLevel *int `json:"chdExceptionLevel,omitempty"` // Coverage Hole Detection Exception Level(%)
}
type ResponseWirelessGetRfProfileByIDV1ResponseRadioType6GHzPropertiesSpatialReuseProperties struct {
	Dot11AxNonSrgObssPacketDetect *bool `json:"dot11axNonSrgObssPacketDetect,omitempty"` // Dot11ax Non SRG OBSS PD

	Dot11AxNonSrgObssPacketDetectMaxThreshold *int `json:"dot11axNonSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax Non SRG OBSS PD Max Threshold

	Dot11AxSrgObssPacketDetect *bool `json:"dot11axSrgObssPacketDetect,omitempty"` // Dot11ax SRG OBSS PD

	Dot11AxSrgObssPacketDetectMinThreshold *int `json:"dot11axSrgObssPacketDetectMinThreshold,omitempty"` // Dot11ax SRG OBSS PD Min Threshold

	Dot11AxSrgObssPacketDetectMaxThreshold *int `json:"dot11axSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax SRG OBSS PD Max Threshold
}
type ResponseWirelessUpdateRfProfileV1 struct {
	Response *ResponseWirelessUpdateRfProfileV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessUpdateRfProfileV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessRetrieveSitesWithOverriddenSSIDsV1 struct {
	Response *[]ResponseWirelessRetrieveSitesWithOverriddenSSIDsV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response version
}
type ResponseWirelessRetrieveSitesWithOverriddenSSIDsV1Response struct {
	SiteID string `json:"siteId,omitempty"` // Site ID

	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Site Name Hierarchy

	SSIDs *[]ResponseWirelessRetrieveSitesWithOverriddenSSIDsV1ResponseSSIDs `json:"ssids,omitempty"` //
}
type ResponseWirelessRetrieveSitesWithOverriddenSSIDsV1ResponseSSIDs struct {
	ID string `json:"id,omitempty"` // SSID ID

	SSID string `json:"ssid,omitempty"` // SSID
}
type ResponseWirelessAssignAnchorManagedApLocationsForWLCV1 struct {
	Response *ResponseWirelessAssignAnchorManagedApLocationsForWLCV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version of the response
}
type ResponseWirelessAssignAnchorManagedApLocationsForWLCV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Unique identifier for the task

	URL string `json:"url,omitempty"` // URL for the task
}
type ResponseWirelessConfigureAccessPointsV2 struct {
	Response *ResponseWirelessConfigureAccessPointsV2Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseWirelessConfigureAccessPointsV2Response struct {
	TaskID string `json:"taskId,omitempty"` //

	URL string `json:"url,omitempty"` //
}
type RequestWirelessCreateAndProvisionSSIDV1 struct {
	ManagedApLocations []string `json:"managedAPLocations,omitempty"` // Managed AP Locations (Enter entire Site(s) hierarchy)

	SSIDDetails *RequestWirelessCreateAndProvisionSSIDV1SSIDDetails `json:"ssidDetails,omitempty"` //

	SSIDType string `json:"ssidType,omitempty"` // SSID Type

	EnableFabric *bool `json:"enableFabric,omitempty"` // Enable SSID for Fabric

	FlexConnect *RequestWirelessCreateAndProvisionSSIDV1FlexConnect `json:"flexConnect,omitempty"` //
}
type RequestWirelessCreateAndProvisionSSIDV1SSIDDetails struct {
	Name string `json:"name,omitempty"` // SSID Name

	SecurityLevel string `json:"securityLevel,omitempty"` // Security Level(For guest SSID OPEN/WEB_AUTH, For Enterprise SSID ENTERPRISE/PERSONAL/OPEN)

	EnableFastLane *bool `json:"enableFastLane,omitempty"` // Enable Fast Lane

	Passphrase string `json:"passphrase,omitempty"` // Pass Phrase ( Only applicable for SSID with PERSONAL auth type )

	TrafficType string `json:"trafficType,omitempty"` // Traffic Type

	EnableBroadcastSSID *bool `json:"enableBroadcastSSID,omitempty"` // Enable Broadcast SSID

	RadioPolicy string `json:"radioPolicy,omitempty"` // Radio Policy

	EnableMacFiltering *bool `json:"enableMACFiltering,omitempty"` // Enable MAC Filtering

	FastTransition string `json:"fastTransition,omitempty"` // Fast Transition

	WebAuthURL string `json:"webAuthURL,omitempty"` // Web Auth URL

	AuthKeyMgmt []string `json:"authKeyMgmt,omitempty"` // Takes string inputs for the AKMs that should be set true. Possible AKM values : dot1x,dot1x_ft, dot1x_sha, psk, psk_ft, psk_sha, owe, sae, sae_ft

	RsnCipherSuiteGcmp256 *bool `json:"rsnCipherSuiteGcmp256,omitempty"` // Rsn Cipher Suite Gcmp256

	RsnCipherSuiteGcmp128 *bool `json:"rsnCipherSuiteGcmp128,omitempty"` // Rsn Cipher Suite  Gcmp128

	RsnCipherSuiteCcmp256 *bool `json:"rsnCipherSuiteCcmp256,omitempty"` // Rsn Cipher Suite Ccmp256

	Ghz6PolicyClientSteering *bool `json:"ghz6PolicyClientSteering,omitempty"` // 6 Ghz Client Steering

	Ghz24Policy string `json:"ghz24Policy,omitempty"` // 2.4 GHz Policy
}
type RequestWirelessCreateAndProvisionSSIDV1FlexConnect struct {
	EnableFlexConnect *bool `json:"enableFlexConnect,omitempty"` // Enable Flex Connect

	LocalToVLAN *int `json:"localToVlan,omitempty"` // Local To Vlan (range is 1 to 4094)
}
type RequestWirelessRebootAccessPointsV1 struct {
	ApMacAddresses []string `json:"apMacAddresses,omitempty"` // The ethernet MAC address of the access point.
}
type RequestWirelessCreateEnterpriseSSIDV1 struct {
	Name string `json:"name,omitempty"` // SSID NAME

	SecurityLevel string `json:"securityLevel,omitempty"` // Security Level

	Passphrase string `json:"passphrase,omitempty"` // Passphrase

	EnableFastLane *bool `json:"enableFastLane,omitempty"` // Enable FastLane

	EnableMacFiltering *bool `json:"enableMACFiltering,omitempty"` // Enable MAC Filtering

	TrafficType string `json:"trafficType,omitempty"` // Traffic Type Enum (voicedata or data )

	RadioPolicy string `json:"radioPolicy,omitempty"` // Radio Policy Enum

	EnableBroadcastSSID *bool `json:"enableBroadcastSSID,omitempty"` // Enable Broadcase SSID

	FastTransition string `json:"fastTransition,omitempty"` // Fast Transition

	EnableSessionTimeOut *bool `json:"enableSessionTimeOut,omitempty"` // Enable Session Timeout

	SessionTimeOut *int `json:"sessionTimeOut,omitempty"` // Session Time Out

	EnableClientExclusion *bool `json:"enableClientExclusion,omitempty"` // Enable Client Exclusion

	ClientExclusionTimeout *int `json:"clientExclusionTimeout,omitempty"` // Client Exclusion Timeout

	EnableBasicServiceSetMaxIDle *bool `json:"enableBasicServiceSetMaxIdle,omitempty"` // Enable Basic Service Set Max Idle

	BasicServiceSetClientIDleTimeout *int `json:"basicServiceSetClientIdleTimeout,omitempty"` // Basic Service Set Client Idle Timeout

	EnableDirectedMulticastService *bool `json:"enableDirectedMulticastService,omitempty"` // Enable Directed Multicast Service

	EnableNeighborList *bool `json:"enableNeighborList,omitempty"` // Enable Neighbor List

	MfpClientProtection string `json:"mfpClientProtection,omitempty"` // Management Frame Protection Client

	NasOptions []string `json:"nasOptions,omitempty"` // Nas Options

	ProfileName string `json:"profileName,omitempty"` // Profile Name

	PolicyProfileName string `json:"policyProfileName,omitempty"` // Policy Profile Name

	AAAOverride *bool `json:"aaaOverride,omitempty"` // Aaa Override

	CoverageHoleDetectionEnable *bool `json:"coverageHoleDetectionEnable,omitempty"` // Coverage Hole Detection Enable

	ProtectedManagementFrame string `json:"protectedManagementFrame,omitempty"` // (Required applicable for Security Type WPA3_PERSONAL, WPA3_ENTERPRISE, OPEN_SECURED) and (Optional, Required Applicable for Security Type WPA2_WPA3_PERSONAL and WPA2_WPA3_ENTERPRISE)

	MultipSKSettings *[]RequestWirelessCreateEnterpriseSSIDV1MultipSKSettings `json:"multiPSKSettings,omitempty"` //

	ClientRateLimit *float64 `json:"clientRateLimit,omitempty"` // Client Rate Limit (in bits per second)

	AuthKeyMgmt []string `json:"authKeyMgmt,omitempty"` // Takes string inputs for the AKMs that should be set true. Possible AKM values : dot1x,dot1x_ft, dot1x_sha, psk, psk_ft, psk_sha, owe, sae, sae_ft

	RsnCipherSuiteGcmp256 *bool `json:"rsnCipherSuiteGcmp256,omitempty"` // Rsn Cipher Suite Gcmp256

	RsnCipherSuiteCcmp256 *bool `json:"rsnCipherSuiteCcmp256,omitempty"` // Rsn Cipher Suite Ccmp256

	RsnCipherSuiteGcmp128 *bool `json:"rsnCipherSuiteGcmp128,omitempty"` // Rsn Cipher Suite Gcmp 128

	Ghz6PolicyClientSteering *bool `json:"ghz6PolicyClientSteering,omitempty"` // Ghz6 Policy Client Steering

	Ghz24Policy string `json:"ghz24Policy,omitempty"` // Ghz24 Policy
}
type RequestWirelessCreateEnterpriseSSIDV1MultipSKSettings struct {
	Priority *int `json:"priority,omitempty"` // Priority

	PassphraseType string `json:"passphraseType,omitempty"` // Passphrase Type

	Passphrase string `json:"passphrase,omitempty"` // Passphrase
}
type RequestWirelessUpdateEnterpriseSSIDV1 struct {
	Name string `json:"name,omitempty"` // SSID NAME

	SecurityLevel string `json:"securityLevel,omitempty"` // Security Level

	Passphrase string `json:"passphrase,omitempty"` // Passphrase

	EnableFastLane *bool `json:"enableFastLane,omitempty"` // Enable FastLane

	EnableMacFiltering *bool `json:"enableMACFiltering,omitempty"` // Enable MAC Filtering

	TrafficType string `json:"trafficType,omitempty"` // Traffic Type Enum (voicedata or data )

	RadioPolicy string `json:"radioPolicy,omitempty"` // Radio Policy Enum

	EnableBroadcastSSID *bool `json:"enableBroadcastSSID,omitempty"` // Enable Broadcase SSID

	FastTransition string `json:"fastTransition,omitempty"` // Fast Transition

	EnableSessionTimeOut *bool `json:"enableSessionTimeOut,omitempty"` // Enable Session Timeout

	SessionTimeOut *int `json:"sessionTimeOut,omitempty"` // Session Time Out

	EnableClientExclusion *bool `json:"enableClientExclusion,omitempty"` // Enable Client Exclusion

	ClientExclusionTimeout *int `json:"clientExclusionTimeout,omitempty"` // Client Exclusion Timeout

	EnableBasicServiceSetMaxIDle *bool `json:"enableBasicServiceSetMaxIdle,omitempty"` // Enable Basic Service Set Max Idle

	BasicServiceSetClientIDleTimeout *int `json:"basicServiceSetClientIdleTimeout,omitempty"` // Basic Service Set Client Idle Timeout

	EnableDirectedMulticastService *bool `json:"enableDirectedMulticastService,omitempty"` // Enable Directed Multicast Service

	EnableNeighborList *bool `json:"enableNeighborList,omitempty"` // Enable Neighbor List

	MfpClientProtection string `json:"mfpClientProtection,omitempty"` // Management Frame Protection Client

	NasOptions []string `json:"nasOptions,omitempty"` // Nas Options

	ProfileName string `json:"profileName,omitempty"` // Profile Name

	PolicyProfileName string `json:"policyProfileName,omitempty"` // Policy Profile Name

	AAAOverride *bool `json:"aaaOverride,omitempty"` // Aaa Override

	CoverageHoleDetectionEnable *bool `json:"coverageHoleDetectionEnable,omitempty"` // Coverage Hole Detection Enable

	ProtectedManagementFrame string `json:"protectedManagementFrame,omitempty"` // (Required applicable for Security Type WPA3_PERSONAL, WPA3_ENTERPRISE, OPEN_SECURED) and (Optional, Required Applicable for Security Type WPA2_WPA3_PERSONAL and WPA2_WPA3_ENTERPRISE)

	MultipSKSettings *[]RequestWirelessUpdateEnterpriseSSIDV1MultipSKSettings `json:"multiPSKSettings,omitempty"` //

	ClientRateLimit *float64 `json:"clientRateLimit,omitempty"` // Client Rate Limit (in bits per second)

	AuthKeyMgmt []string `json:"authKeyMgmt,omitempty"` // Takes string inputs for the AKMs that should be set true. Possible AKM values : dot1x,dot1x_ft, dot1x_sha, psk, psk_ft, psk_sha, owe, sae, sae_ft

	RsnCipherSuiteGcmp256 *bool `json:"rsnCipherSuiteGcmp256,omitempty"` // Rsn Cipher Suite Gcmp256

	RsnCipherSuiteCcmp256 *bool `json:"rsnCipherSuiteCcmp256,omitempty"` // Rsn Cipher Suite Ccmp256

	RsnCipherSuiteGcmp128 *bool `json:"rsnCipherSuiteGcmp128,omitempty"` // Rsn Cipher Suite Gcmp 128

	Ghz6PolicyClientSteering *bool `json:"ghz6PolicyClientSteering,omitempty"` // Ghz6 Policy Client Steering

	Ghz24Policy string `json:"ghz24Policy,omitempty"` // Ghz24 Policy
}
type RequestWirelessUpdateEnterpriseSSIDV1MultipSKSettings struct {
	Priority *int `json:"priority,omitempty"` // Priority

	PassphraseType string `json:"passphraseType,omitempty"` // Passphrase Type

	Passphrase string `json:"passphrase,omitempty"` // Passphrase
}
type RequestWirelessCreateSSIDV1 struct {
	SSID string `json:"ssid,omitempty"` // Name of the SSID

	AuthType string `json:"authType,omitempty"` // L2 Authentication Type (If authType is not open , then atleast one RSN Cipher Suite and corresponding valid AKM must be enabled). Default is L2 Authentication Type if exists else .

	Passphrase string `json:"passphrase,omitempty"` // Passphrase (Only applicable for SSID with PERSONAL security level). Passphrase needs to be between 8 and 63 characters for ASCII type. HEX passphrase needs to be 64 characters

	IsFastLaneEnabled *bool `json:"isFastLaneEnabled,omitempty"` // True if FastLane is enabled, else False

	IsMacFilteringEnabled *bool `json:"isMacFilteringEnabled,omitempty"` // When set to true, MAC Filtering will be activated, allowing control over network access based on the MAC address of the device

	SSIDRadioType string `json:"ssidRadioType,omitempty"` // Radio Policy Enum (default: Triple band operation(2.4GHz, 5GHz and 6GHz))

	IsBroadcastSSID *bool `json:"isBroadcastSSID,omitempty"` // When activated by setting it to true, the Broadcast SSID feature will make the SSID publicly visible to wireless devices searching for available networks

	FastTransition string `json:"fastTransition,omitempty"` // Fast Transition

	SessionTimeOutEnable *bool `json:"sessionTimeOutEnable,omitempty"` // Turn on the feature that imposes a time limit on user sessions

	SessionTimeOut *int `json:"sessionTimeOut,omitempty"` // This denotes the allotted time span, expressed in seconds, before a session is automatically terminated due to inactivity. Default sessionTimeOut is 1800.

	ClientExclusionEnable *bool `json:"clientExclusionEnable,omitempty"` // Activate the feature that allows for the exclusion of clients

	ClientExclusionTimeout *int `json:"clientExclusionTimeout,omitempty"` // This refers to the length of time, in seconds, a client is excluded or blocked from accessing the network after a specified number of unsuccessful attempts. Default is Client Exclusion Timeout if exists else 180.

	BasicServiceSetMaxIDleEnable *bool `json:"basicServiceSetMaxIdleEnable,omitempty"` // Activate the maximum idle feature for the Basic Service Set

	BasicServiceSetClientIDleTimeout *int `json:"basicServiceSetClientIdleTimeout,omitempty"` // This refers to the duration of inactivity, measured in seconds, before a client connected to the Basic Service Set is considered idle and timed out. Default is Basic ServiceSet ClientIdle Timeout if exists else 300.

	DirectedMulticastServiceEnable *bool `json:"directedMulticastServiceEnable,omitempty"` // The Directed Multicast Service feature becomes operational when it is set to true

	NeighborListEnable *bool `json:"neighborListEnable,omitempty"` // The Neighbor List feature is enabled when it is set to true

	ManagementFrameProtectionClientprotection string `json:"managementFrameProtectionClientprotection,omitempty"` // Default is Management Frame Protection Client if exists else Optional.

	NasOptions []string `json:"nasOptions,omitempty"` // Pre-Defined NAS Options : AP ETH Mac Address, AP IP address, AP Location , AP MAC Address, AP Name, AP Policy Tag, AP Site Tag, SSID, System IP Address, System MAC Address, System Name.

	ProfileName string `json:"profileName,omitempty"` // WLAN Profile Name, if not passed autogenerated profile name will be assigned. The same wlanProfileName will also be used for policyProfileName

	AAAOverride *bool `json:"aaaOverride,omitempty"` // Activate the AAA Override feature when set to true

	CoverageHoleDetectionEnable *bool `json:"coverageHoleDetectionEnable,omitempty"` // Activate Coverage Hole Detection feature when set to true

	ProtectedManagementFrame string `json:"protectedManagementFrame,omitempty"` // (REQUIRED is applicable for authType WPA3_PERSONAL, WPA3_ENTERPRISE, OPEN_SECURED) and (OPTIONAL/REQUIRED is applicable for authType WPA2_WPA3_PERSONAL and WPA2_WPA3_ENTERPRISE)

	MultipSKSettings *[]RequestWirelessCreateSSIDV1MultipSKSettings `json:"multiPSKSettings,omitempty"` //

	ClientRateLimit *int `json:"clientRateLimit,omitempty"` // This pertains to the maximum data transfer rate, specified in bits per second, that a client is permitted to achieve. It should be in mutliples of 500 . Default is Client Rate Limit if exists else 0.

	RsnCipherSuiteGcmp256 *bool `json:"rsnCipherSuiteGcmp256,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite GCMP256 encryption protocol is activated

	RsnCipherSuiteCcmp256 *bool `json:"rsnCipherSuiteCcmp256,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite CCMP256 encryption protocol is activated

	RsnCipherSuiteGcmp128 *bool `json:"rsnCipherSuiteGcmp128,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite GCMP128 encryption protocol is activated

	RsnCipherSuiteCcmp128 *bool `json:"rsnCipherSuiteCcmp128,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite CCMP128 encryption protocol is activated

	Ghz6PolicyClientSteering *bool `json:"ghz6PolicyClientSteering,omitempty"` // True if 6 GHz Policy Client Steering is enabled, else False

	IsAuthKey8021X *bool `json:"isAuthKey8021x,omitempty"` // When set to true, the 802.1X authentication key is in use

	IsAuthKey8021XPlusFT *bool `json:"isAuthKey8021xPlusFT,omitempty"` // When set to true, the 802.1X-Plus-FT authentication key is in use

	IsAuthKey8021XSHA256 *bool `json:"isAuthKey8021x_SHA256,omitempty"` // When set to true, the feature that enables 802.1X authentication using the SHA256 algorithm is turned on

	IsAuthKeySae *bool `json:"isAuthKeySae,omitempty"` // When set to true, the feature enabling the Simultaneous Authentication of Equals (SAE) authentication key is activated

	IsAuthKeySaePlusFT *bool `json:"isAuthKeySaePlusFT,omitempty"` // Activating this setting by switching it to true turns on the authentication key feature that supports both Simultaneous Authentication of Equals (SAE) and Fast Transition (FT)

	IsAuthKeyPSK *bool `json:"isAuthKeyPSK,omitempty"` // When set to true, the Pre-shared Key (PSK) authentication feature is enabled

	IsAuthKeyPSKPlusFT *bool `json:"isAuthKeyPSKPlusFT,omitempty"` // When set to true, the feature that enables the combination of Pre-shared Key (PSK) and Fast Transition (FT) authentication keys is activated

	IsAuthKeyOWE *bool `json:"isAuthKeyOWE,omitempty"` // When set to true, the Opportunistic Wireless Encryption (OWE) authentication key feature is turned on

	IsAuthKeyEasyPSK *bool `json:"isAuthKeyEasyPSK,omitempty"` // When set to true, the feature that enables the use of Easy Pre-shared Key (PSK) authentication is activated

	IsAuthKeyPSKSHA256 *bool `json:"isAuthKeyPSKSHA256,omitempty"` // The feature that allows the use of Pre-shared Key (PSK) authentication with the SHA256 algorithm is enabled when it is set to true

	OpenSSID string `json:"openSsid,omitempty"` // Open SSID which is already created in the design and not associated to any other OPEN-SECURED SSID

	WLANBandSelectEnable *bool `json:"wlanBandSelectEnable,omitempty"` // Band select is allowed only when band options selected contains at least 2.4 GHz and 5 GHz band else false.

	IsEnabled *bool `json:"isEnabled,omitempty"` // Set SSID's admin status as 'Enabled' when set to true

	AuthServers []string `json:"authServers,omitempty"` // List of Authentication/Authorization server IpAddresses

	AcctServers []string `json:"acctServers,omitempty"` // List of Accounting server IpAddresses

	EgressQos string `json:"egressQos,omitempty"` // Egress QOS

	IngressQos string `json:"ingressQos,omitempty"` // Ingress QOS

	WLANType string `json:"wlanType,omitempty"` // Wlan Type

	L3AuthType string `json:"l3AuthType,omitempty"` // Default is L3 Authentication Type if exists else .

	AuthServer string `json:"authServer,omitempty"` // Authentication Server, Mandatory for Guest SSIDs with wlanType=Guest and l3AuthType=web_auth

	ExternalAuthIPAddress string `json:"externalAuthIpAddress,omitempty"` // External WebAuth URL (Mandatory for Guest SSIDs with wlanType = Guest, l3AuthType = web_auth and authServer = auth_external)

	WebPassthrough *bool `json:"webPassthrough,omitempty"` // When set to true, the Web-Passthrough feature will be activated for the Guest SSID, allowing guests to bypass certain login requirements

	SleepingClientEnable *bool `json:"sleepingClientEnable,omitempty"` // When set to true, this will activate the timeout settings that apply to clients in sleep mode

	SleepingClientTimeout *int `json:"sleepingClientTimeout,omitempty"` // This refers to the amount of time, measured in minutes, before a sleeping (inactive) client is timed out of the network. Default is Sleeping Client Timeout if exists else 720.

	ACLName string `json:"aclName,omitempty"` // Pre-Auth Access Control List (ACL) Name

	IsPosturingEnabled *bool `json:"isPosturingEnabled,omitempty"` // Applicable only for Enterprise SSIDs. When set to True, Posturing will enabled. Required to be set to True if ACL needs to be mapped for Enterprise SSID.

	IsAuthKeySuiteB1X *bool `json:"isAuthKeySuiteB1x,omitempty"` // When activated by setting it to true, the SuiteB-1x authentication key feature is engaged.

	IsAuthKeySuiteB1921X *bool `json:"isAuthKeySuiteB1921x,omitempty"` // When set to true, the SuiteB192-1x authentication key feature is enabled.

	IsAuthKeySaeExt *bool `json:"isAuthKeySaeExt,omitempty"` // When set to true, the Simultaneous Authentication of Equals (SAE) Extended Authentication key feature is turned on.

	IsAuthKeySaeExtPlusFT *bool `json:"isAuthKeySaeExtPlusFT,omitempty"` // When set to true, the Simultaneous Authentication of Equals (SAE) combined with Fast Transition (FT) Authentication Key feature is enabled.

	IsApBeaconProtectionEnabled *bool `json:"isApBeaconProtectionEnabled,omitempty"` // When set to true, the Access Point (AP) Beacon Protection feature is activated, enhancing the security of the network.

	Ghz24Policy string `json:"ghz24Policy,omitempty"` // 2.4 Ghz Band Policy value. Allowed only when 2.4 Radio Band is enabled in ssidRadioType

	CckmTsfTolerance *int `json:"cckmTsfTolerance,omitempty"` // he default value is the Cckm Timestamp Tolerance (in milliseconds, if specified); otherwise, it is 0.

	IsCckmEnabled *bool `json:"isCckmEnabled,omitempty"` // True if CCKM is enabled, else False

	IsHex *bool `json:"isHex,omitempty"` // True if passphrase is in Hex format, else False.

	IsRandomMacFilterEnabled *bool `json:"isRandomMacFilterEnabled,omitempty"` // Deny clients using randomized MAC addresses when set to true

	FastTransitionOverTheDistributedSystemEnable *bool `json:"fastTransitionOverTheDistributedSystemEnable,omitempty"` // Enable Fast Transition over the Distributed System when set to true
}
type RequestWirelessCreateSSIDV1MultipSKSettings struct {
	Priority *int `json:"priority,omitempty"` // Priority

	PassphraseType string `json:"passphraseType,omitempty"` // Passphrase Type(default: ASCII)

	Passphrase string `json:"passphrase,omitempty"` // Passphrase needs to be between 8 and 63 characters for ASCII type. HEX passphrase needs to be 64 characters
}
type RequestWirelessUpdateSSIDV1 struct {
	SSID string `json:"ssid,omitempty"` // Name of the SSID

	AuthType string `json:"authType,omitempty"` // L2 Authentication Type (If authType is not open , then atleast one RSN Cipher Suite and corresponding valid AKM must be enabled). Default is L2 Authentication Type if exists else .

	Passphrase string `json:"passphrase,omitempty"` // Passphrase (Only applicable for SSID with PERSONAL security level). Passphrase needs to be between 8 and 63 characters for ASCII type. HEX passphrase needs to be 64 characters

	IsFastLaneEnabled *bool `json:"isFastLaneEnabled,omitempty"` // True if FastLane is enabled, else False

	IsMacFilteringEnabled *bool `json:"isMacFilteringEnabled,omitempty"` // When set to true, MAC Filtering will be activated, allowing control over network access based on the MAC address of the device

	SSIDRadioType string `json:"ssidRadioType,omitempty"` // Radio Policy Enum (default: Triple band operation(2.4GHz, 5GHz and 6GHz))

	IsBroadcastSSID *bool `json:"isBroadcastSSID,omitempty"` // When activated by setting it to true, the Broadcast SSID feature will make the SSID publicly visible to wireless devices searching for available networks

	FastTransition string `json:"fastTransition,omitempty"` // Fast Transition

	SessionTimeOutEnable *bool `json:"sessionTimeOutEnable,omitempty"` // Turn on the feature that imposes a time limit on user sessions

	SessionTimeOut *int `json:"sessionTimeOut,omitempty"` // This denotes the allotted time span, expressed in seconds, before a session is automatically terminated due to inactivity. Default sessionTimeOut is 1800.

	ClientExclusionEnable *bool `json:"clientExclusionEnable,omitempty"` // Activate the feature that allows for the exclusion of clients

	ClientExclusionTimeout *int `json:"clientExclusionTimeout,omitempty"` // This refers to the length of time, in seconds, a client is excluded or blocked from accessing the network after a specified number of unsuccessful attempts. Default is Client Exclusion Timeout if exists else 180.

	BasicServiceSetMaxIDleEnable *bool `json:"basicServiceSetMaxIdleEnable,omitempty"` // Activate the maximum idle feature for the Basic Service Set

	BasicServiceSetClientIDleTimeout *int `json:"basicServiceSetClientIdleTimeout,omitempty"` // This refers to the duration of inactivity, measured in seconds, before a client connected to the Basic Service Set is considered idle and timed out. Default is Basic ServiceSet ClientIdle Timeout if exists else 300.

	DirectedMulticastServiceEnable *bool `json:"directedMulticastServiceEnable,omitempty"` // The Directed Multicast Service feature becomes operational when it is set to true

	NeighborListEnable *bool `json:"neighborListEnable,omitempty"` // The Neighbor List feature is enabled when it is set to true

	ManagementFrameProtectionClientprotection string `json:"managementFrameProtectionClientprotection,omitempty"` // Default is Management Frame Protection Client if exists else Optional.

	NasOptions []string `json:"nasOptions,omitempty"` // Pre-Defined NAS Options : AP ETH Mac Address, AP IP address, AP Location , AP MAC Address, AP Name, AP Policy Tag, AP Site Tag, SSID, System IP Address, System MAC Address, System Name.

	ProfileName string `json:"profileName,omitempty"` // WLAN Profile Name, if not passed autogenerated profile name will be assigned. The same wlanProfileName will also be used for policyProfileName

	AAAOverride *bool `json:"aaaOverride,omitempty"` // Activate the AAA Override feature when set to true

	CoverageHoleDetectionEnable *bool `json:"coverageHoleDetectionEnable,omitempty"` // Activate Coverage Hole Detection feature when set to true

	ProtectedManagementFrame string `json:"protectedManagementFrame,omitempty"` // (REQUIRED is applicable for authType WPA3_PERSONAL, WPA3_ENTERPRISE, OPEN_SECURED) and (OPTIONAL/REQUIRED is applicable for authType WPA2_WPA3_PERSONAL and WPA2_WPA3_ENTERPRISE)

	MultipSKSettings *[]RequestWirelessUpdateSSIDV1MultipSKSettings `json:"multiPSKSettings,omitempty"` //

	ClientRateLimit *int `json:"clientRateLimit,omitempty"` // This pertains to the maximum data transfer rate, specified in bits per second, that a client is permitted to achieve. It should be in mutliples of 500 . Default is Client Rate Limit if exists else 0.

	RsnCipherSuiteGcmp256 *bool `json:"rsnCipherSuiteGcmp256,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite GCMP256 encryption protocol is activated

	RsnCipherSuiteCcmp256 *bool `json:"rsnCipherSuiteCcmp256,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite CCMP256 encryption protocol is activated

	RsnCipherSuiteGcmp128 *bool `json:"rsnCipherSuiteGcmp128,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite GCMP128 encryption protocol is activated

	RsnCipherSuiteCcmp128 *bool `json:"rsnCipherSuiteCcmp128,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite CCMP128 encryption protocol is activated

	Ghz6PolicyClientSteering *bool `json:"ghz6PolicyClientSteering,omitempty"` // True if 6 GHz Policy Client Steering is enabled, else False

	IsAuthKey8021X *bool `json:"isAuthKey8021x,omitempty"` // When set to true, the 802.1X authentication key is in use

	IsAuthKey8021XPlusFT *bool `json:"isAuthKey8021xPlusFT,omitempty"` // When set to true, the 802.1X-Plus-FT authentication key is in use

	IsAuthKey8021XSHA256 *bool `json:"isAuthKey8021x_SHA256,omitempty"` // When set to true, the feature that enables 802.1X authentication using the SHA256 algorithm is turned on

	IsAuthKeySae *bool `json:"isAuthKeySae,omitempty"` // When set to true, the feature enabling the Simultaneous Authentication of Equals (SAE) authentication key is activated

	IsAuthKeySaePlusFT *bool `json:"isAuthKeySaePlusFT,omitempty"` // Activating this setting by switching it to true turns on the authentication key feature that supports both Simultaneous Authentication of Equals (SAE) and Fast Transition (FT)

	IsAuthKeyPSK *bool `json:"isAuthKeyPSK,omitempty"` // When set to true, the Pre-shared Key (PSK) authentication feature is enabled

	IsAuthKeyPSKPlusFT *bool `json:"isAuthKeyPSKPlusFT,omitempty"` // When set to true, the feature that enables the combination of Pre-shared Key (PSK) and Fast Transition (FT) authentication keys is activated

	IsAuthKeyOWE *bool `json:"isAuthKeyOWE,omitempty"` // When set to true, the Opportunistic Wireless Encryption (OWE) authentication key feature is turned on

	IsAuthKeyEasyPSK *bool `json:"isAuthKeyEasyPSK,omitempty"` // When set to true, the feature that enables the use of Easy Pre-shared Key (PSK) authentication is activated

	IsAuthKeyPSKSHA256 *bool `json:"isAuthKeyPSKSHA256,omitempty"` // The feature that allows the use of Pre-shared Key (PSK) authentication with the SHA256 algorithm is enabled when it is set to true

	OpenSSID string `json:"openSsid,omitempty"` // Open SSID which is already created in the design and not associated to any other OPEN-SECURED SSID

	WLANBandSelectEnable *bool `json:"wlanBandSelectEnable,omitempty"` // Band select is allowed only when band options selected contains at least 2.4 GHz and 5 GHz band else false.

	IsEnabled *bool `json:"isEnabled,omitempty"` // Set SSID's admin status as 'Enabled' when set to true

	AuthServers []string `json:"authServers,omitempty"` // List of Authentication/Authorization server IpAddresses

	AcctServers []string `json:"acctServers,omitempty"` // List of Accounting server IpAddresses

	EgressQos string `json:"egressQos,omitempty"` // Egress QOS

	IngressQos string `json:"ingressQos,omitempty"` // Ingress QOS

	WLANType string `json:"wlanType,omitempty"` // Wlan Type

	L3AuthType string `json:"l3AuthType,omitempty"` // Default is L3 Authentication Type if exists else .

	AuthServer string `json:"authServer,omitempty"` // Authentication Server, Mandatory for Guest SSIDs with wlanType=Guest and l3AuthType=web_auth

	ExternalAuthIPAddress string `json:"externalAuthIpAddress,omitempty"` // External WebAuth URL (Mandatory for Guest SSIDs with wlanType = Guest, l3AuthType = web_auth and authServer = auth_external)

	WebPassthrough *bool `json:"webPassthrough,omitempty"` // When set to true, the Web-Passthrough feature will be activated for the Guest SSID, allowing guests to bypass certain login requirements

	SleepingClientEnable *bool `json:"sleepingClientEnable,omitempty"` // When set to true, this will activate the timeout settings that apply to clients in sleep mode

	SleepingClientTimeout *int `json:"sleepingClientTimeout,omitempty"` // This refers to the amount of time, measured in minutes, before a sleeping (inactive) client is timed out of the network. Default is Sleeping Client Timeout if exists else 720.

	ACLName string `json:"aclName,omitempty"` // Pre-Auth Access Control List (ACL) Name

	IsPosturingEnabled *bool `json:"isPosturingEnabled,omitempty"` // Applicable only for Enterprise SSIDs. When set to True, Posturing will enabled. Required to be set to True if ACL needs to be mapped for Enterprise SSID.

	IsAuthKeySuiteB1X *bool `json:"isAuthKeySuiteB1x,omitempty"` // When activated by setting it to true, the SuiteB-1x authentication key feature is engaged.

	IsAuthKeySuiteB1921X *bool `json:"isAuthKeySuiteB1921x,omitempty"` // When set to true, the SuiteB192-1x authentication key feature is enabled.

	IsAuthKeySaeExt *bool `json:"isAuthKeySaeExt,omitempty"` // When set to true, the Simultaneous Authentication of Equals (SAE) Extended Authentication key feature is turned on.

	IsAuthKeySaeExtPlusFT *bool `json:"isAuthKeySaeExtPlusFT,omitempty"` // When set to true, the Simultaneous Authentication of Equals (SAE) combined with Fast Transition (FT) Authentication Key feature is enabled.

	IsApBeaconProtectionEnabled *bool `json:"isApBeaconProtectionEnabled,omitempty"` // When set to true, the Access Point (AP) Beacon Protection feature is activated, enhancing the security of the network.

	Ghz24Policy string `json:"ghz24Policy,omitempty"` // 2.4 Ghz Band Policy value. Allowed only when 2.4 Radio Band is enabled in ssidRadioType

	CckmTsfTolerance *int `json:"cckmTsfTolerance,omitempty"` // he default value is the Cckm Timestamp Tolerance (in milliseconds, if specified); otherwise, it is 0.

	IsCckmEnabled *bool `json:"isCckmEnabled,omitempty"` // True if CCKM is enabled, else False

	IsHex *bool `json:"isHex,omitempty"` // True if passphrase is in Hex format, else False.

	IsRandomMacFilterEnabled *bool `json:"isRandomMacFilterEnabled,omitempty"` // Deny clients using randomized MAC addresses when set to true

	FastTransitionOverTheDistributedSystemEnable *bool `json:"fastTransitionOverTheDistributedSystemEnable,omitempty"` // Enable Fast Transition over the Distributed System when set to true
}
type RequestWirelessUpdateSSIDV1MultipSKSettings struct {
	Priority *int `json:"priority,omitempty"` // Priority

	PassphraseType string `json:"passphraseType,omitempty"` // Passphrase Type(default: ASCII)

	Passphrase string `json:"passphrase,omitempty"` // Passphrase needs to be between 8 and 63 characters for ASCII type. HEX passphrase needs to be 64 characters
}
type RequestWirelessUpdateOrOverridessidV1 struct {
	SSID string `json:"ssid,omitempty"` // Name of the SSID

	AuthType string `json:"authType,omitempty"` // L2 Authentication Type (If authType is not open , then atleast one RSN Cipher Suite and corresponding valid AKM must be enabled)

	Passphrase string `json:"passphrase,omitempty"` // Passphrase (Only applicable for SSID with PERSONAL security level). Passphrase needs to be between 8 and 63 characters for ASCII type. HEX passphrase needs to be 64 characters

	IsFastLaneEnabled *bool `json:"isFastLaneEnabled,omitempty"` // True if FastLane is enabled, else False

	IsMacFilteringEnabled *bool `json:"isMacFilteringEnabled,omitempty"` // When set to true, MAC Filtering will be activated, allowing control over network access based on the MAC address of the device

	SSIDRadioType string `json:"ssidRadioType,omitempty"` // Radio Policy Enum (default: Triple band operation(2.4GHz, 5GHz and 6GHz))

	IsBroadcastSSID *bool `json:"isBroadcastSSID,omitempty"` // When activated by setting it to true, the Broadcast SSID feature will make the SSID publicly visible to wireless devices searching for available networks

	FastTransition string `json:"fastTransition,omitempty"` // Fast Transition

	SessionTimeOutEnable *bool `json:"sessionTimeOutEnable,omitempty"` // Turn on the feature that imposes a time limit on user sessions

	SessionTimeOut *int `json:"sessionTimeOut,omitempty"` // This denotes the allotted time span, expressed in seconds, before a session is automatically terminated due to inactivity. Default sessionTimeOut is 1800.

	ClientExclusionEnable *bool `json:"clientExclusionEnable,omitempty"` // Activate the feature that allows for the exclusion of clients

	ClientExclusionTimeout *int `json:"clientExclusionTimeout,omitempty"` // This refers to the length of time, in seconds, a client is excluded or blocked from accessing the network after a specified number of unsuccessful attempts

	BasicServiceSetMaxIDleEnable *bool `json:"basicServiceSetMaxIdleEnable,omitempty"` // Activate the maximum idle feature for the Basic Service Set

	BasicServiceSetClientIDleTimeout *int `json:"basicServiceSetClientIdleTimeout,omitempty"` // This refers to the duration of inactivity, measured in seconds, before a client connected to the Basic Service Set is considered idle and timed out

	DirectedMulticastServiceEnable *bool `json:"directedMulticastServiceEnable,omitempty"` // The Directed Multicast Service feature becomes operational when it is set to true

	NeighborListEnable *bool `json:"neighborListEnable,omitempty"` // The Neighbor List feature is enabled when it is set to true

	ManagementFrameProtectionClientprotection string `json:"managementFrameProtectionClientprotection,omitempty"` // Management Frame Protection Client

	NasOptions []string `json:"nasOptions,omitempty"` // Pre-Defined NAS Options : AP ETH Mac Address, AP IP address, AP Location , AP MAC Address, AP Name, AP Policy Tag, AP Site Tag, SSID, System IP Address, System MAC Address, System Name.

	ProfileName string `json:"profileName,omitempty"` // WLAN Profile Name, if not passed autogenerated profile name will be assigned. The same wlanProfileName will also be used for policyProfileName

	AAAOverride *bool `json:"aaaOverride,omitempty"` // Activate the AAA Override feature when set to true

	CoverageHoleDetectionEnable *bool `json:"coverageHoleDetectionEnable,omitempty"` // Activate Coverage Hole Detection feature when set to true

	ProtectedManagementFrame string `json:"protectedManagementFrame,omitempty"` // (REQUIRED is applicable for authType WPA3_PERSONAL, WPA3_ENTERPRISE, OPEN_SECURED) and (OPTIONAL/REQUIRED is applicable for authType WPA2_WPA3_PERSONAL and WPA2_WPA3_ENTERPRISE)

	MultipSKSettings *[]RequestWirelessUpdateOrOverridessidV1MultipSKSettings `json:"multiPSKSettings,omitempty"` //

	ClientRateLimit *int `json:"clientRateLimit,omitempty"` // This pertains to the maximum data transfer rate, specified in bits per second, that a client is permitted to achieve

	RsnCipherSuiteGcmp256 *bool `json:"rsnCipherSuiteGcmp256,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite GCMP256 encryption protocol is activated

	RsnCipherSuiteCcmp256 *bool `json:"rsnCipherSuiteCcmp256,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite CCMP256 encryption protocol is activated

	RsnCipherSuiteGcmp128 *bool `json:"rsnCipherSuiteGcmp128,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite GCMP128 encryption protocol is activated

	RsnCipherSuiteCcmp128 *bool `json:"rsnCipherSuiteCcmp128,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite CCMP128 encryption protocol is activated

	Ghz6PolicyClientSteering *bool `json:"ghz6PolicyClientSteering,omitempty"` // True if 6 GHz Policy Client Steering is enabled, else False

	IsAuthKey8021X *bool `json:"isAuthKey8021x,omitempty"` // When set to true, the 802.1X authentication key is in use

	IsAuthKey8021XPlusFT *bool `json:"isAuthKey8021xPlusFT,omitempty"` // When set to true, the 802.1X-Plus-FT authentication key is in use

	IsAuthKey8021XSHA256 *bool `json:"isAuthKey8021x_SHA256,omitempty"` // When set to true, the feature that enables 802.1X authentication using the SHA256 algorithm is turned on

	IsAuthKeySae *bool `json:"isAuthKeySae,omitempty"` // When set to true, the feature enabling the Simultaneous Authentication of Equals (SAE) authentication key is activated

	IsAuthKeySaePlusFT *bool `json:"isAuthKeySaePlusFT,omitempty"` // Activating this setting by switching it to true turns on the authentication key feature that supports both Simultaneous Authentication of Equals (SAE) and Fast Transition (FT)

	IsAuthKeyPSK *bool `json:"isAuthKeyPSK,omitempty"` // When set to true, the Pre-shared Key (PSK) authentication feature is enabled

	IsAuthKeyPSKPlusFT *bool `json:"isAuthKeyPSKPlusFT,omitempty"` // When set to true, the feature that enables the combination of Pre-shared Key (PSK) and Fast Transition (FT) authentication keys is activated

	IsAuthKeyOWE *bool `json:"isAuthKeyOWE,omitempty"` // When set to true, the Opportunistic Wireless Encryption (OWE) authentication key feature is turned on

	IsAuthKeyEasyPSK *bool `json:"isAuthKeyEasyPSK,omitempty"` // When set to true, the feature that enables the use of Easy Pre-shared Key (PSK) authentication is activated

	IsAuthKeyPSKSHA256 *bool `json:"isAuthKeyPSKSHA256,omitempty"` // The feature that allows the use of Pre-shared Key (PSK) authentication with the SHA256 algorithm is enabled when it is set to true

	OpenSSID string `json:"openSsid,omitempty"` // Open SSID which is already created in the design and not associated to any other OPEN-SECURED SSID

	WLANBandSelectEnable *bool `json:"wlanBandSelectEnable,omitempty"` // Band select is allowed only when band options selected contains at least 2.4 GHz and 5 GHz band

	IsEnabled *bool `json:"isEnabled,omitempty"` // Set SSID's admin status as 'Enabled' when set to true

	AuthServers []string `json:"authServers,omitempty"` // List of Authentication/Authorization server IpAddresses

	AcctServers []string `json:"acctServers,omitempty"` // List of Accounting server IpAddresses

	EgressQos string `json:"egressQos,omitempty"` // Egress QOS

	IngressQos string `json:"ingressQos,omitempty"` // Ingress QOS

	WLANType string `json:"wlanType,omitempty"` // Wlan Type

	L3AuthType string `json:"l3AuthType,omitempty"` // L3 Authentication Type

	AuthServer string `json:"authServer,omitempty"` // Authentication Server, Mandatory for Guest SSIDs with wlanType=Guest and l3AuthType=web_auth

	ExternalAuthIPAddress string `json:"externalAuthIpAddress,omitempty"` // External WebAuth URL (Mandatory for Guest SSIDs with wlanType = Guest, l3AuthType = web_auth and authServer = auth_external)

	WebPassthrough *bool `json:"webPassthrough,omitempty"` // When set to true, the Web-Passthrough feature will be activated for the Guest SSID, allowing guests to bypass certain login requirements

	SleepingClientEnable *bool `json:"sleepingClientEnable,omitempty"` // When set to true, this will activate the timeout settings that apply to clients in sleep mode

	SleepingClientTimeout *int `json:"sleepingClientTimeout,omitempty"` // This refers to the amount of time, measured in minutes, before a sleeping (inactive) client is timed out of the network

	ACLName string `json:"aclName,omitempty"` // Pre-Auth Access Control List (ACL) Name

	IsPosturingEnabled *bool `json:"isPosturingEnabled,omitempty"` // Applicable only for Enterprise SSIDs. When set to True, Posturing will enabled. Required to be set to True if ACL needs to be mapped for Enterprise SSID.

	IsAuthKeySuiteB1X *bool `json:"isAuthKeySuiteB1x,omitempty"` // When activated by setting it to true, the SuiteB-1x authentication key feature is engaged.

	IsAuthKeySuiteB1921X *bool `json:"isAuthKeySuiteB1921x,omitempty"` // When set to true, the SuiteB192-1x authentication key feature is enabled.

	IsAuthKeySaeExt *bool `json:"isAuthKeySaeExt,omitempty"` // When set to true, the Simultaneous Authentication of Equals (SAE) Extended Authentication key feature is turned on.

	IsAuthKeySaeExtPlusFT *bool `json:"isAuthKeySaeExtPlusFT,omitempty"` // When set to true, the Simultaneous Authentication of Equals (SAE) combined with Fast Transition (FT) Authentication Key feature is enabled.

	IsApBeaconProtectionEnabled *bool `json:"isApBeaconProtectionEnabled,omitempty"` // When set to true, the Access Point (AP) Beacon Protection feature is activated, enhancing the security of the network.

	Ghz24Policy string `json:"ghz24Policy,omitempty"` // 2.4 Ghz Band Policy value. Allowed only when 2.4 Radio Band is enabled in ssidRadioType

	CckmTsfTolerance *int `json:"cckmTsfTolerance,omitempty"` // The default value is the Cckm Timestamp Tolerance (in milliseconds, if specified); otherwise, it is 0.

	IsCckmEnabled *bool `json:"isCckmEnabled,omitempty"` // True if CCKM is enabled, else False

	IsHex *bool `json:"isHex,omitempty"` // True if passphrase is in Hex format, else False.

	IsRandomMacFilterEnabled *bool `json:"isRandomMacFilterEnabled,omitempty"` // Deny clients using randomized MAC addresses when set to true

	FastTransitionOverTheDistributedSystemEnable *bool `json:"fastTransitionOverTheDistributedSystemEnable,omitempty"` // Enable Fast Transition over the Distributed System when set to true

	IsRadiusProfilingEnabled *bool `json:"isRadiusProfilingEnabled,omitempty"` // Enable Radius Profiling. At least one AAA/PSN server is required to enable Radius Profiling on WLAN.
}
type RequestWirelessUpdateOrOverridessidV1MultipSKSettings struct {
	Priority *int `json:"priority,omitempty"` // Priority

	PassphraseType string `json:"passphraseType,omitempty"` // Passphrase Type(default:ASCII)

	Passphrase string `json:"passphrase,omitempty"` // Passphrase needs to be between 8 and 63 characters for ASCII type. HEX passphrase needs to be 64 characters
}
type RequestWirelessConfigureAccessPointsV1 struct {
	ApList *[]RequestWirelessConfigureAccessPointsV1ApList `json:"apList,omitempty"` //

	ConfigureAdminStatus *bool `json:"configureAdminStatus,omitempty"` // To change the access point's admin status, set this parameter's value to "true".

	AdminStatus *bool `json:"adminStatus,omitempty"` // Configure the access point's admin status. Set this parameter's value to "true" to enable it and "false" to disable it.

	ConfigureApMode *bool `json:"configureApMode,omitempty"` // To change the access point's mode, set this parameter's value to "true".

	ApMode *int `json:"apMode,omitempty"` // Configure the access point's mode: for local/flexconnect mode, set "0"; for monitor mode, set "1"; for sniffer mode, set "4"; and for bridge/flex+bridge mode, set "5".

	ConfigureFailoverPriority *bool `json:"configureFailoverPriority,omitempty"` // To change the access point's failover priority, set this parameter's value to "true".

	FailoverPriority *int `json:"failoverPriority,omitempty"` // Configure the acess point's failover priority: for low, set "1"; for medium, set "2"; for high, set "3"; and for critical, set "4".

	ConfigureLedStatus *bool `json:"configureLedStatus,omitempty"` // To change the access point's LED status, set this parameter's value to "true".

	LedStatus *bool `json:"ledStatus,omitempty"` // Configure the access point's LED status. Set "true" to enable its status and "false" to disable it.

	ConfigureLedBrightnessLevel *bool `json:"configureLedBrightnessLevel,omitempty"` // To change the access point's LED brightness level, set this parameter's value to "true".

	LedBrightnessLevel *int `json:"ledBrightnessLevel,omitempty"` // Configure the access point's LED brightness level by setting a value between 1 and 8.

	ConfigureLocation *bool `json:"configureLocation,omitempty"` // To change the access point's location, set this parameter's value to "true".

	Location string `json:"location,omitempty"` // Configure the access point's location.

	ConfigureHAController *bool `json:"configureHAController,omitempty"` // To change the access point's HA controller, set this parameter's value to "true".

	PrimaryControllerName string `json:"primaryControllerName,omitempty"` // Configure the hostname for an access point's primary controller.

	PrimaryIPAddress *RequestWirelessConfigureAccessPointsV1PrimaryIPAddress `json:"primaryIpAddress,omitempty"` //

	SecondaryControllerName string `json:"secondaryControllerName,omitempty"` // Configure the hostname for an access point's secondary controller.

	SecondaryIPAddress *RequestWirelessConfigureAccessPointsV1SecondaryIPAddress `json:"secondaryIpAddress,omitempty"` //

	TertiaryControllerName string `json:"tertiaryControllerName,omitempty"` // Configure the hostname for an access point's tertiary controller.

	TertiaryIPAddress *RequestWirelessConfigureAccessPointsV1TertiaryIPAddress `json:"tertiaryIpAddress,omitempty"` //

	RadioConfigurations *[]RequestWirelessConfigureAccessPointsV1RadioConfigurations `json:"radioConfigurations,omitempty"` //

	IsAssignedSiteAsLocation *bool `json:"isAssignedSiteAsLocation,omitempty"` // If AP is assigned to a site, then to assign AP location as the site name, set this parameter's value to "true".
}
type RequestWirelessConfigureAccessPointsV1ApList struct {
	ApName string `json:"apName,omitempty"` // The current host name of the access point.

	MacAddress string `json:"macAddress,omitempty"` // The ethernet MAC address of the access point.

	ApNameNew string `json:"apNameNew,omitempty"` // The modified hostname of the access point.
}
type RequestWirelessConfigureAccessPointsV1PrimaryIPAddress struct {
	Address string `json:"address,omitempty"` // Configure the IP address for an access point's primary controller.
}
type RequestWirelessConfigureAccessPointsV1SecondaryIPAddress struct {
	Address string `json:"address,omitempty"` // Configure the IP address for an access point's secondary controller.
}
type RequestWirelessConfigureAccessPointsV1TertiaryIPAddress struct {
	Address string `json:"address,omitempty"` // Configure the IP address for an access point's tertiary controller.
}
type RequestWirelessConfigureAccessPointsV1RadioConfigurations struct {
	ConfigureRadioRoleAssignment *bool `json:"configureRadioRoleAssignment,omitempty"` // To change the radio role on the specified radio for an access point, set this parameter's value to "true".

	RadioRoleAssignment string `json:"radioRoleAssignment,omitempty"` // Configure only one of the following roles on the specified radio for an access point as "AUTO", "SERVING", or "MONITOR". Any other string is invalid, including empty string

	RadioBand string `json:"radioBand,omitempty"` // Configure the band on the specified radio for an access point: for 2.4 GHz, set "RADIO24"; for 5 GHz, set "RADIO5". Any other string is invalid, including empty string

	ConfigureAdminStatus *bool `json:"configureAdminStatus,omitempty"` // To change the admin status on the specified radio for an access point, set this parameter's value to "true".

	AdminStatus *bool `json:"adminStatus,omitempty"` // Configure the admin status on the specified radio for an access point. Set this parameter's value to "true" to enable it and "false" to disable it.

	ConfigureAntennaPatternName *bool `json:"configureAntennaPatternName,omitempty"` // To change the antenna gain on the specified radio for an access point, set the value for this parameter to "true".

	AntennaPatternName string `json:"antennaPatternName,omitempty"` // Specify the antenna name on the specified radio for an access point. The antenna name is used to calculate the gain on the radio slot.

	AntennaGain *int `json:"antennaGain,omitempty"` // Configure the antenna gain on the specified radio for an access point by setting a decimal value (in dBi). To configure "antennaGain", set "antennaPatternName" value to "other".

	ConfigureAntennaCable *bool `json:"configureAntennaCable,omitempty"` // To change the antenna cable name on the specified radio for an access point, set this parameter's value to "true".

	AntennaCableName string `json:"antennaCableName,omitempty"` // Configure the antenna cable name on the specified radio for an access point. If cable loss needs to be configured, set this parameter's value to "other".

	CableLoss *float64 `json:"cableLoss,omitempty"` // Configure the cable loss on the specified radio for an access point by setting a decimal value (in dBi).

	ConfigureChannel *bool `json:"configureChannel,omitempty"` // To change the channel on the specified radio for an access point, set this parameter's value to "true".

	ChannelAssignmentMode *int `json:"channelAssignmentMode,omitempty"` // Configure the channel assignment mode on the specified radio for an access point: for global mode, set "1"; and for custom mode, set "2".

	ChannelNumber *int `json:"channelNumber,omitempty"` // Configure the channel number on the specified radio for an access point.

	ConfigureChannelWidth *bool `json:"configureChannelWidth,omitempty"` // To change the channel width on the specified radio for an access point, set this parameter's value to "true".

	ChannelWidth *int `json:"channelWidth,omitempty"` // Configure the channel width on the specified radio for an access point: for 20 MHz, set "3"; for 40 MHz, set "4"; for 80 MHz, set "5"; for 160 MHz, set "6", and for 320 MHz, set "7".

	ConfigurePower *bool `json:"configurePower,omitempty"` // To change the power assignment mode on the specified radio for an access point, set this parameter's value to "true".

	PowerAssignmentMode *int `json:"powerAssignmentMode,omitempty"` // Configure the power assignment mode on the specified radio for an access point: for global mode, set "1"; and for custom mode, set "2".

	Powerlevel *int `json:"powerlevel,omitempty"` // Configure the power level on the specified radio for an access point by setting a value between 1 and 8.

	ConfigureCleanAirSI *bool `json:"configureCleanAirSI,omitempty"` // To enable or disable either CleanAir or Spectrum Intelligence on the specified radio for an access point, set this parameter's value to "true".

	CleanAirSI *int `json:"cleanAirSI,omitempty"` // Configure CleanAir or Spectrum Intelligence on the specified radio for an access point. Set this parameter's value to "0" to disable the feature or "1" to enable it.

	RadioType *int `json:"radioType,omitempty"` // Configure an access point's radio band: for 2.4 GHz, set "1"; for 5 GHz, set "2"; for XOR, set "3"; and for 6 GHz, set "6".
}
type RequestWirelessApProvisionConnectivityV1 []RequestItemWirelessApProvisionConnectivityV1 // Array of RequestWirelessAPProvisionConnectivityV1
type RequestItemWirelessApProvisionConnectivityV1 struct {
	RfProfile string `json:"rfProfile,omitempty"` // Radio frequency profile name

	DeviceName string `json:"deviceName,omitempty"` // Device name

	CustomApGroupName string `json:"customApGroupName,omitempty"` // Custom AP group name

	CustomFlexGroupName []string `json:"customFlexGroupName,omitempty"` // ["Custom flex group name"]

	Type string `json:"type,omitempty"` // ApWirelessConfiguration

	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Site name hierarchy(ex: Global/...)
}
type RequestWirelessCreateUpdateDynamicInterfaceV1 struct {
	InterfaceName string `json:"interfaceName,omitempty"` // dynamic-interface name

	VLANID *float64 `json:"vlanId,omitempty"` // Vlan Id
}
type RequestWirelessUpdateWirelessProfileV1 struct {
	ProfileDetails *RequestWirelessUpdateWirelessProfileV1ProfileDetails `json:"profileDetails,omitempty"` //
}
type RequestWirelessUpdateWirelessProfileV1ProfileDetails struct {
	Name string `json:"name,omitempty"` // Profile Name

	Sites []string `json:"sites,omitempty"` // array of site name hierarchies(eg: ["Global/aaa/zzz", "Global/aaa/zzz"])

	SSIDDetails *[]RequestWirelessUpdateWirelessProfileV1ProfileDetailsSSIDDetails `json:"ssidDetails,omitempty"` //
}
type RequestWirelessUpdateWirelessProfileV1ProfileDetailsSSIDDetails struct {
	Name string `json:"name,omitempty"` // Ssid Name

	EnableFabric *bool `json:"enableFabric,omitempty"` // true if ssid is fabric else false

	FlexConnect *RequestWirelessUpdateWirelessProfileV1ProfileDetailsSSIDDetailsFlexConnect `json:"flexConnect,omitempty"` //

	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name

	WLANProfileName string `json:"wlanProfileName,omitempty"` // WLAN Profile Name

	PolicyProfileName string `json:"policyProfileName,omitempty"` // Policy Profile Name
}
type RequestWirelessUpdateWirelessProfileV1ProfileDetailsSSIDDetailsFlexConnect struct {
	EnableFlexConnect *bool `json:"enableFlexConnect,omitempty"` // true if flex connect is enabled else false

	LocalToVLAN *int `json:"localToVlan,omitempty"` // Local to VLAN ID. Required if enableFlexConnect is true.
}
type RequestWirelessCreateWirelessProfileV1 struct {
	ProfileDetails *RequestWirelessCreateWirelessProfileV1ProfileDetails `json:"profileDetails,omitempty"` //
}
type RequestWirelessCreateWirelessProfileV1ProfileDetails struct {
	Name string `json:"name,omitempty"` // Profile Name

	Sites []string `json:"sites,omitempty"` // array of site name hierarchies(eg: ["Global/aaa/zzz", "Global/aaa/zzz"])

	SSIDDetails *[]RequestWirelessCreateWirelessProfileV1ProfileDetailsSSIDDetails `json:"ssidDetails,omitempty"` //
}
type RequestWirelessCreateWirelessProfileV1ProfileDetailsSSIDDetails struct {
	Name string `json:"name,omitempty"` // Ssid Name

	EnableFabric *bool `json:"enableFabric,omitempty"` // true if ssid is fabric else false

	FlexConnect *RequestWirelessCreateWirelessProfileV1ProfileDetailsSSIDDetailsFlexConnect `json:"flexConnect,omitempty"` //

	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name

	WLANProfileName string `json:"wlanProfileName,omitempty"` // WLAN Profile Name

	PolicyProfileName string `json:"policyProfileName,omitempty"` // Policy Profile Name
}
type RequestWirelessCreateWirelessProfileV1ProfileDetailsSSIDDetailsFlexConnect struct {
	EnableFlexConnect *bool `json:"enableFlexConnect,omitempty"` // true if flex connect is enabled else false

	LocalToVLAN *int `json:"localToVlan,omitempty"` // Local to VLAN ID. Required if enableFlexConnect is true.
}
type RequestWirelessProvisionUpdateV1 []RequestItemWirelessProvisionUpdateV1 // Array of RequestWirelessProvisionUpdateV1
type RequestItemWirelessProvisionUpdateV1 struct {
	DeviceName string `json:"deviceName,omitempty"` // Controller Name

	ManagedApLocations []string `json:"managedAPLocations,omitempty"` // List of managed AP locations (Site Hierarchies)

	DynamicInterfaces *[]RequestItemWirelessProvisionUpdateV1DynamicInterfaces `json:"dynamicInterfaces,omitempty"` //
}
type RequestItemWirelessProvisionUpdateV1DynamicInterfaces struct {
	InterfaceIPAddress string `json:"interfaceIPAddress,omitempty"` // Interface IP Address. Required for AireOS.

	InterfaceNetmaskInCIDR *int `json:"interfaceNetmaskInCIDR,omitempty"` // Interface Netmask In CIDR. Required for AireOS.

	InterfaceGateway string `json:"interfaceGateway,omitempty"` // Interface Gateway. Required for AireOS.

	LagOrPortNumber *int `json:"lagOrPortNumber,omitempty"` // Lag Or Port Number. Required for AireOS.

	VLANID *int `json:"vlanId,omitempty"` // VLAN ID. Required for AireOS and EWLC.

	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name. Required for AireOS and EWLC.
}
type RequestWirelessProvisionV1 []RequestItemWirelessProvisionV1 // Array of RequestWirelessProvisionV1
type RequestItemWirelessProvisionV1 struct {
	DeviceName string `json:"deviceName,omitempty"` // Controller Name

	Site string `json:"site,omitempty"` // Full Site Hierarchy where device has to be assigned

	ManagedApLocations []string `json:"managedAPLocations,omitempty"` // List of managed AP locations (Site Hierarchies)

	DynamicInterfaces *[]RequestItemWirelessProvisionV1DynamicInterfaces `json:"dynamicInterfaces,omitempty"` //
}
type RequestItemWirelessProvisionV1DynamicInterfaces struct {
	InterfaceIPAddress string `json:"interfaceIPAddress,omitempty"` // Interface IP Address. Required for AireOS.

	InterfaceNetmaskInCIDR *int `json:"interfaceNetmaskInCIDR,omitempty"` // Interface Netmask In CIDR. Required for AireOS.

	InterfaceGateway string `json:"interfaceGateway,omitempty"` // Interface Gateway.  Required for AireOS.

	LagOrPortNumber *int `json:"lagOrPortNumber,omitempty"` // Lag Or Port Number.  Required for AireOS.

	VLANID *int `json:"vlanId,omitempty"` // VLAN ID. Required for both AireOS and EWLC.

	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name. Required for both AireOS and EWLC.
}
type RequestWirelessPSKOverrideV1 []RequestItemWirelessPSKOverrideV1 // Array of RequestWirelessPSKOverrideV1
type RequestItemWirelessPSKOverrideV1 struct {
	SSID string `json:"ssid,omitempty"` // enterprise ssid name(already created/present)

	Site string `json:"site,omitempty"` // site name hierarchy (ex: Global/aaa/zzz/...)

	PassPhrase string `json:"passPhrase,omitempty"` // Pass phrase (create/update)

	WLANProfileName string `json:"wlanProfileName,omitempty"` // WLAN Profile Name
}
type RequestWirelessCreateOrUpdateRfProfileV1 struct {
	Name string `json:"name,omitempty"` // RF Profile Name

	DefaultRfProfile *bool `json:"defaultRfProfile,omitempty"` // is Default Rf Profile

	EnableRadioTypeA *bool `json:"enableRadioTypeA,omitempty"` // Enable Radio Type A

	EnableRadioTypeB *bool `json:"enableRadioTypeB,omitempty"` // Enable Radio Type B

	ChannelWidth string `json:"channelWidth,omitempty"` // Channel Width

	EnableCustom *bool `json:"enableCustom,omitempty"` // Enable Custom

	EnableBrownField *bool `json:"enableBrownField,omitempty"` // Enable Brown Field

	RadioTypeAProperties *RequestWirelessCreateOrUpdateRfProfileV1RadioTypeAProperties `json:"radioTypeAProperties,omitempty"` //

	RadioTypeBProperties *RequestWirelessCreateOrUpdateRfProfileV1RadioTypeBProperties `json:"radioTypeBProperties,omitempty"` //

	RadioTypeCProperties *RequestWirelessCreateOrUpdateRfProfileV1RadioTypeCProperties `json:"radioTypeCProperties,omitempty"` //

	EnableRadioTypeC *bool `json:"enableRadioTypeC,omitempty"` // Enable Radio Type C (6GHz)
}
type RequestWirelessCreateOrUpdateRfProfileV1RadioTypeAProperties struct {
	ParentProfile string `json:"parentProfile,omitempty"` // Parent Profile (Default : CUSTOM)

	RadioChannels string `json:"radioChannels,omitempty"` // Radio Channels (Default : "36,40,44,48,52,56,60,64,149,153,157,161")

	DataRates string `json:"dataRates,omitempty"` // Data Rates (Default : "6,9,12,18,24,36,48,54")

	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory Data Rates (Default: "6,12,24")

	PowerThresholdV1 *float64 `json:"powerThresholdV1,omitempty"` // Power Threshold V1 ( (Default: -70)

	RxSopThreshold string `json:"rxSopThreshold,omitempty"` // Rx Sop Threshold  (Default: "AUTO")

	MinPowerLevel *float64 `json:"minPowerLevel,omitempty"` // Rx Sop Threshold  (Default: -10)

	MaxPowerLevel *float64 `json:"maxPowerLevel,omitempty"` // Max Power Level  (Default: 30)
}
type RequestWirelessCreateOrUpdateRfProfileV1RadioTypeBProperties struct {
	ParentProfile string `json:"parentProfile,omitempty"` // Parent Profile (Default : CUSTOM)

	RadioChannels string `json:"radioChannels,omitempty"` // Radio Channels (Default : "9,11,12,18,24,36,48,54")

	DataRates string `json:"dataRates,omitempty"` // Data Rates  (Default: "9,11,12,18,24,36,48,54")

	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory Data Rates  (Default: "12")

	PowerThresholdV1 *float64 `json:"powerThresholdV1,omitempty"` // Power Threshold V1  (Default: -70)

	RxSopThreshold string `json:"rxSopThreshold,omitempty"` // Rx Sop Threshold (Default: "AUTO")

	MinPowerLevel *float64 `json:"minPowerLevel,omitempty"` // Min Power Level  (Default: -10)

	MaxPowerLevel *float64 `json:"maxPowerLevel,omitempty"` // Max Power Level  (Default: 30)
}
type RequestWirelessCreateOrUpdateRfProfileV1RadioTypeCProperties struct {
	ParentProfile string `json:"parentProfile,omitempty"` // Parent Profile (Default : CUSTOM)

	RadioChannels string `json:"radioChannels,omitempty"` // Radio Channels (Default : "5,21,37,53,69,85,101,117,133,149,165,181,197,213,229")

	DataRates string `json:"dataRates,omitempty"` // Data Rates  (Default: "6,9,12,18,24,36,48,54")

	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory Data Rates  (Default: "6,12,24")

	RxSopThreshold string `json:"rxSopThreshold,omitempty"` // Rx Sop Threshold  (Default: "AUTO")

	MinPowerLevel *float64 `json:"minPowerLevel,omitempty"` // Min Power Level  (Default: -10)

	MaxPowerLevel *float64 `json:"maxPowerLevel,omitempty"` // Max Power Level  (Default: 30)

	PowerThresholdV1 *float64 `json:"powerThresholdV1,omitempty"` // Power Threshold V1  (Default: -70)
}
type RequestWirelessFactoryResetAccessPointsV1 struct {
	KeepStaticIPConfig *bool `json:"keepStaticIPConfig,omitempty"` // Set the value of keepStaticIPConfig to false, to clear all configurations from Access Points and set the value of keepStaticIPConfig to true, to clear all configurations from Access Points without clearing static IP configuration.

	ApMacAddresses []string `json:"apMacAddresses,omitempty"` // List of Access Point's Ethernet MAC addresses, set maximum 100 ethernet MAC addresses per request.
}
type RequestWirelessApProvisionV1 struct {
	NetworkDevices *[]RequestWirelessApProvisionV1NetworkDevices `json:"networkDevices,omitempty"` //

	RfProfileName string `json:"rfProfileName,omitempty"` // RF Profile Name. RF Profile is not allowed for custom AP Zones.

	ApZoneName string `json:"apZoneName,omitempty"` // AP Zone Name. A custom AP Zone should be passed if no rfProfileName is provided.

	SiteID string `json:"siteId,omitempty"` // Site ID
}
type RequestWirelessApProvisionV1NetworkDevices struct {
	DeviceID string `json:"deviceId,omitempty"` // Network device ID of access points

	MeshRole string `json:"meshRole,omitempty"` // Mesh Role (Applicable only when AP is in Bridge Mode)
}
type RequestWirelessMobilityProvisionV1 struct {
	MobilityGroupName string `json:"mobilityGroupName,omitempty"` // Self device Group Name. Must be alphanumeric without {!,<,space,?/'}  and maximum of 31 characters.

	MacAddress string `json:"macAddress,omitempty"` // Device mobility MAC Address. Allowed formats are: 0a0b.0c01.0211, 0a0b0c010211, 0a:0b:0c:01:02:11

	ManagementIP string `json:"managementIp,omitempty"` // Self device wireless Management IP.

	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Obtain the network device ID value by using the API call GET: /dna/intent/api/v1/network-device/ip-address/${ipAddress}.

	DtlsHighCipher *bool `json:"dtlsHighCipher,omitempty"` // DTLS High Cipher.

	DataLinkEncryption *bool `json:"dataLinkEncryption,omitempty"` // A secure link in which data is encrypted using CAPWAP DTLS protocol can be established between two controllers. This value will be applied to all peers during POST operation.

	MobilityPeers *[]RequestWirelessMobilityProvisionV1MobilityPeers `json:"mobilityPeers,omitempty"` //
}
type RequestWirelessMobilityProvisionV1MobilityPeers struct {
	PeerIP string `json:"peerIp,omitempty"` // This indicates public ip address.

	PrivateIPAddress string `json:"privateIpAddress,omitempty"` // This indicates private/management ip address.

	PeerDeviceName string `json:"peerDeviceName,omitempty"` // Peer device Host Name.

	PeerNetworkDeviceID string `json:"peerNetworkDeviceId,omitempty"` // The possible values are: UNKNOWN or valid UUID of Network device Id. UNKNOWN represents out of band device which is not managed internally. Valid UUID represents WLC network device id.

	MobilityGroupName string `json:"mobilityGroupName,omitempty"` // Peer Device mobility group Name. Must be alphanumeric without {!,<,space,?/'} <br/> and maximum of 31 characters.

	MemberMacAddress string `json:"memberMacAddress,omitempty"` // Peer device mobility MAC Address.  Allowed formats are: 0a0b.0c01.0211, 0a0b0c010211, 0a:0b:0c:01:02:11

	DeviceSeries string `json:"deviceSeries,omitempty"` // Indicates peer device mobility belongs to AireOS or IOX-XE family. 0 - indicates AireOS and 1 - indicates C9800.

	HashKey string `json:"hashKey,omitempty"` // SSC hash string must be 40 characters.
}
type RequestWirelessMobilityResetV1 struct {
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Network device Id of Cisco wireless controller. Obtain the network device ID value by using the API call GET - /dna/intent/api/v1/network-device/ip-address/${ipAddress}.
}
type RequestWirelessAssignManagedApLocationsForWLCV1 struct {
	PrimaryManagedApLocationsSiteIDs []string `json:"primaryManagedAPLocationsSiteIds,omitempty"` // Site IDs of Primary Managed AP Locations. These values can be obtained by using the API call GET: /dna/intent/api/v1/site

	SecondaryManagedApLocationsSiteIDs []string `json:"secondaryManagedAPLocationsSiteIds,omitempty"` // Site IDs of Secondary Managed AP Locations. These values can be obtained by using the API call GET: /dna/intent/api/v1/site
}
type RequestWirelessWirelessControllerProvisionV1 struct {
	Interfaces *[]RequestWirelessWirelessControllerProvisionV1Interfaces `json:"interfaces,omitempty"` //

	SkipApProvision *bool `json:"skipApProvision,omitempty"` // True if Skip AP Provision is enabled, else False

	RollingApUpgrade *RequestWirelessWirelessControllerProvisionV1RollingApUpgrade `json:"rollingApUpgrade,omitempty"` //

	ApAuthorizationListName string `json:"apAuthorizationListName,omitempty"` // AP Authorization List name. 'Obtain the AP Authorization List names by using the API call GET: /intent/api/v1/wirelessSettings/apAuthorizationLists. During re-provision, obtain the AP Authorization List configured for the given provisioned network device Id using the API call GET: /intent/api/v1/wireless/apAuthorizationLists/{networkDeviceId}'

	AuthorizeMeshAndNonMeshAccessPoints *bool `json:"authorizeMeshAndNonMeshAccessPoints,omitempty"` // True if AP Authorization List should  authorize against All Mesh/Non-Mesh APs, else false if AP Authorization List should only authorize against Mesh APs (Applicable only when Mesh is enabled on sites)
}
type RequestWirelessWirelessControllerProvisionV1Interfaces struct {
	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name

	VLANID *int `json:"vlanId,omitempty"` // VLAN ID range is 1 - 4094

	InterfaceIPAddress string `json:"interfaceIPAddress,omitempty"` // Interface IP Address

	InterfaceNetmaskInCIDR *int `json:"interfaceNetmaskInCIDR,omitempty"` // Interface Netmask In CIDR, range is 1-30

	InterfaceGateway string `json:"interfaceGateway,omitempty"` // Interface Gateway

	LagOrPortNumber *int `json:"lagOrPortNumber,omitempty"` // Lag Or Port Number
}
type RequestWirelessWirelessControllerProvisionV1RollingApUpgrade struct {
	EnableRollingApUpgrade *bool `json:"enableRollingApUpgrade,omitempty"` // True if Rolling AP Upgrade is enabled, else False

	ApRebootPercentage *int `json:"apRebootPercentage,omitempty"` // AP Reboot Percentage. Permissible values - 5, 15, 25
}
type RequestWirelessCreateWirelessProfileConnectivityV1 struct {
	WirelessProfileName string `json:"wirelessProfileName,omitempty"` // Wireless Network Profile Name

	SSIDDetails *[]RequestWirelessCreateWirelessProfileConnectivityV1SSIDDetails `json:"ssidDetails,omitempty"` //

	AdditionalInterfaces []string `json:"additionalInterfaces,omitempty"` // These additional interfaces will be configured on the device as independent interfaces in addition to the interfaces mapped to SSIDs. Max Limit 4094

	ApZones *[]RequestWirelessCreateWirelessProfileConnectivityV1ApZones `json:"apZones,omitempty"` //
}
type RequestWirelessCreateWirelessProfileConnectivityV1SSIDDetails struct {
	SSIDName string `json:"ssidName,omitempty"` // SSID Name

	FlexConnect *RequestWirelessCreateWirelessProfileConnectivityV1SSIDDetailsFlexConnect `json:"flexConnect,omitempty"` //

	EnableFabric *bool `json:"enableFabric,omitempty"` // True if fabric is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time

	WLANProfileName string `json:"wlanProfileName,omitempty"` // WLAN Profile Name

	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name.

	Dot11BeProfileID string `json:"dot11beProfileId,omitempty"` // 802.11be Profile Id. Applicable to IOS controllers with version 17.15 and higher. 802.11be Profiles if passed, should be same across all SSIDs in network profile being configured

	AnchorGroupName string `json:"anchorGroupName,omitempty"` // Anchor Group Name

	VLANGroupName string `json:"vlanGroupName,omitempty"` // VLAN Group Name
}
type RequestWirelessCreateWirelessProfileConnectivityV1SSIDDetailsFlexConnect struct {
	EnableFlexConnect *bool `json:"enableFlexConnect,omitempty"` // True if flex connect is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time

	LocalToVLAN *int `json:"localToVlan,omitempty"` // Local to VLAN ID
}
type RequestWirelessCreateWirelessProfileConnectivityV1ApZones struct {
	ApZoneName string `json:"apZoneName,omitempty"` // AP Zone Name

	RfProfileName string `json:"rfProfileName,omitempty"` // RF Profile Name

	SSIDs []string `json:"ssids,omitempty"` // ssids part of apZone
}
type RequestWirelessUpdateWirelessProfileConnectivityV1 struct {
	WirelessProfileName string `json:"wirelessProfileName,omitempty"` // Wireless Network Profile Name

	SSIDDetails *[]RequestWirelessUpdateWirelessProfileConnectivityV1SSIDDetails `json:"ssidDetails,omitempty"` //

	AdditionalInterfaces []string `json:"additionalInterfaces,omitempty"` // These additional interfaces will be configured on the device as independent interfaces in addition to the interfaces mapped to SSIDs. Max Limit 4094

	ApZones *[]RequestWirelessUpdateWirelessProfileConnectivityV1ApZones `json:"apZones,omitempty"` //
}
type RequestWirelessUpdateWirelessProfileConnectivityV1SSIDDetails struct {
	SSIDName string `json:"ssidName,omitempty"` // SSID Name

	FlexConnect *RequestWirelessUpdateWirelessProfileConnectivityV1SSIDDetailsFlexConnect `json:"flexConnect,omitempty"` //

	EnableFabric *bool `json:"enableFabric,omitempty"` // True if fabric is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time

	WLANProfileName string `json:"wlanProfileName,omitempty"` // WLAN Profile Name

	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name.

	Dot11BeProfileID string `json:"dot11beProfileId,omitempty"` // 802.11be Profile Id. Applicable to IOS controllers with version 17.15 and higher. 802.11be Profiles if passed, should be same across all SSIDs in network profile being configured

	AnchorGroupName string `json:"anchorGroupName,omitempty"` // Anchor Group Name

	VLANGroupName string `json:"vlanGroupName,omitempty"` // VLAN Group Name
}
type RequestWirelessUpdateWirelessProfileConnectivityV1SSIDDetailsFlexConnect struct {
	EnableFlexConnect *bool `json:"enableFlexConnect,omitempty"` // True if flex connect is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time

	LocalToVLAN *int `json:"localToVlan,omitempty"` // Local to VLAN ID
}
type RequestWirelessUpdateWirelessProfileConnectivityV1ApZones struct {
	ApZoneName string `json:"apZoneName,omitempty"` // AP Zone Name

	RfProfileName string `json:"rfProfileName,omitempty"` // RF Profile Name

	SSIDs []string `json:"ssids,omitempty"` // ssids part of apZone
}
type RequestWirelessCreateMultiplePolicyTagsForAWirelessProfileInBulkV1 struct {
	Items *[][]string `json:"items,omitempty"` // Items
}
type RequestWirelessUpdateASpecificPolicyTagForAWirelessProfileV1 struct {
	SiteIDs []string `json:"siteIds,omitempty"` // Site Ids

	PolicyTagName string `json:"policyTagName,omitempty"` // Policy Tag Name

	ApZones []string `json:"apZones,omitempty"` // Ap Zones
}
type RequestWirelessCreateMultipleSiteTagsForAWirelessProfileInBulkV1 struct {
	Items *[][]string `json:"items,omitempty"` // Root
}
type RequestWirelessUpdateASpecificSiteTagForAWirelessProfileV1 struct {
	SiteIDs []string `json:"siteIds,omitempty"` // Site Ids

	SiteTagName string `json:"siteTagName,omitempty"` // Use English letters, numbers, special characters except <, /, '.*', ? and leading/trailing space.

	FlexProfileName string `json:"flexProfileName,omitempty"` // Flex Profile Name

	ApProfileName string `json:"apProfileName,omitempty"` // Ap Profile Name
}
type RequestWirelessCreateAnchorGroupV1 struct {
	AnchorGroupName string `json:"anchorGroupName,omitempty"` // Anchor Group Name. Max length is 32 characters

	MobilityAnchors *[]RequestWirelessCreateAnchorGroupV1MobilityAnchors `json:"mobilityAnchors,omitempty"` //
}
type RequestWirelessCreateAnchorGroupV1MobilityAnchors struct {
	DeviceName string `json:"deviceName,omitempty"` // Peer Host Name

	IPAddress string `json:"ipAddress,omitempty"` // This indicates Mobility public IP address. Allowed formats are 192.168.0.1, 10.0.0.1, 255.255.255.255

	AnchorPriority string `json:"anchorPriority,omitempty"` // This indicates anchor priority.  Priority values range from 1 (high) to 3 (low). Primary, secondary or tertiary and defined priority is displayed with guest anchor. Only one priority value is allowed per anchor WLC.

	ManagedAnchorWlc *bool `json:"managedAnchorWlc,omitempty"` // This indicates whether the Wireless LAN Controller supporting Anchor is managed by the Network Controller or not. True means this is managed by Network Controller.

	PeerDeviceType string `json:"peerDeviceType,omitempty"` // Indicates peer device mobility belongs to AireOS or IOS-XE family.

	MacAddress string `json:"macAddress,omitempty"` // Peer Device mobility MAC address. Allowed formats are:0a0b.0c01.0211, 0a0b0c010211, 0a:0b:0c:01:02:11

	MobilityGroupName string `json:"mobilityGroupName,omitempty"` // Peer Device mobility group Name. Must be alphanumeric without {!,<,space,?/'} and maximum of 31 characters.

	PrivateIP string `json:"privateIp,omitempty"` // This indicates private management IP address. Allowed formats are 192.168.0.1, 10.0.0.1, 255.255.255.255
}
type RequestWirelessUpdateAnchorGroupV1 struct {
	AnchorGroupName string `json:"anchorGroupName,omitempty"` // Anchor Group Name. Max length is 32 characters

	MobilityAnchors *[]RequestWirelessUpdateAnchorGroupV1MobilityAnchors `json:"mobilityAnchors,omitempty"` //
}
type RequestWirelessUpdateAnchorGroupV1MobilityAnchors struct {
	DeviceName string `json:"deviceName,omitempty"` // Peer Host Name

	IPAddress string `json:"ipAddress,omitempty"` // This indicates Mobility public IP address. Allowed formats are 192.168.0.1, 10.0.0.1, 255.255.255.255

	AnchorPriority string `json:"anchorPriority,omitempty"` // This indicates anchor priority.  Priority values range from 1 (high) to 3 (low). Primary, secondary or tertiary and defined priority is displayed with guest anchor. Only one priority value is allowed per anchor WLC.

	ManagedAnchorWlc *bool `json:"managedAnchorWlc,omitempty"` // This indicates whether the Wireless LAN Controller supporting Anchor is managed by the Network Controller or not. True means this is managed by Network Controller.

	PeerDeviceType string `json:"peerDeviceType,omitempty"` // Indicates peer device mobility belongs to AireOS or IOS-XE family.

	MacAddress string `json:"macAddress,omitempty"` // Peer Device mobility MAC address. Allowed formats are:0a0b.0c01.0211, 0a0b0c010211, 0a:0b:0c:01:02:11

	MobilityGroupName string `json:"mobilityGroupName,omitempty"` // Peer Device mobility group Name. Must be alphanumeric without {!,<,space,?/'} and maximum of 31 characters.

	PrivateIP string `json:"privateIp,omitempty"` // This indicates private management IP address. Allowed formats are 192.168.0.1, 10.0.0.1, 255.255.255.255
}
type RequestWirelessCreateApAuthorizationListV1 struct {
	ApAuthorizationListName string `json:"apAuthorizationListName,omitempty"` // AP Authorization List Name. For a AP Authorization List to be created successfully, either Local Authorization or Remote Authorization is mandatory.

	LocalAuthorization *RequestWirelessCreateApAuthorizationListV1LocalAuthorization `json:"localAuthorization,omitempty"` //

	RemoteAuthorization *RequestWirelessCreateApAuthorizationListV1RemoteAuthorization `json:"remoteAuthorization,omitempty"` //
}
type RequestWirelessCreateApAuthorizationListV1LocalAuthorization struct {
	ApMacEntries []string `json:"apMacEntries,omitempty"` // List of Access Point's Ethernet MAC addresses. Allowed formats are:0a0b.0c01.0211, 0a0b0c010211, 0a:0b:0c:01:02:11

	ApSerialNumberEntries []string `json:"apSerialNumberEntries,omitempty"` // List of Access Point's Serial Numbers.
}
type RequestWirelessCreateApAuthorizationListV1RemoteAuthorization struct {
	AAAServers []string `json:"aaaServers,omitempty"` // List of Authorization server IpAddresses. Obtain the AAA servers by using the API GET call '/dna/intent/api/v1/authentication-policy-servers'.

	AuthorizeApWithMac *bool `json:"authorizeApWithMac,omitempty"` // True if AP Authorization List should authorise APs With MAC addresses, else False. (For Non-Mesh Access Points, either of Authorize AP With MAC Address or Serial Number is required to be set to true)

	AuthorizeApWithSerialNumber *bool `json:"authorizeApWithSerialNumber,omitempty"` // True if server IpAddresses are added and AP Authorization List should authorise APs With Serial Numbers, else False (For Non-Mesh Access Points, either of Authorize AP With MAC Address or Serial Number is required to be set to true)
}
type RequestWirelessUpdateApAuthorizationListV1 struct {
	ApAuthorizationListName string `json:"apAuthorizationListName,omitempty"` // AP Authorization List Name. For a AP Authorization List to be created successfully, either Local Authorization or Remote Authorization is mandatory.

	LocalAuthorization *RequestWirelessUpdateApAuthorizationListV1LocalAuthorization `json:"localAuthorization,omitempty"` //

	RemoteAuthorization *RequestWirelessUpdateApAuthorizationListV1RemoteAuthorization `json:"remoteAuthorization,omitempty"` //
}
type RequestWirelessUpdateApAuthorizationListV1LocalAuthorization struct {
	ApMacEntries []string `json:"apMacEntries,omitempty"` // List of Access Point's Ethernet MAC addresses. Allowed formats are:0a0b.0c01.0211, 0a0b0c010211, 0a:0b:0c:01:02:11

	ApSerialNumberEntries []string `json:"apSerialNumberEntries,omitempty"` // List of Access Point's Serial Numbers.
}
type RequestWirelessUpdateApAuthorizationListV1RemoteAuthorization struct {
	AAAServers []string `json:"aaaServers,omitempty"` // List of Authorization server IpAddresses. Obtain the AAA servers by using the API GET call '/dna/intent/api/v1/authentication-policy-servers'.

	AuthorizeApWithMac *bool `json:"authorizeApWithMac,omitempty"` // True if AP Authorization List should authorise APs With MAC addresses, else False. (For Non-Mesh Access Points, either of Authorize AP With MAC Address or Serial Number is required to be set to true)

	AuthorizeApWithSerialNumber *bool `json:"authorizeApWithSerialNumber,omitempty"` // True if server IpAddresses are added and AP Authorization List should authorise APs With Serial Numbers, else False (For Non-Mesh Access Points, either of Authorize AP With MAC Address or Serial Number is required to be set to true)
}
type RequestWirelessCreateApProfileV1 struct {
	ApProfileName string `json:"apProfileName,omitempty"` // Name of the Access Point profile. Max length is 32 characters.

	Description string `json:"description,omitempty"` // Description of the AP profile. Max length is 241 characters

	RemoteWorkerEnabled *bool `json:"remoteWorkerEnabled,omitempty"` // Indicates if remote worker mode is enabled on the AP. Remote teleworker enabled profile cannot support security features like aWIPS,Forensic Capture Enablement, Rogue Detection and Rogue Containment.

	ManagementSetting *RequestWirelessCreateApProfileV1ManagementSetting `json:"managementSetting,omitempty"` //

	AwipsEnabled *bool `json:"awipsEnabled,omitempty"` // Indicates if AWIPS is enabled on the AP.

	AwipsForensicEnabled *bool `json:"awipsForensicEnabled,omitempty"` // Indicates if AWIPS forensic is enabled on the AP. Forensic Capture is supported from IOS-XE version 17.4 and above. Forensic Capture can be activated only if aWIPS is enabled.

	RogueDetectionSetting *RequestWirelessCreateApProfileV1RogueDetectionSetting `json:"rogueDetectionSetting,omitempty"` //

	PmfDenialEnabled *bool `json:"pmfDenialEnabled,omitempty"` // Indicates if PMF denial is active on the AP. PMF Denial is supported from IOS-XE version 17.12 and above.

	MeshEnabled *bool `json:"meshEnabled,omitempty"` // This indicates whether mesh networking is enabled on the AP. For IOS-XE devices, when mesh networking is enabled, a custom mesh profile with the configured parameters will be created and mapped to the AP join profile on the device. When mesh networking is disabled, any existing custom mesh profile will be deleted from the device, and the AP join profile will be mapped to the default mesh profile on the device.

	MeshSetting *RequestWirelessCreateApProfileV1MeshSetting `json:"meshSetting,omitempty"` //

	ApPowerProfileName string `json:"apPowerProfileName,omitempty"` // Name of the existing AP power profile.

	CalendarPowerProfiles *RequestWirelessCreateApProfileV1CalendarPowerProfiles `json:"calendarPowerProfiles,omitempty"` //

	CountryCode string `json:"countryCode,omitempty"` // Country Code

	TimeZone string `json:"timeZone,omitempty"` // In the Time Zone area, choose one of the following options.             Not Configured - APs operate in the UTC time zone.             Controller - APs operate in the Cisco Wireless Controller time zone.             Delta from Controller - APs operate in the offset time from the wireless controller time zone.

	TimeZoneOffsetHour *int `json:"timeZoneOffsetHour,omitempty"` // Enter the hour value (HH). The valid range is from -12 through 14.

	TimeZoneOffsetMinutes *int `json:"timeZoneOffsetMinutes,omitempty"` // Enter the minute value (MM). The valid range is from 0 through 59.

	ClientLimit *int `json:"clientLimit,omitempty"` // Number of clients. Value should be between 0-1200.
}
type RequestWirelessCreateApProfileV1ManagementSetting struct {
	AuthType string `json:"authType,omitempty"` // Authentication type used in the AP profile. These setting are applicable during PnP claim and for day-N authentication of AP. Changing these settings will be service impacting for the PnP onboarded APs and will need a factory-reset for those APs.

	Dot1XUsername string `json:"dot1xUsername,omitempty"` // Username for 802.1X authentication. dot1xUsername must have a minimum of 1 character and a maximum of 32 characters.

	Dot1XPassword string `json:"dot1xPassword,omitempty"` // Password for 802.1X authentication. AP dot1x password length should not exceed 120.

	SSHEnabled *bool `json:"sshEnabled,omitempty"` // Indicates if SSH is enabled on the AP. Enable SSH add credentials for device management.

	TelnetEnabled *bool `json:"telnetEnabled,omitempty"` // Indicates if Telnet is enabled on the AP. Enable Telnet to add credentials for device management.

	ManagementUserName string `json:"managementUserName,omitempty"` // Management username must have a minimum of 1 character and a maximum of 32 characters.

	ManagementPassword string `json:"managementPassword,omitempty"` // Management password for the AP. Length must be 8-120 characters.

	ManagementEnablePassword string `json:"managementEnablePassword,omitempty"` // Enable password for managing the AP. Length must be 8-120 characters.

	CdpState *bool `json:"cdpState,omitempty"` // Indicates if CDP is enabled on the AP. Enable CDP in order to make Cisco Access Points known to its neighboring devices and vice-versa.
}
type RequestWirelessCreateApProfileV1RogueDetectionSetting struct {
	RogueDetection *bool `json:"rogueDetection,omitempty"` // Indicates if rogue detection is enabled on the AP. Detect Access Points that have been installed on a secure network without explicit authorization from a system administrator and configure rogue general configuration parameters

	RogueDetectionMinRssi *int `json:"rogueDetectionMinRssi,omitempty"` // Minimum RSSI for rogue detection. Value should be in range -128 decibel milliwatts and -70 decibel milliwatts

	RogueDetectionTransientInterval *int `json:"rogueDetectionTransientInterval,omitempty"` // Transient interval for rogue detection. Value should be 0 or from 120 to 1800.

	RogueDetectionReportInterval *int `json:"rogueDetectionReportInterval,omitempty"` // Report interval for rogue detection. Value should be in range 10 and 300.
}
type RequestWirelessCreateApProfileV1MeshSetting struct {
	BridgeGroupName string `json:"bridgeGroupName,omitempty"` // Name of the bridge group for mesh settings. If not configured, 'Default' Bridge group name will be used in mesh profile.

	BackhaulClientAccess *bool `json:"backhaulClientAccess,omitempty"` // Indicates if backhaul client access is enabled on the AP.

	Range *int `json:"range,omitempty"` // Range of the mesh network. Value should be between 150-132000

	Ghz5BackhaulDataRates string `json:"ghz5BackhaulDataRates,omitempty"` // 5GHz backhaul data rates.

	Ghz24BackhaulDataRates string `json:"ghz24BackhaulDataRates,omitempty"` // 2.4GHz backhaul data rates.

	RapDownlinkBackhaul string `json:"rapDownlinkBackhaul,omitempty"` // Type of downlink backhaul used.
}
type RequestWirelessCreateApProfileV1CalendarPowerProfiles struct {
	PowerProfileName string `json:"powerProfileName,omitempty"` // Name of the existing AP power profile to be mapped to the calendar power profile. API-/intent/api/v1/wirelessSettings/powerProfiles.

	SchedulerType string `json:"schedulerType,omitempty"` // Type of the scheduler.

	Duration *RequestWirelessCreateApProfileV1CalendarPowerProfilesDuration `json:"duration,omitempty"` //
}
type RequestWirelessCreateApProfileV1CalendarPowerProfilesDuration struct {
	SchedulerStartTime string `json:"schedulerStartTime,omitempty"` // Start time of the duration setting.

	SchedulerEndTime string `json:"schedulerEndTime,omitempty"` // End time of the duration setting.

	SchedulerDay string `json:"schedulerDay,omitempty"` // Applies every week on the selected days

	SchedulerDate string `json:"schedulerDate,omitempty"` // Start and End date of the duration setting, applicable for MONTHLY schedulers.
}
type RequestWirelessUpdateApProfileByIDV1 struct {
	ApProfileName string `json:"apProfileName,omitempty"` // Name of the Access Point profile. Max length is 32 characters.

	Description string `json:"description,omitempty"` // Description of the AP profile. Max length is 241 characters

	RemoteWorkerEnabled *bool `json:"remoteWorkerEnabled,omitempty"` // Indicates if remote worker mode is enabled on the AP. Remote teleworker enabled profile cannot support security features like aWIPS,Forensic Capture Enablement, Rogue Detection and Rogue Containment.

	ManagementSetting *RequestWirelessUpdateApProfileByIDV1ManagementSetting `json:"managementSetting,omitempty"` //

	AwipsEnabled *bool `json:"awipsEnabled,omitempty"` // Indicates if AWIPS is enabled on the AP.

	AwipsForensicEnabled *bool `json:"awipsForensicEnabled,omitempty"` // Indicates if AWIPS forensic is enabled on the AP. Forensic Capture is supported from IOS-XE version 17.4 and above. Forensic Capture can be activated only if aWIPS is enabled.

	RogueDetectionSetting *RequestWirelessUpdateApProfileByIDV1RogueDetectionSetting `json:"rogueDetectionSetting,omitempty"` //

	PmfDenialEnabled *bool `json:"pmfDenialEnabled,omitempty"` // Indicates if PMF denial is active on the AP. PMF Denial is supported from IOS-XE version 17.12 and above.

	MeshEnabled *bool `json:"meshEnabled,omitempty"` // This indicates whether mesh networking is enabled on the AP. For IOS-XE devices, when mesh networking is enabled, a custom mesh profile with the configured parameters will be created and mapped to the AP join profile on the device. When mesh networking is disabled, any existing custom mesh profile will be deleted from the device, and the AP join profile will be mapped to the default mesh profile on the device.

	MeshSetting *RequestWirelessUpdateApProfileByIDV1MeshSetting `json:"meshSetting,omitempty"` //

	ApPowerProfileName string `json:"apPowerProfileName,omitempty"` // Name of the existing AP power profile.

	CalendarPowerProfiles *RequestWirelessUpdateApProfileByIDV1CalendarPowerProfiles `json:"calendarPowerProfiles,omitempty"` //

	CountryCode string `json:"countryCode,omitempty"` // Country Code

	TimeZone string `json:"timeZone,omitempty"` // In the Time Zone area, choose one of the following options.             Not Configured - APs operate in the UTC time zone.             Controller - APs operate in the Cisco Wireless Controller time zone.             Delta from Controller - APs operate in the offset time from the wireless controller time zone.

	TimeZoneOffsetHour *int `json:"timeZoneOffsetHour,omitempty"` // Enter the hour value (HH). The valid range is from -12 through 14.

	TimeZoneOffsetMinutes *int `json:"timeZoneOffsetMinutes,omitempty"` // Enter the minute value (MM). The valid range is from 0 through 59.

	ClientLimit *int `json:"clientLimit,omitempty"` // Number of clients. Value should be between 0-1200.
}
type RequestWirelessUpdateApProfileByIDV1ManagementSetting struct {
	AuthType string `json:"authType,omitempty"` // Authentication type used in the AP profile. These setting are applicable during PnP claim and for day-N authentication of AP. Changing these settings will be service impacting for the PnP onboarded APs and will need a factory-reset for those APs.

	Dot1XUsername string `json:"dot1xUsername,omitempty"` // Username for 802.1X authentication. dot1xUsername must have a minimum of 1 character and a maximum of 32 characters.

	Dot1XPassword string `json:"dot1xPassword,omitempty"` // Password for 802.1X authentication. AP dot1x password length should not exceed 120.

	SSHEnabled *bool `json:"sshEnabled,omitempty"` // Indicates if SSH is enabled on the AP. Enable SSH add credentials for device management.

	TelnetEnabled *bool `json:"telnetEnabled,omitempty"` // Indicates if Telnet is enabled on the AP. Enable Telnet to add credentials for device management.

	ManagementUserName string `json:"managementUserName,omitempty"` // Management username must have a minimum of 1 character and a maximum of 32 characters.

	ManagementPassword string `json:"managementPassword,omitempty"` // Management password for the AP. Length must be 8-120 characters.

	ManagementEnablePassword string `json:"managementEnablePassword,omitempty"` // Enable password for managing the AP. Length must be 8-120 characters.

	CdpState *bool `json:"cdpState,omitempty"` // Indicates if CDP is enabled on the AP. Enable CDP in order to make Cisco Access Points known to its neighboring devices and vice-versa.
}
type RequestWirelessUpdateApProfileByIDV1RogueDetectionSetting struct {
	RogueDetection *bool `json:"rogueDetection,omitempty"` // Indicates if rogue detection is enabled on the AP. Detect Access Points that have been installed on a secure network without explicit authorization from a system administrator and configure rogue general configuration parameters

	RogueDetectionMinRssi *int `json:"rogueDetectionMinRssi,omitempty"` // Minimum RSSI for rogue detection. Value should be in range -128 decibel milliwatts and -70 decibel milliwatts

	RogueDetectionTransientInterval *int `json:"rogueDetectionTransientInterval,omitempty"` // Transient interval for rogue detection. Value should be 0 or from 120 to 1800.

	RogueDetectionReportInterval *int `json:"rogueDetectionReportInterval,omitempty"` // Report interval for rogue detection. Value should be in range 10 and 300.
}
type RequestWirelessUpdateApProfileByIDV1MeshSetting struct {
	BridgeGroupName string `json:"bridgeGroupName,omitempty"` // Name of the bridge group for mesh settings. If not configured, 'Default' Bridge group name will be used in mesh profile.

	BackhaulClientAccess *bool `json:"backhaulClientAccess,omitempty"` // Indicates if backhaul client access is enabled on the AP.

	Range *int `json:"range,omitempty"` // Range of the mesh network. Value should be between 150-132000

	Ghz5BackhaulDataRates string `json:"ghz5BackhaulDataRates,omitempty"` // 5GHz backhaul data rates.

	Ghz24BackhaulDataRates string `json:"ghz24BackhaulDataRates,omitempty"` // 2.4GHz backhaul data rates.

	RapDownlinkBackhaul string `json:"rapDownlinkBackhaul,omitempty"` // Type of downlink backhaul used. Available values 5 GHz, 2.4 GHz.
}
type RequestWirelessUpdateApProfileByIDV1CalendarPowerProfiles struct {
	PowerProfileName string `json:"powerProfileName,omitempty"` // Name of the existing AP power profile to be mapped to the calendar power profile. API-/intent/api/v1/wirelessSettings/powerProfiles.

	SchedulerType string `json:"schedulerType,omitempty"` // Type of the scheduler.

	Duration *RequestWirelessUpdateApProfileByIDV1CalendarPowerProfilesDuration `json:"duration,omitempty"` //
}
type RequestWirelessUpdateApProfileByIDV1CalendarPowerProfilesDuration struct {
	SchedulerStartTime string `json:"schedulerStartTime,omitempty"` // Start time of the duration setting.

	SchedulerEndTime string `json:"schedulerEndTime,omitempty"` // End time of the duration setting.

	SchedulerDay string `json:"schedulerDay,omitempty"` // Applies every week on the selected days

	SchedulerDate string `json:"schedulerDate,omitempty"` // Start and End date of the duration setting, applicable for MONTHLY schedulers.
}
type RequestWirelessCreateA80211BeProfileV1 struct {
	ProfileName string `json:"profileName,omitempty"` // 802.11be Profile Name

	OfdmaDownLink *bool `json:"ofdmaDownLink,omitempty"` // OFDMA Downlink (Default: true)

	OfdmaUpLink *bool `json:"ofdmaUpLink,omitempty"` // OFDMA Uplink (Default: true)

	MuMimoDownLink *bool `json:"muMimoDownLink,omitempty"` // MU-MIMO Downlink (Default: false)

	MuMimoUpLink *bool `json:"muMimoUpLink,omitempty"` // MU-MIMO Uplink (Default: false)

	OfdmaMultiRu *bool `json:"ofdmaMultiRu,omitempty"` // OFDMA Multi-RU (Default: false)
}
type RequestWirelessUpdate80211BeProfileV1 struct {
	ProfileName string `json:"profileName,omitempty"` // 802.11be Profile Name

	OfdmaDownLink *bool `json:"ofdmaDownLink,omitempty"` // OFDMA Downlink (Default: true)

	OfdmaUpLink *bool `json:"ofdmaUpLink,omitempty"` // OFDMA Uplink (Default: true)

	MuMimoDownLink *bool `json:"muMimoDownLink,omitempty"` // MU-MIMO Downlink (Default: false)

	MuMimoUpLink *bool `json:"muMimoUpLink,omitempty"` // MU-MIMO Uplink (Default: false)

	OfdmaMultiRu *bool `json:"ofdmaMultiRu,omitempty"` // OFDMA Multi-RU (Default: false)
}
type RequestWirelessCreateInterfaceV1 struct {
	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name

	VLANID *int `json:"vlanId,omitempty"` // VLAN ID range is 1-4094
}
type RequestWirelessUpdateInterfaceV1 struct {
	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name

	VLANID *int `json:"vlanId,omitempty"` // VLAN ID range is 1-4094
}
type RequestWirelessCreatePowerProfileV1 struct {
	ProfileName string `json:"profileName,omitempty"` // Name of the Power Profile. Max allowed characters is 128

	Description string `json:"description,omitempty"` // Description of the Power Profile. Max allowed characters is 128

	Rules *[]RequestWirelessCreatePowerProfileV1Rules `json:"rules,omitempty"` //
}
type RequestWirelessCreatePowerProfileV1Rules struct {
	InterfaceType string `json:"interfaceType,omitempty"` // Interface Type for the rule.

	InterfaceID string `json:"interfaceId,omitempty"` // Interface Id for the rule.

	ParameterType string `json:"parameterType,omitempty"` // Parameter Type for the rule.

	ParameterValue string `json:"parameterValue,omitempty"` // Parameter Value for the rule.
}
type RequestWirelessUpdatePowerProfileByIDV1 struct {
	ProfileName string `json:"profileName,omitempty"` // Name of the Power Profile. Max length is 32 characters

	Description string `json:"description,omitempty"` // Description of the Power Profile. Max length is 32 characters

	Rules *[]RequestWirelessUpdatePowerProfileByIDV1Rules `json:"rules,omitempty"` //
}
type RequestWirelessUpdatePowerProfileByIDV1Rules struct {
	InterfaceType string `json:"interfaceType,omitempty"` // Interface Type

	InterfaceID string `json:"interfaceID,omitempty"` // Interface ID

	ParameterType string `json:"parameterType,omitempty"` // Parameter Type

	ParameterValue string `json:"parameterValue,omitempty"` // Parameter Value
}
type RequestWirelessCreateRfProfileV1 struct {
	RfProfileName string `json:"rfProfileName,omitempty"` // RF Profile Name

	DefaultRfProfile *bool `json:"defaultRfProfile,omitempty"` // True if RF Profile is default, else False. Maximum of only 1 RF Profile can be marked as default at any given time

	EnableRadioTypeA *bool `json:"enableRadioTypeA,omitempty"` // True if 5 GHz radio band is enabled in the RF Profile, else False

	EnableRadioTypeB *bool `json:"enableRadioTypeB,omitempty"` // True if 2.4 GHz radio band is enabled in the RF Profile, else False

	EnableRadioType6GHz *bool `json:"enableRadioType6GHz,omitempty"` // True if 6 GHz radio band is enabled in the RF Profile, else False

	RadioTypeAProperties *RequestWirelessCreateRfProfileV1RadioTypeAProperties `json:"radioTypeAProperties,omitempty"` //

	RadioTypeBProperties *RequestWirelessCreateRfProfileV1RadioTypeBProperties `json:"radioTypeBProperties,omitempty"` //

	RadioType6GHzProperties *RequestWirelessCreateRfProfileV1RadioType6GHzProperties `json:"radioType6GHzProperties,omitempty"` //
}
type RequestWirelessCreateRfProfileV1RadioTypeAProperties struct {
	ParentProfile string `json:"parentProfile,omitempty"` // Parent profile of 5 GHz radio band

	RadioChannels string `json:"radioChannels,omitempty"` // DCA channels of 5 GHz radio band passed in comma separated format without any spaces. Permissible values: 36, 40, 44, 48, 52, 56, 60, 64, 100, 104, 108, 112, 116, 120, 124, 128, 132, 136, 140, 144, 149, 153, 157, 161, 165, 169, 173

	DataRates string `json:"dataRates,omitempty"` // Data rates of 5 GHz radio band passed in comma separated format without any spaces. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54

	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory data rates of 5 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54

	PowerThresholdV1 *int `json:"powerThresholdV1,omitempty"` // Power threshold of 5 GHz radio band

	RxSopThreshold string `json:"rxSopThreshold,omitempty"` // RX-SOP threshold of 5 GHz radio band

	MinPowerLevel *int `json:"minPowerLevel,omitempty"` // Minimum power level of 5 GHz radio band

	MaxPowerLevel *int `json:"maxPowerLevel,omitempty"` // Maximum power level of 5 GHz radio band

	ChannelWidth string `json:"channelWidth,omitempty"` // Channel Width

	PreamblePuncture *bool `json:"preamblePuncture,omitempty"` // Enable or Disable Preamble Puncturing. This Wifi 7 configuration is applicable to wireless IOS devices supporting 17.15 and higher

	ZeroWaitDfsEnable *bool `json:"zeroWaitDfsEnable,omitempty"` // Zero Wait DFS is applicable only for IOS-XE based Wireless Controllers running 17.9.1 and above versions

	CustomRxSopThreshold *int `json:"customRxSopThreshold,omitempty"` // RX-SOP threshold custom configuration of 5 GHz radio band

	MaxRadioClients *int `json:"maxRadioClients,omitempty"` // Client Limit of 5 GHz radio band

	FraProperties *RequestWirelessCreateRfProfileV1RadioTypeAPropertiesFraProperties `json:"fraProperties,omitempty"` //

	CoverageHoleDetectionProperties *RequestWirelessCreateRfProfileV1RadioTypeAPropertiesCoverageHoleDetectionProperties `json:"coverageHoleDetectionProperties,omitempty"` //

	SpatialReuseProperties *RequestWirelessCreateRfProfileV1RadioTypeAPropertiesSpatialReuseProperties `json:"spatialReuseProperties,omitempty"` //
}
type RequestWirelessCreateRfProfileV1RadioTypeAPropertiesFraProperties struct {
	ClientAware *bool `json:"clientAware,omitempty"` // Client Aware of 5 GHz radio band

	ClientSelect *int `json:"clientSelect,omitempty"` // Client Select(%) of 5 GHz radio band

	ClientReset *int `json:"clientReset,omitempty"` // Client Reset(%) of 5 GHz radio band
}
type RequestWirelessCreateRfProfileV1RadioTypeAPropertiesCoverageHoleDetectionProperties struct {
	ChdClientLevel *int `json:"chdClientLevel,omitempty"` // Coverage Hole Detection Client Level

	ChdDataRssiThreshold *int `json:"chdDataRssiThreshold,omitempty"` // Coverage Hole Detection Data Rssi Threshold

	ChdVoiceRssiThreshold *int `json:"chdVoiceRssiThreshold,omitempty"` // Coverage Hole Detection Voice Rssi Threshold

	ChdExceptionLevel *int `json:"chdExceptionLevel,omitempty"` // Coverage Hole Detection Exception Level(%)
}
type RequestWirelessCreateRfProfileV1RadioTypeAPropertiesSpatialReuseProperties struct {
	Dot11AxNonSrgObssPacketDetect *bool `json:"dot11axNonSrgObssPacketDetect,omitempty"` // Dot11ax Non SRG OBSS PD

	Dot11AxNonSrgObssPacketDetectMaxThreshold *int `json:"dot11axNonSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax Non SRG OBSS PD Max Threshold

	Dot11AxSrgObssPacketDetect *bool `json:"dot11axSrgObssPacketDetect,omitempty"` // Dot11ax SRG OBSS PD

	Dot11AxSrgObssPacketDetectMinThreshold *int `json:"dot11axSrgObssPacketDetectMinThreshold,omitempty"` // Dot11ax SRG OBSS PD Min Threshold

	Dot11AxSrgObssPacketDetectMaxThreshold *int `json:"dot11axSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax SRG OBSS PD Max Threshold
}
type RequestWirelessCreateRfProfileV1RadioTypeBProperties struct {
	ParentProfile string `json:"parentProfile,omitempty"` // Parent profile of 2.4 GHz radio band

	RadioChannels string `json:"radioChannels,omitempty"` // DCA channels of 2.4 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14

	DataRates string `json:"dataRates,omitempty"` // Data rates of 2.4 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 2, 5.5, 6, 9, 11, 12, 18, 24, 36, 48, 54

	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory data rates of 2.4 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 1, 2, 5.5, 6, 9, 11, 12, 18, 24, 36, 48, 54

	PowerThresholdV1 *int `json:"powerThresholdV1,omitempty"` // Power threshold of 2.4 GHz radio band

	RxSopThreshold string `json:"rxSopThreshold,omitempty"` // RX-SOP threshold of 2.4 GHz radio band

	MinPowerLevel *int `json:"minPowerLevel,omitempty"` // Minimum power level of 2.4 GHz radio band

	MaxPowerLevel *int `json:"maxPowerLevel,omitempty"` // Maximum power level of 2.4 GHz radio band

	CustomRxSopThreshold *int `json:"customRxSopThreshold,omitempty"` // RX-SOP threshold custom configuration of 2.4 GHz radio band

	MaxRadioClients *int `json:"maxRadioClients,omitempty"` // Client Limit of 2.4 GHz radio band

	CoverageHoleDetectionProperties *RequestWirelessCreateRfProfileV1RadioTypeBPropertiesCoverageHoleDetectionProperties `json:"coverageHoleDetectionProperties,omitempty"` //

	SpatialReuseProperties *RequestWirelessCreateRfProfileV1RadioTypeBPropertiesSpatialReuseProperties `json:"spatialReuseProperties,omitempty"` //
}
type RequestWirelessCreateRfProfileV1RadioTypeBPropertiesCoverageHoleDetectionProperties struct {
	ChdClientLevel *int `json:"chdClientLevel,omitempty"` // Coverage Hole Detection Client Level

	ChdDataRssiThreshold *int `json:"chdDataRssiThreshold,omitempty"` // Coverage Hole Detection Data Rssi Threshold

	ChdVoiceRssiThreshold *int `json:"chdVoiceRssiThreshold,omitempty"` // Coverage Hole Detection Voice Rssi Threshold

	ChdExceptionLevel *int `json:"chdExceptionLevel,omitempty"` // Coverage Hole Detection Exception Level(%)
}
type RequestWirelessCreateRfProfileV1RadioTypeBPropertiesSpatialReuseProperties struct {
	Dot11AxNonSrgObssPacketDetect *bool `json:"dot11axNonSrgObssPacketDetect,omitempty"` // Dot11ax Non SRG OBSS PD

	Dot11AxNonSrgObssPacketDetectMaxThreshold *int `json:"dot11axNonSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax Non SRG OBSS PD Max Threshold

	Dot11AxSrgObssPacketDetect *bool `json:"dot11axSrgObssPacketDetect,omitempty"` // Dot11ax SRG OBSS PD

	Dot11AxSrgObssPacketDetectMinThreshold *int `json:"dot11axSrgObssPacketDetectMinThreshold,omitempty"` // Dot11ax SRG OBSS PD Min Threshold

	Dot11AxSrgObssPacketDetectMaxThreshold *int `json:"dot11axSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax SRG OBSS PD Max Threshold
}
type RequestWirelessCreateRfProfileV1RadioType6GHzProperties struct {
	ParentProfile string `json:"parentProfile,omitempty"` // Parent profile of 6 GHz radio band

	RadioChannels string `json:"radioChannels,omitempty"` // DCA channels of 6 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 5, 9, 13, 17, 21, 25, 29, 33, 37, 41, 45, 49, 53, 57, 61, 65, 69, 73, 77, 81, 85, 89, 93, 97, 101, 105, 109, 113, 117, 121, 125, 129, 133, 137, 141, 145, 149, 153, 157, 161, 165, 169, 173, 177, 181, 185, 189, 193, 197, 201, 205, 209, 213, 217, 221, 225, 229, 233

	DataRates string `json:"dataRates,omitempty"` // Data rates of 6 GHz radio band passed in comma separated format without any spaces. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54

	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory data rates of 6 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54

	PowerThresholdV1 *int `json:"powerThresholdV1,omitempty"` // Power threshold of 6 GHz radio band

	RxSopThreshold string `json:"rxSopThreshold,omitempty"` // RX-SOP threshold of 6 GHz radio band

	MinPowerLevel *int `json:"minPowerLevel,omitempty"` // Minimum power level of 6 GHz radio band

	MaxPowerLevel *int `json:"maxPowerLevel,omitempty"` // Maximum power level of 6 GHz radio band

	EnableStandardPowerService *bool `json:"enableStandardPowerService,omitempty"` // True if Standard Power Service is enabled, else False

	MultiBssidProperties *RequestWirelessCreateRfProfileV1RadioType6GHzPropertiesMultiBssidProperties `json:"multiBssidProperties,omitempty"` //

	PreamblePuncture *bool `json:"preamblePuncture,omitempty"` // Enable or Disable Preamble Puncturing. This Wifi 7 configuration is applicable to wireless IOS devices supporting 17.15 and higher

	MinDbsWidth *int `json:"minDbsWidth,omitempty"` // Minimum DBS Width (Permissible Values:20,40,80,160,320)

	MaxDbsWidth *int `json:"maxDbsWidth,omitempty"` // Maximum DBS Width (Permissible Values:20,40,80,160,320)

	CustomRxSopThreshold *int `json:"customRxSopThreshold,omitempty"` // RX-SOP threshold custom configuration of 6 GHz radio band

	MaxRadioClients *int `json:"maxRadioClients,omitempty"` // Client Limit of 6 GHz radio band

	PscEnforcingEnabled *bool `json:"pscEnforcingEnabled,omitempty"` // PSC Enforcing Enable for 6 GHz radio band

	DiscoveryFrames6GHz string `json:"discoveryFrames6GHz,omitempty"` // Discovery Frames of 6 GHz radio band

	BroadcastProbeResponseInterval *int `json:"broadcastProbeResponseInterval,omitempty"` // Broadcast Probe Response Interval of 6 GHz radio band

	FraProperties *RequestWirelessCreateRfProfileV1RadioType6GHzPropertiesFraProperties `json:"fraProperties,omitempty"` //

	CoverageHoleDetectionProperties *RequestWirelessCreateRfProfileV1RadioType6GHzPropertiesCoverageHoleDetectionProperties `json:"coverageHoleDetectionProperties,omitempty"` //

	SpatialReuseProperties *RequestWirelessCreateRfProfileV1RadioType6GHzPropertiesSpatialReuseProperties `json:"spatialReuseProperties,omitempty"` //
}
type RequestWirelessCreateRfProfileV1RadioType6GHzPropertiesMultiBssidProperties struct {
	Dot11AxParameters *RequestWirelessCreateRfProfileV1RadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters `json:"dot11axParameters,omitempty"` //

	Dot11BeParameters *RequestWirelessCreateRfProfileV1RadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters `json:"dot11beParameters,omitempty"` //

	TargetWakeTime *bool `json:"targetWakeTime,omitempty"` // Target Wake Time

	TwtBroadcastSupport *bool `json:"twtBroadcastSupport,omitempty"` // TWT Broadcast Support
}
type RequestWirelessCreateRfProfileV1RadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters struct {
	OfdmaDownLink *bool `json:"ofdmaDownLink,omitempty"` // OFDMA Downlink

	OfdmaUpLink *bool `json:"ofdmaUpLink,omitempty"` // OFDMA Uplink

	MuMimoUpLink *bool `json:"muMimoUpLink,omitempty"` // MU-MIMO Uplink

	MuMimoDownLink *bool `json:"muMimoDownLink,omitempty"` // MU-MIMO Downlink
}
type RequestWirelessCreateRfProfileV1RadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters struct {
	OfdmaDownLink *bool `json:"ofdmaDownLink,omitempty"` // OFDMA Downlink

	OfdmaUpLink *bool `json:"ofdmaUpLink,omitempty"` // OFDMA Uplink

	MuMimoUpLink *bool `json:"muMimoUpLink,omitempty"` // MU-MIMO Uplink

	MuMimoDownLink *bool `json:"muMimoDownLink,omitempty"` // MU-MIMO Downlink

	OfdmaMultiRu *bool `json:"ofdmaMultiRu,omitempty"` // OFDMA Multi-RU
}
type RequestWirelessCreateRfProfileV1RadioType6GHzPropertiesFraProperties struct {
	ClientResetCount *int `json:"clientResetCount,omitempty"` // Client Reset Count of 6 GHz radio band

	ClientUtilizationThreshold *int `json:"clientUtilizationThreshold,omitempty"` // Client Utilization Threshold of 6 GHz radio band
}
type RequestWirelessCreateRfProfileV1RadioType6GHzPropertiesCoverageHoleDetectionProperties struct {
	ChdClientLevel *int `json:"chdClientLevel,omitempty"` // Coverage Hole Detection Client Level

	ChdDataRssiThreshold *int `json:"chdDataRssiThreshold,omitempty"` // Coverage Hole Detection Data Rssi Threshold

	ChdVoiceRssiThreshold *int `json:"chdVoiceRssiThreshold,omitempty"` // Coverage Hole Detection Voice Rssi Threshold

	ChdExceptionLevel *int `json:"chdExceptionLevel,omitempty"` // Coverage Hole Detection Exception Level(%)
}
type RequestWirelessCreateRfProfileV1RadioType6GHzPropertiesSpatialReuseProperties struct {
	Dot11AxNonSrgObssPacketDetect *bool `json:"dot11axNonSrgObssPacketDetect,omitempty"` // Dot11ax Non SRG OBSS PD

	Dot11AxNonSrgObssPacketDetectMaxThreshold *int `json:"dot11axNonSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax Non SRG OBSS PD Max Threshold

	Dot11AxSrgObssPacketDetect *bool `json:"dot11axSrgObssPacketDetect,omitempty"` // Dot11ax SRG OBSS PD

	Dot11AxSrgObssPacketDetectMinThreshold *int `json:"dot11axSrgObssPacketDetectMinThreshold,omitempty"` // Dot11ax SRG OBSS PD Min Threshold

	Dot11AxSrgObssPacketDetectMaxThreshold *int `json:"dot11axSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax SRG OBSS PD Max Threshold
}
type RequestWirelessUpdateRfProfileV1 struct {
	RfProfileName string `json:"rfProfileName,omitempty"` // RF Profile Name

	DefaultRfProfile *bool `json:"defaultRfProfile,omitempty"` // True if RF Profile is default, else False. Maximum of only 1 RF Profile can be marked as default at any given time

	EnableRadioTypeA *bool `json:"enableRadioTypeA,omitempty"` // True if 5 GHz radio band is enabled in the RF Profile, else False

	EnableRadioTypeB *bool `json:"enableRadioTypeB,omitempty"` // True if 2.4 GHz radio band is enabled in the RF Profile, else False

	EnableRadioType6GHz *bool `json:"enableRadioType6GHz,omitempty"` // True if 6 GHz radio band is enabled in the RF Profile, else False

	RadioTypeAProperties *RequestWirelessUpdateRfProfileV1RadioTypeAProperties `json:"radioTypeAProperties,omitempty"` //

	RadioTypeBProperties *RequestWirelessUpdateRfProfileV1RadioTypeBProperties `json:"radioTypeBProperties,omitempty"` //

	RadioType6GHzProperties *RequestWirelessUpdateRfProfileV1RadioType6GHzProperties `json:"radioType6GHzProperties,omitempty"` //
}
type RequestWirelessUpdateRfProfileV1RadioTypeAProperties struct {
	ParentProfile string `json:"parentProfile,omitempty"` // Parent profile of 5 GHz radio band. In case of brownfield learnt RF Profile if the parent profile is GLOBAL, any change in RF Profile configurations will not be provisioned to device. Existing parent profile with values of HIGH, TYPICAL, LOW or CUSTOM cannot be modified to GLOBAL

	RadioChannels string `json:"radioChannels,omitempty"` // DCA channels of 5 GHz radio band passed in comma separated format without any spaces. Permissible values: 36, 40, 44, 48, 52, 56, 60, 64, 100, 104, 108, 112, 116, 120, 124, 128, 132, 136, 140, 144, 149, 153, 157, 161, 165, 169, 173

	DataRates string `json:"dataRates,omitempty"` // Data rates of 5 GHz radio band passed in comma separated format without any spaces. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54

	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory data rates of 5 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54

	PowerThresholdV1 *int `json:"powerThresholdV1,omitempty"` // Power threshold of 5 GHz radio band

	RxSopThreshold string `json:"rxSopThreshold,omitempty"` // RX-SOP threshold of 5 GHz radio band

	MinPowerLevel *int `json:"minPowerLevel,omitempty"` // Minimum power level of 5 GHz radio band

	MaxPowerLevel *int `json:"maxPowerLevel,omitempty"` // Maximum power level of 5 GHz radio band

	ChannelWidth string `json:"channelWidth,omitempty"` // Channel Width

	PreamblePuncture *bool `json:"preamblePuncture,omitempty"` // Enable or Disable Preamble Puncturing. This Wifi 7 configuration is applicable to wireless IOS devices supporting 17.15 and higher

	ZeroWaitDfsEnable *bool `json:"zeroWaitDfsEnable,omitempty"` // Zero Wait DFS is applicable only for IOS-XE based Wireless Controllers running 17.9.1 and above versions

	CustomRxSopThreshold *int `json:"customRxSopThreshold,omitempty"` // RX-SOP threshold custom configuration of 5 GHz radio band

	MaxRadioClients *int `json:"maxRadioClients,omitempty"` // Client Limit of 5 GHz radio band

	FraProperties *RequestWirelessUpdateRfProfileV1RadioTypeAPropertiesFraProperties `json:"fraProperties,omitempty"` //

	CoverageHoleDetectionProperties *RequestWirelessUpdateRfProfileV1RadioTypeAPropertiesCoverageHoleDetectionProperties `json:"coverageHoleDetectionProperties,omitempty"` //

	SpatialReuseProperties *RequestWirelessUpdateRfProfileV1RadioTypeAPropertiesSpatialReuseProperties `json:"spatialReuseProperties,omitempty"` //
}
type RequestWirelessUpdateRfProfileV1RadioTypeAPropertiesFraProperties struct {
	ClientAware *bool `json:"clientAware,omitempty"` // Client Aware of 5 GHz radio band

	ClientSelect *int `json:"clientSelect,omitempty"` // Client Select(%) of 5 GHz radio band

	ClientReset *int `json:"clientReset,omitempty"` // Client Reset(%) of 5 GHz radio band
}
type RequestWirelessUpdateRfProfileV1RadioTypeAPropertiesCoverageHoleDetectionProperties struct {
	ChdClientLevel *int `json:"chdClientLevel,omitempty"` // Coverage Hole Detection Client Level

	ChdDataRssiThreshold *int `json:"chdDataRssiThreshold,omitempty"` // Coverage Hole Detection Data Rssi Threshold

	ChdVoiceRssiThreshold *int `json:"chdVoiceRssiThreshold,omitempty"` // Coverage Hole Detection Voice Rssi Threshold

	ChdExceptionLevel *int `json:"chdExceptionLevel,omitempty"` // Coverage Hole Detection Exception Level(%)
}
type RequestWirelessUpdateRfProfileV1RadioTypeAPropertiesSpatialReuseProperties struct {
	Dot11AxNonSrgObssPacketDetect *bool `json:"dot11axNonSrgObssPacketDetect,omitempty"` // Dot11ax Non SRG OBSS PD

	Dot11AxNonSrgObssPacketDetectMaxThreshold *int `json:"dot11axNonSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax Non SRG OBSS PD Max Threshold

	Dot11AxSrgObssPacketDetect *bool `json:"dot11axSrgObssPacketDetect,omitempty"` // Dot11ax SRG OBSS PD

	Dot11AxSrgObssPacketDetectMinThreshold *int `json:"dot11axSrgObssPacketDetectMinThreshold,omitempty"` // Dot11ax SRG OBSS PD Min Threshold

	Dot11AxSrgObssPacketDetectMaxThreshold *int `json:"dot11axSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax SRG OBSS PD Max Threshold
}
type RequestWirelessUpdateRfProfileV1RadioTypeBProperties struct {
	ParentProfile string `json:"parentProfile,omitempty"` // Parent profile of 2.4 GHz radio band. In case of brownfield learnt RF Profile if the parent profile is GLOBAL, any change in RF Profile configurations will not be provisioned to device. Existing parent profile with values of HIGH, TYPICAL, LOW or CUSTOM cannot be modified to GLOBAL

	RadioChannels string `json:"radioChannels,omitempty"` // DCA channels of 2.4 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14

	DataRates string `json:"dataRates,omitempty"` // Data rates of 2.4 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 2, 5.5, 6, 9, 11, 12, 18, 24, 36, 48, 54

	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory data rates of 2.4 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 1, 2, 5.5, 6, 9, 11, 12, 18, 24, 36, 48, 54

	PowerThresholdV1 *int `json:"powerThresholdV1,omitempty"` // Power threshold of 2.4 GHz radio band

	RxSopThreshold string `json:"rxSopThreshold,omitempty"` // RX-SOP threshold of 2.4 GHz radio band

	MinPowerLevel *int `json:"minPowerLevel,omitempty"` // Minimum power level of 2.4 GHz radio band

	MaxPowerLevel *int `json:"maxPowerLevel,omitempty"` // Maximum power level of 2.4 GHz radio band

	CustomRxSopThreshold *int `json:"customRxSopThreshold,omitempty"` // RX-SOP threshold custom configuration of 2.4 GHz radio band

	MaxRadioClients *int `json:"maxRadioClients,omitempty"` // Client Limit of 2.4 GHz radio band

	CoverageHoleDetectionProperties *RequestWirelessUpdateRfProfileV1RadioTypeBPropertiesCoverageHoleDetectionProperties `json:"coverageHoleDetectionProperties,omitempty"` //

	SpatialReuseProperties *RequestWirelessUpdateRfProfileV1RadioTypeBPropertiesSpatialReuseProperties `json:"spatialReuseProperties,omitempty"` //
}
type RequestWirelessUpdateRfProfileV1RadioTypeBPropertiesCoverageHoleDetectionProperties struct {
	ChdClientLevel *int `json:"chdClientLevel,omitempty"` // Coverage Hole Detection Client Level

	ChdDataRssiThreshold *int `json:"chdDataRssiThreshold,omitempty"` // Coverage Hole Detection Data Rssi Threshold

	ChdVoiceRssiThreshold *int `json:"chdVoiceRssiThreshold,omitempty"` // Coverage Hole Detection Voice Rssi Threshold

	ChdExceptionLevel *int `json:"chdExceptionLevel,omitempty"` // Coverage Hole Detection Exception Level(%)
}
type RequestWirelessUpdateRfProfileV1RadioTypeBPropertiesSpatialReuseProperties struct {
	Dot11AxNonSrgObssPacketDetect *bool `json:"dot11axNonSrgObssPacketDetect,omitempty"` // Dot11ax Non SRG OBSS PD

	Dot11AxNonSrgObssPacketDetectMaxThreshold *int `json:"dot11axNonSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax Non SRG OBSS PD Max Threshold

	Dot11AxSrgObssPacketDetect *bool `json:"dot11axSrgObssPacketDetect,omitempty"` // Dot11ax SRG OBSS PD

	Dot11AxSrgObssPacketDetectMinThreshold *int `json:"dot11axSrgObssPacketDetectMinThreshold,omitempty"` // Dot11ax SRG OBSS PD Min Threshold

	Dot11AxSrgObssPacketDetectMaxThreshold *int `json:"dot11axSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax SRG OBSS PD Max Threshold
}
type RequestWirelessUpdateRfProfileV1RadioType6GHzProperties struct {
	ParentProfile string `json:"parentProfile,omitempty"` // Parent profile of 6 GHz radio band

	RadioChannels string `json:"radioChannels,omitempty"` // DCA channels of 6 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 5, 9, 13, 17, 21, 25, 29, 33, 37, 41, 45, 49, 53, 57, 61, 65, 69, 73, 77, 81, 85, 89, 93, 97, 101, 105, 109, 113, 117, 121, 125, 129, 133, 137, 141, 145, 149, 153, 157, 161, 165, 169, 173, 177, 181, 185, 189, 193, 197, 201, 205, 209, 213, 217, 221, 225, 229, 233

	DataRates string `json:"dataRates,omitempty"` // Data rates of 6 GHz radio band passed in comma separated format without any spaces. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54

	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory data rates of 6 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54

	PowerThresholdV1 *int `json:"powerThresholdV1,omitempty"` // Power threshold of 6 GHz radio band

	RxSopThreshold string `json:"rxSopThreshold,omitempty"` // RX-SOP threshold of 6 GHz radio band

	MinPowerLevel *int `json:"minPowerLevel,omitempty"` // Minimum power level of 6 GHz radio band

	MaxPowerLevel *int `json:"maxPowerLevel,omitempty"` // Maximum power level of 6 GHz radio band

	EnableStandardPowerService *bool `json:"enableStandardPowerService,omitempty"` // True if Standard Power Service is enabled, else False

	MultiBssidProperties *RequestWirelessUpdateRfProfileV1RadioType6GHzPropertiesMultiBssidProperties `json:"multiBssidProperties,omitempty"` //

	PreamblePuncture *bool `json:"preamblePuncture,omitempty"` // Enable or Disable Preamble Puncturing. This Wifi 7 configuration is applicable to wireless IOS devices supporting 17.15 and higher

	MinDbsWidth *int `json:"minDbsWidth,omitempty"` // Minimum DBS Width (Permissible Values:20,40,80,160,320)

	MaxDbsWidth *int `json:"maxDbsWidth,omitempty"` // Maximum DBS Width (Permissible Values:20,40,80,160,320)

	CustomRxSopThreshold *int `json:"customRxSopThreshold,omitempty"` // RX-SOP threshold custom configuration of 6 GHz radio band

	MaxRadioClients *int `json:"maxRadioClients,omitempty"` // Client Limit of 6 GHz radio band

	PscEnforcingEnabled *bool `json:"pscEnforcingEnabled,omitempty"` // PSC Enforcing Enable for 6 GHz radio band

	DiscoveryFrames6GHz string `json:"discoveryFrames6GHz,omitempty"` // Discovery Frames of 6 GHz radio band

	BroadcastProbeResponseInterval *int `json:"broadcastProbeResponseInterval,omitempty"` // Broadcast Probe Response Interval of 6 GHz radio band

	FraProperties *RequestWirelessUpdateRfProfileV1RadioType6GHzPropertiesFraProperties `json:"fraProperties,omitempty"` //

	CoverageHoleDetectionProperties *RequestWirelessUpdateRfProfileV1RadioType6GHzPropertiesCoverageHoleDetectionProperties `json:"coverageHoleDetectionProperties,omitempty"` //

	SpatialReuseProperties *RequestWirelessUpdateRfProfileV1RadioType6GHzPropertiesSpatialReuseProperties `json:"spatialReuseProperties,omitempty"` //
}
type RequestWirelessUpdateRfProfileV1RadioType6GHzPropertiesMultiBssidProperties struct {
	Dot11AxParameters *RequestWirelessUpdateRfProfileV1RadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters `json:"dot11axParameters,omitempty"` //

	Dot11BeParameters *RequestWirelessUpdateRfProfileV1RadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters `json:"dot11beParameters,omitempty"` //

	TargetWakeTime *bool `json:"targetWakeTime,omitempty"` // Target Wake Time

	TwtBroadcastSupport *bool `json:"twtBroadcastSupport,omitempty"` // TWT Broadcast Support
}
type RequestWirelessUpdateRfProfileV1RadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters struct {
	OfdmaDownLink *bool `json:"ofdmaDownLink,omitempty"` // OFDMA Downlink

	OfdmaUpLink *bool `json:"ofdmaUpLink,omitempty"` // OFDMA Uplink

	MuMimoUpLink *bool `json:"muMimoUpLink,omitempty"` // MU-MIMO Uplink

	MuMimoDownLink *bool `json:"muMimoDownLink,omitempty"` // MU-MIMO Downlink
}
type RequestWirelessUpdateRfProfileV1RadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters struct {
	OfdmaDownLink *bool `json:"ofdmaDownLink,omitempty"` // OFDMA Downlink

	OfdmaUpLink *bool `json:"ofdmaUpLink,omitempty"` // OFDMA Uplink

	MuMimoUpLink *bool `json:"muMimoUpLink,omitempty"` // MU-MIMO Uplink

	MuMimoDownLink *bool `json:"muMimoDownLink,omitempty"` // MU-MIMO Downlink

	OfdmaMultiRu *bool `json:"ofdmaMultiRu,omitempty"` // OFDMA Multi-RU
}
type RequestWirelessUpdateRfProfileV1RadioType6GHzPropertiesFraProperties struct {
	ClientResetCount *int `json:"clientResetCount,omitempty"` // Client Reset Count of 6 GHz radio band

	ClientUtilizationThreshold *int `json:"clientUtilizationThreshold,omitempty"` // Client Utilization Threshold of 6 GHz radio band
}
type RequestWirelessUpdateRfProfileV1RadioType6GHzPropertiesCoverageHoleDetectionProperties struct {
	ChdClientLevel *int `json:"chdClientLevel,omitempty"` // Coverage Hole Detection Client Level

	ChdDataRssiThreshold *int `json:"chdDataRssiThreshold,omitempty"` // Coverage Hole Detection Data Rssi Threshold

	ChdVoiceRssiThreshold *int `json:"chdVoiceRssiThreshold,omitempty"` // Coverage Hole Detection Voice Rssi Threshold

	ChdExceptionLevel *int `json:"chdExceptionLevel,omitempty"` // Coverage Hole Detection Exception Level(%)
}
type RequestWirelessUpdateRfProfileV1RadioType6GHzPropertiesSpatialReuseProperties struct {
	Dot11AxNonSrgObssPacketDetect *bool `json:"dot11axNonSrgObssPacketDetect,omitempty"` // Dot11ax Non SRG OBSS PD

	Dot11AxNonSrgObssPacketDetectMaxThreshold *int `json:"dot11axNonSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax Non SRG OBSS PD Max Threshold

	Dot11AxSrgObssPacketDetect *bool `json:"dot11axSrgObssPacketDetect,omitempty"` // Dot11ax SRG OBSS PD

	Dot11AxSrgObssPacketDetectMinThreshold *int `json:"dot11axSrgObssPacketDetectMinThreshold,omitempty"` // Dot11ax SRG OBSS PD Min Threshold

	Dot11AxSrgObssPacketDetectMaxThreshold *int `json:"dot11axSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax SRG OBSS PD Max Threshold
}
type RequestWirelessAssignAnchorManagedApLocationsForWLCV1 struct {
	AnchorManagedApLocationsSiteIDs []string `json:"anchorManagedAPLocationsSiteIds,omitempty"` // This API allows user to assign Anchor Managed AP Locations for WLC by device ID. The payload should always be a complete list. The Managed AP Locations included in the payload will be fully processed for both addition and deletion.               -  When anchor managed location array present then it will add the anchor managed locations.
}
type RequestWirelessConfigureAccessPointsV2 struct {
	ApList *[]RequestWirelessConfigureAccessPointsV2ApList `json:"apList,omitempty"` //

	ConfigureAdminStatus *bool `json:"configureAdminStatus,omitempty"` // To change the access point's admin status, set this parameter's value to "true".

	AdminStatus *bool `json:"adminStatus,omitempty"` // Configure the access point's admin status. Set this parameter's value to "true" to enable it and "false" to disable it.

	ConfigureApMode *bool `json:"configureApMode,omitempty"` // To change the access point's mode, set this parameter's value to "true".

	ApMode *int `json:"apMode,omitempty"` // Configure the access point's mode: for local/flexconnect mode, set "0"; for monitor mode, set "1"; for sniffer mode, set "4"; and for bridge/flex+bridge mode, set "5".

	ConfigureFailoverPriority *bool `json:"configureFailoverPriority,omitempty"` // To change the access point's failover priority, set this parameter's value to "true".

	FailoverPriority *int `json:"failoverPriority,omitempty"` // Configure the acess point's failover priority: for low, set "1"; for medium, set "2"; for high, set "3"; and for critical, set "4".

	ConfigureLedStatus *bool `json:"configureLedStatus,omitempty"` // To change the access point's LED status, set this parameter's value to "true".

	LedStatus *bool `json:"ledStatus,omitempty"` // Configure the access point's LED status. Set "true" to enable its status and "false" to disable it.

	ConfigureLedBrightnessLevel *bool `json:"configureLedBrightnessLevel,omitempty"` // To change the access point's LED brightness level, set this parameter's value to "true".

	LedBrightnessLevel *int `json:"ledBrightnessLevel,omitempty"` // Configure the access point's LED brightness level by setting a value between 1 and 8.

	ConfigureLocation *bool `json:"configureLocation,omitempty"` // To change the access point's location, set this parameter's value to "true".

	Location string `json:"location,omitempty"` // Configure the access point's location.

	ConfigureHAController *bool `json:"configureHAController,omitempty"` // To change the access point's HA controller, set this parameter's value to "true".

	PrimaryControllerName string `json:"primaryControllerName,omitempty"` // Configure the hostname for an access point's primary controller.

	PrimaryIPAddress *RequestWirelessConfigureAccessPointsV2PrimaryIPAddress `json:"primaryIpAddress,omitempty"` //

	SecondaryControllerName string `json:"secondaryControllerName,omitempty"` // Configure the hostname for an access point's secondary controller.

	SecondaryIPAddress *RequestWirelessConfigureAccessPointsV2SecondaryIPAddress `json:"secondaryIpAddress,omitempty"` //

	TertiaryControllerName string `json:"tertiaryControllerName,omitempty"` // Configure the hostname for an access point's tertiary controller.

	TertiaryIPAddress *RequestWirelessConfigureAccessPointsV2TertiaryIPAddress `json:"tertiaryIpAddress,omitempty"` //

	RadioConfigurations *[]RequestWirelessConfigureAccessPointsV2RadioConfigurations `json:"radioConfigurations,omitempty"` //

	ConfigureCleanAirSI24Ghz *bool `json:"configureCleanAirSI24Ghz,omitempty"` // To change the clean air status for radios that are in 2.4 Ghz band, set this parameter's value to "true".

	CleanAirSI24 *bool `json:"cleanAirSI24,omitempty"` // Configure clean air status for radios that are in 2.4 Ghz band. Set this parameter's value to "true" to enable it and "false" to disable it.

	ConfigureCleanAirSI5Ghz *bool `json:"configureCleanAirSI5Ghz,omitempty"` // To change the clean air status for radios that are in 5 Ghz band, set this parameter's value to "true".

	CleanAirSI5 *bool `json:"cleanAirSI5,omitempty"` // Configure clean air status for radios that are in 5 Ghz band. Set this parameter's value to "true" to enable it and "false" to disable it.

	ConfigureCleanAirSI6Ghz *bool `json:"configureCleanAirSI6Ghz,omitempty"` // To change the clean air status for radios that are in 6 Ghz band, set this parameter's value to "true".

	CleanAirSI6 *bool `json:"cleanAirSI6,omitempty"` // Configure clean air status for radios that are in 6 Ghz band. Set this parameter's value to "true" to enable it and "false" to disable it.

	IsAssignedSiteAsLocation *bool `json:"isAssignedSiteAsLocation,omitempty"` // To configure the access point's location as the site assigned to the access point, set this parameter's value to "true".
}
type RequestWirelessConfigureAccessPointsV2ApList struct {
	ApName string `json:"apName,omitempty"` // The current host name of the access point.

	MacAddress string `json:"macAddress,omitempty"` // The ethernet MAC address of the access point.

	ApNameNew string `json:"apNameNew,omitempty"` // The modified hostname of the access point.
}
type RequestWirelessConfigureAccessPointsV2PrimaryIPAddress struct {
	Address string `json:"address,omitempty"` // Configure the IP address for an access point's primary controller.
}
type RequestWirelessConfigureAccessPointsV2SecondaryIPAddress struct {
	Address string `json:"address,omitempty"` // Configure the IP address for an access point's secondary controller.
}
type RequestWirelessConfigureAccessPointsV2TertiaryIPAddress struct {
	Address string `json:"address,omitempty"` // Configure the IP address for an access point's tertiary controller.
}
type RequestWirelessConfigureAccessPointsV2RadioConfigurations struct {
	ConfigureRadioRoleAssignment *bool `json:"configureRadioRoleAssignment,omitempty"` // To change the radio role on the specified radio for an access point, set this parameter's value to "true".

	RadioRoleAssignment string `json:"radioRoleAssignment,omitempty"` // Configure only one of the following roles on the specified radio for an access point as "AUTO", "SERVING", or "MONITOR". Any other string is invalid, including empty string

	RadioBand string `json:"radioBand,omitempty"` // Configure the band on the specified radio for an access point: for 2.4 GHz, set "RADIO24"; for 5 GHz, set "RADIO5". Any other string is invalid, including empty string

	ConfigureAdminStatus *bool `json:"configureAdminStatus,omitempty"` // To change the admin status on the specified radio for an access point, set this parameter's value to "true".

	AdminStatus *bool `json:"adminStatus,omitempty"` // Configure the admin status on the specified radio for an access point. Set this parameter's value to "true" to enable it and "false" to disable it.

	ConfigureAntennaPatternName *bool `json:"configureAntennaPatternName,omitempty"` // To change the antenna gain on the specified radio for an access point, set the value for this parameter to "true".

	AntennaPatternName string `json:"antennaPatternName,omitempty"` // Specify the antenna name on the specified radio for an access point. The antenna name is used to calculate the gain on the radio slot.

	AntennaGain *int `json:"antennaGain,omitempty"` // Configure the antenna gain on the specified radio for an access point by setting a decimal value (in dBi). To configure "antennaGain", set "antennaPatternName" value to "other".

	ConfigureAntennaCable *bool `json:"configureAntennaCable,omitempty"` // To change the antenna cable name on the specified radio for an access point, set this parameter's value to "true".

	AntennaCableName string `json:"antennaCableName,omitempty"` // Configure the antenna cable name on the specified radio for an access point. If cable loss needs to be configured, set this parameter's value to "other".

	CableLoss *float64 `json:"cableLoss,omitempty"` // Configure the cable loss on the specified radio for an access point by setting a decimal value (in dBi).

	ConfigureChannel *bool `json:"configureChannel,omitempty"` // To change the channel on the specified radio for an access point, set this parameter's value to "true".

	ChannelAssignmentMode *int `json:"channelAssignmentMode,omitempty"` // Configure the channel assignment mode on the specified radio for an access point: for global mode, set "1"; and for custom mode, set "2".

	ChannelNumber *int `json:"channelNumber,omitempty"` // Configure the channel number on the specified radio for an access point.

	ConfigureChannelWidth *bool `json:"configureChannelWidth,omitempty"` // To change the channel width on the specified radio for an access point, set this parameter's value to "true".

	ChannelWidth *int `json:"channelWidth,omitempty"` // Configure the channel width on the specified radio for an access point: for 20 MHz, set "3"; for 40 MHz, set "4"; for 80 MHz, set "5"; for 160 MHz, set "6", and for 320 MHz, set "7".

	ConfigurePower *bool `json:"configurePower,omitempty"` // To change the power assignment mode on the specified radio for an access point, set this parameter's value to "true".

	PowerAssignmentMode *int `json:"powerAssignmentMode,omitempty"` // Configure the power assignment mode on the specified radio for an access point: for global mode, set "1"; and for custom mode, set "2".

	Powerlevel *int `json:"powerlevel,omitempty"` // Configure the power level on the specified radio for an access point by setting a value between 1 and 8.

	RadioType *int `json:"radioType,omitempty"` // Configure an access point's radio band: for 2.4 GHz, set "1"; for 5 GHz, set "2"; for XOR, set "3"; and for 6 GHz, set "6".
}

//SensorTestResultsV1 Sensor Test Results - 87ae-7b21-4f0b-a838
/* Intent API to get SENSOR test result summary


@param SensorTestResultsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!sensor-test-results
*/
func (s *WirelessService) SensorTestResultsV1(SensorTestResultsV1QueryParams *SensorTestResultsV1QueryParams) (*ResponseWirelessSensorTestResultsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/AssuranceGetSensorTestResults"

	queryString, _ := query.Values(SensorTestResultsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessSensorTestResultsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SensorTestResultsV1(SensorTestResultsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation SensorTestResultsV1")
	}

	result := response.Result().(*ResponseWirelessSensorTestResultsV1)
	return result, response, err

}

//GetAccessPointRebootTaskResultV1 Get Access Point Reboot task result - c4b5-e9ce-460a-a8a3
/* Users can query the access point reboot status using this intent API


@param GetAccessPointRebootTaskResultV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-access-point-reboot-task-result
*/
func (s *WirelessService) GetAccessPointRebootTaskResultV1(GetAccessPointRebootTaskResultV1QueryParams *GetAccessPointRebootTaskResultV1QueryParams) (*ResponseWirelessGetAccessPointRebootTaskResultV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/device-reboot/apreboot/status"

	queryString, _ := query.Values(GetAccessPointRebootTaskResultV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetAccessPointRebootTaskResultV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAccessPointRebootTaskResultV1(GetAccessPointRebootTaskResultV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAccessPointRebootTaskResultV1")
	}

	result := response.Result().(*ResponseWirelessGetAccessPointRebootTaskResultV1)
	return result, response, err

}

//GetEnterpriseSSIDV1 Get Enterprise SSID - cca5-19ba-45eb-b423
/* Get Enterprise SSID


@param GetEnterpriseSSIDV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-enterprise-ssid
*/
func (s *WirelessService) GetEnterpriseSSIDV1(GetEnterpriseSSIDV1QueryParams *GetEnterpriseSSIDV1QueryParams) (*ResponseWirelessGetEnterpriseSSIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/enterprise-ssid"

	queryString, _ := query.Values(GetEnterpriseSSIDV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetEnterpriseSSIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetEnterpriseSSIDV1(GetEnterpriseSSIDV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetEnterpriseSsidV1")
	}

	result := response.Result().(*ResponseWirelessGetEnterpriseSSIDV1)
	return result, response, err

}

//GetSSIDBySiteV1 Get SSID by Site - bb92-f946-4e19-a187
/* This API allows the user to get all SSIDs (Service Set Identifier) at the given site


@param siteID siteId path parameter. Site UUID

@param GetSSIDBySiteV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-ssid-by-site
*/
func (s *WirelessService) GetSSIDBySiteV1(siteID string, GetSSIDBySiteV1QueryParams *GetSSIDBySiteV1QueryParams) (*ResponseWirelessGetSSIDBySiteV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{siteId}/wirelessSettings/ssids"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	queryString, _ := query.Values(GetSSIDBySiteV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetSSIDBySiteV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSSIDBySiteV1(siteID, GetSSIDBySiteV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSsidBySiteV1")
	}

	result := response.Result().(*ResponseWirelessGetSSIDBySiteV1)
	return result, response, err

}

//GetSSIDCountBySiteV1 Get SSID Count by Site - 52ae-589a-48ab-9116
/* This API allows the user to get count of all SSIDs (Service Set Identifier) .


@param siteID siteId path parameter. Site UUID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-ssid-count-by-site
*/
func (s *WirelessService) GetSSIDCountBySiteV1(siteID string) (*ResponseWirelessGetSSIDCountBySiteV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{siteId}/wirelessSettings/ssids/count"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetSSIDCountBySiteV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSSIDCountBySiteV1(siteID)
		}
		return nil, response, fmt.Errorf("error with operation GetSsidCountBySiteV1")
	}

	result := response.Result().(*ResponseWirelessGetSSIDCountBySiteV1)
	return result, response, err

}

//GetSSIDByIDV1 Get SSID by ID - 78a1-2804-47a9-a6a8
/* This API allows the user to get an SSID (Service Set Identifier) by ID at the given site


@param siteID siteId path parameter. Site UUID

@param id id path parameter. SSID ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-ssid-by-id
*/
func (s *WirelessService) GetSSIDByIDV1(siteID string, id string) (*ResponseWirelessGetSSIDByIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{siteId}/wirelessSettings/ssids/{id}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetSSIDByIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSSIDByIDV1(siteID, id)
		}
		return nil, response, fmt.Errorf("error with operation GetSsidByIdV1")
	}

	result := response.Result().(*ResponseWirelessGetSSIDByIDV1)
	return result, response, err

}

//GetAccessPointConfigurationCountV1 Get Access Point Configuration Count - 118b-2898-457b-8d47
/* Get Access Point Configuration Count


@param GetAccessPointConfigurationCountV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-access-point-configuration-count
*/
func (s *WirelessService) GetAccessPointConfigurationCountV1(GetAccessPointConfigurationCountV1QueryParams *GetAccessPointConfigurationCountV1QueryParams) (*ResponseWirelessGetAccessPointConfigurationCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/accesspoint-configuration/count"

	queryString, _ := query.Values(GetAccessPointConfigurationCountV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetAccessPointConfigurationCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAccessPointConfigurationCountV1(GetAccessPointConfigurationCountV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAccessPointConfigurationCountV1")
	}

	result := response.Result().(*ResponseWirelessGetAccessPointConfigurationCountV1)
	return result, response, err

}

//GetAccessPointConfigurationTaskResultV1 Get Access Point Configuration task result - fb90-69dc-4aeb-9afb
/* Users can query the access point configuration result using this intent API


@param taskTypeID task_id path parameter. task id information of ap config


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-access-point-configuration-task-result
*/
func (s *WirelessService) GetAccessPointConfigurationTaskResultV1(taskTypeID string) (*ResponseWirelessGetAccessPointConfigurationTaskResultV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/accesspoint-configuration/details/{task_id}"
	path = strings.Replace(path, "{task_id}", fmt.Sprintf("%v", taskTypeID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetAccessPointConfigurationTaskResultV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAccessPointConfigurationTaskResultV1(taskTypeID)
		}
		return nil, response, fmt.Errorf("error with operation GetAccessPointConfigurationTaskResultV1")
	}

	result := response.Result().(*ResponseWirelessGetAccessPointConfigurationTaskResultV1)
	return result, response, err

}

//GetAccessPointConfigurationV1 Get Access Point Configuration - a191-f9f2-4cb8-9a55
/* Users can query the access point configuration information per device using the ethernet MAC address


@param GetAccessPointConfigurationV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-access-point-configuration
*/
func (s *WirelessService) GetAccessPointConfigurationV1(GetAccessPointConfigurationV1QueryParams *GetAccessPointConfigurationV1QueryParams) (*ResponseWirelessGetAccessPointConfigurationV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/accesspoint-configuration/summary"

	queryString, _ := query.Values(GetAccessPointConfigurationV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetAccessPointConfigurationV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAccessPointConfigurationV1(GetAccessPointConfigurationV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAccessPointConfigurationV1")
	}

	result := response.Result().(*ResponseWirelessGetAccessPointConfigurationV1)
	return result, response, err

}

//GetDynamicInterfaceV1 Get dynamic interface - c5b0-c978-4dfb-90b4
/* Get one or all dynamic interface(s)


@param GetDynamicInterfaceV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-dynamic-interface
*/
func (s *WirelessService) GetDynamicInterfaceV1(GetDynamicInterfaceV1QueryParams *GetDynamicInterfaceV1QueryParams) (*ResponseWirelessGetDynamicInterfaceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/dynamic-interface"

	queryString, _ := query.Values(GetDynamicInterfaceV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetDynamicInterfaceV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDynamicInterfaceV1(GetDynamicInterfaceV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDynamicInterfaceV1")
	}

	result := response.Result().(*ResponseWirelessGetDynamicInterfaceV1)
	return result, response, err

}

//GetWirelessProfileV1 Get Wireless Profile - b3a1-c880-4c8b-9b8b
/* Gets either one or all the wireless network profiles if no name is provided for network-profile.


@param GetWirelessProfileV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-wireless-profile
*/
func (s *WirelessService) GetWirelessProfileV1(GetWirelessProfileV1QueryParams *GetWirelessProfileV1QueryParams) (*ResponseWirelessGetWirelessProfileV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/profile"

	queryString, _ := query.Values(GetWirelessProfileV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetWirelessProfileV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetWirelessProfileV1(GetWirelessProfileV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetWirelessProfileV1")
	}

	result := response.Result().(*ResponseWirelessGetWirelessProfileV1)
	return result, response, err

}

//RetrieveRfProfilesV1 Retrieve RF profiles - 098c-ab91-41c9-a3fe
/* Retrieve all RF profiles


@param RetrieveRFProfilesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-rf-profiles
*/
func (s *WirelessService) RetrieveRfProfilesV1(RetrieveRFProfilesV1QueryParams *RetrieveRfProfilesV1QueryParams) (*ResponseWirelessRetrieveRfProfilesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/rf-profile"

	queryString, _ := query.Values(RetrieveRFProfilesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessRetrieveRfProfilesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveRfProfilesV1(RetrieveRFProfilesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveRfProfilesV1")
	}

	result := response.Result().(*ResponseWirelessRetrieveRfProfilesV1)
	return result, response, err

}

//GetAccessPointsFactoryResetStatusV1 Get Access Point(s) Factory Reset status - 46bf-881b-45b8-a62f
/* This API returns each AP Factory Reset initiation status.


@param GetAccessPointsFactoryResetStatusV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-access-points-factory-reset-status
*/
func (s *WirelessService) GetAccessPointsFactoryResetStatusV1(GetAccessPointsFactoryResetStatusV1QueryParams *GetAccessPointsFactoryResetStatusV1QueryParams) (*ResponseWirelessGetAccessPointsFactoryResetStatusV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessAccessPoints/factoryResetRequestStatus"

	queryString, _ := query.Values(GetAccessPointsFactoryResetStatusV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetAccessPointsFactoryResetStatusV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAccessPointsFactoryResetStatusV1(GetAccessPointsFactoryResetStatusV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAccessPointsFactoryResetStatusV1")
	}

	result := response.Result().(*ResponseWirelessGetAccessPointsFactoryResetStatusV1)
	return result, response, err

}

//GetAnchorCapableDevicesV1 Get Anchor capable devices - a581-0a06-4acb-8f4c
/* This API allows the user to get Anchor capable devices



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-anchor-capable-devices
*/
func (s *WirelessService) GetAnchorCapableDevicesV1() (*ResponseWirelessGetAnchorCapableDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessControllers/anchorCapableDevices"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetAnchorCapableDevicesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAnchorCapableDevicesV1()
		}
		return nil, response, fmt.Errorf("error with operation GetAnchorCapableDevicesV1")
	}

	result := response.Result().(*ResponseWirelessGetAnchorCapableDevicesV1)
	return result, response, err

}

//GetMeshApNeighboursV1 Get Mesh Ap Neighbours - 8a88-98f8-4eca-8300
/* Retrieves all Mesh accesspoint Neighbours details whether child, parent, etc.


@param GetMeshApNeighboursV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-mesh-ap-neighbours
*/
func (s *WirelessService) GetMeshApNeighboursV1(GetMeshApNeighboursV1QueryParams *GetMeshApNeighboursV1QueryParams) (*ResponseWirelessGetMeshApNeighboursV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessControllers/meshApNeighbours"

	queryString, _ := query.Values(GetMeshApNeighboursV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetMeshApNeighboursV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetMeshApNeighboursV1(GetMeshApNeighboursV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetMeshApNeighboursV1")
	}

	result := response.Result().(*ResponseWirelessGetMeshApNeighboursV1)
	return result, response, err

}

//GetMeshApNeighboursCountV1 Get Mesh Ap Neighbours Count - 54b9-09f2-4dd8-b94f
/* This API returns the total number of mesh Ap Neighbours available.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-mesh-ap-neighbours-count
*/
func (s *WirelessService) GetMeshApNeighboursCountV1() (*ResponseWirelessGetMeshApNeighboursCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessControllers/meshApNeighbours/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetMeshApNeighboursCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetMeshApNeighboursCountV1()
		}
		return nil, response, fmt.Errorf("error with operation GetMeshApNeighboursCountV1")
	}

	result := response.Result().(*ResponseWirelessGetMeshApNeighboursCountV1)
	return result, response, err

}

//GetMobilityGroupsV1 Get MobilityGroups - 628f-38bf-4f5a-a48c
/* Retrieve configured mobility groups if no Network Device Id is provided as a query parameter. If a Network Device Id is given and a mobility group is configured for it, return the configured details; otherwise, return the default values from the device.


@param GetMobilityGroupsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-mobility-groups
*/
func (s *WirelessService) GetMobilityGroupsV1(GetMobilityGroupsV1QueryParams *GetMobilityGroupsV1QueryParams) (*ResponseWirelessGetMobilityGroupsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessControllers/wirelessMobilityGroups"

	queryString, _ := query.Values(GetMobilityGroupsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetMobilityGroupsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetMobilityGroupsV1(GetMobilityGroupsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetMobilityGroupsV1")
	}

	result := response.Result().(*ResponseWirelessGetMobilityGroupsV1)
	return result, response, err

}

//GetMobilityGroupsCountV1 Get MobilityGroups Count - 29b2-08fb-420a-8970
/* Retrieves count of mobility groups configured



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-mobility-groups-count
*/
func (s *WirelessService) GetMobilityGroupsCountV1() (*ResponseWirelessGetMobilityGroupsCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessControllers/wirelessMobilityGroups/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetMobilityGroupsCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetMobilityGroupsCountV1()
		}
		return nil, response, fmt.Errorf("error with operation GetMobilityGroupsCountV1")
	}

	result := response.Result().(*ResponseWirelessGetMobilityGroupsCountV1)
	return result, response, err

}

//GetAnchorManagedApLocationsForSpecificWirelessControllerV1 Get Anchor Managed AP Locations for specific Wireless Controller - 8dad-59b4-44b8-8995
/* Retrieves all the details of Anchor Managed AP locations associated with the specific Wireless Controller.


@param networkDeviceID networkDeviceId path parameter. Obtain the network device ID value by using the API call GET: /dna/intent/api/v1/network-device/ip-address/${ipAddress}.

@param GetAnchorManagedAPLocationsForSpecificWirelessControllerV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-anchor-managed-ap-locations-for-specific-wireless-controller
*/
func (s *WirelessService) GetAnchorManagedApLocationsForSpecificWirelessControllerV1(networkDeviceID string, GetAnchorManagedAPLocationsForSpecificWirelessControllerV1QueryParams *GetAnchorManagedApLocationsForSpecificWirelessControllerV1QueryParams) (*ResponseWirelessGetAnchorManagedApLocationsForSpecificWirelessControllerV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessControllers/{networkDeviceId}/anchorManagedApLocations"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	queryString, _ := query.Values(GetAnchorManagedAPLocationsForSpecificWirelessControllerV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetAnchorManagedApLocationsForSpecificWirelessControllerV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAnchorManagedApLocationsForSpecificWirelessControllerV1(networkDeviceID, GetAnchorManagedAPLocationsForSpecificWirelessControllerV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAnchorManagedApLocationsForSpecificWirelessControllerV1")
	}

	result := response.Result().(*ResponseWirelessGetAnchorManagedApLocationsForSpecificWirelessControllerV1)
	return result, response, err

}

//GetApAuthorizationListByNetworkDeviceIDV1 Get AP Authorization List by network device Id - c689-88c5-4128-a366
/* This API allows the user to get an AP Authorization List details configured for the given provisioned network device Id. Obtain the network device ID value by using the API GET call '/dna/intent/api/v1/network-device/ip-address/${ipAddress}'.


@param networkDeviceID networkDeviceId path parameter. Network Device ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-ap-authorization-list-by-network-device-id
*/
func (s *WirelessService) GetApAuthorizationListByNetworkDeviceIDV1(networkDeviceID string) (*ResponseWirelessGetApAuthorizationListByNetworkDeviceIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessControllers/{networkDeviceId}/apAuthorizationLists"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetApAuthorizationListByNetworkDeviceIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApAuthorizationListByNetworkDeviceIDV1(networkDeviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetApAuthorizationListByNetworkDeviceIdV1")
	}

	result := response.Result().(*ResponseWirelessGetApAuthorizationListByNetworkDeviceIDV1)
	return result, response, err

}

//GetManagedApLocationsCountForSpecificWirelessControllerV1 Get Managed AP Locations Count for specific Wireless Controller - f490-6a9b-4c29-bc6a
/* Retrieves the count of Managed AP locations, including Primary Managed AP Locations, Secondary Managed AP Locations, and Anchor Managed AP Locations, associated with the specific Wireless Controller.


@param networkDeviceID networkDeviceId path parameter. Obtain the network device ID value by using the API call GET: /dna/intent/api/v1/network-device/ip-address/${ipAddress}.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-managed-ap-locations-count-for-specific-wireless-controller
*/
func (s *WirelessService) GetManagedApLocationsCountForSpecificWirelessControllerV1(networkDeviceID string) (*ResponseWirelessGetManagedApLocationsCountForSpecificWirelessControllerV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessControllers/{networkDeviceId}/managedApLocations/count"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetManagedApLocationsCountForSpecificWirelessControllerV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetManagedApLocationsCountForSpecificWirelessControllerV1(networkDeviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetManagedApLocationsCountForSpecificWirelessControllerV1")
	}

	result := response.Result().(*ResponseWirelessGetManagedApLocationsCountForSpecificWirelessControllerV1)
	return result, response, err

}

//GetPrimaryManagedApLocationsForSpecificWirelessControllerV1 Get Primary Managed AP Locations for specific Wireless Controller - 1dba-89f4-40ab-abda
/* Retrieves all the details of Primary Managed AP locations associated with the specific Wireless Controller.


@param networkDeviceID networkDeviceId path parameter. Obtain the network device ID value by using the API call GET: /dna/intent/api/v1/network-device/ip-address/${ipAddress}.

@param GetPrimaryManagedAPLocationsForSpecificWirelessControllerV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-primary-managed-ap-locations-for-specific-wireless-controller
*/
func (s *WirelessService) GetPrimaryManagedApLocationsForSpecificWirelessControllerV1(networkDeviceID string, GetPrimaryManagedAPLocationsForSpecificWirelessControllerV1QueryParams *GetPrimaryManagedApLocationsForSpecificWirelessControllerV1QueryParams) (*ResponseWirelessGetPrimaryManagedApLocationsForSpecificWirelessControllerV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessControllers/{networkDeviceId}/primaryManagedApLocations"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	queryString, _ := query.Values(GetPrimaryManagedAPLocationsForSpecificWirelessControllerV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetPrimaryManagedApLocationsForSpecificWirelessControllerV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPrimaryManagedApLocationsForSpecificWirelessControllerV1(networkDeviceID, GetPrimaryManagedAPLocationsForSpecificWirelessControllerV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetPrimaryManagedApLocationsForSpecificWirelessControllerV1")
	}

	result := response.Result().(*ResponseWirelessGetPrimaryManagedApLocationsForSpecificWirelessControllerV1)
	return result, response, err

}

//GetSecondaryManagedApLocationsForSpecificWirelessControllerV1 Get Secondary Managed AP Locations for specific Wireless Controller - b589-7bd6-4f1b-9efb
/* Retrieves all the details of Secondary Managed AP locations associated with the specific Wireless Controller.


@param networkDeviceID networkDeviceId path parameter. Obtain the network device ID value by using the API call GET: /dna/intent/api/v1/network-device/ip-address/${ipAddress}.

@param GetSecondaryManagedAPLocationsForSpecificWirelessControllerV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-secondary-managed-ap-locations-for-specific-wireless-controller
*/
func (s *WirelessService) GetSecondaryManagedApLocationsForSpecificWirelessControllerV1(networkDeviceID string, GetSecondaryManagedAPLocationsForSpecificWirelessControllerV1QueryParams *GetSecondaryManagedApLocationsForSpecificWirelessControllerV1QueryParams) (*ResponseWirelessGetSecondaryManagedApLocationsForSpecificWirelessControllerV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessControllers/{networkDeviceId}/secondaryManagedApLocations"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	queryString, _ := query.Values(GetSecondaryManagedAPLocationsForSpecificWirelessControllerV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetSecondaryManagedApLocationsForSpecificWirelessControllerV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSecondaryManagedApLocationsForSpecificWirelessControllerV1(networkDeviceID, GetSecondaryManagedAPLocationsForSpecificWirelessControllerV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSecondaryManagedApLocationsForSpecificWirelessControllerV1")
	}

	result := response.Result().(*ResponseWirelessGetSecondaryManagedApLocationsForSpecificWirelessControllerV1)
	return result, response, err

}

//GetSSIDDetailsForSpecificWirelessControllerV1 Get SSID Details for specific Wireless Controller - 70b6-393d-4899-ad4d
/* Retrieves all details of SSIDs associated with the specific Wireless Controller.


@param networkDeviceID networkDeviceId path parameter. Obtain the network device ID value by using the API call GET: /dna/intent/api/v1/network-device/ip-address/${ipAddress}.

@param GetSSIDDetailsForSpecificWirelessControllerV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-ssid-details-for-specific-wireless-controller
*/
func (s *WirelessService) GetSSIDDetailsForSpecificWirelessControllerV1(networkDeviceID string, GetSSIDDetailsForSpecificWirelessControllerV1QueryParams *GetSSIDDetailsForSpecificWirelessControllerV1QueryParams) (*ResponseWirelessGetSSIDDetailsForSpecificWirelessControllerV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessControllers/{networkDeviceId}/ssidDetails"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	queryString, _ := query.Values(GetSSIDDetailsForSpecificWirelessControllerV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetSSIDDetailsForSpecificWirelessControllerV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSSIDDetailsForSpecificWirelessControllerV1(networkDeviceID, GetSSIDDetailsForSpecificWirelessControllerV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSsidDetailsForSpecificWirelessControllerV1")
	}

	result := response.Result().(*ResponseWirelessGetSSIDDetailsForSpecificWirelessControllerV1)
	return result, response, err

}

//GetSSIDCountForSpecificWirelessControllerV1 Get SSID Count for specific Wireless Controller - 3e98-c91d-42eb-a469
/* Retrieves the count of SSIDs associated with the specific wireless controller.


@param networkDeviceID networkDeviceId path parameter. Obtain the network device ID value by using the API call GET: /dna/intent/api/v1/network-device/ip-address/${ipAddress}.

@param GetSSIDCountForSpecificWirelessControllerV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-ssid-count-for-specific-wireless-controller
*/
func (s *WirelessService) GetSSIDCountForSpecificWirelessControllerV1(networkDeviceID string, GetSSIDCountForSpecificWirelessControllerV1QueryParams *GetSSIDCountForSpecificWirelessControllerV1QueryParams) (*ResponseWirelessGetSSIDCountForSpecificWirelessControllerV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessControllers/{networkDeviceId}/ssidDetails/count"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	queryString, _ := query.Values(GetSSIDCountForSpecificWirelessControllerV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetSSIDCountForSpecificWirelessControllerV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSSIDCountForSpecificWirelessControllerV1(networkDeviceID, GetSSIDCountForSpecificWirelessControllerV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSsidCountForSpecificWirelessControllerV1")
	}

	result := response.Result().(*ResponseWirelessGetSSIDCountForSpecificWirelessControllerV1)
	return result, response, err

}

//GetWirelessProfilesV1 Get Wireless Profiles - 7988-fac4-447b-8e3d
/* This API allows the user to get all Wireless Network Profiles


@param GetWirelessProfilesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-wireless-profiles
*/
func (s *WirelessService) GetWirelessProfilesV1(GetWirelessProfilesV1QueryParams *GetWirelessProfilesV1QueryParams) (*ResponseWirelessGetWirelessProfilesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessProfiles"

	queryString, _ := query.Values(GetWirelessProfilesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetWirelessProfilesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetWirelessProfilesV1(GetWirelessProfilesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetWirelessProfilesV1")
	}

	result := response.Result().(*ResponseWirelessGetWirelessProfilesV1)
	return result, response, err

}

//GetWirelessProfilesCountV1 Get Wireless Profiles Count - 48a7-1883-48fb-93a5
/* This API allows the user to get count of all wireless profiles



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-wireless-profiles-count
*/
func (s *WirelessService) GetWirelessProfilesCountV1() (*ResponseWirelessGetWirelessProfilesCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessProfiles/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetWirelessProfilesCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetWirelessProfilesCountV1()
		}
		return nil, response, fmt.Errorf("error with operation GetWirelessProfilesCountV1")
	}

	result := response.Result().(*ResponseWirelessGetWirelessProfilesCountV1)
	return result, response, err

}

//GetWirelessProfileByIDV1 Get Wireless Profile by ID - f5b9-fab9-4b79-b0f3
/* This API allows the user to get a Wireless Network Profile by ID


@param id id path parameter. Wireless Profile Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-wireless-profile-by-id
*/
func (s *WirelessService) GetWirelessProfileByIDV1(id string) (*ResponseWirelessGetWirelessProfileByIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessProfiles/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetWirelessProfileByIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetWirelessProfileByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation GetWirelessProfileByIdV1")
	}

	result := response.Result().(*ResponseWirelessGetWirelessProfileByIDV1)
	return result, response, err

}

//RetrieveAllPolicyTagsForAWirelessProfileV1 Retrieve all Policy Tags for a Wireless Profile - 428d-d8f7-4fa9-a0ca
/* This endpoint retrieves a list of all `Policy Tags` associated with a specific `Wireless Profile`. This API requires the `id` of the `Wireless Profile` to be provided as a path parameter.


@param id id path parameter. Wireless Profile Id

@param RetrieveAllPolicyTagsForAWirelessProfileV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-all-policy-tags-for-a-wireless-profile
*/
func (s *WirelessService) RetrieveAllPolicyTagsForAWirelessProfileV1(id string, RetrieveAllPolicyTagsForAWirelessProfileV1QueryParams *RetrieveAllPolicyTagsForAWirelessProfileV1QueryParams) (*ResponseWirelessRetrieveAllPolicyTagsForAWirelessProfileV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessProfiles/{id}/policyTags"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(RetrieveAllPolicyTagsForAWirelessProfileV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessRetrieveAllPolicyTagsForAWirelessProfileV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveAllPolicyTagsForAWirelessProfileV1(id, RetrieveAllPolicyTagsForAWirelessProfileV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveAllPolicyTagsForAWirelessProfileV1")
	}

	result := response.Result().(*ResponseWirelessRetrieveAllPolicyTagsForAWirelessProfileV1)
	return result, response, err

}

//RetrieveTheCountOfPolicyTagsForAWirelessProfileV1 Retrieve the count of Policy Tags for a Wireless Profile - 26be-7947-4a69-be88
/* This endpoint retrieves the total count of `Policy Tags` associated with a specific `Wireless Profile`.This API requires the `id` of the `Wireless Profile` to be provided as a path parameter.


@param id id path parameter. Wireless Profile Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-the-count-of-policy-tags-for-a-wireless-profile
*/
func (s *WirelessService) RetrieveTheCountOfPolicyTagsForAWirelessProfileV1(id string) (*ResponseWirelessRetrieveTheCountOfPolicyTagsForAWirelessProfileV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessProfiles/{id}/policyTags/count"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessRetrieveTheCountOfPolicyTagsForAWirelessProfileV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveTheCountOfPolicyTagsForAWirelessProfileV1(id)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveTheCountOfPolicyTagsForAWirelessProfileV1")
	}

	result := response.Result().(*ResponseWirelessRetrieveTheCountOfPolicyTagsForAWirelessProfileV1)
	return result, response, err

}

//RetrieveASpecificPolicyTagForAWirelessProfileV1 Retrieve a specific Policy Tag for a Wireless Profile - 558d-4b8c-4149-b325
/* This endpoint retrieves the details of a specific `Policy Tag` associated with a given `Wireless Profile`. This API requires the `id` of the `Wireless Profile` and the `policyTagId` of the `Policy Tag`.


@param id id path parameter. Wireless Profile Id

@param policyTagID policyTagId path parameter. Policy Tag Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-a-specific-policy-tag-for-a-wireless-profile
*/
func (s *WirelessService) RetrieveASpecificPolicyTagForAWirelessProfileV1(id string, policyTagID string) (*ResponseWirelessRetrieveASpecificPolicyTagForAWirelessProfileV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessProfiles/{id}/policyTags/{policyTagId}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{policyTagId}", fmt.Sprintf("%v", policyTagID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessRetrieveASpecificPolicyTagForAWirelessProfileV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveASpecificPolicyTagForAWirelessProfileV1(id, policyTagID)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveASpecificPolicyTagForAWirelessProfileV1")
	}

	result := response.Result().(*ResponseWirelessRetrieveASpecificPolicyTagForAWirelessProfileV1)
	return result, response, err

}

//RetrieveAllSiteTagsForAWirelessProfileV1 Retrieve all Site Tags for a Wireless Profile - bd9d-5a4b-4f5a-ac11
/* This endpoint retrieves a list of all `Site Tags` associated with a specific `Wireless Profile`. This API requires the `id` of the `Wireless Profile` to be provided as a path parameter.


@param id id path parameter. Wireless profile id

@param RetrieveAllSiteTagsForAWirelessProfileV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-all-site-tags-for-a-wireless-profile
*/
func (s *WirelessService) RetrieveAllSiteTagsForAWirelessProfileV1(id string, RetrieveAllSiteTagsForAWirelessProfileV1QueryParams *RetrieveAllSiteTagsForAWirelessProfileV1QueryParams) (*ResponseWirelessRetrieveAllSiteTagsForAWirelessProfileV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessProfiles/{id}/siteTags"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(RetrieveAllSiteTagsForAWirelessProfileV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessRetrieveAllSiteTagsForAWirelessProfileV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveAllSiteTagsForAWirelessProfileV1(id, RetrieveAllSiteTagsForAWirelessProfileV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveAllSiteTagsForAWirelessProfileV1")
	}

	result := response.Result().(*ResponseWirelessRetrieveAllSiteTagsForAWirelessProfileV1)
	return result, response, err

}

//RetrieveTheCountOfSiteTagsForAWirelessProfileV1 Retrieve the count of Site Tags for a Wireless Profile - bdab-4896-4fa8-bbcc
/* This endpoint retrieves the total count of `Site Tags` associated with a specific `Wireless Profile`.This API requires the `id` of the `Wireless Profile` to be provided as a path parameter.


@param id id path parameter. Wireless profile id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-the-count-of-site-tags-for-a-wireless-profile
*/
func (s *WirelessService) RetrieveTheCountOfSiteTagsForAWirelessProfileV1(id string) (*ResponseWirelessRetrieveTheCountOfSiteTagsForAWirelessProfileV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessProfiles/{id}/siteTags/count"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessRetrieveTheCountOfSiteTagsForAWirelessProfileV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveTheCountOfSiteTagsForAWirelessProfileV1(id)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveTheCountOfSiteTagsForAWirelessProfileV1")
	}

	result := response.Result().(*ResponseWirelessRetrieveTheCountOfSiteTagsForAWirelessProfileV1)
	return result, response, err

}

//RetrieveASpecificSiteTagForAWirelessProfileV1 Retrieve a specific Site Tag for a Wireless Profile - 1fac-e966-4e6b-96ea
/* This endpoint retrieves the details of a specific `Site Tag` associated with a given `Wireless Profile`. This API requires the `id` of the `Wireless Profile` and the `siteTagId` of the `Site Tag`.


@param id id path parameter. Wireless Profile Id

@param siteTagID siteTagId path parameter. Site Tag Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-a-specific-site-tag-for-a-wireless-profile
*/
func (s *WirelessService) RetrieveASpecificSiteTagForAWirelessProfileV1(id string, siteTagID string) (*ResponseWirelessRetrieveASpecificSiteTagForAWirelessProfileV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessProfiles/{id}/siteTags/{siteTagId}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{siteTagId}", fmt.Sprintf("%v", siteTagID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessRetrieveASpecificSiteTagForAWirelessProfileV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveASpecificSiteTagForAWirelessProfileV1(id, siteTagID)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveASpecificSiteTagForAWirelessProfileV1")
	}

	result := response.Result().(*ResponseWirelessRetrieveASpecificSiteTagForAWirelessProfileV1)
	return result, response, err

}

//GetAnchorGroupsV1 Get AnchorGroups - 32b3-aa83-46db-aae7
/* This API allows the user to get AnchorGroups that captured in wireless settings design.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-anchor-groups
*/
func (s *WirelessService) GetAnchorGroupsV1() (*ResponseWirelessGetAnchorGroupsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/anchorGroups"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetAnchorGroupsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAnchorGroupsV1()
		}
		return nil, response, fmt.Errorf("error with operation GetAnchorGroupsV1")
	}

	result := response.Result().(*ResponseWirelessGetAnchorGroupsV1)
	return result, response, err

}

//GetCountOfAnchorGroupsV1 Get count of AnchorGroups - 5581-4892-4dc8-a66c
/* This API allows the user to get count of all AnchorGroups



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-count-of-anchor-groups
*/
func (s *WirelessService) GetCountOfAnchorGroupsV1() (*ResponseWirelessGetCountOfAnchorGroupsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/anchorGroups/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetCountOfAnchorGroupsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCountOfAnchorGroupsV1()
		}
		return nil, response, fmt.Errorf("error with operation GetCountOfAnchorGroupsV1")
	}

	result := response.Result().(*ResponseWirelessGetCountOfAnchorGroupsV1)
	return result, response, err

}

//GetAnchorGroupByIDV1 Get AnchorGroup by ID - 029f-4acf-420b-9df6
/* This API allows the user to get an AnchorGroup by AnchorGroup ID


@param id id path parameter. AnchorGroup ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-anchor-group-by-id
*/
func (s *WirelessService) GetAnchorGroupByIDV1(id string) (*ResponseWirelessGetAnchorGroupByIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/anchorGroups/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetAnchorGroupByIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAnchorGroupByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation GetAnchorGroupByIdV1")
	}

	result := response.Result().(*ResponseWirelessGetAnchorGroupByIDV1)
	return result, response, err

}

//GetApAuthorizationListsV1 Get AP Authorization Lists - f7b1-b801-4738-937d
/* Retrieves the AP Authorization Lists that are created in the Catalyst Centre network Design for wireless. If an AP Authorization List name is given as query parameter, then returns respective AP Authorization List details including Local and/or Remote authorization.


@param GetAPAuthorizationListsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-ap-authorization-lists
*/
func (s *WirelessService) GetApAuthorizationListsV1(GetAPAuthorizationListsV1QueryParams *GetApAuthorizationListsV1QueryParams) (*ResponseWirelessGetApAuthorizationListsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/apAuthorizationLists"

	queryString, _ := query.Values(GetAPAuthorizationListsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetApAuthorizationListsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApAuthorizationListsV1(GetAPAuthorizationListsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetApAuthorizationListsV1")
	}

	result := response.Result().(*ResponseWirelessGetApAuthorizationListsV1)
	return result, response, err

}

//GetApAuthorizationListCountV1 Get AP Authorization List Count - 51be-d862-47c8-a51a
/* This API allows the user to get count of all AP Authorization lists.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-ap-authorization-list-count
*/
func (s *WirelessService) GetApAuthorizationListCountV1() (*ResponseWirelessGetApAuthorizationListCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/apAuthorizationLists/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetApAuthorizationListCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApAuthorizationListCountV1()
		}
		return nil, response, fmt.Errorf("error with operation GetApAuthorizationListCountV1")
	}

	result := response.Result().(*ResponseWirelessGetApAuthorizationListCountV1)
	return result, response, err

}

//GetApAuthorizationListByIDV1 Get AP Authorization List by ID - 0aa3-1b69-4a58-b0f5
/* This API allows the user to get an AP Authorization List by AP Authorization List ID that captured in wireless settings design.


@param id id path parameter. AP Authorization List ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-ap-authorization-list-by-id
*/
func (s *WirelessService) GetApAuthorizationListByIDV1(id string) (*ResponseWirelessGetApAuthorizationListByIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/apAuthorizationLists/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetApAuthorizationListByIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApAuthorizationListByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation GetApAuthorizationListByIdV1")
	}

	result := response.Result().(*ResponseWirelessGetApAuthorizationListByIDV1)
	return result, response, err

}

//GetApProfilesV1 Get AP Profiles - edad-9bfa-4298-a4cb
/* This API allows the user to get AP profiles that are captured in wireless settings design.


@param GetAPProfilesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-ap-profiles
*/
func (s *WirelessService) GetApProfilesV1(GetAPProfilesV1QueryParams *GetApProfilesV1QueryParams) (*ResponseWirelessGetApProfilesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/apProfiles"

	queryString, _ := query.Values(GetAPProfilesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetApProfilesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApProfilesV1(GetAPProfilesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetApProfilesV1")
	}

	result := response.Result().(*ResponseWirelessGetApProfilesV1)
	return result, response, err

}

//GetApProfilesCountV1 Get AP Profiles Count - a687-f85e-438a-8941
/* This API returns the total number of AP Profiles available.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-ap-profiles-count
*/
func (s *WirelessService) GetApProfilesCountV1() (*ResponseWirelessGetApProfilesCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/apProfiles/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetApProfilesCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApProfilesCountV1()
		}
		return nil, response, fmt.Errorf("error with operation GetApProfilesCountV1")
	}

	result := response.Result().(*ResponseWirelessGetApProfilesCountV1)
	return result, response, err

}

//GetApProfileByIDV1 Get AP Profile by ID - ba9f-5899-4c5b-87f2
/* This API allows the user to get a AP Profile by AP Profile ID that captured in wireless settings design


@param id id path parameter. Ap Profile ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-ap-profile-by-id
*/
func (s *WirelessService) GetApProfileByIDV1(id string) (*ResponseWirelessGetApProfileByIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/apProfiles/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetApProfileByIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApProfileByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation GetApProfileByIdV1")
	}

	result := response.Result().(*ResponseWirelessGetApProfileByIDV1)
	return result, response, err

}

//Get80211BeProfilesV1 Get 802.11be Profiles - 1895-aac1-4428-bd0d
/* This API allows the user to get 802.11be Profile(s) configured under Wireless Settings


@param Get80211beProfilesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get80211be-profiles
*/
func (s *WirelessService) Get80211BeProfilesV1(Get80211beProfilesV1QueryParams *Get80211BeProfilesV1QueryParams) (*ResponseWirelessGet80211BeProfilesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/dot11beProfiles"

	queryString, _ := query.Values(Get80211beProfilesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGet80211BeProfilesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.Get80211BeProfilesV1(Get80211beProfilesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation Get80211BeProfilesV1")
	}

	result := response.Result().(*ResponseWirelessGet80211BeProfilesV1)
	return result, response, err

}

//Get80211BeProfilesCountV1 Get 802.11be Profiles Count - a0b7-da85-4faa-95b7
/* This API allows the user to get count of all 802.11be Profile(s)



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get80211be-profiles-count
*/
func (s *WirelessService) Get80211BeProfilesCountV1() (*ResponseWirelessGet80211BeProfilesCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/dot11beProfiles/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGet80211BeProfilesCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.Get80211BeProfilesCountV1()
		}
		return nil, response, fmt.Errorf("error with operation Get80211BeProfilesCountV1")
	}

	result := response.Result().(*ResponseWirelessGet80211BeProfilesCountV1)
	return result, response, err

}

//Get80211BeProfileByIDV1 Get 802.11be Profile by ID - fa93-88ce-49eb-a5d7
/* This API allows the user to get 802.11be Profile by ID


@param id id path parameter. 802.11be Profile ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get80211be-profile-by-id
*/
func (s *WirelessService) Get80211BeProfileByIDV1(id string) (*ResponseWirelessGet80211BeProfileByIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/dot11beProfiles/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGet80211BeProfileByIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.Get80211BeProfileByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation Get80211BeProfileByIdV1")
	}

	result := response.Result().(*ResponseWirelessGet80211BeProfileByIDV1)
	return result, response, err

}

//GetInterfacesV1 Get Interfaces - 3793-ea73-438a-b243
/* This API allows the user to get all Interfaces


@param GetInterfacesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-interfaces
*/
func (s *WirelessService) GetInterfacesV1(GetInterfacesV1QueryParams *GetInterfacesV1QueryParams) (*ResponseWirelessGetInterfacesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/interfaces"

	queryString, _ := query.Values(GetInterfacesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetInterfacesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetInterfacesV1(GetInterfacesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetInterfacesV1")
	}

	result := response.Result().(*ResponseWirelessGetInterfacesV1)
	return result, response, err

}

//GetInterfacesCountV1 Get Interfaces Count - fd81-f950-424b-b992
/* This API allows the user to get count of all interfaces



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-interfaces-count
*/
func (s *WirelessService) GetInterfacesCountV1() (*ResponseWirelessGetInterfacesCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/interfaces/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetInterfacesCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetInterfacesCountV1()
		}
		return nil, response, fmt.Errorf("error with operation GetInterfacesCountV1")
	}

	result := response.Result().(*ResponseWirelessGetInterfacesCountV1)
	return result, response, err

}

//GetInterfaceByIDV1 Get Interface by ID - 3fa4-19ab-482a-ad07
/* This API allows the user to get an interface by ID


@param id id path parameter. Interface ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-interface-by-id
*/
func (s *WirelessService) GetInterfaceByIDV1(id string) (*ResponseWirelessGetInterfaceByIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/interfaces/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetInterfaceByIDV1{}).
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

	result := response.Result().(*ResponseWirelessGetInterfaceByIDV1)
	return result, response, err

}

//GetPowerProfilesV1 Get Power Profiles - a9b3-d8c2-4b6b-9b01
/* This API allows the user to get Power Profiles that captured in wireless settings design.


@param GetPowerProfilesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-power-profiles
*/
func (s *WirelessService) GetPowerProfilesV1(GetPowerProfilesV1QueryParams *GetPowerProfilesV1QueryParams) (*ResponseWirelessGetPowerProfilesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/powerProfiles"

	queryString, _ := query.Values(GetPowerProfilesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetPowerProfilesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPowerProfilesV1(GetPowerProfilesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetPowerProfilesV1")
	}

	result := response.Result().(*ResponseWirelessGetPowerProfilesV1)
	return result, response, err

}

//GetPowerProfilesCountV1 Get Power Profiles Count - 7091-1b6d-4849-ab2f
/* This API returns the total number of Power Profiles available.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-power-profiles-count
*/
func (s *WirelessService) GetPowerProfilesCountV1() (*ResponseWirelessGetPowerProfilesCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/powerProfiles/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetPowerProfilesCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPowerProfilesCountV1()
		}
		return nil, response, fmt.Errorf("error with operation GetPowerProfilesCountV1")
	}

	result := response.Result().(*ResponseWirelessGetPowerProfilesCountV1)
	return result, response, err

}

//GetPowerProfileByIDV1 Get Power Profile by ID - 6c93-cb96-45b8-a53b
/* This API allows the user to get a Power Profile by Power Profile ID that captured in wireless settings design


@param id id path parameter. Power Profile ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-power-profile-by-id
*/
func (s *WirelessService) GetPowerProfileByIDV1(id string) (*ResponseWirelessGetPowerProfileByIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/powerProfiles/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetPowerProfileByIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPowerProfileByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation GetPowerProfileByIdV1")
	}

	result := response.Result().(*ResponseWirelessGetPowerProfileByIDV1)
	return result, response, err

}

//GetRfProfilesV1 Get RF Profiles - 15a6-e823-49ca-a9cc
/* This API allows the user to get all RF Profiles


@param GetRFProfilesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-rf-profiles
*/
func (s *WirelessService) GetRfProfilesV1(GetRFProfilesV1QueryParams *GetRfProfilesV1QueryParams) (*ResponseWirelessGetRfProfilesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/rfProfiles"

	queryString, _ := query.Values(GetRFProfilesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetRfProfilesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetRfProfilesV1(GetRFProfilesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetRfProfilesV1")
	}

	result := response.Result().(*ResponseWirelessGetRfProfilesV1)
	return result, response, err

}

//GetRfProfilesCountV1 Get RF Profiles Count - f996-2b80-477a-9de2
/* This API allows the user to get count of all RF profiles



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-rf-profiles-count
*/
func (s *WirelessService) GetRfProfilesCountV1() (*ResponseWirelessGetRfProfilesCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/rfProfiles/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetRfProfilesCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetRfProfilesCountV1()
		}
		return nil, response, fmt.Errorf("error with operation GetRfProfilesCountV1")
	}

	result := response.Result().(*ResponseWirelessGetRfProfilesCountV1)
	return result, response, err

}

//GetRfProfileByIDV1 Get RF Profile by ID - 3298-aa56-4ec9-b510
/* This API allows the user to get a RF Profile by RF Profile ID


@param id id path parameter. RF Profile ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-rf-profile-by-id
*/
func (s *WirelessService) GetRfProfileByIDV1(id string) (*ResponseWirelessGetRfProfileByIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/rfProfiles/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetRfProfileByIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetRfProfileByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation GetRfProfileByIdV1")
	}

	result := response.Result().(*ResponseWirelessGetRfProfileByIDV1)
	return result, response, err

}

//RetrieveSitesWithOverriddenSSIDsV1 Retrieve sites with overridden SSIDs - 9a9a-8b8b-4029-a86e
/* Retrieve list of siteId(s) with information of SSID(s) which are overridden


@param RetrieveSitesWithOverriddenSSIDsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-sites-with-overridden-ssids
*/
func (s *WirelessService) RetrieveSitesWithOverriddenSSIDsV1(RetrieveSitesWithOverriddenSSIDsV1QueryParams *RetrieveSitesWithOverriddenSSIDsV1QueryParams) (*ResponseWirelessRetrieveSitesWithOverriddenSSIDsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/ssids/overrideAtSites"

	queryString, _ := query.Values(RetrieveSitesWithOverriddenSSIDsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessRetrieveSitesWithOverriddenSSIDsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveSitesWithOverriddenSSIDsV1(RetrieveSitesWithOverriddenSSIDsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveSitesWithOverriddenSsidsV1")
	}

	result := response.Result().(*ResponseWirelessRetrieveSitesWithOverriddenSSIDsV1)
	return result, response, err

}

//CreateAndProvisionSSIDV1 Create and Provision SSID - 1eb7-2ad3-4e09-8990
/* Creates SSID, updates the SSID to the corresponding site profiles and provision it to the devices matching the given sites


@param CreateAndProvisionSSIDV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-and-provision-ssid
*/
func (s *WirelessService) CreateAndProvisionSSIDV1(requestWirelessCreateAndProvisionSSIDV1 *RequestWirelessCreateAndProvisionSSIDV1, CreateAndProvisionSSIDV1HeaderParams *CreateAndProvisionSSIDV1HeaderParams) (*ResponseWirelessCreateAndProvisionSSIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/ssid"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if CreateAndProvisionSSIDV1HeaderParams != nil {

		if CreateAndProvisionSSIDV1HeaderParams.Persistbapioutput != "" {
			clientRequest = clientRequest.SetHeader("__persistbapioutput", CreateAndProvisionSSIDV1HeaderParams.Persistbapioutput)
		}

	}

	response, err = clientRequest.
		SetBody(requestWirelessCreateAndProvisionSSIDV1).
		SetResult(&ResponseWirelessCreateAndProvisionSSIDV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateAndProvisionSSIDV1(requestWirelessCreateAndProvisionSSIDV1, CreateAndProvisionSSIDV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation CreateAndProvisionSsidV1")
	}

	result := response.Result().(*ResponseWirelessCreateAndProvisionSSIDV1)
	return result, response, err

}

//RebootAccessPointsV1 Reboot Access Points - 6092-d8f1-468b-99ab
/* Users can reboot multiple access points up-to 200 at a time using this API



Documentation Link: https://developer.cisco.com/docs/dna-center/#!reboot-access-points
*/
func (s *WirelessService) RebootAccessPointsV1(requestWirelessRebootAccessPointsV1 *RequestWirelessRebootAccessPointsV1) (*ResponseWirelessRebootAccessPointsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/device-reboot/apreboot"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessRebootAccessPointsV1).
		SetResult(&ResponseWirelessRebootAccessPointsV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RebootAccessPointsV1(requestWirelessRebootAccessPointsV1)
		}

		return nil, response, fmt.Errorf("error with operation RebootAccessPointsV1")
	}

	result := response.Result().(*ResponseWirelessRebootAccessPointsV1)
	return result, response, err

}

//CreateEnterpriseSSIDV1 Create Enterprise SSID - 8a96-fb95-4d09-a349
/* Creates enterprise SSID



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-enterprise-ssid
*/
func (s *WirelessService) CreateEnterpriseSSIDV1(requestWirelessCreateEnterpriseSSIDV1 *RequestWirelessCreateEnterpriseSSIDV1) (*ResponseWirelessCreateEnterpriseSSIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/enterprise-ssid"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessCreateEnterpriseSSIDV1).
		SetResult(&ResponseWirelessCreateEnterpriseSSIDV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateEnterpriseSSIDV1(requestWirelessCreateEnterpriseSSIDV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateEnterpriseSsidV1")
	}

	result := response.Result().(*ResponseWirelessCreateEnterpriseSSIDV1)
	return result, response, err

}

//CreateSSIDV1 Create SSID - 0193-8858-4789-9a53
/* This API allows the user to create an SSID (Service Set Identifier) at the Global site


@param siteID siteId path parameter. Site UUID of Global site


Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-ssid
*/
func (s *WirelessService) CreateSSIDV1(siteID string, requestWirelessCreateSSIDV1 *RequestWirelessCreateSSIDV1) (*ResponseWirelessCreateSSIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{siteId}/wirelessSettings/ssids"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessCreateSSIDV1).
		SetResult(&ResponseWirelessCreateSSIDV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateSSIDV1(siteID, requestWirelessCreateSSIDV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateSsidV1")
	}

	result := response.Result().(*ResponseWirelessCreateSSIDV1)
	return result, response, err

}

//UpdateOrOverridessidV1 Update or Override SSID - 559d-88ff-43c9-9fe3
/* This API allows to either update SSID at global 'siteId' or override SSID at given non-global 'siteId'


@param siteID siteId path parameter. Site UUID

@param id id path parameter. SSID ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!update-or-overridessid
*/
func (s *WirelessService) UpdateOrOverridessidV1(siteID string, id string, requestWirelessUpdateOrOverrideSSIDV1 *RequestWirelessUpdateOrOverridessidV1) (*ResponseWirelessUpdateOrOverridessidV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{siteId}/wirelessSettings/ssids/{id}/update"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateOrOverrideSSIDV1).
		SetResult(&ResponseWirelessUpdateOrOverridessidV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateOrOverridessidV1(siteID, id, requestWirelessUpdateOrOverrideSSIDV1)
		}

		return nil, response, fmt.Errorf("error with operation UpdateOrOverridessidV1")
	}

	result := response.Result().(*ResponseWirelessUpdateOrOverridessidV1)
	return result, response, err

}

//ConfigureAccessPointsV1 Configure Access Points V1 - 0081-cb89-4708-888f
/* User can configure multiple access points with required options using this intent API. This API does not support configuration of CleanAir or SI for IOS-XE devices with version greater than or equal to 17.9



Documentation Link: https://developer.cisco.com/docs/dna-center/#!configure-access-points
*/
func (s *WirelessService) ConfigureAccessPointsV1(requestWirelessConfigureAccessPointsV1 *RequestWirelessConfigureAccessPointsV1) (*ResponseWirelessConfigureAccessPointsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/accesspoint-configuration"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessConfigureAccessPointsV1).
		SetResult(&ResponseWirelessConfigureAccessPointsV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ConfigureAccessPointsV1(requestWirelessConfigureAccessPointsV1)
		}

		return nil, response, fmt.Errorf("error with operation ConfigureAccessPointsV1")
	}

	result := response.Result().(*ResponseWirelessConfigureAccessPointsV1)
	return result, response, err

}

//ApProvisionConnectivityV1 AP Provision - d897-19b8-47aa-a9c4
/* Access Point Provision and ReProvision


@param APProvisionConnectivityV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!ap-provision-connectivity
*/
func (s *WirelessService) ApProvisionConnectivityV1(requestWirelessAPProvisionConnectivityV1 *RequestWirelessApProvisionConnectivityV1, APProvisionConnectivityV1HeaderParams *ApProvisionConnectivityV1HeaderParams) (*ResponseWirelessApProvisionConnectivityV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/ap-provision"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if APProvisionConnectivityV1HeaderParams != nil {

		if APProvisionConnectivityV1HeaderParams.Persistbapioutput != "" {
			clientRequest = clientRequest.SetHeader("__persistbapioutput", APProvisionConnectivityV1HeaderParams.Persistbapioutput)
		}

	}

	response, err = clientRequest.
		SetBody(requestWirelessAPProvisionConnectivityV1).
		SetResult(&ResponseWirelessApProvisionConnectivityV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ApProvisionConnectivityV1(requestWirelessAPProvisionConnectivityV1, APProvisionConnectivityV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation ApProvisionConnectivityV1")
	}

	result := response.Result().(*ResponseWirelessApProvisionConnectivityV1)
	return result, response, err

}

//CreateUpdateDynamicInterfaceV1 Create Update Dynamic interface - daa0-bb75-4e2a-8da6
/* API to create or update an dynamic interface



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-update-dynamic-interface
*/
func (s *WirelessService) CreateUpdateDynamicInterfaceV1(requestWirelessCreateUpdateDynamicInterfaceV1 *RequestWirelessCreateUpdateDynamicInterfaceV1) (*ResponseWirelessCreateUpdateDynamicInterfaceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/dynamic-interface"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessCreateUpdateDynamicInterfaceV1).
		SetResult(&ResponseWirelessCreateUpdateDynamicInterfaceV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateUpdateDynamicInterfaceV1(requestWirelessCreateUpdateDynamicInterfaceV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateUpdateDynamicInterfaceV1")
	}

	result := response.Result().(*ResponseWirelessCreateUpdateDynamicInterfaceV1)
	return result, response, err

}

//CreateWirelessProfileV1 Create Wireless Profile - 7097-6962-4bf9-88d5
/* Creates Wireless Network Profile on Cisco DNA Center and associates sites and SSIDs to it.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-wireless-profile
*/
func (s *WirelessService) CreateWirelessProfileV1(requestWirelessCreateWirelessProfileV1 *RequestWirelessCreateWirelessProfileV1) (*ResponseWirelessCreateWirelessProfileV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/profile"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessCreateWirelessProfileV1).
		SetResult(&ResponseWirelessCreateWirelessProfileV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateWirelessProfileV1(requestWirelessCreateWirelessProfileV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateWirelessProfileV1")
	}

	result := response.Result().(*ResponseWirelessCreateWirelessProfileV1)
	return result, response, err

}

//ProvisionV1 Provision - d09b-08a3-447a-a3b9
/* Provision wireless device



Documentation Link: https://developer.cisco.com/docs/dna-center/#!provision
*/
func (s *WirelessService) ProvisionV1(requestWirelessProvisionV1 *RequestWirelessProvisionV1) (*ResponseWirelessProvisionV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/provision"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessProvisionV1).
		SetResult(&ResponseWirelessProvisionV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ProvisionV1(requestWirelessProvisionV1)
		}

		return nil, response, fmt.Errorf("error with operation ProvisionV1")
	}

	result := response.Result().(*ResponseWirelessProvisionV1)
	return result, response, err

}

//PSKOverrideV1 PSK override - 46ad-ab75-47c9-8762
/* Update/Override passphrase of SSID



Documentation Link: https://developer.cisco.com/docs/dna-center/#!p-s-k-override
*/
func (s *WirelessService) PSKOverrideV1(requestWirelessPSKOverrideV1 *RequestWirelessPSKOverrideV1) (*ResponseWirelessPSKOverrideV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/psk-override"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessPSKOverrideV1).
		SetResult(&ResponseWirelessPSKOverrideV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.PSKOverrideV1(requestWirelessPSKOverrideV1)
		}

		return nil, response, fmt.Errorf("error with operation PSKOverrideV1")
	}

	result := response.Result().(*ResponseWirelessPSKOverrideV1)
	return result, response, err

}

//CreateOrUpdateRfProfileV1 Create or Update RF profile - b783-2967-4878-b815
/* Create or Update RF profile



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-or-update-rf-profile
*/
func (s *WirelessService) CreateOrUpdateRfProfileV1(requestWirelessCreateOrUpdateRFProfileV1 *RequestWirelessCreateOrUpdateRfProfileV1) (*ResponseWirelessCreateOrUpdateRfProfileV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/rf-profile"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessCreateOrUpdateRFProfileV1).
		SetResult(&ResponseWirelessCreateOrUpdateRfProfileV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateOrUpdateRfProfileV1(requestWirelessCreateOrUpdateRFProfileV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateOrUpdateRfProfileV1")
	}

	result := response.Result().(*ResponseWirelessCreateOrUpdateRfProfileV1)
	return result, response, err

}

//FactoryResetAccessPointsV1 Factory Reset Access Point(s) - b09d-4bbc-482b-aeb7
/* This API is used to factory reset Access Points. It is supported for maximum 100 Access Points per request. Factory reset clears all configurations from the Access Points. After factory reset the Access Point may become unreachable from the currently associated Wireless Controller and may or may not join back the same controller.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!factory-reset-access-points
*/
func (s *WirelessService) FactoryResetAccessPointsV1(requestWirelessFactoryResetAccessPointsV1 *RequestWirelessFactoryResetAccessPointsV1) (*ResponseWirelessFactoryResetAccessPointsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessAccessPoints/factoryResetRequest/provision"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessFactoryResetAccessPointsV1).
		SetResult(&ResponseWirelessFactoryResetAccessPointsV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.FactoryResetAccessPointsV1(requestWirelessFactoryResetAccessPointsV1)
		}

		return nil, response, fmt.Errorf("error with operation FactoryResetAccessPointsV1")
	}

	result := response.Result().(*ResponseWirelessFactoryResetAccessPointsV1)
	return result, response, err

}

//ApProvisionV1 AP Provision - 11af-897a-413b-925a
/* This API is used to provision Access Points. Prerequisite: Access Point has to be assigned to the site using the API /dna/intent/api/v1/networkDevices/assignToSite/apply.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!ap-provision
*/
func (s *WirelessService) ApProvisionV1(requestWirelessAPProvisionV1 *RequestWirelessApProvisionV1) (*ResponseWirelessApProvisionV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessAccessPoints/provision"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessAPProvisionV1).
		SetResult(&ResponseWirelessApProvisionV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ApProvisionV1(requestWirelessAPProvisionV1)
		}

		return nil, response, fmt.Errorf("error with operation ApProvisionV1")
	}

	result := response.Result().(*ResponseWirelessApProvisionV1)
	return result, response, err

}

//MobilityProvisionV1 Mobility Provision - 6c8b-6bd5-40bb-ac31
/* This API is used to provision/deploy wireless mobility into Cisco wireless controllers.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!mobility-provision
*/
func (s *WirelessService) MobilityProvisionV1(requestWirelessMobilityProvisionV1 *RequestWirelessMobilityProvisionV1) (*ResponseWirelessMobilityProvisionV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessControllers/wirelessMobilityGroups/mobilityProvision"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessMobilityProvisionV1).
		SetResult(&ResponseWirelessMobilityProvisionV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.MobilityProvisionV1(requestWirelessMobilityProvisionV1)
		}

		return nil, response, fmt.Errorf("error with operation MobilityProvisionV1")
	}

	result := response.Result().(*ResponseWirelessMobilityProvisionV1)
	return result, response, err

}

//MobilityResetV1 Mobility Reset - e589-6baf-4caa-9bbc
/* This API is used to reset wireless mobility which in turn sets mobility group name as 'default'.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!mobility-reset
*/
func (s *WirelessService) MobilityResetV1(requestWirelessMobilityResetV1 *RequestWirelessMobilityResetV1) (*ResponseWirelessMobilityResetV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessControllers/wirelessMobilityGroups/mobilityReset"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessMobilityResetV1).
		SetResult(&ResponseWirelessMobilityResetV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.MobilityResetV1(requestWirelessMobilityResetV1)
		}

		return nil, response, fmt.Errorf("error with operation MobilityResetV1")
	}

	result := response.Result().(*ResponseWirelessMobilityResetV1)
	return result, response, err

}

//AssignManagedApLocationsForWLCV1 Assign Managed AP Locations For WLC - afbd-d880-488a-83e4
/* This API allows user to assign Managed AP Locations for IOS-XE Wireless supported devices by device ID. The payload should always be a complete list. The Managed AP Locations included in the payload will be fully processed for both addition and deletion.


@param deviceID deviceId path parameter. Network Device ID. This value can be obtained by using the API call GET: /dna/intent/api/v1/network-device/ip-address/${ipAddress}


Documentation Link: https://developer.cisco.com/docs/dna-center/#!assign-managed-ap-locations-for-w-l-c
*/
func (s *WirelessService) AssignManagedApLocationsForWLCV1(deviceID string, requestWirelessAssignManagedAPLocationsForWLCV1 *RequestWirelessAssignManagedApLocationsForWLCV1) (*ResponseWirelessAssignManagedApLocationsForWLCV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessControllers/{deviceId}/assignManagedApLocations"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessAssignManagedAPLocationsForWLCV1).
		SetResult(&ResponseWirelessAssignManagedApLocationsForWLCV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AssignManagedApLocationsForWLCV1(deviceID, requestWirelessAssignManagedAPLocationsForWLCV1)
		}

		return nil, response, fmt.Errorf("error with operation AssignManagedApLocationsForWLCV1")
	}

	result := response.Result().(*ResponseWirelessAssignManagedApLocationsForWLCV1)
	return result, response, err

}

//WirelessControllerProvisionV1 Wireless Controller Provision - 9e9c-386b-4069-9e7c
/* This API is used to provision wireless controller


@param deviceID deviceId path parameter. Network Device ID. This value can be obtained by using the API call GET: /dna/intent/api/v1/network-device/ip-address/${ipAddress}


Documentation Link: https://developer.cisco.com/docs/dna-center/#!wireless-controller-provision
*/
func (s *WirelessService) WirelessControllerProvisionV1(deviceID string, requestWirelessWirelessControllerProvisionV1 *RequestWirelessWirelessControllerProvisionV1) (*ResponseWirelessWirelessControllerProvisionV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessControllers/{deviceId}/provision"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessWirelessControllerProvisionV1).
		SetResult(&ResponseWirelessWirelessControllerProvisionV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.WirelessControllerProvisionV1(deviceID, requestWirelessWirelessControllerProvisionV1)
		}

		return nil, response, fmt.Errorf("error with operation WirelessControllerProvisionV1")
	}

	result := response.Result().(*ResponseWirelessWirelessControllerProvisionV1)
	return result, response, err

}

//CreateWirelessProfileConnectivityV1 Create Wireless Profile - dd88-bb37-492a-888b
/* This API allows the user to create a Wireless Network Profile



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-wireless-profile-connectivity
*/
func (s *WirelessService) CreateWirelessProfileConnectivityV1(requestWirelessCreateWirelessProfileConnectivityV1 *RequestWirelessCreateWirelessProfileConnectivityV1) (*ResponseWirelessCreateWirelessProfileConnectivityV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessProfiles"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessCreateWirelessProfileConnectivityV1).
		SetResult(&ResponseWirelessCreateWirelessProfileConnectivityV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateWirelessProfileConnectivityV1(requestWirelessCreateWirelessProfileConnectivityV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateWirelessProfileConnectivityV1")
	}

	result := response.Result().(*ResponseWirelessCreateWirelessProfileConnectivityV1)
	return result, response, err

}

//CreateMultiplePolicyTagsForAWirelessProfileInBulkV1 Create multiple Policy Tags for a Wireless Profile in bulk - 6bbe-ca2b-430a-8665
/* This endpoint allows the creation of multiple `Policy Tags` associated with a specific `Wireless Profile` in a single request. The `id` of the Wireless Profile must be provided as a path parameter, and a list of `Policy Tags` should be included in the request body. Note: Multiple Policy Tags (policyTag) can be configured for the same siteId only if they have different sets of AP Zones (apZones). If multiple Policy Tags are created with the same apZones for the same site or a parent site, only the last one will be saved, overriding the previous ones.


@param id id path parameter. Wireless Profile Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-multiple-policy-tags-for-a-wireless-profile-in-bulk
*/
func (s *WirelessService) CreateMultiplePolicyTagsForAWirelessProfileInBulkV1(id string, requestWirelessCreateMultiplePolicyTagsForAWirelessProfileInBulkV1 *RequestWirelessCreateMultiplePolicyTagsForAWirelessProfileInBulkV1) (*ResponseWirelessCreateMultiplePolicyTagsForAWirelessProfileInBulkV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessProfiles/{id}/policyTags/bulk"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessCreateMultiplePolicyTagsForAWirelessProfileInBulkV1).
		SetResult(&ResponseWirelessCreateMultiplePolicyTagsForAWirelessProfileInBulkV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateMultiplePolicyTagsForAWirelessProfileInBulkV1(id, requestWirelessCreateMultiplePolicyTagsForAWirelessProfileInBulkV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateMultiplePolicyTagsForAWirelessProfileInBulkV1")
	}

	result := response.Result().(*ResponseWirelessCreateMultiplePolicyTagsForAWirelessProfileInBulkV1)
	return result, response, err

}

//CreateMultipleSiteTagsForAWirelessProfileInBulkV1 Create multiple Site Tags for a Wireless Profile in bulk - 7094-e87b-4b2b-9617
/* This endpoint allows the creation of multiple `Site Tags` associated with a specific `Wireless Profile` in a single request. The `id` of the `Wireless Profile` must be provided as a path parameter, and a list of `Site Tags` should be included in the request body. Note: Only one Site Tag (siteTag) can be created per siteId. If multiple siteTags are specified for the same siteId within a request, only the last one will be saved, overriding any previously configured tags. When creating a Site Tag under a Flex-enabled Wireless Profile (i.e., a Wireless Profile with one or more Flex SSIDs), a non-default Flex Profile Name (flexProfileName) will be used. If no custom flexProfileName is defined, the System will automatically generate one and configure it in the controller.


@param id id path parameter. network profile id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-multiple-site-tags-for-a-wireless-profile-in-bulk
*/
func (s *WirelessService) CreateMultipleSiteTagsForAWirelessProfileInBulkV1(id string, requestWirelessCreateMultipleSiteTagsForAWirelessProfileInBulkV1 *RequestWirelessCreateMultipleSiteTagsForAWirelessProfileInBulkV1) (*ResponseWirelessCreateMultipleSiteTagsForAWirelessProfileInBulkV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessProfiles/{id}/siteTags/bulk"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessCreateMultipleSiteTagsForAWirelessProfileInBulkV1).
		SetResult(&ResponseWirelessCreateMultipleSiteTagsForAWirelessProfileInBulkV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateMultipleSiteTagsForAWirelessProfileInBulkV1(id, requestWirelessCreateMultipleSiteTagsForAWirelessProfileInBulkV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateMultipleSiteTagsForAWirelessProfileInBulkV1")
	}

	result := response.Result().(*ResponseWirelessCreateMultipleSiteTagsForAWirelessProfileInBulkV1)
	return result, response, err

}

//CreateAnchorGroupV1 Create AnchorGroup - 3d85-68e5-4909-988a
/* This API allows the user to create an AnchorGroup



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-anchor-group
*/
func (s *WirelessService) CreateAnchorGroupV1(requestWirelessCreateAnchorGroupV1 *RequestWirelessCreateAnchorGroupV1) (*ResponseWirelessCreateAnchorGroupV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/anchorGroups"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessCreateAnchorGroupV1).
		SetResult(&ResponseWirelessCreateAnchorGroupV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateAnchorGroupV1(requestWirelessCreateAnchorGroupV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateAnchorGroupV1")
	}

	result := response.Result().(*ResponseWirelessCreateAnchorGroupV1)
	return result, response, err

}

//CreateApAuthorizationListV1 Create AP Authorization List - 5e9a-4806-489a-91db
/* This API allows the user to create an AP Authorization List.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-ap-authorization-list
*/
func (s *WirelessService) CreateApAuthorizationListV1(requestWirelessCreateAPAuthorizationListV1 *RequestWirelessCreateApAuthorizationListV1) (*ResponseWirelessCreateApAuthorizationListV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/apAuthorizationLists"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessCreateAPAuthorizationListV1).
		SetResult(&ResponseWirelessCreateApAuthorizationListV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateApAuthorizationListV1(requestWirelessCreateAPAuthorizationListV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateApAuthorizationListV1")
	}

	result := response.Result().(*ResponseWirelessCreateApAuthorizationListV1)
	return result, response, err

}

//CreateApProfileV1 Create AP Profile - 3697-68d5-4149-9f02
/* This API allows the user to create a custom AP Profile.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-ap-profile
*/
func (s *WirelessService) CreateApProfileV1(requestWirelessCreateAPProfileV1 *RequestWirelessCreateApProfileV1) (*ResponseWirelessCreateApProfileV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/apProfiles"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessCreateAPProfileV1).
		SetResult(&ResponseWirelessCreateApProfileV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateApProfileV1(requestWirelessCreateAPProfileV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateApProfileV1")
	}

	result := response.Result().(*ResponseWirelessCreateApProfileV1)
	return result, response, err

}

//CreateA80211BeProfileV1 Create a 802.11be Profile - efab-bbaf-4388-a046
/* This API allows the user to create a 802.11be Profile.Catalyst Center will push this profile to device's "default-dot11be-profile”.Also please note , 802.11be Profile is supported only on IOS-XE controllers since device version 17.15



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-a80211be-profile
*/
func (s *WirelessService) CreateA80211BeProfileV1(requestWirelessCreateA80211beProfileV1 *RequestWirelessCreateA80211BeProfileV1) (*ResponseWirelessCreateA80211BeProfileV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/dot11beProfiles"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessCreateA80211beProfileV1).
		SetResult(&ResponseWirelessCreateA80211BeProfileV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateA80211BeProfileV1(requestWirelessCreateA80211beProfileV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateA80211BeProfileV1")
	}

	result := response.Result().(*ResponseWirelessCreateA80211BeProfileV1)
	return result, response, err

}

//CreateInterfaceV1 Create Interface - a098-6877-44e8-ba31
/* This API allows the user to create an interface



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-interface
*/
func (s *WirelessService) CreateInterfaceV1(requestWirelessCreateInterfaceV1 *RequestWirelessCreateInterfaceV1) (*ResponseWirelessCreateInterfaceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/interfaces"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessCreateInterfaceV1).
		SetResult(&ResponseWirelessCreateInterfaceV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateInterfaceV1(requestWirelessCreateInterfaceV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateInterfaceV1")
	}

	result := response.Result().(*ResponseWirelessCreateInterfaceV1)
	return result, response, err

}

//CreatePowerProfileV1 Create Power Profile - 7bac-6bd5-4269-8a3c
/* This API allows the user to create a custom Power Profile.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-power-profile
*/
func (s *WirelessService) CreatePowerProfileV1(requestWirelessCreatePowerProfileV1 *RequestWirelessCreatePowerProfileV1) (*ResponseWirelessCreatePowerProfileV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/powerProfiles"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessCreatePowerProfileV1).
		SetResult(&ResponseWirelessCreatePowerProfileV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreatePowerProfileV1(requestWirelessCreatePowerProfileV1)
		}

		return nil, response, fmt.Errorf("error with operation CreatePowerProfileV1")
	}

	result := response.Result().(*ResponseWirelessCreatePowerProfileV1)
	return result, response, err

}

//CreateRfProfileV1 Create RF Profile - 3cb0-ca20-45d9-8d07
/* This API allows the user to create a custom RF Profile



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-rf-profile
*/
func (s *WirelessService) CreateRfProfileV1(requestWirelessCreateRFProfileV1 *RequestWirelessCreateRfProfileV1) (*ResponseWirelessCreateRfProfileV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/rfProfiles"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessCreateRFProfileV1).
		SetResult(&ResponseWirelessCreateRfProfileV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateRfProfileV1(requestWirelessCreateRFProfileV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateRfProfileV1")
	}

	result := response.Result().(*ResponseWirelessCreateRfProfileV1)
	return result, response, err

}

//AssignAnchorManagedApLocationsForWLCV1 Assign Anchor Managed AP Locations For WLC - 55af-697b-4e28-8167
/* This API allows user to assign Anchor Managed AP Locations for WLC by device ID. The payload should always be a complete list. The Managed AP Locations included in the payload will be fully processed for both addition and deletion.

       When anchor managed location array present then it will add the anchor managed locations.


@param networkDeviceID networkDeviceId path parameter. Network Device ID. This value can be obtained by using the API call GET: /dna/intent/api/v1/network-device/ip-address/${ipAddress}


Documentation Link: https://developer.cisco.com/docs/dna-center/#!assign-anchor-managed-ap-locations-for-w-l-c
*/
func (s *WirelessService) AssignAnchorManagedApLocationsForWLCV1(networkDeviceID string, requestWirelessAssignAnchorManagedAPLocationsForWLCV1 *RequestWirelessAssignAnchorManagedApLocationsForWLCV1) (*ResponseWirelessAssignAnchorManagedApLocationsForWLCV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/{networkDeviceId}/assignAnchorManagedApLocations"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessAssignAnchorManagedAPLocationsForWLCV1).
		SetResult(&ResponseWirelessAssignAnchorManagedApLocationsForWLCV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AssignAnchorManagedApLocationsForWLCV1(networkDeviceID, requestWirelessAssignAnchorManagedAPLocationsForWLCV1)
		}

		return nil, response, fmt.Errorf("error with operation AssignAnchorManagedApLocationsForWLCV1")
	}

	result := response.Result().(*ResponseWirelessAssignAnchorManagedApLocationsForWLCV1)
	return result, response, err

}

//ConfigureAccessPointsV2 Configure Access Points V2 - 5ca7-4a81-4329-9506
/* User can configure multiple access points with required options using this intent API



Documentation Link: https://developer.cisco.com/docs/dna-center/#!configure-access-points-v2
*/
func (s *WirelessService) ConfigureAccessPointsV2(requestWirelessConfigureAccessPointsV2 *RequestWirelessConfigureAccessPointsV2) (*ResponseWirelessConfigureAccessPointsV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/wireless/accesspoint-configuration"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessConfigureAccessPointsV2).
		SetResult(&ResponseWirelessConfigureAccessPointsV2{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ConfigureAccessPointsV2(requestWirelessConfigureAccessPointsV2)
		}

		return nil, response, fmt.Errorf("error with operation ConfigureAccessPointsV2")
	}

	result := response.Result().(*ResponseWirelessConfigureAccessPointsV2)
	return result, response, err

}

//UpdateEnterpriseSSIDV1 Update Enterprise SSID - c493-991f-40ca-ba44
/* Update enterprise SSID


 */
func (s *WirelessService) UpdateEnterpriseSSIDV1(requestWirelessUpdateEnterpriseSSIDV1 *RequestWirelessUpdateEnterpriseSSIDV1) (*ResponseWirelessUpdateEnterpriseSSIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/enterprise-ssid"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateEnterpriseSSIDV1).
		SetResult(&ResponseWirelessUpdateEnterpriseSSIDV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateEnterpriseSSIDV1(requestWirelessUpdateEnterpriseSSIDV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateEnterpriseSsidV1")
	}

	result := response.Result().(*ResponseWirelessUpdateEnterpriseSSIDV1)
	return result, response, err

}

//UpdateSSIDV1 Update SSID - 2496-7ad2-4b8a-913b
/* This API allows the user to update an SSID (Service Set Identifier) at the given site


@param siteID siteId path parameter. Site UUID

@param id id path parameter. SSID ID

*/
func (s *WirelessService) UpdateSSIDV1(siteID string, id string, requestWirelessUpdateSSIDV1 *RequestWirelessUpdateSSIDV1) (*ResponseWirelessUpdateSSIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{siteId}/wirelessSettings/ssids/{id}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateSSIDV1).
		SetResult(&ResponseWirelessUpdateSSIDV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateSSIDV1(siteID, id, requestWirelessUpdateSSIDV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateSsidV1")
	}

	result := response.Result().(*ResponseWirelessUpdateSSIDV1)
	return result, response, err

}

//UpdateWirelessProfileV1 Update Wireless Profile - cfbd-3870-405a-ad55
/* Updates the wireless Network Profile with updated details provided. All sites to be present in the network profile should be provided.


 */
func (s *WirelessService) UpdateWirelessProfileV1(requestWirelessUpdateWirelessProfileV1 *RequestWirelessUpdateWirelessProfileV1) (*ResponseWirelessUpdateWirelessProfileV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/profile"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateWirelessProfileV1).
		SetResult(&ResponseWirelessUpdateWirelessProfileV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateWirelessProfileV1(requestWirelessUpdateWirelessProfileV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateWirelessProfileV1")
	}

	result := response.Result().(*ResponseWirelessUpdateWirelessProfileV1)
	return result, response, err

}

//ProvisionUpdateV1 Provision update - 87a5-ab04-4139-862d
/* Updates wireless provisioning


@param ProvisionUpdateV1HeaderParams Custom header parameters
*/
func (s *WirelessService) ProvisionUpdateV1(requestWirelessProvisionUpdateV1 *RequestWirelessProvisionUpdateV1, ProvisionUpdateV1HeaderParams *ProvisionUpdateV1HeaderParams) (*ResponseWirelessProvisionUpdateV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/provision"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if ProvisionUpdateV1HeaderParams != nil {

		if ProvisionUpdateV1HeaderParams.Persistbapioutput != "" {
			clientRequest = clientRequest.SetHeader("__persistbapioutput", ProvisionUpdateV1HeaderParams.Persistbapioutput)
		}

	}

	response, err = clientRequest.
		SetBody(requestWirelessProvisionUpdateV1).
		SetResult(&ResponseWirelessProvisionUpdateV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ProvisionUpdateV1(requestWirelessProvisionUpdateV1, ProvisionUpdateV1HeaderParams)
		}
		return nil, response, fmt.Errorf("error with operation ProvisionUpdateV1")
	}

	result := response.Result().(*ResponseWirelessProvisionUpdateV1)
	return result, response, err

}

//UpdateWirelessProfileConnectivityV1 Update Wireless Profile - 4f88-d9a3-4ef8-8e2e
/* This API allows the user to update a Wireless Network Profile by ID


@param id id path parameter. Wireless Profile Id

*/
func (s *WirelessService) UpdateWirelessProfileConnectivityV1(id string, requestWirelessUpdateWirelessProfileConnectivityV1 *RequestWirelessUpdateWirelessProfileConnectivityV1) (*ResponseWirelessUpdateWirelessProfileConnectivityV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessProfiles/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateWirelessProfileConnectivityV1).
		SetResult(&ResponseWirelessUpdateWirelessProfileConnectivityV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateWirelessProfileConnectivityV1(id, requestWirelessUpdateWirelessProfileConnectivityV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateWirelessProfileConnectivityV1")
	}

	result := response.Result().(*ResponseWirelessUpdateWirelessProfileConnectivityV1)
	return result, response, err

}

//UpdateASpecificPolicyTagForAWirelessProfileV1 Update a specific Policy Tag for a Wireless Profile - eca1-caf0-49bb-aa9a
/* This endpoint allows updating the details of a specific `Policy Tag` associated with a given `Wireless Profile`. The `id` of the `Wireless Profile` and the `policyTagId` of the Policy Tag must be provided as path parameters, and the request body should contain the updated details of the `Policy Tag`. The `policyTagName` cannot be modified through this endpoint. Note: When updating a Policy Tag, if the same set of AP Zones (apZones) is used for the same site or its parent site, the existing Policy Tag will be overridden by the new one.


@param id id path parameter. Wireless Profile Id

@param policyTagID policyTagId path parameter. Policy Tag Id

*/
func (s *WirelessService) UpdateASpecificPolicyTagForAWirelessProfileV1(id string, policyTagID string, requestWirelessUpdateASpecificPolicyTagForAWirelessProfileV1 *RequestWirelessUpdateASpecificPolicyTagForAWirelessProfileV1) (*ResponseWirelessUpdateASpecificPolicyTagForAWirelessProfileV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessProfiles/{id}/policyTags/{policyTagId}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{policyTagId}", fmt.Sprintf("%v", policyTagID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateASpecificPolicyTagForAWirelessProfileV1).
		SetResult(&ResponseWirelessUpdateASpecificPolicyTagForAWirelessProfileV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateASpecificPolicyTagForAWirelessProfileV1(id, policyTagID, requestWirelessUpdateASpecificPolicyTagForAWirelessProfileV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateASpecificPolicyTagForAWirelessProfileV1")
	}

	result := response.Result().(*ResponseWirelessUpdateASpecificPolicyTagForAWirelessProfileV1)
	return result, response, err

}

//UpdateASpecificSiteTagForAWirelessProfileV1 Update a specific Site Tag for a Wireless Profile - 47b2-1a6c-4158-9527
/* This endpoint allows updating the details of a specific `Site Tag` associated with a given `Wireless Profile`. The `id` of the `Wireless Profile` and the `siteTagId` of the Site Tag must be provided as path parameters, and the request body should contain the updated `Site Tag` details.  The `siteTagName` cannot be modified through this endpoint. Note: When updating a Site Tag (siteTag), if the siteId already has an associated siteTag and the same siteId is included in the update request, the existing siteTag for that siteId will be overridden by the new one. For Flex-enabled Wireless Profiles (i.e., a Wireless Profile with one or more Flex SSIDs), a non-default Flex Profile Name (flexProfileName) will be used. If no custom flexProfileName is provided, the System will automatically generate one and configure it in the controller.


@param id id path parameter. Wireless Profile Id

@param siteTagID siteTagId path parameter. Site Tag Id

*/
func (s *WirelessService) UpdateASpecificSiteTagForAWirelessProfileV1(id string, siteTagID string, requestWirelessUpdateASpecificSiteTagForAWirelessProfileV1 *RequestWirelessUpdateASpecificSiteTagForAWirelessProfileV1) (*ResponseWirelessUpdateASpecificSiteTagForAWirelessProfileV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessProfiles/{id}/siteTags/{siteTagId}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{siteTagId}", fmt.Sprintf("%v", siteTagID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateASpecificSiteTagForAWirelessProfileV1).
		SetResult(&ResponseWirelessUpdateASpecificSiteTagForAWirelessProfileV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateASpecificSiteTagForAWirelessProfileV1(id, siteTagID, requestWirelessUpdateASpecificSiteTagForAWirelessProfileV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateASpecificSiteTagForAWirelessProfileV1")
	}

	result := response.Result().(*ResponseWirelessUpdateASpecificSiteTagForAWirelessProfileV1)
	return result, response, err

}

//UpdateAnchorGroupV1 Update AnchorGroup - 4da3-bae4-484a-8448
/* This API allows the user to update an AnchorGroup


@param id id path parameter. AnchorGroup ID

*/
func (s *WirelessService) UpdateAnchorGroupV1(id string, requestWirelessUpdateAnchorGroupV1 *RequestWirelessUpdateAnchorGroupV1) (*ResponseWirelessUpdateAnchorGroupV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/anchorGroups/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateAnchorGroupV1).
		SetResult(&ResponseWirelessUpdateAnchorGroupV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateAnchorGroupV1(id, requestWirelessUpdateAnchorGroupV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateAnchorGroupV1")
	}

	result := response.Result().(*ResponseWirelessUpdateAnchorGroupV1)
	return result, response, err

}

//UpdateApAuthorizationListV1 Update AP Authorization List - 768b-0b99-4189-986e
/* This API allows the user to update an AP Authorization List.


@param id id path parameter. AP Authorization List ID

*/
func (s *WirelessService) UpdateApAuthorizationListV1(id string, requestWirelessUpdateAPAuthorizationListV1 *RequestWirelessUpdateApAuthorizationListV1) (*ResponseWirelessUpdateApAuthorizationListV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/apAuthorizationLists/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateAPAuthorizationListV1).
		SetResult(&ResponseWirelessUpdateApAuthorizationListV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateApAuthorizationListV1(id, requestWirelessUpdateAPAuthorizationListV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateApAuthorizationListV1")
	}

	result := response.Result().(*ResponseWirelessUpdateApAuthorizationListV1)
	return result, response, err

}

//UpdateApProfileByIDV1 Update AP Profile by ID - 41b5-ea3b-43b8-9d57
/* This API allows the user to update a custom AP Profile


@param id id path parameter. Ap Profile ID

*/
func (s *WirelessService) UpdateApProfileByIDV1(id string, requestWirelessUpdateAPProfileByIDV1 *RequestWirelessUpdateApProfileByIDV1) (*ResponseWirelessUpdateApProfileByIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/apProfiles/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateAPProfileByIDV1).
		SetResult(&ResponseWirelessUpdateApProfileByIDV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateApProfileByIDV1(id, requestWirelessUpdateAPProfileByIDV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateApProfileByIdV1")
	}

	result := response.Result().(*ResponseWirelessUpdateApProfileByIDV1)
	return result, response, err

}

//Update80211BeProfileV1 Update 802.11be Profile - 699b-b8e0-48bb-9b90
/* This API allows the user to update a 802.11be Profile


@param id id path parameter. 802.11be Profile ID

*/
func (s *WirelessService) Update80211BeProfileV1(id string, requestWirelessUpdate80211beProfileV1 *RequestWirelessUpdate80211BeProfileV1) (*ResponseWirelessUpdate80211BeProfileV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/dot11beProfiles/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdate80211beProfileV1).
		SetResult(&ResponseWirelessUpdate80211BeProfileV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.Update80211BeProfileV1(id, requestWirelessUpdate80211beProfileV1)
		}
		return nil, response, fmt.Errorf("error with operation Update80211BeProfileV1")
	}

	result := response.Result().(*ResponseWirelessUpdate80211BeProfileV1)
	return result, response, err

}

//UpdateInterfaceV1 Update Interface - 0f93-9868-454b-a943
/* This API allows the user to update an interface by ID


@param id id path parameter. Interface ID

*/
func (s *WirelessService) UpdateInterfaceV1(id string, requestWirelessUpdateInterfaceV1 *RequestWirelessUpdateInterfaceV1) (*ResponseWirelessUpdateInterfaceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/interfaces/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateInterfaceV1).
		SetResult(&ResponseWirelessUpdateInterfaceV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateInterfaceV1(id, requestWirelessUpdateInterfaceV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateInterfaceV1")
	}

	result := response.Result().(*ResponseWirelessUpdateInterfaceV1)
	return result, response, err

}

//UpdatePowerProfileByIDV1 Update Power Profile by ID - 6f86-19b1-478a-9a5e
/* This API allows the user to update a custom power Profile


@param id id path parameter. Power Profile Id

*/
func (s *WirelessService) UpdatePowerProfileByIDV1(id string, requestWirelessUpdatePowerProfileByIDV1 *RequestWirelessUpdatePowerProfileByIDV1) (*ResponseWirelessUpdatePowerProfileByIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/powerProfiles/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdatePowerProfileByIDV1).
		SetResult(&ResponseWirelessUpdatePowerProfileByIDV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdatePowerProfileByIDV1(id, requestWirelessUpdatePowerProfileByIDV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdatePowerProfileByIdV1")
	}

	result := response.Result().(*ResponseWirelessUpdatePowerProfileByIDV1)
	return result, response, err

}

//UpdateRfProfileV1 Update RF Profile - 2984-b995-4ae9-b3c3
/* This API allows the user to update a custom RF Profile.


@param id id path parameter. RF Profile ID

*/
func (s *WirelessService) UpdateRfProfileV1(id string, requestWirelessUpdateRFProfileV1 *RequestWirelessUpdateRfProfileV1) (*ResponseWirelessUpdateRfProfileV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/rfProfiles/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateRFProfileV1).
		SetResult(&ResponseWirelessUpdateRfProfileV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateRfProfileV1(id, requestWirelessUpdateRFProfileV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateRfProfileV1")
	}

	result := response.Result().(*ResponseWirelessUpdateRfProfileV1)
	return result, response, err

}

//DeleteSSIDAndProvisionItToDevicesV1 Delete SSID and provision it to devices - fc95-38fe-43d9-884d
/* Removes SSID or WLAN from the network profile, reprovision the device(s) and deletes the SSID or WLAN from DNA Center


@param ssidName ssidName path parameter. SSID Name. This parameter needs to be encoded as per UTF-8 encoding.

@param managedAPLocations managedAPLocations path parameter. List of managed AP locations (Site Hierarchies). This parameter needs to be encoded as per UTF-8 encoding

@param DeleteSSIDAndProvisionItToDevicesV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-ssid-and-provision-it-to-devices
*/
func (s *WirelessService) DeleteSSIDAndProvisionItToDevicesV1(ssidName string, managedAPLocations string, DeleteSSIDAndProvisionItToDevicesV1HeaderParams *DeleteSSIDAndProvisionItToDevicesV1HeaderParams) (*ResponseWirelessDeleteSSIDAndProvisionItToDevicesV1, *resty.Response, error) {
	//ssidName string,managedAPLocations string,DeleteSSIDAndProvisionItToDevicesV1HeaderParams *DeleteSSIDAndProvisionItToDevicesV1HeaderParams
	path := "/dna/intent/api/v1/business/ssid/{ssidName}/{managedAPLocations}"
	path = strings.Replace(path, "{ssidName}", fmt.Sprintf("%v", ssidName), -1)
	path = strings.Replace(path, "{managedAPLocations}", fmt.Sprintf("%v", managedAPLocations), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if DeleteSSIDAndProvisionItToDevicesV1HeaderParams != nil {

		if DeleteSSIDAndProvisionItToDevicesV1HeaderParams.Persistbapioutput != "" {
			clientRequest = clientRequest.SetHeader("__persistbapioutput", DeleteSSIDAndProvisionItToDevicesV1HeaderParams.Persistbapioutput)
		}

	}

	response, err = clientRequest.
		SetResult(&ResponseWirelessDeleteSSIDAndProvisionItToDevicesV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteSSIDAndProvisionItToDevicesV1(ssidName, managedAPLocations, DeleteSSIDAndProvisionItToDevicesV1HeaderParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteSsidAndProvisionItToDevicesV1")
	}

	result := response.Result().(*ResponseWirelessDeleteSSIDAndProvisionItToDevicesV1)
	return result, response, err

}

//DeleteEnterpriseSSIDV1 Delete Enterprise SSID - c7a6-592b-4b98-a369
/* Deletes given enterprise SSID


@param ssidName ssidName path parameter. Enter the SSID name to be deleted


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-enterprise-ssid
*/
func (s *WirelessService) DeleteEnterpriseSSIDV1(ssidName string) (*ResponseWirelessDeleteEnterpriseSSIDV1, *resty.Response, error) {
	//ssidName string
	path := "/dna/intent/api/v1/enterprise-ssid/{ssidName}"
	path = strings.Replace(path, "{ssidName}", fmt.Sprintf("%v", ssidName), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessDeleteEnterpriseSSIDV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteEnterpriseSSIDV1(ssidName)
		}
		return nil, response, fmt.Errorf("error with operation DeleteEnterpriseSsidV1")
	}

	result := response.Result().(*ResponseWirelessDeleteEnterpriseSSIDV1)
	return result, response, err

}

//DeleteSSIDV1 Delete SSID - acbe-8b6f-4e8b-9f6a
/* This API allows the user to delete an SSID (Service Set Identifier) at the global level , if the SSID is not mapped to any Wireless Profile, Or remove override from given site Id .


@param siteID siteId path parameter. Site UUID where SSID is to be deleted

@param id id path parameter. SSID ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-ssid
*/
func (s *WirelessService) DeleteSSIDV1(siteID string, id string) (*ResponseWirelessDeleteSSIDV1, *resty.Response, error) {
	//siteID string,id string
	path := "/dna/intent/api/v1/sites/{siteId}/wirelessSettings/ssids/{id}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessDeleteSSIDV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteSSIDV1(siteID, id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteSsidV1")
	}

	result := response.Result().(*ResponseWirelessDeleteSSIDV1)
	return result, response, err

}

//DeleteWirelessProfileV1 Delete Wireless Profile - e395-88a5-4949-82c4
/* Delete the Wireless Profile whose name is provided.


@param wirelessProfileName wirelessProfileName path parameter. Wireless Profile Name


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-wireless-profile
*/
func (s *WirelessService) DeleteWirelessProfileV1(wirelessProfileName string) (*ResponseWirelessDeleteWirelessProfileV1, *resty.Response, error) {
	//wirelessProfileName string
	path := "/dna/intent/api/v1/wireless-profile/{wirelessProfileName}"
	path = strings.Replace(path, "{wirelessProfileName}", fmt.Sprintf("%v", wirelessProfileName), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessDeleteWirelessProfileV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteWirelessProfileV1(wirelessProfileName)
		}
		return nil, response, fmt.Errorf("error with operation DeleteWirelessProfileV1")
	}

	result := response.Result().(*ResponseWirelessDeleteWirelessProfileV1)
	return result, response, err

}

//DeleteDynamicInterfaceV1 Delete dynamic interface - ffb4-abf4-44fb-b70a
/* Delete a dynamic interface


@param DeleteDynamicInterfaceV1HeaderParams Custom header parameters
@param DeleteDynamicInterfaceV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-dynamic-interface
*/
func (s *WirelessService) DeleteDynamicInterfaceV1(DeleteDynamicInterfaceV1HeaderParams *DeleteDynamicInterfaceV1HeaderParams, DeleteDynamicInterfaceV1QueryParams *DeleteDynamicInterfaceV1QueryParams) (*ResponseWirelessDeleteDynamicInterfaceV1, *resty.Response, error) {
	//DeleteDynamicInterfaceV1HeaderParams *DeleteDynamicInterfaceV1HeaderParams,DeleteDynamicInterfaceV1QueryParams *DeleteDynamicInterfaceV1QueryParams
	path := "/dna/intent/api/v1/wireless/dynamic-interface"

	queryString, _ := query.Values(DeleteDynamicInterfaceV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if DeleteDynamicInterfaceV1HeaderParams != nil {

		if DeleteDynamicInterfaceV1HeaderParams.Runsync != "" {
			clientRequest = clientRequest.SetHeader("__runsync", DeleteDynamicInterfaceV1HeaderParams.Runsync)
		}

		if DeleteDynamicInterfaceV1HeaderParams.Timeout != "" {
			clientRequest = clientRequest.SetHeader("__timeout", DeleteDynamicInterfaceV1HeaderParams.Timeout)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessDeleteDynamicInterfaceV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteDynamicInterfaceV1(DeleteDynamicInterfaceV1HeaderParams, DeleteDynamicInterfaceV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteDynamicInterfaceV1")
	}

	result := response.Result().(*ResponseWirelessDeleteDynamicInterfaceV1)
	return result, response, err

}

//DeleteRfProfilesV1 Delete RF profiles - 28b2-4a74-4a99-94be
/* Delete RF profile


@param rfProfileName rfProfileName path parameter. RF profile name to be deleted(required) *non-custom RF profile cannot be deleted


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-rf-profiles
*/
func (s *WirelessService) DeleteRfProfilesV1(rfProfileName string) (*ResponseWirelessDeleteRfProfilesV1, *resty.Response, error) {
	//rfProfileName string
	path := "/dna/intent/api/v1/wireless/rf-profile/{rfProfileName}"
	path = strings.Replace(path, "{rfProfileName}", fmt.Sprintf("%v", rfProfileName), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessDeleteRfProfilesV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteRfProfilesV1(rfProfileName)
		}
		return nil, response, fmt.Errorf("error with operation DeleteRfProfilesV1")
	}

	result := response.Result().(*ResponseWirelessDeleteRfProfilesV1)
	return result, response, err

}

//DeleteWirelessProfileConnectivityV1 Delete Wireless Profile - 289c-f9f5-4f78-b84c
/* This API allows the user to delete Wireless Network Profile by ID


@param id id path parameter. Wireless Profile Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-wireless-profile-connectivity
*/
func (s *WirelessService) DeleteWirelessProfileConnectivityV1(id string) (*ResponseWirelessDeleteWirelessProfileConnectivityV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/wirelessProfiles/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessDeleteWirelessProfileConnectivityV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteWirelessProfileConnectivityV1(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteWirelessProfileConnectivityV1")
	}

	result := response.Result().(*ResponseWirelessDeleteWirelessProfileConnectivityV1)
	return result, response, err

}

//DeleteASpecificPolicyTagFromAWirelessProfileV1 Delete a specific Policy Tag from a Wireless Profile - 4da4-fa50-4b89-a098
/* This endpoint allows for the deletion of a specific `Policy Tag` associated with a given `Wireless Profile`. This API requires the `id` of the `Wireless Profile` and the `policyTagId` of the `Policy Tag` to be provided as path parameters.


@param id id path parameter. Wireless Profile Id

@param policyTagID policyTagId path parameter. Policy Tag Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-a-specific-policy-tag-from-a-wireless-profile
*/
func (s *WirelessService) DeleteASpecificPolicyTagFromAWirelessProfileV1(id string, policyTagID string) (*ResponseWirelessDeleteASpecificPolicyTagFromAWirelessProfileV1, *resty.Response, error) {
	//id string,policyTagID string
	path := "/dna/intent/api/v1/wirelessProfiles/{id}/policyTags/{policyTagId}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{policyTagId}", fmt.Sprintf("%v", policyTagID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessDeleteASpecificPolicyTagFromAWirelessProfileV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteASpecificPolicyTagFromAWirelessProfileV1(id, policyTagID)
		}
		return nil, response, fmt.Errorf("error with operation DeleteASpecificPolicyTagFromAWirelessProfileV1")
	}

	result := response.Result().(*ResponseWirelessDeleteASpecificPolicyTagFromAWirelessProfileV1)
	return result, response, err

}

//DeleteASpecificSiteTagFromAWirelessProfileV1 Delete a specific Site Tag from a Wireless Profile - e090-8a9d-4e2b-858f
/* This endpoint enables the deletion of a specific `Site Tag` associated with a given `Wireless Profile`. This API requires the `id` of the `Wireless Profile` and the `siteTagId` of the `Site Tag` to be provided as path parameters.


@param id id path parameter. Wireless Profile id

@param siteTagID siteTagId path parameter. Site Tag Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-a-specific-site-tag-from-a-wireless-profile
*/
func (s *WirelessService) DeleteASpecificSiteTagFromAWirelessProfileV1(id string, siteTagID string) (*ResponseWirelessDeleteASpecificSiteTagFromAWirelessProfileV1, *resty.Response, error) {
	//id string,siteTagID string
	path := "/dna/intent/api/v1/wirelessProfiles/{id}/siteTags/{siteTagId}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{siteTagId}", fmt.Sprintf("%v", siteTagID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessDeleteASpecificSiteTagFromAWirelessProfileV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteASpecificSiteTagFromAWirelessProfileV1(id, siteTagID)
		}
		return nil, response, fmt.Errorf("error with operation DeleteASpecificSiteTagFromAWirelessProfileV1")
	}

	result := response.Result().(*ResponseWirelessDeleteASpecificSiteTagFromAWirelessProfileV1)
	return result, response, err

}

//DeleteAnchorGroupByIDV1 Delete AnchorGroup by ID - 6d83-eb95-46eb-9276
/* This API allows the user to delete an AnchorGroup  by specifying the AnchorGroup ID


@param id id path parameter. AnchorGroup ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-anchor-group-by-id
*/
func (s *WirelessService) DeleteAnchorGroupByIDV1(id string) (*ResponseWirelessDeleteAnchorGroupByIDV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/wirelessSettings/anchorGroups/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessDeleteAnchorGroupByIDV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteAnchorGroupByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteAnchorGroupByIdV1")
	}

	result := response.Result().(*ResponseWirelessDeleteAnchorGroupByIDV1)
	return result, response, err

}

//DeleteApAuthorizationListV1 Delete AP Authorization List - 96a5-e901-47a8-8a31
/* This API allows the user to delete an AP Authorization List.


@param id id path parameter. AP Authorization List ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-ap-authorization-list
*/
func (s *WirelessService) DeleteApAuthorizationListV1(id string) (*ResponseWirelessDeleteApAuthorizationListV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/wirelessSettings/apAuthorizationLists/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessDeleteApAuthorizationListV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteApAuthorizationListV1(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteApAuthorizationListV1")
	}

	result := response.Result().(*ResponseWirelessDeleteApAuthorizationListV1)
	return result, response, err

}

//DeleteApProfileByIDV1 Delete AP Profile by ID - 5b95-78eb-452b-856b
/* This API allows the user to delete an AP Profile by specifying the AP Profile ID.


@param id id path parameter. AP Profile ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-ap-profile-by-id
*/
func (s *WirelessService) DeleteApProfileByIDV1(id string) (*ResponseWirelessDeleteApProfileByIDV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/wirelessSettings/apProfiles/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessDeleteApProfileByIDV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteApProfileByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteApProfileByIdV1")
	}

	result := response.Result().(*ResponseWirelessDeleteApProfileByIDV1)
	return result, response, err

}

//DeleteA80211BeProfileV1 Delete a 802.11be Profile - e9b0-98c2-4b49-8fe6
/* This API allows the user to delete a 802.11be Profile,if the 802.11be Profile is not mapped to any Wireless Network Profile


@param id id path parameter. 802.11be Profile ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-a80211be-profile
*/
func (s *WirelessService) DeleteA80211BeProfileV1(id string) (*ResponseWirelessDeleteA80211BeProfileV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/wirelessSettings/dot11beProfiles/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessDeleteA80211BeProfileV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteA80211BeProfileV1(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteA80211BeProfileV1")
	}

	result := response.Result().(*ResponseWirelessDeleteA80211BeProfileV1)
	return result, response, err

}

//DeleteInterfaceV1 Delete Interface - 0999-c9cd-4159-a6a1
/* This API allows the user to delete an interface by ID


@param id id path parameter. Interface ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-interface
*/
func (s *WirelessService) DeleteInterfaceV1(id string) (*ResponseWirelessDeleteInterfaceV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/wirelessSettings/interfaces/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessDeleteInterfaceV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteInterfaceV1(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteInterfaceV1")
	}

	result := response.Result().(*ResponseWirelessDeleteInterfaceV1)
	return result, response, err

}

//DeletePowerProfileByIDV1 Delete Power Profile by ID - 0eb7-faa1-41a9-9490
/* This API allows the user to delete an Power Profile by specifying the Power Profile ID.


@param id id path parameter. Power Profile ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-power-profile-by-id
*/
func (s *WirelessService) DeletePowerProfileByIDV1(id string) (*ResponseWirelessDeletePowerProfileByIDV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/wirelessSettings/powerProfiles/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessDeletePowerProfileByIDV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletePowerProfileByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation DeletePowerProfileByIdV1")
	}

	result := response.Result().(*ResponseWirelessDeletePowerProfileByIDV1)
	return result, response, err

}

//DeleteRfProfileV1 Delete RF Profile - 2f8a-799d-4fa9-ac0e
/* This API allows the user to delete a custom RF Profile


@param id id path parameter. RF Profile ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-rf-profile
*/
func (s *WirelessService) DeleteRfProfileV1(id string) (*ResponseWirelessDeleteRfProfileV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/wirelessSettings/rfProfiles/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessDeleteRfProfileV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteRfProfileV1(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteRfProfileV1")
	}

	result := response.Result().(*ResponseWirelessDeleteRfProfileV1)
	return result, response, err

}

// Alias Function
/*
This method acts as an alias for the method `AssignManagedApLocationsForWLCV1`
*/
func (s *WirelessService) AssignManagedApLocationsForWLC(deviceID string, requestWirelessAssignManagedAPLocationsForWLCV1 *RequestWirelessAssignManagedApLocationsForWLCV1) (*ResponseWirelessAssignManagedApLocationsForWLCV1, *resty.Response, error) {
	return s.AssignManagedApLocationsForWLCV1(deviceID, requestWirelessAssignManagedAPLocationsForWLCV1)
}

// Alias Function
/*
This method acts as an alias for the method `CreateRfProfileV1`
*/
func (s *WirelessService) CreateRfProfile(requestWirelessCreateRFProfileV1 *RequestWirelessCreateRfProfileV1) (*ResponseWirelessCreateRfProfileV1, *resty.Response, error) {
	return s.CreateRfProfileV1(requestWirelessCreateRFProfileV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetInterfacesCountV1`
*/
func (s *WirelessService) GetInterfacesCount() (*ResponseWirelessGetInterfacesCountV1, *resty.Response, error) {
	return s.GetInterfacesCountV1()
}

// Alias Function
/*
This method acts as an alias for the method `CreateMultipleSiteTagsForAWirelessProfileInBulkV1`
*/
func (s *WirelessService) CreateMultipleSiteTagsForAWirelessProfileInBulk(id string, requestWirelessCreateMultipleSiteTagsForAWirelessProfileInBulkV1 *RequestWirelessCreateMultipleSiteTagsForAWirelessProfileInBulkV1) (*ResponseWirelessCreateMultipleSiteTagsForAWirelessProfileInBulkV1, *resty.Response, error) {
	return s.CreateMultipleSiteTagsForAWirelessProfileInBulkV1(id, requestWirelessCreateMultipleSiteTagsForAWirelessProfileInBulkV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetAnchorCapableDevicesV1`
*/
func (s *WirelessService) GetAnchorCapableDevices() (*ResponseWirelessGetAnchorCapableDevicesV1, *resty.Response, error) {
	return s.GetAnchorCapableDevicesV1()
}

// Alias Function
/*
This method acts as an alias for the method `DeleteDynamicInterfaceV1`
*/
func (s *WirelessService) DeleteDynamicInterface(DeleteDynamicInterfaceV1HeaderParams *DeleteDynamicInterfaceV1HeaderParams, DeleteDynamicInterfaceV1QueryParams *DeleteDynamicInterfaceV1QueryParams) (*ResponseWirelessDeleteDynamicInterfaceV1, *resty.Response, error) {
	return s.DeleteDynamicInterfaceV1(DeleteDynamicInterfaceV1HeaderParams, DeleteDynamicInterfaceV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `DeleteRfProfileV1`
*/
func (s *WirelessService) DeleteRfProfile(id string) (*ResponseWirelessDeleteRfProfileV1, *resty.Response, error) {
	return s.DeleteRfProfileV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `UpdateWirelessProfileV1`
*/
func (s *WirelessService) UpdateWirelessProfile(requestWirelessUpdateWirelessProfileV1 *RequestWirelessUpdateWirelessProfileV1) (*ResponseWirelessUpdateWirelessProfileV1, *resty.Response, error) {
	return s.UpdateWirelessProfileV1(requestWirelessUpdateWirelessProfileV1)
}

// Alias Function
/*
This method acts as an alias for the method `UpdateWirelessProfileConnectivityV1`
*/
func (s *WirelessService) UpdateWirelessProfileConnectivity(id string, requestWirelessUpdateWirelessProfileConnectivityV1 *RequestWirelessUpdateWirelessProfileConnectivityV1) (*ResponseWirelessUpdateWirelessProfileConnectivityV1, *resty.Response, error) {
	return s.UpdateWirelessProfileConnectivityV1(id, requestWirelessUpdateWirelessProfileConnectivityV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetAnchorManagedApLocationsForSpecificWirelessControllerV1`
*/
func (s *WirelessService) GetAnchorManagedApLocationsForSpecificWirelessController(networkDeviceID string, GetAnchorManagedAPLocationsForSpecificWirelessControllerV1QueryParams *GetAnchorManagedApLocationsForSpecificWirelessControllerV1QueryParams) (*ResponseWirelessGetAnchorManagedApLocationsForSpecificWirelessControllerV1, *resty.Response, error) {
	return s.GetAnchorManagedApLocationsForSpecificWirelessControllerV1(networkDeviceID, GetAnchorManagedAPLocationsForSpecificWirelessControllerV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetMeshApNeighboursCountV1`
*/
func (s *WirelessService) GetMeshApNeighboursCount() (*ResponseWirelessGetMeshApNeighboursCountV1, *resty.Response, error) {
	return s.GetMeshApNeighboursCountV1()
}

// Alias Function
/*
This method acts as an alias for the method `CreateUpdateDynamicInterfaceV1`
*/
func (s *WirelessService) CreateUpdateDynamicInterface(requestWirelessCreateUpdateDynamicInterfaceV1 *RequestWirelessCreateUpdateDynamicInterfaceV1) (*ResponseWirelessCreateUpdateDynamicInterfaceV1, *resty.Response, error) {
	return s.CreateUpdateDynamicInterfaceV1(requestWirelessCreateUpdateDynamicInterfaceV1)
}

// Alias Function
/*
This method acts as an alias for the method `RetrieveTheCountOfPolicyTagsForAWirelessProfileV1`
*/
func (s *WirelessService) RetrieveTheCountOfPolicyTagsForAWirelessProfile(id string) (*ResponseWirelessRetrieveTheCountOfPolicyTagsForAWirelessProfileV1, *resty.Response, error) {
	return s.RetrieveTheCountOfPolicyTagsForAWirelessProfileV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `CreateAnchorGroupV1`
*/
func (s *WirelessService) CreateAnchorGroup(requestWirelessCreateAnchorGroupV1 *RequestWirelessCreateAnchorGroupV1) (*ResponseWirelessCreateAnchorGroupV1, *resty.Response, error) {
	return s.CreateAnchorGroupV1(requestWirelessCreateAnchorGroupV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetInterfacesV1`
*/
func (s *WirelessService) GetInterfaces(GetInterfacesV1QueryParams *GetInterfacesV1QueryParams) (*ResponseWirelessGetInterfacesV1, *resty.Response, error) {
	return s.GetInterfacesV1(GetInterfacesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetEnterpriseSSIDV1`
*/
func (s *WirelessService) GetEnterpriseSSID(GetEnterpriseSSIDV1QueryParams *GetEnterpriseSSIDV1QueryParams) (*ResponseWirelessGetEnterpriseSSIDV1, *resty.Response, error) {
	return s.GetEnterpriseSSIDV1(GetEnterpriseSSIDV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetSSIDByIDV1`
*/
func (s *WirelessService) GetSSIDByID(siteID string, id string) (*ResponseWirelessGetSSIDByIDV1, *resty.Response, error) {
	return s.GetSSIDByIDV1(siteID, id)
}

// Alias Function
/*
This method acts as an alias for the method `RetrieveRfProfilesV1`
*/
func (s *WirelessService) RetrieveRfProfiles(RetrieveRFProfilesV1QueryParams *RetrieveRfProfilesV1QueryParams) (*ResponseWirelessRetrieveRfProfilesV1, *resty.Response, error) {
	return s.RetrieveRfProfilesV1(RetrieveRFProfilesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `ConfigureAccessPointsV1`
*/
func (s *WirelessService) ConfigureAccessPoints(requestWirelessConfigureAccessPointsV1 *RequestWirelessConfigureAccessPointsV1) (*ResponseWirelessConfigureAccessPointsV1, *resty.Response, error) {
	return s.ConfigureAccessPointsV1(requestWirelessConfigureAccessPointsV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetRfProfileByIDV1`
*/
func (s *WirelessService) GetRfProfileByID(id string) (*ResponseWirelessGetRfProfileByIDV1, *resty.Response, error) {
	return s.GetRfProfileByIDV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `CreateWirelessProfileV1`
*/
func (s *WirelessService) CreateWirelessProfile(requestWirelessCreateWirelessProfileV1 *RequestWirelessCreateWirelessProfileV1) (*ResponseWirelessCreateWirelessProfileV1, *resty.Response, error) {
	return s.CreateWirelessProfileV1(requestWirelessCreateWirelessProfileV1)
}

// Alias Function
/*
This method acts as an alias for the method `DeleteASpecificPolicyTagFromAWirelessProfileV1`
*/
func (s *WirelessService) DeleteASpecificPolicyTagFromAWirelessProfile(id string, policyTagID string) (*ResponseWirelessDeleteASpecificPolicyTagFromAWirelessProfileV1, *resty.Response, error) {
	return s.DeleteASpecificPolicyTagFromAWirelessProfileV1(id, policyTagID)
}

// Alias Function
/*
This method acts as an alias for the method `MobilityProvisionV1`
*/
func (s *WirelessService) MobilityProvision(requestWirelessMobilityProvisionV1 *RequestWirelessMobilityProvisionV1) (*ResponseWirelessMobilityProvisionV1, *resty.Response, error) {
	return s.MobilityProvisionV1(requestWirelessMobilityProvisionV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetAnchorGroupByIDV1`
*/
func (s *WirelessService) GetAnchorGroupByID(id string) (*ResponseWirelessGetAnchorGroupByIDV1, *resty.Response, error) {
	return s.GetAnchorGroupByIDV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `DeleteA80211BeProfileV1`
*/
func (s *WirelessService) DeleteA80211BeProfile(id string) (*ResponseWirelessDeleteA80211BeProfileV1, *resty.Response, error) {
	return s.DeleteA80211BeProfileV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `GetMobilityGroupsCountV1`
*/
func (s *WirelessService) GetMobilityGroupsCount() (*ResponseWirelessGetMobilityGroupsCountV1, *resty.Response, error) {
	return s.GetMobilityGroupsCountV1()
}

// Alias Function
/*
This method acts as an alias for the method `GetRfProfilesV1`
*/
func (s *WirelessService) GetRfProfiles(GetRFProfilesV1QueryParams *GetRfProfilesV1QueryParams) (*ResponseWirelessGetRfProfilesV1, *resty.Response, error) {
	return s.GetRfProfilesV1(GetRFProfilesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `DeleteApAuthorizationListV1`
*/
func (s *WirelessService) DeleteApAuthorizationList(id string) (*ResponseWirelessDeleteApAuthorizationListV1, *resty.Response, error) {
	return s.DeleteApAuthorizationListV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `GetMobilityGroupsV1`
*/
func (s *WirelessService) GetMobilityGroups(GetMobilityGroupsV1QueryParams *GetMobilityGroupsV1QueryParams) (*ResponseWirelessGetMobilityGroupsV1, *resty.Response, error) {
	return s.GetMobilityGroupsV1(GetMobilityGroupsV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `AssignAnchorManagedApLocationsForWLCV1`
*/
func (s *WirelessService) AssignAnchorManagedApLocationsForWLC(networkDeviceID string, requestWirelessAssignAnchorManagedAPLocationsForWLCV1 *RequestWirelessAssignAnchorManagedApLocationsForWLCV1) (*ResponseWirelessAssignAnchorManagedApLocationsForWLCV1, *resty.Response, error) {
	return s.AssignAnchorManagedApLocationsForWLCV1(networkDeviceID, requestWirelessAssignAnchorManagedAPLocationsForWLCV1)
}

// Alias Function
/*
This method acts as an alias for the method `DeleteAnchorGroupByIDV1`
*/
func (s *WirelessService) DeleteAnchorGroupByID(id string) (*ResponseWirelessDeleteAnchorGroupByIDV1, *resty.Response, error) {
	return s.DeleteAnchorGroupByIDV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `DeleteWirelessProfileV1`
*/
func (s *WirelessService) DeleteWirelessProfile(wirelessProfileName string) (*ResponseWirelessDeleteWirelessProfileV1, *resty.Response, error) {
	return s.DeleteWirelessProfileV1(wirelessProfileName)
}

// Alias Function
/*
This method acts as an alias for the method `GetRfProfilesCountV1`
*/
func (s *WirelessService) GetRfProfilesCount() (*ResponseWirelessGetRfProfilesCountV1, *resty.Response, error) {
	return s.GetRfProfilesCountV1()
}

// Alias Function
/*
This method acts as an alias for the method `UpdateASpecificPolicyTagForAWirelessProfileV1`
*/
func (s *WirelessService) UpdateASpecificPolicyTagForAWirelessProfile(id string, policyTagID string, requestWirelessUpdateASpecificPolicyTagForAWirelessProfileV1 *RequestWirelessUpdateASpecificPolicyTagForAWirelessProfileV1) (*ResponseWirelessUpdateASpecificPolicyTagForAWirelessProfileV1, *resty.Response, error) {
	return s.UpdateASpecificPolicyTagForAWirelessProfileV1(id, policyTagID, requestWirelessUpdateASpecificPolicyTagForAWirelessProfileV1)
}

// Alias Function
/*
This method acts as an alias for the method `MobilityResetV1`
*/
func (s *WirelessService) MobilityReset(requestWirelessMobilityResetV1 *RequestWirelessMobilityResetV1) (*ResponseWirelessMobilityResetV1, *resty.Response, error) {
	return s.MobilityResetV1(requestWirelessMobilityResetV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetCountOfAnchorGroupsV1`
*/
func (s *WirelessService) GetCountOfAnchorGroups() (*ResponseWirelessGetCountOfAnchorGroupsV1, *resty.Response, error) {
	return s.GetCountOfAnchorGroupsV1()
}

// Alias Function
/*
This method acts as an alias for the method `GetAccessPointConfigurationCountV1`
*/
func (s *WirelessService) GetAccessPointConfigurationCount(GetAccessPointConfigurationCountV1QueryParams *GetAccessPointConfigurationCountV1QueryParams) (*ResponseWirelessGetAccessPointConfigurationCountV1, *resty.Response, error) {
	return s.GetAccessPointConfigurationCountV1(GetAccessPointConfigurationCountV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `UpdateASpecificSiteTagForAWirelessProfileV1`
*/
func (s *WirelessService) UpdateASpecificSiteTagForAWirelessProfile(id string, siteTagID string, requestWirelessUpdateASpecificSiteTagForAWirelessProfileV1 *RequestWirelessUpdateASpecificSiteTagForAWirelessProfileV1) (*ResponseWirelessUpdateASpecificSiteTagForAWirelessProfileV1, *resty.Response, error) {
	return s.UpdateASpecificSiteTagForAWirelessProfileV1(id, siteTagID, requestWirelessUpdateASpecificSiteTagForAWirelessProfileV1)
}

// Alias Function
/*
This method acts as an alias for the method `UpdateApAuthorizationListV1`
*/
func (s *WirelessService) UpdateApAuthorizationList(id string, requestWirelessUpdateAPAuthorizationListV1 *RequestWirelessUpdateApAuthorizationListV1) (*ResponseWirelessUpdateApAuthorizationListV1, *resty.Response, error) {
	return s.UpdateApAuthorizationListV1(id, requestWirelessUpdateAPAuthorizationListV1)
}

// Alias Function
/*
This method acts as an alias for the method `Update80211BeProfileV1`
*/
func (s *WirelessService) Update80211BeProfile(id string, requestWirelessUpdate80211beProfileV1 *RequestWirelessUpdate80211BeProfileV1) (*ResponseWirelessUpdate80211BeProfileV1, *resty.Response, error) {
	return s.Update80211BeProfileV1(id, requestWirelessUpdate80211beProfileV1)
}

// Alias Function
/*
This method acts as an alias for the method `RetrieveTheCountOfSiteTagsForAWirelessProfileV1`
*/
func (s *WirelessService) RetrieveTheCountOfSiteTagsForAWirelessProfile(id string) (*ResponseWirelessRetrieveTheCountOfSiteTagsForAWirelessProfileV1, *resty.Response, error) {
	return s.RetrieveTheCountOfSiteTagsForAWirelessProfileV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `GetDynamicInterfaceV1`
*/
func (s *WirelessService) GetDynamicInterface(GetDynamicInterfaceV1QueryParams *GetDynamicInterfaceV1QueryParams) (*ResponseWirelessGetDynamicInterfaceV1, *resty.Response, error) {
	return s.GetDynamicInterfaceV1(GetDynamicInterfaceV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetInterfaceByIDV1`
*/
func (s *WirelessService) GetInterfaceByID(id string) (*ResponseWirelessGetInterfaceByIDV1, *resty.Response, error) {
	return s.GetInterfaceByIDV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `GetWirelessProfileV1`
*/
func (s *WirelessService) GetWirelessProfile(GetWirelessProfileV1QueryParams *GetWirelessProfileV1QueryParams) (*ResponseWirelessGetWirelessProfileV1, *resty.Response, error) {
	return s.GetWirelessProfileV1(GetWirelessProfileV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `RetrieveASpecificPolicyTagForAWirelessProfileV1`
*/
func (s *WirelessService) RetrieveASpecificPolicyTagForAWirelessProfile(id string, policyTagID string) (*ResponseWirelessRetrieveASpecificPolicyTagForAWirelessProfileV1, *resty.Response, error) {
	return s.RetrieveASpecificPolicyTagForAWirelessProfileV1(id, policyTagID)
}

// Alias Function
/*
This method acts as an alias for the method `CreatePowerProfileV1`
*/
func (s *WirelessService) CreatePowerProfile(requestWirelessCreatePowerProfileV1 *RequestWirelessCreatePowerProfileV1) (*ResponseWirelessCreatePowerProfileV1, *resty.Response, error) {
	return s.CreatePowerProfileV1(requestWirelessCreatePowerProfileV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetPowerProfilesCountV1`
*/
func (s *WirelessService) GetPowerProfilesCount() (*ResponseWirelessGetPowerProfilesCountV1, *resty.Response, error) {
	return s.GetPowerProfilesCountV1()
}

// Alias Function
/*
This method acts as an alias for the method `CreateSSIDV1`
*/
func (s *WirelessService) CreateSSID(siteID string, requestWirelessCreateSSIDV1 *RequestWirelessCreateSSIDV1) (*ResponseWirelessCreateSSIDV1, *resty.Response, error) {
	return s.CreateSSIDV1(siteID, requestWirelessCreateSSIDV1)
}

// Alias Function
/*
This method acts as an alias for the method `CreateEnterpriseSSIDV1`
*/
func (s *WirelessService) CreateEnterpriseSSID(requestWirelessCreateEnterpriseSSIDV1 *RequestWirelessCreateEnterpriseSSIDV1) (*ResponseWirelessCreateEnterpriseSSIDV1, *resty.Response, error) {
	return s.CreateEnterpriseSSIDV1(requestWirelessCreateEnterpriseSSIDV1)
}

// Alias Function
/*
This method acts as an alias for the method `DeletePowerProfileByIDV1`
*/
func (s *WirelessService) DeletePowerProfileByID(id string) (*ResponseWirelessDeletePowerProfileByIDV1, *resty.Response, error) {
	return s.DeletePowerProfileByIDV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `UpdateRfProfileV1`
*/
func (s *WirelessService) UpdateRfProfile(id string, requestWirelessUpdateRFProfileV1 *RequestWirelessUpdateRfProfileV1) (*ResponseWirelessUpdateRfProfileV1, *resty.Response, error) {
	return s.UpdateRfProfileV1(id, requestWirelessUpdateRFProfileV1)
}

// Alias Function
/*
This method acts as an alias for the method `DeleteApProfileByIDV1`
*/
func (s *WirelessService) DeleteApProfileByID(id string) (*ResponseWirelessDeleteApProfileByIDV1, *resty.Response, error) {
	return s.DeleteApProfileByIDV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `GetApAuthorizationListByIDV1`
*/
func (s *WirelessService) GetApAuthorizationListByID(id string) (*ResponseWirelessGetApAuthorizationListByIDV1, *resty.Response, error) {
	return s.GetApAuthorizationListByIDV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `GetApAuthorizationListsV1`
*/
func (s *WirelessService) GetApAuthorizationLists(GetAPAuthorizationListsV1QueryParams *GetApAuthorizationListsV1QueryParams) (*ResponseWirelessGetApAuthorizationListsV1, *resty.Response, error) {
	return s.GetApAuthorizationListsV1(GetAPAuthorizationListsV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetAccessPointConfigurationTaskResultV1`
*/
func (s *WirelessService) GetAccessPointConfigurationTaskResult(taskTypeID string) (*ResponseWirelessGetAccessPointConfigurationTaskResultV1, *resty.Response, error) {
	return s.GetAccessPointConfigurationTaskResultV1(taskTypeID)
}

// Alias Function
/*
This method acts as an alias for the method `GetWirelessProfileByIDV1`
*/
func (s *WirelessService) GetWirelessProfileByID(id string) (*ResponseWirelessGetWirelessProfileByIDV1, *resty.Response, error) {
	return s.GetWirelessProfileByIDV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `UpdatePowerProfileByIDV1`
*/
func (s *WirelessService) UpdatePowerProfileByID(id string, requestWirelessUpdatePowerProfileByIDV1 *RequestWirelessUpdatePowerProfileByIDV1) (*ResponseWirelessUpdatePowerProfileByIDV1, *resty.Response, error) {
	return s.UpdatePowerProfileByIDV1(id, requestWirelessUpdatePowerProfileByIDV1)
}

// Alias Function
/*
This method acts as an alias for the method `Get80211BeProfileByIDV1`
*/
func (s *WirelessService) Get80211BeProfileByID(id string) (*ResponseWirelessGet80211BeProfileByIDV1, *resty.Response, error) {
	return s.Get80211BeProfileByIDV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `GetManagedApLocationsCountForSpecificWirelessControllerV1`
*/
func (s *WirelessService) GetManagedApLocationsCountForSpecificWirelessController(networkDeviceID string) (*ResponseWirelessGetManagedApLocationsCountForSpecificWirelessControllerV1, *resty.Response, error) {
	return s.GetManagedApLocationsCountForSpecificWirelessControllerV1(networkDeviceID)
}

// Alias Function
/*
This method acts as an alias for the method `CreateMultiplePolicyTagsForAWirelessProfileInBulkV1`
*/
func (s *WirelessService) CreateMultiplePolicyTagsForAWirelessProfileInBulk(id string, requestWirelessCreateMultiplePolicyTagsForAWirelessProfileInBulkV1 *RequestWirelessCreateMultiplePolicyTagsForAWirelessProfileInBulkV1) (*ResponseWirelessCreateMultiplePolicyTagsForAWirelessProfileInBulkV1, *resty.Response, error) {
	return s.CreateMultiplePolicyTagsForAWirelessProfileInBulkV1(id, requestWirelessCreateMultiplePolicyTagsForAWirelessProfileInBulkV1)
}

// Alias Function
/*
This method acts as an alias for the method `UpdateEnterpriseSSIDV1`
*/
func (s *WirelessService) UpdateEnterpriseSSID(requestWirelessUpdateEnterpriseSSIDV1 *RequestWirelessUpdateEnterpriseSSIDV1) (*ResponseWirelessUpdateEnterpriseSSIDV1, *resty.Response, error) {
	return s.UpdateEnterpriseSSIDV1(requestWirelessUpdateEnterpriseSSIDV1)
}

// Alias Function
/*
This method acts as an alias for the method `ApProvisionV1`
*/
func (s *WirelessService) ApProvision(requestWirelessAPProvisionV1 *RequestWirelessApProvisionV1) (*ResponseWirelessApProvisionV1, *resty.Response, error) {
	return s.ApProvisionV1(requestWirelessAPProvisionV1)
}

// Alias Function
/*
This method acts as an alias for the method `PSKOverrideV1`
*/
func (s *WirelessService) PSKOverride(requestWirelessPSKOverrideV1 *RequestWirelessPSKOverrideV1) (*ResponseWirelessPSKOverrideV1, *resty.Response, error) {
	return s.PSKOverrideV1(requestWirelessPSKOverrideV1)
}

// Alias Function
/*
This method acts as an alias for the method `ProvisionUpdateV1`
*/
func (s *WirelessService) ProvisionUpdate(requestWirelessProvisionUpdateV1 *RequestWirelessProvisionUpdateV1, ProvisionUpdateV1HeaderParams *ProvisionUpdateV1HeaderParams) (*ResponseWirelessProvisionUpdateV1, *resty.Response, error) {
	return s.ProvisionUpdateV1(requestWirelessProvisionUpdateV1, ProvisionUpdateV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `CreateApProfileV1`
*/
func (s *WirelessService) CreateApProfile(requestWirelessCreateAPProfileV1 *RequestWirelessCreateApProfileV1) (*ResponseWirelessCreateApProfileV1, *resty.Response, error) {
	return s.CreateApProfileV1(requestWirelessCreateAPProfileV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetSSIDBySiteV1`
*/
func (s *WirelessService) GetSSIDBySite(siteID string, GetSSIDBySiteV1QueryParams *GetSSIDBySiteV1QueryParams) (*ResponseWirelessGetSSIDBySiteV1, *resty.Response, error) {
	return s.GetSSIDBySiteV1(siteID, GetSSIDBySiteV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `SensorTestResultsV1`
*/
func (s *WirelessService) SensorTestResults(SensorTestResultsV1QueryParams *SensorTestResultsV1QueryParams) (*ResponseWirelessSensorTestResultsV1, *resty.Response, error) {
	return s.SensorTestResultsV1(SensorTestResultsV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `RetrieveSitesWithOverriddenSSIDsV1`
*/
func (s *WirelessService) RetrieveSitesWithOverriddenSSIDs(RetrieveSitesWithOverriddenSSIDsV1QueryParams *RetrieveSitesWithOverriddenSSIDsV1QueryParams) (*ResponseWirelessRetrieveSitesWithOverriddenSSIDsV1, *resty.Response, error) {
	return s.RetrieveSitesWithOverriddenSSIDsV1(RetrieveSitesWithOverriddenSSIDsV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `RetrieveASpecificSiteTagForAWirelessProfileV1`
*/
func (s *WirelessService) RetrieveASpecificSiteTagForAWirelessProfile(id string, siteTagID string) (*ResponseWirelessRetrieveASpecificSiteTagForAWirelessProfileV1, *resty.Response, error) {
	return s.RetrieveASpecificSiteTagForAWirelessProfileV1(id, siteTagID)
}

// Alias Function
/*
This method acts as an alias for the method `Get80211BeProfilesV1`
*/
func (s *WirelessService) Get80211BeProfiles(Get80211beProfilesV1QueryParams *Get80211BeProfilesV1QueryParams) (*ResponseWirelessGet80211BeProfilesV1, *resty.Response, error) {
	return s.Get80211BeProfilesV1(Get80211beProfilesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetAnchorGroupsV1`
*/
func (s *WirelessService) GetAnchorGroups() (*ResponseWirelessGetAnchorGroupsV1, *resty.Response, error) {
	return s.GetAnchorGroupsV1()
}

// Alias Function
/*
This method acts as an alias for the method `DeleteInterfaceV1`
*/
func (s *WirelessService) DeleteInterface(id string) (*ResponseWirelessDeleteInterfaceV1, *resty.Response, error) {
	return s.DeleteInterfaceV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `GetSSIDCountForSpecificWirelessControllerV1`
*/
func (s *WirelessService) GetSSIDCountForSpecificWirelessController(networkDeviceID string, GetSSIDCountForSpecificWirelessControllerV1QueryParams *GetSSIDCountForSpecificWirelessControllerV1QueryParams) (*ResponseWirelessGetSSIDCountForSpecificWirelessControllerV1, *resty.Response, error) {
	return s.GetSSIDCountForSpecificWirelessControllerV1(networkDeviceID, GetSSIDCountForSpecificWirelessControllerV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `RetrieveAllSiteTagsForAWirelessProfileV1`
*/
func (s *WirelessService) RetrieveAllSiteTagsForAWirelessProfile(id string, RetrieveAllSiteTagsForAWirelessProfileV1QueryParams *RetrieveAllSiteTagsForAWirelessProfileV1QueryParams) (*ResponseWirelessRetrieveAllSiteTagsForAWirelessProfileV1, *resty.Response, error) {
	return s.RetrieveAllSiteTagsForAWirelessProfileV1(id, RetrieveAllSiteTagsForAWirelessProfileV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `ProvisionV1`
*/
func (s *WirelessService) Provision(requestWirelessProvisionV1 *RequestWirelessProvisionV1) (*ResponseWirelessProvisionV1, *resty.Response, error) {
	return s.ProvisionV1(requestWirelessProvisionV1)
}

// Alias Function
/*
This method acts as an alias for the method `UpdateAnchorGroupV1`
*/
func (s *WirelessService) UpdateAnchorGroup(id string, requestWirelessUpdateAnchorGroupV1 *RequestWirelessUpdateAnchorGroupV1) (*ResponseWirelessUpdateAnchorGroupV1, *resty.Response, error) {
	return s.UpdateAnchorGroupV1(id, requestWirelessUpdateAnchorGroupV1)
}

// Alias Function
/*
This method acts as an alias for the method `RebootAccessPointsV1`
*/
func (s *WirelessService) RebootAccessPoints(requestWirelessRebootAccessPointsV1 *RequestWirelessRebootAccessPointsV1) (*ResponseWirelessRebootAccessPointsV1, *resty.Response, error) {
	return s.RebootAccessPointsV1(requestWirelessRebootAccessPointsV1)
}

// Alias Function
/*
This method acts as an alias for the method `ApProvisionConnectivityV1`
*/
func (s *WirelessService) ApProvisionConnectivity(requestWirelessAPProvisionConnectivityV1 *RequestWirelessApProvisionConnectivityV1, APProvisionConnectivityV1HeaderParams *ApProvisionConnectivityV1HeaderParams) (*ResponseWirelessApProvisionConnectivityV1, *resty.Response, error) {
	return s.ApProvisionConnectivityV1(requestWirelessAPProvisionConnectivityV1, APProvisionConnectivityV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetApAuthorizationListByNetworkDeviceIDV1`
*/
func (s *WirelessService) GetApAuthorizationListByNetworkDeviceID(networkDeviceID string) (*ResponseWirelessGetApAuthorizationListByNetworkDeviceIDV1, *resty.Response, error) {
	return s.GetApAuthorizationListByNetworkDeviceIDV1(networkDeviceID)
}

// Alias Function
/*
This method acts as an alias for the method `CreateAndProvisionSSIDV1`
*/
func (s *WirelessService) CreateAndProvisionSSID(requestWirelessCreateAndProvisionSSIDV1 *RequestWirelessCreateAndProvisionSSIDV1, CreateAndProvisionSSIDV1HeaderParams *CreateAndProvisionSSIDV1HeaderParams) (*ResponseWirelessCreateAndProvisionSSIDV1, *resty.Response, error) {
	return s.CreateAndProvisionSSIDV1(requestWirelessCreateAndProvisionSSIDV1, CreateAndProvisionSSIDV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `CreateOrUpdateRfProfileV1`
*/
func (s *WirelessService) CreateOrUpdateRfProfile(requestWirelessCreateOrUpdateRFProfileV1 *RequestWirelessCreateOrUpdateRfProfileV1) (*ResponseWirelessCreateOrUpdateRfProfileV1, *resty.Response, error) {
	return s.CreateOrUpdateRfProfileV1(requestWirelessCreateOrUpdateRFProfileV1)
}

// Alias Function
/*
This method acts as an alias for the method `UpdateOrOverridessidV1`
*/
func (s *WirelessService) UpdateOrOverridessid(siteID string, id string, requestWirelessUpdateOrOverrideSSIDV1 *RequestWirelessUpdateOrOverridessidV1) (*ResponseWirelessUpdateOrOverridessidV1, *resty.Response, error) {
	return s.UpdateOrOverridessidV1(siteID, id, requestWirelessUpdateOrOverrideSSIDV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetApProfilesCountV1`
*/
func (s *WirelessService) GetApProfilesCount() (*ResponseWirelessGetApProfilesCountV1, *resty.Response, error) {
	return s.GetApProfilesCountV1()
}

// Alias Function
/*
This method acts as an alias for the method `CreateApAuthorizationListV1`
*/
func (s *WirelessService) CreateApAuthorizationList(requestWirelessCreateAPAuthorizationListV1 *RequestWirelessCreateApAuthorizationListV1) (*ResponseWirelessCreateApAuthorizationListV1, *resty.Response, error) {
	return s.CreateApAuthorizationListV1(requestWirelessCreateAPAuthorizationListV1)
}

// Alias Function
/*
This method acts as an alias for the method `CreateInterfaceV1`
*/
func (s *WirelessService) CreateInterface(requestWirelessCreateInterfaceV1 *RequestWirelessCreateInterfaceV1) (*ResponseWirelessCreateInterfaceV1, *resty.Response, error) {
	return s.CreateInterfaceV1(requestWirelessCreateInterfaceV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetSSIDDetailsForSpecificWirelessControllerV1`
*/
func (s *WirelessService) GetSSIDDetailsForSpecificWirelessController(networkDeviceID string, GetSSIDDetailsForSpecificWirelessControllerV1QueryParams *GetSSIDDetailsForSpecificWirelessControllerV1QueryParams) (*ResponseWirelessGetSSIDDetailsForSpecificWirelessControllerV1, *resty.Response, error) {
	return s.GetSSIDDetailsForSpecificWirelessControllerV1(networkDeviceID, GetSSIDDetailsForSpecificWirelessControllerV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `UpdateApProfileByIDV1`
*/
func (s *WirelessService) UpdateApProfileByID(id string, requestWirelessUpdateAPProfileByIDV1 *RequestWirelessUpdateApProfileByIDV1) (*ResponseWirelessUpdateApProfileByIDV1, *resty.Response, error) {
	return s.UpdateApProfileByIDV1(id, requestWirelessUpdateAPProfileByIDV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetApProfilesV1`
*/
func (s *WirelessService) GetApProfiles(GetAPProfilesV1QueryParams *GetApProfilesV1QueryParams) (*ResponseWirelessGetApProfilesV1, *resty.Response, error) {
	return s.GetApProfilesV1(GetAPProfilesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `DeleteSSIDAndProvisionItToDevicesV1`
*/
func (s *WirelessService) DeleteSSIDAndProvisionItToDevices(ssidName string, managedAPLocations string, DeleteSSIDAndProvisionItToDevicesV1HeaderParams *DeleteSSIDAndProvisionItToDevicesV1HeaderParams) (*ResponseWirelessDeleteSSIDAndProvisionItToDevicesV1, *resty.Response, error) {
	return s.DeleteSSIDAndProvisionItToDevicesV1(ssidName, managedAPLocations, DeleteSSIDAndProvisionItToDevicesV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `UpdateInterfaceV1`
*/
func (s *WirelessService) UpdateInterface(id string, requestWirelessUpdateInterfaceV1 *RequestWirelessUpdateInterfaceV1) (*ResponseWirelessUpdateInterfaceV1, *resty.Response, error) {
	return s.UpdateInterfaceV1(id, requestWirelessUpdateInterfaceV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetWirelessProfilesV1`
*/
func (s *WirelessService) GetWirelessProfiles(GetWirelessProfilesV1QueryParams *GetWirelessProfilesV1QueryParams) (*ResponseWirelessGetWirelessProfilesV1, *resty.Response, error) {
	return s.GetWirelessProfilesV1(GetWirelessProfilesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `UpdateSSIDV1`
*/
func (s *WirelessService) UpdateSSID(siteID string, id string, requestWirelessUpdateSSIDV1 *RequestWirelessUpdateSSIDV1) (*ResponseWirelessUpdateSSIDV1, *resty.Response, error) {
	return s.UpdateSSIDV1(siteID, id, requestWirelessUpdateSSIDV1)
}

// Alias Function
/*
This method acts as an alias for the method `CreateA80211BeProfileV1`
*/
func (s *WirelessService) CreateA80211BeProfile(requestWirelessCreateA80211beProfileV1 *RequestWirelessCreateA80211BeProfileV1) (*ResponseWirelessCreateA80211BeProfileV1, *resty.Response, error) {
	return s.CreateA80211BeProfileV1(requestWirelessCreateA80211beProfileV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetWirelessProfilesCountV1`
*/
func (s *WirelessService) GetWirelessProfilesCount() (*ResponseWirelessGetWirelessProfilesCountV1, *resty.Response, error) {
	return s.GetWirelessProfilesCountV1()
}

// Alias Function
/*
This method acts as an alias for the method `GetAccessPointConfigurationV1`
*/
func (s *WirelessService) GetAccessPointConfiguration(GetAccessPointConfigurationV1QueryParams *GetAccessPointConfigurationV1QueryParams) (*ResponseWirelessGetAccessPointConfigurationV1, *resty.Response, error) {
	return s.GetAccessPointConfigurationV1(GetAccessPointConfigurationV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `RetrieveAllPolicyTagsForAWirelessProfileV1`
*/
func (s *WirelessService) RetrieveAllPolicyTagsForAWirelessProfile(id string, RetrieveAllPolicyTagsForAWirelessProfileV1QueryParams *RetrieveAllPolicyTagsForAWirelessProfileV1QueryParams) (*ResponseWirelessRetrieveAllPolicyTagsForAWirelessProfileV1, *resty.Response, error) {
	return s.RetrieveAllPolicyTagsForAWirelessProfileV1(id, RetrieveAllPolicyTagsForAWirelessProfileV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `DeleteRfProfilesV1`
*/
func (s *WirelessService) DeleteRfProfiles(rfProfileName string) (*ResponseWirelessDeleteRfProfilesV1, *resty.Response, error) {
	return s.DeleteRfProfilesV1(rfProfileName)
}

// Alias Function
/*
This method acts as an alias for the method `DeleteSSIDV1`
*/
func (s *WirelessService) DeleteSSID(siteID string, id string) (*ResponseWirelessDeleteSSIDV1, *resty.Response, error) {
	return s.DeleteSSIDV1(siteID, id)
}

// Alias Function
/*
This method acts as an alias for the method `DeleteEnterpriseSSIDV1`
*/
func (s *WirelessService) DeleteEnterpriseSSID(ssidName string) (*ResponseWirelessDeleteEnterpriseSSIDV1, *resty.Response, error) {
	return s.DeleteEnterpriseSSIDV1(ssidName)
}

// Alias Function
/*
This method acts as an alias for the method `DeleteWirelessProfileConnectivityV1`
*/
func (s *WirelessService) DeleteWirelessProfileConnectivity(id string) (*ResponseWirelessDeleteWirelessProfileConnectivityV1, *resty.Response, error) {
	return s.DeleteWirelessProfileConnectivityV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `GetPowerProfilesV1`
*/
func (s *WirelessService) GetPowerProfiles(GetPowerProfilesV1QueryParams *GetPowerProfilesV1QueryParams) (*ResponseWirelessGetPowerProfilesV1, *resty.Response, error) {
	return s.GetPowerProfilesV1(GetPowerProfilesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `Get80211BeProfilesCountV1`
*/
func (s *WirelessService) Get80211BeProfilesCount() (*ResponseWirelessGet80211BeProfilesCountV1, *resty.Response, error) {
	return s.Get80211BeProfilesCountV1()
}

// Alias Function
/*
This method acts as an alias for the method `GetPowerProfileByIDV1`
*/
func (s *WirelessService) GetPowerProfileByID(id string) (*ResponseWirelessGetPowerProfileByIDV1, *resty.Response, error) {
	return s.GetPowerProfileByIDV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `GetSSIDCountBySiteV1`
*/
func (s *WirelessService) GetSSIDCountBySite(siteID string) (*ResponseWirelessGetSSIDCountBySiteV1, *resty.Response, error) {
	return s.GetSSIDCountBySiteV1(siteID)
}

// Alias Function
/*
This method acts as an alias for the method `GetMeshApNeighboursV1`
*/
func (s *WirelessService) GetMeshApNeighbours(GetMeshApNeighboursV1QueryParams *GetMeshApNeighboursV1QueryParams) (*ResponseWirelessGetMeshApNeighboursV1, *resty.Response, error) {
	return s.GetMeshApNeighboursV1(GetMeshApNeighboursV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetPrimaryManagedApLocationsForSpecificWirelessControllerV1`
*/
func (s *WirelessService) GetPrimaryManagedApLocationsForSpecificWirelessController(networkDeviceID string, GetPrimaryManagedAPLocationsForSpecificWirelessControllerV1QueryParams *GetPrimaryManagedApLocationsForSpecificWirelessControllerV1QueryParams) (*ResponseWirelessGetPrimaryManagedApLocationsForSpecificWirelessControllerV1, *resty.Response, error) {
	return s.GetPrimaryManagedApLocationsForSpecificWirelessControllerV1(networkDeviceID, GetPrimaryManagedAPLocationsForSpecificWirelessControllerV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `CreateWirelessProfileConnectivityV1`
*/
func (s *WirelessService) CreateWirelessProfileConnectivity(requestWirelessCreateWirelessProfileConnectivityV1 *RequestWirelessCreateWirelessProfileConnectivityV1) (*ResponseWirelessCreateWirelessProfileConnectivityV1, *resty.Response, error) {
	return s.CreateWirelessProfileConnectivityV1(requestWirelessCreateWirelessProfileConnectivityV1)
}

// Alias Function
/*
This method acts as an alias for the method `WirelessControllerProvisionV1`
*/
func (s *WirelessService) WirelessControllerProvision(deviceID string, requestWirelessWirelessControllerProvisionV1 *RequestWirelessWirelessControllerProvisionV1) (*ResponseWirelessWirelessControllerProvisionV1, *resty.Response, error) {
	return s.WirelessControllerProvisionV1(deviceID, requestWirelessWirelessControllerProvisionV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetAccessPointRebootTaskResultV1`
*/
func (s *WirelessService) GetAccessPointRebootTaskResult(GetAccessPointRebootTaskResultV1QueryParams *GetAccessPointRebootTaskResultV1QueryParams) (*ResponseWirelessGetAccessPointRebootTaskResultV1, *resty.Response, error) {
	return s.GetAccessPointRebootTaskResultV1(GetAccessPointRebootTaskResultV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetSecondaryManagedApLocationsForSpecificWirelessControllerV1`
*/
func (s *WirelessService) GetSecondaryManagedApLocationsForSpecificWirelessController(networkDeviceID string, GetSecondaryManagedAPLocationsForSpecificWirelessControllerV1QueryParams *GetSecondaryManagedApLocationsForSpecificWirelessControllerV1QueryParams) (*ResponseWirelessGetSecondaryManagedApLocationsForSpecificWirelessControllerV1, *resty.Response, error) {
	return s.GetSecondaryManagedApLocationsForSpecificWirelessControllerV1(networkDeviceID, GetSecondaryManagedAPLocationsForSpecificWirelessControllerV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `DeleteASpecificSiteTagFromAWirelessProfileV1`
*/
func (s *WirelessService) DeleteASpecificSiteTagFromAWirelessProfile(id string, siteTagID string) (*ResponseWirelessDeleteASpecificSiteTagFromAWirelessProfileV1, *resty.Response, error) {
	return s.DeleteASpecificSiteTagFromAWirelessProfileV1(id, siteTagID)
}

// Alias Function
/*
This method acts as an alias for the method `GetApAuthorizationListCountV1`
*/
func (s *WirelessService) GetApAuthorizationListCount() (*ResponseWirelessGetApAuthorizationListCountV1, *resty.Response, error) {
	return s.GetApAuthorizationListCountV1()
}

// Alias Function
/*
This method acts as an alias for the method `GetApProfileByIDV1`
*/
func (s *WirelessService) GetApProfileByID(id string) (*ResponseWirelessGetApProfileByIDV1, *resty.Response, error) {
	return s.GetApProfileByIDV1(id)
}
