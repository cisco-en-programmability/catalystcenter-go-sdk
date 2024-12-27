package catalyst

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
)

type ItsmIntegrationService service

type ResponseItsmIntegrationCreateItsmIntegrationSettingV1 struct {
	ID                 string                                                                     `json:"id,omitempty"`                 // Id
	DypID              string                                                                     `json:"dypId,omitempty"`              // Dyp Id
	DypName            string                                                                     `json:"dypName,omitempty"`            // Dyp Name
	Name               string                                                                     `json:"name,omitempty"`               // Name
	UniqueKey          string                                                                     `json:"uniqueKey,omitempty"`          // Unique Key
	DypMajorVersion    *int                                                                       `json:"dypMajorVersion,omitempty"`    // Dyp Major Version
	Description        string                                                                     `json:"description,omitempty"`        // Description
	Data               *ResponseItsmIntegrationCreateItsmIntegrationSettingV1Data                 `json:"data,omitempty"`               //
	CreatedDate        *int                                                                       `json:"createdDate,omitempty"`        // Created Date
	CreatedBy          string                                                                     `json:"createdBy,omitempty"`          // Created By
	UpdatedBy          string                                                                     `json:"updatedBy,omitempty"`          // Updated By
	SoftwareVersionLog *[]ResponseItsmIntegrationCreateItsmIntegrationSettingV1SoftwareVersionLog `json:"softwareVersionLog,omitempty"` // Software Version Log
	SchemaVersion      *float64                                                                   `json:"schemaVersion,omitempty"`      // Schema Version
	TenantID           string                                                                     `json:"tenantId,omitempty"`           // Tenant Id
}
type ResponseItsmIntegrationCreateItsmIntegrationSettingV1Data struct {
	ConnectionSettings *ResponseItsmIntegrationCreateItsmIntegrationSettingV1DataConnectionSettings `json:"ConnectionSettings,omitempty"` //
}
type ResponseItsmIntegrationCreateItsmIntegrationSettingV1DataConnectionSettings struct {
	URL          string `json:"Url,omitempty"`           // Url
	AuthUserName string `json:"Auth_UserName,omitempty"` // Auth User Name
	AuthPassword string `json:"Auth_Password,omitempty"` // Auth Password
}
type ResponseItsmIntegrationCreateItsmIntegrationSettingV1SoftwareVersionLog interface{}
type ResponseItsmIntegrationUpdateItsmIntegrationSettingV1 struct {
	TypeID          string                                                     `json:"_id,omitempty"`             // Id
	ID              string                                                     `json:"id,omitempty"`              // Id
	DypID           string                                                     `json:"dypId,omitempty"`           // Dyp Id
	DypName         string                                                     `json:"dypName,omitempty"`         // Dyp Name
	DypMajorVersion *int                                                       `json:"dypMajorVersion,omitempty"` // Dyp Major Version
	Name            string                                                     `json:"name,omitempty"`            // Name
	UniqueKey       string                                                     `json:"uniqueKey,omitempty"`       // Unique Key
	Description     string                                                     `json:"description,omitempty"`     // Description
	Data            *ResponseItsmIntegrationUpdateItsmIntegrationSettingV1Data `json:"data,omitempty"`            //
	UpdatedDate     *int                                                       `json:"updatedDate,omitempty"`     // Updated Date
	UpdatedBy       string                                                     `json:"updatedBy,omitempty"`       // Updated By
	TenantID        string                                                     `json:"tenantId,omitempty"`        // Tenant Id
}
type ResponseItsmIntegrationUpdateItsmIntegrationSettingV1Data struct {
	ConnectionSettings *ResponseItsmIntegrationUpdateItsmIntegrationSettingV1DataConnectionSettings `json:"ConnectionSettings,omitempty"` //
}
type ResponseItsmIntegrationUpdateItsmIntegrationSettingV1DataConnectionSettings struct {
	URL          string `json:"Url,omitempty"`           // Url
	AuthUserName string `json:"Auth_UserName,omitempty"` // Auth User Name
	AuthPassword string `json:"Auth_Password,omitempty"` // Auth Password
}
type ResponseItsmIntegrationGetItsmIntegrationSettingByIDV1 struct {
	TypeID          string                                                      `json:"_id,omitempty"`             // Id
	ID              string                                                      `json:"id,omitempty"`              // Id
	DypID           string                                                      `json:"dypId,omitempty"`           // Dyp Id
	DypName         string                                                      `json:"dypName,omitempty"`         // Dyp Name
	DypMajorVersion *int                                                        `json:"dypMajorVersion,omitempty"` // Dyp Major Version
	Name            string                                                      `json:"name,omitempty"`            // Name
	UniqueKey       string                                                      `json:"uniqueKey,omitempty"`       // Unique Key
	Description     string                                                      `json:"description,omitempty"`     // Description
	Data            *ResponseItsmIntegrationGetItsmIntegrationSettingByIDV1Data `json:"data,omitempty"`            //
	UpdatedDate     *int                                                        `json:"updatedDate,omitempty"`     // Updated Date
	UpdatedBy       string                                                      `json:"updatedBy,omitempty"`       // Updated By
	TenantID        string                                                      `json:"tenantId,omitempty"`        // Tenant Id
}
type ResponseItsmIntegrationGetItsmIntegrationSettingByIDV1Data struct {
	ConnectionSettings *ResponseItsmIntegrationGetItsmIntegrationSettingByIDV1DataConnectionSettings `json:"ConnectionSettings,omitempty"` //
}
type ResponseItsmIntegrationGetItsmIntegrationSettingByIDV1DataConnectionSettings struct {
	URL          string `json:"Url,omitempty"`           // Url
	AuthUserName string `json:"Auth_UserName,omitempty"` // Auth User Name
	AuthPassword string `json:"Auth_Password,omitempty"` // Auth Password
}
type ResponseItsmIntegrationGetAllItsmIntegrationSettingsV1 []ResponseItemItsmIntegrationGetAllItsmIntegrationSettingsV1 // Array of ResponseItsmIntegrationGetAllITSMIntegrationSettingsV1
type ResponseItemItsmIntegrationGetAllItsmIntegrationSettingsV1 struct {
	ID                 string                                                                          `json:"id,omitempty"`                 // Id
	DypID              string                                                                          `json:"dypId,omitempty"`              // Dyp Id
	DypName            string                                                                          `json:"dypName,omitempty"`            // Dyp Name
	Name               string                                                                          `json:"name,omitempty"`               // Name
	UniqueKey          string                                                                          `json:"uniqueKey,omitempty"`          // Unique Key
	DypMajorVersion    *int                                                                            `json:"dypMajorVersion,omitempty"`    // Dyp Major Version
	Description        string                                                                          `json:"description,omitempty"`        // Description
	CreatedDate        *int                                                                            `json:"createdDate,omitempty"`        // Created Date
	CreatedBy          string                                                                          `json:"createdBy,omitempty"`          // Created By
	UpdatedBy          string                                                                          `json:"updatedBy,omitempty"`          // Updated By
	SoftwareVersionLog *[]ResponseItemItsmIntegrationGetAllItsmIntegrationSettingsV1SoftwareVersionLog `json:"softwareVersionLog,omitempty"` // Software Version Log
	SchemaVersion      *float64                                                                        `json:"schemaVersion,omitempty"`      // Schema Version
	TenantID           string                                                                          `json:"tenantId,omitempty"`           // Tenant Id
}
type ResponseItemItsmIntegrationGetAllItsmIntegrationSettingsV1SoftwareVersionLog interface{}
type ResponseItsmIntegrationGetItsmIntegrationStatusV1 struct {
	Response *[]ResponseItsmIntegrationGetItsmIntegrationStatusV1Response `json:"response,omitempty"` //
	Version  string                                                       `json:"version,omitempty"`  // Version
}
type ResponseItsmIntegrationGetItsmIntegrationStatusV1Response struct {
	ID             string                                                                     `json:"id,omitempty"`             // Bundle Id
	Name           string                                                                     `json:"name,omitempty"`           // Bundle name
	Status         string                                                                     `json:"status,omitempty"`         // Bundle Status
	Configurations *[]ResponseItsmIntegrationGetItsmIntegrationStatusV1ResponseConfigurations `json:"configurations,omitempty"` //
}
type ResponseItsmIntegrationGetItsmIntegrationStatusV1ResponseConfigurations struct {
	DypSchemaName string `json:"dypSchemaName,omitempty"` // DYP name of the configuration
	DypInstanceID string `json:"dypInstanceId,omitempty"` // DYP instance Id of the configuration
}
type RequestItsmIntegrationCreateItsmIntegrationSettingV1 struct {
	Name        string                                                    `json:"name,omitempty"`        // Name of the setting instance
	Description string                                                    `json:"description,omitempty"` // Description of the setting instance
	Data        *RequestItsmIntegrationCreateItsmIntegrationSettingV1Data `json:"data,omitempty"`        //
	DypName     string                                                    `json:"dypName,omitempty"`     // It can be ServiceNowConnection
}
type RequestItsmIntegrationCreateItsmIntegrationSettingV1Data struct {
	ConnectionSettings *RequestItsmIntegrationCreateItsmIntegrationSettingV1DataConnectionSettings `json:"ConnectionSettings,omitempty"` //
}
type RequestItsmIntegrationCreateItsmIntegrationSettingV1DataConnectionSettings struct {
	URL          string `json:"Url,omitempty"`           // Url
	AuthUserName string `json:"Auth_UserName,omitempty"` // Auth User Name
	AuthPassword string `json:"Auth_Password,omitempty"` // Auth Password
}
type RequestItsmIntegrationUpdateItsmIntegrationSettingV1 struct {
	Name        string                                                    `json:"name,omitempty"`        // Name of the setting instance
	Description string                                                    `json:"description,omitempty"` // Description of the setting instance
	Data        *RequestItsmIntegrationUpdateItsmIntegrationSettingV1Data `json:"data,omitempty"`        //
	DypName     string                                                    `json:"dypName,omitempty"`     // It can be ServiceNowConnection
}
type RequestItsmIntegrationUpdateItsmIntegrationSettingV1Data struct {
	ConnectionSettings *RequestItsmIntegrationUpdateItsmIntegrationSettingV1DataConnectionSettings `json:"ConnectionSettings,omitempty"` //
}
type RequestItsmIntegrationUpdateItsmIntegrationSettingV1DataConnectionSettings struct {
	URL          string `json:"Url,omitempty"`           // Url
	AuthUserName string `json:"Auth_UserName,omitempty"` // Auth User Name
	AuthPassword string `json:"Auth_Password,omitempty"` // Auth Password
}

