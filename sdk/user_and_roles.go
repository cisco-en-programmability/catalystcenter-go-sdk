package catalyst

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type UserandRolesService service

type GetRolesAPIV1HeaderParams struct {
	InvokeSource string `url:"invokeSource,omitempty"` //Expects type string. The source that invokes this API. The value of this header must be set to "external".
}
type GetUsersAPIV1QueryParams struct {
	InvokeSource string `url:"invokeSource,omitempty"` //The source that invokes this API. The value of this query parameter must be set to "external".
	AuthSource   string `url:"authSource,omitempty"`   //The source that authenticates the user. The value of this query parameter can be set to "internal" or "external". If not provided, then all users will be returned in the response.
}
type GetExternalAuthenticationServersAPIV1QueryParams struct {
	InvokeSource string `url:"invokeSource,omitempty"` //The source that invokes this API. The value of this query parameter must be set to "external".
}

type ResponseUserandRolesAddRoleAPIV1 struct {
	Response *ResponseUserandRolesAddRoleAPIV1Response `json:"response,omitempty"` //
}
type ResponseUserandRolesAddRoleAPIV1Response struct {
	RoleID  string `json:"roleId,omitempty"`  // Role Id
	Message string `json:"message,omitempty"` // Message
}
type ResponseUserandRolesUpdateRoleAPIV1 struct {
	Response *ResponseUserandRolesUpdateRoleAPIV1Response `json:"response,omitempty"` //
}
type ResponseUserandRolesUpdateRoleAPIV1Response struct {
	RoleID  string `json:"roleId,omitempty"`  // Role Id
	Message string `json:"message,omitempty"` // Message
}
type ResponseUserandRolesGetPermissionsAPIV1 struct {
	Response *ResponseUserandRolesGetPermissionsAPIV1Response `json:"response,omitempty"` //
}
type ResponseUserandRolesGetPermissionsAPIV1Response struct {
	ResourceTypes *[]ResponseUserandRolesGetPermissionsAPIV1ResponseResourceTypes `json:"resource-types,omitempty"` //
}
type ResponseUserandRolesGetPermissionsAPIV1ResponseResourceTypes struct {
	Type              string `json:"type,omitempty"`              // Type
	DisplayName       string `json:"displayName,omitempty"`       // Display Name
	Description       string `json:"description,omitempty"`       // Description
	DefaultPermission string `json:"defaultPermission,omitempty"` // Default permission
}
type ResponseUserandRolesDeleteRoleAPIV1 struct {
	Response *ResponseUserandRolesDeleteRoleAPIV1Response `json:"response,omitempty"` //
}
type ResponseUserandRolesDeleteRoleAPIV1Response struct {
	Message string `json:"message,omitempty"` // Message
}
type ResponseUserandRolesGetRolesAPIV1 struct {
	Response *ResponseUserandRolesGetRolesAPIV1Response `json:"response,omitempty"` //
}
type ResponseUserandRolesGetRolesAPIV1Response struct {
	Roles *[]ResponseUserandRolesGetRolesAPIV1ResponseRoles `json:"roles,omitempty"` //
}
type ResponseUserandRolesGetRolesAPIV1ResponseRoles struct {
	ResourceTypes *[]ResponseUserandRolesGetRolesAPIV1ResponseRolesResourceTypes `json:"resourceTypes,omitempty"` //
	Meta          *ResponseUserandRolesGetRolesAPIV1ResponseRolesMeta            `json:"meta,omitempty"`          //
	RoleID        string                                                         `json:"roleId,omitempty"`        // Role Id
	Name          string                                                         `json:"name,omitempty"`          // Role name
	Description   string                                                         `json:"description,omitempty"`   // Description
	Type          string                                                         `json:"type,omitempty"`          // Role type, possible values are: "DEFAULT", "SYSTEM", "CUSTOM"
}
type ResponseUserandRolesGetRolesAPIV1ResponseRolesResourceTypes struct {
	Operations []string `json:"operations,omitempty"` // Operations
	Type       string   `json:"type,omitempty"`       // Type
}
type ResponseUserandRolesGetRolesAPIV1ResponseRolesMeta struct {
	CreatedBy    string `json:"createdBy,omitempty"`    // The user that creates the resource type
	Created      string `json:"created,omitempty"`      // The timestamp that the resource type was created
	LastModified string `json:"lastModified,omitempty"` // The latestest timestamp that the resource type was updated
}
type ResponseUserandRolesGetUsersAPIV1 struct {
	Response *ResponseUserandRolesGetUsersAPIV1Response `json:"response,omitempty"` //
}
type ResponseUserandRolesGetUsersAPIV1Response struct {
	Users *[]ResponseUserandRolesGetUsersAPIV1ResponseUsers `json:"users,omitempty"` //
}
type ResponseUserandRolesGetUsersAPIV1ResponseUsers struct {
	FirstName            string   `json:"firstName,omitempty"`            // First Name
	LastName             string   `json:"lastName,omitempty"`             // Last Name
	AuthSource           string   `json:"authSource,omitempty"`           // Authentiction source, internal or external
	PassphraseUpdateTime string   `json:"passphraseUpdateTime,omitempty"` // Passphrase Update Time
	RoleList             []string `json:"roleList,omitempty"`             // A list of role ids
	UserID               string   `json:"userId,omitempty"`               // User Id
	Email                string   `json:"email,omitempty"`                // Email
	Username             string   `json:"username,omitempty"`             // Username
}
type ResponseUserandRolesAddUserAPIV1 struct {
	Response *ResponseUserandRolesAddUserAPIV1Response `json:"response,omitempty"` //
}
type ResponseUserandRolesAddUserAPIV1Response struct {
	Message string `json:"message,omitempty"` // Message
	UserID  string `json:"userId,omitempty"`  // User Id
}
type ResponseUserandRolesUpdateUserAPIV1 struct {
	Response *ResponseUserandRolesUpdateUserAPIV1Response `json:"response,omitempty"` //
}
type ResponseUserandRolesUpdateUserAPIV1Response struct {
	Message string `json:"message,omitempty"` // Message
}
type ResponseUserandRolesDeleteUserAPIV1 struct {
	Response *ResponseUserandRolesDeleteUserAPIV1Response `json:"response,omitempty"` //
}
type ResponseUserandRolesDeleteUserAPIV1Response struct {
	Message string `json:"message,omitempty"` // Message
}
type ResponseUserandRolesGetExternalAuthenticationSettingAPIV1 struct {
	Response *ResponseUserandRolesGetExternalAuthenticationSettingAPIV1Response `json:"response,omitempty"` //
}
type ResponseUserandRolesGetExternalAuthenticationSettingAPIV1Response struct {
	ExternalAuthenticationFlag *[]ResponseUserandRolesGetExternalAuthenticationSettingAPIV1ResponseExternalAuthenticationFlag `json:"external-authentication-flag,omitempty"` //
}
type ResponseUserandRolesGetExternalAuthenticationSettingAPIV1ResponseExternalAuthenticationFlag struct {
	Enabled *bool `json:"enabled,omitempty"` // External Authentication is enabled/disabled.
}
type ResponseUserandRolesManageExternalAuthenticationSettingAPIV1 struct {
	Response *ResponseUserandRolesManageExternalAuthenticationSettingAPIV1Response `json:"response,omitempty"` //
}
type ResponseUserandRolesManageExternalAuthenticationSettingAPIV1Response struct {
	Message string `json:"message,omitempty"` // Message
}
type ResponseUserandRolesGetExternalAuthenticationServersAPIV1 struct {
	Response *ResponseUserandRolesGetExternalAuthenticationServersAPIV1Response `json:"response,omitempty"` //
}
type ResponseUserandRolesGetExternalAuthenticationServersAPIV1Response struct {
	AAAServers *[]ResponseUserandRolesGetExternalAuthenticationServersAPIV1ResponseAAAServers `json:"aaa-servers,omitempty"` //
}
type ResponseUserandRolesGetExternalAuthenticationServersAPIV1ResponseAAAServers struct {
	AccountingPort     *int   `json:"accountingPort,omitempty"`     // RADIUS server accounting requests port
	Retries            *int   `json:"retries,omitempty"`            // Retries
	Protocol           string `json:"protocol,omitempty"`           // Protocol
	SocketTimeout      *int   `json:"socketTimeout,omitempty"`      // Timeout in seconds
	ServerIP           string `json:"serverIp,omitempty"`           // Server Ip
	SharedSecret       string `json:"sharedSecret,omitempty"`       // Shared Secret
	ServerID           string `json:"serverId,omitempty"`           // Server Id
	AuthenticationPort *int   `json:"authenticationPort,omitempty"` // RADIUS server authorization requests port
	AAAAttribute       string `json:"aaaAttribute,omitempty"`       // Aaa Attribute
	Role               string `json:"role,omitempty"`               // Role of AAA server, primary or secondary server
}
type ResponseUserandRolesAddAndUpdateAAAAttributeAPIV1 struct {
	Response *ResponseUserandRolesAddAndUpdateAAAAttributeAPIV1Response `json:"response,omitempty"` //
}
type ResponseUserandRolesAddAndUpdateAAAAttributeAPIV1Response struct {
	Message string `json:"message,omitempty"` // Message
}
type ResponseUserandRolesDeleteAAAAttributeAPIV1 struct {
	Response *ResponseUserandRolesDeleteAAAAttributeAPIV1Response `json:"response,omitempty"` //
}
type ResponseUserandRolesDeleteAAAAttributeAPIV1Response struct {
	Message string `json:"message,omitempty"` // Message
}
type ResponseUserandRolesGetAAAAttributeAPIV1 struct {
	Response *ResponseUserandRolesGetAAAAttributeAPIV1Response `json:"response,omitempty"` //
}
type ResponseUserandRolesGetAAAAttributeAPIV1Response struct {
	AAAAttributes *[]ResponseUserandRolesGetAAAAttributeAPIV1ResponseAAAAttributes `json:"aaa-attributes,omitempty"` //
}
type ResponseUserandRolesGetAAAAttributeAPIV1ResponseAAAAttributes struct {
	AttributeName string `json:"attributeName,omitempty"` // Value of the custom AAA attribute name
}
type RequestUserandRolesAddRoleAPIV1 struct {
	Role          string                                          `json:"role,omitempty"`          // Name of the role
	Description   string                                          `json:"description,omitempty"`   // Description of role
	ResourceTypes *[]RequestUserandRolesAddRoleAPIV1ResourceTypes `json:"resourceTypes,omitempty"` //
}
type RequestUserandRolesAddRoleAPIV1ResourceTypes struct {
	Type       string   `json:"type,omitempty"`       // Name of the application in Cisco DNA Center System
	Operations []string `json:"operations,omitempty"` // List of operations allowed for the application. Possible values are "gRead", "gCreate", "gUpdate", "gRemove", or some combination of these.
}
type RequestUserandRolesUpdateRoleAPIV1 struct {
	RoleID        string                                             `json:"roleId,omitempty"`        // Id of the role
	Description   string                                             `json:"description,omitempty"`   // Description of the role
	ResourceTypes *[]RequestUserandRolesUpdateRoleAPIV1ResourceTypes `json:"resourceTypes,omitempty"` //
}
type RequestUserandRolesUpdateRoleAPIV1ResourceTypes struct {
	Type       string   `json:"type,omitempty"`       // Name of application in Cisco DNA Center System
	Operations []string `json:"operations,omitempty"` // List of operations allowed for the application. Possible values are "gRead", "gCreate", "gUpdate", "gRemove", or some combination of these.
}
type RequestUserandRolesAddUserAPIV1 struct {
	FirstName string   `json:"firstName,omitempty"` // First Name
	LastName  string   `json:"lastName,omitempty"`  // Last Name
	Username  string   `json:"username,omitempty"`  // Username
	Password  string   `json:"password,omitempty"`  // Password
	Email     string   `json:"email,omitempty"`     // Email
	RoleList  []string `json:"roleList,omitempty"`  // Role id list
}
type RequestUserandRolesUpdateUserAPIV1 struct {
	FirstName string   `json:"firstName,omitempty"` // firstName should be set if the original value is not empty
	LastName  string   `json:"lastName,omitempty"`  // lastName should be set if the original value is not empty
	Email     string   `json:"email,omitempty"`     // email should be set if the original value is not empty
	Username  string   `json:"username,omitempty"`  // Username
	UserID    string   `json:"userId,omitempty"`    // User Id
	RoleList  []string `json:"roleList,omitempty"`  // Role id list
}
type RequestUserandRolesManageExternalAuthenticationSettingAPIV1 struct {
	Enable *bool `json:"enable,omitempty"` // Enable/disable External Authentication.
}
type RequestUserandRolesAddAndUpdateAAAAttributeAPIV1 struct {
	AttributeName string `json:"attributeName,omitempty"` // name of the custom AAA attribute.
}

