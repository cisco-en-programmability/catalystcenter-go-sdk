package catalyst

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SensorsService service

type DeleteSensorTestV1QueryParams struct {
	TemplateName string `url:"templateName,omitempty"` //
}
type SensorsV1QueryParams struct {
	SiteID string `url:"siteId,omitempty"` //
}

type ResponseSensorsEditSensorTestTemplateV1 struct {
	Version  string                                           `json:"version,omitempty"`  // Version
	Response *ResponseSensorsEditSensorTestTemplateV1Response `json:"response,omitempty"` //
}
type ResponseSensorsEditSensorTestTemplateV1Response struct {
	Name                   string                                                             `json:"name,omitempty"`                   // The sensor test template name
	TypeID                 string                                                             `json:"_id,omitempty"`                    // The sensor test template unique identifier
	Version                *int                                                               `json:"version,omitempty"`                // The sensor test template version (must be 2)
	ModelVersion           *int                                                               `json:"modelVersion,omitempty"`           // Test template object model version (must be 2)
	StartTime              *int                                                               `json:"startTime,omitempty"`              // Start time
	LastModifiedTime       *int                                                               `json:"lastModifiedTime,omitempty"`       // Last modify time
	NumAssociatedSensor    *int                                                               `json:"numAssociatedSensor,omitempty"`    // Number of associated sensor
	Location               string                                                             `json:"location,omitempty"`               // Location string
	SiteHierarchy          string                                                             `json:"siteHierarchy,omitempty"`          // Site hierarchy
	Status                 string                                                             `json:"status,omitempty"`                 // Status of the test (RUNNING, NOTRUNNING)
	Connection             string                                                             `json:"connection,omitempty"`             // connection type of test: WIRED, WIRELESS, BOTH
	ActionInProgress       string                                                             `json:"actionInProgress,omitempty"`       // Indication of inprogress action
	Frequency              *ResponseSensorsEditSensorTestTemplateV1ResponseFrequency          `json:"frequency,omitempty"`              //
	RssiThreshold          *int                                                               `json:"rssiThreshold,omitempty"`          // RSSI threshold
	NumNeighborApThreshold *int                                                               `json:"numNeighborAPThreshold,omitempty"` // Number of neighboring AP threshold
	ScheduleInDays         *int                                                               `json:"scheduleInDays,omitempty"`         // Bit-wise value of scheduled test days
	WLANs                  []string                                                           `json:"wlans,omitempty"`                  // WLANs list
	SSIDs                  *[]ResponseSensorsEditSensorTestTemplateV1ResponseSSIDs            `json:"ssids,omitempty"`                  //
	Profiles               *[]ResponseSensorsEditSensorTestTemplateV1ResponseProfiles         `json:"profiles,omitempty"`               //
	TestScheduleMode       string                                                             `json:"testScheduleMode,omitempty"`       // Test schedule mode (ONDEMAND, DEDICATED, SCHEDULED, CONTINUOUS, RUNNOW)
	ShowWlcUpgradeBanner   *bool                                                              `json:"showWlcUpgradeBanner,omitempty"`   // Show WLC upgrade banner
	RadioAsSensorRemoved   *bool                                                              `json:"radioAsSensorRemoved,omitempty"`   // Radio as sensor removed
	EncryptionMode         string                                                             `json:"encryptionMode,omitempty"`         // Encryption mode
	RunNow                 string                                                             `json:"runNow,omitempty"`                 // Run now (YES, NO)
	LocationInfoList       *[]ResponseSensorsEditSensorTestTemplateV1ResponseLocationInfoList `json:"locationInfoList,omitempty"`       //
	Sensors                *[]ResponseSensorsEditSensorTestTemplateV1ResponseSensors          `json:"sensors,omitempty"`                //
	ApCoverage             *[]ResponseSensorsEditSensorTestTemplateV1ResponseApCoverage       `json:"apCoverage,omitempty"`             //
}
type ResponseSensorsEditSensorTestTemplateV1ResponseFrequency struct {
	Value *int   `json:"value,omitempty"` // Value of the unit
	Unit  string `json:"unit,omitempty"`  // Unit of the time value (NANOSECONDS, MICROSECONDS, MILLISECONDS, SECONDS, MINUTES, HOURS, DAYS)
}
type ResponseSensorsEditSensorTestTemplateV1ResponseSSIDs struct {
	Bands                     string                                                                   `json:"bands,omitempty"`                     // WIFI bands: 2.4GHz or 5GHz
	SSID                      string                                                                   `json:"ssid,omitempty"`                      // The SSID string
	ProfileName               string                                                                   `json:"profileName,omitempty"`               // The SSID profile name string
	NumAps                    *int                                                                     `json:"numAps,omitempty"`                    // Number of APs in the test
	NumSensors                *int                                                                     `json:"numSensors,omitempty"`                // Number of Sensors in the test
	Layer3WebAuthsecurity     string                                                                   `json:"layer3webAuthsecurity,omitempty"`     // Layer 3 WEB Auth security
	Layer3WebAuthuserName     string                                                                   `json:"layer3webAuthuserName,omitempty"`     // Layer 3 WEB Auth user name
	Layer3WebAuthpassword     string                                                                   `json:"layer3webAuthpassword,omitempty"`     // Layer 3 WEB Auth password
	Layer3WebAuthEmailAddress string                                                                   `json:"layer3webAuthEmailAddress,omitempty"` // Layer 3 WEB Auth email address
	ThirdParty                *ResponseSensorsEditSensorTestTemplateV1ResponseSSIDsThirdParty          `json:"thirdParty,omitempty"`                //
	ID                        *int                                                                     `json:"id,omitempty"`                        // Identification number
	WLANID                    *int                                                                     `json:"wlanId,omitempty"`                    // WLAN ID
	Wlc                       string                                                                   `json:"wlc,omitempty"`                       // WLC IP addres
	ValidFrom                 *int                                                                     `json:"validFrom,omitempty"`                 // Valid From UTC timestamp
	ValidTo                   *int                                                                     `json:"validTo,omitempty"`                   // Valid To UTC timestamp
	Status                    string                                                                   `json:"status,omitempty"`                    // WLAN status: ENABLED or DISABLED
	ProxyServer               string                                                                   `json:"proxyServer,omitempty"`               // Proxy server for onboarding SSID
	ProxyPort                 string                                                                   `json:"proxyPort,omitempty"`                 // Proxy server port
	ProxyUserName             string                                                                   `json:"proxyUserName,omitempty"`             // Proxy server user name
	ProxyPassword             string                                                                   `json:"proxyPassword,omitempty"`             // Proxy server password
	AuthType                  string                                                                   `json:"authType,omitempty"`                  // Authentication type: OPEN, WPA2_PSK, WPA2_EAP, WEB_AUTH, MAB, DOT1X, OTHER
	Psk                       string                                                                   `json:"psk,omitempty"`                       // Password of SSID when passwordType is ASCII
	Username                  string                                                                   `json:"username,omitempty"`                  // User name string for onboarding SSID
	Password                  string                                                                   `json:"password,omitempty"`                  // Password string for onboarding SSID
	PasswordType              string                                                                   `json:"passwordType,omitempty"`              // SSID password type: ASCII or HEX
	EapMethod                 string                                                                   `json:"eapMethod,omitempty"`                 // WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC
	Scep                      *bool                                                                    `json:"scep,omitempty"`                      // Secure certificate enrollment protocol: true or false or null for not applicable
	AuthProtocol              string                                                                   `json:"authProtocol,omitempty"`              // Auth protocol
	Certfilename              string                                                                   `json:"certfilename,omitempty"`              // Auth certificate file name
	Certxferprotocol          string                                                                   `json:"certxferprotocol,omitempty"`          // Certificate transfering protocol: HTTP or HTTPS
	Certstatus                string                                                                   `json:"certstatus,omitempty"`                // Certificate status: INACTIVE or ACTIVE
	Certpassphrase            string                                                                   `json:"certpassphrase,omitempty"`            // Certificate password phrase
	Certdownloadurl           string                                                                   `json:"certdownloadurl,omitempty"`           // Certificate download URL
	ExtWebAuthVirtualIP       string                                                                   `json:"extWebAuthVirtualIp,omitempty"`       // External WEB Auth virtual IP
	ExtWebAuth                *bool                                                                    `json:"extWebAuth,omitempty"`                // Indication of using external WEB Auth
	WhiteList                 *bool                                                                    `json:"whiteList,omitempty"`                 // Indication of being on allowed list
	ExtWebAuthPortal          string                                                                   `json:"extWebAuthPortal,omitempty"`          // External authentication portal
	ExtWebAuthAccessURL       string                                                                   `json:"extWebAuthAccessUrl,omitempty"`       // External WEB Auth access URL
	ExtWebAuthHTMLTag         *[]ResponseSensorsEditSensorTestTemplateV1ResponseSSIDsExtWebAuthHTMLTag `json:"extWebAuthHtmlTag,omitempty"`         //
	QosPolicy                 string                                                                   `json:"qosPolicy,omitempty"`                 // QoS policy: PlATINUM, GOLD, SILVER, BRONZE
	Tests                     *[]ResponseSensorsEditSensorTestTemplateV1ResponseSSIDsTests             `json:"tests,omitempty"`                     //
}
type ResponseSensorsEditSensorTestTemplateV1ResponseSSIDsThirdParty struct {
	Selected *bool `json:"selected,omitempty"` // true: the SSID is third party
}
type ResponseSensorsEditSensorTestTemplateV1ResponseSSIDsExtWebAuthHTMLTag struct {
	Label string `json:"label,omitempty"` // Label
	Tag   string `json:"tag,omitempty"`   // Tag
	Value string `json:"value,omitempty"` // Value
}
type ResponseSensorsEditSensorTestTemplateV1ResponseSSIDsTests struct {
	Name   string                                                             `json:"name,omitempty"`   // Name of the test
	Config *[]ResponseSensorsEditSensorTestTemplateV1ResponseSSIDsTestsConfig `json:"config,omitempty"` //
}
type ResponseSensorsEditSensorTestTemplateV1ResponseSSIDsTestsConfig struct {
	Domains        []string `json:"domains,omitempty"`        // DNS domain name
	Server         string   `json:"server,omitempty"`         // Ping, file transfer, mail, radius, ssh, or telnet server
	UserName       string   `json:"userName,omitempty"`       // User name
	Password       string   `json:"password,omitempty"`       // Password
	URL            string   `json:"url,omitempty"`            // URL
	Port           *int     `json:"port,omitempty"`           // Radius or WEB server port
	Protocol       string   `json:"protocol,omitempty"`       // Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)
	Servers        []string `json:"servers,omitempty"`        // IPerf server list
	Direction      string   `json:"direction,omitempty"`      // IPerf direction (UPLOAD, DOWNLOAD, BOTH)
	StartPort      *int     `json:"startPort,omitempty"`      // IPerf start port
	EndPort        *int     `json:"endPort,omitempty"`        // IPerf end port
	UDPBandwidth   *int     `json:"udpBandwidth,omitempty"`   // IPerf UDP bandwidth
	ProbeType      string   `json:"probeType,omitempty"`      // Probe type
	NumPackets     *int     `json:"numPackets,omitempty"`     // Number of packets
	PathToDownload string   `json:"pathToDownload,omitempty"` // File path for file transfer
	TransferType   string   `json:"transferType,omitempty"`   // File transfer type (UPLOAD, DOWNLOAD, BOTH)
	SharedSecret   string   `json:"sharedSecret,omitempty"`   // Shared secret
	NdtServer      string   `json:"ndtServer,omitempty"`      // NDT server
	NdtServerPort  string   `json:"ndtServerPort,omitempty"`  // NDT server port
	NdtServerPath  string   `json:"ndtServerPath,omitempty"`  // NDT server path
	UplinkTest     *bool    `json:"uplinkTest,omitempty"`     // Uplink test
	DownlinkTest   *bool    `json:"downlinkTest,omitempty"`   // Downlink test
	ProxyServer    string   `json:"proxyServer,omitempty"`    // Proxy server
	ProxyPort      string   `json:"proxyPort,omitempty"`      // Proxy port
	ProxyUserName  string   `json:"proxyUserName,omitempty"`  // Proxy user name
	ProxyPassword  string   `json:"proxyPassword,omitempty"`  // Proxy password
	UserNamePrompt string   `json:"userNamePrompt,omitempty"` // User name prompt
	PasswordPrompt string   `json:"passwordPrompt,omitempty"` // Password prompt
	ExitCommand    string   `json:"exitCommand,omitempty"`    // Exit command
	FinalPrompt    string   `json:"finalPrompt,omitempty"`    // Final prompt
}
type ResponseSensorsEditSensorTestTemplateV1ResponseProfiles struct {
	AuthType            string                                                                      `json:"authType,omitempty"`            // Authentication type: OPEN, WPA2_PSK, WPA2_EAP, WEB_AUTH, MAB, DOT1X, OTHER
	Psk                 string                                                                      `json:"psk,omitempty"`                 // Password of SSID when passwordType is ASCII
	Username            string                                                                      `json:"username,omitempty"`            // User name string for onboarding SSID
	Password            string                                                                      `json:"password,omitempty"`            // Password string for onboarding SSID
	PasswordType        string                                                                      `json:"passwordType,omitempty"`        // SSID password type: ASCII or HEX
	EapMethod           string                                                                      `json:"eapMethod,omitempty"`           // WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC
	Scep                *bool                                                                       `json:"scep,omitempty"`                // Secure certificate enrollment protocol: true or false or null for not applicable
	AuthProtocol        string                                                                      `json:"authProtocol,omitempty"`        // Auth protocol
	Certfilename        string                                                                      `json:"certfilename,omitempty"`        // Auth certificate file name
	Certxferprotocol    string                                                                      `json:"certxferprotocol,omitempty"`    // Certificate transfering protocol: HTTP or HTTPS
	Certstatus          string                                                                      `json:"certstatus,omitempty"`          // Certificate status: INACTIVE or ACTIVE
	Certpassphrase      string                                                                      `json:"certpassphrase,omitempty"`      // Certificate password phrase
	Certdownloadurl     string                                                                      `json:"certdownloadurl,omitempty"`     // Certificate download URL
	ExtWebAuthVirtualIP string                                                                      `json:"extWebAuthVirtualIp,omitempty"` // External WEB Auth virtual IP
	ExtWebAuth          *bool                                                                       `json:"extWebAuth,omitempty"`          // Indication of using external WEB Auth
	WhiteList           *bool                                                                       `json:"whiteList,omitempty"`           // Indication of being on allowed list
	ExtWebAuthPortal    string                                                                      `json:"extWebAuthPortal,omitempty"`    // External authentication portal
	ExtWebAuthAccessURL string                                                                      `json:"extWebAuthAccessUrl,omitempty"` // External WEB Auth access URL
	ExtWebAuthHTMLTag   *[]ResponseSensorsEditSensorTestTemplateV1ResponseProfilesExtWebAuthHTMLTag `json:"extWebAuthHtmlTag,omitempty"`   //
	QosPolicy           string                                                                      `json:"qosPolicy,omitempty"`           // QoS policy: PlATINUM, GOLD, SILVER, BRONZE
	Tests               *[]ResponseSensorsEditSensorTestTemplateV1ResponseProfilesTests             `json:"tests,omitempty"`               //
	ProfileName         string                                                                      `json:"profileName,omitempty"`         // Profile name
	DeviceType          string                                                                      `json:"deviceType,omitempty"`          // Device Type
	VLAN                string                                                                      `json:"vlan,omitempty"`                // VLAN
	LocationVLANList    *[]ResponseSensorsEditSensorTestTemplateV1ResponseProfilesLocationVLANList  `json:"locationVlanList,omitempty"`    //
}
type ResponseSensorsEditSensorTestTemplateV1ResponseProfilesExtWebAuthHTMLTag struct {
	Label string `json:"label,omitempty"` // Label
	Tag   string `json:"tag,omitempty"`   // Tag
	Value string `json:"value,omitempty"` // Value
}
type ResponseSensorsEditSensorTestTemplateV1ResponseProfilesTests struct {
	Name   string                                                                `json:"name,omitempty"`   // Name of the test
	Config *[]ResponseSensorsEditSensorTestTemplateV1ResponseProfilesTestsConfig `json:"config,omitempty"` //
}
type ResponseSensorsEditSensorTestTemplateV1ResponseProfilesTestsConfig struct {
	Domains        []string `json:"domains,omitempty"`        // DNS domain name
	Server         string   `json:"server,omitempty"`         // Ping, file transfer, mail, radius, ssh, or telnet server
	UserName       string   `json:"userName,omitempty"`       // User name
	Password       string   `json:"password,omitempty"`       // Password
	URL            string   `json:"url,omitempty"`            // URL
	Port           *int     `json:"port,omitempty"`           // Radius or WEB server port
	Protocol       string   `json:"protocol,omitempty"`       // Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)
	Servers        []string `json:"servers,omitempty"`        // IPerf server list
	Direction      string   `json:"direction,omitempty"`      // IPerf direction (UPLOAD, DOWNLOAD, BOTH)
	StartPort      *int     `json:"startPort,omitempty"`      // IPerf start port
	EndPort        *int     `json:"endPort,omitempty"`        // IPerf end port
	UDPBandwidth   *int     `json:"udpBandwidth,omitempty"`   // IPerf UDP bandwidth
	ProbeType      string   `json:"probeType,omitempty"`      // Probe type
	NumPackets     *int     `json:"numPackets,omitempty"`     // Number of packets
	PathToDownload string   `json:"pathToDownload,omitempty"` // File path for file transfer
	TransferType   string   `json:"transferType,omitempty"`   // File transfer type (UPLOAD, DOWNLOAD, BOTH)
	SharedSecret   string   `json:"sharedSecret,omitempty"`   // Shared secret
	NdtServer      string   `json:"ndtServer,omitempty"`      // NDT server
	NdtServerPort  string   `json:"ndtServerPort,omitempty"`  // NDT server port
	NdtServerPath  string   `json:"ndtServerPath,omitempty"`  // NDT server path
	UplinkTest     *bool    `json:"uplinkTest,omitempty"`     // Uplink test
	DownlinkTest   *bool    `json:"downlinkTest,omitempty"`   // Downlink test
	ProxyServer    string   `json:"proxyServer,omitempty"`    // Proxy server
	ProxyPort      string   `json:"proxyPort,omitempty"`      // Proxy port
	ProxyUserName  string   `json:"proxyUserName,omitempty"`  // Proxy user name
	ProxyPassword  string   `json:"proxyPassword,omitempty"`  // Proxy password
	UserNamePrompt string   `json:"userNamePrompt,omitempty"` // User name prompt
	PasswordPrompt string   `json:"passwordPrompt,omitempty"` // Password prompt
	ExitCommand    string   `json:"exitCommand,omitempty"`    // Exit command
	FinalPrompt    string   `json:"finalPrompt,omitempty"`    // Final prompt
}
type ResponseSensorsEditSensorTestTemplateV1ResponseProfilesLocationVLANList struct {
	LocationID string   `json:"locationId,omitempty"` // Site UUID
	VLANs      []string `json:"vlans,omitempty"`      // Array of VLANs
}
type ResponseSensorsEditSensorTestTemplateV1ResponseLocationInfoList struct {
	LocationID           string   `json:"locationId,omitempty"`           // Site UUID
	LocationType         string   `json:"locationType,omitempty"`         // Site type
	AllSensors           *bool    `json:"allSensors,omitempty"`           // Use all sensors in the site for test
	SiteHierarchy        string   `json:"siteHierarchy,omitempty"`        // Site name hierarhy
	MacAddressList       []string `json:"macAddressList,omitempty"`       // MAC addresses
	ManagementVLAN       string   `json:"managementVlan,omitempty"`       // Management VLAN
	CustomManagementVLAN *bool    `json:"customManagementVlan,omitempty"` // Custom Management VLAN
}
type ResponseSensorsEditSensorTestTemplateV1ResponseSensors struct {
	Name                    string                                                                  `json:"name,omitempty"`                    // Sensor name
	MacAddress              string                                                                  `json:"macAddress,omitempty"`              // MAC address
	SwitchMac               string                                                                  `json:"switchMac,omitempty"`               // Switch MAC address
	SwitchUUID              string                                                                  `json:"switchUuid,omitempty"`              // Switch device UUID
	SwitchSerialNumber      string                                                                  `json:"switchSerialNumber,omitempty"`      // Switch serial number
	MarkedForUninstall      *bool                                                                   `json:"markedForUninstall,omitempty"`      // Is marked for uninstall
	IPAddress               string                                                                  `json:"ipAddress,omitempty"`               // IP address
	HostName                string                                                                  `json:"hostName,omitempty"`                // Host name
	WiredApplicationStatus  string                                                                  `json:"wiredApplicationStatus,omitempty"`  // Wired application status
	WiredApplicationMessage string                                                                  `json:"wiredApplicationMessage,omitempty"` // Wired application message
	Assigned                *bool                                                                   `json:"assigned,omitempty"`                // Is assigned
	Status                  string                                                                  `json:"status,omitempty"`                  // Sensor device status: UP, DOWN, REBOOT
	XorSensor               *bool                                                                   `json:"xorSensor,omitempty"`               // Is XOR sensor
	TargetAPs               []string                                                                `json:"targetAPs,omitempty"`               // Array of target APs
	RunNow                  string                                                                  `json:"runNow,omitempty"`                  // Run now: YES, NO
	LocationID              string                                                                  `json:"locationId,omitempty"`              // Site UUID
	AllSensorAddition       *bool                                                                   `json:"allSensorAddition,omitempty"`       // Is all sensor addition
	ConfigUpdated           string                                                                  `json:"configUpdated,omitempty"`           // Configuration updated: YES, NO
	SensorType              string                                                                  `json:"sensorType,omitempty"`              // Sensor type
	TestMacAddresses        *ResponseSensorsEditSensorTestTemplateV1ResponseSensorsTestMacAddresses `json:"testMacAddresses,omitempty"`        // A string-string test MAC address
	ID                      string                                                                  `json:"id,omitempty"`                      // Sensor ID
	ServicePolicy           string                                                                  `json:"servicePolicy,omitempty"`           // Service policy
	IPerfInfo               *ResponseSensorsEditSensorTestTemplateV1ResponseSensorsIPerfInfo        `json:"iPerfInfo,omitempty"`               // A string-stringList iPerf information
}
type ResponseSensorsEditSensorTestTemplateV1ResponseSensorsTestMacAddresses interface{}
type ResponseSensorsEditSensorTestTemplateV1ResponseSensorsIPerfInfo interface{}
type ResponseSensorsEditSensorTestTemplateV1ResponseApCoverage struct {
	Bands             string `json:"bands,omitempty"`             // The WIFI bands
	NumberOfApsToTest *int   `json:"numberOfApsToTest,omitempty"` // Number of APs to test
	RssiThreshold     *int   `json:"rssiThreshold,omitempty"`     // RSSI threshold
}
type ResponseSensorsCreateSensorTestTemplateV1 struct {
	Version  string                                             `json:"version,omitempty"`  // Version
	Response *ResponseSensorsCreateSensorTestTemplateV1Response `json:"response,omitempty"` //
}
type ResponseSensorsCreateSensorTestTemplateV1Response struct {
	Name                   string                                                               `json:"name,omitempty"`                   // The sensor test template name
	TypeID                 string                                                               `json:"_id,omitempty"`                    // (Used in edit only) The sensor test template unique identifier
	Version                *int                                                                 `json:"version,omitempty"`                // The sensor test template version (must be 2)
	ModelVersion           *int                                                                 `json:"modelVersion,omitempty"`           // Test template object model version (must be 2)
	StartTime              *int                                                                 `json:"startTime,omitempty"`              // Start time
	LastModifiedTime       *int                                                                 `json:"lastModifiedTime,omitempty"`       // Last modify time
	NumAssociatedSensor    *int                                                                 `json:"numAssociatedSensor,omitempty"`    // Number of associated sensor
	Location               string                                                               `json:"location,omitempty"`               // Location string
	SiteHierarchy          string                                                               `json:"siteHierarchy,omitempty"`          // Site hierarchy
	Status                 string                                                               `json:"status,omitempty"`                 // Status of the test (RUNNING, NOTRUNNING)
	Connection             string                                                               `json:"connection,omitempty"`             // connection type of test: WIRED, WIRELESS, BOTH
	ActionInProgress       string                                                               `json:"actionInProgress,omitempty"`       // Indication of inprogress action
	Frequency              *ResponseSensorsCreateSensorTestTemplateV1ResponseFrequency          `json:"frequency,omitempty"`              //
	RssiThreshold          *int                                                                 `json:"rssiThreshold,omitempty"`          // RSSI threshold
	NumNeighborApThreshold *int                                                                 `json:"numNeighborAPThreshold,omitempty"` // Number of neighboring AP threshold
	ScheduleInDays         *int                                                                 `json:"scheduleInDays,omitempty"`         // Bit-wise value of scheduled test days
	WLANs                  []string                                                             `json:"wlans,omitempty"`                  // WLANs list
	SSIDs                  *[]ResponseSensorsCreateSensorTestTemplateV1ResponseSSIDs            `json:"ssids,omitempty"`                  //
	Profiles               *[]ResponseSensorsCreateSensorTestTemplateV1ResponseProfiles         `json:"profiles,omitempty"`               //
	TestScheduleMode       string                                                               `json:"testScheduleMode,omitempty"`       // Test schedule mode (ONDEMAND, DEDICATED, SCHEDULED, CONTINUOUS, RUNNOW)
	ShowWlcUpgradeBanner   *bool                                                                `json:"showWlcUpgradeBanner,omitempty"`   // Show WLC upgrade banner
	RadioAsSensorRemoved   *bool                                                                `json:"radioAsSensorRemoved,omitempty"`   // Radio as sensor removed
	EncryptionMode         string                                                               `json:"encryptionMode,omitempty"`         // Encryption mode
	RunNow                 string                                                               `json:"runNow,omitempty"`                 // Run now (YES, NO)
	LocationInfoList       *[]ResponseSensorsCreateSensorTestTemplateV1ResponseLocationInfoList `json:"locationInfoList,omitempty"`       //
	Sensors                *[]ResponseSensorsCreateSensorTestTemplateV1ResponseSensors          `json:"sensors,omitempty"`                //
	ApCoverage             *[]ResponseSensorsCreateSensorTestTemplateV1ResponseApCoverage       `json:"apCoverage,omitempty"`             //
}
type ResponseSensorsCreateSensorTestTemplateV1ResponseFrequency struct {
	Value *int   `json:"value,omitempty"` // Value of the unit
	Unit  string `json:"unit,omitempty"`  // Unit of the time value (NANOSECONDS, MICROSECONDS, MILLISECONDS, SECONDS, MINUTES, HOURS, DAYS)
}
type ResponseSensorsCreateSensorTestTemplateV1ResponseSSIDs struct {
	Bands                     string                                                                     `json:"bands,omitempty"`                     // WIFI bands: 2.4GHz or 5GHz
	SSID                      string                                                                     `json:"ssid,omitempty"`                      // The SSID string
	ProfileName               string                                                                     `json:"profileName,omitempty"`               // The SSID profile name string
	NumAps                    *int                                                                       `json:"numAps,omitempty"`                    // Number of APs in the test
	NumSensors                *int                                                                       `json:"numSensors,omitempty"`                // Number of Sensors in the test
	Layer3WebAuthsecurity     string                                                                     `json:"layer3webAuthsecurity,omitempty"`     // Layer 3 WEB Auth security
	Layer3WebAuthuserName     string                                                                     `json:"layer3webAuthuserName,omitempty"`     // Layer 3 WEB Auth user name
	Layer3WebAuthpassword     string                                                                     `json:"layer3webAuthpassword,omitempty"`     // Layer 3 WEB Auth password
	Layer3WebAuthEmailAddress string                                                                     `json:"layer3webAuthEmailAddress,omitempty"` // Layer 3 WEB Auth email address
	ThirdParty                *ResponseSensorsCreateSensorTestTemplateV1ResponseSSIDsThirdParty          `json:"thirdParty,omitempty"`                //
	ID                        *int                                                                       `json:"id,omitempty"`                        // Identification number
	WLANID                    *int                                                                       `json:"wlanId,omitempty"`                    // WLAN ID
	Wlc                       string                                                                     `json:"wlc,omitempty"`                       // WLC IP addres
	ValidFrom                 *int                                                                       `json:"validFrom,omitempty"`                 // Valid From UTC timestamp
	ValidTo                   *int                                                                       `json:"validTo,omitempty"`                   // Valid To UTC timestamp
	Status                    string                                                                     `json:"status,omitempty"`                    // WLAN status: ENABLED or DISABLED
	ProxyServer               string                                                                     `json:"proxyServer,omitempty"`               // Proxy server for onboarding SSID
	ProxyPort                 string                                                                     `json:"proxyPort,omitempty"`                 // Proxy server port
	ProxyUserName             string                                                                     `json:"proxyUserName,omitempty"`             // Proxy server user name
	ProxyPassword             string                                                                     `json:"proxyPassword,omitempty"`             // Proxy server password
	AuthType                  string                                                                     `json:"authType,omitempty"`                  // Authentication type: OPEN, WPA2_PSK, WPA2_EaP, WEB_AUTH, MAB, DOT1X, OTHER
	Psk                       string                                                                     `json:"psk,omitempty"`                       // Password of SSID when passwordType is ASCII
	Username                  string                                                                     `json:"username,omitempty"`                  // User name string for onboarding SSID
	Password                  string                                                                     `json:"password,omitempty"`                  // Password string for onboarding SSID
	PasswordType              string                                                                     `json:"passwordType,omitempty"`              // SSID password type: ASCII or HEX
	EapMethod                 string                                                                     `json:"eapMethod,omitempty"`                 // WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC
	Scep                      *bool                                                                      `json:"scep,omitempty"`                      // Secure certificate enrollment protocol: true or false or null for not applicable
	AuthProtocol              string                                                                     `json:"authProtocol,omitempty"`              // Auth protocol
	Certfilename              string                                                                     `json:"certfilename,omitempty"`              // Auth certificate file name
	Certxferprotocol          string                                                                     `json:"certxferprotocol,omitempty"`          // Certificate transfering protocol: HTTP or HTTPS
	Certstatus                string                                                                     `json:"certstatus,omitempty"`                // Certificate status: INACTIVE or ACTIVE
	Certpassphrase            string                                                                     `json:"certpassphrase,omitempty"`            // Certificate password phrase
	Certdownloadurl           string                                                                     `json:"certdownloadurl,omitempty"`           // Certificate download URL
	ExtWebAuthVirtualIP       string                                                                     `json:"extWebAuthVirtualIp,omitempty"`       // External WEB Auth virtual IP
	ExtWebAuth                *bool                                                                      `json:"extWebAuth,omitempty"`                // Indication of using external WEB Auth
	WhiteList                 *bool                                                                      `json:"whiteList,omitempty"`                 // Indication of being on allowed list
	ExtWebAuthPortal          string                                                                     `json:"extWebAuthPortal,omitempty"`          // External authentication portal
	ExtWebAuthAccessURL       string                                                                     `json:"extWebAuthAccessUrl,omitempty"`       // External WEB Auth access URL
	ExtWebAuthHTMLTag         *[]ResponseSensorsCreateSensorTestTemplateV1ResponseSSIDsExtWebAuthHTMLTag `json:"extWebAuthHtmlTag,omitempty"`         //
	QosPolicy                 string                                                                     `json:"qosPolicy,omitempty"`                 // QoS policy: PlATINUM, GOLD, SILVER, BRONZE
	Tests                     *[]ResponseSensorsCreateSensorTestTemplateV1ResponseSSIDsTests             `json:"tests,omitempty"`                     //
}
type ResponseSensorsCreateSensorTestTemplateV1ResponseSSIDsThirdParty struct {
	Selected *bool `json:"selected,omitempty"` // true: the SSID is third party
}
type ResponseSensorsCreateSensorTestTemplateV1ResponseSSIDsExtWebAuthHTMLTag struct {
	Label string `json:"label,omitempty"` // Label
	Tag   string `json:"tag,omitempty"`   // Tag
	Value string `json:"value,omitempty"` // Value
}
type ResponseSensorsCreateSensorTestTemplateV1ResponseSSIDsTests struct {
	Name   string                                                               `json:"name,omitempty"`   // Name of the test
	Config *[]ResponseSensorsCreateSensorTestTemplateV1ResponseSSIDsTestsConfig `json:"config,omitempty"` //
}
type ResponseSensorsCreateSensorTestTemplateV1ResponseSSIDsTestsConfig struct {
	Domains        []string `json:"domains,omitempty"`        // DNS domain name
	Server         string   `json:"server,omitempty"`         // Ping, file transfer, mail, radius, ssh, or telnet server
	UserName       string   `json:"userName,omitempty"`       // User name
	Password       string   `json:"password,omitempty"`       // Password
	URL            string   `json:"url,omitempty"`            // URL
	Port           *int     `json:"port,omitempty"`           // Radius or WEB server port
	Protocol       string   `json:"protocol,omitempty"`       // Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)
	Servers        []string `json:"servers,omitempty"`        // IPerf server list
	Direction      string   `json:"direction,omitempty"`      // IPerf direction (UPLOAD, DOWNLOAD, BOTH)
	StartPort      *int     `json:"startPort,omitempty"`      // IPerf start port
	EndPort        *int     `json:"endPort,omitempty"`        // IPerf end port
	UDPBandwidth   *int     `json:"udpBandwidth,omitempty"`   // IPerf UDP bandwidth
	ProbeType      string   `json:"probeType,omitempty"`      // Probe type
	NumPackets     *int     `json:"numPackets,omitempty"`     // Number of packets
	PathToDownload string   `json:"pathToDownload,omitempty"` // File path for file transfer
	TransferType   string   `json:"transferType,omitempty"`   // File transfer type (UPLOAD, DOWNLOAD, BOTH)
	SharedSecret   string   `json:"sharedSecret,omitempty"`   // Shared secret
	NdtServer      string   `json:"ndtServer,omitempty"`      // NDT server
	NdtServerPort  string   `json:"ndtServerPort,omitempty"`  // NDT server port
	NdtServerPath  string   `json:"ndtServerPath,omitempty"`  // NDT server path
	UplinkTest     *bool    `json:"uplinkTest,omitempty"`     // Uplink test
	DownlinkTest   *bool    `json:"downlinkTest,omitempty"`   // Downlink test
	ProxyServer    string   `json:"proxyServer,omitempty"`    // Proxy server
	ProxyPort      string   `json:"proxyPort,omitempty"`      // Proxy port
	ProxyUserName  string   `json:"proxyUserName,omitempty"`  // Proxy user name
	ProxyPassword  string   `json:"proxyPassword,omitempty"`  // Proxy password
	UserNamePrompt string   `json:"userNamePrompt,omitempty"` // User name prompt
	PasswordPrompt string   `json:"passwordPrompt,omitempty"` // Password prompt
	ExitCommand    string   `json:"exitCommand,omitempty"`    // Exit command
	FinalPrompt    string   `json:"finalPrompt,omitempty"`    // Final prompt
}
type ResponseSensorsCreateSensorTestTemplateV1ResponseProfiles struct {
	AuthType            string                                                                        `json:"authType,omitempty"`            // Authentication type: OPEN, WPA2_PSK, WPA2_EaP, WEB_AUTH, MAB, DOT1X, OTHER
	Psk                 string                                                                        `json:"psk,omitempty"`                 // Password of SSID when passwordType is ASCII
	Username            string                                                                        `json:"username,omitempty"`            // User name string for onboarding SSID
	Password            string                                                                        `json:"password,omitempty"`            // Password string for onboarding SSID
	PasswordType        string                                                                        `json:"passwordType,omitempty"`        // SSID password type: ASCII or HEX
	EapMethod           string                                                                        `json:"eapMethod,omitempty"`           // WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC
	Scep                *bool                                                                         `json:"scep,omitempty"`                // Secure certificate enrollment protocol: true or false or null for not applicable
	AuthProtocol        string                                                                        `json:"authProtocol,omitempty"`        // Auth protocol
	Certfilename        string                                                                        `json:"certfilename,omitempty"`        // Auth certificate file name
	Certxferprotocol    string                                                                        `json:"certxferprotocol,omitempty"`    // Certificate transfering protocol: HTTP or HTTPS
	Certstatus          string                                                                        `json:"certstatus,omitempty"`          // Certificate status: INACTIVE or ACTIVE
	Certpassphrase      string                                                                        `json:"certpassphrase,omitempty"`      // Certificate password phrase
	Certdownloadurl     string                                                                        `json:"certdownloadurl,omitempty"`     // Certificate download URL
	ExtWebAuthVirtualIP string                                                                        `json:"extWebAuthVirtualIp,omitempty"` // External WEB Auth virtual IP
	ExtWebAuth          *bool                                                                         `json:"extWebAuth,omitempty"`          // Indication of using external WEB Auth
	WhiteList           *bool                                                                         `json:"whiteList,omitempty"`           // Indication of being on allowed list
	ExtWebAuthPortal    string                                                                        `json:"extWebAuthPortal,omitempty"`    // External authentication portal
	ExtWebAuthAccessURL string                                                                        `json:"extWebAuthAccessUrl,omitempty"` // External WEB Auth access URL
	ExtWebAuthHTMLTag   *[]ResponseSensorsCreateSensorTestTemplateV1ResponseProfilesExtWebAuthHTMLTag `json:"extWebAuthHtmlTag,omitempty"`   //
	QosPolicy           string                                                                        `json:"qosPolicy,omitempty"`           // QoS policy: PlATINUM, GOLD, SILVER, BRONZE
	Tests               *[]ResponseSensorsCreateSensorTestTemplateV1ResponseProfilesTests             `json:"tests,omitempty"`               //
	ProfileName         string                                                                        `json:"profileName,omitempty"`         // Profile name
	DeviceType          string                                                                        `json:"deviceType,omitempty"`          // Device Type
	VLAN                string                                                                        `json:"vlan,omitempty"`                // VLAN
	LocationVLANList    *[]ResponseSensorsCreateSensorTestTemplateV1ResponseProfilesLocationVLANList  `json:"locationVlanList,omitempty"`    //
}
type ResponseSensorsCreateSensorTestTemplateV1ResponseProfilesExtWebAuthHTMLTag struct {
	Label string `json:"label,omitempty"` // Label
	Tag   string `json:"tag,omitempty"`   // Tag
	Value string `json:"value,omitempty"` // Value
}
type ResponseSensorsCreateSensorTestTemplateV1ResponseProfilesTests struct {
	Name   string                                                                  `json:"name,omitempty"`   // Name of the test
	Config *[]ResponseSensorsCreateSensorTestTemplateV1ResponseProfilesTestsConfig `json:"config,omitempty"` //
}
type ResponseSensorsCreateSensorTestTemplateV1ResponseProfilesTestsConfig struct {
	Domains        []string `json:"domains,omitempty"`        // DNS domain name
	Server         string   `json:"server,omitempty"`         // Ping, file transfer, mail, radius, ssh, or telnet server
	UserName       string   `json:"userName,omitempty"`       // User name
	Password       string   `json:"password,omitempty"`       // Password
	URL            string   `json:"url,omitempty"`            // URL
	Port           *int     `json:"port,omitempty"`           // Radius or WEB server port
	Protocol       string   `json:"protocol,omitempty"`       // Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)
	Servers        []string `json:"servers,omitempty"`        // IPerf server list
	Direction      string   `json:"direction,omitempty"`      // IPerf direction (UPLOAD, DOWNLOAD, BOTH)
	StartPort      *int     `json:"startPort,omitempty"`      // IPerf start port
	EndPort        *int     `json:"endPort,omitempty"`        // IPerf end port
	UDPBandwidth   *int     `json:"udpBandwidth,omitempty"`   // IPerf UDP bandwidth
	ProbeType      string   `json:"probeType,omitempty"`      // Probe type
	NumPackets     *int     `json:"numPackets,omitempty"`     // Number of packets
	PathToDownload string   `json:"pathToDownload,omitempty"` // File path for file transfer
	TransferType   string   `json:"transferType,omitempty"`   // File transfer type (UPLOAD, DOWNLOAD, BOTH)
	SharedSecret   string   `json:"sharedSecret,omitempty"`   // Shared secret
	NdtServer      string   `json:"ndtServer,omitempty"`      // NDT server
	NdtServerPort  string   `json:"ndtServerPort,omitempty"`  // NDT server port
	NdtServerPath  string   `json:"ndtServerPath,omitempty"`  // NDT server path
	UplinkTest     *bool    `json:"uplinkTest,omitempty"`     // Uplink test
	DownlinkTest   *bool    `json:"downlinkTest,omitempty"`   // Downlink test
	ProxyServer    string   `json:"proxyServer,omitempty"`    // Proxy server
	ProxyPort      string   `json:"proxyPort,omitempty"`      // Proxy port
	ProxyUserName  string   `json:"proxyUserName,omitempty"`  // Proxy user name
	ProxyPassword  string   `json:"proxyPassword,omitempty"`  // Proxy password
	UserNamePrompt string   `json:"userNamePrompt,omitempty"` // User name prompt
	PasswordPrompt string   `json:"passwordPrompt,omitempty"` // Password prompt
	ExitCommand    string   `json:"exitCommand,omitempty"`    // Exit command
	FinalPrompt    string   `json:"finalPrompt,omitempty"`    // Final prompt
}
type ResponseSensorsCreateSensorTestTemplateV1ResponseProfilesLocationVLANList struct {
	LocationID string   `json:"locationId,omitempty"` // Site UUID
	VLANs      []string `json:"vlans,omitempty"`      // Array of VLANs
}
type ResponseSensorsCreateSensorTestTemplateV1ResponseLocationInfoList struct {
	LocationID           string   `json:"locationId,omitempty"`           // Site UUID
	LocationType         string   `json:"locationType,omitempty"`         // Site type
	AllSensors           *bool    `json:"allSensors,omitempty"`           // Use all sensors in the site for test
	SiteHierarchy        string   `json:"siteHierarchy,omitempty"`        // Site name hierarhy
	MacAddressList       []string `json:"macAddressList,omitempty"`       // MAC addresses
	ManagementVLAN       string   `json:"managementVlan,omitempty"`       // Management VLAN
	CustomManagementVLAN *bool    `json:"customManagementVlan,omitempty"` // Custom Management VLAN
}
type ResponseSensorsCreateSensorTestTemplateV1ResponseSensors struct {
	Name                    string                                                                    `json:"name,omitempty"`                    // Sensor name
	MacAddress              string                                                                    `json:"macAddress,omitempty"`              // MAC address
	SwitchMac               string                                                                    `json:"switchMac,omitempty"`               // Switch MAC address
	SwitchUUID              string                                                                    `json:"switchUuid,omitempty"`              // Switch device UUID
	SwitchSerialNumber      string                                                                    `json:"switchSerialNumber,omitempty"`      // Switch serial number
	MarkedForUninstall      *bool                                                                     `json:"markedForUninstall,omitempty"`      // Is marked for uninstall
	IPAddress               string                                                                    `json:"ipAddress,omitempty"`               // IP address
	HostName                string                                                                    `json:"hostName,omitempty"`                // Host name
	WiredApplicationStatus  string                                                                    `json:"wiredApplicationStatus,omitempty"`  // Wired application status
	WiredApplicationMessage string                                                                    `json:"wiredApplicationMessage,omitempty"` // Wired application message
	Assigned                *bool                                                                     `json:"assigned,omitempty"`                // Is assigned
	Status                  string                                                                    `json:"status,omitempty"`                  // Sensor device status: UP, DOWN, REBOOT
	XorSensor               *bool                                                                     `json:"xorSensor,omitempty"`               // Is XOR sensor
	TargetAPs               []string                                                                  `json:"targetAPs,omitempty"`               // Array of target APs
	RunNow                  string                                                                    `json:"runNow,omitempty"`                  // Run now: YES, NO
	LocationID              string                                                                    `json:"locationId,omitempty"`              // Site UUID
	AllSensorAddition       *bool                                                                     `json:"allSensorAddition,omitempty"`       // Is all sensor addition
	ConfigUpdated           string                                                                    `json:"configUpdated,omitempty"`           // Configuration updated: YES, NO
	SensorType              string                                                                    `json:"sensorType,omitempty"`              // Sensor type
	TestMacAddresses        *ResponseSensorsCreateSensorTestTemplateV1ResponseSensorsTestMacAddresses `json:"testMacAddresses,omitempty"`        // A string-string test MAC address
	ID                      string                                                                    `json:"id,omitempty"`                      // Sensor ID
	ServicePolicy           string                                                                    `json:"servicePolicy,omitempty"`           // Service policy
	IPerfInfo               *ResponseSensorsCreateSensorTestTemplateV1ResponseSensorsIPerfInfo        `json:"iPerfInfo,omitempty"`               // A string-stringList iPerf information
}
type ResponseSensorsCreateSensorTestTemplateV1ResponseSensorsTestMacAddresses interface{}
type ResponseSensorsCreateSensorTestTemplateV1ResponseSensorsIPerfInfo interface{}
type ResponseSensorsCreateSensorTestTemplateV1ResponseApCoverage struct {
	Bands             string `json:"bands,omitempty"`             // The WIFI bands
	NumberOfApsToTest *int   `json:"numberOfApsToTest,omitempty"` // Number of APs to test
	RssiThreshold     *int   `json:"rssiThreshold,omitempty"`     // RSSI threshold
}
type ResponseSensorsDeleteSensorTestV1 struct {
	Version  string                                     `json:"version,omitempty"`  // Version
	Response *ResponseSensorsDeleteSensorTestV1Response `json:"response,omitempty"` //
}
type ResponseSensorsDeleteSensorTestV1Response struct {
	TemplateName string `json:"templateName,omitempty"` // Test template name to be delete
	Status       string `json:"status,omitempty"`       // Status of the DELETE action
}
type ResponseSensorsSensorsV1 struct {
	Version  string                              `json:"version,omitempty"`  // Version string of this API
	Response *[]ResponseSensorsSensorsV1Response `json:"response,omitempty"` //
}
type ResponseSensorsSensorsV1Response struct {
	Name               string                               `json:"name,omitempty"`               // The sensor device name
	Status             string                               `json:"status,omitempty"`             // Status of sensor device (REACHABLE, UNREACHABLE, DELETED, RUNNING, IDLE, UCLAIMED)
	RadioMacAddress    string                               `json:"radioMacAddress,omitempty"`    // Sensor device's radio MAC address
	EthernetMacAddress string                               `json:"ethernetMacAddress,omitempty"` // Sensor device's ethernet MAC address
	Location           string                               `json:"location,omitempty"`           // Site name in hierarchy form
	BackhaulType       string                               `json:"backhaulType,omitempty"`       // Backhall type: WIRED, WIRELESS
	SerialNumber       string                               `json:"serialNumber,omitempty"`       // Serial number
	IPAddress          string                               `json:"ipAddress,omitempty"`          // IP Address
	Version            string                               `json:"version,omitempty"`            // Sensor version
	LastSeen           *int                                 `json:"lastSeen,omitempty"`           // Last seen timestamp
	Type               string                               `json:"type,omitempty"`               // Type
	SSH                *ResponseSensorsSensorsV1ResponseSSH `json:"ssh,omitempty"`                //
	Led                *bool                                `json:"led,omitempty"`                // Is LED Enabled
}
type ResponseSensorsSensorsV1ResponseSSH struct {
	SSHState       string `json:"sshState,omitempty"`       // SSH state
	SSHUserName    string `json:"sshUserName,omitempty"`    // SSH user name
	SSHPassword    string `json:"sshPassword,omitempty"`    // SSH password
	EnablePassword string `json:"enablePassword,omitempty"` // Enable password
}
type ResponseSensorsDuplicateSensorTestTemplateV1 struct {
	Version  string                                                `json:"version,omitempty"`  // Version
	Response *ResponseSensorsDuplicateSensorTestTemplateV1Response `json:"response,omitempty"` //
}
type ResponseSensorsDuplicateSensorTestTemplateV1Response struct {
	Name                   string                                                                  `json:"name,omitempty"`                   // The sensor test template name
	TypeID                 string                                                                  `json:"_id,omitempty"`                    // The sensor test template unique identifier
	Version                *int                                                                    `json:"version,omitempty"`                // The sensor test template version (must be 2)
	ModelVersion           *int                                                                    `json:"modelVersion,omitempty"`           // Test template object model version (must be 2)
	StartTime              *int                                                                    `json:"startTime,omitempty"`              // Start time
	LastModifiedTime       *int                                                                    `json:"lastModifiedTime,omitempty"`       // Last modify time
	NumAssociatedSensor    *int                                                                    `json:"numAssociatedSensor,omitempty"`    // Number of associated sensor
	Location               string                                                                  `json:"location,omitempty"`               // Location string
	SiteHierarchy          string                                                                  `json:"siteHierarchy,omitempty"`          // Site hierarchy
	Status                 string                                                                  `json:"status,omitempty"`                 // Status of the test (RUNNING, NOTRUNNING)
	Connection             string                                                                  `json:"connection,omitempty"`             // connection type of test: WIRED, WIRELESS, BOTH
	ActionInProgress       string                                                                  `json:"actionInProgress,omitempty"`       // Indication of inprogress action
	Frequency              *ResponseSensorsDuplicateSensorTestTemplateV1ResponseFrequency          `json:"frequency,omitempty"`              //
	RssiThreshold          *int                                                                    `json:"rssiThreshold,omitempty"`          // RSSI threshold
	NumNeighborApThreshold *int                                                                    `json:"numNeighborAPThreshold,omitempty"` // Number of neighboring AP threshold
	ScheduleInDays         *int                                                                    `json:"scheduleInDays,omitempty"`         // Bit-wise value of scheduled test days
	WLANs                  []string                                                                `json:"wlans,omitempty"`                  // WLANs list
	SSIDs                  *[]ResponseSensorsDuplicateSensorTestTemplateV1ResponseSSIDs            `json:"ssids,omitempty"`                  //
	Profiles               *[]ResponseSensorsDuplicateSensorTestTemplateV1ResponseProfiles         `json:"profiles,omitempty"`               //
	TestScheduleMode       string                                                                  `json:"testScheduleMode,omitempty"`       // Test schedule mode (ONDEMAND, DEDICATED, SCHEDULED, CONTINUOUS, RUNNOW)
	ShowWlcUpgradeBanner   *bool                                                                   `json:"showWlcUpgradeBanner,omitempty"`   // Show WLC upgrade banner
	RadioAsSensorRemoved   *bool                                                                   `json:"radioAsSensorRemoved,omitempty"`   // Radio as sensor removed
	EncryptionMode         string                                                                  `json:"encryptionMode,omitempty"`         // Encryption mode
	RunNow                 string                                                                  `json:"runNow,omitempty"`                 // Run now (YES, NO)
	LocationInfoList       *[]ResponseSensorsDuplicateSensorTestTemplateV1ResponseLocationInfoList `json:"locationInfoList,omitempty"`       //
	Sensors                *[]ResponseSensorsDuplicateSensorTestTemplateV1ResponseSensors          `json:"sensors,omitempty"`                //
	ApCoverage             *[]ResponseSensorsDuplicateSensorTestTemplateV1ResponseApCoverage       `json:"apCoverage,omitempty"`             //
}
type ResponseSensorsDuplicateSensorTestTemplateV1ResponseFrequency struct {
	Value *int   `json:"value,omitempty"` // Value of the unit
	Unit  string `json:"unit,omitempty"`  // Unit of the time value (NANOSECONDS, MICROSECONDS, MILLISECONDS, SECONDS, MINUTES, HOURS, DAYS)
}
type ResponseSensorsDuplicateSensorTestTemplateV1ResponseSSIDs struct {
	Bands                     string                                                                        `json:"bands,omitempty"`                     // WIFI bands: 2.4GHz or 5GHz
	SSID                      string                                                                        `json:"ssid,omitempty"`                      // The SSID string
	ProfileName               string                                                                        `json:"profileName,omitempty"`               // The SSID profile name string
	NumAps                    *int                                                                          `json:"numAps,omitempty"`                    // Number of APs in the test
	NumSensors                *int                                                                          `json:"numSensors,omitempty"`                // Number of Sensors in the test
	Layer3WebAuthsecurity     string                                                                        `json:"layer3webAuthsecurity,omitempty"`     // Layer 3 WEB Auth security
	Layer3WebAuthuserName     string                                                                        `json:"layer3webAuthuserName,omitempty"`     // Layer 3 WEB Auth user name
	Layer3WebAuthpassword     string                                                                        `json:"layer3webAuthpassword,omitempty"`     // Layer 3 WEB Auth password
	Layer3WebAuthEmailAddress string                                                                        `json:"layer3webAuthEmailAddress,omitempty"` // Layer 3 WEB Auth email address
	ThirdParty                *ResponseSensorsDuplicateSensorTestTemplateV1ResponseSSIDsThirdParty          `json:"thirdParty,omitempty"`                //
	ID                        *int                                                                          `json:"id,omitempty"`                        // Identification number
	WLANID                    *int                                                                          `json:"wlanId,omitempty"`                    // WLAN ID
	Wlc                       string                                                                        `json:"wlc,omitempty"`                       // WLC IP addres
	ValidFrom                 *int                                                                          `json:"validFrom,omitempty"`                 // Valid From UTC timestamp
	ValidTo                   *int                                                                          `json:"validTo,omitempty"`                   // Valid To UTC timestamp
	Status                    string                                                                        `json:"status,omitempty"`                    // WLAN status: ENABLED or DISABLED
	ProxyServer               string                                                                        `json:"proxyServer,omitempty"`               // Proxy server for onboarding SSID
	ProxyPort                 string                                                                        `json:"proxyPort,omitempty"`                 // Proxy server port
	ProxyUserName             string                                                                        `json:"proxyUserName,omitempty"`             // Proxy server user name
	ProxyPassword             string                                                                        `json:"proxyPassword,omitempty"`             // Proxy server password
	AuthType                  string                                                                        `json:"authType,omitempty"`                  // Authentication type: OPEN, WPA2_PSK, WPA2_EaP, WEB_AUTH, MAB, DOT1X, OTHER
	Psk                       string                                                                        `json:"psk,omitempty"`                       // Password of SSID when passwordType is ASCII
	Username                  string                                                                        `json:"username,omitempty"`                  // User name string for onboarding SSID
	Password                  string                                                                        `json:"password,omitempty"`                  // Password string for onboarding SSID
	PasswordType              string                                                                        `json:"passwordType,omitempty"`              // SSID password type: ASCII or HEX
	EapMethod                 string                                                                        `json:"eapMethod,omitempty"`                 // WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC
	Scep                      *bool                                                                         `json:"scep,omitempty"`                      // Secure certificate enrollment protocol: true or false or null for not applicable
	AuthProtocol              string                                                                        `json:"authProtocol,omitempty"`              // Auth protocol
	Certfilename              string                                                                        `json:"certfilename,omitempty"`              // Auth certificate file name
	Certxferprotocol          string                                                                        `json:"certxferprotocol,omitempty"`          // Certificate transfering protocol: HTTP or HTTPS
	Certstatus                string                                                                        `json:"certstatus,omitempty"`                // Certificate status: INACTIVE or ACTIVE
	Certpassphrase            string                                                                        `json:"certpassphrase,omitempty"`            // Certificate password phrase
	Certdownloadurl           string                                                                        `json:"certdownloadurl,omitempty"`           // Certificate download URL
	ExtWebAuthVirtualIP       string                                                                        `json:"extWebAuthVirtualIp,omitempty"`       // External WEB Auth virtual IP
	ExtWebAuth                *bool                                                                         `json:"extWebAuth,omitempty"`                // Indication of using external WEB Auth
	WhiteList                 *bool                                                                         `json:"whiteList,omitempty"`                 // Indication of being on allowed list
	ExtWebAuthPortal          string                                                                        `json:"extWebAuthPortal,omitempty"`          // External authentication portal
	ExtWebAuthAccessURL       string                                                                        `json:"extWebAuthAccessUrl,omitempty"`       // External WEB Auth access URL
	ExtWebAuthHTMLTag         *[]ResponseSensorsDuplicateSensorTestTemplateV1ResponseSSIDsExtWebAuthHTMLTag `json:"extWebAuthHtmlTag,omitempty"`         //
	QosPolicy                 string                                                                        `json:"qosPolicy,omitempty"`                 // QoS policy: PlATINUM, GOLD, SILVER, BRONZE
	Tests                     *[]ResponseSensorsDuplicateSensorTestTemplateV1ResponseSSIDsTests             `json:"tests,omitempty"`                     //
}
type ResponseSensorsDuplicateSensorTestTemplateV1ResponseSSIDsThirdParty struct {
	Selected *bool `json:"selected,omitempty"` // true: the SSID is third party
}
type ResponseSensorsDuplicateSensorTestTemplateV1ResponseSSIDsExtWebAuthHTMLTag struct {
	Label string `json:"label,omitempty"` // Label
	Tag   string `json:"tag,omitempty"`   // Tag
	Value string `json:"value,omitempty"` // Value
}
type ResponseSensorsDuplicateSensorTestTemplateV1ResponseSSIDsTests struct {
	Name   string                                                                  `json:"name,omitempty"`   // Name of the test
	Config *[]ResponseSensorsDuplicateSensorTestTemplateV1ResponseSSIDsTestsConfig `json:"config,omitempty"` //
}
type ResponseSensorsDuplicateSensorTestTemplateV1ResponseSSIDsTestsConfig struct {
	Domains        []string `json:"domains,omitempty"`        // DNS domain name
	Server         string   `json:"server,omitempty"`         // Ping, file transfer, mail, radius, ssh, or telnet server
	UserName       string   `json:"userName,omitempty"`       // User name
	Password       string   `json:"password,omitempty"`       // Password
	URL            string   `json:"url,omitempty"`            // URL
	Port           *int     `json:"port,omitempty"`           // Radius or WEB server port
	Protocol       string   `json:"protocol,omitempty"`       // Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)
	Servers        []string `json:"servers,omitempty"`        // IPerf server list
	Direction      string   `json:"direction,omitempty"`      // IPerf direction (UPLOAD, DOWNLOAD, BOTH)
	StartPort      *int     `json:"startPort,omitempty"`      // IPerf start port
	EndPort        *int     `json:"endPort,omitempty"`        // IPerf end port
	UDPBandwidth   *int     `json:"udpBandwidth,omitempty"`   // IPerf UDP bandwidth
	ProbeType      string   `json:"probeType,omitempty"`      // Probe type
	NumPackets     *int     `json:"numPackets,omitempty"`     // Number of packets
	PathToDownload string   `json:"pathToDownload,omitempty"` // File path for file transfer
	TransferType   string   `json:"transferType,omitempty"`   // File transfer type (UPLOAD, DOWNLOAD, BOTH)
	SharedSecret   string   `json:"sharedSecret,omitempty"`   // Shared secret
	NdtServer      string   `json:"ndtServer,omitempty"`      // NDT server
	NdtServerPort  string   `json:"ndtServerPort,omitempty"`  // NDT server port
	NdtServerPath  string   `json:"ndtServerPath,omitempty"`  // NDT server path
	UplinkTest     *bool    `json:"uplinkTest,omitempty"`     // Uplink test
	DownlinkTest   *bool    `json:"downlinkTest,omitempty"`   // Downlink test
	ProxyServer    string   `json:"proxyServer,omitempty"`    // Proxy server
	ProxyPort      string   `json:"proxyPort,omitempty"`      // Proxy port
	ProxyUserName  string   `json:"proxyUserName,omitempty"`  // Proxy user name
	ProxyPassword  string   `json:"proxyPassword,omitempty"`  // Proxy password
	UserNamePrompt string   `json:"userNamePrompt,omitempty"` // User name prompt
	PasswordPrompt string   `json:"passwordPrompt,omitempty"` // Password prompt
	ExitCommand    string   `json:"exitCommand,omitempty"`    // Exit command
	FinalPrompt    string   `json:"finalPrompt,omitempty"`    // Final prompt
}
type ResponseSensorsDuplicateSensorTestTemplateV1ResponseProfiles struct {
	AuthType            string                                                                           `json:"authType,omitempty"`            // Authentication type: OPEN, WPA2_PSK, WPA2_EaP, WEB_AUTH, MAB, DOT1X, OTHER
	Psk                 string                                                                           `json:"psk,omitempty"`                 // Password of SSID when passwordType is ASCII
	Username            string                                                                           `json:"username,omitempty"`            // User name string for onboarding SSID
	Password            string                                                                           `json:"password,omitempty"`            // Password string for onboarding SSID
	PasswordType        string                                                                           `json:"passwordType,omitempty"`        // SSID password type: ASCII or HEX
	EapMethod           string                                                                           `json:"eapMethod,omitempty"`           // WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC
	Scep                *bool                                                                            `json:"scep,omitempty"`                // Secure certificate enrollment protocol: true or false or null for not applicable
	AuthProtocol        string                                                                           `json:"authProtocol,omitempty"`        // Auth protocol
	Certfilename        string                                                                           `json:"certfilename,omitempty"`        // Auth certificate file name
	Certxferprotocol    string                                                                           `json:"certxferprotocol,omitempty"`    // Certificate transfering protocol: HTTP or HTTPS
	Certstatus          string                                                                           `json:"certstatus,omitempty"`          // Certificate status: INACTIVE or ACTIVE
	Certpassphrase      string                                                                           `json:"certpassphrase,omitempty"`      // Certificate password phrase
	Certdownloadurl     string                                                                           `json:"certdownloadurl,omitempty"`     // Certificate download URL
	ExtWebAuthVirtualIP string                                                                           `json:"extWebAuthVirtualIp,omitempty"` // External WEB Auth virtual IP
	ExtWebAuth          *bool                                                                            `json:"extWebAuth,omitempty"`          // Indication of using external WEB Auth
	WhiteList           *bool                                                                            `json:"whiteList,omitempty"`           // Indication of being on allowed list
	ExtWebAuthPortal    string                                                                           `json:"extWebAuthPortal,omitempty"`    // External authentication portal
	ExtWebAuthAccessURL string                                                                           `json:"extWebAuthAccessUrl,omitempty"` // External WEB Auth access URL
	ExtWebAuthHTMLTag   *[]ResponseSensorsDuplicateSensorTestTemplateV1ResponseProfilesExtWebAuthHTMLTag `json:"extWebAuthHtmlTag,omitempty"`   //
	QosPolicy           string                                                                           `json:"qosPolicy,omitempty"`           // QoS policy: PlATINUM, GOLD, SILVER, BRONZE
	Tests               *[]ResponseSensorsDuplicateSensorTestTemplateV1ResponseProfilesTests             `json:"tests,omitempty"`               //
	ProfileName         string                                                                           `json:"profileName,omitempty"`         // Profile name
	DeviceType          string                                                                           `json:"deviceType,omitempty"`          // Device Type
	VLAN                string                                                                           `json:"vlan,omitempty"`                // VLAN
	LocationVLANList    *[]ResponseSensorsDuplicateSensorTestTemplateV1ResponseProfilesLocationVLANList  `json:"locationVlanList,omitempty"`    //
}
type ResponseSensorsDuplicateSensorTestTemplateV1ResponseProfilesExtWebAuthHTMLTag struct {
	Label string `json:"label,omitempty"` // Label
	Tag   string `json:"tag,omitempty"`   // Tag
	Value string `json:"value,omitempty"` // Value
}
type ResponseSensorsDuplicateSensorTestTemplateV1ResponseProfilesTests struct {
	Name   string                                                                     `json:"name,omitempty"`   // Name of the test
	Config *[]ResponseSensorsDuplicateSensorTestTemplateV1ResponseProfilesTestsConfig `json:"config,omitempty"` //
}
type ResponseSensorsDuplicateSensorTestTemplateV1ResponseProfilesTestsConfig struct {
	Domains        []string `json:"domains,omitempty"`        // DNS domain name
	Server         string   `json:"server,omitempty"`         // Ping, file transfer, mail, radius, ssh, or telnet server
	UserName       string   `json:"userName,omitempty"`       // User name
	Password       string   `json:"password,omitempty"`       // Password
	URL            string   `json:"url,omitempty"`            // URL
	Port           *int     `json:"port,omitempty"`           // Radius or WEB server port
	Protocol       string   `json:"protocol,omitempty"`       // Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)
	Servers        []string `json:"servers,omitempty"`        // IPerf server list
	Direction      string   `json:"direction,omitempty"`      // IPerf direction (UPLOAD, DOWNLOAD, BOTH)
	StartPort      *int     `json:"startPort,omitempty"`      // IPerf start port
	EndPort        *int     `json:"endPort,omitempty"`        // IPerf end port
	UDPBandwidth   *int     `json:"udpBandwidth,omitempty"`   // IPerf UDP bandwidth
	ProbeType      string   `json:"probeType,omitempty"`      // Probe type
	NumPackets     *int     `json:"numPackets,omitempty"`     // Number of packets
	PathToDownload string   `json:"pathToDownload,omitempty"` // File path for file transfer
	TransferType   string   `json:"transferType,omitempty"`   // File transfer type (UPLOAD, DOWNLOAD, BOTH)
	SharedSecret   string   `json:"sharedSecret,omitempty"`   // Shared secret
	NdtServer      string   `json:"ndtServer,omitempty"`      // NDT server
	NdtServerPort  string   `json:"ndtServerPort,omitempty"`  // NDT server port
	NdtServerPath  string   `json:"ndtServerPath,omitempty"`  // NDT server path
	UplinkTest     *bool    `json:"uplinkTest,omitempty"`     // Uplink test
	DownlinkTest   *bool    `json:"downlinkTest,omitempty"`   // Downlink test
	ProxyServer    string   `json:"proxyServer,omitempty"`    // Proxy server
	ProxyPort      string   `json:"proxyPort,omitempty"`      // Proxy port
	ProxyUserName  string   `json:"proxyUserName,omitempty"`  // Proxy user name
	ProxyPassword  string   `json:"proxyPassword,omitempty"`  // Proxy password
	UserNamePrompt string   `json:"userNamePrompt,omitempty"` // User name prompt
	PasswordPrompt string   `json:"passwordPrompt,omitempty"` // Password prompt
	ExitCommand    string   `json:"exitCommand,omitempty"`    // Exit command
	FinalPrompt    string   `json:"finalPrompt,omitempty"`    // Final prompt
}
type ResponseSensorsDuplicateSensorTestTemplateV1ResponseProfilesLocationVLANList struct {
	LocationID string   `json:"locationId,omitempty"` // Site UUID
	VLANs      []string `json:"vlans,omitempty"`      // Array of VLANs
}
type ResponseSensorsDuplicateSensorTestTemplateV1ResponseLocationInfoList struct {
	LocationID           string   `json:"locationId,omitempty"`           // Site UUID
	LocationType         string   `json:"locationType,omitempty"`         // Site type
	AllSensors           *bool    `json:"allSensors,omitempty"`           // Use all sensors in the site for test
	SiteHierarchy        string   `json:"siteHierarchy,omitempty"`        // Site name hierarhy
	MacAddressList       []string `json:"macAddressList,omitempty"`       // MAC addresses
	ManagementVLAN       string   `json:"managementVlan,omitempty"`       // Management VLAN
	CustomManagementVLAN *bool    `json:"customManagementVlan,omitempty"` // Custom Management VLAN
}
type ResponseSensorsDuplicateSensorTestTemplateV1ResponseSensors struct {
	Name                    string                                                                       `json:"name,omitempty"`                    // Sensor name
	MacAddress              string                                                                       `json:"macAddress,omitempty"`              // MAC address
	SwitchMac               string                                                                       `json:"switchMac,omitempty"`               // Switch MAC address
	SwitchUUID              string                                                                       `json:"switchUuid,omitempty"`              // Switch device UUID
	SwitchSerialNumber      string                                                                       `json:"switchSerialNumber,omitempty"`      // Switch serial number
	MarkedForUninstall      *bool                                                                        `json:"markedForUninstall,omitempty"`      // Is marked for uninstall
	IPAddress               string                                                                       `json:"ipAddress,omitempty"`               // IP address
	HostName                string                                                                       `json:"hostName,omitempty"`                // Host name
	WiredApplicationStatus  string                                                                       `json:"wiredApplicationStatus,omitempty"`  // Wired application status
	WiredApplicationMessage string                                                                       `json:"wiredApplicationMessage,omitempty"` // Wired application message
	Assigned                *bool                                                                        `json:"assigned,omitempty"`                // Is assigned
	Status                  string                                                                       `json:"status,omitempty"`                  // Sensor device status: UP, DOWN, REBOOT
	XorSensor               *bool                                                                        `json:"xorSensor,omitempty"`               // Is XOR sensor
	TargetAPs               []string                                                                     `json:"targetAPs,omitempty"`               // Array of target APs
	RunNow                  string                                                                       `json:"runNow,omitempty"`                  // Run now: YES, NO
	LocationID              string                                                                       `json:"locationId,omitempty"`              // Site UUID
	AllSensorAddition       *bool                                                                        `json:"allSensorAddition,omitempty"`       // Is all sensor addition
	ConfigUpdated           string                                                                       `json:"configUpdated,omitempty"`           // Configuration updated: YES, NO
	SensorType              string                                                                       `json:"sensorType,omitempty"`              // Sensor type
	TestMacAddresses        *ResponseSensorsDuplicateSensorTestTemplateV1ResponseSensorsTestMacAddresses `json:"testMacAddresses,omitempty"`        // A string-string test MAC address
	ID                      string                                                                       `json:"id,omitempty"`                      // Sensor ID
	ServicePolicy           string                                                                       `json:"servicePolicy,omitempty"`           // Service policy
	IPerfInfo               *ResponseSensorsDuplicateSensorTestTemplateV1ResponseSensorsIPerfInfo        `json:"iPerfInfo,omitempty"`               // A string-stringList iPerf information
}
type ResponseSensorsDuplicateSensorTestTemplateV1ResponseSensorsTestMacAddresses interface{}
type ResponseSensorsDuplicateSensorTestTemplateV1ResponseSensorsIPerfInfo interface{}
type ResponseSensorsDuplicateSensorTestTemplateV1ResponseApCoverage struct {
	Bands             string `json:"bands,omitempty"`             // The WIFI bands
	NumberOfApsToTest *int   `json:"numberOfApsToTest,omitempty"` // Number of APs to test
	RssiThreshold     *int   `json:"rssiThreshold,omitempty"`     // RSSI threshold
}
type RequestSensorsEditSensorTestTemplateV1 struct {
	TemplateName           string                                                    `json:"templateName,omitempty"`           // The test template name that is to be edited
	Name                   string                                                    `json:"name,omitempty"`                   // The sensor test template name, which is the same as in 'templateName'
	TypeID                 string                                                    `json:"_id,omitempty"`                    // The sensor test template unique identifier, generated at test creation time
	Version                *int                                                      `json:"version,omitempty"`                // The sensor test template version (must be 2)
	ModelVersion           *int                                                      `json:"modelVersion,omitempty"`           // Test template object model version (must be 2)
	StartTime              *int                                                      `json:"startTime,omitempty"`              // Start time
	LastModifiedTime       *int                                                      `json:"lastModifiedTime,omitempty"`       // Last modify time
	NumAssociatedSensor    *int                                                      `json:"numAssociatedSensor,omitempty"`    // Number of associated sensor
	Location               string                                                    `json:"location,omitempty"`               // Location string
	SiteHierarchy          string                                                    `json:"siteHierarchy,omitempty"`          // Site hierarchy
	Status                 string                                                    `json:"status,omitempty"`                 // Status of the test (RUNNING, NOTRUNNING)
	Connection             string                                                    `json:"connection,omitempty"`             // connection type of test: WIRED, WIRELESS, BOTH
	ActionInProgress       string                                                    `json:"actionInProgress,omitempty"`       // Indication of inprogress action
	Frequency              *RequestSensorsEditSensorTestTemplateV1Frequency          `json:"frequency,omitempty"`              //
	RssiThreshold          *int                                                      `json:"rssiThreshold,omitempty"`          // RSSI threshold
	NumNeighborApThreshold *int                                                      `json:"numNeighborAPThreshold,omitempty"` // Number of neighboring AP threshold
	ScheduleInDays         *int                                                      `json:"scheduleInDays,omitempty"`         // Bit-wise value of scheduled test days
	WLANs                  []string                                                  `json:"wlans,omitempty"`                  // WLANs list
	SSIDs                  *[]RequestSensorsEditSensorTestTemplateV1SSIDs            `json:"ssids,omitempty"`                  //
	Profiles               *[]RequestSensorsEditSensorTestTemplateV1Profiles         `json:"profiles,omitempty"`               //
	TestScheduleMode       string                                                    `json:"testScheduleMode,omitempty"`       // Test schedule mode (ONDEMAND, DEDICATED, SCHEDULED, CONTINUOUS, RUNNOW)
	ShowWlcUpgradeBanner   *bool                                                     `json:"showWlcUpgradeBanner,omitempty"`   // Show WLC upgrade banner
	RadioAsSensorRemoved   *bool                                                     `json:"radioAsSensorRemoved,omitempty"`   // Radio as sensor removed
	EncryptionMode         string                                                    `json:"encryptionMode,omitempty"`         // Encryption mode
	RunNow                 string                                                    `json:"runNow,omitempty"`                 // Run now (YES, NO)
	LocationInfoList       *[]RequestSensorsEditSensorTestTemplateV1LocationInfoList `json:"locationInfoList,omitempty"`       //
	Sensors                *[]RequestSensorsEditSensorTestTemplateV1Sensors          `json:"sensors,omitempty"`                //
	ApCoverage             *[]RequestSensorsEditSensorTestTemplateV1ApCoverage       `json:"apCoverage,omitempty"`             //
}
type RequestSensorsEditSensorTestTemplateV1Frequency struct {
	Value *int   `json:"value,omitempty"` // Value of the unit
	Unit  string `json:"unit,omitempty"`  // Unit of the time value (NANOSECONDS, MICROSECONDS, MILLISECONDS, SECONDS, MINUTES, HOURS, DAYS)
}
type RequestSensorsEditSensorTestTemplateV1SSIDs struct {
	Bands                     string                                                          `json:"bands,omitempty"`                     // WIFI bands: 2.4GHz or 5GHz
	SSID                      string                                                          `json:"ssid,omitempty"`                      // The SSID string
	ProfileName               string                                                          `json:"profileName,omitempty"`               // The SSID profile name string
	NumAps                    *int                                                            `json:"numAps,omitempty"`                    // Number of APs in the test
	NumSensors                *int                                                            `json:"numSensors,omitempty"`                // Number of Sensors in the test
	Layer3WebAuthsecurity     string                                                          `json:"layer3webAuthsecurity,omitempty"`     // Layer 3 WEB Auth security
	Layer3WebAuthuserName     string                                                          `json:"layer3webAuthuserName,omitempty"`     // Layer 3 WEB Auth user name
	Layer3WebAuthpassword     string                                                          `json:"layer3webAuthpassword,omitempty"`     // Layer 3 WEB Auth password
	Layer3WebAuthEmailAddress string                                                          `json:"layer3webAuthEmailAddress,omitempty"` // Layer 3 WEB Auth email address
	ThirdParty                *RequestSensorsEditSensorTestTemplateV1SSIDsThirdParty          `json:"thirdParty,omitempty"`                //
	ID                        *int                                                            `json:"id,omitempty"`                        // Identification number
	WLANID                    *int                                                            `json:"wlanId,omitempty"`                    // WLAN ID
	Wlc                       string                                                          `json:"wlc,omitempty"`                       // WLC IP addres
	ValidFrom                 *int                                                            `json:"validFrom,omitempty"`                 // Valid From UTC timestamp
	ValidTo                   *int                                                            `json:"validTo,omitempty"`                   // Valid To UTC timestamp
	Status                    string                                                          `json:"status,omitempty"`                    // WLAN status: ENABLED or DISABLED
	ProxyServer               string                                                          `json:"proxyServer,omitempty"`               // Proxy server for onboarding SSID
	ProxyPort                 string                                                          `json:"proxyPort,omitempty"`                 // Proxy server port
	ProxyUserName             string                                                          `json:"proxyUserName,omitempty"`             // Proxy server user name
	ProxyPassword             string                                                          `json:"proxyPassword,omitempty"`             // Proxy server password
	AuthType                  string                                                          `json:"authType,omitempty"`                  // Authentication type: OPEN, WPA2_PSK, WPA2_EAP, WEB_AUTH, MAB, DOT1X, OTHER
	Psk                       string                                                          `json:"psk,omitempty"`                       // Password of SSID when passwordType is ASCII
	Username                  string                                                          `json:"username,omitempty"`                  // User name string for onboarding SSID
	Password                  string                                                          `json:"password,omitempty"`                  // Password string for onboarding SSID
	PasswordType              string                                                          `json:"passwordType,omitempty"`              // SSID password type: ASCII or HEX
	EapMethod                 string                                                          `json:"eapMethod,omitempty"`                 // WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC
	Scep                      *bool                                                           `json:"scep,omitempty"`                      // Secure certificate enrollment protocol: true or false or null for not applicable
	AuthProtocol              string                                                          `json:"authProtocol,omitempty"`              // Auth protocol
	Certfilename              string                                                          `json:"certfilename,omitempty"`              // Auth certificate file name
	Certxferprotocol          string                                                          `json:"certxferprotocol,omitempty"`          // Certificate transfering protocol: HTTP or HTTPS
	Certstatus                string                                                          `json:"certstatus,omitempty"`                // Certificate status: INACTIVE or ACTIVE
	Certpassphrase            string                                                          `json:"certpassphrase,omitempty"`            // Certificate password phrase
	Certdownloadurl           string                                                          `json:"certdownloadurl,omitempty"`           // Certificate download URL
	ExtWebAuthVirtualIP       string                                                          `json:"extWebAuthVirtualIp,omitempty"`       // External WEB Auth virtual IP
	ExtWebAuth                *bool                                                           `json:"extWebAuth,omitempty"`                // Indication of using external WEB Auth
	WhiteList                 *bool                                                           `json:"whiteList,omitempty"`                 // Indication of being on allowed list
	ExtWebAuthPortal          string                                                          `json:"extWebAuthPortal,omitempty"`          // External authentication portal
	ExtWebAuthAccessURL       string                                                          `json:"extWebAuthAccessUrl,omitempty"`       // External WEB Auth access URL
	ExtWebAuthHTMLTag         *[]RequestSensorsEditSensorTestTemplateV1SSIDsExtWebAuthHTMLTag `json:"extWebAuthHtmlTag,omitempty"`         //
	QosPolicy                 string                                                          `json:"qosPolicy,omitempty"`                 // QoS policy: PlATINUM, GOLD, SILVER, BRONZE
	Tests                     *[]RequestSensorsEditSensorTestTemplateV1SSIDsTests             `json:"tests,omitempty"`                     //
}
type RequestSensorsEditSensorTestTemplateV1SSIDsThirdParty struct {
	Selected *bool `json:"selected,omitempty"` // true: the SSID is third party
}
type RequestSensorsEditSensorTestTemplateV1SSIDsExtWebAuthHTMLTag struct {
	Label string `json:"label,omitempty"` // Label
	Tag   string `json:"tag,omitempty"`   // Tag
	Value string `json:"value,omitempty"` // Value
}
type RequestSensorsEditSensorTestTemplateV1SSIDsTests struct {
	Name   string                                                    `json:"name,omitempty"`   // Name of the test
	Config *[]RequestSensorsEditSensorTestTemplateV1SSIDsTestsConfig `json:"config,omitempty"` //
}
type RequestSensorsEditSensorTestTemplateV1SSIDsTestsConfig struct {
	Domains        []string `json:"domains,omitempty"`        // DNS domain name
	Server         string   `json:"server,omitempty"`         // Ping, file transfer, mail, radius, ssh, or telnet server
	UserName       string   `json:"userName,omitempty"`       // User name
	Password       string   `json:"password,omitempty"`       // Password
	URL            string   `json:"url,omitempty"`            // URL
	Port           *int     `json:"port,omitempty"`           // Radius or WEB server port
	Protocol       string   `json:"protocol,omitempty"`       // Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)
	Servers        []string `json:"servers,omitempty"`        // IPerf server list
	Direction      string   `json:"direction,omitempty"`      // IPerf direction (UPLOAD, DOWNLOAD, BOTH)
	StartPort      *int     `json:"startPort,omitempty"`      // IPerf start port
	EndPort        *int     `json:"endPort,omitempty"`        // IPerf end port
	UDPBandwidth   *int     `json:"udpBandwidth,omitempty"`   // IPerf UDP bandwidth
	ProbeType      string   `json:"probeType,omitempty"`      // Probe type
	NumPackets     *int     `json:"numPackets,omitempty"`     // Number of packets
	PathToDownload string   `json:"pathToDownload,omitempty"` // File path for file transfer
	TransferType   string   `json:"transferType,omitempty"`   // File transfer type (UPLOAD, DOWNLOAD, BOTH)
	SharedSecret   string   `json:"sharedSecret,omitempty"`   // Shared secret
	NdtServer      string   `json:"ndtServer,omitempty"`      // NDT server
	NdtServerPort  string   `json:"ndtServerPort,omitempty"`  // NDT server port
	NdtServerPath  string   `json:"ndtServerPath,omitempty"`  // NDT server path
	UplinkTest     *bool    `json:"uplinkTest,omitempty"`     // Uplink test
	DownlinkTest   *bool    `json:"downlinkTest,omitempty"`   // Downlink test
	ProxyServer    string   `json:"proxyServer,omitempty"`    // Proxy server
	ProxyPort      string   `json:"proxyPort,omitempty"`      // Proxy port
	ProxyUserName  string   `json:"proxyUserName,omitempty"`  // Proxy user name
	ProxyPassword  string   `json:"proxyPassword,omitempty"`  // Proxy password
	UserNamePrompt string   `json:"userNamePrompt,omitempty"` // User name prompt
	PasswordPrompt string   `json:"passwordPrompt,omitempty"` // Password prompt
	ExitCommand    string   `json:"exitCommand,omitempty"`    // Exit command
	FinalPrompt    string   `json:"finalPrompt,omitempty"`    // Final prompt
}
type RequestSensorsEditSensorTestTemplateV1Profiles struct {
	AuthType            string                                                             `json:"authType,omitempty"`            // Authentication type: OPEN, WPA2_PSK, WPA2_EAP, WEB_AUTH, MAB, DOT1X, OTHER
	Psk                 string                                                             `json:"psk,omitempty"`                 // Password of SSID when passwordType is ASCII
	Username            string                                                             `json:"username,omitempty"`            // User name string for onboarding SSID
	Password            string                                                             `json:"password,omitempty"`            // Password string for onboarding SSID
	PasswordType        string                                                             `json:"passwordType,omitempty"`        // SSID password type: ASCII or HEX
	EapMethod           string                                                             `json:"eapMethod,omitempty"`           // WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC
	Scep                *bool                                                              `json:"scep,omitempty"`                // Secure certificate enrollment protocol: true or false or null for not applicable
	AuthProtocol        string                                                             `json:"authProtocol,omitempty"`        // Auth protocol
	Certfilename        string                                                             `json:"certfilename,omitempty"`        // Auth certificate file name
	Certxferprotocol    string                                                             `json:"certxferprotocol,omitempty"`    // Certificate transfering protocol: HTTP or HTTPS
	Certstatus          string                                                             `json:"certstatus,omitempty"`          // Certificate status: INACTIVE or ACTIVE
	Certpassphrase      string                                                             `json:"certpassphrase,omitempty"`      // Certificate password phrase
	Certdownloadurl     string                                                             `json:"certdownloadurl,omitempty"`     // Certificate download URL
	ExtWebAuthVirtualIP string                                                             `json:"extWebAuthVirtualIp,omitempty"` // External WEB Auth virtual IP
	ExtWebAuth          *bool                                                              `json:"extWebAuth,omitempty"`          // Indication of using external WEB Auth
	WhiteList           *bool                                                              `json:"whiteList,omitempty"`           // Indication of being on allowed list
	ExtWebAuthPortal    string                                                             `json:"extWebAuthPortal,omitempty"`    // External authentication portal
	ExtWebAuthAccessURL string                                                             `json:"extWebAuthAccessUrl,omitempty"` // External WEB Auth access URL
	ExtWebAuthHTMLTag   *[]RequestSensorsEditSensorTestTemplateV1ProfilesExtWebAuthHTMLTag `json:"extWebAuthHtmlTag,omitempty"`   //
	QosPolicy           string                                                             `json:"qosPolicy,omitempty"`           // QoS policy: PlATINUM, GOLD, SILVER, BRONZE
	Tests               *[]RequestSensorsEditSensorTestTemplateV1ProfilesTests             `json:"tests,omitempty"`               //
	ProfileName         string                                                             `json:"profileName,omitempty"`         // Profile name
	DeviceType          string                                                             `json:"deviceType,omitempty"`          // Device Type
	VLAN                string                                                             `json:"vlan,omitempty"`                // VLAN
	LocationVLANList    *[]RequestSensorsEditSensorTestTemplateV1ProfilesLocationVLANList  `json:"locationVlanList,omitempty"`    //
}
type RequestSensorsEditSensorTestTemplateV1ProfilesExtWebAuthHTMLTag struct {
	Label string `json:"label,omitempty"` // Label
	Tag   string `json:"tag,omitempty"`   // Tag
	Value string `json:"value,omitempty"` // Value
}
type RequestSensorsEditSensorTestTemplateV1ProfilesTests struct {
	Name   string                                                       `json:"name,omitempty"`   // Name of the test
	Config *[]RequestSensorsEditSensorTestTemplateV1ProfilesTestsConfig `json:"config,omitempty"` //
}
type RequestSensorsEditSensorTestTemplateV1ProfilesTestsConfig struct {
	Domains        []string `json:"domains,omitempty"`        // DNS domain name
	Server         string   `json:"server,omitempty"`         // Ping, file transfer, mail, radius, ssh, or telnet server
	UserName       string   `json:"userName,omitempty"`       // User name
	Password       string   `json:"password,omitempty"`       // Password
	URL            string   `json:"url,omitempty"`            // URL
	Port           *int     `json:"port,omitempty"`           // Radius or WEB server port
	Protocol       string   `json:"protocol,omitempty"`       // Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)
	Servers        []string `json:"servers,omitempty"`        // IPerf server list
	Direction      string   `json:"direction,omitempty"`      // IPerf direction (UPLOAD, DOWNLOAD, BOTH)
	StartPort      *int     `json:"startPort,omitempty"`      // IPerf start port
	EndPort        *int     `json:"endPort,omitempty"`        // IPerf end port
	UDPBandwidth   *int     `json:"udpBandwidth,omitempty"`   // IPerf UDP bandwidth
	ProbeType      string   `json:"probeType,omitempty"`      // Probe type
	NumPackets     *int     `json:"numPackets,omitempty"`     // Number of packets
	PathToDownload string   `json:"pathToDownload,omitempty"` // File path for file transfer
	TransferType   string   `json:"transferType,omitempty"`   // File transfer type (UPLOAD, DOWNLOAD, BOTH)
	SharedSecret   string   `json:"sharedSecret,omitempty"`   // Shared secret
	NdtServer      string   `json:"ndtServer,omitempty"`      // NDT server
	NdtServerPort  string   `json:"ndtServerPort,omitempty"`  // NDT server port
	NdtServerPath  string   `json:"ndtServerPath,omitempty"`  // NDT server path
	UplinkTest     *bool    `json:"uplinkTest,omitempty"`     // Uplink test
	DownlinkTest   *bool    `json:"downlinkTest,omitempty"`   // Downlink test
	ProxyServer    string   `json:"proxyServer,omitempty"`    // Proxy server
	ProxyPort      string   `json:"proxyPort,omitempty"`      // Proxy port
	ProxyUserName  string   `json:"proxyUserName,omitempty"`  // Proxy user name
	ProxyPassword  string   `json:"proxyPassword,omitempty"`  // Proxy password
	UserNamePrompt string   `json:"userNamePrompt,omitempty"` // User name prompt
	PasswordPrompt string   `json:"passwordPrompt,omitempty"` // Password prompt
	ExitCommand    string   `json:"exitCommand,omitempty"`    // Exit command
	FinalPrompt    string   `json:"finalPrompt,omitempty"`    // Final prompt
}
type RequestSensorsEditSensorTestTemplateV1ProfilesLocationVLANList struct {
	LocationID string   `json:"locationId,omitempty"` // Site UUID
	VLANs      []string `json:"vlans,omitempty"`      // Array of VLANs
}
type RequestSensorsEditSensorTestTemplateV1LocationInfoList struct {
	LocationID           string   `json:"locationId,omitempty"`           // Site UUID
	LocationType         string   `json:"locationType,omitempty"`         // Site type
	AllSensors           *bool    `json:"allSensors,omitempty"`           // Use all sensors in the site for test
	SiteHierarchy        string   `json:"siteHierarchy,omitempty"`        // Site name hierarhy
	MacAddressList       []string `json:"macAddressList,omitempty"`       // MAC addresses
	ManagementVLAN       string   `json:"managementVlan,omitempty"`       // Management VLAN
	CustomManagementVLAN *bool    `json:"customManagementVlan,omitempty"` // Custom Management VLAN
}
type RequestSensorsEditSensorTestTemplateV1Sensors struct {
	Name                    string                                                         `json:"name,omitempty"`                    // Sensor name
	MacAddress              string                                                         `json:"macAddress,omitempty"`              // MAC address
	SwitchMac               string                                                         `json:"switchMac,omitempty"`               // Switch MAC address
	SwitchUUID              string                                                         `json:"switchUuid,omitempty"`              // Switch device UUID
	SwitchSerialNumber      string                                                         `json:"switchSerialNumber,omitempty"`      // Switch serial number
	MarkedForUninstall      *bool                                                          `json:"markedForUninstall,omitempty"`      // Is marked for uninstall
	IPAddress               string                                                         `json:"ipAddress,omitempty"`               // IP address
	HostName                string                                                         `json:"hostName,omitempty"`                // Host name
	WiredApplicationStatus  string                                                         `json:"wiredApplicationStatus,omitempty"`  // Wired application status
	WiredApplicationMessage string                                                         `json:"wiredApplicationMessage,omitempty"` // Wired application message
	Assigned                *bool                                                          `json:"assigned,omitempty"`                // Is assigned
	Status                  string                                                         `json:"status,omitempty"`                  // Sensor device status: UP, DOWN, REBOOT
	XorSensor               *bool                                                          `json:"xorSensor,omitempty"`               // Is XOR sensor
	TargetAPs               []string                                                       `json:"targetAPs,omitempty"`               // Array of target APs
	RunNow                  string                                                         `json:"runNow,omitempty"`                  // Run now: YES, NO
	LocationID              string                                                         `json:"locationId,omitempty"`              // Site UUID
	AllSensorAddition       *bool                                                          `json:"allSensorAddition,omitempty"`       // Is all sensor addition
	ConfigUpdated           string                                                         `json:"configUpdated,omitempty"`           // Configuration updated: YES, NO
	SensorType              string                                                         `json:"sensorType,omitempty"`              // Sensor type
	TestMacAddresses        *RequestSensorsEditSensorTestTemplateV1SensorsTestMacAddresses `json:"testMacAddresses,omitempty"`        // A string-string test MAC address
	ID                      string                                                         `json:"id,omitempty"`                      // Sensor ID
	ServicePolicy           string                                                         `json:"servicePolicy,omitempty"`           // Service policy
	IPerfInfo               *RequestSensorsEditSensorTestTemplateV1SensorsIPerfInfo        `json:"iPerfInfo,omitempty"`               // A string-stringList iPerf information
}
type RequestSensorsEditSensorTestTemplateV1SensorsTestMacAddresses interface{}
type RequestSensorsEditSensorTestTemplateV1SensorsIPerfInfo interface{}
type RequestSensorsEditSensorTestTemplateV1ApCoverage struct {
	Bands             string `json:"bands,omitempty"`             // The WIFI bands
	NumberOfApsToTest *int   `json:"numberOfApsToTest,omitempty"` // Number of APs to test
	RssiThreshold     *int   `json:"rssiThreshold,omitempty"`     // RSSI threshold
}
type RequestSensorsCreateSensorTestTemplateV1 struct {
	Name             string                                                      `json:"name,omitempty"`             // The sensor test template name
	Version          *int                                                        `json:"version,omitempty"`          // The sensor test template version (must be 2)
	ModelVersion     *int                                                        `json:"modelVersion,omitempty"`     // Test template object model version (must be 2)
	Connection       string                                                      `json:"connection,omitempty"`       // connection type of test: WIRED, WIRELESS, BOTH
	SSIDs            *[]RequestSensorsCreateSensorTestTemplateV1SSIDs            `json:"ssids,omitempty"`            //
	Profiles         *[]RequestSensorsCreateSensorTestTemplateV1Profiles         `json:"profiles,omitempty"`         //
	EncryptionMode   string                                                      `json:"encryptionMode,omitempty"`   // Encryption mode
	RunNow           string                                                      `json:"runNow,omitempty"`           // Run now (YES, NO)
	LocationInfoList *[]RequestSensorsCreateSensorTestTemplateV1LocationInfoList `json:"locationInfoList,omitempty"` //
	Sensors          *[]RequestSensorsCreateSensorTestTemplateV1Sensors          `json:"sensors,omitempty"`          //
	ApCoverage       *[]RequestSensorsCreateSensorTestTemplateV1ApCoverage       `json:"apCoverage,omitempty"`       //
}
type RequestSensorsCreateSensorTestTemplateV1SSIDs struct {
	Bands                     string                                                            `json:"bands,omitempty"`                     // WIFI bands: 2.4GHz or 5GHz
	SSID                      string                                                            `json:"ssid,omitempty"`                      // The SSID string
	ProfileName               string                                                            `json:"profileName,omitempty"`               // The SSID profile name string
	Layer3WebAuthsecurity     string                                                            `json:"layer3webAuthsecurity,omitempty"`     // Layer 3 WEB Auth security
	Layer3WebAuthuserName     string                                                            `json:"layer3webAuthuserName,omitempty"`     // Layer 3 WEB Auth user name
	Layer3WebAuthpassword     string                                                            `json:"layer3webAuthpassword,omitempty"`     // Layer 3 WEB Auth password
	Layer3WebAuthEmailAddress string                                                            `json:"layer3webAuthEmailAddress,omitempty"` // Layer 3 WEB Auth email address
	ThirdParty                *RequestSensorsCreateSensorTestTemplateV1SSIDsThirdParty          `json:"thirdParty,omitempty"`                //
	WLANID                    *int                                                              `json:"wlanId,omitempty"`                    // WLAN ID
	Wlc                       string                                                            `json:"wlc,omitempty"`                       // WLC IP addres
	ProxyServer               string                                                            `json:"proxyServer,omitempty"`               // Proxy server for onboarding SSID
	ProxyPort                 string                                                            `json:"proxyPort,omitempty"`                 // Proxy server port
	ProxyUserName             string                                                            `json:"proxyUserName,omitempty"`             // Proxy server user name
	ProxyPassword             string                                                            `json:"proxyPassword,omitempty"`             // Proxy server password
	AuthType                  string                                                            `json:"authType,omitempty"`                  // Authentication type: OPEN, WPA2_PSK, WPA2_EaP, WEB_AUTH, MAB, DOT1X, OTHER
	Psk                       string                                                            `json:"psk,omitempty"`                       // Password of SSID when passwordType is ASCII
	Username                  string                                                            `json:"username,omitempty"`                  // User name string for onboarding SSID
	Password                  string                                                            `json:"password,omitempty"`                  // Password string for onboarding SSID
	PasswordType              string                                                            `json:"passwordType,omitempty"`              // SSID password type: ASCII or HEX
	EapMethod                 string                                                            `json:"eapMethod,omitempty"`                 // WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC
	Scep                      *bool                                                             `json:"scep,omitempty"`                      // Secure certificate enrollment protocol: true or false or null for not applicable
	AuthProtocol              string                                                            `json:"authProtocol,omitempty"`              // Auth protocol
	Certfilename              string                                                            `json:"certfilename,omitempty"`              // Auth certificate file name
	Certxferprotocol          string                                                            `json:"certxferprotocol,omitempty"`          // Certificate transfering protocol: HTTP or HTTPS
	Certstatus                string                                                            `json:"certstatus,omitempty"`                // Certificate status: INACTIVE or ACTIVE
	Certpassphrase            string                                                            `json:"certpassphrase,omitempty"`            // Certificate password phrase
	Certdownloadurl           string                                                            `json:"certdownloadurl,omitempty"`           // Certificate download URL
	ExtWebAuthVirtualIP       string                                                            `json:"extWebAuthVirtualIp,omitempty"`       // External WEB Auth virtual IP
	ExtWebAuth                *bool                                                             `json:"extWebAuth,omitempty"`                // Indication of using external WEB Auth
	WhiteList                 *bool                                                             `json:"whiteList,omitempty"`                 // Indication of being on allowed list
	ExtWebAuthPortal          string                                                            `json:"extWebAuthPortal,omitempty"`          // External authentication portal
	ExtWebAuthAccessURL       string                                                            `json:"extWebAuthAccessUrl,omitempty"`       // External WEB Auth access URL
	ExtWebAuthHTMLTag         *[]RequestSensorsCreateSensorTestTemplateV1SSIDsExtWebAuthHTMLTag `json:"extWebAuthHtmlTag,omitempty"`         //
	QosPolicy                 string                                                            `json:"qosPolicy,omitempty"`                 // QoS policy: PlATINUM, GOLD, SILVER, BRONZE
	Tests                     *[]RequestSensorsCreateSensorTestTemplateV1SSIDsTests             `json:"tests,omitempty"`                     //
}
type RequestSensorsCreateSensorTestTemplateV1SSIDsThirdParty struct {
	Selected *bool `json:"selected,omitempty"` // true: the SSID is third party
}
type RequestSensorsCreateSensorTestTemplateV1SSIDsExtWebAuthHTMLTag struct {
	Label string `json:"label,omitempty"` // Label
	Tag   string `json:"tag,omitempty"`   // Tag
	Value string `json:"value,omitempty"` // Value
}
type RequestSensorsCreateSensorTestTemplateV1SSIDsTests struct {
	Name   string                                                      `json:"name,omitempty"`   // Name of the test
	Config *[]RequestSensorsCreateSensorTestTemplateV1SSIDsTestsConfig `json:"config,omitempty"` //
}
type RequestSensorsCreateSensorTestTemplateV1SSIDsTestsConfig struct {
	Domains        []string `json:"domains,omitempty"`        // DNS domain name
	Server         string   `json:"server,omitempty"`         // Ping, file transfer, mail, radius, ssh, or telnet server
	UserName       string   `json:"userName,omitempty"`       // User name
	Password       string   `json:"password,omitempty"`       // Password
	URL            string   `json:"url,omitempty"`            // URL
	Port           *int     `json:"port,omitempty"`           // Radius or WEB server port
	Protocol       string   `json:"protocol,omitempty"`       // Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)
	Servers        []string `json:"servers,omitempty"`        // IPerf server list
	Direction      string   `json:"direction,omitempty"`      // IPerf direction (UPLOAD, DOWNLOAD, BOTH)
	StartPort      *int     `json:"startPort,omitempty"`      // IPerf start port
	EndPort        *int     `json:"endPort,omitempty"`        // IPerf end port
	UDPBandwidth   *int     `json:"udpBandwidth,omitempty"`   // IPerf UDP bandwidth
	ProbeType      string   `json:"probeType,omitempty"`      // Probe type
	NumPackets     *int     `json:"numPackets,omitempty"`     // Number of packets
	PathToDownload string   `json:"pathToDownload,omitempty"` // File path for file transfer
	TransferType   string   `json:"transferType,omitempty"`   // File transfer type (UPLOAD, DOWNLOAD, BOTH)
	SharedSecret   string   `json:"sharedSecret,omitempty"`   // Shared secret
	NdtServer      string   `json:"ndtServer,omitempty"`      // NDT server
	NdtServerPort  string   `json:"ndtServerPort,omitempty"`  // NDT server port
	NdtServerPath  string   `json:"ndtServerPath,omitempty"`  // NDT server path
	UplinkTest     *bool    `json:"uplinkTest,omitempty"`     // Uplink test
	DownlinkTest   *bool    `json:"downlinkTest,omitempty"`   // Downlink test
	ProxyServer    string   `json:"proxyServer,omitempty"`    // Proxy server
	ProxyPort      string   `json:"proxyPort,omitempty"`      // Proxy port
	ProxyUserName  string   `json:"proxyUserName,omitempty"`  // Proxy user name
	ProxyPassword  string   `json:"proxyPassword,omitempty"`  // Proxy password
	UserNamePrompt string   `json:"userNamePrompt,omitempty"` // User name prompt
	PasswordPrompt string   `json:"passwordPrompt,omitempty"` // Password prompt
	ExitCommand    string   `json:"exitCommand,omitempty"`    // Exit command
	FinalPrompt    string   `json:"finalPrompt,omitempty"`    // Final prompt
}
type RequestSensorsCreateSensorTestTemplateV1Profiles struct {
	AuthType            string                                                               `json:"authType,omitempty"`            // Authentication type: OPEN, WPA2_PSK, WPA2_EaP, WEB_AUTH, MAB, DOT1X, OTHER
	Psk                 string                                                               `json:"psk,omitempty"`                 // Password of SSID when passwordType is ASCII
	Username            string                                                               `json:"username,omitempty"`            // User name string for onboarding SSID
	Password            string                                                               `json:"password,omitempty"`            // Password string for onboarding SSID
	PasswordType        string                                                               `json:"passwordType,omitempty"`        // SSID password type: ASCII or HEX
	EapMethod           string                                                               `json:"eapMethod,omitempty"`           // WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC
	Scep                *bool                                                                `json:"scep,omitempty"`                // Secure certificate enrollment protocol: true or false or null for not applicable
	AuthProtocol        string                                                               `json:"authProtocol,omitempty"`        // Auth protocol
	Certfilename        string                                                               `json:"certfilename,omitempty"`        // Auth certificate file name
	Certxferprotocol    string                                                               `json:"certxferprotocol,omitempty"`    // Certificate transfering protocol: HTTP or HTTPS
	Certstatus          string                                                               `json:"certstatus,omitempty"`          // Certificate status: INACTIVE or ACTIVE
	Certpassphrase      string                                                               `json:"certpassphrase,omitempty"`      // Certificate password phrase
	Certdownloadurl     string                                                               `json:"certdownloadurl,omitempty"`     // Certificate download URL
	ExtWebAuthVirtualIP string                                                               `json:"extWebAuthVirtualIp,omitempty"` // External WEB Auth virtual IP
	ExtWebAuth          *bool                                                                `json:"extWebAuth,omitempty"`          // Indication of using external WEB Auth
	WhiteList           *bool                                                                `json:"whiteList,omitempty"`           // Indication of being on allowed list
	ExtWebAuthPortal    string                                                               `json:"extWebAuthPortal,omitempty"`    // External authentication portal
	ExtWebAuthAccessURL string                                                               `json:"extWebAuthAccessUrl,omitempty"` // External WEB Auth access URL
	ExtWebAuthHTMLTag   *[]RequestSensorsCreateSensorTestTemplateV1ProfilesExtWebAuthHTMLTag `json:"extWebAuthHtmlTag,omitempty"`   //
	QosPolicy           string                                                               `json:"qosPolicy,omitempty"`           // QoS policy: PlATINUM, GOLD, SILVER, BRONZE
	Tests               *[]RequestSensorsCreateSensorTestTemplateV1ProfilesTests             `json:"tests,omitempty"`               //
	ProfileName         string                                                               `json:"profileName,omitempty"`         // Profile name
	DeviceType          string                                                               `json:"deviceType,omitempty"`          // Device Type
	VLAN                string                                                               `json:"vlan,omitempty"`                // VLAN
	LocationVLANList    *[]RequestSensorsCreateSensorTestTemplateV1ProfilesLocationVLANList  `json:"locationVlanList,omitempty"`    //
}
type RequestSensorsCreateSensorTestTemplateV1ProfilesExtWebAuthHTMLTag struct {
	Label string `json:"label,omitempty"` // Label
	Tag   string `json:"tag,omitempty"`   // Tag
	Value string `json:"value,omitempty"` // Value
}
type RequestSensorsCreateSensorTestTemplateV1ProfilesTests struct {
	Name   string                                                         `json:"name,omitempty"`   // Name of the test
	Config *[]RequestSensorsCreateSensorTestTemplateV1ProfilesTestsConfig `json:"config,omitempty"` //
}
type RequestSensorsCreateSensorTestTemplateV1ProfilesTestsConfig struct {
	Domains        []string `json:"domains,omitempty"`        // DNS domain name
	Server         string   `json:"server,omitempty"`         // Ping, file transfer, mail, radius, ssh, or telnet server
	UserName       string   `json:"userName,omitempty"`       // User name
	Password       string   `json:"password,omitempty"`       // Password
	URL            string   `json:"url,omitempty"`            // URL
	Port           *int     `json:"port,omitempty"`           // Radius or WEB server port
	Protocol       string   `json:"protocol,omitempty"`       // Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)
	Servers        []string `json:"servers,omitempty"`        // IPerf server list
	Direction      string   `json:"direction,omitempty"`      // IPerf direction (UPLOAD, DOWNLOAD, BOTH)
	StartPort      *int     `json:"startPort,omitempty"`      // IPerf start port
	EndPort        *int     `json:"endPort,omitempty"`        // IPerf end port
	UDPBandwidth   *int     `json:"udpBandwidth,omitempty"`   // IPerf UDP bandwidth
	ProbeType      string   `json:"probeType,omitempty"`      // Probe type
	NumPackets     *int     `json:"numPackets,omitempty"`     // Number of packets
	PathToDownload string   `json:"pathToDownload,omitempty"` // File path for file transfer
	TransferType   string   `json:"transferType,omitempty"`   // File transfer type (UPLOAD, DOWNLOAD, BOTH)
	SharedSecret   string   `json:"sharedSecret,omitempty"`   // Shared secret
	NdtServer      string   `json:"ndtServer,omitempty"`      // NDT server
	NdtServerPort  string   `json:"ndtServerPort,omitempty"`  // NDT server port
	NdtServerPath  string   `json:"ndtServerPath,omitempty"`  // NDT server path
	UplinkTest     *bool    `json:"uplinkTest,omitempty"`     // Uplink test
	DownlinkTest   *bool    `json:"downlinkTest,omitempty"`   // Downlink test
	ProxyServer    string   `json:"proxyServer,omitempty"`    // Proxy server
	ProxyPort      string   `json:"proxyPort,omitempty"`      // Proxy port
	ProxyUserName  string   `json:"proxyUserName,omitempty"`  // Proxy user name
	ProxyPassword  string   `json:"proxyPassword,omitempty"`  // Proxy password
	UserNamePrompt string   `json:"userNamePrompt,omitempty"` // User name prompt
	PasswordPrompt string   `json:"passwordPrompt,omitempty"` // Password prompt
	ExitCommand    string   `json:"exitCommand,omitempty"`    // Exit command
	FinalPrompt    string   `json:"finalPrompt,omitempty"`    // Final prompt
}
type RequestSensorsCreateSensorTestTemplateV1ProfilesLocationVLANList struct {
	LocationID string   `json:"locationId,omitempty"` // Site UUID
	VLANs      []string `json:"vlans,omitempty"`      // Array of VLANs
}
type RequestSensorsCreateSensorTestTemplateV1LocationInfoList struct {
	LocationID           string   `json:"locationId,omitempty"`           // Site UUID
	LocationType         string   `json:"locationType,omitempty"`         // Site type
	AllSensors           *bool    `json:"allSensors,omitempty"`           // Use all sensors in the site for test
	SiteHierarchy        string   `json:"siteHierarchy,omitempty"`        // Site name hierarhy
	MacAddressList       []string `json:"macAddressList,omitempty"`       // MAC addresses
	ManagementVLAN       string   `json:"managementVlan,omitempty"`       // Management VLAN
	CustomManagementVLAN *bool    `json:"customManagementVlan,omitempty"` // Custom Management VLAN
}
type RequestSensorsCreateSensorTestTemplateV1Sensors struct {
	Name                    string                                                           `json:"name,omitempty"`                    // Sensor name
	MacAddress              string                                                           `json:"macAddress,omitempty"`              // MAC address
	SwitchMac               string                                                           `json:"switchMac,omitempty"`               // Switch MAC address
	SwitchUUID              string                                                           `json:"switchUuid,omitempty"`              // Switch device UUID
	SwitchSerialNumber      string                                                           `json:"switchSerialNumber,omitempty"`      // Switch serial number
	MarkedForUninstall      *bool                                                            `json:"markedForUninstall,omitempty"`      // Is marked for uninstall
	IPAddress               string                                                           `json:"ipAddress,omitempty"`               // IP address
	HostName                string                                                           `json:"hostName,omitempty"`                // Host name
	WiredApplicationStatus  string                                                           `json:"wiredApplicationStatus,omitempty"`  // Wired application status
	WiredApplicationMessage string                                                           `json:"wiredApplicationMessage,omitempty"` // Wired application message
	Assigned                *bool                                                            `json:"assigned,omitempty"`                // Is assigned
	Status                  string                                                           `json:"status,omitempty"`                  // Sensor device status: UP, DOWN, REBOOT
	XorSensor               *bool                                                            `json:"xorSensor,omitempty"`               // Is XOR sensor
	TargetAPs               []string                                                         `json:"targetAPs,omitempty"`               // Array of target APs
	RunNow                  string                                                           `json:"runNow,omitempty"`                  // Run now: YES, NO
	LocationID              string                                                           `json:"locationId,omitempty"`              // Site UUID
	AllSensorAddition       *bool                                                            `json:"allSensorAddition,omitempty"`       // Is all sensor addition
	ConfigUpdated           string                                                           `json:"configUpdated,omitempty"`           // Configuration updated: YES, NO
	SensorType              string                                                           `json:"sensorType,omitempty"`              // Sensor type
	TestMacAddresses        *RequestSensorsCreateSensorTestTemplateV1SensorsTestMacAddresses `json:"testMacAddresses,omitempty"`        // A string-string test MAC address
	ID                      string                                                           `json:"id,omitempty"`                      // Sensor ID
	ServicePolicy           string                                                           `json:"servicePolicy,omitempty"`           // Service policy
	IPerfInfo               *RequestSensorsCreateSensorTestTemplateV1SensorsIPerfInfo        `json:"iPerfInfo,omitempty"`               // A string-stringList iPerf information
}
type RequestSensorsCreateSensorTestTemplateV1SensorsTestMacAddresses interface{}
type RequestSensorsCreateSensorTestTemplateV1SensorsIPerfInfo interface{}
type RequestSensorsCreateSensorTestTemplateV1ApCoverage struct {
	Bands             string `json:"bands,omitempty"`             // The WIFI bands
	NumberOfApsToTest *int   `json:"numberOfApsToTest,omitempty"` // Number of APs to test
	RssiThreshold     *int   `json:"rssiThreshold,omitempty"`     // RSSI threshold
}
type RequestSensorsRunNowSensorTestV1 struct {
	TemplateName string `json:"templateName,omitempty"` // Template Name
}
type RequestSensorsDuplicateSensorTestTemplateV1 struct {
	TemplateName    string `json:"templateName,omitempty"`    // Source test template name
	NewTemplateName string `json:"newTemplateName,omitempty"` // Destination test template name
}

