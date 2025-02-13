package catalyst

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type LicensesService service

type DeviceCountDetailsV1QueryParams struct {
	DeviceType         string `url:"device_type,omitempty"`          //Type of device
	RegistrationStatus string `url:"registration_status,omitempty"`  //Smart license registration status of device
	DnaLevel           string `url:"dna_level,omitempty"`            //Device Cisco DNA License Level
	VirtualAccountName string `url:"virtual_account_name,omitempty"` //Virtual account name
	SmartAccountID     string `url:"smart_account_id,omitempty"`     //Smart account id
}
type DeviceLicenseSummaryV1QueryParams struct {
	PageNumber         float64 `url:"page_number,omitempty"`          //Page number of response
	Order              string  `url:"order,omitempty"`                //Sorting order
	SortBy             string  `url:"sort_by,omitempty"`              //Sort result by field
	DnaLevel           string  `url:"dna_level,omitempty"`            //Device Cisco DNA license level
	DeviceType         string  `url:"device_type,omitempty"`          //Type of device
	Limit              float64 `url:"limit,omitempty"`                //Specifies the maximum number of device license summaries to return per page. Must be an integer between 1 and 500, inclusive.
	RegistrationStatus string  `url:"registration_status,omitempty"`  //Smart license registration status of device
	VirtualAccountName string  `url:"virtual_account_name,omitempty"` //Name of virtual account
	SmartAccountID     string  `url:"smart_account_id,omitempty"`     //Id of smart account
	DeviceUUID         string  `url:"device_uuid,omitempty"`          //Id of device
}
type LicenseTermDetailsV1QueryParams struct {
	DeviceType string `url:"device_type,omitempty"` //Type of device like router, switch, wireless or ise
}
type LicenseUsageDetailsV1QueryParams struct {
	DeviceType string `url:"device_type,omitempty"` //Type of device like router, switch, wireless or ise
}

