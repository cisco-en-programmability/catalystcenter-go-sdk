package catalyst

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type PlatformService service

type ResponsePlatformCiscoDnaCenterPackagesSummaryV1 struct {
	Response *[]ResponsePlatformCiscoDnaCenterPackagesSummaryV1Response `json:"response,omitempty"` //
	Version  string                                                     `json:"version,omitempty"`  // The MAGLEV-API version (this field is for internal development purpose)
}
type ResponsePlatformCiscoDnaCenterPackagesSummaryV1Response struct {
	Name    string `json:"name,omitempty"`    // Name of installed package
	Version string `json:"version,omitempty"` // Version of installed package
}
type ResponsePlatformCiscoDnaCenterReleaseSummaryV1 struct {
	Version  string                                                  `json:"version,omitempty"`  // The MAGLEV-API version (this field is for internal development purpose)
	Response *ResponsePlatformCiscoDnaCenterReleaseSummaryV1Response `json:"response,omitempty"` //
}
type ResponsePlatformCiscoDnaCenterReleaseSummaryV1Response struct {
	CorePackages           []string                                                                        `json:"corePackages,omitempty"`           // The set of packages that are mandatory to be installed/upgraded with the release
	Packages               []string                                                                        `json:"packages,omitempty"`               // The set of packages upgrades available with this release that will not be upgraded unless selected for upgrade
	Name                   string                                                                          `json:"name,omitempty"`                   // Name of the release (example "dnac")
	InstalledVersion       string                                                                          `json:"installedVersion,omitempty"`       // The installed Cisco DNAC version
	SystemVersion          string                                                                          `json:"systemVersion,omitempty"`          // The MAGLEV-SYSTEM version
	SupportedDirectUpdates *[]ResponsePlatformCiscoDnaCenterReleaseSummaryV1ResponseSupportedDirectUpdates `json:"supportedDirectUpdates,omitempty"` // The list of earlier releases that can upgrade directly to the current release. If the supportedDirectUpdates value is empty, then no direct upgrades to the current release are allowed.
	TenantID               string                                                                          `json:"tenantId,omitempty"`               // Tenant ID (for multi tenant Cisco DNA Center)
}
type ResponsePlatformCiscoDnaCenterReleaseSummaryV1ResponseSupportedDirectUpdates interface{}
type ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryV1 struct {
	Version  string                                                             `json:"version,omitempty"`  // The MAGLEV-API version (this field is for internal development purpose)
	Response *ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryV1Response `json:"response,omitempty"` //
}
type ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryV1Response struct {
	Nodes *[]ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryV1ResponseNodes `json:"nodes,omitempty"` //
}
type ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryV1ResponseNodes struct {
	Ntp      *ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryV1ResponseNodesNtp       `json:"ntp,omitempty"`      //
	Network  *[]ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryV1ResponseNodesNetwork `json:"network,omitempty"`  //
	Proxy    *ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryV1ResponseNodesProxy     `json:"proxy,omitempty"`    //
	Platform *ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryV1ResponseNodesPlatform  `json:"platform,omitempty"` //
	ID       string                                                                           `json:"id,omitempty"`       // Cluster Identifier
	Name     string                                                                           `json:"name,omitempty"`     // Node name
}
type ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryV1ResponseNodesNtp struct {
	Servers []string `json:"servers,omitempty"` // NTP server
}
type ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryV1ResponseNodesNetwork struct {
	IntraClusterLink *bool                                                                               `json:"intra_cluster_link,omitempty"` // Flag to indicate which interface is configured as the inter-cluster link
	LacpMode         *bool                                                                               `json:"lacp_mode,omitempty"`          // LACP Mode configuration on NIC
	Inet             *ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryV1ResponseNodesNetworkInet  `json:"inet,omitempty"`               //
	Interface        string                                                                              `json:"interface,omitempty"`          // Interface name
	Inet6            *ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryV1ResponseNodesNetworkInet6 `json:"inet6,omitempty"`              //
	LacpSupported    *bool                                                                               `json:"lacp_supported,omitempty"`     // LACP Support configuration on NIC
	SLAve            []string                                                                            `json:"slave,omitempty"`              // Physical interface name
}
type ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryV1ResponseNodesNetworkInet struct {
	Routes     *[]ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryV1ResponseNodesNetworkInetRoutes     `json:"routes,omitempty"`      // Static route
	Gateway    string                                                                                         `json:"gateway,omitempty"`     // Default gateway
	DNSServers *[]ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryV1ResponseNodesNetworkInetDNSServers `json:"dns_servers,omitempty"` // DNS server
	Netmask    string                                                                                         `json:"netmask,omitempty"`     // Subnet mask
	HostIP     string                                                                                         `json:"host_ip,omitempty"`     // IP assigned
}
type ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryV1ResponseNodesNetworkInetRoutes interface{}
type ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryV1ResponseNodesNetworkInetDNSServers interface{}
type ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryV1ResponseNodesNetworkInet6 struct {
	HostIP  string `json:"host_ip,omitempty"` // IP assigned to the host
	Netmask string `json:"netmask,omitempty"` // Subnet mask of the host
}
type ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryV1ResponseNodesProxy struct {
	HTTPSProxy         string   `json:"https_proxy,omitempty"`          // Https Proxy Server
	NoProxy            []string `json:"no_proxy,omitempty"`             // Servers configured to explicitly use no proxy
	HTTPSProxyUsername string   `json:"https_proxy_username,omitempty"` // Configured Https proxy username
	HTTPProxy          string   `json:"http_proxy,omitempty"`           // Not Supported
	HTTPSProxyPassword string   `json:"https_proxy_password,omitempty"` // Configured Https excrypted proxy password.
}
type ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryV1ResponseNodesPlatform struct {
	Vendor  string `json:"vendor,omitempty"`  // Product manufacturer
	Product string `json:"product,omitempty"` // Product Identifier
	Serial  string `json:"serial,omitempty"`  // Serial number of chassis
}

