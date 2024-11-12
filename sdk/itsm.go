package catalyst

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type ItsmService service

type GetCmdbSyncStatusV1QueryParams struct {
	Status string `url:"status,omitempty"` //Supported values are "Success","Failed" and "Unknown". Providing other values will result in all the available sync job status.
	Date   string `url:"date,omitempty"`   //Provide date in "YYYY-MM-DD" format
}
type GetFailedItsmEventsV1QueryParams struct {
	InstanceID string `url:"instanceId,omitempty"` //Instance Id of the failed event as in the Runtime Dashboard
}

type ResponseItsmGetCmdbSyncStatusV1 []ResponseItemItsmGetCmdbSyncStatusV1 // Array of ResponseItsmGetCMDBSyncStatusV1
type ResponseItemItsmGetCmdbSyncStatusV1 struct {
	SuccessCount      string                                        `json:"successCount,omitempty"`      // Successfully synchronized device count
	FailureCount      string                                        `json:"failureCount,omitempty"`      // Failed device count
	Devices           *[]ResponseItemItsmGetCmdbSyncStatusV1Devices `json:"devices,omitempty"`           //
	UnknownErrorCount string                                        `json:"unknownErrorCount,omitempty"` // Unknown error count
	Message           string                                        `json:"message,omitempty"`           // Message
	SyncTime          string                                        `json:"syncTime,omitempty"`          // Synchronization Time
}
type ResponseItemItsmGetCmdbSyncStatusV1Devices struct {
	DeviceID string `json:"deviceId,omitempty"` // Device Id
	Status   string `json:"status,omitempty"`   // Status "Success" or "Failed"
}
type ResponseItsmGetFailedItsmEventsV1 []ResponseItemItsmGetFailedItsmEventsV1 // Array of ResponseItsmGetFailedITSMEventsV1
type ResponseItemItsmGetFailedItsmEventsV1 struct {
	InstanceID     string                                               `json:"instanceId,omitempty"`     // Instance Id
	EventID        string                                               `json:"eventId,omitempty"`        // Event Id
	Name           string                                               `json:"name,omitempty"`           // Name
	Type           string                                               `json:"type,omitempty"`           // Type
	Category       string                                               `json:"category,omitempty"`       // Category
	Domain         string                                               `json:"domain,omitempty"`         // Domain
	SubDomain      string                                               `json:"subDomain,omitempty"`      // Sub Domain
	Severity       string                                               `json:"severity,omitempty"`       // Severity
	Source         string                                               `json:"source,omitempty"`         // Source
	Timestamp      *int                                                 `json:"timestamp,omitempty"`      // Timestamp
	EnrichmentInfo *ResponseItemItsmGetFailedItsmEventsV1EnrichmentInfo `json:"enrichmentInfo,omitempty"` //
	Description    string                                               `json:"description,omitempty"`    // Description
}
type ResponseItemItsmGetFailedItsmEventsV1EnrichmentInfo struct {
	EventStatus                    string                                                                             `json:"eventStatus,omitempty"`                    // Event Status
	ErrorCode                      string                                                                             `json:"errorCode,omitempty"`                      // Error Code
	ErrorDescription               string                                                                             `json:"errorDescription,omitempty"`               // Error Description
	ResponseReceivedFromITSmsystem *ResponseItemItsmGetFailedItsmEventsV1EnrichmentInfoResponseReceivedFromITSmsystem `json:"responseReceivedFromITSMSystem,omitempty"` // Response Received From ITSMSystem
}
type ResponseItemItsmGetFailedItsmEventsV1EnrichmentInfoResponseReceivedFromITSmsystem interface{}
type ResponseItsmRetryIntegrationEventsV1 struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type RequestItsmRetryIntegrationEventsV1 []string // Array of RequestItsmRetryIntegrationEventsV1