//GetPermissionsAPIV1 Get permissions API - 8a9c-6885-455b-a2db
/* Get permissions for a role from Cisco DNA Center System.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-permissions-api-v1
*/
func (s *UserandRolesService) GetPermissionsAPIV1() (*ResponseUserandRolesGetPermissionsAPIV1, *resty.Response, error) {
	path := "/dna/system/api/v1/role/permissions"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseUserandRolesGetPermissionsAPIV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPermissionsAPIV1()
		}
		return nil, response, fmt.Errorf("error with operation GetPermissionsApiV1")
	}

	result := response.Result().(*ResponseUserandRolesGetPermissionsAPIV1)
	return result, response, err

}

//GetRolesAPIV1 Get roles API - 7c86-da3f-4b08-8593
/* Get all roles for the Cisco DNA Center System.


@param GetRolesAPIV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-roles-api-v1
*/
func (s *UserandRolesService) GetRolesAPIV1(GetRolesAPIV1HeaderParams *GetRolesAPIV1HeaderParams) (*ResponseUserandRolesGetRolesAPIV1, *resty.Response, error) {
	path := "/dna/system/api/v1/roles"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetRolesAPIV1HeaderParams != nil {

		if GetRolesAPIV1HeaderParams.InvokeSource != "" {
			clientRequest = clientRequest.SetHeader("invokeSource", GetRolesAPIV1HeaderParams.InvokeSource)
		}

	}

	response, err = clientRequest.
		SetResult(&ResponseUserandRolesGetRolesAPIV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetRolesAPIV1(GetRolesAPIV1HeaderParams)
		}
		return nil, response, fmt.Errorf("error with operation GetRolesApiV1")
	}

	result := response.Result().(*ResponseUserandRolesGetRolesAPIV1)
	return result, response, err

}

