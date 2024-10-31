package catalyst

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type AIEndpointAnalyticsService service

type GetAIEndpointAnalyticsAttributeDictionariesV1QueryParams struct {
	IncludeAttributes bool `url:"includeAttributes,omitempty"` //Flag to indicate whether attribute list for each dictionary should be included in response.
}
type QueryTheEndpointsV1QueryParams struct {
	ProfilingStatus          string   `url:"profilingStatus,omitempty"`          //Profiling status of the endpoint. Possible values are 'profiled', 'partialProfiled', 'notProfiled'.
	MacAddress               string   `url:"macAddress,omitempty"`               //MAC address to search for. Partial string is allowed.
	MacAddresses             []string `url:"macAddresses,omitempty"`             //List of MAC addresses to filter on. Only exact matches will be returned.
	IP                       string   `url:"ip,omitempty"`                       //IP address to search for. Partial string is allowed.
	DeviceType               string   `url:"deviceType,omitempty"`               //Type of device to search for. Partial string is allowed.
	HardwareManufacturer     string   `url:"hardwareManufacturer,omitempty"`     //Hardware manufacturer to search for. Partial string is allowed.
	HardwareModel            string   `url:"hardwareModel,omitempty"`            //Hardware model to search for. Partial string is allowed.
	OperatingSystem          string   `url:"operatingSystem,omitempty"`          //Operating system to search for. Partial string is allowed.
	Registered               bool     `url:"registered,omitempty"`               //Flag to fetch manually registered or non-registered endpoints.
	RandomMac                bool     `url:"randomMac,omitempty"`                //Flag to fetch endpoints having randomized MAC or not.
	TrustScore               string   `url:"trustScore,omitempty"`               //Overall trust score of the endpoint. It can be provided either as a number value (e.g. 5), or as a range (e.g. 3-7). Provide value as '-' if you want to search for all endpoints where trust score is not assigned.
	AuthMethod               string   `url:"authMethod,omitempty"`               //Authentication method. Partial string is allowed.
	PostureStatus            string   `url:"postureStatus,omitempty"`            //Posture status.
	AiSpoofingTrustLevel     string   `url:"aiSpoofingTrustLevel,omitempty"`     //Trust level of the endpoint due to AI spoofing. Possible values are 'low', 'medium', 'high'.
	ChangedProfileTrustLevel string   `url:"changedProfileTrustLevel,omitempty"` //Trust level of the endpoint due to changing profile labels. Possible values are 'low', 'medium', 'high'.
	NatTrustLevel            string   `url:"natTrustLevel,omitempty"`            //Trust level of the endpoint due to NAT access. Possible values are 'low', 'medium', 'high'.
	ConcurrentMacTrustLevel  string   `url:"concurrentMacTrustLevel,omitempty"`  //Trust level of the endpoint due to concurrent MAC address. Possible values are 'low', 'medium', 'high'.
	IPBlocklistDetected      bool     `url:"ipBlocklistDetected,omitempty"`      //Flag to fetch endpoints hitting IP blocklist or not.
	UnauthPortDetected       bool     `url:"unauthPortDetected,omitempty"`       //Flag to fetch endpoints exposing unauthorized ports or not.
	WeakCredDetected         bool     `url:"weakCredDetected,omitempty"`         //Flag to fetch endpoints having weak credentials or not.
	AncPolicy                string   `url:"ancPolicy,omitempty"`                //ANC policy. Only exact match will be returned.
	Limit                    int      `url:"limit,omitempty"`                    //Maximum number of records to be fetched. If not provided, 50 records will be fetched by default. Maximum 1000 records can be fetched at a time. Use pagination if more records need to be fetched.
	Offset                   int      `url:"offset,omitempty"`                   //Record offset to start data fetch at. Offset starts at zero.
	SortBy                   string   `url:"sortBy,omitempty"`                   //Name of the column to sort the results on. Please note that fetch might take more time if sorting is requested. Possible values are 'macAddress', 'ip'.
	Order                    string   `url:"order,omitempty"`                    //Order to be used for sorting. Possible values are 'asc', 'desc'.
	Include                  string   `url:"include,omitempty"`                  //The datasets that should be included in the response. By default, value of this parameter is blank, and the response will include only basic details of the endpoint. To include other datasets or dictionaries, send comma separated list of following values: 'ALL' - Include all attributes. 'CDP', 'DHCP', etc. - Include attributes from given dictionaries. To get full list of dictionaries, use corresponding GET API. 'ANC' - Include ANC policy related details. 'TRUST' - Include trust score details.
}
type FetchTheCountOfEndpointsV1QueryParams struct {
	ProfilingStatus          string   `url:"profilingStatus,omitempty"`          //Profiling status of the endpoint. Possible values are 'profiled', 'partialProfiled', 'notProfiled'.
	MacAddress               string   `url:"macAddress,omitempty"`               //MAC address to search for. Partial string is allowed.
	MacAddresses             []string `url:"macAddresses,omitempty"`             //List of MAC addresses to filter on. Only exact matches will be returned.
	IP                       string   `url:"ip,omitempty"`                       //IP address to search for. Partial string is allowed.
	DeviceType               string   `url:"deviceType,omitempty"`               //Type of device to search for. Partial string is allowed.
	HardwareManufacturer     string   `url:"hardwareManufacturer,omitempty"`     //Hardware manufacturer to search for. Partial string is allowed.
	HardwareModel            string   `url:"hardwareModel,omitempty"`            //Hardware model to search for. Partial string is allowed.
	OperatingSystem          string   `url:"operatingSystem,omitempty"`          //Operating system to search for. Partial string is allowed.
	Registered               bool     `url:"registered,omitempty"`               //Flag to fetch manually registered or non-registered endpoints.
	RandomMac                bool     `url:"randomMac,omitempty"`                //Flag to fetch endpoints having randomized MAC or not.
	TrustScore               string   `url:"trustScore,omitempty"`               //Overall trust score of the endpoint. It can be provided either as a number value (e.g. 5), or as a range (e.g. 3-7). Provide value as '-' if you want to search for all endpoints where trust score is not assigned.
	AuthMethod               string   `url:"authMethod,omitempty"`               //Authentication method. Partial string is allowed.
	PostureStatus            string   `url:"postureStatus,omitempty"`            //Posture status.
	AiSpoofingTrustLevel     string   `url:"aiSpoofingTrustLevel,omitempty"`     //Trust level of the endpoint due to AI spoofing. Possible values are 'low', 'medium', 'high'.
	ChangedProfileTrustLevel string   `url:"changedProfileTrustLevel,omitempty"` //Trust level of the endpoint due to changing profile labels. Possible values are 'low', 'medium', 'high'.
	NatTrustLevel            string   `url:"natTrustLevel,omitempty"`            //Trust level of the endpoint due to NAT access. Possible values are 'low', 'medium', 'high'.
	ConcurrentMacTrustLevel  string   `url:"concurrentMacTrustLevel,omitempty"`  //Trust level of the endpoint due to concurrent MAC address. Possible values are 'low', 'medium', 'high'.
	IPBlocklistDetected      bool     `url:"ipBlocklistDetected,omitempty"`      //Flag to fetch endpoints hitting IP blocklist or not.
	UnauthPortDetected       bool     `url:"unauthPortDetected,omitempty"`       //Flag to fetch endpoints exposing unauthorized ports or not.
	WeakCredDetected         bool     `url:"weakCredDetected,omitempty"`         //Flag to fetch endpoints having weak credentials or not.
	AncPolicy                string   `url:"ancPolicy,omitempty"`                //ANC policy. Only exact match will be returned.
}
type GetEndpointDetailsV1QueryParams struct {
	Include string `url:"include,omitempty"` //The datasets that should be included in the response. By default, value of this parameter is blank, and the response will include only basic details of the endpoint. To include other datasets or dictionaries, send comma separated list of following values: 'ALL' - Include all attributes. 'CDP', 'DHCP', etc. - Include attributes from given dictionaries. To get full list of dictionaries, use corresponding GET API. 'ANC' - Include ANC policy related details. 'TRUST' - Include trust score details.
}
type GetListOfProfilingRulesV1QueryParams struct {
	RuleType       string  `url:"ruleType,omitempty"`       //Use comma-separated list of rule types to filter the data. Defaults to 'Custom Rule'.
	IncludeDeleted bool    `url:"includeDeleted,omitempty"` //Flag to indicate whether deleted rules should be part of the records fetched.
	Limit          float64 `url:"limit,omitempty"`          //Maximum number of records to be fetched. If not provided, 500 records will be fetched by default. To fetch all the records in the system, provide a large value for this parameter.
	Offset         float64 `url:"offset,omitempty"`         //Record offset to start data fetch at. Offset starts at zero.
	SortBy         string  `url:"sortBy,omitempty"`         //Name of the column to sort the results on. Please note that fetch might take more time if sorting is requested.
	Order          string  `url:"order,omitempty"`          //Order to be used for sorting.
}
type GetCountOfProfilingRulesV1QueryParams struct {
	RuleType       string `url:"ruleType,omitempty"`       //Use comma-separated list of rule types to filter the data. Defaults to 'Custom Rule'.
	IncludeDeleted bool   `url:"includeDeleted,omitempty"` //Flag to indicate whether deleted rules should be part of the records fetched.
}