//GetItsmIntegrationSettingByIDV1 Get ITSM Integration setting by Id - 1086-aa18-4cda-8471
/* Fetches ITSM Integration setting by ID


@param instanceID instanceId path parameter. Instance Id of the Integration setting instance


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-itsm-integration-setting-by-id
*/
func (s *ItsmIntegrationService) GetItsmIntegrationSettingByIDV1(instanceID string) (*ResponseItsmIntegrationGetItsmIntegrationSettingByIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/integration-settings/instances/itsm/{instanceId}"
	path = strings.Replace(path, "{instanceId}", fmt.Sprintf("%v", instanceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseItsmIntegrationGetItsmIntegrationSettingByIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetItsmIntegrationSettingByIDV1(instanceID)
		}
		return nil, response, fmt.Errorf("error with operation GetItsmIntegrationSettingByIdV1")
	}

	result := response.Result().(*ResponseItsmIntegrationGetItsmIntegrationSettingByIDV1)
	return result, response, err

}

//GetAllItsmIntegrationSettingsV1 Get all ITSM Integration settings - 468a-fb23-4379-86eb
/* Fetches all ITSM Integration settings



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-all-itsm-integration-settings
*/
func (s *ItsmIntegrationService) GetAllItsmIntegrationSettingsV1() (*ResponseItsmIntegrationGetAllItsmIntegrationSettingsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/integration-settings/itsm/instances"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseItsmIntegrationGetAllItsmIntegrationSettingsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAllItsmIntegrationSettingsV1()
		}
		return nil, response, fmt.Errorf("error with operation GetAllItsmIntegrationSettingsV1")
	}

	result := response.Result().(*ResponseItsmIntegrationGetAllItsmIntegrationSettingsV1)
	return result, response, err

}