//SensorsV1 Sensors - 71a1-2bb7-4569-9cc5
/* Intent API to get a list of SENSOR devices


@param SensorsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!sensors-v1
*/
func (s *SensorsService) SensorsV1(SensorsV1QueryParams *SensorsV1QueryParams) (*ResponseSensorsSensorsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sensor"

	queryString, _ := query.Values(SensorsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSensorsSensorsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SensorsV1(SensorsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation SensorsV1")
	}

	result := response.Result().(*ResponseSensorsSensorsV1)
	return result, response, err

}

//CreateSensorTestTemplateV1 Create sensor test template - 08bd-8883-4a68-a2e6
/* Intent API to create a SENSOR test template with a new SSID, existing SSID, or both new and existing SSID



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-sensor-test-template-v1
*/
func (s *SensorsService) CreateSensorTestTemplateV1(requestSensorsCreateSensorTestTemplateV1 *RequestSensorsCreateSensorTestTemplateV1) (*ResponseSensorsCreateSensorTestTemplateV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sensor"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSensorsCreateSensorTestTemplateV1).
		SetResult(&ResponseSensorsCreateSensorTestTemplateV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateSensorTestTemplateV1(requestSensorsCreateSensorTestTemplateV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateSensorTestTemplateV1")
	}

	result := response.Result().(*ResponseSensorsCreateSensorTestTemplateV1)
	return result, response, err

}

//EditSensorTestTemplateV1 Edit sensor test template - c085-eaf5-4f89-ba34
/* Intent API to deploy, schedule, or edit and existing SENSOR test template


 */
func (s *SensorsService) EditSensorTestTemplateV1(requestSensorsEditSensorTestTemplateV1 *RequestSensorsEditSensorTestTemplateV1) (*ResponseSensorsEditSensorTestTemplateV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/AssuranceScheduleSensorTest"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSensorsEditSensorTestTemplateV1).
		SetResult(&ResponseSensorsEditSensorTestTemplateV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.EditSensorTestTemplateV1(requestSensorsEditSensorTestTemplateV1)
		}
		return nil, response, fmt.Errorf("error with operation EditSensorTestTemplateV1")
	}

	result := response.Result().(*ResponseSensorsEditSensorTestTemplateV1)
	return result, response, err

}

//RunNowSensorTestV1 Run now sensor test - f1a7-a8e7-4cf9-9c8f
/* Intent API to run a deployed SENSOR test


 */
func (s *SensorsService) RunNowSensorTestV1(requestSensorsRunNowSensorTestV1 *RequestSensorsRunNowSensorTestV1) (*resty.Response, error) {
	path := "/dna/intent/api/v1/sensor-run-now"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSensorsRunNowSensorTestV1).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RunNowSensorTestV1(requestSensorsRunNowSensorTestV1)
		}
		return response, fmt.Errorf("error with operation RunNowSensorTestV1")
	}

	return response, err

}

