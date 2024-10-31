package catalyst

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
)

type SecurityAdvisoriesService service

type ResponseSecurityAdvisoriesGetAdvisoriesListV1 struct {
	Response *ResponseSecurityAdvisoriesGetAdvisoriesListV1Response `json:"response,omitempty"` //
	Version  string                                                 `json:"version,omitempty"`  // Version of the response
}
type ResponseSecurityAdvisoriesGetAdvisoriesListV1Response struct {
	AdvisoryID                string                                                              `json:"advisoryId,omitempty"`                // Id of the advisory
	DeviceCount               *int                                                                `json:"deviceCount,omitempty"`               // Number of devices vulnerable to the advisory
	HiddenDeviceCount         *int                                                                `json:"hiddenDeviceCount,omitempty"`         // Number of devices vulnerable to the advisory but were suppressed by the user
	Cves                      []string                                                            `json:"cves,omitempty"`                      // CVE (Common Vulnerabilities and Exposures) IDs of the advisory
	PublicationURL            string                                                              `json:"publicationUrl,omitempty"`            // CISCO publication URL for the advisory
	Sir                       string                                                              `json:"sir,omitempty"`                       // Security Impact Rating of the advisory
	DetectionType             string                                                              `json:"detectionType,omitempty"`             // Criteria for advisory detection
	DefaultDetectionType      string                                                              `json:"defaultDetectionType,omitempty"`      // Original criteria for advisory detection
	DefaultConfigMatchPattern string                                                              `json:"defaultConfigMatchPattern,omitempty"` // Regular expression used by the system to detect the advisory
	FixedVersions             *ResponseSecurityAdvisoriesGetAdvisoriesListV1ResponseFixedVersions `json:"fixedVersions,omitempty"`             // Map where each key is a vulnerable version and the value is a list of versions in which the advisory has been fixed
}
type ResponseSecurityAdvisoriesGetAdvisoriesListV1ResponseFixedVersions interface{}
type ResponseSecurityAdvisoriesGetAdvisoriesSummaryV1 struct {
	Response *ResponseSecurityAdvisoriesGetAdvisoriesSummaryV1Response `json:"response,omitempty"` //
	Version  string                                                    `json:"version,omitempty"`  // Version of the response
}
type ResponseSecurityAdvisoriesGetAdvisoriesSummaryV1Response struct {
	INFORMATIONAL *ResponseSecurityAdvisoriesGetAdvisoriesSummaryV1ResponseINFORMATIONAL `json:"INFORMATIONAL,omitempty"` //
	LOW           *ResponseSecurityAdvisoriesGetAdvisoriesSummaryV1ResponseLOW           `json:"LOW,omitempty"`           //
	MEDIUM        *ResponseSecurityAdvisoriesGetAdvisoriesSummaryV1ResponseMEDIUM        `json:"MEDIUM,omitempty"`        //
	HIGH          *ResponseSecurityAdvisoriesGetAdvisoriesSummaryV1ResponseHIGH          `json:"HIGH,omitempty"`          //
	CRITICaL      *ResponseSecurityAdvisoriesGetAdvisoriesSummaryV1ResponseCRITICaL      `json:"CRITICAL,omitempty"`      //
	NA            *ResponseSecurityAdvisoriesGetAdvisoriesSummaryV1ResponseNA            `json:"NA,omitempty"`            //
}
type ResponseSecurityAdvisoriesGetAdvisoriesSummaryV1ResponseINFORMATIONAL struct {
	CONFIG       *int `json:"CONFIG,omitempty"`        // Number of advisories matched using default config
	CUSTOMCONFIG *int `json:"CUSTOM_CONFIG,omitempty"` // Number of advisories matched using user provided config
	VERSION      *int `json:"VERSION,omitempty"`       // Number of advisories matched using software version
	TOTAL        *int `json:"TOTAL,omitempty"`         // Sum of Config, Custom Config and Version
}
type ResponseSecurityAdvisoriesGetAdvisoriesSummaryV1ResponseLOW struct {
	CONFIG       *int `json:"CONFIG,omitempty"`        // Number of advisories matched using default config
	CUSTOMCONFIG *int `json:"CUSTOM_CONFIG,omitempty"` // Number of advisories matched using user provided config
	VERSION      *int `json:"VERSION,omitempty"`       // Number of advisories matched using software version
	TOTAL        *int `json:"TOTAL,omitempty"`         // Sum of Config, Custom Config and Version
}
type ResponseSecurityAdvisoriesGetAdvisoriesSummaryV1ResponseMEDIUM struct {
	CONFIG       *int `json:"CONFIG,omitempty"`        // Number of advisories matched using default config
	CUSTOMCONFIG *int `json:"CUSTOM_CONFIG,omitempty"` // Number of advisories matched using user provided config
	VERSION      *int `json:"VERSION,omitempty"`       // Number of advisories matched using software version
	TOTAL        *int `json:"TOTAL,omitempty"`         // Sum of Config, Custom Config and Version
}
type ResponseSecurityAdvisoriesGetAdvisoriesSummaryV1ResponseHIGH struct {
	CONFIG       *int `json:"CONFIG,omitempty"`        // Number of advisories matched using default config
	CUSTOMCONFIG *int `json:"CUSTOM_CONFIG,omitempty"` // Number of advisories matched using user provided config
	VERSION      *int `json:"VERSION,omitempty"`       // Number of advisories matched using software version
	TOTAL        *int `json:"TOTAL,omitempty"`         // Sum of Config, Custom Config and Version
}
type ResponseSecurityAdvisoriesGetAdvisoriesSummaryV1ResponseCRITICaL struct {
	CONFIG       *int `json:"CONFIG,omitempty"`        // Number of advisories matched using default config
	CUSTOMCONFIG *int `json:"CUSTOM_CONFIG,omitempty"` // Number of advisories matched using user provided config
	VERSION      *int `json:"VERSION,omitempty"`       // Number of advisories matched using software version
	TOTAL        *int `json:"TOTAL,omitempty"`         // Sum of Config, Custom Config and Version
}
type ResponseSecurityAdvisoriesGetAdvisoriesSummaryV1ResponseNA struct {
	CONFIG       *int `json:"CONFIG,omitempty"`        // Number of advisories matched using default config
	CUSTOMCONFIG *int `json:"CUSTOM_CONFIG,omitempty"` // Number of advisories matched using user provided config
	VERSION      *int `json:"VERSION,omitempty"`       // Number of advisories matched using software version
	TOTAL        *int `json:"TOTAL,omitempty"`         // Sum of Config, Custom Config and Version
}
type ResponseSecurityAdvisoriesGetDevicesPerAdvisoryV1 struct {
	Response []string `json:"response,omitempty"` // List of device IDs vulnerable to the advisory
	Version  string   `json:"version,omitempty"`  // Version of the response
}
type ResponseSecurityAdvisoriesGetAdvisoryDeviceDetailV1 struct {
	Response *ResponseSecurityAdvisoriesGetAdvisoryDeviceDetailV1Response `json:"response,omitempty"` //
	Version  string                                                       `json:"version,omitempty"`  // Version of the response
}
type ResponseSecurityAdvisoriesGetAdvisoryDeviceDetailV1Response struct {
	DeviceID            string   `json:"deviceId,omitempty"`            // Network device ID
	AdvisoryIDs         []string `json:"advisoryIds,omitempty"`         // Advisories detected on the network device
	HiddenAdvisoryCount *int     `json:"hiddenAdvisoryCount,omitempty"` // Number of advisories detected on the network device that were suppressed by the user
	ScanMode            string   `json:"scanMode,omitempty"`            // Criteria on which the network device was scanned
	ScanStatus          string   `json:"scanStatus,omitempty"`          // Status of the scan performed on the network device
	Comments            string   `json:"comments,omitempty"`            // More details about the scan status. Ie:- if the scan status is failed, comments will give the reason for failure
	LastScanTime        *int     `json:"lastScanTime,omitempty"`        // Time at which the network device was scanned. The representation is unix time.
}
type ResponseSecurityAdvisoriesGetAdvisoriesPerDeviceV1 struct {
	Response *ResponseSecurityAdvisoriesGetAdvisoriesPerDeviceV1Response `json:"response,omitempty"` //
	Version  string                                                      `json:"version,omitempty"`  // Version of the response
}
type ResponseSecurityAdvisoriesGetAdvisoriesPerDeviceV1Response struct {
	AdvisoryID                string                                                                   `json:"advisoryId,omitempty"`                // Id of the advisory
	DeviceCount               *int                                                                     `json:"deviceCount,omitempty"`               // Number of devices vulnerable to the advisory
	HiddenDeviceCount         *int                                                                     `json:"hiddenDeviceCount,omitempty"`         // Number of devices vulnerable to the advisory but were suppressed by the user
	Cves                      []string                                                                 `json:"cves,omitempty"`                      // CVE (Common Vulnerabilities and Exposures) IDs of the advisory
	PublicationURL            string                                                                   `json:"publicationUrl,omitempty"`            // CISCO publication URL for the advisory
	Sir                       string                                                                   `json:"sir,omitempty"`                       // Security Impact Rating of the advisory
	DetectionType             string                                                                   `json:"detectionType,omitempty"`             // Criteria for advisory detection
	DefaultDetectionType      string                                                                   `json:"defaultDetectionType,omitempty"`      // Original criteria for advisory detection
	DefaultConfigMatchPattern string                                                                   `json:"defaultConfigMatchPattern,omitempty"` // Regular expression used by the system to detect the advisory
	FixedVersions             *ResponseSecurityAdvisoriesGetAdvisoriesPerDeviceV1ResponseFixedVersions `json:"fixedVersions,omitempty"`             // Map where each key is a vulnerable version and the value is a list of versions in which the advisory has been fixed
}
type ResponseSecurityAdvisoriesGetAdvisoriesPerDeviceV1ResponseFixedVersions interface{}

