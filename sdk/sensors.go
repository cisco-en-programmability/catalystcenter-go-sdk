package catalyst

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SensorsService service

type ListsICapPacketCaptureFilesMatchingSpecifiedCriteriaV1QueryParams struct {
	Type      string  `url:"type,omitempty"`      //Capture Type
	ClientMac string  `url:"clientMac,omitempty"` //The macAddress of client
	ApMac     string  `url:"apMac,omitempty"`     //The base radio macAddress of the access point
	StartTime float64 `url:"startTime,omitempty"` //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime   float64 `url:"endTime,omitempty"`   //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	Limit     float64 `url:"limit,omitempty"`     //Maximum number of records to return
	Offset    float64 `url:"offset,omitempty"`    //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	SortBy    string  `url:"sortBy,omitempty"`    //A field within the response to sort by.
	Order     string  `url:"order,omitempty"`     //The sort order of the field ascending or descending.
}
type ListsICapPacketCaptureFilesMatchingSpecifiedCriteriaV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaV1QueryParams struct {
	Type      string  `url:"type,omitempty"`      //Capture Type
	ClientMac string  `url:"clientMac,omitempty"` //The macAddress of client
	ApMac     string  `url:"apMac,omitempty"`     //The base radio macAddress of the access point
	StartTime float64 `url:"startTime,omitempty"` //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime   float64 `url:"endTime,omitempty"`   //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
}
type RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesDetailsOfASpecificICapPacketCaptureFileV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type DownloadsASpecificICapPacketCaptureFileV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1HeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1HeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMacV1QueryParams struct {
	StartTime     float64 `url:"startTime,omitempty"`     //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime       float64 `url:"endTime,omitempty"`       //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	ApMac         string  `url:"apMac,omitempty"`         //The base ethernet macAddress of the access point
	Limit         float64 `url:"limit,omitempty"`         //Maximum number of records to return
	Offset        float64 `url:"offset,omitempty"`        //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	TimeSortOrder string  `url:"timeSortOrder,omitempty"` //The sort order of the field ascending or descending.
}
type RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMacV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMacV1QueryParams struct {
	StartTime     float64 `url:"startTime,omitempty"`     //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime       float64 `url:"endTime,omitempty"`       //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	ApMac         string  `url:"apMac,omitempty"`         //The base ethernet macAddress of the access point
	DataType      float64 `url:"dataType,omitempty"`      //Data type reported by the sensor |Data Type | Description | | --- | --- | | `0` | Duty Cycle | | `1` | Max Power | | `2` | Average Power | | `3` | Max Power in dBm with adjusted base of +48 | | `4` | Average Power in dBm with adjusted base of +48 |
	Limit         float64 `url:"limit,omitempty"`         //Maximum number of records to return
	Offset        float64 `url:"offset,omitempty"`        //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	TimeSortOrder string  `url:"timeSortOrder,omitempty"` //The sort order of the field ascending or descending.
}
type RetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMacV1HeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesDeployedICapConfigurationsWhileSupportingBasicFilteringV1QueryParams struct {
	CaptureStatus string  `url:"captureStatus,omitempty"` //Catalyst Center ICAP status
	CaptureType   string  `url:"captureType,omitempty"`   //Catalyst Center ICAP type
	ClientMac     string  `url:"clientMac,omitempty"`     //The client device MAC address in ICAP configuration
	APID          string  `url:"apId,omitempty"`          //The AP device's UUID.
	WlcID         string  `url:"wlcId,omitempty"`         //The wireless controller device's UUID
	Offset        float64 `url:"offset,omitempty"`        //The first record to show for this page; the first record is numbered 1.
	Limit         float64 `url:"limit,omitempty"`         //The number of records to show for this page.
}
type CreatesAnICapConfigurationIntentForPreviewApproveV1QueryParams struct {
	PreviewDescription string `url:"previewDescription,omitempty"` //The ICAP intent's preview-deploy description string
}
type RetrievesTheCountOfDeployedICapConfigurationsWhileSupportingBasicFilteringV1QueryParams struct {
	CaptureType   string `url:"captureType,omitempty"`   //Catalyst Center ICAP type
	CaptureStatus string `url:"captureStatus,omitempty"` //Catalyst Center ICAP status
	ClientMac     string `url:"clientMac,omitempty"`     //The client device MAC address in ICAP configuration
	APID          string `url:"apId,omitempty"`          //The AP device's UUID.
	WlcID         string `url:"wlcId,omitempty"`         //The wireless controller device's UUID
}
type DeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1QueryParams struct {
	PreviewDescription string `url:"previewDescription,omitempty"` //The ICAP intent's preview-deploy description string
}
type GetDeviceDeploymentStatusV1QueryParams struct {
	DeployActivityID string  `url:"deployActivityId,omitempty"` //activity from the /deploy task response
	NetworkDeviceIDs string  `url:"networkDeviceIds,omitempty"` //device ids, retrievable from the id attribute in intent/api/v1/network-device
	Offset           float64 `url:"offset,omitempty"`           //The first record to show for this page; the first record is numbered 1.
	Limit            float64 `url:"limit,omitempty"`            //The number of records to show for this page.
	SortBy           string  `url:"sortBy,omitempty"`           //A property within the response to sort by.
	Order            string  `url:"order,omitempty"`            //Whether ascending or descending order should be used to sort the response.
}
type GetDeviceDeploymentStatusCountV1QueryParams struct {
	DeployActivityID string `url:"deployActivityId,omitempty"` //activity from the /deploy task response
	NetworkDeviceIDs string `url:"networkDeviceIds,omitempty"` //device ids, retrievable from the id attribute in intent/api/v1/network-device
}
type DeleteSensorTestV1QueryParams struct {
	TemplateName string `url:"templateName,omitempty"` //
}
type SensorsV1QueryParams struct {
	SiteID string `url:"siteId,omitempty"` //
}

