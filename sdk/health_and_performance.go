package catalyst

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type HealthAndPerformanceService service

type RetrievesAllTheValidationSetsV1QueryParams struct {
	View string `url:"view,omitempty"` //When the query parameter `view=DETAIL` is passed, all validation sets and associated validations will be returned. When the query parameter `view=DEFAULT` is passed, only validation sets metadata will be returned.
}
type RetrievesTheListOfValidationWorkflowsV1QueryParams struct {
	StartTime float64 `url:"startTime,omitempty"` //Workflows started after the given time (as milliseconds since UNIX epoch).
	EndTime   float64 `url:"endTime,omitempty"`   //Workflows started before the given time (as milliseconds since UNIX epoch).
	RunStatus string  `url:"runStatus,omitempty"` //Execution status of the workflow. If the workflow is successfully submitted, runStatus is `PENDING`. If the workflow execution has started, runStatus is `IN_PROGRESS`. If the workflow executed is completed with all validations executed, runStatus is `COMPLETED`. If the workflow execution fails while running validations, runStatus is `FAILED`.
	Offset    float64 `url:"offset,omitempty"`    //The first record to show for this page; the first record is numbered 1.
	Limit     float64 `url:"limit,omitempty"`     //The number of records to show for this page.
}
type RetrievesTheCountOfValidationWorkflowsV1QueryParams struct {
	StartTime float64 `url:"startTime,omitempty"` //Workflows started after the given time (as milliseconds since UNIX epoch).
	EndTime   float64 `url:"endTime,omitempty"`   //Workflows started before the given time (as milliseconds since UNIX epoch).
	RunStatus string  `url:"runStatus,omitempty"` //Execution status of the workflow. If the workflow is successfully submitted, runStatus is `PENDING`. If the workflow execution has started, runStatus is `IN_PROGRESS`. If the workflow executed is completed with all validations executed, runStatus is `COMPLETED`. If the workflow execution fails while running validations, runStatus is `FAILED`.
}
type SystemHealthAPIV1QueryParams struct {
	Summary   bool    `url:"summary,omitempty"`   //Fetch the latest high severity event
	Domain    string  `url:"domain,omitempty"`    //Fetch system events with this domain. Possible values of domain are listed here : /dna/platform/app/consumer-portal/developer-toolkit/events
	Subdomain string  `url:"subdomain,omitempty"` //Fetch system events with this subdomain. Possible values of subdomain are listed here : /dna/platform/app/consumer-portal/developer-toolkit/events
	Limit     float64 `url:"limit,omitempty"`     //limit
	Offset    float64 `url:"offset,omitempty"`    //offset
}
type SystemHealthCountAPIV1QueryParams struct {
	Domain    string `url:"domain,omitempty"`    //Fetch system events with this domain. Possible values of domain are listed here : /dna/platform/app/consumer-portal/developer-toolkit/events
	Subdomain string `url:"subdomain,omitempty"` //Fetch system events with this subdomain. Possible values of subdomain are listed here : /dna/platform/app/consumer-portal/developer-toolkit/events
}
type SystemPerformanceAPIV1QueryParams struct {
	Kpi       string  `url:"kpi,omitempty"`       //Valid values: cpu,memory,network
	Function  string  `url:"function,omitempty"`  //Valid values: sum,average,max
	StartTime float64 `url:"startTime,omitempty"` //This is the epoch start time in milliseconds from which performance indicator need to be fetched
	EndTime   float64 `url:"endTime,omitempty"`   //This is the epoch end time in milliseconds upto which performance indicator need to be fetched
}
type SystemPerformanceHistoricalAPIV1QueryParams struct {
	Kpi       string  `url:"kpi,omitempty"`       //Fetch historical data for this kpi. Valid values: cpu,memory,network
	StartTime float64 `url:"startTime,omitempty"` //This is the epoch start time in milliseconds from which performance indicator need to be fetched
	EndTime   float64 `url:"endTime,omitempty"`   //This is the epoch end time in milliseconds upto which performance indicator need to be fetched
}