//GetAdvisoriesListV1 Get Advisories List - 4295-0bf8-4939-ac35
/* Retrieves list of advisories on the network



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-advisories-list-v1
*/
func (s *SecurityAdvisoriesService) GetAdvisoriesListV1() (*ResponseSecurityAdvisoriesGetAdvisoriesListV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/security-advisory/advisory"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSecurityAdvisoriesGetAdvisoriesListV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAdvisoriesListV1()
		}
		return nil, response, fmt.Errorf("error with operation GetAdvisoriesListV1")
	}

	result := response.Result().(*ResponseSecurityAdvisoriesGetAdvisoriesListV1)
	return result, response, err

}

//GetAdvisoriesSummaryV1 Get Advisories Summary - 3ebf-898d-482b-9207
/* Retrieves summary of advisories on the network.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-advisories-summary-v1
*/
func (s *SecurityAdvisoriesService) GetAdvisoriesSummaryV1() (*ResponseSecurityAdvisoriesGetAdvisoriesSummaryV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/security-advisory/advisory/aggregate"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSecurityAdvisoriesGetAdvisoriesSummaryV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAdvisoriesSummaryV1()
		}
		return nil, response, fmt.Errorf("error with operation GetAdvisoriesSummaryV1")
	}

	result := response.Result().(*ResponseSecurityAdvisoriesGetAdvisoriesSummaryV1)
	return result, response, err

}

