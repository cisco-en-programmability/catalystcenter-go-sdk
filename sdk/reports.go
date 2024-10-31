package catalyst

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type ReportsService service

type GetListOfScheduledReportsV1QueryParams struct {
	ViewGroupID string `url:"viewGroupId,omitempty"` //viewGroupId of viewgroup for report
	ViewID      string `url:"viewId,omitempty"`      //viewId of view for report
}

// # Review unknown case
type ResponseReportsExecutingTheFlexibleReportV1 struct {
	ExecutionID   string                                                 `json:"executionId,omitempty"`   // Report ExecutionId (Unique UUID)
	StartTime     *float64                                               `json:"startTime,omitempty"`     // Report execution start time (Represent the specified number of milliseconds since the epoch time)
	EndTime       *float64                                               `json:"endTime,omitempty"`       // Report execution end time (Represent the specified number of milliseconds since the epoch time)
	RequestStatus string                                                 `json:"requestStatus,omitempty"` // Report  request status
	Errors        []string                                               `json:"errors,omitempty"`        // Errors associated to the report execution
	Warnings      *[]ResponseReportsExecutingTheFlexibleReportV1Warnings `json:"warnings,omitempty"`      // Warnings associated to the report execution
}
type ResponseReportsExecutingTheFlexibleReportV1Warnings interface{}
type ResponseReportsGetExecutionIDByReportIDV1 struct {
	ReportID          string                                                 `json:"reportId,omitempty"`          // Report Id (Unique UUID)
	ReportName        string                                                 `json:"reportName,omitempty"`        // Name of the report
	Executions        *[]ResponseReportsGetExecutionIDByReportIDV1Executions `json:"executions,omitempty"`        //
	ExecutionCount    *int                                                   `json:"executionCount,omitempty"`    // Total number of report executions
	ReportWasExecuted *bool                                                  `json:"reportWasExecuted,omitempty"` // Report execution status flag (true if execution is started, false if the execution is not started)
}
type ResponseReportsGetExecutionIDByReportIDV1Executions struct {
	ExecutionID   string                                                         `json:"executionId,omitempty"`   // Report ExecutionId (Unique UUID)
	StartTime     *int                                                           `json:"startTime,omitempty"`     // Report execution start time (Represent the specified number of milliseconds since the epoch time)
	EndTime       *float64                                                       `json:"endTime,omitempty"`       // Report execution end time (Represent the specified number of milliseconds since the epoch time)
	ProcessStatus string                                                         `json:"processStatus,omitempty"` // Report execution status
	RequestStatus string                                                         `json:"requestStatus,omitempty"` // Report request status
	Errors        []string                                                       `json:"errors,omitempty"`        // Errors associated with the report execution
	Warnings      *[]ResponseReportsGetExecutionIDByReportIDV1ExecutionsWarnings `json:"warnings,omitempty"`      // Warnings associated with the report execution
}
type ResponseReportsGetExecutionIDByReportIDV1ExecutionsWarnings interface{}
type ResponseReportsUpdateScheduleOfFlexibleReportV1 struct {
	Schedule *ResponseReportsUpdateScheduleOfFlexibleReportV1Schedule `json:"schedule,omitempty"` // Schedule information
}
type ResponseReportsUpdateScheduleOfFlexibleReportV1Schedule interface{}
type ResponseReportsGetFlexibleReportScheduleByReportIDV1 struct {
	Schedule *ResponseReportsGetFlexibleReportScheduleByReportIDV1Schedule `json:"schedule,omitempty"` // Schedule information
}
type ResponseReportsGetFlexibleReportScheduleByReportIDV1Schedule interface{}
type ResponseReportsGetAllFlexibleReportSchedulesV1 []ResponseItemReportsGetAllFlexibleReportSchedulesV1 // Array of ResponseReportsGetAllFlexibleReportSchedulesV1
type ResponseItemReportsGetAllFlexibleReportSchedulesV1 struct {
	ReportID   string                                                      `json:"reportId,omitempty"`   // Report Id (Unique UUID)
	Schedule   *ResponseItemReportsGetAllFlexibleReportSchedulesV1Schedule `json:"schedule,omitempty"`   // Schedule information
	ReportName string                                                      `json:"reportName,omitempty"` // Name of the report
}
type ResponseItemReportsGetAllFlexibleReportSchedulesV1Schedule interface{}
type ResponseReportsCreateOrScheduleAReportV1 struct {
	Tags              []string                                              `json:"tags,omitempty"`              // array of tags for report
	DataCategory      string                                                `json:"dataCategory,omitempty"`      // data category of the report
	Deliveries        *[]ResponseReportsCreateOrScheduleAReportV1Deliveries `json:"deliveries,omitempty"`        // Array of available delivery channels
	ExecutionCount    *int                                                  `json:"executionCount,omitempty"`    // Total number of report executions
	Executions        *[]ResponseReportsCreateOrScheduleAReportV1Executions `json:"executions,omitempty"`        //
	Name              string                                                `json:"name,omitempty"`              // report name
	ReportID          string                                                `json:"reportId,omitempty"`          // report Id
	ReportWasExecuted *bool                                                 `json:"reportWasExecuted,omitempty"` // true if atleast one execution has started
	Schedule          *ResponseReportsCreateOrScheduleAReportV1Schedule     `json:"schedule,omitempty"`          //
	View              *ResponseReportsCreateOrScheduleAReportV1View         `json:"view,omitempty"`              //
	ViewGroupID       string                                                `json:"viewGroupId,omitempty"`       // viewGroupId of the viewgroup for the report
	ViewGroupVersion  string                                                `json:"viewGroupVersion,omitempty"`  // version of viewgroup for the report
}
type ResponseReportsCreateOrScheduleAReportV1Deliveries interface{}
type ResponseReportsCreateOrScheduleAReportV1Executions struct {
	EndTime       *int     `json:"endTime,omitempty"`       // Report execution pipeline end time
	Errors        []string `json:"errors,omitempty"`        //
	ExecutionID   string   `json:"executionId,omitempty"`   // Report execution Id.
	ProcessStatus string   `json:"processStatus,omitempty"` // Report execution status
	RequestStatus string   `json:"requestStatus,omitempty"` // Report execution acceptance status from scheduler
	StartTime     *int     `json:"startTime,omitempty"`     // Report execution pipeline start time
	Warnings      []string `json:"warnings,omitempty"`      //
}
type ResponseReportsCreateOrScheduleAReportV1Schedule interface{}
type ResponseReportsCreateOrScheduleAReportV1View struct {
	FieldGroups *[]ResponseReportsCreateOrScheduleAReportV1ViewFieldGroups `json:"fieldGroups,omitempty"` //
	Filters     *[]ResponseReportsCreateOrScheduleAReportV1ViewFilters     `json:"filters,omitempty"`     //
	Format      *ResponseReportsCreateOrScheduleAReportV1ViewFormat        `json:"format,omitempty"`      //
	Name        string                                                     `json:"name,omitempty"`        // view name
	ViewID      string                                                     `json:"viewId,omitempty"`      // view Id
	Description string                                                     `json:"description,omitempty"` // view description
	ViewInfo    string                                                     `json:"viewInfo,omitempty"`    // view filters info
}
type ResponseReportsCreateOrScheduleAReportV1ViewFieldGroups struct {
	FieldGroupDisplayName string                                                           `json:"fieldGroupDisplayName,omitempty"` // Field group label/displayname for user
	FieldGroupName        string                                                           `json:"fieldGroupName,omitempty"`        // Field group name
	Fields                *[]ResponseReportsCreateOrScheduleAReportV1ViewFieldGroupsFields `json:"fields,omitempty"`                //
}
type ResponseReportsCreateOrScheduleAReportV1ViewFieldGroupsFields struct {
	DisplayName string `json:"displayName,omitempty"` // field label/displayname
	Name        string `json:"name,omitempty"`        // field name
}
type ResponseReportsCreateOrScheduleAReportV1ViewFilters struct {
	DisplayName string                                                    `json:"displayName,omitempty"` // filter label/displayname
	Name        string                                                    `json:"name,omitempty"`        // filter name
	Type        string                                                    `json:"type,omitempty"`        // filter type
	Value       *ResponseReportsCreateOrScheduleAReportV1ViewFiltersValue `json:"value,omitempty"`       // value of filter. data type is based on the filter type.
}
type ResponseReportsCreateOrScheduleAReportV1ViewFiltersValue interface{}
type ResponseReportsCreateOrScheduleAReportV1ViewFormat struct {
	FormatType string `json:"formatType,omitempty"` // format type of report
	Name       string `json:"name,omitempty"`       // format name of report
}
type ResponseReportsGetListOfScheduledReportsV1 []ResponseItemReportsGetListOfScheduledReportsV1 // Array of ResponseReportsGetListOfScheduledReportsV1
type ResponseItemReportsGetListOfScheduledReportsV1 struct {
	Tags              []string                                                    `json:"tags,omitempty"`              // array of tags for report
	DataCategory      string                                                      `json:"dataCategory,omitempty"`      // data category of the report
	Deliveries        *[]ResponseItemReportsGetListOfScheduledReportsV1Deliveries `json:"deliveries,omitempty"`        // Array of available delivery channels
	ExecutionCount    *int                                                        `json:"executionCount,omitempty"`    // Total number of report executions
	Executions        *[]ResponseItemReportsGetListOfScheduledReportsV1Executions `json:"executions,omitempty"`        //
	Name              string                                                      `json:"name,omitempty"`              // report name
	ReportID          string                                                      `json:"reportId,omitempty"`          // report Id
	ReportWasExecuted *bool                                                       `json:"reportWasExecuted,omitempty"` // true if atleast one execution has started
	Schedule          *ResponseItemReportsGetListOfScheduledReportsV1Schedule     `json:"schedule,omitempty"`          //
	View              *ResponseItemReportsGetListOfScheduledReportsV1View         `json:"view,omitempty"`              //
	ViewGroupID       string                                                      `json:"viewGroupId,omitempty"`       // viewGroupId of the viewgroup for the report
	ViewGroupVersion  string                                                      `json:"viewGroupVersion,omitempty"`  // version of viewgroup for the report
}
type ResponseItemReportsGetListOfScheduledReportsV1Deliveries interface{}
type ResponseItemReportsGetListOfScheduledReportsV1Executions struct {
	EndTime       *int     `json:"endTime,omitempty"`       // Report execution pipeline end time
	Errors        []string `json:"errors,omitempty"`        //
	ExecutionID   string   `json:"executionId,omitempty"`   // Report execution Id.
	ProcessStatus string   `json:"processStatus,omitempty"` // Report execution status
	RequestStatus string   `json:"requestStatus,omitempty"` // Report execution acceptance status from scheduler
	StartTime     *int     `json:"startTime,omitempty"`     // Report execution pipeline start time
	Warnings      []string `json:"warnings,omitempty"`      //
}
type ResponseItemReportsGetListOfScheduledReportsV1Schedule interface{}
type ResponseItemReportsGetListOfScheduledReportsV1View struct {
	FieldGroups *[]ResponseItemReportsGetListOfScheduledReportsV1ViewFieldGroups `json:"fieldGroups,omitempty"` //
	Filters     *[]ResponseItemReportsGetListOfScheduledReportsV1ViewFilters     `json:"filters,omitempty"`     //
	Format      *ResponseItemReportsGetListOfScheduledReportsV1ViewFormat        `json:"format,omitempty"`      //
	Name        string                                                           `json:"name,omitempty"`        // view name
	ViewID      string                                                           `json:"viewId,omitempty"`      // view Id
	Description string                                                           `json:"description,omitempty"` // view description
	ViewInfo    string                                                           `json:"viewInfo,omitempty"`    // view filters info
}
type ResponseItemReportsGetListOfScheduledReportsV1ViewFieldGroups struct {
	FieldGroupDisplayName string                                                                 `json:"fieldGroupDisplayName,omitempty"` // Field group label/displayname for user
	FieldGroupName        string                                                                 `json:"fieldGroupName,omitempty"`        // Field group name
	Fields                *[]ResponseItemReportsGetListOfScheduledReportsV1ViewFieldGroupsFields `json:"fields,omitempty"`                //
}
type ResponseItemReportsGetListOfScheduledReportsV1ViewFieldGroupsFields struct {
	DisplayName string `json:"displayName,omitempty"` // field label/displayname
	Name        string `json:"name,omitempty"`        // field name
}
type ResponseItemReportsGetListOfScheduledReportsV1ViewFilters struct {
	DisplayName string                                                          `json:"displayName,omitempty"` // filter label/displayname
	Name        string                                                          `json:"name,omitempty"`        // filter name
	Type        string                                                          `json:"type,omitempty"`        // filter type
	Value       *ResponseItemReportsGetListOfScheduledReportsV1ViewFiltersValue `json:"value,omitempty"`       // value of filter. data type is based on the filter type.
}
type ResponseItemReportsGetListOfScheduledReportsV1ViewFiltersValue interface{}
type ResponseItemReportsGetListOfScheduledReportsV1ViewFormat struct {
	FormatType string `json:"formatType,omitempty"` // format type of report
	Name       string `json:"name,omitempty"`       // format name of report
	Default    *bool  `json:"default,omitempty"`    // true, if the format type is considered default
}
type ResponseReportsGetAScheduledReportV1 struct {
	Tags              []string                                          `json:"tags,omitempty"`              // array of tags for report
	DataCategory      string                                            `json:"dataCategory,omitempty"`      // data category of the report
	Deliveries        *[]ResponseReportsGetAScheduledReportV1Deliveries `json:"deliveries,omitempty"`        // Array of available delivery channels
	ExecutionCount    *int                                              `json:"executionCount,omitempty"`    // Total number of report executions
	Executions        *[]ResponseReportsGetAScheduledReportV1Executions `json:"executions,omitempty"`        //
	Name              string                                            `json:"name,omitempty"`              // report name
	ReportID          string                                            `json:"reportId,omitempty"`          // report Id
	ReportWasExecuted *bool                                             `json:"reportWasExecuted,omitempty"` // true if atleast one execution has started
	Schedule          *ResponseReportsGetAScheduledReportV1Schedule     `json:"schedule,omitempty"`          //
	View              *ResponseReportsGetAScheduledReportV1View         `json:"view,omitempty"`              //
	ViewGroupID       string                                            `json:"viewGroupId,omitempty"`       // viewGroupId of the viewgroup for the report
	ViewGroupVersion  string                                            `json:"viewGroupVersion,omitempty"`  // version of viewgroup for the report
}
type ResponseReportsGetAScheduledReportV1Deliveries interface{}
type ResponseReportsGetAScheduledReportV1Executions struct {
	EndTime       *int     `json:"endTime,omitempty"`       // Report execution pipeline end time
	Errors        []string `json:"errors,omitempty"`        //
	ExecutionID   string   `json:"executionId,omitempty"`   // Report execution Id.
	ProcessStatus string   `json:"processStatus,omitempty"` // Report execution status
	RequestStatus string   `json:"requestStatus,omitempty"` // Report execution acceptance status from scheduler
	StartTime     *int     `json:"startTime,omitempty"`     // Report execution pipeline start time
	Warnings      []string `json:"warnings,omitempty"`      //
}
type ResponseReportsGetAScheduledReportV1Schedule interface{}
type ResponseReportsGetAScheduledReportV1View struct {
	FieldGroups *[]ResponseReportsGetAScheduledReportV1ViewFieldGroups `json:"fieldGroups,omitempty"` //
	Filters     *[]ResponseReportsGetAScheduledReportV1ViewFilters     `json:"filters,omitempty"`     //
	Format      *ResponseReportsGetAScheduledReportV1ViewFormat        `json:"format,omitempty"`      //
	Name        string                                                 `json:"name,omitempty"`        // view name
	ViewID      string                                                 `json:"viewId,omitempty"`      // view Id
	Description string                                                 `json:"description,omitempty"` // view description
	ViewInfo    string                                                 `json:"viewInfo,omitempty"`    // view filters info
}
type ResponseReportsGetAScheduledReportV1ViewFieldGroups struct {
	FieldGroupDisplayName string                                                       `json:"fieldGroupDisplayName,omitempty"` // Field group label/displayname for user
	FieldGroupName        string                                                       `json:"fieldGroupName,omitempty"`        // Field group name
	Fields                *[]ResponseReportsGetAScheduledReportV1ViewFieldGroupsFields `json:"fields,omitempty"`                //
}
type ResponseReportsGetAScheduledReportV1ViewFieldGroupsFields struct {
	DisplayName string `json:"displayName,omitempty"` // field label/displayname
	Name        string `json:"name,omitempty"`        // field name
}
type ResponseReportsGetAScheduledReportV1ViewFilters struct {
	DisplayName string                                                `json:"displayName,omitempty"` // filter label/displayname
	Name        string                                                `json:"name,omitempty"`        // filter name
	Type        string                                                `json:"type,omitempty"`        // filter type
	Value       *ResponseReportsGetAScheduledReportV1ViewFiltersValue `json:"value,omitempty"`       // value of filter. data type is based on the filter type.
}
type ResponseReportsGetAScheduledReportV1ViewFiltersValue interface{}
type ResponseReportsGetAScheduledReportV1ViewFormat struct {
	FormatType string `json:"formatType,omitempty"` // format type of report
	Name       string `json:"name,omitempty"`       // format name of report
	Default    *bool  `json:"default,omitempty"`    // true, if the format type is considered default
}
type ResponseReportsDeleteAScheduledReportV1 struct {
	Message string `json:"message,omitempty"` // Response message
	Status  *int   `json:"status,omitempty"`  // Response Status
}
type ResponseReportsGetAllExecutionDetailsForAGivenReportV1 struct {
	Tags              []string                                                            `json:"tags,omitempty"`              // array of tags for report
	DataCategory      string                                                              `json:"dataCategory,omitempty"`      // data category of the report
	Deliveries        *[]ResponseReportsGetAllExecutionDetailsForAGivenReportV1Deliveries `json:"deliveries,omitempty"`        // Array of available delivery channels
	ExecutionCount    *int                                                                `json:"executionCount,omitempty"`    // Total number of report executions
	Executions        *[]ResponseReportsGetAllExecutionDetailsForAGivenReportV1Executions `json:"executions,omitempty"`        //
	Name              string                                                              `json:"name,omitempty"`              // report dataset name
	ReportID          string                                                              `json:"reportId,omitempty"`          // report Id
	ReportWasExecuted *bool                                                               `json:"reportWasExecuted,omitempty"` // true if atleast one execution has started
	Schedule          *ResponseReportsGetAllExecutionDetailsForAGivenReportV1Schedule     `json:"schedule,omitempty"`          //
	View              *ResponseReportsGetAllExecutionDetailsForAGivenReportV1View         `json:"view,omitempty"`              //
	ViewGroupID       string                                                              `json:"viewGroupId,omitempty"`       // viewGroupId of the viewgroup for the report
	ViewGroupVersion  string                                                              `json:"viewGroupVersion,omitempty"`  // version of viewgroup for the report
}
type ResponseReportsGetAllExecutionDetailsForAGivenReportV1Deliveries interface{}
type ResponseReportsGetAllExecutionDetailsForAGivenReportV1Executions struct {
	EndTime       *int     `json:"endTime,omitempty"`       // Report execution pipeline end time
	Errors        []string `json:"errors,omitempty"`        //
	ExecutionID   string   `json:"executionId,omitempty"`   // Report execution Id.
	ProcessStatus string   `json:"processStatus,omitempty"` // Report execution status
	RequestStatus string   `json:"requestStatus,omitempty"` // Report execution acceptance status from scheduler
	StartTime     *int     `json:"startTime,omitempty"`     // Report execution pipeline start time
	Warnings      []string `json:"warnings,omitempty"`      //
}
type ResponseReportsGetAllExecutionDetailsForAGivenReportV1Schedule interface{}
type ResponseReportsGetAllExecutionDetailsForAGivenReportV1View struct {
	FieldGroups *[]ResponseReportsGetAllExecutionDetailsForAGivenReportV1ViewFieldGroups `json:"fieldGroups,omitempty"` //
	Filters     *[]ResponseReportsGetAllExecutionDetailsForAGivenReportV1ViewFilters     `json:"filters,omitempty"`     //
	Format      *ResponseReportsGetAllExecutionDetailsForAGivenReportV1ViewFormat        `json:"format,omitempty"`      //
	Name        string                                                                   `json:"name,omitempty"`        // view name
	ViewID      string                                                                   `json:"viewId,omitempty"`      // view Id
	Description string                                                                   `json:"description,omitempty"` // view description
	ViewInfo    string                                                                   `json:"viewInfo,omitempty"`    // view filters info
}
type ResponseReportsGetAllExecutionDetailsForAGivenReportV1ViewFieldGroups interface{}
type ResponseReportsGetAllExecutionDetailsForAGivenReportV1ViewFilters interface{}
type ResponseReportsGetAllExecutionDetailsForAGivenReportV1ViewFormat interface{} // # Review unknown case
type ResponseReportsGetAllViewGroupsV1 []ResponseItemReportsGetAllViewGroupsV1    // Array of ResponseReportsGetAllViewGroupsV1
type ResponseItemReportsGetAllViewGroupsV1 struct {
	Category    string `json:"category,omitempty"`    // category of the view group
	Description string `json:"description,omitempty"` // view group description
	Name        string `json:"name,omitempty"`        // name of view group
	ViewGroupID string `json:"viewGroupId,omitempty"` // id of viewgroup
}
type ResponseReportsGetViewsForAGivenViewGroupV1 struct {
	ViewGroupID string                                              `json:"viewGroupId,omitempty"` // viewgroup Id
	Views       *[]ResponseReportsGetViewsForAGivenViewGroupV1Views `json:"views,omitempty"`       //
}
type ResponseReportsGetViewsForAGivenViewGroupV1Views struct {
	Description string `json:"description,omitempty"` //
	ViewID      string `json:"viewId,omitempty"`      // Unique id for a view within viewgroup
	ViewName    string `json:"viewName,omitempty"`    // view name
}
type ResponseReportsGetViewDetailsForAGivenViewGroupViewV1 struct {
	Deliveries  *[]ResponseReportsGetViewDetailsForAGivenViewGroupViewV1Deliveries  `json:"deliveries,omitempty"`  //
	Description string                                                              `json:"description,omitempty"` // view description
	FieldGroups *[]ResponseReportsGetViewDetailsForAGivenViewGroupViewV1FieldGroups `json:"fieldGroups,omitempty"` //
	Filters     *[]ResponseReportsGetViewDetailsForAGivenViewGroupViewV1Filters     `json:"filters,omitempty"`     //
	Formats     *[]ResponseReportsGetViewDetailsForAGivenViewGroupViewV1Formats     `json:"formats,omitempty"`     //
	Schedules   *[]ResponseReportsGetViewDetailsForAGivenViewGroupViewV1Schedules   `json:"schedules,omitempty"`   //
	ViewID      string                                                              `json:"viewId,omitempty"`      // Unique view Id
	ViewInfo    string                                                              `json:"viewInfo,omitempty"`    // view filters info
	ViewName    string                                                              `json:"viewName,omitempty"`    // view name
}
type ResponseReportsGetViewDetailsForAGivenViewGroupViewV1Deliveries struct {
	Type    string `json:"type,omitempty"`    // delivery type
	Default *bool  `json:"default,omitempty"` // true, if the delivery type is considered default
}
type ResponseReportsGetViewDetailsForAGivenViewGroupViewV1FieldGroups struct {
	FieldGroupDisplayName string                                                                    `json:"fieldGroupDisplayName,omitempty"` // Field group label/displayname for user
	FieldGroupName        string                                                                    `json:"fieldGroupName,omitempty"`        // Field group name
	Fields                *[]ResponseReportsGetViewDetailsForAGivenViewGroupViewV1FieldGroupsFields `json:"fields,omitempty"`                //
	TableID               string                                                                    `json:"tableId,omitempty"`               // Table Id of the corresponding table mapped to fieldgroup
}
type ResponseReportsGetViewDetailsForAGivenViewGroupViewV1FieldGroupsFields struct {
	DisplayName string `json:"displayName,omitempty"` // field label/displayname
	Name        string `json:"name,omitempty"`        // field name
}
type ResponseReportsGetViewDetailsForAGivenViewGroupViewV1Filters struct {
	AdditionalInfo *ResponseReportsGetViewDetailsForAGivenViewGroupViewV1FiltersAdditionalInfo `json:"additionalInfo,omitempty"` // Additional info for managing filter options
	CacheFilter    *bool                                                                       `json:"cacheFilter,omitempty"`    //
	DataType       string                                                                      `json:"dataType,omitempty"`       // data type of filter value
	DisplayName    string                                                                      `json:"displayName,omitempty"`    // filter label/displayname
	FilterSource   *ResponseReportsGetViewDetailsForAGivenViewGroupViewV1FiltersFilterSource   `json:"filterSource,omitempty"`   //
	Name           string                                                                      `json:"name,omitempty"`           // filter name
	Required       *bool                                                                       `json:"required,omitempty"`       // true if the filter is required
	TimeOptions    *[]ResponseReportsGetViewDetailsForAGivenViewGroupViewV1FiltersTimeOptions  `json:"timeOptions,omitempty"`    //
	Type           string                                                                      `json:"type,omitempty"`           // filter type. Used to handle filter value selection by the client for report configuration.
}
type ResponseReportsGetViewDetailsForAGivenViewGroupViewV1FiltersAdditionalInfo interface{}
type ResponseReportsGetViewDetailsForAGivenViewGroupViewV1FiltersFilterSource struct {
	DataSource       *ResponseReportsGetViewDetailsForAGivenViewGroupViewV1FiltersFilterSourceDataSource `json:"dataSource,omitempty"`       //
	DisplayValuePath string                                                                              `json:"displayValuePath,omitempty"` // JSONPath of the label of filter option from the filter option as root
	RootPath         string                                                                              `json:"rootPath,omitempty"`         // JSONPath of the filter options array in the API response
	ValuePath        string                                                                              `json:"valuePath,omitempty"`        // JSONPath of the value of filter option from the filter option as root
}
type ResponseReportsGetViewDetailsForAGivenViewGroupViewV1FiltersFilterSourceDataSource interface{}
type ResponseReportsGetViewDetailsForAGivenViewGroupViewV1FiltersTimeOptions struct {
	Info     string `json:"info,omitempty"`     // Time range option description
	MaxValue *int   `json:"maxValue,omitempty"` // Maximum number of hours allowed for the time range option. (Client Validation)
	MinValue *int   `json:"minValue,omitempty"` // Minimum number of hours allowed for the time range option. (Client Validation)
	Name     string `json:"name,omitempty"`     // Time range option label
	Value    string `json:"value,omitempty"`    // Time range option value
}
type ResponseReportsGetViewDetailsForAGivenViewGroupViewV1Formats struct {
	Format   string                                                                `json:"format,omitempty"`   // format type
	Name     string                                                                `json:"name,omitempty"`     // format name
	Default  *bool                                                                 `json:"default,omitempty"`  // true, if the format type is considered default
	Template *ResponseReportsGetViewDetailsForAGivenViewGroupViewV1FormatsTemplate `json:"template,omitempty"` //
}
type ResponseReportsGetViewDetailsForAGivenViewGroupViewV1FormatsTemplate struct {
	JsTemplateID string `json:"jsTemplateId,omitempty"` // TemplateId of template
}
type ResponseReportsGetViewDetailsForAGivenViewGroupViewV1Schedules struct {
	Type    string `json:"type,omitempty"`    // schedule type
	Default *bool  `json:"default,omitempty"` // true, if the schedule type is default
}
type RequestReportsUpdateScheduleOfFlexibleReportV1 struct {
	Schedule *RequestReportsUpdateScheduleOfFlexibleReportV1Schedule `json:"schedule,omitempty"` // Schedule information
}
type RequestReportsUpdateScheduleOfFlexibleReportV1Schedule interface{}
type RequestReportsCreateOrScheduleAReportV1 struct {
	Tags             []string                                             `json:"tags,omitempty"`             // array of tags for report
	Deliveries       *[]RequestReportsCreateOrScheduleAReportV1Deliveries `json:"deliveries,omitempty"`       // Array of available delivery channels
	Name             string                                               `json:"name,omitempty"`             // report name
	Schedule         *RequestReportsCreateOrScheduleAReportV1Schedule     `json:"schedule,omitempty"`         //
	View             *RequestReportsCreateOrScheduleAReportV1View         `json:"view,omitempty"`             //
	ViewGroupID      string                                               `json:"viewGroupId,omitempty"`      // viewGroupId of the viewgroup for the report
	ViewGroupVersion string                                               `json:"viewGroupVersion,omitempty"` // version of viewgroup for the report
	DataCategory     string                                               `json:"dataCategory,omitempty"`     // category of viewgroup for the report
}
type RequestReportsCreateOrScheduleAReportV1Deliveries interface{}
type RequestReportsCreateOrScheduleAReportV1Schedule interface{}
type RequestReportsCreateOrScheduleAReportV1View struct {
	FieldGroups *[]RequestReportsCreateOrScheduleAReportV1ViewFieldGroups `json:"fieldGroups,omitempty"` //
	Filters     *[]RequestReportsCreateOrScheduleAReportV1ViewFilters     `json:"filters,omitempty"`     //
	Format      *RequestReportsCreateOrScheduleAReportV1ViewFormat        `json:"format,omitempty"`      //
	Name        string                                                    `json:"name,omitempty"`        // view name
	ViewID      string                                                    `json:"viewId,omitempty"`      // view Id
}
type RequestReportsCreateOrScheduleAReportV1ViewFieldGroups struct {
	FieldGroupDisplayName string                                                          `json:"fieldGroupDisplayName,omitempty"` // Field group label/displayname for user
	FieldGroupName        string                                                          `json:"fieldGroupName,omitempty"`        // Field group name
	Fields                *[]RequestReportsCreateOrScheduleAReportV1ViewFieldGroupsFields `json:"fields,omitempty"`                //
}
type RequestReportsCreateOrScheduleAReportV1ViewFieldGroupsFields struct {
	DisplayName string `json:"displayName,omitempty"` // field label/displayname
	Name        string `json:"name,omitempty"`        // field name
}
type RequestReportsCreateOrScheduleAReportV1ViewFilters struct {
	DisplayName string                                                   `json:"displayName,omitempty"` // filter label/displayname
	Name        string                                                   `json:"name,omitempty"`        // filter name
	Type        string                                                   `json:"type,omitempty"`        // filter type
	Value       *RequestReportsCreateOrScheduleAReportV1ViewFiltersValue `json:"value,omitempty"`       // value of filter. data type is based on the filter type. Use the filter definitions from the view to fetch the options for a filter.
}
type RequestReportsCreateOrScheduleAReportV1ViewFiltersValue interface{}
type RequestReportsCreateOrScheduleAReportV1ViewFormat struct {
	FormatType string `json:"formatType,omitempty"` // format type of report
	Name       string `json:"name,omitempty"`       // format name of report
}

