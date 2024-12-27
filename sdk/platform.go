package catalyst

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type PlatformService service

type ResponsePlatformCiscoCatalystCenterPackagesSummaryV1 struct {
	Response *[]ResponsePlatformCiscoCatalystCenterPackagesSummaryV1Response `json:"response,omitempty"` //
	Version  string                                                          `json:"version,omitempty"`  // The MAGLEV-API version (this field is for internal development purpose)
}
type ResponsePlatformCiscoCatalystCenterPackagesSummaryV1Response struct {
	Name    string `json:"name,omitempty"`    // Name of installed package
	Version string `json:"version,omitempty"` // Version of installed package
}
type ResponsePlatformCiscoCatalystCenterReleaseSummaryV1 struct {
	Version  string                                                       `json:"version,omitempty"`  // The MAGLEV-API version (this field is for internal development purpose)
	Response *ResponsePlatformCiscoCatalystCenterReleaseSummaryV1Response `json:"response,omitempty"` //
}
type ResponsePlatformCiscoCatalystCenterReleaseSummaryV1Response struct {
	CorePackages []string `json:"corePackages,omitempty"` // The set of packages that are mandatory to be installed/upgraded with the release

	Packages []string `json:"packages,omitempty"` // The set of packages upgrades available with this release that will not be upgraded unless selected for upgrade

	Name string `json:"name,omitempty"` // Name of the release (example "dnac")

	InstalledVersion string `json:"installedVersion,omitempty"` // The installed Cisco Catalyst Center version

	SystemVersion string `json:"systemVersion,omitempty"` // The MAGLEV-SYSTEM version

	SupportedDirectUpdates *[]ResponsePlatformCiscoCatalystCenterReleaseSummaryV1ResponseSupportedDirectUpdates `json:"supportedDirectUpdates,omitempty"` // The list of earlier releases that can upgrade directly to the current release. If the supportedDirectUpdates value is empty, then no direct upgrades to the current release are allowed.

	TenantID string `json:"tenantId,omitempty"` // Tenant ID (for multi tenant Cisco Catalyst Center)
}
type ResponsePlatformCiscoCatalystCenterReleaseSummaryV1ResponseSupportedDirectUpdates interface{}
type ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryV1 struct {
	Version string `json:"version,omitempty"` // The MAGLEV-API version (this field is for internal development purpose)

	Response *ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryV1Response `json:"response,omitempty"` //
}
type ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryV1Response struct {
	Nodes *[]ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryV1ResponseNodes `json:"nodes,omitempty"` //
}
type ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryV1ResponseNodes struct {
	Ntp *ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryV1ResponseNodesNtp `json:"ntp,omitempty"` //

	Network *[]ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryV1ResponseNodesNetwork `json:"network,omitempty"` //

	Proxy *ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryV1ResponseNodesProxy `json:"proxy,omitempty"` //

	Platform *ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryV1ResponseNodesPlatform `json:"platform,omitempty"` //

	ID string `json:"id,omitempty"` // Cluster Identifier

	Name string `json:"name,omitempty"` // Node name
}
type ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryV1ResponseNodesNtp struct {
	Servers []string `json:"servers,omitempty"` // NTP server
}
type ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryV1ResponseNodesNetwork struct {
	IntraClusterLink *bool `json:"intra_cluster_link,omitempty"` // Flag to indicate which interface is configured as the inter-cluster link

	LacpMode *bool `json:"lacp_mode,omitempty"` // LACP Mode configuration on NIC

	Inet *ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryV1ResponseNodesNetworkInet `json:"inet,omitempty"` //

	Interface string `json:"interface,omitempty"` // Interface name

	Inet6 *ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryV1ResponseNodesNetworkInet6 `json:"inet6,omitempty"` //

	LacpSupported *bool `json:"lacp_supported,omitempty"` // LACP Support configuration on NIC

	SLAve []string `json:"slave,omitempty"` // Physical interface name
}
type ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryV1ResponseNodesNetworkInet struct {
	Routes *[]ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryV1ResponseNodesNetworkInetRoutes `json:"routes,omitempty"` // Static route

	Gateway string `json:"gateway,omitempty"` // Default gateway

	DNSServers *[]ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryV1ResponseNodesNetworkInetDNSServers `json:"dns_servers,omitempty"` // DNS server

	Netmask string `json:"netmask,omitempty"` // Subnet mask

	HostIP string `json:"host_ip,omitempty"` // IP assigned
}
type ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryV1ResponseNodesNetworkInetRoutes interface{}
type ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryV1ResponseNodesNetworkInetDNSServers interface{}
type ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryV1ResponseNodesNetworkInet6 struct {
	HostIP string `json:"host_ip,omitempty"` // IP assigned to the host

	Netmask string `json:"netmask,omitempty"` // Subnet mask of the host
}
type ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryV1ResponseNodesProxy struct {
	HTTPSProxy string `json:"https_proxy,omitempty"` // Https Proxy Server

	NoProxy []string `json:"no_proxy,omitempty"` // Servers configured to explicitly use no proxy

	HTTPSProxyUsername string `json:"https_proxy_username,omitempty"` // Configured Https proxy username

	HTTPProxy string `json:"http_proxy,omitempty"` // Not Supported

	HTTPSProxyPassword string `json:"https_proxy_password,omitempty"` // Configured Https excrypted proxy password.
}
type ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryV1ResponseNodesPlatform struct {
	Vendor string `json:"vendor,omitempty"` // Product manufacturer

	Product string `json:"product,omitempty"` // Product Identifier

	Serial string `json:"serial,omitempty"` // Serial number of chassis
}