//GetDevicesPerAdvisoryV1 Get Devices Per Advisory - f49c-4ae0-43fa-8352
/* Retrieves list of devices for an advisory


@param advisoryID advisoryId path parameter. Advisory ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-devices-per-advisory-v1
*/
func (s *SecurityAdvisoriesService) GetDevicesPerAdvisoryV1(advisoryID string) (*ResponseSecurityAdvisoriesGetDevicesPerAdvisoryV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/security-advisory/advisory/{advisoryId}/device"
	path = strings.Replace(path, "{advisoryId}", fmt.Sprintf("%v", advisoryID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSecurityAdvisoriesGetDevicesPerAdvisoryV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDevicesPerAdvisoryV1(advisoryID)
		}
		return nil, response, fmt.Errorf("error with operation GetDevicesPerAdvisoryV1")
	}

	result := response.Result().(*ResponseSecurityAdvisoriesGetDevicesPerAdvisoryV1)
	return result, response, err

}

//GetAdvisoryDeviceDetailV1 Get Advisory Device Detail - e295-09d0-420b-8cc4
/* Retrieves advisory device details for a device


@param deviceID deviceId path parameter. Device instance UUID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-advisory-device-detail-v1
*/
func (s *SecurityAdvisoriesService) GetAdvisoryDeviceDetailV1(deviceID string) (*ResponseSecurityAdvisoriesGetAdvisoryDeviceDetailV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/security-advisory/device/{deviceId}"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSecurityAdvisoriesGetAdvisoryDeviceDetailV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAdvisoryDeviceDetailV1(deviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetAdvisoryDeviceDetailV1")
	}

	result := response.Result().(*ResponseSecurityAdvisoriesGetAdvisoryDeviceDetailV1)
	return result, response, err

}

//GetAdvisoriesPerDeviceV1 Get Advisories Per Device - 42a6-c9a1-4ea9-b002
/* Retrieves list of advisories for a device


@param deviceID deviceId path parameter. Device instance UUID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-advisories-per-device-v1
*/
func (s *SecurityAdvisoriesService) GetAdvisoriesPerDeviceV1(deviceID string) (*ResponseSecurityAdvisoriesGetAdvisoriesPerDeviceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/security-advisory/device/{deviceId}/advisory"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSecurityAdvisoriesGetAdvisoriesPerDeviceV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAdvisoriesPerDeviceV1(deviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetAdvisoriesPerDeviceV1")
	}

	result := response.Result().(*ResponseSecurityAdvisoriesGetAdvisoriesPerDeviceV1)
	return result, response, err

}

// Alias Function
func (s *SecurityAdvisoriesService) GetAdvisoriesSummary() (*ResponseSecurityAdvisoriesGetAdvisoriesSummaryV1, *resty.Response, error) {
	return s.GetAdvisoriesSummaryV1()
}

// Alias Function
func (s *SecurityAdvisoriesService) GetAdvisoryDeviceDetail(deviceID string) (*ResponseSecurityAdvisoriesGetAdvisoryDeviceDetailV1, *resty.Response, error) {
	return s.GetAdvisoryDeviceDetailV1(deviceID)
}

// Alias Function
func (s *SecurityAdvisoriesService) GetDevicesPerAdvisory(advisoryID string) (*ResponseSecurityAdvisoriesGetDevicesPerAdvisoryV1, *resty.Response, error) {
	return s.GetDevicesPerAdvisoryV1(advisoryID)
}

// Alias Function
func (s *SecurityAdvisoriesService) GetAdvisoriesPerDevice(deviceID string) (*ResponseSecurityAdvisoriesGetAdvisoriesPerDeviceV1, *resty.Response, error) {
	return s.GetAdvisoriesPerDeviceV1(deviceID)
}

// Alias Function
func (s *SecurityAdvisoriesService) GetAdvisoriesList() (*ResponseSecurityAdvisoriesGetAdvisoriesListV1, *resty.Response, error) {
	return s.GetAdvisoriesListV1()
}