//DownloadFlexibleReportV1 Download Flexible Report - a1bc-fba5-4c1b-849d
/* This is used to download the flexible report. The API returns report content. Save the response to a file by converting the response data as a blob and setting the file format available from content-disposition response header.


@param reportID reportId path parameter. Id of the report

@param executionID executionId path parameter. Id of execution


Documentation Link: https://developer.cisco.com/docs/dna-center/#!download-flexible-report-v1
*/
func (s *ReportsService) DownloadFlexibleReportV1(reportID string, executionID string) (*resty.Response, error) {
	path := "/dna/data/api/v1/flexible-report/report/content/{reportId}/{executionId}"
	path = strings.Replace(path, "{reportId}", fmt.Sprintf("%v", reportID), -1)
	path = strings.Replace(path, "{executionId}", fmt.Sprintf("%v", executionID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DownloadFlexibleReportV1(reportID, executionID)
		}
		return response, fmt.Errorf("error with operation DownloadFlexibleReportV1")
	}

	return response, err

}

//GetExecutionIDByReportIDV1 Get Execution Id by Report Id - 3e91-6aa5-4369-a739
/* Get Execution Id by Report Id


@param reportID reportId path parameter. Id of the report


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-execution-id-by-report-id-v1
*/
func (s *ReportsService) GetExecutionIDByReportIDV1(reportID string) (*ResponseReportsGetExecutionIDByReportIDV1, *resty.Response, error) {
	path := "/dna/data/api/v1/flexible-report/report/{reportId}/executions"
	path = strings.Replace(path, "{reportId}", fmt.Sprintf("%v", reportID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseReportsGetExecutionIDByReportIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetExecutionIDByReportIDV1(reportID)
		}
		return nil, response, fmt.Errorf("error with operation GetExecutionIdByReportIdV1")
	}

	result := response.Result().(*ResponseReportsGetExecutionIDByReportIDV1)
	return result, response, err

}

