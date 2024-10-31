package catalyst

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type ApplicationPolicyService service

type GetApplicationPolicyV1QueryParams struct {
	PolicyScope string `url:"policyScope,omitempty"` //policy scope name
}
type GetApplicationPolicyQueuingProfileV1QueryParams struct {
	Name string `url:"name,omitempty"` //queuing profile name
}
type GetApplicationSetsV1QueryParams struct {
	Offset float64 `url:"offset,omitempty"` //
	Limit  float64 `url:"limit,omitempty"`  //
	Name   string  `url:"name,omitempty"`   //
}
type DeleteApplicationSetV1QueryParams struct {
	ID string `url:"id,omitempty"` //
}
type DeleteApplicationV1QueryParams struct {
	ID string `url:"id,omitempty"` //Application's Id
}
type GetApplicationsV1QueryParams struct {
	Offset float64 `url:"offset,omitempty"` //The offset of the first application to be returned
	Limit  float64 `url:"limit,omitempty"`  //The maximum number of applications to be returned
	Name   string  `url:"name,omitempty"`   //Application's name
}
type GetQosDeviceInterfaceInfoV1QueryParams struct {
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` //network device id
}
type GetApplicationSetsV2QueryParams struct {
	Attributes string  `url:"attributes,omitempty"` //Attributes to retrieve, valid value applicationSet
	Name       string  `url:"name,omitempty"`       //Application set name
	Offset     float64 `url:"offset,omitempty"`     //The starting point or index from where the paginated results should begin.
	Limit      float64 `url:"limit,omitempty"`      //The limit which is the maximum number of items to include in a single page of results, max value 500
}
type GetApplicationSetCountV2QueryParams struct {
	ScalableGroupType string `url:"scalableGroupType,omitempty"` //Scalable group type to retrieve, valid value APPLICATION_GROUP
}
type GetApplicationsV2QueryParams struct {
	Attributes string  `url:"attributes,omitempty"` //Attributes to retrieve, valid value application
	Name       string  `url:"name,omitempty"`       //The application name
	Offset     float64 `url:"offset,omitempty"`     //The starting point or index from where the paginated results should begin.
	Limit      float64 `url:"limit,omitempty"`      //The limit which is the maximum number of items to include in a single page of results, max value 500
}
type GetApplicationCountV2QueryParams struct {
	ScalableGroupType string `url:"scalableGroupType,omitempty"` //scalable group type to retrieve, valid value APPLICATION
}

type ResponseApplicationPolicyGetApplicationPolicyV1 struct {
	Response *[]ResponseApplicationPolicyGetApplicationPolicyV1Response `json:"response,omitempty"` //
	Version  string                                                     `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyGetApplicationPolicyV1Response struct {
	ID                  string                                                                      `json:"id,omitempty"`                  // Id of Group based policy
	InstanceID          *int                                                                        `json:"instanceId,omitempty"`          // Instance id
	DisplayName         string                                                                      `json:"displayName,omitempty"`         // Display name
	InstanceCreatedOn   *int                                                                        `json:"instanceCreatedOn,omitempty"`   // Instance created on
	InstanceUpdatedOn   *int                                                                        `json:"instanceUpdatedOn,omitempty"`   // Instance updated on
	InstanceVersion     *float64                                                                    `json:"instanceVersion,omitempty"`     // Instance version
	CreateTime          *int                                                                        `json:"createTime,omitempty"`          // Create time
	Deployed            *bool                                                                       `json:"deployed,omitempty"`            // Deployed
	IsSeeded            *bool                                                                       `json:"isSeeded,omitempty"`            // Is seeded
	IsStale             *bool                                                                       `json:"isStale,omitempty"`             // Is stale
	LastUpdateTime      *int                                                                        `json:"lastUpdateTime,omitempty"`      // Last update time
	Name                string                                                                      `json:"name,omitempty"`                // Concatination of <polcy name>_<application-set-name> or <polcy name>_global_policy_configuration or <polcy name>_queuing_customization
	Namespace           string                                                                      `json:"namespace,omitempty"`           // Namespace
	ProvisioningState   string                                                                      `json:"provisioningState,omitempty"`   // Provisioning state
	Qualifier           string                                                                      `json:"qualifier,omitempty"`           // Qualifier
	ResourceVersion     *float64                                                                    `json:"resourceVersion,omitempty"`     // Resource version
	TargetIDList        *[]ResponseApplicationPolicyGetApplicationPolicyV1ResponseTargetIDList      `json:"targetIdList,omitempty"`        // Target id list
	Type                string                                                                      `json:"type,omitempty"`                // Type
	CfsChangeInfo       *[]ResponseApplicationPolicyGetApplicationPolicyV1ResponseCfsChangeInfo     `json:"cfsChangeInfo,omitempty"`       // Cfs change info
	CustomProvisions    *[]ResponseApplicationPolicyGetApplicationPolicyV1ResponseCustomProvisions  `json:"customProvisions,omitempty"`    // Custom provisions
	DeletePolicyStatus  string                                                                      `json:"deletePolicyStatus,omitempty"`  // NONE: deployed policy to devices, DELETED: delete policy from devices, RESTORED: restored to original configuration
	Internal            *bool                                                                       `json:"internal,omitempty"`            // Internal
	IsDeleted           *bool                                                                       `json:"isDeleted,omitempty"`           // Is deleted
	IsEnabled           *bool                                                                       `json:"isEnabled,omitempty"`           // Is enabled
	IsScopeStale        *bool                                                                       `json:"isScopeStale,omitempty"`        // Is scope stale
	IseReserved         *bool                                                                       `json:"iseReserved,omitempty"`         // Is reserved
	PolicyScope         string                                                                      `json:"policyScope,omitempty"`         // Policy name
	PolicyStatus        string                                                                      `json:"policyStatus,omitempty"`        // Policy status
	Priority            *int                                                                        `json:"priority,omitempty"`            // Priority
	Pushed              *bool                                                                       `json:"pushed,omitempty"`              // Pushed
	AdvancedPolicyScope *ResponseApplicationPolicyGetApplicationPolicyV1ResponseAdvancedPolicyScope `json:"advancedPolicyScope,omitempty"` //
	ContractList        *[]ResponseApplicationPolicyGetApplicationPolicyV1ResponseContractList      `json:"contractList,omitempty"`        // Contract list
	ExclusiveContract   *ResponseApplicationPolicyGetApplicationPolicyV1ResponseExclusiveContract   `json:"exclusiveContract,omitempty"`   //
	IDentitySource      *ResponseApplicationPolicyGetApplicationPolicyV1ResponseIDentitySource      `json:"identitySource,omitempty"`      //
	Producer            *ResponseApplicationPolicyGetApplicationPolicyV1ResponseProducer            `json:"producer,omitempty"`            //
	Consumer            *ResponseApplicationPolicyGetApplicationPolicyV1ResponseConsumer            `json:"consumer,omitempty"`            //
}
type ResponseApplicationPolicyGetApplicationPolicyV1ResponseTargetIDList interface{}
type ResponseApplicationPolicyGetApplicationPolicyV1ResponseCfsChangeInfo interface{}
type ResponseApplicationPolicyGetApplicationPolicyV1ResponseCustomProvisions interface{}
type ResponseApplicationPolicyGetApplicationPolicyV1ResponseAdvancedPolicyScope struct {
	ID                         string                                                                                                  `json:"id,omitempty"`                         // Id of Advanced policy scope
	InstanceID                 *int                                                                                                    `json:"instanceId,omitempty"`                 // Instance id
	DisplayName                string                                                                                                  `json:"displayName,omitempty"`                // Display name
	InstanceCreatedOn          *int                                                                                                    `json:"instanceCreatedOn,omitempty"`          // Instance created on
	InstanceUpdatedOn          *int                                                                                                    `json:"instanceUpdatedOn,omitempty"`          // Instance updated on
	InstanceVersion            *float64                                                                                                `json:"instanceVersion,omitempty"`            // Instance version
	Name                       string                                                                                                  `json:"name,omitempty"`                       // Policy name
	AdvancedPolicyScopeElement *[]ResponseApplicationPolicyGetApplicationPolicyV1ResponseAdvancedPolicyScopeAdvancedPolicyScopeElement `json:"advancedPolicyScopeElement,omitempty"` //
}
type ResponseApplicationPolicyGetApplicationPolicyV1ResponseAdvancedPolicyScopeAdvancedPolicyScopeElement struct {
	ID                string                                                                                                      `json:"id,omitempty"`                // Id of Advanced policy scope element
	InstanceID        *int                                                                                                        `json:"instanceId,omitempty"`        // Instance id
	DisplayName       string                                                                                                      `json:"displayName,omitempty"`       // Display name
	InstanceCreatedOn *int                                                                                                        `json:"instanceCreatedOn,omitempty"` // Instance created on
	InstanceUpdatedOn *int                                                                                                        `json:"instanceUpdatedOn,omitempty"` // Instance updated on
	InstanceVersion   *float64                                                                                                    `json:"instanceVersion,omitempty"`   // Instance version
	GroupID           []string                                                                                                    `json:"groupId,omitempty"`           // Group id
	SSID              *[]ResponseApplicationPolicyGetApplicationPolicyV1ResponseAdvancedPolicyScopeAdvancedPolicyScopeElementSSID `json:"ssid,omitempty"`              // Ssid
}
type ResponseApplicationPolicyGetApplicationPolicyV1ResponseAdvancedPolicyScopeAdvancedPolicyScopeElementSSID interface{}
type ResponseApplicationPolicyGetApplicationPolicyV1ResponseContractList interface{}
type ResponseApplicationPolicyGetApplicationPolicyV1ResponseExclusiveContract struct {
	ID                string                                                                            `json:"id,omitempty"`                // Id of Exclusive contract
	InstanceID        *int                                                                              `json:"instanceId,omitempty"`        // Instance id
	DisplayName       string                                                                            `json:"displayName,omitempty"`       // Display name
	InstanceCreatedOn *int                                                                              `json:"instanceCreatedOn,omitempty"` // Instance created on
	InstanceUpdatedOn *int                                                                              `json:"instanceUpdatedOn,omitempty"` // Instance updated on
	InstanceVersion   *float64                                                                          `json:"instanceVersion,omitempty"`   // Instance version
	Clause            *[]ResponseApplicationPolicyGetApplicationPolicyV1ResponseExclusiveContractClause `json:"clause,omitempty"`            //
}
type ResponseApplicationPolicyGetApplicationPolicyV1ResponseExclusiveContractClause struct {
	ID                    string   `json:"id,omitempty"`                    // Id of Business relevance or Application policy knobs clause
	InstanceID            *int     `json:"instanceId,omitempty"`            // Instance id
	DisplayName           string   `json:"displayName,omitempty"`           // Display name
	InstanceCreatedOn     *int     `json:"instanceCreatedOn,omitempty"`     // Instance created on
	InstanceUpdatedOn     *int     `json:"instanceUpdatedOn,omitempty"`     // Instance updated on
	InstanceVersion       *float64 `json:"instanceVersion,omitempty"`       // Instance version
	Priority              *int     `json:"priority,omitempty"`              // Priority
	Type                  string   `json:"type,omitempty"`                  // Type
	RelevanceLevel        string   `json:"relevanceLevel,omitempty"`        // Relevance level
	DeviceRemovalBehavior string   `json:"deviceRemovalBehavior,omitempty"` // Device removal behavior
	HostTrackingEnabled   *bool    `json:"hostTrackingEnabled,omitempty"`   // Host tracking enabled
}
type ResponseApplicationPolicyGetApplicationPolicyV1ResponseIDentitySource struct {
	ID                string   `json:"id,omitempty"`                // Id of Identity source
	InstanceID        *int     `json:"instanceId,omitempty"`        // Instance id
	DisplayName       string   `json:"displayName,omitempty"`       // Display name
	InstanceCreatedOn *int     `json:"instanceCreatedOn,omitempty"` // Instance created on
	InstanceUpdatedOn *int     `json:"instanceUpdatedOn,omitempty"` // Instance updated on
	InstanceVersion   *float64 `json:"instanceVersion,omitempty"`   // Instance version
	State             string   `json:"state,omitempty"`             // State
	Type              string   `json:"type,omitempty"`              // Type
}
type ResponseApplicationPolicyGetApplicationPolicyV1ResponseProducer struct {
	ID                string                                                                          `json:"id,omitempty"`                // Id of Producer
	InstanceID        *int                                                                            `json:"instanceId,omitempty"`        // Instance id
	DisplayName       string                                                                          `json:"displayName,omitempty"`       // Display name
	InstanceCreatedOn *int                                                                            `json:"instanceCreatedOn,omitempty"` // Instance created on
	InstanceUpdatedOn *int                                                                            `json:"instanceUpdatedOn,omitempty"` // Instance updated on
	InstanceVersion   *float64                                                                        `json:"instanceVersion,omitempty"`   // Instance version
	ScalableGroup     *[]ResponseApplicationPolicyGetApplicationPolicyV1ResponseProducerScalableGroup `json:"scalableGroup,omitempty"`     //
}
type ResponseApplicationPolicyGetApplicationPolicyV1ResponseProducerScalableGroup struct {
	IDRef string `json:"idRef,omitempty"` // Id ref to application-set or application Scalable group
}
type ResponseApplicationPolicyGetApplicationPolicyV1ResponseConsumer struct {
	ID                string                                                                          `json:"id,omitempty"`                // Id of Consumer
	InstanceID        *int                                                                            `json:"instanceId,omitempty"`        // Instance id
	DisplayName       string                                                                          `json:"displayName,omitempty"`       // Display name
	InstanceCreatedOn *int                                                                            `json:"instanceCreatedOn,omitempty"` // Instance created on
	InstanceUpdatedOn *int                                                                            `json:"instanceUpdatedOn,omitempty"` // Instance updated on
	InstanceVersion   *float64                                                                        `json:"instanceVersion,omitempty"`   // Instance version
	ScalableGroup     *[]ResponseApplicationPolicyGetApplicationPolicyV1ResponseConsumerScalableGroup `json:"scalableGroup,omitempty"`     //
}
type ResponseApplicationPolicyGetApplicationPolicyV1ResponseConsumerScalableGroup struct {
	IDRef string `json:"idRef,omitempty"` // Id ref to application Scalable group
}
type ResponseApplicationPolicyGetApplicationPolicyDefaultV1 struct {
	Response *[]ResponseApplicationPolicyGetApplicationPolicyDefaultV1Response `json:"response,omitempty"` //
	Version  string                                                            `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyGetApplicationPolicyDefaultV1Response struct {
	ID                 string                                                                            `json:"id,omitempty"`                 // Id of Group based policy
	InstanceID         *int                                                                              `json:"instanceId,omitempty"`         // Instance id
	DisplayName        string                                                                            `json:"displayName,omitempty"`        // Display name
	InstanceCreatedOn  *int                                                                              `json:"instanceCreatedOn,omitempty"`  // Instance created on
	InstanceUpdatedOn  *int                                                                              `json:"instanceUpdatedOn,omitempty"`  // Instance updated on
	InstanceVersion    *float64                                                                          `json:"instanceVersion,omitempty"`    // Instance version
	CreateTime         *int                                                                              `json:"createTime,omitempty"`         // Create time
	Deployed           *bool                                                                             `json:"deployed,omitempty"`           // Deployed
	IsSeeded           *bool                                                                             `json:"isSeeded,omitempty"`           // Is seeded
	IsStale            *bool                                                                             `json:"isStale,omitempty"`            // Is stale
	LastUpdateTime     *int                                                                              `json:"lastUpdateTime,omitempty"`     // Last update time
	Name               string                                                                            `json:"name,omitempty"`               // Concatination of <polcy name>_<application-set-name> or <polcy name>_global_policy_configuration or <polcy name>_queuing_customization
	Namespace          string                                                                            `json:"namespace,omitempty"`          // Namespace
	ProvisioningState  string                                                                            `json:"provisioningState,omitempty"`  // Provisioning state
	Qualifier          string                                                                            `json:"qualifier,omitempty"`          // Qualifier
	ResourceVersion    *float64                                                                          `json:"resourceVersion,omitempty"`    // Resource version
	TargetIDList       *[]ResponseApplicationPolicyGetApplicationPolicyDefaultV1ResponseTargetIDList     `json:"targetIdList,omitempty"`       // Target id list
	Type               string                                                                            `json:"type,omitempty"`               // Type
	CfsChangeInfo      *[]ResponseApplicationPolicyGetApplicationPolicyDefaultV1ResponseCfsChangeInfo    `json:"cfsChangeInfo,omitempty"`      // Cfs change info
	CustomProvisions   *[]ResponseApplicationPolicyGetApplicationPolicyDefaultV1ResponseCustomProvisions `json:"customProvisions,omitempty"`   // Custom provisions
	DeletePolicyStatus string                                                                            `json:"deletePolicyStatus,omitempty"` // NONE: deployed policy to devices, DELETED: delete policy from devices, RESTORED: restored to original configuration
	Internal           *bool                                                                             `json:"internal,omitempty"`           // Internal
	IsDeleted          *bool                                                                             `json:"isDeleted,omitempty"`          // Is deleted
	IsEnabled          *bool                                                                             `json:"isEnabled,omitempty"`          // Is enabled
	IsScopeStale       *bool                                                                             `json:"isScopeStale,omitempty"`       // Is scope stale
	IseReserved        *bool                                                                             `json:"iseReserved,omitempty"`        // Is reserved
	PolicyStatus       string                                                                            `json:"policyStatus,omitempty"`       // Policy status
	Priority           *int                                                                              `json:"priority,omitempty"`           // Priority
	Pushed             *bool                                                                             `json:"pushed,omitempty"`             // Pushed
	ContractList       *[]ResponseApplicationPolicyGetApplicationPolicyDefaultV1ResponseContractList     `json:"contractList,omitempty"`       // Contract list
	ExclusiveContract  *ResponseApplicationPolicyGetApplicationPolicyDefaultV1ResponseExclusiveContract  `json:"exclusiveContract,omitempty"`  //
	IDentitySource     *ResponseApplicationPolicyGetApplicationPolicyDefaultV1ResponseIDentitySource     `json:"identitySource,omitempty"`     //
	Producer           *ResponseApplicationPolicyGetApplicationPolicyDefaultV1ResponseProducer           `json:"producer,omitempty"`           //
}
type ResponseApplicationPolicyGetApplicationPolicyDefaultV1ResponseTargetIDList interface{}
type ResponseApplicationPolicyGetApplicationPolicyDefaultV1ResponseCfsChangeInfo interface{}
type ResponseApplicationPolicyGetApplicationPolicyDefaultV1ResponseCustomProvisions interface{}
type ResponseApplicationPolicyGetApplicationPolicyDefaultV1ResponseContractList interface{}
type ResponseApplicationPolicyGetApplicationPolicyDefaultV1ResponseExclusiveContract struct {
	ID                string                                                                                   `json:"id,omitempty"`                // Id of Exclusive contract
	InstanceID        *int                                                                                     `json:"instanceId,omitempty"`        // Instance id
	DisplayName       string                                                                                   `json:"displayName,omitempty"`       // Display name
	InstanceCreatedOn *int                                                                                     `json:"instanceCreatedOn,omitempty"` // Instance created on
	InstanceUpdatedOn *int                                                                                     `json:"instanceUpdatedOn,omitempty"` // Instance updated on
	InstanceVersion   *float64                                                                                 `json:"instanceVersion,omitempty"`   // Instance version
	Clause            *[]ResponseApplicationPolicyGetApplicationPolicyDefaultV1ResponseExclusiveContractClause `json:"clause,omitempty"`            //
}
type ResponseApplicationPolicyGetApplicationPolicyDefaultV1ResponseExclusiveContractClause struct {
	ID                string   `json:"id,omitempty"`                // Id of Business relevance clause
	InstanceID        *int     `json:"instanceId,omitempty"`        // Instance id
	DisplayName       string   `json:"displayName,omitempty"`       // Display name
	InstanceCreatedOn *int     `json:"instanceCreatedOn,omitempty"` // Instance created on
	InstanceUpdatedOn *int     `json:"instanceUpdatedOn,omitempty"` // Instance updated on
	InstanceVersion   *float64 `json:"instanceVersion,omitempty"`   // Instance version
	Priority          *int     `json:"priority,omitempty"`          // Priority
	Type              string   `json:"type,omitempty"`              // Type
	RelevanceLevel    string   `json:"relevanceLevel,omitempty"`    // Relevance level
}
type ResponseApplicationPolicyGetApplicationPolicyDefaultV1ResponseIDentitySource struct {
	ID                string   `json:"id,omitempty"`                // Id of Identity source
	InstanceID        *int     `json:"instanceId,omitempty"`        // Instance id
	DisplayName       string   `json:"displayName,omitempty"`       // Display name
	InstanceCreatedOn *int     `json:"instanceCreatedOn,omitempty"` // Instance created on
	InstanceUpdatedOn *int     `json:"instanceUpdatedOn,omitempty"` // Instance updated on
	InstanceVersion   *float64 `json:"instanceVersion,omitempty"`   // Instance version
	State             string   `json:"state,omitempty"`             // State
	Type              string   `json:"type,omitempty"`              // Type
}
type ResponseApplicationPolicyGetApplicationPolicyDefaultV1ResponseProducer struct {
	ID                string                                                                                 `json:"id,omitempty"`                // Id of Producer
	InstanceID        *int                                                                                   `json:"instanceId,omitempty"`        // Instance id
	DisplayName       string                                                                                 `json:"displayName,omitempty"`       // Display name
	InstanceCreatedOn *int                                                                                   `json:"instanceCreatedOn,omitempty"` // Instance created on
	InstanceUpdatedOn *int                                                                                   `json:"instanceUpdatedOn,omitempty"` // Instance updated on
	InstanceVersion   *float64                                                                               `json:"instanceVersion,omitempty"`   // Instance version
	ScalableGroup     *[]ResponseApplicationPolicyGetApplicationPolicyDefaultV1ResponseProducerScalableGroup `json:"scalableGroup,omitempty"`     //
}
type ResponseApplicationPolicyGetApplicationPolicyDefaultV1ResponseProducerScalableGroup struct {
	IDRef string `json:"idRef,omitempty"` // Id ref to application-set or application Scalable group
}
type ResponseApplicationPolicyApplicationPolicyIntentV1 struct {
	Response *ResponseApplicationPolicyApplicationPolicyIntentV1Response `json:"response,omitempty"` //
	Version  string                                                      `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyApplicationPolicyIntentV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task id
	URL    string `json:"url,omitempty"`    // Task url
}
type ResponseApplicationPolicyGetApplicationPolicyQueuingProfileV1 struct {
	Response *[]ResponseApplicationPolicyGetApplicationPolicyQueuingProfileV1Response `json:"response,omitempty"` //
	Version  string                                                                   `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyGetApplicationPolicyQueuingProfileV1Response struct {
	ID                 string                                                                                     `json:"id,omitempty"`                 // Id of Queueing profile
	InstanceID         *int                                                                                       `json:"instanceId,omitempty"`         // Instance id
	DisplayName        string                                                                                     `json:"displayName,omitempty"`        // Display name
	InstanceCreatedOn  *int                                                                                       `json:"instanceCreatedOn,omitempty"`  // Instance created on
	InstanceUpdatedOn  *int                                                                                       `json:"instanceUpdatedOn,omitempty"`  // Instance updated on
	InstanceVersion    *float64                                                                                   `json:"instanceVersion,omitempty"`    // Instance version
	CreateTime         *int                                                                                       `json:"createTime,omitempty"`         // Create time
	Deployed           *bool                                                                                      `json:"deployed,omitempty"`           // Deployed
	Description        string                                                                                     `json:"description,omitempty"`        // Free test description
	IsSeeded           *bool                                                                                      `json:"isSeeded,omitempty"`           // Is seeded
	IsStale            *bool                                                                                      `json:"isStale,omitempty"`            // Is stale
	LastUpdateTime     *int                                                                                       `json:"lastUpdateTime,omitempty"`     // Last update time
	Name               string                                                                                     `json:"name,omitempty"`               // Queueing profile name
	Namespace          string                                                                                     `json:"namespace,omitempty"`          // Namespace
	ProvisioningState  string                                                                                     `json:"provisioningState,omitempty"`  // Provisioning State
	Qualifier          string                                                                                     `json:"qualifier,omitempty"`          // Qualifier
	ResourceVersion    *float64                                                                                   `json:"resourceVersion,omitempty"`    // Resource version
	TargetIDList       *[]ResponseApplicationPolicyGetApplicationPolicyQueuingProfileV1ResponseTargetIDList       `json:"targetIdList,omitempty"`       // Target id list
	Type               string                                                                                     `json:"type,omitempty"`               // Type
	CfsChangeInfo      *[]ResponseApplicationPolicyGetApplicationPolicyQueuingProfileV1ResponseCfsChangeInfo      `json:"cfsChangeInfo,omitempty"`      // Cfs change info
	CustomProvisions   *[]ResponseApplicationPolicyGetApplicationPolicyQueuingProfileV1ResponseCustomProvisions   `json:"customProvisions,omitempty"`   // Custom provisions
	GenID              *float64                                                                                   `json:"genId,omitempty"`              // Gen id
	Internal           *bool                                                                                      `json:"internal,omitempty"`           // Internal
	IsDeleted          *bool                                                                                      `json:"isDeleted,omitempty"`          // Is deleted
	IseReserved        *bool                                                                                      `json:"iseReserved,omitempty"`        // Is reserved
	Pushed             *bool                                                                                      `json:"pushed,omitempty"`             // Pushed
	Clause             *[]ResponseApplicationPolicyGetApplicationPolicyQueuingProfileV1ResponseClause             `json:"clause,omitempty"`             //
	ContractClassifier *[]ResponseApplicationPolicyGetApplicationPolicyQueuingProfileV1ResponseContractClassifier `json:"contractClassifier,omitempty"` // Contract classifier
}
type ResponseApplicationPolicyGetApplicationPolicyQueuingProfileV1ResponseTargetIDList interface{}
type ResponseApplicationPolicyGetApplicationPolicyQueuingProfileV1ResponseCfsChangeInfo interface{}
type ResponseApplicationPolicyGetApplicationPolicyQueuingProfileV1ResponseCustomProvisions interface{}
type ResponseApplicationPolicyGetApplicationPolicyQueuingProfileV1ResponseClause struct {
	ID                                string                                                                                                       `json:"id,omitempty"`                                // Id
	InstanceID                        *int                                                                                                         `json:"instanceId,omitempty"`                        // Instance id
	DisplayName                       string                                                                                                       `json:"displayName,omitempty"`                       // Display name
	InstanceCreatedOn                 *int                                                                                                         `json:"instanceCreatedOn,omitempty"`                 // Instance created on
	InstanceUpdatedOn                 *int                                                                                                         `json:"instanceUpdatedOn,omitempty"`                 // Instance updated on
	InstanceVersion                   *float64                                                                                                     `json:"instanceVersion,omitempty"`                   // Instance version
	Priority                          *int                                                                                                         `json:"priority,omitempty"`                          // Priority
	Type                              string                                                                                                       `json:"type,omitempty"`                              // Type
	IsCommonBetweenAllInterfaceSpeeds *bool                                                                                                        `json:"isCommonBetweenAllInterfaceSpeeds,omitempty"` // Is common between all interface speeds
	InterfaceSpeedBandwidthClauses    *[]ResponseApplicationPolicyGetApplicationPolicyQueuingProfileV1ResponseClauseInterfaceSpeedBandwidthClauses `json:"interfaceSpeedBandwidthClauses,omitempty"`    //
	TcDscpSettings                    *[]ResponseApplicationPolicyGetApplicationPolicyQueuingProfileV1ResponseClauseTcDscpSettings                 `json:"tcDscpSettings,omitempty"`                    //
}
type ResponseApplicationPolicyGetApplicationPolicyQueuingProfileV1ResponseClauseInterfaceSpeedBandwidthClauses struct {
	ID                  string                                                                                                                          `json:"id,omitempty"`                  // Id
	InstanceID          *int                                                                                                                            `json:"instanceId,omitempty"`          // Instance id
	DisplayName         string                                                                                                                          `json:"displayName,omitempty"`         // Display name
	InstanceCreatedOn   *int                                                                                                                            `json:"instanceCreatedOn,omitempty"`   // Instance created on
	InstanceUpdatedOn   *int                                                                                                                            `json:"instanceUpdatedOn,omitempty"`   // Instance updated on
	InstanceVersion     *float64                                                                                                                        `json:"instanceVersion,omitempty"`     // Instance version
	InterfaceSpeed      string                                                                                                                          `json:"interfaceSpeed,omitempty"`      // Interface speed
	TcBandwidthSettings *[]ResponseApplicationPolicyGetApplicationPolicyQueuingProfileV1ResponseClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings `json:"tcBandwidthSettings,omitempty"` //
}
type ResponseApplicationPolicyGetApplicationPolicyQueuingProfileV1ResponseClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings struct {
	ID                  string   `json:"id,omitempty"`                  // Id
	InstanceID          *int     `json:"instanceId,omitempty"`          // Instance id
	DisplayName         string   `json:"displayName,omitempty"`         // Display name
	InstanceCreatedOn   *int     `json:"instanceCreatedOn,omitempty"`   // Instance created on
	InstanceUpdatedOn   *int     `json:"instanceUpdatedOn,omitempty"`   // Instance updated on
	InstanceVersion     *float64 `json:"instanceVersion,omitempty"`     // Instance version
	BandwidthPercentage *int     `json:"bandwidthPercentage,omitempty"` // Bandwidth percentage
	TrafficClass        string   `json:"trafficClass,omitempty"`        // Traffic Class
}
type ResponseApplicationPolicyGetApplicationPolicyQueuingProfileV1ResponseClauseTcDscpSettings struct {
	ID                string   `json:"id,omitempty"`                // Id
	InstanceID        *int     `json:"instanceId,omitempty"`        // Instance id
	DisplayName       string   `json:"displayName,omitempty"`       // Display name
	InstanceCreatedOn *int     `json:"instanceCreatedOn,omitempty"` // Instance created on
	InstanceUpdatedOn *int     `json:"instanceUpdatedOn,omitempty"` // Instance updated on
	InstanceVersion   *float64 `json:"instanceVersion,omitempty"`   // Instance version
	Dscp              string   `json:"dscp,omitempty"`              // Dscp value
	TrafficClass      string   `json:"trafficClass,omitempty"`      // Traffic Class
}
type ResponseApplicationPolicyGetApplicationPolicyQueuingProfileV1ResponseContractClassifier interface{}
type ResponseApplicationPolicyUpdateApplicationPolicyQueuingProfileV1 struct {
	Response *ResponseApplicationPolicyUpdateApplicationPolicyQueuingProfileV1Response `json:"response,omitempty"` //
	Version  string                                                                    `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyUpdateApplicationPolicyQueuingProfileV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task id
	URL    string `json:"url,omitempty"`    // Task url
}
type ResponseApplicationPolicyCreateApplicationPolicyQueuingProfileV1 struct {
	Response *ResponseApplicationPolicyCreateApplicationPolicyQueuingProfileV1Response `json:"response,omitempty"` //
	Version  string                                                                    `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyCreateApplicationPolicyQueuingProfileV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task id
	URL    string `json:"url,omitempty"`    // Task url
}
type ResponseApplicationPolicyGetApplicationPolicyQueuingProfileCountV1 struct {
	Response *int   `json:"response,omitempty"` // Total number of Queueing Profile
	Version  string `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyDeleteApplicationPolicyQueuingProfileV1 struct {
	Response *ResponseApplicationPolicyDeleteApplicationPolicyQueuingProfileV1Response `json:"response,omitempty"` //
	Version  string                                                                    `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyDeleteApplicationPolicyQueuingProfileV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task id
	URL    string `json:"url,omitempty"`    // Task url
}
type ResponseApplicationPolicyGetApplicationSetsV1 struct {
	Response *[]ResponseApplicationPolicyGetApplicationSetsV1Response `json:"response,omitempty"` //
}
type ResponseApplicationPolicyGetApplicationSetsV1Response struct {
	ID             string                                                               `json:"id,omitempty"`             // Id
	IDentitySource *ResponseApplicationPolicyGetApplicationSetsV1ResponseIDentitySource `json:"identitySource,omitempty"` //
	Name           string                                                               `json:"name,omitempty"`           // Name
}
type ResponseApplicationPolicyGetApplicationSetsV1ResponseIDentitySource struct {
	ID   string `json:"id,omitempty"`   // Id
	Type string `json:"type,omitempty"` // Type
}
type ResponseApplicationPolicyDeleteApplicationSetV1 struct {
	Response *ResponseApplicationPolicyDeleteApplicationSetV1Response `json:"response,omitempty"` //
	Version  string                                                   `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyDeleteApplicationSetV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseApplicationPolicyCreateApplicationSetV1 struct {
	Response *ResponseApplicationPolicyCreateApplicationSetV1Response `json:"response,omitempty"` //
	Version  string                                                   `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyCreateApplicationSetV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseApplicationPolicyGetApplicationSetsCountV1 struct {
	Response string `json:"response,omitempty"` // Response
	Version  string `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyCreateApplicationV1 struct {
	Response *ResponseApplicationPolicyCreateApplicationV1Response `json:"response,omitempty"` //
	Version  string                                                `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyCreateApplicationV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseApplicationPolicyEditApplicationV1 struct {
	Response *ResponseApplicationPolicyEditApplicationV1Response `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyEditApplicationV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseApplicationPolicyDeleteApplicationV1 struct {
	Response *ResponseApplicationPolicyDeleteApplicationV1Response `json:"response,omitempty"` //
	Version  string                                                `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyDeleteApplicationV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseApplicationPolicyGetApplicationsV1 []ResponseItemApplicationPolicyGetApplicationsV1 // Array of ResponseApplicationPolicyGetApplicationsV1
type ResponseItemApplicationPolicyGetApplicationsV1 struct {
	ID                  string                                                               `json:"id,omitempty"`                  // Id
	Name                string                                                               `json:"name,omitempty"`                // Name
	NetworkApplications *[]ResponseItemApplicationPolicyGetApplicationsV1NetworkApplications `json:"networkApplications,omitempty"` //
	NetworkIDentity     *[]ResponseItemApplicationPolicyGetApplicationsV1NetworkIDentity     `json:"networkIdentity,omitempty"`     //
	ApplicationSet      *ResponseItemApplicationPolicyGetApplicationsV1ApplicationSet        `json:"applicationSet,omitempty"`      //
}
type ResponseItemApplicationPolicyGetApplicationsV1NetworkApplications struct {
	ID                 string `json:"id,omitempty"`                 // Id
	AppProtocol        string `json:"appProtocol,omitempty"`        // App Protocol
	ApplicationSubType string `json:"applicationSubType,omitempty"` // Application Sub Type
	ApplicationType    string `json:"applicationType,omitempty"`    // Application Type
	CategoryID         string `json:"categoryId,omitempty"`         // Category Id
	DisplayName        string `json:"displayName,omitempty"`        // Display Name
	EngineID           string `json:"engineId,omitempty"`           // Engine Id
	HelpString         string `json:"helpString,omitempty"`         // Help String
	LongDescription    string `json:"longDescription,omitempty"`    // Long Description
	Name               string `json:"name,omitempty"`               // Name
	Popularity         string `json:"popularity,omitempty"`         // Popularity
	Rank               string `json:"rank,omitempty"`               // Rank
	TrafficClass       string `json:"trafficClass,omitempty"`       // Traffic Class
	ServerName         string `json:"serverName,omitempty"`         // Server Name
	URL                string `json:"url,omitempty"`                // Url
	Dscp               string `json:"dscp,omitempty"`               // Dscp
	IgnoreConflict     string `json:"ignoreConflict,omitempty"`     // Ignore Conflict
}
type ResponseItemApplicationPolicyGetApplicationsV1NetworkIDentity struct {
	ID          string `json:"id,omitempty"`          // Id
	DisplayName string `json:"displayName,omitempty"` // Display Name
	LowerPort   string `json:"lowerPort,omitempty"`   // Lower Port
	Ports       string `json:"ports,omitempty"`       // Ports
	Protocol    string `json:"protocol,omitempty"`    // Protocol
	UpperPort   string `json:"upperPort,omitempty"`   // Upper Port
}
type ResponseItemApplicationPolicyGetApplicationsV1ApplicationSet struct {
	IDRef string `json:"idRef,omitempty"` // Id Ref
}
type ResponseApplicationPolicyGetApplicationsCountV1 struct {
	Response string `json:"response,omitempty"` // Response
	Version  string `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyGetQosDeviceInterfaceInfoV1 struct {
	Response *[]ResponseApplicationPolicyGetQosDeviceInterfaceInfoV1Response `json:"response,omitempty"` //
	Version  string                                                          `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyGetQosDeviceInterfaceInfoV1Response struct {
	ID                     string                                                                                `json:"id,omitempty"`                     // Id of Qos device info
	InstanceID             *int                                                                                  `json:"instanceId,omitempty"`             // Instance id
	DisplayName            string                                                                                `json:"displayName,omitempty"`            // Display name
	InstanceCreatedOn      *int                                                                                  `json:"instanceCreatedOn,omitempty"`      // Instance created on
	InstanceUpdatedOn      *int                                                                                  `json:"instanceUpdatedOn,omitempty"`      // Instance updated on
	InstanceVersion        *int                                                                                  `json:"instanceVersion,omitempty"`        // Instance version
	CreateTime             *int                                                                                  `json:"createTime,omitempty"`             // Create time
	Deployed               *bool                                                                                 `json:"deployed,omitempty"`               // Deployed
	IsSeeded               *bool                                                                                 `json:"isSeeded,omitempty"`               // Is seeded
	IsStale                *bool                                                                                 `json:"isStale,omitempty"`                // Is stale
	LastUpdateTime         *int                                                                                  `json:"lastUpdateTime,omitempty"`         // Last update time
	Name                   string                                                                                `json:"name,omitempty"`                   // Device name
	Namespace              string                                                                                `json:"namespace,omitempty"`              // Namespace
	ProvisioningState      string                                                                                `json:"provisioningState,omitempty"`      // Provisioning state
	Qualifier              string                                                                                `json:"qualifier,omitempty"`              // Qualifier
	ResourceVersion        *int                                                                                  `json:"resourceVersion,omitempty"`        // Resource version
	TargetIDList           *[]ResponseApplicationPolicyGetQosDeviceInterfaceInfoV1ResponseTargetIDList           `json:"targetIdList,omitempty"`           // Target id list
	Type                   string                                                                                `json:"type,omitempty"`                   // Type
	CfsChangeInfo          *[]ResponseApplicationPolicyGetQosDeviceInterfaceInfoV1ResponseCfsChangeInfo          `json:"cfsChangeInfo,omitempty"`          // Cfs change info
	CustomProvisions       *[]ResponseApplicationPolicyGetQosDeviceInterfaceInfoV1ResponseCustomProvisions       `json:"customProvisions,omitempty"`       // Custom provisions
	ExcludedInterfaces     []string                                                                              `json:"excludedInterfaces,omitempty"`     // excluded interfaces ids
	IsExcluded             *bool                                                                                 `json:"isExcluded,omitempty"`             // Is excluded
	NetworkDeviceID        string                                                                                `json:"networkDeviceId,omitempty"`        // Network device id
	QosDeviceInterfaceInfo *[]ResponseApplicationPolicyGetQosDeviceInterfaceInfoV1ResponseQosDeviceInterfaceInfo `json:"qosDeviceInterfaceInfo,omitempty"` //
}
type ResponseApplicationPolicyGetQosDeviceInterfaceInfoV1ResponseTargetIDList interface{}
type ResponseApplicationPolicyGetQosDeviceInterfaceInfoV1ResponseCfsChangeInfo interface{}
type ResponseApplicationPolicyGetQosDeviceInterfaceInfoV1ResponseCustomProvisions interface{}
type ResponseApplicationPolicyGetQosDeviceInterfaceInfoV1ResponseQosDeviceInterfaceInfo struct {
	ID                 string   `json:"id,omitempty"`                 // Id of Qos device interface info
	InstanceID         *int     `json:"instanceId,omitempty"`         // Instance id
	DisplayName        string   `json:"displayName,omitempty"`        // Display name
	InstanceCreatedOn  *int     `json:"instanceCreatedOn,omitempty"`  // Instance created on
	InstanceUpdatedOn  *int     `json:"instanceUpdatedOn,omitempty"`  // Instance updated on
	InstanceVersion    *float64 `json:"instanceVersion,omitempty"`    // Instance version
	DmvpnRemoteSitesBw *[]int   `json:"dmvpnRemoteSitesBw,omitempty"` // Dmvpn remote sites bandwidth
	DownloadBW         *float64 `json:"downloadBW,omitempty"`         // Download bandwidth
	InterfaceID        string   `json:"interfaceId,omitempty"`        // Interface id
	InterfaceName      string   `json:"interfaceName,omitempty"`      // Interface name
	Label              string   `json:"label,omitempty"`              // SP Profile name
	Role               string   `json:"role,omitempty"`               // Interface role
	UploadBW           *int     `json:"uploadBW,omitempty"`           // Upload bandwidth
}
type ResponseApplicationPolicyUpdateQosDeviceInterfaceInfoV1 struct {
	Response *ResponseApplicationPolicyUpdateQosDeviceInterfaceInfoV1Response `json:"response,omitempty"` //
	Version  string                                                           `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyUpdateQosDeviceInterfaceInfoV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task id
	URL    string `json:"url,omitempty"`    // Task url
}
type ResponseApplicationPolicyCreateQosDeviceInterfaceInfoV1 struct {
	Response *ResponseApplicationPolicyCreateQosDeviceInterfaceInfoV1Response `json:"response,omitempty"` //
	Version  string                                                           `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyCreateQosDeviceInterfaceInfoV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task id
	URL    string `json:"url,omitempty"`    // Task url
}
type ResponseApplicationPolicyGetQosDeviceInterfaceInfoCountV1 struct {
	Response *int   `json:"response,omitempty"` // Total number of Qos Device Interface Info
	Version  string `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyDeleteQosDeviceInterfaceInfoV1 struct {
	Response *ResponseApplicationPolicyDeleteQosDeviceInterfaceInfoV1Response `json:"response,omitempty"` //
	Version  string                                                           `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyDeleteQosDeviceInterfaceInfoV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task id
	URL    string `json:"url,omitempty"`    // Task url
}
type ResponseApplicationPolicyCreateApplicationSetsV2 struct {
	Response *ResponseApplicationPolicyCreateApplicationSetsV2Response `json:"response,omitempty"` //
	Version  string                                                    `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyCreateApplicationSetsV2Response struct {
	TaskID string `json:"taskId,omitempty"` // Task id
	URL    string `json:"url,omitempty"`    // Task url
}
type ResponseApplicationPolicyGetApplicationSetsV2 struct {
	Response *[]ResponseApplicationPolicyGetApplicationSetsV2Response `json:"response,omitempty"` //
	Version  string                                                   `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyGetApplicationSetsV2Response struct {
	ID                          string                                                               `json:"id,omitempty"`                          // Id of Application Set
	InstanceID                  *int                                                                 `json:"instanceId,omitempty"`                  // Instance id
	DisplayName                 string                                                               `json:"displayName,omitempty"`                 // Display name
	InstanceVersion             *float64                                                             `json:"instanceVersion,omitempty"`             // Instance version
	DefaultBusinessRelevance    string                                                               `json:"defaultBusinessRelevance,omitempty"`    // Default business relevance
	IDentitySource              *ResponseApplicationPolicyGetApplicationSetsV2ResponseIDentitySource `json:"identitySource,omitempty"`              //
	Name                        string                                                               `json:"name,omitempty"`                        // Application Set name
	Namespace                   string                                                               `json:"namespace,omitempty"`                   // Namespace, valid value scalablegroup:application
	ScalableGroupExternalHandle string                                                               `json:"scalableGroupExternalHandle,omitempty"` // Scalable group external handle, should be equal to Application Set name
	ScalableGroupType           string                                                               `json:"scalableGroupType,omitempty"`           // Scalable group type, valid value APPLICATION_GROUP
	Type                        string                                                               `json:"type,omitempty"`                        // Type, valid value scalablegroup
}
type ResponseApplicationPolicyGetApplicationSetsV2ResponseIDentitySource struct {
	ID   string `json:"id,omitempty"`   // Id
	Type string `json:"type,omitempty"` // Type of identify source. NBAR: build in Application Set, APIC_EM: custom Application Set
}
type ResponseApplicationPolicyGetApplicationSetCountV2 struct {
	Response *int   `json:"response,omitempty"` // Total number of Application Set
	Version  string `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyDeleteApplicationSetV2 struct {
	Response *ResponseApplicationPolicyDeleteApplicationSetV2Response `json:"response,omitempty"` //
	Version  string                                                   `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyDeleteApplicationSetV2Response struct {
	TaskID string `json:"taskId,omitempty"` // Task id
	URL    string `json:"url,omitempty"`    // Task url
}
type ResponseApplicationPolicyEditApplicationsV2 struct {
	Response *ResponseApplicationPolicyEditApplicationsV2Response `json:"response,omitempty"` //
	Version  string                                               `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyEditApplicationsV2Response struct {
	TaskID string `json:"taskId,omitempty"` // Task id
	URL    string `json:"url,omitempty"`    // Task url
}
type ResponseApplicationPolicyCreateApplicationsV2 struct {
	Response *ResponseApplicationPolicyCreateApplicationsV2Response `json:"response,omitempty"` //
	Version  string                                                 `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyCreateApplicationsV2Response struct {
	TaskID string `json:"taskId,omitempty"` // Task id
	URL    string `json:"url,omitempty"`    // Task url
}
type ResponseApplicationPolicyGetApplicationsV2 struct {
	Response *[]ResponseApplicationPolicyGetApplicationsV2Response `json:"response,omitempty"` //
	Version  string                                                `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyGetApplicationsV2Response struct {
	ID                          string                                                                         `json:"id,omitempty"`                          // Id of Application
	InstanceID                  *int                                                                           `json:"instanceId,omitempty"`                  // Instance id
	DisplayName                 string                                                                         `json:"displayName,omitempty"`                 // Display name
	InstanceVersion             *float64                                                                       `json:"instanceVersion,omitempty"`             // Instance version
	IDentitySource              *ResponseApplicationPolicyGetApplicationsV2ResponseIDentitySource              `json:"identitySource,omitempty"`              //
	IndicativeNetworkIDentity   *[]ResponseApplicationPolicyGetApplicationsV2ResponseIndicativeNetworkIDentity `json:"indicativeNetworkIdentity,omitempty"`   //
	Name                        string                                                                         `json:"name,omitempty"`                        // Application name
	Namespace                   string                                                                         `json:"namespace,omitempty"`                   // Namespace, valid value scalablegroup:application
	NetworkApplications         *[]ResponseApplicationPolicyGetApplicationsV2ResponseNetworkApplications       `json:"networkApplications,omitempty"`         //
	NetworkIDentity             *[]ResponseApplicationPolicyGetApplicationsV2ResponseNetworkIDentity           `json:"networkIdentity,omitempty"`             //
	ParentScalableGroup         *ResponseApplicationPolicyGetApplicationsV2ResponseParentScalableGroup         `json:"parentScalableGroup,omitempty"`         //
	Qualifier                   string                                                                         `json:"qualifier,omitempty"`                   // Qualifier, valid value application
	ScalableGroupExternalHandle string                                                                         `json:"scalableGroupExternalHandle,omitempty"` // Scalable group external handle, should be equal to Application name
	ScalableGroupType           string                                                                         `json:"scalableGroupType,omitempty"`           // Scalable group type, valid value APPLICATION
	Type                        string                                                                         `json:"type,omitempty"`                        // Type, valid value scalablegroup
}
type ResponseApplicationPolicyGetApplicationsV2ResponseIDentitySource struct {
	ID   string `json:"id,omitempty"`   // Id
	Type string `json:"type,omitempty"` // Type of identify source. NBAR: build in Application, APIC_EM: custom Application
}
type ResponseApplicationPolicyGetApplicationsV2ResponseIndicativeNetworkIDentity struct {
	ID          string   `json:"id,omitempty"`          // Id
	DisplayName string   `json:"displayName,omitempty"` // Display name
	LowerPort   *float64 `json:"lowerPort,omitempty"`   // Lower port
	Ports       string   `json:"ports,omitempty"`       // Ports
	Protocol    string   `json:"protocol,omitempty"`    // Protocol
	UpperPort   *float64 `json:"upperPort,omitempty"`   // Upper port
}
type ResponseApplicationPolicyGetApplicationsV2ResponseNetworkApplications struct {
	ID                 string   `json:"id,omitempty"`                 // Id
	AppProtocol        string   `json:"appProtocol,omitempty"`        // App protocol
	ApplicationSubType string   `json:"applicationSubType,omitempty"` // Application sub type, LEARNED: discovered application, NONE: nbar and custom application
	ApplicationType    string   `json:"applicationType,omitempty"`    // Application type, DEFAULT: nbar application, DEFAULT_MODIFIED: nbar modified application, CUSTOM: custom application
	CategoryID         string   `json:"categoryId,omitempty"`         // Category id
	DisplayName        string   `json:"displayName,omitempty"`        // Display name
	Dscp               string   `json:"dscp,omitempty"`               // Dscp
	EngineID           string   `json:"engineId,omitempty"`           // Engine id
	HelpString         string   `json:"helpString,omitempty"`         // Help string
	LongDescription    string   `json:"longDescription,omitempty"`    // Long description
	Name               string   `json:"name,omitempty"`               // Application name
	Popularity         *float64 `json:"popularity,omitempty"`         // Popularity
	Rank               *int     `json:"rank,omitempty"`               // Rank, any value between 1 to 65535
	SelectorID         string   `json:"selectorId,omitempty"`         // Selector id
	ServerName         string   `json:"serverName,omitempty"`         // Server name
	URL                string   `json:"url,omitempty"`                // Url
	TrafficClass       string   `json:"trafficClass,omitempty"`       // Traffic class
}
type ResponseApplicationPolicyGetApplicationsV2ResponseNetworkIDentity struct {
	ID          string                                                                         `json:"id,omitempty"`          // Id
	DisplayName string                                                                         `json:"displayName,omitempty"` // Display name
	IPv4Subnet  []string                                                                       `json:"ipv4Subnet,omitempty"`  // Ipv4 subnet
	IPv6Subnet  *[]ResponseApplicationPolicyGetApplicationsV2ResponseNetworkIDentityIPv6Subnet `json:"ipv6Subnet,omitempty"`  // Ipv6 subnet
	LowerPort   *float64                                                                       `json:"lowerPort,omitempty"`   // Lower port
	Ports       string                                                                         `json:"ports,omitempty"`       // Ports
	Protocol    string                                                                         `json:"protocol,omitempty"`    // Protocol
	UpperPort   *float64                                                                       `json:"upperPort,omitempty"`   // Upper port
}
type ResponseApplicationPolicyGetApplicationsV2ResponseNetworkIDentityIPv6Subnet interface{}
type ResponseApplicationPolicyGetApplicationsV2ResponseParentScalableGroup struct {
	ID    string `json:"id,omitempty"`    // Id
	IDRef string `json:"idRef,omitempty"` // Id reference to parent application set
}
type ResponseApplicationPolicyGetApplicationCountV2 struct {
	Response *int   `json:"response,omitempty"` // Total number of Application
	Version  string `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyDeleteApplicationV2 struct {
	Response *ResponseApplicationPolicyDeleteApplicationV2Response `json:"response,omitempty"` //
	Version  string                                                `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyDeleteApplicationV2Response struct {
	TaskID string `json:"taskId,omitempty"` // Task id
	URL    string `json:"url,omitempty"`    // Task url
}
type RequestApplicationPolicyApplicationPolicyIntentV1 struct {
	CreateList *[]RequestApplicationPolicyApplicationPolicyIntentV1CreateList `json:"createList,omitempty"` //
	UpdateList *[]RequestApplicationPolicyApplicationPolicyIntentV1UpdateList `json:"updateList,omitempty"` //
	DeleteList []string                                                       `json:"deleteList,omitempty"` // Delete list of Group Based Policy ids
}
type RequestApplicationPolicyApplicationPolicyIntentV1CreateList struct {
	Name                string                                                                          `json:"name,omitempty"`                // Concatination of <polcy name>_<application-set-name> or <polcy name>_global_policy_configuration or <polcy name>_queuing_customization
	DeletePolicyStatus  string                                                                          `json:"deletePolicyStatus,omitempty"`  // NONE: deployed policy to devices, DELETED: delete policy from devices, RESTORED: restored to original configuration
	PolicyScope         string                                                                          `json:"policyScope,omitempty"`         // Policy name
	Priority            string                                                                          `json:"priority,omitempty"`            // Set to 4095 while producer refer to application Scalable group otherwise 100
	AdvancedPolicyScope *RequestApplicationPolicyApplicationPolicyIntentV1CreateListAdvancedPolicyScope `json:"advancedPolicyScope,omitempty"` //
	ExclusiveContract   *RequestApplicationPolicyApplicationPolicyIntentV1CreateListExclusiveContract   `json:"exclusiveContract,omitempty"`   //
	Contract            *RequestApplicationPolicyApplicationPolicyIntentV1CreateListContract            `json:"contract,omitempty"`            //
	Producer            *RequestApplicationPolicyApplicationPolicyIntentV1CreateListProducer            `json:"producer,omitempty"`            //
	Consumer            *RequestApplicationPolicyApplicationPolicyIntentV1CreateListConsumer            `json:"consumer,omitempty"`            //
}
type RequestApplicationPolicyApplicationPolicyIntentV1CreateListAdvancedPolicyScope struct {
	Name                       string                                                                                                      `json:"name,omitempty"`                       // Policy name
	AdvancedPolicyScopeElement *[]RequestApplicationPolicyApplicationPolicyIntentV1CreateListAdvancedPolicyScopeAdvancedPolicyScopeElement `json:"advancedPolicyScopeElement,omitempty"` //
}
type RequestApplicationPolicyApplicationPolicyIntentV1CreateListAdvancedPolicyScopeAdvancedPolicyScopeElement struct {
	GroupID []string `json:"groupId,omitempty"` // The site(s) ID where the Application QoS Policy will be deployed.
	SSID    []string `json:"ssid,omitempty"`    // Ssid
}
type RequestApplicationPolicyApplicationPolicyIntentV1CreateListExclusiveContract struct {
	Clause *[]RequestApplicationPolicyApplicationPolicyIntentV1CreateListExclusiveContractClause `json:"clause,omitempty"` //
}
type RequestApplicationPolicyApplicationPolicyIntentV1CreateListExclusiveContractClause struct {
	Type                  string `json:"type,omitempty"`                  // Type
	RelevanceLevel        string `json:"relevanceLevel,omitempty"`        // Relevance level
	DeviceRemovalBehavior string `json:"deviceRemovalBehavior,omitempty"` // Device eemoval behavior
	HostTrackingEnabled   *bool  `json:"hostTrackingEnabled,omitempty"`   // Is host tracking enabled
}
type RequestApplicationPolicyApplicationPolicyIntentV1CreateListContract struct {
	IDRef string `json:"idRef,omitempty"` // Id ref to Queueing profile
}
type RequestApplicationPolicyApplicationPolicyIntentV1CreateListProducer struct {
	ScalableGroup *[]RequestApplicationPolicyApplicationPolicyIntentV1CreateListProducerScalableGroup `json:"scalableGroup,omitempty"` //
}
type RequestApplicationPolicyApplicationPolicyIntentV1CreateListProducerScalableGroup struct {
	IDRef string `json:"idRef,omitempty"` // Id ref to application-set or application Scalable group
}
type RequestApplicationPolicyApplicationPolicyIntentV1CreateListConsumer struct {
	ScalableGroup *[]RequestApplicationPolicyApplicationPolicyIntentV1CreateListConsumerScalableGroup `json:"scalableGroup,omitempty"` //
}
type RequestApplicationPolicyApplicationPolicyIntentV1CreateListConsumerScalableGroup struct {
	IDRef string `json:"idRef,omitempty"` // Id ref to application Scalable group
}
type RequestApplicationPolicyApplicationPolicyIntentV1UpdateList struct {
	ID                  string                                                                          `json:"id,omitempty"`                  // Id of Group based policy
	Name                string                                                                          `json:"name,omitempty"`                // Concatination of <polcy name>_<application-set-name> or <polcy name>_global_policy_configuration or <polcy name>_queuing_customization
	DeletePolicyStatus  string                                                                          `json:"deletePolicyStatus,omitempty"`  // NONE: deployed policy to devices, DELETED: delete policy from devices, RESTORED: restored to original configuration
	PolicyScope         string                                                                          `json:"policyScope,omitempty"`         // Policy name
	Priority            string                                                                          `json:"priority,omitempty"`            // Set to 4095 while producer refer to application Scalable group otherwise 100
	AdvancedPolicyScope *RequestApplicationPolicyApplicationPolicyIntentV1UpdateListAdvancedPolicyScope `json:"advancedPolicyScope,omitempty"` //
	ExclusiveContract   *RequestApplicationPolicyApplicationPolicyIntentV1UpdateListExclusiveContract   `json:"exclusiveContract,omitempty"`   //
	Contract            *RequestApplicationPolicyApplicationPolicyIntentV1UpdateListContract            `json:"contract,omitempty"`            //
	Producer            *RequestApplicationPolicyApplicationPolicyIntentV1UpdateListProducer            `json:"producer,omitempty"`            //
	Consumer            *RequestApplicationPolicyApplicationPolicyIntentV1UpdateListConsumer            `json:"consumer,omitempty"`            //
}
type RequestApplicationPolicyApplicationPolicyIntentV1UpdateListAdvancedPolicyScope struct {
	ID                         string                                                                                                      `json:"id,omitempty"`                         // Id of Advance policy scope
	Name                       string                                                                                                      `json:"name,omitempty"`                       // Policy name
	AdvancedPolicyScopeElement *[]RequestApplicationPolicyApplicationPolicyIntentV1UpdateListAdvancedPolicyScopeAdvancedPolicyScopeElement `json:"advancedPolicyScopeElement,omitempty"` //
}
type RequestApplicationPolicyApplicationPolicyIntentV1UpdateListAdvancedPolicyScopeAdvancedPolicyScopeElement struct {
	ID      string   `json:"id,omitempty"`      // Id of Advance policy scope element
	GroupID []string `json:"groupId,omitempty"` // The site(s) ID where the Application QoS Policy will be deployed.
	SSID    []string `json:"ssid,omitempty"`    // Ssid
}
type RequestApplicationPolicyApplicationPolicyIntentV1UpdateListExclusiveContract struct {
	ID     string                                                                                `json:"id,omitempty"`     // Id of Exclusive contract
	Clause *[]RequestApplicationPolicyApplicationPolicyIntentV1UpdateListExclusiveContractClause `json:"clause,omitempty"` //
}
type RequestApplicationPolicyApplicationPolicyIntentV1UpdateListExclusiveContractClause struct {
	ID                    string `json:"id,omitempty"`                    // Id of Business relevance or Application policy knobs clause
	Type                  string `json:"type,omitempty"`                  // Type
	RelevanceLevel        string `json:"relevanceLevel,omitempty"`        // Relevance level
	DeviceRemovalBehavior string `json:"deviceRemovalBehavior,omitempty"` // Device removal behavior
	HostTrackingEnabled   *bool  `json:"hostTrackingEnabled,omitempty"`   // Host tracking enabled
}
type RequestApplicationPolicyApplicationPolicyIntentV1UpdateListContract struct {
	IDRef string `json:"idRef,omitempty"` // Id ref to Queueing profile
}
type RequestApplicationPolicyApplicationPolicyIntentV1UpdateListProducer struct {
	ID            string                                                                              `json:"id,omitempty"`            // Id of Producer
	ScalableGroup *[]RequestApplicationPolicyApplicationPolicyIntentV1UpdateListProducerScalableGroup `json:"scalableGroup,omitempty"` //
}
type RequestApplicationPolicyApplicationPolicyIntentV1UpdateListProducerScalableGroup struct {
	IDRef string `json:"idRef,omitempty"` // Id ref to application-set or application Scalable group
}
type RequestApplicationPolicyApplicationPolicyIntentV1UpdateListConsumer struct {
	ID            string                                                                              `json:"id,omitempty"`            // Id of Consumer
	ScalableGroup *[]RequestApplicationPolicyApplicationPolicyIntentV1UpdateListConsumerScalableGroup `json:"scalableGroup,omitempty"` //
}
type RequestApplicationPolicyApplicationPolicyIntentV1UpdateListConsumerScalableGroup struct {
	IDRef string `json:"idRef,omitempty"` // Id ref to application Scalable group
}
type RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileV1 []RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileV1 // Array of RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileV1
type RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileV1 struct {
	ID          string                                                                       `json:"id,omitempty"`          // Id of Queueing profile
	Description string                                                                       `json:"description,omitempty"` // Free test description
	Name        string                                                                       `json:"name,omitempty"`        // Queueing profile name
	Clause      *[]RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileV1Clause `json:"clause,omitempty"`      //
}
type RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileV1Clause struct {
	InstanceID                        *int                                                                                                       `json:"instanceId,omitempty"`                        // Instance id
	Type                              string                                                                                                     `json:"type,omitempty"`                              // The allowed clause types are: BANDWIDTH, DSCP_CUSTOMIZATION
	IsCommonBetweenAllInterfaceSpeeds *bool                                                                                                      `json:"isCommonBetweenAllInterfaceSpeeds,omitempty"` // Is common between all interface speeds
	InterfaceSpeedBandwidthClauses    *[]RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileV1ClauseInterfaceSpeedBandwidthClauses `json:"interfaceSpeedBandwidthClauses,omitempty"`    //
	TcDscpSettings                    *[]RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileV1ClauseTcDscpSettings                 `json:"tcDscpSettings,omitempty"`                    //
}
type RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileV1ClauseInterfaceSpeedBandwidthClauses struct {
	InstanceID          *int                                                                                                                          `json:"instanceId,omitempty"`          // Instance id
	InterfaceSpeed      string                                                                                                                        `json:"interfaceSpeed,omitempty"`      // Interface speed
	TcBandwidthSettings *[]RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileV1ClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings `json:"tcBandwidthSettings,omitempty"` //
}
type RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileV1ClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings struct {
	InstanceID          *int   `json:"instanceId,omitempty"`          // Instance id
	BandwidthPercentage *int   `json:"bandwidthPercentage,omitempty"` // Bandwidth percentage
	TrafficClass        string `json:"trafficClass,omitempty"`        // Traffic Class
}
type RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileV1ClauseTcDscpSettings struct {
	InstanceID   *int   `json:"instanceId,omitempty"`   // Instance id
	Dscp         string `json:"dscp,omitempty"`         // Dscp value
	TrafficClass string `json:"trafficClass,omitempty"` // Traffic Class
}
type RequestApplicationPolicyCreateApplicationPolicyQueuingProfileV1 []RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileV1 // Array of RequestApplicationPolicyCreateApplicationPolicyQueuingProfileV1
type RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileV1 struct {
	Description string                                                                       `json:"description,omitempty"` // Free test description
	Name        string                                                                       `json:"name,omitempty"`        // Queueing profile name
	Clause      *[]RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileV1Clause `json:"clause,omitempty"`      //
}
type RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileV1Clause struct {
	Type                              string                                                                                                     `json:"type,omitempty"`                              // The allowed clause types are: BANDWIDTH, DSCP_CUSTOMIZATION
	IsCommonBetweenAllInterfaceSpeeds *bool                                                                                                      `json:"isCommonBetweenAllInterfaceSpeeds,omitempty"` // Is common between all interface speeds
	InterfaceSpeedBandwidthClauses    *[]RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileV1ClauseInterfaceSpeedBandwidthClauses `json:"interfaceSpeedBandwidthClauses,omitempty"`    //
	TcDscpSettings                    *[]RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileV1ClauseTcDscpSettings                 `json:"tcDscpSettings,omitempty"`                    //
}
type RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileV1ClauseInterfaceSpeedBandwidthClauses struct {
	InterfaceSpeed      string                                                                                                                        `json:"interfaceSpeed,omitempty"`      // Interface speed
	TcBandwidthSettings *[]RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileV1ClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings `json:"tcBandwidthSettings,omitempty"` //
}
type RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileV1ClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings struct {
	BandwidthPercentage *int   `json:"bandwidthPercentage,omitempty"` // Bandwidth percentage
	TrafficClass        string `json:"trafficClass,omitempty"`        // Traffic Class
}
type RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileV1ClauseTcDscpSettings struct {
	Dscp         string `json:"dscp,omitempty"`         // Dscp value
	TrafficClass string `json:"trafficClass,omitempty"` // Traffic Class
}
type RequestApplicationPolicyCreateApplicationSetV1 []RequestItemApplicationPolicyCreateApplicationSetV1 // Array of RequestApplicationPolicyCreateApplicationSetV1
type RequestItemApplicationPolicyCreateApplicationSetV1 struct {
	Name string `json:"name,omitempty"` // Name
}
type RequestApplicationPolicyCreateApplicationV1 []RequestItemApplicationPolicyCreateApplicationV1 // Array of RequestApplicationPolicyCreateApplicationV1
type RequestItemApplicationPolicyCreateApplicationV1 struct {
	Name                string                                                                `json:"name,omitempty"`                // Name
	NetworkApplications *[]RequestItemApplicationPolicyCreateApplicationV1NetworkApplications `json:"networkApplications,omitempty"` //
	NetworkIDentity     *[]RequestItemApplicationPolicyCreateApplicationV1NetworkIDentity     `json:"networkIdentity,omitempty"`     //
	ApplicationSet      *RequestItemApplicationPolicyCreateApplicationV1ApplicationSet        `json:"applicationSet,omitempty"`      //
}
type RequestItemApplicationPolicyCreateApplicationV1NetworkApplications struct {
	AppProtocol        string `json:"appProtocol,omitempty"`        // App Protocol
	ApplicationSubType string `json:"applicationSubType,omitempty"` // Application Sub Type
	ApplicationType    string `json:"applicationType,omitempty"`    // Application Type
	CategoryID         string `json:"categoryId,omitempty"`         // Category Id
	DisplayName        string `json:"displayName,omitempty"`        // Display Name
	EngineID           string `json:"engineId,omitempty"`           // Engine Id
	HelpString         string `json:"helpString,omitempty"`         // Help String
	LongDescription    string `json:"longDescription,omitempty"`    // Long Description
	Name               string `json:"name,omitempty"`               // Name
	Popularity         string `json:"popularity,omitempty"`         // Popularity
	Rank               string `json:"rank,omitempty"`               // Rank
	TrafficClass       string `json:"trafficClass,omitempty"`       // Traffic Class
	ServerName         string `json:"serverName,omitempty"`         // Server Name
	URL                string `json:"url,omitempty"`                // Url
	Dscp               string `json:"dscp,omitempty"`               // Dscp
	IgnoreConflict     string `json:"ignoreConflict,omitempty"`     // Ignore Conflict
}
type RequestItemApplicationPolicyCreateApplicationV1NetworkIDentity struct {
	DisplayName string `json:"displayName,omitempty"` // Display Name
	LowerPort   string `json:"lowerPort,omitempty"`   // Lower Port
	Ports       string `json:"ports,omitempty"`       // Ports
	Protocol    string `json:"protocol,omitempty"`    // Protocol
	UpperPort   string `json:"upperPort,omitempty"`   // Upper Port
}
type RequestItemApplicationPolicyCreateApplicationV1ApplicationSet struct {
	IDRef string `json:"idRef,omitempty"` // Id Ref
}
type RequestApplicationPolicyEditApplicationV1 []RequestItemApplicationPolicyEditApplicationV1 // Array of RequestApplicationPolicyEditApplicationV1
type RequestItemApplicationPolicyEditApplicationV1 struct {
	ID                  string                                                              `json:"id,omitempty"`                  // Id
	Name                string                                                              `json:"name,omitempty"`                // Name
	NetworkApplications *[]RequestItemApplicationPolicyEditApplicationV1NetworkApplications `json:"networkApplications,omitempty"` //
	NetworkIDentity     *[]RequestItemApplicationPolicyEditApplicationV1NetworkIDentity     `json:"networkIdentity,omitempty"`     //
	ApplicationSet      *RequestItemApplicationPolicyEditApplicationV1ApplicationSet        `json:"applicationSet,omitempty"`      //
}
type RequestItemApplicationPolicyEditApplicationV1NetworkApplications struct {
	ID                 string `json:"id,omitempty"`                 // Id
	AppProtocol        string `json:"appProtocol,omitempty"`        // App Protocol
	ApplicationSubType string `json:"applicationSubType,omitempty"` // Application Sub Type
	ApplicationType    string `json:"applicationType,omitempty"`    // Application Type
	CategoryID         string `json:"categoryId,omitempty"`         // Category Id
	DisplayName        string `json:"displayName,omitempty"`        // Display Name
	EngineID           string `json:"engineId,omitempty"`           // Engine Id
	HelpString         string `json:"helpString,omitempty"`         // Help String
	LongDescription    string `json:"longDescription,omitempty"`    // Long Description
	Name               string `json:"name,omitempty"`               // Name
	Popularity         string `json:"popularity,omitempty"`         // Popularity
	Rank               string `json:"rank,omitempty"`               // Rank
	TrafficClass       string `json:"trafficClass,omitempty"`       // Traffic Class
	ServerName         string `json:"serverName,omitempty"`         // Server Name
	URL                string `json:"url,omitempty"`                // Url
	Dscp               string `json:"dscp,omitempty"`               // Dscp
	IgnoreConflict     string `json:"ignoreConflict,omitempty"`     // Ignore Conflict
}
type RequestItemApplicationPolicyEditApplicationV1NetworkIDentity struct {
	ID          string `json:"id,omitempty"`          // Id
	DisplayName string `json:"displayName,omitempty"` // Display Name
	LowerPort   string `json:"lowerPort,omitempty"`   // Lower Port
	Ports       string `json:"ports,omitempty"`       // Ports
	Protocol    string `json:"protocol,omitempty"`    // Protocol
	UpperPort   string `json:"upperPort,omitempty"`   // Upper Port
}
type RequestItemApplicationPolicyEditApplicationV1ApplicationSet struct {
	IDRef string `json:"idRef,omitempty"` // Id Ref
}
type RequestApplicationPolicyUpdateQosDeviceInterfaceInfoV1 []RequestItemApplicationPolicyUpdateQosDeviceInterfaceInfoV1 // Array of RequestApplicationPolicyUpdateQosDeviceInterfaceInfoV1
type RequestItemApplicationPolicyUpdateQosDeviceInterfaceInfoV1 struct {
	ID                     string                                                                              `json:"id,omitempty"`                     // Id of Qos device info
	Name                   string                                                                              `json:"name,omitempty"`                   // Device name
	ExcludedInterfaces     []string                                                                            `json:"excludedInterfaces,omitempty"`     // Excluded interfaces ids
	NetworkDeviceID        string                                                                              `json:"networkDeviceId,omitempty"`        // Network device id
	QosDeviceInterfaceInfo *[]RequestItemApplicationPolicyUpdateQosDeviceInterfaceInfoV1QosDeviceInterfaceInfo `json:"qosDeviceInterfaceInfo,omitempty"` //
}
type RequestItemApplicationPolicyUpdateQosDeviceInterfaceInfoV1QosDeviceInterfaceInfo struct {
	InstanceID         *int   `json:"instanceId,omitempty"`         // Instance id
	DmvpnRemoteSitesBw *[]int `json:"dmvpnRemoteSitesBw,omitempty"` // Dmvpn remote sites bandwidth
	InterfaceID        string `json:"interfaceId,omitempty"`        // Interface id
	InterfaceName      string `json:"interfaceName,omitempty"`      // Interface name
	Label              string `json:"label,omitempty"`              // SP Profile name
	Role               string `json:"role,omitempty"`               // Interface role
	UploadBW           *int   `json:"uploadBW,omitempty"`           // Upload bandwidth
}
type RequestApplicationPolicyCreateQosDeviceInterfaceInfoV1 []RequestItemApplicationPolicyCreateQosDeviceInterfaceInfoV1 // Array of RequestApplicationPolicyCreateQosDeviceInterfaceInfoV1
type RequestItemApplicationPolicyCreateQosDeviceInterfaceInfoV1 struct {
	Name                   string                                                                              `json:"name,omitempty"`                   // Device name
	ExcludedInterfaces     []string                                                                            `json:"excludedInterfaces,omitempty"`     // Excluded interfaces ids
	NetworkDeviceID        string                                                                              `json:"networkDeviceId,omitempty"`        // Network device id
	QosDeviceInterfaceInfo *[]RequestItemApplicationPolicyCreateQosDeviceInterfaceInfoV1QosDeviceInterfaceInfo `json:"qosDeviceInterfaceInfo,omitempty"` //
}
type RequestItemApplicationPolicyCreateQosDeviceInterfaceInfoV1QosDeviceInterfaceInfo struct {
	DmvpnRemoteSitesBw *[]int `json:"dmvpnRemoteSitesBw,omitempty"` // Dmvpn remote sites bandwidth
	InterfaceID        string `json:"interfaceId,omitempty"`        // Interface id
	InterfaceName      string `json:"interfaceName,omitempty"`      // Interface name
	Label              string `json:"label,omitempty"`              // SP Profile name
	Role               string `json:"role,omitempty"`               // Interface role
	UploadBW           *int   `json:"uploadBW,omitempty"`           // Upload bandwidth
}
type RequestApplicationPolicyCreateApplicationSetsV2 []RequestItemApplicationPolicyCreateApplicationSetsV2 // Array of RequestApplicationPolicyCreateApplicationSetsV2
type RequestItemApplicationPolicyCreateApplicationSetsV2 struct {
	Name                        string `json:"name,omitempty"`                        // Application Set name
	ScalableGroupType           string `json:"scalableGroupType,omitempty"`           // Scalable group type, should be set to APPLICATION_GROUP
	DefaultBusinessRelevance    string `json:"defaultBusinessRelevance,omitempty"`    // Default business relevance
	Namespace                   string `json:"namespace,omitempty"`                   // Namespace, should be set to scalablegroup:application
	Qualifier                   string `json:"qualifier,omitempty"`                   // Qualifier, should be set to application
	Type                        string `json:"type,omitempty"`                        // Type, should be set to scalablegroup
	ScalableGroupExternalHandle string `json:"scalableGroupExternalHandle,omitempty"` // Scalable group external handle, should be set to application set name
}
type RequestApplicationPolicyEditApplicationsV2 []RequestItemApplicationPolicyEditApplicationsV2 // Array of RequestApplicationPolicyEditApplicationsV2
type RequestItemApplicationPolicyEditApplicationsV2 struct {
	ID                          string                                                                     `json:"id,omitempty"`                          // Application id
	InstanceID                  *int                                                                       `json:"instanceId,omitempty"`                  // Instance id
	DisplayName                 string                                                                     `json:"displayName,omitempty"`                 // Display name
	InstanceVersion             *float64                                                                   `json:"instanceVersion,omitempty"`             // Instance version
	IndicativeNetworkIDentity   *[]RequestItemApplicationPolicyEditApplicationsV2IndicativeNetworkIDentity `json:"indicativeNetworkIdentity,omitempty"`   //
	Name                        string                                                                     `json:"name,omitempty"`                        // Application name
	Namespace                   string                                                                     `json:"namespace,omitempty"`                   // Namespace
	NetworkApplications         *[]RequestItemApplicationPolicyEditApplicationsV2NetworkApplications       `json:"networkApplications,omitempty"`         //
	NetworkIDentity             *[]RequestItemApplicationPolicyEditApplicationsV2NetworkIDentity           `json:"networkIdentity,omitempty"`             //
	ParentScalableGroup         *RequestItemApplicationPolicyEditApplicationsV2ParentScalableGroup         `json:"parentScalableGroup,omitempty"`         //
	Qualifier                   string                                                                     `json:"qualifier,omitempty"`                   // Qualifier, valid value application
	ScalableGroupExternalHandle string                                                                     `json:"scalableGroupExternalHandle,omitempty"` // Scalable group external handle, should be equal to Application name
	ScalableGroupType           string                                                                     `json:"scalableGroupType,omitempty"`           // Scalable group type, valid value APPLICATION
	Type                        string                                                                     `json:"type,omitempty"`                        // Type, valid value scalablegroup
}
type RequestItemApplicationPolicyEditApplicationsV2IndicativeNetworkIDentity struct {
	ID          string   `json:"id,omitempty"`          // Id
	DisplayName string   `json:"displayName,omitempty"` // Display name
	LowerPort   *float64 `json:"lowerPort,omitempty"`   // Lower port
	Ports       string   `json:"ports,omitempty"`       // Ports
	Protocol    string   `json:"protocol,omitempty"`    // Protocol
	UpperPort   *float64 `json:"upperPort,omitempty"`   // Upper port
}
type RequestItemApplicationPolicyEditApplicationsV2NetworkApplications struct {
	ID                 string   `json:"id,omitempty"`                 // Id
	ApplicationSubType string   `json:"applicationSubType,omitempty"` // Application sub type, LEARNED: discovered application, NONE: nbar and custom application
	ApplicationType    string   `json:"applicationType,omitempty"`    // Application type, DEFAULT: nbar application, DEFAULT_MODIFIED: nbar modified application, CUSTOM: custom application
	CategoryID         string   `json:"categoryId,omitempty"`         // Category id
	DisplayName        string   `json:"displayName,omitempty"`        // Display name
	EngineID           string   `json:"engineId,omitempty"`           // Engine id
	HelpString         string   `json:"helpString,omitempty"`         // Help string
	LongDescription    string   `json:"longDescription,omitempty"`    // Long description
	Name               string   `json:"name,omitempty"`               // Application name
	Popularity         *float64 `json:"popularity,omitempty"`         // Popularity
	Rank               *int     `json:"rank,omitempty"`               // Rank, any value between 1 to 65535
	SelectorID         string   `json:"selectorId,omitempty"`         // Selector id
	Dscp               string   `json:"dscp,omitempty"`               // Dscp
	AppProtocol        string   `json:"appProtocol,omitempty"`        // App protocol
	ServerName         string   `json:"serverName,omitempty"`         // Server name
	URL                string   `json:"url,omitempty"`                // Url
	TrafficClass       string   `json:"trafficClass,omitempty"`       // Traffic class
	IgnoreConflict     *bool    `json:"ignoreConflict,omitempty"`     // Ignore conflict, true or false
}
type RequestItemApplicationPolicyEditApplicationsV2NetworkIDentity struct {
	ID          string                                                                     `json:"id,omitempty"`          // Id
	DisplayName string                                                                     `json:"displayName,omitempty"` // Display name
	IPv4Subnet  []string                                                                   `json:"ipv4Subnet,omitempty"`  // Ipv4 subnet
	IPv6Subnet  *[]RequestItemApplicationPolicyEditApplicationsV2NetworkIDentityIPv6Subnet `json:"ipv6Subnet,omitempty"`  // Ipv6 subnet
	LowerPort   *float64                                                                   `json:"lowerPort,omitempty"`   // Lower port
	Ports       string                                                                     `json:"ports,omitempty"`       // Ports
	Protocol    string                                                                     `json:"protocol,omitempty"`    // Protocol
	UpperPort   *float64                                                                   `json:"upperPort,omitempty"`   // Upper port
}
type RequestItemApplicationPolicyEditApplicationsV2NetworkIDentityIPv6Subnet interface{}
type RequestItemApplicationPolicyEditApplicationsV2ParentScalableGroup struct {
	IDRef string `json:"idRef,omitempty"` // Id reference to parent application set
}
type RequestApplicationPolicyCreateApplicationsV2 []RequestItemApplicationPolicyCreateApplicationsV2 // Array of RequestApplicationPolicyCreateApplicationsV2
type RequestItemApplicationPolicyCreateApplicationsV2 struct {
	Name                      string                                                                       `json:"name,omitempty"`                      // Application name
	NetworkApplications       *[]RequestItemApplicationPolicyCreateApplicationsV2NetworkApplications       `json:"networkApplications,omitempty"`       //
	ParentScalableGroup       *RequestItemApplicationPolicyCreateApplicationsV2ParentScalableGroup         `json:"parentScalableGroup,omitempty"`       //
	NetworkIDentity           *[]RequestItemApplicationPolicyCreateApplicationsV2NetworkIDentity           `json:"networkIdentity,omitempty"`           //
	IndicativeNetworkIDentity *[]RequestItemApplicationPolicyCreateApplicationsV2IndicativeNetworkIDentity `json:"indicativeNetworkIdentity,omitempty"` //
	ScalableGroupType         string                                                                       `json:"scalableGroupType,omitempty"`         // Scalable group type, valid value APPLICATION
	Type                      string                                                                       `json:"type,omitempty"`                      // Type, valid value scalablegroup
}
type RequestItemApplicationPolicyCreateApplicationsV2NetworkApplications struct {
	HelpString      string `json:"helpString,omitempty"`      // Help string
	ApplicationType string `json:"applicationType,omitempty"` // Application type
	Type            string `json:"type,omitempty"`            // Custom application type
	Dscp            string `json:"dscp,omitempty"`            // Dscp, valid only in case of _server-ip custom application type
	AppProtocol     string `json:"appProtocol,omitempty"`     // App protocol, in case of _servername should not be set, in case of _url should be set to TCP
	ServerName      string `json:"serverName,omitempty"`      // Server name, should be set only in case of _servername
	URL             string `json:"url,omitempty"`             // Url, should be set only in case of _url
	TrafficClass    string `json:"trafficClass,omitempty"`    // Traffic class
	CategoryID      string `json:"categoryId,omitempty"`      // Category id
	IgnoreConflict  *bool  `json:"ignoreConflict,omitempty"`  // Ignore conflict, true or false
	Rank            *int   `json:"rank,omitempty"`            // Rank, should be set to 1
	EngineID        *int   `json:"engineId,omitempty"`        // Engine id, should be set to 6
}
type RequestItemApplicationPolicyCreateApplicationsV2ParentScalableGroup struct {
	IDRef string `json:"idRef,omitempty"` // Id reference to parent application set
}
type RequestItemApplicationPolicyCreateApplicationsV2NetworkIDentity struct {
	Protocol   string   `json:"protocol,omitempty"`   // Protocol
	Ports      string   `json:"ports,omitempty"`      // Ports
	IPv4Subnet []string `json:"ipv4Subnet,omitempty"` // Ipv4 subnet
	LowerPort  *float64 `json:"lowerPort,omitempty"`  // Lower port
	UpperPort  *float64 `json:"upperPort,omitempty"`  // Upper port
}
type RequestItemApplicationPolicyCreateApplicationsV2IndicativeNetworkIDentity struct {
	Protocol   string   `json:"protocol,omitempty"`   // Protocol
	Ports      string   `json:"ports,omitempty"`      // Ports
	IPv4Subnet []string `json:"ipv4Subnet,omitempty"` // Ipv4 subnet
	IPv6Subnet []string `json:"ipv6Subnet,omitempty"` // Ipv6 subnet
	LowerPort  *float64 `json:"lowerPort,omitempty"`  // The minimum port when used as a port range. For single port number, ports attribute should be used.
	UpperPort  *float64 `json:"upperPort,omitempty"`  // The maximum port when used as a port range. For single port number, ports attribute should be used.
}

//GetApplicationPolicyV1 Get Application Policy - 3d9f-6b17-4879-8e45
/* Get all existing application policies


@param GetApplicationPolicyV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-application-policy-v1
*/
func (s *ApplicationPolicyService) GetApplicationPolicyV1(GetApplicationPolicyV1QueryParams *GetApplicationPolicyV1QueryParams) (*ResponseApplicationPolicyGetApplicationPolicyV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/app-policy"

	queryString, _ := query.Values(GetApplicationPolicyV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplicationPolicyGetApplicationPolicyV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApplicationPolicyV1(GetApplicationPolicyV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetApplicationPolicyV1")
	}

	result := response.Result().(*ResponseApplicationPolicyGetApplicationPolicyV1)
	return result, response, err

}

//GetApplicationPolicyDefaultV1 Get Application Policy Default - 21a2-4b5d-4f98-a730
/* Get default application policy



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-application-policy-default-v1
*/
func (s *ApplicationPolicyService) GetApplicationPolicyDefaultV1() (*ResponseApplicationPolicyGetApplicationPolicyDefaultV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/app-policy-default"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplicationPolicyGetApplicationPolicyDefaultV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApplicationPolicyDefaultV1()
		}
		return nil, response, fmt.Errorf("error with operation GetApplicationPolicyDefaultV1")
	}

	result := response.Result().(*ResponseApplicationPolicyGetApplicationPolicyDefaultV1)
	return result, response, err

}

//GetApplicationPolicyQueuingProfileV1 Get Application Policy Queuing Profile - 1698-39cc-4bdb-8ed9
/* Get all or by name, existing application policy queuing profiles


@param GetApplicationPolicyQueuingProfileV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-application-policy-queuing-profile-v1
*/
func (s *ApplicationPolicyService) GetApplicationPolicyQueuingProfileV1(GetApplicationPolicyQueuingProfileV1QueryParams *GetApplicationPolicyQueuingProfileV1QueryParams) (*ResponseApplicationPolicyGetApplicationPolicyQueuingProfileV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/app-policy-queuing-profile"

	queryString, _ := query.Values(GetApplicationPolicyQueuingProfileV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplicationPolicyGetApplicationPolicyQueuingProfileV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApplicationPolicyQueuingProfileV1(GetApplicationPolicyQueuingProfileV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetApplicationPolicyQueuingProfileV1")
	}

	result := response.Result().(*ResponseApplicationPolicyGetApplicationPolicyQueuingProfileV1)
	return result, response, err

}

//GetApplicationPolicyQueuingProfileCountV1 Get Application Policy Queuing Profile Count - efa2-4afa-4578-88e9
/* Get the number of all existing  application policy queuing profile



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-application-policy-queuing-profile-count-v1
*/
func (s *ApplicationPolicyService) GetApplicationPolicyQueuingProfileCountV1() (*ResponseApplicationPolicyGetApplicationPolicyQueuingProfileCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/app-policy-queuing-profile-count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplicationPolicyGetApplicationPolicyQueuingProfileCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApplicationPolicyQueuingProfileCountV1()
		}
		return nil, response, fmt.Errorf("error with operation GetApplicationPolicyQueuingProfileCountV1")
	}

	result := response.Result().(*ResponseApplicationPolicyGetApplicationPolicyQueuingProfileCountV1)
	return result, response, err

}

//GetApplicationSetsV1 Get Application Sets - cb86-8b21-4289-8159
/* Get appllication-sets by offset/limit or by name


@param GetApplicationSetsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-application-sets-v1
*/
func (s *ApplicationPolicyService) GetApplicationSetsV1(GetApplicationSetsV1QueryParams *GetApplicationSetsV1QueryParams) (*ResponseApplicationPolicyGetApplicationSetsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/application-policy-application-set"

	queryString, _ := query.Values(GetApplicationSetsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplicationPolicyGetApplicationSetsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApplicationSetsV1(GetApplicationSetsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetApplicationSetsV1")
	}

	result := response.Result().(*ResponseApplicationPolicyGetApplicationSetsV1)
	return result, response, err

}

//GetApplicationSetsCountV1 Get Application Sets Count - cfa0-49a6-44bb-8a07
/* Get the number of existing application-sets



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-application-sets-count-v1
*/
func (s *ApplicationPolicyService) GetApplicationSetsCountV1() (*ResponseApplicationPolicyGetApplicationSetsCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/application-policy-application-set-count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplicationPolicyGetApplicationSetsCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApplicationSetsCountV1()
		}
		return nil, response, fmt.Errorf("error with operation GetApplicationSetsCountV1")
	}

	result := response.Result().(*ResponseApplicationPolicyGetApplicationSetsCountV1)
	return result, response, err

}

//GetApplicationsV1 Get Applications - 8893-b834-445b-b29c
/* Get applications by offset/limit or by name


@param GetApplicationsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-applications-v1
*/
func (s *ApplicationPolicyService) GetApplicationsV1(GetApplicationsV1QueryParams *GetApplicationsV1QueryParams) (*ResponseApplicationPolicyGetApplicationsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/applications"

	queryString, _ := query.Values(GetApplicationsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplicationPolicyGetApplicationsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApplicationsV1(GetApplicationsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetApplicationsV1")
	}

	result := response.Result().(*ResponseApplicationPolicyGetApplicationsV1)
	return result, response, err

}

//GetApplicationsCountV1 Get Applications Count - 039d-e8b1-47a9-8690
/* Get the number of all existing applications



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-applications-count-v1
*/
func (s *ApplicationPolicyService) GetApplicationsCountV1() (*ResponseApplicationPolicyGetApplicationsCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/applications-count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplicationPolicyGetApplicationsCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApplicationsCountV1()
		}
		return nil, response, fmt.Errorf("error with operation GetApplicationsCountV1")
	}

	result := response.Result().(*ResponseApplicationPolicyGetApplicationsCountV1)
	return result, response, err

}

//GetQosDeviceInterfaceInfoV1 Get Qos Device Interface Info - 42a6-e9eb-46bb-a197
/* Get all or by network device id, existing qos device interface infos


@param GetQosDeviceInterfaceInfoV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-qos-device-interface-info-v1
*/
func (s *ApplicationPolicyService) GetQosDeviceInterfaceInfoV1(GetQosDeviceInterfaceInfoV1QueryParams *GetQosDeviceInterfaceInfoV1QueryParams) (*ResponseApplicationPolicyGetQosDeviceInterfaceInfoV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/qos-device-interface-info"

	queryString, _ := query.Values(GetQosDeviceInterfaceInfoV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplicationPolicyGetQosDeviceInterfaceInfoV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetQosDeviceInterfaceInfoV1(GetQosDeviceInterfaceInfoV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetQosDeviceInterfaceInfoV1")
	}

	result := response.Result().(*ResponseApplicationPolicyGetQosDeviceInterfaceInfoV1)
	return result, response, err

}

//GetQosDeviceInterfaceInfoCountV1 Get Qos Device Interface Info Count - 729a-98e1-4a3b-91f2
/* Get the number of all existing qos device interface infos group by network device id



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-qos-device-interface-info-count-v1
*/
func (s *ApplicationPolicyService) GetQosDeviceInterfaceInfoCountV1() (*ResponseApplicationPolicyGetQosDeviceInterfaceInfoCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/qos-device-interface-info-count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplicationPolicyGetQosDeviceInterfaceInfoCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetQosDeviceInterfaceInfoCountV1()
		}
		return nil, response, fmt.Errorf("error with operation GetQosDeviceInterfaceInfoCountV1")
	}

	result := response.Result().(*ResponseApplicationPolicyGetQosDeviceInterfaceInfoCountV1)
	return result, response, err

}

//GetApplicationSetsV2 Get Application Set/s - 00ac-d849-43aa-bc75
/* Get application set/s by offset/limit or by name


@param GetApplicationSetsV2QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-application-sets-v2
*/
func (s *ApplicationPolicyService) GetApplicationSetsV2(GetApplicationSetsV2QueryParams *GetApplicationSetsV2QueryParams) (*ResponseApplicationPolicyGetApplicationSetsV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/application-policy-application-set"

	queryString, _ := query.Values(GetApplicationSetsV2QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplicationPolicyGetApplicationSetsV2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApplicationSetsV2(GetApplicationSetsV2QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetApplicationSetsV2")
	}

	result := response.Result().(*ResponseApplicationPolicyGetApplicationSetsV2)
	return result, response, err

}

//GetApplicationSetCountV2 Get Application Set Count - 798c-fa61-432b-a67e
/* Get the number of all existing application sets


@param GetApplicationSetCountV2QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-application-set-count-v2
*/
func (s *ApplicationPolicyService) GetApplicationSetCountV2(GetApplicationSetCountV2QueryParams *GetApplicationSetCountV2QueryParams) (*ResponseApplicationPolicyGetApplicationSetCountV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/application-policy-application-set-count"

	queryString, _ := query.Values(GetApplicationSetCountV2QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplicationPolicyGetApplicationSetCountV2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApplicationSetCountV2(GetApplicationSetCountV2QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetApplicationSetCountV2")
	}

	result := response.Result().(*ResponseApplicationPolicyGetApplicationSetCountV2)
	return result, response, err

}

//GetApplicationsV2 Get Application/s - a395-8970-43bb-bedd
/* Get application/s by offset/limit or by name


@param GetApplicationsV2QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-applications-v2
*/
func (s *ApplicationPolicyService) GetApplicationsV2(GetApplicationsV2QueryParams *GetApplicationsV2QueryParams) (*ResponseApplicationPolicyGetApplicationsV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/applications"

	queryString, _ := query.Values(GetApplicationsV2QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplicationPolicyGetApplicationsV2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApplicationsV2(GetApplicationsV2QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetApplicationsV2")
	}

	result := response.Result().(*ResponseApplicationPolicyGetApplicationsV2)
	return result, response, err

}

//GetApplicationCountV2 Get Application Count - cfa3-7a5c-4c08-b24e
/* Get the number of all existing applications


@param GetApplicationCountV2QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-application-count-v2
*/
func (s *ApplicationPolicyService) GetApplicationCountV2(GetApplicationCountV2QueryParams *GetApplicationCountV2QueryParams) (*ResponseApplicationPolicyGetApplicationCountV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/applications-count"

	queryString, _ := query.Values(GetApplicationCountV2QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplicationPolicyGetApplicationCountV2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApplicationCountV2(GetApplicationCountV2QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetApplicationCountV2")
	}

	result := response.Result().(*ResponseApplicationPolicyGetApplicationCountV2)
	return result, response, err

}

//ApplicationPolicyIntentV1 Application Policy Intent - aea4-bb7b-4329-bd06
/* Create/Update/Delete application policy



Documentation Link: https://developer.cisco.com/docs/dna-center/#!application-policy-intent-v1
*/
func (s *ApplicationPolicyService) ApplicationPolicyIntentV1(requestApplicationPolicyApplicationPolicyIntentV1 *RequestApplicationPolicyApplicationPolicyIntentV1) (*ResponseApplicationPolicyApplicationPolicyIntentV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/app-policy-intent"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplicationPolicyApplicationPolicyIntentV1).
		SetResult(&ResponseApplicationPolicyApplicationPolicyIntentV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ApplicationPolicyIntentV1(requestApplicationPolicyApplicationPolicyIntentV1)
		}

		return nil, response, fmt.Errorf("error with operation ApplicationPolicyIntentV1")
	}

	result := response.Result().(*ResponseApplicationPolicyApplicationPolicyIntentV1)
	return result, response, err

}

//CreateApplicationPolicyQueuingProfileV1 Create Application Policy Queuing Profile - 78b7-ba71-4959-bbd2
/* Create new custom application queuing profile



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-application-policy-queuing-profile-v1
*/
func (s *ApplicationPolicyService) CreateApplicationPolicyQueuingProfileV1(requestApplicationPolicyCreateApplicationPolicyQueuingProfileV1 *RequestApplicationPolicyCreateApplicationPolicyQueuingProfileV1) (*ResponseApplicationPolicyCreateApplicationPolicyQueuingProfileV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/app-policy-queuing-profile"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplicationPolicyCreateApplicationPolicyQueuingProfileV1).
		SetResult(&ResponseApplicationPolicyCreateApplicationPolicyQueuingProfileV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateApplicationPolicyQueuingProfileV1(requestApplicationPolicyCreateApplicationPolicyQueuingProfileV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateApplicationPolicyQueuingProfileV1")
	}

	result := response.Result().(*ResponseApplicationPolicyCreateApplicationPolicyQueuingProfileV1)
	return result, response, err

}

//CreateApplicationSetV1 Create Application Set - 3e94-cb1b-485b-8b0e
/* Create new custom application-set/s



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-application-set-v1
*/
func (s *ApplicationPolicyService) CreateApplicationSetV1(requestApplicationPolicyCreateApplicationSetV1 *RequestApplicationPolicyCreateApplicationSetV1) (*ResponseApplicationPolicyCreateApplicationSetV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/application-policy-application-set"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplicationPolicyCreateApplicationSetV1).
		SetResult(&ResponseApplicationPolicyCreateApplicationSetV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateApplicationSetV1(requestApplicationPolicyCreateApplicationSetV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateApplicationSetV1")
	}

	result := response.Result().(*ResponseApplicationPolicyCreateApplicationSetV1)
	return result, response, err

}

//CreateApplicationV1 Create Application - fb9b-f80f-491a-9851
/* Create new Custom application



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-application-v1
*/
func (s *ApplicationPolicyService) CreateApplicationV1(requestApplicationPolicyCreateApplicationV1 *RequestApplicationPolicyCreateApplicationV1) (*ResponseApplicationPolicyCreateApplicationV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/applications"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplicationPolicyCreateApplicationV1).
		SetResult(&ResponseApplicationPolicyCreateApplicationV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateApplicationV1(requestApplicationPolicyCreateApplicationV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateApplicationV1")
	}

	result := response.Result().(*ResponseApplicationPolicyCreateApplicationV1)
	return result, response, err

}

//CreateQosDeviceInterfaceInfoV1 Create Qos Device Interface Info - 3889-59af-4cf8-9fde
/* Create qos device interface infos associate with network device id to allow the user to mark specific interfaces as WAN, to associate WAN interfaces with specific SP Profile and to be able to define a shaper on WAN interfaces



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-qos-device-interface-info-v1
*/
func (s *ApplicationPolicyService) CreateQosDeviceInterfaceInfoV1(requestApplicationPolicyCreateQosDeviceInterfaceInfoV1 *RequestApplicationPolicyCreateQosDeviceInterfaceInfoV1) (*ResponseApplicationPolicyCreateQosDeviceInterfaceInfoV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/qos-device-interface-info"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplicationPolicyCreateQosDeviceInterfaceInfoV1).
		SetResult(&ResponseApplicationPolicyCreateQosDeviceInterfaceInfoV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateQosDeviceInterfaceInfoV1(requestApplicationPolicyCreateQosDeviceInterfaceInfoV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateQosDeviceInterfaceInfoV1")
	}

	result := response.Result().(*ResponseApplicationPolicyCreateQosDeviceInterfaceInfoV1)
	return result, response, err

}

//CreateApplicationSetsV2 Create Application Set/s - e4bf-ca74-45f9-a374
/* Create new custom application set/s



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-application-sets-v2
*/
func (s *ApplicationPolicyService) CreateApplicationSetsV2(requestApplicationPolicyCreateApplicationSetsV2 *RequestApplicationPolicyCreateApplicationSetsV2) (*ResponseApplicationPolicyCreateApplicationSetsV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/application-policy-application-set"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplicationPolicyCreateApplicationSetsV2).
		SetResult(&ResponseApplicationPolicyCreateApplicationSetsV2{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateApplicationSetsV2(requestApplicationPolicyCreateApplicationSetsV2)
		}

		return nil, response, fmt.Errorf("error with operation CreateApplicationSetsV2")
	}

	result := response.Result().(*ResponseApplicationPolicyCreateApplicationSetsV2)
	return result, response, err

}

//CreateApplicationsV2 Create Application/s - b4a6-dae7-4b29-992c
/* Create new custom application/s



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-applications-v2
*/
func (s *ApplicationPolicyService) CreateApplicationsV2(requestApplicationPolicyCreateApplicationsV2 *RequestApplicationPolicyCreateApplicationsV2) (*ResponseApplicationPolicyCreateApplicationsV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/applications"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplicationPolicyCreateApplicationsV2).
		SetResult(&ResponseApplicationPolicyCreateApplicationsV2{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateApplicationsV2(requestApplicationPolicyCreateApplicationsV2)
		}

		return nil, response, fmt.Errorf("error with operation CreateApplicationsV2")
	}

	result := response.Result().(*ResponseApplicationPolicyCreateApplicationsV2)
	return result, response, err

}

//UpdateApplicationPolicyQueuingProfileV1 Update Application Policy Queuing Profile - da98-38ea-44aa-8024
/* Update existing custom application queuing profile


 */
func (s *ApplicationPolicyService) UpdateApplicationPolicyQueuingProfileV1(requestApplicationPolicyUpdateApplicationPolicyQueuingProfileV1 *RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileV1) (*ResponseApplicationPolicyUpdateApplicationPolicyQueuingProfileV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/app-policy-queuing-profile"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplicationPolicyUpdateApplicationPolicyQueuingProfileV1).
		SetResult(&ResponseApplicationPolicyUpdateApplicationPolicyQueuingProfileV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateApplicationPolicyQueuingProfileV1(requestApplicationPolicyUpdateApplicationPolicyQueuingProfileV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateApplicationPolicyQueuingProfileV1")
	}

	result := response.Result().(*ResponseApplicationPolicyUpdateApplicationPolicyQueuingProfileV1)
	return result, response, err

}

//EditApplicationV1 Edit Application - 3986-6887-4439-a41d
/* Edit the attributes of an existing application


 */
func (s *ApplicationPolicyService) EditApplicationV1(requestApplicationPolicyEditApplicationV1 *RequestApplicationPolicyEditApplicationV1) (*ResponseApplicationPolicyEditApplicationV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/applications"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplicationPolicyEditApplicationV1).
		SetResult(&ResponseApplicationPolicyEditApplicationV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.EditApplicationV1(requestApplicationPolicyEditApplicationV1)
		}
		return nil, response, fmt.Errorf("error with operation EditApplicationV1")
	}

	result := response.Result().(*ResponseApplicationPolicyEditApplicationV1)
	return result, response, err

}

//UpdateQosDeviceInterfaceInfoV1 Update Qos Device Interface Info - 71b7-ba8c-47b8-95b6
/* Update existing qos device interface infos associate with network device id


 */
func (s *ApplicationPolicyService) UpdateQosDeviceInterfaceInfoV1(requestApplicationPolicyUpdateQosDeviceInterfaceInfoV1 *RequestApplicationPolicyUpdateQosDeviceInterfaceInfoV1) (*ResponseApplicationPolicyUpdateQosDeviceInterfaceInfoV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/qos-device-interface-info"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplicationPolicyUpdateQosDeviceInterfaceInfoV1).
		SetResult(&ResponseApplicationPolicyUpdateQosDeviceInterfaceInfoV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateQosDeviceInterfaceInfoV1(requestApplicationPolicyUpdateQosDeviceInterfaceInfoV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateQosDeviceInterfaceInfoV1")
	}

	result := response.Result().(*ResponseApplicationPolicyUpdateQosDeviceInterfaceInfoV1)
	return result, response, err

}

//EditApplicationsV2 Edit Application/s - 6995-2aea-4f2b-a053
/* Edit the attributes of an existing application


 */
func (s *ApplicationPolicyService) EditApplicationsV2(requestApplicationPolicyEditApplicationsV2 *RequestApplicationPolicyEditApplicationsV2) (*ResponseApplicationPolicyEditApplicationsV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/applications"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplicationPolicyEditApplicationsV2).
		SetResult(&ResponseApplicationPolicyEditApplicationsV2{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.EditApplicationsV2(requestApplicationPolicyEditApplicationsV2)
		}
		return nil, response, fmt.Errorf("error with operation EditApplicationsV2")
	}

	result := response.Result().(*ResponseApplicationPolicyEditApplicationsV2)
	return result, response, err

}

//DeleteApplicationPolicyQueuingProfileV1 Delete Application Policy Queuing Profile - 09a0-482f-422b-b325
/* Delete existing custom application policy queuing profile by id


@param id id path parameter. Id of custom queuing profile to delete


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-application-policy-queuing-profile-v1
*/
func (s *ApplicationPolicyService) DeleteApplicationPolicyQueuingProfileV1(id string) (*ResponseApplicationPolicyDeleteApplicationPolicyQueuingProfileV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/app-policy-queuing-profile/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplicationPolicyDeleteApplicationPolicyQueuingProfileV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteApplicationPolicyQueuingProfileV1(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteApplicationPolicyQueuingProfileV1")
	}

	result := response.Result().(*ResponseApplicationPolicyDeleteApplicationPolicyQueuingProfileV1)
	return result, response, err

}

//DeleteApplicationSetV1 Delete Application Set - 70b6-f8e1-40b8-b784
/* Delete existing application-set by it's id


@param DeleteApplicationSetV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-application-set-v1
*/
func (s *ApplicationPolicyService) DeleteApplicationSetV1(DeleteApplicationSetV1QueryParams *DeleteApplicationSetV1QueryParams) (*ResponseApplicationPolicyDeleteApplicationSetV1, *resty.Response, error) {
	//DeleteApplicationSetV1QueryParams *DeleteApplicationSetV1QueryParams
	path := "/dna/intent/api/v1/application-policy-application-set"

	queryString, _ := query.Values(DeleteApplicationSetV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplicationPolicyDeleteApplicationSetV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteApplicationSetV1(DeleteApplicationSetV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteApplicationSetV1")
	}

	result := response.Result().(*ResponseApplicationPolicyDeleteApplicationSetV1)
	return result, response, err

}

//DeleteApplicationV1 Delete Application - d49a-f9b8-4c6a-a8ea
/* Delete existing application by its id


@param DeleteApplicationV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-application-v1
*/
func (s *ApplicationPolicyService) DeleteApplicationV1(DeleteApplicationV1QueryParams *DeleteApplicationV1QueryParams) (*ResponseApplicationPolicyDeleteApplicationV1, *resty.Response, error) {
	//DeleteApplicationV1QueryParams *DeleteApplicationV1QueryParams
	path := "/dna/intent/api/v1/applications"

	queryString, _ := query.Values(DeleteApplicationV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplicationPolicyDeleteApplicationV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteApplicationV1(DeleteApplicationV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteApplicationV1")
	}

	result := response.Result().(*ResponseApplicationPolicyDeleteApplicationV1)
	return result, response, err

}

//DeleteQosDeviceInterfaceInfoV1 Delete Qos Device Interface Info - 51b1-b8bd-435b-9842
/* Delete all qos device interface infos associate with network device id


@param id id path parameter. Id of the qos device info, this object holds all qos device interface infos associate with network device id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-qos-device-interface-info-v1
*/
func (s *ApplicationPolicyService) DeleteQosDeviceInterfaceInfoV1(id string) (*ResponseApplicationPolicyDeleteQosDeviceInterfaceInfoV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/qos-device-interface-info/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplicationPolicyDeleteQosDeviceInterfaceInfoV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteQosDeviceInterfaceInfoV1(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteQosDeviceInterfaceInfoV1")
	}

	result := response.Result().(*ResponseApplicationPolicyDeleteQosDeviceInterfaceInfoV1)
	return result, response, err

}

//DeleteApplicationSetV2 Delete Application Set - b2a9-08f1-4879-9f70
/* Delete existing custom application set by id


@param id id path parameter. Id of custom application set to delete


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-application-set-v2
*/
func (s *ApplicationPolicyService) DeleteApplicationSetV2(id string) (*ResponseApplicationPolicyDeleteApplicationSetV2, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v2/application-policy-application-set/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplicationPolicyDeleteApplicationSetV2{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteApplicationSetV2(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteApplicationSetV2")
	}

	result := response.Result().(*ResponseApplicationPolicyDeleteApplicationSetV2)
	return result, response, err

}

//DeleteApplicationV2 Delete Application - 9098-8ada-4abb-bedc
/* Delete existing custom application by id


@param id id path parameter. Id of custom application to delete


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-application-v2
*/
func (s *ApplicationPolicyService) DeleteApplicationV2(id string) (*ResponseApplicationPolicyDeleteApplicationV2, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v2/applications/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplicationPolicyDeleteApplicationV2{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteApplicationV2(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteApplicationV2")
	}

	result := response.Result().(*ResponseApplicationPolicyDeleteApplicationV2)
	return result, response, err

}

// Alias Function
func (s *ApplicationPolicyService) GetApplicationPolicyDefault() (*ResponseApplicationPolicyGetApplicationPolicyDefaultV1, *resty.Response, error) {
	return s.GetApplicationPolicyDefaultV1()
}

// Alias Function
func (s *ApplicationPolicyService) GetQosDeviceInterfaceInfo(GetQosDeviceInterfaceInfoV1QueryParams *GetQosDeviceInterfaceInfoV1QueryParams) (*ResponseApplicationPolicyGetQosDeviceInterfaceInfoV1, *resty.Response, error) {
	return s.GetQosDeviceInterfaceInfoV1(GetQosDeviceInterfaceInfoV1QueryParams)
}

// Alias Function
func (s *ApplicationPolicyService) CreateApplication(requestApplicationPolicyCreateApplicationV1 *RequestApplicationPolicyCreateApplicationV1) (*ResponseApplicationPolicyCreateApplicationV1, *resty.Response, error) {
	return s.CreateApplicationV1(requestApplicationPolicyCreateApplicationV1)
}

// Alias Function
func (s *ApplicationPolicyService) DeleteApplication(DeleteApplicationV1QueryParams *DeleteApplicationV1QueryParams) (*ResponseApplicationPolicyDeleteApplicationV1, *resty.Response, error) {
	return s.DeleteApplicationV1(DeleteApplicationV1QueryParams)
}

// Alias Function
func (s *ApplicationPolicyService) GetApplicationPolicy(GetApplicationPolicyV1QueryParams *GetApplicationPolicyV1QueryParams) (*ResponseApplicationPolicyGetApplicationPolicyV1, *resty.Response, error) {
	return s.GetApplicationPolicyV1(GetApplicationPolicyV1QueryParams)
}

// Alias Function
func (s *ApplicationPolicyService) GetApplicationCount(GetApplicationCountV2QueryParams *GetApplicationCountV2QueryParams) (*ResponseApplicationPolicyGetApplicationCountV2, *resty.Response, error) {
	return s.GetApplicationCountV2(GetApplicationCountV2QueryParams)
}

// Alias Function
func (s *ApplicationPolicyService) GetApplicationSetsCount() (*ResponseApplicationPolicyGetApplicationSetsCountV1, *resty.Response, error) {
	return s.GetApplicationSetsCountV1()
}

// Alias Function
func (s *ApplicationPolicyService) EditApplication(requestApplicationPolicyEditApplicationV1 *RequestApplicationPolicyEditApplicationV1) (*ResponseApplicationPolicyEditApplicationV1, *resty.Response, error) {
	return s.EditApplicationV1(requestApplicationPolicyEditApplicationV1)
}

// Alias Function
func (s *ApplicationPolicyService) GetApplicationPolicyQueuingProfileCount() (*ResponseApplicationPolicyGetApplicationPolicyQueuingProfileCountV1, *resty.Response, error) {
	return s.GetApplicationPolicyQueuingProfileCountV1()
}

// Alias Function
func (s *ApplicationPolicyService) CreateApplicationSets(requestApplicationPolicyCreateApplicationSetsV2 *RequestApplicationPolicyCreateApplicationSetsV2) (*ResponseApplicationPolicyCreateApplicationSetsV2, *resty.Response, error) {
	return s.CreateApplicationSetsV2(requestApplicationPolicyCreateApplicationSetsV2)
}

// Alias Function
func (s *ApplicationPolicyService) EditApplications(requestApplicationPolicyEditApplicationsV2 *RequestApplicationPolicyEditApplicationsV2) (*ResponseApplicationPolicyEditApplicationsV2, *resty.Response, error) {
	return s.EditApplicationsV2(requestApplicationPolicyEditApplicationsV2)
}

// Alias Function
func (s *ApplicationPolicyService) GetApplicationSetCount(GetApplicationSetCountV2QueryParams *GetApplicationSetCountV2QueryParams) (*ResponseApplicationPolicyGetApplicationSetCountV2, *resty.Response, error) {
	return s.GetApplicationSetCountV2(GetApplicationSetCountV2QueryParams)
}

// Alias Function
func (s *ApplicationPolicyService) CreateApplications(requestApplicationPolicyCreateApplicationsV2 *RequestApplicationPolicyCreateApplicationsV2) (*ResponseApplicationPolicyCreateApplicationsV2, *resty.Response, error) {
	return s.CreateApplicationsV2(requestApplicationPolicyCreateApplicationsV2)
}

// Alias Function
func (s *ApplicationPolicyService) GetApplications(GetApplicationsV1QueryParams *GetApplicationsV1QueryParams) (*ResponseApplicationPolicyGetApplicationsV1, *resty.Response, error) {
	return s.GetApplicationsV1(GetApplicationsV1QueryParams)
}

// Alias Function
func (s *ApplicationPolicyService) DeleteQosDeviceInterfaceInfo(id string) (*ResponseApplicationPolicyDeleteQosDeviceInterfaceInfoV1, *resty.Response, error) {
	return s.DeleteQosDeviceInterfaceInfoV1(id)
}

// Alias Function
func (s *ApplicationPolicyService) GetQosDeviceInterfaceInfoCount() (*ResponseApplicationPolicyGetQosDeviceInterfaceInfoCountV1, *resty.Response, error) {
	return s.GetQosDeviceInterfaceInfoCountV1()
}

// Alias Function
func (s *ApplicationPolicyService) ApplicationPolicyIntent(requestApplicationPolicyApplicationPolicyIntentV1 *RequestApplicationPolicyApplicationPolicyIntentV1) (*ResponseApplicationPolicyApplicationPolicyIntentV1, *resty.Response, error) {
	return s.ApplicationPolicyIntentV1(requestApplicationPolicyApplicationPolicyIntentV1)
}

// Alias Function
func (s *ApplicationPolicyService) GetApplicationPolicyQueuingProfile(GetApplicationPolicyQueuingProfileV1QueryParams *GetApplicationPolicyQueuingProfileV1QueryParams) (*ResponseApplicationPolicyGetApplicationPolicyQueuingProfileV1, *resty.Response, error) {
	return s.GetApplicationPolicyQueuingProfileV1(GetApplicationPolicyQueuingProfileV1QueryParams)
}

// Alias Function
func (s *ApplicationPolicyService) UpdateApplicationPolicyQueuingProfile(requestApplicationPolicyUpdateApplicationPolicyQueuingProfileV1 *RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileV1) (*ResponseApplicationPolicyUpdateApplicationPolicyQueuingProfileV1, *resty.Response, error) {
	return s.UpdateApplicationPolicyQueuingProfileV1(requestApplicationPolicyUpdateApplicationPolicyQueuingProfileV1)
}

// Alias Function
func (s *ApplicationPolicyService) DeleteApplicationSet(DeleteApplicationSetV1QueryParams *DeleteApplicationSetV1QueryParams) (*ResponseApplicationPolicyDeleteApplicationSetV1, *resty.Response, error) {
	return s.DeleteApplicationSetV1(DeleteApplicationSetV1QueryParams)
}

// Alias Function
func (s *ApplicationPolicyService) GetApplicationSets(GetApplicationSetsV1QueryParams *GetApplicationSetsV1QueryParams) (*ResponseApplicationPolicyGetApplicationSetsV1, *resty.Response, error) {
	return s.GetApplicationSetsV1(GetApplicationSetsV1QueryParams)
}

// Alias Function
func (s *ApplicationPolicyService) UpdateQosDeviceInterfaceInfo(requestApplicationPolicyUpdateQosDeviceInterfaceInfoV1 *RequestApplicationPolicyUpdateQosDeviceInterfaceInfoV1) (*ResponseApplicationPolicyUpdateQosDeviceInterfaceInfoV1, *resty.Response, error) {
	return s.UpdateQosDeviceInterfaceInfoV1(requestApplicationPolicyUpdateQosDeviceInterfaceInfoV1)
}

// Alias Function
func (s *ApplicationPolicyService) GetApplicationsCount() (*ResponseApplicationPolicyGetApplicationsCountV1, *resty.Response, error) {
	return s.GetApplicationsCountV1()
}

// Alias Function
func (s *ApplicationPolicyService) CreateQosDeviceInterfaceInfo(requestApplicationPolicyCreateQosDeviceInterfaceInfoV1 *RequestApplicationPolicyCreateQosDeviceInterfaceInfoV1) (*ResponseApplicationPolicyCreateQosDeviceInterfaceInfoV1, *resty.Response, error) {
	return s.CreateQosDeviceInterfaceInfoV1(requestApplicationPolicyCreateQosDeviceInterfaceInfoV1)
}

// Alias Function
func (s *ApplicationPolicyService) CreateApplicationPolicyQueuingProfile(requestApplicationPolicyCreateApplicationPolicyQueuingProfileV1 *RequestApplicationPolicyCreateApplicationPolicyQueuingProfileV1) (*ResponseApplicationPolicyCreateApplicationPolicyQueuingProfileV1, *resty.Response, error) {
	return s.CreateApplicationPolicyQueuingProfileV1(requestApplicationPolicyCreateApplicationPolicyQueuingProfileV1)
}

// Alias Function
func (s *ApplicationPolicyService) DeleteApplicationPolicyQueuingProfile(id string) (*ResponseApplicationPolicyDeleteApplicationPolicyQueuingProfileV1, *resty.Response, error) {
	return s.DeleteApplicationPolicyQueuingProfileV1(id)
}

// Alias Function
func (s *ApplicationPolicyService) CreateApplicationSet(requestApplicationPolicyCreateApplicationSetV1 *RequestApplicationPolicyCreateApplicationSetV1) (*ResponseApplicationPolicyCreateApplicationSetV1, *resty.Response, error) {
	return s.CreateApplicationSetV1(requestApplicationPolicyCreateApplicationSetV1)
}
