package catalyst

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type NetworkSettingsService service

type AssignDeviceCredentialToSiteV1HeaderParams struct {
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type string. Persist bapi sync response
}
type GetDeviceCredentialDetailsV1QueryParams struct {
	SiteID string `url:"siteId,omitempty"` //Site id to retrieve the credential details associated with the site.
}
type GetGlobalPoolV1QueryParams struct {
	Offset float64 `url:"offset,omitempty"` //Offset/starting row. Indexed from 1. Default value of 1.
	Limit  float64 `url:"limit,omitempty"`  //Number of Global Pools to be retrieved. Default is 25 if not specified.
}
type RetrievesGlobalIPAddressPoolsV1QueryParams struct {
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1.
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page;The minimum is 1, and the maximum is 500.
	SortBy string  `url:"sortBy,omitempty"` //A property within the response to sort by.
	Order  string  `url:"order,omitempty"`  //Whether ascending or descending order should be used to sort the response.
}
type RetrievesSubpoolsIDsOfAGlobalIPAddressPoolV1QueryParams struct {
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1.
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page;The minimum is 1, and the maximum is 500.
}
type RetrievesIPAddressSubpoolsV1QueryParams struct {
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1.
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page;The minimum is 1, and the maximum is 500.
	SortBy string  `url:"sortBy,omitempty"` //A property within the response to sort by.
	Order  string  `url:"order,omitempty"`  //Whether ascending or descending order should be used to sort the response.
	SiteID string  `url:"siteId,omitempty"` //The `id` of the site for which to retrieve IP address subpools. Only subpools whose `siteId` exactly matches will be fetched, parent or child site matches will not be included.
}
type CountsIPAddressSubpoolsV1QueryParams struct {
	SiteID string `url:"siteId,omitempty"` //The `id` of the site for which to retrieve IP address subpools. Only subpools whose `siteId` matches will be counted.
}
type GetNetworkV1QueryParams struct {
	SiteID string `url:"siteId,omitempty"` //Site id to get the network settings associated with the site.
}
type CreateNetworkV1HeaderParams struct {
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type bool. Persist bapi sync response
}
type GetReserveIPSubpoolV1QueryParams struct {
	SiteID                string  `url:"siteId,omitempty"`                //site id of site from which to retrieve associated reserve pools. Either siteId (per site queries) or ignoreInheritedGroups must be used. They can also be used together.
	Offset                float64 `url:"offset,omitempty"`                //offset/starting row. Indexed from 1.
	Limit                 float64 `url:"limit,omitempty"`                 //Number of reserve pools to be retrieved. Default is 25 if not specified. Maximum allowed limit is 500.
	IgnoreInheritedGroups bool    `url:"ignoreInheritedGroups,omitempty"` //Ignores pools inherited from parent site. Either siteId or ignoreInheritedGroups must be passed. They can also be used together.
	PoolUsage             string  `url:"poolUsage,omitempty"`             //Can take values empty, partially-full or empty-partially-full
	GroupName             string  `url:"groupName,omitempty"`             //Name of the group
}
type UpdateReserveIPSubpoolV1QueryParams struct {
	ID string `url:"id,omitempty"` //Id of subpool group
}
type RetrieveAAASettingsForASiteV1QueryParams struct {
	Inherited bool `url:"_inherited,omitempty"` //Include settings explicitly set for this site and settings inherited from sites higher in the site hierarchy; when `false`, `null` values indicate that the site inherits that setting from the parent site or a site higher in the site hierarchy.
}
type RetrieveBannerSettingsForASiteV1QueryParams struct {
	Inherited bool `url:"_inherited,omitempty"` //Include settings explicitly set for this site and settings inherited from sites higher in the site hierarchy; when `false`, `null` values indicate that the site inherits that setting from the parent site or a site higher in the site hierarchy.
}
type GetDeviceCredentialSettingsForASiteV1QueryParams struct {
	Inherited bool `url:"_inherited,omitempty"` //Include settings explicitly set for this site and settings inherited from sites higher in the site hierarchy; when `false`, `null` values indicate that the site inherits that setting from the parent site or a site higher in the site hierarchy.
}
type RetrieveDHCPSettingsForASiteV1QueryParams struct {
	Inherited bool `url:"_inherited,omitempty"` //Include settings explicitly set for this site and settings inherited from sites higher in the site hierarchy; when `false`, `null` values indicate that the site inherits that setting from the parent site or a site higher in the site hierarchy.
}
type RetrieveDNSSettingsForASiteV1QueryParams struct {
	Inherited bool `url:"_inherited,omitempty"` //Include settings explicitly set for this site and settings inherited from sites higher in the site hierarchy; when `false`, `null` values indicate that the site inherits that setting from the parent site or a site higher in the site hierarchy.
}
type RetrieveImageDistributionSettingsForASiteV1QueryParams struct {
	Inherited bool `url:"_inherited,omitempty"` //Include settings explicitly set for this site and settings inherited from sites higher in the site hierarchy; when `false`, `null` values indicate that the site inherits that setting from the parent site or a site higher in the site hierarchy.
}
type RetrieveNTPSettingsForASiteV1QueryParams struct {
	Inherited bool `url:"_inherited,omitempty"` //Include settings explicitly set for this site and settings inherited from sites higher in the site hierarchy; when `false`, `null` values indicate that the site inherits that setting from the parent site or a site higher in the site hierarchy.
}
type RetrieveTelemetrySettingsForASiteV1QueryParams struct {
	Inherited bool `url:"_inherited,omitempty"` //Include settings explicitly set for this site and settings inherited from sites higher in the site hierarchy; when `false`, `null` values indicate that the site inherits that setting from the parent site or a site higher in the site hierarchy.
}
type RetrieveTimeZoneSettingsForASiteV1QueryParams struct {
	Inherited bool `url:"_inherited,omitempty"` //Include settings explicitly set for this site and settings inherited from sites higher in the site hierarchy; when `false`, `null` values indicate that the site inherits that setting from the parent site or a site higher in the site hierarchy.
}
type GetNetworkV2QueryParams struct {
	SiteID string `url:"siteId,omitempty"` //Site Id to get the network settings associated with the site.
}