//GetFlexibleReportScheduleByReportIDV1 Get flexible report schedule by report id - 2a91-ebd9-4949-a73c
/* Get flexible report schedule by report id


@param reportID reportId path parameter. Id of the report


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-flexible-report-schedule-by-report-id-v1
*/
func (s *ReportsService) GetFlexibleReportScheduleByReportIDV1(reportID string) (*ResponseReportsGetFlexibleReportScheduleByReportIDV1, *resty.Response, error) {
	path := "/dna/data/api/v1/flexible-report/schedule/{reportId}"
	path = strings.Replace(path, "{reportId}", fmt.Sprintf("%v", reportID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseReportsGetFlexibleReportScheduleByReportIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFlexibleReportScheduleByReportIDV1(reportID)
		}
		return nil, response, fmt.Errorf("error with operation GetFlexibleReportScheduleByReportIdV1")
	}

	result := response.Result().(*ResponseReportsGetFlexibleReportScheduleByReportIDV1)
	return result, response, err

}

//GetAllFlexibleReportSchedulesV1 Get all flexible report schedules - 7fa5-299b-4d49-8d6f
/* Get all flexible report schedules



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-all-flexible-report-schedules-v1
*/
func (s *ReportsService) GetAllFlexibleReportSchedulesV1() (*ResponseReportsGetAllFlexibleReportSchedulesV1, *resty.Response, error) {
	path := "/dna/data/api/v1/flexible-report/schedules"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseReportsGetAllFlexibleReportSchedulesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAllFlexibleReportSchedulesV1()
		}
		return nil, response, fmt.Errorf("error with operation GetAllFlexibleReportSchedulesV1")
	}

	result := response.Result().(*ResponseReportsGetAllFlexibleReportSchedulesV1)
	return result, response, err

}