type ResponseLicensesRetrievesCSSMConnectionModeV1 struct {
	Response *ResponseLicensesRetrievesCSSMConnectionModeV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // API version
}
type ResponseLicensesRetrievesCSSMConnectionModeV1Response struct {
	ConnectionMode string `json:"connectionMode,omitempty"` // The CSSM connection modes of Catalyst Center are DIRECT, ON_PREMISE and SMART_PROXY

	Parameters *ResponseLicensesRetrievesCSSMConnectionModeV1ResponseParameters `json:"parameters,omitempty"` //
}
type ResponseLicensesRetrievesCSSMConnectionModeV1ResponseParameters struct {
	OnPremiseHost string `json:"onPremiseHost,omitempty"` // On-premise host

	SmartAccountName string `json:"smartAccountName,omitempty"` // On-premise CSSM local smart account name

	ClientID string `json:"clientId,omitempty"` // On-premise CSSM client id
}
type ResponseLicensesUpdateCSSMConnectionModeV1 struct {
	Version string `json:"version,omitempty"` // Response Version e.g. : 1.0

	Response *ResponseLicensesUpdateCSSMConnectionModeV1Response `json:"response,omitempty"` //
}
type ResponseLicensesUpdateCSSMConnectionModeV1Response struct {
	URL string `json:"url,omitempty"` // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5

	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseLicensesRetrieveLicenseSettingV1 struct {
	Response *ResponseLicensesRetrieveLicenseSettingV1Response `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  // API version
}
type ResponseLicensesRetrieveLicenseSettingV1Response struct {
	DefaultSmartAccountID            string `json:"defaultSmartAccountId,omitempty"`            // Default smart account id
	AutoRegistrationVirtualAccountID string `json:"autoRegistrationVirtualAccountId,omitempty"` // Virtual account id
}
type ResponseLicensesUpdateLicenseSettingV1 struct {
	Response *ResponseLicensesUpdateLicenseSettingV1Response `json:"response,omitempty"` //
	Version  string                                          `json:"version,omitempty"`  // API version
}
type ResponseLicensesUpdateLicenseSettingV1Response struct {
	DefaultSmartAccountID            string `json:"defaultSmartAccountId,omitempty"`            // Default smart account id
	AutoRegistrationVirtualAccountID string `json:"autoRegistrationVirtualAccountId,omitempty"` // Virtual account id
}
type ResponseLicensesDeviceCountDetailsV1 struct {
	Response *int   `json:"response,omitempty"` // Total number of managed device
	Version  string `json:"version,omitempty"`  // Version
}
type ResponseLicensesDeviceLicenseSummaryV1 struct {
	Response *[]ResponseLicensesDeviceLicenseSummaryV1Response `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  // Version
}
type ResponseLicensesDeviceLicenseSummaryV1Response struct {
	AuthorizationStatus              string `json:"authorization_status,omitempty"`                  // Smart license authorization status of device
	LastUpdatedTime                  string `json:"last_updated_time,omitempty"`                     // Time when license information was collected from device
	IsPerformanceAllowed             *bool  `json:"is_performance_allowed,omitempty"`                // Is performance license available
	SleAuthCode                      string `json:"sle_auth_code,omitempty"`                         // SLE Authcode installed or not installed
	ThroughputLevel                  string `json:"throughput_level,omitempty"`                      // Throughput level on device
	HsecStatus                       string `json:"hsec_status,omitempty"`                           // HSEC status of the device
	DeviceUUID                       string `json:"device_uuid,omitempty"`                           // Id of device
	Site                             string `json:"site,omitempty"`                                  // Site of device
	TotalAccessPointCount            *int   `json:"total_access_point_count,omitempty"`              // Total number of Access Points connected
	Model                            string `json:"model,omitempty"`                                 // Model of device
	IsWirelessCapable                *bool  `json:"is_wireless_capable,omitempty"`                   // Is device wireless capable
	RegistrationStatus               string `json:"registration_status,omitempty"`                   // Smart license registration status of device
	SleState                         string `json:"sle_state,omitempty"`                             // SLE state on device
	PerformanceLicense               string `json:"performance_license,omitempty"`                   // Is performance license enabled
	LicenseMode                      string `json:"license_mode,omitempty"`                          // Mode of license
	IsLicenseExpired                 *bool  `json:"is_license_expired,omitempty"`                    // Is device license expired
	SoftwareVersion                  string `json:"software_version,omitempty"`                      // Software image version of device
	ReservationStatus                string `json:"reservation_status,omitempty"`                    // Smart license reservation status
	IsWireless                       *bool  `json:"is_wireless,omitempty"`                           // Is device wireless controller
	NetworkLicense                   string `json:"network_license,omitempty"`                       // Device Network license level
	EvaluationLicenseExpiry          string `json:"evaluation_license_expiry,omitempty"`             // Evaluation period expiry date
	WirelessCapableNetworkLicense    string `json:"wireless_capable_network_license,omitempty"`      // Wireless Cisco Network license value
	DeviceName                       string `json:"device_name,omitempty"`                           // Name of device
	DeviceType                       string `json:"device_type,omitempty"`                           // Type of device
	DnaLevel                         string `json:"dna_level,omitempty"`                             // Device Cisco DNA license level
	VirtualAccountName               string `json:"virtual_account_name,omitempty"`                  // Name of virtual account
	LastSuccessfulRumUsageUploadTime string `json:"last_successful_rum_usage_upload_time,omitempty"` // Last successful rum usage upload time
	IPAddress                        string `json:"ip_address,omitempty"`                            // IP address of device
	WirelessCapableDnaLicense        string `json:"wireless_capable_dna_license,omitempty"`          // Wireless Cisco DNA license value
	MacAddress                       string `json:"mac_address,omitempty"`                           // MAC address of device
	CustomerTag1                     string `json:"customer_tag1,omitempty"`                         // Customer Tag1 set on device
	CustomerTag2                     string `json:"customer_tag2,omitempty"`                         // Customer Tag2 set on device
	CustomerTag3                     string `json:"customer_tag3,omitempty"`                         // Customer Tag3 set on device
	CustomerTag4                     string `json:"customer_tag4,omitempty"`                         // Customer Tag4 set on device
	SmartAccountName                 string `json:"smart_account_name,omitempty"`                    // Name of smart account
}
type ResponseLicensesDeviceLicenseDetailsV1 struct {
	DeviceUUID              string                                                  `json:"device_uuid,omitempty"`               // Id of device
	Site                    string                                                  `json:"site,omitempty"`                      // Site of device
	Model                   string                                                  `json:"model,omitempty"`                     // Model of device
	LicenseMode             string                                                  `json:"license_mode,omitempty"`              // Mode of license
	IsLicenseExpired        *bool                                                   `json:"is_license_expired,omitempty"`        // Is device license expired
	SoftwareVersion         string                                                  `json:"software_version,omitempty"`          // Software image version of device
	NetworkLicense          string                                                  `json:"network_license,omitempty"`           // Device network license level
	EvaluationLicenseExpiry string                                                  `json:"evaluation_license_expiry,omitempty"` // Evaluation period expiry date
	DeviceName              string                                                  `json:"device_name,omitempty"`               // Name of device
	DeviceType              string                                                  `json:"device_type,omitempty"`               // Type of device
	DnaLevel                string                                                  `json:"dna_level,omitempty"`                 // Device Cisco DNA license level
	VirtualAccountName      string                                                  `json:"virtual_account_name,omitempty"`      // Name of virtual account
	IPAddress               string                                                  `json:"ip_address,omitempty"`                // IP address of device
	MacAddress              string                                                  `json:"mac_address,omitempty"`               // MAC address of device
	SntcStatus              string                                                  `json:"sntc_status,omitempty"`               // Valid if device is covered under license contract and invalid if not covered, otherwise unknown.
	FeatureLicense          []string                                                `json:"feature_license,omitempty"`           // Name of feature licenses
	HasSupCards             *bool                                                   `json:"has_sup_cards,omitempty"`             // Whether device has supervisor cards
	Udi                     string                                                  `json:"udi,omitempty"`                       // Unique Device Identifier
	StackedDevices          *[]ResponseLicensesDeviceLicenseDetailsV1StackedDevices `json:"stacked_devices,omitempty"`           //
	IsStackedDevice         *bool                                                   `json:"is_stacked_device,omitempty"`         // Is Stacked Device
	AccessPoints            *[]ResponseLicensesDeviceLicenseDetailsV1AccessPoints   `json:"access_points,omitempty"`             //
	ChassisDetails          *ResponseLicensesDeviceLicenseDetailsV1ChassisDetails   `json:"chassis_details,omitempty"`           //
}
type ResponseLicensesDeviceLicenseDetailsV1StackedDevices struct {
	MacAddress   string `json:"mac_address,omitempty"`   // Stack mac address
	ID           *int   `json:"id,omitempty"`            // Id
	Role         string `json:"role,omitempty"`          // Chassis role
	SerialNumber string `json:"serial_number,omitempty"` // Chassis serial number
}
type ResponseLicensesDeviceLicenseDetailsV1AccessPoints struct {
	ApType string `json:"ap_type,omitempty"` // Type of access point
	Count  string `json:"count,omitempty"`   // Number of access point
}
type ResponseLicensesDeviceLicenseDetailsV1ChassisDetails struct {
	BoardSerialNumber string                                                                 `json:"board_serial_number,omitempty"` // Board serial number
	Modules           *[]ResponseLicensesDeviceLicenseDetailsV1ChassisDetailsModules         `json:"modules,omitempty"`             //
	SupervisorCards   *[]ResponseLicensesDeviceLicenseDetailsV1ChassisDetailsSupervisorCards `json:"supervisor_cards,omitempty"`    //
	Port              *int                                                                   `json:"port,omitempty"`                // Number of port
}
type ResponseLicensesDeviceLicenseDetailsV1ChassisDetailsModules struct {
	ModuleType   string `json:"module_type,omitempty"`   // Type of module
	ModuleName   string `json:"module_name,omitempty"`   // Name of module
	SerialNumber string `json:"serial_number,omitempty"` // Serial number of module
	ID           *int   `json:"id,omitempty"`            // Id of module
}
type ResponseLicensesDeviceLicenseDetailsV1ChassisDetailsSupervisorCards struct {
	SerialNumber       string `json:"serial_number,omitempty"`        // Serial number of supervisor card
	SupervisorCardType string `json:"supervisor_card_type,omitempty"` // Type of supervisor card
	Status             string `json:"status,omitempty"`               // Status of supervisor card
}
type ResponseLicensesDeviceDeregistrationV1 struct {
	Response *ResponseLicensesDeviceDeregistrationV1Response `json:"response,omitempty"` //
}
type ResponseLicensesDeviceDeregistrationV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task id of process
	URL    string `json:"url,omitempty"`    // Task URL of process
}
type ResponseLicensesDeviceRegistrationV1 struct {
	Response *ResponseLicensesDeviceRegistrationV1Response `json:"response,omitempty"` //
}
type ResponseLicensesDeviceRegistrationV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task id of process
	URL    string `json:"url,omitempty"`    // Task URL of process
}
type ResponseLicensesChangeVirtualAccountV1 struct {
	Response *ResponseLicensesChangeVirtualAccountV1Response `json:"response,omitempty"` //
}
type ResponseLicensesChangeVirtualAccountV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task id of process
	URL    string `json:"url,omitempty"`    // Task URL of process
}
type ResponseLicensesVirtualAccountDetailsV1 struct {
	SmartAccountID        string                                                          `json:"smart_account_id,omitempty"`        // Id of smart account
	SmartAccountName      string                                                          `json:"smart_account_name,omitempty"`      // Name of smart account
	VirtualAccountDetails *[]ResponseLicensesVirtualAccountDetailsV1VirtualAccountDetails `json:"virtual_account_details,omitempty"` //
}
type ResponseLicensesVirtualAccountDetailsV1VirtualAccountDetails struct {
	VirtualAccountID   string `json:"virtual_account_id,omitempty"`   // Id of virtual account
	VirtualAccountName string `json:"virtual_account_name,omitempty"` // Name of virtual account
}
type ResponseLicensesSmartAccountDetailsV1 struct {
	Response *[]ResponseLicensesSmartAccountDetailsV1Response `json:"response,omitempty"` //
	Version  string                                           `json:"version,omitempty"`  // Version
}
type ResponseLicensesSmartAccountDetailsV1Response struct {
	Name                 string `json:"name,omitempty"`                    // Name of smart account
	ID                   string `json:"id,omitempty"`                      // Id of smart account
	Domain               string `json:"domain,omitempty"`                  // Domain of smart account
	IsActiveSmartAccount *bool  `json:"is_active_smart_account,omitempty"` // Is active smart account
}
type ResponseLicensesLicenseTermDetailsV1 struct {
	LicenseDetails *[]ResponseLicensesLicenseTermDetailsV1LicenseDetails `json:"license_details,omitempty"` //
}
type ResponseLicensesLicenseTermDetailsV1LicenseDetails struct {
	Model                    string `json:"model,omitempty"`                       // Model of device
	VirtualAccountName       string `json:"virtual_account_name,omitempty"`        // Name of virtual account
	LicenseTermStartDate     string `json:"license_term_start_date,omitempty"`     // Start date of license term
	LicenseTermEndDate       string `json:"license_term_end_date,omitempty"`       // End date of license term
	DnaLevel                 string `json:"dna_level,omitempty"`                   // Cisco DNA license level
	PurchasedDnaLicenseCount string `json:"purchased_dna_license_count,omitempty"` // Number of purchased DNA licenses
	IsLicenseExpired         string `json:"is_license_expired,omitempty"`          // Is license expired
}
type ResponseLicensesLicenseUsageDetailsV1 struct {
	PurchasedDnaLicense *ResponseLicensesLicenseUsageDetailsV1PurchasedDnaLicense `json:"purchased_dna_license,omitempty"` //

	PurchasedNetworkLicense *ResponseLicensesLicenseUsageDetailsV1PurchasedNetworkLicense `json:"purchased_network_license,omitempty"` //

	UsedDnaLicense *ResponseLicensesLicenseUsageDetailsV1UsedDnaLicense `json:"used_dna_license,omitempty"` //

	UsedNetworkLicense *ResponseLicensesLicenseUsageDetailsV1UsedNetworkLicense `json:"used_network_license,omitempty"` //

	PurchasedIseLicense *ResponseLicensesLicenseUsageDetailsPurchasedIseLicense `json:"purchased_ise_license,omitempty"`

	UsedIseLicense *ResponseLicensesLicenseUsageDetailsUsedIseLicense `json:"used_ise_license,omitempty"`
}
type ResponseLicensesLicenseUsageDetailsUsedIseLicense struct {
	TotalLicenseCount  *int                                                                   `json:"total_license_count,omitempty"`
	LicenseCountByType *[]ResponseLicensesLicenseUsageDetailsUsedIseLicenseLicenseCountByType `json:"license_count_by_type,omitempty"`
}
type ResponseLicensesLicenseUsageDetailsUsedIseLicenseLicenseCountByType struct {
	LicenseType  string `json:"license_type,omitempty"`
	LicenseCount *int   `json:"license_count,omitempty"`
}
type ResponseLicensesLicenseUsageDetailsPurchasedIseLicense struct {
	TotalLicenseCount  *int                                                                        `json:"total_license_count,omitempty"`
	LicenseCountByType *[]ResponseLicensesLicenseUsageDetailsPurchasedIseLicenseLicenseCountByType `json:"license_count_by_type,omitempty"`
}
type ResponseLicensesLicenseUsageDetailsPurchasedIseLicenseLicenseCountByType struct {
	LicenseType  string `json:"license_type,omitempty"`
	LicenseCount *int   `json:"license_count,omitempty"`
}
type ResponseLicensesLicenseUsageDetailsV1PurchasedDnaLicense struct {
	TotalLicenseCount  *int                                                                          `json:"total_license_count,omitempty"`   // Total number of licenses
	LicenseCountByType *[]ResponseLicensesLicenseUsageDetailsV1PurchasedDnaLicenseLicenseCountByType `json:"license_count_by_type,omitempty"` //
}
type ResponseLicensesLicenseUsageDetailsV1PurchasedDnaLicenseLicenseCountByType struct {
	LicenseType  string `json:"license_type,omitempty"`  // Type of license
	LicenseCount *int   `json:"license_count,omitempty"` // Number of licenses
}
type ResponseLicensesLicenseUsageDetailsV1PurchasedNetworkLicense struct {
	TotalLicenseCount  *int                                                                              `json:"total_license_count,omitempty"`   // Total number of licenses
	LicenseCountByType *[]ResponseLicensesLicenseUsageDetailsV1PurchasedNetworkLicenseLicenseCountByType `json:"license_count_by_type,omitempty"` //
}
type ResponseLicensesLicenseUsageDetailsV1PurchasedNetworkLicenseLicenseCountByType struct {
	LicenseType  string `json:"license_type,omitempty"`  // Type of license
	LicenseCount *int   `json:"license_count,omitempty"` // Number of licenses
}
type ResponseLicensesLicenseUsageDetailsV1UsedDnaLicense struct {
	TotalLicenseCount  *int                                                                     `json:"total_license_count,omitempty"`   // Total number of licenses
	LicenseCountByType *[]ResponseLicensesLicenseUsageDetailsV1UsedDnaLicenseLicenseCountByType `json:"license_count_by_type,omitempty"` //
}
type ResponseLicensesLicenseUsageDetailsV1UsedDnaLicenseLicenseCountByType struct {
	LicenseType  string `json:"license_type,omitempty"`  // Type of license
	LicenseCount *int   `json:"license_count,omitempty"` // Number of licenses
}
type ResponseLicensesLicenseUsageDetailsV1UsedNetworkLicense struct {
	TotalLicenseCount  *int                                                                         `json:"total_license_count,omitempty"`   // Total number of licenses
	LicenseCountByType *[]ResponseLicensesLicenseUsageDetailsV1UsedNetworkLicenseLicenseCountByType `json:"license_count_by_type,omitempty"` //
}
type ResponseLicensesLicenseUsageDetailsV1UsedNetworkLicenseLicenseCountByType struct {
	LicenseType  string `json:"license_type,omitempty"`  // Type of license
	LicenseCount *int   `json:"license_count,omitempty"` // Number of licenses
}
type ResponseLicensesSmartLicensingDeregistrationV1 struct {
	Response *ResponseLicensesSmartLicensingDeregistrationV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // API version
}
type ResponseLicensesSmartLicensingDeregistrationV1Response struct {
	URL string `json:"url,omitempty"` // URL to track the operation status
}
type ResponseLicensesSystemLicensingLastOperationStatusV1 struct {
	Response *ResponseLicensesSystemLicensingLastOperationStatusV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // API version
}
type ResponseLicensesSystemLicensingLastOperationStatusV1Response struct {
	ID string `json:"id,omitempty"` // The ID of this task

	Status string `json:"status,omitempty"` // Summarizes the status of a task

	IsError *bool `json:"isError,omitempty"` // A boolean indicating if this task has ended with or without error. true indicates a failure, whereas false indicates a success.

	FailureReason string `json:"failureReason,omitempty"` // A textual description indicating the reason why a task has failed

	ErrorCode string `json:"errorCode,omitempty"` // An error code if in case this task has failed in its execution

	LastUpdate *float64 `json:"lastUpdate,omitempty"` // A timestamp of when this task was last updated; as measured in Unix epoch time in milliseconds
}
type ResponseLicensesSystemLicensingRegistrationV1 struct {
	Response *ResponseLicensesSystemLicensingRegistrationV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // API version
}
type ResponseLicensesSystemLicensingRegistrationV1Response struct {
	URL string `json:"url,omitempty"` // URL to track the operation status
}
type ResponseLicensesSmartLicensingRenewOperationV1 struct {
	Response *ResponseLicensesSmartLicensingRenewOperationV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // API version
}
type ResponseLicensesSmartLicensingRenewOperationV1Response struct {
	URL string `json:"url,omitempty"` // URL to track the operation status
}
type ResponseLicensesSystemLicensingStatusV1 struct {
	Response *ResponseLicensesSystemLicensingStatusV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // API version
}
type ResponseLicensesSystemLicensingStatusV1Response struct {
	RegistrationStatus *ResponseLicensesSystemLicensingStatusV1ResponseRegistrationStatus `json:"registrationStatus,omitempty"` //

	AuthorizationStatus *ResponseLicensesSystemLicensingStatusV1ResponseAuthorizationStatus `json:"authorizationStatus,omitempty"` //

	Entitlements *ResponseLicensesSystemLicensingStatusV1ResponseEntitlements `json:"entitlements,omitempty"` //

	SmartAccountID string `json:"smartAccountId,omitempty"` // Smart Account id to which the system is registered

	VirtualAccountID string `json:"virtualAccountId,omitempty"` // Virtual Account id to which the system is registered

	ExportControl string `json:"exportControl,omitempty"` // Export-Controlled setting of Smart Account
}
type ResponseLicensesSystemLicensingStatusV1ResponseRegistrationStatus struct {
	Status string `json:"status,omitempty"` // REGISTERED if the system is registered with CSSM, otherwise UNREGISTERED.

	LastAttemptTimestamp *float64 `json:"lastAttemptTimestamp,omitempty"` // A date and time represented as milliseconds since the Unix epoch.

	ExpiryTimestamp *float64 `json:"expiryTimestamp,omitempty"` // A date and time represented as milliseconds since the Unix epoch.

	NextAttemptTimestamp *float64 `json:"nextAttemptTimestamp,omitempty"` // A date and time represented as milliseconds since the Unix epoch.

	LastAttemptStatus string `json:"lastAttemptStatus,omitempty"` // The last registration request's status

	LastAttemptFailReason string `json:"lastAttemptFailReason,omitempty"` // The reason for last registration request failure
}
type ResponseLicensesSystemLicensingStatusV1ResponseAuthorizationStatus struct {
	Status string `json:"status,omitempty"` // This denotes the authorization status of the system.

	LastAttemptTimestamp *float64 `json:"lastAttemptTimestamp,omitempty"` // A date and time represented as milliseconds since the Unix epoch.

	EvaluationRemainderTimestamp *float64 `json:"evaluationRemainderTimestamp,omitempty"` // A date and time represented as milliseconds since the Unix epoch.

	ExpiryTimestamp *float64 `json:"expiryTimestamp,omitempty"` // A date and time represented as milliseconds since the Unix epoch.

	NextAttemptTimestamp *float64 `json:"nextAttemptTimestamp,omitempty"` // A date and time represented as milliseconds since the Unix epoch.

	LastAttemptStatus string `json:"lastAttemptStatus,omitempty"` // The last authorization request's status

	LastAttemptFailReason string `json:"lastAttemptFailReason,omitempty"` // The reason for last authorization request failure
}
type ResponseLicensesSystemLicensingStatusV1ResponseEntitlements struct {
	Tag string `json:"tag,omitempty"` // Entitlement tag associated with the available licenses

	Description string `json:"description,omitempty"` // Name or description of the license entitlement

	UsageCount *int `json:"usageCount,omitempty"` // Available license count

	Status string `json:"status,omitempty"` // This denotes the authorization status of the available licenses.
}
type RequestLicensesUpdateCSSMConnectionModeV1 struct {
	ConnectionMode string `json:"connectionMode,omitempty"` // The CSSM connection modes of Catalyst Center are DIRECT, ON_PREMISE and SMART_PROXY.

	Parameters *RequestLicensesUpdateCSSMConnectionModeV1Parameters `json:"parameters,omitempty"` //
}
type RequestLicensesUpdateCSSMConnectionModeV1Parameters struct {
	OnPremiseHost string `json:"onPremiseHost,omitempty"` // On-premise CSSM hostname or IP address

	SmartAccountName string `json:"smartAccountName,omitempty"` // On-premise CSSM local smart account name

	ClientID string `json:"clientId,omitempty"` // On-premise CSSM client id

	ClientSecret string `json:"clientSecret,omitempty"` // On-premise CSSM client secret
}
type RequestLicensesUpdateLicenseSettingV1 struct {
	DefaultSmartAccountID string `json:"defaultSmartAccountId,omitempty"` // Default smart account id

	AutoRegistrationVirtualAccountID string `json:"autoRegistrationVirtualAccountId,omitempty"` // Virtual account id
}
type RequestLicensesDeviceDeregistrationV1 struct {
	DeviceUUIDs []string `json:"device_uuids,omitempty"` // Comma separated device ids
}
type RequestLicensesDeviceRegistrationV1 struct {
	DeviceUUIDs []string `json:"device_uuids,omitempty"` // Comma separated device ids
}
type RequestLicensesChangeVirtualAccountV1 struct {
	DeviceUUIDs []string `json:"device_uuids,omitempty"` // Comma separated device ids
}
type RequestLicensesSystemLicensingRegistrationV1 struct {
	SmartAccountID   string `json:"smartAccountId,omitempty"`   // The ID of the Smart Account to which the system is registered
	VirtualAccountID string `json:"virtualAccountId,omitempty"` // The ID of the Virtual Account to which the system is registered
}