type ResponseNetworkSettingsAssignDeviceCredentialToSiteV1 struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseNetworkSettingsCreateDeviceCredentialsV1 struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseNetworkSettingsUpdateDeviceCredentialsV1 struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseNetworkSettingsGetDeviceCredentialDetailsV1 struct {
	SNMPV3      *[]ResponseNetworkSettingsGetDeviceCredentialDetailsV1SNMPV3      `json:"snmp_v3,omitempty"`       //
	HTTPRead    *[]ResponseNetworkSettingsGetDeviceCredentialDetailsV1HTTPRead    `json:"http_read,omitempty"`     //
	HTTPWrite   *[]ResponseNetworkSettingsGetDeviceCredentialDetailsV1HTTPWrite   `json:"http_write,omitempty"`    //
	SNMPV2Write *[]ResponseNetworkSettingsGetDeviceCredentialDetailsV1SNMPV2Write `json:"snmp_v2_write,omitempty"` //
	SNMPV2Read  *[]ResponseNetworkSettingsGetDeviceCredentialDetailsV1SNMPV2Read  `json:"snmp_v2_read,omitempty"`  //
	Cli         *[]ResponseNetworkSettingsGetDeviceCredentialDetailsV1Cli         `json:"cli,omitempty"`           //
}
type ResponseNetworkSettingsGetDeviceCredentialDetailsV1SNMPV3 struct {
	Username         string `json:"username,omitempty"`         // Username
	AuthPassword     string `json:"authPassword,omitempty"`     // Auth Password
	AuthType         string `json:"authType,omitempty"`         // Auth Type
	PrivacyPassword  string `json:"privacyPassword,omitempty"`  // Privacy Password
	PrivacyType      string `json:"privacyType,omitempty"`      // Privacy Type
	SNMPMode         string `json:"snmpMode,omitempty"`         // Snmp Mode
	Comments         string `json:"comments,omitempty"`         // Comments
	Description      string `json:"description,omitempty"`      // Description
	CredentialType   string `json:"credentialType,omitempty"`   // Credential Type
	InstanceUUID     string `json:"instanceUuid,omitempty"`     // Instance Uuid
	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id
	ID               string `json:"id,omitempty"`               // Id
}
type ResponseNetworkSettingsGetDeviceCredentialDetailsV1HTTPRead struct {
	Secure           string `json:"secure,omitempty"`           // Secure
	Username         string `json:"username,omitempty"`         // Username
	Password         string `json:"password,omitempty"`         // Password
	Port             string `json:"port,omitempty"`             // Port
	Comments         string `json:"comments,omitempty"`         // Comments
	Description      string `json:"description,omitempty"`      // Description
	CredentialType   string `json:"credentialType,omitempty"`   // Credential Type
	InstanceUUID     string `json:"instanceUuid,omitempty"`     // Instance Uuid
	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id
	ID               string `json:"id,omitempty"`               // Id
}
type ResponseNetworkSettingsGetDeviceCredentialDetailsV1HTTPWrite struct {
	Secure           string `json:"secure,omitempty"`           // Secure
	Username         string `json:"username,omitempty"`         // Username
	Password         string `json:"password,omitempty"`         // Password
	Port             string `json:"port,omitempty"`             // Port
	Comments         string `json:"comments,omitempty"`         // Comments
	Description      string `json:"description,omitempty"`      // Description
	CredentialType   string `json:"credentialType,omitempty"`   // Credential Type
	InstanceUUID     string `json:"instanceUuid,omitempty"`     // Instance Uuid
	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id
	ID               string `json:"id,omitempty"`               // Id
}
type ResponseNetworkSettingsGetDeviceCredentialDetailsV1SNMPV2Write struct {
	WriteCommunity   string `json:"writeCommunity,omitempty"`   // Write Community
	Comments         string `json:"comments,omitempty"`         // Comments
	Description      string `json:"description,omitempty"`      // Description
	CredentialType   string `json:"credentialType,omitempty"`   // Credential Type
	InstanceUUID     string `json:"instanceUuid,omitempty"`     // Instance Uuid
	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id
	ID               string `json:"id,omitempty"`               // Id
}
type ResponseNetworkSettingsGetDeviceCredentialDetailsV1SNMPV2Read struct {
	ReadCommunity    string `json:"readCommunity,omitempty"`    // Read Community
	Comments         string `json:"comments,omitempty"`         // Comments
	Description      string `json:"description,omitempty"`      // Description
	CredentialType   string `json:"credentialType,omitempty"`   // Credential Type
	InstanceUUID     string `json:"instanceUuid,omitempty"`     // Instance Uuid
	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id
	ID               string `json:"id,omitempty"`               // Id
}
type ResponseNetworkSettingsGetDeviceCredentialDetailsV1Cli struct {
	Username         string `json:"username,omitempty"`         // Username
	EnablePassword   string `json:"enablePassword,omitempty"`   // Enable Password
	Password         string `json:"password,omitempty"`         // Password
	Comments         string `json:"comments,omitempty"`         // Comments
	Description      string `json:"description,omitempty"`      // Description
	CredentialType   string `json:"credentialType,omitempty"`   // Credential Type
	InstanceUUID     string `json:"instanceUuid,omitempty"`     // Instance Uuid
	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id
	ID               string `json:"id,omitempty"`               // Id
}
type ResponseNetworkSettingsDeleteDeviceCredentialV1 struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseNetworkSettingsGetGlobalPoolV1 struct {
	Response *[]ResponseNetworkSettingsGetGlobalPoolV1Response `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  // Version
}
type ResponseNetworkSettingsGetGlobalPoolV1Response struct {
	IPPoolName                    string                                                       `json:"ipPoolName,omitempty"`                    // Ip Pool Name
	DhcpServerIPs                 []string                                                     `json:"dhcpServerIps,omitempty"`                 // Dhcp Server Ips
	Gateways                      []string                                                     `json:"gateways,omitempty"`                      // Gateways
	CreateTime                    *int                                                         `json:"createTime,omitempty"`                    // Create Time
	LastUpdateTime                *int                                                         `json:"lastUpdateTime,omitempty"`                // Last Update Time
	TotalIPAddressCount           *int                                                         `json:"totalIpAddressCount,omitempty"`           // Total Ip Address Count
	UsedIPAddressCount            *int                                                         `json:"usedIpAddressCount,omitempty"`            // Used Ip Address Count
	ParentUUID                    string                                                       `json:"parentUuid,omitempty"`                    // Parent Uuid
	Owner                         string                                                       `json:"owner,omitempty"`                         // Owner
	Shared                        *bool                                                        `json:"shared,omitempty"`                        // Shared
	Overlapping                   *bool                                                        `json:"overlapping,omitempty"`                   // Overlapping
	ConfigureExternalDhcp         *bool                                                        `json:"configureExternalDhcp,omitempty"`         // Configure External Dhcp
	UsedPercentage                string                                                       `json:"usedPercentage,omitempty"`                // Used Percentage
	ClientOptions                 *ResponseNetworkSettingsGetGlobalPoolV1ResponseClientOptions `json:"clientOptions,omitempty"`                 // Client Options
	IPPoolType                    string                                                       `json:"ipPoolType,omitempty"`                    // Ip Pool Type
	UnavailableIPAddressCount     *float64                                                     `json:"unavailableIpAddressCount,omitempty"`     // Unavailable Ip Address Count
	AvailableIPAddressCount       *float64                                                     `json:"availableIpAddressCount,omitempty"`       // Available Ip Address Count
	TotalAssignableIPAddressCount *int                                                         `json:"totalAssignableIpAddressCount,omitempty"` // Total Assignable Ip Address Count
	DNSServerIPs                  []string                                                     `json:"dnsServerIps,omitempty"`                  // Dns Server Ips
	HasSubpools                   *bool                                                        `json:"hasSubpools,omitempty"`                   // Has Subpools
	DefaultAssignedIPAddressCount *int                                                         `json:"defaultAssignedIpAddressCount,omitempty"` // Default Assigned Ip Address Count
	Context                       *[]ResponseNetworkSettingsGetGlobalPoolV1ResponseContext     `json:"context,omitempty"`                       //
	IPv6                          *bool                                                        `json:"ipv6,omitempty"`                          // Ipv6
	ID                            string                                                       `json:"id,omitempty"`                            // Id
	IPPoolCidr                    string                                                       `json:"ipPoolCidr,omitempty"`                    // Ip Pool Cidr
}
type ResponseNetworkSettingsGetGlobalPoolV1ResponseClientOptions interface{}
type ResponseNetworkSettingsGetGlobalPoolV1ResponseContext struct {
	Owner        string `json:"owner,omitempty"`        // Owner
	ContextKey   string `json:"contextKey,omitempty"`   // Context Key
	ContextValue string `json:"contextValue,omitempty"` // Context Value
}
type ResponseNetworkSettingsUpdateGlobalPoolV1 struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseNetworkSettingsCreateGlobalPoolV1 struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseNetworkSettingsDeleteGlobalIPPoolV1 struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseNetworkSettingsCreateAGlobalIPAddressPoolV1 struct {
	Version string `json:"version,omitempty"` // Response Version e.g. : 1.0

	Response *ResponseNetworkSettingsCreateAGlobalIPAddressPoolV1Response `json:"response,omitempty"` //
}
type ResponseNetworkSettingsCreateAGlobalIPAddressPoolV1Response struct {
	URL string `json:"url,omitempty"` // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5

	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseNetworkSettingsRetrievesGlobalIPAddressPoolsV1 struct {
	Response *[]ResponseNetworkSettingsRetrievesGlobalIPAddressPoolsV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseNetworkSettingsRetrievesGlobalIPAddressPoolsV1Response struct {
	AddressSpace *ResponseNetworkSettingsRetrievesGlobalIPAddressPoolsV1ResponseAddressSpace `json:"addressSpace,omitempty"` //

	ID string `json:"id,omitempty"` // The UUID for this global IP pool.

	Name string `json:"name,omitempty"` // The name for this reserve IP pool. Only letters, numbers, '-' (hyphen), '_' (underscore), '.' (period), and '/' (forward slash) are allowed.

	PoolType string `json:"poolType,omitempty"` // Once created, a global pool type cannot be changed. Tunnel: Assigns IP addresses to site-to-site VPN for IPSec tunneling. Generic: used for all other network types.
}
type ResponseNetworkSettingsRetrievesGlobalIPAddressPoolsV1ResponseAddressSpace struct {
	Subnet string `json:"subnet,omitempty"` // The IP address component of the CIDR notation for this subnet.

	PrefixLength *float64 `json:"prefixLength,omitempty"` // The network mask component, as a decimal, for the CIDR notation of this subnet.

	GatewayIPAddress string `json:"gatewayIpAddress,omitempty"` // The gateway IP address for this subnet.

	DhcpServers []string `json:"dhcpServers,omitempty"` // The DHCP server(s) for this subnet.

	DNSServers []string `json:"dnsServers,omitempty"` // The DNS server(s) for this subnet.

	TotalAddresses string `json:"totalAddresses,omitempty"` // The total number of addresses in the pool. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.

	UnassignableAddresses string `json:"unassignableAddresses,omitempty"` // The number of addresses in the pool that cannot be assigned. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.

	AssignedAddresses string `json:"assignedAddresses,omitempty"` // The number of addresses assigned from the pool. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.

	DefaultAssignedAddresses string `json:"defaultAssignedAddresses,omitempty"` // The number of addresses that are assigned from the pool by default. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.
}
type ResponseNetworkSettingsCountsGlobalIPAddressPoolsV1 struct {
	Response *ResponseNetworkSettingsCountsGlobalIPAddressPoolsV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // The version of the response
}
type ResponseNetworkSettingsCountsGlobalIPAddressPoolsV1Response struct {
	Count *int `json:"count,omitempty"` // The reported count.
}
type ResponseNetworkSettingsRetrievesSubpoolsIDsOfAGlobalIPAddressPoolV1 struct {
	Response *[]ResponseNetworkSettingsRetrievesSubpoolsIDsOfAGlobalIPAddressPoolV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseNetworkSettingsRetrievesSubpoolsIDsOfAGlobalIPAddressPoolV1Response struct {
	ID string `json:"id,omitempty"` // ID of the subpool
}
type ResponseNetworkSettingsCountsSubpoolsOfAGlobalIPAddressPoolV1 struct {
	Response *ResponseNetworkSettingsCountsSubpoolsOfAGlobalIPAddressPoolV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // The version of the response
}
type ResponseNetworkSettingsCountsSubpoolsOfAGlobalIPAddressPoolV1Response struct {
	Count *int `json:"count,omitempty"` // The reported count.
}
type ResponseNetworkSettingsRetrievesAGlobalIPAddressPoolV1 struct {
	Response *ResponseNetworkSettingsRetrievesAGlobalIPAddressPoolV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseNetworkSettingsRetrievesAGlobalIPAddressPoolV1Response struct {
	AddressSpace *ResponseNetworkSettingsRetrievesAGlobalIPAddressPoolV1ResponseAddressSpace `json:"addressSpace,omitempty"` //

	ID string `json:"id,omitempty"` // The UUID for this global IP pool.

	Name string `json:"name,omitempty"` // The name for this reserve IP pool. Only letters, numbers, '-' (hyphen), '_' (underscore), '.' (period), and '/' (forward slash) are allowed.

	PoolType string `json:"poolType,omitempty"` // Once created, a global pool type cannot be changed. Tunnel: Assigns IP addresses to site-to-site VPN for IPSec tunneling. Generic: used for all other network types.
}
type ResponseNetworkSettingsRetrievesAGlobalIPAddressPoolV1ResponseAddressSpace struct {
	Subnet string `json:"subnet,omitempty"` // The IP address component of the CIDR notation for this subnet.

	PrefixLength *float64 `json:"prefixLength,omitempty"` // The network mask component, as a decimal, for the CIDR notation of this subnet.

	GatewayIPAddress string `json:"gatewayIpAddress,omitempty"` // The gateway IP address for this subnet.

	DhcpServers []string `json:"dhcpServers,omitempty"` // The DHCP server(s) for this subnet.

	DNSServers []string `json:"dnsServers,omitempty"` // The DNS server(s) for this subnet.

	TotalAddresses string `json:"totalAddresses,omitempty"` // The total number of addresses in the pool. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.

	UnassignableAddresses string `json:"unassignableAddresses,omitempty"` // The number of addresses in the pool that cannot be assigned. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.

	AssignedAddresses string `json:"assignedAddresses,omitempty"` // The number of addresses assigned from the pool. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.

	DefaultAssignedAddresses string `json:"defaultAssignedAddresses,omitempty"` // The number of addresses that are assigned from the pool by default. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.
}
type ResponseNetworkSettingsUpdatesAGlobalIPAddressPoolV1 struct {
	Version string `json:"version,omitempty"` // Response Version e.g. : 1.0

	Response *ResponseNetworkSettingsUpdatesAGlobalIPAddressPoolV1Response `json:"response,omitempty"` //
}
type ResponseNetworkSettingsUpdatesAGlobalIPAddressPoolV1Response struct {
	URL string `json:"url,omitempty"` // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5

	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseNetworkSettingsDeleteAGlobalIPAddressPoolV1 struct {
	Version string `json:"version,omitempty"` // Response Version e.g. : 1.0

	Response *ResponseNetworkSettingsDeleteAGlobalIPAddressPoolV1Response `json:"response,omitempty"` //
}
type ResponseNetworkSettingsDeleteAGlobalIPAddressPoolV1Response struct {
	URL string `json:"url,omitempty"` // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5

	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseNetworkSettingsReservecreateIPAddressSubpoolsV1 struct {
	Version string `json:"version,omitempty"` // Response Version e.g. : 1.0

	Response *ResponseNetworkSettingsReservecreateIPAddressSubpoolsV1Response `json:"response,omitempty"` //
}
type ResponseNetworkSettingsReservecreateIPAddressSubpoolsV1Response struct {
	URL string `json:"url,omitempty"` // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5

	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseNetworkSettingsRetrievesIPAddressSubpoolsV1 struct {
	Response *[]ResponseNetworkSettingsRetrievesIPAddressSubpoolsV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseNetworkSettingsRetrievesIPAddressSubpoolsV1Response struct {
	ID string `json:"id,omitempty"` // The UUID for this reserve IP pool (subpool).

	IPV4AddressSpace *ResponseNetworkSettingsRetrievesIPAddressSubpoolsV1ResponseIPV4AddressSpace `json:"ipV4AddressSpace,omitempty"` //

	IPV6AddressSpace *ResponseNetworkSettingsRetrievesIPAddressSubpoolsV1ResponseIPV6AddressSpace `json:"ipV6AddressSpace,omitempty"` //

	Name string `json:"name,omitempty"` // The name for this reserve IP pool. Only letters, numbers, '-' (hyphen), '_' (underscore), '.' (period), and '/' (forward slash) are allowed.

	PoolType string `json:"poolType,omitempty"` // Once created, a subpool type cannot be changed.  LAN: Assigns IP addresses to LAN interfaces of applicable VNFs and underlay LAN automation.  Management: Assigns IP addresses to management interfaces. A management network is a dedicated network connected to VNFs for VNF management.  Service: Assigns IP addresses to service interfaces. Service networks are used for communication within VNFs.  WAN: Assigns IP addresses to NFVIS for UCS-E provisioning.  Generic: used for all other network types.

	SiteID string `json:"siteId,omitempty"` // The `id` of the site that this subpool belongs to. This must be the `id` of a non-Global site.

	SiteName string `json:"siteName,omitempty"` // The name of the site that this subpool belongs to.
}
type ResponseNetworkSettingsRetrievesIPAddressSubpoolsV1ResponseIPV4AddressSpace struct {
	Subnet string `json:"subnet,omitempty"` // The IP address component of the CIDR notation for this subnet.

	PrefixLength *float64 `json:"prefixLength,omitempty"` // The network mask component, as a decimal, for the CIDR notation of this subnet.

	GatewayIPAddress string `json:"gatewayIpAddress,omitempty"` // The gateway IP address for this subnet.

	DhcpServers []string `json:"dhcpServers,omitempty"` // The DHCP server(s) for this subnet.

	DNSServers []string `json:"dnsServers,omitempty"` // The DNS server(s) for this subnet.

	TotalAddresses string `json:"totalAddresses,omitempty"` // The total number of addresses in the pool. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.

	UnassignableAddresses string `json:"unassignableAddresses,omitempty"` // The number of addresses in the pool that cannot be assigned. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.

	AssignedAddresses string `json:"assignedAddresses,omitempty"` // The number of addresses assigned from the pool. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.

	DefaultAssignedAddresses string `json:"defaultAssignedAddresses,omitempty"` // The number of addresses that are assigned from the pool by default. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.

	SLAacSupport *bool `json:"slaacSupport,omitempty"` // If the prefixLength is 64, this option may be enabled. Stateless Address Auto-configuration (SLAAC) allows network devices to select their IP address without the use of DHCP servers.

	GlobalPoolID string `json:"globalPoolId,omitempty"` // The non-tunnel global pool for this reserve pool (which matches this IP address type). Once added this value cannot be changed.
}
type ResponseNetworkSettingsRetrievesIPAddressSubpoolsV1ResponseIPV6AddressSpace struct {
	Subnet string `json:"subnet,omitempty"` // The IP address component of the CIDR notation for this subnet.

	PrefixLength *float64 `json:"prefixLength,omitempty"` // The network mask component, as a decimal, for the CIDR notation of this subnet.

	GatewayIPAddress string `json:"gatewayIpAddress,omitempty"` // The gateway IP address for this subnet.

	DhcpServers []string `json:"dhcpServers,omitempty"` // The DHCP server(s) for this subnet.

	DNSServers []string `json:"dnsServers,omitempty"` // The DNS server(s) for this subnet.

	TotalAddresses string `json:"totalAddresses,omitempty"` // The total number of addresses in the pool. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.

	UnassignableAddresses string `json:"unassignableAddresses,omitempty"` // The number of addresses in the pool that cannot be assigned. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.

	AssignedAddresses string `json:"assignedAddresses,omitempty"` // The number of addresses assigned from the pool. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.

	DefaultAssignedAddresses string `json:"defaultAssignedAddresses,omitempty"` // The number of addresses that are assigned from the pool by default. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.

	SLAacSupport *bool `json:"slaacSupport,omitempty"` // If the prefixLength is 64, this option may be enabled. Stateless Address Auto-configuration (SLAAC) allows network devices to select their IP address without the use of DHCP servers.

	GlobalPoolID string `json:"globalPoolId,omitempty"` // The non-tunnel global pool for this reserve pool (which matches this IP address type). Once added this value cannot be changed.
}
type ResponseNetworkSettingsCountsIPAddressSubpoolsV1 struct {
	Response *ResponseNetworkSettingsCountsIPAddressSubpoolsV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // The version of the response
}
type ResponseNetworkSettingsCountsIPAddressSubpoolsV1Response struct {
	Count *int `json:"count,omitempty"` // The reported count.
}
type ResponseNetworkSettingsReleaseAnIPAddressSubpoolV1 struct {
	Version string `json:"version,omitempty"` // Response Version e.g. : 1.0

	Response *ResponseNetworkSettingsReleaseAnIPAddressSubpoolV1Response `json:"response,omitempty"` //
}
type ResponseNetworkSettingsReleaseAnIPAddressSubpoolV1Response struct {
	URL string `json:"url,omitempty"` // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5

	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseNetworkSettingsUpdatesAnIPAddressSubpoolV1 struct {
	Version string `json:"version,omitempty"` // Response Version e.g. : 1.0

	Response *ResponseNetworkSettingsUpdatesAnIPAddressSubpoolV1Response `json:"response,omitempty"` //
}
type ResponseNetworkSettingsUpdatesAnIPAddressSubpoolV1Response struct {
	URL string `json:"url,omitempty"` // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5

	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseNetworkSettingsRetrievesAnIPAddressSubpoolV1 struct {
	Response *ResponseNetworkSettingsRetrievesAnIPAddressSubpoolV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseNetworkSettingsRetrievesAnIPAddressSubpoolV1Response struct {
	ID string `json:"id,omitempty"` // The UUID for this reserve IP pool (subpool).

	IPV4AddressSpace *ResponseNetworkSettingsRetrievesAnIPAddressSubpoolV1ResponseIPV4AddressSpace `json:"ipV4AddressSpace,omitempty"` //

	IPV6AddressSpace *ResponseNetworkSettingsRetrievesAnIPAddressSubpoolV1ResponseIPV6AddressSpace `json:"ipV6AddressSpace,omitempty"` //

	Name string `json:"name,omitempty"` // The name for this reserve IP pool. Only letters, numbers, '-' (hyphen), '_' (underscore), '.' (period), and '/' (forward slash) are allowed.

	PoolType string `json:"poolType,omitempty"` // Once created, a subpool type cannot be changed.  LAN: Assigns IP addresses to LAN interfaces of applicable VNFs and underlay LAN automation.  Management: Assigns IP addresses to management interfaces. A management network is a dedicated network connected to VNFs for VNF management.  Service: Assigns IP addresses to service interfaces. Service networks are used for communication within VNFs.  WAN: Assigns IP addresses to NFVIS for UCS-E provisioning.  Generic: used for all other network types.

	SiteID string `json:"siteId,omitempty"` // The `id` of the site that this subpool belongs to. This must be the `id` of a non-Global site.

	SiteName string `json:"siteName,omitempty"` // The name of the site that this subpool belongs to.
}
type ResponseNetworkSettingsRetrievesAnIPAddressSubpoolV1ResponseIPV4AddressSpace struct {
	Subnet string `json:"subnet,omitempty"` // The IP address component of the CIDR notation for this subnet.

	PrefixLength *float64 `json:"prefixLength,omitempty"` // The network mask component, as a decimal, for the CIDR notation of this subnet.

	GatewayIPAddress string `json:"gatewayIpAddress,omitempty"` // The gateway IP address for this subnet.

	DhcpServers []string `json:"dhcpServers,omitempty"` // The DHCP server(s) for this subnet.

	DNSServers []string `json:"dnsServers,omitempty"` // The DNS server(s) for this subnet.

	TotalAddresses string `json:"totalAddresses,omitempty"` // The total number of addresses in the pool. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.

	UnassignableAddresses string `json:"unassignableAddresses,omitempty"` // The number of addresses in the pool that cannot be assigned. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.

	AssignedAddresses string `json:"assignedAddresses,omitempty"` // The number of addresses assigned from the pool. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.

	DefaultAssignedAddresses string `json:"defaultAssignedAddresses,omitempty"` // The number of addresses that are assigned from the pool by default. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.

	SLAacSupport *bool `json:"slaacSupport,omitempty"` // If the prefixLength is 64, this option may be enabled. Stateless Address Auto-configuration (SLAAC) allows network devices to select their IP address without the use of DHCP servers.

	GlobalPoolID string `json:"globalPoolId,omitempty"` // The non-tunnel global pool for this reserve pool (which matches this IP address type). Once added this value cannot be changed.
}
type ResponseNetworkSettingsRetrievesAnIPAddressSubpoolV1ResponseIPV6AddressSpace struct {
	Subnet string `json:"subnet,omitempty"` // The IP address component of the CIDR notation for this subnet.

	PrefixLength *float64 `json:"prefixLength,omitempty"` // The network mask component, as a decimal, for the CIDR notation of this subnet.

	GatewayIPAddress string `json:"gatewayIpAddress,omitempty"` // The gateway IP address for this subnet.

	DhcpServers []string `json:"dhcpServers,omitempty"` // The DHCP server(s) for this subnet.

	DNSServers []string `json:"dnsServers,omitempty"` // The DNS server(s) for this subnet.

	TotalAddresses string `json:"totalAddresses,omitempty"` // The total number of addresses in the pool. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.

	UnassignableAddresses string `json:"unassignableAddresses,omitempty"` // The number of addresses in the pool that cannot be assigned. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.

	AssignedAddresses string `json:"assignedAddresses,omitempty"` // The number of addresses assigned from the pool. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.

	DefaultAssignedAddresses string `json:"defaultAssignedAddresses,omitempty"` // The number of addresses that are assigned from the pool by default. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.

	SLAacSupport *bool `json:"slaacSupport,omitempty"` // If the prefixLength is 64, this option may be enabled. Stateless Address Auto-configuration (SLAAC) allows network devices to select their IP address without the use of DHCP servers.

	GlobalPoolID string `json:"globalPoolId,omitempty"` // The non-tunnel global pool for this reserve pool (which matches this IP address type). Once added this value cannot be changed.
}
type ResponseNetworkSettingsGetNetworkV1 struct {
	Response *[]ResponseNetworkSettingsGetNetworkV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseNetworkSettingsGetNetworkV1Response struct {
	InstanceType string `json:"instanceType,omitempty"` // Instance Type

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid

	Namespace string `json:"namespace,omitempty"` // Namespace

	Type string `json:"type,omitempty"` // Type

	Key string `json:"key,omitempty"` // Key

	Version *int `json:"version,omitempty"` // Version

	Value *[]ResponseNetworkSettingsGetNetworkV1ResponseValue `json:"value,omitempty"` //

	GroupUUID string `json:"groupUuid,omitempty"` // Group Uuid

	InheritedGroupUUID string `json:"inheritedGroupUuid,omitempty"` // Inherited Group Uuid

	InheritedGroupName string `json:"inheritedGroupName,omitempty"` // Inherited Group Name
}
type ResponseNetworkSettingsGetNetworkV1ResponseValue struct {
	IPAddresses []string `json:"ipAddresses,omitempty"` // Ip Addresses

	ConfigureDnacIP *bool `json:"configureDnacIP,omitempty"` // Configure Dnac I P
}
type ResponseNetworkSettingsCreateNetworkV1 struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url

	Message string `json:"message,omitempty"` // Message
}
type ResponseNetworkSettingsUpdateNetworkV1 struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseNetworkSettingsRetrieveCliTemplatesAttachedToANetworkProfileV1 struct {
	Version string `json:"version,omitempty"` // Response Version e.g. : 1.0

	Response *[]ResponseNetworkSettingsRetrieveCliTemplatesAttachedToANetworkProfileV1Response `json:"response,omitempty"` //
}
type ResponseNetworkSettingsRetrieveCliTemplatesAttachedToANetworkProfileV1Response struct {
	ID string `json:"id,omitempty"` // The id of the template attached to the site profile - `/intent/api/v1/templates`

	Name string `json:"name,omitempty"` // The name of the template attached to the site profile - `/intent/api/v1/templates`
}
type ResponseNetworkSettingsRetrieveCountOfCliTemplatesAttachedToANetworkProfileV1 struct {
	Version string `json:"version,omitempty"` // Response Version e.g. : 1.0

	Response *ResponseNetworkSettingsRetrieveCountOfCliTemplatesAttachedToANetworkProfileV1Response `json:"response,omitempty"` //
}
type ResponseNetworkSettingsRetrieveCountOfCliTemplatesAttachedToANetworkProfileV1Response struct {
	Count *int `json:"count,omitempty"` // The reported count
}
type ResponseNetworkSettingsGetReserveIPSubpoolV1 struct {
	Response *[]ResponseNetworkSettingsGetReserveIPSubpoolV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseNetworkSettingsGetReserveIPSubpoolV1Response struct {
	ID string `json:"id,omitempty"` // Id

	GroupName string `json:"groupName,omitempty"` // Group Name

	IPPools *[]ResponseNetworkSettingsGetReserveIPSubpoolV1ResponseIPPools `json:"ipPools,omitempty"` //

	SiteID string `json:"siteId,omitempty"` // Site Id

	SiteHierarchy string `json:"siteHierarchy,omitempty"` // Site Hierarchy

	Type string `json:"type,omitempty"` // Type

	GroupOwner string `json:"groupOwner,omitempty"` // Group Owner
}
type ResponseNetworkSettingsGetReserveIPSubpoolV1ResponseIPPools struct {
	IPPoolName string `json:"ipPoolName,omitempty"` // Ip Pool Name

	DhcpServerIPs *[]ResponseNetworkSettingsGetReserveIPSubpoolV1ResponseIPPoolsDhcpServerIPs `json:"dhcpServerIps,omitempty"` // Dhcp Server Ips

	Gateways []string `json:"gateways,omitempty"` // Gateways

	CreateTime *int `json:"createTime,omitempty"` // Create Time

	LastUpdateTime *int `json:"lastUpdateTime,omitempty"` // Last Update Time

	TotalIPAddressCount *int `json:"totalIpAddressCount,omitempty"` // Total Ip Address Count

	UsedIPAddressCount *int `json:"usedIpAddressCount,omitempty"` // Used Ip Address Count

	ParentUUID string `json:"parentUuid,omitempty"` // Parent Uuid

	Owner string `json:"owner,omitempty"` // Owner

	Shared *bool `json:"shared,omitempty"` // Shared

	Overlapping *bool `json:"overlapping,omitempty"` // Overlapping

	ConfigureExternalDhcp *bool `json:"configureExternalDhcp,omitempty"` // Configure External Dhcp

	UsedPercentage string `json:"usedPercentage,omitempty"` // Used Percentage

	ClientOptions *ResponseNetworkSettingsGetReserveIPSubpoolV1ResponseIPPoolsClientOptions `json:"clientOptions,omitempty"` // Client Options

	GroupUUID string `json:"groupUuid,omitempty"` // Group Uuid

	DNSServerIPs *[]ResponseNetworkSettingsGetReserveIPSubpoolV1ResponseIPPoolsDNSServerIPs `json:"dnsServerIps,omitempty"` // Dns Server Ips

	Context *[]ResponseNetworkSettingsGetReserveIPSubpoolV1ResponseIPPoolsContext `json:"context,omitempty"` //

	IPv6 *bool `json:"ipv6,omitempty"` // Ipv6

	ID string `json:"id,omitempty"` // Id

	IPPoolCidr string `json:"ipPoolCidr,omitempty"` // Ip Pool Cidr
}
type ResponseNetworkSettingsGetReserveIPSubpoolV1ResponseIPPoolsDhcpServerIPs interface{}
type ResponseNetworkSettingsGetReserveIPSubpoolV1ResponseIPPoolsClientOptions interface{}
type ResponseNetworkSettingsGetReserveIPSubpoolV1ResponseIPPoolsDNSServerIPs interface{}
type ResponseNetworkSettingsGetReserveIPSubpoolV1ResponseIPPoolsContext struct {
	Owner string `json:"owner,omitempty"` // Owner

	ContextKey string `json:"contextKey,omitempty"` // Context Key

	ContextValue string `json:"contextValue,omitempty"` // Context Value
}
type ResponseNetworkSettingsReleaseReserveIPSubpoolV1 struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url

	Message string `json:"message,omitempty"` // Message
}
type ResponseNetworkSettingsReserveIPSubpoolV1 struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url

	Message string `json:"message,omitempty"` // Message
}
type ResponseNetworkSettingsUpdateReserveIPSubpoolV1 struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url

	Message string `json:"message,omitempty"` // Message
}
type ResponseNetworkSettingsGetServiceProviderDetailsV1 struct {
	Response *[]ResponseNetworkSettingsGetServiceProviderDetailsV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseNetworkSettingsGetServiceProviderDetailsV1Response struct {
	InstanceType string `json:"instanceType,omitempty"` // Instance Type

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid

	Namespace string `json:"namespace,omitempty"` // Namespace

	Type string `json:"type,omitempty"` // Type

	Key string `json:"key,omitempty"` // Key

	Version *int `json:"version,omitempty"` // Version

	Value *[]ResponseNetworkSettingsGetServiceProviderDetailsV1ResponseValue `json:"value,omitempty"` //

	GroupUUID string `json:"groupUuid,omitempty"` // Group Uuid

	InheritedGroupUUID string `json:"inheritedGroupUuid,omitempty"` // Inherited Group Uuid

	InheritedGroupName string `json:"inheritedGroupName,omitempty"` // Inherited Group Name
}
type ResponseNetworkSettingsGetServiceProviderDetailsV1ResponseValue struct {
	WanProvider string `json:"wanProvider,omitempty"` // Wan Provider

	SpProfileName string `json:"spProfileName,omitempty"` // Sp Profile Name

	SLAProfileName string `json:"slaProfileName,omitempty"` // Sla Profile Name
}
type ResponseNetworkSettingsCreateSpProfileV1 struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url

	Message string `json:"message,omitempty"` // Message
}
type ResponseNetworkSettingsUpdateSpProfileV1 struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url

	Message string `json:"message,omitempty"` // Message
}
type ResponseNetworkSettingsSyncNetworkDevicesCredentialV1 struct {
	Version string `json:"version,omitempty"` // Response Version e.g. : 1.0

	Response *ResponseNetworkSettingsSyncNetworkDevicesCredentialV1Response `json:"response,omitempty"` //
}
type ResponseNetworkSettingsSyncNetworkDevicesCredentialV1Response struct {
	URL string `json:"url,omitempty"` // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5

	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseNetworkSettingsSetAAASettingsForASiteV1 struct {
	Version string `json:"version,omitempty"` // Response Version e.g. : 1.0

	Response *ResponseNetworkSettingsSetAAASettingsForASiteV1Response `json:"response,omitempty"` //
}
type ResponseNetworkSettingsSetAAASettingsForASiteV1Response struct {
	URL string `json:"url,omitempty"` // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5

	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseNetworkSettingsRetrieveAAASettingsForASiteV1 struct {
	Response *ResponseNetworkSettingsRetrieveAAASettingsForASiteV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseNetworkSettingsRetrieveAAASettingsForASiteV1Response struct {
	AAANetwork *ResponseNetworkSettingsRetrieveAAASettingsForASiteV1ResponseAAANetwork `json:"aaaNetwork,omitempty"` //

	AAAClient *ResponseNetworkSettingsRetrieveAAASettingsForASiteV1ResponseAAAClient `json:"aaaClient,omitempty"` //
}
type ResponseNetworkSettingsRetrieveAAASettingsForASiteV1ResponseAAANetwork struct {
	ServerType string `json:"serverType,omitempty"` // Server Type

	Protocol string `json:"protocol,omitempty"` // Protocol

	Pan string `json:"pan,omitempty"` // Administration Node. Required for ISE.

	PrimaryServerIP string `json:"primaryServerIp,omitempty"` // The server to use as a primary.

	SecondaryServerIP string `json:"secondaryServerIp,omitempty"` // The server to use as a secondary.

	SharedSecret string `json:"sharedSecret,omitempty"` // Shared Secret

	InheritedSiteID string `json:"inheritedSiteId,omitempty"` // Inherited Site Id.

	InheritedSiteName string `json:"inheritedSiteName,omitempty"` // Inherited Site Name.
}
type ResponseNetworkSettingsRetrieveAAASettingsForASiteV1ResponseAAAClient struct {
	ServerType string `json:"serverType,omitempty"` // Server Type

	Protocol string `json:"protocol,omitempty"` // Protocol

	Pan string `json:"pan,omitempty"` // Administration Node. Required for ISE.

	PrimaryServerIP string `json:"primaryServerIp,omitempty"` // The server to use as a primary.

	SecondaryServerIP string `json:"secondaryServerIp,omitempty"` // The server to use as a secondary.

	SharedSecret string `json:"sharedSecret,omitempty"` // Shared Secret

	InheritedSiteID string `json:"inheritedSiteId,omitempty"` // Inherited Site Id.

	InheritedSiteName string `json:"inheritedSiteName,omitempty"` // Inherited Site Name.
}
type ResponseNetworkSettingsRetrieveBannerSettingsForASiteV1 struct {
	Response *ResponseNetworkSettingsRetrieveBannerSettingsForASiteV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseNetworkSettingsRetrieveBannerSettingsForASiteV1Response struct {
	Banner *ResponseNetworkSettingsRetrieveBannerSettingsForASiteV1ResponseBanner `json:"banner,omitempty"` //
}
type ResponseNetworkSettingsRetrieveBannerSettingsForASiteV1ResponseBanner struct {
	Type string `json:"type,omitempty"` // Type

	Message string `json:"message,omitempty"` // Custom message that appears when logging into routers, switches, and hubs. Required for custom type.

	InheritedSiteID string `json:"inheritedSiteId,omitempty"` // Inherited Site Id.

	InheritedSiteName string `json:"inheritedSiteName,omitempty"` // Inherited Site Name.
}
type ResponseNetworkSettingsSetBannerSettingsForASiteV1 struct {
	Version string `json:"version,omitempty"` // Response Version e.g. : 1.0

	Response *ResponseNetworkSettingsSetBannerSettingsForASiteV1Response `json:"response,omitempty"` //
}
type ResponseNetworkSettingsSetBannerSettingsForASiteV1Response struct {
	URL string `json:"url,omitempty"` // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5

	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteV1 struct {
	Response *ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteV1Response struct {
	CliCredentialsID *ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteV1ResponseCliCredentialsID `json:"cliCredentialsId,omitempty"` //

	SNMPv2CReadCredentialsID *ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteV1ResponseSNMPv2CReadCredentialsID `json:"snmpv2cReadCredentialsId,omitempty"` //

	SNMPv2CWriteCredentialsID *ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteV1ResponseSNMPv2CWriteCredentialsID `json:"snmpv2cWriteCredentialsId,omitempty"` //

	SNMPv3CredentialsID *ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteV1ResponseSNMPv3CredentialsID `json:"snmpv3CredentialsId,omitempty"` //

	HTTPReadCredentialsID *ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteV1ResponseHTTPReadCredentialsID `json:"httpReadCredentialsId,omitempty"` //

	HTTPWriteCredentialsID *ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteV1ResponseHTTPWriteCredentialsID `json:"httpWriteCredentialsId,omitempty"` //
}
type ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteV1ResponseCliCredentialsID struct {
	CredentialsID string `json:"credentialsId,omitempty"` // The `id` of the credentials.

	InheritedSiteID string `json:"inheritedSiteId,omitempty"` // Inherited Site Id

	InheritedSiteName string `json:"inheritedSiteName,omitempty"` // Inherited Site Name
}
type ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteV1ResponseSNMPv2CReadCredentialsID struct {
	CredentialsID string `json:"credentialsId,omitempty"` // The `id` of the credentials.

	InheritedSiteID string `json:"inheritedSiteId,omitempty"` // Inherited Site Id

	InheritedSiteName string `json:"inheritedSiteName,omitempty"` // Inherited Site Name
}
type ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteV1ResponseSNMPv2CWriteCredentialsID struct {
	CredentialsID string `json:"credentialsId,omitempty"` // The `id` of the credentials.

	InheritedSiteID string `json:"inheritedSiteId,omitempty"` // Inherited Site Id

	InheritedSiteName string `json:"inheritedSiteName,omitempty"` // Inherited Site Name
}
type ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteV1ResponseSNMPv3CredentialsID struct {
	CredentialsID string `json:"credentialsId,omitempty"` // The `id` of the credentials.

	InheritedSiteID string `json:"inheritedSiteId,omitempty"` // Inherited Site Id

	InheritedSiteName string `json:"inheritedSiteName,omitempty"` // Inherited Site Name
}
type ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteV1ResponseHTTPReadCredentialsID struct {
	CredentialsID string `json:"credentialsId,omitempty"` // The `id` of the credentials.

	InheritedSiteID string `json:"inheritedSiteId,omitempty"` // Inherited Site Id

	InheritedSiteName string `json:"inheritedSiteName,omitempty"` // Inherited Site Name
}
type ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteV1ResponseHTTPWriteCredentialsID struct {
	CredentialsID string `json:"credentialsId,omitempty"` // The `id` of the credentials.

	InheritedSiteID string `json:"inheritedSiteId,omitempty"` // Inherited Site Id

	InheritedSiteName string `json:"inheritedSiteName,omitempty"` // Inherited Site Name
}
type ResponseNetworkSettingsUpdateDeviceCredentialSettingsForASiteV1 struct {
	Version string `json:"version,omitempty"` // Response Version e.g. : 1.0

	Response *ResponseNetworkSettingsUpdateDeviceCredentialSettingsForASiteV1Response `json:"response,omitempty"` //
}
type ResponseNetworkSettingsUpdateDeviceCredentialSettingsForASiteV1Response struct {
	URL string `json:"url,omitempty"` // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5

	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseNetworkSettingsGetNetworkDevicesCredentialsSyncStatusV1 struct {
	Response *ResponseNetworkSettingsGetNetworkDevicesCredentialsSyncStatusV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseNetworkSettingsGetNetworkDevicesCredentialsSyncStatusV1Response struct {
	Cli *[]ResponseNetworkSettingsGetNetworkDevicesCredentialsSyncStatusV1ResponseCli `json:"cli,omitempty"` //

	SNMPV2Read *[]ResponseNetworkSettingsGetNetworkDevicesCredentialsSyncStatusV1ResponseSNMPV2Read `json:"snmpV2Read,omitempty"` //

	SNMPV2Write *[]ResponseNetworkSettingsGetNetworkDevicesCredentialsSyncStatusV1ResponseSNMPV2Write `json:"snmpV2Write,omitempty"` //

	SNMPV3 *[]ResponseNetworkSettingsGetNetworkDevicesCredentialsSyncStatusV1ResponseSNMPV3 `json:"snmpV3,omitempty"` //
}
type ResponseNetworkSettingsGetNetworkDevicesCredentialsSyncStatusV1ResponseCli struct {
	DeviceCount *int `json:"deviceCount,omitempty"` // Device count

	Status string `json:"status,omitempty"` // Sync status
}
type ResponseNetworkSettingsGetNetworkDevicesCredentialsSyncStatusV1ResponseSNMPV2Read struct {
	DeviceCount *int `json:"deviceCount,omitempty"` // Device count

	Status string `json:"status,omitempty"` // Sync status
}
type ResponseNetworkSettingsGetNetworkDevicesCredentialsSyncStatusV1ResponseSNMPV2Write struct {
	DeviceCount *int `json:"deviceCount,omitempty"` // Device count

	Status string `json:"status,omitempty"` // Sync status
}
type ResponseNetworkSettingsGetNetworkDevicesCredentialsSyncStatusV1ResponseSNMPV3 struct {
	DeviceCount *int `json:"deviceCount,omitempty"` // Device count

	Status string `json:"status,omitempty"` // Sync status
}
type ResponseNetworkSettingsSetDhcpSettingsForASiteV1 struct {
	Version string `json:"version,omitempty"` // Response Version e.g. : 1.0

	Response *ResponseNetworkSettingsSetDhcpSettingsForASiteV1Response `json:"response,omitempty"` //
}
type ResponseNetworkSettingsSetDhcpSettingsForASiteV1Response struct {
	URL string `json:"url,omitempty"` // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5

	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseNetworkSettingsRetrieveDHCPSettingsForASiteV1 struct {
	Response *ResponseNetworkSettingsRetrieveDHCPSettingsForASiteV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseNetworkSettingsRetrieveDHCPSettingsForASiteV1Response struct {
	Dhcp *ResponseNetworkSettingsRetrieveDHCPSettingsForASiteV1ResponseDhcp `json:"dhcp,omitempty"` //
}
type ResponseNetworkSettingsRetrieveDHCPSettingsForASiteV1ResponseDhcp struct {
	Servers []string `json:"servers,omitempty"` // DHCP servers for managing client device networking configuration.

	InheritedSiteID string `json:"inheritedSiteId,omitempty"` // Inherited Site Id.

	InheritedSiteName string `json:"inheritedSiteName,omitempty"` // Inherited Site Name.
}
type ResponseNetworkSettingsRetrieveDNSSettingsForASiteV1 struct {
	Response *ResponseNetworkSettingsRetrieveDNSSettingsForASiteV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseNetworkSettingsRetrieveDNSSettingsForASiteV1Response struct {
	DNS *ResponseNetworkSettingsRetrieveDNSSettingsForASiteV1ResponseDNS `json:"dns,omitempty"` //
}
type ResponseNetworkSettingsRetrieveDNSSettingsForASiteV1ResponseDNS struct {
	DomainName string `json:"domainName,omitempty"` // Network's domain name.

	DNSServers []string `json:"dnsServers,omitempty"` // DNS servers for hostname resolution.

	InheritedSiteID string `json:"inheritedSiteId,omitempty"` // Inherited Site Id.

	InheritedSiteName string `json:"inheritedSiteName,omitempty"` // Inherited Site Name.
}
type ResponseNetworkSettingsSetDNSSettingsForASiteV1 struct {
	Version string `json:"version,omitempty"` // Response Version e.g. : 1.0

	Response *ResponseNetworkSettingsSetDNSSettingsForASiteV1Response `json:"response,omitempty"` //
}
type ResponseNetworkSettingsSetDNSSettingsForASiteV1Response struct {
	URL string `json:"url,omitempty"` // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5

	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseNetworkSettingsSetImageDistributionSettingsForASiteV1 struct {
	Version string `json:"version,omitempty"` // Response Version e.g. : 1.0

	Response *ResponseNetworkSettingsSetImageDistributionSettingsForASiteV1Response `json:"response,omitempty"` //
}
type ResponseNetworkSettingsSetImageDistributionSettingsForASiteV1Response struct {
	URL string `json:"url,omitempty"` // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5

	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseNetworkSettingsRetrieveImageDistributionSettingsForASiteV1 struct {
	Response *ResponseNetworkSettingsRetrieveImageDistributionSettingsForASiteV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseNetworkSettingsRetrieveImageDistributionSettingsForASiteV1Response struct {
	ImageDistribution *ResponseNetworkSettingsRetrieveImageDistributionSettingsForASiteV1ResponseImageDistribution `json:"imageDistribution,omitempty"` //
}
type ResponseNetworkSettingsRetrieveImageDistributionSettingsForASiteV1ResponseImageDistribution struct {
	Servers []string `json:"servers,omitempty"` // This field holds an array of unique identifiers representing image distribution servers. SFTP servers to act as image distribution servers. A distributed SWIM architecture, using suitably located SFTP servers, can help support large-scale device software image upgrades and conserve WAN bandwidth.

	InheritedSiteID string `json:"inheritedSiteId,omitempty"` // Inherited Site Id.

	InheritedSiteName string `json:"inheritedSiteName,omitempty"` // Inherited Site Name.
}
type ResponseNetworkSettingsSetNTPSettingsForASiteV1 struct {
	Version string `json:"version,omitempty"` // Response Version e.g. : 1.0

	Response *ResponseNetworkSettingsSetNTPSettingsForASiteV1Response `json:"response,omitempty"` //
}
type ResponseNetworkSettingsSetNTPSettingsForASiteV1Response struct {
	URL string `json:"url,omitempty"` // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5

	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseNetworkSettingsRetrieveNTPSettingsForASiteV1 struct {
	Response *ResponseNetworkSettingsRetrieveNTPSettingsForASiteV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseNetworkSettingsRetrieveNTPSettingsForASiteV1Response struct {
	Ntp *ResponseNetworkSettingsRetrieveNTPSettingsForASiteV1ResponseNtp `json:"ntp,omitempty"` //
}
type ResponseNetworkSettingsRetrieveNTPSettingsForASiteV1ResponseNtp struct {
	Servers []string `json:"servers,omitempty"` // NTP servers to facilitate system clock synchronization for your network.

	InheritedSiteID string `json:"inheritedSiteId,omitempty"` // Inherited Site Id.

	InheritedSiteName string `json:"inheritedSiteName,omitempty"` // Inherited Site Name.
}
type ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteV1 struct {
	Response *ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteV1Response `json:"response,omitempty"` //
	Version  string                                                              `json:"version,omitempty"`  // Version
}
type ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteV1Response struct {
	WiredDataCollection   *ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteV1ResponseWiredDataCollection   `json:"wiredDataCollection,omitempty"`   //
	WirelessTelemetry     *ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteV1ResponseWirelessTelemetry     `json:"wirelessTelemetry,omitempty"`     //
	SNMPTraps             *ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteV1ResponseSNMPTraps             `json:"snmpTraps,omitempty"`             //
	Syslogs               *ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteV1ResponseSyslogs               `json:"syslogs,omitempty"`               //
	ApplicationVisibility *ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteV1ResponseApplicationVisibility `json:"applicationVisibility,omitempty"` //
}
type ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteV1ResponseWiredDataCollection struct {
	EnableWiredDataCollectio *bool  `json:"enableWiredDataCollectio,omitempty"` // Track the presence, location, and movement of wired endpoints in the network. Traffic received from endpoints is used to extract and store their identity information (MAC address and IP address). Other features, such as IEEE 802.1X, web authentication, Cisco Security Groups (formerly TrustSec), SD-Access, and Assurance, depend on this identity information to operate properly. Wired Endpoint Data Collection enables Device Tracking policies on devices assigned to the Access role in Inventory.
	InheritedSiteID          string `json:"inheritedSiteId,omitempty"`          // Inherited Site Id
	InheritedSiteName        string `json:"inheritedSiteName,omitempty"`        // Inherited Site Name
}
type ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteV1ResponseWirelessTelemetry struct {
	EnableWirelessTelemetry *bool  `json:"enableWirelessTelemetry,omitempty"` // Enables Streaming Telemetry on your wireless controllers in order to determine the health of your wireless controller, access points and wireless clients.
	InheritedSiteID         string `json:"inheritedSiteId,omitempty"`         // Inherited Site Id
	InheritedSiteName       string `json:"inheritedSiteName,omitempty"`       // Inherited Site Name
}
type ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteV1ResponseSNMPTraps struct {
	UseBuiltinTrapServer *bool    `json:"useBuiltinTrapServer,omitempty"` // Enable this server as a destination server for SNMP traps and messages from your network
	ExternalTrapServers  []string `json:"externalTrapServers,omitempty"`  // External SNMP trap servers. Example: ["250.162.252.170","2001:db8:3c4d:15::1a2f:1a2b"]
	InheritedSiteID      string   `json:"inheritedSiteId,omitempty"`      // Inherited Site Id
	InheritedSiteName    string   `json:"inheritedSiteName,omitempty"`    // Inherited Site Name
}
type ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteV1ResponseSyslogs struct {
	UseBuiltinSyslogServer *bool    `json:"useBuiltinSyslogServer,omitempty"` // Enable this server as a destination server for syslog messages.
	ExternalSyslogServers  []string `json:"externalSyslogServers,omitempty"`  // External syslog servers. Example: ["250.162.252.170", "2001:db8:3c4d:15::1a2f:1a2b"]
	InheritedSiteID        string   `json:"inheritedSiteId,omitempty"`        // Inherited Site Id
	InheritedSiteName      string   `json:"inheritedSiteName,omitempty"`      // Inherited Site Name
}
type ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteV1ResponseApplicationVisibility struct {
	Collector                  *ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteV1ResponseApplicationVisibilityCollector `json:"collector,omitempty"`                  //
	EnableOnWiredAccessDevices *bool                                                                                             `json:"enableOnWiredAccessDevices,omitempty"` // Enable Netflow Application Telemetry and Controller Based Application Recognition (CBAR) by default upon network device site assignment for wired access devices.
	InheritedSiteID            string                                                                                            `json:"inheritedSiteId,omitempty"`            // Inherited Site Id
	InheritedSiteName          string                                                                                            `json:"inheritedSiteName,omitempty"`          // Inherited Site Name
}
type ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteV1ResponseApplicationVisibilityCollector struct {
	CollectorType string `json:"collectorType,omitempty"` // Collector Type
	Address       string `json:"address,omitempty"`       // IP Address. If collection type is 'TelemetryBrokerOrUDPDirector', this field value is mandatory otherwise it is optional. Examples: "250.162.252.170", "2001:db8:3c4d:15::1a2f:1a2b"
	Port          *int   `json:"port,omitempty"`          // Min:1; Max: 65535. If collection type is 'TelemetryBrokerOrUDPDirector', this field value is mandatory otherwise it is optional.
}
type ResponseNetworkSettingsSetTelemetrySettingsForASiteV1 struct {
	Version string `json:"version,omitempty"` // Response Version e.g. : 1.0

	Response *ResponseNetworkSettingsSetTelemetrySettingsForASiteV1Response `json:"response,omitempty"` //
}
type ResponseNetworkSettingsSetTelemetrySettingsForASiteV1Response struct {
	URL string `json:"url,omitempty"` // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5

	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseNetworkSettingsSetTimeZoneForASiteV1 struct {
	Version string `json:"version,omitempty"` // Response Version e.g. : 1.0

	Response *ResponseNetworkSettingsSetTimeZoneForASiteV1Response `json:"response,omitempty"` //
}
type ResponseNetworkSettingsSetTimeZoneForASiteV1Response struct {
	URL string `json:"url,omitempty"` // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5

	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseNetworkSettingsRetrieveTimeZoneSettingsForASiteV1 struct {
	Response *ResponseNetworkSettingsRetrieveTimeZoneSettingsForASiteV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseNetworkSettingsRetrieveTimeZoneSettingsForASiteV1Response struct {
	TimeZone *ResponseNetworkSettingsRetrieveTimeZoneSettingsForASiteV1ResponseTimeZone `json:"timeZone,omitempty"` //
}
type ResponseNetworkSettingsRetrieveTimeZoneSettingsForASiteV1ResponseTimeZone struct {
	IDentifier string `json:"identifier,omitempty"` // Time zone that corresponds to the site's physical location. The site time zone is used when scheduling device provisioning and updates. Example : GMT

	InheritedSiteID string `json:"inheritedSiteId,omitempty"` // Inherited Site Id.

	InheritedSiteName string `json:"inheritedSiteName,omitempty"` // Inherited Site Name.
}
type ResponseNetworkSettingsDeleteSpProfileV1 struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url

	Message string `json:"message,omitempty"` // Message
}
type ResponseNetworkSettingsUpdateADevicesTelemetrySettingsToConformToTheTelemetrySettingsForItsSiteV1 struct {
	Version string `json:"version,omitempty"` // Response Version e.g. : 1.0

	Response *ResponseNetworkSettingsUpdateADevicesTelemetrySettingsToConformToTheTelemetrySettingsForItsSiteV1Response `json:"response,omitempty"` //
}
type ResponseNetworkSettingsUpdateADevicesTelemetrySettingsToConformToTheTelemetrySettingsForItsSiteV1Response struct {
	URL string `json:"url,omitempty"` // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5

	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseNetworkSettingsAssignDeviceCredentialToSiteV2 struct {
	Response *ResponseNetworkSettingsAssignDeviceCredentialToSiteV2Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseNetworkSettingsAssignDeviceCredentialToSiteV2Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	URL string `json:"url,omitempty"` // Url
}
type ResponseNetworkSettingsGetNetworkV2 struct {
	Response *[]ResponseNetworkSettingsGetNetworkV2Response `json:"response,omitempty"` //
}
type ResponseNetworkSettingsGetNetworkV2Response struct {
	InstanceType string `json:"instanceType,omitempty"` // Instance Type

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid

	Namespace string `json:"namespace,omitempty"` // Namespace

	Type string `json:"type,omitempty"` // Type

	Key string `json:"key,omitempty"` // Key

	Version *int `json:"version,omitempty"` // Version

	Value []string `json:"value,omitempty"` // Value

	GroupUUID string `json:"groupUuid,omitempty"` // Group Uuid

	InheritedGroupUUID string `json:"inheritedGroupUuid,omitempty"` // Inherited Group Uuid

	InheritedGroupName string `json:"inheritedGroupName,omitempty"` // Inherited Group Name
}
type ResponseNetworkSettingsCreateNetworkV2 struct {
	Response *ResponseNetworkSettingsCreateNetworkV2Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseNetworkSettingsCreateNetworkV2Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	URL string `json:"url,omitempty"` // Url
}
type ResponseNetworkSettingsUpdateNetworkV2 struct {
	Response *ResponseNetworkSettingsUpdateNetworkV2Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseNetworkSettingsUpdateNetworkV2Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	URL string `json:"url,omitempty"` // Url
}
type ResponseNetworkSettingsCreateSpProfileV2 struct {
	Response *ResponseNetworkSettingsCreateSpProfileV2Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseNetworkSettingsCreateSpProfileV2Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	URL string `json:"url,omitempty"` // Url
}
type ResponseNetworkSettingsUpdateSpProfileV2 struct {
	Response *ResponseNetworkSettingsUpdateSpProfileV2Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseNetworkSettingsUpdateSpProfileV2Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	URL string `json:"url,omitempty"` // Url
}
type ResponseNetworkSettingsGetServiceProviderDetailsV2 struct {
	Response *[]ResponseNetworkSettingsGetServiceProviderDetailsV2Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseNetworkSettingsGetServiceProviderDetailsV2Response struct {
	InstanceType string `json:"instanceType,omitempty"` // Instance Type

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid

	Namespace string `json:"namespace,omitempty"` // Namespace

	Type string `json:"type,omitempty"` // Type

	Key string `json:"key,omitempty"` // Key

	Version *int `json:"version,omitempty"` // Version

	Value *[]ResponseNetworkSettingsGetServiceProviderDetailsV2ResponseValue `json:"value,omitempty"` //

	GroupUUID string `json:"groupUuid,omitempty"` // Group Uuid

	InheritedGroupUUID string `json:"inheritedGroupUuid,omitempty"` // Inherited Group Uuid

	InheritedGroupName string `json:"inheritedGroupName,omitempty"` // Inherited Group Name
}
type ResponseNetworkSettingsGetServiceProviderDetailsV2ResponseValue struct {
	WanProvider string `json:"wanProvider,omitempty"` // Wan Provider

	SpProfileName string `json:"spProfileName,omitempty"` // Sp Profile Name

	SLAProfileName string `json:"slaProfileName,omitempty"` // Sla Profile Name
}
type ResponseNetworkSettingsDeleteSpProfileV2 struct {
	Response *ResponseNetworkSettingsDeleteSpProfileV2Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseNetworkSettingsDeleteSpProfileV2Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	URL string `json:"url,omitempty"` // Url
}
type RequestNetworkSettingsAssignDeviceCredentialToSiteV1 struct {
	CliID string `json:"cliId,omitempty"` // Cli Id

	SNMPV2ReadID string `json:"snmpV2ReadId,omitempty"` // Snmp V2 Read Id

	SNMPV2WriteID string `json:"snmpV2WriteId,omitempty"` // Snmp V2 Write Id

	HTTPRead string `json:"httpRead,omitempty"` // Http Read

	HTTPWrite string `json:"httpWrite,omitempty"` // Http Write

	SNMPV3ID string `json:"snmpV3Id,omitempty"` // Snmp V3 Id
}
type RequestNetworkSettingsCreateDeviceCredentialsV1 struct {
	Settings *RequestNetworkSettingsCreateDeviceCredentialsV1Settings `json:"settings,omitempty"` //
}
type RequestNetworkSettingsCreateDeviceCredentialsV1Settings struct {
	CliCredential *[]RequestNetworkSettingsCreateDeviceCredentialsV1SettingsCliCredential `json:"cliCredential,omitempty"` //

	SNMPV2CRead *[]RequestNetworkSettingsCreateDeviceCredentialsV1SettingsSNMPV2CRead `json:"snmpV2cRead,omitempty"` //

	SNMPV2CWrite *[]RequestNetworkSettingsCreateDeviceCredentialsV1SettingsSNMPV2CWrite `json:"snmpV2cWrite,omitempty"` //

	SNMPV3 *[]RequestNetworkSettingsCreateDeviceCredentialsV1SettingsSNMPV3 `json:"snmpV3,omitempty"` //

	HTTPSRead *[]RequestNetworkSettingsCreateDeviceCredentialsV1SettingsHTTPSRead `json:"httpsRead,omitempty"` //

	HTTPSWrite *[]RequestNetworkSettingsCreateDeviceCredentialsV1SettingsHTTPSWrite `json:"httpsWrite,omitempty"` //
}
type RequestNetworkSettingsCreateDeviceCredentialsV1SettingsCliCredential struct {
	Description string `json:"description,omitempty"` // Name or description for CLI credential

	Username string `json:"username,omitempty"` // User name for CLI credential

	Password string `json:"password,omitempty"` // Password for CLI credential

	EnablePassword string `json:"enablePassword,omitempty"` // Enable password for CLI credential
}
type RequestNetworkSettingsCreateDeviceCredentialsV1SettingsSNMPV2CRead struct {
	Description string `json:"description,omitempty"` // Description for snmp v2 read

	ReadCommunity string `json:"readCommunity,omitempty"` // Ready community for snmp v2 read credential
}
type RequestNetworkSettingsCreateDeviceCredentialsV1SettingsSNMPV2CWrite struct {
	Description string `json:"description,omitempty"` // Description for snmp v2 write

	WriteCommunity string `json:"writeCommunity,omitempty"` // Write community for snmp v2 write credential
}
type RequestNetworkSettingsCreateDeviceCredentialsV1SettingsSNMPV3 struct {
	Description string `json:"description,omitempty"` // Name or description for SNMPV3 credential

	Username string `json:"username,omitempty"` // User name for SNMPv3 credential

	PrivacyType string `json:"privacyType,omitempty"` // Privacy type for snmpv3 credential

	PrivacyPassword string `json:"privacyPassword,omitempty"` // Privacy password for snmpv3 credential

	AuthType string `json:"authType,omitempty"` // Authentication type for snmpv3 credential

	AuthPassword string `json:"authPassword,omitempty"` // Authentication password for snmpv3 credential

	SNMPMode string `json:"snmpMode,omitempty"` // Mode for snmpv3 credential
}
type RequestNetworkSettingsCreateDeviceCredentialsV1SettingsHTTPSRead struct {
	Name string `json:"name,omitempty"` // Name or description of http read credential

	Username string `json:"username,omitempty"` // User name of the http read credential

	Password string `json:"password,omitempty"` // Password for http read credential

	Port *float64 `json:"port,omitempty"` // Port for http read credential
}
type RequestNetworkSettingsCreateDeviceCredentialsV1SettingsHTTPSWrite struct {
	Name string `json:"name,omitempty"` // Name or description of http write credential

	Username string `json:"username,omitempty"` // User name of the http write credential

	Password string `json:"password,omitempty"` // Password for http write credential

	Port *float64 `json:"port,omitempty"` // Port for http write credential
}
type RequestNetworkSettingsUpdateDeviceCredentialsV1 struct {
	Settings *RequestNetworkSettingsUpdateDeviceCredentialsV1Settings `json:"settings,omitempty"` //
}
type RequestNetworkSettingsUpdateDeviceCredentialsV1Settings struct {
	CliCredential *RequestNetworkSettingsUpdateDeviceCredentialsV1SettingsCliCredential `json:"cliCredential,omitempty"` //

	SNMPV2CRead *RequestNetworkSettingsUpdateDeviceCredentialsV1SettingsSNMPV2CRead `json:"snmpV2cRead,omitempty"` //

	SNMPV2CWrite *RequestNetworkSettingsUpdateDeviceCredentialsV1SettingsSNMPV2CWrite `json:"snmpV2cWrite,omitempty"` //

	SNMPV3 *RequestNetworkSettingsUpdateDeviceCredentialsV1SettingsSNMPV3 `json:"snmpV3,omitempty"` //

	HTTPSRead *RequestNetworkSettingsUpdateDeviceCredentialsV1SettingsHTTPSRead `json:"httpsRead,omitempty"` //

	HTTPSWrite *RequestNetworkSettingsUpdateDeviceCredentialsV1SettingsHTTPSWrite `json:"httpsWrite,omitempty"` //
}
type RequestNetworkSettingsUpdateDeviceCredentialsV1SettingsCliCredential struct {
	Description string `json:"description,omitempty"` // Description

	Username string `json:"username,omitempty"` // Username

	Password string `json:"password,omitempty"` // Password

	EnablePassword string `json:"enablePassword,omitempty"` // Enable Password

	ID string `json:"id,omitempty"` // Id
}
type RequestNetworkSettingsUpdateDeviceCredentialsV1SettingsSNMPV2CRead struct {
	Description string `json:"description,omitempty"` // Description

	ReadCommunity string `json:"readCommunity,omitempty"` // Read Community

	ID string `json:"id,omitempty"` // Id
}
type RequestNetworkSettingsUpdateDeviceCredentialsV1SettingsSNMPV2CWrite struct {
	Description string `json:"description,omitempty"` // Description

	WriteCommunity string `json:"writeCommunity,omitempty"` // Write Community

	ID string `json:"id,omitempty"` // Id
}
type RequestNetworkSettingsUpdateDeviceCredentialsV1SettingsSNMPV3 struct {
	AuthPassword string `json:"authPassword,omitempty"` // Auth Password

	AuthType string `json:"authType,omitempty"` // Auth Type

	SNMPMode string `json:"snmpMode,omitempty"` // Snmp Mode

	PrivacyPassword string `json:"privacyPassword,omitempty"` // Privacy Password

	PrivacyType string `json:"privacyType,omitempty"` // Privacy Type

	Username string `json:"username,omitempty"` // Username

	Description string `json:"description,omitempty"` // Description

	ID string `json:"id,omitempty"` // Id
}
type RequestNetworkSettingsUpdateDeviceCredentialsV1SettingsHTTPSRead struct {
	Name string `json:"name,omitempty"` // Name

	Username string `json:"username,omitempty"` // Username

	Password string `json:"password,omitempty"` // Password

	Port string `json:"port,omitempty"` // Port

	ID string `json:"id,omitempty"` // Id
}
type RequestNetworkSettingsUpdateDeviceCredentialsV1SettingsHTTPSWrite struct {
	Name string `json:"name,omitempty"` // Name

	Username string `json:"username,omitempty"` // Username

	Password string `json:"password,omitempty"` // Password

	Port string `json:"port,omitempty"` // Port

	ID string `json:"id,omitempty"` // Id
}
type RequestNetworkSettingsUpdateGlobalPoolV1 struct {
	Settings *RequestNetworkSettingsUpdateGlobalPoolV1Settings `json:"settings,omitempty"` //
}
type RequestNetworkSettingsUpdateGlobalPoolV1Settings struct {
	IPpool *[]RequestNetworkSettingsUpdateGlobalPoolV1SettingsIPpool `json:"ippool,omitempty"` //
}
type RequestNetworkSettingsUpdateGlobalPoolV1SettingsIPpool struct {
	IPPoolName string `json:"ipPoolName,omitempty"` // Ip Pool Name

	Gateway string `json:"gateway,omitempty"` // Gateway

	DhcpServerIPs []string `json:"dhcpServerIps,omitempty"` // Dhcp Server Ips

	DNSServerIPs []string `json:"dnsServerIps,omitempty"` // Dns Server Ips

	ID string `json:"id,omitempty"` // Id
}
type RequestNetworkSettingsCreateGlobalPoolV1 struct {
	Settings *RequestNetworkSettingsCreateGlobalPoolV1Settings `json:"settings,omitempty"` //
}
type RequestNetworkSettingsCreateGlobalPoolV1Settings struct {
	IPpool *[]RequestNetworkSettingsCreateGlobalPoolV1SettingsIPpool `json:"ippool,omitempty"` //
}
type RequestNetworkSettingsCreateGlobalPoolV1SettingsIPpool struct {
	IPPoolName string `json:"ipPoolName,omitempty"` // Ip Pool Name

	Type string `json:"type,omitempty"` // Type

	IPPoolCidr string `json:"ipPoolCidr,omitempty"` // Ip Pool Cidr

	Gateway string `json:"gateway,omitempty"` // Gateway

	DhcpServerIPs []string `json:"dhcpServerIps,omitempty"` // Dhcp Server Ips

	DNSServerIPs []string `json:"dnsServerIps,omitempty"` // Dns Server Ips

	IPAddressSpace string `json:"IpAddressSpace,omitempty"` // Ip Address Space
}
type RequestNetworkSettingsCreateAGlobalIPAddressPoolV1 struct {
	AddressSpace *RequestNetworkSettingsCreateAGlobalIPAddressPoolV1AddressSpace `json:"addressSpace,omitempty"` //

	Name string `json:"name,omitempty"` // The name for this reserve IP pool. Only letters, numbers, '-' (hyphen), '_' (underscore), '.' (period), and '/' (forward slash) are allowed.

	PoolType string `json:"poolType,omitempty"` // Once created, a global pool type cannot be changed. Tunnel: Assigns IP addresses to site-to-site VPN for IPSec tunneling. Generic: used for all other network types.
}
type RequestNetworkSettingsCreateAGlobalIPAddressPoolV1AddressSpace struct {
	Subnet string `json:"subnet,omitempty"` // The IP address component of the CIDR notation for this subnet.

	PrefixLength *float64 `json:"prefixLength,omitempty"` // The network mask component, as a decimal, for the CIDR notation of this subnet.

	GatewayIPAddress string `json:"gatewayIpAddress,omitempty"` // The gateway IP address for this subnet.

	DhcpServers []string `json:"dhcpServers,omitempty"` // The DHCP server(s) for this subnet.

	DNSServers []string `json:"dnsServers,omitempty"` // The DNS server(s) for this subnet.
}
type RequestNetworkSettingsUpdatesAGlobalIPAddressPoolV1 struct {
	AddressSpace *RequestNetworkSettingsUpdatesAGlobalIPAddressPoolV1AddressSpace `json:"addressSpace,omitempty"` //

	Name string `json:"name,omitempty"` // The name for this reserve IP pool. Only letters, numbers, '-' (hyphen), '_' (underscore), '.' (period), and '/' (forward slash) are allowed.

	PoolType string `json:"poolType,omitempty"` // Once created, a global pool type cannot be changed. Tunnel: Assigns IP addresses to site-to-site VPN for IPSec tunneling. Generic: used for all other network types.
}
type RequestNetworkSettingsUpdatesAGlobalIPAddressPoolV1AddressSpace struct {
	Subnet string `json:"subnet,omitempty"` // The IP address component of the CIDR notation for this subnet.

	PrefixLength *float64 `json:"prefixLength,omitempty"` // The network mask component, as a decimal, for the CIDR notation of this subnet.

	GatewayIPAddress string `json:"gatewayIpAddress,omitempty"` // The gateway IP address for this subnet.

	DhcpServers []string `json:"dhcpServers,omitempty"` // The DHCP server(s) for this subnet.

	DNSServers []string `json:"dnsServers,omitempty"` // The DNS server(s) for this subnet.
}
type RequestNetworkSettingsReservecreateIPAddressSubpoolsV1 struct {
	IPV4AddressSpace *RequestNetworkSettingsReservecreateIPAddressSubpoolsV1IPV4AddressSpace `json:"ipV4AddressSpace,omitempty"` //

	IPV6AddressSpace *RequestNetworkSettingsReservecreateIPAddressSubpoolsV1IPV6AddressSpace `json:"ipV6AddressSpace,omitempty"` //

	Name string `json:"name,omitempty"` // The name for this reserve IP pool. Only letters, numbers, '-' (hyphen), '_' (underscore), '.' (period), and '/' (forward slash) are allowed.

	PoolType string `json:"poolType,omitempty"` // Once created, a subpool type cannot be changed.  LAN: Assigns IP addresses to LAN interfaces of applicable VNFs and underlay LAN automation.  Management: Assigns IP addresses to management interfaces. A management network is a dedicated network connected to VNFs for VNF management.  Service: Assigns IP addresses to service interfaces. Service networks are used for communication within VNFs.  WAN: Assigns IP addresses to NFVIS for UCS-E provisioning.  Generic: used for all other network types.

	SiteID string `json:"siteId,omitempty"` // The `id` of the site that this subpool belongs to. This must be the `id` of a non-Global site.

	SiteName string `json:"siteName,omitempty"` // The name of the site that this subpool belongs to.
}
type RequestNetworkSettingsReservecreateIPAddressSubpoolsV1IPV4AddressSpace struct {
	Subnet string `json:"subnet,omitempty"` // The IP address component of the CIDR notation for this subnet.

	PrefixLength *float64 `json:"prefixLength,omitempty"` // The network mask component, as a decimal, for the CIDR notation of this subnet.

	GatewayIPAddress string `json:"gatewayIpAddress,omitempty"` // The gateway IP address for this subnet.

	DhcpServers []string `json:"dhcpServers,omitempty"` // The DHCP server(s) for this subnet.

	DNSServers []string `json:"dnsServers,omitempty"` // The DNS server(s) for this subnet.

	TotalAddresses string `json:"totalAddresses,omitempty"` // The total number of addresses in the pool. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.

	UnassignableAddresses string `json:"unassignableAddresses,omitempty"` // The number of addresses in the pool that cannot be assigned. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.

	AssignedAddresses string `json:"assignedAddresses,omitempty"` // The number of addresses assigned from the pool. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.

	DefaultAssignedAddresses string `json:"defaultAssignedAddresses,omitempty"` // The number of addresses that are assigned from the pool by default. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.

	SLAacSupport *bool `json:"slaacSupport,omitempty"` // If the prefixLength is 64, this option may be enabled. Stateless Address Auto-configuration (SLAAC) allows network devices to select their IP address without the use of DHCP servers.

	GlobalPoolID string `json:"globalPoolId,omitempty"` // The non-tunnel global pool for this reserve pool (which matches this IP address type). Once added this value cannot be changed.
}
type RequestNetworkSettingsReservecreateIPAddressSubpoolsV1IPV6AddressSpace struct {
	Subnet string `json:"subnet,omitempty"` // The IP address component of the CIDR notation for this subnet.

	PrefixLength *float64 `json:"prefixLength,omitempty"` // The network mask component, as a decimal, for the CIDR notation of this subnet.

	GatewayIPAddress string `json:"gatewayIpAddress,omitempty"` // The gateway IP address for this subnet.

	DhcpServers []string `json:"dhcpServers,omitempty"` // The DHCP server(s) for this subnet.

	DNSServers []string `json:"dnsServers,omitempty"` // The DNS server(s) for this subnet.

	TotalAddresses string `json:"totalAddresses,omitempty"` // The total number of addresses in the pool. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.

	UnassignableAddresses string `json:"unassignableAddresses,omitempty"` // The number of addresses in the pool that cannot be assigned. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.

	AssignedAddresses string `json:"assignedAddresses,omitempty"` // The number of addresses assigned from the pool. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.

	DefaultAssignedAddresses string `json:"defaultAssignedAddresses,omitempty"` // The number of addresses that are assigned from the pool by default. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.

	SLAacSupport *bool `json:"slaacSupport,omitempty"` // If the prefixLength is 64, this option may be enabled. Stateless Address Auto-configuration (SLAAC) allows network devices to select their IP address without the use of DHCP servers.

	GlobalPoolID string `json:"globalPoolId,omitempty"` // The non-tunnel global pool for this reserve pool (which matches this IP address type). Once added this value cannot be changed.
}
type RequestNetworkSettingsUpdatesAnIPAddressSubpoolV1 struct {
	IPV4AddressSpace *RequestNetworkSettingsUpdatesAnIPAddressSubpoolV1IPV4AddressSpace `json:"ipV4AddressSpace,omitempty"` //

	IPV6AddressSpace *RequestNetworkSettingsUpdatesAnIPAddressSubpoolV1IPV6AddressSpace `json:"ipV6AddressSpace,omitempty"` //

	Name string `json:"name,omitempty"` // The name for this reserve IP pool. Only letters, numbers, '-' (hyphen), '_' (underscore), '.' (period), and '/' (forward slash) are allowed.

	PoolType string `json:"poolType,omitempty"` // Once created, a subpool type cannot be changed.  LAN: Assigns IP addresses to LAN interfaces of applicable VNFs and underlay LAN automation.  Management: Assigns IP addresses to management interfaces. A management network is a dedicated network connected to VNFs for VNF management.  Service: Assigns IP addresses to service interfaces. Service networks are used for communication within VNFs.  WAN: Assigns IP addresses to NFVIS for UCS-E provisioning.  Generic: used for all other network types.

	SiteID string `json:"siteId,omitempty"` // The `id` of the site that this subpool belongs to. This must be the `id` of a non-Global site.
}
type RequestNetworkSettingsUpdatesAnIPAddressSubpoolV1IPV4AddressSpace struct {
	Subnet string `json:"subnet,omitempty"` // The IP address component of the CIDR notation for this subnet.

	PrefixLength *float64 `json:"prefixLength,omitempty"` // The network mask component, as a decimal, for the CIDR notation of this subnet.

	GatewayIPAddress string `json:"gatewayIpAddress,omitempty"` // The gateway IP address for this subnet.

	DhcpServers []string `json:"dhcpServers,omitempty"` // The DHCP server(s) for this subnet.

	DNSServers []string `json:"dnsServers,omitempty"` // The DNS server(s) for this subnet.

	SLAacSupport *bool `json:"slaacSupport,omitempty"` // If the prefixLength is 64, this option may be enabled. Stateless Address Auto-configuration (SLAAC) allows network devices to select their IP address without the use of DHCP servers.

	GlobalPoolID string `json:"globalPoolId,omitempty"` // The non-tunnel global pool for this reserve pool (which matches this IP address type). Once added this value cannot be changed.
}
type RequestNetworkSettingsUpdatesAnIPAddressSubpoolV1IPV6AddressSpace struct {
	Subnet string `json:"subnet,omitempty"` // The IP address component of the CIDR notation for this subnet.

	PrefixLength *float64 `json:"prefixLength,omitempty"` // The network mask component, as a decimal, for the CIDR notation of this subnet.

	GatewayIPAddress string `json:"gatewayIpAddress,omitempty"` // The gateway IP address for this subnet.

	DhcpServers []string `json:"dhcpServers,omitempty"` // The DHCP server(s) for this subnet.

	DNSServers []string `json:"dnsServers,omitempty"` // The DNS server(s) for this subnet.

	SLAacSupport *bool `json:"slaacSupport,omitempty"` // If the prefixLength is 64, this option may be enabled. Stateless Address Auto-configuration (SLAAC) allows network devices to select their IP address without the use of DHCP servers.

	GlobalPoolID string `json:"globalPoolId,omitempty"` // The non-tunnel global pool for this reserve pool (which matches this IP address type). Once added this value cannot be changed.
}
type RequestNetworkSettingsCreateNetworkV1 struct {
	Settings *RequestNetworkSettingsCreateNetworkV1Settings `json:"settings,omitempty"` //
}
type RequestNetworkSettingsCreateNetworkV1Settings struct {
	DhcpServer []string `json:"dhcpServer,omitempty"` // DHCP Server IP (eg: 1.1.1.1)

	DNSServer *RequestNetworkSettingsCreateNetworkV1SettingsDNSServer `json:"dnsServer,omitempty"` //

	SyslogServer *RequestNetworkSettingsCreateNetworkV1SettingsSyslogServer `json:"syslogServer,omitempty"` //

	SNMPServer *RequestNetworkSettingsCreateNetworkV1SettingsSNMPServer `json:"snmpServer,omitempty"` //

	Netflowcollector *RequestNetworkSettingsCreateNetworkV1SettingsNetflowcollector `json:"netflowcollector,omitempty"` //

	NtpServer []string `json:"ntpServer,omitempty"` // IP address for NTP server (eg: 1.1.1.2)

	Timezone string `json:"timezone,omitempty"` // Input for time zone (eg: Africa/Abidjan)

	MessageOfTheday *RequestNetworkSettingsCreateNetworkV1SettingsMessageOfTheday `json:"messageOfTheday,omitempty"` //

	NetworkAAA *RequestNetworkSettingsCreateNetworkV1SettingsNetworkAAA `json:"network_aaa,omitempty"` //

	ClientAndEndpointAAA *RequestNetworkSettingsCreateNetworkV1SettingsClientAndEndpointAAA `json:"clientAndEndpoint_aaa,omitempty"` //
}
type RequestNetworkSettingsCreateNetworkV1SettingsDNSServer struct {
	DomainName string `json:"domainName,omitempty"` // Domain Name of DHCP (eg; cisco)

	PrimaryIPAddress string `json:"primaryIpAddress,omitempty"` // Primary IP Address for DHCP (eg: 2.2.2.2)

	SecondaryIPAddress string `json:"secondaryIpAddress,omitempty"` // Secondary IP Address for DHCP (eg: 3.3.3.3)
}
type RequestNetworkSettingsCreateNetworkV1SettingsSyslogServer struct {
	IPAddresses []string `json:"ipAddresses,omitempty"` // IP Address for syslog server (eg: 4.4.4.4)

	ConfigureDnacIP *bool `json:"configureDnacIP,omitempty"` // Configuration DNAC IP for syslog server (eg: true)
}
type RequestNetworkSettingsCreateNetworkV1SettingsSNMPServer struct {
	IPAddresses []string `json:"ipAddresses,omitempty"` // IP Address for SNMP Server (eg: 4.4.4.1)

	ConfigureDnacIP *bool `json:"configureDnacIP,omitempty"` // Configuration DNAC IP for SNMP Server (eg: true)
}
type RequestNetworkSettingsCreateNetworkV1SettingsNetflowcollector struct {
	IPAddress string `json:"ipAddress,omitempty"` // IP Address for NetFlow collector (eg: 3.3.3.1)

	Port *float64 `json:"port,omitempty"` // Port for NetFlow Collector (eg; 443)
}
type RequestNetworkSettingsCreateNetworkV1SettingsMessageOfTheday struct {
	BannerMessage string `json:"bannerMessage,omitempty"` // Massage for Banner message (eg; Good day)

	RetainExistingBanner string `json:"retainExistingBanner,omitempty"` // Retain existing Banner Message (eg: "true" or "false")
}
type RequestNetworkSettingsCreateNetworkV1SettingsNetworkAAA struct {
	Servers string `json:"servers,omitempty"` // Server type for AAA Network (eg: AAA)

	IPAddress string `json:"ipAddress,omitempty"` // IP address for AAA and ISE server (eg: 1.1.1.1)

	Network string `json:"network,omitempty"` // IP Address for AAA or ISE server (eg: 2.2.2.2)

	Protocol string `json:"protocol,omitempty"` // Protocol for AAA or ISE serve (eg: RADIUS)

	SharedSecret string `json:"sharedSecret,omitempty"` // Shared secret for ISE Server
}
type RequestNetworkSettingsCreateNetworkV1SettingsClientAndEndpointAAA struct {
	Servers string `json:"servers,omitempty"` // Server type AAA or ISE server (eg: AAA)

	IPAddress string `json:"ipAddress,omitempty"` // IP address for ISE serve (eg: 1.1.1.4)

	Network string `json:"network,omitempty"` // IP address for AAA or ISE server (eg: 2.2.2.1)

	Protocol string `json:"protocol,omitempty"` // Protocol for AAA or ISE serve (eg: RADIUS)

	SharedSecret string `json:"sharedSecret,omitempty"` // Shared secret for ISE server
}
type RequestNetworkSettingsUpdateNetworkV1 struct {
	Settings *RequestNetworkSettingsUpdateNetworkV1Settings `json:"settings,omitempty"` //
}
type RequestNetworkSettingsUpdateNetworkV1Settings struct {
	DhcpServer []string `json:"dhcpServer,omitempty"` // DHCP Server IP (eg: 1.1.1.1)

	DNSServer *RequestNetworkSettingsUpdateNetworkV1SettingsDNSServer `json:"dnsServer,omitempty"` //

	SyslogServer *RequestNetworkSettingsUpdateNetworkV1SettingsSyslogServer `json:"syslogServer,omitempty"` //

	SNMPServer *RequestNetworkSettingsUpdateNetworkV1SettingsSNMPServer `json:"snmpServer,omitempty"` //

	Netflowcollector *RequestNetworkSettingsUpdateNetworkV1SettingsNetflowcollector `json:"netflowcollector,omitempty"` //

	NtpServer []string `json:"ntpServer,omitempty"` // IP address for NTP server (eg: 1.1.1.2)

	Timezone string `json:"timezone,omitempty"` // Input for time zone (eg: Africa/Abidjan)

	MessageOfTheday *RequestNetworkSettingsUpdateNetworkV1SettingsMessageOfTheday `json:"messageOfTheday,omitempty"` //

	NetworkAAA *RequestNetworkSettingsUpdateNetworkV1SettingsNetworkAAA `json:"network_aaa,omitempty"` //

	ClientAndEndpointAAA *RequestNetworkSettingsUpdateNetworkV1SettingsClientAndEndpointAAA `json:"clientAndEndpoint_aaa,omitempty"` //
}
type RequestNetworkSettingsUpdateNetworkV1SettingsDNSServer struct {
	DomainName string `json:"domainName,omitempty"` // Domain Name of DHCP (eg; cisco)

	PrimaryIPAddress string `json:"primaryIpAddress,omitempty"` // Primary IP Address for DHCP (eg: 2.2.2.2)

	SecondaryIPAddress string `json:"secondaryIpAddress,omitempty"` // Secondary IP Address for DHCP (eg: 3.3.3.3)
}
type RequestNetworkSettingsUpdateNetworkV1SettingsSyslogServer struct {
	IPAddresses []string `json:"ipAddresses,omitempty"` // IP Address for syslog server (eg: 4.4.4.4)

	ConfigureDnacIP *bool `json:"configureDnacIP,omitempty"` // Configuration DNAC IP for syslog server (eg: true)
}
type RequestNetworkSettingsUpdateNetworkV1SettingsSNMPServer struct {
	IPAddresses []string `json:"ipAddresses,omitempty"` // IP Address for SNMP Server (eg: 4.4.4.1)

	ConfigureDnacIP *bool `json:"configureDnacIP,omitempty"` // Configuration DNAC IP for SNMP Server (eg: true)
}
type RequestNetworkSettingsUpdateNetworkV1SettingsNetflowcollector struct {
	IPAddress string `json:"ipAddress,omitempty"` // IP Address for NetFlow collector (eg: 3.3.3.1)

	Port *float64 `json:"port,omitempty"` // Port for NetFlow Collector (eg; 443)
}
type RequestNetworkSettingsUpdateNetworkV1SettingsMessageOfTheday struct {
	BannerMessage string `json:"bannerMessage,omitempty"` // Massage for Banner message (eg; Good day)

	RetainExistingBanner string `json:"retainExistingBanner,omitempty"` // Retain existing Banner Message (eg: "true" or "false")
}
type RequestNetworkSettingsUpdateNetworkV1SettingsNetworkAAA struct {
	Servers string `json:"servers,omitempty"` // Server type for AAA Network (eg: AAA)

	IPAddress string `json:"ipAddress,omitempty"` // IP address for AAA and ISE server (eg: 1.1.1.1)

	Network string `json:"network,omitempty"` // IP Address for AAA or ISE server (eg: 2.2.2.2)

	Protocol string `json:"protocol,omitempty"` // Protocol for AAA or ISE serve (eg: RADIUS)

	SharedSecret string `json:"sharedSecret,omitempty"` // Shared secret for ISE Server
}
type RequestNetworkSettingsUpdateNetworkV1SettingsClientAndEndpointAAA struct {
	Servers string `json:"servers,omitempty"` // Server type AAA or ISE server (eg: AAA)

	IPAddress string `json:"ipAddress,omitempty"` // IP address for ISE serve (eg: 1.1.1.4)

	Network string `json:"network,omitempty"` // IP address for AAA or ISE server (eg: 2.2.2.1)

	Protocol string `json:"protocol,omitempty"` // Protocol for AAA or ISE serve (eg: RADIUS)

	SharedSecret string `json:"sharedSecret,omitempty"` // Shared secret for ISE server
}
type RequestNetworkSettingsReserveIPSubpoolV1 struct {
	Name string `json:"name,omitempty"` // Name of the reserve ip sub pool

	Type string `json:"type,omitempty"` // Type of the reserve ip sub pool

	IPv6AddressSpace *bool `json:"ipv6AddressSpace,omitempty"` // If the value is omitted or false only ipv4 input are required, otherwise both ipv6 and ipv4 are required

	IPv4GlobalPool string `json:"ipv4GlobalPool,omitempty"` // IP v4 Global pool address with cidr, example: 175.175.0.0/16

	IPv4Prefix *bool `json:"ipv4Prefix,omitempty"` // IPv4 prefix value is true, the ip4 prefix length input field is enabled , if it is false ipv4 total Host input is enable

	IPv4PrefixLength *int `json:"ipv4PrefixLength,omitempty"` // The ipv4 prefix length is required when ipv4prefix value is true.

	IPv4Subnet string `json:"ipv4Subnet,omitempty"` // IPv4 Subnet address, example: 175.175.0.0. Either ipv4Subnet or ipv4TotalHost needs to be passed if creating IPv4 subpool.

	IPv4GateWay string `json:"ipv4GateWay,omitempty"` // Gateway ip address details, example: 175.175.0.1

	IPv4DhcpServers []string `json:"ipv4DhcpServers,omitempty"` // IPv4 input for dhcp server ip example: ["1.1.1.1"]

	IPv4DNSServers []string `json:"ipv4DnsServers,omitempty"` // IPv4 input for dns server ip example: ["4.4.4.4"]

	IPv6GlobalPool string `json:"ipv6GlobalPool,omitempty"` // IPv6 Global pool address with cidr this is required when Ipv6AddressSpace value is true, example: 2001:db8:85a3::/64

	IPv6Prefix *bool `json:"ipv6Prefix,omitempty"` // Ipv6 prefix value is true, the ip6 prefix length input field is enabled , if it is false ipv6 total Host input is enable

	IPv6PrefixLength *int `json:"ipv6PrefixLength,omitempty"` // IPv6 prefix length is required when the ipv6prefix value is true

	IPv6Subnet string `json:"ipv6Subnet,omitempty"` // IPv6 Subnet address, example :2001:db8:85a3:0:100::. Either ipv6Subnet or ipv6TotalHost needs to be passed if creating IPv6 subpool.

	IPv6GateWay string `json:"ipv6GateWay,omitempty"` // Gateway ip address details, example: 2001:db8:85a3:0:100::1

	IPv6DhcpServers []string `json:"ipv6DhcpServers,omitempty"` // IPv6 format dhcp server as input example : ["2001:db8::1234"]

	IPv6DNSServers []string `json:"ipv6DnsServers,omitempty"` // IPv6 format dns server input example: ["2001:db8::1234"]

	IPv4TotalHost *int `json:"ipv4TotalHost,omitempty"` // IPv4 total host is required when ipv4prefix value is false.

	IPv6TotalHost *int `json:"ipv6TotalHost,omitempty"` // IPv6 total host is required when ipv6prefix value is false.

	SLAacSupport *bool `json:"slaacSupport,omitempty"` // Slaac Support
}
type RequestNetworkSettingsUpdateReserveIPSubpoolV1 struct {
	Name string `json:"name,omitempty"` // Name of the reserve ip sub pool

	IPv6AddressSpace *bool `json:"ipv6AddressSpace,omitempty"` // If the value is false only ipv4 input are required. NOTE if value is false then any existing ipv6 subpool in the group will be removed.

	IPv4DhcpServers []string `json:"ipv4DhcpServers,omitempty"` // IPv4 input for dhcp server ip example: ["1.1.1.1"]

	IPv4DNSServers []string `json:"ipv4DnsServers,omitempty"` // IPv4 input for dns server ip  example: ["4.4.4.4"]

	IPv6GlobalPool string `json:"ipv6GlobalPool,omitempty"` // IPv6 Global pool address with cidr this is required when Ipv6AddressSpace value is true, example: 2001:db8:85a3::/64

	IPv6Prefix *bool `json:"ipv6Prefix,omitempty"` // Ipv6 prefix value is true, the ip6 prefix length input field is enabled, if it is false ipv6 total Host input is enable

	IPv6PrefixLength *int `json:"ipv6PrefixLength,omitempty"` // IPv6 prefix length is required when the ipv6prefix value is true

	IPv6Subnet string `json:"ipv6Subnet,omitempty"` // IPv6 Subnet address, example :2001:db8:85a3:0:100::.

	IPv6TotalHost *int `json:"ipv6TotalHost,omitempty"` // Size of pool in terms of number of IPs. IPv6 total host is required when ipv6prefix value is false.

	IPv6GateWay string `json:"ipv6GateWay,omitempty"` // Gateway ip address details, example: 2001:db8:85a3:0:100::1

	IPv6DhcpServers []string `json:"ipv6DhcpServers,omitempty"` // IPv6 format dhcp server as input example : ["2001:db8::1234"]

	IPv6DNSServers []string `json:"ipv6DnsServers,omitempty"` // IPv6 format dns server input example: ["2001:db8::1234"]

	SLAacSupport *bool `json:"slaacSupport,omitempty"` // Slaac Support

	IPv4GateWay string `json:"ipv4GateWay,omitempty"` // Gateway ip address details, example: 175.175.0.1
}
type RequestNetworkSettingsCreateSpProfileV1 struct {
	Settings *RequestNetworkSettingsCreateSpProfileV1Settings `json:"settings,omitempty"` //
}
type RequestNetworkSettingsCreateSpProfileV1Settings struct {
	Qos *[]RequestNetworkSettingsCreateSpProfileV1SettingsQos `json:"qos,omitempty"` //
}
type RequestNetworkSettingsCreateSpProfileV1SettingsQos struct {
	ProfileName string `json:"profileName,omitempty"` // Profile Name

	Model string `json:"model,omitempty"` // Model

	WanProvider string `json:"wanProvider,omitempty"` // Wan Provider
}
type RequestNetworkSettingsUpdateSpProfileV1 struct {
	Settings *RequestNetworkSettingsUpdateSpProfileV1Settings `json:"settings,omitempty"` //
}
type RequestNetworkSettingsUpdateSpProfileV1Settings struct {
	Qos *[]RequestNetworkSettingsUpdateSpProfileV1SettingsQos `json:"qos,omitempty"` //
}
type RequestNetworkSettingsUpdateSpProfileV1SettingsQos struct {
	ProfileName string `json:"profileName,omitempty"` // Profile Name

	Model string `json:"model,omitempty"` // Model

	WanProvider string `json:"wanProvider,omitempty"` // Wan Provider

	OldProfileName string `json:"oldProfileName,omitempty"` // Old Profile Name
}
type RequestNetworkSettingsSyncNetworkDevicesCredentialV1 struct {
	DeviceCredentialID string `json:"deviceCredentialId,omitempty"` // It must be cli/snmpV2Read/snmpV2Write/snmpV3 Id.

	SiteID string `json:"siteId,omitempty"` // Site Id.
}
type RequestNetworkSettingsSetAAASettingsForASiteV1 struct {
	AAANetwork *RequestNetworkSettingsSetAAASettingsForASiteV1AAANetwork `json:"aaaNetwork,omitempty"` //

	AAAClient *RequestNetworkSettingsSetAAASettingsForASiteV1AAAClient `json:"aaaClient,omitempty"` //
}
type RequestNetworkSettingsSetAAASettingsForASiteV1AAANetwork struct {
	ServerType string `json:"serverType,omitempty"` // Server Type

	Protocol string `json:"protocol,omitempty"` // Protocol

	Pan string `json:"pan,omitempty"` // Administration Node. Required for ISE.

	PrimaryServerIP string `json:"primaryServerIp,omitempty"` // The server to use as a primary.

	SecondaryServerIP string `json:"secondaryServerIp,omitempty"` // The server to use as a secondary.

	SharedSecret string `json:"sharedSecret,omitempty"` // Shared Secret
}
type RequestNetworkSettingsSetAAASettingsForASiteV1AAAClient struct {
	ServerType string `json:"serverType,omitempty"` // Server Type

	Protocol string `json:"protocol,omitempty"` // Protocol

	Pan string `json:"pan,omitempty"` // Administration Node.  Required for ISE.

	PrimaryServerIP string `json:"primaryServerIp,omitempty"` // The server to use as a primary.

	SecondaryServerIP string `json:"secondaryServerIp,omitempty"` // The server to use as a secondary.

	SharedSecret string `json:"sharedSecret,omitempty"` // Shared Secret
}
type RequestNetworkSettingsSetBannerSettingsForASiteV1 struct {
	Banner *RequestNetworkSettingsSetBannerSettingsForASiteV1Banner `json:"banner,omitempty"` //
}
type RequestNetworkSettingsSetBannerSettingsForASiteV1Banner struct {
	Type string `json:"type,omitempty"` // Type

	Message string `json:"message,omitempty"` // Custom message that appears when logging into routers, switches, and hubs. Required for custom type.
}
type RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteV1 struct {
	CliCredentialsID *RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteV1CliCredentialsID `json:"cliCredentialsId,omitempty"` //

	SNMPv2CReadCredentialsID *RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteV1SNMPv2CReadCredentialsID `json:"snmpv2cReadCredentialsId,omitempty"` //

	SNMPv2CWriteCredentialsID *RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteV1SNMPv2CWriteCredentialsID `json:"snmpv2cWriteCredentialsId,omitempty"` //

	SNMPv3CredentialsID *RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteV1SNMPv3CredentialsID `json:"snmpv3CredentialsId,omitempty"` //

	HTTPReadCredentialsID *RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteV1HTTPReadCredentialsID `json:"httpReadCredentialsId,omitempty"` //

	HTTPWriteCredentialsID *RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteV1HTTPWriteCredentialsID `json:"httpWriteCredentialsId,omitempty"` //
}
type RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteV1CliCredentialsID struct {
	CredentialsID string `json:"credentialsId,omitempty"` // The `id` of the credentials.
}
type RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteV1SNMPv2CReadCredentialsID struct {
	CredentialsID string `json:"credentialsId,omitempty"` // The `id` of the credentials.
}
type RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteV1SNMPv2CWriteCredentialsID struct {
	CredentialsID string `json:"credentialsId,omitempty"` // The `id` of the credentials.
}
type RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteV1SNMPv3CredentialsID struct {
	CredentialsID string `json:"credentialsId,omitempty"` // The `id` of the credentials.
}
type RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteV1HTTPReadCredentialsID struct {
	CredentialsID string `json:"credentialsId,omitempty"` // The `id` of the credentials.
}
type RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteV1HTTPWriteCredentialsID struct {
	CredentialsID string `json:"credentialsId,omitempty"` // The `id` of the credentials.
}
type RequestNetworkSettingsSetDhcpSettingsForASiteV1 struct {
	Dhcp *RequestNetworkSettingsSetDhcpSettingsForASiteV1Dhcp `json:"dhcp,omitempty"` //
}
type RequestNetworkSettingsSetDhcpSettingsForASiteV1Dhcp struct {
	Servers []string `json:"servers,omitempty"` // DHCP servers for managing client device networking configuration. Max:10
}
type RequestNetworkSettingsSetDNSSettingsForASiteV1 struct {
	DNS *RequestNetworkSettingsSetDNSSettingsForASiteV1DNS `json:"dns,omitempty"` //
}
type RequestNetworkSettingsSetDNSSettingsForASiteV1DNS struct {
	DomainName string   `json:"domainName,omitempty"` // Network's domain name. Example : myCompnay.com
	DNSServers []string `json:"dnsServers,omitempty"` // DNS servers for hostname resolution.
}
type RequestNetworkSettingsSetImageDistributionSettingsForASiteV1 struct {
	ImageDistribution *RequestNetworkSettingsSetImageDistributionSettingsForASiteV1ImageDistribution `json:"imageDistribution,omitempty"` //
}
type RequestNetworkSettingsSetImageDistributionSettingsForASiteV1ImageDistribution struct {
	Servers []string `json:"servers,omitempty"` // This field holds an array of unique identifiers representing image distribution servers. Use ‘/intent/api/v1/images/distributionServerSettings’ to find the Image distribution server Id. Max:2. Use SFTP servers to act as image distribution servers. A distributed SWIM architecture, using suitably located SFTP servers, can help support large-scale device software image upgrades and conserve WAN bandwidth.
}
type RequestNetworkSettingsSetNTPSettingsForASiteV1 struct {
	Ntp *RequestNetworkSettingsSetNTPSettingsForASiteV1Ntp `json:"ntp,omitempty"` //
}
type RequestNetworkSettingsSetNTPSettingsForASiteV1Ntp struct {
	Servers []string `json:"servers,omitempty"` // NTP servers to facilitate system clock synchronization for your network. Max:10
}
type RequestNetworkSettingsSetTelemetrySettingsForASiteV1 struct {
	WiredDataCollection *RequestNetworkSettingsSetTelemetrySettingsForASiteV1WiredDataCollection `json:"wiredDataCollection,omitempty"` //

	WirelessTelemetry *RequestNetworkSettingsSetTelemetrySettingsForASiteV1WirelessTelemetry `json:"wirelessTelemetry,omitempty"` //

	SNMPTraps *RequestNetworkSettingsSetTelemetrySettingsForASiteV1SNMPTraps `json:"snmpTraps,omitempty"` //

	Syslogs *RequestNetworkSettingsSetTelemetrySettingsForASiteV1Syslogs `json:"syslogs,omitempty"` //

	ApplicationVisibility *RequestNetworkSettingsSetTelemetrySettingsForASiteV1ApplicationVisibility `json:"applicationVisibility,omitempty"` //
}
type RequestNetworkSettingsSetTelemetrySettingsForASiteV1WiredDataCollection struct {
	EnableWiredDataCollectio *bool `json:"enableWiredDataCollectio,omitempty"` // Track the presence, location, and movement of wired endpoints in the network. Traffic received from endpoints is used to extract and store their identity information (MAC address and IP address). Other features, such as IEEE 802.1X, web authentication, Cisco Security Groups (formerly TrustSec), SD-Access, and Assurance, depend on this identity information to operate properly. Wired Endpoint Data Collection enables Device Tracking policies on devices assigned to the Access role in Inventory.
}
type RequestNetworkSettingsSetTelemetrySettingsForASiteV1WirelessTelemetry struct {
	EnableWirelessTelemetry *bool `json:"enableWirelessTelemetry,omitempty"` // Enables Streaming Telemetry on your wireless controllers in order to determine the health of your wireless controller, access points and wireless clients.
}
type RequestNetworkSettingsSetTelemetrySettingsForASiteV1SNMPTraps struct {
	UseBuiltinTrapServer *bool `json:"useBuiltinTrapServer,omitempty"` // Enable this server as a destination server for SNMP traps and messages from your network

	ExternalTrapServers []string `json:"externalTrapServers,omitempty"` // External SNMP trap servers. Example: ["250.162.252.170","2001:db8:3c4d:15::1a2f:1a2b"]
}
type RequestNetworkSettingsSetTelemetrySettingsForASiteV1Syslogs struct {
	UseBuiltinSyslogServer *bool `json:"useBuiltinSyslogServer,omitempty"` // Enable this server as a destination server for syslog messages.

	ExternalSyslogServers []string `json:"externalSyslogServers,omitempty"` // External syslog servers. Example: ["250.162.252.170", "2001:db8:3c4d:15::1a2f:1a2b"]
}
type RequestNetworkSettingsSetTelemetrySettingsForASiteV1ApplicationVisibility struct {
	Collector *RequestNetworkSettingsSetTelemetrySettingsForASiteV1ApplicationVisibilityCollector `json:"collector,omitempty"` //

	EnableOnWiredAccessDevices *bool `json:"enableOnWiredAccessDevices,omitempty"` // Enable Netflow Application Telemetry and Controller Based Application Recognition (CBAR) by default upon network device site assignment for wired access devices.
}
type RequestNetworkSettingsSetTelemetrySettingsForASiteV1ApplicationVisibilityCollector struct {
	CollectorType string `json:"collectorType,omitempty"` // Collector Type

	Address string `json:"address,omitempty"` // IP Address. If collection type is 'TelemetryBrokerOrUDPDirector', this field value is mandatory otherwise it is optional. Examples: "250.162.252.170", "2001:db8:3c4d:15::1a2f:1a2b"

	Port *int `json:"port,omitempty"` // Min:1; Max: 65535. If collection type is 'TelemetryBrokerOrUDPDirector', this field value is mandatory otherwise it is optional.
}
type RequestNetworkSettingsSetTimeZoneForASiteV1 struct {
	TimeZone *RequestNetworkSettingsSetTimeZoneForASiteV1TimeZone `json:"timeZone,omitempty"` //
}
type RequestNetworkSettingsSetTimeZoneForASiteV1TimeZone struct {
	IDentifier string `json:"identifier,omitempty"` // Time zone that corresponds to the site's physical location. The site time zone is used when scheduling device provisioning and updates. Example: GMT
}
type RequestNetworkSettingsUpdateADevicesTelemetrySettingsToConformToTheTelemetrySettingsForItsSiteV1 struct {
	DeviceIDs []string `json:"deviceIds,omitempty"` // The list of device Ids to perform the provisioning against
}
type RequestNetworkSettingsAssignDeviceCredentialToSiteV2 struct {
	CliID string `json:"cliId,omitempty"` // CLI Credential Id

	SNMPV2ReadID string `json:"snmpV2ReadId,omitempty"` // SNMPv2c Read Credential Id

	SNMPV2WriteID string `json:"snmpV2WriteId,omitempty"` // SNMPv2c Write Credential Id

	SNMPV3ID string `json:"snmpV3Id,omitempty"` // SNMPv3 Credential Id

	HTTPRead string `json:"httpRead,omitempty"` // HTTP(S) Read Credential Id

	HTTPWrite string `json:"httpWrite,omitempty"` // HTTP(S) Write Credential Id
}
type RequestNetworkSettingsCreateNetworkV2 struct {
	Settings *RequestNetworkSettingsCreateNetworkV2Settings `json:"settings,omitempty"` //
}
type RequestNetworkSettingsCreateNetworkV2Settings struct {
	DhcpServer []string `json:"dhcpServer,omitempty"` // DHCP Server IP (eg: 1.1.1.1)

	DNSServer *RequestNetworkSettingsCreateNetworkV2SettingsDNSServer `json:"dnsServer,omitempty"` //

	SyslogServer *RequestNetworkSettingsCreateNetworkV2SettingsSyslogServer `json:"syslogServer,omitempty"` //

	SNMPServer *RequestNetworkSettingsCreateNetworkV2SettingsSNMPServer `json:"snmpServer,omitempty"` //

	Netflowcollector *RequestNetworkSettingsCreateNetworkV2SettingsNetflowcollector `json:"netflowcollector,omitempty"` //

	NtpServer []string `json:"ntpServer,omitempty"` // IP address for NTP server (eg: 1.1.1.2)

	Timezone string `json:"timezone,omitempty"` // Input for time zone (eg: Africa/Abidjan)

	MessageOfTheday *RequestNetworkSettingsCreateNetworkV2SettingsMessageOfTheday `json:"messageOfTheday,omitempty"` //

	NetworkAAA *RequestNetworkSettingsCreateNetworkV2SettingsNetworkAAA `json:"network_aaa,omitempty"` //

	ClientAndEndpointAAA *RequestNetworkSettingsCreateNetworkV2SettingsClientAndEndpointAAA `json:"clientAndEndpoint_aaa,omitempty"` //
}
type RequestNetworkSettingsCreateNetworkV2SettingsDNSServer struct {
	DomainName string `json:"domainName,omitempty"` // Domain Name of DHCP (eg; cisco)

	PrimaryIPAddress string `json:"primaryIpAddress,omitempty"` // Primary IP Address for DHCP (eg: 2.2.2.2)

	SecondaryIPAddress string `json:"secondaryIpAddress,omitempty"` // Secondary IP Address for DHCP (eg: 3.3.3.3)
}
type RequestNetworkSettingsCreateNetworkV2SettingsSyslogServer struct {
	IPAddresses []string `json:"ipAddresses,omitempty"` // IP Address for syslog server (eg: 4.4.4.4)

	ConfigureDnacIP *bool `json:"configureDnacIP,omitempty"` // Configuration DNAC IP for syslog server (eg: true)
}
type RequestNetworkSettingsCreateNetworkV2SettingsSNMPServer struct {
	IPAddresses []string `json:"ipAddresses,omitempty"` // IP Address for SNMP Server (eg: 4.4.4.1)

	ConfigureDnacIP *bool `json:"configureDnacIP,omitempty"` // Configuration DNAC IP for SNMP Server (eg: true)
}
type RequestNetworkSettingsCreateNetworkV2SettingsNetflowcollector struct {
	IPAddress string `json:"ipAddress,omitempty"` // IP Address for NetFlow collector (eg: 3.3.3.1)

	Port *float64 `json:"port,omitempty"` // Port for NetFlow Collector (eg; 443)
}
type RequestNetworkSettingsCreateNetworkV2SettingsMessageOfTheday struct {
	BannerMessage string `json:"bannerMessage,omitempty"` // Massage for Banner message (eg; Good day)

	RetainExistingBanner string `json:"retainExistingBanner,omitempty"` // Retain existing Banner Message (eg: "true" or "false")
}
type RequestNetworkSettingsCreateNetworkV2SettingsNetworkAAA struct {
	Servers string `json:"servers,omitempty"` // Server type for AAA Network (eg: AAA)

	IPAddress string `json:"ipAddress,omitempty"` // IP address for AAA and ISE server (eg: 1.1.1.1)

	Network string `json:"network,omitempty"` // IP Address for AAA or ISE server (eg: 2.2.2.2)

	Protocol string `json:"protocol,omitempty"` // Protocol for AAA or ISE serve (eg: RADIUS)

	SharedSecret string `json:"sharedSecret,omitempty"` // Shared secret for ISE Server
}
type RequestNetworkSettingsCreateNetworkV2SettingsClientAndEndpointAAA struct {
	Servers string `json:"servers,omitempty"` // Server type AAA or ISE server (eg: AAA)

	IPAddress string `json:"ipAddress,omitempty"` // IP address for ISE serve (eg: 1.1.1.4)

	Network string `json:"network,omitempty"` // IP address for AAA or ISE server (eg: 2.2.2.1)

	Protocol string `json:"protocol,omitempty"` // Protocol for AAA or ISE serve (eg: RADIUS)

	SharedSecret string `json:"sharedSecret,omitempty"` // Shared secret for ISE server
}
type RequestNetworkSettingsUpdateNetworkV2 struct {
	Settings *RequestNetworkSettingsUpdateNetworkV2Settings `json:"settings,omitempty"` //
}
type RequestNetworkSettingsUpdateNetworkV2Settings struct {
	DhcpServer []string `json:"dhcpServer,omitempty"` // DHCP Server IP (eg: 1.1.1.1)

	DNSServer *RequestNetworkSettingsUpdateNetworkV2SettingsDNSServer `json:"dnsServer,omitempty"` //

	SyslogServer *RequestNetworkSettingsUpdateNetworkV2SettingsSyslogServer `json:"syslogServer,omitempty"` //

	SNMPServer *RequestNetworkSettingsUpdateNetworkV2SettingsSNMPServer `json:"snmpServer,omitempty"` //

	Netflowcollector *RequestNetworkSettingsUpdateNetworkV2SettingsNetflowcollector `json:"netflowcollector,omitempty"` //

	NtpServer []string `json:"ntpServer,omitempty"` // IP address for NTP server (eg: 1.1.1.2)

	Timezone string `json:"timezone,omitempty"` // Input for time zone (eg: Africa/Abidjan)

	MessageOfTheday *RequestNetworkSettingsUpdateNetworkV2SettingsMessageOfTheday `json:"messageOfTheday,omitempty"` //

	NetworkAAA *RequestNetworkSettingsUpdateNetworkV2SettingsNetworkAAA `json:"network_aaa,omitempty"` //

	ClientAndEndpointAAA *RequestNetworkSettingsUpdateNetworkV2SettingsClientAndEndpointAAA `json:"clientAndEndpoint_aaa,omitempty"` //
}
type RequestNetworkSettingsUpdateNetworkV2SettingsDNSServer struct {
	DomainName string `json:"domainName,omitempty"` // Domain Name of DHCP (eg; cisco)

	PrimaryIPAddress string `json:"primaryIpAddress,omitempty"` // Primary IP Address for DHCP (eg: 2.2.2.2)

	SecondaryIPAddress string `json:"secondaryIpAddress,omitempty"` // Secondary IP Address for DHCP (eg: 3.3.3.3)
}
type RequestNetworkSettingsUpdateNetworkV2SettingsSyslogServer struct {
	IPAddresses []string `json:"ipAddresses,omitempty"` // IP Address for syslog server (eg: 4.4.4.4)

	ConfigureDnacIP *bool `json:"configureDnacIP,omitempty"` // Configuration DNAC IP for syslog server (eg: true)
}
type RequestNetworkSettingsUpdateNetworkV2SettingsSNMPServer struct {
	IPAddresses []string `json:"ipAddresses,omitempty"` // IP Address for SNMP Server (eg: 4.4.4.1)

	ConfigureDnacIP *bool `json:"configureDnacIP,omitempty"` // Configuration DNAC IP for SNMP Server (eg: true)
}
type RequestNetworkSettingsUpdateNetworkV2SettingsNetflowcollector struct {
	IPAddress string `json:"ipAddress,omitempty"` // IP Address for NetFlow collector (eg: 3.3.3.1)

	Port *float64 `json:"port,omitempty"` // Port for NetFlow Collector (eg; 443)
}
type RequestNetworkSettingsUpdateNetworkV2SettingsMessageOfTheday struct {
	BannerMessage string `json:"bannerMessage,omitempty"` // Massage for Banner message (eg; Good day)

	RetainExistingBanner string `json:"retainExistingBanner,omitempty"` // Retain existing Banner Message (eg: "true" or "false")
}
type RequestNetworkSettingsUpdateNetworkV2SettingsNetworkAAA struct {
	Servers string `json:"servers,omitempty"` // Server type for AAA Network (eg: AAA)

	IPAddress string `json:"ipAddress,omitempty"` // IP address for AAA and ISE server (eg: 1.1.1.1)

	Network string `json:"network,omitempty"` // IP Address for AAA or ISE server (eg: 2.2.2.2)

	Protocol string `json:"protocol,omitempty"` // Protocol for AAA or ISE serve (eg: RADIUS)

	SharedSecret string `json:"sharedSecret,omitempty"` // Shared secret for ISE Server
}
type RequestNetworkSettingsUpdateNetworkV2SettingsClientAndEndpointAAA struct {
	Servers string `json:"servers,omitempty"` // Server type AAA or ISE server (eg: AAA)

	IPAddress string `json:"ipAddress,omitempty"` // IP address for ISE serve (eg: 1.1.1.4)

	Network string `json:"network,omitempty"` // IP address for AAA or ISE server (eg: 2.2.2.1)

	Protocol string `json:"protocol,omitempty"` // Protocol for AAA or ISE serve (eg: RADIUS)

	SharedSecret string `json:"sharedSecret,omitempty"` // Shared secret for ISE server
}
type RequestNetworkSettingsCreateSpProfileV2 struct {
	Settings *RequestNetworkSettingsCreateSpProfileV2Settings `json:"settings,omitempty"` //
}
type RequestNetworkSettingsCreateSpProfileV2Settings struct {
	Qos *[]RequestNetworkSettingsCreateSpProfileV2SettingsQos `json:"qos,omitempty"` //
}
type RequestNetworkSettingsCreateSpProfileV2SettingsQos struct {
	ProfileName string `json:"profileName,omitempty"` // Profile Name

	Model string `json:"model,omitempty"` // Model

	WanProvider string `json:"wanProvider,omitempty"` // Wan Provider
}
type RequestNetworkSettingsUpdateSpProfileV2 struct {
	Settings *RequestNetworkSettingsUpdateSpProfileV2Settings `json:"settings,omitempty"` //
}
type RequestNetworkSettingsUpdateSpProfileV2Settings struct {
	Qos *[]RequestNetworkSettingsUpdateSpProfileV2SettingsQos `json:"qos,omitempty"` //
}
type RequestNetworkSettingsUpdateSpProfileV2SettingsQos struct {
	ProfileName string `json:"profileName,omitempty"` // Profile Name

	Model string `json:"model,omitempty"` // Model

	WanProvider string `json:"wanProvider,omitempty"` // Wan Provider

	OldProfileName string `json:"oldProfileName,omitempty"` // Old Profile Name
}

//GetDeviceCredentialDetailsV1 Get Device Credential Details - 899f-08e7-401b-82dd
/* API to get device credential details.


@param GetDeviceCredentialDetailsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-credential-details
*/
func (s *NetworkSettingsService) GetDeviceCredentialDetailsV1(GetDeviceCredentialDetailsV1QueryParams *GetDeviceCredentialDetailsV1QueryParams) (*ResponseNetworkSettingsGetDeviceCredentialDetailsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/device-credential"

	queryString, _ := query.Values(GetDeviceCredentialDetailsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworkSettingsGetDeviceCredentialDetailsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceCredentialDetailsV1(GetDeviceCredentialDetailsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceCredentialDetailsV1")
	}

	result := response.Result().(*ResponseNetworkSettingsGetDeviceCredentialDetailsV1)
	return result, response, err

}

//GetGlobalPoolV1 Get Global Pool - c0bc-a856-43c8-b58d
/* API to get the global pool.


@param GetGlobalPoolV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-global-pool
*/
func (s *NetworkSettingsService) GetGlobalPoolV1(GetGlobalPoolV1QueryParams *GetGlobalPoolV1QueryParams) (*ResponseNetworkSettingsGetGlobalPoolV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-pool"

	queryString, _ := query.Values(GetGlobalPoolV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworkSettingsGetGlobalPoolV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetGlobalPoolV1(GetGlobalPoolV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetGlobalPoolV1")
	}

	result := response.Result().(*ResponseNetworkSettingsGetGlobalPoolV1)
	return result, response, err

}

//RetrievesGlobalIPAddressPoolsV1 Retrieves global IP address pools. - 8389-eba4-402a-9892
/* Retrieves global IP address pools. Global pools are not associated with any particular site, but may have portions of their address space reserved by site-specific subpools.


@param RetrievesGlobalIPAddressPoolsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-global-ip-address-pools
*/
func (s *NetworkSettingsService) RetrievesGlobalIPAddressPoolsV1(RetrievesGlobalIPAddressPoolsV1QueryParams *RetrievesGlobalIPAddressPoolsV1QueryParams) (*ResponseNetworkSettingsRetrievesGlobalIPAddressPoolsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/ipam/globalIpAddressPools"

	queryString, _ := query.Values(RetrievesGlobalIPAddressPoolsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworkSettingsRetrievesGlobalIPAddressPoolsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesGlobalIPAddressPoolsV1(RetrievesGlobalIPAddressPoolsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesGlobalIpAddressPoolsV1")
	}

	result := response.Result().(*ResponseNetworkSettingsRetrievesGlobalIPAddressPoolsV1)
	return result, response, err

}

//CountsGlobalIPAddressPoolsV1 Counts global IP address pools. - e0bb-eaf9-4c08-8683
/* Counts global IP address pools. Global pools are not associated with any particular site, but may have portions of their address space reserved by site-specific subpools.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!counts-global-ip-address-pools
*/
func (s *NetworkSettingsService) CountsGlobalIPAddressPoolsV1() (*ResponseNetworkSettingsCountsGlobalIPAddressPoolsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/ipam/globalIpAddressPools/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkSettingsCountsGlobalIPAddressPoolsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.CountsGlobalIPAddressPoolsV1()
		}
		return nil, response, fmt.Errorf("error with operation CountsGlobalIpAddressPoolsV1")
	}

	result := response.Result().(*ResponseNetworkSettingsCountsGlobalIPAddressPoolsV1)
	return result, response, err

}

//RetrievesSubpoolsIDsOfAGlobalIPAddressPoolV1 Retrieves subpools IDs of a global IP address pool. - 8c8d-4821-4c4a-b0fa
/* Retrieves subpools IDs of a global IP address pool.  The IDs can be fetched with `/dna/intent/api/v1/ipam/siteIpAddressPools/{id}`


@param globalIPAddressPoolID globalIpAddressPoolId path parameter. The `id` of the global IP address pool for which to retrieve subpool IDs.

@param RetrievesSubpoolsIDsOfAGlobalIPAddressPoolV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-subpools-ids-of-a-global-ip-address-pool
*/
func (s *NetworkSettingsService) RetrievesSubpoolsIDsOfAGlobalIPAddressPoolV1(globalIPAddressPoolID string, RetrievesSubpoolsIDsOfAGlobalIPAddressPoolV1QueryParams *RetrievesSubpoolsIDsOfAGlobalIPAddressPoolV1QueryParams) (*ResponseNetworkSettingsRetrievesSubpoolsIDsOfAGlobalIPAddressPoolV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/ipam/globalIpAddressPools/{globalIpAddressPoolId}/subpools"
	path = strings.Replace(path, "{globalIpAddressPoolId}", fmt.Sprintf("%v", globalIPAddressPoolID), -1)

	queryString, _ := query.Values(RetrievesSubpoolsIDsOfAGlobalIPAddressPoolV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworkSettingsRetrievesSubpoolsIDsOfAGlobalIPAddressPoolV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesSubpoolsIDsOfAGlobalIPAddressPoolV1(globalIPAddressPoolID, RetrievesSubpoolsIDsOfAGlobalIPAddressPoolV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesSubpoolsIdsOfAGlobalIpAddressPoolV1")
	}

	result := response.Result().(*ResponseNetworkSettingsRetrievesSubpoolsIDsOfAGlobalIPAddressPoolV1)
	return result, response, err

}

//CountsSubpoolsOfAGlobalIPAddressPoolV1 Counts subpools of a global IP address pool. - 3a9b-4bb7-4d5a-b5e1
/* Counts subpools of a global IP address pool.


@param globalIPAddressPoolID globalIpAddressPoolId path parameter. The `id` of the global IP address pool for which to count subpools.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!counts-subpools-of-a-global-ip-address-pool
*/
func (s *NetworkSettingsService) CountsSubpoolsOfAGlobalIPAddressPoolV1(globalIPAddressPoolID string) (*ResponseNetworkSettingsCountsSubpoolsOfAGlobalIPAddressPoolV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/ipam/globalIpAddressPools/{globalIpAddressPoolId}/subpools/count"
	path = strings.Replace(path, "{globalIpAddressPoolId}", fmt.Sprintf("%v", globalIPAddressPoolID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkSettingsCountsSubpoolsOfAGlobalIPAddressPoolV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.CountsSubpoolsOfAGlobalIPAddressPoolV1(globalIPAddressPoolID)
		}
		return nil, response, fmt.Errorf("error with operation CountsSubpoolsOfAGlobalIpAddressPoolV1")
	}

	result := response.Result().(*ResponseNetworkSettingsCountsSubpoolsOfAGlobalIPAddressPoolV1)
	return result, response, err

}

//RetrievesAGlobalIPAddressPoolV1 Retrieves a global IP address pool. - c389-6867-476b-b9a2
/* Retrieves a global IP address pool. Global pools are not associated with any particular site, but may have portions of their address space reserved by site-specific subpools.


@param id id path parameter. The `id` of the global IP address pool to retrieve.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-a-global-ip-address-pool
*/
func (s *NetworkSettingsService) RetrievesAGlobalIPAddressPoolV1(id string) (*ResponseNetworkSettingsRetrievesAGlobalIPAddressPoolV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/ipam/globalIpAddressPools/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkSettingsRetrievesAGlobalIPAddressPoolV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesAGlobalIPAddressPoolV1(id)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesAGlobalIpAddressPoolV1")
	}

	result := response.Result().(*ResponseNetworkSettingsRetrievesAGlobalIPAddressPoolV1)
	return result, response, err

}

//RetrievesIPAddressSubpoolsV1 Retrieves IP address subpools. - aeb2-38a2-4249-a7a0
/* Retrieves IP address subpools, which reserve address space from a global pool (or global pools).


@param RetrievesIPAddressSubpoolsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-ip-address-subpools
*/
func (s *NetworkSettingsService) RetrievesIPAddressSubpoolsV1(RetrievesIPAddressSubpoolsV1QueryParams *RetrievesIPAddressSubpoolsV1QueryParams) (*ResponseNetworkSettingsRetrievesIPAddressSubpoolsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/ipam/siteIpAddressPools"

	queryString, _ := query.Values(RetrievesIPAddressSubpoolsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworkSettingsRetrievesIPAddressSubpoolsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesIPAddressSubpoolsV1(RetrievesIPAddressSubpoolsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesIpAddressSubpoolsV1")
	}

	result := response.Result().(*ResponseNetworkSettingsRetrievesIPAddressSubpoolsV1)
	return result, response, err

}

//CountsIPAddressSubpoolsV1 Counts IP address subpools. - 019b-fb9a-4789-ab20
/* Counts IP address subpools, which reserve address space from a global pool (or global pools).


@param CountsIPAddressSubpoolsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!counts-ip-address-subpools
*/
func (s *NetworkSettingsService) CountsIPAddressSubpoolsV1(CountsIPAddressSubpoolsV1QueryParams *CountsIPAddressSubpoolsV1QueryParams) (*ResponseNetworkSettingsCountsIPAddressSubpoolsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/ipam/siteIpAddressPools/count"

	queryString, _ := query.Values(CountsIPAddressSubpoolsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworkSettingsCountsIPAddressSubpoolsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.CountsIPAddressSubpoolsV1(CountsIPAddressSubpoolsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation CountsIpAddressSubpoolsV1")
	}

	result := response.Result().(*ResponseNetworkSettingsCountsIPAddressSubpoolsV1)
	return result, response, err

}

//RetrievesAnIPAddressSubpoolV1 Retrieves an IP address subpool. - a09a-ca0d-4f38-aada
/* Retrieves an IP address subpool, which reserves address space from a global pool (or global pools) for a particular site.


@param id id path parameter. The `id` of the IP address subpool to retrieve.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-an-ip-address-subpool
*/
func (s *NetworkSettingsService) RetrievesAnIPAddressSubpoolV1(id string) (*ResponseNetworkSettingsRetrievesAnIPAddressSubpoolV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/ipam/siteIpAddressPools/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkSettingsRetrievesAnIPAddressSubpoolV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesAnIPAddressSubpoolV1(id)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesAnIpAddressSubpoolV1")
	}

	result := response.Result().(*ResponseNetworkSettingsRetrievesAnIPAddressSubpoolV1)
	return result, response, err

}

//GetNetworkV1 Get Network - 38b7-eb13-449b-9471
/* API to get  DHCP and DNS center server details.


@param GetNetworkV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network
*/
func (s *NetworkSettingsService) GetNetworkV1(GetNetworkV1QueryParams *GetNetworkV1QueryParams) (*ResponseNetworkSettingsGetNetworkV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network"

	queryString, _ := query.Values(GetNetworkV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworkSettingsGetNetworkV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetNetworkV1(GetNetworkV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetNetworkV1")
	}

	result := response.Result().(*ResponseNetworkSettingsGetNetworkV1)
	return result, response, err

}

//RetrieveCliTemplatesAttachedToANetworkProfileV1 Retrieve CLI templates attached to a network profile - b69f-eafa-4f5b-b286
/* Retrieves a list of CLI templates attached to a network profile based on the network profile ID.


@param profileID profileId path parameter. The `id` of the network profile, retrievable from `GET /intent/api/v1/networkProfilesForSites`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-cli-templates-attached-to-a-network-profile
*/
func (s *NetworkSettingsService) RetrieveCliTemplatesAttachedToANetworkProfileV1(profileID string) (*ResponseNetworkSettingsRetrieveCliTemplatesAttachedToANetworkProfileV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkProfilesForSites/{profileId}/templates"
	path = strings.Replace(path, "{profileId}", fmt.Sprintf("%v", profileID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkSettingsRetrieveCliTemplatesAttachedToANetworkProfileV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveCliTemplatesAttachedToANetworkProfileV1(profileID)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveCliTemplatesAttachedToANetworkProfileV1")
	}

	result := response.Result().(*ResponseNetworkSettingsRetrieveCliTemplatesAttachedToANetworkProfileV1)
	return result, response, err

}

//RetrieveCountOfCliTemplatesAttachedToANetworkProfileV1 Retrieve count of CLI templates attached to a network profile - 60b7-d88a-4a48-a59a
/* Retrieves the count of all CLI templates attached to a network profile by the profile ID.


@param profileID profileId path parameter. The `id` of the network profile, retrievable from `GET /intent/api/v1/networkProfilesForSites`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-count-of-cli-templates-attached-to-a-network-profile
*/
func (s *NetworkSettingsService) RetrieveCountOfCliTemplatesAttachedToANetworkProfileV1(profileID string) (*ResponseNetworkSettingsRetrieveCountOfCliTemplatesAttachedToANetworkProfileV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkProfilesForSites/{profileId}/templates/count"
	path = strings.Replace(path, "{profileId}", fmt.Sprintf("%v", profileID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkSettingsRetrieveCountOfCliTemplatesAttachedToANetworkProfileV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveCountOfCliTemplatesAttachedToANetworkProfileV1(profileID)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveCountOfCliTemplatesAttachedToANetworkProfileV1")
	}

	result := response.Result().(*ResponseNetworkSettingsRetrieveCountOfCliTemplatesAttachedToANetworkProfileV1)
	return result, response, err

}

//GetReserveIPSubpoolV1 Get Reserve IP Subpool - 4586-0917-4fab-87e2
/* API to get the ip subpool info.


@param GetReserveIPSubpoolV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-reserve-ip-subpool
*/
func (s *NetworkSettingsService) GetReserveIPSubpoolV1(GetReserveIPSubpoolV1QueryParams *GetReserveIPSubpoolV1QueryParams) (*ResponseNetworkSettingsGetReserveIPSubpoolV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/reserve-ip-subpool"

	queryString, _ := query.Values(GetReserveIPSubpoolV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworkSettingsGetReserveIPSubpoolV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetReserveIPSubpoolV1(GetReserveIPSubpoolV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetReserveIpSubpoolV1")
	}

	result := response.Result().(*ResponseNetworkSettingsGetReserveIPSubpoolV1)
	return result, response, err

}

//GetServiceProviderDetailsV1 Get Service provider Details - 7084-7bdc-4d89-a437
/* API to get service provider details (QoS).



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-service-provider-details
*/
func (s *NetworkSettingsService) GetServiceProviderDetailsV1() (*ResponseNetworkSettingsGetServiceProviderDetailsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/service-provider"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkSettingsGetServiceProviderDetailsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetServiceProviderDetailsV1()
		}
		return nil, response, fmt.Errorf("error with operation GetServiceProviderDetailsV1")
	}

	result := response.Result().(*ResponseNetworkSettingsGetServiceProviderDetailsV1)
	return result, response, err

}

//RetrieveAAASettingsForASiteV1 Retrieve AAA settings for a site - 3c99-79ea-4ab9-bd33
/* Retrieve AAA settings for a site; `null` values indicate that the setting will be inherited from the parent site; empty objects (`{}`) indicate that the setting is unset at a site.


@param id id path parameter. Site Id

@param RetrieveAAASettingsForASiteV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-a-a-a-settings-for-a-site
*/
func (s *NetworkSettingsService) RetrieveAAASettingsForASiteV1(id string, RetrieveAAASettingsForASiteV1QueryParams *RetrieveAAASettingsForASiteV1QueryParams) (*ResponseNetworkSettingsRetrieveAAASettingsForASiteV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/aaaSettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(RetrieveAAASettingsForASiteV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworkSettingsRetrieveAAASettingsForASiteV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveAAASettingsForASiteV1(id, RetrieveAAASettingsForASiteV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveAAASettingsForASiteV1")
	}

	result := response.Result().(*ResponseNetworkSettingsRetrieveAAASettingsForASiteV1)
	return result, response, err

}

//RetrieveBannerSettingsForASiteV1 Retrieve banner settings for a site - 2a9f-3b2f-4cda-8390
/* Retrieve banner settings for a site; `null` values indicate that the setting will be inherited from the parent site; empty objects (`{}`) indicate that the setting is unset at a site.


@param id id path parameter. Site Id

@param RetrieveBannerSettingsForASiteV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-banner-settings-for-a-site
*/
func (s *NetworkSettingsService) RetrieveBannerSettingsForASiteV1(id string, RetrieveBannerSettingsForASiteV1QueryParams *RetrieveBannerSettingsForASiteV1QueryParams) (*ResponseNetworkSettingsRetrieveBannerSettingsForASiteV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/bannerSettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(RetrieveBannerSettingsForASiteV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworkSettingsRetrieveBannerSettingsForASiteV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveBannerSettingsForASiteV1(id, RetrieveBannerSettingsForASiteV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveBannerSettingsForASiteV1")
	}

	result := response.Result().(*ResponseNetworkSettingsRetrieveBannerSettingsForASiteV1)
	return result, response, err

}

//GetDeviceCredentialSettingsForASiteV1 Get device credential settings for a site - bebf-c9fc-4d3a-be03
/* Gets device credential settings for a site; `null` values indicate that the setting will be inherited from the parent site; empty objects (`{}`) indicate that the credential is unset, and that no credential of that type will be used for the site.


@param id id path parameter. Site Id, retrievable from the `id` attribute in `/dna/intent/api/v1/sites`

@param GetDeviceCredentialSettingsForASiteV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-credential-settings-for-a-site
*/
func (s *NetworkSettingsService) GetDeviceCredentialSettingsForASiteV1(id string, GetDeviceCredentialSettingsForASiteV1QueryParams *GetDeviceCredentialSettingsForASiteV1QueryParams) (*ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/deviceCredentials"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetDeviceCredentialSettingsForASiteV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceCredentialSettingsForASiteV1(id, GetDeviceCredentialSettingsForASiteV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceCredentialSettingsForASiteV1")
	}

	result := response.Result().(*ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteV1)
	return result, response, err

}

//GetNetworkDevicesCredentialsSyncStatusV1 Get network devices credentials sync status - 0f9e-d8a6-45eb-9590
/* Get network devices credentials sync status at a given site.


@param id id path parameter. Site Id.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-devices-credentials-sync-status
*/
func (s *NetworkSettingsService) GetNetworkDevicesCredentialsSyncStatusV1(id string) (*ResponseNetworkSettingsGetNetworkDevicesCredentialsSyncStatusV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/deviceCredentials/status"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkSettingsGetNetworkDevicesCredentialsSyncStatusV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetNetworkDevicesCredentialsSyncStatusV1(id)
		}
		return nil, response, fmt.Errorf("error with operation GetNetworkDevicesCredentialsSyncStatusV1")
	}

	result := response.Result().(*ResponseNetworkSettingsGetNetworkDevicesCredentialsSyncStatusV1)
	return result, response, err

}

//RetrieveDHCPSettingsForASiteV1 Retrieve DHCP settings for a site - cfbb-ca8d-4529-a94b
/* Retrieve DHCP settings for a site; `null` values indicate that the setting will be inherited from the parent site; empty objects (`{}`) indicate that the setting is unset at a site.


@param id id path parameter. Site Id

@param RetrieveDHCPSettingsForASiteV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-d-h-c-p-settings-for-a-site
*/
func (s *NetworkSettingsService) RetrieveDHCPSettingsForASiteV1(id string, RetrieveDHCPSettingsForASiteV1QueryParams *RetrieveDHCPSettingsForASiteV1QueryParams) (*ResponseNetworkSettingsRetrieveDHCPSettingsForASiteV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/dhcpSettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(RetrieveDHCPSettingsForASiteV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworkSettingsRetrieveDHCPSettingsForASiteV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveDHCPSettingsForASiteV1(id, RetrieveDHCPSettingsForASiteV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveDHCPSettingsForASiteV1")
	}

	result := response.Result().(*ResponseNetworkSettingsRetrieveDHCPSettingsForASiteV1)
	return result, response, err

}

//RetrieveDNSSettingsForASiteV1 Retrieve DNS settings for a site - d7a4-0932-41d9-bcf8
/* Retrieve DNS settings for a site; `null` values indicate that the setting will be inherited from the parent site; empty objects (`{}`) indicate that the setting is unset at a site.


@param id id path parameter. Site Id

@param RetrieveDNSSettingsForASiteV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-d-n-s-settings-for-a-site
*/
func (s *NetworkSettingsService) RetrieveDNSSettingsForASiteV1(id string, RetrieveDNSSettingsForASiteV1QueryParams *RetrieveDNSSettingsForASiteV1QueryParams) (*ResponseNetworkSettingsRetrieveDNSSettingsForASiteV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/dnsSettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(RetrieveDNSSettingsForASiteV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworkSettingsRetrieveDNSSettingsForASiteV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveDNSSettingsForASiteV1(id, RetrieveDNSSettingsForASiteV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveDNSSettingsForASiteV1")
	}

	result := response.Result().(*ResponseNetworkSettingsRetrieveDNSSettingsForASiteV1)
	return result, response, err

}

//RetrieveImageDistributionSettingsForASiteV1 Retrieve image distribution settings for a site - d2ad-d9bc-4bcb-9fed
/* Retrieve image distribution settings for a site; `null` values indicate that the setting will be inherited from the parent site; empty objects (`{}`) indicate that the setting is unset at a site.


@param id id path parameter. Site Id

@param RetrieveImageDistributionSettingsForASiteV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-image-distribution-settings-for-a-site
*/
func (s *NetworkSettingsService) RetrieveImageDistributionSettingsForASiteV1(id string, RetrieveImageDistributionSettingsForASiteV1QueryParams *RetrieveImageDistributionSettingsForASiteV1QueryParams) (*ResponseNetworkSettingsRetrieveImageDistributionSettingsForASiteV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/imageDistributionSettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(RetrieveImageDistributionSettingsForASiteV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworkSettingsRetrieveImageDistributionSettingsForASiteV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveImageDistributionSettingsForASiteV1(id, RetrieveImageDistributionSettingsForASiteV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveImageDistributionSettingsForASiteV1")
	}

	result := response.Result().(*ResponseNetworkSettingsRetrieveImageDistributionSettingsForASiteV1)
	return result, response, err

}

//RetrieveNTPSettingsForASiteV1 Retrieve NTP settings for a site - beae-2bf1-4cdb-8f60
/* Retrieve NTP settings for a site; `null` values indicate that the setting will be inherited from the parent site; empty objects (`{}`) indicate that the setting is unset at a site.


@param id id path parameter. Site Id

@param RetrieveNTPSettingsForASiteV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-n-t-p-settings-for-a-site
*/
func (s *NetworkSettingsService) RetrieveNTPSettingsForASiteV1(id string, RetrieveNTPSettingsForASiteV1QueryParams *RetrieveNTPSettingsForASiteV1QueryParams) (*ResponseNetworkSettingsRetrieveNTPSettingsForASiteV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/ntpSettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(RetrieveNTPSettingsForASiteV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworkSettingsRetrieveNTPSettingsForASiteV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveNTPSettingsForASiteV1(id, RetrieveNTPSettingsForASiteV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveNTPSettingsForASiteV1")
	}

	result := response.Result().(*ResponseNetworkSettingsRetrieveNTPSettingsForASiteV1)
	return result, response, err

}

//RetrieveTelemetrySettingsForASiteV1 Retrieve Telemetry settings for a site - 11a7-cbc7-4a9a-bac3
/* Retrieves telemetry settings for the given site. `null` values indicate that the setting will be inherited from the parent site.


@param id id path parameter. Site Id, retrievable from the `id` attribute in `/dna/intent/api/v1/sites`

@param RetrieveTelemetrySettingsForASiteV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-telemetry-settings-for-a-site
*/
func (s *NetworkSettingsService) RetrieveTelemetrySettingsForASiteV1(id string, RetrieveTelemetrySettingsForASiteV1QueryParams *RetrieveTelemetrySettingsForASiteV1QueryParams) (*ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/telemetrySettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(RetrieveTelemetrySettingsForASiteV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveTelemetrySettingsForASiteV1(id, RetrieveTelemetrySettingsForASiteV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveTelemetrySettingsForASiteV1")
	}

	result := response.Result().(*ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteV1)
	return result, response, err

}

//RetrieveTimeZoneSettingsForASiteV1 Retrieve time zone settings for a site - 5ba6-0966-4768-9ae7
/* Retrieve time zone settings for a site; `null` values indicate that the setting will be inherited from the parent site; empty objects (`{}`) indicate that the setting is unset at a site.


@param id id path parameter. Site Id

@param RetrieveTimeZoneSettingsForASiteV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-time-zone-settings-for-a-site
*/
func (s *NetworkSettingsService) RetrieveTimeZoneSettingsForASiteV1(id string, RetrieveTimeZoneSettingsForASiteV1QueryParams *RetrieveTimeZoneSettingsForASiteV1QueryParams) (*ResponseNetworkSettingsRetrieveTimeZoneSettingsForASiteV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/timeZoneSettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(RetrieveTimeZoneSettingsForASiteV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworkSettingsRetrieveTimeZoneSettingsForASiteV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveTimeZoneSettingsForASiteV1(id, RetrieveTimeZoneSettingsForASiteV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveTimeZoneSettingsForASiteV1")
	}

	result := response.Result().(*ResponseNetworkSettingsRetrieveTimeZoneSettingsForASiteV1)
	return result, response, err

}

//GetNetworkV2 Get Network V2 - fdbd-3b7b-4048-bbfd
/* API to get SNMP, NTP, Network AAA, Client and Endpoint AAA, and/or DNS center server settings.


@param GetNetworkV2QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-v2
*/
func (s *NetworkSettingsService) GetNetworkV2(GetNetworkV2QueryParams *GetNetworkV2QueryParams) (*ResponseNetworkSettingsGetNetworkV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/network"

	queryString, _ := query.Values(GetNetworkV2QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworkSettingsGetNetworkV2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetNetworkV2(GetNetworkV2QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetNetworkV2")
	}

	result := response.Result().(*ResponseNetworkSettingsGetNetworkV2)
	return result, response, err

}

//GetServiceProviderDetailsV2 Get Service Provider Details V2 - c89e-0ac4-4279-8591
/* API to get Service Provider details (QoS).



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-service-provider-details-v2
*/
func (s *NetworkSettingsService) GetServiceProviderDetailsV2() (*ResponseNetworkSettingsGetServiceProviderDetailsV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/service-provider"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkSettingsGetServiceProviderDetailsV2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetServiceProviderDetailsV2()
		}
		return nil, response, fmt.Errorf("error with operation GetServiceProviderDetailsV2")
	}

	result := response.Result().(*ResponseNetworkSettingsGetServiceProviderDetailsV2)
	return result, response, err

}

//AssignDeviceCredentialToSiteV1 Assign Device Credential To Site - 4da9-1a54-4e29-842d
/* Assign Device Credential to a site.


@param siteID siteId path parameter. site id to assign credential.

@param AssignDeviceCredentialToSiteV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!assign-device-credential-to-site
*/
func (s *NetworkSettingsService) AssignDeviceCredentialToSiteV1(siteID string, requestNetworkSettingsAssignDeviceCredentialToSiteV1 *RequestNetworkSettingsAssignDeviceCredentialToSiteV1, AssignDeviceCredentialToSiteV1HeaderParams *AssignDeviceCredentialToSiteV1HeaderParams) (*ResponseNetworkSettingsAssignDeviceCredentialToSiteV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/credential-to-site/{siteId}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if AssignDeviceCredentialToSiteV1HeaderParams != nil {

		if AssignDeviceCredentialToSiteV1HeaderParams.Persistbapioutput != "" {
			clientRequest = clientRequest.SetHeader("__persistbapioutput", AssignDeviceCredentialToSiteV1HeaderParams.Persistbapioutput)
		}

	}

	response, err = clientRequest.
		SetBody(requestNetworkSettingsAssignDeviceCredentialToSiteV1).
		SetResult(&ResponseNetworkSettingsAssignDeviceCredentialToSiteV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AssignDeviceCredentialToSiteV1(siteID, requestNetworkSettingsAssignDeviceCredentialToSiteV1, AssignDeviceCredentialToSiteV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation AssignDeviceCredentialToSiteV1")
	}

	result := response.Result().(*ResponseNetworkSettingsAssignDeviceCredentialToSiteV1)
	return result, response, err

}

//CreateDeviceCredentialsV1 Create Device Credentials - fbb9-5b37-484a-9fce
/* API to create device credentials.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-device-credentials
*/
func (s *NetworkSettingsService) CreateDeviceCredentialsV1(requestNetworkSettingsCreateDeviceCredentialsV1 *RequestNetworkSettingsCreateDeviceCredentialsV1) (*ResponseNetworkSettingsCreateDeviceCredentialsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/device-credential"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsCreateDeviceCredentialsV1).
		SetResult(&ResponseNetworkSettingsCreateDeviceCredentialsV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateDeviceCredentialsV1(requestNetworkSettingsCreateDeviceCredentialsV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateDeviceCredentialsV1")
	}

	result := response.Result().(*ResponseNetworkSettingsCreateDeviceCredentialsV1)
	return result, response, err

}

//CreateGlobalPoolV1 Create Global Pool - f793-192a-43da-bed9
/* API to create global pool. There is a limit of creating 25 global pools per request.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-global-pool
*/
func (s *NetworkSettingsService) CreateGlobalPoolV1(requestNetworkSettingsCreateGlobalPoolV1 *RequestNetworkSettingsCreateGlobalPoolV1) (*ResponseNetworkSettingsCreateGlobalPoolV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-pool"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsCreateGlobalPoolV1).
		SetResult(&ResponseNetworkSettingsCreateGlobalPoolV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateGlobalPoolV1(requestNetworkSettingsCreateGlobalPoolV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateGlobalPoolV1")
	}

	result := response.Result().(*ResponseNetworkSettingsCreateGlobalPoolV1)
	return result, response, err

}

//CreateAGlobalIPAddressPoolV1 Create a global IP address pool. - abb0-f98d-4d69-b5ab
/* Creates a global IP address pool, which is not bound to a particular site. A global pool must be either an IPv4 or IPv6 pool.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-a-global-ip-address-pool
*/
func (s *NetworkSettingsService) CreateAGlobalIPAddressPoolV1(requestNetworkSettingsCreateAGlobalIPAddressPoolV1 *RequestNetworkSettingsCreateAGlobalIPAddressPoolV1) (*ResponseNetworkSettingsCreateAGlobalIPAddressPoolV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/ipam/globalIpAddressPools"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsCreateAGlobalIPAddressPoolV1).
		SetResult(&ResponseNetworkSettingsCreateAGlobalIPAddressPoolV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateAGlobalIPAddressPoolV1(requestNetworkSettingsCreateAGlobalIPAddressPoolV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateAGlobalIpAddressPoolV1")
	}

	result := response.Result().(*ResponseNetworkSettingsCreateAGlobalIPAddressPoolV1)
	return result, response, err

}

//ReservecreateIPAddressSubpoolsV1 Reserve (create) IP address subpools. - ecb4-7a7e-43da-91fa
/* Reserves (creates) an IP address subpool, which reserves address space from a global pool (or global pools) for a particular site (and it's child sites). A subpool must be either an IPv4 or dual-stack pool, with `ipV4AddressSpace` and optionally `ipV6AddressSpace` properties specified.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!reservecreate-ip-address-subpools
*/
func (s *NetworkSettingsService) ReservecreateIPAddressSubpoolsV1(requestNetworkSettingsReservecreateIPAddressSubpoolsV1 *RequestNetworkSettingsReservecreateIPAddressSubpoolsV1) (*ResponseNetworkSettingsReservecreateIPAddressSubpoolsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/ipam/siteIpAddressPools"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsReservecreateIPAddressSubpoolsV1).
		SetResult(&ResponseNetworkSettingsReservecreateIPAddressSubpoolsV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReservecreateIPAddressSubpoolsV1(requestNetworkSettingsReservecreateIPAddressSubpoolsV1)
		}

		return nil, response, fmt.Errorf("error with operation ReservecreateIpAddressSubpoolsV1")
	}

	result := response.Result().(*ResponseNetworkSettingsReservecreateIPAddressSubpoolsV1)
	return result, response, err

}

//CreateNetworkV1 Create Network - be89-2bd8-4a78-865a
/* API to create a network for DHCP,  Syslog, SNMP, NTP, Network AAA, Client and EndPoint AAA, and/or DNS center server settings.


@param siteID siteId path parameter. Site id to which site details to associate with the network settings.

@param CreateNetworkV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network
*/
func (s *NetworkSettingsService) CreateNetworkV1(siteID string, requestNetworkSettingsCreateNetworkV1 *RequestNetworkSettingsCreateNetworkV1, CreateNetworkV1HeaderParams *CreateNetworkV1HeaderParams) (*ResponseNetworkSettingsCreateNetworkV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network/{siteId}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if CreateNetworkV1HeaderParams != nil {

		if CreateNetworkV1HeaderParams.Persistbapioutput != "" {
			clientRequest = clientRequest.SetHeader("__persistbapioutput", CreateNetworkV1HeaderParams.Persistbapioutput)
		}

	}

	response, err = clientRequest.
		SetBody(requestNetworkSettingsCreateNetworkV1).
		SetResult(&ResponseNetworkSettingsCreateNetworkV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateNetworkV1(siteID, requestNetworkSettingsCreateNetworkV1, CreateNetworkV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation CreateNetworkV1")
	}

	result := response.Result().(*ResponseNetworkSettingsCreateNetworkV1)
	return result, response, err

}

//ReserveIPSubpoolV1 Reserve IP Subpool - 429f-aa81-4d3b-960a
/* API to reserve an ip subpool from the global pool


@param siteID siteId path parameter. Site id to reserve the ip sub pool.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!reserve-ip-subpool
*/
func (s *NetworkSettingsService) ReserveIPSubpoolV1(siteID string, requestNetworkSettingsReserveIPSubpoolV1 *RequestNetworkSettingsReserveIPSubpoolV1) (*ResponseNetworkSettingsReserveIPSubpoolV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/reserve-ip-subpool/{siteId}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsReserveIPSubpoolV1).
		SetResult(&ResponseNetworkSettingsReserveIPSubpoolV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReserveIPSubpoolV1(siteID, requestNetworkSettingsReserveIPSubpoolV1)
		}

		return nil, response, fmt.Errorf("error with operation ReserveIpSubpoolV1")
	}

	result := response.Result().(*ResponseNetworkSettingsReserveIPSubpoolV1)
	return result, response, err

}

//CreateSpProfileV1 Create SP Profile - a39a-1a21-4deb-b781
/* API to create Service Provider Profile(QOS).



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-sp-profile
*/
func (s *NetworkSettingsService) CreateSpProfileV1(requestNetworkSettingsCreateSPProfileV1 *RequestNetworkSettingsCreateSpProfileV1) (*ResponseNetworkSettingsCreateSpProfileV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/service-provider"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsCreateSPProfileV1).
		SetResult(&ResponseNetworkSettingsCreateSpProfileV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateSpProfileV1(requestNetworkSettingsCreateSPProfileV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateSpProfileV1")
	}

	result := response.Result().(*ResponseNetworkSettingsCreateSpProfileV1)
	return result, response, err

}

//SyncNetworkDevicesCredentialV1 Sync network devices credential - 5ca0-2a36-4d68-92e7
/* When sync is triggered at a site with the credential that are associated to the same site, network devices in impacted sites (child sites which are inheriting the credential) get managed in inventory with the associated site credential. Credential gets configured on network devices before these get managed in inventory. Please make a note that cli credential wouldn't be configured on AAA authenticated devices but they just get managed with the associated site cli credential.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!sync-network-devices-credential
*/
func (s *NetworkSettingsService) SyncNetworkDevicesCredentialV1(requestNetworkSettingsSyncNetworkDevicesCredentialV1 *RequestNetworkSettingsSyncNetworkDevicesCredentialV1) (*ResponseNetworkSettingsSyncNetworkDevicesCredentialV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/deviceCredentials/apply"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsSyncNetworkDevicesCredentialV1).
		SetResult(&ResponseNetworkSettingsSyncNetworkDevicesCredentialV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.SyncNetworkDevicesCredentialV1(requestNetworkSettingsSyncNetworkDevicesCredentialV1)
		}

		return nil, response, fmt.Errorf("error with operation SyncNetworkDevicesCredentialV1")
	}

	result := response.Result().(*ResponseNetworkSettingsSyncNetworkDevicesCredentialV1)
	return result, response, err

}

//UpdateADevicesTelemetrySettingsToConformToTheTelemetrySettingsForItsSiteV1 Update a device(s) telemetry settings to conform to the telemetry settings for its site - 14bf-7997-432b-94a1
/* Update a device(s) telemetry settings to conform to the telemetry settings for its site.  One Task is created to track the update, for more granular status tracking, split your devices into multiple requests.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!update-a-devices-telemetry-settings-to-conform-to-the-telemetry-settings-for-its-site
*/
func (s *NetworkSettingsService) UpdateADevicesTelemetrySettingsToConformToTheTelemetrySettingsForItsSiteV1(requestNetworkSettingsUpdateADevicesTelemetrySettingsToConformToTheTelemetrySettingsForItsSiteV1 *RequestNetworkSettingsUpdateADevicesTelemetrySettingsToConformToTheTelemetrySettingsForItsSiteV1) (*ResponseNetworkSettingsUpdateADevicesTelemetrySettingsToConformToTheTelemetrySettingsForItsSiteV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/telemetrySettings/apply"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsUpdateADevicesTelemetrySettingsToConformToTheTelemetrySettingsForItsSiteV1).
		SetResult(&ResponseNetworkSettingsUpdateADevicesTelemetrySettingsToConformToTheTelemetrySettingsForItsSiteV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateADevicesTelemetrySettingsToConformToTheTelemetrySettingsForItsSiteV1(requestNetworkSettingsUpdateADevicesTelemetrySettingsToConformToTheTelemetrySettingsForItsSiteV1)
		}

		return nil, response, fmt.Errorf("error with operation UpdateADevicesTelemetrySettingsToConformToTheTelemetrySettingsForItsSiteV1")
	}

	result := response.Result().(*ResponseNetworkSettingsUpdateADevicesTelemetrySettingsToConformToTheTelemetrySettingsForItsSiteV1)
	return result, response, err

}

//AssignDeviceCredentialToSiteV2 Assign Device Credential To Site V2 - 0eb2-8bc5-4b99-8d6c
/* API to assign Device Credential to a site.


@param siteID siteId path parameter. Site Id to assign credential.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!assign-device-credential-to-site-v2
*/
func (s *NetworkSettingsService) AssignDeviceCredentialToSiteV2(siteID string, requestNetworkSettingsAssignDeviceCredentialToSiteV2 *RequestNetworkSettingsAssignDeviceCredentialToSiteV2) (*ResponseNetworkSettingsAssignDeviceCredentialToSiteV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/credential-to-site/{siteId}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsAssignDeviceCredentialToSiteV2).
		SetResult(&ResponseNetworkSettingsAssignDeviceCredentialToSiteV2{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AssignDeviceCredentialToSiteV2(siteID, requestNetworkSettingsAssignDeviceCredentialToSiteV2)
		}

		return nil, response, fmt.Errorf("error with operation AssignDeviceCredentialToSiteV2")
	}

	result := response.Result().(*ResponseNetworkSettingsAssignDeviceCredentialToSiteV2)
	return result, response, err

}

//CreateNetworkV2 Create Network V2 - 4696-fb19-48da-ac38
/* API to create network settings for DHCP,  Syslog, SNMP, NTP, Network AAA, Client and Endpoint AAA, and/or DNS center server settings.


@param siteID siteId path parameter. Site Id to which site details to associate with the network settings.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network-v2
*/
func (s *NetworkSettingsService) CreateNetworkV2(siteID string, requestNetworkSettingsCreateNetworkV2 *RequestNetworkSettingsCreateNetworkV2) (*ResponseNetworkSettingsCreateNetworkV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/network/{siteId}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsCreateNetworkV2).
		SetResult(&ResponseNetworkSettingsCreateNetworkV2{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateNetworkV2(siteID, requestNetworkSettingsCreateNetworkV2)
		}

		return nil, response, fmt.Errorf("error with operation CreateNetworkV2")
	}

	result := response.Result().(*ResponseNetworkSettingsCreateNetworkV2)
	return result, response, err

}

//CreateSpProfileV2 Create SP Profile V2 - b6b1-1aa9-4b4b-99e0
/* API to create Service Provider Profile(QOS).



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-sp-profile-v2
*/
func (s *NetworkSettingsService) CreateSpProfileV2(requestNetworkSettingsCreateSPProfileV2 *RequestNetworkSettingsCreateSpProfileV2) (*ResponseNetworkSettingsCreateSpProfileV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/service-provider"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsCreateSPProfileV2).
		SetResult(&ResponseNetworkSettingsCreateSpProfileV2{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateSpProfileV2(requestNetworkSettingsCreateSPProfileV2)
		}

		return nil, response, fmt.Errorf("error with operation CreateSpProfileV2")
	}

	result := response.Result().(*ResponseNetworkSettingsCreateSpProfileV2)
	return result, response, err

}

//UpdateDeviceCredentialsV1 Update Device Credentials - 4f94-7a1c-4fc8-84f6
/* API to update device credentials.


 */
func (s *NetworkSettingsService) UpdateDeviceCredentialsV1(requestNetworkSettingsUpdateDeviceCredentialsV1 *RequestNetworkSettingsUpdateDeviceCredentialsV1) (*ResponseNetworkSettingsUpdateDeviceCredentialsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/device-credential"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsUpdateDeviceCredentialsV1).
		SetResult(&ResponseNetworkSettingsUpdateDeviceCredentialsV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateDeviceCredentialsV1(requestNetworkSettingsUpdateDeviceCredentialsV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateDeviceCredentialsV1")
	}

	result := response.Result().(*ResponseNetworkSettingsUpdateDeviceCredentialsV1)
	return result, response, err

}

//UpdateGlobalPoolV1 Update Global Pool - 03b4-c8b4-4919-b964
/* API to update global pool. There is a limit of updating 25 global pools per request.


 */
func (s *NetworkSettingsService) UpdateGlobalPoolV1(requestNetworkSettingsUpdateGlobalPoolV1 *RequestNetworkSettingsUpdateGlobalPoolV1) (*ResponseNetworkSettingsUpdateGlobalPoolV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-pool"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsUpdateGlobalPoolV1).
		SetResult(&ResponseNetworkSettingsUpdateGlobalPoolV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateGlobalPoolV1(requestNetworkSettingsUpdateGlobalPoolV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateGlobalPoolV1")
	}

	result := response.Result().(*ResponseNetworkSettingsUpdateGlobalPoolV1)
	return result, response, err

}

//UpdatesAGlobalIPAddressPoolV1 Updates a global IP address pool. - 68a9-5861-467a-abd9
/* Updates a global IP address pool.
Restrictions on updating a global IP address pool: The `poolType` cannot be changed. The `subnet` and `prefixLength` within `addressSpace` cannot be changed.


@param id id path parameter. The `id` of the global IP address pool to update.

*/
func (s *NetworkSettingsService) UpdatesAGlobalIPAddressPoolV1(id string, requestNetworkSettingsUpdatesAGlobalIPAddressPoolV1 *RequestNetworkSettingsUpdatesAGlobalIPAddressPoolV1) (*ResponseNetworkSettingsUpdatesAGlobalIPAddressPoolV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/ipam/globalIpAddressPools/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsUpdatesAGlobalIPAddressPoolV1).
		SetResult(&ResponseNetworkSettingsUpdatesAGlobalIPAddressPoolV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdatesAGlobalIPAddressPoolV1(id, requestNetworkSettingsUpdatesAGlobalIPAddressPoolV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdatesAGlobalIpAddressPoolV1")
	}

	result := response.Result().(*ResponseNetworkSettingsUpdatesAGlobalIPAddressPoolV1)
	return result, response, err

}

//UpdatesAnIPAddressSubpoolV1 Updates an IP address subpool. - 968f-d8e4-465b-8ec2
/* Updates an IP address subpool, which reserves address space from a global pool (or global pools) for a particular site.
Restrictions on updating an IP address subpool: The `poolType` cannot be changed. The `siteId` cannot be changed. The `ipV4AddressSpace` may not be removed. The `globalPoolId`, `subnet`, and `prefixLength` cannot be changed once it's already been set. However you may edit a subpool to add an IP address space if it does not already have one.


@param id id path parameter. The `id` of the IP address subpool to update.

*/
func (s *NetworkSettingsService) UpdatesAnIPAddressSubpoolV1(id string, requestNetworkSettingsUpdatesAnIPAddressSubpoolV1 *RequestNetworkSettingsUpdatesAnIPAddressSubpoolV1) (*ResponseNetworkSettingsUpdatesAnIPAddressSubpoolV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/ipam/siteIpAddressPools/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsUpdatesAnIPAddressSubpoolV1).
		SetResult(&ResponseNetworkSettingsUpdatesAnIPAddressSubpoolV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdatesAnIPAddressSubpoolV1(id, requestNetworkSettingsUpdatesAnIPAddressSubpoolV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdatesAnIpAddressSubpoolV1")
	}

	result := response.Result().(*ResponseNetworkSettingsUpdatesAnIPAddressSubpoolV1)
	return result, response, err

}

//UpdateNetworkV1 Update Network - 698b-fbb4-4dcb-9fca
/* API to update network settings for DHCP,  Syslog, SNMP, NTP, Network AAA, Client and EndPoint AAA, and/or DNS server settings.


@param siteID siteId path parameter. Site id to update the network settings which is associated with the site

*/
func (s *NetworkSettingsService) UpdateNetworkV1(siteID string, requestNetworkSettingsUpdateNetworkV1 *RequestNetworkSettingsUpdateNetworkV1) (*ResponseNetworkSettingsUpdateNetworkV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/network/{siteId}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsUpdateNetworkV1).
		SetResult(&ResponseNetworkSettingsUpdateNetworkV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateNetworkV1(siteID, requestNetworkSettingsUpdateNetworkV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateNetworkV1")
	}

	result := response.Result().(*ResponseNetworkSettingsUpdateNetworkV1)
	return result, response, err

}

//UpdateReserveIPSubpoolV1 Update Reserve IP Subpool - 6992-d8ec-42cb-88f1
/* API to update ip subpool from the global pool


@param siteID siteId path parameter. Site id of site to update sub pool.

@param UpdateReserveIPSubpoolV1QueryParams Filtering parameter
*/
func (s *NetworkSettingsService) UpdateReserveIPSubpoolV1(siteID string, requestNetworkSettingsUpdateReserveIPSubpoolV1 *RequestNetworkSettingsUpdateReserveIPSubpoolV1, UpdateReserveIPSubpoolV1QueryParams *UpdateReserveIPSubpoolV1QueryParams) (*ResponseNetworkSettingsUpdateReserveIPSubpoolV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/reserve-ip-subpool/{siteId}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	queryString, _ := query.Values(UpdateReserveIPSubpoolV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetBody(requestNetworkSettingsUpdateReserveIPSubpoolV1).
		SetResult(&ResponseNetworkSettingsUpdateReserveIPSubpoolV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateReserveIPSubpoolV1(siteID, requestNetworkSettingsUpdateReserveIPSubpoolV1, UpdateReserveIPSubpoolV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation UpdateReserveIpSubpoolV1")
	}

	result := response.Result().(*ResponseNetworkSettingsUpdateReserveIPSubpoolV1)
	return result, response, err

}

//UpdateSpProfileV1 Update SP Profile - 5087-daae-4cc9-8566
/* API to update Service Provider Profile (QoS).


 */
func (s *NetworkSettingsService) UpdateSpProfileV1(requestNetworkSettingsUpdateSPProfileV1 *RequestNetworkSettingsUpdateSpProfileV1) (*ResponseNetworkSettingsUpdateSpProfileV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/service-provider"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsUpdateSPProfileV1).
		SetResult(&ResponseNetworkSettingsUpdateSpProfileV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateSpProfileV1(requestNetworkSettingsUpdateSPProfileV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateSpProfileV1")
	}

	result := response.Result().(*ResponseNetworkSettingsUpdateSpProfileV1)
	return result, response, err

}

//SetAAASettingsForASiteV1 Set AAA settings for a site - 3582-ca30-4718-a064
/* Set AAA settings for a site; `null` values indicate that the settings will be inherited from the parent site; empty objects (`{}`) indicate that the settings is unset.


@param id id path parameter. Site Id

*/
func (s *NetworkSettingsService) SetAAASettingsForASiteV1(id string, requestNetworkSettingsSetAAASettingsForASiteV1 *RequestNetworkSettingsSetAAASettingsForASiteV1) (*ResponseNetworkSettingsSetAAASettingsForASiteV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/aaaSettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsSetAAASettingsForASiteV1).
		SetResult(&ResponseNetworkSettingsSetAAASettingsForASiteV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SetAAASettingsForASiteV1(id, requestNetworkSettingsSetAAASettingsForASiteV1)
		}
		return nil, response, fmt.Errorf("error with operation SetAAASettingsForASiteV1")
	}

	result := response.Result().(*ResponseNetworkSettingsSetAAASettingsForASiteV1)
	return result, response, err

}

//SetBannerSettingsForASiteV1 Set banner settings for a site - 0aae-aa56-44f9-95a7
/* Set banner settings for a site; `null` values indicate that the setting will be inherited from the parent site; empty objects (`{}`) indicate that the settings is unset.


@param id id path parameter. Site Id

*/
func (s *NetworkSettingsService) SetBannerSettingsForASiteV1(id string, requestNetworkSettingsSetBannerSettingsForASiteV1 *RequestNetworkSettingsSetBannerSettingsForASiteV1) (*ResponseNetworkSettingsSetBannerSettingsForASiteV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/bannerSettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsSetBannerSettingsForASiteV1).
		SetResult(&ResponseNetworkSettingsSetBannerSettingsForASiteV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SetBannerSettingsForASiteV1(id, requestNetworkSettingsSetBannerSettingsForASiteV1)
		}
		return nil, response, fmt.Errorf("error with operation SetBannerSettingsForASiteV1")
	}

	result := response.Result().(*ResponseNetworkSettingsSetBannerSettingsForASiteV1)
	return result, response, err

}

//UpdateDeviceCredentialSettingsForASiteV1 Update device credential settings for a site. - 5aa0-6949-4dfa-bec5
/* Updates device credential settings for a site; `null` values indicate that the setting will be inherited from the parent site; empty objects (`{}`) indicate that the credential is unset, and that no credential of that type will be used for the site.


@param id id path parameter. Site Id, retrievable from the `id` attribute in `/dna/intent/api/v1/sites`

*/
func (s *NetworkSettingsService) UpdateDeviceCredentialSettingsForASiteV1(id string, requestNetworkSettingsUpdateDeviceCredentialSettingsForASiteV1 *RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteV1) (*ResponseNetworkSettingsUpdateDeviceCredentialSettingsForASiteV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/deviceCredentials"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsUpdateDeviceCredentialSettingsForASiteV1).
		SetResult(&ResponseNetworkSettingsUpdateDeviceCredentialSettingsForASiteV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateDeviceCredentialSettingsForASiteV1(id, requestNetworkSettingsUpdateDeviceCredentialSettingsForASiteV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateDeviceCredentialSettingsForASiteV1")
	}

	result := response.Result().(*ResponseNetworkSettingsUpdateDeviceCredentialSettingsForASiteV1)
	return result, response, err

}

//SetDhcpSettingsForASiteV1 Set dhcp settings for a site - c1ac-194d-40d9-8ae4
/* Set DHCP settings for a site; `null` values indicate that the setting will be inherited from the parent site; empty objects (`{}`) indicate that the settings is unset.


@param id id path parameter. Site Id

*/
func (s *NetworkSettingsService) SetDhcpSettingsForASiteV1(id string, requestNetworkSettingsSetDhcpSettingsForASiteV1 *RequestNetworkSettingsSetDhcpSettingsForASiteV1) (*ResponseNetworkSettingsSetDhcpSettingsForASiteV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/dhcpSettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsSetDhcpSettingsForASiteV1).
		SetResult(&ResponseNetworkSettingsSetDhcpSettingsForASiteV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SetDhcpSettingsForASiteV1(id, requestNetworkSettingsSetDhcpSettingsForASiteV1)
		}
		return nil, response, fmt.Errorf("error with operation SetDhcpSettingsForASiteV1")
	}

	result := response.Result().(*ResponseNetworkSettingsSetDhcpSettingsForASiteV1)
	return result, response, err

}

//SetDNSSettingsForASiteV1 Set DNS settings for a site - 9892-798e-4ed8-a40b
/* Set DNS settings for a site; `null` values indicate that the setting will be inherited from the parent site; empty objects (`{}`) indicate that the settings is unset.


@param id id path parameter. Site Id

*/
func (s *NetworkSettingsService) SetDNSSettingsForASiteV1(id string, requestNetworkSettingsSetDNSSettingsForASiteV1 *RequestNetworkSettingsSetDNSSettingsForASiteV1) (*ResponseNetworkSettingsSetDNSSettingsForASiteV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/dnsSettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsSetDNSSettingsForASiteV1).
		SetResult(&ResponseNetworkSettingsSetDNSSettingsForASiteV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SetDNSSettingsForASiteV1(id, requestNetworkSettingsSetDNSSettingsForASiteV1)
		}
		return nil, response, fmt.Errorf("error with operation SetDNSSettingsForASiteV1")
	}

	result := response.Result().(*ResponseNetworkSettingsSetDNSSettingsForASiteV1)
	return result, response, err

}

//SetImageDistributionSettingsForASiteV1 Set image distribution settings for a site - 8ab7-3889-45f8-9e5d
/* Set image distribution settings for a site; `null` values indicate that the setting will be inherited from the parent site; empty objects (`{}`) indicate that the settings is unset.


@param id id path parameter. Site Id

*/
func (s *NetworkSettingsService) SetImageDistributionSettingsForASiteV1(id string, requestNetworkSettingsSetImageDistributionSettingsForASiteV1 *RequestNetworkSettingsSetImageDistributionSettingsForASiteV1) (*ResponseNetworkSettingsSetImageDistributionSettingsForASiteV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/imageDistributionSettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsSetImageDistributionSettingsForASiteV1).
		SetResult(&ResponseNetworkSettingsSetImageDistributionSettingsForASiteV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SetImageDistributionSettingsForASiteV1(id, requestNetworkSettingsSetImageDistributionSettingsForASiteV1)
		}
		return nil, response, fmt.Errorf("error with operation SetImageDistributionSettingsForASiteV1")
	}

	result := response.Result().(*ResponseNetworkSettingsSetImageDistributionSettingsForASiteV1)
	return result, response, err

}

//SetNTPSettingsForASiteV1 Set NTP settings for a site - 9d80-8815-42a9-b006
/* Set NTP settings for a site; `null` values indicate that the setting will be inherited from the parent site; empty objects (`{}`) indicate that the settings is unset.


@param id id path parameter. Site Id

*/
func (s *NetworkSettingsService) SetNTPSettingsForASiteV1(id string, requestNetworkSettingsSetNTPSettingsForASiteV1 *RequestNetworkSettingsSetNTPSettingsForASiteV1) (*ResponseNetworkSettingsSetNTPSettingsForASiteV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/ntpSettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsSetNTPSettingsForASiteV1).
		SetResult(&ResponseNetworkSettingsSetNTPSettingsForASiteV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SetNTPSettingsForASiteV1(id, requestNetworkSettingsSetNTPSettingsForASiteV1)
		}
		return nil, response, fmt.Errorf("error with operation SetNTPSettingsForASiteV1")
	}

	result := response.Result().(*ResponseNetworkSettingsSetNTPSettingsForASiteV1)
	return result, response, err

}

//SetTelemetrySettingsForASiteV1 Set Telemetry settings for a site - a5a1-6835-40ab-8d2f
/* Sets telemetry settings for the given site; `null` values indicate that the setting will be inherited from the parent site.


@param id id path parameter. Site Id, retrievable from the `id` attribute in `/dna/intent/api/v1/sites`

*/
func (s *NetworkSettingsService) SetTelemetrySettingsForASiteV1(id string, requestNetworkSettingsSetTelemetrySettingsForASiteV1 *RequestNetworkSettingsSetTelemetrySettingsForASiteV1) (*ResponseNetworkSettingsSetTelemetrySettingsForASiteV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/telemetrySettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsSetTelemetrySettingsForASiteV1).
		SetResult(&ResponseNetworkSettingsSetTelemetrySettingsForASiteV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SetTelemetrySettingsForASiteV1(id, requestNetworkSettingsSetTelemetrySettingsForASiteV1)
		}
		return nil, response, fmt.Errorf("error with operation SetTelemetrySettingsForASiteV1")
	}

	result := response.Result().(*ResponseNetworkSettingsSetTelemetrySettingsForASiteV1)
	return result, response, err

}

//SetTimeZoneForASiteV1 Set time zone for a site - 16a7-b874-4b19-88d0
/* Set time zone settings for a site; `null` values indicate that the setting will be inherited from the parent site; empty objects (`{}`) indicate that the settings is unset.


@param id id path parameter. Site Id

*/
func (s *NetworkSettingsService) SetTimeZoneForASiteV1(id string, requestNetworkSettingsSetTimeZoneForASiteV1 *RequestNetworkSettingsSetTimeZoneForASiteV1) (*ResponseNetworkSettingsSetTimeZoneForASiteV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/timeZoneSettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsSetTimeZoneForASiteV1).
		SetResult(&ResponseNetworkSettingsSetTimeZoneForASiteV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SetTimeZoneForASiteV1(id, requestNetworkSettingsSetTimeZoneForASiteV1)
		}
		return nil, response, fmt.Errorf("error with operation SetTimeZoneForASiteV1")
	}

	result := response.Result().(*ResponseNetworkSettingsSetTimeZoneForASiteV1)
	return result, response, err

}

//UpdateNetworkV2 Update Network V2 - ac85-8bd9-4c78-a705
/* API to update network settings for DHCP, Syslog, SNMP, NTP, Network AAA, Client and Endpoint AAA, and/or DNS center server settings.


@param siteID siteId path parameter. Site Id to update the network settings which is associated with the site

*/
func (s *NetworkSettingsService) UpdateNetworkV2(siteID string, requestNetworkSettingsUpdateNetworkV2 *RequestNetworkSettingsUpdateNetworkV2) (*ResponseNetworkSettingsUpdateNetworkV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/network/{siteId}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsUpdateNetworkV2).
		SetResult(&ResponseNetworkSettingsUpdateNetworkV2{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateNetworkV2(siteID, requestNetworkSettingsUpdateNetworkV2)
		}
		return nil, response, fmt.Errorf("error with operation UpdateNetworkV2")
	}

	result := response.Result().(*ResponseNetworkSettingsUpdateNetworkV2)
	return result, response, err

}

//UpdateSpProfileV2 Update SP Profile V2 - 7fa6-78e5-455a-94a6
/* API to update Service Provider Profile (QoS).


 */
func (s *NetworkSettingsService) UpdateSpProfileV2(requestNetworkSettingsUpdateSPProfileV2 *RequestNetworkSettingsUpdateSpProfileV2) (*ResponseNetworkSettingsUpdateSpProfileV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/service-provider"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsUpdateSPProfileV2).
		SetResult(&ResponseNetworkSettingsUpdateSpProfileV2{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateSpProfileV2(requestNetworkSettingsUpdateSPProfileV2)
		}
		return nil, response, fmt.Errorf("error with operation UpdateSpProfileV2")
	}

	result := response.Result().(*ResponseNetworkSettingsUpdateSpProfileV2)
	return result, response, err

}

//DeleteDeviceCredentialV1 Delete Device Credential - 259e-ab30-4598-8958
/* Delete device credential.


@param id id path parameter. global credential id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-device-credential
*/
func (s *NetworkSettingsService) DeleteDeviceCredentialV1(id string) (*ResponseNetworkSettingsDeleteDeviceCredentialV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/device-credential/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkSettingsDeleteDeviceCredentialV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteDeviceCredentialV1(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteDeviceCredentialV1")
	}

	result := response.Result().(*ResponseNetworkSettingsDeleteDeviceCredentialV1)
	return result, response, err

}

//DeleteGlobalIPPoolV1 Delete Global IP Pool - 1eaa-8b21-48ab-81de
/* API to delete global IP pool.


@param id id path parameter. global pool id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-global-ip-pool
*/
func (s *NetworkSettingsService) DeleteGlobalIPPoolV1(id string) (*ResponseNetworkSettingsDeleteGlobalIPPoolV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/global-pool/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkSettingsDeleteGlobalIPPoolV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteGlobalIPPoolV1(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteGlobalIpPoolV1")
	}

	result := response.Result().(*ResponseNetworkSettingsDeleteGlobalIPPoolV1)
	return result, response, err

}

//DeleteAGlobalIPAddressPoolV1 Delete a global IP address pool. - f7ae-db6b-4fea-bfd8
/* Deletes a global IP address pool.  A global IP address pool can only be deleted if there are no subpools reserving address space from it.


@param id id path parameter. The `id` of the global IP address pool to delete.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-a-global-ip-address-pool
*/
func (s *NetworkSettingsService) DeleteAGlobalIPAddressPoolV1(id string) (*ResponseNetworkSettingsDeleteAGlobalIPAddressPoolV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/ipam/globalIpAddressPools/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkSettingsDeleteAGlobalIPAddressPoolV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteAGlobalIPAddressPoolV1(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteAGlobalIpAddressPoolV1")
	}

	result := response.Result().(*ResponseNetworkSettingsDeleteAGlobalIPAddressPoolV1)
	return result, response, err

}

//ReleaseAnIPAddressSubpoolV1 Release an IP address subpool. - e680-0b06-45f9-927d
/* Releases an IP address subpool.
**Release** completely removes the subpool from the site, and from all child sites, and frees the address use from the global pool(s).  Subpools cannot be released when assigned addresses in use.


@param id id path parameter. The `id` of the IP address subpool to delete.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!release-an-ip-address-subpool
*/
func (s *NetworkSettingsService) ReleaseAnIPAddressSubpoolV1(id string) (*ResponseNetworkSettingsReleaseAnIPAddressSubpoolV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/ipam/siteIpAddressPools/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkSettingsReleaseAnIPAddressSubpoolV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReleaseAnIPAddressSubpoolV1(id)
		}
		return nil, response, fmt.Errorf("error with operation ReleaseAnIpAddressSubpoolV1")
	}

	result := response.Result().(*ResponseNetworkSettingsReleaseAnIPAddressSubpoolV1)
	return result, response, err

}

//ReleaseReserveIPSubpoolV1 Release Reserve IP Subpool - 85b2-89e3-4489-9dc1
/* API to delete the reserved ip subpool


@param id id path parameter. Id of reserve ip subpool to be deleted.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!release-reserve-ip-subpool
*/
func (s *NetworkSettingsService) ReleaseReserveIPSubpoolV1(id string) (*ResponseNetworkSettingsReleaseReserveIPSubpoolV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/reserve-ip-subpool/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkSettingsReleaseReserveIPSubpoolV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReleaseReserveIPSubpoolV1(id)
		}
		return nil, response, fmt.Errorf("error with operation ReleaseReserveIpSubpoolV1")
	}

	result := response.Result().(*ResponseNetworkSettingsReleaseReserveIPSubpoolV1)
	return result, response, err

}

//DeleteSpProfileV1 Delete SP Profile - 4ca2-db11-43eb-b5d7
/* API to delete Service Provider Profile (QoS).


@param spProfileName spProfileName path parameter. sp profile name


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-sp-profile
*/
func (s *NetworkSettingsService) DeleteSpProfileV1(spProfileName string) (*ResponseNetworkSettingsDeleteSpProfileV1, *resty.Response, error) {
	//spProfileName string
	path := "/dna/intent/api/v1/sp-profile/{spProfileName}"
	path = strings.Replace(path, "{spProfileName}", fmt.Sprintf("%v", spProfileName), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkSettingsDeleteSpProfileV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteSpProfileV1(spProfileName)
		}
		return nil, response, fmt.Errorf("error with operation DeleteSpProfileV1")
	}

	result := response.Result().(*ResponseNetworkSettingsDeleteSpProfileV1)
	return result, response, err

}

//DeleteSpProfileV2 Delete SP Profile V2 - 5ea0-cb0e-4ed9-9bec
/* API to delete Service Provider Profile (QoS).


@param spProfileName spProfileName path parameter. SP profile name


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-sp-profile-v2
*/
func (s *NetworkSettingsService) DeleteSpProfileV2(spProfileName string) (*ResponseNetworkSettingsDeleteSpProfileV2, *resty.Response, error) {
	//spProfileName string
	path := "/dna/intent/api/v2/sp-profile/{spProfileName}"
	path = strings.Replace(path, "{spProfileName}", fmt.Sprintf("%v", spProfileName), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkSettingsDeleteSpProfileV2{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteSpProfileV2(spProfileName)
		}
		return nil, response, fmt.Errorf("error with operation DeleteSpProfileV2")
	}

	result := response.Result().(*ResponseNetworkSettingsDeleteSpProfileV2)
	return result, response, err

}

// Alias Function
/*
This method acts as an alias for the method `SetImageDistributionSettingsForASiteV1`
*/
func (s *NetworkSettingsService) SetImageDistributionSettingsForASite(id string, requestNetworkSettingsSetImageDistributionSettingsForASiteV1 *RequestNetworkSettingsSetImageDistributionSettingsForASiteV1) (*ResponseNetworkSettingsSetImageDistributionSettingsForASiteV1, *resty.Response, error) {
	return s.SetImageDistributionSettingsForASiteV1(id, requestNetworkSettingsSetImageDistributionSettingsForASiteV1)
}

// Alias Function
/*
This method acts as an alias for the method `CountsIPAddressSubpoolsV1`
*/
func (s *NetworkSettingsService) CountsIPAddressSubpools(CountsIPAddressSubpoolsV1QueryParams *CountsIPAddressSubpoolsV1QueryParams) (*ResponseNetworkSettingsCountsIPAddressSubpoolsV1, *resty.Response, error) {
	return s.CountsIPAddressSubpoolsV1(CountsIPAddressSubpoolsV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `DeleteSpProfileV1`
*/
func (s *NetworkSettingsService) DeleteSpProfile(spProfileName string) (*ResponseNetworkSettingsDeleteSpProfileV1, *resty.Response, error) {
	return s.DeleteSpProfileV1(spProfileName)
}

// Alias Function
/*
This method acts as an alias for the method `SetDhcpSettingsForASiteV1`
*/
func (s *NetworkSettingsService) SetDhcpSettingsForASite(id string, requestNetworkSettingsSetDhcpSettingsForASiteV1 *RequestNetworkSettingsSetDhcpSettingsForASiteV1) (*ResponseNetworkSettingsSetDhcpSettingsForASiteV1, *resty.Response, error) {
	return s.SetDhcpSettingsForASiteV1(id, requestNetworkSettingsSetDhcpSettingsForASiteV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetDeviceCredentialSettingsForASiteV1`
*/
func (s *NetworkSettingsService) GetDeviceCredentialSettingsForASite(id string, GetDeviceCredentialSettingsForASiteV1QueryParams *GetDeviceCredentialSettingsForASiteV1QueryParams) (*ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteV1, *resty.Response, error) {
	return s.GetDeviceCredentialSettingsForASiteV1(id, GetDeviceCredentialSettingsForASiteV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesSubpoolsIDsOfAGlobalIPAddressPoolV1`
*/
func (s *NetworkSettingsService) RetrievesSubpoolsIDsOfAGlobalIPAddressPool(globalIPAddressPoolID string, RetrievesSubpoolsIDsOfAGlobalIPAddressPoolV1QueryParams *RetrievesSubpoolsIDsOfAGlobalIPAddressPoolV1QueryParams) (*ResponseNetworkSettingsRetrievesSubpoolsIDsOfAGlobalIPAddressPoolV1, *resty.Response, error) {
	return s.RetrievesSubpoolsIDsOfAGlobalIPAddressPoolV1(globalIPAddressPoolID, RetrievesSubpoolsIDsOfAGlobalIPAddressPoolV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `RetrieveDHCPSettingsForASiteV1`
*/
func (s *NetworkSettingsService) RetrieveDHCPSettingsForASite(id string, RetrieveDHCPSettingsForASiteV1QueryParams *RetrieveDHCPSettingsForASiteV1QueryParams) (*ResponseNetworkSettingsRetrieveDHCPSettingsForASiteV1, *resty.Response, error) {
	return s.RetrieveDHCPSettingsForASiteV1(id, RetrieveDHCPSettingsForASiteV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `ReleaseReserveIPSubpoolV1`
*/
func (s *NetworkSettingsService) ReleaseReserveIPSubpool(id string) (*ResponseNetworkSettingsReleaseReserveIPSubpoolV1, *resty.Response, error) {
	return s.ReleaseReserveIPSubpoolV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `DeleteDeviceCredentialV1`
*/
func (s *NetworkSettingsService) DeleteDeviceCredential(id string) (*ResponseNetworkSettingsDeleteDeviceCredentialV1, *resty.Response, error) {
	return s.DeleteDeviceCredentialV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesAGlobalIPAddressPoolV1`
*/
func (s *NetworkSettingsService) RetrievesAGlobalIPAddressPool(id string) (*ResponseNetworkSettingsRetrievesAGlobalIPAddressPoolV1, *resty.Response, error) {
	return s.RetrievesAGlobalIPAddressPoolV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `GetServiceProviderDetailsV1`
*/
func (s *NetworkSettingsService) GetServiceProviderDetails() (*ResponseNetworkSettingsGetServiceProviderDetailsV1, *resty.Response, error) {
	return s.GetServiceProviderDetailsV1()
}

// Alias Function
/*
This method acts as an alias for the method `CreateSpProfileV1`
*/
func (s *NetworkSettingsService) CreateSpProfile(requestNetworkSettingsCreateSPProfileV1 *RequestNetworkSettingsCreateSpProfileV1) (*ResponseNetworkSettingsCreateSpProfileV1, *resty.Response, error) {
	return s.CreateSpProfileV1(requestNetworkSettingsCreateSPProfileV1)
}

// Alias Function
/*
This method acts as an alias for the method `RetrieveBannerSettingsForASiteV1`
*/
func (s *NetworkSettingsService) RetrieveBannerSettingsForASite(id string, RetrieveBannerSettingsForASiteV1QueryParams *RetrieveBannerSettingsForASiteV1QueryParams) (*ResponseNetworkSettingsRetrieveBannerSettingsForASiteV1, *resty.Response, error) {
	return s.RetrieveBannerSettingsForASiteV1(id, RetrieveBannerSettingsForASiteV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetNetworkV1`
*/
func (s *NetworkSettingsService) GetNetwork(GetNetworkV1QueryParams *GetNetworkV1QueryParams) (*ResponseNetworkSettingsGetNetworkV1, *resty.Response, error) {
	return s.GetNetworkV1(GetNetworkV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `SetAAASettingsForASiteV1`
*/
func (s *NetworkSettingsService) SetAAASettingsForASite(id string, requestNetworkSettingsSetAAASettingsForASiteV1 *RequestNetworkSettingsSetAAASettingsForASiteV1) (*ResponseNetworkSettingsSetAAASettingsForASiteV1, *resty.Response, error) {
	return s.SetAAASettingsForASiteV1(id, requestNetworkSettingsSetAAASettingsForASiteV1)
}

// Alias Function
/*
This method acts as an alias for the method `RetrieveCliTemplatesAttachedToANetworkProfileV1`
*/
func (s *NetworkSettingsService) RetrieveCliTemplatesAttachedToANetworkProfile(profileID string) (*ResponseNetworkSettingsRetrieveCliTemplatesAttachedToANetworkProfileV1, *resty.Response, error) {
	return s.RetrieveCliTemplatesAttachedToANetworkProfileV1(profileID)
}

// Alias Function
/*
This method acts as an alias for the method `RetrieveImageDistributionSettingsForASiteV1`
*/
func (s *NetworkSettingsService) RetrieveImageDistributionSettingsForASite(id string, RetrieveImageDistributionSettingsForASiteV1QueryParams *RetrieveImageDistributionSettingsForASiteV1QueryParams) (*ResponseNetworkSettingsRetrieveImageDistributionSettingsForASiteV1, *resty.Response, error) {
	return s.RetrieveImageDistributionSettingsForASiteV1(id, RetrieveImageDistributionSettingsForASiteV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `CountsGlobalIPAddressPoolsV1`
*/
func (s *NetworkSettingsService) CountsGlobalIPAddressPools() (*ResponseNetworkSettingsCountsGlobalIPAddressPoolsV1, *resty.Response, error) {
	return s.CountsGlobalIPAddressPoolsV1()
}

// Alias Function
/*
This method acts as an alias for the method `UpdatesAGlobalIPAddressPoolV1`
*/
func (s *NetworkSettingsService) UpdatesAGlobalIPAddressPool(id string, requestNetworkSettingsUpdatesAGlobalIPAddressPoolV1 *RequestNetworkSettingsUpdatesAGlobalIPAddressPoolV1) (*ResponseNetworkSettingsUpdatesAGlobalIPAddressPoolV1, *resty.Response, error) {
	return s.UpdatesAGlobalIPAddressPoolV1(id, requestNetworkSettingsUpdatesAGlobalIPAddressPoolV1)
}

// Alias Function
/*
This method acts as an alias for the method `UpdateSpProfileV1`
*/
func (s *NetworkSettingsService) UpdateSpProfile(requestNetworkSettingsUpdateSPProfileV1 *RequestNetworkSettingsUpdateSpProfileV1) (*ResponseNetworkSettingsUpdateSpProfileV1, *resty.Response, error) {
	return s.UpdateSpProfileV1(requestNetworkSettingsUpdateSPProfileV1)
}

// Alias Function
/*
This method acts as an alias for the method `SetTimeZoneForASiteV1`
*/
func (s *NetworkSettingsService) SetTimeZoneForASite(id string, requestNetworkSettingsSetTimeZoneForASiteV1 *RequestNetworkSettingsSetTimeZoneForASiteV1) (*ResponseNetworkSettingsSetTimeZoneForASiteV1, *resty.Response, error) {
	return s.SetTimeZoneForASiteV1(id, requestNetworkSettingsSetTimeZoneForASiteV1)
}

// Alias Function
/*
This method acts as an alias for the method `CreateDeviceCredentialsV1`
*/
func (s *NetworkSettingsService) CreateDeviceCredentials(requestNetworkSettingsCreateDeviceCredentialsV1 *RequestNetworkSettingsCreateDeviceCredentialsV1) (*ResponseNetworkSettingsCreateDeviceCredentialsV1, *resty.Response, error) {
	return s.CreateDeviceCredentialsV1(requestNetworkSettingsCreateDeviceCredentialsV1)
}

// Alias Function
/*
This method acts as an alias for the method `SyncNetworkDevicesCredentialV1`
*/
func (s *NetworkSettingsService) SyncNetworkDevicesCredential(requestNetworkSettingsSyncNetworkDevicesCredentialV1 *RequestNetworkSettingsSyncNetworkDevicesCredentialV1) (*ResponseNetworkSettingsSyncNetworkDevicesCredentialV1, *resty.Response, error) {
	return s.SyncNetworkDevicesCredentialV1(requestNetworkSettingsSyncNetworkDevicesCredentialV1)
}

// Alias Function
/*
This method acts as an alias for the method `AssignDeviceCredentialToSiteV1`
*/
func (s *NetworkSettingsService) AssignDeviceCredentialToSite(siteID string, requestNetworkSettingsAssignDeviceCredentialToSiteV1 *RequestNetworkSettingsAssignDeviceCredentialToSiteV1, AssignDeviceCredentialToSiteV1HeaderParams *AssignDeviceCredentialToSiteV1HeaderParams) (*ResponseNetworkSettingsAssignDeviceCredentialToSiteV1, *resty.Response, error) {
	return s.AssignDeviceCredentialToSiteV1(siteID, requestNetworkSettingsAssignDeviceCredentialToSiteV1, AssignDeviceCredentialToSiteV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `CreateNetworkV1`
*/
func (s *NetworkSettingsService) CreateNetwork(siteID string, requestNetworkSettingsCreateNetworkV1 *RequestNetworkSettingsCreateNetworkV1, CreateNetworkV1HeaderParams *CreateNetworkV1HeaderParams) (*ResponseNetworkSettingsCreateNetworkV1, *resty.Response, error) {
	return s.CreateNetworkV1(siteID, requestNetworkSettingsCreateNetworkV1, CreateNetworkV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `ReleaseAnIPAddressSubpoolV1`
*/
func (s *NetworkSettingsService) ReleaseAnIPAddressSubpool(id string) (*ResponseNetworkSettingsReleaseAnIPAddressSubpoolV1, *resty.Response, error) {
	return s.ReleaseAnIPAddressSubpoolV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `UpdateDeviceCredentialSettingsForASiteV1`
*/
func (s *NetworkSettingsService) UpdateDeviceCredentialSettingsForASite(id string, requestNetworkSettingsUpdateDeviceCredentialSettingsForASiteV1 *RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteV1) (*ResponseNetworkSettingsUpdateDeviceCredentialSettingsForASiteV1, *resty.Response, error) {
	return s.UpdateDeviceCredentialSettingsForASiteV1(id, requestNetworkSettingsUpdateDeviceCredentialSettingsForASiteV1)
}

// Alias Function
/*
This method acts as an alias for the method `UpdateDeviceCredentialsV1`
*/
func (s *NetworkSettingsService) UpdateDeviceCredentials(requestNetworkSettingsUpdateDeviceCredentialsV1 *RequestNetworkSettingsUpdateDeviceCredentialsV1) (*ResponseNetworkSettingsUpdateDeviceCredentialsV1, *resty.Response, error) {
	return s.UpdateDeviceCredentialsV1(requestNetworkSettingsUpdateDeviceCredentialsV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetGlobalPoolV1`
*/
func (s *NetworkSettingsService) GetGlobalPool(GetGlobalPoolV1QueryParams *GetGlobalPoolV1QueryParams) (*ResponseNetworkSettingsGetGlobalPoolV1, *resty.Response, error) {
	return s.GetGlobalPoolV1(GetGlobalPoolV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `RetrieveTimeZoneSettingsForASiteV1`
*/
func (s *NetworkSettingsService) RetrieveTimeZoneSettingsForASite(id string, RetrieveTimeZoneSettingsForASiteV1QueryParams *RetrieveTimeZoneSettingsForASiteV1QueryParams) (*ResponseNetworkSettingsRetrieveTimeZoneSettingsForASiteV1, *resty.Response, error) {
	return s.RetrieveTimeZoneSettingsForASiteV1(id, RetrieveTimeZoneSettingsForASiteV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `UpdateGlobalPoolV1`
*/
func (s *NetworkSettingsService) UpdateGlobalPool(requestNetworkSettingsUpdateGlobalPoolV1 *RequestNetworkSettingsUpdateGlobalPoolV1) (*ResponseNetworkSettingsUpdateGlobalPoolV1, *resty.Response, error) {
	return s.UpdateGlobalPoolV1(requestNetworkSettingsUpdateGlobalPoolV1)
}

// Alias Function
/*
This method acts as an alias for the method `RetrieveDNSSettingsForASiteV1`
*/
func (s *NetworkSettingsService) RetrieveDNSSettingsForASite(id string, RetrieveDNSSettingsForASiteV1QueryParams *RetrieveDNSSettingsForASiteV1QueryParams) (*ResponseNetworkSettingsRetrieveDNSSettingsForASiteV1, *resty.Response, error) {
	return s.RetrieveDNSSettingsForASiteV1(id, RetrieveDNSSettingsForASiteV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `SetNTPSettingsForASiteV1`
*/
func (s *NetworkSettingsService) SetNTPSettingsForASite(id string, requestNetworkSettingsSetNTPSettingsForASiteV1 *RequestNetworkSettingsSetNTPSettingsForASiteV1) (*ResponseNetworkSettingsSetNTPSettingsForASiteV1, *resty.Response, error) {
	return s.SetNTPSettingsForASiteV1(id, requestNetworkSettingsSetNTPSettingsForASiteV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetDeviceCredentialDetailsV1`
*/
func (s *NetworkSettingsService) GetDeviceCredentialDetails(GetDeviceCredentialDetailsV1QueryParams *GetDeviceCredentialDetailsV1QueryParams) (*ResponseNetworkSettingsGetDeviceCredentialDetailsV1, *resty.Response, error) {
	return s.GetDeviceCredentialDetailsV1(GetDeviceCredentialDetailsV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `CreateAGlobalIPAddressPoolV1`
*/
func (s *NetworkSettingsService) CreateAGlobalIPAddressPool(requestNetworkSettingsCreateAGlobalIPAddressPoolV1 *RequestNetworkSettingsCreateAGlobalIPAddressPoolV1) (*ResponseNetworkSettingsCreateAGlobalIPAddressPoolV1, *resty.Response, error) {
	return s.CreateAGlobalIPAddressPoolV1(requestNetworkSettingsCreateAGlobalIPAddressPoolV1)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesIPAddressSubpoolsV1`
*/
func (s *NetworkSettingsService) RetrievesIPAddressSubpools(RetrievesIPAddressSubpoolsV1QueryParams *RetrievesIPAddressSubpoolsV1QueryParams) (*ResponseNetworkSettingsRetrievesIPAddressSubpoolsV1, *resty.Response, error) {
	return s.RetrievesIPAddressSubpoolsV1(RetrievesIPAddressSubpoolsV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `RetrieveTelemetrySettingsForASiteV1`
*/
func (s *NetworkSettingsService) RetrieveTelemetrySettingsForASite(id string, RetrieveTelemetrySettingsForASiteV1QueryParams *RetrieveTelemetrySettingsForASiteV1QueryParams) (*ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteV1, *resty.Response, error) {
	return s.RetrieveTelemetrySettingsForASiteV1(id, RetrieveTelemetrySettingsForASiteV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `RetrieveAAASettingsForASiteV1`
*/
func (s *NetworkSettingsService) RetrieveAAASettingsForASite(id string, RetrieveAAASettingsForASiteV1QueryParams *RetrieveAAASettingsForASiteV1QueryParams) (*ResponseNetworkSettingsRetrieveAAASettingsForASiteV1, *resty.Response, error) {
	return s.RetrieveAAASettingsForASiteV1(id, RetrieveAAASettingsForASiteV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `UpdatesAnIPAddressSubpoolV1`
*/
func (s *NetworkSettingsService) UpdatesAnIPAddressSubpool(id string, requestNetworkSettingsUpdatesAnIPAddressSubpoolV1 *RequestNetworkSettingsUpdatesAnIPAddressSubpoolV1) (*ResponseNetworkSettingsUpdatesAnIPAddressSubpoolV1, *resty.Response, error) {
	return s.UpdatesAnIPAddressSubpoolV1(id, requestNetworkSettingsUpdatesAnIPAddressSubpoolV1)
}

// Alias Function
/*
This method acts as an alias for the method `CountsSubpoolsOfAGlobalIPAddressPoolV1`
*/
func (s *NetworkSettingsService) CountsSubpoolsOfAGlobalIPAddressPool(globalIPAddressPoolID string) (*ResponseNetworkSettingsCountsSubpoolsOfAGlobalIPAddressPoolV1, *resty.Response, error) {
	return s.CountsSubpoolsOfAGlobalIPAddressPoolV1(globalIPAddressPoolID)
}

// Alias Function
/*
This method acts as an alias for the method `DeleteGlobalIPPoolV1`
*/
func (s *NetworkSettingsService) DeleteGlobalIPPool(id string) (*ResponseNetworkSettingsDeleteGlobalIPPoolV1, *resty.Response, error) {
	return s.DeleteGlobalIPPoolV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `SetTelemetrySettingsForASiteV1`
*/
func (s *NetworkSettingsService) SetTelemetrySettingsForASite(id string, requestNetworkSettingsSetTelemetrySettingsForASiteV1 *RequestNetworkSettingsSetTelemetrySettingsForASiteV1) (*ResponseNetworkSettingsSetTelemetrySettingsForASiteV1, *resty.Response, error) {
	return s.SetTelemetrySettingsForASiteV1(id, requestNetworkSettingsSetTelemetrySettingsForASiteV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetReserveIPSubpoolV1`
*/
func (s *NetworkSettingsService) GetReserveIPSubpool(GetReserveIPSubpoolV1QueryParams *GetReserveIPSubpoolV1QueryParams) (*ResponseNetworkSettingsGetReserveIPSubpoolV1, *resty.Response, error) {
	return s.GetReserveIPSubpoolV1(GetReserveIPSubpoolV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `ReserveIPSubpoolV1`
*/
func (s *NetworkSettingsService) ReserveIPSubpool(siteID string, requestNetworkSettingsReserveIPSubpoolV1 *RequestNetworkSettingsReserveIPSubpoolV1) (*ResponseNetworkSettingsReserveIPSubpoolV1, *resty.Response, error) {
	return s.ReserveIPSubpoolV1(siteID, requestNetworkSettingsReserveIPSubpoolV1)
}

// Alias Function
/*
This method acts as an alias for the method `SetDNSSettingsForASiteV1`
*/
func (s *NetworkSettingsService) SetDNSSettingsForASite(id string, requestNetworkSettingsSetDNSSettingsForASiteV1 *RequestNetworkSettingsSetDNSSettingsForASiteV1) (*ResponseNetworkSettingsSetDNSSettingsForASiteV1, *resty.Response, error) {
	return s.SetDNSSettingsForASiteV1(id, requestNetworkSettingsSetDNSSettingsForASiteV1)
}

// Alias Function
/*
This method acts as an alias for the method `UpdateReserveIPSubpoolV1`
*/
func (s *NetworkSettingsService) UpdateReserveIPSubpool(siteID string, requestNetworkSettingsUpdateReserveIPSubpoolV1 *RequestNetworkSettingsUpdateReserveIPSubpoolV1, UpdateReserveIPSubpoolV1QueryParams *UpdateReserveIPSubpoolV1QueryParams) (*ResponseNetworkSettingsUpdateReserveIPSubpoolV1, *resty.Response, error) {
	return s.UpdateReserveIPSubpoolV1(siteID, requestNetworkSettingsUpdateReserveIPSubpoolV1, UpdateReserveIPSubpoolV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `RetrieveCountOfCliTemplatesAttachedToANetworkProfileV1`
*/
func (s *NetworkSettingsService) RetrieveCountOfCliTemplatesAttachedToANetworkProfile(profileID string) (*ResponseNetworkSettingsRetrieveCountOfCliTemplatesAttachedToANetworkProfileV1, *resty.Response, error) {
	return s.RetrieveCountOfCliTemplatesAttachedToANetworkProfileV1(profileID)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesAnIPAddressSubpoolV1`
*/
func (s *NetworkSettingsService) RetrievesAnIPAddressSubpool(id string) (*ResponseNetworkSettingsRetrievesAnIPAddressSubpoolV1, *resty.Response, error) {
	return s.RetrievesAnIPAddressSubpoolV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `DeleteAGlobalIPAddressPoolV1`
*/
func (s *NetworkSettingsService) DeleteAGlobalIPAddressPool(id string) (*ResponseNetworkSettingsDeleteAGlobalIPAddressPoolV1, *resty.Response, error) {
	return s.DeleteAGlobalIPAddressPoolV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `UpdateNetworkV1`
*/
func (s *NetworkSettingsService) UpdateNetwork(siteID string, requestNetworkSettingsUpdateNetworkV1 *RequestNetworkSettingsUpdateNetworkV1) (*ResponseNetworkSettingsUpdateNetworkV1, *resty.Response, error) {
	return s.UpdateNetworkV1(siteID, requestNetworkSettingsUpdateNetworkV1)
}

// Alias Function
/*
This method acts as an alias for the method `SetBannerSettingsForASiteV1`
*/
func (s *NetworkSettingsService) SetBannerSettingsForASite(id string, requestNetworkSettingsSetBannerSettingsForASiteV1 *RequestNetworkSettingsSetBannerSettingsForASiteV1) (*ResponseNetworkSettingsSetBannerSettingsForASiteV1, *resty.Response, error) {
	return s.SetBannerSettingsForASiteV1(id, requestNetworkSettingsSetBannerSettingsForASiteV1)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesGlobalIPAddressPoolsV1`
*/
func (s *NetworkSettingsService) RetrievesGlobalIPAddressPools(RetrievesGlobalIPAddressPoolsV1QueryParams *RetrievesGlobalIPAddressPoolsV1QueryParams) (*ResponseNetworkSettingsRetrievesGlobalIPAddressPoolsV1, *resty.Response, error) {
	return s.RetrievesGlobalIPAddressPoolsV1(RetrievesGlobalIPAddressPoolsV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetNetworkDevicesCredentialsSyncStatusV1`
*/
func (s *NetworkSettingsService) GetNetworkDevicesCredentialsSyncStatus(id string) (*ResponseNetworkSettingsGetNetworkDevicesCredentialsSyncStatusV1, *resty.Response, error) {
	return s.GetNetworkDevicesCredentialsSyncStatusV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `RetrieveNTPSettingsForASiteV1`
*/
func (s *NetworkSettingsService) RetrieveNTPSettingsForASite(id string, RetrieveNTPSettingsForASiteV1QueryParams *RetrieveNTPSettingsForASiteV1QueryParams) (*ResponseNetworkSettingsRetrieveNTPSettingsForASiteV1, *resty.Response, error) {
	return s.RetrieveNTPSettingsForASiteV1(id, RetrieveNTPSettingsForASiteV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `CreateGlobalPoolV1`
*/
func (s *NetworkSettingsService) CreateGlobalPool(requestNetworkSettingsCreateGlobalPoolV1 *RequestNetworkSettingsCreateGlobalPoolV1) (*ResponseNetworkSettingsCreateGlobalPoolV1, *resty.Response, error) {
	return s.CreateGlobalPoolV1(requestNetworkSettingsCreateGlobalPoolV1)
}

// Alias Function
/*
This method acts as an alias for the method `ReservecreateIPAddressSubpoolsV1`
*/
func (s *NetworkSettingsService) ReservecreateIPAddressSubpools(requestNetworkSettingsReservecreateIPAddressSubpoolsV1 *RequestNetworkSettingsReservecreateIPAddressSubpoolsV1) (*ResponseNetworkSettingsReservecreateIPAddressSubpoolsV1, *resty.Response, error) {
	return s.ReservecreateIPAddressSubpoolsV1(requestNetworkSettingsReservecreateIPAddressSubpoolsV1)
}