type ResponseHealthAndPerformanceRetrievesAllTheValidationSetsV1 struct {
	Response *[]ResponseHealthAndPerformanceRetrievesAllTheValidationSetsV1Response `json:"response,omitempty"` //
	Version  string                                                                 `json:"version,omitempty"`  // The version of the response
}
type ResponseHealthAndPerformanceRetrievesAllTheValidationSetsV1Response struct {
	ID               string                                                                                 `json:"id,omitempty"`               // Validation set id
	Name             string                                                                                 `json:"name,omitempty"`             // Name of the validation set
	Description      string                                                                                 `json:"description,omitempty"`      // Description of the validation set
	Version          string                                                                                 `json:"version,omitempty"`          // Version of the validation set
	ValidationGroups *[]ResponseHealthAndPerformanceRetrievesAllTheValidationSetsV1ResponseValidationGroups `json:"validationGroups,omitempty"` //
}
type ResponseHealthAndPerformanceRetrievesAllTheValidationSetsV1ResponseValidationGroups struct {
	Name        string                                                                                            `json:"name,omitempty"`        // Name of the validation group
	ID          string                                                                                            `json:"id,omitempty"`          // Validation group id
	Description string                                                                                            `json:"description,omitempty"` // Description of the validation group
	Validations *[]ResponseHealthAndPerformanceRetrievesAllTheValidationSetsV1ResponseValidationGroupsValidations `json:"validations,omitempty"` //
}
type ResponseHealthAndPerformanceRetrievesAllTheValidationSetsV1ResponseValidationGroupsValidations struct {
	ID   string `json:"id,omitempty"`   // Validation id
	Name string `json:"name,omitempty"` // Name of the validation
}
type ResponseHealthAndPerformanceRetrievesValidationDetailsForAValidationSetV1 struct {
	Response *ResponseHealthAndPerformanceRetrievesValidationDetailsForAValidationSetV1Response `json:"response,omitempty"` //
	Version  string                                                                             `json:"version,omitempty"`  // The version of the response
}
type ResponseHealthAndPerformanceRetrievesValidationDetailsForAValidationSetV1Response struct {
	ID               string                                                                                               `json:"id,omitempty"`               // Validation set id
	Name             string                                                                                               `json:"name,omitempty"`             // Name of the validation set
	Description      string                                                                                               `json:"description,omitempty"`      // Description of the validation set
	Version          string                                                                                               `json:"version,omitempty"`          // Version of validation set
	ValidationGroups *[]ResponseHealthAndPerformanceRetrievesValidationDetailsForAValidationSetV1ResponseValidationGroups `json:"validationGroups,omitempty"` //
}
type ResponseHealthAndPerformanceRetrievesValidationDetailsForAValidationSetV1ResponseValidationGroups struct {
	Name        string                                                                                                          `json:"name,omitempty"`        // Name of the validation group
	ID          string                                                                                                          `json:"id,omitempty"`          // Validation group id
	Description string                                                                                                          `json:"description,omitempty"` // Description of the validation group
	Validations *[]ResponseHealthAndPerformanceRetrievesValidationDetailsForAValidationSetV1ResponseValidationGroupsValidations `json:"validations,omitempty"` //
}
type ResponseHealthAndPerformanceRetrievesValidationDetailsForAValidationSetV1ResponseValidationGroupsValidations struct {
	ID   string `json:"id,omitempty"`   // Validation id
	Name string `json:"name,omitempty"` // Name of the validation
}
type ResponseHealthAndPerformanceRetrievesTheListOfValidationWorkflowsV1 struct {
	Response *[]ResponseHealthAndPerformanceRetrievesTheListOfValidationWorkflowsV1Response `json:"response,omitempty"` //
	Version  string                                                                         `json:"version,omitempty"`  // The version of the response
}
type ResponseHealthAndPerformanceRetrievesTheListOfValidationWorkflowsV1Response struct {
	ID               string   `json:"id,omitempty"`               // Workflow id
	Name             string   `json:"name,omitempty"`             // Workflow name
	Description      string   `json:"description,omitempty"`      // Workflow description
	RunStatus        string   `json:"runStatus,omitempty"`        // Execution status of the workflow. If the workflow is successfully submitted, runStatus will return `PENDING`. If the workflow execution has started, runStatus will return `IN_PROGRESS`. If the workflow executed is completed with all validations executed, runStatus will return `COMPLETED`. If the workflow execution fails while running validations, runStatus will return `FAILED`.
	SubmitTime       *int     `json:"submitTime,omitempty"`       // Workflow submit time (as milliseconds since UNIX epoch).
	StartTime        *int     `json:"startTime,omitempty"`        // Workflow start time (as milliseconds since UNIX epoch).
	EndTime          *int     `json:"endTime,omitempty"`          // Workflow finish time (as milliseconds since UNIX epoch).
	ValidationStatus string   `json:"validationStatus,omitempty"` // Overall result of execution of the validation workflow
	ValidationSetIDs []string `json:"validationSetIds,omitempty"` // List of validation set ids
}
type ResponseHealthAndPerformanceSubmitsTheWorkflowForExecutingValidationsV1 struct {
	Response *ResponseHealthAndPerformanceSubmitsTheWorkflowForExecutingValidationsV1Response `json:"response,omitempty"` //
	Version  string                                                                           `json:"version,omitempty"`  // The version of the response
}
type ResponseHealthAndPerformanceSubmitsTheWorkflowForExecutingValidationsV1Response struct {
	ID string `json:"id,omitempty"` // UUID of the workflow submitted for executing validations
}
type ResponseHealthAndPerformanceRetrievesTheCountOfValidationWorkflowsV1 struct {
	Response *ResponseHealthAndPerformanceRetrievesTheCountOfValidationWorkflowsV1Response `json:"response,omitempty"` //
	Version  string                                                                        `json:"version,omitempty"`  // The version of the response
}
type ResponseHealthAndPerformanceRetrievesTheCountOfValidationWorkflowsV1Response struct {
	Count *int `json:"count,omitempty"` // The reported count.
}
type ResponseHealthAndPerformanceRetrievesValidationWorkflowDetailsV1 struct {
	Response *ResponseHealthAndPerformanceRetrievesValidationWorkflowDetailsV1Response `json:"response,omitempty"` //
	Version  string                                                                    `json:"version,omitempty"`  // The version of the response
}
type ResponseHealthAndPerformanceRetrievesValidationWorkflowDetailsV1Response struct {
	ID                       string                                                                                              `json:"id,omitempty"`                       // Workflow id
	Name                     string                                                                                              `json:"name,omitempty"`                     // Workflow name
	Description              string                                                                                              `json:"description,omitempty"`              // Workflow description
	RunStatus                string                                                                                              `json:"runStatus,omitempty"`                // Execution status of the workflow. If the workflow is successfully submitted, runStatus will return `PENDING`. If the workflow execution has started, runStatus will return `IN_PROGRESS`. If the workflow executed is completed with all validations executed, runStatus will return `COMPLETED`. If the workflow execution fails while running validations, runStatus will return `FAILED`.
	SubmitTime               *int                                                                                                `json:"submitTime,omitempty"`               // Workflow submit time (as milliseconds since UNIX epoch).
	ValidationSetIDs         []string                                                                                            `json:"validationSetIds,omitempty"`         // List of validation set ids
	ReleaseVersion           string                                                                                              `json:"releaseVersion,omitempty"`           // Product version
	ValidationSetsRunDetails *[]ResponseHealthAndPerformanceRetrievesValidationWorkflowDetailsV1ResponseValidationSetsRunDetails `json:"validationSetsRunDetails,omitempty"` //
	ValidationStatus         string                                                                                              `json:"validationStatus,omitempty"`         // Overall result of the execution of all the validations. If any of the contained validation execution status is `CRITICAL`, this is marked as `CRITICAL`. Else, if any of the contained validation execution status is `WARNING`, this is marked as `WARNING`. Else, this is marked as `INFORMATION`.
}
type ResponseHealthAndPerformanceRetrievesValidationWorkflowDetailsV1ResponseValidationSetsRunDetails struct {
	ValidationSetID      string                                                                                                                  `json:"validationSetId,omitempty"`      // Validation set id
	StartTime            *int                                                                                                                    `json:"startTime,omitempty"`            // Validation set run start time (as milliseconds since UNIX epoch).
	EndTime              *int                                                                                                                    `json:"endTime,omitempty"`              // Validation set run finish time (as milliseconds since UNIX epoch).
	ValidationStatus     string                                                                                                                  `json:"validationStatus,omitempty"`     // Overall result of the validation set execution. If any of the contained validation execution status is `CRITICAL`, this is marked as `CRITICAL`. Else, if any of the contained validation execution status is `WARNING`, this is marked as `WARNING`. Else, this is marked as `INFORMATION`. This is empty when the workflow is in progress.
	Version              string                                                                                                                  `json:"version,omitempty"`              // Validation set version
	ValidationRunDetails *[]ResponseHealthAndPerformanceRetrievesValidationWorkflowDetailsV1ResponseValidationSetsRunDetailsValidationRunDetails `json:"validationRunDetails,omitempty"` //
}
type ResponseHealthAndPerformanceRetrievesValidationWorkflowDetailsV1ResponseValidationSetsRunDetailsValidationRunDetails struct {
	ValidationID      string `json:"validationId,omitempty"`      // Validation id
	ValidationName    string `json:"validationName,omitempty"`    // Validation name
	ValidationMessage string `json:"validationMessage,omitempty"` // Validation execution result detail message
	ValidationStatus  string `json:"validationStatus,omitempty"`  // Validation execution result status
}
type ResponseHealthAndPerformanceSystemHealthAPIV1 struct {
	HealthEvents *[]ResponseHealthAndPerformanceSystemHealthAPIV1HealthEvents `json:"healthEvents,omitempty"` //
	Version      string                                                       `json:"version,omitempty"`      // API version
	HostName     string                                                       `json:"hostName,omitempty"`     // Cluster name
	Cimcaddress  []string                                                     `json:"cimcaddress,omitempty"`  // List of configured cimc addresse(s)
}
type ResponseHealthAndPerformanceSystemHealthAPIV1HealthEvents struct {
	Severity    string `json:"severity,omitempty"`    // Severity of the event
	Hostname    string `json:"hostname,omitempty"`    // Hostname of the event
	Instance    string `json:"instance,omitempty"`    // Instance of the event
	SubDomain   string `json:"subDomain,omitempty"`   // Sub domain of the event
	Domain      string `json:"domain,omitempty"`      // Domain of the event
	Description string `json:"description,omitempty"` // Details of the event
	State       string `json:"state,omitempty"`       // State of the event
	Timestamp   string `json:"timestamp,omitempty"`   // Time of the event occurance
	Status      string `json:"status,omitempty"`      // Event status
}
type ResponseHealthAndPerformanceSystemHealthCountAPIV1 struct {
	Count *float64 `json:"count,omitempty"` // Count of the events
}
type ResponseHealthAndPerformanceSystemPerformanceAPIV1 struct {
	HostName string                                                  `json:"hostName,omitempty"` // Hostname
	Version  string                                                  `json:"version,omitempty"`  // Version of the API
	Kpis     *ResponseHealthAndPerformanceSystemPerformanceAPIV1Kpis `json:"kpis,omitempty"`     //
}
type ResponseHealthAndPerformanceSystemPerformanceAPIV1Kpis struct {
	CPU           *ResponseHealthAndPerformanceSystemPerformanceAPIV1KpisCPU           `json:"cpu,omitempty"`             //
	Memory        *ResponseHealthAndPerformanceSystemPerformanceAPIV1KpisMemory        `json:"memory,omitempty"`          //
	NetworktxRate *ResponseHealthAndPerformanceSystemPerformanceAPIV1KpisNetworktxRate `json:"network tx_rate,omitempty"` //
	NetworkrxRate *ResponseHealthAndPerformanceSystemPerformanceAPIV1KpisNetworkrxRate `json:"network rx_rate,omitempty"` //
}
type ResponseHealthAndPerformanceSystemPerformanceAPIV1KpisCPU struct {
	Units       string `json:"units,omitempty"`       // Units for cpu usage
	Utilization string `json:"utilization,omitempty"` // cpu usage in units
}
type ResponseHealthAndPerformanceSystemPerformanceAPIV1KpisMemory struct {
	Units       string `json:"units,omitempty"`       // Units for memory usage
	Utilization string `json:"utilization,omitempty"` // Memory usage in units
}
type ResponseHealthAndPerformanceSystemPerformanceAPIV1KpisNetworktxRate struct {
	Units       string `json:"units,omitempty"`       // Units for network tx_rate
	Utilization string `json:"utilization,omitempty"` // Network tx_rate in units
}
type ResponseHealthAndPerformanceSystemPerformanceAPIV1KpisNetworkrxRate struct {
	Units       string `json:"units,omitempty"`       // Units for network rx_rate
	Utilization string `json:"utilization,omitempty"` // Network rx_rate in units
}
type ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIV1 struct {
	HostName string                                                            `json:"hostName,omitempty"` // Hostname
	Version  string                                                            `json:"version,omitempty"`  // Version of the API
	Kpis     *ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIV1Kpis `json:"kpis,omitempty"`     //
}
type ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIV1Kpis struct {
	Legends   *ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIV1KpisLegends `json:"legends,omitempty"`   //
	Data      *ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIV1KpisData    `json:"data,omitempty"`      //
	CPUAvg    string                                                                   `json:"cpuAvg,omitempty"`    // CPU average utilization
	MemoryAvg string                                                                   `json:"memoryAvg,omitempty"` // Memory average utilization
}
type ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIV1KpisLegends struct {
	CPU           *ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIV1KpisLegendsCPU           `json:"cpu,omitempty"`             //
	Memory        *ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIV1KpisLegendsMemory        `json:"memory,omitempty"`          //
	NetworktxRate *ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIV1KpisLegendsNetworktxRate `json:"network tx_rate,omitempty"` //
	NetworkrxRate *ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIV1KpisLegendsNetworkrxRate `json:"network rx_rate,omitempty"` //
}
type ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIV1KpisLegendsCPU struct {
	Units string `json:"units,omitempty"` // Units for cpu usage
}
type ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIV1KpisLegendsMemory struct {
	Units string `json:"units,omitempty"` // Units for memory usage
}
type ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIV1KpisLegendsNetworktxRate struct {
	Units string `json:"units,omitempty"` // Units for network tx_rate
}
type ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIV1KpisLegendsNetworkrxRate struct {
	Units string `json:"units,omitempty"` // Units for network rx_rate
}
type ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIV1KpisData map[string][]string // Time in 'YYYY-MM-DDT00:00:00Z' format with values for legends
type RequestHealthAndPerformanceSubmitsTheWorkflowForExecutingValidationsV1 struct {
	Name             string   `json:"name,omitempty"`             // Name of the workflow to run. It must be unique.
	Description      string   `json:"description,omitempty"`      // Description of the workflow to run
	ValidationSetIDs []string `json:"validationSetIds,omitempty"` // List of validation set ids
}

