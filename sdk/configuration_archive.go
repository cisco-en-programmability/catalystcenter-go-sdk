package catalyst

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type ConfigurationArchiveService service

type GetConfigurationArchiveDetailsV1QueryParams struct {
	DeviceID    string  `url:"deviceId,omitempty"`    //comma separated device id for example cf35b0a1-407f-412f-b2f4-f0c3156695f9,aaa38191-0c22-4158-befd-779a09d7cec1 . if device id is not provided it will fetch for all devices
	FileType    string  `url:"fileType,omitempty"`    //Config File Type can be RUNNINGCONFIG or STARTUPCONFIG
	CreatedTime string  `url:"createdTime,omitempty"` //Supported with logical filters GT,GTE,LT,LTE & BT : time in milliseconds (epoc format)
	CreatedBy   string  `url:"createdBy,omitempty"`   //Comma separated values for createdBy - SCHEDULED, USER, CONFIG_CHANGE_EVENT, SCHEDULED_FIRST_TIME, DR_CALL_BACK, PRE_DEPLOY
	Offset      float64 `url:"offset,omitempty"`      //offset
	Limit       float64 `url:"limit,omitempty"`       //The number of records to be retrieved defaults to 500 if not specified, with a maximum allowed limit of 500.
}
type GetNetworkDeviceConfigurationFileDetailsV1QueryParams struct {
	ID              string  `url:"id,omitempty"`              //Unique identifier (UUID) of the configuration file.
	NetworkDeviceID string  `url:"networkDeviceId,omitempty"` //Unique identifier (UUID) of the network devices. The number of networkDeviceId(s) must not exceed 5.
	FileType        string  `url:"fileType,omitempty"`        //Type of device configuration file.Available values : 'RUNNINGCONFIG', 'STARTUPCONFIG', 'VLAN'
	Offset          float64 `url:"offset,omitempty"`          //The first record to show for this page; the first record is numbered 1.
	Limit           float64 `url:"limit,omitempty"`           //The number of records to be retrieved defaults to 500 if not specified, with a maximum allowed limit of 500.
}
type CountOfNetworkDeviceConfigurationFilesV1QueryParams struct {
	ID              string `url:"id,omitempty"`              //Unique identifier (UUID) of the configuration file.
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` //Unique identifier (UUID) of the network devices. The number of networkDeviceId(s) must not exceed 5.
	FileType        string `url:"fileType,omitempty"`        //Type of device configuration file. Available values : 'RUNNINGCONFIG', 'STARTUPCONFIG', 'VLAN'
}

type ResponseConfigurationArchiveExportDeviceConfigurationsV1 struct {
	Version  string                                                            `json:"version,omitempty"`  // Version
	Response *ResponseConfigurationArchiveExportDeviceConfigurationsV1Response `json:"response,omitempty"` //
}
type ResponseConfigurationArchiveExportDeviceConfigurationsV1Response struct {
	URL    string `json:"url,omitempty"`    // Url
	TaskID string `json:"taskId,omitempty"` // Task Id
}
type ResponseConfigurationArchiveGetConfigurationArchiveDetailsV1 []ResponseItemConfigurationArchiveGetConfigurationArchiveDetailsV1 // Array of ResponseConfigurationArchiveGetConfigurationArchiveDetailsV1
type ResponseItemConfigurationArchiveGetConfigurationArchiveDetailsV1 struct {
	IPAddress  string                                                                      `json:"ipAddress,omitempty"`  // IP address of the device.
	DeviceID   string                                                                      `json:"deviceId,omitempty"`   // UUID of the device.
	Versions   *[]ResponseItemConfigurationArchiveGetConfigurationArchiveDetailsV1Versions `json:"versions,omitempty"`   //
	DeviceName string                                                                      `json:"deviceName,omitempty"` // Hostname of the device.
}
type ResponseItemConfigurationArchiveGetConfigurationArchiveDetailsV1Versions struct {
	Files                *[]ResponseItemConfigurationArchiveGetConfigurationArchiveDetailsV1VersionsFiles                `json:"files,omitempty"`                //
	CreatedBy            string                                                                                          `json:"createdBy,omitempty"`            // Reason for archive collection (CONFIG_CHANGE_EVENT - Syslog event based colletion, SCHEDULED - Weekly scheduled collection, SCHEDULED_FIRST_TIME-First Time Managed collection,DR_CALL_BACK- Post Disaster Recovery collection, USER- On Demand Trigger, PRE_DEPLOY- Predeploy collection)
	ConfigChangeType     string                                                                                          `json:"configChangeType,omitempty"`     // Source of configuration change (OUT_OF_BAND - Change was made outside Catalyst Center, IN_BAND - Change was made from Catalyst Center, NOT_APPLICABLE - System Triggered)
	SyslogConfigEventDto *[]ResponseItemConfigurationArchiveGetConfigurationArchiveDetailsV1VersionsSyslogConfigEventDto `json:"syslogConfigEventDto,omitempty"` //
	CreatedTime          *float64                                                                                        `json:"createdTime,omitempty"`          // Source of configuration change (OUT_OF_BAND - Change was made outside Catalyst Center, IN_BAND - Change was made from Catalyst Center, NOT_APPLICABLE - System Triggered)
	StartupRunningStatus string                                                                                          `json:"startupRunningStatus,omitempty"` // Sync status of running and startup configurations (IN_SYNC - if both startup and running config are same excluding dynamic configurations, OUT_OF_SYNC - otherwise).
	ID                   string                                                                                          `json:"id,omitempty"`                   // Unique version ID.
	Tags                 []string                                                                                        `json:"tags,omitempty"`                 // Labels added to a configuration version.
	LastUpdatedTime      *float64                                                                                        `json:"lastUpdatedTime,omitempty"`      // Latest time stamp when the collected configuration is verified to be running on the device. LastUpdatedTime and createdTime will differ when verified configs are the same.
}
type ResponseItemConfigurationArchiveGetConfigurationArchiveDetailsV1VersionsFiles struct {
	FileType     string `json:"fileType,omitempty"`     // Type of configuration file (RUNNINGCONFIG, STARTUPCONFIG, or VLAN).
	FileID       string `json:"fileId,omitempty"`       // Unique file ID.
	DownloadPath string `json:"downloadPath,omitempty"` // File download path (deprecated).
}
type ResponseItemConfigurationArchiveGetConfigurationArchiveDetailsV1VersionsSyslogConfigEventDto struct {
	UserName       string   `json:"userName,omitempty"`       // Name of the user who made the configuration change (if available in Syslog).
	DeviceUUID     string   `json:"deviceUuid,omitempty"`     // UUID of the device as recieved in syslog.
	OutOfBand      *bool    `json:"outOfBand,omitempty"`      // True if configuration changes were made from outside of the Catalist Center. False otherwise.
	ConfigMethod   string   `json:"configMethod,omitempty"`   // Connection mode used to do the changes pn the device (if available in Syslog).
	TerminalName   string   `json:"terminalName,omitempty"`   // Name of the terminal used to make changes on the device (if available in Syslog).
	LoginIPAddress string   `json:"loginIpAddress,omitempty"` // IP address from which the configuration changes were made (if available in Syslog).
	ProcessName    string   `json:"processName,omitempty"`    // Name of the process that made configuration change (only available when configuration got changed by a program such as YANG suite )
	SyslogTime     *float64 `json:"syslogTime,omitempty"`     // Time of configuration change as recorded in the syslog.
}
type ResponseConfigurationArchiveGetNetworkDeviceConfigurationFileDetailsV1 struct {
	Response *[]ResponseConfigurationArchiveGetNetworkDeviceConfigurationFileDetailsV1Response `json:"response,omitempty"` //
	Version  string                                                                            `json:"version,omitempty"`  // The version of API.
}
type ResponseConfigurationArchiveGetNetworkDeviceConfigurationFileDetailsV1Response struct {
	ID              string `json:"id,omitempty"`              // Unique identifier (UUID) of the configuration file.
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Unique identifier (UUID) of the network devices.
	VersionID       string `json:"versionId,omitempty"`       // The version unique identifier triggered after any config change.
	FileType        string `json:"fileType,omitempty"`        // Type of configuration file. Config File Type can be 'RUNNINGCONFIG' or 'STARTUPCONFIG' or 'VLAN'.
	CreatedBy       string `json:"createdBy,omitempty"`       // The entity responsible for creating the configuration changes.
	CreatedTime     *int   `json:"createdTime,omitempty"`     // The UNIX epoch timestamp in milliseconds marking when the resource was created.
}
type ResponseConfigurationArchiveCountOfNetworkDeviceConfigurationFilesV1 struct {
	Response *ResponseConfigurationArchiveCountOfNetworkDeviceConfigurationFilesV1Response `json:"response,omitempty"` //
	Version  string                                                                        `json:"version,omitempty"`  // Version
}
type ResponseConfigurationArchiveCountOfNetworkDeviceConfigurationFilesV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseConfigurationArchiveGetConfigurationFileDetailsByIDV1 struct {
	Response *ResponseConfigurationArchiveGetConfigurationFileDetailsByIDV1Response `json:"response,omitempty"` //
	Version  string                                                                 `json:"version,omitempty"`  // The version of API.
}
type ResponseConfigurationArchiveGetConfigurationFileDetailsByIDV1Response struct {
	ID              string `json:"id,omitempty"`              // Unique identifier (UUID) of the configuration file.
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Unique identifier (UUID) of the network devices.
	VersionID       string `json:"versionId,omitempty"`       // The version unique identifier triggered after any config change.
	FileType        string `json:"fileType,omitempty"`        // Type of configuration file. Config File Type can be 'RUNNINGCONFIG' or 'STARTUPCONFIG' or 'VLAN'.
	CreatedBy       string `json:"createdBy,omitempty"`       // The entity responsible for creating the configuration changes.
	CreatedTime     string `json:"createdTime,omitempty"`     // The UNIX epoch timestamp in milliseconds marking when the resource was created.
}
type ResponseConfigurationArchiveDownloadMaskedDeviceConfigurationV1 interface{}
type ResponseConfigurationArchiveDownloadUnmaskedrawDeviceConfigurationAsZIPV1 interface{}
type RequestConfigurationArchiveExportDeviceConfigurationsV1 struct {
	Password string `json:"password,omitempty"` // Password for the zip file to protect exported configurations. Must contain, at minimum 8 characters, one lowercase letter, one uppercase letter, one number, one special character(-=[];,./~!@#$%^&*()_+{}|:?). It may not contain white space or the characters <>.
	DeviceID string `json:"deviceId,omitempty"` // UUIDs of the devices for which configurations need to be exported.
}
type RequestConfigurationArchiveDownloadUnmaskedrawDeviceConfigurationAsZIPV1 struct {
	Password string `json:"password,omitempty"` // Password for the zip file to protect exported configurations. Must contain, at minimum 8 characters, one lowercase letter, one uppercase letter, one number, one special character(-=[];,./~!@#$%^&*()_+{}|:?). It may not contain white space or the characters <>.
}

//GetConfigurationArchiveDetailsV1 Get configuration archive details - 3bba-48a9-422a-be1e
/* Returns the historical device configurations (running configuration , startup configuration , vlan if applicable) by specified criteria


@param GetConfigurationArchiveDetailsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-configuration-archive-details
*/
func (s *ConfigurationArchiveService) GetConfigurationArchiveDetailsV1(GetConfigurationArchiveDetailsV1QueryParams *GetConfigurationArchiveDetailsV1QueryParams) (*ResponseConfigurationArchiveGetConfigurationArchiveDetailsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device-config"

	queryString, _ := query.Values(GetConfigurationArchiveDetailsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseConfigurationArchiveGetConfigurationArchiveDetailsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetConfigurationArchiveDetailsV1(GetConfigurationArchiveDetailsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetConfigurationArchiveDetailsV1")
	}

	result := response.Result().(*ResponseConfigurationArchiveGetConfigurationArchiveDetailsV1)
	return result, response, err

}

//GetNetworkDeviceConfigurationFileDetailsV1 Get Network Device Configuration File Details - bd95-9a71-4b8a-9442
/* Retrieves the list of network device configuration file details, sorted by createdTime in descending order. Use /intent/api/v1/networkDeviceConfigFiles/{id}/downloadMasked to download masked configurations, or /intent/api/v1/networkDeviceConfigFiles/{id}/downloadUnmasked for unmasked configurations.


@param GetNetworkDeviceConfigurationFileDetailsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-device-configuration-file-details
*/
func (s *ConfigurationArchiveService) GetNetworkDeviceConfigurationFileDetailsV1(GetNetworkDeviceConfigurationFileDetailsV1QueryParams *GetNetworkDeviceConfigurationFileDetailsV1QueryParams) (*ResponseConfigurationArchiveGetNetworkDeviceConfigurationFileDetailsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDeviceConfigFiles"

	queryString, _ := query.Values(GetNetworkDeviceConfigurationFileDetailsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseConfigurationArchiveGetNetworkDeviceConfigurationFileDetailsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetNetworkDeviceConfigurationFileDetailsV1(GetNetworkDeviceConfigurationFileDetailsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetNetworkDeviceConfigurationFileDetailsV1")
	}

	result := response.Result().(*ResponseConfigurationArchiveGetNetworkDeviceConfigurationFileDetailsV1)
	return result, response, err

}

//CountOfNetworkDeviceConfigurationFilesV1 Count of Network Device Configuration Files - d296-cab3-4a6b-b826
/* Retrieves count the details of the network device configuration files.


@param CountOfNetworkDeviceConfigurationFilesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!count-of-network-device-configuration-files
*/
func (s *ConfigurationArchiveService) CountOfNetworkDeviceConfigurationFilesV1(CountOfNetworkDeviceConfigurationFilesV1QueryParams *CountOfNetworkDeviceConfigurationFilesV1QueryParams) (*ResponseConfigurationArchiveCountOfNetworkDeviceConfigurationFilesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDeviceConfigFiles/count"

	queryString, _ := query.Values(CountOfNetworkDeviceConfigurationFilesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseConfigurationArchiveCountOfNetworkDeviceConfigurationFilesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.CountOfNetworkDeviceConfigurationFilesV1(CountOfNetworkDeviceConfigurationFilesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation CountOfNetworkDeviceConfigurationFilesV1")
	}

	result := response.Result().(*ResponseConfigurationArchiveCountOfNetworkDeviceConfigurationFilesV1)
	return result, response, err

}

//GetConfigurationFileDetailsByIDV1 Get Configuration File Details by ID - cc93-5822-44ab-b75f
/* Retrieves the details of a specific network device configuration file using the `id`.


@param id id path parameter. The value of `id` can be obtained from the response of API `/dna/intent/api/v1/networkDeviceConfigFiles`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-configuration-file-details-by-id
*/
func (s *ConfigurationArchiveService) GetConfigurationFileDetailsByIDV1(id string) (*ResponseConfigurationArchiveGetConfigurationFileDetailsByIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDeviceConfigFiles/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseConfigurationArchiveGetConfigurationFileDetailsByIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetConfigurationFileDetailsByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation GetConfigurationFileDetailsByIdV1")
	}

	result := response.Result().(*ResponseConfigurationArchiveGetConfigurationFileDetailsByIDV1)
	return result, response, err

}

//ExportDeviceConfigurationsV1 Export Device configurations - 51a4-0aba-4c68-ac17
/* Export Device configuration for every device that is provided will be included in an encrypted zip file.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!export-device-configurations
*/
func (s *ConfigurationArchiveService) ExportDeviceConfigurationsV1(requestConfigurationArchiveExportDeviceConfigurationsV1 *RequestConfigurationArchiveExportDeviceConfigurationsV1) (*ResponseConfigurationArchiveExportDeviceConfigurationsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device-archive/cleartext"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestConfigurationArchiveExportDeviceConfigurationsV1).
		SetResult(&ResponseConfigurationArchiveExportDeviceConfigurationsV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ExportDeviceConfigurationsV1(requestConfigurationArchiveExportDeviceConfigurationsV1)
		}

		return nil, response, fmt.Errorf("error with operation ExportDeviceConfigurationsV1")
	}

	result := response.Result().(*ResponseConfigurationArchiveExportDeviceConfigurationsV1)
	return result, response, err

}

//DownloadMaskedDeviceConfigurationV1 Download masked device configuration - fe93-185d-4c58-a302
/* Download the masked (sanitized) device configuration by providing the file `id`.


@param id id path parameter. The value of `id` can be obtained from the response of API `/dna/intent/api/v1/networkDeviceConfigFiles`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!download-masked-device-configuration
*/
func (s *ConfigurationArchiveService) DownloadMaskedDeviceConfigurationV1(id string) (*ResponseConfigurationArchiveDownloadMaskedDeviceConfigurationV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDeviceConfigFiles/{id}/downloadMasked"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").

		// SetResult(&ResponseConfigurationArchiveDownloadMaskedDeviceConfigurationV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.DownloadMaskedDeviceConfigurationV1(id)
		}

		return nil, response, fmt.Errorf("error with operation DownloadMaskedDeviceConfigurationV1")
	}

	result := response.Result().(ResponseConfigurationArchiveDownloadMaskedDeviceConfigurationV1)

	return &result, response, err

}

//DownloadUnmaskedrawDeviceConfigurationAsZIPV1 Download Unmasked (raw) Device Configuration as ZIP - 59a7-7a49-4e79-8fde
/* Download the unmasked (raw) device configuration by providing the file `id` and a `password`. The response will be a password-protected zip file containing the unmasked configuration. Password must contain a minimum of 8 characters, one lowercase letter, one uppercase letter, one number, one special character (`-=[];,./~!@#$%^&*()_+{}|:?`). It may not contain white space or the characters `<>`.


@param id id path parameter. The value of `id` can be obtained from the response of API `/dna/intent/api/v1/networkDeviceConfigFiles`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!download-unmaskedraw-device-configuration-as-z-ip
*/
func (s *ConfigurationArchiveService) DownloadUnmaskedrawDeviceConfigurationAsZIPV1(id string, requestConfigurationArchiveDownloadUnmaskedrawDeviceConfigurationAsZIPV1 *RequestConfigurationArchiveDownloadUnmaskedrawDeviceConfigurationAsZIPV1) (*ResponseConfigurationArchiveDownloadUnmaskedrawDeviceConfigurationAsZIPV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDeviceConfigFiles/{id}/downloadUnmasked"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestConfigurationArchiveDownloadUnmaskedrawDeviceConfigurationAsZIPV1).
		// SetResult(&ResponseConfigurationArchiveDownloadUnmaskedrawDeviceConfigurationAsZIPV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.DownloadUnmaskedrawDeviceConfigurationAsZIPV1(id, requestConfigurationArchiveDownloadUnmaskedrawDeviceConfigurationAsZIPV1)
		}

		return nil, response, fmt.Errorf("error with operation DownloadUnmaskedrawDeviceConfigurationAsZIpV1")
	}

	result := response.Result().(ResponseConfigurationArchiveDownloadUnmaskedrawDeviceConfigurationAsZIPV1)

	return &result, response, err

}