//RetrievesCSSMConnectionModeV1 Retrieves CSSM Connection Mode - 1098-3a93-47a9-bf2c
/* Retrieves Cisco Smart Software Manager (CSSM) connection mode setting.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-c-s-s-m-connection-mode
*/
func (s *LicensesService) RetrievesCSSMConnectionModeV1() (*ResponseLicensesRetrievesCSSMConnectionModeV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/connectionModeSetting"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLicensesRetrievesCSSMConnectionModeV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesCSSMConnectionModeV1()
		}
		return nil, response, fmt.Errorf("error with operation RetrievesCSSMConnectionModeV1")
	}

	result := response.Result().(*ResponseLicensesRetrievesCSSMConnectionModeV1)
	return result, response, err

}

//RetrieveLicenseSettingV1 Retrieve license setting - c489-d9a3-4c09-84c7
/* Retrieves license setting Default smart account id and virtual account id for auto registration of devices for smart license flow. If default smart account is not configured, 'defaultSmartAccountId' is 'null'. Similarly, if auto registration of devices for smart license flow is not enabled, 'autoRegistrationVirtualAccountId' is 'null'. For smart proxy connection mode, 'autoRegistrationVirtualAccountId' is always 'null'.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-license-setting
*/
func (s *LicensesService) RetrieveLicenseSettingV1() (*ResponseLicensesRetrieveLicenseSettingV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/licenseSetting"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLicensesRetrieveLicenseSettingV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveLicenseSettingV1()
		}
		return nil, response, fmt.Errorf("error with operation RetrieveLicenseSettingV1")
	}

	result := response.Result().(*ResponseLicensesRetrieveLicenseSettingV1)
	return result, response, err

}