//GetListOfScheduledReportsV1 Get list of scheduled reports - 2ab4-b80d-49ca-ae42
/* Get list of scheduled report configurations.


@param GetListOfScheduledReportsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-list-of-scheduled-reports-v1
*/
func (s *ReportsService) GetListOfScheduledReportsV1(GetListOfScheduledReportsV1QueryParams *GetListOfScheduledReportsV1QueryParams) (*ResponseReportsGetListOfScheduledReportsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/data/reports"

	queryString, _ := query.Values(GetListOfScheduledReportsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseReportsGetListOfScheduledReportsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetListOfScheduledReportsV1(GetListOfScheduledReportsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetListOfScheduledReportsV1")
	}

	result := response.Result().(*ResponseReportsGetListOfScheduledReportsV1)
	return result, response, err

}

//GetAScheduledReportV1 Get a scheduled report - b79a-3910-4e18-9251
/* Get scheduled report configuration by reportId


@param reportID reportId path parameter. reportId of report


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-a-scheduled-report-v1
*/
func (s *ReportsService) GetAScheduledReportV1(reportID string) (*ResponseReportsGetAScheduledReportV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/data/reports/{reportId}"
	path = strings.Replace(path, "{reportId}", fmt.Sprintf("%v", reportID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseReportsGetAScheduledReportV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAScheduledReportV1(reportID)
		}
		return nil, response, fmt.Errorf("error with operation GetAScheduledReportV1")
	}

	result := response.Result().(*ResponseReportsGetAScheduledReportV1)
	return result, response, err

}

//GetAllExecutionDetailsForAGivenReportV1 Get all execution details for a given report - 91b9-d830-4679-a273
/* Get details of all executions for a given report


@param reportID reportId path parameter. reportId of report


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-all-execution-details-for-a-given-report-v1
*/
func (s *ReportsService) GetAllExecutionDetailsForAGivenReportV1(reportID string) (*ResponseReportsGetAllExecutionDetailsForAGivenReportV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/data/reports/{reportId}/executions"
	path = strings.Replace(path, "{reportId}", fmt.Sprintf("%v", reportID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseReportsGetAllExecutionDetailsForAGivenReportV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAllExecutionDetailsForAGivenReportV1(reportID)
		}
		return nil, response, fmt.Errorf("error with operation GetAllExecutionDetailsForAGivenReportV1")
	}

	result := response.Result().(*ResponseReportsGetAllExecutionDetailsForAGivenReportV1)
	return result, response, err

}

//DownloadReportContentV1 Download report content - d6bb-ebd7-4a48-87bd
/* Returns report content. Save the response to a file by converting the response data as a blob and setting the file format available from content-disposition response header.


@param reportID reportId path parameter. reportId of report

@param executionID executionId path parameter. executionId of report execution


Documentation Link: https://developer.cisco.com/docs/dna-center/#!download-report-content-v1
*/
func (s *ReportsService) DownloadReportContentV1(reportID string, executionID string) (FileDownload, *resty.Response, error) {
	path := "/dna/intent/api/v1/data/reports/{reportId}/executions/{executionId}"
	path = strings.Replace(path, "{reportId}", fmt.Sprintf("%v", reportID), -1)
	path = strings.Replace(path, "{executionId}", fmt.Sprintf("%v", executionID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Get(path)

	fdownload := FileDownload{}
	if err != nil {
		return fdownload, nil, err
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DownloadReportContentV1(reportID, executionID)
		}
		return fdownload, response, fmt.Errorf("error with operation ExportTrustedCertificate")
	}

	fdownload.FileData = response.Body()
	headerVal := response.Header()["Content-Disposition"][0]
	fname := strings.SplitAfter(headerVal, "=")
	fdownload.FileName = strings.ReplaceAll(fname[1], "\"", "")

	return fdownload, response, err

}

//GetAllViewGroupsV1 Get all view groups - 2f90-4a35-44ab-b1c9
/* Gives a list of summary of all view groups.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-all-view-groups-v1
*/
func (s *ReportsService) GetAllViewGroupsV1() (*ResponseReportsGetAllViewGroupsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/data/view-groups"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseReportsGetAllViewGroupsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAllViewGroupsV1()
		}
		return nil, response, fmt.Errorf("error with operation GetAllViewGroupsV1")
	}

	result := response.Result().(*ResponseReportsGetAllViewGroupsV1)
	return result, response, err

}