// Alias Function
/*
This method acts as an alias for the method `ExportDeviceConfigurationsV1`
*/
func (s *ConfigurationArchiveService) ExportDeviceConfigurations(requestConfigurationArchiveExportDeviceConfigurationsV1 *RequestConfigurationArchiveExportDeviceConfigurationsV1) (*ResponseConfigurationArchiveExportDeviceConfigurationsV1, *resty.Response, error) {
	return s.ExportDeviceConfigurationsV1(requestConfigurationArchiveExportDeviceConfigurationsV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetNetworkDeviceConfigurationFileDetailsV1`
*/
func (s *ConfigurationArchiveService) GetNetworkDeviceConfigurationFileDetails(GetNetworkDeviceConfigurationFileDetailsV1QueryParams *GetNetworkDeviceConfigurationFileDetailsV1QueryParams) (*ResponseConfigurationArchiveGetNetworkDeviceConfigurationFileDetailsV1, *resty.Response, error) {
	return s.GetNetworkDeviceConfigurationFileDetailsV1(GetNetworkDeviceConfigurationFileDetailsV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `CountOfNetworkDeviceConfigurationFilesV1`
*/
func (s *ConfigurationArchiveService) CountOfNetworkDeviceConfigurationFiles(CountOfNetworkDeviceConfigurationFilesV1QueryParams *CountOfNetworkDeviceConfigurationFilesV1QueryParams) (*ResponseConfigurationArchiveCountOfNetworkDeviceConfigurationFilesV1, *resty.Response, error) {
	return s.CountOfNetworkDeviceConfigurationFilesV1(CountOfNetworkDeviceConfigurationFilesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetConfigurationArchiveDetailsV1`
*/
func (s *ConfigurationArchiveService) GetConfigurationArchiveDetails(GetConfigurationArchiveDetailsV1QueryParams *GetConfigurationArchiveDetailsV1QueryParams) (*ResponseConfigurationArchiveGetConfigurationArchiveDetailsV1, *resty.Response, error) {
	return s.GetConfigurationArchiveDetailsV1(GetConfigurationArchiveDetailsV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetConfigurationFileDetailsByIDV1`
*/
func (s *ConfigurationArchiveService) GetConfigurationFileDetailsByID(id string) (*ResponseConfigurationArchiveGetConfigurationFileDetailsByIDV1, *resty.Response, error) {
	return s.GetConfigurationFileDetailsByIDV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `DownloadMaskedDeviceConfigurationV1`
*/
func (s *ConfigurationArchiveService) DownloadMaskedDeviceConfiguration(id string) (*ResponseConfigurationArchiveDownloadMaskedDeviceConfigurationV1, *resty.Response, error) {
	return s.DownloadMaskedDeviceConfigurationV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `DownloadUnmaskedrawDeviceConfigurationAsZIPV1`
*/
func (s *ConfigurationArchiveService) DownloadUnmaskedrawDeviceConfigurationAsZIP(id string, requestConfigurationArchiveDownloadUnmaskedrawDeviceConfigurationAsZIPV1 *RequestConfigurationArchiveDownloadUnmaskedrawDeviceConfigurationAsZIPV1) (*ResponseConfigurationArchiveDownloadUnmaskedrawDeviceConfigurationAsZIPV1, *resty.Response, error) {
	return s.DownloadUnmaskedrawDeviceConfigurationAsZIPV1(id, requestConfigurationArchiveDownloadUnmaskedrawDeviceConfigurationAsZIPV1)
}
