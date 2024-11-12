package catalyst

import (
	"fmt"
	"net/http"

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
	Limit       float64 `url:"limit,omitempty"`       //limit
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
type RequestConfigurationArchiveExportDeviceConfigurationsV1 struct {
	Password string `json:"password,omitempty"` // Password for the zip file to protect exported configurations. Must contain, at minimum 8 characters, one lowercase letter, one uppercase letter, one number, one special character(-=[];,./~!@#$%^&*()_+{}|:?). It may not contain white space or the characters <>.
	DeviceID string `json:"deviceId,omitempty"` // UUIDs of the devices for which configurations need to be exported.
}

//GetConfigurationArchiveDetailsV1 Get configuration archive details - 3bba-48a9-422a-be1e
/* Returns the historical device configurations (running configuration , startup configuration , vlan if applicable) by specified criteria


@param GetConfigurationArchiveDetailsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-configuration-archive-details-v1
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

//ExportDeviceConfigurationsV1 Export Device configurations - 51a4-0aba-4c68-ac17
/* Export Device configurations to an encrypted zip file



Documentation Link: https://developer.cisco.com/docs/dna-center/#!export-device-configurations-v1
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

// Alias Function
/*
This method acts as an alias for the method `ExportDeviceConfigurationsV1`
*/
func (s *ConfigurationArchiveService) ExportDeviceConfigurations(requestConfigurationArchiveExportDeviceConfigurationsV1 *RequestConfigurationArchiveExportDeviceConfigurationsV1) (*ResponseConfigurationArchiveExportDeviceConfigurationsV1, *resty.Response, error) {
	return s.ExportDeviceConfigurationsV1(requestConfigurationArchiveExportDeviceConfigurationsV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetConfigurationArchiveDetailsV1`
*/
func (s *ConfigurationArchiveService) GetConfigurationArchiveDetails(GetConfigurationArchiveDetailsV1QueryParams *GetConfigurationArchiveDetailsV1QueryParams) (*ResponseConfigurationArchiveGetConfigurationArchiveDetailsV1, *resty.Response, error) {
	return s.GetConfigurationArchiveDetailsV1(GetConfigurationArchiveDetailsV1QueryParams)
}