type ResponseSensorsListsICapPacketCaptureFilesMatchingSpecifiedCriteriaV1 struct {
	Response *[]ResponseSensorsListsICapPacketCaptureFilesMatchingSpecifiedCriteriaV1Response `json:"response,omitempty"` //

	Page *ResponseSensorsListsICapPacketCaptureFilesMatchingSpecifiedCriteriaV1Page `json:"page,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsListsICapPacketCaptureFilesMatchingSpecifiedCriteriaV1Response struct {
	ID string `json:"id,omitempty"` // Id

	FileName string `json:"fileName,omitempty"` // File Name

	FileSize *int `json:"fileSize,omitempty"` // File Size

	Type string `json:"type,omitempty"` // Type

	ClientMac string `json:"clientMac,omitempty"` // Client Mac

	ApMac string `json:"apMac,omitempty"` // Ap Mac

	FileCreationTimestamp *int `json:"fileCreationTimestamp,omitempty"` // File Creation Timestamp

	LastUpdatedTimestamp *int `json:"lastUpdatedTimestamp,omitempty"` // Last Updated Timestamp
}
type ResponseSensorsListsICapPacketCaptureFilesMatchingSpecifiedCriteriaV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	Count *int `json:"count,omitempty"` // Count

	SortBy string `json:"sortBy,omitempty"` // Sort By

	Order string `json:"order,omitempty"` // Order
}
type ResponseSensorsRetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaV1 struct {
	Response *ResponseSensorsRetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsRetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaV1Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseSensorsRetrievesDetailsOfASpecificICapPacketCaptureFileV1 struct {
	Response *ResponseSensorsRetrievesDetailsOfASpecificICapPacketCaptureFileV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsRetrievesDetailsOfASpecificICapPacketCaptureFileV1Response struct {
	ID string `json:"id,omitempty"` // Id

	FileName string `json:"fileName,omitempty"` // File Name

	FileSize *int `json:"fileSize,omitempty"` // File Size

	Type string `json:"type,omitempty"` // Type

	ClientMac string `json:"clientMac,omitempty"` // Client Mac

	ApMac string `json:"apMac,omitempty"` // Ap Mac

	FileCreationTimestamp *int `json:"fileCreationTimestamp,omitempty"` // File Creation Timestamp

	LastUpdatedTimestamp *int `json:"lastUpdatedTimestamp,omitempty"` // Last Updated Timestamp
}
type ResponseSensorsDownloadsASpecificICapPacketCaptureFileV1 struct {
	object string `json:"object,omitempty"` // object
}
type ResponseSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1 struct {
	Response *[]ResponseSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1Response `json:"response,omitempty"` //

	Page *ResponseSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1Page `json:"page,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1Response struct {
	ID string `json:"id,omitempty"` // Id

	ClientMac string `json:"clientMac,omitempty"` // Client Mac

	ApMac string `json:"apMac,omitempty"` // Ap Mac

	RadioID *int `json:"radioId,omitempty"` // Radio Id

	Timestamp *int `json:"timestamp,omitempty"` // Timestamp

	Band string `json:"band,omitempty"` // Band

	SSID string `json:"ssid,omitempty"` // Ssid

	Rssi *int `json:"rssi,omitempty"` // Rssi

	Snr *int `json:"snr,omitempty"` // Snr

	TxBytes *int `json:"txBytes,omitempty"` // Tx Bytes

	RxBytes *int `json:"rxBytes,omitempty"` // Rx Bytes

	RxPackets *int `json:"rxPackets,omitempty"` // Rx Packets

	TxPackets *int `json:"txPackets,omitempty"` // Tx Packets

	RxMgmtPackets *int `json:"rxMgmtPackets,omitempty"` // Rx Mgmt Packets

	TxMgmtPackets *int `json:"txMgmtPackets,omitempty"` // Tx Mgmt Packets

	RxDataPackets *int `json:"rxDataPackets,omitempty"` // Rx Data Packets

	TxDataPackets *int `json:"txDataPackets,omitempty"` // Tx Data Packets

	TxUnicastDataPackets *float64 `json:"txUnicastDataPackets,omitempty"` // Tx Unicast Data Packets

	RxCtrlPackets *int `json:"rxCtrlPackets,omitempty"` // Rx Ctrl Packets

	TxCtrlPackets *int `json:"txCtrlPackets,omitempty"` // Tx Ctrl Packets

	RxRetries *int `json:"rxRetries,omitempty"` // Rx Retries

	RxRate *int `json:"rxRate,omitempty"` // Rx Rate

	TxRate *int `json:"txRate,omitempty"` // Tx Rate

	ClientIP string `json:"clientIp,omitempty"` // Client Ip
}
type ResponseSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	Count *int `json:"count,omitempty"` // Count

	TimeSortOrder string `json:"timeSortOrder,omitempty"` // Time Sort Order
}
type ResponseSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1 struct {
	Response *[]ResponseSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1Response `json:"response,omitempty"` //

	Page *ResponseSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1Page `json:"page,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1Response struct {
	Timestamp *int `json:"timestamp,omitempty"` // Timestamp

	ApMac string `json:"apMac,omitempty"` // Ap Mac

	RadioID *int `json:"radioId,omitempty"` // Radio Id

	Band string `json:"band,omitempty"` // Band

	Utilization *int `json:"utilization,omitempty"` // Utilization

	NonWifiUtilization *float64 `json:"nonWifiUtilization,omitempty"` // Non Wifi Utilization

	RxOtherBSSUtilization *int `json:"rxOtherBSSUtilization,omitempty"` // Rx Other B S S Utilization

	RxInBSSUtilization *float64 `json:"rxInBSSUtilization,omitempty"` // Rx In B S S Utilization

	TxUtilization *int `json:"txUtilization,omitempty"` // Tx Utilization

	NoiseFloor *int `json:"noiseFloor,omitempty"` // Noise Floor

	Channel *int `json:"channel,omitempty"` // Channel

	ChannelWidth *int `json:"channelWidth,omitempty"` // Channel Width

	TxPower *int `json:"txPower,omitempty"` // Tx Power

	MaxTxPower *float64 `json:"maxTxPower,omitempty"` // Max Tx Power

	TxBytes *int `json:"txBytes,omitempty"` // Tx Bytes

	RxBytes *int `json:"rxBytes,omitempty"` // Rx Bytes

	RxPackets *int `json:"rxPackets,omitempty"` // Rx Packets

	TxPackets *int `json:"txPackets,omitempty"` // Tx Packets

	RxMgmtPackets *int `json:"rxMgmtPackets,omitempty"` // Rx Mgmt Packets

	TxMgmtPackets *int `json:"txMgmtPackets,omitempty"` // Tx Mgmt Packets

	RxErrors *int `json:"rxErrors,omitempty"` // Rx Errors

	TxErrors *int `json:"txErrors,omitempty"` // Tx Errors
}
type ResponseSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	Count *int `json:"count,omitempty"` // Count

	TimeSortOrder string `json:"timeSortOrder,omitempty"` // Time Sort Order
}
type ResponseSensorsRetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMacV1 struct {
	Response *[]ResponseSensorsRetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMacV1Response `json:"response,omitempty"` //

	Page *ResponseSensorsRetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMacV1Page `json:"page,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsRetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMacV1Response struct {
	ID string `json:"id,omitempty"` // Id

	ApMac string `json:"apMac,omitempty"` // Ap Mac

	CentralFrequencyKHz *int `json:"centralFrequencyKHz,omitempty"` // Central Frequency K Hz

	BandWidthKHz *int `json:"bandWidthKHz,omitempty"` // Band Width K Hz

	LowEndFrequencyKHz *int `json:"lowEndFrequencyKHz,omitempty"` // Low End Frequency K Hz

	HighEndFrequencyKHz *int `json:"highEndFrequencyKHz,omitempty"` // High End Frequency K Hz

	PowerDbm *float64 `json:"powerDbm,omitempty"` // Power Dbm

	Band string `json:"band,omitempty"` // Band

	DutyCycle *int `json:"dutyCycle,omitempty"` // Duty Cycle

	Timestamp *int `json:"timestamp,omitempty"` // Timestamp

	DeviceType string `json:"deviceType,omitempty"` // Device Type

	SeverityIndex *int `json:"severityIndex,omitempty"` // Severity Index

	DetectedChannels *[]int `json:"detectedChannels,omitempty"` // Detected Channels
}
type ResponseSensorsRetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMacV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	Count *int `json:"count,omitempty"` // Count

	TimeSortOrder string `json:"timeSortOrder,omitempty"` // Time Sort Order
}
type ResponseSensorsRetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMacV1 struct {
	Response *[]ResponseSensorsRetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMacV1Response `json:"response,omitempty"` //

	Page *ResponseSensorsRetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMacV1Page `json:"page,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsRetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMacV1Response struct {
	ID string `json:"id,omitempty"` // Id

	SpanKHz *int `json:"spanKHz,omitempty"` // Span K Hz

	DataType *int `json:"dataType,omitempty"` // Data Type

	ApMac string `json:"apMac,omitempty"` // Ap Mac

	DataAvg *float64 `json:"dataAvg,omitempty"` // Data Avg

	DataMin *float64 `json:"dataMin,omitempty"` // Data Min

	DataMax *float64 `json:"dataMax,omitempty"` // Data Max

	DataUnits string `json:"dataUnits,omitempty"` // Data Units

	CentralFrequencyKHz *int `json:"centralFrequencyKHz,omitempty"` // Central Frequency K Hz

	Band string `json:"band,omitempty"` // Band

	Timestamp *int `json:"timestamp,omitempty"` // Timestamp

	Data *[]float64 `json:"data,omitempty"` // Data

	DataSize *int `json:"dataSize,omitempty"` // Data Size

	Channels *[]int `json:"channels,omitempty"` // Channels
}
type ResponseSensorsRetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMacV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	Count *int `json:"count,omitempty"` // Count

	TimeSortOrder string `json:"timeSortOrder,omitempty"` // Time Sort Order
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
	Bands string `json:"bands,omitempty"` // The WIFI bands

	NumberOfApsToTest *int `json:"numberOfApsToTest,omitempty"` // Number of APs to test

	RssiThreshold *int `json:"rssiThreshold,omitempty"` // RSSI threshold
}
type ResponseSensorsRetrievesDeployedICapConfigurationsWhileSupportingBasicFilteringV1 struct {
	Response *[]ResponseSensorsRetrievesDeployedICapConfigurationsWhileSupportingBasicFilteringV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsRetrievesDeployedICapConfigurationsWhileSupportingBasicFilteringV1Response struct {
	Name string `json:"name,omitempty"` // Name

	Slots *[]float64 `json:"slots,omitempty"` // Slots

	OtaBand string `json:"otaBand,omitempty"` // Ota Band

	OtaChannel *int `json:"otaChannel,omitempty"` // Ota Channel

	OtaChannelWidth *int `json:"otaChannelWidth,omitempty"` // Ota Channel Width

	ID string `json:"id,omitempty"` // Id

	DeployedID string `json:"deployedId,omitempty"` // Deployed Id

	DisableActivityID string `json:"disableActivityId,omitempty"` // Disable Activity Id

	ActivityID string `json:"activityId,omitempty"` // Activity Id

	CaptureType string `json:"captureType,omitempty"` // Capture Type

	APID string `json:"apId,omitempty"` // Ap Id

	WlcID string `json:"wlcId,omitempty"` // Wlc Id

	ClientMac string `json:"clientMac,omitempty"` // Client Mac

	CreateTime *int `json:"createTime,omitempty"` // Create Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	DurationInMins *int `json:"durationInMins,omitempty"` // Duration In Mins

	Status string `json:"status,omitempty"` // Status
}
type ResponseSensorsCreatesAnICapConfigurationIntentForPreviewApproveV1 struct {
	Response *ResponseSensorsCreatesAnICapConfigurationIntentForPreviewApproveV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsCreatesAnICapConfigurationIntentForPreviewApproveV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	URL string `json:"url,omitempty"` // Url
}
type ResponseSensorsCreatesAICapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDeviceV1 struct {
	Response *ResponseSensorsCreatesAICapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDeviceV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsCreatesAICapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDeviceV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	URL string `json:"url,omitempty"` // Url
}
type ResponseSensorsDiscardsTheICapConfigurationIntentByActivityIDV1 struct {
	Response *ResponseSensorsDiscardsTheICapConfigurationIntentByActivityIDV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsDiscardsTheICapConfigurationIntentByActivityIDV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	URL string `json:"url,omitempty"` // Url
}
type ResponseSensorsDeploysTheICapConfigurationIntentByActivityIDV1 struct {
	Response *ResponseSensorsDeploysTheICapConfigurationIntentByActivityIDV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsDeploysTheICapConfigurationIntentByActivityIDV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	URL string `json:"url,omitempty"` // Url
}
type ResponseSensorsGetICapConfigurationStatusPerNetworkDeviceV1 struct {
	Response *[]ResponseSensorsGetICapConfigurationStatusPerNetworkDeviceV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsGetICapConfigurationStatusPerNetworkDeviceV1Response struct {
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Network Device Id

	Status string `json:"status,omitempty"` // Status
}
type ResponseSensorsRetrievesTheDevicesClisOfTheICapintentV1 struct {
	Response *ResponseSensorsRetrievesTheDevicesClisOfTheICapintentV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsRetrievesTheDevicesClisOfTheICapintentV1Response struct {
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Network Device Id

	PreviewItems *[]ResponseSensorsRetrievesTheDevicesClisOfTheICapintentV1ResponsePreviewItems `json:"previewItems,omitempty"` //

	Status string `json:"status,omitempty"` // Status
}
type ResponseSensorsRetrievesTheDevicesClisOfTheICapintentV1ResponsePreviewItems struct {
	ConfigPreview string `json:"configPreview,omitempty"` // Config Preview

	ConfigType string `json:"configType,omitempty"` // Config Type

	ErrorMessages []string `json:"errorMessages,omitempty"` // Error Messages

	Name string `json:"name,omitempty"` // Name
}
type ResponseSensorsGeneratesTheDevicesClisOfTheICapConfigurationIntentV1 struct {
	Response *ResponseSensorsGeneratesTheDevicesClisOfTheICapConfigurationIntentV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsGeneratesTheDevicesClisOfTheICapConfigurationIntentV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	URL string `json:"url,omitempty"` // Url
}
type ResponseSensorsRetrievesTheCountOfDeployedICapConfigurationsWhileSupportingBasicFilteringV1 struct {
	Response *float64 `json:"response,omitempty"` // Response

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1 struct {
	Response *ResponseSensorsDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	URL string `json:"url,omitempty"` // Url
}
type ResponseSensorsRemoveTheICapConfigurationOnTheDeviceWithoutPreviewV1 struct {
	Response *ResponseSensorsRemoveTheICapConfigurationOnTheDeviceWithoutPreviewV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsRemoveTheICapConfigurationOnTheDeviceWithoutPreviewV1Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	URL string `json:"url,omitempty"` // Url
}
type ResponseSensorsGetDeviceDeploymentStatusV1 struct {
	Response *[]ResponseSensorsGetDeviceDeploymentStatusV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsGetDeviceDeploymentStatusV1Response struct {
	DeployActivityID string `json:"deployActivityId,omitempty"` // Deploy Activity Id

	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Network Device Id

	ConfigGroupName string `json:"configGroupName,omitempty"` // Config Group Name

	ConfigGroupVersion *int `json:"configGroupVersion,omitempty"` // Config Group Version

	Status string `json:"status,omitempty"` // Status

	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	Error *ResponseSensorsGetDeviceDeploymentStatusV1ResponseError `json:"error,omitempty"` // Error
}
type ResponseSensorsGetDeviceDeploymentStatusV1ResponseError interface{}
type ResponseSensorsGetDeviceDeploymentStatusCountV1 struct {
	Response *ResponseSensorsGetDeviceDeploymentStatusCountV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsGetDeviceDeploymentStatusCountV1Response struct {
	Count *float64 `json:"count,omitempty"` // Count
}
type ResponseSensorsCreateSensorTestTemplateV1 struct {
	Version string `json:"version,omitempty"` // Version

	Response *ResponseSensorsCreateSensorTestTemplateV1Response `json:"response,omitempty"` //
}
type ResponseSensorsCreateSensorTestTemplateV1Response struct {
	Name   string `json:"name,omitempty"` // The sensor test template name
	TypeID string `json:"_id,omitempty"`  // (Used in edit only) The sensor test template unique identifier

	Version *int `json:"version,omitempty"` // The sensor test template version (must be 2)

	ModelVersion *int `json:"modelVersion,omitempty"` // Test template object model version (must be 2)

	StartTime *int `json:"startTime,omitempty"` // Start time

	LastModifiedTime *int `json:"lastModifiedTime,omitempty"` // Last modify time

	NumAssociatedSensor *int `json:"numAssociatedSensor,omitempty"` // Number of associated sensor

	Location string `json:"location,omitempty"` // Location string

	SiteHierarchy string `json:"siteHierarchy,omitempty"` // Site hierarchy

	Status string `json:"status,omitempty"` // Status of the test (RUNNING, NOTRUNNING)

	Connection string `json:"connection,omitempty"` // connection type of test: WIRED, WIRELESS, BOTH

	ActionInProgress string `json:"actionInProgress,omitempty"` // Indication of inprogress action

	Frequency *ResponseSensorsCreateSensorTestTemplateV1ResponseFrequency `json:"frequency,omitempty"` //

	RssiThreshold *int `json:"rssiThreshold,omitempty"` // RSSI threshold

	NumNeighborApThreshold *int `json:"numNeighborAPThreshold,omitempty"` // Number of neighboring AP threshold

	ScheduleInDays *int `json:"scheduleInDays,omitempty"` // Bit-wise value of scheduled test days

	WLANs []string `json:"wlans,omitempty"` // WLANs list

	SSIDs *[]ResponseSensorsCreateSensorTestTemplateV1ResponseSSIDs `json:"ssids,omitempty"` //

	Profiles *[]ResponseSensorsCreateSensorTestTemplateV1ResponseProfiles `json:"profiles,omitempty"` //

	TestScheduleMode string `json:"testScheduleMode,omitempty"` // Test schedule mode (ONDEMAND, DEDICATED, SCHEDULED, CONTINUOUS, RUNNOW)

	ShowWlcUpgradeBanner *bool `json:"showWlcUpgradeBanner,omitempty"` // Show WLC upgrade banner

	RadioAsSensorRemoved *bool `json:"radioAsSensorRemoved,omitempty"` // Radio as sensor removed

	EncryptionMode string `json:"encryptionMode,omitempty"` // Encryption mode

	RunNow string `json:"runNow,omitempty"` // Run now (YES, NO)

	LocationInfoList *[]ResponseSensorsCreateSensorTestTemplateV1ResponseLocationInfoList `json:"locationInfoList,omitempty"` //

	Sensors *[]ResponseSensorsCreateSensorTestTemplateV1ResponseSensors `json:"sensors,omitempty"` //

	ApCoverage *[]ResponseSensorsCreateSensorTestTemplateV1ResponseApCoverage `json:"apCoverage,omitempty"` //
}
type ResponseSensorsCreateSensorTestTemplateV1ResponseFrequency struct {
	Value *int `json:"value,omitempty"` // Value of the unit

	Unit string `json:"unit,omitempty"` // Unit of the time value (NANOSECONDS, MICROSECONDS, MILLISECONDS, SECONDS, MINUTES, HOURS, DAYS)
}
type ResponseSensorsCreateSensorTestTemplateV1ResponseSSIDs struct {
	Bands string `json:"bands,omitempty"` // WIFI bands: 2.4GHz or 5GHz

	SSID string `json:"ssid,omitempty"` // The SSID string

	ProfileName string `json:"profileName,omitempty"` // The SSID profile name string

	NumAps *int `json:"numAps,omitempty"` // Number of APs in the test

	NumSensors *int `json:"numSensors,omitempty"` // Number of Sensors in the test

	Layer3WebAuthsecurity string `json:"layer3webAuthsecurity,omitempty"` // Layer 3 WEB Auth security

	Layer3WebAuthuserName string `json:"layer3webAuthuserName,omitempty"` // Layer 3 WEB Auth user name

	Layer3WebAuthpassword string `json:"layer3webAuthpassword,omitempty"` // Layer 3 WEB Auth password

	Layer3WebAuthEmailAddress string `json:"layer3webAuthEmailAddress,omitempty"` // Layer 3 WEB Auth email address

	ThirdParty *ResponseSensorsCreateSensorTestTemplateV1ResponseSSIDsThirdParty `json:"thirdParty,omitempty"` //

	ID *int `json:"id,omitempty"` // Identification number

	WLANID *int `json:"wlanId,omitempty"` // WLAN ID

	Wlc string `json:"wlc,omitempty"` // WLC IP addres

	ValidFrom *int `json:"validFrom,omitempty"` // Valid From UTC timestamp

	ValidTo *int `json:"validTo,omitempty"` // Valid To UTC timestamp

	Status string `json:"status,omitempty"` // WLAN status: ENABLED or DISABLED

	ProxyServer string `json:"proxyServer,omitempty"` // Proxy server for onboarding SSID

	ProxyPort string `json:"proxyPort,omitempty"` // Proxy server port

	ProxyUserName string `json:"proxyUserName,omitempty"` // Proxy server user name

	ProxyPassword string `json:"proxyPassword,omitempty"` // Proxy server password

	AuthType string `json:"authType,omitempty"` // Authentication type: OPEN, WPA2_PSK, WPA2_EaP, WEB_AUTH, MAB, DOT1X, OTHER

	Psk string `json:"psk,omitempty"` // Password of SSID when passwordType is ASCII

	Username string `json:"username,omitempty"` // User name string for onboarding SSID

	Password string `json:"password,omitempty"` // Password string for onboarding SSID

	PasswordType string `json:"passwordType,omitempty"` // SSID password type: ASCII or HEX

	EapMethod string `json:"eapMethod,omitempty"` // WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC

	Scep *bool `json:"scep,omitempty"` // Secure certificate enrollment protocol: true or false or null for not applicable

	AuthProtocol string `json:"authProtocol,omitempty"` // Auth protocol

	Certfilename string `json:"certfilename,omitempty"` // Auth certificate file name

	Certxferprotocol string `json:"certxferprotocol,omitempty"` // Certificate transfering protocol: HTTP or HTTPS

	Certstatus string `json:"certstatus,omitempty"` // Certificate status: INACTIVE or ACTIVE

	Certpassphrase string `json:"certpassphrase,omitempty"` // Certificate password phrase

	Certdownloadurl string `json:"certdownloadurl,omitempty"` // Certificate download URL

	ExtWebAuthVirtualIP string `json:"extWebAuthVirtualIp,omitempty"` // External WEB Auth virtual IP

	ExtWebAuth *bool `json:"extWebAuth,omitempty"` // Indication of using external WEB Auth

	WhiteList *bool `json:"whiteList,omitempty"` // Indication of being on allowed list

	ExtWebAuthPortal string `json:"extWebAuthPortal,omitempty"` // External authentication portal

	ExtWebAuthAccessURL string `json:"extWebAuthAccessUrl,omitempty"` // External WEB Auth access URL

	ExtWebAuthHTMLTag *[]ResponseSensorsCreateSensorTestTemplateV1ResponseSSIDsExtWebAuthHTMLTag `json:"extWebAuthHtmlTag,omitempty"` //

	QosPolicy string `json:"qosPolicy,omitempty"` // QoS policy: PlATINUM, GOLD, SILVER, BRONZE

	Tests *[]ResponseSensorsCreateSensorTestTemplateV1ResponseSSIDsTests `json:"tests,omitempty"` //
}
type ResponseSensorsCreateSensorTestTemplateV1ResponseSSIDsThirdParty struct {
	Selected *bool `json:"selected,omitempty"` // true: the SSID is third party
}
type ResponseSensorsCreateSensorTestTemplateV1ResponseSSIDsExtWebAuthHTMLTag struct {
	Label string `json:"label,omitempty"` // Label

	Tag string `json:"tag,omitempty"` // Tag

	Value string `json:"value,omitempty"` // Value
}
type ResponseSensorsCreateSensorTestTemplateV1ResponseSSIDsTests struct {
	Name string `json:"name,omitempty"` // Name of the test

	Config *[]ResponseSensorsCreateSensorTestTemplateV1ResponseSSIDsTestsConfig `json:"config,omitempty"` //
}
type ResponseSensorsCreateSensorTestTemplateV1ResponseSSIDsTestsConfig struct {
	Domains []string `json:"domains,omitempty"` // DNS domain name

	Server string `json:"server,omitempty"` // Ping, file transfer, mail, radius, ssh, or telnet server

	UserName string `json:"userName,omitempty"` // User name

	Password string `json:"password,omitempty"` // Password

	URL string `json:"url,omitempty"` // URL

	Port *int `json:"port,omitempty"` // Radius or WEB server port

	Protocol string `json:"protocol,omitempty"` // Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)

	Servers []string `json:"servers,omitempty"` // IPerf server list

	Direction string `json:"direction,omitempty"` // IPerf direction (UPLOAD, DOWNLOAD, BOTH)

	StartPort *int `json:"startPort,omitempty"` // IPerf start port

	EndPort *int `json:"endPort,omitempty"` // IPerf end port

	UDPBandwidth *int `json:"udpBandwidth,omitempty"` // IPerf UDP bandwidth

	ProbeType string `json:"probeType,omitempty"` // Probe type

	NumPackets *int `json:"numPackets,omitempty"` // Number of packets

	PathToDownload string `json:"pathToDownload,omitempty"` // File path for file transfer

	TransferType string `json:"transferType,omitempty"` // File transfer type (UPLOAD, DOWNLOAD, BOTH)

	SharedSecret string `json:"sharedSecret,omitempty"` // Shared secret

	NdtServer string `json:"ndtServer,omitempty"` // NDT server

	NdtServerPort string `json:"ndtServerPort,omitempty"` // NDT server port

	NdtServerPath string `json:"ndtServerPath,omitempty"` // NDT server path

	UplinkTest *bool `json:"uplinkTest,omitempty"` // Uplink test

	DownlinkTest *bool `json:"downlinkTest,omitempty"` // Downlink test

	ProxyServer string `json:"proxyServer,omitempty"` // Proxy server

	ProxyPort string `json:"proxyPort,omitempty"` // Proxy port

	ProxyUserName string `json:"proxyUserName,omitempty"` // Proxy user name

	ProxyPassword string `json:"proxyPassword,omitempty"` // Proxy password

	UserNamePrompt string `json:"userNamePrompt,omitempty"` // User name prompt

	PasswordPrompt string `json:"passwordPrompt,omitempty"` // Password prompt

	ExitCommand string `json:"exitCommand,omitempty"` // Exit command

	FinalPrompt string `json:"finalPrompt,omitempty"` // Final prompt
}
type ResponseSensorsCreateSensorTestTemplateV1ResponseProfiles struct {
	AuthType string `json:"authType,omitempty"` // Authentication type: OPEN, WPA2_PSK, WPA2_EaP, WEB_AUTH, MAB, DOT1X, OTHER

	Psk string `json:"psk,omitempty"` // Password of SSID when passwordType is ASCII

	Username string `json:"username,omitempty"` // User name string for onboarding SSID

	Password string `json:"password,omitempty"` // Password string for onboarding SSID

	PasswordType string `json:"passwordType,omitempty"` // SSID password type: ASCII or HEX

	EapMethod string `json:"eapMethod,omitempty"` // WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC

	Scep *bool `json:"scep,omitempty"` // Secure certificate enrollment protocol: true or false or null for not applicable

	AuthProtocol string `json:"authProtocol,omitempty"` // Auth protocol

	Certfilename string `json:"certfilename,omitempty"` // Auth certificate file name

	Certxferprotocol string `json:"certxferprotocol,omitempty"` // Certificate transfering protocol: HTTP or HTTPS

	Certstatus string `json:"certstatus,omitempty"` // Certificate status: INACTIVE or ACTIVE

	Certpassphrase string `json:"certpassphrase,omitempty"` // Certificate password phrase

	Certdownloadurl string `json:"certdownloadurl,omitempty"` // Certificate download URL

	ExtWebAuthVirtualIP string `json:"extWebAuthVirtualIp,omitempty"` // External WEB Auth virtual IP

	ExtWebAuth *bool `json:"extWebAuth,omitempty"` // Indication of using external WEB Auth

	WhiteList *bool `json:"whiteList,omitempty"` // Indication of being on allowed list

	ExtWebAuthPortal string `json:"extWebAuthPortal,omitempty"` // External authentication portal

	ExtWebAuthAccessURL string `json:"extWebAuthAccessUrl,omitempty"` // External WEB Auth access URL

	ExtWebAuthHTMLTag *[]ResponseSensorsCreateSensorTestTemplateV1ResponseProfilesExtWebAuthHTMLTag `json:"extWebAuthHtmlTag,omitempty"` //

	QosPolicy string `json:"qosPolicy,omitempty"` // QoS policy: PlATINUM, GOLD, SILVER, BRONZE

	Tests *[]ResponseSensorsCreateSensorTestTemplateV1ResponseProfilesTests `json:"tests,omitempty"` //

	ProfileName string `json:"profileName,omitempty"` // Profile name

	DeviceType string `json:"deviceType,omitempty"` // Device Type

	VLAN string `json:"vlan,omitempty"` // VLAN

	LocationVLANList *[]ResponseSensorsCreateSensorTestTemplateV1ResponseProfilesLocationVLANList `json:"locationVlanList,omitempty"` //
}
type ResponseSensorsCreateSensorTestTemplateV1ResponseProfilesExtWebAuthHTMLTag struct {
	Label string `json:"label,omitempty"` // Label

	Tag string `json:"tag,omitempty"` // Tag

	Value string `json:"value,omitempty"` // Value
}
type ResponseSensorsCreateSensorTestTemplateV1ResponseProfilesTests struct {
	Name string `json:"name,omitempty"` // Name of the test

	Config *[]ResponseSensorsCreateSensorTestTemplateV1ResponseProfilesTestsConfig `json:"config,omitempty"` //
}
type ResponseSensorsCreateSensorTestTemplateV1ResponseProfilesTestsConfig struct {
	Domains []string `json:"domains,omitempty"` // DNS domain name

	Server string `json:"server,omitempty"` // Ping, file transfer, mail, radius, ssh, or telnet server

	UserName string `json:"userName,omitempty"` // User name

	Password string `json:"password,omitempty"` // Password

	URL string `json:"url,omitempty"` // URL

	Port *int `json:"port,omitempty"` // Radius or WEB server port

	Protocol string `json:"protocol,omitempty"` // Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)

	Servers []string `json:"servers,omitempty"` // IPerf server list

	Direction string `json:"direction,omitempty"` // IPerf direction (UPLOAD, DOWNLOAD, BOTH)

	StartPort *int `json:"startPort,omitempty"` // IPerf start port

	EndPort *int `json:"endPort,omitempty"` // IPerf end port

	UDPBandwidth *int `json:"udpBandwidth,omitempty"` // IPerf UDP bandwidth

	ProbeType string `json:"probeType,omitempty"` // Probe type

	NumPackets *int `json:"numPackets,omitempty"` // Number of packets

	PathToDownload string `json:"pathToDownload,omitempty"` // File path for file transfer

	TransferType string `json:"transferType,omitempty"` // File transfer type (UPLOAD, DOWNLOAD, BOTH)

	SharedSecret string `json:"sharedSecret,omitempty"` // Shared secret

	NdtServer string `json:"ndtServer,omitempty"` // NDT server

	NdtServerPort string `json:"ndtServerPort,omitempty"` // NDT server port

	NdtServerPath string `json:"ndtServerPath,omitempty"` // NDT server path

	UplinkTest *bool `json:"uplinkTest,omitempty"` // Uplink test

	DownlinkTest *bool `json:"downlinkTest,omitempty"` // Downlink test

	ProxyServer string `json:"proxyServer,omitempty"` // Proxy server

	ProxyPort string `json:"proxyPort,omitempty"` // Proxy port

	ProxyUserName string `json:"proxyUserName,omitempty"` // Proxy user name

	ProxyPassword string `json:"proxyPassword,omitempty"` // Proxy password

	UserNamePrompt string `json:"userNamePrompt,omitempty"` // User name prompt

	PasswordPrompt string `json:"passwordPrompt,omitempty"` // Password prompt

	ExitCommand string `json:"exitCommand,omitempty"` // Exit command

	FinalPrompt string `json:"finalPrompt,omitempty"` // Final prompt
}
type ResponseSensorsCreateSensorTestTemplateV1ResponseProfilesLocationVLANList struct {
	LocationID string `json:"locationId,omitempty"` // Site UUID

	VLANs []string `json:"vlans,omitempty"` // Array of VLANs
}
type ResponseSensorsCreateSensorTestTemplateV1ResponseLocationInfoList struct {
	LocationID string `json:"locationId,omitempty"` // Site UUID

	LocationType string `json:"locationType,omitempty"` // Site type

	AllSensors *bool `json:"allSensors,omitempty"` // Use all sensors in the site for test

	SiteHierarchy string `json:"siteHierarchy,omitempty"` // Site name hierarhy

	MacAddressList []string `json:"macAddressList,omitempty"` // MAC addresses

	ManagementVLAN string `json:"managementVlan,omitempty"` // Management VLAN

	CustomManagementVLAN *bool `json:"customManagementVlan,omitempty"` // Custom Management VLAN
}
type ResponseSensorsCreateSensorTestTemplateV1ResponseSensors struct {
	Name string `json:"name,omitempty"` // Sensor name

	MacAddress string `json:"macAddress,omitempty"` // MAC address

	SwitchMac string `json:"switchMac,omitempty"` // Switch MAC address

	SwitchUUID string `json:"switchUuid,omitempty"` // Switch device UUID

	SwitchSerialNumber string `json:"switchSerialNumber,omitempty"` // Switch serial number

	MarkedForUninstall *bool `json:"markedForUninstall,omitempty"` // Is marked for uninstall

	IPAddress string `json:"ipAddress,omitempty"` // IP address

	HostName string `json:"hostName,omitempty"` // Host name

	WiredApplicationStatus string `json:"wiredApplicationStatus,omitempty"` // Wired application status

	WiredApplicationMessage string `json:"wiredApplicationMessage,omitempty"` // Wired application message

	Assigned *bool `json:"assigned,omitempty"` // Is assigned

	Status string `json:"status,omitempty"` // Sensor device status: UP, DOWN, REBOOT

	XorSensor *bool `json:"xorSensor,omitempty"` // Is XOR sensor

	TargetAPs []string `json:"targetAPs,omitempty"` // Array of target APs

	RunNow string `json:"runNow,omitempty"` // Run now: YES, NO

	LocationID string `json:"locationId,omitempty"` // Site UUID

	AllSensorAddition *bool `json:"allSensorAddition,omitempty"` // Is all sensor addition

	ConfigUpdated string `json:"configUpdated,omitempty"` // Configuration updated: YES, NO

	SensorType string `json:"sensorType,omitempty"` // Sensor type

	TestMacAddresses *ResponseSensorsCreateSensorTestTemplateV1ResponseSensorsTestMacAddresses `json:"testMacAddresses,omitempty"` // A string-string test MAC address

	ID string `json:"id,omitempty"` // Sensor ID

	ServicePolicy string `json:"servicePolicy,omitempty"` // Service policy

	IPerfInfo *ResponseSensorsCreateSensorTestTemplateV1ResponseSensorsIPerfInfo `json:"iPerfInfo,omitempty"` // A string-stringList iPerf information
}
type ResponseSensorsCreateSensorTestTemplateV1ResponseSensorsTestMacAddresses interface{}
type ResponseSensorsCreateSensorTestTemplateV1ResponseSensorsIPerfInfo interface{}
type ResponseSensorsCreateSensorTestTemplateV1ResponseApCoverage struct {
	Bands string `json:"bands,omitempty"` // The WIFI bands

	NumberOfApsToTest *int `json:"numberOfApsToTest,omitempty"` // Number of APs to test

	RssiThreshold *int `json:"rssiThreshold,omitempty"` // RSSI threshold
}
type ResponseSensorsDeleteSensorTestV1 struct {
	Version string `json:"version,omitempty"` // Version

	Response *ResponseSensorsDeleteSensorTestV1Response `json:"response,omitempty"` //
}
type ResponseSensorsDeleteSensorTestV1Response struct {
	TemplateName string `json:"templateName,omitempty"` // Test template name to be delete

	Status string `json:"status,omitempty"` // Status of the DELETE action
}
type ResponseSensorsSensorsV1 struct {
	Version string `json:"version,omitempty"` // Version string of this API

	Response *[]ResponseSensorsSensorsV1Response `json:"response,omitempty"` //
}
type ResponseSensorsSensorsV1Response struct {
	Name string `json:"name,omitempty"` // The sensor device name

	Status string `json:"status,omitempty"` // Status of sensor device (REACHABLE, UNREACHABLE, DELETED, RUNNING, IDLE, UCLAIMED)

	RadioMacAddress string `json:"radioMacAddress,omitempty"` // Sensor device's radio MAC address

	EthernetMacAddress string `json:"ethernetMacAddress,omitempty"` // Sensor device's ethernet MAC address

	Location string `json:"location,omitempty"` // Site name in hierarchy form

	BackhaulType string `json:"backhaulType,omitempty"` // Backhall type: WIRED, WIRELESS

	SerialNumber string `json:"serialNumber,omitempty"` // Serial number

	IPAddress string `json:"ipAddress,omitempty"` // IP Address

	Version string `json:"version,omitempty"` // Sensor version

	LastSeen *int `json:"lastSeen,omitempty"` // Last seen timestamp

	Type string `json:"type,omitempty"` // Type

	SSH *ResponseSensorsSensorsV1ResponseSSH `json:"ssh,omitempty"` //

	Led *bool `json:"led,omitempty"` // Is LED Enabled
}
type ResponseSensorsSensorsV1ResponseSSH struct {
	SSHState string `json:"sshState,omitempty"` // SSH state

	SSHUserName string `json:"sshUserName,omitempty"` // SSH user name

	SSHPassword string `json:"sshPassword,omitempty"` // SSH password

	EnablePassword string `json:"enablePassword,omitempty"` // Enable password
}
type ResponseSensorsDuplicateSensorTestTemplateV1 struct {
	Version string `json:"version,omitempty"` // Version

	Response *ResponseSensorsDuplicateSensorTestTemplateV1Response `json:"response,omitempty"` //
}
type ResponseSensorsDuplicateSensorTestTemplateV1Response struct {
	Name   string `json:"name,omitempty"` // The sensor test template name
	TypeID string `json:"_id,omitempty"`  // The sensor test template unique identifier

	Version *int `json:"version,omitempty"` // The sensor test template version (must be 2)

	ModelVersion *int `json:"modelVersion,omitempty"` // Test template object model version (must be 2)

	StartTime *int `json:"startTime,omitempty"` // Start time

	LastModifiedTime *int `json:"lastModifiedTime,omitempty"` // Last modify time

	NumAssociatedSensor *int `json:"numAssociatedSensor,omitempty"` // Number of associated sensor

	Location string `json:"location,omitempty"` // Location string

	SiteHierarchy string `json:"siteHierarchy,omitempty"` // Site hierarchy

	Status string `json:"status,omitempty"` // Status of the test (RUNNING, NOTRUNNING)

	Connection string `json:"connection,omitempty"` // connection type of test: WIRED, WIRELESS, BOTH

	ActionInProgress string `json:"actionInProgress,omitempty"` // Indication of inprogress action

	Frequency *ResponseSensorsDuplicateSensorTestTemplateV1ResponseFrequency `json:"frequency,omitempty"` //

	RssiThreshold *int `json:"rssiThreshold,omitempty"` // RSSI threshold

	NumNeighborApThreshold *int `json:"numNeighborAPThreshold,omitempty"` // Number of neighboring AP threshold

	ScheduleInDays *int `json:"scheduleInDays,omitempty"` // Bit-wise value of scheduled test days

	WLANs []string `json:"wlans,omitempty"` // WLANs list

	SSIDs *[]ResponseSensorsDuplicateSensorTestTemplateV1ResponseSSIDs `json:"ssids,omitempty"` //

	Profiles *[]ResponseSensorsDuplicateSensorTestTemplateV1ResponseProfiles `json:"profiles,omitempty"` //

	TestScheduleMode string `json:"testScheduleMode,omitempty"` // Test schedule mode (ONDEMAND, DEDICATED, SCHEDULED, CONTINUOUS, RUNNOW)

	ShowWlcUpgradeBanner *bool `json:"showWlcUpgradeBanner,omitempty"` // Show WLC upgrade banner

	RadioAsSensorRemoved *bool `json:"radioAsSensorRemoved,omitempty"` // Radio as sensor removed

	EncryptionMode string `json:"encryptionMode,omitempty"` // Encryption mode

	RunNow string `json:"runNow,omitempty"` // Run now (YES, NO)

	LocationInfoList *[]ResponseSensorsDuplicateSensorTestTemplateV1ResponseLocationInfoList `json:"locationInfoList,omitempty"` //

	Sensors *[]ResponseSensorsDuplicateSensorTestTemplateV1ResponseSensors `json:"sensors,omitempty"` //

	ApCoverage *[]ResponseSensorsDuplicateSensorTestTemplateV1ResponseApCoverage `json:"apCoverage,omitempty"` //
}
type ResponseSensorsDuplicateSensorTestTemplateV1ResponseFrequency struct {
	Value *int `json:"value,omitempty"` // Value of the unit

	Unit string `json:"unit,omitempty"` // Unit of the time value (NANOSECONDS, MICROSECONDS, MILLISECONDS, SECONDS, MINUTES, HOURS, DAYS)
}
type ResponseSensorsDuplicateSensorTestTemplateV1ResponseSSIDs struct {
	Bands string `json:"bands,omitempty"` // WIFI bands: 2.4GHz or 5GHz

	SSID string `json:"ssid,omitempty"` // The SSID string

	ProfileName string `json:"profileName,omitempty"` // The SSID profile name string

	NumAps *int `json:"numAps,omitempty"` // Number of APs in the test

	NumSensors *int `json:"numSensors,omitempty"` // Number of Sensors in the test

	Layer3WebAuthsecurity string `json:"layer3webAuthsecurity,omitempty"` // Layer 3 WEB Auth security

	Layer3WebAuthuserName string `json:"layer3webAuthuserName,omitempty"` // Layer 3 WEB Auth user name

	Layer3WebAuthpassword string `json:"layer3webAuthpassword,omitempty"` // Layer 3 WEB Auth password

	Layer3WebAuthEmailAddress string `json:"layer3webAuthEmailAddress,omitempty"` // Layer 3 WEB Auth email address

	ThirdParty *ResponseSensorsDuplicateSensorTestTemplateV1ResponseSSIDsThirdParty `json:"thirdParty,omitempty"` //

	ID *int `json:"id,omitempty"` // Identification number

	WLANID *int `json:"wlanId,omitempty"` // WLAN ID

	Wlc string `json:"wlc,omitempty"` // WLC IP addres

	ValidFrom *int `json:"validFrom,omitempty"` // Valid From UTC timestamp

	ValidTo *int `json:"validTo,omitempty"` // Valid To UTC timestamp

	Status string `json:"status,omitempty"` // WLAN status: ENABLED or DISABLED

	ProxyServer string `json:"proxyServer,omitempty"` // Proxy server for onboarding SSID

	ProxyPort string `json:"proxyPort,omitempty"` // Proxy server port

	ProxyUserName string `json:"proxyUserName,omitempty"` // Proxy server user name

	ProxyPassword string `json:"proxyPassword,omitempty"` // Proxy server password

	AuthType string `json:"authType,omitempty"` // Authentication type: OPEN, WPA2_PSK, WPA2_EaP, WEB_AUTH, MAB, DOT1X, OTHER

	Psk string `json:"psk,omitempty"` // Password of SSID when passwordType is ASCII

	Username string `json:"username,omitempty"` // User name string for onboarding SSID

	Password string `json:"password,omitempty"` // Password string for onboarding SSID

	PasswordType string `json:"passwordType,omitempty"` // SSID password type: ASCII or HEX

	EapMethod string `json:"eapMethod,omitempty"` // WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC

	Scep *bool `json:"scep,omitempty"` // Secure certificate enrollment protocol: true or false or null for not applicable

	AuthProtocol string `json:"authProtocol,omitempty"` // Auth protocol

	Certfilename string `json:"certfilename,omitempty"` // Auth certificate file name

	Certxferprotocol string `json:"certxferprotocol,omitempty"` // Certificate transfering protocol: HTTP or HTTPS

	Certstatus string `json:"certstatus,omitempty"` // Certificate status: INACTIVE or ACTIVE

	Certpassphrase string `json:"certpassphrase,omitempty"` // Certificate password phrase

	Certdownloadurl string `json:"certdownloadurl,omitempty"` // Certificate download URL

	ExtWebAuthVirtualIP string `json:"extWebAuthVirtualIp,omitempty"` // External WEB Auth virtual IP

	ExtWebAuth *bool `json:"extWebAuth,omitempty"` // Indication of using external WEB Auth

	WhiteList *bool `json:"whiteList,omitempty"` // Indication of being on allowed list

	ExtWebAuthPortal string `json:"extWebAuthPortal,omitempty"` // External authentication portal

	ExtWebAuthAccessURL string `json:"extWebAuthAccessUrl,omitempty"` // External WEB Auth access URL

	ExtWebAuthHTMLTag *[]ResponseSensorsDuplicateSensorTestTemplateV1ResponseSSIDsExtWebAuthHTMLTag `json:"extWebAuthHtmlTag,omitempty"` //

	QosPolicy string `json:"qosPolicy,omitempty"` // QoS policy: PlATINUM, GOLD, SILVER, BRONZE

	Tests *[]ResponseSensorsDuplicateSensorTestTemplateV1ResponseSSIDsTests `json:"tests,omitempty"` //
}
type ResponseSensorsDuplicateSensorTestTemplateV1ResponseSSIDsThirdParty struct {
	Selected *bool `json:"selected,omitempty"` // true: the SSID is third party
}
type ResponseSensorsDuplicateSensorTestTemplateV1ResponseSSIDsExtWebAuthHTMLTag struct {
	Label string `json:"label,omitempty"` // Label

	Tag string `json:"tag,omitempty"` // Tag

	Value string `json:"value,omitempty"` // Value
}
type ResponseSensorsDuplicateSensorTestTemplateV1ResponseSSIDsTests struct {
	Name string `json:"name,omitempty"` // Name of the test

	Config *[]ResponseSensorsDuplicateSensorTestTemplateV1ResponseSSIDsTestsConfig `json:"config,omitempty"` //
}
type ResponseSensorsDuplicateSensorTestTemplateV1ResponseSSIDsTestsConfig struct {
	Domains []string `json:"domains,omitempty"` // DNS domain name

	Server string `json:"server,omitempty"` // Ping, file transfer, mail, radius, ssh, or telnet server

	UserName string `json:"userName,omitempty"` // User name

	Password string `json:"password,omitempty"` // Password

	URL string `json:"url,omitempty"` // URL

	Port *int `json:"port,omitempty"` // Radius or WEB server port

	Protocol string `json:"protocol,omitempty"` // Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)

	Servers []string `json:"servers,omitempty"` // IPerf server list

	Direction string `json:"direction,omitempty"` // IPerf direction (UPLOAD, DOWNLOAD, BOTH)

	StartPort *int `json:"startPort,omitempty"` // IPerf start port

	EndPort *int `json:"endPort,omitempty"` // IPerf end port

	UDPBandwidth *int `json:"udpBandwidth,omitempty"` // IPerf UDP bandwidth

	ProbeType string `json:"probeType,omitempty"` // Probe type

	NumPackets *int `json:"numPackets,omitempty"` // Number of packets

	PathToDownload string `json:"pathToDownload,omitempty"` // File path for file transfer

	TransferType string `json:"transferType,omitempty"` // File transfer type (UPLOAD, DOWNLOAD, BOTH)

	SharedSecret string `json:"sharedSecret,omitempty"` // Shared secret

	NdtServer string `json:"ndtServer,omitempty"` // NDT server

	NdtServerPort string `json:"ndtServerPort,omitempty"` // NDT server port

	NdtServerPath string `json:"ndtServerPath,omitempty"` // NDT server path

	UplinkTest *bool `json:"uplinkTest,omitempty"` // Uplink test

	DownlinkTest *bool `json:"downlinkTest,omitempty"` // Downlink test

	ProxyServer string `json:"proxyServer,omitempty"` // Proxy server

	ProxyPort string `json:"proxyPort,omitempty"` // Proxy port

	ProxyUserName string `json:"proxyUserName,omitempty"` // Proxy user name

	ProxyPassword string `json:"proxyPassword,omitempty"` // Proxy password

	UserNamePrompt string `json:"userNamePrompt,omitempty"` // User name prompt

	PasswordPrompt string `json:"passwordPrompt,omitempty"` // Password prompt

	ExitCommand string `json:"exitCommand,omitempty"` // Exit command

	FinalPrompt string `json:"finalPrompt,omitempty"` // Final prompt
}
type ResponseSensorsDuplicateSensorTestTemplateV1ResponseProfiles struct {
	AuthType string `json:"authType,omitempty"` // Authentication type: OPEN, WPA2_PSK, WPA2_EaP, WEB_AUTH, MAB, DOT1X, OTHER

	Psk string `json:"psk,omitempty"` // Password of SSID when passwordType is ASCII

	Username string `json:"username,omitempty"` // User name string for onboarding SSID

	Password string `json:"password,omitempty"` // Password string for onboarding SSID

	PasswordType string `json:"passwordType,omitempty"` // SSID password type: ASCII or HEX

	EapMethod string `json:"eapMethod,omitempty"` // WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC

	Scep *bool `json:"scep,omitempty"` // Secure certificate enrollment protocol: true or false or null for not applicable

	AuthProtocol string `json:"authProtocol,omitempty"` // Auth protocol

	Certfilename string `json:"certfilename,omitempty"` // Auth certificate file name

	Certxferprotocol string `json:"certxferprotocol,omitempty"` // Certificate transfering protocol: HTTP or HTTPS

	Certstatus string `json:"certstatus,omitempty"` // Certificate status: INACTIVE or ACTIVE

	Certpassphrase string `json:"certpassphrase,omitempty"` // Certificate password phrase

	Certdownloadurl string `json:"certdownloadurl,omitempty"` // Certificate download URL

	ExtWebAuthVirtualIP string `json:"extWebAuthVirtualIp,omitempty"` // External WEB Auth virtual IP

	ExtWebAuth *bool `json:"extWebAuth,omitempty"` // Indication of using external WEB Auth

	WhiteList *bool `json:"whiteList,omitempty"` // Indication of being on allowed list

	ExtWebAuthPortal string `json:"extWebAuthPortal,omitempty"` // External authentication portal

	ExtWebAuthAccessURL string `json:"extWebAuthAccessUrl,omitempty"` // External WEB Auth access URL

	ExtWebAuthHTMLTag *[]ResponseSensorsDuplicateSensorTestTemplateV1ResponseProfilesExtWebAuthHTMLTag `json:"extWebAuthHtmlTag,omitempty"` //

	QosPolicy string `json:"qosPolicy,omitempty"` // QoS policy: PlATINUM, GOLD, SILVER, BRONZE

	Tests *[]ResponseSensorsDuplicateSensorTestTemplateV1ResponseProfilesTests `json:"tests,omitempty"` //

	ProfileName string `json:"profileName,omitempty"` // Profile name

	DeviceType string `json:"deviceType,omitempty"` // Device Type

	VLAN string `json:"vlan,omitempty"` // VLAN

	LocationVLANList *[]ResponseSensorsDuplicateSensorTestTemplateV1ResponseProfilesLocationVLANList `json:"locationVlanList,omitempty"` //
}
type ResponseSensorsDuplicateSensorTestTemplateV1ResponseProfilesExtWebAuthHTMLTag struct {
	Label string `json:"label,omitempty"` // Label

	Tag string `json:"tag,omitempty"` // Tag

	Value string `json:"value,omitempty"` // Value
}
type ResponseSensorsDuplicateSensorTestTemplateV1ResponseProfilesTests struct {
	Name string `json:"name,omitempty"` // Name of the test

	Config *[]ResponseSensorsDuplicateSensorTestTemplateV1ResponseProfilesTestsConfig `json:"config,omitempty"` //
}
type ResponseSensorsDuplicateSensorTestTemplateV1ResponseProfilesTestsConfig struct {
	Domains []string `json:"domains,omitempty"` // DNS domain name

	Server string `json:"server,omitempty"` // Ping, file transfer, mail, radius, ssh, or telnet server

	UserName string `json:"userName,omitempty"` // User name

	Password string `json:"password,omitempty"` // Password

	URL string `json:"url,omitempty"` // URL

	Port *int `json:"port,omitempty"` // Radius or WEB server port

	Protocol string `json:"protocol,omitempty"` // Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)

	Servers []string `json:"servers,omitempty"` // IPerf server list

	Direction string `json:"direction,omitempty"` // IPerf direction (UPLOAD, DOWNLOAD, BOTH)

	StartPort *int `json:"startPort,omitempty"` // IPerf start port

	EndPort *int `json:"endPort,omitempty"` // IPerf end port

	UDPBandwidth *int `json:"udpBandwidth,omitempty"` // IPerf UDP bandwidth

	ProbeType string `json:"probeType,omitempty"` // Probe type

	NumPackets *int `json:"numPackets,omitempty"` // Number of packets

	PathToDownload string `json:"pathToDownload,omitempty"` // File path for file transfer

	TransferType string `json:"transferType,omitempty"` // File transfer type (UPLOAD, DOWNLOAD, BOTH)

	SharedSecret string `json:"sharedSecret,omitempty"` // Shared secret

	NdtServer string `json:"ndtServer,omitempty"` // NDT server

	NdtServerPort string `json:"ndtServerPort,omitempty"` // NDT server port

	NdtServerPath string `json:"ndtServerPath,omitempty"` // NDT server path

	UplinkTest *bool `json:"uplinkTest,omitempty"` // Uplink test

	DownlinkTest *bool `json:"downlinkTest,omitempty"` // Downlink test

	ProxyServer string `json:"proxyServer,omitempty"` // Proxy server

	ProxyPort string `json:"proxyPort,omitempty"` // Proxy port

	ProxyUserName string `json:"proxyUserName,omitempty"` // Proxy user name

	ProxyPassword string `json:"proxyPassword,omitempty"` // Proxy password

	UserNamePrompt string `json:"userNamePrompt,omitempty"` // User name prompt

	PasswordPrompt string `json:"passwordPrompt,omitempty"` // Password prompt

	ExitCommand string `json:"exitCommand,omitempty"` // Exit command

	FinalPrompt string `json:"finalPrompt,omitempty"` // Final prompt
}
type ResponseSensorsDuplicateSensorTestTemplateV1ResponseProfilesLocationVLANList struct {
	LocationID string `json:"locationId,omitempty"` // Site UUID

	VLANs []string `json:"vlans,omitempty"` // Array of VLANs
}
type ResponseSensorsDuplicateSensorTestTemplateV1ResponseLocationInfoList struct {
	LocationID string `json:"locationId,omitempty"` // Site UUID

	LocationType string `json:"locationType,omitempty"` // Site type

	AllSensors *bool `json:"allSensors,omitempty"` // Use all sensors in the site for test

	SiteHierarchy string `json:"siteHierarchy,omitempty"` // Site name hierarhy

	MacAddressList []string `json:"macAddressList,omitempty"` // MAC addresses

	ManagementVLAN string `json:"managementVlan,omitempty"` // Management VLAN

	CustomManagementVLAN *bool `json:"customManagementVlan,omitempty"` // Custom Management VLAN
}
type ResponseSensorsDuplicateSensorTestTemplateV1ResponseSensors struct {
	Name string `json:"name,omitempty"` // Sensor name

	MacAddress string `json:"macAddress,omitempty"` // MAC address

	SwitchMac string `json:"switchMac,omitempty"` // Switch MAC address

	SwitchUUID string `json:"switchUuid,omitempty"` // Switch device UUID

	SwitchSerialNumber string `json:"switchSerialNumber,omitempty"` // Switch serial number

	MarkedForUninstall *bool `json:"markedForUninstall,omitempty"` // Is marked for uninstall

	IPAddress string `json:"ipAddress,omitempty"` // IP address

	HostName string `json:"hostName,omitempty"` // Host name

	WiredApplicationStatus string `json:"wiredApplicationStatus,omitempty"` // Wired application status

	WiredApplicationMessage string `json:"wiredApplicationMessage,omitempty"` // Wired application message

	Assigned *bool `json:"assigned,omitempty"` // Is assigned

	Status string `json:"status,omitempty"` // Sensor device status: UP, DOWN, REBOOT

	XorSensor *bool `json:"xorSensor,omitempty"` // Is XOR sensor

	TargetAPs []string `json:"targetAPs,omitempty"` // Array of target APs

	RunNow string `json:"runNow,omitempty"` // Run now: YES, NO

	LocationID string `json:"locationId,omitempty"` // Site UUID

	AllSensorAddition *bool `json:"allSensorAddition,omitempty"` // Is all sensor addition

	ConfigUpdated string `json:"configUpdated,omitempty"` // Configuration updated: YES, NO

	SensorType string `json:"sensorType,omitempty"` // Sensor type

	TestMacAddresses *ResponseSensorsDuplicateSensorTestTemplateV1ResponseSensorsTestMacAddresses `json:"testMacAddresses,omitempty"` // A string-string test MAC address

	ID string `json:"id,omitempty"` // Sensor ID

	ServicePolicy string `json:"servicePolicy,omitempty"` // Service policy

	IPerfInfo *ResponseSensorsDuplicateSensorTestTemplateV1ResponseSensorsIPerfInfo `json:"iPerfInfo,omitempty"` // A string-stringList iPerf information
}
type ResponseSensorsDuplicateSensorTestTemplateV1ResponseSensorsTestMacAddresses interface{}
type ResponseSensorsDuplicateSensorTestTemplateV1ResponseSensorsIPerfInfo interface{}
type ResponseSensorsDuplicateSensorTestTemplateV1ResponseApCoverage struct {
	Bands string `json:"bands,omitempty"` // The WIFI bands

	NumberOfApsToTest *int `json:"numberOfApsToTest,omitempty"` // Number of APs to test

	RssiThreshold *int `json:"rssiThreshold,omitempty"` // RSSI threshold
}
type RequestSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1 struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	Filters *[]RequestSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1Filters `json:"filters,omitempty"` //

	Page *RequestSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1Page `json:"page,omitempty"` //
}
type RequestSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1Filters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	Value *int `json:"value,omitempty"` // Value
}
type RequestSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	TimeSortOrder string `json:"timeSortOrder,omitempty"` // Time Sort Order
}
type RequestSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1 struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	Filters *[]RequestSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1Filters `json:"filters,omitempty"` //

	Page *RequestSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1Page `json:"page,omitempty"` //
}
type RequestSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1Filters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	Value *int `json:"value,omitempty"` // Value
}
type RequestSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	TimeSortOrder string `json:"timeSortOrder,omitempty"` // Time Sort Order
}
type RequestSensorsEditSensorTestTemplateV1 struct {
	TemplateName string `json:"templateName,omitempty"` // The test template name that is to be edited

	Name   string `json:"name,omitempty"` // The sensor test template name, which is the same as in 'templateName'
	TypeID string `json:"_id,omitempty"`  // The sensor test template unique identifier, generated at test creation time

	Version *int `json:"version,omitempty"` // The sensor test template version (must be 2)

	ModelVersion *int `json:"modelVersion,omitempty"` // Test template object model version (must be 2)

	StartTime *int `json:"startTime,omitempty"` // Start time

	LastModifiedTime *int `json:"lastModifiedTime,omitempty"` // Last modify time

	NumAssociatedSensor *int `json:"numAssociatedSensor,omitempty"` // Number of associated sensor

	Location string `json:"location,omitempty"` // Location string

	SiteHierarchy string `json:"siteHierarchy,omitempty"` // Site hierarchy

	Status string `json:"status,omitempty"` // Status of the test (RUNNING, NOTRUNNING)

	Connection string `json:"connection,omitempty"` // connection type of test: WIRED, WIRELESS, BOTH

	ActionInProgress string `json:"actionInProgress,omitempty"` // Indication of inprogress action

	Frequency *RequestSensorsEditSensorTestTemplateV1Frequency `json:"frequency,omitempty"` //

	RssiThreshold *int `json:"rssiThreshold,omitempty"` // RSSI threshold

	NumNeighborApThreshold *int `json:"numNeighborAPThreshold,omitempty"` // Number of neighboring AP threshold

	ScheduleInDays *int `json:"scheduleInDays,omitempty"` // Bit-wise value of scheduled test days

	WLANs []string `json:"wlans,omitempty"` // WLANs list

	SSIDs *[]RequestSensorsEditSensorTestTemplateV1SSIDs `json:"ssids,omitempty"` //

	Profiles *[]RequestSensorsEditSensorTestTemplateV1Profiles `json:"profiles,omitempty"` //

	TestScheduleMode string `json:"testScheduleMode,omitempty"` // Test schedule mode (ONDEMAND, DEDICATED, SCHEDULED, CONTINUOUS, RUNNOW)

	ShowWlcUpgradeBanner *bool `json:"showWlcUpgradeBanner,omitempty"` // Show WLC upgrade banner

	RadioAsSensorRemoved *bool `json:"radioAsSensorRemoved,omitempty"` // Radio as sensor removed

	EncryptionMode string `json:"encryptionMode,omitempty"` // Encryption mode

	RunNow string `json:"runNow,omitempty"` // Run now (YES, NO)

	LocationInfoList *[]RequestSensorsEditSensorTestTemplateV1LocationInfoList `json:"locationInfoList,omitempty"` //

	Sensors *[]RequestSensorsEditSensorTestTemplateV1Sensors `json:"sensors,omitempty"` //

	ApCoverage *[]RequestSensorsEditSensorTestTemplateV1ApCoverage `json:"apCoverage,omitempty"` //
}
type RequestSensorsEditSensorTestTemplateV1Frequency struct {
	Value *int `json:"value,omitempty"` // Value of the unit

	Unit string `json:"unit,omitempty"` // Unit of the time value (NANOSECONDS, MICROSECONDS, MILLISECONDS, SECONDS, MINUTES, HOURS, DAYS)
}
type RequestSensorsEditSensorTestTemplateV1SSIDs struct {
	Bands string `json:"bands,omitempty"` // WIFI bands: 2.4GHz or 5GHz

	SSID string `json:"ssid,omitempty"` // The SSID string

	ProfileName string `json:"profileName,omitempty"` // The SSID profile name string

	NumAps *int `json:"numAps,omitempty"` // Number of APs in the test

	NumSensors *int `json:"numSensors,omitempty"` // Number of Sensors in the test

	Layer3WebAuthsecurity string `json:"layer3webAuthsecurity,omitempty"` // Layer 3 WEB Auth security

	Layer3WebAuthuserName string `json:"layer3webAuthuserName,omitempty"` // Layer 3 WEB Auth user name

	Layer3WebAuthpassword string `json:"layer3webAuthpassword,omitempty"` // Layer 3 WEB Auth password

	Layer3WebAuthEmailAddress string `json:"layer3webAuthEmailAddress,omitempty"` // Layer 3 WEB Auth email address

	ThirdParty *RequestSensorsEditSensorTestTemplateV1SSIDsThirdParty `json:"thirdParty,omitempty"` //

	ID *int `json:"id,omitempty"` // Identification number

	WLANID *int `json:"wlanId,omitempty"` // WLAN ID

	Wlc string `json:"wlc,omitempty"` // WLC IP addres

	ValidFrom *int `json:"validFrom,omitempty"` // Valid From UTC timestamp

	ValidTo *int `json:"validTo,omitempty"` // Valid To UTC timestamp

	Status string `json:"status,omitempty"` // WLAN status: ENABLED or DISABLED

	ProxyServer string `json:"proxyServer,omitempty"` // Proxy server for onboarding SSID

	ProxyPort string `json:"proxyPort,omitempty"` // Proxy server port

	ProxyUserName string `json:"proxyUserName,omitempty"` // Proxy server user name

	ProxyPassword string `json:"proxyPassword,omitempty"` // Proxy server password

	AuthType string `json:"authType,omitempty"` // Authentication type: OPEN, WPA2_PSK, WPA2_EAP, WEB_AUTH, MAB, DOT1X, OTHER

	Psk string `json:"psk,omitempty"` // Password of SSID when passwordType is ASCII

	Username string `json:"username,omitempty"` // User name string for onboarding SSID

	Password string `json:"password,omitempty"` // Password string for onboarding SSID

	PasswordType string `json:"passwordType,omitempty"` // SSID password type: ASCII or HEX

	EapMethod string `json:"eapMethod,omitempty"` // WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC

	Scep *bool `json:"scep,omitempty"` // Secure certificate enrollment protocol: true or false or null for not applicable

	AuthProtocol string `json:"authProtocol,omitempty"` // Auth protocol

	Certfilename string `json:"certfilename,omitempty"` // Auth certificate file name

	Certxferprotocol string `json:"certxferprotocol,omitempty"` // Certificate transfering protocol: HTTP or HTTPS

	Certstatus string `json:"certstatus,omitempty"` // Certificate status: INACTIVE or ACTIVE

	Certpassphrase string `json:"certpassphrase,omitempty"` // Certificate password phrase

	Certdownloadurl string `json:"certdownloadurl,omitempty"` // Certificate download URL

	ExtWebAuthVirtualIP string `json:"extWebAuthVirtualIp,omitempty"` // External WEB Auth virtual IP

	ExtWebAuth *bool `json:"extWebAuth,omitempty"` // Indication of using external WEB Auth

	WhiteList *bool `json:"whiteList,omitempty"` // Indication of being on allowed list

	ExtWebAuthPortal string `json:"extWebAuthPortal,omitempty"` // External authentication portal

	ExtWebAuthAccessURL string `json:"extWebAuthAccessUrl,omitempty"` // External WEB Auth access URL

	ExtWebAuthHTMLTag *[]RequestSensorsEditSensorTestTemplateV1SSIDsExtWebAuthHTMLTag `json:"extWebAuthHtmlTag,omitempty"` //

	QosPolicy string `json:"qosPolicy,omitempty"` // QoS policy: PlATINUM, GOLD, SILVER, BRONZE

	Tests *[]RequestSensorsEditSensorTestTemplateV1SSIDsTests `json:"tests,omitempty"` //
}
type RequestSensorsEditSensorTestTemplateV1SSIDsThirdParty struct {
	Selected *bool `json:"selected,omitempty"` // true: the SSID is third party
}
type RequestSensorsEditSensorTestTemplateV1SSIDsExtWebAuthHTMLTag struct {
	Label string `json:"label,omitempty"` // Label

	Tag string `json:"tag,omitempty"` // Tag

	Value string `json:"value,omitempty"` // Value
}
type RequestSensorsEditSensorTestTemplateV1SSIDsTests struct {
	Name string `json:"name,omitempty"` // Name of the test

	Config *[]RequestSensorsEditSensorTestTemplateV1SSIDsTestsConfig `json:"config,omitempty"` //
}
type RequestSensorsEditSensorTestTemplateV1SSIDsTestsConfig struct {
	Domains []string `json:"domains,omitempty"` // DNS domain name

	Server string `json:"server,omitempty"` // Ping, file transfer, mail, radius, ssh, or telnet server

	UserName string `json:"userName,omitempty"` // User name

	Password string `json:"password,omitempty"` // Password

	URL string `json:"url,omitempty"` // URL

	Port *int `json:"port,omitempty"` // Radius or WEB server port

	Protocol string `json:"protocol,omitempty"` // Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)

	Servers []string `json:"servers,omitempty"` // IPerf server list

	Direction string `json:"direction,omitempty"` // IPerf direction (UPLOAD, DOWNLOAD, BOTH)

	StartPort *int `json:"startPort,omitempty"` // IPerf start port

	EndPort *int `json:"endPort,omitempty"` // IPerf end port

	UDPBandwidth *int `json:"udpBandwidth,omitempty"` // IPerf UDP bandwidth

	ProbeType string `json:"probeType,omitempty"` // Probe type

	NumPackets *int `json:"numPackets,omitempty"` // Number of packets

	PathToDownload string `json:"pathToDownload,omitempty"` // File path for file transfer

	TransferType string `json:"transferType,omitempty"` // File transfer type (UPLOAD, DOWNLOAD, BOTH)

	SharedSecret string `json:"sharedSecret,omitempty"` // Shared secret

	NdtServer string `json:"ndtServer,omitempty"` // NDT server

	NdtServerPort string `json:"ndtServerPort,omitempty"` // NDT server port

	NdtServerPath string `json:"ndtServerPath,omitempty"` // NDT server path

	UplinkTest *bool `json:"uplinkTest,omitempty"` // Uplink test

	DownlinkTest *bool `json:"downlinkTest,omitempty"` // Downlink test

	ProxyServer string `json:"proxyServer,omitempty"` // Proxy server

	ProxyPort string `json:"proxyPort,omitempty"` // Proxy port

	ProxyUserName string `json:"proxyUserName,omitempty"` // Proxy user name

	ProxyPassword string `json:"proxyPassword,omitempty"` // Proxy password

	UserNamePrompt string `json:"userNamePrompt,omitempty"` // User name prompt

	PasswordPrompt string `json:"passwordPrompt,omitempty"` // Password prompt

	ExitCommand string `json:"exitCommand,omitempty"` // Exit command

	FinalPrompt string `json:"finalPrompt,omitempty"` // Final prompt
}
type RequestSensorsEditSensorTestTemplateV1Profiles struct {
	AuthType string `json:"authType,omitempty"` // Authentication type: OPEN, WPA2_PSK, WPA2_EAP, WEB_AUTH, MAB, DOT1X, OTHER

	Psk string `json:"psk,omitempty"` // Password of SSID when passwordType is ASCII

	Username string `json:"username,omitempty"` // User name string for onboarding SSID

	Password string `json:"password,omitempty"` // Password string for onboarding SSID

	PasswordType string `json:"passwordType,omitempty"` // SSID password type: ASCII or HEX

	EapMethod string `json:"eapMethod,omitempty"` // WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC

	Scep *bool `json:"scep,omitempty"` // Secure certificate enrollment protocol: true or false or null for not applicable

	AuthProtocol string `json:"authProtocol,omitempty"` // Auth protocol

	Certfilename string `json:"certfilename,omitempty"` // Auth certificate file name

	Certxferprotocol string `json:"certxferprotocol,omitempty"` // Certificate transfering protocol: HTTP or HTTPS

	Certstatus string `json:"certstatus,omitempty"` // Certificate status: INACTIVE or ACTIVE

	Certpassphrase string `json:"certpassphrase,omitempty"` // Certificate password phrase

	Certdownloadurl string `json:"certdownloadurl,omitempty"` // Certificate download URL

	ExtWebAuthVirtualIP string `json:"extWebAuthVirtualIp,omitempty"` // External WEB Auth virtual IP

	ExtWebAuth *bool `json:"extWebAuth,omitempty"` // Indication of using external WEB Auth

	WhiteList *bool `json:"whiteList,omitempty"` // Indication of being on allowed list

	ExtWebAuthPortal string `json:"extWebAuthPortal,omitempty"` // External authentication portal

	ExtWebAuthAccessURL string `json:"extWebAuthAccessUrl,omitempty"` // External WEB Auth access URL

	ExtWebAuthHTMLTag *[]RequestSensorsEditSensorTestTemplateV1ProfilesExtWebAuthHTMLTag `json:"extWebAuthHtmlTag,omitempty"` //

	QosPolicy string `json:"qosPolicy,omitempty"` // QoS policy: PlATINUM, GOLD, SILVER, BRONZE

	Tests *[]RequestSensorsEditSensorTestTemplateV1ProfilesTests `json:"tests,omitempty"` //

	ProfileName string `json:"profileName,omitempty"` // Profile name

	DeviceType string `json:"deviceType,omitempty"` // Device Type

	VLAN string `json:"vlan,omitempty"` // VLAN

	LocationVLANList *[]RequestSensorsEditSensorTestTemplateV1ProfilesLocationVLANList `json:"locationVlanList,omitempty"` //
}
type RequestSensorsEditSensorTestTemplateV1ProfilesExtWebAuthHTMLTag struct {
	Label string `json:"label,omitempty"` // Label

	Tag string `json:"tag,omitempty"` // Tag

	Value string `json:"value,omitempty"` // Value
}
type RequestSensorsEditSensorTestTemplateV1ProfilesTests struct {
	Name string `json:"name,omitempty"` // Name of the test

	Config *[]RequestSensorsEditSensorTestTemplateV1ProfilesTestsConfig `json:"config,omitempty"` //
}
type RequestSensorsEditSensorTestTemplateV1ProfilesTestsConfig struct {
	Domains []string `json:"domains,omitempty"` // DNS domain name

	Server string `json:"server,omitempty"` // Ping, file transfer, mail, radius, ssh, or telnet server

	UserName string `json:"userName,omitempty"` // User name

	Password string `json:"password,omitempty"` // Password

	URL string `json:"url,omitempty"` // URL

	Port *int `json:"port,omitempty"` // Radius or WEB server port

	Protocol string `json:"protocol,omitempty"` // Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)

	Servers []string `json:"servers,omitempty"` // IPerf server list

	Direction string `json:"direction,omitempty"` // IPerf direction (UPLOAD, DOWNLOAD, BOTH)

	StartPort *int `json:"startPort,omitempty"` // IPerf start port

	EndPort *int `json:"endPort,omitempty"` // IPerf end port

	UDPBandwidth *int `json:"udpBandwidth,omitempty"` // IPerf UDP bandwidth

	ProbeType string `json:"probeType,omitempty"` // Probe type

	NumPackets *int `json:"numPackets,omitempty"` // Number of packets

	PathToDownload string `json:"pathToDownload,omitempty"` // File path for file transfer

	TransferType string `json:"transferType,omitempty"` // File transfer type (UPLOAD, DOWNLOAD, BOTH)

	SharedSecret string `json:"sharedSecret,omitempty"` // Shared secret

	NdtServer string `json:"ndtServer,omitempty"` // NDT server

	NdtServerPort string `json:"ndtServerPort,omitempty"` // NDT server port

	NdtServerPath string `json:"ndtServerPath,omitempty"` // NDT server path

	UplinkTest *bool `json:"uplinkTest,omitempty"` // Uplink test

	DownlinkTest *bool `json:"downlinkTest,omitempty"` // Downlink test

	ProxyServer string `json:"proxyServer,omitempty"` // Proxy server

	ProxyPort string `json:"proxyPort,omitempty"` // Proxy port

	ProxyUserName string `json:"proxyUserName,omitempty"` // Proxy user name

	ProxyPassword string `json:"proxyPassword,omitempty"` // Proxy password

	UserNamePrompt string `json:"userNamePrompt,omitempty"` // User name prompt

	PasswordPrompt string `json:"passwordPrompt,omitempty"` // Password prompt

	ExitCommand string `json:"exitCommand,omitempty"` // Exit command

	FinalPrompt string `json:"finalPrompt,omitempty"` // Final prompt
}
type RequestSensorsEditSensorTestTemplateV1ProfilesLocationVLANList struct {
	LocationID string `json:"locationId,omitempty"` // Site UUID

	VLANs []string `json:"vlans,omitempty"` // Array of VLANs
}
type RequestSensorsEditSensorTestTemplateV1LocationInfoList struct {
	LocationID string `json:"locationId,omitempty"` // Site UUID

	LocationType string `json:"locationType,omitempty"` // Site type

	AllSensors *bool `json:"allSensors,omitempty"` // Use all sensors in the site for test

	SiteHierarchy string `json:"siteHierarchy,omitempty"` // Site name hierarhy

	MacAddressList []string `json:"macAddressList,omitempty"` // MAC addresses

	ManagementVLAN string `json:"managementVlan,omitempty"` // Management VLAN

	CustomManagementVLAN *bool `json:"customManagementVlan,omitempty"` // Custom Management VLAN
}
type RequestSensorsEditSensorTestTemplateV1Sensors struct {
	Name string `json:"name,omitempty"` // Sensor name

	MacAddress string `json:"macAddress,omitempty"` // MAC address

	SwitchMac string `json:"switchMac,omitempty"` // Switch MAC address

	SwitchUUID string `json:"switchUuid,omitempty"` // Switch device UUID

	SwitchSerialNumber string `json:"switchSerialNumber,omitempty"` // Switch serial number

	MarkedForUninstall *bool `json:"markedForUninstall,omitempty"` // Is marked for uninstall

	IPAddress string `json:"ipAddress,omitempty"` // IP address

	HostName string `json:"hostName,omitempty"` // Host name

	WiredApplicationStatus string `json:"wiredApplicationStatus,omitempty"` // Wired application status

	WiredApplicationMessage string `json:"wiredApplicationMessage,omitempty"` // Wired application message

	Assigned *bool `json:"assigned,omitempty"` // Is assigned

	Status string `json:"status,omitempty"` // Sensor device status: UP, DOWN, REBOOT

	XorSensor *bool `json:"xorSensor,omitempty"` // Is XOR sensor

	TargetAPs []string `json:"targetAPs,omitempty"` // Array of target APs

	RunNow string `json:"runNow,omitempty"` // Run now: YES, NO

	LocationID string `json:"locationId,omitempty"` // Site UUID

	AllSensorAddition *bool `json:"allSensorAddition,omitempty"` // Is all sensor addition

	ConfigUpdated string `json:"configUpdated,omitempty"` // Configuration updated: YES, NO

	SensorType string `json:"sensorType,omitempty"` // Sensor type

	TestMacAddresses *RequestSensorsEditSensorTestTemplateV1SensorsTestMacAddresses `json:"testMacAddresses,omitempty"` // A string-string test MAC address

	ID string `json:"id,omitempty"` // Sensor ID

	ServicePolicy string `json:"servicePolicy,omitempty"` // Service policy

	IPerfInfo *RequestSensorsEditSensorTestTemplateV1SensorsIPerfInfo `json:"iPerfInfo,omitempty"` // A string-stringList iPerf information
}
type RequestSensorsEditSensorTestTemplateV1SensorsTestMacAddresses interface{}
type RequestSensorsEditSensorTestTemplateV1SensorsIPerfInfo interface{}
type RequestSensorsEditSensorTestTemplateV1ApCoverage struct {
	Bands string `json:"bands,omitempty"` // The WIFI bands

	NumberOfApsToTest *int `json:"numberOfApsToTest,omitempty"` // Number of APs to test

	RssiThreshold *int `json:"rssiThreshold,omitempty"` // RSSI threshold
}
type RequestSensorsCreatesAnICapConfigurationIntentForPreviewApproveV1 []RequestItemSensorsCreatesAnICapConfigurationIntentForPreviewApproveV1 // Array of RequestSensorsCreatesAnICAPConfigurationIntentForPreviewApproveV1
type RequestItemSensorsCreatesAnICapConfigurationIntentForPreviewApproveV1 struct {
	CaptureType string `json:"captureType,omitempty"` // Capture Type

	DurationInMins *int `json:"durationInMins,omitempty"` // Duration In Mins

	ClientMac string `json:"clientMac,omitempty"` // Client Mac

	WlcID string `json:"wlcId,omitempty"` // Wlc Id

	APID string `json:"apId,omitempty"` // Ap Id

	Slot *[]float64 `json:"slot,omitempty"` // Slot

	OtaBand string `json:"otaBand,omitempty"` // Ota Band

	OtaChannel *int `json:"otaChannel,omitempty"` // Ota Channel

	OtaChannelWidth *int `json:"otaChannelWidth,omitempty"` // Ota Channel Width
}
type RequestSensorsCreatesAICapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDeviceV1 struct {
	object string `json:"object,omitempty"` // object
}
type RequestSensorsDeploysTheICapConfigurationIntentByActivityIDV1 struct {
	object string `json:"object,omitempty"` // object
}
type RequestSensorsGeneratesTheDevicesClisOfTheICapConfigurationIntentV1 map[string]interface{}
type RequestSensorsDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1 []RequestItemSensorsDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1 // Array of RequestSensorsDeploysTheGivenICAPConfigurationIntentWithoutPreviewAndApproveV1
type RequestItemSensorsDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1 struct {
	CaptureType string `json:"captureType,omitempty"` // Capture Type

	DurationInMins *int `json:"durationInMins,omitempty"` // Duration In Mins

	ClientMac string `json:"clientMac,omitempty"` // Client Mac

	WlcID string `json:"wlcId,omitempty"` // Wlc Id

	APID string `json:"apId,omitempty"` // Ap Id

	Slot *[]float64 `json:"slot,omitempty"` // Slot

	OtaBand string `json:"otaBand,omitempty"` // Ota Band

	OtaChannel *int `json:"otaChannel,omitempty"` // Ota Channel

	OtaChannelWidth *int `json:"otaChannelWidth,omitempty"` // Ota Channel Width
}
type RequestSensorsRemoveTheICapConfigurationOnTheDeviceWithoutPreviewV1 struct {
	object string `json:"object,omitempty"` // object
}
type RequestSensorsCreateSensorTestTemplateV1 struct {
	Name string `json:"name,omitempty"` // The sensor test template name

	Version *int `json:"version,omitempty"` // The sensor test template version (must be 2)

	ModelVersion *int `json:"modelVersion,omitempty"` // Test template object model version (must be 2)

	Connection string `json:"connection,omitempty"` // connection type of test: WIRED, WIRELESS, BOTH

	SSIDs *[]RequestSensorsCreateSensorTestTemplateV1SSIDs `json:"ssids,omitempty"` //

	Profiles *[]RequestSensorsCreateSensorTestTemplateV1Profiles `json:"profiles,omitempty"` //

	EncryptionMode string `json:"encryptionMode,omitempty"` // Encryption mode

	RunNow string `json:"runNow,omitempty"` // Run now (YES, NO)

	LocationInfoList *[]RequestSensorsCreateSensorTestTemplateV1LocationInfoList `json:"locationInfoList,omitempty"` //

	Sensors *[]RequestSensorsCreateSensorTestTemplateV1Sensors `json:"sensors,omitempty"` //

	ApCoverage *[]RequestSensorsCreateSensorTestTemplateV1ApCoverage `json:"apCoverage,omitempty"` //
}
type RequestSensorsCreateSensorTestTemplateV1SSIDs struct {
	Bands string `json:"bands,omitempty"` // WIFI bands: 2.4GHz or 5GHz

	SSID string `json:"ssid,omitempty"` // The SSID string

	ProfileName string `json:"profileName,omitempty"` // The SSID profile name string

	Layer3WebAuthsecurity string `json:"layer3webAuthsecurity,omitempty"` // Layer 3 WEB Auth security

	Layer3WebAuthuserName string `json:"layer3webAuthuserName,omitempty"` // Layer 3 WEB Auth user name

	Layer3WebAuthpassword string `json:"layer3webAuthpassword,omitempty"` // Layer 3 WEB Auth password

	Layer3WebAuthEmailAddress string `json:"layer3webAuthEmailAddress,omitempty"` // Layer 3 WEB Auth email address

	ThirdParty *RequestSensorsCreateSensorTestTemplateV1SSIDsThirdParty `json:"thirdParty,omitempty"` //

	WLANID *int `json:"wlanId,omitempty"` // WLAN ID

	Wlc string `json:"wlc,omitempty"` // WLC IP addres

	ProxyServer string `json:"proxyServer,omitempty"` // Proxy server for onboarding SSID

	ProxyPort string `json:"proxyPort,omitempty"` // Proxy server port

	ProxyUserName string `json:"proxyUserName,omitempty"` // Proxy server user name

	ProxyPassword string `json:"proxyPassword,omitempty"` // Proxy server password

	AuthType string `json:"authType,omitempty"` // Authentication type: OPEN, WPA2_PSK, WPA2_EaP, WEB_AUTH, MAB, DOT1X, OTHER

	Psk string `json:"psk,omitempty"` // Password of SSID when passwordType is ASCII

	Username string `json:"username,omitempty"` // User name string for onboarding SSID

	Password string `json:"password,omitempty"` // Password string for onboarding SSID

	PasswordType string `json:"passwordType,omitempty"` // SSID password type: ASCII or HEX

	EapMethod string `json:"eapMethod,omitempty"` // WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC

	Scep *bool `json:"scep,omitempty"` // Secure certificate enrollment protocol: true or false or null for not applicable

	AuthProtocol string `json:"authProtocol,omitempty"` // Auth protocol

	Certfilename string `json:"certfilename,omitempty"` // Auth certificate file name

	Certxferprotocol string `json:"certxferprotocol,omitempty"` // Certificate transfering protocol: HTTP or HTTPS

	Certstatus string `json:"certstatus,omitempty"` // Certificate status: INACTIVE or ACTIVE

	Certpassphrase string `json:"certpassphrase,omitempty"` // Certificate password phrase

	Certdownloadurl string `json:"certdownloadurl,omitempty"` // Certificate download URL

	ExtWebAuthVirtualIP string `json:"extWebAuthVirtualIp,omitempty"` // External WEB Auth virtual IP

	ExtWebAuth *bool `json:"extWebAuth,omitempty"` // Indication of using external WEB Auth

	WhiteList *bool `json:"whiteList,omitempty"` // Indication of being on allowed list

	ExtWebAuthPortal string `json:"extWebAuthPortal,omitempty"` // External authentication portal

	ExtWebAuthAccessURL string `json:"extWebAuthAccessUrl,omitempty"` // External WEB Auth access URL

	ExtWebAuthHTMLTag *[]RequestSensorsCreateSensorTestTemplateV1SSIDsExtWebAuthHTMLTag `json:"extWebAuthHtmlTag,omitempty"` //

	QosPolicy string `json:"qosPolicy,omitempty"` // QoS policy: PlATINUM, GOLD, SILVER, BRONZE

	Tests *[]RequestSensorsCreateSensorTestTemplateV1SSIDsTests `json:"tests,omitempty"` //
}
type RequestSensorsCreateSensorTestTemplateV1SSIDsThirdParty struct {
	Selected *bool `json:"selected,omitempty"` // true: the SSID is third party
}
type RequestSensorsCreateSensorTestTemplateV1SSIDsExtWebAuthHTMLTag struct {
	Label string `json:"label,omitempty"` // Label

	Tag string `json:"tag,omitempty"` // Tag

	Value string `json:"value,omitempty"` // Value
}
type RequestSensorsCreateSensorTestTemplateV1SSIDsTests struct {
	Name string `json:"name,omitempty"` // Name of the test

	Config *[]RequestSensorsCreateSensorTestTemplateV1SSIDsTestsConfig `json:"config,omitempty"` //
}
type RequestSensorsCreateSensorTestTemplateV1SSIDsTestsConfig struct {
	Domains []string `json:"domains,omitempty"` // DNS domain name

	Server string `json:"server,omitempty"` // Ping, file transfer, mail, radius, ssh, or telnet server

	UserName string `json:"userName,omitempty"` // User name

	Password string `json:"password,omitempty"` // Password

	URL string `json:"url,omitempty"` // URL

	Port *int `json:"port,omitempty"` // Radius or WEB server port

	Protocol string `json:"protocol,omitempty"` // Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)

	Servers []string `json:"servers,omitempty"` // IPerf server list

	Direction string `json:"direction,omitempty"` // IPerf direction (UPLOAD, DOWNLOAD, BOTH)

	StartPort *int `json:"startPort,omitempty"` // IPerf start port

	EndPort *int `json:"endPort,omitempty"` // IPerf end port

	UDPBandwidth *int `json:"udpBandwidth,omitempty"` // IPerf UDP bandwidth

	ProbeType string `json:"probeType,omitempty"` // Probe type

	NumPackets *int `json:"numPackets,omitempty"` // Number of packets

	PathToDownload string `json:"pathToDownload,omitempty"` // File path for file transfer

	TransferType string `json:"transferType,omitempty"` // File transfer type (UPLOAD, DOWNLOAD, BOTH)

	SharedSecret string `json:"sharedSecret,omitempty"` // Shared secret

	NdtServer string `json:"ndtServer,omitempty"` // NDT server

	NdtServerPort string `json:"ndtServerPort,omitempty"` // NDT server port

	NdtServerPath string `json:"ndtServerPath,omitempty"` // NDT server path

	UplinkTest *bool `json:"uplinkTest,omitempty"` // Uplink test

	DownlinkTest *bool `json:"downlinkTest,omitempty"` // Downlink test

	ProxyServer string `json:"proxyServer,omitempty"` // Proxy server

	ProxyPort string `json:"proxyPort,omitempty"` // Proxy port

	ProxyUserName string `json:"proxyUserName,omitempty"` // Proxy user name

	ProxyPassword string `json:"proxyPassword,omitempty"` // Proxy password

	UserNamePrompt string `json:"userNamePrompt,omitempty"` // User name prompt

	PasswordPrompt string `json:"passwordPrompt,omitempty"` // Password prompt

	ExitCommand string `json:"exitCommand,omitempty"` // Exit command

	FinalPrompt string `json:"finalPrompt,omitempty"` // Final prompt
}
type RequestSensorsCreateSensorTestTemplateV1Profiles struct {
	AuthType string `json:"authType,omitempty"` // Authentication type: OPEN, WPA2_PSK, WPA2_EaP, WEB_AUTH, MAB, DOT1X, OTHER

	Psk string `json:"psk,omitempty"` // Password of SSID when passwordType is ASCII

	Username string `json:"username,omitempty"` // User name string for onboarding SSID

	Password string `json:"password,omitempty"` // Password string for onboarding SSID

	PasswordType string `json:"passwordType,omitempty"` // SSID password type: ASCII or HEX

	EapMethod string `json:"eapMethod,omitempty"` // WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC

	Scep *bool `json:"scep,omitempty"` // Secure certificate enrollment protocol: true or false or null for not applicable

	AuthProtocol string `json:"authProtocol,omitempty"` // Auth protocol

	Certfilename string `json:"certfilename,omitempty"` // Auth certificate file name

	Certxferprotocol string `json:"certxferprotocol,omitempty"` // Certificate transfering protocol: HTTP or HTTPS

	Certstatus string `json:"certstatus,omitempty"` // Certificate status: INACTIVE or ACTIVE

	Certpassphrase string `json:"certpassphrase,omitempty"` // Certificate password phrase

	Certdownloadurl string `json:"certdownloadurl,omitempty"` // Certificate download URL

	ExtWebAuthVirtualIP string `json:"extWebAuthVirtualIp,omitempty"` // External WEB Auth virtual IP

	ExtWebAuth *bool `json:"extWebAuth,omitempty"` // Indication of using external WEB Auth

	WhiteList *bool `json:"whiteList,omitempty"` // Indication of being on allowed list

	ExtWebAuthPortal string `json:"extWebAuthPortal,omitempty"` // External authentication portal

	ExtWebAuthAccessURL string `json:"extWebAuthAccessUrl,omitempty"` // External WEB Auth access URL

	ExtWebAuthHTMLTag *[]RequestSensorsCreateSensorTestTemplateV1ProfilesExtWebAuthHTMLTag `json:"extWebAuthHtmlTag,omitempty"` //

	QosPolicy string `json:"qosPolicy,omitempty"` // QoS policy: PlATINUM, GOLD, SILVER, BRONZE

	Tests *[]RequestSensorsCreateSensorTestTemplateV1ProfilesTests `json:"tests,omitempty"` //

	ProfileName string `json:"profileName,omitempty"` // Profile name

	DeviceType string `json:"deviceType,omitempty"` // Device Type

	VLAN string `json:"vlan,omitempty"` // VLAN

	LocationVLANList *[]RequestSensorsCreateSensorTestTemplateV1ProfilesLocationVLANList `json:"locationVlanList,omitempty"` //
}
type RequestSensorsCreateSensorTestTemplateV1ProfilesExtWebAuthHTMLTag struct {
	Label string `json:"label,omitempty"` // Label

	Tag string `json:"tag,omitempty"` // Tag

	Value string `json:"value,omitempty"` // Value
}
type RequestSensorsCreateSensorTestTemplateV1ProfilesTests struct {
	Name string `json:"name,omitempty"` // Name of the test

	Config *[]RequestSensorsCreateSensorTestTemplateV1ProfilesTestsConfig `json:"config,omitempty"` //
}
type RequestSensorsCreateSensorTestTemplateV1ProfilesTestsConfig struct {
	Domains []string `json:"domains,omitempty"` // DNS domain name

	Server string `json:"server,omitempty"` // Ping, file transfer, mail, radius, ssh, or telnet server

	UserName string `json:"userName,omitempty"` // User name

	Password string `json:"password,omitempty"` // Password

	URL string `json:"url,omitempty"` // URL

	Port *int `json:"port,omitempty"` // Radius or WEB server port

	Protocol string `json:"protocol,omitempty"` // Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)

	Servers []string `json:"servers,omitempty"` // IPerf server list

	Direction string `json:"direction,omitempty"` // IPerf direction (UPLOAD, DOWNLOAD, BOTH)

	StartPort *int `json:"startPort,omitempty"` // IPerf start port

	EndPort *int `json:"endPort,omitempty"` // IPerf end port

	UDPBandwidth *int `json:"udpBandwidth,omitempty"` // IPerf UDP bandwidth

	ProbeType string `json:"probeType,omitempty"` // Probe type

	NumPackets *int `json:"numPackets,omitempty"` // Number of packets

	PathToDownload string `json:"pathToDownload,omitempty"` // File path for file transfer

	TransferType string `json:"transferType,omitempty"` // File transfer type (UPLOAD, DOWNLOAD, BOTH)

	SharedSecret string `json:"sharedSecret,omitempty"` // Shared secret

	NdtServer string `json:"ndtServer,omitempty"` // NDT server

	NdtServerPort string `json:"ndtServerPort,omitempty"` // NDT server port

	NdtServerPath string `json:"ndtServerPath,omitempty"` // NDT server path

	UplinkTest *bool `json:"uplinkTest,omitempty"` // Uplink test

	DownlinkTest *bool `json:"downlinkTest,omitempty"` // Downlink test

	ProxyServer string `json:"proxyServer,omitempty"` // Proxy server

	ProxyPort string `json:"proxyPort,omitempty"` // Proxy port

	ProxyUserName string `json:"proxyUserName,omitempty"` // Proxy user name

	ProxyPassword string `json:"proxyPassword,omitempty"` // Proxy password

	UserNamePrompt string `json:"userNamePrompt,omitempty"` // User name prompt

	PasswordPrompt string `json:"passwordPrompt,omitempty"` // Password prompt

	ExitCommand string `json:"exitCommand,omitempty"` // Exit command

	FinalPrompt string `json:"finalPrompt,omitempty"` // Final prompt
}
type RequestSensorsCreateSensorTestTemplateV1ProfilesLocationVLANList struct {
	LocationID string `json:"locationId,omitempty"` // Site UUID

	VLANs []string `json:"vlans,omitempty"` // Array of VLANs
}
type RequestSensorsCreateSensorTestTemplateV1LocationInfoList struct {
	LocationID string `json:"locationId,omitempty"` // Site UUID

	LocationType string `json:"locationType,omitempty"` // Site type

	AllSensors *bool `json:"allSensors,omitempty"` // Use all sensors in the site for test

	SiteHierarchy string `json:"siteHierarchy,omitempty"` // Site name hierarhy

	MacAddressList []string `json:"macAddressList,omitempty"` // MAC addresses

	ManagementVLAN string `json:"managementVlan,omitempty"` // Management VLAN

	CustomManagementVLAN *bool `json:"customManagementVlan,omitempty"` // Custom Management VLAN
}
type RequestSensorsCreateSensorTestTemplateV1Sensors struct {
	Name string `json:"name,omitempty"` // Sensor name

	MacAddress string `json:"macAddress,omitempty"` // MAC address

	SwitchMac string `json:"switchMac,omitempty"` // Switch MAC address

	SwitchUUID string `json:"switchUuid,omitempty"` // Switch device UUID

	SwitchSerialNumber string `json:"switchSerialNumber,omitempty"` // Switch serial number

	MarkedForUninstall *bool `json:"markedForUninstall,omitempty"` // Is marked for uninstall

	IPAddress string `json:"ipAddress,omitempty"` // IP address

	HostName string `json:"hostName,omitempty"` // Host name

	WiredApplicationStatus string `json:"wiredApplicationStatus,omitempty"` // Wired application status

	WiredApplicationMessage string `json:"wiredApplicationMessage,omitempty"` // Wired application message

	Assigned *bool `json:"assigned,omitempty"` // Is assigned

	Status string `json:"status,omitempty"` // Sensor device status: UP, DOWN, REBOOT

	XorSensor *bool `json:"xorSensor,omitempty"` // Is XOR sensor

	TargetAPs []string `json:"targetAPs,omitempty"` // Array of target APs

	RunNow string `json:"runNow,omitempty"` // Run now: YES, NO

	LocationID string `json:"locationId,omitempty"` // Site UUID

	AllSensorAddition *bool `json:"allSensorAddition,omitempty"` // Is all sensor addition

	ConfigUpdated string `json:"configUpdated,omitempty"` // Configuration updated: YES, NO

	SensorType string `json:"sensorType,omitempty"` // Sensor type

	TestMacAddresses *RequestSensorsCreateSensorTestTemplateV1SensorsTestMacAddresses `json:"testMacAddresses,omitempty"` // A string-string test MAC address

	ID string `json:"id,omitempty"` // Sensor ID

	ServicePolicy string `json:"servicePolicy,omitempty"` // Service policy

	IPerfInfo *RequestSensorsCreateSensorTestTemplateV1SensorsIPerfInfo `json:"iPerfInfo,omitempty"` // A string-stringList iPerf information
}
type RequestSensorsCreateSensorTestTemplateV1SensorsTestMacAddresses interface{}
type RequestSensorsCreateSensorTestTemplateV1SensorsIPerfInfo interface{}
type RequestSensorsCreateSensorTestTemplateV1ApCoverage struct {
	Bands string `json:"bands,omitempty"` // The WIFI bands

	NumberOfApsToTest *int `json:"numberOfApsToTest,omitempty"` // Number of APs to test

	RssiThreshold *int `json:"rssiThreshold,omitempty"` // RSSI threshold
}
type RequestSensorsRunNowSensorTestV1 struct {
	TemplateName string `json:"templateName,omitempty"` // Template Name
}
type RequestSensorsDuplicateSensorTestTemplateV1 struct {
	TemplateName string `json:"templateName,omitempty"` // Source test template name

	NewTemplateName string `json:"newTemplateName,omitempty"` // Destination test template name
}

//ListsICapPacketCaptureFilesMatchingSpecifiedCriteriaV1 Lists ICAP packet capture files matching specified criteria - b88d-fbd6-4639-9b2f
/* Lists the ICAP packet capture (pcap) files matching the specified criteria. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-icap-1.0.0-resolved.yaml


@param ListsICAPPacketCaptureFilesMatchingSpecifiedCriteriaV1HeaderParams Custom header parameters
@param ListsICAPPacketCaptureFilesMatchingSpecifiedCriteriaV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!lists-i-cap-packet-capture-files-matching-specified-criteria
*/
func (s *SensorsService) ListsICapPacketCaptureFilesMatchingSpecifiedCriteriaV1(ListsICAPPacketCaptureFilesMatchingSpecifiedCriteriaV1HeaderParams *ListsICapPacketCaptureFilesMatchingSpecifiedCriteriaV1HeaderParams, ListsICAPPacketCaptureFilesMatchingSpecifiedCriteriaV1QueryParams *ListsICapPacketCaptureFilesMatchingSpecifiedCriteriaV1QueryParams) (*ResponseSensorsListsICapPacketCaptureFilesMatchingSpecifiedCriteriaV1, *resty.Response, error) {
	path := "/dna/data/api/v1/icap/captureFiles"

	queryString, _ := query.Values(ListsICAPPacketCaptureFilesMatchingSpecifiedCriteriaV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if ListsICAPPacketCaptureFilesMatchingSpecifiedCriteriaV1HeaderParams != nil {

		if ListsICAPPacketCaptureFilesMatchingSpecifiedCriteriaV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", ListsICAPPacketCaptureFilesMatchingSpecifiedCriteriaV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSensorsListsICapPacketCaptureFilesMatchingSpecifiedCriteriaV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ListsICapPacketCaptureFilesMatchingSpecifiedCriteriaV1(ListsICAPPacketCaptureFilesMatchingSpecifiedCriteriaV1HeaderParams, ListsICAPPacketCaptureFilesMatchingSpecifiedCriteriaV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ListsICapPacketCaptureFilesMatchingSpecifiedCriteriaV1")
	}

	result := response.Result().(*ResponseSensorsListsICapPacketCaptureFilesMatchingSpecifiedCriteriaV1)
	return result, response, err

}

//RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaV1 Retrieves the total number of packet capture files matching specified criteria - 5483-e972-4b1a-98c6
/* Retrieves the total number of packet capture files matching the specified criteria. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-icap-1.0.0-resolved.yaml


@param RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaV1HeaderParams Custom header parameters
@param RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-total-number-of-packet-capture-files-matching-specified-criteria
*/
func (s *SensorsService) RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaV1(RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaV1HeaderParams *RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaV1HeaderParams, RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaV1QueryParams *RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaV1QueryParams) (*ResponseSensorsRetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaV1, *resty.Response, error) {
	path := "/dna/data/api/v1/icap/captureFiles/count"

	queryString, _ := query.Values(RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaV1HeaderParams != nil {

		if RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSensorsRetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaV1(RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaV1HeaderParams, RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaV1")
	}

	result := response.Result().(*ResponseSensorsRetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaV1)
	return result, response, err

}

//RetrievesDetailsOfASpecificICapPacketCaptureFileV1 Retrieves details of a specific ICAP packet capture file - 7bb7-5afe-4cc8-be27
/* Retrieves details of a specific ICAP packet capture file. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-icap-1.0.0-resolved.yaml


@param id id path parameter. The name of the packet capture file, as given by the GET /captureFiles API response.

@param RetrievesDetailsOfASpecificICAPPacketCaptureFileV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-details-of-a-specific-i-cap-packet-capture-file
*/
func (s *SensorsService) RetrievesDetailsOfASpecificICapPacketCaptureFileV1(id string, RetrievesDetailsOfASpecificICAPPacketCaptureFileV1HeaderParams *RetrievesDetailsOfASpecificICapPacketCaptureFileV1HeaderParams) (*ResponseSensorsRetrievesDetailsOfASpecificICapPacketCaptureFileV1, *resty.Response, error) {
	path := "/dna/data/api/v1/icap/captureFiles/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesDetailsOfASpecificICAPPacketCaptureFileV1HeaderParams != nil {

		if RetrievesDetailsOfASpecificICAPPacketCaptureFileV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesDetailsOfASpecificICAPPacketCaptureFileV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetResult(&ResponseSensorsRetrievesDetailsOfASpecificICapPacketCaptureFileV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesDetailsOfASpecificICapPacketCaptureFileV1(id, RetrievesDetailsOfASpecificICAPPacketCaptureFileV1HeaderParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesDetailsOfASpecificICapPacketCaptureFileV1")
	}

	result := response.Result().(*ResponseSensorsRetrievesDetailsOfASpecificICapPacketCaptureFileV1)
	return result, response, err

}

//DownloadsASpecificICapPacketCaptureFileV1 Downloads a specific ICAP packet capture file - 47b2-db4f-4468-ba5d
/* Downloads a specific ICAP packet capture file. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-icap-1.0.0-resolved.yaml


@param id id path parameter. The name of the packet capture file, as given by the GET /captureFiles API response.

@param DownloadsASpecificICAPPacketCaptureFileV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!downloads-a-specific-i-cap-packet-capture-file
*/
func (s *SensorsService) DownloadsASpecificICapPacketCaptureFileV1(id string, DownloadsASpecificICAPPacketCaptureFileV1HeaderParams *DownloadsASpecificICapPacketCaptureFileV1HeaderParams) (FileDownload, *resty.Response, error) {
	path := "/dna/data/api/v1/icap/captureFiles/{id}/download"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if DownloadsASpecificICAPPacketCaptureFileV1HeaderParams != nil {

		if DownloadsASpecificICAPPacketCaptureFileV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", DownloadsASpecificICAPPacketCaptureFileV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetResult(&ResponseSensorsDownloadsASpecificICapPacketCaptureFileV1{}).
		SetError(&Error).
		Get(path)
	fdownload := FileDownload{}
	if err != nil {
		return fdownload, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DownloadsASpecificICapPacketCaptureFileV1(id, DownloadsASpecificICAPPacketCaptureFileV1HeaderParams)
		}
		return fdownload, response, fmt.Errorf("error with operation DownloadsASpecificICapPacketCaptureFileV1")
	}

	fdownload.FileData = response.Body()
	headerVal := response.Header()["Content-Disposition"][0]
	fname := strings.SplitAfter(headerVal, "=")
	fdownload.FileName = strings.ReplaceAll(fname[1], "\"", "")

	return fdownload, response, err

}

//RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMacV1 Retrieves the spectrum interference devices reports sent by WLC for provided AP Mac - 78a6-8aec-4278-8cdc
/* Retrieves the spectrum interference devices reports sent by WLC for provided AP Mac. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-icap-1.0.0-resolved.yaml


@param RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedAPMacV1HeaderParams Custom header parameters
@param RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedAPMacV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-spectrum-interference-devices-reports-sent-by-w-l-c-for-provided-ap-mac
*/
func (s *SensorsService) RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMacV1(RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedAPMacV1HeaderParams *RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMacV1HeaderParams, RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedAPMacV1QueryParams *RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMacV1QueryParams) (*ResponseSensorsRetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMacV1, *resty.Response, error) {
	path := "/dna/data/api/v1/icap/spectrumInterferenceDeviceReports"

	queryString, _ := query.Values(RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedAPMacV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedAPMacV1HeaderParams != nil {

		if RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedAPMacV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedAPMacV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSensorsRetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMacV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMacV1(RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedAPMacV1HeaderParams, RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedAPMacV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMacV1")
	}

	result := response.Result().(*ResponseSensorsRetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMacV1)
	return result, response, err

}

//RetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMacV1 Retrieves the spectrum sensor reports sent by WLC for provided AP Mac - f583-d956-4c1a-8cc3
/* Retrieves the spectrum sensor reports sent by WLC for provided AP Mac. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-icap-1.0.0-resolved.yaml


@param RetrievesTheSpectrumSensorReportsSentByWLCForProvidedAPMacV1HeaderParams Custom header parameters
@param RetrievesTheSpectrumSensorReportsSentByWLCForProvidedAPMacV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-spectrum-sensor-reports-sent-by-w-l-c-for-provided-ap-mac
*/
func (s *SensorsService) RetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMacV1(RetrievesTheSpectrumSensorReportsSentByWLCForProvidedAPMacV1HeaderParams *RetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMacV1HeaderParams, RetrievesTheSpectrumSensorReportsSentByWLCForProvidedAPMacV1QueryParams *RetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMacV1QueryParams) (*ResponseSensorsRetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMacV1, *resty.Response, error) {
	path := "/dna/data/api/v1/icap/spectrumSensorReports"

	queryString, _ := query.Values(RetrievesTheSpectrumSensorReportsSentByWLCForProvidedAPMacV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheSpectrumSensorReportsSentByWLCForProvidedAPMacV1HeaderParams != nil {

		if RetrievesTheSpectrumSensorReportsSentByWLCForProvidedAPMacV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheSpectrumSensorReportsSentByWLCForProvidedAPMacV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSensorsRetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMacV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMacV1(RetrievesTheSpectrumSensorReportsSentByWLCForProvidedAPMacV1HeaderParams, RetrievesTheSpectrumSensorReportsSentByWLCForProvidedAPMacV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMacV1")
	}

	result := response.Result().(*ResponseSensorsRetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMacV1)
	return result, response, err

}

//RetrievesDeployedICapConfigurationsWhileSupportingBasicFilteringV1 Retrieves deployed ICAP configurations while supporting basic filtering. - c780-8b33-44c8-858c
/* Retrieves deployed ICAP configurations while supporting basic filtering. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-ICAP_APIs-1.0.0-resolved.yaml


@param RetrievesDeployedICAPConfigurationsWhileSupportingBasicFilteringV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-deployed-i-cap-configurations-while-supporting-basic-filtering
*/
func (s *SensorsService) RetrievesDeployedICapConfigurationsWhileSupportingBasicFilteringV1(RetrievesDeployedICAPConfigurationsWhileSupportingBasicFilteringV1QueryParams *RetrievesDeployedICapConfigurationsWhileSupportingBasicFilteringV1QueryParams) (*ResponseSensorsRetrievesDeployedICapConfigurationsWhileSupportingBasicFilteringV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/icapSettings"

	queryString, _ := query.Values(RetrievesDeployedICAPConfigurationsWhileSupportingBasicFilteringV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSensorsRetrievesDeployedICapConfigurationsWhileSupportingBasicFilteringV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesDeployedICapConfigurationsWhileSupportingBasicFilteringV1(RetrievesDeployedICAPConfigurationsWhileSupportingBasicFilteringV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesDeployedICapConfigurationsWhileSupportingBasicFilteringV1")
	}

	result := response.Result().(*ResponseSensorsRetrievesDeployedICapConfigurationsWhileSupportingBasicFilteringV1)
	return result, response, err

}

//GetICapConfigurationStatusPerNetworkDeviceV1 Get ICAP configuration status per network device. - 5291-5b76-46ab-98eb
/* Get ICAP configuration status per network device using the activity ID, which was returned in property "taskId" of the TaskResponse of the POST. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-ICAP_APIs-1.0.0-resolved.yaml


@param previewActivityID previewActivityId path parameter. activity from the POST /deviceConfigugrationModels task response


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-i-cap-configuration-status-per-network-device
*/
func (s *SensorsService) GetICapConfigurationStatusPerNetworkDeviceV1(previewActivityID string) (*ResponseSensorsGetICapConfigurationStatusPerNetworkDeviceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/icapSettings/configurationModels/{previewActivityId}/networkDeviceStatusDetails"
	path = strings.Replace(path, "{previewActivityId}", fmt.Sprintf("%v", previewActivityID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSensorsGetICapConfigurationStatusPerNetworkDeviceV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetICapConfigurationStatusPerNetworkDeviceV1(previewActivityID)
		}
		return nil, response, fmt.Errorf("error with operation GetICapConfigurationStatusPerNetworkDeviceV1")
	}

	result := response.Result().(*ResponseSensorsGetICapConfigurationStatusPerNetworkDeviceV1)
	return result, response, err

}

//RetrievesTheDevicesClisOfTheICapintentV1 Retrieves the device's CLIs of the ICAP intent. - 4288-bbf8-4b59-8fcf
/* Returns the device's CLIs of the ICAP intent. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-ICAP_APIs-1.0.0-resolved.yaml


@param previewActivityID previewActivityId path parameter. activity from the POST /deviceConfigugrationModels task response

@param networkDeviceID networkDeviceId path parameter. device id from intent/api/v1/network-device


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-devices-clis-of-the-i-capintent
*/
func (s *SensorsService) RetrievesTheDevicesClisOfTheICapintentV1(previewActivityID string, networkDeviceID string) (*ResponseSensorsRetrievesTheDevicesClisOfTheICapintentV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/icapSettings/configurationModels/{previewActivityId}/networkDevices/{networkDeviceId}/config"
	path = strings.Replace(path, "{previewActivityId}", fmt.Sprintf("%v", previewActivityID), -1)
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSensorsRetrievesTheDevicesClisOfTheICapintentV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheDevicesClisOfTheICapintentV1(previewActivityID, networkDeviceID)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheDevicesClisOfTheICapintentV1")
	}

	result := response.Result().(*ResponseSensorsRetrievesTheDevicesClisOfTheICapintentV1)
	return result, response, err

}

//RetrievesTheCountOfDeployedICapConfigurationsWhileSupportingBasicFilteringV1 Retrieves the count of deployed ICAP configurations while supporting basic filtering. - 46b9-694b-41ca-8155
/* Retrieves the count of deployed ICAP configurations while supporting basic filtering. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-ICAP_APIs-1.0.0-resolved.yaml


@param RetrievesTheCountOfDeployedICAPConfigurationsWhileSupportingBasicFilteringV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-count-of-deployed-i-cap-configurations-while-supporting-basic-filtering
*/
func (s *SensorsService) RetrievesTheCountOfDeployedICapConfigurationsWhileSupportingBasicFilteringV1(RetrievesTheCountOfDeployedICAPConfigurationsWhileSupportingBasicFilteringV1QueryParams *RetrievesTheCountOfDeployedICapConfigurationsWhileSupportingBasicFilteringV1QueryParams) (*ResponseSensorsRetrievesTheCountOfDeployedICapConfigurationsWhileSupportingBasicFilteringV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/icapSettings/count"

	queryString, _ := query.Values(RetrievesTheCountOfDeployedICAPConfigurationsWhileSupportingBasicFilteringV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSensorsRetrievesTheCountOfDeployedICapConfigurationsWhileSupportingBasicFilteringV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheCountOfDeployedICapConfigurationsWhileSupportingBasicFilteringV1(RetrievesTheCountOfDeployedICAPConfigurationsWhileSupportingBasicFilteringV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheCountOfDeployedICapConfigurationsWhileSupportingBasicFilteringV1")
	}

	result := response.Result().(*ResponseSensorsRetrievesTheCountOfDeployedICapConfigurationsWhileSupportingBasicFilteringV1)
	return result, response, err

}

//GetDeviceDeploymentStatusV1 Get device deployment status. - 5aa1-1ac0-4d28-af86
/* Retrieves ICAP configuration deployment status(s) per device based on filter criteria. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-ICAP_APIs-1.0.0-resolved.yaml


@param GetDeviceDeploymentStatusV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-deployment-status
*/
func (s *SensorsService) GetDeviceDeploymentStatusV1(GetDeviceDeploymentStatusV1QueryParams *GetDeviceDeploymentStatusV1QueryParams) (*ResponseSensorsGetDeviceDeploymentStatusV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/icapSettings/deviceDeployments"

	queryString, _ := query.Values(GetDeviceDeploymentStatusV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSensorsGetDeviceDeploymentStatusV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceDeploymentStatusV1(GetDeviceDeploymentStatusV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceDeploymentStatusV1")
	}

	result := response.Result().(*ResponseSensorsGetDeviceDeploymentStatusV1)
	return result, response, err

}

//GetDeviceDeploymentStatusCountV1 Get device deployment status count. - 6d80-6a7b-4459-a1a3
/* Returns the count of device deployment status(s) based on filter criteria. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-ICAP_APIs-1.0.0-resolved.yaml


@param GetDeviceDeploymentStatusCountV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-deployment-status-count
*/
func (s *SensorsService) GetDeviceDeploymentStatusCountV1(GetDeviceDeploymentStatusCountV1QueryParams *GetDeviceDeploymentStatusCountV1QueryParams) (*ResponseSensorsGetDeviceDeploymentStatusCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/icapSettings/deviceDeployments/count"

	queryString, _ := query.Values(GetDeviceDeploymentStatusCountV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSensorsGetDeviceDeploymentStatusCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceDeploymentStatusCountV1(GetDeviceDeploymentStatusCountV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceDeploymentStatusCountV1")
	}

	result := response.Result().(*ResponseSensorsGetDeviceDeploymentStatusCountV1)
	return result, response, err

}

//SensorsV1 Sensors - 71a1-2bb7-4569-9cc5
/* Intent API to get a list of SENSOR devices


@param SensorsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!sensors
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

//RetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1 Retrieves specific client statistics over specified period of time. - 8e8d-29bf-4019-95fb
/* Retrieves the time series statistics of a specific client by applying complex filters. If startTime and endTime are not provided, the API defaults to the last 24 hours. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-icap-1.0.0-resolved.yaml


@param id id path parameter. id is the client mac address. It can be specified in one of the notational conventions 01:23:45:67:89:AB or 01-23-45-67-89-AB or 0123.4567.89AB and is case insensitive

@param RetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-specific-client-statistics-over-specified-period-of-time
*/
func (s *SensorsService) RetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1(id string, requestSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1 *RequestSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1, RetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1HeaderParams *RetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1HeaderParams) (*ResponseSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1, *resty.Response, error) {
	path := "/dna/data/api/v1/icap/clients/{id}/stats"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1HeaderParams != nil {

		if RetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1).
		SetResult(&ResponseSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1(id, requestSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1, RetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation RetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1")
	}

	result := response.Result().(*ResponseSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1)
	return result, response, err

}

//RetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1 Retrieves specific radio statistics over specified period of time. - 9fac-b838-4c29-b3b5
/* Retrieves the time series statistics of a specific radio by applying complex filters. If startTime and endTime are not provided, the API defaults to the last 24 hours. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-icap-1.0.0-resolved.yaml


@param id id path parameter. id is the composite key made of AP Base Ethernet macAddress and Radio Slot Id. Format apMac_RadioId

@param RetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-specific-radio-statistics-over-specified-period-of-time
*/
func (s *SensorsService) RetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1(id string, requestSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1 *RequestSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1, RetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1HeaderParams *RetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1HeaderParams) (*ResponseSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1, *resty.Response, error) {
	path := "/dna/data/api/v1/icap/radios/{id}/stats"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1HeaderParams != nil {

		if RetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1HeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1HeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1).
		SetResult(&ResponseSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1(id, requestSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1, RetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation RetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1")
	}

	result := response.Result().(*ResponseSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1)
	return result, response, err

}

//CreatesAnICapConfigurationIntentForPreviewApproveV1 Creates an ICAP configuration intent for preview-approve. - be85-b906-4ea8-acfa
/* This creates an ICAP configuration intent for preview approval. The intent is not deployed to the device until further preview-approve APIs are applied. This API is the first step in the preview-approve workflow, which consists of several APIs. Skipping any API in the process is not recommended for a complete preview-approve use case. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-ICAP_APIs-1.0.0-resolved.yaml


@param CreatesAnICAPConfigurationIntentForPreviewApproveV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!creates-an-i-cap-configuration-intent-for-preview-approve
*/
func (s *SensorsService) CreatesAnICapConfigurationIntentForPreviewApproveV1(requestSensorsCreatesAnICAPConfigurationIntentForPreviewApproveV1 *RequestSensorsCreatesAnICapConfigurationIntentForPreviewApproveV1, CreatesAnICAPConfigurationIntentForPreviewApproveV1QueryParams *CreatesAnICapConfigurationIntentForPreviewApproveV1QueryParams) (*ResponseSensorsCreatesAnICapConfigurationIntentForPreviewApproveV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/icapSettings/configurationModels"

	queryString, _ := query.Values(CreatesAnICAPConfigurationIntentForPreviewApproveV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetBody(requestSensorsCreatesAnICAPConfigurationIntentForPreviewApproveV1).
		SetResult(&ResponseSensorsCreatesAnICapConfigurationIntentForPreviewApproveV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreatesAnICapConfigurationIntentForPreviewApproveV1(requestSensorsCreatesAnICAPConfigurationIntentForPreviewApproveV1, CreatesAnICAPConfigurationIntentForPreviewApproveV1QueryParams)
		}

		return nil, response, fmt.Errorf("error with operation CreatesAnICapConfigurationIntentForPreviewApproveV1")
	}

	result := response.Result().(*ResponseSensorsCreatesAnICapConfigurationIntentForPreviewApproveV1)
	return result, response, err

}

//CreatesAICapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDeviceV1 Creates a ICAP configuration workflow for ICAP intent to remove the ICAP configuration on the device. - e681-4857-4298-9079
/* Creates a ICAP configuration intent to remove the ICAP RFSTATS or ANOMALY configuration from the device. The task has not been applied to the device yet. Subsequent preview-approve workflow APIs must be used to complete the preview-approve process.  The path parameter 'id' can be retrieved from **GET /dna/intent/api/v1/icapSettings** API. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-ICAP_APIs-1.0.0-resolved.yaml


@param id id path parameter. A unique ID of the deployed ICAP object, which can be obtained from **GET /dna/intent/api/v1/icapSettings**


Documentation Link: https://developer.cisco.com/docs/dna-center/#!creates-a-i-cap-configuration-workflow-for-i-capintent-to-remove-the-i-cap-configuration-on-the-device
*/
func (s *SensorsService) CreatesAICapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDeviceV1(id string, requestSensorsCreatesAICAPConfigurationWorkflowForICAPIntentToRemoveTheICAPConfigurationOnTheDeviceV1 *RequestSensorsCreatesAICapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDeviceV1) (*ResponseSensorsCreatesAICapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDeviceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/icapSettings/configurationModels/{id}/deleteDeploy"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSensorsCreatesAICAPConfigurationWorkflowForICAPIntentToRemoveTheICAPConfigurationOnTheDeviceV1).
		SetResult(&ResponseSensorsCreatesAICapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDeviceV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreatesAICapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDeviceV1(id, requestSensorsCreatesAICAPConfigurationWorkflowForICAPIntentToRemoveTheICAPConfigurationOnTheDeviceV1)
		}

		return nil, response, fmt.Errorf("error with operation CreatesAICapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDeviceV1")
	}

	result := response.Result().(*ResponseSensorsCreatesAICapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDeviceV1)
	return result, response, err

}

//DeploysTheICapConfigurationIntentByActivityIDV1 Deploys the ICAP configuration intent by activity ID. - f58e-6b46-42ba-ae2f
/* Deploys the ICAP configuration intent by activity ID, which was returned in property "taskId" of the TaskResponse of the POST.  POST'ing the intent prior to generating the intent CLI for preview-approve has the same effect as direct-deploy'ing the intent to the device.
Generating of device's CLIs for preview-approve is not available for this activity ID after using this POST API. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-ICAP_APIs-1.0.0-resolved.yaml


@param previewActivityID previewActivityId path parameter. activity from the POST /deviceConfigugrationModels task response


Documentation Link: https://developer.cisco.com/docs/dna-center/#!deploys-the-i-cap-configuration-intent-by-activity-id
*/
func (s *SensorsService) DeploysTheICapConfigurationIntentByActivityIDV1(previewActivityID string, requestSensorsDeploysTheICAPConfigurationIntentByActivityIDV1 *RequestSensorsDeploysTheICapConfigurationIntentByActivityIDV1) (*ResponseSensorsDeploysTheICapConfigurationIntentByActivityIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/icapSettings/configurationModels/{previewActivityId}/deploy"
	path = strings.Replace(path, "{previewActivityId}", fmt.Sprintf("%v", previewActivityID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSensorsDeploysTheICAPConfigurationIntentByActivityIDV1).
		SetResult(&ResponseSensorsDeploysTheICapConfigurationIntentByActivityIDV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeploysTheICapConfigurationIntentByActivityIDV1(previewActivityID, requestSensorsDeploysTheICAPConfigurationIntentByActivityIDV1)
		}

		return nil, response, fmt.Errorf("error with operation DeploysTheICapConfigurationIntentByActivityIdV1")
	}

	result := response.Result().(*ResponseSensorsDeploysTheICapConfigurationIntentByActivityIDV1)
	return result, response, err

}

//GeneratesTheDevicesClisOfTheICapConfigurationIntentV1 Generates the device's CLIs of the ICAP configuration intent. - e1b2-aaf0-4358-9325
/* Generates the device's CLIs of the ICAP intent for preview and approve prior to deploying the ICAP configuration intent to the device.  After deploying the configuration intent, generating intent CLIs will not be available for preview.


@param previewActivityID previewActivityId path parameter. activity from the POST /deviceConfigugrationModels task response

@param networkDeviceID networkDeviceId path parameter. device id from intent/api/v1/network-device


Documentation Link: https://developer.cisco.com/docs/dna-center/#!generates-the-devices-clis-of-the-i-cap-configuration-intent
*/
func (s *SensorsService) GeneratesTheDevicesClisOfTheICapConfigurationIntentV1(previewActivityID string, networkDeviceID string, requestSensorsGeneratesTheDevicesCLIsOfTheICAPConfigurationIntentV1 *RequestSensorsGeneratesTheDevicesClisOfTheICapConfigurationIntentV1) (*ResponseSensorsGeneratesTheDevicesClisOfTheICapConfigurationIntentV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/icapSettings/configurationModels/{previewActivityId}/networkDevices/{networkDeviceId}/config"
	path = strings.Replace(path, "{previewActivityId}", fmt.Sprintf("%v", previewActivityID), -1)
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSensorsGeneratesTheDevicesCLIsOfTheICAPConfigurationIntentV1).
		SetResult(&ResponseSensorsGeneratesTheDevicesClisOfTheICapConfigurationIntentV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GeneratesTheDevicesClisOfTheICapConfigurationIntentV1(previewActivityID, networkDeviceID, requestSensorsGeneratesTheDevicesCLIsOfTheICAPConfigurationIntentV1)
		}

		return nil, response, fmt.Errorf("error with operation GeneratesTheDevicesClisOfTheICapConfigurationIntentV1")
	}

	result := response.Result().(*ResponseSensorsGeneratesTheDevicesClisOfTheICapConfigurationIntentV1)
	return result, response, err

}

//DeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1 Deploys the given ICAP configuration intent without preview and approve. - 78ba-0a27-4808-b722
/* Deploys the given ICAP intent without preview and approval. The response body contains a task object with a taskId and a URL for more information about the task. The deployment status of this ICAP intent can be found in the output of the URL.  For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-ICAP_APIs-1.0.0-resolved.yaml


@param DeploysTheGivenICAPConfigurationIntentWithoutPreviewAndApproveV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!deploys-the-given-i-cap-configuration-intent-without-preview-and-approve
*/
func (s *SensorsService) DeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1(requestSensorsDeploysTheGivenICAPConfigurationIntentWithoutPreviewAndApproveV1 *RequestSensorsDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1, DeploysTheGivenICAPConfigurationIntentWithoutPreviewAndApproveV1QueryParams *DeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1QueryParams) (*ResponseSensorsDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/icapSettings/deploy"

	queryString, _ := query.Values(DeploysTheGivenICAPConfigurationIntentWithoutPreviewAndApproveV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetBody(requestSensorsDeploysTheGivenICAPConfigurationIntentWithoutPreviewAndApproveV1).
		SetResult(&ResponseSensorsDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1(requestSensorsDeploysTheGivenICAPConfigurationIntentWithoutPreviewAndApproveV1, DeploysTheGivenICAPConfigurationIntentWithoutPreviewAndApproveV1QueryParams)
		}

		return nil, response, fmt.Errorf("error with operation DeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1")
	}

	result := response.Result().(*ResponseSensorsDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1)
	return result, response, err

}

//RemoveTheICapConfigurationOnTheDeviceWithoutPreviewV1 Remove the ICAP configuration on the device without preview - 779f-7b3d-40e9-b736
/* Remove the ICAP configuration from the device by *id* without preview-deploy. The path parameter *id* can be retrieved from the **GET /dna/intent/api/v1/icapSettings** API. The response body contains a task object with a taskId and a URL. Use the URL to check the task status. ICAP FULL, ONBOARDING, OTA, and SPECTRUM configurations have a durationInMins field. A disable task is scheduled to remove the configuration from the device. Removing the ICAP intent should be done after the pre-scheduled disable task has been deployed. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-ICAP_APIs-1.0.0-resolved.yaml


@param id id path parameter. A unique ID of the deployed ICAP object, which can be obtained from **GET /dna/intent/api/v1/icapSettings**


Documentation Link: https://developer.cisco.com/docs/dna-center/#!remove-the-i-cap-configuration-on-the-device-without-preview
*/
func (s *SensorsService) RemoveTheICapConfigurationOnTheDeviceWithoutPreviewV1(id string, requestSensorsRemoveTheICAPConfigurationOnTheDeviceWithoutPreviewV1 *RequestSensorsRemoveTheICapConfigurationOnTheDeviceWithoutPreviewV1) (*ResponseSensorsRemoveTheICapConfigurationOnTheDeviceWithoutPreviewV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/icapSettings/deploy/{id}/deleteDeploy"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSensorsRemoveTheICAPConfigurationOnTheDeviceWithoutPreviewV1).
		SetResult(&ResponseSensorsRemoveTheICapConfigurationOnTheDeviceWithoutPreviewV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RemoveTheICapConfigurationOnTheDeviceWithoutPreviewV1(id, requestSensorsRemoveTheICAPConfigurationOnTheDeviceWithoutPreviewV1)
		}

		return nil, response, fmt.Errorf("error with operation RemoveTheICapConfigurationOnTheDeviceWithoutPreviewV1")
	}

	result := response.Result().(*ResponseSensorsRemoveTheICapConfigurationOnTheDeviceWithoutPreviewV1)
	return result, response, err

}

//CreateSensorTestTemplateV1 Create sensor test template - 08bd-8883-4a68-a2e6
/* Intent API to create a SENSOR test template with a new SSID, existing SSID, or both new and existing SSID



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-sensor-test-template
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

//DiscardsTheICapConfigurationIntentByActivityIDV1 Discards the ICAP configuration intent by activity ID. - 02b1-1ba1-4afb-9155
/* Discard the ICAP configuration intent by activity ID, which was returned in TaskResponse's property "taskId" at the beginning of the preview-approve workflow.  Discarding the intent can only be applied to intent activities that have not been deployed.
Note that ICAP type FULL, ONBOARDING, OTA, and SPECTRUM for the scheduled-disabled task cannot be discarded or cancelled because they have already deployed.  The feature can only be disabled by sending in a direct-deploy DELETE with API /dna/intent/api/v1/icapSettings/deploy/deployedId/{icapDeployedId} For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-ICAP_APIs-1.0.0-resolved.yaml


@param previewActivityID previewActivityId path parameter. activity from the POST /deviceConfigugrationModels task response


Documentation Link: https://developer.cisco.com/docs/dna-center/#!discards-the-i-cap-configuration-intent-by-activity-id
*/
func (s *SensorsService) DiscardsTheICapConfigurationIntentByActivityIDV1(previewActivityID string) (*ResponseSensorsDiscardsTheICapConfigurationIntentByActivityIDV1, *resty.Response, error) {
	//previewActivityID string
	path := "/dna/intent/api/v1/icapSettings/configurationModels/{previewActivityId}"
	path = strings.Replace(path, "{previewActivityId}", fmt.Sprintf("%v", previewActivityID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSensorsDiscardsTheICapConfigurationIntentByActivityIDV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DiscardsTheICapConfigurationIntentByActivityIDV1(previewActivityID)
		}
		return nil, response, fmt.Errorf("error with operation DiscardsTheICapConfigurationIntentByActivityIdV1")
	}

	result := response.Result().(*ResponseSensorsDiscardsTheICapConfigurationIntentByActivityIDV1)
	return result, response, err

}

//DeleteSensorTestV1 Delete sensor test - 5bbb-28ff-442a-825f
/* Intent API to delete an existing SENSOR test template


@param DeleteSensorTestV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-sensor-test
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
/*
This method acts as an alias for the method `RetrievesTheDevicesClisOfTheICapintentV1`
*/
func (s *SensorsService) RetrievesTheDevicesClisOfTheICapintent(previewActivityID string, networkDeviceID string) (*ResponseSensorsRetrievesTheDevicesClisOfTheICapintentV1, *resty.Response, error) {
	return s.RetrievesTheDevicesClisOfTheICapintentV1(previewActivityID, networkDeviceID)
}

// Alias Function
/*
This method acts as an alias for the method `GetDeviceDeploymentStatusCountV1`
*/
func (s *SensorsService) GetDeviceDeploymentStatusCount(GetDeviceDeploymentStatusCountV1QueryParams *GetDeviceDeploymentStatusCountV1QueryParams) (*ResponseSensorsGetDeviceDeploymentStatusCountV1, *resty.Response, error) {
	return s.GetDeviceDeploymentStatusCountV1(GetDeviceDeploymentStatusCountV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `DiscardsTheICapConfigurationIntentByActivityIDV1`
*/
func (s *SensorsService) DiscardsTheICapConfigurationIntentByActivityID(previewActivityID string) (*ResponseSensorsDiscardsTheICapConfigurationIntentByActivityIDV1, *resty.Response, error) {
	return s.DiscardsTheICapConfigurationIntentByActivityIDV1(previewActivityID)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1`
*/
func (s *SensorsService) RetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTime(id string, requestSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1 *RequestSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1, RetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1HeaderParams *RetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1HeaderParams) (*ResponseSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1, *resty.Response, error) {
	return s.RetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1(id, requestSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1, RetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `CreatesAICapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDeviceV1`
*/
func (s *SensorsService) CreatesAICapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDevice(id string, requestSensorsCreatesAICAPConfigurationWorkflowForICAPIntentToRemoveTheICAPConfigurationOnTheDeviceV1 *RequestSensorsCreatesAICapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDeviceV1) (*ResponseSensorsCreatesAICapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDeviceV1, *resty.Response, error) {
	return s.CreatesAICapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDeviceV1(id, requestSensorsCreatesAICAPConfigurationWorkflowForICAPIntentToRemoveTheICAPConfigurationOnTheDeviceV1)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaV1`
*/
func (s *SensorsService) RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteria(RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaV1HeaderParams *RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaV1HeaderParams, RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaV1QueryParams *RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaV1QueryParams) (*ResponseSensorsRetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaV1, *resty.Response, error) {
	return s.RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaV1(RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaV1HeaderParams, RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `CreateSensorTestTemplateV1`
*/
func (s *SensorsService) CreateSensorTestTemplate(requestSensorsCreateSensorTestTemplateV1 *RequestSensorsCreateSensorTestTemplateV1) (*ResponseSensorsCreateSensorTestTemplateV1, *resty.Response, error) {
	return s.CreateSensorTestTemplateV1(requestSensorsCreateSensorTestTemplateV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetDeviceDeploymentStatusV1`
*/
func (s *SensorsService) GetDeviceDeploymentStatus(GetDeviceDeploymentStatusV1QueryParams *GetDeviceDeploymentStatusV1QueryParams) (*ResponseSensorsGetDeviceDeploymentStatusV1, *resty.Response, error) {
	return s.GetDeviceDeploymentStatusV1(GetDeviceDeploymentStatusV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `ListsICapPacketCaptureFilesMatchingSpecifiedCriteriaV1`
*/
func (s *SensorsService) ListsICapPacketCaptureFilesMatchingSpecifiedCriteria(ListsICAPPacketCaptureFilesMatchingSpecifiedCriteriaV1HeaderParams *ListsICapPacketCaptureFilesMatchingSpecifiedCriteriaV1HeaderParams, ListsICAPPacketCaptureFilesMatchingSpecifiedCriteriaV1QueryParams *ListsICapPacketCaptureFilesMatchingSpecifiedCriteriaV1QueryParams) (*ResponseSensorsListsICapPacketCaptureFilesMatchingSpecifiedCriteriaV1, *resty.Response, error) {
	return s.ListsICapPacketCaptureFilesMatchingSpecifiedCriteriaV1(ListsICAPPacketCaptureFilesMatchingSpecifiedCriteriaV1HeaderParams, ListsICAPPacketCaptureFilesMatchingSpecifiedCriteriaV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `RunNowSensorTestV1`
*/
func (s *SensorsService) RunNowSensorTest(requestSensorsRunNowSensorTestV1 *RequestSensorsRunNowSensorTestV1) (*resty.Response, error) {
	return s.RunNowSensorTestV1(requestSensorsRunNowSensorTestV1)
}

// Alias Function
/*
This method acts as an alias for the method `DeleteSensorTestV1`
*/
func (s *SensorsService) DeleteSensorTest(DeleteSensorTestV1QueryParams *DeleteSensorTestV1QueryParams) (*ResponseSensorsDeleteSensorTestV1, *resty.Response, error) {
	return s.DeleteSensorTestV1(DeleteSensorTestV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `DuplicateSensorTestTemplateV1`
*/
func (s *SensorsService) DuplicateSensorTestTemplate(requestSensorsDuplicateSensorTestTemplateV1 *RequestSensorsDuplicateSensorTestTemplateV1) (*ResponseSensorsDuplicateSensorTestTemplateV1, *resty.Response, error) {
	return s.DuplicateSensorTestTemplateV1(requestSensorsDuplicateSensorTestTemplateV1)
}

// Alias Function
/*
This method acts as an alias for the method `DeploysTheICapConfigurationIntentByActivityIDV1`
*/
func (s *SensorsService) DeploysTheICapConfigurationIntentByActivityID(previewActivityID string, requestSensorsDeploysTheICAPConfigurationIntentByActivityIDV1 *RequestSensorsDeploysTheICapConfigurationIntentByActivityIDV1) (*ResponseSensorsDeploysTheICapConfigurationIntentByActivityIDV1, *resty.Response, error) {
	return s.DeploysTheICapConfigurationIntentByActivityIDV1(previewActivityID, requestSensorsDeploysTheICAPConfigurationIntentByActivityIDV1)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesDeployedICapConfigurationsWhileSupportingBasicFilteringV1`
*/
func (s *SensorsService) RetrievesDeployedICapConfigurationsWhileSupportingBasicFiltering(RetrievesDeployedICAPConfigurationsWhileSupportingBasicFilteringV1QueryParams *RetrievesDeployedICapConfigurationsWhileSupportingBasicFilteringV1QueryParams) (*ResponseSensorsRetrievesDeployedICapConfigurationsWhileSupportingBasicFilteringV1, *resty.Response, error) {
	return s.RetrievesDeployedICapConfigurationsWhileSupportingBasicFilteringV1(RetrievesDeployedICAPConfigurationsWhileSupportingBasicFilteringV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesTheCountOfDeployedICapConfigurationsWhileSupportingBasicFilteringV1`
*/
func (s *SensorsService) RetrievesTheCountOfDeployedICapConfigurationsWhileSupportingBasicFiltering(RetrievesTheCountOfDeployedICAPConfigurationsWhileSupportingBasicFilteringV1QueryParams *RetrievesTheCountOfDeployedICapConfigurationsWhileSupportingBasicFilteringV1QueryParams) (*ResponseSensorsRetrievesTheCountOfDeployedICapConfigurationsWhileSupportingBasicFilteringV1, *resty.Response, error) {
	return s.RetrievesTheCountOfDeployedICapConfigurationsWhileSupportingBasicFilteringV1(RetrievesTheCountOfDeployedICAPConfigurationsWhileSupportingBasicFilteringV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `CreatesAnICapConfigurationIntentForPreviewApproveV1`
*/
func (s *SensorsService) CreatesAnICapConfigurationIntentForPreviewApprove(requestSensorsCreatesAnICAPConfigurationIntentForPreviewApproveV1 *RequestSensorsCreatesAnICapConfigurationIntentForPreviewApproveV1, CreatesAnICAPConfigurationIntentForPreviewApproveV1QueryParams *CreatesAnICapConfigurationIntentForPreviewApproveV1QueryParams) (*ResponseSensorsCreatesAnICapConfigurationIntentForPreviewApproveV1, *resty.Response, error) {
	return s.CreatesAnICapConfigurationIntentForPreviewApproveV1(requestSensorsCreatesAnICAPConfigurationIntentForPreviewApproveV1, CreatesAnICAPConfigurationIntentForPreviewApproveV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1`
*/
func (s *SensorsService) RetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTime(id string, requestSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1 *RequestSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1, RetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1HeaderParams *RetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1HeaderParams) (*ResponseSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1, *resty.Response, error) {
	return s.RetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1(id, requestSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1, RetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMacV1`
*/
func (s *SensorsService) RetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMac(RetrievesTheSpectrumSensorReportsSentByWLCForProvidedAPMacV1HeaderParams *RetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMacV1HeaderParams, RetrievesTheSpectrumSensorReportsSentByWLCForProvidedAPMacV1QueryParams *RetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMacV1QueryParams) (*ResponseSensorsRetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMacV1, *resty.Response, error) {
	return s.RetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMacV1(RetrievesTheSpectrumSensorReportsSentByWLCForProvidedAPMacV1HeaderParams, RetrievesTheSpectrumSensorReportsSentByWLCForProvidedAPMacV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `SensorsV1`
*/
func (s *SensorsService) Sensors(SensorsV1QueryParams *SensorsV1QueryParams) (*ResponseSensorsSensorsV1, *resty.Response, error) {
	return s.SensorsV1(SensorsV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesDetailsOfASpecificICapPacketCaptureFileV1`
*/
func (s *SensorsService) RetrievesDetailsOfASpecificICapPacketCaptureFile(id string, RetrievesDetailsOfASpecificICAPPacketCaptureFileV1HeaderParams *RetrievesDetailsOfASpecificICapPacketCaptureFileV1HeaderParams) (*ResponseSensorsRetrievesDetailsOfASpecificICapPacketCaptureFileV1, *resty.Response, error) {
	return s.RetrievesDetailsOfASpecificICapPacketCaptureFileV1(id, RetrievesDetailsOfASpecificICAPPacketCaptureFileV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `RemoveTheICapConfigurationOnTheDeviceWithoutPreviewV1`
*/
func (s *SensorsService) RemoveTheICapConfigurationOnTheDeviceWithoutPreview(id string, requestSensorsRemoveTheICAPConfigurationOnTheDeviceWithoutPreviewV1 *RequestSensorsRemoveTheICapConfigurationOnTheDeviceWithoutPreviewV1) (*ResponseSensorsRemoveTheICapConfigurationOnTheDeviceWithoutPreviewV1, *resty.Response, error) {
	return s.RemoveTheICapConfigurationOnTheDeviceWithoutPreviewV1(id, requestSensorsRemoveTheICAPConfigurationOnTheDeviceWithoutPreviewV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetICapConfigurationStatusPerNetworkDeviceV1`
*/
func (s *SensorsService) GetICapConfigurationStatusPerNetworkDevice(previewActivityID string) (*ResponseSensorsGetICapConfigurationStatusPerNetworkDeviceV1, *resty.Response, error) {
	return s.GetICapConfigurationStatusPerNetworkDeviceV1(previewActivityID)
}

// Alias Function
/*
This method acts as an alias for the method `DeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1`
*/
func (s *SensorsService) DeploysTheGivenICapConfigurationIntentWithoutPreviewAndApprove(requestSensorsDeploysTheGivenICAPConfigurationIntentWithoutPreviewAndApproveV1 *RequestSensorsDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1, DeploysTheGivenICAPConfigurationIntentWithoutPreviewAndApproveV1QueryParams *DeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1QueryParams) (*ResponseSensorsDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1, *resty.Response, error) {
	return s.DeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1(requestSensorsDeploysTheGivenICAPConfigurationIntentWithoutPreviewAndApproveV1, DeploysTheGivenICAPConfigurationIntentWithoutPreviewAndApproveV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMacV1`
*/
func (s *SensorsService) RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMac(RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedAPMacV1HeaderParams *RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMacV1HeaderParams, RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedAPMacV1QueryParams *RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMacV1QueryParams) (*ResponseSensorsRetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMacV1, *resty.Response, error) {
	return s.RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMacV1(RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedAPMacV1HeaderParams, RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedAPMacV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `DownloadsASpecificICapPacketCaptureFileV1`
*/
func (s *SensorsService) DownloadsASpecificICapPacketCaptureFile(id string, DownloadsASpecificICAPPacketCaptureFileV1HeaderParams *DownloadsASpecificICapPacketCaptureFileV1HeaderParams) (FileDownload, *resty.Response, error) {
	return s.DownloadsASpecificICapPacketCaptureFileV1(id, DownloadsASpecificICAPPacketCaptureFileV1HeaderParams)
}

// Alias Function
/*
This method acts as an alias for the method `EditSensorTestTemplateV1`
*/
func (s *SensorsService) EditSensorTestTemplate(requestSensorsEditSensorTestTemplateV1 *RequestSensorsEditSensorTestTemplateV1) (*ResponseSensorsEditSensorTestTemplateV1, *resty.Response, error) {
	return s.EditSensorTestTemplateV1(requestSensorsEditSensorTestTemplateV1)
}