type ResponseAIEndpointAnalyticsGetAncPoliciesV1 []ResponseItemAIEndpointAnalyticsGetAncPoliciesV1 // Array of ResponseAIEndpointAnalyticsGetANCPoliciesV1
type ResponseItemAIEndpointAnalyticsGetAncPoliciesV1 struct {
	Name string `json:"name,omitempty"` // Name of the ANC policy.
}
type ResponseAIEndpointAnalyticsGetAIEndpointAnalyticsAttributeDictionariesV1 []ResponseItemAIEndpointAnalyticsGetAIEndpointAnalyticsAttributeDictionariesV1 // Array of ResponseAIEndpointAnalyticsGetAIEndpointAnalyticsAttributeDictionariesV1
type ResponseItemAIEndpointAnalyticsGetAIEndpointAnalyticsAttributeDictionariesV1 struct {
	Name        string                                                                                    `json:"name,omitempty"`        // Name of the dictionary.
	Description string                                                                                    `json:"description,omitempty"` // Description of the dictionary.
	Attributes  *[]ResponseItemAIEndpointAnalyticsGetAIEndpointAnalyticsAttributeDictionariesV1Attributes `json:"attributes,omitempty"`  //
}
type ResponseItemAIEndpointAnalyticsGetAIEndpointAnalyticsAttributeDictionariesV1Attributes struct {
	Name        string `json:"name,omitempty"`        // Name of the attribute.
	Description string `json:"description,omitempty"` // Description of the attribute.
}
type ResponseAIEndpointAnalyticsQueryTheEndpointsV1 struct {
	TotalResults   *int                                                   `json:"totalResults,omitempty"`   // Total number of records matching the given filter criteria.
	HasMoreResults *bool                                                  `json:"hasMoreResults,omitempty"` // Flag to indicate whether more results are available than what is currently in the response.
	Items          *[]ResponseAIEndpointAnalyticsQueryTheEndpointsV1Items `json:"items,omitempty"`          //
}
type ResponseAIEndpointAnalyticsQueryTheEndpointsV1Items struct {
	ID                           string                                                                  `json:"id,omitempty"`                           // Unique identifier for the endpoint.
	Duid                         string                                                                  `json:"duid,omitempty"`                         // Unique DUID.
	MacAddress                   string                                                                  `json:"macAddress,omitempty"`                   // MAC address of the endpoint.
	DeviceType                   []string                                                                `json:"deviceType,omitempty"`                   // Type of the device represented by this endpoint.
	HardwareManufacturer         []string                                                                `json:"hardwareManufacturer,omitempty"`         // Hardware manufacturer for the endpoint.
	HardwareModel                []string                                                                `json:"hardwareModel,omitempty"`                // Hardware model of the endpoint.
	OperatingSystem              []string                                                                `json:"operatingSystem,omitempty"`              // Operating system of the endpoint.
	LastProbeCollectionTimestamp *int                                                                    `json:"lastProbeCollectionTimestamp,omitempty"` // Last probe collection timestamp in epoch milliseconds.
	RandomMac                    *bool                                                                   `json:"randomMac,omitempty"`                    // Flag to indicate whether MAC address is a randomized one or not.
	Registered                   *bool                                                                   `json:"registered,omitempty"`                   // Flag to indicate whether this is a manually registered endpoint or not.
	Attributes                   *ResponseAIEndpointAnalyticsQueryTheEndpointsV1ItemsAttributes          `json:"attributes,omitempty"`                   // Various endpoint attributes grouped by dictionaries (e.g. IP, DHCP, etc).
	TrustData                    *ResponseAIEndpointAnalyticsQueryTheEndpointsV1ItemsTrustData           `json:"trustData,omitempty"`                    //
	AncPolicy                    string                                                                  `json:"ancPolicy,omitempty"`                    // ANC policy currently applied to the endpoint in ISE.
	GranularAncPolicy            *[]ResponseAIEndpointAnalyticsQueryTheEndpointsV1ItemsGranularAncPolicy `json:"granularAncPolicy,omitempty"`            //
}
type ResponseAIEndpointAnalyticsQueryTheEndpointsV1ItemsAttributes interface{}
type ResponseAIEndpointAnalyticsQueryTheEndpointsV1ItemsTrustData struct {
	TrustScore               *int   `json:"trustScore,omitempty"`               // Overall trust score of the endpoint.
	AuthMethod               string `json:"authMethod,omitempty"`               // Authentication method.
	PostureStatus            string `json:"postureStatus,omitempty"`            // Posture status.
	AiSpoofingTrustLevel     string `json:"aiSpoofingTrustLevel,omitempty"`     // Trust level of the endpoint due to AI spoofing.
	ChangedProfileTrustLevel string `json:"changedProfileTrustLevel,omitempty"` // Trust level of the endpoint due to changing profile labels.
	NatTrustLevel            string `json:"natTrustLevel,omitempty"`            // Trust level of the endpoint due to NAT access.
	ConcurrentMacTrustLevel  string `json:"concurrentMacTrustLevel,omitempty"`  // Trust level of the endpoint due to concurrent MAC address.
	IPBlocklistDetected      *bool  `json:"ipBlocklistDetected,omitempty"`      // Flag to fetch endpoints hitting IP blocklist or not.
	UnauthPortDetected       *bool  `json:"unauthPortDetected,omitempty"`       // Flag to fetch endpoints exposing unauthorized ports or not.
	WeakCredDetected         *bool  `json:"weakCredDetected,omitempty"`         // Flag to fetch endpoints having weak credentials or not.
}
type ResponseAIEndpointAnalyticsQueryTheEndpointsV1ItemsGranularAncPolicy struct {
	Name         string `json:"name,omitempty"`         // Name of the granular ANC policy.
	NasIPAddress string `json:"nasIpAddress,omitempty"` // IP address of the network device where endpoint is attached.
}
type ResponseAIEndpointAnalyticsFetchTheCountOfEndpointsV1 struct {
	Count *int `json:"count,omitempty"` //
}
type ResponseAIEndpointAnalyticsGetEndpointDetailsV1 struct {
	ID                           string                                                              `json:"id,omitempty"`                           // Unique identifier for the endpoint.
	Duid                         string                                                              `json:"duid,omitempty"`                         // Unique DUID.
	MacAddress                   string                                                              `json:"macAddress,omitempty"`                   // MAC address of the endpoint.
	DeviceType                   []string                                                            `json:"deviceType,omitempty"`                   // Type of the device represented by this endpoint.
	HardwareManufacturer         []string                                                            `json:"hardwareManufacturer,omitempty"`         // Hardware manufacturer for the endpoint.
	HardwareModel                []string                                                            `json:"hardwareModel,omitempty"`                // Hardware model of the endpoint.
	OperatingSystem              []string                                                            `json:"operatingSystem,omitempty"`              // Operating system of the endpoint.
	LastProbeCollectionTimestamp *int                                                                `json:"lastProbeCollectionTimestamp,omitempty"` // Last probe collection timestamp in epoch milliseconds.
	RandomMac                    *bool                                                               `json:"randomMac,omitempty"`                    // Flag to indicate whether MAC address is a randomized one or not.
	Registered                   *bool                                                               `json:"registered,omitempty"`                   // Flag to indicate whether this is a manually registered endpoint or not.
	Attributes                   *ResponseAIEndpointAnalyticsGetEndpointDetailsV1Attributes          `json:"attributes,omitempty"`                   // Various endpoint attributes grouped by dictionaries (e.g. IP, DHCP, etc).
	TrustData                    *ResponseAIEndpointAnalyticsGetEndpointDetailsV1TrustData           `json:"trustData,omitempty"`                    //
	AncPolicy                    string                                                              `json:"ancPolicy,omitempty"`                    // ANC policy currently applied to the endpoint in ISE.
	GranularAncPolicy            *[]ResponseAIEndpointAnalyticsGetEndpointDetailsV1GranularAncPolicy `json:"granularAncPolicy,omitempty"`            //
}
type ResponseAIEndpointAnalyticsGetEndpointDetailsV1Attributes interface{}
type ResponseAIEndpointAnalyticsGetEndpointDetailsV1TrustData struct {
	TrustScore               *int   `json:"trustScore,omitempty"`               // Overall trust score of the endpoint.
	AuthMethod               string `json:"authMethod,omitempty"`               // Authentication method.
	PostureStatus            string `json:"postureStatus,omitempty"`            // Posture status.
	AiSpoofingTrustLevel     string `json:"aiSpoofingTrustLevel,omitempty"`     // Trust level of the endpoint due to AI spoofing.
	ChangedProfileTrustLevel string `json:"changedProfileTrustLevel,omitempty"` // Trust level of the endpoint due to changing profile labels.
	NatTrustLevel            string `json:"natTrustLevel,omitempty"`            // Trust level of the endpoint due to NAT access.
	ConcurrentMacTrustLevel  string `json:"concurrentMacTrustLevel,omitempty"`  // Trust level of the endpoint due to concurrent MAC address.
	IPBlocklistDetected      *bool  `json:"ipBlocklistDetected,omitempty"`      // Flag to fetch endpoints hitting IP blocklist or not.
	UnauthPortDetected       *bool  `json:"unauthPortDetected,omitempty"`       // Flag to fetch endpoints exposing unauthorized ports or not.
	WeakCredDetected         *bool  `json:"weakCredDetected,omitempty"`         // Flag to fetch endpoints having weak credentials or not.
}
type ResponseAIEndpointAnalyticsGetEndpointDetailsV1GranularAncPolicy struct {
	Name         string `json:"name,omitempty"`         // Name of the granular ANC policy.
	NasIPAddress string `json:"nasIpAddress,omitempty"` // IP address of the network device where endpoint is attached.
}
type ResponseAIEndpointAnalyticsCreateAProfilingRuleV1 struct {
	ID   string `json:"id,omitempty"`   // Unique identifier for the newly created resource.
	Link string `json:"link,omitempty"` // Link to the newly created resource.
}
type ResponseAIEndpointAnalyticsGetListOfProfilingRulesV1 struct {
	ProfilingRules *[]ResponseAIEndpointAnalyticsGetListOfProfilingRulesV1ProfilingRules `json:"profilingRules,omitempty"` //
}
type ResponseAIEndpointAnalyticsGetListOfProfilingRulesV1ProfilingRules struct {
	RuleID          string                                                                             `json:"ruleId,omitempty"`          // Unique identifier for the rule. This is normally generated by the system, and client does not need to provide it for rules that need to be newly created.
	RuleName        string                                                                             `json:"ruleName,omitempty"`        // Human readable name for the rule.
	RuleType        string                                                                             `json:"ruleType,omitempty"`        // Type of the rule.
	RuleVersion     *int                                                                               `json:"ruleVersion,omitempty"`     // Version of the rule.
	RulePriority    *int                                                                               `json:"rulePriority,omitempty"`    // Priority for the rule.
	SourcePriority  *int                                                                               `json:"sourcePriority,omitempty"`  // Source priority for the rule.
	IsDeleted       *bool                                                                              `json:"isDeleted,omitempty"`       // Flag to indicate whether the rule was deleted.
	LastModifiedBy  string                                                                             `json:"lastModifiedBy,omitempty"`  // User that last modified the rule. It is read-only, and is ignored if provided as part of input request.
	LastModifiedOn  *int                                                                               `json:"lastModifiedOn,omitempty"`  // Timestamp (in epoch milliseconds) of last modification. It is read-only, and is ignored if provided as part of input request.
	PluginID        string                                                                             `json:"pluginId,omitempty"`        // Plugin for the rule. Only applicable for 'Cisco Default' rules.
	ClusterID       string                                                                             `json:"clusterId,omitempty"`       // Unique identifier for ML cluster. Only applicable for 'ML Rule'.
	Rejected        *bool                                                                              `json:"rejected,omitempty"`        // Flag to indicate whether rule has been rejected by user or not. Only applicable for 'ML Rule'.
	Result          *ResponseAIEndpointAnalyticsGetListOfProfilingRulesV1ProfilingRulesResult          `json:"result,omitempty"`          //
	ConditionGroups *ResponseAIEndpointAnalyticsGetListOfProfilingRulesV1ProfilingRulesConditionGroups `json:"conditionGroups,omitempty"` //
	UsedAttributes  []string                                                                           `json:"usedAttributes,omitempty"`  // List of attributes used in the rule. Only applicable for 'Cisco Default' rules.
}
type ResponseAIEndpointAnalyticsGetListOfProfilingRulesV1ProfilingRulesResult struct {
	DeviceType           []string `json:"deviceType,omitempty"`           // List of device types determined by the current rule.
	HardwareManufacturer []string `json:"hardwareManufacturer,omitempty"` // List of hardware manufacturers determined by the current rule.
	HardwareModel        []string `json:"hardwareModel,omitempty"`        // List of hardware models determined by the current rule.
	OperatingSystem      []string `json:"operatingSystem,omitempty"`      // List of operating systems determined by the current rule.
}
type ResponseAIEndpointAnalyticsGetListOfProfilingRulesV1ProfilingRulesConditionGroups struct {
	Type           string                                                                                      `json:"type,omitempty"`           //
	Condition      *ResponseAIEndpointAnalyticsGetListOfProfilingRulesV1ProfilingRulesConditionGroupsCondition `json:"condition,omitempty"`      //
	Operator       string                                                                                      `json:"operator,omitempty"`       //
	ConditionGroup []string                                                                                    `json:"conditionGroup,omitempty"` //
}
type ResponseAIEndpointAnalyticsGetListOfProfilingRulesV1ProfilingRulesConditionGroupsCondition struct {
	Attribute           string `json:"attribute,omitempty"`           //
	Operator            string `json:"operator,omitempty"`            //
	Value               string `json:"value,omitempty"`               //
	AttributeDictionary string `json:"attributeDictionary,omitempty"` //
}
type ResponseAIEndpointAnalyticsGetCountOfProfilingRulesV1 struct {
	Count *int `json:"count,omitempty"` //
}
type ResponseAIEndpointAnalyticsGetDetailsOfASingleProfilingRuleV1 struct {
	RuleID          string                                                                        `json:"ruleId,omitempty"`          // Unique identifier for the rule. This is normally generated by the system, and client does not need to provide it for rules that need to be newly created.
	RuleName        string                                                                        `json:"ruleName,omitempty"`        // Human readable name for the rule.
	RuleType        string                                                                        `json:"ruleType,omitempty"`        // Type of the rule.
	RuleVersion     *int                                                                          `json:"ruleVersion,omitempty"`     // Version of the rule.
	RulePriority    *int                                                                          `json:"rulePriority,omitempty"`    // Priority for the rule.
	SourcePriority  *int                                                                          `json:"sourcePriority,omitempty"`  // Source priority for the rule.
	IsDeleted       *bool                                                                         `json:"isDeleted,omitempty"`       // Flag to indicate whether the rule was deleted.
	LastModifiedBy  string                                                                        `json:"lastModifiedBy,omitempty"`  // User that last modified the rule. It is read-only, and is ignored if provided as part of input request.
	LastModifiedOn  *int                                                                          `json:"lastModifiedOn,omitempty"`  // Timestamp (in epoch milliseconds) of last modification. It is read-only, and is ignored if provided as part of input request.
	PluginID        string                                                                        `json:"pluginId,omitempty"`        // Plugin for the rule. Only applicable for 'Cisco Default' rules.
	ClusterID       string                                                                        `json:"clusterId,omitempty"`       // Unique identifier for ML cluster. Only applicable for 'ML Rule'.
	Rejected        *bool                                                                         `json:"rejected,omitempty"`        // Flag to indicate whether rule has been rejected by user or not. Only applicable for 'ML Rule'.
	Result          *ResponseAIEndpointAnalyticsGetDetailsOfASingleProfilingRuleV1Result          `json:"result,omitempty"`          //
	ConditionGroups *ResponseAIEndpointAnalyticsGetDetailsOfASingleProfilingRuleV1ConditionGroups `json:"conditionGroups,omitempty"` //
	UsedAttributes  []string                                                                      `json:"usedAttributes,omitempty"`  // List of attributes used in the rule. Only applicable for 'Cisco Default' rules.
}
type ResponseAIEndpointAnalyticsGetDetailsOfASingleProfilingRuleV1Result struct {
	DeviceType           []string `json:"deviceType,omitempty"`           // List of device types determined by the current rule.
	HardwareManufacturer []string `json:"hardwareManufacturer,omitempty"` // List of hardware manufacturers determined by the current rule.
	HardwareModel        []string `json:"hardwareModel,omitempty"`        // List of hardware models determined by the current rule.
	OperatingSystem      []string `json:"operatingSystem,omitempty"`      // List of operating systems determined by the current rule.
}
type ResponseAIEndpointAnalyticsGetDetailsOfASingleProfilingRuleV1ConditionGroups struct {
	Type           string                                                                                 `json:"type,omitempty"`           //
	Condition      *ResponseAIEndpointAnalyticsGetDetailsOfASingleProfilingRuleV1ConditionGroupsCondition `json:"condition,omitempty"`      //
	Operator       string                                                                                 `json:"operator,omitempty"`       //
	ConditionGroup []string                                                                               `json:"conditionGroup,omitempty"` //
}
type ResponseAIEndpointAnalyticsGetDetailsOfASingleProfilingRuleV1ConditionGroupsCondition struct {
	Attribute           string `json:"attribute,omitempty"`           //
	Operator            string `json:"operator,omitempty"`            //
	Value               string `json:"value,omitempty"`               //
	AttributeDictionary string `json:"attributeDictionary,omitempty"` //
}
type ResponseAIEndpointAnalyticsGetTaskDetailsV1 struct {
	ID             string                                                     `json:"id,omitempty"`             // Unique identifier for the task.
	Name           string                                                     `json:"name,omitempty"`           // Name of the task.
	Status         string                                                     `json:"status,omitempty"`         // Status of the task.
	Errors         *[]ResponseAIEndpointAnalyticsGetTaskDetailsV1Errors       `json:"errors,omitempty"`         //
	AdditionalInfo *ResponseAIEndpointAnalyticsGetTaskDetailsV1AdditionalInfo `json:"additionalInfo,omitempty"` // Additional information about the task.
	CreatedBy      string                                                     `json:"createdBy,omitempty"`      // Name of the user that created the task.
	CreatedOn      *int                                                       `json:"createdOn,omitempty"`      // Task creation timestamp in epoch milliseconds.
	LastUpdatedOn  *int                                                       `json:"lastUpdatedOn,omitempty"`  // Last update timestamp in epoch milliseconds.
}
type ResponseAIEndpointAnalyticsGetTaskDetailsV1Errors struct {
	Index   *int   `json:"index,omitempty"`   // Index of the input records which had error during processing. In case the input is not an array, or the error is not record specific, this will be -1.
	Code    *int   `json:"code,omitempty"`    // Error code.
	Message string `json:"message,omitempty"` // Error message.
	Details string `json:"details,omitempty"` // Optional details about the error.
}
type ResponseAIEndpointAnalyticsGetTaskDetailsV1AdditionalInfo interface{}
type RequestAIEndpointAnalyticsProcessCmdbEndpointsV1 []RequestItemAIEndpointAnalyticsProcessCmdbEndpointsV1 // Array of RequestAIEndpointAnalyticsProcessCMDBEndpointsV1
type RequestItemAIEndpointAnalyticsProcessCmdbEndpointsV1 struct {
	MacAddress          string `json:"macAddress,omitempty"`          // MAC address of the endpoint.
	SerialNumber        string `json:"serialNumber,omitempty"`        // Serial number of the endpoint.
	AssetTag            string `json:"assetTag,omitempty"`            // Asset tag.
	ModelCategory       string `json:"modelCategory,omitempty"`       // Category of the model.
	Model               string `json:"model,omitempty"`               // Asset model.
	DisplayName         string `json:"displayName,omitempty"`         // Display name of the asset.
	Department          string `json:"department,omitempty"`          // Department that asset belongs to.
	Location            string `json:"location,omitempty"`            // Location of the asset.
	ManagedBy           string `json:"managedBy,omitempty"`           // Asset managed by.
	LastUpdateTimestamp *int   `json:"lastUpdateTimestamp,omitempty"` // Last update timestamp in epoch milliseconds.
}
type RequestAIEndpointAnalyticsRegisterAnEndpointV1 struct {
	MacAddress           string `json:"macAddress,omitempty"`           // MAC address of the endpoint.
	DeviceType           string `json:"deviceType,omitempty"`           // Type of the device represented by this endpoint.
	HardwareManufacturer string `json:"hardwareManufacturer,omitempty"` // Hardware manufacturer for the endpoint.
	HardwareModel        string `json:"hardwareModel,omitempty"`        // Hardware model of the endpoint.
}
type RequestAIEndpointAnalyticsUpdateARegisteredEndpointV1 struct {
	DeviceType           string `json:"deviceType,omitempty"`           // Type of the device represented by this endpoint.
	HardwareManufacturer string `json:"hardwareManufacturer,omitempty"` // Hardware manufacturer for the endpoint.
	HardwareModel        string `json:"hardwareModel,omitempty"`        // Hardware model of the endpoint.
}
type RequestAIEndpointAnalyticsApplyAncPolicyV1 struct {
	AncPolicy         string                                                         `json:"ancPolicy,omitempty"`         // ANC policy name.
	GranularAncPolicy *[]RequestAIEndpointAnalyticsApplyAncPolicyV1GranularAncPolicy `json:"granularAncPolicy,omitempty"` //
}
type RequestAIEndpointAnalyticsApplyAncPolicyV1GranularAncPolicy struct {
	Name         string `json:"name,omitempty"`         // Name of the granular ANC policy.
	NasIPAddress string `json:"nasIpAddress,omitempty"` // IP address of the network device where endpoint is attached.
}
type RequestAIEndpointAnalyticsCreateAProfilingRuleV1 struct {
	RuleID          string                                                           `json:"ruleId,omitempty"`          // Unique identifier for the rule. This is normally generated by the system, and client does not need to provide it for rules that need to be newly created.
	RuleName        string                                                           `json:"ruleName,omitempty"`        // Human readable name for the rule.
	RuleType        string                                                           `json:"ruleType,omitempty"`        // Type of the rule.
	RuleVersion     *int                                                             `json:"ruleVersion,omitempty"`     // Version of the rule.
	RulePriority    *int                                                             `json:"rulePriority,omitempty"`    // Priority for the rule.
	SourcePriority  *int                                                             `json:"sourcePriority,omitempty"`  // Source priority for the rule.
	IsDeleted       *bool                                                            `json:"isDeleted,omitempty"`       // Flag to indicate whether the rule was deleted.
	LastModifiedBy  string                                                           `json:"lastModifiedBy,omitempty"`  // User that last modified the rule. It is read-only, and is ignored if provided as part of input request.
	LastModifiedOn  *int                                                             `json:"lastModifiedOn,omitempty"`  // Timestamp (in epoch milliseconds) of last modification. It is read-only, and is ignored if provided as part of input request.
	PluginID        string                                                           `json:"pluginId,omitempty"`        // Plugin for the rule. Only applicable for 'Cisco Default' rules.
	ClusterID       string                                                           `json:"clusterId,omitempty"`       // Unique identifier for ML cluster. Only applicable for 'ML Rule'.
	Rejected        *bool                                                            `json:"rejected,omitempty"`        // Flag to indicate whether rule has been rejected by user or not. Only applicable for 'ML Rule'.
	Result          *RequestAIEndpointAnalyticsCreateAProfilingRuleV1Result          `json:"result,omitempty"`          //
	ConditionGroups *RequestAIEndpointAnalyticsCreateAProfilingRuleV1ConditionGroups `json:"conditionGroups,omitempty"` //
	UsedAttributes  []string                                                         `json:"usedAttributes,omitempty"`  // List of attributes used in the rule. Only applicable for 'Cisco Default' rules.
}
type RequestAIEndpointAnalyticsCreateAProfilingRuleV1Result struct {
	DeviceType           []string `json:"deviceType,omitempty"`           // List of device types determined by the current rule.
	HardwareManufacturer []string `json:"hardwareManufacturer,omitempty"` // List of hardware manufacturers determined by the current rule.
	HardwareModel        []string `json:"hardwareModel,omitempty"`        // List of hardware models determined by the current rule.
	OperatingSystem      []string `json:"operatingSystem,omitempty"`      // List of operating systems determined by the current rule.
}
type RequestAIEndpointAnalyticsCreateAProfilingRuleV1ConditionGroups struct {
	Type           string                                                                    `json:"type,omitempty"`           //
	Condition      *RequestAIEndpointAnalyticsCreateAProfilingRuleV1ConditionGroupsCondition `json:"condition,omitempty"`      //
	Operator       string                                                                    `json:"operator,omitempty"`       //
	ConditionGroup []string                                                                  `json:"conditionGroup,omitempty"` //
}
type RequestAIEndpointAnalyticsCreateAProfilingRuleV1ConditionGroupsCondition struct {
	Attribute           string `json:"attribute,omitempty"`           //
	Operator            string `json:"operator,omitempty"`            //
	Value               string `json:"value,omitempty"`               //
	AttributeDictionary string `json:"attributeDictionary,omitempty"` //
}
type RequestAIEndpointAnalyticsImportProfilingRulesInBulkV1 struct {
	ProfilingRules *[]RequestAIEndpointAnalyticsImportProfilingRulesInBulkV1ProfilingRules `json:"profilingRules,omitempty"` //
}
type RequestAIEndpointAnalyticsImportProfilingRulesInBulkV1ProfilingRules struct {
	RuleID          string                                                                               `json:"ruleId,omitempty"`          // Unique identifier for the rule. This is normally generated by the system, and client does not need to provide it for rules that need to be newly created.
	RuleName        string                                                                               `json:"ruleName,omitempty"`        // Human readable name for the rule.
	RuleType        string                                                                               `json:"ruleType,omitempty"`        // Type of the rule.
	RuleVersion     *int                                                                                 `json:"ruleVersion,omitempty"`     // Version of the rule.
	RulePriority    *int                                                                                 `json:"rulePriority,omitempty"`    // Priority for the rule.
	SourcePriority  *int                                                                                 `json:"sourcePriority,omitempty"`  // Source priority for the rule.
	IsDeleted       *bool                                                                                `json:"isDeleted,omitempty"`       // Flag to indicate whether the rule was deleted.
	LastModifiedBy  string                                                                               `json:"lastModifiedBy,omitempty"`  // User that last modified the rule. It is read-only, and is ignored if provided as part of input request.
	LastModifiedOn  *int                                                                                 `json:"lastModifiedOn,omitempty"`  // Timestamp (in epoch milliseconds) of last modification. It is read-only, and is ignored if provided as part of input request.
	PluginID        string                                                                               `json:"pluginId,omitempty"`        // Plugin for the rule. Only applicable for 'Cisco Default' rules.
	ClusterID       string                                                                               `json:"clusterId,omitempty"`       // Unique identifier for ML cluster. Only applicable for 'ML Rule'.
	Rejected        *bool                                                                                `json:"rejected,omitempty"`        // Flag to indicate whether rule has been rejected by user or not. Only applicable for 'ML Rule'.
	Result          *RequestAIEndpointAnalyticsImportProfilingRulesInBulkV1ProfilingRulesResult          `json:"result,omitempty"`          //
	ConditionGroups *RequestAIEndpointAnalyticsImportProfilingRulesInBulkV1ProfilingRulesConditionGroups `json:"conditionGroups,omitempty"` //
	UsedAttributes  []string                                                                             `json:"usedAttributes,omitempty"`  // List of attributes used in the rule. Only applicable for 'Cisco Default' rules.
}
type RequestAIEndpointAnalyticsImportProfilingRulesInBulkV1ProfilingRulesResult struct {
	DeviceType           []string `json:"deviceType,omitempty"`           // List of device types determined by the current rule.
	HardwareManufacturer []string `json:"hardwareManufacturer,omitempty"` // List of hardware manufacturers determined by the current rule.
	HardwareModel        []string `json:"hardwareModel,omitempty"`        // List of hardware models determined by the current rule.
	OperatingSystem      []string `json:"operatingSystem,omitempty"`      // List of operating systems determined by the current rule.
}
type RequestAIEndpointAnalyticsImportProfilingRulesInBulkV1ProfilingRulesConditionGroups struct {
	Type           string                                                                                        `json:"type,omitempty"`           //
	Condition      *RequestAIEndpointAnalyticsImportProfilingRulesInBulkV1ProfilingRulesConditionGroupsCondition `json:"condition,omitempty"`      //
	Operator       string                                                                                        `json:"operator,omitempty"`       //
	ConditionGroup []string                                                                                      `json:"conditionGroup,omitempty"` //
}
type RequestAIEndpointAnalyticsImportProfilingRulesInBulkV1ProfilingRulesConditionGroupsCondition struct {
	Attribute           string `json:"attribute,omitempty"`           //
	Operator            string `json:"operator,omitempty"`            //
	Value               string `json:"value,omitempty"`               //
	AttributeDictionary string `json:"attributeDictionary,omitempty"` //
}
type RequestAIEndpointAnalyticsUpdateAnExistingProfilingRuleV1 struct {
	RuleID          string                                                                    `json:"ruleId,omitempty"`          // Unique identifier for the rule. This is normally generated by the system, and client does not need to provide it for rules that need to be newly created.
	RuleName        string                                                                    `json:"ruleName,omitempty"`        // Human readable name for the rule.
	RuleType        string                                                                    `json:"ruleType,omitempty"`        // Type of the rule.
	RuleVersion     *int                                                                      `json:"ruleVersion,omitempty"`     // Version of the rule.
	RulePriority    *int                                                                      `json:"rulePriority,omitempty"`    // Priority for the rule.
	SourcePriority  *int                                                                      `json:"sourcePriority,omitempty"`  // Source priority for the rule.
	IsDeleted       *bool                                                                     `json:"isDeleted,omitempty"`       // Flag to indicate whether the rule was deleted.
	LastModifiedBy  string                                                                    `json:"lastModifiedBy,omitempty"`  // User that last modified the rule. It is read-only, and is ignored if provided as part of input request.
	LastModifiedOn  *int                                                                      `json:"lastModifiedOn,omitempty"`  // Timestamp (in epoch milliseconds) of last modification. It is read-only, and is ignored if provided as part of input request.
	PluginID        string                                                                    `json:"pluginId,omitempty"`        // Plugin for the rule. Only applicable for 'Cisco Default' rules.
	ClusterID       string                                                                    `json:"clusterId,omitempty"`       // Unique identifier for ML cluster. Only applicable for 'ML Rule'.
	Rejected        *bool                                                                     `json:"rejected,omitempty"`        // Flag to indicate whether rule has been rejected by user or not. Only applicable for 'ML Rule'.
	Result          *RequestAIEndpointAnalyticsUpdateAnExistingProfilingRuleV1Result          `json:"result,omitempty"`          //
	ConditionGroups *RequestAIEndpointAnalyticsUpdateAnExistingProfilingRuleV1ConditionGroups `json:"conditionGroups,omitempty"` //
	UsedAttributes  []string                                                                  `json:"usedAttributes,omitempty"`  // List of attributes used in the rule. Only applicable for 'Cisco Default' rules.
}
type RequestAIEndpointAnalyticsUpdateAnExistingProfilingRuleV1Result struct {
	DeviceType           []string `json:"deviceType,omitempty"`           // List of device types determined by the current rule.
	HardwareManufacturer []string `json:"hardwareManufacturer,omitempty"` // List of hardware manufacturers determined by the current rule.
	HardwareModel        []string `json:"hardwareModel,omitempty"`        // List of hardware models determined by the current rule.
	OperatingSystem      []string `json:"operatingSystem,omitempty"`      // List of operating systems determined by the current rule.
}
type RequestAIEndpointAnalyticsUpdateAnExistingProfilingRuleV1ConditionGroups struct {
	Type           string                                                                             `json:"type,omitempty"`           //
	Condition      *RequestAIEndpointAnalyticsUpdateAnExistingProfilingRuleV1ConditionGroupsCondition `json:"condition,omitempty"`      //
	Operator       string                                                                             `json:"operator,omitempty"`       //
	ConditionGroup []string                                                                           `json:"conditionGroup,omitempty"` //
}
type RequestAIEndpointAnalyticsUpdateAnExistingProfilingRuleV1ConditionGroupsCondition struct {
	Attribute           string `json:"attribute,omitempty"`           //
	Operator            string `json:"operator,omitempty"`            //
	Value               string `json:"value,omitempty"`               //
	AttributeDictionary string `json:"attributeDictionary,omitempty"` //
}