//DeviceCountDetailsV1 Device Count Details - 949a-6983-4c38-9d59
/* Get total number of managed device(s).


@param DeviceCountDetailsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!device-count-details
*/
func (s *LicensesService) DeviceCountDetailsV1(DeviceCountDetailsV1QueryParams *DeviceCountDetailsV1QueryParams) (*ResponseLicensesDeviceCountDetailsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/licenses/device/count"

	queryString, _ := query.Values(DeviceCountDetailsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseLicensesDeviceCountDetailsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeviceCountDetailsV1(DeviceCountDetailsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeviceCountDetailsV1")
	}

	result := response.Result().(*ResponseLicensesDeviceCountDetailsV1)
	return result, response, err

}

//DeviceLicenseSummaryV1 Device License Summary - 529e-2ae0-4a9b-89b1
/* Show license summary of device(s).


@param DeviceLicenseSummaryV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!device-license-summary
*/
func (s *LicensesService) DeviceLicenseSummaryV1(DeviceLicenseSummaryV1QueryParams *DeviceLicenseSummaryV1QueryParams) (*ResponseLicensesDeviceLicenseSummaryV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/licenses/device/summary"

	queryString, _ := query.Values(DeviceLicenseSummaryV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseLicensesDeviceLicenseSummaryV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeviceLicenseSummaryV1(DeviceLicenseSummaryV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeviceLicenseSummaryV1")
	}

	result := response.Result().(*ResponseLicensesDeviceLicenseSummaryV1)
	return result, response, err

}

//DeviceLicenseDetailsV1 Device License Details - dca1-1bc2-4e7b-8c5d
/* Get detailed license information of a device.


@param deviceuuid device_uuid path parameter. Id of device


Documentation Link: https://developer.cisco.com/docs/dna-center/#!device-license-details
*/
func (s *LicensesService) DeviceLicenseDetailsV1(deviceuuid string) (*ResponseLicensesDeviceLicenseDetailsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/licenses/device/{device_uuid}/details"
	path = strings.Replace(path, "{device_uuid}", fmt.Sprintf("%v", deviceuuid), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLicensesDeviceLicenseDetailsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeviceLicenseDetailsV1(deviceuuid)
		}
		return nil, response, fmt.Errorf("error with operation DeviceLicenseDetailsV1")
	}

	result := response.Result().(*ResponseLicensesDeviceLicenseDetailsV1)
	return result, response, err

}

//VirtualAccountDetailsV1 Virtual Account Details - 829e-daf5-4d49-a581
/* Get virtual account details of a smart account.


@param smartaccountTypeID smart_account_id path parameter. Id of smart account


Documentation Link: https://developer.cisco.com/docs/dna-center/#!virtual-account-details
*/
func (s *LicensesService) VirtualAccountDetailsV1(smartaccountTypeID string) (*ResponseLicensesVirtualAccountDetailsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/licenses/smartAccount/{smart_account_id}/virtualAccounts"
	path = strings.Replace(path, "{smart_account_id}", fmt.Sprintf("%v", smartaccountTypeID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLicensesVirtualAccountDetailsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.VirtualAccountDetailsV1(smartaccountTypeID)
		}
		return nil, response, fmt.Errorf("error with operation VirtualAccountDetailsV1")
	}

	result := response.Result().(*ResponseLicensesVirtualAccountDetailsV1)
	return result, response, err

}

//SmartAccountDetailsV1 Smart Account Details - c181-b8db-4c89-adf0
/* Retrieve details of all smart accounts.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!smart-account-details
*/
func (s *LicensesService) SmartAccountDetailsV1() (*ResponseLicensesSmartAccountDetailsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/licenses/smartAccounts"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLicensesSmartAccountDetailsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SmartAccountDetailsV1()
		}
		return nil, response, fmt.Errorf("error with operation SmartAccountDetailsV1")
	}

	result := response.Result().(*ResponseLicensesSmartAccountDetailsV1)
	return result, response, err

}