//GetViewsForAGivenViewGroupV1 Get views for a given view group - 03b6-aa2b-4dda-a555
/* Gives a list of summary of all views in a viewgroup. Use "Get all view groups" API to get the viewGroupIds (required as a query param for this API) for available viewgroups.


@param viewGroupID viewGroupId path parameter. viewGroupId of viewgroup.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-views-for-a-given-view-group-v1
*/
func (s *ReportsService) GetViewsForAGivenViewGroupV1(viewGroupID string) (*ResponseReportsGetViewsForAGivenViewGroupV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/data/view-groups/{viewGroupId}"
	path = strings.Replace(path, "{viewGroupId}", fmt.Sprintf("%v", viewGroupID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseReportsGetViewsForAGivenViewGroupV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetViewsForAGivenViewGroupV1(viewGroupID)
		}
		return nil, response, fmt.Errorf("error with operation GetViewsForAGivenViewGroupV1")
	}

	result := response.Result().(*ResponseReportsGetViewsForAGivenViewGroupV1)
	return result, response, err

}

//GetViewDetailsForAGivenViewGroupViewV1 Get view details for a given view group & view - 1d9a-ba2f-4f89-ae51
/* Gives complete information of the view that is required to configure a report. Use "Get views for a given view group" API to get the viewIds  (required as a query param for this API) for available views.


@param viewGroupID viewGroupId path parameter. viewGroupId of viewgroup

@param viewID viewId path parameter. view id of view


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-view-details-for-a-given-view-group-view-v1
*/
func (s *ReportsService) GetViewDetailsForAGivenViewGroupViewV1(viewGroupID string, viewID string) (*ResponseReportsGetViewDetailsForAGivenViewGroupViewV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/data/view-groups/{viewGroupId}/views/{viewId}"
	path = strings.Replace(path, "{viewGroupId}", fmt.Sprintf("%v", viewGroupID), -1)
	path = strings.Replace(path, "{viewId}", fmt.Sprintf("%v", viewID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseReportsGetViewDetailsForAGivenViewGroupViewV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetViewDetailsForAGivenViewGroupViewV1(viewGroupID, viewID)
		}
		return nil, response, fmt.Errorf("error with operation GetViewDetailsForAGivenViewGroupViewV1")
	}

	result := response.Result().(*ResponseReportsGetViewDetailsForAGivenViewGroupViewV1)
	return result, response, err

}

