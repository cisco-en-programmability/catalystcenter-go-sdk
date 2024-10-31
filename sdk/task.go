package catalyst

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type TaskService service

type GetTasksOperationalTasksV1QueryParams struct {
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
	Offset        int    `url:"offset,omitempty"`        //offset
	Limit         int    `url:"limit,omitempty"`         //limit
	SortBy        string `url:"sortBy,omitempty"`        //Sort results by this field
	Order         string `url:"order,omitempty"`         //Sort order - asc or dsc
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
	Offset    int    `url:"offset,omitempty"`    //The first record to show for this page; the first record is numbered 1.
	Limit     int    `url:"limit,omitempty"`     //The number of records to show for this page.
	SortBy    string `url:"sortBy,omitempty"`    //A property within the response to sort by.
	Order     string `url:"order,omitempty"`     //Whether ascending or descending order should be used to sort the response.
	StartTime int    `url:"startTime,omitempty"` //This is the epoch millisecond start time from which tasks need to be fetched
	EndTime   int    `url:"endTime,omitempty"`   //This is the epoch millisecond end time upto which task records need to be fetched
	ParentID  string `url:"parentId,omitempty"`  //Fetch tasks that have this parent Id
	RootID    string `url:"rootId,omitempty"`    //Fetch tasks that have this root Id
	Status    string `url:"status,omitempty"`    //Fetch tasks that have this status. Available values : PENDING, FAILURE, SUCCESS
}
type GetTasksCountV1QueryParams struct {
	StartTime int    `url:"startTime,omitempty"` //This is the epoch millisecond start time from which tasks need to be fetched
	EndTime   int    `url:"endTime,omitempty"`   //This is the epoch millisecond end time upto which task records need to be fetched
	ParentID  string `url:"parentId,omitempty"`  //Fetch tasks that have this parent Id
	RootID    string `url:"rootId,omitempty"`    //Fetch tasks that have this root Id
	Status    string `url:"status,omitempty"`    //Fetch tasks that have this status. Available values : PENDING, FAILURE, SUCCESS
}

type ResponseTaskGetBusinessAPIExecutionDetailsV1 struct {
	BapiKey           string `json:"bapiKey,omitempty"`           // Business API Key (UUID)
	BapiName          string `json:"bapiName,omitempty"`          // Name of the Business API
	BapiExecutionID   string `json:"bapiExecutionId,omitempty"`   // Execution Id of the Business API (UUID)
	StartTime         string `json:"startTime,omitempty"`         // Execution Start Time of the Business API (Date Time Format)
	StartTimeEpoch    *int   `json:"startTimeEpoch,omitempty"`    // Execution Start Time of the Business API (Epoch Milliseconds)
	EndTime           string `json:"endTime,omitempty"`           // Execution End Time of the Business API (Date Time Format)
	EndTimeEpoch      *int   `json:"endTimeEpoch,omitempty"`      // Execution End Time of the Business API (Epoch Milliseconds)
	TimeDuration      *int   `json:"timeDuration,omitempty"`      // Time taken for Business API Execution (Milliseconds)
	Status            string `json:"status,omitempty"`            // Execution status of the Business API
	RuntimeInstanceID string `json:"runtimeInstanceId,omitempty"` // Pod Id in which the Business API is executed
	BapiError         string `json:"bapiError,omitempty"`         // Business API error message
}
type ResponseTaskGetTasksOperationalTasksV1 struct {
	Response *[]ResponseTaskGetTasksOperationalTasksV1Response `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  //
}
type ResponseTaskGetTasksOperationalTasksV1Response struct {
	AdditionalStatusURL string   `json:"additionalStatusURL,omitempty"` //
	Data                string   `json:"data,omitempty"`                //
	EndTime             *int     `json:"endTime,omitempty"`             //
	ErrorCode           string   `json:"errorCode,omitempty"`           //
	ErrorKey            string   `json:"errorKey,omitempty"`            //
	FailureReason       string   `json:"failureReason,omitempty"`       //
	ID                  string   `json:"id,omitempty"`                  //
	InstanceTenantID    string   `json:"instanceTenantId,omitempty"`    //
	IsError             *bool    `json:"isError,omitempty"`             //
	LastUpdate          *int     `json:"lastUpdate,omitempty"`          //
	OperationIDList     []string `json:"operationIdList,omitempty"`     //
	ParentID            string   `json:"parentId,omitempty"`            //
	Progress            string   `json:"progress,omitempty"`            //
	RootID              string   `json:"rootId,omitempty"`              //
	ServiceType         string   `json:"serviceType,omitempty"`         //
	StartTime           *int     `json:"startTime,omitempty"`           //
	Username            string   `json:"username,omitempty"`            //
	Version             *int     `json:"version,omitempty"`             //
}
type ResponseTaskGetTaskCountV1 struct {
	Response *int   `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}