//LicenseTermDetailsV1 License Term Details - e891-4b73-4ad9-b414
/* Get license term details.


@param smartaccountTypeID smart_account_id path parameter. Id of smart account

@param virtualaccountname virtual_account_name path parameter. Name of virtual account. Putting "All" will give license term detail for all virtual accounts.

@param LicenseTermDetailsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!license-term-details
*/
func (s *LicensesService) LicenseTermDetailsV1(smartaccountTypeID string, virtualaccountname string, LicenseTermDetailsV1QueryParams *LicenseTermDetailsV1QueryParams) (*ResponseLicensesLicenseTermDetailsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/licenses/term/smartAccount/{smart_account_id}/virtualAccount/{virtual_account_name}"
	path = strings.Replace(path, "{smart_account_id}", fmt.Sprintf("%v", smartaccountTypeID), -1)
	path = strings.Replace(path, "{virtual_account_name}", fmt.Sprintf("%v", virtualaccountname), -1)

	queryString, _ := query.Values(LicenseTermDetailsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseLicensesLicenseTermDetailsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.LicenseTermDetailsV1(smartaccountTypeID, virtualaccountname, LicenseTermDetailsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation LicenseTermDetailsV1")
	}

	result := response.Result().(*ResponseLicensesLicenseTermDetailsV1)
	return result, response, err

}

//LicenseUsageDetailsV1 License Usage Details - 418a-6b43-4e29-bfe5
/* Get count of purchased and in use Cisco DNA and Network licenses.


@param smartaccountTypeID smart_account_id path parameter. Id of smart account

@param virtualaccountname virtual_account_name path parameter. Name of virtual account. Putting "All" will give license term detail for all virtual accounts.

@param LicenseUsageDetailsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!license-usage-details
*/
func (s *LicensesService) LicenseUsageDetailsV1(smartaccountTypeID string, virtualaccountname string, LicenseUsageDetailsV1QueryParams *LicenseUsageDetailsV1QueryParams) (*ResponseLicensesLicenseUsageDetailsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/licenses/usage/smartAccount/{smart_account_id}/virtualAccount/{virtual_account_name}"
	path = strings.Replace(path, "{smart_account_id}", fmt.Sprintf("%v", smartaccountTypeID), -1)
	path = strings.Replace(path, "{virtual_account_name}", fmt.Sprintf("%v", virtualaccountname), -1)

	queryString, _ := query.Values(LicenseUsageDetailsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseLicensesLicenseUsageDetailsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.LicenseUsageDetailsV1(smartaccountTypeID, virtualaccountname, LicenseUsageDetailsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation LicenseUsageDetailsV1")
	}

	result := response.Result().(*ResponseLicensesLicenseUsageDetailsV1)
	return result, response, err

}

//SystemLicensingLastOperationStatusV1 System Licensing Last Operation Status - 458d-da00-4e18-b4ec
/* Retrieves the status of the last system licensing operation.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!system-licensing-last-operation-status
*/
func (s *LicensesService) SystemLicensingLastOperationStatusV1() (*ResponseLicensesSystemLicensingLastOperationStatusV1, *resty.Response, error) {
	path := "/dna/system/api/v1/license/lastOperationStatus"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLicensesSystemLicensingLastOperationStatusV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SystemLicensingLastOperationStatusV1()
		}
		return nil, response, fmt.Errorf("error with operation SystemLicensingLastOperationStatusV1")
	}

	result := response.Result().(*ResponseLicensesSystemLicensingLastOperationStatusV1)
	return result, response, err

}

