package catalyst

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SystemSettingsService service

type GetAuthenticationAndPolicyServersV1QueryParams struct {
	IsIseEnabled bool   `url:"isIseEnabled,omitempty"` //Valid values are : true, false
	State        string `url:"state,omitempty"`        //Valid values are: ACTIVE, INACTIVE, RBAC_SUCCESS, RBAC_FAILURE, DELETED, FAILED, INPROGRESS
	Role         string `url:"role,omitempty"`         //Authentication and Policy Server Role (Example: primary, secondary)
}

type ResponseSystemSettingsAddAuthenticationAndPolicyServerAccessConfigurationV1 struct {
	Response *ResponseSystemSettingsAddAuthenticationAndPolicyServerAccessConfigurationV1Response `json:"response,omitempty"` //
	Version  string                                                                               `json:"version,omitempty"`  // Version
}
type ResponseSystemSettingsAddAuthenticationAndPolicyServerAccessConfigurationV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseSystemSettingsGetAuthenticationAndPolicyServersV1 struct {
	Response *[]ResponseSystemSettingsGetAuthenticationAndPolicyServersV1Response `json:"response,omitempty"` //
}
type ResponseSystemSettingsGetAuthenticationAndPolicyServersV1Response struct {
	IPAddress            string                                                                           `json:"ipAddress,omitempty"`            // IP address of authentication and policy server
	SharedSecret         string                                                                           `json:"sharedSecret,omitempty"`         // Shared secret between devices and authentication and policy server
	Protocol             string                                                                           `json:"protocol,omitempty"`             // Type of protocol for authentication and policy server
	Role                 string                                                                           `json:"role,omitempty"`                 // Role of authentication and policy server (Example: primary, secondary)
	Port                 *int                                                                             `json:"port,omitempty"`                 // Port of TACACS server (Default: 49)
	AuthenticationPort   *int                                                                             `json:"authenticationPort,omitempty"`   // Authentication port of RADIUS server (Default: 1812)
	AccountingPort       *int                                                                             `json:"accountingPort,omitempty"`       // Accounting port of RADIUS server (Default: 1813)
	Retries              *int                                                                             `json:"retries,omitempty"`              // Number of communication retries between devices and authentication and policy server (Default: 3)
	TimeoutSeconds       *int                                                                             `json:"timeoutSeconds,omitempty"`       // Number of seconds before timing out between devices and authentication and policy server (Default: 4 seconds)
	IsIseEnabled         *bool                                                                            `json:"isIseEnabled,omitempty"`         // If server type is ISE, value will be true otherwise false
	InstanceUUID         string                                                                           `json:"instanceUuid,omitempty"`         // Internal record identifier
	State                string                                                                           `json:"state,omitempty"`                // State of authentication and policy server
	CiscoIseDtos         *[]ResponseSystemSettingsGetAuthenticationAndPolicyServersV1ResponseCiscoIseDtos `json:"ciscoIseDtos,omitempty"`         //
	EncryptionScheme     string                                                                           `json:"encryptionScheme,omitempty"`     // Type of encryption scheme for additional security
	MessageKey           string                                                                           `json:"messageKey,omitempty"`           // Message key used to encrypt shared secret
	EncryptionKey        string                                                                           `json:"encryptionKey,omitempty"`        // Encryption key used to encrypt shared secret
	UseDnacCertForPxgrid *bool                                                                            `json:"useDnacCertForPxgrid,omitempty"` // Use DNAC Certificate For Pxgrid
	IseEnabled           *bool                                                                            `json:"iseEnabled,omitempty"`           // If server type is ISE, value will be true otherwise false
	PxgridEnabled        *bool                                                                            `json:"pxgridEnabled,omitempty"`        // If pxgrid enabled, value will be true otherwise false
	RbacUUID             string                                                                           `json:"rbacUuid,omitempty"`             // Internal use only
	MultiDnacEnabled     *bool                                                                            `json:"multiDnacEnabled,omitempty"`     // Internal use only
}
type ResponseSystemSettingsGetAuthenticationAndPolicyServersV1ResponseCiscoIseDtos struct {
	SubscriberName             string                                                                                                   `json:"subscriberName,omitempty"`             // Subscriber name of the ISE server (Example: pxgrid_client_1662589467)
	Description                string                                                                                                   `json:"description,omitempty"`                // Description about the ISE server
	Password                   string                                                                                                   `json:"password,omitempty"`                   // For security reasons the value will always be empty
	UserName                   string                                                                                                   `json:"userName,omitempty"`                   // User name of the ISE server
	Fqdn                       string                                                                                                   `json:"fqdn,omitempty"`                       // Fully-qualified domain name of the ISE server (Example: xi-62.my.com)
	IPAddress                  string                                                                                                   `json:"ipAddress,omitempty"`                  // IP Address of the ISE server
	TrustState                 string                                                                                                   `json:"trustState,omitempty"`                 // Trust State between DNAC and the ISE server
	InstanceUUID               string                                                                                                   `json:"instanceUuid,omitempty"`               // Internal record identifier
	SSHkey                     string                                                                                                   `json:"sshkey,omitempty"`                     // For security reasons the value will always be empty
	Type                       string                                                                                                   `json:"type,omitempty"`                       // Type (Example: ISE)
	FailureReason              string                                                                                                   `json:"failureReason,omitempty"`              // Reason for integration failure between DNAC and the ISE server
	Role                       string                                                                                                   `json:"role,omitempty"`                       // Role of the ISE server
	ExternalCiscoIseIPAddrDtos *ResponseSystemSettingsGetAuthenticationAndPolicyServersV1ResponseCiscoIseDtosExternalCiscoIseIPAddrDtos `json:"externalCiscoIseIpAddrDtos,omitempty"` //
}
type ResponseSystemSettingsGetAuthenticationAndPolicyServersV1ResponseCiscoIseDtosExternalCiscoIseIPAddrDtos struct {
	Type                        string                                                                                                                                `json:"type,omitempty"`                        // Type
	ExternalCiscoIseIPAddresses *[]ResponseSystemSettingsGetAuthenticationAndPolicyServersV1ResponseCiscoIseDtosExternalCiscoIseIPAddrDtosExternalCiscoIseIPAddresses `json:"externalCiscoIseIpAddresses,omitempty"` //
}
type ResponseSystemSettingsGetAuthenticationAndPolicyServersV1ResponseCiscoIseDtosExternalCiscoIseIPAddrDtosExternalCiscoIseIPAddresses struct {
	ExternalIPAddress string `json:"externalIpAddress,omitempty"` // External IP Address
}
type ResponseSystemSettingsDeleteAuthenticationAndPolicyServerAccessConfigurationV1 struct {
	Response *ResponseSystemSettingsDeleteAuthenticationAndPolicyServerAccessConfigurationV1Response `json:"response,omitempty"` //
	Version  string                                                                                  `json:"version,omitempty"`  // Version
}
type ResponseSystemSettingsDeleteAuthenticationAndPolicyServerAccessConfigurationV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseSystemSettingsEditAuthenticationAndPolicyServerAccessConfigurationV1 struct {
	Response *ResponseSystemSettingsEditAuthenticationAndPolicyServerAccessConfigurationV1Response `json:"response,omitempty"` //
	Version  string                                                                                `json:"version,omitempty"`  // Version
}
type ResponseSystemSettingsEditAuthenticationAndPolicyServerAccessConfigurationV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseSystemSettingsAcceptCiscoIseServerCertificateForCiscoIseServerIntegrationV1 interface{}
type ResponseSystemSettingsCreatesConfigurationDetailsOfTheExternalIPAMServerV1 struct {
	Response *ResponseSystemSettingsCreatesConfigurationDetailsOfTheExternalIPAMServerV1Response `json:"response,omitempty"` //
	Version  string                                                                              `json:"version,omitempty"`  // Version
}
type ResponseSystemSettingsCreatesConfigurationDetailsOfTheExternalIPAMServerV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseSystemSettingsRetrievesConfigurationDetailsOfTheExternalIPAMServerV1 struct {
	Response *ResponseSystemSettingsRetrievesConfigurationDetailsOfTheExternalIPAMServerV1Response `json:"response,omitempty"` //
	Version  string                                                                                `json:"version,omitempty"`  // Version
}
type ResponseSystemSettingsRetrievesConfigurationDetailsOfTheExternalIPAMServerV1Response struct {
	Provider   string `json:"provider,omitempty"`   // Type of external IPAM. Can be either INFOBLOX, BLUECAT or GENERIC.
	ServerName string `json:"serverName,omitempty"` // A descriptive name of this external server, used for identification purposes
	ServerURL  string `json:"serverUrl,omitempty"`  // The URL of this external server
	State      string `json:"state,omitempty"`      // State of the the external IPAM.* OK indicates success of most recent periodic communication check with external IPAM.* CRITICAL indicates failure of most recent attempt to communicate with the external IPAM.* SYNCHRONIZING indicates that the process of synchronizing the external IPAM database with the local IPAM database is running and all other IPAM processes will be blocked until the completes.* DISCONNECTED indicates the external IPAM is no longer being used.
	UserName   string `json:"userName,omitempty"`   // The external IPAM server login username
	View       string `json:"view,omitempty"`       // The view under which pools are created in the external IPAM server.
}
type ResponseSystemSettingsDeletesConfigurationDetailsOfTheExternalIPAMServerV1 struct {
	Response *ResponseSystemSettingsDeletesConfigurationDetailsOfTheExternalIPAMServerV1Response `json:"response,omitempty"` //
	Version  string                                                                              `json:"version,omitempty"`  // Version
}
type ResponseSystemSettingsDeletesConfigurationDetailsOfTheExternalIPAMServerV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseSystemSettingsUpdatesConfigurationDetailsOfTheExternalIPAMServerV1 struct {
	Response *ResponseSystemSettingsUpdatesConfigurationDetailsOfTheExternalIPAMServerV1Response `json:"response,omitempty"` //
	Version  string                                                                              `json:"version,omitempty"`  // Version
}
type ResponseSystemSettingsUpdatesConfigurationDetailsOfTheExternalIPAMServerV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseSystemSettingsCiscoIseServerIntegrationStatusV1 struct {
	AAAServerSettingID  string                                                          `json:"aaaServerSettingId,omitempty"`  // Cisco ISE Server setting identifier (E.g. 867e46c9-f8f5-40b1-8de2-62f7744f75f6)
	OverallStatus       string                                                          `json:"overallStatus,omitempty"`       // Cisco ISE Server integration status
	OverallErrorMessage string                                                          `json:"overallErrorMessage,omitempty"` // Cisco ISE Server integration failure message
	Steps               *[]ResponseSystemSettingsCiscoIseServerIntegrationStatusV1Steps `json:"steps,omitempty"`               //
}
type ResponseSystemSettingsCiscoIseServerIntegrationStatusV1Steps struct {
	StepID             string `json:"stepId,omitempty"`             // Cisco ISE Server integration step identifier (E.g. 1)
	StepOrder          *int   `json:"stepOrder,omitempty"`          // Cisco ISE Server integration step order (E.g. 1)
	StepName           string `json:"stepName,omitempty"`           // Cisco ISE Server integration step name
	StepDescription    string `json:"stepDescription,omitempty"`    // Cisco ISE Server step description
	StepStatus         string `json:"stepStatus,omitempty"`         // Cisco ISE Server integration step status
	CertAcceptedByUser *bool  `json:"certAcceptedByUser,omitempty"` // If user accept Cisco ISE Server certificate, value will be true otherwise it will be false
	StepTime           *int   `json:"stepTime,omitempty"`           // Last updated epoc time  by the step (E.g. 1677745739314)
}
type ResponseSystemSettingsCustomPromptSupportGETAPIV1 struct {
	Response *ResponseSystemSettingsCustomPromptSupportGETAPIV1Response `json:"response,omitempty"` //
	Version  string                                                     `json:"version,omitempty"`  // Version
}
type ResponseSystemSettingsCustomPromptSupportGETAPIV1Response struct {
	CustomUsernamePrompt  string `json:"customUsernamePrompt,omitempty"`  // Username for Custom Prompt
	CustomPasswordPrompt  string `json:"customPasswordPrompt,omitempty"`  // Password for Custom Prompt
	DefaultUsernamePrompt string `json:"defaultUsernamePrompt,omitempty"` // Default Username for Custom Prompt
	DefaultPasswordPrompt string `json:"defaultPasswordPrompt,omitempty"` // Default Password for Custom Prompt
}
type ResponseSystemSettingsCustomPromptPOSTAPIV1 struct {
	Response *ResponseSystemSettingsCustomPromptPOSTAPIV1Response `json:"response,omitempty"` //
	Version  string                                               `json:"version,omitempty"`  // Version
}
type ResponseSystemSettingsCustomPromptPOSTAPIV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseSystemSettingsSetProvisioningSettingsV1 struct {
	Version  string                                                   `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSystemSettingsSetProvisioningSettingsV1Response `json:"response,omitempty"` //
}
type ResponseSystemSettingsSetProvisioningSettingsV1Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSystemSettingsGetProvisioningSettingsV1 struct {
	Response *ResponseSystemSettingsGetProvisioningSettingsV1Response `json:"response,omitempty"` //
	Version  string                                                   `json:"version,omitempty"`  //
}
type ResponseSystemSettingsGetProvisioningSettingsV1Response struct {
	RequireItsmApproval *bool `json:"requireItsmApproval,omitempty"` // If require ITSM approval is enabled, the planned configurations must be submitted for ITSM approval. Also if enabled, requirePreview will default to enabled.
	RequirePreview      *bool `json:"requirePreview,omitempty"`      // If require preview is enabled, the device configurations must be reviewed before deploying them
}
type RequestSystemSettingsAddAuthenticationAndPolicyServerAccessConfigurationV1 struct {
	AuthenticationPort         *int                                                                                                    `json:"authenticationPort,omitempty"`         // Authentication port of RADIUS server (readonly). The range is from 1 to 65535. E.g. 1812
	AccountingPort             *int                                                                                                    `json:"accountingPort,omitempty"`             // Accounting port of RADIUS server (readonly). The range is from 1 to 65535. E.g. 1813
	CiscoIseDtos               *[]RequestSystemSettingsAddAuthenticationAndPolicyServerAccessConfigurationV1CiscoIseDtos               `json:"ciscoIseDtos,omitempty"`               //
	IPAddress                  string                                                                                                  `json:"ipAddress,omitempty"`                  // IP address of authentication and policy server (readonly)
	PxgridEnabled              *bool                                                                                                   `json:"pxgridEnabled,omitempty"`              // Value true for enable, false for disable. Default value is true
	UseDnacCertForPxgrid       *bool                                                                                                   `json:"useDnacCertForPxgrid,omitempty"`       // Value true to use DNAC certificate for Pxgrid. Default value is false
	IsIseEnabled               *bool                                                                                                   `json:"isIseEnabled,omitempty"`               // Value true for Cisco ISE Server (readonly). Default value is false
	Port                       *int                                                                                                    `json:"port,omitempty"`                       // Port of TACACS server (readonly). The range is from 1 to 65535
	Protocol                   string                                                                                                  `json:"protocol,omitempty"`                   // Type of protocol for authentication and policy server. If already saved with RADIUS, can update to RADIUS_TACACS. If already saved with TACACS, can update to RADIUS_TACACS
	Retries                    string                                                                                                  `json:"retries,omitempty"`                    // Number of communication retries between devices and authentication and policy server. The range is from 1 to 3
	Role                       string                                                                                                  `json:"role,omitempty"`                       // Role of authentication and policy server (readonly). E.g. primary, secondary
	SharedSecret               string                                                                                                  `json:"sharedSecret,omitempty"`               // Shared secret between devices and authentication and policy server (readonly)
	TimeoutSeconds             string                                                                                                  `json:"timeoutSeconds,omitempty"`             // Number of seconds before timing out between devices and authentication and policy server. The range is from 2 to 20
	EncryptionScheme           string                                                                                                  `json:"encryptionScheme,omitempty"`           // Type of encryption scheme for additional security (readonly)
	MessageKey                 string                                                                                                  `json:"messageKey,omitempty"`                 // Message key used to encrypt shared secret (readonly)
	EncryptionKey              string                                                                                                  `json:"encryptionKey,omitempty"`              // Encryption key used to encrypt shared secret (readonly)
	ExternalCiscoIseIPAddrDtos *[]RequestSystemSettingsAddAuthenticationAndPolicyServerAccessConfigurationV1ExternalCiscoIseIPAddrDtos `json:"externalCiscoIseIpAddrDtos,omitempty"` //
}
type RequestSystemSettingsAddAuthenticationAndPolicyServerAccessConfigurationV1CiscoIseDtos struct {
	Description    string `json:"description,omitempty"`    // Description about the Cisco ISE server
	Fqdn           string `json:"fqdn,omitempty"`           // Fully-qualified domain name of the Cisco ISE server (readonly). E.g. xi-62.my.com
	Password       string `json:"password,omitempty"`       // Password of the Cisco ISE server
	SSHkey         string `json:"sshkey,omitempty"`         // SSH key of the Cisco ISE server
	IPAddress      string `json:"ipAddress,omitempty"`      // IP Address of the Cisco ISE Server (readonly)
	SubscriberName string `json:"subscriberName,omitempty"` // Subscriber name of the Cisco ISE server (readonly). E.g. pxgrid_client_1662589467
	UserName       string `json:"userName,omitempty"`       // User name of the Cisco ISE server
}
type RequestSystemSettingsAddAuthenticationAndPolicyServerAccessConfigurationV1ExternalCiscoIseIPAddrDtos struct {
	ExternalCiscoIseIPAddresses *[]RequestSystemSettingsAddAuthenticationAndPolicyServerAccessConfigurationV1ExternalCiscoIseIPAddrDtosExternalCiscoIseIPAddresses `json:"externalCiscoIseIpAddresses,omitempty"` //
	Type                        string                                                                                                                             `json:"type,omitempty"`                        // Type
}
type RequestSystemSettingsAddAuthenticationAndPolicyServerAccessConfigurationV1ExternalCiscoIseIPAddrDtosExternalCiscoIseIPAddresses struct {
	ExternalIPAddress string `json:"externalIpAddress,omitempty"` // External IP Address
}
type RequestSystemSettingsEditAuthenticationAndPolicyServerAccessConfigurationV1 struct {
	AuthenticationPort         *int                                                                                                     `json:"authenticationPort,omitempty"`         // Authentication port of RADIUS server (readonly). The range is from 1 to 65535. E.g. 1812
	AccountingPort             *int                                                                                                     `json:"accountingPort,omitempty"`             // Accounting port of RADIUS server (readonly). The range is from 1 to 65535. E.g. 1813
	CiscoIseDtos               *[]RequestSystemSettingsEditAuthenticationAndPolicyServerAccessConfigurationV1CiscoIseDtos               `json:"ciscoIseDtos,omitempty"`               //
	IPAddress                  string                                                                                                   `json:"ipAddress,omitempty"`                  // IP address of authentication and policy server (readonly)
	PxgridEnabled              *bool                                                                                                    `json:"pxgridEnabled,omitempty"`              // Value true for enable, false for disable. Default value is true
	UseDnacCertForPxgrid       *bool                                                                                                    `json:"useDnacCertForPxgrid,omitempty"`       // Value true to use DNAC certificate for Pxgrid. Default value is false
	IsIseEnabled               *bool                                                                                                    `json:"isIseEnabled,omitempty"`               // Value true for Cisco ISE Server (readonly). Default value is false
	Port                       *int                                                                                                     `json:"port,omitempty"`                       // Port of TACACS server (readonly). The range is from 1 to 65535
	Protocol                   string                                                                                                   `json:"protocol,omitempty"`                   // Type of protocol for authentication and policy server. If already saved with RADIUS, can update to RADIUS_TACACS. If already saved with TACACS, can update to RADIUS_TACACS
	Retries                    string                                                                                                   `json:"retries,omitempty"`                    // Number of communication retries between devices and authentication and policy server. The range is from 1 to 3
	Role                       string                                                                                                   `json:"role,omitempty"`                       // Role of authentication and policy server (readonly). E.g. primary, secondary
	SharedSecret               string                                                                                                   `json:"sharedSecret,omitempty"`               // Shared secret between devices and authentication and policy server (readonly)
	TimeoutSeconds             string                                                                                                   `json:"timeoutSeconds,omitempty"`             // Number of seconds before timing out between devices and authentication and policy server. The range is from 2 to 20
	EncryptionScheme           string                                                                                                   `json:"encryptionScheme,omitempty"`           // Type of encryption scheme for additional security (readonly)
	MessageKey                 string                                                                                                   `json:"messageKey,omitempty"`                 // Message key used to encrypt shared secret (readonly)
	EncryptionKey              string                                                                                                   `json:"encryptionKey,omitempty"`              // Encryption key used to encrypt shared secret (readonly)
	ExternalCiscoIseIPAddrDtos *[]RequestSystemSettingsEditAuthenticationAndPolicyServerAccessConfigurationV1ExternalCiscoIseIPAddrDtos `json:"externalCiscoIseIpAddrDtos,omitempty"` //
}
type RequestSystemSettingsEditAuthenticationAndPolicyServerAccessConfigurationV1CiscoIseDtos struct {
	Description    string `json:"description,omitempty"`    // Description about the Cisco ISE server
	Fqdn           string `json:"fqdn,omitempty"`           // Fully-qualified domain name of the Cisco ISE server (readonly). E.g. xi-62.my.com
	Password       string `json:"password,omitempty"`       // Password of the Cisco ISE server
	SSHkey         string `json:"sshkey,omitempty"`         // SSH key of the Cisco ISE server
	IPAddress      string `json:"ipAddress,omitempty"`      // IP Address of the Cisco ISE Server (readonly)
	SubscriberName string `json:"subscriberName,omitempty"` // Subscriber name of the Cisco ISE server (readonly). E.g. pxgrid_client_1662589467
	UserName       string `json:"userName,omitempty"`       // User name of the Cisco ISE server
}
type RequestSystemSettingsEditAuthenticationAndPolicyServerAccessConfigurationV1ExternalCiscoIseIPAddrDtos struct {
	ExternalCiscoIseIPAddresses *[]RequestSystemSettingsEditAuthenticationAndPolicyServerAccessConfigurationV1ExternalCiscoIseIPAddrDtosExternalCiscoIseIPAddresses `json:"externalCiscoIseIpAddresses,omitempty"` //
	Type                        string                                                                                                                              `json:"type,omitempty"`                        // Type
}
type RequestSystemSettingsEditAuthenticationAndPolicyServerAccessConfigurationV1ExternalCiscoIseIPAddrDtosExternalCiscoIseIPAddresses struct {
	ExternalIPAddress string `json:"externalIpAddress,omitempty"` // External IP Address
}
type RequestSystemSettingsAcceptCiscoIseServerCertificateForCiscoIseServerIntegrationV1 struct {
	IsCertAcceptedByUser *bool `json:"isCertAcceptedByUser,omitempty"` // Value true for accept, false for deny. Remove this field and send empty request payload ( {} ) to retry the failed integration
}
type RequestSystemSettingsCreatesConfigurationDetailsOfTheExternalIPAMServerV1 struct {
	ServerName string `json:"serverName,omitempty"` // A descriptive name of this external server, used for identification purposes
	ServerURL  string `json:"serverUrl,omitempty"`  // The URL of this external server
	Password   string `json:"password,omitempty"`   // The password for the external IPAM server login username
	UserName   string `json:"userName,omitempty"`   // The external IPAM server login username
	Provider   string `json:"provider,omitempty"`   // Type of external IPAM. Can be either INFOBLOX, BLUECAT or GENERIC.
	View       string `json:"view,omitempty"`       // The view under which pools are created in the external IPAM server.
	SyncView   *bool  `json:"syncView,omitempty"`   // Synchronize the IP pools from the local IPAM to this external server
}
type RequestSystemSettingsUpdatesConfigurationDetailsOfTheExternalIPAMServerV1 struct {
	ServerName string `json:"serverName,omitempty"` // A descriptive name of this external server, used for identification purposes
	ServerURL  string `json:"serverUrl,omitempty"`  // The URL of this external server
	Password   string `json:"password,omitempty"`   // The password for the external IPAM server login username
	UserName   string `json:"userName,omitempty"`   // The external IPAM server login username
	View       string `json:"view,omitempty"`       // The view under which pools are created in the external IPAM server.
	SyncView   *bool  `json:"syncView,omitempty"`   // Synchronize the IP pools from the local IPAM to this external server
}
type RequestSystemSettingsCustomPromptPOSTAPIV1 struct {
	UsernamePrompt string `json:"usernamePrompt,omitempty"` // Username for Custom Prompt
	PasswordPrompt string `json:"passwordPrompt,omitempty"` // Password for Custom Prompt
}
type RequestSystemSettingsSetProvisioningSettingsV1 struct {
	RequireItsmApproval *bool `json:"requireItsmApproval,omitempty"` // If require ITSM approval is enabled, the planned configurations must be submitted for ITSM approval. Also if enabled, requirePreview will default to enabled.
	RequirePreview      *bool `json:"requirePreview,omitempty"`      // If require preview is enabled, the device configurations must be reviewed before deploying them
}

//GetAuthenticationAndPolicyServersV1 Get Authentication and Policy Servers - a4b4-c849-4be8-b362
/* API to get Authentication and Policy Servers


@param GetAuthenticationAndPolicyServersV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-authentication-and-policy-servers-v1
*/
func (s *SystemSettingsService) GetAuthenticationAndPolicyServersV1(GetAuthenticationAndPolicyServersV1QueryParams *GetAuthenticationAndPolicyServersV1QueryParams) (*ResponseSystemSettingsGetAuthenticationAndPolicyServersV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/authentication-policy-servers"

	queryString, _ := query.Values(GetAuthenticationAndPolicyServersV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSystemSettingsGetAuthenticationAndPolicyServersV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAuthenticationAndPolicyServersV1(GetAuthenticationAndPolicyServersV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAuthenticationAndPolicyServersV1")
	}

	result := response.Result().(*ResponseSystemSettingsGetAuthenticationAndPolicyServersV1)
	return result, response, err

}

//RetrievesConfigurationDetailsOfTheExternalIPAMServerV1 Retrieves configuration details of the external IPAM server. - 3ebf-1bc3-4c8a-95e4
/* Retrieves configuration details of the external IPAM server.  If an external IPAM server has not been created, this resource will return a `404` response.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-configuration-details-of-the-external-ip-a-m-server-v1
*/
func (s *SystemSettingsService) RetrievesConfigurationDetailsOfTheExternalIPAMServerV1() (*ResponseSystemSettingsRetrievesConfigurationDetailsOfTheExternalIPAMServerV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/ipam/serverSetting"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSystemSettingsRetrievesConfigurationDetailsOfTheExternalIPAMServerV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesConfigurationDetailsOfTheExternalIPAMServerV1()
		}
		return nil, response, fmt.Errorf("error with operation RetrievesConfigurationDetailsOfTheExternalIpAMServerV1")
	}

	result := response.Result().(*ResponseSystemSettingsRetrievesConfigurationDetailsOfTheExternalIPAMServerV1)
	return result, response, err

}

//CiscoIseServerIntegrationStatusV1 Cisco ISE Server Integration Status - c1a4-f8fb-448a-8135
/* API to check Cisco ISE server integration status.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!cisco-ise-server-integration-status-v1
*/
func (s *SystemSettingsService) CiscoIseServerIntegrationStatusV1() (*ResponseSystemSettingsCiscoIseServerIntegrationStatusV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/ise-integration-status"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSystemSettingsCiscoIseServerIntegrationStatusV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.CiscoIseServerIntegrationStatusV1()
		}
		return nil, response, fmt.Errorf("error with operation CiscoIseServerIntegrationStatusV1")
	}

	result := response.Result().(*ResponseSystemSettingsCiscoIseServerIntegrationStatusV1)
	return result, response, err

}

//CustomPromptSupportGETAPIV1 Custom-prompt support GET API - 2aa8-f90e-4ebb-a629
/* Returns supported custom prompts by Catalyst Center



Documentation Link: https://developer.cisco.com/docs/dna-center/#!custom-prompt-support-g-e-t-api-v1
*/
func (s *SystemSettingsService) CustomPromptSupportGETAPIV1() (*ResponseSystemSettingsCustomPromptSupportGETAPIV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/custom-prompt"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSystemSettingsCustomPromptSupportGETAPIV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.CustomPromptSupportGETAPIV1()
		}
		return nil, response, fmt.Errorf("error with operation CustomPromptSupportGETApiV1")
	}

	result := response.Result().(*ResponseSystemSettingsCustomPromptSupportGETAPIV1)
	return result, response, err

}

//GetProvisioningSettingsV1 Get provisioning settings - b9b5-db54-4788-82e4
/* Returns provisioning settings



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-provisioning-settings-v1
*/
func (s *SystemSettingsService) GetProvisioningSettingsV1() (*ResponseSystemSettingsGetProvisioningSettingsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/provisioningSettings"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSystemSettingsGetProvisioningSettingsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetProvisioningSettingsV1()
		}
		return nil, response, fmt.Errorf("error with operation GetProvisioningSettingsV1")
	}

	result := response.Result().(*ResponseSystemSettingsGetProvisioningSettingsV1)
	return result, response, err

}

//AddAuthenticationAndPolicyServerAccessConfigurationV1 Add Authentication and Policy Server Access Configuration - 5282-78a3-4fbb-a82c
/* API to add AAA/ISE server access configuration. Protocol can be configured as either RADIUS OR TACACS OR RADIUS_TACACS. If configuring Cisco ISE server, after configuration, use ‘Cisco ISE Server Integration Status’ Intent API to check the integration status. Based on integration status, if require use 'Accept Cisco ISE Server Certificate for Cisco ISE Server Integration' Intent API to accept the Cisco ISE certificate for Cisco ISE server integration, then use again ‘Cisco ISE Server Integration Status’ Intent API to check the integration status.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-authentication-and-policy-server-access-configuration-v1
*/
func (s *SystemSettingsService) AddAuthenticationAndPolicyServerAccessConfigurationV1(requestSystemSettingsAddAuthenticationAndPolicyServerAccessConfigurationV1 *RequestSystemSettingsAddAuthenticationAndPolicyServerAccessConfigurationV1) (*ResponseSystemSettingsAddAuthenticationAndPolicyServerAccessConfigurationV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/authentication-policy-servers"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSystemSettingsAddAuthenticationAndPolicyServerAccessConfigurationV1).
		SetResult(&ResponseSystemSettingsAddAuthenticationAndPolicyServerAccessConfigurationV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddAuthenticationAndPolicyServerAccessConfigurationV1(requestSystemSettingsAddAuthenticationAndPolicyServerAccessConfigurationV1)
		}

		return nil, response, fmt.Errorf("error with operation AddAuthenticationAndPolicyServerAccessConfigurationV1")
	}

	result := response.Result().(*ResponseSystemSettingsAddAuthenticationAndPolicyServerAccessConfigurationV1)
	return result, response, err

}

//CreatesConfigurationDetailsOfTheExternalIPAMServerV1 Creates configuration details of the external IPAM server. - 36a4-38d6-4589-8f87
/* Creates configuration details of the external IPAM server. You should only create one external IPAM server; delete any existing external server before creating a new one.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!creates-configuration-details-of-the-external-ip-a-m-server-v1
*/
func (s *SystemSettingsService) CreatesConfigurationDetailsOfTheExternalIPAMServerV1(requestSystemSettingsCreatesConfigurationDetailsOfTheExternalIPAMServerV1 *RequestSystemSettingsCreatesConfigurationDetailsOfTheExternalIPAMServerV1) (*ResponseSystemSettingsCreatesConfigurationDetailsOfTheExternalIPAMServerV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/ipam/serverSetting"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSystemSettingsCreatesConfigurationDetailsOfTheExternalIPAMServerV1).
		SetResult(&ResponseSystemSettingsCreatesConfigurationDetailsOfTheExternalIPAMServerV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreatesConfigurationDetailsOfTheExternalIPAMServerV1(requestSystemSettingsCreatesConfigurationDetailsOfTheExternalIPAMServerV1)
		}

		return nil, response, fmt.Errorf("error with operation CreatesConfigurationDetailsOfTheExternalIpAMServerV1")
	}

	result := response.Result().(*ResponseSystemSettingsCreatesConfigurationDetailsOfTheExternalIPAMServerV1)
	return result, response, err

}

//CustomPromptPOSTAPIV1 Custom Prompt POST API - f4b9-1a8a-4718-aa97
/* Save custom prompt added by user in Catalyst Center. API will always override the existing prompts. User should provide all the custom prompt in case of any update



Documentation Link: https://developer.cisco.com/docs/dna-center/#!custom-prompt-p-o-s-t-api-v1
*/
func (s *SystemSettingsService) CustomPromptPOSTAPIV1(requestSystemSettingsCustomPromptPOSTAPIV1 *RequestSystemSettingsCustomPromptPOSTAPIV1) (*ResponseSystemSettingsCustomPromptPOSTAPIV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/custom-prompt"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSystemSettingsCustomPromptPOSTAPIV1).
		SetResult(&ResponseSystemSettingsCustomPromptPOSTAPIV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CustomPromptPOSTAPIV1(requestSystemSettingsCustomPromptPOSTAPIV1)
		}

		return nil, response, fmt.Errorf("error with operation CustomPromptPOSTApiV1")
	}

	result := response.Result().(*ResponseSystemSettingsCustomPromptPOSTAPIV1)
	return result, response, err

}

//EditAuthenticationAndPolicyServerAccessConfigurationV1 Edit Authentication and Policy Server Access Configuration - e4bf-3bf1-48c9-ba11
/* API to edit AAA/ISE server access configuration. After edit, use ‘Cisco ISE Server Integration Status’ Intent API to check the integration status.


@param id id path parameter. Authentication and Policy Server Identifier. Use 'Get Authentication and Policy Servers' intent API to find the identifier.

*/
func (s *SystemSettingsService) EditAuthenticationAndPolicyServerAccessConfigurationV1(id string, requestSystemSettingsEditAuthenticationAndPolicyServerAccessConfigurationV1 *RequestSystemSettingsEditAuthenticationAndPolicyServerAccessConfigurationV1) (*ResponseSystemSettingsEditAuthenticationAndPolicyServerAccessConfigurationV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/authentication-policy-servers/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSystemSettingsEditAuthenticationAndPolicyServerAccessConfigurationV1).
		SetResult(&ResponseSystemSettingsEditAuthenticationAndPolicyServerAccessConfigurationV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.EditAuthenticationAndPolicyServerAccessConfigurationV1(id, requestSystemSettingsEditAuthenticationAndPolicyServerAccessConfigurationV1)
		}
		return nil, response, fmt.Errorf("error with operation EditAuthenticationAndPolicyServerAccessConfigurationV1")
	}

	result := response.Result().(*ResponseSystemSettingsEditAuthenticationAndPolicyServerAccessConfigurationV1)
	return result, response, err

}

//AcceptCiscoIseServerCertificateForCiscoIseServerIntegrationV1 Accept Cisco ISE Server Certificate for Cisco ISE Server Integration - c8a6-aae2-48b8-b41c
/* API to accept Cisco ISE server certificate for Cisco ISE server integration. Use ‘Cisco ISE Server Integration Status’ Intent API to check the integration status. This API can be used to retry the failed integration.


@param id id path parameter. Cisco ISE Server Identifier. Use 'Get Authentication and Policy Servers' intent API to find the identifier.

*/
func (s *SystemSettingsService) AcceptCiscoIseServerCertificateForCiscoIseServerIntegrationV1(id string, requestSystemSettingsAcceptCiscoISEServerCertificateForCiscoISEServerIntegrationV1 *RequestSystemSettingsAcceptCiscoIseServerCertificateForCiscoIseServerIntegrationV1) (*resty.Response, error) {
	path := "/dna/intent/api/v1/integrate-ise/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSystemSettingsAcceptCiscoISEServerCertificateForCiscoISEServerIntegrationV1).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.AcceptCiscoIseServerCertificateForCiscoIseServerIntegrationV1(id, requestSystemSettingsAcceptCiscoISEServerCertificateForCiscoISEServerIntegrationV1)
		}
		return response, fmt.Errorf("error with operation AcceptCiscoIseServerCertificateForCiscoIseServerIntegrationV1")
	}

	return response, err

}

//UpdatesConfigurationDetailsOfTheExternalIPAMServerV1 Updates configuration details of the external IPAM server. - 5a99-0bfe-4c99-a0a4
/* Updates configuration details of the external IPAM server.


 */
func (s *SystemSettingsService) UpdatesConfigurationDetailsOfTheExternalIPAMServerV1(requestSystemSettingsUpdatesConfigurationDetailsOfTheExternalIPAMServerV1 *RequestSystemSettingsUpdatesConfigurationDetailsOfTheExternalIPAMServerV1) (*ResponseSystemSettingsUpdatesConfigurationDetailsOfTheExternalIPAMServerV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/ipam/serverSetting"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSystemSettingsUpdatesConfigurationDetailsOfTheExternalIPAMServerV1).
		SetResult(&ResponseSystemSettingsUpdatesConfigurationDetailsOfTheExternalIPAMServerV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdatesConfigurationDetailsOfTheExternalIPAMServerV1(requestSystemSettingsUpdatesConfigurationDetailsOfTheExternalIPAMServerV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdatesConfigurationDetailsOfTheExternalIpAMServerV1")
	}

	result := response.Result().(*ResponseSystemSettingsUpdatesConfigurationDetailsOfTheExternalIPAMServerV1)
	return result, response, err

}

//SetProvisioningSettingsV1 Set provisioning settings - e5ab-1811-450a-bb01
/* Sets provisioning settings


 */
func (s *SystemSettingsService) SetProvisioningSettingsV1(requestSystemSettingsSetProvisioningSettingsV1 *RequestSystemSettingsSetProvisioningSettingsV1) (*ResponseSystemSettingsSetProvisioningSettingsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/provisioningSettings"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSystemSettingsSetProvisioningSettingsV1).
		SetResult(&ResponseSystemSettingsSetProvisioningSettingsV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SetProvisioningSettingsV1(requestSystemSettingsSetProvisioningSettingsV1)
		}
		return nil, response, fmt.Errorf("error with operation SetProvisioningSettingsV1")
	}

	result := response.Result().(*ResponseSystemSettingsSetProvisioningSettingsV1)
	return result, response, err

}

//DeleteAuthenticationAndPolicyServerAccessConfigurationV1 Delete Authentication and Policy Server Access Configuration - 0b92-bb8a-477a-a942
/* API to delete AAA/ISE server access configuration.


@param id id path parameter. Authentication and Policy Server Identifier. Use 'Get Authentication and Policy Servers' intent API to find the identifier.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-authentication-and-policy-server-access-configuration-v1
*/
func (s *SystemSettingsService) DeleteAuthenticationAndPolicyServerAccessConfigurationV1(id string) (*ResponseSystemSettingsDeleteAuthenticationAndPolicyServerAccessConfigurationV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/authentication-policy-servers/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSystemSettingsDeleteAuthenticationAndPolicyServerAccessConfigurationV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteAuthenticationAndPolicyServerAccessConfigurationV1(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteAuthenticationAndPolicyServerAccessConfigurationV1")
	}

	result := response.Result().(*ResponseSystemSettingsDeleteAuthenticationAndPolicyServerAccessConfigurationV1)
	return result, response, err

}

//DeletesConfigurationDetailsOfTheExternalIPAMServerV1 Deletes configuration details of the external IPAM server. - 67b6-eb01-4688-a164
/* Deletes configuration details of the external IPAM server.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!deletes-configuration-details-of-the-external-ip-a-m-server-v1
*/
func (s *SystemSettingsService) DeletesConfigurationDetailsOfTheExternalIPAMServerV1() (*ResponseSystemSettingsDeletesConfigurationDetailsOfTheExternalIPAMServerV1, *resty.Response, error) {
	//
	path := "/dna/intent/api/v1/ipam/serverSetting"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSystemSettingsDeletesConfigurationDetailsOfTheExternalIPAMServerV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletesConfigurationDetailsOfTheExternalIPAMServerV1()
		}
		return nil, response, fmt.Errorf("error with operation DeletesConfigurationDetailsOfTheExternalIpAMServerV1")
	}

	result := response.Result().(*ResponseSystemSettingsDeletesConfigurationDetailsOfTheExternalIPAMServerV1)
	return result, response, err

}

// Alias Function
/*
This method acts as an alias for the method `CreatesConfigurationDetailsOfTheExternalIPAMServerV1`
*/
func (s *SystemSettingsService) CreatesConfigurationDetailsOfTheExternalIPAMServer(requestSystemSettingsCreatesConfigurationDetailsOfTheExternalIPAMServerV1 *RequestSystemSettingsCreatesConfigurationDetailsOfTheExternalIPAMServerV1) (*ResponseSystemSettingsCreatesConfigurationDetailsOfTheExternalIPAMServerV1, *resty.Response, error) {
	return s.CreatesConfigurationDetailsOfTheExternalIPAMServerV1(requestSystemSettingsCreatesConfigurationDetailsOfTheExternalIPAMServerV1)
}

// Alias Function
/*
This method acts as an alias for the method `DeleteAuthenticationAndPolicyServerAccessConfigurationV1`
*/
func (s *SystemSettingsService) DeleteAuthenticationAndPolicyServerAccessConfiguration(id string) (*ResponseSystemSettingsDeleteAuthenticationAndPolicyServerAccessConfigurationV1, *resty.Response, error) {
	return s.DeleteAuthenticationAndPolicyServerAccessConfigurationV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `SetProvisioningSettingsV1`
*/
func (s *SystemSettingsService) SetProvisioningSettings(requestSystemSettingsSetProvisioningSettingsV1 *RequestSystemSettingsSetProvisioningSettingsV1) (*ResponseSystemSettingsSetProvisioningSettingsV1, *resty.Response, error) {
	return s.SetProvisioningSettingsV1(requestSystemSettingsSetProvisioningSettingsV1)
}

// Alias Function
/*
This method acts as an alias for the method `CustomPromptSupportGETAPIV1`
*/
func (s *SystemSettingsService) CustomPromptSupportGETAPI() (*ResponseSystemSettingsCustomPromptSupportGETAPIV1, *resty.Response, error) {
	return s.CustomPromptSupportGETAPIV1()
}

// Alias Function
/*
This method acts as an alias for the method `AddAuthenticationAndPolicyServerAccessConfigurationV1`
*/
func (s *SystemSettingsService) AddAuthenticationAndPolicyServerAccessConfiguration(requestSystemSettingsAddAuthenticationAndPolicyServerAccessConfigurationV1 *RequestSystemSettingsAddAuthenticationAndPolicyServerAccessConfigurationV1) (*ResponseSystemSettingsAddAuthenticationAndPolicyServerAccessConfigurationV1, *resty.Response, error) {
	return s.AddAuthenticationAndPolicyServerAccessConfigurationV1(requestSystemSettingsAddAuthenticationAndPolicyServerAccessConfigurationV1)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesConfigurationDetailsOfTheExternalIPAMServerV1`
*/
func (s *SystemSettingsService) RetrievesConfigurationDetailsOfTheExternalIPAMServer() (*ResponseSystemSettingsRetrievesConfigurationDetailsOfTheExternalIPAMServerV1, *resty.Response, error) {
	return s.RetrievesConfigurationDetailsOfTheExternalIPAMServerV1()
}

// Alias Function
/*
This method acts as an alias for the method `CustomPromptPOSTAPIV1`
*/
func (s *SystemSettingsService) CustomPromptPOSTAPI(requestSystemSettingsCustomPromptPOSTAPIV1 *RequestSystemSettingsCustomPromptPOSTAPIV1) (*ResponseSystemSettingsCustomPromptPOSTAPIV1, *resty.Response, error) {
	return s.CustomPromptPOSTAPIV1(requestSystemSettingsCustomPromptPOSTAPIV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetProvisioningSettingsV1`
*/
func (s *SystemSettingsService) GetProvisioningSettings() (*ResponseSystemSettingsGetProvisioningSettingsV1, *resty.Response, error) {
	return s.GetProvisioningSettingsV1()
}

// Alias Function
/*
This method acts as an alias for the method `AcceptCiscoIseServerCertificateForCiscoIseServerIntegrationV1`
*/
func (s *SystemSettingsService) AcceptCiscoIseServerCertificateForCiscoIseServerIntegration(id string, requestSystemSettingsAcceptCiscoISEServerCertificateForCiscoISEServerIntegrationV1 *RequestSystemSettingsAcceptCiscoIseServerCertificateForCiscoIseServerIntegrationV1) (*resty.Response, error) {
	return s.AcceptCiscoIseServerCertificateForCiscoIseServerIntegrationV1(id, requestSystemSettingsAcceptCiscoISEServerCertificateForCiscoISEServerIntegrationV1)
}

// Alias Function
/*
This method acts as an alias for the method `DeletesConfigurationDetailsOfTheExternalIPAMServerV1`
*/
func (s *SystemSettingsService) DeletesConfigurationDetailsOfTheExternalIPAMServer() (*ResponseSystemSettingsDeletesConfigurationDetailsOfTheExternalIPAMServerV1, *resty.Response, error) {
	return s.DeletesConfigurationDetailsOfTheExternalIPAMServerV1()
}

// Alias Function
/*
This method acts as an alias for the method `CiscoIseServerIntegrationStatusV1`
*/
func (s *SystemSettingsService) CiscoIseServerIntegrationStatus() (*ResponseSystemSettingsCiscoIseServerIntegrationStatusV1, *resty.Response, error) {
	return s.CiscoIseServerIntegrationStatusV1()
}

// Alias Function
/*
This method acts as an alias for the method `EditAuthenticationAndPolicyServerAccessConfigurationV1`
*/
func (s *SystemSettingsService) EditAuthenticationAndPolicyServerAccessConfiguration(id string, requestSystemSettingsEditAuthenticationAndPolicyServerAccessConfigurationV1 *RequestSystemSettingsEditAuthenticationAndPolicyServerAccessConfigurationV1) (*ResponseSystemSettingsEditAuthenticationAndPolicyServerAccessConfigurationV1, *resty.Response, error) {
	return s.EditAuthenticationAndPolicyServerAccessConfigurationV1(id, requestSystemSettingsEditAuthenticationAndPolicyServerAccessConfigurationV1)
}

// Alias Function
/*
This method acts as an alias for the method `UpdatesConfigurationDetailsOfTheExternalIPAMServerV1`
*/
func (s *SystemSettingsService) UpdatesConfigurationDetailsOfTheExternalIPAMServer(requestSystemSettingsUpdatesConfigurationDetailsOfTheExternalIPAMServerV1 *RequestSystemSettingsUpdatesConfigurationDetailsOfTheExternalIPAMServerV1) (*ResponseSystemSettingsUpdatesConfigurationDetailsOfTheExternalIPAMServerV1, *resty.Response, error) {
	return s.UpdatesConfigurationDetailsOfTheExternalIPAMServerV1(requestSystemSettingsUpdatesConfigurationDetailsOfTheExternalIPAMServerV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetAuthenticationAndPolicyServersV1`
*/
func (s *SystemSettingsService) GetAuthenticationAndPolicyServers(GetAuthenticationAndPolicyServersV1QueryParams *GetAuthenticationAndPolicyServersV1QueryParams) (*ResponseSystemSettingsGetAuthenticationAndPolicyServersV1, *resty.Response, error) {
	return s.GetAuthenticationAndPolicyServersV1(GetAuthenticationAndPolicyServersV1QueryParams)
}