//GetUsersAPIV1 Get users API - 918c-89fa-4a98-a528
/* Get all users for the Cisco DNA Center System.


@param GetUsersAPIV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-users-api-v1
*/
func (s *UserandRolesService) GetUsersAPIV1(GetUsersAPIV1QueryParams *GetUsersAPIV1QueryParams) (*ResponseUserandRolesGetUsersAPIV1, *resty.Response, error) {
	path := "/dna/system/api/v1/user"

	queryString, _ := query.Values(GetUsersAPIV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseUserandRolesGetUsersAPIV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetUsersAPIV1(GetUsersAPIV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetUsersApiV1")
	}

	result := response.Result().(*ResponseUserandRolesGetUsersAPIV1)
	return result, response, err

}

//GetExternalAuthenticationSettingAPIV1 Get External Authentication Setting API - e0a8-aa75-49cb-815c
/* Get the External Authentication setting.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-external-authentication-setting-api-v1
*/
func (s *UserandRolesService) GetExternalAuthenticationSettingAPIV1() (*ResponseUserandRolesGetExternalAuthenticationSettingAPIV1, *resty.Response, error) {
	path := "/dna/system/api/v1/users/external-authentication"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseUserandRolesGetExternalAuthenticationSettingAPIV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetExternalAuthenticationSettingAPIV1()
		}
		return nil, response, fmt.Errorf("error with operation GetExternalAuthenticationSettingApiV1")
	}

	result := response.Result().(*ResponseUserandRolesGetExternalAuthenticationSettingAPIV1)
	return result, response, err

}