//SystemLicensingStatusV1 System Licensing Status - 64a5-0bcb-414b-89e1
/* Fetches registration status, authorization status and entitlements of the system with Cisco Smart Software Manage (CSSM).



Documentation Link: https://developer.cisco.com/docs/dna-center/#!system-licensing-status
*/
func (s *LicensesService) SystemLicensingStatusV1() (*ResponseLicensesSystemLicensingStatusV1, *resty.Response, error) {
	path := "/dna/system/api/v1/license/status"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLicensesSystemLicensingStatusV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SystemLicensingStatusV1()
		}
		return nil, response, fmt.Errorf("error with operation SystemLicensingStatusV1")
	}

	result := response.Result().(*ResponseLicensesSystemLicensingStatusV1)
	return result, response, err

}

//ChangeVirtualAccountV1 Change Virtual Account - bea7-4a0b-4778-8c89
/* Transfer device(s) from one virtual account to another within same smart account.


@param smartaccountTypeID smart_account_id path parameter. Id of smart account

@param virtualaccountname virtual_account_name path parameter. Name of target virtual account


Documentation Link: https://developer.cisco.com/docs/dna-center/#!change-virtual-account
*/
func (s *LicensesService) ChangeVirtualAccountV1(smartaccountTypeID string, virtualaccountname string, requestLicensesChangeVirtualAccountV1 *RequestLicensesChangeVirtualAccountV1) (*ResponseLicensesChangeVirtualAccountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/licenses/smartAccount/{smart_account_id}/virtualAccount/{virtual_account_name}/device/transfer"
	path = strings.Replace(path, "{smart_account_id}", fmt.Sprintf("%v", smartaccountTypeID), -1)
	path = strings.Replace(path, "{virtual_account_name}", fmt.Sprintf("%v", virtualaccountname), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestLicensesChangeVirtualAccountV1).
		SetResult(&ResponseLicensesChangeVirtualAccountV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ChangeVirtualAccountV1(smartaccountTypeID, virtualaccountname, requestLicensesChangeVirtualAccountV1)
		}

		return nil, response, fmt.Errorf("error with operation ChangeVirtualAccountV1")
	}

	result := response.Result().(*ResponseLicensesChangeVirtualAccountV1)
	return result, response, err

}

//SmartLicensingDeregistrationV1 Smart Licensing Deregistration - 8489-ea73-4838-a40e
/* Deregisters the system with Cisco Smart Software Manager (CSSM)



Documentation Link: https://developer.cisco.com/docs/dna-center/#!smart-licensing-deregistration
*/
func (s *LicensesService) SmartLicensingDeregistrationV1() (*ResponseLicensesSmartLicensingDeregistrationV1, *resty.Response, error) {
	path := "/dna/system/api/v1/license/deregister"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLicensesSmartLicensingDeregistrationV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.SmartLicensingDeregistrationV1()
		}

		return nil, response, fmt.Errorf("error with operation SmartLicensingDeregistrationV1")
	}

	result := response.Result().(*ResponseLicensesSmartLicensingDeregistrationV1)
	return result, response, err

}