//CiscoCatalystCenterPackagesSummaryV1 Cisco Catalyst Center Packages Summary - f3aa-697a-453a-bba0
/* Provides information such as name, version of packages installed on the Catalyst center.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!cisco-dna-center-packages-summary
*/
func (s *PlatformService) CiscoCatalystCenterPackagesSummaryV1() (*ResponsePlatformCiscoCatalystCenterPackagesSummaryV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/dnac-packages"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePlatformCiscoCatalystCenterPackagesSummaryV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.CiscoCatalystCenterPackagesSummaryV1()
		}
		return nil, response, fmt.Errorf("error with operation CiscoCatalystCenterPackagesSummaryV1")
	}

	result := response.Result().(*ResponsePlatformCiscoCatalystCenterPackagesSummaryV1)
	return result, response, err

}

//CiscoCatalystCenterReleaseSummaryV1 Cisco Catalyst Center Release Summary - 5b87-e929-418b-8550
/* Provides information such as API version, mandatory core packages for installation or upgrade, optional packages, Cisco Catalyst Center name and version, supported direct updates, and tenant ID.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!cisco-dna-center-release-summary
*/
func (s *PlatformService) CiscoCatalystCenterReleaseSummaryV1() (*ResponsePlatformCiscoCatalystCenterReleaseSummaryV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/dnac-release"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePlatformCiscoCatalystCenterReleaseSummaryV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.CiscoCatalystCenterReleaseSummaryV1()
		}
		return nil, response, fmt.Errorf("error with operation CiscoCatalystCenterReleaseSummaryV1")
	}

	result := response.Result().(*ResponsePlatformCiscoCatalystCenterReleaseSummaryV1)
	return result, response, err

}

//CiscoCatalystCenterNodesConfigurationSummaryV1 Cisco Catalyst Center Nodes Configuration Summary - d8b0-fb13-4f08-a967
/* Provides details about the current Cisco Catalyst Center node configuration, such as API version, node name, NTP server, intracluster link, LACP mode, network static routes, DNS server, subnet mask, host IP, default gateway, and interface information.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!cisco-dna-center-nodes-configuration-summary
*/
func (s *PlatformService) CiscoCatalystCenterNodesConfigurationSummaryV1() (*ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/nodes-config"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.CiscoCatalystCenterNodesConfigurationSummaryV1()
		}
		return nil, response, fmt.Errorf("error with operation CiscoCatalystCenterNodesConfigurationSummaryV1")
	}

	result := response.Result().(*ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryV1)
	return result, response, err

}

// Alias Function
/*
This method acts as an alias for the method `CiscoCatalystCenterNodesConfigurationSummaryV1`
*/
func (s *PlatformService) CiscoCatalystCenterNodesConfigurationSummary() (*ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryV1, *resty.Response, error) {
	return s.CiscoCatalystCenterNodesConfigurationSummaryV1()
}

// Alias Function
/*
This method acts as an alias for the method `CiscoCatalystCenterPackagesSummaryV1`
*/
func (s *PlatformService) CiscoCatalystCenterPackagesSummary() (*ResponsePlatformCiscoCatalystCenterPackagesSummaryV1, *resty.Response, error) {
	return s.CiscoCatalystCenterPackagesSummaryV1()
}

// Alias Function
/*
This method acts as an alias for the method `CiscoCatalystCenterReleaseSummaryV1`
*/
func (s *PlatformService) CiscoCatalystCenterReleaseSummary() (*ResponsePlatformCiscoCatalystCenterReleaseSummaryV1, *resty.Response, error) {
	return s.CiscoCatalystCenterReleaseSummaryV1()
}