//GetExternalAuthenticationServersAPIV1 Get external authentication servers API - 9dbd-0b01-4758-bde4
/* Get external users authentication servers.


@param GetExternalAuthenticationServersAPIV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-external-authentication-servers-api-v1
*/
func (s *UserandRolesService) GetExternalAuthenticationServersAPIV1(GetExternalAuthenticationServersAPIV1QueryParams *GetExternalAuthenticationServersAPIV1QueryParams) (*ResponseUserandRolesGetExternalAuthenticationServersAPIV1, *resty.Response, error) {
	path := "/dna/system/api/v1/users/external-servers"

	queryString, _ := query.Values(GetExternalAuthenticationServersAPIV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseUserandRolesGetExternalAuthenticationServersAPIV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetExternalAuthenticationServersAPIV1(GetExternalAuthenticationServersAPIV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetExternalAuthenticationServersApiV1")
	}

	result := response.Result().(*ResponseUserandRolesGetExternalAuthenticationServersAPIV1)
	return result, response, err

}

//GetAAAAttributeAPIV1 Get AAA Attribute API - 2eb5-ea84-4d29-bf8b
/* Get the current value of the custom AAA attribute.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-a-a-a-attribute-api-v1
*/
func (s *UserandRolesService) GetAAAAttributeAPIV1() (*ResponseUserandRolesGetAAAAttributeAPIV1, *resty.Response, error) {
	path := "/dna/system/api/v1/users/external-servers/aaa-attribute"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseUserandRolesGetAAAAttributeAPIV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAAAAttributeAPIV1()
		}
		return nil, response, fmt.Errorf("error with operation GetAAAAttributeApiV1")
	}

	result := response.Result().(*ResponseUserandRolesGetAAAAttributeAPIV1)
	return result, response, err

}