//RetrievesAllTheValidationSetsV1 Retrieves all the validation sets - 11bb-4b03-4059-a001
/* Retrieves all the validation sets and optionally the contained validations


@param RetrievesAllTheValidationSetsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-all-the-validation-sets-v1
*/
func (s *HealthAndPerformanceService) RetrievesAllTheValidationSetsV1(RetrievesAllTheValidationSetsV1QueryParams *RetrievesAllTheValidationSetsV1QueryParams) (*ResponseHealthAndPerformanceRetrievesAllTheValidationSetsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/diagnosticValidationSets"

	queryString, _ := query.Values(RetrievesAllTheValidationSetsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseHealthAndPerformanceRetrievesAllTheValidationSetsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesAllTheValidationSetsV1(RetrievesAllTheValidationSetsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesAllTheValidationSetsV1")
	}

	result := response.Result().(*ResponseHealthAndPerformanceRetrievesAllTheValidationSetsV1)
	return result, response, err

}

//RetrievesValidationDetailsForAValidationSetV1 Retrieves validation details for a validation set - 37b7-88bd-47b9-8533
/* Retrieves validation details for the given validation set id


@param id id path parameter. Validation set id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-validation-details-for-a-validation-set-v1
*/
func (s *HealthAndPerformanceService) RetrievesValidationDetailsForAValidationSetV1(id string) (*ResponseHealthAndPerformanceRetrievesValidationDetailsForAValidationSetV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/diagnosticValidationSets/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseHealthAndPerformanceRetrievesValidationDetailsForAValidationSetV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesValidationDetailsForAValidationSetV1(id)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesValidationDetailsForAValidationSetV1")
	}

	result := response.Result().(*ResponseHealthAndPerformanceRetrievesValidationDetailsForAValidationSetV1)
	return result, response, err

}