//SystemLicensingRegistrationV1 System Licensing Registration - eda9-ea64-4ce8-ad1e
/* Registers the system with Cisco Smart Software Manager (CSSM)



Documentation Link: https://developer.cisco.com/docs/dna-center/#!system-licensing-registration
*/
func (s *LicensesService) SystemLicensingRegistrationV1(requestLicensesSystemLicensingRegistrationV1 *RequestLicensesSystemLicensingRegistrationV1) (*ResponseLicensesSystemLicensingRegistrationV1, *resty.Response, error) {
	path := "/dna/system/api/v1/license/register"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestLicensesSystemLicensingRegistrationV1).
		SetResult(&ResponseLicensesSystemLicensingRegistrationV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.SystemLicensingRegistrationV1(requestLicensesSystemLicensingRegistrationV1)
		}

		return nil, response, fmt.Errorf("error with operation SystemLicensingRegistrationV1")
	}

	result := response.Result().(*ResponseLicensesSystemLicensingRegistrationV1)
	return result, response, err

}

//SmartLicensingRenewOperationV1 Smart Licensing Renew Operation - d288-fb7f-45b8-87d7
/* Renews license registration and authorization status of the system with Cisco Smart Software Manager (CSSM)



Documentation Link: https://developer.cisco.com/docs/dna-center/#!smart-licensing-renew-operation
*/
func (s *LicensesService) SmartLicensingRenewOperationV1() (*ResponseLicensesSmartLicensingRenewOperationV1, *resty.Response, error) {
	path := "/dna/system/api/v1/license/renew"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLicensesSmartLicensingRenewOperationV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.SmartLicensingRenewOperationV1()
		}

		return nil, response, fmt.Errorf("error with operation SmartLicensingRenewOperationV1")
	}

	result := response.Result().(*ResponseLicensesSmartLicensingRenewOperationV1)
	return result, response, err

}

//UpdateCSSMConnectionModeV1 Update CSSM Connection Mode - cfb4-18c6-4eb8-959a
/* Update Cisco Smart Software Manager (CSSM) connection mode for the system.


 */
func (s *LicensesService) UpdateCSSMConnectionModeV1(requestLicensesUpdateCSSMConnectionModeV1 *RequestLicensesUpdateCSSMConnectionModeV1) (*ResponseLicensesUpdateCSSMConnectionModeV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/connectionModeSetting"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestLicensesUpdateCSSMConnectionModeV1).
		SetResult(&ResponseLicensesUpdateCSSMConnectionModeV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateCSSMConnectionModeV1(requestLicensesUpdateCSSMConnectionModeV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateCSSMConnectionModeV1")
	}

	result := response.Result().(*ResponseLicensesUpdateCSSMConnectionModeV1)
	return result, response, err

}

//UpdateLicenseSettingV1 Update license setting - 97ae-8980-475a-961e
/* Update license setting Configure default smart account id  and/or virtual account id for auto registration of devices for smart license flow. Virtual account should be part of default smart account. Default smart account id cannot be set to 'null'. Auto registration of devices for smart license flow is applicable only for direct or on-prem SSM connection mode.


 */
func (s *LicensesService) UpdateLicenseSettingV1(requestLicensesUpdateLicenseSettingV1 *RequestLicensesUpdateLicenseSettingV1) (*ResponseLicensesUpdateLicenseSettingV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/licenseSetting"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestLicensesUpdateLicenseSettingV1).
		SetResult(&ResponseLicensesUpdateLicenseSettingV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateLicenseSettingV1(requestLicensesUpdateLicenseSettingV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateLicenseSettingV1")
	}

	result := response.Result().(*ResponseLicensesUpdateLicenseSettingV1)
	return result, response, err

}

//DeviceDeregistrationV1 Device Deregistration - 8c82-dad4-49ba-b8eb
/* Deregister device(s) from CSSM(Cisco Smart Software Manager).


 */
func (s *LicensesService) DeviceDeregistrationV1(requestLicensesDeviceDeregistrationV1 *RequestLicensesDeviceDeregistrationV1) (*ResponseLicensesDeviceDeregistrationV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/licenses/smartAccount/virtualAccount/deregister"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestLicensesDeviceDeregistrationV1).
		SetResult(&ResponseLicensesDeviceDeregistrationV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeviceDeregistrationV1(requestLicensesDeviceDeregistrationV1)
		}
		return nil, response, fmt.Errorf("error with operation DeviceDeregistrationV1")
	}

	result := response.Result().(*ResponseLicensesDeviceDeregistrationV1)
	return result, response, err

}

//DeviceRegistrationV1 Device Registration - a08b-eae5-47fb-95e3
/* Register device(s) in CSSM(Cisco Smart Software Manager).


@param virtualaccountname virtual_account_name path parameter. Name of virtual account

*/
func (s *LicensesService) DeviceRegistrationV1(virtualaccountname string, requestLicensesDeviceRegistrationV1 *RequestLicensesDeviceRegistrationV1) (*ResponseLicensesDeviceRegistrationV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/licenses/smartAccount/virtualAccount/{virtual_account_name}/register"
	path = strings.Replace(path, "{virtual_account_name}", fmt.Sprintf("%v", virtualaccountname), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestLicensesDeviceRegistrationV1).
		SetResult(&ResponseLicensesDeviceRegistrationV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeviceRegistrationV1(virtualaccountname, requestLicensesDeviceRegistrationV1)
		}
		return nil, response, fmt.Errorf("error with operation DeviceRegistrationV1")
	}

	result := response.Result().(*ResponseLicensesDeviceRegistrationV1)
	return result, response, err

}

// Alias Function
/*
This method acts as an alias for the method `VirtualAccountDetailsV1`
*/
func (s *LicensesService) VirtualAccountDetails(smartaccountTypeID string) (*ResponseLicensesVirtualAccountDetailsV1, *resty.Response, error) {
	return s.VirtualAccountDetailsV1(smartaccountTypeID)
}