//AddRoleAPIV1 Add role API - b697-0a1e-46a9-b542
/* Add a new role in Cisco DNA Center System.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-role-api-v1
*/
func (s *UserandRolesService) AddRoleAPIV1(requestUserandRolesAddRoleAPIV1 *RequestUserandRolesAddRoleAPIV1) (*ResponseUserandRolesAddRoleAPIV1, *resty.Response, error) {
	path := "/dna/system/api/v1/role"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestUserandRolesAddRoleAPIV1).
		SetResult(&ResponseUserandRolesAddRoleAPIV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddRoleAPIV1(requestUserandRolesAddRoleAPIV1)
		}

		return nil, response, fmt.Errorf("error with operation AddRoleApiV1")
	}

	result := response.Result().(*ResponseUserandRolesAddRoleAPIV1)
	return result, response, err

}

//AddUserAPIV1 Add user API - 6c9a-09c4-4a39-9e2b
/* Add a new user for Cisco DNA Center System.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-user-api-v1
*/
func (s *UserandRolesService) AddUserAPIV1(requestUserandRolesAddUserAPIV1 *RequestUserandRolesAddUserAPIV1) (*ResponseUserandRolesAddUserAPIV1, *resty.Response, error) {
	path := "/dna/system/api/v1/user"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestUserandRolesAddUserAPIV1).
		SetResult(&ResponseUserandRolesAddUserAPIV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddUserAPIV1(requestUserandRolesAddUserAPIV1)
		}

		return nil, response, fmt.Errorf("error with operation AddUserApiV1")
	}

	result := response.Result().(*ResponseUserandRolesAddUserAPIV1)
	return result, response, err

}