//GetAncPoliciesV1 Get ANC policies - ae9a-f945-47a8-871e
/* Fetches the list of ANC policies available in ISE.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-anc-policies-v1
*/
func (s *AIEndpointAnalyticsService) GetAncPoliciesV1() (*ResponseAIEndpointAnalyticsGetAncPoliciesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/endpoint-analytics/anc-policies"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseAIEndpointAnalyticsGetAncPoliciesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAncPoliciesV1()
		}
		return nil, response, fmt.Errorf("error with operation GetAncPoliciesV1")
	}

	result := response.Result().(*ResponseAIEndpointAnalyticsGetAncPoliciesV1)
	return result, response, err

}

//GetAIEndpointAnalyticsAttributeDictionariesV1 Get AI Endpoint Analytics attribute dictionaries - 409f-1aff-482a-ae1e
/* Fetches the list of attribute dictionaries.


@param GetAIEndpointAnalyticsAttributeDictionariesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-a-i-endpoint-analytics-attribute-dictionaries-v1
*/
func (s *AIEndpointAnalyticsService) GetAIEndpointAnalyticsAttributeDictionariesV1(GetAIEndpointAnalyticsAttributeDictionariesV1QueryParams *GetAIEndpointAnalyticsAttributeDictionariesV1QueryParams) (*ResponseAIEndpointAnalyticsGetAIEndpointAnalyticsAttributeDictionariesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/endpoint-analytics/dictionaries"

	queryString, _ := query.Values(GetAIEndpointAnalyticsAttributeDictionariesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseAIEndpointAnalyticsGetAIEndpointAnalyticsAttributeDictionariesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAIEndpointAnalyticsAttributeDictionariesV1(GetAIEndpointAnalyticsAttributeDictionariesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAIEndpointAnalyticsAttributeDictionariesV1")
	}

	result := response.Result().(*ResponseAIEndpointAnalyticsGetAIEndpointAnalyticsAttributeDictionariesV1)
	return result, response, err

}

//QueryTheEndpointsV1 Query the endpoints - aeb4-7a77-425b-b30f
/* Query the endpoints, optionally using various filter and pagination criteria. 'GET /endpoints/count' API can be used to find out the total number of endpoints matching the filter criteria.


@param QueryTheEndpointsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!query-the-endpoints-v1
*/
func (s *AIEndpointAnalyticsService) QueryTheEndpointsV1(QueryTheEndpointsV1QueryParams *QueryTheEndpointsV1QueryParams) (*ResponseAIEndpointAnalyticsQueryTheEndpointsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/endpoint-analytics/endpoints"

	queryString, _ := query.Values(QueryTheEndpointsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseAIEndpointAnalyticsQueryTheEndpointsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.QueryTheEndpointsV1(QueryTheEndpointsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation QueryTheEndpointsV1")
	}

	result := response.Result().(*ResponseAIEndpointAnalyticsQueryTheEndpointsV1)
	return result, response, err

}

//FetchTheCountOfEndpointsV1 Fetch the count of endpoints - 04b2-bbb0-472b-9ce0
/* Fetch the total count of endpoints that match the given filter criteria.


@param FetchTheCountOfEndpointsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!fetch-the-count-of-endpoints-v1
*/
func (s *AIEndpointAnalyticsService) FetchTheCountOfEndpointsV1(FetchTheCountOfEndpointsV1QueryParams *FetchTheCountOfEndpointsV1QueryParams) (*ResponseAIEndpointAnalyticsFetchTheCountOfEndpointsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/endpoint-analytics/endpoints/count"

	queryString, _ := query.Values(FetchTheCountOfEndpointsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseAIEndpointAnalyticsFetchTheCountOfEndpointsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.FetchTheCountOfEndpointsV1(FetchTheCountOfEndpointsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation FetchTheCountOfEndpointsV1")
	}

	result := response.Result().(*ResponseAIEndpointAnalyticsFetchTheCountOfEndpointsV1)
	return result, response, err

}

//GetEndpointDetailsV1 Get endpoint details - 5881-9a5e-41a8-8cce
/* Fetches details of the endpoint for the given unique identifier 'epId'.


@param epID epId path parameter. Unique identifier for the endpoint.

@param GetEndpointDetailsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-endpoint-details-v1
*/
func (s *AIEndpointAnalyticsService) GetEndpointDetailsV1(epID string, GetEndpointDetailsV1QueryParams *GetEndpointDetailsV1QueryParams) (*ResponseAIEndpointAnalyticsGetEndpointDetailsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/endpoint-analytics/endpoints/{epId}"
	path = strings.Replace(path, "{epId}", fmt.Sprintf("%v", epID), -1)

	queryString, _ := query.Values(GetEndpointDetailsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseAIEndpointAnalyticsGetEndpointDetailsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetEndpointDetailsV1(epID, GetEndpointDetailsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetEndpointDetailsV1")
	}

	result := response.Result().(*ResponseAIEndpointAnalyticsGetEndpointDetailsV1)
	return result, response, err

}

//GetListOfProfilingRulesV1 Get list of profiling rules - 07b4-eb60-435a-bf90
/* This API fetches the list of profiling rules. It can be used to show profiling rules in client applications, or export those from an environment. 'POST /profiling-rules/bulk' API can be used to import such exported rules into another environment. If this API is used to export rules to be imported into another Cisco DNA Center system, then ensure that 'includeDeleted' parameter is 'true', so that deleted rules get synchronized correctly. Use query parameters to filter the data, as required. If no filter is provided, then it will include only rules of type 'Custom Rule' in the response. By default, the response is limited to 500 records. Use 'limit' parameter to fetch higher number of records, if required. 'GET /profiling-rules/count' API can be used to find out the total number of rules in the system.


@param GetListOfProfilingRulesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-list-of-profiling-rules-v1
*/
func (s *AIEndpointAnalyticsService) GetListOfProfilingRulesV1(GetListOfProfilingRulesV1QueryParams *GetListOfProfilingRulesV1QueryParams) (*ResponseAIEndpointAnalyticsGetListOfProfilingRulesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/endpoint-analytics/profiling-rules"

	queryString, _ := query.Values(GetListOfProfilingRulesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseAIEndpointAnalyticsGetListOfProfilingRulesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetListOfProfilingRulesV1(GetListOfProfilingRulesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetListOfProfilingRulesV1")
	}

	result := response.Result().(*ResponseAIEndpointAnalyticsGetListOfProfilingRulesV1)
	return result, response, err

}

//GetCountOfProfilingRulesV1 Get count of profiling rules - 4dad-ba2c-4968-b494
/* This API fetches the count of profiling rules based on the filter values provided in the query parameters. The filter parameters are same as that of 'GET /profiling-rules' API, excluding the pagination and sort parameters.


@param GetCountOfProfilingRulesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-count-of-profiling-rules-v1
*/
func (s *AIEndpointAnalyticsService) GetCountOfProfilingRulesV1(GetCountOfProfilingRulesV1QueryParams *GetCountOfProfilingRulesV1QueryParams) (*ResponseAIEndpointAnalyticsGetCountOfProfilingRulesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/endpoint-analytics/profiling-rules/count"

	queryString, _ := query.Values(GetCountOfProfilingRulesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseAIEndpointAnalyticsGetCountOfProfilingRulesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCountOfProfilingRulesV1(GetCountOfProfilingRulesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetCountOfProfilingRulesV1")
	}

	result := response.Result().(*ResponseAIEndpointAnalyticsGetCountOfProfilingRulesV1)
	return result, response, err

}

//GetDetailsOfASingleProfilingRuleV1 Get details of a single profiling rule - 20bc-6a22-4a4b-bead
/* Fetches details of the profiling rule for the given 'ruleId'.


@param ruleID ruleId path parameter. Unique rule identifier


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-details-of-a-single-profiling-rule-v1
*/
func (s *AIEndpointAnalyticsService) GetDetailsOfASingleProfilingRuleV1(ruleID string) (*ResponseAIEndpointAnalyticsGetDetailsOfASingleProfilingRuleV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/endpoint-analytics/profiling-rules/{ruleId}"
	path = strings.Replace(path, "{ruleId}", fmt.Sprintf("%v", ruleID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseAIEndpointAnalyticsGetDetailsOfASingleProfilingRuleV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDetailsOfASingleProfilingRuleV1(ruleID)
		}
		return nil, response, fmt.Errorf("error with operation GetDetailsOfASingleProfilingRuleV1")
	}

	result := response.Result().(*ResponseAIEndpointAnalyticsGetDetailsOfASingleProfilingRuleV1)
	return result, response, err

}

//GetTaskDetailsV1 Get task details - 2689-39c4-43fa-a2f2
/* Fetches the details of backend task. Task is typically created by making call to some other API that takes longer time to execute.


@param taskID taskId path parameter. Unique identifier for the task.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-task-details-v1
*/
func (s *AIEndpointAnalyticsService) GetTaskDetailsV1(taskID string) (*ResponseAIEndpointAnalyticsGetTaskDetailsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/endpoint-analytics/tasks/{taskId}"
	path = strings.Replace(path, "{taskId}", fmt.Sprintf("%v", taskID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseAIEndpointAnalyticsGetTaskDetailsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTaskDetailsV1(taskID)
		}
		return nil, response, fmt.Errorf("error with operation GetTaskDetailsV1")
	}

	result := response.Result().(*ResponseAIEndpointAnalyticsGetTaskDetailsV1)
	return result, response, err

}

//ProcessCmdbEndpointsV1 Process CMDB endpoints - fa9f-f839-42fb-9e38
/* Processes incoming CMDB endpoints data and imports the same in AI Endpoint Analytics.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!process-cmdb-endpoints-v1
*/
func (s *AIEndpointAnalyticsService) ProcessCmdbEndpointsV1(requestAIEndpointAnalyticsProcessCMDBEndpointsV1 *RequestAIEndpointAnalyticsProcessCmdbEndpointsV1) (*resty.Response, error) {
	path := "/dna/intent/api/v1/endpoint-analytics/cmdb/endpoints"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestAIEndpointAnalyticsProcessCMDBEndpointsV1).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ProcessCmdbEndpointsV1(requestAIEndpointAnalyticsProcessCMDBEndpointsV1)
		}

		return response, fmt.Errorf("error with operation ProcessCmdbEndpointsV1")
	}

	return response, err

}

//RegisterAnEndpointV1 Register an endpoint - a895-f856-4089-92fd
/* Register a new endpoint in the system.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!register-an-endpoint-v1
*/
func (s *AIEndpointAnalyticsService) RegisterAnEndpointV1(requestAIEndpointAnalyticsRegisterAnEndpointV1 *RequestAIEndpointAnalyticsRegisterAnEndpointV1) (*resty.Response, error) {
	path := "/dna/intent/api/v1/endpoint-analytics/endpoints"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestAIEndpointAnalyticsRegisterAnEndpointV1).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RegisterAnEndpointV1(requestAIEndpointAnalyticsRegisterAnEndpointV1)
		}

		return response, fmt.Errorf("error with operation RegisterAnEndpointV1")
	}

	return response, err

}