//GetItsmIntegrationStatusV1 Get ITSM Integration status - b7a2-ea02-4e69-abdf
/* Fetches ITSM Integration status



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-itsm-integration-status
*/
func (s *ItsmIntegrationService) GetItsmIntegrationStatusV1() (*ResponseItsmIntegrationGetItsmIntegrationStatusV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/integration-settings/status"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseItsmIntegrationGetItsmIntegrationStatusV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetItsmIntegrationStatusV1()
		}
		return nil, response, fmt.Errorf("error with operation GetItsmIntegrationStatusV1")
	}

	result := response.Result().(*ResponseItsmIntegrationGetItsmIntegrationStatusV1)
	return result, response, err

}

//CreateItsmIntegrationSettingV1 Create ITSM Integration setting - 0cb0-1a15-4d79-a440
/* Creates ITSM Integration setting



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-itsm-integration-setting
*/
func (s *ItsmIntegrationService) CreateItsmIntegrationSettingV1(requestItsmIntegrationCreateITSMIntegrationSettingV1 *RequestItsmIntegrationCreateItsmIntegrationSettingV1) (*ResponseItsmIntegrationCreateItsmIntegrationSettingV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/integration-settings/instances/itsm"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestItsmIntegrationCreateITSMIntegrationSettingV1).
		SetResult(&ResponseItsmIntegrationCreateItsmIntegrationSettingV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateItsmIntegrationSettingV1(requestItsmIntegrationCreateITSMIntegrationSettingV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateItsmIntegrationSettingV1")
	}

	result := response.Result().(*ResponseItsmIntegrationCreateItsmIntegrationSettingV1)
	return result, response, err

}

//UpdateItsmIntegrationSettingV1 Update ITSM Integration setting - 5fbe-680a-4208-92e6
/* Updates the ITSM Integration setting


@param instanceID instanceId path parameter. Instance Id of the Integration setting instance

*/
func (s *ItsmIntegrationService) UpdateItsmIntegrationSettingV1(instanceID string, requestItsmIntegrationUpdateITSMIntegrationSettingV1 *RequestItsmIntegrationUpdateItsmIntegrationSettingV1) (*ResponseItsmIntegrationUpdateItsmIntegrationSettingV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/integration-settings/instances/itsm/{instanceId}"
	path = strings.Replace(path, "{instanceId}", fmt.Sprintf("%v", instanceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestItsmIntegrationUpdateITSMIntegrationSettingV1).
		SetResult(&ResponseItsmIntegrationUpdateItsmIntegrationSettingV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateItsmIntegrationSettingV1(instanceID, requestItsmIntegrationUpdateITSMIntegrationSettingV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateItsmIntegrationSettingV1")
	}

	result := response.Result().(*ResponseItsmIntegrationUpdateItsmIntegrationSettingV1)
	return result, response, err

}

//DeleteItsmIntegrationSettingV1 Delete ITSM Integration setting - e8b9-ba8b-4c29-b637
/* Deletes the ITSM Integration setting


@param instanceID instanceId path parameter. Instance Id of the Integration setting instance


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-itsm-integration-setting
*/
func (s *ItsmIntegrationService) DeleteItsmIntegrationSettingV1(instanceID string) (*resty.Response, error) {
	//instanceID string
	path := "/dna/intent/api/v1/integration-settings/instances/itsm/{instanceId}"
	path = strings.Replace(path, "{instanceId}", fmt.Sprintf("%v", instanceID), -1)

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
			return s.DeleteItsmIntegrationSettingV1(instanceID)
		}
		return response, fmt.Errorf("error with operation DeleteItsmIntegrationSettingV1")
	}

	return response, err

}

// Alias Function
/*
This method acts as an alias for the method `CreateItsmIntegrationSettingV1`
*/
func (s *ItsmIntegrationService) CreateItsmIntegrationSetting(requestItsmIntegrationCreateITSMIntegrationSettingV1 *RequestItsmIntegrationCreateItsmIntegrationSettingV1) (*ResponseItsmIntegrationCreateItsmIntegrationSettingV1, *resty.Response, error) {
	return s.CreateItsmIntegrationSettingV1(requestItsmIntegrationCreateITSMIntegrationSettingV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetAllItsmIntegrationSettingsV1`
*/
func (s *ItsmIntegrationService) GetAllItsmIntegrationSettings() (*ResponseItsmIntegrationGetAllItsmIntegrationSettingsV1, *resty.Response, error) {
	return s.GetAllItsmIntegrationSettingsV1()
}

// Alias Function
/*
This method acts as an alias for the method `GetItsmIntegrationStatusV1`
*/
func (s *ItsmIntegrationService) GetItsmIntegrationStatus() (*ResponseItsmIntegrationGetItsmIntegrationStatusV1, *resty.Response, error) {
	return s.GetItsmIntegrationStatusV1()
}

// Alias Function
/*
This method acts as an alias for the method `UpdateItsmIntegrationSettingV1`
*/
func (s *ItsmIntegrationService) UpdateItsmIntegrationSetting(instanceID string, requestItsmIntegrationUpdateITSMIntegrationSettingV1 *RequestItsmIntegrationUpdateItsmIntegrationSettingV1) (*ResponseItsmIntegrationUpdateItsmIntegrationSettingV1, *resty.Response, error) {
	return s.UpdateItsmIntegrationSettingV1(instanceID, requestItsmIntegrationUpdateITSMIntegrationSettingV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetItsmIntegrationSettingByIDV1`
*/
func (s *ItsmIntegrationService) GetItsmIntegrationSettingByID(instanceID string) (*ResponseItsmIntegrationGetItsmIntegrationSettingByIDV1, *resty.Response, error) {
	return s.GetItsmIntegrationSettingByIDV1(instanceID)
}

// Alias Function
/*
This method acts as an alias for the method `DeleteItsmIntegrationSettingV1`
*/
func (s *ItsmIntegrationService) DeleteItsmIntegrationSetting(instanceID string) (*resty.Response, error) {
	return s.DeleteItsmIntegrationSettingV1(instanceID)
}