//RetrievesTheListOfValidationWorkflowsV1 Retrieves the list of validation workflows - 0fab-cafd-440b-98f8
/* Retrieves the workflows that have been successfully submitted and are currently available. This is sorted by `submitTime`


@param RetrievesTheListOfValidationWorkflowsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-list-of-validation-workflows-v1
*/
func (s *HealthAndPerformanceService) RetrievesTheListOfValidationWorkflowsV1(RetrievesTheListOfValidationWorkflowsV1QueryParams *RetrievesTheListOfValidationWorkflowsV1QueryParams) (*ResponseHealthAndPerformanceRetrievesTheListOfValidationWorkflowsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/diagnosticValidationWorkflows"

	queryString, _ := query.Values(RetrievesTheListOfValidationWorkflowsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseHealthAndPerformanceRetrievesTheListOfValidationWorkflowsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheListOfValidationWorkflowsV1(RetrievesTheListOfValidationWorkflowsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheListOfValidationWorkflowsV1")
	}

	result := response.Result().(*ResponseHealthAndPerformanceRetrievesTheListOfValidationWorkflowsV1)
	return result, response, err

}

//RetrievesTheCountOfValidationWorkflowsV1 Retrieves the count of validation workflows - 4f91-8ac9-44c9-baef
/* Retrieves the count of workflows that have been successfully submitted and are currently available.


@param RetrievesTheCountOfValidationWorkflowsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-count-of-validation-workflows-v1
*/
func (s *HealthAndPerformanceService) RetrievesTheCountOfValidationWorkflowsV1(RetrievesTheCountOfValidationWorkflowsV1QueryParams *RetrievesTheCountOfValidationWorkflowsV1QueryParams) (*ResponseHealthAndPerformanceRetrievesTheCountOfValidationWorkflowsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/diagnosticValidationWorkflows/count"

	queryString, _ := query.Values(RetrievesTheCountOfValidationWorkflowsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseHealthAndPerformanceRetrievesTheCountOfValidationWorkflowsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheCountOfValidationWorkflowsV1(RetrievesTheCountOfValidationWorkflowsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheCountOfValidationWorkflowsV1")
	}

	result := response.Result().(*ResponseHealthAndPerformanceRetrievesTheCountOfValidationWorkflowsV1)
	return result, response, err

}

//RetrievesValidationWorkflowDetailsV1 Retrieves validation workflow details - eb8b-eaad-451a-9c09
/* Retrieves workflow details for a workflow id


@param id id path parameter. Workflow id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-validation-workflow-details-v1
*/
func (s *HealthAndPerformanceService) RetrievesValidationWorkflowDetailsV1(id string) (*ResponseHealthAndPerformanceRetrievesValidationWorkflowDetailsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/diagnosticValidationWorkflows/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseHealthAndPerformanceRetrievesValidationWorkflowDetailsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesValidationWorkflowDetailsV1(id)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesValidationWorkflowDetailsV1")
	}

	result := response.Result().(*ResponseHealthAndPerformanceRetrievesValidationWorkflowDetailsV1)
	return result, response, err

}

//SystemHealthAPIV1 System Health API - 6085-eb1b-4f48-9740
/* This API retrieves the latest system events


@param SystemHealthAPIV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!system-health-api-v1
*/
func (s *HealthAndPerformanceService) SystemHealthAPIV1(SystemHealthAPIV1QueryParams *SystemHealthAPIV1QueryParams) (*ResponseHealthAndPerformanceSystemHealthAPIV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/diagnostics/system/health"

	queryString, _ := query.Values(SystemHealthAPIV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseHealthAndPerformanceSystemHealthAPIV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SystemHealthAPIV1(SystemHealthAPIV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation SystemHealthApiV1")
	}

	result := response.Result().(*ResponseHealthAndPerformanceSystemHealthAPIV1)
	return result, response, err

}

//SystemHealthCountAPIV1 System Health Count API - 5289-0891-4729-8714
/* This API gives the count of the latest system events


@param SystemHealthCountAPIV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!system-health-count-api-v1
*/
func (s *HealthAndPerformanceService) SystemHealthCountAPIV1(SystemHealthCountAPIV1QueryParams *SystemHealthCountAPIV1QueryParams) (*ResponseHealthAndPerformanceSystemHealthCountAPIV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/diagnostics/system/health/count"

	queryString, _ := query.Values(SystemHealthCountAPIV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseHealthAndPerformanceSystemHealthCountAPIV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SystemHealthCountAPIV1(SystemHealthCountAPIV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation SystemHealthCountApiV1")
	}

	result := response.Result().(*ResponseHealthAndPerformanceSystemHealthCountAPIV1)
	return result, response, err

}

//SystemPerformanceAPIV1 System Performance API - f2a9-5b4d-48eb-a4f8
/* Retrieves the aggregated metrics (total, average or maximum) of cluster key performance indicators (KPIs), such as CPU utilization, memory utilization or network rates recorded within a specified time period. The data will be available from the past 24 hours.


@param SystemPerformanceAPIV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!system-performance-api-v1
*/
func (s *HealthAndPerformanceService) SystemPerformanceAPIV1(SystemPerformanceAPIV1QueryParams *SystemPerformanceAPIV1QueryParams) (*ResponseHealthAndPerformanceSystemPerformanceAPIV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/diagnostics/system/performance"

	queryString, _ := query.Values(SystemPerformanceAPIV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseHealthAndPerformanceSystemPerformanceAPIV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SystemPerformanceAPIV1(SystemPerformanceAPIV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation SystemPerformanceApiV1")
	}

	result := response.Result().(*ResponseHealthAndPerformanceSystemPerformanceAPIV1)
	return result, response, err

}

//SystemPerformanceHistoricalAPIV1 System Performance Historical API - 879b-ea1e-4389-83d7
/* Retrieves the average values of cluster key performance indicators (KPIs), like CPU utilization, memory utilization or network rates grouped by time intervals within a specified time range. The data will be available from the past 24 hours.


@param SystemPerformanceHistoricalAPIV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!system-performance-historical-api-v1
*/
func (s *HealthAndPerformanceService) SystemPerformanceHistoricalAPIV1(SystemPerformanceHistoricalAPIV1QueryParams *SystemPerformanceHistoricalAPIV1QueryParams) (*ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/diagnostics/system/performance/history"

	queryString, _ := query.Values(SystemPerformanceHistoricalAPIV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SystemPerformanceHistoricalAPIV1(SystemPerformanceHistoricalAPIV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation SystemPerformanceHistoricalApiV1")
	}

	result := response.Result().(*ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIV1)
	return result, response, err

}

//SubmitsTheWorkflowForExecutingValidationsV1 Submits the workflow for executing validations - 52a0-7981-41ab-81d8
/* Submits the workflow for executing the validations for the given validation specifications



Documentation Link: https://developer.cisco.com/docs/dna-center/#!submits-the-workflow-for-executing-validations-v1
*/
func (s *HealthAndPerformanceService) SubmitsTheWorkflowForExecutingValidationsV1(requestHealthAndPerformanceSubmitsTheWorkflowForExecutingValidationsV1 *RequestHealthAndPerformanceSubmitsTheWorkflowForExecutingValidationsV1) (*ResponseHealthAndPerformanceSubmitsTheWorkflowForExecutingValidationsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/diagnosticValidationWorkflows"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestHealthAndPerformanceSubmitsTheWorkflowForExecutingValidationsV1).
		SetResult(&ResponseHealthAndPerformanceSubmitsTheWorkflowForExecutingValidationsV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.SubmitsTheWorkflowForExecutingValidationsV1(requestHealthAndPerformanceSubmitsTheWorkflowForExecutingValidationsV1)
		}

		return nil, response, fmt.Errorf("error with operation SubmitsTheWorkflowForExecutingValidationsV1")
	}

	result := response.Result().(*ResponseHealthAndPerformanceSubmitsTheWorkflowForExecutingValidationsV1)
	return result, response, err

}

//DeletesAValidationWorkflowV1 Deletes a validation workflow - 8eb3-3959-47fa-9c50
/* Deletes the workflow for the given id


@param id id path parameter. Workflow id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!deletes-a-validation-workflow-v1
*/
func (s *HealthAndPerformanceService) DeletesAValidationWorkflowV1(id string) (*resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/diagnosticValidationWorkflows/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletesAValidationWorkflowV1(
				id)
		}
		return response, fmt.Errorf("error with operation DeletesAValidationWorkflowV1")
	}

	return response, err

}

// Alias Function
func (s *HealthAndPerformanceService) SystemPerformanceAPI(SystemPerformanceAPIV1QueryParams *SystemPerformanceAPIV1QueryParams) (*ResponseHealthAndPerformanceSystemPerformanceAPIV1, *resty.Response, error) {
	return s.SystemPerformanceAPIV1(SystemPerformanceAPIV1QueryParams)
}

// Alias Function
func (s *HealthAndPerformanceService) SystemPerformanceHistoricalAPI(SystemPerformanceHistoricalAPIV1QueryParams *SystemPerformanceHistoricalAPIV1QueryParams) (*ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIV1, *resty.Response, error) {
	return s.SystemPerformanceHistoricalAPIV1(SystemPerformanceHistoricalAPIV1QueryParams)
}

// Alias Function
func (s *HealthAndPerformanceService) SystemHealthAPI(SystemHealthAPIV1QueryParams *SystemHealthAPIV1QueryParams) (*ResponseHealthAndPerformanceSystemHealthAPIV1, *resty.Response, error) {
	return s.SystemHealthAPIV1(SystemHealthAPIV1QueryParams)
}

// Alias Function
func (s *HealthAndPerformanceService) DeletesAValidationWorkflow(id string) (*resty.Response, error) {
	return s.DeletesAValidationWorkflowV1(id)
}

// Alias Function
func (s *HealthAndPerformanceService) RetrievesValidationDetailsForAValidationSet(id string) (*ResponseHealthAndPerformanceRetrievesValidationDetailsForAValidationSetV1, *resty.Response, error) {
	return s.RetrievesValidationDetailsForAValidationSetV1(id)
}

// Alias Function
func (s *HealthAndPerformanceService) RetrievesTheCountOfValidationWorkflows(RetrievesTheCountOfValidationWorkflowsV1QueryParams *RetrievesTheCountOfValidationWorkflowsV1QueryParams) (*ResponseHealthAndPerformanceRetrievesTheCountOfValidationWorkflowsV1, *resty.Response, error) {
	return s.RetrievesTheCountOfValidationWorkflowsV1(RetrievesTheCountOfValidationWorkflowsV1QueryParams)
}

// Alias Function
func (s *HealthAndPerformanceService) RetrievesTheListOfValidationWorkflows(RetrievesTheListOfValidationWorkflowsV1QueryParams *RetrievesTheListOfValidationWorkflowsV1QueryParams) (*ResponseHealthAndPerformanceRetrievesTheListOfValidationWorkflowsV1, *resty.Response, error) {
	return s.RetrievesTheListOfValidationWorkflowsV1(RetrievesTheListOfValidationWorkflowsV1QueryParams)
}

// Alias Function
func (s *HealthAndPerformanceService) SystemHealthCountAPI(SystemHealthCountAPIV1QueryParams *SystemHealthCountAPIV1QueryParams) (*ResponseHealthAndPerformanceSystemHealthCountAPIV1, *resty.Response, error) {
	return s.SystemHealthCountAPIV1(SystemHealthCountAPIV1QueryParams)
}

// Alias Function
func (s *HealthAndPerformanceService) RetrievesValidationWorkflowDetails(id string) (*ResponseHealthAndPerformanceRetrievesValidationWorkflowDetailsV1, *resty.Response, error) {
	return s.RetrievesValidationWorkflowDetailsV1(id)
}

// Alias Function
func (s *HealthAndPerformanceService) SubmitsTheWorkflowForExecutingValidations(requestHealthAndPerformanceSubmitsTheWorkflowForExecutingValidationsV1 *RequestHealthAndPerformanceSubmitsTheWorkflowForExecutingValidationsV1) (*ResponseHealthAndPerformanceSubmitsTheWorkflowForExecutingValidationsV1, *resty.Response, error) {
	return s.SubmitsTheWorkflowForExecutingValidationsV1(requestHealthAndPerformanceSubmitsTheWorkflowForExecutingValidationsV1)
}

// Alias Function
func (s *HealthAndPerformanceService) RetrievesAllTheValidationSets(RetrievesAllTheValidationSetsV1QueryParams *RetrievesAllTheValidationSetsV1QueryParams) (*ResponseHealthAndPerformanceRetrievesAllTheValidationSetsV1, *resty.Response, error) {
	return s.RetrievesAllTheValidationSetsV1(RetrievesAllTheValidationSetsV1QueryParams)
}