//ManageExternalAuthenticationSettingAPIV1 Manage External Authentication Setting API - e09c-1806-48da-bb40
/* Enable or disable external authentication on Cisco DNA Center System.
Please find the Administrator Guide for your particular release from the list linked below and follow the steps required to enable external authentication before trying to do so from this API.
https://www.cisco.com/c/en/us/support/cloud-systems-management/dna-center/products-maintenance-guides-list.html



Documentation Link: https://developer.cisco.com/docs/dna-center/#!manage-external-authentication-setting-api-v1
*/
func (s *UserandRolesService) ManageExternalAuthenticationSettingAPIV1(requestUserandRolesManageExternalAuthenticationSettingAPIV1 *RequestUserandRolesManageExternalAuthenticationSettingAPIV1) (*ResponseUserandRolesManageExternalAuthenticationSettingAPIV1, *resty.Response, error) {
	path := "/dna/system/api/v1/users/external-authentication"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestUserandRolesManageExternalAuthenticationSettingAPIV1).
		SetResult(&ResponseUserandRolesManageExternalAuthenticationSettingAPIV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ManageExternalAuthenticationSettingAPIV1(requestUserandRolesManageExternalAuthenticationSettingAPIV1)
		}

		return nil, response, fmt.Errorf("error with operation ManageExternalAuthenticationSettingApiV1")
	}

	result := response.Result().(*ResponseUserandRolesManageExternalAuthenticationSettingAPIV1)
	return result, response, err

}

//AddAndUpdateAAAAttributeAPIV1 Add and Update AAA Attribute API - 808a-0aa0-491b-891f
/* Add or update the custom AAA attribute for external authentication. Note that if you decide not to set the custom AAA attribute, a default AAA attribute will be used for authentication based on the protocol supported by your server. For TACACS servers it will be "cisco-av-pair" and for RADIUS servers it will be "Cisco-AVPair".



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-and-update-a-a-a-attribute-api-v1
*/
func (s *UserandRolesService) AddAndUpdateAAAAttributeAPIV1(requestUserandRolesAddAndUpdateAAAAttributeAPIV1 *RequestUserandRolesAddAndUpdateAAAAttributeAPIV1) (*ResponseUserandRolesAddAndUpdateAAAAttributeAPIV1, *resty.Response, error) {
	path := "/dna/system/api/v1/users/external-servers/aaa-attribute"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestUserandRolesAddAndUpdateAAAAttributeAPIV1).
		SetResult(&ResponseUserandRolesAddAndUpdateAAAAttributeAPIV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddAndUpdateAAAAttributeAPIV1(requestUserandRolesAddAndUpdateAAAAttributeAPIV1)
		}

		return nil, response, fmt.Errorf("error with operation AddAndUpdateAAAAttributeApiV1")
	}

	result := response.Result().(*ResponseUserandRolesAddAndUpdateAAAAttributeAPIV1)
	return result, response, err

}

//UpdateRoleAPIV1 Update role API - 539c-ea73-400b-bf20
/* Update a role in Cisco DNA Center System.


 */
func (s *UserandRolesService) UpdateRoleAPIV1(requestUserandRolesUpdateRoleAPIV1 *RequestUserandRolesUpdateRoleAPIV1) (*ResponseUserandRolesUpdateRoleAPIV1, *resty.Response, error) {
	path := "/dna/system/api/v1/role"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestUserandRolesUpdateRoleAPIV1).
		SetResult(&ResponseUserandRolesUpdateRoleAPIV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateRoleAPIV1(requestUserandRolesUpdateRoleAPIV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateRoleApiV1")
	}

	result := response.Result().(*ResponseUserandRolesUpdateRoleAPIV1)
	return result, response, err

}

//UpdateUserAPIV1 Update user API - f596-6adc-492b-a2ff
/* Update a user for Cisco DNA Center System.


 */
func (s *UserandRolesService) UpdateUserAPIV1(requestUserandRolesUpdateUserAPIV1 *RequestUserandRolesUpdateUserAPIV1) (*ResponseUserandRolesUpdateUserAPIV1, *resty.Response, error) {
	path := "/dna/system/api/v1/user"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestUserandRolesUpdateUserAPIV1).
		SetResult(&ResponseUserandRolesUpdateUserAPIV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateUserAPIV1(requestUserandRolesUpdateUserAPIV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateUserApiV1")
	}

	result := response.Result().(*ResponseUserandRolesUpdateUserAPIV1)
	return result, response, err

}