//CreateAProfilingRuleV1 Create a profiling rule - 6cb9-98bb-47ea-90f6
/* Creates profiling rule from the request body.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-a-profiling-rule-v1
*/
func (s *AIEndpointAnalyticsService) CreateAProfilingRuleV1(requestAIEndpointAnalyticsCreateAProfilingRuleV1 *RequestAIEndpointAnalyticsCreateAProfilingRuleV1) (*ResponseAIEndpointAnalyticsCreateAProfilingRuleV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/endpoint-analytics/profiling-rules"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestAIEndpointAnalyticsCreateAProfilingRuleV1).
		SetResult(&ResponseAIEndpointAnalyticsCreateAProfilingRuleV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateAProfilingRuleV1(requestAIEndpointAnalyticsCreateAProfilingRuleV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateAProfilingRuleV1")
	}

	result := response.Result().(*ResponseAIEndpointAnalyticsCreateAProfilingRuleV1)
	return result, response, err

}

//ImportProfilingRulesInBulkV1 Import profiling rules in bulk - 70bf-885f-408a-9c74
/* This API imports the given list of profiling rules. For each record, 1) If 'ruleType' for a record is not 'Custom Rule', then it is rejected. 2) If 'ruleId' is provided in the input record,
  2a) Record with same 'ruleId' exists in the system, then it is replaced with new data.
  2b) Record with same 'ruleId' does not exist, then it is inserted in the database.
3) If 'ruleId' is not provided in the input record, then new 'ruleId' is generated by the system, and record is inserted.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!import-profiling-rules-in-bulk-v1
*/
func (s *AIEndpointAnalyticsService) ImportProfilingRulesInBulkV1(requestAIEndpointAnalyticsImportProfilingRulesInBulkV1 *RequestAIEndpointAnalyticsImportProfilingRulesInBulkV1) (*resty.Response, error) {
	path := "/dna/intent/api/v1/endpoint-analytics/profiling-rules/bulk"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestAIEndpointAnalyticsImportProfilingRulesInBulkV1).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ImportProfilingRulesInBulkV1(requestAIEndpointAnalyticsImportProfilingRulesInBulkV1)
		}

		return response, fmt.Errorf("error with operation ImportProfilingRulesInBulkV1")
	}

	return response, err

}

