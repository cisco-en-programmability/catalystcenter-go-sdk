package catalyst

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type EventManagementService service

type GetAuditLogParentRecordsV1QueryParams struct {
	InstanceID     string  `url:"instanceId,omitempty"`     //InstanceID of the Audit Log.
	Name           string  `url:"name,omitempty"`           //Audit Log notification event name.
	EventID        string  `url:"eventId,omitempty"`        //Audit Log notification's event ID.
	Category       string  `url:"category,omitempty"`       //Audit Log notification's event category. Supported values: INFO, WARN, ERROR, ALERT, TASK_PROGRESS, TASK_FAILURE, TASK_COMPLETE, COMMAND, QUERY, CONVERSATION
	Severity       string  `url:"severity,omitempty"`       //Audit Log notification's event severity. Supported values: 1, 2, 3, 4, 5.
	Domain         string  `url:"domain,omitempty"`         //Audit Log notification's event domain.
	SubDomain      string  `url:"subDomain,omitempty"`      //Audit Log notification's event sub-domain.
	Source         string  `url:"source,omitempty"`         //Audit Log notification's event source.
	UserID         string  `url:"userId,omitempty"`         //Audit Log notification's event userId.
	Context        string  `url:"context,omitempty"`        //Audit Log notification's event correlationId.
	EventHierarchy string  `url:"eventHierarchy,omitempty"` //Audit Log notification's event eventHierarchy. Example: "US.CA.San Jose" OR "US.CA" OR "CA.San Jose" - Delimiter for hierarchy separation is ".".
	SiteID         string  `url:"siteId,omitempty"`         //Audit Log notification's siteId.
	DeviceID       string  `url:"deviceId,omitempty"`       //Audit Log notification's deviceId.
	IsSystemEvents bool    `url:"isSystemEvents,omitempty"` //Parameter to filter system generated audit-logs.
	Description    string  `url:"description,omitempty"`    //String full/partial search - (Provided input string is case insensitively matched for records).
	Offset         float64 `url:"offset,omitempty"`         //Position of a particular Audit Log record in the data.
	Limit          float64 `url:"limit,omitempty"`          //Number of Audit Log records to be returned per page.
	StartTime      float64 `url:"startTime,omitempty"`      //Start Time in milliseconds since Epoch Eg. 1597950637211 (when provided endTime is mandatory)
	EndTime        float64 `url:"endTime,omitempty"`        //End Time in milliseconds since Epoch Eg. 1597961437211 (when provided startTime is mandatory)
	SortBy         string  `url:"sortBy,omitempty"`         //Sort the Audit Logs by certain fields. Supported values are event notification header attributes.
	Order          string  `url:"order,omitempty"`          //Order of the sorted Audit Log records. Default value is desc by timestamp. Supported values: asc, desc.
}
type GetAuditLogSummaryV1QueryParams struct {
	ParentInstanceID string  `url:"parentInstanceId,omitempty"` //Parent Audit Log record's instanceID.
	IsParentOnly     bool    `url:"isParentOnly,omitempty"`     //Parameter to filter parent only audit-logs.
	InstanceID       string  `url:"instanceId,omitempty"`       //InstanceID of the Audit Log.
	Name             string  `url:"name,omitempty"`             //Audit Log notification event name.
	EventID          string  `url:"eventId,omitempty"`          //Audit Log notification's event ID.
	Category         string  `url:"category,omitempty"`         //Audit Log notification's event category. Supported values: INFO, WARN, ERROR, ALERT, TASK_PROGRESS, TASK_FAILURE, TASK_COMPLETE, COMMAND, QUERY, CONVERSATION
	Severity         string  `url:"severity,omitempty"`         //Audit Log notification's event severity. Supported values: 1, 2, 3, 4, 5.
	Domain           string  `url:"domain,omitempty"`           //Audit Log notification's event domain.
	SubDomain        string  `url:"subDomain,omitempty"`        //Audit Log notification's event sub-domain.
	Source           string  `url:"source,omitempty"`           //Audit Log notification's event source.
	UserID           string  `url:"userId,omitempty"`           //Audit Log notification's event userId.
	Context          string  `url:"context,omitempty"`          //Audit Log notification's event correlationId.
	EventHierarchy   string  `url:"eventHierarchy,omitempty"`   //Audit Log notification's event eventHierarchy. Example: "US.CA.San Jose" OR "US.CA" OR "CA.San Jose" - Delimiter for hierarchy separation is ".".
	SiteID           string  `url:"siteId,omitempty"`           //Audit Log notification's siteId.
	DeviceID         string  `url:"deviceId,omitempty"`         //Audit Log notification's deviceId.
	IsSystemEvents   bool    `url:"isSystemEvents,omitempty"`   //Parameter to filter system generated audit-logs.
	Description      string  `url:"description,omitempty"`      //String full/partial search - (Provided input string is case insensitively matched for records).
	StartTime        float64 `url:"startTime,omitempty"`        //Start Time in milliseconds since Epoch Eg. 1597950637211 (when provided endTime is mandatory)
	EndTime          float64 `url:"endTime,omitempty"`          //End Time in milliseconds since Epoch Eg. 1597961437211 (when provided startTime is mandatory)
}
type GetAuditLogRecordsV1QueryParams struct {
	ParentInstanceID string  `url:"parentInstanceId,omitempty"` //Parent Audit Log record's instanceID.
	InstanceID       string  `url:"instanceId,omitempty"`       //InstanceID of the Audit Log.
	Name             string  `url:"name,omitempty"`             //Audit Log notification event name.
	EventID          string  `url:"eventId,omitempty"`          //Audit Log notification's event ID.
	Category         string  `url:"category,omitempty"`         //Audit Log notification's event category. Supported values: INFO, WARN, ERROR, ALERT, TASK_PROGRESS, TASK_FAILURE, TASK_COMPLETE, COMMAND, QUERY, CONVERSATION
	Severity         string  `url:"severity,omitempty"`         //Audit Log notification's event severity. Supported values: 1, 2, 3, 4, 5.
	Domain           string  `url:"domain,omitempty"`           //Audit Log notification's event domain.
	SubDomain        string  `url:"subDomain,omitempty"`        //Audit Log notification's event sub-domain.
	Source           string  `url:"source,omitempty"`           //Audit Log notification's event source.
	UserID           string  `url:"userId,omitempty"`           //Audit Log notification's event userId.
	Context          string  `url:"context,omitempty"`          //Audit Log notification's event correlationId.
	EventHierarchy   string  `url:"eventHierarchy,omitempty"`   //Audit Log notification's event eventHierarchy. Example: "US.CA.San Jose" OR "US.CA" OR "CA.San Jose" - Delimiter for hierarchy separation is ".".
	SiteID           string  `url:"siteId,omitempty"`           //Audit Log notification's siteId.
	DeviceID         string  `url:"deviceId,omitempty"`         //Audit Log notification's deviceId.
	IsSystemEvents   bool    `url:"isSystemEvents,omitempty"`   //Parameter to filter system generated audit-logs.
	Description      string  `url:"description,omitempty"`      //String full/partial search - (Provided input string is case insensitively matched for records).
	Offset           float64 `url:"offset,omitempty"`           //Position of a particular Audit Log record in the data.
	Limit            float64 `url:"limit,omitempty"`            //Number of Audit Log records to be returned per page.
	StartTime        float64 `url:"startTime,omitempty"`        //Start Time in milliseconds since Epoch Eg. 1597950637211 (when provided endTime is mandatory)
	EndTime          float64 `url:"endTime,omitempty"`          //End Time in milliseconds since Epoch Eg. 1597961437211 (when provided startTime is mandatory)
	SortBy           string  `url:"sortBy,omitempty"`           //Sort the Audit Logs by certain fields. Supported values are event notification header attributes.
	Order            string  `url:"order,omitempty"`            //Order of the sorted Audit Log records. Default value is desc by timestamp. Supported values: asc, desc.
}
type GetSNMPDestinationV1QueryParams struct {
	ConfigID string  `url:"configId,omitempty"` //List of SNMP configurations
	Offset   float64 `url:"offset,omitempty"`   //The number of SNMP configuration's to offset in the resultset whose default value 0
	Limit    float64 `url:"limit,omitempty"`    //The number of SNMP configuration's to limit in the resultset whose default value 10
	SortBy   string  `url:"sortBy,omitempty"`   //SortBy field name
	Order    string  `url:"order,omitempty"`    //order(asc/desc)
}
type GetNotificationsV1QueryParams struct {
	EventIDs  string  `url:"eventIds,omitempty"`  //The registered EventId should be provided
	StartTime float64 `url:"startTime,omitempty"` //Start Time in milliseconds
	EndTime   float64 `url:"endTime,omitempty"`   //End Time in milliseconds
	Category  string  `url:"category,omitempty"`  //Category
	Type      string  `url:"type,omitempty"`      //Type
	Severity  string  `url:"severity,omitempty"`  //Severity
	Domain    string  `url:"domain,omitempty"`    //Domain
	SubDomain string  `url:"subDomain,omitempty"` //Sub Domain
	Source    string  `url:"source,omitempty"`    //Source
	Offset    float64 `url:"offset,omitempty"`    //Start Offset
	Limit     float64 `url:"limit,omitempty"`     //# of records
	SortBy    string  `url:"sortBy,omitempty"`    //Sort By column
	Order     string  `url:"order,omitempty"`     //Ascending/Descending order [asc/desc]
	Tags      string  `url:"tags,omitempty"`      //Tags
	Namespace string  `url:"namespace,omitempty"` //Namespace
	SiteID    string  `url:"siteId,omitempty"`    //Site Id
}
type CountOfNotificationsV1QueryParams struct {
	EventIDs  string  `url:"eventIds,omitempty"`  //The registered EventId should be provided
	StartTime float64 `url:"startTime,omitempty"` //Start Time in milliseconds
	EndTime   float64 `url:"endTime,omitempty"`   //End Time in milliseconds
	Category  string  `url:"category,omitempty"`  //Category
	Type      string  `url:"type,omitempty"`      //Type
	Severity  string  `url:"severity,omitempty"`  //Severity
	Domain    string  `url:"domain,omitempty"`    //Domain
	SubDomain string  `url:"subDomain,omitempty"` //Sub Domain
	Source    string  `url:"source,omitempty"`    //Source
}
type GetEventSubscriptionsV1QueryParams struct {
	EventIDs string  `url:"eventIds,omitempty"` //List of subscriptions related to the respective eventIds
	Offset   float64 `url:"offset,omitempty"`   //The number of Subscriptions's to offset in the resultset whose default value 0
	Limit    float64 `url:"limit,omitempty"`    //The number of Subscriptions's to limit in the resultset whose default value 10
	SortBy   string  `url:"sortBy,omitempty"`   //SortBy field name
	Order    string  `url:"order,omitempty"`    //order(asc/desc)
}
type DeleteEventSubscriptionsV1QueryParams struct {
	Subscriptions string `url:"subscriptions,omitempty"` //List of EventSubscriptionId's for removal
}
type GetEmailSubscriptionDetailsV1QueryParams struct {
	Name       string  `url:"name,omitempty"`       //Name of the specific configuration
	InstanceID string  `url:"instanceId,omitempty"` //Instance Id of the specific configuration
	Offset     float64 `url:"offset,omitempty"`     //The number of Email Subscription detail's to offset in the resultset whose default value 0
	Limit      float64 `url:"limit,omitempty"`      //The number of Email Subscription detail's to limit in the resultset whose default value 10
	SortBy     string  `url:"sortBy,omitempty"`     //SortBy field name
	Order      string  `url:"order,omitempty"`      //order(asc/desc)
}
type GetRestWebhookSubscriptionDetailsV1QueryParams struct {
	Name       string  `url:"name,omitempty"`       //Name of the specific configuration
	InstanceID string  `url:"instanceId,omitempty"` //Instance Id of the specific configuration
	Offset     float64 `url:"offset,omitempty"`     //The number of Rest/Webhook Subscription detail's to offset in the resultset whose default value 0
	Limit      float64 `url:"limit,omitempty"`      //The number of Rest/Webhook Subscription detail's to limit in the resultset whose default value 10
	SortBy     string  `url:"sortBy,omitempty"`     //SortBy field name
	Order      string  `url:"order,omitempty"`      //order(asc/desc)
}
type GetSyslogSubscriptionDetailsV1QueryParams struct {
	Name       string  `url:"name,omitempty"`       //Name of the specific configuration
	InstanceID string  `url:"instanceId,omitempty"` //Instance Id of the specific configuration
	Offset     float64 `url:"offset,omitempty"`     //The number of Syslog Subscription detail's to offset in the resultset whose default value 0
	Limit      float64 `url:"limit,omitempty"`      //The number of Syslog Subscription detail's to limit in the resultset whose default value 10
	SortBy     string  `url:"sortBy,omitempty"`     //SortBy field name
	Order      string  `url:"order,omitempty"`      //order(asc/desc)
}
type CountOfEventSubscriptionsV1QueryParams struct {
	EventIDs string `url:"eventIds,omitempty"` //List of subscriptions related to the respective eventIds
}
type GetEmailEventSubscriptionsV1QueryParams struct {
	EventIDs  string  `url:"eventIds,omitempty"`  //List of email subscriptions related to the respective eventIds (Comma separated event ids)
	Offset    float64 `url:"offset,omitempty"`    //The number of Subscriptions's to offset in the resultset whose default value 0
	Limit     float64 `url:"limit,omitempty"`     //The number of Subscriptions's to limit in the resultset whose default value 10
	SortBy    string  `url:"sortBy,omitempty"`    //SortBy field name
	Order     string  `url:"order,omitempty"`     //order(asc/desc)
	Domain    string  `url:"domain,omitempty"`    //List of email subscriptions related to the respective domain
	SubDomain string  `url:"subDomain,omitempty"` //List of email subscriptions related to the respective sub-domain
	Category  string  `url:"category,omitempty"`  //List of email subscriptions related to the respective category
	Type      string  `url:"type,omitempty"`      //List of email subscriptions related to the respective type
	Name      string  `url:"name,omitempty"`      //List of email subscriptions related to the respective name
}
type GetRestWebhookEventSubscriptionsV1QueryParams struct {
	EventIDs  string  `url:"eventIds,omitempty"`  //List of subscriptions related to the respective eventIds (Comma separated event ids)
	Offset    float64 `url:"offset,omitempty"`    //The number of Subscriptions's to offset in the resultset whose default value 0
	Limit     float64 `url:"limit,omitempty"`     //The number of Subscriptions's to limit in the resultset whose default value 10
	SortBy    string  `url:"sortBy,omitempty"`    //SortBy field name
	Order     string  `url:"order,omitempty"`     //order(asc/desc)
	Domain    string  `url:"domain,omitempty"`    //List of subscriptions related to the respective domain
	SubDomain string  `url:"subDomain,omitempty"` //List of subscriptions related to the respective sub-domain
	Category  string  `url:"category,omitempty"`  //List of subscriptions related to the respective category
	Type      string  `url:"type,omitempty"`      //List of subscriptions related to the respective type
	Name      string  `url:"name,omitempty"`      //List of subscriptions related to the respective name
}
type GetSyslogEventSubscriptionsV1QueryParams struct {
	EventIDs  string  `url:"eventIds,omitempty"`  //List of subscriptions related to the respective eventIds (Comma separated event ids)
	Offset    float64 `url:"offset,omitempty"`    //The number of Subscriptions's to offset in the resultset whose default value 0
	Limit     float64 `url:"limit,omitempty"`     //The number of Subscriptions's to limit in the resultset whose default value 10
	SortBy    string  `url:"sortBy,omitempty"`    //SortBy field name
	Order     string  `url:"order,omitempty"`     //order(asc/desc)
	Domain    string  `url:"domain,omitempty"`    //List of subscriptions related to the respective domain
	SubDomain string  `url:"subDomain,omitempty"` //List of subscriptions related to the respective sub-domain
	Category  string  `url:"category,omitempty"`  //List of subscriptions related to the respective category
	Type      string  `url:"type,omitempty"`      //List of subscriptions related to the respective type
	Name      string  `url:"name,omitempty"`      //List of subscriptions related to the respective name
}
type GetSyslogDestinationV1QueryParams struct {
	ConfigID string  `url:"configId,omitempty"` //Config id of syslog server
	Name     string  `url:"name,omitempty"`     //Name of syslog server
	Protocol string  `url:"protocol,omitempty"` //Protocol of syslog server
	Offset   float64 `url:"offset,omitempty"`   //The number of syslog configuration's to offset in the resultset whose default value 0
	Limit    float64 `url:"limit,omitempty"`    //The number of syslog configuration's to limit in the resultset whose default value 10
	SortBy   string  `url:"sortBy,omitempty"`   //SortBy field name
	Order    string  `url:"order,omitempty"`    //order(asc/desc)
}
type GetWebhookDestinationV1QueryParams struct {
	WebhookIDs string  `url:"webhookIds,omitempty"` //List of webhook configurations
	Offset     float64 `url:"offset,omitempty"`     //The number of webhook configuration's to offset in the resultset whose default value 0
	Limit      float64 `url:"limit,omitempty"`      //The number of webhook configuration's to limit in the resultset whose default value 10
	SortBy     string  `url:"sortBy,omitempty"`     //SortBy field name
	Order      string  `url:"order,omitempty"`      //order(asc/desc)
}
type GetEventsV1QueryParams struct {
	EventID string  `url:"eventId,omitempty"` //The registered EventId should be provided
	Tags    string  `url:"tags,omitempty"`    //The registered Tags should be provided
	Offset  float64 `url:"offset,omitempty"`  //The number of Registries to offset in the resultset whose default value 0
	Limit   float64 `url:"limit,omitempty"`   //The number of Registries to limit in the resultset whose default value 10
	SortBy  string  `url:"sortBy,omitempty"`  //SortBy field name
	Order   string  `url:"order,omitempty"`   //order(asc/desc)
}
type CountOfEventsV1QueryParams struct {
	EventID string `url:"eventId,omitempty"` //The registered EventId should be provided
	Tags    string `url:"tags,omitempty"`    //The registered Tags should be provided
}
type GetEventArtifactsV1QueryParams struct {
	EventIDs string  `url:"eventIds,omitempty"` //List of eventIds
	Tags     string  `url:"tags,omitempty"`     //Tags defined
	Offset   float64 `url:"offset,omitempty"`   //Record start offset
	Limit    float64 `url:"limit,omitempty"`    //# of records to return in result set
	SortBy   string  `url:"sortBy,omitempty"`   //Sort by field
	Order    string  `url:"order,omitempty"`    //sorting order (asc/desc)
	Search   string  `url:"search,omitempty"`   //findd matches in name, description, eventId, type, category
}

