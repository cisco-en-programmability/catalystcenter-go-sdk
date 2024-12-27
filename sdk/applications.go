package catalyst

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type ApplicationsService service

type RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1QueryParams struct {
	StartTime         float64 `url:"startTime,omitempty"`         //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime           float64 `url:"endTime,omitempty"`           //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	Limit             float64 `url:"limit,omitempty"`             //Maximum number of records to return
	Offset            float64 `url:"offset,omitempty"`            //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	SortBy            string  `url:"sortBy,omitempty"`            //A field within the response to sort by.
	Order             string  `url:"order,omitempty"`             //The sort order of the field ascending or descending.
	SiteID            string  `url:"siteId,omitempty"`            //The site UUID without the top level hierarchy. `siteId` is mandatory. `siteId` must be a site UUID of a building. (Ex."buildingUuid") Examples: `siteId=buildingUuid` (single siteId requested) `siteId=buildingUuid1&siteId=buildingUuid2` (multiple siteId requested)
	SSID              string  `url:"ssid,omitempty"`              //In the context of a network application, SSID refers to the name of the wireless network to which the client connects. Examples: `ssid=Alpha` (single ssid requested) `ssid=Alpha&ssid=Guest` (multiple ssid requested)
	ApplicationName   string  `url:"applicationName,omitempty"`   //Name of the application for which the experience data is intended. Examples: `applicationName=webex` (single applicationName requested) `applicationName=webex&applicationName=teams` (multiple applicationName requested)
	BusinessRelevance string  `url:"businessRelevance,omitempty"` //The application can be chosen to be categorized as business-relevant, irrelevant, or default (neutral). By doing so, the assurance application prioritizes the monitoring and analysis of business-relevant data, ensuring critical insights are captured. Applications marked as irrelevant or default are selectively excluded from certain data sets, streamlining focus on what's most important for business outcomes.
	Attribute         string  `url:"attribute,omitempty"`         //List of attributes related to resource that can be requested to only be part of the response along with the required attributes. Supported attributes are applicationName, siteId, exporterIpAddress, exporterNetworkDeviceId, healthScore, businessRelevance, usage, throughput, packetLossPercent, networkLatency, applicationServerLatency, clientNetworkLatency, serverNetworkLatency, trafficClass, jitter, ssid Examples: `attribute=healthScore` (single attribute requested) `attribute=healthScore&attribute=ssid&attribute=jitter` (multiple attribute requested)
}
type RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringV1QueryParams struct {
	StartTime         float64 `url:"startTime,omitempty"`         //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime           float64 `url:"endTime,omitempty"`           //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	SiteID            string  `url:"siteId,omitempty"`            //The site UUID without the top level hierarchy. `siteId` is mandatory. `siteId` must be a site UUID of a building. (Ex."buildingUuid") Examples: `siteId=buildingUuid` (single siteId requested) `siteId=buildingUuid1&siteId=buildingUuid2` (multiple siteId requested)
	SSID              string  `url:"ssid,omitempty"`              //In the context of a network application, SSID refers to the name of the wireless network to which the client connects. Examples: `ssid=Alpha` (single ssid requested) `ssid=Alpha&ssid=Guest` (multiple ssid requested)
	ApplicationName   string  `url:"applicationName,omitempty"`   //Name of the application for which the experience data is intended. Examples: `applicationName=webex` (single applicationName requested) `applicationName=webex&applicationName=teams` (multiple applicationName requested)
	BusinessRelevance string  `url:"businessRelevance,omitempty"` //The application can be chosen to be categorized as business-relevant, irrelevant, or default (neutral). By doing so, the assurance application prioritizes the monitoring and analysis of business-relevant data, ensuring critical insights are captured. Applications marked as irrelevant or default are selectively excluded from certain data sets, streamlining focus on what's most important for business outcomes.
}
type RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1HeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type ApplicationsV1QueryParams struct {
	SiteID            string  `url:"siteId,omitempty"`            //Assurance site UUID value (Cannot be submitted together with deviceId and clientMac)
	DeviceID          string  `url:"deviceId,omitempty"`          //Assurance device UUID value (Cannot be submitted together with siteId and clientMac)
	MacAddress        string  `url:"macAddress,omitempty"`        //Client device's MAC address (Cannot be submitted together with siteId and deviceId)
	StartTime         float64 `url:"startTime,omitempty"`         //Starting epoch time in milliseconds of time window
	EndTime           float64 `url:"endTime,omitempty"`           //Ending epoch time in milliseconds of time window
	ApplicationHealth string  `url:"applicationHealth,omitempty"` //Application health category (POOR, FAIR, or GOOD.  Optionally use with siteId only)
	Offset            float64 `url:"offset,omitempty"`            //The offset of the first application in the returned data (optionally used with siteId only)
	Limit             float64 `url:"limit,omitempty"`             //The max number of application entries in returned data [1, 1000] (optionally used with siteId only)
	ApplicationName   string  `url:"applicationName,omitempty"`   //The name of the application to get information on
}

type ResponseApplicationsRetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1 struct {
	Response *[]ResponseApplicationsRetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1Response `json:"response,omitempty"` //

	Page *ResponseApplicationsRetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1Page `json:"page,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseApplicationsRetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1Response struct {
	ID string `json:"id,omitempty"` // Id

	ApplicationName string `json:"applicationName,omitempty"` // Application Name

	BusinessRelevance string `json:"businessRelevance,omitempty"` // Business Relevance

	SiteID string `json:"siteId,omitempty"` // Site Id

	ExporterIPAddress string `json:"exporterIpAddress,omitempty"` // Exporter Ip Address

	ExporterNetworkDeviceID string `json:"exporterNetworkDeviceId,omitempty"` // Exporter Network Device Id

	HealthScore *int `json:"healthScore,omitempty"` // Health Score

	Usage *float64 `json:"usage,omitempty"` // Usage

	Throughput *float64 `json:"throughput,omitempty"` // Throughput

	PacketLossPercent *float64 `json:"packetLossPercent,omitempty"` // Packet Loss Percent

	NetworkLatency *float64 `json:"networkLatency,omitempty"` // Network Latency

	ApplicationServerLatency *float64 `json:"applicationServerLatency,omitempty"` // Application Server Latency

	ClientNetworkLatency *float64 `json:"clientNetworkLatency,omitempty"` // Client Network Latency

	ServerNetworkLatency *float64 `json:"serverNetworkLatency,omitempty"` // Server Network Latency

	TrafficClass string `json:"trafficClass,omitempty"` // Traffic Class

	Jitter *float64 `json:"jitter,omitempty"` // Jitter

	SSID string `json:"ssid,omitempty"` // Ssid
}
type ResponseApplicationsRetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	Count *int `json:"count,omitempty"` // Count

	SortBy *[]ResponseApplicationsRetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1PageSortBy `json:"sortBy,omitempty"` //
}
type ResponseApplicationsRetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1PageSortBy struct {
	Name string `json:"name,omitempty"` // Name

	Order string `json:"order,omitempty"` // Order
}
type ResponseApplicationsRetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringV1 struct {
	Response *ResponseApplicationsRetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseApplicationsRetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1 struct {
	Response *[]ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1Response `json:"response,omitempty"` //

	Page *ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1Page `json:"page,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1Response struct {
	Timestamp *int `json:"timestamp,omitempty"` // Timestamp

	Attributes *[]ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1ResponseAttributes `json:"attributes,omitempty"` //

	AggregateAttributes *[]ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1ResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Groups *[]ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1ResponseGroups `json:"groups,omitempty"` //
}
type ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1ResponseAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Value string `json:"value,omitempty"` // Value
}
type ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1ResponseAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function

	Value string `json:"value,omitempty"` // Value
}
type ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1ResponseGroups struct {
	ID string `json:"id,omitempty"` // Id

	Attributes *[]ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1ResponseGroupsAttributes `json:"attributes,omitempty"` //

	AggregateAttributes *[]ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1ResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1ResponseGroupsAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Value string `json:"value,omitempty"` // Value
}
type ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1ResponseGroupsAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function

	Value string `json:"value,omitempty"` // Value
}
type ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Cursor string `json:"cursor,omitempty"` // Cursor

	Count *int `json:"count,omitempty"` // Count

	TimeSortOrder string `json:"timeSortOrder,omitempty"` // Time Sort Order
}
type ResponseApplicationsApplicationsV1 struct {
	Version    string                                        `json:"version,omitempty"`    // API version
	TotalCount *int                                          `json:"totalCount,omitempty"` // Count of items in response
	Response   *[]ResponseApplicationsApplicationsV1Response `json:"response,omitempty"`   //
}
type ResponseApplicationsApplicationsV1Response struct {
	Name                     string                                                              `json:"name,omitempty"`                     // Application name
	Health                   *int                                                                `json:"health,omitempty"`                   // Health score
	BusinessRelevance        string                                                              `json:"businessRelevance,omitempty"`        // Application's business relevance
	TrafficClass             string                                                              `json:"trafficClass,omitempty"`             // Application's traffic class
	UsageBytes               *int                                                                `json:"usageBytes,omitempty"`               // Usage amount in bytes
	AverageThroughput        *float64                                                            `json:"averageThroughput,omitempty"`        // Average throughput of application
	PacketLossPercent        *ResponseApplicationsApplicationsV1ResponsePacketLossPercent        `json:"packetLossPercent,omitempty"`        // Packet loss percentage for application
	NetworkLatency           *ResponseApplicationsApplicationsV1ResponseNetworkLatency           `json:"networkLatency,omitempty"`           // Network latency for application
	Jitter                   *ResponseApplicationsApplicationsV1ResponseJitter                   `json:"jitter,omitempty"`                   // Jitter for application
	ApplicationServerLatency *ResponseApplicationsApplicationsV1ResponseApplicationServerLatency `json:"applicationServerLatency,omitempty"` // Latency of application server
	ClientNetworkLatency     *ResponseApplicationsApplicationsV1ResponseClientNetworkLatency     `json:"clientNetworkLatency,omitempty"`     // Latency of client network
	ServerNetworkLatency     *ResponseApplicationsApplicationsV1ResponseServerNetworkLatency     `json:"serverNetworkLatency,omitempty"`     // Latency of server network
	ExporterIPAddress        string                                                              `json:"exporterIpAddress,omitempty"`        // Ip address of exporter device
	ExporterName             string                                                              `json:"exporterName,omitempty"`             // Name of exporter device
	ExporterUUID             string                                                              `json:"exporterUUID,omitempty"`             // UUID of exporter device
	ExporterFamily           string                                                              `json:"exporterFamily,omitempty"`           // Devices family of exporter device
	ClientName               string                                                              `json:"clientName,omitempty"`               // Endpoint client name
	ClientIP                 string                                                              `json:"clientIp,omitempty"`                 // Endpoint client ip
	Location                 string                                                              `json:"location,omitempty"`                 // Site location
	OperatingSystem          string                                                              `json:"operatingSystem,omitempty"`          // Endpoint's operating system
	DeviceType               string                                                              `json:"deviceType,omitempty"`               // Type of device
	ClientMacAddress         string                                                              `json:"clientMacAddress,omitempty"`         // Endpoint mac address
	IssueID                  string                                                              `json:"issueId,omitempty"`                  // Id number of issue
	IssueName                string                                                              `json:"issueName,omitempty"`                // Name of issue
	Application              string                                                              `json:"application,omitempty"`              // Issue reltaed application
	Severity                 string                                                              `json:"severity,omitempty"`                 // Issue severity
	Summary                  string                                                              `json:"summary,omitempty"`                  // Issue summary
	RootCause                string                                                              `json:"rootCause,omitempty"`                // Issue's root cause
	Timestamp                *int                                                                `json:"timestamp,omitempty"`                // Issue's timestamp
	Occurrences              *int                                                                `json:"occurrences,omitempty"`              // Issue occurrences
	Priority                 string                                                              `json:"priority,omitempty"`                 // Issue priority
}
type ResponseApplicationsApplicationsV1ResponsePacketLossPercent interface{}
type ResponseApplicationsApplicationsV1ResponseNetworkLatency interface{}
type ResponseApplicationsApplicationsV1ResponseJitter interface{}
type ResponseApplicationsApplicationsV1ResponseApplicationServerLatency interface{}
type ResponseApplicationsApplicationsV1ResponseClientNetworkLatency interface{}
type ResponseApplicationsApplicationsV1ResponseServerNetworkLatency interface{}
type RequestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1 struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	SiteIDs []string `json:"siteIds,omitempty"` // Site Ids

	TrendInterval string `json:"trendInterval,omitempty"` // Trend Interval

	GroupBy []string `json:"groupBy,omitempty"` // Group By

	Attributes []string `json:"attributes,omitempty"` // Attributes

	Filters *[]RequestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1Filters `json:"filters,omitempty"` //

	AggregateAttributes *[]RequestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1AggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1Page `json:"page,omitempty"` //
}
type RequestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1Filters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	Value *int `json:"value,omitempty"` // Value
}
type RequestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1AggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Cursor string `json:"cursor,omitempty"` // Cursor

	TimeSortOrder string `json:"timeSortOrder,omitempty"` // Time Sort Order
}

//RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1 Retrieves the list of network applications along with experience and health metrics - 0d8c-6945-49f8-845d
/* Retrieves the list of network applications along with experience and health metrics. If startTime and endTime are not provided, the API defaults to the last 24 hours. `siteId` is mandatory. `siteId` must be a site UUID of a building. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-NetworkApplications-1.0.0-resolved.yaml


@param RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1HeaderParams Custom header parameters
@param RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-list-of-network-applications-along-with-experience-and-health-metrics
*/
func (s *ApplicationsService) RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1(RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1HeaderParams *RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1HeaderParams, RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1QueryParams *RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1QueryParams) (*ResponseApplicationsRetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1, *resty.Response, error) {
	path := "/dna/data/api/v1/networkApplications"

	queryString, _ := query.Values(RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1HeaderParams != nil {

		if RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplicationsRetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1(RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1HeaderParams, RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1")
	}

	result := response.Result().(*ResponseApplicationsRetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1)
	return result, response, err

}

//RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringV1 Retrieves the total count of network applications by applying basic filtering - 90b0-5a40-422a-8446
/* Retrieves the number of network applications by applying basic filtering. If startTime and endTime are not provided, the API defaults to the last 24 hours. `siteId` is mandatory. `siteId` must be a site UUID of a building. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-NetworkApplications-1.0.0-resolved.yaml


@param RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringV1HeaderParams Custom header parameters
@param RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-total-count-of-network-applications-by-applying-basic-filtering
*/
func (s *ApplicationsService) RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringV1(RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringV1HeaderParams *RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringV1HeaderParams, RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringV1QueryParams *RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringV1QueryParams) (*ResponseApplicationsRetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringV1, *resty.Response, error) {
	path := "/dna/data/api/v1/networkApplications/count"

	queryString, _ := query.Values(RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringV1HeaderParams != nil {

		if RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplicationsRetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringV1(RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringV1HeaderParams, RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringV1")
	}

	result := response.Result().(*ResponseApplicationsRetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringV1)
	return result, response, err

}

//ApplicationsV1 Applications - 2db5-8a1f-4fea-9242
/* Intent API to get a list of applications for a specific site, a device, or a client device's MAC address. For a combination of a specific application with site and/or device the API gets list of issues/devices/endpoints.


@param ApplicationsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!applications
*/
func (s *ApplicationsService) ApplicationsV1(ApplicationsV1QueryParams *ApplicationsV1QueryParams) (*ResponseApplicationsApplicationsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/application-health"

	queryString, _ := query.Values(ApplicationsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplicationsApplicationsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ApplicationsV1(ApplicationsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ApplicationsV1")
	}

	result := response.Result().(*ResponseApplicationsApplicationsV1)
	return result, response, err

}

//RetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1 Retrieves the Trend analytics data related to network applications. - 22b6-58d5-41ba-ac1a
/* Retrieves the trend analytics of applications experience data for the specified time range. The data will be grouped based on the given trend time interval. This API facilitates obtaining consolidated insights into the performance and status of the network applications over the specified start and end time. If startTime and endTime are not provided, the API defaults to the last 24 hours. `siteId` and `trendInterval` are mandatory. `siteId` must be a site UUID of a building. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-NetworkApplications-1.0.0-resolved.yaml


@param RetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-trend-analytics-data-related-to-network-applications
*/
func (s *ApplicationsService) RetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1(requestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1 *RequestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1, RetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1HeaderParams *RetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1HeaderParams) (*ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1, *resty.Response, error) {
	path := "/dna/data/api/v1/networkApplications/trendAnalytics"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1HeaderParams != nil {

		if RetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1).
		SetResult(&ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1(requestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1, RetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation RetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1")
	}

	result := response.Result().(*ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1)
	return result, response, err

}

// Alias Function
/*
This method acts as an alias for the method `RetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1`
*/
func (s *ApplicationsService) RetrievesTheTrendAnalyticsDataRelatedToNetworkApplications(requestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1 *RequestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1, RetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1HeaderParams *RetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1HeaderParams) (*ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1, *resty.Response, error) {
	return s.RetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1(requestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1, RetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1`
*/
func (s *ApplicationsService) RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetrics(RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1HeaderParams *RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1HeaderParams, RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1QueryParams *RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1QueryParams) (*ResponseApplicationsRetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1, *resty.Response, error) {
	return s.RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1(RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1HeaderParams, RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringV1`
*/
func (s *ApplicationsService) RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFiltering(RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringV1HeaderParams *RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringV1HeaderParams, RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringV1QueryParams *RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringV1QueryParams) (*ResponseApplicationsRetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringV1, *resty.Response, error) {
	return s.RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringV1(RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringV1HeaderParams, RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `ApplicationsV1`
*/
func (s *ApplicationsService) Applications(ApplicationsV1QueryParams *ApplicationsV1QueryParams) (*ResponseApplicationsApplicationsV1, *resty.Response, error) {
	return s.ApplicationsV1(ApplicationsV1QueryParams)
}