//UpdateARegisteredEndpointV1 Update a registered endpoint - e5af-892c-40e9-a2a1
/* Update attributes of a registered endpoint.


@param epID epId path parameter. Unique identifier for the endpoint.

*/
func (s *AIEndpointAnalyticsService) UpdateARegisteredEndpointV1(epID string, requestAIEndpointAnalyticsUpdateARegisteredEndpointV1 *RequestAIEndpointAnalyticsUpdateARegisteredEndpointV1) (*resty.Response, error) {
	path := "/dna/intent/api/v1/endpoint-analytics/endpoints/{epId}"
	path = strings.Replace(path, "{epId}", fmt.Sprintf("%v", epID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestAIEndpointAnalyticsUpdateARegisteredEndpointV1).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateARegisteredEndpointV1(epID, requestAIEndpointAnalyticsUpdateARegisteredEndpointV1)
		}
		return response, fmt.Errorf("error with operation UpdateARegisteredEndpointV1")
	}

	return response, err

}

//ApplyAncPolicyV1 Apply ANC Policy - 2ebb-79f2-4489-8b73
/* Applies given ANC policy to the endpoint.


@param epID epId path parameter. Unique identifier for the endpoint.

*/
func (s *AIEndpointAnalyticsService) ApplyAncPolicyV1(epID string, requestAIEndpointAnalyticsApplyANCPolicyV1 *RequestAIEndpointAnalyticsApplyAncPolicyV1) (*resty.Response, error) {
	path := "/dna/intent/api/v1/endpoint-analytics/endpoints/{epId}/anc-policy"
	path = strings.Replace(path, "{epId}", fmt.Sprintf("%v", epID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestAIEndpointAnalyticsApplyANCPolicyV1).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ApplyAncPolicyV1(epID, requestAIEndpointAnalyticsApplyANCPolicyV1)
		}
		return response, fmt.Errorf("error with operation ApplyAncPolicyV1")
	}

	return response, err

}

