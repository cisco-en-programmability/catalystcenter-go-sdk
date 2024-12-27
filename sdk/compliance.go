package catalyst

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type ComplianceService service

type GetComplianceStatusV1QueryParams struct {
	ComplianceStatus string  `url:"complianceStatus,omitempty"` //Specify "Compliance status(es)" separated by commas. The Compliance status can be 'COMPLIANT', 'NON_COMPLIANT', 'IN_PROGRESS', 'NOT_AVAILABLE', 'NOT_APPLICABLE', 'ERROR'.
	DeviceUUID       string  `url:"deviceUuid,omitempty"`       //Comma separated 'Device Ids'
	Offset           float64 `url:"offset,omitempty"`           //offset/starting row	number
	Limit            float64 `url:"limit,omitempty"`            //The number of records to be retrieved defaults to 500 if not specified, with a maximum allowed limit of 500.
}
type GetComplianceStatusCountV1QueryParams struct {
	ComplianceStatus string `url:"complianceStatus,omitempty"` //Specify "Compliance status(es)" separated by commas. The Compliance status can be 'COMPLIANT', 'NON_COMPLIANT', 'IN_PROGRESS', 'NOT_AVAILABLE', 'NOT_APPLICABLE', 'ERROR'.
}
type GetComplianceDetailV1QueryParams struct {
	ComplianceType   string  `url:"complianceType,omitempty"`   //Specify "Compliance type(s)" in commas. The Compliance type can be 'NETWORK_PROFILE', 'IMAGE', 'FABRIC', 'APPLICATION_VISIBILITY', 'FABRIC', RUNNING_CONFIG', 'NETWORK_SETTINGS', 'WORKFLOW' , 'EoX'.
	ComplianceStatus string  `url:"complianceStatus,omitempty"` //Specify "Compliance status(es)" in commas. The Compliance status can be 'COMPLIANT', 'NON_COMPLIANT', 'IN_PROGRESS', 'NOT_AVAILABLE', 'NOT_APPLICABLE', 'ERROR'.
	DeviceUUID       string  `url:"deviceUuid,omitempty"`       //Comma separated "Device Id(s)"
	Offset           float64 `url:"offset,omitempty"`           //offset/starting row
	Limit            float64 `url:"limit,omitempty"`            //The number of records to be retrieved defaults to 500 if not specified, with a maximum allowed limit of 500.
}
type GetComplianceDetailCountV1QueryParams struct {
	ComplianceType   string `url:"complianceType,omitempty"`   //Specify "Compliance type(s)" separated by commas. The Compliance type can be 'APPLICATION_VISIBILITY', 'EoX', 'FABRIC', 'IMAGE', 'NETWORK_PROFILE', 'NETWORK_SETTINGS', 'PSIRT', 'RUNNING_CONFIG', 'WORKFLOW'.
	ComplianceStatus string `url:"complianceStatus,omitempty"` //Specify "Compliance status(es)" separated by commas. The Compliance status can be 'COMPLIANT', 'NON_COMPLIANT', 'IN_PROGRESS', 'NOT_AVAILABLE', 'NOT_APPLICABLE', 'ERROR'.
}
type ComplianceDetailsOfDeviceV1QueryParams struct {
	Category             string `url:"category,omitempty"`             //category can have any value among 'INTENT', 'RUNNING_CONFIG' , 'IMAGE' , 'PSIRT' , 'DESIGN_OOD' , 'EOX' , 'NETWORK_SETTINGS'
	ComplianceType       string `url:"complianceType,omitempty"`       //Specify "Compliance type(s)" separated by commas. The Compliance type can be 'APPLICATION_VISIBILITY', 'EOX', 'FABRIC', 'IMAGE', 'NETWORK_PROFILE', 'NETWORK_SETTINGS', 'PSIRT', 'RUNNING_CONFIG', 'WORKFLOW'.
	DiffList             bool   `url:"diffList,omitempty"`             //diff list [ pass true to fetch the diff list ]
	Status               string `url:"status,omitempty"`               //'COMPLIANT', 'NON_COMPLIANT', 'ERROR', 'IN_PROGRESS', 'NOT_APPLICABLE', 'NOT_AVAILABLE', 'WARNING', 'REMEDIATION_IN_PROGRESS' can be the value of the compliance 'status' parameter. [COMPLIANT: Device currently meets the compliance requirements.  NON_COMPLIANT: One of the compliance requirements like Software Image, PSIRT, Network Profile, Startup vs Running, etc. are not met. ERROR: Compliance is unable to compute status due to underlying errors. IN_PROGRESS: Compliance check is in progress for the device. NOT_APPLICABLE: Device is not supported for compliance, or minimum license requirement is not met. NOT_AVAILABLE: Compliance is not available for the device. COMPLIANT_WARNING: The device is compliant with warning if the last date of support is nearing. REMEDIATION_IN_PROGRESS: Compliance remediation is in progress for the device.]
	RemediationSupported bool   `url:"remediationSupported,omitempty"` //The 'remediationSupported' parameter can be set to 'true' or 'false'. The result will be a combination of both values if it is not provided.
}
type GetFieldNoticeNetworkDevicesV1QueryParams struct {
	NetworkDeviceID string  `url:"networkDeviceId,omitempty"` //Id of the network device
	ScanStatus      string  `url:"scanStatus,omitempty"`      //Status of the scan on the network device. Available values : NOT_SCANNED, IN_PROGRESS, SUCCESS, FAILED.
	NoticeCount     float64 `url:"noticeCount,omitempty"`     //Return network devices with noticeCount greater than this noticeCount
	Offset          float64 `url:"offset,omitempty"`          //The first record to show for this page; the first record is numbered 1. Default value is 1.
	Limit           float64 `url:"limit,omitempty"`           //The number of records to show for this page. Minimum value is 1. Maximum value is 500. Default value is 500.
	SortBy          string  `url:"sortBy,omitempty"`          //A property within the response to sort by.
	Order           string  `url:"order,omitempty"`           //Whether ascending or descending order should be used to sort the response. Available values : asc, desc. Default value is asc.
}
type GetCountOfFieldNoticeNetworkDevicesV1QueryParams struct {
	NetworkDeviceID string  `url:"networkDeviceId,omitempty"` //Id of the network device
	ScanStatus      string  `url:"scanStatus,omitempty"`      //Status of the scan on the network device. Available values : NOT_SCANNED, IN_PROGRESS, SUCCESS, FAILED.
	NoticeCount     float64 `url:"noticeCount,omitempty"`     //Return network devices with noticeCount greater than this noticeCount
}
type GetFieldNoticesAffectingTheNetworkDeviceV1QueryParams struct {
	ID     string  `url:"id,omitempty"`     //Id of the field notice
	Type   string  `url:"type,omitempty"`   //Return field notices with this type. Available values : SOFTWARE, HARDWARE.
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1. Default value is 1.
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page. Minimum value is 1. Maximum value is 500. Default value is 500.
	SortBy string  `url:"sortBy,omitempty"` //A property within the response to sort by.
	Order  string  `url:"order,omitempty"`  //Whether ascending or descending order should be used to sort the response. Available values : asc, desc. Default value is asc.
}
type GetCountOfFieldNoticesAffectingTheNetworkDeviceV1QueryParams struct {
	ID   string `url:"id,omitempty"`   //Id of the field notice
	Type string `url:"type,omitempty"` //Return field notices with this type. Available values : SOFTWARE, HARDWARE
}
type GetFieldNoticesV1QueryParams struct {
	ID          string  `url:"id,omitempty"`          //Id of the field notice
	DeviceCount float64 `url:"deviceCount,omitempty"` //Return field notices with deviceCount greater than this deviceCount
	Type        string  `url:"type,omitempty"`        //Return field notices with this type. Available values : SOFTWARE, HARDWARE.
	Offset      float64 `url:"offset,omitempty"`      //The first record to show for this page; the first record is numbered 1. Default value is 1.
	Limit       float64 `url:"limit,omitempty"`       //The number of records to show for this page. Minimum value is 1. Maximum value is 500. Default value is 500.
	SortBy      string  `url:"sortBy,omitempty"`      //A property within the response to sort by.
	Order       string  `url:"order,omitempty"`       //Whether ascending or descending order should be used to sort the response. Available values : asc, desc. Default value is asc.
}
type GetCountOfFieldNoticesV1QueryParams struct {
	ID          string  `url:"id,omitempty"`          //Id of the field notice
	DeviceCount float64 `url:"deviceCount,omitempty"` //Return field notices with deviceCount greater than this deviceCount
	Type        string  `url:"type,omitempty"`        //Return field notices with this type. Available values : SOFTWARE, HARDWARE
}
type GetFieldNoticeNetworkDevicesForTheNoticeV1QueryParams struct {
	NetworkDeviceID string  `url:"networkDeviceId,omitempty"` //Id of the network device
	ScanStatus      string  `url:"scanStatus,omitempty"`      //Status of the scan on the network device. Available values : NOT_SCANNED, IN_PROGRESS, SUCCESS, FAILED.
	Offset          float64 `url:"offset,omitempty"`          //The first record to show for this page; the first record is numbered 1. Default value is 1.
	Limit           float64 `url:"limit,omitempty"`           //The number of records to show for this page. Minimum value is 1. Maximum value is 500. Default value is 500.
	SortBy          string  `url:"sortBy,omitempty"`          //A property within the response to sort by.
	Order           string  `url:"order,omitempty"`           //Whether ascending or descending order should be used to sort the response. Available values : asc, desc. Default value is asc.
}
type GetCountOfFieldNoticeNetworkDevicesForTheNoticeV1QueryParams struct {
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` //id of the network device
	ScanStatus      string `url:"scanStatus,omitempty"`      //status of the scan on the network device. Available values : NOT_SCANNED, IN_PROGRESS, SUCCESS, FAILED.
}
type GetFieldNoticesResultsTrendOverTimeV1QueryParams struct {
	ScanTime float64 `url:"scanTime,omitempty"` //Return field notices trend with scanTime greater than this scanTime
	Offset   float64 `url:"offset,omitempty"`   //The first record to show for this page; the first record is numbered 1. Default value is 1.
	Limit    float64 `url:"limit,omitempty"`    //The number of records to show for this page. Minimum value is 1. Maximum value is 500. Default value is 500.
}
type GetCountOfFieldNoticesResultsTrendOverTimeV1QueryParams struct {
	ScanTime float64 `url:"scanTime,omitempty"` //Return field notices trend with scanTime greater than this scanTime
}
type TriggersAFieldNoticesScanForTheSupportedNetworkDevicesV1QueryParams struct {
	FailedDevicesOnly bool `url:"failedDevicesOnly,omitempty"` //Used to specify if the scan should run only for the network devices that failed during the previous scan. If not specified, this parameter defaults to false.
}
type GetConfigTaskDetailsV1QueryParams struct {
	ParentTaskID string `url:"parentTaskId,omitempty"` //task Id
}
type GetNetworkBugsV1QueryParams struct {
	ID          string  `url:"id,omitempty"`          //The id of the network bug
	DeviceCount float64 `url:"deviceCount,omitempty"` //Return network bugs with deviceCount greater than this deviceCount
	Severity    string  `url:"severity,omitempty"`    //Return network bugs with this severity. Available values : CATASTROPHIC, SEVERE, MODERATE
	Offset      float64 `url:"offset,omitempty"`      //The first record to show for this page; the first record is numbered 1. Default value is 1
	Limit       float64 `url:"limit,omitempty"`       //The number of records to show for this page. Minimum value is 1. Maximum value is 500. Default value is 500.
	SortBy      string  `url:"sortBy,omitempty"`      //A property within the response to sort by.
	Order       string  `url:"order,omitempty"`       //Whether ascending or descending order should be used to sort the response. Available values : asc, desc. Default value is asc
}
type GetCountOfNetworkBugsV1QueryParams struct {
	ID          string  `url:"id,omitempty"`          //Id of the network bug
	DeviceCount float64 `url:"deviceCount,omitempty"` //Return network bugs with deviceCount greater than this deviceCount
	Severity    string  `url:"severity,omitempty"`    //Return network bugs with this severity. Available values : CATASTROPHIC, SEVERE, MODERATE
}
type GetNetworkBugDevicesForTheBugV1QueryParams struct {
	NetworkDeviceID string  `url:"networkDeviceId,omitempty"` //Id of the network device
	ScanMode        string  `url:"scanMode,omitempty"`        //Mode or the criteria using which the network device was scanned. Available values : ESSENTIALS, ADVANTAGE, CX_CLOUD, NOT_AVAILABLE
	ScanStatus      string  `url:"scanStatus,omitempty"`      //Status of the scan on the network device. Available values : NOT_SCANNED, IN_PROGRESS, SUCCESS, FAILED, FALL_BACK
	Offset          float64 `url:"offset,omitempty"`          //The first record to show for this page; the first record is numbered 1. Default value is 1
	Limit           float64 `url:"limit,omitempty"`           //The number of records to show for this page. Minimum value is 1. Maximum value is 500. Default value is 500.
	SortBy          string  `url:"sortBy,omitempty"`          //A property within the response to sort by.
	Order           string  `url:"order,omitempty"`           //Whether ascending or descending order should be used to sort the response. Available values : asc, desc. Default value is asc
}
type GetCountOfNetworkBugDevicesForTheBugV1QueryParams struct {
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` //Id of the network device
	ScanMode        string `url:"scanMode,omitempty"`        //Mode or the criteria using which the network device was scanned. Available values : ESSENTIALS, ADVANTAGE, CX_CLOUD, NOT_AVAILABLE
	ScanStatus      string `url:"scanStatus,omitempty"`      //Status of the scan on the network device. Available values : NOT_SCANNED, IN_PROGRESS, SUCCESS, FAILED, FALL_BACK
}
type GetNetworkBugDevicesV1QueryParams struct {
	NetworkDeviceID string  `url:"networkDeviceId,omitempty"` //Id of the network device
	ScanMode        string  `url:"scanMode,omitempty"`        //Mode or the criteria using which the network device was scanned. Available values : ESSENTIALS, ADVANTAGE, CX_CLOUD, NOT_AVAILABLE
	ScanStatus      string  `url:"scanStatus,omitempty"`      //Status of the scan on the network device. Available values : NOT_SCANNED, IN_PROGRESS, SUCCESS, FAILED, FALL_BACK
	BugCount        float64 `url:"bugCount,omitempty"`        //Return network devices with bugCount greater than this bugCount
	Offset          float64 `url:"offset,omitempty"`          //The first record to show for this page; the first record is numbered 1. Default value is 1.
	Limit           float64 `url:"limit,omitempty"`           //The number of records to show for this page. Minimum value is 1. Maximum value is 500. Default value is 500.
	SortBy          string  `url:"sortBy,omitempty"`          //A property within the response to sort by.
	Order           string  `url:"order,omitempty"`           //Whether ascending or descending order should be used to sort the response. Available values : asc, desc. Default value is asc
}
type GetCountOfNetworkBugDevicesV1QueryParams struct {
	NetworkDeviceID string  `url:"networkDeviceId,omitempty"` //Id of the network device
	ScanMode        string  `url:"scanMode,omitempty"`        //Mode or the criteria using which the network device was scanned. Available values : ESSENTIALS, ADVANTAGE, CX_CLOUD, NOT_AVAILABLE
	ScanStatus      string  `url:"scanStatus,omitempty"`      //Status of the scan on the network device. Available values : NOT_SCANNED, IN_PROGRESS, SUCCESS, FAILED, FALL_BACK
	BugCount        float64 `url:"bugCount,omitempty"`        //Return network devices with bugCount greater than this bugCount
}
type GetBugsAffectingTheNetworkDeviceV1QueryParams struct {
	ID       string  `url:"id,omitempty"`       //Id of the network bug
	Severity string  `url:"severity,omitempty"` //Return network bugs with this severity. Available values : CATASTROPHIC, SEVERE, MODERATE.
	Offset   float64 `url:"offset,omitempty"`   //The first record to show for this page; the first record is numbered 1. Default value is 1.
	Limit    float64 `url:"limit,omitempty"`    //The number of records to show for this page. Minimum value is 1. Maximum value is 500. Default value is 500.
	SortBy   string  `url:"sortBy,omitempty"`   //A property within the response to sort by.
	Order    string  `url:"order,omitempty"`    //Whether ascending or descending order should be used to sort the response. Available values : asc, desc. Default value is asc.
}
type GetCountOfBugsAffectingTheNetworkDeviceV1QueryParams struct {
	ID       string `url:"id,omitempty"`       //Id of the network bug
	Severity string `url:"severity,omitempty"` //Return network bugs with this severity. Available values : CATASTROPHIC, SEVERE, MODERATE
}
type GetNetworkBugsResultsTrendOverTimeV1QueryParams struct {
	ScanTime float64 `url:"scanTime,omitempty"` //Return bugs trend with scanTime greater than this scanTime
	Offset   float64 `url:"offset,omitempty"`   //The first record to show for this page; the first record is numbered 1. Default value is 1.
	Limit    float64 `url:"limit,omitempty"`    //The number of records to show for this page. Minimum value is 1. Maximum value is 500. Default value is 500.
}
type GetCountOfNetworkBugsResultsTrendOverTimeV1QueryParams struct {
	ScanTime float64 `url:"scanTime,omitempty"` //Return bugs trend with scanTime greater than this scanTime
}
type TriggersABugsScanForTheSupportedNetworkDevicesV1QueryParams struct {
	FailedDevicesOnly bool `url:"failedDevicesOnly,omitempty"` //Used to specify if the scan should run only for the network devices that failed during the previous scan. If not specified, this parameter defaults to false.
}
type GetSecurityAdvisoriesAffectingTheNetworkDevicesV1QueryParams struct {
	ID                   string  `url:"id,omitempty"`                   //Id of the advisory
	DeviceCount          float64 `url:"deviceCount,omitempty"`          //Return advisories with deviceCount greater than this deviceCount
	CvssBaseScore        string  `url:"cvssBaseScore,omitempty"`        //Return advisories with cvssBaseScore greater than this cvssBaseScore. E.g. : 8.5
	SecurityImpactRating string  `url:"securityImpactRating,omitempty"` //Return advisories with this securityImpactRating. Available values : CRITICAL, HIGH.
	Offset               float64 `url:"offset,omitempty"`               //The first record to show for this page; the first record is numbered 1. Default value is 1.
	Limit                float64 `url:"limit,omitempty"`                //The number of records to show for this page. Minimum value is 1. Maximum value is 500. Default value is 500.
	SortBy               string  `url:"sortBy,omitempty"`               //A property within the response to sort by.
	Order                string  `url:"order,omitempty"`                //Whether ascending or descending order should be used to sort the response. Available values : asc, desc. Default value is asc.
}
type GetCountOfSecurityAdvisoriesAffectingTheNetworkDevicesV1QueryParams struct {
	ID                   string  `url:"id,omitempty"`                   //Id of the security advisory
	DeviceCount          float64 `url:"deviceCount,omitempty"`          //Return advisories with deviceCount greater than this deviceCount
	CvssBaseScore        string  `url:"cvssBaseScore,omitempty"`        //Return advisories with cvssBaseScore greater than this cvssBaseScore. E.g. : 8.5
	SecurityImpactRating string  `url:"securityImpactRating,omitempty"` //Return advisories with this securityImpactRating. Available values : CRITICAL, HIGH.
}
type GetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1QueryParams struct {
	NetworkDeviceID string  `url:"networkDeviceId,omitempty"` //Id of the network device
	ScanMode        string  `url:"scanMode,omitempty"`        //Mode or the criteria using which the network device was scanned. Available values : ESSENTIALS, ADVANTAGE, CX_CLOUD, NOT_AVAILABLE
	ScanStatus      string  `url:"scanStatus,omitempty"`      //Status of the scan on the network device. Available values : NOT_SCANNED, IN_PROGRESS, SUCCESS, FAILED, FALL_BACK.
	Offset          float64 `url:"offset,omitempty"`          //The first record to show for this page; the first record is numbered 1. Default value is 1.
	Limit           float64 `url:"limit,omitempty"`           //The number of records to show for this page. Minimum value is 1. Maximum value is 500. Default value is 500.
	SortBy          string  `url:"sortBy,omitempty"`          //A property within the response to sort by.
	Order           string  `url:"order,omitempty"`           //Whether ascending or descending order should be used to sort the response. Available values : asc, desc. Default value is asc.
}
type GetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1QueryParams struct {
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` //Id of the network device
	ScanMode        string `url:"scanMode,omitempty"`        //Mode or the criteria using which the network device was scanned. Available values : ESSENTIALS, ADVANTAGE, CX_CLOUD, NOT_AVAILABLE
	ScanStatus      string `url:"scanStatus,omitempty"`      //Status of the scan on the network device. Available values : NOT_SCANNED, IN_PROGRESS, SUCCESS, FAILED, FALL_BACK.
}
type GetSecurityAdvisoryNetworkDevicesV1QueryParams struct {
	NetworkDeviceID string  `url:"networkDeviceId,omitempty"` //Id of the network device
	ScanMode        string  `url:"scanMode,omitempty"`        //Mode or the criteria using which the network device was scanned. Available values : ESSENTIALS, ADVANTAGE, CX_CLOUD, NOT_AVAILABLE
	ScanStatus      string  `url:"scanStatus,omitempty"`      //Status of the scan on the network device. Available values : NOT_SCANNED, IN_PROGRESS, SUCCESS, FAILED, FALL_BACK.
	AdvisoryCount   string  `url:"advisoryCount,omitempty"`   //Return network devices with advisoryCount greater than this advisoryCount
	Offset          float64 `url:"offset,omitempty"`          //The first record to show for this page; the first record is numbered 1. Default value is 1.
	Limit           float64 `url:"limit,omitempty"`           //The number of records to show for this page. Minimum value is 1. Maximum value is 500. Default value is 500.
	SortBy          string  `url:"sortBy,omitempty"`          //A property within the response to sort by.
	Order           string  `url:"order,omitempty"`           //Whether ascending or descending order should be used to sort the response. Available values : asc, desc. Default value is asc.
}
type GetCountOfSecurityAdvisoryNetworkDevicesV1QueryParams struct {
	NetworkDeviceID string  `url:"networkDeviceId,omitempty"` //Id of the network device
	ScanMode        string  `url:"scanMode,omitempty"`        //Mode or the criteria using which the network device was scanned. Available values : ESSENTIALS, ADVANTAGE, CX_CLOUD, NOT_AVAILABLE
	ScanStatus      string  `url:"scanStatus,omitempty"`      //Status of the scan on the network device. Available values : NOT_SCANNED, IN_PROGRESS, SUCCESS, FAILED, FALL_BACK.
	AdvisoryCount   float64 `url:"advisoryCount,omitempty"`   //Return network devices with advisoryCount greater than this advisoryCount
}
type GetSecurityAdvisoriesAffectingTheNetworkDeviceV1QueryParams struct {
	ID                   string  `url:"id,omitempty"`                   //Id of the security advisory
	CvssBaseScore        string  `url:"cvssBaseScore,omitempty"`        //Return advisories with cvssBaseScore greater than this cvssBaseScore. E.g. : 8.5
	SecurityImpactRating string  `url:"securityImpactRating,omitempty"` //Return advisories with this securityImpactRating. Available values : CRITICAL, HIGH.
	Offset               float64 `url:"offset,omitempty"`               //The first record to show for this page; the first record is numbered 1. Default value is 1.
	Limit                float64 `url:"limit,omitempty"`                //The number of records to show for this page. Minimum value is 1. Maximum value is 500. Default value is 500.
	SortBy               string  `url:"sortBy,omitempty"`               //A property within the response to sort by.
	Order                string  `url:"order,omitempty"`                //Whether ascending or descending order should be used to sort the response. Available values : asc, desc. Default value is asc.
}
type GetCountOfSecurityAdvisoriesAffectingTheNetworkDeviceV1QueryParams struct {
	ID                   string `url:"id,omitempty"`                   //Id of the security advisory
	CvssBaseScore        string `url:"cvssBaseScore,omitempty"`        //Return advisories with cvssBaseScore greater than this cvssBaseScore. E.g. : 8.5
	SecurityImpactRating string `url:"securityImpactRating,omitempty"` //Return advisories with this securityImpactRating. Available values : CRITICAL, HIGH.
}
type GetSecurityAdvisoriesResultsTrendOverTimeV1QueryParams struct {
	ScanTime float64 `url:"scanTime,omitempty"` //Return advisories trend with scanTime greater than this scanTime
	Offset   float64 `url:"offset,omitempty"`   //The first record to show for this page; the first record is numbered 1. Default value is 1.
	Limit    float64 `url:"limit,omitempty"`    //The number of records to show for this page. Minimum value is 1. Maximum value is 500. Default value is 500.
}
type GetCountOfSecurityAdvisoriesResultsTrendOverTimeV1QueryParams struct {
	ScanTime float64 `url:"scanTime,omitempty"` //Return advisories trend with scanTime greater than this scanTime
}
type TriggersASecurityAdvisoriesScanForTheSupportedNetworkDevicesV1QueryParams struct {
	FailedDevicesOnly bool `url:"failedDevicesOnly,omitempty"` //Used to specify if the scan should run only for the network devices that failed during the previous scan. If not specified, this parameter defaults to false.
}

type ResponseComplianceGetComplianceStatusV1 struct {
	Version  string                                             `json:"version,omitempty"`  // Version of the API.
	Response *[]ResponseComplianceGetComplianceStatusV1Response `json:"response,omitempty"` //
}
type ResponseComplianceGetComplianceStatusV1Response struct {
	DeviceUUID       string   `json:"deviceUuid,omitempty"`       // UUID of the device.
	ComplianceStatus string   `json:"complianceStatus,omitempty"` // Current compliance status for the compliance type that will be one of COMPLIANT, NON_COMPLIANT, ERROR, IN_PROGRESS, NOT_APPLICABLE, NOT_AVAILABLE, COMPLIANT_WARNING, REMEDIATION_IN_PROGRESS, or ABORTED.
	Message          string   `json:"message,omitempty"`          // Additional message of compliance status for the compliance type.
	ScheduleTime     *float64 `json:"scheduleTime,omitempty"`     // Timestamp when compliance is scheduled to run.
	LastUpdateTime   *float64 `json:"lastUpdateTime,omitempty"`   // Timestamp when the latest compliance checks ran.
}
type ResponseComplianceRunComplianceV1 struct {
	Version  string                                     `json:"version,omitempty"`  // Version of the API.
	Response *ResponseComplianceRunComplianceV1Response `json:"response,omitempty"` //
}
type ResponseComplianceRunComplianceV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task id.
	URL    string `json:"url,omitempty"`    // Additional url for task id.
}
type ResponseComplianceGetComplianceStatusCountV1 struct {
	Version  string   `json:"version,omitempty"`  // Version of the API.
	Response *float64 `json:"response,omitempty"` // Returns count of compliant status
}
type ResponseComplianceGetComplianceDetailV1 struct {
	Version  string                                             `json:"version,omitempty"`  // Version of the API.
	Response *[]ResponseComplianceGetComplianceDetailV1Response `json:"response,omitempty"` //
}
type ResponseComplianceGetComplianceDetailV1Response struct {
	ComplianceType       string   `json:"complianceType,omitempty"`       // Compliance type corresponds to a tile on the UI. Will be one of NETWORK_PROFILE, IMAGE, APPLICATION_VISIBILITY, FABRIC, PSIRT, RUNNING_CONFIG, NETWORK_SETTINGS, WORKFLOW, or EoX.
	LastSyncTime         *float64 `json:"lastSyncTime,omitempty"`         // Timestamp when the status changed from different value to the current value.
	DeviceUUID           string   `json:"deviceUuid,omitempty"`           // UUID of the device.
	DisplayName          string   `json:"displayName,omitempty"`          // User friendly name for the configuration.
	Status               string   `json:"status,omitempty"`               // Current status of compliance for the complianceType. Will be one of COMPLIANT, NON_COMPLIANT, ERROR, IN_PROGRESS, NOT_APPLICABLE, NOT_AVAILABLE, COMPLIANT_WARNING, REMEDIATION_IN_PROGRESS, or ABORTED.
	Category             string   `json:"category,omitempty"`             // category can have any value among 'INTENT'(mapped to compliance types: NETWORK_SETTINGS,NETWORK_PROFILE,WORKFLOW,FABRIC,APPLICATION_VISIBILITY), 'RUNNING_CONFIG' , 'IMAGE' , 'PSIRT' , 'EoX' , 'NETWORK_SETTINGS'.
	LastUpdateTime       *float64 `json:"lastUpdateTime,omitempty"`       // Timestamp when the latest compliance checks ran.
	State                string   `json:"state,omitempty"`                // State of latest compliance check for the complianceType. Will be one of SUCCESS, FAILED, or IN_PROGRESS.
	RemediationSupported *bool    `json:"remediationSupported,omitempty"` // Indicates whether remediation is supported for this compliance type or not
}
type ResponseComplianceGetComplianceDetailCountV1 struct {
	Version  string   `json:"version,omitempty"`  // Version of API.
	Response *float64 `json:"response,omitempty"` // Count of all devices or devices that match the query parameters.
}
type ResponseComplianceComplianceRemediationV1 struct {
	Response *ResponseComplianceComplianceRemediationV1Response `json:"response,omitempty"` //
	Version  string                                             `json:"version,omitempty"`  // Version of API.
}
type ResponseComplianceComplianceRemediationV1Response struct {
	TaskID string `json:"taskId,omitempty"` // The UUID of the task.
	URL    string `json:"url,omitempty"`    // The path to the API endpoint to GET for information on the task.
}
type ResponseComplianceDeviceComplianceStatusV1 struct {
	Response *ResponseComplianceDeviceComplianceStatusV1Response `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  // Version of the API.
}
type ResponseComplianceDeviceComplianceStatusV1Response struct {
	DeviceUUID       string   `json:"deviceUuid,omitempty"`       // UUID of the device.
	ComplianceStatus string   `json:"complianceStatus,omitempty"` // Current compliance status of the device that will be one of COMPLIANT, NON_COMPLIANT, ERROR, IN_PROGRESS, NOT_APPLICABLE, NOT_AVAILABLE, COMPLIANT_WARNING, REMEDIATION_IN_PROGRESS, or ABORTED.
	LastUpdateTime   *float64 `json:"lastUpdateTime,omitempty"`   // Timestamp when the latest compliance checks ran.
	ScheduleTime     string   `json:"scheduleTime,omitempty"`     // Timestamp when the next compliance checks will run.
}
type ResponseComplianceComplianceDetailsOfDeviceV1 struct {
	Response   *[]ResponseComplianceComplianceDetailsOfDeviceV1Response `json:"response,omitempty"`   //
	DeviceUUID string                                                   `json:"deviceUuid,omitempty"` // UUID of the device.
}
type ResponseComplianceComplianceDetailsOfDeviceV1Response struct {
	DeviceUUID     string                                                                 `json:"deviceUuid,omitempty"`     // UUID of the device.
	ComplianceType string                                                                 `json:"complianceType,omitempty"` // Compliance type corresponds to a tile on the UI that will be one of NETWORK_PROFILE, IMAGE, APPLICATION_VISIBILITY, FABRIC, PSIRT, RUNNING_CONFIG, NETWORK_SETTINGS, WORKFLOW, or EoX.
	Status         string                                                                 `json:"status,omitempty"`         // Status of compliance for the compliance type, will be one of COMPLIANT, NON_COMPLIANT, ERROR, IN_PROGRESS, NOT_APPLICABLE, NOT_AVAILABLE, COMPLIANT_WARNING, REMEDIATION_IN_PROGRESS, or ABORTED.
	State          string                                                                 `json:"state,omitempty"`          // State of the compliance check for the compliance type, will be one of SUCCESS, FAILED, or IN_PROGRESS.
	LastSyncTime   *float64                                                               `json:"lastSyncTime,omitempty"`   // Timestamp when the status changed from a different value to the current value.
	LastUpdateTime *float64                                                               `json:"lastUpdateTime,omitempty"` // Timestamp of the latest compliance check that was run.
	SourceInfoList *[]ResponseComplianceComplianceDetailsOfDeviceV1ResponseSourceInfoList `json:"sourceInfoList,omitempty"` //
	AckStatus      string                                                                 `json:"ackStatus,omitempty"`      // Acknowledgment status of the compliance type. UNACKNOWLEDGED if none of the violations under the compliance type are acknowledged. Otherwise it will be ACKNOWLEDGED.
	Version        string                                                                 `json:"version,omitempty"`        // Version of the API.
}
type ResponseComplianceComplianceDetailsOfDeviceV1ResponseSourceInfoList struct {
	Name                string                                                                          `json:"name,omitempty"`                // Name of the type of top level configuration.
	NameWithBusinessKey string                                                                          `json:"nameWithBusinessKey,omitempty"` // Name With Business Key
	SourceEnum          string                                                                          `json:"sourceEnum,omitempty"`          // Will be same as compliance type.
	Type                string                                                                          `json:"type,omitempty"`                // Type of the top level configuration.
	AppName             string                                                                          `json:"appName,omitempty"`             // Application name that is used to club the violations.
	Count               *float64                                                                        `json:"count,omitempty"`               // Number of violations present.
	AckStatus           string                                                                          `json:"ackStatus,omitempty"`           // Acknowledgment status of violations. UNACKNOWLEDGED if none of the violations are acknowledged. Otherwise it will be ACKNOWLEDGED.
	BusinessKey         *ResponseComplianceComplianceDetailsOfDeviceV1ResponseSourceInfoListBusinessKey `json:"businessKey,omitempty"`         //
	DiffList            *[]ResponseComplianceComplianceDetailsOfDeviceV1ResponseSourceInfoListDiffList  `json:"diffList,omitempty"`            //
	DisplayName         string                                                                          `json:"displayName,omitempty"`         // Model display name.
}
type ResponseComplianceComplianceDetailsOfDeviceV1ResponseSourceInfoListBusinessKey struct {
	ResourceName          string                                                                                               `json:"resourceName,omitempty"`          // Name of the top level resource. Same as name above.
	BusinessKeyAttributes *ResponseComplianceComplianceDetailsOfDeviceV1ResponseSourceInfoListBusinessKeyBusinessKeyAttributes `json:"businessKeyAttributes,omitempty"` // Attributes that together uniquely identify the configuration instance.
	OtherAttributes       *ResponseComplianceComplianceDetailsOfDeviceV1ResponseSourceInfoListBusinessKeyOtherAttributes       `json:"otherAttributes,omitempty"`       //
}
type ResponseComplianceComplianceDetailsOfDeviceV1ResponseSourceInfoListBusinessKeyBusinessKeyAttributes interface{}
type ResponseComplianceComplianceDetailsOfDeviceV1ResponseSourceInfoListBusinessKeyOtherAttributes struct {
	Name          string                                                                                                      `json:"name,omitempty"`          // Name of the attributes.
	CfsAttributes *ResponseComplianceComplianceDetailsOfDeviceV1ResponseSourceInfoListBusinessKeyOtherAttributesCfsAttributes `json:"cfsAttributes,omitempty"` //
}
type ResponseComplianceComplianceDetailsOfDeviceV1ResponseSourceInfoListBusinessKeyOtherAttributesCfsAttributes struct {
	DisplayName string `json:"displayName,omitempty"` // User friendly name for the configuration.
	AppName     string `json:"appName,omitempty"`     // Same as appName above.
	Description string `json:"description,omitempty"` // Description for the configuration, if available.
	Source      string `json:"source,omitempty"`      // Will be same as compliance type.
	Type        string `json:"type,omitempty"`        // The type of this attribute (for example, type can be Intent).
}
type ResponseComplianceComplianceDetailsOfDeviceV1ResponseSourceInfoListDiffList struct {
	Op                 string                                                                                         `json:"op,omitempty"`                 // Type of change (add, remove, or update).
	ConfiguredValue    string                                                                                         `json:"configuredValue,omitempty"`    // Configured value i.e. running / current value. It will be empty for the template violations due to potentially large size of the template. Use a dedicated API to get the template data.
	IntendedValue      string                                                                                         `json:"intendedValue,omitempty"`      // Enable", Intended value. It will be empty for the template violations due to potentially large size of the template. Use a dedicated API to get the template data.
	MoveFromPath       string                                                                                         `json:"moveFromPath,omitempty"`       // Additional URI to fetch more details, if available.
	BusinessKey        string                                                                                         `json:"businessKey,omitempty"`        // The Unique key of the individual violation does not change after every compliance check, as long as the deployment data doesn't change.
	Path               string                                                                                         `json:"path,omitempty"`               // Path of the configuration relative to the top-level configuration. Use it along with a name to identify certain set of violations.
	ExtendedAttributes *ResponseComplianceComplianceDetailsOfDeviceV1ResponseSourceInfoListDiffListExtendedAttributes `json:"extendedAttributes,omitempty"` //
	AckStatus          string                                                                                         `json:"ackStatus,omitempty"`          // Acknowledgment status of the violation. ACKNOWLEDGED if the violation is acknowledged or at the top-level configuration. Otherwise it will be UNACKNOWLEDGED.
	InstanceUUID       string                                                                                         `json:"instanceUUID,omitempty"`       // UUID of the individual violation. Changes after every compliance check.
	DisplayName        string                                                                                         `json:"displayName,omitempty"`        // Display name for attribute in ui .If business key is null or of type owning entity type.
}
type ResponseComplianceComplianceDetailsOfDeviceV1ResponseSourceInfoListDiffListExtendedAttributes struct {
	AttributeDisplayName string `json:"attributeDisplayName,omitempty"` // Display name for attribute in ui .if business key is null or only owning entity type.
	Path                 string `json:"path,omitempty"`                 // Path to be displayed on the UI, instead of the above path, if available.
	DataConverter        string `json:"dataConverter,omitempty"`        // Name of the converter used to display configurations in user-friendly format, if available.
	Type                 string `json:"type,omitempty"`                 // Type of this attribute.(example type can be Intent)
}
type ResponseComplianceGetFieldNoticeNetworkDevicesV1 struct {
	Response *[]ResponseComplianceGetFieldNoticeNetworkDevicesV1Response `json:"response,omitempty"` //
	Version  string                                                      `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetFieldNoticeNetworkDevicesV1Response struct {
	NetworkDeviceID      string `json:"networkDeviceId,omitempty"`      // Id of the device
	NoticeCount          *int   `json:"noticeCount,omitempty"`          // Number of field notices to which the network device is vulnerable
	PotentialNoticeCount *int   `json:"potentialNoticeCount,omitempty"` // Number of potential field notices to which the network device is vulnerable
	ScanStatus           string `json:"scanStatus,omitempty"`           // 'NOT_SCANNED' - the device was not scanned. 'IN_PROGRESS' - a scan is in progress for the device. 'SUCCESS' - device scan was successful. 'FAILED' - device scan failed
	Comments             string `json:"comments,omitempty"`             // More details about the scan status. ie:- if the scan status is failed, comments will give the reason for failure
	LastScanTime         *int   `json:"lastScanTime,omitempty"`         // Time at which the device was scanned
}
type ResponseComplianceGetCountOfFieldNoticeNetworkDevicesV1 struct {
	Version  string                                                           `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetCountOfFieldNoticeNetworkDevicesV1Response `json:"response,omitempty"` //
}
type ResponseComplianceGetCountOfFieldNoticeNetworkDevicesV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseComplianceGetFieldNoticeNetworkDeviceByDeviceIDV1 struct {
	Response *ResponseComplianceGetFieldNoticeNetworkDeviceByDeviceIDV1Response `json:"response,omitempty"` //
	Version  string                                                             `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetFieldNoticeNetworkDeviceByDeviceIDV1Response struct {
	NetworkDeviceID      string `json:"networkDeviceId,omitempty"`      // Id of the device
	NoticeCount          *int   `json:"noticeCount,omitempty"`          // Number of field notices to which the network device is vulnerable
	PotentialNoticeCount *int   `json:"potentialNoticeCount,omitempty"` // Number of potential field notices to which the network device is vulnerable
	ScanStatus           string `json:"scanStatus,omitempty"`           // 'NOT_SCANNED' - the device was not scanned. 'IN_PROGRESS' - a scan is in progress for the device. 'SUCCESS' - device scan was successful. 'FAILED' - device scan failed
	Comments             string `json:"comments,omitempty"`             // More details about the scan status. ie:- if the scan status is failed, comments will give the reason for failure
	LastScanTime         *int   `json:"lastScanTime,omitempty"`         // Time at which the device was scanned
}
type ResponseComplianceGetFieldNoticesAffectingTheNetworkDeviceV1 struct {
	Response *[]ResponseComplianceGetFieldNoticesAffectingTheNetworkDeviceV1Response `json:"response,omitempty"` //
	Version  string                                                                  `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetFieldNoticesAffectingTheNetworkDeviceV1Response struct {
	ID                   string `json:"id,omitempty"`                   // Id of the field notice
	Name                 string `json:"name,omitempty"`                 // Name of the field notice
	PublicationURL       *int   `json:"publicationUrl,omitempty"`       // Url for getting field notice details on cisco website
	DeviceCount          *int   `json:"deviceCount,omitempty"`          // Number of devices which are vulnerable to this field notice
	PotentialDeviceCount *int   `json:"potentialDeviceCount,omitempty"` // Number of devices which are potentially vulnerable to this field notice
	Type                 string `json:"type,omitempty"`                 // 'SOFTWARE' - field notice is for the network device software. 'HARDWARE' - field notice is for the network device hardware
	FirstPublishDate     *int   `json:"firstPublishDate,omitempty"`     // Time at which the field notice was published
	LastUpdatedDate      *int   `json:"lastUpdatedDate,omitempty"`      // Time at which the field notice was last updated
	MatchConfidence      string `json:"matchConfidence,omitempty"`      // 'VULNERABLE' - network device is vulnerable to the field notice. 'POTENTIALLY_VULNERABLE' - network device is potentially vulnerable to the field notice. additional manual verifications are needed.
	MatchReason          string `json:"matchReason,omitempty"`          // If the MatchConfidence is POTENTIALLY_VULNERABLE, this gives more details such as what was matched and if additional manual verifications are needed.
	NetworkDeviceID      string `json:"networkDeviceId,omitempty"`      // Id of the device
}
type ResponseComplianceGetCountOfFieldNoticesAffectingTheNetworkDeviceV1 struct {
	Version  string                                                                       `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetCountOfFieldNoticesAffectingTheNetworkDeviceV1Response `json:"response,omitempty"` //
}
type ResponseComplianceGetCountOfFieldNoticesAffectingTheNetworkDeviceV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseComplianceGetFieldNoticeAffectingTheNetworkDeviceByDeviceIDAndNoticeIDV1 struct {
	Response *ResponseComplianceGetFieldNoticeAffectingTheNetworkDeviceByDeviceIDAndNoticeIDV1Response `json:"response,omitempty"` //
	Version  string                                                                                    `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetFieldNoticeAffectingTheNetworkDeviceByDeviceIDAndNoticeIDV1Response struct {
	ID                   string `json:"id,omitempty"`                   // Id of the field notice
	Name                 string `json:"name,omitempty"`                 // Name of the field notice
	PublicationURL       string `json:"publicationUrl,omitempty"`       // Url for getting field notice details on cisco website
	DeviceCount          *int   `json:"deviceCount,omitempty"`          // Number of devices which are vulnerable to this field notice
	PotentialDeviceCount *int   `json:"potentialDeviceCount,omitempty"` // Number of devices which are potentially vulnerable to this field notice
	Type                 string `json:"type,omitempty"`                 // 'SOFTWARE' - field notice is for the network device software. 'HARDWARE' - field notice is for the network device hardware
	FirstPublishDate     *int   `json:"firstPublishDate,omitempty"`     // Time at which the field notice was published
	LastUpdatedDate      *int   `json:"lastUpdatedDate,omitempty"`      // Time at which the field notice was last updated
	MatchConfidence      string `json:"matchConfidence,omitempty"`      // 'VULNERABLE' - network device is vulnerable to the field notice. 'POTENTIALLY_VULNERABLE' - network device is potentially vulnerable to the field notice. additional manual verifications are needed.
	MatchReason          string `json:"matchReason,omitempty"`          // If the MatchConfidence is POTENTIALLY_VULNERABLE, this gives more details such as what was matched and if additional manual verifications are needed.
	NetworkDeviceID      string `json:"networkDeviceId,omitempty"`      // Id of the device
}
type ResponseComplianceGetFieldNoticesV1 struct {
	Response *[]ResponseComplianceGetFieldNoticesV1Response `json:"response,omitempty"` //
	Version  string                                         `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetFieldNoticesV1Response struct {
	ID                   string `json:"id,omitempty"`                   // Id of the field notice
	Name                 string `json:"name,omitempty"`                 // Name of the field notice
	PublicationURL       string `json:"publicationUrl,omitempty"`       // Url for getting field notice details on cisco website
	DeviceCount          *int   `json:"deviceCount,omitempty"`          // Number of devices which are vulnerable to this field notice
	PotentialDeviceCount *int   `json:"potentialDeviceCount,omitempty"` // Number of devices which are potentially vulnerable to this field notice
	Type                 string `json:"type,omitempty"`                 // 'SOFTWARE' - field notice is for the network device software. 'HARDWARE' - field notice is for the network device hardware
	FirstPublishDate     *int   `json:"firstPublishDate,omitempty"`     // Time at which the field notice was published
	LastUpdatedDate      *int   `json:"lastUpdatedDate,omitempty"`      // Time at which the field notice was last updated
}
type ResponseComplianceGetCountOfFieldNoticesV1 struct {
	Version  string                                              `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetCountOfFieldNoticesV1Response `json:"response,omitempty"` //
}
type ResponseComplianceGetCountOfFieldNoticesV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseComplianceGetFieldNoticeByIDV1 struct {
	Response *ResponseComplianceGetFieldNoticeByIDV1Response `json:"response,omitempty"` //
	Version  string                                          `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetFieldNoticeByIDV1Response struct {
	ID                   string `json:"id,omitempty"`                   // Id of the field notice
	Name                 string `json:"name,omitempty"`                 // Name of the field notice
	PublicationURL       string `json:"publicationUrl,omitempty"`       // Url for getting field notice details on cisco website
	DeviceCount          *int   `json:"deviceCount,omitempty"`          // Number of devices which are vulnerable to this field notice
	PotentialDeviceCount *int   `json:"potentialDeviceCount,omitempty"` // Number of devices which are potentially vulnerable to this field notice
	Type                 string `json:"type,omitempty"`                 // 'SOFTWARE' - field notice is for the network device software. 'HARDWARE' - field notice is for the network device hardware
	FirstPublishDate     *int   `json:"firstPublishDate,omitempty"`     // Time at which the field notice was published
	LastUpdatedDate      *int   `json:"lastUpdatedDate,omitempty"`      // Time at which the field notice was last updated
}
type ResponseComplianceGetFieldNoticeNetworkDevicesForTheNoticeV1 struct {
	Response *[]ResponseComplianceGetFieldNoticeNetworkDevicesForTheNoticeV1Response `json:"response,omitempty"` //
	Version  string                                                                  `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetFieldNoticeNetworkDevicesForTheNoticeV1Response struct {
	NetworkDeviceID      string `json:"networkDeviceId,omitempty"`      // Id of the device
	NoticeCount          *int   `json:"noticeCount,omitempty"`          // Number of field notices to which the network device is vulnerable
	PotentialNoticeCount *int   `json:"potentialNoticeCount,omitempty"` // Number of potential field notices to which the network device is vulnerable
	ScanStatus           string `json:"scanStatus,omitempty"`           // 'NOT_SCANNED' - the device was not scanned. 'IN_PROGRESS' - a scan is in progress for the device. 'SUCCESS' - device scan was successful. 'FAILED' - device scan failed
	Comments             string `json:"comments,omitempty"`             // More details about the scan status. ie:- if the scan status is failed, comments will give the reason for failure
	LastScanTime         *int   `json:"lastScanTime,omitempty"`         // Time at which the device was scanned
}
type ResponseComplianceGetCountOfFieldNoticeNetworkDevicesForTheNoticeV1 struct {
	Version  string                                                                       `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetCountOfFieldNoticeNetworkDevicesForTheNoticeV1Response `json:"response,omitempty"` //
}
type ResponseComplianceGetCountOfFieldNoticeNetworkDevicesForTheNoticeV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseComplianceGetFieldNoticeNetworkDeviceForTheNoticeByNetworkDeviceIDV1 struct {
	Response *ResponseComplianceGetFieldNoticeNetworkDeviceForTheNoticeByNetworkDeviceIDV1Response `json:"response,omitempty"` //
	Version  string                                                                                `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetFieldNoticeNetworkDeviceForTheNoticeByNetworkDeviceIDV1Response struct {
	NetworkDeviceID      string `json:"networkDeviceId,omitempty"`      // Id of the device
	NoticeCount          *int   `json:"noticeCount,omitempty"`          // Number of field notices to which the network device is vulnerable
	PotentialNoticeCount *int   `json:"potentialNoticeCount,omitempty"` // Number of potential field notices to which the network device is vulnerable
	ScanStatus           string `json:"scanStatus,omitempty"`           // 'NOT_SCANNED' - the device was not scanned. 'IN_PROGRESS' - a scan is in progress for the device. 'SUCCESS' - device scan was successful. 'FAILED' - device scan failed
	Comments             string `json:"comments,omitempty"`             // More details about the scan status. ie:- if the scan status is failed, comments will give the reason for failure
	LastScanTime         *int   `json:"lastScanTime,omitempty"`         // Time at which the device was scanned
}
type ResponseComplianceGetFieldNoticesResultsTrendOverTimeV1 struct {
	Response *[]ResponseComplianceGetFieldNoticesResultsTrendOverTimeV1Response `json:"response,omitempty"` //
	Version  string                                                             `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetFieldNoticesResultsTrendOverTimeV1Response struct {
	ScanTime                  *int `json:"scanTime,omitempty"`                  // End time for the scan
	SoftwareFieldNoticesCount *int `json:"softwareFieldNoticesCount,omitempty"` // Number of field notices of type SOFTWARE
	HardwareFieldNoticesCount *int `json:"hardwareFieldNoticesCount,omitempty"` // Number of field notices of type HARDWARE
}
type ResponseComplianceGetCountOfFieldNoticesResultsTrendOverTimeV1 struct {
	Version  string                                                                  `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetCountOfFieldNoticesResultsTrendOverTimeV1Response `json:"response,omitempty"` //
}
type ResponseComplianceGetCountOfFieldNoticesResultsTrendOverTimeV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseComplianceCreatesATrialForFieldNoticesDetectionOnNetworkDevicesV1 struct {
	Version  string                                                                             `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseComplianceCreatesATrialForFieldNoticesDetectionOnNetworkDevicesV1Response `json:"response,omitempty"` //
}
type ResponseComplianceCreatesATrialForFieldNoticesDetectionOnNetworkDevicesV1Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseComplianceGetTrialDetailsForFieldNoticesDetectionOnNetworkDevicesV1 struct {
	Version  string                                                                               `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetTrialDetailsForFieldNoticesDetectionOnNetworkDevicesV1Response `json:"response,omitempty"` //
}
type ResponseComplianceGetTrialDetailsForFieldNoticesDetectionOnNetworkDevicesV1Response struct {
	Type                     string `json:"type,omitempty"`                     // Type of trial: * 'feature - the trial is of type feature. this is the currently supported type. * 'contract' - the trial is of type contract. this was used in older versions and exists only for compatibility.
	Feature                  string `json:"feature,omitempty"`                  // Name of the feature for which trial was created. for older versions that created contract type trials, this field will be absent.
	ContractLevel            string `json:"contractLevel,omitempty"`            // Contract level for which trial was created. this was used in older versions and exists only for compatibility.
	Active                   *bool  `json:"active,omitempty"`                   // Indicates if the trial is active
	StartTime                *int   `json:"startTime,omitempty"`                // Trial start time; as measured in Unix epoch time in milliseconds
	EndTime                  *int   `json:"endTime,omitempty"`                  // Trial end time; as measured in Unix epoch time in milliseconds
	SecondsRemainingToExpiry *int   `json:"secondsRemainingToExpiry,omitempty"` // Seconds remaining in the trial before it expires. for expired trials this will be 0.
	SecondsSinceExpired      *int   `json:"secondsSinceExpired,omitempty"`      // Seconds elapsed after the trial has expired. for active trials this will be 0.
}
type ResponseComplianceTriggersAFieldNoticesScanForTheSupportedNetworkDevicesV1 struct {
	Version  string                                                                              `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseComplianceTriggersAFieldNoticesScanForTheSupportedNetworkDevicesV1Response `json:"response,omitempty"` //
}
type ResponseComplianceTriggersAFieldNoticesScanForTheSupportedNetworkDevicesV1Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseComplianceGetConfigTaskDetailsV1 struct {
	Version  string                                              `json:"version,omitempty"`  // Version of the API.
	Response *[]ResponseComplianceGetConfigTaskDetailsV1Response `json:"response,omitempty"` //
}
type ResponseComplianceGetConfigTaskDetailsV1Response struct {
	StartTime       *int   `json:"startTime,omitempty"`       // Timestamp when the task started.
	ErrorCode       string `json:"errorCode,omitempty"`       // Error code if the task failed.
	DeviceID        string `json:"deviceId,omitempty"`        // UUID of the device.
	TaskID          string `json:"taskId,omitempty"`          // UUID of the task.
	TaskStatus      string `json:"taskStatus,omitempty"`      // Status of the task.
	ParentTaskID    string `json:"parentTaskId,omitempty"`    // UUID of the parent task.
	DeviceIPAddress string `json:"deviceIpAddress,omitempty"` // IP address of the device.
	DetailMessage   string `json:"detailMessage,omitempty"`   // Details of the task, if available.
	FailureMessage  string `json:"failureMessage,omitempty"`  // Failure message, if the task failed.
	TaskType        string `json:"taskType,omitempty"`        // Task type can be 0,1,2 etc(ARCHIVE_RUNNING(0),ARCHIVE_STARTUP(1),ARCHIVE_VLAN(2),DEPLOY_RUNNING(3),DEPLOY_STARTUP(4),DEPLOY_VLAN(5),COPY_RUNNING_TO_STARTUP(6))
	CompletionTime  *int   `json:"completionTime,omitempty"`  // Timestamp when the task was completed.
	HostName        string `json:"hostName,omitempty"`        // Host name of the device.
}
type ResponseComplianceCommitDeviceConfigurationV1 struct {
	Version  string                                                 `json:"version,omitempty"`  // Version of the API.
	Response *ResponseComplianceCommitDeviceConfigurationV1Response `json:"response,omitempty"` //
}
type ResponseComplianceCommitDeviceConfigurationV1Response struct {
	URL    string `json:"url,omitempty"`    // Task Id url.
	TaskID string `json:"taskId,omitempty"` // Unique Id of task.
}
type ResponseComplianceGetNetworkBugsV1 struct {
	Response *[]ResponseComplianceGetNetworkBugsV1Response `json:"response,omitempty"` //
	Version  string                                        `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetNetworkBugsV1Response struct {
	ID                 string   `json:"id,omitempty"`                 // Id of the network bug
	Headline           string   `json:"headline,omitempty"`           // Title of the network bug
	PublicationURL     string   `json:"publicationUrl,omitempty"`     // Url for getting network bug details on cisco website
	DeviceCount        *int     `json:"deviceCount,omitempty"`        // Number of devices which are vulnerable to this network bug
	Severity           string   `json:"severity,omitempty"`           // 'CATASTROPHIC' - Reasonably common circumstances cause the entire system to fail, or a major subsystem to stop working. 'SEVERE' - Important functions are unusable. 'MODERATE' - Failures occur in unusual circumstances, or minor features do not work at all.
	HasWorkaround      *bool    `json:"hasWorkaround,omitempty"`      // Indicates if the network bug has a workaround
	Workaround         string   `json:"workaround,omitempty"`         // Workaround if any that exists for the network bug
	AffectedVersions   []string `json:"affectedVersions,omitempty"`   // Versions that are affected by the network bug
	IntegratedReleases []string `json:"integratedReleases,omitempty"` // Versions that have the fix for the network bug
}
type ResponseComplianceGetCountOfNetworkBugsV1 struct {
	Version  string                                             `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetCountOfNetworkBugsV1Response `json:"response,omitempty"` //
}
type ResponseComplianceGetCountOfNetworkBugsV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseComplianceGetNetworkBugByIDV1 struct {
	Version  string                                         `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetNetworkBugByIDV1Response `json:"response,omitempty"` //
}
type ResponseComplianceGetNetworkBugByIDV1Response struct {
	ID                 string   `json:"id,omitempty"`                 // Id of the network bug
	Headline           string   `json:"headline,omitempty"`           // Title of the network bug
	PublicationURL     string   `json:"publicationUrl,omitempty"`     // Url for getting network bug details on cisco website
	DeviceCount        *int     `json:"deviceCount,omitempty"`        // Number of devices which are vulnerable to this network bug
	Severity           string   `json:"severity,omitempty"`           // 'CATASTROPHIC' - Reasonably common circumstances cause the entire system to fail, or a major subsystem to stop working. 'SEVERE' - Important functions are unusable. 'MODERATE' - Failures occur in unusual circumstances, or minor features do not work at all.
	HasWorkaround      *bool    `json:"hasWorkaround,omitempty"`      // Indicates if the network bug has a workaround
	Workaround         string   `json:"workaround,omitempty"`         // Workaround if any that exists for the network bug
	AffectedVersions   []string `json:"affectedVersions,omitempty"`   // Versions that are affected by the network bug
	IntegratedReleases []string `json:"integratedReleases,omitempty"` // Versions that have the fix for the network bug
}
type ResponseComplianceGetNetworkBugDevicesForTheBugV1 struct {
	Response *[]ResponseComplianceGetNetworkBugDevicesForTheBugV1Response `json:"response,omitempty"` //
	Version  string                                                       `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetNetworkBugDevicesForTheBugV1Response struct {
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Id of the device
	BugCount        *int   `json:"bugCount,omitempty"`        // Number of bugs to which the network device is vulnerable
	ScanMode        string `json:"scanMode,omitempty"`        // 'ESSENTIALS' - the device was scanned using a version based match criteria. 'CX_CLOUD' - the device was scanned using CX cloud engine which uses advanced matching criteria which eliminates false positives. 'ADVANTAGE' - the device was scanned using a version based match and user provided config match criteria. ‘NOT_AVAILABLE’ - scan mode is not available. e.g. when the device is not scanned
	ScanStatus      string `json:"scanStatus,omitempty"`      // 'NOT_SCANNED' - the device was not scanned. 'IN_PROGRESS' - a scan is in progress for the device. 'SUCCESS' - device scan was successful. 'FAILED' - device scan failed. 'FALL_BACK' - the device was supposed to be scanned using CX_CLOUD but because of connectivity issues fell back to a ESSENTIALS scan
	Comments        string `json:"comments,omitempty"`        // More details about the scan status. ie:- if the scan status is failed, comments will give the reason for failure
	LastScanTime    *int   `json:"lastScanTime,omitempty"`    // Time at which the device was scanned
}
type ResponseComplianceGetCountOfNetworkBugDevicesForTheBugV1 struct {
	Version  string                                                            `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetCountOfNetworkBugDevicesForTheBugV1Response `json:"response,omitempty"` //
}
type ResponseComplianceGetCountOfNetworkBugDevicesForTheBugV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseComplianceGetNetworkBugDeviceForTheBugByNetworkDeviceIDV1 []ResponseItemComplianceGetNetworkBugDeviceForTheBugByNetworkDeviceIDV1 // Array of ResponseComplianceGetNetworkBugDeviceForTheBugByNetworkDeviceIdV1
type ResponseItemComplianceGetNetworkBugDeviceForTheBugByNetworkDeviceIDV1 struct {
	Response *ResponseItemComplianceGetNetworkBugDeviceForTheBugByNetworkDeviceIDV1Response `json:"response,omitempty"` //
	Version  string                                                                         `json:"version,omitempty"`  // Version
}
type ResponseItemComplianceGetNetworkBugDeviceForTheBugByNetworkDeviceIDV1Response struct {
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Id of the device
	BugCount        *int   `json:"bugCount,omitempty"`        // Number of bugs to which the network device is vulnerable
	ScanMode        string `json:"scanMode,omitempty"`        // 'ESSENTIALS' - the device was scanned using a version based match criteria. 'CX_CLOUD' - the device was scanned using CX cloud engine which uses advanced matching criteria which eliminates false positives. 'ADVANTAGE' - the device was scanned using a version based match and user provided config match criteria. ‘NOT_AVAILABLE’ - scan mode is not available. e.g. when the device is not scanned
	ScanStatus      string `json:"scanStatus,omitempty"`      // 'NOT_SCANNED' - the device was not scanned. 'IN_PROGRESS' - a scan is in progress for the device. 'SUCCESS' - device scan was successful. 'FAILED' - device scan failed. 'FALL_BACK' - the device was supposed to be scanned using CX_CLOUD but because of connectivity issues fell back to a ESSENTIALS scan
	Comments        string `json:"comments,omitempty"`        // More details about the scan status. ie:- if the scan status is failed, comments will give the reason for failure
	LastScanTime    *int   `json:"lastScanTime,omitempty"`    // Time at which the device was scanned
}
type ResponseComplianceGetNetworkBugDevicesV1 struct {
	Response *[]ResponseComplianceGetNetworkBugDevicesV1Response `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetNetworkBugDevicesV1Response struct {
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Id of the device
	BugCount        *int   `json:"bugCount,omitempty"`        // Number of bugs to which the network device is vulnerable
	ScanMode        string `json:"scanMode,omitempty"`        // 'ESSENTIALS' - the device was scanned using a version based match criteria. 'CX_CLOUD' - the device was scanned using CX cloud engine which uses advanced matching criteria which eliminates false positives. 'ADVANTAGE' - the device was scanned using a version based match and user provided config match criteria. ‘NOT_AVAILABLE’ - scan mode is not available. e.g. when the device is not scanned
	ScanStatus      string `json:"scanStatus,omitempty"`      // 'NOT_SCANNED' - the device was not scanned. 'IN_PROGRESS' - a scan is in progress for the device. 'SUCCESS' - device scan was successful. 'FAILED' - device scan failed. 'FALL_BACK' - the device was supposed to be scanned using CX_CLOUD but because of connectivity issues fell back to a ESSENTIALS scan
	Comments        string `json:"comments,omitempty"`        // More details about the scan status. ie:- if the scan status is failed, comments will give the reason for failure
	LastScanTime    *int   `json:"lastScanTime,omitempty"`    // Time at which the device was scanned
}
type ResponseComplianceGetCountOfNetworkBugDevicesV1 struct {
	Version  string                                                   `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetCountOfNetworkBugDevicesV1Response `json:"response,omitempty"` //
}
type ResponseComplianceGetCountOfNetworkBugDevicesV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseComplianceGetNetworkBugDeviceByDeviceIDV1 []ResponseItemComplianceGetNetworkBugDeviceByDeviceIDV1 // Array of ResponseComplianceGetNetworkBugDeviceByDeviceIdV1
type ResponseItemComplianceGetNetworkBugDeviceByDeviceIDV1 struct {
	Response *ResponseItemComplianceGetNetworkBugDeviceByDeviceIDV1Response `json:"response,omitempty"` //
	Version  string                                                         `json:"version,omitempty"`  // Version
}
type ResponseItemComplianceGetNetworkBugDeviceByDeviceIDV1Response struct {
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Id of the device
	BugCount        *int   `json:"bugCount,omitempty"`        // Number of bugs to which the network device is vulnerable
	ScanMode        string `json:"scanMode,omitempty"`        // 'ESSENTIALS' - the device was scanned using a version based match criteria. 'CX_CLOUD' - the device was scanned using CX cloud engine which uses advanced matching criteria which eliminates false positives. 'ADVANTAGE' - the device was scanned using a version based match and user provided config match criteria. ‘NOT_AVAILABLE’ - scan mode is not available. e.g. when the device is not scanned
	ScanStatus      string `json:"scanStatus,omitempty"`      // 'NOT_SCANNED' - the device was not scanned. 'IN_PROGRESS' - a scan is in progress for the device. 'SUCCESS' - device scan was successful. 'FAILED' - device scan failed. 'FALL_BACK' - the device was supposed to be scanned using CX_CLOUD but because of connectivity issues fell back to a ESSENTIALS scan
	Comments        string `json:"comments,omitempty"`        // More details about the scan status. ie:- if the scan status is failed, comments will give the reason for failure
	LastScanTime    *int   `json:"lastScanTime,omitempty"`    // Time at which the device was scanned
}
type ResponseComplianceGetBugsAffectingTheNetworkDeviceV1 struct {
	Response *[]ResponseComplianceGetBugsAffectingTheNetworkDeviceV1Response `json:"response,omitempty"` //
	Version  string                                                          `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetBugsAffectingTheNetworkDeviceV1Response struct {
	ID                 string   `json:"id,omitempty"`                 // Id of the network bug
	Headline           string   `json:"headline,omitempty"`           // Title of the network bug
	PublicationURL     string   `json:"publicationUrl,omitempty"`     // Url for getting network bug details on cisco website
	DeviceCount        *int     `json:"deviceCount,omitempty"`        // Number of devices which are vulnerable to this network bug
	Severity           string   `json:"severity,omitempty"`           // 'CATASTROPHIC' - Reasonably common circumstances cause the entire system to fail, or a major subsystem to stop working. 'SEVERE' - Important functions are unusable. 'MODERATE' - Failures occur in unusual circumstances, or minor features do not work at all.
	HasWorkaround      *bool    `json:"hasWorkaround,omitempty"`      // Indicates if the network bug has a workaround
	Workaround         string   `json:"workaround,omitempty"`         // Workaround if any that exists for the network bug
	AffectedVersions   []string `json:"affectedVersions,omitempty"`   // Versions that are affected by the network bug
	IntegratedReleases []string `json:"integratedReleases,omitempty"` // Versions that have the fix for the network bug
}
type ResponseComplianceGetCountOfBugsAffectingTheNetworkDeviceV1 struct {
	Version  string                                                               `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetCountOfBugsAffectingTheNetworkDeviceV1Response `json:"response,omitempty"` //
}
type ResponseComplianceGetCountOfBugsAffectingTheNetworkDeviceV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseComplianceGetBugAffectingTheNetworkDeviceByDeviceIDAndBugIDV1 struct {
	Version  string                                                                         `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetBugAffectingTheNetworkDeviceByDeviceIDAndBugIDV1Response `json:"response,omitempty"` //
}
type ResponseComplianceGetBugAffectingTheNetworkDeviceByDeviceIDAndBugIDV1Response struct {
	ID                 string   `json:"id,omitempty"`                 // Id of the network bug
	Headline           string   `json:"headline,omitempty"`           // Title of the network bug
	PublicationURL     string   `json:"publicationUrl,omitempty"`     // Url for getting network bug details on cisco website
	DeviceCount        *int     `json:"deviceCount,omitempty"`        // Number of devices which are vulnerable to this network bug
	Severity           string   `json:"severity,omitempty"`           // 'CATASTROPHIC' - Reasonably common circumstances cause the entire system to fail, or a major subsystem to stop working. 'SEVERE' - Important functions are unusable. 'MODERATE' - Failures occur in unusual circumstances, or minor features do not work at all.
	HasWorkaround      *bool    `json:"hasWorkaround,omitempty"`      // Indicates if the network bug has a workaround
	Workaround         string   `json:"workaround,omitempty"`         // Workaround if any that exists for the network bug
	AffectedVersions   []string `json:"affectedVersions,omitempty"`   // Versions that are affected by the network bug
	IntegratedReleases []string `json:"integratedReleases,omitempty"` // Versions that have the fix for the network bug
}
type ResponseComplianceGetNetworkBugsResultsTrendOverTimeV1 struct {
	Response *[]ResponseComplianceGetNetworkBugsResultsTrendOverTimeV1Response `json:"response,omitempty"` //
	Version  string                                                            `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetNetworkBugsResultsTrendOverTimeV1Response struct {
	CatastrophicBugsCount *int `json:"catastrophicBugsCount,omitempty"` // Number of network bugs which have a severity of CATASTROPHIC
	SevereBugsCount       *int `json:"severeBugsCount,omitempty"`       // Number of network bugs which have a severity of SEVERE
	ModerateBugsCount     *int `json:"moderateBugsCount,omitempty"`     // Number of network bugs which have a severity of MODERATE
	ScanTime              *int `json:"scanTime,omitempty"`              // End time for the scan
}
type ResponseComplianceGetCountOfNetworkBugsResultsTrendOverTimeV1 struct {
	Version  string                                                                 `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetCountOfNetworkBugsResultsTrendOverTimeV1Response `json:"response,omitempty"` //
}
type ResponseComplianceGetCountOfNetworkBugsResultsTrendOverTimeV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseComplianceCreatesATrialForBugsDetectionOnNetworkDevicesV1 struct {
	Version  string                                                                     `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseComplianceCreatesATrialForBugsDetectionOnNetworkDevicesV1Response `json:"response,omitempty"` //
}
type ResponseComplianceCreatesATrialForBugsDetectionOnNetworkDevicesV1Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseComplianceGetTrialDetailsForBugsDetectionOnNetworkDevicesV1 struct {
	Version  string                                                                       `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetTrialDetailsForBugsDetectionOnNetworkDevicesV1Response `json:"response,omitempty"` //
}
type ResponseComplianceGetTrialDetailsForBugsDetectionOnNetworkDevicesV1Response struct {
	Type                     string `json:"type,omitempty"`                     // Type of trial: * 'feature - the trial is of type feature. this is the currently supported type. * 'contract' - the trial is of type contract. this was used in older versions and exists only for compatibility.
	Feature                  string `json:"feature,omitempty"`                  // Name of the feature for which trial was created. for older versions that created contract type trials, this field will be absent.
	ContractLevel            string `json:"contractLevel,omitempty"`            // Contract level for which trial was created. this was used in older versions and exists only for compatibility.
	Active                   *bool  `json:"active,omitempty"`                   // Indicates if the trial is active
	StartTime                *int   `json:"startTime,omitempty"`                // Trial start time; as measured in Unix epoch time in milliseconds
	EndTime                  *int   `json:"endTime,omitempty"`                  // Trial end time; as measured in Unix epoch time in milliseconds
	SecondsRemainingToExpiry *int   `json:"secondsRemainingToExpiry,omitempty"` // Seconds remaining in the trial before it expires. for expired trials this will be 0.
	SecondsSinceExpired      *int   `json:"secondsSinceExpired,omitempty"`      // Seconds elapsed after the trial has expired. for active trials this will be 0.
}
type ResponseComplianceTriggersABugsScanForTheSupportedNetworkDevicesV1 struct {
	Version  string                                                                      `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseComplianceTriggersABugsScanForTheSupportedNetworkDevicesV1Response `json:"response,omitempty"` //
}
type ResponseComplianceTriggersABugsScanForTheSupportedNetworkDevicesV1Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseComplianceGetSecurityAdvisoriesAffectingTheNetworkDevicesV1 struct {
	Response *[]ResponseComplianceGetSecurityAdvisoriesAffectingTheNetworkDevicesV1Response `json:"response,omitempty"` //
	Version  string                                                                         `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetSecurityAdvisoriesAffectingTheNetworkDevicesV1Response struct {
	ID                     string                                                                                               `json:"id,omitempty"`                     // Id of the security advisory
	DeviceCount            *int                                                                                                 `json:"deviceCount,omitempty"`            // Number of devices which are vulnerable to this advisory
	CveIDs                 []string                                                                                             `json:"cveIds,omitempty"`                 // CVE (Common Vulnerabilities and Exposures) ID of the advisory
	PublicationURL         string                                                                                               `json:"publicationUrl,omitempty"`         // Url for getting advisory details on cisco website
	CvssBaseScore          *float64                                                                                             `json:"cvssBaseScore,omitempty"`          // Common Vulnerability Scoring System(CVSS) base score
	SecurityImpactRating   string                                                                                               `json:"securityImpactRating,omitempty"`   // 'CRITICAL' - the advisory requires immediate mitigation. 'HIGH' - the advisory requires priority mitigation
	FirstFixedVersionsList *[]ResponseComplianceGetSecurityAdvisoriesAffectingTheNetworkDevicesV1ResponseFirstFixedVersionsList `json:"firstFixedVersionsList,omitempty"` //
}
type ResponseComplianceGetSecurityAdvisoriesAffectingTheNetworkDevicesV1ResponseFirstFixedVersionsList struct {
	VulnerableVersion string   `json:"vulnerableVersion,omitempty"` // Version that is vulnerable to the advisory
	FixedVersions     []string `json:"fixedVersions,omitempty"`     // First versions that have the fix for the advisory
}
type ResponseComplianceGetCountOfSecurityAdvisoriesAffectingTheNetworkDevicesV1 struct {
	Version  string                                                                              `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetCountOfSecurityAdvisoriesAffectingTheNetworkDevicesV1Response `json:"response,omitempty"` //
}
type ResponseComplianceGetCountOfSecurityAdvisoriesAffectingTheNetworkDevicesV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseComplianceGetSecurityAdvisoryAffectingTheNetworkDevicesByIDV1 struct {
	Response *ResponseComplianceGetSecurityAdvisoryAffectingTheNetworkDevicesByIDV1Response `json:"response,omitempty"` //
	Version  string                                                                         `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetSecurityAdvisoryAffectingTheNetworkDevicesByIDV1Response struct {
	ID                     string                                                                                                 `json:"id,omitempty"`                     // Id of the security advisory
	DeviceCount            *int                                                                                                   `json:"deviceCount,omitempty"`            // Number of devices which are vulnerable to this advisory
	CveIDs                 []string                                                                                               `json:"cveIds,omitempty"`                 // CVE (Common Vulnerabilities and Exposures) ID of the advisory
	PublicationURL         string                                                                                                 `json:"publicationUrl,omitempty"`         // Url for getting advisory details on cisco website
	CvssBaseScore          *float64                                                                                               `json:"cvssBaseScore,omitempty"`          // Common Vulnerability Scoring System(CVSS) base score
	SecurityImpactRating   string                                                                                                 `json:"securityImpactRating,omitempty"`   // 'CRITICAL' - the advisory requires immediate mitigation. 'HIGH' - the advisory requires priority mitigation
	FirstFixedVersionsList *[]ResponseComplianceGetSecurityAdvisoryAffectingTheNetworkDevicesByIDV1ResponseFirstFixedVersionsList `json:"firstFixedVersionsList,omitempty"` //
}
type ResponseComplianceGetSecurityAdvisoryAffectingTheNetworkDevicesByIDV1ResponseFirstFixedVersionsList struct {
	VulnerableVersion string   `json:"vulnerableVersion,omitempty"` // Version that is vulnerable to the advisory
	FixedVersions     []string `json:"fixedVersions,omitempty"`     // First versions that have the fix for the advisory
}
type ResponseComplianceGetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1 struct {
	Response *[]ResponseComplianceGetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1Response `json:"response,omitempty"` //
	Version  string                                                                                 `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1Response struct {
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Id of the device
	AdvisoryCount   *int   `json:"advisoryCount,omitempty"`   // Number of advisories to which the network device is vulnerable
	ScanMode        string `json:"scanMode,omitempty"`        // 'ESSENTIALS' - the device was scanned using a version based match criteria. 'ADVANTAGE' - the device was scanned using a version based match and user provided config match criteria. 'CX_CLOUD' - the device was scanned using CX cloud engine which uses advanced matching criteria which eliminates false positives.  ‘NOT_AVAILABLE’ - scan mode is not available. e.g. when the device is not scanned
	ScanStatus      string `json:"scanStatus,omitempty"`      // 'NOT_SCANNED' - the device was not scanned. 'IN_PROGRESS' - a scan is in progress for the device. 'SUCCESS' - device scan was successful. 'FAILED' - device scan failed. 'FALL_BACK' - the device was supposed to be scanned using CX_CLOUD but because of connectivity issues fell back to a ESSENTIALS or ADVANTAGE scan
	Comments        string `json:"comments,omitempty"`        // More details about the scan status. ie:- if the scan status is failed, comments will give the reason for failure
	LastScanTime    *int   `json:"lastScanTime,omitempty"`    // Time at which the device was scanned
}
type ResponseComplianceGetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1 struct {
	Version  string                                                                                      `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1Response `json:"response,omitempty"` //
}
type ResponseComplianceGetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseComplianceGetSecurityAdvisoryNetworkDeviceForTheSecurityAdvisoryByNetworkDeviceIDV1 struct {
	Response *ResponseComplianceGetSecurityAdvisoryNetworkDeviceForTheSecurityAdvisoryByNetworkDeviceIDV1Response `json:"response,omitempty"` //
	Version  string                                                                                               `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetSecurityAdvisoryNetworkDeviceForTheSecurityAdvisoryByNetworkDeviceIDV1Response struct {
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Id of the device
	AdvisoryCount   *int   `json:"advisoryCount,omitempty"`   // Number of advisories to which the network device is vulnerable
	ScanMode        string `json:"scanMode,omitempty"`        // 'ESSENTIALS' - the device was scanned using a version based match criteria. 'ADVANTAGE' - the device was scanned using a version based match and user provided config match criteria. 'CX_CLOUD' - the device was scanned using CX cloud engine which uses advanced matching criteria which eliminates false positives.  ‘NOT_AVAILABLE’ - scan mode is not available. e.g. when the device is not scanned
	ScanStatus      string `json:"scanStatus,omitempty"`      // 'NOT_SCANNED' - the device was not scanned. 'IN_PROGRESS' - a scan is in progress for the device. 'SUCCESS' - device scan was successful. 'FAILED' - device scan failed. 'FALL_BACK' - the device was supposed to be scanned using CX_CLOUD but because of connectivity issues fell back to a ESSENTIALS or ADVANTAGE scan
	Comments        string `json:"comments,omitempty"`        // More details about the scan status. ie:- if the scan status is failed, comments will give the reason for failure
	LastScanTime    *int   `json:"lastScanTime,omitempty"`    // Time at which the device was scanned
}
type ResponseComplianceGetSecurityAdvisoryNetworkDevicesV1 struct {
	Response *[]ResponseComplianceGetSecurityAdvisoryNetworkDevicesV1Response `json:"response,omitempty"` //
	Version  string                                                           `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetSecurityAdvisoryNetworkDevicesV1Response struct {
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Id of the device
	AdvisoryCount   *int   `json:"advisoryCount,omitempty"`   // Number of advisories to which the network device is vulnerable
	ScanMode        string `json:"scanMode,omitempty"`        // 'ESSENTIALS' - the device was scanned using a version based match criteria. 'ADVANTAGE' - the device was scanned using a version based match and user provided config match criteria. 'CX_CLOUD' - the device was scanned using CX cloud engine which uses advanced matching criteria which eliminates false positives.  ‘NOT_AVAILABLE’ - scan mode is not available. e.g. when the device is not scanned
	ScanStatus      string `json:"scanStatus,omitempty"`      // 'NOT_SCANNED' - the device was not scanned. 'IN_PROGRESS' - a scan is in progress for the device. 'SUCCESS' - device scan was successful. 'FAILED' - device scan failed. 'FALL_BACK' - the device was supposed to be scanned using CX_CLOUD but because of connectivity issues fell back to a ESSENTIALS or ADVANTAGE scan
	Comments        string `json:"comments,omitempty"`        // More details about the scan status. ie:- if the scan status is failed, comments will give the reason for failure
	LastScanTime    *int   `json:"lastScanTime,omitempty"`    // Time at which the device was scanned
}
type ResponseComplianceGetCountOfSecurityAdvisoryNetworkDevicesV1 struct {
	Version  string                                                                `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetCountOfSecurityAdvisoryNetworkDevicesV1Response `json:"response,omitempty"` //
}
type ResponseComplianceGetCountOfSecurityAdvisoryNetworkDevicesV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseComplianceGetSecurityAdvisoryNetworkDeviceByNetworkDeviceIDV1 struct {
	Response *ResponseComplianceGetSecurityAdvisoryNetworkDeviceByNetworkDeviceIDV1Response `json:"response,omitempty"` //
	Version  string                                                                         `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetSecurityAdvisoryNetworkDeviceByNetworkDeviceIDV1Response struct {
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Id of the device
	AdvisoryCount   *int   `json:"advisoryCount,omitempty"`   // Number of advisories to which the network device is vulnerable
	ScanMode        string `json:"scanMode,omitempty"`        // 'ESSENTIALS' - the device was scanned using a version based match criteria. 'ADVANTAGE' - the device was scanned using a version based match and user provided config match criteria. 'CX_CLOUD' - the device was scanned using CX cloud engine which uses advanced matching criteria which eliminates false positives.  ‘NOT_AVAILABLE’ - scan mode is not available. e.g. when the device is not scanned
	ScanStatus      string `json:"scanStatus,omitempty"`      // 'NOT_SCANNED' - the device was not scanned. 'IN_PROGRESS' - a scan is in progress for the device. 'SUCCESS' - device scan was successful. 'FAILED' - device scan failed. 'FALL_BACK' - the device was supposed to be scanned using CX_CLOUD but because of connectivity issues fell back to a ESSENTIALS or ADVANTAGE scan
	Comments        string `json:"comments,omitempty"`        // More details about the scan status. ie:- if the scan status is failed, comments will give the reason for failure
	LastScanTime    *int   `json:"lastScanTime,omitempty"`    // Time at which the device was scanned
}
type ResponseComplianceGetSecurityAdvisoriesAffectingTheNetworkDeviceV1 struct {
	Response *[]ResponseComplianceGetSecurityAdvisoriesAffectingTheNetworkDeviceV1Response `json:"response,omitempty"` //
	Version  string                                                                        `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetSecurityAdvisoriesAffectingTheNetworkDeviceV1Response struct {
	ID                     string                                                                                              `json:"id,omitempty"`                     // Id of the security advisory
	DeviceCount            *int                                                                                                `json:"deviceCount,omitempty"`            // Number of devices which are vulnerable to this advisory
	CveIDs                 []string                                                                                            `json:"cveIds,omitempty"`                 // CVE (Common Vulnerabilities and Exposures) ID of the advisory
	PublicationURL         string                                                                                              `json:"publicationUrl,omitempty"`         // Url for getting advisory details on cisco website
	CvssBaseScore          *float64                                                                                            `json:"cvssBaseScore,omitempty"`          // Common Vulnerability Scoring System(CVSS) base score
	SecurityImpactRating   string                                                                                              `json:"securityImpactRating,omitempty"`   // 'CRITICAL' - the advisory requires immediate mitigation. 'HIGH' - the advisory requires priority mitigation
	FirstFixedVersionsList *[]ResponseComplianceGetSecurityAdvisoriesAffectingTheNetworkDeviceV1ResponseFirstFixedVersionsList `json:"firstFixedVersionsList,omitempty"` //
}
type ResponseComplianceGetSecurityAdvisoriesAffectingTheNetworkDeviceV1ResponseFirstFixedVersionsList struct {
	VulnerableVersion string   `json:"vulnerableVersion,omitempty"` // Version that is vulnerable to the advisory
	FixedVersions     []string `json:"fixedVersions,omitempty"`     // First versions that have the fix for the advisory
}
type ResponseComplianceGetCountOfSecurityAdvisoriesAffectingTheNetworkDeviceV1 struct {
	Version  string                                                                             `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetCountOfSecurityAdvisoriesAffectingTheNetworkDeviceV1Response `json:"response,omitempty"` //
}
type ResponseComplianceGetCountOfSecurityAdvisoriesAffectingTheNetworkDeviceV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseComplianceGetSecurityAdvisoryAffectingTheNetworkDeviceByDeviceIDAndAdvisoryIDV1 struct {
	Response *ResponseComplianceGetSecurityAdvisoryAffectingTheNetworkDeviceByDeviceIDAndAdvisoryIDV1Response `json:"response,omitempty"` //
	Version  string                                                                                           `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetSecurityAdvisoryAffectingTheNetworkDeviceByDeviceIDAndAdvisoryIDV1Response struct {
	ID                     string                                                                                                                   `json:"id,omitempty"`                     // Id of the security advisory
	DeviceCount            *int                                                                                                                     `json:"deviceCount,omitempty"`            // Number of devices which are vulnerable to this advisory
	CveIDs                 []string                                                                                                                 `json:"cveIds,omitempty"`                 // CVE (Common Vulnerabilities and Exposures) ID of the advisory
	PublicationURL         string                                                                                                                   `json:"publicationUrl,omitempty"`         // Url for getting advisory details on cisco website
	CvssBaseScore          *float64                                                                                                                 `json:"cvssBaseScore,omitempty"`          // Common Vulnerability Scoring System(CVSS) base score
	SecurityImpactRating   string                                                                                                                   `json:"securityImpactRating,omitempty"`   // 'CRITICAL' - the advisory requires immediate mitigation. 'HIGH' - the advisory requires priority mitigation
	FirstFixedVersionsList *[]ResponseComplianceGetSecurityAdvisoryAffectingTheNetworkDeviceByDeviceIDAndAdvisoryIDV1ResponseFirstFixedVersionsList `json:"firstFixedVersionsList,omitempty"` //
}
type ResponseComplianceGetSecurityAdvisoryAffectingTheNetworkDeviceByDeviceIDAndAdvisoryIDV1ResponseFirstFixedVersionsList struct {
	VulnerableVersion string   `json:"vulnerableVersion,omitempty"` // Version that is vulnerable to the advisory
	FixedVersions     []string `json:"fixedVersions,omitempty"`     // First versions that have the fix for the advisory
}
type ResponseComplianceGetSecurityAdvisoriesResultsTrendOverTimeV1 struct {
	Response *[]ResponseComplianceGetSecurityAdvisoriesResultsTrendOverTimeV1Response `json:"response,omitempty"` //
	Version  string                                                                   `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetSecurityAdvisoriesResultsTrendOverTimeV1Response struct {
	CriticalSecurityImpactRatingAdvisoriesCount *int `json:"criticalSecurityImpactRatingAdvisoriesCount,omitempty"` // Number of advisories which have a security impact rating of critical
	HighSecurityImpactRatingAdvisoriesCount     *int `json:"highSecurityImpactRatingAdvisoriesCount,omitempty"`     // Number of advisories which have a security impact rating of high
	ScanTime                                    *int `json:"scanTime,omitempty"`                                    // End time for the scan
}
type ResponseComplianceGetCountOfSecurityAdvisoriesResultsTrendOverTimeV1 struct {
	Version  string                                                                        `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetCountOfSecurityAdvisoriesResultsTrendOverTimeV1Response `json:"response,omitempty"` //
}
type ResponseComplianceGetCountOfSecurityAdvisoriesResultsTrendOverTimeV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseComplianceGetTrialDetailsForSecurityAdvisoriesDetectionOnNetworkDevicesV1 struct {
	Version  string                                                                                     `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetTrialDetailsForSecurityAdvisoriesDetectionOnNetworkDevicesV1Response `json:"response,omitempty"` //
}
type ResponseComplianceGetTrialDetailsForSecurityAdvisoriesDetectionOnNetworkDevicesV1Response struct {
	Type                     string `json:"type,omitempty"`                     // Type of trial: * 'feature - the trial is of type feature. this is the currently supported type. * 'contract' - the trial is of type contract. this was used in older versions and exists only for compatibility.
	Feature                  string `json:"feature,omitempty"`                  // Name of the feature for which trial was created. for older versions that created contract type trials, this field will be absent.
	ContractLevel            string `json:"contractLevel,omitempty"`            // Contract level for which trial was created. this was used in older versions and exists only for compatibility.
	Active                   *bool  `json:"active,omitempty"`                   // Indicates if the trial is active
	StartTime                *int   `json:"startTime,omitempty"`                // Trial start time; as measured in Unix epoch time in milliseconds
	EndTime                  *int   `json:"endTime,omitempty"`                  // Trial end time; as measured in Unix epoch time in milliseconds
	SecondsRemainingToExpiry *int   `json:"secondsRemainingToExpiry,omitempty"` // Seconds remaining in the trial before it expires. for expired trials this will be 0.
	SecondsSinceExpired      *int   `json:"secondsSinceExpired,omitempty"`      // Seconds elapsed after the trial has expired. for active trials this will be 0.
}
type ResponseComplianceCreatesATrialForSecurityAdvisoriesDetectionOnNetworkDevicesV1 struct {
	Version  string                                                                                   `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseComplianceCreatesATrialForSecurityAdvisoriesDetectionOnNetworkDevicesV1Response `json:"response,omitempty"` //
}
type ResponseComplianceCreatesATrialForSecurityAdvisoriesDetectionOnNetworkDevicesV1Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseComplianceTriggersASecurityAdvisoriesScanForTheSupportedNetworkDevicesV1 struct {
	Version  string                                                                                    `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseComplianceTriggersASecurityAdvisoriesScanForTheSupportedNetworkDevicesV1Response `json:"response,omitempty"` //
}
type ResponseComplianceTriggersASecurityAdvisoriesScanForTheSupportedNetworkDevicesV1Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type RequestComplianceRunComplianceV1 struct {
	TriggerFull *bool    `json:"triggerFull,omitempty"` // if it is true then compliance will be triggered for all categories. If it is false then compliance will be triggered for categories mentioned in categories section .
	Categories  []string `json:"categories,omitempty"`  // Category can have any value among 'INTENT'(mapped to compliance types: NETWORK_SETTINGS,NETWORK_PROFILE,WORKFLOW,FABRIC,APPLICATION_VISIBILITY), 'RUNNING_CONFIG' , 'IMAGE' , 'PSIRT' , 'EoX' , 'NETWORK_SETTINGS'
	DeviceUUIDs []string `json:"deviceUuids,omitempty"` // UUID of the device.
}
type RequestComplianceCommitDeviceConfigurationV1 struct {
	DeviceID []string `json:"deviceId,omitempty"` // UUID of the device.
}

//GetComplianceStatusV1 Get Compliance Status  - dda5-cb9a-49aa-aef6
/* Return compliance status of device(s).


@param GetComplianceStatusV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-compliance-status
*/
func (s *ComplianceService) GetComplianceStatusV1(GetComplianceStatusV1QueryParams *GetComplianceStatusV1QueryParams) (*ResponseComplianceGetComplianceStatusV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/compliance"

	queryString, _ := query.Values(GetComplianceStatusV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetComplianceStatusV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetComplianceStatusV1(GetComplianceStatusV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetComplianceStatusV1")
	}

	result := response.Result().(*ResponseComplianceGetComplianceStatusV1)
	return result, response, err

}

//GetComplianceStatusCountV1 Get Compliance Status Count - db99-f919-424a-9f83
/* Return Compliance Status Count


@param GetComplianceStatusCountV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-compliance-status-count
*/
func (s *ComplianceService) GetComplianceStatusCountV1(GetComplianceStatusCountV1QueryParams *GetComplianceStatusCountV1QueryParams) (*ResponseComplianceGetComplianceStatusCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/compliance/count"

	queryString, _ := query.Values(GetComplianceStatusCountV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetComplianceStatusCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetComplianceStatusCountV1(GetComplianceStatusCountV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetComplianceStatusCountV1")
	}

	result := response.Result().(*ResponseComplianceGetComplianceStatusCountV1)
	return result, response, err

}

//GetComplianceDetailV1 Get Compliance Detail  - dda1-1ae7-4788-9d49
/* Return Compliance Detail


@param GetComplianceDetailV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-compliance-detail
*/
func (s *ComplianceService) GetComplianceDetailV1(GetComplianceDetailV1QueryParams *GetComplianceDetailV1QueryParams) (*ResponseComplianceGetComplianceDetailV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/compliance/detail"

	queryString, _ := query.Values(GetComplianceDetailV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetComplianceDetailV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetComplianceDetailV1(GetComplianceDetailV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetComplianceDetailV1")
	}

	result := response.Result().(*ResponseComplianceGetComplianceDetailV1)
	return result, response, err

}

//GetComplianceDetailCountV1 Get Compliance Detail Count - 3eb6-58c3-4549-94df
/* Return  Compliance Count Detail


@param GetComplianceDetailCountV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-compliance-detail-count
*/
func (s *ComplianceService) GetComplianceDetailCountV1(GetComplianceDetailCountV1QueryParams *GetComplianceDetailCountV1QueryParams) (*ResponseComplianceGetComplianceDetailCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/compliance/detail/count"

	queryString, _ := query.Values(GetComplianceDetailCountV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetComplianceDetailCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetComplianceDetailCountV1(GetComplianceDetailCountV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetComplianceDetailCountV1")
	}

	result := response.Result().(*ResponseComplianceGetComplianceDetailCountV1)
	return result, response, err

}

//DeviceComplianceStatusV1 Device Compliance Status - 7aa8-5ad5-48ea-94a7
/* Return compliance status of a device.


@param deviceUUID deviceUuid path parameter. Device Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!device-compliance-status
*/
func (s *ComplianceService) DeviceComplianceStatusV1(deviceUUID string) (*ResponseComplianceDeviceComplianceStatusV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/compliance/{deviceUuid}"
	path = strings.Replace(path, "{deviceUuid}", fmt.Sprintf("%v", deviceUUID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceDeviceComplianceStatusV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeviceComplianceStatusV1(deviceUUID)
		}
		return nil, response, fmt.Errorf("error with operation DeviceComplianceStatusV1")
	}

	result := response.Result().(*ResponseComplianceDeviceComplianceStatusV1)
	return result, response, err

}

//ComplianceDetailsOfDeviceV1 Compliance Details of Device - 52bf-e904-45aa-b017
/* Return compliance detailed report for a device.


@param deviceUUID deviceUuid path parameter. Device Id

@param ComplianceDetailsOfDeviceV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!compliance-details-of-device
*/
func (s *ComplianceService) ComplianceDetailsOfDeviceV1(deviceUUID string, ComplianceDetailsOfDeviceV1QueryParams *ComplianceDetailsOfDeviceV1QueryParams) (*ResponseComplianceComplianceDetailsOfDeviceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/compliance/{deviceUuid}/detail"
	path = strings.Replace(path, "{deviceUuid}", fmt.Sprintf("%v", deviceUUID), -1)

	queryString, _ := query.Values(ComplianceDetailsOfDeviceV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceComplianceDetailsOfDeviceV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ComplianceDetailsOfDeviceV1(deviceUUID, ComplianceDetailsOfDeviceV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ComplianceDetailsOfDeviceV1")
	}

	result := response.Result().(*ResponseComplianceComplianceDetailsOfDeviceV1)
	return result, response, err

}

//GetFieldNoticeNetworkDevicesV1 Get field notice network devices - e8b3-68d9-483b-8e07
/* Get field notice network devices


@param GetFieldNoticeNetworkDevicesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-field-notice-network-devices
*/
func (s *ComplianceService) GetFieldNoticeNetworkDevicesV1(GetFieldNoticeNetworkDevicesV1QueryParams *GetFieldNoticeNetworkDevicesV1QueryParams) (*ResponseComplianceGetFieldNoticeNetworkDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/fieldNotices/results/networkDevices"

	queryString, _ := query.Values(GetFieldNoticeNetworkDevicesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetFieldNoticeNetworkDevicesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFieldNoticeNetworkDevicesV1(GetFieldNoticeNetworkDevicesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetFieldNoticeNetworkDevicesV1")
	}

	result := response.Result().(*ResponseComplianceGetFieldNoticeNetworkDevicesV1)
	return result, response, err

}

//GetCountOfFieldNoticeNetworkDevicesV1 Get count of field notice network devices - 23bd-3911-4cc9-987c
/* Get count of field notice network devices


@param GetCountOfFieldNoticeNetworkDevicesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-count-of-field-notice-network-devices
*/
func (s *ComplianceService) GetCountOfFieldNoticeNetworkDevicesV1(GetCountOfFieldNoticeNetworkDevicesV1QueryParams *GetCountOfFieldNoticeNetworkDevicesV1QueryParams) (*ResponseComplianceGetCountOfFieldNoticeNetworkDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/fieldNotices/results/networkDevices/count"

	queryString, _ := query.Values(GetCountOfFieldNoticeNetworkDevicesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetCountOfFieldNoticeNetworkDevicesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCountOfFieldNoticeNetworkDevicesV1(GetCountOfFieldNoticeNetworkDevicesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetCountOfFieldNoticeNetworkDevicesV1")
	}

	result := response.Result().(*ResponseComplianceGetCountOfFieldNoticeNetworkDevicesV1)
	return result, response, err

}

//GetFieldNoticeNetworkDeviceByDeviceIDV1 Get field notice network device by device id - db80-68db-4b1b-976d
/* Get field notice network device by device id


@param networkDeviceID networkDeviceId path parameter. Id of the network device


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-field-notice-network-device-by-device-id
*/
func (s *ComplianceService) GetFieldNoticeNetworkDeviceByDeviceIDV1(networkDeviceID string) (*ResponseComplianceGetFieldNoticeNetworkDeviceByDeviceIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/fieldNotices/results/networkDevices/{networkDeviceId}"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceGetFieldNoticeNetworkDeviceByDeviceIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFieldNoticeNetworkDeviceByDeviceIDV1(networkDeviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetFieldNoticeNetworkDeviceByDeviceIdV1")
	}

	result := response.Result().(*ResponseComplianceGetFieldNoticeNetworkDeviceByDeviceIDV1)
	return result, response, err

}

//GetFieldNoticesAffectingTheNetworkDeviceV1 Get field notices affecting the network device - e5a6-3887-44c9-95d6
/* Get field notices affecting the network device


@param networkDeviceID networkDeviceId path parameter. Id of the network device

@param GetFieldNoticesAffectingTheNetworkDeviceV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-field-notices-affecting-the-network-device
*/
func (s *ComplianceService) GetFieldNoticesAffectingTheNetworkDeviceV1(networkDeviceID string, GetFieldNoticesAffectingTheNetworkDeviceV1QueryParams *GetFieldNoticesAffectingTheNetworkDeviceV1QueryParams) (*ResponseComplianceGetFieldNoticesAffectingTheNetworkDeviceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/fieldNotices/results/networkDevices/{networkDeviceId}/notices"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	queryString, _ := query.Values(GetFieldNoticesAffectingTheNetworkDeviceV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetFieldNoticesAffectingTheNetworkDeviceV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFieldNoticesAffectingTheNetworkDeviceV1(networkDeviceID, GetFieldNoticesAffectingTheNetworkDeviceV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetFieldNoticesAffectingTheNetworkDeviceV1")
	}

	result := response.Result().(*ResponseComplianceGetFieldNoticesAffectingTheNetworkDeviceV1)
	return result, response, err

}

//GetCountOfFieldNoticesAffectingTheNetworkDeviceV1 Get count of field notices affecting the network device - 5494-098c-414b-bf34
/* Get count of field notices affecting the network device


@param networkDeviceID networkDeviceId path parameter. Id of the network device

@param GetCountOfFieldNoticesAffectingTheNetworkDeviceV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-count-of-field-notices-affecting-the-network-device
*/
func (s *ComplianceService) GetCountOfFieldNoticesAffectingTheNetworkDeviceV1(networkDeviceID string, GetCountOfFieldNoticesAffectingTheNetworkDeviceV1QueryParams *GetCountOfFieldNoticesAffectingTheNetworkDeviceV1QueryParams) (*ResponseComplianceGetCountOfFieldNoticesAffectingTheNetworkDeviceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/fieldNotices/results/networkDevices/{networkDeviceId}/notices/count"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	queryString, _ := query.Values(GetCountOfFieldNoticesAffectingTheNetworkDeviceV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetCountOfFieldNoticesAffectingTheNetworkDeviceV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCountOfFieldNoticesAffectingTheNetworkDeviceV1(networkDeviceID, GetCountOfFieldNoticesAffectingTheNetworkDeviceV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetCountOfFieldNoticesAffectingTheNetworkDeviceV1")
	}

	result := response.Result().(*ResponseComplianceGetCountOfFieldNoticesAffectingTheNetworkDeviceV1)
	return result, response, err

}

//GetFieldNoticeAffectingTheNetworkDeviceByDeviceIDAndNoticeIDV1 Get field notice affecting the network device by device Id and notice id - 86a9-2ad8-436a-9299
/* Get field notice affecting the network device by device Id and notice id


@param networkDeviceID networkDeviceId path parameter. Id of the network device

@param id id path parameter. Id of the field notice


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-field-notice-affecting-the-network-device-by-device-id-and-notice-id
*/
func (s *ComplianceService) GetFieldNoticeAffectingTheNetworkDeviceByDeviceIDAndNoticeIDV1(networkDeviceID string, id string) (*ResponseComplianceGetFieldNoticeAffectingTheNetworkDeviceByDeviceIDAndNoticeIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/fieldNotices/results/networkDevices/{networkDeviceId}/notices/{id}"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceGetFieldNoticeAffectingTheNetworkDeviceByDeviceIDAndNoticeIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFieldNoticeAffectingTheNetworkDeviceByDeviceIDAndNoticeIDV1(networkDeviceID, id)
		}
		return nil, response, fmt.Errorf("error with operation GetFieldNoticeAffectingTheNetworkDeviceByDeviceIdAndNoticeIdV1")
	}

	result := response.Result().(*ResponseComplianceGetFieldNoticeAffectingTheNetworkDeviceByDeviceIDAndNoticeIDV1)
	return result, response, err

}

//GetFieldNoticesV1 Get field notices - 6989-39f3-4279-ae61
/* Get field notices


@param GetFieldNoticesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-field-notices
*/
func (s *ComplianceService) GetFieldNoticesV1(GetFieldNoticesV1QueryParams *GetFieldNoticesV1QueryParams) (*ResponseComplianceGetFieldNoticesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/fieldNotices/results/notices"

	queryString, _ := query.Values(GetFieldNoticesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetFieldNoticesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFieldNoticesV1(GetFieldNoticesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetFieldNoticesV1")
	}

	result := response.Result().(*ResponseComplianceGetFieldNoticesV1)
	return result, response, err

}

//GetCountOfFieldNoticesV1 Get count of field notices - ba99-e9ba-40cb-99e1
/* Get count of field notices


@param GetCountOfFieldNoticesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-count-of-field-notices
*/
func (s *ComplianceService) GetCountOfFieldNoticesV1(GetCountOfFieldNoticesV1QueryParams *GetCountOfFieldNoticesV1QueryParams) (*ResponseComplianceGetCountOfFieldNoticesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/fieldNotices/results/notices/count"

	queryString, _ := query.Values(GetCountOfFieldNoticesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetCountOfFieldNoticesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCountOfFieldNoticesV1(GetCountOfFieldNoticesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetCountOfFieldNoticesV1")
	}

	result := response.Result().(*ResponseComplianceGetCountOfFieldNoticesV1)
	return result, response, err

}

//GetFieldNoticeByIDV1 Get field notice by Id - 7c90-9909-4a19-8d77
/* Get field notice by Id


@param id id path parameter. Id of the field notice


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-field-notice-by-id
*/
func (s *ComplianceService) GetFieldNoticeByIDV1(id string) (*ResponseComplianceGetFieldNoticeByIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/fieldNotices/results/notices/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceGetFieldNoticeByIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFieldNoticeByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation GetFieldNoticeByIdV1")
	}

	result := response.Result().(*ResponseComplianceGetFieldNoticeByIDV1)
	return result, response, err

}

//GetFieldNoticeNetworkDevicesForTheNoticeV1 Get field notice network devices for the notice - ddaa-9b91-4cfa-8943
/* Get field notice network devices for the notice


@param id id path parameter. Id of the field notice

@param GetFieldNoticeNetworkDevicesForTheNoticeV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-field-notice-network-devices-for-the-notice
*/
func (s *ComplianceService) GetFieldNoticeNetworkDevicesForTheNoticeV1(id string, GetFieldNoticeNetworkDevicesForTheNoticeV1QueryParams *GetFieldNoticeNetworkDevicesForTheNoticeV1QueryParams) (*ResponseComplianceGetFieldNoticeNetworkDevicesForTheNoticeV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/fieldNotices/results/notices/{id}/networkDevices"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetFieldNoticeNetworkDevicesForTheNoticeV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetFieldNoticeNetworkDevicesForTheNoticeV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFieldNoticeNetworkDevicesForTheNoticeV1(id, GetFieldNoticeNetworkDevicesForTheNoticeV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetFieldNoticeNetworkDevicesForTheNoticeV1")
	}

	result := response.Result().(*ResponseComplianceGetFieldNoticeNetworkDevicesForTheNoticeV1)
	return result, response, err

}

//GetCountOfFieldNoticeNetworkDevicesForTheNoticeV1 Get count of field notice network devices for the notice - 4a9e-d86a-421b-890a
/* Get count of field notice network devices for the notice


@param id id path parameter. Id of the field notice

@param GetCountOfFieldNoticeNetworkDevicesForTheNoticeV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-count-of-field-notice-network-devices-for-the-notice
*/
func (s *ComplianceService) GetCountOfFieldNoticeNetworkDevicesForTheNoticeV1(id string, GetCountOfFieldNoticeNetworkDevicesForTheNoticeV1QueryParams *GetCountOfFieldNoticeNetworkDevicesForTheNoticeV1QueryParams) (*ResponseComplianceGetCountOfFieldNoticeNetworkDevicesForTheNoticeV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/fieldNotices/results/notices/{id}/networkDevices/count"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetCountOfFieldNoticeNetworkDevicesForTheNoticeV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetCountOfFieldNoticeNetworkDevicesForTheNoticeV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCountOfFieldNoticeNetworkDevicesForTheNoticeV1(id, GetCountOfFieldNoticeNetworkDevicesForTheNoticeV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetCountOfFieldNoticeNetworkDevicesForTheNoticeV1")
	}

	result := response.Result().(*ResponseComplianceGetCountOfFieldNoticeNetworkDevicesForTheNoticeV1)
	return result, response, err

}

//GetFieldNoticeNetworkDeviceForTheNoticeByNetworkDeviceIDV1 Get field notice network device for the notice by network device id - e4b4-cb51-46cb-9925
/* Get field notice network device for the notice by network device id


@param id id path parameter. Id of the field notice

@param networkDeviceID networkDeviceId path parameter. Id of the network device


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-field-notice-network-device-for-the-notice-by-network-device-id
*/
func (s *ComplianceService) GetFieldNoticeNetworkDeviceForTheNoticeByNetworkDeviceIDV1(id string, networkDeviceID string) (*ResponseComplianceGetFieldNoticeNetworkDeviceForTheNoticeByNetworkDeviceIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/fieldNotices/results/notices/{id}/networkDevices/{networkDeviceId}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceGetFieldNoticeNetworkDeviceForTheNoticeByNetworkDeviceIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFieldNoticeNetworkDeviceForTheNoticeByNetworkDeviceIDV1(id, networkDeviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetFieldNoticeNetworkDeviceForTheNoticeByNetworkDeviceIdV1")
	}

	result := response.Result().(*ResponseComplianceGetFieldNoticeNetworkDeviceForTheNoticeByNetworkDeviceIDV1)
	return result, response, err

}

//GetFieldNoticesResultsTrendOverTimeV1 Get field notices results trend over time - 6690-8bd1-4b58-ba8e
/* Get field notices results trend over time. The default sort is by scan time descending.


@param GetFieldNoticesResultsTrendOverTimeV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-field-notices-results-trend-over-time
*/
func (s *ComplianceService) GetFieldNoticesResultsTrendOverTimeV1(GetFieldNoticesResultsTrendOverTimeV1QueryParams *GetFieldNoticesResultsTrendOverTimeV1QueryParams) (*ResponseComplianceGetFieldNoticesResultsTrendOverTimeV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/fieldNotices/resultsTrend"

	queryString, _ := query.Values(GetFieldNoticesResultsTrendOverTimeV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetFieldNoticesResultsTrendOverTimeV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFieldNoticesResultsTrendOverTimeV1(GetFieldNoticesResultsTrendOverTimeV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetFieldNoticesResultsTrendOverTimeV1")
	}

	result := response.Result().(*ResponseComplianceGetFieldNoticesResultsTrendOverTimeV1)
	return result, response, err

}

//GetCountOfFieldNoticesResultsTrendOverTimeV1 Get count of field notices results trend over time - d285-c901-46b9-a120
/* Get count of field notices results trend over time


@param GetCountOfFieldNoticesResultsTrendOverTimeV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-count-of-field-notices-results-trend-over-time
*/
func (s *ComplianceService) GetCountOfFieldNoticesResultsTrendOverTimeV1(GetCountOfFieldNoticesResultsTrendOverTimeV1QueryParams *GetCountOfFieldNoticesResultsTrendOverTimeV1QueryParams) (*ResponseComplianceGetCountOfFieldNoticesResultsTrendOverTimeV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/fieldNotices/resultsTrend/count"

	queryString, _ := query.Values(GetCountOfFieldNoticesResultsTrendOverTimeV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetCountOfFieldNoticesResultsTrendOverTimeV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCountOfFieldNoticesResultsTrendOverTimeV1(GetCountOfFieldNoticesResultsTrendOverTimeV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetCountOfFieldNoticesResultsTrendOverTimeV1")
	}

	result := response.Result().(*ResponseComplianceGetCountOfFieldNoticesResultsTrendOverTimeV1)
	return result, response, err

}

//GetTrialDetailsForFieldNoticesDetectionOnNetworkDevicesV1 Get trial details for field notices detection on network devices - 92b9-d9a8-4a09-ad05
/* Get trial details for field notices detection on network devices



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-trial-details-for-field-notices-detection-on-network-devices
*/
func (s *ComplianceService) GetTrialDetailsForFieldNoticesDetectionOnNetworkDevicesV1() (*ResponseComplianceGetTrialDetailsForFieldNoticesDetectionOnNetworkDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/fieldNotices/trials"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceGetTrialDetailsForFieldNoticesDetectionOnNetworkDevicesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTrialDetailsForFieldNoticesDetectionOnNetworkDevicesV1()
		}
		return nil, response, fmt.Errorf("error with operation GetTrialDetailsForFieldNoticesDetectionOnNetworkDevicesV1")
	}

	result := response.Result().(*ResponseComplianceGetTrialDetailsForFieldNoticesDetectionOnNetworkDevicesV1)
	return result, response, err

}

//GetConfigTaskDetailsV1 Get config task details - 8183-1a90-4788-b8c5
/* Returns a config task result details by specified id


@param GetConfigTaskDetailsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-config-task-details
*/
func (s *ComplianceService) GetConfigTaskDetailsV1(GetConfigTaskDetailsV1QueryParams *GetConfigTaskDetailsV1QueryParams) (*ResponseComplianceGetConfigTaskDetailsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device-config/task"

	queryString, _ := query.Values(GetConfigTaskDetailsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetConfigTaskDetailsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetConfigTaskDetailsV1(GetConfigTaskDetailsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetConfigTaskDetailsV1")
	}

	result := response.Result().(*ResponseComplianceGetConfigTaskDetailsV1)
	return result, response, err

}

//GetNetworkBugsV1 Get network bugs - 3e8a-0a51-423a-968f
/* Get network bugs


@param GetNetworkBugsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-bugs
*/
func (s *ComplianceService) GetNetworkBugsV1(GetNetworkBugsV1QueryParams *GetNetworkBugsV1QueryParams) (*ResponseComplianceGetNetworkBugsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkBugs/results/bugs"

	queryString, _ := query.Values(GetNetworkBugsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetNetworkBugsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetNetworkBugsV1(GetNetworkBugsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetNetworkBugsV1")
	}

	result := response.Result().(*ResponseComplianceGetNetworkBugsV1)
	return result, response, err

}

//GetCountOfNetworkBugsV1 Get count of network bugs - 3cad-684d-4508-aa15
/* Get count of network bugs


@param GetCountOfNetworkBugsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-count-of-network-bugs
*/
func (s *ComplianceService) GetCountOfNetworkBugsV1(GetCountOfNetworkBugsV1QueryParams *GetCountOfNetworkBugsV1QueryParams) (*ResponseComplianceGetCountOfNetworkBugsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkBugs/results/bugs/count"

	queryString, _ := query.Values(GetCountOfNetworkBugsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetCountOfNetworkBugsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCountOfNetworkBugsV1(GetCountOfNetworkBugsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetCountOfNetworkBugsV1")
	}

	result := response.Result().(*ResponseComplianceGetCountOfNetworkBugsV1)
	return result, response, err

}

//GetNetworkBugByIDV1 Get network bug by Id - ec93-a9c6-48d9-9050
/* Get network bug by Id


@param id id path parameter. Id of the network bug


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-bug-by-id
*/
func (s *ComplianceService) GetNetworkBugByIDV1(id string) (*ResponseComplianceGetNetworkBugByIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkBugs/results/bugs/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceGetNetworkBugByIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetNetworkBugByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation GetNetworkBugByIdV1")
	}

	result := response.Result().(*ResponseComplianceGetNetworkBugByIDV1)
	return result, response, err

}

//GetNetworkBugDevicesForTheBugV1 Get network bug devices for the bug - a18c-2be4-4a1b-bbf7
/* Get network bug devices for the bug


@param id id path parameter. Id of the network bug

@param GetNetworkBugDevicesForTheBugV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-bug-devices-for-the-bug
*/
func (s *ComplianceService) GetNetworkBugDevicesForTheBugV1(id string, GetNetworkBugDevicesForTheBugV1QueryParams *GetNetworkBugDevicesForTheBugV1QueryParams) (*ResponseComplianceGetNetworkBugDevicesForTheBugV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkBugs/results/bugs/{id}/networkDevices"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetNetworkBugDevicesForTheBugV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetNetworkBugDevicesForTheBugV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetNetworkBugDevicesForTheBugV1(id, GetNetworkBugDevicesForTheBugV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetNetworkBugDevicesForTheBugV1")
	}

	result := response.Result().(*ResponseComplianceGetNetworkBugDevicesForTheBugV1)
	return result, response, err

}

//GetCountOfNetworkBugDevicesForTheBugV1 Get count of network bug devices for the bug - 269a-d906-4e5b-aad1
/* Get count of network bug devices for the bug


@param id id path parameter. Id of the network bug

@param GetCountOfNetworkBugDevicesForTheBugV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-count-of-network-bug-devices-for-the-bug
*/
func (s *ComplianceService) GetCountOfNetworkBugDevicesForTheBugV1(id string, GetCountOfNetworkBugDevicesForTheBugV1QueryParams *GetCountOfNetworkBugDevicesForTheBugV1QueryParams) (*ResponseComplianceGetCountOfNetworkBugDevicesForTheBugV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkBugs/results/bugs/{id}/networkDevices/count"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetCountOfNetworkBugDevicesForTheBugV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetCountOfNetworkBugDevicesForTheBugV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCountOfNetworkBugDevicesForTheBugV1(id, GetCountOfNetworkBugDevicesForTheBugV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetCountOfNetworkBugDevicesForTheBugV1")
	}

	result := response.Result().(*ResponseComplianceGetCountOfNetworkBugDevicesForTheBugV1)
	return result, response, err

}

//GetNetworkBugDeviceForTheBugByNetworkDeviceIDV1 Get network bug device for the bug by network device id - 7594-39b9-4d78-8144
/* Get network bug device for the bug by network device id


@param id id path parameter. Id of the network bug

@param networkDeviceID networkDeviceId path parameter. Id of the network device


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-bug-device-for-the-bug-by-network-device-id
*/
func (s *ComplianceService) GetNetworkBugDeviceForTheBugByNetworkDeviceIDV1(id string, networkDeviceID string) (*ResponseComplianceGetNetworkBugDeviceForTheBugByNetworkDeviceIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkBugs/results/bugs/{id}/networkDevices/{networkDeviceId}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceGetNetworkBugDeviceForTheBugByNetworkDeviceIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetNetworkBugDeviceForTheBugByNetworkDeviceIDV1(id, networkDeviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetNetworkBugDeviceForTheBugByNetworkDeviceIdV1")
	}

	result := response.Result().(*ResponseComplianceGetNetworkBugDeviceForTheBugByNetworkDeviceIDV1)
	return result, response, err

}

//GetNetworkBugDevicesV1 Get network bug devices - f9ad-f991-4c0a-9248
/* Get network bug devices


@param GetNetworkBugDevicesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-bug-devices
*/
func (s *ComplianceService) GetNetworkBugDevicesV1(GetNetworkBugDevicesV1QueryParams *GetNetworkBugDevicesV1QueryParams) (*ResponseComplianceGetNetworkBugDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkBugs/results/networkDevices"

	queryString, _ := query.Values(GetNetworkBugDevicesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetNetworkBugDevicesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetNetworkBugDevicesV1(GetNetworkBugDevicesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetNetworkBugDevicesV1")
	}

	result := response.Result().(*ResponseComplianceGetNetworkBugDevicesV1)
	return result, response, err

}

//GetCountOfNetworkBugDevicesV1 Get count of network bug devices - e583-ba39-4a4b-b8bd
/* Get count of network bug devices


@param GetCountOfNetworkBugDevicesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-count-of-network-bug-devices
*/
func (s *ComplianceService) GetCountOfNetworkBugDevicesV1(GetCountOfNetworkBugDevicesV1QueryParams *GetCountOfNetworkBugDevicesV1QueryParams) (*ResponseComplianceGetCountOfNetworkBugDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkBugs/results/networkDevices/count"

	queryString, _ := query.Values(GetCountOfNetworkBugDevicesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetCountOfNetworkBugDevicesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCountOfNetworkBugDevicesV1(GetCountOfNetworkBugDevicesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetCountOfNetworkBugDevicesV1")
	}

	result := response.Result().(*ResponseComplianceGetCountOfNetworkBugDevicesV1)
	return result, response, err

}

//GetNetworkBugDeviceByDeviceIDV1 Get network bug device by device id - 3eae-5b2d-43eb-b61f
/* Get network bug device by device id


@param networkDeviceID networkDeviceId path parameter. Id of the network device


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-bug-device-by-device-id
*/
func (s *ComplianceService) GetNetworkBugDeviceByDeviceIDV1(networkDeviceID string) (*ResponseComplianceGetNetworkBugDeviceByDeviceIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkBugs/results/networkDevices/{networkDeviceId}"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceGetNetworkBugDeviceByDeviceIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetNetworkBugDeviceByDeviceIDV1(networkDeviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetNetworkBugDeviceByDeviceIdV1")
	}

	result := response.Result().(*ResponseComplianceGetNetworkBugDeviceByDeviceIDV1)
	return result, response, err

}

//GetBugsAffectingTheNetworkDeviceV1 Get bugs affecting the network device - 199a-c80d-4138-8853
/* Get bugs affecting the network device


@param networkDeviceID networkDeviceId path parameter. Id of the network device

@param GetBugsAffectingTheNetworkDeviceV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-bugs-affecting-the-network-device
*/
func (s *ComplianceService) GetBugsAffectingTheNetworkDeviceV1(networkDeviceID string, GetBugsAffectingTheNetworkDeviceV1QueryParams *GetBugsAffectingTheNetworkDeviceV1QueryParams) (*ResponseComplianceGetBugsAffectingTheNetworkDeviceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkBugs/results/networkDevices/{networkDeviceId}/bugs"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	queryString, _ := query.Values(GetBugsAffectingTheNetworkDeviceV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetBugsAffectingTheNetworkDeviceV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetBugsAffectingTheNetworkDeviceV1(networkDeviceID, GetBugsAffectingTheNetworkDeviceV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetBugsAffectingTheNetworkDeviceV1")
	}

	result := response.Result().(*ResponseComplianceGetBugsAffectingTheNetworkDeviceV1)
	return result, response, err

}

//GetCountOfBugsAffectingTheNetworkDeviceV1 Get count of bugs affecting the network device - 86a4-d898-466a-a7e0
/* Get count of bugs affecting the network device


@param networkDeviceID networkDeviceId path parameter. Id of the network device

@param GetCountOfBugsAffectingTheNetworkDeviceV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-count-of-bugs-affecting-the-network-device
*/
func (s *ComplianceService) GetCountOfBugsAffectingTheNetworkDeviceV1(networkDeviceID string, GetCountOfBugsAffectingTheNetworkDeviceV1QueryParams *GetCountOfBugsAffectingTheNetworkDeviceV1QueryParams) (*ResponseComplianceGetCountOfBugsAffectingTheNetworkDeviceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkBugs/results/networkDevices/{networkDeviceId}/bugs/count"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	queryString, _ := query.Values(GetCountOfBugsAffectingTheNetworkDeviceV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetCountOfBugsAffectingTheNetworkDeviceV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCountOfBugsAffectingTheNetworkDeviceV1(networkDeviceID, GetCountOfBugsAffectingTheNetworkDeviceV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetCountOfBugsAffectingTheNetworkDeviceV1")
	}

	result := response.Result().(*ResponseComplianceGetCountOfBugsAffectingTheNetworkDeviceV1)
	return result, response, err

}

//GetBugAffectingTheNetworkDeviceByDeviceIDAndBugIDV1 Get bug affecting the network device by device Id and bug id - e293-6b08-49c9-b715
/* Get bug affecting the network device by device Id and bug id


@param networkDeviceID networkDeviceId path parameter. Id of the network device

@param id id path parameter. Id of the network bug


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-bug-affecting-the-network-device-by-device-id-and-bug-id
*/
func (s *ComplianceService) GetBugAffectingTheNetworkDeviceByDeviceIDAndBugIDV1(networkDeviceID string, id string) (*ResponseComplianceGetBugAffectingTheNetworkDeviceByDeviceIDAndBugIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkBugs/results/networkDevices/{networkDeviceId}/bugs/{id}"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceGetBugAffectingTheNetworkDeviceByDeviceIDAndBugIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetBugAffectingTheNetworkDeviceByDeviceIDAndBugIDV1(networkDeviceID, id)
		}
		return nil, response, fmt.Errorf("error with operation GetBugAffectingTheNetworkDeviceByDeviceIdAndBugIdV1")
	}

	result := response.Result().(*ResponseComplianceGetBugAffectingTheNetworkDeviceByDeviceIDAndBugIDV1)
	return result, response, err

}

//GetNetworkBugsResultsTrendOverTimeV1 Get network bugs results trend over time - 708d-2bc8-42da-b062
/* Get network bugs results trend over time. The default sort is by scan time descending.


@param GetNetworkBugsResultsTrendOverTimeV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-bugs-results-trend-over-time
*/
func (s *ComplianceService) GetNetworkBugsResultsTrendOverTimeV1(GetNetworkBugsResultsTrendOverTimeV1QueryParams *GetNetworkBugsResultsTrendOverTimeV1QueryParams) (*ResponseComplianceGetNetworkBugsResultsTrendOverTimeV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkBugs/resultsTrend"

	queryString, _ := query.Values(GetNetworkBugsResultsTrendOverTimeV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetNetworkBugsResultsTrendOverTimeV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetNetworkBugsResultsTrendOverTimeV1(GetNetworkBugsResultsTrendOverTimeV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetNetworkBugsResultsTrendOverTimeV1")
	}

	result := response.Result().(*ResponseComplianceGetNetworkBugsResultsTrendOverTimeV1)
	return result, response, err

}

//GetCountOfNetworkBugsResultsTrendOverTimeV1 Get count of network bugs results trend over time - 6791-696c-4199-86e1
/* Get count of network bugs results trend over time


@param GetCountOfNetworkBugsResultsTrendOverTimeV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-count-of-network-bugs-results-trend-over-time
*/
func (s *ComplianceService) GetCountOfNetworkBugsResultsTrendOverTimeV1(GetCountOfNetworkBugsResultsTrendOverTimeV1QueryParams *GetCountOfNetworkBugsResultsTrendOverTimeV1QueryParams) (*ResponseComplianceGetCountOfNetworkBugsResultsTrendOverTimeV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkBugs/resultsTrend/count"

	queryString, _ := query.Values(GetCountOfNetworkBugsResultsTrendOverTimeV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetCountOfNetworkBugsResultsTrendOverTimeV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCountOfNetworkBugsResultsTrendOverTimeV1(GetCountOfNetworkBugsResultsTrendOverTimeV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetCountOfNetworkBugsResultsTrendOverTimeV1")
	}

	result := response.Result().(*ResponseComplianceGetCountOfNetworkBugsResultsTrendOverTimeV1)
	return result, response, err

}

//GetTrialDetailsForBugsDetectionOnNetworkDevicesV1 Get trial details for bugs detection on network devices - 11a4-a89b-430b-93cd
/* Get trial details for bugs detection on network devices



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-trial-details-for-bugs-detection-on-network-devices
*/
func (s *ComplianceService) GetTrialDetailsForBugsDetectionOnNetworkDevicesV1() (*ResponseComplianceGetTrialDetailsForBugsDetectionOnNetworkDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkBugs/trials"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceGetTrialDetailsForBugsDetectionOnNetworkDevicesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTrialDetailsForBugsDetectionOnNetworkDevicesV1()
		}
		return nil, response, fmt.Errorf("error with operation GetTrialDetailsForBugsDetectionOnNetworkDevicesV1")
	}

	result := response.Result().(*ResponseComplianceGetTrialDetailsForBugsDetectionOnNetworkDevicesV1)
	return result, response, err

}

//GetSecurityAdvisoriesAffectingTheNetworkDevicesV1 Get security advisories affecting the network devices - ef91-f8be-47d8-8fbf
/* Get security advisories affecting the network devices


@param GetSecurityAdvisoriesAffectingTheNetworkDevicesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-security-advisories-affecting-the-network-devices
*/
func (s *ComplianceService) GetSecurityAdvisoriesAffectingTheNetworkDevicesV1(GetSecurityAdvisoriesAffectingTheNetworkDevicesV1QueryParams *GetSecurityAdvisoriesAffectingTheNetworkDevicesV1QueryParams) (*ResponseComplianceGetSecurityAdvisoriesAffectingTheNetworkDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/securityAdvisories/results/advisories"

	queryString, _ := query.Values(GetSecurityAdvisoriesAffectingTheNetworkDevicesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetSecurityAdvisoriesAffectingTheNetworkDevicesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSecurityAdvisoriesAffectingTheNetworkDevicesV1(GetSecurityAdvisoriesAffectingTheNetworkDevicesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSecurityAdvisoriesAffectingTheNetworkDevicesV1")
	}

	result := response.Result().(*ResponseComplianceGetSecurityAdvisoriesAffectingTheNetworkDevicesV1)
	return result, response, err

}

//GetCountOfSecurityAdvisoriesAffectingTheNetworkDevicesV1 Get count of security advisories affecting the network devices - 129c-9b1f-4dd8-9173
/* Get count of security advisories affecting the network devices


@param GetCountOfSecurityAdvisoriesAffectingTheNetworkDevicesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-count-of-security-advisories-affecting-the-network-devices
*/
func (s *ComplianceService) GetCountOfSecurityAdvisoriesAffectingTheNetworkDevicesV1(GetCountOfSecurityAdvisoriesAffectingTheNetworkDevicesV1QueryParams *GetCountOfSecurityAdvisoriesAffectingTheNetworkDevicesV1QueryParams) (*ResponseComplianceGetCountOfSecurityAdvisoriesAffectingTheNetworkDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/securityAdvisories/results/advisories/count"

	queryString, _ := query.Values(GetCountOfSecurityAdvisoriesAffectingTheNetworkDevicesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetCountOfSecurityAdvisoriesAffectingTheNetworkDevicesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCountOfSecurityAdvisoriesAffectingTheNetworkDevicesV1(GetCountOfSecurityAdvisoriesAffectingTheNetworkDevicesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetCountOfSecurityAdvisoriesAffectingTheNetworkDevicesV1")
	}

	result := response.Result().(*ResponseComplianceGetCountOfSecurityAdvisoriesAffectingTheNetworkDevicesV1)
	return result, response, err

}

//GetSecurityAdvisoryAffectingTheNetworkDevicesByIDV1 Get security advisory affecting the network devices by Id - 51aa-ea19-4c88-bea6
/* Get security advisory affecting the network devices by Id


@param id id path parameter. Id of the security advisory


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-security-advisory-affecting-the-network-devices-by-id
*/
func (s *ComplianceService) GetSecurityAdvisoryAffectingTheNetworkDevicesByIDV1(id string) (*ResponseComplianceGetSecurityAdvisoryAffectingTheNetworkDevicesByIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/securityAdvisories/results/advisories/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceGetSecurityAdvisoryAffectingTheNetworkDevicesByIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSecurityAdvisoryAffectingTheNetworkDevicesByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation GetSecurityAdvisoryAffectingTheNetworkDevicesByIdV1")
	}

	result := response.Result().(*ResponseComplianceGetSecurityAdvisoryAffectingTheNetworkDevicesByIDV1)
	return result, response, err

}

//GetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1 Get security advisory network devices for the security advisory - ee81-e9ad-40bb-b3d1
/* Get security advisory network devices for the security advisory


@param id id path parameter. Id of the security advisory

@param GetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-security-advisory-network-devices-for-the-security-advisory
*/
func (s *ComplianceService) GetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1(id string, GetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1QueryParams *GetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1QueryParams) (*ResponseComplianceGetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/securityAdvisories/results/advisories/{id}/networkDevices"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1(id, GetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1")
	}

	result := response.Result().(*ResponseComplianceGetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1)
	return result, response, err

}

//GetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1 Get count of security advisory network devices for the security advisory - 969b-bb96-404b-b905
/* Get count of security advisory network devices for the security advisory


@param id id path parameter. Id of the security advisory

@param GetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-count-of-security-advisory-network-devices-for-the-security-advisory
*/
func (s *ComplianceService) GetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1(id string, GetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1QueryParams *GetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1QueryParams) (*ResponseComplianceGetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/securityAdvisories/results/advisories/{id}/networkDevices/count"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1(id, GetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1")
	}

	result := response.Result().(*ResponseComplianceGetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1)
	return result, response, err

}

//GetSecurityAdvisoryNetworkDeviceForTheSecurityAdvisoryByNetworkDeviceIDV1 Get security advisory network device for the security advisory by network device id - 15ac-59b6-4668-a848
/* Get security advisory network device for the security advisory by network device id


@param id id path parameter. Id of the security advisory

@param networkDeviceID networkDeviceId path parameter. Id of the network device


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-security-advisory-network-device-for-the-security-advisory-by-network-device-id
*/
func (s *ComplianceService) GetSecurityAdvisoryNetworkDeviceForTheSecurityAdvisoryByNetworkDeviceIDV1(id string, networkDeviceID string) (*ResponseComplianceGetSecurityAdvisoryNetworkDeviceForTheSecurityAdvisoryByNetworkDeviceIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/securityAdvisories/results/advisories/{id}/networkDevices/{networkDeviceId}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceGetSecurityAdvisoryNetworkDeviceForTheSecurityAdvisoryByNetworkDeviceIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSecurityAdvisoryNetworkDeviceForTheSecurityAdvisoryByNetworkDeviceIDV1(id, networkDeviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetSecurityAdvisoryNetworkDeviceForTheSecurityAdvisoryByNetworkDeviceIdV1")
	}

	result := response.Result().(*ResponseComplianceGetSecurityAdvisoryNetworkDeviceForTheSecurityAdvisoryByNetworkDeviceIDV1)
	return result, response, err

}

//GetSecurityAdvisoryNetworkDevicesV1 Get security advisory network devices - af83-89a1-43da-9337
/* Get security advisory network devices


@param GetSecurityAdvisoryNetworkDevicesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-security-advisory-network-devices
*/
func (s *ComplianceService) GetSecurityAdvisoryNetworkDevicesV1(GetSecurityAdvisoryNetworkDevicesV1QueryParams *GetSecurityAdvisoryNetworkDevicesV1QueryParams) (*ResponseComplianceGetSecurityAdvisoryNetworkDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/securityAdvisories/results/networkDevices"

	queryString, _ := query.Values(GetSecurityAdvisoryNetworkDevicesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetSecurityAdvisoryNetworkDevicesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSecurityAdvisoryNetworkDevicesV1(GetSecurityAdvisoryNetworkDevicesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSecurityAdvisoryNetworkDevicesV1")
	}

	result := response.Result().(*ResponseComplianceGetSecurityAdvisoryNetworkDevicesV1)
	return result, response, err

}

//GetCountOfSecurityAdvisoryNetworkDevicesV1 Get count of security advisory network devices - 93a6-8af1-438a-8f39
/* Get count of security advisory network devices


@param GetCountOfSecurityAdvisoryNetworkDevicesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-count-of-security-advisory-network-devices
*/
func (s *ComplianceService) GetCountOfSecurityAdvisoryNetworkDevicesV1(GetCountOfSecurityAdvisoryNetworkDevicesV1QueryParams *GetCountOfSecurityAdvisoryNetworkDevicesV1QueryParams) (*ResponseComplianceGetCountOfSecurityAdvisoryNetworkDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/securityAdvisories/results/networkDevices/count"

	queryString, _ := query.Values(GetCountOfSecurityAdvisoryNetworkDevicesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetCountOfSecurityAdvisoryNetworkDevicesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCountOfSecurityAdvisoryNetworkDevicesV1(GetCountOfSecurityAdvisoryNetworkDevicesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetCountOfSecurityAdvisoryNetworkDevicesV1")
	}

	result := response.Result().(*ResponseComplianceGetCountOfSecurityAdvisoryNetworkDevicesV1)
	return result, response, err

}

//GetSecurityAdvisoryNetworkDeviceByNetworkDeviceIDV1 Get security advisory network device by network device id - a5bb-ca1a-4abb-8a7f
/* Get security advisory network device by network device id


@param networkDeviceID networkDeviceId path parameter. Id of the network device


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-security-advisory-network-device-by-network-device-id
*/
func (s *ComplianceService) GetSecurityAdvisoryNetworkDeviceByNetworkDeviceIDV1(networkDeviceID string) (*ResponseComplianceGetSecurityAdvisoryNetworkDeviceByNetworkDeviceIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/securityAdvisories/results/networkDevices/{networkDeviceId}"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceGetSecurityAdvisoryNetworkDeviceByNetworkDeviceIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSecurityAdvisoryNetworkDeviceByNetworkDeviceIDV1(networkDeviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetSecurityAdvisoryNetworkDeviceByNetworkDeviceIdV1")
	}

	result := response.Result().(*ResponseComplianceGetSecurityAdvisoryNetworkDeviceByNetworkDeviceIDV1)
	return result, response, err

}

//GetSecurityAdvisoriesAffectingTheNetworkDeviceV1 Get security advisories affecting the network device - 20a9-3b0d-4769-8091
/* Get security advisories affecting the network device


@param networkDeviceID networkDeviceId path parameter. Id of the network device

@param GetSecurityAdvisoriesAffectingTheNetworkDeviceV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-security-advisories-affecting-the-network-device
*/
func (s *ComplianceService) GetSecurityAdvisoriesAffectingTheNetworkDeviceV1(networkDeviceID string, GetSecurityAdvisoriesAffectingTheNetworkDeviceV1QueryParams *GetSecurityAdvisoriesAffectingTheNetworkDeviceV1QueryParams) (*ResponseComplianceGetSecurityAdvisoriesAffectingTheNetworkDeviceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/securityAdvisories/results/networkDevices/{networkDeviceId}/advisories"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	queryString, _ := query.Values(GetSecurityAdvisoriesAffectingTheNetworkDeviceV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetSecurityAdvisoriesAffectingTheNetworkDeviceV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSecurityAdvisoriesAffectingTheNetworkDeviceV1(networkDeviceID, GetSecurityAdvisoriesAffectingTheNetworkDeviceV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSecurityAdvisoriesAffectingTheNetworkDeviceV1")
	}

	result := response.Result().(*ResponseComplianceGetSecurityAdvisoriesAffectingTheNetworkDeviceV1)
	return result, response, err

}

//GetCountOfSecurityAdvisoriesAffectingTheNetworkDeviceV1 Get count of security advisories affecting the network device - d4ba-db3a-4488-9d47
/* Get count of security advisories affecting the network device


@param networkDeviceID networkDeviceId path parameter. Id of the network device

@param GetCountOfSecurityAdvisoriesAffectingTheNetworkDeviceV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-count-of-security-advisories-affecting-the-network-device
*/
func (s *ComplianceService) GetCountOfSecurityAdvisoriesAffectingTheNetworkDeviceV1(networkDeviceID string, GetCountOfSecurityAdvisoriesAffectingTheNetworkDeviceV1QueryParams *GetCountOfSecurityAdvisoriesAffectingTheNetworkDeviceV1QueryParams) (*ResponseComplianceGetCountOfSecurityAdvisoriesAffectingTheNetworkDeviceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/securityAdvisories/results/networkDevices/{networkDeviceId}/advisories/count"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	queryString, _ := query.Values(GetCountOfSecurityAdvisoriesAffectingTheNetworkDeviceV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetCountOfSecurityAdvisoriesAffectingTheNetworkDeviceV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCountOfSecurityAdvisoriesAffectingTheNetworkDeviceV1(networkDeviceID, GetCountOfSecurityAdvisoriesAffectingTheNetworkDeviceV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetCountOfSecurityAdvisoriesAffectingTheNetworkDeviceV1")
	}

	result := response.Result().(*ResponseComplianceGetCountOfSecurityAdvisoriesAffectingTheNetworkDeviceV1)
	return result, response, err

}

//GetSecurityAdvisoryAffectingTheNetworkDeviceByDeviceIDAndAdvisoryIDV1 Get security advisory affecting the network device by device Id and advisory id - c9a3-d93e-4fe8-959d
/* Get security advisory affecting the network device by device Id and advisory id


@param networkDeviceID networkDeviceId path parameter. Id of the network device

@param id id path parameter. Id of the security advisory


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-security-advisory-affecting-the-network-device-by-device-id-and-advisory-id
*/
func (s *ComplianceService) GetSecurityAdvisoryAffectingTheNetworkDeviceByDeviceIDAndAdvisoryIDV1(networkDeviceID string, id string) (*ResponseComplianceGetSecurityAdvisoryAffectingTheNetworkDeviceByDeviceIDAndAdvisoryIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/securityAdvisories/results/networkDevices/{networkDeviceId}/advisories/{id}"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceGetSecurityAdvisoryAffectingTheNetworkDeviceByDeviceIDAndAdvisoryIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSecurityAdvisoryAffectingTheNetworkDeviceByDeviceIDAndAdvisoryIDV1(networkDeviceID, id)
		}
		return nil, response, fmt.Errorf("error with operation GetSecurityAdvisoryAffectingTheNetworkDeviceByDeviceIdAndAdvisoryIdV1")
	}

	result := response.Result().(*ResponseComplianceGetSecurityAdvisoryAffectingTheNetworkDeviceByDeviceIDAndAdvisoryIDV1)
	return result, response, err

}

//GetSecurityAdvisoriesResultsTrendOverTimeV1 Get security advisories results trend over time - b584-aa2b-4158-bc5a
/* Get security advisories results trend over time. The default sort is by scan time descending.


@param GetSecurityAdvisoriesResultsTrendOverTimeV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-security-advisories-results-trend-over-time
*/
func (s *ComplianceService) GetSecurityAdvisoriesResultsTrendOverTimeV1(GetSecurityAdvisoriesResultsTrendOverTimeV1QueryParams *GetSecurityAdvisoriesResultsTrendOverTimeV1QueryParams) (*ResponseComplianceGetSecurityAdvisoriesResultsTrendOverTimeV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/securityAdvisories/resultsTrend"

	queryString, _ := query.Values(GetSecurityAdvisoriesResultsTrendOverTimeV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetSecurityAdvisoriesResultsTrendOverTimeV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSecurityAdvisoriesResultsTrendOverTimeV1(GetSecurityAdvisoriesResultsTrendOverTimeV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSecurityAdvisoriesResultsTrendOverTimeV1")
	}

	result := response.Result().(*ResponseComplianceGetSecurityAdvisoriesResultsTrendOverTimeV1)
	return result, response, err

}

//GetCountOfSecurityAdvisoriesResultsTrendOverTimeV1 Get count of security advisories results trend over time - a9af-78ef-46aa-8534
/* Get count of security advisories results trend over time.


@param GetCountOfSecurityAdvisoriesResultsTrendOverTimeV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-count-of-security-advisories-results-trend-over-time
*/
func (s *ComplianceService) GetCountOfSecurityAdvisoriesResultsTrendOverTimeV1(GetCountOfSecurityAdvisoriesResultsTrendOverTimeV1QueryParams *GetCountOfSecurityAdvisoriesResultsTrendOverTimeV1QueryParams) (*ResponseComplianceGetCountOfSecurityAdvisoriesResultsTrendOverTimeV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/securityAdvisories/resultsTrend/count"

	queryString, _ := query.Values(GetCountOfSecurityAdvisoriesResultsTrendOverTimeV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetCountOfSecurityAdvisoriesResultsTrendOverTimeV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCountOfSecurityAdvisoriesResultsTrendOverTimeV1(GetCountOfSecurityAdvisoriesResultsTrendOverTimeV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetCountOfSecurityAdvisoriesResultsTrendOverTimeV1")
	}

	result := response.Result().(*ResponseComplianceGetCountOfSecurityAdvisoriesResultsTrendOverTimeV1)
	return result, response, err

}

//GetTrialDetailsForSecurityAdvisoriesDetectionOnNetworkDevicesV1 Get trial details for security advisories detection on network devices - f6ba-8a34-4c4a-ba48
/* Get trial details for security advisories detection on network devices



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-trial-details-for-security-advisories-detection-on-network-devices
*/
func (s *ComplianceService) GetTrialDetailsForSecurityAdvisoriesDetectionOnNetworkDevicesV1() (*ResponseComplianceGetTrialDetailsForSecurityAdvisoriesDetectionOnNetworkDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/securityAdvisories/trials"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceGetTrialDetailsForSecurityAdvisoriesDetectionOnNetworkDevicesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTrialDetailsForSecurityAdvisoriesDetectionOnNetworkDevicesV1()
		}
		return nil, response, fmt.Errorf("error with operation GetTrialDetailsForSecurityAdvisoriesDetectionOnNetworkDevicesV1")
	}

	result := response.Result().(*ResponseComplianceGetTrialDetailsForSecurityAdvisoriesDetectionOnNetworkDevicesV1)
	return result, response, err

}

//RunComplianceV1 Run Compliance - f6ae-c8a7-4428-a9ff
/* Run compliance check for device(s).



Documentation Link: https://developer.cisco.com/docs/dna-center/#!run-compliance
*/
func (s *ComplianceService) RunComplianceV1(requestComplianceRunComplianceV1 *RequestComplianceRunComplianceV1) (*ResponseComplianceRunComplianceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/compliance/"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestComplianceRunComplianceV1).
		SetResult(&ResponseComplianceRunComplianceV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RunComplianceV1(requestComplianceRunComplianceV1)
		}

		return nil, response, fmt.Errorf("error with operation RunComplianceV1")
	}

	result := response.Result().(*ResponseComplianceRunComplianceV1)
	return result, response, err

}

//ComplianceRemediationV1 Compliance Remediation - 7d80-2867-4179-8488
/* Remediates configuration compliance issues. Compliance issues related to 'Routing', 'HA Remediation', 'Software Image', 'Securities Advisories', 'SD-Access Unsupported Configuration', 'Workflow', etc. will not be addressed by this API.
Warning: Fixing compliance mismatches could result in a possible network flap.


@param id id path parameter. Network device identifier


Documentation Link: https://developer.cisco.com/docs/dna-center/#!compliance-remediation
*/
func (s *ComplianceService) ComplianceRemediationV1(id string) (*ResponseComplianceComplianceRemediationV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/compliance/networkDevices/{id}/issues/remediation/provision"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceComplianceRemediationV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ComplianceRemediationV1(id)
		}

		return nil, response, fmt.Errorf("error with operation ComplianceRemediationV1")
	}

	result := response.Result().(*ResponseComplianceComplianceRemediationV1)
	return result, response, err

}

//CreatesATrialForFieldNoticesDetectionOnNetworkDevicesV1 Creates a trial for field notices detection on network devices - 3a9a-88e2-4c3a-9db8
/* Creates a trial for field notices detection on network devices. The consent to connect agreement must have been accepted in the UI for this to succeed. Please refer to the user guide at
 for more details on consent to connect.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!creates-a-trial-for-field-notices-detection-on-network-devices
*/
func (s *ComplianceService) CreatesATrialForFieldNoticesDetectionOnNetworkDevicesV1() (*ResponseComplianceCreatesATrialForFieldNoticesDetectionOnNetworkDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/fieldNotices/trials"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceCreatesATrialForFieldNoticesDetectionOnNetworkDevicesV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreatesATrialForFieldNoticesDetectionOnNetworkDevicesV1()
		}

		return nil, response, fmt.Errorf("error with operation CreatesATrialForFieldNoticesDetectionOnNetworkDevicesV1")
	}

	result := response.Result().(*ResponseComplianceCreatesATrialForFieldNoticesDetectionOnNetworkDevicesV1)
	return result, response, err

}

//TriggersAFieldNoticesScanForTheSupportedNetworkDevicesV1 Triggers a field notices scan for the supported network devices - d4b4-5ae2-4e68-bb04
/* Triggers a field notices scan for the supported network devices. The supported devices are switches, routers and wireless controllers. If a device is not supported, the FieldNoticeNetworkDevice scanStatus will be Failed with appropriate comments. The consent to connect agreement must have been accepted in the UI for this to succeed. Please refer to the user guide at
 for more details on consent to connect.


@param TriggersAFieldNoticesScanForTheSupportedNetworkDevicesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!triggers-a-field-notices-scan-for-the-supported-network-devices
*/
func (s *ComplianceService) TriggersAFieldNoticesScanForTheSupportedNetworkDevicesV1(TriggersAFieldNoticesScanForTheSupportedNetworkDevicesV1QueryParams *TriggersAFieldNoticesScanForTheSupportedNetworkDevicesV1QueryParams) (*ResponseComplianceTriggersAFieldNoticesScanForTheSupportedNetworkDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/fieldNotices/triggerScan"

	queryString, _ := query.Values(TriggersAFieldNoticesScanForTheSupportedNetworkDevicesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).
		SetResult(&ResponseComplianceTriggersAFieldNoticesScanForTheSupportedNetworkDevicesV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.TriggersAFieldNoticesScanForTheSupportedNetworkDevicesV1(TriggersAFieldNoticesScanForTheSupportedNetworkDevicesV1QueryParams)
		}

		return nil, response, fmt.Errorf("error with operation TriggersAFieldNoticesScanForTheSupportedNetworkDevicesV1")
	}

	result := response.Result().(*ResponseComplianceTriggersAFieldNoticesScanForTheSupportedNetworkDevicesV1)
	return result, response, err

}

//CommitDeviceConfigurationV1 Commit device configuration - 53a3-5a70-4e3b-87b5
/* This operation would commit device running configuration to startup by issuing "write memory" to device



Documentation Link: https://developer.cisco.com/docs/dna-center/#!commit-device-configuration
*/
func (s *ComplianceService) CommitDeviceConfigurationV1(requestComplianceCommitDeviceConfigurationV1 *RequestComplianceCommitDeviceConfigurationV1) (*ResponseComplianceCommitDeviceConfigurationV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device-config/write-memory"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestComplianceCommitDeviceConfigurationV1).
		SetResult(&ResponseComplianceCommitDeviceConfigurationV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CommitDeviceConfigurationV1(requestComplianceCommitDeviceConfigurationV1)
		}

		return nil, response, fmt.Errorf("error with operation CommitDeviceConfigurationV1")
	}

	result := response.Result().(*ResponseComplianceCommitDeviceConfigurationV1)
	return result, response, err

}

//CreatesATrialForBugsDetectionOnNetworkDevicesV1 Creates a trial for bugs detection on network devices - b080-6bcf-402b-ad8e
/* Creates a trial for bugs detection on network devices. The consent to connect agreement must have been accepted in the UI for this to succeed. Please refer to the user guide at
 for more details on consent to connect.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!creates-a-trial-for-bugs-detection-on-network-devices
*/
func (s *ComplianceService) CreatesATrialForBugsDetectionOnNetworkDevicesV1() (*ResponseComplianceCreatesATrialForBugsDetectionOnNetworkDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkBugs/trials"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceCreatesATrialForBugsDetectionOnNetworkDevicesV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreatesATrialForBugsDetectionOnNetworkDevicesV1()
		}

		return nil, response, fmt.Errorf("error with operation CreatesATrialForBugsDetectionOnNetworkDevicesV1")
	}

	result := response.Result().(*ResponseComplianceCreatesATrialForBugsDetectionOnNetworkDevicesV1)
	return result, response, err

}

//TriggersABugsScanForTheSupportedNetworkDevicesV1 Triggers a bugs scan for the supported network devices - 5296-db34-457b-b233
/* Triggers a bugs scan for the supported network devices. The supported devices are switches and routers. If a device is not supported, the NetworkBugsDevice scanStatus will be Failed with appropriate comments.


@param TriggersABugsScanForTheSupportedNetworkDevicesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!triggers-a-bugs-scan-for-the-supported-network-devices
*/
func (s *ComplianceService) TriggersABugsScanForTheSupportedNetworkDevicesV1(TriggersABugsScanForTheSupportedNetworkDevicesV1QueryParams *TriggersABugsScanForTheSupportedNetworkDevicesV1QueryParams) (*ResponseComplianceTriggersABugsScanForTheSupportedNetworkDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkBugs/triggerScan"

	queryString, _ := query.Values(TriggersABugsScanForTheSupportedNetworkDevicesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).
		SetResult(&ResponseComplianceTriggersABugsScanForTheSupportedNetworkDevicesV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.TriggersABugsScanForTheSupportedNetworkDevicesV1(TriggersABugsScanForTheSupportedNetworkDevicesV1QueryParams)
		}

		return nil, response, fmt.Errorf("error with operation TriggersABugsScanForTheSupportedNetworkDevicesV1")
	}

	result := response.Result().(*ResponseComplianceTriggersABugsScanForTheSupportedNetworkDevicesV1)
	return result, response, err

}

//CreatesATrialForSecurityAdvisoriesDetectionOnNetworkDevicesV1 Creates a trial for security advisories detection on network devices - 0190-1b7a-4edb-91d8
/* Creates a trial for security advisories detection on network devices. The consent to connect agreement must have been accepted in the UI for this to succeed. Please refer to the user guide at
 for more details on consent to connect.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!creates-a-trial-for-security-advisories-detection-on-network-devices
*/
func (s *ComplianceService) CreatesATrialForSecurityAdvisoriesDetectionOnNetworkDevicesV1() (*ResponseComplianceCreatesATrialForSecurityAdvisoriesDetectionOnNetworkDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/securityAdvisories/trials"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceCreatesATrialForSecurityAdvisoriesDetectionOnNetworkDevicesV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreatesATrialForSecurityAdvisoriesDetectionOnNetworkDevicesV1()
		}

		return nil, response, fmt.Errorf("error with operation CreatesATrialForSecurityAdvisoriesDetectionOnNetworkDevicesV1")
	}

	result := response.Result().(*ResponseComplianceCreatesATrialForSecurityAdvisoriesDetectionOnNetworkDevicesV1)
	return result, response, err

}

//TriggersASecurityAdvisoriesScanForTheSupportedNetworkDevicesV1 Triggers a security advisories scan for the supported network devices - a1a1-7b93-481b-9e03
/* Triggers a security advisories scan for the supported network devices. The supported devices are switches, routers and wireless controllers with IOS and IOS-XE. If a device is not supported, the SecurityAdvisoryNetworkDevice scanStatus will be Failed with appropriate comments.


@param TriggersASecurityAdvisoriesScanForTheSupportedNetworkDevicesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!triggers-a-security-advisories-scan-for-the-supported-network-devices
*/
func (s *ComplianceService) TriggersASecurityAdvisoriesScanForTheSupportedNetworkDevicesV1(TriggersASecurityAdvisoriesScanForTheSupportedNetworkDevicesV1QueryParams *TriggersASecurityAdvisoriesScanForTheSupportedNetworkDevicesV1QueryParams) (*ResponseComplianceTriggersASecurityAdvisoriesScanForTheSupportedNetworkDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/securityAdvisories/triggerScan"

	queryString, _ := query.Values(TriggersASecurityAdvisoriesScanForTheSupportedNetworkDevicesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).
		SetResult(&ResponseComplianceTriggersASecurityAdvisoriesScanForTheSupportedNetworkDevicesV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.TriggersASecurityAdvisoriesScanForTheSupportedNetworkDevicesV1(TriggersASecurityAdvisoriesScanForTheSupportedNetworkDevicesV1QueryParams)
		}

		return nil, response, fmt.Errorf("error with operation TriggersASecurityAdvisoriesScanForTheSupportedNetworkDevicesV1")
	}

	result := response.Result().(*ResponseComplianceTriggersASecurityAdvisoriesScanForTheSupportedNetworkDevicesV1)
	return result, response, err

}

// Alias Function
/*
This method acts as an alias for the method `GetComplianceDetailCountV1`
*/
func (s *ComplianceService) GetComplianceDetailCount(GetComplianceDetailCountV1QueryParams *GetComplianceDetailCountV1QueryParams) (*ResponseComplianceGetComplianceDetailCountV1, *resty.Response, error) {
	return s.GetComplianceDetailCountV1(GetComplianceDetailCountV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetCountOfFieldNoticesAffectingTheNetworkDeviceV1`
*/
func (s *ComplianceService) GetCountOfFieldNoticesAffectingTheNetworkDevice(networkDeviceID string, GetCountOfFieldNoticesAffectingTheNetworkDeviceV1QueryParams *GetCountOfFieldNoticesAffectingTheNetworkDeviceV1QueryParams) (*ResponseComplianceGetCountOfFieldNoticesAffectingTheNetworkDeviceV1, *resty.Response, error) {
	return s.GetCountOfFieldNoticesAffectingTheNetworkDeviceV1(networkDeviceID, GetCountOfFieldNoticesAffectingTheNetworkDeviceV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetCountOfNetworkBugsV1`
*/
func (s *ComplianceService) GetCountOfNetworkBugs(GetCountOfNetworkBugsV1QueryParams *GetCountOfNetworkBugsV1QueryParams) (*ResponseComplianceGetCountOfNetworkBugsV1, *resty.Response, error) {
	return s.GetCountOfNetworkBugsV1(GetCountOfNetworkBugsV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `ComplianceRemediationV1`
*/
func (s *ComplianceService) ComplianceRemediation(id string) (*ResponseComplianceComplianceRemediationV1, *resty.Response, error) {
	return s.ComplianceRemediationV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `GetComplianceDetailV1`
*/
func (s *ComplianceService) GetComplianceDetail(GetComplianceDetailV1QueryParams *GetComplianceDetailV1QueryParams) (*ResponseComplianceGetComplianceDetailV1, *resty.Response, error) {
	return s.GetComplianceDetailV1(GetComplianceDetailV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetFieldNoticeByIDV1`
*/
func (s *ComplianceService) GetFieldNoticeByID(id string) (*ResponseComplianceGetFieldNoticeByIDV1, *resty.Response, error) {
	return s.GetFieldNoticeByIDV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `GetNetworkBugDeviceByDeviceIDV1`
*/
func (s *ComplianceService) GetNetworkBugDeviceByDeviceID(networkDeviceID string) (*ResponseComplianceGetNetworkBugDeviceByDeviceIDV1, *resty.Response, error) {
	return s.GetNetworkBugDeviceByDeviceIDV1(networkDeviceID)
}

// Alias Function
/*
This method acts as an alias for the method `GetSecurityAdvisoryNetworkDeviceByNetworkDeviceIDV1`
*/
func (s *ComplianceService) GetSecurityAdvisoryNetworkDeviceByNetworkDeviceID(networkDeviceID string) (*ResponseComplianceGetSecurityAdvisoryNetworkDeviceByNetworkDeviceIDV1, *resty.Response, error) {
	return s.GetSecurityAdvisoryNetworkDeviceByNetworkDeviceIDV1(networkDeviceID)
}

// Alias Function
/*
This method acts as an alias for the method `GetCountOfSecurityAdvisoryNetworkDevicesV1`
*/
func (s *ComplianceService) GetCountOfSecurityAdvisoryNetworkDevices(GetCountOfSecurityAdvisoryNetworkDevicesV1QueryParams *GetCountOfSecurityAdvisoryNetworkDevicesV1QueryParams) (*ResponseComplianceGetCountOfSecurityAdvisoryNetworkDevicesV1, *resty.Response, error) {
	return s.GetCountOfSecurityAdvisoryNetworkDevicesV1(GetCountOfSecurityAdvisoryNetworkDevicesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetCountOfFieldNoticeNetworkDevicesForTheNoticeV1`
*/
func (s *ComplianceService) GetCountOfFieldNoticeNetworkDevicesForTheNotice(id string, GetCountOfFieldNoticeNetworkDevicesForTheNoticeV1QueryParams *GetCountOfFieldNoticeNetworkDevicesForTheNoticeV1QueryParams) (*ResponseComplianceGetCountOfFieldNoticeNetworkDevicesForTheNoticeV1, *resty.Response, error) {
	return s.GetCountOfFieldNoticeNetworkDevicesForTheNoticeV1(id, GetCountOfFieldNoticeNetworkDevicesForTheNoticeV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetNetworkBugByIDV1`
*/
func (s *ComplianceService) GetNetworkBugByID(id string) (*ResponseComplianceGetNetworkBugByIDV1, *resty.Response, error) {
	return s.GetNetworkBugByIDV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `GetCountOfNetworkBugDevicesV1`
*/
func (s *ComplianceService) GetCountOfNetworkBugDevices(GetCountOfNetworkBugDevicesV1QueryParams *GetCountOfNetworkBugDevicesV1QueryParams) (*ResponseComplianceGetCountOfNetworkBugDevicesV1, *resty.Response, error) {
	return s.GetCountOfNetworkBugDevicesV1(GetCountOfNetworkBugDevicesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetConfigTaskDetailsV1`
*/
func (s *ComplianceService) GetConfigTaskDetails(GetConfigTaskDetailsV1QueryParams *GetConfigTaskDetailsV1QueryParams) (*ResponseComplianceGetConfigTaskDetailsV1, *resty.Response, error) {
	return s.GetConfigTaskDetailsV1(GetConfigTaskDetailsV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetComplianceStatusV1`
*/
func (s *ComplianceService) GetComplianceStatus(GetComplianceStatusV1QueryParams *GetComplianceStatusV1QueryParams) (*ResponseComplianceGetComplianceStatusV1, *resty.Response, error) {
	return s.GetComplianceStatusV1(GetComplianceStatusV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1`
*/
func (s *ComplianceService) GetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisory(id string, GetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1QueryParams *GetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1QueryParams) (*ResponseComplianceGetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1, *resty.Response, error) {
	return s.GetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1(id, GetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetCountOfBugsAffectingTheNetworkDeviceV1`
*/
func (s *ComplianceService) GetCountOfBugsAffectingTheNetworkDevice(networkDeviceID string, GetCountOfBugsAffectingTheNetworkDeviceV1QueryParams *GetCountOfBugsAffectingTheNetworkDeviceV1QueryParams) (*ResponseComplianceGetCountOfBugsAffectingTheNetworkDeviceV1, *resty.Response, error) {
	return s.GetCountOfBugsAffectingTheNetworkDeviceV1(networkDeviceID, GetCountOfBugsAffectingTheNetworkDeviceV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetCountOfFieldNoticeNetworkDevicesV1`
*/
func (s *ComplianceService) GetCountOfFieldNoticeNetworkDevices(GetCountOfFieldNoticeNetworkDevicesV1QueryParams *GetCountOfFieldNoticeNetworkDevicesV1QueryParams) (*ResponseComplianceGetCountOfFieldNoticeNetworkDevicesV1, *resty.Response, error) {
	return s.GetCountOfFieldNoticeNetworkDevicesV1(GetCountOfFieldNoticeNetworkDevicesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetCountOfSecurityAdvisoriesAffectingTheNetworkDevicesV1`
*/
func (s *ComplianceService) GetCountOfSecurityAdvisoriesAffectingTheNetworkDevices(GetCountOfSecurityAdvisoriesAffectingTheNetworkDevicesV1QueryParams *GetCountOfSecurityAdvisoriesAffectingTheNetworkDevicesV1QueryParams) (*ResponseComplianceGetCountOfSecurityAdvisoriesAffectingTheNetworkDevicesV1, *resty.Response, error) {
	return s.GetCountOfSecurityAdvisoriesAffectingTheNetworkDevicesV1(GetCountOfSecurityAdvisoriesAffectingTheNetworkDevicesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetSecurityAdvisoriesAffectingTheNetworkDeviceV1`
*/
func (s *ComplianceService) GetSecurityAdvisoriesAffectingTheNetworkDevice(networkDeviceID string, GetSecurityAdvisoriesAffectingTheNetworkDeviceV1QueryParams *GetSecurityAdvisoriesAffectingTheNetworkDeviceV1QueryParams) (*ResponseComplianceGetSecurityAdvisoriesAffectingTheNetworkDeviceV1, *resty.Response, error) {
	return s.GetSecurityAdvisoriesAffectingTheNetworkDeviceV1(networkDeviceID, GetSecurityAdvisoriesAffectingTheNetworkDeviceV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `TriggersABugsScanForTheSupportedNetworkDevicesV1`
*/
func (s *ComplianceService) TriggersABugsScanForTheSupportedNetworkDevices(TriggersABugsScanForTheSupportedNetworkDevicesV1QueryParams *TriggersABugsScanForTheSupportedNetworkDevicesV1QueryParams) (*ResponseComplianceTriggersABugsScanForTheSupportedNetworkDevicesV1, *resty.Response, error) {
	return s.TriggersABugsScanForTheSupportedNetworkDevicesV1(TriggersABugsScanForTheSupportedNetworkDevicesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetSecurityAdvisoryAffectingTheNetworkDeviceByDeviceIDAndAdvisoryIDV1`
*/
func (s *ComplianceService) GetSecurityAdvisoryAffectingTheNetworkDeviceByDeviceIDAndAdvisoryID(networkDeviceID string, id string) (*ResponseComplianceGetSecurityAdvisoryAffectingTheNetworkDeviceByDeviceIDAndAdvisoryIDV1, *resty.Response, error) {
	return s.GetSecurityAdvisoryAffectingTheNetworkDeviceByDeviceIDAndAdvisoryIDV1(networkDeviceID, id)
}

// Alias Function
/*
This method acts as an alias for the method `CreatesATrialForBugsDetectionOnNetworkDevicesV1`
*/
func (s *ComplianceService) CreatesATrialForBugsDetectionOnNetworkDevices() (*ResponseComplianceCreatesATrialForBugsDetectionOnNetworkDevicesV1, *resty.Response, error) {
	return s.CreatesATrialForBugsDetectionOnNetworkDevicesV1()
}

// Alias Function
/*
This method acts as an alias for the method `GetCountOfSecurityAdvisoriesAffectingTheNetworkDeviceV1`
*/
func (s *ComplianceService) GetCountOfSecurityAdvisoriesAffectingTheNetworkDevice(networkDeviceID string, GetCountOfSecurityAdvisoriesAffectingTheNetworkDeviceV1QueryParams *GetCountOfSecurityAdvisoriesAffectingTheNetworkDeviceV1QueryParams) (*ResponseComplianceGetCountOfSecurityAdvisoriesAffectingTheNetworkDeviceV1, *resty.Response, error) {
	return s.GetCountOfSecurityAdvisoriesAffectingTheNetworkDeviceV1(networkDeviceID, GetCountOfSecurityAdvisoriesAffectingTheNetworkDeviceV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetSecurityAdvisoriesResultsTrendOverTimeV1`
*/
func (s *ComplianceService) GetSecurityAdvisoriesResultsTrendOverTime(GetSecurityAdvisoriesResultsTrendOverTimeV1QueryParams *GetSecurityAdvisoriesResultsTrendOverTimeV1QueryParams) (*ResponseComplianceGetSecurityAdvisoriesResultsTrendOverTimeV1, *resty.Response, error) {
	return s.GetSecurityAdvisoriesResultsTrendOverTimeV1(GetSecurityAdvisoriesResultsTrendOverTimeV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetNetworkBugsV1`
*/
func (s *ComplianceService) GetNetworkBugs(GetNetworkBugsV1QueryParams *GetNetworkBugsV1QueryParams) (*ResponseComplianceGetNetworkBugsV1, *resty.Response, error) {
	return s.GetNetworkBugsV1(GetNetworkBugsV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetCountOfFieldNoticesV1`
*/
func (s *ComplianceService) GetCountOfFieldNotices(GetCountOfFieldNoticesV1QueryParams *GetCountOfFieldNoticesV1QueryParams) (*ResponseComplianceGetCountOfFieldNoticesV1, *resty.Response, error) {
	return s.GetCountOfFieldNoticesV1(GetCountOfFieldNoticesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetTrialDetailsForSecurityAdvisoriesDetectionOnNetworkDevicesV1`
*/
func (s *ComplianceService) GetTrialDetailsForSecurityAdvisoriesDetectionOnNetworkDevices() (*ResponseComplianceGetTrialDetailsForSecurityAdvisoriesDetectionOnNetworkDevicesV1, *resty.Response, error) {
	return s.GetTrialDetailsForSecurityAdvisoriesDetectionOnNetworkDevicesV1()
}

// Alias Function
/*
This method acts as an alias for the method `GetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1`
*/
func (s *ComplianceService) GetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisory(id string, GetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1QueryParams *GetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1QueryParams) (*ResponseComplianceGetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1, *resty.Response, error) {
	return s.GetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1(id, GetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetFieldNoticesV1`
*/
func (s *ComplianceService) GetFieldNotices(GetFieldNoticesV1QueryParams *GetFieldNoticesV1QueryParams) (*ResponseComplianceGetFieldNoticesV1, *resty.Response, error) {
	return s.GetFieldNoticesV1(GetFieldNoticesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetCountOfSecurityAdvisoriesResultsTrendOverTimeV1`
*/
func (s *ComplianceService) GetCountOfSecurityAdvisoriesResultsTrendOverTime(GetCountOfSecurityAdvisoriesResultsTrendOverTimeV1QueryParams *GetCountOfSecurityAdvisoriesResultsTrendOverTimeV1QueryParams) (*ResponseComplianceGetCountOfSecurityAdvisoriesResultsTrendOverTimeV1, *resty.Response, error) {
	return s.GetCountOfSecurityAdvisoriesResultsTrendOverTimeV1(GetCountOfSecurityAdvisoriesResultsTrendOverTimeV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetNetworkBugDevicesForTheBugV1`
*/
func (s *ComplianceService) GetNetworkBugDevicesForTheBug(id string, GetNetworkBugDevicesForTheBugV1QueryParams *GetNetworkBugDevicesForTheBugV1QueryParams) (*ResponseComplianceGetNetworkBugDevicesForTheBugV1, *resty.Response, error) {
	return s.GetNetworkBugDevicesForTheBugV1(id, GetNetworkBugDevicesForTheBugV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetFieldNoticesAffectingTheNetworkDeviceV1`
*/
func (s *ComplianceService) GetFieldNoticesAffectingTheNetworkDevice(networkDeviceID string, GetFieldNoticesAffectingTheNetworkDeviceV1QueryParams *GetFieldNoticesAffectingTheNetworkDeviceV1QueryParams) (*ResponseComplianceGetFieldNoticesAffectingTheNetworkDeviceV1, *resty.Response, error) {
	return s.GetFieldNoticesAffectingTheNetworkDeviceV1(networkDeviceID, GetFieldNoticesAffectingTheNetworkDeviceV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetCountOfNetworkBugsResultsTrendOverTimeV1`
*/
func (s *ComplianceService) GetCountOfNetworkBugsResultsTrendOverTime(GetCountOfNetworkBugsResultsTrendOverTimeV1QueryParams *GetCountOfNetworkBugsResultsTrendOverTimeV1QueryParams) (*ResponseComplianceGetCountOfNetworkBugsResultsTrendOverTimeV1, *resty.Response, error) {
	return s.GetCountOfNetworkBugsResultsTrendOverTimeV1(GetCountOfNetworkBugsResultsTrendOverTimeV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetTrialDetailsForBugsDetectionOnNetworkDevicesV1`
*/
func (s *ComplianceService) GetTrialDetailsForBugsDetectionOnNetworkDevices() (*ResponseComplianceGetTrialDetailsForBugsDetectionOnNetworkDevicesV1, *resty.Response, error) {
	return s.GetTrialDetailsForBugsDetectionOnNetworkDevicesV1()
}

// Alias Function
/*
This method acts as an alias for the method `GetFieldNoticeNetworkDeviceByDeviceIDV1`
*/
func (s *ComplianceService) GetFieldNoticeNetworkDeviceByDeviceID(networkDeviceID string) (*ResponseComplianceGetFieldNoticeNetworkDeviceByDeviceIDV1, *resty.Response, error) {
	return s.GetFieldNoticeNetworkDeviceByDeviceIDV1(networkDeviceID)
}

// Alias Function
/*
This method acts as an alias for the method `GetBugsAffectingTheNetworkDeviceV1`
*/
func (s *ComplianceService) GetBugsAffectingTheNetworkDevice(networkDeviceID string, GetBugsAffectingTheNetworkDeviceV1QueryParams *GetBugsAffectingTheNetworkDeviceV1QueryParams) (*ResponseComplianceGetBugsAffectingTheNetworkDeviceV1, *resty.Response, error) {
	return s.GetBugsAffectingTheNetworkDeviceV1(networkDeviceID, GetBugsAffectingTheNetworkDeviceV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `DeviceComplianceStatusV1`
*/
func (s *ComplianceService) DeviceComplianceStatus(deviceUUID string) (*ResponseComplianceDeviceComplianceStatusV1, *resty.Response, error) {
	return s.DeviceComplianceStatusV1(deviceUUID)
}

// Alias Function
/*
This method acts as an alias for the method `GetCountOfFieldNoticesResultsTrendOverTimeV1`
*/
func (s *ComplianceService) GetCountOfFieldNoticesResultsTrendOverTime(GetCountOfFieldNoticesResultsTrendOverTimeV1QueryParams *GetCountOfFieldNoticesResultsTrendOverTimeV1QueryParams) (*ResponseComplianceGetCountOfFieldNoticesResultsTrendOverTimeV1, *resty.Response, error) {
	return s.GetCountOfFieldNoticesResultsTrendOverTimeV1(GetCountOfFieldNoticesResultsTrendOverTimeV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `CommitDeviceConfigurationV1`
*/
func (s *ComplianceService) CommitDeviceConfiguration(requestComplianceCommitDeviceConfigurationV1 *RequestComplianceCommitDeviceConfigurationV1) (*ResponseComplianceCommitDeviceConfigurationV1, *resty.Response, error) {
	return s.CommitDeviceConfigurationV1(requestComplianceCommitDeviceConfigurationV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetSecurityAdvisoryAffectingTheNetworkDevicesByIDV1`
*/
func (s *ComplianceService) GetSecurityAdvisoryAffectingTheNetworkDevicesByID(id string) (*ResponseComplianceGetSecurityAdvisoryAffectingTheNetworkDevicesByIDV1, *resty.Response, error) {
	return s.GetSecurityAdvisoryAffectingTheNetworkDevicesByIDV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `GetNetworkBugDevicesV1`
*/
func (s *ComplianceService) GetNetworkBugDevices(GetNetworkBugDevicesV1QueryParams *GetNetworkBugDevicesV1QueryParams) (*ResponseComplianceGetNetworkBugDevicesV1, *resty.Response, error) {
	return s.GetNetworkBugDevicesV1(GetNetworkBugDevicesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `ComplianceDetailsOfDeviceV1`
*/
func (s *ComplianceService) ComplianceDetailsOfDevice(deviceUUID string, ComplianceDetailsOfDeviceV1QueryParams *ComplianceDetailsOfDeviceV1QueryParams) (*ResponseComplianceComplianceDetailsOfDeviceV1, *resty.Response, error) {
	return s.ComplianceDetailsOfDeviceV1(deviceUUID, ComplianceDetailsOfDeviceV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetCountOfNetworkBugDevicesForTheBugV1`
*/
func (s *ComplianceService) GetCountOfNetworkBugDevicesForTheBug(id string, GetCountOfNetworkBugDevicesForTheBugV1QueryParams *GetCountOfNetworkBugDevicesForTheBugV1QueryParams) (*ResponseComplianceGetCountOfNetworkBugDevicesForTheBugV1, *resty.Response, error) {
	return s.GetCountOfNetworkBugDevicesForTheBugV1(id, GetCountOfNetworkBugDevicesForTheBugV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetFieldNoticesResultsTrendOverTimeV1`
*/
func (s *ComplianceService) GetFieldNoticesResultsTrendOverTime(GetFieldNoticesResultsTrendOverTimeV1QueryParams *GetFieldNoticesResultsTrendOverTimeV1QueryParams) (*ResponseComplianceGetFieldNoticesResultsTrendOverTimeV1, *resty.Response, error) {
	return s.GetFieldNoticesResultsTrendOverTimeV1(GetFieldNoticesResultsTrendOverTimeV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetFieldNoticeNetworkDevicesForTheNoticeV1`
*/
func (s *ComplianceService) GetFieldNoticeNetworkDevicesForTheNotice(id string, GetFieldNoticeNetworkDevicesForTheNoticeV1QueryParams *GetFieldNoticeNetworkDevicesForTheNoticeV1QueryParams) (*ResponseComplianceGetFieldNoticeNetworkDevicesForTheNoticeV1, *resty.Response, error) {
	return s.GetFieldNoticeNetworkDevicesForTheNoticeV1(id, GetFieldNoticeNetworkDevicesForTheNoticeV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetSecurityAdvisoryNetworkDeviceForTheSecurityAdvisoryByNetworkDeviceIDV1`
*/
func (s *ComplianceService) GetSecurityAdvisoryNetworkDeviceForTheSecurityAdvisoryByNetworkDeviceID(id string, networkDeviceID string) (*ResponseComplianceGetSecurityAdvisoryNetworkDeviceForTheSecurityAdvisoryByNetworkDeviceIDV1, *resty.Response, error) {
	return s.GetSecurityAdvisoryNetworkDeviceForTheSecurityAdvisoryByNetworkDeviceIDV1(id, networkDeviceID)
}

// Alias Function
/*
This method acts as an alias for the method `GetFieldNoticeAffectingTheNetworkDeviceByDeviceIDAndNoticeIDV1`
*/
func (s *ComplianceService) GetFieldNoticeAffectingTheNetworkDeviceByDeviceIDAndNoticeID(networkDeviceID string, id string) (*ResponseComplianceGetFieldNoticeAffectingTheNetworkDeviceByDeviceIDAndNoticeIDV1, *resty.Response, error) {
	return s.GetFieldNoticeAffectingTheNetworkDeviceByDeviceIDAndNoticeIDV1(networkDeviceID, id)
}

// Alias Function
/*
This method acts as an alias for the method `GetSecurityAdvisoriesAffectingTheNetworkDevicesV1`
*/
func (s *ComplianceService) GetSecurityAdvisoriesAffectingTheNetworkDevices(GetSecurityAdvisoriesAffectingTheNetworkDevicesV1QueryParams *GetSecurityAdvisoriesAffectingTheNetworkDevicesV1QueryParams) (*ResponseComplianceGetSecurityAdvisoriesAffectingTheNetworkDevicesV1, *resty.Response, error) {
	return s.GetSecurityAdvisoriesAffectingTheNetworkDevicesV1(GetSecurityAdvisoriesAffectingTheNetworkDevicesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `TriggersASecurityAdvisoriesScanForTheSupportedNetworkDevicesV1`
*/
func (s *ComplianceService) TriggersASecurityAdvisoriesScanForTheSupportedNetworkDevices(TriggersASecurityAdvisoriesScanForTheSupportedNetworkDevicesV1QueryParams *TriggersASecurityAdvisoriesScanForTheSupportedNetworkDevicesV1QueryParams) (*ResponseComplianceTriggersASecurityAdvisoriesScanForTheSupportedNetworkDevicesV1, *resty.Response, error) {
	return s.TriggersASecurityAdvisoriesScanForTheSupportedNetworkDevicesV1(TriggersASecurityAdvisoriesScanForTheSupportedNetworkDevicesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetComplianceStatusCountV1`
*/
func (s *ComplianceService) GetComplianceStatusCount(GetComplianceStatusCountV1QueryParams *GetComplianceStatusCountV1QueryParams) (*ResponseComplianceGetComplianceStatusCountV1, *resty.Response, error) {
	return s.GetComplianceStatusCountV1(GetComplianceStatusCountV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetNetworkBugDeviceForTheBugByNetworkDeviceIDV1`
*/
func (s *ComplianceService) GetNetworkBugDeviceForTheBugByNetworkDeviceID(id string, networkDeviceID string) (*ResponseComplianceGetNetworkBugDeviceForTheBugByNetworkDeviceIDV1, *resty.Response, error) {
	return s.GetNetworkBugDeviceForTheBugByNetworkDeviceIDV1(id, networkDeviceID)
}

// Alias Function
/*
This method acts as an alias for the method `GetNetworkBugsResultsTrendOverTimeV1`
*/
func (s *ComplianceService) GetNetworkBugsResultsTrendOverTime(GetNetworkBugsResultsTrendOverTimeV1QueryParams *GetNetworkBugsResultsTrendOverTimeV1QueryParams) (*ResponseComplianceGetNetworkBugsResultsTrendOverTimeV1, *resty.Response, error) {
	return s.GetNetworkBugsResultsTrendOverTimeV1(GetNetworkBugsResultsTrendOverTimeV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetBugAffectingTheNetworkDeviceByDeviceIDAndBugIDV1`
*/
func (s *ComplianceService) GetBugAffectingTheNetworkDeviceByDeviceIDAndBugID(networkDeviceID string, id string) (*ResponseComplianceGetBugAffectingTheNetworkDeviceByDeviceIDAndBugIDV1, *resty.Response, error) {
	return s.GetBugAffectingTheNetworkDeviceByDeviceIDAndBugIDV1(networkDeviceID, id)
}

// Alias Function
/*
This method acts as an alias for the method `RunComplianceV1`
*/
func (s *ComplianceService) RunCompliance(requestComplianceRunComplianceV1 *RequestComplianceRunComplianceV1) (*ResponseComplianceRunComplianceV1, *resty.Response, error) {
	return s.RunComplianceV1(requestComplianceRunComplianceV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetFieldNoticeNetworkDevicesV1`
*/
func (s *ComplianceService) GetFieldNoticeNetworkDevices(GetFieldNoticeNetworkDevicesV1QueryParams *GetFieldNoticeNetworkDevicesV1QueryParams) (*ResponseComplianceGetFieldNoticeNetworkDevicesV1, *resty.Response, error) {
	return s.GetFieldNoticeNetworkDevicesV1(GetFieldNoticeNetworkDevicesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `CreatesATrialForFieldNoticesDetectionOnNetworkDevicesV1`
*/
func (s *ComplianceService) CreatesATrialForFieldNoticesDetectionOnNetworkDevices() (*ResponseComplianceCreatesATrialForFieldNoticesDetectionOnNetworkDevicesV1, *resty.Response, error) {
	return s.CreatesATrialForFieldNoticesDetectionOnNetworkDevicesV1()
}

// Alias Function
/*
This method acts as an alias for the method `CreatesATrialForSecurityAdvisoriesDetectionOnNetworkDevicesV1`
*/
func (s *ComplianceService) CreatesATrialForSecurityAdvisoriesDetectionOnNetworkDevices() (*ResponseComplianceCreatesATrialForSecurityAdvisoriesDetectionOnNetworkDevicesV1, *resty.Response, error) {
	return s.CreatesATrialForSecurityAdvisoriesDetectionOnNetworkDevicesV1()
}

// Alias Function
/*
This method acts as an alias for the method `GetSecurityAdvisoryNetworkDevicesV1`
*/
func (s *ComplianceService) GetSecurityAdvisoryNetworkDevices(GetSecurityAdvisoryNetworkDevicesV1QueryParams *GetSecurityAdvisoryNetworkDevicesV1QueryParams) (*ResponseComplianceGetSecurityAdvisoryNetworkDevicesV1, *resty.Response, error) {
	return s.GetSecurityAdvisoryNetworkDevicesV1(GetSecurityAdvisoryNetworkDevicesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `TriggersAFieldNoticesScanForTheSupportedNetworkDevicesV1`
*/
func (s *ComplianceService) TriggersAFieldNoticesScanForTheSupportedNetworkDevices(TriggersAFieldNoticesScanForTheSupportedNetworkDevicesV1QueryParams *TriggersAFieldNoticesScanForTheSupportedNetworkDevicesV1QueryParams) (*ResponseComplianceTriggersAFieldNoticesScanForTheSupportedNetworkDevicesV1, *resty.Response, error) {
	return s.TriggersAFieldNoticesScanForTheSupportedNetworkDevicesV1(TriggersAFieldNoticesScanForTheSupportedNetworkDevicesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetTrialDetailsForFieldNoticesDetectionOnNetworkDevicesV1`
*/
func (s *ComplianceService) GetTrialDetailsForFieldNoticesDetectionOnNetworkDevices() (*ResponseComplianceGetTrialDetailsForFieldNoticesDetectionOnNetworkDevicesV1, *resty.Response, error) {
	return s.GetTrialDetailsForFieldNoticesDetectionOnNetworkDevicesV1()
}

// Alias Function
/*
This method acts as an alias for the method `GetFieldNoticeNetworkDeviceForTheNoticeByNetworkDeviceIDV1`
*/
func (s *ComplianceService) GetFieldNoticeNetworkDeviceForTheNoticeByNetworkDeviceID(id string, networkDeviceID string) (*ResponseComplianceGetFieldNoticeNetworkDeviceForTheNoticeByNetworkDeviceIDV1, *resty.Response, error) {
	return s.GetFieldNoticeNetworkDeviceForTheNoticeByNetworkDeviceIDV1(id, networkDeviceID)
}