type ResponseTaskGetTaskByOperationIDV1 struct {
	Response *[]ResponseTaskGetTaskByOperationIDV1Response `json:"response,omitempty"` //
	Version  string                                        `json:"version,omitempty"`  //
}
type ResponseTaskGetTaskByOperationIDV1Response struct {
	AdditionalStatusURL string   `json:"additionalStatusURL,omitempty"` //
	Data                string   `json:"data,omitempty"`                //
	EndTime             *int     `json:"endTime,omitempty"`             //
	ErrorCode           string   `json:"errorCode,omitempty"`           //
	ErrorKey            string   `json:"errorKey,omitempty"`            //
	FailureReason       string   `json:"failureReason,omitempty"`       //
	ID                  string   `json:"id,omitempty"`                  //
	InstanceTenantID    string   `json:"instanceTenantId,omitempty"`    //
	IsError             *bool    `json:"isError,omitempty"`             //
	LastUpdate          *int     `json:"lastUpdate,omitempty"`          //
	OperationIDList     []string `json:"operationIdList,omitempty"`     //
	ParentID            string   `json:"parentId,omitempty"`            //
	Progress            string   `json:"progress,omitempty"`            //
	RootID              string   `json:"rootId,omitempty"`              //
	ServiceType         string   `json:"serviceType,omitempty"`         //
	StartTime           *int     `json:"startTime,omitempty"`           //
	Username            string   `json:"username,omitempty"`            //
	Version             *int     `json:"version,omitempty"`             //
}
type ResponseTaskGetTaskByIDV1 struct {
	Response *ResponseTaskGetTaskByIDV1Response `json:"response,omitempty"` //
	Version  string                             `json:"version,omitempty"`  //
}
type ResponseTaskGetTaskByIDV1Response struct {
	AdditionalStatusURL string   `json:"additionalStatusURL,omitempty"` //
	Data                string   `json:"data,omitempty"`                //
	EndTime             *int     `json:"endTime,omitempty"`             //
	ErrorCode           string   `json:"errorCode,omitempty"`           //
	ErrorKey            string   `json:"errorKey,omitempty"`            //
	FailureReason       string   `json:"failureReason,omitempty"`       //
	ID                  string   `json:"id,omitempty"`                  //
	InstanceTenantID    string   `json:"instanceTenantId,omitempty"`    //
	IsError             *bool    `json:"isError,omitempty"`             //
	LastUpdate          *int     `json:"lastUpdate,omitempty"`          //
	OperationIDList     []string `json:"operationIdList,omitempty"`     //
	ParentID            string   `json:"parentId,omitempty"`            //
	Progress            string   `json:"progress,omitempty"`            //
	RootID              string   `json:"rootId,omitempty"`              //
	ServiceType         string   `json:"serviceType,omitempty"`         //
	StartTime           *int     `json:"startTime,omitempty"`           //
	Username            string   `json:"username,omitempty"`            //
	Version             *int     `json:"version,omitempty"`             //
}
type ResponseTaskGetTaskTreeV1 struct {
	Response *[]ResponseTaskGetTaskTreeV1Response `json:"response,omitempty"` //
	Version  string                               `json:"version,omitempty"`  //
}
type ResponseTaskGetTaskTreeV1Response struct {
	AdditionalStatusURL string   `json:"additionalStatusURL,omitempty"` //
	Data                string   `json:"data,omitempty"`                //
	EndTime             *int     `json:"endTime,omitempty"`             //
	ErrorCode           string   `json:"errorCode,omitempty"`           //
	ErrorKey            string   `json:"errorKey,omitempty"`            //
	FailureReason       string   `json:"failureReason,omitempty"`       //
	ID                  string   `json:"id,omitempty"`                  //
	InstanceTenantID    string   `json:"instanceTenantId,omitempty"`    //
	IsError             *bool    `json:"isError,omitempty"`             //
	LastUpdate          *int     `json:"lastUpdate,omitempty"`          //
	OperationIDList     []string `json:"operationIdList,omitempty"`     //
	ParentID            string   `json:"parentId,omitempty"`            //
	Progress            string   `json:"progress,omitempty"`            //
	RootID              string   `json:"rootId,omitempty"`              //
	ServiceType         string   `json:"serviceType,omitempty"`         //
	StartTime           *int     `json:"startTime,omitempty"`           //
	Username            string   `json:"username,omitempty"`            //
	Version             *int     `json:"version,omitempty"`             //
}
type ResponseTaskGetTasksV1 struct {
	Response *[]ResponseTaskGetTasksV1Response `json:"response,omitempty"` //
	Version  string                            `json:"version,omitempty"`  //
}
type ResponseTaskGetTasksV1Response struct {
	EndTime        *int   `json:"endTime,omitempty"`        // An approximate time of when this task has been marked completed; as measured in Unix epoch time in milliseconds
	ID             string `json:"id,omitempty"`             // The ID of this task
	UpdatedTime    *int   `json:"updatedTime,omitempty"`    // A timestamp of when this task was last updated; as measured in Unix epoch time in milliseconds
	ParentID       string `json:"parentId,omitempty"`       // The ID of the parent task if this happens to be a subtask. In case this task is not a subtask, then the parentId is expected to be null.  To construct a task tree, this task will be the child of the task with the ID listed here, or the root of the tree if this task has no parentId.
	ResultLocation string `json:"resultLocation,omitempty"` // A server-relative URL indicating where additional task-specific details may be found
	RootID         string `json:"rootId,omitempty"`         // The ID of the task representing the root node of the tree which this task belongs to.  In some cases, this may be the same as the ID or null, which indicates that this task is the root task.
	StartTime      *int   `json:"startTime,omitempty"`      // An approximate time of when the task creation was triggered; as measured in Unix epoch time in milliseconds
	Status         string `json:"status,omitempty"`         //
}
type ResponseTaskGetTasksCountV1 struct {
	Response *ResponseTaskGetTasksCountV1Response `json:"response,omitempty"` //
	Version  string                               `json:"version,omitempty"`  // The version of the response
}
type ResponseTaskGetTasksCountV1Response struct {
	Count *int `json:"count,omitempty"` // The reported count.
}
type ResponseTaskGetTasksByIDV1 struct {
	Response *ResponseTaskGetTasksByIDV1Response `json:"response,omitempty"` //
	Version  string                              `json:"version,omitempty"`  //
}
type ResponseTaskGetTasksByIDV1Response struct {
	EndTime        *int   `json:"endTime,omitempty"`        // An approximate time of when this task has been marked completed; as measured in Unix epoch time in milliseconds
	ID             string `json:"id,omitempty"`             // The ID of this task
	UpdatedTime    *int   `json:"updatedTime,omitempty"`    // A timestamp of when this task was last updated; as measured in Unix epoch time in milliseconds
	ParentID       string `json:"parentId,omitempty"`       // The ID of the parent task if this happens to be a subtask. In case this task is not a subtask, then the parentId is expected to be null.  To construct a task tree, this task will be the child of the task with the ID listed here, or the root of the tree if this task has no parentId.
	ResultLocation string `json:"resultLocation,omitempty"` // A server-relative URL indicating where additional task-specific details may be found
	RootID         string `json:"rootId,omitempty"`         // The ID of the task representing the root node of the tree which this task belongs to.  In some cases, this may be the same as the ID or null, which indicates that this task is the root task.
	StartTime      *int   `json:"startTime,omitempty"`      // An approximate time of when the task creation was triggered; as measured in Unix epoch time in milliseconds
	Status         string `json:"status,omitempty"`         //
}
type ResponseTaskGetTaskDetailsByIDV1 struct {
	Response *ResponseTaskGetTaskDetailsByIDV1Response `json:"response,omitempty"` //
	Version  string                                    `json:"version,omitempty"`  //
}
type ResponseTaskGetTaskDetailsByIDV1Response struct {
	Data          string `json:"data,omitempty"`          // Any data associated with this task; the value may vary significantly across different tasks
	Progress      string `json:"progress,omitempty"`      // A textual representation which indicates the progress of this task; the value may vary significantly across different tasks
	ErrorCode     string `json:"errorCode,omitempty"`     // An error code if in case this task has failed in its execution
	FailureReason string `json:"failureReason,omitempty"` // A textual description indicating the reason why a task has failed
}