//UpdateAnExistingProfilingRuleV1 Update an existing profiling rule - c197-6aa2-4fd9-82d7
/* Updates the profiling rule for the given 'ruleId'.


@param ruleID ruleId path parameter. Unique rule identifier

*/
func (s *AIEndpointAnalyticsService) UpdateAnExistingProfilingRuleV1(ruleID string, requestAIEndpointAnalyticsUpdateAnExistingProfilingRuleV1 *RequestAIEndpointAnalyticsUpdateAnExistingProfilingRuleV1) (*resty.Response, error) {
	path := "/dna/intent/api/v1/endpoint-analytics/profiling-rules/{ruleId}"
	path = strings.Replace(path, "{ruleId}", fmt.Sprintf("%v", ruleID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestAIEndpointAnalyticsUpdateAnExistingProfilingRuleV1).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateAnExistingProfilingRuleV1(ruleID, requestAIEndpointAnalyticsUpdateAnExistingProfilingRuleV1)
		}
		return response, fmt.Errorf("error with operation UpdateAnExistingProfilingRuleV1")
	}

	return response, err

}

//DeleteAnEndpointV1 Delete an endpoint - 689d-e83b-442a-9435
/* Deletes the endpoint for the given unique identifier 'epId'.


@param epID epId path parameter. Unique identifier for the endpoint.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-an-endpoint-v1
*/
func (s *AIEndpointAnalyticsService) DeleteAnEndpointV1(epID string) (*resty.Response, error) {
	//epID string
	path := "/dna/intent/api/v1/endpoint-analytics/endpoints/{epId}"
	path = strings.Replace(path, "{epId}", fmt.Sprintf("%v", epID), -1)

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
			return s.DeleteAnEndpointV1(
				epID)
		}
		return response, fmt.Errorf("error with operation DeleteAnEndpointV1")
	}

	return response, err

}