//ExecutingTheFlexibleReportV1 Executing the Flexible report - 4886-9a2c-4c5a-b570
/* This API is used for executing the report


@param reportID reportId path parameter. Id of the Report


Documentation Link: https://developer.cisco.com/docs/dna-center/#!executing-the-flexible-report-v1
*/
func (s *ReportsService) ExecutingTheFlexibleReportV1(reportID string) (*ResponseReportsExecutingTheFlexibleReportV1, *resty.Response, error) {
	path := "/dna/data/api/v1/flexible-report/report/{reportId}/execute"
	path = strings.Replace(path, "{reportId}", fmt.Sprintf("%v", reportID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseReportsExecutingTheFlexibleReportV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ExecutingTheFlexibleReportV1(reportID)
		}

		return nil, response, fmt.Errorf("error with operation ExecutingTheFlexibleReportV1")
	}

	result := response.Result().(*ResponseReportsExecutingTheFlexibleReportV1)
	return result, response, err

}

//CreateOrScheduleAReportV1 Create or Schedule a report - 8abf-291a-42aa-8860
/* Create/Schedule a report configuration. Use "Get view details for a given view group & view" API to get the metadata required to configure a report.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-or-schedule-a-report-v1
*/
func (s *ReportsService) CreateOrScheduleAReportV1(requestReportsCreateOrScheduleAReportV1 *RequestReportsCreateOrScheduleAReportV1) (*ResponseReportsCreateOrScheduleAReportV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/data/reports"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestReportsCreateOrScheduleAReportV1).
		SetResult(&ResponseReportsCreateOrScheduleAReportV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateOrScheduleAReportV1(requestReportsCreateOrScheduleAReportV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateOrScheduleAReportV1")
	}

	result := response.Result().(*ResponseReportsCreateOrScheduleAReportV1)
	return result, response, err

}

//UpdateScheduleOfFlexibleReportV1 Update schedule of flexible report - 498f-2b3d-4cd8-bd9d
/* Update schedule of flexible report


@param reportID reportId path parameter. Id of the report

*/
func (s *ReportsService) UpdateScheduleOfFlexibleReportV1(reportID string, requestReportsUpdateScheduleOfFlexibleReportV1 *RequestReportsUpdateScheduleOfFlexibleReportV1) (*ResponseReportsUpdateScheduleOfFlexibleReportV1, *resty.Response, error) {
	path := "/dna/data/api/v1/flexible-report/schedule/{reportId}"
	path = strings.Replace(path, "{reportId}", fmt.Sprintf("%v", reportID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestReportsUpdateScheduleOfFlexibleReportV1).
		SetResult(&ResponseReportsUpdateScheduleOfFlexibleReportV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateScheduleOfFlexibleReportV1(reportID, requestReportsUpdateScheduleOfFlexibleReportV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateScheduleOfFlexibleReportV1")
	}

	result := response.Result().(*ResponseReportsUpdateScheduleOfFlexibleReportV1)
	return result, response, err

}

//DeleteAScheduledReportV1 Delete a scheduled report - 239c-6921-4f9b-b12e
/* Delete a scheduled report configuration. Deletes the report executions also.


@param reportID reportId path parameter. reportId of report


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-a-scheduled-report-v1
*/
func (s *ReportsService) DeleteAScheduledReportV1(reportID string) (*ResponseReportsDeleteAScheduledReportV1, *resty.Response, error) {
	//reportID string
	path := "/dna/intent/api/v1/data/reports/{reportId}"
	path = strings.Replace(path, "{reportId}", fmt.Sprintf("%v", reportID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseReportsDeleteAScheduledReportV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteAScheduledReportV1(reportID)
		}
		return nil, response, fmt.Errorf("error with operation DeleteAScheduledReportV1")
	}

	result := response.Result().(*ResponseReportsDeleteAScheduledReportV1)
	return result, response, err

}

// Alias Function
func (s *ReportsService) CreateOrScheduleAReport(requestReportsCreateOrScheduleAReportV1 *RequestReportsCreateOrScheduleAReportV1) (*ResponseReportsCreateOrScheduleAReportV1, *resty.Response, error) {
	return s.CreateOrScheduleAReportV1(requestReportsCreateOrScheduleAReportV1)
}

// Alias Function
func (s *ReportsService) GetExecutionIDByReportID(reportID string) (*ResponseReportsGetExecutionIDByReportIDV1, *resty.Response, error) {
	return s.GetExecutionIDByReportIDV1(reportID)
}

// Alias Function
func (s *ReportsService) GetAllViewGroups() (*ResponseReportsGetAllViewGroupsV1, *resty.Response, error) {
	return s.GetAllViewGroupsV1()
}

// Alias Function
func (s *ReportsService) ExecutingTheFlexibleReport(reportID string) (*ResponseReportsExecutingTheFlexibleReportV1, *resty.Response, error) {
	return s.ExecutingTheFlexibleReportV1(reportID)
}

// Alias Function
func (s *ReportsService) GetAScheduledReport(reportID string) (*ResponseReportsGetAScheduledReportV1, *resty.Response, error) {
	return s.GetAScheduledReportV1(reportID)
}

// Alias Function
func (s *ReportsService) UpdateScheduleOfFlexibleReport(reportID string, requestReportsUpdateScheduleOfFlexibleReportV1 *RequestReportsUpdateScheduleOfFlexibleReportV1) (*ResponseReportsUpdateScheduleOfFlexibleReportV1, *resty.Response, error) {
	return s.UpdateScheduleOfFlexibleReportV1(reportID, requestReportsUpdateScheduleOfFlexibleReportV1)
}

// Alias Function
func (s *ReportsService) GetAllExecutionDetailsForAGivenReport(reportID string) (*ResponseReportsGetAllExecutionDetailsForAGivenReportV1, *resty.Response, error) {
	return s.GetAllExecutionDetailsForAGivenReportV1(reportID)
}

// Alias Function
func (s *ReportsService) GetAllFlexibleReportSchedules() (*ResponseReportsGetAllFlexibleReportSchedulesV1, *resty.Response, error) {
	return s.GetAllFlexibleReportSchedulesV1()
}

// Alias Function
func (s *ReportsService) DownloadFlexibleReport(reportID string, executionID string) (*resty.Response, error) {
	return s.DownloadFlexibleReportV1(reportID, executionID)
}

// Alias Function
func (s *ReportsService) GetListOfScheduledReports(GetListOfScheduledReportsV1QueryParams *GetListOfScheduledReportsV1QueryParams) (*ResponseReportsGetListOfScheduledReportsV1, *resty.Response, error) {
	return s.GetListOfScheduledReportsV1(GetListOfScheduledReportsV1QueryParams)
}

// Alias Function
func (s *ReportsService) DownloadReportContent(reportID string, executionID string) (FileDownload, *resty.Response, error) {
	return s.DownloadReportContentV1(reportID, executionID)
}

// Alias Function
func (s *ReportsService) GetViewsForAGivenViewGroup(viewGroupID string) (*ResponseReportsGetViewsForAGivenViewGroupV1, *resty.Response, error) {
	return s.GetViewsForAGivenViewGroupV1(viewGroupID)
}

// Alias Function
func (s *ReportsService) DeleteAScheduledReport(reportID string) (*ResponseReportsDeleteAScheduledReportV1, *resty.Response, error) {
	return s.DeleteAScheduledReportV1(reportID)
}

// Alias Function
func (s *ReportsService) GetFlexibleReportScheduleByReportID(reportID string) (*ResponseReportsGetFlexibleReportScheduleByReportIDV1, *resty.Response, error) {
	return s.GetFlexibleReportScheduleByReportIDV1(reportID)
}