//DeleteRoleAPIV1 Delete role API - d3b9-8bdc-472b-b236
/* Delete a role in Cisco DNA Center System


@param roleID roleId path parameter. The Id of the role to be deleted


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-role-api-v1
*/
func (s *UserandRolesService) DeleteRoleAPIV1(roleID string) (*ResponseUserandRolesDeleteRoleAPIV1, *resty.Response, error) {
	//roleID string
	path := "/dna/system/api/v1/role/{roleId}"
	path = strings.Replace(path, "{roleId}", fmt.Sprintf("%v", roleID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseUserandRolesDeleteRoleAPIV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteRoleAPIV1(
				roleID)
		}
		return nil, response, fmt.Errorf("error with operation DeleteRoleApiV1")
	}

	result := response.Result().(*ResponseUserandRolesDeleteRoleAPIV1)
	return result, response, err

}

//DeleteUserAPIV1 Delete user API - 69b4-ba37-4aca-8e86
/* Delete a user from Cisco DNA Center System.


@param userID userId path parameter. The id of the user to be deleted


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-user-api-v1
*/
func (s *UserandRolesService) DeleteUserAPIV1(userID string) (*ResponseUserandRolesDeleteUserAPIV1, *resty.Response, error) {
	//userID string
	path := "/dna/system/api/v1/user/{userId}"
	path = strings.Replace(path, "{userId}", fmt.Sprintf("%v", userID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseUserandRolesDeleteUserAPIV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteUserAPIV1(userID)
		}
		return nil, response, fmt.Errorf("error with operation DeleteUserApiV1")
	}

	result := response.Result().(*ResponseUserandRolesDeleteUserAPIV1)
	return result, response, err

}

//DeleteAAAAttributeAPIV1 Delete AAA Attribute API - d99e-c8df-4f1b-98fe
/* Delete the custom AAA attribute that was added. Note that by deleting the AAA attribute, a default AAA attribute will be used for authentication based on the protocol supported by your server. For TACACS servers it will be "cisco-av-pair" and for RADIUS servers it will be "Cisco-AVPair".



Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-a-a-a-attribute-api-v1
*/
func (s *UserandRolesService) DeleteAAAAttributeAPIV1() (*ResponseUserandRolesDeleteAAAAttributeAPIV1, *resty.Response, error) {
	//
	path := "/dna/system/api/v1/users/external-servers/aaa-attribute"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseUserandRolesDeleteAAAAttributeAPIV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteAAAAttributeAPIV1()
		}
		return nil, response, fmt.Errorf("error with operation DeleteAAAAttributeApiV1")
	}

	result := response.Result().(*ResponseUserandRolesDeleteAAAAttributeAPIV1)
	return result, response, err

}