// Alias Function
/*
This method acts as an alias for the method `LicenseTermDetailsV1`
*/
func (s *LicensesService) LicenseTermDetails(smartaccountTypeID string, virtualaccountname string, LicenseTermDetailsV1QueryParams *LicenseTermDetailsV1QueryParams) (*ResponseLicensesLicenseTermDetailsV1, *resty.Response, error) {
	return s.LicenseTermDetailsV1(smartaccountTypeID, virtualaccountname, LicenseTermDetailsV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `SystemLicensingRegistrationV1`
*/
func (s *LicensesService) SystemLicensingRegistration(requestLicensesSystemLicensingRegistrationV1 *RequestLicensesSystemLicensingRegistrationV1) (*ResponseLicensesSystemLicensingRegistrationV1, *resty.Response, error) {
	return s.SystemLicensingRegistrationV1(requestLicensesSystemLicensingRegistrationV1)
}

// Alias Function
/*
This method acts as an alias for the method `DeviceCountDetailsV1`
*/
func (s *LicensesService) DeviceCountDetails(DeviceCountDetailsV1QueryParams *DeviceCountDetailsV1QueryParams) (*ResponseLicensesDeviceCountDetailsV1, *resty.Response, error) {
	return s.DeviceCountDetailsV1(DeviceCountDetailsV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `SmartLicensingDeregistrationV1`
*/
func (s *LicensesService) SmartLicensingDeregistration() (*ResponseLicensesSmartLicensingDeregistrationV1, *resty.Response, error) {
	return s.SmartLicensingDeregistrationV1()
}

// Alias Function
/*
This method acts as an alias for the method `RetrieveLicenseSettingV1`
*/
func (s *LicensesService) RetrieveLicenseSetting() (*ResponseLicensesRetrieveLicenseSettingV1, *resty.Response, error) {
	return s.RetrieveLicenseSettingV1()
}

// Alias Function
/*
This method acts as an alias for the method `SmartLicensingRenewOperationV1`
*/
func (s *LicensesService) SmartLicensingRenewOperation() (*ResponseLicensesSmartLicensingRenewOperationV1, *resty.Response, error) {
	return s.SmartLicensingRenewOperationV1()
}

// Alias Function
/*
This method acts as an alias for the method `DeviceLicenseSummaryV1`
*/
func (s *LicensesService) DeviceLicenseSummary(DeviceLicenseSummaryV1QueryParams *DeviceLicenseSummaryV1QueryParams) (*ResponseLicensesDeviceLicenseSummaryV1, *resty.Response, error) {
	return s.DeviceLicenseSummaryV1(DeviceLicenseSummaryV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `ChangeVirtualAccountV1`
*/
func (s *LicensesService) ChangeVirtualAccount(smartaccountTypeID string, virtualaccountname string, requestLicensesChangeVirtualAccountV1 *RequestLicensesChangeVirtualAccountV1) (*ResponseLicensesChangeVirtualAccountV1, *resty.Response, error) {
	return s.ChangeVirtualAccountV1(smartaccountTypeID, virtualaccountname, requestLicensesChangeVirtualAccountV1)
}

// Alias Function
/*
This method acts as an alias for the method `DeviceRegistrationV1`
*/
func (s *LicensesService) DeviceRegistration(virtualaccountname string, requestLicensesDeviceRegistrationV1 *RequestLicensesDeviceRegistrationV1) (*ResponseLicensesDeviceRegistrationV1, *resty.Response, error) {
	return s.DeviceRegistrationV1(virtualaccountname, requestLicensesDeviceRegistrationV1)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesCSSMConnectionModeV1`
*/
func (s *LicensesService) RetrievesCSSMConnectionMode() (*ResponseLicensesRetrievesCSSMConnectionModeV1, *resty.Response, error) {
	return s.RetrievesCSSMConnectionModeV1()
}

// Alias Function
/*
This method acts as an alias for the method `UpdateCSSMConnectionModeV1`
*/
func (s *LicensesService) UpdateCSSMConnectionMode(requestLicensesUpdateCSSMConnectionModeV1 *RequestLicensesUpdateCSSMConnectionModeV1) (*ResponseLicensesUpdateCSSMConnectionModeV1, *resty.Response, error) {
	return s.UpdateCSSMConnectionModeV1(requestLicensesUpdateCSSMConnectionModeV1)
}

// Alias Function
/*
This method acts as an alias for the method `SmartAccountDetailsV1`
*/
func (s *LicensesService) SmartAccountDetails() (*ResponseLicensesSmartAccountDetailsV1, *resty.Response, error) {
	return s.SmartAccountDetailsV1()
}

// Alias Function
/*
This method acts as an alias for the method `SystemLicensingLastOperationStatusV1`
*/
func (s *LicensesService) SystemLicensingLastOperationStatus() (*ResponseLicensesSystemLicensingLastOperationStatusV1, *resty.Response, error) {
	return s.SystemLicensingLastOperationStatusV1()
}

// Alias Function
/*
This method acts as an alias for the method `UpdateLicenseSettingV1`
*/
func (s *LicensesService) UpdateLicenseSetting(requestLicensesUpdateLicenseSettingV1 *RequestLicensesUpdateLicenseSettingV1) (*ResponseLicensesUpdateLicenseSettingV1, *resty.Response, error) {
	return s.UpdateLicenseSettingV1(requestLicensesUpdateLicenseSettingV1)
}

// Alias Function
/*
This method acts as an alias for the method `SystemLicensingStatusV1`
*/
func (s *LicensesService) SystemLicensingStatus() (*ResponseLicensesSystemLicensingStatusV1, *resty.Response, error) {
	return s.SystemLicensingStatusV1()
}

// Alias Function
/*
This method acts as an alias for the method `DeviceLicenseDetailsV1`
*/
func (s *LicensesService) DeviceLicenseDetails(deviceuuid string) (*ResponseLicensesDeviceLicenseDetailsV1, *resty.Response, error) {
	return s.DeviceLicenseDetailsV1(deviceuuid)
}

// Alias Function
/*
This method acts as an alias for the method `DeviceDeregistrationV1`
*/
func (s *LicensesService) DeviceDeregistration(requestLicensesDeviceDeregistrationV1 *RequestLicensesDeviceDeregistrationV1) (*ResponseLicensesDeviceDeregistrationV1, *resty.Response, error) {
	return s.DeviceDeregistrationV1(requestLicensesDeviceDeregistrationV1)
}

// Alias Function
/*
This method acts as an alias for the method `LicenseUsageDetailsV1`
*/
func (s *LicensesService) LicenseUsageDetails(smartaccountTypeID string, virtualaccountname string, LicenseUsageDetailsV1QueryParams *LicenseUsageDetailsV1QueryParams) (*ResponseLicensesLicenseUsageDetailsV1, *resty.Response, error) {
	return s.LicenseUsageDetailsV1(smartaccountTypeID, virtualaccountname, LicenseUsageDetailsV1QueryParams)
}