//RevokeAncPolicyV1 Revoke ANC policy - 8982-89f3-4e1b-b3dc
/* Revokes given ANC policy from the endpoint.


@param epID epId path parameter. Unique identifier for the endpoint.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!revoke-anc-policy-v1
*/
func (s *AIEndpointAnalyticsService) RevokeAncPolicyV1(epID string) (*resty.Response, error) {
	//epID string
	path := "/dna/intent/api/v1/endpoint-analytics/endpoints/{epId}/anc-policy"
	path = strings.Replace(path, "{epId}", fmt.Sprintf("%v", epID), -1)

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
			return s.RevokeAncPolicyV1(
				epID)
		}
		return response, fmt.Errorf("error with operation RevokeAncPolicyV1")
	}

	return response, err

}

//DeleteAnExistingProfilingRuleV1 Delete an existing profiling rule - 6f9f-98d4-4b0b-9e7c
/* Deletes the profiling rule for the given 'ruleId'.


@param ruleID ruleId path parameter. Unique rule identifier


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-an-existing-profiling-rule-v1
*/
func (s *AIEndpointAnalyticsService) DeleteAnExistingProfilingRuleV1(ruleID string) (*resty.Response, error) {
	//ruleID string
	path := "/dna/intent/api/v1/endpoint-analytics/profiling-rules/{ruleId}"
	path = strings.Replace(path, "{ruleId}", fmt.Sprintf("%v", ruleID), -1)

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
			return s.DeleteAnExistingProfilingRuleV1(
				ruleID)
		}
		return response, fmt.Errorf("error with operation DeleteAnExistingProfilingRuleV1")
	}

	return response, err

}