//GetCmdbSyncStatusV1 Get CMDB Sync Status - a492-8993-4948-b86c
/* This API allows to retrieve the detail of CMDB sync status.It accepts two query parameter "status","date".The supported values for status field are "Success","Failed","Unknown" and date field should be in "YYYY-MM-DD" format. By default all the cmdb sync status will be send as response and based on the query parameter filtered detail will be send as response.


@param GetCMDBSyncStatusV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-cmdb-sync-status-v1
*/
func (s *ItsmService) GetCmdbSyncStatusV1(GetCMDBSyncStatusV1QueryParams *GetCmdbSyncStatusV1QueryParams) (*ResponseItsmGetCmdbSyncStatusV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/cmdb-sync/detail"

	queryString, _ := query.Values(GetCMDBSyncStatusV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseItsmGetCmdbSyncStatusV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCmdbSyncStatusV1(GetCMDBSyncStatusV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetCmdbSyncStatusV1")
	}

	result := response.Result().(*ResponseItsmGetCmdbSyncStatusV1)
	return result, response, err

}

//GetFailedItsmEventsV1 Get Failed ITSM Events - a293-b82a-42a8-ab15
/* Used to retrieve the list of integration events that failed to create tickets in ITSM


@param GetFailedITSMEventsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-failed-itsm-events-v1
*/
func (s *ItsmService) GetFailedItsmEventsV1(GetFailedITSMEventsV1QueryParams *GetFailedItsmEventsV1QueryParams) (*ResponseItsmGetFailedItsmEventsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/integration/events"

	queryString, _ := query.Values(GetFailedITSMEventsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseItsmGetFailedItsmEventsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFailedItsmEventsV1(GetFailedITSMEventsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetFailedItsmEventsV1")
	}

	result := response.Result().(*ResponseItsmGetFailedItsmEventsV1)
	return result, response, err

}

//RetryIntegrationEventsV1 Retry Integration Events - fa9a-9817-4129-af50
/* Allows retry of multiple failed ITSM event instances. The retry request payload can be given as a list of strings: ["instance1","instance2","instance3",..] A minimum of one instance Id is mandatory. The list of failed event instance Ids can be retrieved using the 'Get Failed ITSM Events' API in the 'instanceId' attribute.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!retry-integration-events-v1
*/
func (s *ItsmService) RetryIntegrationEventsV1(requestItsmRetryIntegrationEventsV1 *RequestItsmRetryIntegrationEventsV1) (*ResponseItsmRetryIntegrationEventsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/integration/events"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestItsmRetryIntegrationEventsV1).
		SetResult(&ResponseItsmRetryIntegrationEventsV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetryIntegrationEventsV1(requestItsmRetryIntegrationEventsV1)
		}

		return nil, response, fmt.Errorf("error with operation RetryIntegrationEventsV1")
	}

	result := response.Result().(*ResponseItsmRetryIntegrationEventsV1)
	return result, response, err

}

// Alias Function
/*
This method acts as an alias for the method `GetCmdbSyncStatusV1`
*/
func (s *ItsmService) GetCmdbSyncStatus(GetCMDBSyncStatusV1QueryParams *GetCmdbSyncStatusV1QueryParams) (*ResponseItsmGetCmdbSyncStatusV1, *resty.Response, error) {
	return s.GetCmdbSyncStatusV1(GetCMDBSyncStatusV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetFailedItsmEventsV1`
*/
func (s *ItsmService) GetFailedItsmEvents(GetFailedITSMEventsV1QueryParams *GetFailedItsmEventsV1QueryParams) (*ResponseItsmGetFailedItsmEventsV1, *resty.Response, error) {
	return s.GetFailedItsmEventsV1(GetFailedITSMEventsV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `RetryIntegrationEventsV1`
*/
func (s *ItsmService) RetryIntegrationEvents(requestItsmRetryIntegrationEventsV1 *RequestItsmRetryIntegrationEventsV1) (*ResponseItsmRetryIntegrationEventsV1, *resty.Response, error) {
	return s.RetryIntegrationEventsV1(requestItsmRetryIntegrationEventsV1)
}