// Alias Function
/*
This method acts as an alias for the method `AddAndUpdateAAAAttributeAPIV1`
*/
func (s *UserandRolesService) AddAndUpdateAAAAttributeAPI(requestUserandRolesAddAndUpdateAAAAttributeAPIV1 *RequestUserandRolesAddAndUpdateAAAAttributeAPIV1) (*ResponseUserandRolesAddAndUpdateAAAAttributeAPIV1, *resty.Response, error) {
	return s.AddAndUpdateAAAAttributeAPIV1(requestUserandRolesAddAndUpdateAAAAttributeAPIV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetUsersAPIV1`
*/
func (s *UserandRolesService) GetUsersAPI(GetUsersAPIV1QueryParams *GetUsersAPIV1QueryParams) (*ResponseUserandRolesGetUsersAPIV1, *resty.Response, error) {
	return s.GetUsersAPIV1(GetUsersAPIV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetExternalAuthenticationServersAPIV1`
*/
func (s *UserandRolesService) GetExternalAuthenticationServersAPI(GetExternalAuthenticationServersAPIV1QueryParams *GetExternalAuthenticationServersAPIV1QueryParams) (*ResponseUserandRolesGetExternalAuthenticationServersAPIV1, *resty.Response, error) {
	return s.GetExternalAuthenticationServersAPIV1(GetExternalAuthenticationServersAPIV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `AddUserAPIV1`
*/
func (s *UserandRolesService) AddUserAPI(requestUserandRolesAddUserAPIV1 *RequestUserandRolesAddUserAPIV1) (*ResponseUserandRolesAddUserAPIV1, *resty.Response, error) {
	return s.AddUserAPIV1(requestUserandRolesAddUserAPIV1)
}

// Alias Function
/*
This method acts as an alias for the method `DeleteAAAAttributeAPIV1`
*/
func (s *UserandRolesService) DeleteAAAAttributeAPI() (*ResponseUserandRolesDeleteAAAAttributeAPIV1, *resty.Response, error) {
	return s.DeleteAAAAttributeAPIV1()
}

// Alias Function
/*
This method acts as an alias for the method `DeleteUserAPIV1`
*/
func (s *UserandRolesService) DeleteUserAPI(userID string) (*ResponseUserandRolesDeleteUserAPIV1, *resty.Response, error) {
	return s.DeleteUserAPIV1(userID)
}

// Alias Function
/*
This method acts as an alias for the method `GetExternalAuthenticationSettingAPIV1`
*/
func (s *UserandRolesService) GetExternalAuthenticationSettingAPI() (*ResponseUserandRolesGetExternalAuthenticationSettingAPIV1, *resty.Response, error) {
	return s.GetExternalAuthenticationSettingAPIV1()
}

// Alias Function
/*
This method acts as an alias for the method `GetRolesAPIV1`
*/
func (s *UserandRolesService) GetRolesAPI(GetRolesAPIV1HeaderParams *GetRolesAPIV1HeaderParams) (*ResponseUserandRolesGetRolesAPIV1, *resty.Response, error) {
	return s.GetRolesAPIV1(GetRolesAPIV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `AddRoleAPIV1`
*/
func (s *UserandRolesService) AddRoleAPI(requestUserandRolesAddRoleAPIV1 *RequestUserandRolesAddRoleAPIV1) (*ResponseUserandRolesAddRoleAPIV1, *resty.Response, error) {
	return s.AddRoleAPIV1(requestUserandRolesAddRoleAPIV1)
}

// Alias Function
/*
This method acts as an alias for the method `UpdateUserAPIV1`
*/
func (s *UserandRolesService) UpdateUserAPI(requestUserandRolesUpdateUserAPIV1 *RequestUserandRolesUpdateUserAPIV1) (*ResponseUserandRolesUpdateUserAPIV1, *resty.Response, error) {
	return s.UpdateUserAPIV1(requestUserandRolesUpdateUserAPIV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetPermissionsAPIV1`
*/
func (s *UserandRolesService) GetPermissionsAPI() (*ResponseUserandRolesGetPermissionsAPIV1, *resty.Response, error) {
	return s.GetPermissionsAPIV1()
}

// Alias Function
/*
This method acts as an alias for the method `GetAAAAttributeAPIV1`
*/
func (s *UserandRolesService) GetAAAAttributeAPI() (*ResponseUserandRolesGetAAAAttributeAPIV1, *resty.Response, error) {
	return s.GetAAAAttributeAPIV1()
}

// Alias Function
/*
This method acts as an alias for the method `ManageExternalAuthenticationSettingAPIV1`
*/
func (s *UserandRolesService) ManageExternalAuthenticationSettingAPI(requestUserandRolesManageExternalAuthenticationSettingAPIV1 *RequestUserandRolesManageExternalAuthenticationSettingAPIV1) (*ResponseUserandRolesManageExternalAuthenticationSettingAPIV1, *resty.Response, error) {
	return s.ManageExternalAuthenticationSettingAPIV1(requestUserandRolesManageExternalAuthenticationSettingAPIV1)
}

// Alias Function
/*
This method acts as an alias for the method `DeleteRoleAPIV1`
*/
func (s *UserandRolesService) DeleteRoleAPI(roleID string) (*ResponseUserandRolesDeleteRoleAPIV1, *resty.Response, error) {
	return s.DeleteRoleAPIV1(roleID)
}

// Alias Function
/*
This method acts as an alias for the method `UpdateRoleAPIV1`
*/
func (s *UserandRolesService) UpdateRoleAPI(requestUserandRolesUpdateRoleAPIV1 *RequestUserandRolesUpdateRoleAPIV1) (*ResponseUserandRolesUpdateRoleAPIV1, *resty.Response, error) {
	return s.UpdateRoleAPIV1(requestUserandRolesUpdateRoleAPIV1)
}