// Alias Function
func (s *AIEndpointAnalyticsService) UpdateAnExistingProfilingRule(ruleID string, requestAIEndpointAnalyticsUpdateAnExistingProfilingRuleV1 *RequestAIEndpointAnalyticsUpdateAnExistingProfilingRuleV1) (*resty.Response, error) {
	return s.UpdateAnExistingProfilingRuleV1(ruleID, requestAIEndpointAnalyticsUpdateAnExistingProfilingRuleV1)
}

// Alias Function
func (s *AIEndpointAnalyticsService) UpdateARegisteredEndpoint(epID string, requestAIEndpointAnalyticsUpdateARegisteredEndpointV1 *RequestAIEndpointAnalyticsUpdateARegisteredEndpointV1) (*resty.Response, error) {
	return s.UpdateARegisteredEndpointV1(epID, requestAIEndpointAnalyticsUpdateARegisteredEndpointV1)
}

// Alias Function
func (s *AIEndpointAnalyticsService) RegisterAnEndpoint(requestAIEndpointAnalyticsRegisterAnEndpointV1 *RequestAIEndpointAnalyticsRegisterAnEndpointV1) (*resty.Response, error) {
	return s.RegisterAnEndpointV1(requestAIEndpointAnalyticsRegisterAnEndpointV1)
}

// Alias Function
func (s *AIEndpointAnalyticsService) ApplyAncPolicy(epID string, requestAIEndpointAnalyticsApplyANCPolicyV1 *RequestAIEndpointAnalyticsApplyAncPolicyV1) (*resty.Response, error) {
	return s.ApplyAncPolicyV1(epID, requestAIEndpointAnalyticsApplyANCPolicyV1)
}

// Alias Function
func (s *AIEndpointAnalyticsService) ProcessCmdbEndpoints(requestAIEndpointAnalyticsProcessCMDBEndpointsV1 *RequestAIEndpointAnalyticsProcessCmdbEndpointsV1) (*resty.Response, error) {
	return s.ProcessCmdbEndpointsV1(requestAIEndpointAnalyticsProcessCMDBEndpointsV1)
}

// Alias Function
func (s *AIEndpointAnalyticsService) GetDetailsOfASingleProfilingRule(ruleID string) (*ResponseAIEndpointAnalyticsGetDetailsOfASingleProfilingRuleV1, *resty.Response, error) {
	return s.GetDetailsOfASingleProfilingRuleV1(ruleID)
}

// Alias Function
func (s *AIEndpointAnalyticsService) GetAncPolicies() (*ResponseAIEndpointAnalyticsGetAncPoliciesV1, *resty.Response, error) {
	return s.GetAncPoliciesV1()
}

// Alias Function
func (s *AIEndpointAnalyticsService) CreateAProfilingRule(requestAIEndpointAnalyticsCreateAProfilingRuleV1 *RequestAIEndpointAnalyticsCreateAProfilingRuleV1) (*ResponseAIEndpointAnalyticsCreateAProfilingRuleV1, *resty.Response, error) {
	return s.CreateAProfilingRuleV1(requestAIEndpointAnalyticsCreateAProfilingRuleV1)
}

// Alias Function
func (s *AIEndpointAnalyticsService) QueryTheEndpoints(QueryTheEndpointsV1QueryParams *QueryTheEndpointsV1QueryParams) (*ResponseAIEndpointAnalyticsQueryTheEndpointsV1, *resty.Response, error) {
	return s.QueryTheEndpointsV1(QueryTheEndpointsV1QueryParams)
}

// Alias Function
func (s *AIEndpointAnalyticsService) RevokeAncPolicy(epID string) (*resty.Response, error) {
	return s.RevokeAncPolicyV1(epID)
}

// Alias Function
func (s *AIEndpointAnalyticsService) GetListOfProfilingRules(GetListOfProfilingRulesV1QueryParams *GetListOfProfilingRulesV1QueryParams) (*ResponseAIEndpointAnalyticsGetListOfProfilingRulesV1, *resty.Response, error) {
	return s.GetListOfProfilingRulesV1(GetListOfProfilingRulesV1QueryParams)
}

// Alias Function
func (s *AIEndpointAnalyticsService) GetTaskDetails(taskID string) (*ResponseAIEndpointAnalyticsGetTaskDetailsV1, *resty.Response, error) {
	return s.GetTaskDetailsV1(taskID)
}

// Alias Function
func (s *AIEndpointAnalyticsService) GetAIEndpointAnalyticsAttributeDictionaries(GetAIEndpointAnalyticsAttributeDictionariesV1QueryParams *GetAIEndpointAnalyticsAttributeDictionariesV1QueryParams) (*ResponseAIEndpointAnalyticsGetAIEndpointAnalyticsAttributeDictionariesV1, *resty.Response, error) {
	return s.GetAIEndpointAnalyticsAttributeDictionariesV1(GetAIEndpointAnalyticsAttributeDictionariesV1QueryParams)
}

// Alias Function
func (s *AIEndpointAnalyticsService) FetchTheCountOfEndpoints(FetchTheCountOfEndpointsV1QueryParams *FetchTheCountOfEndpointsV1QueryParams) (*ResponseAIEndpointAnalyticsFetchTheCountOfEndpointsV1, *resty.Response, error) {
	return s.FetchTheCountOfEndpointsV1(FetchTheCountOfEndpointsV1QueryParams)
}

// Alias Function
func (s *AIEndpointAnalyticsService) ImportProfilingRulesInBulk(requestAIEndpointAnalyticsImportProfilingRulesInBulkV1 *RequestAIEndpointAnalyticsImportProfilingRulesInBulkV1) (*resty.Response, error) {
	return s.ImportProfilingRulesInBulkV1(requestAIEndpointAnalyticsImportProfilingRulesInBulkV1)
}

// Alias Function
func (s *AIEndpointAnalyticsService) DeleteAnExistingProfilingRule(ruleID string) (*resty.Response, error) {
	return s.DeleteAnExistingProfilingRuleV1(ruleID)
}

// Alias Function
func (s *AIEndpointAnalyticsService) GetEndpointDetails(epID string, GetEndpointDetailsV1QueryParams *GetEndpointDetailsV1QueryParams) (*ResponseAIEndpointAnalyticsGetEndpointDetailsV1, *resty.Response, error) {
	return s.GetEndpointDetailsV1(epID, GetEndpointDetailsV1QueryParams)
}

// Alias Function
func (s *AIEndpointAnalyticsService) GetCountOfProfilingRules(GetCountOfProfilingRulesV1QueryParams *GetCountOfProfilingRulesV1QueryParams) (*ResponseAIEndpointAnalyticsGetCountOfProfilingRulesV1, *resty.Response, error) {
	return s.GetCountOfProfilingRulesV1(GetCountOfProfilingRulesV1QueryParams)
}

// Alias Function
func (s *AIEndpointAnalyticsService) DeleteAnEndpoint(epID string) (*resty.Response, error) {
	return s.DeleteAnEndpointV1(epID)
}
