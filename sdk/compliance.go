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
	ComplianceStatus string `url:"complianceStatus,omitempty"` //Specify "Compliance status(es)" separated by commas. The Compliance status can be 'COMPLIANT', 'NON_COMPLIANT', 'IN_PROGRESS', 'NOT_AVAILABLE', 'NOT_APPLICABLE', 'ERROR'.
	DeviceUUID       string `url:"deviceUuid,omitempty"`       //Comma separated 'Device Ids'
}
type GetComplianceStatusCountV1QueryParams struct {
	ComplianceStatus string `url:"complianceStatus,omitempty"` //Specify "Compliance status(es)" separated by commas. The Compliance status can be 'COMPLIANT', 'NON_COMPLIANT', 'IN_PROGRESS', 'NOT_AVAILABLE', 'NOT_APPLICABLE', 'ERROR'.
}
type GetComplianceDetailV1QueryParams struct {
	ComplianceType   string  `url:"complianceType,omitempty"`   //Specify "Compliance type(s)" in commas. The Compliance type can be 'NETWORK_PROFILE', 'IMAGE', 'FABRIC', 'APPLICATION_VISIBILITY', 'FABRIC', RUNNING_CONFIG', 'NETWORK_SETTINGS', 'WORKFLOW' , 'EoX'.
	ComplianceStatus string  `url:"complianceStatus,omitempty"` //Specify "Compliance status(es)" in commas. The Compliance status can be 'COMPLIANT', 'NON_COMPLIANT', 'IN_PROGRESS', 'NOT_AVAILABLE', 'NOT_APPLICABLE', 'ERROR'.
	DeviceUUID       string  `url:"deviceUuid,omitempty"`       //Comma separated "Device Id(s)"
	Offset           float64 `url:"offset,omitempty"`           //offset/starting row
	Limit            float64 `url:"limit,omitempty"`            //Number of records to be retrieved
}
type GetComplianceDetailCountV1QueryParams struct {
	ComplianceType   string `url:"complianceType,omitempty"`   //Specify "Compliance type(s)" separated by commas. The Compliance type can be 'APPLICATION_VISIBILITY', 'EoX', 'FABRIC', 'IMAGE', 'NETWORK_PROFILE', 'NETWORK_SETTINGS', 'PSIRT', 'RUNNING_CONFIG', 'WORKFLOW'.
	ComplianceStatus string `url:"complianceStatus,omitempty"` //Specify "Compliance status(es)" separated by commas. The Compliance status can be 'COMPLIANT', 'NON_COMPLIANT', 'IN_PROGRESS', 'NOT_AVAILABLE', 'NOT_APPLICABLE', 'ERROR'.
}
type ComplianceDetailsOfDeviceV1QueryParams struct {
	Category       string `url:"category,omitempty"`       //category can have any value among 'INTENT', 'RUNNING_CONFIG' , 'IMAGE' , 'PSIRT' , 'DESIGN_OOD' , 'EoX' , 'NETWORK_SETTINGS'
	ComplianceType string `url:"complianceType,omitempty"` //Specify "Compliance type(s)" separated by commas. The Compliance type can be 'APPLICATION_VISIBILITY', 'EoX', 'FABRIC', 'IMAGE', 'NETWORK_PROFILE', 'NETWORK_SETTINGS', 'PSIRT', 'RUNNING_CONFIG', 'WORKFLOW'.
	DiffList       bool   `url:"diffList,omitempty"`       //diff list [ pass true to fetch the diff list ]
}
type GetConfigTaskDetailsV1QueryParams struct {
	ParentTaskID string `url:"parentTaskId,omitempty"` //task Id
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
	ComplianceType string   `json:"complianceType,omitempty"` // Compliance type corresponds to a tile on the UI. Will be one of NETWORK_PROFILE, IMAGE, APPLICATION_VISIBILITY, FABRIC, PSIRT, RUNNING_CONFIG, NETWORK_SETTINGS, WORKFLOW, or EoX.
	LastSyncTime   *float64 `json:"lastSyncTime,omitempty"`   // Timestamp when the status changed from different value to the current value.
	DeviceUUID     string   `json:"deviceUuid,omitempty"`     // UUID of the device.
	DisplayName    string   `json:"displayName,omitempty"`    // User friendly name for the configuration.
	Status         string   `json:"status,omitempty"`         // Current status of compliance for the complianceType. Will be one of COMPLIANT, NON_COMPLIANT, ERROR, IN_PROGRESS, NOT_APPLICABLE, NOT_AVAILABLE, COMPLIANT_WARNING, REMEDIATION_IN_PROGRESS, or ABORTED.
	Category       string   `json:"category,omitempty"`       // category can have any value among 'INTENT'(mapped to compliance types: NETWORK_SETTINGS,NETWORK_PROFILE,WORKFLOW,FABRIC,APPLICATION_VISIBILITY), 'RUNNING_CONFIG' , 'IMAGE' , 'PSIRT' , 'EoX' , 'NETWORK_SETTINGS'.
	LastUpdateTime *float64 `json:"lastUpdateTime,omitempty"` // Timestamp when the latest compliance checks ran.
	State          string   `json:"state,omitempty"`          // State of latest compliance check for the complianceType. Will be one of SUCCESS, FAILED, or IN_PROGRESS.
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

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-compliance-status-v1
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

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-compliance-status-count-v1
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

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-compliance-detail-v1
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

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-compliance-detail-count-v1
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


Documentation Link: https://developer.cisco.com/docs/dna-center/#!device-compliance-status-v1
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

Documentation Link: https://developer.cisco.com/docs/dna-center/#!compliance-details-of-device-v1
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

//GetConfigTaskDetailsV1 Get config task details - 8183-1a90-4788-b8c5
/* Returns a config task result details by specified id


@param GetConfigTaskDetailsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-config-task-details-v1
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

//RunComplianceV1 Run Compliance - f6ae-c8a7-4428-a9ff
/* Run compliance check for device(s).



Documentation Link: https://developer.cisco.com/docs/dna-center/#!run-compliance-v1
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


Documentation Link: https://developer.cisco.com/docs/dna-center/#!compliance-remediation-v1
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

//CommitDeviceConfigurationV1 Commit device configuration - 53a3-5a70-4e3b-87b5
/* This operation would commit device running configuration to startup by issuing "write memory" to device



Documentation Link: https://developer.cisco.com/docs/dna-center/#!commit-device-configuration-v1
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

// Alias Function
/*
This method acts as an alias for the method `GetComplianceDetailCountV1`
*/
func (s *ComplianceService) GetComplianceDetailCount(GetComplianceDetailCountV1QueryParams *GetComplianceDetailCountV1QueryParams) (*ResponseComplianceGetComplianceDetailCountV1, *resty.Response, error) {
	return s.GetComplianceDetailCountV1(GetComplianceDetailCountV1QueryParams)
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
This method acts as an alias for the method `DeviceComplianceStatusV1`
*/
func (s *ComplianceService) DeviceComplianceStatus(deviceUUID string) (*ResponseComplianceDeviceComplianceStatusV1, *resty.Response, error) {
	return s.DeviceComplianceStatusV1(deviceUUID)
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
This method acts as an alias for the method `ComplianceDetailsOfDeviceV1`
*/
func (s *ComplianceService) ComplianceDetailsOfDevice(deviceUUID string, ComplianceDetailsOfDeviceV1QueryParams *ComplianceDetailsOfDeviceV1QueryParams) (*ResponseComplianceComplianceDetailsOfDeviceV1, *resty.Response, error) {
	return s.ComplianceDetailsOfDeviceV1(deviceUUID, ComplianceDetailsOfDeviceV1QueryParams)
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
This method acts as an alias for the method `RunComplianceV1`
*/
func (s *ComplianceService) RunCompliance(requestComplianceRunComplianceV1 *RequestComplianceRunComplianceV1) (*ResponseComplianceRunComplianceV1, *resty.Response, error) {
	return s.RunComplianceV1(requestComplianceRunComplianceV1)
}
