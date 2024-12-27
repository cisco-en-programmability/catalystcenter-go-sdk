package catalyst

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type FabricWirelessService service

type GetSSIDToIPPoolMappingV1QueryParams struct {
	VLANName          string `url:"vlanName,omitempty"`          //VLAN Name
	SiteNameHierarchy string `url:"siteNameHierarchy,omitempty"` //Site Name Heirarchy
}
type RemoveWLCFromFabricDomainV1QueryParams struct {
	DeviceIPAddress string `url:"deviceIPAddress,omitempty"` //Device Management IP Address
}
type RemoveWLCFromFabricDomainV1HeaderParams struct {
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type string. Enable this parameter to execute the API and return a response asynchronously.
}
type ReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingV1QueryParams struct {
	Limit  float64 `url:"limit,omitempty"`  //Return only this many IP Pool to SSID Mapping. Default is 500 if not specified. Maximum allowed limit is 500.
	Offset float64 `url:"offset,omitempty"` //Number of records to skip for pagination
}
type RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteV1QueryParams struct {
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500.
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1.
}

type ResponseFabricWirelessAddSSIDToIPPoolMappingV1 struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseFabricWirelessUpdateSSIDToIPPoolMappingV1 struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseFabricWirelessGetSSIDToIPPoolMappingV1 struct {
	VLANName    string                                                       `json:"vlanName,omitempty"`    // VLAN Name
	SSIDDetails *[]ResponseFabricWirelessGetSSIDToIPPoolMappingV1SSIDDetails `json:"ssidDetails,omitempty"` //
}
type ResponseFabricWirelessGetSSIDToIPPoolMappingV1SSIDDetails struct {
	Name              string `json:"name,omitempty"`              // SSID Name
	ScalableGroupName string `json:"scalableGroupName,omitempty"` // Scalable Group Name
}
type ResponseFabricWirelessRemoveWLCFromFabricDomainV1 struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseFabricWirelessAddWLCToFabricDomainV1 struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseFabricWirelessReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingV1 struct {
	Response *[]ResponseFabricWirelessReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingV1Response `json:"response,omitempty"` //
	Version  string                                                                               `json:"version,omitempty"`  // Version
}
type ResponseFabricWirelessReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingV1Response struct {
	FabricID    string                                                                                          `json:"fabricId,omitempty"`    // Fabric Id
	VLANDetails *[]ResponseFabricWirelessReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingV1ResponseVLANDetails `json:"vlanDetails,omitempty"` //
}
type ResponseFabricWirelessReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingV1ResponseVLANDetails struct {
	VLANName    string                                                                                                     `json:"vlanName,omitempty"`    // Vlan Name
	SSIDDetails *[]ResponseFabricWirelessReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingV1ResponseVLANDetailsSSIDDetails `json:"ssidDetails,omitempty"` //
}
type ResponseFabricWirelessReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingV1ResponseVLANDetailsSSIDDetails struct {
	Name             string `json:"name,omitempty"`             // Name of the SSID.
	SecurityGroupTag string `json:"securityGroupTag,omitempty"` // Represents the name of the Security Group. Example: Auditors, BYOD, Developers, etc.
}
type ResponseFabricWirelessReturnTheCountOfAllTheFabricSiteWhichHasSSIDToIPPoolMappingV1 struct {
	Response *ResponseFabricWirelessReturnTheCountOfAllTheFabricSiteWhichHasSSIDToIPPoolMappingV1Response `json:"response,omitempty"` //
	Version  string                                                                                       `json:"version,omitempty"`  // Response Version
}
type ResponseFabricWirelessReturnTheCountOfAllTheFabricSiteWhichHasSSIDToIPPoolMappingV1Response struct {
	Count *int `json:"count,omitempty"` // Return the count of all the fabric site which has SSID to IP Pool mapping
}
type ResponseFabricWirelessSwitchWirelessSettingAndRollingApUpgradeManagementV1 struct {
	Response *ResponseFabricWirelessSwitchWirelessSettingAndRollingApUpgradeManagementV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version of the response
}
type ResponseFabricWirelessSwitchWirelessSettingAndRollingApUpgradeManagementV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Unique identifier for the task

	URL string `json:"url,omitempty"` // URL for the task
}
type ResponseFabricWirelessGetSdaWirelessDetailsFromSwitchesV1 struct {
	Response *[]ResponseFabricWirelessGetSdaWirelessDetailsFromSwitchesV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseFabricWirelessGetSdaWirelessDetailsFromSwitchesV1Response struct {
	ID string `json:"id,omitempty"` // Network Device ID of the Wireless Capable Switch

	EnableWireless *bool `json:"enableWireless,omitempty"` // Enable Wireless

	RollingApUpgrade *ResponseFabricWirelessGetSdaWirelessDetailsFromSwitchesV1ResponseRollingApUpgrade `json:"rollingApUpgrade,omitempty"` //
}
type ResponseFabricWirelessGetSdaWirelessDetailsFromSwitchesV1ResponseRollingApUpgrade struct {
	EnableRollingApUpgrade *bool `json:"enableRollingApUpgrade,omitempty"` // Enable Rolling Ap Upgrade

	ApRebootPercentage *int `json:"apRebootPercentage,omitempty"` // AP Reboot Percentage. Permissible values - 5, 15, 25
}
type ResponseFabricWirelessReloadSwitchForWirelessControllerCleanupV1 struct {
	Response *ResponseFabricWirelessReloadSwitchForWirelessControllerCleanupV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseFabricWirelessReloadSwitchForWirelessControllerCleanupV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLANV1 struct {
	Response *ResponseFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLANV1Response `json:"response,omitempty"` //
	Version  string                                                               `json:"version,omitempty"`  // Version
}
type ResponseFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLANV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID
	URL    string `json:"url,omitempty"`    // Task URL
}
type ResponseFabricWirelessRetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteV1 struct {
	Response *[]ResponseFabricWirelessRetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteV1Response `json:"response,omitempty"` //
	Version  string                                                                                      `json:"version,omitempty"`  // Version
}
type ResponseFabricWirelessRetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteV1Response struct {
	VLANName    string                                                                                                 `json:"vlanName,omitempty"`    // Vlan Name
	SSIDDetails *[]ResponseFabricWirelessRetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteV1ResponseSSIDDetails `json:"ssidDetails,omitempty"` //
}
type ResponseFabricWirelessRetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteV1ResponseSSIDDetails struct {
	Name             string `json:"name,omitempty"`             // Name of the SSID
	SecurityGroupTag string `json:"securityGroupTag,omitempty"` // Represents the name of the Security Group. Example: Auditors, BYOD, Developers, etc.
}
type ResponseFabricWirelessReturnsTheCountOfVLANsMappedToSSIDsInAFabricSiteV1 struct {
	Response *ResponseFabricWirelessReturnsTheCountOfVLANsMappedToSSIDsInAFabricSiteV1Response `json:"response,omitempty"` //
	Version  string                                                                            `json:"version,omitempty"`  // Response Version
}
type ResponseFabricWirelessReturnsTheCountOfVLANsMappedToSSIDsInAFabricSiteV1Response struct {
	Count *int `json:"count,omitempty"` // Returns the count of VLANs mapped to SSIDs in a Fabric Site
}
type ResponseFabricWirelessUpdateSdaWirelessMulticastV1 struct {
	Response *ResponseFabricWirelessUpdateSdaWirelessMulticastV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseFabricWirelessUpdateSdaWirelessMulticastV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseFabricWirelessGetSdaWirelessMulticastV1 struct {
	Response *ResponseFabricWirelessGetSdaWirelessMulticastV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseFabricWirelessGetSdaWirelessMulticastV1Response struct {
	MulticastEnabled *bool `json:"multicastEnabled,omitempty"` // The setting indicates whether multicast is enabled (true) or disabled (false).
}
type RequestFabricWirelessAddSSIDToIPPoolMappingV1 struct {
	VLANName          string   `json:"vlanName,omitempty"`          // VLAN Name
	ScalableGroupName string   `json:"scalableGroupName,omitempty"` // Scalable Group Name
	SSIDNames         []string `json:"ssidNames,omitempty"`         // List of SSIDs
	SiteNameHierarchy string   `json:"siteNameHierarchy,omitempty"` // Site Name Hierarchy
}
type RequestFabricWirelessUpdateSSIDToIPPoolMappingV1 struct {
	VLANName          string   `json:"vlanName,omitempty"`          // VLAN Name
	ScalableGroupName string   `json:"scalableGroupName,omitempty"` // Scalable Group Name
	SSIDNames         []string `json:"ssidNames,omitempty"`         // List of SSIDs
	SiteNameHierarchy string   `json:"siteNameHierarchy,omitempty"` // Site Name Hierarchy
}
type RequestFabricWirelessAddWLCToFabricDomainV1 struct {
	DeviceName        string `json:"deviceName,omitempty"`        // WLC Device Name
	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Fabric Site Name Hierarchy
}
type RequestFabricWirelessSwitchWirelessSettingAndRollingApUpgradeManagementV1 struct {
	ID string `json:"id,omitempty"` // Network Device ID of the wireless capable switch

	EnableWireless *bool `json:"enableWireless,omitempty"` // Enable Wireless

	RollingApUpgrade *RequestFabricWirelessSwitchWirelessSettingAndRollingApUpgradeManagementV1RollingApUpgrade `json:"rollingApUpgrade,omitempty"` //
}
type RequestFabricWirelessSwitchWirelessSettingAndRollingApUpgradeManagementV1RollingApUpgrade struct {
	EnableRollingApUpgrade *bool `json:"enableRollingApUpgrade,omitempty"` // Enable Rolling Ap Upgrade

	ApRebootPercentage *int `json:"apRebootPercentage,omitempty"` // AP Reboot Percentage. Permissible values - 5, 15, 25
}
type RequestFabricWirelessReloadSwitchForWirelessControllerCleanupV1 struct {
	DeviceID string `json:"deviceId,omitempty"` // Network Device ID
}
type RequestFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLANV1 []RequestItemFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLANV1 // Array of RequestFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLANV1
type RequestItemFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLANV1 struct {
	VLANName    string                                                                       `json:"vlanName,omitempty"`    // Vlan Name
	SSIDDetails *[]RequestItemFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLANV1SSIDDetails `json:"ssidDetails,omitempty"` //
}
type RequestItemFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLANV1SSIDDetails struct {
	Name             string `json:"name,omitempty"`             // Name of the SSID
	SecurityGroupTag string `json:"securityGroupTag,omitempty"` // Represents the name of the Security Group. Example: Auditors, BYOD, Developers, etc.
}
type RequestFabricWirelessUpdateSdaWirelessMulticastV1 struct {
	MulticastEnabled *bool `json:"multicastEnabled,omitempty"` // Multicast Enabled
}

//GetSSIDToIPPoolMappingV1 Get SSID to IP Pool Mapping - d891-8a44-4b6a-ad19
/* Get SSID to IP Pool Mapping


@param GetSSIDToIPPoolMappingV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-ssid-to-ip-pool-mapping
*/
func (s *FabricWirelessService) GetSSIDToIPPoolMappingV1(GetSSIDToIPPoolMappingV1QueryParams *GetSSIDToIPPoolMappingV1QueryParams) (*ResponseFabricWirelessGetSSIDToIPPoolMappingV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/hostonboarding/ssid-ippool"

	queryString, _ := query.Values(GetSSIDToIPPoolMappingV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseFabricWirelessGetSSIDToIPPoolMappingV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSSIDToIPPoolMappingV1(GetSSIDToIPPoolMappingV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSsidToIpPoolMappingV1")
	}

	result := response.Result().(*ResponseFabricWirelessGetSSIDToIPPoolMappingV1)
	return result, response, err

}

//ReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingV1 Returns all the Fabric Sites that have VLAN to SSID mapping. - 7a96-98ce-400a-99ce
/* It will return all vlan to SSID mapping across all the fabric site


@param ReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!returns-all-the-fabric-sites-that-have-vlan-to-ssid-mapping
*/
func (s *FabricWirelessService) ReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingV1(ReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingV1QueryParams *ReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingV1QueryParams) (*ResponseFabricWirelessReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabrics/vlanToSsids"

	queryString, _ := query.Values(ReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseFabricWirelessReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingV1(ReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReturnsAllTheFabricSitesThatHaveVlanToSsidMappingV1")
	}

	result := response.Result().(*ResponseFabricWirelessReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingV1)
	return result, response, err

}

//ReturnTheCountOfAllTheFabricSiteWhichHasSSIDToIPPoolMappingV1 Return the count of all the fabric site which has SSID to IP Pool mapping  - 36b0-0b14-44fa-8c4b
/* Return the count of all the fabric site which has SSID to IP Pool mapping



Documentation Link: https://developer.cisco.com/docs/dna-center/#!return-the-count-of-all-the-fabric-site-which-has-ssid-to-ip-pool-mapping
*/
func (s *FabricWirelessService) ReturnTheCountOfAllTheFabricSiteWhichHasSSIDToIPPoolMappingV1() (*ResponseFabricWirelessReturnTheCountOfAllTheFabricSiteWhichHasSSIDToIPPoolMappingV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabrics/vlanToSsids/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseFabricWirelessReturnTheCountOfAllTheFabricSiteWhichHasSSIDToIPPoolMappingV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReturnTheCountOfAllTheFabricSiteWhichHasSSIDToIPPoolMappingV1()
		}
		return nil, response, fmt.Errorf("error with operation ReturnTheCountOfAllTheFabricSiteWhichHasSsidToIpPoolMappingV1")
	}

	result := response.Result().(*ResponseFabricWirelessReturnTheCountOfAllTheFabricSiteWhichHasSSIDToIPPoolMappingV1)
	return result, response, err

}

//GetSdaWirelessDetailsFromSwitchesV1 Get SDA Wireless details from Switches - e48c-0bd1-459b-ad8d
/* Get the SDA Wireless details from the switches on the fabric site that have wireless capability enabled. A maximum of two switches can have a wireless role in a fabric site.


@param fabricID fabricId path parameter. The 'fabricId' represents the Fabric ID of a particular Fabric Site. The 'fabricId' can be obtained using the api /dna/intent/api/v1/sda/fabricSites.  Example : e290f1ee-6c54-4b01-90e6-d701748f0851


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-sda-wireless-details-from-switches
*/
func (s *FabricWirelessService) GetSdaWirelessDetailsFromSwitchesV1(fabricID string) (*ResponseFabricWirelessGetSdaWirelessDetailsFromSwitchesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabrics/{fabricId}/switchWirelessSetting"
	path = strings.Replace(path, "{fabricId}", fmt.Sprintf("%v", fabricID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseFabricWirelessGetSdaWirelessDetailsFromSwitchesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSdaWirelessDetailsFromSwitchesV1(fabricID)
		}
		return nil, response, fmt.Errorf("error with operation GetSdaWirelessDetailsFromSwitchesV1")
	}

	result := response.Result().(*ResponseFabricWirelessGetSdaWirelessDetailsFromSwitchesV1)
	return result, response, err

}

//RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteV1 Retrieve the VLANs and SSIDs mapped to the VLAN within a Fabric Site. - edbe-baa6-46cb-9f5e
/* Retrieve the VLANs and SSIDs mapped to the VLAN, within a Fabric Site. The 'fabricId' represents the Fabric ID of a particular Fabric Site.


@param fabricID fabricId path parameter. The 'fabricId' represents the Fabric ID of a particular Fabric Site

@param RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-the-vlans-and-ssids-mapped-to-the-vlan-within-a-fabric-site
*/
func (s *FabricWirelessService) RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteV1(fabricID string, RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteV1QueryParams *RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteV1QueryParams) (*ResponseFabricWirelessRetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabrics/{fabricId}/vlanToSsids"
	path = strings.Replace(path, "{fabricId}", fmt.Sprintf("%v", fabricID), -1)

	queryString, _ := query.Values(RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseFabricWirelessRetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteV1(fabricID, RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveTheVlansAndSsidsMappedToTheVlanWithinAFabricSiteV1")
	}

	result := response.Result().(*ResponseFabricWirelessRetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteV1)
	return result, response, err

}

//ReturnsTheCountOfVLANsMappedToSSIDsInAFabricSiteV1 Returns the count of VLANs mapped to SSIDs in a Fabric Site. - e0ab-88b3-4198-a152
/* Returns the count of VLANs mapped to SSIDs in a Fabric Site. The 'fabricId' represents the Fabric ID of a particular Fabric Site.


@param fabricID fabricId path parameter. The 'fabricId' represents the Fabric ID of a particular Fabric Site


Documentation Link: https://developer.cisco.com/docs/dna-center/#!returns-the-count-of-vlans-mapped-to-ssids-in-a-fabric-site
*/
func (s *FabricWirelessService) ReturnsTheCountOfVLANsMappedToSSIDsInAFabricSiteV1(fabricID string) (*ResponseFabricWirelessReturnsTheCountOfVLANsMappedToSSIDsInAFabricSiteV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabrics/{fabricId}/vlanToSsids/count"
	path = strings.Replace(path, "{fabricId}", fmt.Sprintf("%v", fabricID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseFabricWirelessReturnsTheCountOfVLANsMappedToSSIDsInAFabricSiteV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReturnsTheCountOfVLANsMappedToSSIDsInAFabricSiteV1(fabricID)
		}
		return nil, response, fmt.Errorf("error with operation ReturnsTheCountOfVlansMappedToSsidsInAFabricSiteV1")
	}

	result := response.Result().(*ResponseFabricWirelessReturnsTheCountOfVLANsMappedToSSIDsInAFabricSiteV1)
	return result, response, err

}

//GetSdaWirelessMulticastV1 Get SDA Wireless Multicast - a58d-4bd1-4f89-8424
/* Retrieves the current Software-Defined Access (SDA) Wireless Multicast setting for a specified fabric site. The setting indicates whether multicast is enabled (true) or disabled (false). For optimal performance, ensure wired multicast is also enabled.


@param fabricID fabricId path parameter. The unique identifier of the fabric site for which the multicast setting is being requested. The identifier should be in the format of a UUID. The 'fabricId' can be obtained using the api /dna/intent/api/v1/sda/fabricSites.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-sda-wireless-multicast
*/
func (s *FabricWirelessService) GetSdaWirelessMulticastV1(fabricID string) (*ResponseFabricWirelessGetSdaWirelessMulticastV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabrics/{fabricId}/wirelessMulticast"
	path = strings.Replace(path, "{fabricId}", fmt.Sprintf("%v", fabricID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseFabricWirelessGetSdaWirelessMulticastV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSdaWirelessMulticastV1(fabricID)
		}
		return nil, response, fmt.Errorf("error with operation GetSdaWirelessMulticastV1")
	}

	result := response.Result().(*ResponseFabricWirelessGetSdaWirelessMulticastV1)
	return result, response, err

}

//AddSSIDToIPPoolMappingV1 Add SSID to IP Pool Mapping - b783-49d9-463a-98dd
/* Add SSID to IP Pool Mapping



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-ssid-to-ip-pool-mapping
*/
func (s *FabricWirelessService) AddSSIDToIPPoolMappingV1(requestFabricWirelessAddSSIDToIPPoolMappingV1 *RequestFabricWirelessAddSSIDToIPPoolMappingV1) (*ResponseFabricWirelessAddSSIDToIPPoolMappingV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/hostonboarding/ssid-ippool"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestFabricWirelessAddSSIDToIPPoolMappingV1).
		SetResult(&ResponseFabricWirelessAddSSIDToIPPoolMappingV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddSSIDToIPPoolMappingV1(requestFabricWirelessAddSSIDToIPPoolMappingV1)
		}

		return nil, response, fmt.Errorf("error with operation AddSsidToIpPoolMappingV1")
	}

	result := response.Result().(*ResponseFabricWirelessAddSSIDToIPPoolMappingV1)
	return result, response, err

}

//AddWLCToFabricDomainV1 Add WLC to Fabric Domain - f4ad-b85b-4f19-ae86
/* Add WLC to Fabric Domain



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-w-l-c-to-fabric-domain
*/
func (s *FabricWirelessService) AddWLCToFabricDomainV1(requestFabricWirelessAddWLCToFabricDomainV1 *RequestFabricWirelessAddWLCToFabricDomainV1) (*ResponseFabricWirelessAddWLCToFabricDomainV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/wireless-controller"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestFabricWirelessAddWLCToFabricDomainV1).
		SetResult(&ResponseFabricWirelessAddWLCToFabricDomainV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddWLCToFabricDomainV1(requestFabricWirelessAddWLCToFabricDomainV1)
		}

		return nil, response, fmt.Errorf("error with operation AddWLCToFabricDomainV1")
	}

	result := response.Result().(*ResponseFabricWirelessAddWLCToFabricDomainV1)
	return result, response, err

}

//ReloadSwitchForWirelessControllerCleanupV1 Reload Switch for Wireless Controller Cleanup - 7ab4-3994-429b-827c
/* This API is used to reload switches after disabling wireless to remove the wireless-controller configuration on the device. When wireless is disabled on a switch, all wireless configurations are removed except for the wireless-controller configuration. To completely remove the wireless-controller configuration, you can use this API. Please note that using this API will cause a reload of the device(s). This API should only be used for devices that have wireless disabled but still have the 'wireless-controller' configuration present. The reload payload can have a maximum of two switches as only two switches can have a wireless role in a fabric site.


@param fabricID fabricId path parameter. The 'fabricId' represents the Fabric ID of a particular Fabric Site. The 'fabricId' can be obtained using the api /dna/intent/api/v1/sda/fabricSites.  Example : e290f1ee-6c54-4b01-90e6-d701748f0851


Documentation Link: https://developer.cisco.com/docs/dna-center/#!reload-switch-for-wireless-controller-cleanup
*/
func (s *FabricWirelessService) ReloadSwitchForWirelessControllerCleanupV1(fabricID string, requestFabricWirelessReloadSwitchForWirelessControllerCleanupV1 *RequestFabricWirelessReloadSwitchForWirelessControllerCleanupV1) (*ResponseFabricWirelessReloadSwitchForWirelessControllerCleanupV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabrics/{fabricId}/switchWirelessSetting/reload"
	path = strings.Replace(path, "{fabricId}", fmt.Sprintf("%v", fabricID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestFabricWirelessReloadSwitchForWirelessControllerCleanupV1).
		SetResult(&ResponseFabricWirelessReloadSwitchForWirelessControllerCleanupV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReloadSwitchForWirelessControllerCleanupV1(fabricID, requestFabricWirelessReloadSwitchForWirelessControllerCleanupV1)
		}

		return nil, response, fmt.Errorf("error with operation ReloadSwitchForWirelessControllerCleanupV1")
	}

	result := response.Result().(*ResponseFabricWirelessReloadSwitchForWirelessControllerCleanupV1)
	return result, response, err

}

//UpdateSSIDToIPPoolMappingV1 Update SSID to IP Pool Mapping - a7bf-1936-424b-91f0
/* Update SSID to IP Pool Mapping


 */
func (s *FabricWirelessService) UpdateSSIDToIPPoolMappingV1(requestFabricWirelessUpdateSSIDToIPPoolMappingV1 *RequestFabricWirelessUpdateSSIDToIPPoolMappingV1) (*ResponseFabricWirelessUpdateSSIDToIPPoolMappingV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/hostonboarding/ssid-ippool"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestFabricWirelessUpdateSSIDToIPPoolMappingV1).
		SetResult(&ResponseFabricWirelessUpdateSSIDToIPPoolMappingV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateSSIDToIPPoolMappingV1(requestFabricWirelessUpdateSSIDToIPPoolMappingV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateSsidToIpPoolMappingV1")
	}

	result := response.Result().(*ResponseFabricWirelessUpdateSSIDToIPPoolMappingV1)
	return result, response, err

}

//SwitchWirelessSettingAndRollingApUpgradeManagementV1 Switch Wireless Setting and Rolling AP Upgrade Management - 2d95-183a-46db-87ed
/* This API is used to enable or disable wireless capabilities on switch devices, along with configuring rolling AP upgrades on the fabric site. Reboot action is required to remove wireless configurations.


@param fabricID fabricId path parameter. The 'fabricId' represents the Fabric ID of a particular Fabric Site. The 'fabricId' can be obtained using the api /dna/intent/api/v1/sda/fabricSites.  Example : e290f1ee-6c54-4b01-90e6-d701748f0851

*/
func (s *FabricWirelessService) SwitchWirelessSettingAndRollingApUpgradeManagementV1(fabricID string, requestFabricWirelessSwitchWirelessSettingAndRollingAPUpgradeManagementV1 *RequestFabricWirelessSwitchWirelessSettingAndRollingApUpgradeManagementV1) (*ResponseFabricWirelessSwitchWirelessSettingAndRollingApUpgradeManagementV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabrics/{fabricId}/switchWirelessSetting"
	path = strings.Replace(path, "{fabricId}", fmt.Sprintf("%v", fabricID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestFabricWirelessSwitchWirelessSettingAndRollingAPUpgradeManagementV1).
		SetResult(&ResponseFabricWirelessSwitchWirelessSettingAndRollingApUpgradeManagementV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SwitchWirelessSettingAndRollingApUpgradeManagementV1(fabricID, requestFabricWirelessSwitchWirelessSettingAndRollingAPUpgradeManagementV1)
		}
		return nil, response, fmt.Errorf("error with operation SwitchWirelessSettingAndRollingApUpgradeManagementV1")
	}

	result := response.Result().(*ResponseFabricWirelessSwitchWirelessSettingAndRollingApUpgradeManagementV1)
	return result, response, err

}

//AddUpdateOrRemoveSSIDMappingToAVLANV1 Add, Update or Remove SSID mapping to a VLAN - 07af-b879-4c2a-983b
/* Add, update, or remove SSID mappings to a VLAN. If the payload doesn't contain a 'vlanName' which has SSIDs mapping done earlier then all the mapped SSIDs of the 'vlanName' is cleared. The request must include all SSIDs currently mapped to a VLAN, as determined by the response from the GET operation for the same fabricId used in the request. If an already-mapped SSID is not included in the payload, its mapping will be removed by this API. Conversely, if a new SSID is provided, it will be added to the Mapping. Ensure that any new SSID added is a Fabric SSID. This API can also be used to add a VLAN and associate the relevant SSIDs with it. The 'vlanName' must be 'Fabric Wireless Enabled' and should be part of the Fabric Site representing 'Fabric ID' specified in the API request.


@param fabricID fabricId path parameter. The 'fabricId' represents the Fabric ID of a particular Fabric Site

*/
func (s *FabricWirelessService) AddUpdateOrRemoveSSIDMappingToAVLANV1(fabricID string, requestFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLANV1 *RequestFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLANV1) (*ResponseFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLANV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabrics/{fabricId}/vlanToSsids"
	path = strings.Replace(path, "{fabricId}", fmt.Sprintf("%v", fabricID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLANV1).
		SetResult(&ResponseFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLANV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddUpdateOrRemoveSSIDMappingToAVLANV1(fabricID, requestFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLANV1)
		}
		return nil, response, fmt.Errorf("error with operation AddUpdateOrRemoveSsidMappingToAVlanV1")
	}

	result := response.Result().(*ResponseFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLANV1)
	return result, response, err

}

//UpdateSdaWirelessMulticastV1 Update SDA Wireless Multicast - 908b-d97a-4458-8352
/* Updates the Software-Defined Access (SDA) Wireless Multicast setting for a specified fabric site. This API allows you to enable or disable the multicast feature. For optimal performance, ensure wired multicast is also enabled.


@param fabricID fabricId path parameter. The unique identifier of the fabric site for which the multicast setting is being requested. The identifier should be in the format of a UUID. The 'fabricId' can be obtained using the api /dna/intent/api/v1/sda/fabricSites.

*/
func (s *FabricWirelessService) UpdateSdaWirelessMulticastV1(fabricID string, requestFabricWirelessUpdateSDAWirelessMulticastV1 *RequestFabricWirelessUpdateSdaWirelessMulticastV1) (*ResponseFabricWirelessUpdateSdaWirelessMulticastV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabrics/{fabricId}/wirelessMulticast"
	path = strings.Replace(path, "{fabricId}", fmt.Sprintf("%v", fabricID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestFabricWirelessUpdateSDAWirelessMulticastV1).
		SetResult(&ResponseFabricWirelessUpdateSdaWirelessMulticastV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateSdaWirelessMulticastV1(fabricID, requestFabricWirelessUpdateSDAWirelessMulticastV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateSdaWirelessMulticastV1")
	}

	result := response.Result().(*ResponseFabricWirelessUpdateSdaWirelessMulticastV1)
	return result, response, err

}

//RemoveWLCFromFabricDomainV1 Remove WLC from Fabric Domain - 10bb-1ae9-46e9-840f
/* Remove WLC from Fabric Domain


@param RemoveWLCFromFabricDomainV1HeaderParams Custom header parameters
@param RemoveWLCFromFabricDomainV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!remove-w-l-c-from-fabric-domain
*/
func (s *FabricWirelessService) RemoveWLCFromFabricDomainV1(RemoveWLCFromFabricDomainV1HeaderParams *RemoveWLCFromFabricDomainV1HeaderParams, RemoveWLCFromFabricDomainV1QueryParams *RemoveWLCFromFabricDomainV1QueryParams) (*ResponseFabricWirelessRemoveWLCFromFabricDomainV1, *resty.Response, error) {
	//RemoveWLCFromFabricDomainV1HeaderParams *RemoveWLCFromFabricDomainV1HeaderParams,RemoveWLCFromFabricDomainV1QueryParams *RemoveWLCFromFabricDomainV1QueryParams
	path := "/dna/intent/api/v1/business/sda/wireless-controller"

	queryString, _ := query.Values(RemoveWLCFromFabricDomainV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RemoveWLCFromFabricDomainV1HeaderParams != nil {

		if RemoveWLCFromFabricDomainV1HeaderParams.Persistbapioutput != "" {
			clientRequest = clientRequest.SetHeader("__persistbapioutput", RemoveWLCFromFabricDomainV1HeaderParams.Persistbapioutput)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseFabricWirelessRemoveWLCFromFabricDomainV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RemoveWLCFromFabricDomainV1(RemoveWLCFromFabricDomainV1HeaderParams, RemoveWLCFromFabricDomainV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RemoveWLCFromFabricDomainV1")
	}

	result := response.Result().(*ResponseFabricWirelessRemoveWLCFromFabricDomainV1)
	return result, response, err

}

// Alias Function
/*
This method acts as an alias for the method `RemoveWLCFromFabricDomainV1`
*/
func (s *FabricWirelessService) RemoveWLCFromFabricDomain(RemoveWLCFromFabricDomainV1HeaderParams *RemoveWLCFromFabricDomainV1HeaderParams, RemoveWLCFromFabricDomainV1QueryParams *RemoveWLCFromFabricDomainV1QueryParams) (*ResponseFabricWirelessRemoveWLCFromFabricDomainV1, *resty.Response, error) {
	return s.RemoveWLCFromFabricDomainV1(RemoveWLCFromFabricDomainV1HeaderParams, RemoveWLCFromFabricDomainV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `ReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingV1`
*/
func (s *FabricWirelessService) ReturnsAllTheFabricSitesThatHaveVLANToSSIDMapping(ReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingV1QueryParams *ReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingV1QueryParams) (*ResponseFabricWirelessReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingV1, *resty.Response, error) {
	return s.ReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingV1(ReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `UpdateSdaWirelessMulticastV1`
*/
func (s *FabricWirelessService) UpdateSdaWirelessMulticast(fabricID string, requestFabricWirelessUpdateSDAWirelessMulticastV1 *RequestFabricWirelessUpdateSdaWirelessMulticastV1) (*ResponseFabricWirelessUpdateSdaWirelessMulticastV1, *resty.Response, error) {
	return s.UpdateSdaWirelessMulticastV1(fabricID, requestFabricWirelessUpdateSDAWirelessMulticastV1)
}

// Alias Function
/*
This method acts as an alias for the method `SwitchWirelessSettingAndRollingApUpgradeManagementV1`
*/
func (s *FabricWirelessService) SwitchWirelessSettingAndRollingApUpgradeManagement(fabricID string, requestFabricWirelessSwitchWirelessSettingAndRollingAPUpgradeManagementV1 *RequestFabricWirelessSwitchWirelessSettingAndRollingApUpgradeManagementV1) (*ResponseFabricWirelessSwitchWirelessSettingAndRollingApUpgradeManagementV1, *resty.Response, error) {
	return s.SwitchWirelessSettingAndRollingApUpgradeManagementV1(fabricID, requestFabricWirelessSwitchWirelessSettingAndRollingAPUpgradeManagementV1)
}

// Alias Function
/*
This method acts as an alias for the method `AddWLCToFabricDomainV1`
*/
func (s *FabricWirelessService) AddWLCToFabricDomain(requestFabricWirelessAddWLCToFabricDomainV1 *RequestFabricWirelessAddWLCToFabricDomainV1) (*ResponseFabricWirelessAddWLCToFabricDomainV1, *resty.Response, error) {
	return s.AddWLCToFabricDomainV1(requestFabricWirelessAddWLCToFabricDomainV1)
}

// Alias Function
/*
This method acts as an alias for the method `AddSSIDToIPPoolMappingV1`
*/
func (s *FabricWirelessService) AddSSIDToIPPoolMapping(requestFabricWirelessAddSSIDToIPPoolMappingV1 *RequestFabricWirelessAddSSIDToIPPoolMappingV1) (*ResponseFabricWirelessAddSSIDToIPPoolMappingV1, *resty.Response, error) {
	return s.AddSSIDToIPPoolMappingV1(requestFabricWirelessAddSSIDToIPPoolMappingV1)
}

// Alias Function
/*
This method acts as an alias for the method `UpdateSSIDToIPPoolMappingV1`
*/
func (s *FabricWirelessService) UpdateSSIDToIPPoolMapping(requestFabricWirelessUpdateSSIDToIPPoolMappingV1 *RequestFabricWirelessUpdateSSIDToIPPoolMappingV1) (*ResponseFabricWirelessUpdateSSIDToIPPoolMappingV1, *resty.Response, error) {
	return s.UpdateSSIDToIPPoolMappingV1(requestFabricWirelessUpdateSSIDToIPPoolMappingV1)
}

// Alias Function
/*
This method acts as an alias for the method `AddUpdateOrRemoveSSIDMappingToAVLANV1`
*/
func (s *FabricWirelessService) AddUpdateOrRemoveSSIDMappingToAVLAN(fabricID string, requestFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLANV1 *RequestFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLANV1) (*ResponseFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLANV1, *resty.Response, error) {
	return s.AddUpdateOrRemoveSSIDMappingToAVLANV1(fabricID, requestFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLANV1)
}

// Alias Function
/*
This method acts as an alias for the method `ReturnsTheCountOfVLANsMappedToSSIDsInAFabricSiteV1`
*/
func (s *FabricWirelessService) ReturnsTheCountOfVLANsMappedToSSIDsInAFabricSite(fabricID string) (*ResponseFabricWirelessReturnsTheCountOfVLANsMappedToSSIDsInAFabricSiteV1, *resty.Response, error) {
	return s.ReturnsTheCountOfVLANsMappedToSSIDsInAFabricSiteV1(fabricID)
}

// Alias Function
/*
This method acts as an alias for the method `GetSdaWirelessMulticastV1`
*/
func (s *FabricWirelessService) GetSdaWirelessMulticast(fabricID string) (*ResponseFabricWirelessGetSdaWirelessMulticastV1, *resty.Response, error) {
	return s.GetSdaWirelessMulticastV1(fabricID)
}

// Alias Function
/*
This method acts as an alias for the method `ReturnTheCountOfAllTheFabricSiteWhichHasSSIDToIPPoolMappingV1`
*/
func (s *FabricWirelessService) ReturnTheCountOfAllTheFabricSiteWhichHasSSIDToIPPoolMapping() (*ResponseFabricWirelessReturnTheCountOfAllTheFabricSiteWhichHasSSIDToIPPoolMappingV1, *resty.Response, error) {
	return s.ReturnTheCountOfAllTheFabricSiteWhichHasSSIDToIPPoolMappingV1()
}

// Alias Function
/*
This method acts as an alias for the method `ReloadSwitchForWirelessControllerCleanupV1`
*/
func (s *FabricWirelessService) ReloadSwitchForWirelessControllerCleanup(fabricID string, requestFabricWirelessReloadSwitchForWirelessControllerCleanupV1 *RequestFabricWirelessReloadSwitchForWirelessControllerCleanupV1) (*ResponseFabricWirelessReloadSwitchForWirelessControllerCleanupV1, *resty.Response, error) {
	return s.ReloadSwitchForWirelessControllerCleanupV1(fabricID, requestFabricWirelessReloadSwitchForWirelessControllerCleanupV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetSSIDToIPPoolMappingV1`
*/
func (s *FabricWirelessService) GetSSIDToIPPoolMapping(GetSSIDToIPPoolMappingV1QueryParams *GetSSIDToIPPoolMappingV1QueryParams) (*ResponseFabricWirelessGetSSIDToIPPoolMappingV1, *resty.Response, error) {
	return s.GetSSIDToIPPoolMappingV1(GetSSIDToIPPoolMappingV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetSdaWirelessDetailsFromSwitchesV1`
*/
func (s *FabricWirelessService) GetSdaWirelessDetailsFromSwitches(fabricID string) (*ResponseFabricWirelessGetSdaWirelessDetailsFromSwitchesV1, *resty.Response, error) {
	return s.GetSdaWirelessDetailsFromSwitchesV1(fabricID)
}

// Alias Function
/*
This method acts as an alias for the method `RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteV1`
*/
func (s *FabricWirelessService) RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSite(fabricID string, RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteV1QueryParams *RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteV1QueryParams) (*ResponseFabricWirelessRetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteV1, *resty.Response, error) {
	return s.RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteV1(fabricID, RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteV1QueryParams)
}
