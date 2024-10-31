package catalyst

import (
	"fmt"
	"net/http"

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

Documentation Link: https://developer.cisco.com/docs/dna-center/#!return-list-of-replacement-devices-with-replacement-details-v1
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

Documentation Link: https://developer.cisco.com/docs/dna-center/#!return-replacement-devices-count-v1
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

//MarkDeviceForReplacementV1 Mark device for replacement - 64b9-dad0-403a-aca1
/* Marks device for replacement



Documentation Link: https://developer.cisco.com/docs/dna-center/#!mark-device-for-replacement-v1
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



Documentation Link: https://developer.cisco.com/docs/dna-center/#!deploy-device-replacement-workflow-v1
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
func (s *DeviceReplacementService) ReturnListOfReplacementDevicesWithReplacementDetails(ReturnListOfReplacementDevicesWithReplacementDetailsV1QueryParams *ReturnListOfReplacementDevicesWithReplacementDetailsV1QueryParams) (*ResponseDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetailsV1, *resty.Response, error) {
	return s.ReturnListOfReplacementDevicesWithReplacementDetailsV1(ReturnListOfReplacementDevicesWithReplacementDetailsV1QueryParams)
}

// Alias Function
func (s *DeviceReplacementService) UnmarkDeviceForReplacement(requestDeviceReplacementUnMarkDeviceForReplacementV1 *RequestDeviceReplacementUnmarkDeviceForReplacementV1) (*ResponseDeviceReplacementUnmarkDeviceForReplacementV1, *resty.Response, error) {
	return s.UnmarkDeviceForReplacementV1(requestDeviceReplacementUnMarkDeviceForReplacementV1)
}

// Alias Function
func (s *DeviceReplacementService) ReturnReplacementDevicesCount(ReturnReplacementDevicesCountV1QueryParams *ReturnReplacementDevicesCountV1QueryParams) (*ResponseDeviceReplacementReturnReplacementDevicesCountV1, *resty.Response, error) {
	return s.ReturnReplacementDevicesCountV1(ReturnReplacementDevicesCountV1QueryParams)
}

// Alias Function
func (s *DeviceReplacementService) MarkDeviceForReplacement(requestDeviceReplacementMarkDeviceForReplacementV1 *RequestDeviceReplacementMarkDeviceForReplacementV1) (*ResponseDeviceReplacementMarkDeviceForReplacementV1, *resty.Response, error) {
	return s.MarkDeviceForReplacementV1(requestDeviceReplacementMarkDeviceForReplacementV1)
}

// Alias Function
func (s *DeviceReplacementService) DeployDeviceReplacementWorkflow(requestDeviceReplacementDeployDeviceReplacementWorkflowV1 *RequestDeviceReplacementDeployDeviceReplacementWorkflowV1) (*ResponseDeviceReplacementDeployDeviceReplacementWorkflowV1, *resty.Response, error) {
	return s.DeployDeviceReplacementWorkflowV1(requestDeviceReplacementDeployDeviceReplacementWorkflowV1)
}
