package catalyst

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
)

type EoxService service

type ResponseEoxGetEoxStatusForAllDevicesV1 struct {
	Response *[]ResponseEoxGetEoxStatusForAllDevicesV1Response `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  // Version of the response
}
type ResponseEoxGetEoxStatusForAllDevicesV1Response struct {
	DeviceID     string                                                   `json:"deviceId,omitempty"`     // Device instance UUID
	AlertCount   *int                                                     `json:"alertCount,omitempty"`   // Number of EoX alerts on the network device
	Summary      *[]ResponseEoxGetEoxStatusForAllDevicesV1ResponseSummary `json:"summary,omitempty"`      //
	ScanStatus   string                                                   `json:"scanStatus,omitempty"`   // Status of the scan performed on the network device
	Comments     []string                                                 `json:"comments,omitempty"`     // More details about the scan status. ie:- if the scan status is failed, comments will give the reason for failure.
	LastScanTime *int                                                     `json:"lastScanTime,omitempty"` // Time at which the network device was scanned. The representation is unix time.
}
type ResponseEoxGetEoxStatusForAllDevicesV1ResponseSummary struct {
	EoxType string `json:"eoxType,omitempty"` // Type of EoX Alert
}
type ResponseEoxGetEoxDetailsPerDeviceV1 struct {
	Response *ResponseEoxGetEoxDetailsPerDeviceV1Response `json:"response,omitempty"` //
	Version  string                                       `json:"version,omitempty"`  // Version of the response
}
type ResponseEoxGetEoxDetailsPerDeviceV1Response struct {
	DeviceID     string                                                   `json:"deviceId,omitempty"`     // Device instance UUID
	AlertCount   *int                                                     `json:"alertCount,omitempty"`   // Number of EoX alerts on the network device
	EoxDetails   *[]ResponseEoxGetEoxDetailsPerDeviceV1ResponseEoxDetails `json:"eoxDetails,omitempty"`   //
	ScanStatus   string                                                   `json:"scanStatus,omitempty"`   // Status of the scan performed on the network device
	Comments     []string                                                 `json:"comments,omitempty"`     // More details about the scan status. ie:- if the scan status is failed, comments will give the reason for failure.
	LastScanTime *int                                                     `json:"lastScanTime,omitempty"` // Time at which the network device was scanned. The representation is unix time.
}
type ResponseEoxGetEoxDetailsPerDeviceV1ResponseEoxDetails struct {
	Name                                              string `json:"name,omitempty"`                                              // Name of the EoX alert. Every EoX announcement has a unique name. ie:- EOL13873
	BulletinHeadline                                  string `json:"bulletinHeadline,omitempty"`                                  // Title of the EoX bulletin
	BulletinName                                      string `json:"bulletinName,omitempty"`                                      // Name of the EoX bulletin
	BulletinNumber                                    string `json:"bulletinNumber,omitempty"`                                    // Identifier of the EoX bulletin. Usually the same as name.
	BulletinURL                                       string `json:"bulletinURL,omitempty"`                                       // URL where the EoX bulletin is posted
	EndOfHardwareNewServiceAttachmentDate             string `json:"endOfHardwareNewServiceAttachmentDate,omitempty"`             // For equipment and software that is not covered by a service-and-support contract, this is the last date to order a new service-and-support contract or add the equipment and/or software to an existing service-and-support contract
	EndOfHardwareServiceContractRenewalDate           string `json:"endOfHardwareServiceContractRenewalDate,omitempty"`           // The last date to extend or renew a service contract for the product
	EndOfLastHardwareShipDate                         string `json:"endOfLastHardwareShipDate,omitempty"`                         // The last-possible ship date that can be requested of Cisco and/or its contract manufacturers
	EndOfLifeExternalAnnouncementDate                 string `json:"endOfLifeExternalAnnouncementDate,omitempty"`                 // The date the document that announces the end-of-sale and end-of-life of a product is distributed to the general public
	EndOfSignatureReleasesDate                        string `json:"endOfSignatureReleasesDate,omitempty"`                        // The date after which there will be no more signature update release for the product
	EndOfSoftwareVulnerabilityOrSecuritySupportDate   string `json:"endOfSoftwareVulnerabilityOrSecuritySupportDate,omitempty"`   // The last date that Cisco Engineering may release bug fixes for Vulnerability or Security issues for the product. This will be populated for software alerts only.
	EndOfSoftwareVulnerabilityOrSecuritySupportDateHw string `json:"endOfSoftwareVulnerabilityOrSecuritySupportDateHw,omitempty"` // The last date that Cisco Engineering may release bug fixes for Vulnerability or Security issues for the product. This will be populated for hardware or module alerts only.
	EndOfSaleDate                                     string `json:"endOfSaleDate,omitempty"`                                     // The last date to order the product through Cisco point-of-sale mechanisms
	EndOfLifeDate                                     string `json:"endOfLifeDate,omitempty"`                                     // The last date to receive applicable service and support for the product as entitled by active service contracts or by warranty terms and conditions. This will be populated for software alerts only.
	LastDateOfSupport                                 string `json:"lastDateOfSupport,omitempty"`                                 // The last date to receive applicable service and support for the product as entitled by active service contracts or by warranty terms and conditions. This will be populated for hardware and module alerts only.
	EndOfSoftwareMaintenanceReleasesDate              string `json:"endOfSoftwareMaintenanceReleasesDate,omitempty"`              // The last date that Cisco Engineering may release any final software maintenance releases or bug fixes for the product
	EoxAlertType                                      string `json:"eoxAlertType,omitempty"`                                      // Type of EoX alert
	EoxPhysicalType                                   string `json:"eoXPhysicalType,omitempty"`                                   // The type of part for EoX alert. eg:- Power Supply, Chassis, Fan etc.
	BulletinPID                                       string `json:"bulletinPID,omitempty"`                                       // The part number for the EoX alert. eg:- PWR-C1-1100WAC
}
type ResponseEoxGetEoxSummaryV1 struct {
	Response *ResponseEoxGetEoxSummaryV1Response `json:"response,omitempty"` //
	Version  string                              `json:"version,omitempty"`  // Version of the response
}
type ResponseEoxGetEoxSummaryV1Response struct {
	HardwareCount *int `json:"hardwareCount,omitempty"` // Number of hardware EoX alerts detected on the network
	SoftwareCount *int `json:"softwareCount,omitempty"` // Number of software EoX alerts detected on the network
	ModuleCount   *int `json:"moduleCount,omitempty"`   // Number of module EoX alerts detected on the network
	TotalCount    *int `json:"totalCount,omitempty"`    // Total number of EoX alerts detected on the network. This is the sum of hardwareCount, softwareCount and moduleCount.
}

//GetEoxStatusForAllDevicesV1 Get EoX Status For All Devices - 3281-fa04-49ba-87d9
/* Retrieves EoX status for all devices in the network



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-eox-status-for-all-devices-v1
*/
func (s *EoxService) GetEoxStatusForAllDevicesV1() (*ResponseEoxGetEoxStatusForAllDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/eox-status/device"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseEoxGetEoxStatusForAllDevicesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetEoxStatusForAllDevicesV1()
		}
		return nil, response, fmt.Errorf("error with operation GetEoxStatusForAllDevicesV1")
	}

	result := response.Result().(*ResponseEoxGetEoxStatusForAllDevicesV1)
	return result, response, err

}

//GetEoxDetailsPerDeviceV1 Get EoX Details Per Device - dc80-099e-4d59-986d
/* Retrieves EoX details for a device


@param deviceID deviceId path parameter. Device instance UUID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-eox-details-per-device-v1
*/
func (s *EoxService) GetEoxDetailsPerDeviceV1(deviceID string) (*ResponseEoxGetEoxDetailsPerDeviceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/eox-status/device/{deviceId}"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseEoxGetEoxDetailsPerDeviceV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetEoxDetailsPerDeviceV1(deviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetEoxDetailsPerDeviceV1")
	}

	result := response.Result().(*ResponseEoxGetEoxDetailsPerDeviceV1)
	return result, response, err

}

//GetEoxSummaryV1 Get EoX Summary - f0b2-7a23-4fea-96fc
/* Retrieves EoX summary for all devices in the network



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-eox-summary-v1
*/
func (s *EoxService) GetEoxSummaryV1() (*ResponseEoxGetEoxSummaryV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/eox-status/summary"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseEoxGetEoxSummaryV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetEoxSummaryV1()
		}
		return nil, response, fmt.Errorf("error with operation GetEoxSummaryV1")
	}

	result := response.Result().(*ResponseEoxGetEoxSummaryV1)
	return result, response, err

}

// Alias Function
func (s *EoxService) GetEoxDetailsPerDevice(deviceID string) (*ResponseEoxGetEoxDetailsPerDeviceV1, *resty.Response, error) {
	return s.GetEoxDetailsPerDeviceV1(deviceID)
}

// Alias Function
func (s *EoxService) GetEoxStatusForAllDevices() (*ResponseEoxGetEoxStatusForAllDevicesV1, *resty.Response, error) {
	return s.GetEoxStatusForAllDevicesV1()
}

// Alias Function
func (s *EoxService) GetEoxSummary() (*ResponseEoxGetEoxSummaryV1, *resty.Response, error) {
	return s.GetEoxSummaryV1()
}