type ResponseEventManagementGetAuditLogParentRecordsV1 []ResponseItemEventManagementGetAuditLogParentRecordsV1 // Array of ResponseEventManagementGetAuditLogParentRecordsV1
type ResponseItemEventManagementGetAuditLogParentRecordsV1 struct {
	Version           string                                                                  `json:"version,omitempty"`           // Version
	InstanceID        string                                                                  `json:"instanceId,omitempty"`        // Instance Id
	EventID           string                                                                  `json:"eventId,omitempty"`           // Event Id
	Namespace         string                                                                  `json:"namespace,omitempty"`         // Namespace
	Name              string                                                                  `json:"name,omitempty"`              // Name
	Description       string                                                                  `json:"description,omitempty"`       // Description
	Type              string                                                                  `json:"type,omitempty"`              // Type
	Category          string                                                                  `json:"category,omitempty"`          // Category
	Domain            string                                                                  `json:"domain,omitempty"`            // Domain
	SubDomain         string                                                                  `json:"subDomain,omitempty"`         // Sub Domain
	Severity          *int                                                                    `json:"severity,omitempty"`          // Severity
	Source            string                                                                  `json:"source,omitempty"`            // Source
	Timestamp         *int                                                                    `json:"timestamp,omitempty"`         // Timestamp
	Tags              *[]ResponseItemEventManagementGetAuditLogParentRecordsV1Tags            `json:"tags,omitempty"`              // Tags
	Details           *ResponseItemEventManagementGetAuditLogParentRecordsV1Details           `json:"details,omitempty"`           // Details
	CiscoDnaEventLink string                                                                  `json:"ciscoDnaEventLink,omitempty"` // Cisco Dna Event Link
	Note              string                                                                  `json:"note,omitempty"`              // Note
	TntID             string                                                                  `json:"tntId,omitempty"`             // Tnt Id
	Context           string                                                                  `json:"context,omitempty"`           // Context
	UserID            string                                                                  `json:"userId,omitempty"`            // User Id
	I18N              string                                                                  `json:"i18n,omitempty"`              // I18n
	EventHierarchy    string                                                                  `json:"eventHierarchy,omitempty"`    // Event Hierarchy
	Message           string                                                                  `json:"message,omitempty"`           // Message
	MessageParams     string                                                                  `json:"messageParams,omitempty"`     // Message Params
	AdditionalDetails *ResponseItemEventManagementGetAuditLogParentRecordsV1AdditionalDetails `json:"additionalDetails,omitempty"` // Additional Details
	ParentInstanceID  string                                                                  `json:"parentInstanceId,omitempty"`  // Parent Instance Id
	Network           string                                                                  `json:"network,omitempty"`           // Network
	ChildCount        *float64                                                                `json:"childCount,omitempty"`        // Child Count
	TenantID          string                                                                  `json:"tenantId,omitempty"`          // Tenant Id
}
type ResponseItemEventManagementGetAuditLogParentRecordsV1Tags interface{}
type ResponseItemEventManagementGetAuditLogParentRecordsV1Details interface{}
type ResponseItemEventManagementGetAuditLogParentRecordsV1AdditionalDetails interface{}
type ResponseEventManagementGetAuditLogSummaryV1 []ResponseItemEventManagementGetAuditLogSummaryV1 // Array of ResponseEventManagementGetAuditLogSummaryV1
type ResponseItemEventManagementGetAuditLogSummaryV1 struct {
	Count        *int `json:"count,omitempty"`        // Count
	MaxTimestamp *int `json:"maxTimestamp,omitempty"` // Max Timestamp
	MinTimestamp *int `json:"minTimestamp,omitempty"` // Min Timestamp
}
type ResponseEventManagementGetAuditLogRecordsV1 []ResponseItemEventManagementGetAuditLogRecordsV1 // Array of ResponseEventManagementGetAuditLogRecordsV1
type ResponseItemEventManagementGetAuditLogRecordsV1 struct {
	Version           string                                                            `json:"version,omitempty"`           // Version
	InstanceID        string                                                            `json:"instanceId,omitempty"`        // Instance Id
	EventID           string                                                            `json:"eventId,omitempty"`           // Event Id
	Namespace         string                                                            `json:"namespace,omitempty"`         // Namespace
	Name              string                                                            `json:"name,omitempty"`              // Name
	Description       string                                                            `json:"description,omitempty"`       // Description
	Type              string                                                            `json:"type,omitempty"`              // Type
	Category          string                                                            `json:"category,omitempty"`          // Category
	Domain            string                                                            `json:"domain,omitempty"`            // Domain
	SubDomain         string                                                            `json:"subDomain,omitempty"`         // Sub Domain
	Severity          *int                                                              `json:"severity,omitempty"`          // Severity
	Source            string                                                            `json:"source,omitempty"`            // Source
	Timestamp         *int                                                              `json:"timestamp,omitempty"`         // Timestamp
	Tags              *[]ResponseItemEventManagementGetAuditLogRecordsV1Tags            `json:"tags,omitempty"`              // Tags
	Details           *ResponseItemEventManagementGetAuditLogRecordsV1Details           `json:"details,omitempty"`           // Details
	CiscoDnaEventLink string                                                            `json:"ciscoDnaEventLink,omitempty"` // Cisco Dna Event Link
	Note              string                                                            `json:"note,omitempty"`              // Note
	TntID             string                                                            `json:"tntId,omitempty"`             // Tnt Id
	Context           string                                                            `json:"context,omitempty"`           // Context
	UserID            string                                                            `json:"userId,omitempty"`            // User Id
	I18N              string                                                            `json:"i18n,omitempty"`              // I18n
	EventHierarchy    string                                                            `json:"eventHierarchy,omitempty"`    // Event Hierarchy
	Message           string                                                            `json:"message,omitempty"`           // Message
	MessageParams     string                                                            `json:"messageParams,omitempty"`     // Message Params
	AdditionalDetails *ResponseItemEventManagementGetAuditLogRecordsV1AdditionalDetails `json:"additionalDetails,omitempty"` // Additional Details
	ParentInstanceID  string                                                            `json:"parentInstanceId,omitempty"`  // Parent Instance Id
	Network           string                                                            `json:"network,omitempty"`           // Network
	ChildCount        *float64                                                          `json:"childCount,omitempty"`        // Child Count
	TenantID          string                                                            `json:"tenantId,omitempty"`          // Tenant Id
}
type ResponseItemEventManagementGetAuditLogRecordsV1Tags interface{}
type ResponseItemEventManagementGetAuditLogRecordsV1Details interface{}
type ResponseItemEventManagementGetAuditLogRecordsV1AdditionalDetails interface{}
type ResponseEventManagementGetSNMPDestinationV1 []ResponseItemEventManagementGetSNMPDestinationV1 // Array of ResponseEventManagementGetSNMPDestinationV1
type ResponseItemEventManagementGetSNMPDestinationV1 struct {
	Version         string `json:"version,omitempty"`         // Version
	TenantID        string `json:"tenantId,omitempty"`        // Tenant Id
	ConfigID        string `json:"configId,omitempty"`        // Config Id
	Name            string `json:"name,omitempty"`            // Name
	Description     string `json:"description,omitempty"`     // Description
	IPAddress       string `json:"ipAddress,omitempty"`       // Ip Address
	Port            *int   `json:"port,omitempty"`            // Port
	SNMPVersion     string `json:"snmpVersion,omitempty"`     // Snmp Version
	Community       string `json:"community,omitempty"`       // Community
	UserName        string `json:"userName,omitempty"`        // User Name
	SNMPMode        string `json:"snmpMode,omitempty"`        // Snmp Mode
	SNMPAuthType    string `json:"snmpAuthType,omitempty"`    // Snmp Auth Type
	AuthPassword    string `json:"authPassword,omitempty"`    // Auth Password
	SNMPPrivacyType string `json:"snmpPrivacyType,omitempty"` // Snmp Privacy Type
	PrivacyPassword string `json:"privacyPassword,omitempty"` // Privacy Password
}
type ResponseEventManagementGetStatusAPIForEventsV1 struct {
	ErrorMessage  *ResponseEventManagementGetStatusAPIForEventsV1ErrorMessage `json:"errorMessage,omitempty"`  // Error Message
	APIStatus     string                                                      `json:"apiStatus,omitempty"`     // Api Status
	StatusMessage string                                                      `json:"statusMessage,omitempty"` // Status Message
}
type ResponseEventManagementGetStatusAPIForEventsV1ErrorMessage interface{}
type ResponseEventManagementUpdateEmailDestinationV1 struct {
	StatusURI string `json:"statusUri,omitempty"` // Status Uri
}
type ResponseEventManagementGetEmailDestinationV1 []ResponseItemEventManagementGetEmailDestinationV1 // Array of ResponseEventManagementGetEmailDestinationV1
type ResponseItemEventManagementGetEmailDestinationV1 struct {
	EmailConfigID       string                                                               `json:"emailConfigId,omitempty"`       // UUID
	PrimarySmtpConfig   *ResponseItemEventManagementGetEmailDestinationV1PrimarySmtpConfig   `json:"primarySMTPConfig,omitempty"`   //
	SecondarySmtpConfig *ResponseItemEventManagementGetEmailDestinationV1SecondarySmtpConfig `json:"secondarySMTPConfig,omitempty"` //
	FromEmail           string                                                               `json:"fromEmail,omitempty"`           // From Email
	ToEmail             string                                                               `json:"toEmail,omitempty"`             // To Email
	Subject             string                                                               `json:"subject,omitempty"`             // Subject
	Version             string                                                               `json:"version,omitempty"`             // Version
	TenantID            string                                                               `json:"tenantId,omitempty"`            // Tenant Id
}
type ResponseItemEventManagementGetEmailDestinationV1PrimarySmtpConfig struct {
	HostName string `json:"hostName,omitempty"` // Host Name
	Port     string `json:"port,omitempty"`     // Port
	UserName string `json:"userName,omitempty"` // User Name
	Password string `json:"password,omitempty"` // Password
	SmtpType string `json:"smtpType,omitempty"` // smtpType
}
type ResponseItemEventManagementGetEmailDestinationV1SecondarySmtpConfig struct {
	HostName string `json:"hostName,omitempty"` // Host Name
	Port     string `json:"port,omitempty"`     // Port
	UserName string `json:"userName,omitempty"` // User Name
	Password string `json:"password,omitempty"` // Password
	SmtpType string `json:"smtpType,omitempty"` // smtpType
}
type ResponseEventManagementCreateEmailDestinationV1 struct {
	StatusURI string `json:"statusUri,omitempty"` // Status Uri
}
type ResponseEventManagementGetNotificationsV1 []ResponseItemEventManagementGetNotificationsV1 // Array of ResponseEventManagementGetNotificationsV1
type ResponseItemEventManagementGetNotificationsV1 struct {
	EventID        string                                                `json:"eventId,omitempty"`        // Event Id
	InstanceID     string                                                `json:"instanceId,omitempty"`     // Instance Id
	Namespace      string                                                `json:"namespace,omitempty"`      // Namespace
	Name           string                                                `json:"name,omitempty"`           // Name
	Description    string                                                `json:"description,omitempty"`    // Description
	Version        string                                                `json:"version,omitempty"`        // Version
	Category       string                                                `json:"category,omitempty"`       // Category
	Domain         string                                                `json:"domain,omitempty"`         // Domain
	SubDomain      string                                                `json:"subDomain,omitempty"`      // Sub Domain
	Type           string                                                `json:"type,omitempty"`           // Type
	Severity       string                                                `json:"severity,omitempty"`       // Severity
	Source         string                                                `json:"source,omitempty"`         // Source
	Timestamp      string                                                `json:"timestamp,omitempty"`      // Timestamp
	Details        string                                                `json:"details,omitempty"`        // Details
	EventHierarchy string                                                `json:"eventHierarchy,omitempty"` // Event Hierarchy
	Network        *ResponseItemEventManagementGetNotificationsV1Network `json:"network,omitempty"`        //
}
type ResponseItemEventManagementGetNotificationsV1Network struct {
	SiteID   string `json:"siteId,omitempty"`   // Site Id
	DeviceID string `json:"deviceId,omitempty"` // Device Id
}
type ResponseEventManagementCountOfNotificationsV1 struct {
	Response string `json:"response,omitempty"` // Response
}
type ResponseEventManagementCreateSNMPDestinationV1 struct {
	ErrorMessage  *ResponseEventManagementCreateSNMPDestinationV1ErrorMessage `json:"errorMessage,omitempty"`  //
	APIStatus     string                                                      `json:"apiStatus,omitempty"`     // Api Status
	StatusMessage string                                                      `json:"statusMessage,omitempty"` // Status Message
}
type ResponseEventManagementCreateSNMPDestinationV1ErrorMessage struct {
	Errors *[]ResponseEventManagementCreateSNMPDestinationV1ErrorMessageErrors `json:"errors,omitempty"` // Errors
}
type ResponseEventManagementCreateSNMPDestinationV1ErrorMessageErrors interface{}
type ResponseEventManagementUpdateSNMPDestinationV1 struct {
	ErrorMessage  *ResponseEventManagementUpdateSNMPDestinationV1ErrorMessage `json:"errorMessage,omitempty"`  //
	APIStatus     string                                                      `json:"apiStatus,omitempty"`     // Api Status
	StatusMessage string                                                      `json:"statusMessage,omitempty"` // Status Message
}
type ResponseEventManagementUpdateSNMPDestinationV1ErrorMessage struct {
	Errors []string `json:"errors,omitempty"` // Errors
}
type ResponseEventManagementGetEventSubscriptionsV1 []ResponseItemEventManagementGetEventSubscriptionsV1 // Array of ResponseEventManagementGetEventSubscriptionsV1
type ResponseItemEventManagementGetEventSubscriptionsV1 struct {
	Version               string                                                                     `json:"version,omitempty"`               // Version
	SubscriptionID        string                                                                     `json:"subscriptionId,omitempty"`        // Subscription Id
	Name                  string                                                                     `json:"name,omitempty"`                  // Name
	Description           string                                                                     `json:"description,omitempty"`           // Description
	SubscriptionEndpoints *[]ResponseItemEventManagementGetEventSubscriptionsV1SubscriptionEndpoints `json:"subscriptionEndpoints,omitempty"` //
	Filter                *ResponseItemEventManagementGetEventSubscriptionsV1Filter                  `json:"filter,omitempty"`                //
	IsPrivate             *bool                                                                      `json:"isPrivate,omitempty"`             // Is Private
	TenantID              string                                                                     `json:"tenantId,omitempty"`              // Tenant Id
}
type ResponseItemEventManagementGetEventSubscriptionsV1SubscriptionEndpoints struct {
	InstanceID          string                                                                                      `json:"instanceId,omitempty"`          // Instance Id
	SubscriptionDetails *ResponseItemEventManagementGetEventSubscriptionsV1SubscriptionEndpointsSubscriptionDetails `json:"subscriptionDetails,omitempty"` //
	ConnectorType       string                                                                                      `json:"connectorType,omitempty"`       // Connector Type
}
type ResponseItemEventManagementGetEventSubscriptionsV1SubscriptionEndpointsSubscriptionDetails struct {
	ConnectorType  string                                                                                                   `json:"connectorType,omitempty"`  // Connector Type
	InstanceID     string                                                                                                   `json:"instanceId,omitempty"`     // Instance Id
	Name           string                                                                                                   `json:"name,omitempty"`           // Name
	Description    string                                                                                                   `json:"description,omitempty"`    // Description
	URL            string                                                                                                   `json:"url,omitempty"`            // Url
	BasePath       string                                                                                                   `json:"basePath,omitempty"`       // Base Path
	Resource       string                                                                                                   `json:"resource,omitempty"`       // Resource
	Method         string                                                                                                   `json:"method,omitempty"`         // Method
	TrustCert      *bool                                                                                                    `json:"trustCert,omitempty"`      // Trust Cert
	Headers        *[]ResponseItemEventManagementGetEventSubscriptionsV1SubscriptionEndpointsSubscriptionDetailsHeaders     `json:"headers,omitempty"`        //
	QueryParams    *[]ResponseItemEventManagementGetEventSubscriptionsV1SubscriptionEndpointsSubscriptionDetailsQueryParams `json:"queryParams,omitempty"`    //
	PathParams     *[]ResponseItemEventManagementGetEventSubscriptionsV1SubscriptionEndpointsSubscriptionDetailsPathParams  `json:"pathParams,omitempty"`     //
	Body           string                                                                                                   `json:"body,omitempty"`           // Body
	ConnectTimeout *int                                                                                                     `json:"connectTimeout,omitempty"` // Connect Timeout
	ReadTimeout    *int                                                                                                     `json:"readTimeout,omitempty"`    // Read Timeout
}
type ResponseItemEventManagementGetEventSubscriptionsV1SubscriptionEndpointsSubscriptionDetailsHeaders struct {
	String string `json:"string,omitempty"` // String
}
type ResponseItemEventManagementGetEventSubscriptionsV1SubscriptionEndpointsSubscriptionDetailsQueryParams struct {
	String string `json:"string,omitempty"` // String
}
type ResponseItemEventManagementGetEventSubscriptionsV1SubscriptionEndpointsSubscriptionDetailsPathParams struct {
	String string `json:"string,omitempty"` // String
}
type ResponseItemEventManagementGetEventSubscriptionsV1Filter struct {
	EventIDs          []string                                                                     `json:"eventIds,omitempty"`          // Event Ids
	Others            []string                                                                     `json:"others,omitempty"`            // Others
	DomainsSubdomains *[]ResponseItemEventManagementGetEventSubscriptionsV1FilterDomainsSubdomains `json:"domainsSubdomains,omitempty"` //
	Types             []string                                                                     `json:"types,omitempty"`             // Types
	Categories        []string                                                                     `json:"categories,omitempty"`        // Categories
	Severities        []string                                                                     `json:"severities,omitempty"`        // Severities
	Sources           []string                                                                     `json:"sources,omitempty"`           // Sources
	SiteIDs           []string                                                                     `json:"siteIds,omitempty"`           // Site Ids
}
type ResponseItemEventManagementGetEventSubscriptionsV1FilterDomainsSubdomains struct {
	Domain     string   `json:"domain,omitempty"`     // Domain
	SubDomains []string `json:"subDomains,omitempty"` // Sub Domains
}
type ResponseEventManagementDeleteEventSubscriptionsV1 struct {
	StatusURI string `json:"statusUri,omitempty"` // Status Uri
}
type ResponseEventManagementUpdateEventSubscriptionsV1 struct {
	StatusURI string `json:"statusUri,omitempty"` // Status Uri
}
type ResponseEventManagementCreateEventSubscriptionsV1 struct {
	StatusURI string `json:"statusUri,omitempty"` // Status Uri
}
type ResponseEventManagementGetEmailSubscriptionDetailsV1 []ResponseItemEventManagementGetEmailSubscriptionDetailsV1 // Array of ResponseEventManagementGetEmailSubscriptionDetailsV1
type ResponseItemEventManagementGetEmailSubscriptionDetailsV1 struct {
	InstanceID       string   `json:"instanceId,omitempty"`       // Instance Id
	Name             string   `json:"name,omitempty"`             // Name
	Description      string   `json:"description,omitempty"`      // Description
	ConnectorType    string   `json:"connectorType,omitempty"`    // Connector Type
	FromEmailAddress string   `json:"fromEmailAddress,omitempty"` // From Email Address
	ToEmailAddresses []string `json:"toEmailAddresses,omitempty"` // To Email Addresses
	Subject          string   `json:"subject,omitempty"`          // Subject
}
type ResponseEventManagementGetRestWebhookSubscriptionDetailsV1 []ResponseItemEventManagementGetRestWebhookSubscriptionDetailsV1 // Array of ResponseEventManagementGetRestWebhookSubscriptionDetailsV1
type ResponseItemEventManagementGetRestWebhookSubscriptionDetailsV1 struct {
	InstanceID     string                                                                   `json:"instanceId,omitempty"`     // Instance Id
	Name           string                                                                   `json:"name,omitempty"`           // Name
	Description    string                                                                   `json:"description,omitempty"`    // Description
	ConnectorType  string                                                                   `json:"connectorType,omitempty"`  // Connector Type
	URL            string                                                                   `json:"url,omitempty"`            // Url
	Method         string                                                                   `json:"method,omitempty"`         // Method
	TrustCert      *bool                                                                    `json:"trustCert,omitempty"`      // Trust Cert
	Headers        *[]ResponseItemEventManagementGetRestWebhookSubscriptionDetailsV1Headers `json:"headers,omitempty"`        //
	QueryParams    []string                                                                 `json:"queryParams,omitempty"`    // Query Params
	PathParams     []string                                                                 `json:"pathParams,omitempty"`     // Path Params
	Body           string                                                                   `json:"body,omitempty"`           // Body
	ConnectTimeout *int                                                                     `json:"connectTimeout,omitempty"` // Connect Timeout
	ReadTimeout    *int                                                                     `json:"readTimeout,omitempty"`    // Read Timeout
	ServiceName    string                                                                   `json:"serviceName,omitempty"`    // Service Name
	ServicePort    string                                                                   `json:"servicePort,omitempty"`    // Service Port
	Namespace      string                                                                   `json:"namespace,omitempty"`      // Namespace
	ProxyRoute     *bool                                                                    `json:"proxyRoute,omitempty"`     // Proxy Route
}
type ResponseItemEventManagementGetRestWebhookSubscriptionDetailsV1Headers struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseEventManagementGetSyslogSubscriptionDetailsV1 []ResponseItemEventManagementGetSyslogSubscriptionDetailsV1 // Array of ResponseEventManagementGetSyslogSubscriptionDetailsV1
type ResponseItemEventManagementGetSyslogSubscriptionDetailsV1 struct {
	InstanceID    string                                                                 `json:"instanceId,omitempty"`    // Instance Id
	Name          string                                                                 `json:"name,omitempty"`          // Name
	Description   string                                                                 `json:"description,omitempty"`   // Description
	ConnectorType string                                                                 `json:"connectorType,omitempty"` // Connector Type
	SyslogConfig  *ResponseItemEventManagementGetSyslogSubscriptionDetailsV1SyslogConfig `json:"syslogConfig,omitempty"`  //
}
type ResponseItemEventManagementGetSyslogSubscriptionDetailsV1SyslogConfig struct {
	ConfigID    string `json:"configId,omitempty"`    // Config Id
	Name        string `json:"name,omitempty"`        // Name
	Description string `json:"description,omitempty"` // Description
	Host        string `json:"host,omitempty"`        // Host
	Port        string `json:"port,omitempty"`        // Port
	Protocol    string `json:"protocol,omitempty"`    // Protocol
}
type ResponseEventManagementCountOfEventSubscriptionsV1 struct {
	Response *float64 `json:"response,omitempty"` // Response
}
type ResponseEventManagementCreateEmailEventSubscriptionV1 struct {
	StatusURI string `json:"statusUri,omitempty"` // Status Uri
}
type ResponseEventManagementUpdateEmailEventSubscriptionV1 struct {
	StatusURI string `json:"statusUri,omitempty"` // Status Uri
}
type ResponseEventManagementGetEmailEventSubscriptionsV1 []ResponseItemEventManagementGetEmailEventSubscriptionsV1 // Array of ResponseEventManagementGetEmailEventSubscriptionsV1
type ResponseItemEventManagementGetEmailEventSubscriptionsV1 struct {
	Version               string                                                                          `json:"version,omitempty"`               // Version
	SubscriptionID        string                                                                          `json:"subscriptionId,omitempty"`        // Subscription Id
	Name                  string                                                                          `json:"name,omitempty"`                  // Name
	Description           string                                                                          `json:"description,omitempty"`           // Description
	SubscriptionEndpoints *[]ResponseItemEventManagementGetEmailEventSubscriptionsV1SubscriptionEndpoints `json:"subscriptionEndpoints,omitempty"` //
	Filter                *ResponseItemEventManagementGetEmailEventSubscriptionsV1Filter                  `json:"filter,omitempty"`                //
	IsPrivate             *bool                                                                           `json:"isPrivate,omitempty"`             // Is Private
	TenantID              string                                                                          `json:"tenantId,omitempty"`              // Tenant Id
}
type ResponseItemEventManagementGetEmailEventSubscriptionsV1SubscriptionEndpoints struct {
	InstanceID          string                                                                                           `json:"instanceId,omitempty"`          // Instance Id
	SubscriptionDetails *ResponseItemEventManagementGetEmailEventSubscriptionsV1SubscriptionEndpointsSubscriptionDetails `json:"subscriptionDetails,omitempty"` //
	ConnectorType       string                                                                                           `json:"connectorType,omitempty"`       // Connector Type
}
type ResponseItemEventManagementGetEmailEventSubscriptionsV1SubscriptionEndpointsSubscriptionDetails struct {
	ConnectorType    string   `json:"connectorType,omitempty"`    // Connector Type
	InstanceID       string   `json:"instanceId,omitempty"`       // Instance Id
	Name             string   `json:"name,omitempty"`             // Name
	Description      string   `json:"description,omitempty"`      // Description
	FromEmailAddress string   `json:"fromEmailAddress,omitempty"` // From Email Address
	ToEmailAddresses []string `json:"toEmailAddresses,omitempty"` // To Email Addresses
	Subject          string   `json:"subject,omitempty"`          // Subject
}
type ResponseItemEventManagementGetEmailEventSubscriptionsV1Filter struct {
	EventIDs          []string                                                                          `json:"eventIds,omitempty"`          // Event Ids
	Others            []string                                                                          `json:"others,omitempty"`            // Others
	DomainsSubdomains *[]ResponseItemEventManagementGetEmailEventSubscriptionsV1FilterDomainsSubdomains `json:"domainsSubdomains,omitempty"` //
	Types             []string                                                                          `json:"types,omitempty"`             // Types
	Categories        []string                                                                          `json:"categories,omitempty"`        // Categories
	Severities        []string                                                                          `json:"severities,omitempty"`        // Severities
	Sources           []string                                                                          `json:"sources,omitempty"`           // Sources
	SiteIDs           []string                                                                          `json:"siteIds,omitempty"`           // Site Ids
}
type ResponseItemEventManagementGetEmailEventSubscriptionsV1FilterDomainsSubdomains struct {
	Domain     string   `json:"domain,omitempty"`     // Domain
	SubDomains []string `json:"subDomains,omitempty"` // Sub Domains
}
type ResponseEventManagementCreateRestWebhookEventSubscriptionV1 struct {
	StatusURI string `json:"statusUri,omitempty"` // Status Uri
}
type ResponseEventManagementGetRestWebhookEventSubscriptionsV1 []ResponseItemEventManagementGetRestWebhookEventSubscriptionsV1 // Array of ResponseEventManagementGetRestWebhookEventSubscriptionsV1
type ResponseItemEventManagementGetRestWebhookEventSubscriptionsV1 struct {
	Version               string                                                                                `json:"version,omitempty"`               // Version
	SubscriptionID        string                                                                                `json:"subscriptionId,omitempty"`        // Subscription Id
	Name                  string                                                                                `json:"name,omitempty"`                  // Name
	Description           string                                                                                `json:"description,omitempty"`           // Description
	SubscriptionEndpoints *[]ResponseItemEventManagementGetRestWebhookEventSubscriptionsV1SubscriptionEndpoints `json:"subscriptionEndpoints,omitempty"` //
	Filter                *ResponseItemEventManagementGetRestWebhookEventSubscriptionsV1Filter                  `json:"filter,omitempty"`                //
	IsPrivate             string                                                                                `json:"isPrivate,omitempty"`             // Is Private
	TenantID              string                                                                                `json:"tenantId,omitempty"`              // Tenant Id
}
type ResponseItemEventManagementGetRestWebhookEventSubscriptionsV1SubscriptionEndpoints struct {
	InstanceID          string                                                                                                 `json:"instanceId,omitempty"`          // Instance Id
	SubscriptionDetails *ResponseItemEventManagementGetRestWebhookEventSubscriptionsV1SubscriptionEndpointsSubscriptionDetails `json:"subscriptionDetails,omitempty"` //
	ConnectorType       string                                                                                                 `json:"connectorType,omitempty"`       // Connector Type
}
type ResponseItemEventManagementGetRestWebhookEventSubscriptionsV1SubscriptionEndpointsSubscriptionDetails struct {
	ConnectorType  string                                                                                                              `json:"connectorType,omitempty"`  // Connector Type
	InstanceID     string                                                                                                              `json:"instanceId,omitempty"`     // Instance Id
	Name           string                                                                                                              `json:"name,omitempty"`           // Name
	Description    string                                                                                                              `json:"description,omitempty"`    // Description
	URL            string                                                                                                              `json:"url,omitempty"`            // Url
	BasePath       string                                                                                                              `json:"basePath,omitempty"`       // Base Path
	Resource       string                                                                                                              `json:"resource,omitempty"`       // Resource
	Method         string                                                                                                              `json:"method,omitempty"`         // Method
	TrustCert      string                                                                                                              `json:"trustCert,omitempty"`      // Trust Cert
	Headers        *[]ResponseItemEventManagementGetRestWebhookEventSubscriptionsV1SubscriptionEndpointsSubscriptionDetailsHeaders     `json:"headers,omitempty"`        //
	QueryParams    *[]ResponseItemEventManagementGetRestWebhookEventSubscriptionsV1SubscriptionEndpointsSubscriptionDetailsQueryParams `json:"queryParams,omitempty"`    //
	PathParams     *[]ResponseItemEventManagementGetRestWebhookEventSubscriptionsV1SubscriptionEndpointsSubscriptionDetailsPathParams  `json:"pathParams,omitempty"`     //
	Body           string                                                                                                              `json:"body,omitempty"`           // Body
	ConnectTimeout string                                                                                                              `json:"connectTimeout,omitempty"` // Connect Timeout
	ReadTimeout    string                                                                                                              `json:"readTimeout,omitempty"`    // Read Timeout
}
type ResponseItemEventManagementGetRestWebhookEventSubscriptionsV1SubscriptionEndpointsSubscriptionDetailsHeaders struct {
	String string `json:"string,omitempty"` // String
}
type ResponseItemEventManagementGetRestWebhookEventSubscriptionsV1SubscriptionEndpointsSubscriptionDetailsQueryParams struct {
	String string `json:"string,omitempty"` // String
}
type ResponseItemEventManagementGetRestWebhookEventSubscriptionsV1SubscriptionEndpointsSubscriptionDetailsPathParams struct {
	String string `json:"string,omitempty"` // String
}
type ResponseItemEventManagementGetRestWebhookEventSubscriptionsV1Filter struct {
	EventIDs          []string                                                                                `json:"eventIds,omitempty"`          // Event Ids
	Others            []string                                                                                `json:"others,omitempty"`            // Others
	DomainsSubdomains *[]ResponseItemEventManagementGetRestWebhookEventSubscriptionsV1FilterDomainsSubdomains `json:"domainsSubdomains,omitempty"` //
	Types             []string                                                                                `json:"types,omitempty"`             // Types
	Categories        []string                                                                                `json:"categories,omitempty"`        // Categories
	Severities        []string                                                                                `json:"severities,omitempty"`        // Severities
	Sources           []string                                                                                `json:"sources,omitempty"`           // Sources
	SiteIDs           []string                                                                                `json:"siteIds,omitempty"`           // Site Ids
}
type ResponseItemEventManagementGetRestWebhookEventSubscriptionsV1FilterDomainsSubdomains struct {
	Domain     string   `json:"domain,omitempty"`     // Domain
	SubDomains []string `json:"subDomains,omitempty"` // Sub Domains
}
type ResponseEventManagementUpdateRestWebhookEventSubscriptionV1 struct {
	StatusURI string `json:"statusUri,omitempty"` // Status Uri
}
type ResponseEventManagementUpdateSyslogEventSubscriptionV1 struct {
	StatusURI string `json:"statusUri,omitempty"` // Status Uri
}
type ResponseEventManagementCreateSyslogEventSubscriptionV1 struct {
	StatusURI string `json:"statusUri,omitempty"` // Status Uri
}
type ResponseEventManagementGetSyslogEventSubscriptionsV1 []ResponseItemEventManagementGetSyslogEventSubscriptionsV1 // Array of ResponseEventManagementGetSyslogEventSubscriptionsV1
type ResponseItemEventManagementGetSyslogEventSubscriptionsV1 struct {
	Version               string                                                                           `json:"version,omitempty"`               // Version
	SubscriptionID        string                                                                           `json:"subscriptionId,omitempty"`        // Subscription Id
	Name                  string                                                                           `json:"name,omitempty"`                  // Name
	Description           string                                                                           `json:"description,omitempty"`           // Description
	SubscriptionEndpoints *[]ResponseItemEventManagementGetSyslogEventSubscriptionsV1SubscriptionEndpoints `json:"subscriptionEndpoints,omitempty"` //
	Filter                *ResponseItemEventManagementGetSyslogEventSubscriptionsV1Filter                  `json:"filter,omitempty"`                //
	IsPrivate             *bool                                                                            `json:"isPrivate,omitempty"`             // Is Private
	TenantID              string                                                                           `json:"tenantId,omitempty"`              // Tenant Id
}
type ResponseItemEventManagementGetSyslogEventSubscriptionsV1SubscriptionEndpoints struct {
	InstanceID          string                                                                                            `json:"instanceId,omitempty"`          // Instance Id
	SubscriptionDetails *ResponseItemEventManagementGetSyslogEventSubscriptionsV1SubscriptionEndpointsSubscriptionDetails `json:"subscriptionDetails,omitempty"` //
	ConnectorType       string                                                                                            `json:"connectorType,omitempty"`       // Connector Type
}
type ResponseItemEventManagementGetSyslogEventSubscriptionsV1SubscriptionEndpointsSubscriptionDetails struct {
	ConnectorType string                                                                                                        `json:"connectorType,omitempty"` // Connector Type
	InstanceID    string                                                                                                        `json:"instanceId,omitempty"`    // Instance Id
	Name          string                                                                                                        `json:"name,omitempty"`          // Name
	Description   string                                                                                                        `json:"description,omitempty"`   // Description
	SyslogConfig  *ResponseItemEventManagementGetSyslogEventSubscriptionsV1SubscriptionEndpointsSubscriptionDetailsSyslogConfig `json:"syslogConfig,omitempty"`  //
}
type ResponseItemEventManagementGetSyslogEventSubscriptionsV1SubscriptionEndpointsSubscriptionDetailsSyslogConfig struct {
	Version     string `json:"version,omitempty"`     // Version
	TenantID    string `json:"tenantId,omitempty"`    // Tenant Id
	ConfigID    string `json:"configId,omitempty"`    // Config Id
	Name        string `json:"name,omitempty"`        // Name
	Description string `json:"description,omitempty"` // Description
	Host        string `json:"host,omitempty"`        // Host
	Port        *int   `json:"port,omitempty"`        // Port
}
type ResponseItemEventManagementGetSyslogEventSubscriptionsV1Filter struct {
	EventIDs          []string                                                                           `json:"eventIds,omitempty"`          // Event Ids
	Others            []string                                                                           `json:"others,omitempty"`            // Others
	DomainsSubdomains *[]ResponseItemEventManagementGetSyslogEventSubscriptionsV1FilterDomainsSubdomains `json:"domainsSubdomains,omitempty"` //
	Types             []string                                                                           `json:"types,omitempty"`             // Types
	Categories        []string                                                                           `json:"categories,omitempty"`        // Categories
	Severities        *[]ResponseItemEventManagementGetSyslogEventSubscriptionsV1FilterSeverities        `json:"severities,omitempty"`        // Severities
	Sources           []string                                                                           `json:"sources,omitempty"`           // Sources
	SiteIDs           []string                                                                           `json:"siteIds,omitempty"`           // Site Ids
}
type ResponseItemEventManagementGetSyslogEventSubscriptionsV1FilterDomainsSubdomains struct {
	Domain     string   `json:"domain,omitempty"`     // Domain
	SubDomains []string `json:"subDomains,omitempty"` // Sub Domains
}
type ResponseItemEventManagementGetSyslogEventSubscriptionsV1FilterSeverities interface{}
type ResponseEventManagementUpdateSyslogDestinationV1 struct {
	ErrorMessage  *ResponseEventManagementUpdateSyslogDestinationV1ErrorMessage `json:"errorMessage,omitempty"`  //
	APIStatus     string                                                        `json:"apiStatus,omitempty"`     // Api Status
	StatusMessage string                                                        `json:"statusMessage,omitempty"` // Status Message
}
type ResponseEventManagementUpdateSyslogDestinationV1ErrorMessage struct {
	Errors []string `json:"errors,omitempty"` // Errors
}
type ResponseEventManagementGetSyslogDestinationV1 struct {
	ErrorMessage  *ResponseEventManagementGetSyslogDestinationV1ErrorMessage    `json:"errorMessage,omitempty"`  //
	APIStatus     string                                                        `json:"apiStatus,omitempty"`     // Status
	StatusMessage *[]ResponseEventManagementGetSyslogDestinationV1StatusMessage `json:"statusMessage,omitempty"` //
}
type ResponseEventManagementGetSyslogDestinationV1ErrorMessage struct {
	Errors []string `json:"errors,omitempty"` // Errors
}
type ResponseEventManagementGetSyslogDestinationV1StatusMessage struct {
	Version     string `json:"version,omitempty"`     // Version
	TenantID    string `json:"tenantId,omitempty"`    // Tenant Id
	ConfigID    string `json:"configId,omitempty"`    // UUID
	Name        string `json:"name,omitempty"`        // Name
	Description string `json:"description,omitempty"` // Description
	Host        string `json:"host,omitempty"`        // Host
	Port        *int   `json:"port,omitempty"`        // Port
	Protocol    string `json:"protocol,omitempty"`    // Protocol
}
type ResponseEventManagementCreateSyslogDestinationV1 struct {
	ErrorMessage  *ResponseEventManagementCreateSyslogDestinationV1ErrorMessage `json:"errorMessage,omitempty"`  //
	APIStatus     string                                                        `json:"apiStatus,omitempty"`     // Api Status
	StatusMessage string                                                        `json:"statusMessage,omitempty"` // Status Message
}
type ResponseEventManagementCreateSyslogDestinationV1ErrorMessage struct {
	Errors []string `json:"errors,omitempty"` // Errors
}
type ResponseEventManagementCreateWebhookDestinationV1 struct {
	ErrorMessage  *ResponseEventManagementCreateWebhookDestinationV1ErrorMessage `json:"errorMessage,omitempty"`  //
	APIStatus     string                                                         `json:"apiStatus,omitempty"`     // Api Status
	StatusMessage string                                                         `json:"statusMessage,omitempty"` // Status Message
}
type ResponseEventManagementCreateWebhookDestinationV1ErrorMessage struct {
	Errors []string `json:"errors,omitempty"` // Errors
}
type ResponseEventManagementUpdateWebhookDestinationV1 struct {
	ErrorMessage  *ResponseEventManagementUpdateWebhookDestinationV1ErrorMessage `json:"errorMessage,omitempty"`  //
	APIStatus     string                                                         `json:"apiStatus,omitempty"`     // Api Status
	StatusMessage string                                                         `json:"statusMessage,omitempty"` // Status Message
}
type ResponseEventManagementUpdateWebhookDestinationV1ErrorMessage struct {
	Errors []string `json:"errors,omitempty"` // Errors
}
type ResponseEventManagementGetWebhookDestinationV1 struct {
	ErrorMessage  *ResponseEventManagementGetWebhookDestinationV1ErrorMessage    `json:"errorMessage,omitempty"`  //
	APIStatus     string                                                         `json:"apiStatus,omitempty"`     // Status
	StatusMessage *[]ResponseEventManagementGetWebhookDestinationV1StatusMessage `json:"statusMessage,omitempty"` //
}
type ResponseEventManagementGetWebhookDestinationV1ErrorMessage struct {
	Errors []string `json:"errors,omitempty"` // Errors
}
type ResponseEventManagementGetWebhookDestinationV1StatusMessage struct {
	Version      string                                                                `json:"version,omitempty"`      // Version
	TenantID     string                                                                `json:"tenantId,omitempty"`     // Tenant Id
	WebhookID    string                                                                `json:"webhookId,omitempty"`    // Webhook Id
	Name         string                                                                `json:"name,omitempty"`         // Name
	Description  string                                                                `json:"description,omitempty"`  // Description
	URL          string                                                                `json:"url,omitempty"`          // Url
	Method       string                                                                `json:"method,omitempty"`       // Method
	TrustCert    *bool                                                                 `json:"trustCert,omitempty"`    // Trust Cert
	Headers      *[]ResponseEventManagementGetWebhookDestinationV1StatusMessageHeaders `json:"headers,omitempty"`      //
	IsProxyRoute *bool                                                                 `json:"isProxyRoute,omitempty"` // Is Proxy Route
}
type ResponseEventManagementGetWebhookDestinationV1StatusMessageHeaders struct {
	Name         string `json:"name,omitempty"`         // Name
	Value        string `json:"value,omitempty"`        // Value
	DefaultValue string `json:"defaultValue,omitempty"` // Default Value
	Encrypt      *bool  `json:"encrypt,omitempty"`      // Encrypt
}
type ResponseEventManagementGetEventsV1 []ResponseItemEventManagementGetEventsV1 // Array of ResponseEventManagementGetEventsV1
type ResponseItemEventManagementGetEventsV1 struct {
	EventID           string                                         `json:"eventId,omitempty"`           // Event Id
	NameSpace         string                                         `json:"nameSpace,omitempty"`         // Name Space
	Name              string                                         `json:"name,omitempty"`              // Name
	Description       string                                         `json:"description,omitempty"`       // Description
	Version           string                                         `json:"version,omitempty"`           // Version
	Category          string                                         `json:"category,omitempty"`          // Category
	Domain            string                                         `json:"domain,omitempty"`            // Domain
	SubDomain         string                                         `json:"subDomain,omitempty"`         // Sub Domain
	Type              string                                         `json:"type,omitempty"`              // Type
	Tags              []string                                       `json:"tags,omitempty"`              // Tags
	Severity          *float64                                       `json:"severity,omitempty"`          // Severity
	Details           *ResponseItemEventManagementGetEventsV1Details `json:"details,omitempty"`           // Details
	SubscriptionTypes []string                                       `json:"subscriptionTypes,omitempty"` // Subscription Types
}
type ResponseItemEventManagementGetEventsV1Details interface{}
type ResponseEventManagementCountOfEventsV1 struct {
	Response *float64 `json:"response,omitempty"` // Response
}
type ResponseEventManagementGetEventArtifactsV1 []ResponseItemEventManagementGetEventArtifactsV1 // Array of ResponseEventManagementGetEventArtifactsV1
type ResponseItemEventManagementGetEventArtifactsV1 struct {
	Version                 string                                                          `json:"version,omitempty"`                 // Version
	ArtifactID              string                                                          `json:"artifactId,omitempty"`              // Artifact Id
	Namespace               string                                                          `json:"namespace,omitempty"`               // Namespace
	Name                    string                                                          `json:"name,omitempty"`                    // Name
	Description             string                                                          `json:"description,omitempty"`             // Description
	Domain                  string                                                          `json:"domain,omitempty"`                  // Domain
	SubDomain               string                                                          `json:"subDomain,omitempty"`               // Sub Domain
	DeprecationMessage      string                                                          `json:"deprecationMessage,omitempty"`      // Deprecation Message
	Deprecated              *bool                                                           `json:"deprecated,omitempty"`              // Deprecated
	Tags                    []string                                                        `json:"tags,omitempty"`                    // Tags
	IsTemplateEnabled       *bool                                                           `json:"isTemplateEnabled,omitempty"`       // Is Template Enabled
	CiscoDnaEventLink       string                                                          `json:"ciscoDNAEventLink,omitempty"`       // Cisco D N A Event Link
	Note                    string                                                          `json:"note,omitempty"`                    // Note
	IsPrivate               *bool                                                           `json:"isPrivate,omitempty"`               // Is Private
	EventPayload            *ResponseItemEventManagementGetEventArtifactsV1EventPayload     `json:"eventPayload,omitempty"`            //
	EventTemplates          *[]ResponseItemEventManagementGetEventArtifactsV1EventTemplates `json:"eventTemplates,omitempty"`          // Event Templates
	IsTenantAware           *bool                                                           `json:"isTenantAware,omitempty"`           // Is Tenant Aware
	SupportedConnectorTypes []string                                                        `json:"supportedConnectorTypes,omitempty"` // Supported Connector Types
	TenantID                string                                                          `json:"tenantId,omitempty"`                // Tenant Id
}
type ResponseItemEventManagementGetEventArtifactsV1EventPayload struct {
	EventID           string                                                                       `json:"eventId,omitempty"`           // Event Id
	Version           string                                                                       `json:"version,omitempty"`           // Version
	Category          string                                                                       `json:"category,omitempty"`          // Category
	Type              string                                                                       `json:"type,omitempty"`              // Type
	Source            string                                                                       `json:"source,omitempty"`            // Source
	Severity          string                                                                       `json:"severity,omitempty"`          // Severity
	Details           *ResponseItemEventManagementGetEventArtifactsV1EventPayloadDetails           `json:"details,omitempty"`           //
	AdditionalDetails *ResponseItemEventManagementGetEventArtifactsV1EventPayloadAdditionalDetails `json:"additionalDetails,omitempty"` // Additional Details
}
type ResponseItemEventManagementGetEventArtifactsV1EventPayloadDetails struct {
	DeviceIP string `json:"device_ip,omitempty"` // Device Ip
	Message  string `json:"message,omitempty"`   // Message
}
type ResponseItemEventManagementGetEventArtifactsV1EventPayloadAdditionalDetails interface{}
type ResponseItemEventManagementGetEventArtifactsV1EventTemplates interface{}
type ResponseEventManagementEventArtifactCountV1 struct {
	Response *float64 `json:"response,omitempty"` // Response
}
type ResponseEventManagementGetConnectorTypesV1 []ResponseItemEventManagementGetConnectorTypesV1 // Array of ResponseEventManagementGetConnectorTypesV1
type ResponseItemEventManagementGetConnectorTypesV1 struct {
	ConnectorType      string `json:"connectorType,omitempty"`      // Connector Type
	DisplayName        string `json:"displayName,omitempty"`        // Display Name
	IsDefaultSupported *bool  `json:"isDefaultSupported,omitempty"` // Is Default Supported
	IsCustomConnector  *bool  `json:"isCustomConnector,omitempty"`  // Is Custom Connector
}
type RequestEventManagementUpdateEmailDestinationV1 struct {
	EmailConfigID       string                                                             `json:"emailConfigId,omitempty"`       // Required only for update email configuration
	PrimarySmtpConfig   *RequestEventManagementUpdateEmailDestinationV1PrimarySmtpConfig   `json:"primarySMTPConfig,omitempty"`   //
	SecondarySmtpConfig *RequestEventManagementUpdateEmailDestinationV1SecondarySmtpConfig `json:"secondarySMTPConfig,omitempty"` //
	FromEmail           string                                                             `json:"fromEmail,omitempty"`           // From Email
	ToEmail             string                                                             `json:"toEmail,omitempty"`             // To Email
	Subject             string                                                             `json:"subject,omitempty"`             // Subject
}
type RequestEventManagementUpdateEmailDestinationV1PrimarySmtpConfig struct {
	HostName string `json:"hostName,omitempty"` // Host Name
	Port     string `json:"port,omitempty"`     // Port
	UserName string `json:"userName,omitempty"` // User Name
	Password string `json:"password,omitempty"` // Password
	SmtpType string `json:"smtpType,omitempty"` // smtpType
}
type RequestEventManagementUpdateEmailDestinationV1SecondarySmtpConfig struct {
	HostName string `json:"hostName,omitempty"` // Host Name
	Port     string `json:"port,omitempty"`     // Port
	UserName string `json:"userName,omitempty"` // User Name
	Password string `json:"password,omitempty"` // Password
	SmtpType string `json:"smtpType,omitempty"` // smtpType
}
type RequestEventManagementCreateEmailDestinationV1 struct {
	EmailConfigID       string                                                             `json:"emailConfigId,omitempty"`       // Required only for update email configuration
	PrimarySmtpConfig   *RequestEventManagementCreateEmailDestinationV1PrimarySmtpConfig   `json:"primarySMTPConfig,omitempty"`   //
	SecondarySmtpConfig *RequestEventManagementCreateEmailDestinationV1SecondarySmtpConfig `json:"secondarySMTPConfig,omitempty"` //
	FromEmail           string                                                             `json:"fromEmail,omitempty"`           // From Email
	ToEmail             string                                                             `json:"toEmail,omitempty"`             // To Email
	Subject             string                                                             `json:"subject,omitempty"`             // Subject
}
type RequestEventManagementCreateEmailDestinationV1PrimarySmtpConfig struct {
	HostName string `json:"hostName,omitempty"` // Host Name
	Port     string `json:"port,omitempty"`     // Port
	UserName string `json:"userName,omitempty"` // User Name
	Password string `json:"password,omitempty"` // Password
	SmtpType string `json:"smtpType,omitempty"` // smtpType
}
type RequestEventManagementCreateEmailDestinationV1SecondarySmtpConfig struct {
	HostName string `json:"hostName,omitempty"` // Host Name
	Port     string `json:"port,omitempty"`     // Port
	UserName string `json:"userName,omitempty"` // User Name
	Password string `json:"password,omitempty"` // Password
	SmtpType string `json:"smtpType,omitempty"` // smtpType
}
type RequestEventManagementCreateSNMPDestinationV1 struct {
	Name            string `json:"name,omitempty"`            // Name
	Description     string `json:"description,omitempty"`     // Description
	IPAddress       string `json:"ipAddress,omitempty"`       // Ip Address
	Port            string `json:"port,omitempty"`            // Port
	SNMPVersion     string `json:"snmpVersion,omitempty"`     // Snmp Version
	Community       string `json:"community,omitempty"`       // Required only if snmpVersion is V2C
	UserName        string `json:"userName,omitempty"`        // Required only if snmpVersion is V3
	SNMPMode        string `json:"snmpMode,omitempty"`        // If snmpVersion is V3 it is required and cannot be NONE
	SNMPAuthType    string `json:"snmpAuthType,omitempty"`    // Snmp Auth Type
	AuthPassword    string `json:"authPassword,omitempty"`    // Auth Password
	SNMPPrivacyType string `json:"snmpPrivacyType,omitempty"` // Snmp Privacy Type
	PrivacyPassword string `json:"privacyPassword,omitempty"` // Privacy Password
}
type RequestEventManagementUpdateSNMPDestinationV1 struct {
	ConfigID        string `json:"configId,omitempty"`        // Config Id
	Name            string `json:"name,omitempty"`            // Name
	Description     string `json:"description,omitempty"`     //
	IPAddress       string `json:"ipAddress,omitempty"`       // Ip Address
	Port            string `json:"port,omitempty"`            // Port
	SNMPVersion     string `json:"snmpVersion,omitempty"`     // Snmp Version
	Community       string `json:"community,omitempty"`       // Required only if snmpVersion is V2C
	UserName        string `json:"userName,omitempty"`        // Required only if snmpVersion is V3
	SNMPMode        string `json:"snmpMode,omitempty"`        // If snmpVersion is V3 it is required and cannot be NONE
	SNMPAuthType    string `json:"snmpAuthType,omitempty"`    // Snmp Auth Type
	AuthPassword    string `json:"authPassword,omitempty"`    // Auth Password
	SNMPPrivacyType string `json:"snmpPrivacyType,omitempty"` // Snmp Privacy Type
	PrivacyPassword string `json:"privacyPassword,omitempty"` // Privacy Password
}
type RequestEventManagementUpdateEventSubscriptionsV1 []RequestItemEventManagementUpdateEventSubscriptionsV1 // Array of RequestEventManagementUpdateEventSubscriptionsV1
type RequestItemEventManagementUpdateEventSubscriptionsV1 struct {
	SubscriptionID        string                                                                       `json:"subscriptionId,omitempty"`        // Subscription Id (Unique UUID)
	Version               string                                                                       `json:"version,omitempty"`               // Version
	Name                  string                                                                       `json:"name,omitempty"`                  // Name
	Description           string                                                                       `json:"description,omitempty"`           // Description
	SubscriptionEndpoints *[]RequestItemEventManagementUpdateEventSubscriptionsV1SubscriptionEndpoints `json:"subscriptionEndpoints,omitempty"` //
	Filter                *RequestItemEventManagementUpdateEventSubscriptionsV1Filter                  `json:"filter,omitempty"`                //
}
type RequestItemEventManagementUpdateEventSubscriptionsV1SubscriptionEndpoints struct {
	InstanceID          string                                                                                        `json:"instanceId,omitempty"`          // (From 	Get Rest/Webhook Subscription Details --> pick instanceId)
	SubscriptionDetails *RequestItemEventManagementUpdateEventSubscriptionsV1SubscriptionEndpointsSubscriptionDetails `json:"subscriptionDetails,omitempty"` //
}
type RequestItemEventManagementUpdateEventSubscriptionsV1SubscriptionEndpointsSubscriptionDetails struct {
	ConnectorType string `json:"connectorType,omitempty"` // Connector Type (Must be REST)
}
type RequestItemEventManagementUpdateEventSubscriptionsV1Filter struct {
	EventIDs          []string                                                                       `json:"eventIds,omitempty"`          // Event Ids (Comma separated event ids)
	DomainsSubdomains *[]RequestItemEventManagementUpdateEventSubscriptionsV1FilterDomainsSubdomains `json:"domainsSubdomains,omitempty"` //
	Types             []string                                                                       `json:"types,omitempty"`             // Types
	Categories        []string                                                                       `json:"categories,omitempty"`        // Categories
	Severities        []string                                                                       `json:"severities,omitempty"`        // Severities
	Sources           []string                                                                       `json:"sources,omitempty"`           // Sources
	SiteIDs           []string                                                                       `json:"siteIds,omitempty"`           // Site Ids
}
type RequestItemEventManagementUpdateEventSubscriptionsV1FilterDomainsSubdomains struct {
	Domain     string   `json:"domain,omitempty"`     // Domain
	SubDomains []string `json:"subDomains,omitempty"` // Sub Domains
}
type RequestEventManagementCreateEventSubscriptionsV1 []RequestItemEventManagementCreateEventSubscriptionsV1 // Array of RequestEventManagementCreateEventSubscriptionsV1
type RequestItemEventManagementCreateEventSubscriptionsV1 struct {
	SubscriptionID        string                                                                       `json:"subscriptionId,omitempty"`        // Subscription Id (Unique UUID)
	Version               string                                                                       `json:"version,omitempty"`               // Version
	Name                  string                                                                       `json:"name,omitempty"`                  // Name
	Description           string                                                                       `json:"description,omitempty"`           // Description
	SubscriptionEndpoints *[]RequestItemEventManagementCreateEventSubscriptionsV1SubscriptionEndpoints `json:"subscriptionEndpoints,omitempty"` //
	Filter                *RequestItemEventManagementCreateEventSubscriptionsV1Filter                  `json:"filter,omitempty"`                //
}
type RequestItemEventManagementCreateEventSubscriptionsV1SubscriptionEndpoints struct {
	InstanceID          string                                                                                        `json:"instanceId,omitempty"`          // (From 	Get Rest/Webhook Subscription Details --> pick instanceId)
	SubscriptionDetails *RequestItemEventManagementCreateEventSubscriptionsV1SubscriptionEndpointsSubscriptionDetails `json:"subscriptionDetails,omitempty"` //
}
type RequestItemEventManagementCreateEventSubscriptionsV1SubscriptionEndpointsSubscriptionDetails struct {
	ConnectorType string `json:"connectorType,omitempty"` // Connector Type (Must be REST)
}
type RequestItemEventManagementCreateEventSubscriptionsV1Filter struct {
	EventIDs          []string                                                                       `json:"eventIds,omitempty"`          // Event Ids (Comma separated event ids)
	DomainsSubdomains *[]RequestItemEventManagementCreateEventSubscriptionsV1FilterDomainsSubdomains `json:"domainsSubdomains,omitempty"` //
	Types             []string                                                                       `json:"types,omitempty"`             // Types
	Categories        []string                                                                       `json:"categories,omitempty"`        // Categories
	Severities        []string                                                                       `json:"severities,omitempty"`        // Severities
	Sources           []string                                                                       `json:"sources,omitempty"`           // Sources
	SiteIDs           []string                                                                       `json:"siteIds,omitempty"`           // Site Ids
}
type RequestItemEventManagementCreateEventSubscriptionsV1FilterDomainsSubdomains struct {
	Domain     string   `json:"domain,omitempty"`     // Domain
	SubDomains []string `json:"subDomains,omitempty"` // Sub Domains
}
type RequestEventManagementCreateEmailEventSubscriptionV1 []RequestItemEventManagementCreateEmailEventSubscriptionV1 // Array of RequestEventManagementCreateEmailEventSubscriptionV1
type RequestItemEventManagementCreateEmailEventSubscriptionV1 struct {
	SubscriptionID        string                                                                           `json:"subscriptionId,omitempty"`        // Subscription Id (Unique UUID)
	Version               string                                                                           `json:"version,omitempty"`               // Version
	Name                  string                                                                           `json:"name,omitempty"`                  // Name
	Description           string                                                                           `json:"description,omitempty"`           // Description
	SubscriptionEndpoints *[]RequestItemEventManagementCreateEmailEventSubscriptionV1SubscriptionEndpoints `json:"subscriptionEndpoints,omitempty"` //
	Filter                *RequestItemEventManagementCreateEmailEventSubscriptionV1Filter                  `json:"filter,omitempty"`                //
}
type RequestItemEventManagementCreateEmailEventSubscriptionV1SubscriptionEndpoints struct {
	InstanceID          string                                                                                            `json:"instanceId,omitempty"`          // (From Get Email Subscription Details --> pick InstanceId if available)
	SubscriptionDetails *RequestItemEventManagementCreateEmailEventSubscriptionV1SubscriptionEndpointsSubscriptionDetails `json:"subscriptionDetails,omitempty"` //
}
type RequestItemEventManagementCreateEmailEventSubscriptionV1SubscriptionEndpointsSubscriptionDetails struct {
	ConnectorType    string   `json:"connectorType,omitempty"`    // Connector Type (Must be EMAIL)
	FromEmailAddress string   `json:"fromEmailAddress,omitempty"` // Senders Email Address
	ToEmailAddresses []string `json:"toEmailAddresses,omitempty"` // Recipient's Email Addresses (Comma separated)
	Subject          string   `json:"subject,omitempty"`          // Email Subject
	Name             string   `json:"name,omitempty"`             // Name
	Description      string   `json:"description,omitempty"`      // Description
}
type RequestItemEventManagementCreateEmailEventSubscriptionV1Filter struct {
	EventIDs          []string                                                                           `json:"eventIds,omitempty"`          // Event Ids
	DomainsSubdomains *[]RequestItemEventManagementCreateEmailEventSubscriptionV1FilterDomainsSubdomains `json:"domainsSubdomains,omitempty"` //
	Types             []string                                                                           `json:"types,omitempty"`             // Types
	Categories        []string                                                                           `json:"categories,omitempty"`        // Categories
	Severities        *[]int                                                                             `json:"severities,omitempty"`        // Severities
	Sources           []string                                                                           `json:"sources,omitempty"`           // Sources
	SiteIDs           []string                                                                           `json:"siteIds,omitempty"`           // Site Ids
}
type RequestItemEventManagementCreateEmailEventSubscriptionV1FilterDomainsSubdomains struct {
	Domain     string   `json:"domain,omitempty"`     // Domain
	SubDomains []string `json:"subDomains,omitempty"` // Sub Domains
}
type RequestEventManagementUpdateEmailEventSubscriptionV1 []RequestItemEventManagementUpdateEmailEventSubscriptionV1 // Array of RequestEventManagementUpdateEmailEventSubscriptionV1
type RequestItemEventManagementUpdateEmailEventSubscriptionV1 struct {
	SubscriptionID        string                                                                           `json:"subscriptionId,omitempty"`        // Subscription Id (Unique UUID)
	Version               string                                                                           `json:"version,omitempty"`               // Version
	Name                  string                                                                           `json:"name,omitempty"`                  // Name
	Description           string                                                                           `json:"description,omitempty"`           // Description
	SubscriptionEndpoints *[]RequestItemEventManagementUpdateEmailEventSubscriptionV1SubscriptionEndpoints `json:"subscriptionEndpoints,omitempty"` //
	Filter                *RequestItemEventManagementUpdateEmailEventSubscriptionV1Filter                  `json:"filter,omitempty"`                //
}
type RequestItemEventManagementUpdateEmailEventSubscriptionV1SubscriptionEndpoints struct {
	InstanceID          string                                                                                            `json:"instanceId,omitempty"`          // (From Get Email Subscription Details --> pick InstanceId)
	SubscriptionDetails *RequestItemEventManagementUpdateEmailEventSubscriptionV1SubscriptionEndpointsSubscriptionDetails `json:"subscriptionDetails,omitempty"` //
}
type RequestItemEventManagementUpdateEmailEventSubscriptionV1SubscriptionEndpointsSubscriptionDetails struct {
	ConnectorType    string   `json:"connectorType,omitempty"`    // Connector Type (Must be EMAIL)
	FromEmailAddress string   `json:"fromEmailAddress,omitempty"` // Senders Email Address
	ToEmailAddresses []string `json:"toEmailAddresses,omitempty"` // Recipient's Email Addresses (Comma separated)
	Subject          string   `json:"subject,omitempty"`          // Email Subject
	Name             string   `json:"name,omitempty"`             // Name
	Description      string   `json:"description,omitempty"`      // Description
}
type RequestItemEventManagementUpdateEmailEventSubscriptionV1Filter struct {
	EventIDs          []string                                                                           `json:"eventIds,omitempty"`          // Event Ids
	DomainsSubdomains *[]RequestItemEventManagementUpdateEmailEventSubscriptionV1FilterDomainsSubdomains `json:"domainsSubdomains,omitempty"` //
	Types             []string                                                                           `json:"types,omitempty"`             // Types
	Categories        []string                                                                           `json:"categories,omitempty"`        // Categories
	Severities        *[]int                                                                             `json:"severities,omitempty"`        // Severities
	Sources           []string                                                                           `json:"sources,omitempty"`           // Sources
	SiteIDs           []string                                                                           `json:"siteIds,omitempty"`           // Site Ids
}
type RequestItemEventManagementUpdateEmailEventSubscriptionV1FilterDomainsSubdomains struct {
	Domain     string   `json:"domain,omitempty"`     // Domain
	SubDomains []string `json:"subDomains,omitempty"` // Sub Domains
}
type RequestEventManagementCreateRestWebhookEventSubscriptionV1 []RequestItemEventManagementCreateRestWebhookEventSubscriptionV1 // Array of RequestEventManagementCreateRestWebhookEventSubscriptionV1
type RequestItemEventManagementCreateRestWebhookEventSubscriptionV1 struct {
	SubscriptionID        string                                                                                 `json:"subscriptionId,omitempty"`        // Subscription Id (Unique UUID)
	Version               string                                                                                 `json:"version,omitempty"`               // Version
	Name                  string                                                                                 `json:"name,omitempty"`                  // Name
	Description           string                                                                                 `json:"description,omitempty"`           // Description
	SubscriptionEndpoints *[]RequestItemEventManagementCreateRestWebhookEventSubscriptionV1SubscriptionEndpoints `json:"subscriptionEndpoints,omitempty"` //
	Filter                *RequestItemEventManagementCreateRestWebhookEventSubscriptionV1Filter                  `json:"filter,omitempty"`                //
}
type RequestItemEventManagementCreateRestWebhookEventSubscriptionV1SubscriptionEndpoints struct {
	InstanceID          string                                                                                                  `json:"instanceId,omitempty"`          // (From 	Get Rest/Webhook Subscription Details --> pick instanceId)
	SubscriptionDetails *RequestItemEventManagementCreateRestWebhookEventSubscriptionV1SubscriptionEndpointsSubscriptionDetails `json:"subscriptionDetails,omitempty"` //
}
type RequestItemEventManagementCreateRestWebhookEventSubscriptionV1SubscriptionEndpointsSubscriptionDetails struct {
	ConnectorType string `json:"connectorType,omitempty"` // Connector Type (Must be REST)
}
type RequestItemEventManagementCreateRestWebhookEventSubscriptionV1Filter struct {
	EventIDs          []string                                                                                 `json:"eventIds,omitempty"`          // Event Ids (Comma separated event ids)
	DomainsSubdomains *[]RequestItemEventManagementCreateRestWebhookEventSubscriptionV1FilterDomainsSubdomains `json:"domainsSubdomains,omitempty"` //
	Types             []string                                                                                 `json:"types,omitempty"`             // Types
	Categories        []string                                                                                 `json:"categories,omitempty"`        // Categories
	Severities        []string                                                                                 `json:"severities,omitempty"`        // Severities
	Sources           []string                                                                                 `json:"sources,omitempty"`           // Sources
	SiteIDs           []string                                                                                 `json:"siteIds,omitempty"`           // Site Ids
}
type RequestItemEventManagementCreateRestWebhookEventSubscriptionV1FilterDomainsSubdomains struct {
	Domain     string   `json:"domain,omitempty"`     // Domain
	SubDomains []string `json:"subDomains,omitempty"` // Sub Domains
}
type RequestEventManagementUpdateRestWebhookEventSubscriptionV1 []RequestItemEventManagementUpdateRestWebhookEventSubscriptionV1 // Array of RequestEventManagementUpdateRestWebhookEventSubscriptionV1
type RequestItemEventManagementUpdateRestWebhookEventSubscriptionV1 struct {
	SubscriptionID        string                                                                                 `json:"subscriptionId,omitempty"`        // Subscription Id (Unique UUID)
	Version               string                                                                                 `json:"version,omitempty"`               // Version
	Name                  string                                                                                 `json:"name,omitempty"`                  // Name
	Description           string                                                                                 `json:"description,omitempty"`           // Description
	SubscriptionEndpoints *[]RequestItemEventManagementUpdateRestWebhookEventSubscriptionV1SubscriptionEndpoints `json:"subscriptionEndpoints,omitempty"` //
	Filter                *RequestItemEventManagementUpdateRestWebhookEventSubscriptionV1Filter                  `json:"filter,omitempty"`                //
}
type RequestItemEventManagementUpdateRestWebhookEventSubscriptionV1SubscriptionEndpoints struct {
	InstanceID          string                                                                                                  `json:"instanceId,omitempty"`          // (From 	Get Rest/Webhook Subscription Details --> pick instanceId)
	SubscriptionDetails *RequestItemEventManagementUpdateRestWebhookEventSubscriptionV1SubscriptionEndpointsSubscriptionDetails `json:"subscriptionDetails,omitempty"` //
}
type RequestItemEventManagementUpdateRestWebhookEventSubscriptionV1SubscriptionEndpointsSubscriptionDetails struct {
	ConnectorType string `json:"connectorType,omitempty"` // Connector Type (Must be REST)
}
type RequestItemEventManagementUpdateRestWebhookEventSubscriptionV1Filter struct {
	EventIDs          []string                                                                                 `json:"eventIds,omitempty"`          // Event Ids (Comma separated event ids)
	DomainsSubdomains *[]RequestItemEventManagementUpdateRestWebhookEventSubscriptionV1FilterDomainsSubdomains `json:"domainsSubdomains,omitempty"` //
	Types             []string                                                                                 `json:"types,omitempty"`             // Types
	Categories        []string                                                                                 `json:"categories,omitempty"`        // Categories
	Severities        []string                                                                                 `json:"severities,omitempty"`        // Severities
	Sources           []string                                                                                 `json:"sources,omitempty"`           // Sources
	SiteIDs           []string                                                                                 `json:"siteIds,omitempty"`           // Site Ids
}
type RequestItemEventManagementUpdateRestWebhookEventSubscriptionV1FilterDomainsSubdomains struct {
	Domain     string   `json:"domain,omitempty"`     // Domain
	SubDomains []string `json:"subDomains,omitempty"` // Sub Domains
}
type RequestEventManagementUpdateSyslogEventSubscriptionV1 []RequestItemEventManagementUpdateSyslogEventSubscriptionV1 // Array of RequestEventManagementUpdateSyslogEventSubscriptionV1
type RequestItemEventManagementUpdateSyslogEventSubscriptionV1 struct {
	SubscriptionID        string                                                                            `json:"subscriptionId,omitempty"`        // Subscription Id (Unique UUID)
	Version               string                                                                            `json:"version,omitempty"`               // Version
	Name                  string                                                                            `json:"name,omitempty"`                  // Name
	Description           string                                                                            `json:"description,omitempty"`           // Description
	SubscriptionEndpoints *[]RequestItemEventManagementUpdateSyslogEventSubscriptionV1SubscriptionEndpoints `json:"subscriptionEndpoints,omitempty"` //
	Filter                *RequestItemEventManagementUpdateSyslogEventSubscriptionV1Filter                  `json:"filter,omitempty"`                //
}
type RequestItemEventManagementUpdateSyslogEventSubscriptionV1SubscriptionEndpoints struct {
	InstanceID          string                                                                                             `json:"instanceId,omitempty"`          // (From Get Syslog Subscription Details --> pick instanceId)
	SubscriptionDetails *RequestItemEventManagementUpdateSyslogEventSubscriptionV1SubscriptionEndpointsSubscriptionDetails `json:"subscriptionDetails,omitempty"` //
}
type RequestItemEventManagementUpdateSyslogEventSubscriptionV1SubscriptionEndpointsSubscriptionDetails struct {
	ConnectorType string `json:"connectorType,omitempty"` // Connector Type (Must be SYSLOG)
}
type RequestItemEventManagementUpdateSyslogEventSubscriptionV1Filter struct {
	EventIDs          []string                                                                            `json:"eventIds,omitempty"`          // Event Ids (Comma separated event ids)
	DomainsSubdomains *[]RequestItemEventManagementUpdateSyslogEventSubscriptionV1FilterDomainsSubdomains `json:"domainsSubdomains,omitempty"` //
	Types             []string                                                                            `json:"types,omitempty"`             // Types
	Categories        []string                                                                            `json:"categories,omitempty"`        // Categories
	Severities        []string                                                                            `json:"severities,omitempty"`        // Severities
	Sources           []string                                                                            `json:"sources,omitempty"`           // Sources
	SiteIDs           []string                                                                            `json:"siteIds,omitempty"`           // Site Ids
}
type RequestItemEventManagementUpdateSyslogEventSubscriptionV1FilterDomainsSubdomains struct {
	Domain     string   `json:"domain,omitempty"`     // Domain
	SubDomains []string `json:"subDomains,omitempty"` // Sub Domains
}
type RequestEventManagementCreateSyslogEventSubscriptionV1 []RequestItemEventManagementCreateSyslogEventSubscriptionV1 // Array of RequestEventManagementCreateSyslogEventSubscriptionV1
type RequestItemEventManagementCreateSyslogEventSubscriptionV1 struct {
	SubscriptionID        string                                                                            `json:"subscriptionId,omitempty"`        // Subscription Id (Unique UUID)
	Version               string                                                                            `json:"version,omitempty"`               // Version
	Name                  string                                                                            `json:"name,omitempty"`                  // Name
	Description           string                                                                            `json:"description,omitempty"`           // Description
	SubscriptionEndpoints *[]RequestItemEventManagementCreateSyslogEventSubscriptionV1SubscriptionEndpoints `json:"subscriptionEndpoints,omitempty"` //
	Filter                *RequestItemEventManagementCreateSyslogEventSubscriptionV1Filter                  `json:"filter,omitempty"`                //
}
type RequestItemEventManagementCreateSyslogEventSubscriptionV1SubscriptionEndpoints struct {
	InstanceID          string                                                                                             `json:"instanceId,omitempty"`          // (From Get Syslog Subscription Details --> pick instanceId)
	SubscriptionDetails *RequestItemEventManagementCreateSyslogEventSubscriptionV1SubscriptionEndpointsSubscriptionDetails `json:"subscriptionDetails,omitempty"` //
}
type RequestItemEventManagementCreateSyslogEventSubscriptionV1SubscriptionEndpointsSubscriptionDetails struct {
	ConnectorType string `json:"connectorType,omitempty"` // Connector Type (Must be SYSLOG)
}
type RequestItemEventManagementCreateSyslogEventSubscriptionV1Filter struct {
	EventIDs          []string                                                                            `json:"eventIds,omitempty"`          // Event Ids (Comma separated event ids)
	DomainsSubdomains *[]RequestItemEventManagementCreateSyslogEventSubscriptionV1FilterDomainsSubdomains `json:"domainsSubdomains,omitempty"` //
	Types             []string                                                                            `json:"types,omitempty"`             // Types
	Categories        []string                                                                            `json:"categories,omitempty"`        // Categories
	Severities        []string                                                                            `json:"severities,omitempty"`        // Severities
	Sources           []string                                                                            `json:"sources,omitempty"`           // Sources
	SiteIDs           []string                                                                            `json:"siteIds,omitempty"`           // Site Ids
}
type RequestItemEventManagementCreateSyslogEventSubscriptionV1FilterDomainsSubdomains struct {
	Domain     string   `json:"domain,omitempty"`     // Domain
	SubDomains []string `json:"subDomains,omitempty"` // Sub Domains
}
type RequestEventManagementUpdateSyslogDestinationV1 struct {
	ConfigID    string `json:"configId,omitempty"`    // Required only for update syslog configuration
	Name        string `json:"name,omitempty"`        // Name
	Description string `json:"description,omitempty"` // Description
	Host        string `json:"host,omitempty"`        // Host
	Protocol    string `json:"protocol,omitempty"`    // Protocol
	Port        *int   `json:"port,omitempty"`        // Port
}
type RequestEventManagementCreateSyslogDestinationV1 struct {
	ConfigID    string `json:"configId,omitempty"`    // Required only for update syslog configuration
	Name        string `json:"name,omitempty"`        // Name
	Description string `json:"description,omitempty"` // Description
	Host        string `json:"host,omitempty"`        // Host
	Protocol    string `json:"protocol,omitempty"`    // Protocol
	Port        *int   `json:"port,omitempty"`        // Port
}
type RequestEventManagementCreateWebhookDestinationV1 struct {
	WebhookID    string                                                     `json:"webhookId,omitempty"`    // Required only for update webhook configuration
	Name         string                                                     `json:"name,omitempty"`         // Name
	Description  string                                                     `json:"description,omitempty"`  // Description
	URL          string                                                     `json:"url,omitempty"`          // Url
	Method       string                                                     `json:"method,omitempty"`       // Method
	TrustCert    *bool                                                      `json:"trustCert,omitempty"`    // Trust Cert
	Headers      *[]RequestEventManagementCreateWebhookDestinationV1Headers `json:"headers,omitempty"`      //
	IsProxyRoute *bool                                                      `json:"isProxyRoute,omitempty"` // Is Proxy Route
}
type RequestEventManagementCreateWebhookDestinationV1Headers struct {
	Name         string `json:"name,omitempty"`         // Name
	Value        string `json:"value,omitempty"`        // Value
	DefaultValue string `json:"defaultValue,omitempty"` // Default Value
	Encrypt      *bool  `json:"encrypt,omitempty"`      // Encrypt
}
type RequestEventManagementUpdateWebhookDestinationV1 struct {
	WebhookID    string                                                     `json:"webhookId,omitempty"`    // Required only for update webhook configuration
	Name         string                                                     `json:"name,omitempty"`         // Name
	Description  string                                                     `json:"description,omitempty"`  // Description
	URL          string                                                     `json:"url,omitempty"`          // Url
	Method       string                                                     `json:"method,omitempty"`       // Method
	TrustCert    *bool                                                      `json:"trustCert,omitempty"`    // Trust Cert
	Headers      *[]RequestEventManagementUpdateWebhookDestinationV1Headers `json:"headers,omitempty"`      //
	IsProxyRoute *bool                                                      `json:"isProxyRoute,omitempty"` // Is Proxy Route
}
type RequestEventManagementUpdateWebhookDestinationV1Headers struct {
	Name         string `json:"name,omitempty"`         // Name
	Value        string `json:"value,omitempty"`        // Value
	DefaultValue string `json:"defaultValue,omitempty"` // Default Value
	Encrypt      *bool  `json:"encrypt,omitempty"`      // Encrypt
}