//DuplicateSensorTestTemplateV1 Duplicate sensor test template - 85a2-8837-4909-9021
/* Intent API to duplicate an existing SENSOR test template


 */
func (s *SensorsService) DuplicateSensorTestTemplateV1(requestSensorsDuplicateSensorTestTemplateV1 *RequestSensorsDuplicateSensorTestTemplateV1) (*ResponseSensorsDuplicateSensorTestTemplateV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/sensorTestTemplate"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSensorsDuplicateSensorTestTemplateV1).
		SetResult(&ResponseSensorsDuplicateSensorTestTemplateV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DuplicateSensorTestTemplateV1(requestSensorsDuplicateSensorTestTemplateV1)
		}
		return nil, response, fmt.Errorf("error with operation DuplicateSensorTestTemplateV1")
	}

	result := response.Result().(*ResponseSensorsDuplicateSensorTestTemplateV1)
	return result, response, err

}

//DeleteSensorTestV1 Delete sensor test - 5bbb-28ff-442a-825f
/* Intent API to delete an existing SENSOR test template


@param DeleteSensorTestV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-sensor-test-v1
*/
func (s *SensorsService) DeleteSensorTestV1(DeleteSensorTestV1QueryParams *DeleteSensorTestV1QueryParams) (*ResponseSensorsDeleteSensorTestV1, *resty.Response, error) {
	//DeleteSensorTestV1QueryParams *DeleteSensorTestV1QueryParams
	path := "/dna/intent/api/v1/sensor"

	queryString, _ := query.Values(DeleteSensorTestV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSensorsDeleteSensorTestV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteSensorTestV1(DeleteSensorTestV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteSensorTestV1")
	}

	result := response.Result().(*ResponseSensorsDeleteSensorTestV1)
	return result, response, err

}

// Alias Function
func (s *SensorsService) CreateSensorTestTemplate(requestSensorsCreateSensorTestTemplateV1 *RequestSensorsCreateSensorTestTemplateV1) (*ResponseSensorsCreateSensorTestTemplateV1, *resty.Response, error) {
	return s.CreateSensorTestTemplateV1(requestSensorsCreateSensorTestTemplateV1)
}

// Alias Function
func (s *SensorsService) RunNowSensorTest(requestSensorsRunNowSensorTestV1 *RequestSensorsRunNowSensorTestV1) (*resty.Response, error) {
	return s.RunNowSensorTestV1(requestSensorsRunNowSensorTestV1)
}

// Alias Function
func (s *SensorsService) DeleteSensorTest(DeleteSensorTestV1QueryParams *DeleteSensorTestV1QueryParams) (*ResponseSensorsDeleteSensorTestV1, *resty.Response, error) {
	return s.DeleteSensorTestV1(DeleteSensorTestV1QueryParams)
}

// Alias Function
func (s *SensorsService) DuplicateSensorTestTemplate(requestSensorsDuplicateSensorTestTemplateV1 *RequestSensorsDuplicateSensorTestTemplateV1) (*ResponseSensorsDuplicateSensorTestTemplateV1, *resty.Response, error) {
	return s.DuplicateSensorTestTemplateV1(requestSensorsDuplicateSensorTestTemplateV1)
}

// Alias Function
func (s *SensorsService) Sensors(SensorsV1QueryParams *SensorsV1QueryParams) (*ResponseSensorsSensorsV1, *resty.Response, error) {
	return s.SensorsV1(SensorsV1QueryParams)
}

// Alias Function
func (s *SensorsService) EditSensorTestTemplate(requestSensorsEditSensorTestTemplateV1 *RequestSensorsEditSensorTestTemplateV1) (*ResponseSensorsEditSensorTestTemplateV1, *resty.Response, error) {
	return s.EditSensorTestTemplateV1(requestSensorsEditSensorTestTemplateV1)
}