//CiscoDnaCenterPackagesSummary Cisco Catalyst Center Packages Summary - f3aa-697a-453a-bba0
/* Provides information such as name, version of packages installed on the Catalyst Center.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!cisco-dna-center-packages-summary-v1
*/
func (s *PlatformService) CiscoDnaCenterPackagesSummaryV1() (*ResponsePlatformCiscoDnaCenterPackagesSummaryV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/dnac-packages"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePlatformCiscoDnaCenterPackagesSummaryV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.CiscoDnaCenterPackagesSummaryV1()
		}
		return nil, response, fmt.Errorf("error with operation CiscoDnaCenterPackagesSummaryV1")
	}

	result := response.Result().(*ResponsePlatformCiscoDnaCenterPackagesSummaryV1)
	return result, response, err

}

//CiscoDnaCenterReleaseSummary Cisco Catalyst Center Release Summary - 5b87-e929-418b-8550
/* Provides information such as API version, mandatory core packages for installation or upgrade, optional packages, Cisco Catalyst Center name and version, supported direct updates, and tenant ID.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!cisco-dna-center-release-summary-v1
*/
func (s *PlatformService) CiscoDnaCenterReleaseSummaryV1() (*ResponsePlatformCiscoDnaCenterReleaseSummaryV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/dnac-release"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePlatformCiscoDnaCenterReleaseSummaryV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.CiscoDnaCenterReleaseSummaryV1()
		}
		return nil, response, fmt.Errorf("error with operation CiscoDnaCenterReleaseSummaryV1")
	}

	result := response.Result().(*ResponsePlatformCiscoDnaCenterReleaseSummaryV1)
	return result, response, err

}

//CiscoDnaCenterNodesConfigurationSummary Cisco Catalyst Center Nodes Configuration Summary - d8b0-fb13-4f08-a967
/* Provides details about the current Cisco Catalyst Center node configuration, such as API version, node name, NTP server, intracluster link, LACP mode, network static routes, DNS server, subnet mask, host IP, default gateway, and interface information.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!cisco-dna-center-nodes-configuration-summary-v1
*/
func (s *PlatformService) CiscoDnaCenterNodesConfigurationSummaryV1() (*ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/nodes-config"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.CiscoDnaCenterNodesConfigurationSummaryV1()
		}
		return nil, response, fmt.Errorf("error with operation CiscoDnaCenterNodesConfigurationSummaryV1")
	}

	result := response.Result().(*ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryV1)
	return result, response, err

}

// Alias Function
/*
This method acts as an alias for the method `CiscoDnaCenterNodesConfigurationSummaryV1`
*/
func (s *PlatformService) CiscoDnaCenterNodesConfigurationSummary() (*ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryV1, *resty.Response, error) {
	return s.CiscoDnaCenterNodesConfigurationSummaryV1()
}

// Alias Function
/*
This method acts as an alias for the method `CiscoDnaCenterPackagesSummaryV1`
*/
func (s *PlatformService) CiscoDnaCenterPackagesSummary() (*ResponsePlatformCiscoDnaCenterPackagesSummaryV1, *resty.Response, error) {
	return s.CiscoDnaCenterPackagesSummaryV1()
}

// Alias Function
/*
This method acts as an alias for the method `CiscoDnaCenterReleaseSummaryV1`
*/
func (s *PlatformService) CiscoDnaCenterReleaseSummary() (*ResponsePlatformCiscoDnaCenterReleaseSummaryV1, *resty.Response, error) {
	return s.CiscoDnaCenterReleaseSummaryV1()
}
