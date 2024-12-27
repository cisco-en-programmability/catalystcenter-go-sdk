package catalyst

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type TaskService service

type RetrieveAListOfAssuranceTasksV1QueryParams struct {
	Limit  float64 `url:"limit,omitempty"`  //Maximum number of records to return
	Offset float64 `url:"offset,omitempty"` //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	SortBy string  `url:"sortBy,omitempty"` //A field within the response to sort by.
	Order  string  `url:"order,omitempty"`  //The sort order of the field ascending or descending.
	Status string  `url:"status,omitempty"` //used to get a subset of tasks by their status
}
type RetrieveAListOfAssuranceTasksV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistV1QueryParams struct {
	Status string `url:"status,omitempty"` //used to get a subset of tasks by their status
}
type RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrieveASpecificAssuranceTaskByIDV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetTasksOperationalTasksV1QueryParams struct {
	StartTime     string  `url:"startTime,omitempty"`     //This is the epoch start time from which tasks need to be fetched
	EndTime       string  `url:"endTime,omitempty"`       //This is the epoch end time upto which audit records need to be fetched
	Data          string  `url:"data,omitempty"`          //Fetch tasks that contains this data
	ErrorCode     string  `url:"errorCode,omitempty"`     //Fetch tasks that have this error code
	ServiceType   string  `url:"serviceType,omitempty"`   //Fetch tasks with this service type
	Username      string  `url:"username,omitempty"`      //Fetch tasks with this username
	Progress      string  `url:"progress,omitempty"`      //Fetch tasks that contains this progress
	IsError       string  `url:"isError,omitempty"`       //Fetch tasks ended as success or failure. Valid values: true, false
	FailureReason string  `url:"failureReason,omitempty"` //Fetch tasks that contains this failure reason
	ParentID      string  `url:"parentId,omitempty"`      //Fetch tasks that have this parent Id
	Offset        float64 `url:"offset,omitempty"`        //The first record to show for this page; the first record is numbered 1.
	Limit         float64 `url:"limit,omitempty"`         //The number of records to show for this page;The minimum is 1, and the maximum is 500.
	SortBy        string  `url:"sortBy,omitempty"`        //Sort results by this field
	Order         string  `url:"order,omitempty"`         //Sort order - asc or dsc
}
type GetTaskCountV1QueryParams struct {
	StartTime     string `url:"startTime,omitempty"`     //This is the epoch start time from which tasks need to be fetched
	EndTime       string `url:"endTime,omitempty"`       //This is the epoch end time upto which audit records need to be fetched
	Data          string `url:"data,omitempty"`          //Fetch tasks that contains this data
	ErrorCode     string `url:"errorCode,omitempty"`     //Fetch tasks that have this error code
	ServiceType   string `url:"serviceType,omitempty"`   //Fetch tasks with this service type
	Username      string `url:"username,omitempty"`      //Fetch tasks with this username
	Progress      string `url:"progress,omitempty"`      //Fetch tasks that contains this progress
	IsError       string `url:"isError,omitempty"`       //Fetch tasks ended as success or failure. Valid values: true, false
	FailureReason string `url:"failureReason,omitempty"` //Fetch tasks that contains this failure reason
	ParentID      string `url:"parentId,omitempty"`      //Fetch tasks that have this parent Id
}
type GetTasksV1QueryParams struct {
	Offset    float64 `url:"offset,omitempty"`    //The first record to show for this page; the first record is numbered 1.
	Limit     float64 `url:"limit,omitempty"`     //The number of records to show for this page;The minimum is 1, and the maximum is 500.
	SortBy    string  `url:"sortBy,omitempty"`    //A property within the response to sort by.
	Order     string  `url:"order,omitempty"`     //Whether ascending or descending order should be used to sort the response.
	StartTime int     `url:"startTime,omitempty"` //This is the epoch millisecond start time from which tasks need to be fetched
	EndTime   int     `url:"endTime,omitempty"`   //This is the epoch millisecond end time upto which task records need to be fetched
	ParentID  string  `url:"parentId,omitempty"`  //Fetch tasks that have this parent Id
	RootID    string  `url:"rootId,omitempty"`    //Fetch tasks that have this root Id
	Status    string  `url:"status,omitempty"`    //Fetch tasks that have this status. Available values : PENDING, FAILURE, SUCCESS
}
type GetTasksCountV1QueryParams struct {
	StartTime int    `url:"startTime,omitempty"` //This is the epoch millisecond start time from which tasks need to be fetched
	EndTime   int    `url:"endTime,omitempty"`   //This is the epoch millisecond end time upto which task records need to be fetched
	ParentID  string `url:"parentId,omitempty"`  //Fetch tasks that have this parent Id
	RootID    string `url:"rootId,omitempty"`    //Fetch tasks that have this root Id
	Status    string `url:"status,omitempty"`    //Fetch tasks that have this status. Available values : PENDING, FAILURE, SUCCESS
}