//GetBusinessAPIExecutionDetailsV1 Get Business API Execution Details - c1bc-a8c1-41fb-9f75
/* Retrieves the execution details of a Business API


@param executionID executionId path parameter. Execution Id of API


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-business-api-execution-details-v1
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

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-tasks-operational-tasks-v1
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

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-task-count-v1
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


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-task-by-operation-id-v1
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


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-task-by-id-v1
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


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-task-tree-v1
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

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-tasks-v1
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

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-tasks-count-v1
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


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-tasks-by-id-v1
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


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-task-details-by-id-v1
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
func (s *TaskService) GetTaskCount(GetTaskCountV1QueryParams *GetTaskCountV1QueryParams) (*ResponseTaskGetTaskCountV1, *resty.Response, error) {
	return s.GetTaskCountV1(GetTaskCountV1QueryParams)
}

// Alias Function
func (s *TaskService) GetTaskTree(taskID string) (*ResponseTaskGetTaskTreeV1, *resty.Response, error) {
	return s.GetTaskTreeV1(taskID)
}

// Alias Function
func (s *TaskService) GetTasks(GetTasksV1QueryParams *GetTasksV1QueryParams) (*ResponseTaskGetTasksV1, *resty.Response, error) {
	return s.GetTasksV1(GetTasksV1QueryParams)
}

// Alias Function
func (s *TaskService) GetTaskDetailsByID(id string) (*ResponseTaskGetTaskDetailsByIDV1, *resty.Response, error) {
	return s.GetTaskDetailsByIDV1(id)
}

// Alias Function
func (s *TaskService) GetTasksOperationalTasks(GetTasksOperationalTasksV1QueryParams *GetTasksOperationalTasksV1QueryParams) (*ResponseTaskGetTasksOperationalTasksV1, *resty.Response, error) {
	return s.GetTasksOperationalTasksV1(GetTasksOperationalTasksV1QueryParams)
}

// Alias Function
func (s *TaskService) GetTaskByID(taskID string) (*ResponseTaskGetTaskByIDV1, *resty.Response, error) {
	return s.GetTaskByIDV1(taskID)
}

// Alias Function
func (s *TaskService) GetBusinessAPIExecutionDetails(executionID string) (*ResponseTaskGetBusinessAPIExecutionDetailsV1, *resty.Response, error) {
	return s.GetBusinessAPIExecutionDetailsV1(executionID)
}

// Alias Function
func (s *TaskService) GetTasksCount(GetTasksCountV1QueryParams *GetTasksCountV1QueryParams) (*ResponseTaskGetTasksCountV1, *resty.Response, error) {
	return s.GetTasksCountV1(GetTasksCountV1QueryParams)
}

// Alias Function
func (s *TaskService) GetTaskByOperationID(operationID string, offset int, limit int) (*ResponseTaskGetTaskByOperationIDV1, *resty.Response, error) {
	return s.GetTaskByOperationIDV1(operationID, offset, limit)
}

// Alias Function
func (s *TaskService) GetTasksByID(id string) (*ResponseTaskGetTasksByIDV1, *resty.Response, error) {
	return s.GetTasksByIDV1(id)
}
