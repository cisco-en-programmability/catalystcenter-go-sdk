package catalyst

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type DeviceReplacementService service

type ReturnListOfReplacementDevicesWithReplacementDetailsV1QueryParams struct {
	FaultyDeviceName              string   `url:"faultyDeviceName,omitempty"`              //Faulty Device Name
	FaultyDevicePlatform          string   `url:"faultyDevicePlatform,omitempty"`          //Faulty Device Platform
	ReplacementDevicePlatform     string   `url:"replacementDevicePlatform,omitempty"`     //Replacement Device Platform
	FaultyDeviceSerialNumber      string   `url:"faultyDeviceSerialNumber,omitempty"`      //Faulty Device Serial Number
	ReplacementDeviceSerialNumber string   `url:"replacementDeviceSerialNumber,omitempty"` //Replacement Device Serial Number
	ReplacementStatus             []string `url:"replacementStatus,omitempty"`             //Device Replacement status [READY-FOR-REPLACEMENT, REPLACEMENT-IN-PROGRESS, REPLACEMENT-SCHEDULED, REPLACED, ERROR, NETWORK_READINESS_REQUESTED, NETWORK_READINESS_FAILED]
	Family                        []string `url:"family,omitempty"`                        //List of families[Routers, Switches and Hubs, AP]
	SortBy                        string   `url:"sortBy,omitempty"`                        //SortBy this field. SortBy is mandatory when order is used.
	SortOrder                     string   `url:"sortOrder,omitempty"`                     //Order on displayName[ASC,DESC]
	Offset                        int      `url:"offset,omitempty"`                        //offset
	Limit                         int      `url:"limit,omitempty"`                         //limit
}
type ReturnReplacementDevicesCountV1QueryParams struct {
	ReplacementStatus []string `url:"replacementStatus,omitempty"` //Device Replacement status list[READY-FOR-REPLACEMENT, REPLACEMENT-IN-PROGRESS, REPLACEMENT-SCHEDULED, REPLACED, ERROR]
}
type RetrieveTheStatusOfAllTheDeviceReplacementWorkflowsV1QueryParams struct {
	Family                        string  `url:"family,omitempty"`                        //Faulty device family.
	FaultyDeviceName              string  `url:"faultyDeviceName,omitempty"`              //Faulty device name.
	FaultyDevicePlatform          string  `url:"faultyDevicePlatform,omitempty"`          //Faulty device platform.
	FaultyDeviceSerialNumber      string  `url:"faultyDeviceSerialNumber,omitempty"`      //Faulty device serial number.
	ReplacementDevicePlatform     string  `url:"replacementDevicePlatform,omitempty"`     //Replacement device platform.
	ReplacementDeviceSerialNumber string  `url:"replacementDeviceSerialNumber,omitempty"` //Replacement device serial number.
	ReplacementStatus             string  `url:"replacementStatus,omitempty"`             //Device replacement status. Available values : MARKED_FOR_REPLACEMENT, NETWORK_READINESS_REQUESTED, NETWORK_READINESS_FAILED, READY_FOR_REPLACEMENT, REPLACEMENT_SCHEDULED, REPLACEMENT_IN_PROGRESS, REPLACED, ERROR. Replacement status: 'MARKED_FOR_REPLACEMENT' - The faulty device has been marked for replacement. 'NETWORK_READINESS_REQUESTED' - Initiated steps to shut down neighboring device interfaces and create a DHCP server on the uplink neighbor if the faulty device is part of a fabric setup. 'NETWORK_READINESS_FAILED' - Preparation of the network failed. Neighboring device interfaces were not shut down, and the DHCP server on the uplink neighbor was not created. 'READY_FOR_REPLACEMENT' - The network is prepared for the faulty device replacement. Neighboring device interfaces are shut down, and the DHCP server on the uplink neighbor is set up. 'REPLACEMENT_SCHEDULED' - Device replacement has been scheduled. 'REPLACEMENT_IN_PROGRESS' - Device replacement is currently in progress. 'REPLACED' - Device replacement was successful. 'ERROR' - Device replacement has failed.
	Offset                        float64 `url:"offset,omitempty"`                        //The first record to show for this page; the first record is numbered 1.
	Limit                         float64 `url:"limit,omitempty"`                         //The number of records to show for this page. Maximum value can be 500.
	SortBy                        string  `url:"sortBy,omitempty"`                        //A property within the response to sort by. Available values : id, creationTime, family, faultyDeviceId, fautyDeviceName, faultyDevicePlatform, faultyDeviceSerialNumber, replacementDevicePlatform, replacementDeviceSerialNumber, replacementTime.
	SortOrder                     string  `url:"sortOrder,omitempty"`                     //Whether ascending or descending order should be used to sort the response. Available values : ASC, DESC
}

type ResponseDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetailsV1 struct {
	Response *[]ResponseDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetailsV1Response `json:"response,omitempty"` //
	Version  string                                                                                     `json:"version,omitempty"`  //
}
type ResponseDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetailsV1Response struct {
	CreationTime                  *int   `json:"creationTime,omitempty"`                  // Date and time of marking the device for replacement
	Family                        string `json:"family,omitempty"`                        // Faulty device family
	FaultyDeviceID                string `json:"faultyDeviceId,omitempty"`                // Unique identifier of the faulty device
	FaultyDeviceName              string `json:"faultyDeviceName,omitempty"`              // Faulty device name
	FaultyDevicePlatform          string `json:"faultyDevicePlatform,omitempty"`          // Faulty device platform
	FaultyDeviceSerialNumber      string `json:"faultyDeviceSerialNumber,omitempty"`      // Faulty device serial number
	ID                            string `json:"id,omitempty"`                            // Unique identifier of the device replacement resource
	NeighbourDeviceID             string `json:"neighbourDeviceId,omitempty"`             // Unique identifier of the neighbor device to create the DHCP server
	NetworkReadinessTaskID        string `json:"networkReadinessTaskId,omitempty"`        // Unique identifier of network readiness task
	ReplacementDevicePlatform     string `json:"replacementDevicePlatform,omitempty"`     // Replacement device platform
	ReplacementDeviceSerialNumber string `json:"replacementDeviceSerialNumber,omitempty"` // Replacement device serial number
	ReplacementStatus             string `json:"replacementStatus,omitempty"`             // Device Replacement status
	ReplacementTime               *int   `json:"replacementTime,omitempty"`               // Date and time of device replacement
	WorkflowID                    string `json:"workflowId,omitempty"`                    // Unique identifier of the device replacement workflow
	WorkflowFailedStep            string `json:"workflowFailedStep,omitempty"`            // Step in which the device replacement failed
	ReadinesscheckTaskID          string `json:"readinesscheckTaskId,omitempty"`          // Unique identifier of the readiness check task for the replacement device
}
type ResponseDeviceReplacementUnmarkDeviceForReplacementV1 struct {
	Response *ResponseDeviceReplacementUnmarkDeviceForReplacementV1Response `json:"response,omitempty"` //
	Version  string                                                         `json:"version,omitempty"`  //
}
type ResponseDeviceReplacementUnmarkDeviceForReplacementV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDeviceReplacementMarkDeviceForReplacementV1 struct {
	Response *ResponseDeviceReplacementMarkDeviceForReplacementV1Response `json:"response,omitempty"` //
	Version  string                                                       `json:"version,omitempty"`  //
}
type ResponseDeviceReplacementMarkDeviceForReplacementV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDeviceReplacementReturnReplacementDevicesCountV1 struct {
	Response *int   `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}
type ResponseDeviceReplacementDeployDeviceReplacementWorkflowV1 struct {
	Response *ResponseDeviceReplacementDeployDeviceReplacementWorkflowV1Response `json:"response,omitempty"` //
	Version  string                                                              `json:"version,omitempty"`  //
}
type ResponseDeviceReplacementDeployDeviceReplacementWorkflowV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDeviceReplacementRetrieveTheStatusOfAllTheDeviceReplacementWorkflowsV1 struct {
	Response *[]ResponseDeviceReplacementRetrieveTheStatusOfAllTheDeviceReplacementWorkflowsV1Response `json:"response,omitempty"` //
	Version  string                                                                                    `json:"version,omitempty"`  // The version of the response.
}
type ResponseDeviceReplacementRetrieveTheStatusOfAllTheDeviceReplacementWorkflowsV1Response struct {
	CreationTime                  *int                                                                                            `json:"creationTime,omitempty"`                  // Time of marking the device for replacement in Unix epoch time in milliseconds
	Family                        string                                                                                          `json:"family,omitempty"`                        // Faulty device family
	FaultyDeviceID                string                                                                                          `json:"faultyDeviceId,omitempty"`                // Faulty device id
	FaultyDeviceName              string                                                                                          `json:"faultyDeviceName,omitempty"`              // Faulty device name
	FaultyDevicePlatform          string                                                                                          `json:"faultyDevicePlatform,omitempty"`          // Faulty device platform
	FaultyDeviceSerialNumber      string                                                                                          `json:"faultyDeviceSerialNumber,omitempty"`      // Faulty device serial number
	ID                            string                                                                                          `json:"id,omitempty"`                            // Unique identifier of the device replacement resource
	NeighborDeviceID              string                                                                                          `json:"neighborDeviceId,omitempty"`              // Unique identifier of the neighbor device to create the DHCP server
	ReplacementDevicePlatform     string                                                                                          `json:"replacementDevicePlatform,omitempty"`     // Replacement device platform
	ReplacementDeviceSerialNumber string                                                                                          `json:"replacementDeviceSerialNumber,omitempty"` // Replacement device serial number
	ReplacementStatus             string                                                                                          `json:"replacementStatus,omitempty"`             // Device Replacement status. 'MARKED_FOR_REPLACEMENT' - The faulty device has been marked for replacement. 'NETWORK_READINESS_REQUESTED' - Initiated steps to shut down neighboring device interfaces and create a DHCP server on the uplink neighbor if the faulty device is part of a fabric setup. 'NETWORK_READINESS_FAILED' - Preparation of the network failed. Neighboring device interfaces were not shut down, and the DHCP server on the uplink neighbor was not created. 'READY_FOR_REPLACEMENT' - The network is prepared for the faulty device replacement. Neighboring device interfaces are shut down, and the DHCP server on the uplink neighbor is set up. 'REPLACEMENT_SCHEDULED' - Device replacement has been scheduled. 'REPLACEMENT_IN_PROGRESS' - Device replacement is currently in progress. 'REPLACED' - Device replacement was successful. 'ERROR' - Device replacement has failed.
	ReplacementTime               *int                                                                                            `json:"replacementTime,omitempty"`               // The Unix epoch time in milliseconds at which the device was replaced successfully
	Workflow                      *ResponseDeviceReplacementRetrieveTheStatusOfAllTheDeviceReplacementWorkflowsV1ResponseWorkflow `json:"workflow,omitempty"`                      //
}
type ResponseDeviceReplacementRetrieveTheStatusOfAllTheDeviceReplacementWorkflowsV1ResponseWorkflow struct {
	ID             string                                                                                                 `json:"id,omitempty"`             // Workflow id
	Name           string                                                                                                 `json:"name,omitempty"`           // Name of the workflow
	WorkflowStatus string                                                                                                 `json:"workflowStatus,omitempty"` // Workflow status. 'RUNNING' - Workflow is currently in progress. 'SUCCESS' - Workflow completed successfully. 'FAILED' - Workflow completed with failure.
	StartTime      *int                                                                                                   `json:"startTime,omitempty"`      // Start time of the workflow in Unix epoch time in milliseconds
	EndTime        *int                                                                                                   `json:"endTime,omitempty"`        // Completion time of the workflow in Unix epoch time in milliseconds
	Steps          *[]ResponseDeviceReplacementRetrieveTheStatusOfAllTheDeviceReplacementWorkflowsV1ResponseWorkflowSteps `json:"steps,omitempty"`          //
}
type ResponseDeviceReplacementRetrieveTheStatusOfAllTheDeviceReplacementWorkflowsV1ResponseWorkflowSteps struct {
	Name          string `json:"name,omitempty"`          // Workflow step name
	Status        string `json:"status,omitempty"`        // Workflow step status. 'INIT' - Workflow step has not started execution. 'RUNNING' - Workflow step is currently in progress. 'SUCCESS' - Workflow step completed successfully. 'FAILED' - Workflow step completed with failure. 'ABORTED' - Workflow step aborted execution due to failure of the previous step. 'TIMEOUT' - Workflow step timedout to complete execution.
	StatusMessage string `json:"statusMessage,omitempty"` // Detailed status message for the step
	StartTime     *int   `json:"startTime,omitempty"`     // Start time of the workflow step in Unix epoch time in milliseconds
	EndTime       *int   `json:"endTime,omitempty"`       // Completion time of the workflow step in Unix epoch time in milliseconds
}
type ResponseDeviceReplacementRetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDeviceV1 struct {
	Response *ResponseDeviceReplacementRetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDeviceV1Response `json:"response,omitempty"` //
	Version  string                                                                                                                          `json:"version,omitempty"`  // The version of the response.
}
type ResponseDeviceReplacementRetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDeviceV1Response struct {
	CreationTime                  *int                                                                                                                                    `json:"creationTime,omitempty"`                  // Time of marking the device for replacement in Unix epoch time in milliseconds
	Family                        string                                                                                                                                  `json:"family,omitempty"`                        // Faulty device family
	FaultyDeviceID                string                                                                                                                                  `json:"faultyDeviceId,omitempty"`                // Faulty device id
	FaultyDeviceName              string                                                                                                                                  `json:"faultyDeviceName,omitempty"`              // Faulty device name
	FaultyDevicePlatform          string                                                                                                                                  `json:"faultyDevicePlatform,omitempty"`          // Faulty device platform
	FaultyDeviceSerialNumber      string                                                                                                                                  `json:"faultyDeviceSerialNumber,omitempty"`      // Faulty device serial number
	ID                            string                                                                                                                                  `json:"id,omitempty"`                            // Unique identifier of the device replacement resource
	NeighborDeviceID              string                                                                                                                                  `json:"neighborDeviceId,omitempty"`              // Unique identifier of the neighbor device to create the DHCP server
	ReplacementDevicePlatform     string                                                                                                                                  `json:"replacementDevicePlatform,omitempty"`     // Replacement device platform
	ReplacementDeviceSerialNumber string                                                                                                                                  `json:"replacementDeviceSerialNumber,omitempty"` // Replacement device serial number
	ReplacementStatus             string                                                                                                                                  `json:"replacementStatus,omitempty"`             // Device Replacement status. 'MARKED_FOR_REPLACEMENT' - The faulty device has been marked for replacement. 'NETWORK_READINESS_REQUESTED' - Initiated steps to shut down neighboring device interfaces and create a DHCP server on the uplink neighbor if the faulty device is part of a fabric setup. 'NETWORK_READINESS_FAILED' - Preparation of the network failed. Neighboring device interfaces were not shut down, and the DHCP server on the uplink neighbor was not created. 'READY_FOR_REPLACEMENT' - The network is prepared for the faulty device replacement. Neighboring device interfaces are shut down, and the DHCP server on the uplink neighbor is set up. 'REPLACEMENT_SCHEDULED' - Device replacement has been scheduled. 'REPLACEMENT_IN_PROGRESS' - Device replacement is currently in progress. 'REPLACED' - Device replacement was successful. 'ERROR' - Device replacement has failed.
	ReplacementTime               *int                                                                                                                                    `json:"replacementTime,omitempty"`               // The Unix epoch time in milliseconds at which the device was replaced successfully
	Workflow                      *ResponseDeviceReplacementRetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDeviceV1ResponseWorkflow `json:"workflow,omitempty"`                      //
}
type ResponseDeviceReplacementRetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDeviceV1ResponseWorkflow struct {
	ID             string                                                                                                                                         `json:"id,omitempty"`             // Workflow id
	Name           string                                                                                                                                         `json:"name,omitempty"`           // Name of the workflow
	WorkflowStatus string                                                                                                                                         `json:"workflowStatus,omitempty"` // Workflow status. 'RUNNING' - Workflow is currently in progress. 'SUCCESS' - Workflow completed successfully. 'FAILED' - Workflow completed with failure.
	StartTime      *int                                                                                                                                           `json:"startTime,omitempty"`      // Start time of the workflow in Unix epoch time in milliseconds
	EndTime        *int                                                                                                                                           `json:"endTime,omitempty"`        // Completion time of the workflow in Unix epoch time in milliseconds
	Steps          *[]ResponseDeviceReplacementRetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDeviceV1ResponseWorkflowSteps `json:"steps,omitempty"`          //
}
type ResponseDeviceReplacementRetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDeviceV1ResponseWorkflowSteps struct {
	Name          string `json:"name,omitempty"`          // Workflow step name
	Status        string `json:"status,omitempty"`        // Workflow step status. 'INIT' - Workflow step has not started execution. 'RUNNING' - Workflow step is currently in progress. 'SUCCESS' - Workflow step completed successfully. 'FAILED' - Workflow step completed with failure. 'ABORTED' - Workflow step aborted execution due to failure of the previous step. 'TIMEOUT' - Workflow step timedout to complete execution.
	StatusMessage string `json:"statusMessage,omitempty"` // Detailed status message for the step
	StartTime     *int   `json:"startTime,omitempty"`     // Start time of the workflow step in Unix epoch time in milliseconds
	EndTime       *int   `json:"endTime,omitempty"`       // Completion time of the workflow step in Unix epoch time in milliseconds
}
type RequestDeviceReplacementUnmarkDeviceForReplacementV1 []RequestItemDeviceReplacementUnmarkDeviceForReplacementV1 // Array of RequestDeviceReplacementUnMarkDeviceForReplacementV1
type RequestItemDeviceReplacementUnmarkDeviceForReplacementV1 struct {
	CreationTime                  *int   `json:"creationTime,omitempty"`                  // Date and time of marking the device for replacement
	Family                        string `json:"family,omitempty"`                        // Faulty device family
	FaultyDeviceID                string `json:"faultyDeviceId,omitempty"`                // Unique identifier of the faulty device
	FaultyDeviceName              string `json:"faultyDeviceName,omitempty"`              // Faulty device name
	FaultyDevicePlatform          string `json:"faultyDevicePlatform,omitempty"`          // Faulty device platform
	FaultyDeviceSerialNumber      string `json:"faultyDeviceSerialNumber,omitempty"`      // Faulty device serial number
	ID                            string `json:"id,omitempty"`                            // Unique identifier of the device replacement resource
	NeighbourDeviceID             string `json:"neighbourDeviceId,omitempty"`             // Unique identifier of the neighbor device to create the DHCP server
	NetworkReadinessTaskID        string `json:"networkReadinessTaskId,omitempty"`        // Unique identifier of network readiness task
	ReplacementDevicePlatform     string `json:"replacementDevicePlatform,omitempty"`     // Replacement device platform
	ReplacementDeviceSerialNumber string `json:"replacementDeviceSerialNumber,omitempty"` // Replacement device serial number
	ReplacementStatus             string `json:"replacementStatus,omitempty"`             // Device replacement status. Use NON-FAULTY to unmark the device for replacement.
	ReplacementTime               *int   `json:"replacementTime,omitempty"`               // Date and time of device replacement
	WorkflowID                    string `json:"workflowId,omitempty"`                    // Unique identifier of the device replacement workflow
}
type RequestDeviceReplacementMarkDeviceForReplacementV1 []RequestItemDeviceReplacementMarkDeviceForReplacementV1 // Array of RequestDeviceReplacementMarkDeviceForReplacementV1
type RequestItemDeviceReplacementMarkDeviceForReplacementV1 struct {
	CreationTime                  *int   `json:"creationTime,omitempty"`                  // Date and time of marking the device for replacement
	Family                        string `json:"family,omitempty"`                        // Faulty device family
	FaultyDeviceID                string `json:"faultyDeviceId,omitempty"`                // Unique identifier of the faulty device
	FaultyDeviceName              string `json:"faultyDeviceName,omitempty"`              // Faulty device name
	FaultyDevicePlatform          string `json:"faultyDevicePlatform,omitempty"`          // Faulty device platform
	FaultyDeviceSerialNumber      string `json:"faultyDeviceSerialNumber,omitempty"`      // Faulty device serial number
	ID                            string `json:"id,omitempty"`                            // Unique identifier of the device replacement resource
	NeighbourDeviceID             string `json:"neighbourDeviceId,omitempty"`             // Unique identifier of the neighbor device to create the DHCP server
	NetworkReadinessTaskID        string `json:"networkReadinessTaskId,omitempty"`        // Unique identifier of network readiness task
	ReplacementDevicePlatform     string `json:"replacementDevicePlatform,omitempty"`     // Replacement device platform
	ReplacementDeviceSerialNumber string `json:"replacementDeviceSerialNumber,omitempty"` // Replacement device serial number
	ReplacementStatus             string `json:"replacementStatus,omitempty"`             // Device replacement status. Use MARKED-FOR-REPLACEMENT to mark the device for replacement.
	ReplacementTime               *int   `json:"replacementTime,omitempty"`               // Date and time of device replacement
	WorkflowID                    string `json:"workflowId,omitempty"`                    // Unique identifier of the device replacement workflow
}
type RequestDeviceReplacementDeployDeviceReplacementWorkflowV1 struct {
	FaultyDeviceSerialNumber      string `json:"faultyDeviceSerialNumber,omitempty"`      // Faulty device serial number
	ReplacementDeviceSerialNumber string `json:"replacementDeviceSerialNumber,omitempty"` // Replacement device serial number
}

//ReturnListOfReplacementDevicesWithReplacementDetailsV1 Return list of replacement devices with replacement details - 809c-2956-4bc9-97d0
/* Get list of replacement devices with replacement details and it can filter replacement devices based on Faulty Device Name,Faulty Device Platform, Replacement Device Platform, Faulty Device Serial Number,Replacement Device Serial Number, Device Replacement status, Product Family.


@param ReturnListOfReplacementDevicesWithReplacementDetailsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!return-list-of-replacement-devices-with-replacement-details
*/
func (s *DeviceReplacementService) ReturnListOfReplacementDevicesWithReplacementDetailsV1(ReturnListOfReplacementDevicesWithReplacementDetailsV1QueryParams *ReturnListOfReplacementDevicesWithReplacementDetailsV1QueryParams) (*ResponseDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetailsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/device-replacement"

	queryString, _ := query.Values(ReturnListOfReplacementDevicesWithReplacementDetailsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetailsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReturnListOfReplacementDevicesWithReplacementDetailsV1(ReturnListOfReplacementDevicesWithReplacementDetailsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReturnListOfReplacementDevicesWithReplacementDetailsV1")
	}

	result := response.Result().(*ResponseDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetailsV1)
	return result, response, err

}

//ReturnReplacementDevicesCountV1 Return replacement devices count - 9eb8-4ba5-4929-a2a2
/* Get replacement devices count


@param ReturnReplacementDevicesCountV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!return-replacement-devices-count
*/
func (s *DeviceReplacementService) ReturnReplacementDevicesCountV1(ReturnReplacementDevicesCountV1QueryParams *ReturnReplacementDevicesCountV1QueryParams) (*ResponseDeviceReplacementReturnReplacementDevicesCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/device-replacement/count"

	queryString, _ := query.Values(ReturnReplacementDevicesCountV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDeviceReplacementReturnReplacementDevicesCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReturnReplacementDevicesCountV1(ReturnReplacementDevicesCountV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReturnReplacementDevicesCountV1")
	}

	result := response.Result().(*ResponseDeviceReplacementReturnReplacementDevicesCountV1)
	return result, response, err

}

//RetrieveTheStatusOfAllTheDeviceReplacementWorkflowsV1 Retrieve the status of all the device replacement workflows. - e6b8-0a1a-4929-a7a9
/* Retrieve the list of device replacements with replacement details. Filters can be applied based on faulty device name, faulty device platform, faulty device serial number, replacement device platform, replacement device serial number, device replacement status, device family.


@param RetrieveTheStatusOfAllTheDeviceReplacementWorkflowsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-the-status-of-all-the-device-replacement-workflows
*/
func (s *DeviceReplacementService) RetrieveTheStatusOfAllTheDeviceReplacementWorkflowsV1(RetrieveTheStatusOfAllTheDeviceReplacementWorkflowsV1QueryParams *RetrieveTheStatusOfAllTheDeviceReplacementWorkflowsV1QueryParams) (*ResponseDeviceReplacementRetrieveTheStatusOfAllTheDeviceReplacementWorkflowsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDeviceReplacements"

	queryString, _ := query.Values(RetrieveTheStatusOfAllTheDeviceReplacementWorkflowsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDeviceReplacementRetrieveTheStatusOfAllTheDeviceReplacementWorkflowsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveTheStatusOfAllTheDeviceReplacementWorkflowsV1(RetrieveTheStatusOfAllTheDeviceReplacementWorkflowsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveTheStatusOfAllTheDeviceReplacementWorkflowsV1")
	}

	result := response.Result().(*ResponseDeviceReplacementRetrieveTheStatusOfAllTheDeviceReplacementWorkflowsV1)
	return result, response, err

}

//RetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDeviceV1 Retrieve the status of device replacement workflow that replaces a faulty device with a replacement device. - 92ba-aa03-43c8-9d62
/* Fetches the status of the device replacement workflow for a given device replacement `id`. Invoke the API `/dna/intent/api/v1/networkDeviceReplacements` to `GET` the list of all device replacements and use the `id` field data as input to this API.


@param id id path parameter. Instance UUID of the device replacement


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-the-status-of-device-replacement-workflow-that-replaces-a-faulty-device-with-a-replacement-device
*/
func (s *DeviceReplacementService) RetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDeviceV1(id string) (*ResponseDeviceReplacementRetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDeviceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDeviceReplacements/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceReplacementRetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDeviceV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDeviceV1(id)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDeviceV1")
	}

	result := response.Result().(*ResponseDeviceReplacementRetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDeviceV1)
	return result, response, err

}

//MarkDeviceForReplacementV1 Mark device for replacement - 64b9-dad0-403a-aca1
/* Marks device for replacement



Documentation Link: https://developer.cisco.com/docs/dna-center/#!mark-device-for-replacement
*/
func (s *DeviceReplacementService) MarkDeviceForReplacementV1(requestDeviceReplacementMarkDeviceForReplacementV1 *RequestDeviceReplacementMarkDeviceForReplacementV1) (*ResponseDeviceReplacementMarkDeviceForReplacementV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/device-replacement"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceReplacementMarkDeviceForReplacementV1).
		SetResult(&ResponseDeviceReplacementMarkDeviceForReplacementV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.MarkDeviceForReplacementV1(requestDeviceReplacementMarkDeviceForReplacementV1)
		}

		return nil, response, fmt.Errorf("error with operation MarkDeviceForReplacementV1")
	}

	result := response.Result().(*ResponseDeviceReplacementMarkDeviceForReplacementV1)
	return result, response, err

}

//DeployDeviceReplacementWorkflowV1 Deploy device replacement workflow - 3faa-a994-4b49-bc9f
/* API to trigger RMA workflow that will replace faulty device with replacement device with same configuration and images



Documentation Link: https://developer.cisco.com/docs/dna-center/#!deploy-device-replacement-workflow
*/
func (s *DeviceReplacementService) DeployDeviceReplacementWorkflowV1(requestDeviceReplacementDeployDeviceReplacementWorkflowV1 *RequestDeviceReplacementDeployDeviceReplacementWorkflowV1) (*ResponseDeviceReplacementDeployDeviceReplacementWorkflowV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/device-replacement/workflow"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceReplacementDeployDeviceReplacementWorkflowV1).
		SetResult(&ResponseDeviceReplacementDeployDeviceReplacementWorkflowV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeployDeviceReplacementWorkflowV1(requestDeviceReplacementDeployDeviceReplacementWorkflowV1)
		}

		return nil, response, fmt.Errorf("error with operation DeployDeviceReplacementWorkflowV1")
	}

	result := response.Result().(*ResponseDeviceReplacementDeployDeviceReplacementWorkflowV1)
	return result, response, err

}

//UnmarkDeviceForReplacementV1 UnMark device for replacement - 4aba-ba75-489a-b24b
/* UnMarks device for replacement


 */
func (s *DeviceReplacementService) UnmarkDeviceForReplacementV1(requestDeviceReplacementUnMarkDeviceForReplacementV1 *RequestDeviceReplacementUnmarkDeviceForReplacementV1) (*ResponseDeviceReplacementUnmarkDeviceForReplacementV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/device-replacement"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceReplacementUnMarkDeviceForReplacementV1).
		SetResult(&ResponseDeviceReplacementUnmarkDeviceForReplacementV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UnmarkDeviceForReplacementV1(requestDeviceReplacementUnMarkDeviceForReplacementV1)
		}
		return nil, response, fmt.Errorf("error with operation UnmarkDeviceForReplacementV1")
	}

	result := response.Result().(*ResponseDeviceReplacementUnmarkDeviceForReplacementV1)
	return result, response, err

}

// Alias Function
/*
This method acts as an alias for the method `ReturnListOfReplacementDevicesWithReplacementDetailsV1`
*/
func (s *DeviceReplacementService) ReturnListOfReplacementDevicesWithReplacementDetails(ReturnListOfReplacementDevicesWithReplacementDetailsV1QueryParams *ReturnListOfReplacementDevicesWithReplacementDetailsV1QueryParams) (*ResponseDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetailsV1, *resty.Response, error) {
	return s.ReturnListOfReplacementDevicesWithReplacementDetailsV1(ReturnListOfReplacementDevicesWithReplacementDetailsV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `UnmarkDeviceForReplacementV1`
*/
func (s *DeviceReplacementService) UnmarkDeviceForReplacement(requestDeviceReplacementUnMarkDeviceForReplacementV1 *RequestDeviceReplacementUnmarkDeviceForReplacementV1) (*ResponseDeviceReplacementUnmarkDeviceForReplacementV1, *resty.Response, error) {
	return s.UnmarkDeviceForReplacementV1(requestDeviceReplacementUnMarkDeviceForReplacementV1)
}

// Alias Function
/*
This method acts as an alias for the method `ReturnReplacementDevicesCountV1`
*/
func (s *DeviceReplacementService) ReturnReplacementDevicesCount(ReturnReplacementDevicesCountV1QueryParams *ReturnReplacementDevicesCountV1QueryParams) (*ResponseDeviceReplacementReturnReplacementDevicesCountV1, *resty.Response, error) {
	return s.ReturnReplacementDevicesCountV1(ReturnReplacementDevicesCountV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `RetrieveTheStatusOfAllTheDeviceReplacementWorkflowsV1`
*/
func (s *DeviceReplacementService) RetrieveTheStatusOfAllTheDeviceReplacementWorkflows(RetrieveTheStatusOfAllTheDeviceReplacementWorkflowsV1QueryParams *RetrieveTheStatusOfAllTheDeviceReplacementWorkflowsV1QueryParams) (*ResponseDeviceReplacementRetrieveTheStatusOfAllTheDeviceReplacementWorkflowsV1, *resty.Response, error) {
	return s.RetrieveTheStatusOfAllTheDeviceReplacementWorkflowsV1(RetrieveTheStatusOfAllTheDeviceReplacementWorkflowsV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `RetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDeviceV1`
*/
func (s *DeviceReplacementService) RetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDevice(id string) (*ResponseDeviceReplacementRetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDeviceV1, *resty.Response, error) {
	return s.RetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDeviceV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `MarkDeviceForReplacementV1`
*/
func (s *DeviceReplacementService) MarkDeviceForReplacement(requestDeviceReplacementMarkDeviceForReplacementV1 *RequestDeviceReplacementMarkDeviceForReplacementV1) (*ResponseDeviceReplacementMarkDeviceForReplacementV1, *resty.Response, error) {
	return s.MarkDeviceForReplacementV1(requestDeviceReplacementMarkDeviceForReplacementV1)
}

// Alias Function
/*
This method acts as an alias for the method `DeployDeviceReplacementWorkflowV1`
*/
func (s *DeviceReplacementService) DeployDeviceReplacementWorkflow(requestDeviceReplacementDeployDeviceReplacementWorkflowV1 *RequestDeviceReplacementDeployDeviceReplacementWorkflowV1) (*ResponseDeviceReplacementDeployDeviceReplacementWorkflowV1, *resty.Response, error) {
	return s.DeployDeviceReplacementWorkflowV1(requestDeviceReplacementDeployDeviceReplacementWorkflowV1)
}