type ResponseTaskRetrieveAListOfAssuranceTasksV1 struct {
	Response *[]ResponseTaskRetrieveAListOfAssuranceTasksV1Response `json:"response,omitempty"` //

	Page *ResponseTaskRetrieveAListOfAssuranceTasksV1Page `json:"page,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseTaskRetrieveAListOfAssuranceTasksV1Response struct {
	ID string `json:"id,omitempty"` // Id

	Status string `json:"status,omitempty"` // Status

	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	UpdateTime *int `json:"updateTime,omitempty"` // Update Time

	Progress string `json:"progress,omitempty"` // Progress

	FailureReason string `json:"failureReason,omitempty"` // Failure Reason

	ErrorCode string `json:"errorCode,omitempty"` // Error Code

	RequestType string `json:"requestType,omitempty"` // Request Type

	Data *ResponseTaskRetrieveAListOfAssuranceTasksV1ResponseData `json:"data,omitempty"` // Data

	ResultURL string `json:"resultUrl,omitempty"` // Result Url
}
type ResponseTaskRetrieveAListOfAssuranceTasksV1ResponseData interface{}
type ResponseTaskRetrieveAListOfAssuranceTasksV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	Count *int `json:"count,omitempty"` // Count

	SortBy *[]ResponseTaskRetrieveAListOfAssuranceTasksV1PageSortBy `json:"sortBy,omitempty"` //
}
type ResponseTaskRetrieveAListOfAssuranceTasksV1PageSortBy struct {
	Name string `json:"name,omitempty"` // Name

	Order string `json:"order,omitempty"` // Order
}
type ResponseTaskRetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistV1 struct {
	Response *ResponseTaskRetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseTaskRetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseTaskRetrieveASpecificAssuranceTaskByIDV1 struct {
	Response *ResponseTaskRetrieveASpecificAssuranceTaskByIDV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseTaskRetrieveASpecificAssuranceTaskByIDV1Response struct {
	ID string `json:"id,omitempty"` // Id

	Status string `json:"status,omitempty"` // Status

	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	UpdateTime *int `json:"updateTime,omitempty"` // Update Time

	Progress string `json:"progress,omitempty"` // Progress

	FailureReason string `json:"failureReason,omitempty"` // Failure Reason

	ErrorCode string `json:"errorCode,omitempty"` // Error Code

	RequestType string `json:"requestType,omitempty"` // Request Type

	Data *ResponseTaskRetrieveASpecificAssuranceTaskByIDV1ResponseData `json:"data,omitempty"` // Data

	ResultURL string `json:"resultUrl,omitempty"` // Result Url
}
type ResponseTaskRetrieveASpecificAssuranceTaskByIDV1ResponseData interface{}
type ResponseTaskGetBusinessAPIExecutionDetailsV1 struct {
	BapiKey string `json:"bapiKey,omitempty"` // Business API Key (UUID)

	BapiName string `json:"bapiName,omitempty"` // Name of the Business API

	BapiExecutionID string `json:"bapiExecutionId,omitempty"` // Execution Id of the Business API (UUID)

	StartTime string `json:"startTime,omitempty"` // Execution Start Time of the Business API (Date Time Format)

	StartTimeEpoch *int `json:"startTimeEpoch,omitempty"` // Execution Start Time of the Business API (Epoch Milliseconds)

	EndTime string `json:"endTime,omitempty"` // Execution End Time of the Business API (Date Time Format)

	EndTimeEpoch *int `json:"endTimeEpoch,omitempty"` // Execution End Time of the Business API (Epoch Milliseconds)

	TimeDuration *int `json:"timeDuration,omitempty"` // Time taken for Business API Execution (Milliseconds)

	Status string `json:"status,omitempty"` // Execution status of the Business API

	BapiSyncResponse string `json:"bapiSyncResponse,omitempty"` // Returns the actual response of the original API as a string

	BapiSyncResponseJSON *ResponseTaskGetBusinessAPIExecutionDetailsV1BapiSyncResponseJSON `json:"bapiSyncResponseJson,omitempty"` // Returns the actual response of the original API  as a json

	RuntimeInstanceID string `json:"runtimeInstanceId,omitempty"` // Pod Id in which the Business API is executed

	BapiError string `json:"bapiError,omitempty"` // Returns the error response of the original API  as a string
}
type ResponseTaskGetBusinessAPIExecutionDetailsV1BapiSyncResponseJSON interface{}
type ResponseTaskGetTasksOperationalTasksV1 struct {
	Response *[]ResponseTaskGetTasksOperationalTasksV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseTaskGetTasksOperationalTasksV1Response struct {
	AdditionalStatusURL string `json:"additionalStatusURL,omitempty"` //

	Data string `json:"data,omitempty"` //

	EndTime *int `json:"endTime,omitempty"` //

	ErrorCode string `json:"errorCode,omitempty"` //

	ErrorKey string `json:"errorKey,omitempty"` //

	FailureReason string `json:"failureReason,omitempty"` //

	ID string `json:"id,omitempty"` //

	InstanceTenantID string `json:"instanceTenantId,omitempty"` //

	IsError *bool `json:"isError,omitempty"` //

	LastUpdate *int `json:"lastUpdate,omitempty"` //

	OperationIDList []string `json:"operationIdList,omitempty"` //

	ParentID string `json:"parentId,omitempty"` //

	Progress string `json:"progress,omitempty"` //

	RootID string `json:"rootId,omitempty"` //

	ServiceType string `json:"serviceType,omitempty"` //

	StartTime *int `json:"startTime,omitempty"` //

	Username string `json:"username,omitempty"` //

	Version *int `json:"version,omitempty"` //
}
type ResponseTaskGetTaskCountV1 struct {
	Response *int `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseTaskGetTaskByOperationIDV1 struct {
	Response *[]ResponseTaskGetTaskByOperationIDV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseTaskGetTaskByOperationIDV1Response struct {
	AdditionalStatusURL string `json:"additionalStatusURL,omitempty"` //

	Data string `json:"data,omitempty"` //

	EndTime *int `json:"endTime,omitempty"` //

	ErrorCode string `json:"errorCode,omitempty"` //

	ErrorKey string `json:"errorKey,omitempty"` //

	FailureReason string `json:"failureReason,omitempty"` //

	ID string `json:"id,omitempty"` //

	InstanceTenantID string `json:"instanceTenantId,omitempty"` //

	IsError *bool `json:"isError,omitempty"` //

	LastUpdate *int `json:"lastUpdate,omitempty"` //

	OperationIDList []string `json:"operationIdList,omitempty"` //

	ParentID string `json:"parentId,omitempty"` //

	Progress string `json:"progress,omitempty"` //

	RootID string `json:"rootId,omitempty"` //

	ServiceType string `json:"serviceType,omitempty"` //

	StartTime *int `json:"startTime,omitempty"` //

	Username string `json:"username,omitempty"` //

	Version *int `json:"version,omitempty"` //
}
type ResponseTaskGetTaskByIDV1 struct {
	Response *ResponseTaskGetTaskByIDV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseTaskGetTaskByIDV1Response struct {
	AdditionalStatusURL string `json:"additionalStatusURL,omitempty"` //

	Data string `json:"data,omitempty"` //

	EndTime *int `json:"endTime,omitempty"` //

	ErrorCode string `json:"errorCode,omitempty"` //

	ErrorKey string `json:"errorKey,omitempty"` //

	FailureReason string `json:"failureReason,omitempty"` //

	ID string `json:"id,omitempty"` //

	InstanceTenantID string `json:"instanceTenantId,omitempty"` //

	IsError *bool `json:"isError,omitempty"` //

	LastUpdate *int `json:"lastUpdate,omitempty"` //

	OperationIDList []string `json:"operationIdList,omitempty"` //

	ParentID string `json:"parentId,omitempty"` //

	Progress string `json:"progress,omitempty"` //

	RootID string `json:"rootId,omitempty"` //

	ServiceType string `json:"serviceType,omitempty"` //

	StartTime *int `json:"startTime,omitempty"` //

	Username string `json:"username,omitempty"` //

	Version *int `json:"version,omitempty"` //
}
type ResponseTaskGetTaskTreeV1 struct {
	Response *[]ResponseTaskGetTaskTreeV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseTaskGetTaskTreeV1Response struct {
	AdditionalStatusURL string `json:"additionalStatusURL,omitempty"` //

	Data string `json:"data,omitempty"` //

	EndTime *int `json:"endTime,omitempty"` //

	ErrorCode string `json:"errorCode,omitempty"` //

	ErrorKey string `json:"errorKey,omitempty"` //

	FailureReason string `json:"failureReason,omitempty"` //

	ID string `json:"id,omitempty"` //

	InstanceTenantID string `json:"instanceTenantId,omitempty"` //

	IsError *bool `json:"isError,omitempty"` //

	LastUpdate *int `json:"lastUpdate,omitempty"` //

	OperationIDList []string `json:"operationIdList,omitempty"` //

	ParentID string `json:"parentId,omitempty"` //

	Progress string `json:"progress,omitempty"` //

	RootID string `json:"rootId,omitempty"` //

	ServiceType string `json:"serviceType,omitempty"` //

	StartTime *int `json:"startTime,omitempty"` //

	Username string `json:"username,omitempty"` //

	Version *int `json:"version,omitempty"` //
}
type ResponseTaskGetTasksV1 struct {
	Response *[]ResponseTaskGetTasksV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseTaskGetTasksV1Response struct {
	EndTime *int `json:"endTime,omitempty"` // An approximate time of when this task has been marked completed; as measured in Unix epoch time in milliseconds

	ID string `json:"id,omitempty"` // The ID of this task

	UpdatedTime *int `json:"updatedTime,omitempty"` // A timestamp of when this task was last updated; as measured in Unix epoch time in milliseconds

	ParentID string `json:"parentId,omitempty"` // The ID of the parent task if this happens to be a subtask. In case this task is not a subtask, then the parentId is expected to be null.  To construct a task tree, this task will be the child of the task with the ID listed here, or the root of the tree if this task has no parentId.

	ResultLocation string `json:"resultLocation,omitempty"` // A server-relative URL indicating where additional task-specific details may be found

	RootID string `json:"rootId,omitempty"` // The ID of the task representing the root node of the tree which this task belongs to.  In some cases, this may be the same as the ID or null, which indicates that this task is the root task.

	StartTime *int `json:"startTime,omitempty"` // An approximate time of when the task creation was triggered; as measured in Unix epoch time in milliseconds

	Status string `json:"status,omitempty"` //
}
type ResponseTaskGetTasksCountV1 struct {
	Response *ResponseTaskGetTasksCountV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // The version of the response
}
type ResponseTaskGetTasksCountV1Response struct {
	Count *int `json:"count,omitempty"` // The reported count.
}
type ResponseTaskGetTasksByIDV1 struct {
	Response *ResponseTaskGetTasksByIDV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseTaskGetTasksByIDV1Response struct {
	EndTime *int `json:"endTime,omitempty"` // An approximate time of when this task has been marked completed; as measured in Unix epoch time in milliseconds

	ID string `json:"id,omitempty"` // The ID of this task

	UpdatedTime *int `json:"updatedTime,omitempty"` // A timestamp of when this task was last updated; as measured in Unix epoch time in milliseconds

	ParentID string `json:"parentId,omitempty"` // The ID of the parent task if this happens to be a subtask. In case this task is not a subtask, then the parentId is expected to be null.  To construct a task tree, this task will be the child of the task with the ID listed here, or the root of the tree if this task has no parentId.

	ResultLocation string `json:"resultLocation,omitempty"` // A server-relative URL indicating where additional task-specific details may be found

	RootID string `json:"rootId,omitempty"` // The ID of the task representing the root node of the tree which this task belongs to.  In some cases, this may be the same as the ID or null, which indicates that this task is the root task.

	StartTime *int `json:"startTime,omitempty"` // An approximate time of when the task creation was triggered; as measured in Unix epoch time in milliseconds

	Status string `json:"status,omitempty"` //
}
type ResponseTaskGetTaskDetailsByIDV1 struct {
	Response *ResponseTaskGetTaskDetailsByIDV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseTaskGetTaskDetailsByIDV1Response struct {
	Data string `json:"data,omitempty"` // Any data associated with this task; the value may vary significantly across different tasks

	Progress string `json:"progress,omitempty"` // A textual representation which indicates the progress of this task; the value may vary significantly across different tasks

	ErrorCode string `json:"errorCode,omitempty"` // An error code if in case this task has failed in its execution

	FailureReason string `json:"failureReason,omitempty"` // A textual description indicating the reason why a task has failed
}

//RetrieveAListOfAssuranceTasksV1 Retrieve a list of assurance tasks - ee8a-e874-40ca-8154
/* returns all existing tasks in a paginated list
default sorting of list is `startTime`, `asc`
valid field to sort by are [`startTime`,`endTime`,`updateTime`,`status`] For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceTasks-1.0.0-resolved.yaml


@param RetrieveAListOfAssuranceTasksV1HeaderParams Custom header parameters
@param RetrieveAListOfAssuranceTasksV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-a-list-of-assurance-tasks
*/
func (s *TaskService) RetrieveAListOfAssuranceTasksV1(RetrieveAListOfAssuranceTasksV1HeaderParams *RetrieveAListOfAssuranceTasksV1HeaderParams, RetrieveAListOfAssuranceTasksV1QueryParams *RetrieveAListOfAssuranceTasksV1QueryParams) (*ResponseTaskRetrieveAListOfAssuranceTasksV1, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceTasks"

	queryString, _ := query.Values(RetrieveAListOfAssuranceTasksV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrieveAListOfAssuranceTasksV1HeaderParams != nil {

		if RetrieveAListOfAssuranceTasksV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrieveAListOfAssuranceTasksV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseTaskRetrieveAListOfAssuranceTasksV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveAListOfAssuranceTasksV1(RetrieveAListOfAssuranceTasksV1HeaderParams, RetrieveAListOfAssuranceTasksV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveAListOfAssuranceTasksV1")
	}

	result := response.Result().(*ResponseTaskRetrieveAListOfAssuranceTasksV1)
	return result, response, err

}

//RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistV1 Retrieve a count of the number of assurance tasks that currently exist - b094-0b13-423b-bfb4
/* returns a count of the number of assurance tasks that are not expired For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceTasks-1.0.0-resolved.yaml


@param RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistV1HeaderParams Custom header parameters
@param RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-a-count-of-the-number-of-assurance-tasks-that-currently-exist
*/
func (s *TaskService) RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistV1(RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistV1HeaderParams *RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistV1HeaderParams, RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistV1QueryParams *RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistV1QueryParams) (*ResponseTaskRetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistV1, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceTasks/count"

	queryString, _ := query.Values(RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistV1HeaderParams != nil {

		if RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseTaskRetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistV1(RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistV1HeaderParams, RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistV1")
	}

	result := response.Result().(*ResponseTaskRetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistV1)
	return result, response, err

}

//RetrieveASpecificAssuranceTaskByIDV1 Retrieve a specific assurance task by id - 1e8f-7ae5-4798-88f1
/* returns a task given a specific task id For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceTasks-1.0.0-resolved.yaml


@param id id path parameter. unique task id

@param RetrieveASpecificAssuranceTaskByIdV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-a-specific-assurance-task-by-id
*/
func (s *TaskService) RetrieveASpecificAssuranceTaskByIDV1(id string, RetrieveASpecificAssuranceTaskByIdV1HeaderParams *RetrieveASpecificAssuranceTaskByIDV1HeaderParams) (*ResponseTaskRetrieveASpecificAssuranceTaskByIDV1, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceTasks/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrieveASpecificAssuranceTaskByIdV1HeaderParams != nil {

		if RetrieveASpecificAssuranceTaskByIdV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrieveASpecificAssuranceTaskByIdV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetResult(&ResponseTaskRetrieveASpecificAssuranceTaskByIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveASpecificAssuranceTaskByIDV1(id, RetrieveASpecificAssuranceTaskByIdV1HeaderParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveASpecificAssuranceTaskByIdV1")
	}

	result := response.Result().(*ResponseTaskRetrieveASpecificAssuranceTaskByIDV1)
	return result, response, err

}

//GetBusinessAPIExecutionDetailsV1 Get Business API Execution Details - c1bc-a8c1-41fb-9f75
/* Retrieves the execution details of a Business API


@param executionID executionId path parameter. Execution Id of API


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-business-api-execution-details
*/
func (s *TaskService) GetBusinessAPIExecutionDetailsV1(executionID string) (*ResponseTaskGetBusinessAPIExecutionDetailsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/dnacaap/management/execution-status/{executionId}"
	path = strings.Replace(path, "{executionId}", fmt.Sprintf("%v", executionID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTaskGetBusinessAPIExecutionDetailsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetBusinessAPIExecutionDetailsV1(executionID)
		}
		return nil, response, fmt.Errorf("error with operation GetBusinessApiExecutionDetailsV1")
	}

	result := response.Result().(*ResponseTaskGetBusinessAPIExecutionDetailsV1)
	return result, response, err

}

//GetTasksOperationalTasksV1 Get tasks - e78b-b8a2-449b-9eed
/* Returns task(s) based on filter criteria


@param GetTasksOperationalTasksV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-tasks-operational-tasks
*/
func (s *TaskService) GetTasksOperationalTasksV1(GetTasksOperationalTasksV1QueryParams *GetTasksOperationalTasksV1QueryParams) (*ResponseTaskGetTasksOperationalTasksV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/task"

	queryString, _ := query.Values(GetTasksOperationalTasksV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseTaskGetTasksOperationalTasksV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTasksOperationalTasksV1(GetTasksOperationalTasksV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTasksOperationalTasksV1")
	}

	result := response.Result().(*ResponseTaskGetTasksOperationalTasksV1)
	return result, response, err

}

//GetTaskCountV1 Get task count - 26b4-4ab0-4649-a183
/* Returns Task count


@param GetTaskCountV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-task-count
*/
func (s *TaskService) GetTaskCountV1(GetTaskCountV1QueryParams *GetTaskCountV1QueryParams) (*ResponseTaskGetTaskCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/task/count"

	queryString, _ := query.Values(GetTaskCountV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseTaskGetTaskCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTaskCountV1(GetTaskCountV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTaskCountV1")
	}

	result := response.Result().(*ResponseTaskGetTaskCountV1)
	return result, response, err

}

//GetTaskByOperationIDV1 Get task by OperationId - e487-f8d3-481b-94f2
/* Returns root tasks associated with an Operationid


@param operationID operationId path parameter.
@param offset offset path parameter. Index, minimum value is 0

@param limit limit path parameter. The maximum value of {limit} supported is 500. <br/> Base 1 indexing for {limit}, minimum value is 1


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-task-by-operation-id
*/
func (s *TaskService) GetTaskByOperationIDV1(operationID string, offset int, limit int) (*ResponseTaskGetTaskByOperationIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/task/operation/{operationId}/{offset}/{limit}"
	path = strings.Replace(path, "{operationId}", fmt.Sprintf("%v", operationID), -1)
	path = strings.Replace(path, "{offset}", fmt.Sprintf("%v", offset), -1)
	path = strings.Replace(path, "{limit}", fmt.Sprintf("%v", limit), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTaskGetTaskByOperationIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTaskByOperationIDV1(operationID, offset, limit)
		}
		return nil, response, fmt.Errorf("error with operation GetTaskByOperationIdV1")
	}

	result := response.Result().(*ResponseTaskGetTaskByOperationIDV1)
	return result, response, err

}

//GetTaskByIDV1 Get task by Id - a1a9-3873-46ba-92b1
/* Returns a task by specified id


@param taskID taskId path parameter. UUID of the Task


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-task-by-id
*/
func (s *TaskService) GetTaskByIDV1(taskID string) (*ResponseTaskGetTaskByIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/task/{taskId}"
	path = strings.Replace(path, "{taskId}", fmt.Sprintf("%v", taskID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTaskGetTaskByIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTaskByIDV1(taskID)
		}
		return nil, response, fmt.Errorf("error with operation GetTaskByIdV1")
	}

	result := response.Result().(*ResponseTaskGetTaskByIDV1)
	return result, response, err

}

//GetTaskTreeV1 Get task tree - f5a2-69c4-4f2a-95fa
/* Returns a task with its children tasks by based on their id


@param taskID taskId path parameter. UUID of the Task


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-task-tree
*/
func (s *TaskService) GetTaskTreeV1(taskID string) (*ResponseTaskGetTaskTreeV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/task/{taskId}/tree"
	path = strings.Replace(path, "{taskId}", fmt.Sprintf("%v", taskID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTaskGetTaskTreeV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTaskTreeV1(taskID)
		}
		return nil, response, fmt.Errorf("error with operation GetTaskTreeV1")
	}

	result := response.Result().(*ResponseTaskGetTaskTreeV1)
	return result, response, err

}

//GetTasksV1 Get tasks - b7bf-c860-466b-aaa7
/* Returns task(s) based on filter criteria


@param GetTasksV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-tasks
*/
func (s *TaskService) GetTasksV1(GetTasksV1QueryParams *GetTasksV1QueryParams) (*ResponseTaskGetTasksV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/tasks"

	queryString, _ := query.Values(GetTasksV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseTaskGetTasksV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTasksV1(GetTasksV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTasksV1")
	}

	result := response.Result().(*ResponseTaskGetTasksV1)
	return result, response, err

}

//GetTasksCountV1 Get tasks count - 6bb9-395b-4af9-8285
/* Returns the number of tasks that meet the filter criteria


@param GetTasksCountV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-tasks-count
*/
func (s *TaskService) GetTasksCountV1(GetTasksCountV1QueryParams *GetTasksCountV1QueryParams) (*ResponseTaskGetTasksCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/tasks/count"

	queryString, _ := query.Values(GetTasksCountV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseTaskGetTasksCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTasksCountV1(GetTasksCountV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTasksCountV1")
	}

	result := response.Result().(*ResponseTaskGetTasksCountV1)
	return result, response, err

}

//GetTasksByIDV1 Get tasks by ID - e493-ea85-4038-8183
/* Returns the task with the given ID


@param id id path parameter. the `id` of the task to retrieve


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-tasks-by-id
*/
func (s *TaskService) GetTasksByIDV1(id string) (*ResponseTaskGetTasksByIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/tasks/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTaskGetTasksByIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTasksByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation GetTasksByIdV1")
	}

	result := response.Result().(*ResponseTaskGetTasksByIDV1)
	return result, response, err

}

//GetTaskDetailsByIDV1 Get task details by ID - 408d-8acf-43fb-92c2
/* Returns the task details for the given task ID


@param id id path parameter. the `id` of the task to retrieve details for


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-task-details-by-id
*/
func (s *TaskService) GetTaskDetailsByIDV1(id string) (*ResponseTaskGetTaskDetailsByIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/tasks/{id}/detail"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTaskGetTaskDetailsByIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTaskDetailsByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation GetTaskDetailsByIdV1")
	}

	result := response.Result().(*ResponseTaskGetTaskDetailsByIDV1)
	return result, response, err

}

// Alias Function
/*
This method acts as an alias for the method `GetTaskCountV1`
*/
func (s *TaskService) GetTaskCount(GetTaskCountV1QueryParams *GetTaskCountV1QueryParams) (*ResponseTaskGetTaskCountV1, *resty.Response, error) {
	return s.GetTaskCountV1(GetTaskCountV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `RetrieveASpecificAssuranceTaskByIDV1`
*/
func (s *TaskService) RetrieveASpecificAssuranceTaskByID(id string, RetrieveASpecificAssuranceTaskByIdV1HeaderParams *RetrieveASpecificAssuranceTaskByIDV1HeaderParams) (*ResponseTaskRetrieveASpecificAssuranceTaskByIDV1, *resty.Response, error) {
	return s.RetrieveASpecificAssuranceTaskByIDV1(id, RetrieveASpecificAssuranceTaskByIdV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetTaskTreeV1`
*/
func (s *TaskService) GetTaskTree(taskID string) (*ResponseTaskGetTaskTreeV1, *resty.Response, error) {
	return s.GetTaskTreeV1(taskID)
}

// Alias Function
/*
This method acts as an alias for the method `GetTasksV1`
*/
func (s *TaskService) GetTasks(GetTasksV1QueryParams *GetTasksV1QueryParams) (*ResponseTaskGetTasksV1, *resty.Response, error) {
	return s.GetTasksV1(GetTasksV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetTaskDetailsByIDV1`
*/
func (s *TaskService) GetTaskDetailsByID(id string) (*ResponseTaskGetTaskDetailsByIDV1, *resty.Response, error) {
	return s.GetTaskDetailsByIDV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `GetTasksOperationalTasksV1`
*/
func (s *TaskService) GetTasksOperationalTasks(GetTasksOperationalTasksV1QueryParams *GetTasksOperationalTasksV1QueryParams) (*ResponseTaskGetTasksOperationalTasksV1, *resty.Response, error) {
	return s.GetTasksOperationalTasksV1(GetTasksOperationalTasksV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetTaskByIDV1`
*/
func (s *TaskService) GetTaskByID(taskID string) (*ResponseTaskGetTaskByIDV1, *resty.Response, error) {
	return s.GetTaskByIDV1(taskID)
}

// Alias Function
/*
This method acts as an alias for the method `RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistV1`
*/
func (s *TaskService) RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExist(RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistV1HeaderParams *RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistV1HeaderParams, RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistV1QueryParams *RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistV1QueryParams) (*ResponseTaskRetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistV1, *resty.Response, error) {
	return s.RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistV1(RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistV1HeaderParams, RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetBusinessAPIExecutionDetailsV1`
*/
func (s *TaskService) GetBusinessAPIExecutionDetails(executionID string) (*ResponseTaskGetBusinessAPIExecutionDetailsV1, *resty.Response, error) {
	return s.GetBusinessAPIExecutionDetailsV1(executionID)
}

// Alias Function
/*
This method acts as an alias for the method `GetTasksCountV1`
*/
func (s *TaskService) GetTasksCount(GetTasksCountV1QueryParams *GetTasksCountV1QueryParams) (*ResponseTaskGetTasksCountV1, *resty.Response, error) {
	return s.GetTasksCountV1(GetTasksCountV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `RetrieveAListOfAssuranceTasksV1`
*/
func (s *TaskService) RetrieveAListOfAssuranceTasks(RetrieveAListOfAssuranceTasksV1HeaderParams *RetrieveAListOfAssuranceTasksV1HeaderParams, RetrieveAListOfAssuranceTasksV1QueryParams *RetrieveAListOfAssuranceTasksV1QueryParams) (*ResponseTaskRetrieveAListOfAssuranceTasksV1, *resty.Response, error) {
	return s.RetrieveAListOfAssuranceTasksV1(RetrieveAListOfAssuranceTasksV1HeaderParams, RetrieveAListOfAssuranceTasksV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetTaskByOperationIDV1`
*/
func (s *TaskService) GetTaskByOperationID(operationID string, offset int, limit int) (*ResponseTaskGetTaskByOperationIDV1, *resty.Response, error) {
	return s.GetTaskByOperationIDV1(operationID, offset, limit)
}

// Alias Function
/*
This method acts as an alias for the method `GetTasksByIDV1`
*/
func (s *TaskService) GetTasksByID(id string) (*ResponseTaskGetTasksByIDV1, *resty.Response, error) {
	return s.GetTasksByIDV1(id)
}