//GetAuditLogParentRecordsV1 Get AuditLog Parent Records - 9590-7ae9-46ea-b1c6
/* Get Parent Audit Log Event instances from the Event-Hub


@param GetAuditLogParentRecordsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-audit-log-parent-records-v1
*/
func (s *EventManagementService) GetAuditLogParentRecordsV1(GetAuditLogParentRecordsV1QueryParams *GetAuditLogParentRecordsV1QueryParams) (*ResponseEventManagementGetAuditLogParentRecordsV1, *resty.Response, error) {
	path := "/dna/data/api/v1/event/event-series/audit-log/parent-records"

	queryString, _ := query.Values(GetAuditLogParentRecordsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementGetAuditLogParentRecordsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAuditLogParentRecordsV1(GetAuditLogParentRecordsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAuditLogParentRecordsV1")
	}

	result := response.Result().(*ResponseEventManagementGetAuditLogParentRecordsV1)
	return result, response, err

}

//GetAuditLogSummaryV1 Get AuditLog Summary - 4a87-484a-4df9-819e
/* Get Audit Log Summary from the Event-Hub


@param GetAuditLogSummaryV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-audit-log-summary-v1
*/
func (s *EventManagementService) GetAuditLogSummaryV1(GetAuditLogSummaryV1QueryParams *GetAuditLogSummaryV1QueryParams) (*ResponseEventManagementGetAuditLogSummaryV1, *resty.Response, error) {
	path := "/dna/data/api/v1/event/event-series/audit-log/summary"

	queryString, _ := query.Values(GetAuditLogSummaryV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementGetAuditLogSummaryV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAuditLogSummaryV1(GetAuditLogSummaryV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAuditLogSummaryV1")
	}

	result := response.Result().(*ResponseEventManagementGetAuditLogSummaryV1)
	return result, response, err

}

//GetAuditLogRecordsV1 Get AuditLog Records - 89a9-fafb-4d49-bd86
/* Get Audit Log Event instances from the Event-Hub


@param GetAuditLogRecordsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-audit-log-records-v1
*/
func (s *EventManagementService) GetAuditLogRecordsV1(GetAuditLogRecordsV1QueryParams *GetAuditLogRecordsV1QueryParams) (*ResponseEventManagementGetAuditLogRecordsV1, *resty.Response, error) {
	path := "/dna/data/api/v1/event/event-series/audit-logs"

	queryString, _ := query.Values(GetAuditLogRecordsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementGetAuditLogRecordsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAuditLogRecordsV1(GetAuditLogRecordsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAuditLogRecordsV1")
	}

	result := response.Result().(*ResponseEventManagementGetAuditLogRecordsV1)
	return result, response, err

}

//GetSNMPDestinationV1 Get SNMP Destination - ffbb-e92a-40e9-9ae6
/* Get SNMP Destination


@param GetSNMPDestinationV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-snmp-destination-v1
*/
func (s *EventManagementService) GetSNMPDestinationV1(GetSNMPDestinationV1QueryParams *GetSNMPDestinationV1QueryParams) (*ResponseEventManagementGetSNMPDestinationV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/dna-event/snmp-config"

	queryString, _ := query.Values(GetSNMPDestinationV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementGetSNMPDestinationV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSNMPDestinationV1(GetSNMPDestinationV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSnmpDestinationV1")
	}

	result := response.Result().(*ResponseEventManagementGetSNMPDestinationV1)
	return result, response, err

}

//GetStatusAPIForEventsV1 Get Status API for Events - f9bd-99c7-4bba-8832
/* Get the Status of events API calls with provided executionId as mandatory path parameter


@param executionID executionId path parameter. Execution ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-status-api-for-events-v1
*/
func (s *EventManagementService) GetStatusAPIForEventsV1(executionID string) (*ResponseEventManagementGetStatusAPIForEventsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/api-status/{executionId}"
	path = strings.Replace(path, "{executionId}", fmt.Sprintf("%v", executionID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseEventManagementGetStatusAPIForEventsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetStatusAPIForEventsV1(executionID)
		}
		return nil, response, fmt.Errorf("error with operation GetStatusApiForEventsV1")
	}

	result := response.Result().(*ResponseEventManagementGetStatusAPIForEventsV1)
	return result, response, err

}

//GetEmailDestinationV1 Get Email Destination - aebc-3bec-4858-8488
/* Get Email Destination



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-email-destination-v1
*/
func (s *EventManagementService) GetEmailDestinationV1() (*ResponseEventManagementGetEmailDestinationV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/email-config"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseEventManagementGetEmailDestinationV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetEmailDestinationV1()
		}
		return nil, response, fmt.Errorf("error with operation GetEmailDestinationV1")
	}

	result := response.Result().(*ResponseEventManagementGetEmailDestinationV1)
	return result, response, err

}

//GetNotificationsV1 Get Notifications - 8499-9b56-4afb-8657
/* Get the list of Published Notifications


@param GetNotificationsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-notifications-v1
*/
func (s *EventManagementService) GetNotificationsV1(GetNotificationsV1QueryParams *GetNotificationsV1QueryParams) (*ResponseEventManagementGetNotificationsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/event-series"

	queryString, _ := query.Values(GetNotificationsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementGetNotificationsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetNotificationsV1(GetNotificationsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetNotificationsV1")
	}

	result := response.Result().(*ResponseEventManagementGetNotificationsV1)
	return result, response, err

}

//CountOfNotificationsV1 Count of Notifications - 0eb8-faf7-42aa-abb7
/* Get the Count of Published Notifications


@param CountOfNotificationsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!count-of-notifications-v1
*/
func (s *EventManagementService) CountOfNotificationsV1(CountOfNotificationsV1QueryParams *CountOfNotificationsV1QueryParams) (*ResponseEventManagementCountOfNotificationsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/event-series/count"

	queryString, _ := query.Values(CountOfNotificationsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementCountOfNotificationsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.CountOfNotificationsV1(CountOfNotificationsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation CountOfNotificationsV1")
	}

	result := response.Result().(*ResponseEventManagementCountOfNotificationsV1)
	return result, response, err

}

//GetEventSubscriptionsV1 Get Event Subscriptions - dcaa-6bde-4feb-9152
/* Gets the list of Subscriptions's based on provided offset and limit (Deprecated)


@param GetEventSubscriptionsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-event-subscriptions-v1
*/
func (s *EventManagementService) GetEventSubscriptionsV1(GetEventSubscriptionsV1QueryParams *GetEventSubscriptionsV1QueryParams) (*ResponseEventManagementGetEventSubscriptionsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/subscription"

	queryString, _ := query.Values(GetEventSubscriptionsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementGetEventSubscriptionsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetEventSubscriptionsV1(GetEventSubscriptionsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetEventSubscriptionsV1")
	}

	result := response.Result().(*ResponseEventManagementGetEventSubscriptionsV1)
	return result, response, err

}

//GetEmailSubscriptionDetailsV1 Get Email Subscription Details - 339f-d9f5-4719-a410
/* Gets the list of subscription details for specified connectorType


@param GetEmailSubscriptionDetailsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-email-subscription-details-v1
*/
func (s *EventManagementService) GetEmailSubscriptionDetailsV1(GetEmailSubscriptionDetailsV1QueryParams *GetEmailSubscriptionDetailsV1QueryParams) (*ResponseEventManagementGetEmailSubscriptionDetailsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/subscription-details/email"

	queryString, _ := query.Values(GetEmailSubscriptionDetailsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementGetEmailSubscriptionDetailsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetEmailSubscriptionDetailsV1(GetEmailSubscriptionDetailsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetEmailSubscriptionDetailsV1")
	}

	result := response.Result().(*ResponseEventManagementGetEmailSubscriptionDetailsV1)
	return result, response, err

}

//GetRestWebhookSubscriptionDetailsV1 Get Rest/Webhook Subscription Details - eeb6-8baf-4338-bb23
/* Gets the list of subscription details for specified connectorType


@param GetRestWebhookSubscriptionDetailsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-rest-webhook-subscription-details-v1
*/
func (s *EventManagementService) GetRestWebhookSubscriptionDetailsV1(GetRestWebhookSubscriptionDetailsV1QueryParams *GetRestWebhookSubscriptionDetailsV1QueryParams) (*ResponseEventManagementGetRestWebhookSubscriptionDetailsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/subscription-details/rest"

	queryString, _ := query.Values(GetRestWebhookSubscriptionDetailsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementGetRestWebhookSubscriptionDetailsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetRestWebhookSubscriptionDetailsV1(GetRestWebhookSubscriptionDetailsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetRestWebhookSubscriptionDetailsV1")
	}

	result := response.Result().(*ResponseEventManagementGetRestWebhookSubscriptionDetailsV1)
	return result, response, err

}

//GetSyslogSubscriptionDetailsV1 Get Syslog Subscription Details - 1785-5b4e-4e69-a497
/* Gets the list of subscription details for specified connectorType


@param GetSyslogSubscriptionDetailsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-syslog-subscription-details-v1
*/
func (s *EventManagementService) GetSyslogSubscriptionDetailsV1(GetSyslogSubscriptionDetailsV1QueryParams *GetSyslogSubscriptionDetailsV1QueryParams) (*ResponseEventManagementGetSyslogSubscriptionDetailsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/subscription-details/syslog"

	queryString, _ := query.Values(GetSyslogSubscriptionDetailsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementGetSyslogSubscriptionDetailsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSyslogSubscriptionDetailsV1(GetSyslogSubscriptionDetailsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSyslogSubscriptionDetailsV1")
	}

	result := response.Result().(*ResponseEventManagementGetSyslogSubscriptionDetailsV1)
	return result, response, err

}

//CountOfEventSubscriptionsV1 Count of Event Subscriptions - 149b-7ba0-4e58-90b2
/* Returns the Count of EventSubscriptions


@param CountOfEventSubscriptionsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!count-of-event-subscriptions-v1
*/
func (s *EventManagementService) CountOfEventSubscriptionsV1(CountOfEventSubscriptionsV1QueryParams *CountOfEventSubscriptionsV1QueryParams) (*ResponseEventManagementCountOfEventSubscriptionsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/subscription/count"

	queryString, _ := query.Values(CountOfEventSubscriptionsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementCountOfEventSubscriptionsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.CountOfEventSubscriptionsV1(CountOfEventSubscriptionsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation CountOfEventSubscriptionsV1")
	}

	result := response.Result().(*ResponseEventManagementCountOfEventSubscriptionsV1)
	return result, response, err

}

//GetEmailEventSubscriptionsV1 Get Email Event Subscriptions - 39b2-0851-4b39-837e
/* Gets the list of email Subscriptions's based on provided query params


@param GetEmailEventSubscriptionsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-email-event-subscriptions-v1
*/
func (s *EventManagementService) GetEmailEventSubscriptionsV1(GetEmailEventSubscriptionsV1QueryParams *GetEmailEventSubscriptionsV1QueryParams) (*ResponseEventManagementGetEmailEventSubscriptionsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/subscription/email"

	queryString, _ := query.Values(GetEmailEventSubscriptionsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementGetEmailEventSubscriptionsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetEmailEventSubscriptionsV1(GetEmailEventSubscriptionsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetEmailEventSubscriptionsV1")
	}

	result := response.Result().(*ResponseEventManagementGetEmailEventSubscriptionsV1)
	return result, response, err

}

//GetRestWebhookEventSubscriptionsV1 Get Rest/Webhook Event Subscriptions - dcaa-6bde-4feb-9153
/* Gets the list of Rest/Webhook Subscriptions's based on provided query params


@param GetRestWebhookEventSubscriptionsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-rest-webhook-event-subscriptions-v1
*/
func (s *EventManagementService) GetRestWebhookEventSubscriptionsV1(GetRestWebhookEventSubscriptionsV1QueryParams *GetRestWebhookEventSubscriptionsV1QueryParams) (*ResponseEventManagementGetRestWebhookEventSubscriptionsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/subscription/rest"

	queryString, _ := query.Values(GetRestWebhookEventSubscriptionsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementGetRestWebhookEventSubscriptionsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetRestWebhookEventSubscriptionsV1(GetRestWebhookEventSubscriptionsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetRestWebhookEventSubscriptionsV1")
	}

	result := response.Result().(*ResponseEventManagementGetRestWebhookEventSubscriptionsV1)
	return result, response, err

}

//GetSyslogEventSubscriptionsV1 Get Syslog Event Subscriptions - c5a9-2a5b-4e6a-852e
/* Gets the list of Syslog Subscriptions's based on provided offset and limit


@param GetSyslogEventSubscriptionsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-syslog-event-subscriptions-v1
*/
func (s *EventManagementService) GetSyslogEventSubscriptionsV1(GetSyslogEventSubscriptionsV1QueryParams *GetSyslogEventSubscriptionsV1QueryParams) (*ResponseEventManagementGetSyslogEventSubscriptionsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/subscription/syslog"

	queryString, _ := query.Values(GetSyslogEventSubscriptionsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementGetSyslogEventSubscriptionsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSyslogEventSubscriptionsV1(GetSyslogEventSubscriptionsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSyslogEventSubscriptionsV1")
	}

	result := response.Result().(*ResponseEventManagementGetSyslogEventSubscriptionsV1)
	return result, response, err

}

//GetSyslogDestinationV1 Get Syslog Destination - 86a2-6b24-4828-b9dc
/* Get Syslog Destination


@param GetSyslogDestinationV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-syslog-destination-v1
*/
func (s *EventManagementService) GetSyslogDestinationV1(GetSyslogDestinationV1QueryParams *GetSyslogDestinationV1QueryParams) (*ResponseEventManagementGetSyslogDestinationV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/syslog-config"

	queryString, _ := query.Values(GetSyslogDestinationV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementGetSyslogDestinationV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSyslogDestinationV1(GetSyslogDestinationV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSyslogDestinationV1")
	}

	result := response.Result().(*ResponseEventManagementGetSyslogDestinationV1)
	return result, response, err

}

//GetWebhookDestinationV1 Get Webhook Destination - 2395-8ba8-4a7b-bd72
/* Get Webhook Destination


@param GetWebhookDestinationV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-webhook-destination-v1
*/
func (s *EventManagementService) GetWebhookDestinationV1(GetWebhookDestinationV1QueryParams *GetWebhookDestinationV1QueryParams) (*ResponseEventManagementGetWebhookDestinationV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/webhook"

	queryString, _ := query.Values(GetWebhookDestinationV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementGetWebhookDestinationV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetWebhookDestinationV1(GetWebhookDestinationV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetWebhookDestinationV1")
	}

	result := response.Result().(*ResponseEventManagementGetWebhookDestinationV1)
	return result, response, err

}

//GetEventsV1 Get Events - 44a3-9a07-4a6a-82a2
/* Gets the list of registered Events with provided eventIds or tags as mandatory


@param GetEventsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-events-v1
*/
func (s *EventManagementService) GetEventsV1(GetEventsV1QueryParams *GetEventsV1QueryParams) (*ResponseEventManagementGetEventsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/events"

	queryString, _ := query.Values(GetEventsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementGetEventsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetEventsV1(GetEventsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetEventsV1")
	}

	result := response.Result().(*ResponseEventManagementGetEventsV1)
	return result, response, err

}

//CountOfEventsV1 Count of Events - 6a9e-dac1-49ba-86cf
/* Get the count of registered events with provided eventIds or tags as mandatory


@param CountOfEventsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!count-of-events-v1
*/
func (s *EventManagementService) CountOfEventsV1(CountOfEventsV1QueryParams *CountOfEventsV1QueryParams) (*ResponseEventManagementCountOfEventsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/events/count"

	queryString, _ := query.Values(CountOfEventsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementCountOfEventsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.CountOfEventsV1(CountOfEventsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation CountOfEventsV1")
	}

	result := response.Result().(*ResponseEventManagementCountOfEventsV1)
	return result, response, err

}

//GetEventArtifactsV1 Get EventArtifacts - 73b1-d832-4c98-bc22
/* Gets the list of artifacts based on provided offset and limit


@param GetEventArtifactsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-event-artifacts-v1
*/
func (s *EventManagementService) GetEventArtifactsV1(GetEventArtifactsV1QueryParams *GetEventArtifactsV1QueryParams) (*ResponseEventManagementGetEventArtifactsV1, *resty.Response, error) {
	path := "/dna/system/api/v1/event/artifact"

	queryString, _ := query.Values(GetEventArtifactsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementGetEventArtifactsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetEventArtifactsV1(GetEventArtifactsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetEventArtifactsV1")
	}

	result := response.Result().(*ResponseEventManagementGetEventArtifactsV1)
	return result, response, err

}

//EventArtifactCountV1 EventArtifact Count - b78e-9bf7-4f8a-8321
/* Get the count of registered event artifacts.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!event-artifact-count-v1
*/
func (s *EventManagementService) EventArtifactCountV1() (*ResponseEventManagementEventArtifactCountV1, *resty.Response, error) {
	path := "/dna/system/api/v1/event/artifact/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseEventManagementEventArtifactCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.EventArtifactCountV1()
		}
		return nil, response, fmt.Errorf("error with operation EventArtifactCountV1")
	}

	result := response.Result().(*ResponseEventManagementEventArtifactCountV1)
	return result, response, err

}

//GetConnectorTypesV1 Get Connector Types - c0be-2a2b-4dc9-8bfb
/* Get the list of connector types



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-connector-types-v1
*/
func (s *EventManagementService) GetConnectorTypesV1() (*ResponseEventManagementGetConnectorTypesV1, *resty.Response, error) {
	path := "/dna/system/api/v1/event/config/connector-types"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseEventManagementGetConnectorTypesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetConnectorTypesV1()
		}
		return nil, response, fmt.Errorf("error with operation GetConnectorTypesV1")
	}

	result := response.Result().(*ResponseEventManagementGetConnectorTypesV1)
	return result, response, err

}

//CreateEmailDestinationV1 Create Email Destination - 84ab-8bff-4769-a7d6
/* Create Email Destination



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-email-destination-v1
*/
func (s *EventManagementService) CreateEmailDestinationV1(requestEventManagementCreateEmailDestinationV1 *RequestEventManagementCreateEmailDestinationV1) (*ResponseEventManagementCreateEmailDestinationV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/email-config"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEventManagementCreateEmailDestinationV1).
		SetResult(&ResponseEventManagementCreateEmailDestinationV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateEmailDestinationV1(requestEventManagementCreateEmailDestinationV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateEmailDestinationV1")
	}

	result := response.Result().(*ResponseEventManagementCreateEmailDestinationV1)
	return result, response, err

}

//CreateSNMPDestinationV1 Create SNMP Destination - 0e81-69c1-4a7a-965d
/* Create SNMP Destination



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-snmp-destination-v1
*/
func (s *EventManagementService) CreateSNMPDestinationV1(requestEventManagementCreateSNMPDestinationV1 *RequestEventManagementCreateSNMPDestinationV1) (*ResponseEventManagementCreateSNMPDestinationV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/snmp-config"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEventManagementCreateSNMPDestinationV1).
		SetResult(&ResponseEventManagementCreateSNMPDestinationV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateSNMPDestinationV1(requestEventManagementCreateSNMPDestinationV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateSnmpDestinationV1")
	}

	result := response.Result().(*ResponseEventManagementCreateSNMPDestinationV1)
	return result, response, err

}

//CreateEventSubscriptionsV1 Create Event Subscriptions - 4f9f-7a7b-40f9-90de
/* Subscribe SubscriptionEndpoint to list of registered events (Deprecated)



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-event-subscriptions-v1
*/
func (s *EventManagementService) CreateEventSubscriptionsV1(requestEventManagementCreateEventSubscriptionsV1 *RequestEventManagementCreateEventSubscriptionsV1) (*ResponseEventManagementCreateEventSubscriptionsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/subscription"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEventManagementCreateEventSubscriptionsV1).
		SetResult(&ResponseEventManagementCreateEventSubscriptionsV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateEventSubscriptionsV1(requestEventManagementCreateEventSubscriptionsV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateEventSubscriptionsV1")
	}

	result := response.Result().(*ResponseEventManagementCreateEventSubscriptionsV1)
	return result, response, err

}

//CreateEmailEventSubscriptionV1 Create Email Event Subscription - 7bbc-88c8-424a-840f
/* Create Email Subscription Endpoint for list of registered events.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-email-event-subscription-v1
*/
func (s *EventManagementService) CreateEmailEventSubscriptionV1(requestEventManagementCreateEmailEventSubscriptionV1 *RequestEventManagementCreateEmailEventSubscriptionV1) (*ResponseEventManagementCreateEmailEventSubscriptionV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/subscription/email"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEventManagementCreateEmailEventSubscriptionV1).
		SetResult(&ResponseEventManagementCreateEmailEventSubscriptionV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateEmailEventSubscriptionV1(requestEventManagementCreateEmailEventSubscriptionV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateEmailEventSubscriptionV1")
	}

	result := response.Result().(*ResponseEventManagementCreateEmailEventSubscriptionV1)
	return result, response, err

}

//CreateRestWebhookEventSubscriptionV1 Create Rest/Webhook Event Subscription - 9584-d988-45eb-b4b0
/* Create Rest/Webhook Subscription Endpoint for list of registered events



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-rest-webhook-event-subscription-v1
*/
func (s *EventManagementService) CreateRestWebhookEventSubscriptionV1(requestEventManagementCreateRestWebhookEventSubscriptionV1 *RequestEventManagementCreateRestWebhookEventSubscriptionV1) (*ResponseEventManagementCreateRestWebhookEventSubscriptionV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/subscription/rest"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEventManagementCreateRestWebhookEventSubscriptionV1).
		SetResult(&ResponseEventManagementCreateRestWebhookEventSubscriptionV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateRestWebhookEventSubscriptionV1(requestEventManagementCreateRestWebhookEventSubscriptionV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateRestWebhookEventSubscriptionV1")
	}

	result := response.Result().(*ResponseEventManagementCreateRestWebhookEventSubscriptionV1)
	return result, response, err

}

//CreateSyslogEventSubscriptionV1 Create Syslog Event Subscription - 919a-8bb7-445a-88fe
/* Create Syslog Subscription Endpoint for list of registered events



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-syslog-event-subscription-v1
*/
func (s *EventManagementService) CreateSyslogEventSubscriptionV1(requestEventManagementCreateSyslogEventSubscriptionV1 *RequestEventManagementCreateSyslogEventSubscriptionV1) (*ResponseEventManagementCreateSyslogEventSubscriptionV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/subscription/syslog"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEventManagementCreateSyslogEventSubscriptionV1).
		SetResult(&ResponseEventManagementCreateSyslogEventSubscriptionV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateSyslogEventSubscriptionV1(requestEventManagementCreateSyslogEventSubscriptionV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateSyslogEventSubscriptionV1")
	}

	result := response.Result().(*ResponseEventManagementCreateSyslogEventSubscriptionV1)
	return result, response, err

}

//CreateSyslogDestinationV1 Create Syslog Destination - 08ad-2bd5-47fa-9227
/* Create Syslog Destination



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-syslog-destination-v1
*/
func (s *EventManagementService) CreateSyslogDestinationV1(requestEventManagementCreateSyslogDestinationV1 *RequestEventManagementCreateSyslogDestinationV1) (*ResponseEventManagementCreateSyslogDestinationV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/syslog-config"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEventManagementCreateSyslogDestinationV1).
		SetResult(&ResponseEventManagementCreateSyslogDestinationV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateSyslogDestinationV1(requestEventManagementCreateSyslogDestinationV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateSyslogDestinationV1")
	}

	result := response.Result().(*ResponseEventManagementCreateSyslogDestinationV1)
	return result, response, err

}

//CreateWebhookDestinationV1 Create Webhook Destination - 4bbd-580a-4bab-80b7
/* Create Webhook Destination



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-webhook-destination-v1
*/
func (s *EventManagementService) CreateWebhookDestinationV1(requestEventManagementCreateWebhookDestinationV1 *RequestEventManagementCreateWebhookDestinationV1) (*ResponseEventManagementCreateWebhookDestinationV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/webhook"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEventManagementCreateWebhookDestinationV1).
		SetResult(&ResponseEventManagementCreateWebhookDestinationV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateWebhookDestinationV1(requestEventManagementCreateWebhookDestinationV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateWebhookDestinationV1")
	}

	result := response.Result().(*ResponseEventManagementCreateWebhookDestinationV1)
	return result, response, err

}

//UpdateEmailDestinationV1 Update Email Destination - 8285-3bf9-4aba-a65f
/* Update Email Destination


 */
func (s *EventManagementService) UpdateEmailDestinationV1(requestEventManagementUpdateEmailDestinationV1 *RequestEventManagementUpdateEmailDestinationV1) (*ResponseEventManagementUpdateEmailDestinationV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/email-config"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEventManagementUpdateEmailDestinationV1).
		SetResult(&ResponseEventManagementUpdateEmailDestinationV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateEmailDestinationV1(requestEventManagementUpdateEmailDestinationV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateEmailDestinationV1")
	}

	result := response.Result().(*ResponseEventManagementUpdateEmailDestinationV1)
	return result, response, err

}

//UpdateSNMPDestinationV1 Update SNMP Destination - 26b7-ab81-4c5a-811b
/* Update SNMP Destination


 */
func (s *EventManagementService) UpdateSNMPDestinationV1(requestEventManagementUpdateSNMPDestinationV1 *RequestEventManagementUpdateSNMPDestinationV1) (*ResponseEventManagementUpdateSNMPDestinationV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/snmp-config"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEventManagementUpdateSNMPDestinationV1).
		SetResult(&ResponseEventManagementUpdateSNMPDestinationV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateSNMPDestinationV1(requestEventManagementUpdateSNMPDestinationV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateSnmpDestinationV1")
	}

	result := response.Result().(*ResponseEventManagementUpdateSNMPDestinationV1)
	return result, response, err

}

//UpdateEventSubscriptionsV1 Update Event Subscriptions - 579a-6a72-48cb-94cf
/* Update SubscriptionEndpoint to list of registered events(Deprecated)


 */
func (s *EventManagementService) UpdateEventSubscriptionsV1(requestEventManagementUpdateEventSubscriptionsV1 *RequestEventManagementUpdateEventSubscriptionsV1) (*ResponseEventManagementUpdateEventSubscriptionsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/subscription"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEventManagementUpdateEventSubscriptionsV1).
		SetResult(&ResponseEventManagementUpdateEventSubscriptionsV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateEventSubscriptionsV1(requestEventManagementUpdateEventSubscriptionsV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateEventSubscriptionsV1")
	}

	result := response.Result().(*ResponseEventManagementUpdateEventSubscriptionsV1)
	return result, response, err

}

//UpdateEmailEventSubscriptionV1 Update Email Event Subscription - 87b2-2b83-46bb-8983
/* Update Email Subscription Endpoint for list of registered events


 */
func (s *EventManagementService) UpdateEmailEventSubscriptionV1(requestEventManagementUpdateEmailEventSubscriptionV1 *RequestEventManagementUpdateEmailEventSubscriptionV1) (*ResponseEventManagementUpdateEmailEventSubscriptionV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/subscription/email"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEventManagementUpdateEmailEventSubscriptionV1).
		SetResult(&ResponseEventManagementUpdateEmailEventSubscriptionV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateEmailEventSubscriptionV1(requestEventManagementUpdateEmailEventSubscriptionV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateEmailEventSubscriptionV1")
	}

	result := response.Result().(*ResponseEventManagementUpdateEmailEventSubscriptionV1)
	return result, response, err

}

//UpdateRestWebhookEventSubscriptionV1 Update Rest/Webhook Event Subscription - ce81-f9c5-4fc8-b576
/* Update Rest/Webhook Subscription Endpoint for list of registered events


 */
func (s *EventManagementService) UpdateRestWebhookEventSubscriptionV1(requestEventManagementUpdateRestWebhookEventSubscriptionV1 *RequestEventManagementUpdateRestWebhookEventSubscriptionV1) (*ResponseEventManagementUpdateRestWebhookEventSubscriptionV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/subscription/rest"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEventManagementUpdateRestWebhookEventSubscriptionV1).
		SetResult(&ResponseEventManagementUpdateRestWebhookEventSubscriptionV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateRestWebhookEventSubscriptionV1(requestEventManagementUpdateRestWebhookEventSubscriptionV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateRestWebhookEventSubscriptionV1")
	}

	result := response.Result().(*ResponseEventManagementUpdateRestWebhookEventSubscriptionV1)
	return result, response, err

}

//UpdateSyslogEventSubscriptionV1 Update Syslog Event Subscription - 6285-cbc1-4039-9ace
/* Update Syslog Subscription Endpoint for list of registered events


 */
func (s *EventManagementService) UpdateSyslogEventSubscriptionV1(requestEventManagementUpdateSyslogEventSubscriptionV1 *RequestEventManagementUpdateSyslogEventSubscriptionV1) (*ResponseEventManagementUpdateSyslogEventSubscriptionV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/subscription/syslog"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEventManagementUpdateSyslogEventSubscriptionV1).
		SetResult(&ResponseEventManagementUpdateSyslogEventSubscriptionV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateSyslogEventSubscriptionV1(requestEventManagementUpdateSyslogEventSubscriptionV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateSyslogEventSubscriptionV1")
	}

	result := response.Result().(*ResponseEventManagementUpdateSyslogEventSubscriptionV1)
	return result, response, err

}

//UpdateSyslogDestinationV1 Update Syslog Destination - eeac-fbe9-4518-b3fb
/* Update Syslog Destination


 */
func (s *EventManagementService) UpdateSyslogDestinationV1(requestEventManagementUpdateSyslogDestinationV1 *RequestEventManagementUpdateSyslogDestinationV1) (*ResponseEventManagementUpdateSyslogDestinationV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/syslog-config"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEventManagementUpdateSyslogDestinationV1).
		SetResult(&ResponseEventManagementUpdateSyslogDestinationV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateSyslogDestinationV1(requestEventManagementUpdateSyslogDestinationV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateSyslogDestinationV1")
	}

	result := response.Result().(*ResponseEventManagementUpdateSyslogDestinationV1)
	return result, response, err

}

//UpdateWebhookDestinationV1 Update Webhook Destination - 06b5-da28-423a-9bf6
/* Update Webhook Destination


 */
func (s *EventManagementService) UpdateWebhookDestinationV1(requestEventManagementUpdateWebhookDestinationV1 *RequestEventManagementUpdateWebhookDestinationV1) (*ResponseEventManagementUpdateWebhookDestinationV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/webhook"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEventManagementUpdateWebhookDestinationV1).
		SetResult(&ResponseEventManagementUpdateWebhookDestinationV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateWebhookDestinationV1(requestEventManagementUpdateWebhookDestinationV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateWebhookDestinationV1")
	}

	result := response.Result().(*ResponseEventManagementUpdateWebhookDestinationV1)
	return result, response, err

}

//DeleteEventSubscriptionsV1 Delete Event Subscriptions - 9398-1baa-4079-9483
/* Delete EventSubscriptions


@param DeleteEventSubscriptionsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-event-subscriptions-v1
*/
func (s *EventManagementService) DeleteEventSubscriptionsV1(DeleteEventSubscriptionsV1QueryParams *DeleteEventSubscriptionsV1QueryParams) (*ResponseEventManagementDeleteEventSubscriptionsV1, *resty.Response, error) {
	//DeleteEventSubscriptionsV1QueryParams *DeleteEventSubscriptionsV1QueryParams
	path := "/dna/intent/api/v1/event/subscription"

	queryString, _ := query.Values(DeleteEventSubscriptionsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementDeleteEventSubscriptionsV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteEventSubscriptionsV1(
				DeleteEventSubscriptionsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteEventSubscriptionsV1")
	}

	result := response.Result().(*ResponseEventManagementDeleteEventSubscriptionsV1)
	return result, response, err

}

// Alias Function
func (s *EventManagementService) GetEventSubscriptions(GetEventSubscriptionsV1QueryParams *GetEventSubscriptionsV1QueryParams) (*ResponseEventManagementGetEventSubscriptionsV1, *resty.Response, error) {
	return s.GetEventSubscriptionsV1(GetEventSubscriptionsV1QueryParams)
}

// Alias Function
func (s *EventManagementService) UpdateSyslogDestination(requestEventManagementUpdateSyslogDestinationV1 *RequestEventManagementUpdateSyslogDestinationV1) (*ResponseEventManagementUpdateSyslogDestinationV1, *resty.Response, error) {
	return s.UpdateSyslogDestinationV1(requestEventManagementUpdateSyslogDestinationV1)
}

// Alias Function
func (s *EventManagementService) CreateRestWebhookEventSubscription(requestEventManagementCreateRestWebhookEventSubscriptionV1 *RequestEventManagementCreateRestWebhookEventSubscriptionV1) (*ResponseEventManagementCreateRestWebhookEventSubscriptionV1, *resty.Response, error) {
	return s.CreateRestWebhookEventSubscriptionV1(requestEventManagementCreateRestWebhookEventSubscriptionV1)
}

// Alias Function
func (s *EventManagementService) UpdateEmailDestination(requestEventManagementUpdateEmailDestinationV1 *RequestEventManagementUpdateEmailDestinationV1) (*ResponseEventManagementUpdateEmailDestinationV1, *resty.Response, error) {
	return s.UpdateEmailDestinationV1(requestEventManagementUpdateEmailDestinationV1)
}

// Alias Function
func (s *EventManagementService) GetRestWebhookSubscriptionDetails(GetRestWebhookSubscriptionDetailsV1QueryParams *GetRestWebhookSubscriptionDetailsV1QueryParams) (*ResponseEventManagementGetRestWebhookSubscriptionDetailsV1, *resty.Response, error) {
	return s.GetRestWebhookSubscriptionDetailsV1(GetRestWebhookSubscriptionDetailsV1QueryParams)
}

// Alias Function
func (s *EventManagementService) GetRestWebhookEventSubscriptions(GetRestWebhookEventSubscriptionsV1QueryParams *GetRestWebhookEventSubscriptionsV1QueryParams) (*ResponseEventManagementGetRestWebhookEventSubscriptionsV1, *resty.Response, error) {
	return s.GetRestWebhookEventSubscriptionsV1(GetRestWebhookEventSubscriptionsV1QueryParams)
}

// Alias Function
func (s *EventManagementService) DeleteEventSubscriptions(DeleteEventSubscriptionsV1QueryParams *DeleteEventSubscriptionsV1QueryParams) (*ResponseEventManagementDeleteEventSubscriptionsV1, *resty.Response, error) {
	return s.DeleteEventSubscriptionsV1(DeleteEventSubscriptionsV1QueryParams)
}

// Alias Function
func (s *EventManagementService) GetEmailDestination() (*ResponseEventManagementGetEmailDestinationV1, *resty.Response, error) {
	return s.GetEmailDestinationV1()
}

// Alias Function
func (s *EventManagementService) CreateSNMPDestination(requestEventManagementCreateSNMPDestinationV1 *RequestEventManagementCreateSNMPDestinationV1) (*ResponseEventManagementCreateSNMPDestinationV1, *resty.Response, error) {
	return s.CreateSNMPDestinationV1(requestEventManagementCreateSNMPDestinationV1)
}

// Alias Function
func (s *EventManagementService) UpdateEventSubscriptions(requestEventManagementUpdateEventSubscriptionsV1 *RequestEventManagementUpdateEventSubscriptionsV1) (*ResponseEventManagementUpdateEventSubscriptionsV1, *resty.Response, error) {
	return s.UpdateEventSubscriptionsV1(requestEventManagementUpdateEventSubscriptionsV1)
}

// Alias Function
func (s *EventManagementService) EventArtifactCount() (*ResponseEventManagementEventArtifactCountV1, *resty.Response, error) {
	return s.EventArtifactCountV1()
}

// Alias Function
func (s *EventManagementService) GetSNMPDestination(GetSNMPDestinationV1QueryParams *GetSNMPDestinationV1QueryParams) (*ResponseEventManagementGetSNMPDestinationV1, *resty.Response, error) {
	return s.GetSNMPDestinationV1(GetSNMPDestinationV1QueryParams)
}

// Alias Function
func (s *EventManagementService) CreateEmailEventSubscription(requestEventManagementCreateEmailEventSubscriptionV1 *RequestEventManagementCreateEmailEventSubscriptionV1) (*ResponseEventManagementCreateEmailEventSubscriptionV1, *resty.Response, error) {
	return s.CreateEmailEventSubscriptionV1(requestEventManagementCreateEmailEventSubscriptionV1)
}

// Alias Function
func (s *EventManagementService) CountOfEventSubscriptions(CountOfEventSubscriptionsV1QueryParams *CountOfEventSubscriptionsV1QueryParams) (*ResponseEventManagementCountOfEventSubscriptionsV1, *resty.Response, error) {
	return s.CountOfEventSubscriptionsV1(CountOfEventSubscriptionsV1QueryParams)
}

// Alias Function
func (s *EventManagementService) CreateWebhookDestination(requestEventManagementCreateWebhookDestinationV1 *RequestEventManagementCreateWebhookDestinationV1) (*ResponseEventManagementCreateWebhookDestinationV1, *resty.Response, error) {
	return s.CreateWebhookDestinationV1(requestEventManagementCreateWebhookDestinationV1)
}

// Alias Function
func (s *EventManagementService) GetEmailSubscriptionDetails(GetEmailSubscriptionDetailsV1QueryParams *GetEmailSubscriptionDetailsV1QueryParams) (*ResponseEventManagementGetEmailSubscriptionDetailsV1, *resty.Response, error) {
	return s.GetEmailSubscriptionDetailsV1(GetEmailSubscriptionDetailsV1QueryParams)
}

// Alias Function
func (s *EventManagementService) GetStatusAPIForEvents(executionID string) (*ResponseEventManagementGetStatusAPIForEventsV1, *resty.Response, error) {
	return s.GetStatusAPIForEventsV1(executionID)
}

// Alias Function
func (s *EventManagementService) GetEvents(GetEventsV1QueryParams *GetEventsV1QueryParams) (*ResponseEventManagementGetEventsV1, *resty.Response, error) {
	return s.GetEventsV1(GetEventsV1QueryParams)
}

// Alias Function
func (s *EventManagementService) UpdateSNMPDestination(requestEventManagementUpdateSNMPDestinationV1 *RequestEventManagementUpdateSNMPDestinationV1) (*ResponseEventManagementUpdateSNMPDestinationV1, *resty.Response, error) {
	return s.UpdateSNMPDestinationV1(requestEventManagementUpdateSNMPDestinationV1)
}

// Alias Function
func (s *EventManagementService) UpdateSyslogEventSubscription(requestEventManagementUpdateSyslogEventSubscriptionV1 *RequestEventManagementUpdateSyslogEventSubscriptionV1) (*ResponseEventManagementUpdateSyslogEventSubscriptionV1, *resty.Response, error) {
	return s.UpdateSyslogEventSubscriptionV1(requestEventManagementUpdateSyslogEventSubscriptionV1)
}

// Alias Function
func (s *EventManagementService) CreateEmailDestination(requestEventManagementCreateEmailDestinationV1 *RequestEventManagementCreateEmailDestinationV1) (*ResponseEventManagementCreateEmailDestinationV1, *resty.Response, error) {
	return s.CreateEmailDestinationV1(requestEventManagementCreateEmailDestinationV1)
}

// Alias Function
func (s *EventManagementService) CreateSyslogEventSubscription(requestEventManagementCreateSyslogEventSubscriptionV1 *RequestEventManagementCreateSyslogEventSubscriptionV1) (*ResponseEventManagementCreateSyslogEventSubscriptionV1, *resty.Response, error) {
	return s.CreateSyslogEventSubscriptionV1(requestEventManagementCreateSyslogEventSubscriptionV1)
}

// Alias Function
func (s *EventManagementService) UpdateEmailEventSubscription(requestEventManagementUpdateEmailEventSubscriptionV1 *RequestEventManagementUpdateEmailEventSubscriptionV1) (*ResponseEventManagementUpdateEmailEventSubscriptionV1, *resty.Response, error) {
	return s.UpdateEmailEventSubscriptionV1(requestEventManagementUpdateEmailEventSubscriptionV1)
}

// Alias Function
func (s *EventManagementService) UpdateWebhookDestination(requestEventManagementUpdateWebhookDestinationV1 *RequestEventManagementUpdateWebhookDestinationV1) (*ResponseEventManagementUpdateWebhookDestinationV1, *resty.Response, error) {
	return s.UpdateWebhookDestinationV1(requestEventManagementUpdateWebhookDestinationV1)
}

// Alias Function
func (s *EventManagementService) GetSyslogDestination(GetSyslogDestinationV1QueryParams *GetSyslogDestinationV1QueryParams) (*ResponseEventManagementGetSyslogDestinationV1, *resty.Response, error) {
	return s.GetSyslogDestinationV1(GetSyslogDestinationV1QueryParams)
}

// Alias Function
func (s *EventManagementService) CreateEventSubscriptions(requestEventManagementCreateEventSubscriptionsV1 *RequestEventManagementCreateEventSubscriptionsV1) (*ResponseEventManagementCreateEventSubscriptionsV1, *resty.Response, error) {
	return s.CreateEventSubscriptionsV1(requestEventManagementCreateEventSubscriptionsV1)
}

// Alias Function
func (s *EventManagementService) GetNotifications(GetNotificationsV1QueryParams *GetNotificationsV1QueryParams) (*ResponseEventManagementGetNotificationsV1, *resty.Response, error) {
	return s.GetNotificationsV1(GetNotificationsV1QueryParams)
}

// Alias Function
func (s *EventManagementService) GetSyslogSubscriptionDetails(GetSyslogSubscriptionDetailsV1QueryParams *GetSyslogSubscriptionDetailsV1QueryParams) (*ResponseEventManagementGetSyslogSubscriptionDetailsV1, *resty.Response, error) {
	return s.GetSyslogSubscriptionDetailsV1(GetSyslogSubscriptionDetailsV1QueryParams)
}

// Alias Function
func (s *EventManagementService) GetConnectorTypes() (*ResponseEventManagementGetConnectorTypesV1, *resty.Response, error) {
	return s.GetConnectorTypesV1()
}

// Alias Function
func (s *EventManagementService) GetEmailEventSubscriptions(GetEmailEventSubscriptionsV1QueryParams *GetEmailEventSubscriptionsV1QueryParams) (*ResponseEventManagementGetEmailEventSubscriptionsV1, *resty.Response, error) {
	return s.GetEmailEventSubscriptionsV1(GetEmailEventSubscriptionsV1QueryParams)
}

// Alias Function
func (s *EventManagementService) GetAuditLogRecords(GetAuditLogRecordsV1QueryParams *GetAuditLogRecordsV1QueryParams) (*ResponseEventManagementGetAuditLogRecordsV1, *resty.Response, error) {
	return s.GetAuditLogRecordsV1(GetAuditLogRecordsV1QueryParams)
}

// Alias Function
func (s *EventManagementService) UpdateRestWebhookEventSubscription(requestEventManagementUpdateRestWebhookEventSubscriptionV1 *RequestEventManagementUpdateRestWebhookEventSubscriptionV1) (*ResponseEventManagementUpdateRestWebhookEventSubscriptionV1, *resty.Response, error) {
	return s.UpdateRestWebhookEventSubscriptionV1(requestEventManagementUpdateRestWebhookEventSubscriptionV1)
}

// Alias Function
func (s *EventManagementService) CountOfNotifications(CountOfNotificationsV1QueryParams *CountOfNotificationsV1QueryParams) (*ResponseEventManagementCountOfNotificationsV1, *resty.Response, error) {
	return s.CountOfNotificationsV1(CountOfNotificationsV1QueryParams)
}

// Alias Function
func (s *EventManagementService) CreateSyslogDestination(requestEventManagementCreateSyslogDestinationV1 *RequestEventManagementCreateSyslogDestinationV1) (*ResponseEventManagementCreateSyslogDestinationV1, *resty.Response, error) {
	return s.CreateSyslogDestinationV1(requestEventManagementCreateSyslogDestinationV1)
}

// Alias Function
func (s *EventManagementService) GetSyslogEventSubscriptions(GetSyslogEventSubscriptionsV1QueryParams *GetSyslogEventSubscriptionsV1QueryParams) (*ResponseEventManagementGetSyslogEventSubscriptionsV1, *resty.Response, error) {
	return s.GetSyslogEventSubscriptionsV1(GetSyslogEventSubscriptionsV1QueryParams)
}

// Alias Function
func (s *EventManagementService) GetAuditLogSummary(GetAuditLogSummaryV1QueryParams *GetAuditLogSummaryV1QueryParams) (*ResponseEventManagementGetAuditLogSummaryV1, *resty.Response, error) {
	return s.GetAuditLogSummaryV1(GetAuditLogSummaryV1QueryParams)
}

// Alias Function
func (s *EventManagementService) GetEventArtifacts(GetEventArtifactsV1QueryParams *GetEventArtifactsV1QueryParams) (*ResponseEventManagementGetEventArtifactsV1, *resty.Response, error) {
	return s.GetEventArtifactsV1(GetEventArtifactsV1QueryParams)
}

// Alias Function
func (s *EventManagementService) GetAuditLogParentRecords(GetAuditLogParentRecordsV1QueryParams *GetAuditLogParentRecordsV1QueryParams) (*ResponseEventManagementGetAuditLogParentRecordsV1, *resty.Response, error) {
	return s.GetAuditLogParentRecordsV1(GetAuditLogParentRecordsV1QueryParams)
}

// Alias Function
func (s *EventManagementService) CountOfEvents(CountOfEventsV1QueryParams *CountOfEventsV1QueryParams) (*ResponseEventManagementCountOfEventsV1, *resty.Response, error) {
	return s.CountOfEventsV1(CountOfEventsV1QueryParams)
}

// Alias Function
func (s *EventManagementService) GetWebhookDestination(GetWebhookDestinationV1QueryParams *GetWebhookDestinationV1QueryParams) (*ResponseEventManagementGetWebhookDestinationV1, *resty.Response, error) {
	return s.GetWebhookDestinationV1(GetWebhookDestinationV1QueryParams)
}
